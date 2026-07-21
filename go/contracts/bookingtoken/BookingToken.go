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
	ABI: "[{\"inputs\":[],\"name\":\"AccessControlBadConfirmation\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"},{\"internalType\":\"bytes32\",\"name\":\"neededRole\",\"type\":\"bytes32\"}],\"name\":\"AccessControlUnauthorizedAccount\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"target\",\"type\":\"address\"}],\"name\":\"AddressEmptyCode\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"implementation\",\"type\":\"address\"}],\"name\":\"ERC1967InvalidImplementation\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"ERC1967NonPayable\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"ERC721EnumerableForbiddenBatchMint\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"sender\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"tokenId\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"owner\",\"type\":\"address\"}],\"name\":\"ERC721IncorrectOwner\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"operator\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"tokenId\",\"type\":\"uint256\"}],\"name\":\"ERC721InsufficientApproval\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"approver\",\"type\":\"address\"}],\"name\":\"ERC721InvalidApprover\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"operator\",\"type\":\"address\"}],\"name\":\"ERC721InvalidOperator\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"owner\",\"type\":\"address\"}],\"name\":\"ERC721InvalidOwner\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"receiver\",\"type\":\"address\"}],\"name\":\"ERC721InvalidReceiver\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"sender\",\"type\":\"address\"}],\"name\":\"ERC721InvalidSender\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"tokenId\",\"type\":\"uint256\"}],\"name\":\"ERC721NonexistentToken\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"owner\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"index\",\"type\":\"uint256\"}],\"name\":\"ERC721OutOfBoundsIndex\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"EnforcedPause\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"ExpectedPause\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"expirationTimestamp\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"minExpirationTimestampDiff\",\"type\":\"uint256\"}],\"name\":\"ExpirationTimestampTooSoon\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"FailedCall\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"price\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"reservationPrice\",\"type\":\"uint256\"}],\"name\":\"IncorrectPrice\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"tokenId\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"existing\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"checked\",\"type\":\"uint256\"}],\"name\":\"IncorrectRefundAmount\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"balance\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"needed\",\"type\":\"uint256\"}],\"name\":\"InsufficientBalance\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"tokenId\",\"type\":\"uint256\"},{\"internalType\":\"enumCancellationProposalStatus\",\"name\":\"status\",\"type\":\"uint8\"}],\"name\":\"InvalidCancellationProposalStatus\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"InvalidInitialization\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"tokenId\",\"type\":\"uint256\"},{\"internalType\":\"enumBookingToken.BookingStatus\",\"name\":\"status\",\"type\":\"uint8\"}],\"name\":\"InvalidTokenStatus\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"NotInitializing\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"NotOwnerOrSupplier\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"}],\"name\":\"NotTTMAccount\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"tokenId\",\"type\":\"uint256\"}],\"name\":\"OnlyCurrentProposerCanWithdrawCancellation\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"tokenId\",\"type\":\"uint256\"}],\"name\":\"OnlySupplierCanFinalizeCancellation\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"tokenId\",\"type\":\"uint256\"}],\"name\":\"OwnerNotAcceptedCancellation\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"tokenId\",\"type\":\"uint256\"}],\"name\":\"ProposerCanNotRejectCancellation\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"ReentrancyGuardReentrantCall\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"tokenId\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"expirationTimestamp\",\"type\":\"uint256\"}],\"name\":\"ReservationExpired\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"reservedFor\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"buyer\",\"type\":\"address\"}],\"name\":\"ReservationMismatch\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"token\",\"type\":\"address\"}],\"name\":\"SafeERC20FailedOperation\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"tokenId\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"supplier\",\"type\":\"address\"}],\"name\":\"SupplierIsNotOwner\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"tokenId\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"reservedFor\",\"type\":\"address\"}],\"name\":\"TokenIsReserved\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"UUPSUnauthorizedCallContext\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"slot\",\"type\":\"bytes32\"}],\"name\":\"UUPSUnsupportedProxiableUUID\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"UnexpectedNativePayment\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"offchainPaymentCurrency\",\"type\":\"uint256\"}],\"name\":\"UnexpectedOffchainPaymentCurrency\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"ZeroAddress\",\"type\":\"error\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"owner\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"approved\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"uint256\",\"name\":\"tokenId\",\"type\":\"uint256\"}],\"name\":\"Approval\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"owner\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"operator\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"bool\",\"name\":\"approved\",\"type\":\"bool\"}],\"name\":\"ApprovalForAll\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"_fromTokenId\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"_toTokenId\",\"type\":\"uint256\"}],\"name\":\"BatchMetadataUpdate\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"uint256\",\"name\":\"tokenId\",\"type\":\"uint256\"}],\"name\":\"CancellationFinalized\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"uint256\",\"name\":\"tokenId\",\"type\":\"uint256\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"initialProposer\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"currentProposer\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"refundAmount\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"bool\",\"name\":\"ownerAccepted\",\"type\":\"bool\"},{\"indexed\":false,\"internalType\":\"bool\",\"name\":\"supplierAccepted\",\"type\":\"bool\"},{\"indexed\":false,\"internalType\":\"uint32\",\"name\":\"timesCountered\",\"type\":\"uint32\"},{\"indexed\":false,\"internalType\":\"uint32\",\"name\":\"timesRejected\",\"type\":\"uint32\"}],\"name\":\"CancellationPending\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"uint256\",\"name\":\"tokenId\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint16\",\"name\":\"cancellationReason\",\"type\":\"uint16\"},{\"indexed\":false,\"internalType\":\"uint16\",\"name\":\"cancellationReasonVersion\",\"type\":\"uint16\"},{\"indexed\":false,\"internalType\":\"uint16\",\"name\":\"rejectionReason\",\"type\":\"uint16\"},{\"indexed\":false,\"internalType\":\"uint16\",\"name\":\"rejectionVersion\",\"type\":\"uint16\"},{\"indexed\":false,\"internalType\":\"uint16\",\"name\":\"counterReason\",\"type\":\"uint16\"},{\"indexed\":false,\"internalType\":\"uint16\",\"name\":\"counterVersion\",\"type\":\"uint16\"},{\"indexed\":false,\"internalType\":\"uint16\",\"name\":\"withdrawalReason\",\"type\":\"uint16\"},{\"indexed\":false,\"internalType\":\"uint16\",\"name\":\"withdrawalVersion\",\"type\":\"uint16\"}],\"name\":\"CancellationReasons\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"uint256\",\"name\":\"tokenId\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint16\",\"name\":\"rejectionReason\",\"type\":\"uint16\"},{\"indexed\":false,\"internalType\":\"uint16\",\"name\":\"rejectionVersion\",\"type\":\"uint16\"}],\"name\":\"CancellationRejected\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"uint256\",\"name\":\"tokenId\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint16\",\"name\":\"withdrawalReason\",\"type\":\"uint16\"},{\"indexed\":false,\"internalType\":\"uint16\",\"name\":\"withdrawalVersion\",\"type\":\"uint16\"}],\"name\":\"CancellationWithdrawn\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint64\",\"name\":\"version\",\"type\":\"uint64\"}],\"name\":\"Initialized\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"oldManager\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"newManager\",\"type\":\"address\"}],\"name\":\"ManagerAddressUpdated\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"_tokenId\",\"type\":\"uint256\"}],\"name\":\"MetadataUpdate\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"oldDiff\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"newDiff\",\"type\":\"uint256\"}],\"name\":\"MinExpirationTimestampDiffUpdated\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"}],\"name\":\"Paused\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"bytes32\",\"name\":\"role\",\"type\":\"bytes32\"},{\"indexed\":true,\"internalType\":\"bytes32\",\"name\":\"previousAdminRole\",\"type\":\"bytes32\"},{\"indexed\":true,\"internalType\":\"bytes32\",\"name\":\"newAdminRole\",\"type\":\"bytes32\"}],\"name\":\"RoleAdminChanged\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"bytes32\",\"name\":\"role\",\"type\":\"bytes32\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"sender\",\"type\":\"address\"}],\"name\":\"RoleGranted\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"bytes32\",\"name\":\"role\",\"type\":\"bytes32\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"sender\",\"type\":\"address\"}],\"name\":\"RoleRevoked\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"uint256\",\"name\":\"tokenId\",\"type\":\"uint256\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"buyer\",\"type\":\"address\"}],\"name\":\"TokenBought\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"uint256\",\"name\":\"tokenId\",\"type\":\"uint256\"}],\"name\":\"TokenReservationExpired\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"uint256\",\"name\":\"tokenId\",\"type\":\"uint256\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"reservedFor\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"supplier\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"expirationTimestamp\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"price\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"contractIERC20\",\"name\":\"paymentToken\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"offchainPaymentCurrency\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"bool\",\"name\":\"cancellable\",\"type\":\"bool\"}],\"name\":\"TokenReserved\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"from\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"uint256\",\"name\":\"tokenId\",\"type\":\"uint256\"}],\"name\":\"Transfer\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"}],\"name\":\"Unpaused\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"implementation\",\"type\":\"address\"}],\"name\":\"Upgraded\",\"type\":\"event\"},{\"inputs\":[],\"name\":\"DEFAULT_ADMIN_ROLE\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"MIN_EXPIRATION_ADMIN_ROLE\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"NATIVE_PAYMENT\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"OFFCHAIN_PAYMENT\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"PAUSER_ROLE\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"UPGRADER_ROLE\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"UPGRADE_INTERFACE_VERSION\",\"outputs\":[{\"internalType\":\"string\",\"name\":\"\",\"type\":\"string\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"tokenId\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"refundAmount\",\"type\":\"uint256\"}],\"name\":\"acceptCancellation\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"tokenId\",\"type\":\"uint256\"}],\"name\":\"approve\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"owner\",\"type\":\"address\"}],\"name\":\"balanceOf\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"tokenId\",\"type\":\"uint256\"}],\"name\":\"buyReservedToken\",\"outputs\":[],\"stateMutability\":\"payable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"tokenId\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"refundAmount\",\"type\":\"uint256\"},{\"internalType\":\"uint16\",\"name\":\"counterReason\",\"type\":\"uint16\"},{\"internalType\":\"uint16\",\"name\":\"counterReasonVersion\",\"type\":\"uint16\"}],\"name\":\"counterCancellation\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"tokenId\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"checkRefundAmount\",\"type\":\"uint256\"}],\"name\":\"finalizeCancellation\",\"outputs\":[],\"stateMutability\":\"payable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"tokenId\",\"type\":\"uint256\"}],\"name\":\"getApproved\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"tokenId\",\"type\":\"uint256\"}],\"name\":\"getBookingStatus\",\"outputs\":[{\"internalType\":\"enumBookingToken.BookingStatus\",\"name\":\"\",\"type\":\"uint8\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"tokenId\",\"type\":\"uint256\"}],\"name\":\"getCancellationProposal\",\"outputs\":[{\"internalType\":\"enumCancellationProposalStatus\",\"name\":\"\",\"type\":\"uint8\"},{\"internalType\":\"uint256\",\"name\":\"refundAmount\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"initialProposer\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"currentProposer\",\"type\":\"address\"},{\"internalType\":\"bool\",\"name\":\"ownerAccepted\",\"type\":\"bool\"},{\"internalType\":\"bool\",\"name\":\"supplierAccepted\",\"type\":\"bool\"},{\"internalType\":\"uint32\",\"name\":\"timesCountered\",\"type\":\"uint32\"},{\"internalType\":\"uint32\",\"name\":\"timesRejected\",\"type\":\"uint32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"tokenId\",\"type\":\"uint256\"}],\"name\":\"getCancellationReasons\",\"outputs\":[{\"internalType\":\"uint16\",\"name\":\"cancellationReason\",\"type\":\"uint16\"},{\"internalType\":\"uint16\",\"name\":\"cancellationVersion\",\"type\":\"uint16\"},{\"internalType\":\"uint16\",\"name\":\"rejectionReason\",\"type\":\"uint16\"},{\"internalType\":\"uint16\",\"name\":\"rejectionVersion\",\"type\":\"uint16\"},{\"internalType\":\"uint16\",\"name\":\"counterReason\",\"type\":\"uint16\"},{\"internalType\":\"uint16\",\"name\":\"counterVersion\",\"type\":\"uint16\"},{\"internalType\":\"uint16\",\"name\":\"withdrawalReason\",\"type\":\"uint16\"},{\"internalType\":\"uint16\",\"name\":\"withdrawalVersion\",\"type\":\"uint16\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"getManagerAddress\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"getMinExpirationTimestampDiff\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"tokenId\",\"type\":\"uint256\"}],\"name\":\"getReservationPaymentToken\",\"outputs\":[{\"internalType\":\"contractIERC20\",\"name\":\"paymentToken\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"tokenId\",\"type\":\"uint256\"}],\"name\":\"getReservationPrice\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"price\",\"type\":\"uint256\"},{\"internalType\":\"contractIERC20\",\"name\":\"paymentToken\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"role\",\"type\":\"bytes32\"}],\"name\":\"getRoleAdmin\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"role\",\"type\":\"bytes32\"},{\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"}],\"name\":\"grantRole\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"role\",\"type\":\"bytes32\"},{\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"}],\"name\":\"hasRole\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"manager\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"defaultAdmin\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"upgrader\",\"type\":\"address\"}],\"name\":\"initialize\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"tokenId\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"refundAmount\",\"type\":\"uint256\"},{\"internalType\":\"uint16\",\"name\":\"cancellationReason\",\"type\":\"uint16\"},{\"internalType\":\"uint16\",\"name\":\"cancellationReasonVersion\",\"type\":\"uint16\"}],\"name\":\"initiateCancellation\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"owner\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"operator\",\"type\":\"address\"}],\"name\":\"isApprovedForAll\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"tokenId\",\"type\":\"uint256\"}],\"name\":\"isCancellable\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"}],\"name\":\"isTTMAccount\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"name\",\"outputs\":[{\"internalType\":\"string\",\"name\":\"\",\"type\":\"string\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"tokenId\",\"type\":\"uint256\"}],\"name\":\"ownerOf\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"pause\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"paused\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"proxiableUUID\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"tokenId\",\"type\":\"uint256\"}],\"name\":\"recordExpiration\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"tokenId\",\"type\":\"uint256\"},{\"internalType\":\"uint16\",\"name\":\"rejectionReason\",\"type\":\"uint16\"},{\"internalType\":\"uint16\",\"name\":\"rejectionReasonVersion\",\"type\":\"uint16\"}],\"name\":\"rejectCancellation\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"role\",\"type\":\"bytes32\"},{\"internalType\":\"address\",\"name\":\"callerConfirmation\",\"type\":\"address\"}],\"name\":\"renounceRole\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"role\",\"type\":\"bytes32\"},{\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"}],\"name\":\"revokeRole\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"reservedFor\",\"type\":\"address\"},{\"internalType\":\"string\",\"name\":\"uri\",\"type\":\"string\"},{\"internalType\":\"uint256\",\"name\":\"expirationTimestamp\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"price\",\"type\":\"uint256\"},{\"internalType\":\"contractIERC20\",\"name\":\"paymentToken\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"offchainPaymentCurrency\",\"type\":\"uint256\"},{\"internalType\":\"bool\",\"name\":\"cancellable\",\"type\":\"bool\"}],\"name\":\"safeMintWithReservation\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"from\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"tokenId\",\"type\":\"uint256\"}],\"name\":\"safeTransferFrom\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"from\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"tokenId\",\"type\":\"uint256\"},{\"internalType\":\"bytes\",\"name\":\"data\",\"type\":\"bytes\"}],\"name\":\"safeTransferFrom\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"operator\",\"type\":\"address\"},{\"internalType\":\"bool\",\"name\":\"approved\",\"type\":\"bool\"}],\"name\":\"setApprovalForAll\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"manager\",\"type\":\"address\"}],\"name\":\"setManagerAddress\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"minExpirationTimestampDiff\",\"type\":\"uint256\"}],\"name\":\"setMinExpirationTimestampDiff\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes4\",\"name\":\"interfaceId\",\"type\":\"bytes4\"}],\"name\":\"supportsInterface\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"symbol\",\"outputs\":[{\"internalType\":\"string\",\"name\":\"\",\"type\":\"string\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"index\",\"type\":\"uint256\"}],\"name\":\"tokenByIndex\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"owner\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"index\",\"type\":\"uint256\"}],\"name\":\"tokenOfOwnerByIndex\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"tokenId\",\"type\":\"uint256\"}],\"name\":\"tokenURI\",\"outputs\":[{\"internalType\":\"string\",\"name\":\"\",\"type\":\"string\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"totalSupply\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"from\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"tokenId\",\"type\":\"uint256\"}],\"name\":\"transferFrom\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"unpause\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"newImplementation\",\"type\":\"address\"},{\"internalType\":\"bytes\",\"name\":\"data\",\"type\":\"bytes\"}],\"name\":\"upgradeToAndCall\",\"outputs\":[],\"stateMutability\":\"payable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"version\",\"outputs\":[{\"internalType\":\"uint16\",\"name\":\"major\",\"type\":\"uint16\"},{\"internalType\":\"uint16\",\"name\":\"minor\",\"type\":\"uint16\"},{\"internalType\":\"uint16\",\"name\":\"patch\",\"type\":\"uint16\"}],\"stateMutability\":\"pure\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"tokenId\",\"type\":\"uint256\"},{\"internalType\":\"uint16\",\"name\":\"withdrawalReason\",\"type\":\"uint16\"},{\"internalType\":\"uint16\",\"name\":\"withdrawalReasonVersion\",\"type\":\"uint16\"}],\"name\":\"withdrawCancellation\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"}]",
	Bin: "0x60a060405230608052348015610013575f80fd5b506080516153a361003a5f395f81816129c0015281816129e90152612b8901526153a35ff3fe60806040526004361061034f575f3560e01c80636352211e116101bd578063b88d4fde116100f2578063d547741f11610092578063e63ab1e91161006d578063e63ab1e914610c74578063e985e9c514610ca7578063f72c0d8b14610d0d578063f7e45f0914610d40575f80fd5b8063d547741f14610c17578063db2b268214610c36578063e5a6725c14610c55575f80fd5b8063bfb26c06116100cd578063bfb26c0614610b9c578063c0c53b8b14610bb0578063c162d7da14610bcf578063c87b56dd14610bf8575f80fd5b8063b88d4fde14610ab0578063bb520b4714610acf578063be66718814610b7d575f80fd5b806396591edd1161015d578063a22cb46511610138578063a22cb46514610903578063a9bc55a214610922578063ad3cb1cc14610a12578063b191d09214610a5a575f80fd5b806396591edd146108ca578063a0f07c74146108dd578063a217fddf146108f0575f80fd5b806374fe60e91161019857806374fe60e9146108205780638456cb591461083f57806391d148541461085357806395d89b41146108b6575f80fd5b80636352211e146107c357806370a08231146107e257806374aa204814610801575f80fd5b80632f2ff15d116102935780634f1ef2861161023357806352d1902d1161020e57806352d1902d1461072f57806354fd4d50146107435780635c975abb1461076e5780635f7290f1146107a4575f80fd5b80634f1ef286146106de5780634f6ccce7146106f1578063516a82b814610710575f80fd5b80633c15b31c1161026e5780633c15b31c146106325780633f4ba83a1461068c57806341431908146106a057806342842e0e146106bf575f80fd5b80632f2ff15d146105d55780632f745c59146105f457806336568abe14610613575f80fd5b806318160ddd116102fe578063248a9ca3116102d9578063248a9ca3146104e65780632a119380146105335780632d3a6329146105525780632edf5e2c146105a2575f80fd5b806318160ddd146104a05780631c54f0f7146104b457806323b872dd146104c7575f80fd5b8063081812fc1161032e578063081812fc14610426578063095ea7b31461045d5780630e75c1a81461047e575f80fd5b80624fdd3c1461035357806301ffc9a7146103d657806306fdde0314610405575b5f80fd5b34801561035e575f80fd5b506103b461036d366004614b2b565b5f9081527f54347f9bbfbc2e7b5786abab9693d2ba67834fa4787a7740d72712e87e20740360205260409020600381015460049091015490916001600160a01b0390911690565b604080519283526001600160a01b039091166020830152015b60405180910390f35b3480156103e1575f80fd5b506103f56103f0366004614b57565b610d5f565b60405190151581526020016103cd565b348015610410575f80fd5b50610419610d6f565b6040516103cd9190614bbf565b348015610431575f80fd5b50610445610440366004614b2b565b610e23565b6040516001600160a01b0390911681526020016103cd565b348015610468575f80fd5b5061047c610477366004614be5565b610e69565b005b348015610489575f80fd5b50610492610e78565b6040519081526020016103cd565b3480156104ab575f80fd5b50610492610e93565b61047c6104c2366004614c0f565b610ebb565b3480156104d2575f80fd5b5061047c6104e1366004614c2f565b610f79565b3480156104f1575f80fd5b50610492610500366004614b2b565b5f9081527f02dd7bc7dec4dceedda775e58dd541e08a116c6c53815c0bd028192f7b626800602052604090206001015490565b34801561053e575f80fd5b5061047c61054d366004614c83565b610f92565b34801561055d575f80fd5b506103f561056c366004614b2b565b5f9081527f54347f9bbfbc2e7b5786abab9693d2ba67834fa4787a7740d72712e87e207403602052604090206006015460ff1690565b3480156105ad575f80fd5b506104927f3ae8648b97d3fd425d26286fc6bb1d50724a93a6a5763921dd2b90405a83b4a481565b3480156105e0575f80fd5b5061047c6105ef366004614cbc565b610fc0565b3480156105ff575f80fd5b5061049261060e366004614be5565b611009565b34801561061e575f80fd5b5061047c61062d366004614cbc565b611092565b34801561063d575f80fd5b5061067f61064c366004614b2b565b5f9081527f54347f9bbfbc2e7b5786abab9693d2ba67834fa4787a7740d72712e87e207404602052604090205460ff1690565b6040516103cd9190614d1a565b348015610697575f80fd5b5061047c6110de565b3480156106ab575f80fd5b5061047c6106ba366004614d2d565b611113565b3480156106ca575f80fd5b5061047c6106d9366004614c2f565b61117b565b61047c6106ec366004614ded565b611195565b3480156106fc575f80fd5b5061049261070b366004614b2b565b6111b0565b34801561071b575f80fd5b5061047c61072a366004614b2b565b611228565b34801561073a575f80fd5b506104926112c7565b34801561074e575f80fd5b5060408051600181525f60208201819052918101919091526060016103cd565b348015610779575f80fd5b507fcd5ed15c6e187e77e9aee88184c21f4f2182ab5827cb3b7e07fbedcd63f033005460ff166103f5565b3480156107af575f80fd5b506103f56107be366004614d2d565b6112f5565b3480156107ce575f80fd5b506104456107dd366004614b2b565b611397565b3480156107ed575f80fd5b506104926107fc366004614d2d565b6113a1565b34801561080c575f80fd5b5061047c61081b366004614e3a565b611425565b34801561082b575f80fd5b5061047c61083a366004614c83565b611455565b34801561084a575f80fd5b5061047c61147b565b34801561085e575f80fd5b506103f561086d366004614cbc565b5f9182527f02dd7bc7dec4dceedda775e58dd541e08a116c6c53815c0bd028192f7b626800602090815260408084206001600160a01b0393909316845291905290205460ff1690565b3480156108c1575f80fd5b506104196114ad565b61047c6108d8366004614b2b565b6114fe565b3480156108e8575f80fd5b506104455f81565b3480156108fb575f80fd5b506104925f81565b34801561090e575f80fd5b5061047c61091d366004614e8a565b61175e565b34801561092d575f80fd5b506109c561093c366004614b2b565b5f9081527f1973af827d98d3fe49ea0c9179b69c8f267978d372df89173953530663b7220060205260409020600281015460039091015461ffff600160c81b8304811693600160d81b8404821693600160e81b900482169282811692620100008204811692640100000000830482169266010000000000008104831692600160401b9091041690565b6040805161ffff998a16815297891660208901529588169587019590955292861660608601529085166080850152841660a0840152831660c083015290911660e0820152610100016103cd565b348015610a1d575f80fd5b506104196040518060400160405280600581526020017f352e302e3000000000000000000000000000000000000000000000000000000081525081565b348015610a65575f80fd5b50610445610a74366004614b2b565b5f9081527f54347f9bbfbc2e7b5786abab9693d2ba67834fa4787a7740d72712e87e20740360205260409020600401546001600160a01b031690565b348015610abb575f80fd5b5061047c610aca366004614eb6565b611769565b348015610ada575f80fd5b50610b69610ae9366004614b2b565b5f9081527f1973af827d98d3fe49ea0c9179b69c8f267978d372df89173953530663b72200602052604090206002810154815460019092015460ff600160c01b8084048216956001600160a01b038085169590811694928304841693600160c81b8404169263ffffffff600160a01b918290048116939190920490911690565b6040516103cd989796959493929190614f1e565b348015610b88575f80fd5b5061047c610b97366004614c0f565b611781565b348015610ba7575f80fd5b50610445600181565b348015610bbb575f80fd5b5061047c610bca366004614f76565b6117ad565b348015610bda575f80fd5b505f8051602061534e833981519152546001600160a01b0316610445565b348015610c03575f80fd5b50610419610c12366004614b2b565b611a5a565b348015610c22575f80fd5b5061047c610c31366004614cbc565b611a65565b348015610c41575f80fd5b5061047c610c50366004614fbe565b611aa8565b348015610c60575f80fd5b5061047c610c6f366004614b2b565b611d22565b348015610c7f575f80fd5b506104927f65d7a28e3265b37a6474929f336521b332c1681b933f6cb9f3376673440d862a81565b348015610cb2575f80fd5b506103f5610cc136600461505f565b6001600160a01b039182165f9081527f80bb2b638cc20bc4d0a60d66940f3ab4a00c1d7b313497ca82fb0b4ab00793056020908152604080832093909416825291909152205460ff1690565b348015610d18575f80fd5b506104927f189ab7a9244df0848122154315af71fe140f3db0fe014031783b0946b8c9d2e381565b348015610d4b575f80fd5b5061047c610d5a366004614e3a565b611e7b565b5f610d6982611ea2565b92915050565b7f80bb2b638cc20bc4d0a60d66940f3ab4a00c1d7b313497ca82fb0b4ab00793008054606091908190610da19061508b565b80601f0160208091040260200160405190810160405280929190818152602001828054610dcd9061508b565b8015610e185780601f10610def57610100808354040283529160200191610e18565b820191905f5260205f20905b815481529060010190602001808311610dfb57829003601f168201915b505050505091505090565b5f610e2d82611edf565b505f8281527f80bb2b638cc20bc4d0a60d66940f3ab4a00c1d7b313497ca82fb0b4ab007930460205260409020546001600160a01b0316610d69565b610e74828233611f36565b5050565b5f805f8051602061534e8339815191525b6002015492915050565b5f807f645e039705490088daad89bae25049a34f4a9072d398537b1ab2425f24cbed00610e89565b610ec3611f43565b33610ecd81611fa6565b610ed5611ff0565b5f80610ee08561204e565b915091505f610ef95f8051602061534e83398151915290565b90505f610f07838888612109565b5f88815260038401602090815260408083206004908101548188019093529220805460ff19169092179091559091506001600160a01b0316610f4a8183876122d6565b505050505050610e7460017f9b779b17422d0df92223018b32b4d1fa46e071723d6817e2486d003becc55f0055565b610f82816123a1565b610f8d8383836125b8565b505050565b33610f9c81611fa6565b5f80610fa78661204e565b91509150610fb8828288888861263b565b505050505050565b5f8281527f02dd7bc7dec4dceedda775e58dd541e08a116c6c53815c0bd028192f7b6268006020526040902060010154610ff9816127c9565b61100383836127d3565b50505050565b5f7f645e039705490088daad89bae25049a34f4a9072d398537b1ab2425f24cbed00611034846113a1565b831061106a5760405163295f44f760e21b81526001600160a01b0385166004820152602481018490526044015b60405180910390fd5b6001600160a01b0384165f908152602091825260408082208583529092522054905092915050565b6001600160a01b03811633146110d4576040517f6697b23200000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b610f8d828261289f565b7f65d7a28e3265b37a6474929f336521b332c1681b933f6cb9f3376673440d862a611108816127c9565b611110612943565b50565b5f61111d816127c9565b5f8051602061534e83398151915280546001600160a01b031981166001600160a01b03858116918217845560405192169182907f9462e60b9d7b78dcca266b08b885d2cd87178de9a5c63e600065b86e530f0b9b905f90a350505050565b610f8d83838360405180602001604052805f815250611769565b61119d6129b5565b6111a682612a6c565b610e748282612a96565b5f7f645e039705490088daad89bae25049a34f4a9072d398537b1ab2425f24cbed006111da610e93565b83106112025760405163295f44f760e21b81525f600482015260248101849052604401611061565b806002018381548110611217576112176150c3565b905f5260205f200154915050919050565b7f3ae8648b97d3fd425d26286fc6bb1d50724a93a6a5763921dd2b90405a83b4a4611252816127c9565b7f54347f9bbfbc2e7b5786abab9693d2ba67834fa4787a7740d72712e87e207402805490839055604080518281526020810185905281515f8051602061534e83398151915293927f6175237436049150327545b616ef840ddeb5def6cd197617a415f10fc838fdb5928290030190a150505050565b5f6112d0612b7e565b507f360894a13ba1a3210667c828492db98dca3e2076cc3735a920a3ca505d382bbc90565b5f6113145f8051602061534e833981519152546001600160a01b031690565b6040517f5f7290f10000000000000000000000000000000000000000000000000000000081526001600160a01b0384811660048301529190911690635f7290f190602401602060405180830381865afa158015611373573d5f803e3d5ffd5b505050506040513d601f19601f82011682018060405250810190610d6991906150d7565b5f610d6982611edf565b5f7f80bb2b638cc20bc4d0a60d66940f3ab4a00c1d7b313497ca82fb0b4ab00793006001600160a01b038316611405576040517f89c62b640000000000000000000000000000000000000000000000000000000081525f6004820152602401611061565b6001600160a01b039092165f908152600390920160205250604090205490565b3361142f81611fa6565b5f8061143a8761204e565b9150915061144c828289898989612bc7565b50505050505050565b3361145f81611fa6565b5f8061146a8661204e565b91509150610fb88282888888612e7e565b7f65d7a28e3265b37a6474929f336521b332c1681b933f6cb9f3376673440d862a6114a5816127c9565b611110613041565b7f80bb2b638cc20bc4d0a60d66940f3ab4a00c1d7b313497ca82fb0b4ab007930180546060917f80bb2b638cc20bc4d0a60d66940f3ab4a00c1d7b313497ca82fb0b4ab007930091610da19061508b565b611506611f43565b61150e611ff0565b3361151881611fa6565b5f8281527f54347f9bbfbc2e7b5786abab9693d2ba67834fa4787a7740d72712e87e2074036020908152604091829020825160e08101845281546001600160a01b0390811680835260018401548216948301949094526002830154948201949094526003820154606082015260048201549093166080840152600581015460a08401526006015460ff16151560c08301525f8051602061534e833981519152919033146116055780516040517f4cc7538a0000000000000000000000000000000000000000000000000000000081526001600160a01b039091166004820152336024820152604401611061565b8060400151421115611654578381604001516040517f527ae76e000000000000000000000000000000000000000000000000000000008152600401611061929190918252602082015260400190565b5f61165e85611397565b905081602001516001600160a01b0316816001600160a01b0316146116c75760208201516040517f81e8a2c8000000000000000000000000000000000000000000000000000000008152600481018790526001600160a01b039091166024820152604401611061565b6116d68260200151338761309c565b6116ed8260800151836060015184602001516122d6565b5f858152600484016020526040808220805460ff1916600317905551339187917fa751fb02c318279a22135a408663ae08ea45eafa950a4351c14ae543cbb950409190a35050505061111060017f9b779b17422d0df92223018b32b4d1fa46e071723d6817e2486d003becc55f0055565b610e74338383613149565b611774848484610f79565b6110033385858585613224565b3361178b81611fa6565b5f806117968561204e565b915091506117a682828787613343565b5050505050565b7ff0c57e16840df040f15088dc2f81fe391c3923bec73e23a9662efc9c229c6a008054600160401b810460ff16159067ffffffffffffffff165f811580156117f25750825b90505f8267ffffffffffffffff16600114801561180e5750303b155b90508115801561181c575080155b15611853576040517ff92ee8a900000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b845467ffffffffffffffff19166001178555831561188257845468ff00000000000000001916600160401b1785555b6001600160a01b038816158061189f57506001600160a01b038716155b806118b157506001600160a01b038616155b156118e8576040517fd92e233d00000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b61195c6040518060400160405280600c81526020017f426f6f6b696e67546f6b656e00000000000000000000000000000000000000008152506040518060400160405280600681526020017f42546f6b656e000000000000000000000000000000000000000000000000000081525061358b565b61196461359d565b61196c61359d565b61197461359d565b61197c6135a5565b61198461359d565b61198e5f886127d3565b506119b97f189ab7a9244df0848122154315af71fe140f3db0fe014031783b0946b8c9d2e3876127d3565b505f8051602061534e83398151915280546001600160a01b0319166001600160a01b038a16179055603c7f54347f9bbfbc2e7b5786abab9693d2ba67834fa4787a7740d72712e87e207402558315611a5057845468ff000000000000000019168555604051600181527fc7f505b2f371ae2175ee4913f4499e1f2633a7b5936321eed1cdaeb6115181d29060200160405180910390a15b5050505050505050565b6060610d69826135b5565b5f8281527f02dd7bc7dec4dceedda775e58dd541e08a116c6c53815c0bd028192f7b6268006020526040902060010154611a9e816127c9565b611003838361289f565b33611ab281611fa6565b611aba611ff0565b611ac388611fa6565b7f54347f9bbfbc2e7b5786abab9693d2ba67834fa4787a7740d72712e87e207402545f8051602061534e83398151915290611afe8142615106565b8811611b40576040517f999f7d700000000000000000000000000000000000000000000000000000000081526004810189905260248101829052604401611061565b5f85118015611b5957506001600160a01b038616600114155b15611b93576040517f8fe757e700000000000000000000000000000000000000000000000000000000815260048101869052602401611061565b6001820180545f9182611ba583615119565b919050559050611bb533826136e4565b611bbf818b6136fd565b6040805160e0810182526001600160a01b03808e1682523360208084019182528385018e8152606085018e8152848e166080870190815260a087018e81528d151560c089019081525f8b81527f54347f9bbfbc2e7b5786abab9693d2ba67834fa4787a7740d72712e87e20740390965298909420965187546001600160a01b0319908116918816919091178855945160018801805487169188169190911790559151600287015551600386015551600485018054909316931692909217905551600582015590516006909101805460ff19169115159190911790555f818152600484016020908152604091829020805460ff1916600117905581518b81529081018a90526001600160a01b0389811682840152606082018990528715156080830152915133928e169184917f1424af4f4cb40d8a1a2d00b2324cb122ba73eac426f98b62c33ff31ca045f0679160a0908290030190a45050505050505050505050565b5f8181527f54347f9bbfbc2e7b5786abab9693d2ba67834fa4787a7740d72712e87e207403602090815260408083207f54347f9bbfbc2e7b5786abab9693d2ba67834fa4787a7740d72712e87e207404909252909120545f8051602061534e833981519152919060ff166002816004811115611da057611da0614cea565b1480611dbd57506003816004811115611dbb57611dbb614cea565b145b80611dd957506004816004811115611dd757611dd7614cea565b145b15611dfb57838160405163e4e3b53b60e01b8152600401611061929190615131565b8160020154421115611e4d575f848152600484016020526040808220805460ff191660021790555185917fc47ab59d5c41a0220b594ece6f7e87863b07f8b33579130f3a59b31d8f7b6eed91a2611003565b815460405163d4cde2af60e01b8152600481018690526001600160a01b039091166024820152604401611061565b33611e8581611fa6565b5f80611e908761204e565b9150915061144c82828989898961376f565b5f6001600160e01b031982167f7965db0b000000000000000000000000000000000000000000000000000000001480610d695750610d6982613aa0565b5f8181527f80bb2b638cc20bc4d0a60d66940f3ab4a00c1d7b313497ca82fb0b4ab007930260205260408120546001600160a01b031680610d6957604051637e27328960e01b815260048101849052602401611061565b610f8d8383836001613add565b7f9b779b17422d0df92223018b32b4d1fa46e071723d6817e2486d003becc55f00805460011901611fa0576040517f3ee5aeb500000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b60029055565b611faf816112f5565b611110576040517f27b12cd70000000000000000000000000000000000000000000000000000000081526001600160a01b0382166004820152602401611061565b7fcd5ed15c6e187e77e9aee88184c21f4f2182ab5827cb3b7e07fbedcd63f033005460ff161561204c576040517fd93c066500000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b565b5f8061205983611edf565b5f8481527f54347f9bbfbc2e7b5786abab9693d2ba67834fa4787a7740d72712e87e20740460205260409020549092505f8051602061534e8339815191529060039060ff1660048111156120af576120af614cea565b146120e5575f848152600480830160205260409182902054915163e4e3b53b60e01b815261106192879260ff9091169101615131565b5f9384526003016020525060409091206001015490916001600160a01b0390911690565b5f336001600160a01b0385161461214f576040517f6c83fb1b00000000000000000000000000000000000000000000000000000000815260048101849052602401611061565b5f8381527f1973af827d98d3fe49ea0c9179b69c8f267978d372df89173953530663b72200602052604090206002810154600190600160c01b900460ff16600481111561219e5761219e614cea565b146121d157838160020160189054906101000a900460ff1660405163d5d8eacb60e01b8152600401611061929190615131565b8054831461220657805460405163cc45283760e01b815260048101869052602481019190915260448101849052606401611061565b6001810154600160c01b900460ff1661224e576040517fc84052f900000000000000000000000000000000000000000000000000000000815260048101859052602401611061565b6001810154600160c81b900460ff166122775760018101805460ff60c81b1916600160c81b1790555b60028101805460ff60c01b1916780400000000000000000000000000000000000000000000000017905560405184907f17c3690813e5ff9135b87fd91848109978b23db8e471498d18886560da7f2867905f90a25490505b9392505050565b6001600160a01b03831661233957813414612326576040517f0515845400000000000000000000000000000000000000000000000000000000815234600482015260248101839052604401611061565b610f8d6001600160a01b03821634613c59565b5f196001600160a01b0384160161236b573415610f8d576040516347d6729960e01b8152346004820152602401611061565b341561238c576040516347d6729960e01b8152346004820152602401611061565b610f8d6001600160a01b038416338385613d0c565b5f8181527f54347f9bbfbc2e7b5786abab9693d2ba67834fa4787a7740d72712e87e20740460209081526040808320547f1973af827d98d3fe49ea0c9179b69c8f267978d372df89173953530663b72200909252909120600201545f8051602061534e8339815191529160ff90811691600160c01b8104909116906001600160a01b0316600182600481111561243957612439614cea565b03612498575f61244886611edf565b5f8781526003870160205260409020600101549091506001600160a01b0390811690831633146124865761248182828960636001612e7e565b612495565b6124958282896063600161263b565b50505b60038360048111156124ac576124ac614cea565b14806124c9575060028360048111156124c7576124c7614cea565b145b806124e457505f8360048111156124e2576124e2614cea565b145b156124f0575050505050565b600483600481111561250457612504614cea565b0361252657848360405163e4e3b53b60e01b8152600401611061929190615131565b5f8581526003850160205260409020600281015442111561258a575f868152600486016020526040808220805460ff191660021790555187917fc47ab59d5c41a0220b594ece6f7e87863b07f8b33579130f3a59b31d8f7b6eed91a2505050505050565b805460405163d4cde2af60e01b8152600481018890526001600160a01b039091166024820152604401611061565b6001600160a01b0382166125e157604051633250574960e11b81525f6004820152602401611061565b5f6125ed838333613d94565b9050836001600160a01b0316816001600160a01b031614611003576040516364283d7b60e01b81526001600160a01b0380861660048301526024820184905282166044820152606401611061565b84846126478282613da8565b5f8581527f1973af827d98d3fe49ea0c9179b69c8f267978d372df89173953530663b72200602052604090206002810154600190600160c01b900460ff16600481111561269657612696614cea565b146126c957858160020160189054906101000a900460ff1660405163d5d8eacb60e01b8152600401611061929190615131565b60028101546001600160a01b03163314612712576040517f8c37ede800000000000000000000000000000000000000000000000000000000815260048101879052602401611061565b60038101805469ffffffff0000000000001916660100000000000061ffff88811691820269ffff0000000000000000191692909217600160401b92881692830217909255600283018054780300000000000000000000000000000000000000000000000060ff60c01b1990911617905560408051928352602083019190915287917f48e256ce3da490e3bbba80f056bb54ec3d7264f8ad7d152b77bf8c2eca3db5a591015b60405180910390a25050505050505050565b6111108133613e01565b5f8281527f02dd7bc7dec4dceedda775e58dd541e08a116c6c53815c0bd028192f7b626800602081815260408084206001600160a01b038616855290915282205460ff16612896575f848152602082815260408083206001600160a01b03871684529091529020805460ff1916600117905561284c3390565b6001600160a01b0316836001600160a01b0316857f2f8788117e7eff1d82e926ec794901d17c78024a50270940304540a733656f0d60405160405180910390a46001915050610d69565b5f915050610d69565b5f8281527f02dd7bc7dec4dceedda775e58dd541e08a116c6c53815c0bd028192f7b626800602081815260408084206001600160a01b038616855290915282205460ff1615612896575f848152602082815260408083206001600160a01b0387168085529252808320805460ff1916905551339287917ff6391f5c32d9c69d2a47ea670b442974b53935d1edc7fd64eb21e047a839171b9190a46001915050610d69565b61294b613e8d565b7fcd5ed15c6e187e77e9aee88184c21f4f2182ab5827cb3b7e07fbedcd63f03300805460ff191681557f5db9ee0a495bf2e6ff9c91a7834c1ba4fdd244a5e8aa4e537bd38aeae4b073aa335b6040516001600160a01b03909116815260200160405180910390a150565b306001600160a01b037f0000000000000000000000000000000000000000000000000000000000000000161480612a4e57507f00000000000000000000000000000000000000000000000000000000000000006001600160a01b0316612a427f360894a13ba1a3210667c828492db98dca3e2076cc3735a920a3ca505d382bbc546001600160a01b031690565b6001600160a01b031614155b1561204c5760405163703e46dd60e11b815260040160405180910390fd5b7f189ab7a9244df0848122154315af71fe140f3db0fe014031783b0946b8c9d2e3610e74816127c9565b816001600160a01b03166352d1902d6040518163ffffffff1660e01b8152600401602060405180830381865afa925050508015612af0575060408051601f3d908101601f19168201909252612aed9181019061514e565b60015b612b1857604051634c9c8ce360e01b81526001600160a01b0383166004820152602401611061565b7f360894a13ba1a3210667c828492db98dca3e2076cc3735a920a3ca505d382bbc8114612b74576040517faa1d49a400000000000000000000000000000000000000000000000000000000815260048101829052602401611061565b610f8d8383613ee8565b306001600160a01b037f0000000000000000000000000000000000000000000000000000000000000000161461204c5760405163703e46dd60e11b815260040160405180910390fd5b8585612bd38282613da8565b5f8681527f1973af827d98d3fe49ea0c9179b69c8f267978d372df89173953530663b72200602052604090206002810154600190600160c01b900460ff166004811115612c2257612c22614cea565b14612c5557868160020160189054906101000a900460ff1660405163d5d8eacb60e01b8152600401611061929190615131565b858155600281018054336001600160a01b031990911681179091556001820180547fffffffffffff0000ffffffffffffffffffffffffffffffffffffffffffffffff166001600160a01b038c81168414600160c01b0260ff60c81b191691909117908b16909214600160c81b029190911780825560038301805465ffffffff000019166201000061ffff8a81169190910265ffff0000000019169190911764010000000091891691909102179055600160a01b900463ffffffff16906014612d1c83615165565b82546101009290920a63ffffffff81810219909316918316021790915560028301546001840154845460408051918252600160c01b830460ff90811615156020840152600160c81b840416151590820152600160a01b80830485166060830152830490931660808401526001600160a01b039182169350169089907f6aafaf6639750382d6dc665bc6a3c73d54e2076a8739a69c860a2314032eb9819060a0015b60405180910390a46002810154600382015460408051600160c81b840461ffff9081168252600160d81b850481166020830152600160e81b9094048416818301528383166060820152620100008304841660808201526401000000008304841660a082015266010000000000008304841660c0820152600160401b90920490921660e0820152905188917fade2ce8b39037f5140109f53fef47c0491427b8b785ab83d35aa2931295cb79091908190036101000190a2505050505050505050565b8484612e8a8282613da8565b5f8581527f1973af827d98d3fe49ea0c9179b69c8f267978d372df89173953530663b72200602052604090206002810154600190600160c01b900460ff166004811115612ed957612ed9614cea565b14612f0c57858160020160189054906101000a900460ff1660405163d5d8eacb60e01b8152600401611061929190615131565b60028101546001600160a01b03163303612f55576040517fc18363dc00000000000000000000000000000000000000000000000000000000815260048101879052602401611061565b60028101805460038301805461ffff88811661ffff19909216919091179091558716600160e81b0260ff60c01b19167fff0000ffffffff00ffffffffffffffffffffffffffffffffffffffffffffffff90911617780200000000000000000000000000000000000000000000000017808255600160a01b900463ffffffff16906014612fe083615165565b91906101000a81548163ffffffff021916908363ffffffff16021790555050857fab78ba855f2fdb28beb212a9b3f41a33cda034729848cd452f0cc96528c23a8086866040516127b792919061ffff92831681529116602082015260400190565b613049611ff0565b7fcd5ed15c6e187e77e9aee88184c21f4f2182ab5827cb3b7e07fbedcd63f03300805460ff191660011781557f62e78cea01bee320cd4e420270b5ea74000d11b0c9f74754ebdbfc544b05a25833612997565b6001600160a01b0382166130c557604051633250574960e11b81525f6004820152602401611061565b5f6130d183835f613d94565b90506001600160a01b0381166130fd57604051637e27328960e01b815260048101839052602401611061565b836001600160a01b0316816001600160a01b031614611003576040516364283d7b60e01b81526001600160a01b0380861660048301526024820184905282166044820152606401611061565b7f80bb2b638cc20bc4d0a60d66940f3ab4a00c1d7b313497ca82fb0b4ab00793006001600160a01b0383166131b5576040517f5b08ba180000000000000000000000000000000000000000000000000000000081526001600160a01b0384166004820152602401611061565b6001600160a01b038481165f818152600584016020908152604080832094881680845294825291829020805460ff191687151590811790915591519182527f17307eab39ab6107e8899845ad3d59bd9653f200f220920489ca2b5937696c31910160405180910390a350505050565b6001600160a01b0383163b156117a657604051630a85bd0160e11b81526001600160a01b0384169063150b7a0290613266908890889087908790600401615187565b6020604051808303815f875af19250505080156132a0575060408051601f3d908101601f1916820190925261329d918101906151c2565b60015b613307573d8080156132cd576040519150601f19603f3d011682016040523d82523d5f602084013e6132d2565b606091505b5080515f036132ff57604051633250574960e11b81526001600160a01b0385166004820152602401611061565b805181602001fd5b6001600160e01b03198116630a85bd0160e11b14610fb857604051633250574960e11b81526001600160a01b0385166004820152602401611061565b838361334f8282613da8565b5f8481527f1973af827d98d3fe49ea0c9179b69c8f267978d372df89173953530663b72200602052604090206002810154600190600160c01b900460ff16600481111561339e5761339e614cea565b146133d157848160020160189054906101000a900460ff1660405163d5d8eacb60e01b8152600401611061929190615131565b8054841461340657805460405163cc45283760e01b815260048101879052602481019190915260448101859052606401611061565b6001600160a01b03871633036134305760018101805460ff60c01b1916600160c01b179055613446565b60018101805460ff60c81b1916600160c81b1790555b6002810154600182015482546040805191825260ff600160c01b8404811615156020840152600160c81b84041615159082015263ffffffff600160a01b8084048216606084015284041660808201526001600160a01b03928316929091169087907f6aafaf6639750382d6dc665bc6a3c73d54e2076a8739a69c860a2314032eb9819060a00160405180910390a46002810154600382015460408051600160c81b840461ffff9081168252600160d81b850481166020830152600160e81b9094048416818301528383166060820152620100008304841660808201526401000000008304841660a082015266010000000000008304841660c0820152600160401b90920490921660e0820152905186917fade2ce8b39037f5140109f53fef47c0491427b8b785ab83d35aa2931295cb79091908190036101000190a250505050505050565b613593613f3d565b610e748282613f9f565b61204c613f3d565b6135ad613f3d565b61204c613fe2565b60607f0542a41881ee128a365a727b282c86fa859579490b9bb45aab8503648c8e79006135e183611edf565b505f83815260208290526040812080546135fa9061508b565b80601f01602080910402602001604051908101604052809291908181526020018280546136269061508b565b80156136715780601f1061364857610100808354040283529160200191613671565b820191905f5260205f20905b81548152906001019060200180831161365457829003601f168201915b505050505090505f61368d60408051602081019091525f815290565b905080515f0361369f57509392505050565b8151156136d25780826040516020016136b99291906151dd565b6040516020818303038152906040529350505050919050565b6136db85614015565b95945050505050565b610e74828260405180602001604052805f815250614085565b5f8281527f0542a41881ee128a365a727b282c86fa859579490b9bb45aab8503648c8e790060208190526040909120613736838261524f565b506040518381527ff8e1a15aba9398e019f0b49df1a4fde98ee17ae345cb5f6b5e2c27f5033e8ce79060200160405180910390a1505050565b858561377b8282613da8565b5f8681527f1973af827d98d3fe49ea0c9179b69c8f267978d372df89173953530663b72200602052604090206002810154600490600160c01b900460ff16818111156137c9576137c9614cea565b14806137f4575060016002820154600160c01b900460ff1660048111156137f2576137f2614cea565b145b1561382757868160020160189054906101000a900460ff1660405163d5d8eacb60e01b8152600401611061929190615131565b5f6002820154600160c01b900460ff16600481111561384857613848614cea565b03613862576001810180546001600160a01b031916331790555b33816002015f6101000a8154816001600160a01b0302191690836001600160a01b0316021790555085815f0181905550886001600160a01b0316336001600160a01b0316148160010160186101000a81548160ff021916908315150217905550876001600160a01b0316336001600160a01b0316148160010160196101000a81548160ff021916908315150217905550848160020160196101000a81548161ffff021916908361ffff1602179055508381600201601b6101000a81548161ffff021916908361ffff1602179055505f81600201601d6101000a81548161ffff021916908361ffff1602179055505f816003015f6101000a81548161ffff021916908361ffff1602179055505f8160030160026101000a81548161ffff021916908361ffff1602179055505f8160030160046101000a81548161ffff021916908361ffff1602179055505f8160030160066101000a81548161ffff021916908361ffff1602179055505f8160030160086101000a81548161ffff021916908361ffff16021790555060018160020160186101000a81548160ff02191690836004811115613a1057613a10614cea565b02179055506002810154600182015482546040805191825260ff600160c01b8404811615156020840152600160c81b84041615159082015263ffffffff600160a01b8084048216606084015284041660808201526001600160a01b03928316929091169089907f6aafaf6639750382d6dc665bc6a3c73d54e2076a8739a69c860a2314032eb9819060a001612dbd565b5f6001600160e01b031982167f49064906000000000000000000000000000000000000000000000000000000001480610d695750610d698261409c565b7f80bb2b638cc20bc4d0a60d66940f3ab4a00c1d7b313497ca82fb0b4ab00793008180613b1257506001600160a01b03831615155b15613c29575f613b2185611edf565b90506001600160a01b03841615801590613b4d5750836001600160a01b0316816001600160a01b031614155b8015613b9d57506001600160a01b038082165f9081527f80bb2b638cc20bc4d0a60d66940f3ab4a00c1d7b313497ca82fb0b4ab0079305602090815260408083209388168352929052205460ff16155b15613bdf576040517fa9fbf51f0000000000000000000000000000000000000000000000000000000081526001600160a01b0385166004820152602401611061565b8215613c275784866001600160a01b0316826001600160a01b03167f8c5be1e5ebec7d5bd14f71427d1e84f3dd0314c0f7b2291e5b200ac8c7c3b92560405160405180910390a45b505b5f93845260040160205250506040902080546001600160a01b0319166001600160a01b0392909216919091179055565b80471015613c9c576040517fcf47918100000000000000000000000000000000000000000000000000000000815247600482015260248101829052604401611061565b5f826001600160a01b0316826040515f6040518083038185875af1925050503d805f8114613ce5576040519150601f19603f3d011682016040523d82523d5f602084013e613cea565b606091505b5050905080610f8d5760405163d6bda27560e01b815260040160405180910390fd5b604080516001600160a01b0385811660248301528416604482015260648082018490528251808303909101815260849091019091526020810180517bffffffffffffffffffffffffffffffffffffffffffffffffffffffff167f23b872dd000000000000000000000000000000000000000000000000000000001790526110039085906140d9565b5f613da084848461415e565b949350505050565b336001600160a01b03831614801590613dca5750336001600160a01b03821614155b15610e74576040517f4793d28100000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b5f8281527f02dd7bc7dec4dceedda775e58dd541e08a116c6c53815c0bd028192f7b626800602090815260408083206001600160a01b038516845290915290205460ff16610e74576040517fe2517d3f0000000000000000000000000000000000000000000000000000000081526001600160a01b038216600482015260248101839052604401611061565b7fcd5ed15c6e187e77e9aee88184c21f4f2182ab5827cb3b7e07fbedcd63f033005460ff1661204c576040517f8dfc202b00000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b613ef182614267565b6040516001600160a01b038316907fbc7cd75a20ee27fd9adebab32041f755214dbc6bffa90cc0225b39da2e5c2d3b905f90a2805115613f3557610f8d82826142dd565b610e74614346565b7ff0c57e16840df040f15088dc2f81fe391c3923bec73e23a9662efc9c229c6a0054600160401b900460ff1661204c576040517fd7e6bcf800000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b613fa7613f3d565b7f80bb2b638cc20bc4d0a60d66940f3ab4a00c1d7b313497ca82fb0b4ab007930080613fd3848261524f565b5060018101611003838261524f565b613fea613f3d565b7fcd5ed15c6e187e77e9aee88184c21f4f2182ab5827cb3b7e07fbedcd63f03300805460ff19169055565b606061402082611edf565b505f61403660408051602081019091525f815290565b90505f8151116140545760405180602001604052805f8152506122cf565b8061405e8461437e565b60405160200161406f9291906151dd565b6040516020818303038152906040529392505050565b61408f838361441b565b610f8d335f858585613224565b5f6001600160e01b031982167f780e9d63000000000000000000000000000000000000000000000000000000001480610d695750610d6982614495565b5f8060205f8451602086015f885af1806140f8576040513d5f823e3d81fd5b50505f513d9150811561410f57806001141561411c565b6001600160a01b0384163b155b15611003576040517f5274afe70000000000000000000000000000000000000000000000000000000081526001600160a01b0385166004820152602401611061565b5f8061416b85858561452f565b90506001600160a01b03811661420557614200847f645e039705490088daad89bae25049a34f4a9072d398537b1ab2425f24cbed0280545f8381527f645e039705490088daad89bae25049a34f4a9072d398537b1ab2425f24cbed0360205260408120829055600182018355919091527fa42f15e5d656f8155fd7419d740a6073999f19cd6e061449ce4a257150545bf20155565b614228565b846001600160a01b0316816001600160a01b031614614228576142288185614669565b6001600160a01b0385166142445761423f84614713565b613da0565b846001600160a01b0316816001600160a01b031614613da057613da08585614806565b806001600160a01b03163b5f0361429c57604051634c9c8ce360e01b81526001600160a01b0382166004820152602401611061565b7f360894a13ba1a3210667c828492db98dca3e2076cc3735a920a3ca505d382bbc80546001600160a01b0319166001600160a01b0392909216919091179055565b60605f80846001600160a01b0316846040516142f9919061530b565b5f60405180830381855af49150503d805f8114614331576040519150601f19603f3d011682016040523d82523d5f602084013e614336565b606091505b50915091506136db858383614871565b341561204c576040517fb398979f00000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b60605f61438a836148e6565b60010190505f8167ffffffffffffffff8111156143a9576143a9614d48565b6040519080825280601f01601f1916602001820160405280156143d3576020820181803683370190505b5090508181016020015b5f19017f3031323334353637383961626364656600000000000000000000000000000000600a86061a8153600a85049450846143dd57509392505050565b6001600160a01b03821661444457604051633250574960e11b81525f6004820152602401611061565b5f61445083835f613d94565b90506001600160a01b03811615610f8d576040517f73c6ac6e0000000000000000000000000000000000000000000000000000000081525f6004820152602401611061565b5f6001600160e01b031982167f80ac58cd0000000000000000000000000000000000000000000000000000000014806144f757506001600160e01b031982167f5b5e139f00000000000000000000000000000000000000000000000000000000145b80610d6957507f01ffc9a7000000000000000000000000000000000000000000000000000000006001600160e01b0319831614610d69565b5f8281527f80bb2b638cc20bc4d0a60d66940f3ab4a00c1d7b313497ca82fb0b4ab007930260205260408120547f80bb2b638cc20bc4d0a60d66940f3ab4a00c1d7b313497ca82fb0b4ab0079300906001600160a01b039081169084161561459c5761459c8185876149c7565b6001600160a01b038116156145d8576145b75f865f80613add565b6001600160a01b0381165f908152600383016020526040902080545f190190555b6001600160a01b03861615614608576001600160a01b0386165f9081526003830160205260409020805460010190555b5f85815260028301602052604080822080546001600160a01b0319166001600160a01b038a811691821790925591518893918516917fddf252ad1be2c89b69c2b068fc378daa952ba7f163c4a11628f55a4df523b3ef91a495945050505050565b7f645e039705490088daad89bae25049a34f4a9072d398537b1ab2425f24cbed005f614694846113a1565b5f8481526001840160209081526040808320546001600160a01b038916845291869052909120919250908183146146ec575f838152602082815260408083205485845281842081905583526001870190915290208290555b5f948552600190930160209081526040808620869055928552929092528220919091555050565b7f645e039705490088daad89bae25049a34f4a9072d398537b1ab2425f24cbed02547f645e039705490088daad89bae25049a34f4a9072d398537b1ab2425f24cbed00905f9061476590600190615326565b5f848152600384016020526040812054600285018054939450909284908110614790576147906150c3565b905f5260205f2001549050808460020183815481106147b1576147b16150c3565b5f918252602080832090910192909255828152600386019091526040808220849055868252812055600284018054806147ec576147ec615339565b600190038181905f5260205f20015f905590555050505050565b7f645e039705490088daad89bae25049a34f4a9072d398537b1ab2425f24cbed005f6001614833856113a1565b61483d9190615326565b6001600160a01b039094165f9081526020838152604080832087845282528083208690559482526001909301909252502055565b6060826148865761488182614a44565b6122cf565b815115801561489d57506001600160a01b0384163b155b156148df576040517f9996b3150000000000000000000000000000000000000000000000000000000081526001600160a01b0385166004820152602401611061565b50806122cf565b5f807a184f03e93ff9f4daa797ed6e38ed64bf6a1f010000000000000000831061492e577a184f03e93ff9f4daa797ed6e38ed64bf6a1f010000000000000000830492506040015b6d04ee2d6d415b85acef8100000000831061495a576d04ee2d6d415b85acef8100000000830492506020015b662386f26fc10000831061497857662386f26fc10000830492506010015b6305f5e1008310614990576305f5e100830492506008015b61271083106149a457612710830492506004015b606483106149b6576064830492506002015b600a8310610d695760010192915050565b6149d2838383614a6d565b610f8d576001600160a01b038316614a0057604051637e27328960e01b815260048101829052602401611061565b6040517f177e802f0000000000000000000000000000000000000000000000000000000081526001600160a01b038316600482015260248101829052604401611061565b805115614a545780518082602001fd5b60405163d6bda27560e01b815260040160405180910390fd5b5f6001600160a01b03831615801590613da05750826001600160a01b0316846001600160a01b03161480614ae457506001600160a01b038085165f9081527f80bb2b638cc20bc4d0a60d66940f3ab4a00c1d7b313497ca82fb0b4ab0079305602090815260408083209387168352929052205460ff165b80613da05750505f9081527f80bb2b638cc20bc4d0a60d66940f3ab4a00c1d7b313497ca82fb0b4ab007930460205260409020546001600160a01b03908116911614919050565b5f60208284031215614b3b575f80fd5b5035919050565b6001600160e01b031981168114611110575f80fd5b5f60208284031215614b67575f80fd5b81356122cf81614b42565b5f5b83811015614b8c578181015183820152602001614b74565b50505f910152565b5f8151808452614bab816020860160208601614b72565b601f01601f19169290920160200192915050565b602081525f6122cf6020830184614b94565b6001600160a01b0381168114611110575f80fd5b5f8060408385031215614bf6575f80fd5b8235614c0181614bd1565b946020939093013593505050565b5f8060408385031215614c20575f80fd5b50508035926020909101359150565b5f805f60608486031215614c41575f80fd5b8335614c4c81614bd1565b92506020840135614c5c81614bd1565b929592945050506040919091013590565b803561ffff81168114614c7e575f80fd5b919050565b5f805f60608486031215614c95575f80fd5b83359250614ca560208501614c6d565b9150614cb360408501614c6d565b90509250925092565b5f8060408385031215614ccd575f80fd5b823591506020830135614cdf81614bd1565b809150509250929050565b634e487b7160e01b5f52602160045260245ffd5b6005811061111057634e487b7160e01b5f52602160045260245ffd5b60208101614d2783614cfe565b91905290565b5f60208284031215614d3d575f80fd5b81356122cf81614bd1565b634e487b7160e01b5f52604160045260245ffd5b5f67ffffffffffffffff80841115614d7657614d76614d48565b604051601f8501601f19908116603f01168101908282118183101715614d9e57614d9e614d48565b81604052809350858152868686011115614db6575f80fd5b858560208301375f602087830101525050509392505050565b5f82601f830112614dde575f80fd5b6122cf83833560208501614d5c565b5f8060408385031215614dfe575f80fd5b8235614e0981614bd1565b9150602083013567ffffffffffffffff811115614e24575f80fd5b614e3085828601614dcf565b9150509250929050565b5f805f8060808587031215614e4d575f80fd5b8435935060208501359250614e6460408601614c6d565b9150614e7260608601614c6d565b905092959194509250565b8015158114611110575f80fd5b5f8060408385031215614e9b575f80fd5b8235614ea681614bd1565b91506020830135614cdf81614e7d565b5f805f8060808587031215614ec9575f80fd5b8435614ed481614bd1565b93506020850135614ee481614bd1565b925060408501359150606085013567ffffffffffffffff811115614f06575f80fd5b614f1287828801614dcf565b91505092959194509250565b6101008101614f2c8a614cfe565b98815260208101979097526001600160a01b0395861660408801529390941660608601529015156080850152151560a084015263ffffffff91821660c08401521660e09091015290565b5f805f60608486031215614f88575f80fd5b8335614f9381614bd1565b92506020840135614fa381614bd1565b91506040840135614fb381614bd1565b809150509250925092565b5f805f805f805f60e0888a031215614fd4575f80fd5b8735614fdf81614bd1565b9650602088013567ffffffffffffffff811115614ffa575f80fd5b8801601f81018a1361500a575f80fd5b6150198a823560208401614d5c565b9650506040880135945060608801359350608088013561503881614bd1565b925060a0880135915060c088013561504f81614e7d565b8091505092959891949750929550565b5f8060408385031215615070575f80fd5b823561507b81614bd1565b91506020830135614cdf81614bd1565b600181811c9082168061509f57607f821691505b6020821081036150bd57634e487b7160e01b5f52602260045260245ffd5b50919050565b634e487b7160e01b5f52603260045260245ffd5b5f602082840312156150e7575f80fd5b81516122cf81614e7d565b634e487b7160e01b5f52601160045260245ffd5b80820180821115610d6957610d696150f2565b5f6001820161512a5761512a6150f2565b5060010190565b8281526040810161514183614cfe565b8260208301529392505050565b5f6020828403121561515e575f80fd5b5051919050565b5f63ffffffff80831681810361517d5761517d6150f2565b6001019392505050565b5f6001600160a01b038087168352808616602084015250836040830152608060608301526151b86080830184614b94565b9695505050505050565b5f602082840312156151d2575f80fd5b81516122cf81614b42565b5f83516151ee818460208801614b72565b835190830190615202818360208801614b72565b01949350505050565b601f821115610f8d57805f5260205f20601f840160051c810160208510156152305750805b601f840160051c820191505b818110156117a6575f815560010161523c565b815167ffffffffffffffff81111561526957615269614d48565b61527d81615277845461508b565b8461520b565b602080601f8311600181146152b0575f84156152995750858301515b5f19600386901b1c1916600185901b178555610fb8565b5f85815260208120601f198616915b828110156152de578886015182559484019460019091019084016152bf565b50858210156152fb57878501515f19600388901b60f8161c191681555b5050505050600190811b01905550565b5f825161531c818460208701614b72565b9190910192915050565b81810381811115610d6957610d696150f2565b634e487b7160e01b5f52603160045260245ffdfe54347f9bbfbc2e7b5786abab9693d2ba67834fa4787a7740d72712e87e207400a2646970667358221220bd32c850b9916a0486802cdad4aefcbe63bff6d5f0af229ca710deb07bc9fba564736f6c63430008180033",
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

// BookingtokenManagerAddressUpdatedIterator is returned from FilterManagerAddressUpdated and is used to iterate over the raw logs and unpacked data for ManagerAddressUpdated events raised by the Bookingtoken contract.
type BookingtokenManagerAddressUpdatedIterator struct {
	Event *BookingtokenManagerAddressUpdated // Event containing the contract specifics and raw log

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
func (it *BookingtokenManagerAddressUpdatedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(BookingtokenManagerAddressUpdated)
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
		it.Event = new(BookingtokenManagerAddressUpdated)
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
func (it *BookingtokenManagerAddressUpdatedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *BookingtokenManagerAddressUpdatedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// BookingtokenManagerAddressUpdated represents a ManagerAddressUpdated event raised by the Bookingtoken contract.
type BookingtokenManagerAddressUpdated struct {
	OldManager common.Address
	NewManager common.Address
	Raw        types.Log // Blockchain specific contextual infos
}

// FilterManagerAddressUpdated is a free log retrieval operation binding the contract event 0x9462e60b9d7b78dcca266b08b885d2cd87178de9a5c63e600065b86e530f0b9b.
//
// Solidity: event ManagerAddressUpdated(address indexed oldManager, address indexed newManager)
func (_Bookingtoken *BookingtokenFilterer) FilterManagerAddressUpdated(opts *bind.FilterOpts, oldManager []common.Address, newManager []common.Address) (*BookingtokenManagerAddressUpdatedIterator, error) {

	var oldManagerRule []interface{}
	for _, oldManagerItem := range oldManager {
		oldManagerRule = append(oldManagerRule, oldManagerItem)
	}
	var newManagerRule []interface{}
	for _, newManagerItem := range newManager {
		newManagerRule = append(newManagerRule, newManagerItem)
	}

	logs, sub, err := _Bookingtoken.contract.FilterLogs(opts, "ManagerAddressUpdated", oldManagerRule, newManagerRule)
	if err != nil {
		return nil, err
	}
	return &BookingtokenManagerAddressUpdatedIterator{contract: _Bookingtoken.contract, event: "ManagerAddressUpdated", logs: logs, sub: sub}, nil
}

// WatchManagerAddressUpdated is a free log subscription operation binding the contract event 0x9462e60b9d7b78dcca266b08b885d2cd87178de9a5c63e600065b86e530f0b9b.
//
// Solidity: event ManagerAddressUpdated(address indexed oldManager, address indexed newManager)
func (_Bookingtoken *BookingtokenFilterer) WatchManagerAddressUpdated(opts *bind.WatchOpts, sink chan<- *BookingtokenManagerAddressUpdated, oldManager []common.Address, newManager []common.Address) (event.Subscription, error) {

	var oldManagerRule []interface{}
	for _, oldManagerItem := range oldManager {
		oldManagerRule = append(oldManagerRule, oldManagerItem)
	}
	var newManagerRule []interface{}
	for _, newManagerItem := range newManager {
		newManagerRule = append(newManagerRule, newManagerItem)
	}

	logs, sub, err := _Bookingtoken.contract.WatchLogs(opts, "ManagerAddressUpdated", oldManagerRule, newManagerRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(BookingtokenManagerAddressUpdated)
				if err := _Bookingtoken.contract.UnpackLog(event, "ManagerAddressUpdated", log); err != nil {
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

// ParseManagerAddressUpdated is a log parse operation binding the contract event 0x9462e60b9d7b78dcca266b08b885d2cd87178de9a5c63e600065b86e530f0b9b.
//
// Solidity: event ManagerAddressUpdated(address indexed oldManager, address indexed newManager)
func (_Bookingtoken *BookingtokenFilterer) ParseManagerAddressUpdated(log types.Log) (*BookingtokenManagerAddressUpdated, error) {
	event := new(BookingtokenManagerAddressUpdated)
	if err := _Bookingtoken.contract.UnpackLog(event, "ManagerAddressUpdated", log); err != nil {
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

// BookingtokenMinExpirationTimestampDiffUpdatedIterator is returned from FilterMinExpirationTimestampDiffUpdated and is used to iterate over the raw logs and unpacked data for MinExpirationTimestampDiffUpdated events raised by the Bookingtoken contract.
type BookingtokenMinExpirationTimestampDiffUpdatedIterator struct {
	Event *BookingtokenMinExpirationTimestampDiffUpdated // Event containing the contract specifics and raw log

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
func (it *BookingtokenMinExpirationTimestampDiffUpdatedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(BookingtokenMinExpirationTimestampDiffUpdated)
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
		it.Event = new(BookingtokenMinExpirationTimestampDiffUpdated)
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
func (it *BookingtokenMinExpirationTimestampDiffUpdatedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *BookingtokenMinExpirationTimestampDiffUpdatedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// BookingtokenMinExpirationTimestampDiffUpdated represents a MinExpirationTimestampDiffUpdated event raised by the Bookingtoken contract.
type BookingtokenMinExpirationTimestampDiffUpdated struct {
	OldDiff *big.Int
	NewDiff *big.Int
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterMinExpirationTimestampDiffUpdated is a free log retrieval operation binding the contract event 0x6175237436049150327545b616ef840ddeb5def6cd197617a415f10fc838fdb5.
//
// Solidity: event MinExpirationTimestampDiffUpdated(uint256 oldDiff, uint256 newDiff)
func (_Bookingtoken *BookingtokenFilterer) FilterMinExpirationTimestampDiffUpdated(opts *bind.FilterOpts) (*BookingtokenMinExpirationTimestampDiffUpdatedIterator, error) {

	logs, sub, err := _Bookingtoken.contract.FilterLogs(opts, "MinExpirationTimestampDiffUpdated")
	if err != nil {
		return nil, err
	}
	return &BookingtokenMinExpirationTimestampDiffUpdatedIterator{contract: _Bookingtoken.contract, event: "MinExpirationTimestampDiffUpdated", logs: logs, sub: sub}, nil
}

// WatchMinExpirationTimestampDiffUpdated is a free log subscription operation binding the contract event 0x6175237436049150327545b616ef840ddeb5def6cd197617a415f10fc838fdb5.
//
// Solidity: event MinExpirationTimestampDiffUpdated(uint256 oldDiff, uint256 newDiff)
func (_Bookingtoken *BookingtokenFilterer) WatchMinExpirationTimestampDiffUpdated(opts *bind.WatchOpts, sink chan<- *BookingtokenMinExpirationTimestampDiffUpdated) (event.Subscription, error) {

	logs, sub, err := _Bookingtoken.contract.WatchLogs(opts, "MinExpirationTimestampDiffUpdated")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(BookingtokenMinExpirationTimestampDiffUpdated)
				if err := _Bookingtoken.contract.UnpackLog(event, "MinExpirationTimestampDiffUpdated", log); err != nil {
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

// ParseMinExpirationTimestampDiffUpdated is a log parse operation binding the contract event 0x6175237436049150327545b616ef840ddeb5def6cd197617a415f10fc838fdb5.
//
// Solidity: event MinExpirationTimestampDiffUpdated(uint256 oldDiff, uint256 newDiff)
func (_Bookingtoken *BookingtokenFilterer) ParseMinExpirationTimestampDiffUpdated(log types.Log) (*BookingtokenMinExpirationTimestampDiffUpdated, error) {
	event := new(BookingtokenMinExpirationTimestampDiffUpdated)
	if err := _Bookingtoken.contract.UnpackLog(event, "MinExpirationTimestampDiffUpdated", log); err != nil {
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
