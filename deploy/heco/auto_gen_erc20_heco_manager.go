// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package heco

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

// ERC20HecoManagerABI is the input ABI used to generate the binding from.
const ERC20HecoManagerABI = "[{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_wallet\",\"type\":\"address\"}],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"constructor\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"token\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"sender\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"recipient\",\"type\":\"address\"}],\"name\":\"Burned\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address\",\"name\":\"token\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"recipient\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"bytes32\",\"name\":\"receiptId\",\"type\":\"bytes32\"}],\"name\":\"Minted\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"previousOwner\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"newOwner\",\"type\":\"address\"}],\"name\":\"OwnershipTransferred\",\"type\":\"event\"},{\"constant\":true,\"inputs\":[],\"name\":\"isOwner\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"name\":\"mappings\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"owner\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[],\"name\":\"renounceOwnership\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"internalType\":\"address\",\"name\":\"newOwner\",\"type\":\"address\"}],\"name\":\"transferOwnership\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"name\":\"usedEvents_\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"wallet\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"internalType\":\"address\",\"name\":\"tokenManager\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"hynTokenAddr\",\"type\":\"address\"},{\"internalType\":\"string\",\"name\":\"name\",\"type\":\"string\"},{\"internalType\":\"string\",\"name\":\"symbol\",\"type\":\"string\"},{\"internalType\":\"uint8\",\"name\":\"decimals\",\"type\":\"uint8\"}],\"name\":\"addToken\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"internalType\":\"address\",\"name\":\"tokenManager\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"hynTokenAddr\",\"type\":\"address\"}],\"name\":\"removeToken\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"internalType\":\"address\",\"name\":\"heco20Token\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"recipient\",\"type\":\"address\"}],\"name\":\"burnToken\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"internalType\":\"address\",\"name\":\"heco20Token\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"recipient\",\"type\":\"address\"},{\"internalType\":\"bytes32\",\"name\":\"receiptId\",\"type\":\"bytes32\"}],\"name\":\"mintToken\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"}]"

// ERC20HecoManagerBin is the compiled bytecode used for deploying new contracts.
var ERC20HecoManagerBin = "0x608060405234801561001057600080fd5b50604051610bcd380380610bcd8339818101604052602081101561003357600080fd5b505160006100486001600160e01b036100b716565b600080546001600160a01b0319166001600160a01b0383169081178255604051929350917f8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0908290a350600380546001600160a01b0319166001600160a01b03929092169190911790556100bb565b3390565b610b03806100ca6000396000f3fe608060405234801561001057600080fd5b50600436106100a95760003560e01c80639735b0e3116100715780639735b0e314610126578063a0e3d1a014610154578063bccc9fcf1461018a578063e8162bda146101a7578063f2fde38b146102f2578063f633be1e14610318576100a9565b806341102619146100ae578063521eb273146100f0578063715018a6146100f85780638da5cb5b146101025780638f32d59b1461010a575b600080fd5b6100d4600480360360208110156100c457600080fd5b50356001600160a01b0316610352565b604080516001600160a01b039092168252519081900360200190f35b6100d461036d565b61010061037c565b005b6100d461041f565b61011261042e565b604080519115158252519081900360200190f35b6101006004803603604081101561013c57600080fd5b506001600160a01b0381358116916020013516610452565b6101006004803603606081101561016a57600080fd5b506001600160a01b03813581169160208101359160409091013516610519565b610112600480360360208110156101a057600080fd5b50356105ce565b610100600480360360a08110156101bd57600080fd5b6001600160a01b0382358116926020810135909116918101906060810160408201356401000000008111156101f157600080fd5b82018360208201111561020357600080fd5b8035906020019184600183028401116401000000008311171561022557600080fd5b91908080601f016020809104026020016040519081016040528093929190818152602001838380828437600092019190915250929594936020810193503591505064010000000081111561027857600080fd5b82018360208201111561028a57600080fd5b803590602001918460018302840111640100000000831117156102ac57600080fd5b91908080601f0160208091040260200160405190810160405280939291908181526020018383808284376000920191909152509295505050903560ff1691506105e39050565b6101006004803603602081101561030857600080fd5b50356001600160a01b03166107db565b6101006004803603608081101561032e57600080fd5b506001600160a01b0381358116916020810135916040820135169060600135610840565b6002602052600090815260409020546001600160a01b031681565b6003546001600160a01b031681565b61038461042e565b6103d5576040805162461bcd60e51b815260206004820181905260248201527f4f776e61626c653a2063616c6c6572206973206e6f7420746865206f776e6572604482015290519081900360640190fd5b600080546040516001600160a01b03909116907f8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0908390a3600080546001600160a01b0319169055565b6000546001600160a01b031690565b600080546001600160a01b03166104436109d9565b6001600160a01b031614905090565b6003546001600160a01b031633146104ae576040805162461bcd60e51b815260206004820152601a602482015279121958dbd3585b9859d95c8bdb9bdd0b585d5d1a1bdc9a5e995960321b604482015290519081900360640190fd5b604080516309dd78f360e11b81526001600160a01b0383811660048301526000602483018190529251908516926313baf1e6926044808201939182900301818387803b1580156104fd57600080fd5b505af1158015610511573d6000803e3d6000fd5b505050505050565b6040805163079cc67960e41b81523360048201526024810184905290516001600160a01b038516916379cc679091604480830192600092919082900301818387803b15801561056757600080fd5b505af115801561057b573d6000803e3d6000fd5b5050604080518581526001600160a01b038581166020830152825133955090881693507f47b4128ad231bee46bdb1f8614450b5ffe6eb96340e87519a0c3a8b91ee714ec929181900390910190a3505050565b60016020526000908152604090205460ff1681565b6003546001600160a01b0316331461063f576040805162461bcd60e51b815260206004820152601a602482015279121958dbd3585b9859d95c8bdb9bdd0b585d5d1a1bdc9a5e995960321b604482015290519081900360640190fd5b6000856001600160a01b0316638e404e9c868686866040518563ffffffff1660e01b815260040180856001600160a01b03166001600160a01b0316815260200180602001806020018460ff1660ff168152602001838103835286818151815260200191508051906020019080838360005b838110156106c85781810151838201526020016106b0565b50505050905090810190601f1680156106f55780820380516001836020036101000a031916815260200191505b50838103825285518152855160209182019187019080838360005b83811015610728578181015183820152602001610710565b50505050905090810190601f1680156107555780820380516001836020036101000a031916815260200191505b509650505050505050602060405180830381600087803b15801561077857600080fd5b505af115801561078c573d6000803e3d6000fd5b505050506040513d60208110156107a257600080fd5b50516001600160a01b03958616600090815260026020526040902080546001600160a01b03191696909116959095179094555050505050565b6107e361042e565b610834576040805162461bcd60e51b815260206004820181905260248201527f4f776e61626c653a2063616c6c6572206973206e6f7420746865206f776e6572604482015290519081900360640190fd5b61083d816109dd565b50565b6003546001600160a01b0316331461089c576040805162461bcd60e51b815260206004820152601a602482015279121958dbd3585b9859d95c8bdb9bdd0b585d5d1a1bdc9a5e995960321b604482015290519081900360640190fd5b60008181526001602052604090205460ff16156108ea5760405162461bcd60e51b815260040180806020018281038252602b815260200180610a7e602b913960400191505060405180910390fd5b6000818152600160208181526040808420805460ff191690931790925581516340c10f1960e01b81526001600160a01b038681166004830152602482018890529251928816936340c10f1993604480840194939192918390030190829087803b15801561095657600080fd5b505af115801561096a573d6000803e3d6000fd5b505050506040513d602081101561098057600080fd5b5050604080516001600160a01b038087168252602082018690528416818301526060810183905290517ffdc0a3af821280dd73e4f6683e16c5afc31a629d07af623f9a4689d6855938869181900360800190a150505050565b3390565b6001600160a01b038116610a225760405162461bcd60e51b8152600401808060200182810382526026815260200180610aa96026913960400191505060405180910390fd5b600080546040516001600160a01b03808516939216917f8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e091a3600080546001600160a01b0319166001600160a01b039290921691909117905556fe4865636f4d616e616765722f546865206c6f636b206576656e742063616e6e6f74206265207265757365644f776e61626c653a206e6577206f776e657220697320746865207a65726f2061646472657373a265627a7a723158204d76999bb218a796879c34342f1451e1022b1a56be8c70569693205e9cfd822d64736f6c63430005110032"

// DeployERC20HecoManager deploys a new Ethereum contract, binding an instance of ERC20HecoManager to it.
func DeployERC20HecoManager(auth *bind.TransactOpts, backend bind.ContractBackend, _wallet common.Address) (common.Address, *types.Transaction, *ERC20HecoManager, error) {
	parsed, err := abi.JSON(strings.NewReader(ERC20HecoManagerABI))
	if err != nil {
		return common.Address{}, nil, nil, err
	}

	address, tx, contract, err := bind.DeployContract(auth, parsed, common.FromHex(ERC20HecoManagerBin), backend, _wallet)
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	return address, tx, &ERC20HecoManager{ERC20HecoManagerCaller: ERC20HecoManagerCaller{contract: contract}, ERC20HecoManagerTransactor: ERC20HecoManagerTransactor{contract: contract}, ERC20HecoManagerFilterer: ERC20HecoManagerFilterer{contract: contract}}, nil
}

// ERC20HecoManager is an auto generated Go binding around an Ethereum contract.
type ERC20HecoManager struct {
	ERC20HecoManagerCaller     // Read-only binding to the contract
	ERC20HecoManagerTransactor // Write-only binding to the contract
	ERC20HecoManagerFilterer   // Log filterer for contract events
}

// ERC20HecoManagerCaller is an auto generated read-only Go binding around an Ethereum contract.
type ERC20HecoManagerCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// ERC20HecoManagerTransactor is an auto generated write-only Go binding around an Ethereum contract.
type ERC20HecoManagerTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// ERC20HecoManagerFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type ERC20HecoManagerFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// ERC20HecoManagerSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type ERC20HecoManagerSession struct {
	Contract     *ERC20HecoManager // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// ERC20HecoManagerCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type ERC20HecoManagerCallerSession struct {
	Contract *ERC20HecoManagerCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts           // Call options to use throughout this session
}

// ERC20HecoManagerTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type ERC20HecoManagerTransactorSession struct {
	Contract     *ERC20HecoManagerTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts           // Transaction auth options to use throughout this session
}

// ERC20HecoManagerRaw is an auto generated low-level Go binding around an Ethereum contract.
type ERC20HecoManagerRaw struct {
	Contract *ERC20HecoManager // Generic contract binding to access the raw methods on
}

// ERC20HecoManagerCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type ERC20HecoManagerCallerRaw struct {
	Contract *ERC20HecoManagerCaller // Generic read-only contract binding to access the raw methods on
}

// ERC20HecoManagerTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type ERC20HecoManagerTransactorRaw struct {
	Contract *ERC20HecoManagerTransactor // Generic write-only contract binding to access the raw methods on
}

// NewERC20HecoManager creates a new instance of ERC20HecoManager, bound to a specific deployed contract.
func NewERC20HecoManager(address common.Address, backend bind.ContractBackend) (*ERC20HecoManager, error) {
	contract, err := bindERC20HecoManager(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &ERC20HecoManager{ERC20HecoManagerCaller: ERC20HecoManagerCaller{contract: contract}, ERC20HecoManagerTransactor: ERC20HecoManagerTransactor{contract: contract}, ERC20HecoManagerFilterer: ERC20HecoManagerFilterer{contract: contract}}, nil
}

// NewERC20HecoManagerCaller creates a new read-only instance of ERC20HecoManager, bound to a specific deployed contract.
func NewERC20HecoManagerCaller(address common.Address, caller bind.ContractCaller) (*ERC20HecoManagerCaller, error) {
	contract, err := bindERC20HecoManager(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &ERC20HecoManagerCaller{contract: contract}, nil
}

// NewERC20HecoManagerTransactor creates a new write-only instance of ERC20HecoManager, bound to a specific deployed contract.
func NewERC20HecoManagerTransactor(address common.Address, transactor bind.ContractTransactor) (*ERC20HecoManagerTransactor, error) {
	contract, err := bindERC20HecoManager(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &ERC20HecoManagerTransactor{contract: contract}, nil
}

// NewERC20HecoManagerFilterer creates a new log filterer instance of ERC20HecoManager, bound to a specific deployed contract.
func NewERC20HecoManagerFilterer(address common.Address, filterer bind.ContractFilterer) (*ERC20HecoManagerFilterer, error) {
	contract, err := bindERC20HecoManager(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &ERC20HecoManagerFilterer{contract: contract}, nil
}

// bindERC20HecoManager binds a generic wrapper to an already deployed contract.
func bindERC20HecoManager(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(ERC20HecoManagerABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_ERC20HecoManager *ERC20HecoManagerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _ERC20HecoManager.Contract.ERC20HecoManagerCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_ERC20HecoManager *ERC20HecoManagerRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _ERC20HecoManager.Contract.ERC20HecoManagerTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_ERC20HecoManager *ERC20HecoManagerRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _ERC20HecoManager.Contract.ERC20HecoManagerTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_ERC20HecoManager *ERC20HecoManagerCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _ERC20HecoManager.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_ERC20HecoManager *ERC20HecoManagerTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _ERC20HecoManager.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_ERC20HecoManager *ERC20HecoManagerTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _ERC20HecoManager.Contract.contract.Transact(opts, method, params...)
}

// IsOwner is a free data retrieval call binding the contract method 0x8f32d59b.
//
// Solidity: function isOwner() view returns(bool)
func (_ERC20HecoManager *ERC20HecoManagerCaller) IsOwner(opts *bind.CallOpts) (bool, error) {
	var out []interface{}
	err := _ERC20HecoManager.contract.Call(opts, &out, "isOwner")

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// IsOwner is a free data retrieval call binding the contract method 0x8f32d59b.
//
// Solidity: function isOwner() view returns(bool)
func (_ERC20HecoManager *ERC20HecoManagerSession) IsOwner() (bool, error) {
	return _ERC20HecoManager.Contract.IsOwner(&_ERC20HecoManager.CallOpts)
}

// IsOwner is a free data retrieval call binding the contract method 0x8f32d59b.
//
// Solidity: function isOwner() view returns(bool)
func (_ERC20HecoManager *ERC20HecoManagerCallerSession) IsOwner() (bool, error) {
	return _ERC20HecoManager.Contract.IsOwner(&_ERC20HecoManager.CallOpts)
}

// Mappings is a free data retrieval call binding the contract method 0x41102619.
//
// Solidity: function mappings(address ) view returns(address)
func (_ERC20HecoManager *ERC20HecoManagerCaller) Mappings(opts *bind.CallOpts, arg0 common.Address) (common.Address, error) {
	var out []interface{}
	err := _ERC20HecoManager.contract.Call(opts, &out, "mappings", arg0)

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Mappings is a free data retrieval call binding the contract method 0x41102619.
//
// Solidity: function mappings(address ) view returns(address)
func (_ERC20HecoManager *ERC20HecoManagerSession) Mappings(arg0 common.Address) (common.Address, error) {
	return _ERC20HecoManager.Contract.Mappings(&_ERC20HecoManager.CallOpts, arg0)
}

// Mappings is a free data retrieval call binding the contract method 0x41102619.
//
// Solidity: function mappings(address ) view returns(address)
func (_ERC20HecoManager *ERC20HecoManagerCallerSession) Mappings(arg0 common.Address) (common.Address, error) {
	return _ERC20HecoManager.Contract.Mappings(&_ERC20HecoManager.CallOpts, arg0)
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_ERC20HecoManager *ERC20HecoManagerCaller) Owner(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _ERC20HecoManager.contract.Call(opts, &out, "owner")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_ERC20HecoManager *ERC20HecoManagerSession) Owner() (common.Address, error) {
	return _ERC20HecoManager.Contract.Owner(&_ERC20HecoManager.CallOpts)
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_ERC20HecoManager *ERC20HecoManagerCallerSession) Owner() (common.Address, error) {
	return _ERC20HecoManager.Contract.Owner(&_ERC20HecoManager.CallOpts)
}

// UsedEvents is a free data retrieval call binding the contract method 0xbccc9fcf.
//
// Solidity: function usedEvents_(bytes32 ) view returns(bool)
func (_ERC20HecoManager *ERC20HecoManagerCaller) UsedEvents(opts *bind.CallOpts, arg0 [32]byte) (bool, error) {
	var out []interface{}
	err := _ERC20HecoManager.contract.Call(opts, &out, "usedEvents_", arg0)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// UsedEvents is a free data retrieval call binding the contract method 0xbccc9fcf.
//
// Solidity: function usedEvents_(bytes32 ) view returns(bool)
func (_ERC20HecoManager *ERC20HecoManagerSession) UsedEvents(arg0 [32]byte) (bool, error) {
	return _ERC20HecoManager.Contract.UsedEvents(&_ERC20HecoManager.CallOpts, arg0)
}

// UsedEvents is a free data retrieval call binding the contract method 0xbccc9fcf.
//
// Solidity: function usedEvents_(bytes32 ) view returns(bool)
func (_ERC20HecoManager *ERC20HecoManagerCallerSession) UsedEvents(arg0 [32]byte) (bool, error) {
	return _ERC20HecoManager.Contract.UsedEvents(&_ERC20HecoManager.CallOpts, arg0)
}

// Wallet is a free data retrieval call binding the contract method 0x521eb273.
//
// Solidity: function wallet() view returns(address)
func (_ERC20HecoManager *ERC20HecoManagerCaller) Wallet(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _ERC20HecoManager.contract.Call(opts, &out, "wallet")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Wallet is a free data retrieval call binding the contract method 0x521eb273.
//
// Solidity: function wallet() view returns(address)
func (_ERC20HecoManager *ERC20HecoManagerSession) Wallet() (common.Address, error) {
	return _ERC20HecoManager.Contract.Wallet(&_ERC20HecoManager.CallOpts)
}

// Wallet is a free data retrieval call binding the contract method 0x521eb273.
//
// Solidity: function wallet() view returns(address)
func (_ERC20HecoManager *ERC20HecoManagerCallerSession) Wallet() (common.Address, error) {
	return _ERC20HecoManager.Contract.Wallet(&_ERC20HecoManager.CallOpts)
}

// AddToken is a paid mutator transaction binding the contract method 0xe8162bda.
//
// Solidity: function addToken(address tokenManager, address hynTokenAddr, string name, string symbol, uint8 decimals) returns()
func (_ERC20HecoManager *ERC20HecoManagerTransactor) AddToken(opts *bind.TransactOpts, tokenManager common.Address, hynTokenAddr common.Address, name string, symbol string, decimals uint8) (*types.Transaction, error) {
	return _ERC20HecoManager.contract.Transact(opts, "addToken", tokenManager, hynTokenAddr, name, symbol, decimals)
}

// AddToken is a paid mutator transaction binding the contract method 0xe8162bda.
//
// Solidity: function addToken(address tokenManager, address hynTokenAddr, string name, string symbol, uint8 decimals) returns()
func (_ERC20HecoManager *ERC20HecoManagerSession) AddToken(tokenManager common.Address, hynTokenAddr common.Address, name string, symbol string, decimals uint8) (*types.Transaction, error) {
	return _ERC20HecoManager.Contract.AddToken(&_ERC20HecoManager.TransactOpts, tokenManager, hynTokenAddr, name, symbol, decimals)
}

// AddToken is a paid mutator transaction binding the contract method 0xe8162bda.
//
// Solidity: function addToken(address tokenManager, address hynTokenAddr, string name, string symbol, uint8 decimals) returns()
func (_ERC20HecoManager *ERC20HecoManagerTransactorSession) AddToken(tokenManager common.Address, hynTokenAddr common.Address, name string, symbol string, decimals uint8) (*types.Transaction, error) {
	return _ERC20HecoManager.Contract.AddToken(&_ERC20HecoManager.TransactOpts, tokenManager, hynTokenAddr, name, symbol, decimals)
}

// BurnToken is a paid mutator transaction binding the contract method 0xa0e3d1a0.
//
// Solidity: function burnToken(address heco20Token, uint256 amount, address recipient) returns()
func (_ERC20HecoManager *ERC20HecoManagerTransactor) BurnToken(opts *bind.TransactOpts, heco20Token common.Address, amount *big.Int, recipient common.Address) (*types.Transaction, error) {
	return _ERC20HecoManager.contract.Transact(opts, "burnToken", heco20Token, amount, recipient)
}

// BurnToken is a paid mutator transaction binding the contract method 0xa0e3d1a0.
//
// Solidity: function burnToken(address heco20Token, uint256 amount, address recipient) returns()
func (_ERC20HecoManager *ERC20HecoManagerSession) BurnToken(heco20Token common.Address, amount *big.Int, recipient common.Address) (*types.Transaction, error) {
	return _ERC20HecoManager.Contract.BurnToken(&_ERC20HecoManager.TransactOpts, heco20Token, amount, recipient)
}

// BurnToken is a paid mutator transaction binding the contract method 0xa0e3d1a0.
//
// Solidity: function burnToken(address heco20Token, uint256 amount, address recipient) returns()
func (_ERC20HecoManager *ERC20HecoManagerTransactorSession) BurnToken(heco20Token common.Address, amount *big.Int, recipient common.Address) (*types.Transaction, error) {
	return _ERC20HecoManager.Contract.BurnToken(&_ERC20HecoManager.TransactOpts, heco20Token, amount, recipient)
}

// MintToken is a paid mutator transaction binding the contract method 0xf633be1e.
//
// Solidity: function mintToken(address heco20Token, uint256 amount, address recipient, bytes32 receiptId) returns()
func (_ERC20HecoManager *ERC20HecoManagerTransactor) MintToken(opts *bind.TransactOpts, heco20Token common.Address, amount *big.Int, recipient common.Address, receiptId [32]byte) (*types.Transaction, error) {
	return _ERC20HecoManager.contract.Transact(opts, "mintToken", heco20Token, amount, recipient, receiptId)
}

// MintToken is a paid mutator transaction binding the contract method 0xf633be1e.
//
// Solidity: function mintToken(address heco20Token, uint256 amount, address recipient, bytes32 receiptId) returns()
func (_ERC20HecoManager *ERC20HecoManagerSession) MintToken(heco20Token common.Address, amount *big.Int, recipient common.Address, receiptId [32]byte) (*types.Transaction, error) {
	return _ERC20HecoManager.Contract.MintToken(&_ERC20HecoManager.TransactOpts, heco20Token, amount, recipient, receiptId)
}

// MintToken is a paid mutator transaction binding the contract method 0xf633be1e.
//
// Solidity: function mintToken(address heco20Token, uint256 amount, address recipient, bytes32 receiptId) returns()
func (_ERC20HecoManager *ERC20HecoManagerTransactorSession) MintToken(heco20Token common.Address, amount *big.Int, recipient common.Address, receiptId [32]byte) (*types.Transaction, error) {
	return _ERC20HecoManager.Contract.MintToken(&_ERC20HecoManager.TransactOpts, heco20Token, amount, recipient, receiptId)
}

// RemoveToken is a paid mutator transaction binding the contract method 0x9735b0e3.
//
// Solidity: function removeToken(address tokenManager, address hynTokenAddr) returns()
func (_ERC20HecoManager *ERC20HecoManagerTransactor) RemoveToken(opts *bind.TransactOpts, tokenManager common.Address, hynTokenAddr common.Address) (*types.Transaction, error) {
	return _ERC20HecoManager.contract.Transact(opts, "removeToken", tokenManager, hynTokenAddr)
}

// RemoveToken is a paid mutator transaction binding the contract method 0x9735b0e3.
//
// Solidity: function removeToken(address tokenManager, address hynTokenAddr) returns()
func (_ERC20HecoManager *ERC20HecoManagerSession) RemoveToken(tokenManager common.Address, hynTokenAddr common.Address) (*types.Transaction, error) {
	return _ERC20HecoManager.Contract.RemoveToken(&_ERC20HecoManager.TransactOpts, tokenManager, hynTokenAddr)
}

// RemoveToken is a paid mutator transaction binding the contract method 0x9735b0e3.
//
// Solidity: function removeToken(address tokenManager, address hynTokenAddr) returns()
func (_ERC20HecoManager *ERC20HecoManagerTransactorSession) RemoveToken(tokenManager common.Address, hynTokenAddr common.Address) (*types.Transaction, error) {
	return _ERC20HecoManager.Contract.RemoveToken(&_ERC20HecoManager.TransactOpts, tokenManager, hynTokenAddr)
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_ERC20HecoManager *ERC20HecoManagerTransactor) RenounceOwnership(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _ERC20HecoManager.contract.Transact(opts, "renounceOwnership")
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_ERC20HecoManager *ERC20HecoManagerSession) RenounceOwnership() (*types.Transaction, error) {
	return _ERC20HecoManager.Contract.RenounceOwnership(&_ERC20HecoManager.TransactOpts)
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_ERC20HecoManager *ERC20HecoManagerTransactorSession) RenounceOwnership() (*types.Transaction, error) {
	return _ERC20HecoManager.Contract.RenounceOwnership(&_ERC20HecoManager.TransactOpts)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_ERC20HecoManager *ERC20HecoManagerTransactor) TransferOwnership(opts *bind.TransactOpts, newOwner common.Address) (*types.Transaction, error) {
	return _ERC20HecoManager.contract.Transact(opts, "transferOwnership", newOwner)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_ERC20HecoManager *ERC20HecoManagerSession) TransferOwnership(newOwner common.Address) (*types.Transaction, error) {
	return _ERC20HecoManager.Contract.TransferOwnership(&_ERC20HecoManager.TransactOpts, newOwner)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_ERC20HecoManager *ERC20HecoManagerTransactorSession) TransferOwnership(newOwner common.Address) (*types.Transaction, error) {
	return _ERC20HecoManager.Contract.TransferOwnership(&_ERC20HecoManager.TransactOpts, newOwner)
}

// ERC20HecoManagerBurnedIterator is returned from FilterBurned and is used to iterate over the raw logs and unpacked data for Burned events raised by the ERC20HecoManager contract.
type ERC20HecoManagerBurnedIterator struct {
	Event *ERC20HecoManagerBurned // Event containing the contract specifics and raw log

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
func (it *ERC20HecoManagerBurnedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ERC20HecoManagerBurned)
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
		it.Event = new(ERC20HecoManagerBurned)
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
func (it *ERC20HecoManagerBurnedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ERC20HecoManagerBurnedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ERC20HecoManagerBurned represents a Burned event raised by the ERC20HecoManager contract.
type ERC20HecoManagerBurned struct {
	Token     common.Address
	Sender    common.Address
	Amount    *big.Int
	Recipient common.Address
	Raw       types.Log // Blockchain specific contextual infos
}

// FilterBurned is a free log retrieval operation binding the contract event 0x47b4128ad231bee46bdb1f8614450b5ffe6eb96340e87519a0c3a8b91ee714ec.
//
// Solidity: event Burned(address indexed token, address indexed sender, uint256 amount, address recipient)
func (_ERC20HecoManager *ERC20HecoManagerFilterer) FilterBurned(opts *bind.FilterOpts, token []common.Address, sender []common.Address) (*ERC20HecoManagerBurnedIterator, error) {

	var tokenRule []interface{}
	for _, tokenItem := range token {
		tokenRule = append(tokenRule, tokenItem)
	}
	var senderRule []interface{}
	for _, senderItem := range sender {
		senderRule = append(senderRule, senderItem)
	}

	logs, sub, err := _ERC20HecoManager.contract.FilterLogs(opts, "Burned", tokenRule, senderRule)
	if err != nil {
		return nil, err
	}
	return &ERC20HecoManagerBurnedIterator{contract: _ERC20HecoManager.contract, event: "Burned", logs: logs, sub: sub}, nil
}

// WatchBurned is a free log subscription operation binding the contract event 0x47b4128ad231bee46bdb1f8614450b5ffe6eb96340e87519a0c3a8b91ee714ec.
//
// Solidity: event Burned(address indexed token, address indexed sender, uint256 amount, address recipient)
func (_ERC20HecoManager *ERC20HecoManagerFilterer) WatchBurned(opts *bind.WatchOpts, sink chan<- *ERC20HecoManagerBurned, token []common.Address, sender []common.Address) (event.Subscription, error) {

	var tokenRule []interface{}
	for _, tokenItem := range token {
		tokenRule = append(tokenRule, tokenItem)
	}
	var senderRule []interface{}
	for _, senderItem := range sender {
		senderRule = append(senderRule, senderItem)
	}

	logs, sub, err := _ERC20HecoManager.contract.WatchLogs(opts, "Burned", tokenRule, senderRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ERC20HecoManagerBurned)
				if err := _ERC20HecoManager.contract.UnpackLog(event, "Burned", log); err != nil {
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

// ParseBurned is a log parse operation binding the contract event 0x47b4128ad231bee46bdb1f8614450b5ffe6eb96340e87519a0c3a8b91ee714ec.
//
// Solidity: event Burned(address indexed token, address indexed sender, uint256 amount, address recipient)
func (_ERC20HecoManager *ERC20HecoManagerFilterer) ParseBurned(log types.Log) (*ERC20HecoManagerBurned, error) {
	event := new(ERC20HecoManagerBurned)
	if err := _ERC20HecoManager.contract.UnpackLog(event, "Burned", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// ERC20HecoManagerMintedIterator is returned from FilterMinted and is used to iterate over the raw logs and unpacked data for Minted events raised by the ERC20HecoManager contract.
type ERC20HecoManagerMintedIterator struct {
	Event *ERC20HecoManagerMinted // Event containing the contract specifics and raw log

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
func (it *ERC20HecoManagerMintedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ERC20HecoManagerMinted)
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
		it.Event = new(ERC20HecoManagerMinted)
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
func (it *ERC20HecoManagerMintedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ERC20HecoManagerMintedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ERC20HecoManagerMinted represents a Minted event raised by the ERC20HecoManager contract.
type ERC20HecoManagerMinted struct {
	Token     common.Address
	Amount    *big.Int
	Recipient common.Address
	ReceiptId [32]byte
	Raw       types.Log // Blockchain specific contextual infos
}

// FilterMinted is a free log retrieval operation binding the contract event 0xfdc0a3af821280dd73e4f6683e16c5afc31a629d07af623f9a4689d685593886.
//
// Solidity: event Minted(address token, uint256 amount, address recipient, bytes32 receiptId)
func (_ERC20HecoManager *ERC20HecoManagerFilterer) FilterMinted(opts *bind.FilterOpts) (*ERC20HecoManagerMintedIterator, error) {

	logs, sub, err := _ERC20HecoManager.contract.FilterLogs(opts, "Minted")
	if err != nil {
		return nil, err
	}
	return &ERC20HecoManagerMintedIterator{contract: _ERC20HecoManager.contract, event: "Minted", logs: logs, sub: sub}, nil
}

// WatchMinted is a free log subscription operation binding the contract event 0xfdc0a3af821280dd73e4f6683e16c5afc31a629d07af623f9a4689d685593886.
//
// Solidity: event Minted(address token, uint256 amount, address recipient, bytes32 receiptId)
func (_ERC20HecoManager *ERC20HecoManagerFilterer) WatchMinted(opts *bind.WatchOpts, sink chan<- *ERC20HecoManagerMinted) (event.Subscription, error) {

	logs, sub, err := _ERC20HecoManager.contract.WatchLogs(opts, "Minted")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ERC20HecoManagerMinted)
				if err := _ERC20HecoManager.contract.UnpackLog(event, "Minted", log); err != nil {
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

// ParseMinted is a log parse operation binding the contract event 0xfdc0a3af821280dd73e4f6683e16c5afc31a629d07af623f9a4689d685593886.
//
// Solidity: event Minted(address token, uint256 amount, address recipient, bytes32 receiptId)
func (_ERC20HecoManager *ERC20HecoManagerFilterer) ParseMinted(log types.Log) (*ERC20HecoManagerMinted, error) {
	event := new(ERC20HecoManagerMinted)
	if err := _ERC20HecoManager.contract.UnpackLog(event, "Minted", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// ERC20HecoManagerOwnershipTransferredIterator is returned from FilterOwnershipTransferred and is used to iterate over the raw logs and unpacked data for OwnershipTransferred events raised by the ERC20HecoManager contract.
type ERC20HecoManagerOwnershipTransferredIterator struct {
	Event *ERC20HecoManagerOwnershipTransferred // Event containing the contract specifics and raw log

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
func (it *ERC20HecoManagerOwnershipTransferredIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ERC20HecoManagerOwnershipTransferred)
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
		it.Event = new(ERC20HecoManagerOwnershipTransferred)
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
func (it *ERC20HecoManagerOwnershipTransferredIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ERC20HecoManagerOwnershipTransferredIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ERC20HecoManagerOwnershipTransferred represents a OwnershipTransferred event raised by the ERC20HecoManager contract.
type ERC20HecoManagerOwnershipTransferred struct {
	PreviousOwner common.Address
	NewOwner      common.Address
	Raw           types.Log // Blockchain specific contextual infos
}

// FilterOwnershipTransferred is a free log retrieval operation binding the contract event 0x8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0.
//
// Solidity: event OwnershipTransferred(address indexed previousOwner, address indexed newOwner)
func (_ERC20HecoManager *ERC20HecoManagerFilterer) FilterOwnershipTransferred(opts *bind.FilterOpts, previousOwner []common.Address, newOwner []common.Address) (*ERC20HecoManagerOwnershipTransferredIterator, error) {

	var previousOwnerRule []interface{}
	for _, previousOwnerItem := range previousOwner {
		previousOwnerRule = append(previousOwnerRule, previousOwnerItem)
	}
	var newOwnerRule []interface{}
	for _, newOwnerItem := range newOwner {
		newOwnerRule = append(newOwnerRule, newOwnerItem)
	}

	logs, sub, err := _ERC20HecoManager.contract.FilterLogs(opts, "OwnershipTransferred", previousOwnerRule, newOwnerRule)
	if err != nil {
		return nil, err
	}
	return &ERC20HecoManagerOwnershipTransferredIterator{contract: _ERC20HecoManager.contract, event: "OwnershipTransferred", logs: logs, sub: sub}, nil
}

// WatchOwnershipTransferred is a free log subscription operation binding the contract event 0x8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0.
//
// Solidity: event OwnershipTransferred(address indexed previousOwner, address indexed newOwner)
func (_ERC20HecoManager *ERC20HecoManagerFilterer) WatchOwnershipTransferred(opts *bind.WatchOpts, sink chan<- *ERC20HecoManagerOwnershipTransferred, previousOwner []common.Address, newOwner []common.Address) (event.Subscription, error) {

	var previousOwnerRule []interface{}
	for _, previousOwnerItem := range previousOwner {
		previousOwnerRule = append(previousOwnerRule, previousOwnerItem)
	}
	var newOwnerRule []interface{}
	for _, newOwnerItem := range newOwner {
		newOwnerRule = append(newOwnerRule, newOwnerItem)
	}

	logs, sub, err := _ERC20HecoManager.contract.WatchLogs(opts, "OwnershipTransferred", previousOwnerRule, newOwnerRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ERC20HecoManagerOwnershipTransferred)
				if err := _ERC20HecoManager.contract.UnpackLog(event, "OwnershipTransferred", log); err != nil {
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

// ParseOwnershipTransferred is a log parse operation binding the contract event 0x8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0.
//
// Solidity: event OwnershipTransferred(address indexed previousOwner, address indexed newOwner)
func (_ERC20HecoManager *ERC20HecoManagerFilterer) ParseOwnershipTransferred(log types.Log) (*ERC20HecoManagerOwnershipTransferred, error) {
	event := new(ERC20HecoManagerOwnershipTransferred)
	if err := _ERC20HecoManager.contract.UnpackLog(event, "OwnershipTransferred", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}
