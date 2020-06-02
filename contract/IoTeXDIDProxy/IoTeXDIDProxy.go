// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package IoTeXDIDProxy

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

// IoTeXDIDProxyABI is the input ABI used to generate the binding from.
const IoTeXDIDProxyABI = "[{\"inputs\":[{\"internalType\":\"address\",\"name\":\"addr\",\"type\":\"address\"}],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"constructor\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"previousOwner\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"newOwner\",\"type\":\"address\"}],\"name\":\"OwnershipTransferred\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"string\",\"name\":\"version\",\"type\":\"string\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"addr\",\"type\":\"address\"}],\"name\":\"Upgrade\",\"type\":\"event\"},{\"payable\":true,\"stateMutability\":\"payable\",\"type\":\"fallback\"},{\"constant\":true,\"inputs\":[{\"internalType\":\"string\",\"name\":\"\",\"type\":\"string\"}],\"name\":\"allVersions\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"currentDIDAddress\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"currentVersion\",\"outputs\":[{\"internalType\":\"string\",\"name\":\"\",\"type\":\"string\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"internalType\":\"string\",\"name\":\"\",\"type\":\"string\"}],\"name\":\"dids\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"exist\",\"type\":\"bool\"},{\"internalType\":\"bytes32\",\"name\":\"hash\",\"type\":\"bytes32\"},{\"internalType\":\"string\",\"name\":\"uri\",\"type\":\"string\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"internalType\":\"string\",\"name\":\"version\",\"type\":\"string\"}],\"name\":\"getDIDAddress\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"owner\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"internalType\":\"address\",\"name\":\"newOwner\",\"type\":\"address\"}],\"name\":\"transferOwnership\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"internalType\":\"string\",\"name\":\"version\",\"type\":\"string\"},{\"internalType\":\"address\",\"name\":\"addr\",\"type\":\"address\"}],\"name\":\"upgrade\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"name\":\"versionList\",\"outputs\":[{\"internalType\":\"string\",\"name\":\"\",\"type\":\"string\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"}]"

// IoTeXDIDProxyFuncSigs maps the 4-byte function signature to its string representation.
var IoTeXDIDProxyFuncSigs = map[string]string{
	"264ecb71": "allVersions(string)",
	"8e35d737": "currentDIDAddress()",
	"9d888e86": "currentVersion()",
	"f44ab516": "dids(string)",
	"3e805761": "getDIDAddress(string)",
	"8da5cb5b": "owner()",
	"f2fde38b": "transferOwnership(address)",
	"5332c74d": "upgrade(string,address)",
	"1d8f5f18": "versionList(uint256)",
}

// IoTeXDIDProxyBin is the compiled bytecode used for deploying new contracts.
var IoTeXDIDProxyBin = "0x60806040523480156200001157600080fd5b506040516200141e3803806200141e833981810160405260208110156200003757600080fd5b5051600180546001600160a01b03191633179055604080518082019091526005815264302e302e3160d81b60208201526200007c90826001600160e01b036200008316565b5062000554565b6001546001600160a01b031633146200009b57600080fd5b6001600160a01b038116620000b86001600160e01b036200037316565b6001600160a01b031614158015620000d857506001600160a01b03811615155b620001155760405162461bcd60e51b8152600401808060200182810382526040815260200180620013a96040913960400191505060405180910390fd5b62000129816001600160e01b03620003fb16565b620001665760405162461bcd60e51b8152600401808060200182810382526035815260200180620013e96035913960400191505060405180910390fd5b6000825111620001a85760405162461bcd60e51b8152600401808060200182810382526022815260200180620013876022913960400191505060405180910390fd5b8151620001bd90600490602085019062000438565b5080600260046040518082805460018160011615610100020316600290048015620002225780601f10620001ff57610100808354040283529182019162000222565b820191906000526020600020905b8154815290600101906020018083116200020d575b5050928352505060405190819003602001902080546001600160a01b03929092166001600160a01b031990921691909117905560038054600181810180845560009390935260048054620002ab937fc2575a0e9e593c00f959f8c92f12db2869c3395a3b0502d05e2516446f71f85b0192600290821615610100026000190190911604620004bd565b5050604080516001600160a01b038316602082015281815260048054600260001961010060018416150201909116049282018390527f6f48f7d6926b162cb7a11839a8d06721e35a9124ee35ee79d9faae8d34ab82ce92909184918190606082019085908015620003605780601f10620003345761010080835404028352916020019162000360565b820191906000526020600020905b8154815290600101906020018083116200034257829003601f168201915b5050935050505060405180910390a15050565b6000600260046040518082805460018160011615610100020316600290048015620003d85780601f10620003b5576101008083540402835291820191620003d8565b820191906000526020600020905b815481529060010190602001808311620003c3575b50509283525050604051908190036020019020546001600160a01b031690505b90565b6000813f7fc5d2460186f7233c927e7db2dcc703c0e500b653ca82273b7bfad8045d85a4708181148015906200043057508115155b949350505050565b828054600181600116156101000203166002900490600052602060002090601f016020900481019282601f106200047b57805160ff1916838001178555620004ab565b82800160010185558215620004ab579182015b82811115620004ab5782518255916020019190600101906200048e565b50620004b992915062000537565b5090565b828054600181600116156101000203166002900490600052602060002090601f016020900481019282601f10620004f85780548555620004ab565b82800160010185558215620004ab57600052602060002091601f016020900482015b82811115620004ab5782548255916001019190600101906200051a565b620003f891905b80821115620004b957600081556001016200053e565b610e2380620005646000396000f3fe6080604052600436106100865760003560e01c80638da5cb5b116100595780638da5cb5b146103ed5780638e35d737146104025780639d888e8614610417578063f2fde38b1461042c578063f44ab5161461045f57610086565b80631d8f5f1814610112578063264ecb71146101b15780633e8057611461027e5780635332c74d1461032f575b600061009061059a565b90506001600160a01b0381166100ed576040805162461bcd60e51b815260206004820152601c60248201527f64696420636f6e74726163742061646472657373206e6f742073657400000000604482015290519081900360640190fd5b60405136600082376000803683855af43d806000843e81801561010e578184f35b8184fd5b34801561011e57600080fd5b5061013c6004803603602081101561013557600080fd5b503561061e565b6040805160208082528351818301528351919283929083019185019080838360005b8381101561017657818101518382015260200161015e565b50505050905090810190601f1680156101a35780820380516001836020036101000a031916815260200191505b509250505060405180910390f35b3480156101bd57600080fd5b50610262600480360360208110156101d457600080fd5b810190602081018135600160201b8111156101ee57600080fd5b82018360208201111561020057600080fd5b803590602001918460018302840111600160201b8311171561022157600080fd5b91908080601f0160208091040260200160405190810160405280939291908181526020018383808284376000920191909152509295506106c4945050505050565b604080516001600160a01b039092168252519081900360200190f35b34801561028a57600080fd5b50610262600480360360208110156102a157600080fd5b810190602081018135600160201b8111156102bb57600080fd5b8201836020820111156102cd57600080fd5b803590602001918460018302840111600160201b831117156102ee57600080fd5b91908080601f0160208091040260200160405190810160405280939291908181526020018383808284376000920191909152509295506106ea945050505050565b34801561033b57600080fd5b506103eb6004803603604081101561035257600080fd5b810190602081018135600160201b81111561036c57600080fd5b82018360208201111561037e57600080fd5b803590602001918460018302840111600160201b8311171561039f57600080fd5b91908080601f016020809104026020016040519081016040528093929190818152602001838380828437600092019190915250929550505090356001600160a01b0316915061079a9050565b005b3480156103f957600080fd5b50610262610a60565b34801561040e57600080fd5b5061026261059a565b34801561042357600080fd5b5061013c610a6f565b34801561043857600080fd5b506103eb6004803603602081101561044f57600080fd5b50356001600160a01b0316610aca565b34801561046b57600080fd5b506105106004803603602081101561048257600080fd5b810190602081018135600160201b81111561049c57600080fd5b8201836020820111156104ae57600080fd5b803590602001918460018302840111600160201b831117156104cf57600080fd5b91908080601f016020809104026020016040519081016040528093929190818152602001838380828437600092019190915250929550610b50945050505050565b604051808415151515815260200183815260200180602001828103825283818151815260200191508051906020019080838360005b8381101561055d578181015183820152602001610545565b50505050905090810190601f16801561058a5780820380516001836020036101000a031916815260200191505b5094505050505060405180910390f35b60006002600460405180828054600181600116156101000203166002900480156105fb5780601f106105d95761010080835404028352918201916105fb565b820191906000526020600020905b8154815290600101906020018083116105e7575b50509283525050604051908190036020019020546001600160a01b031690505b90565b6003818154811061062b57fe5b600091825260209182902001805460408051601f60026000196101006001871615020190941693909304928301859004850281018501909152818152935090918301828280156106bc5780601f10610691576101008083540402835291602001916106bc565b820191906000526020600020905b81548152906001019060200180831161069f57829003601f168201915b505050505081565b80516020818301810180516002825292820191909301209152546001600160a01b031681565b60008082511161072b5760405162461bcd60e51b8152600401808060200182810382526022815260200180610d586022913960400191505060405180910390fd5b6002826040518082805190602001908083835b6020831061075d5780518252601f19909201916020918201910161073e565b51815160209384036101000a60001901801990921691161790529201948552506040519384900301909220546001600160a01b0316949350505050565b6001546001600160a01b031633146107b157600080fd5b806001600160a01b03166107c361059a565b6001600160a01b0316141580156107e257506001600160a01b03811615155b61081d5760405162461bcd60e51b8152600401808060200182810382526040815260200180610d7a6040913960400191505060405180910390fd5b61082681610c0e565b6108615760405162461bcd60e51b8152600401808060200182810382526035815260200180610dba6035913960400191505060405180910390fd5b60008251116108a15760405162461bcd60e51b8152600401808060200182810382526022815260200180610d586022913960400191505060405180910390fd5b81516108b4906004906020850190610c4a565b50806002600460405180828054600181600116156101000203166002900480156109155780601f106108f3576101008083540402835291820191610915565b820191906000526020600020905b815481529060010190602001808311610901575b5050928352505060405190819003602001902080546001600160a01b03929092166001600160a01b03199092169190911790556003805460018181018084556000939093526004805461099c937fc2575a0e9e593c00f959f8c92f12db2869c3395a3b0502d05e2516446f71f85b0192600290821615610100026000190190911604610cc8565b5050604080516001600160a01b038316602082015281815260048054600260001961010060018416150201909116049282018390527f6f48f7d6926b162cb7a11839a8d06721e35a9124ee35ee79d9faae8d34ab82ce92909184918190606082019085908015610a4d5780601f10610a2257610100808354040283529160200191610a4d565b820191906000526020600020905b815481529060010190602001808311610a3057829003601f168201915b5050935050505060405180910390a15050565b6001546001600160a01b031681565b6004805460408051602060026001851615610100026000190190941693909304601f810184900484028201840190925281815292918301828280156106bc5780601f10610691576101008083540402835291602001916106bc565b6001546001600160a01b03163314610ae157600080fd5b6001600160a01b038116610af457600080fd5b6001546040516001600160a01b038084169216907f8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e090600090a3600180546001600160a01b0319166001600160a01b0392909216919091179055565b80516020818301810180516000825292820193820193909320919092528054600180830154600280850180546040805161010096831615969096026000190190911692909204601f810188900488028501880190925281845260ff9094169591949391830182828015610c045780601f10610bd957610100808354040283529160200191610c04565b820191906000526020600020905b815481529060010190602001808311610be757829003601f168201915b5050505050905083565b6000813f7fc5d2460186f7233c927e7db2dcc703c0e500b653ca82273b7bfad8045d85a470818114801590610c4257508115155b949350505050565b828054600181600116156101000203166002900490600052602060002090601f016020900481019282601f10610c8b57805160ff1916838001178555610cb8565b82800160010185558215610cb8579182015b82811115610cb8578251825591602001919060010190610c9d565b50610cc4929150610d3d565b5090565b828054600181600116156101000203166002900490600052602060002090601f016020900481019282601f10610d015780548555610cb8565b82800160010185558215610cb857600052602060002091601f016020900482015b82811115610cb8578254825591600101919060010190610d22565b61061b91905b80821115610cc45760008155600101610d4356fe56657273696f6e2073686f756c64206e6f7420626520656d70747920737472696e674f6c642061646472657373206973206e6f7420616c6c6f77656420616e6420636f6e747261637420616464726573732073686f756c64206e6f7420626520307843616e6e6f74207365742061206c6f67696320636f6e747261637420746f2061206e6f6e2d636f6e74726163742061646472657373a265627a7a72315820d1fca25ef17b4dd81c5d0d08f9f815e5527e0d1899e80160d862d394fea75d4a64736f6c6343000510003256657273696f6e2073686f756c64206e6f7420626520656d70747920737472696e674f6c642061646472657373206973206e6f7420616c6c6f77656420616e6420636f6e747261637420616464726573732073686f756c64206e6f7420626520307843616e6e6f74207365742061206c6f67696320636f6e747261637420746f2061206e6f6e2d636f6e74726163742061646472657373"

// DeployIoTeXDIDProxy deploys a new Ethereum contract, binding an instance of IoTeXDIDProxy to it.
func DeployIoTeXDIDProxy(auth *bind.TransactOpts, backend bind.ContractBackend, addr common.Address) (common.Address, *types.Transaction, *IoTeXDIDProxy, error) {
	parsed, err := abi.JSON(strings.NewReader(IoTeXDIDProxyABI))
	if err != nil {
		return common.Address{}, nil, nil, err
	}

	address, tx, contract, err := bind.DeployContract(auth, parsed, common.FromHex(IoTeXDIDProxyBin), backend, addr)
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	return address, tx, &IoTeXDIDProxy{IoTeXDIDProxyCaller: IoTeXDIDProxyCaller{contract: contract}, IoTeXDIDProxyTransactor: IoTeXDIDProxyTransactor{contract: contract}, IoTeXDIDProxyFilterer: IoTeXDIDProxyFilterer{contract: contract}}, nil
}

// IoTeXDIDProxy is an auto generated Go binding around an Ethereum contract.
type IoTeXDIDProxy struct {
	IoTeXDIDProxyCaller     // Read-only binding to the contract
	IoTeXDIDProxyTransactor // Write-only binding to the contract
	IoTeXDIDProxyFilterer   // Log filterer for contract events
}

// IoTeXDIDProxyCaller is an auto generated read-only Go binding around an Ethereum contract.
type IoTeXDIDProxyCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// IoTeXDIDProxyTransactor is an auto generated write-only Go binding around an Ethereum contract.
type IoTeXDIDProxyTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// IoTeXDIDProxyFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type IoTeXDIDProxyFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// IoTeXDIDProxySession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type IoTeXDIDProxySession struct {
	Contract     *IoTeXDIDProxy    // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// IoTeXDIDProxyCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type IoTeXDIDProxyCallerSession struct {
	Contract *IoTeXDIDProxyCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts        // Call options to use throughout this session
}

// IoTeXDIDProxyTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type IoTeXDIDProxyTransactorSession struct {
	Contract     *IoTeXDIDProxyTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts        // Transaction auth options to use throughout this session
}

// IoTeXDIDProxyRaw is an auto generated low-level Go binding around an Ethereum contract.
type IoTeXDIDProxyRaw struct {
	Contract *IoTeXDIDProxy // Generic contract binding to access the raw methods on
}

// IoTeXDIDProxyCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type IoTeXDIDProxyCallerRaw struct {
	Contract *IoTeXDIDProxyCaller // Generic read-only contract binding to access the raw methods on
}

// IoTeXDIDProxyTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type IoTeXDIDProxyTransactorRaw struct {
	Contract *IoTeXDIDProxyTransactor // Generic write-only contract binding to access the raw methods on
}

// NewIoTeXDIDProxy creates a new instance of IoTeXDIDProxy, bound to a specific deployed contract.
func NewIoTeXDIDProxy(address common.Address, backend bind.ContractBackend) (*IoTeXDIDProxy, error) {
	contract, err := bindIoTeXDIDProxy(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &IoTeXDIDProxy{IoTeXDIDProxyCaller: IoTeXDIDProxyCaller{contract: contract}, IoTeXDIDProxyTransactor: IoTeXDIDProxyTransactor{contract: contract}, IoTeXDIDProxyFilterer: IoTeXDIDProxyFilterer{contract: contract}}, nil
}

// NewIoTeXDIDProxyCaller creates a new read-only instance of IoTeXDIDProxy, bound to a specific deployed contract.
func NewIoTeXDIDProxyCaller(address common.Address, caller bind.ContractCaller) (*IoTeXDIDProxyCaller, error) {
	contract, err := bindIoTeXDIDProxy(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &IoTeXDIDProxyCaller{contract: contract}, nil
}

// NewIoTeXDIDProxyTransactor creates a new write-only instance of IoTeXDIDProxy, bound to a specific deployed contract.
func NewIoTeXDIDProxyTransactor(address common.Address, transactor bind.ContractTransactor) (*IoTeXDIDProxyTransactor, error) {
	contract, err := bindIoTeXDIDProxy(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &IoTeXDIDProxyTransactor{contract: contract}, nil
}

// NewIoTeXDIDProxyFilterer creates a new log filterer instance of IoTeXDIDProxy, bound to a specific deployed contract.
func NewIoTeXDIDProxyFilterer(address common.Address, filterer bind.ContractFilterer) (*IoTeXDIDProxyFilterer, error) {
	contract, err := bindIoTeXDIDProxy(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &IoTeXDIDProxyFilterer{contract: contract}, nil
}

// bindIoTeXDIDProxy binds a generic wrapper to an already deployed contract.
func bindIoTeXDIDProxy(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(IoTeXDIDProxyABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_IoTeXDIDProxy *IoTeXDIDProxyRaw) Call(opts *bind.CallOpts, result interface{}, method string, params ...interface{}) error {
	return _IoTeXDIDProxy.Contract.IoTeXDIDProxyCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_IoTeXDIDProxy *IoTeXDIDProxyRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _IoTeXDIDProxy.Contract.IoTeXDIDProxyTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_IoTeXDIDProxy *IoTeXDIDProxyRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _IoTeXDIDProxy.Contract.IoTeXDIDProxyTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_IoTeXDIDProxy *IoTeXDIDProxyCallerRaw) Call(opts *bind.CallOpts, result interface{}, method string, params ...interface{}) error {
	return _IoTeXDIDProxy.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_IoTeXDIDProxy *IoTeXDIDProxyTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _IoTeXDIDProxy.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_IoTeXDIDProxy *IoTeXDIDProxyTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _IoTeXDIDProxy.Contract.contract.Transact(opts, method, params...)
}

// AllVersions is a free data retrieval call binding the contract method 0x264ecb71.
//
// Solidity: function allVersions(string ) view returns(address)
func (_IoTeXDIDProxy *IoTeXDIDProxyCaller) AllVersions(opts *bind.CallOpts, arg0 string) (common.Address, error) {
	var (
		ret0 = new(common.Address)
	)
	out := ret0
	err := _IoTeXDIDProxy.contract.Call(opts, out, "allVersions", arg0)
	return *ret0, err
}

// AllVersions is a free data retrieval call binding the contract method 0x264ecb71.
//
// Solidity: function allVersions(string ) view returns(address)
func (_IoTeXDIDProxy *IoTeXDIDProxySession) AllVersions(arg0 string) (common.Address, error) {
	return _IoTeXDIDProxy.Contract.AllVersions(&_IoTeXDIDProxy.CallOpts, arg0)
}

// AllVersions is a free data retrieval call binding the contract method 0x264ecb71.
//
// Solidity: function allVersions(string ) view returns(address)
func (_IoTeXDIDProxy *IoTeXDIDProxyCallerSession) AllVersions(arg0 string) (common.Address, error) {
	return _IoTeXDIDProxy.Contract.AllVersions(&_IoTeXDIDProxy.CallOpts, arg0)
}

// CurrentDIDAddress is a free data retrieval call binding the contract method 0x8e35d737.
//
// Solidity: function currentDIDAddress() view returns(address)
func (_IoTeXDIDProxy *IoTeXDIDProxyCaller) CurrentDIDAddress(opts *bind.CallOpts) (common.Address, error) {
	var (
		ret0 = new(common.Address)
	)
	out := ret0
	err := _IoTeXDIDProxy.contract.Call(opts, out, "currentDIDAddress")
	return *ret0, err
}

// CurrentDIDAddress is a free data retrieval call binding the contract method 0x8e35d737.
//
// Solidity: function currentDIDAddress() view returns(address)
func (_IoTeXDIDProxy *IoTeXDIDProxySession) CurrentDIDAddress() (common.Address, error) {
	return _IoTeXDIDProxy.Contract.CurrentDIDAddress(&_IoTeXDIDProxy.CallOpts)
}

// CurrentDIDAddress is a free data retrieval call binding the contract method 0x8e35d737.
//
// Solidity: function currentDIDAddress() view returns(address)
func (_IoTeXDIDProxy *IoTeXDIDProxyCallerSession) CurrentDIDAddress() (common.Address, error) {
	return _IoTeXDIDProxy.Contract.CurrentDIDAddress(&_IoTeXDIDProxy.CallOpts)
}

// CurrentVersion is a free data retrieval call binding the contract method 0x9d888e86.
//
// Solidity: function currentVersion() view returns(string)
func (_IoTeXDIDProxy *IoTeXDIDProxyCaller) CurrentVersion(opts *bind.CallOpts) (string, error) {
	var (
		ret0 = new(string)
	)
	out := ret0
	err := _IoTeXDIDProxy.contract.Call(opts, out, "currentVersion")
	return *ret0, err
}

// CurrentVersion is a free data retrieval call binding the contract method 0x9d888e86.
//
// Solidity: function currentVersion() view returns(string)
func (_IoTeXDIDProxy *IoTeXDIDProxySession) CurrentVersion() (string, error) {
	return _IoTeXDIDProxy.Contract.CurrentVersion(&_IoTeXDIDProxy.CallOpts)
}

// CurrentVersion is a free data retrieval call binding the contract method 0x9d888e86.
//
// Solidity: function currentVersion() view returns(string)
func (_IoTeXDIDProxy *IoTeXDIDProxyCallerSession) CurrentVersion() (string, error) {
	return _IoTeXDIDProxy.Contract.CurrentVersion(&_IoTeXDIDProxy.CallOpts)
}

// Dids is a free data retrieval call binding the contract method 0xf44ab516.
//
// Solidity: function dids(string ) view returns(bool exist, bytes32 hash, string uri)
func (_IoTeXDIDProxy *IoTeXDIDProxyCaller) Dids(opts *bind.CallOpts, arg0 string) (struct {
	Exist bool
	Hash  [32]byte
	Uri   string
}, error) {
	ret := new(struct {
		Exist bool
		Hash  [32]byte
		Uri   string
	})
	out := ret
	err := _IoTeXDIDProxy.contract.Call(opts, out, "dids", arg0)
	return *ret, err
}

// Dids is a free data retrieval call binding the contract method 0xf44ab516.
//
// Solidity: function dids(string ) view returns(bool exist, bytes32 hash, string uri)
func (_IoTeXDIDProxy *IoTeXDIDProxySession) Dids(arg0 string) (struct {
	Exist bool
	Hash  [32]byte
	Uri   string
}, error) {
	return _IoTeXDIDProxy.Contract.Dids(&_IoTeXDIDProxy.CallOpts, arg0)
}

// Dids is a free data retrieval call binding the contract method 0xf44ab516.
//
// Solidity: function dids(string ) view returns(bool exist, bytes32 hash, string uri)
func (_IoTeXDIDProxy *IoTeXDIDProxyCallerSession) Dids(arg0 string) (struct {
	Exist bool
	Hash  [32]byte
	Uri   string
}, error) {
	return _IoTeXDIDProxy.Contract.Dids(&_IoTeXDIDProxy.CallOpts, arg0)
}

// GetDIDAddress is a free data retrieval call binding the contract method 0x3e805761.
//
// Solidity: function getDIDAddress(string version) view returns(address)
func (_IoTeXDIDProxy *IoTeXDIDProxyCaller) GetDIDAddress(opts *bind.CallOpts, version string) (common.Address, error) {
	var (
		ret0 = new(common.Address)
	)
	out := ret0
	err := _IoTeXDIDProxy.contract.Call(opts, out, "getDIDAddress", version)
	return *ret0, err
}

// GetDIDAddress is a free data retrieval call binding the contract method 0x3e805761.
//
// Solidity: function getDIDAddress(string version) view returns(address)
func (_IoTeXDIDProxy *IoTeXDIDProxySession) GetDIDAddress(version string) (common.Address, error) {
	return _IoTeXDIDProxy.Contract.GetDIDAddress(&_IoTeXDIDProxy.CallOpts, version)
}

// GetDIDAddress is a free data retrieval call binding the contract method 0x3e805761.
//
// Solidity: function getDIDAddress(string version) view returns(address)
func (_IoTeXDIDProxy *IoTeXDIDProxyCallerSession) GetDIDAddress(version string) (common.Address, error) {
	return _IoTeXDIDProxy.Contract.GetDIDAddress(&_IoTeXDIDProxy.CallOpts, version)
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_IoTeXDIDProxy *IoTeXDIDProxyCaller) Owner(opts *bind.CallOpts) (common.Address, error) {
	var (
		ret0 = new(common.Address)
	)
	out := ret0
	err := _IoTeXDIDProxy.contract.Call(opts, out, "owner")
	return *ret0, err
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_IoTeXDIDProxy *IoTeXDIDProxySession) Owner() (common.Address, error) {
	return _IoTeXDIDProxy.Contract.Owner(&_IoTeXDIDProxy.CallOpts)
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_IoTeXDIDProxy *IoTeXDIDProxyCallerSession) Owner() (common.Address, error) {
	return _IoTeXDIDProxy.Contract.Owner(&_IoTeXDIDProxy.CallOpts)
}

// VersionList is a free data retrieval call binding the contract method 0x1d8f5f18.
//
// Solidity: function versionList(uint256 ) view returns(string)
func (_IoTeXDIDProxy *IoTeXDIDProxyCaller) VersionList(opts *bind.CallOpts, arg0 *big.Int) (string, error) {
	var (
		ret0 = new(string)
	)
	out := ret0
	err := _IoTeXDIDProxy.contract.Call(opts, out, "versionList", arg0)
	return *ret0, err
}

// VersionList is a free data retrieval call binding the contract method 0x1d8f5f18.
//
// Solidity: function versionList(uint256 ) view returns(string)
func (_IoTeXDIDProxy *IoTeXDIDProxySession) VersionList(arg0 *big.Int) (string, error) {
	return _IoTeXDIDProxy.Contract.VersionList(&_IoTeXDIDProxy.CallOpts, arg0)
}

// VersionList is a free data retrieval call binding the contract method 0x1d8f5f18.
//
// Solidity: function versionList(uint256 ) view returns(string)
func (_IoTeXDIDProxy *IoTeXDIDProxyCallerSession) VersionList(arg0 *big.Int) (string, error) {
	return _IoTeXDIDProxy.Contract.VersionList(&_IoTeXDIDProxy.CallOpts, arg0)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_IoTeXDIDProxy *IoTeXDIDProxyTransactor) TransferOwnership(opts *bind.TransactOpts, newOwner common.Address) (*types.Transaction, error) {
	return _IoTeXDIDProxy.contract.Transact(opts, "transferOwnership", newOwner)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_IoTeXDIDProxy *IoTeXDIDProxySession) TransferOwnership(newOwner common.Address) (*types.Transaction, error) {
	return _IoTeXDIDProxy.Contract.TransferOwnership(&_IoTeXDIDProxy.TransactOpts, newOwner)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_IoTeXDIDProxy *IoTeXDIDProxyTransactorSession) TransferOwnership(newOwner common.Address) (*types.Transaction, error) {
	return _IoTeXDIDProxy.Contract.TransferOwnership(&_IoTeXDIDProxy.TransactOpts, newOwner)
}

// Upgrade is a paid mutator transaction binding the contract method 0x5332c74d.
//
// Solidity: function upgrade(string version, address addr) returns()
func (_IoTeXDIDProxy *IoTeXDIDProxyTransactor) Upgrade(opts *bind.TransactOpts, version string, addr common.Address) (*types.Transaction, error) {
	return _IoTeXDIDProxy.contract.Transact(opts, "upgrade", version, addr)
}

// Upgrade is a paid mutator transaction binding the contract method 0x5332c74d.
//
// Solidity: function upgrade(string version, address addr) returns()
func (_IoTeXDIDProxy *IoTeXDIDProxySession) Upgrade(version string, addr common.Address) (*types.Transaction, error) {
	return _IoTeXDIDProxy.Contract.Upgrade(&_IoTeXDIDProxy.TransactOpts, version, addr)
}

// Upgrade is a paid mutator transaction binding the contract method 0x5332c74d.
//
// Solidity: function upgrade(string version, address addr) returns()
func (_IoTeXDIDProxy *IoTeXDIDProxyTransactorSession) Upgrade(version string, addr common.Address) (*types.Transaction, error) {
	return _IoTeXDIDProxy.Contract.Upgrade(&_IoTeXDIDProxy.TransactOpts, version, addr)
}

// Fallback is a paid mutator transaction binding the contract fallback function.
//
// Solidity: fallback() payable returns()
func (_IoTeXDIDProxy *IoTeXDIDProxyTransactor) Fallback(opts *bind.TransactOpts, calldata []byte) (*types.Transaction, error) {
	return _IoTeXDIDProxy.contract.RawTransact(opts, calldata)
}

// Fallback is a paid mutator transaction binding the contract fallback function.
//
// Solidity: fallback() payable returns()
func (_IoTeXDIDProxy *IoTeXDIDProxySession) Fallback(calldata []byte) (*types.Transaction, error) {
	return _IoTeXDIDProxy.Contract.Fallback(&_IoTeXDIDProxy.TransactOpts, calldata)
}

// Fallback is a paid mutator transaction binding the contract fallback function.
//
// Solidity: fallback() payable returns()
func (_IoTeXDIDProxy *IoTeXDIDProxyTransactorSession) Fallback(calldata []byte) (*types.Transaction, error) {
	return _IoTeXDIDProxy.Contract.Fallback(&_IoTeXDIDProxy.TransactOpts, calldata)
}

// IoTeXDIDProxyOwnershipTransferredIterator is returned from FilterOwnershipTransferred and is used to iterate over the raw logs and unpacked data for OwnershipTransferred events raised by the IoTeXDIDProxy contract.
type IoTeXDIDProxyOwnershipTransferredIterator struct {
	Event *IoTeXDIDProxyOwnershipTransferred // Event containing the contract specifics and raw log

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
func (it *IoTeXDIDProxyOwnershipTransferredIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(IoTeXDIDProxyOwnershipTransferred)
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
		it.Event = new(IoTeXDIDProxyOwnershipTransferred)
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
func (it *IoTeXDIDProxyOwnershipTransferredIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *IoTeXDIDProxyOwnershipTransferredIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// IoTeXDIDProxyOwnershipTransferred represents a OwnershipTransferred event raised by the IoTeXDIDProxy contract.
type IoTeXDIDProxyOwnershipTransferred struct {
	PreviousOwner common.Address
	NewOwner      common.Address
	Raw           types.Log // Blockchain specific contextual infos
}

// FilterOwnershipTransferred is a free log retrieval operation binding the contract event 0x8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0.
//
// Solidity: event OwnershipTransferred(address indexed previousOwner, address indexed newOwner)
func (_IoTeXDIDProxy *IoTeXDIDProxyFilterer) FilterOwnershipTransferred(opts *bind.FilterOpts, previousOwner []common.Address, newOwner []common.Address) (*IoTeXDIDProxyOwnershipTransferredIterator, error) {

	var previousOwnerRule []interface{}
	for _, previousOwnerItem := range previousOwner {
		previousOwnerRule = append(previousOwnerRule, previousOwnerItem)
	}
	var newOwnerRule []interface{}
	for _, newOwnerItem := range newOwner {
		newOwnerRule = append(newOwnerRule, newOwnerItem)
	}

	logs, sub, err := _IoTeXDIDProxy.contract.FilterLogs(opts, "OwnershipTransferred", previousOwnerRule, newOwnerRule)
	if err != nil {
		return nil, err
	}
	return &IoTeXDIDProxyOwnershipTransferredIterator{contract: _IoTeXDIDProxy.contract, event: "OwnershipTransferred", logs: logs, sub: sub}, nil
}

// WatchOwnershipTransferred is a free log subscription operation binding the contract event 0x8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0.
//
// Solidity: event OwnershipTransferred(address indexed previousOwner, address indexed newOwner)
func (_IoTeXDIDProxy *IoTeXDIDProxyFilterer) WatchOwnershipTransferred(opts *bind.WatchOpts, sink chan<- *IoTeXDIDProxyOwnershipTransferred, previousOwner []common.Address, newOwner []common.Address) (event.Subscription, error) {

	var previousOwnerRule []interface{}
	for _, previousOwnerItem := range previousOwner {
		previousOwnerRule = append(previousOwnerRule, previousOwnerItem)
	}
	var newOwnerRule []interface{}
	for _, newOwnerItem := range newOwner {
		newOwnerRule = append(newOwnerRule, newOwnerItem)
	}

	logs, sub, err := _IoTeXDIDProxy.contract.WatchLogs(opts, "OwnershipTransferred", previousOwnerRule, newOwnerRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(IoTeXDIDProxyOwnershipTransferred)
				if err := _IoTeXDIDProxy.contract.UnpackLog(event, "OwnershipTransferred", log); err != nil {
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
func (_IoTeXDIDProxy *IoTeXDIDProxyFilterer) ParseOwnershipTransferred(log types.Log) (*IoTeXDIDProxyOwnershipTransferred, error) {
	event := new(IoTeXDIDProxyOwnershipTransferred)
	if err := _IoTeXDIDProxy.contract.UnpackLog(event, "OwnershipTransferred", log); err != nil {
		return nil, err
	}
	return event, nil
}

// IoTeXDIDProxyUpgradeIterator is returned from FilterUpgrade and is used to iterate over the raw logs and unpacked data for Upgrade events raised by the IoTeXDIDProxy contract.
type IoTeXDIDProxyUpgradeIterator struct {
	Event *IoTeXDIDProxyUpgrade // Event containing the contract specifics and raw log

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
func (it *IoTeXDIDProxyUpgradeIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(IoTeXDIDProxyUpgrade)
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
		it.Event = new(IoTeXDIDProxyUpgrade)
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
func (it *IoTeXDIDProxyUpgradeIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *IoTeXDIDProxyUpgradeIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// IoTeXDIDProxyUpgrade represents a Upgrade event raised by the IoTeXDIDProxy contract.
type IoTeXDIDProxyUpgrade struct {
	Version string
	Addr    common.Address
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterUpgrade is a free log retrieval operation binding the contract event 0x6f48f7d6926b162cb7a11839a8d06721e35a9124ee35ee79d9faae8d34ab82ce.
//
// Solidity: event Upgrade(string version, address addr)
func (_IoTeXDIDProxy *IoTeXDIDProxyFilterer) FilterUpgrade(opts *bind.FilterOpts) (*IoTeXDIDProxyUpgradeIterator, error) {

	logs, sub, err := _IoTeXDIDProxy.contract.FilterLogs(opts, "Upgrade")
	if err != nil {
		return nil, err
	}
	return &IoTeXDIDProxyUpgradeIterator{contract: _IoTeXDIDProxy.contract, event: "Upgrade", logs: logs, sub: sub}, nil
}

// WatchUpgrade is a free log subscription operation binding the contract event 0x6f48f7d6926b162cb7a11839a8d06721e35a9124ee35ee79d9faae8d34ab82ce.
//
// Solidity: event Upgrade(string version, address addr)
func (_IoTeXDIDProxy *IoTeXDIDProxyFilterer) WatchUpgrade(opts *bind.WatchOpts, sink chan<- *IoTeXDIDProxyUpgrade) (event.Subscription, error) {

	logs, sub, err := _IoTeXDIDProxy.contract.WatchLogs(opts, "Upgrade")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(IoTeXDIDProxyUpgrade)
				if err := _IoTeXDIDProxy.contract.UnpackLog(event, "Upgrade", log); err != nil {
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

// ParseUpgrade is a log parse operation binding the contract event 0x6f48f7d6926b162cb7a11839a8d06721e35a9124ee35ee79d9faae8d34ab82ce.
//
// Solidity: event Upgrade(string version, address addr)
func (_IoTeXDIDProxy *IoTeXDIDProxyFilterer) ParseUpgrade(log types.Log) (*IoTeXDIDProxyUpgrade, error) {
	event := new(IoTeXDIDProxyUpgrade)
	if err := _IoTeXDIDProxy.contract.UnpackLog(event, "Upgrade", log); err != nil {
		return nil, err
	}
	return event, nil
}

// IoTeXDIDStorageABI is the input ABI used to generate the binding from.
const IoTeXDIDStorageABI = "[{\"constant\":true,\"inputs\":[{\"internalType\":\"string\",\"name\":\"\",\"type\":\"string\"}],\"name\":\"dids\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"exist\",\"type\":\"bool\"},{\"internalType\":\"bytes32\",\"name\":\"hash\",\"type\":\"bytes32\"},{\"internalType\":\"string\",\"name\":\"uri\",\"type\":\"string\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"}]"

// IoTeXDIDStorageFuncSigs maps the 4-byte function signature to its string representation.
var IoTeXDIDStorageFuncSigs = map[string]string{
	"f44ab516": "dids(string)",
}

// IoTeXDIDStorageBin is the compiled bytecode used for deploying new contracts.
var IoTeXDIDStorageBin = "0x608060405234801561001057600080fd5b50610253806100206000396000f3fe608060405234801561001057600080fd5b506004361061002b5760003560e01c8063f44ab51614610030575b600080fd5b6100d66004803603602081101561004657600080fd5b81019060208101813564010000000081111561006157600080fd5b82018360208201111561007357600080fd5b8035906020019184600183028401116401000000008311171561009557600080fd5b91908080601f016020809104026020016040519081016040528093929190818152602001838380828437600092019190915250929550610160945050505050565b604051808415151515815260200183815260200180602001828103825283818151815260200191508051906020019080838360005b8381101561012357818101518382015260200161010b565b50505050905090810190601f1680156101505780820380516001836020036101000a031916815260200191505b5094505050505060405180910390f35b80516020818301810180516000825292820193820193909320919092528054600180830154600280850180546040805161010096831615969096026000190190911692909204601f810188900488028501880190925281845260ff90941695919493918301828280156102145780601f106101e957610100808354040283529160200191610214565b820191906000526020600020905b8154815290600101906020018083116101f757829003601f168201915b505050505090508356fea265627a7a72315820d71850a24eab2c865e99a2ae943a6c9ccf23bf7ab1c8e3a449c8efe1d5c7122b64736f6c63430005100032"

// DeployIoTeXDIDStorage deploys a new Ethereum contract, binding an instance of IoTeXDIDStorage to it.
func DeployIoTeXDIDStorage(auth *bind.TransactOpts, backend bind.ContractBackend) (common.Address, *types.Transaction, *IoTeXDIDStorage, error) {
	parsed, err := abi.JSON(strings.NewReader(IoTeXDIDStorageABI))
	if err != nil {
		return common.Address{}, nil, nil, err
	}

	address, tx, contract, err := bind.DeployContract(auth, parsed, common.FromHex(IoTeXDIDStorageBin), backend)
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	return address, tx, &IoTeXDIDStorage{IoTeXDIDStorageCaller: IoTeXDIDStorageCaller{contract: contract}, IoTeXDIDStorageTransactor: IoTeXDIDStorageTransactor{contract: contract}, IoTeXDIDStorageFilterer: IoTeXDIDStorageFilterer{contract: contract}}, nil
}

// IoTeXDIDStorage is an auto generated Go binding around an Ethereum contract.
type IoTeXDIDStorage struct {
	IoTeXDIDStorageCaller     // Read-only binding to the contract
	IoTeXDIDStorageTransactor // Write-only binding to the contract
	IoTeXDIDStorageFilterer   // Log filterer for contract events
}

// IoTeXDIDStorageCaller is an auto generated read-only Go binding around an Ethereum contract.
type IoTeXDIDStorageCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// IoTeXDIDStorageTransactor is an auto generated write-only Go binding around an Ethereum contract.
type IoTeXDIDStorageTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// IoTeXDIDStorageFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type IoTeXDIDStorageFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// IoTeXDIDStorageSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type IoTeXDIDStorageSession struct {
	Contract     *IoTeXDIDStorage  // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// IoTeXDIDStorageCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type IoTeXDIDStorageCallerSession struct {
	Contract *IoTeXDIDStorageCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts          // Call options to use throughout this session
}

// IoTeXDIDStorageTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type IoTeXDIDStorageTransactorSession struct {
	Contract     *IoTeXDIDStorageTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts          // Transaction auth options to use throughout this session
}

// IoTeXDIDStorageRaw is an auto generated low-level Go binding around an Ethereum contract.
type IoTeXDIDStorageRaw struct {
	Contract *IoTeXDIDStorage // Generic contract binding to access the raw methods on
}

// IoTeXDIDStorageCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type IoTeXDIDStorageCallerRaw struct {
	Contract *IoTeXDIDStorageCaller // Generic read-only contract binding to access the raw methods on
}

// IoTeXDIDStorageTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type IoTeXDIDStorageTransactorRaw struct {
	Contract *IoTeXDIDStorageTransactor // Generic write-only contract binding to access the raw methods on
}

// NewIoTeXDIDStorage creates a new instance of IoTeXDIDStorage, bound to a specific deployed contract.
func NewIoTeXDIDStorage(address common.Address, backend bind.ContractBackend) (*IoTeXDIDStorage, error) {
	contract, err := bindIoTeXDIDStorage(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &IoTeXDIDStorage{IoTeXDIDStorageCaller: IoTeXDIDStorageCaller{contract: contract}, IoTeXDIDStorageTransactor: IoTeXDIDStorageTransactor{contract: contract}, IoTeXDIDStorageFilterer: IoTeXDIDStorageFilterer{contract: contract}}, nil
}

// NewIoTeXDIDStorageCaller creates a new read-only instance of IoTeXDIDStorage, bound to a specific deployed contract.
func NewIoTeXDIDStorageCaller(address common.Address, caller bind.ContractCaller) (*IoTeXDIDStorageCaller, error) {
	contract, err := bindIoTeXDIDStorage(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &IoTeXDIDStorageCaller{contract: contract}, nil
}

// NewIoTeXDIDStorageTransactor creates a new write-only instance of IoTeXDIDStorage, bound to a specific deployed contract.
func NewIoTeXDIDStorageTransactor(address common.Address, transactor bind.ContractTransactor) (*IoTeXDIDStorageTransactor, error) {
	contract, err := bindIoTeXDIDStorage(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &IoTeXDIDStorageTransactor{contract: contract}, nil
}

// NewIoTeXDIDStorageFilterer creates a new log filterer instance of IoTeXDIDStorage, bound to a specific deployed contract.
func NewIoTeXDIDStorageFilterer(address common.Address, filterer bind.ContractFilterer) (*IoTeXDIDStorageFilterer, error) {
	contract, err := bindIoTeXDIDStorage(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &IoTeXDIDStorageFilterer{contract: contract}, nil
}

// bindIoTeXDIDStorage binds a generic wrapper to an already deployed contract.
func bindIoTeXDIDStorage(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(IoTeXDIDStorageABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_IoTeXDIDStorage *IoTeXDIDStorageRaw) Call(opts *bind.CallOpts, result interface{}, method string, params ...interface{}) error {
	return _IoTeXDIDStorage.Contract.IoTeXDIDStorageCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_IoTeXDIDStorage *IoTeXDIDStorageRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _IoTeXDIDStorage.Contract.IoTeXDIDStorageTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_IoTeXDIDStorage *IoTeXDIDStorageRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _IoTeXDIDStorage.Contract.IoTeXDIDStorageTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_IoTeXDIDStorage *IoTeXDIDStorageCallerRaw) Call(opts *bind.CallOpts, result interface{}, method string, params ...interface{}) error {
	return _IoTeXDIDStorage.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_IoTeXDIDStorage *IoTeXDIDStorageTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _IoTeXDIDStorage.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_IoTeXDIDStorage *IoTeXDIDStorageTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _IoTeXDIDStorage.Contract.contract.Transact(opts, method, params...)
}

// Dids is a free data retrieval call binding the contract method 0xf44ab516.
//
// Solidity: function dids(string ) view returns(bool exist, bytes32 hash, string uri)
func (_IoTeXDIDStorage *IoTeXDIDStorageCaller) Dids(opts *bind.CallOpts, arg0 string) (struct {
	Exist bool
	Hash  [32]byte
	Uri   string
}, error) {
	ret := new(struct {
		Exist bool
		Hash  [32]byte
		Uri   string
	})
	out := ret
	err := _IoTeXDIDStorage.contract.Call(opts, out, "dids", arg0)
	return *ret, err
}

// Dids is a free data retrieval call binding the contract method 0xf44ab516.
//
// Solidity: function dids(string ) view returns(bool exist, bytes32 hash, string uri)
func (_IoTeXDIDStorage *IoTeXDIDStorageSession) Dids(arg0 string) (struct {
	Exist bool
	Hash  [32]byte
	Uri   string
}, error) {
	return _IoTeXDIDStorage.Contract.Dids(&_IoTeXDIDStorage.CallOpts, arg0)
}

// Dids is a free data retrieval call binding the contract method 0xf44ab516.
//
// Solidity: function dids(string ) view returns(bool exist, bytes32 hash, string uri)
func (_IoTeXDIDStorage *IoTeXDIDStorageCallerSession) Dids(arg0 string) (struct {
	Exist bool
	Hash  [32]byte
	Uri   string
}, error) {
	return _IoTeXDIDStorage.Contract.Dids(&_IoTeXDIDStorage.CallOpts, arg0)
}

// OwnableABI is the input ABI used to generate the binding from.
const OwnableABI = "[{\"inputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"constructor\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"previousOwner\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"newOwner\",\"type\":\"address\"}],\"name\":\"OwnershipTransferred\",\"type\":\"event\"},{\"constant\":true,\"inputs\":[],\"name\":\"owner\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"internalType\":\"address\",\"name\":\"newOwner\",\"type\":\"address\"}],\"name\":\"transferOwnership\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"}]"

// OwnableFuncSigs maps the 4-byte function signature to its string representation.
var OwnableFuncSigs = map[string]string{
	"8da5cb5b": "owner()",
	"f2fde38b": "transferOwnership(address)",
}

// OwnableBin is the compiled bytecode used for deploying new contracts.
var OwnableBin = "0x608060405234801561001057600080fd5b50600080546001600160a01b03191633179055610150806100326000396000f3fe608060405234801561001057600080fd5b50600436106100365760003560e01c80638da5cb5b1461003b578063f2fde38b1461005f575b600080fd5b610043610087565b604080516001600160a01b039092168252519081900360200190f35b6100856004803603602081101561007557600080fd5b50356001600160a01b0316610096565b005b6000546001600160a01b031681565b6000546001600160a01b031633146100ad57600080fd5b6001600160a01b0381166100c057600080fd5b600080546040516001600160a01b03808516939216917f8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e091a3600080546001600160a01b0319166001600160a01b039290921691909117905556fea265627a7a72315820c59a5fbbe60c1ce0179d594550cdfd1ab16bc39d661c5757b83ec57e1b6b4e0664736f6c63430005100032"

// DeployOwnable deploys a new Ethereum contract, binding an instance of Ownable to it.
func DeployOwnable(auth *bind.TransactOpts, backend bind.ContractBackend) (common.Address, *types.Transaction, *Ownable, error) {
	parsed, err := abi.JSON(strings.NewReader(OwnableABI))
	if err != nil {
		return common.Address{}, nil, nil, err
	}

	address, tx, contract, err := bind.DeployContract(auth, parsed, common.FromHex(OwnableBin), backend)
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	return address, tx, &Ownable{OwnableCaller: OwnableCaller{contract: contract}, OwnableTransactor: OwnableTransactor{contract: contract}, OwnableFilterer: OwnableFilterer{contract: contract}}, nil
}

// Ownable is an auto generated Go binding around an Ethereum contract.
type Ownable struct {
	OwnableCaller     // Read-only binding to the contract
	OwnableTransactor // Write-only binding to the contract
	OwnableFilterer   // Log filterer for contract events
}

// OwnableCaller is an auto generated read-only Go binding around an Ethereum contract.
type OwnableCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// OwnableTransactor is an auto generated write-only Go binding around an Ethereum contract.
type OwnableTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// OwnableFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type OwnableFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// OwnableSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type OwnableSession struct {
	Contract     *Ownable          // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// OwnableCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type OwnableCallerSession struct {
	Contract *OwnableCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts  // Call options to use throughout this session
}

// OwnableTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type OwnableTransactorSession struct {
	Contract     *OwnableTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts  // Transaction auth options to use throughout this session
}

// OwnableRaw is an auto generated low-level Go binding around an Ethereum contract.
type OwnableRaw struct {
	Contract *Ownable // Generic contract binding to access the raw methods on
}

// OwnableCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type OwnableCallerRaw struct {
	Contract *OwnableCaller // Generic read-only contract binding to access the raw methods on
}

// OwnableTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type OwnableTransactorRaw struct {
	Contract *OwnableTransactor // Generic write-only contract binding to access the raw methods on
}

// NewOwnable creates a new instance of Ownable, bound to a specific deployed contract.
func NewOwnable(address common.Address, backend bind.ContractBackend) (*Ownable, error) {
	contract, err := bindOwnable(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &Ownable{OwnableCaller: OwnableCaller{contract: contract}, OwnableTransactor: OwnableTransactor{contract: contract}, OwnableFilterer: OwnableFilterer{contract: contract}}, nil
}

// NewOwnableCaller creates a new read-only instance of Ownable, bound to a specific deployed contract.
func NewOwnableCaller(address common.Address, caller bind.ContractCaller) (*OwnableCaller, error) {
	contract, err := bindOwnable(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &OwnableCaller{contract: contract}, nil
}

// NewOwnableTransactor creates a new write-only instance of Ownable, bound to a specific deployed contract.
func NewOwnableTransactor(address common.Address, transactor bind.ContractTransactor) (*OwnableTransactor, error) {
	contract, err := bindOwnable(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &OwnableTransactor{contract: contract}, nil
}

// NewOwnableFilterer creates a new log filterer instance of Ownable, bound to a specific deployed contract.
func NewOwnableFilterer(address common.Address, filterer bind.ContractFilterer) (*OwnableFilterer, error) {
	contract, err := bindOwnable(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &OwnableFilterer{contract: contract}, nil
}

// bindOwnable binds a generic wrapper to an already deployed contract.
func bindOwnable(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(OwnableABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Ownable *OwnableRaw) Call(opts *bind.CallOpts, result interface{}, method string, params ...interface{}) error {
	return _Ownable.Contract.OwnableCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Ownable *OwnableRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Ownable.Contract.OwnableTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Ownable *OwnableRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Ownable.Contract.OwnableTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Ownable *OwnableCallerRaw) Call(opts *bind.CallOpts, result interface{}, method string, params ...interface{}) error {
	return _Ownable.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Ownable *OwnableTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Ownable.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Ownable *OwnableTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Ownable.Contract.contract.Transact(opts, method, params...)
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_Ownable *OwnableCaller) Owner(opts *bind.CallOpts) (common.Address, error) {
	var (
		ret0 = new(common.Address)
	)
	out := ret0
	err := _Ownable.contract.Call(opts, out, "owner")
	return *ret0, err
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_Ownable *OwnableSession) Owner() (common.Address, error) {
	return _Ownable.Contract.Owner(&_Ownable.CallOpts)
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_Ownable *OwnableCallerSession) Owner() (common.Address, error) {
	return _Ownable.Contract.Owner(&_Ownable.CallOpts)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_Ownable *OwnableTransactor) TransferOwnership(opts *bind.TransactOpts, newOwner common.Address) (*types.Transaction, error) {
	return _Ownable.contract.Transact(opts, "transferOwnership", newOwner)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_Ownable *OwnableSession) TransferOwnership(newOwner common.Address) (*types.Transaction, error) {
	return _Ownable.Contract.TransferOwnership(&_Ownable.TransactOpts, newOwner)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_Ownable *OwnableTransactorSession) TransferOwnership(newOwner common.Address) (*types.Transaction, error) {
	return _Ownable.Contract.TransferOwnership(&_Ownable.TransactOpts, newOwner)
}

// OwnableOwnershipTransferredIterator is returned from FilterOwnershipTransferred and is used to iterate over the raw logs and unpacked data for OwnershipTransferred events raised by the Ownable contract.
type OwnableOwnershipTransferredIterator struct {
	Event *OwnableOwnershipTransferred // Event containing the contract specifics and raw log

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
func (it *OwnableOwnershipTransferredIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(OwnableOwnershipTransferred)
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
		it.Event = new(OwnableOwnershipTransferred)
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
func (it *OwnableOwnershipTransferredIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *OwnableOwnershipTransferredIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// OwnableOwnershipTransferred represents a OwnershipTransferred event raised by the Ownable contract.
type OwnableOwnershipTransferred struct {
	PreviousOwner common.Address
	NewOwner      common.Address
	Raw           types.Log // Blockchain specific contextual infos
}

// FilterOwnershipTransferred is a free log retrieval operation binding the contract event 0x8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0.
//
// Solidity: event OwnershipTransferred(address indexed previousOwner, address indexed newOwner)
func (_Ownable *OwnableFilterer) FilterOwnershipTransferred(opts *bind.FilterOpts, previousOwner []common.Address, newOwner []common.Address) (*OwnableOwnershipTransferredIterator, error) {

	var previousOwnerRule []interface{}
	for _, previousOwnerItem := range previousOwner {
		previousOwnerRule = append(previousOwnerRule, previousOwnerItem)
	}
	var newOwnerRule []interface{}
	for _, newOwnerItem := range newOwner {
		newOwnerRule = append(newOwnerRule, newOwnerItem)
	}

	logs, sub, err := _Ownable.contract.FilterLogs(opts, "OwnershipTransferred", previousOwnerRule, newOwnerRule)
	if err != nil {
		return nil, err
	}
	return &OwnableOwnershipTransferredIterator{contract: _Ownable.contract, event: "OwnershipTransferred", logs: logs, sub: sub}, nil
}

// WatchOwnershipTransferred is a free log subscription operation binding the contract event 0x8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0.
//
// Solidity: event OwnershipTransferred(address indexed previousOwner, address indexed newOwner)
func (_Ownable *OwnableFilterer) WatchOwnershipTransferred(opts *bind.WatchOpts, sink chan<- *OwnableOwnershipTransferred, previousOwner []common.Address, newOwner []common.Address) (event.Subscription, error) {

	var previousOwnerRule []interface{}
	for _, previousOwnerItem := range previousOwner {
		previousOwnerRule = append(previousOwnerRule, previousOwnerItem)
	}
	var newOwnerRule []interface{}
	for _, newOwnerItem := range newOwner {
		newOwnerRule = append(newOwnerRule, newOwnerItem)
	}

	logs, sub, err := _Ownable.contract.WatchLogs(opts, "OwnershipTransferred", previousOwnerRule, newOwnerRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(OwnableOwnershipTransferred)
				if err := _Ownable.contract.UnpackLog(event, "OwnershipTransferred", log); err != nil {
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
func (_Ownable *OwnableFilterer) ParseOwnershipTransferred(log types.Log) (*OwnableOwnershipTransferred, error) {
	event := new(OwnableOwnershipTransferred)
	if err := _Ownable.contract.UnpackLog(event, "OwnershipTransferred", log); err != nil {
		return nil, err
	}
	return event, nil
}
