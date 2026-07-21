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
	ABI: "[{\"inputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"constructor\"},{\"inputs\":[],\"name\":\"AccessControlBadConfirmation\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"},{\"internalType\":\"bytes32\",\"name\":\"neededRole\",\"type\":\"bytes32\"}],\"name\":\"AccessControlUnauthorizedAccount\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"target\",\"type\":\"address\"}],\"name\":\"AddressEmptyCode\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"serviceHash\",\"type\":\"bytes32\"},{\"internalType\":\"string\",\"name\":\"capability\",\"type\":\"string\"}],\"name\":\"CapabilityDoesNotExist\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"implementation\",\"type\":\"address\"}],\"name\":\"ERC1967InvalidImplementation\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"ERC1967NonPayable\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"FailedCall\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"limit\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"period\",\"type\":\"uint256\"}],\"name\":\"GasMoneyValueOutOfRange\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"balance\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"needed\",\"type\":\"uint256\"}],\"name\":\"InsufficientBalance\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"InvalidInitialization\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"NotInitializing\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"token\",\"type\":\"address\"}],\"name\":\"PaymentTokenAlreadyExists\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"token\",\"type\":\"address\"}],\"name\":\"PaymentTokenDoesNotExist\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"pubKeyAddress\",\"type\":\"address\"}],\"name\":\"PublicKeyAlreadyExists\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"pubKeyAddress\",\"type\":\"address\"}],\"name\":\"PublicKeyDoesNotExist\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"ReentrancyGuardReentrantCall\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"token\",\"type\":\"address\"}],\"name\":\"SafeERC20FailedOperation\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"serviceHash\",\"type\":\"bytes32\"}],\"name\":\"ServiceAlreadyExists\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"serviceHash\",\"type\":\"bytes32\"}],\"name\":\"ServiceDoesNotExist\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"latestImplementation\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"newImplementation\",\"type\":\"address\"}],\"name\":\"TTMAccountImplementationMismatch\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"oldImplementation\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"newImplementation\",\"type\":\"address\"}],\"name\":\"TTMAccountNoUpgradeNeeded\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"TransferToZeroAddress\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"UUPSUnauthorizedCallContext\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"slot\",\"type\":\"bytes32\"}],\"name\":\"UUPSUnsupportedProxiableUUID\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"serviceHash\",\"type\":\"bytes32\"}],\"name\":\"WantedServiceAlreadyExists\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"serviceHash\",\"type\":\"bytes32\"}],\"name\":\"WantedServiceDoesNotExist\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"limit\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"WithdrawalLimitExceeded\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"limit\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"WithdrawalLimitExceededForPeriod\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"ZeroAddress\",\"type\":\"error\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"sender\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"Deposit\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"withdrawer\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"GasMoneyWithdrawal\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"limit\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"period\",\"type\":\"uint256\"}],\"name\":\"GasMoneyWithdrawalUpdated\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint64\",\"name\":\"version\",\"type\":\"uint64\"}],\"name\":\"Initialized\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"bot\",\"type\":\"address\"}],\"name\":\"MessengerBotAdded\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"bot\",\"type\":\"address\"}],\"name\":\"MessengerBotRemoved\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"bool\",\"name\":\"supportsOffChainPayment\",\"type\":\"bool\"}],\"name\":\"OffChainPaymentSupportUpdated\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"token\",\"type\":\"address\"}],\"name\":\"PaymentTokenAdded\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"token\",\"type\":\"address\"}],\"name\":\"PaymentTokenRemoved\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"pubKeyAddress\",\"type\":\"address\"}],\"name\":\"PublicKeyAdded\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"pubKeyAddress\",\"type\":\"address\"}],\"name\":\"PublicKeyRemoved\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"bytes32\",\"name\":\"role\",\"type\":\"bytes32\"},{\"indexed\":true,\"internalType\":\"bytes32\",\"name\":\"previousAdminRole\",\"type\":\"bytes32\"},{\"indexed\":true,\"internalType\":\"bytes32\",\"name\":\"newAdminRole\",\"type\":\"bytes32\"}],\"name\":\"RoleAdminChanged\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"bytes32\",\"name\":\"role\",\"type\":\"bytes32\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"sender\",\"type\":\"address\"}],\"name\":\"RoleGranted\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"bytes32\",\"name\":\"role\",\"type\":\"bytes32\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"sender\",\"type\":\"address\"}],\"name\":\"RoleRevoked\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"string\",\"name\":\"serviceName\",\"type\":\"string\"}],\"name\":\"ServiceAdded\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"string\",\"name\":\"serviceName\",\"type\":\"string\"}],\"name\":\"ServiceCapabilitiesUpdated\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"string\",\"name\":\"serviceName\",\"type\":\"string\"},{\"indexed\":false,\"internalType\":\"string\",\"name\":\"capability\",\"type\":\"string\"}],\"name\":\"ServiceCapabilityAdded\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"string\",\"name\":\"serviceName\",\"type\":\"string\"},{\"indexed\":false,\"internalType\":\"string\",\"name\":\"capability\",\"type\":\"string\"}],\"name\":\"ServiceCapabilityRemoved\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"string\",\"name\":\"serviceName\",\"type\":\"string\"}],\"name\":\"ServiceRemoved\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"string\",\"name\":\"serviceName\",\"type\":\"string\"},{\"indexed\":false,\"internalType\":\"bool\",\"name\":\"restrictedRate\",\"type\":\"bool\"}],\"name\":\"ServiceRestrictedRateUpdated\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"oldImplementation\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"newImplementation\",\"type\":\"address\"}],\"name\":\"TTMAccountUpgraded\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"implementation\",\"type\":\"address\"}],\"name\":\"Upgraded\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"string\",\"name\":\"serviceName\",\"type\":\"string\"}],\"name\":\"WantedServiceAdded\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"string\",\"name\":\"serviceName\",\"type\":\"string\"}],\"name\":\"WantedServiceRemoved\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"receiver\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"Withdraw\",\"type\":\"event\"},{\"inputs\":[],\"name\":\"BOOKING_OPERATOR_ROLE\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"BOT_ADMIN_ROLE\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"DEFAULT_ADMIN_ROLE\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"GAS_WITHDRAWER_ROLE\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"MESSENGER_BOT_ROLE\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"SERVICE_ADMIN_ROLE\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"UPGRADER_ROLE\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"UPGRADE_INTERFACE_VERSION\",\"outputs\":[{\"internalType\":\"string\",\"name\":\"\",\"type\":\"string\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"WITHDRAWER_ROLE\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"tokenId\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"refundAmount\",\"type\":\"uint256\"}],\"name\":\"acceptCancellation\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"bot\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"gasMoney\",\"type\":\"uint256\"}],\"name\":\"addMessengerBot\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"pubKeyAddress\",\"type\":\"address\"},{\"internalType\":\"bytes\",\"name\":\"data\",\"type\":\"bytes\"}],\"name\":\"addPublicKey\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"string\",\"name\":\"serviceName\",\"type\":\"string\"},{\"internalType\":\"bool\",\"name\":\"restrictedRate\",\"type\":\"bool\"},{\"internalType\":\"string[]\",\"name\":\"capabilities\",\"type\":\"string[]\"}],\"name\":\"addService\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"string\",\"name\":\"serviceName\",\"type\":\"string\"},{\"internalType\":\"string\",\"name\":\"capability\",\"type\":\"string\"}],\"name\":\"addServiceCapability\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_supportedToken\",\"type\":\"address\"}],\"name\":\"addSupportedToken\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"string[]\",\"name\":\"serviceNames\",\"type\":\"string[]\"}],\"name\":\"addWantedServices\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"tokenId\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"expectedPrice\",\"type\":\"uint256\"},{\"internalType\":\"contractIERC20\",\"name\":\"expectedPaymentToken\",\"type\":\"address\"}],\"name\":\"buyBookingToken\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"tokenId\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"refundAmount\",\"type\":\"uint256\"},{\"internalType\":\"uint16\",\"name\":\"counterReason\",\"type\":\"uint16\"},{\"internalType\":\"uint16\",\"name\":\"counterReasonVersion\",\"type\":\"uint16\"}],\"name\":\"counterCancellation\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"tokenId\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"refundAmount\",\"type\":\"uint256\"}],\"name\":\"finalizeCancellation\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"getAllServiceHashes\",\"outputs\":[{\"internalType\":\"bytes32[]\",\"name\":\"serviceHashes\",\"type\":\"bytes32[]\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"getBookingTokenAddress\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"getGasMoneyWithdrawal\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"withdrawalLimit\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"withdrawalPeriod\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"}],\"name\":\"getGasMoneyWithdrawalForAccount\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"periodStart\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"withdrawnAmount\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"getManagerAddress\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"pubKeyAddress\",\"type\":\"address\"}],\"name\":\"getPublicKey\",\"outputs\":[{\"internalType\":\"bytes\",\"name\":\"data\",\"type\":\"bytes\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"getPublicKeysAddresses\",\"outputs\":[{\"internalType\":\"address[]\",\"name\":\"pubKeyAddresses\",\"type\":\"address[]\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"role\",\"type\":\"bytes32\"}],\"name\":\"getRoleAdmin\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"role\",\"type\":\"bytes32\"},{\"internalType\":\"uint256\",\"name\":\"index\",\"type\":\"uint256\"}],\"name\":\"getRoleMember\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"role\",\"type\":\"bytes32\"}],\"name\":\"getRoleMemberCount\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"role\",\"type\":\"bytes32\"}],\"name\":\"getRoleMembers\",\"outputs\":[{\"internalType\":\"address[]\",\"name\":\"\",\"type\":\"address[]\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"serviceHash\",\"type\":\"bytes32\"}],\"name\":\"getService\",\"outputs\":[{\"components\":[{\"internalType\":\"bool\",\"name\":\"_restrictedRate\",\"type\":\"bool\"},{\"internalType\":\"string[]\",\"name\":\"_capabilities\",\"type\":\"string[]\"}],\"internalType\":\"structPartnerConfiguration.Service\",\"name\":\"service\",\"type\":\"tuple\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"string\",\"name\":\"serviceName\",\"type\":\"string\"}],\"name\":\"getServiceCapabilities\",\"outputs\":[{\"internalType\":\"string[]\",\"name\":\"capabilities\",\"type\":\"string[]\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"serviceHash\",\"type\":\"bytes32\"}],\"name\":\"getServiceCapabilities\",\"outputs\":[{\"internalType\":\"string[]\",\"name\":\"capabilities\",\"type\":\"string[]\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"serviceHash\",\"type\":\"bytes32\"}],\"name\":\"getServiceRestrictedRate\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"restrictedRate\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"string\",\"name\":\"serviceName\",\"type\":\"string\"}],\"name\":\"getServiceRestrictedRate\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"restrictedRate\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"getSupportedServices\",\"outputs\":[{\"internalType\":\"string[]\",\"name\":\"serviceNames\",\"type\":\"string[]\"},{\"components\":[{\"internalType\":\"bool\",\"name\":\"_restrictedRate\",\"type\":\"bool\"},{\"internalType\":\"string[]\",\"name\":\"_capabilities\",\"type\":\"string[]\"}],\"internalType\":\"structPartnerConfiguration.Service[]\",\"name\":\"services\",\"type\":\"tuple[]\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"getSupportedTokens\",\"outputs\":[{\"internalType\":\"address[]\",\"name\":\"tokens\",\"type\":\"address[]\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"getWantedServiceHashes\",\"outputs\":[{\"internalType\":\"bytes32[]\",\"name\":\"serviceHashes\",\"type\":\"bytes32[]\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"getWantedServices\",\"outputs\":[{\"internalType\":\"string[]\",\"name\":\"serviceNames\",\"type\":\"string[]\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"role\",\"type\":\"bytes32\"},{\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"}],\"name\":\"grantRole\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"role\",\"type\":\"bytes32\"},{\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"}],\"name\":\"hasRole\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"manager\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"bookingToken\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"defaultAdmin\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"upgrader\",\"type\":\"address\"}],\"name\":\"initialize\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"tokenId\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"refundAmount\",\"type\":\"uint256\"},{\"internalType\":\"uint16\",\"name\":\"cancellationReason\",\"type\":\"uint16\"},{\"internalType\":\"uint16\",\"name\":\"cancellationReasonVersion\",\"type\":\"uint16\"}],\"name\":\"initiateCancellation\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"bot\",\"type\":\"address\"}],\"name\":\"isBotAllowed\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"string\",\"name\":\"serviceName\",\"type\":\"string\"}],\"name\":\"isServiceSupported\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"reservedFor\",\"type\":\"address\"},{\"internalType\":\"string\",\"name\":\"uri\",\"type\":\"string\"},{\"internalType\":\"uint256\",\"name\":\"expirationTimestamp\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"price\",\"type\":\"uint256\"},{\"internalType\":\"contractIERC20\",\"name\":\"paymentToken\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"offchainPaymentCurrency\",\"type\":\"uint256\"},{\"internalType\":\"bool\",\"name\":\"cancellable\",\"type\":\"bool\"}],\"name\":\"mintBookingToken\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"offChainPaymentSupported\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"},{\"internalType\":\"bytes\",\"name\":\"\",\"type\":\"bytes\"}],\"name\":\"onERC721Received\",\"outputs\":[{\"internalType\":\"bytes4\",\"name\":\"\",\"type\":\"bytes4\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"proxiableUUID\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"tokenId\",\"type\":\"uint256\"}],\"name\":\"recordExpiration\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"tokenId\",\"type\":\"uint256\"},{\"internalType\":\"uint16\",\"name\":\"rejectionReason\",\"type\":\"uint16\"},{\"internalType\":\"uint16\",\"name\":\"rejectionReasonVersion\",\"type\":\"uint16\"}],\"name\":\"rejectCancellation\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"removeAllServices\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"bot\",\"type\":\"address\"}],\"name\":\"removeMessengerBot\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"pubKeyAddress\",\"type\":\"address\"}],\"name\":\"removePublicKey\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"string\",\"name\":\"serviceName\",\"type\":\"string\"}],\"name\":\"removeService\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"string\",\"name\":\"serviceName\",\"type\":\"string\"},{\"internalType\":\"string\",\"name\":\"capability\",\"type\":\"string\"}],\"name\":\"removeServiceCapability\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_supportedToken\",\"type\":\"address\"}],\"name\":\"removeSupportedToken\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"string[]\",\"name\":\"serviceNames\",\"type\":\"string[]\"}],\"name\":\"removeWantedServices\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"role\",\"type\":\"bytes32\"},{\"internalType\":\"address\",\"name\":\"callerConfirmation\",\"type\":\"address\"}],\"name\":\"renounceRole\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"role\",\"type\":\"bytes32\"},{\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"}],\"name\":\"revokeRole\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"limit\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"period\",\"type\":\"uint256\"}],\"name\":\"setGasMoneyWithdrawal\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bool\",\"name\":\"_isSupported\",\"type\":\"bool\"}],\"name\":\"setOffChainPaymentSupported\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"string\",\"name\":\"serviceName\",\"type\":\"string\"},{\"internalType\":\"string[]\",\"name\":\"capabilities\",\"type\":\"string[]\"}],\"name\":\"setServiceCapabilities\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"string\",\"name\":\"serviceName\",\"type\":\"string\"},{\"internalType\":\"bool\",\"name\":\"restrictedRate\",\"type\":\"bool\"}],\"name\":\"setServiceRestrictedRate\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes4\",\"name\":\"interfaceId\",\"type\":\"bytes4\"}],\"name\":\"supportsInterface\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"contractIERC20\",\"name\":\"token\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"transferERC20\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"contractIERC721\",\"name\":\"token\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"tokenId\",\"type\":\"uint256\"}],\"name\":\"transferERC721\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"newImplementation\",\"type\":\"address\"},{\"internalType\":\"bytes\",\"name\":\"data\",\"type\":\"bytes\"}],\"name\":\"upgradeToAndCall\",\"outputs\":[],\"stateMutability\":\"payable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"addresspayable\",\"name\":\"recipient\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"withdraw\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"tokenId\",\"type\":\"uint256\"},{\"internalType\":\"uint16\",\"name\":\"reason\",\"type\":\"uint16\"},{\"internalType\":\"uint16\",\"name\":\"reasonVersion\",\"type\":\"uint16\"}],\"name\":\"withdrawCancellation\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"withdrawGasMoney\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"stateMutability\":\"payable\",\"type\":\"receive\"}]",
	Bin: "0x60a06040523060805234801562000014575f80fd5b506200001f62000025565b620000d9565b7ff0c57e16840df040f15088dc2f81fe391c3923bec73e23a9662efc9c229c6a00805468010000000000000000900460ff1615620000765760405163f92ee8a960e01b815260040160405180910390fd5b80546001600160401b0390811614620000d65780546001600160401b0319166001600160401b0390811782556040519081527fc7f505b2f371ae2175ee4913f4499e1f2633a7b5936321eed1cdaeb6115181d29060200160405180910390a15b50565b60805161547c620001005f395f8181612dde01528181612e070152613225015261547c5ff3fe60806040526004361061046c575f3560e01c8063857cdbb811610251578063ccde65dc1161013c578063e7bfce9a116100b7578063f3fef3a311610087578063f72c0d8b1161006d578063f72c0d8b14610f10578063f7e45f0914610f43578063f8c8765e14610f62575f80fd5b8063f3fef3a314610ed2578063f51acaea14610ef1575f80fd5b8063e7bfce9a14610dfa578063ea79d07a14610e19578063ebc20d2014610e2d578063ee3b641f14610e4c575f80fd5b8063d547741f1161010c578063e0b78add116100f2578063e0b78add14610d9d578063e26a61bb14610dbc578063e5a6725c14610ddb575f80fd5b8063d547741f14610d52578063da47d85614610d71575f80fd5b8063ccde65dc14610ce0578063cd9ef91414610cff578063d09445c214610d1e578063d3c7c2c714610d3e575f80fd5b8063a3246ad3116101cc578063b82923fb1161019c578063c162d7da11610182578063c162d7da14610c66578063c6640e6814610ca2578063ca15c87314610cc1575f80fd5b8063b82923fb14610c33578063be66718814610c47575f80fd5b8063a3246ad314610b81578063a7d022f814610bad578063ad3cb1cc14610bcc578063b512463514610c14575f80fd5b80639010d07c116102215780639db5dbe4116102075780639db5dbe414610b30578063a217fddf14610b4f578063a31aa03914610b62575f80fd5b80639010d07c14610aae57806391d1485414610acd575f80fd5b8063857cdbb814610a1157806385f438c114610a3d5780638c20f57414610a705780638f69347d14610a8f575f80fd5b80634f1ef286116103715780636d69fcaf116102ec57806374fe60e9116102bc57806376319190116102a257806376319190146109b05780637eec56c7146109cf578063852b3ccb146109f1575f80fd5b806374fe60e9146109725780637512e55b14610991575f80fd5b80636d69fcaf146108e25780636fc22cd11461090157806372afa3281461092057806374aa204814610953575f80fd5b8063581ed290116103415780635c988994116103275780635c988994146108445780635e07f86914610863578063658db0af14610882575f80fd5b8063581ed2901461080657806358c0a4c214610825575f80fd5b80634f1ef286146107835780634f3f46391461079657806351889d6b146107d357806352d1902d146107f2575f80fd5b8063248a9ca31161040157806333746274116103d1578063383aba87116103b7578063383aba871461071d57806339e4c7051461075057806342072bbd1461076f575f80fd5b806333746274146106cb57806336568abe146106fe575f80fd5b8063248a9ca3146106135780632a1193801461066e5780632f2ff15d1461068d578063319d13f3146106ac575f80fd5b80631aca63761161043c5780631aca63761461057e5780631c54f0f71461059f5780631c5db99e146105be578063241bbbfc146105dd575f80fd5b806301ffc9a7146104ac57806308564c19146104e0578063136f50ca14610501578063150b7a0214610522575f80fd5b366104a85760405134815233907fe1fffcc4923d04b559f4d29a8bfc6cda04eb5b0d3c460751c2402c5c5cc9109c9060200160405180910390a2005b5f80fd5b3480156104b7575f80fd5b506104cb6104c63660046146b9565b610f81565b60405190151581526020015b60405180910390f35b3480156104eb575f80fd5b506104f4610fc4565b6040516104d79190614785565b34801561050c575f80fd5b5061051561107b565b6040516104d79190614797565b34801561052d575f80fd5b5061056561053c3660046148ac565b7f150b7a0200000000000000000000000000000000000000000000000000000000949350505050565b6040516001600160e01b031990911681526020016104d7565b348015610589575f80fd5b5061059d610598366004614914565b6110ba565b005b3480156105aa575f80fd5b5061059d6105b9366004614952565b61118d565b3480156105c9575f80fd5b5061059d6105d8366004614a0c565b611242565b3480156105e8575f80fd5b507f39fc6f3ebcc11656ef8cd451e9e1f6a26855304f0f1787c5b86c527a1e2d86035460ff166104cb565b34801561061e575f80fd5b5061066061062d366004614a3e565b5f9081527f02dd7bc7dec4dceedda775e58dd541e08a116c6c53815c0bd028192f7b626800602052604090206001015490565b6040519081526020016104d7565b348015610679575f80fd5b5061059d610688366004614a6b565b6112f6565b348015610698575f80fd5b5061059d6106a7366004614aa4565b6113ad565b3480156106b7575f80fd5b506104f46106c6366004614ad2565b6113f6565b3480156106d6575f80fd5b506106607fc6ae66da7afbd0becbe49ddf95a9256901d87f1bb5e3d43236030e5e4a585b6d81565b348015610709575f80fd5b5061059d610718366004614aa4565b611404565b348015610728575f80fd5b506106607fe562760eaa817d85ec1bf58364c4d65adb65d99d113c6785ef9aa66567076c9581565b34801561075b575f80fd5b5061059d61076a366004614a0c565b611450565b34801561077a575f80fd5b506105156114ff565b61059d610791366004614b04565b611518565b3480156107a1575f80fd5b505f80516020615427833981519152546001600160a01b03165b6040516001600160a01b0390911681526020016104d7565b3480156107de575f80fd5b5061059d6107ed366004614b51565b611537565b3480156107fd575f80fd5b5061066061163c565b348015610811575f80fd5b5061059d610820366004614b8a565b61166a565b348015610830575f80fd5b506104cb61083f366004614ad2565b6116d7565b34801561084f575f80fd5b5061059d61085e366004614a3e565b6116e9565b34801561086e575f80fd5b506104f461087d366004614a3e565b611751565b34801561088d575f80fd5b507fc7ae2c65fdae475b1bb2dd4079b252a7d893d0822b91a40dd25c763e583f7701546001600160801b03811690600160801b900467ffffffffffffffff165b604080519283526020830191909152016104d7565b3480156108ed575f80fd5b5061059d6108fc366004614bf9565b61184c565b34801561090c575f80fd5b5061059d61091b366004614952565b61186c565b34801561092b575f80fd5b506106607fe9e633d75f5fd429fed1611b4a64c1a1d6af654653894bd204eb46dbd457621a81565b34801561095e575f80fd5b5061059d61096d366004614c14565b6118a0565b34801561097d575f80fd5b5061059d61098c366004614a6b565b611969565b34801561099c575f80fd5b5061059d6109ab366004614c57565b6119b8565b3480156109bb575f80fd5b5061059d6109ca366004614bf9565b611a33565b3480156109da575f80fd5b506109e3611a53565b6040516104d7929190614d1f565b3480156109fc575f80fd5b506106605f8051602061540783398151915281565b348015610a1c575f80fd5b50610a30610a2b366004614bf9565b611ba4565b6040516104d79190614d8f565b348015610a48575f80fd5b506106607f10dac8c06a04bec0b551627dad28bc00d6516b0caacd1c7b345fcdb5211334e481565b348015610a7b575f80fd5b5061059d610a8a366004614c57565b611cb3565b348015610a9a575f80fd5b506104cb610aa9366004614a3e565b611d21565b348015610ab9575f80fd5b506107bb610ac8366004614952565b611d50565b348015610ad8575f80fd5b506104cb610ae7366004614aa4565b5f9182527f02dd7bc7dec4dceedda775e58dd541e08a116c6c53815c0bd028192f7b626800602090815260408084206001600160a01b0393909316845291905290205460ff1690565b348015610b3b575f80fd5b5061059d610b4a366004614914565b611d90565b348015610b5a575f80fd5b506106605f81565b348015610b6d575f80fd5b5061059d610b7c366004614da1565b611df5565b348015610b8c575f80fd5b50610ba0610b9b366004614a3e565b611e15565b6040516104d79190614dba565b348015610bb8575f80fd5b5061059d610bc7366004614dfa565b611e58565b348015610bd7575f80fd5b50610a306040518060400160405280600581526020017f352e302e3000000000000000000000000000000000000000000000000000000081525081565b348015610c1f575f80fd5b506104cb610c2e366004614ad2565b611ec8565b348015610c3e575f80fd5b5061059d611ed5565b348015610c52575f80fd5b5061059d610c61366004614952565b611f7e565b348015610c71575f80fd5b507f17fbd8e22750b6dc617ceae09f600bf2810b53ccc1f790a9a4f78a3f7169b900546001600160a01b03166107bb565b348015610cad575f80fd5b5061059d610cbc366004614bf9565b611fcd565b348015610ccc575f80fd5b50610660610cdb366004614a3e565b61209c565b348015610ceb575f80fd5b5061059d610cfa366004614b04565b6120d3565b348015610d0a575f80fd5b5061059d610d19366004614e45565b6120f4565b348015610d29575f80fd5b506106605f805160206153e783398151915281565b348015610d49575f80fd5b50610ba06121de565b348015610d5d575f80fd5b5061059d610d6c366004614aa4565b612217565b348015610d7c575f80fd5b50610d90610d8b366004614a3e565b61225a565b6040516104d79190614e7b565b348015610da8575f80fd5b506104cb610db7366004614bf9565b612384565b348015610dc7575f80fd5b5061059d610dd6366004614e8d565b6123c3565b348015610de6575f80fd5b5061059d610df5366004614a3e565b612472565b348015610e05575f80fd5b5061059d610e14366004614bf9565b61251f565b348015610e24575f80fd5b50610ba061253f565b348015610e38575f80fd5b5061059d610e47366004614f17565b612578565b348015610e57575f80fd5b506108cd610e66366004614bf9565b6001600160a01b03165f9081527fc7ae2c65fdae475b1bb2dd4079b252a7d893d0822b91a40dd25c763e583f770060209081526040918290208251808401909352546001600160801b038116808452600160801b90910467ffffffffffffffff16929091018290529091565b348015610edd575f80fd5b5061059d610eec366004614b51565b6125e3565b348015610efc575f80fd5b5061059d610f0b366004614ad2565b6126bc565b348015610f1b575f80fd5b506106607f189ab7a9244df0848122154315af71fe140f3db0fe014031783b0946b8c9d2e381565b348015610f4e575f80fd5b5061059d610f5d366004614c14565b612720565b348015610f6d575f80fd5b5061059d610f7c366004614f6d565b61276f565b5f6001600160e01b031982167f5a05180f000000000000000000000000000000000000000000000000000000001480610fbe5750610fbe82612a32565b92915050565b60605f610fcf61107b565b90505f815167ffffffffffffffff811115610fec57610fec6147ee565b60405190808252806020026020018201604052801561101f57816020015b606081526020019060019003908161100a5790505b5090505f5b82518110156110745761104f83828151811061104257611042614fc6565b6020026020010151612a98565b82828151811061106157611061614fc6565b6020908102919091010152600101611024565b5092915050565b60605f805160206153c78339815191526110b47f39fc6f3ebcc11656ef8cd451e9e1f6a26855304f0f1787c5b86c527a1e2d8609612b39565b91505090565b7f10dac8c06a04bec0b551627dad28bc00d6516b0caacd1c7b345fcdb5211334e46110e481612b45565b6001600160a01b03831661110b57604051633a954ecd60e21b815260040160405180910390fd5b6040517f42842e0e0000000000000000000000000000000000000000000000000000000081523060048201526001600160a01b038481166024830152604482018490528516906342842e0e906064015f604051808303815f87803b158015611171575f80fd5b505af1158015611183573d5f803e3d5ffd5b5050505050505050565b5f805160206154078339815191526111a481612b45565b73__$12bd2f62b73a470fe0f6e02c33045f3191$__6321b87f3a6111dc5f80516020615427833981519152546001600160a01b031690565b6040516001600160e01b031960e084901b1681526001600160a01b03909116600482015260248101869052604481018590526064015f6040518083038186803b158015611227575f80fd5b505af4158015611239573d5f803e3d5ffd5b50505050505050565b5f805160206153e783398151915261125981612b45565b5f5b82518110156112f1575f61128784838151811061127a5761127a614fc6565b6020026020010151612b4f565b905061129281612beb565b8382815181106112a4576112a4614fc6565b60200260200101516040516112b99190614fda565b604051908190038120907f50cc5f9d56177aa0de269c136f2d2ffd45d7b66c82f0a82f8f840db54d9801f8905f90a25060010161125b565b505050565b5f8051602061540783398151915261130d81612b45565b73__$12bd2f62b73a470fe0f6e02c33045f3191$__6307e473166113455f80516020615427833981519152546001600160a01b031690565b6040516001600160e01b031960e084901b1681526001600160a01b0390911660048201526024810187905261ffff8087166044830152851660648201526084015f6040518083038186803b15801561139b575f80fd5b505af4158015611183573d5f803e3d5ffd5b5f8281527f02dd7bc7dec4dceedda775e58dd541e08a116c6c53815c0bd028192f7b62680060205260409020600101546113e681612b45565b6113f08383612c60565b50505050565b6060610fbe61087d83612cb5565b6001600160a01b0381163314611446576040517f6697b23200000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b6112f18282612d12565b5f805160206153e783398151915261146781612b45565b5f5b82518110156112f1575f61149584838151811061148857611488614fc6565b6020026020010151612cb5565b90506114a081612d5e565b8382815181106114b2576114b2614fc6565b60200260200101516040516114c79190614fda565b604051908190038120907f0edb7a081e6ff720ad9e97b837c352ef0548c4d09ec421b9b930b1e0c708e39e905f90a250600101611469565b60605f805160206153c78339815191526110b481612b39565b611520612dd3565b61152982612e8c565b611533828261307f565b5050565b7fc6ae66da7afbd0becbe49ddf95a9256901d87f1bb5e3d43236030e5e4a585b6d61156181612b45565b6001600160a01b03831661158857604051633a954ecd60e21b815260040160405180910390fd5b6115b27fe9e633d75f5fd429fed1611b4a64c1a1d6af654653894bd204eb46dbd457621a84612c60565b506115ca5f8051602061540783398151915284612c60565b506115f57fe562760eaa817d85ec1bf58364c4d65adb65d99d113c6785ef9aa66567076c9584612c60565b506040516001600160a01b038416907fdb3e11ba26e83d528bf96a2167061674c1ce7777c61376d852d172594a873994905f90a26112f16001600160a01b03841683613167565b5f61164561321a565b507f360894a13ba1a3210667c828492db98dca3e2076cc3735a920a3ca505d382bbc90565b5f805160206153e783398151915261168181612b45565b61169461168d85612b4f565b8385613263565b836040516116a29190614fda565b604051908190038120907f763f2f41e0c407dd0a7067f44e5468a0db74da9fdb6cd1cb20c7b6dae9375279905f90a250505050565b5f610fbe6116e483612cb5565b613301565b6116f161331a565b7fe562760eaa817d85ec1bf58364c4d65adb65d99d113c6785ef9aa66567076c9561171b81612b45565b6117248261337d565b5061174e60017f9b779b17422d0df92223018b32b4d1fa46e071723d6817e2486d003becc55f0055565b50565b60605f805160206153c783398151915261176b8382613600565b5f838152600282016020908152604080832060010180548251818502810185019093528083529193909284015b82821015611840578382905f5260205f200180546117b590614ff5565b80601f01602080910402602001604051908101604052809291908181526020018280546117e190614ff5565b801561182c5780601f106118035761010080835404028352916020019161182c565b820191905f5260205f20905b81548152906001019060200180831161180f57829003601f168201915b505050505081526020019060010190611798565b50505050915050919050565b5f805160206153e783398151915261186381612b45565b6115338261362a565b7fc6ae66da7afbd0becbe49ddf95a9256901d87f1bb5e3d43236030e5e4a585b6d61189681612b45565b6112f183836136df565b5f805160206154078339815191526118b781612b45565b73__$12bd2f62b73a470fe0f6e02c33045f3191$__63fd13a43e6118ef5f80516020615427833981519152546001600160a01b031690565b6040516001600160e01b031960e084901b1681526001600160a01b039091166004820152602481018890526044810187905261ffff80871660648301528516608482015260a4015f6040518083038186803b15801561194c575f80fd5b505af415801561195e573d5f803e3d5ffd5b505050505050505050565b5f8051602061540783398151915261198081612b45565b73__$12bd2f62b73a470fe0f6e02c33045f3191$__63793dddac6113455f80516020615427833981519152546001600160a01b031690565b5f805160206153e78339815191526119cf81612b45565b6119e16119db84612cb5565b836137c2565b826040516119ef9190614fda565b60405180910390207f498a5f4e6f3921f63e6863032989bdb7bb41e5cf5cbde5437c7322c5c8dc46bf83604051611a269190614d8f565b60405180910390a2505050565b5f805160206153e7833981519152611a4a81612b45565b61153382613806565b6060805f611a5f6114ff565b90505f815167ffffffffffffffff811115611a7c57611a7c6147ee565b604051908082528060200260200182016040528015611aaf57816020015b6060815260200190600190039081611a9a5790505b5090505f825167ffffffffffffffff811115611acd57611acd6147ee565b604051908082528060200260200182016040528015611b1257816020015b604080518082019091525f815260606020820152815260200190600190039081611aeb5790505b5090505f5b8351811015611b9957611b3584828151811061104257611042614fc6565b838281518110611b4757611b47614fc6565b6020026020010181905250611b74848281518110611b6757611b67614fc6565b602002602001015161225a565b828281518110611b8657611b86614fc6565b6020908102919091010152600101611b17565b509094909350915050565b60605f805160206153c7833981519152611bde7f39fc6f3ebcc11656ef8cd451e9e1f6a26855304f0f1787c5b86c527a1e2d8606846138bb565b611c0b5760405163ba650b5f60e01b81526001600160a01b03841660048201526024015b60405180910390fd5b6001600160a01b0383165f90815260088201602052604090208054611c2f90614ff5565b80601f0160208091040260200160405190810160405280929190818152602001828054611c5b90614ff5565b8015611ca65780601f10611c7d57610100808354040283529160200191611ca6565b820191905f5260205f20905b815481529060010190602001808311611c8957829003601f168201915b5050505050915050919050565b5f805160206153e7833981519152611cca81612b45565b611cdc611cd684612cb5565b836138dc565b82604051611cea9190614fda565b60405180910390207fba851faec9e30a9961f0adb49fe025cda6c8d7d0fb9bad99f89c37d05702326483604051611a269190614d8f565b5f5f805160206153c7833981519152611d3a8382613600565b5f92835260020160205250604090205460ff1690565b5f8281527fc1f6fe24621ce81ec5827caf0253cadb74709b061630e6b55e82371705932000602081905260408220611d889084613a3c565b949350505050565b7f10dac8c06a04bec0b551627dad28bc00d6516b0caacd1c7b345fcdb5211334e4611dba81612b45565b6001600160a01b038316611de157604051633a954ecd60e21b815260040160405180910390fd5b6113f06001600160a01b0385168484613a47565b5f805160206153e7833981519152611e0c81612b45565b61153382613ac7565b5f8181527fc1f6fe24621ce81ec5827caf0253cadb74709b061630e6b55e823717059320006020819052604090912060609190611e5190612b39565b9392505050565b5f805160206153e7833981519152611e6f81612b45565b611e81611e7b84612cb5565b83613b3e565b82604051611e8f9190614fda565b6040519081900381208315158252907f23960b931eb4b63e2e47d040f51cc0de6eef2e865639eb674fbb1890ece3a0ab90602001611a26565b5f610fbe610aa983612cb5565b5f805160206153e7833981519152611eec81612b45565b5f611ef5611a53565b5090505f5b81518110156112f157611f20611f1b83838151811061148857611488614fc6565b613b76565b818181518110611f3257611f32614fc6565b6020026020010151604051611f479190614fda565b604051908190038120907f52f6e0779195109314dfb8cf301d33491c63f136afac4c5d4f35aa934b254813905f90a2600101611efa565b5f80516020615407833981519152611f9581612b45565b73__$12bd2f62b73a470fe0f6e02c33045f3191$__63348e06dd6111dc5f80516020615427833981519152546001600160a01b031690565b7fc6ae66da7afbd0becbe49ddf95a9256901d87f1bb5e3d43236030e5e4a585b6d611ff781612b45565b6120217fe9e633d75f5fd429fed1611b4a64c1a1d6af654653894bd204eb46dbd457621a83612d12565b506120395f8051602061540783398151915283612d12565b506120647fe562760eaa817d85ec1bf58364c4d65adb65d99d113c6785ef9aa66567076c9583612d12565b506040516001600160a01b038316907fd124523a9cfa28c5dd01826c4fa56192ec7d56859943082e0ca46c3b9dc62913905f90a25050565b5f8181527fc1f6fe24621ce81ec5827caf0253cadb74709b061630e6b55e82371705932000602081905260408220611e5190613bde565b5f805160206153e78339815191526120ea81612b45565b6112f18383613be7565b6120fc61331a565b5f8051602061540783398151915261211381612b45565b73__$12bd2f62b73a470fe0f6e02c33045f3191$__637adf63b761214b5f80516020615427833981519152546001600160a01b031690565b6040516001600160e01b031960e084901b1681526001600160a01b039182166004820152602481018890526044810187905290851660648201526084015f6040518083038186803b15801561219e575f80fd5b505af41580156121b0573d5f803e3d5ffd5b50505050506112f160017f9b779b17422d0df92223018b32b4d1fa46e071723d6817e2486d003becc55f0055565b60605f805160206153c78339815191526110b47f39fc6f3ebcc11656ef8cd451e9e1f6a26855304f0f1787c5b86c527a1e2d8604612b39565b5f8281527f02dd7bc7dec4dceedda775e58dd541e08a116c6c53815c0bd028192f7b626800602052604090206001015461225081612b45565b6113f08383612d12565b604080518082019091525f8152606060208201525f805160206153c78339815191526122868382613600565b5f838152600282016020908152604080832081518083018352815460ff16151581526001820180548451818702810187019095528085529195929486810194939192919084015b82821015612375578382905f5260205f200180546122ea90614ff5565b80601f016020809104026020016040519081016040528092919081815260200182805461231690614ff5565b80156123615780601f1061233857610100808354040283529160200191612361565b820191905f5260205f20905b81548152906001019060200180831161234457829003601f168201915b5050505050815260200190600101906122cd565b50505091525090949350505050565b6001600160a01b0381165f9081527f439a7b1b33d79c367c7c6755d8bb3d3ca77b7bca0d68cd209dcbe6cb4f5db4da602052604081205460ff16610fbe565b5f805160206154078339815191526123da81612b45565b73__$12bd2f62b73a470fe0f6e02c33045f3191$__63e4c225696124125f80516020615427833981519152546001600160a01b031690565b8a8a8a8a8a8a8a6040518963ffffffff1660e01b815260040161243c98979695949392919061502d565b5f6040518083038186803b158015612452575f80fd5b505af4158015612464573d5f803e3d5ffd5b505050505050505050505050565b5f8051602061540783398151915261248981612b45565b73__$12bd2f62b73a470fe0f6e02c33045f3191$__63c7bffa966124c15f80516020615427833981519152546001600160a01b031690565b6040516001600160e01b031960e084901b1681526001600160a01b039091166004820152602481018590526044015f6040518083038186803b158015612505575f80fd5b505af4158015612517573d5f803e3d5ffd5b505050505050565b5f805160206153e783398151915261253681612b45565b61153382613cc1565b60605f805160206153c78339815191526110b47f39fc6f3ebcc11656ef8cd451e9e1f6a26855304f0f1787c5b86c527a1e2d8606612b39565b5f805160206153e783398151915261258f81612b45565b6125a161259b84612cb5565b83613d7f565b826040516125af9190614fda565b604051908190038120907fd52aef6010d6b6303240865274298b7c5784b14ebf9df788047b34c69c531371905f90a2505050565b6125eb61331a565b7f10dac8c06a04bec0b551627dad28bc00d6516b0caacd1c7b345fcdb5211334e461261581612b45565b6001600160a01b03831661263c57604051633a954ecd60e21b815260040160405180910390fd5b61264f6001600160a01b03841683613167565b826001600160a01b03167f884edad9ce6fa2440d8a54cc123490eb96d2768479d49ff9c7366125a94243648360405161268a91815260200190565b60405180910390a25061153360017f9b779b17422d0df92223018b32b4d1fa46e071723d6817e2486d003becc55f0055565b5f805160206153e78339815191526126d381612b45565b6126df611f1b83612cb5565b816040516126ed9190614fda565b604051908190038120907f52f6e0779195109314dfb8cf301d33491c63f136afac4c5d4f35aa934b254813905f90a25050565b5f8051602061540783398151915261273781612b45565b73__$12bd2f62b73a470fe0f6e02c33045f3191$__63b54e72d86118ef5f80516020615427833981519152546001600160a01b031690565b7ff0c57e16840df040f15088dc2f81fe391c3923bec73e23a9662efc9c229c6a00805468010000000000000000810460ff16159067ffffffffffffffff165f811580156127b95750825b90505f8267ffffffffffffffff1660011480156127d55750303b155b9050811580156127e3575080155b1561281a576040517ff92ee8a900000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b845467ffffffffffffffff19166001178555831561284e57845468ff00000000000000001916680100000000000000001785555b6001600160a01b038916158061286b57506001600160a01b038816155b8061287d57506001600160a01b038716155b8061288f57506001600160a01b038616155b156128c6576040517fd92e233d00000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b6128ce613dbd565b6128d6613dbd565b6128de613dc5565b6128e6613dbd565b6128f05f88612c60565b506129085f805160206153e783398151915288612c60565b506129337fc6ae66da7afbd0becbe49ddf95a9256901d87f1bb5e3d43236030e5e4a585b6d88612c60565b5061295e7f189ab7a9244df0848122154315af71fe140f3db0fe014031783b0946b8c9d2e387612c60565b507f17fbd8e22750b6dc617ceae09f600bf2810b53ccc1f790a9a4f78a3f7169b90080546001600160a01b038b811673ffffffffffffffffffffffffffffffffffffffff199283161783555f805160206154278339815191528054918c1691909216179055678ac7230489e80000620151806129da8282613dd5565b505050831561195e57845468ff000000000000000019168555604051600181527fc7f505b2f371ae2175ee4913f4499e1f2633a7b5936321eed1cdaeb6115181d29060200160405180910390a1505050505050505050565b5f6001600160e01b031982167f7965db0b000000000000000000000000000000000000000000000000000000001480610fbe57507f01ffc9a7000000000000000000000000000000000000000000000000000000006001600160e01b0319831614610fbe565b6060612acb7f17fbd8e22750b6dc617ceae09f600bf2810b53ccc1f790a9a4f78a3f7169b900546001600160a01b031690565b6001600160a01b031663306ade09836040518263ffffffff1660e01b8152600401612af891815260200190565b5f60405180830381865afa158015612b12573d5f803e3d5ffd5b505050506040513d5f823e601f3d908101601f19168201604052610fbe9190810190615088565b60605f611e5183613e6e565b61174e8133613ec7565b5f612b817f17fbd8e22750b6dc617ceae09f600bf2810b53ccc1f790a9a4f78a3f7169b900546001600160a01b031690565b6001600160a01b031663352af39a836040518263ffffffff1660e01b8152600401612bac9190614d8f565b602060405180830381865afa158015612bc7573d5f803e3d5ffd5b505050506040513d601f19601f82011682018060405250810190610fbe91906150f1565b5f805160206153c78339815191525f612c247f39fc6f3ebcc11656ef8cd451e9e1f6a26855304f0f1787c5b86c527a1e2d860984613f53565b9050806112f1576040517f1a1e056900000000000000000000000000000000000000000000000000000000815260048101849052602401611c02565b5f7fc1f6fe24621ce81ec5827caf0253cadb74709b061630e6b55e8237170593200081612c8d8585613f5e565b90508015611d88575f858152602083905260409020612cac908561402a565b50949350505050565b5f612ce77f17fbd8e22750b6dc617ceae09f600bf2810b53ccc1f790a9a4f78a3f7169b900546001600160a01b031690565b6001600160a01b0316631ca0e943836040518263ffffffff1660e01b8152600401612bac9190614d8f565b5f7fc1f6fe24621ce81ec5827caf0253cadb74709b061630e6b55e8237170593200081612d3f858561403e565b90508015611d88575f858152602083905260409020612cac90856140e2565b5f805160206153c78339815191525f612d977f39fc6f3ebcc11656ef8cd451e9e1f6a26855304f0f1787c5b86c527a1e2d8609846140f6565b9050806112f1576040517ffd5cb3e200000000000000000000000000000000000000000000000000000000815260048101849052602401611c02565b306001600160a01b037f0000000000000000000000000000000000000000000000000000000000000000161480612e6c57507f00000000000000000000000000000000000000000000000000000000000000006001600160a01b0316612e607f360894a13ba1a3210667c828492db98dca3e2076cc3735a920a3ca505d382bbc546001600160a01b031690565b6001600160a01b031614155b15612e8a5760405163703e46dd60e11b815260040160405180910390fd5b565b7f189ab7a9244df0848122154315af71fe140f3db0fe014031783b0946b8c9d2e3612eb681612b45565b5f612ee87f17fbd8e22750b6dc617ceae09f600bf2810b53ccc1f790a9a4f78a3f7169b900546001600160a01b031690565b6001600160a01b0316639d825bc56040518163ffffffff1660e01b8152600401602060405180830381865afa158015612f23573d5f803e3d5ffd5b505050506040513d601f19601f82011682018060405250810190612f479190615108565b90505f612f7b7f360894a13ba1a3210667c828492db98dca3e2076cc3735a920a3ca505d382bbc546001600160a01b031690565b9050836001600160a01b0316816001600160a01b031603612fdb576040517ffe51a0290000000000000000000000000000000000000000000000000000000081526001600160a01b03808316600483015285166024820152604401611c02565b816001600160a01b0316846001600160a01b031614613039576040517f08811c0c0000000000000000000000000000000000000000000000000000000081526001600160a01b03808416600483015285166024820152604401611c02565b836001600160a01b0316816001600160a01b03167f897c7778b6095182ea48ee84760832efeae452e4c42d863ea35b271a3aaae75960405160405180910390a350505050565b816001600160a01b03166352d1902d6040518163ffffffff1660e01b8152600401602060405180830381865afa9250505080156130d9575060408051601f3d908101601f191682019092526130d6918101906150f1565b60015b61310157604051634c9c8ce360e01b81526001600160a01b0383166004820152602401611c02565b7f360894a13ba1a3210667c828492db98dca3e2076cc3735a920a3ca505d382bbc811461315d576040517faa1d49a400000000000000000000000000000000000000000000000000000000815260048101829052602401611c02565b6112f18383614101565b804710156131aa576040517fcf47918100000000000000000000000000000000000000000000000000000000815247600482015260248101829052604401611c02565b5f826001600160a01b0316826040515f6040518083038185875af1925050503d805f81146131f3576040519150601f19603f3d011682016040523d82523d5f602084013e6131f8565b606091505b50509050806112f15760405163d6bda27560e01b815260040160405180910390fd5b306001600160a01b037f00000000000000000000000000000000000000000000000000000000000000001614612e8a5760405163703e46dd60e11b815260040160405180910390fd5b5f805160206153c78339815191525f61327c8286613f53565b9050806132b8576040517f010f1dd800000000000000000000000000000000000000000000000000000000815260048101869052602401611c02565b604080518082018252841515815260208082018781525f898152600287018352939093208251815460ff19169015151781559251805192939261118392600185019201906145e7565b5f5f805160206153c7833981519152611e518184614156565b7f9b779b17422d0df92223018b32b4d1fa46e071723d6817e2486d003becc55f00805460011901613377576040517f3ee5aeb500000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b60029055565b7fc7ae2c65fdae475b1bb2dd4079b252a7d893d0822b91a40dd25c763e583f7701547fc7ae2c65fdae475b1bb2dd4079b252a7d893d0822b91a40dd25c763e583f7700906001600160801b03168083111561340e576040517fb945e2d80000000000000000000000000000000000000000000000000000000081526004810182905260248101849052604401611c02565b335f90815260208381526040918290208251808401909352546001600160801b038116835267ffffffffffffffff600160801b918290048116928401839052600186015442936134649390910490911690615137565b8111156134a1575f825260018401546134919082908590600160801b900467ffffffffffffffff1661416d565b67ffffffffffffffff1660208301525b815183906134b99087906001600160801b0316615137565b11156134fb576040517fd54b18870000000000000000000000000000000000000000000000000000000081526004810184905260248101869052604401611c02565b8151613532906135159087906001600160801b0316615137565b60018601548590600160801b900467ffffffffffffffff166141a9565b6001600160801b039081168352335f81815260208781526040909120855181549287015167ffffffffffffffff16600160801b027fffffffffffffffff0000000000000000000000000000000000000000000000009093169416939093171790915561359e9086613167565b60405185815233907fb9ec638398bbdcd0844ca414d8ce760939fa88b9258b9764b3fc6c12ea2605c29060200160405180910390a25050505050565b60017f9b779b17422d0df92223018b32b4d1fa46e071723d6817e2486d003becc55f0055565b61360a8183614156565b61153357604051631e96f6ed60e21b815260048101839052602401611c02565b5f805160206153c78339815191525f6136637f39fc6f3ebcc11656ef8cd451e9e1f6a26855304f0f1787c5b86c527a1e2d86048461402a565b9050806136a7576040517f50e5f7f20000000000000000000000000000000000000000000000000000000081526001600160a01b0384166004820152602401611c02565b6040516001600160a01b038416907fa317c10673baf4f03b3c1041bd5ddbb537d0333a86fec3607c75f9dbb630f48f905f90a2505050565b7fc7ae2c65fdae475b1bb2dd4079b252a7d893d0822b91a40dd25c763e583f770061370b8380846141a9565b6001820180546fffffffffffffffffffffffffffffffff19166001600160801b039290921691909117905561374182848161416d565b60018201805467ffffffffffffffff92909216600160801b027fffffffffffffffff0000000000000000ffffffffffffffffffffffffffffffff90921691909117905560408051848152602081018490527f8d4925b196ae6b935035a27ed36c6bd9c7a8fbddc7a3f55f493aa8e230be373e910160405180910390a1505050565b5f805160206153c78339815191526137da8382613600565b5f8381526002820160209081526040822060019081018054918201815583529120016113f0838261518e565b5f805160206153c78339815191525f61383f7f39fc6f3ebcc11656ef8cd451e9e1f6a26855304f0f1787c5b86c527a1e2d8604846140e2565b905080613883576040517f54cb99c40000000000000000000000000000000000000000000000000000000081526001600160a01b0384166004820152602401611c02565b6040516001600160a01b038416907f85a3e72f8dd6db3794f93109c3c5f5b79d6112f6979431c45f98b26134b42af2905f90a2505050565b6001600160a01b0381165f9081526001830160205260408120541515611e51565b5f805160206153c78339815191526138f48382613600565b5f8381526002820160205260408120600101805490915b81811015613a0557846040516020016139249190614fda565b6040516020818303038152906040528051906020012083828154811061394c5761394c614fc6565b905f5260205f2001604051602001613964919061524a565b60405160208183030381529060405280519060200120036139fd578261398b6001846152bc565b8154811061399b5761399b614fc6565b905f5260205f20018382815481106139b5576139b5614fc6565b905f5260205f200190816139c991906152cf565b50828054806139da576139da61539a565b600190038181905f5260205f20015f6139f3919061463b565b9055505050505050565b60010161390b565b5084846040517fe879f039000000000000000000000000000000000000000000000000000000008152600401611c029291906153ae565b5f611e5183836141dc565b604080516001600160a01b038416602482015260448082018490528251808303909101815260649091019091526020810180517bffffffffffffffffffffffffffffffffffffffffffffffffffffffff167fa9059cbb000000000000000000000000000000000000000000000000000000001790526112f1908490614202565b7f39fc6f3ebcc11656ef8cd451e9e1f6a26855304f0f1787c5b86c527a1e2d8603805482151560ff19909116811790915560408051918252515f805160206153c7833981519152917fe93ceb76efb130156c6aa39fa4ac986b3f683b6da926496fca3f95ea7fe715e3919081900360200190a15050565b5f805160206153c7833981519152613b568382613600565b5f9283526002016020526040909120805460ff1916911515919091179055565b5f805160206153c78339815191525f613b8f82846140f6565b905080613bb257604051631e96f6ed60e21b815260048101849052602401611c02565b5f8381526002830160205260408120805460ff1916815590613bd76001830182614672565b5050505050565b5f610fbe825490565b5f805160206153c78339815191525f613c207f39fc6f3ebcc11656ef8cd451e9e1f6a26855304f0f1787c5b86c527a1e2d86068561402a565b905080613c64576040517fd3083f180000000000000000000000000000000000000000000000000000000081526001600160a01b0385166004820152602401611c02565b6001600160a01b0384165f9081526008830160205260409020613c87848261518e565b506040516001600160a01b038516907f928ec246afda323bc23c2815ca3f516e9fc6a7b7179772235c221e132545fe82905f90a250505050565b5f805160206153c78339815191525f613cfa7f39fc6f3ebcc11656ef8cd451e9e1f6a26855304f0f1787c5b86c527a1e2d8606846140e2565b905080613d255760405163ba650b5f60e01b81526001600160a01b0384166004820152602401611c02565b6001600160a01b0383165f9081526008830160205260408120613d479161463b565b6040516001600160a01b038416907fc5a9b8041ef11732e7dd7043167d8c22db5c7ea99dcd38dce401effacf8a29bf905f90a2505050565b5f805160206153c7833981519152613d978382613600565b5f838152600282016020908152604090912083516113f0926001909201918501906145e7565b612e8a614287565b613dcd614287565b612e8a6142ee565b613ddd614287565b7fc7ae2c65fdae475b1bb2dd4079b252a7d893d0822b91a40dd25c763e583f7700613e098380846141a9565b6001820180546fffffffffffffffffffffffffffffffff19166001600160801b0392909216919091179055613e3f82848161416d565b8160010160106101000a81548167ffffffffffffffff021916908367ffffffffffffffff160217905550505050565b6060815f01805480602002602001604051908101604052809291908181526020018280548015613ebb57602002820191905f5260205f20905b815481526020019060010190808311613ea7575b50505050509050919050565b5f8281527f02dd7bc7dec4dceedda775e58dd541e08a116c6c53815c0bd028192f7b626800602090815260408083206001600160a01b038516845290915290205460ff16611533576040517fe2517d3f0000000000000000000000000000000000000000000000000000000081526001600160a01b038216600482015260248101839052604401611c02565b5f611e5183836142f6565b5f8281527f02dd7bc7dec4dceedda775e58dd541e08a116c6c53815c0bd028192f7b626800602081815260408084206001600160a01b038616855290915282205460ff16614021575f848152602082815260408083206001600160a01b03871684529091529020805460ff19166001179055613fd73390565b6001600160a01b0316836001600160a01b0316857f2f8788117e7eff1d82e926ec794901d17c78024a50270940304540a733656f0d60405160405180910390a46001915050610fbe565b5f915050610fbe565b5f611e51836001600160a01b0384166142f6565b5f8281527f02dd7bc7dec4dceedda775e58dd541e08a116c6c53815c0bd028192f7b626800602081815260408084206001600160a01b038616855290915282205460ff1615614021575f848152602082815260408083206001600160a01b0387168085529252808320805460ff1916905551339287917ff6391f5c32d9c69d2a47ea670b442974b53935d1edc7fd64eb21e047a839171b9190a46001915050610fbe565b5f611e51836001600160a01b038416614342565b5f611e518383614342565b61410a8261441c565b6040516001600160a01b038316907fbc7cd75a20ee27fd9adebab32041f755214dbc6bffa90cc0225b39da2e5c2d3b905f90a280511561414e576112f1828261449f565b611533614511565b5f8181526001830160205260408120541515611e51565b5f67ffffffffffffffff8411156141a15760405163d450716560e01b81526004810184905260248101839052604401611c02565b509192915050565b5f6001600160801b038411156141a15760405163d450716560e01b81526004810184905260248101839052604401611c02565b5f825f0182815481106141f1576141f1614fc6565b905f5260205f200154905092915050565b5f8060205f8451602086015f885af180614221576040513d5f823e3d81fd5b50505f513d91508115614238578060011415614245565b6001600160a01b0384163b155b156113f0576040517f5274afe70000000000000000000000000000000000000000000000000000000081526001600160a01b0385166004820152602401611c02565b7ff0c57e16840df040f15088dc2f81fe391c3923bec73e23a9662efc9c229c6a005468010000000000000000900460ff16612e8a576040517fd7e6bcf800000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b6135da614287565b5f81815260018301602052604081205461433b57508154600181810184555f848152602080822090930184905584548482528286019093526040902091909155610fbe565b505f610fbe565b5f8181526001830160205260408120548015614021575f6143646001836152bc565b85549091505f90614377906001906152bc565b90508082146143d6575f865f01828154811061439557614395614fc6565b905f5260205f200154905080875f0184815481106143b5576143b5614fc6565b5f918252602080832090910192909255918252600188019052604090208390555b85548690806143e7576143e761539a565b600190038181905f5260205f20015f90559055856001015f8681526020019081526020015f205f905560019350505050610fbe565b806001600160a01b03163b5f0361445157604051634c9c8ce360e01b81526001600160a01b0382166004820152602401611c02565b7f360894a13ba1a3210667c828492db98dca3e2076cc3735a920a3ca505d382bbc805473ffffffffffffffffffffffffffffffffffffffff19166001600160a01b0392909216919091179055565b60605f80846001600160a01b0316846040516144bb9190614fda565b5f60405180830381855af49150503d805f81146144f3576040519150601f19603f3d011682016040523d82523d5f602084013e6144f8565b606091505b5091509150614508858383614549565b95945050505050565b3415612e8a576040517fb398979f00000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b60608261455e57614559826145be565b611e51565b815115801561457557506001600160a01b0384163b155b156145b7576040517f9996b3150000000000000000000000000000000000000000000000000000000081526001600160a01b0385166004820152602401611c02565b5080611e51565b8051156145ce5780518082602001fd5b60405163d6bda27560e01b815260040160405180910390fd5b828054828255905f5260205f2090810192821561462b579160200282015b8281111561462b578251829061461b908261518e565b5091602001919060010190614605565b50614637929150614689565b5090565b50805461464790614ff5565b5f825580601f10614656575050565b601f0160209004905f5260205f209081019061174e91906146a5565b5080545f8255905f5260205f209081019061174e91905b80821115614637575f61469c828261463b565b50600101614689565b5b80821115614637575f81556001016146a6565b5f602082840312156146c9575f80fd5b81356001600160e01b031981168114611e51575f80fd5b5f5b838110156146fa5781810151838201526020016146e2565b50505f910152565b5f81518084526147198160208601602086016146e0565b601f01601f19169290920160200192915050565b5f8282518085526020808601955060208260051b840101602086015f5b8481101561477857601f19868403018952614766838351614702565b9884019892509083019060010161474a565b5090979650505050505050565b602081525f611e51602083018461472d565b602080825282518282018190525f9190848201906040850190845b818110156147ce578351835292840192918401916001016147b2565b50909695505050505050565b6001600160a01b038116811461174e575f80fd5b634e487b7160e01b5f52604160045260245ffd5b604051601f8201601f1916810167ffffffffffffffff8111828210171561482b5761482b6147ee565b604052919050565b5f67ffffffffffffffff82111561484c5761484c6147ee565b50601f01601f191660200190565b5f82601f830112614869575f80fd5b813561487c61487782614833565b614802565b818152846020838601011115614890575f80fd5b816020850160208301375f918101602001919091529392505050565b5f805f80608085870312156148bf575f80fd5b84356148ca816147da565b935060208501356148da816147da565b925060408501359150606085013567ffffffffffffffff8111156148fc575f80fd5b6149088782880161485a565b91505092959194509250565b5f805f60608486031215614926575f80fd5b8335614931816147da565b92506020840135614941816147da565b929592945050506040919091013590565b5f8060408385031215614963575f80fd5b50508035926020909101359150565b5f82601f830112614981575f80fd5b8135602067ffffffffffffffff8083111561499e5761499e6147ee565b8260051b6149ad838201614802565b93845285810183019383810190888611156149c6575f80fd5b84880192505b85831015614a00578235848111156149e2575f80fd5b6149f08a87838c010161485a565b83525091840191908401906149cc565b98975050505050505050565b5f60208284031215614a1c575f80fd5b813567ffffffffffffffff811115614a32575f80fd5b611d8884828501614972565b5f60208284031215614a4e575f80fd5b5035919050565b803561ffff81168114614a66575f80fd5b919050565b5f805f60608486031215614a7d575f80fd5b83359250614a8d60208501614a55565b9150614a9b60408501614a55565b90509250925092565b5f8060408385031215614ab5575f80fd5b823591506020830135614ac7816147da565b809150509250929050565b5f60208284031215614ae2575f80fd5b813567ffffffffffffffff811115614af8575f80fd5b611d888482850161485a565b5f8060408385031215614b15575f80fd5b8235614b20816147da565b9150602083013567ffffffffffffffff811115614b3b575f80fd5b614b478582860161485a565b9150509250929050565b5f8060408385031215614b62575f80fd5b8235614b6d816147da565b946020939093013593505050565b80358015158114614a66575f80fd5b5f805f60608486031215614b9c575f80fd5b833567ffffffffffffffff80821115614bb3575f80fd5b614bbf8783880161485a565b9450614bcd60208701614b7b565b93506040860135915080821115614be2575f80fd5b50614bef86828701614972565b9150509250925092565b5f60208284031215614c09575f80fd5b8135611e51816147da565b5f805f8060808587031215614c27575f80fd5b8435935060208501359250614c3e60408601614a55565b9150614c4c60608601614a55565b905092959194509250565b5f8060408385031215614c68575f80fd5b823567ffffffffffffffff80821115614c7f575f80fd5b614c8b8683870161485a565b93506020850135915080821115614ca0575f80fd5b50614b478582860161485a565b5f604083018251151584526020808401516040602087015282815180855260608801915060608160051b89010194506020830192505f5b81811015614d1257605f19898703018352614d00868551614702565b95509284019291840191600101614ce4565b5093979650505050505050565b604081525f614d31604083018561472d565b6020838203818501528185518084528284019150828160051b8501018388015f5b83811015614d8057601f19878403018552614d6e838351614cad565b94860194925090850190600101614d52565b50909998505050505050505050565b602081525f611e516020830184614702565b5f60208284031215614db1575f80fd5b611e5182614b7b565b602080825282518282018190525f9190848201906040850190845b818110156147ce5783516001600160a01b031683529284019291840191600101614dd5565b5f8060408385031215614e0b575f80fd5b823567ffffffffffffffff811115614e21575f80fd5b614e2d8582860161485a565b925050614e3c60208401614b7b565b90509250929050565b5f805f60608486031215614e57575f80fd5b83359250602084013591506040840135614e70816147da565b809150509250925092565b602081525f611e516020830184614cad565b5f805f805f805f60e0888a031215614ea3575f80fd5b8735614eae816147da565b9650602088013567ffffffffffffffff811115614ec9575f80fd5b614ed58a828b0161485a565b96505060408801359450606088013593506080880135614ef4816147da565b925060a08801359150614f0960c08901614b7b565b905092959891949750929550565b5f8060408385031215614f28575f80fd5b823567ffffffffffffffff80821115614f3f575f80fd5b614f4b8683870161485a565b93506020850135915080821115614f60575f80fd5b50614b4785828601614972565b5f805f8060808587031215614f80575f80fd5b8435614f8b816147da565b93506020850135614f9b816147da565b92506040850135614fab816147da565b91506060850135614fbb816147da565b939692955090935050565b634e487b7160e01b5f52603260045260245ffd5b5f8251614feb8184602087016146e0565b9190910192915050565b600181811c9082168061500957607f821691505b60208210810361502757634e487b7160e01b5f52602260045260245ffd5b50919050565b5f6101006001600160a01b03808c168452808b1660208501528160408501526150588285018b614702565b6060850199909952608084019790975250509290931660a083015260c082015290151560e0909101529392505050565b5f60208284031215615098575f80fd5b815167ffffffffffffffff8111156150ae575f80fd5b8201601f810184136150be575f80fd5b80516150cc61487782614833565b8181528560208385010111156150e0575f80fd5b6145088260208301602086016146e0565b5f60208284031215615101575f80fd5b5051919050565b5f60208284031215615118575f80fd5b8151611e51816147da565b634e487b7160e01b5f52601160045260245ffd5b80820180821115610fbe57610fbe615123565b601f8211156112f157805f5260205f20601f840160051c8101602085101561516f5750805b601f840160051c820191505b81811015613bd7575f815560010161517b565b815167ffffffffffffffff8111156151a8576151a86147ee565b6151bc816151b68454614ff5565b8461514a565b602080601f8311600181146151ef575f84156151d85750858301515b5f19600386901b1c1916600185901b178555612517565b5f85815260208120601f198616915b8281101561521d578886015182559484019460019091019084016151fe565b508582101561523a57878501515f19600388901b60f8161c191681555b5050505050600190811b01905550565b5f80835461525781614ff5565b6001828116801561526f5760018114615284576152b0565b60ff19841687528215158302870194506152b0565b875f526020805f205f5b858110156152a75781548a82015290840190820161528e565b50505082870194505b50929695505050505050565b81810381811115610fbe57610fbe615123565b8181036152da575050565b6152e48254614ff5565b67ffffffffffffffff8111156152fc576152fc6147ee565b61530a816151b68454614ff5565b5f601f82116001811461533b575f83156153245750848201545b5f19600385901b1c1916600184901b178455613bd7565b5f8581526020808220868352908220601f198616925b838110156153715782860154825560019586019590910190602001615351565b508583101561523a579301545f1960f8600387901b161c19169092555050600190811b01905550565b634e487b7160e01b5f52603160045260245ffd5b828152604060208201525f611d88604083018461470256fe39fc6f3ebcc11656ef8cd451e9e1f6a26855304f0f1787c5b86c527a1e2d86009a95e87c5af084bf5db8491c3a6515da9dd6da39b24b0eb0af08d7b9cd808d913acdf00ba9ef08b5f2c22768276611b9af078bf6c24fa36b34ec5e9f2eb061fa17fbd8e22750b6dc617ceae09f600bf2810b53ccc1f790a9a4f78a3f7169b901a2646970667358221220ecefd9b53ba77aa476f887d78a5d272512f6f8a3f888919b54df88e401254d4b64736f6c63430008180033",
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

// GetServiceCapabilities is a free data retrieval call binding the contract method 0x319d13f3.
//
// Solidity: function getServiceCapabilities(string serviceName) view returns(string[] capabilities)
func (_Ttmaccount *TtmaccountCaller) GetServiceCapabilities(opts *bind.CallOpts, serviceName string) ([]string, error) {
	var out []interface{}
	err := _Ttmaccount.contract.Call(opts, &out, "getServiceCapabilities", serviceName)

	if err != nil {
		return *new([]string), err
	}

	out0 := *abi.ConvertType(out[0], new([]string)).(*[]string)

	return out0, err

}

// GetServiceCapabilities is a free data retrieval call binding the contract method 0x319d13f3.
//
// Solidity: function getServiceCapabilities(string serviceName) view returns(string[] capabilities)
func (_Ttmaccount *TtmaccountSession) GetServiceCapabilities(serviceName string) ([]string, error) {
	return _Ttmaccount.Contract.GetServiceCapabilities(&_Ttmaccount.CallOpts, serviceName)
}

// GetServiceCapabilities is a free data retrieval call binding the contract method 0x319d13f3.
//
// Solidity: function getServiceCapabilities(string serviceName) view returns(string[] capabilities)
func (_Ttmaccount *TtmaccountCallerSession) GetServiceCapabilities(serviceName string) ([]string, error) {
	return _Ttmaccount.Contract.GetServiceCapabilities(&_Ttmaccount.CallOpts, serviceName)
}

// GetServiceCapabilities0 is a free data retrieval call binding the contract method 0x5e07f869.
//
// Solidity: function getServiceCapabilities(bytes32 serviceHash) view returns(string[] capabilities)
func (_Ttmaccount *TtmaccountCaller) GetServiceCapabilities0(opts *bind.CallOpts, serviceHash [32]byte) ([]string, error) {
	var out []interface{}
	err := _Ttmaccount.contract.Call(opts, &out, "getServiceCapabilities0", serviceHash)

	if err != nil {
		return *new([]string), err
	}

	out0 := *abi.ConvertType(out[0], new([]string)).(*[]string)

	return out0, err

}

// GetServiceCapabilities0 is a free data retrieval call binding the contract method 0x5e07f869.
//
// Solidity: function getServiceCapabilities(bytes32 serviceHash) view returns(string[] capabilities)
func (_Ttmaccount *TtmaccountSession) GetServiceCapabilities0(serviceHash [32]byte) ([]string, error) {
	return _Ttmaccount.Contract.GetServiceCapabilities0(&_Ttmaccount.CallOpts, serviceHash)
}

// GetServiceCapabilities0 is a free data retrieval call binding the contract method 0x5e07f869.
//
// Solidity: function getServiceCapabilities(bytes32 serviceHash) view returns(string[] capabilities)
func (_Ttmaccount *TtmaccountCallerSession) GetServiceCapabilities0(serviceHash [32]byte) ([]string, error) {
	return _Ttmaccount.Contract.GetServiceCapabilities0(&_Ttmaccount.CallOpts, serviceHash)
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

// GetServiceRestrictedRate0 is a free data retrieval call binding the contract method 0xb5124635.
//
// Solidity: function getServiceRestrictedRate(string serviceName) view returns(bool restrictedRate)
func (_Ttmaccount *TtmaccountCaller) GetServiceRestrictedRate0(opts *bind.CallOpts, serviceName string) (bool, error) {
	var out []interface{}
	err := _Ttmaccount.contract.Call(opts, &out, "getServiceRestrictedRate0", serviceName)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// GetServiceRestrictedRate0 is a free data retrieval call binding the contract method 0xb5124635.
//
// Solidity: function getServiceRestrictedRate(string serviceName) view returns(bool restrictedRate)
func (_Ttmaccount *TtmaccountSession) GetServiceRestrictedRate0(serviceName string) (bool, error) {
	return _Ttmaccount.Contract.GetServiceRestrictedRate0(&_Ttmaccount.CallOpts, serviceName)
}

// GetServiceRestrictedRate0 is a free data retrieval call binding the contract method 0xb5124635.
//
// Solidity: function getServiceRestrictedRate(string serviceName) view returns(bool restrictedRate)
func (_Ttmaccount *TtmaccountCallerSession) GetServiceRestrictedRate0(serviceName string) (bool, error) {
	return _Ttmaccount.Contract.GetServiceRestrictedRate0(&_Ttmaccount.CallOpts, serviceName)
}

// GetSupportedServices is a free data retrieval call binding the contract method 0x7eec56c7.
//
// Solidity: function getSupportedServices() view returns(string[] serviceNames, (bool,string[])[] services)
func (_Ttmaccount *TtmaccountCaller) GetSupportedServices(opts *bind.CallOpts) (struct {
	ServiceNames []string
	Services     []PartnerConfigurationService
}, error) {
	var out []interface{}
	err := _Ttmaccount.contract.Call(opts, &out, "getSupportedServices")

	outstruct := new(struct {
		ServiceNames []string
		Services     []PartnerConfigurationService
	})
	if err != nil {
		return *outstruct, err
	}

	outstruct.ServiceNames = *abi.ConvertType(out[0], new([]string)).(*[]string)
	outstruct.Services = *abi.ConvertType(out[1], new([]PartnerConfigurationService)).(*[]PartnerConfigurationService)

	return *outstruct, err

}

// GetSupportedServices is a free data retrieval call binding the contract method 0x7eec56c7.
//
// Solidity: function getSupportedServices() view returns(string[] serviceNames, (bool,string[])[] services)
func (_Ttmaccount *TtmaccountSession) GetSupportedServices() (struct {
	ServiceNames []string
	Services     []PartnerConfigurationService
}, error) {
	return _Ttmaccount.Contract.GetSupportedServices(&_Ttmaccount.CallOpts)
}

// GetSupportedServices is a free data retrieval call binding the contract method 0x7eec56c7.
//
// Solidity: function getSupportedServices() view returns(string[] serviceNames, (bool,string[])[] services)
func (_Ttmaccount *TtmaccountCallerSession) GetSupportedServices() (struct {
	ServiceNames []string
	Services     []PartnerConfigurationService
}, error) {
	return _Ttmaccount.Contract.GetSupportedServices(&_Ttmaccount.CallOpts)
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

// GetWantedServices is a free data retrieval call binding the contract method 0x08564c19.
//
// Solidity: function getWantedServices() view returns(string[] serviceNames)
func (_Ttmaccount *TtmaccountCaller) GetWantedServices(opts *bind.CallOpts) ([]string, error) {
	var out []interface{}
	err := _Ttmaccount.contract.Call(opts, &out, "getWantedServices")

	if err != nil {
		return *new([]string), err
	}

	out0 := *abi.ConvertType(out[0], new([]string)).(*[]string)

	return out0, err

}

// GetWantedServices is a free data retrieval call binding the contract method 0x08564c19.
//
// Solidity: function getWantedServices() view returns(string[] serviceNames)
func (_Ttmaccount *TtmaccountSession) GetWantedServices() ([]string, error) {
	return _Ttmaccount.Contract.GetWantedServices(&_Ttmaccount.CallOpts)
}

// GetWantedServices is a free data retrieval call binding the contract method 0x08564c19.
//
// Solidity: function getWantedServices() view returns(string[] serviceNames)
func (_Ttmaccount *TtmaccountCallerSession) GetWantedServices() ([]string, error) {
	return _Ttmaccount.Contract.GetWantedServices(&_Ttmaccount.CallOpts)
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

// IsServiceSupported is a free data retrieval call binding the contract method 0x58c0a4c2.
//
// Solidity: function isServiceSupported(string serviceName) view returns(bool)
func (_Ttmaccount *TtmaccountCaller) IsServiceSupported(opts *bind.CallOpts, serviceName string) (bool, error) {
	var out []interface{}
	err := _Ttmaccount.contract.Call(opts, &out, "isServiceSupported", serviceName)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// IsServiceSupported is a free data retrieval call binding the contract method 0x58c0a4c2.
//
// Solidity: function isServiceSupported(string serviceName) view returns(bool)
func (_Ttmaccount *TtmaccountSession) IsServiceSupported(serviceName string) (bool, error) {
	return _Ttmaccount.Contract.IsServiceSupported(&_Ttmaccount.CallOpts, serviceName)
}

// IsServiceSupported is a free data retrieval call binding the contract method 0x58c0a4c2.
//
// Solidity: function isServiceSupported(string serviceName) view returns(bool)
func (_Ttmaccount *TtmaccountCallerSession) IsServiceSupported(serviceName string) (bool, error) {
	return _Ttmaccount.Contract.IsServiceSupported(&_Ttmaccount.CallOpts, serviceName)
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

// AddService is a paid mutator transaction binding the contract method 0x581ed290.
//
// Solidity: function addService(string serviceName, bool restrictedRate, string[] capabilities) returns()
func (_Ttmaccount *TtmaccountTransactor) AddService(opts *bind.TransactOpts, serviceName string, restrictedRate bool, capabilities []string) (*types.Transaction, error) {
	return _Ttmaccount.contract.Transact(opts, "addService", serviceName, restrictedRate, capabilities)
}

// AddService is a paid mutator transaction binding the contract method 0x581ed290.
//
// Solidity: function addService(string serviceName, bool restrictedRate, string[] capabilities) returns()
func (_Ttmaccount *TtmaccountSession) AddService(serviceName string, restrictedRate bool, capabilities []string) (*types.Transaction, error) {
	return _Ttmaccount.Contract.AddService(&_Ttmaccount.TransactOpts, serviceName, restrictedRate, capabilities)
}

// AddService is a paid mutator transaction binding the contract method 0x581ed290.
//
// Solidity: function addService(string serviceName, bool restrictedRate, string[] capabilities) returns()
func (_Ttmaccount *TtmaccountTransactorSession) AddService(serviceName string, restrictedRate bool, capabilities []string) (*types.Transaction, error) {
	return _Ttmaccount.Contract.AddService(&_Ttmaccount.TransactOpts, serviceName, restrictedRate, capabilities)
}

// AddServiceCapability is a paid mutator transaction binding the contract method 0x7512e55b.
//
// Solidity: function addServiceCapability(string serviceName, string capability) returns()
func (_Ttmaccount *TtmaccountTransactor) AddServiceCapability(opts *bind.TransactOpts, serviceName string, capability string) (*types.Transaction, error) {
	return _Ttmaccount.contract.Transact(opts, "addServiceCapability", serviceName, capability)
}

// AddServiceCapability is a paid mutator transaction binding the contract method 0x7512e55b.
//
// Solidity: function addServiceCapability(string serviceName, string capability) returns()
func (_Ttmaccount *TtmaccountSession) AddServiceCapability(serviceName string, capability string) (*types.Transaction, error) {
	return _Ttmaccount.Contract.AddServiceCapability(&_Ttmaccount.TransactOpts, serviceName, capability)
}

// AddServiceCapability is a paid mutator transaction binding the contract method 0x7512e55b.
//
// Solidity: function addServiceCapability(string serviceName, string capability) returns()
func (_Ttmaccount *TtmaccountTransactorSession) AddServiceCapability(serviceName string, capability string) (*types.Transaction, error) {
	return _Ttmaccount.Contract.AddServiceCapability(&_Ttmaccount.TransactOpts, serviceName, capability)
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

// AddWantedServices is a paid mutator transaction binding the contract method 0x1c5db99e.
//
// Solidity: function addWantedServices(string[] serviceNames) returns()
func (_Ttmaccount *TtmaccountTransactor) AddWantedServices(opts *bind.TransactOpts, serviceNames []string) (*types.Transaction, error) {
	return _Ttmaccount.contract.Transact(opts, "addWantedServices", serviceNames)
}

// AddWantedServices is a paid mutator transaction binding the contract method 0x1c5db99e.
//
// Solidity: function addWantedServices(string[] serviceNames) returns()
func (_Ttmaccount *TtmaccountSession) AddWantedServices(serviceNames []string) (*types.Transaction, error) {
	return _Ttmaccount.Contract.AddWantedServices(&_Ttmaccount.TransactOpts, serviceNames)
}

// AddWantedServices is a paid mutator transaction binding the contract method 0x1c5db99e.
//
// Solidity: function addWantedServices(string[] serviceNames) returns()
func (_Ttmaccount *TtmaccountTransactorSession) AddWantedServices(serviceNames []string) (*types.Transaction, error) {
	return _Ttmaccount.Contract.AddWantedServices(&_Ttmaccount.TransactOpts, serviceNames)
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

// RemoveService is a paid mutator transaction binding the contract method 0xf51acaea.
//
// Solidity: function removeService(string serviceName) returns()
func (_Ttmaccount *TtmaccountTransactor) RemoveService(opts *bind.TransactOpts, serviceName string) (*types.Transaction, error) {
	return _Ttmaccount.contract.Transact(opts, "removeService", serviceName)
}

// RemoveService is a paid mutator transaction binding the contract method 0xf51acaea.
//
// Solidity: function removeService(string serviceName) returns()
func (_Ttmaccount *TtmaccountSession) RemoveService(serviceName string) (*types.Transaction, error) {
	return _Ttmaccount.Contract.RemoveService(&_Ttmaccount.TransactOpts, serviceName)
}

// RemoveService is a paid mutator transaction binding the contract method 0xf51acaea.
//
// Solidity: function removeService(string serviceName) returns()
func (_Ttmaccount *TtmaccountTransactorSession) RemoveService(serviceName string) (*types.Transaction, error) {
	return _Ttmaccount.Contract.RemoveService(&_Ttmaccount.TransactOpts, serviceName)
}

// RemoveServiceCapability is a paid mutator transaction binding the contract method 0x8c20f574.
//
// Solidity: function removeServiceCapability(string serviceName, string capability) returns()
func (_Ttmaccount *TtmaccountTransactor) RemoveServiceCapability(opts *bind.TransactOpts, serviceName string, capability string) (*types.Transaction, error) {
	return _Ttmaccount.contract.Transact(opts, "removeServiceCapability", serviceName, capability)
}

// RemoveServiceCapability is a paid mutator transaction binding the contract method 0x8c20f574.
//
// Solidity: function removeServiceCapability(string serviceName, string capability) returns()
func (_Ttmaccount *TtmaccountSession) RemoveServiceCapability(serviceName string, capability string) (*types.Transaction, error) {
	return _Ttmaccount.Contract.RemoveServiceCapability(&_Ttmaccount.TransactOpts, serviceName, capability)
}

// RemoveServiceCapability is a paid mutator transaction binding the contract method 0x8c20f574.
//
// Solidity: function removeServiceCapability(string serviceName, string capability) returns()
func (_Ttmaccount *TtmaccountTransactorSession) RemoveServiceCapability(serviceName string, capability string) (*types.Transaction, error) {
	return _Ttmaccount.Contract.RemoveServiceCapability(&_Ttmaccount.TransactOpts, serviceName, capability)
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

// RemoveWantedServices is a paid mutator transaction binding the contract method 0x39e4c705.
//
// Solidity: function removeWantedServices(string[] serviceNames) returns()
func (_Ttmaccount *TtmaccountTransactor) RemoveWantedServices(opts *bind.TransactOpts, serviceNames []string) (*types.Transaction, error) {
	return _Ttmaccount.contract.Transact(opts, "removeWantedServices", serviceNames)
}

// RemoveWantedServices is a paid mutator transaction binding the contract method 0x39e4c705.
//
// Solidity: function removeWantedServices(string[] serviceNames) returns()
func (_Ttmaccount *TtmaccountSession) RemoveWantedServices(serviceNames []string) (*types.Transaction, error) {
	return _Ttmaccount.Contract.RemoveWantedServices(&_Ttmaccount.TransactOpts, serviceNames)
}

// RemoveWantedServices is a paid mutator transaction binding the contract method 0x39e4c705.
//
// Solidity: function removeWantedServices(string[] serviceNames) returns()
func (_Ttmaccount *TtmaccountTransactorSession) RemoveWantedServices(serviceNames []string) (*types.Transaction, error) {
	return _Ttmaccount.Contract.RemoveWantedServices(&_Ttmaccount.TransactOpts, serviceNames)
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

// SetServiceCapabilities is a paid mutator transaction binding the contract method 0xebc20d20.
//
// Solidity: function setServiceCapabilities(string serviceName, string[] capabilities) returns()
func (_Ttmaccount *TtmaccountTransactor) SetServiceCapabilities(opts *bind.TransactOpts, serviceName string, capabilities []string) (*types.Transaction, error) {
	return _Ttmaccount.contract.Transact(opts, "setServiceCapabilities", serviceName, capabilities)
}

// SetServiceCapabilities is a paid mutator transaction binding the contract method 0xebc20d20.
//
// Solidity: function setServiceCapabilities(string serviceName, string[] capabilities) returns()
func (_Ttmaccount *TtmaccountSession) SetServiceCapabilities(serviceName string, capabilities []string) (*types.Transaction, error) {
	return _Ttmaccount.Contract.SetServiceCapabilities(&_Ttmaccount.TransactOpts, serviceName, capabilities)
}

// SetServiceCapabilities is a paid mutator transaction binding the contract method 0xebc20d20.
//
// Solidity: function setServiceCapabilities(string serviceName, string[] capabilities) returns()
func (_Ttmaccount *TtmaccountTransactorSession) SetServiceCapabilities(serviceName string, capabilities []string) (*types.Transaction, error) {
	return _Ttmaccount.Contract.SetServiceCapabilities(&_Ttmaccount.TransactOpts, serviceName, capabilities)
}

// SetServiceRestrictedRate is a paid mutator transaction binding the contract method 0xa7d022f8.
//
// Solidity: function setServiceRestrictedRate(string serviceName, bool restrictedRate) returns()
func (_Ttmaccount *TtmaccountTransactor) SetServiceRestrictedRate(opts *bind.TransactOpts, serviceName string, restrictedRate bool) (*types.Transaction, error) {
	return _Ttmaccount.contract.Transact(opts, "setServiceRestrictedRate", serviceName, restrictedRate)
}

// SetServiceRestrictedRate is a paid mutator transaction binding the contract method 0xa7d022f8.
//
// Solidity: function setServiceRestrictedRate(string serviceName, bool restrictedRate) returns()
func (_Ttmaccount *TtmaccountSession) SetServiceRestrictedRate(serviceName string, restrictedRate bool) (*types.Transaction, error) {
	return _Ttmaccount.Contract.SetServiceRestrictedRate(&_Ttmaccount.TransactOpts, serviceName, restrictedRate)
}

// SetServiceRestrictedRate is a paid mutator transaction binding the contract method 0xa7d022f8.
//
// Solidity: function setServiceRestrictedRate(string serviceName, bool restrictedRate) returns()
func (_Ttmaccount *TtmaccountTransactorSession) SetServiceRestrictedRate(serviceName string, restrictedRate bool) (*types.Transaction, error) {
	return _Ttmaccount.Contract.SetServiceRestrictedRate(&_Ttmaccount.TransactOpts, serviceName, restrictedRate)
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
	ServiceName common.Hash
	Raw         types.Log // Blockchain specific contextual infos
}

// FilterServiceAdded is a free log retrieval operation binding the contract event 0x763f2f41e0c407dd0a7067f44e5468a0db74da9fdb6cd1cb20c7b6dae9375279.
//
// Solidity: event ServiceAdded(string indexed serviceName)
func (_Ttmaccount *TtmaccountFilterer) FilterServiceAdded(opts *bind.FilterOpts, serviceName []string) (*TtmaccountServiceAddedIterator, error) {

	var serviceNameRule []interface{}
	for _, serviceNameItem := range serviceName {
		serviceNameRule = append(serviceNameRule, serviceNameItem)
	}

	logs, sub, err := _Ttmaccount.contract.FilterLogs(opts, "ServiceAdded", serviceNameRule)
	if err != nil {
		return nil, err
	}
	return &TtmaccountServiceAddedIterator{contract: _Ttmaccount.contract, event: "ServiceAdded", logs: logs, sub: sub}, nil
}

// WatchServiceAdded is a free log subscription operation binding the contract event 0x763f2f41e0c407dd0a7067f44e5468a0db74da9fdb6cd1cb20c7b6dae9375279.
//
// Solidity: event ServiceAdded(string indexed serviceName)
func (_Ttmaccount *TtmaccountFilterer) WatchServiceAdded(opts *bind.WatchOpts, sink chan<- *TtmaccountServiceAdded, serviceName []string) (event.Subscription, error) {

	var serviceNameRule []interface{}
	for _, serviceNameItem := range serviceName {
		serviceNameRule = append(serviceNameRule, serviceNameItem)
	}

	logs, sub, err := _Ttmaccount.contract.WatchLogs(opts, "ServiceAdded", serviceNameRule)
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

// ParseServiceAdded is a log parse operation binding the contract event 0x763f2f41e0c407dd0a7067f44e5468a0db74da9fdb6cd1cb20c7b6dae9375279.
//
// Solidity: event ServiceAdded(string indexed serviceName)
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
	ServiceName common.Hash
	Raw         types.Log // Blockchain specific contextual infos
}

// FilterServiceCapabilitiesUpdated is a free log retrieval operation binding the contract event 0xd52aef6010d6b6303240865274298b7c5784b14ebf9df788047b34c69c531371.
//
// Solidity: event ServiceCapabilitiesUpdated(string indexed serviceName)
func (_Ttmaccount *TtmaccountFilterer) FilterServiceCapabilitiesUpdated(opts *bind.FilterOpts, serviceName []string) (*TtmaccountServiceCapabilitiesUpdatedIterator, error) {

	var serviceNameRule []interface{}
	for _, serviceNameItem := range serviceName {
		serviceNameRule = append(serviceNameRule, serviceNameItem)
	}

	logs, sub, err := _Ttmaccount.contract.FilterLogs(opts, "ServiceCapabilitiesUpdated", serviceNameRule)
	if err != nil {
		return nil, err
	}
	return &TtmaccountServiceCapabilitiesUpdatedIterator{contract: _Ttmaccount.contract, event: "ServiceCapabilitiesUpdated", logs: logs, sub: sub}, nil
}

// WatchServiceCapabilitiesUpdated is a free log subscription operation binding the contract event 0xd52aef6010d6b6303240865274298b7c5784b14ebf9df788047b34c69c531371.
//
// Solidity: event ServiceCapabilitiesUpdated(string indexed serviceName)
func (_Ttmaccount *TtmaccountFilterer) WatchServiceCapabilitiesUpdated(opts *bind.WatchOpts, sink chan<- *TtmaccountServiceCapabilitiesUpdated, serviceName []string) (event.Subscription, error) {

	var serviceNameRule []interface{}
	for _, serviceNameItem := range serviceName {
		serviceNameRule = append(serviceNameRule, serviceNameItem)
	}

	logs, sub, err := _Ttmaccount.contract.WatchLogs(opts, "ServiceCapabilitiesUpdated", serviceNameRule)
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

// ParseServiceCapabilitiesUpdated is a log parse operation binding the contract event 0xd52aef6010d6b6303240865274298b7c5784b14ebf9df788047b34c69c531371.
//
// Solidity: event ServiceCapabilitiesUpdated(string indexed serviceName)
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
	ServiceName common.Hash
	Capability  string
	Raw         types.Log // Blockchain specific contextual infos
}

// FilterServiceCapabilityAdded is a free log retrieval operation binding the contract event 0x498a5f4e6f3921f63e6863032989bdb7bb41e5cf5cbde5437c7322c5c8dc46bf.
//
// Solidity: event ServiceCapabilityAdded(string indexed serviceName, string capability)
func (_Ttmaccount *TtmaccountFilterer) FilterServiceCapabilityAdded(opts *bind.FilterOpts, serviceName []string) (*TtmaccountServiceCapabilityAddedIterator, error) {

	var serviceNameRule []interface{}
	for _, serviceNameItem := range serviceName {
		serviceNameRule = append(serviceNameRule, serviceNameItem)
	}

	logs, sub, err := _Ttmaccount.contract.FilterLogs(opts, "ServiceCapabilityAdded", serviceNameRule)
	if err != nil {
		return nil, err
	}
	return &TtmaccountServiceCapabilityAddedIterator{contract: _Ttmaccount.contract, event: "ServiceCapabilityAdded", logs: logs, sub: sub}, nil
}

// WatchServiceCapabilityAdded is a free log subscription operation binding the contract event 0x498a5f4e6f3921f63e6863032989bdb7bb41e5cf5cbde5437c7322c5c8dc46bf.
//
// Solidity: event ServiceCapabilityAdded(string indexed serviceName, string capability)
func (_Ttmaccount *TtmaccountFilterer) WatchServiceCapabilityAdded(opts *bind.WatchOpts, sink chan<- *TtmaccountServiceCapabilityAdded, serviceName []string) (event.Subscription, error) {

	var serviceNameRule []interface{}
	for _, serviceNameItem := range serviceName {
		serviceNameRule = append(serviceNameRule, serviceNameItem)
	}

	logs, sub, err := _Ttmaccount.contract.WatchLogs(opts, "ServiceCapabilityAdded", serviceNameRule)
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

// ParseServiceCapabilityAdded is a log parse operation binding the contract event 0x498a5f4e6f3921f63e6863032989bdb7bb41e5cf5cbde5437c7322c5c8dc46bf.
//
// Solidity: event ServiceCapabilityAdded(string indexed serviceName, string capability)
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
	ServiceName common.Hash
	Capability  string
	Raw         types.Log // Blockchain specific contextual infos
}

// FilterServiceCapabilityRemoved is a free log retrieval operation binding the contract event 0xba851faec9e30a9961f0adb49fe025cda6c8d7d0fb9bad99f89c37d057023264.
//
// Solidity: event ServiceCapabilityRemoved(string indexed serviceName, string capability)
func (_Ttmaccount *TtmaccountFilterer) FilterServiceCapabilityRemoved(opts *bind.FilterOpts, serviceName []string) (*TtmaccountServiceCapabilityRemovedIterator, error) {

	var serviceNameRule []interface{}
	for _, serviceNameItem := range serviceName {
		serviceNameRule = append(serviceNameRule, serviceNameItem)
	}

	logs, sub, err := _Ttmaccount.contract.FilterLogs(opts, "ServiceCapabilityRemoved", serviceNameRule)
	if err != nil {
		return nil, err
	}
	return &TtmaccountServiceCapabilityRemovedIterator{contract: _Ttmaccount.contract, event: "ServiceCapabilityRemoved", logs: logs, sub: sub}, nil
}

// WatchServiceCapabilityRemoved is a free log subscription operation binding the contract event 0xba851faec9e30a9961f0adb49fe025cda6c8d7d0fb9bad99f89c37d057023264.
//
// Solidity: event ServiceCapabilityRemoved(string indexed serviceName, string capability)
func (_Ttmaccount *TtmaccountFilterer) WatchServiceCapabilityRemoved(opts *bind.WatchOpts, sink chan<- *TtmaccountServiceCapabilityRemoved, serviceName []string) (event.Subscription, error) {

	var serviceNameRule []interface{}
	for _, serviceNameItem := range serviceName {
		serviceNameRule = append(serviceNameRule, serviceNameItem)
	}

	logs, sub, err := _Ttmaccount.contract.WatchLogs(opts, "ServiceCapabilityRemoved", serviceNameRule)
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

// ParseServiceCapabilityRemoved is a log parse operation binding the contract event 0xba851faec9e30a9961f0adb49fe025cda6c8d7d0fb9bad99f89c37d057023264.
//
// Solidity: event ServiceCapabilityRemoved(string indexed serviceName, string capability)
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
	ServiceName common.Hash
	Raw         types.Log // Blockchain specific contextual infos
}

// FilterServiceRemoved is a free log retrieval operation binding the contract event 0x52f6e0779195109314dfb8cf301d33491c63f136afac4c5d4f35aa934b254813.
//
// Solidity: event ServiceRemoved(string indexed serviceName)
func (_Ttmaccount *TtmaccountFilterer) FilterServiceRemoved(opts *bind.FilterOpts, serviceName []string) (*TtmaccountServiceRemovedIterator, error) {

	var serviceNameRule []interface{}
	for _, serviceNameItem := range serviceName {
		serviceNameRule = append(serviceNameRule, serviceNameItem)
	}

	logs, sub, err := _Ttmaccount.contract.FilterLogs(opts, "ServiceRemoved", serviceNameRule)
	if err != nil {
		return nil, err
	}
	return &TtmaccountServiceRemovedIterator{contract: _Ttmaccount.contract, event: "ServiceRemoved", logs: logs, sub: sub}, nil
}

// WatchServiceRemoved is a free log subscription operation binding the contract event 0x52f6e0779195109314dfb8cf301d33491c63f136afac4c5d4f35aa934b254813.
//
// Solidity: event ServiceRemoved(string indexed serviceName)
func (_Ttmaccount *TtmaccountFilterer) WatchServiceRemoved(opts *bind.WatchOpts, sink chan<- *TtmaccountServiceRemoved, serviceName []string) (event.Subscription, error) {

	var serviceNameRule []interface{}
	for _, serviceNameItem := range serviceName {
		serviceNameRule = append(serviceNameRule, serviceNameItem)
	}

	logs, sub, err := _Ttmaccount.contract.WatchLogs(opts, "ServiceRemoved", serviceNameRule)
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

// ParseServiceRemoved is a log parse operation binding the contract event 0x52f6e0779195109314dfb8cf301d33491c63f136afac4c5d4f35aa934b254813.
//
// Solidity: event ServiceRemoved(string indexed serviceName)
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
	ServiceName    common.Hash
	RestrictedRate bool
	Raw            types.Log // Blockchain specific contextual infos
}

// FilterServiceRestrictedRateUpdated is a free log retrieval operation binding the contract event 0x23960b931eb4b63e2e47d040f51cc0de6eef2e865639eb674fbb1890ece3a0ab.
//
// Solidity: event ServiceRestrictedRateUpdated(string indexed serviceName, bool restrictedRate)
func (_Ttmaccount *TtmaccountFilterer) FilterServiceRestrictedRateUpdated(opts *bind.FilterOpts, serviceName []string) (*TtmaccountServiceRestrictedRateUpdatedIterator, error) {

	var serviceNameRule []interface{}
	for _, serviceNameItem := range serviceName {
		serviceNameRule = append(serviceNameRule, serviceNameItem)
	}

	logs, sub, err := _Ttmaccount.contract.FilterLogs(opts, "ServiceRestrictedRateUpdated", serviceNameRule)
	if err != nil {
		return nil, err
	}
	return &TtmaccountServiceRestrictedRateUpdatedIterator{contract: _Ttmaccount.contract, event: "ServiceRestrictedRateUpdated", logs: logs, sub: sub}, nil
}

// WatchServiceRestrictedRateUpdated is a free log subscription operation binding the contract event 0x23960b931eb4b63e2e47d040f51cc0de6eef2e865639eb674fbb1890ece3a0ab.
//
// Solidity: event ServiceRestrictedRateUpdated(string indexed serviceName, bool restrictedRate)
func (_Ttmaccount *TtmaccountFilterer) WatchServiceRestrictedRateUpdated(opts *bind.WatchOpts, sink chan<- *TtmaccountServiceRestrictedRateUpdated, serviceName []string) (event.Subscription, error) {

	var serviceNameRule []interface{}
	for _, serviceNameItem := range serviceName {
		serviceNameRule = append(serviceNameRule, serviceNameItem)
	}

	logs, sub, err := _Ttmaccount.contract.WatchLogs(opts, "ServiceRestrictedRateUpdated", serviceNameRule)
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

// ParseServiceRestrictedRateUpdated is a log parse operation binding the contract event 0x23960b931eb4b63e2e47d040f51cc0de6eef2e865639eb674fbb1890ece3a0ab.
//
// Solidity: event ServiceRestrictedRateUpdated(string indexed serviceName, bool restrictedRate)
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
	ServiceName common.Hash
	Raw         types.Log // Blockchain specific contextual infos
}

// FilterWantedServiceAdded is a free log retrieval operation binding the contract event 0x50cc5f9d56177aa0de269c136f2d2ffd45d7b66c82f0a82f8f840db54d9801f8.
//
// Solidity: event WantedServiceAdded(string indexed serviceName)
func (_Ttmaccount *TtmaccountFilterer) FilterWantedServiceAdded(opts *bind.FilterOpts, serviceName []string) (*TtmaccountWantedServiceAddedIterator, error) {

	var serviceNameRule []interface{}
	for _, serviceNameItem := range serviceName {
		serviceNameRule = append(serviceNameRule, serviceNameItem)
	}

	logs, sub, err := _Ttmaccount.contract.FilterLogs(opts, "WantedServiceAdded", serviceNameRule)
	if err != nil {
		return nil, err
	}
	return &TtmaccountWantedServiceAddedIterator{contract: _Ttmaccount.contract, event: "WantedServiceAdded", logs: logs, sub: sub}, nil
}

// WatchWantedServiceAdded is a free log subscription operation binding the contract event 0x50cc5f9d56177aa0de269c136f2d2ffd45d7b66c82f0a82f8f840db54d9801f8.
//
// Solidity: event WantedServiceAdded(string indexed serviceName)
func (_Ttmaccount *TtmaccountFilterer) WatchWantedServiceAdded(opts *bind.WatchOpts, sink chan<- *TtmaccountWantedServiceAdded, serviceName []string) (event.Subscription, error) {

	var serviceNameRule []interface{}
	for _, serviceNameItem := range serviceName {
		serviceNameRule = append(serviceNameRule, serviceNameItem)
	}

	logs, sub, err := _Ttmaccount.contract.WatchLogs(opts, "WantedServiceAdded", serviceNameRule)
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

// ParseWantedServiceAdded is a log parse operation binding the contract event 0x50cc5f9d56177aa0de269c136f2d2ffd45d7b66c82f0a82f8f840db54d9801f8.
//
// Solidity: event WantedServiceAdded(string indexed serviceName)
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
	ServiceName common.Hash
	Raw         types.Log // Blockchain specific contextual infos
}

// FilterWantedServiceRemoved is a free log retrieval operation binding the contract event 0x0edb7a081e6ff720ad9e97b837c352ef0548c4d09ec421b9b930b1e0c708e39e.
//
// Solidity: event WantedServiceRemoved(string indexed serviceName)
func (_Ttmaccount *TtmaccountFilterer) FilterWantedServiceRemoved(opts *bind.FilterOpts, serviceName []string) (*TtmaccountWantedServiceRemovedIterator, error) {

	var serviceNameRule []interface{}
	for _, serviceNameItem := range serviceName {
		serviceNameRule = append(serviceNameRule, serviceNameItem)
	}

	logs, sub, err := _Ttmaccount.contract.FilterLogs(opts, "WantedServiceRemoved", serviceNameRule)
	if err != nil {
		return nil, err
	}
	return &TtmaccountWantedServiceRemovedIterator{contract: _Ttmaccount.contract, event: "WantedServiceRemoved", logs: logs, sub: sub}, nil
}

// WatchWantedServiceRemoved is a free log subscription operation binding the contract event 0x0edb7a081e6ff720ad9e97b837c352ef0548c4d09ec421b9b930b1e0c708e39e.
//
// Solidity: event WantedServiceRemoved(string indexed serviceName)
func (_Ttmaccount *TtmaccountFilterer) WatchWantedServiceRemoved(opts *bind.WatchOpts, sink chan<- *TtmaccountWantedServiceRemoved, serviceName []string) (event.Subscription, error) {

	var serviceNameRule []interface{}
	for _, serviceNameItem := range serviceName {
		serviceNameRule = append(serviceNameRule, serviceNameItem)
	}

	logs, sub, err := _Ttmaccount.contract.WatchLogs(opts, "WantedServiceRemoved", serviceNameRule)
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

// ParseWantedServiceRemoved is a log parse operation binding the contract event 0x0edb7a081e6ff720ad9e97b837c352ef0548c4d09ec421b9b930b1e0c708e39e.
//
// Solidity: event WantedServiceRemoved(string indexed serviceName)
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
