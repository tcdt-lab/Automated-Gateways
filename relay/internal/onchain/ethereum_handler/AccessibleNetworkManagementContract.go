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

// AccessibleNetworkManagementContarctAccessibleNetwork is an auto generated low-level Go binding around an user-defined struct.
type AccessibleNetworkManagementContarctAccessibleNetwork struct {
	NetworkName         string
	NetworkIP           string
	NetworkAddress      string
	CompanyName         string
	AccessibleNetworkId *big.Int
}

// AccessibleNetworkManagementMetaData contains all meta data concerning the AccessibleNetworkManagement contract.
var AccessibleNetworkManagementMetaData = &bind.MetaData{
	ABI: "[{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"name\":\"accessibleNetworkIndexs\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"name\":\"accessibleNetworks\",\"outputs\":[{\"internalType\":\"string\",\"name\":\"networkName\",\"type\":\"string\"},{\"internalType\":\"string\",\"name\":\"networkIP\",\"type\":\"string\"},{\"internalType\":\"string\",\"name\":\"networkAddress\",\"type\":\"string\"},{\"internalType\":\"string\",\"name\":\"companyName\",\"type\":\"string\"},{\"internalType\":\"uint256\",\"name\":\"accessibleNetworkId\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"string\",\"name\":\"networkName\",\"type\":\"string\"},{\"internalType\":\"string\",\"name\":\"networkIP\",\"type\":\"string\"},{\"internalType\":\"string\",\"name\":\"networkAddress\",\"type\":\"string\"},{\"internalType\":\"string\",\"name\":\"companyName\",\"type\":\"string\"}],\"name\":\"createAccessibleNetwork\",\"outputs\":[{\"components\":[{\"internalType\":\"string\",\"name\":\"networkName\",\"type\":\"string\"},{\"internalType\":\"string\",\"name\":\"networkIP\",\"type\":\"string\"},{\"internalType\":\"string\",\"name\":\"networkAddress\",\"type\":\"string\"},{\"internalType\":\"string\",\"name\":\"companyName\",\"type\":\"string\"},{\"internalType\":\"uint256\",\"name\":\"accessibleNetworkId\",\"type\":\"uint256\"}],\"internalType\":\"structAccessibleNetworkManagementContarct.AccessibleNetwork\",\"name\":\"\",\"type\":\"tuple\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"accessibleNetworkId\",\"type\":\"uint256\"}],\"name\":\"deleteAccessibleNetwork\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"existingAccessibleNetworks\",\"outputs\":[{\"internalType\":\"uint16\",\"name\":\"\",\"type\":\"uint16\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"generateAccessibleNetworkId\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"accessibleNetworkId\",\"type\":\"uint256\"}],\"name\":\"getAccessibleNetwork\",\"outputs\":[{\"components\":[{\"internalType\":\"string\",\"name\":\"networkName\",\"type\":\"string\"},{\"internalType\":\"string\",\"name\":\"networkIP\",\"type\":\"string\"},{\"internalType\":\"string\",\"name\":\"networkAddress\",\"type\":\"string\"},{\"internalType\":\"string\",\"name\":\"companyName\",\"type\":\"string\"},{\"internalType\":\"uint256\",\"name\":\"accessibleNetworkId\",\"type\":\"uint256\"}],\"internalType\":\"structAccessibleNetworkManagementContarct.AccessibleNetwork\",\"name\":\"\",\"type\":\"tuple\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"getAccessibleNetworksNumber\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"getAllAccessibleNetworks\",\"outputs\":[{\"components\":[{\"internalType\":\"string\",\"name\":\"networkName\",\"type\":\"string\"},{\"internalType\":\"string\",\"name\":\"networkIP\",\"type\":\"string\"},{\"internalType\":\"string\",\"name\":\"networkAddress\",\"type\":\"string\"},{\"internalType\":\"string\",\"name\":\"companyName\",\"type\":\"string\"},{\"internalType\":\"uint256\",\"name\":\"accessibleNetworkId\",\"type\":\"uint256\"}],\"internalType\":\"structAccessibleNetworkManagementContarct.AccessibleNetwork[]\",\"name\":\"\",\"type\":\"tuple[]\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"accessibleNetworkId\",\"type\":\"uint256\"}],\"name\":\"isAccessibleNetworkExist\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"accessibleNetworkId\",\"type\":\"uint256\"},{\"internalType\":\"string\",\"name\":\"networkName\",\"type\":\"string\"},{\"internalType\":\"string\",\"name\":\"networkIP\",\"type\":\"string\"},{\"internalType\":\"string\",\"name\":\"networkAddress\",\"type\":\"string\"},{\"internalType\":\"string\",\"name\":\"companyName\",\"type\":\"string\"}],\"name\":\"updateAccessibleNetwork\",\"outputs\":[{\"components\":[{\"internalType\":\"string\",\"name\":\"networkName\",\"type\":\"string\"},{\"internalType\":\"string\",\"name\":\"networkIP\",\"type\":\"string\"},{\"internalType\":\"string\",\"name\":\"networkAddress\",\"type\":\"string\"},{\"internalType\":\"string\",\"name\":\"companyName\",\"type\":\"string\"},{\"internalType\":\"uint256\",\"name\":\"accessibleNetworkId\",\"type\":\"uint256\"}],\"internalType\":\"structAccessibleNetworkManagementContarct.AccessibleNetwork\",\"name\":\"\",\"type\":\"tuple\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"}]",
}

// AccessibleNetworkManagementABI is the input ABI used to generate the binding from.
// Deprecated: Use AccessibleNetworkManagementMetaData.ABI instead.
var AccessibleNetworkManagementABI = AccessibleNetworkManagementMetaData.ABI

// AccessibleNetworkManagement is an auto generated Go binding around an Ethereum contract.
type AccessibleNetworkManagement struct {
	AccessibleNetworkManagementCaller     // Read-only binding to the contract
	AccessibleNetworkManagementTransactor // Write-only binding to the contract
	AccessibleNetworkManagementFilterer   // Log filterer for contract events
}

// AccessibleNetworkManagementCaller is an auto generated read-only Go binding around an Ethereum contract.
type AccessibleNetworkManagementCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// AccessibleNetworkManagementTransactor is an auto generated write-only Go binding around an Ethereum contract.
type AccessibleNetworkManagementTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// AccessibleNetworkManagementFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type AccessibleNetworkManagementFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// AccessibleNetworkManagementSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type AccessibleNetworkManagementSession struct {
	Contract     *AccessibleNetworkManagement // Generic contract binding to set the session for
	CallOpts     bind.CallOpts                // Call options to use throughout this session
	TransactOpts bind.TransactOpts            // Transaction auth options to use throughout this session
}

// AccessibleNetworkManagementCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type AccessibleNetworkManagementCallerSession struct {
	Contract *AccessibleNetworkManagementCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts                      // Call options to use throughout this session
}

// AccessibleNetworkManagementTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type AccessibleNetworkManagementTransactorSession struct {
	Contract     *AccessibleNetworkManagementTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts                      // Transaction auth options to use throughout this session
}

// AccessibleNetworkManagementRaw is an auto generated low-level Go binding around an Ethereum contract.
type AccessibleNetworkManagementRaw struct {
	Contract *AccessibleNetworkManagement // Generic contract binding to access the raw methods on
}

// AccessibleNetworkManagementCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type AccessibleNetworkManagementCallerRaw struct {
	Contract *AccessibleNetworkManagementCaller // Generic read-only contract binding to access the raw methods on
}

// AccessibleNetworkManagementTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type AccessibleNetworkManagementTransactorRaw struct {
	Contract *AccessibleNetworkManagementTransactor // Generic write-only contract binding to access the raw methods on
}

// NewAccessibleNetworkManagement creates a new instance of AccessibleNetworkManagement, bound to a specific deployed contract.
func NewAccessibleNetworkManagement(address common.Address, backend bind.ContractBackend) (*AccessibleNetworkManagement, error) {
	contract, err := bindAccessibleNetworkManagement(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &AccessibleNetworkManagement{AccessibleNetworkManagementCaller: AccessibleNetworkManagementCaller{contract: contract}, AccessibleNetworkManagementTransactor: AccessibleNetworkManagementTransactor{contract: contract}, AccessibleNetworkManagementFilterer: AccessibleNetworkManagementFilterer{contract: contract}}, nil
}

// NewAccessibleNetworkManagementCaller creates a new read-only instance of AccessibleNetworkManagement, bound to a specific deployed contract.
func NewAccessibleNetworkManagementCaller(address common.Address, caller bind.ContractCaller) (*AccessibleNetworkManagementCaller, error) {
	contract, err := bindAccessibleNetworkManagement(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &AccessibleNetworkManagementCaller{contract: contract}, nil
}

// NewAccessibleNetworkManagementTransactor creates a new write-only instance of AccessibleNetworkManagement, bound to a specific deployed contract.
func NewAccessibleNetworkManagementTransactor(address common.Address, transactor bind.ContractTransactor) (*AccessibleNetworkManagementTransactor, error) {
	contract, err := bindAccessibleNetworkManagement(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &AccessibleNetworkManagementTransactor{contract: contract}, nil
}

// NewAccessibleNetworkManagementFilterer creates a new log filterer instance of AccessibleNetworkManagement, bound to a specific deployed contract.
func NewAccessibleNetworkManagementFilterer(address common.Address, filterer bind.ContractFilterer) (*AccessibleNetworkManagementFilterer, error) {
	contract, err := bindAccessibleNetworkManagement(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &AccessibleNetworkManagementFilterer{contract: contract}, nil
}

// bindAccessibleNetworkManagement binds a generic wrapper to an already deployed contract.
func bindAccessibleNetworkManagement(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := AccessibleNetworkManagementMetaData.GetAbi()
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, *parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_AccessibleNetworkManagement *AccessibleNetworkManagementRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _AccessibleNetworkManagement.Contract.AccessibleNetworkManagementCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_AccessibleNetworkManagement *AccessibleNetworkManagementRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _AccessibleNetworkManagement.Contract.AccessibleNetworkManagementTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_AccessibleNetworkManagement *AccessibleNetworkManagementRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _AccessibleNetworkManagement.Contract.AccessibleNetworkManagementTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_AccessibleNetworkManagement *AccessibleNetworkManagementCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _AccessibleNetworkManagement.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_AccessibleNetworkManagement *AccessibleNetworkManagementTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _AccessibleNetworkManagement.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_AccessibleNetworkManagement *AccessibleNetworkManagementTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _AccessibleNetworkManagement.Contract.contract.Transact(opts, method, params...)
}

// AccessibleNetworkIndexs is a free data retrieval call binding the contract method 0x903c1263.
//
// Solidity: function accessibleNetworkIndexs(uint256 ) view returns(uint256)
func (_AccessibleNetworkManagement *AccessibleNetworkManagementCaller) AccessibleNetworkIndexs(opts *bind.CallOpts, arg0 *big.Int) (*big.Int, error) {
	var out []interface{}
	err := _AccessibleNetworkManagement.contract.Call(opts, &out, "accessibleNetworkIndexs", arg0)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// AccessibleNetworkIndexs is a free data retrieval call binding the contract method 0x903c1263.
//
// Solidity: function accessibleNetworkIndexs(uint256 ) view returns(uint256)
func (_AccessibleNetworkManagement *AccessibleNetworkManagementSession) AccessibleNetworkIndexs(arg0 *big.Int) (*big.Int, error) {
	return _AccessibleNetworkManagement.Contract.AccessibleNetworkIndexs(&_AccessibleNetworkManagement.CallOpts, arg0)
}

// AccessibleNetworkIndexs is a free data retrieval call binding the contract method 0x903c1263.
//
// Solidity: function accessibleNetworkIndexs(uint256 ) view returns(uint256)
func (_AccessibleNetworkManagement *AccessibleNetworkManagementCallerSession) AccessibleNetworkIndexs(arg0 *big.Int) (*big.Int, error) {
	return _AccessibleNetworkManagement.Contract.AccessibleNetworkIndexs(&_AccessibleNetworkManagement.CallOpts, arg0)
}

// AccessibleNetworks is a free data retrieval call binding the contract method 0x691b3614.
//
// Solidity: function accessibleNetworks(uint256 ) view returns(string networkName, string networkIP, string networkAddress, string companyName, uint256 accessibleNetworkId)
func (_AccessibleNetworkManagement *AccessibleNetworkManagementCaller) AccessibleNetworks(opts *bind.CallOpts, arg0 *big.Int) (struct {
	NetworkName         string
	NetworkIP           string
	NetworkAddress      string
	CompanyName         string
	AccessibleNetworkId *big.Int
}, error) {
	var out []interface{}
	err := _AccessibleNetworkManagement.contract.Call(opts, &out, "accessibleNetworks", arg0)

	outstruct := new(struct {
		NetworkName         string
		NetworkIP           string
		NetworkAddress      string
		CompanyName         string
		AccessibleNetworkId *big.Int
	})
	if err != nil {
		return *outstruct, err
	}

	outstruct.NetworkName = *abi.ConvertType(out[0], new(string)).(*string)
	outstruct.NetworkIP = *abi.ConvertType(out[1], new(string)).(*string)
	outstruct.NetworkAddress = *abi.ConvertType(out[2], new(string)).(*string)
	outstruct.CompanyName = *abi.ConvertType(out[3], new(string)).(*string)
	outstruct.AccessibleNetworkId = *abi.ConvertType(out[4], new(*big.Int)).(**big.Int)

	return *outstruct, err

}

// AccessibleNetworks is a free data retrieval call binding the contract method 0x691b3614.
//
// Solidity: function accessibleNetworks(uint256 ) view returns(string networkName, string networkIP, string networkAddress, string companyName, uint256 accessibleNetworkId)
func (_AccessibleNetworkManagement *AccessibleNetworkManagementSession) AccessibleNetworks(arg0 *big.Int) (struct {
	NetworkName         string
	NetworkIP           string
	NetworkAddress      string
	CompanyName         string
	AccessibleNetworkId *big.Int
}, error) {
	return _AccessibleNetworkManagement.Contract.AccessibleNetworks(&_AccessibleNetworkManagement.CallOpts, arg0)
}

// AccessibleNetworks is a free data retrieval call binding the contract method 0x691b3614.
//
// Solidity: function accessibleNetworks(uint256 ) view returns(string networkName, string networkIP, string networkAddress, string companyName, uint256 accessibleNetworkId)
func (_AccessibleNetworkManagement *AccessibleNetworkManagementCallerSession) AccessibleNetworks(arg0 *big.Int) (struct {
	NetworkName         string
	NetworkIP           string
	NetworkAddress      string
	CompanyName         string
	AccessibleNetworkId *big.Int
}, error) {
	return _AccessibleNetworkManagement.Contract.AccessibleNetworks(&_AccessibleNetworkManagement.CallOpts, arg0)
}

// ExistingAccessibleNetworks is a free data retrieval call binding the contract method 0x35fb246b.
//
// Solidity: function existingAccessibleNetworks() view returns(uint16)
func (_AccessibleNetworkManagement *AccessibleNetworkManagementCaller) ExistingAccessibleNetworks(opts *bind.CallOpts) (uint16, error) {
	var out []interface{}
	err := _AccessibleNetworkManagement.contract.Call(opts, &out, "existingAccessibleNetworks")

	if err != nil {
		return *new(uint16), err
	}

	out0 := *abi.ConvertType(out[0], new(uint16)).(*uint16)

	return out0, err

}

// ExistingAccessibleNetworks is a free data retrieval call binding the contract method 0x35fb246b.
//
// Solidity: function existingAccessibleNetworks() view returns(uint16)
func (_AccessibleNetworkManagement *AccessibleNetworkManagementSession) ExistingAccessibleNetworks() (uint16, error) {
	return _AccessibleNetworkManagement.Contract.ExistingAccessibleNetworks(&_AccessibleNetworkManagement.CallOpts)
}

// ExistingAccessibleNetworks is a free data retrieval call binding the contract method 0x35fb246b.
//
// Solidity: function existingAccessibleNetworks() view returns(uint16)
func (_AccessibleNetworkManagement *AccessibleNetworkManagementCallerSession) ExistingAccessibleNetworks() (uint16, error) {
	return _AccessibleNetworkManagement.Contract.ExistingAccessibleNetworks(&_AccessibleNetworkManagement.CallOpts)
}

// GenerateAccessibleNetworkId is a free data retrieval call binding the contract method 0x66be824e.
//
// Solidity: function generateAccessibleNetworkId() view returns(uint256)
func (_AccessibleNetworkManagement *AccessibleNetworkManagementCaller) GenerateAccessibleNetworkId(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _AccessibleNetworkManagement.contract.Call(opts, &out, "generateAccessibleNetworkId")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// GenerateAccessibleNetworkId is a free data retrieval call binding the contract method 0x66be824e.
//
// Solidity: function generateAccessibleNetworkId() view returns(uint256)
func (_AccessibleNetworkManagement *AccessibleNetworkManagementSession) GenerateAccessibleNetworkId() (*big.Int, error) {
	return _AccessibleNetworkManagement.Contract.GenerateAccessibleNetworkId(&_AccessibleNetworkManagement.CallOpts)
}

// GenerateAccessibleNetworkId is a free data retrieval call binding the contract method 0x66be824e.
//
// Solidity: function generateAccessibleNetworkId() view returns(uint256)
func (_AccessibleNetworkManagement *AccessibleNetworkManagementCallerSession) GenerateAccessibleNetworkId() (*big.Int, error) {
	return _AccessibleNetworkManagement.Contract.GenerateAccessibleNetworkId(&_AccessibleNetworkManagement.CallOpts)
}

// GetAccessibleNetwork is a free data retrieval call binding the contract method 0x4bdc1f3f.
//
// Solidity: function getAccessibleNetwork(uint256 accessibleNetworkId) view returns((string,string,string,string,uint256))
func (_AccessibleNetworkManagement *AccessibleNetworkManagementCaller) GetAccessibleNetwork(opts *bind.CallOpts, accessibleNetworkId *big.Int) (AccessibleNetworkManagementContarctAccessibleNetwork, error) {
	var out []interface{}
	err := _AccessibleNetworkManagement.contract.Call(opts, &out, "getAccessibleNetwork", accessibleNetworkId)

	if err != nil {
		return *new(AccessibleNetworkManagementContarctAccessibleNetwork), err
	}

	out0 := *abi.ConvertType(out[0], new(AccessibleNetworkManagementContarctAccessibleNetwork)).(*AccessibleNetworkManagementContarctAccessibleNetwork)

	return out0, err

}

// GetAccessibleNetwork is a free data retrieval call binding the contract method 0x4bdc1f3f.
//
// Solidity: function getAccessibleNetwork(uint256 accessibleNetworkId) view returns((string,string,string,string,uint256))
func (_AccessibleNetworkManagement *AccessibleNetworkManagementSession) GetAccessibleNetwork(accessibleNetworkId *big.Int) (AccessibleNetworkManagementContarctAccessibleNetwork, error) {
	return _AccessibleNetworkManagement.Contract.GetAccessibleNetwork(&_AccessibleNetworkManagement.CallOpts, accessibleNetworkId)
}

// GetAccessibleNetwork is a free data retrieval call binding the contract method 0x4bdc1f3f.
//
// Solidity: function getAccessibleNetwork(uint256 accessibleNetworkId) view returns((string,string,string,string,uint256))
func (_AccessibleNetworkManagement *AccessibleNetworkManagementCallerSession) GetAccessibleNetwork(accessibleNetworkId *big.Int) (AccessibleNetworkManagementContarctAccessibleNetwork, error) {
	return _AccessibleNetworkManagement.Contract.GetAccessibleNetwork(&_AccessibleNetworkManagement.CallOpts, accessibleNetworkId)
}

// GetAccessibleNetworksNumber is a free data retrieval call binding the contract method 0x33cc4c8b.
//
// Solidity: function getAccessibleNetworksNumber() view returns(uint256)
func (_AccessibleNetworkManagement *AccessibleNetworkManagementCaller) GetAccessibleNetworksNumber(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _AccessibleNetworkManagement.contract.Call(opts, &out, "getAccessibleNetworksNumber")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// GetAccessibleNetworksNumber is a free data retrieval call binding the contract method 0x33cc4c8b.
//
// Solidity: function getAccessibleNetworksNumber() view returns(uint256)
func (_AccessibleNetworkManagement *AccessibleNetworkManagementSession) GetAccessibleNetworksNumber() (*big.Int, error) {
	return _AccessibleNetworkManagement.Contract.GetAccessibleNetworksNumber(&_AccessibleNetworkManagement.CallOpts)
}

// GetAccessibleNetworksNumber is a free data retrieval call binding the contract method 0x33cc4c8b.
//
// Solidity: function getAccessibleNetworksNumber() view returns(uint256)
func (_AccessibleNetworkManagement *AccessibleNetworkManagementCallerSession) GetAccessibleNetworksNumber() (*big.Int, error) {
	return _AccessibleNetworkManagement.Contract.GetAccessibleNetworksNumber(&_AccessibleNetworkManagement.CallOpts)
}

// GetAllAccessibleNetworks is a free data retrieval call binding the contract method 0xda233319.
//
// Solidity: function getAllAccessibleNetworks() view returns((string,string,string,string,uint256)[])
func (_AccessibleNetworkManagement *AccessibleNetworkManagementCaller) GetAllAccessibleNetworks(opts *bind.CallOpts) ([]AccessibleNetworkManagementContarctAccessibleNetwork, error) {
	var out []interface{}
	err := _AccessibleNetworkManagement.contract.Call(opts, &out, "getAllAccessibleNetworks")

	if err != nil {
		return *new([]AccessibleNetworkManagementContarctAccessibleNetwork), err
	}

	out0 := *abi.ConvertType(out[0], new([]AccessibleNetworkManagementContarctAccessibleNetwork)).(*[]AccessibleNetworkManagementContarctAccessibleNetwork)

	return out0, err

}

// GetAllAccessibleNetworks is a free data retrieval call binding the contract method 0xda233319.
//
// Solidity: function getAllAccessibleNetworks() view returns((string,string,string,string,uint256)[])
func (_AccessibleNetworkManagement *AccessibleNetworkManagementSession) GetAllAccessibleNetworks() ([]AccessibleNetworkManagementContarctAccessibleNetwork, error) {
	return _AccessibleNetworkManagement.Contract.GetAllAccessibleNetworks(&_AccessibleNetworkManagement.CallOpts)
}

// GetAllAccessibleNetworks is a free data retrieval call binding the contract method 0xda233319.
//
// Solidity: function getAllAccessibleNetworks() view returns((string,string,string,string,uint256)[])
func (_AccessibleNetworkManagement *AccessibleNetworkManagementCallerSession) GetAllAccessibleNetworks() ([]AccessibleNetworkManagementContarctAccessibleNetwork, error) {
	return _AccessibleNetworkManagement.Contract.GetAllAccessibleNetworks(&_AccessibleNetworkManagement.CallOpts)
}

// IsAccessibleNetworkExist is a free data retrieval call binding the contract method 0xaf928456.
//
// Solidity: function isAccessibleNetworkExist(uint256 accessibleNetworkId) view returns(bool)
func (_AccessibleNetworkManagement *AccessibleNetworkManagementCaller) IsAccessibleNetworkExist(opts *bind.CallOpts, accessibleNetworkId *big.Int) (bool, error) {
	var out []interface{}
	err := _AccessibleNetworkManagement.contract.Call(opts, &out, "isAccessibleNetworkExist", accessibleNetworkId)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// IsAccessibleNetworkExist is a free data retrieval call binding the contract method 0xaf928456.
//
// Solidity: function isAccessibleNetworkExist(uint256 accessibleNetworkId) view returns(bool)
func (_AccessibleNetworkManagement *AccessibleNetworkManagementSession) IsAccessibleNetworkExist(accessibleNetworkId *big.Int) (bool, error) {
	return _AccessibleNetworkManagement.Contract.IsAccessibleNetworkExist(&_AccessibleNetworkManagement.CallOpts, accessibleNetworkId)
}

// IsAccessibleNetworkExist is a free data retrieval call binding the contract method 0xaf928456.
//
// Solidity: function isAccessibleNetworkExist(uint256 accessibleNetworkId) view returns(bool)
func (_AccessibleNetworkManagement *AccessibleNetworkManagementCallerSession) IsAccessibleNetworkExist(accessibleNetworkId *big.Int) (bool, error) {
	return _AccessibleNetworkManagement.Contract.IsAccessibleNetworkExist(&_AccessibleNetworkManagement.CallOpts, accessibleNetworkId)
}

// CreateAccessibleNetwork is a paid mutator transaction binding the contract method 0x0675b00b.
//
// Solidity: function createAccessibleNetwork(string networkName, string networkIP, string networkAddress, string companyName) returns((string,string,string,string,uint256))
func (_AccessibleNetworkManagement *AccessibleNetworkManagementTransactor) CreateAccessibleNetwork(opts *bind.TransactOpts, networkName string, networkIP string, networkAddress string, companyName string) (*types.Transaction, error) {
	return _AccessibleNetworkManagement.contract.Transact(opts, "createAccessibleNetwork", networkName, networkIP, networkAddress, companyName)
}

// CreateAccessibleNetwork is a paid mutator transaction binding the contract method 0x0675b00b.
//
// Solidity: function createAccessibleNetwork(string networkName, string networkIP, string networkAddress, string companyName) returns((string,string,string,string,uint256))
func (_AccessibleNetworkManagement *AccessibleNetworkManagementSession) CreateAccessibleNetwork(networkName string, networkIP string, networkAddress string, companyName string) (*types.Transaction, error) {
	return _AccessibleNetworkManagement.Contract.CreateAccessibleNetwork(&_AccessibleNetworkManagement.TransactOpts, networkName, networkIP, networkAddress, companyName)
}

// CreateAccessibleNetwork is a paid mutator transaction binding the contract method 0x0675b00b.
//
// Solidity: function createAccessibleNetwork(string networkName, string networkIP, string networkAddress, string companyName) returns((string,string,string,string,uint256))
func (_AccessibleNetworkManagement *AccessibleNetworkManagementTransactorSession) CreateAccessibleNetwork(networkName string, networkIP string, networkAddress string, companyName string) (*types.Transaction, error) {
	return _AccessibleNetworkManagement.Contract.CreateAccessibleNetwork(&_AccessibleNetworkManagement.TransactOpts, networkName, networkIP, networkAddress, companyName)
}

// DeleteAccessibleNetwork is a paid mutator transaction binding the contract method 0x46e4e3d0.
//
// Solidity: function deleteAccessibleNetwork(uint256 accessibleNetworkId) returns()
func (_AccessibleNetworkManagement *AccessibleNetworkManagementTransactor) DeleteAccessibleNetwork(opts *bind.TransactOpts, accessibleNetworkId *big.Int) (*types.Transaction, error) {
	return _AccessibleNetworkManagement.contract.Transact(opts, "deleteAccessibleNetwork", accessibleNetworkId)
}

// DeleteAccessibleNetwork is a paid mutator transaction binding the contract method 0x46e4e3d0.
//
// Solidity: function deleteAccessibleNetwork(uint256 accessibleNetworkId) returns()
func (_AccessibleNetworkManagement *AccessibleNetworkManagementSession) DeleteAccessibleNetwork(accessibleNetworkId *big.Int) (*types.Transaction, error) {
	return _AccessibleNetworkManagement.Contract.DeleteAccessibleNetwork(&_AccessibleNetworkManagement.TransactOpts, accessibleNetworkId)
}

// DeleteAccessibleNetwork is a paid mutator transaction binding the contract method 0x46e4e3d0.
//
// Solidity: function deleteAccessibleNetwork(uint256 accessibleNetworkId) returns()
func (_AccessibleNetworkManagement *AccessibleNetworkManagementTransactorSession) DeleteAccessibleNetwork(accessibleNetworkId *big.Int) (*types.Transaction, error) {
	return _AccessibleNetworkManagement.Contract.DeleteAccessibleNetwork(&_AccessibleNetworkManagement.TransactOpts, accessibleNetworkId)
}

// UpdateAccessibleNetwork is a paid mutator transaction binding the contract method 0xae59854e.
//
// Solidity: function updateAccessibleNetwork(uint256 accessibleNetworkId, string networkName, string networkIP, string networkAddress, string companyName) returns((string,string,string,string,uint256))
func (_AccessibleNetworkManagement *AccessibleNetworkManagementTransactor) UpdateAccessibleNetwork(opts *bind.TransactOpts, accessibleNetworkId *big.Int, networkName string, networkIP string, networkAddress string, companyName string) (*types.Transaction, error) {
	return _AccessibleNetworkManagement.contract.Transact(opts, "updateAccessibleNetwork", accessibleNetworkId, networkName, networkIP, networkAddress, companyName)
}

// UpdateAccessibleNetwork is a paid mutator transaction binding the contract method 0xae59854e.
//
// Solidity: function updateAccessibleNetwork(uint256 accessibleNetworkId, string networkName, string networkIP, string networkAddress, string companyName) returns((string,string,string,string,uint256))
func (_AccessibleNetworkManagement *AccessibleNetworkManagementSession) UpdateAccessibleNetwork(accessibleNetworkId *big.Int, networkName string, networkIP string, networkAddress string, companyName string) (*types.Transaction, error) {
	return _AccessibleNetworkManagement.Contract.UpdateAccessibleNetwork(&_AccessibleNetworkManagement.TransactOpts, accessibleNetworkId, networkName, networkIP, networkAddress, companyName)
}

// UpdateAccessibleNetwork is a paid mutator transaction binding the contract method 0xae59854e.
//
// Solidity: function updateAccessibleNetwork(uint256 accessibleNetworkId, string networkName, string networkIP, string networkAddress, string companyName) returns((string,string,string,string,uint256))
func (_AccessibleNetworkManagement *AccessibleNetworkManagementTransactorSession) UpdateAccessibleNetwork(accessibleNetworkId *big.Int, networkName string, networkIP string, networkAddress string, companyName string) (*types.Transaction, error) {
	return _AccessibleNetworkManagement.Contract.UpdateAccessibleNetwork(&_AccessibleNetworkManagement.TransactOpts, accessibleNetworkId, networkName, networkIP, networkAddress, companyName)
}
