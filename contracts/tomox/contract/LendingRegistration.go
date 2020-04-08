// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package contract

import (
	"math/big"
	"strings"

	"github.com/tomochain/tomochain/accounts/abi"
	"github.com/tomochain/tomochain/accounts/abi/bind"
	"github.com/tomochain/tomochain/common"
	"github.com/tomochain/tomochain/core/types"
)

// LAbstractRegistrationABI is the input ABI used to generate the binding from.
const LAbstractRegistrationABI = "[{\"constant\":true,\"inputs\":[{\"name\":\"\",\"type\":\"address\"}],\"name\":\"RESIGN_REQUESTS\",\"outputs\":[{\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"name\":\"\",\"type\":\"address\"}],\"name\":\"getRelayerByCoinbase\",\"outputs\":[{\"name\":\"\",\"type\":\"uint256\"},{\"name\":\"\",\"type\":\"address\"},{\"name\":\"\",\"type\":\"uint256\"},{\"name\":\"\",\"type\":\"uint16\"},{\"name\":\"\",\"type\":\"address[]\"},{\"name\":\"\",\"type\":\"address[]\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"}]"

// LAbstractRegistrationBin is the compiled bytecode used for deploying new contracts.
const LAbstractRegistrationBin = `0x`

// DeployLAbstractRegistration deploys a new Ethereum contract, binding an instance of LAbstractRegistration to it.
func DeployLAbstractRegistration(auth *bind.TransactOpts, backend bind.ContractBackend) (common.Address, *types.Transaction, *LAbstractRegistration, error) {
	parsed, err := abi.JSON(strings.NewReader(LAbstractRegistrationABI))
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	address, tx, contract, err := bind.DeployContract(auth, parsed, common.FromHex(LAbstractRegistrationBin), backend)
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	return address, tx, &LAbstractRegistration{LAbstractRegistrationCaller: LAbstractRegistrationCaller{contract: contract}, LAbstractRegistrationTransactor: LAbstractRegistrationTransactor{contract: contract}, LAbstractRegistrationFilterer: LAbstractRegistrationFilterer{contract: contract}}, nil
}

// LAbstractRegistration is an auto generated Go binding around an Ethereum contract.
type LAbstractRegistration struct {
	LAbstractRegistrationCaller     // Read-only binding to the contract
	LAbstractRegistrationTransactor // Write-only binding to the contract
	LAbstractRegistrationFilterer   // Log filterer for contract events
}

// LAbstractRegistrationCaller is an auto generated read-only Go binding around an Ethereum contract.
type LAbstractRegistrationCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// LAbstractRegistrationTransactor is an auto generated write-only Go binding around an Ethereum contract.
type LAbstractRegistrationTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// LAbstractRegistrationFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type LAbstractRegistrationFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// LAbstractRegistrationSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type LAbstractRegistrationSession struct {
	Contract     *LAbstractRegistration // Generic contract binding to set the session for
	CallOpts     bind.CallOpts          // Call options to use throughout this session
	TransactOpts bind.TransactOpts      // Transaction auth options to use throughout this session
}

// LAbstractRegistrationCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type LAbstractRegistrationCallerSession struct {
	Contract *LAbstractRegistrationCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts                // Call options to use throughout this session
}

// LAbstractRegistrationTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type LAbstractRegistrationTransactorSession struct {
	Contract     *LAbstractRegistrationTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts                // Transaction auth options to use throughout this session
}

// LAbstractRegistrationRaw is an auto generated low-level Go binding around an Ethereum contract.
type LAbstractRegistrationRaw struct {
	Contract *LAbstractRegistration // Generic contract binding to access the raw methods on
}

// LAbstractRegistrationCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type LAbstractRegistrationCallerRaw struct {
	Contract *LAbstractRegistrationCaller // Generic read-only contract binding to access the raw methods on
}

// LAbstractRegistrationTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type LAbstractRegistrationTransactorRaw struct {
	Contract *LAbstractRegistrationTransactor // Generic write-only contract binding to access the raw methods on
}

// NewLAbstractRegistration creates a new instance of LAbstractRegistration, bound to a specific deployed contract.
func NewLAbstractRegistration(address common.Address, backend bind.ContractBackend) (*LAbstractRegistration, error) {
	contract, err := bindLAbstractRegistration(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &LAbstractRegistration{LAbstractRegistrationCaller: LAbstractRegistrationCaller{contract: contract}, LAbstractRegistrationTransactor: LAbstractRegistrationTransactor{contract: contract}, LAbstractRegistrationFilterer: LAbstractRegistrationFilterer{contract: contract}}, nil
}

// NewLAbstractRegistrationCaller creates a new read-only instance of LAbstractRegistration, bound to a specific deployed contract.
func NewLAbstractRegistrationCaller(address common.Address, caller bind.ContractCaller) (*LAbstractRegistrationCaller, error) {
	contract, err := bindLAbstractRegistration(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &LAbstractRegistrationCaller{contract: contract}, nil
}

// NewLAbstractRegistrationTransactor creates a new write-only instance of LAbstractRegistration, bound to a specific deployed contract.
func NewLAbstractRegistrationTransactor(address common.Address, transactor bind.ContractTransactor) (*LAbstractRegistrationTransactor, error) {
	contract, err := bindLAbstractRegistration(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &LAbstractRegistrationTransactor{contract: contract}, nil
}

// NewLAbstractRegistrationFilterer creates a new log filterer instance of LAbstractRegistration, bound to a specific deployed contract.
func NewLAbstractRegistrationFilterer(address common.Address, filterer bind.ContractFilterer) (*LAbstractRegistrationFilterer, error) {
	contract, err := bindLAbstractRegistration(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &LAbstractRegistrationFilterer{contract: contract}, nil
}

// bindLAbstractRegistration binds a generic wrapper to an already deployed contract.
func bindLAbstractRegistration(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(LAbstractRegistrationABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_LAbstractRegistration *LAbstractRegistrationRaw) Call(opts *bind.CallOpts, result interface{}, method string, params ...interface{}) error {
	return _LAbstractRegistration.Contract.LAbstractRegistrationCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_LAbstractRegistration *LAbstractRegistrationRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _LAbstractRegistration.Contract.LAbstractRegistrationTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_LAbstractRegistration *LAbstractRegistrationRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _LAbstractRegistration.Contract.LAbstractRegistrationTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_LAbstractRegistration *LAbstractRegistrationCallerRaw) Call(opts *bind.CallOpts, result interface{}, method string, params ...interface{}) error {
	return _LAbstractRegistration.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_LAbstractRegistration *LAbstractRegistrationTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _LAbstractRegistration.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_LAbstractRegistration *LAbstractRegistrationTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _LAbstractRegistration.Contract.contract.Transact(opts, method, params...)
}

// RESIGNREQUESTS is a free data retrieval call binding the contract method 0x500f99f7.
//
// Solidity: function RESIGN_REQUESTS( address) constant returns(uint256)
func (_LAbstractRegistration *LAbstractRegistrationCaller) RESIGNREQUESTS(opts *bind.CallOpts, arg0 common.Address) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _LAbstractRegistration.contract.Call(opts, out, "RESIGN_REQUESTS", arg0)
	return *ret0, err
}

// RESIGNREQUESTS is a free data retrieval call binding the contract method 0x500f99f7.
//
// Solidity: function RESIGN_REQUESTS( address) constant returns(uint256)
func (_LAbstractRegistration *LAbstractRegistrationSession) RESIGNREQUESTS(arg0 common.Address) (*big.Int, error) {
	return _LAbstractRegistration.Contract.RESIGNREQUESTS(&_LAbstractRegistration.CallOpts, arg0)
}

// RESIGNREQUESTS is a free data retrieval call binding the contract method 0x500f99f7.
//
// Solidity: function RESIGN_REQUESTS( address) constant returns(uint256)
func (_LAbstractRegistration *LAbstractRegistrationCallerSession) RESIGNREQUESTS(arg0 common.Address) (*big.Int, error) {
	return _LAbstractRegistration.Contract.RESIGNREQUESTS(&_LAbstractRegistration.CallOpts, arg0)
}

// GetRelayerByCoinbase is a free data retrieval call binding the contract method 0x540105c7.
//
// Solidity: function getRelayerByCoinbase( address) constant returns(uint256, address, uint256, uint16, address[], address[])
func (_LAbstractRegistration *LAbstractRegistrationCaller) GetRelayerByCoinbase(opts *bind.CallOpts, arg0 common.Address) (*big.Int, common.Address, *big.Int, uint16, []common.Address, []common.Address, error) {
	var (
		ret0 = new(*big.Int)
		ret1 = new(common.Address)
		ret2 = new(*big.Int)
		ret3 = new(uint16)
		ret4 = new([]common.Address)
		ret5 = new([]common.Address)
	)
	out := &[]interface{}{
		ret0,
		ret1,
		ret2,
		ret3,
		ret4,
		ret5,
	}
	err := _LAbstractRegistration.contract.Call(opts, out, "getRelayerByCoinbase", arg0)
	return *ret0, *ret1, *ret2, *ret3, *ret4, *ret5, err
}

// GetRelayerByCoinbase is a free data retrieval call binding the contract method 0x540105c7.
//
// Solidity: function getRelayerByCoinbase( address) constant returns(uint256, address, uint256, uint16, address[], address[])
func (_LAbstractRegistration *LAbstractRegistrationSession) GetRelayerByCoinbase(arg0 common.Address) (*big.Int, common.Address, *big.Int, uint16, []common.Address, []common.Address, error) {
	return _LAbstractRegistration.Contract.GetRelayerByCoinbase(&_LAbstractRegistration.CallOpts, arg0)
}

// GetRelayerByCoinbase is a free data retrieval call binding the contract method 0x540105c7.
//
// Solidity: function getRelayerByCoinbase( address) constant returns(uint256, address, uint256, uint16, address[], address[])
func (_LAbstractRegistration *LAbstractRegistrationCallerSession) GetRelayerByCoinbase(arg0 common.Address) (*big.Int, common.Address, *big.Int, uint16, []common.Address, []common.Address, error) {
	return _LAbstractRegistration.Contract.GetRelayerByCoinbase(&_LAbstractRegistration.CallOpts, arg0)
}

// LAbstractTOMOXListingABI is the input ABI used to generate the binding from.
const LAbstractTOMOXListingABI = "[{\"constant\":true,\"inputs\":[{\"name\":\"\",\"type\":\"address\"}],\"name\":\"getTokenStatus\",\"outputs\":[{\"name\":\"\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"}]"

// LAbstractTOMOXListingBin is the compiled bytecode used for deploying new contracts.
const LAbstractTOMOXListingBin = `0x`

// DeployLAbstractTOMOXListing deploys a new Ethereum contract, binding an instance of LAbstractTOMOXListing to it.
func DeployLAbstractTOMOXListing(auth *bind.TransactOpts, backend bind.ContractBackend) (common.Address, *types.Transaction, *LAbstractTOMOXListing, error) {
	parsed, err := abi.JSON(strings.NewReader(LAbstractTOMOXListingABI))
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	address, tx, contract, err := bind.DeployContract(auth, parsed, common.FromHex(LAbstractTOMOXListingBin), backend)
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	return address, tx, &LAbstractTOMOXListing{LAbstractTOMOXListingCaller: LAbstractTOMOXListingCaller{contract: contract}, LAbstractTOMOXListingTransactor: LAbstractTOMOXListingTransactor{contract: contract}, LAbstractTOMOXListingFilterer: LAbstractTOMOXListingFilterer{contract: contract}}, nil
}

// LAbstractTOMOXListing is an auto generated Go binding around an Ethereum contract.
type LAbstractTOMOXListing struct {
	LAbstractTOMOXListingCaller     // Read-only binding to the contract
	LAbstractTOMOXListingTransactor // Write-only binding to the contract
	LAbstractTOMOXListingFilterer   // Log filterer for contract events
}

// LAbstractTOMOXListingCaller is an auto generated read-only Go binding around an Ethereum contract.
type LAbstractTOMOXListingCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// LAbstractTOMOXListingTransactor is an auto generated write-only Go binding around an Ethereum contract.
type LAbstractTOMOXListingTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// LAbstractTOMOXListingFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type LAbstractTOMOXListingFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// LAbstractTOMOXListingSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type LAbstractTOMOXListingSession struct {
	Contract     *LAbstractTOMOXListing // Generic contract binding to set the session for
	CallOpts     bind.CallOpts          // Call options to use throughout this session
	TransactOpts bind.TransactOpts      // Transaction auth options to use throughout this session
}

// LAbstractTOMOXListingCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type LAbstractTOMOXListingCallerSession struct {
	Contract *LAbstractTOMOXListingCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts                // Call options to use throughout this session
}

// LAbstractTOMOXListingTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type LAbstractTOMOXListingTransactorSession struct {
	Contract     *LAbstractTOMOXListingTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts                // Transaction auth options to use throughout this session
}

// LAbstractTOMOXListingRaw is an auto generated low-level Go binding around an Ethereum contract.
type LAbstractTOMOXListingRaw struct {
	Contract *LAbstractTOMOXListing // Generic contract binding to access the raw methods on
}

// LAbstractTOMOXListingCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type LAbstractTOMOXListingCallerRaw struct {
	Contract *LAbstractTOMOXListingCaller // Generic read-only contract binding to access the raw methods on
}

// LAbstractTOMOXListingTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type LAbstractTOMOXListingTransactorRaw struct {
	Contract *LAbstractTOMOXListingTransactor // Generic write-only contract binding to access the raw methods on
}

// NewLAbstractTOMOXListing creates a new instance of LAbstractTOMOXListing, bound to a specific deployed contract.
func NewLAbstractTOMOXListing(address common.Address, backend bind.ContractBackend) (*LAbstractTOMOXListing, error) {
	contract, err := bindLAbstractTOMOXListing(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &LAbstractTOMOXListing{LAbstractTOMOXListingCaller: LAbstractTOMOXListingCaller{contract: contract}, LAbstractTOMOXListingTransactor: LAbstractTOMOXListingTransactor{contract: contract}, LAbstractTOMOXListingFilterer: LAbstractTOMOXListingFilterer{contract: contract}}, nil
}

// NewLAbstractTOMOXListingCaller creates a new read-only instance of LAbstractTOMOXListing, bound to a specific deployed contract.
func NewLAbstractTOMOXListingCaller(address common.Address, caller bind.ContractCaller) (*LAbstractTOMOXListingCaller, error) {
	contract, err := bindLAbstractTOMOXListing(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &LAbstractTOMOXListingCaller{contract: contract}, nil
}

// NewLAbstractTOMOXListingTransactor creates a new write-only instance of LAbstractTOMOXListing, bound to a specific deployed contract.
func NewLAbstractTOMOXListingTransactor(address common.Address, transactor bind.ContractTransactor) (*LAbstractTOMOXListingTransactor, error) {
	contract, err := bindLAbstractTOMOXListing(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &LAbstractTOMOXListingTransactor{contract: contract}, nil
}

// NewLAbstractTOMOXListingFilterer creates a new log filterer instance of LAbstractTOMOXListing, bound to a specific deployed contract.
func NewLAbstractTOMOXListingFilterer(address common.Address, filterer bind.ContractFilterer) (*LAbstractTOMOXListingFilterer, error) {
	contract, err := bindLAbstractTOMOXListing(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &LAbstractTOMOXListingFilterer{contract: contract}, nil
}

// bindLAbstractTOMOXListing binds a generic wrapper to an already deployed contract.
func bindLAbstractTOMOXListing(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(LAbstractTOMOXListingABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_LAbstractTOMOXListing *LAbstractTOMOXListingRaw) Call(opts *bind.CallOpts, result interface{}, method string, params ...interface{}) error {
	return _LAbstractTOMOXListing.Contract.LAbstractTOMOXListingCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_LAbstractTOMOXListing *LAbstractTOMOXListingRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _LAbstractTOMOXListing.Contract.LAbstractTOMOXListingTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_LAbstractTOMOXListing *LAbstractTOMOXListingRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _LAbstractTOMOXListing.Contract.LAbstractTOMOXListingTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_LAbstractTOMOXListing *LAbstractTOMOXListingCallerRaw) Call(opts *bind.CallOpts, result interface{}, method string, params ...interface{}) error {
	return _LAbstractTOMOXListing.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_LAbstractTOMOXListing *LAbstractTOMOXListingTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _LAbstractTOMOXListing.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_LAbstractTOMOXListing *LAbstractTOMOXListingTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _LAbstractTOMOXListing.Contract.contract.Transact(opts, method, params...)
}

// GetTokenStatus is a free data retrieval call binding the contract method 0xa3ff31b5.
//
// Solidity: function getTokenStatus( address) constant returns(bool)
func (_LAbstractTOMOXListing *LAbstractTOMOXListingCaller) GetTokenStatus(opts *bind.CallOpts, arg0 common.Address) (bool, error) {
	var (
		ret0 = new(bool)
	)
	out := ret0
	err := _LAbstractTOMOXListing.contract.Call(opts, out, "getTokenStatus", arg0)
	return *ret0, err
}

// GetTokenStatus is a free data retrieval call binding the contract method 0xa3ff31b5.
//
// Solidity: function getTokenStatus( address) constant returns(bool)
func (_LAbstractTOMOXListing *LAbstractTOMOXListingSession) GetTokenStatus(arg0 common.Address) (bool, error) {
	return _LAbstractTOMOXListing.Contract.GetTokenStatus(&_LAbstractTOMOXListing.CallOpts, arg0)
}

// GetTokenStatus is a free data retrieval call binding the contract method 0xa3ff31b5.
//
// Solidity: function getTokenStatus( address) constant returns(bool)
func (_LAbstractTOMOXListing *LAbstractTOMOXListingCallerSession) GetTokenStatus(arg0 common.Address) (bool, error) {
	return _LAbstractTOMOXListing.Contract.GetTokenStatus(&_LAbstractTOMOXListing.CallOpts, arg0)
}

// LAbstractTokenTRC21ABI is the input ABI used to generate the binding from.
const LAbstractTokenTRC21ABI = "[{\"constant\":true,\"inputs\":[],\"name\":\"issuer\",\"outputs\":[{\"name\":\"\",\"type\":\"address\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"}]"

// LAbstractTokenTRC21Bin is the compiled bytecode used for deploying new contracts.
const LAbstractTokenTRC21Bin = `0x`

// DeployLAbstractTokenTRC21 deploys a new Ethereum contract, binding an instance of LAbstractTokenTRC21 to it.
func DeployLAbstractTokenTRC21(auth *bind.TransactOpts, backend bind.ContractBackend) (common.Address, *types.Transaction, *LAbstractTokenTRC21, error) {
	parsed, err := abi.JSON(strings.NewReader(LAbstractTokenTRC21ABI))
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	address, tx, contract, err := bind.DeployContract(auth, parsed, common.FromHex(LAbstractTokenTRC21Bin), backend)
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	return address, tx, &LAbstractTokenTRC21{LAbstractTokenTRC21Caller: LAbstractTokenTRC21Caller{contract: contract}, LAbstractTokenTRC21Transactor: LAbstractTokenTRC21Transactor{contract: contract}, LAbstractTokenTRC21Filterer: LAbstractTokenTRC21Filterer{contract: contract}}, nil
}

// LAbstractTokenTRC21 is an auto generated Go binding around an Ethereum contract.
type LAbstractTokenTRC21 struct {
	LAbstractTokenTRC21Caller     // Read-only binding to the contract
	LAbstractTokenTRC21Transactor // Write-only binding to the contract
	LAbstractTokenTRC21Filterer   // Log filterer for contract events
}

// LAbstractTokenTRC21Caller is an auto generated read-only Go binding around an Ethereum contract.
type LAbstractTokenTRC21Caller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// LAbstractTokenTRC21Transactor is an auto generated write-only Go binding around an Ethereum contract.
type LAbstractTokenTRC21Transactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// LAbstractTokenTRC21Filterer is an auto generated log filtering Go binding around an Ethereum contract events.
type LAbstractTokenTRC21Filterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// LAbstractTokenTRC21Session is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type LAbstractTokenTRC21Session struct {
	Contract     *LAbstractTokenTRC21 // Generic contract binding to set the session for
	CallOpts     bind.CallOpts        // Call options to use throughout this session
	TransactOpts bind.TransactOpts    // Transaction auth options to use throughout this session
}

// LAbstractTokenTRC21CallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type LAbstractTokenTRC21CallerSession struct {
	Contract *LAbstractTokenTRC21Caller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts              // Call options to use throughout this session
}

// LAbstractTokenTRC21TransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type LAbstractTokenTRC21TransactorSession struct {
	Contract     *LAbstractTokenTRC21Transactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts              // Transaction auth options to use throughout this session
}

// LAbstractTokenTRC21Raw is an auto generated low-level Go binding around an Ethereum contract.
type LAbstractTokenTRC21Raw struct {
	Contract *LAbstractTokenTRC21 // Generic contract binding to access the raw methods on
}

// LAbstractTokenTRC21CallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type LAbstractTokenTRC21CallerRaw struct {
	Contract *LAbstractTokenTRC21Caller // Generic read-only contract binding to access the raw methods on
}

// LAbstractTokenTRC21TransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type LAbstractTokenTRC21TransactorRaw struct {
	Contract *LAbstractTokenTRC21Transactor // Generic write-only contract binding to access the raw methods on
}

// NewLAbstractTokenTRC21 creates a new instance of LAbstractTokenTRC21, bound to a specific deployed contract.
func NewLAbstractTokenTRC21(address common.Address, backend bind.ContractBackend) (*LAbstractTokenTRC21, error) {
	contract, err := bindLAbstractTokenTRC21(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &LAbstractTokenTRC21{LAbstractTokenTRC21Caller: LAbstractTokenTRC21Caller{contract: contract}, LAbstractTokenTRC21Transactor: LAbstractTokenTRC21Transactor{contract: contract}, LAbstractTokenTRC21Filterer: LAbstractTokenTRC21Filterer{contract: contract}}, nil
}

// NewLAbstractTokenTRC21Caller creates a new read-only instance of LAbstractTokenTRC21, bound to a specific deployed contract.
func NewLAbstractTokenTRC21Caller(address common.Address, caller bind.ContractCaller) (*LAbstractTokenTRC21Caller, error) {
	contract, err := bindLAbstractTokenTRC21(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &LAbstractTokenTRC21Caller{contract: contract}, nil
}

// NewLAbstractTokenTRC21Transactor creates a new write-only instance of LAbstractTokenTRC21, bound to a specific deployed contract.
func NewLAbstractTokenTRC21Transactor(address common.Address, transactor bind.ContractTransactor) (*LAbstractTokenTRC21Transactor, error) {
	contract, err := bindLAbstractTokenTRC21(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &LAbstractTokenTRC21Transactor{contract: contract}, nil
}

// NewLAbstractTokenTRC21Filterer creates a new log filterer instance of LAbstractTokenTRC21, bound to a specific deployed contract.
func NewLAbstractTokenTRC21Filterer(address common.Address, filterer bind.ContractFilterer) (*LAbstractTokenTRC21Filterer, error) {
	contract, err := bindLAbstractTokenTRC21(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &LAbstractTokenTRC21Filterer{contract: contract}, nil
}

// bindLAbstractTokenTRC21 binds a generic wrapper to an already deployed contract.
func bindLAbstractTokenTRC21(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(LAbstractTokenTRC21ABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_LAbstractTokenTRC21 *LAbstractTokenTRC21Raw) Call(opts *bind.CallOpts, result interface{}, method string, params ...interface{}) error {
	return _LAbstractTokenTRC21.Contract.LAbstractTokenTRC21Caller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_LAbstractTokenTRC21 *LAbstractTokenTRC21Raw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _LAbstractTokenTRC21.Contract.LAbstractTokenTRC21Transactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_LAbstractTokenTRC21 *LAbstractTokenTRC21Raw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _LAbstractTokenTRC21.Contract.LAbstractTokenTRC21Transactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_LAbstractTokenTRC21 *LAbstractTokenTRC21CallerRaw) Call(opts *bind.CallOpts, result interface{}, method string, params ...interface{}) error {
	return _LAbstractTokenTRC21.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_LAbstractTokenTRC21 *LAbstractTokenTRC21TransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _LAbstractTokenTRC21.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_LAbstractTokenTRC21 *LAbstractTokenTRC21TransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _LAbstractTokenTRC21.Contract.contract.Transact(opts, method, params...)
}

// Issuer is a free data retrieval call binding the contract method 0x1d143848.
//
// Solidity: function issuer() constant returns(address)
func (_LAbstractTokenTRC21 *LAbstractTokenTRC21Caller) Issuer(opts *bind.CallOpts) (common.Address, error) {
	var (
		ret0 = new(common.Address)
	)
	out := ret0
	err := _LAbstractTokenTRC21.contract.Call(opts, out, "issuer")
	return *ret0, err
}

// Issuer is a free data retrieval call binding the contract method 0x1d143848.
//
// Solidity: function issuer() constant returns(address)
func (_LAbstractTokenTRC21 *LAbstractTokenTRC21Session) Issuer() (common.Address, error) {
	return _LAbstractTokenTRC21.Contract.Issuer(&_LAbstractTokenTRC21.CallOpts)
}

// Issuer is a free data retrieval call binding the contract method 0x1d143848.
//
// Solidity: function issuer() constant returns(address)
func (_LAbstractTokenTRC21 *LAbstractTokenTRC21CallerSession) Issuer() (common.Address, error) {
	return _LAbstractTokenTRC21.Contract.Issuer(&_LAbstractTokenTRC21.CallOpts)
}

// LendingABI is the input ABI used to generate the binding from.
const LendingABI = "[{\"constant\":true,\"inputs\":[{\"name\":\"\",\"type\":\"uint256\"}],\"name\":\"COLLATERALS\",\"outputs\":[{\"name\":\"\",\"type\":\"address\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"term\",\"type\":\"uint256\"}],\"name\":\"addTerm\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"name\":\"\",\"type\":\"address\"}],\"name\":\"LENDINGRELAYER_LIST\",\"outputs\":[{\"name\":\"_tradeFee\",\"type\":\"uint16\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"Relayer\",\"outputs\":[{\"name\":\"\",\"type\":\"address\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"TomoXListing\",\"outputs\":[{\"name\":\"\",\"type\":\"address\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"coinbase\",\"type\":\"address\"},{\"name\":\"tradeFee\",\"type\":\"uint16\"},{\"name\":\"baseTokens\",\"type\":\"address[]\"},{\"name\":\"terms\",\"type\":\"uint256[]\"},{\"name\":\"collaterals\",\"type\":\"address[]\"}],\"name\":\"update\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"token\",\"type\":\"address\"},{\"name\":\"depositRate\",\"type\":\"uint256\"},{\"name\":\"liquidationRate\",\"type\":\"uint256\"},{\"name\":\"recallRate\",\"type\":\"uint256\"}],\"name\":\"addILOCollateral\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"name\":\"\",\"type\":\"uint256\"}],\"name\":\"TERMS\",\"outputs\":[{\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"name\":\"\",\"type\":\"uint256\"}],\"name\":\"BASES\",\"outputs\":[{\"name\":\"\",\"type\":\"address\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"name\":\"\",\"type\":\"address\"}],\"name\":\"COLLATERAL_LIST\",\"outputs\":[{\"name\":\"_depositRate\",\"type\":\"uint256\"},{\"name\":\"_liquidationRate\",\"type\":\"uint256\"},{\"name\":\"_recallRate\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"token\",\"type\":\"address\"}],\"name\":\"addBaseToken\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"token\",\"type\":\"address\"},{\"name\":\"lendingToken\",\"type\":\"address\"},{\"name\":\"price\",\"type\":\"uint256\"}],\"name\":\"setCollateralPrice\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"name\":\"\",\"type\":\"uint256\"}],\"name\":\"ILO_COLLATERALS\",\"outputs\":[{\"name\":\"\",\"type\":\"address\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"token\",\"type\":\"address\"},{\"name\":\"depositRate\",\"type\":\"uint256\"},{\"name\":\"liquidationRate\",\"type\":\"uint256\"},{\"name\":\"recallRate\",\"type\":\"uint256\"}],\"name\":\"addCollateral\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"CONTRACT_OWNER\",\"outputs\":[{\"name\":\"\",\"type\":\"address\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"name\":\"coinbase\",\"type\":\"address\"}],\"name\":\"getLendingRelayerByCoinbase\",\"outputs\":[{\"name\":\"\",\"type\":\"uint16\"},{\"name\":\"\",\"type\":\"address[]\"},{\"name\":\"\",\"type\":\"uint256[]\"},{\"name\":\"\",\"type\":\"address[]\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"name\":\"r\",\"type\":\"address\"},{\"name\":\"t\",\"type\":\"address\"}],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"constructor\"}]"

// LendingBin is the compiled bytecode used for deploying new contracts.
const LendingBin = `0x608060405234801561001057600080fd5b50604051604080611e0f83398101604052805160209091015160068054600160a060020a03938416600160a060020a03199182161790915560088054939092169281169290921790556007805490911633179055611d9c806100736000396000f3006080604052600436106100cc5763ffffffff60e060020a6000350416630811f05a81146100d15780630c655955146101055780630faf292c1461011f578063264949d81461015757806329a4ddec1461016c5780632ddada4c146101815780633b8748271461025e57806356327f57146102885780636d1dc42a146102b257806382250701146102ca57806383e280d914610309578063acb8cd921461032a578063b8687ec414610354578063e5eecf681461036c578063fd301c4914610396578063fe824700146103ab575b600080fd5b3480156100dd57600080fd5b506100e96004356104b9565b60408051600160a060020a039092168252519081900360200190f35b34801561011157600080fd5b5061011d6004356104e1565b005b34801561012b57600080fd5b50610140600160a060020a0360043516610635565b6040805161ffff9092168252519081900360200190f35b34801561016357600080fd5b506100e961064b565b34801561017857600080fd5b506100e961065a565b34801561018d57600080fd5b50604080516020600460443581810135838102808601850190965280855261011d958335600160a060020a0316956024803561ffff1696369695606495939492019291829185019084908082843750506040805187358901803560208181028481018201909552818452989b9a998901989297509082019550935083925085019084908082843750506040805187358901803560208181028481018201909552818452989b9a9989019892975090820195509350839250850190849080828437509497506106699650505050505050565b34801561026a57600080fd5b5061011d600160a060020a0360043516602435604435606435610da8565b34801561029457600080fd5b506102a0600435611167565b60408051918252519081900360200190f35b3480156102be57600080fd5b506100e9600435611186565b3480156102d657600080fd5b506102eb600160a060020a0360043516611194565b60408051938452602084019290925282820152519081900360600190f35b34801561031557600080fd5b5061011d600160a060020a03600435166111b4565b34801561033657600080fd5b5061011d600160a060020a03600435811690602435166044356113cb565b34801561036057600080fd5b506100e96004356116f2565b34801561037857600080fd5b5061011d600160a060020a0360043516602435604435606435611700565b3480156103a257600080fd5b506100e9611a59565b3480156103b757600080fd5b506103cc600160a060020a0360043516611a68565b604051808561ffff1661ffff168152602001806020018060200180602001848103845287818151815260200191508051906020019060200280838360005b8381101561042257818101518382015260200161040a565b50505050905001848103835286818151815260200191508051906020019060200280838360005b83811015610461578181015183820152602001610449565b50505050905001848103825285818151815260200191508051906020019060200280838360005b838110156104a0578181015183820152602001610488565b5050505090500197505050505050505060405180910390f35b60028054829081106104c757fe5b600091825260209091200154600160a060020a0316905081565b600754600160a060020a03163314610543576040805160e560020a62461bcd02815260206004820152601460248201527f436f6e7472616374204f776e6572204f6e6c792e000000000000000000000000604482015290519081900360640190fd5b603c81101561059c576040805160e560020a62461bcd02815260206004820152600c60248201527f496e76616c6964207465726d0000000000000000000000000000000000000000604482015290519081900360640190fd5b6105f660048054806020026020016040519081016040528092919081815260200182805480156105eb57602002820191906000526020600020905b8154815260200190600101908083116105d7575b505050505082611bb1565b151561063257600480546001810182556000919091527f8a35acfbc15ff81a39ae7d344fd709f28e8600b4aa8c65c6b64bfe7fe36bd19b018190555b50565b60006020819052908152604090205461ffff1681565b600654600160a060020a031681565b600854600160a060020a031681565b600654604080517f540105c7000000000000000000000000000000000000000000000000000000008152600160a060020a03888116600483015291516000938493849391169163540105c791602480820192869290919082900301818387803b1580156106d557600080fd5b505af11580156106e9573d6000803e3d6000fd5b505050506040513d6000823e601f3d908101601f1916820160405260c081101561071257600080fd5b81516020830151604084015160608501516080860180519496939592949193928301929164010000000081111561074857600080fd5b8201602081018481111561075b57600080fd5b815185602082028301116401000000008211171561077857600080fd5b5050929190602001805164010000000081111561079457600080fd5b820160208101848111156107a757600080fd5b81518560208202830111640100000000821117156107c457600080fd5b50979b505050600160a060020a038a163314965061083395505050505050576040805160e560020a62461bcd02815260206004820152601660248201527f52656c61796572206f776e657220726571756972656400000000000000000000604482015290519081900360640190fd5b600654604080517f500f99f7000000000000000000000000000000000000000000000000000000008152600160a060020a038b811660048301529151919092169163500f99f79160248083019260209291908290030181600087803b15801561089b57600080fd5b505af11580156108af573d6000803e3d6000fd5b505050506040513d60208110156108c557600080fd5b50511561091c576040805160e560020a62461bcd02815260206004820152601960248201527f52656c6179657220726571756972656420746f20636c6f736500000000000000604482015290519081900360640190fd5b60008761ffff161015801561093657506103e88761ffff16105b151561098c576040805160e560020a62461bcd02815260206004820152601160248201527f496e76616c696420747261646520466565000000000000000000000000000000604482015290519081900360640190fd5b84518651146109e5576040805160e560020a62461bcd02815260206004820152601960248201527f4e6f742076616c6964206e756d626572206f66207465726d7300000000000000604482015290519081900360640190fd5b8351865114610a3e576040805160e560020a62461bcd02815260206004820152601f60248201527f4e6f742076616c6964206e756d626572206f6620636f6c6c61746572616c7300604482015290519081900360640190fd5b5060009050805b8551811015610b2d57610ac96003805480602002602001604051908101604052809291908181526020018280548015610aa757602002820191906000526020600020905b8154600160a060020a03168152600190910190602001808311610a89575b50505050508783815181101515610aba57fe5b90602001906020020151611bfa565b9150600182151514610b25576040805160e560020a62461bcd02815260206004820152601560248201527f496e76616c6964206c656e64696e6720746f6b656e0000000000000000000000604482015290519081900360640190fd5b600101610a45565b5060005b8451811015610c0f57610bab6004805480602002602001604051908101604052809291908181526020018280548015610b8957602002820191906000526020600020905b815481526020019060010190808311610b75575b50505050508683815181101515610b9c57fe5b90602001906020020151611bb1565b9150600182151514610c07576040805160e560020a62461bcd02815260206004820152600c60248201527f496e76616c6964207465726d0000000000000000000000000000000000000000604482015290519081900360640190fd5b600101610b31565b5060005b8351811015610cfd578351600090859083908110610c2d57fe5b60209081029091010151600160a060020a031614610cf557610cb16005805480602002602001604051908101604052809291908181526020018280548015610c9e57602002820191906000526020600020905b8154600160a060020a03168152600190910190602001808311610c80575b50505050508583815181101515610aba57fe5b1515610cf5576040805160e560020a62461bcd0281526020600482015260126024820152600080516020611d51833981519152604482015290519081900360640190fd5b600101610c13565b6040805160808101825261ffff898116825260208083018a81528385018a905260608401899052600160a060020a038d166000908152808352949094208351815461ffff191693169290921782559251805192939192610d639260018501920190611c49565b5060408201518051610d7f916002840191602090910190611cbb565b5060608201518051610d9b916003840191602090910190611c49565b5050505050505050505050565b60008060648510158015610dbc5750606484115b1515610e12576040805160e560020a62461bcd02815260206004820152600d60248201527f496e76616c696420726174657300000000000000000000000000000000000000604482015290519081900360640190fd5b838511610e69576040805160e560020a62461bcd02815260206004820152601560248201527f496e76616c6964206465706f7369742072617465730000000000000000000000604482015290519081900360640190fd5b848311610ec0576040805160e560020a62461bcd02815260206004820152601460248201527f496e76616c696420726563616c6c207261746573000000000000000000000000604482015290519081900360640190fd5b6008546040805160e060020a63a3ff31b5028152600160a060020a0389811660048301529151919092169163a3ff31b59160248083019260209291908290030181600087803b158015610f1257600080fd5b505af1158015610f26573d6000803e3d6000fd5b505050506040513d6020811015610f3c57600080fd5b50519150811515610f85576040805160e560020a62461bcd0281526020600482015260126024820152600080516020611d51833981519152604482015290519081900360640190fd5b85905033600160a060020a031681600160a060020a0316631d1438486040518163ffffffff1660e060020a028152600401602060405180830381600087803b158015610fd057600080fd5b505af1158015610fe4573d6000803e3d6000fd5b505050506040513d6020811015610ffa57600080fd5b5051600160a060020a03161461105a576040805160e560020a62461bcd02815260206004820152601560248201527f526571756972656420746f6b656e206973737565720000000000000000000000604482015290519081900360640190fd5b604080516060810182528681526020808201878152828401878152600160a060020a038b16600090815260018085529086902094518555915191840191909155516002909201919091556005805483518184028101840190945280845261110093928301828280156110f557602002820191906000526020600020905b8154600160a060020a031681526001909101906020018083116110d7575b505050505087611bfa565b151561115f57600580546001810182556000919091527f036b6384b5eca791c62761152d0c79bb0604c104a5fb6f4eb0703f3154bb3db001805473ffffffffffffffffffffffffffffffffffffffff1916600160a060020a0388161790555b505050505050565b600480548290811061117557fe5b600091825260209091200154905081565b60038054829081106104c757fe5b600160208190526000918252604090912080549181015460029091015483565b600754600090600160a060020a03163314611219576040805160e560020a62461bcd02815260206004820152601460248201527f436f6e7472616374204f776e6572204f6e6c792e000000000000000000000000604482015290519081900360640190fd5b6008546040805160e060020a63a3ff31b5028152600160a060020a0385811660048301529151919092169163a3ff31b59160248083019260209291908290030181600087803b15801561126b57600080fd5b505af115801561127f573d6000803e3d6000fd5b505050506040513d602081101561129557600080fd5b5051806112ab5750600160a060020a0382166001145b9050801515611304576040805160e560020a62461bcd02815260206004820152601260248201527f496e76616c6964206261736520746f6b656e0000000000000000000000000000604482015290519081900360640190fd5b611368600380548060200260200160405190810160405280929190818152602001828054801561135d57602002820191906000526020600020905b8154600160a060020a0316815260019091019060200180831161133f575b505050505083611bfa565b15156113c757600380546001810182556000919091527fc2575a0e9e593c00f959f8c92f12db2869c3395a3b0502d05e2516446f71f85b01805473ffffffffffffffffffffffffffffffffffffffff1916600160a060020a0384161790555b5050565b6008546040805160e060020a63a3ff31b5028152600160a060020a03868116600483015291516000938493169163a3ff31b591602480830192602092919082900301818787803b15801561141e57600080fd5b505af1158015611432573d6000803e3d6000fd5b505050506040513d602081101561144857600080fd5b50518061145e5750600160a060020a0385166001145b91508115156114a5576040805160e560020a62461bcd0281526020600482015260126024820152600080516020611d51833981519152604482015290519081900360640190fd5b600160a060020a03851660009081526001602052604090205460641115611504576040805160e560020a62461bcd0281526020600482015260126024820152600080516020611d51833981519152604482015290519081900360640190fd5b611568600280548060200260200160405190810160405280929190818152602001828054801561155d57602002820191906000526020600020905b8154600160a060020a0316815260019091019060200180831161153f575b505050505086611bfa565b156115d457600754600160a060020a031633146115cf576040805160e560020a62461bcd02815260206004820152601760248201527f436f6e7472616374206f776e6572207265717569726564000000000000000000604482015290519081900360640190fd5b6116a9565b84905033600160a060020a031681600160a060020a0316631d1438486040518163ffffffff1660e060020a028152600401602060405180830381600087803b15801561161f57600080fd5b505af1158015611633573d6000803e3d6000fd5b505050506040513d602081101561164957600080fd5b5051600160a060020a0316146116a9576040805160e560020a62461bcd02815260206004820152601560248201527f526571756972656420746f6b656e206973737565720000000000000000000000604482015290519081900360640190fd5b5050604080518082018252918252436020808401918252600160a060020a0395861660009081526001808352848220969097168152600390950190529220905181559051910155565b60058054829081106104c757fe5b600754600090600160a060020a03163314611765576040805160e560020a62461bcd02815260206004820152601460248201527f436f6e7472616374204f776e6572204f6e6c792e000000000000000000000000604482015290519081900360640190fd5b606484101580156117765750606483115b15156117cc576040805160e560020a62461bcd02815260206004820152600d60248201527f496e76616c696420726174657300000000000000000000000000000000000000604482015290519081900360640190fd5b828411611823576040805160e560020a62461bcd02815260206004820152601560248201527f496e76616c6964206465706f7369742072617465730000000000000000000000604482015290519081900360640190fd5b83821161187a576040805160e560020a62461bcd02815260206004820152601460248201527f496e76616c696420726563616c6c207261746573000000000000000000000000604482015290519081900360640190fd5b6008546040805160e060020a63a3ff31b5028152600160a060020a0388811660048301529151919092169163a3ff31b59160248083019260209291908290030181600087803b1580156118cc57600080fd5b505af11580156118e0573d6000803e3d6000fd5b505050506040513d60208110156118f657600080fd5b50518061190c5750600160a060020a0385166001145b9050801515611953576040805160e560020a62461bcd0281526020600482015260126024820152600080516020611d51833981519152604482015290519081900360640190fd5b604080516060810182528581526020808201868152828401868152600160a060020a038a166000908152600180855290869020945185559151918401919091555160029283015581548351818302810183019094528084526119f393929183018282801561155d57602002820191906000526020600020908154600160a060020a0316815260019091019060200180831161153f57505050505086611bfa565b1515611a5257600280546001810182556000919091527f405787fa12a823e0f2b7631cc41b3ba8828b3321ca811111fa75cd3aa3bb5ace01805473ffffffffffffffffffffffffffffffffffffffff1916600160a060020a0387161790555b5050505050565b600754600160a060020a031681565b600160a060020a03811660009081526020818152604080832080546001820180548451818702810187019095528085526060958695869561ffff909516946002810193600390910192859190830182828015611aed57602002820191906000526020600020905b8154600160a060020a03168152600190910190602001808311611acf575b5050505050925081805480602002602001604051908101604052809291908181526020018280548015611b3f57602002820191906000526020600020905b815481526020019060010190808311611b2b575b5050505050915080805480602002602001604051908101604052809291908181526020018280548015611b9b57602002820191906000526020600020905b8154600160a060020a03168152600190910190602001808311611b7d575b5050505050905093509350935093509193509193565b6000805b8351811015611bee57828482815181101515611bcd57fe5b906020019060200201511415611be65760019150611bf3565b600101611bb5565b600091505b5092915050565b6000805b8351811015611bee5782600160a060020a03168482815181101515611c1f57fe5b90602001906020020151600160a060020a03161415611c415760019150611bf3565b600101611bfe565b828054828255906000526020600020908101928215611cab579160200282015b82811115611cab578251825473ffffffffffffffffffffffffffffffffffffffff1916600160a060020a03909116178255602090920191600190910190611c69565b50611cb7929150611d02565b5090565b828054828255906000526020600020908101928215611cf6579160200282015b82811115611cf6578251825591602001919060010190611cdb565b50611cb7929150611d36565b611d3391905b80821115611cb757805473ffffffffffffffffffffffffffffffffffffffff19168155600101611d08565b90565b611d3391905b80821115611cb75760008155600101611d3c5600496e76616c696420636f6c6c61746572616c0000000000000000000000000000a165627a7a72305820d78300513578cdc9238c16ef5404e780c1cc9b58fc61a93f1b69797f50ee1e740029`

// DeployLending deploys a new Ethereum contract, binding an instance of Lending to it.
func DeployLending(auth *bind.TransactOpts, backend bind.ContractBackend, r common.Address, t common.Address) (common.Address, *types.Transaction, *Lending, error) {
	parsed, err := abi.JSON(strings.NewReader(LendingABI))
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	address, tx, contract, err := bind.DeployContract(auth, parsed, common.FromHex(LendingBin), backend, r, t)
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	return address, tx, &Lending{LendingCaller: LendingCaller{contract: contract}, LendingTransactor: LendingTransactor{contract: contract}, LendingFilterer: LendingFilterer{contract: contract}}, nil
}

// Lending is an auto generated Go binding around an Ethereum contract.
type Lending struct {
	LendingCaller     // Read-only binding to the contract
	LendingTransactor // Write-only binding to the contract
	LendingFilterer   // Log filterer for contract events
}

// LendingCaller is an auto generated read-only Go binding around an Ethereum contract.
type LendingCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// LendingTransactor is an auto generated write-only Go binding around an Ethereum contract.
type LendingTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// LendingFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type LendingFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// LendingSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type LendingSession struct {
	Contract     *Lending          // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// LendingCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type LendingCallerSession struct {
	Contract *LendingCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts  // Call options to use throughout this session
}

// LendingTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type LendingTransactorSession struct {
	Contract     *LendingTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts  // Transaction auth options to use throughout this session
}

// LendingRaw is an auto generated low-level Go binding around an Ethereum contract.
type LendingRaw struct {
	Contract *Lending // Generic contract binding to access the raw methods on
}

// LendingCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type LendingCallerRaw struct {
	Contract *LendingCaller // Generic read-only contract binding to access the raw methods on
}

// LendingTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type LendingTransactorRaw struct {
	Contract *LendingTransactor // Generic write-only contract binding to access the raw methods on
}

// NewLending creates a new instance of Lending, bound to a specific deployed contract.
func NewLending(address common.Address, backend bind.ContractBackend) (*Lending, error) {
	contract, err := bindLending(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &Lending{LendingCaller: LendingCaller{contract: contract}, LendingTransactor: LendingTransactor{contract: contract}, LendingFilterer: LendingFilterer{contract: contract}}, nil
}

// NewLendingCaller creates a new read-only instance of Lending, bound to a specific deployed contract.
func NewLendingCaller(address common.Address, caller bind.ContractCaller) (*LendingCaller, error) {
	contract, err := bindLending(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &LendingCaller{contract: contract}, nil
}

// NewLendingTransactor creates a new write-only instance of Lending, bound to a specific deployed contract.
func NewLendingTransactor(address common.Address, transactor bind.ContractTransactor) (*LendingTransactor, error) {
	contract, err := bindLending(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &LendingTransactor{contract: contract}, nil
}

// NewLendingFilterer creates a new log filterer instance of Lending, bound to a specific deployed contract.
func NewLendingFilterer(address common.Address, filterer bind.ContractFilterer) (*LendingFilterer, error) {
	contract, err := bindLending(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &LendingFilterer{contract: contract}, nil
}

// bindLending binds a generic wrapper to an already deployed contract.
func bindLending(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(LendingABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Lending *LendingRaw) Call(opts *bind.CallOpts, result interface{}, method string, params ...interface{}) error {
	return _Lending.Contract.LendingCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Lending *LendingRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Lending.Contract.LendingTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Lending *LendingRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Lending.Contract.LendingTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Lending *LendingCallerRaw) Call(opts *bind.CallOpts, result interface{}, method string, params ...interface{}) error {
	return _Lending.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Lending *LendingTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Lending.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Lending *LendingTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Lending.Contract.contract.Transact(opts, method, params...)
}

// BASES is a free data retrieval call binding the contract method 0x6d1dc42a.
//
// Solidity: function BASES( uint256) constant returns(address)
func (_Lending *LendingCaller) BASES(opts *bind.CallOpts, arg0 *big.Int) (common.Address, error) {
	var (
		ret0 = new(common.Address)
	)
	out := ret0
	err := _Lending.contract.Call(opts, out, "BASES", arg0)
	return *ret0, err
}

// BASES is a free data retrieval call binding the contract method 0x6d1dc42a.
//
// Solidity: function BASES( uint256) constant returns(address)
func (_Lending *LendingSession) BASES(arg0 *big.Int) (common.Address, error) {
	return _Lending.Contract.BASES(&_Lending.CallOpts, arg0)
}

// BASES is a free data retrieval call binding the contract method 0x6d1dc42a.
//
// Solidity: function BASES( uint256) constant returns(address)
func (_Lending *LendingCallerSession) BASES(arg0 *big.Int) (common.Address, error) {
	return _Lending.Contract.BASES(&_Lending.CallOpts, arg0)
}

// COLLATERALS is a free data retrieval call binding the contract method 0x0811f05a.
//
// Solidity: function COLLATERALS( uint256) constant returns(address)
func (_Lending *LendingCaller) COLLATERALS(opts *bind.CallOpts, arg0 *big.Int) (common.Address, error) {
	var (
		ret0 = new(common.Address)
	)
	out := ret0
	err := _Lending.contract.Call(opts, out, "COLLATERALS", arg0)
	return *ret0, err
}

// COLLATERALS is a free data retrieval call binding the contract method 0x0811f05a.
//
// Solidity: function COLLATERALS( uint256) constant returns(address)
func (_Lending *LendingSession) COLLATERALS(arg0 *big.Int) (common.Address, error) {
	return _Lending.Contract.COLLATERALS(&_Lending.CallOpts, arg0)
}

// COLLATERALS is a free data retrieval call binding the contract method 0x0811f05a.
//
// Solidity: function COLLATERALS( uint256) constant returns(address)
func (_Lending *LendingCallerSession) COLLATERALS(arg0 *big.Int) (common.Address, error) {
	return _Lending.Contract.COLLATERALS(&_Lending.CallOpts, arg0)
}

// COLLATERALLIST is a free data retrieval call binding the contract method 0x82250701.
//
// Solidity: function COLLATERAL_LIST( address) constant returns(_depositRate uint256, _liquidationRate uint256, _recallRate uint256)
func (_Lending *LendingCaller) COLLATERALLIST(opts *bind.CallOpts, arg0 common.Address) (struct {
	DepositRate     *big.Int
	LiquidationRate *big.Int
	RecallRate      *big.Int
}, error) {
	ret := new(struct {
		DepositRate     *big.Int
		LiquidationRate *big.Int
		RecallRate      *big.Int
	})
	out := ret
	err := _Lending.contract.Call(opts, out, "COLLATERAL_LIST", arg0)
	return *ret, err
}

// COLLATERALLIST is a free data retrieval call binding the contract method 0x82250701.
//
// Solidity: function COLLATERAL_LIST( address) constant returns(_depositRate uint256, _liquidationRate uint256, _recallRate uint256)
func (_Lending *LendingSession) COLLATERALLIST(arg0 common.Address) (struct {
	DepositRate     *big.Int
	LiquidationRate *big.Int
	RecallRate      *big.Int
}, error) {
	return _Lending.Contract.COLLATERALLIST(&_Lending.CallOpts, arg0)
}

// COLLATERALLIST is a free data retrieval call binding the contract method 0x82250701.
//
// Solidity: function COLLATERAL_LIST( address) constant returns(_depositRate uint256, _liquidationRate uint256, _recallRate uint256)
func (_Lending *LendingCallerSession) COLLATERALLIST(arg0 common.Address) (struct {
	DepositRate     *big.Int
	LiquidationRate *big.Int
	RecallRate      *big.Int
}, error) {
	return _Lending.Contract.COLLATERALLIST(&_Lending.CallOpts, arg0)
}

// CONTRACTOWNER is a free data retrieval call binding the contract method 0xfd301c49.
//
// Solidity: function CONTRACT_OWNER() constant returns(address)
func (_Lending *LendingCaller) CONTRACTOWNER(opts *bind.CallOpts) (common.Address, error) {
	var (
		ret0 = new(common.Address)
	)
	out := ret0
	err := _Lending.contract.Call(opts, out, "CONTRACT_OWNER")
	return *ret0, err
}

// CONTRACTOWNER is a free data retrieval call binding the contract method 0xfd301c49.
//
// Solidity: function CONTRACT_OWNER() constant returns(address)
func (_Lending *LendingSession) CONTRACTOWNER() (common.Address, error) {
	return _Lending.Contract.CONTRACTOWNER(&_Lending.CallOpts)
}

// CONTRACTOWNER is a free data retrieval call binding the contract method 0xfd301c49.
//
// Solidity: function CONTRACT_OWNER() constant returns(address)
func (_Lending *LendingCallerSession) CONTRACTOWNER() (common.Address, error) {
	return _Lending.Contract.CONTRACTOWNER(&_Lending.CallOpts)
}

// ILOCOLLATERALS is a free data retrieval call binding the contract method 0xb8687ec4.
//
// Solidity: function ILO_COLLATERALS( uint256) constant returns(address)
func (_Lending *LendingCaller) ILOCOLLATERALS(opts *bind.CallOpts, arg0 *big.Int) (common.Address, error) {
	var (
		ret0 = new(common.Address)
	)
	out := ret0
	err := _Lending.contract.Call(opts, out, "ILO_COLLATERALS", arg0)
	return *ret0, err
}

// ILOCOLLATERALS is a free data retrieval call binding the contract method 0xb8687ec4.
//
// Solidity: function ILO_COLLATERALS( uint256) constant returns(address)
func (_Lending *LendingSession) ILOCOLLATERALS(arg0 *big.Int) (common.Address, error) {
	return _Lending.Contract.ILOCOLLATERALS(&_Lending.CallOpts, arg0)
}

// ILOCOLLATERALS is a free data retrieval call binding the contract method 0xb8687ec4.
//
// Solidity: function ILO_COLLATERALS( uint256) constant returns(address)
func (_Lending *LendingCallerSession) ILOCOLLATERALS(arg0 *big.Int) (common.Address, error) {
	return _Lending.Contract.ILOCOLLATERALS(&_Lending.CallOpts, arg0)
}

// LENDINGRELAYERLIST is a free data retrieval call binding the contract method 0x0faf292c.
//
// Solidity: function LENDINGRELAYER_LIST( address) constant returns(_tradeFee uint16)
func (_Lending *LendingCaller) LENDINGRELAYERLIST(opts *bind.CallOpts, arg0 common.Address) (uint16, error) {
	var (
		ret0 = new(uint16)
	)
	out := ret0
	err := _Lending.contract.Call(opts, out, "LENDINGRELAYER_LIST", arg0)
	return *ret0, err
}

// LENDINGRELAYERLIST is a free data retrieval call binding the contract method 0x0faf292c.
//
// Solidity: function LENDINGRELAYER_LIST( address) constant returns(_tradeFee uint16)
func (_Lending *LendingSession) LENDINGRELAYERLIST(arg0 common.Address) (uint16, error) {
	return _Lending.Contract.LENDINGRELAYERLIST(&_Lending.CallOpts, arg0)
}

// LENDINGRELAYERLIST is a free data retrieval call binding the contract method 0x0faf292c.
//
// Solidity: function LENDINGRELAYER_LIST( address) constant returns(_tradeFee uint16)
func (_Lending *LendingCallerSession) LENDINGRELAYERLIST(arg0 common.Address) (uint16, error) {
	return _Lending.Contract.LENDINGRELAYERLIST(&_Lending.CallOpts, arg0)
}

// Relayer is a free data retrieval call binding the contract method 0x264949d8.
//
// Solidity: function Relayer() constant returns(address)
func (_Lending *LendingCaller) Relayer(opts *bind.CallOpts) (common.Address, error) {
	var (
		ret0 = new(common.Address)
	)
	out := ret0
	err := _Lending.contract.Call(opts, out, "Relayer")
	return *ret0, err
}

// Relayer is a free data retrieval call binding the contract method 0x264949d8.
//
// Solidity: function Relayer() constant returns(address)
func (_Lending *LendingSession) Relayer() (common.Address, error) {
	return _Lending.Contract.Relayer(&_Lending.CallOpts)
}

// Relayer is a free data retrieval call binding the contract method 0x264949d8.
//
// Solidity: function Relayer() constant returns(address)
func (_Lending *LendingCallerSession) Relayer() (common.Address, error) {
	return _Lending.Contract.Relayer(&_Lending.CallOpts)
}

// TERMS is a free data retrieval call binding the contract method 0x56327f57.
//
// Solidity: function TERMS( uint256) constant returns(uint256)
func (_Lending *LendingCaller) TERMS(opts *bind.CallOpts, arg0 *big.Int) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _Lending.contract.Call(opts, out, "TERMS", arg0)
	return *ret0, err
}

// TERMS is a free data retrieval call binding the contract method 0x56327f57.
//
// Solidity: function TERMS( uint256) constant returns(uint256)
func (_Lending *LendingSession) TERMS(arg0 *big.Int) (*big.Int, error) {
	return _Lending.Contract.TERMS(&_Lending.CallOpts, arg0)
}

// TERMS is a free data retrieval call binding the contract method 0x56327f57.
//
// Solidity: function TERMS( uint256) constant returns(uint256)
func (_Lending *LendingCallerSession) TERMS(arg0 *big.Int) (*big.Int, error) {
	return _Lending.Contract.TERMS(&_Lending.CallOpts, arg0)
}

// TomoXListing is a free data retrieval call binding the contract method 0x29a4ddec.
//
// Solidity: function TomoXListing() constant returns(address)
func (_Lending *LendingCaller) TomoXListing(opts *bind.CallOpts) (common.Address, error) {
	var (
		ret0 = new(common.Address)
	)
	out := ret0
	err := _Lending.contract.Call(opts, out, "TomoXListing")
	return *ret0, err
}

// TomoXListing is a free data retrieval call binding the contract method 0x29a4ddec.
//
// Solidity: function TomoXListing() constant returns(address)
func (_Lending *LendingSession) TomoXListing() (common.Address, error) {
	return _Lending.Contract.TomoXListing(&_Lending.CallOpts)
}

// TomoXListing is a free data retrieval call binding the contract method 0x29a4ddec.
//
// Solidity: function TomoXListing() constant returns(address)
func (_Lending *LendingCallerSession) TomoXListing() (common.Address, error) {
	return _Lending.Contract.TomoXListing(&_Lending.CallOpts)
}

// GetLendingRelayerByCoinbase is a free data retrieval call binding the contract method 0xfe824700.
//
// Solidity: function getLendingRelayerByCoinbase(coinbase address) constant returns(uint16, address[], uint256[], address[])
func (_Lending *LendingCaller) GetLendingRelayerByCoinbase(opts *bind.CallOpts, coinbase common.Address) (uint16, []common.Address, []*big.Int, []common.Address, error) {
	var (
		ret0 = new(uint16)
		ret1 = new([]common.Address)
		ret2 = new([]*big.Int)
		ret3 = new([]common.Address)
	)
	out := &[]interface{}{
		ret0,
		ret1,
		ret2,
		ret3,
	}
	err := _Lending.contract.Call(opts, out, "getLendingRelayerByCoinbase", coinbase)
	return *ret0, *ret1, *ret2, *ret3, err
}

// GetLendingRelayerByCoinbase is a free data retrieval call binding the contract method 0xfe824700.
//
// Solidity: function getLendingRelayerByCoinbase(coinbase address) constant returns(uint16, address[], uint256[], address[])
func (_Lending *LendingSession) GetLendingRelayerByCoinbase(coinbase common.Address) (uint16, []common.Address, []*big.Int, []common.Address, error) {
	return _Lending.Contract.GetLendingRelayerByCoinbase(&_Lending.CallOpts, coinbase)
}

// GetLendingRelayerByCoinbase is a free data retrieval call binding the contract method 0xfe824700.
//
// Solidity: function getLendingRelayerByCoinbase(coinbase address) constant returns(uint16, address[], uint256[], address[])
func (_Lending *LendingCallerSession) GetLendingRelayerByCoinbase(coinbase common.Address) (uint16, []common.Address, []*big.Int, []common.Address, error) {
	return _Lending.Contract.GetLendingRelayerByCoinbase(&_Lending.CallOpts, coinbase)
}

// AddBaseToken is a paid mutator transaction binding the contract method 0x83e280d9.
//
// Solidity: function addBaseToken(token address) returns()
func (_Lending *LendingTransactor) AddBaseToken(opts *bind.TransactOpts, token common.Address) (*types.Transaction, error) {
	return _Lending.contract.Transact(opts, "addBaseToken", token)
}

// AddBaseToken is a paid mutator transaction binding the contract method 0x83e280d9.
//
// Solidity: function addBaseToken(token address) returns()
func (_Lending *LendingSession) AddBaseToken(token common.Address) (*types.Transaction, error) {
	return _Lending.Contract.AddBaseToken(&_Lending.TransactOpts, token)
}

// AddBaseToken is a paid mutator transaction binding the contract method 0x83e280d9.
//
// Solidity: function addBaseToken(token address) returns()
func (_Lending *LendingTransactorSession) AddBaseToken(token common.Address) (*types.Transaction, error) {
	return _Lending.Contract.AddBaseToken(&_Lending.TransactOpts, token)
}

// AddCollateral is a paid mutator transaction binding the contract method 0xe5eecf68.
//
// Solidity: function addCollateral(token address, depositRate uint256, liquidationRate uint256, recallRate uint256) returns()
func (_Lending *LendingTransactor) AddCollateral(opts *bind.TransactOpts, token common.Address, depositRate *big.Int, liquidationRate *big.Int, recallRate *big.Int) (*types.Transaction, error) {
	return _Lending.contract.Transact(opts, "addCollateral", token, depositRate, liquidationRate, recallRate)
}

// AddCollateral is a paid mutator transaction binding the contract method 0xe5eecf68.
//
// Solidity: function addCollateral(token address, depositRate uint256, liquidationRate uint256, recallRate uint256) returns()
func (_Lending *LendingSession) AddCollateral(token common.Address, depositRate *big.Int, liquidationRate *big.Int, recallRate *big.Int) (*types.Transaction, error) {
	return _Lending.Contract.AddCollateral(&_Lending.TransactOpts, token, depositRate, liquidationRate, recallRate)
}

// AddCollateral is a paid mutator transaction binding the contract method 0xe5eecf68.
//
// Solidity: function addCollateral(token address, depositRate uint256, liquidationRate uint256, recallRate uint256) returns()
func (_Lending *LendingTransactorSession) AddCollateral(token common.Address, depositRate *big.Int, liquidationRate *big.Int, recallRate *big.Int) (*types.Transaction, error) {
	return _Lending.Contract.AddCollateral(&_Lending.TransactOpts, token, depositRate, liquidationRate, recallRate)
}

// AddILOCollateral is a paid mutator transaction binding the contract method 0x3b874827.
//
// Solidity: function addILOCollateral(token address, depositRate uint256, liquidationRate uint256, recallRate uint256) returns()
func (_Lending *LendingTransactor) AddILOCollateral(opts *bind.TransactOpts, token common.Address, depositRate *big.Int, liquidationRate *big.Int, recallRate *big.Int) (*types.Transaction, error) {
	return _Lending.contract.Transact(opts, "addILOCollateral", token, depositRate, liquidationRate, recallRate)
}

// AddILOCollateral is a paid mutator transaction binding the contract method 0x3b874827.
//
// Solidity: function addILOCollateral(token address, depositRate uint256, liquidationRate uint256, recallRate uint256) returns()
func (_Lending *LendingSession) AddILOCollateral(token common.Address, depositRate *big.Int, liquidationRate *big.Int, recallRate *big.Int) (*types.Transaction, error) {
	return _Lending.Contract.AddILOCollateral(&_Lending.TransactOpts, token, depositRate, liquidationRate, recallRate)
}

// AddILOCollateral is a paid mutator transaction binding the contract method 0x3b874827.
//
// Solidity: function addILOCollateral(token address, depositRate uint256, liquidationRate uint256, recallRate uint256) returns()
func (_Lending *LendingTransactorSession) AddILOCollateral(token common.Address, depositRate *big.Int, liquidationRate *big.Int, recallRate *big.Int) (*types.Transaction, error) {
	return _Lending.Contract.AddILOCollateral(&_Lending.TransactOpts, token, depositRate, liquidationRate, recallRate)
}

// AddTerm is a paid mutator transaction binding the contract method 0x0c655955.
//
// Solidity: function addTerm(term uint256) returns()
func (_Lending *LendingTransactor) AddTerm(opts *bind.TransactOpts, term *big.Int) (*types.Transaction, error) {
	return _Lending.contract.Transact(opts, "addTerm", term)
}

// AddTerm is a paid mutator transaction binding the contract method 0x0c655955.
//
// Solidity: function addTerm(term uint256) returns()
func (_Lending *LendingSession) AddTerm(term *big.Int) (*types.Transaction, error) {
	return _Lending.Contract.AddTerm(&_Lending.TransactOpts, term)
}

// AddTerm is a paid mutator transaction binding the contract method 0x0c655955.
//
// Solidity: function addTerm(term uint256) returns()
func (_Lending *LendingTransactorSession) AddTerm(term *big.Int) (*types.Transaction, error) {
	return _Lending.Contract.AddTerm(&_Lending.TransactOpts, term)
}

// SetCollateralPrice is a paid mutator transaction binding the contract method 0xacb8cd92.
//
// Solidity: function setCollateralPrice(token address, lendingToken address, price uint256) returns()
func (_Lending *LendingTransactor) SetCollateralPrice(opts *bind.TransactOpts, token common.Address, lendingToken common.Address, price *big.Int) (*types.Transaction, error) {
	return _Lending.contract.Transact(opts, "setCollateralPrice", token, lendingToken, price)
}

// SetCollateralPrice is a paid mutator transaction binding the contract method 0xacb8cd92.
//
// Solidity: function setCollateralPrice(token address, lendingToken address, price uint256) returns()
func (_Lending *LendingSession) SetCollateralPrice(token common.Address, lendingToken common.Address, price *big.Int) (*types.Transaction, error) {
	return _Lending.Contract.SetCollateralPrice(&_Lending.TransactOpts, token, lendingToken, price)
}

// SetCollateralPrice is a paid mutator transaction binding the contract method 0xacb8cd92.
//
// Solidity: function setCollateralPrice(token address, lendingToken address, price uint256) returns()
func (_Lending *LendingTransactorSession) SetCollateralPrice(token common.Address, lendingToken common.Address, price *big.Int) (*types.Transaction, error) {
	return _Lending.Contract.SetCollateralPrice(&_Lending.TransactOpts, token, lendingToken, price)
}

// Update is a paid mutator transaction binding the contract method 0x2ddada4c.
//
// Solidity: function update(coinbase address, tradeFee uint16, baseTokens address[], terms uint256[], collaterals address[]) returns()
func (_Lending *LendingTransactor) Update(opts *bind.TransactOpts, coinbase common.Address, tradeFee uint16, baseTokens []common.Address, terms []*big.Int, collaterals []common.Address) (*types.Transaction, error) {
	return _Lending.contract.Transact(opts, "update", coinbase, tradeFee, baseTokens, terms, collaterals)
}

// Update is a paid mutator transaction binding the contract method 0x2ddada4c.
//
// Solidity: function update(coinbase address, tradeFee uint16, baseTokens address[], terms uint256[], collaterals address[]) returns()
func (_Lending *LendingSession) Update(coinbase common.Address, tradeFee uint16, baseTokens []common.Address, terms []*big.Int, collaterals []common.Address) (*types.Transaction, error) {
	return _Lending.Contract.Update(&_Lending.TransactOpts, coinbase, tradeFee, baseTokens, terms, collaterals)
}

// Update is a paid mutator transaction binding the contract method 0x2ddada4c.
//
// Solidity: function update(coinbase address, tradeFee uint16, baseTokens address[], terms uint256[], collaterals address[]) returns()
func (_Lending *LendingTransactorSession) Update(coinbase common.Address, tradeFee uint16, baseTokens []common.Address, terms []*big.Int, collaterals []common.Address) (*types.Transaction, error) {
	return _Lending.Contract.Update(&_Lending.TransactOpts, coinbase, tradeFee, baseTokens, terms, collaterals)
}
