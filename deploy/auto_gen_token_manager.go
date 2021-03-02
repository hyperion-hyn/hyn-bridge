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

// TokenManagerABI is the input ABI used to generate the binding from.
const TokenManagerABI = "[{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"previousOwner\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"newOwner\",\"type\":\"address\"}],\"name\":\"OwnershipTransferred\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"tokenReq\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"tokenAck\",\"type\":\"address\"}],\"name\":\"TokenMapAck\",\"type\":\"event\"},{\"constant\":true,\"inputs\":[],\"name\":\"isOwner\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"name\":\"mappedTokens\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"owner\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[],\"name\":\"renounceOwnership\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"internalType\":\"address\",\"name\":\"newOwner\",\"type\":\"address\"}],\"name\":\"transferOwnership\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"name\":\"wards\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"internalType\":\"address\",\"name\":\"guy\",\"type\":\"address\"}],\"name\":\"rely\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"internalType\":\"address\",\"name\":\"guy\",\"type\":\"address\"}],\"name\":\"deny\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"internalType\":\"address\",\"name\":\"ethTokenAddr\",\"type\":\"address\"},{\"internalType\":\"string\",\"name\":\"name\",\"type\":\"string\"},{\"internalType\":\"string\",\"name\":\"symbol\",\"type\":\"string\"},{\"internalType\":\"uint8\",\"name\":\"decimals\",\"type\":\"uint8\"}],\"name\":\"addToken\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"internalType\":\"address\",\"name\":\"ethTokenAddr\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"hecoTokenAddr\",\"type\":\"address\"}],\"name\":\"registerToken\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"internalType\":\"address\",\"name\":\"ethTokenAddr\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"supply\",\"type\":\"uint256\"}],\"name\":\"removeToken\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"}]"

// TokenManagerBin is the compiled bytecode used for deploying new contracts.
var TokenManagerBin = "0x608060405260006100176001600160e01b0361006616565b600080546001600160a01b0319166001600160a01b0383169081178255604051929350917f8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0908290a35061006a565b3390565b612416806100796000396000f3fe60806040523480156200001157600080fd5b5060043610620000b85760003560e01c80638e404e9c116200007b5780638e404e9c146200018c5780638f32d59b14620002d75780639c52a7f114620002e1578063bf353dbb146200030a578063e6ede14d1462000345578063f2fde38b146200036e57620000b8565b806313baf1e614620000bd5780634739f7e514620000ee57806365fae35e1462000133578063715018a6146200015c5780638da5cb5b1462000166575b600080fd5b620000ec60048036036040811015620000d557600080fd5b506001600160a01b03813516906020013562000397565b005b6200011f600480360360408110156200010657600080fd5b506001600160a01b03813581169160200135166200053d565b604080519115158252519081900360200190f35b620000ec600480360360208110156200014b57600080fd5b50356001600160a01b031662000675565b620000ec620006dd565b6200017062000772565b604080516001600160a01b039092168252519081900360200190f35b6200017060048036036080811015620001a457600080fd5b6001600160a01b038235169190810190604081016020820135640100000000811115620001d057600080fd5b820183602082011115620001e357600080fd5b803590602001918460018302840111640100000000831117156200020657600080fd5b91908080601f01602080910402602001604051908101604052809392919081815260200183838082843760009201919091525092959493602081019350359150506401000000008111156200025a57600080fd5b8201836020820111156200026d57600080fd5b803590602001918460018302840111640100000000831117156200029057600080fd5b91908080601f0160208091040260200160405190810160405280939291908181526020018383808284376000920191909152509295505050903560ff169150620007819050565b6200011f62000a71565b620000ec60048036036020811015620002f957600080fd5b50356001600160a01b031662000a97565b62000333600480360360208110156200032257600080fd5b50356001600160a01b031662000b58565b60408051918252519081900360200190f35b62000170600480360360208110156200035d57600080fd5b50356001600160a01b031662000b6a565b620000ec600480360360208110156200038657600080fd5b50356001600160a01b031662000b85565b33600090815260026020526040902054600114620003fc576040805162461bcd60e51b815260206004820152601b60248201527f546f6b656e4d616e616765722f6e6f742d617574686f72697a65640000000000604482015290519081900360640190fd5b6001600160a01b0382811660009081526001602052604090205416620004545760405162461bcd60e51b815260040180806020018281038252602d81526020018062002346602d913960400191505060405180910390fd5b6001600160a01b038083166000908152600160209081526040918290205482516318160ddd60e01b81529251931692849284926318160ddd92600480840193829003018186803b158015620004a857600080fd5b505afa158015620004bd573d6000803e3d6000fd5b505050506040513d6020811015620004d457600080fd5b505114620005145760405162461bcd60e51b8152600401808060200182810382526027815260200180620023bb6027913960400191505060405180910390fd5b50506001600160a01b0316600090815260016020526040902080546001600160a01b0319169055565b33600090815260026020526040812054600114620005a2576040805162461bcd60e51b815260206004820152601b60248201527f546f6b656e4d616e616765722f6e6f742d617574686f72697a65640000000000604482015290519081900360640190fd5b6001600160a01b038316620005e95760405162461bcd60e51b815260040180806020018281038252602b81526020018062002288602b913960400191505060405180910390fd5b6001600160a01b038381166000908152600160205260409020541615620006425760405162461bcd60e51b8152600401808060200182810382526028815260200180620023936028913960400191505060405180910390fd5b6001600160a01b03928316600090815260016020526040902080546001600160a01b031916929093169190911790915590565b6200067f62000a71565b620006c0576040805162461bcd60e51b8152602060048201819052602482015260008051602062002373833981519152604482015290519081900360640190fd5b6001600160a01b0316600090815260026020526040902060019055565b620006e762000a71565b62000728576040805162461bcd60e51b8152602060048201819052602482015260008051602062002373833981519152604482015290519081900360640190fd5b600080546040516001600160a01b03909116907f8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0908390a3600080546001600160a01b0319169055565b6000546001600160a01b031690565b33600090815260026020526040812054600114620007e6576040805162461bcd60e51b815260206004820152601b60248201527f546f6b656e4d616e616765722f6e6f742d617574686f72697a65640000000000604482015290519081900360640190fd5b6001600160a01b0385166200082d5760405162461bcd60e51b8152600401808060200182810382526027815260200180620022d96027913960400191505060405180910390fd5b6001600160a01b038581166000908152600160205260409020541615620008865760405162461bcd60e51b8152600401808060200182810382526024815260200180620023006024913960400191505060405180910390fd5b6000858585856040516200089a9062000c84565b6001600160a01b038516815260ff82166060820152608060208083018281528651928401929092528551604084019160a08501919088019080838360005b83811015620008f2578181015183820152602001620008d8565b50505050905090810190601f168015620009205780820380516001836020036101000a031916815260200191505b50838103825285518152855160209182019187019080838360005b83811015620009555781810151838201526020016200093b565b50505050905090810190601f168015620009835780820380516001836020036101000a031916815260200191505b509650505050505050604051809103906000f080158015620009a9573d6000803e3d6000fd5b506001600160a01b0387811660009081526001602052604080822080546001600160a01b03191693851693841790558051634c1d96ab60e11b81523360048201529051939450849363983b2d569260248084019391929182900301818387803b15801562000a1657600080fd5b505af115801562000a2b573d6000803e3d6000fd5b50506040516001600160a01b0380851693508a1691507f78591f651c27eef63481b7fd779e44c2426cbe82c7050fe9bc90b6707efbc3d390600090a39695505050505050565b600080546001600160a01b031662000a8862000bde565b6001600160a01b031614905090565b62000aa162000a71565b62000ae2576040805162461bcd60e51b8152602060048201819052602482015260008051602062002373833981519152604482015290519081900360640190fd5b62000aec62000772565b6001600160a01b0316816001600160a01b0316141562000b3e5760405162461bcd60e51b8152600401808060200182810382526022815260200180620023246022913960400191505060405180910390fd5b6001600160a01b0316600090815260026020526040812055565b60026020526000908152604090205481565b6001602052600090815260409020546001600160a01b031681565b62000b8f62000a71565b62000bd0576040805162461bcd60e51b8152602060048201819052602482015260008051602062002373833981519152604482015290519081900360640190fd5b62000bdb8162000be2565b50565b3390565b6001600160a01b03811662000c295760405162461bcd60e51b8152600401808060200182810382526026815260200180620022b36026913960400191505060405180910390fd5b600080546040516001600160a01b03808516939216917f8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e091a3600080546001600160a01b0319166001600160a01b0392909216919091179055565b6115f58062000c938339019056fe60806040523480156200001157600080fd5b50604051620015f5380380620015f5833981810160405260808110156200003757600080fd5b8151602083018051604051929492938301929190846401000000008211156200005f57600080fd5b9083019060208201858111156200007557600080fd5b82516401000000008111828201881017156200009057600080fd5b82525081516020918201929091019080838360005b83811015620000bf578181015183820152602001620000a5565b50505050905090810190601f168015620000ed5780820380516001836020036101000a031916815260200191505b50604052602001805160405193929190846401000000008211156200011157600080fd5b9083019060208201858111156200012757600080fd5b82516401000000008111828201881017156200014257600080fd5b82525081516020918201929091019080838360005b838110156200017157818101518382015260200162000157565b50505050905090810190601f1680156200019f5780820380516001836020036101000a031916815260200191505b50604052602090810151855190935085925084918491620001c7916003919086019062000389565b508151620001dd90600490602085019062000389565b506005805460ff191660ff9290921691909117905550620002129050620002036200023c565b6001600160e01b036200024116565b5050600780546001600160a01b0319166001600160a01b039390931692909217909155506200042b565b335b90565b6200025c8160066200029360201b62000ea01790919060201c565b6040516001600160a01b038216907f6ae172837ea30b801fbfcdd4108aa1d5bf8ff775444fd70256b44e6bf3dfc3f690600090a250565b620002a882826001600160e01b036200032016565b15620002fb576040805162461bcd60e51b815260206004820152601f60248201527f526f6c65733a206163636f756e7420616c72656164792068617320726f6c6500604482015290519081900360640190fd5b6001600160a01b0316600090815260209190915260409020805460ff19166001179055565b60006001600160a01b038216620003695760405162461bcd60e51b8152600401808060200182810382526022815260200180620015d36022913960400191505060405180910390fd5b506001600160a01b03166000908152602091909152604090205460ff1690565b828054600181600116156101000203166002900490600052602060002090601f016020900481019282601f10620003cc57805160ff1916838001178555620003fc565b82800160010185558215620003fc579182015b82811115620003fc578251825591602001919060010190620003df565b506200040a9291506200040e565b5090565b6200023e91905b808211156200040a576000815560010162000415565b611198806200043b6000396000f3fe608060405234801561001057600080fd5b50600436106101165760003560e01c806370a08231116100a257806398650275116100715780639865027514610361578063a457c2d714610369578063a9059cbb14610395578063aa271e1a146103c1578063dd62ed3e146103e757610116565b806370a08231146102e157806379cc67901461030757806395d89b4114610333578063983b2d561461033b57610116565b806323b872dd116100e957806323b872dd14610216578063313ce5671461024c578063395093511461026a57806340c10f191461029657806342966c68146102c257610116565b806306fdde031461011b578063095ea7b31461019857806311c9adc8146101d857806318160ddd146101fc575b600080fd5b610123610415565b6040805160208082528351818301528351919283929083019185019080838360005b8381101561015d578181015183820152602001610145565b50505050905090810190601f16801561018a5780820380516001836020036101000a031916815260200191505b509250505060405180910390f35b6101c4600480360360408110156101ae57600080fd5b506001600160a01b0381351690602001356104ab565b604080519115158252519081900360200190f35b6101e06104c8565b604080516001600160a01b039092168252519081900360200190f35b6102046104d7565b60408051918252519081900360200190f35b6101c46004803603606081101561022c57600080fd5b506001600160a01b038135811691602081013590911690604001356104dd565b61025461056a565b6040805160ff9092168252519081900360200190f35b6101c46004803603604081101561028057600080fd5b506001600160a01b038135169060200135610573565b6101c4600480360360408110156102ac57600080fd5b506001600160a01b0381351690602001356105c7565b6102df600480360360208110156102d857600080fd5b503561061e565b005b610204600480360360208110156102f757600080fd5b50356001600160a01b0316610632565b6102df6004803603604081101561031d57600080fd5b506001600160a01b03813516906020013561064d565b61012361065b565b6102df6004803603602081101561035157600080fd5b50356001600160a01b03166106bc565b6102df61070b565b6101c46004803603604081101561037f57600080fd5b506001600160a01b03813516906020013561071d565b6101c4600480360360408110156103ab57600080fd5b506001600160a01b03813516906020013561078b565b6101c4600480360360208110156103d757600080fd5b50356001600160a01b031661079f565b610204600480360360408110156103fd57600080fd5b506001600160a01b03813581169160200135166107b8565b60038054604080516020601f60026000196101006001881615020190951694909404938401819004810282018101909252828152606093909290918301828280156104a15780601f10610476576101008083540402835291602001916104a1565b820191906000526020600020905b81548152906001019060200180831161048457829003601f168201915b5050505050905090565b60006104bf6104b86107e3565b84846107e7565b50600192915050565b6007546001600160a01b031681565b60025490565b60006104ea8484846108d3565b610560846104f66107e3565b61055b85604051806060016040528060288152602001611067602891396001600160a01b038a166000908152600160205260408120906105346107e3565b6001600160a01b03168152602081019190915260400160002054919063ffffffff610a2f16565b6107e7565b5060019392505050565b60055460ff1690565b60006104bf6105806107e3565b8461055b85600160006105916107e3565b6001600160a01b03908116825260208083019390935260409182016000908120918c16815292529020549063ffffffff610ac616565b60006105d96105d46107e3565b61079f565b6106145760405162461bcd60e51b81526004018080602001828103825260308152602001806110166030913960400191505060405180910390fd5b6104bf8383610b27565b61062f6106296107e3565b82610c17565b50565b6001600160a01b031660009081526020819052604090205490565b6106578282610d13565b5050565b60048054604080516020601f60026000196101006001881615020190951694909404938401819004810282018101909252828152606093909290918301828280156104a15780601f10610476576101008083540402835291602001916104a1565b6106c76105d46107e3565b6107025760405162461bcd60e51b81526004018080602001828103825260308152602001806110166030913960400191505060405180910390fd5b61062f81610d67565b61071b6107166107e3565b610daf565b565b60006104bf61072a6107e3565b8461055b8560405180606001604052806025815260200161113f60259139600160006107546107e3565b6001600160a01b03908116825260208083019390935260409182016000908120918d1681529252902054919063ffffffff610a2f16565b60006104bf6107986107e3565b84846108d3565b60006107b260068363ffffffff610df716565b92915050565b6001600160a01b03918216600090815260016020908152604080832093909416825291909152205490565b3390565b6001600160a01b03831661082c5760405162461bcd60e51b815260040180806020018281038252602481526020018061111b6024913960400191505060405180910390fd5b6001600160a01b0382166108715760405162461bcd60e51b8152600401808060200182810382526022815260200180610fce6022913960400191505060405180910390fd5b6001600160a01b03808416600081815260016020908152604080832094871680845294825291829020859055815185815291517f8c5be1e5ebec7d5bd14f71427d1e84f3dd0314c0f7b2291e5b200ac8c7c3b9259281900390910190a3505050565b6001600160a01b0383166109185760405162461bcd60e51b81526004018080602001828103825260258152602001806110f66025913960400191505060405180910390fd5b6001600160a01b03821661095d5760405162461bcd60e51b8152600401808060200182810382526023815260200180610f896023913960400191505060405180910390fd5b6109a081604051806060016040528060268152602001610ff0602691396001600160a01b038616600090815260208190526040902054919063ffffffff610a2f16565b6001600160a01b0380851660009081526020819052604080822093909355908416815220546109d5908263ffffffff610ac616565b6001600160a01b038084166000818152602081815260409182902094909455805185815290519193928716927fddf252ad1be2c89b69c2b068fc378daa952ba7f163c4a11628f55a4df523b3ef92918290030190a3505050565b60008184841115610abe5760405162461bcd60e51b81526004018080602001828103825283818151815260200191508051906020019080838360005b83811015610a83578181015183820152602001610a6b565b50505050905090810190601f168015610ab05780820380516001836020036101000a031916815260200191505b509250505060405180910390fd5b505050900390565b600082820183811015610b20576040805162461bcd60e51b815260206004820152601b60248201527f536166654d6174683a206164646974696f6e206f766572666c6f770000000000604482015290519081900360640190fd5b9392505050565b6001600160a01b038216610b82576040805162461bcd60e51b815260206004820152601f60248201527f45524332303a206d696e7420746f20746865207a65726f206164647265737300604482015290519081900360640190fd5b600254610b95908263ffffffff610ac616565b6002556001600160a01b038216600090815260208190526040902054610bc1908263ffffffff610ac616565b6001600160a01b0383166000818152602081815260408083209490945583518581529351929391927fddf252ad1be2c89b69c2b068fc378daa952ba7f163c4a11628f55a4df523b3ef9281900390910190a35050565b6001600160a01b038216610c5c5760405162461bcd60e51b81526004018080602001828103825260218152602001806110d56021913960400191505060405180910390fd5b610c9f81604051806060016040528060228152602001610fac602291396001600160a01b038516600090815260208190526040902054919063ffffffff610a2f16565b6001600160a01b038316600090815260208190526040902055600254610ccb908263ffffffff610e5e16565b6002556040805182815290516000916001600160a01b038516917fddf252ad1be2c89b69c2b068fc378daa952ba7f163c4a11628f55a4df523b3ef9181900360200190a35050565b610d1d8282610c17565b61065782610d296107e3565b61055b846040518060600160405280602481526020016110b1602491396001600160a01b0388166000908152600160205260408120906105346107e3565b610d7860068263ffffffff610ea016565b6040516001600160a01b038216907f6ae172837ea30b801fbfcdd4108aa1d5bf8ff775444fd70256b44e6bf3dfc3f690600090a250565b610dc060068263ffffffff610f2116565b6040516001600160a01b038216907fe94479a9f7e1952cc78f2d6baab678adc1b772d936c6583def489e524cb6669290600090a250565b60006001600160a01b038216610e3e5760405162461bcd60e51b815260040180806020018281038252602281526020018061108f6022913960400191505060405180910390fd5b506001600160a01b03166000908152602091909152604090205460ff1690565b6000610b2083836040518060400160405280601e81526020017f536166654d6174683a207375627472616374696f6e206f766572666c6f770000815250610a2f565b610eaa8282610df7565b15610efc576040805162461bcd60e51b815260206004820152601f60248201527f526f6c65733a206163636f756e7420616c72656164792068617320726f6c6500604482015290519081900360640190fd5b6001600160a01b0316600090815260209190915260409020805460ff19166001179055565b610f2b8282610df7565b610f665760405162461bcd60e51b81526004018080602001828103825260218152602001806110466021913960400191505060405180910390fd5b6001600160a01b0316600090815260209190915260409020805460ff1916905556fe45524332303a207472616e7366657220746f20746865207a65726f206164647265737345524332303a206275726e20616d6f756e7420657863656564732062616c616e636545524332303a20617070726f766520746f20746865207a65726f206164647265737345524332303a207472616e7366657220616d6f756e7420657863656564732062616c616e63654d696e746572526f6c653a2063616c6c657220646f6573206e6f74206861766520746865204d696e74657220726f6c65526f6c65733a206163636f756e7420646f6573206e6f74206861766520726f6c6545524332303a207472616e7366657220616d6f756e74206578636565647320616c6c6f77616e6365526f6c65733a206163636f756e7420697320746865207a65726f206164647265737345524332303a206275726e20616d6f756e74206578636565647320616c6c6f77616e636545524332303a206275726e2066726f6d20746865207a65726f206164647265737345524332303a207472616e736665722066726f6d20746865207a65726f206164647265737345524332303a20617070726f76652066726f6d20746865207a65726f206164647265737345524332303a2064656372656173656420616c6c6f77616e63652062656c6f77207a65726fa265627a7a7231582039322b3502377ec18d10b34596ef13427799abeac7c04923e908b3613ba48ed864736f6c63430005110032526f6c65733a206163636f756e7420697320746865207a65726f2061646472657373546f6b656e4d616e616765722f657468546f6b656e416464722069732061207a65726f20616464726573734f776e61626c653a206e6577206f776e657220697320746865207a65726f2061646472657373546f6b656e4d616e616765722f657468546f6b656e2069732061207a65726f2061646472657373546f6b656e4d616e616765722f657468546f6b656e20616c7265616479206d6170706564546f6b656e4d616e616765722f63616e6e6f742064656e7920746865206f776e6572546f6b656e4d616e616765722f657468546f6b656e206d617070696e6720646f6573206e6f74206578697374734f776e61626c653a2063616c6c6572206973206e6f7420746865206f776e6572546f6b656e4d616e616765722f657468546f6b656e4164647220616c7265616479206d6170706564546f6b656e4d616e616765722f72656d6f766520686173206e6f6e2d7a65726f20737570706c79a265627a7a7231582059ca995c79d1976bcf2f03b33068a665222190ff47179dc873f95649d8e99b5b64736f6c63430005110032"

// DeployTokenManager deploys a new Ethereum contract, binding an instance of TokenManager to it.
func DeployTokenManager(auth *bind.TransactOpts, backend bind.ContractBackend) (common.Address, *types.Transaction, *TokenManager, error) {
	parsed, err := abi.JSON(strings.NewReader(TokenManagerABI))
	if err != nil {
		return common.Address{}, nil, nil, err
	}

	address, tx, contract, err := bind.DeployContract(auth, parsed, common.FromHex(TokenManagerBin), backend)
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	return address, tx, &TokenManager{TokenManagerCaller: TokenManagerCaller{contract: contract}, TokenManagerTransactor: TokenManagerTransactor{contract: contract}, TokenManagerFilterer: TokenManagerFilterer{contract: contract}}, nil
}

// TokenManager is an auto generated Go binding around an Ethereum contract.
type TokenManager struct {
	TokenManagerCaller     // Read-only binding to the contract
	TokenManagerTransactor // Write-only binding to the contract
	TokenManagerFilterer   // Log filterer for contract events
}

// TokenManagerCaller is an auto generated read-only Go binding around an Ethereum contract.
type TokenManagerCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// TokenManagerTransactor is an auto generated write-only Go binding around an Ethereum contract.
type TokenManagerTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// TokenManagerFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type TokenManagerFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// TokenManagerSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type TokenManagerSession struct {
	Contract     *TokenManager     // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// TokenManagerCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type TokenManagerCallerSession struct {
	Contract *TokenManagerCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts       // Call options to use throughout this session
}

// TokenManagerTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type TokenManagerTransactorSession struct {
	Contract     *TokenManagerTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts       // Transaction auth options to use throughout this session
}

// TokenManagerRaw is an auto generated low-level Go binding around an Ethereum contract.
type TokenManagerRaw struct {
	Contract *TokenManager // Generic contract binding to access the raw methods on
}

// TokenManagerCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type TokenManagerCallerRaw struct {
	Contract *TokenManagerCaller // Generic read-only contract binding to access the raw methods on
}

// TokenManagerTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type TokenManagerTransactorRaw struct {
	Contract *TokenManagerTransactor // Generic write-only contract binding to access the raw methods on
}

// NewTokenManager creates a new instance of TokenManager, bound to a specific deployed contract.
func NewTokenManager(address common.Address, backend bind.ContractBackend) (*TokenManager, error) {
	contract, err := bindTokenManager(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &TokenManager{TokenManagerCaller: TokenManagerCaller{contract: contract}, TokenManagerTransactor: TokenManagerTransactor{contract: contract}, TokenManagerFilterer: TokenManagerFilterer{contract: contract}}, nil
}

// NewTokenManagerCaller creates a new read-only instance of TokenManager, bound to a specific deployed contract.
func NewTokenManagerCaller(address common.Address, caller bind.ContractCaller) (*TokenManagerCaller, error) {
	contract, err := bindTokenManager(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &TokenManagerCaller{contract: contract}, nil
}

// NewTokenManagerTransactor creates a new write-only instance of TokenManager, bound to a specific deployed contract.
func NewTokenManagerTransactor(address common.Address, transactor bind.ContractTransactor) (*TokenManagerTransactor, error) {
	contract, err := bindTokenManager(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &TokenManagerTransactor{contract: contract}, nil
}

// NewTokenManagerFilterer creates a new log filterer instance of TokenManager, bound to a specific deployed contract.
func NewTokenManagerFilterer(address common.Address, filterer bind.ContractFilterer) (*TokenManagerFilterer, error) {
	contract, err := bindTokenManager(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &TokenManagerFilterer{contract: contract}, nil
}

// bindTokenManager binds a generic wrapper to an already deployed contract.
func bindTokenManager(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(TokenManagerABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_TokenManager *TokenManagerRaw) Call(opts *bind.CallOpts, result interface{}, method string, params ...interface{}) error {
	return _TokenManager.Contract.TokenManagerCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_TokenManager *TokenManagerRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _TokenManager.Contract.TokenManagerTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_TokenManager *TokenManagerRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _TokenManager.Contract.TokenManagerTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_TokenManager *TokenManagerCallerRaw) Call(opts *bind.CallOpts, result interface{}, method string, params ...interface{}) error {
	return _TokenManager.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_TokenManager *TokenManagerTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _TokenManager.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_TokenManager *TokenManagerTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _TokenManager.Contract.contract.Transact(opts, method, params...)
}

// IsOwner is a free data retrieval call binding the contract method 0x8f32d59b.
//
// Solidity: function isOwner() view returns(bool)
func (_TokenManager *TokenManagerCaller) IsOwner(opts *bind.CallOpts) (bool, error) {
	var (
		ret0 = new(bool)
	)
	out := ret0
	err := _TokenManager.contract.Call(opts, out, "isOwner")
	return *ret0, err
}

// IsOwner is a free data retrieval call binding the contract method 0x8f32d59b.
//
// Solidity: function isOwner() view returns(bool)
func (_TokenManager *TokenManagerSession) IsOwner() (bool, error) {
	return _TokenManager.Contract.IsOwner(&_TokenManager.CallOpts)
}

// IsOwner is a free data retrieval call binding the contract method 0x8f32d59b.
//
// Solidity: function isOwner() view returns(bool)
func (_TokenManager *TokenManagerCallerSession) IsOwner() (bool, error) {
	return _TokenManager.Contract.IsOwner(&_TokenManager.CallOpts)
}

// MappedTokens is a free data retrieval call binding the contract method 0xe6ede14d.
//
// Solidity: function mappedTokens(address ) view returns(address)
func (_TokenManager *TokenManagerCaller) MappedTokens(opts *bind.CallOpts, arg0 common.Address) (common.Address, error) {
	var (
		ret0 = new(common.Address)
	)
	out := ret0
	err := _TokenManager.contract.Call(opts, out, "mappedTokens", arg0)
	return *ret0, err
}

// MappedTokens is a free data retrieval call binding the contract method 0xe6ede14d.
//
// Solidity: function mappedTokens(address ) view returns(address)
func (_TokenManager *TokenManagerSession) MappedTokens(arg0 common.Address) (common.Address, error) {
	return _TokenManager.Contract.MappedTokens(&_TokenManager.CallOpts, arg0)
}

// MappedTokens is a free data retrieval call binding the contract method 0xe6ede14d.
//
// Solidity: function mappedTokens(address ) view returns(address)
func (_TokenManager *TokenManagerCallerSession) MappedTokens(arg0 common.Address) (common.Address, error) {
	return _TokenManager.Contract.MappedTokens(&_TokenManager.CallOpts, arg0)
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_TokenManager *TokenManagerCaller) Owner(opts *bind.CallOpts) (common.Address, error) {
	var (
		ret0 = new(common.Address)
	)
	out := ret0
	err := _TokenManager.contract.Call(opts, out, "owner")
	return *ret0, err
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_TokenManager *TokenManagerSession) Owner() (common.Address, error) {
	return _TokenManager.Contract.Owner(&_TokenManager.CallOpts)
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_TokenManager *TokenManagerCallerSession) Owner() (common.Address, error) {
	return _TokenManager.Contract.Owner(&_TokenManager.CallOpts)
}

// Wards is a free data retrieval call binding the contract method 0xbf353dbb.
//
// Solidity: function wards(address ) view returns(uint256)
func (_TokenManager *TokenManagerCaller) Wards(opts *bind.CallOpts, arg0 common.Address) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _TokenManager.contract.Call(opts, out, "wards", arg0)
	return *ret0, err
}

// Wards is a free data retrieval call binding the contract method 0xbf353dbb.
//
// Solidity: function wards(address ) view returns(uint256)
func (_TokenManager *TokenManagerSession) Wards(arg0 common.Address) (*big.Int, error) {
	return _TokenManager.Contract.Wards(&_TokenManager.CallOpts, arg0)
}

// Wards is a free data retrieval call binding the contract method 0xbf353dbb.
//
// Solidity: function wards(address ) view returns(uint256)
func (_TokenManager *TokenManagerCallerSession) Wards(arg0 common.Address) (*big.Int, error) {
	return _TokenManager.Contract.Wards(&_TokenManager.CallOpts, arg0)
}

// AddToken is a paid mutator transaction binding the contract method 0x8e404e9c.
//
// Solidity: function addToken(address ethTokenAddr, string name, string symbol, uint8 decimals) returns(address)
func (_TokenManager *TokenManagerTransactor) AddToken(opts *bind.TransactOpts, ethTokenAddr common.Address, name string, symbol string, decimals uint8) (*types.Transaction, error) {
	return _TokenManager.contract.Transact(opts, "addToken", ethTokenAddr, name, symbol, decimals)
}

// AddToken is a paid mutator transaction binding the contract method 0x8e404e9c.
//
// Solidity: function addToken(address ethTokenAddr, string name, string symbol, uint8 decimals) returns(address)
func (_TokenManager *TokenManagerSession) AddToken(ethTokenAddr common.Address, name string, symbol string, decimals uint8) (*types.Transaction, error) {
	return _TokenManager.Contract.AddToken(&_TokenManager.TransactOpts, ethTokenAddr, name, symbol, decimals)
}

// AddToken is a paid mutator transaction binding the contract method 0x8e404e9c.
//
// Solidity: function addToken(address ethTokenAddr, string name, string symbol, uint8 decimals) returns(address)
func (_TokenManager *TokenManagerTransactorSession) AddToken(ethTokenAddr common.Address, name string, symbol string, decimals uint8) (*types.Transaction, error) {
	return _TokenManager.Contract.AddToken(&_TokenManager.TransactOpts, ethTokenAddr, name, symbol, decimals)
}

// Deny is a paid mutator transaction binding the contract method 0x9c52a7f1.
//
// Solidity: function deny(address guy) returns()
func (_TokenManager *TokenManagerTransactor) Deny(opts *bind.TransactOpts, guy common.Address) (*types.Transaction, error) {
	return _TokenManager.contract.Transact(opts, "deny", guy)
}

// Deny is a paid mutator transaction binding the contract method 0x9c52a7f1.
//
// Solidity: function deny(address guy) returns()
func (_TokenManager *TokenManagerSession) Deny(guy common.Address) (*types.Transaction, error) {
	return _TokenManager.Contract.Deny(&_TokenManager.TransactOpts, guy)
}

// Deny is a paid mutator transaction binding the contract method 0x9c52a7f1.
//
// Solidity: function deny(address guy) returns()
func (_TokenManager *TokenManagerTransactorSession) Deny(guy common.Address) (*types.Transaction, error) {
	return _TokenManager.Contract.Deny(&_TokenManager.TransactOpts, guy)
}

// RegisterToken is a paid mutator transaction binding the contract method 0x4739f7e5.
//
// Solidity: function registerToken(address ethTokenAddr, address hecoTokenAddr) returns(bool)
func (_TokenManager *TokenManagerTransactor) RegisterToken(opts *bind.TransactOpts, ethTokenAddr common.Address, hecoTokenAddr common.Address) (*types.Transaction, error) {
	return _TokenManager.contract.Transact(opts, "registerToken", ethTokenAddr, hecoTokenAddr)
}

// RegisterToken is a paid mutator transaction binding the contract method 0x4739f7e5.
//
// Solidity: function registerToken(address ethTokenAddr, address hecoTokenAddr) returns(bool)
func (_TokenManager *TokenManagerSession) RegisterToken(ethTokenAddr common.Address, hecoTokenAddr common.Address) (*types.Transaction, error) {
	return _TokenManager.Contract.RegisterToken(&_TokenManager.TransactOpts, ethTokenAddr, hecoTokenAddr)
}

// RegisterToken is a paid mutator transaction binding the contract method 0x4739f7e5.
//
// Solidity: function registerToken(address ethTokenAddr, address hecoTokenAddr) returns(bool)
func (_TokenManager *TokenManagerTransactorSession) RegisterToken(ethTokenAddr common.Address, hecoTokenAddr common.Address) (*types.Transaction, error) {
	return _TokenManager.Contract.RegisterToken(&_TokenManager.TransactOpts, ethTokenAddr, hecoTokenAddr)
}

// Rely is a paid mutator transaction binding the contract method 0x65fae35e.
//
// Solidity: function rely(address guy) returns()
func (_TokenManager *TokenManagerTransactor) Rely(opts *bind.TransactOpts, guy common.Address) (*types.Transaction, error) {
	return _TokenManager.contract.Transact(opts, "rely", guy)
}

// Rely is a paid mutator transaction binding the contract method 0x65fae35e.
//
// Solidity: function rely(address guy) returns()
func (_TokenManager *TokenManagerSession) Rely(guy common.Address) (*types.Transaction, error) {
	return _TokenManager.Contract.Rely(&_TokenManager.TransactOpts, guy)
}

// Rely is a paid mutator transaction binding the contract method 0x65fae35e.
//
// Solidity: function rely(address guy) returns()
func (_TokenManager *TokenManagerTransactorSession) Rely(guy common.Address) (*types.Transaction, error) {
	return _TokenManager.Contract.Rely(&_TokenManager.TransactOpts, guy)
}

// RemoveToken is a paid mutator transaction binding the contract method 0x13baf1e6.
//
// Solidity: function removeToken(address ethTokenAddr, uint256 supply) returns()
func (_TokenManager *TokenManagerTransactor) RemoveToken(opts *bind.TransactOpts, ethTokenAddr common.Address, supply *big.Int) (*types.Transaction, error) {
	return _TokenManager.contract.Transact(opts, "removeToken", ethTokenAddr, supply)
}

// RemoveToken is a paid mutator transaction binding the contract method 0x13baf1e6.
//
// Solidity: function removeToken(address ethTokenAddr, uint256 supply) returns()
func (_TokenManager *TokenManagerSession) RemoveToken(ethTokenAddr common.Address, supply *big.Int) (*types.Transaction, error) {
	return _TokenManager.Contract.RemoveToken(&_TokenManager.TransactOpts, ethTokenAddr, supply)
}

// RemoveToken is a paid mutator transaction binding the contract method 0x13baf1e6.
//
// Solidity: function removeToken(address ethTokenAddr, uint256 supply) returns()
func (_TokenManager *TokenManagerTransactorSession) RemoveToken(ethTokenAddr common.Address, supply *big.Int) (*types.Transaction, error) {
	return _TokenManager.Contract.RemoveToken(&_TokenManager.TransactOpts, ethTokenAddr, supply)
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_TokenManager *TokenManagerTransactor) RenounceOwnership(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _TokenManager.contract.Transact(opts, "renounceOwnership")
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_TokenManager *TokenManagerSession) RenounceOwnership() (*types.Transaction, error) {
	return _TokenManager.Contract.RenounceOwnership(&_TokenManager.TransactOpts)
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_TokenManager *TokenManagerTransactorSession) RenounceOwnership() (*types.Transaction, error) {
	return _TokenManager.Contract.RenounceOwnership(&_TokenManager.TransactOpts)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_TokenManager *TokenManagerTransactor) TransferOwnership(opts *bind.TransactOpts, newOwner common.Address) (*types.Transaction, error) {
	return _TokenManager.contract.Transact(opts, "transferOwnership", newOwner)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_TokenManager *TokenManagerSession) TransferOwnership(newOwner common.Address) (*types.Transaction, error) {
	return _TokenManager.Contract.TransferOwnership(&_TokenManager.TransactOpts, newOwner)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_TokenManager *TokenManagerTransactorSession) TransferOwnership(newOwner common.Address) (*types.Transaction, error) {
	return _TokenManager.Contract.TransferOwnership(&_TokenManager.TransactOpts, newOwner)
}

// TokenManagerOwnershipTransferredIterator is returned from FilterOwnershipTransferred and is used to iterate over the raw logs and unpacked data for OwnershipTransferred events raised by the TokenManager contract.
type TokenManagerOwnershipTransferredIterator struct {
	Event *TokenManagerOwnershipTransferred // Event containing the contract specifics and raw log

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
func (it *TokenManagerOwnershipTransferredIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(TokenManagerOwnershipTransferred)
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
		it.Event = new(TokenManagerOwnershipTransferred)
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
func (it *TokenManagerOwnershipTransferredIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *TokenManagerOwnershipTransferredIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// TokenManagerOwnershipTransferred represents a OwnershipTransferred event raised by the TokenManager contract.
type TokenManagerOwnershipTransferred struct {
	PreviousOwner common.Address
	NewOwner      common.Address
	Raw           types.Log // Blockchain specific contextual infos
}

// FilterOwnershipTransferred is a free log retrieval operation binding the contract event 0x8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0.
//
// Solidity: event OwnershipTransferred(address indexed previousOwner, address indexed newOwner)
func (_TokenManager *TokenManagerFilterer) FilterOwnershipTransferred(opts *bind.FilterOpts, previousOwner []common.Address, newOwner []common.Address) (*TokenManagerOwnershipTransferredIterator, error) {

	var previousOwnerRule []interface{}
	for _, previousOwnerItem := range previousOwner {
		previousOwnerRule = append(previousOwnerRule, previousOwnerItem)
	}
	var newOwnerRule []interface{}
	for _, newOwnerItem := range newOwner {
		newOwnerRule = append(newOwnerRule, newOwnerItem)
	}

	logs, sub, err := _TokenManager.contract.FilterLogs(opts, "OwnershipTransferred", previousOwnerRule, newOwnerRule)
	if err != nil {
		return nil, err
	}
	return &TokenManagerOwnershipTransferredIterator{contract: _TokenManager.contract, event: "OwnershipTransferred", logs: logs, sub: sub}, nil
}

// WatchOwnershipTransferred is a free log subscription operation binding the contract event 0x8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0.
//
// Solidity: event OwnershipTransferred(address indexed previousOwner, address indexed newOwner)
func (_TokenManager *TokenManagerFilterer) WatchOwnershipTransferred(opts *bind.WatchOpts, sink chan<- *TokenManagerOwnershipTransferred, previousOwner []common.Address, newOwner []common.Address) (event.Subscription, error) {

	var previousOwnerRule []interface{}
	for _, previousOwnerItem := range previousOwner {
		previousOwnerRule = append(previousOwnerRule, previousOwnerItem)
	}
	var newOwnerRule []interface{}
	for _, newOwnerItem := range newOwner {
		newOwnerRule = append(newOwnerRule, newOwnerItem)
	}

	logs, sub, err := _TokenManager.contract.WatchLogs(opts, "OwnershipTransferred", previousOwnerRule, newOwnerRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(TokenManagerOwnershipTransferred)
				if err := _TokenManager.contract.UnpackLog(event, "OwnershipTransferred", log); err != nil {
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
func (_TokenManager *TokenManagerFilterer) ParseOwnershipTransferred(log types.Log) (*TokenManagerOwnershipTransferred, error) {
	event := new(TokenManagerOwnershipTransferred)
	if err := _TokenManager.contract.UnpackLog(event, "OwnershipTransferred", log); err != nil {
		return nil, err
	}
	return event, nil
}

// TokenManagerTokenMapAckIterator is returned from FilterTokenMapAck and is used to iterate over the raw logs and unpacked data for TokenMapAck events raised by the TokenManager contract.
type TokenManagerTokenMapAckIterator struct {
	Event *TokenManagerTokenMapAck // Event containing the contract specifics and raw log

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
func (it *TokenManagerTokenMapAckIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(TokenManagerTokenMapAck)
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
		it.Event = new(TokenManagerTokenMapAck)
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
func (it *TokenManagerTokenMapAckIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *TokenManagerTokenMapAckIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// TokenManagerTokenMapAck represents a TokenMapAck event raised by the TokenManager contract.
type TokenManagerTokenMapAck struct {
	TokenReq common.Address
	TokenAck common.Address
	Raw      types.Log // Blockchain specific contextual infos
}

// FilterTokenMapAck is a free log retrieval operation binding the contract event 0x78591f651c27eef63481b7fd779e44c2426cbe82c7050fe9bc90b6707efbc3d3.
//
// Solidity: event TokenMapAck(address indexed tokenReq, address indexed tokenAck)
func (_TokenManager *TokenManagerFilterer) FilterTokenMapAck(opts *bind.FilterOpts, tokenReq []common.Address, tokenAck []common.Address) (*TokenManagerTokenMapAckIterator, error) {

	var tokenReqRule []interface{}
	for _, tokenReqItem := range tokenReq {
		tokenReqRule = append(tokenReqRule, tokenReqItem)
	}
	var tokenAckRule []interface{}
	for _, tokenAckItem := range tokenAck {
		tokenAckRule = append(tokenAckRule, tokenAckItem)
	}

	logs, sub, err := _TokenManager.contract.FilterLogs(opts, "TokenMapAck", tokenReqRule, tokenAckRule)
	if err != nil {
		return nil, err
	}
	return &TokenManagerTokenMapAckIterator{contract: _TokenManager.contract, event: "TokenMapAck", logs: logs, sub: sub}, nil
}

// WatchTokenMapAck is a free log subscription operation binding the contract event 0x78591f651c27eef63481b7fd779e44c2426cbe82c7050fe9bc90b6707efbc3d3.
//
// Solidity: event TokenMapAck(address indexed tokenReq, address indexed tokenAck)
func (_TokenManager *TokenManagerFilterer) WatchTokenMapAck(opts *bind.WatchOpts, sink chan<- *TokenManagerTokenMapAck, tokenReq []common.Address, tokenAck []common.Address) (event.Subscription, error) {

	var tokenReqRule []interface{}
	for _, tokenReqItem := range tokenReq {
		tokenReqRule = append(tokenReqRule, tokenReqItem)
	}
	var tokenAckRule []interface{}
	for _, tokenAckItem := range tokenAck {
		tokenAckRule = append(tokenAckRule, tokenAckItem)
	}

	logs, sub, err := _TokenManager.contract.WatchLogs(opts, "TokenMapAck", tokenReqRule, tokenAckRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(TokenManagerTokenMapAck)
				if err := _TokenManager.contract.UnpackLog(event, "TokenMapAck", log); err != nil {
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

// ParseTokenMapAck is a log parse operation binding the contract event 0x78591f651c27eef63481b7fd779e44c2426cbe82c7050fe9bc90b6707efbc3d3.
//
// Solidity: event TokenMapAck(address indexed tokenReq, address indexed tokenAck)
func (_TokenManager *TokenManagerFilterer) ParseTokenMapAck(log types.Log) (*TokenManagerTokenMapAck, error) {
	event := new(TokenManagerTokenMapAck)
	if err := _TokenManager.contract.UnpackLog(event, "TokenMapAck", log); err != nil {
		return nil, err
	}
	return event, nil
}
