// SPDX-License-Identifier: LGPL-3.0-or-later
//
// Travel Token Messenger Account Implementation

pragma solidity 0.8.24;

import { Initializable } from "@openzeppelin/contracts-upgradeable/proxy/utils/Initializable.sol";
import { UUPSUpgradeable } from "@openzeppelin/contracts-upgradeable/proxy/utils/UUPSUpgradeable.sol";
import { AccessControlEnumerableUpgradeable } from "@openzeppelin/contracts-upgradeable/access/extensions/AccessControlEnumerableUpgradeable.sol";
import { ReentrancyGuardUpgradeable } from "@openzeppelin/contracts-upgradeable/utils/ReentrancyGuardUpgradeable.sol";
import { Address } from "@openzeppelin/contracts/utils/Address.sol";
import { IERC20 } from "@openzeppelin/contracts/token/ERC20/IERC20.sol";
import { SafeERC20 } from "@openzeppelin/contracts/token/ERC20/utils/SafeERC20.sol";
import { IERC721 } from "@openzeppelin/contracts/token/ERC721/IERC721.sol";
import { IERC721Receiver } from "@openzeppelin/contracts/token/ERC721/IERC721Receiver.sol";
import { ERC1967Utils } from "@openzeppelin/contracts/proxy/ERC1967/ERC1967Utils.sol";

import { ITTMAccountManager } from "../manager/ITTMAccountManager.sol";
import { BookingTokenOperator } from "../booking-token/BookingTokenOperator.sol";
import { PartnerConfiguration } from "../partner/PartnerConfiguration.sol";
import { GasMoneyManager } from "./GasMoneyManager.sol";

/**
 * @title Travel Token Messenger Account
 * @notice A TTM Account manages funds, minting/buying of booking tokens, provided
 * or wanted services, and multiple bots for distributors and suppliers on
 * Travel Token Messenger ecosystem.
 *
 * Registering bots is done by role based access control. Bot's with
 * `MESSENGER_BOT_ROLE` are authorized to represent the TTMAccount.
 * Bot can also have `GAS_WITHDRAWER_ROLE` and `BOOKING_OPERATOR_ROLE`.
 *
 * `GAS_WITHDRAWER_ROLE` enables a bot to withdraw native coins (ETH) from the
 * contract to be used as gas money. This is restricted with a `limit` (wei)
 * and `period` (seconds) set by the `BOT_ADMIN_ROLE`. The limit and period
 * apply per bot address: each bot tracks its own withdrawals against the
 * same limit, independently of every other bot on the account. Default
 * starting values are 10 ETH per 24 hours.
 *
 * `BOOKING_OPERATOR_ROLE` enables a bot to mint and buy Booking Tokens by
 * calling the corresponding functions on the {BookingToken} contract. The buy
 * operation pays the price of the Booking Token with the funds on the
 * {TTMAccount} contract.
 *
 * @dev This contract uses UUPS style upgradeability. The authorization function
 * `_authorizeUpgrade(address)` can be called by the `UPGRADER_ROLE` and is
 * restricted to only upgrade to the implementation address registered on the
 * {TTMAccountManager} contract.
 */
contract TTMAccount is
    Initializable,
    AccessControlEnumerableUpgradeable,
    UUPSUpgradeable,
    ReentrancyGuardUpgradeable,
    IERC721Receiver,
    PartnerConfiguration,
    GasMoneyManager
{
    using Address for address payable;
    using SafeERC20 for IERC20;

    /***************************************************
     *                    ROLES                        *
     ***************************************************/

    /**
     * @notice Upgrader role can upgrade the contract to a new implementation.
     */
    bytes32 public constant UPGRADER_ROLE = keccak256("UPGRADER_ROLE");

    /**
     * @notice Bot admin role can add & remove bots and set gas money withdrawal
     * parameters.
     */
    bytes32 public constant BOT_ADMIN_ROLE = keccak256("BOT_ADMIN_ROLE");

    /**
     * @notice Messenger bot role can interact on behalf of this TTMAccount
     * contract.
     */
    bytes32 public constant MESSENGER_BOT_ROLE = keccak256("MESSENGER_BOT_ROLE");

    /**
     * @notice Gas withdrawer role can withdraw gas money from the contract. This is
     * intended to be used by the bots and is granted when `addMessengerBot` is
     * called.
     */
    bytes32 public constant GAS_WITHDRAWER_ROLE = keccak256("GAS_WITHDRAWER_ROLE");

    /**
     * @notice Withdrawer role can withdraw funds from the contract.
     */
    bytes32 public constant WITHDRAWER_ROLE = keccak256("WITHDRAWER_ROLE");

    /**
     * @notice Booking operator role can mint and buy booking tokens using the
     * functions on this contract. This is generally used by the bots. The
     * price for the booking token is paid by this contract.
     */
    bytes32 public constant BOOKING_OPERATOR_ROLE = keccak256("BOOKING_OPERATOR_ROLE");

    /**
     * @notice Service admin role can add & remove supported & wanted services.
     */
    bytes32 public constant SERVICE_ADMIN_ROLE = keccak256("SERVICE_ADMIN_ROLE");

    /***************************************************
     *                   STORAGE                       *
     ***************************************************/

    /// @custom:storage-location erc7201:traveltoken.messenger.storage.TTMAccount
    struct TTMAccountStorage {
        /**
         * @dev Address of the TTMAccountManager
         */
        address _manager;
        /**
         * @dev Address of the BookingToken contract
         */
        address _bookingToken;
    }

    // keccak256(abi.encode(uint256(keccak256("traveltoken.messenger.storage.TTMAccount")) - 1)) & ~bytes32(uint256(0xff));
    bytes32 private constant TTMAccountStorageLocation =
        0x17fbd8e22750b6dc617ceae09f600bf2810b53ccc1f790a9a4f78a3f7169b900;

    function _getTTMAccountStorage() private pure returns (TTMAccountStorage storage $) {
        assembly {
            $.slot := TTMAccountStorageLocation
        }
    }

    /***************************************************
     *                    EVENTS                       *
     ***************************************************/

    /**
     * @notice TTMAccount upgrade event. Emitted when the TTMAccount implementation is upgraded.
     */
    event TTMAccountUpgraded(address indexed oldImplementation, address indexed newImplementation);

    /**
     * @notice Deposit event, emitted when there is a new deposit
     */
    event Deposit(address indexed sender, uint256 amount);

    /**
     * @notice Withdraw event, emitted when there is a new withdrawal
     */
    event Withdraw(address indexed receiver, uint256 amount);

    /**
     * @notice Messenger bot added
     */
    event MessengerBotAdded(address indexed bot);

    /**
     * @notice Messenger bot removed
     */
    event MessengerBotRemoved(address indexed bot);

    // Partner Config Events

    /**
     * @dev Service events carry the service hash only. Indexing a dynamic `string`
     * stores just its keccak hash in the topic and nothing in the data section, so the
     * old `string indexed serviceName` form published a hash while pretending to
     * publish a name. Consumers resolve names from `ServiceRegistry`'s
     * `ServiceRegistered` / `ServiceUnregistered` events, which do carry them.
     *
     * Capability strings stay readable: capabilities are free-form partner text with
     * no registry to resolve against.
     */
    event ServiceAdded(bytes32 indexed serviceHash);
    event ServiceRemoved(bytes32 indexed serviceHash);

    event WantedServiceAdded(string indexed serviceName);
    event WantedServiceRemoved(string indexed serviceName);

    event ServiceRestrictedRateUpdated(bytes32 indexed serviceHash, bool restrictedRate);

    event ServiceCapabilitiesUpdated(bytes32 indexed serviceHash);
    event ServiceCapabilityAdded(bytes32 indexed serviceHash, string capability);
    event ServiceCapabilityRemoved(bytes32 indexed serviceHash, string capability);

    /***************************************************
     *                    ERRORS                       *
     ***************************************************/

    /**
     * @notice TTMAccount implementation address does not match the one in the manager
     */
    error TTMAccountImplementationMismatch(address latestImplementation, address newImplementation);

    /**
     * @notice New implementation is the same as the current implementation, no update needed
     */
    error TTMAccountNoUpgradeNeeded(address oldImplementation, address newImplementation);

    /**
     * @notice Error to revert if transfer to zero address
     */
    error TransferToZeroAddress();

    /**
     * @notice A required address parameter was the zero address.
     */
    error ZeroAddress();

    /**
     * @notice The given service hash is not registered in the manager's ServiceRegistry.
     *
     * @dev Same selector as ServiceRegistry's `ServiceNotRegistered()` (identical, argument-less
     * signature) since this error is what actually bubbles up from the staticcall in
     * {_requireRegisteredService}; declaring it here as well only lets this contract's ABI
     * name it directly.
     */
    error ServiceNotRegistered();

    /***************************************************
     *         CONSTRUCTOR & INITIALIZATION            *
     ***************************************************/

    /// @custom:oz-upgrades-unsafe-allow constructor
    constructor() {
        _disableInitializers();
    }

    function initialize(
        address manager,
        address bookingToken,
        address defaultAdmin,
        address upgrader
    ) public initializer {
        if (
            manager == address(0) || bookingToken == address(0) || defaultAdmin == address(0) || upgrader == address(0)
        ) {
            revert ZeroAddress();
        }

        __AccessControl_init();
        __UUPSUpgradeable_init();
        __ReentrancyGuard_init();
        __PartnerConfiguration_init();

        _grantRole(DEFAULT_ADMIN_ROLE, defaultAdmin);
        _grantRole(SERVICE_ADMIN_ROLE, defaultAdmin);
        _grantRole(BOT_ADMIN_ROLE, defaultAdmin);
        _grantRole(UPGRADER_ROLE, upgrader);

        TTMAccountStorage storage $ = _getTTMAccountStorage();

        $._manager = manager;
        $._bookingToken = bookingToken;

        // Initialize GasMoneyManager
        uint256 withdrawalLimit = 10 ether; // 10 ETH
        uint256 withdrawalPeriod = 24 hours; // per 24 hours
        __GasMoneyManager_init(withdrawalLimit, withdrawalPeriod);
    }

    receive() external payable {
        emit Deposit(msg.sender, msg.value);
    }

    /***************************************************
     *                    Getters                      *
     ***************************************************/

    /**
     * @notice Returns the TTMAccountManager address.
     *
     * @return TTMAccountManager address
     */
    function getManagerAddress() public view returns (address) {
        TTMAccountStorage storage $ = _getTTMAccountStorage();
        return $._manager;
    }

    /**
     * @notice Returns the booking token address.
     *
     * @return BookingToken address
     */
    function getBookingTokenAddress() public view returns (address) {
        TTMAccountStorage storage $ = _getTTMAccountStorage();
        return $._bookingToken;
    }

    /***************************************************
     *                    Account                      *
     ***************************************************/

    /**
     * @notice Authorizes the upgrade of the TTMAccount.
     *
     * Reverts if the new implementation is the same as the old one.
     *
     * Reverts if the new implementation does not match the implementation address
     * in the manager. Only implementations registered at the manager are allowed.
     *
     * @dev Emits a {TTMAccountUpgraded} event.
     *
     * @param newImplementation The new implementation address
     */
    function _authorizeUpgrade(address newImplementation) internal override onlyRole(UPGRADER_ROLE) {
        // Get the implementation address from the manager
        address managerImplementation = ITTMAccountManager(getManagerAddress()).getAccountImplementation();
        address oldImplementation = ERC1967Utils.getImplementation();

        // Revert if the new implementation is the same as the old one
        if (oldImplementation == newImplementation) {
            revert TTMAccountNoUpgradeNeeded(oldImplementation, newImplementation);
        }

        // Check if new implementation matches the implementation address in the manager
        if (newImplementation != managerImplementation) {
            revert TTMAccountImplementationMismatch(managerImplementation, newImplementation);
        }

        emit TTMAccountUpgraded(oldImplementation, newImplementation);
    }

    /**
     * @notice Returns true if an address is an authorized messenger bot
     *
     * @param bot The bot's address
     */
    function isBotAllowed(address bot) public view returns (bool) {
        return hasRole(MESSENGER_BOT_ROLE, bot);
    }

    /**
     * @notice Withdraw ETH from the TTMAccount
     *
     * @param recipient The recipient of the withdrawal
     * @param amount The amount to withdraw
     */
    function withdraw(address payable recipient, uint256 amount) external nonReentrant onlyRole(WITHDRAWER_ROLE) {
        if (recipient == address(0)) {
            revert TransferToZeroAddress();
        }
        recipient.sendValue(amount);
        emit Withdraw(recipient, amount);
    }

    /***************************************************
     *                 BOOKING TOKEN                   *
     ***************************************************/

    /**
     * @notice Mints booking token.
     *
     * @param reservedFor The account to reserve the token for
     * @param uri The URI of the token
     * @param expirationTimestamp The expiration timestamp
     * @param price The price of the token
     * @param paymentToken The payment token, if address(0) then native
     * @param offchainPaymentCurrency The offchain payment currency
     * @param cancellable If the token is cancellable
     */
    function mintBookingToken(
        address reservedFor,
        string memory uri,
        uint256 expirationTimestamp,
        uint256 price,
        IERC20 paymentToken,
        uint256 offchainPaymentCurrency,
        bool cancellable
    ) external onlyRole(BOOKING_OPERATOR_ROLE) {
        BookingTokenOperator.mintBookingToken(
            getBookingTokenAddress(),
            reservedFor,
            uri,
            expirationTimestamp,
            price,
            paymentToken,
            offchainPaymentCurrency,
            cancellable
        );
    }

    /**
     * @notice Buys booking token.
     *
     * @param tokenId The token id
     */
    function buyBookingToken(
        uint256 tokenId,
        uint256 expectedPrice,
        IERC20 expectedPaymentToken
    ) external nonReentrant onlyRole(BOOKING_OPERATOR_ROLE) {
        BookingTokenOperator.buyBookingToken(getBookingTokenAddress(), tokenId, expectedPrice, expectedPaymentToken);
    }

    /**
     * @notice Record expiration status if the token is expired
     */
    function recordExpiration(uint256 tokenId) external onlyRole(BOOKING_OPERATOR_ROLE) {
        BookingTokenOperator.recordExpiration(getBookingTokenAddress(), tokenId);
    }

    /**
     * @notice Always returns `IERC721Receiver.onERC721Received.selector`.
     *
     * @dev See {IERC721Receiver-onERC721Received}.
     */
    function onERC721Received(address, address, uint256, bytes memory) public virtual returns (bytes4) {
        return this.onERC721Received.selector;
    }

    /***************************************************
     *                ERC20 & ERC721                   *
     ***************************************************/

    /**
     * @notice Transfers ERC20 tokens.
     *
     * This function reverts if `to` is the zero address.
     *
     * @param token The ERC20 token
     * @param to The address to transfer the tokens to
     * @param amount The amount of tokens to transfer
     */
    function transferERC20(IERC20 token, address to, uint256 amount) external onlyRole(WITHDRAWER_ROLE) {
        if (to == address(0)) {
            revert TransferToZeroAddress();
        }
        token.safeTransfer(to, amount);
    }

    /**
     * @notice Transfers ERC721 tokens.
     *
     * This function reverts if `to` is the zero address.
     *
     * @param token The ERC721 token
     * @param to The address to transfer the tokens to
     * @param tokenId The token id of the token
     */
    function transferERC721(IERC721 token, address to, uint256 tokenId) external onlyRole(WITHDRAWER_ROLE) {
        if (to == address(0)) {
            revert TransferToZeroAddress();
        }
        token.safeTransferFrom(address(this), to, tokenId);
    }

    /***************************************************
     *                PARTNER CONFIG                   *
     ***************************************************/

    /**
     * @notice Adds a service to the account as a supported service.
     *
     * `serviceHash` is `keccak256(abi.encodePacked(serviceName))`, where the name is
     * pkg + service name as defined in the Travel Token Messenger Protocol's protobuf
     * definitions. For example:
     *
     * ```text
     *  ┌────────────── pkg ─────────────┐ ┌───── service name ─────┐
     * "ttm.services.accommodation.v1alpha.AccommodationSearchService")
     * ```
     *
     * @dev The hash must be registered in the manager's `ServiceRegistry`. That check is
     * the one manager staticcall left on this path: it is a write, called rarely, and
     * without it an account could advertise a service that does not exist. Reads carry
     * no manager dependency at all.
     *
     * @param serviceHash Hash of the service name to support
     * @param restrictedRate Whether the service is restricted to pre-agreement
     * @param capabilities Capabilities of the service (optional)
     */
    function addService(
        bytes32 serviceHash,
        bool restrictedRate,
        string[] memory capabilities
    ) public onlyRole(SERVICE_ADMIN_ROLE) {
        _requireRegisteredService(serviceHash);
        _addService(serviceHash, capabilities, restrictedRate);
        emit ServiceAdded(serviceHash);
    }

    /**
     * @notice Reverts unless `serviceHash` is registered in the manager's ServiceRegistry.
     */
    function _requireRegisteredService(bytes32 serviceHash) private view {
        // Reverts with ServiceNotRegistered if the hash is unknown to the registry.
        ITTMAccountManager(getManagerAddress()).getRegisteredServiceNameByHash(serviceHash);
    }

    /**
     * @notice Removes a service from the account by its hash.
     */
    function removeService(bytes32 serviceHash) public onlyRole(SERVICE_ADMIN_ROLE) {
        _removeService(serviceHash);
        emit ServiceRemoved(serviceHash);
    }

    /**
     * @notice Removes all supported services from the account.
     */
    function removeAllServices() public onlyRole(SERVICE_ADMIN_ROLE) {
        bytes32[] memory serviceHashes = getAllServiceHashes();

        for (uint256 i = 0; i < serviceHashes.length; i++) {
            _removeService(serviceHashes[i]);
            emit ServiceRemoved(serviceHashes[i]);
        }
    }

    /**
     * @notice Sets whether a service is offered at a restricted (non-rack) rate.
     */
    function setServiceRestrictedRate(bytes32 serviceHash, bool restrictedRate) public onlyRole(SERVICE_ADMIN_ROLE) {
        _setServiceRestrictedRate(serviceHash, restrictedRate);
        emit ServiceRestrictedRateUpdated(serviceHash, restrictedRate);
    }

    /**
     * @notice Replaces the capability list of a service.
     */
    function setServiceCapabilities(
        bytes32 serviceHash,
        string[] memory capabilities
    ) public onlyRole(SERVICE_ADMIN_ROLE) {
        _setServiceCapabilities(serviceHash, capabilities);
        emit ServiceCapabilitiesUpdated(serviceHash);
    }

    /**
     * @notice Adds a single capability to a service.
     */
    function addServiceCapability(bytes32 serviceHash, string memory capability) public onlyRole(SERVICE_ADMIN_ROLE) {
        _addServiceCapability(serviceHash, capability);
        emit ServiceCapabilityAdded(serviceHash, capability);
    }

    /**
     * @notice Removes a single capability from a service.
     */
    function removeServiceCapability(
        bytes32 serviceHash,
        string memory capability
    ) public onlyRole(SERVICE_ADMIN_ROLE) {
        _removeServiceCapability(serviceHash, capability);
        emit ServiceCapabilityRemoved(serviceHash, capability);
    }

    /**
     * @notice Get service hash by name. Returns the keccak256 hash of the
     * registered service name from the account manager
     */
    function getRegisteredServiceHash(string memory serviceName) private view returns (bytes32 serviceHash) {
        return ITTMAccountManager(getManagerAddress()).getRegisteredServiceHashByName(serviceName);
    }

    /**
     * @notice Get service hash by name. Returns the keccak256 hash of the service name
     * from the account manager
     */
    function getServiceHash(string memory serviceName) private view returns (bytes32 serviceHash) {
        return ITTMAccountManager(getManagerAddress()).getServiceHashByName(serviceName);
    }

    /**
     * @notice Get service name by hash. Returns the service name from the account manager
     */
    function getServiceName(bytes32 serviceHash) private view returns (string memory serviceName) {
        return ITTMAccountManager(getManagerAddress()).getServiceNameByHash(serviceHash);
    }

    /***************************************************
     *           SERVICES WITH RESOLVED NAMES          *
     ***************************************************/

    /**
     * @notice Get all supported services. Return a list of service names and a list of service objects.
     */
    function getSupportedServices() public view returns (string[] memory serviceNames, Service[] memory services) {
        // Get all hashes and create a list with predefined length
        bytes32[] memory _serviceHashes = getAllServiceHashes();
        string[] memory _serviceNames = new string[](_serviceHashes.length);
        Service[] memory _allSupportedServicesList = new Service[](_serviceHashes.length);

        for (uint256 i = 0; i < _serviceHashes.length; i++) {
            _serviceNames[i] = getServiceName(_serviceHashes[i]);
            _allSupportedServicesList[i] = getService(_serviceHashes[i]);
        }

        return (_serviceNames, _allSupportedServicesList);
    }

    /**
     * @notice Check if a service is registered and supported.
     *
     * @param serviceName Service name to check
     */
    function isServiceSupported(string memory serviceName) public view returns (bool) {
        return _isServiceSupported(getServiceHash(serviceName));
    }

    /**
     * @notice Get service restricted rate by name. Overloading the getServiceRestrictedRate function.
     */
    function getServiceRestrictedRate(string memory serviceName) public view returns (bool restrictedRate) {
        return getServiceRestrictedRate(getServiceHash(serviceName));
    }

    /**
     * @notice Get service capabilities by name. Overloading the getServiceCapabilities function.
     */
    function getServiceCapabilities(string memory serviceName) public view returns (string[] memory capabilities) {
        return getServiceCapabilities(getServiceHash(serviceName));
    }

    /***************************************************
     *                WANTED SERVICES                  *
     ***************************************************/

    /**
     * @notice Adds wanted services.
     *
     * @param serviceNames List of service names
     */
    function addWantedServices(string[] memory serviceNames) public onlyRole(SERVICE_ADMIN_ROLE) {
        for (uint256 i = 0; i < serviceNames.length; i++) {
            bytes32 serviceHash = getRegisteredServiceHash(serviceNames[i]);
            _addWantedService(serviceHash);
            emit WantedServiceAdded(serviceNames[i]);
        }
    }

    /**
     * @notice Removes wanted services.
     *
     * @param serviceNames List of service names
     */
    function removeWantedServices(string[] memory serviceNames) public onlyRole(SERVICE_ADMIN_ROLE) {
        for (uint256 i = 0; i < serviceNames.length; i++) {
            bytes32 serviceHash = getServiceHash(serviceNames[i]);
            _removeWantedService(serviceHash);
            emit WantedServiceRemoved(serviceNames[i]);
        }
    }

    /**
     * @notice Get all wanted services.
     *
     * @return serviceNames List of service names
     */
    function getWantedServices() public view returns (string[] memory serviceNames) {
        bytes32[] memory _wantedServiceHashes = getWantedServiceHashes();

        string[] memory _wantedServiceNames = new string[](_wantedServiceHashes.length);

        for (uint256 i = 0; i < _wantedServiceHashes.length; i++) {
            _wantedServiceNames[i] = getServiceName(_wantedServiceHashes[i]);
        }

        return _wantedServiceNames;
    }

    /***************************************************
     *                   PAYMENT                       *
     ***************************************************/

    /**
     * @notice Sets if off-chain payment is supported.
     *
     * @param _isSupported true if off-chain payment is supported
     */
    function setOffChainPaymentSupported(bool _isSupported) public onlyRole(SERVICE_ADMIN_ROLE) {
        _setOffChainPaymentSupported(_isSupported);
    }

    /**
     * @notice Adds a supported payment token.
     *
     * @param _supportedToken address of the token
     */
    function addSupportedToken(address _supportedToken) public onlyRole(SERVICE_ADMIN_ROLE) {
        _addSupportedToken(_supportedToken);
    }

    /**
     * @notice Removes a supported payment token.
     *
     * @param _supportedToken address of the token
     */
    function removeSupportedToken(address _supportedToken) public onlyRole(SERVICE_ADMIN_ROLE) {
        _removeSupportedToken(_supportedToken);
    }

    /***************************************************
     *                  PUBLIC KEY                     *
     ***************************************************/

    /**
     * @notice Add public key with address
     *
     * These public keys are intended to be used with for off-chain encryption of private booking data.
     *
     * @param pubKeyAddress address of the public key
     * @param data public key data
     */
    function addPublicKey(address pubKeyAddress, bytes memory data) public onlyRole(SERVICE_ADMIN_ROLE) {
        _addPublicKey(pubKeyAddress, data);
    }

    /**
     * @notice Remove public key by address
     */
    function removePublicKey(address pubKeyAddress) public onlyRole(SERVICE_ADMIN_ROLE) {
        _removePublicKey(pubKeyAddress);
    }

    /***************************************************
     *                MESSENGER BOTS                   *
     ***************************************************/

    /**
     * @notice Adds messenger bot with initial gas money. The amount of `gasMoney`
     * need to be present in the contract.
     */
    function addMessengerBot(address bot, uint256 gasMoney) public onlyRole(BOT_ADMIN_ROLE) {
        // Check if bot is valid to prevent accidental transfers of funds to zero address
        if (bot == address(0)) revert TransferToZeroAddress();

        // Grant roles to bot
        _grantRole(MESSENGER_BOT_ROLE, bot);
        _grantRole(BOOKING_OPERATOR_ROLE, bot);
        _grantRole(GAS_WITHDRAWER_ROLE, bot);

        emit MessengerBotAdded(bot);

        // Send gasMoney to bot
        payable(bot).sendValue(gasMoney);
    }

    /**
     * @notice Removes messenger bot by revoking the roles.
     */
    function removeMessengerBot(address bot) public onlyRole(BOT_ADMIN_ROLE) {
        _revokeRole(MESSENGER_BOT_ROLE, bot);
        _revokeRole(BOOKING_OPERATOR_ROLE, bot);
        _revokeRole(GAS_WITHDRAWER_ROLE, bot);

        emit MessengerBotRemoved(bot);
    }

    /***************************************************
     *              GAS MONEY WITHDRAW                 *
     ***************************************************/

    /**
     * @notice Withdraw gas money. Requires the `GAS_WITHDRAWER_ROLE`.
     *
     * @param amount The amount to withdraw in wei
     */
    function withdrawGasMoney(uint256 amount) public nonReentrant onlyRole(GAS_WITHDRAWER_ROLE) {
        _withdrawGasMoney(amount);
    }

    /**
     * @notice Set gas money withdrawal parameters. Requires the `BOT_ADMIN_ROLE`.
     *
     * @param limit Amount of gas money to withdraw in wei per period
     * @param period Duration of the withdrawal period in seconds
     */
    function setGasMoneyWithdrawal(uint256 limit, uint256 period) public onlyRole(BOT_ADMIN_ROLE) {
        _setGasMoneyWithdrawal(limit, period);
    }

    /***************************************************
     *                 CANCELLATION                    *
     ***************************************************/

    function initiateCancellation(
        uint256 tokenId,
        uint256 refundAmount,
        uint16 cancellationReason,
        uint16 cancellationReasonVersion
    ) external onlyRole(BOOKING_OPERATOR_ROLE) {
        BookingTokenOperator.initiateCancellation(
            getBookingTokenAddress(),
            tokenId,
            refundAmount,
            cancellationReason,
            cancellationReasonVersion
        );
    }

    function acceptCancellation(uint256 tokenId, uint256 refundAmount) external onlyRole(BOOKING_OPERATOR_ROLE) {
        BookingTokenOperator.acceptCancellation(getBookingTokenAddress(), tokenId, refundAmount);
    }

    function rejectCancellation(
        uint256 tokenId,
        uint16 rejectionReason,
        uint16 rejectionReasonVersion
    ) external onlyRole(BOOKING_OPERATOR_ROLE) {
        BookingTokenOperator.rejectCancellation(
            getBookingTokenAddress(),
            tokenId,
            rejectionReason,
            rejectionReasonVersion
        );
    }

    function counterCancellation(
        uint256 tokenId,
        uint256 refundAmount,
        uint16 counterReason,
        uint16 counterReasonVersion
    ) external onlyRole(BOOKING_OPERATOR_ROLE) {
        BookingTokenOperator.counterCancellation(
            getBookingTokenAddress(),
            tokenId,
            refundAmount,
            counterReason,
            counterReasonVersion
        );
    }

    /**
     * @notice Withdraws an active cancellation proposal. Only the initiator can withdraw.
     *
     * @param tokenId The token id for which to withdraw the proposal
     * @param reason The reason for withdrawing the proposal
     * @param reasonVersion The version of the withdrawal reason from the Travel Token Messenger Protocol
     */
    function withdrawCancellation(
        uint256 tokenId,
        uint16 reason,
        uint16 reasonVersion
    ) external onlyRole(BOOKING_OPERATOR_ROLE) {
        BookingTokenOperator.withdrawCancellation(getBookingTokenAddress(), tokenId, reason, reasonVersion);
    }

    /**
     * @notice Finalizes a cancellation proposal. Only the supplier of the token can finalize.
     *
     * @param tokenId The token id for which to finalize the proposal
     * @param refundAmount The refund amount to check, this is to prevent front-running attacks
     */
    function finalizeCancellation(uint256 tokenId, uint256 refundAmount) external onlyRole(BOOKING_OPERATOR_ROLE) {
        BookingTokenOperator.finalizeCancellation(getBookingTokenAddress(), tokenId, refundAmount);
    }
}
