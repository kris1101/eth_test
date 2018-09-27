// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package mypxc

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

// PxcABI is the input ABI used to generate the binding from.
const PxcABI = "[{\"constant\":true,\"inputs\":[],\"name\":\"name\",\"outputs\":[{\"name\":\"\",\"type\":\"string\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"_spender\",\"type\":\"address\"},{\"name\":\"_value\",\"type\":\"uint256\"}],\"name\":\"approve\",\"outputs\":[{\"name\":\"success\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"totalSupply\",\"outputs\":[{\"name\":\"totalSupply\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"_from\",\"type\":\"address\"},{\"name\":\"_to\",\"type\":\"address\"},{\"name\":\"_value\",\"type\":\"uint256\"}],\"name\":\"transferFrom\",\"outputs\":[{\"name\":\"success\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"foundation\",\"outputs\":[{\"name\":\"\",\"type\":\"address\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"name\":\"_owner\",\"type\":\"address\"}],\"name\":\"balanceOf\",\"outputs\":[{\"name\":\"balance\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"sym\",\"outputs\":[{\"name\":\"\",\"type\":\"string\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"_to\",\"type\":\"address\"},{\"name\":\"_value\",\"type\":\"uint256\"}],\"name\":\"transfer\",\"outputs\":[{\"name\":\"success\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"name\":\"_owner\",\"type\":\"address\"},{\"name\":\"_spender\",\"type\":\"address\"}],\"name\":\"allowance\",\"outputs\":[{\"name\":\"remaining\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"name\":\"totalSupply\",\"type\":\"uint256\"},{\"name\":\"addr\",\"type\":\"address\"}],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"constructor\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"name\":\"_from\",\"type\":\"address\"},{\"indexed\":true,\"name\":\"_to\",\"type\":\"address\"},{\"indexed\":false,\"name\":\"_value\",\"type\":\"uint256\"}],\"name\":\"Transfer\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"name\":\"_owner\",\"type\":\"address\"},{\"indexed\":true,\"name\":\"_spender\",\"type\":\"address\"},{\"indexed\":false,\"name\":\"_value\",\"type\":\"uint256\"}],\"name\":\"Approval\",\"type\":\"event\"}]"

// PxcBin is the compiled bytecode used for deploying new contracts.
const PxcBin = `60806040526040805190810160405280600781526020017f70646a636f696e000000000000000000000000000000000000000000000000008152506000908051906020019061004f92919061018a565b506040805190810160405280600381526020017f70786300000000000000000000000000000000000000000000000000000000008152506001908051906020019061009b92919061018a565b503480156100a857600080fd5b50604051604080610f9083398101806040528101908080519060200190929190805190602001909291905050508160028190555080600360006101000a81548173ffffffffffffffffffffffffffffffffffffffff021916908373ffffffffffffffffffffffffffffffffffffffff1602179055508160046000600360009054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200190815260200160002081905550505061022f565b828054600181600116156101000203166002900490600052602060002090601f016020900481019282601f106101cb57805160ff19168380011785556101f9565b828001600101855582156101f9579182015b828111156101f85782518255916020019190600101906101dd565b5b509050610206919061020a565b5090565b61022c91905b80821115610228576000816000905550600101610210565b5090565b90565b610d528061023e6000396000f300608060405260043610610099576000357c0100000000000000000000000000000000000000000000000000000000900463ffffffff16806306fdde031461009e578063095ea7b31461012e57806318160ddd1461019357806323b872dd146101be57806341fbb0501461024357806370a082311461029a57806383813a28146102f1578063a9059cbb14610381578063dd62ed3e146103e6575b600080fd5b3480156100aa57600080fd5b506100b361045d565b6040518080602001828103825283818151815260200191508051906020019080838360005b838110156100f35780820151818401526020810190506100d8565b50505050905090810190601f1680156101205780820380516001836020036101000a031916815260200191505b509250505060405180910390f35b34801561013a57600080fd5b50610179600480360381019080803573ffffffffffffffffffffffffffffffffffffffff169060200190929190803590602001909291905050506104fb565b604051808215151515815260200191505060405180910390f35b34801561019f57600080fd5b506101a86106d5565b6040518082815260200191505060405180910390f35b3480156101ca57600080fd5b50610229600480360381019080803573ffffffffffffffffffffffffffffffffffffffff169060200190929190803573ffffffffffffffffffffffffffffffffffffffff169060200190929190803590602001909291905050506106e2565b604051808215151515815260200191505060405180910390f35b34801561024f57600080fd5b5061025861099e565b604051808273ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200191505060405180910390f35b3480156102a657600080fd5b506102db600480360381019080803573ffffffffffffffffffffffffffffffffffffffff1690602001909291905050506109c4565b6040518082815260200191505060405180910390f35b3480156102fd57600080fd5b50610306610a0d565b6040518080602001828103825283818151815260200191508051906020019080838360005b8381101561034657808201518184015260208101905061032b565b50505050905090810190601f1680156103735780820380516001836020036101000a031916815260200191505b509250505060405180910390f35b34801561038d57600080fd5b506103cc600480360381019080803573ffffffffffffffffffffffffffffffffffffffff16906020019092919080359060200190929190505050610aab565b604051808215151515815260200191505060405180910390f35b3480156103f257600080fd5b50610447600480360381019080803573ffffffffffffffffffffffffffffffffffffffff169060200190929190803573ffffffffffffffffffffffffffffffffffffffff169060200190929190505050610c9f565b6040518082815260200191505060405180910390f35b60008054600181600116156101000203166002900480601f0160208091040260200160405190810160405280929190818152602001828054600181600116156101000203166002900480156104f35780601f106104c8576101008083540402835291602001916104f3565b820191906000526020600020905b8154815290600101906020018083116104d657829003601f168201915b505050505081565b600081600460003373ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff168152602001908152602001600020541015801561054c5750600082115b80156105d65750600082600560003373ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200190815260200160002060008673ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1681526020019081526020016000205401115b156106ca5781600560003373ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200190815260200160002060008573ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff168152602001908152602001600020819055508273ffffffffffffffffffffffffffffffffffffffff163373ffffffffffffffffffffffffffffffffffffffff167f8c5be1e5ebec7d5bd14f71427d1e84f3dd0314c0f7b2291e5b200ac8c7c3b925846040518082815260200191505060405180910390a3600190506106cf565b600090505b92915050565b6000600254905080905090565b600081600560008673ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200190815260200160002060008573ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200190815260200160002054101580156107705750600082115b80156107fb5750600460008473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1681526020019081526020016000205482600460008673ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1681526020019081526020016000205401115b156109925781600460008573ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1681526020019081526020016000206000828254019250508190555081600460008673ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1681526020019081526020016000206000828254039250508190555081600560008673ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200190815260200160002060008573ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff168152602001908152602001600020600082825403925050819055508273ffffffffffffffffffffffffffffffffffffffff168473ffffffffffffffffffffffffffffffffffffffff167fddf252ad1be2c89b69c2b068fc378daa952ba7f163c4a11628f55a4df523b3ef846040518082815260200191505060405180910390a360019050610997565b600090505b9392505050565b600360009054906101000a900473ffffffffffffffffffffffffffffffffffffffff1681565b6000600460008373ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff168152602001908152602001600020549050919050565b60018054600181600116156101000203166002900480601f016020809104026020016040519081016040528092919081815260200182805460018160011615610100020316600290048015610aa35780601f10610a7857610100808354040283529160200191610aa3565b820191906000526020600020905b815481529060010190602001808311610a8657829003601f168201915b505050505081565b600081600460003373ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1681526020019081526020016000205410158015610afc5750600082115b8015610b875750600460008473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1681526020019081526020016000205482600460008673ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1681526020019081526020016000205401115b15610c945781600460008573ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1681526020019081526020016000206000828254019250508190555081600460003373ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff168152602001908152602001600020600082825403925050819055508273ffffffffffffffffffffffffffffffffffffffff163373ffffffffffffffffffffffffffffffffffffffff167fddf252ad1be2c89b69c2b068fc378daa952ba7f163c4a11628f55a4df523b3ef846040518082815260200191505060405180910390a360019050610c99565b600090505b92915050565b6000600560008473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200190815260200160002060008373ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff168152602001908152602001600020549050929150505600a165627a7a72305820f584a211bc335f23279fa8318aec1f27b3ba406814b5c45b6d000cd69e708d290029`

// DeployPxc deploys a new Ethereum contract, binding an instance of Pxc to it.
func DeployPxc(auth *bind.TransactOpts, backend bind.ContractBackend, totalSupply *big.Int, addr common.Address) (common.Address, *types.Transaction, *Pxc, error) {
	parsed, err := abi.JSON(strings.NewReader(PxcABI))
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	address, tx, contract, err := bind.DeployContract(auth, parsed, common.FromHex(PxcBin), backend, totalSupply, addr)
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	return address, tx, &Pxc{PxcCaller: PxcCaller{contract: contract}, PxcTransactor: PxcTransactor{contract: contract}, PxcFilterer: PxcFilterer{contract: contract}}, nil
}

// Pxc is an auto generated Go binding around an Ethereum contract.
type Pxc struct {
	PxcCaller     // Read-only binding to the contract
	PxcTransactor // Write-only binding to the contract
	PxcFilterer   // Log filterer for contract events
}

// PxcCaller is an auto generated read-only Go binding around an Ethereum contract.
type PxcCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// PxcTransactor is an auto generated write-only Go binding around an Ethereum contract.
type PxcTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// PxcFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type PxcFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// PxcSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type PxcSession struct {
	Contract     *Pxc              // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// PxcCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type PxcCallerSession struct {
	Contract *PxcCaller    // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts // Call options to use throughout this session
}

// PxcTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type PxcTransactorSession struct {
	Contract     *PxcTransactor    // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// PxcRaw is an auto generated low-level Go binding around an Ethereum contract.
type PxcRaw struct {
	Contract *Pxc // Generic contract binding to access the raw methods on
}

// PxcCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type PxcCallerRaw struct {
	Contract *PxcCaller // Generic read-only contract binding to access the raw methods on
}

// PxcTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type PxcTransactorRaw struct {
	Contract *PxcTransactor // Generic write-only contract binding to access the raw methods on
}

// NewPxc creates a new instance of Pxc, bound to a specific deployed contract.
func NewPxc(address common.Address, backend bind.ContractBackend) (*Pxc, error) {
	contract, err := bindPxc(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &Pxc{PxcCaller: PxcCaller{contract: contract}, PxcTransactor: PxcTransactor{contract: contract}, PxcFilterer: PxcFilterer{contract: contract}}, nil
}

// NewPxcCaller creates a new read-only instance of Pxc, bound to a specific deployed contract.
func NewPxcCaller(address common.Address, caller bind.ContractCaller) (*PxcCaller, error) {
	contract, err := bindPxc(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &PxcCaller{contract: contract}, nil
}

// NewPxcTransactor creates a new write-only instance of Pxc, bound to a specific deployed contract.
func NewPxcTransactor(address common.Address, transactor bind.ContractTransactor) (*PxcTransactor, error) {
	contract, err := bindPxc(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &PxcTransactor{contract: contract}, nil
}

// NewPxcFilterer creates a new log filterer instance of Pxc, bound to a specific deployed contract.
func NewPxcFilterer(address common.Address, filterer bind.ContractFilterer) (*PxcFilterer, error) {
	contract, err := bindPxc(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &PxcFilterer{contract: contract}, nil
}

// bindPxc binds a generic wrapper to an already deployed contract.
func bindPxc(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(PxcABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Pxc *PxcRaw) Call(opts *bind.CallOpts, result interface{}, method string, params ...interface{}) error {
	return _Pxc.Contract.PxcCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Pxc *PxcRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Pxc.Contract.PxcTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Pxc *PxcRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Pxc.Contract.PxcTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Pxc *PxcCallerRaw) Call(opts *bind.CallOpts, result interface{}, method string, params ...interface{}) error {
	return _Pxc.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Pxc *PxcTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Pxc.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Pxc *PxcTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Pxc.Contract.contract.Transact(opts, method, params...)
}

// Allowance is a free data retrieval call binding the contract method 0xdd62ed3e.
//
// Solidity: function allowance(_owner address, _spender address) constant returns(remaining uint256)
func (_Pxc *PxcCaller) Allowance(opts *bind.CallOpts, _owner common.Address, _spender common.Address) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _Pxc.contract.Call(opts, out, "allowance", _owner, _spender)
	return *ret0, err
}

// Allowance is a free data retrieval call binding the contract method 0xdd62ed3e.
//
// Solidity: function allowance(_owner address, _spender address) constant returns(remaining uint256)
func (_Pxc *PxcSession) Allowance(_owner common.Address, _spender common.Address) (*big.Int, error) {
	return _Pxc.Contract.Allowance(&_Pxc.CallOpts, _owner, _spender)
}

// Allowance is a free data retrieval call binding the contract method 0xdd62ed3e.
//
// Solidity: function allowance(_owner address, _spender address) constant returns(remaining uint256)
func (_Pxc *PxcCallerSession) Allowance(_owner common.Address, _spender common.Address) (*big.Int, error) {
	return _Pxc.Contract.Allowance(&_Pxc.CallOpts, _owner, _spender)
}

// BalanceOf is a free data retrieval call binding the contract method 0x70a08231.
//
// Solidity: function balanceOf(_owner address) constant returns(balance uint256)
func (_Pxc *PxcCaller) BalanceOf(opts *bind.CallOpts, _owner common.Address) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _Pxc.contract.Call(opts, out, "balanceOf", _owner)
	return *ret0, err
}

// BalanceOf is a free data retrieval call binding the contract method 0x70a08231.
//
// Solidity: function balanceOf(_owner address) constant returns(balance uint256)
func (_Pxc *PxcSession) BalanceOf(_owner common.Address) (*big.Int, error) {
	return _Pxc.Contract.BalanceOf(&_Pxc.CallOpts, _owner)
}

// BalanceOf is a free data retrieval call binding the contract method 0x70a08231.
//
// Solidity: function balanceOf(_owner address) constant returns(balance uint256)
func (_Pxc *PxcCallerSession) BalanceOf(_owner common.Address) (*big.Int, error) {
	return _Pxc.Contract.BalanceOf(&_Pxc.CallOpts, _owner)
}

// Foundation is a free data retrieval call binding the contract method 0x41fbb050.
//
// Solidity: function foundation() constant returns(address)
func (_Pxc *PxcCaller) Foundation(opts *bind.CallOpts) (common.Address, error) {
	var (
		ret0 = new(common.Address)
	)
	out := ret0
	err := _Pxc.contract.Call(opts, out, "foundation")
	return *ret0, err
}

// Foundation is a free data retrieval call binding the contract method 0x41fbb050.
//
// Solidity: function foundation() constant returns(address)
func (_Pxc *PxcSession) Foundation() (common.Address, error) {
	return _Pxc.Contract.Foundation(&_Pxc.CallOpts)
}

// Foundation is a free data retrieval call binding the contract method 0x41fbb050.
//
// Solidity: function foundation() constant returns(address)
func (_Pxc *PxcCallerSession) Foundation() (common.Address, error) {
	return _Pxc.Contract.Foundation(&_Pxc.CallOpts)
}

// Name is a free data retrieval call binding the contract method 0x06fdde03.
//
// Solidity: function name() constant returns(string)
func (_Pxc *PxcCaller) Name(opts *bind.CallOpts) (string, error) {
	var (
		ret0 = new(string)
	)
	out := ret0
	err := _Pxc.contract.Call(opts, out, "name")
	return *ret0, err
}

// Name is a free data retrieval call binding the contract method 0x06fdde03.
//
// Solidity: function name() constant returns(string)
func (_Pxc *PxcSession) Name() (string, error) {
	return _Pxc.Contract.Name(&_Pxc.CallOpts)
}

// Name is a free data retrieval call binding the contract method 0x06fdde03.
//
// Solidity: function name() constant returns(string)
func (_Pxc *PxcCallerSession) Name() (string, error) {
	return _Pxc.Contract.Name(&_Pxc.CallOpts)
}

// Sym is a free data retrieval call binding the contract method 0x83813a28.
//
// Solidity: function sym() constant returns(string)
func (_Pxc *PxcCaller) Sym(opts *bind.CallOpts) (string, error) {
	var (
		ret0 = new(string)
	)
	out := ret0
	err := _Pxc.contract.Call(opts, out, "sym")
	return *ret0, err
}

// Sym is a free data retrieval call binding the contract method 0x83813a28.
//
// Solidity: function sym() constant returns(string)
func (_Pxc *PxcSession) Sym() (string, error) {
	return _Pxc.Contract.Sym(&_Pxc.CallOpts)
}

// Sym is a free data retrieval call binding the contract method 0x83813a28.
//
// Solidity: function sym() constant returns(string)
func (_Pxc *PxcCallerSession) Sym() (string, error) {
	return _Pxc.Contract.Sym(&_Pxc.CallOpts)
}

// TotalSupply is a free data retrieval call binding the contract method 0x18160ddd.
//
// Solidity: function totalSupply() constant returns(totalSupply uint256)
func (_Pxc *PxcCaller) TotalSupply(opts *bind.CallOpts) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _Pxc.contract.Call(opts, out, "totalSupply")
	return *ret0, err
}

// TotalSupply is a free data retrieval call binding the contract method 0x18160ddd.
//
// Solidity: function totalSupply() constant returns(totalSupply uint256)
func (_Pxc *PxcSession) TotalSupply() (*big.Int, error) {
	return _Pxc.Contract.TotalSupply(&_Pxc.CallOpts)
}

// TotalSupply is a free data retrieval call binding the contract method 0x18160ddd.
//
// Solidity: function totalSupply() constant returns(totalSupply uint256)
func (_Pxc *PxcCallerSession) TotalSupply() (*big.Int, error) {
	return _Pxc.Contract.TotalSupply(&_Pxc.CallOpts)
}

// Approve is a paid mutator transaction binding the contract method 0x095ea7b3.
//
// Solidity: function approve(_spender address, _value uint256) returns(success bool)
func (_Pxc *PxcTransactor) Approve(opts *bind.TransactOpts, _spender common.Address, _value *big.Int) (*types.Transaction, error) {
	return _Pxc.contract.Transact(opts, "approve", _spender, _value)
}

// Approve is a paid mutator transaction binding the contract method 0x095ea7b3.
//
// Solidity: function approve(_spender address, _value uint256) returns(success bool)
func (_Pxc *PxcSession) Approve(_spender common.Address, _value *big.Int) (*types.Transaction, error) {
	return _Pxc.Contract.Approve(&_Pxc.TransactOpts, _spender, _value)
}

// Approve is a paid mutator transaction binding the contract method 0x095ea7b3.
//
// Solidity: function approve(_spender address, _value uint256) returns(success bool)
func (_Pxc *PxcTransactorSession) Approve(_spender common.Address, _value *big.Int) (*types.Transaction, error) {
	return _Pxc.Contract.Approve(&_Pxc.TransactOpts, _spender, _value)
}

// Transfer is a paid mutator transaction binding the contract method 0xa9059cbb.
//
// Solidity: function transfer(_to address, _value uint256) returns(success bool)
func (_Pxc *PxcTransactor) Transfer(opts *bind.TransactOpts, _to common.Address, _value *big.Int) (*types.Transaction, error) {
	return _Pxc.contract.Transact(opts, "transfer", _to, _value)
}

// Transfer is a paid mutator transaction binding the contract method 0xa9059cbb.
//
// Solidity: function transfer(_to address, _value uint256) returns(success bool)
func (_Pxc *PxcSession) Transfer(_to common.Address, _value *big.Int) (*types.Transaction, error) {
	return _Pxc.Contract.Transfer(&_Pxc.TransactOpts, _to, _value)
}

// Transfer is a paid mutator transaction binding the contract method 0xa9059cbb.
//
// Solidity: function transfer(_to address, _value uint256) returns(success bool)
func (_Pxc *PxcTransactorSession) Transfer(_to common.Address, _value *big.Int) (*types.Transaction, error) {
	return _Pxc.Contract.Transfer(&_Pxc.TransactOpts, _to, _value)
}

// TransferFrom is a paid mutator transaction binding the contract method 0x23b872dd.
//
// Solidity: function transferFrom(_from address, _to address, _value uint256) returns(success bool)
func (_Pxc *PxcTransactor) TransferFrom(opts *bind.TransactOpts, _from common.Address, _to common.Address, _value *big.Int) (*types.Transaction, error) {
	return _Pxc.contract.Transact(opts, "transferFrom", _from, _to, _value)
}

// TransferFrom is a paid mutator transaction binding the contract method 0x23b872dd.
//
// Solidity: function transferFrom(_from address, _to address, _value uint256) returns(success bool)
func (_Pxc *PxcSession) TransferFrom(_from common.Address, _to common.Address, _value *big.Int) (*types.Transaction, error) {
	return _Pxc.Contract.TransferFrom(&_Pxc.TransactOpts, _from, _to, _value)
}

// TransferFrom is a paid mutator transaction binding the contract method 0x23b872dd.
//
// Solidity: function transferFrom(_from address, _to address, _value uint256) returns(success bool)
func (_Pxc *PxcTransactorSession) TransferFrom(_from common.Address, _to common.Address, _value *big.Int) (*types.Transaction, error) {
	return _Pxc.Contract.TransferFrom(&_Pxc.TransactOpts, _from, _to, _value)
}

// PxcApprovalIterator is returned from FilterApproval and is used to iterate over the raw logs and unpacked data for Approval events raised by the Pxc contract.
type PxcApprovalIterator struct {
	Event *PxcApproval // Event containing the contract specifics and raw log

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
func (it *PxcApprovalIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(PxcApproval)
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
		it.Event = new(PxcApproval)
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
func (it *PxcApprovalIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *PxcApprovalIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// PxcApproval represents a Approval event raised by the Pxc contract.
type PxcApproval struct {
	Owner   common.Address
	Spender common.Address
	Value   *big.Int
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterApproval is a free log retrieval operation binding the contract event 0x8c5be1e5ebec7d5bd14f71427d1e84f3dd0314c0f7b2291e5b200ac8c7c3b925.
//
// Solidity: e Approval(_owner indexed address, _spender indexed address, _value uint256)
func (_Pxc *PxcFilterer) FilterApproval(opts *bind.FilterOpts, _owner []common.Address, _spender []common.Address) (*PxcApprovalIterator, error) {

	var _ownerRule []interface{}
	for _, _ownerItem := range _owner {
		_ownerRule = append(_ownerRule, _ownerItem)
	}
	var _spenderRule []interface{}
	for _, _spenderItem := range _spender {
		_spenderRule = append(_spenderRule, _spenderItem)
	}

	logs, sub, err := _Pxc.contract.FilterLogs(opts, "Approval", _ownerRule, _spenderRule)
	if err != nil {
		return nil, err
	}
	return &PxcApprovalIterator{contract: _Pxc.contract, event: "Approval", logs: logs, sub: sub}, nil
}

// WatchApproval is a free log subscription operation binding the contract event 0x8c5be1e5ebec7d5bd14f71427d1e84f3dd0314c0f7b2291e5b200ac8c7c3b925.
//
// Solidity: e Approval(_owner indexed address, _spender indexed address, _value uint256)
func (_Pxc *PxcFilterer) WatchApproval(opts *bind.WatchOpts, sink chan<- *PxcApproval, _owner []common.Address, _spender []common.Address) (event.Subscription, error) {

	var _ownerRule []interface{}
	for _, _ownerItem := range _owner {
		_ownerRule = append(_ownerRule, _ownerItem)
	}
	var _spenderRule []interface{}
	for _, _spenderItem := range _spender {
		_spenderRule = append(_spenderRule, _spenderItem)
	}

	logs, sub, err := _Pxc.contract.WatchLogs(opts, "Approval", _ownerRule, _spenderRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(PxcApproval)
				if err := _Pxc.contract.UnpackLog(event, "Approval", log); err != nil {
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

// PxcTransferIterator is returned from FilterTransfer and is used to iterate over the raw logs and unpacked data for Transfer events raised by the Pxc contract.
type PxcTransferIterator struct {
	Event *PxcTransfer // Event containing the contract specifics and raw log

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
func (it *PxcTransferIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(PxcTransfer)
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
		it.Event = new(PxcTransfer)
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
func (it *PxcTransferIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *PxcTransferIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// PxcTransfer represents a Transfer event raised by the Pxc contract.
type PxcTransfer struct {
	From  common.Address
	To    common.Address
	Value *big.Int
	Raw   types.Log // Blockchain specific contextual infos
}

// FilterTransfer is a free log retrieval operation binding the contract event 0xddf252ad1be2c89b69c2b068fc378daa952ba7f163c4a11628f55a4df523b3ef.
//
// Solidity: e Transfer(_from indexed address, _to indexed address, _value uint256)
func (_Pxc *PxcFilterer) FilterTransfer(opts *bind.FilterOpts, _from []common.Address, _to []common.Address) (*PxcTransferIterator, error) {

	var _fromRule []interface{}
	for _, _fromItem := range _from {
		_fromRule = append(_fromRule, _fromItem)
	}
	var _toRule []interface{}
	for _, _toItem := range _to {
		_toRule = append(_toRule, _toItem)
	}

	logs, sub, err := _Pxc.contract.FilterLogs(opts, "Transfer", _fromRule, _toRule)
	if err != nil {
		return nil, err
	}
	return &PxcTransferIterator{contract: _Pxc.contract, event: "Transfer", logs: logs, sub: sub}, nil
}

// WatchTransfer is a free log subscription operation binding the contract event 0xddf252ad1be2c89b69c2b068fc378daa952ba7f163c4a11628f55a4df523b3ef.
//
// Solidity: e Transfer(_from indexed address, _to indexed address, _value uint256)
func (_Pxc *PxcFilterer) WatchTransfer(opts *bind.WatchOpts, sink chan<- *PxcTransfer, _from []common.Address, _to []common.Address) (event.Subscription, error) {

	var _fromRule []interface{}
	for _, _fromItem := range _from {
		_fromRule = append(_fromRule, _fromItem)
	}
	var _toRule []interface{}
	for _, _toItem := range _to {
		_toRule = append(_toRule, _toItem)
	}

	logs, sub, err := _Pxc.contract.WatchLogs(opts, "Transfer", _fromRule, _toRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(PxcTransfer)
				if err := _Pxc.contract.UnpackLog(event, "Transfer", log); err != nil {
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
