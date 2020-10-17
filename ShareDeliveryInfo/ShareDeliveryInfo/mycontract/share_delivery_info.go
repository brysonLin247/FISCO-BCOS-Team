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
const ShareDeliveryInfoABI = "[{\"constant\":false,\"inputs\":[{\"name\":\"_pkg_uid\",\"type\":\"uint64\"},{\"name\":\"station\",\"type\":\"string\"}],\"name\":\"set\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"name\":\"\",\"type\":\"uint64\"}],\"name\":\"delivery_path\",\"outputs\":[{\"name\":\"\",\"type\":\"string\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"name\":\"_pkg_uid\",\"type\":\"uint64\"}],\"name\":\"get\",\"outputs\":[{\"name\":\"\",\"type\":\"string\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"constructor\"}]"

// ShareDeliveryInfoBin is the compiled bytecode used for deploying new contracts.
var ShareDeliveryInfoBin = "0x608060405234801561001057600080fd5b50610894806100206000396000f300608060405260043610610056576000357c0100000000000000000000000000000000000000000000000000000000900463ffffffff168062e19f951461005b5780638b117e13146100d8578063ada8679814610188575b600080fd5b34801561006757600080fd5b506100d6600480360381019080803567ffffffffffffffff169060200190929190803590602001908201803590602001908080601f0160208091040260200160405190810160405280939291908181526020018383808284378201915050505050509192919290505050610238565b005b3480156100e457600080fd5b5061010d600480360381019080803567ffffffffffffffff169060200190929190505050610479565b6040518080602001828103825283818151815260200191508051906020019080838360005b8381101561014d578082015181840152602081019050610132565b50505050905090810190601f16801561017a5780820380516001836020036101000a031916815260200191505b509250505060405180910390f35b34801561019457600080fd5b506101bd600480360381019080803567ffffffffffffffff169060200190929190505050610529565b6040518080602001828103825283818151815260200191508051906020019080838360005b838110156101fd5780820151818401526020810190506101e2565b50505050905090810190601f16801561022a5780820380516001836020036101000a031916815260200191505b509250505060405180910390f35b610336600160008467ffffffffffffffff1667ffffffffffffffff1681526020019081526020016000208054600181600116156101000203166002900480601f0160208091040260200160405190810160405280929190818152602001828054600181600116156101000203166002900480156102f65780601f106102cb576101008083540402835291602001916102f6565b820191906000526020600020905b8154815290600101906020018083116102d957829003601f168201915b50505050506040805190810160405280601381526020017fefbc9b20e588b0e8bebee7ab99e782b9efbc9a000000000000000000000000008152506105f2565b600160008467ffffffffffffffff1667ffffffffffffffff16815260200190815260200160002090805190602001906103709291906107c3565b5061043a600160008467ffffffffffffffff1667ffffffffffffffff1681526020019081526020016000208054600181600116156101000203166002900480601f01602080910402602001604051908101604052809291908181526020018280546001816001161561010002031660029004801561042f5780601f106104045761010080835404028352916020019161042f565b820191906000526020600020905b81548152906001019060200180831161041257829003601f168201915b5050505050826105f2565b600160008467ffffffffffffffff1667ffffffffffffffff16815260200190815260200160002090805190602001906104749291906107c3565b505050565b60016020528060005260406000206000915090508054600181600116156101000203166002900480601f0160208091040260200160405190810160405280929190818152602001828054600181600116156101000203166002900480156105215780601f106104f657610100808354040283529160200191610521565b820191906000526020600020905b81548152906001019060200180831161050457829003601f168201915b505050505081565b6060600160008367ffffffffffffffff1667ffffffffffffffff1681526020019081526020016000208054600181600116156101000203166002900480601f0160208091040260200160405190810160405280929190818152602001828054600181600116156101000203166002900480156105e65780601f106105bb576101008083540402835291602001916105e6565b820191906000526020600020905b8154815290600101906020018083116105c957829003601f168201915b50505050509050919050565b606080606080606060008088955087945084518651016040519080825280601f01601f19166020018201604052801561063a5781602001602082028038833980820191505090505b50935083925060009150600090505b85518110156106fc57858181518110151561066057fe5b9060200101517f010000000000000000000000000000000000000000000000000000000000000090047f01000000000000000000000000000000000000000000000000000000000000000283838060010194508151811015156106bf57fe5b9060200101907effffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff1916908160001a9053508080600101915050610649565b600090505b84518110156107b457848181518110151561071857fe5b9060200101517f010000000000000000000000000000000000000000000000000000000000000090047f010000000000000000000000000000000000000000000000000000000000000002838380600101945081518110151561077757fe5b9060200101907effffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff1916908160001a9053508080600101915050610701565b83965050505050505092915050565b828054600181600116156101000203166002900490600052602060002090601f016020900481019282601f1061080457805160ff1916838001178555610832565b82800160010185558215610832579182015b82811115610831578251825591602001919060010190610816565b5b50905061083f9190610843565b5090565b61086591905b80821115610861576000816000905550600101610849565b5090565b905600a165627a7a723058207db2e6a4f9dbac460f40a3a39df9e6ca70df95ac8461761a5e106fd81e26e96c0029"

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

// DeliveryPath is a free data retrieval call binding the contract method 0x8b117e13.
//
// Solidity: function delivery_path(uint64 ) constant returns(string)
func (_ShareDeliveryInfo *ShareDeliveryInfoCaller) DeliveryPath(opts *bind.CallOpts, arg0 uint64) (string, error) {
	var (
		ret0 = new(string)
	)
	out := ret0
	err := _ShareDeliveryInfo.contract.Call(opts, out, "delivery_path", arg0)
	return *ret0, err
}

// DeliveryPath is a free data retrieval call binding the contract method 0x8b117e13.
//
// Solidity: function delivery_path(uint64 ) constant returns(string)
func (_ShareDeliveryInfo *ShareDeliveryInfoSession) DeliveryPath(arg0 uint64) (string, error) {
	return _ShareDeliveryInfo.Contract.DeliveryPath(&_ShareDeliveryInfo.CallOpts, arg0)
}

// DeliveryPath is a free data retrieval call binding the contract method 0x8b117e13.
//
// Solidity: function delivery_path(uint64 ) constant returns(string)
func (_ShareDeliveryInfo *ShareDeliveryInfoCallerSession) DeliveryPath(arg0 uint64) (string, error) {
	return _ShareDeliveryInfo.Contract.DeliveryPath(&_ShareDeliveryInfo.CallOpts, arg0)
}

// Get is a free data retrieval call binding the contract method 0xada86798.
//
// Solidity: function get(uint64 _pkg_uid) constant returns(string)
func (_ShareDeliveryInfo *ShareDeliveryInfoCaller) Get(opts *bind.CallOpts, _pkg_uid uint64) (string, error) {
	var (
		ret0 = new(string)
	)
	out := ret0
	err := _ShareDeliveryInfo.contract.Call(opts, out, "get", _pkg_uid)
	return *ret0, err
}

// Get is a free data retrieval call binding the contract method 0xada86798.
//
// Solidity: function get(uint64 _pkg_uid) constant returns(string)
func (_ShareDeliveryInfo *ShareDeliveryInfoSession) Get(_pkg_uid uint64) (string, error) {
	return _ShareDeliveryInfo.Contract.Get(&_ShareDeliveryInfo.CallOpts, _pkg_uid)
}

// Get is a free data retrieval call binding the contract method 0xada86798.
//
// Solidity: function get(uint64 _pkg_uid) constant returns(string)
func (_ShareDeliveryInfo *ShareDeliveryInfoCallerSession) Get(_pkg_uid uint64) (string, error) {
	return _ShareDeliveryInfo.Contract.Get(&_ShareDeliveryInfo.CallOpts, _pkg_uid)
}

// Set is a paid mutator transaction binding the contract method 0x00e19f95.
//
// Solidity: function set(uint64 _pkg_uid, string station) returns()
func (_ShareDeliveryInfo *ShareDeliveryInfoTransactor) Set(opts *bind.TransactOpts, _pkg_uid uint64, station string) (*types.Transaction, error) {
	return _ShareDeliveryInfo.contract.Transact(opts, "set", _pkg_uid, station)
}

// Set is a paid mutator transaction binding the contract method 0x00e19f95.
//
// Solidity: function set(uint64 _pkg_uid, string station) returns()
func (_ShareDeliveryInfo *ShareDeliveryInfoSession) Set(_pkg_uid uint64, station string) (*types.Transaction, error) {
	return _ShareDeliveryInfo.Contract.Set(&_ShareDeliveryInfo.TransactOpts, _pkg_uid, station)
}

// Set is a paid mutator transaction binding the contract method 0x00e19f95.
//
// Solidity: function set(uint64 _pkg_uid, string station) returns()
func (_ShareDeliveryInfo *ShareDeliveryInfoTransactorSession) Set(_pkg_uid uint64, station string) (*types.Transaction, error) {
	return _ShareDeliveryInfo.Contract.Set(&_ShareDeliveryInfo.TransactOpts, _pkg_uid, station)
}
