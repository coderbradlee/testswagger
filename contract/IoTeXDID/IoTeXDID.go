// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package IoTeXDID

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

// IoTeXDIDABI is the input ABI used to generate the binding from.
const IoTeXDIDABI = "[{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"string\",\"name\":\"id\",\"type\":\"string\"},{\"indexed\":false,\"internalType\":\"string\",\"name\":\"didString\",\"type\":\"string\"}],\"name\":\"CreateDID\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"string\",\"name\":\"didString\",\"type\":\"string\"}],\"name\":\"DeleteDID\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"string\",\"name\":\"didString\",\"type\":\"string\"},{\"indexed\":false,\"internalType\":\"bytes32\",\"name\":\"hash\",\"type\":\"bytes32\"}],\"name\":\"UpdateHash\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"string\",\"name\":\"didString\",\"type\":\"string\"},{\"indexed\":false,\"internalType\":\"string\",\"name\":\"uri\",\"type\":\"string\"}],\"name\":\"UpdateURI\",\"type\":\"event\"},{\"constant\":false,\"inputs\":[{\"internalType\":\"string\",\"name\":\"id\",\"type\":\"string\"},{\"internalType\":\"bytes32\",\"name\":\"hash\",\"type\":\"bytes32\"},{\"internalType\":\"string\",\"name\":\"uri\",\"type\":\"string\"}],\"name\":\"createDID\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"internalType\":\"string\",\"name\":\"did\",\"type\":\"string\"}],\"name\":\"deleteDID\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"internalType\":\"string\",\"name\":\"\",\"type\":\"string\"}],\"name\":\"dids\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"exist\",\"type\":\"bool\"},{\"internalType\":\"bytes32\",\"name\":\"hash\",\"type\":\"bytes32\"},{\"internalType\":\"string\",\"name\":\"uri\",\"type\":\"string\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"internalType\":\"string\",\"name\":\"did\",\"type\":\"string\"}],\"name\":\"getHash\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"internalType\":\"string\",\"name\":\"did\",\"type\":\"string\"}],\"name\":\"getURI\",\"outputs\":[{\"internalType\":\"string\",\"name\":\"\",\"type\":\"string\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"internalType\":\"string\",\"name\":\"did\",\"type\":\"string\"},{\"internalType\":\"bytes32\",\"name\":\"hash\",\"type\":\"bytes32\"}],\"name\":\"updateHash\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"internalType\":\"string\",\"name\":\"did\",\"type\":\"string\"},{\"internalType\":\"string\",\"name\":\"uri\",\"type\":\"string\"}],\"name\":\"updateURI\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"}]"

// IoTeXDIDFuncSigs maps the 4-byte function signature to its string representation.
var IoTeXDIDFuncSigs = map[string]string{
	"3e05e066": "createDID(string,bytes32,string)",
	"42e643bc": "deleteDID(string)",
	"f44ab516": "dids(string)",
	"5b6beeb9": "getHash(string)",
	"93ff5d3e": "getURI(string)",
	"78eab45a": "updateHash(string,bytes32)",
	"1789e2d8": "updateURI(string,string)",
}

// IoTeXDIDBin is the compiled bytecode used for deploying new contracts.
var IoTeXDIDBin = "0x608060405234801561001057600080fd5b50611943806100206000396000f3fe608060405234801561001057600080fd5b506004361061007d5760003560e01c80635b6beeb91161005b5780635b6beeb91461038257806378eab45a1461043857806393ff5d3e146104de578063f44ab516146105f75761007d565b80631789e2d8146100825780633e05e066146101ad57806342e643bc146102de575b600080fd5b6101ab6004803603604081101561009857600080fd5b810190602081018135600160201b8111156100b257600080fd5b8201836020820111156100c457600080fd5b803590602001918460018302840111600160201b831117156100e557600080fd5b91908080601f0160208091040260200160405190810160405280939291908181526020018383808284376000920191909152509295949360208101935035915050600160201b81111561013757600080fd5b82018360208201111561014957600080fd5b803590602001918460018302840111600160201b8311171561016a57600080fd5b91908080601f016020809104026020016040519081016040528093929190818152602001838380828437600092019190915250929550610725945050505050565b005b6101ab600480360360608110156101c357600080fd5b810190602081018135600160201b8111156101dd57600080fd5b8201836020820111156101ef57600080fd5b803590602001918460018302840111600160201b8311171561021057600080fd5b91908080601f01602080910402602001604051908101604052809392919081815260200183838082843760009201919091525092958435959094909350604081019250602001359050600160201b81111561026a57600080fd5b82018360208201111561027c57600080fd5b803590602001918460018302840111600160201b8311171561029d57600080fd5b91908080601f0160208091040260200160405190810160405280939291908181526020018383808284376000920191909152509295506109ab945050505050565b6101ab600480360360208110156102f457600080fd5b810190602081018135600160201b81111561030e57600080fd5b82018360208201111561032057600080fd5b803590602001918460018302840111600160201b8311171561034157600080fd5b91908080601f016020809104026020016040519081016040528093929190818152602001838380828437600092019190915250929550610c3e945050505050565b6104266004803603602081101561039857600080fd5b810190602081018135600160201b8111156103b257600080fd5b8201836020820111156103c457600080fd5b803590602001918460018302840111600160201b831117156103e557600080fd5b91908080601f016020809104026020016040519081016040528093929190818152602001838380828437600092019190915250929550610e51945050505050565b60408051918252519081900360200190f35b6101ab6004803603604081101561044e57600080fd5b810190602081018135600160201b81111561046857600080fd5b82018360208201111561047a57600080fd5b803590602001918460018302840111600160201b8311171561049b57600080fd5b91908080601f0160208091040260200160405190810160405280939291908181526020018383808284376000920191909152509295505091359250610f75915050565b610582600480360360208110156104f457600080fd5b810190602081018135600160201b81111561050e57600080fd5b82018360208201111561052057600080fd5b803590602001918460018302840111600160201b8311171561054157600080fd5b91908080601f016020809104026020016040519081016040528093929190818152602001838380828437600092019190915250929550611188945050505050565b6040805160208082528351818301528351919283929083019185019080838360005b838110156105bc5781810151838201526020016105a4565b50505050905090810190601f1680156105e95780820380516001836020036101000a031916815260200191505b509250505060405180910390f35b61069b6004803603602081101561060d57600080fd5b810190602081018135600160201b81111561062757600080fd5b82018360208201111561063957600080fd5b803590602001918460018302840111600160201b8311171561065a57600080fd5b91908080601f016020809104026020016040519081016040528093929190818152602001838380828437600092019190915250929550611337945050505050565b604051808415151515815260200183815260200180602001828103825283818151815260200191508051906020019080838360005b838110156106e85781810151838201526020016106d0565b50505050905090810190601f1680156107155780820380516001836020036101000a031916815260200191505b5094505050505060405180910390f35b8160606107306113f5565b82519091501561077f5761074482826114d5565b61077f5760405162461bcd60e51b81526004018080602001828103825260218152602001806118ee6021913960400191505060405180910390fd5b6000816040518082805190602001908083835b602083106107b15780518252601f199092019160209182019101610792565b51815160209384036101000a600019018019909216911617905292019485525060405193849003019092205460ff1691506108319050576040805162461bcd60e51b815260206004820152601960248201527831b0b63632b91034b9903737ba1030903234b21037bbb732b960391b604482015290519081900360640190fd5b82600061083c6113f5565b6040518082805190602001908083835b6020831061086b5780518252601f19909201916020918201910161084c565b51815160209384036101000a600019018019909216911617905292019485525060405193849003810190932084516108b0956002909201949190910192509050611855565b506108b96113f5565b6040518082805190602001908083835b602083106108e85780518252601f1990920191602091820191016108c9565b51815160209384036101000a60001901801990921691161790526040805192909401829003822081835289518383015289519096507fa55ea16d9a68fded3c5b007d1fbba6ed1524d9354d44c66bcc54c721108586c495508994929350839283019185019080838360005b8381101561096b578181015183820152602001610953565b50505050905090810190601f1680156109985780820380516001836020036101000a031916815260200191505b509250505060405180910390a250505050565b825115610a15576109c4836109bf336115cc565b6114d5565b610a15576040805162461bcd60e51b815260206004820152601960248201527f696420646f6573206e6f74206d617463682063726561746f7200000000000000604482015290519081900360640190fd5b6060610a1f6113f5565b90506000816040518082805190602001908083835b60208310610a535780518252601f199092019160209182019101610a34565b51815160209384036101000a600019018019909216911617905292019485525060405193849003019092205460ff16159150610acd9050576040805162461bcd60e51b815260206004820152601260248201527164696420616c72656164792065786973747360701b604482015290519081900360640190fd5b6040518060600160405280600115158152602001848152602001838152506000826040518082805190602001908083835b60208310610b1d5780518252601f199092019160209182019101610afe565b51815160209384036101000a6000190180199092169116179052920194855250604080519485900382019094208551815460ff1916901515178155858201516001820155938501518051610b7a9450600286019350910190611855565b50905050610b8f610b8a336115cc565b611742565b6040518082805190602001908083835b60208310610bbe5780518252601f199092019160209182019101610b9f565b51815160209384036101000a60001901801990921691161790526040805192909401829003822081835287518383015287519096507fedcad282d2fc969322aa68caeb2b9bf9de762094ec288a9965535b8968dca17f9550879492935083928301918501908083836000831561096b578181015183820152602001610953565b806060610c496113f5565b825190915015610c9857610c5d82826114d5565b610c985760405162461bcd60e51b81526004018080602001828103825260218152602001806118ee6021913960400191505060405180910390fd5b6000816040518082805190602001908083835b60208310610cca5780518252601f199092019160209182019101610cab565b51815160209384036101000a600019018019909216911617905292019485525060405193849003019092205460ff169150610d4a9050576040805162461bcd60e51b815260206004820152601960248201527831b0b63632b91034b9903737ba1030903234b21037bbb732b960391b604482015290519081900360640190fd5b600080610d556113f5565b6040518082805190602001908083835b60208310610d845780518252601f199092019160209182019101610d65565b51815160209384036101000a60001901801990921691161790529201948552506040519384900301909220805460ff19169315159390931790925550610dca90506113f5565b6040518082805190602001908083835b60208310610df95780518252601f199092019160209182019101610dda565b5181516020939093036101000a60001901801990911692169190911790526040519201829003822093507ff8910781c5d425c8a634b0a60a6567ef96cae8b2aa97b0db44389fff2bbcc64f92506000919050a2505050565b60006060610e5e83611742565b90506000816040518082805190602001908083835b60208310610e925780518252601f199092019160209182019101610e73565b51815160209384036101000a600019018019909216911617905292019485525060405193849003019092205460ff169150610f0b9050576040805162461bcd60e51b8152602060048201526012602482015271191a5908191bd95cc81b9bdd08195e1a5cdd60721b604482015290519081900360640190fd5b6000816040518082805190602001908083835b60208310610f3d5780518252601f199092019160209182019101610f1e565b51815160209384036101000a600019018019909216911617905292019485525060405193849003019092206001015495945050505050565b816060610f806113f5565b825190915015610fcf57610f9482826114d5565b610fcf5760405162461bcd60e51b81526004018080602001828103825260218152602001806118ee6021913960400191505060405180910390fd5b6000816040518082805190602001908083835b602083106110015780518252601f199092019160209182019101610fe2565b51815160209384036101000a600019018019909216911617905292019485525060405193849003019092205460ff1691506110819050576040805162461bcd60e51b815260206004820152601960248201527831b0b63632b91034b9903737ba1030903234b21037bbb732b960391b604482015290519081900360640190fd5b82600061108c6113f5565b6040518082805190602001908083835b602083106110bb5780518252601f19909201916020918201910161109c565b51815160209384036101000a6000190180199092169116179052920194855250604051938490030190922060010192909255506110f890506113f5565b6040518082805190602001908083835b602083106111275780518252601f199092019160209182019101611108565b51815160209384036101000a60001901801990921691161790526040805192909401829003822089835293519395507fad2d38b24c03f1f0db04b9345ca745dcf8904f27f8977599ec4769743d5e898d94509083900301919050a250505050565b60608061119483611742565b90506000816040518082805190602001908083835b602083106111c85780518252601f1990920191602091820191016111a9565b51815160209384036101000a600019018019909216911617905292019485525060405193849003019092205460ff1691506112419050576040805162461bcd60e51b8152602060048201526012602482015271191a5908191bd95cc81b9bdd08195e1a5cdd60721b604482015290519081900360640190fd5b6000816040518082805190602001908083835b602083106112735780518252601f199092019160209182019101611254565b518151600019602094850361010090810a8201928316921993909316919091179092529490920196875260408051978890038201882060029081018054601f60018216159098029095019094160494850182900482028801820190528387529094509192505083018282801561132a5780601f106112ff5761010080835404028352916020019161132a565b820191906000526020600020905b81548152906001019060200180831161130d57829003601f168201915b5050505050915050919050565b80516020818301810180516000825292820193820193909320919092528054600180830154600280850180546040805161010096831615969096026000190190911692909204601f810188900488028501880190925281845260ff90941695919493918301828280156113eb5780601f106113c0576101008083540402835291602001916113eb565b820191906000526020600020905b8154815290600101906020018083116113ce57829003601f168201915b5050505050905083565b6060604051806040016040528060078152602001663234b21d34b79d60c91b815250611420336115cc565b6040516020018083805190602001908083835b602083106114525780518252601f199092019160209182019101611433565b51815160209384036101000a600019018019909216911617905285519190930192850191508083835b6020831061149a5780518252601f19909201916020918201910161147b565b6001836020036101000a0380198251168184511680821785525050505050509050019250505060405160208183030381529060405290505b90565b60006114e082611742565b6040516020018082805190602001908083835b602083106115125780518252601f1990920191602091820191016114f3565b6001836020036101000a0380198251168184511680821785525050505050509050019150506040516020818303038152906040528051906020012061155684611742565b6040516020018082805190602001908083835b602083106115885780518252601f199092019160209182019101611569565b6001836020036101000a0380198251168184511680821785525050505050509050019150506040516020818303038152906040528051906020012014905092915050565b604080518082018252601081526f181899199a1a9b1b9c1cb0b131b232b360811b60208201528151602a80825260608281019094526001600160a01b03851692918491602082018180388339019050509050600360fc1b8160008151811061163057fe5b60200101906001600160f81b031916908160001a905350600f60fb1b8160018151811061165957fe5b60200101906001600160f81b031916908160001a90535060005b6014811015611739578260048583600c016020811061168e57fe5b835191901a90911c60ff169081106116a257fe5b602001015160f81c60f81b8282600202600201815181106116bf57fe5b60200101906001600160f81b031916908160001a905350828482600c01602081106116e657fe5b1a60f81b60f81c600f1660ff16815181106116fd57fe5b602001015160f81c60f81b82826002026003018151811061171a57fe5b60200101906001600160f81b031916908160001a905350600101611673565b50949350505050565b606080829050606081516040519080825280601f01601f191660200182016040528015611776576020820181803883390190505b50905060005b825181101561184d57604183828151811061179357fe5b016020015160f81c108015906117bd5750605a8382815181106117b257fe5b016020015160f81c11155b1561180a578281815181106117ce57fe5b602001015160f81c60f81b60f81c60200160f81b8282815181106117ee57fe5b60200101906001600160f81b031916908160001a905350611845565b82818151811061181657fe5b602001015160f81c60f81b82828151811061182d57fe5b60200101906001600160f81b031916908160001a9053505b60010161177c565b509392505050565b828054600181600116156101000203166002900490600052602060002090601f016020900481019282601f1061189657805160ff19168380011785556118c3565b828001600101855582156118c3579182015b828111156118c35782518255916020019190600101906118a8565b506118cf9291506118d3565b5090565b6114d291905b808211156118cf57600081556001016118d956fe63616c6c657220646f6573206e6f74206f776e2074686520676976656e20646964a265627a7a72315820f1717a95aa86ec9d954d349e6b0c374444b222db3b744750f219b031e519ecf164736f6c63430005100032"

// DeployIoTeXDID deploys a new Ethereum contract, binding an instance of IoTeXDID to it.
func DeployIoTeXDID(auth *bind.TransactOpts, backend bind.ContractBackend) (common.Address, *types.Transaction, *IoTeXDID, error) {
	parsed, err := abi.JSON(strings.NewReader(IoTeXDIDABI))
	if err != nil {
		return common.Address{}, nil, nil, err
	}

	address, tx, contract, err := bind.DeployContract(auth, parsed, common.FromHex(IoTeXDIDBin), backend)
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	return address, tx, &IoTeXDID{IoTeXDIDCaller: IoTeXDIDCaller{contract: contract}, IoTeXDIDTransactor: IoTeXDIDTransactor{contract: contract}, IoTeXDIDFilterer: IoTeXDIDFilterer{contract: contract}}, nil
}

// IoTeXDID is an auto generated Go binding around an Ethereum contract.
type IoTeXDID struct {
	IoTeXDIDCaller     // Read-only binding to the contract
	IoTeXDIDTransactor // Write-only binding to the contract
	IoTeXDIDFilterer   // Log filterer for contract events
}

// IoTeXDIDCaller is an auto generated read-only Go binding around an Ethereum contract.
type IoTeXDIDCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// IoTeXDIDTransactor is an auto generated write-only Go binding around an Ethereum contract.
type IoTeXDIDTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// IoTeXDIDFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type IoTeXDIDFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// IoTeXDIDSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type IoTeXDIDSession struct {
	Contract     *IoTeXDID         // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// IoTeXDIDCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type IoTeXDIDCallerSession struct {
	Contract *IoTeXDIDCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts   // Call options to use throughout this session
}

// IoTeXDIDTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type IoTeXDIDTransactorSession struct {
	Contract     *IoTeXDIDTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts   // Transaction auth options to use throughout this session
}

// IoTeXDIDRaw is an auto generated low-level Go binding around an Ethereum contract.
type IoTeXDIDRaw struct {
	Contract *IoTeXDID // Generic contract binding to access the raw methods on
}

// IoTeXDIDCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type IoTeXDIDCallerRaw struct {
	Contract *IoTeXDIDCaller // Generic read-only contract binding to access the raw methods on
}

// IoTeXDIDTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type IoTeXDIDTransactorRaw struct {
	Contract *IoTeXDIDTransactor // Generic write-only contract binding to access the raw methods on
}

// NewIoTeXDID creates a new instance of IoTeXDID, bound to a specific deployed contract.
func NewIoTeXDID(address common.Address, backend bind.ContractBackend) (*IoTeXDID, error) {
	contract, err := bindIoTeXDID(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &IoTeXDID{IoTeXDIDCaller: IoTeXDIDCaller{contract: contract}, IoTeXDIDTransactor: IoTeXDIDTransactor{contract: contract}, IoTeXDIDFilterer: IoTeXDIDFilterer{contract: contract}}, nil
}

// NewIoTeXDIDCaller creates a new read-only instance of IoTeXDID, bound to a specific deployed contract.
func NewIoTeXDIDCaller(address common.Address, caller bind.ContractCaller) (*IoTeXDIDCaller, error) {
	contract, err := bindIoTeXDID(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &IoTeXDIDCaller{contract: contract}, nil
}

// NewIoTeXDIDTransactor creates a new write-only instance of IoTeXDID, bound to a specific deployed contract.
func NewIoTeXDIDTransactor(address common.Address, transactor bind.ContractTransactor) (*IoTeXDIDTransactor, error) {
	contract, err := bindIoTeXDID(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &IoTeXDIDTransactor{contract: contract}, nil
}

// NewIoTeXDIDFilterer creates a new log filterer instance of IoTeXDID, bound to a specific deployed contract.
func NewIoTeXDIDFilterer(address common.Address, filterer bind.ContractFilterer) (*IoTeXDIDFilterer, error) {
	contract, err := bindIoTeXDID(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &IoTeXDIDFilterer{contract: contract}, nil
}

// bindIoTeXDID binds a generic wrapper to an already deployed contract.
func bindIoTeXDID(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(IoTeXDIDABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_IoTeXDID *IoTeXDIDRaw) Call(opts *bind.CallOpts, result interface{}, method string, params ...interface{}) error {
	return _IoTeXDID.Contract.IoTeXDIDCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_IoTeXDID *IoTeXDIDRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _IoTeXDID.Contract.IoTeXDIDTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_IoTeXDID *IoTeXDIDRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _IoTeXDID.Contract.IoTeXDIDTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_IoTeXDID *IoTeXDIDCallerRaw) Call(opts *bind.CallOpts, result interface{}, method string, params ...interface{}) error {
	return _IoTeXDID.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_IoTeXDID *IoTeXDIDTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _IoTeXDID.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_IoTeXDID *IoTeXDIDTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _IoTeXDID.Contract.contract.Transact(opts, method, params...)
}

// Dids is a free data retrieval call binding the contract method 0xf44ab516.
//
// Solidity: function dids(string ) view returns(bool exist, bytes32 hash, string uri)
func (_IoTeXDID *IoTeXDIDCaller) Dids(opts *bind.CallOpts, arg0 string) (struct {
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
	err := _IoTeXDID.contract.Call(opts, out, "dids", arg0)
	return *ret, err
}

// Dids is a free data retrieval call binding the contract method 0xf44ab516.
//
// Solidity: function dids(string ) view returns(bool exist, bytes32 hash, string uri)
func (_IoTeXDID *IoTeXDIDSession) Dids(arg0 string) (struct {
	Exist bool
	Hash  [32]byte
	Uri   string
}, error) {
	return _IoTeXDID.Contract.Dids(&_IoTeXDID.CallOpts, arg0)
}

// Dids is a free data retrieval call binding the contract method 0xf44ab516.
//
// Solidity: function dids(string ) view returns(bool exist, bytes32 hash, string uri)
func (_IoTeXDID *IoTeXDIDCallerSession) Dids(arg0 string) (struct {
	Exist bool
	Hash  [32]byte
	Uri   string
}, error) {
	return _IoTeXDID.Contract.Dids(&_IoTeXDID.CallOpts, arg0)
}

// GetHash is a free data retrieval call binding the contract method 0x5b6beeb9.
//
// Solidity: function getHash(string did) view returns(bytes32)
func (_IoTeXDID *IoTeXDIDCaller) GetHash(opts *bind.CallOpts, did string) ([32]byte, error) {
	var (
		ret0 = new([32]byte)
	)
	out := ret0
	err := _IoTeXDID.contract.Call(opts, out, "getHash", did)
	return *ret0, err
}

// GetHash is a free data retrieval call binding the contract method 0x5b6beeb9.
//
// Solidity: function getHash(string did) view returns(bytes32)
func (_IoTeXDID *IoTeXDIDSession) GetHash(did string) ([32]byte, error) {
	return _IoTeXDID.Contract.GetHash(&_IoTeXDID.CallOpts, did)
}

// GetHash is a free data retrieval call binding the contract method 0x5b6beeb9.
//
// Solidity: function getHash(string did) view returns(bytes32)
func (_IoTeXDID *IoTeXDIDCallerSession) GetHash(did string) ([32]byte, error) {
	return _IoTeXDID.Contract.GetHash(&_IoTeXDID.CallOpts, did)
}

// GetURI is a free data retrieval call binding the contract method 0x93ff5d3e.
//
// Solidity: function getURI(string did) view returns(string)
func (_IoTeXDID *IoTeXDIDCaller) GetURI(opts *bind.CallOpts, did string) (string, error) {
	var (
		ret0 = new(string)
	)
	out := ret0
	err := _IoTeXDID.contract.Call(opts, out, "getURI", did)
	return *ret0, err
}

// GetURI is a free data retrieval call binding the contract method 0x93ff5d3e.
//
// Solidity: function getURI(string did) view returns(string)
func (_IoTeXDID *IoTeXDIDSession) GetURI(did string) (string, error) {
	return _IoTeXDID.Contract.GetURI(&_IoTeXDID.CallOpts, did)
}

// GetURI is a free data retrieval call binding the contract method 0x93ff5d3e.
//
// Solidity: function getURI(string did) view returns(string)
func (_IoTeXDID *IoTeXDIDCallerSession) GetURI(did string) (string, error) {
	return _IoTeXDID.Contract.GetURI(&_IoTeXDID.CallOpts, did)
}

// CreateDID is a paid mutator transaction binding the contract method 0x3e05e066.
//
// Solidity: function createDID(string id, bytes32 hash, string uri) returns()
func (_IoTeXDID *IoTeXDIDTransactor) CreateDID(opts *bind.TransactOpts, id string, hash [32]byte, uri string) (*types.Transaction, error) {
	return _IoTeXDID.contract.Transact(opts, "createDID", id, hash, uri)
}

// CreateDID is a paid mutator transaction binding the contract method 0x3e05e066.
//
// Solidity: function createDID(string id, bytes32 hash, string uri) returns()
func (_IoTeXDID *IoTeXDIDSession) CreateDID(id string, hash [32]byte, uri string) (*types.Transaction, error) {
	return _IoTeXDID.Contract.CreateDID(&_IoTeXDID.TransactOpts, id, hash, uri)
}

// CreateDID is a paid mutator transaction binding the contract method 0x3e05e066.
//
// Solidity: function createDID(string id, bytes32 hash, string uri) returns()
func (_IoTeXDID *IoTeXDIDTransactorSession) CreateDID(id string, hash [32]byte, uri string) (*types.Transaction, error) {
	return _IoTeXDID.Contract.CreateDID(&_IoTeXDID.TransactOpts, id, hash, uri)
}

// DeleteDID is a paid mutator transaction binding the contract method 0x42e643bc.
//
// Solidity: function deleteDID(string did) returns()
func (_IoTeXDID *IoTeXDIDTransactor) DeleteDID(opts *bind.TransactOpts, did string) (*types.Transaction, error) {
	return _IoTeXDID.contract.Transact(opts, "deleteDID", did)
}

// DeleteDID is a paid mutator transaction binding the contract method 0x42e643bc.
//
// Solidity: function deleteDID(string did) returns()
func (_IoTeXDID *IoTeXDIDSession) DeleteDID(did string) (*types.Transaction, error) {
	return _IoTeXDID.Contract.DeleteDID(&_IoTeXDID.TransactOpts, did)
}

// DeleteDID is a paid mutator transaction binding the contract method 0x42e643bc.
//
// Solidity: function deleteDID(string did) returns()
func (_IoTeXDID *IoTeXDIDTransactorSession) DeleteDID(did string) (*types.Transaction, error) {
	return _IoTeXDID.Contract.DeleteDID(&_IoTeXDID.TransactOpts, did)
}

// UpdateHash is a paid mutator transaction binding the contract method 0x78eab45a.
//
// Solidity: function updateHash(string did, bytes32 hash) returns()
func (_IoTeXDID *IoTeXDIDTransactor) UpdateHash(opts *bind.TransactOpts, did string, hash [32]byte) (*types.Transaction, error) {
	return _IoTeXDID.contract.Transact(opts, "updateHash", did, hash)
}

// UpdateHash is a paid mutator transaction binding the contract method 0x78eab45a.
//
// Solidity: function updateHash(string did, bytes32 hash) returns()
func (_IoTeXDID *IoTeXDIDSession) UpdateHash(did string, hash [32]byte) (*types.Transaction, error) {
	return _IoTeXDID.Contract.UpdateHash(&_IoTeXDID.TransactOpts, did, hash)
}

// UpdateHash is a paid mutator transaction binding the contract method 0x78eab45a.
//
// Solidity: function updateHash(string did, bytes32 hash) returns()
func (_IoTeXDID *IoTeXDIDTransactorSession) UpdateHash(did string, hash [32]byte) (*types.Transaction, error) {
	return _IoTeXDID.Contract.UpdateHash(&_IoTeXDID.TransactOpts, did, hash)
}

// UpdateURI is a paid mutator transaction binding the contract method 0x1789e2d8.
//
// Solidity: function updateURI(string did, string uri) returns()
func (_IoTeXDID *IoTeXDIDTransactor) UpdateURI(opts *bind.TransactOpts, did string, uri string) (*types.Transaction, error) {
	return _IoTeXDID.contract.Transact(opts, "updateURI", did, uri)
}

// UpdateURI is a paid mutator transaction binding the contract method 0x1789e2d8.
//
// Solidity: function updateURI(string did, string uri) returns()
func (_IoTeXDID *IoTeXDIDSession) UpdateURI(did string, uri string) (*types.Transaction, error) {
	return _IoTeXDID.Contract.UpdateURI(&_IoTeXDID.TransactOpts, did, uri)
}

// UpdateURI is a paid mutator transaction binding the contract method 0x1789e2d8.
//
// Solidity: function updateURI(string did, string uri) returns()
func (_IoTeXDID *IoTeXDIDTransactorSession) UpdateURI(did string, uri string) (*types.Transaction, error) {
	return _IoTeXDID.Contract.UpdateURI(&_IoTeXDID.TransactOpts, did, uri)
}

// IoTeXDIDCreateDIDIterator is returned from FilterCreateDID and is used to iterate over the raw logs and unpacked data for CreateDID events raised by the IoTeXDID contract.
type IoTeXDIDCreateDIDIterator struct {
	Event *IoTeXDIDCreateDID // Event containing the contract specifics and raw log

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
func (it *IoTeXDIDCreateDIDIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(IoTeXDIDCreateDID)
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
		it.Event = new(IoTeXDIDCreateDID)
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
func (it *IoTeXDIDCreateDIDIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *IoTeXDIDCreateDIDIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// IoTeXDIDCreateDID represents a CreateDID event raised by the IoTeXDID contract.
type IoTeXDIDCreateDID struct {
	Id        common.Hash
	DidString string
	Raw       types.Log // Blockchain specific contextual infos
}

// FilterCreateDID is a free log retrieval operation binding the contract event 0xedcad282d2fc969322aa68caeb2b9bf9de762094ec288a9965535b8968dca17f.
//
// Solidity: event CreateDID(string indexed id, string didString)
func (_IoTeXDID *IoTeXDIDFilterer) FilterCreateDID(opts *bind.FilterOpts, id []string) (*IoTeXDIDCreateDIDIterator, error) {

	var idRule []interface{}
	for _, idItem := range id {
		idRule = append(idRule, idItem)
	}

	logs, sub, err := _IoTeXDID.contract.FilterLogs(opts, "CreateDID", idRule)
	if err != nil {
		return nil, err
	}
	return &IoTeXDIDCreateDIDIterator{contract: _IoTeXDID.contract, event: "CreateDID", logs: logs, sub: sub}, nil
}

// WatchCreateDID is a free log subscription operation binding the contract event 0xedcad282d2fc969322aa68caeb2b9bf9de762094ec288a9965535b8968dca17f.
//
// Solidity: event CreateDID(string indexed id, string didString)
func (_IoTeXDID *IoTeXDIDFilterer) WatchCreateDID(opts *bind.WatchOpts, sink chan<- *IoTeXDIDCreateDID, id []string) (event.Subscription, error) {

	var idRule []interface{}
	for _, idItem := range id {
		idRule = append(idRule, idItem)
	}

	logs, sub, err := _IoTeXDID.contract.WatchLogs(opts, "CreateDID", idRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(IoTeXDIDCreateDID)
				if err := _IoTeXDID.contract.UnpackLog(event, "CreateDID", log); err != nil {
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

// ParseCreateDID is a log parse operation binding the contract event 0xedcad282d2fc969322aa68caeb2b9bf9de762094ec288a9965535b8968dca17f.
//
// Solidity: event CreateDID(string indexed id, string didString)
func (_IoTeXDID *IoTeXDIDFilterer) ParseCreateDID(log types.Log) (*IoTeXDIDCreateDID, error) {
	event := new(IoTeXDIDCreateDID)
	if err := _IoTeXDID.contract.UnpackLog(event, "CreateDID", log); err != nil {
		return nil, err
	}
	return event, nil
}

// IoTeXDIDDeleteDIDIterator is returned from FilterDeleteDID and is used to iterate over the raw logs and unpacked data for DeleteDID events raised by the IoTeXDID contract.
type IoTeXDIDDeleteDIDIterator struct {
	Event *IoTeXDIDDeleteDID // Event containing the contract specifics and raw log

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
func (it *IoTeXDIDDeleteDIDIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(IoTeXDIDDeleteDID)
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
		it.Event = new(IoTeXDIDDeleteDID)
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
func (it *IoTeXDIDDeleteDIDIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *IoTeXDIDDeleteDIDIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// IoTeXDIDDeleteDID represents a DeleteDID event raised by the IoTeXDID contract.
type IoTeXDIDDeleteDID struct {
	DidString common.Hash
	Raw       types.Log // Blockchain specific contextual infos
}

// FilterDeleteDID is a free log retrieval operation binding the contract event 0xf8910781c5d425c8a634b0a60a6567ef96cae8b2aa97b0db44389fff2bbcc64f.
//
// Solidity: event DeleteDID(string indexed didString)
func (_IoTeXDID *IoTeXDIDFilterer) FilterDeleteDID(opts *bind.FilterOpts, didString []string) (*IoTeXDIDDeleteDIDIterator, error) {

	var didStringRule []interface{}
	for _, didStringItem := range didString {
		didStringRule = append(didStringRule, didStringItem)
	}

	logs, sub, err := _IoTeXDID.contract.FilterLogs(opts, "DeleteDID", didStringRule)
	if err != nil {
		return nil, err
	}
	return &IoTeXDIDDeleteDIDIterator{contract: _IoTeXDID.contract, event: "DeleteDID", logs: logs, sub: sub}, nil
}

// WatchDeleteDID is a free log subscription operation binding the contract event 0xf8910781c5d425c8a634b0a60a6567ef96cae8b2aa97b0db44389fff2bbcc64f.
//
// Solidity: event DeleteDID(string indexed didString)
func (_IoTeXDID *IoTeXDIDFilterer) WatchDeleteDID(opts *bind.WatchOpts, sink chan<- *IoTeXDIDDeleteDID, didString []string) (event.Subscription, error) {

	var didStringRule []interface{}
	for _, didStringItem := range didString {
		didStringRule = append(didStringRule, didStringItem)
	}

	logs, sub, err := _IoTeXDID.contract.WatchLogs(opts, "DeleteDID", didStringRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(IoTeXDIDDeleteDID)
				if err := _IoTeXDID.contract.UnpackLog(event, "DeleteDID", log); err != nil {
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

// ParseDeleteDID is a log parse operation binding the contract event 0xf8910781c5d425c8a634b0a60a6567ef96cae8b2aa97b0db44389fff2bbcc64f.
//
// Solidity: event DeleteDID(string indexed didString)
func (_IoTeXDID *IoTeXDIDFilterer) ParseDeleteDID(log types.Log) (*IoTeXDIDDeleteDID, error) {
	event := new(IoTeXDIDDeleteDID)
	if err := _IoTeXDID.contract.UnpackLog(event, "DeleteDID", log); err != nil {
		return nil, err
	}
	return event, nil
}

// IoTeXDIDUpdateHashIterator is returned from FilterUpdateHash and is used to iterate over the raw logs and unpacked data for UpdateHash events raised by the IoTeXDID contract.
type IoTeXDIDUpdateHashIterator struct {
	Event *IoTeXDIDUpdateHash // Event containing the contract specifics and raw log

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
func (it *IoTeXDIDUpdateHashIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(IoTeXDIDUpdateHash)
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
		it.Event = new(IoTeXDIDUpdateHash)
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
func (it *IoTeXDIDUpdateHashIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *IoTeXDIDUpdateHashIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// IoTeXDIDUpdateHash represents a UpdateHash event raised by the IoTeXDID contract.
type IoTeXDIDUpdateHash struct {
	DidString common.Hash
	Hash      [32]byte
	Raw       types.Log // Blockchain specific contextual infos
}

// FilterUpdateHash is a free log retrieval operation binding the contract event 0xad2d38b24c03f1f0db04b9345ca745dcf8904f27f8977599ec4769743d5e898d.
//
// Solidity: event UpdateHash(string indexed didString, bytes32 hash)
func (_IoTeXDID *IoTeXDIDFilterer) FilterUpdateHash(opts *bind.FilterOpts, didString []string) (*IoTeXDIDUpdateHashIterator, error) {

	var didStringRule []interface{}
	for _, didStringItem := range didString {
		didStringRule = append(didStringRule, didStringItem)
	}

	logs, sub, err := _IoTeXDID.contract.FilterLogs(opts, "UpdateHash", didStringRule)
	if err != nil {
		return nil, err
	}
	return &IoTeXDIDUpdateHashIterator{contract: _IoTeXDID.contract, event: "UpdateHash", logs: logs, sub: sub}, nil
}

// WatchUpdateHash is a free log subscription operation binding the contract event 0xad2d38b24c03f1f0db04b9345ca745dcf8904f27f8977599ec4769743d5e898d.
//
// Solidity: event UpdateHash(string indexed didString, bytes32 hash)
func (_IoTeXDID *IoTeXDIDFilterer) WatchUpdateHash(opts *bind.WatchOpts, sink chan<- *IoTeXDIDUpdateHash, didString []string) (event.Subscription, error) {

	var didStringRule []interface{}
	for _, didStringItem := range didString {
		didStringRule = append(didStringRule, didStringItem)
	}

	logs, sub, err := _IoTeXDID.contract.WatchLogs(opts, "UpdateHash", didStringRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(IoTeXDIDUpdateHash)
				if err := _IoTeXDID.contract.UnpackLog(event, "UpdateHash", log); err != nil {
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

// ParseUpdateHash is a log parse operation binding the contract event 0xad2d38b24c03f1f0db04b9345ca745dcf8904f27f8977599ec4769743d5e898d.
//
// Solidity: event UpdateHash(string indexed didString, bytes32 hash)
func (_IoTeXDID *IoTeXDIDFilterer) ParseUpdateHash(log types.Log) (*IoTeXDIDUpdateHash, error) {
	event := new(IoTeXDIDUpdateHash)
	if err := _IoTeXDID.contract.UnpackLog(event, "UpdateHash", log); err != nil {
		return nil, err
	}
	return event, nil
}

// IoTeXDIDUpdateURIIterator is returned from FilterUpdateURI and is used to iterate over the raw logs and unpacked data for UpdateURI events raised by the IoTeXDID contract.
type IoTeXDIDUpdateURIIterator struct {
	Event *IoTeXDIDUpdateURI // Event containing the contract specifics and raw log

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
func (it *IoTeXDIDUpdateURIIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(IoTeXDIDUpdateURI)
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
		it.Event = new(IoTeXDIDUpdateURI)
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
func (it *IoTeXDIDUpdateURIIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *IoTeXDIDUpdateURIIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// IoTeXDIDUpdateURI represents a UpdateURI event raised by the IoTeXDID contract.
type IoTeXDIDUpdateURI struct {
	DidString common.Hash
	Uri       string
	Raw       types.Log // Blockchain specific contextual infos
}

// FilterUpdateURI is a free log retrieval operation binding the contract event 0xa55ea16d9a68fded3c5b007d1fbba6ed1524d9354d44c66bcc54c721108586c4.
//
// Solidity: event UpdateURI(string indexed didString, string uri)
func (_IoTeXDID *IoTeXDIDFilterer) FilterUpdateURI(opts *bind.FilterOpts, didString []string) (*IoTeXDIDUpdateURIIterator, error) {

	var didStringRule []interface{}
	for _, didStringItem := range didString {
		didStringRule = append(didStringRule, didStringItem)
	}

	logs, sub, err := _IoTeXDID.contract.FilterLogs(opts, "UpdateURI", didStringRule)
	if err != nil {
		return nil, err
	}
	return &IoTeXDIDUpdateURIIterator{contract: _IoTeXDID.contract, event: "UpdateURI", logs: logs, sub: sub}, nil
}

// WatchUpdateURI is a free log subscription operation binding the contract event 0xa55ea16d9a68fded3c5b007d1fbba6ed1524d9354d44c66bcc54c721108586c4.
//
// Solidity: event UpdateURI(string indexed didString, string uri)
func (_IoTeXDID *IoTeXDIDFilterer) WatchUpdateURI(opts *bind.WatchOpts, sink chan<- *IoTeXDIDUpdateURI, didString []string) (event.Subscription, error) {

	var didStringRule []interface{}
	for _, didStringItem := range didString {
		didStringRule = append(didStringRule, didStringItem)
	}

	logs, sub, err := _IoTeXDID.contract.WatchLogs(opts, "UpdateURI", didStringRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(IoTeXDIDUpdateURI)
				if err := _IoTeXDID.contract.UnpackLog(event, "UpdateURI", log); err != nil {
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

// ParseUpdateURI is a log parse operation binding the contract event 0xa55ea16d9a68fded3c5b007d1fbba6ed1524d9354d44c66bcc54c721108586c4.
//
// Solidity: event UpdateURI(string indexed didString, string uri)
func (_IoTeXDID *IoTeXDIDFilterer) ParseUpdateURI(log types.Log) (*IoTeXDIDUpdateURI, error) {
	event := new(IoTeXDIDUpdateURI)
	if err := _IoTeXDID.contract.UnpackLog(event, "UpdateURI", log); err != nil {
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
