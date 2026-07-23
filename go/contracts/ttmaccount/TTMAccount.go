// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package ttmaccount

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

// PartnerConfigurationService is an auto generated low-level Go binding around an user-defined struct.
type PartnerConfigurationService struct {
	RestrictedRate bool
	Capabilities   []string
}

// TtmaccountMetaData contains all meta data concerning the Ttmaccount contract.
var TtmaccountMetaData = &bind.MetaData{
	ABI: "[{\"inputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"constructor\"},{\"inputs\":[],\"name\":\"AccessControlBadConfirmation\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"},{\"internalType\":\"bytes32\",\"name\":\"neededRole\",\"type\":\"bytes32\"}],\"name\":\"AccessControlUnauthorizedAccount\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"target\",\"type\":\"address\"}],\"name\":\"AddressEmptyCode\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"serviceHash\",\"type\":\"bytes32\"},{\"internalType\":\"string\",\"name\":\"capability\",\"type\":\"string\"}],\"name\":\"CapabilityDoesNotExist\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"implementation\",\"type\":\"address\"}],\"name\":\"ERC1967InvalidImplementation\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"ERC1967NonPayable\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"FailedCall\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"limit\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"period\",\"type\":\"uint256\"}],\"name\":\"GasMoneyValueOutOfRange\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"balance\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"needed\",\"type\":\"uint256\"}],\"name\":\"InsufficientBalance\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"InvalidInitialization\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"NotInitializing\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"token\",\"type\":\"address\"}],\"name\":\"PaymentTokenAlreadyExists\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"token\",\"type\":\"address\"}],\"name\":\"PaymentTokenDoesNotExist\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"pubKeyAddress\",\"type\":\"address\"}],\"name\":\"PublicKeyAlreadyExists\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"pubKeyAddress\",\"type\":\"address\"}],\"name\":\"PublicKeyDoesNotExist\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"ReentrancyGuardReentrantCall\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"token\",\"type\":\"address\"}],\"name\":\"SafeERC20FailedOperation\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"serviceHash\",\"type\":\"bytes32\"}],\"name\":\"ServiceAlreadyExists\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"serviceHash\",\"type\":\"bytes32\"}],\"name\":\"ServiceDoesNotExist\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"ServiceNotRegistered\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"latestImplementation\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"newImplementation\",\"type\":\"address\"}],\"name\":\"TTMAccountImplementationMismatch\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"oldImplementation\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"newImplementation\",\"type\":\"address\"}],\"name\":\"TTMAccountNoUpgradeNeeded\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"TransferToZeroAddress\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"UUPSUnauthorizedCallContext\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"slot\",\"type\":\"bytes32\"}],\"name\":\"UUPSUnsupportedProxiableUUID\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"serviceHash\",\"type\":\"bytes32\"}],\"name\":\"WantedServiceAlreadyExists\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"serviceHash\",\"type\":\"bytes32\"}],\"name\":\"WantedServiceDoesNotExist\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"limit\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"WithdrawalLimitExceeded\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"limit\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"WithdrawalLimitExceededForPeriod\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"ZeroAddress\",\"type\":\"error\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"sender\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"Deposit\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"withdrawer\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"GasMoneyWithdrawal\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"limit\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"period\",\"type\":\"uint256\"}],\"name\":\"GasMoneyWithdrawalUpdated\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint64\",\"name\":\"version\",\"type\":\"uint64\"}],\"name\":\"Initialized\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"bot\",\"type\":\"address\"}],\"name\":\"MessengerBotAdded\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"bot\",\"type\":\"address\"}],\"name\":\"MessengerBotRemoved\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"token\",\"type\":\"address\"}],\"name\":\"PaymentTokenAdded\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"token\",\"type\":\"address\"}],\"name\":\"PaymentTokenRemoved\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"pubKeyAddress\",\"type\":\"address\"}],\"name\":\"PublicKeyAdded\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"pubKeyAddress\",\"type\":\"address\"}],\"name\":\"PublicKeyRemoved\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"bytes32\",\"name\":\"role\",\"type\":\"bytes32\"},{\"indexed\":true,\"internalType\":\"bytes32\",\"name\":\"previousAdminRole\",\"type\":\"bytes32\"},{\"indexed\":true,\"internalType\":\"bytes32\",\"name\":\"newAdminRole\",\"type\":\"bytes32\"}],\"name\":\"RoleAdminChanged\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"bytes32\",\"name\":\"role\",\"type\":\"bytes32\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"sender\",\"type\":\"address\"}],\"name\":\"RoleGranted\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"bytes32\",\"name\":\"role\",\"type\":\"bytes32\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"sender\",\"type\":\"address\"}],\"name\":\"RoleRevoked\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"bytes32\",\"name\":\"serviceHash\",\"type\":\"bytes32\"}],\"name\":\"ServiceAdded\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"bytes32\",\"name\":\"serviceHash\",\"type\":\"bytes32\"}],\"name\":\"ServiceCapabilitiesUpdated\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"bytes32\",\"name\":\"serviceHash\",\"type\":\"bytes32\"},{\"indexed\":false,\"internalType\":\"string\",\"name\":\"capability\",\"type\":\"string\"}],\"name\":\"ServiceCapabilityAdded\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"bytes32\",\"name\":\"serviceHash\",\"type\":\"bytes32\"},{\"indexed\":false,\"internalType\":\"string\",\"name\":\"capability\",\"type\":\"string\"}],\"name\":\"ServiceCapabilityRemoved\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"bytes32\",\"name\":\"serviceHash\",\"type\":\"bytes32\"}],\"name\":\"ServiceRemoved\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"bytes32\",\"name\":\"serviceHash\",\"type\":\"bytes32\"},{\"indexed\":false,\"internalType\":\"bool\",\"name\":\"restrictedRate\",\"type\":\"bool\"}],\"name\":\"ServiceRestrictedRateUpdated\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"oldImplementation\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"newImplementation\",\"type\":\"address\"}],\"name\":\"TTMAccountUpgraded\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"implementation\",\"type\":\"address\"}],\"name\":\"Upgraded\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"bytes32\",\"name\":\"serviceHash\",\"type\":\"bytes32\"}],\"name\":\"WantedServiceAdded\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"bytes32\",\"name\":\"serviceHash\",\"type\":\"bytes32\"}],\"name\":\"WantedServiceRemoved\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"receiver\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"Withdraw\",\"type\":\"event\"},{\"inputs\":[],\"name\":\"BOOKING_OPERATOR_ROLE\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"BOT_ADMIN_ROLE\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"DEFAULT_ADMIN_ROLE\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"GAS_WITHDRAWER_ROLE\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"MESSENGER_BOT_ROLE\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"SERVICE_ADMIN_ROLE\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"UPGRADER_ROLE\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"UPGRADE_INTERFACE_VERSION\",\"outputs\":[{\"internalType\":\"string\",\"name\":\"\",\"type\":\"string\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"WITHDRAWER_ROLE\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"tokenId\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"refundAmount\",\"type\":\"uint256\"}],\"name\":\"acceptCancellation\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"bot\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"gasMoney\",\"type\":\"uint256\"}],\"name\":\"addMessengerBot\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"pubKeyAddress\",\"type\":\"address\"},{\"internalType\":\"bytes\",\"name\":\"data\",\"type\":\"bytes\"}],\"name\":\"addPublicKey\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"serviceHash\",\"type\":\"bytes32\"},{\"internalType\":\"bool\",\"name\":\"restrictedRate\",\"type\":\"bool\"},{\"internalType\":\"string[]\",\"name\":\"capabilities\",\"type\":\"string[]\"}],\"name\":\"addService\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"serviceHash\",\"type\":\"bytes32\"},{\"internalType\":\"string\",\"name\":\"capability\",\"type\":\"string\"}],\"name\":\"addServiceCapability\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_supportedToken\",\"type\":\"address\"}],\"name\":\"addSupportedToken\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32[]\",\"name\":\"serviceHashes\",\"type\":\"bytes32[]\"}],\"name\":\"addWantedServices\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"contractIERC721\",\"name\":\"token\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"tokenId\",\"type\":\"uint256\"}],\"name\":\"approveERC721\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"tokenId\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"expectedPrice\",\"type\":\"uint256\"},{\"internalType\":\"contractIERC20\",\"name\":\"expectedPaymentToken\",\"type\":\"address\"}],\"name\":\"buyBookingToken\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"tokenId\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"refundAmount\",\"type\":\"uint256\"},{\"internalType\":\"uint16\",\"name\":\"counterReason\",\"type\":\"uint16\"},{\"internalType\":\"uint16\",\"name\":\"counterReasonVersion\",\"type\":\"uint16\"}],\"name\":\"counterCancellation\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"tokenId\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"refundAmount\",\"type\":\"uint256\"}],\"name\":\"finalizeCancellation\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"getAllServiceHashes\",\"outputs\":[{\"internalType\":\"bytes32[]\",\"name\":\"serviceHashes\",\"type\":\"bytes32[]\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"getBookingTokenAddress\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"getGasMoneyWithdrawal\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"withdrawalLimit\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"withdrawalPeriod\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"}],\"name\":\"getGasMoneyWithdrawalForAccount\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"periodStart\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"withdrawnAmount\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"getManagerAddress\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"pubKeyAddress\",\"type\":\"address\"}],\"name\":\"getPublicKey\",\"outputs\":[{\"internalType\":\"bytes\",\"name\":\"data\",\"type\":\"bytes\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"getPublicKeysAddresses\",\"outputs\":[{\"internalType\":\"address[]\",\"name\":\"pubKeyAddresses\",\"type\":\"address[]\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"role\",\"type\":\"bytes32\"}],\"name\":\"getRoleAdmin\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"role\",\"type\":\"bytes32\"},{\"internalType\":\"uint256\",\"name\":\"index\",\"type\":\"uint256\"}],\"name\":\"getRoleMember\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"role\",\"type\":\"bytes32\"}],\"name\":\"getRoleMemberCount\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"role\",\"type\":\"bytes32\"}],\"name\":\"getRoleMembers\",\"outputs\":[{\"internalType\":\"address[]\",\"name\":\"\",\"type\":\"address[]\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"serviceHash\",\"type\":\"bytes32\"}],\"name\":\"getService\",\"outputs\":[{\"components\":[{\"internalType\":\"bool\",\"name\":\"_restrictedRate\",\"type\":\"bool\"},{\"internalType\":\"string[]\",\"name\":\"_capabilities\",\"type\":\"string[]\"}],\"internalType\":\"structPartnerConfiguration.Service\",\"name\":\"service\",\"type\":\"tuple\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"serviceHash\",\"type\":\"bytes32\"}],\"name\":\"getServiceCapabilities\",\"outputs\":[{\"internalType\":\"string[]\",\"name\":\"capabilities\",\"type\":\"string[]\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"serviceHash\",\"type\":\"bytes32\"}],\"name\":\"getServiceRestrictedRate\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"restrictedRate\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"getSupportedServices\",\"outputs\":[{\"internalType\":\"bytes32[]\",\"name\":\"serviceHashes\",\"type\":\"bytes32[]\"},{\"components\":[{\"internalType\":\"bool\",\"name\":\"_restrictedRate\",\"type\":\"bool\"},{\"internalType\":\"string[]\",\"name\":\"_capabilities\",\"type\":\"string[]\"}],\"internalType\":\"structPartnerConfiguration.Service[]\",\"name\":\"services\",\"type\":\"tuple[]\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"offset\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"limit\",\"type\":\"uint256\"}],\"name\":\"getSupportedServicesSlice\",\"outputs\":[{\"internalType\":\"bytes32[]\",\"name\":\"serviceHashes\",\"type\":\"bytes32[]\"},{\"components\":[{\"internalType\":\"bool\",\"name\":\"_restrictedRate\",\"type\":\"bool\"},{\"internalType\":\"string[]\",\"name\":\"_capabilities\",\"type\":\"string[]\"}],\"internalType\":\"structPartnerConfiguration.Service[]\",\"name\":\"services\",\"type\":\"tuple[]\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"getSupportedTokens\",\"outputs\":[{\"internalType\":\"address[]\",\"name\":\"tokens\",\"type\":\"address[]\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"getWantedServiceHashes\",\"outputs\":[{\"internalType\":\"bytes32[]\",\"name\":\"serviceHashes\",\"type\":\"bytes32[]\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"role\",\"type\":\"bytes32\"},{\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"}],\"name\":\"grantRole\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"role\",\"type\":\"bytes32\"},{\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"}],\"name\":\"hasRole\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"manager\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"bookingToken\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"defaultAdmin\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"upgrader\",\"type\":\"address\"}],\"name\":\"initialize\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"tokenId\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"refundAmount\",\"type\":\"uint256\"},{\"internalType\":\"uint16\",\"name\":\"cancellationReason\",\"type\":\"uint16\"},{\"internalType\":\"uint16\",\"name\":\"cancellationReasonVersion\",\"type\":\"uint16\"}],\"name\":\"initiateCancellation\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"bot\",\"type\":\"address\"}],\"name\":\"isBotAllowed\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"serviceHash\",\"type\":\"bytes32\"}],\"name\":\"isServiceSupported\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_token\",\"type\":\"address\"}],\"name\":\"isSupportedToken\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"supported\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"reservedFor\",\"type\":\"address\"},{\"internalType\":\"string\",\"name\":\"uri\",\"type\":\"string\"},{\"internalType\":\"uint256\",\"name\":\"expirationTimestamp\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"price\",\"type\":\"uint256\"},{\"internalType\":\"contractIERC20\",\"name\":\"paymentToken\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"offchainPaymentCurrency\",\"type\":\"uint256\"},{\"internalType\":\"bool\",\"name\":\"cancellable\",\"type\":\"bool\"}],\"name\":\"mintBookingToken\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"},{\"internalType\":\"bytes\",\"name\":\"\",\"type\":\"bytes\"}],\"name\":\"onERC721Received\",\"outputs\":[{\"internalType\":\"bytes4\",\"name\":\"\",\"type\":\"bytes4\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"proxiableUUID\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"tokenId\",\"type\":\"uint256\"}],\"name\":\"recordExpiration\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"tokenId\",\"type\":\"uint256\"},{\"internalType\":\"uint16\",\"name\":\"rejectionReason\",\"type\":\"uint16\"},{\"internalType\":\"uint16\",\"name\":\"rejectionReasonVersion\",\"type\":\"uint16\"}],\"name\":\"rejectCancellation\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"removeAllServices\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"bot\",\"type\":\"address\"}],\"name\":\"removeMessengerBot\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"pubKeyAddress\",\"type\":\"address\"}],\"name\":\"removePublicKey\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"serviceHash\",\"type\":\"bytes32\"}],\"name\":\"removeService\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"serviceHash\",\"type\":\"bytes32\"},{\"internalType\":\"string\",\"name\":\"capability\",\"type\":\"string\"}],\"name\":\"removeServiceCapability\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_supportedToken\",\"type\":\"address\"}],\"name\":\"removeSupportedToken\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32[]\",\"name\":\"serviceHashes\",\"type\":\"bytes32[]\"}],\"name\":\"removeWantedServices\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"role\",\"type\":\"bytes32\"},{\"internalType\":\"address\",\"name\":\"callerConfirmation\",\"type\":\"address\"}],\"name\":\"renounceRole\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"role\",\"type\":\"bytes32\"},{\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"}],\"name\":\"revokeRole\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"limit\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"period\",\"type\":\"uint256\"}],\"name\":\"setGasMoneyWithdrawal\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"serviceHash\",\"type\":\"bytes32\"},{\"internalType\":\"string[]\",\"name\":\"capabilities\",\"type\":\"string[]\"}],\"name\":\"setServiceCapabilities\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"serviceHash\",\"type\":\"bytes32\"},{\"internalType\":\"bool\",\"name\":\"restrictedRate\",\"type\":\"bool\"}],\"name\":\"setServiceRestrictedRate\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes4\",\"name\":\"interfaceId\",\"type\":\"bytes4\"}],\"name\":\"supportsInterface\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"contractIERC20\",\"name\":\"token\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"transferERC20\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"contractIERC721\",\"name\":\"token\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"tokenId\",\"type\":\"uint256\"}],\"name\":\"transferERC721\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"newImplementation\",\"type\":\"address\"},{\"internalType\":\"bytes\",\"name\":\"data\",\"type\":\"bytes\"}],\"name\":\"upgradeToAndCall\",\"outputs\":[],\"stateMutability\":\"payable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"addresspayable\",\"name\":\"recipient\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"withdraw\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"tokenId\",\"type\":\"uint256\"},{\"internalType\":\"uint16\",\"name\":\"reason\",\"type\":\"uint16\"},{\"internalType\":\"uint16\",\"name\":\"reasonVersion\",\"type\":\"uint16\"}],\"name\":\"withdrawCancellation\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"withdrawGasMoney\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"stateMutability\":\"payable\",\"type\":\"receive\"}]",
	Bin: "0x60a06040523060805234801562000014575f80fd5b506200001f62000025565b620000d9565b7ff0c57e16840df040f15088dc2f81fe391c3923bec73e23a9662efc9c229c6a00805468010000000000000000900460ff1615620000765760405163f92ee8a960e01b815260040160405180910390fd5b80546001600160401b0390811614620000d65780546001600160401b0319166001600160401b0390811782556040519081527fc7f505b2f371ae2175ee4913f4499e1f2633a7b5936321eed1cdaeb6115181d29060200160405180910390a15b50565b6080516151ee620001005f395f8181612be001528181612c09015261302701526151ee5ff3fe608060405260043610610437575f3560e01c8063852b3ccb11610237578063cd9ef9141161013c578063e7bfce9a116100b7578063f6b5174011610087578063f7e45f091161006d578063f7e45f0914610e8c578063f8c8765e14610eab578063f8e191ac14610eca575f80fd5b8063f6b5174014610e3a578063f72c0d8b14610e59575f80fd5b8063e7bfce9a14610d62578063ea79d07a14610d81578063ee3b641f14610d95578063f3fef3a314610e1b575f80fd5b8063d547741f1161010c578063e0b78add116100f2578063e0b78add14610d05578063e26a61bb14610d24578063e5a6725c14610d43575f80fd5b8063d547741f14610cba578063da47d85614610cd9575f80fd5b8063cd9ef91414610c48578063d09445c214610c67578063d3884c3f14610c87578063d3c7c2c714610ca6575f80fd5b8063a3246ad3116101cc578063be6671881161019c578063c6640e6811610182578063c6640e6814610beb578063ca15c87314610c0a578063ccde65dc14610c29575f80fd5b8063be66718814610b90578063c162d7da14610baf575f80fd5b8063a3246ad314610ae9578063ad3cb1cc14610b15578063b82923fb14610b5d578063bd252c1c14610b71575f80fd5b80639010d07c116102075780639010d07c14610a3557806391d1485414610a545780639db5dbe414610ab7578063a217fddf14610ad6575f80fd5b8063852b3ccb14610997578063857cdbb8146109b757806385f438c1146109e35780638f69347d14610a16575f80fd5b80634f1ef2861161033d5780636d69fcaf116102d257806374fe60e9116102a25780637c5d62b3116102885780637c5d62b3146109455780637eec56c71461096457806382010fb114610978575f80fd5b806374fe60e9146109075780637631919014610926575f80fd5b80636d69fcaf146108775780636fc22cd11461089657806372afa328146108b557806374aa2048146108e8575f80fd5b80635ae733c91161030d5780635ae733c9146107ad5780635c988994146107cc5780635e07f869146107eb578063658db0af14610817575f80fd5b80634f1ef2861461072a5780634f3f46391461073d57806351889d6b1461077a57806352d1902d14610799575f80fd5b80631c54f0f7116103cd5780632f2ff15d1161039d57806336568abe1161038357806336568abe146106c4578063383aba87146106e357806342072bbd14610716575f80fd5b80632f2ff15d146106725780633374627414610691575f80fd5b80631c54f0f7146105ba578063240028e8146105d9578063248a9ca3146105f85780632a11938014610653575f80fd5b80630d5e2e16116104085780630d5e2e1614610518578063136f50ca14610537578063150b7a02146105585780631aca63761461059b575f80fd5b8062a7230a1461047757806301ffc9a71461049857806304a3d81e146104cc57806308b1565a146104eb575f80fd5b366104735760405134815233907fe1fffcc4923d04b559f4d29a8bfc6cda04eb5b0d3c460751c2402c5c5cc9109c9060200160405180910390a2005b5f80fd5b348015610482575f80fd5b506104966104913660046144d7565b610ee9565b005b3480156104a3575f80fd5b506104b76104b2366004614515565b610f90565b60405190151581526020015b60405180910390f35b3480156104d7575f80fd5b506104966104e63660046145a4565b610fba565b3480156104f6575f80fd5b5061050a610505366004614635565b611073565b6040516104c3929190614754565b348015610523575f80fd5b506104966105323660046147d8565b611240565b348015610542575f80fd5b5061054b6112a2565b6040516104c39190614802565b348015610563575f80fd5b50610582610572366004614888565b630a85bd0160e11b949350505050565b6040516001600160e01b031990911681526020016104c3565b3480156105a6575f80fd5b506104966105b53660046144d7565b6112e1565b3480156105c5575f80fd5b506104966105d4366004614635565b611386565b3480156105e4575f80fd5b506104b76105f33660046148f0565b61143b565b348015610603575f80fd5b5061064561061236600461490b565b5f9081527f02dd7bc7dec4dceedda775e58dd541e08a116c6c53815c0bd028192f7b626800602052604090206001015490565b6040519081526020016104c3565b34801561065e575f80fd5b5061049661066d366004614933565b61147b565b34801561067d575f80fd5b5061049661068c36600461496c565b611532565b34801561069c575f80fd5b506106457fc6ae66da7afbd0becbe49ddf95a9256901d87f1bb5e3d43236030e5e4a585b6d81565b3480156106cf575f80fd5b506104966106de36600461496c565b61157b565b3480156106ee575f80fd5b506106457fe562760eaa817d85ec1bf58364c4d65adb65d99d113c6785ef9aa66567076c9581565b348015610721575f80fd5b5061054b6115c7565b61049661073836600461499a565b6115e0565b348015610748575f80fd5b505f80516020615199833981519152546001600160a01b03165b6040516001600160a01b0390911681526020016104c3565b348015610785575f80fd5b506104966107943660046149e7565b6115ff565b3480156107a4575f80fd5b506106456116d9565b3480156107b8575f80fd5b506104966107c7366004614a11565b611707565b3480156107d7575f80fd5b506104966107e636600461490b565b611758565b3480156107f6575f80fd5b5061080a61080536600461490b565b6117c0565b6040516104c39190614a3f565b348015610822575f80fd5b507fc7ae2c65fdae475b1bb2dd4079b252a7d893d0822b91a40dd25c763e583f7701546001600160801b03811690600160801b900467ffffffffffffffff165b604080519283526020830191909152016104c3565b348015610882575f80fd5b506104966108913660046148f0565b6118bb565b3480156108a1575f80fd5b506104966108b0366004614635565b6118db565b3480156108c0575f80fd5b506106457fe9e633d75f5fd429fed1611b4a64c1a1d6af654653894bd204eb46dbd457621a81565b3480156108f3575f80fd5b50610496610902366004614a51565b61190f565b348015610912575f80fd5b50610496610921366004614933565b6119d8565b348015610931575f80fd5b506104966109403660046148f0565b611a27565b348015610950575f80fd5b5061049661095f366004614b1a565b611a47565b34801561096f575f80fd5b5061050a611aa2565b348015610983575f80fd5b506104966109923660046145a4565b611b5b565b3480156109a2575f80fd5b506106455f8051602061517983398151915281565b3480156109c2575f80fd5b506109d66109d13660046148f0565b611bed565b6040516104c39190614b6d565b3480156109ee575f80fd5b506106457f10dac8c06a04bec0b551627dad28bc00d6516b0caacd1c7b345fcdb5211334e481565b348015610a21575f80fd5b506104b7610a3036600461490b565b611cfc565b348015610a40575f80fd5b50610762610a4f366004614635565b611d2b565b348015610a5f575f80fd5b506104b7610a6e36600461496c565b5f9182527f02dd7bc7dec4dceedda775e58dd541e08a116c6c53815c0bd028192f7b626800602090815260408084206001600160a01b0393909316845291905290205460ff1690565b348015610ac2575f80fd5b50610496610ad13660046144d7565b611d6b565b348015610ae1575f80fd5b506106455f81565b348015610af4575f80fd5b50610b08610b0336600461490b565b611dd0565b6040516104c39190614b7f565b348015610b20575f80fd5b506109d66040518060400160405280600581526020017f352e302e3000000000000000000000000000000000000000000000000000000081525081565b348015610b68575f80fd5b50610496611e0c565b348015610b7c575f80fd5b506104b7610b8b36600461490b565b611ea9565b348015610b9b575f80fd5b50610496610baa366004614635565b611eb3565b348015610bba575f80fd5b507f17fbd8e22750b6dc617ceae09f600bf2810b53ccc1f790a9a4f78a3f7169b900546001600160a01b0316610762565b348015610bf6575f80fd5b50610496610c053660046148f0565b611f02565b348015610c15575f80fd5b50610645610c2436600461490b565b611fd1565b348015610c34575f80fd5b50610496610c4336600461499a565b612008565b348015610c53575f80fd5b50610496610c62366004614bcb565b612029565b348015610c72575f80fd5b506106455f8051602061515983398151915281565b348015610c92575f80fd5b50610496610ca136600461490b565b612113565b348015610cb1575f80fd5b50610b08612161565b348015610cc5575f80fd5b50610496610cd436600461496c565b61219a565b348015610ce4575f80fd5b50610cf8610cf336600461490b565b6121dd565b6040516104c39190614c01565b348015610d10575f80fd5b506104b7610d1f3660046148f0565b612307565b348015610d2f575f80fd5b50610496610d3e366004614c13565b612346565b348015610d4e575f80fd5b50610496610d5d36600461490b565b6123f5565b348015610d6d575f80fd5b50610496610d7c3660046148f0565b61248a565b348015610d8c575f80fd5b50610b086124aa565b348015610da0575f80fd5b50610862610daf3660046148f0565b6001600160a01b03165f9081527fc7ae2c65fdae475b1bb2dd4079b252a7d893d0822b91a40dd25c763e583f770060209081526040918290208251808401909352546001600160801b038116808452600160801b90910467ffffffffffffffff16929091018290529091565b348015610e26575f80fd5b50610496610e353660046149e7565b6124e3565b348015610e45575f80fd5b50610496610e54366004614a11565b6125bc565b348015610e64575f80fd5b506106457f189ab7a9244df0848122154315af71fe140f3db0fe014031783b0946b8c9d2e381565b348015610e97575f80fd5b50610496610ea6366004614a51565b61260d565b348015610eb6575f80fd5b50610496610ec5366004614c9d565b61265c565b348015610ed5575f80fd5b50610496610ee4366004614cf6565b61291e565b7f10dac8c06a04bec0b551627dad28bc00d6516b0caacd1c7b345fcdb5211334e4610f138161296e565b6040517f095ea7b30000000000000000000000000000000000000000000000000000000081526001600160a01b0384811660048301526024820184905285169063095ea7b3906044015b5f604051808303815f87803b158015610f74575f80fd5b505af1158015610f86573d5f803e3d5ffd5b5050505050505050565b5f6001600160e01b03198216630a85bd0160e11b1480610fb45750610fb482612978565b92915050565b5f80516020615159833981519152610fd18161296e565b5f5b825181101561106e57610ffe838281518110610ff157610ff1614d30565b60200260200101516129b5565b61102083828151811061101357611013614d30565b6020026020010151612a5a565b82818151811061103257611032614d30565b60200260200101517f7acacfd576383587962277516962c289d19f807be443f4e303ab45ace24931ac60405160405180910390a2600101610fd3565b505050565b6060805f61107f6115c7565b80519091508086106110db57604080515f808252602082018181528284019093529091906110cf565b604080518082019091525f8152606060208201528152602001906001900390816110a85790505b50935093505050611239565b5f6110e68783614d58565b9050808611156110f4578095505b8567ffffffffffffffff81111561110d5761110d61453c565b604051908082528060200260200182016040528015611136578160200160208202803683370190505b5094508567ffffffffffffffff8111156111525761115261453c565b60405190808252806020026020018201604052801561119757816020015b604080518082019091525f8152606060208201528152602001906001900390816111705790505b5093505f5b8681101561123457836111af828a614d6b565b815181106111bf576111bf614d30565b60200260200101518682815181106111d9576111d9614d30565b602090810291909101015261120f846111f2838b614d6b565b8151811061120257611202614d30565b60200260200101516121dd565b85828151811061122157611221614d30565b602090810291909101015260010161119c565b505050505b9250929050565b5f805160206151598339815191526112578161296e565b6112618383612acf565b827f1b76230b39d2d0c1a2a77a90c170190d2280796ed56b280177256ce39df1a66483604051611295911515815260200190565b60405180910390a2505050565b60605f805160206151398339815191526112db7f39fc6f3ebcc11656ef8cd451e9e1f6a26855304f0f1787c5b86c527a1e2d8608612b07565b91505090565b7f10dac8c06a04bec0b551627dad28bc00d6516b0caacd1c7b345fcdb5211334e461130b8161296e565b6001600160a01b03831661133257604051633a954ecd60e21b815260040160405180910390fd5b6040517f42842e0e0000000000000000000000000000000000000000000000000000000081523060048201526001600160a01b038481166024830152604482018490528516906342842e0e90606401610f5d565b5f8051602061517983398151915261139d8161296e565b73__$12bd2f62b73a470fe0f6e02c33045f3191$__6321b87f3a6113d55f80516020615199833981519152546001600160a01b031690565b6040516001600160e01b031960e084901b1681526001600160a01b03909116600482015260248101869052604481018590526064015f6040518083038186803b158015611420575f80fd5b505af4158015611432573d5f803e3d5ffd5b50505050505050565b5f5f805160206151398339815191526114747f39fc6f3ebcc11656ef8cd451e9e1f6a26855304f0f1787c5b86c527a1e2d860384612b13565b9392505050565b5f805160206151798339815191526114928161296e565b73__$12bd2f62b73a470fe0f6e02c33045f3191$__6307e473166114ca5f80516020615199833981519152546001600160a01b031690565b6040516001600160e01b031960e084901b1681526001600160a01b0390911660048201526024810187905261ffff8087166044830152851660648201526084015f6040518083038186803b158015611520575f80fd5b505af4158015610f86573d5f803e3d5ffd5b5f8281527f02dd7bc7dec4dceedda775e58dd541e08a116c6c53815c0bd028192f7b626800602052604090206001015461156b8161296e565b6115758383612b34565b50505050565b6001600160a01b03811633146115bd576040517f6697b23200000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b61106e8282612b89565b60605f805160206151398339815191526112db81612b07565b6115e8612bd5565b6115f182612c8e565b6115fb8282612e81565b5050565b7fc6ae66da7afbd0becbe49ddf95a9256901d87f1bb5e3d43236030e5e4a585b6d6116298161296e565b6001600160a01b03831661165057604051633a954ecd60e21b815260040160405180910390fd5b61167a7fe9e633d75f5fd429fed1611b4a64c1a1d6af654653894bd204eb46dbd457621a84612b34565b506116925f8051602061517983398151915284612b34565b506040516001600160a01b038416907fdb3e11ba26e83d528bf96a2167061674c1ce7777c61376d852d172594a873994905f90a261106e6001600160a01b03841683612f69565b5f6116e261301c565b507f360894a13ba1a3210667c828492db98dca3e2076cc3735a920a3ca505d382bbc90565b5f8051602061515983398151915261171e8161296e565b6117288383613065565b827f1cd139430ed537ab9e8086952076cce01edd5ba6e30907af0ffe3709fd3139e6836040516112959190614b6d565b6117606130a9565b7fe562760eaa817d85ec1bf58364c4d65adb65d99d113c6785ef9aa66567076c9561178a8161296e565b6117938261310c565b506117bd60017f9b779b17422d0df92223018b32b4d1fa46e071723d6817e2486d003becc55f0055565b50565b60605f805160206151398339815191526117da838261338f565b5f838152600282016020908152604080832060010180548251818502810185019093528083529193909284015b828210156118af578382905f5260205f2001805461182490614d7e565b80601f016020809104026020016040519081016040528092919081815260200182805461185090614d7e565b801561189b5780601f106118725761010080835404028352916020019161189b565b820191905f5260205f20905b81548152906001019060200180831161187e57829003601f168201915b505050505081526020019060010190611807565b50505050915050919050565b5f805160206151598339815191526118d28161296e565b6115fb826133b9565b7fc6ae66da7afbd0becbe49ddf95a9256901d87f1bb5e3d43236030e5e4a585b6d6119058161296e565b61106e838361346e565b5f805160206151798339815191526119268161296e565b73__$12bd2f62b73a470fe0f6e02c33045f3191$__63fd13a43e61195e5f80516020615199833981519152546001600160a01b031690565b6040516001600160e01b031960e084901b1681526001600160a01b039091166004820152602481018890526044810187905261ffff80871660648301528516608482015260a4015f6040518083038186803b1580156119bb575f80fd5b505af41580156119cd573d5f803e3d5ffd5b505050505050505050565b5f805160206151798339815191526119ef8161296e565b73__$12bd2f62b73a470fe0f6e02c33045f3191$__63793dddac6114ca5f80516020615199833981519152546001600160a01b031690565b5f80516020615159833981519152611a3e8161296e565b6115fb82613551565b5f80516020615159833981519152611a5e8161296e565b611a67846129b5565b611a72848385613606565b60405184907f8f531e5ede07d5741fd086bb787ed399a64704eb757b87cc80cf6635b274e5b5905f90a250505050565b606080611aad6115c7565b9150815167ffffffffffffffff811115611ac957611ac961453c565b604051908082528060200260200182016040528015611b0e57816020015b604080518082019091525f815260606020820152815260200190600190039081611ae75790505b5090505f5b8251811015611b5657611b3183828151811061120257611202614d30565b828281518110611b4357611b43614d30565b6020908102919091010152600101611b13565b509091565b5f80516020615159833981519152611b728161296e565b5f5b825181101561106e57611b9f838281518110611b9257611b92614d30565b60200260200101516136a4565b828181518110611bb157611bb1614d30565b60200260200101517ff0dd3de472ddcd75ae2c17728a45801355fb6dd8615a7c53c15504b4279c09be60405160405180910390a2600101611b74565b60605f80516020615139833981519152611c277f39fc6f3ebcc11656ef8cd451e9e1f6a26855304f0f1787c5b86c527a1e2d860584612b13565b611c545760405163ba650b5f60e01b81526001600160a01b03841660048201526024015b60405180910390fd5b6001600160a01b0383165f90815260078201602052604090208054611c7890614d7e565b80601f0160208091040260200160405190810160405280929190818152602001828054611ca490614d7e565b8015611cef5780601f10611cc657610100808354040283529160200191611cef565b820191905f5260205f20905b815481529060010190602001808311611cd257829003601f168201915b5050505050915050919050565b5f5f80516020615139833981519152611d15838261338f565b5f92835260020160205250604090205460ff1690565b5f8281527fc1f6fe24621ce81ec5827caf0253cadb74709b061630e6b55e82371705932000602081905260408220611d639084613719565b949350505050565b7f10dac8c06a04bec0b551627dad28bc00d6516b0caacd1c7b345fcdb5211334e4611d958161296e565b6001600160a01b038316611dbc57604051633a954ecd60e21b815260040160405180910390fd5b6115756001600160a01b0385168484613724565b5f8181527fc1f6fe24621ce81ec5827caf0253cadb74709b061630e6b55e82371705932000602081905260409091206060919061147490612b07565b5f80516020615159833981519152611e238161296e565b5f611e2c6115c7565b90505f5b815181101561106e57611e5b828281518110611e4e57611e4e614d30565b60200260200101516137a4565b818181518110611e6d57611e6d614d30565b60200260200101517f94da5eeca10d4d6ee8455f99240c10b0c74b0cf5bf754afb81c81e2704b9c42760405160405180910390a2600101611e30565b5f610fb482613805565b5f80516020615179833981519152611eca8161296e565b73__$12bd2f62b73a470fe0f6e02c33045f3191$__63348e06dd6113d55f80516020615199833981519152546001600160a01b031690565b7fc6ae66da7afbd0becbe49ddf95a9256901d87f1bb5e3d43236030e5e4a585b6d611f2c8161296e565b611f567fe9e633d75f5fd429fed1611b4a64c1a1d6af654653894bd204eb46dbd457621a83612b89565b50611f6e5f8051602061517983398151915283612b89565b50611f997fe562760eaa817d85ec1bf58364c4d65adb65d99d113c6785ef9aa66567076c9583612b89565b506040516001600160a01b038316907fd124523a9cfa28c5dd01826c4fa56192ec7d56859943082e0ca46c3b9dc62913905f90a25050565b5f8181527fc1f6fe24621ce81ec5827caf0253cadb74709b061630e6b55e823717059320006020819052604082206114749061381e565b5f8051602061515983398151915261201f8161296e565b61106e8383613827565b6120316130a9565b5f805160206151798339815191526120488161296e565b73__$12bd2f62b73a470fe0f6e02c33045f3191$__637adf63b76120805f80516020615199833981519152546001600160a01b031690565b6040516001600160e01b031960e084901b1681526001600160a01b039182166004820152602481018890526044810187905290851660648201526084015f6040518083038186803b1580156120d3575f80fd5b505af41580156120e5573d5f803e3d5ffd5b505050505061106e60017f9b779b17422d0df92223018b32b4d1fa46e071723d6817e2486d003becc55f0055565b5f8051602061515983398151915261212a8161296e565b612133826137a4565b60405182907f94da5eeca10d4d6ee8455f99240c10b0c74b0cf5bf754afb81c81e2704b9c427905f90a25050565b60605f805160206151398339815191526112db7f39fc6f3ebcc11656ef8cd451e9e1f6a26855304f0f1787c5b86c527a1e2d8603612b07565b5f8281527f02dd7bc7dec4dceedda775e58dd541e08a116c6c53815c0bd028192f7b62680060205260409020600101546121d38161296e565b6115758383612b89565b604080518082019091525f8152606060208201525f80516020615139833981519152612209838261338f565b5f838152600282016020908152604080832081518083018352815460ff16151581526001820180548451818702810187019095528085529195929486810194939192919084015b828210156122f8578382905f5260205f2001805461226d90614d7e565b80601f016020809104026020016040519081016040528092919081815260200182805461229990614d7e565b80156122e45780601f106122bb576101008083540402835291602001916122e4565b820191905f5260205f20905b8154815290600101906020018083116122c757829003601f168201915b505050505081526020019060010190612250565b50505091525090949350505050565b6001600160a01b0381165f9081527f439a7b1b33d79c367c7c6755d8bb3d3ca77b7bca0d68cd209dcbe6cb4f5db4da602052604081205460ff16610fb4565b5f8051602061517983398151915261235d8161296e565b73__$12bd2f62b73a470fe0f6e02c33045f3191$__63e4c225696123955f80516020615199833981519152546001600160a01b031690565b8a8a8a8a8a8a8a6040518963ffffffff1660e01b81526004016123bf989796959493929190614db6565b5f6040518083038186803b1580156123d5575f80fd5b505af41580156123e7573d5f803e3d5ffd5b505050505050505050505050565b73__$12bd2f62b73a470fe0f6e02c33045f3191$__63c7bffa9661242d5f80516020615199833981519152546001600160a01b031690565b6040516001600160e01b031960e084901b1681526001600160a01b039091166004820152602481018490526044015f6040518083038186803b158015612471575f80fd5b505af4158015612483573d5f803e3d5ffd5b5050505050565b5f805160206151598339815191526124a18161296e565b6115fb82613901565b60605f805160206151398339815191526112db7f39fc6f3ebcc11656ef8cd451e9e1f6a26855304f0f1787c5b86c527a1e2d8605612b07565b6124eb6130a9565b7f10dac8c06a04bec0b551627dad28bc00d6516b0caacd1c7b345fcdb5211334e46125158161296e565b6001600160a01b03831661253c57604051633a954ecd60e21b815260040160405180910390fd5b61254f6001600160a01b03841683612f69565b826001600160a01b03167f884edad9ce6fa2440d8a54cc123490eb96d2768479d49ff9c7366125a94243648360405161258a91815260200190565b60405180910390a2506115fb60017f9b779b17422d0df92223018b32b4d1fa46e071723d6817e2486d003becc55f0055565b5f805160206151598339815191526125d38161296e565b6125dd83836139bf565b827ffc8d82c9e7e7938446da05458183efa5916c443a2bab87f97f94a8d47742b014836040516112959190614b6d565b5f805160206151798339815191526126248161296e565b73__$12bd2f62b73a470fe0f6e02c33045f3191$__63b54e72d861195e5f80516020615199833981519152546001600160a01b031690565b7ff0c57e16840df040f15088dc2f81fe391c3923bec73e23a9662efc9c229c6a00805468010000000000000000810460ff16159067ffffffffffffffff165f811580156126a65750825b90505f8267ffffffffffffffff1660011480156126c25750303b155b9050811580156126d0575080155b15612707576040517ff92ee8a900000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b845467ffffffffffffffff19166001178555831561273b57845468ff00000000000000001916680100000000000000001785555b6001600160a01b038916158061275857506001600160a01b038816155b8061276a57506001600160a01b038716155b8061277c57506001600160a01b038616155b156127b3576040517fd92e233d00000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b6127bb613b1f565b6127c3613b1f565b6127cb613b27565b6127d3613b1f565b6127dd5f88612b34565b506127f55f8051602061515983398151915288612b34565b506128207fc6ae66da7afbd0becbe49ddf95a9256901d87f1bb5e3d43236030e5e4a585b6d88612b34565b5061284b7f189ab7a9244df0848122154315af71fe140f3db0fe014031783b0946b8c9d2e387612b34565b507f17fbd8e22750b6dc617ceae09f600bf2810b53ccc1f790a9a4f78a3f7169b90080546001600160a01b038b811673ffffffffffffffffffffffffffffffffffffffff199283161783555f805160206151998339815191528054918c1691909216179055662386f26fc10000620151806128c68282613b37565b50505083156119cd57845468ff000000000000000019168555604051600181527fc7f505b2f371ae2175ee4913f4499e1f2633a7b5936321eed1cdaeb6115181d29060200160405180910390a1505050505050505050565b5f805160206151598339815191526129358161296e565b61293f8383613bd0565b60405183907fa616bfc5bb0e46c6cad727e1b55e3685067e1296d962a7f37017874a27aa0098905f90a2505050565b6117bd8133613c0e565b5f6001600160e01b031982167f5a05180f000000000000000000000000000000000000000000000000000000001480610fb45750610fb482613c9a565b7f17fbd8e22750b6dc617ceae09f600bf2810b53ccc1f790a9a4f78a3f7169b900546040517f5a81a626000000000000000000000000000000000000000000000000000000008152600481018390526001600160a01b0390911690635a81a626906024015f60405180830381865afa158015612a33573d5f803e3d5ffd5b505050506040513d5f823e601f3d908101601f191682016040526115fb9190810190614e11565b5f805160206151398339815191525f612a937f39fc6f3ebcc11656ef8cd451e9e1f6a26855304f0f1787c5b86c527a1e2d860884613d00565b90508061106e576040517f1a1e056900000000000000000000000000000000000000000000000000000000815260048101849052602401611c4b565b5f80516020615139833981519152612ae7838261338f565b5f9283526002016020526040909120805460ff1916911515919091179055565b60605f61147483613d0b565b6001600160a01b0381165f9081526001830160205260408120541515611474565b5f7fc1f6fe24621ce81ec5827caf0253cadb74709b061630e6b55e8237170593200081612b618585613d64565b90508015611d63575f858152602083905260409020612b809085613e30565b50949350505050565b5f7fc1f6fe24621ce81ec5827caf0253cadb74709b061630e6b55e8237170593200081612bb68585613e44565b90508015611d63575f858152602083905260409020612b809085613ee8565b306001600160a01b037f0000000000000000000000000000000000000000000000000000000000000000161480612c6e57507f00000000000000000000000000000000000000000000000000000000000000006001600160a01b0316612c627f360894a13ba1a3210667c828492db98dca3e2076cc3735a920a3ca505d382bbc546001600160a01b031690565b6001600160a01b031614155b15612c8c5760405163703e46dd60e11b815260040160405180910390fd5b565b7f189ab7a9244df0848122154315af71fe140f3db0fe014031783b0946b8c9d2e3612cb88161296e565b5f612cea7f17fbd8e22750b6dc617ceae09f600bf2810b53ccc1f790a9a4f78a3f7169b900546001600160a01b031690565b6001600160a01b0316639d825bc56040518163ffffffff1660e01b8152600401602060405180830381865afa158015612d25573d5f803e3d5ffd5b505050506040513d601f19601f82011682018060405250810190612d499190614e7a565b90505f612d7d7f360894a13ba1a3210667c828492db98dca3e2076cc3735a920a3ca505d382bbc546001600160a01b031690565b9050836001600160a01b0316816001600160a01b031603612ddd576040517ffe51a0290000000000000000000000000000000000000000000000000000000081526001600160a01b03808316600483015285166024820152604401611c4b565b816001600160a01b0316846001600160a01b031614612e3b576040517f08811c0c0000000000000000000000000000000000000000000000000000000081526001600160a01b03808416600483015285166024820152604401611c4b565b836001600160a01b0316816001600160a01b03167f897c7778b6095182ea48ee84760832efeae452e4c42d863ea35b271a3aaae75960405160405180910390a350505050565b816001600160a01b03166352d1902d6040518163ffffffff1660e01b8152600401602060405180830381865afa925050508015612edb575060408051601f3d908101601f19168201909252612ed891810190614e95565b60015b612f0357604051634c9c8ce360e01b81526001600160a01b0383166004820152602401611c4b565b7f360894a13ba1a3210667c828492db98dca3e2076cc3735a920a3ca505d382bbc8114612f5f576040517faa1d49a400000000000000000000000000000000000000000000000000000000815260048101829052602401611c4b565b61106e8383613efc565b80471015612fac576040517fcf47918100000000000000000000000000000000000000000000000000000000815247600482015260248101829052604401611c4b565b5f826001600160a01b0316826040515f6040518083038185875af1925050503d805f8114612ff5576040519150601f19603f3d011682016040523d82523d5f602084013e612ffa565b606091505b505090508061106e5760405163d6bda27560e01b815260040160405180910390fd5b306001600160a01b037f00000000000000000000000000000000000000000000000000000000000000001614612c8c5760405163703e46dd60e11b815260040160405180910390fd5b5f8051602061513983398151915261307d838261338f565b5f8381526002820160209081526040822060019081018054918201815583529120016115758382614ef0565b7f9b779b17422d0df92223018b32b4d1fa46e071723d6817e2486d003becc55f00805460011901613106576040517f3ee5aeb500000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b60029055565b7fc7ae2c65fdae475b1bb2dd4079b252a7d893d0822b91a40dd25c763e583f7701547fc7ae2c65fdae475b1bb2dd4079b252a7d893d0822b91a40dd25c763e583f7700906001600160801b03168083111561319d576040517fb945e2d80000000000000000000000000000000000000000000000000000000081526004810182905260248101849052604401611c4b565b335f90815260208381526040918290208251808401909352546001600160801b038116835267ffffffffffffffff600160801b918290048116928401839052600186015442936131f39390910490911690614d6b565b811115613230575f825260018401546132209082908590600160801b900467ffffffffffffffff16613f51565b67ffffffffffffffff1660208301525b815183906132489087906001600160801b0316614d6b565b111561328a576040517fd54b18870000000000000000000000000000000000000000000000000000000081526004810184905260248101869052604401611c4b565b81516132c1906132a49087906001600160801b0316614d6b565b60018601548590600160801b900467ffffffffffffffff16613f8d565b6001600160801b039081168352335f81815260208781526040909120855181549287015167ffffffffffffffff16600160801b027fffffffffffffffff0000000000000000000000000000000000000000000000009093169416939093171790915561332d9086612f69565b60405185815233907fb9ec638398bbdcd0844ca414d8ce760939fa88b9258b9764b3fc6c12ea2605c29060200160405180910390a25050505050565b60017f9b779b17422d0df92223018b32b4d1fa46e071723d6817e2486d003becc55f0055565b6133998183613fc0565b6115fb57604051631e96f6ed60e21b815260048101839052602401611c4b565b5f805160206151398339815191525f6133f27f39fc6f3ebcc11656ef8cd451e9e1f6a26855304f0f1787c5b86c527a1e2d860384613e30565b905080613436576040517f50e5f7f20000000000000000000000000000000000000000000000000000000081526001600160a01b0384166004820152602401611c4b565b6040516001600160a01b038416907fa317c10673baf4f03b3c1041bd5ddbb537d0333a86fec3607c75f9dbb630f48f905f90a2505050565b7fc7ae2c65fdae475b1bb2dd4079b252a7d893d0822b91a40dd25c763e583f770061349a838084613f8d565b6001820180546fffffffffffffffffffffffffffffffff19166001600160801b03929092169190911790556134d0828481613f51565b60018201805467ffffffffffffffff92909216600160801b027fffffffffffffffff0000000000000000ffffffffffffffffffffffffffffffff90921691909117905560408051848152602081018490527f8d4925b196ae6b935035a27ed36c6bd9c7a8fbddc7a3f55f493aa8e230be373e910160405180910390a1505050565b5f805160206151398339815191525f61358a7f39fc6f3ebcc11656ef8cd451e9e1f6a26855304f0f1787c5b86c527a1e2d860384613ee8565b9050806135ce576040517f54cb99c40000000000000000000000000000000000000000000000000000000081526001600160a01b0384166004820152602401611c4b565b6040516001600160a01b038416907f85a3e72f8dd6db3794f93109c3c5f5b79d6112f6979431c45f98b26134b42af2905f90a2505050565b5f805160206151398339815191525f61361f8286613d00565b90508061365b576040517f010f1dd800000000000000000000000000000000000000000000000000000000815260048101869052602401611c4b565b604080518082018252841515815260208082018781525f898152600287018352939093208251815460ff191690151517815592518051929392610f8692600185019201906143ed565b5f805160206151398339815191525f6136dd7f39fc6f3ebcc11656ef8cd451e9e1f6a26855304f0f1787c5b86c527a1e2d860884613fd7565b90508061106e576040517ffd5cb3e200000000000000000000000000000000000000000000000000000000815260048101849052602401611c4b565b5f6114748383613fe2565b604080516001600160a01b038416602482015260448082018490528251808303909101815260649091019091526020810180517bffffffffffffffffffffffffffffffffffffffffffffffffffffffff167fa9059cbb0000000000000000000000000000000000000000000000000000000017905261106e908490614008565b5f805160206151398339815191525f6137bd8284613fd7565b9050806137e057604051631e96f6ed60e21b815260048101849052602401611c4b565b5f8381526002830160205260408120805460ff19168155906124836001830182614441565b5f5f805160206151398339815191526114748184613fc0565b5f610fb4825490565b5f805160206151398339815191525f6138607f39fc6f3ebcc11656ef8cd451e9e1f6a26855304f0f1787c5b86c527a1e2d860585613e30565b9050806138a4576040517fd3083f180000000000000000000000000000000000000000000000000000000081526001600160a01b0385166004820152602401611c4b565b6001600160a01b0384165f90815260078301602052604090206138c78482614ef0565b506040516001600160a01b038516907f928ec246afda323bc23c2815ca3f516e9fc6a7b7179772235c221e132545fe82905f90a250505050565b5f805160206151398339815191525f61393a7f39fc6f3ebcc11656ef8cd451e9e1f6a26855304f0f1787c5b86c527a1e2d860584613ee8565b9050806139655760405163ba650b5f60e01b81526001600160a01b0384166004820152602401611c4b565b6001600160a01b0383165f90815260078301602052604081206139879161445c565b6040516001600160a01b038416907fc5a9b8041ef11732e7dd7043167d8c22db5c7ea99dcd38dce401effacf8a29bf905f90a2505050565b5f805160206151398339815191526139d7838261338f565b5f8381526002820160205260408120600101805490915b81811015613ae85784604051602001613a079190614fb0565b60405160208183030381529060405280519060200120838281548110613a2f57613a2f614d30565b905f5260205f2001604051602001613a479190614fcb565b6040516020818303038152906040528051906020012003613ae05782613a6e600184614d58565b81548110613a7e57613a7e614d30565b905f5260205f2001838281548110613a9857613a98614d30565b905f5260205f20019081613aac919061503d565b5082805480613abd57613abd61510c565b600190038181905f5260205f20015f613ad6919061445c565b9055505050505050565b6001016139ee565b5084846040517fe879f039000000000000000000000000000000000000000000000000000000008152600401611c4b929190615120565b612c8c61408d565b613b2f61408d565b612c8c6140f4565b613b3f61408d565b7fc7ae2c65fdae475b1bb2dd4079b252a7d893d0822b91a40dd25c763e583f7700613b6b838084613f8d565b6001820180546fffffffffffffffffffffffffffffffff19166001600160801b0392909216919091179055613ba1828481613f51565b8160010160106101000a81548167ffffffffffffffff021916908367ffffffffffffffff160217905550505050565b5f80516020615139833981519152613be8838261338f565b5f83815260028201602090815260409091208351611575926001909201918501906143ed565b5f8281527f02dd7bc7dec4dceedda775e58dd541e08a116c6c53815c0bd028192f7b626800602090815260408083206001600160a01b038516845290915290205460ff166115fb576040517fe2517d3f0000000000000000000000000000000000000000000000000000000081526001600160a01b038216600482015260248101839052604401611c4b565b5f6001600160e01b031982167f7965db0b000000000000000000000000000000000000000000000000000000001480610fb457507f01ffc9a7000000000000000000000000000000000000000000000000000000006001600160e01b0319831614610fb4565b5f61147483836140fc565b6060815f01805480602002602001604051908101604052809291908181526020018280548015613d5857602002820191905f5260205f20905b815481526020019060010190808311613d44575b50505050509050919050565b5f8281527f02dd7bc7dec4dceedda775e58dd541e08a116c6c53815c0bd028192f7b626800602081815260408084206001600160a01b038616855290915282205460ff16613e27575f848152602082815260408083206001600160a01b03871684529091529020805460ff19166001179055613ddd3390565b6001600160a01b0316836001600160a01b0316857f2f8788117e7eff1d82e926ec794901d17c78024a50270940304540a733656f0d60405160405180910390a46001915050610fb4565b5f915050610fb4565b5f611474836001600160a01b0384166140fc565b5f8281527f02dd7bc7dec4dceedda775e58dd541e08a116c6c53815c0bd028192f7b626800602081815260408084206001600160a01b038616855290915282205460ff1615613e27575f848152602082815260408083206001600160a01b0387168085529252808320805460ff1916905551339287917ff6391f5c32d9c69d2a47ea670b442974b53935d1edc7fd64eb21e047a839171b9190a46001915050610fb4565b5f611474836001600160a01b038416614148565b613f0582614222565b6040516001600160a01b038316907fbc7cd75a20ee27fd9adebab32041f755214dbc6bffa90cc0225b39da2e5c2d3b905f90a2805115613f495761106e82826142a5565b6115fb614317565b5f67ffffffffffffffff841115613f855760405163d450716560e01b81526004810184905260248101839052604401611c4b565b509192915050565b5f6001600160801b03841115613f855760405163d450716560e01b81526004810184905260248101839052604401611c4b565b5f8181526001830160205260408120541515611474565b5f6114748383614148565b5f825f018281548110613ff757613ff7614d30565b905f5260205f200154905092915050565b5f8060205f8451602086015f885af180614027576040513d5f823e3d81fd5b50505f513d9150811561403e57806001141561404b565b6001600160a01b0384163b155b15611575576040517f5274afe70000000000000000000000000000000000000000000000000000000081526001600160a01b0385166004820152602401611c4b565b7ff0c57e16840df040f15088dc2f81fe391c3923bec73e23a9662efc9c229c6a005468010000000000000000900460ff16612c8c576040517fd7e6bcf800000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b61336961408d565b5f81815260018301602052604081205461414157508154600181810184555f848152602080822090930184905584548482528286019093526040902091909155610fb4565b505f610fb4565b5f8181526001830160205260408120548015613e27575f61416a600183614d58565b85549091505f9061417d90600190614d58565b90508082146141dc575f865f01828154811061419b5761419b614d30565b905f5260205f200154905080875f0184815481106141bb576141bb614d30565b5f918252602080832090910192909255918252600188019052604090208390555b85548690806141ed576141ed61510c565b600190038181905f5260205f20015f90559055856001015f8681526020019081526020015f205f905560019350505050610fb4565b806001600160a01b03163b5f0361425757604051634c9c8ce360e01b81526001600160a01b0382166004820152602401611c4b565b7f360894a13ba1a3210667c828492db98dca3e2076cc3735a920a3ca505d382bbc805473ffffffffffffffffffffffffffffffffffffffff19166001600160a01b0392909216919091179055565b60605f80846001600160a01b0316846040516142c19190614fb0565b5f60405180830381855af49150503d805f81146142f9576040519150601f19603f3d011682016040523d82523d5f602084013e6142fe565b606091505b509150915061430e85838361434f565b95945050505050565b3415612c8c576040517fb398979f00000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b6060826143645761435f826143c4565b611474565b815115801561437b57506001600160a01b0384163b155b156143bd576040517f9996b3150000000000000000000000000000000000000000000000000000000081526001600160a01b0385166004820152602401611c4b565b5080611474565b8051156143d45780518082602001fd5b60405163d6bda27560e01b815260040160405180910390fd5b828054828255905f5260205f20908101928215614431579160200282015b8281111561443157825182906144219082614ef0565b509160200191906001019061440b565b5061443d929150614493565b5090565b5080545f8255905f5260205f20908101906117bd9190614493565b50805461446890614d7e565b5f825580601f10614477575050565b601f0160209004905f5260205f20908101906117bd91906144af565b8082111561443d575f6144a6828261445c565b50600101614493565b5b8082111561443d575f81556001016144b0565b6001600160a01b03811681146117bd575f80fd5b5f805f606084860312156144e9575f80fd5b83356144f4816144c3565b92506020840135614504816144c3565b929592945050506040919091013590565b5f60208284031215614525575f80fd5b81356001600160e01b031981168114611474575f80fd5b634e487b7160e01b5f52604160045260245ffd5b604051601f8201601f1916810167ffffffffffffffff811182821017156145795761457961453c565b604052919050565b5f67ffffffffffffffff82111561459a5761459a61453c565b5060051b60200190565b5f60208083850312156145b5575f80fd5b823567ffffffffffffffff8111156145cb575f80fd5b8301601f810185136145db575f80fd5b80356145ee6145e982614581565b614550565b81815260059190911b8201830190838101908783111561460c575f80fd5b928401925b8284101561462a57833582529284019290840190614611565b979650505050505050565b5f8060408385031215614646575f80fd5b50508035926020909101359150565b5f815180845260208085019450602084015f5b8381101561468457815187529582019590820190600101614668565b509495945050505050565b5f5b838110156146a9578181015183820152602001614691565b50505f910152565b5f81518084526146c881602086016020860161468f565b601f01601f19169290920160200192915050565b5f8282518085526020808601955060208260051b840101602086015f5b8481101561472757601f198684030189526147158383516146b1565b988401989250908301906001016146f9565b5090979650505050505050565b8051151582525f602082015160406020850152611d6360408501826146dc565b604081525f6147666040830185614655565b6020838203818501528185518084528284019150828160051b8501018388015f5b838110156147b557601f198784030185526147a3838351614734565b94860194925090850190600101614787565b50909998505050505050505050565b803580151581146147d3575f80fd5b919050565b5f80604083850312156147e9575f80fd5b823591506147f9602084016147c4565b90509250929050565b602081525f6114746020830184614655565b5f67ffffffffffffffff82111561482d5761482d61453c565b50601f01601f191660200190565b5f82601f83011261484a575f80fd5b81356148586145e982614814565b81815284602083860101111561486c575f80fd5b816020850160208301375f918101602001919091529392505050565b5f805f806080858703121561489b575f80fd5b84356148a6816144c3565b935060208501356148b6816144c3565b925060408501359150606085013567ffffffffffffffff8111156148d8575f80fd5b6148e48782880161483b565b91505092959194509250565b5f60208284031215614900575f80fd5b8135611474816144c3565b5f6020828403121561491b575f80fd5b5035919050565b803561ffff811681146147d3575f80fd5b5f805f60608486031215614945575f80fd5b8335925061495560208501614922565b915061496360408501614922565b90509250925092565b5f806040838503121561497d575f80fd5b82359150602083013561498f816144c3565b809150509250929050565b5f80604083850312156149ab575f80fd5b82356149b6816144c3565b9150602083013567ffffffffffffffff8111156149d1575f80fd5b6149dd8582860161483b565b9150509250929050565b5f80604083850312156149f8575f80fd5b8235614a03816144c3565b946020939093013593505050565b5f8060408385031215614a22575f80fd5b82359150602083013567ffffffffffffffff8111156149d1575f80fd5b602081525f61147460208301846146dc565b5f805f8060808587031215614a64575f80fd5b8435935060208501359250614a7b60408601614922565b9150614a8960608601614922565b905092959194509250565b5f82601f830112614aa3575f80fd5b81356020614ab36145e983614581565b82815260059290921b84018101918181019086841115614ad1575f80fd5b8286015b84811015614b0f57803567ffffffffffffffff811115614af3575f80fd5b614b018986838b010161483b565b845250918301918301614ad5565b509695505050505050565b5f805f60608486031215614b2c575f80fd5b83359250614b3c602085016147c4565b9150604084013567ffffffffffffffff811115614b57575f80fd5b614b6386828701614a94565b9150509250925092565b602081525f61147460208301846146b1565b602080825282518282018190525f9190848201906040850190845b81811015614bbf5783516001600160a01b031683529284019291840191600101614b9a565b50909695505050505050565b5f805f60608486031215614bdd575f80fd5b83359250602084013591506040840135614bf6816144c3565b809150509250925092565b602081525f6114746020830184614734565b5f805f805f805f60e0888a031215614c29575f80fd5b8735614c34816144c3565b9650602088013567ffffffffffffffff811115614c4f575f80fd5b614c5b8a828b0161483b565b96505060408801359450606088013593506080880135614c7a816144c3565b925060a08801359150614c8f60c089016147c4565b905092959891949750929550565b5f805f8060808587031215614cb0575f80fd5b8435614cbb816144c3565b93506020850135614ccb816144c3565b92506040850135614cdb816144c3565b91506060850135614ceb816144c3565b939692955090935050565b5f8060408385031215614d07575f80fd5b82359150602083013567ffffffffffffffff811115614d24575f80fd5b6149dd85828601614a94565b634e487b7160e01b5f52603260045260245ffd5b634e487b7160e01b5f52601160045260245ffd5b81810381811115610fb457610fb4614d44565b80820180821115610fb457610fb4614d44565b600181811c90821680614d9257607f821691505b602082108103614db057634e487b7160e01b5f52602260045260245ffd5b50919050565b5f6101006001600160a01b03808c168452808b166020850152816040850152614de18285018b6146b1565b6060850199909952608084019790975250509290931660a083015260c082015290151560e0909101529392505050565b5f60208284031215614e21575f80fd5b815167ffffffffffffffff811115614e37575f80fd5b8201601f81018413614e47575f80fd5b8051614e556145e982614814565b818152856020838501011115614e69575f80fd5b61430e82602083016020860161468f565b5f60208284031215614e8a575f80fd5b8151611474816144c3565b5f60208284031215614ea5575f80fd5b5051919050565b601f82111561106e57805f5260205f20601f840160051c81016020851015614ed15750805b601f840160051c820191505b81811015612483575f8155600101614edd565b815167ffffffffffffffff811115614f0a57614f0a61453c565b614f1e81614f188454614d7e565b84614eac565b602080601f831160018114614f51575f8415614f3a5750858301515b5f19600386901b1c1916600185901b178555614fa8565b5f85815260208120601f198616915b82811015614f7f57888601518255948401946001909101908401614f60565b5085821015614f9c57878501515f19600388901b60f8161c191681555b505060018460011b0185555b505050505050565b5f8251614fc181846020870161468f565b9190910192915050565b5f808354614fd881614d7e565b60018281168015614ff0576001811461500557615031565b60ff1984168752821515830287019450615031565b875f526020805f205f5b858110156150285781548a82015290840190820161500f565b50505082870194505b50929695505050505050565b818103615048575050565b6150528254614d7e565b67ffffffffffffffff81111561506a5761506a61453c565b61507881614f188454614d7e565b5f601f8211600181146150a9575f83156150925750848201545b5f19600385901b1c1916600184901b178455612483565b5f8581526020808220868352908220601f198616925b838110156150df57828601548255600195860195909101906020016150bf565b50858310156150fc57818501545f19600388901b60f8161c191681555b5050505050600190811b01905550565b634e487b7160e01b5f52603160045260245ffd5b828152604060208201525f611d6360408301846146b156fe39fc6f3ebcc11656ef8cd451e9e1f6a26855304f0f1787c5b86c527a1e2d86009a95e87c5af084bf5db8491c3a6515da9dd6da39b24b0eb0af08d7b9cd808d913acdf00ba9ef08b5f2c22768276611b9af078bf6c24fa36b34ec5e9f2eb061fa17fbd8e22750b6dc617ceae09f600bf2810b53ccc1f790a9a4f78a3f7169b901a2646970667358221220acb3eaad38c4bac260b934b98afd8d6addb9a858fdc89eed00736febf6b544c264736f6c63430008180033",
}

// TtmaccountABI is the input ABI used to generate the binding from.
// Deprecated: Use TtmaccountMetaData.ABI instead.
var TtmaccountABI = TtmaccountMetaData.ABI

// TtmaccountBin is the compiled bytecode used for deploying new contracts.
// Deprecated: Use TtmaccountMetaData.Bin instead.
var TtmaccountBin = TtmaccountMetaData.Bin

// DeployTtmaccount deploys a new Ethereum contract, binding an instance of Ttmaccount to it.
func DeployTtmaccount(auth *bind.TransactOpts, backend bind.ContractBackend) (common.Address, *types.Transaction, *Ttmaccount, error) {
	parsed, err := TtmaccountMetaData.GetAbi()
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	if parsed == nil {
		return common.Address{}, nil, nil, errors.New("GetABI returned nil")
	}

	address, tx, contract, err := bind.DeployContract(auth, *parsed, common.FromHex(TtmaccountBin), backend)
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	return address, tx, &Ttmaccount{TtmaccountCaller: TtmaccountCaller{contract: contract}, TtmaccountTransactor: TtmaccountTransactor{contract: contract}, TtmaccountFilterer: TtmaccountFilterer{contract: contract}}, nil
}

// Ttmaccount is an auto generated Go binding around an Ethereum contract.
type Ttmaccount struct {
	TtmaccountCaller     // Read-only binding to the contract
	TtmaccountTransactor // Write-only binding to the contract
	TtmaccountFilterer   // Log filterer for contract events
}

// TtmaccountCaller is an auto generated read-only Go binding around an Ethereum contract.
type TtmaccountCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// TtmaccountTransactor is an auto generated write-only Go binding around an Ethereum contract.
type TtmaccountTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// TtmaccountFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type TtmaccountFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// TtmaccountSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type TtmaccountSession struct {
	Contract     *Ttmaccount       // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// TtmaccountCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type TtmaccountCallerSession struct {
	Contract *TtmaccountCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts     // Call options to use throughout this session
}

// TtmaccountTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type TtmaccountTransactorSession struct {
	Contract     *TtmaccountTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts     // Transaction auth options to use throughout this session
}

// TtmaccountRaw is an auto generated low-level Go binding around an Ethereum contract.
type TtmaccountRaw struct {
	Contract *Ttmaccount // Generic contract binding to access the raw methods on
}

// TtmaccountCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type TtmaccountCallerRaw struct {
	Contract *TtmaccountCaller // Generic read-only contract binding to access the raw methods on
}

// TtmaccountTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type TtmaccountTransactorRaw struct {
	Contract *TtmaccountTransactor // Generic write-only contract binding to access the raw methods on
}

// NewTtmaccount creates a new instance of Ttmaccount, bound to a specific deployed contract.
func NewTtmaccount(address common.Address, backend bind.ContractBackend) (*Ttmaccount, error) {
	contract, err := bindTtmaccount(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &Ttmaccount{TtmaccountCaller: TtmaccountCaller{contract: contract}, TtmaccountTransactor: TtmaccountTransactor{contract: contract}, TtmaccountFilterer: TtmaccountFilterer{contract: contract}}, nil
}

// NewTtmaccountCaller creates a new read-only instance of Ttmaccount, bound to a specific deployed contract.
func NewTtmaccountCaller(address common.Address, caller bind.ContractCaller) (*TtmaccountCaller, error) {
	contract, err := bindTtmaccount(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &TtmaccountCaller{contract: contract}, nil
}

// NewTtmaccountTransactor creates a new write-only instance of Ttmaccount, bound to a specific deployed contract.
func NewTtmaccountTransactor(address common.Address, transactor bind.ContractTransactor) (*TtmaccountTransactor, error) {
	contract, err := bindTtmaccount(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &TtmaccountTransactor{contract: contract}, nil
}

// NewTtmaccountFilterer creates a new log filterer instance of Ttmaccount, bound to a specific deployed contract.
func NewTtmaccountFilterer(address common.Address, filterer bind.ContractFilterer) (*TtmaccountFilterer, error) {
	contract, err := bindTtmaccount(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &TtmaccountFilterer{contract: contract}, nil
}

// bindTtmaccount binds a generic wrapper to an already deployed contract.
func bindTtmaccount(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := TtmaccountMetaData.GetAbi()
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, *parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Ttmaccount *TtmaccountRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _Ttmaccount.Contract.TtmaccountCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Ttmaccount *TtmaccountRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Ttmaccount.Contract.TtmaccountTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Ttmaccount *TtmaccountRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Ttmaccount.Contract.TtmaccountTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Ttmaccount *TtmaccountCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _Ttmaccount.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Ttmaccount *TtmaccountTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Ttmaccount.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Ttmaccount *TtmaccountTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Ttmaccount.Contract.contract.Transact(opts, method, params...)
}

// BOOKINGOPERATORROLE is a free data retrieval call binding the contract method 0x852b3ccb.
//
// Solidity: function BOOKING_OPERATOR_ROLE() view returns(bytes32)
func (_Ttmaccount *TtmaccountCaller) BOOKINGOPERATORROLE(opts *bind.CallOpts) ([32]byte, error) {
	var out []interface{}
	err := _Ttmaccount.contract.Call(opts, &out, "BOOKING_OPERATOR_ROLE")

	if err != nil {
		return *new([32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)

	return out0, err

}

// BOOKINGOPERATORROLE is a free data retrieval call binding the contract method 0x852b3ccb.
//
// Solidity: function BOOKING_OPERATOR_ROLE() view returns(bytes32)
func (_Ttmaccount *TtmaccountSession) BOOKINGOPERATORROLE() ([32]byte, error) {
	return _Ttmaccount.Contract.BOOKINGOPERATORROLE(&_Ttmaccount.CallOpts)
}

// BOOKINGOPERATORROLE is a free data retrieval call binding the contract method 0x852b3ccb.
//
// Solidity: function BOOKING_OPERATOR_ROLE() view returns(bytes32)
func (_Ttmaccount *TtmaccountCallerSession) BOOKINGOPERATORROLE() ([32]byte, error) {
	return _Ttmaccount.Contract.BOOKINGOPERATORROLE(&_Ttmaccount.CallOpts)
}

// BOTADMINROLE is a free data retrieval call binding the contract method 0x33746274.
//
// Solidity: function BOT_ADMIN_ROLE() view returns(bytes32)
func (_Ttmaccount *TtmaccountCaller) BOTADMINROLE(opts *bind.CallOpts) ([32]byte, error) {
	var out []interface{}
	err := _Ttmaccount.contract.Call(opts, &out, "BOT_ADMIN_ROLE")

	if err != nil {
		return *new([32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)

	return out0, err

}

// BOTADMINROLE is a free data retrieval call binding the contract method 0x33746274.
//
// Solidity: function BOT_ADMIN_ROLE() view returns(bytes32)
func (_Ttmaccount *TtmaccountSession) BOTADMINROLE() ([32]byte, error) {
	return _Ttmaccount.Contract.BOTADMINROLE(&_Ttmaccount.CallOpts)
}

// BOTADMINROLE is a free data retrieval call binding the contract method 0x33746274.
//
// Solidity: function BOT_ADMIN_ROLE() view returns(bytes32)
func (_Ttmaccount *TtmaccountCallerSession) BOTADMINROLE() ([32]byte, error) {
	return _Ttmaccount.Contract.BOTADMINROLE(&_Ttmaccount.CallOpts)
}

// DEFAULTADMINROLE is a free data retrieval call binding the contract method 0xa217fddf.
//
// Solidity: function DEFAULT_ADMIN_ROLE() view returns(bytes32)
func (_Ttmaccount *TtmaccountCaller) DEFAULTADMINROLE(opts *bind.CallOpts) ([32]byte, error) {
	var out []interface{}
	err := _Ttmaccount.contract.Call(opts, &out, "DEFAULT_ADMIN_ROLE")

	if err != nil {
		return *new([32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)

	return out0, err

}

// DEFAULTADMINROLE is a free data retrieval call binding the contract method 0xa217fddf.
//
// Solidity: function DEFAULT_ADMIN_ROLE() view returns(bytes32)
func (_Ttmaccount *TtmaccountSession) DEFAULTADMINROLE() ([32]byte, error) {
	return _Ttmaccount.Contract.DEFAULTADMINROLE(&_Ttmaccount.CallOpts)
}

// DEFAULTADMINROLE is a free data retrieval call binding the contract method 0xa217fddf.
//
// Solidity: function DEFAULT_ADMIN_ROLE() view returns(bytes32)
func (_Ttmaccount *TtmaccountCallerSession) DEFAULTADMINROLE() ([32]byte, error) {
	return _Ttmaccount.Contract.DEFAULTADMINROLE(&_Ttmaccount.CallOpts)
}

// GASWITHDRAWERROLE is a free data retrieval call binding the contract method 0x383aba87.
//
// Solidity: function GAS_WITHDRAWER_ROLE() view returns(bytes32)
func (_Ttmaccount *TtmaccountCaller) GASWITHDRAWERROLE(opts *bind.CallOpts) ([32]byte, error) {
	var out []interface{}
	err := _Ttmaccount.contract.Call(opts, &out, "GAS_WITHDRAWER_ROLE")

	if err != nil {
		return *new([32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)

	return out0, err

}

// GASWITHDRAWERROLE is a free data retrieval call binding the contract method 0x383aba87.
//
// Solidity: function GAS_WITHDRAWER_ROLE() view returns(bytes32)
func (_Ttmaccount *TtmaccountSession) GASWITHDRAWERROLE() ([32]byte, error) {
	return _Ttmaccount.Contract.GASWITHDRAWERROLE(&_Ttmaccount.CallOpts)
}

// GASWITHDRAWERROLE is a free data retrieval call binding the contract method 0x383aba87.
//
// Solidity: function GAS_WITHDRAWER_ROLE() view returns(bytes32)
func (_Ttmaccount *TtmaccountCallerSession) GASWITHDRAWERROLE() ([32]byte, error) {
	return _Ttmaccount.Contract.GASWITHDRAWERROLE(&_Ttmaccount.CallOpts)
}

// MESSENGERBOTROLE is a free data retrieval call binding the contract method 0x72afa328.
//
// Solidity: function MESSENGER_BOT_ROLE() view returns(bytes32)
func (_Ttmaccount *TtmaccountCaller) MESSENGERBOTROLE(opts *bind.CallOpts) ([32]byte, error) {
	var out []interface{}
	err := _Ttmaccount.contract.Call(opts, &out, "MESSENGER_BOT_ROLE")

	if err != nil {
		return *new([32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)

	return out0, err

}

// MESSENGERBOTROLE is a free data retrieval call binding the contract method 0x72afa328.
//
// Solidity: function MESSENGER_BOT_ROLE() view returns(bytes32)
func (_Ttmaccount *TtmaccountSession) MESSENGERBOTROLE() ([32]byte, error) {
	return _Ttmaccount.Contract.MESSENGERBOTROLE(&_Ttmaccount.CallOpts)
}

// MESSENGERBOTROLE is a free data retrieval call binding the contract method 0x72afa328.
//
// Solidity: function MESSENGER_BOT_ROLE() view returns(bytes32)
func (_Ttmaccount *TtmaccountCallerSession) MESSENGERBOTROLE() ([32]byte, error) {
	return _Ttmaccount.Contract.MESSENGERBOTROLE(&_Ttmaccount.CallOpts)
}

// SERVICEADMINROLE is a free data retrieval call binding the contract method 0xd09445c2.
//
// Solidity: function SERVICE_ADMIN_ROLE() view returns(bytes32)
func (_Ttmaccount *TtmaccountCaller) SERVICEADMINROLE(opts *bind.CallOpts) ([32]byte, error) {
	var out []interface{}
	err := _Ttmaccount.contract.Call(opts, &out, "SERVICE_ADMIN_ROLE")

	if err != nil {
		return *new([32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)

	return out0, err

}

// SERVICEADMINROLE is a free data retrieval call binding the contract method 0xd09445c2.
//
// Solidity: function SERVICE_ADMIN_ROLE() view returns(bytes32)
func (_Ttmaccount *TtmaccountSession) SERVICEADMINROLE() ([32]byte, error) {
	return _Ttmaccount.Contract.SERVICEADMINROLE(&_Ttmaccount.CallOpts)
}

// SERVICEADMINROLE is a free data retrieval call binding the contract method 0xd09445c2.
//
// Solidity: function SERVICE_ADMIN_ROLE() view returns(bytes32)
func (_Ttmaccount *TtmaccountCallerSession) SERVICEADMINROLE() ([32]byte, error) {
	return _Ttmaccount.Contract.SERVICEADMINROLE(&_Ttmaccount.CallOpts)
}

// UPGRADERROLE is a free data retrieval call binding the contract method 0xf72c0d8b.
//
// Solidity: function UPGRADER_ROLE() view returns(bytes32)
func (_Ttmaccount *TtmaccountCaller) UPGRADERROLE(opts *bind.CallOpts) ([32]byte, error) {
	var out []interface{}
	err := _Ttmaccount.contract.Call(opts, &out, "UPGRADER_ROLE")

	if err != nil {
		return *new([32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)

	return out0, err

}

// UPGRADERROLE is a free data retrieval call binding the contract method 0xf72c0d8b.
//
// Solidity: function UPGRADER_ROLE() view returns(bytes32)
func (_Ttmaccount *TtmaccountSession) UPGRADERROLE() ([32]byte, error) {
	return _Ttmaccount.Contract.UPGRADERROLE(&_Ttmaccount.CallOpts)
}

// UPGRADERROLE is a free data retrieval call binding the contract method 0xf72c0d8b.
//
// Solidity: function UPGRADER_ROLE() view returns(bytes32)
func (_Ttmaccount *TtmaccountCallerSession) UPGRADERROLE() ([32]byte, error) {
	return _Ttmaccount.Contract.UPGRADERROLE(&_Ttmaccount.CallOpts)
}

// UPGRADEINTERFACEVERSION is a free data retrieval call binding the contract method 0xad3cb1cc.
//
// Solidity: function UPGRADE_INTERFACE_VERSION() view returns(string)
func (_Ttmaccount *TtmaccountCaller) UPGRADEINTERFACEVERSION(opts *bind.CallOpts) (string, error) {
	var out []interface{}
	err := _Ttmaccount.contract.Call(opts, &out, "UPGRADE_INTERFACE_VERSION")

	if err != nil {
		return *new(string), err
	}

	out0 := *abi.ConvertType(out[0], new(string)).(*string)

	return out0, err

}

// UPGRADEINTERFACEVERSION is a free data retrieval call binding the contract method 0xad3cb1cc.
//
// Solidity: function UPGRADE_INTERFACE_VERSION() view returns(string)
func (_Ttmaccount *TtmaccountSession) UPGRADEINTERFACEVERSION() (string, error) {
	return _Ttmaccount.Contract.UPGRADEINTERFACEVERSION(&_Ttmaccount.CallOpts)
}

// UPGRADEINTERFACEVERSION is a free data retrieval call binding the contract method 0xad3cb1cc.
//
// Solidity: function UPGRADE_INTERFACE_VERSION() view returns(string)
func (_Ttmaccount *TtmaccountCallerSession) UPGRADEINTERFACEVERSION() (string, error) {
	return _Ttmaccount.Contract.UPGRADEINTERFACEVERSION(&_Ttmaccount.CallOpts)
}

// WITHDRAWERROLE is a free data retrieval call binding the contract method 0x85f438c1.
//
// Solidity: function WITHDRAWER_ROLE() view returns(bytes32)
func (_Ttmaccount *TtmaccountCaller) WITHDRAWERROLE(opts *bind.CallOpts) ([32]byte, error) {
	var out []interface{}
	err := _Ttmaccount.contract.Call(opts, &out, "WITHDRAWER_ROLE")

	if err != nil {
		return *new([32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)

	return out0, err

}

// WITHDRAWERROLE is a free data retrieval call binding the contract method 0x85f438c1.
//
// Solidity: function WITHDRAWER_ROLE() view returns(bytes32)
func (_Ttmaccount *TtmaccountSession) WITHDRAWERROLE() ([32]byte, error) {
	return _Ttmaccount.Contract.WITHDRAWERROLE(&_Ttmaccount.CallOpts)
}

// WITHDRAWERROLE is a free data retrieval call binding the contract method 0x85f438c1.
//
// Solidity: function WITHDRAWER_ROLE() view returns(bytes32)
func (_Ttmaccount *TtmaccountCallerSession) WITHDRAWERROLE() ([32]byte, error) {
	return _Ttmaccount.Contract.WITHDRAWERROLE(&_Ttmaccount.CallOpts)
}

// GetAllServiceHashes is a free data retrieval call binding the contract method 0x42072bbd.
//
// Solidity: function getAllServiceHashes() view returns(bytes32[] serviceHashes)
func (_Ttmaccount *TtmaccountCaller) GetAllServiceHashes(opts *bind.CallOpts) ([][32]byte, error) {
	var out []interface{}
	err := _Ttmaccount.contract.Call(opts, &out, "getAllServiceHashes")

	if err != nil {
		return *new([][32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([][32]byte)).(*[][32]byte)

	return out0, err

}

// GetAllServiceHashes is a free data retrieval call binding the contract method 0x42072bbd.
//
// Solidity: function getAllServiceHashes() view returns(bytes32[] serviceHashes)
func (_Ttmaccount *TtmaccountSession) GetAllServiceHashes() ([][32]byte, error) {
	return _Ttmaccount.Contract.GetAllServiceHashes(&_Ttmaccount.CallOpts)
}

// GetAllServiceHashes is a free data retrieval call binding the contract method 0x42072bbd.
//
// Solidity: function getAllServiceHashes() view returns(bytes32[] serviceHashes)
func (_Ttmaccount *TtmaccountCallerSession) GetAllServiceHashes() ([][32]byte, error) {
	return _Ttmaccount.Contract.GetAllServiceHashes(&_Ttmaccount.CallOpts)
}

// GetBookingTokenAddress is a free data retrieval call binding the contract method 0x4f3f4639.
//
// Solidity: function getBookingTokenAddress() view returns(address)
func (_Ttmaccount *TtmaccountCaller) GetBookingTokenAddress(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _Ttmaccount.contract.Call(opts, &out, "getBookingTokenAddress")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// GetBookingTokenAddress is a free data retrieval call binding the contract method 0x4f3f4639.
//
// Solidity: function getBookingTokenAddress() view returns(address)
func (_Ttmaccount *TtmaccountSession) GetBookingTokenAddress() (common.Address, error) {
	return _Ttmaccount.Contract.GetBookingTokenAddress(&_Ttmaccount.CallOpts)
}

// GetBookingTokenAddress is a free data retrieval call binding the contract method 0x4f3f4639.
//
// Solidity: function getBookingTokenAddress() view returns(address)
func (_Ttmaccount *TtmaccountCallerSession) GetBookingTokenAddress() (common.Address, error) {
	return _Ttmaccount.Contract.GetBookingTokenAddress(&_Ttmaccount.CallOpts)
}

// GetGasMoneyWithdrawal is a free data retrieval call binding the contract method 0x658db0af.
//
// Solidity: function getGasMoneyWithdrawal() view returns(uint256 withdrawalLimit, uint256 withdrawalPeriod)
func (_Ttmaccount *TtmaccountCaller) GetGasMoneyWithdrawal(opts *bind.CallOpts) (struct {
	WithdrawalLimit  *big.Int
	WithdrawalPeriod *big.Int
}, error) {
	var out []interface{}
	err := _Ttmaccount.contract.Call(opts, &out, "getGasMoneyWithdrawal")

	outstruct := new(struct {
		WithdrawalLimit  *big.Int
		WithdrawalPeriod *big.Int
	})
	if err != nil {
		return *outstruct, err
	}

	outstruct.WithdrawalLimit = *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)
	outstruct.WithdrawalPeriod = *abi.ConvertType(out[1], new(*big.Int)).(**big.Int)

	return *outstruct, err

}

// GetGasMoneyWithdrawal is a free data retrieval call binding the contract method 0x658db0af.
//
// Solidity: function getGasMoneyWithdrawal() view returns(uint256 withdrawalLimit, uint256 withdrawalPeriod)
func (_Ttmaccount *TtmaccountSession) GetGasMoneyWithdrawal() (struct {
	WithdrawalLimit  *big.Int
	WithdrawalPeriod *big.Int
}, error) {
	return _Ttmaccount.Contract.GetGasMoneyWithdrawal(&_Ttmaccount.CallOpts)
}

// GetGasMoneyWithdrawal is a free data retrieval call binding the contract method 0x658db0af.
//
// Solidity: function getGasMoneyWithdrawal() view returns(uint256 withdrawalLimit, uint256 withdrawalPeriod)
func (_Ttmaccount *TtmaccountCallerSession) GetGasMoneyWithdrawal() (struct {
	WithdrawalLimit  *big.Int
	WithdrawalPeriod *big.Int
}, error) {
	return _Ttmaccount.Contract.GetGasMoneyWithdrawal(&_Ttmaccount.CallOpts)
}

// GetGasMoneyWithdrawalForAccount is a free data retrieval call binding the contract method 0xee3b641f.
//
// Solidity: function getGasMoneyWithdrawalForAccount(address account) view returns(uint256 periodStart, uint256 withdrawnAmount)
func (_Ttmaccount *TtmaccountCaller) GetGasMoneyWithdrawalForAccount(opts *bind.CallOpts, account common.Address) (struct {
	PeriodStart     *big.Int
	WithdrawnAmount *big.Int
}, error) {
	var out []interface{}
	err := _Ttmaccount.contract.Call(opts, &out, "getGasMoneyWithdrawalForAccount", account)

	outstruct := new(struct {
		PeriodStart     *big.Int
		WithdrawnAmount *big.Int
	})
	if err != nil {
		return *outstruct, err
	}

	outstruct.PeriodStart = *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)
	outstruct.WithdrawnAmount = *abi.ConvertType(out[1], new(*big.Int)).(**big.Int)

	return *outstruct, err

}

// GetGasMoneyWithdrawalForAccount is a free data retrieval call binding the contract method 0xee3b641f.
//
// Solidity: function getGasMoneyWithdrawalForAccount(address account) view returns(uint256 periodStart, uint256 withdrawnAmount)
func (_Ttmaccount *TtmaccountSession) GetGasMoneyWithdrawalForAccount(account common.Address) (struct {
	PeriodStart     *big.Int
	WithdrawnAmount *big.Int
}, error) {
	return _Ttmaccount.Contract.GetGasMoneyWithdrawalForAccount(&_Ttmaccount.CallOpts, account)
}

// GetGasMoneyWithdrawalForAccount is a free data retrieval call binding the contract method 0xee3b641f.
//
// Solidity: function getGasMoneyWithdrawalForAccount(address account) view returns(uint256 periodStart, uint256 withdrawnAmount)
func (_Ttmaccount *TtmaccountCallerSession) GetGasMoneyWithdrawalForAccount(account common.Address) (struct {
	PeriodStart     *big.Int
	WithdrawnAmount *big.Int
}, error) {
	return _Ttmaccount.Contract.GetGasMoneyWithdrawalForAccount(&_Ttmaccount.CallOpts, account)
}

// GetManagerAddress is a free data retrieval call binding the contract method 0xc162d7da.
//
// Solidity: function getManagerAddress() view returns(address)
func (_Ttmaccount *TtmaccountCaller) GetManagerAddress(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _Ttmaccount.contract.Call(opts, &out, "getManagerAddress")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// GetManagerAddress is a free data retrieval call binding the contract method 0xc162d7da.
//
// Solidity: function getManagerAddress() view returns(address)
func (_Ttmaccount *TtmaccountSession) GetManagerAddress() (common.Address, error) {
	return _Ttmaccount.Contract.GetManagerAddress(&_Ttmaccount.CallOpts)
}

// GetManagerAddress is a free data retrieval call binding the contract method 0xc162d7da.
//
// Solidity: function getManagerAddress() view returns(address)
func (_Ttmaccount *TtmaccountCallerSession) GetManagerAddress() (common.Address, error) {
	return _Ttmaccount.Contract.GetManagerAddress(&_Ttmaccount.CallOpts)
}

// GetPublicKey is a free data retrieval call binding the contract method 0x857cdbb8.
//
// Solidity: function getPublicKey(address pubKeyAddress) view returns(bytes data)
func (_Ttmaccount *TtmaccountCaller) GetPublicKey(opts *bind.CallOpts, pubKeyAddress common.Address) ([]byte, error) {
	var out []interface{}
	err := _Ttmaccount.contract.Call(opts, &out, "getPublicKey", pubKeyAddress)

	if err != nil {
		return *new([]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([]byte)).(*[]byte)

	return out0, err

}

// GetPublicKey is a free data retrieval call binding the contract method 0x857cdbb8.
//
// Solidity: function getPublicKey(address pubKeyAddress) view returns(bytes data)
func (_Ttmaccount *TtmaccountSession) GetPublicKey(pubKeyAddress common.Address) ([]byte, error) {
	return _Ttmaccount.Contract.GetPublicKey(&_Ttmaccount.CallOpts, pubKeyAddress)
}

// GetPublicKey is a free data retrieval call binding the contract method 0x857cdbb8.
//
// Solidity: function getPublicKey(address pubKeyAddress) view returns(bytes data)
func (_Ttmaccount *TtmaccountCallerSession) GetPublicKey(pubKeyAddress common.Address) ([]byte, error) {
	return _Ttmaccount.Contract.GetPublicKey(&_Ttmaccount.CallOpts, pubKeyAddress)
}

// GetPublicKeysAddresses is a free data retrieval call binding the contract method 0xea79d07a.
//
// Solidity: function getPublicKeysAddresses() view returns(address[] pubKeyAddresses)
func (_Ttmaccount *TtmaccountCaller) GetPublicKeysAddresses(opts *bind.CallOpts) ([]common.Address, error) {
	var out []interface{}
	err := _Ttmaccount.contract.Call(opts, &out, "getPublicKeysAddresses")

	if err != nil {
		return *new([]common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new([]common.Address)).(*[]common.Address)

	return out0, err

}

// GetPublicKeysAddresses is a free data retrieval call binding the contract method 0xea79d07a.
//
// Solidity: function getPublicKeysAddresses() view returns(address[] pubKeyAddresses)
func (_Ttmaccount *TtmaccountSession) GetPublicKeysAddresses() ([]common.Address, error) {
	return _Ttmaccount.Contract.GetPublicKeysAddresses(&_Ttmaccount.CallOpts)
}

// GetPublicKeysAddresses is a free data retrieval call binding the contract method 0xea79d07a.
//
// Solidity: function getPublicKeysAddresses() view returns(address[] pubKeyAddresses)
func (_Ttmaccount *TtmaccountCallerSession) GetPublicKeysAddresses() ([]common.Address, error) {
	return _Ttmaccount.Contract.GetPublicKeysAddresses(&_Ttmaccount.CallOpts)
}

// GetRoleAdmin is a free data retrieval call binding the contract method 0x248a9ca3.
//
// Solidity: function getRoleAdmin(bytes32 role) view returns(bytes32)
func (_Ttmaccount *TtmaccountCaller) GetRoleAdmin(opts *bind.CallOpts, role [32]byte) ([32]byte, error) {
	var out []interface{}
	err := _Ttmaccount.contract.Call(opts, &out, "getRoleAdmin", role)

	if err != nil {
		return *new([32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)

	return out0, err

}

// GetRoleAdmin is a free data retrieval call binding the contract method 0x248a9ca3.
//
// Solidity: function getRoleAdmin(bytes32 role) view returns(bytes32)
func (_Ttmaccount *TtmaccountSession) GetRoleAdmin(role [32]byte) ([32]byte, error) {
	return _Ttmaccount.Contract.GetRoleAdmin(&_Ttmaccount.CallOpts, role)
}

// GetRoleAdmin is a free data retrieval call binding the contract method 0x248a9ca3.
//
// Solidity: function getRoleAdmin(bytes32 role) view returns(bytes32)
func (_Ttmaccount *TtmaccountCallerSession) GetRoleAdmin(role [32]byte) ([32]byte, error) {
	return _Ttmaccount.Contract.GetRoleAdmin(&_Ttmaccount.CallOpts, role)
}

// GetRoleMember is a free data retrieval call binding the contract method 0x9010d07c.
//
// Solidity: function getRoleMember(bytes32 role, uint256 index) view returns(address)
func (_Ttmaccount *TtmaccountCaller) GetRoleMember(opts *bind.CallOpts, role [32]byte, index *big.Int) (common.Address, error) {
	var out []interface{}
	err := _Ttmaccount.contract.Call(opts, &out, "getRoleMember", role, index)

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// GetRoleMember is a free data retrieval call binding the contract method 0x9010d07c.
//
// Solidity: function getRoleMember(bytes32 role, uint256 index) view returns(address)
func (_Ttmaccount *TtmaccountSession) GetRoleMember(role [32]byte, index *big.Int) (common.Address, error) {
	return _Ttmaccount.Contract.GetRoleMember(&_Ttmaccount.CallOpts, role, index)
}

// GetRoleMember is a free data retrieval call binding the contract method 0x9010d07c.
//
// Solidity: function getRoleMember(bytes32 role, uint256 index) view returns(address)
func (_Ttmaccount *TtmaccountCallerSession) GetRoleMember(role [32]byte, index *big.Int) (common.Address, error) {
	return _Ttmaccount.Contract.GetRoleMember(&_Ttmaccount.CallOpts, role, index)
}

// GetRoleMemberCount is a free data retrieval call binding the contract method 0xca15c873.
//
// Solidity: function getRoleMemberCount(bytes32 role) view returns(uint256)
func (_Ttmaccount *TtmaccountCaller) GetRoleMemberCount(opts *bind.CallOpts, role [32]byte) (*big.Int, error) {
	var out []interface{}
	err := _Ttmaccount.contract.Call(opts, &out, "getRoleMemberCount", role)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// GetRoleMemberCount is a free data retrieval call binding the contract method 0xca15c873.
//
// Solidity: function getRoleMemberCount(bytes32 role) view returns(uint256)
func (_Ttmaccount *TtmaccountSession) GetRoleMemberCount(role [32]byte) (*big.Int, error) {
	return _Ttmaccount.Contract.GetRoleMemberCount(&_Ttmaccount.CallOpts, role)
}

// GetRoleMemberCount is a free data retrieval call binding the contract method 0xca15c873.
//
// Solidity: function getRoleMemberCount(bytes32 role) view returns(uint256)
func (_Ttmaccount *TtmaccountCallerSession) GetRoleMemberCount(role [32]byte) (*big.Int, error) {
	return _Ttmaccount.Contract.GetRoleMemberCount(&_Ttmaccount.CallOpts, role)
}

// GetRoleMembers is a free data retrieval call binding the contract method 0xa3246ad3.
//
// Solidity: function getRoleMembers(bytes32 role) view returns(address[])
func (_Ttmaccount *TtmaccountCaller) GetRoleMembers(opts *bind.CallOpts, role [32]byte) ([]common.Address, error) {
	var out []interface{}
	err := _Ttmaccount.contract.Call(opts, &out, "getRoleMembers", role)

	if err != nil {
		return *new([]common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new([]common.Address)).(*[]common.Address)

	return out0, err

}

// GetRoleMembers is a free data retrieval call binding the contract method 0xa3246ad3.
//
// Solidity: function getRoleMembers(bytes32 role) view returns(address[])
func (_Ttmaccount *TtmaccountSession) GetRoleMembers(role [32]byte) ([]common.Address, error) {
	return _Ttmaccount.Contract.GetRoleMembers(&_Ttmaccount.CallOpts, role)
}

// GetRoleMembers is a free data retrieval call binding the contract method 0xa3246ad3.
//
// Solidity: function getRoleMembers(bytes32 role) view returns(address[])
func (_Ttmaccount *TtmaccountCallerSession) GetRoleMembers(role [32]byte) ([]common.Address, error) {
	return _Ttmaccount.Contract.GetRoleMembers(&_Ttmaccount.CallOpts, role)
}

// GetService is a free data retrieval call binding the contract method 0xda47d856.
//
// Solidity: function getService(bytes32 serviceHash) view returns((bool,string[]) service)
func (_Ttmaccount *TtmaccountCaller) GetService(opts *bind.CallOpts, serviceHash [32]byte) (PartnerConfigurationService, error) {
	var out []interface{}
	err := _Ttmaccount.contract.Call(opts, &out, "getService", serviceHash)

	if err != nil {
		return *new(PartnerConfigurationService), err
	}

	out0 := *abi.ConvertType(out[0], new(PartnerConfigurationService)).(*PartnerConfigurationService)

	return out0, err

}

// GetService is a free data retrieval call binding the contract method 0xda47d856.
//
// Solidity: function getService(bytes32 serviceHash) view returns((bool,string[]) service)
func (_Ttmaccount *TtmaccountSession) GetService(serviceHash [32]byte) (PartnerConfigurationService, error) {
	return _Ttmaccount.Contract.GetService(&_Ttmaccount.CallOpts, serviceHash)
}

// GetService is a free data retrieval call binding the contract method 0xda47d856.
//
// Solidity: function getService(bytes32 serviceHash) view returns((bool,string[]) service)
func (_Ttmaccount *TtmaccountCallerSession) GetService(serviceHash [32]byte) (PartnerConfigurationService, error) {
	return _Ttmaccount.Contract.GetService(&_Ttmaccount.CallOpts, serviceHash)
}

// GetServiceCapabilities is a free data retrieval call binding the contract method 0x5e07f869.
//
// Solidity: function getServiceCapabilities(bytes32 serviceHash) view returns(string[] capabilities)
func (_Ttmaccount *TtmaccountCaller) GetServiceCapabilities(opts *bind.CallOpts, serviceHash [32]byte) ([]string, error) {
	var out []interface{}
	err := _Ttmaccount.contract.Call(opts, &out, "getServiceCapabilities", serviceHash)

	if err != nil {
		return *new([]string), err
	}

	out0 := *abi.ConvertType(out[0], new([]string)).(*[]string)

	return out0, err

}

// GetServiceCapabilities is a free data retrieval call binding the contract method 0x5e07f869.
//
// Solidity: function getServiceCapabilities(bytes32 serviceHash) view returns(string[] capabilities)
func (_Ttmaccount *TtmaccountSession) GetServiceCapabilities(serviceHash [32]byte) ([]string, error) {
	return _Ttmaccount.Contract.GetServiceCapabilities(&_Ttmaccount.CallOpts, serviceHash)
}

// GetServiceCapabilities is a free data retrieval call binding the contract method 0x5e07f869.
//
// Solidity: function getServiceCapabilities(bytes32 serviceHash) view returns(string[] capabilities)
func (_Ttmaccount *TtmaccountCallerSession) GetServiceCapabilities(serviceHash [32]byte) ([]string, error) {
	return _Ttmaccount.Contract.GetServiceCapabilities(&_Ttmaccount.CallOpts, serviceHash)
}

// GetServiceRestrictedRate is a free data retrieval call binding the contract method 0x8f69347d.
//
// Solidity: function getServiceRestrictedRate(bytes32 serviceHash) view returns(bool restrictedRate)
func (_Ttmaccount *TtmaccountCaller) GetServiceRestrictedRate(opts *bind.CallOpts, serviceHash [32]byte) (bool, error) {
	var out []interface{}
	err := _Ttmaccount.contract.Call(opts, &out, "getServiceRestrictedRate", serviceHash)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// GetServiceRestrictedRate is a free data retrieval call binding the contract method 0x8f69347d.
//
// Solidity: function getServiceRestrictedRate(bytes32 serviceHash) view returns(bool restrictedRate)
func (_Ttmaccount *TtmaccountSession) GetServiceRestrictedRate(serviceHash [32]byte) (bool, error) {
	return _Ttmaccount.Contract.GetServiceRestrictedRate(&_Ttmaccount.CallOpts, serviceHash)
}

// GetServiceRestrictedRate is a free data retrieval call binding the contract method 0x8f69347d.
//
// Solidity: function getServiceRestrictedRate(bytes32 serviceHash) view returns(bool restrictedRate)
func (_Ttmaccount *TtmaccountCallerSession) GetServiceRestrictedRate(serviceHash [32]byte) (bool, error) {
	return _Ttmaccount.Contract.GetServiceRestrictedRate(&_Ttmaccount.CallOpts, serviceHash)
}

// GetSupportedServices is a free data retrieval call binding the contract method 0x7eec56c7.
//
// Solidity: function getSupportedServices() view returns(bytes32[] serviceHashes, (bool,string[])[] services)
func (_Ttmaccount *TtmaccountCaller) GetSupportedServices(opts *bind.CallOpts) (struct {
	ServiceHashes [][32]byte
	Services      []PartnerConfigurationService
}, error) {
	var out []interface{}
	err := _Ttmaccount.contract.Call(opts, &out, "getSupportedServices")

	outstruct := new(struct {
		ServiceHashes [][32]byte
		Services      []PartnerConfigurationService
	})
	if err != nil {
		return *outstruct, err
	}

	outstruct.ServiceHashes = *abi.ConvertType(out[0], new([][32]byte)).(*[][32]byte)
	outstruct.Services = *abi.ConvertType(out[1], new([]PartnerConfigurationService)).(*[]PartnerConfigurationService)

	return *outstruct, err

}

// GetSupportedServices is a free data retrieval call binding the contract method 0x7eec56c7.
//
// Solidity: function getSupportedServices() view returns(bytes32[] serviceHashes, (bool,string[])[] services)
func (_Ttmaccount *TtmaccountSession) GetSupportedServices() (struct {
	ServiceHashes [][32]byte
	Services      []PartnerConfigurationService
}, error) {
	return _Ttmaccount.Contract.GetSupportedServices(&_Ttmaccount.CallOpts)
}

// GetSupportedServices is a free data retrieval call binding the contract method 0x7eec56c7.
//
// Solidity: function getSupportedServices() view returns(bytes32[] serviceHashes, (bool,string[])[] services)
func (_Ttmaccount *TtmaccountCallerSession) GetSupportedServices() (struct {
	ServiceHashes [][32]byte
	Services      []PartnerConfigurationService
}, error) {
	return _Ttmaccount.Contract.GetSupportedServices(&_Ttmaccount.CallOpts)
}

// GetSupportedServicesSlice is a free data retrieval call binding the contract method 0x08b1565a.
//
// Solidity: function getSupportedServicesSlice(uint256 offset, uint256 limit) view returns(bytes32[] serviceHashes, (bool,string[])[] services)
func (_Ttmaccount *TtmaccountCaller) GetSupportedServicesSlice(opts *bind.CallOpts, offset *big.Int, limit *big.Int) (struct {
	ServiceHashes [][32]byte
	Services      []PartnerConfigurationService
}, error) {
	var out []interface{}
	err := _Ttmaccount.contract.Call(opts, &out, "getSupportedServicesSlice", offset, limit)

	outstruct := new(struct {
		ServiceHashes [][32]byte
		Services      []PartnerConfigurationService
	})
	if err != nil {
		return *outstruct, err
	}

	outstruct.ServiceHashes = *abi.ConvertType(out[0], new([][32]byte)).(*[][32]byte)
	outstruct.Services = *abi.ConvertType(out[1], new([]PartnerConfigurationService)).(*[]PartnerConfigurationService)

	return *outstruct, err

}

// GetSupportedServicesSlice is a free data retrieval call binding the contract method 0x08b1565a.
//
// Solidity: function getSupportedServicesSlice(uint256 offset, uint256 limit) view returns(bytes32[] serviceHashes, (bool,string[])[] services)
func (_Ttmaccount *TtmaccountSession) GetSupportedServicesSlice(offset *big.Int, limit *big.Int) (struct {
	ServiceHashes [][32]byte
	Services      []PartnerConfigurationService
}, error) {
	return _Ttmaccount.Contract.GetSupportedServicesSlice(&_Ttmaccount.CallOpts, offset, limit)
}

// GetSupportedServicesSlice is a free data retrieval call binding the contract method 0x08b1565a.
//
// Solidity: function getSupportedServicesSlice(uint256 offset, uint256 limit) view returns(bytes32[] serviceHashes, (bool,string[])[] services)
func (_Ttmaccount *TtmaccountCallerSession) GetSupportedServicesSlice(offset *big.Int, limit *big.Int) (struct {
	ServiceHashes [][32]byte
	Services      []PartnerConfigurationService
}, error) {
	return _Ttmaccount.Contract.GetSupportedServicesSlice(&_Ttmaccount.CallOpts, offset, limit)
}

// GetSupportedTokens is a free data retrieval call binding the contract method 0xd3c7c2c7.
//
// Solidity: function getSupportedTokens() view returns(address[] tokens)
func (_Ttmaccount *TtmaccountCaller) GetSupportedTokens(opts *bind.CallOpts) ([]common.Address, error) {
	var out []interface{}
	err := _Ttmaccount.contract.Call(opts, &out, "getSupportedTokens")

	if err != nil {
		return *new([]common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new([]common.Address)).(*[]common.Address)

	return out0, err

}

// GetSupportedTokens is a free data retrieval call binding the contract method 0xd3c7c2c7.
//
// Solidity: function getSupportedTokens() view returns(address[] tokens)
func (_Ttmaccount *TtmaccountSession) GetSupportedTokens() ([]common.Address, error) {
	return _Ttmaccount.Contract.GetSupportedTokens(&_Ttmaccount.CallOpts)
}

// GetSupportedTokens is a free data retrieval call binding the contract method 0xd3c7c2c7.
//
// Solidity: function getSupportedTokens() view returns(address[] tokens)
func (_Ttmaccount *TtmaccountCallerSession) GetSupportedTokens() ([]common.Address, error) {
	return _Ttmaccount.Contract.GetSupportedTokens(&_Ttmaccount.CallOpts)
}

// GetWantedServiceHashes is a free data retrieval call binding the contract method 0x136f50ca.
//
// Solidity: function getWantedServiceHashes() view returns(bytes32[] serviceHashes)
func (_Ttmaccount *TtmaccountCaller) GetWantedServiceHashes(opts *bind.CallOpts) ([][32]byte, error) {
	var out []interface{}
	err := _Ttmaccount.contract.Call(opts, &out, "getWantedServiceHashes")

	if err != nil {
		return *new([][32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([][32]byte)).(*[][32]byte)

	return out0, err

}

// GetWantedServiceHashes is a free data retrieval call binding the contract method 0x136f50ca.
//
// Solidity: function getWantedServiceHashes() view returns(bytes32[] serviceHashes)
func (_Ttmaccount *TtmaccountSession) GetWantedServiceHashes() ([][32]byte, error) {
	return _Ttmaccount.Contract.GetWantedServiceHashes(&_Ttmaccount.CallOpts)
}

// GetWantedServiceHashes is a free data retrieval call binding the contract method 0x136f50ca.
//
// Solidity: function getWantedServiceHashes() view returns(bytes32[] serviceHashes)
func (_Ttmaccount *TtmaccountCallerSession) GetWantedServiceHashes() ([][32]byte, error) {
	return _Ttmaccount.Contract.GetWantedServiceHashes(&_Ttmaccount.CallOpts)
}

// HasRole is a free data retrieval call binding the contract method 0x91d14854.
//
// Solidity: function hasRole(bytes32 role, address account) view returns(bool)
func (_Ttmaccount *TtmaccountCaller) HasRole(opts *bind.CallOpts, role [32]byte, account common.Address) (bool, error) {
	var out []interface{}
	err := _Ttmaccount.contract.Call(opts, &out, "hasRole", role, account)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// HasRole is a free data retrieval call binding the contract method 0x91d14854.
//
// Solidity: function hasRole(bytes32 role, address account) view returns(bool)
func (_Ttmaccount *TtmaccountSession) HasRole(role [32]byte, account common.Address) (bool, error) {
	return _Ttmaccount.Contract.HasRole(&_Ttmaccount.CallOpts, role, account)
}

// HasRole is a free data retrieval call binding the contract method 0x91d14854.
//
// Solidity: function hasRole(bytes32 role, address account) view returns(bool)
func (_Ttmaccount *TtmaccountCallerSession) HasRole(role [32]byte, account common.Address) (bool, error) {
	return _Ttmaccount.Contract.HasRole(&_Ttmaccount.CallOpts, role, account)
}

// IsBotAllowed is a free data retrieval call binding the contract method 0xe0b78add.
//
// Solidity: function isBotAllowed(address bot) view returns(bool)
func (_Ttmaccount *TtmaccountCaller) IsBotAllowed(opts *bind.CallOpts, bot common.Address) (bool, error) {
	var out []interface{}
	err := _Ttmaccount.contract.Call(opts, &out, "isBotAllowed", bot)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// IsBotAllowed is a free data retrieval call binding the contract method 0xe0b78add.
//
// Solidity: function isBotAllowed(address bot) view returns(bool)
func (_Ttmaccount *TtmaccountSession) IsBotAllowed(bot common.Address) (bool, error) {
	return _Ttmaccount.Contract.IsBotAllowed(&_Ttmaccount.CallOpts, bot)
}

// IsBotAllowed is a free data retrieval call binding the contract method 0xe0b78add.
//
// Solidity: function isBotAllowed(address bot) view returns(bool)
func (_Ttmaccount *TtmaccountCallerSession) IsBotAllowed(bot common.Address) (bool, error) {
	return _Ttmaccount.Contract.IsBotAllowed(&_Ttmaccount.CallOpts, bot)
}

// IsServiceSupported is a free data retrieval call binding the contract method 0xbd252c1c.
//
// Solidity: function isServiceSupported(bytes32 serviceHash) view returns(bool)
func (_Ttmaccount *TtmaccountCaller) IsServiceSupported(opts *bind.CallOpts, serviceHash [32]byte) (bool, error) {
	var out []interface{}
	err := _Ttmaccount.contract.Call(opts, &out, "isServiceSupported", serviceHash)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// IsServiceSupported is a free data retrieval call binding the contract method 0xbd252c1c.
//
// Solidity: function isServiceSupported(bytes32 serviceHash) view returns(bool)
func (_Ttmaccount *TtmaccountSession) IsServiceSupported(serviceHash [32]byte) (bool, error) {
	return _Ttmaccount.Contract.IsServiceSupported(&_Ttmaccount.CallOpts, serviceHash)
}

// IsServiceSupported is a free data retrieval call binding the contract method 0xbd252c1c.
//
// Solidity: function isServiceSupported(bytes32 serviceHash) view returns(bool)
func (_Ttmaccount *TtmaccountCallerSession) IsServiceSupported(serviceHash [32]byte) (bool, error) {
	return _Ttmaccount.Contract.IsServiceSupported(&_Ttmaccount.CallOpts, serviceHash)
}

// IsSupportedToken is a free data retrieval call binding the contract method 0x240028e8.
//
// Solidity: function isSupportedToken(address _token) view returns(bool supported)
func (_Ttmaccount *TtmaccountCaller) IsSupportedToken(opts *bind.CallOpts, _token common.Address) (bool, error) {
	var out []interface{}
	err := _Ttmaccount.contract.Call(opts, &out, "isSupportedToken", _token)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// IsSupportedToken is a free data retrieval call binding the contract method 0x240028e8.
//
// Solidity: function isSupportedToken(address _token) view returns(bool supported)
func (_Ttmaccount *TtmaccountSession) IsSupportedToken(_token common.Address) (bool, error) {
	return _Ttmaccount.Contract.IsSupportedToken(&_Ttmaccount.CallOpts, _token)
}

// IsSupportedToken is a free data retrieval call binding the contract method 0x240028e8.
//
// Solidity: function isSupportedToken(address _token) view returns(bool supported)
func (_Ttmaccount *TtmaccountCallerSession) IsSupportedToken(_token common.Address) (bool, error) {
	return _Ttmaccount.Contract.IsSupportedToken(&_Ttmaccount.CallOpts, _token)
}

// ProxiableUUID is a free data retrieval call binding the contract method 0x52d1902d.
//
// Solidity: function proxiableUUID() view returns(bytes32)
func (_Ttmaccount *TtmaccountCaller) ProxiableUUID(opts *bind.CallOpts) ([32]byte, error) {
	var out []interface{}
	err := _Ttmaccount.contract.Call(opts, &out, "proxiableUUID")

	if err != nil {
		return *new([32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)

	return out0, err

}

// ProxiableUUID is a free data retrieval call binding the contract method 0x52d1902d.
//
// Solidity: function proxiableUUID() view returns(bytes32)
func (_Ttmaccount *TtmaccountSession) ProxiableUUID() ([32]byte, error) {
	return _Ttmaccount.Contract.ProxiableUUID(&_Ttmaccount.CallOpts)
}

// ProxiableUUID is a free data retrieval call binding the contract method 0x52d1902d.
//
// Solidity: function proxiableUUID() view returns(bytes32)
func (_Ttmaccount *TtmaccountCallerSession) ProxiableUUID() ([32]byte, error) {
	return _Ttmaccount.Contract.ProxiableUUID(&_Ttmaccount.CallOpts)
}

// SupportsInterface is a free data retrieval call binding the contract method 0x01ffc9a7.
//
// Solidity: function supportsInterface(bytes4 interfaceId) view returns(bool)
func (_Ttmaccount *TtmaccountCaller) SupportsInterface(opts *bind.CallOpts, interfaceId [4]byte) (bool, error) {
	var out []interface{}
	err := _Ttmaccount.contract.Call(opts, &out, "supportsInterface", interfaceId)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// SupportsInterface is a free data retrieval call binding the contract method 0x01ffc9a7.
//
// Solidity: function supportsInterface(bytes4 interfaceId) view returns(bool)
func (_Ttmaccount *TtmaccountSession) SupportsInterface(interfaceId [4]byte) (bool, error) {
	return _Ttmaccount.Contract.SupportsInterface(&_Ttmaccount.CallOpts, interfaceId)
}

// SupportsInterface is a free data retrieval call binding the contract method 0x01ffc9a7.
//
// Solidity: function supportsInterface(bytes4 interfaceId) view returns(bool)
func (_Ttmaccount *TtmaccountCallerSession) SupportsInterface(interfaceId [4]byte) (bool, error) {
	return _Ttmaccount.Contract.SupportsInterface(&_Ttmaccount.CallOpts, interfaceId)
}

// AcceptCancellation is a paid mutator transaction binding the contract method 0xbe667188.
//
// Solidity: function acceptCancellation(uint256 tokenId, uint256 refundAmount) returns()
func (_Ttmaccount *TtmaccountTransactor) AcceptCancellation(opts *bind.TransactOpts, tokenId *big.Int, refundAmount *big.Int) (*types.Transaction, error) {
	return _Ttmaccount.contract.Transact(opts, "acceptCancellation", tokenId, refundAmount)
}

// AcceptCancellation is a paid mutator transaction binding the contract method 0xbe667188.
//
// Solidity: function acceptCancellation(uint256 tokenId, uint256 refundAmount) returns()
func (_Ttmaccount *TtmaccountSession) AcceptCancellation(tokenId *big.Int, refundAmount *big.Int) (*types.Transaction, error) {
	return _Ttmaccount.Contract.AcceptCancellation(&_Ttmaccount.TransactOpts, tokenId, refundAmount)
}

// AcceptCancellation is a paid mutator transaction binding the contract method 0xbe667188.
//
// Solidity: function acceptCancellation(uint256 tokenId, uint256 refundAmount) returns()
func (_Ttmaccount *TtmaccountTransactorSession) AcceptCancellation(tokenId *big.Int, refundAmount *big.Int) (*types.Transaction, error) {
	return _Ttmaccount.Contract.AcceptCancellation(&_Ttmaccount.TransactOpts, tokenId, refundAmount)
}

// AddMessengerBot is a paid mutator transaction binding the contract method 0x51889d6b.
//
// Solidity: function addMessengerBot(address bot, uint256 gasMoney) returns()
func (_Ttmaccount *TtmaccountTransactor) AddMessengerBot(opts *bind.TransactOpts, bot common.Address, gasMoney *big.Int) (*types.Transaction, error) {
	return _Ttmaccount.contract.Transact(opts, "addMessengerBot", bot, gasMoney)
}

// AddMessengerBot is a paid mutator transaction binding the contract method 0x51889d6b.
//
// Solidity: function addMessengerBot(address bot, uint256 gasMoney) returns()
func (_Ttmaccount *TtmaccountSession) AddMessengerBot(bot common.Address, gasMoney *big.Int) (*types.Transaction, error) {
	return _Ttmaccount.Contract.AddMessengerBot(&_Ttmaccount.TransactOpts, bot, gasMoney)
}

// AddMessengerBot is a paid mutator transaction binding the contract method 0x51889d6b.
//
// Solidity: function addMessengerBot(address bot, uint256 gasMoney) returns()
func (_Ttmaccount *TtmaccountTransactorSession) AddMessengerBot(bot common.Address, gasMoney *big.Int) (*types.Transaction, error) {
	return _Ttmaccount.Contract.AddMessengerBot(&_Ttmaccount.TransactOpts, bot, gasMoney)
}

// AddPublicKey is a paid mutator transaction binding the contract method 0xccde65dc.
//
// Solidity: function addPublicKey(address pubKeyAddress, bytes data) returns()
func (_Ttmaccount *TtmaccountTransactor) AddPublicKey(opts *bind.TransactOpts, pubKeyAddress common.Address, data []byte) (*types.Transaction, error) {
	return _Ttmaccount.contract.Transact(opts, "addPublicKey", pubKeyAddress, data)
}

// AddPublicKey is a paid mutator transaction binding the contract method 0xccde65dc.
//
// Solidity: function addPublicKey(address pubKeyAddress, bytes data) returns()
func (_Ttmaccount *TtmaccountSession) AddPublicKey(pubKeyAddress common.Address, data []byte) (*types.Transaction, error) {
	return _Ttmaccount.Contract.AddPublicKey(&_Ttmaccount.TransactOpts, pubKeyAddress, data)
}

// AddPublicKey is a paid mutator transaction binding the contract method 0xccde65dc.
//
// Solidity: function addPublicKey(address pubKeyAddress, bytes data) returns()
func (_Ttmaccount *TtmaccountTransactorSession) AddPublicKey(pubKeyAddress common.Address, data []byte) (*types.Transaction, error) {
	return _Ttmaccount.Contract.AddPublicKey(&_Ttmaccount.TransactOpts, pubKeyAddress, data)
}

// AddService is a paid mutator transaction binding the contract method 0x7c5d62b3.
//
// Solidity: function addService(bytes32 serviceHash, bool restrictedRate, string[] capabilities) returns()
func (_Ttmaccount *TtmaccountTransactor) AddService(opts *bind.TransactOpts, serviceHash [32]byte, restrictedRate bool, capabilities []string) (*types.Transaction, error) {
	return _Ttmaccount.contract.Transact(opts, "addService", serviceHash, restrictedRate, capabilities)
}

// AddService is a paid mutator transaction binding the contract method 0x7c5d62b3.
//
// Solidity: function addService(bytes32 serviceHash, bool restrictedRate, string[] capabilities) returns()
func (_Ttmaccount *TtmaccountSession) AddService(serviceHash [32]byte, restrictedRate bool, capabilities []string) (*types.Transaction, error) {
	return _Ttmaccount.Contract.AddService(&_Ttmaccount.TransactOpts, serviceHash, restrictedRate, capabilities)
}

// AddService is a paid mutator transaction binding the contract method 0x7c5d62b3.
//
// Solidity: function addService(bytes32 serviceHash, bool restrictedRate, string[] capabilities) returns()
func (_Ttmaccount *TtmaccountTransactorSession) AddService(serviceHash [32]byte, restrictedRate bool, capabilities []string) (*types.Transaction, error) {
	return _Ttmaccount.Contract.AddService(&_Ttmaccount.TransactOpts, serviceHash, restrictedRate, capabilities)
}

// AddServiceCapability is a paid mutator transaction binding the contract method 0x5ae733c9.
//
// Solidity: function addServiceCapability(bytes32 serviceHash, string capability) returns()
func (_Ttmaccount *TtmaccountTransactor) AddServiceCapability(opts *bind.TransactOpts, serviceHash [32]byte, capability string) (*types.Transaction, error) {
	return _Ttmaccount.contract.Transact(opts, "addServiceCapability", serviceHash, capability)
}

// AddServiceCapability is a paid mutator transaction binding the contract method 0x5ae733c9.
//
// Solidity: function addServiceCapability(bytes32 serviceHash, string capability) returns()
func (_Ttmaccount *TtmaccountSession) AddServiceCapability(serviceHash [32]byte, capability string) (*types.Transaction, error) {
	return _Ttmaccount.Contract.AddServiceCapability(&_Ttmaccount.TransactOpts, serviceHash, capability)
}

// AddServiceCapability is a paid mutator transaction binding the contract method 0x5ae733c9.
//
// Solidity: function addServiceCapability(bytes32 serviceHash, string capability) returns()
func (_Ttmaccount *TtmaccountTransactorSession) AddServiceCapability(serviceHash [32]byte, capability string) (*types.Transaction, error) {
	return _Ttmaccount.Contract.AddServiceCapability(&_Ttmaccount.TransactOpts, serviceHash, capability)
}

// AddSupportedToken is a paid mutator transaction binding the contract method 0x6d69fcaf.
//
// Solidity: function addSupportedToken(address _supportedToken) returns()
func (_Ttmaccount *TtmaccountTransactor) AddSupportedToken(opts *bind.TransactOpts, _supportedToken common.Address) (*types.Transaction, error) {
	return _Ttmaccount.contract.Transact(opts, "addSupportedToken", _supportedToken)
}

// AddSupportedToken is a paid mutator transaction binding the contract method 0x6d69fcaf.
//
// Solidity: function addSupportedToken(address _supportedToken) returns()
func (_Ttmaccount *TtmaccountSession) AddSupportedToken(_supportedToken common.Address) (*types.Transaction, error) {
	return _Ttmaccount.Contract.AddSupportedToken(&_Ttmaccount.TransactOpts, _supportedToken)
}

// AddSupportedToken is a paid mutator transaction binding the contract method 0x6d69fcaf.
//
// Solidity: function addSupportedToken(address _supportedToken) returns()
func (_Ttmaccount *TtmaccountTransactorSession) AddSupportedToken(_supportedToken common.Address) (*types.Transaction, error) {
	return _Ttmaccount.Contract.AddSupportedToken(&_Ttmaccount.TransactOpts, _supportedToken)
}

// AddWantedServices is a paid mutator transaction binding the contract method 0x04a3d81e.
//
// Solidity: function addWantedServices(bytes32[] serviceHashes) returns()
func (_Ttmaccount *TtmaccountTransactor) AddWantedServices(opts *bind.TransactOpts, serviceHashes [][32]byte) (*types.Transaction, error) {
	return _Ttmaccount.contract.Transact(opts, "addWantedServices", serviceHashes)
}

// AddWantedServices is a paid mutator transaction binding the contract method 0x04a3d81e.
//
// Solidity: function addWantedServices(bytes32[] serviceHashes) returns()
func (_Ttmaccount *TtmaccountSession) AddWantedServices(serviceHashes [][32]byte) (*types.Transaction, error) {
	return _Ttmaccount.Contract.AddWantedServices(&_Ttmaccount.TransactOpts, serviceHashes)
}

// AddWantedServices is a paid mutator transaction binding the contract method 0x04a3d81e.
//
// Solidity: function addWantedServices(bytes32[] serviceHashes) returns()
func (_Ttmaccount *TtmaccountTransactorSession) AddWantedServices(serviceHashes [][32]byte) (*types.Transaction, error) {
	return _Ttmaccount.Contract.AddWantedServices(&_Ttmaccount.TransactOpts, serviceHashes)
}

// ApproveERC721 is a paid mutator transaction binding the contract method 0x00a7230a.
//
// Solidity: function approveERC721(address token, address to, uint256 tokenId) returns()
func (_Ttmaccount *TtmaccountTransactor) ApproveERC721(opts *bind.TransactOpts, token common.Address, to common.Address, tokenId *big.Int) (*types.Transaction, error) {
	return _Ttmaccount.contract.Transact(opts, "approveERC721", token, to, tokenId)
}

// ApproveERC721 is a paid mutator transaction binding the contract method 0x00a7230a.
//
// Solidity: function approveERC721(address token, address to, uint256 tokenId) returns()
func (_Ttmaccount *TtmaccountSession) ApproveERC721(token common.Address, to common.Address, tokenId *big.Int) (*types.Transaction, error) {
	return _Ttmaccount.Contract.ApproveERC721(&_Ttmaccount.TransactOpts, token, to, tokenId)
}

// ApproveERC721 is a paid mutator transaction binding the contract method 0x00a7230a.
//
// Solidity: function approveERC721(address token, address to, uint256 tokenId) returns()
func (_Ttmaccount *TtmaccountTransactorSession) ApproveERC721(token common.Address, to common.Address, tokenId *big.Int) (*types.Transaction, error) {
	return _Ttmaccount.Contract.ApproveERC721(&_Ttmaccount.TransactOpts, token, to, tokenId)
}

// BuyBookingToken is a paid mutator transaction binding the contract method 0xcd9ef914.
//
// Solidity: function buyBookingToken(uint256 tokenId, uint256 expectedPrice, address expectedPaymentToken) returns()
func (_Ttmaccount *TtmaccountTransactor) BuyBookingToken(opts *bind.TransactOpts, tokenId *big.Int, expectedPrice *big.Int, expectedPaymentToken common.Address) (*types.Transaction, error) {
	return _Ttmaccount.contract.Transact(opts, "buyBookingToken", tokenId, expectedPrice, expectedPaymentToken)
}

// BuyBookingToken is a paid mutator transaction binding the contract method 0xcd9ef914.
//
// Solidity: function buyBookingToken(uint256 tokenId, uint256 expectedPrice, address expectedPaymentToken) returns()
func (_Ttmaccount *TtmaccountSession) BuyBookingToken(tokenId *big.Int, expectedPrice *big.Int, expectedPaymentToken common.Address) (*types.Transaction, error) {
	return _Ttmaccount.Contract.BuyBookingToken(&_Ttmaccount.TransactOpts, tokenId, expectedPrice, expectedPaymentToken)
}

// BuyBookingToken is a paid mutator transaction binding the contract method 0xcd9ef914.
//
// Solidity: function buyBookingToken(uint256 tokenId, uint256 expectedPrice, address expectedPaymentToken) returns()
func (_Ttmaccount *TtmaccountTransactorSession) BuyBookingToken(tokenId *big.Int, expectedPrice *big.Int, expectedPaymentToken common.Address) (*types.Transaction, error) {
	return _Ttmaccount.Contract.BuyBookingToken(&_Ttmaccount.TransactOpts, tokenId, expectedPrice, expectedPaymentToken)
}

// CounterCancellation is a paid mutator transaction binding the contract method 0x74aa2048.
//
// Solidity: function counterCancellation(uint256 tokenId, uint256 refundAmount, uint16 counterReason, uint16 counterReasonVersion) returns()
func (_Ttmaccount *TtmaccountTransactor) CounterCancellation(opts *bind.TransactOpts, tokenId *big.Int, refundAmount *big.Int, counterReason uint16, counterReasonVersion uint16) (*types.Transaction, error) {
	return _Ttmaccount.contract.Transact(opts, "counterCancellation", tokenId, refundAmount, counterReason, counterReasonVersion)
}

// CounterCancellation is a paid mutator transaction binding the contract method 0x74aa2048.
//
// Solidity: function counterCancellation(uint256 tokenId, uint256 refundAmount, uint16 counterReason, uint16 counterReasonVersion) returns()
func (_Ttmaccount *TtmaccountSession) CounterCancellation(tokenId *big.Int, refundAmount *big.Int, counterReason uint16, counterReasonVersion uint16) (*types.Transaction, error) {
	return _Ttmaccount.Contract.CounterCancellation(&_Ttmaccount.TransactOpts, tokenId, refundAmount, counterReason, counterReasonVersion)
}

// CounterCancellation is a paid mutator transaction binding the contract method 0x74aa2048.
//
// Solidity: function counterCancellation(uint256 tokenId, uint256 refundAmount, uint16 counterReason, uint16 counterReasonVersion) returns()
func (_Ttmaccount *TtmaccountTransactorSession) CounterCancellation(tokenId *big.Int, refundAmount *big.Int, counterReason uint16, counterReasonVersion uint16) (*types.Transaction, error) {
	return _Ttmaccount.Contract.CounterCancellation(&_Ttmaccount.TransactOpts, tokenId, refundAmount, counterReason, counterReasonVersion)
}

// FinalizeCancellation is a paid mutator transaction binding the contract method 0x1c54f0f7.
//
// Solidity: function finalizeCancellation(uint256 tokenId, uint256 refundAmount) returns()
func (_Ttmaccount *TtmaccountTransactor) FinalizeCancellation(opts *bind.TransactOpts, tokenId *big.Int, refundAmount *big.Int) (*types.Transaction, error) {
	return _Ttmaccount.contract.Transact(opts, "finalizeCancellation", tokenId, refundAmount)
}

// FinalizeCancellation is a paid mutator transaction binding the contract method 0x1c54f0f7.
//
// Solidity: function finalizeCancellation(uint256 tokenId, uint256 refundAmount) returns()
func (_Ttmaccount *TtmaccountSession) FinalizeCancellation(tokenId *big.Int, refundAmount *big.Int) (*types.Transaction, error) {
	return _Ttmaccount.Contract.FinalizeCancellation(&_Ttmaccount.TransactOpts, tokenId, refundAmount)
}

// FinalizeCancellation is a paid mutator transaction binding the contract method 0x1c54f0f7.
//
// Solidity: function finalizeCancellation(uint256 tokenId, uint256 refundAmount) returns()
func (_Ttmaccount *TtmaccountTransactorSession) FinalizeCancellation(tokenId *big.Int, refundAmount *big.Int) (*types.Transaction, error) {
	return _Ttmaccount.Contract.FinalizeCancellation(&_Ttmaccount.TransactOpts, tokenId, refundAmount)
}

// GrantRole is a paid mutator transaction binding the contract method 0x2f2ff15d.
//
// Solidity: function grantRole(bytes32 role, address account) returns()
func (_Ttmaccount *TtmaccountTransactor) GrantRole(opts *bind.TransactOpts, role [32]byte, account common.Address) (*types.Transaction, error) {
	return _Ttmaccount.contract.Transact(opts, "grantRole", role, account)
}

// GrantRole is a paid mutator transaction binding the contract method 0x2f2ff15d.
//
// Solidity: function grantRole(bytes32 role, address account) returns()
func (_Ttmaccount *TtmaccountSession) GrantRole(role [32]byte, account common.Address) (*types.Transaction, error) {
	return _Ttmaccount.Contract.GrantRole(&_Ttmaccount.TransactOpts, role, account)
}

// GrantRole is a paid mutator transaction binding the contract method 0x2f2ff15d.
//
// Solidity: function grantRole(bytes32 role, address account) returns()
func (_Ttmaccount *TtmaccountTransactorSession) GrantRole(role [32]byte, account common.Address) (*types.Transaction, error) {
	return _Ttmaccount.Contract.GrantRole(&_Ttmaccount.TransactOpts, role, account)
}

// Initialize is a paid mutator transaction binding the contract method 0xf8c8765e.
//
// Solidity: function initialize(address manager, address bookingToken, address defaultAdmin, address upgrader) returns()
func (_Ttmaccount *TtmaccountTransactor) Initialize(opts *bind.TransactOpts, manager common.Address, bookingToken common.Address, defaultAdmin common.Address, upgrader common.Address) (*types.Transaction, error) {
	return _Ttmaccount.contract.Transact(opts, "initialize", manager, bookingToken, defaultAdmin, upgrader)
}

// Initialize is a paid mutator transaction binding the contract method 0xf8c8765e.
//
// Solidity: function initialize(address manager, address bookingToken, address defaultAdmin, address upgrader) returns()
func (_Ttmaccount *TtmaccountSession) Initialize(manager common.Address, bookingToken common.Address, defaultAdmin common.Address, upgrader common.Address) (*types.Transaction, error) {
	return _Ttmaccount.Contract.Initialize(&_Ttmaccount.TransactOpts, manager, bookingToken, defaultAdmin, upgrader)
}

// Initialize is a paid mutator transaction binding the contract method 0xf8c8765e.
//
// Solidity: function initialize(address manager, address bookingToken, address defaultAdmin, address upgrader) returns()
func (_Ttmaccount *TtmaccountTransactorSession) Initialize(manager common.Address, bookingToken common.Address, defaultAdmin common.Address, upgrader common.Address) (*types.Transaction, error) {
	return _Ttmaccount.Contract.Initialize(&_Ttmaccount.TransactOpts, manager, bookingToken, defaultAdmin, upgrader)
}

// InitiateCancellation is a paid mutator transaction binding the contract method 0xf7e45f09.
//
// Solidity: function initiateCancellation(uint256 tokenId, uint256 refundAmount, uint16 cancellationReason, uint16 cancellationReasonVersion) returns()
func (_Ttmaccount *TtmaccountTransactor) InitiateCancellation(opts *bind.TransactOpts, tokenId *big.Int, refundAmount *big.Int, cancellationReason uint16, cancellationReasonVersion uint16) (*types.Transaction, error) {
	return _Ttmaccount.contract.Transact(opts, "initiateCancellation", tokenId, refundAmount, cancellationReason, cancellationReasonVersion)
}

// InitiateCancellation is a paid mutator transaction binding the contract method 0xf7e45f09.
//
// Solidity: function initiateCancellation(uint256 tokenId, uint256 refundAmount, uint16 cancellationReason, uint16 cancellationReasonVersion) returns()
func (_Ttmaccount *TtmaccountSession) InitiateCancellation(tokenId *big.Int, refundAmount *big.Int, cancellationReason uint16, cancellationReasonVersion uint16) (*types.Transaction, error) {
	return _Ttmaccount.Contract.InitiateCancellation(&_Ttmaccount.TransactOpts, tokenId, refundAmount, cancellationReason, cancellationReasonVersion)
}

// InitiateCancellation is a paid mutator transaction binding the contract method 0xf7e45f09.
//
// Solidity: function initiateCancellation(uint256 tokenId, uint256 refundAmount, uint16 cancellationReason, uint16 cancellationReasonVersion) returns()
func (_Ttmaccount *TtmaccountTransactorSession) InitiateCancellation(tokenId *big.Int, refundAmount *big.Int, cancellationReason uint16, cancellationReasonVersion uint16) (*types.Transaction, error) {
	return _Ttmaccount.Contract.InitiateCancellation(&_Ttmaccount.TransactOpts, tokenId, refundAmount, cancellationReason, cancellationReasonVersion)
}

// MintBookingToken is a paid mutator transaction binding the contract method 0xe26a61bb.
//
// Solidity: function mintBookingToken(address reservedFor, string uri, uint256 expirationTimestamp, uint256 price, address paymentToken, uint256 offchainPaymentCurrency, bool cancellable) returns()
func (_Ttmaccount *TtmaccountTransactor) MintBookingToken(opts *bind.TransactOpts, reservedFor common.Address, uri string, expirationTimestamp *big.Int, price *big.Int, paymentToken common.Address, offchainPaymentCurrency *big.Int, cancellable bool) (*types.Transaction, error) {
	return _Ttmaccount.contract.Transact(opts, "mintBookingToken", reservedFor, uri, expirationTimestamp, price, paymentToken, offchainPaymentCurrency, cancellable)
}

// MintBookingToken is a paid mutator transaction binding the contract method 0xe26a61bb.
//
// Solidity: function mintBookingToken(address reservedFor, string uri, uint256 expirationTimestamp, uint256 price, address paymentToken, uint256 offchainPaymentCurrency, bool cancellable) returns()
func (_Ttmaccount *TtmaccountSession) MintBookingToken(reservedFor common.Address, uri string, expirationTimestamp *big.Int, price *big.Int, paymentToken common.Address, offchainPaymentCurrency *big.Int, cancellable bool) (*types.Transaction, error) {
	return _Ttmaccount.Contract.MintBookingToken(&_Ttmaccount.TransactOpts, reservedFor, uri, expirationTimestamp, price, paymentToken, offchainPaymentCurrency, cancellable)
}

// MintBookingToken is a paid mutator transaction binding the contract method 0xe26a61bb.
//
// Solidity: function mintBookingToken(address reservedFor, string uri, uint256 expirationTimestamp, uint256 price, address paymentToken, uint256 offchainPaymentCurrency, bool cancellable) returns()
func (_Ttmaccount *TtmaccountTransactorSession) MintBookingToken(reservedFor common.Address, uri string, expirationTimestamp *big.Int, price *big.Int, paymentToken common.Address, offchainPaymentCurrency *big.Int, cancellable bool) (*types.Transaction, error) {
	return _Ttmaccount.Contract.MintBookingToken(&_Ttmaccount.TransactOpts, reservedFor, uri, expirationTimestamp, price, paymentToken, offchainPaymentCurrency, cancellable)
}

// OnERC721Received is a paid mutator transaction binding the contract method 0x150b7a02.
//
// Solidity: function onERC721Received(address , address , uint256 , bytes ) returns(bytes4)
func (_Ttmaccount *TtmaccountTransactor) OnERC721Received(opts *bind.TransactOpts, arg0 common.Address, arg1 common.Address, arg2 *big.Int, arg3 []byte) (*types.Transaction, error) {
	return _Ttmaccount.contract.Transact(opts, "onERC721Received", arg0, arg1, arg2, arg3)
}

// OnERC721Received is a paid mutator transaction binding the contract method 0x150b7a02.
//
// Solidity: function onERC721Received(address , address , uint256 , bytes ) returns(bytes4)
func (_Ttmaccount *TtmaccountSession) OnERC721Received(arg0 common.Address, arg1 common.Address, arg2 *big.Int, arg3 []byte) (*types.Transaction, error) {
	return _Ttmaccount.Contract.OnERC721Received(&_Ttmaccount.TransactOpts, arg0, arg1, arg2, arg3)
}

// OnERC721Received is a paid mutator transaction binding the contract method 0x150b7a02.
//
// Solidity: function onERC721Received(address , address , uint256 , bytes ) returns(bytes4)
func (_Ttmaccount *TtmaccountTransactorSession) OnERC721Received(arg0 common.Address, arg1 common.Address, arg2 *big.Int, arg3 []byte) (*types.Transaction, error) {
	return _Ttmaccount.Contract.OnERC721Received(&_Ttmaccount.TransactOpts, arg0, arg1, arg2, arg3)
}

// RecordExpiration is a paid mutator transaction binding the contract method 0xe5a6725c.
//
// Solidity: function recordExpiration(uint256 tokenId) returns()
func (_Ttmaccount *TtmaccountTransactor) RecordExpiration(opts *bind.TransactOpts, tokenId *big.Int) (*types.Transaction, error) {
	return _Ttmaccount.contract.Transact(opts, "recordExpiration", tokenId)
}

// RecordExpiration is a paid mutator transaction binding the contract method 0xe5a6725c.
//
// Solidity: function recordExpiration(uint256 tokenId) returns()
func (_Ttmaccount *TtmaccountSession) RecordExpiration(tokenId *big.Int) (*types.Transaction, error) {
	return _Ttmaccount.Contract.RecordExpiration(&_Ttmaccount.TransactOpts, tokenId)
}

// RecordExpiration is a paid mutator transaction binding the contract method 0xe5a6725c.
//
// Solidity: function recordExpiration(uint256 tokenId) returns()
func (_Ttmaccount *TtmaccountTransactorSession) RecordExpiration(tokenId *big.Int) (*types.Transaction, error) {
	return _Ttmaccount.Contract.RecordExpiration(&_Ttmaccount.TransactOpts, tokenId)
}

// RejectCancellation is a paid mutator transaction binding the contract method 0x74fe60e9.
//
// Solidity: function rejectCancellation(uint256 tokenId, uint16 rejectionReason, uint16 rejectionReasonVersion) returns()
func (_Ttmaccount *TtmaccountTransactor) RejectCancellation(opts *bind.TransactOpts, tokenId *big.Int, rejectionReason uint16, rejectionReasonVersion uint16) (*types.Transaction, error) {
	return _Ttmaccount.contract.Transact(opts, "rejectCancellation", tokenId, rejectionReason, rejectionReasonVersion)
}

// RejectCancellation is a paid mutator transaction binding the contract method 0x74fe60e9.
//
// Solidity: function rejectCancellation(uint256 tokenId, uint16 rejectionReason, uint16 rejectionReasonVersion) returns()
func (_Ttmaccount *TtmaccountSession) RejectCancellation(tokenId *big.Int, rejectionReason uint16, rejectionReasonVersion uint16) (*types.Transaction, error) {
	return _Ttmaccount.Contract.RejectCancellation(&_Ttmaccount.TransactOpts, tokenId, rejectionReason, rejectionReasonVersion)
}

// RejectCancellation is a paid mutator transaction binding the contract method 0x74fe60e9.
//
// Solidity: function rejectCancellation(uint256 tokenId, uint16 rejectionReason, uint16 rejectionReasonVersion) returns()
func (_Ttmaccount *TtmaccountTransactorSession) RejectCancellation(tokenId *big.Int, rejectionReason uint16, rejectionReasonVersion uint16) (*types.Transaction, error) {
	return _Ttmaccount.Contract.RejectCancellation(&_Ttmaccount.TransactOpts, tokenId, rejectionReason, rejectionReasonVersion)
}

// RemoveAllServices is a paid mutator transaction binding the contract method 0xb82923fb.
//
// Solidity: function removeAllServices() returns()
func (_Ttmaccount *TtmaccountTransactor) RemoveAllServices(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Ttmaccount.contract.Transact(opts, "removeAllServices")
}

// RemoveAllServices is a paid mutator transaction binding the contract method 0xb82923fb.
//
// Solidity: function removeAllServices() returns()
func (_Ttmaccount *TtmaccountSession) RemoveAllServices() (*types.Transaction, error) {
	return _Ttmaccount.Contract.RemoveAllServices(&_Ttmaccount.TransactOpts)
}

// RemoveAllServices is a paid mutator transaction binding the contract method 0xb82923fb.
//
// Solidity: function removeAllServices() returns()
func (_Ttmaccount *TtmaccountTransactorSession) RemoveAllServices() (*types.Transaction, error) {
	return _Ttmaccount.Contract.RemoveAllServices(&_Ttmaccount.TransactOpts)
}

// RemoveMessengerBot is a paid mutator transaction binding the contract method 0xc6640e68.
//
// Solidity: function removeMessengerBot(address bot) returns()
func (_Ttmaccount *TtmaccountTransactor) RemoveMessengerBot(opts *bind.TransactOpts, bot common.Address) (*types.Transaction, error) {
	return _Ttmaccount.contract.Transact(opts, "removeMessengerBot", bot)
}

// RemoveMessengerBot is a paid mutator transaction binding the contract method 0xc6640e68.
//
// Solidity: function removeMessengerBot(address bot) returns()
func (_Ttmaccount *TtmaccountSession) RemoveMessengerBot(bot common.Address) (*types.Transaction, error) {
	return _Ttmaccount.Contract.RemoveMessengerBot(&_Ttmaccount.TransactOpts, bot)
}

// RemoveMessengerBot is a paid mutator transaction binding the contract method 0xc6640e68.
//
// Solidity: function removeMessengerBot(address bot) returns()
func (_Ttmaccount *TtmaccountTransactorSession) RemoveMessengerBot(bot common.Address) (*types.Transaction, error) {
	return _Ttmaccount.Contract.RemoveMessengerBot(&_Ttmaccount.TransactOpts, bot)
}

// RemovePublicKey is a paid mutator transaction binding the contract method 0xe7bfce9a.
//
// Solidity: function removePublicKey(address pubKeyAddress) returns()
func (_Ttmaccount *TtmaccountTransactor) RemovePublicKey(opts *bind.TransactOpts, pubKeyAddress common.Address) (*types.Transaction, error) {
	return _Ttmaccount.contract.Transact(opts, "removePublicKey", pubKeyAddress)
}

// RemovePublicKey is a paid mutator transaction binding the contract method 0xe7bfce9a.
//
// Solidity: function removePublicKey(address pubKeyAddress) returns()
func (_Ttmaccount *TtmaccountSession) RemovePublicKey(pubKeyAddress common.Address) (*types.Transaction, error) {
	return _Ttmaccount.Contract.RemovePublicKey(&_Ttmaccount.TransactOpts, pubKeyAddress)
}

// RemovePublicKey is a paid mutator transaction binding the contract method 0xe7bfce9a.
//
// Solidity: function removePublicKey(address pubKeyAddress) returns()
func (_Ttmaccount *TtmaccountTransactorSession) RemovePublicKey(pubKeyAddress common.Address) (*types.Transaction, error) {
	return _Ttmaccount.Contract.RemovePublicKey(&_Ttmaccount.TransactOpts, pubKeyAddress)
}

// RemoveService is a paid mutator transaction binding the contract method 0xd3884c3f.
//
// Solidity: function removeService(bytes32 serviceHash) returns()
func (_Ttmaccount *TtmaccountTransactor) RemoveService(opts *bind.TransactOpts, serviceHash [32]byte) (*types.Transaction, error) {
	return _Ttmaccount.contract.Transact(opts, "removeService", serviceHash)
}

// RemoveService is a paid mutator transaction binding the contract method 0xd3884c3f.
//
// Solidity: function removeService(bytes32 serviceHash) returns()
func (_Ttmaccount *TtmaccountSession) RemoveService(serviceHash [32]byte) (*types.Transaction, error) {
	return _Ttmaccount.Contract.RemoveService(&_Ttmaccount.TransactOpts, serviceHash)
}

// RemoveService is a paid mutator transaction binding the contract method 0xd3884c3f.
//
// Solidity: function removeService(bytes32 serviceHash) returns()
func (_Ttmaccount *TtmaccountTransactorSession) RemoveService(serviceHash [32]byte) (*types.Transaction, error) {
	return _Ttmaccount.Contract.RemoveService(&_Ttmaccount.TransactOpts, serviceHash)
}

// RemoveServiceCapability is a paid mutator transaction binding the contract method 0xf6b51740.
//
// Solidity: function removeServiceCapability(bytes32 serviceHash, string capability) returns()
func (_Ttmaccount *TtmaccountTransactor) RemoveServiceCapability(opts *bind.TransactOpts, serviceHash [32]byte, capability string) (*types.Transaction, error) {
	return _Ttmaccount.contract.Transact(opts, "removeServiceCapability", serviceHash, capability)
}

// RemoveServiceCapability is a paid mutator transaction binding the contract method 0xf6b51740.
//
// Solidity: function removeServiceCapability(bytes32 serviceHash, string capability) returns()
func (_Ttmaccount *TtmaccountSession) RemoveServiceCapability(serviceHash [32]byte, capability string) (*types.Transaction, error) {
	return _Ttmaccount.Contract.RemoveServiceCapability(&_Ttmaccount.TransactOpts, serviceHash, capability)
}

// RemoveServiceCapability is a paid mutator transaction binding the contract method 0xf6b51740.
//
// Solidity: function removeServiceCapability(bytes32 serviceHash, string capability) returns()
func (_Ttmaccount *TtmaccountTransactorSession) RemoveServiceCapability(serviceHash [32]byte, capability string) (*types.Transaction, error) {
	return _Ttmaccount.Contract.RemoveServiceCapability(&_Ttmaccount.TransactOpts, serviceHash, capability)
}

// RemoveSupportedToken is a paid mutator transaction binding the contract method 0x76319190.
//
// Solidity: function removeSupportedToken(address _supportedToken) returns()
func (_Ttmaccount *TtmaccountTransactor) RemoveSupportedToken(opts *bind.TransactOpts, _supportedToken common.Address) (*types.Transaction, error) {
	return _Ttmaccount.contract.Transact(opts, "removeSupportedToken", _supportedToken)
}

// RemoveSupportedToken is a paid mutator transaction binding the contract method 0x76319190.
//
// Solidity: function removeSupportedToken(address _supportedToken) returns()
func (_Ttmaccount *TtmaccountSession) RemoveSupportedToken(_supportedToken common.Address) (*types.Transaction, error) {
	return _Ttmaccount.Contract.RemoveSupportedToken(&_Ttmaccount.TransactOpts, _supportedToken)
}

// RemoveSupportedToken is a paid mutator transaction binding the contract method 0x76319190.
//
// Solidity: function removeSupportedToken(address _supportedToken) returns()
func (_Ttmaccount *TtmaccountTransactorSession) RemoveSupportedToken(_supportedToken common.Address) (*types.Transaction, error) {
	return _Ttmaccount.Contract.RemoveSupportedToken(&_Ttmaccount.TransactOpts, _supportedToken)
}

// RemoveWantedServices is a paid mutator transaction binding the contract method 0x82010fb1.
//
// Solidity: function removeWantedServices(bytes32[] serviceHashes) returns()
func (_Ttmaccount *TtmaccountTransactor) RemoveWantedServices(opts *bind.TransactOpts, serviceHashes [][32]byte) (*types.Transaction, error) {
	return _Ttmaccount.contract.Transact(opts, "removeWantedServices", serviceHashes)
}

// RemoveWantedServices is a paid mutator transaction binding the contract method 0x82010fb1.
//
// Solidity: function removeWantedServices(bytes32[] serviceHashes) returns()
func (_Ttmaccount *TtmaccountSession) RemoveWantedServices(serviceHashes [][32]byte) (*types.Transaction, error) {
	return _Ttmaccount.Contract.RemoveWantedServices(&_Ttmaccount.TransactOpts, serviceHashes)
}

// RemoveWantedServices is a paid mutator transaction binding the contract method 0x82010fb1.
//
// Solidity: function removeWantedServices(bytes32[] serviceHashes) returns()
func (_Ttmaccount *TtmaccountTransactorSession) RemoveWantedServices(serviceHashes [][32]byte) (*types.Transaction, error) {
	return _Ttmaccount.Contract.RemoveWantedServices(&_Ttmaccount.TransactOpts, serviceHashes)
}

// RenounceRole is a paid mutator transaction binding the contract method 0x36568abe.
//
// Solidity: function renounceRole(bytes32 role, address callerConfirmation) returns()
func (_Ttmaccount *TtmaccountTransactor) RenounceRole(opts *bind.TransactOpts, role [32]byte, callerConfirmation common.Address) (*types.Transaction, error) {
	return _Ttmaccount.contract.Transact(opts, "renounceRole", role, callerConfirmation)
}

// RenounceRole is a paid mutator transaction binding the contract method 0x36568abe.
//
// Solidity: function renounceRole(bytes32 role, address callerConfirmation) returns()
func (_Ttmaccount *TtmaccountSession) RenounceRole(role [32]byte, callerConfirmation common.Address) (*types.Transaction, error) {
	return _Ttmaccount.Contract.RenounceRole(&_Ttmaccount.TransactOpts, role, callerConfirmation)
}

// RenounceRole is a paid mutator transaction binding the contract method 0x36568abe.
//
// Solidity: function renounceRole(bytes32 role, address callerConfirmation) returns()
func (_Ttmaccount *TtmaccountTransactorSession) RenounceRole(role [32]byte, callerConfirmation common.Address) (*types.Transaction, error) {
	return _Ttmaccount.Contract.RenounceRole(&_Ttmaccount.TransactOpts, role, callerConfirmation)
}

// RevokeRole is a paid mutator transaction binding the contract method 0xd547741f.
//
// Solidity: function revokeRole(bytes32 role, address account) returns()
func (_Ttmaccount *TtmaccountTransactor) RevokeRole(opts *bind.TransactOpts, role [32]byte, account common.Address) (*types.Transaction, error) {
	return _Ttmaccount.contract.Transact(opts, "revokeRole", role, account)
}

// RevokeRole is a paid mutator transaction binding the contract method 0xd547741f.
//
// Solidity: function revokeRole(bytes32 role, address account) returns()
func (_Ttmaccount *TtmaccountSession) RevokeRole(role [32]byte, account common.Address) (*types.Transaction, error) {
	return _Ttmaccount.Contract.RevokeRole(&_Ttmaccount.TransactOpts, role, account)
}

// RevokeRole is a paid mutator transaction binding the contract method 0xd547741f.
//
// Solidity: function revokeRole(bytes32 role, address account) returns()
func (_Ttmaccount *TtmaccountTransactorSession) RevokeRole(role [32]byte, account common.Address) (*types.Transaction, error) {
	return _Ttmaccount.Contract.RevokeRole(&_Ttmaccount.TransactOpts, role, account)
}

// SetGasMoneyWithdrawal is a paid mutator transaction binding the contract method 0x6fc22cd1.
//
// Solidity: function setGasMoneyWithdrawal(uint256 limit, uint256 period) returns()
func (_Ttmaccount *TtmaccountTransactor) SetGasMoneyWithdrawal(opts *bind.TransactOpts, limit *big.Int, period *big.Int) (*types.Transaction, error) {
	return _Ttmaccount.contract.Transact(opts, "setGasMoneyWithdrawal", limit, period)
}

// SetGasMoneyWithdrawal is a paid mutator transaction binding the contract method 0x6fc22cd1.
//
// Solidity: function setGasMoneyWithdrawal(uint256 limit, uint256 period) returns()
func (_Ttmaccount *TtmaccountSession) SetGasMoneyWithdrawal(limit *big.Int, period *big.Int) (*types.Transaction, error) {
	return _Ttmaccount.Contract.SetGasMoneyWithdrawal(&_Ttmaccount.TransactOpts, limit, period)
}

// SetGasMoneyWithdrawal is a paid mutator transaction binding the contract method 0x6fc22cd1.
//
// Solidity: function setGasMoneyWithdrawal(uint256 limit, uint256 period) returns()
func (_Ttmaccount *TtmaccountTransactorSession) SetGasMoneyWithdrawal(limit *big.Int, period *big.Int) (*types.Transaction, error) {
	return _Ttmaccount.Contract.SetGasMoneyWithdrawal(&_Ttmaccount.TransactOpts, limit, period)
}

// SetServiceCapabilities is a paid mutator transaction binding the contract method 0xf8e191ac.
//
// Solidity: function setServiceCapabilities(bytes32 serviceHash, string[] capabilities) returns()
func (_Ttmaccount *TtmaccountTransactor) SetServiceCapabilities(opts *bind.TransactOpts, serviceHash [32]byte, capabilities []string) (*types.Transaction, error) {
	return _Ttmaccount.contract.Transact(opts, "setServiceCapabilities", serviceHash, capabilities)
}

// SetServiceCapabilities is a paid mutator transaction binding the contract method 0xf8e191ac.
//
// Solidity: function setServiceCapabilities(bytes32 serviceHash, string[] capabilities) returns()
func (_Ttmaccount *TtmaccountSession) SetServiceCapabilities(serviceHash [32]byte, capabilities []string) (*types.Transaction, error) {
	return _Ttmaccount.Contract.SetServiceCapabilities(&_Ttmaccount.TransactOpts, serviceHash, capabilities)
}

// SetServiceCapabilities is a paid mutator transaction binding the contract method 0xf8e191ac.
//
// Solidity: function setServiceCapabilities(bytes32 serviceHash, string[] capabilities) returns()
func (_Ttmaccount *TtmaccountTransactorSession) SetServiceCapabilities(serviceHash [32]byte, capabilities []string) (*types.Transaction, error) {
	return _Ttmaccount.Contract.SetServiceCapabilities(&_Ttmaccount.TransactOpts, serviceHash, capabilities)
}

// SetServiceRestrictedRate is a paid mutator transaction binding the contract method 0x0d5e2e16.
//
// Solidity: function setServiceRestrictedRate(bytes32 serviceHash, bool restrictedRate) returns()
func (_Ttmaccount *TtmaccountTransactor) SetServiceRestrictedRate(opts *bind.TransactOpts, serviceHash [32]byte, restrictedRate bool) (*types.Transaction, error) {
	return _Ttmaccount.contract.Transact(opts, "setServiceRestrictedRate", serviceHash, restrictedRate)
}

// SetServiceRestrictedRate is a paid mutator transaction binding the contract method 0x0d5e2e16.
//
// Solidity: function setServiceRestrictedRate(bytes32 serviceHash, bool restrictedRate) returns()
func (_Ttmaccount *TtmaccountSession) SetServiceRestrictedRate(serviceHash [32]byte, restrictedRate bool) (*types.Transaction, error) {
	return _Ttmaccount.Contract.SetServiceRestrictedRate(&_Ttmaccount.TransactOpts, serviceHash, restrictedRate)
}

// SetServiceRestrictedRate is a paid mutator transaction binding the contract method 0x0d5e2e16.
//
// Solidity: function setServiceRestrictedRate(bytes32 serviceHash, bool restrictedRate) returns()
func (_Ttmaccount *TtmaccountTransactorSession) SetServiceRestrictedRate(serviceHash [32]byte, restrictedRate bool) (*types.Transaction, error) {
	return _Ttmaccount.Contract.SetServiceRestrictedRate(&_Ttmaccount.TransactOpts, serviceHash, restrictedRate)
}

// TransferERC20 is a paid mutator transaction binding the contract method 0x9db5dbe4.
//
// Solidity: function transferERC20(address token, address to, uint256 amount) returns()
func (_Ttmaccount *TtmaccountTransactor) TransferERC20(opts *bind.TransactOpts, token common.Address, to common.Address, amount *big.Int) (*types.Transaction, error) {
	return _Ttmaccount.contract.Transact(opts, "transferERC20", token, to, amount)
}

// TransferERC20 is a paid mutator transaction binding the contract method 0x9db5dbe4.
//
// Solidity: function transferERC20(address token, address to, uint256 amount) returns()
func (_Ttmaccount *TtmaccountSession) TransferERC20(token common.Address, to common.Address, amount *big.Int) (*types.Transaction, error) {
	return _Ttmaccount.Contract.TransferERC20(&_Ttmaccount.TransactOpts, token, to, amount)
}

// TransferERC20 is a paid mutator transaction binding the contract method 0x9db5dbe4.
//
// Solidity: function transferERC20(address token, address to, uint256 amount) returns()
func (_Ttmaccount *TtmaccountTransactorSession) TransferERC20(token common.Address, to common.Address, amount *big.Int) (*types.Transaction, error) {
	return _Ttmaccount.Contract.TransferERC20(&_Ttmaccount.TransactOpts, token, to, amount)
}

// TransferERC721 is a paid mutator transaction binding the contract method 0x1aca6376.
//
// Solidity: function transferERC721(address token, address to, uint256 tokenId) returns()
func (_Ttmaccount *TtmaccountTransactor) TransferERC721(opts *bind.TransactOpts, token common.Address, to common.Address, tokenId *big.Int) (*types.Transaction, error) {
	return _Ttmaccount.contract.Transact(opts, "transferERC721", token, to, tokenId)
}

// TransferERC721 is a paid mutator transaction binding the contract method 0x1aca6376.
//
// Solidity: function transferERC721(address token, address to, uint256 tokenId) returns()
func (_Ttmaccount *TtmaccountSession) TransferERC721(token common.Address, to common.Address, tokenId *big.Int) (*types.Transaction, error) {
	return _Ttmaccount.Contract.TransferERC721(&_Ttmaccount.TransactOpts, token, to, tokenId)
}

// TransferERC721 is a paid mutator transaction binding the contract method 0x1aca6376.
//
// Solidity: function transferERC721(address token, address to, uint256 tokenId) returns()
func (_Ttmaccount *TtmaccountTransactorSession) TransferERC721(token common.Address, to common.Address, tokenId *big.Int) (*types.Transaction, error) {
	return _Ttmaccount.Contract.TransferERC721(&_Ttmaccount.TransactOpts, token, to, tokenId)
}

// UpgradeToAndCall is a paid mutator transaction binding the contract method 0x4f1ef286.
//
// Solidity: function upgradeToAndCall(address newImplementation, bytes data) payable returns()
func (_Ttmaccount *TtmaccountTransactor) UpgradeToAndCall(opts *bind.TransactOpts, newImplementation common.Address, data []byte) (*types.Transaction, error) {
	return _Ttmaccount.contract.Transact(opts, "upgradeToAndCall", newImplementation, data)
}

// UpgradeToAndCall is a paid mutator transaction binding the contract method 0x4f1ef286.
//
// Solidity: function upgradeToAndCall(address newImplementation, bytes data) payable returns()
func (_Ttmaccount *TtmaccountSession) UpgradeToAndCall(newImplementation common.Address, data []byte) (*types.Transaction, error) {
	return _Ttmaccount.Contract.UpgradeToAndCall(&_Ttmaccount.TransactOpts, newImplementation, data)
}

// UpgradeToAndCall is a paid mutator transaction binding the contract method 0x4f1ef286.
//
// Solidity: function upgradeToAndCall(address newImplementation, bytes data) payable returns()
func (_Ttmaccount *TtmaccountTransactorSession) UpgradeToAndCall(newImplementation common.Address, data []byte) (*types.Transaction, error) {
	return _Ttmaccount.Contract.UpgradeToAndCall(&_Ttmaccount.TransactOpts, newImplementation, data)
}

// Withdraw is a paid mutator transaction binding the contract method 0xf3fef3a3.
//
// Solidity: function withdraw(address recipient, uint256 amount) returns()
func (_Ttmaccount *TtmaccountTransactor) Withdraw(opts *bind.TransactOpts, recipient common.Address, amount *big.Int) (*types.Transaction, error) {
	return _Ttmaccount.contract.Transact(opts, "withdraw", recipient, amount)
}

// Withdraw is a paid mutator transaction binding the contract method 0xf3fef3a3.
//
// Solidity: function withdraw(address recipient, uint256 amount) returns()
func (_Ttmaccount *TtmaccountSession) Withdraw(recipient common.Address, amount *big.Int) (*types.Transaction, error) {
	return _Ttmaccount.Contract.Withdraw(&_Ttmaccount.TransactOpts, recipient, amount)
}

// Withdraw is a paid mutator transaction binding the contract method 0xf3fef3a3.
//
// Solidity: function withdraw(address recipient, uint256 amount) returns()
func (_Ttmaccount *TtmaccountTransactorSession) Withdraw(recipient common.Address, amount *big.Int) (*types.Transaction, error) {
	return _Ttmaccount.Contract.Withdraw(&_Ttmaccount.TransactOpts, recipient, amount)
}

// WithdrawCancellation is a paid mutator transaction binding the contract method 0x2a119380.
//
// Solidity: function withdrawCancellation(uint256 tokenId, uint16 reason, uint16 reasonVersion) returns()
func (_Ttmaccount *TtmaccountTransactor) WithdrawCancellation(opts *bind.TransactOpts, tokenId *big.Int, reason uint16, reasonVersion uint16) (*types.Transaction, error) {
	return _Ttmaccount.contract.Transact(opts, "withdrawCancellation", tokenId, reason, reasonVersion)
}

// WithdrawCancellation is a paid mutator transaction binding the contract method 0x2a119380.
//
// Solidity: function withdrawCancellation(uint256 tokenId, uint16 reason, uint16 reasonVersion) returns()
func (_Ttmaccount *TtmaccountSession) WithdrawCancellation(tokenId *big.Int, reason uint16, reasonVersion uint16) (*types.Transaction, error) {
	return _Ttmaccount.Contract.WithdrawCancellation(&_Ttmaccount.TransactOpts, tokenId, reason, reasonVersion)
}

// WithdrawCancellation is a paid mutator transaction binding the contract method 0x2a119380.
//
// Solidity: function withdrawCancellation(uint256 tokenId, uint16 reason, uint16 reasonVersion) returns()
func (_Ttmaccount *TtmaccountTransactorSession) WithdrawCancellation(tokenId *big.Int, reason uint16, reasonVersion uint16) (*types.Transaction, error) {
	return _Ttmaccount.Contract.WithdrawCancellation(&_Ttmaccount.TransactOpts, tokenId, reason, reasonVersion)
}

// WithdrawGasMoney is a paid mutator transaction binding the contract method 0x5c988994.
//
// Solidity: function withdrawGasMoney(uint256 amount) returns()
func (_Ttmaccount *TtmaccountTransactor) WithdrawGasMoney(opts *bind.TransactOpts, amount *big.Int) (*types.Transaction, error) {
	return _Ttmaccount.contract.Transact(opts, "withdrawGasMoney", amount)
}

// WithdrawGasMoney is a paid mutator transaction binding the contract method 0x5c988994.
//
// Solidity: function withdrawGasMoney(uint256 amount) returns()
func (_Ttmaccount *TtmaccountSession) WithdrawGasMoney(amount *big.Int) (*types.Transaction, error) {
	return _Ttmaccount.Contract.WithdrawGasMoney(&_Ttmaccount.TransactOpts, amount)
}

// WithdrawGasMoney is a paid mutator transaction binding the contract method 0x5c988994.
//
// Solidity: function withdrawGasMoney(uint256 amount) returns()
func (_Ttmaccount *TtmaccountTransactorSession) WithdrawGasMoney(amount *big.Int) (*types.Transaction, error) {
	return _Ttmaccount.Contract.WithdrawGasMoney(&_Ttmaccount.TransactOpts, amount)
}

// Receive is a paid mutator transaction binding the contract receive function.
//
// Solidity: receive() payable returns()
func (_Ttmaccount *TtmaccountTransactor) Receive(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Ttmaccount.contract.RawTransact(opts, nil) // calldata is disallowed for receive function
}

// Receive is a paid mutator transaction binding the contract receive function.
//
// Solidity: receive() payable returns()
func (_Ttmaccount *TtmaccountSession) Receive() (*types.Transaction, error) {
	return _Ttmaccount.Contract.Receive(&_Ttmaccount.TransactOpts)
}

// Receive is a paid mutator transaction binding the contract receive function.
//
// Solidity: receive() payable returns()
func (_Ttmaccount *TtmaccountTransactorSession) Receive() (*types.Transaction, error) {
	return _Ttmaccount.Contract.Receive(&_Ttmaccount.TransactOpts)
}

// TtmaccountDepositIterator is returned from FilterDeposit and is used to iterate over the raw logs and unpacked data for Deposit events raised by the Ttmaccount contract.
type TtmaccountDepositIterator struct {
	Event *TtmaccountDeposit // Event containing the contract specifics and raw log

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
func (it *TtmaccountDepositIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(TtmaccountDeposit)
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
		it.Event = new(TtmaccountDeposit)
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
func (it *TtmaccountDepositIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *TtmaccountDepositIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// TtmaccountDeposit represents a Deposit event raised by the Ttmaccount contract.
type TtmaccountDeposit struct {
	Sender common.Address
	Amount *big.Int
	Raw    types.Log // Blockchain specific contextual infos
}

// FilterDeposit is a free log retrieval operation binding the contract event 0xe1fffcc4923d04b559f4d29a8bfc6cda04eb5b0d3c460751c2402c5c5cc9109c.
//
// Solidity: event Deposit(address indexed sender, uint256 amount)
func (_Ttmaccount *TtmaccountFilterer) FilterDeposit(opts *bind.FilterOpts, sender []common.Address) (*TtmaccountDepositIterator, error) {

	var senderRule []interface{}
	for _, senderItem := range sender {
		senderRule = append(senderRule, senderItem)
	}

	logs, sub, err := _Ttmaccount.contract.FilterLogs(opts, "Deposit", senderRule)
	if err != nil {
		return nil, err
	}
	return &TtmaccountDepositIterator{contract: _Ttmaccount.contract, event: "Deposit", logs: logs, sub: sub}, nil
}

// WatchDeposit is a free log subscription operation binding the contract event 0xe1fffcc4923d04b559f4d29a8bfc6cda04eb5b0d3c460751c2402c5c5cc9109c.
//
// Solidity: event Deposit(address indexed sender, uint256 amount)
func (_Ttmaccount *TtmaccountFilterer) WatchDeposit(opts *bind.WatchOpts, sink chan<- *TtmaccountDeposit, sender []common.Address) (event.Subscription, error) {

	var senderRule []interface{}
	for _, senderItem := range sender {
		senderRule = append(senderRule, senderItem)
	}

	logs, sub, err := _Ttmaccount.contract.WatchLogs(opts, "Deposit", senderRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(TtmaccountDeposit)
				if err := _Ttmaccount.contract.UnpackLog(event, "Deposit", log); err != nil {
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

// ParseDeposit is a log parse operation binding the contract event 0xe1fffcc4923d04b559f4d29a8bfc6cda04eb5b0d3c460751c2402c5c5cc9109c.
//
// Solidity: event Deposit(address indexed sender, uint256 amount)
func (_Ttmaccount *TtmaccountFilterer) ParseDeposit(log types.Log) (*TtmaccountDeposit, error) {
	event := new(TtmaccountDeposit)
	if err := _Ttmaccount.contract.UnpackLog(event, "Deposit", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// TtmaccountGasMoneyWithdrawalIterator is returned from FilterGasMoneyWithdrawal and is used to iterate over the raw logs and unpacked data for GasMoneyWithdrawal events raised by the Ttmaccount contract.
type TtmaccountGasMoneyWithdrawalIterator struct {
	Event *TtmaccountGasMoneyWithdrawal // Event containing the contract specifics and raw log

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
func (it *TtmaccountGasMoneyWithdrawalIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(TtmaccountGasMoneyWithdrawal)
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
		it.Event = new(TtmaccountGasMoneyWithdrawal)
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
func (it *TtmaccountGasMoneyWithdrawalIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *TtmaccountGasMoneyWithdrawalIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// TtmaccountGasMoneyWithdrawal represents a GasMoneyWithdrawal event raised by the Ttmaccount contract.
type TtmaccountGasMoneyWithdrawal struct {
	Withdrawer common.Address
	Amount     *big.Int
	Raw        types.Log // Blockchain specific contextual infos
}

// FilterGasMoneyWithdrawal is a free log retrieval operation binding the contract event 0xb9ec638398bbdcd0844ca414d8ce760939fa88b9258b9764b3fc6c12ea2605c2.
//
// Solidity: event GasMoneyWithdrawal(address indexed withdrawer, uint256 amount)
func (_Ttmaccount *TtmaccountFilterer) FilterGasMoneyWithdrawal(opts *bind.FilterOpts, withdrawer []common.Address) (*TtmaccountGasMoneyWithdrawalIterator, error) {

	var withdrawerRule []interface{}
	for _, withdrawerItem := range withdrawer {
		withdrawerRule = append(withdrawerRule, withdrawerItem)
	}

	logs, sub, err := _Ttmaccount.contract.FilterLogs(opts, "GasMoneyWithdrawal", withdrawerRule)
	if err != nil {
		return nil, err
	}
	return &TtmaccountGasMoneyWithdrawalIterator{contract: _Ttmaccount.contract, event: "GasMoneyWithdrawal", logs: logs, sub: sub}, nil
}

// WatchGasMoneyWithdrawal is a free log subscription operation binding the contract event 0xb9ec638398bbdcd0844ca414d8ce760939fa88b9258b9764b3fc6c12ea2605c2.
//
// Solidity: event GasMoneyWithdrawal(address indexed withdrawer, uint256 amount)
func (_Ttmaccount *TtmaccountFilterer) WatchGasMoneyWithdrawal(opts *bind.WatchOpts, sink chan<- *TtmaccountGasMoneyWithdrawal, withdrawer []common.Address) (event.Subscription, error) {

	var withdrawerRule []interface{}
	for _, withdrawerItem := range withdrawer {
		withdrawerRule = append(withdrawerRule, withdrawerItem)
	}

	logs, sub, err := _Ttmaccount.contract.WatchLogs(opts, "GasMoneyWithdrawal", withdrawerRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(TtmaccountGasMoneyWithdrawal)
				if err := _Ttmaccount.contract.UnpackLog(event, "GasMoneyWithdrawal", log); err != nil {
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

// ParseGasMoneyWithdrawal is a log parse operation binding the contract event 0xb9ec638398bbdcd0844ca414d8ce760939fa88b9258b9764b3fc6c12ea2605c2.
//
// Solidity: event GasMoneyWithdrawal(address indexed withdrawer, uint256 amount)
func (_Ttmaccount *TtmaccountFilterer) ParseGasMoneyWithdrawal(log types.Log) (*TtmaccountGasMoneyWithdrawal, error) {
	event := new(TtmaccountGasMoneyWithdrawal)
	if err := _Ttmaccount.contract.UnpackLog(event, "GasMoneyWithdrawal", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// TtmaccountGasMoneyWithdrawalUpdatedIterator is returned from FilterGasMoneyWithdrawalUpdated and is used to iterate over the raw logs and unpacked data for GasMoneyWithdrawalUpdated events raised by the Ttmaccount contract.
type TtmaccountGasMoneyWithdrawalUpdatedIterator struct {
	Event *TtmaccountGasMoneyWithdrawalUpdated // Event containing the contract specifics and raw log

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
func (it *TtmaccountGasMoneyWithdrawalUpdatedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(TtmaccountGasMoneyWithdrawalUpdated)
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
		it.Event = new(TtmaccountGasMoneyWithdrawalUpdated)
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
func (it *TtmaccountGasMoneyWithdrawalUpdatedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *TtmaccountGasMoneyWithdrawalUpdatedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// TtmaccountGasMoneyWithdrawalUpdated represents a GasMoneyWithdrawalUpdated event raised by the Ttmaccount contract.
type TtmaccountGasMoneyWithdrawalUpdated struct {
	Limit  *big.Int
	Period *big.Int
	Raw    types.Log // Blockchain specific contextual infos
}

// FilterGasMoneyWithdrawalUpdated is a free log retrieval operation binding the contract event 0x8d4925b196ae6b935035a27ed36c6bd9c7a8fbddc7a3f55f493aa8e230be373e.
//
// Solidity: event GasMoneyWithdrawalUpdated(uint256 limit, uint256 period)
func (_Ttmaccount *TtmaccountFilterer) FilterGasMoneyWithdrawalUpdated(opts *bind.FilterOpts) (*TtmaccountGasMoneyWithdrawalUpdatedIterator, error) {

	logs, sub, err := _Ttmaccount.contract.FilterLogs(opts, "GasMoneyWithdrawalUpdated")
	if err != nil {
		return nil, err
	}
	return &TtmaccountGasMoneyWithdrawalUpdatedIterator{contract: _Ttmaccount.contract, event: "GasMoneyWithdrawalUpdated", logs: logs, sub: sub}, nil
}

// WatchGasMoneyWithdrawalUpdated is a free log subscription operation binding the contract event 0x8d4925b196ae6b935035a27ed36c6bd9c7a8fbddc7a3f55f493aa8e230be373e.
//
// Solidity: event GasMoneyWithdrawalUpdated(uint256 limit, uint256 period)
func (_Ttmaccount *TtmaccountFilterer) WatchGasMoneyWithdrawalUpdated(opts *bind.WatchOpts, sink chan<- *TtmaccountGasMoneyWithdrawalUpdated) (event.Subscription, error) {

	logs, sub, err := _Ttmaccount.contract.WatchLogs(opts, "GasMoneyWithdrawalUpdated")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(TtmaccountGasMoneyWithdrawalUpdated)
				if err := _Ttmaccount.contract.UnpackLog(event, "GasMoneyWithdrawalUpdated", log); err != nil {
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

// ParseGasMoneyWithdrawalUpdated is a log parse operation binding the contract event 0x8d4925b196ae6b935035a27ed36c6bd9c7a8fbddc7a3f55f493aa8e230be373e.
//
// Solidity: event GasMoneyWithdrawalUpdated(uint256 limit, uint256 period)
func (_Ttmaccount *TtmaccountFilterer) ParseGasMoneyWithdrawalUpdated(log types.Log) (*TtmaccountGasMoneyWithdrawalUpdated, error) {
	event := new(TtmaccountGasMoneyWithdrawalUpdated)
	if err := _Ttmaccount.contract.UnpackLog(event, "GasMoneyWithdrawalUpdated", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// TtmaccountInitializedIterator is returned from FilterInitialized and is used to iterate over the raw logs and unpacked data for Initialized events raised by the Ttmaccount contract.
type TtmaccountInitializedIterator struct {
	Event *TtmaccountInitialized // Event containing the contract specifics and raw log

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
func (it *TtmaccountInitializedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(TtmaccountInitialized)
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
		it.Event = new(TtmaccountInitialized)
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
func (it *TtmaccountInitializedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *TtmaccountInitializedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// TtmaccountInitialized represents a Initialized event raised by the Ttmaccount contract.
type TtmaccountInitialized struct {
	Version uint64
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterInitialized is a free log retrieval operation binding the contract event 0xc7f505b2f371ae2175ee4913f4499e1f2633a7b5936321eed1cdaeb6115181d2.
//
// Solidity: event Initialized(uint64 version)
func (_Ttmaccount *TtmaccountFilterer) FilterInitialized(opts *bind.FilterOpts) (*TtmaccountInitializedIterator, error) {

	logs, sub, err := _Ttmaccount.contract.FilterLogs(opts, "Initialized")
	if err != nil {
		return nil, err
	}
	return &TtmaccountInitializedIterator{contract: _Ttmaccount.contract, event: "Initialized", logs: logs, sub: sub}, nil
}

// WatchInitialized is a free log subscription operation binding the contract event 0xc7f505b2f371ae2175ee4913f4499e1f2633a7b5936321eed1cdaeb6115181d2.
//
// Solidity: event Initialized(uint64 version)
func (_Ttmaccount *TtmaccountFilterer) WatchInitialized(opts *bind.WatchOpts, sink chan<- *TtmaccountInitialized) (event.Subscription, error) {

	logs, sub, err := _Ttmaccount.contract.WatchLogs(opts, "Initialized")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(TtmaccountInitialized)
				if err := _Ttmaccount.contract.UnpackLog(event, "Initialized", log); err != nil {
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

// ParseInitialized is a log parse operation binding the contract event 0xc7f505b2f371ae2175ee4913f4499e1f2633a7b5936321eed1cdaeb6115181d2.
//
// Solidity: event Initialized(uint64 version)
func (_Ttmaccount *TtmaccountFilterer) ParseInitialized(log types.Log) (*TtmaccountInitialized, error) {
	event := new(TtmaccountInitialized)
	if err := _Ttmaccount.contract.UnpackLog(event, "Initialized", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// TtmaccountMessengerBotAddedIterator is returned from FilterMessengerBotAdded and is used to iterate over the raw logs and unpacked data for MessengerBotAdded events raised by the Ttmaccount contract.
type TtmaccountMessengerBotAddedIterator struct {
	Event *TtmaccountMessengerBotAdded // Event containing the contract specifics and raw log

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
func (it *TtmaccountMessengerBotAddedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(TtmaccountMessengerBotAdded)
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
		it.Event = new(TtmaccountMessengerBotAdded)
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
func (it *TtmaccountMessengerBotAddedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *TtmaccountMessengerBotAddedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// TtmaccountMessengerBotAdded represents a MessengerBotAdded event raised by the Ttmaccount contract.
type TtmaccountMessengerBotAdded struct {
	Bot common.Address
	Raw types.Log // Blockchain specific contextual infos
}

// FilterMessengerBotAdded is a free log retrieval operation binding the contract event 0xdb3e11ba26e83d528bf96a2167061674c1ce7777c61376d852d172594a873994.
//
// Solidity: event MessengerBotAdded(address indexed bot)
func (_Ttmaccount *TtmaccountFilterer) FilterMessengerBotAdded(opts *bind.FilterOpts, bot []common.Address) (*TtmaccountMessengerBotAddedIterator, error) {

	var botRule []interface{}
	for _, botItem := range bot {
		botRule = append(botRule, botItem)
	}

	logs, sub, err := _Ttmaccount.contract.FilterLogs(opts, "MessengerBotAdded", botRule)
	if err != nil {
		return nil, err
	}
	return &TtmaccountMessengerBotAddedIterator{contract: _Ttmaccount.contract, event: "MessengerBotAdded", logs: logs, sub: sub}, nil
}

// WatchMessengerBotAdded is a free log subscription operation binding the contract event 0xdb3e11ba26e83d528bf96a2167061674c1ce7777c61376d852d172594a873994.
//
// Solidity: event MessengerBotAdded(address indexed bot)
func (_Ttmaccount *TtmaccountFilterer) WatchMessengerBotAdded(opts *bind.WatchOpts, sink chan<- *TtmaccountMessengerBotAdded, bot []common.Address) (event.Subscription, error) {

	var botRule []interface{}
	for _, botItem := range bot {
		botRule = append(botRule, botItem)
	}

	logs, sub, err := _Ttmaccount.contract.WatchLogs(opts, "MessengerBotAdded", botRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(TtmaccountMessengerBotAdded)
				if err := _Ttmaccount.contract.UnpackLog(event, "MessengerBotAdded", log); err != nil {
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

// ParseMessengerBotAdded is a log parse operation binding the contract event 0xdb3e11ba26e83d528bf96a2167061674c1ce7777c61376d852d172594a873994.
//
// Solidity: event MessengerBotAdded(address indexed bot)
func (_Ttmaccount *TtmaccountFilterer) ParseMessengerBotAdded(log types.Log) (*TtmaccountMessengerBotAdded, error) {
	event := new(TtmaccountMessengerBotAdded)
	if err := _Ttmaccount.contract.UnpackLog(event, "MessengerBotAdded", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// TtmaccountMessengerBotRemovedIterator is returned from FilterMessengerBotRemoved and is used to iterate over the raw logs and unpacked data for MessengerBotRemoved events raised by the Ttmaccount contract.
type TtmaccountMessengerBotRemovedIterator struct {
	Event *TtmaccountMessengerBotRemoved // Event containing the contract specifics and raw log

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
func (it *TtmaccountMessengerBotRemovedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(TtmaccountMessengerBotRemoved)
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
		it.Event = new(TtmaccountMessengerBotRemoved)
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
func (it *TtmaccountMessengerBotRemovedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *TtmaccountMessengerBotRemovedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// TtmaccountMessengerBotRemoved represents a MessengerBotRemoved event raised by the Ttmaccount contract.
type TtmaccountMessengerBotRemoved struct {
	Bot common.Address
	Raw types.Log // Blockchain specific contextual infos
}

// FilterMessengerBotRemoved is a free log retrieval operation binding the contract event 0xd124523a9cfa28c5dd01826c4fa56192ec7d56859943082e0ca46c3b9dc62913.
//
// Solidity: event MessengerBotRemoved(address indexed bot)
func (_Ttmaccount *TtmaccountFilterer) FilterMessengerBotRemoved(opts *bind.FilterOpts, bot []common.Address) (*TtmaccountMessengerBotRemovedIterator, error) {

	var botRule []interface{}
	for _, botItem := range bot {
		botRule = append(botRule, botItem)
	}

	logs, sub, err := _Ttmaccount.contract.FilterLogs(opts, "MessengerBotRemoved", botRule)
	if err != nil {
		return nil, err
	}
	return &TtmaccountMessengerBotRemovedIterator{contract: _Ttmaccount.contract, event: "MessengerBotRemoved", logs: logs, sub: sub}, nil
}

// WatchMessengerBotRemoved is a free log subscription operation binding the contract event 0xd124523a9cfa28c5dd01826c4fa56192ec7d56859943082e0ca46c3b9dc62913.
//
// Solidity: event MessengerBotRemoved(address indexed bot)
func (_Ttmaccount *TtmaccountFilterer) WatchMessengerBotRemoved(opts *bind.WatchOpts, sink chan<- *TtmaccountMessengerBotRemoved, bot []common.Address) (event.Subscription, error) {

	var botRule []interface{}
	for _, botItem := range bot {
		botRule = append(botRule, botItem)
	}

	logs, sub, err := _Ttmaccount.contract.WatchLogs(opts, "MessengerBotRemoved", botRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(TtmaccountMessengerBotRemoved)
				if err := _Ttmaccount.contract.UnpackLog(event, "MessengerBotRemoved", log); err != nil {
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

// ParseMessengerBotRemoved is a log parse operation binding the contract event 0xd124523a9cfa28c5dd01826c4fa56192ec7d56859943082e0ca46c3b9dc62913.
//
// Solidity: event MessengerBotRemoved(address indexed bot)
func (_Ttmaccount *TtmaccountFilterer) ParseMessengerBotRemoved(log types.Log) (*TtmaccountMessengerBotRemoved, error) {
	event := new(TtmaccountMessengerBotRemoved)
	if err := _Ttmaccount.contract.UnpackLog(event, "MessengerBotRemoved", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// TtmaccountPaymentTokenAddedIterator is returned from FilterPaymentTokenAdded and is used to iterate over the raw logs and unpacked data for PaymentTokenAdded events raised by the Ttmaccount contract.
type TtmaccountPaymentTokenAddedIterator struct {
	Event *TtmaccountPaymentTokenAdded // Event containing the contract specifics and raw log

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
func (it *TtmaccountPaymentTokenAddedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(TtmaccountPaymentTokenAdded)
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
		it.Event = new(TtmaccountPaymentTokenAdded)
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
func (it *TtmaccountPaymentTokenAddedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *TtmaccountPaymentTokenAddedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// TtmaccountPaymentTokenAdded represents a PaymentTokenAdded event raised by the Ttmaccount contract.
type TtmaccountPaymentTokenAdded struct {
	Token common.Address
	Raw   types.Log // Blockchain specific contextual infos
}

// FilterPaymentTokenAdded is a free log retrieval operation binding the contract event 0xa317c10673baf4f03b3c1041bd5ddbb537d0333a86fec3607c75f9dbb630f48f.
//
// Solidity: event PaymentTokenAdded(address indexed token)
func (_Ttmaccount *TtmaccountFilterer) FilterPaymentTokenAdded(opts *bind.FilterOpts, token []common.Address) (*TtmaccountPaymentTokenAddedIterator, error) {

	var tokenRule []interface{}
	for _, tokenItem := range token {
		tokenRule = append(tokenRule, tokenItem)
	}

	logs, sub, err := _Ttmaccount.contract.FilterLogs(opts, "PaymentTokenAdded", tokenRule)
	if err != nil {
		return nil, err
	}
	return &TtmaccountPaymentTokenAddedIterator{contract: _Ttmaccount.contract, event: "PaymentTokenAdded", logs: logs, sub: sub}, nil
}

// WatchPaymentTokenAdded is a free log subscription operation binding the contract event 0xa317c10673baf4f03b3c1041bd5ddbb537d0333a86fec3607c75f9dbb630f48f.
//
// Solidity: event PaymentTokenAdded(address indexed token)
func (_Ttmaccount *TtmaccountFilterer) WatchPaymentTokenAdded(opts *bind.WatchOpts, sink chan<- *TtmaccountPaymentTokenAdded, token []common.Address) (event.Subscription, error) {

	var tokenRule []interface{}
	for _, tokenItem := range token {
		tokenRule = append(tokenRule, tokenItem)
	}

	logs, sub, err := _Ttmaccount.contract.WatchLogs(opts, "PaymentTokenAdded", tokenRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(TtmaccountPaymentTokenAdded)
				if err := _Ttmaccount.contract.UnpackLog(event, "PaymentTokenAdded", log); err != nil {
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

// ParsePaymentTokenAdded is a log parse operation binding the contract event 0xa317c10673baf4f03b3c1041bd5ddbb537d0333a86fec3607c75f9dbb630f48f.
//
// Solidity: event PaymentTokenAdded(address indexed token)
func (_Ttmaccount *TtmaccountFilterer) ParsePaymentTokenAdded(log types.Log) (*TtmaccountPaymentTokenAdded, error) {
	event := new(TtmaccountPaymentTokenAdded)
	if err := _Ttmaccount.contract.UnpackLog(event, "PaymentTokenAdded", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// TtmaccountPaymentTokenRemovedIterator is returned from FilterPaymentTokenRemoved and is used to iterate over the raw logs and unpacked data for PaymentTokenRemoved events raised by the Ttmaccount contract.
type TtmaccountPaymentTokenRemovedIterator struct {
	Event *TtmaccountPaymentTokenRemoved // Event containing the contract specifics and raw log

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
func (it *TtmaccountPaymentTokenRemovedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(TtmaccountPaymentTokenRemoved)
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
		it.Event = new(TtmaccountPaymentTokenRemoved)
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
func (it *TtmaccountPaymentTokenRemovedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *TtmaccountPaymentTokenRemovedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// TtmaccountPaymentTokenRemoved represents a PaymentTokenRemoved event raised by the Ttmaccount contract.
type TtmaccountPaymentTokenRemoved struct {
	Token common.Address
	Raw   types.Log // Blockchain specific contextual infos
}

// FilterPaymentTokenRemoved is a free log retrieval operation binding the contract event 0x85a3e72f8dd6db3794f93109c3c5f5b79d6112f6979431c45f98b26134b42af2.
//
// Solidity: event PaymentTokenRemoved(address indexed token)
func (_Ttmaccount *TtmaccountFilterer) FilterPaymentTokenRemoved(opts *bind.FilterOpts, token []common.Address) (*TtmaccountPaymentTokenRemovedIterator, error) {

	var tokenRule []interface{}
	for _, tokenItem := range token {
		tokenRule = append(tokenRule, tokenItem)
	}

	logs, sub, err := _Ttmaccount.contract.FilterLogs(opts, "PaymentTokenRemoved", tokenRule)
	if err != nil {
		return nil, err
	}
	return &TtmaccountPaymentTokenRemovedIterator{contract: _Ttmaccount.contract, event: "PaymentTokenRemoved", logs: logs, sub: sub}, nil
}

// WatchPaymentTokenRemoved is a free log subscription operation binding the contract event 0x85a3e72f8dd6db3794f93109c3c5f5b79d6112f6979431c45f98b26134b42af2.
//
// Solidity: event PaymentTokenRemoved(address indexed token)
func (_Ttmaccount *TtmaccountFilterer) WatchPaymentTokenRemoved(opts *bind.WatchOpts, sink chan<- *TtmaccountPaymentTokenRemoved, token []common.Address) (event.Subscription, error) {

	var tokenRule []interface{}
	for _, tokenItem := range token {
		tokenRule = append(tokenRule, tokenItem)
	}

	logs, sub, err := _Ttmaccount.contract.WatchLogs(opts, "PaymentTokenRemoved", tokenRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(TtmaccountPaymentTokenRemoved)
				if err := _Ttmaccount.contract.UnpackLog(event, "PaymentTokenRemoved", log); err != nil {
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

// ParsePaymentTokenRemoved is a log parse operation binding the contract event 0x85a3e72f8dd6db3794f93109c3c5f5b79d6112f6979431c45f98b26134b42af2.
//
// Solidity: event PaymentTokenRemoved(address indexed token)
func (_Ttmaccount *TtmaccountFilterer) ParsePaymentTokenRemoved(log types.Log) (*TtmaccountPaymentTokenRemoved, error) {
	event := new(TtmaccountPaymentTokenRemoved)
	if err := _Ttmaccount.contract.UnpackLog(event, "PaymentTokenRemoved", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// TtmaccountPublicKeyAddedIterator is returned from FilterPublicKeyAdded and is used to iterate over the raw logs and unpacked data for PublicKeyAdded events raised by the Ttmaccount contract.
type TtmaccountPublicKeyAddedIterator struct {
	Event *TtmaccountPublicKeyAdded // Event containing the contract specifics and raw log

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
func (it *TtmaccountPublicKeyAddedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(TtmaccountPublicKeyAdded)
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
		it.Event = new(TtmaccountPublicKeyAdded)
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
func (it *TtmaccountPublicKeyAddedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *TtmaccountPublicKeyAddedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// TtmaccountPublicKeyAdded represents a PublicKeyAdded event raised by the Ttmaccount contract.
type TtmaccountPublicKeyAdded struct {
	PubKeyAddress common.Address
	Raw           types.Log // Blockchain specific contextual infos
}

// FilterPublicKeyAdded is a free log retrieval operation binding the contract event 0x928ec246afda323bc23c2815ca3f516e9fc6a7b7179772235c221e132545fe82.
//
// Solidity: event PublicKeyAdded(address indexed pubKeyAddress)
func (_Ttmaccount *TtmaccountFilterer) FilterPublicKeyAdded(opts *bind.FilterOpts, pubKeyAddress []common.Address) (*TtmaccountPublicKeyAddedIterator, error) {

	var pubKeyAddressRule []interface{}
	for _, pubKeyAddressItem := range pubKeyAddress {
		pubKeyAddressRule = append(pubKeyAddressRule, pubKeyAddressItem)
	}

	logs, sub, err := _Ttmaccount.contract.FilterLogs(opts, "PublicKeyAdded", pubKeyAddressRule)
	if err != nil {
		return nil, err
	}
	return &TtmaccountPublicKeyAddedIterator{contract: _Ttmaccount.contract, event: "PublicKeyAdded", logs: logs, sub: sub}, nil
}

// WatchPublicKeyAdded is a free log subscription operation binding the contract event 0x928ec246afda323bc23c2815ca3f516e9fc6a7b7179772235c221e132545fe82.
//
// Solidity: event PublicKeyAdded(address indexed pubKeyAddress)
func (_Ttmaccount *TtmaccountFilterer) WatchPublicKeyAdded(opts *bind.WatchOpts, sink chan<- *TtmaccountPublicKeyAdded, pubKeyAddress []common.Address) (event.Subscription, error) {

	var pubKeyAddressRule []interface{}
	for _, pubKeyAddressItem := range pubKeyAddress {
		pubKeyAddressRule = append(pubKeyAddressRule, pubKeyAddressItem)
	}

	logs, sub, err := _Ttmaccount.contract.WatchLogs(opts, "PublicKeyAdded", pubKeyAddressRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(TtmaccountPublicKeyAdded)
				if err := _Ttmaccount.contract.UnpackLog(event, "PublicKeyAdded", log); err != nil {
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

// ParsePublicKeyAdded is a log parse operation binding the contract event 0x928ec246afda323bc23c2815ca3f516e9fc6a7b7179772235c221e132545fe82.
//
// Solidity: event PublicKeyAdded(address indexed pubKeyAddress)
func (_Ttmaccount *TtmaccountFilterer) ParsePublicKeyAdded(log types.Log) (*TtmaccountPublicKeyAdded, error) {
	event := new(TtmaccountPublicKeyAdded)
	if err := _Ttmaccount.contract.UnpackLog(event, "PublicKeyAdded", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// TtmaccountPublicKeyRemovedIterator is returned from FilterPublicKeyRemoved and is used to iterate over the raw logs and unpacked data for PublicKeyRemoved events raised by the Ttmaccount contract.
type TtmaccountPublicKeyRemovedIterator struct {
	Event *TtmaccountPublicKeyRemoved // Event containing the contract specifics and raw log

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
func (it *TtmaccountPublicKeyRemovedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(TtmaccountPublicKeyRemoved)
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
		it.Event = new(TtmaccountPublicKeyRemoved)
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
func (it *TtmaccountPublicKeyRemovedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *TtmaccountPublicKeyRemovedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// TtmaccountPublicKeyRemoved represents a PublicKeyRemoved event raised by the Ttmaccount contract.
type TtmaccountPublicKeyRemoved struct {
	PubKeyAddress common.Address
	Raw           types.Log // Blockchain specific contextual infos
}

// FilterPublicKeyRemoved is a free log retrieval operation binding the contract event 0xc5a9b8041ef11732e7dd7043167d8c22db5c7ea99dcd38dce401effacf8a29bf.
//
// Solidity: event PublicKeyRemoved(address indexed pubKeyAddress)
func (_Ttmaccount *TtmaccountFilterer) FilterPublicKeyRemoved(opts *bind.FilterOpts, pubKeyAddress []common.Address) (*TtmaccountPublicKeyRemovedIterator, error) {

	var pubKeyAddressRule []interface{}
	for _, pubKeyAddressItem := range pubKeyAddress {
		pubKeyAddressRule = append(pubKeyAddressRule, pubKeyAddressItem)
	}

	logs, sub, err := _Ttmaccount.contract.FilterLogs(opts, "PublicKeyRemoved", pubKeyAddressRule)
	if err != nil {
		return nil, err
	}
	return &TtmaccountPublicKeyRemovedIterator{contract: _Ttmaccount.contract, event: "PublicKeyRemoved", logs: logs, sub: sub}, nil
}

// WatchPublicKeyRemoved is a free log subscription operation binding the contract event 0xc5a9b8041ef11732e7dd7043167d8c22db5c7ea99dcd38dce401effacf8a29bf.
//
// Solidity: event PublicKeyRemoved(address indexed pubKeyAddress)
func (_Ttmaccount *TtmaccountFilterer) WatchPublicKeyRemoved(opts *bind.WatchOpts, sink chan<- *TtmaccountPublicKeyRemoved, pubKeyAddress []common.Address) (event.Subscription, error) {

	var pubKeyAddressRule []interface{}
	for _, pubKeyAddressItem := range pubKeyAddress {
		pubKeyAddressRule = append(pubKeyAddressRule, pubKeyAddressItem)
	}

	logs, sub, err := _Ttmaccount.contract.WatchLogs(opts, "PublicKeyRemoved", pubKeyAddressRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(TtmaccountPublicKeyRemoved)
				if err := _Ttmaccount.contract.UnpackLog(event, "PublicKeyRemoved", log); err != nil {
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

// ParsePublicKeyRemoved is a log parse operation binding the contract event 0xc5a9b8041ef11732e7dd7043167d8c22db5c7ea99dcd38dce401effacf8a29bf.
//
// Solidity: event PublicKeyRemoved(address indexed pubKeyAddress)
func (_Ttmaccount *TtmaccountFilterer) ParsePublicKeyRemoved(log types.Log) (*TtmaccountPublicKeyRemoved, error) {
	event := new(TtmaccountPublicKeyRemoved)
	if err := _Ttmaccount.contract.UnpackLog(event, "PublicKeyRemoved", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// TtmaccountRoleAdminChangedIterator is returned from FilterRoleAdminChanged and is used to iterate over the raw logs and unpacked data for RoleAdminChanged events raised by the Ttmaccount contract.
type TtmaccountRoleAdminChangedIterator struct {
	Event *TtmaccountRoleAdminChanged // Event containing the contract specifics and raw log

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
func (it *TtmaccountRoleAdminChangedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(TtmaccountRoleAdminChanged)
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
		it.Event = new(TtmaccountRoleAdminChanged)
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
func (it *TtmaccountRoleAdminChangedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *TtmaccountRoleAdminChangedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// TtmaccountRoleAdminChanged represents a RoleAdminChanged event raised by the Ttmaccount contract.
type TtmaccountRoleAdminChanged struct {
	Role              [32]byte
	PreviousAdminRole [32]byte
	NewAdminRole      [32]byte
	Raw               types.Log // Blockchain specific contextual infos
}

// FilterRoleAdminChanged is a free log retrieval operation binding the contract event 0xbd79b86ffe0ab8e8776151514217cd7cacd52c909f66475c3af44e129f0b00ff.
//
// Solidity: event RoleAdminChanged(bytes32 indexed role, bytes32 indexed previousAdminRole, bytes32 indexed newAdminRole)
func (_Ttmaccount *TtmaccountFilterer) FilterRoleAdminChanged(opts *bind.FilterOpts, role [][32]byte, previousAdminRole [][32]byte, newAdminRole [][32]byte) (*TtmaccountRoleAdminChangedIterator, error) {

	var roleRule []interface{}
	for _, roleItem := range role {
		roleRule = append(roleRule, roleItem)
	}
	var previousAdminRoleRule []interface{}
	for _, previousAdminRoleItem := range previousAdminRole {
		previousAdminRoleRule = append(previousAdminRoleRule, previousAdminRoleItem)
	}
	var newAdminRoleRule []interface{}
	for _, newAdminRoleItem := range newAdminRole {
		newAdminRoleRule = append(newAdminRoleRule, newAdminRoleItem)
	}

	logs, sub, err := _Ttmaccount.contract.FilterLogs(opts, "RoleAdminChanged", roleRule, previousAdminRoleRule, newAdminRoleRule)
	if err != nil {
		return nil, err
	}
	return &TtmaccountRoleAdminChangedIterator{contract: _Ttmaccount.contract, event: "RoleAdminChanged", logs: logs, sub: sub}, nil
}

// WatchRoleAdminChanged is a free log subscription operation binding the contract event 0xbd79b86ffe0ab8e8776151514217cd7cacd52c909f66475c3af44e129f0b00ff.
//
// Solidity: event RoleAdminChanged(bytes32 indexed role, bytes32 indexed previousAdminRole, bytes32 indexed newAdminRole)
func (_Ttmaccount *TtmaccountFilterer) WatchRoleAdminChanged(opts *bind.WatchOpts, sink chan<- *TtmaccountRoleAdminChanged, role [][32]byte, previousAdminRole [][32]byte, newAdminRole [][32]byte) (event.Subscription, error) {

	var roleRule []interface{}
	for _, roleItem := range role {
		roleRule = append(roleRule, roleItem)
	}
	var previousAdminRoleRule []interface{}
	for _, previousAdminRoleItem := range previousAdminRole {
		previousAdminRoleRule = append(previousAdminRoleRule, previousAdminRoleItem)
	}
	var newAdminRoleRule []interface{}
	for _, newAdminRoleItem := range newAdminRole {
		newAdminRoleRule = append(newAdminRoleRule, newAdminRoleItem)
	}

	logs, sub, err := _Ttmaccount.contract.WatchLogs(opts, "RoleAdminChanged", roleRule, previousAdminRoleRule, newAdminRoleRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(TtmaccountRoleAdminChanged)
				if err := _Ttmaccount.contract.UnpackLog(event, "RoleAdminChanged", log); err != nil {
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

// ParseRoleAdminChanged is a log parse operation binding the contract event 0xbd79b86ffe0ab8e8776151514217cd7cacd52c909f66475c3af44e129f0b00ff.
//
// Solidity: event RoleAdminChanged(bytes32 indexed role, bytes32 indexed previousAdminRole, bytes32 indexed newAdminRole)
func (_Ttmaccount *TtmaccountFilterer) ParseRoleAdminChanged(log types.Log) (*TtmaccountRoleAdminChanged, error) {
	event := new(TtmaccountRoleAdminChanged)
	if err := _Ttmaccount.contract.UnpackLog(event, "RoleAdminChanged", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// TtmaccountRoleGrantedIterator is returned from FilterRoleGranted and is used to iterate over the raw logs and unpacked data for RoleGranted events raised by the Ttmaccount contract.
type TtmaccountRoleGrantedIterator struct {
	Event *TtmaccountRoleGranted // Event containing the contract specifics and raw log

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
func (it *TtmaccountRoleGrantedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(TtmaccountRoleGranted)
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
		it.Event = new(TtmaccountRoleGranted)
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
func (it *TtmaccountRoleGrantedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *TtmaccountRoleGrantedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// TtmaccountRoleGranted represents a RoleGranted event raised by the Ttmaccount contract.
type TtmaccountRoleGranted struct {
	Role    [32]byte
	Account common.Address
	Sender  common.Address
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterRoleGranted is a free log retrieval operation binding the contract event 0x2f8788117e7eff1d82e926ec794901d17c78024a50270940304540a733656f0d.
//
// Solidity: event RoleGranted(bytes32 indexed role, address indexed account, address indexed sender)
func (_Ttmaccount *TtmaccountFilterer) FilterRoleGranted(opts *bind.FilterOpts, role [][32]byte, account []common.Address, sender []common.Address) (*TtmaccountRoleGrantedIterator, error) {

	var roleRule []interface{}
	for _, roleItem := range role {
		roleRule = append(roleRule, roleItem)
	}
	var accountRule []interface{}
	for _, accountItem := range account {
		accountRule = append(accountRule, accountItem)
	}
	var senderRule []interface{}
	for _, senderItem := range sender {
		senderRule = append(senderRule, senderItem)
	}

	logs, sub, err := _Ttmaccount.contract.FilterLogs(opts, "RoleGranted", roleRule, accountRule, senderRule)
	if err != nil {
		return nil, err
	}
	return &TtmaccountRoleGrantedIterator{contract: _Ttmaccount.contract, event: "RoleGranted", logs: logs, sub: sub}, nil
}

// WatchRoleGranted is a free log subscription operation binding the contract event 0x2f8788117e7eff1d82e926ec794901d17c78024a50270940304540a733656f0d.
//
// Solidity: event RoleGranted(bytes32 indexed role, address indexed account, address indexed sender)
func (_Ttmaccount *TtmaccountFilterer) WatchRoleGranted(opts *bind.WatchOpts, sink chan<- *TtmaccountRoleGranted, role [][32]byte, account []common.Address, sender []common.Address) (event.Subscription, error) {

	var roleRule []interface{}
	for _, roleItem := range role {
		roleRule = append(roleRule, roleItem)
	}
	var accountRule []interface{}
	for _, accountItem := range account {
		accountRule = append(accountRule, accountItem)
	}
	var senderRule []interface{}
	for _, senderItem := range sender {
		senderRule = append(senderRule, senderItem)
	}

	logs, sub, err := _Ttmaccount.contract.WatchLogs(opts, "RoleGranted", roleRule, accountRule, senderRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(TtmaccountRoleGranted)
				if err := _Ttmaccount.contract.UnpackLog(event, "RoleGranted", log); err != nil {
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

// ParseRoleGranted is a log parse operation binding the contract event 0x2f8788117e7eff1d82e926ec794901d17c78024a50270940304540a733656f0d.
//
// Solidity: event RoleGranted(bytes32 indexed role, address indexed account, address indexed sender)
func (_Ttmaccount *TtmaccountFilterer) ParseRoleGranted(log types.Log) (*TtmaccountRoleGranted, error) {
	event := new(TtmaccountRoleGranted)
	if err := _Ttmaccount.contract.UnpackLog(event, "RoleGranted", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// TtmaccountRoleRevokedIterator is returned from FilterRoleRevoked and is used to iterate over the raw logs and unpacked data for RoleRevoked events raised by the Ttmaccount contract.
type TtmaccountRoleRevokedIterator struct {
	Event *TtmaccountRoleRevoked // Event containing the contract specifics and raw log

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
func (it *TtmaccountRoleRevokedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(TtmaccountRoleRevoked)
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
		it.Event = new(TtmaccountRoleRevoked)
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
func (it *TtmaccountRoleRevokedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *TtmaccountRoleRevokedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// TtmaccountRoleRevoked represents a RoleRevoked event raised by the Ttmaccount contract.
type TtmaccountRoleRevoked struct {
	Role    [32]byte
	Account common.Address
	Sender  common.Address
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterRoleRevoked is a free log retrieval operation binding the contract event 0xf6391f5c32d9c69d2a47ea670b442974b53935d1edc7fd64eb21e047a839171b.
//
// Solidity: event RoleRevoked(bytes32 indexed role, address indexed account, address indexed sender)
func (_Ttmaccount *TtmaccountFilterer) FilterRoleRevoked(opts *bind.FilterOpts, role [][32]byte, account []common.Address, sender []common.Address) (*TtmaccountRoleRevokedIterator, error) {

	var roleRule []interface{}
	for _, roleItem := range role {
		roleRule = append(roleRule, roleItem)
	}
	var accountRule []interface{}
	for _, accountItem := range account {
		accountRule = append(accountRule, accountItem)
	}
	var senderRule []interface{}
	for _, senderItem := range sender {
		senderRule = append(senderRule, senderItem)
	}

	logs, sub, err := _Ttmaccount.contract.FilterLogs(opts, "RoleRevoked", roleRule, accountRule, senderRule)
	if err != nil {
		return nil, err
	}
	return &TtmaccountRoleRevokedIterator{contract: _Ttmaccount.contract, event: "RoleRevoked", logs: logs, sub: sub}, nil
}

// WatchRoleRevoked is a free log subscription operation binding the contract event 0xf6391f5c32d9c69d2a47ea670b442974b53935d1edc7fd64eb21e047a839171b.
//
// Solidity: event RoleRevoked(bytes32 indexed role, address indexed account, address indexed sender)
func (_Ttmaccount *TtmaccountFilterer) WatchRoleRevoked(opts *bind.WatchOpts, sink chan<- *TtmaccountRoleRevoked, role [][32]byte, account []common.Address, sender []common.Address) (event.Subscription, error) {

	var roleRule []interface{}
	for _, roleItem := range role {
		roleRule = append(roleRule, roleItem)
	}
	var accountRule []interface{}
	for _, accountItem := range account {
		accountRule = append(accountRule, accountItem)
	}
	var senderRule []interface{}
	for _, senderItem := range sender {
		senderRule = append(senderRule, senderItem)
	}

	logs, sub, err := _Ttmaccount.contract.WatchLogs(opts, "RoleRevoked", roleRule, accountRule, senderRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(TtmaccountRoleRevoked)
				if err := _Ttmaccount.contract.UnpackLog(event, "RoleRevoked", log); err != nil {
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

// ParseRoleRevoked is a log parse operation binding the contract event 0xf6391f5c32d9c69d2a47ea670b442974b53935d1edc7fd64eb21e047a839171b.
//
// Solidity: event RoleRevoked(bytes32 indexed role, address indexed account, address indexed sender)
func (_Ttmaccount *TtmaccountFilterer) ParseRoleRevoked(log types.Log) (*TtmaccountRoleRevoked, error) {
	event := new(TtmaccountRoleRevoked)
	if err := _Ttmaccount.contract.UnpackLog(event, "RoleRevoked", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// TtmaccountServiceAddedIterator is returned from FilterServiceAdded and is used to iterate over the raw logs and unpacked data for ServiceAdded events raised by the Ttmaccount contract.
type TtmaccountServiceAddedIterator struct {
	Event *TtmaccountServiceAdded // Event containing the contract specifics and raw log

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
func (it *TtmaccountServiceAddedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(TtmaccountServiceAdded)
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
		it.Event = new(TtmaccountServiceAdded)
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
func (it *TtmaccountServiceAddedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *TtmaccountServiceAddedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// TtmaccountServiceAdded represents a ServiceAdded event raised by the Ttmaccount contract.
type TtmaccountServiceAdded struct {
	ServiceHash [32]byte
	Raw         types.Log // Blockchain specific contextual infos
}

// FilterServiceAdded is a free log retrieval operation binding the contract event 0x8f531e5ede07d5741fd086bb787ed399a64704eb757b87cc80cf6635b274e5b5.
//
// Solidity: event ServiceAdded(bytes32 indexed serviceHash)
func (_Ttmaccount *TtmaccountFilterer) FilterServiceAdded(opts *bind.FilterOpts, serviceHash [][32]byte) (*TtmaccountServiceAddedIterator, error) {

	var serviceHashRule []interface{}
	for _, serviceHashItem := range serviceHash {
		serviceHashRule = append(serviceHashRule, serviceHashItem)
	}

	logs, sub, err := _Ttmaccount.contract.FilterLogs(opts, "ServiceAdded", serviceHashRule)
	if err != nil {
		return nil, err
	}
	return &TtmaccountServiceAddedIterator{contract: _Ttmaccount.contract, event: "ServiceAdded", logs: logs, sub: sub}, nil
}

// WatchServiceAdded is a free log subscription operation binding the contract event 0x8f531e5ede07d5741fd086bb787ed399a64704eb757b87cc80cf6635b274e5b5.
//
// Solidity: event ServiceAdded(bytes32 indexed serviceHash)
func (_Ttmaccount *TtmaccountFilterer) WatchServiceAdded(opts *bind.WatchOpts, sink chan<- *TtmaccountServiceAdded, serviceHash [][32]byte) (event.Subscription, error) {

	var serviceHashRule []interface{}
	for _, serviceHashItem := range serviceHash {
		serviceHashRule = append(serviceHashRule, serviceHashItem)
	}

	logs, sub, err := _Ttmaccount.contract.WatchLogs(opts, "ServiceAdded", serviceHashRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(TtmaccountServiceAdded)
				if err := _Ttmaccount.contract.UnpackLog(event, "ServiceAdded", log); err != nil {
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

// ParseServiceAdded is a log parse operation binding the contract event 0x8f531e5ede07d5741fd086bb787ed399a64704eb757b87cc80cf6635b274e5b5.
//
// Solidity: event ServiceAdded(bytes32 indexed serviceHash)
func (_Ttmaccount *TtmaccountFilterer) ParseServiceAdded(log types.Log) (*TtmaccountServiceAdded, error) {
	event := new(TtmaccountServiceAdded)
	if err := _Ttmaccount.contract.UnpackLog(event, "ServiceAdded", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// TtmaccountServiceCapabilitiesUpdatedIterator is returned from FilterServiceCapabilitiesUpdated and is used to iterate over the raw logs and unpacked data for ServiceCapabilitiesUpdated events raised by the Ttmaccount contract.
type TtmaccountServiceCapabilitiesUpdatedIterator struct {
	Event *TtmaccountServiceCapabilitiesUpdated // Event containing the contract specifics and raw log

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
func (it *TtmaccountServiceCapabilitiesUpdatedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(TtmaccountServiceCapabilitiesUpdated)
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
		it.Event = new(TtmaccountServiceCapabilitiesUpdated)
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
func (it *TtmaccountServiceCapabilitiesUpdatedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *TtmaccountServiceCapabilitiesUpdatedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// TtmaccountServiceCapabilitiesUpdated represents a ServiceCapabilitiesUpdated event raised by the Ttmaccount contract.
type TtmaccountServiceCapabilitiesUpdated struct {
	ServiceHash [32]byte
	Raw         types.Log // Blockchain specific contextual infos
}

// FilterServiceCapabilitiesUpdated is a free log retrieval operation binding the contract event 0xa616bfc5bb0e46c6cad727e1b55e3685067e1296d962a7f37017874a27aa0098.
//
// Solidity: event ServiceCapabilitiesUpdated(bytes32 indexed serviceHash)
func (_Ttmaccount *TtmaccountFilterer) FilterServiceCapabilitiesUpdated(opts *bind.FilterOpts, serviceHash [][32]byte) (*TtmaccountServiceCapabilitiesUpdatedIterator, error) {

	var serviceHashRule []interface{}
	for _, serviceHashItem := range serviceHash {
		serviceHashRule = append(serviceHashRule, serviceHashItem)
	}

	logs, sub, err := _Ttmaccount.contract.FilterLogs(opts, "ServiceCapabilitiesUpdated", serviceHashRule)
	if err != nil {
		return nil, err
	}
	return &TtmaccountServiceCapabilitiesUpdatedIterator{contract: _Ttmaccount.contract, event: "ServiceCapabilitiesUpdated", logs: logs, sub: sub}, nil
}

// WatchServiceCapabilitiesUpdated is a free log subscription operation binding the contract event 0xa616bfc5bb0e46c6cad727e1b55e3685067e1296d962a7f37017874a27aa0098.
//
// Solidity: event ServiceCapabilitiesUpdated(bytes32 indexed serviceHash)
func (_Ttmaccount *TtmaccountFilterer) WatchServiceCapabilitiesUpdated(opts *bind.WatchOpts, sink chan<- *TtmaccountServiceCapabilitiesUpdated, serviceHash [][32]byte) (event.Subscription, error) {

	var serviceHashRule []interface{}
	for _, serviceHashItem := range serviceHash {
		serviceHashRule = append(serviceHashRule, serviceHashItem)
	}

	logs, sub, err := _Ttmaccount.contract.WatchLogs(opts, "ServiceCapabilitiesUpdated", serviceHashRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(TtmaccountServiceCapabilitiesUpdated)
				if err := _Ttmaccount.contract.UnpackLog(event, "ServiceCapabilitiesUpdated", log); err != nil {
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

// ParseServiceCapabilitiesUpdated is a log parse operation binding the contract event 0xa616bfc5bb0e46c6cad727e1b55e3685067e1296d962a7f37017874a27aa0098.
//
// Solidity: event ServiceCapabilitiesUpdated(bytes32 indexed serviceHash)
func (_Ttmaccount *TtmaccountFilterer) ParseServiceCapabilitiesUpdated(log types.Log) (*TtmaccountServiceCapabilitiesUpdated, error) {
	event := new(TtmaccountServiceCapabilitiesUpdated)
	if err := _Ttmaccount.contract.UnpackLog(event, "ServiceCapabilitiesUpdated", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// TtmaccountServiceCapabilityAddedIterator is returned from FilterServiceCapabilityAdded and is used to iterate over the raw logs and unpacked data for ServiceCapabilityAdded events raised by the Ttmaccount contract.
type TtmaccountServiceCapabilityAddedIterator struct {
	Event *TtmaccountServiceCapabilityAdded // Event containing the contract specifics and raw log

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
func (it *TtmaccountServiceCapabilityAddedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(TtmaccountServiceCapabilityAdded)
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
		it.Event = new(TtmaccountServiceCapabilityAdded)
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
func (it *TtmaccountServiceCapabilityAddedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *TtmaccountServiceCapabilityAddedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// TtmaccountServiceCapabilityAdded represents a ServiceCapabilityAdded event raised by the Ttmaccount contract.
type TtmaccountServiceCapabilityAdded struct {
	ServiceHash [32]byte
	Capability  string
	Raw         types.Log // Blockchain specific contextual infos
}

// FilterServiceCapabilityAdded is a free log retrieval operation binding the contract event 0x1cd139430ed537ab9e8086952076cce01edd5ba6e30907af0ffe3709fd3139e6.
//
// Solidity: event ServiceCapabilityAdded(bytes32 indexed serviceHash, string capability)
func (_Ttmaccount *TtmaccountFilterer) FilterServiceCapabilityAdded(opts *bind.FilterOpts, serviceHash [][32]byte) (*TtmaccountServiceCapabilityAddedIterator, error) {

	var serviceHashRule []interface{}
	for _, serviceHashItem := range serviceHash {
		serviceHashRule = append(serviceHashRule, serviceHashItem)
	}

	logs, sub, err := _Ttmaccount.contract.FilterLogs(opts, "ServiceCapabilityAdded", serviceHashRule)
	if err != nil {
		return nil, err
	}
	return &TtmaccountServiceCapabilityAddedIterator{contract: _Ttmaccount.contract, event: "ServiceCapabilityAdded", logs: logs, sub: sub}, nil
}

// WatchServiceCapabilityAdded is a free log subscription operation binding the contract event 0x1cd139430ed537ab9e8086952076cce01edd5ba6e30907af0ffe3709fd3139e6.
//
// Solidity: event ServiceCapabilityAdded(bytes32 indexed serviceHash, string capability)
func (_Ttmaccount *TtmaccountFilterer) WatchServiceCapabilityAdded(opts *bind.WatchOpts, sink chan<- *TtmaccountServiceCapabilityAdded, serviceHash [][32]byte) (event.Subscription, error) {

	var serviceHashRule []interface{}
	for _, serviceHashItem := range serviceHash {
		serviceHashRule = append(serviceHashRule, serviceHashItem)
	}

	logs, sub, err := _Ttmaccount.contract.WatchLogs(opts, "ServiceCapabilityAdded", serviceHashRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(TtmaccountServiceCapabilityAdded)
				if err := _Ttmaccount.contract.UnpackLog(event, "ServiceCapabilityAdded", log); err != nil {
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

// ParseServiceCapabilityAdded is a log parse operation binding the contract event 0x1cd139430ed537ab9e8086952076cce01edd5ba6e30907af0ffe3709fd3139e6.
//
// Solidity: event ServiceCapabilityAdded(bytes32 indexed serviceHash, string capability)
func (_Ttmaccount *TtmaccountFilterer) ParseServiceCapabilityAdded(log types.Log) (*TtmaccountServiceCapabilityAdded, error) {
	event := new(TtmaccountServiceCapabilityAdded)
	if err := _Ttmaccount.contract.UnpackLog(event, "ServiceCapabilityAdded", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// TtmaccountServiceCapabilityRemovedIterator is returned from FilterServiceCapabilityRemoved and is used to iterate over the raw logs and unpacked data for ServiceCapabilityRemoved events raised by the Ttmaccount contract.
type TtmaccountServiceCapabilityRemovedIterator struct {
	Event *TtmaccountServiceCapabilityRemoved // Event containing the contract specifics and raw log

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
func (it *TtmaccountServiceCapabilityRemovedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(TtmaccountServiceCapabilityRemoved)
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
		it.Event = new(TtmaccountServiceCapabilityRemoved)
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
func (it *TtmaccountServiceCapabilityRemovedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *TtmaccountServiceCapabilityRemovedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// TtmaccountServiceCapabilityRemoved represents a ServiceCapabilityRemoved event raised by the Ttmaccount contract.
type TtmaccountServiceCapabilityRemoved struct {
	ServiceHash [32]byte
	Capability  string
	Raw         types.Log // Blockchain specific contextual infos
}

// FilterServiceCapabilityRemoved is a free log retrieval operation binding the contract event 0xfc8d82c9e7e7938446da05458183efa5916c443a2bab87f97f94a8d47742b014.
//
// Solidity: event ServiceCapabilityRemoved(bytes32 indexed serviceHash, string capability)
func (_Ttmaccount *TtmaccountFilterer) FilterServiceCapabilityRemoved(opts *bind.FilterOpts, serviceHash [][32]byte) (*TtmaccountServiceCapabilityRemovedIterator, error) {

	var serviceHashRule []interface{}
	for _, serviceHashItem := range serviceHash {
		serviceHashRule = append(serviceHashRule, serviceHashItem)
	}

	logs, sub, err := _Ttmaccount.contract.FilterLogs(opts, "ServiceCapabilityRemoved", serviceHashRule)
	if err != nil {
		return nil, err
	}
	return &TtmaccountServiceCapabilityRemovedIterator{contract: _Ttmaccount.contract, event: "ServiceCapabilityRemoved", logs: logs, sub: sub}, nil
}

// WatchServiceCapabilityRemoved is a free log subscription operation binding the contract event 0xfc8d82c9e7e7938446da05458183efa5916c443a2bab87f97f94a8d47742b014.
//
// Solidity: event ServiceCapabilityRemoved(bytes32 indexed serviceHash, string capability)
func (_Ttmaccount *TtmaccountFilterer) WatchServiceCapabilityRemoved(opts *bind.WatchOpts, sink chan<- *TtmaccountServiceCapabilityRemoved, serviceHash [][32]byte) (event.Subscription, error) {

	var serviceHashRule []interface{}
	for _, serviceHashItem := range serviceHash {
		serviceHashRule = append(serviceHashRule, serviceHashItem)
	}

	logs, sub, err := _Ttmaccount.contract.WatchLogs(opts, "ServiceCapabilityRemoved", serviceHashRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(TtmaccountServiceCapabilityRemoved)
				if err := _Ttmaccount.contract.UnpackLog(event, "ServiceCapabilityRemoved", log); err != nil {
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

// ParseServiceCapabilityRemoved is a log parse operation binding the contract event 0xfc8d82c9e7e7938446da05458183efa5916c443a2bab87f97f94a8d47742b014.
//
// Solidity: event ServiceCapabilityRemoved(bytes32 indexed serviceHash, string capability)
func (_Ttmaccount *TtmaccountFilterer) ParseServiceCapabilityRemoved(log types.Log) (*TtmaccountServiceCapabilityRemoved, error) {
	event := new(TtmaccountServiceCapabilityRemoved)
	if err := _Ttmaccount.contract.UnpackLog(event, "ServiceCapabilityRemoved", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// TtmaccountServiceRemovedIterator is returned from FilterServiceRemoved and is used to iterate over the raw logs and unpacked data for ServiceRemoved events raised by the Ttmaccount contract.
type TtmaccountServiceRemovedIterator struct {
	Event *TtmaccountServiceRemoved // Event containing the contract specifics and raw log

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
func (it *TtmaccountServiceRemovedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(TtmaccountServiceRemoved)
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
		it.Event = new(TtmaccountServiceRemoved)
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
func (it *TtmaccountServiceRemovedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *TtmaccountServiceRemovedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// TtmaccountServiceRemoved represents a ServiceRemoved event raised by the Ttmaccount contract.
type TtmaccountServiceRemoved struct {
	ServiceHash [32]byte
	Raw         types.Log // Blockchain specific contextual infos
}

// FilterServiceRemoved is a free log retrieval operation binding the contract event 0x94da5eeca10d4d6ee8455f99240c10b0c74b0cf5bf754afb81c81e2704b9c427.
//
// Solidity: event ServiceRemoved(bytes32 indexed serviceHash)
func (_Ttmaccount *TtmaccountFilterer) FilterServiceRemoved(opts *bind.FilterOpts, serviceHash [][32]byte) (*TtmaccountServiceRemovedIterator, error) {

	var serviceHashRule []interface{}
	for _, serviceHashItem := range serviceHash {
		serviceHashRule = append(serviceHashRule, serviceHashItem)
	}

	logs, sub, err := _Ttmaccount.contract.FilterLogs(opts, "ServiceRemoved", serviceHashRule)
	if err != nil {
		return nil, err
	}
	return &TtmaccountServiceRemovedIterator{contract: _Ttmaccount.contract, event: "ServiceRemoved", logs: logs, sub: sub}, nil
}

// WatchServiceRemoved is a free log subscription operation binding the contract event 0x94da5eeca10d4d6ee8455f99240c10b0c74b0cf5bf754afb81c81e2704b9c427.
//
// Solidity: event ServiceRemoved(bytes32 indexed serviceHash)
func (_Ttmaccount *TtmaccountFilterer) WatchServiceRemoved(opts *bind.WatchOpts, sink chan<- *TtmaccountServiceRemoved, serviceHash [][32]byte) (event.Subscription, error) {

	var serviceHashRule []interface{}
	for _, serviceHashItem := range serviceHash {
		serviceHashRule = append(serviceHashRule, serviceHashItem)
	}

	logs, sub, err := _Ttmaccount.contract.WatchLogs(opts, "ServiceRemoved", serviceHashRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(TtmaccountServiceRemoved)
				if err := _Ttmaccount.contract.UnpackLog(event, "ServiceRemoved", log); err != nil {
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

// ParseServiceRemoved is a log parse operation binding the contract event 0x94da5eeca10d4d6ee8455f99240c10b0c74b0cf5bf754afb81c81e2704b9c427.
//
// Solidity: event ServiceRemoved(bytes32 indexed serviceHash)
func (_Ttmaccount *TtmaccountFilterer) ParseServiceRemoved(log types.Log) (*TtmaccountServiceRemoved, error) {
	event := new(TtmaccountServiceRemoved)
	if err := _Ttmaccount.contract.UnpackLog(event, "ServiceRemoved", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// TtmaccountServiceRestrictedRateUpdatedIterator is returned from FilterServiceRestrictedRateUpdated and is used to iterate over the raw logs and unpacked data for ServiceRestrictedRateUpdated events raised by the Ttmaccount contract.
type TtmaccountServiceRestrictedRateUpdatedIterator struct {
	Event *TtmaccountServiceRestrictedRateUpdated // Event containing the contract specifics and raw log

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
func (it *TtmaccountServiceRestrictedRateUpdatedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(TtmaccountServiceRestrictedRateUpdated)
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
		it.Event = new(TtmaccountServiceRestrictedRateUpdated)
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
func (it *TtmaccountServiceRestrictedRateUpdatedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *TtmaccountServiceRestrictedRateUpdatedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// TtmaccountServiceRestrictedRateUpdated represents a ServiceRestrictedRateUpdated event raised by the Ttmaccount contract.
type TtmaccountServiceRestrictedRateUpdated struct {
	ServiceHash    [32]byte
	RestrictedRate bool
	Raw            types.Log // Blockchain specific contextual infos
}

// FilterServiceRestrictedRateUpdated is a free log retrieval operation binding the contract event 0x1b76230b39d2d0c1a2a77a90c170190d2280796ed56b280177256ce39df1a664.
//
// Solidity: event ServiceRestrictedRateUpdated(bytes32 indexed serviceHash, bool restrictedRate)
func (_Ttmaccount *TtmaccountFilterer) FilterServiceRestrictedRateUpdated(opts *bind.FilterOpts, serviceHash [][32]byte) (*TtmaccountServiceRestrictedRateUpdatedIterator, error) {

	var serviceHashRule []interface{}
	for _, serviceHashItem := range serviceHash {
		serviceHashRule = append(serviceHashRule, serviceHashItem)
	}

	logs, sub, err := _Ttmaccount.contract.FilterLogs(opts, "ServiceRestrictedRateUpdated", serviceHashRule)
	if err != nil {
		return nil, err
	}
	return &TtmaccountServiceRestrictedRateUpdatedIterator{contract: _Ttmaccount.contract, event: "ServiceRestrictedRateUpdated", logs: logs, sub: sub}, nil
}

// WatchServiceRestrictedRateUpdated is a free log subscription operation binding the contract event 0x1b76230b39d2d0c1a2a77a90c170190d2280796ed56b280177256ce39df1a664.
//
// Solidity: event ServiceRestrictedRateUpdated(bytes32 indexed serviceHash, bool restrictedRate)
func (_Ttmaccount *TtmaccountFilterer) WatchServiceRestrictedRateUpdated(opts *bind.WatchOpts, sink chan<- *TtmaccountServiceRestrictedRateUpdated, serviceHash [][32]byte) (event.Subscription, error) {

	var serviceHashRule []interface{}
	for _, serviceHashItem := range serviceHash {
		serviceHashRule = append(serviceHashRule, serviceHashItem)
	}

	logs, sub, err := _Ttmaccount.contract.WatchLogs(opts, "ServiceRestrictedRateUpdated", serviceHashRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(TtmaccountServiceRestrictedRateUpdated)
				if err := _Ttmaccount.contract.UnpackLog(event, "ServiceRestrictedRateUpdated", log); err != nil {
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

// ParseServiceRestrictedRateUpdated is a log parse operation binding the contract event 0x1b76230b39d2d0c1a2a77a90c170190d2280796ed56b280177256ce39df1a664.
//
// Solidity: event ServiceRestrictedRateUpdated(bytes32 indexed serviceHash, bool restrictedRate)
func (_Ttmaccount *TtmaccountFilterer) ParseServiceRestrictedRateUpdated(log types.Log) (*TtmaccountServiceRestrictedRateUpdated, error) {
	event := new(TtmaccountServiceRestrictedRateUpdated)
	if err := _Ttmaccount.contract.UnpackLog(event, "ServiceRestrictedRateUpdated", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// TtmaccountTTMAccountUpgradedIterator is returned from FilterTTMAccountUpgraded and is used to iterate over the raw logs and unpacked data for TTMAccountUpgraded events raised by the Ttmaccount contract.
type TtmaccountTTMAccountUpgradedIterator struct {
	Event *TtmaccountTTMAccountUpgraded // Event containing the contract specifics and raw log

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
func (it *TtmaccountTTMAccountUpgradedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(TtmaccountTTMAccountUpgraded)
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
		it.Event = new(TtmaccountTTMAccountUpgraded)
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
func (it *TtmaccountTTMAccountUpgradedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *TtmaccountTTMAccountUpgradedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// TtmaccountTTMAccountUpgraded represents a TTMAccountUpgraded event raised by the Ttmaccount contract.
type TtmaccountTTMAccountUpgraded struct {
	OldImplementation common.Address
	NewImplementation common.Address
	Raw               types.Log // Blockchain specific contextual infos
}

// FilterTTMAccountUpgraded is a free log retrieval operation binding the contract event 0x897c7778b6095182ea48ee84760832efeae452e4c42d863ea35b271a3aaae759.
//
// Solidity: event TTMAccountUpgraded(address indexed oldImplementation, address indexed newImplementation)
func (_Ttmaccount *TtmaccountFilterer) FilterTTMAccountUpgraded(opts *bind.FilterOpts, oldImplementation []common.Address, newImplementation []common.Address) (*TtmaccountTTMAccountUpgradedIterator, error) {

	var oldImplementationRule []interface{}
	for _, oldImplementationItem := range oldImplementation {
		oldImplementationRule = append(oldImplementationRule, oldImplementationItem)
	}
	var newImplementationRule []interface{}
	for _, newImplementationItem := range newImplementation {
		newImplementationRule = append(newImplementationRule, newImplementationItem)
	}

	logs, sub, err := _Ttmaccount.contract.FilterLogs(opts, "TTMAccountUpgraded", oldImplementationRule, newImplementationRule)
	if err != nil {
		return nil, err
	}
	return &TtmaccountTTMAccountUpgradedIterator{contract: _Ttmaccount.contract, event: "TTMAccountUpgraded", logs: logs, sub: sub}, nil
}

// WatchTTMAccountUpgraded is a free log subscription operation binding the contract event 0x897c7778b6095182ea48ee84760832efeae452e4c42d863ea35b271a3aaae759.
//
// Solidity: event TTMAccountUpgraded(address indexed oldImplementation, address indexed newImplementation)
func (_Ttmaccount *TtmaccountFilterer) WatchTTMAccountUpgraded(opts *bind.WatchOpts, sink chan<- *TtmaccountTTMAccountUpgraded, oldImplementation []common.Address, newImplementation []common.Address) (event.Subscription, error) {

	var oldImplementationRule []interface{}
	for _, oldImplementationItem := range oldImplementation {
		oldImplementationRule = append(oldImplementationRule, oldImplementationItem)
	}
	var newImplementationRule []interface{}
	for _, newImplementationItem := range newImplementation {
		newImplementationRule = append(newImplementationRule, newImplementationItem)
	}

	logs, sub, err := _Ttmaccount.contract.WatchLogs(opts, "TTMAccountUpgraded", oldImplementationRule, newImplementationRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(TtmaccountTTMAccountUpgraded)
				if err := _Ttmaccount.contract.UnpackLog(event, "TTMAccountUpgraded", log); err != nil {
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

// ParseTTMAccountUpgraded is a log parse operation binding the contract event 0x897c7778b6095182ea48ee84760832efeae452e4c42d863ea35b271a3aaae759.
//
// Solidity: event TTMAccountUpgraded(address indexed oldImplementation, address indexed newImplementation)
func (_Ttmaccount *TtmaccountFilterer) ParseTTMAccountUpgraded(log types.Log) (*TtmaccountTTMAccountUpgraded, error) {
	event := new(TtmaccountTTMAccountUpgraded)
	if err := _Ttmaccount.contract.UnpackLog(event, "TTMAccountUpgraded", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// TtmaccountUpgradedIterator is returned from FilterUpgraded and is used to iterate over the raw logs and unpacked data for Upgraded events raised by the Ttmaccount contract.
type TtmaccountUpgradedIterator struct {
	Event *TtmaccountUpgraded // Event containing the contract specifics and raw log

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
func (it *TtmaccountUpgradedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(TtmaccountUpgraded)
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
		it.Event = new(TtmaccountUpgraded)
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
func (it *TtmaccountUpgradedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *TtmaccountUpgradedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// TtmaccountUpgraded represents a Upgraded event raised by the Ttmaccount contract.
type TtmaccountUpgraded struct {
	Implementation common.Address
	Raw            types.Log // Blockchain specific contextual infos
}

// FilterUpgraded is a free log retrieval operation binding the contract event 0xbc7cd75a20ee27fd9adebab32041f755214dbc6bffa90cc0225b39da2e5c2d3b.
//
// Solidity: event Upgraded(address indexed implementation)
func (_Ttmaccount *TtmaccountFilterer) FilterUpgraded(opts *bind.FilterOpts, implementation []common.Address) (*TtmaccountUpgradedIterator, error) {

	var implementationRule []interface{}
	for _, implementationItem := range implementation {
		implementationRule = append(implementationRule, implementationItem)
	}

	logs, sub, err := _Ttmaccount.contract.FilterLogs(opts, "Upgraded", implementationRule)
	if err != nil {
		return nil, err
	}
	return &TtmaccountUpgradedIterator{contract: _Ttmaccount.contract, event: "Upgraded", logs: logs, sub: sub}, nil
}

// WatchUpgraded is a free log subscription operation binding the contract event 0xbc7cd75a20ee27fd9adebab32041f755214dbc6bffa90cc0225b39da2e5c2d3b.
//
// Solidity: event Upgraded(address indexed implementation)
func (_Ttmaccount *TtmaccountFilterer) WatchUpgraded(opts *bind.WatchOpts, sink chan<- *TtmaccountUpgraded, implementation []common.Address) (event.Subscription, error) {

	var implementationRule []interface{}
	for _, implementationItem := range implementation {
		implementationRule = append(implementationRule, implementationItem)
	}

	logs, sub, err := _Ttmaccount.contract.WatchLogs(opts, "Upgraded", implementationRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(TtmaccountUpgraded)
				if err := _Ttmaccount.contract.UnpackLog(event, "Upgraded", log); err != nil {
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

// ParseUpgraded is a log parse operation binding the contract event 0xbc7cd75a20ee27fd9adebab32041f755214dbc6bffa90cc0225b39da2e5c2d3b.
//
// Solidity: event Upgraded(address indexed implementation)
func (_Ttmaccount *TtmaccountFilterer) ParseUpgraded(log types.Log) (*TtmaccountUpgraded, error) {
	event := new(TtmaccountUpgraded)
	if err := _Ttmaccount.contract.UnpackLog(event, "Upgraded", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// TtmaccountWantedServiceAddedIterator is returned from FilterWantedServiceAdded and is used to iterate over the raw logs and unpacked data for WantedServiceAdded events raised by the Ttmaccount contract.
type TtmaccountWantedServiceAddedIterator struct {
	Event *TtmaccountWantedServiceAdded // Event containing the contract specifics and raw log

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
func (it *TtmaccountWantedServiceAddedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(TtmaccountWantedServiceAdded)
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
		it.Event = new(TtmaccountWantedServiceAdded)
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
func (it *TtmaccountWantedServiceAddedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *TtmaccountWantedServiceAddedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// TtmaccountWantedServiceAdded represents a WantedServiceAdded event raised by the Ttmaccount contract.
type TtmaccountWantedServiceAdded struct {
	ServiceHash [32]byte
	Raw         types.Log // Blockchain specific contextual infos
}

// FilterWantedServiceAdded is a free log retrieval operation binding the contract event 0x7acacfd576383587962277516962c289d19f807be443f4e303ab45ace24931ac.
//
// Solidity: event WantedServiceAdded(bytes32 indexed serviceHash)
func (_Ttmaccount *TtmaccountFilterer) FilterWantedServiceAdded(opts *bind.FilterOpts, serviceHash [][32]byte) (*TtmaccountWantedServiceAddedIterator, error) {

	var serviceHashRule []interface{}
	for _, serviceHashItem := range serviceHash {
		serviceHashRule = append(serviceHashRule, serviceHashItem)
	}

	logs, sub, err := _Ttmaccount.contract.FilterLogs(opts, "WantedServiceAdded", serviceHashRule)
	if err != nil {
		return nil, err
	}
	return &TtmaccountWantedServiceAddedIterator{contract: _Ttmaccount.contract, event: "WantedServiceAdded", logs: logs, sub: sub}, nil
}

// WatchWantedServiceAdded is a free log subscription operation binding the contract event 0x7acacfd576383587962277516962c289d19f807be443f4e303ab45ace24931ac.
//
// Solidity: event WantedServiceAdded(bytes32 indexed serviceHash)
func (_Ttmaccount *TtmaccountFilterer) WatchWantedServiceAdded(opts *bind.WatchOpts, sink chan<- *TtmaccountWantedServiceAdded, serviceHash [][32]byte) (event.Subscription, error) {

	var serviceHashRule []interface{}
	for _, serviceHashItem := range serviceHash {
		serviceHashRule = append(serviceHashRule, serviceHashItem)
	}

	logs, sub, err := _Ttmaccount.contract.WatchLogs(opts, "WantedServiceAdded", serviceHashRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(TtmaccountWantedServiceAdded)
				if err := _Ttmaccount.contract.UnpackLog(event, "WantedServiceAdded", log); err != nil {
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

// ParseWantedServiceAdded is a log parse operation binding the contract event 0x7acacfd576383587962277516962c289d19f807be443f4e303ab45ace24931ac.
//
// Solidity: event WantedServiceAdded(bytes32 indexed serviceHash)
func (_Ttmaccount *TtmaccountFilterer) ParseWantedServiceAdded(log types.Log) (*TtmaccountWantedServiceAdded, error) {
	event := new(TtmaccountWantedServiceAdded)
	if err := _Ttmaccount.contract.UnpackLog(event, "WantedServiceAdded", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// TtmaccountWantedServiceRemovedIterator is returned from FilterWantedServiceRemoved and is used to iterate over the raw logs and unpacked data for WantedServiceRemoved events raised by the Ttmaccount contract.
type TtmaccountWantedServiceRemovedIterator struct {
	Event *TtmaccountWantedServiceRemoved // Event containing the contract specifics and raw log

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
func (it *TtmaccountWantedServiceRemovedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(TtmaccountWantedServiceRemoved)
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
		it.Event = new(TtmaccountWantedServiceRemoved)
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
func (it *TtmaccountWantedServiceRemovedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *TtmaccountWantedServiceRemovedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// TtmaccountWantedServiceRemoved represents a WantedServiceRemoved event raised by the Ttmaccount contract.
type TtmaccountWantedServiceRemoved struct {
	ServiceHash [32]byte
	Raw         types.Log // Blockchain specific contextual infos
}

// FilterWantedServiceRemoved is a free log retrieval operation binding the contract event 0xf0dd3de472ddcd75ae2c17728a45801355fb6dd8615a7c53c15504b4279c09be.
//
// Solidity: event WantedServiceRemoved(bytes32 indexed serviceHash)
func (_Ttmaccount *TtmaccountFilterer) FilterWantedServiceRemoved(opts *bind.FilterOpts, serviceHash [][32]byte) (*TtmaccountWantedServiceRemovedIterator, error) {

	var serviceHashRule []interface{}
	for _, serviceHashItem := range serviceHash {
		serviceHashRule = append(serviceHashRule, serviceHashItem)
	}

	logs, sub, err := _Ttmaccount.contract.FilterLogs(opts, "WantedServiceRemoved", serviceHashRule)
	if err != nil {
		return nil, err
	}
	return &TtmaccountWantedServiceRemovedIterator{contract: _Ttmaccount.contract, event: "WantedServiceRemoved", logs: logs, sub: sub}, nil
}

// WatchWantedServiceRemoved is a free log subscription operation binding the contract event 0xf0dd3de472ddcd75ae2c17728a45801355fb6dd8615a7c53c15504b4279c09be.
//
// Solidity: event WantedServiceRemoved(bytes32 indexed serviceHash)
func (_Ttmaccount *TtmaccountFilterer) WatchWantedServiceRemoved(opts *bind.WatchOpts, sink chan<- *TtmaccountWantedServiceRemoved, serviceHash [][32]byte) (event.Subscription, error) {

	var serviceHashRule []interface{}
	for _, serviceHashItem := range serviceHash {
		serviceHashRule = append(serviceHashRule, serviceHashItem)
	}

	logs, sub, err := _Ttmaccount.contract.WatchLogs(opts, "WantedServiceRemoved", serviceHashRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(TtmaccountWantedServiceRemoved)
				if err := _Ttmaccount.contract.UnpackLog(event, "WantedServiceRemoved", log); err != nil {
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

// ParseWantedServiceRemoved is a log parse operation binding the contract event 0xf0dd3de472ddcd75ae2c17728a45801355fb6dd8615a7c53c15504b4279c09be.
//
// Solidity: event WantedServiceRemoved(bytes32 indexed serviceHash)
func (_Ttmaccount *TtmaccountFilterer) ParseWantedServiceRemoved(log types.Log) (*TtmaccountWantedServiceRemoved, error) {
	event := new(TtmaccountWantedServiceRemoved)
	if err := _Ttmaccount.contract.UnpackLog(event, "WantedServiceRemoved", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// TtmaccountWithdrawIterator is returned from FilterWithdraw and is used to iterate over the raw logs and unpacked data for Withdraw events raised by the Ttmaccount contract.
type TtmaccountWithdrawIterator struct {
	Event *TtmaccountWithdraw // Event containing the contract specifics and raw log

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
func (it *TtmaccountWithdrawIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(TtmaccountWithdraw)
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
		it.Event = new(TtmaccountWithdraw)
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
func (it *TtmaccountWithdrawIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *TtmaccountWithdrawIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// TtmaccountWithdraw represents a Withdraw event raised by the Ttmaccount contract.
type TtmaccountWithdraw struct {
	Receiver common.Address
	Amount   *big.Int
	Raw      types.Log // Blockchain specific contextual infos
}

// FilterWithdraw is a free log retrieval operation binding the contract event 0x884edad9ce6fa2440d8a54cc123490eb96d2768479d49ff9c7366125a9424364.
//
// Solidity: event Withdraw(address indexed receiver, uint256 amount)
func (_Ttmaccount *TtmaccountFilterer) FilterWithdraw(opts *bind.FilterOpts, receiver []common.Address) (*TtmaccountWithdrawIterator, error) {

	var receiverRule []interface{}
	for _, receiverItem := range receiver {
		receiverRule = append(receiverRule, receiverItem)
	}

	logs, sub, err := _Ttmaccount.contract.FilterLogs(opts, "Withdraw", receiverRule)
	if err != nil {
		return nil, err
	}
	return &TtmaccountWithdrawIterator{contract: _Ttmaccount.contract, event: "Withdraw", logs: logs, sub: sub}, nil
}

// WatchWithdraw is a free log subscription operation binding the contract event 0x884edad9ce6fa2440d8a54cc123490eb96d2768479d49ff9c7366125a9424364.
//
// Solidity: event Withdraw(address indexed receiver, uint256 amount)
func (_Ttmaccount *TtmaccountFilterer) WatchWithdraw(opts *bind.WatchOpts, sink chan<- *TtmaccountWithdraw, receiver []common.Address) (event.Subscription, error) {

	var receiverRule []interface{}
	for _, receiverItem := range receiver {
		receiverRule = append(receiverRule, receiverItem)
	}

	logs, sub, err := _Ttmaccount.contract.WatchLogs(opts, "Withdraw", receiverRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(TtmaccountWithdraw)
				if err := _Ttmaccount.contract.UnpackLog(event, "Withdraw", log); err != nil {
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

// ParseWithdraw is a log parse operation binding the contract event 0x884edad9ce6fa2440d8a54cc123490eb96d2768479d49ff9c7366125a9424364.
//
// Solidity: event Withdraw(address indexed receiver, uint256 amount)
func (_Ttmaccount *TtmaccountFilterer) ParseWithdraw(log types.Log) (*TtmaccountWithdraw, error) {
	event := new(TtmaccountWithdraw)
	if err := _Ttmaccount.contract.UnpackLog(event, "Withdraw", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}
