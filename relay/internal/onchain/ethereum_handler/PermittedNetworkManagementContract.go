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

// PermittedNetworkManagementContractPermittedNetwork is an auto generated low-level Go binding around an user-defined struct.
type PermittedNetworkManagementContractPermittedNetwork struct {
	NetworkName        string
	NetworkIP          string
	NetworkAddress     string
	CompanyName        string
	PermittedNetworkId *big.Int
}

// PermittedNetworkManagementMetaData contains all meta data concerning the PermittedNetworkManagement contract.
var PermittedNetworkManagementMetaData = &bind.MetaData{
	ABI: "[{\"inputs\":[{\"internalType\":\"string\",\"name\":\"networkName\",\"type\":\"string\"},{\"internalType\":\"string\",\"name\":\"networkIP\",\"type\":\"string\"},{\"internalType\":\"string\",\"name\":\"networkAddress\",\"type\":\"string\"},{\"internalType\":\"string\",\"name\":\"companyName\",\"type\":\"string\"}],\"name\":\"createPermittedNetwork\",\"outputs\":[{\"components\":[{\"internalType\":\"string\",\"name\":\"networkName\",\"type\":\"string\"},{\"internalType\":\"string\",\"name\":\"networkIP\",\"type\":\"string\"},{\"internalType\":\"string\",\"name\":\"networkAddress\",\"type\":\"string\"},{\"internalType\":\"string\",\"name\":\"companyName\",\"type\":\"string\"},{\"internalType\":\"uint256\",\"name\":\"permittedNetworkId\",\"type\":\"uint256\"}],\"internalType\":\"structPermittedNetworkManagementContract.PermittedNetwork\",\"name\":\"\",\"type\":\"tuple\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"permittedNetworkId\",\"type\":\"uint256\"}],\"name\":\"deletePermittedNetwork\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"existingPermittedNetworks\",\"outputs\":[{\"internalType\":\"uint16\",\"name\":\"\",\"type\":\"uint16\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"generatePermittedNetworkId\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"getAllPermittedNetworks\",\"outputs\":[{\"components\":[{\"internalType\":\"string\",\"name\":\"networkName\",\"type\":\"string\"},{\"internalType\":\"string\",\"name\":\"networkIP\",\"type\":\"string\"},{\"internalType\":\"string\",\"name\":\"networkAddress\",\"type\":\"string\"},{\"internalType\":\"string\",\"name\":\"companyName\",\"type\":\"string\"},{\"internalType\":\"uint256\",\"name\":\"permittedNetworkId\",\"type\":\"uint256\"}],\"internalType\":\"structPermittedNetworkManagementContract.PermittedNetwork[]\",\"name\":\"\",\"type\":\"tuple[]\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"permittedNetworkId\",\"type\":\"uint256\"}],\"name\":\"getPermittedNetwork\",\"outputs\":[{\"components\":[{\"internalType\":\"string\",\"name\":\"networkName\",\"type\":\"string\"},{\"internalType\":\"string\",\"name\":\"networkIP\",\"type\":\"string\"},{\"internalType\":\"string\",\"name\":\"networkAddress\",\"type\":\"string\"},{\"internalType\":\"string\",\"name\":\"companyName\",\"type\":\"string\"},{\"internalType\":\"uint256\",\"name\":\"permittedNetworkId\",\"type\":\"uint256\"}],\"internalType\":\"structPermittedNetworkManagementContract.PermittedNetwork\",\"name\":\"\",\"type\":\"tuple\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"getPermittedNetworkIndexArray\",\"outputs\":[{\"internalType\":\"uint256[]\",\"name\":\"\",\"type\":\"uint256[]\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"getPermittedNetworksNumber\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"permittedNetworkId\",\"type\":\"uint256\"}],\"name\":\"permittedNetworkExists\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"name\":\"permittedNetworkIndexs\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"name\":\"permittedNetworks\",\"outputs\":[{\"internalType\":\"string\",\"name\":\"networkName\",\"type\":\"string\"},{\"internalType\":\"string\",\"name\":\"networkIP\",\"type\":\"string\"},{\"internalType\":\"string\",\"name\":\"networkAddress\",\"type\":\"string\"},{\"internalType\":\"string\",\"name\":\"companyName\",\"type\":\"string\"},{\"internalType\":\"uint256\",\"name\":\"permittedNetworkId\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"permittedNetworkId\",\"type\":\"uint256\"},{\"internalType\":\"string\",\"name\":\"networkName\",\"type\":\"string\"},{\"internalType\":\"string\",\"name\":\"networkIP\",\"type\":\"string\"},{\"internalType\":\"string\",\"name\":\"networkAddress\",\"type\":\"string\"},{\"internalType\":\"string\",\"name\":\"companyName\",\"type\":\"string\"}],\"name\":\"updatePermittedNetwork\",\"outputs\":[{\"components\":[{\"internalType\":\"string\",\"name\":\"networkName\",\"type\":\"string\"},{\"internalType\":\"string\",\"name\":\"networkIP\",\"type\":\"string\"},{\"internalType\":\"string\",\"name\":\"networkAddress\",\"type\":\"string\"},{\"internalType\":\"string\",\"name\":\"companyName\",\"type\":\"string\"},{\"internalType\":\"uint256\",\"name\":\"permittedNetworkId\",\"type\":\"uint256\"}],\"internalType\":\"structPermittedNetworkManagementContract.PermittedNetwork\",\"name\":\"\",\"type\":\"tuple\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"}]",
}

// PermittedNetworkManagementABI is the input ABI used to generate the binding from.
// Deprecated: Use PermittedNetworkManagementMetaData.ABI instead.
var PermittedNetworkManagementABI = PermittedNetworkManagementMetaData.ABI

// PermittedNetworkManagement is an auto generated Go binding around an Ethereum contract.
type PermittedNetworkManagement struct {
	PermittedNetworkManagementCaller     // Read-only binding to the contract
	PermittedNetworkManagementTransactor // Write-only binding to the contract
	PermittedNetworkManagementFilterer   // Log filterer for contract events
}

// PermittedNetworkManagementCaller is an auto generated read-only Go binding around an Ethereum contract.
type PermittedNetworkManagementCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// PermittedNetworkManagementTransactor is an auto generated write-only Go binding around an Ethereum contract.
type PermittedNetworkManagementTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// PermittedNetworkManagementFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type PermittedNetworkManagementFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// PermittedNetworkManagementSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type PermittedNetworkManagementSession struct {
	Contract     *PermittedNetworkManagement // Generic contract binding to set the session for
	CallOpts     bind.CallOpts               // Call options to use throughout this session
	TransactOpts bind.TransactOpts           // Transaction auth options to use throughout this session
}

// PermittedNetworkManagementCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type PermittedNetworkManagementCallerSession struct {
	Contract *PermittedNetworkManagementCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts                     // Call options to use throughout this session
}

// PermittedNetworkManagementTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type PermittedNetworkManagementTransactorSession struct {
	Contract     *PermittedNetworkManagementTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts                     // Transaction auth options to use throughout this session
}

// PermittedNetworkManagementRaw is an auto generated low-level Go binding around an Ethereum contract.
type PermittedNetworkManagementRaw struct {
	Contract *PermittedNetworkManagement // Generic contract binding to access the raw methods on
}

// PermittedNetworkManagementCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type PermittedNetworkManagementCallerRaw struct {
	Contract *PermittedNetworkManagementCaller // Generic read-only contract binding to access the raw methods on
}

// PermittedNetworkManagementTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type PermittedNetworkManagementTransactorRaw struct {
	Contract *PermittedNetworkManagementTransactor // Generic write-only contract binding to access the raw methods on
}

// NewPermittedNetworkManagement creates a new instance of PermittedNetworkManagement, bound to a specific deployed contract.
func NewPermittedNetworkManagement(address common.Address, backend bind.ContractBackend) (*PermittedNetworkManagement, error) {
	contract, err := bindPermittedNetworkManagement(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &PermittedNetworkManagement{PermittedNetworkManagementCaller: PermittedNetworkManagementCaller{contract: contract}, PermittedNetworkManagementTransactor: PermittedNetworkManagementTransactor{contract: contract}, PermittedNetworkManagementFilterer: PermittedNetworkManagementFilterer{contract: contract}}, nil
}

// NewPermittedNetworkManagementCaller creates a new read-only instance of PermittedNetworkManagement, bound to a specific deployed contract.
func NewPermittedNetworkManagementCaller(address common.Address, caller bind.ContractCaller) (*PermittedNetworkManagementCaller, error) {
	contract, err := bindPermittedNetworkManagement(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &PermittedNetworkManagementCaller{contract: contract}, nil
}

// NewPermittedNetworkManagementTransactor creates a new write-only instance of PermittedNetworkManagement, bound to a specific deployed contract.
func NewPermittedNetworkManagementTransactor(address common.Address, transactor bind.ContractTransactor) (*PermittedNetworkManagementTransactor, error) {
	contract, err := bindPermittedNetworkManagement(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &PermittedNetworkManagementTransactor{contract: contract}, nil
}

// NewPermittedNetworkManagementFilterer creates a new log filterer instance of PermittedNetworkManagement, bound to a specific deployed contract.
func NewPermittedNetworkManagementFilterer(address common.Address, filterer bind.ContractFilterer) (*PermittedNetworkManagementFilterer, error) {
	contract, err := bindPermittedNetworkManagement(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &PermittedNetworkManagementFilterer{contract: contract}, nil
}

// bindPermittedNetworkManagement binds a generic wrapper to an already deployed contract.
func bindPermittedNetworkManagement(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := PermittedNetworkManagementMetaData.GetAbi()
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, *parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_PermittedNetworkManagement *PermittedNetworkManagementRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _PermittedNetworkManagement.Contract.PermittedNetworkManagementCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_PermittedNetworkManagement *PermittedNetworkManagementRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _PermittedNetworkManagement.Contract.PermittedNetworkManagementTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_PermittedNetworkManagement *PermittedNetworkManagementRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _PermittedNetworkManagement.Contract.PermittedNetworkManagementTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_PermittedNetworkManagement *PermittedNetworkManagementCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _PermittedNetworkManagement.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_PermittedNetworkManagement *PermittedNetworkManagementTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _PermittedNetworkManagement.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_PermittedNetworkManagement *PermittedNetworkManagementTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _PermittedNetworkManagement.Contract.contract.Transact(opts, method, params...)
}

// ExistingPermittedNetworks is a free data retrieval call binding the contract method 0x7419fe64.
//
// Solidity: function existingPermittedNetworks() view returns(uint16)
func (_PermittedNetworkManagement *PermittedNetworkManagementCaller) ExistingPermittedNetworks(opts *bind.CallOpts) (uint16, error) {
	var out []interface{}
	err := _PermittedNetworkManagement.contract.Call(opts, &out, "existingPermittedNetworks")

	if err != nil {
		return *new(uint16), err
	}

	out0 := *abi.ConvertType(out[0], new(uint16)).(*uint16)

	return out0, err

}

// ExistingPermittedNetworks is a free data retrieval call binding the contract method 0x7419fe64.
//
// Solidity: function existingPermittedNetworks() view returns(uint16)
func (_PermittedNetworkManagement *PermittedNetworkManagementSession) ExistingPermittedNetworks() (uint16, error) {
	return _PermittedNetworkManagement.Contract.ExistingPermittedNetworks(&_PermittedNetworkManagement.CallOpts)
}

// ExistingPermittedNetworks is a free data retrieval call binding the contract method 0x7419fe64.
//
// Solidity: function existingPermittedNetworks() view returns(uint16)
func (_PermittedNetworkManagement *PermittedNetworkManagementCallerSession) ExistingPermittedNetworks() (uint16, error) {
	return _PermittedNetworkManagement.Contract.ExistingPermittedNetworks(&_PermittedNetworkManagement.CallOpts)
}

// GeneratePermittedNetworkId is a free data retrieval call binding the contract method 0x8361868a.
//
// Solidity: function generatePermittedNetworkId() view returns(uint256)
func (_PermittedNetworkManagement *PermittedNetworkManagementCaller) GeneratePermittedNetworkId(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _PermittedNetworkManagement.contract.Call(opts, &out, "generatePermittedNetworkId")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// GeneratePermittedNetworkId is a free data retrieval call binding the contract method 0x8361868a.
//
// Solidity: function generatePermittedNetworkId() view returns(uint256)
func (_PermittedNetworkManagement *PermittedNetworkManagementSession) GeneratePermittedNetworkId() (*big.Int, error) {
	return _PermittedNetworkManagement.Contract.GeneratePermittedNetworkId(&_PermittedNetworkManagement.CallOpts)
}

// GeneratePermittedNetworkId is a free data retrieval call binding the contract method 0x8361868a.
//
// Solidity: function generatePermittedNetworkId() view returns(uint256)
func (_PermittedNetworkManagement *PermittedNetworkManagementCallerSession) GeneratePermittedNetworkId() (*big.Int, error) {
	return _PermittedNetworkManagement.Contract.GeneratePermittedNetworkId(&_PermittedNetworkManagement.CallOpts)
}

// GetAllPermittedNetworks is a free data retrieval call binding the contract method 0xd4d75201.
//
// Solidity: function getAllPermittedNetworks() view returns((string,string,string,string,uint256)[])
func (_PermittedNetworkManagement *PermittedNetworkManagementCaller) GetAllPermittedNetworks(opts *bind.CallOpts) ([]PermittedNetworkManagementContractPermittedNetwork, error) {
	var out []interface{}
	err := _PermittedNetworkManagement.contract.Call(opts, &out, "getAllPermittedNetworks")

	if err != nil {
		return *new([]PermittedNetworkManagementContractPermittedNetwork), err
	}

	out0 := *abi.ConvertType(out[0], new([]PermittedNetworkManagementContractPermittedNetwork)).(*[]PermittedNetworkManagementContractPermittedNetwork)

	return out0, err

}

// GetAllPermittedNetworks is a free data retrieval call binding the contract method 0xd4d75201.
//
// Solidity: function getAllPermittedNetworks() view returns((string,string,string,string,uint256)[])
func (_PermittedNetworkManagement *PermittedNetworkManagementSession) GetAllPermittedNetworks() ([]PermittedNetworkManagementContractPermittedNetwork, error) {
	return _PermittedNetworkManagement.Contract.GetAllPermittedNetworks(&_PermittedNetworkManagement.CallOpts)
}

// GetAllPermittedNetworks is a free data retrieval call binding the contract method 0xd4d75201.
//
// Solidity: function getAllPermittedNetworks() view returns((string,string,string,string,uint256)[])
func (_PermittedNetworkManagement *PermittedNetworkManagementCallerSession) GetAllPermittedNetworks() ([]PermittedNetworkManagementContractPermittedNetwork, error) {
	return _PermittedNetworkManagement.Contract.GetAllPermittedNetworks(&_PermittedNetworkManagement.CallOpts)
}

// GetPermittedNetwork is a free data retrieval call binding the contract method 0xd320f926.
//
// Solidity: function getPermittedNetwork(uint256 permittedNetworkId) view returns((string,string,string,string,uint256))
func (_PermittedNetworkManagement *PermittedNetworkManagementCaller) GetPermittedNetwork(opts *bind.CallOpts, permittedNetworkId *big.Int) (PermittedNetworkManagementContractPermittedNetwork, error) {
	var out []interface{}
	err := _PermittedNetworkManagement.contract.Call(opts, &out, "getPermittedNetwork", permittedNetworkId)

	if err != nil {
		return *new(PermittedNetworkManagementContractPermittedNetwork), err
	}

	out0 := *abi.ConvertType(out[0], new(PermittedNetworkManagementContractPermittedNetwork)).(*PermittedNetworkManagementContractPermittedNetwork)

	return out0, err

}

// GetPermittedNetwork is a free data retrieval call binding the contract method 0xd320f926.
//
// Solidity: function getPermittedNetwork(uint256 permittedNetworkId) view returns((string,string,string,string,uint256))
func (_PermittedNetworkManagement *PermittedNetworkManagementSession) GetPermittedNetwork(permittedNetworkId *big.Int) (PermittedNetworkManagementContractPermittedNetwork, error) {
	return _PermittedNetworkManagement.Contract.GetPermittedNetwork(&_PermittedNetworkManagement.CallOpts, permittedNetworkId)
}

// GetPermittedNetwork is a free data retrieval call binding the contract method 0xd320f926.
//
// Solidity: function getPermittedNetwork(uint256 permittedNetworkId) view returns((string,string,string,string,uint256))
func (_PermittedNetworkManagement *PermittedNetworkManagementCallerSession) GetPermittedNetwork(permittedNetworkId *big.Int) (PermittedNetworkManagementContractPermittedNetwork, error) {
	return _PermittedNetworkManagement.Contract.GetPermittedNetwork(&_PermittedNetworkManagement.CallOpts, permittedNetworkId)
}

// GetPermittedNetworkIndexArray is a free data retrieval call binding the contract method 0xb1f84722.
//
// Solidity: function getPermittedNetworkIndexArray() view returns(uint256[])
func (_PermittedNetworkManagement *PermittedNetworkManagementCaller) GetPermittedNetworkIndexArray(opts *bind.CallOpts) ([]*big.Int, error) {
	var out []interface{}
	err := _PermittedNetworkManagement.contract.Call(opts, &out, "getPermittedNetworkIndexArray")

	if err != nil {
		return *new([]*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new([]*big.Int)).(*[]*big.Int)

	return out0, err

}

// GetPermittedNetworkIndexArray is a free data retrieval call binding the contract method 0xb1f84722.
//
// Solidity: function getPermittedNetworkIndexArray() view returns(uint256[])
func (_PermittedNetworkManagement *PermittedNetworkManagementSession) GetPermittedNetworkIndexArray() ([]*big.Int, error) {
	return _PermittedNetworkManagement.Contract.GetPermittedNetworkIndexArray(&_PermittedNetworkManagement.CallOpts)
}

// GetPermittedNetworkIndexArray is a free data retrieval call binding the contract method 0xb1f84722.
//
// Solidity: function getPermittedNetworkIndexArray() view returns(uint256[])
func (_PermittedNetworkManagement *PermittedNetworkManagementCallerSession) GetPermittedNetworkIndexArray() ([]*big.Int, error) {
	return _PermittedNetworkManagement.Contract.GetPermittedNetworkIndexArray(&_PermittedNetworkManagement.CallOpts)
}

// GetPermittedNetworksNumber is a free data retrieval call binding the contract method 0xe349bd21.
//
// Solidity: function getPermittedNetworksNumber() view returns(uint256)
func (_PermittedNetworkManagement *PermittedNetworkManagementCaller) GetPermittedNetworksNumber(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _PermittedNetworkManagement.contract.Call(opts, &out, "getPermittedNetworksNumber")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// GetPermittedNetworksNumber is a free data retrieval call binding the contract method 0xe349bd21.
//
// Solidity: function getPermittedNetworksNumber() view returns(uint256)
func (_PermittedNetworkManagement *PermittedNetworkManagementSession) GetPermittedNetworksNumber() (*big.Int, error) {
	return _PermittedNetworkManagement.Contract.GetPermittedNetworksNumber(&_PermittedNetworkManagement.CallOpts)
}

// GetPermittedNetworksNumber is a free data retrieval call binding the contract method 0xe349bd21.
//
// Solidity: function getPermittedNetworksNumber() view returns(uint256)
func (_PermittedNetworkManagement *PermittedNetworkManagementCallerSession) GetPermittedNetworksNumber() (*big.Int, error) {
	return _PermittedNetworkManagement.Contract.GetPermittedNetworksNumber(&_PermittedNetworkManagement.CallOpts)
}

// PermittedNetworkExists is a free data retrieval call binding the contract method 0x8eb60baa.
//
// Solidity: function permittedNetworkExists(uint256 permittedNetworkId) view returns(bool)
func (_PermittedNetworkManagement *PermittedNetworkManagementCaller) PermittedNetworkExists(opts *bind.CallOpts, permittedNetworkId *big.Int) (bool, error) {
	var out []interface{}
	err := _PermittedNetworkManagement.contract.Call(opts, &out, "permittedNetworkExists", permittedNetworkId)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// PermittedNetworkExists is a free data retrieval call binding the contract method 0x8eb60baa.
//
// Solidity: function permittedNetworkExists(uint256 permittedNetworkId) view returns(bool)
func (_PermittedNetworkManagement *PermittedNetworkManagementSession) PermittedNetworkExists(permittedNetworkId *big.Int) (bool, error) {
	return _PermittedNetworkManagement.Contract.PermittedNetworkExists(&_PermittedNetworkManagement.CallOpts, permittedNetworkId)
}

// PermittedNetworkExists is a free data retrieval call binding the contract method 0x8eb60baa.
//
// Solidity: function permittedNetworkExists(uint256 permittedNetworkId) view returns(bool)
func (_PermittedNetworkManagement *PermittedNetworkManagementCallerSession) PermittedNetworkExists(permittedNetworkId *big.Int) (bool, error) {
	return _PermittedNetworkManagement.Contract.PermittedNetworkExists(&_PermittedNetworkManagement.CallOpts, permittedNetworkId)
}

// PermittedNetworkIndexs is a free data retrieval call binding the contract method 0x83685c98.
//
// Solidity: function permittedNetworkIndexs(uint256 ) view returns(uint256)
func (_PermittedNetworkManagement *PermittedNetworkManagementCaller) PermittedNetworkIndexs(opts *bind.CallOpts, arg0 *big.Int) (*big.Int, error) {
	var out []interface{}
	err := _PermittedNetworkManagement.contract.Call(opts, &out, "permittedNetworkIndexs", arg0)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// PermittedNetworkIndexs is a free data retrieval call binding the contract method 0x83685c98.
//
// Solidity: function permittedNetworkIndexs(uint256 ) view returns(uint256)
func (_PermittedNetworkManagement *PermittedNetworkManagementSession) PermittedNetworkIndexs(arg0 *big.Int) (*big.Int, error) {
	return _PermittedNetworkManagement.Contract.PermittedNetworkIndexs(&_PermittedNetworkManagement.CallOpts, arg0)
}

// PermittedNetworkIndexs is a free data retrieval call binding the contract method 0x83685c98.
//
// Solidity: function permittedNetworkIndexs(uint256 ) view returns(uint256)
func (_PermittedNetworkManagement *PermittedNetworkManagementCallerSession) PermittedNetworkIndexs(arg0 *big.Int) (*big.Int, error) {
	return _PermittedNetworkManagement.Contract.PermittedNetworkIndexs(&_PermittedNetworkManagement.CallOpts, arg0)
}

// PermittedNetworks is a free data retrieval call binding the contract method 0x2ef1c75d.
//
// Solidity: function permittedNetworks(uint256 ) view returns(string networkName, string networkIP, string networkAddress, string companyName, uint256 permittedNetworkId)
func (_PermittedNetworkManagement *PermittedNetworkManagementCaller) PermittedNetworks(opts *bind.CallOpts, arg0 *big.Int) (struct {
	NetworkName        string
	NetworkIP          string
	NetworkAddress     string
	CompanyName        string
	PermittedNetworkId *big.Int
}, error) {
	var out []interface{}
	err := _PermittedNetworkManagement.contract.Call(opts, &out, "permittedNetworks", arg0)

	outstruct := new(struct {
		NetworkName        string
		NetworkIP          string
		NetworkAddress     string
		CompanyName        string
		PermittedNetworkId *big.Int
	})
	if err != nil {
		return *outstruct, err
	}

	outstruct.NetworkName = *abi.ConvertType(out[0], new(string)).(*string)
	outstruct.NetworkIP = *abi.ConvertType(out[1], new(string)).(*string)
	outstruct.NetworkAddress = *abi.ConvertType(out[2], new(string)).(*string)
	outstruct.CompanyName = *abi.ConvertType(out[3], new(string)).(*string)
	outstruct.PermittedNetworkId = *abi.ConvertType(out[4], new(*big.Int)).(**big.Int)

	return *outstruct, err

}

// PermittedNetworks is a free data retrieval call binding the contract method 0x2ef1c75d.
//
// Solidity: function permittedNetworks(uint256 ) view returns(string networkName, string networkIP, string networkAddress, string companyName, uint256 permittedNetworkId)
func (_PermittedNetworkManagement *PermittedNetworkManagementSession) PermittedNetworks(arg0 *big.Int) (struct {
	NetworkName        string
	NetworkIP          string
	NetworkAddress     string
	CompanyName        string
	PermittedNetworkId *big.Int
}, error) {
	return _PermittedNetworkManagement.Contract.PermittedNetworks(&_PermittedNetworkManagement.CallOpts, arg0)
}

// PermittedNetworks is a free data retrieval call binding the contract method 0x2ef1c75d.
//
// Solidity: function permittedNetworks(uint256 ) view returns(string networkName, string networkIP, string networkAddress, string companyName, uint256 permittedNetworkId)
func (_PermittedNetworkManagement *PermittedNetworkManagementCallerSession) PermittedNetworks(arg0 *big.Int) (struct {
	NetworkName        string
	NetworkIP          string
	NetworkAddress     string
	CompanyName        string
	PermittedNetworkId *big.Int
}, error) {
	return _PermittedNetworkManagement.Contract.PermittedNetworks(&_PermittedNetworkManagement.CallOpts, arg0)
}

// CreatePermittedNetwork is a paid mutator transaction binding the contract method 0x3f7408a5.
//
// Solidity: function createPermittedNetwork(string networkName, string networkIP, string networkAddress, string companyName) returns((string,string,string,string,uint256))
func (_PermittedNetworkManagement *PermittedNetworkManagementTransactor) CreatePermittedNetwork(opts *bind.TransactOpts, networkName string, networkIP string, networkAddress string, companyName string) (*types.Transaction, error) {
	return _PermittedNetworkManagement.contract.Transact(opts, "createPermittedNetwork", networkName, networkIP, networkAddress, companyName)
}

// CreatePermittedNetwork is a paid mutator transaction binding the contract method 0x3f7408a5.
//
// Solidity: function createPermittedNetwork(string networkName, string networkIP, string networkAddress, string companyName) returns((string,string,string,string,uint256))
func (_PermittedNetworkManagement *PermittedNetworkManagementSession) CreatePermittedNetwork(networkName string, networkIP string, networkAddress string, companyName string) (*types.Transaction, error) {
	return _PermittedNetworkManagement.Contract.CreatePermittedNetwork(&_PermittedNetworkManagement.TransactOpts, networkName, networkIP, networkAddress, companyName)
}

// CreatePermittedNetwork is a paid mutator transaction binding the contract method 0x3f7408a5.
//
// Solidity: function createPermittedNetwork(string networkName, string networkIP, string networkAddress, string companyName) returns((string,string,string,string,uint256))
func (_PermittedNetworkManagement *PermittedNetworkManagementTransactorSession) CreatePermittedNetwork(networkName string, networkIP string, networkAddress string, companyName string) (*types.Transaction, error) {
	return _PermittedNetworkManagement.Contract.CreatePermittedNetwork(&_PermittedNetworkManagement.TransactOpts, networkName, networkIP, networkAddress, companyName)
}

// DeletePermittedNetwork is a paid mutator transaction binding the contract method 0xfa63b2b7.
//
// Solidity: function deletePermittedNetwork(uint256 permittedNetworkId) returns()
func (_PermittedNetworkManagement *PermittedNetworkManagementTransactor) DeletePermittedNetwork(opts *bind.TransactOpts, permittedNetworkId *big.Int) (*types.Transaction, error) {
	return _PermittedNetworkManagement.contract.Transact(opts, "deletePermittedNetwork", permittedNetworkId)
}

// DeletePermittedNetwork is a paid mutator transaction binding the contract method 0xfa63b2b7.
//
// Solidity: function deletePermittedNetwork(uint256 permittedNetworkId) returns()
func (_PermittedNetworkManagement *PermittedNetworkManagementSession) DeletePermittedNetwork(permittedNetworkId *big.Int) (*types.Transaction, error) {
	return _PermittedNetworkManagement.Contract.DeletePermittedNetwork(&_PermittedNetworkManagement.TransactOpts, permittedNetworkId)
}

// DeletePermittedNetwork is a paid mutator transaction binding the contract method 0xfa63b2b7.
//
// Solidity: function deletePermittedNetwork(uint256 permittedNetworkId) returns()
func (_PermittedNetworkManagement *PermittedNetworkManagementTransactorSession) DeletePermittedNetwork(permittedNetworkId *big.Int) (*types.Transaction, error) {
	return _PermittedNetworkManagement.Contract.DeletePermittedNetwork(&_PermittedNetworkManagement.TransactOpts, permittedNetworkId)
}

// UpdatePermittedNetwork is a paid mutator transaction binding the contract method 0xcdd6b7a1.
//
// Solidity: function updatePermittedNetwork(uint256 permittedNetworkId, string networkName, string networkIP, string networkAddress, string companyName) returns((string,string,string,string,uint256))
func (_PermittedNetworkManagement *PermittedNetworkManagementTransactor) UpdatePermittedNetwork(opts *bind.TransactOpts, permittedNetworkId *big.Int, networkName string, networkIP string, networkAddress string, companyName string) (*types.Transaction, error) {
	return _PermittedNetworkManagement.contract.Transact(opts, "updatePermittedNetwork", permittedNetworkId, networkName, networkIP, networkAddress, companyName)
}

// UpdatePermittedNetwork is a paid mutator transaction binding the contract method 0xcdd6b7a1.
//
// Solidity: function updatePermittedNetwork(uint256 permittedNetworkId, string networkName, string networkIP, string networkAddress, string companyName) returns((string,string,string,string,uint256))
func (_PermittedNetworkManagement *PermittedNetworkManagementSession) UpdatePermittedNetwork(permittedNetworkId *big.Int, networkName string, networkIP string, networkAddress string, companyName string) (*types.Transaction, error) {
	return _PermittedNetworkManagement.Contract.UpdatePermittedNetwork(&_PermittedNetworkManagement.TransactOpts, permittedNetworkId, networkName, networkIP, networkAddress, companyName)
}

// UpdatePermittedNetwork is a paid mutator transaction binding the contract method 0xcdd6b7a1.
//
// Solidity: function updatePermittedNetwork(uint256 permittedNetworkId, string networkName, string networkIP, string networkAddress, string companyName) returns((string,string,string,string,uint256))
func (_PermittedNetworkManagement *PermittedNetworkManagementTransactorSession) UpdatePermittedNetwork(permittedNetworkId *big.Int, networkName string, networkIP string, networkAddress string, companyName string) (*types.Transaction, error) {
	return _PermittedNetworkManagement.Contract.UpdatePermittedNetwork(&_PermittedNetworkManagement.TransactOpts, permittedNetworkId, networkName, networkIP, networkAddress, companyName)
}
