// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package nullusd

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

// NullusdMetaData contains all meta data concerning the Nullusd contract.
var NullusdMetaData = &bind.MetaData{
	ABI: "[{\"inputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"constructor\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"spender\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"allowance\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"needed\",\"type\":\"uint256\"}],\"name\":\"ERC20InsufficientAllowance\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"sender\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"balance\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"needed\",\"type\":\"uint256\"}],\"name\":\"ERC20InsufficientBalance\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"approver\",\"type\":\"address\"}],\"name\":\"ERC20InvalidApprover\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"receiver\",\"type\":\"address\"}],\"name\":\"ERC20InvalidReceiver\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"sender\",\"type\":\"address\"}],\"name\":\"ERC20InvalidSender\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"spender\",\"type\":\"address\"}],\"name\":\"ERC20InvalidSpender\",\"type\":\"error\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"owner\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"spender\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"value\",\"type\":\"uint256\"}],\"name\":\"Approval\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"from\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"value\",\"type\":\"uint256\"}],\"name\":\"Transfer\",\"type\":\"event\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"owner\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"spender\",\"type\":\"address\"}],\"name\":\"allowance\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"spender\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"value\",\"type\":\"uint256\"}],\"name\":\"approve\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"}],\"name\":\"balanceOf\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"decimals\",\"outputs\":[{\"internalType\":\"uint8\",\"name\":\"\",\"type\":\"uint8\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"name\",\"outputs\":[{\"internalType\":\"string\",\"name\":\"\",\"type\":\"string\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"symbol\",\"outputs\":[{\"internalType\":\"string\",\"name\":\"\",\"type\":\"string\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"totalSupply\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"value\",\"type\":\"uint256\"}],\"name\":\"transfer\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"from\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"value\",\"type\":\"uint256\"}],\"name\":\"transferFrom\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"}]",
	Bin: "0x608060405234801562000010575f80fd5b5060405180604001604052806007815260200166139d5b1b1554d160ca1b81525060405180604001604052806004815260200163139554d160e21b81525081600390816200005f9190620002c0565b5060046200006e8282620002c0565b505050620000a83362000086620000ae60201b60201c565b6200009390600a6200049b565b620000a290620f4240620004b2565b620000b3565b620004e2565b601290565b6001600160a01b038216620000e25760405163ec442f0560e01b81525f60048201526024015b60405180910390fd5b620000ef5f8383620000f3565b5050565b6001600160a01b03831662000121578060025f828254620001159190620004cc565b90915550620001939050565b6001600160a01b0383165f9081526020819052604090205481811015620001755760405163391434e360e21b81526001600160a01b03851660048201526024810182905260448101839052606401620000d9565b6001600160a01b0384165f9081526020819052604090209082900390555b6001600160a01b038216620001b157600280548290039055620001cf565b6001600160a01b0382165f9081526020819052604090208054820190555b816001600160a01b0316836001600160a01b03167fddf252ad1be2c89b69c2b068fc378daa952ba7f163c4a11628f55a4df523b3ef836040516200021591815260200190565b60405180910390a3505050565b634e487b7160e01b5f52604160045260245ffd5b600181811c908216806200024b57607f821691505b6020821081036200026a57634e487b7160e01b5f52602260045260245ffd5b50919050565b601f821115620002bb57805f5260205f20601f840160051c81016020851015620002975750805b601f840160051c820191505b81811015620002b8575f8155600101620002a3565b50505b505050565b81516001600160401b03811115620002dc57620002dc62000222565b620002f481620002ed845462000236565b8462000270565b602080601f8311600181146200032a575f8415620003125750858301515b5f19600386901b1c1916600185901b17855562000384565b5f85815260208120601f198616915b828110156200035a5788860151825594840194600190910190840162000339565b50858210156200037857878501515f19600388901b60f8161c191681555b505060018460011b0185555b505050505050565b634e487b7160e01b5f52601160045260245ffd5b600181815b80851115620003e057815f1904821115620003c457620003c46200038c565b80851615620003d257918102915b93841c9390800290620003a5565b509250929050565b5f82620003f85750600162000495565b816200040657505f62000495565b81600181146200041f57600281146200042a576200044a565b600191505062000495565b60ff8411156200043e576200043e6200038c565b50506001821b62000495565b5060208310610133831016604e8410600b84101617156200046f575081810a62000495565b6200047b8383620003a0565b805f19048211156200049157620004916200038c565b0290505b92915050565b5f620004ab60ff841683620003e8565b9392505050565b80820281158282048414176200049557620004956200038c565b808201808211156200049557620004956200038c565b61079e80620004f05f395ff3fe608060405234801561000f575f80fd5b506004361061009f575f3560e01c8063313ce5671161007257806395d89b411161005857806395d89b4114610140578063a9059cbb14610148578063dd62ed3e1461015b575f80fd5b8063313ce5671461010957806370a0823114610118575f80fd5b806306fdde03146100a3578063095ea7b3146100c157806318160ddd146100e457806323b872dd146100f6575b5f80fd5b6100ab610193565b6040516100b891906105f8565b60405180910390f35b6100d46100cf36600461065f565b610223565b60405190151581526020016100b8565b6002545b6040519081526020016100b8565b6100d4610104366004610687565b61023c565b604051601281526020016100b8565b6100e86101263660046106c0565b6001600160a01b03165f9081526020819052604090205490565b6100ab61025f565b6100d461015636600461065f565b61026e565b6100e86101693660046106e0565b6001600160a01b039182165f90815260016020908152604080832093909416825291909152205490565b6060600380546101a290610711565b80601f01602080910402602001604051908101604052809291908181526020018280546101ce90610711565b80156102195780601f106101f057610100808354040283529160200191610219565b820191905f5260205f20905b8154815290600101906020018083116101fc57829003601f168201915b5050505050905090565b5f3361023081858561027b565b60019150505b92915050565b5f3361024985828561028d565b610254858585610326565b506001949350505050565b6060600480546101a290610711565b5f33610230818585610326565b61028883838360016103b5565b505050565b6001600160a01b038381165f908152600160209081526040808320938616835292905220545f1981146103205781811015610312576040517ffb8f41b20000000000000000000000000000000000000000000000000000000081526001600160a01b038416600482015260248101829052604481018390526064015b60405180910390fd5b61032084848484035f6103b5565b50505050565b6001600160a01b038316610368576040517f96c6fd1e0000000000000000000000000000000000000000000000000000000081525f6004820152602401610309565b6001600160a01b0382166103aa576040517fec442f050000000000000000000000000000000000000000000000000000000081525f6004820152602401610309565b6102888383836104b9565b6001600160a01b0384166103f7576040517fe602df050000000000000000000000000000000000000000000000000000000081525f6004820152602401610309565b6001600160a01b038316610439576040517f94280d620000000000000000000000000000000000000000000000000000000081525f6004820152602401610309565b6001600160a01b038085165f908152600160209081526040808320938716835292905220829055801561032057826001600160a01b0316846001600160a01b03167f8c5be1e5ebec7d5bd14f71427d1e84f3dd0314c0f7b2291e5b200ac8c7c3b925846040516104ab91815260200190565b60405180910390a350505050565b6001600160a01b0383166104e3578060025f8282546104d89190610749565b9091555061056c9050565b6001600160a01b0383165f908152602081905260409020548181101561054e576040517fe450d38c0000000000000000000000000000000000000000000000000000000081526001600160a01b03851660048201526024810182905260448101839052606401610309565b6001600160a01b0384165f9081526020819052604090209082900390555b6001600160a01b038216610588576002805482900390556105a6565b6001600160a01b0382165f9081526020819052604090208054820190555b816001600160a01b0316836001600160a01b03167fddf252ad1be2c89b69c2b068fc378daa952ba7f163c4a11628f55a4df523b3ef836040516105eb91815260200190565b60405180910390a3505050565b5f602080835283518060208501525f5b8181101561062457858101830151858201604001528201610608565b505f604082860101526040601f19601f8301168501019250505092915050565b80356001600160a01b038116811461065a575f80fd5b919050565b5f8060408385031215610670575f80fd5b61067983610644565b946020939093013593505050565b5f805f60608486031215610699575f80fd5b6106a284610644565b92506106b060208501610644565b9150604084013590509250925092565b5f602082840312156106d0575f80fd5b6106d982610644565b9392505050565b5f80604083850312156106f1575f80fd5b6106fa83610644565b915061070860208401610644565b90509250929050565b600181811c9082168061072557607f821691505b60208210810361074357634e487b7160e01b5f52602260045260245ffd5b50919050565b8082018082111561023657634e487b7160e01b5f52601160045260245ffdfea2646970667358221220ce5dd0abdb9710b3b299adea8cfcd7b8957939009c3ac5dd9f6118ca8aea784064736f6c63430008180033",
}

// NullusdABI is the input ABI used to generate the binding from.
// Deprecated: Use NullusdMetaData.ABI instead.
var NullusdABI = NullusdMetaData.ABI

// NullusdBin is the compiled bytecode used for deploying new contracts.
// Deprecated: Use NullusdMetaData.Bin instead.
var NullusdBin = NullusdMetaData.Bin

// DeployNullusd deploys a new Ethereum contract, binding an instance of Nullusd to it.
func DeployNullusd(auth *bind.TransactOpts, backend bind.ContractBackend) (common.Address, *types.Transaction, *Nullusd, error) {
	parsed, err := NullusdMetaData.GetAbi()
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	if parsed == nil {
		return common.Address{}, nil, nil, errors.New("GetABI returned nil")
	}

	address, tx, contract, err := bind.DeployContract(auth, *parsed, common.FromHex(NullusdBin), backend)
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	return address, tx, &Nullusd{NullusdCaller: NullusdCaller{contract: contract}, NullusdTransactor: NullusdTransactor{contract: contract}, NullusdFilterer: NullusdFilterer{contract: contract}}, nil
}

// Nullusd is an auto generated Go binding around an Ethereum contract.
type Nullusd struct {
	NullusdCaller     // Read-only binding to the contract
	NullusdTransactor // Write-only binding to the contract
	NullusdFilterer   // Log filterer for contract events
}

// NullusdCaller is an auto generated read-only Go binding around an Ethereum contract.
type NullusdCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// NullusdTransactor is an auto generated write-only Go binding around an Ethereum contract.
type NullusdTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// NullusdFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type NullusdFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// NullusdSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type NullusdSession struct {
	Contract     *Nullusd          // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// NullusdCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type NullusdCallerSession struct {
	Contract *NullusdCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts  // Call options to use throughout this session
}

// NullusdTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type NullusdTransactorSession struct {
	Contract     *NullusdTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts  // Transaction auth options to use throughout this session
}

// NullusdRaw is an auto generated low-level Go binding around an Ethereum contract.
type NullusdRaw struct {
	Contract *Nullusd // Generic contract binding to access the raw methods on
}

// NullusdCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type NullusdCallerRaw struct {
	Contract *NullusdCaller // Generic read-only contract binding to access the raw methods on
}

// NullusdTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type NullusdTransactorRaw struct {
	Contract *NullusdTransactor // Generic write-only contract binding to access the raw methods on
}

// NewNullusd creates a new instance of Nullusd, bound to a specific deployed contract.
func NewNullusd(address common.Address, backend bind.ContractBackend) (*Nullusd, error) {
	contract, err := bindNullusd(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &Nullusd{NullusdCaller: NullusdCaller{contract: contract}, NullusdTransactor: NullusdTransactor{contract: contract}, NullusdFilterer: NullusdFilterer{contract: contract}}, nil
}

// NewNullusdCaller creates a new read-only instance of Nullusd, bound to a specific deployed contract.
func NewNullusdCaller(address common.Address, caller bind.ContractCaller) (*NullusdCaller, error) {
	contract, err := bindNullusd(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &NullusdCaller{contract: contract}, nil
}

// NewNullusdTransactor creates a new write-only instance of Nullusd, bound to a specific deployed contract.
func NewNullusdTransactor(address common.Address, transactor bind.ContractTransactor) (*NullusdTransactor, error) {
	contract, err := bindNullusd(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &NullusdTransactor{contract: contract}, nil
}

// NewNullusdFilterer creates a new log filterer instance of Nullusd, bound to a specific deployed contract.
func NewNullusdFilterer(address common.Address, filterer bind.ContractFilterer) (*NullusdFilterer, error) {
	contract, err := bindNullusd(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &NullusdFilterer{contract: contract}, nil
}

// bindNullusd binds a generic wrapper to an already deployed contract.
func bindNullusd(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := NullusdMetaData.GetAbi()
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, *parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Nullusd *NullusdRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _Nullusd.Contract.NullusdCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Nullusd *NullusdRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Nullusd.Contract.NullusdTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Nullusd *NullusdRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Nullusd.Contract.NullusdTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Nullusd *NullusdCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _Nullusd.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Nullusd *NullusdTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Nullusd.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Nullusd *NullusdTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Nullusd.Contract.contract.Transact(opts, method, params...)
}

// Allowance is a free data retrieval call binding the contract method 0xdd62ed3e.
//
// Solidity: function allowance(address owner, address spender) view returns(uint256)
func (_Nullusd *NullusdCaller) Allowance(opts *bind.CallOpts, owner common.Address, spender common.Address) (*big.Int, error) {
	var out []interface{}
	err := _Nullusd.contract.Call(opts, &out, "allowance", owner, spender)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// Allowance is a free data retrieval call binding the contract method 0xdd62ed3e.
//
// Solidity: function allowance(address owner, address spender) view returns(uint256)
func (_Nullusd *NullusdSession) Allowance(owner common.Address, spender common.Address) (*big.Int, error) {
	return _Nullusd.Contract.Allowance(&_Nullusd.CallOpts, owner, spender)
}

// Allowance is a free data retrieval call binding the contract method 0xdd62ed3e.
//
// Solidity: function allowance(address owner, address spender) view returns(uint256)
func (_Nullusd *NullusdCallerSession) Allowance(owner common.Address, spender common.Address) (*big.Int, error) {
	return _Nullusd.Contract.Allowance(&_Nullusd.CallOpts, owner, spender)
}

// BalanceOf is a free data retrieval call binding the contract method 0x70a08231.
//
// Solidity: function balanceOf(address account) view returns(uint256)
func (_Nullusd *NullusdCaller) BalanceOf(opts *bind.CallOpts, account common.Address) (*big.Int, error) {
	var out []interface{}
	err := _Nullusd.contract.Call(opts, &out, "balanceOf", account)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// BalanceOf is a free data retrieval call binding the contract method 0x70a08231.
//
// Solidity: function balanceOf(address account) view returns(uint256)
func (_Nullusd *NullusdSession) BalanceOf(account common.Address) (*big.Int, error) {
	return _Nullusd.Contract.BalanceOf(&_Nullusd.CallOpts, account)
}

// BalanceOf is a free data retrieval call binding the contract method 0x70a08231.
//
// Solidity: function balanceOf(address account) view returns(uint256)
func (_Nullusd *NullusdCallerSession) BalanceOf(account common.Address) (*big.Int, error) {
	return _Nullusd.Contract.BalanceOf(&_Nullusd.CallOpts, account)
}

// Decimals is a free data retrieval call binding the contract method 0x313ce567.
//
// Solidity: function decimals() view returns(uint8)
func (_Nullusd *NullusdCaller) Decimals(opts *bind.CallOpts) (uint8, error) {
	var out []interface{}
	err := _Nullusd.contract.Call(opts, &out, "decimals")

	if err != nil {
		return *new(uint8), err
	}

	out0 := *abi.ConvertType(out[0], new(uint8)).(*uint8)

	return out0, err

}

// Decimals is a free data retrieval call binding the contract method 0x313ce567.
//
// Solidity: function decimals() view returns(uint8)
func (_Nullusd *NullusdSession) Decimals() (uint8, error) {
	return _Nullusd.Contract.Decimals(&_Nullusd.CallOpts)
}

// Decimals is a free data retrieval call binding the contract method 0x313ce567.
//
// Solidity: function decimals() view returns(uint8)
func (_Nullusd *NullusdCallerSession) Decimals() (uint8, error) {
	return _Nullusd.Contract.Decimals(&_Nullusd.CallOpts)
}

// Name is a free data retrieval call binding the contract method 0x06fdde03.
//
// Solidity: function name() view returns(string)
func (_Nullusd *NullusdCaller) Name(opts *bind.CallOpts) (string, error) {
	var out []interface{}
	err := _Nullusd.contract.Call(opts, &out, "name")

	if err != nil {
		return *new(string), err
	}

	out0 := *abi.ConvertType(out[0], new(string)).(*string)

	return out0, err

}

// Name is a free data retrieval call binding the contract method 0x06fdde03.
//
// Solidity: function name() view returns(string)
func (_Nullusd *NullusdSession) Name() (string, error) {
	return _Nullusd.Contract.Name(&_Nullusd.CallOpts)
}

// Name is a free data retrieval call binding the contract method 0x06fdde03.
//
// Solidity: function name() view returns(string)
func (_Nullusd *NullusdCallerSession) Name() (string, error) {
	return _Nullusd.Contract.Name(&_Nullusd.CallOpts)
}

// Symbol is a free data retrieval call binding the contract method 0x95d89b41.
//
// Solidity: function symbol() view returns(string)
func (_Nullusd *NullusdCaller) Symbol(opts *bind.CallOpts) (string, error) {
	var out []interface{}
	err := _Nullusd.contract.Call(opts, &out, "symbol")

	if err != nil {
		return *new(string), err
	}

	out0 := *abi.ConvertType(out[0], new(string)).(*string)

	return out0, err

}

// Symbol is a free data retrieval call binding the contract method 0x95d89b41.
//
// Solidity: function symbol() view returns(string)
func (_Nullusd *NullusdSession) Symbol() (string, error) {
	return _Nullusd.Contract.Symbol(&_Nullusd.CallOpts)
}

// Symbol is a free data retrieval call binding the contract method 0x95d89b41.
//
// Solidity: function symbol() view returns(string)
func (_Nullusd *NullusdCallerSession) Symbol() (string, error) {
	return _Nullusd.Contract.Symbol(&_Nullusd.CallOpts)
}

// TotalSupply is a free data retrieval call binding the contract method 0x18160ddd.
//
// Solidity: function totalSupply() view returns(uint256)
func (_Nullusd *NullusdCaller) TotalSupply(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _Nullusd.contract.Call(opts, &out, "totalSupply")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// TotalSupply is a free data retrieval call binding the contract method 0x18160ddd.
//
// Solidity: function totalSupply() view returns(uint256)
func (_Nullusd *NullusdSession) TotalSupply() (*big.Int, error) {
	return _Nullusd.Contract.TotalSupply(&_Nullusd.CallOpts)
}

// TotalSupply is a free data retrieval call binding the contract method 0x18160ddd.
//
// Solidity: function totalSupply() view returns(uint256)
func (_Nullusd *NullusdCallerSession) TotalSupply() (*big.Int, error) {
	return _Nullusd.Contract.TotalSupply(&_Nullusd.CallOpts)
}

// Approve is a paid mutator transaction binding the contract method 0x095ea7b3.
//
// Solidity: function approve(address spender, uint256 value) returns(bool)
func (_Nullusd *NullusdTransactor) Approve(opts *bind.TransactOpts, spender common.Address, value *big.Int) (*types.Transaction, error) {
	return _Nullusd.contract.Transact(opts, "approve", spender, value)
}

// Approve is a paid mutator transaction binding the contract method 0x095ea7b3.
//
// Solidity: function approve(address spender, uint256 value) returns(bool)
func (_Nullusd *NullusdSession) Approve(spender common.Address, value *big.Int) (*types.Transaction, error) {
	return _Nullusd.Contract.Approve(&_Nullusd.TransactOpts, spender, value)
}

// Approve is a paid mutator transaction binding the contract method 0x095ea7b3.
//
// Solidity: function approve(address spender, uint256 value) returns(bool)
func (_Nullusd *NullusdTransactorSession) Approve(spender common.Address, value *big.Int) (*types.Transaction, error) {
	return _Nullusd.Contract.Approve(&_Nullusd.TransactOpts, spender, value)
}

// Transfer is a paid mutator transaction binding the contract method 0xa9059cbb.
//
// Solidity: function transfer(address to, uint256 value) returns(bool)
func (_Nullusd *NullusdTransactor) Transfer(opts *bind.TransactOpts, to common.Address, value *big.Int) (*types.Transaction, error) {
	return _Nullusd.contract.Transact(opts, "transfer", to, value)
}

// Transfer is a paid mutator transaction binding the contract method 0xa9059cbb.
//
// Solidity: function transfer(address to, uint256 value) returns(bool)
func (_Nullusd *NullusdSession) Transfer(to common.Address, value *big.Int) (*types.Transaction, error) {
	return _Nullusd.Contract.Transfer(&_Nullusd.TransactOpts, to, value)
}

// Transfer is a paid mutator transaction binding the contract method 0xa9059cbb.
//
// Solidity: function transfer(address to, uint256 value) returns(bool)
func (_Nullusd *NullusdTransactorSession) Transfer(to common.Address, value *big.Int) (*types.Transaction, error) {
	return _Nullusd.Contract.Transfer(&_Nullusd.TransactOpts, to, value)
}

// TransferFrom is a paid mutator transaction binding the contract method 0x23b872dd.
//
// Solidity: function transferFrom(address from, address to, uint256 value) returns(bool)
func (_Nullusd *NullusdTransactor) TransferFrom(opts *bind.TransactOpts, from common.Address, to common.Address, value *big.Int) (*types.Transaction, error) {
	return _Nullusd.contract.Transact(opts, "transferFrom", from, to, value)
}

// TransferFrom is a paid mutator transaction binding the contract method 0x23b872dd.
//
// Solidity: function transferFrom(address from, address to, uint256 value) returns(bool)
func (_Nullusd *NullusdSession) TransferFrom(from common.Address, to common.Address, value *big.Int) (*types.Transaction, error) {
	return _Nullusd.Contract.TransferFrom(&_Nullusd.TransactOpts, from, to, value)
}

// TransferFrom is a paid mutator transaction binding the contract method 0x23b872dd.
//
// Solidity: function transferFrom(address from, address to, uint256 value) returns(bool)
func (_Nullusd *NullusdTransactorSession) TransferFrom(from common.Address, to common.Address, value *big.Int) (*types.Transaction, error) {
	return _Nullusd.Contract.TransferFrom(&_Nullusd.TransactOpts, from, to, value)
}

// NullusdApprovalIterator is returned from FilterApproval and is used to iterate over the raw logs and unpacked data for Approval events raised by the Nullusd contract.
type NullusdApprovalIterator struct {
	Event *NullusdApproval // Event containing the contract specifics and raw log

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
func (it *NullusdApprovalIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(NullusdApproval)
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
		it.Event = new(NullusdApproval)
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
func (it *NullusdApprovalIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *NullusdApprovalIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// NullusdApproval represents a Approval event raised by the Nullusd contract.
type NullusdApproval struct {
	Owner   common.Address
	Spender common.Address
	Value   *big.Int
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterApproval is a free log retrieval operation binding the contract event 0x8c5be1e5ebec7d5bd14f71427d1e84f3dd0314c0f7b2291e5b200ac8c7c3b925.
//
// Solidity: event Approval(address indexed owner, address indexed spender, uint256 value)
func (_Nullusd *NullusdFilterer) FilterApproval(opts *bind.FilterOpts, owner []common.Address, spender []common.Address) (*NullusdApprovalIterator, error) {

	var ownerRule []interface{}
	for _, ownerItem := range owner {
		ownerRule = append(ownerRule, ownerItem)
	}
	var spenderRule []interface{}
	for _, spenderItem := range spender {
		spenderRule = append(spenderRule, spenderItem)
	}

	logs, sub, err := _Nullusd.contract.FilterLogs(opts, "Approval", ownerRule, spenderRule)
	if err != nil {
		return nil, err
	}
	return &NullusdApprovalIterator{contract: _Nullusd.contract, event: "Approval", logs: logs, sub: sub}, nil
}

// WatchApproval is a free log subscription operation binding the contract event 0x8c5be1e5ebec7d5bd14f71427d1e84f3dd0314c0f7b2291e5b200ac8c7c3b925.
//
// Solidity: event Approval(address indexed owner, address indexed spender, uint256 value)
func (_Nullusd *NullusdFilterer) WatchApproval(opts *bind.WatchOpts, sink chan<- *NullusdApproval, owner []common.Address, spender []common.Address) (event.Subscription, error) {

	var ownerRule []interface{}
	for _, ownerItem := range owner {
		ownerRule = append(ownerRule, ownerItem)
	}
	var spenderRule []interface{}
	for _, spenderItem := range spender {
		spenderRule = append(spenderRule, spenderItem)
	}

	logs, sub, err := _Nullusd.contract.WatchLogs(opts, "Approval", ownerRule, spenderRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(NullusdApproval)
				if err := _Nullusd.contract.UnpackLog(event, "Approval", log); err != nil {
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

// ParseApproval is a log parse operation binding the contract event 0x8c5be1e5ebec7d5bd14f71427d1e84f3dd0314c0f7b2291e5b200ac8c7c3b925.
//
// Solidity: event Approval(address indexed owner, address indexed spender, uint256 value)
func (_Nullusd *NullusdFilterer) ParseApproval(log types.Log) (*NullusdApproval, error) {
	event := new(NullusdApproval)
	if err := _Nullusd.contract.UnpackLog(event, "Approval", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// NullusdTransferIterator is returned from FilterTransfer and is used to iterate over the raw logs and unpacked data for Transfer events raised by the Nullusd contract.
type NullusdTransferIterator struct {
	Event *NullusdTransfer // Event containing the contract specifics and raw log

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
func (it *NullusdTransferIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(NullusdTransfer)
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
		it.Event = new(NullusdTransfer)
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
func (it *NullusdTransferIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *NullusdTransferIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// NullusdTransfer represents a Transfer event raised by the Nullusd contract.
type NullusdTransfer struct {
	From  common.Address
	To    common.Address
	Value *big.Int
	Raw   types.Log // Blockchain specific contextual infos
}

// FilterTransfer is a free log retrieval operation binding the contract event 0xddf252ad1be2c89b69c2b068fc378daa952ba7f163c4a11628f55a4df523b3ef.
//
// Solidity: event Transfer(address indexed from, address indexed to, uint256 value)
func (_Nullusd *NullusdFilterer) FilterTransfer(opts *bind.FilterOpts, from []common.Address, to []common.Address) (*NullusdTransferIterator, error) {

	var fromRule []interface{}
	for _, fromItem := range from {
		fromRule = append(fromRule, fromItem)
	}
	var toRule []interface{}
	for _, toItem := range to {
		toRule = append(toRule, toItem)
	}

	logs, sub, err := _Nullusd.contract.FilterLogs(opts, "Transfer", fromRule, toRule)
	if err != nil {
		return nil, err
	}
	return &NullusdTransferIterator{contract: _Nullusd.contract, event: "Transfer", logs: logs, sub: sub}, nil
}

// WatchTransfer is a free log subscription operation binding the contract event 0xddf252ad1be2c89b69c2b068fc378daa952ba7f163c4a11628f55a4df523b3ef.
//
// Solidity: event Transfer(address indexed from, address indexed to, uint256 value)
func (_Nullusd *NullusdFilterer) WatchTransfer(opts *bind.WatchOpts, sink chan<- *NullusdTransfer, from []common.Address, to []common.Address) (event.Subscription, error) {

	var fromRule []interface{}
	for _, fromItem := range from {
		fromRule = append(fromRule, fromItem)
	}
	var toRule []interface{}
	for _, toItem := range to {
		toRule = append(toRule, toItem)
	}

	logs, sub, err := _Nullusd.contract.WatchLogs(opts, "Transfer", fromRule, toRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(NullusdTransfer)
				if err := _Nullusd.contract.UnpackLog(event, "Transfer", log); err != nil {
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

// ParseTransfer is a log parse operation binding the contract event 0xddf252ad1be2c89b69c2b068fc378daa952ba7f163c4a11628f55a4df523b3ef.
//
// Solidity: event Transfer(address indexed from, address indexed to, uint256 value)
func (_Nullusd *NullusdFilterer) ParseTransfer(log types.Log) (*NullusdTransfer, error) {
	event := new(NullusdTransfer)
	if err := _Nullusd.contract.UnpackLog(event, "Transfer", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}
