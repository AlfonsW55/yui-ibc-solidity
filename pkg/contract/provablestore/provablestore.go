// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package provablestore

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

// ChannelCounterpartyData is an auto generated low-level Go binding around an user-defined struct.
type ChannelCounterpartyData struct {
	PortId    string
	ChannelId string
}

// ChannelData is an auto generated low-level Go binding around an user-defined struct.
type ChannelData struct {
	State          uint8
	Ordering       uint8
	Counterparty   ChannelCounterpartyData
	ConnectionHops []string
	Version        string
}

// ClientStateData is an auto generated low-level Go binding around an user-defined struct.
type ClientStateData struct {
	ChainId              string
	ProvableStoreAddress []byte
	LatestHeight         uint64
}

// ConnectionEndData is an auto generated low-level Go binding around an user-defined struct.
type ConnectionEndData struct {
	ClientId     string
	Versions     []VersionData
	State        uint8
	DelayPeriod  uint64
	Counterparty CounterpartyData
}

// ConsensusStateData is an auto generated low-level Go binding around an user-defined struct.
type ConsensusStateData struct {
	Timestamp  uint64
	Root       []byte
	Validators [][]byte
}

// CounterpartyData is an auto generated low-level Go binding around an user-defined struct.
type CounterpartyData struct {
	ClientId     string
	ConnectionId string
	Prefix       MerklePrefixData
}

// MerklePrefixData is an auto generated low-level Go binding around an user-defined struct.
type MerklePrefixData struct {
	KeyPrefix []byte
}

// VersionData is an auto generated low-level Go binding around an user-defined struct.
type VersionData struct {
	Identifier string
	Features   []string
}

// ProvablestoreABI is the input ABI used to generate the binding from.
const ProvablestoreABI = "[{\"inputs\":[{\"internalType\":\"string\",\"name\":\"clientId\",\"type\":\"string\"}],\"name\":\"clientCommitmentKey\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"stateMutability\":\"pure\",\"type\":\"function\",\"constant\":true},{\"inputs\":[{\"internalType\":\"string\",\"name\":\"clientId\",\"type\":\"string\"},{\"internalType\":\"uint64\",\"name\":\"height\",\"type\":\"uint64\"}],\"name\":\"consensusCommitmentKey\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"stateMutability\":\"pure\",\"type\":\"function\",\"constant\":true},{\"inputs\":[{\"internalType\":\"string\",\"name\":\"connectionId\",\"type\":\"string\"}],\"name\":\"connectionCommitmentKey\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"stateMutability\":\"pure\",\"type\":\"function\",\"constant\":true},{\"inputs\":[{\"internalType\":\"string\",\"name\":\"channelId\",\"type\":\"string\"}],\"name\":\"channelCommitmentKey\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"stateMutability\":\"pure\",\"type\":\"function\",\"constant\":true},{\"inputs\":[{\"internalType\":\"string\",\"name\":\"clientId\",\"type\":\"string\"}],\"name\":\"clientStateCommitmentSlot\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"stateMutability\":\"pure\",\"type\":\"function\",\"constant\":true},{\"inputs\":[{\"internalType\":\"string\",\"name\":\"clientId\",\"type\":\"string\"},{\"internalType\":\"uint64\",\"name\":\"height\",\"type\":\"uint64\"}],\"name\":\"consensusStateCommitmentSlot\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"stateMutability\":\"pure\",\"type\":\"function\",\"constant\":true},{\"inputs\":[{\"internalType\":\"string\",\"name\":\"connectionId\",\"type\":\"string\"}],\"name\":\"connectionCommitmentSlot\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"stateMutability\":\"pure\",\"type\":\"function\",\"constant\":true},{\"inputs\":[{\"internalType\":\"string\",\"name\":\"channelId\",\"type\":\"string\"}],\"name\":\"channelCommitmentSlot\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"stateMutability\":\"pure\",\"type\":\"function\",\"constant\":true},{\"inputs\":[{\"internalType\":\"string\",\"name\":\"clientId\",\"type\":\"string\"},{\"components\":[{\"internalType\":\"string\",\"name\":\"chain_id\",\"type\":\"string\"},{\"internalType\":\"bytes\",\"name\":\"provable_store_address\",\"type\":\"bytes\"},{\"internalType\":\"uint64\",\"name\":\"latest_height\",\"type\":\"uint64\"}],\"internalType\":\"structClientState.Data\",\"name\":\"data\",\"type\":\"tuple\"}],\"name\":\"setClientState\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"string\",\"name\":\"clientId\",\"type\":\"string\"}],\"name\":\"getClientState\",\"outputs\":[{\"components\":[{\"internalType\":\"string\",\"name\":\"chain_id\",\"type\":\"string\"},{\"internalType\":\"bytes\",\"name\":\"provable_store_address\",\"type\":\"bytes\"},{\"internalType\":\"uint64\",\"name\":\"latest_height\",\"type\":\"uint64\"}],\"internalType\":\"structClientState.Data\",\"name\":\"\",\"type\":\"tuple\"},{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\",\"constant\":true},{\"inputs\":[{\"internalType\":\"string\",\"name\":\"clientId\",\"type\":\"string\"}],\"name\":\"hasClientState\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\",\"constant\":true},{\"inputs\":[{\"internalType\":\"string\",\"name\":\"clientId\",\"type\":\"string\"},{\"internalType\":\"uint64\",\"name\":\"height\",\"type\":\"uint64\"},{\"components\":[{\"internalType\":\"uint64\",\"name\":\"timestamp\",\"type\":\"uint64\"},{\"internalType\":\"bytes\",\"name\":\"root\",\"type\":\"bytes\"},{\"internalType\":\"bytes[]\",\"name\":\"validators\",\"type\":\"bytes[]\"}],\"internalType\":\"structConsensusState.Data\",\"name\":\"consensusState\",\"type\":\"tuple\"}],\"name\":\"setConsensusState\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"string\",\"name\":\"clientId\",\"type\":\"string\"},{\"internalType\":\"uint64\",\"name\":\"height\",\"type\":\"uint64\"}],\"name\":\"getConsensusState\",\"outputs\":[{\"components\":[{\"internalType\":\"uint64\",\"name\":\"timestamp\",\"type\":\"uint64\"},{\"internalType\":\"bytes\",\"name\":\"root\",\"type\":\"bytes\"},{\"internalType\":\"bytes[]\",\"name\":\"validators\",\"type\":\"bytes[]\"}],\"internalType\":\"structConsensusState.Data\",\"name\":\"consensusState\",\"type\":\"tuple\"},{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\",\"constant\":true},{\"inputs\":[{\"internalType\":\"string\",\"name\":\"connectionId\",\"type\":\"string\"},{\"components\":[{\"internalType\":\"string\",\"name\":\"client_id\",\"type\":\"string\"},{\"components\":[{\"internalType\":\"string\",\"name\":\"identifier\",\"type\":\"string\"},{\"internalType\":\"string[]\",\"name\":\"features\",\"type\":\"string[]\"}],\"internalType\":\"structVersion.Data[]\",\"name\":\"versions\",\"type\":\"tuple[]\"},{\"internalType\":\"enumConnectionEnd.State\",\"name\":\"state\",\"type\":\"uint8\"},{\"internalType\":\"uint64\",\"name\":\"delay_period\",\"type\":\"uint64\"},{\"components\":[{\"internalType\":\"string\",\"name\":\"client_id\",\"type\":\"string\"},{\"internalType\":\"string\",\"name\":\"connection_id\",\"type\":\"string\"},{\"components\":[{\"internalType\":\"bytes\",\"name\":\"key_prefix\",\"type\":\"bytes\"}],\"internalType\":\"structMerklePrefix.Data\",\"name\":\"prefix\",\"type\":\"tuple\"}],\"internalType\":\"structCounterparty.Data\",\"name\":\"counterparty\",\"type\":\"tuple\"}],\"internalType\":\"structConnectionEnd.Data\",\"name\":\"connection\",\"type\":\"tuple\"}],\"name\":\"setConnection\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"string\",\"name\":\"connectionId\",\"type\":\"string\"}],\"name\":\"getConnection\",\"outputs\":[{\"components\":[{\"internalType\":\"string\",\"name\":\"client_id\",\"type\":\"string\"},{\"components\":[{\"internalType\":\"string\",\"name\":\"identifier\",\"type\":\"string\"},{\"internalType\":\"string[]\",\"name\":\"features\",\"type\":\"string[]\"}],\"internalType\":\"structVersion.Data[]\",\"name\":\"versions\",\"type\":\"tuple[]\"},{\"internalType\":\"enumConnectionEnd.State\",\"name\":\"state\",\"type\":\"uint8\"},{\"internalType\":\"uint64\",\"name\":\"delay_period\",\"type\":\"uint64\"},{\"components\":[{\"internalType\":\"string\",\"name\":\"client_id\",\"type\":\"string\"},{\"internalType\":\"string\",\"name\":\"connection_id\",\"type\":\"string\"},{\"components\":[{\"internalType\":\"bytes\",\"name\":\"key_prefix\",\"type\":\"bytes\"}],\"internalType\":\"structMerklePrefix.Data\",\"name\":\"prefix\",\"type\":\"tuple\"}],\"internalType\":\"structCounterparty.Data\",\"name\":\"counterparty\",\"type\":\"tuple\"}],\"internalType\":\"structConnectionEnd.Data\",\"name\":\"connection\",\"type\":\"tuple\"},{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\",\"constant\":true},{\"inputs\":[{\"internalType\":\"string\",\"name\":\"channelId\",\"type\":\"string\"},{\"components\":[{\"internalType\":\"enumChannel.State\",\"name\":\"state\",\"type\":\"uint8\"},{\"internalType\":\"enumChannel.Order\",\"name\":\"ordering\",\"type\":\"uint8\"},{\"components\":[{\"internalType\":\"string\",\"name\":\"port_id\",\"type\":\"string\"},{\"internalType\":\"string\",\"name\":\"channel_id\",\"type\":\"string\"}],\"internalType\":\"structChannelCounterparty.Data\",\"name\":\"counterparty\",\"type\":\"tuple\"},{\"internalType\":\"string[]\",\"name\":\"connection_hops\",\"type\":\"string[]\"},{\"internalType\":\"string\",\"name\":\"version\",\"type\":\"string\"}],\"internalType\":\"structChannel.Data\",\"name\":\"channel\",\"type\":\"tuple\"}],\"name\":\"setChannel\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"string\",\"name\":\"channelId\",\"type\":\"string\"}],\"name\":\"getChannel\",\"outputs\":[{\"components\":[{\"internalType\":\"enumChannel.State\",\"name\":\"state\",\"type\":\"uint8\"},{\"internalType\":\"enumChannel.Order\",\"name\":\"ordering\",\"type\":\"uint8\"},{\"components\":[{\"internalType\":\"string\",\"name\":\"port_id\",\"type\":\"string\"},{\"internalType\":\"string\",\"name\":\"channel_id\",\"type\":\"string\"}],\"internalType\":\"structChannelCounterparty.Data\",\"name\":\"counterparty\",\"type\":\"tuple\"},{\"internalType\":\"string[]\",\"name\":\"connection_hops\",\"type\":\"string[]\"},{\"internalType\":\"string\",\"name\":\"version\",\"type\":\"string\"}],\"internalType\":\"structChannel.Data\",\"name\":\"channel\",\"type\":\"tuple\"},{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\",\"constant\":true},{\"inputs\":[{\"internalType\":\"string\",\"name\":\"channelId\",\"type\":\"string\"}],\"name\":\"hasChannel\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\",\"constant\":true},{\"inputs\":[{\"internalType\":\"string\",\"name\":\"portId\",\"type\":\"string\"},{\"internalType\":\"string\",\"name\":\"channelId\",\"type\":\"string\"},{\"internalType\":\"uint64\",\"name\":\"sequence\",\"type\":\"uint64\"}],\"name\":\"setNextSequenceSend\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"string\",\"name\":\"portId\",\"type\":\"string\"},{\"internalType\":\"string\",\"name\":\"channelId\",\"type\":\"string\"},{\"internalType\":\"uint64\",\"name\":\"sequence\",\"type\":\"uint64\"}],\"name\":\"setNextSequenceRecv\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"string\",\"name\":\"portId\",\"type\":\"string\"},{\"internalType\":\"string\",\"name\":\"channelId\",\"type\":\"string\"},{\"internalType\":\"uint64\",\"name\":\"sequence\",\"type\":\"uint64\"}],\"name\":\"setNextSequenceAck\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"string\",\"name\":\"clientId\",\"type\":\"string\"}],\"name\":\"getClientStateBytes\",\"outputs\":[{\"internalType\":\"bytes\",\"name\":\"\",\"type\":\"bytes\"},{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\",\"constant\":true},{\"inputs\":[{\"internalType\":\"string\",\"name\":\"connectionId\",\"type\":\"string\"}],\"name\":\"getConnectionBytes\",\"outputs\":[{\"internalType\":\"bytes\",\"name\":\"\",\"type\":\"bytes\"},{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\",\"constant\":true},{\"inputs\":[{\"internalType\":\"bytes\",\"name\":\"connectionBytes\",\"type\":\"bytes\"}],\"name\":\"parseConnectionBytes\",\"outputs\":[{\"components\":[{\"internalType\":\"string\",\"name\":\"client_id\",\"type\":\"string\"},{\"components\":[{\"internalType\":\"string\",\"name\":\"identifier\",\"type\":\"string\"},{\"internalType\":\"string[]\",\"name\":\"features\",\"type\":\"string[]\"}],\"internalType\":\"structVersion.Data[]\",\"name\":\"versions\",\"type\":\"tuple[]\"},{\"internalType\":\"enumConnectionEnd.State\",\"name\":\"state\",\"type\":\"uint8\"},{\"internalType\":\"uint64\",\"name\":\"delay_period\",\"type\":\"uint64\"},{\"components\":[{\"internalType\":\"string\",\"name\":\"client_id\",\"type\":\"string\"},{\"internalType\":\"string\",\"name\":\"connection_id\",\"type\":\"string\"},{\"components\":[{\"internalType\":\"bytes\",\"name\":\"key_prefix\",\"type\":\"bytes\"}],\"internalType\":\"structMerklePrefix.Data\",\"name\":\"prefix\",\"type\":\"tuple\"}],\"internalType\":\"structCounterparty.Data\",\"name\":\"counterparty\",\"type\":\"tuple\"}],\"internalType\":\"structConnectionEnd.Data\",\"name\":\"connection\",\"type\":\"tuple\"}],\"stateMutability\":\"view\",\"type\":\"function\",\"constant\":true},{\"inputs\":[{\"internalType\":\"string\",\"name\":\"connectionId\",\"type\":\"string\"}],\"name\":\"getCommitment\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"stateMutability\":\"view\",\"type\":\"function\",\"constant\":true}]"

// Provablestore is an auto generated Go binding around an Ethereum contract.
type Provablestore struct {
	ProvablestoreCaller     // Read-only binding to the contract
	ProvablestoreTransactor // Write-only binding to the contract
	ProvablestoreFilterer   // Log filterer for contract events
}

// ProvablestoreCaller is an auto generated read-only Go binding around an Ethereum contract.
type ProvablestoreCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// ProvablestoreTransactor is an auto generated write-only Go binding around an Ethereum contract.
type ProvablestoreTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// ProvablestoreFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type ProvablestoreFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// ProvablestoreSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type ProvablestoreSession struct {
	Contract     *Provablestore    // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// ProvablestoreCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type ProvablestoreCallerSession struct {
	Contract *ProvablestoreCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts        // Call options to use throughout this session
}

// ProvablestoreTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type ProvablestoreTransactorSession struct {
	Contract     *ProvablestoreTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts        // Transaction auth options to use throughout this session
}

// ProvablestoreRaw is an auto generated low-level Go binding around an Ethereum contract.
type ProvablestoreRaw struct {
	Contract *Provablestore // Generic contract binding to access the raw methods on
}

// ProvablestoreCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type ProvablestoreCallerRaw struct {
	Contract *ProvablestoreCaller // Generic read-only contract binding to access the raw methods on
}

// ProvablestoreTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type ProvablestoreTransactorRaw struct {
	Contract *ProvablestoreTransactor // Generic write-only contract binding to access the raw methods on
}

// NewProvablestore creates a new instance of Provablestore, bound to a specific deployed contract.
func NewProvablestore(address common.Address, backend bind.ContractBackend) (*Provablestore, error) {
	contract, err := bindProvablestore(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &Provablestore{ProvablestoreCaller: ProvablestoreCaller{contract: contract}, ProvablestoreTransactor: ProvablestoreTransactor{contract: contract}, ProvablestoreFilterer: ProvablestoreFilterer{contract: contract}}, nil
}

// NewProvablestoreCaller creates a new read-only instance of Provablestore, bound to a specific deployed contract.
func NewProvablestoreCaller(address common.Address, caller bind.ContractCaller) (*ProvablestoreCaller, error) {
	contract, err := bindProvablestore(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &ProvablestoreCaller{contract: contract}, nil
}

// NewProvablestoreTransactor creates a new write-only instance of Provablestore, bound to a specific deployed contract.
func NewProvablestoreTransactor(address common.Address, transactor bind.ContractTransactor) (*ProvablestoreTransactor, error) {
	contract, err := bindProvablestore(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &ProvablestoreTransactor{contract: contract}, nil
}

// NewProvablestoreFilterer creates a new log filterer instance of Provablestore, bound to a specific deployed contract.
func NewProvablestoreFilterer(address common.Address, filterer bind.ContractFilterer) (*ProvablestoreFilterer, error) {
	contract, err := bindProvablestore(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &ProvablestoreFilterer{contract: contract}, nil
}

// bindProvablestore binds a generic wrapper to an already deployed contract.
func bindProvablestore(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(ProvablestoreABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Provablestore *ProvablestoreRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _Provablestore.Contract.ProvablestoreCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Provablestore *ProvablestoreRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Provablestore.Contract.ProvablestoreTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Provablestore *ProvablestoreRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Provablestore.Contract.ProvablestoreTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Provablestore *ProvablestoreCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _Provablestore.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Provablestore *ProvablestoreTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Provablestore.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Provablestore *ProvablestoreTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Provablestore.Contract.contract.Transact(opts, method, params...)
}

// ChannelCommitmentKey is a free data retrieval call binding the contract method 0x1f676483.
//
// Solidity: function channelCommitmentKey(string channelId) pure returns(bytes32)
func (_Provablestore *ProvablestoreCaller) ChannelCommitmentKey(opts *bind.CallOpts, channelId string) ([32]byte, error) {
	var out []interface{}
	err := _Provablestore.contract.Call(opts, &out, "channelCommitmentKey", channelId)

	if err != nil {
		return *new([32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)

	return out0, err

}

// ChannelCommitmentKey is a free data retrieval call binding the contract method 0x1f676483.
//
// Solidity: function channelCommitmentKey(string channelId) pure returns(bytes32)
func (_Provablestore *ProvablestoreSession) ChannelCommitmentKey(channelId string) ([32]byte, error) {
	return _Provablestore.Contract.ChannelCommitmentKey(&_Provablestore.CallOpts, channelId)
}

// ChannelCommitmentKey is a free data retrieval call binding the contract method 0x1f676483.
//
// Solidity: function channelCommitmentKey(string channelId) pure returns(bytes32)
func (_Provablestore *ProvablestoreCallerSession) ChannelCommitmentKey(channelId string) ([32]byte, error) {
	return _Provablestore.Contract.ChannelCommitmentKey(&_Provablestore.CallOpts, channelId)
}

// ChannelCommitmentSlot is a free data retrieval call binding the contract method 0xb78af694.
//
// Solidity: function channelCommitmentSlot(string channelId) pure returns(bytes32)
func (_Provablestore *ProvablestoreCaller) ChannelCommitmentSlot(opts *bind.CallOpts, channelId string) ([32]byte, error) {
	var out []interface{}
	err := _Provablestore.contract.Call(opts, &out, "channelCommitmentSlot", channelId)

	if err != nil {
		return *new([32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)

	return out0, err

}

// ChannelCommitmentSlot is a free data retrieval call binding the contract method 0xb78af694.
//
// Solidity: function channelCommitmentSlot(string channelId) pure returns(bytes32)
func (_Provablestore *ProvablestoreSession) ChannelCommitmentSlot(channelId string) ([32]byte, error) {
	return _Provablestore.Contract.ChannelCommitmentSlot(&_Provablestore.CallOpts, channelId)
}

// ChannelCommitmentSlot is a free data retrieval call binding the contract method 0xb78af694.
//
// Solidity: function channelCommitmentSlot(string channelId) pure returns(bytes32)
func (_Provablestore *ProvablestoreCallerSession) ChannelCommitmentSlot(channelId string) ([32]byte, error) {
	return _Provablestore.Contract.ChannelCommitmentSlot(&_Provablestore.CallOpts, channelId)
}

// ClientCommitmentKey is a free data retrieval call binding the contract method 0xc7ddc606.
//
// Solidity: function clientCommitmentKey(string clientId) pure returns(bytes32)
func (_Provablestore *ProvablestoreCaller) ClientCommitmentKey(opts *bind.CallOpts, clientId string) ([32]byte, error) {
	var out []interface{}
	err := _Provablestore.contract.Call(opts, &out, "clientCommitmentKey", clientId)

	if err != nil {
		return *new([32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)

	return out0, err

}

// ClientCommitmentKey is a free data retrieval call binding the contract method 0xc7ddc606.
//
// Solidity: function clientCommitmentKey(string clientId) pure returns(bytes32)
func (_Provablestore *ProvablestoreSession) ClientCommitmentKey(clientId string) ([32]byte, error) {
	return _Provablestore.Contract.ClientCommitmentKey(&_Provablestore.CallOpts, clientId)
}

// ClientCommitmentKey is a free data retrieval call binding the contract method 0xc7ddc606.
//
// Solidity: function clientCommitmentKey(string clientId) pure returns(bytes32)
func (_Provablestore *ProvablestoreCallerSession) ClientCommitmentKey(clientId string) ([32]byte, error) {
	return _Provablestore.Contract.ClientCommitmentKey(&_Provablestore.CallOpts, clientId)
}

// ClientStateCommitmentSlot is a free data retrieval call binding the contract method 0xb0cbb120.
//
// Solidity: function clientStateCommitmentSlot(string clientId) pure returns(bytes32)
func (_Provablestore *ProvablestoreCaller) ClientStateCommitmentSlot(opts *bind.CallOpts, clientId string) ([32]byte, error) {
	var out []interface{}
	err := _Provablestore.contract.Call(opts, &out, "clientStateCommitmentSlot", clientId)

	if err != nil {
		return *new([32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)

	return out0, err

}

// ClientStateCommitmentSlot is a free data retrieval call binding the contract method 0xb0cbb120.
//
// Solidity: function clientStateCommitmentSlot(string clientId) pure returns(bytes32)
func (_Provablestore *ProvablestoreSession) ClientStateCommitmentSlot(clientId string) ([32]byte, error) {
	return _Provablestore.Contract.ClientStateCommitmentSlot(&_Provablestore.CallOpts, clientId)
}

// ClientStateCommitmentSlot is a free data retrieval call binding the contract method 0xb0cbb120.
//
// Solidity: function clientStateCommitmentSlot(string clientId) pure returns(bytes32)
func (_Provablestore *ProvablestoreCallerSession) ClientStateCommitmentSlot(clientId string) ([32]byte, error) {
	return _Provablestore.Contract.ClientStateCommitmentSlot(&_Provablestore.CallOpts, clientId)
}

// ConnectionCommitmentKey is a free data retrieval call binding the contract method 0xa9dd3eb3.
//
// Solidity: function connectionCommitmentKey(string connectionId) pure returns(bytes32)
func (_Provablestore *ProvablestoreCaller) ConnectionCommitmentKey(opts *bind.CallOpts, connectionId string) ([32]byte, error) {
	var out []interface{}
	err := _Provablestore.contract.Call(opts, &out, "connectionCommitmentKey", connectionId)

	if err != nil {
		return *new([32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)

	return out0, err

}

// ConnectionCommitmentKey is a free data retrieval call binding the contract method 0xa9dd3eb3.
//
// Solidity: function connectionCommitmentKey(string connectionId) pure returns(bytes32)
func (_Provablestore *ProvablestoreSession) ConnectionCommitmentKey(connectionId string) ([32]byte, error) {
	return _Provablestore.Contract.ConnectionCommitmentKey(&_Provablestore.CallOpts, connectionId)
}

// ConnectionCommitmentKey is a free data retrieval call binding the contract method 0xa9dd3eb3.
//
// Solidity: function connectionCommitmentKey(string connectionId) pure returns(bytes32)
func (_Provablestore *ProvablestoreCallerSession) ConnectionCommitmentKey(connectionId string) ([32]byte, error) {
	return _Provablestore.Contract.ConnectionCommitmentKey(&_Provablestore.CallOpts, connectionId)
}

// ConnectionCommitmentSlot is a free data retrieval call binding the contract method 0x8b89bf24.
//
// Solidity: function connectionCommitmentSlot(string connectionId) pure returns(bytes32)
func (_Provablestore *ProvablestoreCaller) ConnectionCommitmentSlot(opts *bind.CallOpts, connectionId string) ([32]byte, error) {
	var out []interface{}
	err := _Provablestore.contract.Call(opts, &out, "connectionCommitmentSlot", connectionId)

	if err != nil {
		return *new([32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)

	return out0, err

}

// ConnectionCommitmentSlot is a free data retrieval call binding the contract method 0x8b89bf24.
//
// Solidity: function connectionCommitmentSlot(string connectionId) pure returns(bytes32)
func (_Provablestore *ProvablestoreSession) ConnectionCommitmentSlot(connectionId string) ([32]byte, error) {
	return _Provablestore.Contract.ConnectionCommitmentSlot(&_Provablestore.CallOpts, connectionId)
}

// ConnectionCommitmentSlot is a free data retrieval call binding the contract method 0x8b89bf24.
//
// Solidity: function connectionCommitmentSlot(string connectionId) pure returns(bytes32)
func (_Provablestore *ProvablestoreCallerSession) ConnectionCommitmentSlot(connectionId string) ([32]byte, error) {
	return _Provablestore.Contract.ConnectionCommitmentSlot(&_Provablestore.CallOpts, connectionId)
}

// ConsensusCommitmentKey is a free data retrieval call binding the contract method 0xc1f0b643.
//
// Solidity: function consensusCommitmentKey(string clientId, uint64 height) pure returns(bytes32)
func (_Provablestore *ProvablestoreCaller) ConsensusCommitmentKey(opts *bind.CallOpts, clientId string, height uint64) ([32]byte, error) {
	var out []interface{}
	err := _Provablestore.contract.Call(opts, &out, "consensusCommitmentKey", clientId, height)

	if err != nil {
		return *new([32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)

	return out0, err

}

// ConsensusCommitmentKey is a free data retrieval call binding the contract method 0xc1f0b643.
//
// Solidity: function consensusCommitmentKey(string clientId, uint64 height) pure returns(bytes32)
func (_Provablestore *ProvablestoreSession) ConsensusCommitmentKey(clientId string, height uint64) ([32]byte, error) {
	return _Provablestore.Contract.ConsensusCommitmentKey(&_Provablestore.CallOpts, clientId, height)
}

// ConsensusCommitmentKey is a free data retrieval call binding the contract method 0xc1f0b643.
//
// Solidity: function consensusCommitmentKey(string clientId, uint64 height) pure returns(bytes32)
func (_Provablestore *ProvablestoreCallerSession) ConsensusCommitmentKey(clientId string, height uint64) ([32]byte, error) {
	return _Provablestore.Contract.ConsensusCommitmentKey(&_Provablestore.CallOpts, clientId, height)
}

// ConsensusStateCommitmentSlot is a free data retrieval call binding the contract method 0xad30116c.
//
// Solidity: function consensusStateCommitmentSlot(string clientId, uint64 height) pure returns(bytes32)
func (_Provablestore *ProvablestoreCaller) ConsensusStateCommitmentSlot(opts *bind.CallOpts, clientId string, height uint64) ([32]byte, error) {
	var out []interface{}
	err := _Provablestore.contract.Call(opts, &out, "consensusStateCommitmentSlot", clientId, height)

	if err != nil {
		return *new([32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)

	return out0, err

}

// ConsensusStateCommitmentSlot is a free data retrieval call binding the contract method 0xad30116c.
//
// Solidity: function consensusStateCommitmentSlot(string clientId, uint64 height) pure returns(bytes32)
func (_Provablestore *ProvablestoreSession) ConsensusStateCommitmentSlot(clientId string, height uint64) ([32]byte, error) {
	return _Provablestore.Contract.ConsensusStateCommitmentSlot(&_Provablestore.CallOpts, clientId, height)
}

// ConsensusStateCommitmentSlot is a free data retrieval call binding the contract method 0xad30116c.
//
// Solidity: function consensusStateCommitmentSlot(string clientId, uint64 height) pure returns(bytes32)
func (_Provablestore *ProvablestoreCallerSession) ConsensusStateCommitmentSlot(clientId string, height uint64) ([32]byte, error) {
	return _Provablestore.Contract.ConsensusStateCommitmentSlot(&_Provablestore.CallOpts, clientId, height)
}

// GetChannel is a free data retrieval call binding the contract method 0x7cd7ee3d.
//
// Solidity: function getChannel(string channelId) view returns((uint8,uint8,(string,string),string[],string) channel, bool)
func (_Provablestore *ProvablestoreCaller) GetChannel(opts *bind.CallOpts, channelId string) (ChannelData, bool, error) {
	var out []interface{}
	err := _Provablestore.contract.Call(opts, &out, "getChannel", channelId)

	if err != nil {
		return *new(ChannelData), *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(ChannelData)).(*ChannelData)
	out1 := *abi.ConvertType(out[1], new(bool)).(*bool)

	return out0, out1, err

}

// GetChannel is a free data retrieval call binding the contract method 0x7cd7ee3d.
//
// Solidity: function getChannel(string channelId) view returns((uint8,uint8,(string,string),string[],string) channel, bool)
func (_Provablestore *ProvablestoreSession) GetChannel(channelId string) (ChannelData, bool, error) {
	return _Provablestore.Contract.GetChannel(&_Provablestore.CallOpts, channelId)
}

// GetChannel is a free data retrieval call binding the contract method 0x7cd7ee3d.
//
// Solidity: function getChannel(string channelId) view returns((uint8,uint8,(string,string),string[],string) channel, bool)
func (_Provablestore *ProvablestoreCallerSession) GetChannel(channelId string) (ChannelData, bool, error) {
	return _Provablestore.Contract.GetChannel(&_Provablestore.CallOpts, channelId)
}

// GetClientState is a free data retrieval call binding the contract method 0x76c81c42.
//
// Solidity: function getClientState(string clientId) view returns((string,bytes,uint64), bool)
func (_Provablestore *ProvablestoreCaller) GetClientState(opts *bind.CallOpts, clientId string) (ClientStateData, bool, error) {
	var out []interface{}
	err := _Provablestore.contract.Call(opts, &out, "getClientState", clientId)

	if err != nil {
		return *new(ClientStateData), *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(ClientStateData)).(*ClientStateData)
	out1 := *abi.ConvertType(out[1], new(bool)).(*bool)

	return out0, out1, err

}

// GetClientState is a free data retrieval call binding the contract method 0x76c81c42.
//
// Solidity: function getClientState(string clientId) view returns((string,bytes,uint64), bool)
func (_Provablestore *ProvablestoreSession) GetClientState(clientId string) (ClientStateData, bool, error) {
	return _Provablestore.Contract.GetClientState(&_Provablestore.CallOpts, clientId)
}

// GetClientState is a free data retrieval call binding the contract method 0x76c81c42.
//
// Solidity: function getClientState(string clientId) view returns((string,bytes,uint64), bool)
func (_Provablestore *ProvablestoreCallerSession) GetClientState(clientId string) (ClientStateData, bool, error) {
	return _Provablestore.Contract.GetClientState(&_Provablestore.CallOpts, clientId)
}

// GetClientStateBytes is a free data retrieval call binding the contract method 0x78e9a180.
//
// Solidity: function getClientStateBytes(string clientId) view returns(bytes, bool)
func (_Provablestore *ProvablestoreCaller) GetClientStateBytes(opts *bind.CallOpts, clientId string) ([]byte, bool, error) {
	var out []interface{}
	err := _Provablestore.contract.Call(opts, &out, "getClientStateBytes", clientId)

	if err != nil {
		return *new([]byte), *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new([]byte)).(*[]byte)
	out1 := *abi.ConvertType(out[1], new(bool)).(*bool)

	return out0, out1, err

}

// GetClientStateBytes is a free data retrieval call binding the contract method 0x78e9a180.
//
// Solidity: function getClientStateBytes(string clientId) view returns(bytes, bool)
func (_Provablestore *ProvablestoreSession) GetClientStateBytes(clientId string) ([]byte, bool, error) {
	return _Provablestore.Contract.GetClientStateBytes(&_Provablestore.CallOpts, clientId)
}

// GetClientStateBytes is a free data retrieval call binding the contract method 0x78e9a180.
//
// Solidity: function getClientStateBytes(string clientId) view returns(bytes, bool)
func (_Provablestore *ProvablestoreCallerSession) GetClientStateBytes(clientId string) ([]byte, bool, error) {
	return _Provablestore.Contract.GetClientStateBytes(&_Provablestore.CallOpts, clientId)
}

// GetCommitment is a free data retrieval call binding the contract method 0xa7b09e2a.
//
// Solidity: function getCommitment(string connectionId) view returns(bytes32)
func (_Provablestore *ProvablestoreCaller) GetCommitment(opts *bind.CallOpts, connectionId string) ([32]byte, error) {
	var out []interface{}
	err := _Provablestore.contract.Call(opts, &out, "getCommitment", connectionId)

	if err != nil {
		return *new([32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)

	return out0, err

}

// GetCommitment is a free data retrieval call binding the contract method 0xa7b09e2a.
//
// Solidity: function getCommitment(string connectionId) view returns(bytes32)
func (_Provablestore *ProvablestoreSession) GetCommitment(connectionId string) ([32]byte, error) {
	return _Provablestore.Contract.GetCommitment(&_Provablestore.CallOpts, connectionId)
}

// GetCommitment is a free data retrieval call binding the contract method 0xa7b09e2a.
//
// Solidity: function getCommitment(string connectionId) view returns(bytes32)
func (_Provablestore *ProvablestoreCallerSession) GetCommitment(connectionId string) ([32]byte, error) {
	return _Provablestore.Contract.GetCommitment(&_Provablestore.CallOpts, connectionId)
}

// GetConnection is a free data retrieval call binding the contract method 0x27711a69.
//
// Solidity: function getConnection(string connectionId) view returns((string,(string,string[])[],uint8,uint64,(string,string,(bytes))) connection, bool)
func (_Provablestore *ProvablestoreCaller) GetConnection(opts *bind.CallOpts, connectionId string) (ConnectionEndData, bool, error) {
	var out []interface{}
	err := _Provablestore.contract.Call(opts, &out, "getConnection", connectionId)

	if err != nil {
		return *new(ConnectionEndData), *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(ConnectionEndData)).(*ConnectionEndData)
	out1 := *abi.ConvertType(out[1], new(bool)).(*bool)

	return out0, out1, err

}

// GetConnection is a free data retrieval call binding the contract method 0x27711a69.
//
// Solidity: function getConnection(string connectionId) view returns((string,(string,string[])[],uint8,uint64,(string,string,(bytes))) connection, bool)
func (_Provablestore *ProvablestoreSession) GetConnection(connectionId string) (ConnectionEndData, bool, error) {
	return _Provablestore.Contract.GetConnection(&_Provablestore.CallOpts, connectionId)
}

// GetConnection is a free data retrieval call binding the contract method 0x27711a69.
//
// Solidity: function getConnection(string connectionId) view returns((string,(string,string[])[],uint8,uint64,(string,string,(bytes))) connection, bool)
func (_Provablestore *ProvablestoreCallerSession) GetConnection(connectionId string) (ConnectionEndData, bool, error) {
	return _Provablestore.Contract.GetConnection(&_Provablestore.CallOpts, connectionId)
}

// GetConnectionBytes is a free data retrieval call binding the contract method 0xc6c886d1.
//
// Solidity: function getConnectionBytes(string connectionId) view returns(bytes, bool)
func (_Provablestore *ProvablestoreCaller) GetConnectionBytes(opts *bind.CallOpts, connectionId string) ([]byte, bool, error) {
	var out []interface{}
	err := _Provablestore.contract.Call(opts, &out, "getConnectionBytes", connectionId)

	if err != nil {
		return *new([]byte), *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new([]byte)).(*[]byte)
	out1 := *abi.ConvertType(out[1], new(bool)).(*bool)

	return out0, out1, err

}

// GetConnectionBytes is a free data retrieval call binding the contract method 0xc6c886d1.
//
// Solidity: function getConnectionBytes(string connectionId) view returns(bytes, bool)
func (_Provablestore *ProvablestoreSession) GetConnectionBytes(connectionId string) ([]byte, bool, error) {
	return _Provablestore.Contract.GetConnectionBytes(&_Provablestore.CallOpts, connectionId)
}

// GetConnectionBytes is a free data retrieval call binding the contract method 0xc6c886d1.
//
// Solidity: function getConnectionBytes(string connectionId) view returns(bytes, bool)
func (_Provablestore *ProvablestoreCallerSession) GetConnectionBytes(connectionId string) ([]byte, bool, error) {
	return _Provablestore.Contract.GetConnectionBytes(&_Provablestore.CallOpts, connectionId)
}

// GetConsensusState is a free data retrieval call binding the contract method 0xa37a45c8.
//
// Solidity: function getConsensusState(string clientId, uint64 height) view returns((uint64,bytes,bytes[]) consensusState, bool)
func (_Provablestore *ProvablestoreCaller) GetConsensusState(opts *bind.CallOpts, clientId string, height uint64) (ConsensusStateData, bool, error) {
	var out []interface{}
	err := _Provablestore.contract.Call(opts, &out, "getConsensusState", clientId, height)

	if err != nil {
		return *new(ConsensusStateData), *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(ConsensusStateData)).(*ConsensusStateData)
	out1 := *abi.ConvertType(out[1], new(bool)).(*bool)

	return out0, out1, err

}

// GetConsensusState is a free data retrieval call binding the contract method 0xa37a45c8.
//
// Solidity: function getConsensusState(string clientId, uint64 height) view returns((uint64,bytes,bytes[]) consensusState, bool)
func (_Provablestore *ProvablestoreSession) GetConsensusState(clientId string, height uint64) (ConsensusStateData, bool, error) {
	return _Provablestore.Contract.GetConsensusState(&_Provablestore.CallOpts, clientId, height)
}

// GetConsensusState is a free data retrieval call binding the contract method 0xa37a45c8.
//
// Solidity: function getConsensusState(string clientId, uint64 height) view returns((uint64,bytes,bytes[]) consensusState, bool)
func (_Provablestore *ProvablestoreCallerSession) GetConsensusState(clientId string, height uint64) (ConsensusStateData, bool, error) {
	return _Provablestore.Contract.GetConsensusState(&_Provablestore.CallOpts, clientId, height)
}

// HasChannel is a free data retrieval call binding the contract method 0x726f7844.
//
// Solidity: function hasChannel(string channelId) view returns(bool)
func (_Provablestore *ProvablestoreCaller) HasChannel(opts *bind.CallOpts, channelId string) (bool, error) {
	var out []interface{}
	err := _Provablestore.contract.Call(opts, &out, "hasChannel", channelId)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// HasChannel is a free data retrieval call binding the contract method 0x726f7844.
//
// Solidity: function hasChannel(string channelId) view returns(bool)
func (_Provablestore *ProvablestoreSession) HasChannel(channelId string) (bool, error) {
	return _Provablestore.Contract.HasChannel(&_Provablestore.CallOpts, channelId)
}

// HasChannel is a free data retrieval call binding the contract method 0x726f7844.
//
// Solidity: function hasChannel(string channelId) view returns(bool)
func (_Provablestore *ProvablestoreCallerSession) HasChannel(channelId string) (bool, error) {
	return _Provablestore.Contract.HasChannel(&_Provablestore.CallOpts, channelId)
}

// HasClientState is a free data retrieval call binding the contract method 0x41a879d8.
//
// Solidity: function hasClientState(string clientId) view returns(bool)
func (_Provablestore *ProvablestoreCaller) HasClientState(opts *bind.CallOpts, clientId string) (bool, error) {
	var out []interface{}
	err := _Provablestore.contract.Call(opts, &out, "hasClientState", clientId)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// HasClientState is a free data retrieval call binding the contract method 0x41a879d8.
//
// Solidity: function hasClientState(string clientId) view returns(bool)
func (_Provablestore *ProvablestoreSession) HasClientState(clientId string) (bool, error) {
	return _Provablestore.Contract.HasClientState(&_Provablestore.CallOpts, clientId)
}

// HasClientState is a free data retrieval call binding the contract method 0x41a879d8.
//
// Solidity: function hasClientState(string clientId) view returns(bool)
func (_Provablestore *ProvablestoreCallerSession) HasClientState(clientId string) (bool, error) {
	return _Provablestore.Contract.HasClientState(&_Provablestore.CallOpts, clientId)
}

// ParseConnectionBytes is a free data retrieval call binding the contract method 0xd9800869.
//
// Solidity: function parseConnectionBytes(bytes connectionBytes) view returns((string,(string,string[])[],uint8,uint64,(string,string,(bytes))) connection)
func (_Provablestore *ProvablestoreCaller) ParseConnectionBytes(opts *bind.CallOpts, connectionBytes []byte) (ConnectionEndData, error) {
	var out []interface{}
	err := _Provablestore.contract.Call(opts, &out, "parseConnectionBytes", connectionBytes)

	if err != nil {
		return *new(ConnectionEndData), err
	}

	out0 := *abi.ConvertType(out[0], new(ConnectionEndData)).(*ConnectionEndData)

	return out0, err

}

// ParseConnectionBytes is a free data retrieval call binding the contract method 0xd9800869.
//
// Solidity: function parseConnectionBytes(bytes connectionBytes) view returns((string,(string,string[])[],uint8,uint64,(string,string,(bytes))) connection)
func (_Provablestore *ProvablestoreSession) ParseConnectionBytes(connectionBytes []byte) (ConnectionEndData, error) {
	return _Provablestore.Contract.ParseConnectionBytes(&_Provablestore.CallOpts, connectionBytes)
}

// ParseConnectionBytes is a free data retrieval call binding the contract method 0xd9800869.
//
// Solidity: function parseConnectionBytes(bytes connectionBytes) view returns((string,(string,string[])[],uint8,uint64,(string,string,(bytes))) connection)
func (_Provablestore *ProvablestoreCallerSession) ParseConnectionBytes(connectionBytes []byte) (ConnectionEndData, error) {
	return _Provablestore.Contract.ParseConnectionBytes(&_Provablestore.CallOpts, connectionBytes)
}

// SetChannel is a paid mutator transaction binding the contract method 0xd8aa58f1.
//
// Solidity: function setChannel(string channelId, (uint8,uint8,(string,string),string[],string) channel) returns()
func (_Provablestore *ProvablestoreTransactor) SetChannel(opts *bind.TransactOpts, channelId string, channel ChannelData) (*types.Transaction, error) {
	return _Provablestore.contract.Transact(opts, "setChannel", channelId, channel)
}

// SetChannel is a paid mutator transaction binding the contract method 0xd8aa58f1.
//
// Solidity: function setChannel(string channelId, (uint8,uint8,(string,string),string[],string) channel) returns()
func (_Provablestore *ProvablestoreSession) SetChannel(channelId string, channel ChannelData) (*types.Transaction, error) {
	return _Provablestore.Contract.SetChannel(&_Provablestore.TransactOpts, channelId, channel)
}

// SetChannel is a paid mutator transaction binding the contract method 0xd8aa58f1.
//
// Solidity: function setChannel(string channelId, (uint8,uint8,(string,string),string[],string) channel) returns()
func (_Provablestore *ProvablestoreTransactorSession) SetChannel(channelId string, channel ChannelData) (*types.Transaction, error) {
	return _Provablestore.Contract.SetChannel(&_Provablestore.TransactOpts, channelId, channel)
}

// SetClientState is a paid mutator transaction binding the contract method 0xe0ca210d.
//
// Solidity: function setClientState(string clientId, (string,bytes,uint64) data) returns()
func (_Provablestore *ProvablestoreTransactor) SetClientState(opts *bind.TransactOpts, clientId string, data ClientStateData) (*types.Transaction, error) {
	return _Provablestore.contract.Transact(opts, "setClientState", clientId, data)
}

// SetClientState is a paid mutator transaction binding the contract method 0xe0ca210d.
//
// Solidity: function setClientState(string clientId, (string,bytes,uint64) data) returns()
func (_Provablestore *ProvablestoreSession) SetClientState(clientId string, data ClientStateData) (*types.Transaction, error) {
	return _Provablestore.Contract.SetClientState(&_Provablestore.TransactOpts, clientId, data)
}

// SetClientState is a paid mutator transaction binding the contract method 0xe0ca210d.
//
// Solidity: function setClientState(string clientId, (string,bytes,uint64) data) returns()
func (_Provablestore *ProvablestoreTransactorSession) SetClientState(clientId string, data ClientStateData) (*types.Transaction, error) {
	return _Provablestore.Contract.SetClientState(&_Provablestore.TransactOpts, clientId, data)
}

// SetConnection is a paid mutator transaction binding the contract method 0x5e483f60.
//
// Solidity: function setConnection(string connectionId, (string,(string,string[])[],uint8,uint64,(string,string,(bytes))) connection) returns()
func (_Provablestore *ProvablestoreTransactor) SetConnection(opts *bind.TransactOpts, connectionId string, connection ConnectionEndData) (*types.Transaction, error) {
	return _Provablestore.contract.Transact(opts, "setConnection", connectionId, connection)
}

// SetConnection is a paid mutator transaction binding the contract method 0x5e483f60.
//
// Solidity: function setConnection(string connectionId, (string,(string,string[])[],uint8,uint64,(string,string,(bytes))) connection) returns()
func (_Provablestore *ProvablestoreSession) SetConnection(connectionId string, connection ConnectionEndData) (*types.Transaction, error) {
	return _Provablestore.Contract.SetConnection(&_Provablestore.TransactOpts, connectionId, connection)
}

// SetConnection is a paid mutator transaction binding the contract method 0x5e483f60.
//
// Solidity: function setConnection(string connectionId, (string,(string,string[])[],uint8,uint64,(string,string,(bytes))) connection) returns()
func (_Provablestore *ProvablestoreTransactorSession) SetConnection(connectionId string, connection ConnectionEndData) (*types.Transaction, error) {
	return _Provablestore.Contract.SetConnection(&_Provablestore.TransactOpts, connectionId, connection)
}

// SetConsensusState is a paid mutator transaction binding the contract method 0xbcce780d.
//
// Solidity: function setConsensusState(string clientId, uint64 height, (uint64,bytes,bytes[]) consensusState) returns()
func (_Provablestore *ProvablestoreTransactor) SetConsensusState(opts *bind.TransactOpts, clientId string, height uint64, consensusState ConsensusStateData) (*types.Transaction, error) {
	return _Provablestore.contract.Transact(opts, "setConsensusState", clientId, height, consensusState)
}

// SetConsensusState is a paid mutator transaction binding the contract method 0xbcce780d.
//
// Solidity: function setConsensusState(string clientId, uint64 height, (uint64,bytes,bytes[]) consensusState) returns()
func (_Provablestore *ProvablestoreSession) SetConsensusState(clientId string, height uint64, consensusState ConsensusStateData) (*types.Transaction, error) {
	return _Provablestore.Contract.SetConsensusState(&_Provablestore.TransactOpts, clientId, height, consensusState)
}

// SetConsensusState is a paid mutator transaction binding the contract method 0xbcce780d.
//
// Solidity: function setConsensusState(string clientId, uint64 height, (uint64,bytes,bytes[]) consensusState) returns()
func (_Provablestore *ProvablestoreTransactorSession) SetConsensusState(clientId string, height uint64, consensusState ConsensusStateData) (*types.Transaction, error) {
	return _Provablestore.Contract.SetConsensusState(&_Provablestore.TransactOpts, clientId, height, consensusState)
}

// SetNextSequenceAck is a paid mutator transaction binding the contract method 0x24fe5ee8.
//
// Solidity: function setNextSequenceAck(string portId, string channelId, uint64 sequence) returns()
func (_Provablestore *ProvablestoreTransactor) SetNextSequenceAck(opts *bind.TransactOpts, portId string, channelId string, sequence uint64) (*types.Transaction, error) {
	return _Provablestore.contract.Transact(opts, "setNextSequenceAck", portId, channelId, sequence)
}

// SetNextSequenceAck is a paid mutator transaction binding the contract method 0x24fe5ee8.
//
// Solidity: function setNextSequenceAck(string portId, string channelId, uint64 sequence) returns()
func (_Provablestore *ProvablestoreSession) SetNextSequenceAck(portId string, channelId string, sequence uint64) (*types.Transaction, error) {
	return _Provablestore.Contract.SetNextSequenceAck(&_Provablestore.TransactOpts, portId, channelId, sequence)
}

// SetNextSequenceAck is a paid mutator transaction binding the contract method 0x24fe5ee8.
//
// Solidity: function setNextSequenceAck(string portId, string channelId, uint64 sequence) returns()
func (_Provablestore *ProvablestoreTransactorSession) SetNextSequenceAck(portId string, channelId string, sequence uint64) (*types.Transaction, error) {
	return _Provablestore.Contract.SetNextSequenceAck(&_Provablestore.TransactOpts, portId, channelId, sequence)
}

// SetNextSequenceRecv is a paid mutator transaction binding the contract method 0xff171745.
//
// Solidity: function setNextSequenceRecv(string portId, string channelId, uint64 sequence) returns()
func (_Provablestore *ProvablestoreTransactor) SetNextSequenceRecv(opts *bind.TransactOpts, portId string, channelId string, sequence uint64) (*types.Transaction, error) {
	return _Provablestore.contract.Transact(opts, "setNextSequenceRecv", portId, channelId, sequence)
}

// SetNextSequenceRecv is a paid mutator transaction binding the contract method 0xff171745.
//
// Solidity: function setNextSequenceRecv(string portId, string channelId, uint64 sequence) returns()
func (_Provablestore *ProvablestoreSession) SetNextSequenceRecv(portId string, channelId string, sequence uint64) (*types.Transaction, error) {
	return _Provablestore.Contract.SetNextSequenceRecv(&_Provablestore.TransactOpts, portId, channelId, sequence)
}

// SetNextSequenceRecv is a paid mutator transaction binding the contract method 0xff171745.
//
// Solidity: function setNextSequenceRecv(string portId, string channelId, uint64 sequence) returns()
func (_Provablestore *ProvablestoreTransactorSession) SetNextSequenceRecv(portId string, channelId string, sequence uint64) (*types.Transaction, error) {
	return _Provablestore.Contract.SetNextSequenceRecv(&_Provablestore.TransactOpts, portId, channelId, sequence)
}

// SetNextSequenceSend is a paid mutator transaction binding the contract method 0xddc090dd.
//
// Solidity: function setNextSequenceSend(string portId, string channelId, uint64 sequence) returns()
func (_Provablestore *ProvablestoreTransactor) SetNextSequenceSend(opts *bind.TransactOpts, portId string, channelId string, sequence uint64) (*types.Transaction, error) {
	return _Provablestore.contract.Transact(opts, "setNextSequenceSend", portId, channelId, sequence)
}

// SetNextSequenceSend is a paid mutator transaction binding the contract method 0xddc090dd.
//
// Solidity: function setNextSequenceSend(string portId, string channelId, uint64 sequence) returns()
func (_Provablestore *ProvablestoreSession) SetNextSequenceSend(portId string, channelId string, sequence uint64) (*types.Transaction, error) {
	return _Provablestore.Contract.SetNextSequenceSend(&_Provablestore.TransactOpts, portId, channelId, sequence)
}

// SetNextSequenceSend is a paid mutator transaction binding the contract method 0xddc090dd.
//
// Solidity: function setNextSequenceSend(string portId, string channelId, uint64 sequence) returns()
func (_Provablestore *ProvablestoreTransactorSession) SetNextSequenceSend(portId string, channelId string, sequence uint64) (*types.Transaction, error) {
	return _Provablestore.Contract.SetNextSequenceSend(&_Provablestore.TransactOpts, portId, channelId, sequence)
}
