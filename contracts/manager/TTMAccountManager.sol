// SPDX-License-Identifier: LGPL-3.0-or-later
//
// Travel Token Messenger Account Manager

pragma solidity 0.8.24;

// UUPS Proxy
import { Initializable } from "@openzeppelin/contracts-upgradeable/proxy/utils/Initializable.sol";
import { ERC1967Proxy } from "@openzeppelin/contracts/proxy/ERC1967/ERC1967Proxy.sol";
import { UUPSUpgradeable } from "@openzeppelin/contracts-upgradeable/proxy/utils/UUPSUpgradeable.sol";

// Access
import { PausableUpgradeable } from "@openzeppelin/contracts-upgradeable/utils/PausableUpgradeable.sol";
import { AccessControlEnumerableUpgradeable } from "@openzeppelin/contracts-upgradeable/access/extensions/AccessControlEnumerableUpgradeable.sol";
import { ReentrancyGuardUpgradeable } from "@openzeppelin/contracts-upgradeable/utils/ReentrancyGuardUpgradeable.sol";

// ABI of the TTMAccount implementation contract
import { ITTMAccount } from "../account/ITTMAccount.sol";

// Utils
import { Address } from "@openzeppelin/contracts/utils/Address.sol";
import { EnumerableSet } from "@openzeppelin/contracts/utils/structs/EnumerableSet.sol";

// Service Registry
import { ServiceRegistry } from "../partner/ServiceRegistry.sol";

/**
 * @title Travel Token Messenger Account Manager
 * @notice This contract manages the creation of the Travel Token Messenger accounts by
 * deploying {ERC1967Proxy} proxies that point to the{TTMAccount} implementation
 * address.
 *
 * Create TTM Account: Users who want to create an account should call
 * `createTTMAccount(address admin, address upgrader)` function with addresses of
 * the accounts admin and upgrader roles.
 *
 * When the manager contract is paused, account creation is stopped.
 *
 * Service Registry: {TTMAccountManager} also acts as a registry for the services
 * that {TTMAccount} contracts add as a supported or wanted service. Registry
 * works by hashing (keccak256) the service name (string) and creating a mapping
 * as keccak256(serviceName) => serviceName. And provides functions that
 * {TTMAccount} function uses to register services. The {TTMAccount} only keeps
 * the hashes (byte32) of the registered services.
 */
contract TTMAccountManager is
    Initializable,
    PausableUpgradeable,
    AccessControlEnumerableUpgradeable,
    UUPSUpgradeable,
    ReentrancyGuardUpgradeable,
    ServiceRegistry
{
    using Address for address payable;
    using EnumerableSet for EnumerableSet.AddressSet;

    /**
     * @notice Pauser role can pause the contract. Currently this only affects the
     * creation of TTM Accounts. When paused, account creation is stopped.
     */
    bytes32 public constant PAUSER_ROLE = keccak256("PAUSER_ROLE");

    /**
     * @notice Upgrader role can upgrade the contract to a new implementation.
     */
    bytes32 public constant UPGRADER_ROLE = keccak256("UPGRADER_ROLE");

    /**
     * @notice Versioner role can set new {TTMAccount} implementation address. When a
     * new implementation address is set, it is used for the new {TTMAccount}
     * creations.
     *
     * The old {TTMAccount} contracts are not affected by this. Owners of those
     * should do the upgrade manually by calling the `upgradeToAndCall(address)`
     * function on the account.
     */
    bytes32 public constant VERSIONER_ROLE = keccak256("VERSIONER_ROLE");

    /**
     * @notice Service registry admin role can add and remove services to the service
     * registry mapping. Implemented by {ServiceRegistry} contract.
     */
    bytes32 public constant SERVICE_REGISTRY_ADMIN_ROLE = keccak256("SERVICE_REGISTRY_ADMIN_ROLE");

    /***************************************************
     *                   STORAGE                       *
     ***************************************************/

    /// @custom:storage-location erc7201:traveltoken.messenger.storage.TTMAccountManager
    struct TTMAccountManagerStorage {
        /**
         * @dev TTM Account implementation address to be used by the TTMAccount contract to restrict
         * the implementation address for the UUPS proxies.
         */
        address _latestAccountImplementation;
        /**
         * @dev BookingToken address.
         */
        address _bookingToken;
        /**
         * @dev Every TTM Account created by this manager. This is the single source of
         * truth for account identity. `_createTTMAccount` is the only writer and there
         * is no external mutator, so this set cannot diverge from reality.
         */
        EnumerableSet.AddressSet _ttmAccounts;
        /**
         * @dev Creator of each TTM Account, keyed by the account address.
         */
        mapping(address account => address creator) _ttmAccountCreator;
    }

    // keccak256(abi.encode(uint256(keccak256("traveltoken.messenger.storage.TTMAccountManager")) - 1)) & ~bytes32(uint256(0xff));
    bytes32 private constant TTMAccountManagerStorageLocation =
        0x82fd17ead72ea2acbf4028d3b6fb6ced2f61d2d2be6f5996ded313320918c700;

    function _getTTMAccountManagerStorage() private pure returns (TTMAccountManagerStorage storage $) {
        assembly {
            $.slot := TTMAccountManagerStorageLocation
        }
    }

    /***************************************************
     *                    EVENTS                       *
     ***************************************************/

    /**
     * @notice Emitted when a TTM Account is created.
     *
     * @dev Carries creator and admin so indexers need no follow-up call per account.
     */
    event TTMAccountCreated(address indexed account, address indexed creator, address indexed admin);

    /**
     * @notice TTM Account implementation address updated event.
     * @param oldImplementation The old implementation address
     * @param newImplementation The new implementation address
     */
    event TTMAccountImplementationUpdated(address indexed oldImplementation, address indexed newImplementation);

    /**
     * @notice Booking token address updated event.
     * @param oldBookingToken The old booking token address
     * @param newBookingToken The new booking token address
     */
    event BookingTokenAddressUpdated(address indexed oldBookingToken, address indexed newBookingToken);

    /***************************************************
     *                    ERRORS                       *
     ***************************************************/

    /**
     * @notice The implementation of the TTMAccount is invalid.
     * @param implementation The implementation address of the TTMAccount
     */
    error TTMAccountInvalidImplementation(address implementation);

    /**
     * @notice The admin address is invalid.
     * @param admin The admin address
     */
    error TTMAccountInvalidAdmin(address admin);

    /**
     * @notice Invalid booking token address.
     * @param bookingToken The booking token address
     */
    error InvalidBookingTokenAddress(address bookingToken);

    /**
     * @notice A required address parameter was the zero address.
     */
    error ZeroAddress();

    /***************************************************
     *                    FUNCS                        *
     ***************************************************/

    /// @custom:oz-upgrades-unsafe-allow constructor
    constructor() {
        _disableInitializers();
    }

    function initialize(
        address defaultAdmin, // can grant roles
        address pauser, // can pause the manager
        address upgrader, // can upgrade the manager (this contract)
        address versioner // can set TTMAccount implementation address
    ) public initializer {
        if (defaultAdmin == address(0) || pauser == address(0) || upgrader == address(0) || versioner == address(0)) {
            revert ZeroAddress();
        }

        __Pausable_init();
        __AccessControl_init();
        __UUPSUpgradeable_init();
        __ReentrancyGuard_init();
        __ServiceRegistry_init();

        _grantRole(DEFAULT_ADMIN_ROLE, defaultAdmin);
        _grantRole(PAUSER_ROLE, pauser);
        _grantRole(UPGRADER_ROLE, upgrader);
        _grantRole(VERSIONER_ROLE, versioner);
    }

    /**
     * @notice Pauses the TTMAccountManager contract. Currently this only affects the
     * creation of TTMAccount. When paused, account creation is stopped.
     */
    function pause() public onlyRole(PAUSER_ROLE) {
        _pause();
    }

    /**
     * @notice Unpauses the TTMAccountManager contract.
     */
    function unpause() public onlyRole(PAUSER_ROLE) {
        _unpause();
    }

    /**
     * @notice Authorization for the TTMAccountManager contract upgrade.
     */
    function _authorizeUpgrade(address newImplementation) internal override onlyRole(UPGRADER_ROLE) {}

    /***************************************************
     *                    ACCOUNT                      *
     ***************************************************/

    /**
     * @notice Creates a new TTMAccount.
     *
     * This function is currently permissionless: any address may create an
     * account. See docs/decisions/2026-07-21-contract-design-decisions.md
     * (Decision 1) -- gating must be resolved before Base mainnet.
     *
     * @dev Emits a {TTMAccountCreated} event.
     */
    function createTTMAccount(
        address admin,
        address upgrader
    ) external payable nonReentrant whenNotPaused returns (address) {
        return _createTTMAccount(admin, upgrader);
    }

    /**
     * @notice Private function to create a `TTMAccount`.
     */
    function _createTTMAccount(address admin, address upgrader) private returns (address) {
        // Checks
        if (admin == address(0)) {
            revert TTMAccountInvalidAdmin(admin);
        }

        address latestAccountImplementation = getAccountImplementation();
        if (latestAccountImplementation.code.length == 0) {
            revert TTMAccountInvalidImplementation(latestAccountImplementation);
        }

        address bookingToken = getBookingTokenAddress();
        if (bookingToken.code.length == 0) {
            revert InvalidBookingTokenAddress(bookingToken);
        }

        // Create TTMAccount Proxy and set the implementation address
        ERC1967Proxy ttmAccountProxy = new ERC1967Proxy(latestAccountImplementation, "");

        // Initialize the TTMAccount
        ITTMAccount(address(ttmAccountProxy)).initialize(address(this), bookingToken, admin, upgrader);

        // Record the account and its creator. This is the only writer.
        TTMAccountManagerStorage storage $ = _getTTMAccountManagerStorage();
        $._ttmAccounts.add(address(ttmAccountProxy));
        $._ttmAccountCreator[address(ttmAccountProxy)] = msg.sender;

        // [ETH] Send the msg.value to the TTMAccount
        payable(ttmAccountProxy).sendValue(msg.value);

        emit TTMAccountCreated(address(ttmAccountProxy), msg.sender, admin);

        return address(ttmAccountProxy);
    }

    /**
     * @notice Returns the given account's creator, or the zero address if the
     * address is not a TTM Account.
     *
     * @param account The account address
     */
    function getTTMAccountCreator(address account) public view returns (address) {
        return _getTTMAccountManagerStorage()._ttmAccountCreator[account];
    }

    /**
     * @notice Returns whether the given address is a TTM Account created by this manager.
     *
     * @param account The address to check
     */
    function isTTMAccount(address account) public view returns (bool) {
        return _getTTMAccountManagerStorage()._ttmAccounts.contains(account);
    }

    /**
     * @notice Returns the number of TTM Accounts created by this manager.
     */
    function getTTMAccountCount() public view returns (uint256) {
        return _getTTMAccountManagerStorage()._ttmAccounts.length();
    }

    /**
     * @notice Returns every TTM Account created by this manager.
     *
     * @dev Unbounded. Prefer {getTTMAccountsSlice} against a public RPC once the
     * ecosystem grows past a few hundred accounts.
     */
    function getTTMAccounts() public view returns (address[] memory) {
        return _getTTMAccountManagerStorage()._ttmAccounts.values();
    }

    /**
     * @notice Returns a bounded window of TTM Accounts, for callers that cannot
     * afford an unbounded read.
     *
     * Returns an empty array if `offset` is at or past the end. The window is
     * clamped to the end of the set, so a `limit` larger than the remainder is
     * not an error.
     *
     * @param offset Index to start at
     * @param limit Maximum number of accounts to return
     */
    function getTTMAccountsSlice(uint256 offset, uint256 limit) public view returns (address[] memory accounts) {
        EnumerableSet.AddressSet storage ttmAccounts = _getTTMAccountManagerStorage()._ttmAccounts;
        uint256 total = ttmAccounts.length();
        if (offset >= total) {
            return new address[](0);
        }

        // Clamp by subtraction, not by computing `offset + limit`: under checked
        // arithmetic that sum reverts for a large `limit`, which would contradict
        // the "an oversized limit is not an error" contract above. `offset < total`
        // here, so `total - offset` cannot underflow.
        uint256 remaining = total - offset;
        if (limit > remaining) {
            limit = remaining;
        }

        accounts = new address[](limit);
        for (uint256 i = 0; i < limit; i++) {
            accounts[i] = ttmAccounts.at(offset + i);
        }
    }

    /***************************************************
     *             ACCOUNT IMPLEMENTATION              *
     ***************************************************/

    /**
     * @notice Returns the TTMAccount implementation address.
     */
    function getAccountImplementation() public view returns (address) {
        TTMAccountManagerStorage storage $ = _getTTMAccountManagerStorage();
        return $._latestAccountImplementation;
    }

    /**
     * @notice Set a new TTMAccount implementation address.
     * @param newImplementation The new implementation address
     */
    function setAccountImplementation(address newImplementation) public onlyRole(VERSIONER_ROLE) {
        if (newImplementation.code.length == 0) {
            revert TTMAccountInvalidImplementation(newImplementation);
        }

        address oldImplementation = getAccountImplementation();
        _setAccountImplementation(newImplementation);
        emit TTMAccountImplementationUpdated(oldImplementation, newImplementation);
    }

    function _setAccountImplementation(address newImplementation) internal {
        TTMAccountManagerStorage storage $ = _getTTMAccountManagerStorage();
        $._latestAccountImplementation = newImplementation;
    }

    /***************************************************
     *                  BOOKING TOKEN                  *
     ***************************************************/

    /**
     * @notice Returns the booking token address.
     */
    function getBookingTokenAddress() public view returns (address) {
        TTMAccountManagerStorage storage $ = _getTTMAccountManagerStorage();
        return $._bookingToken;
    }

    /**
     * @notice Sets booking token address.
     */
    function setBookingTokenAddress(address token) public onlyRole(VERSIONER_ROLE) {
        if (token.code.length == 0) {
            revert InvalidBookingTokenAddress(token);
        }

        address oldToken = getBookingTokenAddress();
        _setBookingTokenAddress(token);
        emit BookingTokenAddressUpdated(oldToken, token);
    }

    function _setBookingTokenAddress(address token) internal {
        TTMAccountManagerStorage storage $ = _getTTMAccountManagerStorage();
        $._bookingToken = token;
    }

    /***************************************************
     *               SERVICE REGISTRY                  *
     ***************************************************/

    /**
     * @notice Registers a given service name. TTM Accounts can only register services
     * if they are also registered in the service registry on the manager contract.
     *
     * @param serviceName Name of the service
     */
    function registerService(string memory serviceName) public onlyRole(SERVICE_REGISTRY_ADMIN_ROLE) {
        _registerServiceName(serviceName);
    }

    /**
     * @notice Unregisters a given service name. TTM Accounts will not be able to register
     * the service anymore.
     *
     * @param serviceName Name of the service
     */
    function unregisterService(string memory serviceName) public onlyRole(SERVICE_REGISTRY_ADMIN_ROLE) {
        _unregisterServiceName(serviceName);
    }
}
