// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package contracts

import (
	"errors"
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
	_ = errors.New
	_ = big.NewInt
	_ = strings.NewReader
	_ = ethereum.NotFound
	_ = bind.Bind
	_ = common.Big1
	_ = types.BloomLookup
	_ = event.NewSubscription
	_ = abi.ConvertType
)

// PackRegistryMetaData contains all meta data concerning the PackRegistry contract.
var PackRegistryMetaData = &bind.MetaData{
	ABI: "[{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_initialStock\",\"type\":\"uint256\"}],\"stateMutability\":\"nonpayable\",\"type\":\"constructor\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"string\",\"name\":\"cardId\",\"type\":\"string\"},{\"indexed\":false,\"internalType\":\"string\",\"name\":\"ownerId\",\"type\":\"string\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"timestamp\",\"type\":\"uint256\"}],\"name\":\"CardAssigned\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"string\",\"name\":\"cardId\",\"type\":\"string\"},{\"indexed\":false,\"internalType\":\"string\",\"name\":\"fromId\",\"type\":\"string\"},{\"indexed\":false,\"internalType\":\"string\",\"name\":\"toId\",\"type\":\"string\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"timestamp\",\"type\":\"uint256\"}],\"name\":\"CardTransferred\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"string\",\"name\":\"matchId\",\"type\":\"string\"},{\"indexed\":false,\"internalType\":\"string\",\"name\":\"winner\",\"type\":\"string\"},{\"indexed\":false,\"internalType\":\"string\",\"name\":\"loser\",\"type\":\"string\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"timestamp\",\"type\":\"uint256\"}],\"name\":\"MatchRecorded\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"newStock\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"string\",\"name\":\"reason\",\"type\":\"string\"}],\"name\":\"StockUpdated\",\"type\":\"event\"},{\"inputs\":[{\"internalType\":\"string\",\"name\":\"_playerId\",\"type\":\"string\"},{\"internalType\":\"string[]\",\"name\":\"_cardIds\",\"type\":\"string[]\"}],\"name\":\"assignCards\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"string\",\"name\":\"\",\"type\":\"string\"}],\"name\":\"cardOwner\",\"outputs\":[{\"internalType\":\"string\",\"name\":\"\",\"type\":\"string\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"decrementStock\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"string\",\"name\":\"_cardId\",\"type\":\"string\"}],\"name\":\"getCardOwner\",\"outputs\":[{\"internalType\":\"string\",\"name\":\"\",\"type\":\"string\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"getMatchCount\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"getStock\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"name\":\"matches\",\"outputs\":[{\"internalType\":\"string\",\"name\":\"matchId\",\"type\":\"string\"},{\"internalType\":\"string\",\"name\":\"winnerId\",\"type\":\"string\"},{\"internalType\":\"string\",\"name\":\"loserId\",\"type\":\"string\"},{\"internalType\":\"uint256\",\"name\":\"timestamp\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"string\",\"name\":\"_matchId\",\"type\":\"string\"},{\"internalType\":\"string\",\"name\":\"_winnerId\",\"type\":\"string\"},{\"internalType\":\"string\",\"name\":\"_loserId\",\"type\":\"string\"}],\"name\":\"recordMatch\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"totalPacks\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"string\",\"name\":\"_fromId\",\"type\":\"string\"},{\"internalType\":\"string\",\"name\":\"_toId\",\"type\":\"string\"},{\"internalType\":\"string\",\"name\":\"_cardId\",\"type\":\"string\"}],\"name\":\"transferCard\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"}]",
	Bin: "0x60806040523480156200001157600080fd5b506040516200156338038062001563833981810160405281019062000037919062000085565b8060008190555050620000b7565b600080fd5b6000819050919050565b6200005f816200004a565b81146200006b57600080fd5b50565b6000815190506200007f8162000054565b92915050565b6000602082840312156200009e576200009d62000045565b5b6000620000ae848285016200006e565b91505092915050565b61149c80620000c76000396000f3fe608060405234801561001057600080fd5b506004361061009d5760003560e01c80637281df12116100665780637281df1214610133578063888ad898146101515780638c4f7dae14610181578063a79e52671461019f578063e38ffcd8146101cf5761009d565b8062375a44146100a2578063118c54ec146100ac57806312636a62146100c85780634768d4ef146100e45780636d34550e14610117575b600080fd5b6100aa6101ed565b005b6100c660048036038101906100c19190610a10565b610285565b005b6100e260048036038101906100dd9190610a10565b61036e565b005b6100fe60048036038101906100f99190610aed565b61044a565b60405161010e9493929190610ba8565b60405180910390f35b610131600480360381019061012c9190610ce8565b610622565b005b61013b610734565b6040516101489190610d60565b60405180910390f35b61016b60048036038101906101669190610d7b565b61073a565b6040516101789190610dc4565b60405180910390f35b6101896107f0565b6040516101969190610d60565b60405180910390f35b6101b960048036038101906101b49190610d7b565b6107fd565b6040516101c69190610dc4565b60405180910390f35b6101d76108ad565b6040516101e49190610d60565b60405180910390f35b6000805411610231576040517f08c379a000000000000000000000000000000000000000000000000000000000815260040161022890610e32565b60405180910390fd5b60016000808282546102439190610e81565b925050819055507f2926060a4c291371fea4258fa171860291f0144cd93a2accc8a0da15fc05c7a060005460405161027b9190610f01565b60405180910390a1565b828051906020012060028260405161029d9190610f6b565b90815260200160405180910390206040516102b89190611085565b604051809103902014610300576040517f08c379a00000000000000000000000000000000000000000000000000000000081526004016102f7906110e8565b60405180910390fd5b816002826040516103119190610f6b565b9081526020016040518091039020908161032b91906112b4565b507f6b557a50618841e0cdca1b7d2e68c490faf639331b28e17efa4930ea03d88e9d818484426040516103619493929190610ba8565b60405180910390a1505050565b6001604051806080016040528085815260200184815260200183815260200142815250908060018154018082558091505060019003906000526020600020906004020160009091909190915060008201518160000190816103cf91906112b4565b5060208201518160010190816103e591906112b4565b5060408201518160020190816103fb91906112b4565b506060820151816003015550507fe6bfbd050def41eb3a8f15a21087f802b211548d57828abcb3741ec236bc1b378383834260405161043d9493929190610ba8565b60405180910390a1505050565b6001818154811061045a57600080fd5b906000526020600020906004020160009150905080600001805461047d90610fb1565b80601f01602080910402602001604051908101604052809291908181526020018280546104a990610fb1565b80156104f65780601f106104cb576101008083540402835291602001916104f6565b820191906000526020600020905b8154815290600101906020018083116104d957829003601f168201915b50505050509080600101805461050b90610fb1565b80601f016020809104026020016040519081016040528092919081815260200182805461053790610fb1565b80156105845780601f1061055957610100808354040283529160200191610584565b820191906000526020600020905b81548152906001019060200180831161056757829003601f168201915b50505050509080600201805461059990610fb1565b80601f01602080910402602001604051908101604052809291908181526020018280546105c590610fb1565b80156106125780601f106105e757610100808354040283529160200191610612565b820191906000526020600020905b8154815290600101906020018083116105f557829003601f168201915b5050505050908060030154905084565b60005b815181101561072f57600082828151811061064357610642611386565b5b60200260200101519050600060028260405161065f9190610f6b565b9081526020016040518091039020805461067890610fb1565b9050146106ba576040517f08c379a00000000000000000000000000000000000000000000000000000000081526004016106b190611401565b60405180910390fd5b836002826040516106cb9190610f6b565b908152602001604051809103902090816106e591906112b4565b507f3c3087eb611c2be646224be60e6856109eeb94b7d90e6e707fed5c096a82703e81854260405161071993929190611421565b60405180910390a1508080600101915050610625565b505050565b60005481565b600281805160208101820180518482526020830160208501208183528095505050505050600091509050805461076f90610fb1565b80601f016020809104026020016040519081016040528092919081815260200182805461079b90610fb1565b80156107e85780601f106107bd576101008083540402835291602001916107e8565b820191906000526020600020905b8154815290600101906020018083116107cb57829003601f168201915b505050505081565b6000600180549050905090565b606060028260405161080f9190610f6b565b9081526020016040518091039020805461082890610fb1565b80601f016020809104026020016040519081016040528092919081815260200182805461085490610fb1565b80156108a15780601f10610876576101008083540402835291602001916108a1565b820191906000526020600020905b81548152906001019060200180831161088457829003601f168201915b50505050509050919050565b60008054905090565b6000604051905090565b600080fd5b600080fd5b600080fd5b600080fd5b6000601f19601f8301169050919050565b7f4e487b7100000000000000000000000000000000000000000000000000000000600052604160045260246000fd5b61091d826108d4565b810181811067ffffffffffffffff8211171561093c5761093b6108e5565b5b80604052505050565b600061094f6108b6565b905061095b8282610914565b919050565b600067ffffffffffffffff82111561097b5761097a6108e5565b5b610984826108d4565b9050602081019050919050565b82818337600083830152505050565b60006109b36109ae84610960565b610945565b9050828152602081018484840111156109cf576109ce6108cf565b5b6109da848285610991565b509392505050565b600082601f8301126109f7576109f66108ca565b5b8135610a078482602086016109a0565b91505092915050565b600080600060608486031215610a2957610a286108c0565b5b600084013567ffffffffffffffff811115610a4757610a466108c5565b5b610a53868287016109e2565b935050602084013567ffffffffffffffff811115610a7457610a736108c5565b5b610a80868287016109e2565b925050604084013567ffffffffffffffff811115610aa157610aa06108c5565b5b610aad868287016109e2565b9150509250925092565b6000819050919050565b610aca81610ab7565b8114610ad557600080fd5b50565b600081359050610ae781610ac1565b92915050565b600060208284031215610b0357610b026108c0565b5b6000610b1184828501610ad8565b91505092915050565b600081519050919050565b600082825260208201905092915050565b60005b83811015610b54578082015181840152602081019050610b39565b60008484015250505050565b6000610b6b82610b1a565b610b758185610b25565b9350610b85818560208601610b36565b610b8e816108d4565b840191505092915050565b610ba281610ab7565b82525050565b60006080820190508181036000830152610bc28187610b60565b90508181036020830152610bd68186610b60565b90508181036040830152610bea8185610b60565b9050610bf96060830184610b99565b95945050505050565b600067ffffffffffffffff821115610c1d57610c1c6108e5565b5b602082029050602081019050919050565b600080fd5b6000610c46610c4184610c02565b610945565b90508083825260208201905060208402830185811115610c6957610c68610c2e565b5b835b81811015610cb057803567ffffffffffffffff811115610c8e57610c8d6108ca565b5b808601610c9b89826109e2565b85526020850194505050602081019050610c6b565b5050509392505050565b600082601f830112610ccf57610cce6108ca565b5b8135610cdf848260208601610c33565b91505092915050565b60008060408385031215610cff57610cfe6108c0565b5b600083013567ffffffffffffffff811115610d1d57610d1c6108c5565b5b610d29858286016109e2565b925050602083013567ffffffffffffffff811115610d4a57610d496108c5565b5b610d5685828601610cba565b9150509250929050565b6000602082019050610d756000830184610b99565b92915050565b600060208284031215610d9157610d906108c0565b5b600082013567ffffffffffffffff811115610daf57610dae6108c5565b5b610dbb848285016109e2565b91505092915050565b60006020820190508181036000830152610dde8184610b60565b905092915050565b7f4573746f717565206573676f7461646f206e6f20426c6f636b636861696e0000600082015250565b6000610e1c601e83610b25565b9150610e2782610de6565b602082019050919050565b60006020820190508181036000830152610e4b81610e0f565b9050919050565b7f4e487b7100000000000000000000000000000000000000000000000000000000600052601160045260246000fd5b6000610e8c82610ab7565b9150610e9783610ab7565b9250828203905081811115610eaf57610eae610e52565b5b92915050565b7f5061636f74652061626572746f00000000000000000000000000000000000000600082015250565b6000610eeb600d83610b25565b9150610ef682610eb5565b602082019050919050565b6000604082019050610f166000830184610b99565b8181036020830152610f2781610ede565b905092915050565b600081905092915050565b6000610f4582610b1a565b610f4f8185610f2f565b9350610f5f818560208601610b36565b80840191505092915050565b6000610f778284610f3a565b915081905092915050565b7f4e487b7100000000000000000000000000000000000000000000000000000000600052602260045260246000fd5b60006002820490506001821680610fc957607f821691505b602082108103610fdc57610fdb610f82565b5b50919050565b600081905092915050565b60008190508160005260206000209050919050565b6000815461100f81610fb1565b6110198186610fe2565b9450600182166000811461103457600181146110495761107c565b60ff198316865281151582028601935061107c565b61105285610fed565b60005b8381101561107457815481890152600182019150602081019050611055565b838801955050505b50505092915050565b60006110918284611002565b915081905092915050565b7f52656d6574656e7465206e616f20706f73737569206573746120636172746100600082015250565b60006110d2601f83610b25565b91506110dd8261109c565b602082019050919050565b60006020820190508181036000830152611101816110c5565b9050919050565b60008190508160005260206000209050919050565b60006020601f8301049050919050565b600082821b905092915050565b60006008830261116a7fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff8261112d565b611174868361112d565b95508019841693508086168417925050509392505050565b6000819050919050565b60006111b16111ac6111a784610ab7565b61118c565b610ab7565b9050919050565b6000819050919050565b6111cb83611196565b6111df6111d7826111b8565b84845461113a565b825550505050565b600090565b6111f46111e7565b6111ff8184846111c2565b505050565b5b81811015611223576112186000826111ec565b600181019050611205565b5050565b601f8211156112685761123981611108565b6112428461111d565b81016020851015611251578190505b61126561125d8561111d565b830182611204565b50505b505050565b600082821c905092915050565b600061128b6000198460080261126d565b1980831691505092915050565b60006112a4838361127a565b9150826002028217905092915050565b6112bd82610b1a565b67ffffffffffffffff8111156112d6576112d56108e5565b5b6112e08254610fb1565b6112eb828285611227565b600060209050601f83116001811461131e576000841561130c578287015190505b6113168582611298565b86555061137e565b601f19841661132c86611108565b60005b828110156113545784890151825560018201915060208501945060208101905061132f565b86831015611371578489015161136d601f89168261127a565b8355505b6001600288020188555050505b505050505050565b7f4e487b7100000000000000000000000000000000000000000000000000000000600052603260045260246000fd5b7f4361727461206a612074656d20646f6e6f000000000000000000000000000000600082015250565b60006113eb601183610b25565b91506113f6826113b5565b602082019050919050565b6000602082019050818103600083015261141a816113de565b9050919050565b6000606082019050818103600083015261143b8186610b60565b9050818103602083015261144f8185610b60565b905061145e6040830184610b99565b94935050505056fea2646970667358221220f6226ee2f79ed430c6fc320a96bbf5b5de64ada65cb3a1e69ef8f8e1bdcdfd1764736f6c63430008180033",
}

// PackRegistryABI is the input ABI used to generate the binding from.
// Deprecated: Use PackRegistryMetaData.ABI instead.
var PackRegistryABI = PackRegistryMetaData.ABI

// PackRegistryBin is the compiled bytecode used for deploying new contracts.
// Deprecated: Use PackRegistryMetaData.Bin instead.
var PackRegistryBin = PackRegistryMetaData.Bin

// DeployPackRegistry deploys a new Ethereum contract, binding an instance of PackRegistry to it.
func DeployPackRegistry(auth *bind.TransactOpts, backend bind.ContractBackend, _initialStock *big.Int) (common.Address, *types.Transaction, *PackRegistry, error) {
	parsed, err := PackRegistryMetaData.GetAbi()
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	if parsed == nil {
		return common.Address{}, nil, nil, errors.New("GetABI returned nil")
	}

	address, tx, contract, err := bind.DeployContract(auth, *parsed, common.FromHex(PackRegistryBin), backend, _initialStock)
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	return address, tx, &PackRegistry{PackRegistryCaller: PackRegistryCaller{contract: contract}, PackRegistryTransactor: PackRegistryTransactor{contract: contract}, PackRegistryFilterer: PackRegistryFilterer{contract: contract}}, nil
}

// PackRegistry is an auto generated Go binding around an Ethereum contract.
type PackRegistry struct {
	PackRegistryCaller     // Read-only binding to the contract
	PackRegistryTransactor // Write-only binding to the contract
	PackRegistryFilterer   // Log filterer for contract events
}

// PackRegistryCaller is an auto generated read-only Go binding around an Ethereum contract.
type PackRegistryCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// PackRegistryTransactor is an auto generated write-only Go binding around an Ethereum contract.
type PackRegistryTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// PackRegistryFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type PackRegistryFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// PackRegistrySession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type PackRegistrySession struct {
	Contract     *PackRegistry     // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// PackRegistryCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type PackRegistryCallerSession struct {
	Contract *PackRegistryCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts       // Call options to use throughout this session
}

// PackRegistryTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type PackRegistryTransactorSession struct {
	Contract     *PackRegistryTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts       // Transaction auth options to use throughout this session
}

// PackRegistryRaw is an auto generated low-level Go binding around an Ethereum contract.
type PackRegistryRaw struct {
	Contract *PackRegistry // Generic contract binding to access the raw methods on
}

// PackRegistryCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type PackRegistryCallerRaw struct {
	Contract *PackRegistryCaller // Generic read-only contract binding to access the raw methods on
}

// PackRegistryTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type PackRegistryTransactorRaw struct {
	Contract *PackRegistryTransactor // Generic write-only contract binding to access the raw methods on
}

// NewPackRegistry creates a new instance of PackRegistry, bound to a specific deployed contract.
func NewPackRegistry(address common.Address, backend bind.ContractBackend) (*PackRegistry, error) {
	contract, err := bindPackRegistry(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &PackRegistry{PackRegistryCaller: PackRegistryCaller{contract: contract}, PackRegistryTransactor: PackRegistryTransactor{contract: contract}, PackRegistryFilterer: PackRegistryFilterer{contract: contract}}, nil
}

// NewPackRegistryCaller creates a new read-only instance of PackRegistry, bound to a specific deployed contract.
func NewPackRegistryCaller(address common.Address, caller bind.ContractCaller) (*PackRegistryCaller, error) {
	contract, err := bindPackRegistry(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &PackRegistryCaller{contract: contract}, nil
}

// NewPackRegistryTransactor creates a new write-only instance of PackRegistry, bound to a specific deployed contract.
func NewPackRegistryTransactor(address common.Address, transactor bind.ContractTransactor) (*PackRegistryTransactor, error) {
	contract, err := bindPackRegistry(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &PackRegistryTransactor{contract: contract}, nil
}

// NewPackRegistryFilterer creates a new log filterer instance of PackRegistry, bound to a specific deployed contract.
func NewPackRegistryFilterer(address common.Address, filterer bind.ContractFilterer) (*PackRegistryFilterer, error) {
	contract, err := bindPackRegistry(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &PackRegistryFilterer{contract: contract}, nil
}

// bindPackRegistry binds a generic wrapper to an already deployed contract.
func bindPackRegistry(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := PackRegistryMetaData.GetAbi()
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, *parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_PackRegistry *PackRegistryRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _PackRegistry.Contract.PackRegistryCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_PackRegistry *PackRegistryRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _PackRegistry.Contract.PackRegistryTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_PackRegistry *PackRegistryRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _PackRegistry.Contract.PackRegistryTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_PackRegistry *PackRegistryCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _PackRegistry.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_PackRegistry *PackRegistryTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _PackRegistry.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_PackRegistry *PackRegistryTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _PackRegistry.Contract.contract.Transact(opts, method, params...)
}

// CardOwner is a free data retrieval call binding the contract method 0x888ad898.
//
// Solidity: function cardOwner(string ) view returns(string)
func (_PackRegistry *PackRegistryCaller) CardOwner(opts *bind.CallOpts, arg0 string) (string, error) {
	var out []interface{}
	err := _PackRegistry.contract.Call(opts, &out, "cardOwner", arg0)

	if err != nil {
		return *new(string), err
	}

	out0 := *abi.ConvertType(out[0], new(string)).(*string)

	return out0, err

}

// CardOwner is a free data retrieval call binding the contract method 0x888ad898.
//
// Solidity: function cardOwner(string ) view returns(string)
func (_PackRegistry *PackRegistrySession) CardOwner(arg0 string) (string, error) {
	return _PackRegistry.Contract.CardOwner(&_PackRegistry.CallOpts, arg0)
}

// CardOwner is a free data retrieval call binding the contract method 0x888ad898.
//
// Solidity: function cardOwner(string ) view returns(string)
func (_PackRegistry *PackRegistryCallerSession) CardOwner(arg0 string) (string, error) {
	return _PackRegistry.Contract.CardOwner(&_PackRegistry.CallOpts, arg0)
}

// GetCardOwner is a free data retrieval call binding the contract method 0xa79e5267.
//
// Solidity: function getCardOwner(string _cardId) view returns(string)
func (_PackRegistry *PackRegistryCaller) GetCardOwner(opts *bind.CallOpts, _cardId string) (string, error) {
	var out []interface{}
	err := _PackRegistry.contract.Call(opts, &out, "getCardOwner", _cardId)

	if err != nil {
		return *new(string), err
	}

	out0 := *abi.ConvertType(out[0], new(string)).(*string)

	return out0, err

}

// GetCardOwner is a free data retrieval call binding the contract method 0xa79e5267.
//
// Solidity: function getCardOwner(string _cardId) view returns(string)
func (_PackRegistry *PackRegistrySession) GetCardOwner(_cardId string) (string, error) {
	return _PackRegistry.Contract.GetCardOwner(&_PackRegistry.CallOpts, _cardId)
}

// GetCardOwner is a free data retrieval call binding the contract method 0xa79e5267.
//
// Solidity: function getCardOwner(string _cardId) view returns(string)
func (_PackRegistry *PackRegistryCallerSession) GetCardOwner(_cardId string) (string, error) {
	return _PackRegistry.Contract.GetCardOwner(&_PackRegistry.CallOpts, _cardId)
}

// GetMatchCount is a free data retrieval call binding the contract method 0x8c4f7dae.
//
// Solidity: function getMatchCount() view returns(uint256)
func (_PackRegistry *PackRegistryCaller) GetMatchCount(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _PackRegistry.contract.Call(opts, &out, "getMatchCount")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// GetMatchCount is a free data retrieval call binding the contract method 0x8c4f7dae.
//
// Solidity: function getMatchCount() view returns(uint256)
func (_PackRegistry *PackRegistrySession) GetMatchCount() (*big.Int, error) {
	return _PackRegistry.Contract.GetMatchCount(&_PackRegistry.CallOpts)
}

// GetMatchCount is a free data retrieval call binding the contract method 0x8c4f7dae.
//
// Solidity: function getMatchCount() view returns(uint256)
func (_PackRegistry *PackRegistryCallerSession) GetMatchCount() (*big.Int, error) {
	return _PackRegistry.Contract.GetMatchCount(&_PackRegistry.CallOpts)
}

// GetStock is a free data retrieval call binding the contract method 0xe38ffcd8.
//
// Solidity: function getStock() view returns(uint256)
func (_PackRegistry *PackRegistryCaller) GetStock(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _PackRegistry.contract.Call(opts, &out, "getStock")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// GetStock is a free data retrieval call binding the contract method 0xe38ffcd8.
//
// Solidity: function getStock() view returns(uint256)
func (_PackRegistry *PackRegistrySession) GetStock() (*big.Int, error) {
	return _PackRegistry.Contract.GetStock(&_PackRegistry.CallOpts)
}

// GetStock is a free data retrieval call binding the contract method 0xe38ffcd8.
//
// Solidity: function getStock() view returns(uint256)
func (_PackRegistry *PackRegistryCallerSession) GetStock() (*big.Int, error) {
	return _PackRegistry.Contract.GetStock(&_PackRegistry.CallOpts)
}

// Matches is a free data retrieval call binding the contract method 0x4768d4ef.
//
// Solidity: function matches(uint256 ) view returns(string matchId, string winnerId, string loserId, uint256 timestamp)
func (_PackRegistry *PackRegistryCaller) Matches(opts *bind.CallOpts, arg0 *big.Int) (struct {
	MatchId   string
	WinnerId  string
	LoserId   string
	Timestamp *big.Int
}, error) {
	var out []interface{}
	err := _PackRegistry.contract.Call(opts, &out, "matches", arg0)

	outstruct := new(struct {
		MatchId   string
		WinnerId  string
		LoserId   string
		Timestamp *big.Int
	})
	if err != nil {
		return *outstruct, err
	}

	outstruct.MatchId = *abi.ConvertType(out[0], new(string)).(*string)
	outstruct.WinnerId = *abi.ConvertType(out[1], new(string)).(*string)
	outstruct.LoserId = *abi.ConvertType(out[2], new(string)).(*string)
	outstruct.Timestamp = *abi.ConvertType(out[3], new(*big.Int)).(**big.Int)

	return *outstruct, err

}

// Matches is a free data retrieval call binding the contract method 0x4768d4ef.
//
// Solidity: function matches(uint256 ) view returns(string matchId, string winnerId, string loserId, uint256 timestamp)
func (_PackRegistry *PackRegistrySession) Matches(arg0 *big.Int) (struct {
	MatchId   string
	WinnerId  string
	LoserId   string
	Timestamp *big.Int
}, error) {
	return _PackRegistry.Contract.Matches(&_PackRegistry.CallOpts, arg0)
}

// Matches is a free data retrieval call binding the contract method 0x4768d4ef.
//
// Solidity: function matches(uint256 ) view returns(string matchId, string winnerId, string loserId, uint256 timestamp)
func (_PackRegistry *PackRegistryCallerSession) Matches(arg0 *big.Int) (struct {
	MatchId   string
	WinnerId  string
	LoserId   string
	Timestamp *big.Int
}, error) {
	return _PackRegistry.Contract.Matches(&_PackRegistry.CallOpts, arg0)
}

// TotalPacks is a free data retrieval call binding the contract method 0x7281df12.
//
// Solidity: function totalPacks() view returns(uint256)
func (_PackRegistry *PackRegistryCaller) TotalPacks(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _PackRegistry.contract.Call(opts, &out, "totalPacks")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// TotalPacks is a free data retrieval call binding the contract method 0x7281df12.
//
// Solidity: function totalPacks() view returns(uint256)
func (_PackRegistry *PackRegistrySession) TotalPacks() (*big.Int, error) {
	return _PackRegistry.Contract.TotalPacks(&_PackRegistry.CallOpts)
}

// TotalPacks is a free data retrieval call binding the contract method 0x7281df12.
//
// Solidity: function totalPacks() view returns(uint256)
func (_PackRegistry *PackRegistryCallerSession) TotalPacks() (*big.Int, error) {
	return _PackRegistry.Contract.TotalPacks(&_PackRegistry.CallOpts)
}

// AssignCards is a paid mutator transaction binding the contract method 0x6d34550e.
//
// Solidity: function assignCards(string _playerId, string[] _cardIds) returns()
func (_PackRegistry *PackRegistryTransactor) AssignCards(opts *bind.TransactOpts, _playerId string, _cardIds []string) (*types.Transaction, error) {
	return _PackRegistry.contract.Transact(opts, "assignCards", _playerId, _cardIds)
}

// AssignCards is a paid mutator transaction binding the contract method 0x6d34550e.
//
// Solidity: function assignCards(string _playerId, string[] _cardIds) returns()
func (_PackRegistry *PackRegistrySession) AssignCards(_playerId string, _cardIds []string) (*types.Transaction, error) {
	return _PackRegistry.Contract.AssignCards(&_PackRegistry.TransactOpts, _playerId, _cardIds)
}

// AssignCards is a paid mutator transaction binding the contract method 0x6d34550e.
//
// Solidity: function assignCards(string _playerId, string[] _cardIds) returns()
func (_PackRegistry *PackRegistryTransactorSession) AssignCards(_playerId string, _cardIds []string) (*types.Transaction, error) {
	return _PackRegistry.Contract.AssignCards(&_PackRegistry.TransactOpts, _playerId, _cardIds)
}

// DecrementStock is a paid mutator transaction binding the contract method 0x00375a44.
//
// Solidity: function decrementStock() returns()
func (_PackRegistry *PackRegistryTransactor) DecrementStock(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _PackRegistry.contract.Transact(opts, "decrementStock")
}

// DecrementStock is a paid mutator transaction binding the contract method 0x00375a44.
//
// Solidity: function decrementStock() returns()
func (_PackRegistry *PackRegistrySession) DecrementStock() (*types.Transaction, error) {
	return _PackRegistry.Contract.DecrementStock(&_PackRegistry.TransactOpts)
}

// DecrementStock is a paid mutator transaction binding the contract method 0x00375a44.
//
// Solidity: function decrementStock() returns()
func (_PackRegistry *PackRegistryTransactorSession) DecrementStock() (*types.Transaction, error) {
	return _PackRegistry.Contract.DecrementStock(&_PackRegistry.TransactOpts)
}

// RecordMatch is a paid mutator transaction binding the contract method 0x12636a62.
//
// Solidity: function recordMatch(string _matchId, string _winnerId, string _loserId) returns()
func (_PackRegistry *PackRegistryTransactor) RecordMatch(opts *bind.TransactOpts, _matchId string, _winnerId string, _loserId string) (*types.Transaction, error) {
	return _PackRegistry.contract.Transact(opts, "recordMatch", _matchId, _winnerId, _loserId)
}

// RecordMatch is a paid mutator transaction binding the contract method 0x12636a62.
//
// Solidity: function recordMatch(string _matchId, string _winnerId, string _loserId) returns()
func (_PackRegistry *PackRegistrySession) RecordMatch(_matchId string, _winnerId string, _loserId string) (*types.Transaction, error) {
	return _PackRegistry.Contract.RecordMatch(&_PackRegistry.TransactOpts, _matchId, _winnerId, _loserId)
}

// RecordMatch is a paid mutator transaction binding the contract method 0x12636a62.
//
// Solidity: function recordMatch(string _matchId, string _winnerId, string _loserId) returns()
func (_PackRegistry *PackRegistryTransactorSession) RecordMatch(_matchId string, _winnerId string, _loserId string) (*types.Transaction, error) {
	return _PackRegistry.Contract.RecordMatch(&_PackRegistry.TransactOpts, _matchId, _winnerId, _loserId)
}

// TransferCard is a paid mutator transaction binding the contract method 0x118c54ec.
//
// Solidity: function transferCard(string _fromId, string _toId, string _cardId) returns()
func (_PackRegistry *PackRegistryTransactor) TransferCard(opts *bind.TransactOpts, _fromId string, _toId string, _cardId string) (*types.Transaction, error) {
	return _PackRegistry.contract.Transact(opts, "transferCard", _fromId, _toId, _cardId)
}

// TransferCard is a paid mutator transaction binding the contract method 0x118c54ec.
//
// Solidity: function transferCard(string _fromId, string _toId, string _cardId) returns()
func (_PackRegistry *PackRegistrySession) TransferCard(_fromId string, _toId string, _cardId string) (*types.Transaction, error) {
	return _PackRegistry.Contract.TransferCard(&_PackRegistry.TransactOpts, _fromId, _toId, _cardId)
}

// TransferCard is a paid mutator transaction binding the contract method 0x118c54ec.
//
// Solidity: function transferCard(string _fromId, string _toId, string _cardId) returns()
func (_PackRegistry *PackRegistryTransactorSession) TransferCard(_fromId string, _toId string, _cardId string) (*types.Transaction, error) {
	return _PackRegistry.Contract.TransferCard(&_PackRegistry.TransactOpts, _fromId, _toId, _cardId)
}

// PackRegistryCardAssignedIterator is returned from FilterCardAssigned and is used to iterate over the raw logs and unpacked data for CardAssigned events raised by the PackRegistry contract.
type PackRegistryCardAssignedIterator struct {
	Event *PackRegistryCardAssigned // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *PackRegistryCardAssignedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(PackRegistryCardAssigned)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(PackRegistryCardAssigned)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *PackRegistryCardAssignedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *PackRegistryCardAssignedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// PackRegistryCardAssigned represents a CardAssigned event raised by the PackRegistry contract.
type PackRegistryCardAssigned struct {
	CardId    string
	OwnerId   string
	Timestamp *big.Int
	Raw       types.Log // Blockchain specific contextual infos
}

// FilterCardAssigned is a free log retrieval operation binding the contract event 0x3c3087eb611c2be646224be60e6856109eeb94b7d90e6e707fed5c096a82703e.
//
// Solidity: event CardAssigned(string cardId, string ownerId, uint256 timestamp)
func (_PackRegistry *PackRegistryFilterer) FilterCardAssigned(opts *bind.FilterOpts) (*PackRegistryCardAssignedIterator, error) {

	logs, sub, err := _PackRegistry.contract.FilterLogs(opts, "CardAssigned")
	if err != nil {
		return nil, err
	}
	return &PackRegistryCardAssignedIterator{contract: _PackRegistry.contract, event: "CardAssigned", logs: logs, sub: sub}, nil
}

// WatchCardAssigned is a free log subscription operation binding the contract event 0x3c3087eb611c2be646224be60e6856109eeb94b7d90e6e707fed5c096a82703e.
//
// Solidity: event CardAssigned(string cardId, string ownerId, uint256 timestamp)
func (_PackRegistry *PackRegistryFilterer) WatchCardAssigned(opts *bind.WatchOpts, sink chan<- *PackRegistryCardAssigned) (event.Subscription, error) {

	logs, sub, err := _PackRegistry.contract.WatchLogs(opts, "CardAssigned")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(PackRegistryCardAssigned)
				if err := _PackRegistry.contract.UnpackLog(event, "CardAssigned", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// ParseCardAssigned is a log parse operation binding the contract event 0x3c3087eb611c2be646224be60e6856109eeb94b7d90e6e707fed5c096a82703e.
//
// Solidity: event CardAssigned(string cardId, string ownerId, uint256 timestamp)
func (_PackRegistry *PackRegistryFilterer) ParseCardAssigned(log types.Log) (*PackRegistryCardAssigned, error) {
	event := new(PackRegistryCardAssigned)
	if err := _PackRegistry.contract.UnpackLog(event, "CardAssigned", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// PackRegistryCardTransferredIterator is returned from FilterCardTransferred and is used to iterate over the raw logs and unpacked data for CardTransferred events raised by the PackRegistry contract.
type PackRegistryCardTransferredIterator struct {
	Event *PackRegistryCardTransferred // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *PackRegistryCardTransferredIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(PackRegistryCardTransferred)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(PackRegistryCardTransferred)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *PackRegistryCardTransferredIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *PackRegistryCardTransferredIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// PackRegistryCardTransferred represents a CardTransferred event raised by the PackRegistry contract.
type PackRegistryCardTransferred struct {
	CardId    string
	FromId    string
	ToId      string
	Timestamp *big.Int
	Raw       types.Log // Blockchain specific contextual infos
}

// FilterCardTransferred is a free log retrieval operation binding the contract event 0x6b557a50618841e0cdca1b7d2e68c490faf639331b28e17efa4930ea03d88e9d.
//
// Solidity: event CardTransferred(string cardId, string fromId, string toId, uint256 timestamp)
func (_PackRegistry *PackRegistryFilterer) FilterCardTransferred(opts *bind.FilterOpts) (*PackRegistryCardTransferredIterator, error) {

	logs, sub, err := _PackRegistry.contract.FilterLogs(opts, "CardTransferred")
	if err != nil {
		return nil, err
	}
	return &PackRegistryCardTransferredIterator{contract: _PackRegistry.contract, event: "CardTransferred", logs: logs, sub: sub}, nil
}

// WatchCardTransferred is a free log subscription operation binding the contract event 0x6b557a50618841e0cdca1b7d2e68c490faf639331b28e17efa4930ea03d88e9d.
//
// Solidity: event CardTransferred(string cardId, string fromId, string toId, uint256 timestamp)
func (_PackRegistry *PackRegistryFilterer) WatchCardTransferred(opts *bind.WatchOpts, sink chan<- *PackRegistryCardTransferred) (event.Subscription, error) {

	logs, sub, err := _PackRegistry.contract.WatchLogs(opts, "CardTransferred")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(PackRegistryCardTransferred)
				if err := _PackRegistry.contract.UnpackLog(event, "CardTransferred", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// ParseCardTransferred is a log parse operation binding the contract event 0x6b557a50618841e0cdca1b7d2e68c490faf639331b28e17efa4930ea03d88e9d.
//
// Solidity: event CardTransferred(string cardId, string fromId, string toId, uint256 timestamp)
func (_PackRegistry *PackRegistryFilterer) ParseCardTransferred(log types.Log) (*PackRegistryCardTransferred, error) {
	event := new(PackRegistryCardTransferred)
	if err := _PackRegistry.contract.UnpackLog(event, "CardTransferred", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// PackRegistryMatchRecordedIterator is returned from FilterMatchRecorded and is used to iterate over the raw logs and unpacked data for MatchRecorded events raised by the PackRegistry contract.
type PackRegistryMatchRecordedIterator struct {
	Event *PackRegistryMatchRecorded // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *PackRegistryMatchRecordedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(PackRegistryMatchRecorded)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(PackRegistryMatchRecorded)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *PackRegistryMatchRecordedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *PackRegistryMatchRecordedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// PackRegistryMatchRecorded represents a MatchRecorded event raised by the PackRegistry contract.
type PackRegistryMatchRecorded struct {
	MatchId   string
	Winner    string
	Loser     string
	Timestamp *big.Int
	Raw       types.Log // Blockchain specific contextual infos
}

// FilterMatchRecorded is a free log retrieval operation binding the contract event 0xe6bfbd050def41eb3a8f15a21087f802b211548d57828abcb3741ec236bc1b37.
//
// Solidity: event MatchRecorded(string matchId, string winner, string loser, uint256 timestamp)
func (_PackRegistry *PackRegistryFilterer) FilterMatchRecorded(opts *bind.FilterOpts) (*PackRegistryMatchRecordedIterator, error) {

	logs, sub, err := _PackRegistry.contract.FilterLogs(opts, "MatchRecorded")
	if err != nil {
		return nil, err
	}
	return &PackRegistryMatchRecordedIterator{contract: _PackRegistry.contract, event: "MatchRecorded", logs: logs, sub: sub}, nil
}

// WatchMatchRecorded is a free log subscription operation binding the contract event 0xe6bfbd050def41eb3a8f15a21087f802b211548d57828abcb3741ec236bc1b37.
//
// Solidity: event MatchRecorded(string matchId, string winner, string loser, uint256 timestamp)
func (_PackRegistry *PackRegistryFilterer) WatchMatchRecorded(opts *bind.WatchOpts, sink chan<- *PackRegistryMatchRecorded) (event.Subscription, error) {

	logs, sub, err := _PackRegistry.contract.WatchLogs(opts, "MatchRecorded")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(PackRegistryMatchRecorded)
				if err := _PackRegistry.contract.UnpackLog(event, "MatchRecorded", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// ParseMatchRecorded is a log parse operation binding the contract event 0xe6bfbd050def41eb3a8f15a21087f802b211548d57828abcb3741ec236bc1b37.
//
// Solidity: event MatchRecorded(string matchId, string winner, string loser, uint256 timestamp)
func (_PackRegistry *PackRegistryFilterer) ParseMatchRecorded(log types.Log) (*PackRegistryMatchRecorded, error) {
	event := new(PackRegistryMatchRecorded)
	if err := _PackRegistry.contract.UnpackLog(event, "MatchRecorded", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// PackRegistryStockUpdatedIterator is returned from FilterStockUpdated and is used to iterate over the raw logs and unpacked data for StockUpdated events raised by the PackRegistry contract.
type PackRegistryStockUpdatedIterator struct {
	Event *PackRegistryStockUpdated // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *PackRegistryStockUpdatedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(PackRegistryStockUpdated)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(PackRegistryStockUpdated)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *PackRegistryStockUpdatedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *PackRegistryStockUpdatedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// PackRegistryStockUpdated represents a StockUpdated event raised by the PackRegistry contract.
type PackRegistryStockUpdated struct {
	NewStock *big.Int
	Reason   string
	Raw      types.Log // Blockchain specific contextual infos
}

// FilterStockUpdated is a free log retrieval operation binding the contract event 0x2926060a4c291371fea4258fa171860291f0144cd93a2accc8a0da15fc05c7a0.
//
// Solidity: event StockUpdated(uint256 newStock, string reason)
func (_PackRegistry *PackRegistryFilterer) FilterStockUpdated(opts *bind.FilterOpts) (*PackRegistryStockUpdatedIterator, error) {

	logs, sub, err := _PackRegistry.contract.FilterLogs(opts, "StockUpdated")
	if err != nil {
		return nil, err
	}
	return &PackRegistryStockUpdatedIterator{contract: _PackRegistry.contract, event: "StockUpdated", logs: logs, sub: sub}, nil
}

// WatchStockUpdated is a free log subscription operation binding the contract event 0x2926060a4c291371fea4258fa171860291f0144cd93a2accc8a0da15fc05c7a0.
//
// Solidity: event StockUpdated(uint256 newStock, string reason)
func (_PackRegistry *PackRegistryFilterer) WatchStockUpdated(opts *bind.WatchOpts, sink chan<- *PackRegistryStockUpdated) (event.Subscription, error) {

	logs, sub, err := _PackRegistry.contract.WatchLogs(opts, "StockUpdated")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(PackRegistryStockUpdated)
				if err := _PackRegistry.contract.UnpackLog(event, "StockUpdated", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// ParseStockUpdated is a log parse operation binding the contract event 0x2926060a4c291371fea4258fa171860291f0144cd93a2accc8a0da15fc05c7a0.
//
// Solidity: event StockUpdated(uint256 newStock, string reason)
func (_PackRegistry *PackRegistryFilterer) ParseStockUpdated(log types.Log) (*PackRegistryStockUpdated, error) {
	event := new(PackRegistryStockUpdated)
	if err := _PackRegistry.contract.UnpackLog(event, "StockUpdated", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}
