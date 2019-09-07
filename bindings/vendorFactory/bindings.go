// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package vendorfactory

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
	_ = abi.U256
	_ = bind.Bind
	_ = common.Big1
	_ = types.BloomLookup
	_ = event.NewSubscription
)

// VendorfactoryABI is the input ABI used to generate the binding from.
const VendorfactoryABI = "[{\"constant\":true,\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"name\":\"vendorContracts\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"name\":\"vendorContract\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[],\"name\":\"newVendor\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"name\":\"registeredVendors\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"}]"

// VendorfactoryBin is the compiled bytecode used for deploying new contracts.
var VendorfactoryBin = "0x608060405234801561001057600080fd5b50610df5806100206000396000f3fe608060405234801561001057600080fd5b506004361061004c5760003560e01c80631e6765c71461005157806340b90f4b1461008a578063c27f2112146100b0578063e7b8c718146100cc575b600080fd5b61006e6004803603602081101561006757600080fd5b50356100f2565b604080516001600160a01b039092168252519081900360200190f35b61006e600480360360208110156100a057600080fd5b50356001600160a01b0316610119565b6100b8610134565b604080519115158252519081900360200190f35b6100b8600480360360208110156100e257600080fd5b50356001600160a01b031661023a565b600081815481106100ff57fe5b6000918252602090912001546001600160a01b0316905081565b6002602052600090815260409020546001600160a01b031681565b3360009081526001602052604081205460ff1615610192576040805162461bcd60e51b81526020600482015260166024820152751b5d5cdd081b9bdd081899481c9959da5cdd195c995960521b604482015290519081900360640190fd5b60006040516101a09061024f565b604051809103906000f0801580156101bc573d6000803e3d6000fd5b5060008054600181810183557f290decd9548b62a8d60345a988386fc84ba6bc95484008f6362f93160ef3e56390910180546001600160a01b039094166001600160a01b0319948516811790915533835260208281526040808520805460ff1916851790556002909152909220805490931690911790915591505090565b60016020526000908152604090205460ff1681565b610b648061025d8339019056fe608060405234801561001057600080fd5b5033604051602001610022919061006b565b60408051601f198184030181529190528051602090910120600155600080546001600160a01b031916331790556100af565b61006561006082610080565b61009d565b82525050565b60006100778284610054565b50601401919050565b600061008b82610091565b92915050565b6001600160a01b031690565b600061008b82600061008b8260601b90565b610aa6806100be6000396000f3fe608060405234801561001057600080fd5b506004361061007d5760003560e01c80636c8e745c1161005b5780636c8e745c146100df5780638da5cb5b146100f2578063af640d0f14610107578063c43df6aa1461011c5761007d565b80630186a42314610082578063338a6d10146100ac5780635d85ed13146100cc575b600080fd5b610095610090366004610646565b61012f565b6040516100a3929190610932565b60405180910390f35b6100bf6100ba3660046106ff565b6101dd565b6040516100a391906108bd565b6100bf6100da3660046106ff565b61021a565b6100bf6100ed3660046106ff565b6102cd565b6100fa610371565b6040516100a391906108af565b61010f610380565b6040516100a391906108cb565b6100bf61012a366004610683565b610386565b805180820160209081018051600280835293830194830194909420939052825460408051601f6000196101006001861615020190931694909404918201839004830284018301905280835283918301828280156101cd5780601f106101a2576101008083540402835291602001916101cd565b820191906000526020600020905b8154815290600101906020018083116101b057829003601f168201915b5050505050908060010154905082565b8151602081840181018051600382529282019482019490942091909352815180830184018051928152908401929093019190912091525460ff1681565b60006102246104c5565b6102495760405162461bcd60e51b815260040161024090610952565b60405180910390fd5b60038360405161025991906108a3565b90815260200160405180910390208260405161027591906108a3565b908152604051908190036020018120805460ff191690557f6f87e5eed57f39feb4e7480e0eaa01353c1806e3a63675ebe5095b8f338cd62f906102bb908590859061090d565b60405180910390a15060015b92915050565b60006102d76104c5565b6102f35760405162461bcd60e51b815260040161024090610952565b600160038460405161030591906108a3565b90815260200160405180910390208360405161032191906108a3565b908152604051908190036020018120805492151560ff19909316929092179091557f20eca5c8f895538bd583d98f6dc289ffd029bf3d0a3b5362dffeace9ab6634dc906102bb908590859061090d565b6000546001600160a01b031681565b60015481565b60006103906104c5565b6103ac5760405162461bcd60e51b815260040161024090610952565b6040518060400160405280858152602001838152506002856040516103d191906108a3565b908152602001604051809103902060008201518160000190805190602001906103fb9291906104e8565b506020919091015160019091015560005b835181101561047e57600160038660405161042791906108a3565b908152602001604051809103902085838151811061044157fe5b602002602001015160405161045691906108a3565b908152604051908190036020019020805491151560ff1990921691909117905560010161040c565b507f3c16f324d383d6feb7b8939d45155634c4ba851efa4ea07b79cc280489899d1d8484846040516104b2939291906108d9565b60405180910390a15060015b9392505050565b600080546001600160a01b03163314156104e1575060016104e5565b5060005b90565b828054600181600116156101000203166002900490600052602060002090601f016020900481019282601f1061052957805160ff1916838001178555610556565b82800160010185558215610556579182015b8281111561055657825182559160200191906001019061053b565b50610562929150610566565b5090565b6104e591905b80821115610562576000815560010161056c565b600082601f83011261059157600080fd5b81356105a461059f82610989565b610962565b81815260209384019390925082018360005b838110156105e257813586016105cc88826105ec565b84525060209283019291909101906001016105b6565b5050505092915050565b600082601f8301126105fd57600080fd5b813561060b61059f826109aa565b9150808252602083016020830185838301111561062757600080fd5b610632838284610a06565b50505092915050565b80356102c781610a4c565b60006020828403121561065857600080fd5b813567ffffffffffffffff81111561066f57600080fd5b61067b848285016105ec565b949350505050565b60008060006060848603121561069857600080fd5b833567ffffffffffffffff8111156106af57600080fd5b6106bb868287016105ec565b935050602084013567ffffffffffffffff8111156106d857600080fd5b6106e486828701610580565b92505060406106f58682870161063b565b9150509250925092565b6000806040838503121561071257600080fd5b823567ffffffffffffffff81111561072957600080fd5b610735858286016105ec565b925050602083013567ffffffffffffffff81111561075257600080fd5b61075e858286016105ec565b9150509250929050565b60006104be8383610803565b61077d816109ea565b82525050565b600061078e826109d8565b61079881856109dc565b9350836020820285016107aa856109d2565b8060005b858110156107e457848403895281516107c78582610768565b94506107d2836109d2565b60209a909a01999250506001016107ae565b5091979650505050505050565b61077d816109f5565b61077d816104e5565b600061080e826109d8565b61081881856109dc565b9350610828818560208601610a12565b61083181610a42565b9093019392505050565b6000610846826109d8565b61085081856109e5565b9350610860818560208601610a12565b9290920192915050565b60006108776017836109dc565b7f63616c6c6572206d7573742062652076656e646f726564000000000000000000815260200192915050565b60006104be828461083b565b602081016102c78284610774565b602081016102c782846107f1565b602081016102c782846107fa565b606080825281016108ea8186610803565b905081810360208301526108fe8185610783565b905061067b60408301846107fa565b6040808252810161091e8185610803565b9050818103602083015261067b8184610803565b604080825281016109438185610803565b90506104be60208301846107fa565b602080825281016102c78161086a565b60405181810167ffffffffffffffff8111828210171561098157600080fd5b604052919050565b600067ffffffffffffffff8211156109a057600080fd5b5060209081020190565b600067ffffffffffffffff8211156109c157600080fd5b506020601f91909101601f19160190565b60200190565b5190565b90815260200190565b919050565b60006102c7826109fa565b151590565b6001600160a01b031690565b82818337506000910152565b60005b83811015610a2d578181015183820152602001610a15565b83811115610a3c576000848401525b50505050565b601f01601f191690565b610a55816104e5565b8114610a6057600080fd5b5056fea365627a7a7231582060a98dd25adb7f93363c18c1509ce6f354c7dd6cc6c2b950b6e4f2f4fb2f78b16c6578706572696d656e74616cf564736f6c634300050b0040a265627a7a72315820de57ba7f0347899918e83efce809a8a5be20eccc8eb1daf630b4a75c80a365a164736f6c634300050b0032"

// DeployVendorfactory deploys a new Ethereum contract, binding an instance of Vendorfactory to it.
func DeployVendorfactory(auth *bind.TransactOpts, backend bind.ContractBackend) (common.Address, *types.Transaction, *Vendorfactory, error) {
	parsed, err := abi.JSON(strings.NewReader(VendorfactoryABI))
	if err != nil {
		return common.Address{}, nil, nil, err
	}

	address, tx, contract, err := bind.DeployContract(auth, parsed, common.FromHex(VendorfactoryBin), backend)
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	return address, tx, &Vendorfactory{VendorfactoryCaller: VendorfactoryCaller{contract: contract}, VendorfactoryTransactor: VendorfactoryTransactor{contract: contract}, VendorfactoryFilterer: VendorfactoryFilterer{contract: contract}}, nil
}

// Vendorfactory is an auto generated Go binding around an Ethereum contract.
type Vendorfactory struct {
	VendorfactoryCaller     // Read-only binding to the contract
	VendorfactoryTransactor // Write-only binding to the contract
	VendorfactoryFilterer   // Log filterer for contract events
}

// VendorfactoryCaller is an auto generated read-only Go binding around an Ethereum contract.
type VendorfactoryCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// VendorfactoryTransactor is an auto generated write-only Go binding around an Ethereum contract.
type VendorfactoryTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// VendorfactoryFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type VendorfactoryFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// VendorfactorySession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type VendorfactorySession struct {
	Contract     *Vendorfactory    // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// VendorfactoryCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type VendorfactoryCallerSession struct {
	Contract *VendorfactoryCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts        // Call options to use throughout this session
}

// VendorfactoryTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type VendorfactoryTransactorSession struct {
	Contract     *VendorfactoryTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts        // Transaction auth options to use throughout this session
}

// VendorfactoryRaw is an auto generated low-level Go binding around an Ethereum contract.
type VendorfactoryRaw struct {
	Contract *Vendorfactory // Generic contract binding to access the raw methods on
}

// VendorfactoryCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type VendorfactoryCallerRaw struct {
	Contract *VendorfactoryCaller // Generic read-only contract binding to access the raw methods on
}

// VendorfactoryTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type VendorfactoryTransactorRaw struct {
	Contract *VendorfactoryTransactor // Generic write-only contract binding to access the raw methods on
}

// NewVendorfactory creates a new instance of Vendorfactory, bound to a specific deployed contract.
func NewVendorfactory(address common.Address, backend bind.ContractBackend) (*Vendorfactory, error) {
	contract, err := bindVendorfactory(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &Vendorfactory{VendorfactoryCaller: VendorfactoryCaller{contract: contract}, VendorfactoryTransactor: VendorfactoryTransactor{contract: contract}, VendorfactoryFilterer: VendorfactoryFilterer{contract: contract}}, nil
}

// NewVendorfactoryCaller creates a new read-only instance of Vendorfactory, bound to a specific deployed contract.
func NewVendorfactoryCaller(address common.Address, caller bind.ContractCaller) (*VendorfactoryCaller, error) {
	contract, err := bindVendorfactory(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &VendorfactoryCaller{contract: contract}, nil
}

// NewVendorfactoryTransactor creates a new write-only instance of Vendorfactory, bound to a specific deployed contract.
func NewVendorfactoryTransactor(address common.Address, transactor bind.ContractTransactor) (*VendorfactoryTransactor, error) {
	contract, err := bindVendorfactory(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &VendorfactoryTransactor{contract: contract}, nil
}

// NewVendorfactoryFilterer creates a new log filterer instance of Vendorfactory, bound to a specific deployed contract.
func NewVendorfactoryFilterer(address common.Address, filterer bind.ContractFilterer) (*VendorfactoryFilterer, error) {
	contract, err := bindVendorfactory(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &VendorfactoryFilterer{contract: contract}, nil
}

// bindVendorfactory binds a generic wrapper to an already deployed contract.
func bindVendorfactory(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(VendorfactoryABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Vendorfactory *VendorfactoryRaw) Call(opts *bind.CallOpts, result interface{}, method string, params ...interface{}) error {
	return _Vendorfactory.Contract.VendorfactoryCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Vendorfactory *VendorfactoryRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Vendorfactory.Contract.VendorfactoryTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Vendorfactory *VendorfactoryRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Vendorfactory.Contract.VendorfactoryTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Vendorfactory *VendorfactoryCallerRaw) Call(opts *bind.CallOpts, result interface{}, method string, params ...interface{}) error {
	return _Vendorfactory.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Vendorfactory *VendorfactoryTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Vendorfactory.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Vendorfactory *VendorfactoryTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Vendorfactory.Contract.contract.Transact(opts, method, params...)
}

// RegisteredVendors is a free data retrieval call binding the contract method 0xe7b8c718.
//
// Solidity: function registeredVendors(address ) constant returns(bool)
func (_Vendorfactory *VendorfactoryCaller) RegisteredVendors(opts *bind.CallOpts, arg0 common.Address) (bool, error) {
	var (
		ret0 = new(bool)
	)
	out := ret0
	err := _Vendorfactory.contract.Call(opts, out, "registeredVendors", arg0)
	return *ret0, err
}

// RegisteredVendors is a free data retrieval call binding the contract method 0xe7b8c718.
//
// Solidity: function registeredVendors(address ) constant returns(bool)
func (_Vendorfactory *VendorfactorySession) RegisteredVendors(arg0 common.Address) (bool, error) {
	return _Vendorfactory.Contract.RegisteredVendors(&_Vendorfactory.CallOpts, arg0)
}

// RegisteredVendors is a free data retrieval call binding the contract method 0xe7b8c718.
//
// Solidity: function registeredVendors(address ) constant returns(bool)
func (_Vendorfactory *VendorfactoryCallerSession) RegisteredVendors(arg0 common.Address) (bool, error) {
	return _Vendorfactory.Contract.RegisteredVendors(&_Vendorfactory.CallOpts, arg0)
}

// VendorContract is a free data retrieval call binding the contract method 0x40b90f4b.
//
// Solidity: function vendorContract(address ) constant returns(address)
func (_Vendorfactory *VendorfactoryCaller) VendorContract(opts *bind.CallOpts, arg0 common.Address) (common.Address, error) {
	var (
		ret0 = new(common.Address)
	)
	out := ret0
	err := _Vendorfactory.contract.Call(opts, out, "vendorContract", arg0)
	return *ret0, err
}

// VendorContract is a free data retrieval call binding the contract method 0x40b90f4b.
//
// Solidity: function vendorContract(address ) constant returns(address)
func (_Vendorfactory *VendorfactorySession) VendorContract(arg0 common.Address) (common.Address, error) {
	return _Vendorfactory.Contract.VendorContract(&_Vendorfactory.CallOpts, arg0)
}

// VendorContract is a free data retrieval call binding the contract method 0x40b90f4b.
//
// Solidity: function vendorContract(address ) constant returns(address)
func (_Vendorfactory *VendorfactoryCallerSession) VendorContract(arg0 common.Address) (common.Address, error) {
	return _Vendorfactory.Contract.VendorContract(&_Vendorfactory.CallOpts, arg0)
}

// VendorContracts is a free data retrieval call binding the contract method 0x1e6765c7.
//
// Solidity: function vendorContracts(uint256 ) constant returns(address)
func (_Vendorfactory *VendorfactoryCaller) VendorContracts(opts *bind.CallOpts, arg0 *big.Int) (common.Address, error) {
	var (
		ret0 = new(common.Address)
	)
	out := ret0
	err := _Vendorfactory.contract.Call(opts, out, "vendorContracts", arg0)
	return *ret0, err
}

// VendorContracts is a free data retrieval call binding the contract method 0x1e6765c7.
//
// Solidity: function vendorContracts(uint256 ) constant returns(address)
func (_Vendorfactory *VendorfactorySession) VendorContracts(arg0 *big.Int) (common.Address, error) {
	return _Vendorfactory.Contract.VendorContracts(&_Vendorfactory.CallOpts, arg0)
}

// VendorContracts is a free data retrieval call binding the contract method 0x1e6765c7.
//
// Solidity: function vendorContracts(uint256 ) constant returns(address)
func (_Vendorfactory *VendorfactoryCallerSession) VendorContracts(arg0 *big.Int) (common.Address, error) {
	return _Vendorfactory.Contract.VendorContracts(&_Vendorfactory.CallOpts, arg0)
}

// NewVendor is a paid mutator transaction binding the contract method 0xc27f2112.
//
// Solidity: function newVendor() returns(bool)
func (_Vendorfactory *VendorfactoryTransactor) NewVendor(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Vendorfactory.contract.Transact(opts, "newVendor")
}

// NewVendor is a paid mutator transaction binding the contract method 0xc27f2112.
//
// Solidity: function newVendor() returns(bool)
func (_Vendorfactory *VendorfactorySession) NewVendor() (*types.Transaction, error) {
	return _Vendorfactory.Contract.NewVendor(&_Vendorfactory.TransactOpts)
}

// NewVendor is a paid mutator transaction binding the contract method 0xc27f2112.
//
// Solidity: function newVendor() returns(bool)
func (_Vendorfactory *VendorfactoryTransactorSession) NewVendor() (*types.Transaction, error) {
	return _Vendorfactory.Contract.NewVendor(&_Vendorfactory.TransactOpts)
}
