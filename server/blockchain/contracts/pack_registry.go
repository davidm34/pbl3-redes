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
	ABI: "[{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_initialStock\",\"type\":\"uint256\"}],\"stateMutability\":\"nonpayable\",\"type\":\"constructor\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"string\",\"name\":\"matchId\",\"type\":\"string\"},{\"indexed\":false,\"internalType\":\"string\",\"name\":\"winner\",\"type\":\"string\"},{\"indexed\":false,\"internalType\":\"string\",\"name\":\"loser\",\"type\":\"string\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"timestamp\",\"type\":\"uint256\"}],\"name\":\"MatchRecorded\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"newStock\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"string\",\"name\":\"reason\",\"type\":\"string\"}],\"name\":\"StockUpdated\",\"type\":\"event\"},{\"inputs\":[],\"name\":\"decrementStock\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"getMatchCount\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"getStock\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"name\":\"matches\",\"outputs\":[{\"internalType\":\"string\",\"name\":\"matchId\",\"type\":\"string\"},{\"internalType\":\"string\",\"name\":\"winnerId\",\"type\":\"string\"},{\"internalType\":\"string\",\"name\":\"loserId\",\"type\":\"string\"},{\"internalType\":\"uint256\",\"name\":\"timestamp\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"string\",\"name\":\"_matchId\",\"type\":\"string\"},{\"internalType\":\"string\",\"name\":\"_winnerId\",\"type\":\"string\"},{\"internalType\":\"string\",\"name\":\"_loserId\",\"type\":\"string\"}],\"name\":\"recordMatch\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"totalPacks\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"}]",
	Bin: "0x608060405234801561001057600080fd5b50604051610cfb380380610cfb8339818101604052810190610032919061007a565b80600081905550506100a7565b600080fd5b6000819050919050565b61005781610044565b811461006257600080fd5b50565b6000815190506100748161004e565b92915050565b6000602082840312156100905761008f61003f565b5b600061009e84828501610065565b91505092915050565b610c45806100b66000396000f3fe608060405234801561001057600080fd5b50600436106100615760003560e01c8062375a441461006657806312636a62146100705780634768d4ef1461008c5780637281df12146100bf5780638c4f7dae146100dd578063e38ffcd8146100fb575b600080fd5b61006e610119565b005b61008a600480360381019061008591906105db565b6101b1565b005b6100a660048036038101906100a191906106b8565b61028d565b6040516100b69493929190610773565b60405180910390f35b6100c7610465565b6040516100d491906107cd565b60405180910390f35b6100e561046b565b6040516100f291906107cd565b60405180910390f35b610103610478565b60405161011091906107cd565b60405180910390f35b600080541161015d576040517f08c379a000000000000000000000000000000000000000000000000000000000815260040161015490610834565b60405180910390fd5b600160008082825461016f9190610883565b925050819055507f2926060a4c291371fea4258fa171860291f0144cd93a2accc8a0da15fc05c7a06000546040516101a79190610903565b60405180910390a1565b6001604051806080016040528085815260200184815260200183815260200142815250908060018154018082558091505060019003906000526020600020906004020160009091909190915060008201518160000190816102129190610b3d565b5060208201518160010190816102289190610b3d565b50604082015181600201908161023e9190610b3d565b506060820151816003015550507fe6bfbd050def41eb3a8f15a21087f802b211548d57828abcb3741ec236bc1b37838383426040516102809493929190610773565b60405180910390a1505050565b6001818154811061029d57600080fd5b90600052602060002090600402016000915090508060000180546102c090610960565b80601f01602080910402602001604051908101604052809291908181526020018280546102ec90610960565b80156103395780601f1061030e57610100808354040283529160200191610339565b820191906000526020600020905b81548152906001019060200180831161031c57829003601f168201915b50505050509080600101805461034e90610960565b80601f016020809104026020016040519081016040528092919081815260200182805461037a90610960565b80156103c75780601f1061039c576101008083540402835291602001916103c7565b820191906000526020600020905b8154815290600101906020018083116103aa57829003601f168201915b5050505050908060020180546103dc90610960565b80601f016020809104026020016040519081016040528092919081815260200182805461040890610960565b80156104555780601f1061042a57610100808354040283529160200191610455565b820191906000526020600020905b81548152906001019060200180831161043857829003601f168201915b5050505050908060030154905084565b60005481565b6000600180549050905090565b60008054905090565b6000604051905090565b600080fd5b600080fd5b600080fd5b600080fd5b6000601f19601f8301169050919050565b7f4e487b7100000000000000000000000000000000000000000000000000000000600052604160045260246000fd5b6104e88261049f565b810181811067ffffffffffffffff82111715610507576105066104b0565b5b80604052505050565b600061051a610481565b905061052682826104df565b919050565b600067ffffffffffffffff821115610546576105456104b0565b5b61054f8261049f565b9050602081019050919050565b82818337600083830152505050565b600061057e6105798461052b565b610510565b90508281526020810184848401111561059a5761059961049a565b5b6105a584828561055c565b509392505050565b600082601f8301126105c2576105c1610495565b5b81356105d284826020860161056b565b91505092915050565b6000806000606084860312156105f4576105f361048b565b5b600084013567ffffffffffffffff81111561061257610611610490565b5b61061e868287016105ad565b935050602084013567ffffffffffffffff81111561063f5761063e610490565b5b61064b868287016105ad565b925050604084013567ffffffffffffffff81111561066c5761066b610490565b5b610678868287016105ad565b9150509250925092565b6000819050919050565b61069581610682565b81146106a057600080fd5b50565b6000813590506106b28161068c565b92915050565b6000602082840312156106ce576106cd61048b565b5b60006106dc848285016106a3565b91505092915050565b600081519050919050565b600082825260208201905092915050565b60005b8381101561071f578082015181840152602081019050610704565b60008484015250505050565b6000610736826106e5565b61074081856106f0565b9350610750818560208601610701565b6107598161049f565b840191505092915050565b61076d81610682565b82525050565b6000608082019050818103600083015261078d818761072b565b905081810360208301526107a1818661072b565b905081810360408301526107b5818561072b565b90506107c46060830184610764565b95945050505050565b60006020820190506107e26000830184610764565b92915050565b7f4573746f717565206573676f7461646f206e6f20426c6f636b636861696e0000600082015250565b600061081e601e836106f0565b9150610829826107e8565b602082019050919050565b6000602082019050818103600083015261084d81610811565b9050919050565b7f4e487b7100000000000000000000000000000000000000000000000000000000600052601160045260246000fd5b600061088e82610682565b915061089983610682565b92508282039050818111156108b1576108b0610854565b5b92915050565b7f5061636f74652061626572746f00000000000000000000000000000000000000600082015250565b60006108ed600d836106f0565b91506108f8826108b7565b602082019050919050565b60006040820190506109186000830184610764565b8181036020830152610929816108e0565b905092915050565b7f4e487b7100000000000000000000000000000000000000000000000000000000600052602260045260246000fd5b6000600282049050600182168061097857607f821691505b60208210810361098b5761098a610931565b5b50919050565b60008190508160005260206000209050919050565b60006020601f8301049050919050565b600082821b905092915050565b6000600883026109f37fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff826109b6565b6109fd86836109b6565b95508019841693508086168417925050509392505050565b6000819050919050565b6000610a3a610a35610a3084610682565b610a15565b610682565b9050919050565b6000819050919050565b610a5483610a1f565b610a68610a6082610a41565b8484546109c3565b825550505050565b600090565b610a7d610a70565b610a88818484610a4b565b505050565b5b81811015610aac57610aa1600082610a75565b600181019050610a8e565b5050565b601f821115610af157610ac281610991565b610acb846109a6565b81016020851015610ada578190505b610aee610ae6856109a6565b830182610a8d565b50505b505050565b600082821c905092915050565b6000610b1460001984600802610af6565b1980831691505092915050565b6000610b2d8383610b03565b9150826002028217905092915050565b610b46826106e5565b67ffffffffffffffff811115610b5f57610b5e6104b0565b5b610b698254610960565b610b74828285610ab0565b600060209050601f831160018114610ba75760008415610b95578287015190505b610b9f8582610b21565b865550610c07565b601f198416610bb586610991565b60005b82811015610bdd57848901518255600182019150602085019450602081019050610bb8565b86831015610bfa5784890151610bf6601f891682610b03565b8355505b6001600288020188555050505b50505050505056fea26469706673582212201cedde51889035b926798ba237575bb9f3895380d4b4811ad110d960255062f664736f6c63430008180033",
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
