// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package ethereum_handler

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

// PermittedMethodManagementContractPermittedMethod is an auto generated low-level Go binding around an user-defined struct.
type PermittedMethodManagementContractPermittedMethod struct {
	PermittedMethodId *big.Int
	Name              string
	Chaincode         string
	InputArgs         string
	OutputArgs        string
}

// PermittedMethodManagementMetaData contains all meta data concerning the PermittedMethodManagement contract.
var PermittedMethodManagementMetaData = &bind.MetaData{
	ABI: "[{\"inputs\":[{\"internalType\":\"string\",\"name\":\"name\",\"type\":\"string\"},{\"internalType\":\"string\",\"name\":\"chaincode\",\"type\":\"string\"},{\"internalType\":\"string\",\"name\":\"inputArgs\",\"type\":\"string\"},{\"internalType\":\"string\",\"name\":\"outputArgs\",\"type\":\"string\"}],\"name\":\"createPermittedMethod\",\"outputs\":[{\"components\":[{\"internalType\":\"uint256\",\"name\":\"permittedMethodId\",\"type\":\"uint256\"},{\"internalType\":\"string\",\"name\":\"name\",\"type\":\"string\"},{\"internalType\":\"string\",\"name\":\"chaincode\",\"type\":\"string\"},{\"internalType\":\"string\",\"name\":\"inputArgs\",\"type\":\"string\"},{\"internalType\":\"string\",\"name\":\"outputArgs\",\"type\":\"string\"}],\"internalType\":\"structPermittedMethodManagementContract.PermittedMethod\",\"name\":\"\",\"type\":\"tuple\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"permittedMethodId\",\"type\":\"uint256\"}],\"name\":\"deletePermittedMethod\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"existingPermittedMethods\",\"outputs\":[{\"internalType\":\"uint16\",\"name\":\"\",\"type\":\"uint16\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"generatePermittedMethodId\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"getAllPermittedMethods\",\"outputs\":[{\"components\":[{\"internalType\":\"uint256\",\"name\":\"permittedMethodId\",\"type\":\"uint256\"},{\"internalType\":\"string\",\"name\":\"name\",\"type\":\"string\"},{\"internalType\":\"string\",\"name\":\"chaincode\",\"type\":\"string\"},{\"internalType\":\"string\",\"name\":\"inputArgs\",\"type\":\"string\"},{\"internalType\":\"string\",\"name\":\"outputArgs\",\"type\":\"string\"}],\"internalType\":\"structPermittedMethodManagementContract.PermittedMethod[]\",\"name\":\"\",\"type\":\"tuple[]\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"permittedMethodId\",\"type\":\"uint256\"}],\"name\":\"getPermittedMethod\",\"outputs\":[{\"components\":[{\"internalType\":\"uint256\",\"name\":\"permittedMethodId\",\"type\":\"uint256\"},{\"internalType\":\"string\",\"name\":\"name\",\"type\":\"string\"},{\"internalType\":\"string\",\"name\":\"chaincode\",\"type\":\"string\"},{\"internalType\":\"string\",\"name\":\"inputArgs\",\"type\":\"string\"},{\"internalType\":\"string\",\"name\":\"outputArgs\",\"type\":\"string\"}],\"internalType\":\"structPermittedMethodManagementContract.PermittedMethod\",\"name\":\"\",\"type\":\"tuple\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"id\",\"type\":\"uint256\"}],\"name\":\"getPermittedMethodById\",\"outputs\":[{\"components\":[{\"internalType\":\"uint256\",\"name\":\"permittedMethodId\",\"type\":\"uint256\"},{\"internalType\":\"string\",\"name\":\"name\",\"type\":\"string\"},{\"internalType\":\"string\",\"name\":\"chaincode\",\"type\":\"string\"},{\"internalType\":\"string\",\"name\":\"inputArgs\",\"type\":\"string\"},{\"internalType\":\"string\",\"name\":\"outputArgs\",\"type\":\"string\"}],\"internalType\":\"structPermittedMethodManagementContract.PermittedMethod\",\"name\":\"\",\"type\":\"tuple\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"getPermittedMethodsNumber\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"accessibleNetworkId\",\"type\":\"uint256\"}],\"name\":\"isPermittedMethodExist\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"name\":\"permittedMethods\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"permittedMethodId\",\"type\":\"uint256\"},{\"internalType\":\"string\",\"name\":\"name\",\"type\":\"string\"},{\"internalType\":\"string\",\"name\":\"chaincode\",\"type\":\"string\"},{\"internalType\":\"string\",\"name\":\"inputArgs\",\"type\":\"string\"},{\"internalType\":\"string\",\"name\":\"outputArgs\",\"type\":\"string\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"name\":\"permittedMethodsIndexs\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"permittedMethodId\",\"type\":\"uint256\"},{\"internalType\":\"string\",\"name\":\"name\",\"type\":\"string\"},{\"internalType\":\"string\",\"name\":\"chaincode\",\"type\":\"string\"},{\"internalType\":\"string\",\"name\":\"inputArgs\",\"type\":\"string\"},{\"internalType\":\"string\",\"name\":\"outputArgs\",\"type\":\"string\"}],\"name\":\"updatePermittedMethod\",\"outputs\":[{\"components\":[{\"internalType\":\"uint256\",\"name\":\"permittedMethodId\",\"type\":\"uint256\"},{\"internalType\":\"string\",\"name\":\"name\",\"type\":\"string\"},{\"internalType\":\"string\",\"name\":\"chaincode\",\"type\":\"string\"},{\"internalType\":\"string\",\"name\":\"inputArgs\",\"type\":\"string\"},{\"internalType\":\"string\",\"name\":\"outputArgs\",\"type\":\"string\"}],\"internalType\":\"structPermittedMethodManagementContract.PermittedMethod\",\"name\":\"\",\"type\":\"tuple\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"}]",
}

// PermittedMethodManagementABI is the input ABI used to generate the binding from.
// Deprecated: Use PermittedMethodManagementMetaData.ABI instead.
var PermittedMethodManagementABI = PermittedMethodManagementMetaData.ABI

// PermittedMethodManagement is an auto generated Go binding around an Ethereum contract.
type PermittedMethodManagement struct {
	PermittedMethodManagementCaller     // Read-only binding to the contract
	PermittedMethodManagementTransactor // Write-only binding to the contract
	PermittedMethodManagementFilterer   // Log filterer for contract events
}

// PermittedMethodManagementCaller is an auto generated read-only Go binding around an Ethereum contract.
type PermittedMethodManagementCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// PermittedMethodManagementTransactor is an auto generated write-only Go binding around an Ethereum contract.
type PermittedMethodManagementTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// PermittedMethodManagementFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type PermittedMethodManagementFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// PermittedMethodManagementSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type PermittedMethodManagementSession struct {
	Contract     *PermittedMethodManagement // Generic contract binding to set the session for
	CallOpts     bind.CallOpts              // Call options to use throughout this session
	TransactOpts bind.TransactOpts          // Transaction auth options to use throughout this session
}

// PermittedMethodManagementCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type PermittedMethodManagementCallerSession struct {
	Contract *PermittedMethodManagementCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts                    // Call options to use throughout this session
}

// PermittedMethodManagementTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type PermittedMethodManagementTransactorSession struct {
	Contract     *PermittedMethodManagementTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts                    // Transaction auth options to use throughout this session
}

// PermittedMethodManagementRaw is an auto generated low-level Go binding around an Ethereum contract.
type PermittedMethodManagementRaw struct {
	Contract *PermittedMethodManagement // Generic contract binding to access the raw methods on
}

// PermittedMethodManagementCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type PermittedMethodManagementCallerRaw struct {
	Contract *PermittedMethodManagementCaller // Generic read-only contract binding to access the raw methods on
}

// PermittedMethodManagementTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type PermittedMethodManagementTransactorRaw struct {
	Contract *PermittedMethodManagementTransactor // Generic write-only contract binding to access the raw methods on
}

// NewPermittedMethodManagement creates a new instance of PermittedMethodManagement, bound to a specific deployed contract.
func NewPermittedMethodManagement(address common.Address, backend bind.ContractBackend) (*PermittedMethodManagement, error) {
	contract, err := bindPermittedMethodManagement(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &PermittedMethodManagement{PermittedMethodManagementCaller: PermittedMethodManagementCaller{contract: contract}, PermittedMethodManagementTransactor: PermittedMethodManagementTransactor{contract: contract}, PermittedMethodManagementFilterer: PermittedMethodManagementFilterer{contract: contract}}, nil
}

// NewPermittedMethodManagementCaller creates a new read-only instance of PermittedMethodManagement, bound to a specific deployed contract.
func NewPermittedMethodManagementCaller(address common.Address, caller bind.ContractCaller) (*PermittedMethodManagementCaller, error) {
	contract, err := bindPermittedMethodManagement(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &PermittedMethodManagementCaller{contract: contract}, nil
}

// NewPermittedMethodManagementTransactor creates a new write-only instance of PermittedMethodManagement, bound to a specific deployed contract.
func NewPermittedMethodManagementTransactor(address common.Address, transactor bind.ContractTransactor) (*PermittedMethodManagementTransactor, error) {
	contract, err := bindPermittedMethodManagement(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &PermittedMethodManagementTransactor{contract: contract}, nil
}

// NewPermittedMethodManagementFilterer creates a new log filterer instance of PermittedMethodManagement, bound to a specific deployed contract.
func NewPermittedMethodManagementFilterer(address common.Address, filterer bind.ContractFilterer) (*PermittedMethodManagementFilterer, error) {
	contract, err := bindPermittedMethodManagement(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &PermittedMethodManagementFilterer{contract: contract}, nil
}

// bindPermittedMethodManagement binds a generic wrapper to an already deployed contract.
func bindPermittedMethodManagement(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := PermittedMethodManagementMetaData.GetAbi()
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, *parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_PermittedMethodManagement *PermittedMethodManagementRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _PermittedMethodManagement.Contract.PermittedMethodManagementCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_PermittedMethodManagement *PermittedMethodManagementRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _PermittedMethodManagement.Contract.PermittedMethodManagementTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_PermittedMethodManagement *PermittedMethodManagementRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _PermittedMethodManagement.Contract.PermittedMethodManagementTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_PermittedMethodManagement *PermittedMethodManagementCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _PermittedMethodManagement.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_PermittedMethodManagement *PermittedMethodManagementTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _PermittedMethodManagement.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_PermittedMethodManagement *PermittedMethodManagementTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _PermittedMethodManagement.Contract.contract.Transact(opts, method, params...)
}

// ExistingPermittedMethods is a free data retrieval call binding the contract method 0xcd1d234b.
//
// Solidity: function existingPermittedMethods() view returns(uint16)
func (_PermittedMethodManagement *PermittedMethodManagementCaller) ExistingPermittedMethods(opts *bind.CallOpts) (uint16, error) {
	var out []interface{}
	err := _PermittedMethodManagement.contract.Call(opts, &out, "existingPermittedMethods")

	if err != nil {
		return *new(uint16), err
	}

	out0 := *abi.ConvertType(out[0], new(uint16)).(*uint16)

	return out0, err

}

// ExistingPermittedMethods is a free data retrieval call binding the contract method 0xcd1d234b.
//
// Solidity: function existingPermittedMethods() view returns(uint16)
func (_PermittedMethodManagement *PermittedMethodManagementSession) ExistingPermittedMethods() (uint16, error) {
	return _PermittedMethodManagement.Contract.ExistingPermittedMethods(&_PermittedMethodManagement.CallOpts)
}

// ExistingPermittedMethods is a free data retrieval call binding the contract method 0xcd1d234b.
//
// Solidity: function existingPermittedMethods() view returns(uint16)
func (_PermittedMethodManagement *PermittedMethodManagementCallerSession) ExistingPermittedMethods() (uint16, error) {
	return _PermittedMethodManagement.Contract.ExistingPermittedMethods(&_PermittedMethodManagement.CallOpts)
}

// GeneratePermittedMethodId is a free data retrieval call binding the contract method 0x3641d42d.
//
// Solidity: function generatePermittedMethodId() view returns(uint256)
func (_PermittedMethodManagement *PermittedMethodManagementCaller) GeneratePermittedMethodId(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _PermittedMethodManagement.contract.Call(opts, &out, "generatePermittedMethodId")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// GeneratePermittedMethodId is a free data retrieval call binding the contract method 0x3641d42d.
//
// Solidity: function generatePermittedMethodId() view returns(uint256)
func (_PermittedMethodManagement *PermittedMethodManagementSession) GeneratePermittedMethodId() (*big.Int, error) {
	return _PermittedMethodManagement.Contract.GeneratePermittedMethodId(&_PermittedMethodManagement.CallOpts)
}

// GeneratePermittedMethodId is a free data retrieval call binding the contract method 0x3641d42d.
//
// Solidity: function generatePermittedMethodId() view returns(uint256)
func (_PermittedMethodManagement *PermittedMethodManagementCallerSession) GeneratePermittedMethodId() (*big.Int, error) {
	return _PermittedMethodManagement.Contract.GeneratePermittedMethodId(&_PermittedMethodManagement.CallOpts)
}

// GetAllPermittedMethods is a free data retrieval call binding the contract method 0x77e34bc9.
//
// Solidity: function getAllPermittedMethods() view returns((uint256,string,string,string,string)[])
func (_PermittedMethodManagement *PermittedMethodManagementCaller) GetAllPermittedMethods(opts *bind.CallOpts) ([]PermittedMethodManagementContractPermittedMethod, error) {
	var out []interface{}
	err := _PermittedMethodManagement.contract.Call(opts, &out, "getAllPermittedMethods")

	if err != nil {
		return *new([]PermittedMethodManagementContractPermittedMethod), err
	}

	out0 := *abi.ConvertType(out[0], new([]PermittedMethodManagementContractPermittedMethod)).(*[]PermittedMethodManagementContractPermittedMethod)

	return out0, err

}

// GetAllPermittedMethods is a free data retrieval call binding the contract method 0x77e34bc9.
//
// Solidity: function getAllPermittedMethods() view returns((uint256,string,string,string,string)[])
func (_PermittedMethodManagement *PermittedMethodManagementSession) GetAllPermittedMethods() ([]PermittedMethodManagementContractPermittedMethod, error) {
	return _PermittedMethodManagement.Contract.GetAllPermittedMethods(&_PermittedMethodManagement.CallOpts)
}

// GetAllPermittedMethods is a free data retrieval call binding the contract method 0x77e34bc9.
//
// Solidity: function getAllPermittedMethods() view returns((uint256,string,string,string,string)[])
func (_PermittedMethodManagement *PermittedMethodManagementCallerSession) GetAllPermittedMethods() ([]PermittedMethodManagementContractPermittedMethod, error) {
	return _PermittedMethodManagement.Contract.GetAllPermittedMethods(&_PermittedMethodManagement.CallOpts)
}

// GetPermittedMethod is a free data retrieval call binding the contract method 0xdfead6f0.
//
// Solidity: function getPermittedMethod(uint256 permittedMethodId) view returns((uint256,string,string,string,string))
func (_PermittedMethodManagement *PermittedMethodManagementCaller) GetPermittedMethod(opts *bind.CallOpts, permittedMethodId *big.Int) (PermittedMethodManagementContractPermittedMethod, error) {
	var out []interface{}
	err := _PermittedMethodManagement.contract.Call(opts, &out, "getPermittedMethod", permittedMethodId)

	if err != nil {
		return *new(PermittedMethodManagementContractPermittedMethod), err
	}

	out0 := *abi.ConvertType(out[0], new(PermittedMethodManagementContractPermittedMethod)).(*PermittedMethodManagementContractPermittedMethod)

	return out0, err

}

// GetPermittedMethod is a free data retrieval call binding the contract method 0xdfead6f0.
//
// Solidity: function getPermittedMethod(uint256 permittedMethodId) view returns((uint256,string,string,string,string))
func (_PermittedMethodManagement *PermittedMethodManagementSession) GetPermittedMethod(permittedMethodId *big.Int) (PermittedMethodManagementContractPermittedMethod, error) {
	return _PermittedMethodManagement.Contract.GetPermittedMethod(&_PermittedMethodManagement.CallOpts, permittedMethodId)
}

// GetPermittedMethod is a free data retrieval call binding the contract method 0xdfead6f0.
//
// Solidity: function getPermittedMethod(uint256 permittedMethodId) view returns((uint256,string,string,string,string))
func (_PermittedMethodManagement *PermittedMethodManagementCallerSession) GetPermittedMethod(permittedMethodId *big.Int) (PermittedMethodManagementContractPermittedMethod, error) {
	return _PermittedMethodManagement.Contract.GetPermittedMethod(&_PermittedMethodManagement.CallOpts, permittedMethodId)
}

// GetPermittedMethodById is a free data retrieval call binding the contract method 0x975c7805.
//
// Solidity: function getPermittedMethodById(uint256 id) view returns((uint256,string,string,string,string))
func (_PermittedMethodManagement *PermittedMethodManagementCaller) GetPermittedMethodById(opts *bind.CallOpts, id *big.Int) (PermittedMethodManagementContractPermittedMethod, error) {
	var out []interface{}
	err := _PermittedMethodManagement.contract.Call(opts, &out, "getPermittedMethodById", id)

	if err != nil {
		return *new(PermittedMethodManagementContractPermittedMethod), err
	}

	out0 := *abi.ConvertType(out[0], new(PermittedMethodManagementContractPermittedMethod)).(*PermittedMethodManagementContractPermittedMethod)

	return out0, err

}

// GetPermittedMethodById is a free data retrieval call binding the contract method 0x975c7805.
//
// Solidity: function getPermittedMethodById(uint256 id) view returns((uint256,string,string,string,string))
func (_PermittedMethodManagement *PermittedMethodManagementSession) GetPermittedMethodById(id *big.Int) (PermittedMethodManagementContractPermittedMethod, error) {
	return _PermittedMethodManagement.Contract.GetPermittedMethodById(&_PermittedMethodManagement.CallOpts, id)
}

// GetPermittedMethodById is a free data retrieval call binding the contract method 0x975c7805.
//
// Solidity: function getPermittedMethodById(uint256 id) view returns((uint256,string,string,string,string))
func (_PermittedMethodManagement *PermittedMethodManagementCallerSession) GetPermittedMethodById(id *big.Int) (PermittedMethodManagementContractPermittedMethod, error) {
	return _PermittedMethodManagement.Contract.GetPermittedMethodById(&_PermittedMethodManagement.CallOpts, id)
}

// GetPermittedMethodsNumber is a free data retrieval call binding the contract method 0x901ae07e.
//
// Solidity: function getPermittedMethodsNumber() view returns(uint256)
func (_PermittedMethodManagement *PermittedMethodManagementCaller) GetPermittedMethodsNumber(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _PermittedMethodManagement.contract.Call(opts, &out, "getPermittedMethodsNumber")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// GetPermittedMethodsNumber is a free data retrieval call binding the contract method 0x901ae07e.
//
// Solidity: function getPermittedMethodsNumber() view returns(uint256)
func (_PermittedMethodManagement *PermittedMethodManagementSession) GetPermittedMethodsNumber() (*big.Int, error) {
	return _PermittedMethodManagement.Contract.GetPermittedMethodsNumber(&_PermittedMethodManagement.CallOpts)
}

// GetPermittedMethodsNumber is a free data retrieval call binding the contract method 0x901ae07e.
//
// Solidity: function getPermittedMethodsNumber() view returns(uint256)
func (_PermittedMethodManagement *PermittedMethodManagementCallerSession) GetPermittedMethodsNumber() (*big.Int, error) {
	return _PermittedMethodManagement.Contract.GetPermittedMethodsNumber(&_PermittedMethodManagement.CallOpts)
}

// IsPermittedMethodExist is a free data retrieval call binding the contract method 0x355d9ea1.
//
// Solidity: function isPermittedMethodExist(uint256 accessibleNetworkId) view returns(bool)
func (_PermittedMethodManagement *PermittedMethodManagementCaller) IsPermittedMethodExist(opts *bind.CallOpts, accessibleNetworkId *big.Int) (bool, error) {
	var out []interface{}
	err := _PermittedMethodManagement.contract.Call(opts, &out, "isPermittedMethodExist", accessibleNetworkId)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// IsPermittedMethodExist is a free data retrieval call binding the contract method 0x355d9ea1.
//
// Solidity: function isPermittedMethodExist(uint256 accessibleNetworkId) view returns(bool)
func (_PermittedMethodManagement *PermittedMethodManagementSession) IsPermittedMethodExist(accessibleNetworkId *big.Int) (bool, error) {
	return _PermittedMethodManagement.Contract.IsPermittedMethodExist(&_PermittedMethodManagement.CallOpts, accessibleNetworkId)
}

// IsPermittedMethodExist is a free data retrieval call binding the contract method 0x355d9ea1.
//
// Solidity: function isPermittedMethodExist(uint256 accessibleNetworkId) view returns(bool)
func (_PermittedMethodManagement *PermittedMethodManagementCallerSession) IsPermittedMethodExist(accessibleNetworkId *big.Int) (bool, error) {
	return _PermittedMethodManagement.Contract.IsPermittedMethodExist(&_PermittedMethodManagement.CallOpts, accessibleNetworkId)
}

// PermittedMethods is a free data retrieval call binding the contract method 0xf7295dfa.
//
// Solidity: function permittedMethods(uint256 ) view returns(uint256 permittedMethodId, string name, string chaincode, string inputArgs, string outputArgs)
func (_PermittedMethodManagement *PermittedMethodManagementCaller) PermittedMethods(opts *bind.CallOpts, arg0 *big.Int) (struct {
	PermittedMethodId *big.Int
	Name              string
	Chaincode         string
	InputArgs         string
	OutputArgs        string
}, error) {
	var out []interface{}
	err := _PermittedMethodManagement.contract.Call(opts, &out, "permittedMethods", arg0)

	outstruct := new(struct {
		PermittedMethodId *big.Int
		Name              string
		Chaincode         string
		InputArgs         string
		OutputArgs        string
	})
	if err != nil {
		return *outstruct, err
	}

	outstruct.PermittedMethodId = *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)
	outstruct.Name = *abi.ConvertType(out[1], new(string)).(*string)
	outstruct.Chaincode = *abi.ConvertType(out[2], new(string)).(*string)
	outstruct.InputArgs = *abi.ConvertType(out[3], new(string)).(*string)
	outstruct.OutputArgs = *abi.ConvertType(out[4], new(string)).(*string)

	return *outstruct, err

}

// PermittedMethods is a free data retrieval call binding the contract method 0xf7295dfa.
//
// Solidity: function permittedMethods(uint256 ) view returns(uint256 permittedMethodId, string name, string chaincode, string inputArgs, string outputArgs)
func (_PermittedMethodManagement *PermittedMethodManagementSession) PermittedMethods(arg0 *big.Int) (struct {
	PermittedMethodId *big.Int
	Name              string
	Chaincode         string
	InputArgs         string
	OutputArgs        string
}, error) {
	return _PermittedMethodManagement.Contract.PermittedMethods(&_PermittedMethodManagement.CallOpts, arg0)
}

// PermittedMethods is a free data retrieval call binding the contract method 0xf7295dfa.
//
// Solidity: function permittedMethods(uint256 ) view returns(uint256 permittedMethodId, string name, string chaincode, string inputArgs, string outputArgs)
func (_PermittedMethodManagement *PermittedMethodManagementCallerSession) PermittedMethods(arg0 *big.Int) (struct {
	PermittedMethodId *big.Int
	Name              string
	Chaincode         string
	InputArgs         string
	OutputArgs        string
}, error) {
	return _PermittedMethodManagement.Contract.PermittedMethods(&_PermittedMethodManagement.CallOpts, arg0)
}

// PermittedMethodsIndexs is a free data retrieval call binding the contract method 0xa06e3fb6.
//
// Solidity: function permittedMethodsIndexs(uint256 ) view returns(uint256)
func (_PermittedMethodManagement *PermittedMethodManagementCaller) PermittedMethodsIndexs(opts *bind.CallOpts, arg0 *big.Int) (*big.Int, error) {
	var out []interface{}
	err := _PermittedMethodManagement.contract.Call(opts, &out, "permittedMethodsIndexs", arg0)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// PermittedMethodsIndexs is a free data retrieval call binding the contract method 0xa06e3fb6.
//
// Solidity: function permittedMethodsIndexs(uint256 ) view returns(uint256)
func (_PermittedMethodManagement *PermittedMethodManagementSession) PermittedMethodsIndexs(arg0 *big.Int) (*big.Int, error) {
	return _PermittedMethodManagement.Contract.PermittedMethodsIndexs(&_PermittedMethodManagement.CallOpts, arg0)
}

// PermittedMethodsIndexs is a free data retrieval call binding the contract method 0xa06e3fb6.
//
// Solidity: function permittedMethodsIndexs(uint256 ) view returns(uint256)
func (_PermittedMethodManagement *PermittedMethodManagementCallerSession) PermittedMethodsIndexs(arg0 *big.Int) (*big.Int, error) {
	return _PermittedMethodManagement.Contract.PermittedMethodsIndexs(&_PermittedMethodManagement.CallOpts, arg0)
}

// CreatePermittedMethod is a paid mutator transaction binding the contract method 0x08fc64e3.
//
// Solidity: function createPermittedMethod(string name, string chaincode, string inputArgs, string outputArgs) returns((uint256,string,string,string,string))
func (_PermittedMethodManagement *PermittedMethodManagementTransactor) CreatePermittedMethod(opts *bind.TransactOpts, name string, chaincode string, inputArgs string, outputArgs string) (*types.Transaction, error) {
	return _PermittedMethodManagement.contract.Transact(opts, "createPermittedMethod", name, chaincode, inputArgs, outputArgs)
}

// CreatePermittedMethod is a paid mutator transaction binding the contract method 0x08fc64e3.
//
// Solidity: function createPermittedMethod(string name, string chaincode, string inputArgs, string outputArgs) returns((uint256,string,string,string,string))
func (_PermittedMethodManagement *PermittedMethodManagementSession) CreatePermittedMethod(name string, chaincode string, inputArgs string, outputArgs string) (*types.Transaction, error) {
	return _PermittedMethodManagement.Contract.CreatePermittedMethod(&_PermittedMethodManagement.TransactOpts, name, chaincode, inputArgs, outputArgs)
}

// CreatePermittedMethod is a paid mutator transaction binding the contract method 0x08fc64e3.
//
// Solidity: function createPermittedMethod(string name, string chaincode, string inputArgs, string outputArgs) returns((uint256,string,string,string,string))
func (_PermittedMethodManagement *PermittedMethodManagementTransactorSession) CreatePermittedMethod(name string, chaincode string, inputArgs string, outputArgs string) (*types.Transaction, error) {
	return _PermittedMethodManagement.Contract.CreatePermittedMethod(&_PermittedMethodManagement.TransactOpts, name, chaincode, inputArgs, outputArgs)
}

// DeletePermittedMethod is a paid mutator transaction binding the contract method 0xf222f2de.
//
// Solidity: function deletePermittedMethod(uint256 permittedMethodId) returns()
func (_PermittedMethodManagement *PermittedMethodManagementTransactor) DeletePermittedMethod(opts *bind.TransactOpts, permittedMethodId *big.Int) (*types.Transaction, error) {
	return _PermittedMethodManagement.contract.Transact(opts, "deletePermittedMethod", permittedMethodId)
}

// DeletePermittedMethod is a paid mutator transaction binding the contract method 0xf222f2de.
//
// Solidity: function deletePermittedMethod(uint256 permittedMethodId) returns()
func (_PermittedMethodManagement *PermittedMethodManagementSession) DeletePermittedMethod(permittedMethodId *big.Int) (*types.Transaction, error) {
	return _PermittedMethodManagement.Contract.DeletePermittedMethod(&_PermittedMethodManagement.TransactOpts, permittedMethodId)
}

// DeletePermittedMethod is a paid mutator transaction binding the contract method 0xf222f2de.
//
// Solidity: function deletePermittedMethod(uint256 permittedMethodId) returns()
func (_PermittedMethodManagement *PermittedMethodManagementTransactorSession) DeletePermittedMethod(permittedMethodId *big.Int) (*types.Transaction, error) {
	return _PermittedMethodManagement.Contract.DeletePermittedMethod(&_PermittedMethodManagement.TransactOpts, permittedMethodId)
}

// UpdatePermittedMethod is a paid mutator transaction binding the contract method 0x8d2634c8.
//
// Solidity: function updatePermittedMethod(uint256 permittedMethodId, string name, string chaincode, string inputArgs, string outputArgs) returns((uint256,string,string,string,string))
func (_PermittedMethodManagement *PermittedMethodManagementTransactor) UpdatePermittedMethod(opts *bind.TransactOpts, permittedMethodId *big.Int, name string, chaincode string, inputArgs string, outputArgs string) (*types.Transaction, error) {
	return _PermittedMethodManagement.contract.Transact(opts, "updatePermittedMethod", permittedMethodId, name, chaincode, inputArgs, outputArgs)
}

// UpdatePermittedMethod is a paid mutator transaction binding the contract method 0x8d2634c8.
//
// Solidity: function updatePermittedMethod(uint256 permittedMethodId, string name, string chaincode, string inputArgs, string outputArgs) returns((uint256,string,string,string,string))
func (_PermittedMethodManagement *PermittedMethodManagementSession) UpdatePermittedMethod(permittedMethodId *big.Int, name string, chaincode string, inputArgs string, outputArgs string) (*types.Transaction, error) {
	return _PermittedMethodManagement.Contract.UpdatePermittedMethod(&_PermittedMethodManagement.TransactOpts, permittedMethodId, name, chaincode, inputArgs, outputArgs)
}

// UpdatePermittedMethod is a paid mutator transaction binding the contract method 0x8d2634c8.
//
// Solidity: function updatePermittedMethod(uint256 permittedMethodId, string name, string chaincode, string inputArgs, string outputArgs) returns((uint256,string,string,string,string))
func (_PermittedMethodManagement *PermittedMethodManagementTransactorSession) UpdatePermittedMethod(permittedMethodId *big.Int, name string, chaincode string, inputArgs string, outputArgs string) (*types.Transaction, error) {
	return _PermittedMethodManagement.Contract.UpdatePermittedMethod(&_PermittedMethodManagement.TransactOpts, permittedMethodId, name, chaincode, inputArgs, outputArgs)
}
