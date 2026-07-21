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
	ABI: "[{\"inputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"constructor\"},{\"inputs\":[],\"name\":\"AccessControlBadConfirmation\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"},{\"internalType\":\"bytes32\",\"name\":\"neededRole\",\"type\":\"bytes32\"}],\"name\":\"AccessControlUnauthorizedAccount\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"target\",\"type\":\"address\"}],\"name\":\"AddressEmptyCode\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"serviceHash\",\"type\":\"bytes32\"},{\"internalType\":\"string\",\"name\":\"capability\",\"type\":\"string\"}],\"name\":\"CapabilityDoesNotExist\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"implementation\",\"type\":\"address\"}],\"name\":\"ERC1967InvalidImplementation\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"ERC1967NonPayable\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"FailedCall\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"limit\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"period\",\"type\":\"uint256\"}],\"name\":\"GasMoneyValueOutOfRange\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"balance\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"needed\",\"type\":\"uint256\"}],\"name\":\"InsufficientBalance\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"InvalidInitialization\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"NotInitializing\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"token\",\"type\":\"address\"}],\"name\":\"PaymentTokenAlreadyExists\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"token\",\"type\":\"address\"}],\"name\":\"PaymentTokenDoesNotExist\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"pubKeyAddress\",\"type\":\"address\"}],\"name\":\"PublicKeyAlreadyExists\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"pubKeyAddress\",\"type\":\"address\"}],\"name\":\"PublicKeyDoesNotExist\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"ReentrancyGuardReentrantCall\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"token\",\"type\":\"address\"}],\"name\":\"SafeERC20FailedOperation\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"serviceHash\",\"type\":\"bytes32\"}],\"name\":\"ServiceAlreadyExists\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"serviceHash\",\"type\":\"bytes32\"}],\"name\":\"ServiceDoesNotExist\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"ServiceNotRegistered\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"latestImplementation\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"newImplementation\",\"type\":\"address\"}],\"name\":\"TTMAccountImplementationMismatch\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"oldImplementation\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"newImplementation\",\"type\":\"address\"}],\"name\":\"TTMAccountNoUpgradeNeeded\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"TransferToZeroAddress\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"UUPSUnauthorizedCallContext\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"slot\",\"type\":\"bytes32\"}],\"name\":\"UUPSUnsupportedProxiableUUID\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"serviceHash\",\"type\":\"bytes32\"}],\"name\":\"WantedServiceAlreadyExists\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"serviceHash\",\"type\":\"bytes32\"}],\"name\":\"WantedServiceDoesNotExist\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"limit\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"WithdrawalLimitExceeded\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"limit\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"WithdrawalLimitExceededForPeriod\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"ZeroAddress\",\"type\":\"error\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"sender\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"Deposit\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"withdrawer\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"GasMoneyWithdrawal\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"limit\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"period\",\"type\":\"uint256\"}],\"name\":\"GasMoneyWithdrawalUpdated\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint64\",\"name\":\"version\",\"type\":\"uint64\"}],\"name\":\"Initialized\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"bot\",\"type\":\"address\"}],\"name\":\"MessengerBotAdded\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"bot\",\"type\":\"address\"}],\"name\":\"MessengerBotRemoved\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"bool\",\"name\":\"supportsOffChainPayment\",\"type\":\"bool\"}],\"name\":\"OffChainPaymentSupportUpdated\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"token\",\"type\":\"address\"}],\"name\":\"PaymentTokenAdded\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"token\",\"type\":\"address\"}],\"name\":\"PaymentTokenRemoved\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"pubKeyAddress\",\"type\":\"address\"}],\"name\":\"PublicKeyAdded\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"pubKeyAddress\",\"type\":\"address\"}],\"name\":\"PublicKeyRemoved\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"bytes32\",\"name\":\"role\",\"type\":\"bytes32\"},{\"indexed\":true,\"internalType\":\"bytes32\",\"name\":\"previousAdminRole\",\"type\":\"bytes32\"},{\"indexed\":true,\"internalType\":\"bytes32\",\"name\":\"newAdminRole\",\"type\":\"bytes32\"}],\"name\":\"RoleAdminChanged\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"bytes32\",\"name\":\"role\",\"type\":\"bytes32\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"sender\",\"type\":\"address\"}],\"name\":\"RoleGranted\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"bytes32\",\"name\":\"role\",\"type\":\"bytes32\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"sender\",\"type\":\"address\"}],\"name\":\"RoleRevoked\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"bytes32\",\"name\":\"serviceHash\",\"type\":\"bytes32\"}],\"name\":\"ServiceAdded\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"bytes32\",\"name\":\"serviceHash\",\"type\":\"bytes32\"}],\"name\":\"ServiceCapabilitiesUpdated\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"bytes32\",\"name\":\"serviceHash\",\"type\":\"bytes32\"},{\"indexed\":false,\"internalType\":\"string\",\"name\":\"capability\",\"type\":\"string\"}],\"name\":\"ServiceCapabilityAdded\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"bytes32\",\"name\":\"serviceHash\",\"type\":\"bytes32\"},{\"indexed\":false,\"internalType\":\"string\",\"name\":\"capability\",\"type\":\"string\"}],\"name\":\"ServiceCapabilityRemoved\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"bytes32\",\"name\":\"serviceHash\",\"type\":\"bytes32\"}],\"name\":\"ServiceRemoved\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"bytes32\",\"name\":\"serviceHash\",\"type\":\"bytes32\"},{\"indexed\":false,\"internalType\":\"bool\",\"name\":\"restrictedRate\",\"type\":\"bool\"}],\"name\":\"ServiceRestrictedRateUpdated\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"oldImplementation\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"newImplementation\",\"type\":\"address\"}],\"name\":\"TTMAccountUpgraded\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"implementation\",\"type\":\"address\"}],\"name\":\"Upgraded\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"bytes32\",\"name\":\"serviceHash\",\"type\":\"bytes32\"}],\"name\":\"WantedServiceAdded\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"bytes32\",\"name\":\"serviceHash\",\"type\":\"bytes32\"}],\"name\":\"WantedServiceRemoved\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"receiver\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"Withdraw\",\"type\":\"event\"},{\"inputs\":[],\"name\":\"BOOKING_OPERATOR_ROLE\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"BOT_ADMIN_ROLE\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"DEFAULT_ADMIN_ROLE\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"GAS_WITHDRAWER_ROLE\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"MESSENGER_BOT_ROLE\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"SERVICE_ADMIN_ROLE\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"UPGRADER_ROLE\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"UPGRADE_INTERFACE_VERSION\",\"outputs\":[{\"internalType\":\"string\",\"name\":\"\",\"type\":\"string\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"WITHDRAWER_ROLE\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"tokenId\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"refundAmount\",\"type\":\"uint256\"}],\"name\":\"acceptCancellation\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"bot\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"gasMoney\",\"type\":\"uint256\"}],\"name\":\"addMessengerBot\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"pubKeyAddress\",\"type\":\"address\"},{\"internalType\":\"bytes\",\"name\":\"data\",\"type\":\"bytes\"}],\"name\":\"addPublicKey\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"serviceHash\",\"type\":\"bytes32\"},{\"internalType\":\"bool\",\"name\":\"restrictedRate\",\"type\":\"bool\"},{\"internalType\":\"string[]\",\"name\":\"capabilities\",\"type\":\"string[]\"}],\"name\":\"addService\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"serviceHash\",\"type\":\"bytes32\"},{\"internalType\":\"string\",\"name\":\"capability\",\"type\":\"string\"}],\"name\":\"addServiceCapability\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_supportedToken\",\"type\":\"address\"}],\"name\":\"addSupportedToken\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32[]\",\"name\":\"serviceHashes\",\"type\":\"bytes32[]\"}],\"name\":\"addWantedServices\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"tokenId\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"expectedPrice\",\"type\":\"uint256\"},{\"internalType\":\"contractIERC20\",\"name\":\"expectedPaymentToken\",\"type\":\"address\"}],\"name\":\"buyBookingToken\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"tokenId\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"refundAmount\",\"type\":\"uint256\"},{\"internalType\":\"uint16\",\"name\":\"counterReason\",\"type\":\"uint16\"},{\"internalType\":\"uint16\",\"name\":\"counterReasonVersion\",\"type\":\"uint16\"}],\"name\":\"counterCancellation\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"tokenId\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"refundAmount\",\"type\":\"uint256\"}],\"name\":\"finalizeCancellation\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"getAllServiceHashes\",\"outputs\":[{\"internalType\":\"bytes32[]\",\"name\":\"serviceHashes\",\"type\":\"bytes32[]\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"getBookingTokenAddress\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"getGasMoneyWithdrawal\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"withdrawalLimit\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"withdrawalPeriod\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"}],\"name\":\"getGasMoneyWithdrawalForAccount\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"periodStart\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"withdrawnAmount\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"getManagerAddress\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"pubKeyAddress\",\"type\":\"address\"}],\"name\":\"getPublicKey\",\"outputs\":[{\"internalType\":\"bytes\",\"name\":\"data\",\"type\":\"bytes\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"getPublicKeysAddresses\",\"outputs\":[{\"internalType\":\"address[]\",\"name\":\"pubKeyAddresses\",\"type\":\"address[]\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"role\",\"type\":\"bytes32\"}],\"name\":\"getRoleAdmin\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"role\",\"type\":\"bytes32\"},{\"internalType\":\"uint256\",\"name\":\"index\",\"type\":\"uint256\"}],\"name\":\"getRoleMember\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"role\",\"type\":\"bytes32\"}],\"name\":\"getRoleMemberCount\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"role\",\"type\":\"bytes32\"}],\"name\":\"getRoleMembers\",\"outputs\":[{\"internalType\":\"address[]\",\"name\":\"\",\"type\":\"address[]\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"serviceHash\",\"type\":\"bytes32\"}],\"name\":\"getService\",\"outputs\":[{\"components\":[{\"internalType\":\"bool\",\"name\":\"_restrictedRate\",\"type\":\"bool\"},{\"internalType\":\"string[]\",\"name\":\"_capabilities\",\"type\":\"string[]\"}],\"internalType\":\"structPartnerConfiguration.Service\",\"name\":\"service\",\"type\":\"tuple\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"serviceHash\",\"type\":\"bytes32\"}],\"name\":\"getServiceCapabilities\",\"outputs\":[{\"internalType\":\"string[]\",\"name\":\"capabilities\",\"type\":\"string[]\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"serviceHash\",\"type\":\"bytes32\"}],\"name\":\"getServiceRestrictedRate\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"restrictedRate\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"getSupportedServices\",\"outputs\":[{\"internalType\":\"bytes32[]\",\"name\":\"serviceHashes\",\"type\":\"bytes32[]\"},{\"components\":[{\"internalType\":\"bool\",\"name\":\"_restrictedRate\",\"type\":\"bool\"},{\"internalType\":\"string[]\",\"name\":\"_capabilities\",\"type\":\"string[]\"}],\"internalType\":\"structPartnerConfiguration.Service[]\",\"name\":\"services\",\"type\":\"tuple[]\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"offset\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"limit\",\"type\":\"uint256\"}],\"name\":\"getSupportedServicesSlice\",\"outputs\":[{\"internalType\":\"bytes32[]\",\"name\":\"serviceHashes\",\"type\":\"bytes32[]\"},{\"components\":[{\"internalType\":\"bool\",\"name\":\"_restrictedRate\",\"type\":\"bool\"},{\"internalType\":\"string[]\",\"name\":\"_capabilities\",\"type\":\"string[]\"}],\"internalType\":\"structPartnerConfiguration.Service[]\",\"name\":\"services\",\"type\":\"tuple[]\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"getSupportedTokens\",\"outputs\":[{\"internalType\":\"address[]\",\"name\":\"tokens\",\"type\":\"address[]\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"getWantedServiceHashes\",\"outputs\":[{\"internalType\":\"bytes32[]\",\"name\":\"serviceHashes\",\"type\":\"bytes32[]\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"role\",\"type\":\"bytes32\"},{\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"}],\"name\":\"grantRole\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"role\",\"type\":\"bytes32\"},{\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"}],\"name\":\"hasRole\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"manager\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"bookingToken\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"defaultAdmin\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"upgrader\",\"type\":\"address\"}],\"name\":\"initialize\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"tokenId\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"refundAmount\",\"type\":\"uint256\"},{\"internalType\":\"uint16\",\"name\":\"cancellationReason\",\"type\":\"uint16\"},{\"internalType\":\"uint16\",\"name\":\"cancellationReasonVersion\",\"type\":\"uint16\"}],\"name\":\"initiateCancellation\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"bot\",\"type\":\"address\"}],\"name\":\"isBotAllowed\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"serviceHash\",\"type\":\"bytes32\"}],\"name\":\"isServiceSupported\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"reservedFor\",\"type\":\"address\"},{\"internalType\":\"string\",\"name\":\"uri\",\"type\":\"string\"},{\"internalType\":\"uint256\",\"name\":\"expirationTimestamp\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"price\",\"type\":\"uint256\"},{\"internalType\":\"contractIERC20\",\"name\":\"paymentToken\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"offchainPaymentCurrency\",\"type\":\"uint256\"},{\"internalType\":\"bool\",\"name\":\"cancellable\",\"type\":\"bool\"}],\"name\":\"mintBookingToken\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"offChainPaymentSupported\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"},{\"internalType\":\"bytes\",\"name\":\"\",\"type\":\"bytes\"}],\"name\":\"onERC721Received\",\"outputs\":[{\"internalType\":\"bytes4\",\"name\":\"\",\"type\":\"bytes4\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"proxiableUUID\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"tokenId\",\"type\":\"uint256\"}],\"name\":\"recordExpiration\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"tokenId\",\"type\":\"uint256\"},{\"internalType\":\"uint16\",\"name\":\"rejectionReason\",\"type\":\"uint16\"},{\"internalType\":\"uint16\",\"name\":\"rejectionReasonVersion\",\"type\":\"uint16\"}],\"name\":\"rejectCancellation\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"removeAllServices\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"bot\",\"type\":\"address\"}],\"name\":\"removeMessengerBot\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"pubKeyAddress\",\"type\":\"address\"}],\"name\":\"removePublicKey\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"serviceHash\",\"type\":\"bytes32\"}],\"name\":\"removeService\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"serviceHash\",\"type\":\"bytes32\"},{\"internalType\":\"string\",\"name\":\"capability\",\"type\":\"string\"}],\"name\":\"removeServiceCapability\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_supportedToken\",\"type\":\"address\"}],\"name\":\"removeSupportedToken\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32[]\",\"name\":\"serviceHashes\",\"type\":\"bytes32[]\"}],\"name\":\"removeWantedServices\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"role\",\"type\":\"bytes32\"},{\"internalType\":\"address\",\"name\":\"callerConfirmation\",\"type\":\"address\"}],\"name\":\"renounceRole\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"role\",\"type\":\"bytes32\"},{\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"}],\"name\":\"revokeRole\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"limit\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"period\",\"type\":\"uint256\"}],\"name\":\"setGasMoneyWithdrawal\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bool\",\"name\":\"_isSupported\",\"type\":\"bool\"}],\"name\":\"setOffChainPaymentSupported\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"serviceHash\",\"type\":\"bytes32\"},{\"internalType\":\"string[]\",\"name\":\"capabilities\",\"type\":\"string[]\"}],\"name\":\"setServiceCapabilities\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"serviceHash\",\"type\":\"bytes32\"},{\"internalType\":\"bool\",\"name\":\"restrictedRate\",\"type\":\"bool\"}],\"name\":\"setServiceRestrictedRate\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes4\",\"name\":\"interfaceId\",\"type\":\"bytes4\"}],\"name\":\"supportsInterface\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"contractIERC20\",\"name\":\"token\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"transferERC20\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"contractIERC721\",\"name\":\"token\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"tokenId\",\"type\":\"uint256\"}],\"name\":\"transferERC721\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"newImplementation\",\"type\":\"address\"},{\"internalType\":\"bytes\",\"name\":\"data\",\"type\":\"bytes\"}],\"name\":\"upgradeToAndCall\",\"outputs\":[],\"stateMutability\":\"payable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"addresspayable\",\"name\":\"recipient\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"withdraw\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"tokenId\",\"type\":\"uint256\"},{\"internalType\":\"uint16\",\"name\":\"reason\",\"type\":\"uint16\"},{\"internalType\":\"uint16\",\"name\":\"reasonVersion\",\"type\":\"uint16\"}],\"name\":\"withdrawCancellation\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"withdrawGasMoney\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"stateMutability\":\"payable\",\"type\":\"receive\"}]",
	Bin: "0x60a06040523060805234801562000014575f80fd5b506200001f62000025565b620000d9565b7ff0c57e16840df040f15088dc2f81fe391c3923bec73e23a9662efc9c229c6a00805468010000000000000000900460ff1615620000765760405163f92ee8a960e01b815260040160405180910390fd5b80546001600160401b0390811614620000d65780546001600160401b0319166001600160401b0390811782556040519081527fc7f505b2f371ae2175ee4913f4499e1f2633a7b5936321eed1cdaeb6115181d29060200160405180910390a15b50565b60805161527a620001005f395f8181612b7101528181612b9a0152612fb8015261527a5ff3fe608060405260043610610438575f3560e01c8063857cdbb811610237578063cd9ef9141161013c578063e7bfce9a116100b7578063f6b5174011610087578063f7e45f091161006d578063f7e45f0914610ea4578063f8c8765e14610ec3578063f8e191ac14610ee2575f80fd5b8063f6b5174014610e52578063f72c0d8b14610e71575f80fd5b8063e7bfce9a14610d7a578063ea79d07a14610d99578063ee3b641f14610dad578063f3fef3a314610e33575f80fd5b8063d547741f1161010c578063e0b78add116100f2578063e0b78add14610d1d578063e26a61bb14610d3c578063e5a6725c14610d5b575f80fd5b8063d547741f14610cd2578063da47d85614610cf1575f80fd5b8063cd9ef91414610c60578063d09445c214610c7f578063d3884c3f14610c9f578063d3c7c2c714610cbe575f80fd5b8063a3246ad3116101cc578063be6671881161019c578063c6640e6811610182578063c6640e6814610c03578063ca15c87314610c22578063ccde65dc14610c41575f80fd5b8063be66718814610ba8578063c162d7da14610bc7575f80fd5b8063a3246ad314610b01578063ad3cb1cc14610b2d578063b82923fb14610b75578063bd252c1c14610b89575f80fd5b806391d148541161020757806391d1485414610a4d5780639db5dbe414610ab0578063a217fddf14610acf578063a31aa03914610ae2575f80fd5b8063857cdbb8146109b057806385f438c1146109dc5780638f69347d14610a0f5780639010d07c14610a2e575f80fd5b80634f3f46391161033d5780636fc22cd1116102d257806376319190116102a25780637eec56c7116102885780637eec56c71461095d57806382010fb114610971578063852b3ccb14610990575f80fd5b8063763191901461091f5780637c5d62b31461093e575f80fd5b80636fc22cd11461088f57806372afa328146108ae57806374aa2048146108e157806374fe60e914610900575f80fd5b80635c9889941161030d5780635c988994146107c55780635e07f869146107e4578063658db0af146108105780636d69fcaf14610870575f80fd5b80634f3f46391461073657806351889d6b1461077357806352d1902d146107925780635ae733c9146107a6575f80fd5b8063241bbbfc116103cd578063337462741161039d578063383aba8711610383578063383aba87146106dc57806342072bbd1461070f5780634f1ef28614610723575f80fd5b8063337462741461068a57806336568abe146106bd575f80fd5b8063241bbbfc146105bb578063248a9ca3146105f15780632a1193801461064c5780632f2ff15d1461066b575f80fd5b8063136f50ca11610408578063136f50ca14610519578063150b7a021461053a5780631aca63761461057d5780631c54f0f71461059c575f80fd5b806301ffc9a71461047857806304a3d81e146104ac57806308b1565a146104cd5780630d5e2e16146104fa575f80fd5b366104745760405134815233907fe1fffcc4923d04b559f4d29a8bfc6cda04eb5b0d3c460751c2402c5c5cc9109c9060200160405180910390a2005b5f80fd5b348015610483575f80fd5b506104976104923660046144ec565b610f01565b60405190151581526020015b60405180910390f35b3480156104b7575f80fd5b506104cb6104c636600461457b565b610f2b565b005b3480156104d8575f80fd5b506104ec6104e736600461460c565b610fe4565b6040516104a3929190614725565b348015610505575f80fd5b506104cb6105143660046147a9565b6111b1565b348015610524575f80fd5b5061052d611213565b6040516104a391906147d3565b348015610545575f80fd5b5061056461055436600461486d565b630a85bd0160e11b949350505050565b6040516001600160e01b031990911681526020016104a3565b348015610588575f80fd5b506104cb6105973660046148d5565b611252565b3480156105a7575f80fd5b506104cb6105b636600461460c565b611325565b3480156105c6575f80fd5b507f39fc6f3ebcc11656ef8cd451e9e1f6a26855304f0f1787c5b86c527a1e2d86035460ff16610497565b3480156105fc575f80fd5b5061063e61060b366004614913565b5f9081527f02dd7bc7dec4dceedda775e58dd541e08a116c6c53815c0bd028192f7b626800602052604090206001015490565b6040519081526020016104a3565b348015610657575f80fd5b506104cb61066636600461493b565b6113da565b348015610676575f80fd5b506104cb610685366004614974565b611491565b348015610695575f80fd5b5061063e7fc6ae66da7afbd0becbe49ddf95a9256901d87f1bb5e3d43236030e5e4a585b6d81565b3480156106c8575f80fd5b506104cb6106d7366004614974565b6114da565b3480156106e7575f80fd5b5061063e7fe562760eaa817d85ec1bf58364c4d65adb65d99d113c6785ef9aa66567076c9581565b34801561071a575f80fd5b5061052d611526565b6104cb6107313660046149a2565b61153f565b348015610741575f80fd5b505f80516020615225833981519152546001600160a01b03165b6040516001600160a01b0390911681526020016104a3565b34801561077e575f80fd5b506104cb61078d3660046149ef565b61155e565b34801561079d575f80fd5b5061063e611663565b3480156107b1575f80fd5b506104cb6107c0366004614a19565b611691565b3480156107d0575f80fd5b506104cb6107df366004614913565b6116e2565b3480156107ef575f80fd5b506108036107fe366004614913565b61174a565b6040516104a39190614a47565b34801561081b575f80fd5b507fc7ae2c65fdae475b1bb2dd4079b252a7d893d0822b91a40dd25c763e583f7701546001600160801b03811690600160801b900467ffffffffffffffff165b604080519283526020830191909152016104a3565b34801561087b575f80fd5b506104cb61088a366004614aa9565b611845565b34801561089a575f80fd5b506104cb6108a936600461460c565b611865565b3480156108b9575f80fd5b5061063e7fe9e633d75f5fd429fed1611b4a64c1a1d6af654653894bd204eb46dbd457621a81565b3480156108ec575f80fd5b506104cb6108fb366004614ac4565b611899565b34801561090b575f80fd5b506104cb61091a36600461493b565b611962565b34801561092a575f80fd5b506104cb610939366004614aa9565b6119b1565b348015610949575f80fd5b506104cb610958366004614b8d565b6119d1565b348015610968575f80fd5b506104ec611a2c565b34801561097c575f80fd5b506104cb61098b36600461457b565b611ae5565b34801561099b575f80fd5b5061063e5f8051602061520583398151915281565b3480156109bb575f80fd5b506109cf6109ca366004614aa9565b611b77565b6040516104a39190614be0565b3480156109e7575f80fd5b5061063e7f10dac8c06a04bec0b551627dad28bc00d6516b0caacd1c7b345fcdb5211334e481565b348015610a1a575f80fd5b50610497610a29366004614913565b611c86565b348015610a39575f80fd5b5061075b610a4836600461460c565b611cb5565b348015610a58575f80fd5b50610497610a67366004614974565b5f9182527f02dd7bc7dec4dceedda775e58dd541e08a116c6c53815c0bd028192f7b626800602090815260408084206001600160a01b0393909316845291905290205460ff1690565b348015610abb575f80fd5b506104cb610aca3660046148d5565b611cf5565b348015610ada575f80fd5b5061063e5f81565b348015610aed575f80fd5b506104cb610afc366004614bf2565b611d5a565b348015610b0c575f80fd5b50610b20610b1b366004614913565b611d7a565b6040516104a39190614c0b565b348015610b38575f80fd5b506109cf6040518060400160405280600581526020017f352e302e3000000000000000000000000000000000000000000000000000000081525081565b348015610b80575f80fd5b506104cb611dbd565b348015610b94575f80fd5b50610497610ba3366004614913565b611e5a565b348015610bb3575f80fd5b506104cb610bc236600461460c565b611e64565b348015610bd2575f80fd5b507f17fbd8e22750b6dc617ceae09f600bf2810b53ccc1f790a9a4f78a3f7169b900546001600160a01b031661075b565b348015610c0e575f80fd5b506104cb610c1d366004614aa9565b611eb3565b348015610c2d575f80fd5b5061063e610c3c366004614913565b611f82565b348015610c4c575f80fd5b506104cb610c5b3660046149a2565b611fb9565b348015610c6b575f80fd5b506104cb610c7a366004614c57565b611fda565b348015610c8a575f80fd5b5061063e5f805160206151e583398151915281565b348015610caa575f80fd5b506104cb610cb9366004614913565b6120c4565b348015610cc9575f80fd5b50610b20612112565b348015610cdd575f80fd5b506104cb610cec366004614974565b61214b565b348015610cfc575f80fd5b50610d10610d0b366004614913565b61218e565b6040516104a39190614c8d565b348015610d28575f80fd5b50610497610d37366004614aa9565b6122b8565b348015610d47575f80fd5b506104cb610d56366004614c9f565b6122f7565b348015610d66575f80fd5b506104cb610d75366004614913565b6123a6565b348015610d85575f80fd5b506104cb610d94366004614aa9565b61243b565b348015610da4575f80fd5b50610b2061245b565b348015610db8575f80fd5b5061085b610dc7366004614aa9565b6001600160a01b03165f9081527fc7ae2c65fdae475b1bb2dd4079b252a7d893d0822b91a40dd25c763e583f770060209081526040918290208251808401909352546001600160801b038116808452600160801b90910467ffffffffffffffff16929091018290529091565b348015610e3e575f80fd5b506104cb610e4d3660046149ef565b612494565b348015610e5d575f80fd5b506104cb610e6c366004614a19565b61256d565b348015610e7c575f80fd5b5061063e7f189ab7a9244df0848122154315af71fe140f3db0fe014031783b0946b8c9d2e381565b348015610eaf575f80fd5b506104cb610ebe366004614ac4565b6125be565b348015610ece575f80fd5b506104cb610edd366004614d29565b61260d565b348015610eed575f80fd5b506104cb610efc366004614d82565b6128d0565b5f6001600160e01b03198216630a85bd0160e11b1480610f255750610f2582612920565b92915050565b5f805160206151e5833981519152610f428161295d565b5f5b8251811015610fdf57610f6f838281518110610f6257610f62614dbc565b6020026020010151612967565b610f91838281518110610f8457610f84614dbc565b6020026020010151612a0c565b828181518110610fa357610fa3614dbc565b60200260200101517f7acacfd576383587962277516962c289d19f807be443f4e303ab45ace24931ac60405160405180910390a2600101610f44565b505050565b6060805f610ff0611526565b805190915080861061104c57604080515f80825260208201818152828401909352909190611040565b604080518082019091525f8152606060208201528152602001906001900390816110195790505b509350935050506111aa565b5f6110578783614de4565b905080861115611065578095505b8567ffffffffffffffff81111561107e5761107e614513565b6040519080825280602002602001820160405280156110a7578160200160208202803683370190505b5094508567ffffffffffffffff8111156110c3576110c3614513565b60405190808252806020026020018201604052801561110857816020015b604080518082019091525f8152606060208201528152602001906001900390816110e15790505b5093505f5b868110156111a55783611120828a614df7565b8151811061113057611130614dbc565b602002602001015186828151811061114a5761114a614dbc565b602090810291909101015261118084611163838b614df7565b8151811061117357611173614dbc565b602002602001015161218e565b85828151811061119257611192614dbc565b602090810291909101015260010161110d565b505050505b9250929050565b5f805160206151e58339815191526111c88161295d565b6111d28383612a81565b827f1b76230b39d2d0c1a2a77a90c170190d2280796ed56b280177256ce39df1a66483604051611206911515815260200190565b60405180910390a2505050565b60605f805160206151c583398151915261124c7f39fc6f3ebcc11656ef8cd451e9e1f6a26855304f0f1787c5b86c527a1e2d8609612ab9565b91505090565b7f10dac8c06a04bec0b551627dad28bc00d6516b0caacd1c7b345fcdb5211334e461127c8161295d565b6001600160a01b0383166112a357604051633a954ecd60e21b815260040160405180910390fd5b6040517f42842e0e0000000000000000000000000000000000000000000000000000000081523060048201526001600160a01b038481166024830152604482018490528516906342842e0e906064015f604051808303815f87803b158015611309575f80fd5b505af115801561131b573d5f803e3d5ffd5b5050505050505050565b5f8051602061520583398151915261133c8161295d565b73__$12bd2f62b73a470fe0f6e02c33045f3191$__6321b87f3a6113745f80516020615225833981519152546001600160a01b031690565b6040516001600160e01b031960e084901b1681526001600160a01b03909116600482015260248101869052604481018590526064015f6040518083038186803b1580156113bf575f80fd5b505af41580156113d1573d5f803e3d5ffd5b50505050505050565b5f805160206152058339815191526113f18161295d565b73__$12bd2f62b73a470fe0f6e02c33045f3191$__6307e473166114295f80516020615225833981519152546001600160a01b031690565b6040516001600160e01b031960e084901b1681526001600160a01b0390911660048201526024810187905261ffff8087166044830152851660648201526084015f6040518083038186803b15801561147f575f80fd5b505af415801561131b573d5f803e3d5ffd5b5f8281527f02dd7bc7dec4dceedda775e58dd541e08a116c6c53815c0bd028192f7b62680060205260409020600101546114ca8161295d565b6114d48383612ac5565b50505050565b6001600160a01b038116331461151c576040517f6697b23200000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b610fdf8282612b1a565b60605f805160206151c583398151915261124c81612ab9565b611547612b66565b61155082612c1f565b61155a8282612e12565b5050565b7fc6ae66da7afbd0becbe49ddf95a9256901d87f1bb5e3d43236030e5e4a585b6d6115888161295d565b6001600160a01b0383166115af57604051633a954ecd60e21b815260040160405180910390fd5b6115d97fe9e633d75f5fd429fed1611b4a64c1a1d6af654653894bd204eb46dbd457621a84612ac5565b506115f15f8051602061520583398151915284612ac5565b5061161c7fe562760eaa817d85ec1bf58364c4d65adb65d99d113c6785ef9aa66567076c9584612ac5565b506040516001600160a01b038416907fdb3e11ba26e83d528bf96a2167061674c1ce7777c61376d852d172594a873994905f90a2610fdf6001600160a01b03841683612efa565b5f61166c612fad565b507f360894a13ba1a3210667c828492db98dca3e2076cc3735a920a3ca505d382bbc90565b5f805160206151e58339815191526116a88161295d565b6116b28383612ff6565b827f1cd139430ed537ab9e8086952076cce01edd5ba6e30907af0ffe3709fd3139e6836040516112069190614be0565b6116ea61303a565b7fe562760eaa817d85ec1bf58364c4d65adb65d99d113c6785ef9aa66567076c956117148161295d565b61171d8261309d565b5061174760017f9b779b17422d0df92223018b32b4d1fa46e071723d6817e2486d003becc55f0055565b50565b60605f805160206151c58339815191526117648382613320565b5f838152600282016020908152604080832060010180548251818502810185019093528083529193909284015b82821015611839578382905f5260205f200180546117ae90614e0a565b80601f01602080910402602001604051908101604052809291908181526020018280546117da90614e0a565b80156118255780601f106117fc57610100808354040283529160200191611825565b820191905f5260205f20905b81548152906001019060200180831161180857829003601f168201915b505050505081526020019060010190611791565b50505050915050919050565b5f805160206151e583398151915261185c8161295d565b61155a8261334a565b7fc6ae66da7afbd0becbe49ddf95a9256901d87f1bb5e3d43236030e5e4a585b6d61188f8161295d565b610fdf83836133ff565b5f805160206152058339815191526118b08161295d565b73__$12bd2f62b73a470fe0f6e02c33045f3191$__63fd13a43e6118e85f80516020615225833981519152546001600160a01b031690565b6040516001600160e01b031960e084901b1681526001600160a01b039091166004820152602481018890526044810187905261ffff80871660648301528516608482015260a4015f6040518083038186803b158015611945575f80fd5b505af4158015611957573d5f803e3d5ffd5b505050505050505050565b5f805160206152058339815191526119798161295d565b73__$12bd2f62b73a470fe0f6e02c33045f3191$__63793dddac6114295f80516020615225833981519152546001600160a01b031690565b5f805160206151e58339815191526119c88161295d565b61155a826134e2565b5f805160206151e58339815191526119e88161295d565b6119f184612967565b6119fc848385613597565b60405184907f8f531e5ede07d5741fd086bb787ed399a64704eb757b87cc80cf6635b274e5b5905f90a250505050565b606080611a37611526565b9150815167ffffffffffffffff811115611a5357611a53614513565b604051908082528060200260200182016040528015611a9857816020015b604080518082019091525f815260606020820152815260200190600190039081611a715790505b5090505f5b8251811015611ae057611abb83828151811061117357611173614dbc565b828281518110611acd57611acd614dbc565b6020908102919091010152600101611a9d565b509091565b5f805160206151e5833981519152611afc8161295d565b5f5b8251811015610fdf57611b29838281518110611b1c57611b1c614dbc565b6020026020010151613635565b828181518110611b3b57611b3b614dbc565b60200260200101517ff0dd3de472ddcd75ae2c17728a45801355fb6dd8615a7c53c15504b4279c09be60405160405180910390a2600101611afe565b60605f805160206151c5833981519152611bb17f39fc6f3ebcc11656ef8cd451e9e1f6a26855304f0f1787c5b86c527a1e2d8606846136aa565b611bde5760405163ba650b5f60e01b81526001600160a01b03841660048201526024015b60405180910390fd5b6001600160a01b0383165f90815260088201602052604090208054611c0290614e0a565b80601f0160208091040260200160405190810160405280929190818152602001828054611c2e90614e0a565b8015611c795780601f10611c5057610100808354040283529160200191611c79565b820191905f5260205f20905b815481529060010190602001808311611c5c57829003601f168201915b5050505050915050919050565b5f5f805160206151c5833981519152611c9f8382613320565b5f92835260020160205250604090205460ff1690565b5f8281527fc1f6fe24621ce81ec5827caf0253cadb74709b061630e6b55e82371705932000602081905260408220611ced90846136cb565b949350505050565b7f10dac8c06a04bec0b551627dad28bc00d6516b0caacd1c7b345fcdb5211334e4611d1f8161295d565b6001600160a01b038316611d4657604051633a954ecd60e21b815260040160405180910390fd5b6114d46001600160a01b03851684846136d6565b5f805160206151e5833981519152611d718161295d565b61155a82613756565b5f8181527fc1f6fe24621ce81ec5827caf0253cadb74709b061630e6b55e823717059320006020819052604090912060609190611db690612ab9565b9392505050565b5f805160206151e5833981519152611dd48161295d565b5f611ddd611526565b90505f5b8151811015610fdf57611e0c828281518110611dff57611dff614dbc565b60200260200101516137cd565b818181518110611e1e57611e1e614dbc565b60200260200101517f94da5eeca10d4d6ee8455f99240c10b0c74b0cf5bf754afb81c81e2704b9c42760405160405180910390a2600101611de1565b5f610f258261382e565b5f80516020615205833981519152611e7b8161295d565b73__$12bd2f62b73a470fe0f6e02c33045f3191$__63348e06dd6113745f80516020615225833981519152546001600160a01b031690565b7fc6ae66da7afbd0becbe49ddf95a9256901d87f1bb5e3d43236030e5e4a585b6d611edd8161295d565b611f077fe9e633d75f5fd429fed1611b4a64c1a1d6af654653894bd204eb46dbd457621a83612b1a565b50611f1f5f8051602061520583398151915283612b1a565b50611f4a7fe562760eaa817d85ec1bf58364c4d65adb65d99d113c6785ef9aa66567076c9583612b1a565b506040516001600160a01b038316907fd124523a9cfa28c5dd01826c4fa56192ec7d56859943082e0ca46c3b9dc62913905f90a25050565b5f8181527fc1f6fe24621ce81ec5827caf0253cadb74709b061630e6b55e82371705932000602081905260408220611db690613847565b5f805160206151e5833981519152611fd08161295d565b610fdf8383613850565b611fe261303a565b5f80516020615205833981519152611ff98161295d565b73__$12bd2f62b73a470fe0f6e02c33045f3191$__637adf63b76120315f80516020615225833981519152546001600160a01b031690565b6040516001600160e01b031960e084901b1681526001600160a01b039182166004820152602481018890526044810187905290851660648201526084015f6040518083038186803b158015612084575f80fd5b505af4158015612096573d5f803e3d5ffd5b5050505050610fdf60017f9b779b17422d0df92223018b32b4d1fa46e071723d6817e2486d003becc55f0055565b5f805160206151e58339815191526120db8161295d565b6120e4826137cd565b60405182907f94da5eeca10d4d6ee8455f99240c10b0c74b0cf5bf754afb81c81e2704b9c427905f90a25050565b60605f805160206151c583398151915261124c7f39fc6f3ebcc11656ef8cd451e9e1f6a26855304f0f1787c5b86c527a1e2d8604612ab9565b5f8281527f02dd7bc7dec4dceedda775e58dd541e08a116c6c53815c0bd028192f7b62680060205260409020600101546121848161295d565b6114d48383612b1a565b604080518082019091525f8152606060208201525f805160206151c58339815191526121ba8382613320565b5f838152600282016020908152604080832081518083018352815460ff16151581526001820180548451818702810187019095528085529195929486810194939192919084015b828210156122a9578382905f5260205f2001805461221e90614e0a565b80601f016020809104026020016040519081016040528092919081815260200182805461224a90614e0a565b80156122955780601f1061226c57610100808354040283529160200191612295565b820191905f5260205f20905b81548152906001019060200180831161227857829003601f168201915b505050505081526020019060010190612201565b50505091525090949350505050565b6001600160a01b0381165f9081527f439a7b1b33d79c367c7c6755d8bb3d3ca77b7bca0d68cd209dcbe6cb4f5db4da602052604081205460ff16610f25565b5f8051602061520583398151915261230e8161295d565b73__$12bd2f62b73a470fe0f6e02c33045f3191$__63e4c225696123465f80516020615225833981519152546001600160a01b031690565b8a8a8a8a8a8a8a6040518963ffffffff1660e01b8152600401612370989796959493929190614e42565b5f6040518083038186803b158015612386575f80fd5b505af4158015612398573d5f803e3d5ffd5b505050505050505050505050565b73__$12bd2f62b73a470fe0f6e02c33045f3191$__63c7bffa966123de5f80516020615225833981519152546001600160a01b031690565b6040516001600160e01b031960e084901b1681526001600160a01b039091166004820152602481018490526044015f6040518083038186803b158015612422575f80fd5b505af4158015612434573d5f803e3d5ffd5b5050505050565b5f805160206151e58339815191526124528161295d565b61155a8261392a565b60605f805160206151c583398151915261124c7f39fc6f3ebcc11656ef8cd451e9e1f6a26855304f0f1787c5b86c527a1e2d8606612ab9565b61249c61303a565b7f10dac8c06a04bec0b551627dad28bc00d6516b0caacd1c7b345fcdb5211334e46124c68161295d565b6001600160a01b0383166124ed57604051633a954ecd60e21b815260040160405180910390fd5b6125006001600160a01b03841683612efa565b826001600160a01b03167f884edad9ce6fa2440d8a54cc123490eb96d2768479d49ff9c7366125a94243648360405161253b91815260200190565b60405180910390a25061155a60017f9b779b17422d0df92223018b32b4d1fa46e071723d6817e2486d003becc55f0055565b5f805160206151e58339815191526125848161295d565b61258e83836139e8565b827ffc8d82c9e7e7938446da05458183efa5916c443a2bab87f97f94a8d47742b014836040516112069190614be0565b5f805160206152058339815191526125d58161295d565b73__$12bd2f62b73a470fe0f6e02c33045f3191$__63b54e72d86118e85f80516020615225833981519152546001600160a01b031690565b7ff0c57e16840df040f15088dc2f81fe391c3923bec73e23a9662efc9c229c6a00805468010000000000000000810460ff16159067ffffffffffffffff165f811580156126575750825b90505f8267ffffffffffffffff1660011480156126735750303b155b905081158015612681575080155b156126b8576040517ff92ee8a900000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b845467ffffffffffffffff1916600117855583156126ec57845468ff00000000000000001916680100000000000000001785555b6001600160a01b038916158061270957506001600160a01b038816155b8061271b57506001600160a01b038716155b8061272d57506001600160a01b038616155b15612764576040517fd92e233d00000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b61276c613b48565b612774613b48565b61277c613b50565b612784613b48565b61278e5f88612ac5565b506127a65f805160206151e583398151915288612ac5565b506127d17fc6ae66da7afbd0becbe49ddf95a9256901d87f1bb5e3d43236030e5e4a585b6d88612ac5565b506127fc7f189ab7a9244df0848122154315af71fe140f3db0fe014031783b0946b8c9d2e387612ac5565b507f17fbd8e22750b6dc617ceae09f600bf2810b53ccc1f790a9a4f78a3f7169b90080546001600160a01b038b811673ffffffffffffffffffffffffffffffffffffffff199283161783555f805160206152258339815191528054918c1691909216179055678ac7230489e80000620151806128788282613b60565b505050831561195757845468ff000000000000000019168555604051600181527fc7f505b2f371ae2175ee4913f4499e1f2633a7b5936321eed1cdaeb6115181d29060200160405180910390a1505050505050505050565b5f805160206151e58339815191526128e78161295d565b6128f18383613bf9565b60405183907fa616bfc5bb0e46c6cad727e1b55e3685067e1296d962a7f37017874a27aa0098905f90a2505050565b5f6001600160e01b031982167f5a05180f000000000000000000000000000000000000000000000000000000001480610f255750610f2582613c37565b6117478133613c9d565b7f17fbd8e22750b6dc617ceae09f600bf2810b53ccc1f790a9a4f78a3f7169b900546040517f5a81a626000000000000000000000000000000000000000000000000000000008152600481018390526001600160a01b0390911690635a81a626906024015f60405180830381865afa1580156129e5573d5f803e3d5ffd5b505050506040513d5f823e601f3d908101601f1916820160405261155a9190810190614e9d565b5f805160206151c58339815191525f612a457f39fc6f3ebcc11656ef8cd451e9e1f6a26855304f0f1787c5b86c527a1e2d860984613d29565b905080610fdf576040517f1a1e056900000000000000000000000000000000000000000000000000000000815260048101849052602401611bd5565b5f805160206151c5833981519152612a998382613320565b5f9283526002016020526040909120805460ff1916911515919091179055565b60605f611db683613d34565b5f7fc1f6fe24621ce81ec5827caf0253cadb74709b061630e6b55e8237170593200081612af28585613d8d565b90508015611ced575f858152602083905260409020612b119085613e59565b50949350505050565b5f7fc1f6fe24621ce81ec5827caf0253cadb74709b061630e6b55e8237170593200081612b478585613e6d565b90508015611ced575f858152602083905260409020612b119085613f11565b306001600160a01b037f0000000000000000000000000000000000000000000000000000000000000000161480612bff57507f00000000000000000000000000000000000000000000000000000000000000006001600160a01b0316612bf37f360894a13ba1a3210667c828492db98dca3e2076cc3735a920a3ca505d382bbc546001600160a01b031690565b6001600160a01b031614155b15612c1d5760405163703e46dd60e11b815260040160405180910390fd5b565b7f189ab7a9244df0848122154315af71fe140f3db0fe014031783b0946b8c9d2e3612c498161295d565b5f612c7b7f17fbd8e22750b6dc617ceae09f600bf2810b53ccc1f790a9a4f78a3f7169b900546001600160a01b031690565b6001600160a01b0316639d825bc56040518163ffffffff1660e01b8152600401602060405180830381865afa158015612cb6573d5f803e3d5ffd5b505050506040513d601f19601f82011682018060405250810190612cda9190614f06565b90505f612d0e7f360894a13ba1a3210667c828492db98dca3e2076cc3735a920a3ca505d382bbc546001600160a01b031690565b9050836001600160a01b0316816001600160a01b031603612d6e576040517ffe51a0290000000000000000000000000000000000000000000000000000000081526001600160a01b03808316600483015285166024820152604401611bd5565b816001600160a01b0316846001600160a01b031614612dcc576040517f08811c0c0000000000000000000000000000000000000000000000000000000081526001600160a01b03808416600483015285166024820152604401611bd5565b836001600160a01b0316816001600160a01b03167f897c7778b6095182ea48ee84760832efeae452e4c42d863ea35b271a3aaae75960405160405180910390a350505050565b816001600160a01b03166352d1902d6040518163ffffffff1660e01b8152600401602060405180830381865afa925050508015612e6c575060408051601f3d908101601f19168201909252612e6991810190614f21565b60015b612e9457604051634c9c8ce360e01b81526001600160a01b0383166004820152602401611bd5565b7f360894a13ba1a3210667c828492db98dca3e2076cc3735a920a3ca505d382bbc8114612ef0576040517faa1d49a400000000000000000000000000000000000000000000000000000000815260048101829052602401611bd5565b610fdf8383613f25565b80471015612f3d576040517fcf47918100000000000000000000000000000000000000000000000000000000815247600482015260248101829052604401611bd5565b5f826001600160a01b0316826040515f6040518083038185875af1925050503d805f8114612f86576040519150601f19603f3d011682016040523d82523d5f602084013e612f8b565b606091505b5050905080610fdf5760405163d6bda27560e01b815260040160405180910390fd5b306001600160a01b037f00000000000000000000000000000000000000000000000000000000000000001614612c1d5760405163703e46dd60e11b815260040160405180910390fd5b5f805160206151c583398151915261300e8382613320565b5f8381526002820160209081526040822060019081018054918201815583529120016114d48382614f7c565b7f9b779b17422d0df92223018b32b4d1fa46e071723d6817e2486d003becc55f00805460011901613097576040517f3ee5aeb500000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b60029055565b7fc7ae2c65fdae475b1bb2dd4079b252a7d893d0822b91a40dd25c763e583f7701547fc7ae2c65fdae475b1bb2dd4079b252a7d893d0822b91a40dd25c763e583f7700906001600160801b03168083111561312e576040517fb945e2d80000000000000000000000000000000000000000000000000000000081526004810182905260248101849052604401611bd5565b335f90815260208381526040918290208251808401909352546001600160801b038116835267ffffffffffffffff600160801b918290048116928401839052600186015442936131849390910490911690614df7565b8111156131c1575f825260018401546131b19082908590600160801b900467ffffffffffffffff16613f7a565b67ffffffffffffffff1660208301525b815183906131d99087906001600160801b0316614df7565b111561321b576040517fd54b18870000000000000000000000000000000000000000000000000000000081526004810184905260248101869052604401611bd5565b8151613252906132359087906001600160801b0316614df7565b60018601548590600160801b900467ffffffffffffffff16613fb6565b6001600160801b039081168352335f81815260208781526040909120855181549287015167ffffffffffffffff16600160801b027fffffffffffffffff000000000000000000000000000000000000000000000000909316941693909317179091556132be9086612efa565b60405185815233907fb9ec638398bbdcd0844ca414d8ce760939fa88b9258b9764b3fc6c12ea2605c29060200160405180910390a25050505050565b60017f9b779b17422d0df92223018b32b4d1fa46e071723d6817e2486d003becc55f0055565b61332a8183613fe9565b61155a57604051631e96f6ed60e21b815260048101839052602401611bd5565b5f805160206151c58339815191525f6133837f39fc6f3ebcc11656ef8cd451e9e1f6a26855304f0f1787c5b86c527a1e2d860484613e59565b9050806133c7576040517f50e5f7f20000000000000000000000000000000000000000000000000000000081526001600160a01b0384166004820152602401611bd5565b6040516001600160a01b038416907fa317c10673baf4f03b3c1041bd5ddbb537d0333a86fec3607c75f9dbb630f48f905f90a2505050565b7fc7ae2c65fdae475b1bb2dd4079b252a7d893d0822b91a40dd25c763e583f770061342b838084613fb6565b6001820180546fffffffffffffffffffffffffffffffff19166001600160801b0392909216919091179055613461828481613f7a565b60018201805467ffffffffffffffff92909216600160801b027fffffffffffffffff0000000000000000ffffffffffffffffffffffffffffffff90921691909117905560408051848152602081018490527f8d4925b196ae6b935035a27ed36c6bd9c7a8fbddc7a3f55f493aa8e230be373e910160405180910390a1505050565b5f805160206151c58339815191525f61351b7f39fc6f3ebcc11656ef8cd451e9e1f6a26855304f0f1787c5b86c527a1e2d860484613f11565b90508061355f576040517f54cb99c40000000000000000000000000000000000000000000000000000000081526001600160a01b0384166004820152602401611bd5565b6040516001600160a01b038416907f85a3e72f8dd6db3794f93109c3c5f5b79d6112f6979431c45f98b26134b42af2905f90a2505050565b5f805160206151c58339815191525f6135b08286613d29565b9050806135ec576040517f010f1dd800000000000000000000000000000000000000000000000000000000815260048101869052602401611bd5565b604080518082018252841515815260208082018781525f898152600287018352939093208251815460ff19169015151781559251805192939261131b9260018501920190614416565b5f805160206151c58339815191525f61366e7f39fc6f3ebcc11656ef8cd451e9e1f6a26855304f0f1787c5b86c527a1e2d860984614000565b905080610fdf576040517ffd5cb3e200000000000000000000000000000000000000000000000000000000815260048101849052602401611bd5565b6001600160a01b0381165f9081526001830160205260408120541515611db6565b5f611db6838361400b565b604080516001600160a01b038416602482015260448082018490528251808303909101815260649091019091526020810180517bffffffffffffffffffffffffffffffffffffffffffffffffffffffff167fa9059cbb00000000000000000000000000000000000000000000000000000000179052610fdf908490614031565b7f39fc6f3ebcc11656ef8cd451e9e1f6a26855304f0f1787c5b86c527a1e2d8603805482151560ff19909116811790915560408051918252515f805160206151c5833981519152917fe93ceb76efb130156c6aa39fa4ac986b3f683b6da926496fca3f95ea7fe715e3919081900360200190a15050565b5f805160206151c58339815191525f6137e68284614000565b90508061380957604051631e96f6ed60e21b815260048101849052602401611bd5565b5f8381526002830160205260408120805460ff1916815590612434600183018261446a565b5f5f805160206151c5833981519152611db68184613fe9565b5f610f25825490565b5f805160206151c58339815191525f6138897f39fc6f3ebcc11656ef8cd451e9e1f6a26855304f0f1787c5b86c527a1e2d860685613e59565b9050806138cd576040517fd3083f180000000000000000000000000000000000000000000000000000000081526001600160a01b0385166004820152602401611bd5565b6001600160a01b0384165f90815260088301602052604090206138f08482614f7c565b506040516001600160a01b038516907f928ec246afda323bc23c2815ca3f516e9fc6a7b7179772235c221e132545fe82905f90a250505050565b5f805160206151c58339815191525f6139637f39fc6f3ebcc11656ef8cd451e9e1f6a26855304f0f1787c5b86c527a1e2d860684613f11565b90508061398e5760405163ba650b5f60e01b81526001600160a01b0384166004820152602401611bd5565b6001600160a01b0383165f90815260088301602052604081206139b091614485565b6040516001600160a01b038416907fc5a9b8041ef11732e7dd7043167d8c22db5c7ea99dcd38dce401effacf8a29bf905f90a2505050565b5f805160206151c5833981519152613a008382613320565b5f8381526002820160205260408120600101805490915b81811015613b115784604051602001613a30919061503c565b60405160208183030381529060405280519060200120838281548110613a5857613a58614dbc565b905f5260205f2001604051602001613a709190615057565b6040516020818303038152906040528051906020012003613b095782613a97600184614de4565b81548110613aa757613aa7614dbc565b905f5260205f2001838281548110613ac157613ac1614dbc565b905f5260205f20019081613ad591906150c9565b5082805480613ae657613ae6615198565b600190038181905f5260205f20015f613aff9190614485565b9055505050505050565b600101613a17565b5084846040517fe879f039000000000000000000000000000000000000000000000000000000008152600401611bd59291906151ac565b612c1d6140b6565b613b586140b6565b612c1d61411d565b613b686140b6565b7fc7ae2c65fdae475b1bb2dd4079b252a7d893d0822b91a40dd25c763e583f7700613b94838084613fb6565b6001820180546fffffffffffffffffffffffffffffffff19166001600160801b0392909216919091179055613bca828481613f7a565b8160010160106101000a81548167ffffffffffffffff021916908367ffffffffffffffff160217905550505050565b5f805160206151c5833981519152613c118382613320565b5f838152600282016020908152604090912083516114d492600190920191850190614416565b5f6001600160e01b031982167f7965db0b000000000000000000000000000000000000000000000000000000001480610f2557507f01ffc9a7000000000000000000000000000000000000000000000000000000006001600160e01b0319831614610f25565b5f8281527f02dd7bc7dec4dceedda775e58dd541e08a116c6c53815c0bd028192f7b626800602090815260408083206001600160a01b038516845290915290205460ff1661155a576040517fe2517d3f0000000000000000000000000000000000000000000000000000000081526001600160a01b038216600482015260248101839052604401611bd5565b5f611db68383614125565b6060815f01805480602002602001604051908101604052809291908181526020018280548015613d8157602002820191905f5260205f20905b815481526020019060010190808311613d6d575b50505050509050919050565b5f8281527f02dd7bc7dec4dceedda775e58dd541e08a116c6c53815c0bd028192f7b626800602081815260408084206001600160a01b038616855290915282205460ff16613e50575f848152602082815260408083206001600160a01b03871684529091529020805460ff19166001179055613e063390565b6001600160a01b0316836001600160a01b0316857f2f8788117e7eff1d82e926ec794901d17c78024a50270940304540a733656f0d60405160405180910390a46001915050610f25565b5f915050610f25565b5f611db6836001600160a01b038416614125565b5f8281527f02dd7bc7dec4dceedda775e58dd541e08a116c6c53815c0bd028192f7b626800602081815260408084206001600160a01b038616855290915282205460ff1615613e50575f848152602082815260408083206001600160a01b0387168085529252808320805460ff1916905551339287917ff6391f5c32d9c69d2a47ea670b442974b53935d1edc7fd64eb21e047a839171b9190a46001915050610f25565b5f611db6836001600160a01b038416614171565b613f2e8261424b565b6040516001600160a01b038316907fbc7cd75a20ee27fd9adebab32041f755214dbc6bffa90cc0225b39da2e5c2d3b905f90a2805115613f7257610fdf82826142ce565b61155a614340565b5f67ffffffffffffffff841115613fae5760405163d450716560e01b81526004810184905260248101839052604401611bd5565b509192915050565b5f6001600160801b03841115613fae5760405163d450716560e01b81526004810184905260248101839052604401611bd5565b5f8181526001830160205260408120541515611db6565b5f611db68383614171565b5f825f01828154811061402057614020614dbc565b905f5260205f200154905092915050565b5f8060205f8451602086015f885af180614050576040513d5f823e3d81fd5b50505f513d91508115614067578060011415614074565b6001600160a01b0384163b155b156114d4576040517f5274afe70000000000000000000000000000000000000000000000000000000081526001600160a01b0385166004820152602401611bd5565b7ff0c57e16840df040f15088dc2f81fe391c3923bec73e23a9662efc9c229c6a005468010000000000000000900460ff16612c1d576040517fd7e6bcf800000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b6132fa6140b6565b5f81815260018301602052604081205461416a57508154600181810184555f848152602080822090930184905584548482528286019093526040902091909155610f25565b505f610f25565b5f8181526001830160205260408120548015613e50575f614193600183614de4565b85549091505f906141a690600190614de4565b9050808214614205575f865f0182815481106141c4576141c4614dbc565b905f5260205f200154905080875f0184815481106141e4576141e4614dbc565b5f918252602080832090910192909255918252600188019052604090208390555b855486908061421657614216615198565b600190038181905f5260205f20015f90559055856001015f8681526020019081526020015f205f905560019350505050610f25565b806001600160a01b03163b5f0361428057604051634c9c8ce360e01b81526001600160a01b0382166004820152602401611bd5565b7f360894a13ba1a3210667c828492db98dca3e2076cc3735a920a3ca505d382bbc805473ffffffffffffffffffffffffffffffffffffffff19166001600160a01b0392909216919091179055565b60605f80846001600160a01b0316846040516142ea919061503c565b5f60405180830381855af49150503d805f8114614322576040519150601f19603f3d011682016040523d82523d5f602084013e614327565b606091505b5091509150614337858383614378565b95945050505050565b3415612c1d576040517fb398979f00000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b60608261438d57614388826143ed565b611db6565b81511580156143a457506001600160a01b0384163b155b156143e6576040517f9996b3150000000000000000000000000000000000000000000000000000000081526001600160a01b0385166004820152602401611bd5565b5080611db6565b8051156143fd5780518082602001fd5b60405163d6bda27560e01b815260040160405180910390fd5b828054828255905f5260205f2090810192821561445a579160200282015b8281111561445a578251829061444a9082614f7c565b5091602001919060010190614434565b506144669291506144bc565b5090565b5080545f8255905f5260205f209081019061174791906144bc565b50805461449190614e0a565b5f825580601f106144a0575050565b601f0160209004905f5260205f209081019061174791906144d8565b80821115614466575f6144cf8282614485565b506001016144bc565b5b80821115614466575f81556001016144d9565b5f602082840312156144fc575f80fd5b81356001600160e01b031981168114611db6575f80fd5b634e487b7160e01b5f52604160045260245ffd5b604051601f8201601f1916810167ffffffffffffffff8111828210171561455057614550614513565b604052919050565b5f67ffffffffffffffff82111561457157614571614513565b5060051b60200190565b5f602080838503121561458c575f80fd5b823567ffffffffffffffff8111156145a2575f80fd5b8301601f810185136145b2575f80fd5b80356145c56145c082614558565b614527565b81815260059190911b820183019083810190878311156145e3575f80fd5b928401925b82841015614601578335825292840192908401906145e8565b979650505050505050565b5f806040838503121561461d575f80fd5b50508035926020909101359150565b5f815180845260208085019450602084015f5b8381101561465b5781518752958201959082019060010161463f565b509495945050505050565b5f5b83811015614680578181015183820152602001614668565b50505f910152565b5f815180845261469f816020860160208601614666565b601f01601f19169290920160200192915050565b5f604083018251151584526020808401516040602087015282815180855260608801915060608160051b89010194506020830192505f5b8181101561471857605f19898703018352614706868551614688565b955092840192918401916001016146ea565b5093979650505050505050565b604081525f614737604083018561462c565b6020838203818501528185518084528284019150828160051b8501018388015f5b8381101561478657601f198784030185526147748383516146b3565b94860194925090850190600101614758565b50909998505050505050505050565b803580151581146147a4575f80fd5b919050565b5f80604083850312156147ba575f80fd5b823591506147ca60208401614795565b90509250929050565b602081525f611db6602083018461462c565b6001600160a01b0381168114611747575f80fd5b5f67ffffffffffffffff82111561481257614812614513565b50601f01601f191660200190565b5f82601f83011261482f575f80fd5b813561483d6145c0826147f9565b818152846020838601011115614851575f80fd5b816020850160208301375f918101602001919091529392505050565b5f805f8060808587031215614880575f80fd5b843561488b816147e5565b9350602085013561489b816147e5565b925060408501359150606085013567ffffffffffffffff8111156148bd575f80fd5b6148c987828801614820565b91505092959194509250565b5f805f606084860312156148e7575f80fd5b83356148f2816147e5565b92506020840135614902816147e5565b929592945050506040919091013590565b5f60208284031215614923575f80fd5b5035919050565b803561ffff811681146147a4575f80fd5b5f805f6060848603121561494d575f80fd5b8335925061495d6020850161492a565b915061496b6040850161492a565b90509250925092565b5f8060408385031215614985575f80fd5b823591506020830135614997816147e5565b809150509250929050565b5f80604083850312156149b3575f80fd5b82356149be816147e5565b9150602083013567ffffffffffffffff8111156149d9575f80fd5b6149e585828601614820565b9150509250929050565b5f8060408385031215614a00575f80fd5b8235614a0b816147e5565b946020939093013593505050565b5f8060408385031215614a2a575f80fd5b82359150602083013567ffffffffffffffff8111156149d9575f80fd5b5f60208083016020845280855180835260408601915060408160051b8701019250602087015f5b82811015614a9c57603f19888603018452614a8a858351614688565b94509285019290850190600101614a6e565b5092979650505050505050565b5f60208284031215614ab9575f80fd5b8135611db6816147e5565b5f805f8060808587031215614ad7575f80fd5b8435935060208501359250614aee6040860161492a565b9150614afc6060860161492a565b905092959194509250565b5f82601f830112614b16575f80fd5b81356020614b266145c083614558565b82815260059290921b84018101918181019086841115614b44575f80fd5b8286015b84811015614b8257803567ffffffffffffffff811115614b66575f80fd5b614b748986838b0101614820565b845250918301918301614b48565b509695505050505050565b5f805f60608486031215614b9f575f80fd5b83359250614baf60208501614795565b9150604084013567ffffffffffffffff811115614bca575f80fd5b614bd686828701614b07565b9150509250925092565b602081525f611db66020830184614688565b5f60208284031215614c02575f80fd5b611db682614795565b602080825282518282018190525f9190848201906040850190845b81811015614c4b5783516001600160a01b031683529284019291840191600101614c26565b50909695505050505050565b5f805f60608486031215614c69575f80fd5b83359250602084013591506040840135614c82816147e5565b809150509250925092565b602081525f611db660208301846146b3565b5f805f805f805f60e0888a031215614cb5575f80fd5b8735614cc0816147e5565b9650602088013567ffffffffffffffff811115614cdb575f80fd5b614ce78a828b01614820565b96505060408801359450606088013593506080880135614d06816147e5565b925060a08801359150614d1b60c08901614795565b905092959891949750929550565b5f805f8060808587031215614d3c575f80fd5b8435614d47816147e5565b93506020850135614d57816147e5565b92506040850135614d67816147e5565b91506060850135614d77816147e5565b939692955090935050565b5f8060408385031215614d93575f80fd5b82359150602083013567ffffffffffffffff811115614db0575f80fd5b6149e585828601614b07565b634e487b7160e01b5f52603260045260245ffd5b634e487b7160e01b5f52601160045260245ffd5b81810381811115610f2557610f25614dd0565b80820180821115610f2557610f25614dd0565b600181811c90821680614e1e57607f821691505b602082108103614e3c57634e487b7160e01b5f52602260045260245ffd5b50919050565b5f6101006001600160a01b03808c168452808b166020850152816040850152614e6d8285018b614688565b6060850199909952608084019790975250509290931660a083015260c082015290151560e0909101529392505050565b5f60208284031215614ead575f80fd5b815167ffffffffffffffff811115614ec3575f80fd5b8201601f81018413614ed3575f80fd5b8051614ee16145c0826147f9565b818152856020838501011115614ef5575f80fd5b614337826020830160208601614666565b5f60208284031215614f16575f80fd5b8151611db6816147e5565b5f60208284031215614f31575f80fd5b5051919050565b601f821115610fdf57805f5260205f20601f840160051c81016020851015614f5d5750805b601f840160051c820191505b81811015612434575f8155600101614f69565b815167ffffffffffffffff811115614f9657614f96614513565b614faa81614fa48454614e0a565b84614f38565b602080601f831160018114614fdd575f8415614fc65750858301515b5f19600386901b1c1916600185901b178555615034565b5f85815260208120601f198616915b8281101561500b57888601518255948401946001909101908401614fec565b508582101561502857878501515f19600388901b60f8161c191681555b505060018460011b0185555b505050505050565b5f825161504d818460208701614666565b9190910192915050565b5f80835461506481614e0a565b6001828116801561507c5760018114615091576150bd565b60ff19841687528215158302870194506150bd565b875f526020805f205f5b858110156150b45781548a82015290840190820161509b565b50505082870194505b50929695505050505050565b8181036150d4575050565b6150de8254614e0a565b67ffffffffffffffff8111156150f6576150f6614513565b61510481614fa48454614e0a565b5f601f821160018114615135575f831561511e5750848201545b5f19600385901b1c1916600184901b178455612434565b5f8581526020808220868352908220601f198616925b8381101561516b578286015482556001958601959091019060200161514b565b508583101561518857818501545f19600388901b60f8161c191681555b5050505050600190811b01905550565b634e487b7160e01b5f52603160045260245ffd5b828152604060208201525f611ced604083018461468856fe39fc6f3ebcc11656ef8cd451e9e1f6a26855304f0f1787c5b86c527a1e2d86009a95e87c5af084bf5db8491c3a6515da9dd6da39b24b0eb0af08d7b9cd808d913acdf00ba9ef08b5f2c22768276611b9af078bf6c24fa36b34ec5e9f2eb061fa17fbd8e22750b6dc617ceae09f600bf2810b53ccc1f790a9a4f78a3f7169b901a264697066735822122014f7981089975b283ac6f9370f2be595a80b8218e31f0b9bbdad5f33f6b6e02064736f6c63430008180033",
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

// OffChainPaymentSupported is a free data retrieval call binding the contract method 0x241bbbfc.
//
// Solidity: function offChainPaymentSupported() view returns(bool)
func (_Ttmaccount *TtmaccountCaller) OffChainPaymentSupported(opts *bind.CallOpts) (bool, error) {
	var out []interface{}
	err := _Ttmaccount.contract.Call(opts, &out, "offChainPaymentSupported")

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// OffChainPaymentSupported is a free data retrieval call binding the contract method 0x241bbbfc.
//
// Solidity: function offChainPaymentSupported() view returns(bool)
func (_Ttmaccount *TtmaccountSession) OffChainPaymentSupported() (bool, error) {
	return _Ttmaccount.Contract.OffChainPaymentSupported(&_Ttmaccount.CallOpts)
}

// OffChainPaymentSupported is a free data retrieval call binding the contract method 0x241bbbfc.
//
// Solidity: function offChainPaymentSupported() view returns(bool)
func (_Ttmaccount *TtmaccountCallerSession) OffChainPaymentSupported() (bool, error) {
	return _Ttmaccount.Contract.OffChainPaymentSupported(&_Ttmaccount.CallOpts)
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

// SetOffChainPaymentSupported is a paid mutator transaction binding the contract method 0xa31aa039.
//
// Solidity: function setOffChainPaymentSupported(bool _isSupported) returns()
func (_Ttmaccount *TtmaccountTransactor) SetOffChainPaymentSupported(opts *bind.TransactOpts, _isSupported bool) (*types.Transaction, error) {
	return _Ttmaccount.contract.Transact(opts, "setOffChainPaymentSupported", _isSupported)
}

// SetOffChainPaymentSupported is a paid mutator transaction binding the contract method 0xa31aa039.
//
// Solidity: function setOffChainPaymentSupported(bool _isSupported) returns()
func (_Ttmaccount *TtmaccountSession) SetOffChainPaymentSupported(_isSupported bool) (*types.Transaction, error) {
	return _Ttmaccount.Contract.SetOffChainPaymentSupported(&_Ttmaccount.TransactOpts, _isSupported)
}

// SetOffChainPaymentSupported is a paid mutator transaction binding the contract method 0xa31aa039.
//
// Solidity: function setOffChainPaymentSupported(bool _isSupported) returns()
func (_Ttmaccount *TtmaccountTransactorSession) SetOffChainPaymentSupported(_isSupported bool) (*types.Transaction, error) {
	return _Ttmaccount.Contract.SetOffChainPaymentSupported(&_Ttmaccount.TransactOpts, _isSupported)
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

// TtmaccountOffChainPaymentSupportUpdatedIterator is returned from FilterOffChainPaymentSupportUpdated and is used to iterate over the raw logs and unpacked data for OffChainPaymentSupportUpdated events raised by the Ttmaccount contract.
type TtmaccountOffChainPaymentSupportUpdatedIterator struct {
	Event *TtmaccountOffChainPaymentSupportUpdated // Event containing the contract specifics and raw log

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
func (it *TtmaccountOffChainPaymentSupportUpdatedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(TtmaccountOffChainPaymentSupportUpdated)
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
		it.Event = new(TtmaccountOffChainPaymentSupportUpdated)
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
func (it *TtmaccountOffChainPaymentSupportUpdatedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *TtmaccountOffChainPaymentSupportUpdatedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// TtmaccountOffChainPaymentSupportUpdated represents a OffChainPaymentSupportUpdated event raised by the Ttmaccount contract.
type TtmaccountOffChainPaymentSupportUpdated struct {
	SupportsOffChainPayment bool
	Raw                     types.Log // Blockchain specific contextual infos
}

// FilterOffChainPaymentSupportUpdated is a free log retrieval operation binding the contract event 0xe93ceb76efb130156c6aa39fa4ac986b3f683b6da926496fca3f95ea7fe715e3.
//
// Solidity: event OffChainPaymentSupportUpdated(bool supportsOffChainPayment)
func (_Ttmaccount *TtmaccountFilterer) FilterOffChainPaymentSupportUpdated(opts *bind.FilterOpts) (*TtmaccountOffChainPaymentSupportUpdatedIterator, error) {

	logs, sub, err := _Ttmaccount.contract.FilterLogs(opts, "OffChainPaymentSupportUpdated")
	if err != nil {
		return nil, err
	}
	return &TtmaccountOffChainPaymentSupportUpdatedIterator{contract: _Ttmaccount.contract, event: "OffChainPaymentSupportUpdated", logs: logs, sub: sub}, nil
}

// WatchOffChainPaymentSupportUpdated is a free log subscription operation binding the contract event 0xe93ceb76efb130156c6aa39fa4ac986b3f683b6da926496fca3f95ea7fe715e3.
//
// Solidity: event OffChainPaymentSupportUpdated(bool supportsOffChainPayment)
func (_Ttmaccount *TtmaccountFilterer) WatchOffChainPaymentSupportUpdated(opts *bind.WatchOpts, sink chan<- *TtmaccountOffChainPaymentSupportUpdated) (event.Subscription, error) {

	logs, sub, err := _Ttmaccount.contract.WatchLogs(opts, "OffChainPaymentSupportUpdated")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(TtmaccountOffChainPaymentSupportUpdated)
				if err := _Ttmaccount.contract.UnpackLog(event, "OffChainPaymentSupportUpdated", log); err != nil {
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

// ParseOffChainPaymentSupportUpdated is a log parse operation binding the contract event 0xe93ceb76efb130156c6aa39fa4ac986b3f683b6da926496fca3f95ea7fe715e3.
//
// Solidity: event OffChainPaymentSupportUpdated(bool supportsOffChainPayment)
func (_Ttmaccount *TtmaccountFilterer) ParseOffChainPaymentSupportUpdated(log types.Log) (*TtmaccountOffChainPaymentSupportUpdated, error) {
	event := new(TtmaccountOffChainPaymentSupportUpdated)
	if err := _Ttmaccount.contract.UnpackLog(event, "OffChainPaymentSupportUpdated", log); err != nil {
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
