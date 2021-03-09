// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package main

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

// ERC20AtlasManagerABI is the input ABI used to generate the binding from.
const ERC20AtlasManagerABI = "[{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_wallet\",\"type\":\"address\"}],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"constructor\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"token\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"sender\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"recipient\",\"type\":\"address\"}],\"name\":\"Locked\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"previousOwner\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"newOwner\",\"type\":\"address\"}],\"name\":\"OwnershipTransferred\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address\",\"name\":\"token\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"recipient\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"bytes32\",\"name\":\"receiptId\",\"type\":\"bytes32\"}],\"name\":\"Unlocked\",\"type\":\"event\"},{\"constant\":true,\"inputs\":[],\"name\":\"HYN_ADDRESS\",\"outputs\":[{\"internalType\":\"contractIERC20\",\"name\":\"\",\"type\":\"address\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"isOwner\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"owner\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[],\"name\":\"renounceOwnership\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"internalType\":\"address\",\"name\":\"newOwner\",\"type\":\"address\"}],\"name\":\"transferOwnership\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"name\":\"usedEvents_\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"wallet\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"internalType\":\"address\",\"name\":\"hynTokenAddr\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"recipient\",\"type\":\"address\"}],\"name\":\"lockToken\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"internalType\":\"address\",\"name\":\"hynTokenAddr\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"userAddr\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"recipient\",\"type\":\"address\"}],\"name\":\"lockTokenFor\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"internalType\":\"address\",\"name\":\"hynTokenAddr\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"recipient\",\"type\":\"address\"},{\"internalType\":\"bytes32\",\"name\":\"receiptId\",\"type\":\"bytes32\"}],\"name\":\"unlockToken\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"recipient\",\"type\":\"address\"}],\"name\":\"lockHyn\",\"outputs\":[],\"payable\":true,\"stateMutability\":\"payable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"},{\"internalType\":\"addresspayable\",\"name\":\"recipient\",\"type\":\"address\"},{\"internalType\":\"bytes32\",\"name\":\"receiptId\",\"type\":\"bytes32\"}],\"name\":\"unlockHyn\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"}]"

// ERC20AtlasManagerBin is the compiled bytecode used for deploying new contracts.
var ERC20AtlasManagerBin = "0x608060405234801561001057600080fd5b5060405161125c38038061125c8339818101604052602081101561003357600080fd5b505160006100486001600160e01b036100b716565b600080546001600160a01b0319166001600160a01b0383169081178255604051929350917f8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0908290a350600280546001600160a01b0319166001600160a01b03929092169190911790556100bb565b3390565b611192806100ca6000396000f3fe6080604052600436106100a75760003560e01c80638f32d59b116100645780638f32d59b146101d4578063bccc9fcf146101fd578063be76511f14610227578063d70d485814610253578063f2fde38b14610268578063fe7f61ea1461029b576100a7565b80630296be63146100ac57806309eb2728146100f75780633d2a59961461013a578063521eb27314610179578063715018a6146101aa5780638da5cb5b146101bf575b600080fd5b3480156100b857600080fd5b506100f5600480360360808110156100cf57600080fd5b506001600160a01b038135811691602081013582169160408201359160600135166102e2565b005b34801561010357600080fd5b506100f56004803603606081101561011a57600080fd5b506001600160a01b0381358116916020810135916040909101351661056f565b34801561014657600080fd5b506100f56004803603606081101561015d57600080fd5b508035906001600160a01b03602082013516906040013561077a565b34801561018557600080fd5b5061018e6108df565b604080516001600160a01b039092168252519081900360200190f35b3480156101b657600080fd5b506100f56108ee565b3480156101cb57600080fd5b5061018e610991565b3480156101e057600080fd5b506101e96109a0565b604080519115158252519081900360200190f35b34801561020957600080fd5b506101e96004803603602081101561022057600080fd5b50356109c4565b6100f56004803603604081101561023d57600080fd5b50803590602001356001600160a01b03166109d9565b34801561025f57600080fd5b5061018e610ad1565b34801561027457600080fd5b506100f56004803603602081101561028b57600080fd5b50356001600160a01b0316610ae9565b3480156102a757600080fd5b506100f5600480360360808110156102be57600080fd5b506001600160a01b0381358116916020810135916040820135169060600135610b4e565b6002546001600160a01b03163314610341576040805162461bcd60e51b815260206004820152601b60248201527f41746c61734d616e616765722f6e6f742d617574686f72697a65640000000000604482015290519081900360640190fd5b6001600160a01b0381166103865760405162461bcd60e51b815260040180806020018281038252602881526020018061110c6028913960400191505060405180910390fd5b600082116103db576040805162461bcd60e51b815260206004820152601e60248201527f41746c61734d616e616765722f7a65726f20746f6b656e206c6f636b65640000604482015290519081900360640190fd5b60008490506000816001600160a01b03166370a08231866040518263ffffffff1660e01b815260040180826001600160a01b03166001600160a01b0316815260200191505060206040518083038186803b15801561043857600080fd5b505afa15801561044c573d6000803e3d6000fd5b505050506040513d602081101561046257600080fd5b505190506104816001600160a01b03831686308763ffffffff610c8a16565b6000826001600160a01b03166370a08231876040518263ffffffff1660e01b815260040180826001600160a01b03166001600160a01b0316815260200191505060206040518083038186803b1580156104d957600080fd5b505afa1580156104ed573d6000803e3d6000fd5b505050506040513d602081101561050357600080fd5b505190506000610519838363ffffffff610cea16565b604080518281526001600160a01b0388811660208301528251939450808b1693908816927f4c6ab40ee4cfa212a441d32ee2897945b4a52461284f9369e23fdf8faa6cdd69928290030190a35050505050505050565b6001600160a01b0381166105b45760405162461bcd60e51b815260040180806020018281038252602881526020018061110c6028913960400191505060405180910390fd5b60008211610609576040805162461bcd60e51b815260206004820152601e60248201527f41746c61734d616e616765722f7a65726f20746f6b656e206c6f636b65640000604482015290519081900360640190fd5b604080516370a0823160e01b8152336004820152905184916000916001600160a01b038416916370a08231916024808301926020929190829003018186803b15801561065457600080fd5b505afa158015610668573d6000803e3d6000fd5b505050506040513d602081101561067e57600080fd5b5051905061069d6001600160a01b03831633308763ffffffff610c8a16565b604080516370a0823160e01b815233600482015290516000916001600160a01b038516916370a0823191602480820192602092909190829003018186803b1580156106e757600080fd5b505afa1580156106fb573d6000803e3d6000fd5b505050506040513d602081101561071157600080fd5b505190506000610727838363ffffffff610cea16565b604080518281526001600160a01b03888116602083015282519394503393908816927f4c6ab40ee4cfa212a441d32ee2897945b4a52461284f9369e23fdf8faa6cdd69928290030190a350505050505050565b6002546001600160a01b031633146107d9576040805162461bcd60e51b815260206004820152601b60248201527f41746c61734d616e616765722f6e6f742d617574686f72697a65640000000000604482015290519081900360640190fd5b60008181526001602052604090205460ff16156108275760405162461bcd60e51b815260040180806020018281038252602c8152602001806110e0602c913960400191505060405180910390fd5b6000818152600160208190526040808320805460ff1916909217909155516001600160a01b0384169185156108fc02918691818181858888f19350505050158015610876573d6000803e3d6000fd5b506040805173eeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeee8152602081018590526001600160a01b038416818301526060810183905290517fb24e65d2501e29a3ce014b0cc2283699c081ad27f10d64f036f96912b6f8943e9181900360800190a1505050565b6002546001600160a01b031681565b6108f66109a0565b610947576040805162461bcd60e51b815260206004820181905260248201527f4f776e61626c653a2063616c6c6572206973206e6f7420746865206f776e6572604482015290519081900360640190fd5b600080546040516001600160a01b03909116907f8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0908390a3600080546001600160a01b0319169055565b6000546001600160a01b031690565b600080546001600160a01b03166109b5610d33565b6001600160a01b031614905090565b60016020526000908152604090205460ff1681565b6001600160a01b038116610a1e5760405162461bcd60e51b815260040180806020018281038252602881526020018061110c6028913960400191505060405180910390fd5b813414610a72576040805162461bcd60e51b815260206004820152601c60248201527f48796e4d616e616765722f7a65726f20746f6b656e206c6f636b656400000000604482015290519081900360640190fd5b604080518381526001600160a01b03831660208201528151339273eeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeee927f4c6ab40ee4cfa212a441d32ee2897945b4a52461284f9369e23fdf8faa6cdd69929081900390910190a35050565b73eeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeee81565b610af16109a0565b610b42576040805162461bcd60e51b815260206004820181905260248201527f4f776e61626c653a2063616c6c6572206973206e6f7420746865206f776e6572604482015290519081900360640190fd5b610b4b81610d37565b50565b6002546001600160a01b03163314610bad576040805162461bcd60e51b815260206004820152601b60248201527f41746c61734d616e616765722f6e6f742d617574686f72697a65640000000000604482015290519081900360640190fd5b60008181526001602052604090205460ff1615610bfb5760405162461bcd60e51b815260040180806020018281038252602c8152602001806110e0602c913960400191505060405180910390fd5b6000818152600160208190526040909120805460ff1916909117905583610c326001600160a01b038216848663ffffffff610dd716565b604080516001600160a01b038088168252602082018790528516818301526060810184905290517fb24e65d2501e29a3ce014b0cc2283699c081ad27f10d64f036f96912b6f8943e9181900360800190a15050505050565b604080516001600160a01b0385811660248301528416604482015260648082018490528251808303909101815260849091019091526020810180516001600160e01b03166323b872dd60e01b179052610ce4908590610e2e565b50505050565b6000610d2c83836040518060400160405280601e81526020017f536166654d6174683a207375627472616374696f6e206f766572666c6f770000815250610fe6565b9392505050565b3390565b6001600160a01b038116610d7c5760405162461bcd60e51b81526004018080602001828103825260268152602001806110ba6026913960400191505060405180910390fd5b600080546040516001600160a01b03808516939216917f8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e091a3600080546001600160a01b0319166001600160a01b0392909216919091179055565b604080516001600160a01b038416602482015260448082018490528251808303909101815260649091019091526020810180516001600160e01b031663a9059cbb60e01b179052610e29908490610e2e565b505050565b610e40826001600160a01b031661107d565b610e91576040805162461bcd60e51b815260206004820152601f60248201527f5361666545524332303a2063616c6c20746f206e6f6e2d636f6e747261637400604482015290519081900360640190fd5b60006060836001600160a01b0316836040518082805190602001908083835b60208310610ecf5780518252601f199092019160209182019101610eb0565b6001836020036101000a0380198251168184511680821785525050505050509050019150506000604051808303816000865af19150503d8060008114610f31576040519150601f19603f3d011682016040523d82523d6000602084013e610f36565b606091505b509150915081610f8d576040805162461bcd60e51b815260206004820181905260248201527f5361666545524332303a206c6f772d6c6576656c2063616c6c206661696c6564604482015290519081900360640190fd5b805115610ce457808060200190516020811015610fa957600080fd5b5051610ce45760405162461bcd60e51b815260040180806020018281038252602a815260200180611134602a913960400191505060405180910390fd5b600081848411156110755760405162461bcd60e51b81526004018080602001828103825283818151815260200191508051906020019080838360005b8381101561103a578181015183820152602001611022565b50505050905090810190601f1680156110675780820380516001836020036101000a031916815260200191505b509250505060405180910390fd5b505050900390565b6000813f7fc5d2460186f7233c927e7db2dcc703c0e500b653ca82273b7bfad8045d85a4708181148015906110b157508115155b94935050505056fe4f776e61626c653a206e6577206f776e657220697320746865207a65726f206164647265737341746c61734d616e616765722f546865206275726e206576656e742063616e6e6f742062652072657573656441746c61734d616e616765722f726563697069656e742069732061207a65726f20616464726573735361666545524332303a204552433230206f7065726174696f6e20646964206e6f742073756363656564a265627a7a72315820a540170d9feb0b6c2861295a025b8c107dd989aafaec32429147257a31dd1f1f64736f6c63430005110032"

// DeployERC20AtlasManager deploys a new Ethereum contract, binding an instance of ERC20AtlasManager to it.
func DeployERC20AtlasManager(auth *bind.TransactOpts, backend bind.ContractBackend, _wallet common.Address) (common.Address, *types.Transaction, *ERC20AtlasManager, error) {
	parsed, err := abi.JSON(strings.NewReader(ERC20AtlasManagerABI))
	if err != nil {
		return common.Address{}, nil, nil, err
	}

	address, tx, contract, err := bind.DeployContract(auth, parsed, common.FromHex(ERC20AtlasManagerBin), backend, _wallet)
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	return address, tx, &ERC20AtlasManager{ERC20AtlasManagerCaller: ERC20AtlasManagerCaller{contract: contract}, ERC20AtlasManagerTransactor: ERC20AtlasManagerTransactor{contract: contract}, ERC20AtlasManagerFilterer: ERC20AtlasManagerFilterer{contract: contract}}, nil
}

// ERC20AtlasManager is an auto generated Go binding around an Ethereum contract.
type ERC20AtlasManager struct {
	ERC20AtlasManagerCaller     // Read-only binding to the contract
	ERC20AtlasManagerTransactor // Write-only binding to the contract
	ERC20AtlasManagerFilterer   // Log filterer for contract events
}

// ERC20AtlasManagerCaller is an auto generated read-only Go binding around an Ethereum contract.
type ERC20AtlasManagerCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// ERC20AtlasManagerTransactor is an auto generated write-only Go binding around an Ethereum contract.
type ERC20AtlasManagerTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// ERC20AtlasManagerFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type ERC20AtlasManagerFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// ERC20AtlasManagerSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type ERC20AtlasManagerSession struct {
	Contract     *ERC20AtlasManager // Generic contract binding to set the session for
	CallOpts     bind.CallOpts      // Call options to use throughout this session
	TransactOpts bind.TransactOpts  // Transaction auth options to use throughout this session
}

// ERC20AtlasManagerCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type ERC20AtlasManagerCallerSession struct {
	Contract *ERC20AtlasManagerCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts            // Call options to use throughout this session
}

// ERC20AtlasManagerTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type ERC20AtlasManagerTransactorSession struct {
	Contract     *ERC20AtlasManagerTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts            // Transaction auth options to use throughout this session
}

// ERC20AtlasManagerRaw is an auto generated low-level Go binding around an Ethereum contract.
type ERC20AtlasManagerRaw struct {
	Contract *ERC20AtlasManager // Generic contract binding to access the raw methods on
}

// ERC20AtlasManagerCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type ERC20AtlasManagerCallerRaw struct {
	Contract *ERC20AtlasManagerCaller // Generic read-only contract binding to access the raw methods on
}

// ERC20AtlasManagerTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type ERC20AtlasManagerTransactorRaw struct {
	Contract *ERC20AtlasManagerTransactor // Generic write-only contract binding to access the raw methods on
}

// NewERC20AtlasManager creates a new instance of ERC20AtlasManager, bound to a specific deployed contract.
func NewERC20AtlasManager(address common.Address, backend bind.ContractBackend) (*ERC20AtlasManager, error) {
	contract, err := bindERC20AtlasManager(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &ERC20AtlasManager{ERC20AtlasManagerCaller: ERC20AtlasManagerCaller{contract: contract}, ERC20AtlasManagerTransactor: ERC20AtlasManagerTransactor{contract: contract}, ERC20AtlasManagerFilterer: ERC20AtlasManagerFilterer{contract: contract}}, nil
}

// NewERC20AtlasManagerCaller creates a new read-only instance of ERC20AtlasManager, bound to a specific deployed contract.
func NewERC20AtlasManagerCaller(address common.Address, caller bind.ContractCaller) (*ERC20AtlasManagerCaller, error) {
	contract, err := bindERC20AtlasManager(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &ERC20AtlasManagerCaller{contract: contract}, nil
}

// NewERC20AtlasManagerTransactor creates a new write-only instance of ERC20AtlasManager, bound to a specific deployed contract.
func NewERC20AtlasManagerTransactor(address common.Address, transactor bind.ContractTransactor) (*ERC20AtlasManagerTransactor, error) {
	contract, err := bindERC20AtlasManager(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &ERC20AtlasManagerTransactor{contract: contract}, nil
}

// NewERC20AtlasManagerFilterer creates a new log filterer instance of ERC20AtlasManager, bound to a specific deployed contract.
func NewERC20AtlasManagerFilterer(address common.Address, filterer bind.ContractFilterer) (*ERC20AtlasManagerFilterer, error) {
	contract, err := bindERC20AtlasManager(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &ERC20AtlasManagerFilterer{contract: contract}, nil
}

// bindERC20AtlasManager binds a generic wrapper to an already deployed contract.
func bindERC20AtlasManager(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(ERC20AtlasManagerABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_ERC20AtlasManager *ERC20AtlasManagerRaw) Call(opts *bind.CallOpts, result interface{}, method string, params ...interface{}) error {
	return _ERC20AtlasManager.Contract.ERC20AtlasManagerCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_ERC20AtlasManager *ERC20AtlasManagerRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _ERC20AtlasManager.Contract.ERC20AtlasManagerTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_ERC20AtlasManager *ERC20AtlasManagerRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _ERC20AtlasManager.Contract.ERC20AtlasManagerTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_ERC20AtlasManager *ERC20AtlasManagerCallerRaw) Call(opts *bind.CallOpts, result interface{}, method string, params ...interface{}) error {
	return _ERC20AtlasManager.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_ERC20AtlasManager *ERC20AtlasManagerTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _ERC20AtlasManager.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_ERC20AtlasManager *ERC20AtlasManagerTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _ERC20AtlasManager.Contract.contract.Transact(opts, method, params...)
}

// HYNADDRESS is a free data retrieval call binding the contract method 0xd70d4858.
//
// Solidity: function HYN_ADDRESS() view returns(address)
func (_ERC20AtlasManager *ERC20AtlasManagerCaller) HYNADDRESS(opts *bind.CallOpts) (common.Address, error) {
	var (
		ret0 = new(common.Address)
	)
	out := ret0
	err := _ERC20AtlasManager.contract.Call(opts, out, "HYN_ADDRESS")
	return *ret0, err
}

// HYNADDRESS is a free data retrieval call binding the contract method 0xd70d4858.
//
// Solidity: function HYN_ADDRESS() view returns(address)
func (_ERC20AtlasManager *ERC20AtlasManagerSession) HYNADDRESS() (common.Address, error) {
	return _ERC20AtlasManager.Contract.HYNADDRESS(&_ERC20AtlasManager.CallOpts)
}

// HYNADDRESS is a free data retrieval call binding the contract method 0xd70d4858.
//
// Solidity: function HYN_ADDRESS() view returns(address)
func (_ERC20AtlasManager *ERC20AtlasManagerCallerSession) HYNADDRESS() (common.Address, error) {
	return _ERC20AtlasManager.Contract.HYNADDRESS(&_ERC20AtlasManager.CallOpts)
}

// IsOwner is a free data retrieval call binding the contract method 0x8f32d59b.
//
// Solidity: function isOwner() view returns(bool)
func (_ERC20AtlasManager *ERC20AtlasManagerCaller) IsOwner(opts *bind.CallOpts) (bool, error) {
	var (
		ret0 = new(bool)
	)
	out := ret0
	err := _ERC20AtlasManager.contract.Call(opts, out, "isOwner")
	return *ret0, err
}

// IsOwner is a free data retrieval call binding the contract method 0x8f32d59b.
//
// Solidity: function isOwner() view returns(bool)
func (_ERC20AtlasManager *ERC20AtlasManagerSession) IsOwner() (bool, error) {
	return _ERC20AtlasManager.Contract.IsOwner(&_ERC20AtlasManager.CallOpts)
}

// IsOwner is a free data retrieval call binding the contract method 0x8f32d59b.
//
// Solidity: function isOwner() view returns(bool)
func (_ERC20AtlasManager *ERC20AtlasManagerCallerSession) IsOwner() (bool, error) {
	return _ERC20AtlasManager.Contract.IsOwner(&_ERC20AtlasManager.CallOpts)
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_ERC20AtlasManager *ERC20AtlasManagerCaller) Owner(opts *bind.CallOpts) (common.Address, error) {
	var (
		ret0 = new(common.Address)
	)
	out := ret0
	err := _ERC20AtlasManager.contract.Call(opts, out, "owner")
	return *ret0, err
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_ERC20AtlasManager *ERC20AtlasManagerSession) Owner() (common.Address, error) {
	return _ERC20AtlasManager.Contract.Owner(&_ERC20AtlasManager.CallOpts)
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_ERC20AtlasManager *ERC20AtlasManagerCallerSession) Owner() (common.Address, error) {
	return _ERC20AtlasManager.Contract.Owner(&_ERC20AtlasManager.CallOpts)
}

// UsedEvents is a free data retrieval call binding the contract method 0xbccc9fcf.
//
// Solidity: function usedEvents_(bytes32 ) view returns(bool)
func (_ERC20AtlasManager *ERC20AtlasManagerCaller) UsedEvents(opts *bind.CallOpts, arg0 [32]byte) (bool, error) {
	var (
		ret0 = new(bool)
	)
	out := ret0
	err := _ERC20AtlasManager.contract.Call(opts, out, "usedEvents_", arg0)
	return *ret0, err
}

// UsedEvents is a free data retrieval call binding the contract method 0xbccc9fcf.
//
// Solidity: function usedEvents_(bytes32 ) view returns(bool)
func (_ERC20AtlasManager *ERC20AtlasManagerSession) UsedEvents(arg0 [32]byte) (bool, error) {
	return _ERC20AtlasManager.Contract.UsedEvents(&_ERC20AtlasManager.CallOpts, arg0)
}

// UsedEvents is a free data retrieval call binding the contract method 0xbccc9fcf.
//
// Solidity: function usedEvents_(bytes32 ) view returns(bool)
func (_ERC20AtlasManager *ERC20AtlasManagerCallerSession) UsedEvents(arg0 [32]byte) (bool, error) {
	return _ERC20AtlasManager.Contract.UsedEvents(&_ERC20AtlasManager.CallOpts, arg0)
}

// Wallet is a free data retrieval call binding the contract method 0x521eb273.
//
// Solidity: function wallet() view returns(address)
func (_ERC20AtlasManager *ERC20AtlasManagerCaller) Wallet(opts *bind.CallOpts) (common.Address, error) {
	var (
		ret0 = new(common.Address)
	)
	out := ret0
	err := _ERC20AtlasManager.contract.Call(opts, out, "wallet")
	return *ret0, err
}

// Wallet is a free data retrieval call binding the contract method 0x521eb273.
//
// Solidity: function wallet() view returns(address)
func (_ERC20AtlasManager *ERC20AtlasManagerSession) Wallet() (common.Address, error) {
	return _ERC20AtlasManager.Contract.Wallet(&_ERC20AtlasManager.CallOpts)
}

// Wallet is a free data retrieval call binding the contract method 0x521eb273.
//
// Solidity: function wallet() view returns(address)
func (_ERC20AtlasManager *ERC20AtlasManagerCallerSession) Wallet() (common.Address, error) {
	return _ERC20AtlasManager.Contract.Wallet(&_ERC20AtlasManager.CallOpts)
}

// LockHyn is a paid mutator transaction binding the contract method 0xbe76511f.
//
// Solidity: function lockHyn(uint256 amount, address recipient) payable returns()
func (_ERC20AtlasManager *ERC20AtlasManagerTransactor) LockHyn(opts *bind.TransactOpts, amount *big.Int, recipient common.Address) (*types.Transaction, error) {
	return _ERC20AtlasManager.contract.Transact(opts, "lockHyn", amount, recipient)
}

// LockHyn is a paid mutator transaction binding the contract method 0xbe76511f.
//
// Solidity: function lockHyn(uint256 amount, address recipient) payable returns()
func (_ERC20AtlasManager *ERC20AtlasManagerSession) LockHyn(amount *big.Int, recipient common.Address) (*types.Transaction, error) {
	return _ERC20AtlasManager.Contract.LockHyn(&_ERC20AtlasManager.TransactOpts, amount, recipient)
}

// LockHyn is a paid mutator transaction binding the contract method 0xbe76511f.
//
// Solidity: function lockHyn(uint256 amount, address recipient) payable returns()
func (_ERC20AtlasManager *ERC20AtlasManagerTransactorSession) LockHyn(amount *big.Int, recipient common.Address) (*types.Transaction, error) {
	return _ERC20AtlasManager.Contract.LockHyn(&_ERC20AtlasManager.TransactOpts, amount, recipient)
}

// LockToken is a paid mutator transaction binding the contract method 0x09eb2728.
//
// Solidity: function lockToken(address hynTokenAddr, uint256 amount, address recipient) returns()
func (_ERC20AtlasManager *ERC20AtlasManagerTransactor) LockToken(opts *bind.TransactOpts, hynTokenAddr common.Address, amount *big.Int, recipient common.Address) (*types.Transaction, error) {
	return _ERC20AtlasManager.contract.Transact(opts, "lockToken", hynTokenAddr, amount, recipient)
}

// LockToken is a paid mutator transaction binding the contract method 0x09eb2728.
//
// Solidity: function lockToken(address hynTokenAddr, uint256 amount, address recipient) returns()
func (_ERC20AtlasManager *ERC20AtlasManagerSession) LockToken(hynTokenAddr common.Address, amount *big.Int, recipient common.Address) (*types.Transaction, error) {
	return _ERC20AtlasManager.Contract.LockToken(&_ERC20AtlasManager.TransactOpts, hynTokenAddr, amount, recipient)
}

// LockToken is a paid mutator transaction binding the contract method 0x09eb2728.
//
// Solidity: function lockToken(address hynTokenAddr, uint256 amount, address recipient) returns()
func (_ERC20AtlasManager *ERC20AtlasManagerTransactorSession) LockToken(hynTokenAddr common.Address, amount *big.Int, recipient common.Address) (*types.Transaction, error) {
	return _ERC20AtlasManager.Contract.LockToken(&_ERC20AtlasManager.TransactOpts, hynTokenAddr, amount, recipient)
}

// LockTokenFor is a paid mutator transaction binding the contract method 0x0296be63.
//
// Solidity: function lockTokenFor(address hynTokenAddr, address userAddr, uint256 amount, address recipient) returns()
func (_ERC20AtlasManager *ERC20AtlasManagerTransactor) LockTokenFor(opts *bind.TransactOpts, hynTokenAddr common.Address, userAddr common.Address, amount *big.Int, recipient common.Address) (*types.Transaction, error) {
	return _ERC20AtlasManager.contract.Transact(opts, "lockTokenFor", hynTokenAddr, userAddr, amount, recipient)
}

// LockTokenFor is a paid mutator transaction binding the contract method 0x0296be63.
//
// Solidity: function lockTokenFor(address hynTokenAddr, address userAddr, uint256 amount, address recipient) returns()
func (_ERC20AtlasManager *ERC20AtlasManagerSession) LockTokenFor(hynTokenAddr common.Address, userAddr common.Address, amount *big.Int, recipient common.Address) (*types.Transaction, error) {
	return _ERC20AtlasManager.Contract.LockTokenFor(&_ERC20AtlasManager.TransactOpts, hynTokenAddr, userAddr, amount, recipient)
}

// LockTokenFor is a paid mutator transaction binding the contract method 0x0296be63.
//
// Solidity: function lockTokenFor(address hynTokenAddr, address userAddr, uint256 amount, address recipient) returns()
func (_ERC20AtlasManager *ERC20AtlasManagerTransactorSession) LockTokenFor(hynTokenAddr common.Address, userAddr common.Address, amount *big.Int, recipient common.Address) (*types.Transaction, error) {
	return _ERC20AtlasManager.Contract.LockTokenFor(&_ERC20AtlasManager.TransactOpts, hynTokenAddr, userAddr, amount, recipient)
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_ERC20AtlasManager *ERC20AtlasManagerTransactor) RenounceOwnership(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _ERC20AtlasManager.contract.Transact(opts, "renounceOwnership")
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_ERC20AtlasManager *ERC20AtlasManagerSession) RenounceOwnership() (*types.Transaction, error) {
	return _ERC20AtlasManager.Contract.RenounceOwnership(&_ERC20AtlasManager.TransactOpts)
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_ERC20AtlasManager *ERC20AtlasManagerTransactorSession) RenounceOwnership() (*types.Transaction, error) {
	return _ERC20AtlasManager.Contract.RenounceOwnership(&_ERC20AtlasManager.TransactOpts)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_ERC20AtlasManager *ERC20AtlasManagerTransactor) TransferOwnership(opts *bind.TransactOpts, newOwner common.Address) (*types.Transaction, error) {
	return _ERC20AtlasManager.contract.Transact(opts, "transferOwnership", newOwner)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_ERC20AtlasManager *ERC20AtlasManagerSession) TransferOwnership(newOwner common.Address) (*types.Transaction, error) {
	return _ERC20AtlasManager.Contract.TransferOwnership(&_ERC20AtlasManager.TransactOpts, newOwner)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_ERC20AtlasManager *ERC20AtlasManagerTransactorSession) TransferOwnership(newOwner common.Address) (*types.Transaction, error) {
	return _ERC20AtlasManager.Contract.TransferOwnership(&_ERC20AtlasManager.TransactOpts, newOwner)
}

// UnlockHyn is a paid mutator transaction binding the contract method 0x3d2a5996.
//
// Solidity: function unlockHyn(uint256 amount, address recipient, bytes32 receiptId) returns()
func (_ERC20AtlasManager *ERC20AtlasManagerTransactor) UnlockHyn(opts *bind.TransactOpts, amount *big.Int, recipient common.Address, receiptId [32]byte) (*types.Transaction, error) {
	return _ERC20AtlasManager.contract.Transact(opts, "unlockHyn", amount, recipient, receiptId)
}

// UnlockHyn is a paid mutator transaction binding the contract method 0x3d2a5996.
//
// Solidity: function unlockHyn(uint256 amount, address recipient, bytes32 receiptId) returns()
func (_ERC20AtlasManager *ERC20AtlasManagerSession) UnlockHyn(amount *big.Int, recipient common.Address, receiptId [32]byte) (*types.Transaction, error) {
	return _ERC20AtlasManager.Contract.UnlockHyn(&_ERC20AtlasManager.TransactOpts, amount, recipient, receiptId)
}

// UnlockHyn is a paid mutator transaction binding the contract method 0x3d2a5996.
//
// Solidity: function unlockHyn(uint256 amount, address recipient, bytes32 receiptId) returns()
func (_ERC20AtlasManager *ERC20AtlasManagerTransactorSession) UnlockHyn(amount *big.Int, recipient common.Address, receiptId [32]byte) (*types.Transaction, error) {
	return _ERC20AtlasManager.Contract.UnlockHyn(&_ERC20AtlasManager.TransactOpts, amount, recipient, receiptId)
}

// UnlockToken is a paid mutator transaction binding the contract method 0xfe7f61ea.
//
// Solidity: function unlockToken(address hynTokenAddr, uint256 amount, address recipient, bytes32 receiptId) returns()
func (_ERC20AtlasManager *ERC20AtlasManagerTransactor) UnlockToken(opts *bind.TransactOpts, hynTokenAddr common.Address, amount *big.Int, recipient common.Address, receiptId [32]byte) (*types.Transaction, error) {
	return _ERC20AtlasManager.contract.Transact(opts, "unlockToken", hynTokenAddr, amount, recipient, receiptId)
}

// UnlockToken is a paid mutator transaction binding the contract method 0xfe7f61ea.
//
// Solidity: function unlockToken(address hynTokenAddr, uint256 amount, address recipient, bytes32 receiptId) returns()
func (_ERC20AtlasManager *ERC20AtlasManagerSession) UnlockToken(hynTokenAddr common.Address, amount *big.Int, recipient common.Address, receiptId [32]byte) (*types.Transaction, error) {
	return _ERC20AtlasManager.Contract.UnlockToken(&_ERC20AtlasManager.TransactOpts, hynTokenAddr, amount, recipient, receiptId)
}

// UnlockToken is a paid mutator transaction binding the contract method 0xfe7f61ea.
//
// Solidity: function unlockToken(address hynTokenAddr, uint256 amount, address recipient, bytes32 receiptId) returns()
func (_ERC20AtlasManager *ERC20AtlasManagerTransactorSession) UnlockToken(hynTokenAddr common.Address, amount *big.Int, recipient common.Address, receiptId [32]byte) (*types.Transaction, error) {
	return _ERC20AtlasManager.Contract.UnlockToken(&_ERC20AtlasManager.TransactOpts, hynTokenAddr, amount, recipient, receiptId)
}

// ERC20AtlasManagerLockedIterator is returned from FilterLocked and is used to iterate over the raw logs and unpacked data for Locked events raised by the ERC20AtlasManager contract.
type ERC20AtlasManagerLockedIterator struct {
	Event *ERC20AtlasManagerLocked // Event containing the contract specifics and raw log

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
func (it *ERC20AtlasManagerLockedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ERC20AtlasManagerLocked)
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
		it.Event = new(ERC20AtlasManagerLocked)
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
func (it *ERC20AtlasManagerLockedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ERC20AtlasManagerLockedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ERC20AtlasManagerLocked represents a Locked event raised by the ERC20AtlasManager contract.
type ERC20AtlasManagerLocked struct {
	Token     common.Address
	Sender    common.Address
	Amount    *big.Int
	Recipient common.Address
	Raw       types.Log // Blockchain specific contextual infos
}

// FilterLocked is a free log retrieval operation binding the contract event 0x4c6ab40ee4cfa212a441d32ee2897945b4a52461284f9369e23fdf8faa6cdd69.
//
// Solidity: event Locked(address indexed token, address indexed sender, uint256 amount, address recipient)
func (_ERC20AtlasManager *ERC20AtlasManagerFilterer) FilterLocked(opts *bind.FilterOpts, token []common.Address, sender []common.Address) (*ERC20AtlasManagerLockedIterator, error) {

	var tokenRule []interface{}
	for _, tokenItem := range token {
		tokenRule = append(tokenRule, tokenItem)
	}
	var senderRule []interface{}
	for _, senderItem := range sender {
		senderRule = append(senderRule, senderItem)
	}

	logs, sub, err := _ERC20AtlasManager.contract.FilterLogs(opts, "Locked", tokenRule, senderRule)
	if err != nil {
		return nil, err
	}
	return &ERC20AtlasManagerLockedIterator{contract: _ERC20AtlasManager.contract, event: "Locked", logs: logs, sub: sub}, nil
}

// WatchLocked is a free log subscription operation binding the contract event 0x4c6ab40ee4cfa212a441d32ee2897945b4a52461284f9369e23fdf8faa6cdd69.
//
// Solidity: event Locked(address indexed token, address indexed sender, uint256 amount, address recipient)
func (_ERC20AtlasManager *ERC20AtlasManagerFilterer) WatchLocked(opts *bind.WatchOpts, sink chan<- *ERC20AtlasManagerLocked, token []common.Address, sender []common.Address) (event.Subscription, error) {

	var tokenRule []interface{}
	for _, tokenItem := range token {
		tokenRule = append(tokenRule, tokenItem)
	}
	var senderRule []interface{}
	for _, senderItem := range sender {
		senderRule = append(senderRule, senderItem)
	}

	logs, sub, err := _ERC20AtlasManager.contract.WatchLogs(opts, "Locked", tokenRule, senderRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ERC20AtlasManagerLocked)
				if err := _ERC20AtlasManager.contract.UnpackLog(event, "Locked", log); err != nil {
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

// ParseLocked is a log parse operation binding the contract event 0x4c6ab40ee4cfa212a441d32ee2897945b4a52461284f9369e23fdf8faa6cdd69.
//
// Solidity: event Locked(address indexed token, address indexed sender, uint256 amount, address recipient)
func (_ERC20AtlasManager *ERC20AtlasManagerFilterer) ParseLocked(log types.Log) (*ERC20AtlasManagerLocked, error) {
	event := new(ERC20AtlasManagerLocked)
	if err := _ERC20AtlasManager.contract.UnpackLog(event, "Locked", log); err != nil {
		return nil, err
	}
	return event, nil
}

// ERC20AtlasManagerOwnershipTransferredIterator is returned from FilterOwnershipTransferred and is used to iterate over the raw logs and unpacked data for OwnershipTransferred events raised by the ERC20AtlasManager contract.
type ERC20AtlasManagerOwnershipTransferredIterator struct {
	Event *ERC20AtlasManagerOwnershipTransferred // Event containing the contract specifics and raw log

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
func (it *ERC20AtlasManagerOwnershipTransferredIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ERC20AtlasManagerOwnershipTransferred)
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
		it.Event = new(ERC20AtlasManagerOwnershipTransferred)
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
func (it *ERC20AtlasManagerOwnershipTransferredIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ERC20AtlasManagerOwnershipTransferredIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ERC20AtlasManagerOwnershipTransferred represents a OwnershipTransferred event raised by the ERC20AtlasManager contract.
type ERC20AtlasManagerOwnershipTransferred struct {
	PreviousOwner common.Address
	NewOwner      common.Address
	Raw           types.Log // Blockchain specific contextual infos
}

// FilterOwnershipTransferred is a free log retrieval operation binding the contract event 0x8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0.
//
// Solidity: event OwnershipTransferred(address indexed previousOwner, address indexed newOwner)
func (_ERC20AtlasManager *ERC20AtlasManagerFilterer) FilterOwnershipTransferred(opts *bind.FilterOpts, previousOwner []common.Address, newOwner []common.Address) (*ERC20AtlasManagerOwnershipTransferredIterator, error) {

	var previousOwnerRule []interface{}
	for _, previousOwnerItem := range previousOwner {
		previousOwnerRule = append(previousOwnerRule, previousOwnerItem)
	}
	var newOwnerRule []interface{}
	for _, newOwnerItem := range newOwner {
		newOwnerRule = append(newOwnerRule, newOwnerItem)
	}

	logs, sub, err := _ERC20AtlasManager.contract.FilterLogs(opts, "OwnershipTransferred", previousOwnerRule, newOwnerRule)
	if err != nil {
		return nil, err
	}
	return &ERC20AtlasManagerOwnershipTransferredIterator{contract: _ERC20AtlasManager.contract, event: "OwnershipTransferred", logs: logs, sub: sub}, nil
}

// WatchOwnershipTransferred is a free log subscription operation binding the contract event 0x8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0.
//
// Solidity: event OwnershipTransferred(address indexed previousOwner, address indexed newOwner)
func (_ERC20AtlasManager *ERC20AtlasManagerFilterer) WatchOwnershipTransferred(opts *bind.WatchOpts, sink chan<- *ERC20AtlasManagerOwnershipTransferred, previousOwner []common.Address, newOwner []common.Address) (event.Subscription, error) {

	var previousOwnerRule []interface{}
	for _, previousOwnerItem := range previousOwner {
		previousOwnerRule = append(previousOwnerRule, previousOwnerItem)
	}
	var newOwnerRule []interface{}
	for _, newOwnerItem := range newOwner {
		newOwnerRule = append(newOwnerRule, newOwnerItem)
	}

	logs, sub, err := _ERC20AtlasManager.contract.WatchLogs(opts, "OwnershipTransferred", previousOwnerRule, newOwnerRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ERC20AtlasManagerOwnershipTransferred)
				if err := _ERC20AtlasManager.contract.UnpackLog(event, "OwnershipTransferred", log); err != nil {
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
func (_ERC20AtlasManager *ERC20AtlasManagerFilterer) ParseOwnershipTransferred(log types.Log) (*ERC20AtlasManagerOwnershipTransferred, error) {
	event := new(ERC20AtlasManagerOwnershipTransferred)
	if err := _ERC20AtlasManager.contract.UnpackLog(event, "OwnershipTransferred", log); err != nil {
		return nil, err
	}
	return event, nil
}

// ERC20AtlasManagerUnlockedIterator is returned from FilterUnlocked and is used to iterate over the raw logs and unpacked data for Unlocked events raised by the ERC20AtlasManager contract.
type ERC20AtlasManagerUnlockedIterator struct {
	Event *ERC20AtlasManagerUnlocked // Event containing the contract specifics and raw log

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
func (it *ERC20AtlasManagerUnlockedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ERC20AtlasManagerUnlocked)
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
		it.Event = new(ERC20AtlasManagerUnlocked)
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
func (it *ERC20AtlasManagerUnlockedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ERC20AtlasManagerUnlockedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ERC20AtlasManagerUnlocked represents a Unlocked event raised by the ERC20AtlasManager contract.
type ERC20AtlasManagerUnlocked struct {
	Token     common.Address
	Amount    *big.Int
	Recipient common.Address
	ReceiptId [32]byte
	Raw       types.Log // Blockchain specific contextual infos
}

// FilterUnlocked is a free log retrieval operation binding the contract event 0xb24e65d2501e29a3ce014b0cc2283699c081ad27f10d64f036f96912b6f8943e.
//
// Solidity: event Unlocked(address token, uint256 amount, address recipient, bytes32 receiptId)
func (_ERC20AtlasManager *ERC20AtlasManagerFilterer) FilterUnlocked(opts *bind.FilterOpts) (*ERC20AtlasManagerUnlockedIterator, error) {

	logs, sub, err := _ERC20AtlasManager.contract.FilterLogs(opts, "Unlocked")
	if err != nil {
		return nil, err
	}
	return &ERC20AtlasManagerUnlockedIterator{contract: _ERC20AtlasManager.contract, event: "Unlocked", logs: logs, sub: sub}, nil
}

// WatchUnlocked is a free log subscription operation binding the contract event 0xb24e65d2501e29a3ce014b0cc2283699c081ad27f10d64f036f96912b6f8943e.
//
// Solidity: event Unlocked(address token, uint256 amount, address recipient, bytes32 receiptId)
func (_ERC20AtlasManager *ERC20AtlasManagerFilterer) WatchUnlocked(opts *bind.WatchOpts, sink chan<- *ERC20AtlasManagerUnlocked) (event.Subscription, error) {

	logs, sub, err := _ERC20AtlasManager.contract.WatchLogs(opts, "Unlocked")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ERC20AtlasManagerUnlocked)
				if err := _ERC20AtlasManager.contract.UnpackLog(event, "Unlocked", log); err != nil {
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

// ParseUnlocked is a log parse operation binding the contract event 0xb24e65d2501e29a3ce014b0cc2283699c081ad27f10d64f036f96912b6f8943e.
//
// Solidity: event Unlocked(address token, uint256 amount, address recipient, bytes32 receiptId)
func (_ERC20AtlasManager *ERC20AtlasManagerFilterer) ParseUnlocked(log types.Log) (*ERC20AtlasManagerUnlocked, error) {
	event := new(ERC20AtlasManagerUnlocked)
	if err := _ERC20AtlasManager.contract.UnpackLog(event, "Unlocked", log); err != nil {
		return nil, err
	}
	return event, nil
}
