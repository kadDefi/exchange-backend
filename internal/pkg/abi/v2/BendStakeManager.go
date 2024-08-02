// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package v2

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
)

// IApeCoinStakingPairNft is an auto generated low-level Go binding around an user-defined struct.
type IApeCoinStakingPairNft struct {
	MainTokenId *big.Int
	BakcTokenId *big.Int
}

// IStakeManagerCompoundArgs is an auto generated low-level Go binding around an user-defined struct.
type IStakeManagerCompoundArgs struct {
	ClaimCoinPool      bool
	Claim              IStakeManagerNftArgs
	Unstake            IStakeManagerNftArgs
	Stake              IStakeManagerNftArgs
	CoinStakeThreshold *big.Int
}

// IStakeManagerNftArgs is an auto generated low-level Go binding around an user-defined struct.
type IStakeManagerNftArgs struct {
	Bayc      []*big.Int
	Mayc      []*big.Int
	BaycPairs []IApeCoinStakingPairNft
	MaycPairs []IApeCoinStakingPairNft
}

// BendStakeManagerMetaData contains all meta data concerning the BendStakeManager contract.
var BendStakeManagerMetaData = &bind.MetaData{
	ABI: "[{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint8\",\"name\":\"version\",\"type\":\"uint8\"}],\"name\":\"Initialized\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"previousOwner\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"newOwner\",\"type\":\"address\"}],\"name\":\"OwnershipTransferred\",\"type\":\"event\"},{\"inputs\":[],\"name\":\"MAX_FEE\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"MAX_PENDING_FEE\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"PERCENTAGE_FACTOR\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"botAdmin\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"rewardsAmount_\",\"type\":\"uint256\"}],\"name\":\"calculateFee\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"feeAmount\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"claimApeCoin\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"components\":[{\"internalType\":\"uint128\",\"name\":\"mainTokenId\",\"type\":\"uint128\"},{\"internalType\":\"uint128\",\"name\":\"bakcTokenId\",\"type\":\"uint128\"}],\"internalType\":\"structIApeCoinStaking.PairNft[]\",\"name\":\"baycPairs_\",\"type\":\"tuple[]\"},{\"components\":[{\"internalType\":\"uint128\",\"name\":\"mainTokenId\",\"type\":\"uint128\"},{\"internalType\":\"uint128\",\"name\":\"bakcTokenId\",\"type\":\"uint128\"}],\"internalType\":\"structIApeCoinStaking.PairNft[]\",\"name\":\"maycPairs_\",\"type\":\"tuple[]\"}],\"name\":\"claimBakc\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256[]\",\"name\":\"tokenIds_\",\"type\":\"uint256[]\"}],\"name\":\"claimBayc\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256[]\",\"name\":\"tokenIds_\",\"type\":\"uint256[]\"}],\"name\":\"claimMayc\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"components\":[{\"internalType\":\"bool\",\"name\":\"claimCoinPool\",\"type\":\"bool\"},{\"components\":[{\"internalType\":\"uint256[]\",\"name\":\"bayc\",\"type\":\"uint256[]\"},{\"internalType\":\"uint256[]\",\"name\":\"mayc\",\"type\":\"uint256[]\"},{\"components\":[{\"internalType\":\"uint128\",\"name\":\"mainTokenId\",\"type\":\"uint128\"},{\"internalType\":\"uint128\",\"name\":\"bakcTokenId\",\"type\":\"uint128\"}],\"internalType\":\"structIApeCoinStaking.PairNft[]\",\"name\":\"baycPairs\",\"type\":\"tuple[]\"},{\"components\":[{\"internalType\":\"uint128\",\"name\":\"mainTokenId\",\"type\":\"uint128\"},{\"internalType\":\"uint128\",\"name\":\"bakcTokenId\",\"type\":\"uint128\"}],\"internalType\":\"structIApeCoinStaking.PairNft[]\",\"name\":\"maycPairs\",\"type\":\"tuple[]\"}],\"internalType\":\"structIStakeManager.NftArgs\",\"name\":\"claim\",\"type\":\"tuple\"},{\"components\":[{\"internalType\":\"uint256[]\",\"name\":\"bayc\",\"type\":\"uint256[]\"},{\"internalType\":\"uint256[]\",\"name\":\"mayc\",\"type\":\"uint256[]\"},{\"components\":[{\"internalType\":\"uint128\",\"name\":\"mainTokenId\",\"type\":\"uint128\"},{\"internalType\":\"uint128\",\"name\":\"bakcTokenId\",\"type\":\"uint128\"}],\"internalType\":\"structIApeCoinStaking.PairNft[]\",\"name\":\"baycPairs\",\"type\":\"tuple[]\"},{\"components\":[{\"internalType\":\"uint128\",\"name\":\"mainTokenId\",\"type\":\"uint128\"},{\"internalType\":\"uint128\",\"name\":\"bakcTokenId\",\"type\":\"uint128\"}],\"internalType\":\"structIApeCoinStaking.PairNft[]\",\"name\":\"maycPairs\",\"type\":\"tuple[]\"}],\"internalType\":\"structIStakeManager.NftArgs\",\"name\":\"unstake\",\"type\":\"tuple\"},{\"components\":[{\"internalType\":\"uint256[]\",\"name\":\"bayc\",\"type\":\"uint256[]\"},{\"internalType\":\"uint256[]\",\"name\":\"mayc\",\"type\":\"uint256[]\"},{\"components\":[{\"internalType\":\"uint128\",\"name\":\"mainTokenId\",\"type\":\"uint128\"},{\"internalType\":\"uint128\",\"name\":\"bakcTokenId\",\"type\":\"uint128\"}],\"internalType\":\"structIApeCoinStaking.PairNft[]\",\"name\":\"baycPairs\",\"type\":\"tuple[]\"},{\"components\":[{\"internalType\":\"uint128\",\"name\":\"mainTokenId\",\"type\":\"uint128\"},{\"internalType\":\"uint128\",\"name\":\"bakcTokenId\",\"type\":\"uint128\"}],\"internalType\":\"structIApeCoinStaking.PairNft[]\",\"name\":\"maycPairs\",\"type\":\"tuple[]\"}],\"internalType\":\"structIStakeManager.NftArgs\",\"name\":\"stake\",\"type\":\"tuple\"},{\"internalType\":\"uint256\",\"name\":\"coinStakeThreshold\",\"type\":\"uint256\"}],\"internalType\":\"structIStakeManager.CompoundArgs\",\"name\":\"args_\",\"type\":\"tuple\"}],\"name\":\"compound\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"fee\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"feeRecipient\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"nft_\",\"type\":\"address\"}],\"name\":\"getNftRewardsShare\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"nftShare\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"contractIApeCoinStaking\",\"name\":\"apeStaking_\",\"type\":\"address\"},{\"internalType\":\"contractICoinPool\",\"name\":\"coinPool_\",\"type\":\"address\"},{\"internalType\":\"contractINftPool\",\"name\":\"nftPool_\",\"type\":\"address\"},{\"internalType\":\"contractINftVault\",\"name\":\"nftVault_\",\"type\":\"address\"},{\"internalType\":\"contractIStakedNft\",\"name\":\"stBayc_\",\"type\":\"address\"},{\"internalType\":\"contractIStakedNft\",\"name\":\"stMayc_\",\"type\":\"address\"},{\"internalType\":\"contractIStakedNft\",\"name\":\"stBakc_\",\"type\":\"address\"}],\"name\":\"initialize\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"contractIStakedNft\",\"name\":\"stNft_\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"to_\",\"type\":\"address\"},{\"internalType\":\"uint256[]\",\"name\":\"tokenIds_\",\"type\":\"uint256[]\"}],\"name\":\"mintStNft\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"},{\"internalType\":\"bytes\",\"name\":\"\",\"type\":\"bytes\"}],\"name\":\"onERC721Received\",\"outputs\":[{\"internalType\":\"bytes4\",\"name\":\"\",\"type\":\"bytes4\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"owner\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"pendingFeeAmount\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"poolId_\",\"type\":\"uint256\"}],\"name\":\"pendingRewards\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"nft_\",\"type\":\"address\"}],\"name\":\"refundOf\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"renounceOwnership\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"nft_\",\"type\":\"address\"}],\"name\":\"rewardsStrategies\",\"outputs\":[{\"internalType\":\"contractIRewardsStrategy\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"stBakc\",\"outputs\":[{\"internalType\":\"contractIStakedNft\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"stBayc\",\"outputs\":[{\"internalType\":\"contractIStakedNft\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"stMayc\",\"outputs\":[{\"internalType\":\"contractIStakedNft\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"amount_\",\"type\":\"uint256\"}],\"name\":\"stakeApeCoin\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"components\":[{\"internalType\":\"uint128\",\"name\":\"mainTokenId\",\"type\":\"uint128\"},{\"internalType\":\"uint128\",\"name\":\"bakcTokenId\",\"type\":\"uint128\"}],\"internalType\":\"structIApeCoinStaking.PairNft[]\",\"name\":\"baycPairs_\",\"type\":\"tuple[]\"},{\"components\":[{\"internalType\":\"uint128\",\"name\":\"mainTokenId\",\"type\":\"uint128\"},{\"internalType\":\"uint128\",\"name\":\"bakcTokenId\",\"type\":\"uint128\"}],\"internalType\":\"structIApeCoinStaking.PairNft[]\",\"name\":\"maycPairs_\",\"type\":\"tuple[]\"}],\"name\":\"stakeBakc\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256[]\",\"name\":\"tokenIds_\",\"type\":\"uint256[]\"}],\"name\":\"stakeBayc\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256[]\",\"name\":\"tokenIds_\",\"type\":\"uint256[]\"}],\"name\":\"stakeMayc\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"poolId_\",\"type\":\"uint256\"}],\"name\":\"stakedApeCoin\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"totalPendingRewards\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"totalRefund\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"totalStakedApeCoin\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"newOwner\",\"type\":\"address\"}],\"name\":\"transferOwnership\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"amount_\",\"type\":\"uint256\"}],\"name\":\"unstakeApeCoin\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"components\":[{\"internalType\":\"uint128\",\"name\":\"mainTokenId\",\"type\":\"uint128\"},{\"internalType\":\"uint128\",\"name\":\"bakcTokenId\",\"type\":\"uint128\"}],\"internalType\":\"structIApeCoinStaking.PairNft[]\",\"name\":\"baycPairs_\",\"type\":\"tuple[]\"},{\"components\":[{\"internalType\":\"uint128\",\"name\":\"mainTokenId\",\"type\":\"uint128\"},{\"internalType\":\"uint128\",\"name\":\"bakcTokenId\",\"type\":\"uint128\"}],\"internalType\":\"structIApeCoinStaking.PairNft[]\",\"name\":\"maycPairs_\",\"type\":\"tuple[]\"}],\"name\":\"unstakeBakc\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256[]\",\"name\":\"tokenIds_\",\"type\":\"uint256[]\"}],\"name\":\"unstakeBayc\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256[]\",\"name\":\"tokenIds_\",\"type\":\"uint256[]\"}],\"name\":\"unstakeMayc\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"botAdmin_\",\"type\":\"address\"}],\"name\":\"updateBotAdmin\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"fee_\",\"type\":\"uint256\"}],\"name\":\"updateFee\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"recipient_\",\"type\":\"address\"}],\"name\":\"updateFeeRecipient\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"nft_\",\"type\":\"address\"},{\"internalType\":\"contractIRewardsStrategy\",\"name\":\"rewardsStrategy_\",\"type\":\"address\"}],\"name\":\"updateRewardsStrategy\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"contractIWithdrawStrategy\",\"name\":\"withdrawStrategy_\",\"type\":\"address\"}],\"name\":\"updateWithdrawStrategy\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"required\",\"type\":\"uint256\"}],\"name\":\"withdrawApeCoin\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"withdrawn\",\"type\":\"uint256\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"nft_\",\"type\":\"address\"}],\"name\":\"withdrawRefund\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"withdrawTotalRefund\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"}]",
}

// BendStakeManagerABI is the input ABI used to generate the binding from.
// Deprecated: Use BendStakeManagerMetaData.ABI instead.
var BendStakeManagerABI = BendStakeManagerMetaData.ABI

// BendStakeManager is an auto generated Go binding around an Ethereum contract.
type BendStakeManager struct {
	BendStakeManagerCaller     // Read-only binding to the contract
	BendStakeManagerTransactor // Write-only binding to the contract
	BendStakeManagerFilterer   // Log filterer for contract events
}

// BendStakeManagerCaller is an auto generated read-only Go binding around an Ethereum contract.
type BendStakeManagerCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// BendStakeManagerTransactor is an auto generated write-only Go binding around an Ethereum contract.
type BendStakeManagerTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// BendStakeManagerFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type BendStakeManagerFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// BendStakeManagerSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type BendStakeManagerSession struct {
	Contract     *BendStakeManager // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// BendStakeManagerCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type BendStakeManagerCallerSession struct {
	Contract *BendStakeManagerCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts           // Call options to use throughout this session
}

// BendStakeManagerTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type BendStakeManagerTransactorSession struct {
	Contract     *BendStakeManagerTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts           // Transaction auth options to use throughout this session
}

// BendStakeManagerRaw is an auto generated low-level Go binding around an Ethereum contract.
type BendStakeManagerRaw struct {
	Contract *BendStakeManager // Generic contract binding to access the raw methods on
}

// BendStakeManagerCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type BendStakeManagerCallerRaw struct {
	Contract *BendStakeManagerCaller // Generic read-only contract binding to access the raw methods on
}

// BendStakeManagerTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type BendStakeManagerTransactorRaw struct {
	Contract *BendStakeManagerTransactor // Generic write-only contract binding to access the raw methods on
}

// NewBendStakeManager creates a new instance of BendStakeManager, bound to a specific deployed contract.
func NewBendStakeManager(address common.Address, backend bind.ContractBackend) (*BendStakeManager, error) {
	contract, err := bindBendStakeManager(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &BendStakeManager{BendStakeManagerCaller: BendStakeManagerCaller{contract: contract}, BendStakeManagerTransactor: BendStakeManagerTransactor{contract: contract}, BendStakeManagerFilterer: BendStakeManagerFilterer{contract: contract}}, nil
}

// NewBendStakeManagerCaller creates a new read-only instance of BendStakeManager, bound to a specific deployed contract.
func NewBendStakeManagerCaller(address common.Address, caller bind.ContractCaller) (*BendStakeManagerCaller, error) {
	contract, err := bindBendStakeManager(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &BendStakeManagerCaller{contract: contract}, nil
}

// NewBendStakeManagerTransactor creates a new write-only instance of BendStakeManager, bound to a specific deployed contract.
func NewBendStakeManagerTransactor(address common.Address, transactor bind.ContractTransactor) (*BendStakeManagerTransactor, error) {
	contract, err := bindBendStakeManager(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &BendStakeManagerTransactor{contract: contract}, nil
}

// NewBendStakeManagerFilterer creates a new log filterer instance of BendStakeManager, bound to a specific deployed contract.
func NewBendStakeManagerFilterer(address common.Address, filterer bind.ContractFilterer) (*BendStakeManagerFilterer, error) {
	contract, err := bindBendStakeManager(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &BendStakeManagerFilterer{contract: contract}, nil
}

// bindBendStakeManager binds a generic wrapper to an already deployed contract.
func bindBendStakeManager(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(BendStakeManagerABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_BendStakeManager *BendStakeManagerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _BendStakeManager.Contract.BendStakeManagerCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_BendStakeManager *BendStakeManagerRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _BendStakeManager.Contract.BendStakeManagerTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_BendStakeManager *BendStakeManagerRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _BendStakeManager.Contract.BendStakeManagerTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_BendStakeManager *BendStakeManagerCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _BendStakeManager.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_BendStakeManager *BendStakeManagerTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _BendStakeManager.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_BendStakeManager *BendStakeManagerTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _BendStakeManager.Contract.contract.Transact(opts, method, params...)
}

// MAXFEE is a free data retrieval call binding the contract method 0xbc063e1a.
//
// Solidity: function MAX_FEE() view returns(uint256)
func (_BendStakeManager *BendStakeManagerCaller) MAXFEE(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _BendStakeManager.contract.Call(opts, &out, "MAX_FEE")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// MAXFEE is a free data retrieval call binding the contract method 0xbc063e1a.
//
// Solidity: function MAX_FEE() view returns(uint256)
func (_BendStakeManager *BendStakeManagerSession) MAXFEE() (*big.Int, error) {
	return _BendStakeManager.Contract.MAXFEE(&_BendStakeManager.CallOpts)
}

// MAXFEE is a free data retrieval call binding the contract method 0xbc063e1a.
//
// Solidity: function MAX_FEE() view returns(uint256)
func (_BendStakeManager *BendStakeManagerCallerSession) MAXFEE() (*big.Int, error) {
	return _BendStakeManager.Contract.MAXFEE(&_BendStakeManager.CallOpts)
}

// MAXPENDINGFEE is a free data retrieval call binding the contract method 0x9d3f9531.
//
// Solidity: function MAX_PENDING_FEE() view returns(uint256)
func (_BendStakeManager *BendStakeManagerCaller) MAXPENDINGFEE(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _BendStakeManager.contract.Call(opts, &out, "MAX_PENDING_FEE")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// MAXPENDINGFEE is a free data retrieval call binding the contract method 0x9d3f9531.
//
// Solidity: function MAX_PENDING_FEE() view returns(uint256)
func (_BendStakeManager *BendStakeManagerSession) MAXPENDINGFEE() (*big.Int, error) {
	return _BendStakeManager.Contract.MAXPENDINGFEE(&_BendStakeManager.CallOpts)
}

// MAXPENDINGFEE is a free data retrieval call binding the contract method 0x9d3f9531.
//
// Solidity: function MAX_PENDING_FEE() view returns(uint256)
func (_BendStakeManager *BendStakeManagerCallerSession) MAXPENDINGFEE() (*big.Int, error) {
	return _BendStakeManager.Contract.MAXPENDINGFEE(&_BendStakeManager.CallOpts)
}

// PERCENTAGEFACTOR is a free data retrieval call binding the contract method 0xee01e5e7.
//
// Solidity: function PERCENTAGE_FACTOR() view returns(uint256)
func (_BendStakeManager *BendStakeManagerCaller) PERCENTAGEFACTOR(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _BendStakeManager.contract.Call(opts, &out, "PERCENTAGE_FACTOR")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// PERCENTAGEFACTOR is a free data retrieval call binding the contract method 0xee01e5e7.
//
// Solidity: function PERCENTAGE_FACTOR() view returns(uint256)
func (_BendStakeManager *BendStakeManagerSession) PERCENTAGEFACTOR() (*big.Int, error) {
	return _BendStakeManager.Contract.PERCENTAGEFACTOR(&_BendStakeManager.CallOpts)
}

// PERCENTAGEFACTOR is a free data retrieval call binding the contract method 0xee01e5e7.
//
// Solidity: function PERCENTAGE_FACTOR() view returns(uint256)
func (_BendStakeManager *BendStakeManagerCallerSession) PERCENTAGEFACTOR() (*big.Int, error) {
	return _BendStakeManager.Contract.PERCENTAGEFACTOR(&_BendStakeManager.CallOpts)
}

// BotAdmin is a free data retrieval call binding the contract method 0xf3c717b8.
//
// Solidity: function botAdmin() view returns(address)
func (_BendStakeManager *BendStakeManagerCaller) BotAdmin(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _BendStakeManager.contract.Call(opts, &out, "botAdmin")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// BotAdmin is a free data retrieval call binding the contract method 0xf3c717b8.
//
// Solidity: function botAdmin() view returns(address)
func (_BendStakeManager *BendStakeManagerSession) BotAdmin() (common.Address, error) {
	return _BendStakeManager.Contract.BotAdmin(&_BendStakeManager.CallOpts)
}

// BotAdmin is a free data retrieval call binding the contract method 0xf3c717b8.
//
// Solidity: function botAdmin() view returns(address)
func (_BendStakeManager *BendStakeManagerCallerSession) BotAdmin() (common.Address, error) {
	return _BendStakeManager.Contract.BotAdmin(&_BendStakeManager.CallOpts)
}

// CalculateFee is a free data retrieval call binding the contract method 0x99a5d747.
//
// Solidity: function calculateFee(uint256 rewardsAmount_) view returns(uint256 feeAmount)
func (_BendStakeManager *BendStakeManagerCaller) CalculateFee(opts *bind.CallOpts, rewardsAmount_ *big.Int) (*big.Int, error) {
	var out []interface{}
	err := _BendStakeManager.contract.Call(opts, &out, "calculateFee", rewardsAmount_)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// CalculateFee is a free data retrieval call binding the contract method 0x99a5d747.
//
// Solidity: function calculateFee(uint256 rewardsAmount_) view returns(uint256 feeAmount)
func (_BendStakeManager *BendStakeManagerSession) CalculateFee(rewardsAmount_ *big.Int) (*big.Int, error) {
	return _BendStakeManager.Contract.CalculateFee(&_BendStakeManager.CallOpts, rewardsAmount_)
}

// CalculateFee is a free data retrieval call binding the contract method 0x99a5d747.
//
// Solidity: function calculateFee(uint256 rewardsAmount_) view returns(uint256 feeAmount)
func (_BendStakeManager *BendStakeManagerCallerSession) CalculateFee(rewardsAmount_ *big.Int) (*big.Int, error) {
	return _BendStakeManager.Contract.CalculateFee(&_BendStakeManager.CallOpts, rewardsAmount_)
}

// Fee is a free data retrieval call binding the contract method 0xddca3f43.
//
// Solidity: function fee() view returns(uint256)
func (_BendStakeManager *BendStakeManagerCaller) Fee(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _BendStakeManager.contract.Call(opts, &out, "fee")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// Fee is a free data retrieval call binding the contract method 0xddca3f43.
//
// Solidity: function fee() view returns(uint256)
func (_BendStakeManager *BendStakeManagerSession) Fee() (*big.Int, error) {
	return _BendStakeManager.Contract.Fee(&_BendStakeManager.CallOpts)
}

// Fee is a free data retrieval call binding the contract method 0xddca3f43.
//
// Solidity: function fee() view returns(uint256)
func (_BendStakeManager *BendStakeManagerCallerSession) Fee() (*big.Int, error) {
	return _BendStakeManager.Contract.Fee(&_BendStakeManager.CallOpts)
}

// FeeRecipient is a free data retrieval call binding the contract method 0x46904840.
//
// Solidity: function feeRecipient() view returns(address)
func (_BendStakeManager *BendStakeManagerCaller) FeeRecipient(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _BendStakeManager.contract.Call(opts, &out, "feeRecipient")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// FeeRecipient is a free data retrieval call binding the contract method 0x46904840.
//
// Solidity: function feeRecipient() view returns(address)
func (_BendStakeManager *BendStakeManagerSession) FeeRecipient() (common.Address, error) {
	return _BendStakeManager.Contract.FeeRecipient(&_BendStakeManager.CallOpts)
}

// FeeRecipient is a free data retrieval call binding the contract method 0x46904840.
//
// Solidity: function feeRecipient() view returns(address)
func (_BendStakeManager *BendStakeManagerCallerSession) FeeRecipient() (common.Address, error) {
	return _BendStakeManager.Contract.FeeRecipient(&_BendStakeManager.CallOpts)
}

// GetNftRewardsShare is a free data retrieval call binding the contract method 0x4ae1596a.
//
// Solidity: function getNftRewardsShare(address nft_) view returns(uint256 nftShare)
func (_BendStakeManager *BendStakeManagerCaller) GetNftRewardsShare(opts *bind.CallOpts, nft_ common.Address) (*big.Int, error) {
	var out []interface{}
	err := _BendStakeManager.contract.Call(opts, &out, "getNftRewardsShare", nft_)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// GetNftRewardsShare is a free data retrieval call binding the contract method 0x4ae1596a.
//
// Solidity: function getNftRewardsShare(address nft_) view returns(uint256 nftShare)
func (_BendStakeManager *BendStakeManagerSession) GetNftRewardsShare(nft_ common.Address) (*big.Int, error) {
	return _BendStakeManager.Contract.GetNftRewardsShare(&_BendStakeManager.CallOpts, nft_)
}

// GetNftRewardsShare is a free data retrieval call binding the contract method 0x4ae1596a.
//
// Solidity: function getNftRewardsShare(address nft_) view returns(uint256 nftShare)
func (_BendStakeManager *BendStakeManagerCallerSession) GetNftRewardsShare(nft_ common.Address) (*big.Int, error) {
	return _BendStakeManager.Contract.GetNftRewardsShare(&_BendStakeManager.CallOpts, nft_)
}

// OnERC721Received is a free data retrieval call binding the contract method 0x150b7a02.
//
// Solidity: function onERC721Received(address , address , uint256 , bytes ) view returns(bytes4)
func (_BendStakeManager *BendStakeManagerCaller) OnERC721Received(opts *bind.CallOpts, arg0 common.Address, arg1 common.Address, arg2 *big.Int, arg3 []byte) ([4]byte, error) {
	var out []interface{}
	err := _BendStakeManager.contract.Call(opts, &out, "onERC721Received", arg0, arg1, arg2, arg3)

	if err != nil {
		return *new([4]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([4]byte)).(*[4]byte)

	return out0, err

}

// OnERC721Received is a free data retrieval call binding the contract method 0x150b7a02.
//
// Solidity: function onERC721Received(address , address , uint256 , bytes ) view returns(bytes4)
func (_BendStakeManager *BendStakeManagerSession) OnERC721Received(arg0 common.Address, arg1 common.Address, arg2 *big.Int, arg3 []byte) ([4]byte, error) {
	return _BendStakeManager.Contract.OnERC721Received(&_BendStakeManager.CallOpts, arg0, arg1, arg2, arg3)
}

// OnERC721Received is a free data retrieval call binding the contract method 0x150b7a02.
//
// Solidity: function onERC721Received(address , address , uint256 , bytes ) view returns(bytes4)
func (_BendStakeManager *BendStakeManagerCallerSession) OnERC721Received(arg0 common.Address, arg1 common.Address, arg2 *big.Int, arg3 []byte) ([4]byte, error) {
	return _BendStakeManager.Contract.OnERC721Received(&_BendStakeManager.CallOpts, arg0, arg1, arg2, arg3)
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_BendStakeManager *BendStakeManagerCaller) Owner(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _BendStakeManager.contract.Call(opts, &out, "owner")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_BendStakeManager *BendStakeManagerSession) Owner() (common.Address, error) {
	return _BendStakeManager.Contract.Owner(&_BendStakeManager.CallOpts)
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_BendStakeManager *BendStakeManagerCallerSession) Owner() (common.Address, error) {
	return _BendStakeManager.Contract.Owner(&_BendStakeManager.CallOpts)
}

// PendingFeeAmount is a free data retrieval call binding the contract method 0xc0ee83b9.
//
// Solidity: function pendingFeeAmount() view returns(uint256)
func (_BendStakeManager *BendStakeManagerCaller) PendingFeeAmount(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _BendStakeManager.contract.Call(opts, &out, "pendingFeeAmount")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// PendingFeeAmount is a free data retrieval call binding the contract method 0xc0ee83b9.
//
// Solidity: function pendingFeeAmount() view returns(uint256)
func (_BendStakeManager *BendStakeManagerSession) PendingFeeAmount() (*big.Int, error) {
	return _BendStakeManager.Contract.PendingFeeAmount(&_BendStakeManager.CallOpts)
}

// PendingFeeAmount is a free data retrieval call binding the contract method 0xc0ee83b9.
//
// Solidity: function pendingFeeAmount() view returns(uint256)
func (_BendStakeManager *BendStakeManagerCallerSession) PendingFeeAmount() (*big.Int, error) {
	return _BendStakeManager.Contract.PendingFeeAmount(&_BendStakeManager.CallOpts)
}

// PendingRewards is a free data retrieval call binding the contract method 0x7dcb2abf.
//
// Solidity: function pendingRewards(uint256 poolId_) view returns(uint256 amount)
func (_BendStakeManager *BendStakeManagerCaller) PendingRewards(opts *bind.CallOpts, poolId_ *big.Int) (*big.Int, error) {
	var out []interface{}
	err := _BendStakeManager.contract.Call(opts, &out, "pendingRewards", poolId_)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// PendingRewards is a free data retrieval call binding the contract method 0x7dcb2abf.
//
// Solidity: function pendingRewards(uint256 poolId_) view returns(uint256 amount)
func (_BendStakeManager *BendStakeManagerSession) PendingRewards(poolId_ *big.Int) (*big.Int, error) {
	return _BendStakeManager.Contract.PendingRewards(&_BendStakeManager.CallOpts, poolId_)
}

// PendingRewards is a free data retrieval call binding the contract method 0x7dcb2abf.
//
// Solidity: function pendingRewards(uint256 poolId_) view returns(uint256 amount)
func (_BendStakeManager *BendStakeManagerCallerSession) PendingRewards(poolId_ *big.Int) (*big.Int, error) {
	return _BendStakeManager.Contract.PendingRewards(&_BendStakeManager.CallOpts, poolId_)
}

// RefundOf is a free data retrieval call binding the contract method 0x07738ad2.
//
// Solidity: function refundOf(address nft_) view returns(uint256 amount)
func (_BendStakeManager *BendStakeManagerCaller) RefundOf(opts *bind.CallOpts, nft_ common.Address) (*big.Int, error) {
	var out []interface{}
	err := _BendStakeManager.contract.Call(opts, &out, "refundOf", nft_)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// RefundOf is a free data retrieval call binding the contract method 0x07738ad2.
//
// Solidity: function refundOf(address nft_) view returns(uint256 amount)
func (_BendStakeManager *BendStakeManagerSession) RefundOf(nft_ common.Address) (*big.Int, error) {
	return _BendStakeManager.Contract.RefundOf(&_BendStakeManager.CallOpts, nft_)
}

// RefundOf is a free data retrieval call binding the contract method 0x07738ad2.
//
// Solidity: function refundOf(address nft_) view returns(uint256 amount)
func (_BendStakeManager *BendStakeManagerCallerSession) RefundOf(nft_ common.Address) (*big.Int, error) {
	return _BendStakeManager.Contract.RefundOf(&_BendStakeManager.CallOpts, nft_)
}

// RewardsStrategies is a free data retrieval call binding the contract method 0x46d0d554.
//
// Solidity: function rewardsStrategies(address nft_) view returns(address)
func (_BendStakeManager *BendStakeManagerCaller) RewardsStrategies(opts *bind.CallOpts, nft_ common.Address) (common.Address, error) {
	var out []interface{}
	err := _BendStakeManager.contract.Call(opts, &out, "rewardsStrategies", nft_)

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// RewardsStrategies is a free data retrieval call binding the contract method 0x46d0d554.
//
// Solidity: function rewardsStrategies(address nft_) view returns(address)
func (_BendStakeManager *BendStakeManagerSession) RewardsStrategies(nft_ common.Address) (common.Address, error) {
	return _BendStakeManager.Contract.RewardsStrategies(&_BendStakeManager.CallOpts, nft_)
}

// RewardsStrategies is a free data retrieval call binding the contract method 0x46d0d554.
//
// Solidity: function rewardsStrategies(address nft_) view returns(address)
func (_BendStakeManager *BendStakeManagerCallerSession) RewardsStrategies(nft_ common.Address) (common.Address, error) {
	return _BendStakeManager.Contract.RewardsStrategies(&_BendStakeManager.CallOpts, nft_)
}

// StBakc is a free data retrieval call binding the contract method 0x167ef616.
//
// Solidity: function stBakc() view returns(address)
func (_BendStakeManager *BendStakeManagerCaller) StBakc(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _BendStakeManager.contract.Call(opts, &out, "stBakc")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// StBakc is a free data retrieval call binding the contract method 0x167ef616.
//
// Solidity: function stBakc() view returns(address)
func (_BendStakeManager *BendStakeManagerSession) StBakc() (common.Address, error) {
	return _BendStakeManager.Contract.StBakc(&_BendStakeManager.CallOpts)
}

// StBakc is a free data retrieval call binding the contract method 0x167ef616.
//
// Solidity: function stBakc() view returns(address)
func (_BendStakeManager *BendStakeManagerCallerSession) StBakc() (common.Address, error) {
	return _BendStakeManager.Contract.StBakc(&_BendStakeManager.CallOpts)
}

// StBayc is a free data retrieval call binding the contract method 0xee22f4ca.
//
// Solidity: function stBayc() view returns(address)
func (_BendStakeManager *BendStakeManagerCaller) StBayc(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _BendStakeManager.contract.Call(opts, &out, "stBayc")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// StBayc is a free data retrieval call binding the contract method 0xee22f4ca.
//
// Solidity: function stBayc() view returns(address)
func (_BendStakeManager *BendStakeManagerSession) StBayc() (common.Address, error) {
	return _BendStakeManager.Contract.StBayc(&_BendStakeManager.CallOpts)
}

// StBayc is a free data retrieval call binding the contract method 0xee22f4ca.
//
// Solidity: function stBayc() view returns(address)
func (_BendStakeManager *BendStakeManagerCallerSession) StBayc() (common.Address, error) {
	return _BendStakeManager.Contract.StBayc(&_BendStakeManager.CallOpts)
}

// StMayc is a free data retrieval call binding the contract method 0xff6b0f81.
//
// Solidity: function stMayc() view returns(address)
func (_BendStakeManager *BendStakeManagerCaller) StMayc(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _BendStakeManager.contract.Call(opts, &out, "stMayc")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// StMayc is a free data retrieval call binding the contract method 0xff6b0f81.
//
// Solidity: function stMayc() view returns(address)
func (_BendStakeManager *BendStakeManagerSession) StMayc() (common.Address, error) {
	return _BendStakeManager.Contract.StMayc(&_BendStakeManager.CallOpts)
}

// StMayc is a free data retrieval call binding the contract method 0xff6b0f81.
//
// Solidity: function stMayc() view returns(address)
func (_BendStakeManager *BendStakeManagerCallerSession) StMayc() (common.Address, error) {
	return _BendStakeManager.Contract.StMayc(&_BendStakeManager.CallOpts)
}

// StakedApeCoin is a free data retrieval call binding the contract method 0xcecfb447.
//
// Solidity: function stakedApeCoin(uint256 poolId_) view returns(uint256)
func (_BendStakeManager *BendStakeManagerCaller) StakedApeCoin(opts *bind.CallOpts, poolId_ *big.Int) (*big.Int, error) {
	var out []interface{}
	err := _BendStakeManager.contract.Call(opts, &out, "stakedApeCoin", poolId_)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// StakedApeCoin is a free data retrieval call binding the contract method 0xcecfb447.
//
// Solidity: function stakedApeCoin(uint256 poolId_) view returns(uint256)
func (_BendStakeManager *BendStakeManagerSession) StakedApeCoin(poolId_ *big.Int) (*big.Int, error) {
	return _BendStakeManager.Contract.StakedApeCoin(&_BendStakeManager.CallOpts, poolId_)
}

// StakedApeCoin is a free data retrieval call binding the contract method 0xcecfb447.
//
// Solidity: function stakedApeCoin(uint256 poolId_) view returns(uint256)
func (_BendStakeManager *BendStakeManagerCallerSession) StakedApeCoin(poolId_ *big.Int) (*big.Int, error) {
	return _BendStakeManager.Contract.StakedApeCoin(&_BendStakeManager.CallOpts, poolId_)
}

// TotalPendingRewards is a free data retrieval call binding the contract method 0xec1371f2.
//
// Solidity: function totalPendingRewards() view returns(uint256 amount)
func (_BendStakeManager *BendStakeManagerCaller) TotalPendingRewards(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _BendStakeManager.contract.Call(opts, &out, "totalPendingRewards")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// TotalPendingRewards is a free data retrieval call binding the contract method 0xec1371f2.
//
// Solidity: function totalPendingRewards() view returns(uint256 amount)
func (_BendStakeManager *BendStakeManagerSession) TotalPendingRewards() (*big.Int, error) {
	return _BendStakeManager.Contract.TotalPendingRewards(&_BendStakeManager.CallOpts)
}

// TotalPendingRewards is a free data retrieval call binding the contract method 0xec1371f2.
//
// Solidity: function totalPendingRewards() view returns(uint256 amount)
func (_BendStakeManager *BendStakeManagerCallerSession) TotalPendingRewards() (*big.Int, error) {
	return _BendStakeManager.Contract.TotalPendingRewards(&_BendStakeManager.CallOpts)
}

// TotalRefund is a free data retrieval call binding the contract method 0x01541fe3.
//
// Solidity: function totalRefund() view returns(uint256 amount)
func (_BendStakeManager *BendStakeManagerCaller) TotalRefund(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _BendStakeManager.contract.Call(opts, &out, "totalRefund")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// TotalRefund is a free data retrieval call binding the contract method 0x01541fe3.
//
// Solidity: function totalRefund() view returns(uint256 amount)
func (_BendStakeManager *BendStakeManagerSession) TotalRefund() (*big.Int, error) {
	return _BendStakeManager.Contract.TotalRefund(&_BendStakeManager.CallOpts)
}

// TotalRefund is a free data retrieval call binding the contract method 0x01541fe3.
//
// Solidity: function totalRefund() view returns(uint256 amount)
func (_BendStakeManager *BendStakeManagerCallerSession) TotalRefund() (*big.Int, error) {
	return _BendStakeManager.Contract.TotalRefund(&_BendStakeManager.CallOpts)
}

// TotalStakedApeCoin is a free data retrieval call binding the contract method 0x0ea1fc88.
//
// Solidity: function totalStakedApeCoin() view returns(uint256 amount)
func (_BendStakeManager *BendStakeManagerCaller) TotalStakedApeCoin(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _BendStakeManager.contract.Call(opts, &out, "totalStakedApeCoin")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// TotalStakedApeCoin is a free data retrieval call binding the contract method 0x0ea1fc88.
//
// Solidity: function totalStakedApeCoin() view returns(uint256 amount)
func (_BendStakeManager *BendStakeManagerSession) TotalStakedApeCoin() (*big.Int, error) {
	return _BendStakeManager.Contract.TotalStakedApeCoin(&_BendStakeManager.CallOpts)
}

// TotalStakedApeCoin is a free data retrieval call binding the contract method 0x0ea1fc88.
//
// Solidity: function totalStakedApeCoin() view returns(uint256 amount)
func (_BendStakeManager *BendStakeManagerCallerSession) TotalStakedApeCoin() (*big.Int, error) {
	return _BendStakeManager.Contract.TotalStakedApeCoin(&_BendStakeManager.CallOpts)
}

// ClaimApeCoin is a paid mutator transaction binding the contract method 0x9b3cd1be.
//
// Solidity: function claimApeCoin() returns()
func (_BendStakeManager *BendStakeManagerTransactor) ClaimApeCoin(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _BendStakeManager.contract.Transact(opts, "claimApeCoin")
}

// ClaimApeCoin is a paid mutator transaction binding the contract method 0x9b3cd1be.
//
// Solidity: function claimApeCoin() returns()
func (_BendStakeManager *BendStakeManagerSession) ClaimApeCoin() (*types.Transaction, error) {
	return _BendStakeManager.Contract.ClaimApeCoin(&_BendStakeManager.TransactOpts)
}

// ClaimApeCoin is a paid mutator transaction binding the contract method 0x9b3cd1be.
//
// Solidity: function claimApeCoin() returns()
func (_BendStakeManager *BendStakeManagerTransactorSession) ClaimApeCoin() (*types.Transaction, error) {
	return _BendStakeManager.Contract.ClaimApeCoin(&_BendStakeManager.TransactOpts)
}

// ClaimBakc is a paid mutator transaction binding the contract method 0xc7b8f9c8.
//
// Solidity: function claimBakc((uint128,uint128)[] baycPairs_, (uint128,uint128)[] maycPairs_) returns()
func (_BendStakeManager *BendStakeManagerTransactor) ClaimBakc(opts *bind.TransactOpts, baycPairs_ []IApeCoinStakingPairNft, maycPairs_ []IApeCoinStakingPairNft) (*types.Transaction, error) {
	return _BendStakeManager.contract.Transact(opts, "claimBakc", baycPairs_, maycPairs_)
}

// ClaimBakc is a paid mutator transaction binding the contract method 0xc7b8f9c8.
//
// Solidity: function claimBakc((uint128,uint128)[] baycPairs_, (uint128,uint128)[] maycPairs_) returns()
func (_BendStakeManager *BendStakeManagerSession) ClaimBakc(baycPairs_ []IApeCoinStakingPairNft, maycPairs_ []IApeCoinStakingPairNft) (*types.Transaction, error) {
	return _BendStakeManager.Contract.ClaimBakc(&_BendStakeManager.TransactOpts, baycPairs_, maycPairs_)
}

// ClaimBakc is a paid mutator transaction binding the contract method 0xc7b8f9c8.
//
// Solidity: function claimBakc((uint128,uint128)[] baycPairs_, (uint128,uint128)[] maycPairs_) returns()
func (_BendStakeManager *BendStakeManagerTransactorSession) ClaimBakc(baycPairs_ []IApeCoinStakingPairNft, maycPairs_ []IApeCoinStakingPairNft) (*types.Transaction, error) {
	return _BendStakeManager.Contract.ClaimBakc(&_BendStakeManager.TransactOpts, baycPairs_, maycPairs_)
}

// ClaimBayc is a paid mutator transaction binding the contract method 0x35daa911.
//
// Solidity: function claimBayc(uint256[] tokenIds_) returns()
func (_BendStakeManager *BendStakeManagerTransactor) ClaimBayc(opts *bind.TransactOpts, tokenIds_ []*big.Int) (*types.Transaction, error) {
	return _BendStakeManager.contract.Transact(opts, "claimBayc", tokenIds_)
}

// ClaimBayc is a paid mutator transaction binding the contract method 0x35daa911.
//
// Solidity: function claimBayc(uint256[] tokenIds_) returns()
func (_BendStakeManager *BendStakeManagerSession) ClaimBayc(tokenIds_ []*big.Int) (*types.Transaction, error) {
	return _BendStakeManager.Contract.ClaimBayc(&_BendStakeManager.TransactOpts, tokenIds_)
}

// ClaimBayc is a paid mutator transaction binding the contract method 0x35daa911.
//
// Solidity: function claimBayc(uint256[] tokenIds_) returns()
func (_BendStakeManager *BendStakeManagerTransactorSession) ClaimBayc(tokenIds_ []*big.Int) (*types.Transaction, error) {
	return _BendStakeManager.Contract.ClaimBayc(&_BendStakeManager.TransactOpts, tokenIds_)
}

// ClaimMayc is a paid mutator transaction binding the contract method 0xb94d43b7.
//
// Solidity: function claimMayc(uint256[] tokenIds_) returns()
func (_BendStakeManager *BendStakeManagerTransactor) ClaimMayc(opts *bind.TransactOpts, tokenIds_ []*big.Int) (*types.Transaction, error) {
	return _BendStakeManager.contract.Transact(opts, "claimMayc", tokenIds_)
}

// ClaimMayc is a paid mutator transaction binding the contract method 0xb94d43b7.
//
// Solidity: function claimMayc(uint256[] tokenIds_) returns()
func (_BendStakeManager *BendStakeManagerSession) ClaimMayc(tokenIds_ []*big.Int) (*types.Transaction, error) {
	return _BendStakeManager.Contract.ClaimMayc(&_BendStakeManager.TransactOpts, tokenIds_)
}

// ClaimMayc is a paid mutator transaction binding the contract method 0xb94d43b7.
//
// Solidity: function claimMayc(uint256[] tokenIds_) returns()
func (_BendStakeManager *BendStakeManagerTransactorSession) ClaimMayc(tokenIds_ []*big.Int) (*types.Transaction, error) {
	return _BendStakeManager.Contract.ClaimMayc(&_BendStakeManager.TransactOpts, tokenIds_)
}

// Compound is a paid mutator transaction binding the contract method 0xbeaa4540.
//
// Solidity: function compound((bool,(uint256[],uint256[],(uint128,uint128)[],(uint128,uint128)[]),(uint256[],uint256[],(uint128,uint128)[],(uint128,uint128)[]),(uint256[],uint256[],(uint128,uint128)[],(uint128,uint128)[]),uint256) args_) returns()
func (_BendStakeManager *BendStakeManagerTransactor) Compound(opts *bind.TransactOpts, args_ IStakeManagerCompoundArgs) (*types.Transaction, error) {
	return _BendStakeManager.contract.Transact(opts, "compound", args_)
}

// Compound is a paid mutator transaction binding the contract method 0xbeaa4540.
//
// Solidity: function compound((bool,(uint256[],uint256[],(uint128,uint128)[],(uint128,uint128)[]),(uint256[],uint256[],(uint128,uint128)[],(uint128,uint128)[]),(uint256[],uint256[],(uint128,uint128)[],(uint128,uint128)[]),uint256) args_) returns()
func (_BendStakeManager *BendStakeManagerSession) Compound(args_ IStakeManagerCompoundArgs) (*types.Transaction, error) {
	return _BendStakeManager.Contract.Compound(&_BendStakeManager.TransactOpts, args_)
}

// Compound is a paid mutator transaction binding the contract method 0xbeaa4540.
//
// Solidity: function compound((bool,(uint256[],uint256[],(uint128,uint128)[],(uint128,uint128)[]),(uint256[],uint256[],(uint128,uint128)[],(uint128,uint128)[]),(uint256[],uint256[],(uint128,uint128)[],(uint128,uint128)[]),uint256) args_) returns()
func (_BendStakeManager *BendStakeManagerTransactorSession) Compound(args_ IStakeManagerCompoundArgs) (*types.Transaction, error) {
	return _BendStakeManager.Contract.Compound(&_BendStakeManager.TransactOpts, args_)
}

// Initialize is a paid mutator transaction binding the contract method 0x35876476.
//
// Solidity: function initialize(address apeStaking_, address coinPool_, address nftPool_, address nftVault_, address stBayc_, address stMayc_, address stBakc_) returns()
func (_BendStakeManager *BendStakeManagerTransactor) Initialize(opts *bind.TransactOpts, apeStaking_ common.Address, coinPool_ common.Address, nftPool_ common.Address, nftVault_ common.Address, stBayc_ common.Address, stMayc_ common.Address, stBakc_ common.Address) (*types.Transaction, error) {
	return _BendStakeManager.contract.Transact(opts, "initialize", apeStaking_, coinPool_, nftPool_, nftVault_, stBayc_, stMayc_, stBakc_)
}

// Initialize is a paid mutator transaction binding the contract method 0x35876476.
//
// Solidity: function initialize(address apeStaking_, address coinPool_, address nftPool_, address nftVault_, address stBayc_, address stMayc_, address stBakc_) returns()
func (_BendStakeManager *BendStakeManagerSession) Initialize(apeStaking_ common.Address, coinPool_ common.Address, nftPool_ common.Address, nftVault_ common.Address, stBayc_ common.Address, stMayc_ common.Address, stBakc_ common.Address) (*types.Transaction, error) {
	return _BendStakeManager.Contract.Initialize(&_BendStakeManager.TransactOpts, apeStaking_, coinPool_, nftPool_, nftVault_, stBayc_, stMayc_, stBakc_)
}

// Initialize is a paid mutator transaction binding the contract method 0x35876476.
//
// Solidity: function initialize(address apeStaking_, address coinPool_, address nftPool_, address nftVault_, address stBayc_, address stMayc_, address stBakc_) returns()
func (_BendStakeManager *BendStakeManagerTransactorSession) Initialize(apeStaking_ common.Address, coinPool_ common.Address, nftPool_ common.Address, nftVault_ common.Address, stBayc_ common.Address, stMayc_ common.Address, stBakc_ common.Address) (*types.Transaction, error) {
	return _BendStakeManager.Contract.Initialize(&_BendStakeManager.TransactOpts, apeStaking_, coinPool_, nftPool_, nftVault_, stBayc_, stMayc_, stBakc_)
}

// MintStNft is a paid mutator transaction binding the contract method 0x3d6f4334.
//
// Solidity: function mintStNft(address stNft_, address to_, uint256[] tokenIds_) returns()
func (_BendStakeManager *BendStakeManagerTransactor) MintStNft(opts *bind.TransactOpts, stNft_ common.Address, to_ common.Address, tokenIds_ []*big.Int) (*types.Transaction, error) {
	return _BendStakeManager.contract.Transact(opts, "mintStNft", stNft_, to_, tokenIds_)
}

// MintStNft is a paid mutator transaction binding the contract method 0x3d6f4334.
//
// Solidity: function mintStNft(address stNft_, address to_, uint256[] tokenIds_) returns()
func (_BendStakeManager *BendStakeManagerSession) MintStNft(stNft_ common.Address, to_ common.Address, tokenIds_ []*big.Int) (*types.Transaction, error) {
	return _BendStakeManager.Contract.MintStNft(&_BendStakeManager.TransactOpts, stNft_, to_, tokenIds_)
}

// MintStNft is a paid mutator transaction binding the contract method 0x3d6f4334.
//
// Solidity: function mintStNft(address stNft_, address to_, uint256[] tokenIds_) returns()
func (_BendStakeManager *BendStakeManagerTransactorSession) MintStNft(stNft_ common.Address, to_ common.Address, tokenIds_ []*big.Int) (*types.Transaction, error) {
	return _BendStakeManager.Contract.MintStNft(&_BendStakeManager.TransactOpts, stNft_, to_, tokenIds_)
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_BendStakeManager *BendStakeManagerTransactor) RenounceOwnership(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _BendStakeManager.contract.Transact(opts, "renounceOwnership")
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_BendStakeManager *BendStakeManagerSession) RenounceOwnership() (*types.Transaction, error) {
	return _BendStakeManager.Contract.RenounceOwnership(&_BendStakeManager.TransactOpts)
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_BendStakeManager *BendStakeManagerTransactorSession) RenounceOwnership() (*types.Transaction, error) {
	return _BendStakeManager.Contract.RenounceOwnership(&_BendStakeManager.TransactOpts)
}

// StakeApeCoin is a paid mutator transaction binding the contract method 0xadf31c80.
//
// Solidity: function stakeApeCoin(uint256 amount_) returns()
func (_BendStakeManager *BendStakeManagerTransactor) StakeApeCoin(opts *bind.TransactOpts, amount_ *big.Int) (*types.Transaction, error) {
	return _BendStakeManager.contract.Transact(opts, "stakeApeCoin", amount_)
}

// StakeApeCoin is a paid mutator transaction binding the contract method 0xadf31c80.
//
// Solidity: function stakeApeCoin(uint256 amount_) returns()
func (_BendStakeManager *BendStakeManagerSession) StakeApeCoin(amount_ *big.Int) (*types.Transaction, error) {
	return _BendStakeManager.Contract.StakeApeCoin(&_BendStakeManager.TransactOpts, amount_)
}

// StakeApeCoin is a paid mutator transaction binding the contract method 0xadf31c80.
//
// Solidity: function stakeApeCoin(uint256 amount_) returns()
func (_BendStakeManager *BendStakeManagerTransactorSession) StakeApeCoin(amount_ *big.Int) (*types.Transaction, error) {
	return _BendStakeManager.Contract.StakeApeCoin(&_BendStakeManager.TransactOpts, amount_)
}

// StakeBakc is a paid mutator transaction binding the contract method 0x79e52ef0.
//
// Solidity: function stakeBakc((uint128,uint128)[] baycPairs_, (uint128,uint128)[] maycPairs_) returns()
func (_BendStakeManager *BendStakeManagerTransactor) StakeBakc(opts *bind.TransactOpts, baycPairs_ []IApeCoinStakingPairNft, maycPairs_ []IApeCoinStakingPairNft) (*types.Transaction, error) {
	return _BendStakeManager.contract.Transact(opts, "stakeBakc", baycPairs_, maycPairs_)
}

// StakeBakc is a paid mutator transaction binding the contract method 0x79e52ef0.
//
// Solidity: function stakeBakc((uint128,uint128)[] baycPairs_, (uint128,uint128)[] maycPairs_) returns()
func (_BendStakeManager *BendStakeManagerSession) StakeBakc(baycPairs_ []IApeCoinStakingPairNft, maycPairs_ []IApeCoinStakingPairNft) (*types.Transaction, error) {
	return _BendStakeManager.Contract.StakeBakc(&_BendStakeManager.TransactOpts, baycPairs_, maycPairs_)
}

// StakeBakc is a paid mutator transaction binding the contract method 0x79e52ef0.
//
// Solidity: function stakeBakc((uint128,uint128)[] baycPairs_, (uint128,uint128)[] maycPairs_) returns()
func (_BendStakeManager *BendStakeManagerTransactorSession) StakeBakc(baycPairs_ []IApeCoinStakingPairNft, maycPairs_ []IApeCoinStakingPairNft) (*types.Transaction, error) {
	return _BendStakeManager.Contract.StakeBakc(&_BendStakeManager.TransactOpts, baycPairs_, maycPairs_)
}

// StakeBayc is a paid mutator transaction binding the contract method 0x28e54014.
//
// Solidity: function stakeBayc(uint256[] tokenIds_) returns()
func (_BendStakeManager *BendStakeManagerTransactor) StakeBayc(opts *bind.TransactOpts, tokenIds_ []*big.Int) (*types.Transaction, error) {
	return _BendStakeManager.contract.Transact(opts, "stakeBayc", tokenIds_)
}

// StakeBayc is a paid mutator transaction binding the contract method 0x28e54014.
//
// Solidity: function stakeBayc(uint256[] tokenIds_) returns()
func (_BendStakeManager *BendStakeManagerSession) StakeBayc(tokenIds_ []*big.Int) (*types.Transaction, error) {
	return _BendStakeManager.Contract.StakeBayc(&_BendStakeManager.TransactOpts, tokenIds_)
}

// StakeBayc is a paid mutator transaction binding the contract method 0x28e54014.
//
// Solidity: function stakeBayc(uint256[] tokenIds_) returns()
func (_BendStakeManager *BendStakeManagerTransactorSession) StakeBayc(tokenIds_ []*big.Int) (*types.Transaction, error) {
	return _BendStakeManager.Contract.StakeBayc(&_BendStakeManager.TransactOpts, tokenIds_)
}

// StakeMayc is a paid mutator transaction binding the contract method 0x2718fa2d.
//
// Solidity: function stakeMayc(uint256[] tokenIds_) returns()
func (_BendStakeManager *BendStakeManagerTransactor) StakeMayc(opts *bind.TransactOpts, tokenIds_ []*big.Int) (*types.Transaction, error) {
	return _BendStakeManager.contract.Transact(opts, "stakeMayc", tokenIds_)
}

// StakeMayc is a paid mutator transaction binding the contract method 0x2718fa2d.
//
// Solidity: function stakeMayc(uint256[] tokenIds_) returns()
func (_BendStakeManager *BendStakeManagerSession) StakeMayc(tokenIds_ []*big.Int) (*types.Transaction, error) {
	return _BendStakeManager.Contract.StakeMayc(&_BendStakeManager.TransactOpts, tokenIds_)
}

// StakeMayc is a paid mutator transaction binding the contract method 0x2718fa2d.
//
// Solidity: function stakeMayc(uint256[] tokenIds_) returns()
func (_BendStakeManager *BendStakeManagerTransactorSession) StakeMayc(tokenIds_ []*big.Int) (*types.Transaction, error) {
	return _BendStakeManager.Contract.StakeMayc(&_BendStakeManager.TransactOpts, tokenIds_)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_BendStakeManager *BendStakeManagerTransactor) TransferOwnership(opts *bind.TransactOpts, newOwner common.Address) (*types.Transaction, error) {
	return _BendStakeManager.contract.Transact(opts, "transferOwnership", newOwner)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_BendStakeManager *BendStakeManagerSession) TransferOwnership(newOwner common.Address) (*types.Transaction, error) {
	return _BendStakeManager.Contract.TransferOwnership(&_BendStakeManager.TransactOpts, newOwner)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_BendStakeManager *BendStakeManagerTransactorSession) TransferOwnership(newOwner common.Address) (*types.Transaction, error) {
	return _BendStakeManager.Contract.TransferOwnership(&_BendStakeManager.TransactOpts, newOwner)
}

// UnstakeApeCoin is a paid mutator transaction binding the contract method 0x36c7657c.
//
// Solidity: function unstakeApeCoin(uint256 amount_) returns()
func (_BendStakeManager *BendStakeManagerTransactor) UnstakeApeCoin(opts *bind.TransactOpts, amount_ *big.Int) (*types.Transaction, error) {
	return _BendStakeManager.contract.Transact(opts, "unstakeApeCoin", amount_)
}

// UnstakeApeCoin is a paid mutator transaction binding the contract method 0x36c7657c.
//
// Solidity: function unstakeApeCoin(uint256 amount_) returns()
func (_BendStakeManager *BendStakeManagerSession) UnstakeApeCoin(amount_ *big.Int) (*types.Transaction, error) {
	return _BendStakeManager.Contract.UnstakeApeCoin(&_BendStakeManager.TransactOpts, amount_)
}

// UnstakeApeCoin is a paid mutator transaction binding the contract method 0x36c7657c.
//
// Solidity: function unstakeApeCoin(uint256 amount_) returns()
func (_BendStakeManager *BendStakeManagerTransactorSession) UnstakeApeCoin(amount_ *big.Int) (*types.Transaction, error) {
	return _BendStakeManager.Contract.UnstakeApeCoin(&_BendStakeManager.TransactOpts, amount_)
}

// UnstakeBakc is a paid mutator transaction binding the contract method 0xd1d08e2d.
//
// Solidity: function unstakeBakc((uint128,uint128)[] baycPairs_, (uint128,uint128)[] maycPairs_) returns()
func (_BendStakeManager *BendStakeManagerTransactor) UnstakeBakc(opts *bind.TransactOpts, baycPairs_ []IApeCoinStakingPairNft, maycPairs_ []IApeCoinStakingPairNft) (*types.Transaction, error) {
	return _BendStakeManager.contract.Transact(opts, "unstakeBakc", baycPairs_, maycPairs_)
}

// UnstakeBakc is a paid mutator transaction binding the contract method 0xd1d08e2d.
//
// Solidity: function unstakeBakc((uint128,uint128)[] baycPairs_, (uint128,uint128)[] maycPairs_) returns()
func (_BendStakeManager *BendStakeManagerSession) UnstakeBakc(baycPairs_ []IApeCoinStakingPairNft, maycPairs_ []IApeCoinStakingPairNft) (*types.Transaction, error) {
	return _BendStakeManager.Contract.UnstakeBakc(&_BendStakeManager.TransactOpts, baycPairs_, maycPairs_)
}

// UnstakeBakc is a paid mutator transaction binding the contract method 0xd1d08e2d.
//
// Solidity: function unstakeBakc((uint128,uint128)[] baycPairs_, (uint128,uint128)[] maycPairs_) returns()
func (_BendStakeManager *BendStakeManagerTransactorSession) UnstakeBakc(baycPairs_ []IApeCoinStakingPairNft, maycPairs_ []IApeCoinStakingPairNft) (*types.Transaction, error) {
	return _BendStakeManager.Contract.UnstakeBakc(&_BendStakeManager.TransactOpts, baycPairs_, maycPairs_)
}

// UnstakeBayc is a paid mutator transaction binding the contract method 0xf12bde8c.
//
// Solidity: function unstakeBayc(uint256[] tokenIds_) returns()
func (_BendStakeManager *BendStakeManagerTransactor) UnstakeBayc(opts *bind.TransactOpts, tokenIds_ []*big.Int) (*types.Transaction, error) {
	return _BendStakeManager.contract.Transact(opts, "unstakeBayc", tokenIds_)
}

// UnstakeBayc is a paid mutator transaction binding the contract method 0xf12bde8c.
//
// Solidity: function unstakeBayc(uint256[] tokenIds_) returns()
func (_BendStakeManager *BendStakeManagerSession) UnstakeBayc(tokenIds_ []*big.Int) (*types.Transaction, error) {
	return _BendStakeManager.Contract.UnstakeBayc(&_BendStakeManager.TransactOpts, tokenIds_)
}

// UnstakeBayc is a paid mutator transaction binding the contract method 0xf12bde8c.
//
// Solidity: function unstakeBayc(uint256[] tokenIds_) returns()
func (_BendStakeManager *BendStakeManagerTransactorSession) UnstakeBayc(tokenIds_ []*big.Int) (*types.Transaction, error) {
	return _BendStakeManager.Contract.UnstakeBayc(&_BendStakeManager.TransactOpts, tokenIds_)
}

// UnstakeMayc is a paid mutator transaction binding the contract method 0x9b76c697.
//
// Solidity: function unstakeMayc(uint256[] tokenIds_) returns()
func (_BendStakeManager *BendStakeManagerTransactor) UnstakeMayc(opts *bind.TransactOpts, tokenIds_ []*big.Int) (*types.Transaction, error) {
	return _BendStakeManager.contract.Transact(opts, "unstakeMayc", tokenIds_)
}

// UnstakeMayc is a paid mutator transaction binding the contract method 0x9b76c697.
//
// Solidity: function unstakeMayc(uint256[] tokenIds_) returns()
func (_BendStakeManager *BendStakeManagerSession) UnstakeMayc(tokenIds_ []*big.Int) (*types.Transaction, error) {
	return _BendStakeManager.Contract.UnstakeMayc(&_BendStakeManager.TransactOpts, tokenIds_)
}

// UnstakeMayc is a paid mutator transaction binding the contract method 0x9b76c697.
//
// Solidity: function unstakeMayc(uint256[] tokenIds_) returns()
func (_BendStakeManager *BendStakeManagerTransactorSession) UnstakeMayc(tokenIds_ []*big.Int) (*types.Transaction, error) {
	return _BendStakeManager.Contract.UnstakeMayc(&_BendStakeManager.TransactOpts, tokenIds_)
}

// UpdateBotAdmin is a paid mutator transaction binding the contract method 0x0362dde5.
//
// Solidity: function updateBotAdmin(address botAdmin_) returns()
func (_BendStakeManager *BendStakeManagerTransactor) UpdateBotAdmin(opts *bind.TransactOpts, botAdmin_ common.Address) (*types.Transaction, error) {
	return _BendStakeManager.contract.Transact(opts, "updateBotAdmin", botAdmin_)
}

// UpdateBotAdmin is a paid mutator transaction binding the contract method 0x0362dde5.
//
// Solidity: function updateBotAdmin(address botAdmin_) returns()
func (_BendStakeManager *BendStakeManagerSession) UpdateBotAdmin(botAdmin_ common.Address) (*types.Transaction, error) {
	return _BendStakeManager.Contract.UpdateBotAdmin(&_BendStakeManager.TransactOpts, botAdmin_)
}

// UpdateBotAdmin is a paid mutator transaction binding the contract method 0x0362dde5.
//
// Solidity: function updateBotAdmin(address botAdmin_) returns()
func (_BendStakeManager *BendStakeManagerTransactorSession) UpdateBotAdmin(botAdmin_ common.Address) (*types.Transaction, error) {
	return _BendStakeManager.Contract.UpdateBotAdmin(&_BendStakeManager.TransactOpts, botAdmin_)
}

// UpdateFee is a paid mutator transaction binding the contract method 0x9012c4a8.
//
// Solidity: function updateFee(uint256 fee_) returns()
func (_BendStakeManager *BendStakeManagerTransactor) UpdateFee(opts *bind.TransactOpts, fee_ *big.Int) (*types.Transaction, error) {
	return _BendStakeManager.contract.Transact(opts, "updateFee", fee_)
}

// UpdateFee is a paid mutator transaction binding the contract method 0x9012c4a8.
//
// Solidity: function updateFee(uint256 fee_) returns()
func (_BendStakeManager *BendStakeManagerSession) UpdateFee(fee_ *big.Int) (*types.Transaction, error) {
	return _BendStakeManager.Contract.UpdateFee(&_BendStakeManager.TransactOpts, fee_)
}

// UpdateFee is a paid mutator transaction binding the contract method 0x9012c4a8.
//
// Solidity: function updateFee(uint256 fee_) returns()
func (_BendStakeManager *BendStakeManagerTransactorSession) UpdateFee(fee_ *big.Int) (*types.Transaction, error) {
	return _BendStakeManager.Contract.UpdateFee(&_BendStakeManager.TransactOpts, fee_)
}

// UpdateFeeRecipient is a paid mutator transaction binding the contract method 0xf160d369.
//
// Solidity: function updateFeeRecipient(address recipient_) returns()
func (_BendStakeManager *BendStakeManagerTransactor) UpdateFeeRecipient(opts *bind.TransactOpts, recipient_ common.Address) (*types.Transaction, error) {
	return _BendStakeManager.contract.Transact(opts, "updateFeeRecipient", recipient_)
}

// UpdateFeeRecipient is a paid mutator transaction binding the contract method 0xf160d369.
//
// Solidity: function updateFeeRecipient(address recipient_) returns()
func (_BendStakeManager *BendStakeManagerSession) UpdateFeeRecipient(recipient_ common.Address) (*types.Transaction, error) {
	return _BendStakeManager.Contract.UpdateFeeRecipient(&_BendStakeManager.TransactOpts, recipient_)
}

// UpdateFeeRecipient is a paid mutator transaction binding the contract method 0xf160d369.
//
// Solidity: function updateFeeRecipient(address recipient_) returns()
func (_BendStakeManager *BendStakeManagerTransactorSession) UpdateFeeRecipient(recipient_ common.Address) (*types.Transaction, error) {
	return _BendStakeManager.Contract.UpdateFeeRecipient(&_BendStakeManager.TransactOpts, recipient_)
}

// UpdateRewardsStrategy is a paid mutator transaction binding the contract method 0x1a1eac8e.
//
// Solidity: function updateRewardsStrategy(address nft_, address rewardsStrategy_) returns()
func (_BendStakeManager *BendStakeManagerTransactor) UpdateRewardsStrategy(opts *bind.TransactOpts, nft_ common.Address, rewardsStrategy_ common.Address) (*types.Transaction, error) {
	return _BendStakeManager.contract.Transact(opts, "updateRewardsStrategy", nft_, rewardsStrategy_)
}

// UpdateRewardsStrategy is a paid mutator transaction binding the contract method 0x1a1eac8e.
//
// Solidity: function updateRewardsStrategy(address nft_, address rewardsStrategy_) returns()
func (_BendStakeManager *BendStakeManagerSession) UpdateRewardsStrategy(nft_ common.Address, rewardsStrategy_ common.Address) (*types.Transaction, error) {
	return _BendStakeManager.Contract.UpdateRewardsStrategy(&_BendStakeManager.TransactOpts, nft_, rewardsStrategy_)
}

// UpdateRewardsStrategy is a paid mutator transaction binding the contract method 0x1a1eac8e.
//
// Solidity: function updateRewardsStrategy(address nft_, address rewardsStrategy_) returns()
func (_BendStakeManager *BendStakeManagerTransactorSession) UpdateRewardsStrategy(nft_ common.Address, rewardsStrategy_ common.Address) (*types.Transaction, error) {
	return _BendStakeManager.Contract.UpdateRewardsStrategy(&_BendStakeManager.TransactOpts, nft_, rewardsStrategy_)
}

// UpdateWithdrawStrategy is a paid mutator transaction binding the contract method 0xae9accd0.
//
// Solidity: function updateWithdrawStrategy(address withdrawStrategy_) returns()
func (_BendStakeManager *BendStakeManagerTransactor) UpdateWithdrawStrategy(opts *bind.TransactOpts, withdrawStrategy_ common.Address) (*types.Transaction, error) {
	return _BendStakeManager.contract.Transact(opts, "updateWithdrawStrategy", withdrawStrategy_)
}

// UpdateWithdrawStrategy is a paid mutator transaction binding the contract method 0xae9accd0.
//
// Solidity: function updateWithdrawStrategy(address withdrawStrategy_) returns()
func (_BendStakeManager *BendStakeManagerSession) UpdateWithdrawStrategy(withdrawStrategy_ common.Address) (*types.Transaction, error) {
	return _BendStakeManager.Contract.UpdateWithdrawStrategy(&_BendStakeManager.TransactOpts, withdrawStrategy_)
}

// UpdateWithdrawStrategy is a paid mutator transaction binding the contract method 0xae9accd0.
//
// Solidity: function updateWithdrawStrategy(address withdrawStrategy_) returns()
func (_BendStakeManager *BendStakeManagerTransactorSession) UpdateWithdrawStrategy(withdrawStrategy_ common.Address) (*types.Transaction, error) {
	return _BendStakeManager.Contract.UpdateWithdrawStrategy(&_BendStakeManager.TransactOpts, withdrawStrategy_)
}

// WithdrawApeCoin is a paid mutator transaction binding the contract method 0xc0d103c8.
//
// Solidity: function withdrawApeCoin(uint256 required) returns(uint256 withdrawn)
func (_BendStakeManager *BendStakeManagerTransactor) WithdrawApeCoin(opts *bind.TransactOpts, required *big.Int) (*types.Transaction, error) {
	return _BendStakeManager.contract.Transact(opts, "withdrawApeCoin", required)
}

// WithdrawApeCoin is a paid mutator transaction binding the contract method 0xc0d103c8.
//
// Solidity: function withdrawApeCoin(uint256 required) returns(uint256 withdrawn)
func (_BendStakeManager *BendStakeManagerSession) WithdrawApeCoin(required *big.Int) (*types.Transaction, error) {
	return _BendStakeManager.Contract.WithdrawApeCoin(&_BendStakeManager.TransactOpts, required)
}

// WithdrawApeCoin is a paid mutator transaction binding the contract method 0xc0d103c8.
//
// Solidity: function withdrawApeCoin(uint256 required) returns(uint256 withdrawn)
func (_BendStakeManager *BendStakeManagerTransactorSession) WithdrawApeCoin(required *big.Int) (*types.Transaction, error) {
	return _BendStakeManager.Contract.WithdrawApeCoin(&_BendStakeManager.TransactOpts, required)
}

// WithdrawRefund is a paid mutator transaction binding the contract method 0xa16c86f7.
//
// Solidity: function withdrawRefund(address nft_) returns()
func (_BendStakeManager *BendStakeManagerTransactor) WithdrawRefund(opts *bind.TransactOpts, nft_ common.Address) (*types.Transaction, error) {
	return _BendStakeManager.contract.Transact(opts, "withdrawRefund", nft_)
}

// WithdrawRefund is a paid mutator transaction binding the contract method 0xa16c86f7.
//
// Solidity: function withdrawRefund(address nft_) returns()
func (_BendStakeManager *BendStakeManagerSession) WithdrawRefund(nft_ common.Address) (*types.Transaction, error) {
	return _BendStakeManager.Contract.WithdrawRefund(&_BendStakeManager.TransactOpts, nft_)
}

// WithdrawRefund is a paid mutator transaction binding the contract method 0xa16c86f7.
//
// Solidity: function withdrawRefund(address nft_) returns()
func (_BendStakeManager *BendStakeManagerTransactorSession) WithdrawRefund(nft_ common.Address) (*types.Transaction, error) {
	return _BendStakeManager.Contract.WithdrawRefund(&_BendStakeManager.TransactOpts, nft_)
}

// WithdrawTotalRefund is a paid mutator transaction binding the contract method 0x6128eb4f.
//
// Solidity: function withdrawTotalRefund() returns()
func (_BendStakeManager *BendStakeManagerTransactor) WithdrawTotalRefund(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _BendStakeManager.contract.Transact(opts, "withdrawTotalRefund")
}

// WithdrawTotalRefund is a paid mutator transaction binding the contract method 0x6128eb4f.
//
// Solidity: function withdrawTotalRefund() returns()
func (_BendStakeManager *BendStakeManagerSession) WithdrawTotalRefund() (*types.Transaction, error) {
	return _BendStakeManager.Contract.WithdrawTotalRefund(&_BendStakeManager.TransactOpts)
}

// WithdrawTotalRefund is a paid mutator transaction binding the contract method 0x6128eb4f.
//
// Solidity: function withdrawTotalRefund() returns()
func (_BendStakeManager *BendStakeManagerTransactorSession) WithdrawTotalRefund() (*types.Transaction, error) {
	return _BendStakeManager.Contract.WithdrawTotalRefund(&_BendStakeManager.TransactOpts)
}

// BendStakeManagerInitializedIterator is returned from FilterInitialized and is used to iterate over the raw logs and unpacked data for Initialized events raised by the BendStakeManager contract.
type BendStakeManagerInitializedIterator struct {
	Event *BendStakeManagerInitialized // Event containing the contract specifics and raw log

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
func (it *BendStakeManagerInitializedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(BendStakeManagerInitialized)
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
		it.Event = new(BendStakeManagerInitialized)
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
func (it *BendStakeManagerInitializedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *BendStakeManagerInitializedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// BendStakeManagerInitialized represents a Initialized event raised by the BendStakeManager contract.
type BendStakeManagerInitialized struct {
	Version uint8
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterInitialized is a free log retrieval operation binding the contract event 0x7f26b83ff96e1f2b6a682f133852f6798a09c465da95921460cefb3847402498.
//
// Solidity: event Initialized(uint8 version)
func (_BendStakeManager *BendStakeManagerFilterer) FilterInitialized(opts *bind.FilterOpts) (*BendStakeManagerInitializedIterator, error) {

	logs, sub, err := _BendStakeManager.contract.FilterLogs(opts, "Initialized")
	if err != nil {
		return nil, err
	}
	return &BendStakeManagerInitializedIterator{contract: _BendStakeManager.contract, event: "Initialized", logs: logs, sub: sub}, nil
}

// WatchInitialized is a free log subscription operation binding the contract event 0x7f26b83ff96e1f2b6a682f133852f6798a09c465da95921460cefb3847402498.
//
// Solidity: event Initialized(uint8 version)
func (_BendStakeManager *BendStakeManagerFilterer) WatchInitialized(opts *bind.WatchOpts, sink chan<- *BendStakeManagerInitialized) (event.Subscription, error) {

	logs, sub, err := _BendStakeManager.contract.WatchLogs(opts, "Initialized")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(BendStakeManagerInitialized)
				if err := _BendStakeManager.contract.UnpackLog(event, "Initialized", log); err != nil {
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

// ParseInitialized is a log parse operation binding the contract event 0x7f26b83ff96e1f2b6a682f133852f6798a09c465da95921460cefb3847402498.
//
// Solidity: event Initialized(uint8 version)
func (_BendStakeManager *BendStakeManagerFilterer) ParseInitialized(log types.Log) (*BendStakeManagerInitialized, error) {
	event := new(BendStakeManagerInitialized)
	if err := _BendStakeManager.contract.UnpackLog(event, "Initialized", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// BendStakeManagerOwnershipTransferredIterator is returned from FilterOwnershipTransferred and is used to iterate over the raw logs and unpacked data for OwnershipTransferred events raised by the BendStakeManager contract.
type BendStakeManagerOwnershipTransferredIterator struct {
	Event *BendStakeManagerOwnershipTransferred // Event containing the contract specifics and raw log

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
func (it *BendStakeManagerOwnershipTransferredIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(BendStakeManagerOwnershipTransferred)
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
		it.Event = new(BendStakeManagerOwnershipTransferred)
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
func (it *BendStakeManagerOwnershipTransferredIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *BendStakeManagerOwnershipTransferredIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// BendStakeManagerOwnershipTransferred represents a OwnershipTransferred event raised by the BendStakeManager contract.
type BendStakeManagerOwnershipTransferred struct {
	PreviousOwner common.Address
	NewOwner      common.Address
	Raw           types.Log // Blockchain specific contextual infos
}

// FilterOwnershipTransferred is a free log retrieval operation binding the contract event 0x8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0.
//
// Solidity: event OwnershipTransferred(address indexed previousOwner, address indexed newOwner)
func (_BendStakeManager *BendStakeManagerFilterer) FilterOwnershipTransferred(opts *bind.FilterOpts, previousOwner []common.Address, newOwner []common.Address) (*BendStakeManagerOwnershipTransferredIterator, error) {

	var previousOwnerRule []interface{}
	for _, previousOwnerItem := range previousOwner {
		previousOwnerRule = append(previousOwnerRule, previousOwnerItem)
	}
	var newOwnerRule []interface{}
	for _, newOwnerItem := range newOwner {
		newOwnerRule = append(newOwnerRule, newOwnerItem)
	}

	logs, sub, err := _BendStakeManager.contract.FilterLogs(opts, "OwnershipTransferred", previousOwnerRule, newOwnerRule)
	if err != nil {
		return nil, err
	}
	return &BendStakeManagerOwnershipTransferredIterator{contract: _BendStakeManager.contract, event: "OwnershipTransferred", logs: logs, sub: sub}, nil
}

// WatchOwnershipTransferred is a free log subscription operation binding the contract event 0x8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0.
//
// Solidity: event OwnershipTransferred(address indexed previousOwner, address indexed newOwner)
func (_BendStakeManager *BendStakeManagerFilterer) WatchOwnershipTransferred(opts *bind.WatchOpts, sink chan<- *BendStakeManagerOwnershipTransferred, previousOwner []common.Address, newOwner []common.Address) (event.Subscription, error) {

	var previousOwnerRule []interface{}
	for _, previousOwnerItem := range previousOwner {
		previousOwnerRule = append(previousOwnerRule, previousOwnerItem)
	}
	var newOwnerRule []interface{}
	for _, newOwnerItem := range newOwner {
		newOwnerRule = append(newOwnerRule, newOwnerItem)
	}

	logs, sub, err := _BendStakeManager.contract.WatchLogs(opts, "OwnershipTransferred", previousOwnerRule, newOwnerRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(BendStakeManagerOwnershipTransferred)
				if err := _BendStakeManager.contract.UnpackLog(event, "OwnershipTransferred", log); err != nil {
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
func (_BendStakeManager *BendStakeManagerFilterer) ParseOwnershipTransferred(log types.Log) (*BendStakeManagerOwnershipTransferred, error) {
	event := new(BendStakeManagerOwnershipTransferred)
	if err := _BendStakeManager.contract.UnpackLog(event, "OwnershipTransferred", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}
