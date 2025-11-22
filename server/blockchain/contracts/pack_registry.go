// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package contracts

import (
	"errors"
	"math/big"
	"strings"

	"github.com/ethereum/go-ethereum"
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
	ABI: "[{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_initialStock\",\"type\":\"uint256\"}],\"stateMutability\":\"nonpayable\",\"type\":\"constructor\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"newStock\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"string\",\"name\":\"reason\",\"type\":\"string\"}],\"name\":\"StockUpdated\",\"type\":\"event\"},{\"inputs\":[],\"name\":\"decrementStock\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"getStock\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_newStock\",\"type\":\"uint256\"},{\"internalType\":\"string\",\"name\":\"_reason\",\"type\":\"string\"}],\"name\":\"setStock\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"totalPacks\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"}]",
	Bin: "0x608060405234801561001057600080fd5b506040516106973803806106978339818101604052810190610032919061007a565b80600081905550506100a7565b600080fd5b6000819050919050565b61005781610044565b811461006257600080fd5b50565b6000815190506100748161004e565b92915050565b6000602082840312156100905761008f61003f565b5b600061009e84828501610065565b91505092915050565b6105e1806100b66000396000f3fe608060405234801561001057600080fd5b506004361061004b5760003560e01c8062375a44146100505780637281df121461005a578063de9a977414610078578063e38ffcd814610094575b600080fd5b6100586100b2565b005b61006261014a565b60405161006f91906101b6565b60405180910390f35b610092600480360381019061008d9190610357565b610150565b005b61009c610194565b6040516100a991906101b6565b60405180910390f35b60008054116100f6576040517f08c379a00000000000000000000000000000000000000000000000000000000081526004016100ed90610410565b60405180910390fd5b6001600080828254610108919061045f565b925050819055507f2926060a4c291371fea4258fa171860291f0144cd93a2accc8a0da15fc05c7a060005460405161014091906104df565b60405180910390a1565b60005481565b816000819055507f2926060a4c291371fea4258fa171860291f0144cd93a2accc8a0da15fc05c7a0828260405161018892919061057b565b60405180910390a15050565b60008054905090565b6000819050919050565b6101b08161019d565b82525050565b60006020820190506101cb60008301846101a7565b92915050565b6000604051905090565b600080fd5b600080fd5b6101ee8161019d565b81146101f957600080fd5b50565b60008135905061020b816101e5565b92915050565b600080fd5b600080fd5b6000601f19601f8301169050919050565b7f4e487b7100000000000000000000000000000000000000000000000000000000600052604160045260246000fd5b6102648261021b565b810181811067ffffffffffffffff821117156102835761028261022c565b5b80604052505050565b60006102966101d1565b90506102a2828261025b565b919050565b600067ffffffffffffffff8211156102c2576102c161022c565b5b6102cb8261021b565b9050602081019050919050565b82818337600083830152505050565b60006102fa6102f5846102a7565b61028c565b90508281526020810184848401111561031657610315610216565b5b6103218482856102d8565b509392505050565b600082601f83011261033e5761033d610211565b5b813561034e8482602086016102e7565b91505092915050565b6000806040838503121561036e5761036d6101db565b5b600061037c858286016101fc565b925050602083013567ffffffffffffffff81111561039d5761039c6101e0565b5b6103a985828601610329565b9150509250929050565b600082825260208201905092915050565b7f4573746f717565206573676f7461646f206e6f20426c6f636b636861696e0000600082015250565b60006103fa601e836103b3565b9150610405826103c4565b602082019050919050565b60006020820190508181036000830152610429816103ed565b9050919050565b7f4e487b7100000000000000000000000000000000000000000000000000000000600052601160045260246000fd5b600061046a8261019d565b91506104758361019d565b925082820390508181111561048d5761048c610430565b5b92915050565b7f5061636f74652061626572746f00000000000000000000000000000000000000600082015250565b60006104c9600d836103b3565b91506104d482610493565b602082019050919050565b60006040820190506104f460008301846101a7565b8181036020830152610505816104bc565b905092915050565b600081519050919050565b60005b8381101561053657808201518184015260208101905061051b565b60008484015250505050565b600061054d8261050d565b61055781856103b3565b9350610567818560208601610518565b6105708161021b565b840191505092915050565b600060408201905061059060008301856101a7565b81810360208301526105a28184610542565b9050939250505056fea26469706673582212208b86e61fa5832807bf0066aab32e09cbe2b6e8737467e813d70c3f4245f47e0b64736f6c63430008180033",
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

// SetStock is a paid mutator transaction binding the contract method 0xde9a9774.
//
// Solidity: function setStock(uint256 _newStock, string _reason) returns()
func (_PackRegistry *PackRegistryTransactor) SetStock(opts *bind.TransactOpts, _newStock *big.Int, _reason string) (*types.Transaction, error) {
	return _PackRegistry.contract.Transact(opts, "setStock", _newStock, _reason)
}

// SetStock is a paid mutator transaction binding the contract method 0xde9a9774.
//
// Solidity: function setStock(uint256 _newStock, string _reason) returns()
func (_PackRegistry *PackRegistrySession) SetStock(_newStock *big.Int, _reason string) (*types.Transaction, error) {
	return _PackRegistry.Contract.SetStock(&_PackRegistry.TransactOpts, _newStock, _reason)
}

// SetStock is a paid mutator transaction binding the contract method 0xde9a9774.
//
// Solidity: function setStock(uint256 _newStock, string _reason) returns()
func (_PackRegistry *PackRegistryTransactorSession) SetStock(_newStock *big.Int, _reason string) (*types.Transaction, error) {
	return _PackRegistry.Contract.SetStock(&_PackRegistry.TransactOpts, _newStock, _reason)
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
