// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package bookingtokenoperator

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

// BookingtokenoperatorMetaData contains all meta data concerning the Bookingtokenoperator contract.
var BookingtokenoperatorMetaData = &bind.MetaData{
	ABI: "[{\"inputs\":[{\"internalType\":\"address\",\"name\":\"token\",\"type\":\"address\"}],\"name\":\"SafeERC20FailedOperation\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"tokenId\",\"type\":\"uint256\"},{\"internalType\":\"contractIERC20\",\"name\":\"actualPaymentToken\",\"type\":\"address\"},{\"internalType\":\"contractIERC20\",\"name\":\"expectedPaymentToken\",\"type\":\"address\"}],\"name\":\"UnexpectedPaymentToken\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"tokenId\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"actualPrice\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"expectedPrice\",\"type\":\"uint256\"}],\"name\":\"UnexpectedPrice\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"NATIVE_PAYMENT\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"OFFCHAIN_PAYMENT\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"}]",
	Bin: "0x610f17610035600b8282823980515f1a60731461002957634e487b7160e01b5f525f60045260245ffd5b305f52607381538281f3fe73000000000000000000000000000000000000000030146080604052600436106100cf575f3560e01c8063a0f07c7411610088578063c7bffa9611610063578063c7bffa96146101ba578063e4c22569146101d9578063fd13a43e146101f8575f80fd5b8063a0f07c7414610170578063b54e72d814610193578063bfb26c06146101b2575f80fd5b8063348e06dd116100b8578063348e06dd14610113578063793dddac146101325780637adf63b714610151575f80fd5b806307e47316146100d357806321b87f3a146100f4575b5f80fd5b8180156100de575f80fd5b506100f26100ed366004610b5f565b610217565b005b8180156100ff575f80fd5b506100f261010e366004610bab565b61029e565b81801561011e575f80fd5b506100f261012d366004610bab565b610454565b81801561013d575f80fd5b506100f261014c366004610b5f565b6104ce565b81801561015c575f80fd5b506100f261016b366004610be8565b610526565b6101775f81565b6040516001600160a01b03909116815260200160405180910390f35b81801561019e575f80fd5b506100f26101ad366004610c2f565b6107a7565b610177600181565b8180156101c5575f80fd5b506100f26101d4366004610c84565b610836565b8180156101e4575f80fd5b506100f26101f3366004610cea565b6108a0565b818015610203575f80fd5b506100f2610212366004610c2f565b610928565b6040517f2a1193800000000000000000000000000000000000000000000000000000000081526004810184905261ffff8084166024830152821660448201526001600160a01b03851690632a119380906064015b5f604051808303815f87803b158015610282575f80fd5b505af1158015610294573d5f803e3d5ffd5b5050505050505050565b6040517fb191d092000000000000000000000000000000000000000000000000000000008152600481018390525f906001600160a01b0385169063b191d09290602401602060405180830381865afa1580156102fc573d5f803e3d5ffd5b505050506040513d601f19601f820116820180604052508101906103209190610df6565b90506001600160a01b03811661039457604051631c54f0f760e01b815260048101849052602481018390526001600160a01b03851690631c54f0f79084906044015f604051808303818588803b158015610378575f80fd5b505af115801561038a573d5f803e3d5ffd5b505050505061044e565b5f196001600160a01b0382160161040657604051631c54f0f760e01b815260048101849052602481018390526001600160a01b03851690631c54f0f7906044015f604051808303815f87803b1580156103eb575f80fd5b505af11580156103fd573d5f803e3d5ffd5b5050505061044e565b61041a6001600160a01b0382168584610987565b604051631c54f0f760e01b815260048101849052602481018390526001600160a01b03851690631c54f0f79060440161026b565b50505050565b6040517fbe66718800000000000000000000000000000000000000000000000000000000815260048101839052602481018290526001600160a01b0384169063be667188906044015f604051808303815f87803b1580156104b3575f80fd5b505af11580156104c5573d5f803e3d5ffd5b50505050505050565b6040517f74fe60e90000000000000000000000000000000000000000000000000000000081526004810184905261ffff8084166024830152821660448201526001600160a01b038516906374fe60e99060640161026b565b6040517e4fdd3c000000000000000000000000000000000000000000000000000000008152600481018490525f9081906001600160a01b03871690624fdd3c906024016040805180830381865afa158015610583573d5f803e3d5ffd5b505050506040513d601f19601f820116820180604052508101906105a79190610e18565b915091508382146105fa576040517f30cb9c410000000000000000000000000000000000000000000000000000000081526004810186905260248101839052604481018590526064015b60405180910390fd5b826001600160a01b0316816001600160a01b03161461065f576040517f41fe0054000000000000000000000000000000000000000000000000000000008152600481018690526001600160a01b038083166024830152841660448201526064016105f1565b6001600160a01b0381166106ca576040516396591edd60e01b8152600481018690526001600160a01b038716906396591edd9084906024015f604051808303818588803b1580156106ae575f80fd5b505af11580156106c0573d5f803e3d5ffd5b505050505061079f565b5f196001600160a01b03821601610735576040516396591edd60e01b8152600481018690526001600160a01b038716906396591edd906024015f604051808303815f87803b15801561071a575f80fd5b505af115801561072c573d5f803e3d5ffd5b5050505061079f565b6107496001600160a01b0382168784610987565b6040516396591edd60e01b8152600481018690526001600160a01b038716906396591edd906024015f604051808303815f87803b158015610788575f80fd5b505af115801561079a573d5f803e3d5ffd5b505050505b505050505050565b6040517ff7e45f09000000000000000000000000000000000000000000000000000000008152600481018590526024810184905261ffff8084166044830152821660648201526001600160a01b0386169063f7e45f09906084015b5f604051808303815f87803b158015610819575f80fd5b505af115801561082b573d5f803e3d5ffd5b505050505050505050565b6040517fe5a6725c000000000000000000000000000000000000000000000000000000008152600481018290526001600160a01b0383169063e5a6725c906024015f604051808303815f87803b15801561088e575f80fd5b505af115801561079f573d5f803e3d5ffd5b6040517fdb2b26820000000000000000000000000000000000000000000000000000000081526001600160a01b0389169063db2b2682906108f1908a908a908a908a908a908a908a90600401610e46565b5f604051808303815f87803b158015610908575f80fd5b505af115801561091a573d5f803e3d5ffd5b505050505050505050505050565b6040517f74aa2048000000000000000000000000000000000000000000000000000000008152600481018590526024810184905261ffff8084166044830152821660648201526001600160a01b038616906374aa204890608401610802565b604080516001600160a01b038416602482015260448082018490528251808303909101815260649091019091526020810180517bffffffffffffffffffffffffffffffffffffffffffffffffffffffff1663095ea7b360e01b1790526109ed8482610a64565b61044e57604080516001600160a01b03851660248201525f6044808301919091528251808303909101815260649091019091526020810180517bffffffffffffffffffffffffffffffffffffffffffffffffffffffff1663095ea7b360e01b179052610a5a908590610aad565b61044e8482610aad565b5f805f8060205f8651602088015f8a5af192503d91505f519050828015610aa357508115610a955780600114610aa3565b5f866001600160a01b03163b115b9695505050505050565b5f8060205f8451602086015f885af180610acc576040513d5f823e3d81fd5b50505f513d91508115610ae3578060011415610af0565b6001600160a01b0384163b155b1561044e576040517f5274afe70000000000000000000000000000000000000000000000000000000081526001600160a01b03851660048201526024016105f1565b6001600160a01b0381168114610b46575f80fd5b50565b803561ffff81168114610b5a575f80fd5b919050565b5f805f8060808587031215610b72575f80fd5b8435610b7d81610b32565b935060208501359250610b9260408601610b49565b9150610ba060608601610b49565b905092959194509250565b5f805f60608486031215610bbd575f80fd5b8335610bc881610b32565b95602085013595506040909401359392505050565b8035610b5a81610b32565b5f805f8060808587031215610bfb575f80fd5b8435610c0681610b32565b935060208501359250604085013591506060850135610c2481610b32565b939692955090935050565b5f805f805f60a08688031215610c43575f80fd5b8535610c4e81610b32565b94506020860135935060408601359250610c6a60608701610b49565b9150610c7860808701610b49565b90509295509295909350565b5f8060408385031215610c95575f80fd5b8235610ca081610b32565b946020939093013593505050565b7f4e487b71000000000000000000000000000000000000000000000000000000005f52604160045260245ffd5b80358015158114610b5a575f80fd5b5f805f805f805f80610100898b031215610d02575f80fd5b8835610d0d81610b32565b97506020890135610d1d81610b32565b9650604089013567ffffffffffffffff80821115610d39575f80fd5b818b0191508b601f830112610d4c575f80fd5b813581811115610d5e57610d5e610cae565b604051601f8201601f19908116603f01168101908382118183101715610d8657610d86610cae565b816040528281528e6020848701011115610d9e575f80fd5b826020860160208301375f60208483010152809a5050505050506060890135945060808901359350610dd260a08a01610bdd565b925060c08901359150610de760e08a01610cdb565b90509295985092959890939650565b5f60208284031215610e06575f80fd5b8151610e1181610b32565b9392505050565b5f8060408385031215610e29575f80fd5b825191506020830151610e3b81610b32565b809150509250929050565b6001600160a01b03881681525f602060e0602084015288518060e08501525f5b81811015610e83578a810183015185820161010001528201610e66565b5061010091505f82828601015281601f19601f83011685010192505050866040830152856060830152610ec160808301866001600160a01b03169052565b8360a0830152610ed560c083018415159052565b9897505050505050505056fea2646970667358221220646ac6dbee1c6eb243b3291d6315c702a940acac58a583eb89c54437d63fdff864736f6c63430008180033",
}

// BookingtokenoperatorABI is the input ABI used to generate the binding from.
// Deprecated: Use BookingtokenoperatorMetaData.ABI instead.
var BookingtokenoperatorABI = BookingtokenoperatorMetaData.ABI

// BookingtokenoperatorBin is the compiled bytecode used for deploying new contracts.
// Deprecated: Use BookingtokenoperatorMetaData.Bin instead.
var BookingtokenoperatorBin = BookingtokenoperatorMetaData.Bin

// DeployBookingtokenoperator deploys a new Ethereum contract, binding an instance of Bookingtokenoperator to it.
func DeployBookingtokenoperator(auth *bind.TransactOpts, backend bind.ContractBackend) (common.Address, *types.Transaction, *Bookingtokenoperator, error) {
	parsed, err := BookingtokenoperatorMetaData.GetAbi()
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	if parsed == nil {
		return common.Address{}, nil, nil, errors.New("GetABI returned nil")
	}

	address, tx, contract, err := bind.DeployContract(auth, *parsed, common.FromHex(BookingtokenoperatorBin), backend)
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	return address, tx, &Bookingtokenoperator{BookingtokenoperatorCaller: BookingtokenoperatorCaller{contract: contract}, BookingtokenoperatorTransactor: BookingtokenoperatorTransactor{contract: contract}, BookingtokenoperatorFilterer: BookingtokenoperatorFilterer{contract: contract}}, nil
}

// Bookingtokenoperator is an auto generated Go binding around an Ethereum contract.
type Bookingtokenoperator struct {
	BookingtokenoperatorCaller     // Read-only binding to the contract
	BookingtokenoperatorTransactor // Write-only binding to the contract
	BookingtokenoperatorFilterer   // Log filterer for contract events
}

// BookingtokenoperatorCaller is an auto generated read-only Go binding around an Ethereum contract.
type BookingtokenoperatorCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// BookingtokenoperatorTransactor is an auto generated write-only Go binding around an Ethereum contract.
type BookingtokenoperatorTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// BookingtokenoperatorFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type BookingtokenoperatorFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// BookingtokenoperatorSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type BookingtokenoperatorSession struct {
	Contract     *Bookingtokenoperator // Generic contract binding to set the session for
	CallOpts     bind.CallOpts         // Call options to use throughout this session
	TransactOpts bind.TransactOpts     // Transaction auth options to use throughout this session
}

// BookingtokenoperatorCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type BookingtokenoperatorCallerSession struct {
	Contract *BookingtokenoperatorCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts               // Call options to use throughout this session
}

// BookingtokenoperatorTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type BookingtokenoperatorTransactorSession struct {
	Contract     *BookingtokenoperatorTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts               // Transaction auth options to use throughout this session
}

// BookingtokenoperatorRaw is an auto generated low-level Go binding around an Ethereum contract.
type BookingtokenoperatorRaw struct {
	Contract *Bookingtokenoperator // Generic contract binding to access the raw methods on
}

// BookingtokenoperatorCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type BookingtokenoperatorCallerRaw struct {
	Contract *BookingtokenoperatorCaller // Generic read-only contract binding to access the raw methods on
}

// BookingtokenoperatorTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type BookingtokenoperatorTransactorRaw struct {
	Contract *BookingtokenoperatorTransactor // Generic write-only contract binding to access the raw methods on
}

// NewBookingtokenoperator creates a new instance of Bookingtokenoperator, bound to a specific deployed contract.
func NewBookingtokenoperator(address common.Address, backend bind.ContractBackend) (*Bookingtokenoperator, error) {
	contract, err := bindBookingtokenoperator(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &Bookingtokenoperator{BookingtokenoperatorCaller: BookingtokenoperatorCaller{contract: contract}, BookingtokenoperatorTransactor: BookingtokenoperatorTransactor{contract: contract}, BookingtokenoperatorFilterer: BookingtokenoperatorFilterer{contract: contract}}, nil
}

// NewBookingtokenoperatorCaller creates a new read-only instance of Bookingtokenoperator, bound to a specific deployed contract.
func NewBookingtokenoperatorCaller(address common.Address, caller bind.ContractCaller) (*BookingtokenoperatorCaller, error) {
	contract, err := bindBookingtokenoperator(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &BookingtokenoperatorCaller{contract: contract}, nil
}

// NewBookingtokenoperatorTransactor creates a new write-only instance of Bookingtokenoperator, bound to a specific deployed contract.
func NewBookingtokenoperatorTransactor(address common.Address, transactor bind.ContractTransactor) (*BookingtokenoperatorTransactor, error) {
	contract, err := bindBookingtokenoperator(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &BookingtokenoperatorTransactor{contract: contract}, nil
}

// NewBookingtokenoperatorFilterer creates a new log filterer instance of Bookingtokenoperator, bound to a specific deployed contract.
func NewBookingtokenoperatorFilterer(address common.Address, filterer bind.ContractFilterer) (*BookingtokenoperatorFilterer, error) {
	contract, err := bindBookingtokenoperator(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &BookingtokenoperatorFilterer{contract: contract}, nil
}

// bindBookingtokenoperator binds a generic wrapper to an already deployed contract.
func bindBookingtokenoperator(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := BookingtokenoperatorMetaData.GetAbi()
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, *parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Bookingtokenoperator *BookingtokenoperatorRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _Bookingtokenoperator.Contract.BookingtokenoperatorCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Bookingtokenoperator *BookingtokenoperatorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Bookingtokenoperator.Contract.BookingtokenoperatorTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Bookingtokenoperator *BookingtokenoperatorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Bookingtokenoperator.Contract.BookingtokenoperatorTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Bookingtokenoperator *BookingtokenoperatorCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _Bookingtokenoperator.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Bookingtokenoperator *BookingtokenoperatorTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Bookingtokenoperator.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Bookingtokenoperator *BookingtokenoperatorTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Bookingtokenoperator.Contract.contract.Transact(opts, method, params...)
}

// NATIVEPAYMENT is a free data retrieval call binding the contract method 0xa0f07c74.
//
// Solidity: function NATIVE_PAYMENT() view returns(address)
func (_Bookingtokenoperator *BookingtokenoperatorCaller) NATIVEPAYMENT(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _Bookingtokenoperator.contract.Call(opts, &out, "NATIVE_PAYMENT")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// NATIVEPAYMENT is a free data retrieval call binding the contract method 0xa0f07c74.
//
// Solidity: function NATIVE_PAYMENT() view returns(address)
func (_Bookingtokenoperator *BookingtokenoperatorSession) NATIVEPAYMENT() (common.Address, error) {
	return _Bookingtokenoperator.Contract.NATIVEPAYMENT(&_Bookingtokenoperator.CallOpts)
}

// NATIVEPAYMENT is a free data retrieval call binding the contract method 0xa0f07c74.
//
// Solidity: function NATIVE_PAYMENT() view returns(address)
func (_Bookingtokenoperator *BookingtokenoperatorCallerSession) NATIVEPAYMENT() (common.Address, error) {
	return _Bookingtokenoperator.Contract.NATIVEPAYMENT(&_Bookingtokenoperator.CallOpts)
}

// OFFCHAINPAYMENT is a free data retrieval call binding the contract method 0xbfb26c06.
//
// Solidity: function OFFCHAIN_PAYMENT() view returns(address)
func (_Bookingtokenoperator *BookingtokenoperatorCaller) OFFCHAINPAYMENT(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _Bookingtokenoperator.contract.Call(opts, &out, "OFFCHAIN_PAYMENT")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// OFFCHAINPAYMENT is a free data retrieval call binding the contract method 0xbfb26c06.
//
// Solidity: function OFFCHAIN_PAYMENT() view returns(address)
func (_Bookingtokenoperator *BookingtokenoperatorSession) OFFCHAINPAYMENT() (common.Address, error) {
	return _Bookingtokenoperator.Contract.OFFCHAINPAYMENT(&_Bookingtokenoperator.CallOpts)
}

// OFFCHAINPAYMENT is a free data retrieval call binding the contract method 0xbfb26c06.
//
// Solidity: function OFFCHAIN_PAYMENT() view returns(address)
func (_Bookingtokenoperator *BookingtokenoperatorCallerSession) OFFCHAINPAYMENT() (common.Address, error) {
	return _Bookingtokenoperator.Contract.OFFCHAINPAYMENT(&_Bookingtokenoperator.CallOpts)
}
