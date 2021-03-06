// SPDX-License-Identifier: MIT
pragma solidity <=0.8.6;

import "./System.sol";
import "./Interface/IValidator.sol";
import "./Interface/IDelegation.sol";
import "./BKCValidatorSet.sol";
import "./StakePool.sol";
import "./ValidatorPool.sol";

contract SystemReward is System, IValidator, IDelegation {
    mapping(address => uint256) public rewardMapping; // the mapping for reward of each validator.

    address private _BKCValidatorSetAddress;
    address private _StakePoolAddress;
    address private _ValidatorPoolAddress;

    uint8 constant PERCENTAGE_OF_REWARD_KEPT_FOR_MAINTENANCE = 90;

    BKCValidatorSet private validatorset;
    StakePool private stakepool;
    ValidatorPool private vldpool;

    modifier onlyBKCValidatorSetContract {
        require(
            msg.sender == _BKCValidatorSetAddress,
            "can only be called from the validator set contract"
        );
        _;
    }

    function init(
        address BKCValidatorSetAddress,
        address StakePoolAddress,
        address ValidatorPoolAddress
    ) external onlyNotInit {
        _BKCValidatorSetAddress = BKCValidatorSetAddress;
        _StakePoolAddress = StakePoolAddress;
        _ValidatorPoolAddress = ValidatorPoolAddress;
        validatorset = BKCValidatorSet(BKCValidatorSetAddress);
        stakepool = StakePool(StakePoolAddress);
        vldpool = ValidatorPool(ValidatorPoolAddress);
        // initialize for testing (need to be funded 3 ether)
        // rewardMapping[0x17F6AD8Ef982297579C203069C1DbfFE4348c372] =
        //     3 *
        //     (10**18);
        alreadyInit = true;
    }

    function getBalance() external view onlyInit returns (uint256) {
        return address(this).balance;
    }

    function fund() external payable onlyInit {
        // funding algo here
    }

    function addReward(address consensusAddress) external payable onlyInit {
        require(
            validatorset.currentValidatorSetMap(consensusAddress) > 0,
            "can't add reward to a non-active validator"
        );
        require(msg.value > 0, "can't add reward with 0 value");
        rewardMapping[consensusAddress] +=
            (msg.value * PERCENTAGE_OF_REWARD_KEPT_FOR_MAINTENANCE) /
            100; // the reward is kept at 10% for system maintenance.
    }

    function distributeReward() external onlyInit onlyBKCValidatorSetContract {
        uint256 totalRewardToBeDistributed = 0;

        for (uint256 i = 0; i < validatorset.number_of_validators(); i++) {
            (address a, , , ) = validatorset.currentValidatorSet(i);
            totalRewardToBeDistributed += rewardMapping[a];
        }

        require(
            address(this).balance >= totalRewardToBeDistributed,
            "can't distribute reward, not enough fund in the account, for some reason..."
        );

        for (uint256 i = 0; i < validatorset.number_of_validators(); i++) {
            (address a, , , ) = validatorset.currentValidatorSet(i);
            if (rewardMapping[a] > 0) {
                calculateAndTransferReward(a); // calculate and transfer reward for validator address a and its delegators.
                rewardMapping[a] = 0; // set the reward back to 0
            }
        }
    }

    function getTotalPower(address consensusAddress)
        public
        view
        returns (uint256)
    {
        return vldpool.getTotalPower(consensusAddress);
    }

    function getTotalPowerExcludeUnbonding(address consensusAddress)
        public
        view
        returns (uint256)
    {
        return vldpool.getTotalPowerExcludeUnbonding(consensusAddress);
    }

    function AllocateJailValidatorReward(address consensusAddress)
        external
        onlyInit
        onlyBKCValidatorSetContract
    {
        uint256 reward = rewardMapping[consensusAddress];
        uint256 totalPowerOfValidatorSet = 0;
        for (uint256 i = 0; i < validatorset.number_of_validators(); i++) {
            (address a, , , bool isJail) = validatorset.currentValidatorSet(i);
            if (!isJail) {
                uint256 power = getTotalPowerExcludeUnbonding(a);
                totalPowerOfValidatorSet += power;
            }
        }
        reward = (reward * PERCENTAGE_OF_REWARD_KEPT_FOR_MAINTENANCE) / 100;
        for (uint256 i = 0; i < validatorset.number_of_validators(); i++) {
            (address a, , , bool isJail) = validatorset.currentValidatorSet(i);
            if (!isJail) {
                rewardMapping[a] +=
                    (reward * getTotalPowerExcludeUnbonding(a)) /
                    totalPowerOfValidatorSet;
            }
        }
        rewardMapping[consensusAddress] = 0;
    }

    function calculateAndTransferReward(address consensusAddress)
        private
        onlyInit
    {
        (, uint256 stakeAmount, , ) = validatorset.currentValidatorSet(
            validatorset.currentValidatorSetMap(consensusAddress) - 1
        );
        uint256 totalPower = getTotalPowerExcludeUnbonding(consensusAddress);
        uint256 validatorGets = (stakeAmount *
            rewardMapping[consensusAddress]) / totalPower;
        payable(consensusAddress).transfer(validatorGets); // validator gets
        address[] memory delegators = stakepool.getDelegators(consensusAddress);
        for (uint256 i = 0; i < delegators.length; i++) {
            address delegator = delegators[i];
            uint256 delegateAmount = stakepool
            .getUserDelegationBondedAmountCallable(delegator, consensusAddress);
            uint256 delegatorGets = (delegateAmount *
                rewardMapping[consensusAddress]) / totalPower;
            payable(delegator).transfer(delegatorGets); // delegator gets
        }
    }
}
