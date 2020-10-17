// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package mycontract

import (
	"math/big"
	"strings"

	"github.com/FISCO-BCOS/go-sdk/abi"
	"github.com/FISCO-BCOS/go-sdk/abi/bind"
	"github.com/FISCO-BCOS/go-sdk/core/types"
	"github.com/FISCO-BCOS/go-sdk/event"
	ethereum "github.com/ethereum/go-ethereum"
	"github.com/ethereum/go-ethereum/common"
)

// Reference imports to suppress errors if they are not otherwise used.
var (
	_ = big.NewInt
	_ = strings.NewReader
	_ = ethereum.NotFound
	_ = abi.U256
	_ = bind.Bind
	_ = common.Big1
	_ = types.BloomLookup
	_ = event.NewSubscription
)

// ShareDeliveryInfoABI is the input ABI used to generate the binding from.
const ShareDeliveryInfoABI = "[{\"constant\":false,\"inputs\":[{\"name\":\"_pkg_uid\",\"type\":\"uint256\"},{\"name\":\"station\",\"type\":\"string\"}],\"name\":\"set\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"name\":\"_pkg_uid\",\"type\":\"uint256\"}],\"name\":\"get\",\"outputs\":[{\"name\":\"\",\"type\":\"string\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"name\":\"\",\"type\":\"uint256\"}],\"name\":\"delivery_path\",\"outputs\":[{\"name\":\"\",\"type\":\"string\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"constructor\"}]"

// ShareDeliveryInfoBin is the compiled bytecode used for deploying new contracts.
var ShareDeliveryInfoBin = "0x608060405234801561001057600080fd5b50610813806100206000396000f300608060405260043610610057576000357c0100000000000000000000000000000000000000000000000000000000900463ffffffff168063643719771461005c5780639507d39a146100cf578063a0afbd9f14610175575b600080fd5b34801561006857600080fd5b506100cd60048036038101908080359060200190929190803590602001908201803590602001908080601f016020809104026020016040519081016040528093929190818152602001838380828437820191505050505050919291929050505061021b565b005b3480156100db57600080fd5b506100fa6004803603810190808035906020019092919050505061040c565b6040518080602001828103825283818151815260200191508051906020019080838360005b8381101561013a57808201518184015260208101905061011f565b50505050905090810190601f1680156101675780820380516001836020036101000a031916815260200191505b509250505060405180910390f35b34801561018157600080fd5b506101a0600480360381019080803590602001909291905050506104c1565b6040518080602001828103825283818151815260200191508051906020019080838360005b838110156101e05780820151818401526020810190506101c5565b50505050905090810190601f16801561020d5780820380516001836020036101000a031916815260200191505b509250505060405180910390f35b610305600160008481526020019081526020016000208054600181600116156101000203166002900480601f0160208091040260200160405190810160405280929190818152602001828054600181600116156101000203166002900480156102c55780601f1061029a576101008083540402835291602001916102c5565b820191906000526020600020905b8154815290600101906020018083116102a857829003601f168201915b50505050506040805190810160405280601381526020017fefbc9b20e588b0e8bebee7ab99e782b9efbc9a00000000000000000000000000815250610571565b60016000848152602001908152602001600020908051906020019061032b929190610742565b506103e1600160008481526020019081526020016000208054600181600116156101000203166002900480601f0160208091040260200160405190810160405280929190818152602001828054600181600116156101000203166002900480156103d65780601f106103ab576101008083540402835291602001916103d6565b820191906000526020600020905b8154815290600101906020018083116103b957829003601f168201915b505050505082610571565b600160008481526020019081526020016000209080519060200190610407929190610742565b505050565b6060600160008381526020019081526020016000208054600181600116156101000203166002900480601f0160208091040260200160405190810160405280929190818152602001828054600181600116156101000203166002900480156104b55780601f1061048a576101008083540402835291602001916104b5565b820191906000526020600020905b81548152906001019060200180831161049857829003601f168201915b50505050509050919050565b60016020528060005260406000206000915090508054600181600116156101000203166002900480601f0160208091040260200160405190810160405280929190818152602001828054600181600116156101000203166002900480156105695780601f1061053e57610100808354040283529160200191610569565b820191906000526020600020905b81548152906001019060200180831161054c57829003601f168201915b505050505081565b606080606080606060008088955087945084518651016040519080825280601f01601f1916602001820160405280156105b95781602001602082028038833980820191505090505b50935083925060009150600090505b855181101561067b5785818151811015156105df57fe5b9060200101517f010000000000000000000000000000000000000000000000000000000000000090047f010000000000000000000000000000000000000000000000000000000000000002838380600101945081518110151561063e57fe5b9060200101907effffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff1916908160001a90535080806001019150506105c8565b600090505b845181101561073357848181518110151561069757fe5b9060200101517f010000000000000000000000000000000000000000000000000000000000000090047f01000000000000000000000000000000000000000000000000000000000000000283838060010194508151811015156106f657fe5b9060200101907effffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff1916908160001a9053508080600101915050610680565b83965050505050505092915050565b828054600181600116156101000203166002900490600052602060002090601f016020900481019282601f1061078357805160ff19168380011785556107b1565b828001600101855582156107b1579182015b828111156107b0578251825591602001919060010190610795565b5b5090506107be91906107c2565b5090565b6107e491905b808211156107e05760008160009055506001016107c8565b5090565b905600a165627a7a72305820cb70a5e7f03a01558a4f8042b7046bb3b7789d5997021d75c288253bcccf37de0029"

// DeployShareDeliveryInfo deploys a new contract, binding an instance of ShareDeliveryInfo to it.
func DeployShareDeliveryInfo(auth *bind.TransactOpts, backend bind.ContractBackend) (common.Address, *types.Transaction, *ShareDeliveryInfo, error) {
	parsed, err := abi.JSON(strings.NewReader(ShareDeliveryInfoABI))
	if err != nil {
		return common.Address{}, nil, nil, err
	}

	address, tx, contract, err := bind.DeployContract(auth, parsed, common.FromHex(ShareDeliveryInfoBin), backend)
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	return address, tx, &ShareDeliveryInfo{ShareDeliveryInfoCaller: ShareDeliveryInfoCaller{contract: contract}, ShareDeliveryInfoTransactor: ShareDeliveryInfoTransactor{contract: contract}, ShareDeliveryInfoFilterer: ShareDeliveryInfoFilterer{contract: contract}}, nil
}

// ShareDeliveryInfo is an auto generated Go binding around a Solidity contract.
type ShareDeliveryInfo struct {
	ShareDeliveryInfoCaller     // Read-only binding to the contract
	ShareDeliveryInfoTransactor // Write-only binding to the contract
	ShareDeliveryInfoFilterer   // Log filterer for contract events
}

// ShareDeliveryInfoCaller is an auto generated read-only Go binding around a Solidity contract.
type ShareDeliveryInfoCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// ShareDeliveryInfoTransactor is an auto generated write-only Go binding around a Solidity contract.
type ShareDeliveryInfoTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// ShareDeliveryInfoFilterer is an auto generated log filtering Go binding around a Solidity contract events.
type ShareDeliveryInfoFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// ShareDeliveryInfoSession is an auto generated Go binding around a Solidity contract,
// with pre-set call and transact options.
type ShareDeliveryInfoSession struct {
	Contract     *ShareDeliveryInfo // Generic contract binding to set the session for
	CallOpts     bind.CallOpts      // Call options to use throughout this session
	TransactOpts bind.TransactOpts  // Transaction auth options to use throughout this session
}

// ShareDeliveryInfoCallerSession is an auto generated read-only Go binding around a Solidity contract,
// with pre-set call options.
type ShareDeliveryInfoCallerSession struct {
	Contract *ShareDeliveryInfoCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts            // Call options to use throughout this session
}

// ShareDeliveryInfoTransactorSession is an auto generated write-only Go binding around a Solidity contract,
// with pre-set transact options.
type ShareDeliveryInfoTransactorSession struct {
	Contract     *ShareDeliveryInfoTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts            // Transaction auth options to use throughout this session
}

// ShareDeliveryInfoRaw is an auto generated low-level Go binding around a Solidity contract.
type ShareDeliveryInfoRaw struct {
	Contract *ShareDeliveryInfo // Generic contract binding to access the raw methods on
}

// ShareDeliveryInfoCallerRaw is an auto generated low-level read-only Go binding around a Solidity contract.
type ShareDeliveryInfoCallerRaw struct {
	Contract *ShareDeliveryInfoCaller // Generic read-only contract binding to access the raw methods on
}

// ShareDeliveryInfoTransactorRaw is an auto generated low-level write-only Go binding around a Solidity contract.
type ShareDeliveryInfoTransactorRaw struct {
	Contract *ShareDeliveryInfoTransactor // Generic write-only contract binding to access the raw methods on
}

// NewShareDeliveryInfo creates a new instance of ShareDeliveryInfo, bound to a specific deployed contract.
func NewShareDeliveryInfo(address common.Address, backend bind.ContractBackend) (*ShareDeliveryInfo, error) {
	contract, err := bindShareDeliveryInfo(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &ShareDeliveryInfo{ShareDeliveryInfoCaller: ShareDeliveryInfoCaller{contract: contract}, ShareDeliveryInfoTransactor: ShareDeliveryInfoTransactor{contract: contract}, ShareDeliveryInfoFilterer: ShareDeliveryInfoFilterer{contract: contract}}, nil
}

// NewShareDeliveryInfoCaller creates a new read-only instance of ShareDeliveryInfo, bound to a specific deployed contract.
func NewShareDeliveryInfoCaller(address common.Address, caller bind.ContractCaller) (*ShareDeliveryInfoCaller, error) {
	contract, err := bindShareDeliveryInfo(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &ShareDeliveryInfoCaller{contract: contract}, nil
}

// NewShareDeliveryInfoTransactor creates a new write-only instance of ShareDeliveryInfo, bound to a specific deployed contract.
func NewShareDeliveryInfoTransactor(address common.Address, transactor bind.ContractTransactor) (*ShareDeliveryInfoTransactor, error) {
	contract, err := bindShareDeliveryInfo(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &ShareDeliveryInfoTransactor{contract: contract}, nil
}

// NewShareDeliveryInfoFilterer creates a new log filterer instance of ShareDeliveryInfo, bound to a specific deployed contract.
func NewShareDeliveryInfoFilterer(address common.Address, filterer bind.ContractFilterer) (*ShareDeliveryInfoFilterer, error) {
	contract, err := bindShareDeliveryInfo(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &ShareDeliveryInfoFilterer{contract: contract}, nil
}

// bindShareDeliveryInfo binds a generic wrapper to an already deployed contract.
func bindShareDeliveryInfo(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(ShareDeliveryInfoABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_ShareDeliveryInfo *ShareDeliveryInfoRaw) Call(opts *bind.CallOpts, result interface{}, method string, params ...interface{}) error {
	return _ShareDeliveryInfo.Contract.ShareDeliveryInfoCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_ShareDeliveryInfo *ShareDeliveryInfoRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _ShareDeliveryInfo.Contract.ShareDeliveryInfoTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_ShareDeliveryInfo *ShareDeliveryInfoRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _ShareDeliveryInfo.Contract.ShareDeliveryInfoTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_ShareDeliveryInfo *ShareDeliveryInfoCallerRaw) Call(opts *bind.CallOpts, result interface{}, method string, params ...interface{}) error {
	return _ShareDeliveryInfo.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_ShareDeliveryInfo *ShareDeliveryInfoTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _ShareDeliveryInfo.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_ShareDeliveryInfo *ShareDeliveryInfoTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _ShareDeliveryInfo.Contract.contract.Transact(opts, method, params...)
}

// DeliveryPath is a free data retrieval call binding the contract method 0xa0afbd9f.
//
// Solidity: function delivery_path(uint256 ) constant returns(string)
func (_ShareDeliveryInfo *ShareDeliveryInfoCaller) DeliveryPath(opts *bind.CallOpts, arg0 *big.Int) (string, error) {
	var (
		ret0 = new(string)
	)
	out := ret0
	err := _ShareDeliveryInfo.contract.Call(opts, out, "delivery_path", arg0)
	return *ret0, err
}

// DeliveryPath is a free data retrieval call binding the contract method 0xa0afbd9f.
//
// Solidity: function delivery_path(uint256 ) constant returns(string)
func (_ShareDeliveryInfo *ShareDeliveryInfoSession) DeliveryPath(arg0 *big.Int) (string, error) {
	return _ShareDeliveryInfo.Contract.DeliveryPath(&_ShareDeliveryInfo.CallOpts, arg0)
}

// DeliveryPath is a free data retrieval call binding the contract method 0xa0afbd9f.
//
// Solidity: function delivery_path(uint256 ) constant returns(string)
func (_ShareDeliveryInfo *ShareDeliveryInfoCallerSession) DeliveryPath(arg0 *big.Int) (string, error) {
	return _ShareDeliveryInfo.Contract.DeliveryPath(&_ShareDeliveryInfo.CallOpts, arg0)
}

// Get is a free data retrieval call binding the contract method 0x9507d39a.
//
// Solidity: function get(uint256 _pkg_uid) constant returns(string)
func (_ShareDeliveryInfo *ShareDeliveryInfoCaller) Get(opts *bind.CallOpts, _pkg_uid *big.Int) (string, error) {
	var (
		ret0 = new(string)
	)
	out := ret0
	err := _ShareDeliveryInfo.contract.Call(opts, out, "get", _pkg_uid)
	return *ret0, err
}

// Get is a free data retrieval call binding the contract method 0x9507d39a.
//
// Solidity: function get(uint256 _pkg_uid) constant returns(string)
func (_ShareDeliveryInfo *ShareDeliveryInfoSession) Get(_pkg_uid *big.Int) (string, error) {
	return _ShareDeliveryInfo.Contract.Get(&_ShareDeliveryInfo.CallOpts, _pkg_uid)
}

// Get is a free data retrieval call binding the contract method 0x9507d39a.
//
// Solidity: function get(uint256 _pkg_uid) constant returns(string)
func (_ShareDeliveryInfo *ShareDeliveryInfoCallerSession) Get(_pkg_uid *big.Int) (string, error) {
	return _ShareDeliveryInfo.Contract.Get(&_ShareDeliveryInfo.CallOpts, _pkg_uid)
}

// Set is a paid mutator transaction binding the contract method 0x64371977.
//
// Solidity: function set(uint256 _pkg_uid, string station) returns()
func (_ShareDeliveryInfo *ShareDeliveryInfoTransactor) Set(opts *bind.TransactOpts, _pkg_uid *big.Int, station string) (*types.Transaction, error) {
	return _ShareDeliveryInfo.contract.Transact(opts, "set", _pkg_uid, station)
}

// Set is a paid mutator transaction binding the contract method 0x64371977.
//
// Solidity: function set(uint256 _pkg_uid, string station) returns()
func (_ShareDeliveryInfo *ShareDeliveryInfoSession) Set(_pkg_uid *big.Int, station string) (*types.Transaction, error) {
	return _ShareDeliveryInfo.Contract.Set(&_ShareDeliveryInfo.TransactOpts, _pkg_uid, station)
}

// Set is a paid mutator transaction binding the contract method 0x64371977.
//
// Solidity: function set(uint256 _pkg_uid, string station) returns()
func (_ShareDeliveryInfo *ShareDeliveryInfoTransactorSession) Set(_pkg_uid *big.Int, station string) (*types.Transaction, error) {
	return _ShareDeliveryInfo.Contract.Set(&_ShareDeliveryInfo.TransactOpts, _pkg_uid, station)
}
