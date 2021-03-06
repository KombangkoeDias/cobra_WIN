// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package systemreward

import (
	"math/big"
	"strings"

	ethereum "github.com/ethereum/go-ethereum"
	"github.com/ethereum/go-ethereum/accounts/abi"
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/event"
)

// Reference imports to suppress errors if they are not otherwise used.
var (
	_ = big.NewInt
	_ = strings.NewReader
	_ = ethereum.NotFound
	_ = bind.Bind
	_ = common.Big1
	_ = types.BloomLookup
	_ = event.NewSubscription
)

// SystemrewardABI is the input ABI used to generate the binding from.
const SystemrewardABI = "[{\"inputs\":[{\"internalType\":\"address\",\"name\":\"consensusAddress\",\"type\":\"address\"}],\"name\":\"AllocateJailValidatorReward\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"consensusAddress\",\"type\":\"address\"}],\"name\":\"addReward\",\"outputs\":[],\"stateMutability\":\"payable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"alreadyInit\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"distributeReward\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"fund\",\"outputs\":[],\"stateMutability\":\"payable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"getBalance\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"consensusAddress\",\"type\":\"address\"}],\"name\":\"getTotalPower\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"consensusAddress\",\"type\":\"address\"}],\"name\":\"getTotalPowerExcludeUnbonding\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"BKCValidatorSetAddress\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"StakePoolAddress\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"ValidatorPoolAddress\",\"type\":\"address\"}],\"name\":\"init\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"name\":\"rewardMapping\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"}]"

// SystemrewardBin is the compiled bytecode used for deploying new contracts.
var SystemrewardBin = "0x6080604052600a60015534801561001557600080fd5b506120c6806100256000396000f3fe6080604052600436106100915760003560e01c80638f73c5ae116100595780638f73c5ae1461018d5780639c9b2e21146101a4578063a78abc16146101c0578063b60d4288146101eb578063d7dc81cd146101f557610091565b806312065fe014610096578063184b9559146100c1578063322923b9146100ea578063348749e014610127578063571a915414610150575b600080fd5b3480156100a257600080fd5b506100ab610232565b6040516100b89190611b78565b60405180910390f35b3480156100cd57600080fd5b506100e860048036038101906100e3919061182a565b610288565b005b3480156100f657600080fd5b50610111600480360381019061010c91906117fd565b61047b565b60405161011e9190611b78565b60405180910390f35b34801561013357600080fd5b5061014e600480360381019061014991906117fd565b61052f565b005b34801561015c57600080fd5b50610177600480360381019061017291906117fd565b610a3d565b6040516101849190611b78565b60405180910390f35b34801561019957600080fd5b506101a2610a55565b005b6101be60048036038101906101b991906117fd565b610f3f565b005b3480156101cc57600080fd5b506101d5611130565b6040516101e29190611a9d565b60405180910390f35b6101f3611141565b005b34801561020157600080fd5b5061021c600480360381019061021791906117fd565b611190565b6040516102299190611b78565b60405180910390f35b60008060009054906101000a900460ff16610282576040517f08c379a000000000000000000000000000000000000000000000000000000000815260040161027990611ab8565b60405180910390fd5b47905090565b60008054906101000a900460ff16156102d6576040517f08c379a00000000000000000000000000000000000000000000000000000000081526004016102cd90611b18565b60405180910390fd5b82600360006101000a81548173ffffffffffffffffffffffffffffffffffffffff021916908373ffffffffffffffffffffffffffffffffffffffff16021790555081600460006101000a81548173ffffffffffffffffffffffffffffffffffffffff021916908373ffffffffffffffffffffffffffffffffffffffff16021790555080600560006101000a81548173ffffffffffffffffffffffffffffffffffffffff021916908373ffffffffffffffffffffffffffffffffffffffff16021790555082600660006101000a81548173ffffffffffffffffffffffffffffffffffffffff021916908373ffffffffffffffffffffffffffffffffffffffff16021790555081600760006101000a81548173ffffffffffffffffffffffffffffffffffffffff021916908373ffffffffffffffffffffffffffffffffffffffff16021790555080600860006101000a81548173ffffffffffffffffffffffffffffffffffffffff021916908373ffffffffffffffffffffffffffffffffffffffff16021790555060016000806101000a81548160ff021916908315150217905550505050565b6000600860009054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1663322923b9836040518263ffffffff1660e01b81526004016104d89190611a59565b60206040518083038186803b1580156104f057600080fd5b505afa158015610504573d6000803e3d6000fd5b505050506040513d601f19601f82011682018060405250810190610528919061192d565b9050919050565b60008054906101000a900460ff1661057c576040517f08c379a000000000000000000000000000000000000000000000000000000000815260040161057390611ab8565b60405180910390fd5b600360009054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff163373ffffffffffffffffffffffffffffffffffffffff161461060c576040517f08c379a000000000000000000000000000000000000000000000000000000000815260040161060390611ad8565b60405180910390fd5b6000600260008373ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1681526020019081526020016000205490506000805b600660009054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff166374de16696040518163ffffffff1660e01b815260040160206040518083038186803b1580156106bc57600080fd5b505afa1580156106d0573d6000803e3d6000fd5b505050506040513d601f19601f820116820180604052508101906106f4919061192d565b8110156107e657600080600660009054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16636969a25c846040518263ffffffff1660e01b81526004016107599190611b78565b60806040518083038186803b15801561077157600080fd5b505afa158015610785573d6000803e3d6000fd5b505050506040513d601f19601f820116820180604052508101906107a9919061187d565b935050509150806107d15760006107bf8361047b565b905080856107cd9190611bf5565b9450505b505080806107de90611d83565b915050610654565b506064605a60ff16836107f99190611c7c565b6108039190611c4b565b915060005b600660009054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff166374de16696040518163ffffffff1660e01b815260040160206040518083038186803b15801561087057600080fd5b505afa158015610884573d6000803e3d6000fd5b505050506040513d601f19601f820116820180604052508101906108a8919061192d565b8110156109f257600080600660009054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16636969a25c846040518263ffffffff1660e01b815260040161090d9190611b78565b60806040518083038186803b15801561092557600080fd5b505afa158015610939573d6000803e3d6000fd5b505050506040513d601f19601f8201168201806040525081019061095d919061187d565b935050509150806109dd57836109728361047b565b8661097d9190611c7c565b6109879190611c4b565b600260008473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200190815260200160002060008282546109d59190611bf5565b925050819055505b505080806109ea90611d83565b915050610808565b506000600260008573ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200190815260200160002081905550505050565b60026020528060005260406000206000915090505481565b60008054906101000a900460ff16610aa2576040517f08c379a0000000000000000000000000000000000000000000000000000000008152600401610a9990611ab8565b60405180910390fd5b600360009054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff163373ffffffffffffffffffffffffffffffffffffffff1614610b32576040517f08c379a0000000000000000000000000000000000000000000000000000000008152600401610b2990611ad8565b60405180910390fd5b6000805b600660009054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff166374de16696040518163ffffffff1660e01b815260040160206040518083038186803b158015610b9e57600080fd5b505afa158015610bb2573d6000803e3d6000fd5b505050506040513d601f19601f82011682018060405250810190610bd6919061192d565b811015610cf0576000600660009054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16636969a25c836040518263ffffffff1660e01b8152600401610c3a9190611b78565b60806040518083038186803b158015610c5257600080fd5b505afa158015610c66573d6000803e3d6000fd5b505050506040513d601f19601f82011682018060405250810190610c8a919061187d565b5050509050600260008273ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1681526020019081526020016000205483610cda9190611bf5565b9250508080610ce890611d83565b915050610b36565b5080471015610d34576040517f08c379a0000000000000000000000000000000000000000000000000000000008152600401610d2b90611b38565b60405180910390fd5b60005b600660009054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff166374de16696040518163ffffffff1660e01b815260040160206040518083038186803b158015610d9f57600080fd5b505afa158015610db3573d6000803e3d6000fd5b505050506040513d601f19601f82011682018060405250810190610dd7919061192d565b811015610f3b576000600660009054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16636969a25c836040518263ffffffff1660e01b8152600401610e3b9190611b78565b60806040518083038186803b158015610e5357600080fd5b505afa158015610e67573d6000803e3d6000fd5b505050506040513d601f19601f82011682018060405250810190610e8b919061187d565b50505090506000600260008373ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff168152602001908152602001600020541115610f2757610ee181611244565b6000600260008373ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff168152602001908152602001600020819055505b508080610f3390611d83565b915050610d37565b5050565b60008054906101000a900460ff16610f8c576040517f08c379a0000000000000000000000000000000000000000000000000000000008152600401610f8390611ab8565b60405180910390fd5b6000600660009054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1663ad3c9da6836040518263ffffffff1660e01b8152600401610fe99190611a59565b60206040518083038186803b15801561100157600080fd5b505afa158015611015573d6000803e3d6000fd5b505050506040513d601f19601f82011682018060405250810190611039919061192d565b11611079576040517f08c379a000000000000000000000000000000000000000000000000000000000815260040161107090611af8565b60405180910390fd5b600034116110bc576040517f08c379a00000000000000000000000000000000000000000000000000000000081526004016110b390611b58565b60405180910390fd5b6064605a60ff16346110ce9190611c7c565b6110d89190611c4b565b600260008373ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200190815260200160002060008282546111269190611bf5565b9250508190555050565b60008054906101000a900460ff1681565b60008054906101000a900460ff1661118e576040517f08c379a000000000000000000000000000000000000000000000000000000000815260040161118590611ab8565b60405180910390fd5b565b6000600860009054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1663d7dc81cd836040518263ffffffff1660e01b81526004016111ed9190611a59565b60206040518083038186803b15801561120557600080fd5b505afa158015611219573d6000803e3d6000fd5b505050506040513d601f19601f8201168201806040525081019061123d919061192d565b9050919050565b60008054906101000a900460ff16611291576040517f08c379a000000000000000000000000000000000000000000000000000000000815260040161128890611ab8565b60405180910390fd5b6000600660009054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16636969a25c6001600660009054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1663ad3c9da6866040518263ffffffff1660e01b815260040161132e9190611a59565b60206040518083038186803b15801561134657600080fd5b505afa15801561135a573d6000803e3d6000fd5b505050506040513d601f19601f8201168201806040525081019061137e919061192d565b6113889190611cd6565b6040518263ffffffff1660e01b81526004016113a49190611b78565b60806040518083038186803b1580156113bc57600080fd5b505afa1580156113d0573d6000803e3d6000fd5b505050506040513d601f19601f820116820180604052508101906113f4919061187d565b505091505060006114048361047b565b9050600081600260008673ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200190815260200160002054846114549190611c7c565b61145e9190611c4b565b90508373ffffffffffffffffffffffffffffffffffffffff166108fc829081150290604051600060405180830381858888f193505050501580156114a6573d6000803e3d6000fd5b506000600760009054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff166368e76346866040518263ffffffff1660e01b81526004016115049190611a59565b60006040518083038186803b15801561151c57600080fd5b505afa158015611530573d6000803e3d6000fd5b505050506040513d6000823e3d601f19601f8201168201806040525081019061155991906118e4565b905060005b81518110156116ee57600082828151811061157c5761157b611e2a565b5b602002602001015190506000600760009054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1663f9353fe0838a6040518363ffffffff1660e01b81526004016115e5929190611a74565b60206040518083038186803b1580156115fd57600080fd5b505afa158015611611573d6000803e3d6000fd5b505050506040513d601f19601f82011682018060405250810190611635919061192d565b9050600086600260008b73ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200190815260200160002054836116859190611c7c565b61168f9190611c4b565b90508273ffffffffffffffffffffffffffffffffffffffff166108fc829081150290604051600060405180830381858888f193505050501580156116d7573d6000803e3d6000fd5b5050505080806116e690611d83565b91505061155e565b505050505050565b600061170961170484611bb8565b611b93565b9050808382526020820190508285602086028201111561172c5761172b611e8d565b5b60005b8581101561175c5781611742888261177b565b84526020840193506020830192505060018101905061172f565b5050509392505050565b6000813590506117758161203b565b92915050565b60008151905061178a8161203b565b92915050565b600082601f8301126117a5576117a4611e88565b5b81516117b58482602086016116f6565b91505092915050565b6000815190506117cd81612052565b92915050565b6000815190506117e281612069565b92915050565b6000815190506117f781612079565b92915050565b60006020828403121561181357611812611e97565b5b600061182184828501611766565b91505092915050565b60008060006060848603121561184357611842611e97565b5b600061185186828701611766565b935050602061186286828701611766565b925050604061187386828701611766565b9150509250925092565b6000806000806080858703121561189757611896611e97565b5b60006118a58782880161177b565b94505060206118b6878288016117e8565b93505060406118c7878288016117d3565b92505060606118d8878288016117be565b91505092959194509250565b6000602082840312156118fa576118f9611e97565b5b600082015167ffffffffffffffff81111561191857611917611e92565b5b61192484828501611790565b91505092915050565b60006020828403121561194357611942611e97565b5b6000611951848285016117e8565b91505092915050565b61196381611d0a565b82525050565b61197281611d1c565b82525050565b6000611985601983611be4565b915061199082611ead565b602082019050919050565b60006119a8603283611be4565b91506119b382611ed6565b604082019050919050565b60006119cb602a83611be4565b91506119d682611f25565b604082019050919050565b60006119ee601983611be4565b91506119f982611f74565b602082019050919050565b6000611a11604b83611be4565b9150611a1c82611f9d565b606082019050919050565b6000611a34601d83611be4565b9150611a3f82612012565b602082019050919050565b611a5381611d48565b82525050565b6000602082019050611a6e600083018461195a565b92915050565b6000604082019050611a89600083018561195a565b611a96602083018461195a565b9392505050565b6000602082019050611ab26000830184611969565b92915050565b60006020820190508181036000830152611ad181611978565b9050919050565b60006020820190508181036000830152611af18161199b565b9050919050565b60006020820190508181036000830152611b11816119be565b9050919050565b60006020820190508181036000830152611b31816119e1565b9050919050565b60006020820190508181036000830152611b5181611a04565b9050919050565b60006020820190508181036000830152611b7181611a27565b9050919050565b6000602082019050611b8d6000830184611a4a565b92915050565b6000611b9d611bae565b9050611ba98282611d52565b919050565b6000604051905090565b600067ffffffffffffffff821115611bd357611bd2611e59565b5b602082029050602081019050919050565b600082825260208201905092915050565b6000611c0082611d48565b9150611c0b83611d48565b9250827fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff03821115611c4057611c3f611dcc565b5b828201905092915050565b6000611c5682611d48565b9150611c6183611d48565b925082611c7157611c70611dfb565b5b828204905092915050565b6000611c8782611d48565b9150611c9283611d48565b9250817fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff0483118215151615611ccb57611cca611dcc565b5b828202905092915050565b6000611ce182611d48565b9150611cec83611d48565b925082821015611cff57611cfe611dcc565b5b828203905092915050565b6000611d1582611d28565b9050919050565b60008115159050919050565b600073ffffffffffffffffffffffffffffffffffffffff82169050919050565b6000819050919050565b611d5b82611e9c565b810181811067ffffffffffffffff82111715611d7a57611d79611e59565b5b80604052505050565b6000611d8e82611d48565b91507fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff821415611dc157611dc0611dcc565b5b600182019050919050565b7f4e487b7100000000000000000000000000000000000000000000000000000000600052601160045260246000fd5b7f4e487b7100000000000000000000000000000000000000000000000000000000600052601260045260246000fd5b7f4e487b7100000000000000000000000000000000000000000000000000000000600052603260045260246000fd5b7f4e487b7100000000000000000000000000000000000000000000000000000000600052604160045260246000fd5b600080fd5b600080fd5b600080fd5b600080fd5b6000601f19601f8301169050919050565b7f74686520636f6e7472616374206e6f7420696e69742079657400000000000000600082015250565b7f63616e206f6e6c792062652063616c6c65642066726f6d207468652076616c6960008201527f6461746f722073657420636f6e74726163740000000000000000000000000000602082015250565b7f63616e2774206164642072657761726420746f2061206e6f6e2d61637469766560008201527f2076616c696461746f7200000000000000000000000000000000000000000000602082015250565b7f74686520636f6e747261637420616c726561647920696e697400000000000000600082015250565b7f63616e27742064697374726962757465207265776172642c206e6f7420656e6f60008201527f7567682066756e6420696e20746865206163636f756e742c20666f7220736f6d60208201527f6520726561736f6e2e2e2e000000000000000000000000000000000000000000604082015250565b7f63616e27742061646420726577617264207769746820302076616c7565000000600082015250565b61204481611d0a565b811461204f57600080fd5b50565b61205b81611d1c565b811461206657600080fd5b50565b6003811061207657600080fd5b50565b61208281611d48565b811461208d57600080fd5b5056fea264697066735822122074b8701f80c13856cfd3d6ceb60ec839b6c90c6521d4e1c43f2852645d619ba964736f6c63430008060033"

// DeploySystemreward deploys a new Ethereum contract, binding an instance of Systemreward to it.
func DeploySystemreward(auth *bind.TransactOpts, backend bind.ContractBackend) (common.Address, *types.Transaction, *Systemreward, error) {
	parsed, err := abi.JSON(strings.NewReader(SystemrewardABI))
	if err != nil {
		return common.Address{}, nil, nil, err
	}

	address, tx, contract, err := bind.DeployContract(auth, parsed, common.FromHex(SystemrewardBin), backend)
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	return address, tx, &Systemreward{SystemrewardCaller: SystemrewardCaller{contract: contract}, SystemrewardTransactor: SystemrewardTransactor{contract: contract}, SystemrewardFilterer: SystemrewardFilterer{contract: contract}}, nil
}

// Systemreward is an auto generated Go binding around an Ethereum contract.
type Systemreward struct {
	SystemrewardCaller     // Read-only binding to the contract
	SystemrewardTransactor // Write-only binding to the contract
	SystemrewardFilterer   // Log filterer for contract events
}

// SystemrewardCaller is an auto generated read-only Go binding around an Ethereum contract.
type SystemrewardCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// SystemrewardTransactor is an auto generated write-only Go binding around an Ethereum contract.
type SystemrewardTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// SystemrewardFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type SystemrewardFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// SystemrewardSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type SystemrewardSession struct {
	Contract     *Systemreward     // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// SystemrewardCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type SystemrewardCallerSession struct {
	Contract *SystemrewardCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts       // Call options to use throughout this session
}

// SystemrewardTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type SystemrewardTransactorSession struct {
	Contract     *SystemrewardTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts       // Transaction auth options to use throughout this session
}

// SystemrewardRaw is an auto generated low-level Go binding around an Ethereum contract.
type SystemrewardRaw struct {
	Contract *Systemreward // Generic contract binding to access the raw methods on
}

// SystemrewardCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type SystemrewardCallerRaw struct {
	Contract *SystemrewardCaller // Generic read-only contract binding to access the raw methods on
}

// SystemrewardTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type SystemrewardTransactorRaw struct {
	Contract *SystemrewardTransactor // Generic write-only contract binding to access the raw methods on
}

// NewSystemreward creates a new instance of Systemreward, bound to a specific deployed contract.
func NewSystemreward(address common.Address, backend bind.ContractBackend) (*Systemreward, error) {
	contract, err := bindSystemreward(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &Systemreward{SystemrewardCaller: SystemrewardCaller{contract: contract}, SystemrewardTransactor: SystemrewardTransactor{contract: contract}, SystemrewardFilterer: SystemrewardFilterer{contract: contract}}, nil
}

// NewSystemrewardCaller creates a new read-only instance of Systemreward, bound to a specific deployed contract.
func NewSystemrewardCaller(address common.Address, caller bind.ContractCaller) (*SystemrewardCaller, error) {
	contract, err := bindSystemreward(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &SystemrewardCaller{contract: contract}, nil
}

// NewSystemrewardTransactor creates a new write-only instance of Systemreward, bound to a specific deployed contract.
func NewSystemrewardTransactor(address common.Address, transactor bind.ContractTransactor) (*SystemrewardTransactor, error) {
	contract, err := bindSystemreward(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &SystemrewardTransactor{contract: contract}, nil
}

// NewSystemrewardFilterer creates a new log filterer instance of Systemreward, bound to a specific deployed contract.
func NewSystemrewardFilterer(address common.Address, filterer bind.ContractFilterer) (*SystemrewardFilterer, error) {
	contract, err := bindSystemreward(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &SystemrewardFilterer{contract: contract}, nil
}

// bindSystemreward binds a generic wrapper to an already deployed contract.
func bindSystemreward(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(SystemrewardABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Systemreward *SystemrewardRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _Systemreward.Contract.SystemrewardCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Systemreward *SystemrewardRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Systemreward.Contract.SystemrewardTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Systemreward *SystemrewardRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Systemreward.Contract.SystemrewardTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Systemreward *SystemrewardCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _Systemreward.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Systemreward *SystemrewardTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Systemreward.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Systemreward *SystemrewardTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Systemreward.Contract.contract.Transact(opts, method, params...)
}

// AlreadyInit is a free data retrieval call binding the contract method 0xa78abc16.
//
// Solidity: function alreadyInit() view returns(bool)
func (_Systemreward *SystemrewardCaller) AlreadyInit(opts *bind.CallOpts) (bool, error) {
	var out []interface{}
	err := _Systemreward.contract.Call(opts, &out, "alreadyInit")

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// AlreadyInit is a free data retrieval call binding the contract method 0xa78abc16.
//
// Solidity: function alreadyInit() view returns(bool)
func (_Systemreward *SystemrewardSession) AlreadyInit() (bool, error) {
	return _Systemreward.Contract.AlreadyInit(&_Systemreward.CallOpts)
}

// AlreadyInit is a free data retrieval call binding the contract method 0xa78abc16.
//
// Solidity: function alreadyInit() view returns(bool)
func (_Systemreward *SystemrewardCallerSession) AlreadyInit() (bool, error) {
	return _Systemreward.Contract.AlreadyInit(&_Systemreward.CallOpts)
}

// GetBalance is a free data retrieval call binding the contract method 0x12065fe0.
//
// Solidity: function getBalance() view returns(uint256)
func (_Systemreward *SystemrewardCaller) GetBalance(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _Systemreward.contract.Call(opts, &out, "getBalance")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// GetBalance is a free data retrieval call binding the contract method 0x12065fe0.
//
// Solidity: function getBalance() view returns(uint256)
func (_Systemreward *SystemrewardSession) GetBalance() (*big.Int, error) {
	return _Systemreward.Contract.GetBalance(&_Systemreward.CallOpts)
}

// GetBalance is a free data retrieval call binding the contract method 0x12065fe0.
//
// Solidity: function getBalance() view returns(uint256)
func (_Systemreward *SystemrewardCallerSession) GetBalance() (*big.Int, error) {
	return _Systemreward.Contract.GetBalance(&_Systemreward.CallOpts)
}

// GetTotalPower is a free data retrieval call binding the contract method 0xd7dc81cd.
//
// Solidity: function getTotalPower(address consensusAddress) view returns(uint256)
func (_Systemreward *SystemrewardCaller) GetTotalPower(opts *bind.CallOpts, consensusAddress common.Address) (*big.Int, error) {
	var out []interface{}
	err := _Systemreward.contract.Call(opts, &out, "getTotalPower", consensusAddress)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// GetTotalPower is a free data retrieval call binding the contract method 0xd7dc81cd.
//
// Solidity: function getTotalPower(address consensusAddress) view returns(uint256)
func (_Systemreward *SystemrewardSession) GetTotalPower(consensusAddress common.Address) (*big.Int, error) {
	return _Systemreward.Contract.GetTotalPower(&_Systemreward.CallOpts, consensusAddress)
}

// GetTotalPower is a free data retrieval call binding the contract method 0xd7dc81cd.
//
// Solidity: function getTotalPower(address consensusAddress) view returns(uint256)
func (_Systemreward *SystemrewardCallerSession) GetTotalPower(consensusAddress common.Address) (*big.Int, error) {
	return _Systemreward.Contract.GetTotalPower(&_Systemreward.CallOpts, consensusAddress)
}

// GetTotalPowerExcludeUnbonding is a free data retrieval call binding the contract method 0x322923b9.
//
// Solidity: function getTotalPowerExcludeUnbonding(address consensusAddress) view returns(uint256)
func (_Systemreward *SystemrewardCaller) GetTotalPowerExcludeUnbonding(opts *bind.CallOpts, consensusAddress common.Address) (*big.Int, error) {
	var out []interface{}
	err := _Systemreward.contract.Call(opts, &out, "getTotalPowerExcludeUnbonding", consensusAddress)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// GetTotalPowerExcludeUnbonding is a free data retrieval call binding the contract method 0x322923b9.
//
// Solidity: function getTotalPowerExcludeUnbonding(address consensusAddress) view returns(uint256)
func (_Systemreward *SystemrewardSession) GetTotalPowerExcludeUnbonding(consensusAddress common.Address) (*big.Int, error) {
	return _Systemreward.Contract.GetTotalPowerExcludeUnbonding(&_Systemreward.CallOpts, consensusAddress)
}

// GetTotalPowerExcludeUnbonding is a free data retrieval call binding the contract method 0x322923b9.
//
// Solidity: function getTotalPowerExcludeUnbonding(address consensusAddress) view returns(uint256)
func (_Systemreward *SystemrewardCallerSession) GetTotalPowerExcludeUnbonding(consensusAddress common.Address) (*big.Int, error) {
	return _Systemreward.Contract.GetTotalPowerExcludeUnbonding(&_Systemreward.CallOpts, consensusAddress)
}

// RewardMapping is a free data retrieval call binding the contract method 0x571a9154.
//
// Solidity: function rewardMapping(address ) view returns(uint256)
func (_Systemreward *SystemrewardCaller) RewardMapping(opts *bind.CallOpts, arg0 common.Address) (*big.Int, error) {
	var out []interface{}
	err := _Systemreward.contract.Call(opts, &out, "rewardMapping", arg0)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// RewardMapping is a free data retrieval call binding the contract method 0x571a9154.
//
// Solidity: function rewardMapping(address ) view returns(uint256)
func (_Systemreward *SystemrewardSession) RewardMapping(arg0 common.Address) (*big.Int, error) {
	return _Systemreward.Contract.RewardMapping(&_Systemreward.CallOpts, arg0)
}

// RewardMapping is a free data retrieval call binding the contract method 0x571a9154.
//
// Solidity: function rewardMapping(address ) view returns(uint256)
func (_Systemreward *SystemrewardCallerSession) RewardMapping(arg0 common.Address) (*big.Int, error) {
	return _Systemreward.Contract.RewardMapping(&_Systemreward.CallOpts, arg0)
}

// AllocateJailValidatorReward is a paid mutator transaction binding the contract method 0x348749e0.
//
// Solidity: function AllocateJailValidatorReward(address consensusAddress) returns()
func (_Systemreward *SystemrewardTransactor) AllocateJailValidatorReward(opts *bind.TransactOpts, consensusAddress common.Address) (*types.Transaction, error) {
	return _Systemreward.contract.Transact(opts, "AllocateJailValidatorReward", consensusAddress)
}

// AllocateJailValidatorReward is a paid mutator transaction binding the contract method 0x348749e0.
//
// Solidity: function AllocateJailValidatorReward(address consensusAddress) returns()
func (_Systemreward *SystemrewardSession) AllocateJailValidatorReward(consensusAddress common.Address) (*types.Transaction, error) {
	return _Systemreward.Contract.AllocateJailValidatorReward(&_Systemreward.TransactOpts, consensusAddress)
}

// AllocateJailValidatorReward is a paid mutator transaction binding the contract method 0x348749e0.
//
// Solidity: function AllocateJailValidatorReward(address consensusAddress) returns()
func (_Systemreward *SystemrewardTransactorSession) AllocateJailValidatorReward(consensusAddress common.Address) (*types.Transaction, error) {
	return _Systemreward.Contract.AllocateJailValidatorReward(&_Systemreward.TransactOpts, consensusAddress)
}

// AddReward is a paid mutator transaction binding the contract method 0x9c9b2e21.
//
// Solidity: function addReward(address consensusAddress) payable returns()
func (_Systemreward *SystemrewardTransactor) AddReward(opts *bind.TransactOpts, consensusAddress common.Address) (*types.Transaction, error) {
	return _Systemreward.contract.Transact(opts, "addReward", consensusAddress)
}

// AddReward is a paid mutator transaction binding the contract method 0x9c9b2e21.
//
// Solidity: function addReward(address consensusAddress) payable returns()
func (_Systemreward *SystemrewardSession) AddReward(consensusAddress common.Address) (*types.Transaction, error) {
	return _Systemreward.Contract.AddReward(&_Systemreward.TransactOpts, consensusAddress)
}

// AddReward is a paid mutator transaction binding the contract method 0x9c9b2e21.
//
// Solidity: function addReward(address consensusAddress) payable returns()
func (_Systemreward *SystemrewardTransactorSession) AddReward(consensusAddress common.Address) (*types.Transaction, error) {
	return _Systemreward.Contract.AddReward(&_Systemreward.TransactOpts, consensusAddress)
}

// DistributeReward is a paid mutator transaction binding the contract method 0x8f73c5ae.
//
// Solidity: function distributeReward() returns()
func (_Systemreward *SystemrewardTransactor) DistributeReward(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Systemreward.contract.Transact(opts, "distributeReward")
}

// DistributeReward is a paid mutator transaction binding the contract method 0x8f73c5ae.
//
// Solidity: function distributeReward() returns()
func (_Systemreward *SystemrewardSession) DistributeReward() (*types.Transaction, error) {
	return _Systemreward.Contract.DistributeReward(&_Systemreward.TransactOpts)
}

// DistributeReward is a paid mutator transaction binding the contract method 0x8f73c5ae.
//
// Solidity: function distributeReward() returns()
func (_Systemreward *SystemrewardTransactorSession) DistributeReward() (*types.Transaction, error) {
	return _Systemreward.Contract.DistributeReward(&_Systemreward.TransactOpts)
}

// Fund is a paid mutator transaction binding the contract method 0xb60d4288.
//
// Solidity: function fund() payable returns()
func (_Systemreward *SystemrewardTransactor) Fund(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Systemreward.contract.Transact(opts, "fund")
}

// Fund is a paid mutator transaction binding the contract method 0xb60d4288.
//
// Solidity: function fund() payable returns()
func (_Systemreward *SystemrewardSession) Fund() (*types.Transaction, error) {
	return _Systemreward.Contract.Fund(&_Systemreward.TransactOpts)
}

// Fund is a paid mutator transaction binding the contract method 0xb60d4288.
//
// Solidity: function fund() payable returns()
func (_Systemreward *SystemrewardTransactorSession) Fund() (*types.Transaction, error) {
	return _Systemreward.Contract.Fund(&_Systemreward.TransactOpts)
}

// Init is a paid mutator transaction binding the contract method 0x184b9559.
//
// Solidity: function init(address BKCValidatorSetAddress, address StakePoolAddress, address ValidatorPoolAddress) returns()
func (_Systemreward *SystemrewardTransactor) Init(opts *bind.TransactOpts, BKCValidatorSetAddress common.Address, StakePoolAddress common.Address, ValidatorPoolAddress common.Address) (*types.Transaction, error) {
	return _Systemreward.contract.Transact(opts, "init", BKCValidatorSetAddress, StakePoolAddress, ValidatorPoolAddress)
}

// Init is a paid mutator transaction binding the contract method 0x184b9559.
//
// Solidity: function init(address BKCValidatorSetAddress, address StakePoolAddress, address ValidatorPoolAddress) returns()
func (_Systemreward *SystemrewardSession) Init(BKCValidatorSetAddress common.Address, StakePoolAddress common.Address, ValidatorPoolAddress common.Address) (*types.Transaction, error) {
	return _Systemreward.Contract.Init(&_Systemreward.TransactOpts, BKCValidatorSetAddress, StakePoolAddress, ValidatorPoolAddress)
}

// Init is a paid mutator transaction binding the contract method 0x184b9559.
//
// Solidity: function init(address BKCValidatorSetAddress, address StakePoolAddress, address ValidatorPoolAddress) returns()
func (_Systemreward *SystemrewardTransactorSession) Init(BKCValidatorSetAddress common.Address, StakePoolAddress common.Address, ValidatorPoolAddress common.Address) (*types.Transaction, error) {
	return _Systemreward.Contract.Init(&_Systemreward.TransactOpts, BKCValidatorSetAddress, StakePoolAddress, ValidatorPoolAddress)
}
