// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package bookingtoken

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

// BookingtokenMetaData contains all meta data concerning the Bookingtoken contract.
var BookingtokenMetaData = &bind.MetaData{
	ABI: "[{\"inputs\":[],\"name\":\"AccessControlBadConfirmation\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"},{\"internalType\":\"bytes32\",\"name\":\"neededRole\",\"type\":\"bytes32\"}],\"name\":\"AccessControlUnauthorizedAccount\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"target\",\"type\":\"address\"}],\"name\":\"AddressEmptyCode\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"implementation\",\"type\":\"address\"}],\"name\":\"ERC1967InvalidImplementation\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"ERC1967NonPayable\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"ERC721EnumerableForbiddenBatchMint\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"sender\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"tokenId\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"owner\",\"type\":\"address\"}],\"name\":\"ERC721IncorrectOwner\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"operator\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"tokenId\",\"type\":\"uint256\"}],\"name\":\"ERC721InsufficientApproval\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"approver\",\"type\":\"address\"}],\"name\":\"ERC721InvalidApprover\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"operator\",\"type\":\"address\"}],\"name\":\"ERC721InvalidOperator\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"owner\",\"type\":\"address\"}],\"name\":\"ERC721InvalidOwner\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"receiver\",\"type\":\"address\"}],\"name\":\"ERC721InvalidReceiver\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"sender\",\"type\":\"address\"}],\"name\":\"ERC721InvalidSender\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"tokenId\",\"type\":\"uint256\"}],\"name\":\"ERC721NonexistentToken\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"owner\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"index\",\"type\":\"uint256\"}],\"name\":\"ERC721OutOfBoundsIndex\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"EnforcedPause\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"ExpectedPause\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"expirationTimestamp\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"minExpirationTimestampDiff\",\"type\":\"uint256\"}],\"name\":\"ExpirationTimestampTooSoon\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"FailedCall\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"price\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"reservationPrice\",\"type\":\"uint256\"}],\"name\":\"IncorrectPrice\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"tokenId\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"existing\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"checked\",\"type\":\"uint256\"}],\"name\":\"IncorrectRefundAmount\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"balance\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"needed\",\"type\":\"uint256\"}],\"name\":\"InsufficientBalance\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"tokenId\",\"type\":\"uint256\"},{\"internalType\":\"enumCancellationProposalStatus\",\"name\":\"status\",\"type\":\"uint8\"}],\"name\":\"InvalidCancellationProposalStatus\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"InvalidInitialization\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"tokenId\",\"type\":\"uint256\"},{\"internalType\":\"enumBookingToken.BookingStatus\",\"name\":\"status\",\"type\":\"uint8\"}],\"name\":\"InvalidTokenStatus\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"NotInitializing\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"NotOwnerOrSupplier\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"}],\"name\":\"NotTTMAccount\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"tokenId\",\"type\":\"uint256\"}],\"name\":\"OnlyCurrentProposerCanWithdrawCancellation\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"tokenId\",\"type\":\"uint256\"}],\"name\":\"OnlySupplierCanFinalizeCancellation\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"tokenId\",\"type\":\"uint256\"}],\"name\":\"OwnerNotAcceptedCancellation\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"tokenId\",\"type\":\"uint256\"}],\"name\":\"ProposerCanNotRejectCancellation\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"ReentrancyGuardReentrantCall\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"tokenId\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"expirationTimestamp\",\"type\":\"uint256\"}],\"name\":\"ReservationExpired\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"reservedFor\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"buyer\",\"type\":\"address\"}],\"name\":\"ReservationMismatch\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"token\",\"type\":\"address\"}],\"name\":\"SafeERC20FailedOperation\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"tokenId\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"supplier\",\"type\":\"address\"}],\"name\":\"SupplierIsNotOwner\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"tokenId\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"reservedFor\",\"type\":\"address\"}],\"name\":\"TokenIsReserved\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"UUPSUnauthorizedCallContext\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"slot\",\"type\":\"bytes32\"}],\"name\":\"UUPSUnsupportedProxiableUUID\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"UnexpectedNativePayment\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"offchainPaymentCurrency\",\"type\":\"uint256\"}],\"name\":\"UnexpectedOffchainPaymentCurrency\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"ZeroAddress\",\"type\":\"error\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"owner\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"approved\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"uint256\",\"name\":\"tokenId\",\"type\":\"uint256\"}],\"name\":\"Approval\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"owner\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"operator\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"bool\",\"name\":\"approved\",\"type\":\"bool\"}],\"name\":\"ApprovalForAll\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"_fromTokenId\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"_toTokenId\",\"type\":\"uint256\"}],\"name\":\"BatchMetadataUpdate\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"uint256\",\"name\":\"tokenId\",\"type\":\"uint256\"}],\"name\":\"CancellationFinalized\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"uint256\",\"name\":\"tokenId\",\"type\":\"uint256\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"initialProposer\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"currentProposer\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"refundAmount\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"bool\",\"name\":\"ownerAccepted\",\"type\":\"bool\"},{\"indexed\":false,\"internalType\":\"bool\",\"name\":\"supplierAccepted\",\"type\":\"bool\"},{\"indexed\":false,\"internalType\":\"uint32\",\"name\":\"timesCountered\",\"type\":\"uint32\"},{\"indexed\":false,\"internalType\":\"uint32\",\"name\":\"timesRejected\",\"type\":\"uint32\"}],\"name\":\"CancellationPending\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"uint256\",\"name\":\"tokenId\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint16\",\"name\":\"cancellationReason\",\"type\":\"uint16\"},{\"indexed\":false,\"internalType\":\"uint16\",\"name\":\"cancellationReasonVersion\",\"type\":\"uint16\"},{\"indexed\":false,\"internalType\":\"uint16\",\"name\":\"rejectionReason\",\"type\":\"uint16\"},{\"indexed\":false,\"internalType\":\"uint16\",\"name\":\"rejectionVersion\",\"type\":\"uint16\"},{\"indexed\":false,\"internalType\":\"uint16\",\"name\":\"counterReason\",\"type\":\"uint16\"},{\"indexed\":false,\"internalType\":\"uint16\",\"name\":\"counterVersion\",\"type\":\"uint16\"},{\"indexed\":false,\"internalType\":\"uint16\",\"name\":\"withdrawalReason\",\"type\":\"uint16\"},{\"indexed\":false,\"internalType\":\"uint16\",\"name\":\"withdrawalVersion\",\"type\":\"uint16\"}],\"name\":\"CancellationReasons\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"uint256\",\"name\":\"tokenId\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint16\",\"name\":\"rejectionReason\",\"type\":\"uint16\"},{\"indexed\":false,\"internalType\":\"uint16\",\"name\":\"rejectionVersion\",\"type\":\"uint16\"}],\"name\":\"CancellationRejected\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"uint256\",\"name\":\"tokenId\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint16\",\"name\":\"withdrawalReason\",\"type\":\"uint16\"},{\"indexed\":false,\"internalType\":\"uint16\",\"name\":\"withdrawalVersion\",\"type\":\"uint16\"}],\"name\":\"CancellationWithdrawn\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint64\",\"name\":\"version\",\"type\":\"uint64\"}],\"name\":\"Initialized\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"_tokenId\",\"type\":\"uint256\"}],\"name\":\"MetadataUpdate\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"}],\"name\":\"Paused\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"bytes32\",\"name\":\"role\",\"type\":\"bytes32\"},{\"indexed\":true,\"internalType\":\"bytes32\",\"name\":\"previousAdminRole\",\"type\":\"bytes32\"},{\"indexed\":true,\"internalType\":\"bytes32\",\"name\":\"newAdminRole\",\"type\":\"bytes32\"}],\"name\":\"RoleAdminChanged\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"bytes32\",\"name\":\"role\",\"type\":\"bytes32\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"sender\",\"type\":\"address\"}],\"name\":\"RoleGranted\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"bytes32\",\"name\":\"role\",\"type\":\"bytes32\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"sender\",\"type\":\"address\"}],\"name\":\"RoleRevoked\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"uint256\",\"name\":\"tokenId\",\"type\":\"uint256\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"buyer\",\"type\":\"address\"}],\"name\":\"TokenBought\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"uint256\",\"name\":\"tokenId\",\"type\":\"uint256\"}],\"name\":\"TokenReservationExpired\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"uint256\",\"name\":\"tokenId\",\"type\":\"uint256\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"reservedFor\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"supplier\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"expirationTimestamp\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"price\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"contractIERC20\",\"name\":\"paymentToken\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"offchainPaymentCurrency\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"bool\",\"name\":\"cancellable\",\"type\":\"bool\"}],\"name\":\"TokenReserved\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"from\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"uint256\",\"name\":\"tokenId\",\"type\":\"uint256\"}],\"name\":\"Transfer\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"}],\"name\":\"Unpaused\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"implementation\",\"type\":\"address\"}],\"name\":\"Upgraded\",\"type\":\"event\"},{\"inputs\":[],\"name\":\"DEFAULT_ADMIN_ROLE\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"MIN_EXPIRATION_ADMIN_ROLE\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"NATIVE_PAYMENT\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"OFFCHAIN_PAYMENT\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"PAUSER_ROLE\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"UPGRADER_ROLE\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"UPGRADE_INTERFACE_VERSION\",\"outputs\":[{\"internalType\":\"string\",\"name\":\"\",\"type\":\"string\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"tokenId\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"refundAmount\",\"type\":\"uint256\"}],\"name\":\"acceptCancellation\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"tokenId\",\"type\":\"uint256\"}],\"name\":\"approve\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"owner\",\"type\":\"address\"}],\"name\":\"balanceOf\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"tokenId\",\"type\":\"uint256\"}],\"name\":\"buyReservedToken\",\"outputs\":[],\"stateMutability\":\"payable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"tokenId\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"refundAmount\",\"type\":\"uint256\"},{\"internalType\":\"uint16\",\"name\":\"counterReason\",\"type\":\"uint16\"},{\"internalType\":\"uint16\",\"name\":\"counterReasonVersion\",\"type\":\"uint16\"}],\"name\":\"counterCancellation\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"tokenId\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"checkRefundAmount\",\"type\":\"uint256\"}],\"name\":\"finalizeCancellation\",\"outputs\":[],\"stateMutability\":\"payable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"tokenId\",\"type\":\"uint256\"}],\"name\":\"getApproved\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"tokenId\",\"type\":\"uint256\"}],\"name\":\"getBookingStatus\",\"outputs\":[{\"internalType\":\"enumBookingToken.BookingStatus\",\"name\":\"\",\"type\":\"uint8\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"tokenId\",\"type\":\"uint256\"}],\"name\":\"getCancellationProposal\",\"outputs\":[{\"internalType\":\"enumCancellationProposalStatus\",\"name\":\"\",\"type\":\"uint8\"},{\"internalType\":\"uint256\",\"name\":\"refundAmount\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"initialProposer\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"currentProposer\",\"type\":\"address\"},{\"internalType\":\"bool\",\"name\":\"ownerAccepted\",\"type\":\"bool\"},{\"internalType\":\"bool\",\"name\":\"supplierAccepted\",\"type\":\"bool\"},{\"internalType\":\"uint32\",\"name\":\"timesCountered\",\"type\":\"uint32\"},{\"internalType\":\"uint32\",\"name\":\"timesRejected\",\"type\":\"uint32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"tokenId\",\"type\":\"uint256\"}],\"name\":\"getCancellationReasons\",\"outputs\":[{\"internalType\":\"uint16\",\"name\":\"cancellationReason\",\"type\":\"uint16\"},{\"internalType\":\"uint16\",\"name\":\"cancellationVersion\",\"type\":\"uint16\"},{\"internalType\":\"uint16\",\"name\":\"rejectionReason\",\"type\":\"uint16\"},{\"internalType\":\"uint16\",\"name\":\"rejectionVersion\",\"type\":\"uint16\"},{\"internalType\":\"uint16\",\"name\":\"counterReason\",\"type\":\"uint16\"},{\"internalType\":\"uint16\",\"name\":\"counterVersion\",\"type\":\"uint16\"},{\"internalType\":\"uint16\",\"name\":\"withdrawalReason\",\"type\":\"uint16\"},{\"internalType\":\"uint16\",\"name\":\"withdrawalVersion\",\"type\":\"uint16\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"getManagerAddress\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"getMinExpirationTimestampDiff\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"tokenId\",\"type\":\"uint256\"}],\"name\":\"getReservationPaymentToken\",\"outputs\":[{\"internalType\":\"contractIERC20\",\"name\":\"paymentToken\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"tokenId\",\"type\":\"uint256\"}],\"name\":\"getReservationPrice\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"price\",\"type\":\"uint256\"},{\"internalType\":\"contractIERC20\",\"name\":\"paymentToken\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"role\",\"type\":\"bytes32\"}],\"name\":\"getRoleAdmin\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"role\",\"type\":\"bytes32\"},{\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"}],\"name\":\"grantRole\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"role\",\"type\":\"bytes32\"},{\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"}],\"name\":\"hasRole\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"manager\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"defaultAdmin\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"upgrader\",\"type\":\"address\"}],\"name\":\"initialize\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"tokenId\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"refundAmount\",\"type\":\"uint256\"},{\"internalType\":\"uint16\",\"name\":\"cancellationReason\",\"type\":\"uint16\"},{\"internalType\":\"uint16\",\"name\":\"cancellationReasonVersion\",\"type\":\"uint16\"}],\"name\":\"initiateCancellation\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"owner\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"operator\",\"type\":\"address\"}],\"name\":\"isApprovedForAll\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"tokenId\",\"type\":\"uint256\"}],\"name\":\"isCancellable\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"}],\"name\":\"isTTMAccount\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"name\",\"outputs\":[{\"internalType\":\"string\",\"name\":\"\",\"type\":\"string\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"tokenId\",\"type\":\"uint256\"}],\"name\":\"ownerOf\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"pause\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"paused\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"proxiableUUID\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"tokenId\",\"type\":\"uint256\"}],\"name\":\"recordExpiration\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"tokenId\",\"type\":\"uint256\"},{\"internalType\":\"uint16\",\"name\":\"rejectionReason\",\"type\":\"uint16\"},{\"internalType\":\"uint16\",\"name\":\"rejectionReasonVersion\",\"type\":\"uint16\"}],\"name\":\"rejectCancellation\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"role\",\"type\":\"bytes32\"},{\"internalType\":\"address\",\"name\":\"callerConfirmation\",\"type\":\"address\"}],\"name\":\"renounceRole\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"role\",\"type\":\"bytes32\"},{\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"}],\"name\":\"revokeRole\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"reservedFor\",\"type\":\"address\"},{\"internalType\":\"string\",\"name\":\"uri\",\"type\":\"string\"},{\"internalType\":\"uint256\",\"name\":\"expirationTimestamp\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"price\",\"type\":\"uint256\"},{\"internalType\":\"contractIERC20\",\"name\":\"paymentToken\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"offchainPaymentCurrency\",\"type\":\"uint256\"},{\"internalType\":\"bool\",\"name\":\"cancellable\",\"type\":\"bool\"}],\"name\":\"safeMintWithReservation\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"from\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"tokenId\",\"type\":\"uint256\"}],\"name\":\"safeTransferFrom\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"from\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"tokenId\",\"type\":\"uint256\"},{\"internalType\":\"bytes\",\"name\":\"data\",\"type\":\"bytes\"}],\"name\":\"safeTransferFrom\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"operator\",\"type\":\"address\"},{\"internalType\":\"bool\",\"name\":\"approved\",\"type\":\"bool\"}],\"name\":\"setApprovalForAll\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"manager\",\"type\":\"address\"}],\"name\":\"setManagerAddress\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"minExpirationTimestampDiff\",\"type\":\"uint256\"}],\"name\":\"setMinExpirationTimestampDiff\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes4\",\"name\":\"interfaceId\",\"type\":\"bytes4\"}],\"name\":\"supportsInterface\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"symbol\",\"outputs\":[{\"internalType\":\"string\",\"name\":\"\",\"type\":\"string\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"index\",\"type\":\"uint256\"}],\"name\":\"tokenByIndex\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"owner\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"index\",\"type\":\"uint256\"}],\"name\":\"tokenOfOwnerByIndex\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"tokenId\",\"type\":\"uint256\"}],\"name\":\"tokenURI\",\"outputs\":[{\"internalType\":\"string\",\"name\":\"\",\"type\":\"string\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"totalSupply\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"from\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"tokenId\",\"type\":\"uint256\"}],\"name\":\"transferFrom\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"unpause\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"newImplementation\",\"type\":\"address\"},{\"internalType\":\"bytes\",\"name\":\"data\",\"type\":\"bytes\"}],\"name\":\"upgradeToAndCall\",\"outputs\":[],\"stateMutability\":\"payable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"version\",\"outputs\":[{\"internalType\":\"uint16\",\"name\":\"major\",\"type\":\"uint16\"},{\"internalType\":\"uint16\",\"name\":\"minor\",\"type\":\"uint16\"},{\"internalType\":\"uint16\",\"name\":\"patch\",\"type\":\"uint16\"}],\"stateMutability\":\"pure\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"tokenId\",\"type\":\"uint256\"},{\"internalType\":\"uint16\",\"name\":\"withdrawalReason\",\"type\":\"uint16\"},{\"internalType\":\"uint16\",\"name\":\"withdrawalReasonVersion\",\"type\":\"uint16\"}],\"name\":\"withdrawCancellation\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"}]",
	Bin: "0x60a060405230608052348015610013575f80fd5b506080516155fa6200003b5f395f8181612bad01528181612bd60152612d7601526155fa5ff3fe60806040526004361061034f575f3560e01c80636352211e116101bd578063b88d4fde116100f2578063d547741f11610092578063e63ab1e91161006d578063e63ab1e914610c74578063e985e9c514610ca7578063f72c0d8b14610d0d578063f7e45f0914610d40575f80fd5b8063d547741f14610c17578063db2b268214610c36578063e5a6725c14610c55575f80fd5b8063bfb26c06116100cd578063bfb26c0614610b9c578063c0c53b8b14610bb0578063c162d7da14610bcf578063c87b56dd14610bf8575f80fd5b8063b88d4fde14610ab0578063bb520b4714610acf578063be66718814610b7d575f80fd5b806396591edd1161015d578063a22cb46511610138578063a22cb46514610903578063a9bc55a214610922578063ad3cb1cc14610a12578063b191d09214610a5a575f80fd5b806396591edd146108ca578063a0f07c74146108dd578063a217fddf146108f0575f80fd5b806374fe60e91161019857806374fe60e9146108205780638456cb591461083f57806391d148541461085357806395d89b41146108b6575f80fd5b80636352211e146107c357806370a08231146107e257806374aa204814610801575f80fd5b80632f2ff15d116102935780634f1ef2861161023357806352d1902d1161020e57806352d1902d1461072f57806354fd4d50146107435780635c975abb1461076e5780635f7290f1146107a4575f80fd5b80634f1ef286146106de5780634f6ccce7146106f1578063516a82b814610710575f80fd5b80633c15b31c1161026e5780633c15b31c146106325780633f4ba83a1461068c57806341431908146106a057806342842e0e146106bf575f80fd5b80632f2ff15d146105d55780632f745c59146105f457806336568abe14610613575f80fd5b806318160ddd116102fe578063248a9ca3116102d9578063248a9ca3146104e65780632a119380146105335780632d3a6329146105525780632edf5e2c146105a2575f80fd5b806318160ddd146104a05780631c54f0f7146104b457806323b872dd146104c7575f80fd5b8063081812fc1161032e578063081812fc14610426578063095ea7b31461045d5780630e75c1a81461047e575f80fd5b80624fdd3c1461035357806301ffc9a7146103d657806306fdde0314610405575b5f80fd5b34801561035e575f80fd5b506103b461036d366004614d82565b5f9081527f54347f9bbfbc2e7b5786abab9693d2ba67834fa4787a7740d72712e87e20740360205260409020600381015460049091015490916001600160a01b0390911690565b604080519283526001600160a01b039091166020830152015b60405180910390f35b3480156103e1575f80fd5b506103f56103f0366004614dae565b610d5f565b60405190151581526020016103cd565b348015610410575f80fd5b50610419610d6f565b6040516103cd9190614e16565b348015610431575f80fd5b50610445610440366004614d82565b610e23565b6040516001600160a01b0390911681526020016103cd565b348015610468575f80fd5b5061047c610477366004614e3c565b610e69565b005b348015610489575f80fd5b50610492610e78565b6040519081526020016103cd565b3480156104ab575f80fd5b50610492610e93565b61047c6104c2366004614e66565b610ebb565b3480156104d2575f80fd5b5061047c6104e1366004614e86565b610fe1565b3480156104f1575f80fd5b50610492610500366004614d82565b5f9081527f02dd7bc7dec4dceedda775e58dd541e08a116c6c53815c0bd028192f7b626800602052604090206001015490565b34801561053e575f80fd5b5061047c61054d366004614eda565b610ffa565b34801561055d575f80fd5b506103f561056c366004614d82565b5f9081527f54347f9bbfbc2e7b5786abab9693d2ba67834fa4787a7740d72712e87e207403602052604090206006015460ff1690565b3480156105ad575f80fd5b506104927f3ae8648b97d3fd425d26286fc6bb1d50724a93a6a5763921dd2b90405a83b4a481565b3480156105e0575f80fd5b5061047c6105ef366004614f13565b6110cc565b3480156105ff575f80fd5b5061049261060e366004614e3c565b611115565b34801561061e575f80fd5b5061047c61062d366004614f13565b611199565b34801561063d575f80fd5b5061067f61064c366004614d82565b5f9081527f54347f9bbfbc2e7b5786abab9693d2ba67834fa4787a7740d72712e87e207404602052604090205460ff1690565b6040516103cd9190614f71565b348015610697575f80fd5b5061047c6111e5565b3480156106ab575f80fd5b5061047c6106ba366004614f84565b61121a565b3480156106ca575f80fd5b5061047c6106d9366004614e86565b611253565b61047c6106ec366004615044565b61126d565b3480156106fc575f80fd5b5061049261070b366004614d82565b611288565b34801561071b575f80fd5b5061047c61072a366004614d82565b611300565b34801561073a575f80fd5b5061049261134f565b34801561074e575f80fd5b5060408051600181525f60208201819052918101919091526060016103cd565b348015610779575f80fd5b507fcd5ed15c6e187e77e9aee88184c21f4f2182ab5827cb3b7e07fbedcd63f033005460ff166103f5565b3480156107af575f80fd5b506103f56107be366004614f84565b61137d565b3480156107ce575f80fd5b506104456107dd366004614d82565b61141f565b3480156107ed575f80fd5b506104926107fc366004614f84565b611429565b34801561080c575f80fd5b5061047c61081b366004615091565b6114ad565b34801561082b575f80fd5b5061047c61083a366004614eda565b611577565b34801561084a575f80fd5b5061047c611640565b34801561085e575f80fd5b506103f561086d366004614f13565b5f9182527f02dd7bc7dec4dceedda775e58dd541e08a116c6c53815c0bd028192f7b626800602090815260408084206001600160a01b0393909316845291905290205460ff1690565b3480156108c1575f80fd5b50610419611672565b61047c6108d8366004614d82565b6116c3565b3480156108e8575f80fd5b506104455f81565b3480156108fb575f80fd5b506104925f81565b34801561090e575f80fd5b5061047c61091d3660046150e1565b611923565b34801561092d575f80fd5b506109c561093c366004614d82565b5f9081527f1973af827d98d3fe49ea0c9179b69c8f267978d372df89173953530663b7220060205260409020600281015460039091015461ffff600160c81b8304811693600160d81b8404821693600160e81b900482169282811692620100008204811692640100000000830482169266010000000000008104831692600160401b9091041690565b6040805161ffff998a16815297891660208901529588169587019590955292861660608601529085166080850152841660a0840152831660c083015290911660e0820152610100016103cd565b348015610a1d575f80fd5b506104196040518060400160405280600581526020017f352e302e3000000000000000000000000000000000000000000000000000000081525081565b348015610a65575f80fd5b50610445610a74366004614d82565b5f9081527f54347f9bbfbc2e7b5786abab9693d2ba67834fa4787a7740d72712e87e20740360205260409020600401546001600160a01b031690565b348015610abb575f80fd5b5061047c610aca36600461510d565b61192e565b348015610ada575f80fd5b50610b69610ae9366004614d82565b5f9081527f1973af827d98d3fe49ea0c9179b69c8f267978d372df89173953530663b72200602052604090206002810154815460019092015460ff600160c01b8084048216956001600160a01b038085169590811694928304841693600160c81b8404169263ffffffff600160a01b918290048116939190920490911690565b6040516103cd989796959493929190615175565b348015610b88575f80fd5b5061047c610b97366004614e66565b611946565b348015610ba7575f80fd5b50610445600181565b348015610bbb575f80fd5b5061047c610bca3660046151cd565b611a16565b348015610bda575f80fd5b505f805160206155a5833981519152546001600160a01b0316610445565b348015610c03575f80fd5b50610419610c12366004614d82565b611cc2565b348015610c22575f80fd5b5061047c610c31366004614f13565b611ccd565b348015610c41575f80fd5b5061047c610c50366004615215565b611d10565b348015610c60575f80fd5b5061047c610c6f366004614d82565b611f8a565b348015610c7f575f80fd5b506104927f65d7a28e3265b37a6474929f336521b332c1681b933f6cb9f3376673440d862a81565b348015610cb2575f80fd5b506103f5610cc13660046152b6565b6001600160a01b039182165f9081527f80bb2b638cc20bc4d0a60d66940f3ab4a00c1d7b313497ca82fb0b4ab00793056020908152604080832093909416825291909152205460ff1690565b348015610d18575f80fd5b506104927f189ab7a9244df0848122154315af71fe140f3db0fe014031783b0946b8c9d2e381565b348015610d4b575f80fd5b5061047c610d5a366004615091565b6120e3565b5f610d69826121ad565b92915050565b7f80bb2b638cc20bc4d0a60d66940f3ab4a00c1d7b313497ca82fb0b4ab00793008054606091908190610da1906152e2565b80601f0160208091040260200160405190810160405280929190818152602001828054610dcd906152e2565b8015610e185780601f10610def57610100808354040283529160200191610e18565b820191905f5260205f20905b815481529060010190602001808311610dfb57829003601f168201915b505050505091505090565b5f610e2d826121ea565b505f8281527f80bb2b638cc20bc4d0a60d66940f3ab4a00c1d7b313497ca82fb0b4ab007930460205260409020546001600160a01b0316610d69565b610e74828233612241565b5050565b5f805f805160206155a58339815191525b6002015492915050565b5f807f645e039705490088daad89bae25049a34f4a9072d398537b1ab2425f24cbed00610e89565b33610ec58161224e565b610ecd612298565b5f610ed7846121ea565b5f8581527f54347f9bbfbc2e7b5786abab9693d2ba67834fa4787a7740d72712e87e20740460205260409020549091505f805160206155a58339815191529060039060ff166004811115610f2d57610f2d614f41565b14610f6c575f858152600480830160205260409182902054915163e4e3b53b60e01b8152610f6392889260ff909116910161531a565b60405180910390fd5b5f8581526003820160205260408120600101546001600160a01b031690610f948288886122f6565b5f88815260038501602090815260408083206004908101548189019093529220805460ff19169092179091559091506001600160a01b0316610fd78183876124c3565b5050505050505050565b610fea8161258e565b610ff58383836127a5565b505050565b336110048161224e565b5f61100e856121ea565b5f8681527f54347f9bbfbc2e7b5786abab9693d2ba67834fa4787a7740d72712e87e20740460205260409020549091505f805160206155a58339815191529060039060ff16600481111561106457611064614f41565b1461109a575f868152600480830160205260409182902054915163e4e3b53b60e01b8152610f6392899260ff909116910161531a565b5f8681526003820160205260409020600101546001600160a01b03166110c38382898989612828565b50505050505050565b5f8281527f02dd7bc7dec4dceedda775e58dd541e08a116c6c53815c0bd028192f7b6268006020526040902060010154611105816129b6565b61110f83836129c0565b50505050565b5f7f645e039705490088daad89bae25049a34f4a9072d398537b1ab2425f24cbed0061114084611429565b83106111715760405163295f44f760e21b81526001600160a01b038516600482015260248101849052604401610f63565b6001600160a01b0384165f908152602091825260408082208583529092522054905092915050565b6001600160a01b03811633146111db576040517f6697b23200000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b610ff58282612a8c565b7f65d7a28e3265b37a6474929f336521b332c1681b933f6cb9f3376673440d862a61120f816129b6565b611217612b30565b50565b5f611224816129b6565b505f805160206155a583398151915280546001600160a01b0319166001600160a01b0392909216919091179055565b610ff583838360405180602001604052805f81525061192e565b611275612ba2565b61127e82612c59565b610e748282612c83565b5f7f645e039705490088daad89bae25049a34f4a9072d398537b1ab2425f24cbed006112b2610e93565b83106112da5760405163295f44f760e21b81525f600482015260248101849052604401610f63565b8060020183815481106112ef576112ef615337565b905f5260205f200154915050919050565b7f3ae8648b97d3fd425d26286fc6bb1d50724a93a6a5763921dd2b90405a83b4a461132a816129b6565b507f54347f9bbfbc2e7b5786abab9693d2ba67834fa4787a7740d72712e87e20740255565b5f611358612d6b565b507f360894a13ba1a3210667c828492db98dca3e2076cc3735a920a3ca505d382bbc90565b5f61139c5f805160206155a5833981519152546001600160a01b031690565b6040517f5f7290f10000000000000000000000000000000000000000000000000000000081526001600160a01b0384811660048301529190911690635f7290f190602401602060405180830381865afa1580156113fb573d5f803e3d5ffd5b505050506040513d601f19601f82011682018060405250810190610d69919061534b565b5f610d69826121ea565b5f7f80bb2b638cc20bc4d0a60d66940f3ab4a00c1d7b313497ca82fb0b4ab00793006001600160a01b03831661148d576040517f89c62b640000000000000000000000000000000000000000000000000000000081525f6004820152602401610f63565b6001600160a01b039092165f908152600390920160205250604090205490565b336114b78161224e565b5f6114c1866121ea565b5f8781527f54347f9bbfbc2e7b5786abab9693d2ba67834fa4787a7740d72712e87e20740460205260409020549091505f805160206155a58339815191529060039060ff16600481111561151757611517614f41565b1461154d575f878152600480830160205260409182902054915163e4e3b53b60e01b8152610f63928a9260ff909116910161531a565b5f8781526003820160205260409020600101546001600160a01b0316610fd783828a8a8a8a612db4565b336115818161224e565b5f61158b856121ea565b5f8681527f54347f9bbfbc2e7b5786abab9693d2ba67834fa4787a7740d72712e87e20740460205260409020549091505f805160206155a58339815191529060039060ff1660048111156115e1576115e1614f41565b14611617575f868152600480830160205260409182902054915163e4e3b53b60e01b8152610f6392899260ff909116910161531a565b5f8681526003820160205260409020600101546001600160a01b03166110c3838289898961306b565b7f65d7a28e3265b37a6474929f336521b332c1681b933f6cb9f3376673440d862a61166a816129b6565b61121761322e565b7f80bb2b638cc20bc4d0a60d66940f3ab4a00c1d7b313497ca82fb0b4ab007930180546060917f80bb2b638cc20bc4d0a60d66940f3ab4a00c1d7b313497ca82fb0b4ab007930091610da1906152e2565b6116cb613289565b6116d3612298565b336116dd8161224e565b5f8281527f54347f9bbfbc2e7b5786abab9693d2ba67834fa4787a7740d72712e87e2074036020908152604091829020825160e08101845281546001600160a01b0390811680835260018401548216948301949094526002830154948201949094526003820154606082015260048201549093166080840152600581015460a08401526006015460ff16151560c08301525f805160206155a5833981519152919033146117ca5780516040517f4cc7538a0000000000000000000000000000000000000000000000000000000081526001600160a01b039091166004820152336024820152604401610f63565b8060400151421115611819578381604001516040517f527ae76e000000000000000000000000000000000000000000000000000000008152600401610f63929190918252602082015260400190565b5f6118238561141f565b905081602001516001600160a01b0316816001600160a01b03161461188c5760208201516040517f81e8a2c8000000000000000000000000000000000000000000000000000000008152600481018790526001600160a01b039091166024820152604401610f63565b61189b826020015133876132ec565b6118b28260800151836060015184602001516124c3565b5f858152600484016020526040808220805460ff1916600317905551339187917fa751fb02c318279a22135a408663ae08ea45eafa950a4351c14ae543cbb950409190a35050505061121760017f9b779b17422d0df92223018b32b4d1fa46e071723d6817e2486d003becc55f0055565b610e74338383613399565b611939848484610fe1565b61110f3385858585613474565b336119508161224e565b5f61195a846121ea565b5f8581527f54347f9bbfbc2e7b5786abab9693d2ba67834fa4787a7740d72712e87e20740460205260409020549091505f805160206155a58339815191529060039060ff1660048111156119b0576119b0614f41565b146119e6575f858152600480830160205260409182902054915163e4e3b53b60e01b8152610f6392889260ff909116910161531a565b5f8581526003820160205260409020600101546001600160a01b0316611a0e8382888861359a565b505050505050565b7ff0c57e16840df040f15088dc2f81fe391c3923bec73e23a9662efc9c229c6a008054600160401b810460ff16159067ffffffffffffffff165f81158015611a5b5750825b90505f8267ffffffffffffffff166001148015611a775750303b155b905081158015611a85575080155b15611abc576040517ff92ee8a900000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b845467ffffffffffffffff191660011785558315611aeb57845468ff00000000000000001916600160401b1785555b6001600160a01b0388161580611b0857506001600160a01b038716155b80611b1a57506001600160a01b038616155b15611b51576040517fd92e233d00000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b611bc56040518060400160405280600c81526020017f426f6f6b696e67546f6b656e00000000000000000000000000000000000000008152506040518060400160405280600681526020017f42546f6b656e00000000000000000000000000000000000000000000000000008152506137e2565b611bcd6137f4565b611bd56137f4565b611bdd6137f4565b611be56137fc565b611bed6137f4565b611bf75f886129c0565b50611c227f189ab7a9244df0848122154315af71fe140f3db0fe014031783b0946b8c9d2e3876129c0565b505f805160206155a583398151915280546001600160a01b0319166001600160a01b038a16179055603c7f54347f9bbfbc2e7b5786abab9693d2ba67834fa4787a7740d72712e87e207402558315610fd757845468ff000000000000000019168555604051600181527fc7f505b2f371ae2175ee4913f4499e1f2633a7b5936321eed1cdaeb6115181d29060200160405180910390a15050505050505050565b6060610d698261380c565b5f8281527f02dd7bc7dec4dceedda775e58dd541e08a116c6c53815c0bd028192f7b6268006020526040902060010154611d06816129b6565b61110f8383612a8c565b33611d1a8161224e565b611d22612298565b611d2b8861224e565b7f54347f9bbfbc2e7b5786abab9693d2ba67834fa4787a7740d72712e87e207402545f805160206155a583398151915290611d66814261537a565b8811611da8576040517f999f7d700000000000000000000000000000000000000000000000000000000081526004810189905260248101829052604401610f63565b5f85118015611dc157506001600160a01b038616600114155b15611dfb576040517f8fe757e700000000000000000000000000000000000000000000000000000000815260048101869052602401610f63565b6001820180545f9182611e0d8361538d565b919050559050611e1d338261393b565b611e27818b613954565b6040805160e0810182526001600160a01b03808e1682523360208084019182528385018e8152606085018e8152848e166080870190815260a087018e81528d151560c089019081525f8b81527f54347f9bbfbc2e7b5786abab9693d2ba67834fa4787a7740d72712e87e20740390965298909420965187546001600160a01b0319908116918816919091178855945160018801805487169188169190911790559151600287015551600386015551600485018054909316931692909217905551600582015590516006909101805460ff19169115159190911790555f818152600484016020908152604091829020805460ff1916600117905581518b81529081018a90526001600160a01b0389811682840152606082018990528715156080830152915133928e169184917f1424af4f4cb40d8a1a2d00b2324cb122ba73eac426f98b62c33ff31ca045f0679160a0908290030190a45050505050505050505050565b5f8181527f54347f9bbfbc2e7b5786abab9693d2ba67834fa4787a7740d72712e87e207403602090815260408083207f54347f9bbfbc2e7b5786abab9693d2ba67834fa4787a7740d72712e87e207404909252909120545f805160206155a5833981519152919060ff16600281600481111561200857612008614f41565b14806120255750600381600481111561202357612023614f41565b145b806120415750600481600481111561203f5761203f614f41565b145b1561206357838160405163e4e3b53b60e01b8152600401610f6392919061531a565b81600201544211156120b5575f848152600484016020526040808220805460ff191660021790555185917fc47ab59d5c41a0220b594ece6f7e87863b07f8b33579130f3a59b31d8f7b6eed91a261110f565b815460405163d4cde2af60e01b8152600481018690526001600160a01b039091166024820152604401610f63565b336120ed8161224e565b5f6120f7866121ea565b5f8781527f54347f9bbfbc2e7b5786abab9693d2ba67834fa4787a7740d72712e87e20740460205260409020549091505f805160206155a58339815191529060039060ff16600481111561214d5761214d614f41565b14612183575f878152600480830160205260409182902054915163e4e3b53b60e01b8152610f63928a9260ff909116910161531a565b5f8781526003820160205260409020600101546001600160a01b0316610fd783828a8a8a8a6139c6565b5f6001600160e01b031982167f7965db0b000000000000000000000000000000000000000000000000000000001480610d695750610d6982613cf7565b5f8181527f80bb2b638cc20bc4d0a60d66940f3ab4a00c1d7b313497ca82fb0b4ab007930260205260408120546001600160a01b031680610d6957604051637e27328960e01b815260048101849052602401610f63565b610ff58383836001613d34565b6122578161137d565b611217576040517f27b12cd70000000000000000000000000000000000000000000000000000000081526001600160a01b0382166004820152602401610f63565b7fcd5ed15c6e187e77e9aee88184c21f4f2182ab5827cb3b7e07fbedcd63f033005460ff16156122f4576040517fd93c066500000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b565b5f336001600160a01b0385161461233c576040517f6c83fb1b00000000000000000000000000000000000000000000000000000000815260048101849052602401610f63565b5f8381527f1973af827d98d3fe49ea0c9179b69c8f267978d372df89173953530663b72200602052604090206002810154600190600160c01b900460ff16600481111561238b5761238b614f41565b146123be57838160020160189054906101000a900460ff1660405163d5d8eacb60e01b8152600401610f6392919061531a565b805483146123f357805460405163cc45283760e01b815260048101869052602481019190915260448101849052606401610f63565b6001810154600160c01b900460ff1661243b576040517fc84052f900000000000000000000000000000000000000000000000000000000815260048101859052602401610f63565b6001810154600160c81b900460ff166124645760018101805460ff60c81b1916600160c81b1790555b60028101805460ff60c01b1916780400000000000000000000000000000000000000000000000017905560405184907f17c3690813e5ff9135b87fd91848109978b23db8e471498d18886560da7f2867905f90a25490505b9392505050565b6001600160a01b03831661252657813414612513576040517f0515845400000000000000000000000000000000000000000000000000000000815234600482015260248101839052604401610f63565b610ff56001600160a01b03821634613eb0565b5f196001600160a01b03841601612558573415610ff5576040516347d6729960e01b8152346004820152602401610f63565b3415612579576040516347d6729960e01b8152346004820152602401610f63565b610ff56001600160a01b038416338385613f63565b5f8181527f54347f9bbfbc2e7b5786abab9693d2ba67834fa4787a7740d72712e87e20740460209081526040808320547f1973af827d98d3fe49ea0c9179b69c8f267978d372df89173953530663b72200909252909120600201545f805160206155a58339815191529160ff90811691600160c01b8104909116906001600160a01b0316600182600481111561262657612626614f41565b03612685575f612635866121ea565b5f8781526003870160205260409020600101549091506001600160a01b0390811690831633146126735761266e8282896063600161306b565b612682565b61268282828960636001612828565b50505b600383600481111561269957612699614f41565b14806126b6575060028360048111156126b4576126b4614f41565b145b806126d157505f8360048111156126cf576126cf614f41565b145b156126dd575050505050565b60048360048111156126f1576126f1614f41565b0361271357848360405163e4e3b53b60e01b8152600401610f6392919061531a565b5f85815260038501602052604090206002810154421115612777575f868152600486016020526040808220805460ff191660021790555187917fc47ab59d5c41a0220b594ece6f7e87863b07f8b33579130f3a59b31d8f7b6eed91a2505050505050565b805460405163d4cde2af60e01b8152600481018890526001600160a01b039091166024820152604401610f63565b6001600160a01b0382166127ce57604051633250574960e11b81525f6004820152602401610f63565b5f6127da838333613feb565b9050836001600160a01b0316816001600160a01b03161461110f576040516364283d7b60e01b81526001600160a01b0380861660048301526024820184905282166044820152606401610f63565b84846128348282613fff565b5f8581527f1973af827d98d3fe49ea0c9179b69c8f267978d372df89173953530663b72200602052604090206002810154600190600160c01b900460ff16600481111561288357612883614f41565b146128b657858160020160189054906101000a900460ff1660405163d5d8eacb60e01b8152600401610f6392919061531a565b60028101546001600160a01b031633146128ff576040517f8c37ede800000000000000000000000000000000000000000000000000000000815260048101879052602401610f63565b60038101805469ffffffff0000000000001916660100000000000061ffff88811691820269ffff0000000000000000191692909217600160401b92881692830217909255600283018054780300000000000000000000000000000000000000000000000060ff60c01b1990911617905560408051928352602083019190915287917f48e256ce3da490e3bbba80f056bb54ec3d7264f8ad7d152b77bf8c2eca3db5a591015b60405180910390a25050505050505050565b6112178133614058565b5f8281527f02dd7bc7dec4dceedda775e58dd541e08a116c6c53815c0bd028192f7b626800602081815260408084206001600160a01b038616855290915282205460ff16612a83575f848152602082815260408083206001600160a01b03871684529091529020805460ff19166001179055612a393390565b6001600160a01b0316836001600160a01b0316857f2f8788117e7eff1d82e926ec794901d17c78024a50270940304540a733656f0d60405160405180910390a46001915050610d69565b5f915050610d69565b5f8281527f02dd7bc7dec4dceedda775e58dd541e08a116c6c53815c0bd028192f7b626800602081815260408084206001600160a01b038616855290915282205460ff1615612a83575f848152602082815260408083206001600160a01b0387168085529252808320805460ff1916905551339287917ff6391f5c32d9c69d2a47ea670b442974b53935d1edc7fd64eb21e047a839171b9190a46001915050610d69565b612b386140e4565b7fcd5ed15c6e187e77e9aee88184c21f4f2182ab5827cb3b7e07fbedcd63f03300805460ff191681557f5db9ee0a495bf2e6ff9c91a7834c1ba4fdd244a5e8aa4e537bd38aeae4b073aa335b6040516001600160a01b03909116815260200160405180910390a150565b306001600160a01b037f0000000000000000000000000000000000000000000000000000000000000000161480612c3b57507f00000000000000000000000000000000000000000000000000000000000000006001600160a01b0316612c2f7f360894a13ba1a3210667c828492db98dca3e2076cc3735a920a3ca505d382bbc546001600160a01b031690565b6001600160a01b031614155b156122f45760405163703e46dd60e11b815260040160405180910390fd5b7f189ab7a9244df0848122154315af71fe140f3db0fe014031783b0946b8c9d2e3610e74816129b6565b816001600160a01b03166352d1902d6040518163ffffffff1660e01b8152600401602060405180830381865afa925050508015612cdd575060408051601f3d908101601f19168201909252612cda918101906153a5565b60015b612d0557604051634c9c8ce360e01b81526001600160a01b0383166004820152602401610f63565b7f360894a13ba1a3210667c828492db98dca3e2076cc3735a920a3ca505d382bbc8114612d61576040517faa1d49a400000000000000000000000000000000000000000000000000000000815260048101829052602401610f63565b610ff5838361413f565b306001600160a01b037f000000000000000000000000000000000000000000000000000000000000000016146122f45760405163703e46dd60e11b815260040160405180910390fd5b8585612dc08282613fff565b5f8681527f1973af827d98d3fe49ea0c9179b69c8f267978d372df89173953530663b72200602052604090206002810154600190600160c01b900460ff166004811115612e0f57612e0f614f41565b14612e4257868160020160189054906101000a900460ff1660405163d5d8eacb60e01b8152600401610f6392919061531a565b858155600281018054336001600160a01b031990911681179091556001820180547fffffffffffff0000ffffffffffffffffffffffffffffffffffffffffffffffff166001600160a01b038c81168414600160c01b0260ff60c81b191691909117908b16909214600160c81b029190911780825560038301805465ffffffff000019166201000061ffff8a81169190910265ffff0000000019169190911764010000000091891691909102179055600160a01b900463ffffffff16906014612f09836153bc565b82546101009290920a63ffffffff81810219909316918316021790915560028301546001840154845460408051918252600160c01b830460ff90811615156020840152600160c81b840416151590820152600160a01b80830485166060830152830490931660808401526001600160a01b039182169350169089907f6aafaf6639750382d6dc665bc6a3c73d54e2076a8739a69c860a2314032eb9819060a0015b60405180910390a46002810154600382015460408051600160c81b840461ffff9081168252600160d81b850481166020830152600160e81b9094048416818301528383166060820152620100008304841660808201526401000000008304841660a082015266010000000000008304841660c0820152600160401b90920490921660e0820152905188917fade2ce8b39037f5140109f53fef47c0491427b8b785ab83d35aa2931295cb79091908190036101000190a2505050505050505050565b84846130778282613fff565b5f8581527f1973af827d98d3fe49ea0c9179b69c8f267978d372df89173953530663b72200602052604090206002810154600190600160c01b900460ff1660048111156130c6576130c6614f41565b146130f957858160020160189054906101000a900460ff1660405163d5d8eacb60e01b8152600401610f6392919061531a565b60028101546001600160a01b03163303613142576040517fc18363dc00000000000000000000000000000000000000000000000000000000815260048101879052602401610f63565b60028101805460038301805461ffff88811661ffff19909216919091179091558716600160e81b0260ff60c01b19167fff0000ffffffff00ffffffffffffffffffffffffffffffffffffffffffffffff90911617780200000000000000000000000000000000000000000000000017808255600160a01b900463ffffffff169060146131cd836153bc565b91906101000a81548163ffffffff021916908363ffffffff16021790555050857fab78ba855f2fdb28beb212a9b3f41a33cda034729848cd452f0cc96528c23a8086866040516129a492919061ffff92831681529116602082015260400190565b613236612298565b7fcd5ed15c6e187e77e9aee88184c21f4f2182ab5827cb3b7e07fbedcd63f03300805460ff191660011781557f62e78cea01bee320cd4e420270b5ea74000d11b0c9f74754ebdbfc544b05a25833612b84565b7f9b779b17422d0df92223018b32b4d1fa46e071723d6817e2486d003becc55f008054600119016132e6576040517f3ee5aeb500000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b60029055565b6001600160a01b03821661331557604051633250574960e11b81525f6004820152602401610f63565b5f61332183835f613feb565b90506001600160a01b03811661334d57604051637e27328960e01b815260048101839052602401610f63565b836001600160a01b0316816001600160a01b03161461110f576040516364283d7b60e01b81526001600160a01b0380861660048301526024820184905282166044820152606401610f63565b7f80bb2b638cc20bc4d0a60d66940f3ab4a00c1d7b313497ca82fb0b4ab00793006001600160a01b038316613405576040517f5b08ba180000000000000000000000000000000000000000000000000000000081526001600160a01b0384166004820152602401610f63565b6001600160a01b038481165f818152600584016020908152604080832094881680845294825291829020805460ff191687151590811790915591519182527f17307eab39ab6107e8899845ad3d59bd9653f200f220920489ca2b5937696c31910160405180910390a350505050565b6001600160a01b0383163b1561359357604051630a85bd0160e11b81526001600160a01b0384169063150b7a02906134b69088908890879087906004016153de565b6020604051808303815f875af19250505080156134f0575060408051601f3d908101601f191682019092526134ed91810190615419565b60015b613557573d80801561351d576040519150601f19603f3d011682016040523d82523d5f602084013e613522565b606091505b5080515f0361354f57604051633250574960e11b81526001600160a01b0385166004820152602401610f63565b805181602001fd5b6001600160e01b03198116630a85bd0160e11b14611a0e57604051633250574960e11b81526001600160a01b0385166004820152602401610f63565b5050505050565b83836135a68282613fff565b5f8481527f1973af827d98d3fe49ea0c9179b69c8f267978d372df89173953530663b72200602052604090206002810154600190600160c01b900460ff1660048111156135f5576135f5614f41565b1461362857848160020160189054906101000a900460ff1660405163d5d8eacb60e01b8152600401610f6392919061531a565b8054841461365d57805460405163cc45283760e01b815260048101879052602481019190915260448101859052606401610f63565b6001600160a01b03871633036136875760018101805460ff60c01b1916600160c01b17905561369d565b60018101805460ff60c81b1916600160c81b1790555b6002810154600182015482546040805191825260ff600160c01b8404811615156020840152600160c81b84041615159082015263ffffffff600160a01b8084048216606084015284041660808201526001600160a01b03928316929091169087907f6aafaf6639750382d6dc665bc6a3c73d54e2076a8739a69c860a2314032eb9819060a00160405180910390a46002810154600382015460408051600160c81b840461ffff9081168252600160d81b850481166020830152600160e81b9094048416818301528383166060820152620100008304841660808201526401000000008304841660a082015266010000000000008304841660c0820152600160401b90920490921660e0820152905186917fade2ce8b39037f5140109f53fef47c0491427b8b785ab83d35aa2931295cb79091908190036101000190a250505050505050565b6137ea614194565b610e7482826141f6565b6122f4614194565b613804614194565b6122f4614239565b60607f0542a41881ee128a365a727b282c86fa859579490b9bb45aab8503648c8e7900613838836121ea565b505f8381526020829052604081208054613851906152e2565b80601f016020809104026020016040519081016040528092919081815260200182805461387d906152e2565b80156138c85780601f1061389f576101008083540402835291602001916138c8565b820191905f5260205f20905b8154815290600101906020018083116138ab57829003601f168201915b505050505090505f6138e460408051602081019091525f815290565b905080515f036138f657509392505050565b815115613929578082604051602001613910929190615434565b6040516020818303038152906040529350505050919050565b6139328561426c565b95945050505050565b610e74828260405180602001604052805f8152506142dc565b5f8281527f0542a41881ee128a365a727b282c86fa859579490b9bb45aab8503648c8e79006020819052604090912061398d83826154a6565b506040518381527ff8e1a15aba9398e019f0b49df1a4fde98ee17ae345cb5f6b5e2c27f5033e8ce79060200160405180910390a1505050565b85856139d28282613fff565b5f8681527f1973af827d98d3fe49ea0c9179b69c8f267978d372df89173953530663b72200602052604090206002810154600490600160c01b900460ff1681811115613a2057613a20614f41565b1480613a4b575060016002820154600160c01b900460ff166004811115613a4957613a49614f41565b145b15613a7e57868160020160189054906101000a900460ff1660405163d5d8eacb60e01b8152600401610f6392919061531a565b5f6002820154600160c01b900460ff166004811115613a9f57613a9f614f41565b03613ab9576001810180546001600160a01b031916331790555b33816002015f6101000a8154816001600160a01b0302191690836001600160a01b0316021790555085815f0181905550886001600160a01b0316336001600160a01b0316148160010160186101000a81548160ff021916908315150217905550876001600160a01b0316336001600160a01b0316148160010160196101000a81548160ff021916908315150217905550848160020160196101000a81548161ffff021916908361ffff1602179055508381600201601b6101000a81548161ffff021916908361ffff1602179055505f81600201601d6101000a81548161ffff021916908361ffff1602179055505f816003015f6101000a81548161ffff021916908361ffff1602179055505f8160030160026101000a81548161ffff021916908361ffff1602179055505f8160030160046101000a81548161ffff021916908361ffff1602179055505f8160030160066101000a81548161ffff021916908361ffff1602179055505f8160030160086101000a81548161ffff021916908361ffff16021790555060018160020160186101000a81548160ff02191690836004811115613c6757613c67614f41565b02179055506002810154600182015482546040805191825260ff600160c01b8404811615156020840152600160c81b84041615159082015263ffffffff600160a01b8084048216606084015284041660808201526001600160a01b03928316929091169089907f6aafaf6639750382d6dc665bc6a3c73d54e2076a8739a69c860a2314032eb9819060a001612faa565b5f6001600160e01b031982167f49064906000000000000000000000000000000000000000000000000000000001480610d695750610d69826142f3565b7f80bb2b638cc20bc4d0a60d66940f3ab4a00c1d7b313497ca82fb0b4ab00793008180613d6957506001600160a01b03831615155b15613e80575f613d78856121ea565b90506001600160a01b03841615801590613da45750836001600160a01b0316816001600160a01b031614155b8015613df457506001600160a01b038082165f9081527f80bb2b638cc20bc4d0a60d66940f3ab4a00c1d7b313497ca82fb0b4ab0079305602090815260408083209388168352929052205460ff16155b15613e36576040517fa9fbf51f0000000000000000000000000000000000000000000000000000000081526001600160a01b0385166004820152602401610f63565b8215613e7e5784866001600160a01b0316826001600160a01b03167f8c5be1e5ebec7d5bd14f71427d1e84f3dd0314c0f7b2291e5b200ac8c7c3b92560405160405180910390a45b505b5f93845260040160205250506040902080546001600160a01b0319166001600160a01b0392909216919091179055565b80471015613ef3576040517fcf47918100000000000000000000000000000000000000000000000000000000815247600482015260248101829052604401610f63565b5f826001600160a01b0316826040515f6040518083038185875af1925050503d805f8114613f3c576040519150601f19603f3d011682016040523d82523d5f602084013e613f41565b606091505b5050905080610ff55760405163d6bda27560e01b815260040160405180910390fd5b604080516001600160a01b0385811660248301528416604482015260648082018490528251808303909101815260849091019091526020810180517bffffffffffffffffffffffffffffffffffffffffffffffffffffffff167f23b872dd0000000000000000000000000000000000000000000000000000000017905261110f908590614330565b5f613ff78484846143b5565b949350505050565b336001600160a01b038316148015906140215750336001600160a01b03821614155b15610e74576040517f4793d28100000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b5f8281527f02dd7bc7dec4dceedda775e58dd541e08a116c6c53815c0bd028192f7b626800602090815260408083206001600160a01b038516845290915290205460ff16610e74576040517fe2517d3f0000000000000000000000000000000000000000000000000000000081526001600160a01b038216600482015260248101839052604401610f63565b7fcd5ed15c6e187e77e9aee88184c21f4f2182ab5827cb3b7e07fbedcd63f033005460ff166122f4576040517f8dfc202b00000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b614148826144be565b6040516001600160a01b038316907fbc7cd75a20ee27fd9adebab32041f755214dbc6bffa90cc0225b39da2e5c2d3b905f90a280511561418c57610ff58282614534565b610e7461459d565b7ff0c57e16840df040f15088dc2f81fe391c3923bec73e23a9662efc9c229c6a0054600160401b900460ff166122f4576040517fd7e6bcf800000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b6141fe614194565b7f80bb2b638cc20bc4d0a60d66940f3ab4a00c1d7b313497ca82fb0b4ab00793008061422a84826154a6565b506001810161110f83826154a6565b614241614194565b7fcd5ed15c6e187e77e9aee88184c21f4f2182ab5827cb3b7e07fbedcd63f03300805460ff19169055565b6060614277826121ea565b505f61428d60408051602081019091525f815290565b90505f8151116142ab5760405180602001604052805f8152506124bc565b806142b5846145d5565b6040516020016142c6929190615434565b6040516020818303038152906040529392505050565b6142e68383614672565b610ff5335f858585613474565b5f6001600160e01b031982167f780e9d63000000000000000000000000000000000000000000000000000000001480610d695750610d69826146ec565b5f8060205f8451602086015f885af18061434f576040513d5f823e3d81fd5b50505f513d91508115614366578060011415614373565b6001600160a01b0384163b155b1561110f576040517f5274afe70000000000000000000000000000000000000000000000000000000081526001600160a01b0385166004820152602401610f63565b5f806143c2858585614786565b90506001600160a01b03811661445c57614457847f645e039705490088daad89bae25049a34f4a9072d398537b1ab2425f24cbed0280545f8381527f645e039705490088daad89bae25049a34f4a9072d398537b1ab2425f24cbed0360205260408120829055600182018355919091527fa42f15e5d656f8155fd7419d740a6073999f19cd6e061449ce4a257150545bf20155565b61447f565b846001600160a01b0316816001600160a01b03161461447f5761447f81856148c0565b6001600160a01b03851661449b576144968461496a565b613ff7565b846001600160a01b0316816001600160a01b031614613ff757613ff78585614a5d565b806001600160a01b03163b5f036144f357604051634c9c8ce360e01b81526001600160a01b0382166004820152602401610f63565b7f360894a13ba1a3210667c828492db98dca3e2076cc3735a920a3ca505d382bbc80546001600160a01b0319166001600160a01b0392909216919091179055565b60605f80846001600160a01b0316846040516145509190615562565b5f60405180830381855af49150503d805f8114614588576040519150601f19603f3d011682016040523d82523d5f602084013e61458d565b606091505b5091509150613932858383614ac8565b34156122f4576040517fb398979f00000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b60605f6145e183614b3d565b60010190505f8167ffffffffffffffff81111561460057614600614f9f565b6040519080825280601f01601f19166020018201604052801561462a576020820181803683370190505b5090508181016020015b5f19017f3031323334353637383961626364656600000000000000000000000000000000600a86061a8153600a850494508461463457509392505050565b6001600160a01b03821661469b57604051633250574960e11b81525f6004820152602401610f63565b5f6146a783835f613feb565b90506001600160a01b03811615610ff5576040517f73c6ac6e0000000000000000000000000000000000000000000000000000000081525f6004820152602401610f63565b5f6001600160e01b031982167f80ac58cd00000000000000000000000000000000000000000000000000000000148061474e57506001600160e01b031982167f5b5e139f00000000000000000000000000000000000000000000000000000000145b80610d6957507f01ffc9a7000000000000000000000000000000000000000000000000000000006001600160e01b0319831614610d69565b5f8281527f80bb2b638cc20bc4d0a60d66940f3ab4a00c1d7b313497ca82fb0b4ab007930260205260408120547f80bb2b638cc20bc4d0a60d66940f3ab4a00c1d7b313497ca82fb0b4ab0079300906001600160a01b03908116908416156147f3576147f3818587614c1e565b6001600160a01b0381161561482f5761480e5f865f80613d34565b6001600160a01b0381165f908152600383016020526040902080545f190190555b6001600160a01b0386161561485f576001600160a01b0386165f9081526003830160205260409020805460010190555b5f85815260028301602052604080822080546001600160a01b0319166001600160a01b038a811691821790925591518893918516917fddf252ad1be2c89b69c2b068fc378daa952ba7f163c4a11628f55a4df523b3ef91a495945050505050565b7f645e039705490088daad89bae25049a34f4a9072d398537b1ab2425f24cbed005f6148eb84611429565b5f8481526001840160209081526040808320546001600160a01b03891684529186905290912091925090818314614943575f838152602082815260408083205485845281842081905583526001870190915290208290555b5f948552600190930160209081526040808620869055928552929092528220919091555050565b7f645e039705490088daad89bae25049a34f4a9072d398537b1ab2425f24cbed02547f645e039705490088daad89bae25049a34f4a9072d398537b1ab2425f24cbed00905f906149bc9060019061557d565b5f8481526003840160205260408120546002850180549394509092849081106149e7576149e7615337565b905f5260205f200154905080846002018381548110614a0857614a08615337565b5f91825260208083209091019290925582815260038601909152604080822084905586825281205560028401805480614a4357614a43615590565b600190038181905f5260205f20015f905590555050505050565b7f645e039705490088daad89bae25049a34f4a9072d398537b1ab2425f24cbed005f6001614a8a85611429565b614a94919061557d565b6001600160a01b039094165f9081526020838152604080832087845282528083208690559482526001909301909252502055565b606082614add57614ad882614c9b565b6124bc565b8151158015614af457506001600160a01b0384163b155b15614b36576040517f9996b3150000000000000000000000000000000000000000000000000000000081526001600160a01b0385166004820152602401610f63565b50806124bc565b5f807a184f03e93ff9f4daa797ed6e38ed64bf6a1f0100000000000000008310614b85577a184f03e93ff9f4daa797ed6e38ed64bf6a1f010000000000000000830492506040015b6d04ee2d6d415b85acef81000000008310614bb1576d04ee2d6d415b85acef8100000000830492506020015b662386f26fc100008310614bcf57662386f26fc10000830492506010015b6305f5e1008310614be7576305f5e100830492506008015b6127108310614bfb57612710830492506004015b60648310614c0d576064830492506002015b600a8310610d695760010192915050565b614c29838383614cc4565b610ff5576001600160a01b038316614c5757604051637e27328960e01b815260048101829052602401610f63565b6040517f177e802f0000000000000000000000000000000000000000000000000000000081526001600160a01b038316600482015260248101829052604401610f63565b805115614cab5780518082602001fd5b60405163d6bda27560e01b815260040160405180910390fd5b5f6001600160a01b03831615801590613ff75750826001600160a01b0316846001600160a01b03161480614d3b57506001600160a01b038085165f9081527f80bb2b638cc20bc4d0a60d66940f3ab4a00c1d7b313497ca82fb0b4ab0079305602090815260408083209387168352929052205460ff165b80613ff75750505f9081527f80bb2b638cc20bc4d0a60d66940f3ab4a00c1d7b313497ca82fb0b4ab007930460205260409020546001600160a01b03908116911614919050565b5f60208284031215614d92575f80fd5b5035919050565b6001600160e01b031981168114611217575f80fd5b5f60208284031215614dbe575f80fd5b81356124bc81614d99565b5f5b83811015614de3578181015183820152602001614dcb565b50505f910152565b5f8151808452614e02816020860160208601614dc9565b601f01601f19169290920160200192915050565b602081525f6124bc6020830184614deb565b6001600160a01b0381168114611217575f80fd5b5f8060408385031215614e4d575f80fd5b8235614e5881614e28565b946020939093013593505050565b5f8060408385031215614e77575f80fd5b50508035926020909101359150565b5f805f60608486031215614e98575f80fd5b8335614ea381614e28565b92506020840135614eb381614e28565b929592945050506040919091013590565b803561ffff81168114614ed5575f80fd5b919050565b5f805f60608486031215614eec575f80fd5b83359250614efc60208501614ec4565b9150614f0a60408501614ec4565b90509250925092565b5f8060408385031215614f24575f80fd5b823591506020830135614f3681614e28565b809150509250929050565b634e487b7160e01b5f52602160045260245ffd5b6005811061121757634e487b7160e01b5f52602160045260245ffd5b60208101614f7e83614f55565b91905290565b5f60208284031215614f94575f80fd5b81356124bc81614e28565b634e487b7160e01b5f52604160045260245ffd5b5f67ffffffffffffffff80841115614fcd57614fcd614f9f565b604051601f8501601f19908116603f01168101908282118183101715614ff557614ff5614f9f565b8160405280935085815286868601111561500d575f80fd5b858560208301375f602087830101525050509392505050565b5f82601f830112615035575f80fd5b6124bc83833560208501614fb3565b5f8060408385031215615055575f80fd5b823561506081614e28565b9150602083013567ffffffffffffffff81111561507b575f80fd5b61508785828601615026565b9150509250929050565b5f805f80608085870312156150a4575f80fd5b84359350602085013592506150bb60408601614ec4565b91506150c960608601614ec4565b905092959194509250565b8015158114611217575f80fd5b5f80604083850312156150f2575f80fd5b82356150fd81614e28565b91506020830135614f36816150d4565b5f805f8060808587031215615120575f80fd5b843561512b81614e28565b9350602085013561513b81614e28565b925060408501359150606085013567ffffffffffffffff81111561515d575f80fd5b61516987828801615026565b91505092959194509250565b61010081016151838a614f55565b98815260208101979097526001600160a01b0395861660408801529390941660608601529015156080850152151560a084015263ffffffff91821660c08401521660e09091015290565b5f805f606084860312156151df575f80fd5b83356151ea81614e28565b925060208401356151fa81614e28565b9150604084013561520a81614e28565b809150509250925092565b5f805f805f805f60e0888a03121561522b575f80fd5b873561523681614e28565b9650602088013567ffffffffffffffff811115615251575f80fd5b8801601f81018a13615261575f80fd5b6152708a823560208401614fb3565b9650506040880135945060608801359350608088013561528f81614e28565b925060a0880135915060c08801356152a6816150d4565b8091505092959891949750929550565b5f80604083850312156152c7575f80fd5b82356152d281614e28565b91506020830135614f3681614e28565b600181811c908216806152f657607f821691505b60208210810361531457634e487b7160e01b5f52602260045260245ffd5b50919050565b8281526040810161532a83614f55565b8260208301529392505050565b634e487b7160e01b5f52603260045260245ffd5b5f6020828403121561535b575f80fd5b81516124bc816150d4565b634e487b7160e01b5f52601160045260245ffd5b80820180821115610d6957610d69615366565b5f6001820161539e5761539e615366565b5060010190565b5f602082840312156153b5575f80fd5b5051919050565b5f63ffffffff8083168181036153d4576153d4615366565b6001019392505050565b5f6001600160a01b0380871683528086166020840152508360408301526080606083015261540f6080830184614deb565b9695505050505050565b5f60208284031215615429575f80fd5b81516124bc81614d99565b5f8351615445818460208801614dc9565b835190830190615459818360208801614dc9565b01949350505050565b601f821115610ff557805f5260205f20601f840160051c810160208510156154875750805b601f840160051c820191505b81811015613593575f8155600101615493565b815167ffffffffffffffff8111156154c0576154c0614f9f565b6154d4816154ce84546152e2565b84615462565b602080601f831160018114615507575f84156154f05750858301515b5f19600386901b1c1916600185901b178555611a0e565b5f85815260208120601f198616915b8281101561553557888601518255948401946001909101908401615516565b508582101561555257878501515f19600388901b60f8161c191681555b5050505050600190811b01905550565b5f8251615573818460208701614dc9565b9190910192915050565b81810381811115610d6957610d69615366565b634e487b7160e01b5f52603160045260245ffdfe54347f9bbfbc2e7b5786abab9693d2ba67834fa4787a7740d72712e87e207400a2646970667358221220fe6181c89ac5fa38ae3d6b48f6001e67b4846a08b05762e434048d6e9b677ec564736f6c63430008180033",
}

// BookingtokenABI is the input ABI used to generate the binding from.
// Deprecated: Use BookingtokenMetaData.ABI instead.
var BookingtokenABI = BookingtokenMetaData.ABI

// BookingtokenBin is the compiled bytecode used for deploying new contracts.
// Deprecated: Use BookingtokenMetaData.Bin instead.
var BookingtokenBin = BookingtokenMetaData.Bin

// DeployBookingtoken deploys a new Ethereum contract, binding an instance of Bookingtoken to it.
func DeployBookingtoken(auth *bind.TransactOpts, backend bind.ContractBackend) (common.Address, *types.Transaction, *Bookingtoken, error) {
	parsed, err := BookingtokenMetaData.GetAbi()
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	if parsed == nil {
		return common.Address{}, nil, nil, errors.New("GetABI returned nil")
	}

	address, tx, contract, err := bind.DeployContract(auth, *parsed, common.FromHex(BookingtokenBin), backend)
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	return address, tx, &Bookingtoken{BookingtokenCaller: BookingtokenCaller{contract: contract}, BookingtokenTransactor: BookingtokenTransactor{contract: contract}, BookingtokenFilterer: BookingtokenFilterer{contract: contract}}, nil
}

// Bookingtoken is an auto generated Go binding around an Ethereum contract.
type Bookingtoken struct {
	BookingtokenCaller     // Read-only binding to the contract
	BookingtokenTransactor // Write-only binding to the contract
	BookingtokenFilterer   // Log filterer for contract events
}

// BookingtokenCaller is an auto generated read-only Go binding around an Ethereum contract.
type BookingtokenCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// BookingtokenTransactor is an auto generated write-only Go binding around an Ethereum contract.
type BookingtokenTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// BookingtokenFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type BookingtokenFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// BookingtokenSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type BookingtokenSession struct {
	Contract     *Bookingtoken     // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// BookingtokenCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type BookingtokenCallerSession struct {
	Contract *BookingtokenCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts       // Call options to use throughout this session
}

// BookingtokenTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type BookingtokenTransactorSession struct {
	Contract     *BookingtokenTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts       // Transaction auth options to use throughout this session
}

// BookingtokenRaw is an auto generated low-level Go binding around an Ethereum contract.
type BookingtokenRaw struct {
	Contract *Bookingtoken // Generic contract binding to access the raw methods on
}

// BookingtokenCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type BookingtokenCallerRaw struct {
	Contract *BookingtokenCaller // Generic read-only contract binding to access the raw methods on
}

// BookingtokenTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type BookingtokenTransactorRaw struct {
	Contract *BookingtokenTransactor // Generic write-only contract binding to access the raw methods on
}

// NewBookingtoken creates a new instance of Bookingtoken, bound to a specific deployed contract.
func NewBookingtoken(address common.Address, backend bind.ContractBackend) (*Bookingtoken, error) {
	contract, err := bindBookingtoken(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &Bookingtoken{BookingtokenCaller: BookingtokenCaller{contract: contract}, BookingtokenTransactor: BookingtokenTransactor{contract: contract}, BookingtokenFilterer: BookingtokenFilterer{contract: contract}}, nil
}

// NewBookingtokenCaller creates a new read-only instance of Bookingtoken, bound to a specific deployed contract.
func NewBookingtokenCaller(address common.Address, caller bind.ContractCaller) (*BookingtokenCaller, error) {
	contract, err := bindBookingtoken(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &BookingtokenCaller{contract: contract}, nil
}

// NewBookingtokenTransactor creates a new write-only instance of Bookingtoken, bound to a specific deployed contract.
func NewBookingtokenTransactor(address common.Address, transactor bind.ContractTransactor) (*BookingtokenTransactor, error) {
	contract, err := bindBookingtoken(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &BookingtokenTransactor{contract: contract}, nil
}

// NewBookingtokenFilterer creates a new log filterer instance of Bookingtoken, bound to a specific deployed contract.
func NewBookingtokenFilterer(address common.Address, filterer bind.ContractFilterer) (*BookingtokenFilterer, error) {
	contract, err := bindBookingtoken(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &BookingtokenFilterer{contract: contract}, nil
}

// bindBookingtoken binds a generic wrapper to an already deployed contract.
func bindBookingtoken(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := BookingtokenMetaData.GetAbi()
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, *parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Bookingtoken *BookingtokenRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _Bookingtoken.Contract.BookingtokenCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Bookingtoken *BookingtokenRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Bookingtoken.Contract.BookingtokenTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Bookingtoken *BookingtokenRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Bookingtoken.Contract.BookingtokenTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Bookingtoken *BookingtokenCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _Bookingtoken.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Bookingtoken *BookingtokenTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Bookingtoken.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Bookingtoken *BookingtokenTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Bookingtoken.Contract.contract.Transact(opts, method, params...)
}

// DEFAULTADMINROLE is a free data retrieval call binding the contract method 0xa217fddf.
//
// Solidity: function DEFAULT_ADMIN_ROLE() view returns(bytes32)
func (_Bookingtoken *BookingtokenCaller) DEFAULTADMINROLE(opts *bind.CallOpts) ([32]byte, error) {
	var out []interface{}
	err := _Bookingtoken.contract.Call(opts, &out, "DEFAULT_ADMIN_ROLE")

	if err != nil {
		return *new([32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)

	return out0, err

}

// DEFAULTADMINROLE is a free data retrieval call binding the contract method 0xa217fddf.
//
// Solidity: function DEFAULT_ADMIN_ROLE() view returns(bytes32)
func (_Bookingtoken *BookingtokenSession) DEFAULTADMINROLE() ([32]byte, error) {
	return _Bookingtoken.Contract.DEFAULTADMINROLE(&_Bookingtoken.CallOpts)
}

// DEFAULTADMINROLE is a free data retrieval call binding the contract method 0xa217fddf.
//
// Solidity: function DEFAULT_ADMIN_ROLE() view returns(bytes32)
func (_Bookingtoken *BookingtokenCallerSession) DEFAULTADMINROLE() ([32]byte, error) {
	return _Bookingtoken.Contract.DEFAULTADMINROLE(&_Bookingtoken.CallOpts)
}

// MINEXPIRATIONADMINROLE is a free data retrieval call binding the contract method 0x2edf5e2c.
//
// Solidity: function MIN_EXPIRATION_ADMIN_ROLE() view returns(bytes32)
func (_Bookingtoken *BookingtokenCaller) MINEXPIRATIONADMINROLE(opts *bind.CallOpts) ([32]byte, error) {
	var out []interface{}
	err := _Bookingtoken.contract.Call(opts, &out, "MIN_EXPIRATION_ADMIN_ROLE")

	if err != nil {
		return *new([32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)

	return out0, err

}

// MINEXPIRATIONADMINROLE is a free data retrieval call binding the contract method 0x2edf5e2c.
//
// Solidity: function MIN_EXPIRATION_ADMIN_ROLE() view returns(bytes32)
func (_Bookingtoken *BookingtokenSession) MINEXPIRATIONADMINROLE() ([32]byte, error) {
	return _Bookingtoken.Contract.MINEXPIRATIONADMINROLE(&_Bookingtoken.CallOpts)
}

// MINEXPIRATIONADMINROLE is a free data retrieval call binding the contract method 0x2edf5e2c.
//
// Solidity: function MIN_EXPIRATION_ADMIN_ROLE() view returns(bytes32)
func (_Bookingtoken *BookingtokenCallerSession) MINEXPIRATIONADMINROLE() ([32]byte, error) {
	return _Bookingtoken.Contract.MINEXPIRATIONADMINROLE(&_Bookingtoken.CallOpts)
}

// NATIVEPAYMENT is a free data retrieval call binding the contract method 0xa0f07c74.
//
// Solidity: function NATIVE_PAYMENT() view returns(address)
func (_Bookingtoken *BookingtokenCaller) NATIVEPAYMENT(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _Bookingtoken.contract.Call(opts, &out, "NATIVE_PAYMENT")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// NATIVEPAYMENT is a free data retrieval call binding the contract method 0xa0f07c74.
//
// Solidity: function NATIVE_PAYMENT() view returns(address)
func (_Bookingtoken *BookingtokenSession) NATIVEPAYMENT() (common.Address, error) {
	return _Bookingtoken.Contract.NATIVEPAYMENT(&_Bookingtoken.CallOpts)
}

// NATIVEPAYMENT is a free data retrieval call binding the contract method 0xa0f07c74.
//
// Solidity: function NATIVE_PAYMENT() view returns(address)
func (_Bookingtoken *BookingtokenCallerSession) NATIVEPAYMENT() (common.Address, error) {
	return _Bookingtoken.Contract.NATIVEPAYMENT(&_Bookingtoken.CallOpts)
}

// OFFCHAINPAYMENT is a free data retrieval call binding the contract method 0xbfb26c06.
//
// Solidity: function OFFCHAIN_PAYMENT() view returns(address)
func (_Bookingtoken *BookingtokenCaller) OFFCHAINPAYMENT(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _Bookingtoken.contract.Call(opts, &out, "OFFCHAIN_PAYMENT")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// OFFCHAINPAYMENT is a free data retrieval call binding the contract method 0xbfb26c06.
//
// Solidity: function OFFCHAIN_PAYMENT() view returns(address)
func (_Bookingtoken *BookingtokenSession) OFFCHAINPAYMENT() (common.Address, error) {
	return _Bookingtoken.Contract.OFFCHAINPAYMENT(&_Bookingtoken.CallOpts)
}

// OFFCHAINPAYMENT is a free data retrieval call binding the contract method 0xbfb26c06.
//
// Solidity: function OFFCHAIN_PAYMENT() view returns(address)
func (_Bookingtoken *BookingtokenCallerSession) OFFCHAINPAYMENT() (common.Address, error) {
	return _Bookingtoken.Contract.OFFCHAINPAYMENT(&_Bookingtoken.CallOpts)
}

// PAUSERROLE is a free data retrieval call binding the contract method 0xe63ab1e9.
//
// Solidity: function PAUSER_ROLE() view returns(bytes32)
func (_Bookingtoken *BookingtokenCaller) PAUSERROLE(opts *bind.CallOpts) ([32]byte, error) {
	var out []interface{}
	err := _Bookingtoken.contract.Call(opts, &out, "PAUSER_ROLE")

	if err != nil {
		return *new([32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)

	return out0, err

}

// PAUSERROLE is a free data retrieval call binding the contract method 0xe63ab1e9.
//
// Solidity: function PAUSER_ROLE() view returns(bytes32)
func (_Bookingtoken *BookingtokenSession) PAUSERROLE() ([32]byte, error) {
	return _Bookingtoken.Contract.PAUSERROLE(&_Bookingtoken.CallOpts)
}

// PAUSERROLE is a free data retrieval call binding the contract method 0xe63ab1e9.
//
// Solidity: function PAUSER_ROLE() view returns(bytes32)
func (_Bookingtoken *BookingtokenCallerSession) PAUSERROLE() ([32]byte, error) {
	return _Bookingtoken.Contract.PAUSERROLE(&_Bookingtoken.CallOpts)
}

// UPGRADERROLE is a free data retrieval call binding the contract method 0xf72c0d8b.
//
// Solidity: function UPGRADER_ROLE() view returns(bytes32)
func (_Bookingtoken *BookingtokenCaller) UPGRADERROLE(opts *bind.CallOpts) ([32]byte, error) {
	var out []interface{}
	err := _Bookingtoken.contract.Call(opts, &out, "UPGRADER_ROLE")

	if err != nil {
		return *new([32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)

	return out0, err

}

// UPGRADERROLE is a free data retrieval call binding the contract method 0xf72c0d8b.
//
// Solidity: function UPGRADER_ROLE() view returns(bytes32)
func (_Bookingtoken *BookingtokenSession) UPGRADERROLE() ([32]byte, error) {
	return _Bookingtoken.Contract.UPGRADERROLE(&_Bookingtoken.CallOpts)
}

// UPGRADERROLE is a free data retrieval call binding the contract method 0xf72c0d8b.
//
// Solidity: function UPGRADER_ROLE() view returns(bytes32)
func (_Bookingtoken *BookingtokenCallerSession) UPGRADERROLE() ([32]byte, error) {
	return _Bookingtoken.Contract.UPGRADERROLE(&_Bookingtoken.CallOpts)
}

// UPGRADEINTERFACEVERSION is a free data retrieval call binding the contract method 0xad3cb1cc.
//
// Solidity: function UPGRADE_INTERFACE_VERSION() view returns(string)
func (_Bookingtoken *BookingtokenCaller) UPGRADEINTERFACEVERSION(opts *bind.CallOpts) (string, error) {
	var out []interface{}
	err := _Bookingtoken.contract.Call(opts, &out, "UPGRADE_INTERFACE_VERSION")

	if err != nil {
		return *new(string), err
	}

	out0 := *abi.ConvertType(out[0], new(string)).(*string)

	return out0, err

}

// UPGRADEINTERFACEVERSION is a free data retrieval call binding the contract method 0xad3cb1cc.
//
// Solidity: function UPGRADE_INTERFACE_VERSION() view returns(string)
func (_Bookingtoken *BookingtokenSession) UPGRADEINTERFACEVERSION() (string, error) {
	return _Bookingtoken.Contract.UPGRADEINTERFACEVERSION(&_Bookingtoken.CallOpts)
}

// UPGRADEINTERFACEVERSION is a free data retrieval call binding the contract method 0xad3cb1cc.
//
// Solidity: function UPGRADE_INTERFACE_VERSION() view returns(string)
func (_Bookingtoken *BookingtokenCallerSession) UPGRADEINTERFACEVERSION() (string, error) {
	return _Bookingtoken.Contract.UPGRADEINTERFACEVERSION(&_Bookingtoken.CallOpts)
}

// BalanceOf is a free data retrieval call binding the contract method 0x70a08231.
//
// Solidity: function balanceOf(address owner) view returns(uint256)
func (_Bookingtoken *BookingtokenCaller) BalanceOf(opts *bind.CallOpts, owner common.Address) (*big.Int, error) {
	var out []interface{}
	err := _Bookingtoken.contract.Call(opts, &out, "balanceOf", owner)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// BalanceOf is a free data retrieval call binding the contract method 0x70a08231.
//
// Solidity: function balanceOf(address owner) view returns(uint256)
func (_Bookingtoken *BookingtokenSession) BalanceOf(owner common.Address) (*big.Int, error) {
	return _Bookingtoken.Contract.BalanceOf(&_Bookingtoken.CallOpts, owner)
}

// BalanceOf is a free data retrieval call binding the contract method 0x70a08231.
//
// Solidity: function balanceOf(address owner) view returns(uint256)
func (_Bookingtoken *BookingtokenCallerSession) BalanceOf(owner common.Address) (*big.Int, error) {
	return _Bookingtoken.Contract.BalanceOf(&_Bookingtoken.CallOpts, owner)
}

// GetApproved is a free data retrieval call binding the contract method 0x081812fc.
//
// Solidity: function getApproved(uint256 tokenId) view returns(address)
func (_Bookingtoken *BookingtokenCaller) GetApproved(opts *bind.CallOpts, tokenId *big.Int) (common.Address, error) {
	var out []interface{}
	err := _Bookingtoken.contract.Call(opts, &out, "getApproved", tokenId)

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// GetApproved is a free data retrieval call binding the contract method 0x081812fc.
//
// Solidity: function getApproved(uint256 tokenId) view returns(address)
func (_Bookingtoken *BookingtokenSession) GetApproved(tokenId *big.Int) (common.Address, error) {
	return _Bookingtoken.Contract.GetApproved(&_Bookingtoken.CallOpts, tokenId)
}

// GetApproved is a free data retrieval call binding the contract method 0x081812fc.
//
// Solidity: function getApproved(uint256 tokenId) view returns(address)
func (_Bookingtoken *BookingtokenCallerSession) GetApproved(tokenId *big.Int) (common.Address, error) {
	return _Bookingtoken.Contract.GetApproved(&_Bookingtoken.CallOpts, tokenId)
}

// GetBookingStatus is a free data retrieval call binding the contract method 0x3c15b31c.
//
// Solidity: function getBookingStatus(uint256 tokenId) view returns(uint8)
func (_Bookingtoken *BookingtokenCaller) GetBookingStatus(opts *bind.CallOpts, tokenId *big.Int) (uint8, error) {
	var out []interface{}
	err := _Bookingtoken.contract.Call(opts, &out, "getBookingStatus", tokenId)

	if err != nil {
		return *new(uint8), err
	}

	out0 := *abi.ConvertType(out[0], new(uint8)).(*uint8)

	return out0, err

}

// GetBookingStatus is a free data retrieval call binding the contract method 0x3c15b31c.
//
// Solidity: function getBookingStatus(uint256 tokenId) view returns(uint8)
func (_Bookingtoken *BookingtokenSession) GetBookingStatus(tokenId *big.Int) (uint8, error) {
	return _Bookingtoken.Contract.GetBookingStatus(&_Bookingtoken.CallOpts, tokenId)
}

// GetBookingStatus is a free data retrieval call binding the contract method 0x3c15b31c.
//
// Solidity: function getBookingStatus(uint256 tokenId) view returns(uint8)
func (_Bookingtoken *BookingtokenCallerSession) GetBookingStatus(tokenId *big.Int) (uint8, error) {
	return _Bookingtoken.Contract.GetBookingStatus(&_Bookingtoken.CallOpts, tokenId)
}

// GetCancellationProposal is a free data retrieval call binding the contract method 0xbb520b47.
//
// Solidity: function getCancellationProposal(uint256 tokenId) view returns(uint8, uint256 refundAmount, address initialProposer, address currentProposer, bool ownerAccepted, bool supplierAccepted, uint32 timesCountered, uint32 timesRejected)
func (_Bookingtoken *BookingtokenCaller) GetCancellationProposal(opts *bind.CallOpts, tokenId *big.Int) (uint8, *big.Int, common.Address, common.Address, bool, bool, uint32, uint32, error) {
	var out []interface{}
	err := _Bookingtoken.contract.Call(opts, &out, "getCancellationProposal", tokenId)

	if err != nil {
		return *new(uint8), *new(*big.Int), *new(common.Address), *new(common.Address), *new(bool), *new(bool), *new(uint32), *new(uint32), err
	}

	out0 := *abi.ConvertType(out[0], new(uint8)).(*uint8)
	out1 := *abi.ConvertType(out[1], new(*big.Int)).(**big.Int)
	out2 := *abi.ConvertType(out[2], new(common.Address)).(*common.Address)
	out3 := *abi.ConvertType(out[3], new(common.Address)).(*common.Address)
	out4 := *abi.ConvertType(out[4], new(bool)).(*bool)
	out5 := *abi.ConvertType(out[5], new(bool)).(*bool)
	out6 := *abi.ConvertType(out[6], new(uint32)).(*uint32)
	out7 := *abi.ConvertType(out[7], new(uint32)).(*uint32)

	return out0, out1, out2, out3, out4, out5, out6, out7, err

}

// GetCancellationProposal is a free data retrieval call binding the contract method 0xbb520b47.
//
// Solidity: function getCancellationProposal(uint256 tokenId) view returns(uint8, uint256 refundAmount, address initialProposer, address currentProposer, bool ownerAccepted, bool supplierAccepted, uint32 timesCountered, uint32 timesRejected)
func (_Bookingtoken *BookingtokenSession) GetCancellationProposal(tokenId *big.Int) (uint8, *big.Int, common.Address, common.Address, bool, bool, uint32, uint32, error) {
	return _Bookingtoken.Contract.GetCancellationProposal(&_Bookingtoken.CallOpts, tokenId)
}

// GetCancellationProposal is a free data retrieval call binding the contract method 0xbb520b47.
//
// Solidity: function getCancellationProposal(uint256 tokenId) view returns(uint8, uint256 refundAmount, address initialProposer, address currentProposer, bool ownerAccepted, bool supplierAccepted, uint32 timesCountered, uint32 timesRejected)
func (_Bookingtoken *BookingtokenCallerSession) GetCancellationProposal(tokenId *big.Int) (uint8, *big.Int, common.Address, common.Address, bool, bool, uint32, uint32, error) {
	return _Bookingtoken.Contract.GetCancellationProposal(&_Bookingtoken.CallOpts, tokenId)
}

// GetCancellationReasons is a free data retrieval call binding the contract method 0xa9bc55a2.
//
// Solidity: function getCancellationReasons(uint256 tokenId) view returns(uint16 cancellationReason, uint16 cancellationVersion, uint16 rejectionReason, uint16 rejectionVersion, uint16 counterReason, uint16 counterVersion, uint16 withdrawalReason, uint16 withdrawalVersion)
func (_Bookingtoken *BookingtokenCaller) GetCancellationReasons(opts *bind.CallOpts, tokenId *big.Int) (struct {
	CancellationReason  uint16
	CancellationVersion uint16
	RejectionReason     uint16
	RejectionVersion    uint16
	CounterReason       uint16
	CounterVersion      uint16
	WithdrawalReason    uint16
	WithdrawalVersion   uint16
}, error) {
	var out []interface{}
	err := _Bookingtoken.contract.Call(opts, &out, "getCancellationReasons", tokenId)

	outstruct := new(struct {
		CancellationReason  uint16
		CancellationVersion uint16
		RejectionReason     uint16
		RejectionVersion    uint16
		CounterReason       uint16
		CounterVersion      uint16
		WithdrawalReason    uint16
		WithdrawalVersion   uint16
	})
	if err != nil {
		return *outstruct, err
	}

	outstruct.CancellationReason = *abi.ConvertType(out[0], new(uint16)).(*uint16)
	outstruct.CancellationVersion = *abi.ConvertType(out[1], new(uint16)).(*uint16)
	outstruct.RejectionReason = *abi.ConvertType(out[2], new(uint16)).(*uint16)
	outstruct.RejectionVersion = *abi.ConvertType(out[3], new(uint16)).(*uint16)
	outstruct.CounterReason = *abi.ConvertType(out[4], new(uint16)).(*uint16)
	outstruct.CounterVersion = *abi.ConvertType(out[5], new(uint16)).(*uint16)
	outstruct.WithdrawalReason = *abi.ConvertType(out[6], new(uint16)).(*uint16)
	outstruct.WithdrawalVersion = *abi.ConvertType(out[7], new(uint16)).(*uint16)

	return *outstruct, err

}

// GetCancellationReasons is a free data retrieval call binding the contract method 0xa9bc55a2.
//
// Solidity: function getCancellationReasons(uint256 tokenId) view returns(uint16 cancellationReason, uint16 cancellationVersion, uint16 rejectionReason, uint16 rejectionVersion, uint16 counterReason, uint16 counterVersion, uint16 withdrawalReason, uint16 withdrawalVersion)
func (_Bookingtoken *BookingtokenSession) GetCancellationReasons(tokenId *big.Int) (struct {
	CancellationReason  uint16
	CancellationVersion uint16
	RejectionReason     uint16
	RejectionVersion    uint16
	CounterReason       uint16
	CounterVersion      uint16
	WithdrawalReason    uint16
	WithdrawalVersion   uint16
}, error) {
	return _Bookingtoken.Contract.GetCancellationReasons(&_Bookingtoken.CallOpts, tokenId)
}

// GetCancellationReasons is a free data retrieval call binding the contract method 0xa9bc55a2.
//
// Solidity: function getCancellationReasons(uint256 tokenId) view returns(uint16 cancellationReason, uint16 cancellationVersion, uint16 rejectionReason, uint16 rejectionVersion, uint16 counterReason, uint16 counterVersion, uint16 withdrawalReason, uint16 withdrawalVersion)
func (_Bookingtoken *BookingtokenCallerSession) GetCancellationReasons(tokenId *big.Int) (struct {
	CancellationReason  uint16
	CancellationVersion uint16
	RejectionReason     uint16
	RejectionVersion    uint16
	CounterReason       uint16
	CounterVersion      uint16
	WithdrawalReason    uint16
	WithdrawalVersion   uint16
}, error) {
	return _Bookingtoken.Contract.GetCancellationReasons(&_Bookingtoken.CallOpts, tokenId)
}

// GetManagerAddress is a free data retrieval call binding the contract method 0xc162d7da.
//
// Solidity: function getManagerAddress() view returns(address)
func (_Bookingtoken *BookingtokenCaller) GetManagerAddress(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _Bookingtoken.contract.Call(opts, &out, "getManagerAddress")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// GetManagerAddress is a free data retrieval call binding the contract method 0xc162d7da.
//
// Solidity: function getManagerAddress() view returns(address)
func (_Bookingtoken *BookingtokenSession) GetManagerAddress() (common.Address, error) {
	return _Bookingtoken.Contract.GetManagerAddress(&_Bookingtoken.CallOpts)
}

// GetManagerAddress is a free data retrieval call binding the contract method 0xc162d7da.
//
// Solidity: function getManagerAddress() view returns(address)
func (_Bookingtoken *BookingtokenCallerSession) GetManagerAddress() (common.Address, error) {
	return _Bookingtoken.Contract.GetManagerAddress(&_Bookingtoken.CallOpts)
}

// GetMinExpirationTimestampDiff is a free data retrieval call binding the contract method 0x0e75c1a8.
//
// Solidity: function getMinExpirationTimestampDiff() view returns(uint256)
func (_Bookingtoken *BookingtokenCaller) GetMinExpirationTimestampDiff(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _Bookingtoken.contract.Call(opts, &out, "getMinExpirationTimestampDiff")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// GetMinExpirationTimestampDiff is a free data retrieval call binding the contract method 0x0e75c1a8.
//
// Solidity: function getMinExpirationTimestampDiff() view returns(uint256)
func (_Bookingtoken *BookingtokenSession) GetMinExpirationTimestampDiff() (*big.Int, error) {
	return _Bookingtoken.Contract.GetMinExpirationTimestampDiff(&_Bookingtoken.CallOpts)
}

// GetMinExpirationTimestampDiff is a free data retrieval call binding the contract method 0x0e75c1a8.
//
// Solidity: function getMinExpirationTimestampDiff() view returns(uint256)
func (_Bookingtoken *BookingtokenCallerSession) GetMinExpirationTimestampDiff() (*big.Int, error) {
	return _Bookingtoken.Contract.GetMinExpirationTimestampDiff(&_Bookingtoken.CallOpts)
}

// GetReservationPaymentToken is a free data retrieval call binding the contract method 0xb191d092.
//
// Solidity: function getReservationPaymentToken(uint256 tokenId) view returns(address paymentToken)
func (_Bookingtoken *BookingtokenCaller) GetReservationPaymentToken(opts *bind.CallOpts, tokenId *big.Int) (common.Address, error) {
	var out []interface{}
	err := _Bookingtoken.contract.Call(opts, &out, "getReservationPaymentToken", tokenId)

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// GetReservationPaymentToken is a free data retrieval call binding the contract method 0xb191d092.
//
// Solidity: function getReservationPaymentToken(uint256 tokenId) view returns(address paymentToken)
func (_Bookingtoken *BookingtokenSession) GetReservationPaymentToken(tokenId *big.Int) (common.Address, error) {
	return _Bookingtoken.Contract.GetReservationPaymentToken(&_Bookingtoken.CallOpts, tokenId)
}

// GetReservationPaymentToken is a free data retrieval call binding the contract method 0xb191d092.
//
// Solidity: function getReservationPaymentToken(uint256 tokenId) view returns(address paymentToken)
func (_Bookingtoken *BookingtokenCallerSession) GetReservationPaymentToken(tokenId *big.Int) (common.Address, error) {
	return _Bookingtoken.Contract.GetReservationPaymentToken(&_Bookingtoken.CallOpts, tokenId)
}

// GetReservationPrice is a free data retrieval call binding the contract method 0x004fdd3c.
//
// Solidity: function getReservationPrice(uint256 tokenId) view returns(uint256 price, address paymentToken)
func (_Bookingtoken *BookingtokenCaller) GetReservationPrice(opts *bind.CallOpts, tokenId *big.Int) (struct {
	Price        *big.Int
	PaymentToken common.Address
}, error) {
	var out []interface{}
	err := _Bookingtoken.contract.Call(opts, &out, "getReservationPrice", tokenId)

	outstruct := new(struct {
		Price        *big.Int
		PaymentToken common.Address
	})
	if err != nil {
		return *outstruct, err
	}

	outstruct.Price = *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)
	outstruct.PaymentToken = *abi.ConvertType(out[1], new(common.Address)).(*common.Address)

	return *outstruct, err

}

// GetReservationPrice is a free data retrieval call binding the contract method 0x004fdd3c.
//
// Solidity: function getReservationPrice(uint256 tokenId) view returns(uint256 price, address paymentToken)
func (_Bookingtoken *BookingtokenSession) GetReservationPrice(tokenId *big.Int) (struct {
	Price        *big.Int
	PaymentToken common.Address
}, error) {
	return _Bookingtoken.Contract.GetReservationPrice(&_Bookingtoken.CallOpts, tokenId)
}

// GetReservationPrice is a free data retrieval call binding the contract method 0x004fdd3c.
//
// Solidity: function getReservationPrice(uint256 tokenId) view returns(uint256 price, address paymentToken)
func (_Bookingtoken *BookingtokenCallerSession) GetReservationPrice(tokenId *big.Int) (struct {
	Price        *big.Int
	PaymentToken common.Address
}, error) {
	return _Bookingtoken.Contract.GetReservationPrice(&_Bookingtoken.CallOpts, tokenId)
}

// GetRoleAdmin is a free data retrieval call binding the contract method 0x248a9ca3.
//
// Solidity: function getRoleAdmin(bytes32 role) view returns(bytes32)
func (_Bookingtoken *BookingtokenCaller) GetRoleAdmin(opts *bind.CallOpts, role [32]byte) ([32]byte, error) {
	var out []interface{}
	err := _Bookingtoken.contract.Call(opts, &out, "getRoleAdmin", role)

	if err != nil {
		return *new([32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)

	return out0, err

}

// GetRoleAdmin is a free data retrieval call binding the contract method 0x248a9ca3.
//
// Solidity: function getRoleAdmin(bytes32 role) view returns(bytes32)
func (_Bookingtoken *BookingtokenSession) GetRoleAdmin(role [32]byte) ([32]byte, error) {
	return _Bookingtoken.Contract.GetRoleAdmin(&_Bookingtoken.CallOpts, role)
}

// GetRoleAdmin is a free data retrieval call binding the contract method 0x248a9ca3.
//
// Solidity: function getRoleAdmin(bytes32 role) view returns(bytes32)
func (_Bookingtoken *BookingtokenCallerSession) GetRoleAdmin(role [32]byte) ([32]byte, error) {
	return _Bookingtoken.Contract.GetRoleAdmin(&_Bookingtoken.CallOpts, role)
}

// HasRole is a free data retrieval call binding the contract method 0x91d14854.
//
// Solidity: function hasRole(bytes32 role, address account) view returns(bool)
func (_Bookingtoken *BookingtokenCaller) HasRole(opts *bind.CallOpts, role [32]byte, account common.Address) (bool, error) {
	var out []interface{}
	err := _Bookingtoken.contract.Call(opts, &out, "hasRole", role, account)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// HasRole is a free data retrieval call binding the contract method 0x91d14854.
//
// Solidity: function hasRole(bytes32 role, address account) view returns(bool)
func (_Bookingtoken *BookingtokenSession) HasRole(role [32]byte, account common.Address) (bool, error) {
	return _Bookingtoken.Contract.HasRole(&_Bookingtoken.CallOpts, role, account)
}

// HasRole is a free data retrieval call binding the contract method 0x91d14854.
//
// Solidity: function hasRole(bytes32 role, address account) view returns(bool)
func (_Bookingtoken *BookingtokenCallerSession) HasRole(role [32]byte, account common.Address) (bool, error) {
	return _Bookingtoken.Contract.HasRole(&_Bookingtoken.CallOpts, role, account)
}

// IsApprovedForAll is a free data retrieval call binding the contract method 0xe985e9c5.
//
// Solidity: function isApprovedForAll(address owner, address operator) view returns(bool)
func (_Bookingtoken *BookingtokenCaller) IsApprovedForAll(opts *bind.CallOpts, owner common.Address, operator common.Address) (bool, error) {
	var out []interface{}
	err := _Bookingtoken.contract.Call(opts, &out, "isApprovedForAll", owner, operator)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// IsApprovedForAll is a free data retrieval call binding the contract method 0xe985e9c5.
//
// Solidity: function isApprovedForAll(address owner, address operator) view returns(bool)
func (_Bookingtoken *BookingtokenSession) IsApprovedForAll(owner common.Address, operator common.Address) (bool, error) {
	return _Bookingtoken.Contract.IsApprovedForAll(&_Bookingtoken.CallOpts, owner, operator)
}

// IsApprovedForAll is a free data retrieval call binding the contract method 0xe985e9c5.
//
// Solidity: function isApprovedForAll(address owner, address operator) view returns(bool)
func (_Bookingtoken *BookingtokenCallerSession) IsApprovedForAll(owner common.Address, operator common.Address) (bool, error) {
	return _Bookingtoken.Contract.IsApprovedForAll(&_Bookingtoken.CallOpts, owner, operator)
}

// IsCancellable is a free data retrieval call binding the contract method 0x2d3a6329.
//
// Solidity: function isCancellable(uint256 tokenId) view returns(bool)
func (_Bookingtoken *BookingtokenCaller) IsCancellable(opts *bind.CallOpts, tokenId *big.Int) (bool, error) {
	var out []interface{}
	err := _Bookingtoken.contract.Call(opts, &out, "isCancellable", tokenId)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// IsCancellable is a free data retrieval call binding the contract method 0x2d3a6329.
//
// Solidity: function isCancellable(uint256 tokenId) view returns(bool)
func (_Bookingtoken *BookingtokenSession) IsCancellable(tokenId *big.Int) (bool, error) {
	return _Bookingtoken.Contract.IsCancellable(&_Bookingtoken.CallOpts, tokenId)
}

// IsCancellable is a free data retrieval call binding the contract method 0x2d3a6329.
//
// Solidity: function isCancellable(uint256 tokenId) view returns(bool)
func (_Bookingtoken *BookingtokenCallerSession) IsCancellable(tokenId *big.Int) (bool, error) {
	return _Bookingtoken.Contract.IsCancellable(&_Bookingtoken.CallOpts, tokenId)
}

// IsTTMAccount is a free data retrieval call binding the contract method 0x5f7290f1.
//
// Solidity: function isTTMAccount(address account) view returns(bool)
func (_Bookingtoken *BookingtokenCaller) IsTTMAccount(opts *bind.CallOpts, account common.Address) (bool, error) {
	var out []interface{}
	err := _Bookingtoken.contract.Call(opts, &out, "isTTMAccount", account)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// IsTTMAccount is a free data retrieval call binding the contract method 0x5f7290f1.
//
// Solidity: function isTTMAccount(address account) view returns(bool)
func (_Bookingtoken *BookingtokenSession) IsTTMAccount(account common.Address) (bool, error) {
	return _Bookingtoken.Contract.IsTTMAccount(&_Bookingtoken.CallOpts, account)
}

// IsTTMAccount is a free data retrieval call binding the contract method 0x5f7290f1.
//
// Solidity: function isTTMAccount(address account) view returns(bool)
func (_Bookingtoken *BookingtokenCallerSession) IsTTMAccount(account common.Address) (bool, error) {
	return _Bookingtoken.Contract.IsTTMAccount(&_Bookingtoken.CallOpts, account)
}

// Name is a free data retrieval call binding the contract method 0x06fdde03.
//
// Solidity: function name() view returns(string)
func (_Bookingtoken *BookingtokenCaller) Name(opts *bind.CallOpts) (string, error) {
	var out []interface{}
	err := _Bookingtoken.contract.Call(opts, &out, "name")

	if err != nil {
		return *new(string), err
	}

	out0 := *abi.ConvertType(out[0], new(string)).(*string)

	return out0, err

}

// Name is a free data retrieval call binding the contract method 0x06fdde03.
//
// Solidity: function name() view returns(string)
func (_Bookingtoken *BookingtokenSession) Name() (string, error) {
	return _Bookingtoken.Contract.Name(&_Bookingtoken.CallOpts)
}

// Name is a free data retrieval call binding the contract method 0x06fdde03.
//
// Solidity: function name() view returns(string)
func (_Bookingtoken *BookingtokenCallerSession) Name() (string, error) {
	return _Bookingtoken.Contract.Name(&_Bookingtoken.CallOpts)
}

// OwnerOf is a free data retrieval call binding the contract method 0x6352211e.
//
// Solidity: function ownerOf(uint256 tokenId) view returns(address)
func (_Bookingtoken *BookingtokenCaller) OwnerOf(opts *bind.CallOpts, tokenId *big.Int) (common.Address, error) {
	var out []interface{}
	err := _Bookingtoken.contract.Call(opts, &out, "ownerOf", tokenId)

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// OwnerOf is a free data retrieval call binding the contract method 0x6352211e.
//
// Solidity: function ownerOf(uint256 tokenId) view returns(address)
func (_Bookingtoken *BookingtokenSession) OwnerOf(tokenId *big.Int) (common.Address, error) {
	return _Bookingtoken.Contract.OwnerOf(&_Bookingtoken.CallOpts, tokenId)
}

// OwnerOf is a free data retrieval call binding the contract method 0x6352211e.
//
// Solidity: function ownerOf(uint256 tokenId) view returns(address)
func (_Bookingtoken *BookingtokenCallerSession) OwnerOf(tokenId *big.Int) (common.Address, error) {
	return _Bookingtoken.Contract.OwnerOf(&_Bookingtoken.CallOpts, tokenId)
}

// Paused is a free data retrieval call binding the contract method 0x5c975abb.
//
// Solidity: function paused() view returns(bool)
func (_Bookingtoken *BookingtokenCaller) Paused(opts *bind.CallOpts) (bool, error) {
	var out []interface{}
	err := _Bookingtoken.contract.Call(opts, &out, "paused")

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// Paused is a free data retrieval call binding the contract method 0x5c975abb.
//
// Solidity: function paused() view returns(bool)
func (_Bookingtoken *BookingtokenSession) Paused() (bool, error) {
	return _Bookingtoken.Contract.Paused(&_Bookingtoken.CallOpts)
}

// Paused is a free data retrieval call binding the contract method 0x5c975abb.
//
// Solidity: function paused() view returns(bool)
func (_Bookingtoken *BookingtokenCallerSession) Paused() (bool, error) {
	return _Bookingtoken.Contract.Paused(&_Bookingtoken.CallOpts)
}

// ProxiableUUID is a free data retrieval call binding the contract method 0x52d1902d.
//
// Solidity: function proxiableUUID() view returns(bytes32)
func (_Bookingtoken *BookingtokenCaller) ProxiableUUID(opts *bind.CallOpts) ([32]byte, error) {
	var out []interface{}
	err := _Bookingtoken.contract.Call(opts, &out, "proxiableUUID")

	if err != nil {
		return *new([32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)

	return out0, err

}

// ProxiableUUID is a free data retrieval call binding the contract method 0x52d1902d.
//
// Solidity: function proxiableUUID() view returns(bytes32)
func (_Bookingtoken *BookingtokenSession) ProxiableUUID() ([32]byte, error) {
	return _Bookingtoken.Contract.ProxiableUUID(&_Bookingtoken.CallOpts)
}

// ProxiableUUID is a free data retrieval call binding the contract method 0x52d1902d.
//
// Solidity: function proxiableUUID() view returns(bytes32)
func (_Bookingtoken *BookingtokenCallerSession) ProxiableUUID() ([32]byte, error) {
	return _Bookingtoken.Contract.ProxiableUUID(&_Bookingtoken.CallOpts)
}

// SupportsInterface is a free data retrieval call binding the contract method 0x01ffc9a7.
//
// Solidity: function supportsInterface(bytes4 interfaceId) view returns(bool)
func (_Bookingtoken *BookingtokenCaller) SupportsInterface(opts *bind.CallOpts, interfaceId [4]byte) (bool, error) {
	var out []interface{}
	err := _Bookingtoken.contract.Call(opts, &out, "supportsInterface", interfaceId)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// SupportsInterface is a free data retrieval call binding the contract method 0x01ffc9a7.
//
// Solidity: function supportsInterface(bytes4 interfaceId) view returns(bool)
func (_Bookingtoken *BookingtokenSession) SupportsInterface(interfaceId [4]byte) (bool, error) {
	return _Bookingtoken.Contract.SupportsInterface(&_Bookingtoken.CallOpts, interfaceId)
}

// SupportsInterface is a free data retrieval call binding the contract method 0x01ffc9a7.
//
// Solidity: function supportsInterface(bytes4 interfaceId) view returns(bool)
func (_Bookingtoken *BookingtokenCallerSession) SupportsInterface(interfaceId [4]byte) (bool, error) {
	return _Bookingtoken.Contract.SupportsInterface(&_Bookingtoken.CallOpts, interfaceId)
}

// Symbol is a free data retrieval call binding the contract method 0x95d89b41.
//
// Solidity: function symbol() view returns(string)
func (_Bookingtoken *BookingtokenCaller) Symbol(opts *bind.CallOpts) (string, error) {
	var out []interface{}
	err := _Bookingtoken.contract.Call(opts, &out, "symbol")

	if err != nil {
		return *new(string), err
	}

	out0 := *abi.ConvertType(out[0], new(string)).(*string)

	return out0, err

}

// Symbol is a free data retrieval call binding the contract method 0x95d89b41.
//
// Solidity: function symbol() view returns(string)
func (_Bookingtoken *BookingtokenSession) Symbol() (string, error) {
	return _Bookingtoken.Contract.Symbol(&_Bookingtoken.CallOpts)
}

// Symbol is a free data retrieval call binding the contract method 0x95d89b41.
//
// Solidity: function symbol() view returns(string)
func (_Bookingtoken *BookingtokenCallerSession) Symbol() (string, error) {
	return _Bookingtoken.Contract.Symbol(&_Bookingtoken.CallOpts)
}

// TokenByIndex is a free data retrieval call binding the contract method 0x4f6ccce7.
//
// Solidity: function tokenByIndex(uint256 index) view returns(uint256)
func (_Bookingtoken *BookingtokenCaller) TokenByIndex(opts *bind.CallOpts, index *big.Int) (*big.Int, error) {
	var out []interface{}
	err := _Bookingtoken.contract.Call(opts, &out, "tokenByIndex", index)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// TokenByIndex is a free data retrieval call binding the contract method 0x4f6ccce7.
//
// Solidity: function tokenByIndex(uint256 index) view returns(uint256)
func (_Bookingtoken *BookingtokenSession) TokenByIndex(index *big.Int) (*big.Int, error) {
	return _Bookingtoken.Contract.TokenByIndex(&_Bookingtoken.CallOpts, index)
}

// TokenByIndex is a free data retrieval call binding the contract method 0x4f6ccce7.
//
// Solidity: function tokenByIndex(uint256 index) view returns(uint256)
func (_Bookingtoken *BookingtokenCallerSession) TokenByIndex(index *big.Int) (*big.Int, error) {
	return _Bookingtoken.Contract.TokenByIndex(&_Bookingtoken.CallOpts, index)
}

// TokenOfOwnerByIndex is a free data retrieval call binding the contract method 0x2f745c59.
//
// Solidity: function tokenOfOwnerByIndex(address owner, uint256 index) view returns(uint256)
func (_Bookingtoken *BookingtokenCaller) TokenOfOwnerByIndex(opts *bind.CallOpts, owner common.Address, index *big.Int) (*big.Int, error) {
	var out []interface{}
	err := _Bookingtoken.contract.Call(opts, &out, "tokenOfOwnerByIndex", owner, index)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// TokenOfOwnerByIndex is a free data retrieval call binding the contract method 0x2f745c59.
//
// Solidity: function tokenOfOwnerByIndex(address owner, uint256 index) view returns(uint256)
func (_Bookingtoken *BookingtokenSession) TokenOfOwnerByIndex(owner common.Address, index *big.Int) (*big.Int, error) {
	return _Bookingtoken.Contract.TokenOfOwnerByIndex(&_Bookingtoken.CallOpts, owner, index)
}

// TokenOfOwnerByIndex is a free data retrieval call binding the contract method 0x2f745c59.
//
// Solidity: function tokenOfOwnerByIndex(address owner, uint256 index) view returns(uint256)
func (_Bookingtoken *BookingtokenCallerSession) TokenOfOwnerByIndex(owner common.Address, index *big.Int) (*big.Int, error) {
	return _Bookingtoken.Contract.TokenOfOwnerByIndex(&_Bookingtoken.CallOpts, owner, index)
}

// TokenURI is a free data retrieval call binding the contract method 0xc87b56dd.
//
// Solidity: function tokenURI(uint256 tokenId) view returns(string)
func (_Bookingtoken *BookingtokenCaller) TokenURI(opts *bind.CallOpts, tokenId *big.Int) (string, error) {
	var out []interface{}
	err := _Bookingtoken.contract.Call(opts, &out, "tokenURI", tokenId)

	if err != nil {
		return *new(string), err
	}

	out0 := *abi.ConvertType(out[0], new(string)).(*string)

	return out0, err

}

// TokenURI is a free data retrieval call binding the contract method 0xc87b56dd.
//
// Solidity: function tokenURI(uint256 tokenId) view returns(string)
func (_Bookingtoken *BookingtokenSession) TokenURI(tokenId *big.Int) (string, error) {
	return _Bookingtoken.Contract.TokenURI(&_Bookingtoken.CallOpts, tokenId)
}

// TokenURI is a free data retrieval call binding the contract method 0xc87b56dd.
//
// Solidity: function tokenURI(uint256 tokenId) view returns(string)
func (_Bookingtoken *BookingtokenCallerSession) TokenURI(tokenId *big.Int) (string, error) {
	return _Bookingtoken.Contract.TokenURI(&_Bookingtoken.CallOpts, tokenId)
}

// TotalSupply is a free data retrieval call binding the contract method 0x18160ddd.
//
// Solidity: function totalSupply() view returns(uint256)
func (_Bookingtoken *BookingtokenCaller) TotalSupply(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _Bookingtoken.contract.Call(opts, &out, "totalSupply")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// TotalSupply is a free data retrieval call binding the contract method 0x18160ddd.
//
// Solidity: function totalSupply() view returns(uint256)
func (_Bookingtoken *BookingtokenSession) TotalSupply() (*big.Int, error) {
	return _Bookingtoken.Contract.TotalSupply(&_Bookingtoken.CallOpts)
}

// TotalSupply is a free data retrieval call binding the contract method 0x18160ddd.
//
// Solidity: function totalSupply() view returns(uint256)
func (_Bookingtoken *BookingtokenCallerSession) TotalSupply() (*big.Int, error) {
	return _Bookingtoken.Contract.TotalSupply(&_Bookingtoken.CallOpts)
}

// Version is a free data retrieval call binding the contract method 0x54fd4d50.
//
// Solidity: function version() pure returns(uint16 major, uint16 minor, uint16 patch)
func (_Bookingtoken *BookingtokenCaller) Version(opts *bind.CallOpts) (struct {
	Major uint16
	Minor uint16
	Patch uint16
}, error) {
	var out []interface{}
	err := _Bookingtoken.contract.Call(opts, &out, "version")

	outstruct := new(struct {
		Major uint16
		Minor uint16
		Patch uint16
	})
	if err != nil {
		return *outstruct, err
	}

	outstruct.Major = *abi.ConvertType(out[0], new(uint16)).(*uint16)
	outstruct.Minor = *abi.ConvertType(out[1], new(uint16)).(*uint16)
	outstruct.Patch = *abi.ConvertType(out[2], new(uint16)).(*uint16)

	return *outstruct, err

}

// Version is a free data retrieval call binding the contract method 0x54fd4d50.
//
// Solidity: function version() pure returns(uint16 major, uint16 minor, uint16 patch)
func (_Bookingtoken *BookingtokenSession) Version() (struct {
	Major uint16
	Minor uint16
	Patch uint16
}, error) {
	return _Bookingtoken.Contract.Version(&_Bookingtoken.CallOpts)
}

// Version is a free data retrieval call binding the contract method 0x54fd4d50.
//
// Solidity: function version() pure returns(uint16 major, uint16 minor, uint16 patch)
func (_Bookingtoken *BookingtokenCallerSession) Version() (struct {
	Major uint16
	Minor uint16
	Patch uint16
}, error) {
	return _Bookingtoken.Contract.Version(&_Bookingtoken.CallOpts)
}

// AcceptCancellation is a paid mutator transaction binding the contract method 0xbe667188.
//
// Solidity: function acceptCancellation(uint256 tokenId, uint256 refundAmount) returns()
func (_Bookingtoken *BookingtokenTransactor) AcceptCancellation(opts *bind.TransactOpts, tokenId *big.Int, refundAmount *big.Int) (*types.Transaction, error) {
	return _Bookingtoken.contract.Transact(opts, "acceptCancellation", tokenId, refundAmount)
}

// AcceptCancellation is a paid mutator transaction binding the contract method 0xbe667188.
//
// Solidity: function acceptCancellation(uint256 tokenId, uint256 refundAmount) returns()
func (_Bookingtoken *BookingtokenSession) AcceptCancellation(tokenId *big.Int, refundAmount *big.Int) (*types.Transaction, error) {
	return _Bookingtoken.Contract.AcceptCancellation(&_Bookingtoken.TransactOpts, tokenId, refundAmount)
}

// AcceptCancellation is a paid mutator transaction binding the contract method 0xbe667188.
//
// Solidity: function acceptCancellation(uint256 tokenId, uint256 refundAmount) returns()
func (_Bookingtoken *BookingtokenTransactorSession) AcceptCancellation(tokenId *big.Int, refundAmount *big.Int) (*types.Transaction, error) {
	return _Bookingtoken.Contract.AcceptCancellation(&_Bookingtoken.TransactOpts, tokenId, refundAmount)
}

// Approve is a paid mutator transaction binding the contract method 0x095ea7b3.
//
// Solidity: function approve(address to, uint256 tokenId) returns()
func (_Bookingtoken *BookingtokenTransactor) Approve(opts *bind.TransactOpts, to common.Address, tokenId *big.Int) (*types.Transaction, error) {
	return _Bookingtoken.contract.Transact(opts, "approve", to, tokenId)
}

// Approve is a paid mutator transaction binding the contract method 0x095ea7b3.
//
// Solidity: function approve(address to, uint256 tokenId) returns()
func (_Bookingtoken *BookingtokenSession) Approve(to common.Address, tokenId *big.Int) (*types.Transaction, error) {
	return _Bookingtoken.Contract.Approve(&_Bookingtoken.TransactOpts, to, tokenId)
}

// Approve is a paid mutator transaction binding the contract method 0x095ea7b3.
//
// Solidity: function approve(address to, uint256 tokenId) returns()
func (_Bookingtoken *BookingtokenTransactorSession) Approve(to common.Address, tokenId *big.Int) (*types.Transaction, error) {
	return _Bookingtoken.Contract.Approve(&_Bookingtoken.TransactOpts, to, tokenId)
}

// BuyReservedToken is a paid mutator transaction binding the contract method 0x96591edd.
//
// Solidity: function buyReservedToken(uint256 tokenId) payable returns()
func (_Bookingtoken *BookingtokenTransactor) BuyReservedToken(opts *bind.TransactOpts, tokenId *big.Int) (*types.Transaction, error) {
	return _Bookingtoken.contract.Transact(opts, "buyReservedToken", tokenId)
}

// BuyReservedToken is a paid mutator transaction binding the contract method 0x96591edd.
//
// Solidity: function buyReservedToken(uint256 tokenId) payable returns()
func (_Bookingtoken *BookingtokenSession) BuyReservedToken(tokenId *big.Int) (*types.Transaction, error) {
	return _Bookingtoken.Contract.BuyReservedToken(&_Bookingtoken.TransactOpts, tokenId)
}

// BuyReservedToken is a paid mutator transaction binding the contract method 0x96591edd.
//
// Solidity: function buyReservedToken(uint256 tokenId) payable returns()
func (_Bookingtoken *BookingtokenTransactorSession) BuyReservedToken(tokenId *big.Int) (*types.Transaction, error) {
	return _Bookingtoken.Contract.BuyReservedToken(&_Bookingtoken.TransactOpts, tokenId)
}

// CounterCancellation is a paid mutator transaction binding the contract method 0x74aa2048.
//
// Solidity: function counterCancellation(uint256 tokenId, uint256 refundAmount, uint16 counterReason, uint16 counterReasonVersion) returns()
func (_Bookingtoken *BookingtokenTransactor) CounterCancellation(opts *bind.TransactOpts, tokenId *big.Int, refundAmount *big.Int, counterReason uint16, counterReasonVersion uint16) (*types.Transaction, error) {
	return _Bookingtoken.contract.Transact(opts, "counterCancellation", tokenId, refundAmount, counterReason, counterReasonVersion)
}

// CounterCancellation is a paid mutator transaction binding the contract method 0x74aa2048.
//
// Solidity: function counterCancellation(uint256 tokenId, uint256 refundAmount, uint16 counterReason, uint16 counterReasonVersion) returns()
func (_Bookingtoken *BookingtokenSession) CounterCancellation(tokenId *big.Int, refundAmount *big.Int, counterReason uint16, counterReasonVersion uint16) (*types.Transaction, error) {
	return _Bookingtoken.Contract.CounterCancellation(&_Bookingtoken.TransactOpts, tokenId, refundAmount, counterReason, counterReasonVersion)
}

// CounterCancellation is a paid mutator transaction binding the contract method 0x74aa2048.
//
// Solidity: function counterCancellation(uint256 tokenId, uint256 refundAmount, uint16 counterReason, uint16 counterReasonVersion) returns()
func (_Bookingtoken *BookingtokenTransactorSession) CounterCancellation(tokenId *big.Int, refundAmount *big.Int, counterReason uint16, counterReasonVersion uint16) (*types.Transaction, error) {
	return _Bookingtoken.Contract.CounterCancellation(&_Bookingtoken.TransactOpts, tokenId, refundAmount, counterReason, counterReasonVersion)
}

// FinalizeCancellation is a paid mutator transaction binding the contract method 0x1c54f0f7.
//
// Solidity: function finalizeCancellation(uint256 tokenId, uint256 checkRefundAmount) payable returns()
func (_Bookingtoken *BookingtokenTransactor) FinalizeCancellation(opts *bind.TransactOpts, tokenId *big.Int, checkRefundAmount *big.Int) (*types.Transaction, error) {
	return _Bookingtoken.contract.Transact(opts, "finalizeCancellation", tokenId, checkRefundAmount)
}

// FinalizeCancellation is a paid mutator transaction binding the contract method 0x1c54f0f7.
//
// Solidity: function finalizeCancellation(uint256 tokenId, uint256 checkRefundAmount) payable returns()
func (_Bookingtoken *BookingtokenSession) FinalizeCancellation(tokenId *big.Int, checkRefundAmount *big.Int) (*types.Transaction, error) {
	return _Bookingtoken.Contract.FinalizeCancellation(&_Bookingtoken.TransactOpts, tokenId, checkRefundAmount)
}

// FinalizeCancellation is a paid mutator transaction binding the contract method 0x1c54f0f7.
//
// Solidity: function finalizeCancellation(uint256 tokenId, uint256 checkRefundAmount) payable returns()
func (_Bookingtoken *BookingtokenTransactorSession) FinalizeCancellation(tokenId *big.Int, checkRefundAmount *big.Int) (*types.Transaction, error) {
	return _Bookingtoken.Contract.FinalizeCancellation(&_Bookingtoken.TransactOpts, tokenId, checkRefundAmount)
}

// GrantRole is a paid mutator transaction binding the contract method 0x2f2ff15d.
//
// Solidity: function grantRole(bytes32 role, address account) returns()
func (_Bookingtoken *BookingtokenTransactor) GrantRole(opts *bind.TransactOpts, role [32]byte, account common.Address) (*types.Transaction, error) {
	return _Bookingtoken.contract.Transact(opts, "grantRole", role, account)
}

// GrantRole is a paid mutator transaction binding the contract method 0x2f2ff15d.
//
// Solidity: function grantRole(bytes32 role, address account) returns()
func (_Bookingtoken *BookingtokenSession) GrantRole(role [32]byte, account common.Address) (*types.Transaction, error) {
	return _Bookingtoken.Contract.GrantRole(&_Bookingtoken.TransactOpts, role, account)
}

// GrantRole is a paid mutator transaction binding the contract method 0x2f2ff15d.
//
// Solidity: function grantRole(bytes32 role, address account) returns()
func (_Bookingtoken *BookingtokenTransactorSession) GrantRole(role [32]byte, account common.Address) (*types.Transaction, error) {
	return _Bookingtoken.Contract.GrantRole(&_Bookingtoken.TransactOpts, role, account)
}

// Initialize is a paid mutator transaction binding the contract method 0xc0c53b8b.
//
// Solidity: function initialize(address manager, address defaultAdmin, address upgrader) returns()
func (_Bookingtoken *BookingtokenTransactor) Initialize(opts *bind.TransactOpts, manager common.Address, defaultAdmin common.Address, upgrader common.Address) (*types.Transaction, error) {
	return _Bookingtoken.contract.Transact(opts, "initialize", manager, defaultAdmin, upgrader)
}

// Initialize is a paid mutator transaction binding the contract method 0xc0c53b8b.
//
// Solidity: function initialize(address manager, address defaultAdmin, address upgrader) returns()
func (_Bookingtoken *BookingtokenSession) Initialize(manager common.Address, defaultAdmin common.Address, upgrader common.Address) (*types.Transaction, error) {
	return _Bookingtoken.Contract.Initialize(&_Bookingtoken.TransactOpts, manager, defaultAdmin, upgrader)
}

// Initialize is a paid mutator transaction binding the contract method 0xc0c53b8b.
//
// Solidity: function initialize(address manager, address defaultAdmin, address upgrader) returns()
func (_Bookingtoken *BookingtokenTransactorSession) Initialize(manager common.Address, defaultAdmin common.Address, upgrader common.Address) (*types.Transaction, error) {
	return _Bookingtoken.Contract.Initialize(&_Bookingtoken.TransactOpts, manager, defaultAdmin, upgrader)
}

// InitiateCancellation is a paid mutator transaction binding the contract method 0xf7e45f09.
//
// Solidity: function initiateCancellation(uint256 tokenId, uint256 refundAmount, uint16 cancellationReason, uint16 cancellationReasonVersion) returns()
func (_Bookingtoken *BookingtokenTransactor) InitiateCancellation(opts *bind.TransactOpts, tokenId *big.Int, refundAmount *big.Int, cancellationReason uint16, cancellationReasonVersion uint16) (*types.Transaction, error) {
	return _Bookingtoken.contract.Transact(opts, "initiateCancellation", tokenId, refundAmount, cancellationReason, cancellationReasonVersion)
}

// InitiateCancellation is a paid mutator transaction binding the contract method 0xf7e45f09.
//
// Solidity: function initiateCancellation(uint256 tokenId, uint256 refundAmount, uint16 cancellationReason, uint16 cancellationReasonVersion) returns()
func (_Bookingtoken *BookingtokenSession) InitiateCancellation(tokenId *big.Int, refundAmount *big.Int, cancellationReason uint16, cancellationReasonVersion uint16) (*types.Transaction, error) {
	return _Bookingtoken.Contract.InitiateCancellation(&_Bookingtoken.TransactOpts, tokenId, refundAmount, cancellationReason, cancellationReasonVersion)
}

// InitiateCancellation is a paid mutator transaction binding the contract method 0xf7e45f09.
//
// Solidity: function initiateCancellation(uint256 tokenId, uint256 refundAmount, uint16 cancellationReason, uint16 cancellationReasonVersion) returns()
func (_Bookingtoken *BookingtokenTransactorSession) InitiateCancellation(tokenId *big.Int, refundAmount *big.Int, cancellationReason uint16, cancellationReasonVersion uint16) (*types.Transaction, error) {
	return _Bookingtoken.Contract.InitiateCancellation(&_Bookingtoken.TransactOpts, tokenId, refundAmount, cancellationReason, cancellationReasonVersion)
}

// Pause is a paid mutator transaction binding the contract method 0x8456cb59.
//
// Solidity: function pause() returns()
func (_Bookingtoken *BookingtokenTransactor) Pause(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Bookingtoken.contract.Transact(opts, "pause")
}

// Pause is a paid mutator transaction binding the contract method 0x8456cb59.
//
// Solidity: function pause() returns()
func (_Bookingtoken *BookingtokenSession) Pause() (*types.Transaction, error) {
	return _Bookingtoken.Contract.Pause(&_Bookingtoken.TransactOpts)
}

// Pause is a paid mutator transaction binding the contract method 0x8456cb59.
//
// Solidity: function pause() returns()
func (_Bookingtoken *BookingtokenTransactorSession) Pause() (*types.Transaction, error) {
	return _Bookingtoken.Contract.Pause(&_Bookingtoken.TransactOpts)
}

// RecordExpiration is a paid mutator transaction binding the contract method 0xe5a6725c.
//
// Solidity: function recordExpiration(uint256 tokenId) returns()
func (_Bookingtoken *BookingtokenTransactor) RecordExpiration(opts *bind.TransactOpts, tokenId *big.Int) (*types.Transaction, error) {
	return _Bookingtoken.contract.Transact(opts, "recordExpiration", tokenId)
}

// RecordExpiration is a paid mutator transaction binding the contract method 0xe5a6725c.
//
// Solidity: function recordExpiration(uint256 tokenId) returns()
func (_Bookingtoken *BookingtokenSession) RecordExpiration(tokenId *big.Int) (*types.Transaction, error) {
	return _Bookingtoken.Contract.RecordExpiration(&_Bookingtoken.TransactOpts, tokenId)
}

// RecordExpiration is a paid mutator transaction binding the contract method 0xe5a6725c.
//
// Solidity: function recordExpiration(uint256 tokenId) returns()
func (_Bookingtoken *BookingtokenTransactorSession) RecordExpiration(tokenId *big.Int) (*types.Transaction, error) {
	return _Bookingtoken.Contract.RecordExpiration(&_Bookingtoken.TransactOpts, tokenId)
}

// RejectCancellation is a paid mutator transaction binding the contract method 0x74fe60e9.
//
// Solidity: function rejectCancellation(uint256 tokenId, uint16 rejectionReason, uint16 rejectionReasonVersion) returns()
func (_Bookingtoken *BookingtokenTransactor) RejectCancellation(opts *bind.TransactOpts, tokenId *big.Int, rejectionReason uint16, rejectionReasonVersion uint16) (*types.Transaction, error) {
	return _Bookingtoken.contract.Transact(opts, "rejectCancellation", tokenId, rejectionReason, rejectionReasonVersion)
}

// RejectCancellation is a paid mutator transaction binding the contract method 0x74fe60e9.
//
// Solidity: function rejectCancellation(uint256 tokenId, uint16 rejectionReason, uint16 rejectionReasonVersion) returns()
func (_Bookingtoken *BookingtokenSession) RejectCancellation(tokenId *big.Int, rejectionReason uint16, rejectionReasonVersion uint16) (*types.Transaction, error) {
	return _Bookingtoken.Contract.RejectCancellation(&_Bookingtoken.TransactOpts, tokenId, rejectionReason, rejectionReasonVersion)
}

// RejectCancellation is a paid mutator transaction binding the contract method 0x74fe60e9.
//
// Solidity: function rejectCancellation(uint256 tokenId, uint16 rejectionReason, uint16 rejectionReasonVersion) returns()
func (_Bookingtoken *BookingtokenTransactorSession) RejectCancellation(tokenId *big.Int, rejectionReason uint16, rejectionReasonVersion uint16) (*types.Transaction, error) {
	return _Bookingtoken.Contract.RejectCancellation(&_Bookingtoken.TransactOpts, tokenId, rejectionReason, rejectionReasonVersion)
}

// RenounceRole is a paid mutator transaction binding the contract method 0x36568abe.
//
// Solidity: function renounceRole(bytes32 role, address callerConfirmation) returns()
func (_Bookingtoken *BookingtokenTransactor) RenounceRole(opts *bind.TransactOpts, role [32]byte, callerConfirmation common.Address) (*types.Transaction, error) {
	return _Bookingtoken.contract.Transact(opts, "renounceRole", role, callerConfirmation)
}

// RenounceRole is a paid mutator transaction binding the contract method 0x36568abe.
//
// Solidity: function renounceRole(bytes32 role, address callerConfirmation) returns()
func (_Bookingtoken *BookingtokenSession) RenounceRole(role [32]byte, callerConfirmation common.Address) (*types.Transaction, error) {
	return _Bookingtoken.Contract.RenounceRole(&_Bookingtoken.TransactOpts, role, callerConfirmation)
}

// RenounceRole is a paid mutator transaction binding the contract method 0x36568abe.
//
// Solidity: function renounceRole(bytes32 role, address callerConfirmation) returns()
func (_Bookingtoken *BookingtokenTransactorSession) RenounceRole(role [32]byte, callerConfirmation common.Address) (*types.Transaction, error) {
	return _Bookingtoken.Contract.RenounceRole(&_Bookingtoken.TransactOpts, role, callerConfirmation)
}

// RevokeRole is a paid mutator transaction binding the contract method 0xd547741f.
//
// Solidity: function revokeRole(bytes32 role, address account) returns()
func (_Bookingtoken *BookingtokenTransactor) RevokeRole(opts *bind.TransactOpts, role [32]byte, account common.Address) (*types.Transaction, error) {
	return _Bookingtoken.contract.Transact(opts, "revokeRole", role, account)
}

// RevokeRole is a paid mutator transaction binding the contract method 0xd547741f.
//
// Solidity: function revokeRole(bytes32 role, address account) returns()
func (_Bookingtoken *BookingtokenSession) RevokeRole(role [32]byte, account common.Address) (*types.Transaction, error) {
	return _Bookingtoken.Contract.RevokeRole(&_Bookingtoken.TransactOpts, role, account)
}

// RevokeRole is a paid mutator transaction binding the contract method 0xd547741f.
//
// Solidity: function revokeRole(bytes32 role, address account) returns()
func (_Bookingtoken *BookingtokenTransactorSession) RevokeRole(role [32]byte, account common.Address) (*types.Transaction, error) {
	return _Bookingtoken.Contract.RevokeRole(&_Bookingtoken.TransactOpts, role, account)
}

// SafeMintWithReservation is a paid mutator transaction binding the contract method 0xdb2b2682.
//
// Solidity: function safeMintWithReservation(address reservedFor, string uri, uint256 expirationTimestamp, uint256 price, address paymentToken, uint256 offchainPaymentCurrency, bool cancellable) returns()
func (_Bookingtoken *BookingtokenTransactor) SafeMintWithReservation(opts *bind.TransactOpts, reservedFor common.Address, uri string, expirationTimestamp *big.Int, price *big.Int, paymentToken common.Address, offchainPaymentCurrency *big.Int, cancellable bool) (*types.Transaction, error) {
	return _Bookingtoken.contract.Transact(opts, "safeMintWithReservation", reservedFor, uri, expirationTimestamp, price, paymentToken, offchainPaymentCurrency, cancellable)
}

// SafeMintWithReservation is a paid mutator transaction binding the contract method 0xdb2b2682.
//
// Solidity: function safeMintWithReservation(address reservedFor, string uri, uint256 expirationTimestamp, uint256 price, address paymentToken, uint256 offchainPaymentCurrency, bool cancellable) returns()
func (_Bookingtoken *BookingtokenSession) SafeMintWithReservation(reservedFor common.Address, uri string, expirationTimestamp *big.Int, price *big.Int, paymentToken common.Address, offchainPaymentCurrency *big.Int, cancellable bool) (*types.Transaction, error) {
	return _Bookingtoken.Contract.SafeMintWithReservation(&_Bookingtoken.TransactOpts, reservedFor, uri, expirationTimestamp, price, paymentToken, offchainPaymentCurrency, cancellable)
}

// SafeMintWithReservation is a paid mutator transaction binding the contract method 0xdb2b2682.
//
// Solidity: function safeMintWithReservation(address reservedFor, string uri, uint256 expirationTimestamp, uint256 price, address paymentToken, uint256 offchainPaymentCurrency, bool cancellable) returns()
func (_Bookingtoken *BookingtokenTransactorSession) SafeMintWithReservation(reservedFor common.Address, uri string, expirationTimestamp *big.Int, price *big.Int, paymentToken common.Address, offchainPaymentCurrency *big.Int, cancellable bool) (*types.Transaction, error) {
	return _Bookingtoken.Contract.SafeMintWithReservation(&_Bookingtoken.TransactOpts, reservedFor, uri, expirationTimestamp, price, paymentToken, offchainPaymentCurrency, cancellable)
}

// SafeTransferFrom is a paid mutator transaction binding the contract method 0x42842e0e.
//
// Solidity: function safeTransferFrom(address from, address to, uint256 tokenId) returns()
func (_Bookingtoken *BookingtokenTransactor) SafeTransferFrom(opts *bind.TransactOpts, from common.Address, to common.Address, tokenId *big.Int) (*types.Transaction, error) {
	return _Bookingtoken.contract.Transact(opts, "safeTransferFrom", from, to, tokenId)
}

// SafeTransferFrom is a paid mutator transaction binding the contract method 0x42842e0e.
//
// Solidity: function safeTransferFrom(address from, address to, uint256 tokenId) returns()
func (_Bookingtoken *BookingtokenSession) SafeTransferFrom(from common.Address, to common.Address, tokenId *big.Int) (*types.Transaction, error) {
	return _Bookingtoken.Contract.SafeTransferFrom(&_Bookingtoken.TransactOpts, from, to, tokenId)
}

// SafeTransferFrom is a paid mutator transaction binding the contract method 0x42842e0e.
//
// Solidity: function safeTransferFrom(address from, address to, uint256 tokenId) returns()
func (_Bookingtoken *BookingtokenTransactorSession) SafeTransferFrom(from common.Address, to common.Address, tokenId *big.Int) (*types.Transaction, error) {
	return _Bookingtoken.Contract.SafeTransferFrom(&_Bookingtoken.TransactOpts, from, to, tokenId)
}

// SafeTransferFrom0 is a paid mutator transaction binding the contract method 0xb88d4fde.
//
// Solidity: function safeTransferFrom(address from, address to, uint256 tokenId, bytes data) returns()
func (_Bookingtoken *BookingtokenTransactor) SafeTransferFrom0(opts *bind.TransactOpts, from common.Address, to common.Address, tokenId *big.Int, data []byte) (*types.Transaction, error) {
	return _Bookingtoken.contract.Transact(opts, "safeTransferFrom0", from, to, tokenId, data)
}

// SafeTransferFrom0 is a paid mutator transaction binding the contract method 0xb88d4fde.
//
// Solidity: function safeTransferFrom(address from, address to, uint256 tokenId, bytes data) returns()
func (_Bookingtoken *BookingtokenSession) SafeTransferFrom0(from common.Address, to common.Address, tokenId *big.Int, data []byte) (*types.Transaction, error) {
	return _Bookingtoken.Contract.SafeTransferFrom0(&_Bookingtoken.TransactOpts, from, to, tokenId, data)
}

// SafeTransferFrom0 is a paid mutator transaction binding the contract method 0xb88d4fde.
//
// Solidity: function safeTransferFrom(address from, address to, uint256 tokenId, bytes data) returns()
func (_Bookingtoken *BookingtokenTransactorSession) SafeTransferFrom0(from common.Address, to common.Address, tokenId *big.Int, data []byte) (*types.Transaction, error) {
	return _Bookingtoken.Contract.SafeTransferFrom0(&_Bookingtoken.TransactOpts, from, to, tokenId, data)
}

// SetApprovalForAll is a paid mutator transaction binding the contract method 0xa22cb465.
//
// Solidity: function setApprovalForAll(address operator, bool approved) returns()
func (_Bookingtoken *BookingtokenTransactor) SetApprovalForAll(opts *bind.TransactOpts, operator common.Address, approved bool) (*types.Transaction, error) {
	return _Bookingtoken.contract.Transact(opts, "setApprovalForAll", operator, approved)
}

// SetApprovalForAll is a paid mutator transaction binding the contract method 0xa22cb465.
//
// Solidity: function setApprovalForAll(address operator, bool approved) returns()
func (_Bookingtoken *BookingtokenSession) SetApprovalForAll(operator common.Address, approved bool) (*types.Transaction, error) {
	return _Bookingtoken.Contract.SetApprovalForAll(&_Bookingtoken.TransactOpts, operator, approved)
}

// SetApprovalForAll is a paid mutator transaction binding the contract method 0xa22cb465.
//
// Solidity: function setApprovalForAll(address operator, bool approved) returns()
func (_Bookingtoken *BookingtokenTransactorSession) SetApprovalForAll(operator common.Address, approved bool) (*types.Transaction, error) {
	return _Bookingtoken.Contract.SetApprovalForAll(&_Bookingtoken.TransactOpts, operator, approved)
}

// SetManagerAddress is a paid mutator transaction binding the contract method 0x41431908.
//
// Solidity: function setManagerAddress(address manager) returns()
func (_Bookingtoken *BookingtokenTransactor) SetManagerAddress(opts *bind.TransactOpts, manager common.Address) (*types.Transaction, error) {
	return _Bookingtoken.contract.Transact(opts, "setManagerAddress", manager)
}

// SetManagerAddress is a paid mutator transaction binding the contract method 0x41431908.
//
// Solidity: function setManagerAddress(address manager) returns()
func (_Bookingtoken *BookingtokenSession) SetManagerAddress(manager common.Address) (*types.Transaction, error) {
	return _Bookingtoken.Contract.SetManagerAddress(&_Bookingtoken.TransactOpts, manager)
}

// SetManagerAddress is a paid mutator transaction binding the contract method 0x41431908.
//
// Solidity: function setManagerAddress(address manager) returns()
func (_Bookingtoken *BookingtokenTransactorSession) SetManagerAddress(manager common.Address) (*types.Transaction, error) {
	return _Bookingtoken.Contract.SetManagerAddress(&_Bookingtoken.TransactOpts, manager)
}

// SetMinExpirationTimestampDiff is a paid mutator transaction binding the contract method 0x516a82b8.
//
// Solidity: function setMinExpirationTimestampDiff(uint256 minExpirationTimestampDiff) returns()
func (_Bookingtoken *BookingtokenTransactor) SetMinExpirationTimestampDiff(opts *bind.TransactOpts, minExpirationTimestampDiff *big.Int) (*types.Transaction, error) {
	return _Bookingtoken.contract.Transact(opts, "setMinExpirationTimestampDiff", minExpirationTimestampDiff)
}

// SetMinExpirationTimestampDiff is a paid mutator transaction binding the contract method 0x516a82b8.
//
// Solidity: function setMinExpirationTimestampDiff(uint256 minExpirationTimestampDiff) returns()
func (_Bookingtoken *BookingtokenSession) SetMinExpirationTimestampDiff(minExpirationTimestampDiff *big.Int) (*types.Transaction, error) {
	return _Bookingtoken.Contract.SetMinExpirationTimestampDiff(&_Bookingtoken.TransactOpts, minExpirationTimestampDiff)
}

// SetMinExpirationTimestampDiff is a paid mutator transaction binding the contract method 0x516a82b8.
//
// Solidity: function setMinExpirationTimestampDiff(uint256 minExpirationTimestampDiff) returns()
func (_Bookingtoken *BookingtokenTransactorSession) SetMinExpirationTimestampDiff(minExpirationTimestampDiff *big.Int) (*types.Transaction, error) {
	return _Bookingtoken.Contract.SetMinExpirationTimestampDiff(&_Bookingtoken.TransactOpts, minExpirationTimestampDiff)
}

// TransferFrom is a paid mutator transaction binding the contract method 0x23b872dd.
//
// Solidity: function transferFrom(address from, address to, uint256 tokenId) returns()
func (_Bookingtoken *BookingtokenTransactor) TransferFrom(opts *bind.TransactOpts, from common.Address, to common.Address, tokenId *big.Int) (*types.Transaction, error) {
	return _Bookingtoken.contract.Transact(opts, "transferFrom", from, to, tokenId)
}

// TransferFrom is a paid mutator transaction binding the contract method 0x23b872dd.
//
// Solidity: function transferFrom(address from, address to, uint256 tokenId) returns()
func (_Bookingtoken *BookingtokenSession) TransferFrom(from common.Address, to common.Address, tokenId *big.Int) (*types.Transaction, error) {
	return _Bookingtoken.Contract.TransferFrom(&_Bookingtoken.TransactOpts, from, to, tokenId)
}

// TransferFrom is a paid mutator transaction binding the contract method 0x23b872dd.
//
// Solidity: function transferFrom(address from, address to, uint256 tokenId) returns()
func (_Bookingtoken *BookingtokenTransactorSession) TransferFrom(from common.Address, to common.Address, tokenId *big.Int) (*types.Transaction, error) {
	return _Bookingtoken.Contract.TransferFrom(&_Bookingtoken.TransactOpts, from, to, tokenId)
}

// Unpause is a paid mutator transaction binding the contract method 0x3f4ba83a.
//
// Solidity: function unpause() returns()
func (_Bookingtoken *BookingtokenTransactor) Unpause(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Bookingtoken.contract.Transact(opts, "unpause")
}

// Unpause is a paid mutator transaction binding the contract method 0x3f4ba83a.
//
// Solidity: function unpause() returns()
func (_Bookingtoken *BookingtokenSession) Unpause() (*types.Transaction, error) {
	return _Bookingtoken.Contract.Unpause(&_Bookingtoken.TransactOpts)
}

// Unpause is a paid mutator transaction binding the contract method 0x3f4ba83a.
//
// Solidity: function unpause() returns()
func (_Bookingtoken *BookingtokenTransactorSession) Unpause() (*types.Transaction, error) {
	return _Bookingtoken.Contract.Unpause(&_Bookingtoken.TransactOpts)
}

// UpgradeToAndCall is a paid mutator transaction binding the contract method 0x4f1ef286.
//
// Solidity: function upgradeToAndCall(address newImplementation, bytes data) payable returns()
func (_Bookingtoken *BookingtokenTransactor) UpgradeToAndCall(opts *bind.TransactOpts, newImplementation common.Address, data []byte) (*types.Transaction, error) {
	return _Bookingtoken.contract.Transact(opts, "upgradeToAndCall", newImplementation, data)
}

// UpgradeToAndCall is a paid mutator transaction binding the contract method 0x4f1ef286.
//
// Solidity: function upgradeToAndCall(address newImplementation, bytes data) payable returns()
func (_Bookingtoken *BookingtokenSession) UpgradeToAndCall(newImplementation common.Address, data []byte) (*types.Transaction, error) {
	return _Bookingtoken.Contract.UpgradeToAndCall(&_Bookingtoken.TransactOpts, newImplementation, data)
}

// UpgradeToAndCall is a paid mutator transaction binding the contract method 0x4f1ef286.
//
// Solidity: function upgradeToAndCall(address newImplementation, bytes data) payable returns()
func (_Bookingtoken *BookingtokenTransactorSession) UpgradeToAndCall(newImplementation common.Address, data []byte) (*types.Transaction, error) {
	return _Bookingtoken.Contract.UpgradeToAndCall(&_Bookingtoken.TransactOpts, newImplementation, data)
}

// WithdrawCancellation is a paid mutator transaction binding the contract method 0x2a119380.
//
// Solidity: function withdrawCancellation(uint256 tokenId, uint16 withdrawalReason, uint16 withdrawalReasonVersion) returns()
func (_Bookingtoken *BookingtokenTransactor) WithdrawCancellation(opts *bind.TransactOpts, tokenId *big.Int, withdrawalReason uint16, withdrawalReasonVersion uint16) (*types.Transaction, error) {
	return _Bookingtoken.contract.Transact(opts, "withdrawCancellation", tokenId, withdrawalReason, withdrawalReasonVersion)
}

// WithdrawCancellation is a paid mutator transaction binding the contract method 0x2a119380.
//
// Solidity: function withdrawCancellation(uint256 tokenId, uint16 withdrawalReason, uint16 withdrawalReasonVersion) returns()
func (_Bookingtoken *BookingtokenSession) WithdrawCancellation(tokenId *big.Int, withdrawalReason uint16, withdrawalReasonVersion uint16) (*types.Transaction, error) {
	return _Bookingtoken.Contract.WithdrawCancellation(&_Bookingtoken.TransactOpts, tokenId, withdrawalReason, withdrawalReasonVersion)
}

// WithdrawCancellation is a paid mutator transaction binding the contract method 0x2a119380.
//
// Solidity: function withdrawCancellation(uint256 tokenId, uint16 withdrawalReason, uint16 withdrawalReasonVersion) returns()
func (_Bookingtoken *BookingtokenTransactorSession) WithdrawCancellation(tokenId *big.Int, withdrawalReason uint16, withdrawalReasonVersion uint16) (*types.Transaction, error) {
	return _Bookingtoken.Contract.WithdrawCancellation(&_Bookingtoken.TransactOpts, tokenId, withdrawalReason, withdrawalReasonVersion)
}

// BookingtokenApprovalIterator is returned from FilterApproval and is used to iterate over the raw logs and unpacked data for Approval events raised by the Bookingtoken contract.
type BookingtokenApprovalIterator struct {
	Event *BookingtokenApproval // Event containing the contract specifics and raw log

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
func (it *BookingtokenApprovalIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(BookingtokenApproval)
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
		it.Event = new(BookingtokenApproval)
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
func (it *BookingtokenApprovalIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *BookingtokenApprovalIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// BookingtokenApproval represents a Approval event raised by the Bookingtoken contract.
type BookingtokenApproval struct {
	Owner    common.Address
	Approved common.Address
	TokenId  *big.Int
	Raw      types.Log // Blockchain specific contextual infos
}

// FilterApproval is a free log retrieval operation binding the contract event 0x8c5be1e5ebec7d5bd14f71427d1e84f3dd0314c0f7b2291e5b200ac8c7c3b925.
//
// Solidity: event Approval(address indexed owner, address indexed approved, uint256 indexed tokenId)
func (_Bookingtoken *BookingtokenFilterer) FilterApproval(opts *bind.FilterOpts, owner []common.Address, approved []common.Address, tokenId []*big.Int) (*BookingtokenApprovalIterator, error) {

	var ownerRule []interface{}
	for _, ownerItem := range owner {
		ownerRule = append(ownerRule, ownerItem)
	}
	var approvedRule []interface{}
	for _, approvedItem := range approved {
		approvedRule = append(approvedRule, approvedItem)
	}
	var tokenIdRule []interface{}
	for _, tokenIdItem := range tokenId {
		tokenIdRule = append(tokenIdRule, tokenIdItem)
	}

	logs, sub, err := _Bookingtoken.contract.FilterLogs(opts, "Approval", ownerRule, approvedRule, tokenIdRule)
	if err != nil {
		return nil, err
	}
	return &BookingtokenApprovalIterator{contract: _Bookingtoken.contract, event: "Approval", logs: logs, sub: sub}, nil
}

// WatchApproval is a free log subscription operation binding the contract event 0x8c5be1e5ebec7d5bd14f71427d1e84f3dd0314c0f7b2291e5b200ac8c7c3b925.
//
// Solidity: event Approval(address indexed owner, address indexed approved, uint256 indexed tokenId)
func (_Bookingtoken *BookingtokenFilterer) WatchApproval(opts *bind.WatchOpts, sink chan<- *BookingtokenApproval, owner []common.Address, approved []common.Address, tokenId []*big.Int) (event.Subscription, error) {

	var ownerRule []interface{}
	for _, ownerItem := range owner {
		ownerRule = append(ownerRule, ownerItem)
	}
	var approvedRule []interface{}
	for _, approvedItem := range approved {
		approvedRule = append(approvedRule, approvedItem)
	}
	var tokenIdRule []interface{}
	for _, tokenIdItem := range tokenId {
		tokenIdRule = append(tokenIdRule, tokenIdItem)
	}

	logs, sub, err := _Bookingtoken.contract.WatchLogs(opts, "Approval", ownerRule, approvedRule, tokenIdRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(BookingtokenApproval)
				if err := _Bookingtoken.contract.UnpackLog(event, "Approval", log); err != nil {
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
// Solidity: event Approval(address indexed owner, address indexed approved, uint256 indexed tokenId)
func (_Bookingtoken *BookingtokenFilterer) ParseApproval(log types.Log) (*BookingtokenApproval, error) {
	event := new(BookingtokenApproval)
	if err := _Bookingtoken.contract.UnpackLog(event, "Approval", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// BookingtokenApprovalForAllIterator is returned from FilterApprovalForAll and is used to iterate over the raw logs and unpacked data for ApprovalForAll events raised by the Bookingtoken contract.
type BookingtokenApprovalForAllIterator struct {
	Event *BookingtokenApprovalForAll // Event containing the contract specifics and raw log

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
func (it *BookingtokenApprovalForAllIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(BookingtokenApprovalForAll)
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
		it.Event = new(BookingtokenApprovalForAll)
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
func (it *BookingtokenApprovalForAllIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *BookingtokenApprovalForAllIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// BookingtokenApprovalForAll represents a ApprovalForAll event raised by the Bookingtoken contract.
type BookingtokenApprovalForAll struct {
	Owner    common.Address
	Operator common.Address
	Approved bool
	Raw      types.Log // Blockchain specific contextual infos
}

// FilterApprovalForAll is a free log retrieval operation binding the contract event 0x17307eab39ab6107e8899845ad3d59bd9653f200f220920489ca2b5937696c31.
//
// Solidity: event ApprovalForAll(address indexed owner, address indexed operator, bool approved)
func (_Bookingtoken *BookingtokenFilterer) FilterApprovalForAll(opts *bind.FilterOpts, owner []common.Address, operator []common.Address) (*BookingtokenApprovalForAllIterator, error) {

	var ownerRule []interface{}
	for _, ownerItem := range owner {
		ownerRule = append(ownerRule, ownerItem)
	}
	var operatorRule []interface{}
	for _, operatorItem := range operator {
		operatorRule = append(operatorRule, operatorItem)
	}

	logs, sub, err := _Bookingtoken.contract.FilterLogs(opts, "ApprovalForAll", ownerRule, operatorRule)
	if err != nil {
		return nil, err
	}
	return &BookingtokenApprovalForAllIterator{contract: _Bookingtoken.contract, event: "ApprovalForAll", logs: logs, sub: sub}, nil
}

// WatchApprovalForAll is a free log subscription operation binding the contract event 0x17307eab39ab6107e8899845ad3d59bd9653f200f220920489ca2b5937696c31.
//
// Solidity: event ApprovalForAll(address indexed owner, address indexed operator, bool approved)
func (_Bookingtoken *BookingtokenFilterer) WatchApprovalForAll(opts *bind.WatchOpts, sink chan<- *BookingtokenApprovalForAll, owner []common.Address, operator []common.Address) (event.Subscription, error) {

	var ownerRule []interface{}
	for _, ownerItem := range owner {
		ownerRule = append(ownerRule, ownerItem)
	}
	var operatorRule []interface{}
	for _, operatorItem := range operator {
		operatorRule = append(operatorRule, operatorItem)
	}

	logs, sub, err := _Bookingtoken.contract.WatchLogs(opts, "ApprovalForAll", ownerRule, operatorRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(BookingtokenApprovalForAll)
				if err := _Bookingtoken.contract.UnpackLog(event, "ApprovalForAll", log); err != nil {
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

// ParseApprovalForAll is a log parse operation binding the contract event 0x17307eab39ab6107e8899845ad3d59bd9653f200f220920489ca2b5937696c31.
//
// Solidity: event ApprovalForAll(address indexed owner, address indexed operator, bool approved)
func (_Bookingtoken *BookingtokenFilterer) ParseApprovalForAll(log types.Log) (*BookingtokenApprovalForAll, error) {
	event := new(BookingtokenApprovalForAll)
	if err := _Bookingtoken.contract.UnpackLog(event, "ApprovalForAll", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// BookingtokenBatchMetadataUpdateIterator is returned from FilterBatchMetadataUpdate and is used to iterate over the raw logs and unpacked data for BatchMetadataUpdate events raised by the Bookingtoken contract.
type BookingtokenBatchMetadataUpdateIterator struct {
	Event *BookingtokenBatchMetadataUpdate // Event containing the contract specifics and raw log

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
func (it *BookingtokenBatchMetadataUpdateIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(BookingtokenBatchMetadataUpdate)
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
		it.Event = new(BookingtokenBatchMetadataUpdate)
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
func (it *BookingtokenBatchMetadataUpdateIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *BookingtokenBatchMetadataUpdateIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// BookingtokenBatchMetadataUpdate represents a BatchMetadataUpdate event raised by the Bookingtoken contract.
type BookingtokenBatchMetadataUpdate struct {
	FromTokenId *big.Int
	ToTokenId   *big.Int
	Raw         types.Log // Blockchain specific contextual infos
}

// FilterBatchMetadataUpdate is a free log retrieval operation binding the contract event 0x6bd5c950a8d8df17f772f5af37cb3655737899cbf903264b9795592da439661c.
//
// Solidity: event BatchMetadataUpdate(uint256 _fromTokenId, uint256 _toTokenId)
func (_Bookingtoken *BookingtokenFilterer) FilterBatchMetadataUpdate(opts *bind.FilterOpts) (*BookingtokenBatchMetadataUpdateIterator, error) {

	logs, sub, err := _Bookingtoken.contract.FilterLogs(opts, "BatchMetadataUpdate")
	if err != nil {
		return nil, err
	}
	return &BookingtokenBatchMetadataUpdateIterator{contract: _Bookingtoken.contract, event: "BatchMetadataUpdate", logs: logs, sub: sub}, nil
}

// WatchBatchMetadataUpdate is a free log subscription operation binding the contract event 0x6bd5c950a8d8df17f772f5af37cb3655737899cbf903264b9795592da439661c.
//
// Solidity: event BatchMetadataUpdate(uint256 _fromTokenId, uint256 _toTokenId)
func (_Bookingtoken *BookingtokenFilterer) WatchBatchMetadataUpdate(opts *bind.WatchOpts, sink chan<- *BookingtokenBatchMetadataUpdate) (event.Subscription, error) {

	logs, sub, err := _Bookingtoken.contract.WatchLogs(opts, "BatchMetadataUpdate")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(BookingtokenBatchMetadataUpdate)
				if err := _Bookingtoken.contract.UnpackLog(event, "BatchMetadataUpdate", log); err != nil {
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

// ParseBatchMetadataUpdate is a log parse operation binding the contract event 0x6bd5c950a8d8df17f772f5af37cb3655737899cbf903264b9795592da439661c.
//
// Solidity: event BatchMetadataUpdate(uint256 _fromTokenId, uint256 _toTokenId)
func (_Bookingtoken *BookingtokenFilterer) ParseBatchMetadataUpdate(log types.Log) (*BookingtokenBatchMetadataUpdate, error) {
	event := new(BookingtokenBatchMetadataUpdate)
	if err := _Bookingtoken.contract.UnpackLog(event, "BatchMetadataUpdate", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// BookingtokenCancellationFinalizedIterator is returned from FilterCancellationFinalized and is used to iterate over the raw logs and unpacked data for CancellationFinalized events raised by the Bookingtoken contract.
type BookingtokenCancellationFinalizedIterator struct {
	Event *BookingtokenCancellationFinalized // Event containing the contract specifics and raw log

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
func (it *BookingtokenCancellationFinalizedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(BookingtokenCancellationFinalized)
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
		it.Event = new(BookingtokenCancellationFinalized)
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
func (it *BookingtokenCancellationFinalizedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *BookingtokenCancellationFinalizedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// BookingtokenCancellationFinalized represents a CancellationFinalized event raised by the Bookingtoken contract.
type BookingtokenCancellationFinalized struct {
	TokenId *big.Int
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterCancellationFinalized is a free log retrieval operation binding the contract event 0x17c3690813e5ff9135b87fd91848109978b23db8e471498d18886560da7f2867.
//
// Solidity: event CancellationFinalized(uint256 indexed tokenId)
func (_Bookingtoken *BookingtokenFilterer) FilterCancellationFinalized(opts *bind.FilterOpts, tokenId []*big.Int) (*BookingtokenCancellationFinalizedIterator, error) {

	var tokenIdRule []interface{}
	for _, tokenIdItem := range tokenId {
		tokenIdRule = append(tokenIdRule, tokenIdItem)
	}

	logs, sub, err := _Bookingtoken.contract.FilterLogs(opts, "CancellationFinalized", tokenIdRule)
	if err != nil {
		return nil, err
	}
	return &BookingtokenCancellationFinalizedIterator{contract: _Bookingtoken.contract, event: "CancellationFinalized", logs: logs, sub: sub}, nil
}

// WatchCancellationFinalized is a free log subscription operation binding the contract event 0x17c3690813e5ff9135b87fd91848109978b23db8e471498d18886560da7f2867.
//
// Solidity: event CancellationFinalized(uint256 indexed tokenId)
func (_Bookingtoken *BookingtokenFilterer) WatchCancellationFinalized(opts *bind.WatchOpts, sink chan<- *BookingtokenCancellationFinalized, tokenId []*big.Int) (event.Subscription, error) {

	var tokenIdRule []interface{}
	for _, tokenIdItem := range tokenId {
		tokenIdRule = append(tokenIdRule, tokenIdItem)
	}

	logs, sub, err := _Bookingtoken.contract.WatchLogs(opts, "CancellationFinalized", tokenIdRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(BookingtokenCancellationFinalized)
				if err := _Bookingtoken.contract.UnpackLog(event, "CancellationFinalized", log); err != nil {
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

// ParseCancellationFinalized is a log parse operation binding the contract event 0x17c3690813e5ff9135b87fd91848109978b23db8e471498d18886560da7f2867.
//
// Solidity: event CancellationFinalized(uint256 indexed tokenId)
func (_Bookingtoken *BookingtokenFilterer) ParseCancellationFinalized(log types.Log) (*BookingtokenCancellationFinalized, error) {
	event := new(BookingtokenCancellationFinalized)
	if err := _Bookingtoken.contract.UnpackLog(event, "CancellationFinalized", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// BookingtokenCancellationPendingIterator is returned from FilterCancellationPending and is used to iterate over the raw logs and unpacked data for CancellationPending events raised by the Bookingtoken contract.
type BookingtokenCancellationPendingIterator struct {
	Event *BookingtokenCancellationPending // Event containing the contract specifics and raw log

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
func (it *BookingtokenCancellationPendingIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(BookingtokenCancellationPending)
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
		it.Event = new(BookingtokenCancellationPending)
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
func (it *BookingtokenCancellationPendingIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *BookingtokenCancellationPendingIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// BookingtokenCancellationPending represents a CancellationPending event raised by the Bookingtoken contract.
type BookingtokenCancellationPending struct {
	TokenId          *big.Int
	InitialProposer  common.Address
	CurrentProposer  common.Address
	RefundAmount     *big.Int
	OwnerAccepted    bool
	SupplierAccepted bool
	TimesCountered   uint32
	TimesRejected    uint32
	Raw              types.Log // Blockchain specific contextual infos
}

// FilterCancellationPending is a free log retrieval operation binding the contract event 0x6aafaf6639750382d6dc665bc6a3c73d54e2076a8739a69c860a2314032eb981.
//
// Solidity: event CancellationPending(uint256 indexed tokenId, address indexed initialProposer, address indexed currentProposer, uint256 refundAmount, bool ownerAccepted, bool supplierAccepted, uint32 timesCountered, uint32 timesRejected)
func (_Bookingtoken *BookingtokenFilterer) FilterCancellationPending(opts *bind.FilterOpts, tokenId []*big.Int, initialProposer []common.Address, currentProposer []common.Address) (*BookingtokenCancellationPendingIterator, error) {

	var tokenIdRule []interface{}
	for _, tokenIdItem := range tokenId {
		tokenIdRule = append(tokenIdRule, tokenIdItem)
	}
	var initialProposerRule []interface{}
	for _, initialProposerItem := range initialProposer {
		initialProposerRule = append(initialProposerRule, initialProposerItem)
	}
	var currentProposerRule []interface{}
	for _, currentProposerItem := range currentProposer {
		currentProposerRule = append(currentProposerRule, currentProposerItem)
	}

	logs, sub, err := _Bookingtoken.contract.FilterLogs(opts, "CancellationPending", tokenIdRule, initialProposerRule, currentProposerRule)
	if err != nil {
		return nil, err
	}
	return &BookingtokenCancellationPendingIterator{contract: _Bookingtoken.contract, event: "CancellationPending", logs: logs, sub: sub}, nil
}

// WatchCancellationPending is a free log subscription operation binding the contract event 0x6aafaf6639750382d6dc665bc6a3c73d54e2076a8739a69c860a2314032eb981.
//
// Solidity: event CancellationPending(uint256 indexed tokenId, address indexed initialProposer, address indexed currentProposer, uint256 refundAmount, bool ownerAccepted, bool supplierAccepted, uint32 timesCountered, uint32 timesRejected)
func (_Bookingtoken *BookingtokenFilterer) WatchCancellationPending(opts *bind.WatchOpts, sink chan<- *BookingtokenCancellationPending, tokenId []*big.Int, initialProposer []common.Address, currentProposer []common.Address) (event.Subscription, error) {

	var tokenIdRule []interface{}
	for _, tokenIdItem := range tokenId {
		tokenIdRule = append(tokenIdRule, tokenIdItem)
	}
	var initialProposerRule []interface{}
	for _, initialProposerItem := range initialProposer {
		initialProposerRule = append(initialProposerRule, initialProposerItem)
	}
	var currentProposerRule []interface{}
	for _, currentProposerItem := range currentProposer {
		currentProposerRule = append(currentProposerRule, currentProposerItem)
	}

	logs, sub, err := _Bookingtoken.contract.WatchLogs(opts, "CancellationPending", tokenIdRule, initialProposerRule, currentProposerRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(BookingtokenCancellationPending)
				if err := _Bookingtoken.contract.UnpackLog(event, "CancellationPending", log); err != nil {
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

// ParseCancellationPending is a log parse operation binding the contract event 0x6aafaf6639750382d6dc665bc6a3c73d54e2076a8739a69c860a2314032eb981.
//
// Solidity: event CancellationPending(uint256 indexed tokenId, address indexed initialProposer, address indexed currentProposer, uint256 refundAmount, bool ownerAccepted, bool supplierAccepted, uint32 timesCountered, uint32 timesRejected)
func (_Bookingtoken *BookingtokenFilterer) ParseCancellationPending(log types.Log) (*BookingtokenCancellationPending, error) {
	event := new(BookingtokenCancellationPending)
	if err := _Bookingtoken.contract.UnpackLog(event, "CancellationPending", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// BookingtokenCancellationReasonsIterator is returned from FilterCancellationReasons and is used to iterate over the raw logs and unpacked data for CancellationReasons events raised by the Bookingtoken contract.
type BookingtokenCancellationReasonsIterator struct {
	Event *BookingtokenCancellationReasons // Event containing the contract specifics and raw log

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
func (it *BookingtokenCancellationReasonsIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(BookingtokenCancellationReasons)
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
		it.Event = new(BookingtokenCancellationReasons)
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
func (it *BookingtokenCancellationReasonsIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *BookingtokenCancellationReasonsIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// BookingtokenCancellationReasons represents a CancellationReasons event raised by the Bookingtoken contract.
type BookingtokenCancellationReasons struct {
	TokenId                   *big.Int
	CancellationReason        uint16
	CancellationReasonVersion uint16
	RejectionReason           uint16
	RejectionVersion          uint16
	CounterReason             uint16
	CounterVersion            uint16
	WithdrawalReason          uint16
	WithdrawalVersion         uint16
	Raw                       types.Log // Blockchain specific contextual infos
}

// FilterCancellationReasons is a free log retrieval operation binding the contract event 0xade2ce8b39037f5140109f53fef47c0491427b8b785ab83d35aa2931295cb790.
//
// Solidity: event CancellationReasons(uint256 indexed tokenId, uint16 cancellationReason, uint16 cancellationReasonVersion, uint16 rejectionReason, uint16 rejectionVersion, uint16 counterReason, uint16 counterVersion, uint16 withdrawalReason, uint16 withdrawalVersion)
func (_Bookingtoken *BookingtokenFilterer) FilterCancellationReasons(opts *bind.FilterOpts, tokenId []*big.Int) (*BookingtokenCancellationReasonsIterator, error) {

	var tokenIdRule []interface{}
	for _, tokenIdItem := range tokenId {
		tokenIdRule = append(tokenIdRule, tokenIdItem)
	}

	logs, sub, err := _Bookingtoken.contract.FilterLogs(opts, "CancellationReasons", tokenIdRule)
	if err != nil {
		return nil, err
	}
	return &BookingtokenCancellationReasonsIterator{contract: _Bookingtoken.contract, event: "CancellationReasons", logs: logs, sub: sub}, nil
}

// WatchCancellationReasons is a free log subscription operation binding the contract event 0xade2ce8b39037f5140109f53fef47c0491427b8b785ab83d35aa2931295cb790.
//
// Solidity: event CancellationReasons(uint256 indexed tokenId, uint16 cancellationReason, uint16 cancellationReasonVersion, uint16 rejectionReason, uint16 rejectionVersion, uint16 counterReason, uint16 counterVersion, uint16 withdrawalReason, uint16 withdrawalVersion)
func (_Bookingtoken *BookingtokenFilterer) WatchCancellationReasons(opts *bind.WatchOpts, sink chan<- *BookingtokenCancellationReasons, tokenId []*big.Int) (event.Subscription, error) {

	var tokenIdRule []interface{}
	for _, tokenIdItem := range tokenId {
		tokenIdRule = append(tokenIdRule, tokenIdItem)
	}

	logs, sub, err := _Bookingtoken.contract.WatchLogs(opts, "CancellationReasons", tokenIdRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(BookingtokenCancellationReasons)
				if err := _Bookingtoken.contract.UnpackLog(event, "CancellationReasons", log); err != nil {
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

// ParseCancellationReasons is a log parse operation binding the contract event 0xade2ce8b39037f5140109f53fef47c0491427b8b785ab83d35aa2931295cb790.
//
// Solidity: event CancellationReasons(uint256 indexed tokenId, uint16 cancellationReason, uint16 cancellationReasonVersion, uint16 rejectionReason, uint16 rejectionVersion, uint16 counterReason, uint16 counterVersion, uint16 withdrawalReason, uint16 withdrawalVersion)
func (_Bookingtoken *BookingtokenFilterer) ParseCancellationReasons(log types.Log) (*BookingtokenCancellationReasons, error) {
	event := new(BookingtokenCancellationReasons)
	if err := _Bookingtoken.contract.UnpackLog(event, "CancellationReasons", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// BookingtokenCancellationRejectedIterator is returned from FilterCancellationRejected and is used to iterate over the raw logs and unpacked data for CancellationRejected events raised by the Bookingtoken contract.
type BookingtokenCancellationRejectedIterator struct {
	Event *BookingtokenCancellationRejected // Event containing the contract specifics and raw log

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
func (it *BookingtokenCancellationRejectedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(BookingtokenCancellationRejected)
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
		it.Event = new(BookingtokenCancellationRejected)
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
func (it *BookingtokenCancellationRejectedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *BookingtokenCancellationRejectedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// BookingtokenCancellationRejected represents a CancellationRejected event raised by the Bookingtoken contract.
type BookingtokenCancellationRejected struct {
	TokenId          *big.Int
	RejectionReason  uint16
	RejectionVersion uint16
	Raw              types.Log // Blockchain specific contextual infos
}

// FilterCancellationRejected is a free log retrieval operation binding the contract event 0xab78ba855f2fdb28beb212a9b3f41a33cda034729848cd452f0cc96528c23a80.
//
// Solidity: event CancellationRejected(uint256 indexed tokenId, uint16 rejectionReason, uint16 rejectionVersion)
func (_Bookingtoken *BookingtokenFilterer) FilterCancellationRejected(opts *bind.FilterOpts, tokenId []*big.Int) (*BookingtokenCancellationRejectedIterator, error) {

	var tokenIdRule []interface{}
	for _, tokenIdItem := range tokenId {
		tokenIdRule = append(tokenIdRule, tokenIdItem)
	}

	logs, sub, err := _Bookingtoken.contract.FilterLogs(opts, "CancellationRejected", tokenIdRule)
	if err != nil {
		return nil, err
	}
	return &BookingtokenCancellationRejectedIterator{contract: _Bookingtoken.contract, event: "CancellationRejected", logs: logs, sub: sub}, nil
}

// WatchCancellationRejected is a free log subscription operation binding the contract event 0xab78ba855f2fdb28beb212a9b3f41a33cda034729848cd452f0cc96528c23a80.
//
// Solidity: event CancellationRejected(uint256 indexed tokenId, uint16 rejectionReason, uint16 rejectionVersion)
func (_Bookingtoken *BookingtokenFilterer) WatchCancellationRejected(opts *bind.WatchOpts, sink chan<- *BookingtokenCancellationRejected, tokenId []*big.Int) (event.Subscription, error) {

	var tokenIdRule []interface{}
	for _, tokenIdItem := range tokenId {
		tokenIdRule = append(tokenIdRule, tokenIdItem)
	}

	logs, sub, err := _Bookingtoken.contract.WatchLogs(opts, "CancellationRejected", tokenIdRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(BookingtokenCancellationRejected)
				if err := _Bookingtoken.contract.UnpackLog(event, "CancellationRejected", log); err != nil {
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

// ParseCancellationRejected is a log parse operation binding the contract event 0xab78ba855f2fdb28beb212a9b3f41a33cda034729848cd452f0cc96528c23a80.
//
// Solidity: event CancellationRejected(uint256 indexed tokenId, uint16 rejectionReason, uint16 rejectionVersion)
func (_Bookingtoken *BookingtokenFilterer) ParseCancellationRejected(log types.Log) (*BookingtokenCancellationRejected, error) {
	event := new(BookingtokenCancellationRejected)
	if err := _Bookingtoken.contract.UnpackLog(event, "CancellationRejected", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// BookingtokenCancellationWithdrawnIterator is returned from FilterCancellationWithdrawn and is used to iterate over the raw logs and unpacked data for CancellationWithdrawn events raised by the Bookingtoken contract.
type BookingtokenCancellationWithdrawnIterator struct {
	Event *BookingtokenCancellationWithdrawn // Event containing the contract specifics and raw log

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
func (it *BookingtokenCancellationWithdrawnIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(BookingtokenCancellationWithdrawn)
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
		it.Event = new(BookingtokenCancellationWithdrawn)
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
func (it *BookingtokenCancellationWithdrawnIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *BookingtokenCancellationWithdrawnIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// BookingtokenCancellationWithdrawn represents a CancellationWithdrawn event raised by the Bookingtoken contract.
type BookingtokenCancellationWithdrawn struct {
	TokenId           *big.Int
	WithdrawalReason  uint16
	WithdrawalVersion uint16
	Raw               types.Log // Blockchain specific contextual infos
}

// FilterCancellationWithdrawn is a free log retrieval operation binding the contract event 0x48e256ce3da490e3bbba80f056bb54ec3d7264f8ad7d152b77bf8c2eca3db5a5.
//
// Solidity: event CancellationWithdrawn(uint256 indexed tokenId, uint16 withdrawalReason, uint16 withdrawalVersion)
func (_Bookingtoken *BookingtokenFilterer) FilterCancellationWithdrawn(opts *bind.FilterOpts, tokenId []*big.Int) (*BookingtokenCancellationWithdrawnIterator, error) {

	var tokenIdRule []interface{}
	for _, tokenIdItem := range tokenId {
		tokenIdRule = append(tokenIdRule, tokenIdItem)
	}

	logs, sub, err := _Bookingtoken.contract.FilterLogs(opts, "CancellationWithdrawn", tokenIdRule)
	if err != nil {
		return nil, err
	}
	return &BookingtokenCancellationWithdrawnIterator{contract: _Bookingtoken.contract, event: "CancellationWithdrawn", logs: logs, sub: sub}, nil
}

// WatchCancellationWithdrawn is a free log subscription operation binding the contract event 0x48e256ce3da490e3bbba80f056bb54ec3d7264f8ad7d152b77bf8c2eca3db5a5.
//
// Solidity: event CancellationWithdrawn(uint256 indexed tokenId, uint16 withdrawalReason, uint16 withdrawalVersion)
func (_Bookingtoken *BookingtokenFilterer) WatchCancellationWithdrawn(opts *bind.WatchOpts, sink chan<- *BookingtokenCancellationWithdrawn, tokenId []*big.Int) (event.Subscription, error) {

	var tokenIdRule []interface{}
	for _, tokenIdItem := range tokenId {
		tokenIdRule = append(tokenIdRule, tokenIdItem)
	}

	logs, sub, err := _Bookingtoken.contract.WatchLogs(opts, "CancellationWithdrawn", tokenIdRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(BookingtokenCancellationWithdrawn)
				if err := _Bookingtoken.contract.UnpackLog(event, "CancellationWithdrawn", log); err != nil {
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

// ParseCancellationWithdrawn is a log parse operation binding the contract event 0x48e256ce3da490e3bbba80f056bb54ec3d7264f8ad7d152b77bf8c2eca3db5a5.
//
// Solidity: event CancellationWithdrawn(uint256 indexed tokenId, uint16 withdrawalReason, uint16 withdrawalVersion)
func (_Bookingtoken *BookingtokenFilterer) ParseCancellationWithdrawn(log types.Log) (*BookingtokenCancellationWithdrawn, error) {
	event := new(BookingtokenCancellationWithdrawn)
	if err := _Bookingtoken.contract.UnpackLog(event, "CancellationWithdrawn", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// BookingtokenInitializedIterator is returned from FilterInitialized and is used to iterate over the raw logs and unpacked data for Initialized events raised by the Bookingtoken contract.
type BookingtokenInitializedIterator struct {
	Event *BookingtokenInitialized // Event containing the contract specifics and raw log

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
func (it *BookingtokenInitializedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(BookingtokenInitialized)
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
		it.Event = new(BookingtokenInitialized)
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
func (it *BookingtokenInitializedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *BookingtokenInitializedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// BookingtokenInitialized represents a Initialized event raised by the Bookingtoken contract.
type BookingtokenInitialized struct {
	Version uint64
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterInitialized is a free log retrieval operation binding the contract event 0xc7f505b2f371ae2175ee4913f4499e1f2633a7b5936321eed1cdaeb6115181d2.
//
// Solidity: event Initialized(uint64 version)
func (_Bookingtoken *BookingtokenFilterer) FilterInitialized(opts *bind.FilterOpts) (*BookingtokenInitializedIterator, error) {

	logs, sub, err := _Bookingtoken.contract.FilterLogs(opts, "Initialized")
	if err != nil {
		return nil, err
	}
	return &BookingtokenInitializedIterator{contract: _Bookingtoken.contract, event: "Initialized", logs: logs, sub: sub}, nil
}

// WatchInitialized is a free log subscription operation binding the contract event 0xc7f505b2f371ae2175ee4913f4499e1f2633a7b5936321eed1cdaeb6115181d2.
//
// Solidity: event Initialized(uint64 version)
func (_Bookingtoken *BookingtokenFilterer) WatchInitialized(opts *bind.WatchOpts, sink chan<- *BookingtokenInitialized) (event.Subscription, error) {

	logs, sub, err := _Bookingtoken.contract.WatchLogs(opts, "Initialized")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(BookingtokenInitialized)
				if err := _Bookingtoken.contract.UnpackLog(event, "Initialized", log); err != nil {
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
func (_Bookingtoken *BookingtokenFilterer) ParseInitialized(log types.Log) (*BookingtokenInitialized, error) {
	event := new(BookingtokenInitialized)
	if err := _Bookingtoken.contract.UnpackLog(event, "Initialized", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// BookingtokenMetadataUpdateIterator is returned from FilterMetadataUpdate and is used to iterate over the raw logs and unpacked data for MetadataUpdate events raised by the Bookingtoken contract.
type BookingtokenMetadataUpdateIterator struct {
	Event *BookingtokenMetadataUpdate // Event containing the contract specifics and raw log

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
func (it *BookingtokenMetadataUpdateIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(BookingtokenMetadataUpdate)
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
		it.Event = new(BookingtokenMetadataUpdate)
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
func (it *BookingtokenMetadataUpdateIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *BookingtokenMetadataUpdateIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// BookingtokenMetadataUpdate represents a MetadataUpdate event raised by the Bookingtoken contract.
type BookingtokenMetadataUpdate struct {
	TokenId *big.Int
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterMetadataUpdate is a free log retrieval operation binding the contract event 0xf8e1a15aba9398e019f0b49df1a4fde98ee17ae345cb5f6b5e2c27f5033e8ce7.
//
// Solidity: event MetadataUpdate(uint256 _tokenId)
func (_Bookingtoken *BookingtokenFilterer) FilterMetadataUpdate(opts *bind.FilterOpts) (*BookingtokenMetadataUpdateIterator, error) {

	logs, sub, err := _Bookingtoken.contract.FilterLogs(opts, "MetadataUpdate")
	if err != nil {
		return nil, err
	}
	return &BookingtokenMetadataUpdateIterator{contract: _Bookingtoken.contract, event: "MetadataUpdate", logs: logs, sub: sub}, nil
}

// WatchMetadataUpdate is a free log subscription operation binding the contract event 0xf8e1a15aba9398e019f0b49df1a4fde98ee17ae345cb5f6b5e2c27f5033e8ce7.
//
// Solidity: event MetadataUpdate(uint256 _tokenId)
func (_Bookingtoken *BookingtokenFilterer) WatchMetadataUpdate(opts *bind.WatchOpts, sink chan<- *BookingtokenMetadataUpdate) (event.Subscription, error) {

	logs, sub, err := _Bookingtoken.contract.WatchLogs(opts, "MetadataUpdate")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(BookingtokenMetadataUpdate)
				if err := _Bookingtoken.contract.UnpackLog(event, "MetadataUpdate", log); err != nil {
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

// ParseMetadataUpdate is a log parse operation binding the contract event 0xf8e1a15aba9398e019f0b49df1a4fde98ee17ae345cb5f6b5e2c27f5033e8ce7.
//
// Solidity: event MetadataUpdate(uint256 _tokenId)
func (_Bookingtoken *BookingtokenFilterer) ParseMetadataUpdate(log types.Log) (*BookingtokenMetadataUpdate, error) {
	event := new(BookingtokenMetadataUpdate)
	if err := _Bookingtoken.contract.UnpackLog(event, "MetadataUpdate", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// BookingtokenPausedIterator is returned from FilterPaused and is used to iterate over the raw logs and unpacked data for Paused events raised by the Bookingtoken contract.
type BookingtokenPausedIterator struct {
	Event *BookingtokenPaused // Event containing the contract specifics and raw log

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
func (it *BookingtokenPausedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(BookingtokenPaused)
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
		it.Event = new(BookingtokenPaused)
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
func (it *BookingtokenPausedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *BookingtokenPausedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// BookingtokenPaused represents a Paused event raised by the Bookingtoken contract.
type BookingtokenPaused struct {
	Account common.Address
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterPaused is a free log retrieval operation binding the contract event 0x62e78cea01bee320cd4e420270b5ea74000d11b0c9f74754ebdbfc544b05a258.
//
// Solidity: event Paused(address account)
func (_Bookingtoken *BookingtokenFilterer) FilterPaused(opts *bind.FilterOpts) (*BookingtokenPausedIterator, error) {

	logs, sub, err := _Bookingtoken.contract.FilterLogs(opts, "Paused")
	if err != nil {
		return nil, err
	}
	return &BookingtokenPausedIterator{contract: _Bookingtoken.contract, event: "Paused", logs: logs, sub: sub}, nil
}

// WatchPaused is a free log subscription operation binding the contract event 0x62e78cea01bee320cd4e420270b5ea74000d11b0c9f74754ebdbfc544b05a258.
//
// Solidity: event Paused(address account)
func (_Bookingtoken *BookingtokenFilterer) WatchPaused(opts *bind.WatchOpts, sink chan<- *BookingtokenPaused) (event.Subscription, error) {

	logs, sub, err := _Bookingtoken.contract.WatchLogs(opts, "Paused")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(BookingtokenPaused)
				if err := _Bookingtoken.contract.UnpackLog(event, "Paused", log); err != nil {
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

// ParsePaused is a log parse operation binding the contract event 0x62e78cea01bee320cd4e420270b5ea74000d11b0c9f74754ebdbfc544b05a258.
//
// Solidity: event Paused(address account)
func (_Bookingtoken *BookingtokenFilterer) ParsePaused(log types.Log) (*BookingtokenPaused, error) {
	event := new(BookingtokenPaused)
	if err := _Bookingtoken.contract.UnpackLog(event, "Paused", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// BookingtokenRoleAdminChangedIterator is returned from FilterRoleAdminChanged and is used to iterate over the raw logs and unpacked data for RoleAdminChanged events raised by the Bookingtoken contract.
type BookingtokenRoleAdminChangedIterator struct {
	Event *BookingtokenRoleAdminChanged // Event containing the contract specifics and raw log

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
func (it *BookingtokenRoleAdminChangedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(BookingtokenRoleAdminChanged)
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
		it.Event = new(BookingtokenRoleAdminChanged)
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
func (it *BookingtokenRoleAdminChangedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *BookingtokenRoleAdminChangedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// BookingtokenRoleAdminChanged represents a RoleAdminChanged event raised by the Bookingtoken contract.
type BookingtokenRoleAdminChanged struct {
	Role              [32]byte
	PreviousAdminRole [32]byte
	NewAdminRole      [32]byte
	Raw               types.Log // Blockchain specific contextual infos
}

// FilterRoleAdminChanged is a free log retrieval operation binding the contract event 0xbd79b86ffe0ab8e8776151514217cd7cacd52c909f66475c3af44e129f0b00ff.
//
// Solidity: event RoleAdminChanged(bytes32 indexed role, bytes32 indexed previousAdminRole, bytes32 indexed newAdminRole)
func (_Bookingtoken *BookingtokenFilterer) FilterRoleAdminChanged(opts *bind.FilterOpts, role [][32]byte, previousAdminRole [][32]byte, newAdminRole [][32]byte) (*BookingtokenRoleAdminChangedIterator, error) {

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

	logs, sub, err := _Bookingtoken.contract.FilterLogs(opts, "RoleAdminChanged", roleRule, previousAdminRoleRule, newAdminRoleRule)
	if err != nil {
		return nil, err
	}
	return &BookingtokenRoleAdminChangedIterator{contract: _Bookingtoken.contract, event: "RoleAdminChanged", logs: logs, sub: sub}, nil
}

// WatchRoleAdminChanged is a free log subscription operation binding the contract event 0xbd79b86ffe0ab8e8776151514217cd7cacd52c909f66475c3af44e129f0b00ff.
//
// Solidity: event RoleAdminChanged(bytes32 indexed role, bytes32 indexed previousAdminRole, bytes32 indexed newAdminRole)
func (_Bookingtoken *BookingtokenFilterer) WatchRoleAdminChanged(opts *bind.WatchOpts, sink chan<- *BookingtokenRoleAdminChanged, role [][32]byte, previousAdminRole [][32]byte, newAdminRole [][32]byte) (event.Subscription, error) {

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

	logs, sub, err := _Bookingtoken.contract.WatchLogs(opts, "RoleAdminChanged", roleRule, previousAdminRoleRule, newAdminRoleRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(BookingtokenRoleAdminChanged)
				if err := _Bookingtoken.contract.UnpackLog(event, "RoleAdminChanged", log); err != nil {
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
func (_Bookingtoken *BookingtokenFilterer) ParseRoleAdminChanged(log types.Log) (*BookingtokenRoleAdminChanged, error) {
	event := new(BookingtokenRoleAdminChanged)
	if err := _Bookingtoken.contract.UnpackLog(event, "RoleAdminChanged", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// BookingtokenRoleGrantedIterator is returned from FilterRoleGranted and is used to iterate over the raw logs and unpacked data for RoleGranted events raised by the Bookingtoken contract.
type BookingtokenRoleGrantedIterator struct {
	Event *BookingtokenRoleGranted // Event containing the contract specifics and raw log

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
func (it *BookingtokenRoleGrantedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(BookingtokenRoleGranted)
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
		it.Event = new(BookingtokenRoleGranted)
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
func (it *BookingtokenRoleGrantedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *BookingtokenRoleGrantedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// BookingtokenRoleGranted represents a RoleGranted event raised by the Bookingtoken contract.
type BookingtokenRoleGranted struct {
	Role    [32]byte
	Account common.Address
	Sender  common.Address
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterRoleGranted is a free log retrieval operation binding the contract event 0x2f8788117e7eff1d82e926ec794901d17c78024a50270940304540a733656f0d.
//
// Solidity: event RoleGranted(bytes32 indexed role, address indexed account, address indexed sender)
func (_Bookingtoken *BookingtokenFilterer) FilterRoleGranted(opts *bind.FilterOpts, role [][32]byte, account []common.Address, sender []common.Address) (*BookingtokenRoleGrantedIterator, error) {

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

	logs, sub, err := _Bookingtoken.contract.FilterLogs(opts, "RoleGranted", roleRule, accountRule, senderRule)
	if err != nil {
		return nil, err
	}
	return &BookingtokenRoleGrantedIterator{contract: _Bookingtoken.contract, event: "RoleGranted", logs: logs, sub: sub}, nil
}

// WatchRoleGranted is a free log subscription operation binding the contract event 0x2f8788117e7eff1d82e926ec794901d17c78024a50270940304540a733656f0d.
//
// Solidity: event RoleGranted(bytes32 indexed role, address indexed account, address indexed sender)
func (_Bookingtoken *BookingtokenFilterer) WatchRoleGranted(opts *bind.WatchOpts, sink chan<- *BookingtokenRoleGranted, role [][32]byte, account []common.Address, sender []common.Address) (event.Subscription, error) {

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

	logs, sub, err := _Bookingtoken.contract.WatchLogs(opts, "RoleGranted", roleRule, accountRule, senderRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(BookingtokenRoleGranted)
				if err := _Bookingtoken.contract.UnpackLog(event, "RoleGranted", log); err != nil {
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
func (_Bookingtoken *BookingtokenFilterer) ParseRoleGranted(log types.Log) (*BookingtokenRoleGranted, error) {
	event := new(BookingtokenRoleGranted)
	if err := _Bookingtoken.contract.UnpackLog(event, "RoleGranted", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// BookingtokenRoleRevokedIterator is returned from FilterRoleRevoked and is used to iterate over the raw logs and unpacked data for RoleRevoked events raised by the Bookingtoken contract.
type BookingtokenRoleRevokedIterator struct {
	Event *BookingtokenRoleRevoked // Event containing the contract specifics and raw log

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
func (it *BookingtokenRoleRevokedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(BookingtokenRoleRevoked)
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
		it.Event = new(BookingtokenRoleRevoked)
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
func (it *BookingtokenRoleRevokedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *BookingtokenRoleRevokedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// BookingtokenRoleRevoked represents a RoleRevoked event raised by the Bookingtoken contract.
type BookingtokenRoleRevoked struct {
	Role    [32]byte
	Account common.Address
	Sender  common.Address
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterRoleRevoked is a free log retrieval operation binding the contract event 0xf6391f5c32d9c69d2a47ea670b442974b53935d1edc7fd64eb21e047a839171b.
//
// Solidity: event RoleRevoked(bytes32 indexed role, address indexed account, address indexed sender)
func (_Bookingtoken *BookingtokenFilterer) FilterRoleRevoked(opts *bind.FilterOpts, role [][32]byte, account []common.Address, sender []common.Address) (*BookingtokenRoleRevokedIterator, error) {

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

	logs, sub, err := _Bookingtoken.contract.FilterLogs(opts, "RoleRevoked", roleRule, accountRule, senderRule)
	if err != nil {
		return nil, err
	}
	return &BookingtokenRoleRevokedIterator{contract: _Bookingtoken.contract, event: "RoleRevoked", logs: logs, sub: sub}, nil
}

// WatchRoleRevoked is a free log subscription operation binding the contract event 0xf6391f5c32d9c69d2a47ea670b442974b53935d1edc7fd64eb21e047a839171b.
//
// Solidity: event RoleRevoked(bytes32 indexed role, address indexed account, address indexed sender)
func (_Bookingtoken *BookingtokenFilterer) WatchRoleRevoked(opts *bind.WatchOpts, sink chan<- *BookingtokenRoleRevoked, role [][32]byte, account []common.Address, sender []common.Address) (event.Subscription, error) {

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

	logs, sub, err := _Bookingtoken.contract.WatchLogs(opts, "RoleRevoked", roleRule, accountRule, senderRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(BookingtokenRoleRevoked)
				if err := _Bookingtoken.contract.UnpackLog(event, "RoleRevoked", log); err != nil {
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
func (_Bookingtoken *BookingtokenFilterer) ParseRoleRevoked(log types.Log) (*BookingtokenRoleRevoked, error) {
	event := new(BookingtokenRoleRevoked)
	if err := _Bookingtoken.contract.UnpackLog(event, "RoleRevoked", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// BookingtokenTokenBoughtIterator is returned from FilterTokenBought and is used to iterate over the raw logs and unpacked data for TokenBought events raised by the Bookingtoken contract.
type BookingtokenTokenBoughtIterator struct {
	Event *BookingtokenTokenBought // Event containing the contract specifics and raw log

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
func (it *BookingtokenTokenBoughtIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(BookingtokenTokenBought)
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
		it.Event = new(BookingtokenTokenBought)
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
func (it *BookingtokenTokenBoughtIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *BookingtokenTokenBoughtIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// BookingtokenTokenBought represents a TokenBought event raised by the Bookingtoken contract.
type BookingtokenTokenBought struct {
	TokenId *big.Int
	Buyer   common.Address
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterTokenBought is a free log retrieval operation binding the contract event 0xa751fb02c318279a22135a408663ae08ea45eafa950a4351c14ae543cbb95040.
//
// Solidity: event TokenBought(uint256 indexed tokenId, address indexed buyer)
func (_Bookingtoken *BookingtokenFilterer) FilterTokenBought(opts *bind.FilterOpts, tokenId []*big.Int, buyer []common.Address) (*BookingtokenTokenBoughtIterator, error) {

	var tokenIdRule []interface{}
	for _, tokenIdItem := range tokenId {
		tokenIdRule = append(tokenIdRule, tokenIdItem)
	}
	var buyerRule []interface{}
	for _, buyerItem := range buyer {
		buyerRule = append(buyerRule, buyerItem)
	}

	logs, sub, err := _Bookingtoken.contract.FilterLogs(opts, "TokenBought", tokenIdRule, buyerRule)
	if err != nil {
		return nil, err
	}
	return &BookingtokenTokenBoughtIterator{contract: _Bookingtoken.contract, event: "TokenBought", logs: logs, sub: sub}, nil
}

// WatchTokenBought is a free log subscription operation binding the contract event 0xa751fb02c318279a22135a408663ae08ea45eafa950a4351c14ae543cbb95040.
//
// Solidity: event TokenBought(uint256 indexed tokenId, address indexed buyer)
func (_Bookingtoken *BookingtokenFilterer) WatchTokenBought(opts *bind.WatchOpts, sink chan<- *BookingtokenTokenBought, tokenId []*big.Int, buyer []common.Address) (event.Subscription, error) {

	var tokenIdRule []interface{}
	for _, tokenIdItem := range tokenId {
		tokenIdRule = append(tokenIdRule, tokenIdItem)
	}
	var buyerRule []interface{}
	for _, buyerItem := range buyer {
		buyerRule = append(buyerRule, buyerItem)
	}

	logs, sub, err := _Bookingtoken.contract.WatchLogs(opts, "TokenBought", tokenIdRule, buyerRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(BookingtokenTokenBought)
				if err := _Bookingtoken.contract.UnpackLog(event, "TokenBought", log); err != nil {
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

// ParseTokenBought is a log parse operation binding the contract event 0xa751fb02c318279a22135a408663ae08ea45eafa950a4351c14ae543cbb95040.
//
// Solidity: event TokenBought(uint256 indexed tokenId, address indexed buyer)
func (_Bookingtoken *BookingtokenFilterer) ParseTokenBought(log types.Log) (*BookingtokenTokenBought, error) {
	event := new(BookingtokenTokenBought)
	if err := _Bookingtoken.contract.UnpackLog(event, "TokenBought", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// BookingtokenTokenReservationExpiredIterator is returned from FilterTokenReservationExpired and is used to iterate over the raw logs and unpacked data for TokenReservationExpired events raised by the Bookingtoken contract.
type BookingtokenTokenReservationExpiredIterator struct {
	Event *BookingtokenTokenReservationExpired // Event containing the contract specifics and raw log

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
func (it *BookingtokenTokenReservationExpiredIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(BookingtokenTokenReservationExpired)
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
		it.Event = new(BookingtokenTokenReservationExpired)
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
func (it *BookingtokenTokenReservationExpiredIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *BookingtokenTokenReservationExpiredIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// BookingtokenTokenReservationExpired represents a TokenReservationExpired event raised by the Bookingtoken contract.
type BookingtokenTokenReservationExpired struct {
	TokenId *big.Int
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterTokenReservationExpired is a free log retrieval operation binding the contract event 0xc47ab59d5c41a0220b594ece6f7e87863b07f8b33579130f3a59b31d8f7b6eed.
//
// Solidity: event TokenReservationExpired(uint256 indexed tokenId)
func (_Bookingtoken *BookingtokenFilterer) FilterTokenReservationExpired(opts *bind.FilterOpts, tokenId []*big.Int) (*BookingtokenTokenReservationExpiredIterator, error) {

	var tokenIdRule []interface{}
	for _, tokenIdItem := range tokenId {
		tokenIdRule = append(tokenIdRule, tokenIdItem)
	}

	logs, sub, err := _Bookingtoken.contract.FilterLogs(opts, "TokenReservationExpired", tokenIdRule)
	if err != nil {
		return nil, err
	}
	return &BookingtokenTokenReservationExpiredIterator{contract: _Bookingtoken.contract, event: "TokenReservationExpired", logs: logs, sub: sub}, nil
}

// WatchTokenReservationExpired is a free log subscription operation binding the contract event 0xc47ab59d5c41a0220b594ece6f7e87863b07f8b33579130f3a59b31d8f7b6eed.
//
// Solidity: event TokenReservationExpired(uint256 indexed tokenId)
func (_Bookingtoken *BookingtokenFilterer) WatchTokenReservationExpired(opts *bind.WatchOpts, sink chan<- *BookingtokenTokenReservationExpired, tokenId []*big.Int) (event.Subscription, error) {

	var tokenIdRule []interface{}
	for _, tokenIdItem := range tokenId {
		tokenIdRule = append(tokenIdRule, tokenIdItem)
	}

	logs, sub, err := _Bookingtoken.contract.WatchLogs(opts, "TokenReservationExpired", tokenIdRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(BookingtokenTokenReservationExpired)
				if err := _Bookingtoken.contract.UnpackLog(event, "TokenReservationExpired", log); err != nil {
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

// ParseTokenReservationExpired is a log parse operation binding the contract event 0xc47ab59d5c41a0220b594ece6f7e87863b07f8b33579130f3a59b31d8f7b6eed.
//
// Solidity: event TokenReservationExpired(uint256 indexed tokenId)
func (_Bookingtoken *BookingtokenFilterer) ParseTokenReservationExpired(log types.Log) (*BookingtokenTokenReservationExpired, error) {
	event := new(BookingtokenTokenReservationExpired)
	if err := _Bookingtoken.contract.UnpackLog(event, "TokenReservationExpired", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// BookingtokenTokenReservedIterator is returned from FilterTokenReserved and is used to iterate over the raw logs and unpacked data for TokenReserved events raised by the Bookingtoken contract.
type BookingtokenTokenReservedIterator struct {
	Event *BookingtokenTokenReserved // Event containing the contract specifics and raw log

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
func (it *BookingtokenTokenReservedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(BookingtokenTokenReserved)
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
		it.Event = new(BookingtokenTokenReserved)
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
func (it *BookingtokenTokenReservedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *BookingtokenTokenReservedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// BookingtokenTokenReserved represents a TokenReserved event raised by the Bookingtoken contract.
type BookingtokenTokenReserved struct {
	TokenId                 *big.Int
	ReservedFor             common.Address
	Supplier                common.Address
	ExpirationTimestamp     *big.Int
	Price                   *big.Int
	PaymentToken            common.Address
	OffchainPaymentCurrency *big.Int
	Cancellable             bool
	Raw                     types.Log // Blockchain specific contextual infos
}

// FilterTokenReserved is a free log retrieval operation binding the contract event 0x1424af4f4cb40d8a1a2d00b2324cb122ba73eac426f98b62c33ff31ca045f067.
//
// Solidity: event TokenReserved(uint256 indexed tokenId, address indexed reservedFor, address indexed supplier, uint256 expirationTimestamp, uint256 price, address paymentToken, uint256 offchainPaymentCurrency, bool cancellable)
func (_Bookingtoken *BookingtokenFilterer) FilterTokenReserved(opts *bind.FilterOpts, tokenId []*big.Int, reservedFor []common.Address, supplier []common.Address) (*BookingtokenTokenReservedIterator, error) {

	var tokenIdRule []interface{}
	for _, tokenIdItem := range tokenId {
		tokenIdRule = append(tokenIdRule, tokenIdItem)
	}
	var reservedForRule []interface{}
	for _, reservedForItem := range reservedFor {
		reservedForRule = append(reservedForRule, reservedForItem)
	}
	var supplierRule []interface{}
	for _, supplierItem := range supplier {
		supplierRule = append(supplierRule, supplierItem)
	}

	logs, sub, err := _Bookingtoken.contract.FilterLogs(opts, "TokenReserved", tokenIdRule, reservedForRule, supplierRule)
	if err != nil {
		return nil, err
	}
	return &BookingtokenTokenReservedIterator{contract: _Bookingtoken.contract, event: "TokenReserved", logs: logs, sub: sub}, nil
}

// WatchTokenReserved is a free log subscription operation binding the contract event 0x1424af4f4cb40d8a1a2d00b2324cb122ba73eac426f98b62c33ff31ca045f067.
//
// Solidity: event TokenReserved(uint256 indexed tokenId, address indexed reservedFor, address indexed supplier, uint256 expirationTimestamp, uint256 price, address paymentToken, uint256 offchainPaymentCurrency, bool cancellable)
func (_Bookingtoken *BookingtokenFilterer) WatchTokenReserved(opts *bind.WatchOpts, sink chan<- *BookingtokenTokenReserved, tokenId []*big.Int, reservedFor []common.Address, supplier []common.Address) (event.Subscription, error) {

	var tokenIdRule []interface{}
	for _, tokenIdItem := range tokenId {
		tokenIdRule = append(tokenIdRule, tokenIdItem)
	}
	var reservedForRule []interface{}
	for _, reservedForItem := range reservedFor {
		reservedForRule = append(reservedForRule, reservedForItem)
	}
	var supplierRule []interface{}
	for _, supplierItem := range supplier {
		supplierRule = append(supplierRule, supplierItem)
	}

	logs, sub, err := _Bookingtoken.contract.WatchLogs(opts, "TokenReserved", tokenIdRule, reservedForRule, supplierRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(BookingtokenTokenReserved)
				if err := _Bookingtoken.contract.UnpackLog(event, "TokenReserved", log); err != nil {
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

// ParseTokenReserved is a log parse operation binding the contract event 0x1424af4f4cb40d8a1a2d00b2324cb122ba73eac426f98b62c33ff31ca045f067.
//
// Solidity: event TokenReserved(uint256 indexed tokenId, address indexed reservedFor, address indexed supplier, uint256 expirationTimestamp, uint256 price, address paymentToken, uint256 offchainPaymentCurrency, bool cancellable)
func (_Bookingtoken *BookingtokenFilterer) ParseTokenReserved(log types.Log) (*BookingtokenTokenReserved, error) {
	event := new(BookingtokenTokenReserved)
	if err := _Bookingtoken.contract.UnpackLog(event, "TokenReserved", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// BookingtokenTransferIterator is returned from FilterTransfer and is used to iterate over the raw logs and unpacked data for Transfer events raised by the Bookingtoken contract.
type BookingtokenTransferIterator struct {
	Event *BookingtokenTransfer // Event containing the contract specifics and raw log

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
func (it *BookingtokenTransferIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(BookingtokenTransfer)
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
		it.Event = new(BookingtokenTransfer)
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
func (it *BookingtokenTransferIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *BookingtokenTransferIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// BookingtokenTransfer represents a Transfer event raised by the Bookingtoken contract.
type BookingtokenTransfer struct {
	From    common.Address
	To      common.Address
	TokenId *big.Int
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterTransfer is a free log retrieval operation binding the contract event 0xddf252ad1be2c89b69c2b068fc378daa952ba7f163c4a11628f55a4df523b3ef.
//
// Solidity: event Transfer(address indexed from, address indexed to, uint256 indexed tokenId)
func (_Bookingtoken *BookingtokenFilterer) FilterTransfer(opts *bind.FilterOpts, from []common.Address, to []common.Address, tokenId []*big.Int) (*BookingtokenTransferIterator, error) {

	var fromRule []interface{}
	for _, fromItem := range from {
		fromRule = append(fromRule, fromItem)
	}
	var toRule []interface{}
	for _, toItem := range to {
		toRule = append(toRule, toItem)
	}
	var tokenIdRule []interface{}
	for _, tokenIdItem := range tokenId {
		tokenIdRule = append(tokenIdRule, tokenIdItem)
	}

	logs, sub, err := _Bookingtoken.contract.FilterLogs(opts, "Transfer", fromRule, toRule, tokenIdRule)
	if err != nil {
		return nil, err
	}
	return &BookingtokenTransferIterator{contract: _Bookingtoken.contract, event: "Transfer", logs: logs, sub: sub}, nil
}

// WatchTransfer is a free log subscription operation binding the contract event 0xddf252ad1be2c89b69c2b068fc378daa952ba7f163c4a11628f55a4df523b3ef.
//
// Solidity: event Transfer(address indexed from, address indexed to, uint256 indexed tokenId)
func (_Bookingtoken *BookingtokenFilterer) WatchTransfer(opts *bind.WatchOpts, sink chan<- *BookingtokenTransfer, from []common.Address, to []common.Address, tokenId []*big.Int) (event.Subscription, error) {

	var fromRule []interface{}
	for _, fromItem := range from {
		fromRule = append(fromRule, fromItem)
	}
	var toRule []interface{}
	for _, toItem := range to {
		toRule = append(toRule, toItem)
	}
	var tokenIdRule []interface{}
	for _, tokenIdItem := range tokenId {
		tokenIdRule = append(tokenIdRule, tokenIdItem)
	}

	logs, sub, err := _Bookingtoken.contract.WatchLogs(opts, "Transfer", fromRule, toRule, tokenIdRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(BookingtokenTransfer)
				if err := _Bookingtoken.contract.UnpackLog(event, "Transfer", log); err != nil {
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
// Solidity: event Transfer(address indexed from, address indexed to, uint256 indexed tokenId)
func (_Bookingtoken *BookingtokenFilterer) ParseTransfer(log types.Log) (*BookingtokenTransfer, error) {
	event := new(BookingtokenTransfer)
	if err := _Bookingtoken.contract.UnpackLog(event, "Transfer", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// BookingtokenUnpausedIterator is returned from FilterUnpaused and is used to iterate over the raw logs and unpacked data for Unpaused events raised by the Bookingtoken contract.
type BookingtokenUnpausedIterator struct {
	Event *BookingtokenUnpaused // Event containing the contract specifics and raw log

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
func (it *BookingtokenUnpausedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(BookingtokenUnpaused)
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
		it.Event = new(BookingtokenUnpaused)
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
func (it *BookingtokenUnpausedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *BookingtokenUnpausedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// BookingtokenUnpaused represents a Unpaused event raised by the Bookingtoken contract.
type BookingtokenUnpaused struct {
	Account common.Address
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterUnpaused is a free log retrieval operation binding the contract event 0x5db9ee0a495bf2e6ff9c91a7834c1ba4fdd244a5e8aa4e537bd38aeae4b073aa.
//
// Solidity: event Unpaused(address account)
func (_Bookingtoken *BookingtokenFilterer) FilterUnpaused(opts *bind.FilterOpts) (*BookingtokenUnpausedIterator, error) {

	logs, sub, err := _Bookingtoken.contract.FilterLogs(opts, "Unpaused")
	if err != nil {
		return nil, err
	}
	return &BookingtokenUnpausedIterator{contract: _Bookingtoken.contract, event: "Unpaused", logs: logs, sub: sub}, nil
}

// WatchUnpaused is a free log subscription operation binding the contract event 0x5db9ee0a495bf2e6ff9c91a7834c1ba4fdd244a5e8aa4e537bd38aeae4b073aa.
//
// Solidity: event Unpaused(address account)
func (_Bookingtoken *BookingtokenFilterer) WatchUnpaused(opts *bind.WatchOpts, sink chan<- *BookingtokenUnpaused) (event.Subscription, error) {

	logs, sub, err := _Bookingtoken.contract.WatchLogs(opts, "Unpaused")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(BookingtokenUnpaused)
				if err := _Bookingtoken.contract.UnpackLog(event, "Unpaused", log); err != nil {
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

// ParseUnpaused is a log parse operation binding the contract event 0x5db9ee0a495bf2e6ff9c91a7834c1ba4fdd244a5e8aa4e537bd38aeae4b073aa.
//
// Solidity: event Unpaused(address account)
func (_Bookingtoken *BookingtokenFilterer) ParseUnpaused(log types.Log) (*BookingtokenUnpaused, error) {
	event := new(BookingtokenUnpaused)
	if err := _Bookingtoken.contract.UnpackLog(event, "Unpaused", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// BookingtokenUpgradedIterator is returned from FilterUpgraded and is used to iterate over the raw logs and unpacked data for Upgraded events raised by the Bookingtoken contract.
type BookingtokenUpgradedIterator struct {
	Event *BookingtokenUpgraded // Event containing the contract specifics and raw log

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
func (it *BookingtokenUpgradedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(BookingtokenUpgraded)
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
		it.Event = new(BookingtokenUpgraded)
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
func (it *BookingtokenUpgradedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *BookingtokenUpgradedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// BookingtokenUpgraded represents a Upgraded event raised by the Bookingtoken contract.
type BookingtokenUpgraded struct {
	Implementation common.Address
	Raw            types.Log // Blockchain specific contextual infos
}

// FilterUpgraded is a free log retrieval operation binding the contract event 0xbc7cd75a20ee27fd9adebab32041f755214dbc6bffa90cc0225b39da2e5c2d3b.
//
// Solidity: event Upgraded(address indexed implementation)
func (_Bookingtoken *BookingtokenFilterer) FilterUpgraded(opts *bind.FilterOpts, implementation []common.Address) (*BookingtokenUpgradedIterator, error) {

	var implementationRule []interface{}
	for _, implementationItem := range implementation {
		implementationRule = append(implementationRule, implementationItem)
	}

	logs, sub, err := _Bookingtoken.contract.FilterLogs(opts, "Upgraded", implementationRule)
	if err != nil {
		return nil, err
	}
	return &BookingtokenUpgradedIterator{contract: _Bookingtoken.contract, event: "Upgraded", logs: logs, sub: sub}, nil
}

// WatchUpgraded is a free log subscription operation binding the contract event 0xbc7cd75a20ee27fd9adebab32041f755214dbc6bffa90cc0225b39da2e5c2d3b.
//
// Solidity: event Upgraded(address indexed implementation)
func (_Bookingtoken *BookingtokenFilterer) WatchUpgraded(opts *bind.WatchOpts, sink chan<- *BookingtokenUpgraded, implementation []common.Address) (event.Subscription, error) {

	var implementationRule []interface{}
	for _, implementationItem := range implementation {
		implementationRule = append(implementationRule, implementationItem)
	}

	logs, sub, err := _Bookingtoken.contract.WatchLogs(opts, "Upgraded", implementationRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(BookingtokenUpgraded)
				if err := _Bookingtoken.contract.UnpackLog(event, "Upgraded", log); err != nil {
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
func (_Bookingtoken *BookingtokenFilterer) ParseUpgraded(log types.Log) (*BookingtokenUpgraded, error) {
	event := new(BookingtokenUpgraded)
	if err := _Bookingtoken.contract.UnpackLog(event, "Upgraded", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}
