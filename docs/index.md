# Solidity API

## GasMoneyManager

GasMoneyManager manages gas money withdrawals for a {TTMAccount}.

Gas money withdrawals are restricted to a withdrawal limit and period.

### GasMoneyWithdrawalRecord

Per-account withdrawal accounting, packed into a single slot.

```solidity
struct GasMoneyWithdrawalRecord {
    uint128 amount;
    uint64 periodStart;
}
```

### GasMoneyStorage

```solidity
struct GasMoneyStorage {
  mapping(address => struct GasMoneyManager.GasMoneyWithdrawalRecord) _withdrawals;
  uint128 _withdrawalLimit;
  uint64 _withdrawalPeriod;
}
```

### GasMoneyWithdrawal

```solidity
event GasMoneyWithdrawal(address withdrawer, uint256 amount)
```

Gas money withdrawal event

#### Parameters

| Name       | Type    | Description                   |
| ---------- | ------- | ----------------------------- |
| withdrawer | address | the address of the withdrawer |
| amount     | uint256 | the amount withdrawn          |

### GasMoneyWithdrawalUpdated

```solidity
event GasMoneyWithdrawalUpdated(uint256 limit, uint256 period)
```

Gas money withdrawal limit and period updated event

#### Parameters

| Name   | Type    | Description                         |
| ------ | ------- | ----------------------------------- |
| limit  | uint256 | the withdrawal limit for the period |
| period | uint256 | the withdrawal period in seconds    |

### WithdrawalLimitExceeded

```solidity
error WithdrawalLimitExceeded(uint256 limit, uint256 amount)
```

### WithdrawalLimitExceededForPeriod

```solidity
error WithdrawalLimitExceededForPeriod(uint256 limit, uint256 amount)
```

### GasMoneyValueOutOfRange

```solidity
error GasMoneyValueOutOfRange(uint256 limit, uint256 period)
```

### \_\_GasMoneyManager_init

```solidity
function __GasMoneyManager_init(uint256 withdrawalLimit, uint256 withdrawalPeriod) internal
```

### \_withdrawGasMoney

```solidity
function _withdrawGasMoney(uint256 amount) internal
```

Withdraws gas money.

This functions is intended to be called by the bot to withdraw gas money.
Inheriting contract should restrict who can call this with a public
function.

### \_setGasMoneyWithdrawal

```solidity
function _setGasMoneyWithdrawal(uint256 limit, uint256 period) internal
```

Sets the gas money withdrawal limit and period.

#### Parameters

| Name   | Type    | Description                         |
| ------ | ------- | ----------------------------------- |
| limit  | uint256 | the withdrawal limit for the period |
| period | uint256 | the withdrawal period in seconds    |

### getGasMoneyWithdrawal

```solidity
function getGasMoneyWithdrawal() public view returns (uint256 withdrawalLimit, uint256 withdrawalPeriod)
```

Returns the gas money withdrawal restrictions.

#### Return Values

| Name             | Type    | Description |
| ---------------- | ------- | ----------- |
| withdrawalLimit  | uint256 |             |
| withdrawalPeriod | uint256 |             |

### getGasMoneyWithdrawalForAccount

```solidity
function getGasMoneyWithdrawalForAccount(address account) public view returns (uint256 periodStart, uint256 withdrawnAmount)
```

Returns the gas money withdrawal details for an account.

#### Parameters

| Name    | Type    | Description            |
| ------- | ------- | ---------------------- |
| account | address | address of the account |

#### Return Values

| Name            | Type    | Description                              |
| --------------- | ------- | ---------------------------------------- |
| periodStart     | uint256 | timestamp of the withdrawal period start |
| withdrawnAmount | uint256 | amount withdrawn within the period       |

## ITTMAccount

### initialize

```solidity
function initialize(address manager, address bookingToken, address owner, address upgrader) external
```

### isSupportedToken

```solidity
function isSupportedToken(address _token) external view returns (bool)
```

Whether a payment token is declared as supported by this account.

Payment mode is encoded as an address, matching BookingToken:
`address(0)` is native currency, `address(1)` is off-chain payment, and
any other value is an ERC-20 address. All three are declared through the
same allowlist.

## TTMAccount

A TTM Account manages funds, minting/buying of booking tokens, provided
or wanted services, and multiple bots for distributors and suppliers on
Travel Token Messenger ecosystem.

Registering bots is done by role based access control. Bot's with
`MESSENGER_BOT_ROLE` are authorized to represent the TTMAccount.
Bot can also have `GAS_WITHDRAWER_ROLE` and `BOOKING_OPERATOR_ROLE`.

`GAS_WITHDRAWER_ROLE` enables a bot to withdraw native coins (ETH) from the
contract to be used as gas money. This is restricted with a `limit` (wei)
and `period` (seconds) set by the `BOT_ADMIN_ROLE`. The limit and period
apply per bot address: each bot tracks its own withdrawals against the
same limit, independently of every other bot on the account. Default
starting values are 0.01 ETH per 24 hours.

Unlike `MESSENGER_BOT_ROLE` and `BOOKING_OPERATOR_ROLE`, `GAS_WITHDRAWER_ROLE`
is not granted by `addMessengerBot`. It must be granted explicitly by
`DEFAULT_ADMIN_ROLE` via the inherited `grantRole`. This means a compromised
bot key does not, by itself, come with standing authority to withdraw gas
money: that authority is opt-in per bot, decided separately by
`DEFAULT_ADMIN_ROLE`. It is not a fund-safety boundary against
`BOT_ADMIN_ROLE` itself — `addMessengerBot` takes a `gasMoney` argument and
sends that amount to the new bot immediately, and `BOT_ADMIN_ROLE` can also
raise the default withdrawal limit via `setGasMoneyWithdrawal`.

`BOOKING_OPERATOR_ROLE` enables a bot to mint and buy Booking Tokens by
calling the corresponding functions on the {BookingToken} contract. The buy
operation pays the price of the Booking Token with the funds on the
{TTMAccount} contract.

_This contract uses UUPS style upgradeability. The authorization function
`_authorizeUpgrade(address)` can be called by the `UPGRADER_ROLE` and is
restricted to only upgrade to the implementation address registered on the
{TTMAccountManager} contract._

### UPGRADER_ROLE

```solidity
bytes32 UPGRADER_ROLE
```

Upgrader role can upgrade the contract to a new implementation.

### BOT_ADMIN_ROLE

```solidity
bytes32 BOT_ADMIN_ROLE
```

Bot admin role can add & remove bots and set gas money withdrawal
parameters.

### MESSENGER_BOT_ROLE

```solidity
bytes32 MESSENGER_BOT_ROLE
```

Messenger bot role can interact on behalf of this TTMAccount
contract.

### GAS_WITHDRAWER_ROLE

```solidity
bytes32 GAS_WITHDRAWER_ROLE
```

Gas withdrawer role can withdraw gas money from the contract. This is
intended to be used by the bots, but is not granted by `addMessengerBot`.
`DEFAULT_ADMIN_ROLE` must grant it explicitly, so a `BOT_ADMIN_ROLE`
holder can onboard and remove bots but cannot give them access to funds.

### WITHDRAWER_ROLE

```solidity
bytes32 WITHDRAWER_ROLE
```

Withdrawer role can withdraw funds from the contract.

### BOOKING_OPERATOR_ROLE

```solidity
bytes32 BOOKING_OPERATOR_ROLE
```

Booking operator role can mint and buy booking tokens using the
functions on this contract. This is generally used by the bots. The
price for the booking token is paid by this contract.

### SERVICE_ADMIN_ROLE

```solidity
bytes32 SERVICE_ADMIN_ROLE
```

Service admin role can add & remove supported & wanted services.

### TTMAccountStorage

```solidity
struct TTMAccountStorage {
    address _manager;
    address _bookingToken;
}
```

### TTMAccountUpgraded

```solidity
event TTMAccountUpgraded(address oldImplementation, address newImplementation)
```

TTMAccount upgrade event. Emitted when the TTMAccount implementation is upgraded.

### Deposit

```solidity
event Deposit(address sender, uint256 amount)
```

Deposit event, emitted when there is a new deposit

### Withdraw

```solidity
event Withdraw(address receiver, uint256 amount)
```

Withdraw event, emitted when there is a new withdrawal

### MessengerBotAdded

```solidity
event MessengerBotAdded(address bot)
```

Messenger bot added

### MessengerBotRemoved

```solidity
event MessengerBotRemoved(address bot)
```

Messenger bot removed

### ServiceAdded

```solidity
event ServiceAdded(bytes32 serviceHash)
```

\_Service events carry the service hash only. Indexing a dynamic `string`
stores just its keccak hash in the topic and nothing in the data section, so the
old `string indexed serviceName` form published a hash while pretending to
publish a name. Consumers resolve names from `ServiceRegistry`'s
`ServiceRegistered` / `ServiceUnregistered` events, which do carry them.

Capability strings stay readable: capabilities are free-form partner text with
no registry to resolve against.\_

### ServiceRemoved

```solidity
event ServiceRemoved(bytes32 serviceHash)
```

### WantedServiceAdded

```solidity
event WantedServiceAdded(bytes32 serviceHash)
```

### WantedServiceRemoved

```solidity
event WantedServiceRemoved(bytes32 serviceHash)
```

### ServiceRestrictedRateUpdated

```solidity
event ServiceRestrictedRateUpdated(bytes32 serviceHash, bool restrictedRate)
```

### ServiceCapabilitiesUpdated

```solidity
event ServiceCapabilitiesUpdated(bytes32 serviceHash)
```

### ServiceCapabilityAdded

```solidity
event ServiceCapabilityAdded(bytes32 serviceHash, string capability)
```

### ServiceCapabilityRemoved

```solidity
event ServiceCapabilityRemoved(bytes32 serviceHash, string capability)
```

### TTMAccountImplementationMismatch

```solidity
error TTMAccountImplementationMismatch(address latestImplementation, address newImplementation)
```

TTMAccount implementation address does not match the one in the manager

### TTMAccountNoUpgradeNeeded

```solidity
error TTMAccountNoUpgradeNeeded(address oldImplementation, address newImplementation)
```

New implementation is the same as the current implementation, no update needed

### TransferToZeroAddress

```solidity
error TransferToZeroAddress()
```

Error to revert if transfer to zero address

### ZeroAddress

```solidity
error ZeroAddress()
```

A required address parameter was the zero address.

### ServiceNotRegistered

```solidity
error ServiceNotRegistered()
```

The given service hash is not registered in the manager's ServiceRegistry.

_Same selector as ServiceRegistry's `ServiceNotRegistered()` (identical, argument-less
signature) since this error is what actually bubbles up from the staticcall in
{\_requireRegisteredService}; declaring it here as well only lets this contract's ABI
name it directly._

### constructor

```solidity
constructor() public
```

### initialize

```solidity
function initialize(address manager, address bookingToken, address defaultAdmin, address upgrader) public
```

### receive

```solidity
receive() external payable
```

### getManagerAddress

```solidity
function getManagerAddress() public view returns (address)
```

Returns the TTMAccountManager address.

#### Return Values

| Name | Type    | Description               |
| ---- | ------- | ------------------------- |
| [0]  | address | TTMAccountManager address |

### getBookingTokenAddress

```solidity
function getBookingTokenAddress() public view returns (address)
```

Returns the booking token address.

#### Return Values

| Name | Type    | Description          |
| ---- | ------- | -------------------- |
| [0]  | address | BookingToken address |

### \_authorizeUpgrade

```solidity
function _authorizeUpgrade(address newImplementation) internal
```

Authorizes the upgrade of the TTMAccount.

Reverts if the new implementation is the same as the old one.

Reverts if the new implementation does not match the implementation address
in the manager. Only implementations registered at the manager are allowed.

_Emits a {TTMAccountUpgraded} event._

#### Parameters

| Name              | Type    | Description                    |
| ----------------- | ------- | ------------------------------ |
| newImplementation | address | The new implementation address |

### isBotAllowed

```solidity
function isBotAllowed(address bot) public view returns (bool)
```

Returns true if an address is an authorized messenger bot

#### Parameters

| Name | Type    | Description       |
| ---- | ------- | ----------------- |
| bot  | address | The bot's address |

### withdraw

```solidity
function withdraw(address payable recipient, uint256 amount) external
```

Withdraw ETH from the TTMAccount

#### Parameters

| Name      | Type            | Description                     |
| --------- | --------------- | ------------------------------- |
| recipient | address payable | The recipient of the withdrawal |
| amount    | uint256         | The amount to withdraw          |

### mintBookingToken

```solidity
function mintBookingToken(address reservedFor, string uri, uint256 expirationTimestamp, uint256 price, contract IERC20 paymentToken, uint256 offchainPaymentCurrency, bool cancellable) external
```

Mints booking token.

#### Parameters

| Name                    | Type            | Description                                  |
| ----------------------- | --------------- | -------------------------------------------- |
| reservedFor             | address         | The account to reserve the token for         |
| uri                     | string          | The URI of the token                         |
| expirationTimestamp     | uint256         | The expiration timestamp                     |
| price                   | uint256         | The price of the token                       |
| paymentToken            | contract IERC20 | The payment token, if address(0) then native |
| offchainPaymentCurrency | uint256         | The offchain payment currency                |
| cancellable             | bool            | If the token is cancellable                  |

### buyBookingToken

```solidity
function buyBookingToken(uint256 tokenId, uint256 expectedPrice, contract IERC20 expectedPaymentToken) external
```

Buys booking token.

#### Parameters

| Name                 | Type            | Description  |
| -------------------- | --------------- | ------------ |
| tokenId              | uint256         | The token id |
| expectedPrice        | uint256         |              |
| expectedPaymentToken | contract IERC20 |              |

### recordExpiration

```solidity
function recordExpiration(uint256 tokenId) external
```

Marks an expired reservation as expired on the BookingToken.

_Deliberately permissionless. The underlying `BookingToken.recordExpiration`
is public and unrestricted, so a role gate here would protect nothing - it only
created the false impression that one was needed. The operation is objective
housekeeping: it succeeds only once `block.timestamp` has genuinely passed the
reservation's expiry, so there is nothing for an attacker to gain._

#### Parameters

| Name    | Type    | Description                       |
| ------- | ------- | --------------------------------- |
| tokenId | uint256 | The booking token to mark expired |

### onERC721Received

```solidity
function onERC721Received(address, address, uint256, bytes) public virtual returns (bytes4)
```

Always returns `IERC721Receiver.onERC721Received.selector`.

_See {IERC721Receiver-onERC721Received}._

### supportsInterface

```solidity
function supportsInterface(bytes4 interfaceId) public view virtual returns (bool)
```

See {IERC165-supportsInterface}.

_This contract implements {IERC721Receiver}, so it must say so - counterparties
that capability-detect before transferring an NFT would otherwise conclude it
cannot receive one._

### transferERC20

```solidity
function transferERC20(contract IERC20 token, address to, uint256 amount) external
```

Transfers ERC20 tokens.

This function reverts if `to` is the zero address.

#### Parameters

| Name   | Type            | Description                           |
| ------ | --------------- | ------------------------------------- |
| token  | contract IERC20 | The ERC20 token                       |
| to     | address         | The address to transfer the tokens to |
| amount | uint256         | The amount of tokens to transfer      |

### transferERC721

```solidity
function transferERC721(contract IERC721 token, address to, uint256 tokenId) external
```

Transfers ERC721 tokens.

This function reverts if `to` is the zero address.

#### Parameters

| Name    | Type             | Description                           |
| ------- | ---------------- | ------------------------------------- |
| token   | contract IERC721 | The ERC721 token                      |
| to      | address          | The address to transfer the tokens to |
| tokenId | uint256          | The token id of the token             |

### approveERC721

```solidity
function approveERC721(contract IERC721 token, address to, uint256 tokenId) external
```

Approves an operator to transfer a specific ERC721 token held by
this account. Required for listing a booking token on a marketplace or
handing it to a custody provider.

_`token` is not restricted to ERC721 contracts: `approve(address,uint256)`
shares its selector with ERC-20's `approve`, so calling this with an
IERC20 address cast as IERC721 grants an ERC-20 allowance instead. This is
not a privilege escalation — `WITHDRAWER_ROLE` can already move the
account's ERC-20 balances outright via `transferERC20`._

#### Parameters

| Name    | Type             | Description                                                                                     |
| ------- | ---------------- | ----------------------------------------------------------------------------------------------- |
| token   | contract IERC721 | The ERC721 contract (or, due to the shared selector noted above, any ERC20-compatible contract) |
| to      | address          | The operator being approved                                                                     |
| tokenId | uint256          | The token id                                                                                    |

### addService

```solidity
function addService(bytes32 serviceHash, bool restrictedRate, string[] capabilities) public
```

Adds a service to the account as a supported service.

`serviceHash` is `keccak256(abi.encodePacked(serviceName))`, where the name is
pkg + service name as defined in the Travel Token Messenger Protocol's protobuf
definitions. For example:

```text
 ┌────────────── pkg ─────────────┐ ┌───── service name ─────┐
"ttm.services.accommodation.v1alpha.AccommodationSearchService")
```

_The hash must be registered in the manager's `ServiceRegistry`. That check is
the one manager staticcall left on this path: it is a write, called rarely, and
without it an account could advertise a service that does not exist. Reads carry
no manager dependency at all._

#### Parameters

| Name           | Type     | Description                                        |
| -------------- | -------- | -------------------------------------------------- |
| serviceHash    | bytes32  | Hash of the service name to support                |
| restrictedRate | bool     | Whether the service is restricted to pre-agreement |
| capabilities   | string[] | Capabilities of the service (optional)             |

### removeService

```solidity
function removeService(bytes32 serviceHash) public
```

Removes a service from the account by its hash.

### removeAllServices

```solidity
function removeAllServices() public
```

Removes all supported services from the account.

### setServiceRestrictedRate

```solidity
function setServiceRestrictedRate(bytes32 serviceHash, bool restrictedRate) public
```

Sets whether a service is offered at a restricted (non-rack) rate.

### setServiceCapabilities

```solidity
function setServiceCapabilities(bytes32 serviceHash, string[] capabilities) public
```

Replaces the capability list of a service.

### addServiceCapability

```solidity
function addServiceCapability(bytes32 serviceHash, string capability) public
```

Adds a single capability to a service.

### removeServiceCapability

```solidity
function removeServiceCapability(bytes32 serviceHash, string capability) public
```

Removes a single capability from a service.

### getSupportedServices

```solidity
function getSupportedServices() public view returns (bytes32[] serviceHashes, struct PartnerConfiguration.Service[] services)
```

Returns every supported service as a hash plus its stored record.

_Reads no longer touch the manager. Resolve hashes to names client-side from
the registry's `ServiceRegistered` events or `getAllRegisteredServiceNames()`.
Unbounded - prefer {getSupportedServicesSlice} against a public RPC._

### getSupportedServicesSlice

```solidity
function getSupportedServicesSlice(uint256 offset, uint256 limit) public view returns (bytes32[] serviceHashes, struct PartnerConfiguration.Service[] services)
```

Returns a bounded window of supported services.

Returns empty arrays if `offset` is at or past the end; the window is clamped to
the end of the list, so an oversized `limit` is not an error.

#### Parameters

| Name   | Type    | Description                          |
| ------ | ------- | ------------------------------------ |
| offset | uint256 | Index to start at                    |
| limit  | uint256 | Maximum number of services to return |

### isServiceSupported

```solidity
function isServiceSupported(bytes32 serviceHash) public view returns (bool)
```

Checks whether a service is supported by this account.

#### Parameters

| Name        | Type    | Description                       |
| ----------- | ------- | --------------------------------- |
| serviceHash | bytes32 | Hash of the service name to check |

### addWantedServices

```solidity
function addWantedServices(bytes32[] serviceHashes) public
```

Declares services this account wants to consume from other partners.

_Each hash must be registered in the manager's ServiceRegistry, for the same
reason as {addService}._

#### Parameters

| Name          | Type      | Description                         |
| ------------- | --------- | ----------------------------------- |
| serviceHashes | bytes32[] | Hashes of the service names to want |

### removeWantedServices

```solidity
function removeWantedServices(bytes32[] serviceHashes) public
```

Removes services from this account's wanted list.

#### Parameters

| Name          | Type      | Description                                 |
| ------------- | --------- | ------------------------------------------- |
| serviceHashes | bytes32[] | Hashes of the service names to stop wanting |

### addSupportedToken

```solidity
function addSupportedToken(address _supportedToken) public
```

Adds a supported payment token.

#### Parameters

| Name             | Type    | Description          |
| ---------------- | ------- | -------------------- |
| \_supportedToken | address | address of the token |

### removeSupportedToken

```solidity
function removeSupportedToken(address _supportedToken) public
```

Removes a supported payment token.

#### Parameters

| Name             | Type    | Description          |
| ---------------- | ------- | -------------------- |
| \_supportedToken | address | address of the token |

### addPublicKey

```solidity
function addPublicKey(address pubKeyAddress, bytes data) public
```

Add public key with address

These public keys are intended to be used with for off-chain encryption of private booking data.

#### Parameters

| Name          | Type    | Description               |
| ------------- | ------- | ------------------------- |
| pubKeyAddress | address | address of the public key |
| data          | bytes   | public key data           |

### removePublicKey

```solidity
function removePublicKey(address pubKeyAddress) public
```

Remove public key by address

### addMessengerBot

```solidity
function addMessengerBot(address bot, uint256 gasMoney) public
```

Adds messenger bot with initial gas money. The amount of `gasMoney`
need to be present in the contract.

### removeMessengerBot

```solidity
function removeMessengerBot(address bot) public
```

Removes messenger bot by revoking the roles.

### withdrawGasMoney

```solidity
function withdrawGasMoney(uint256 amount) public
```

Withdraw gas money. Requires the `GAS_WITHDRAWER_ROLE`.

#### Parameters

| Name   | Type    | Description                   |
| ------ | ------- | ----------------------------- |
| amount | uint256 | The amount to withdraw in wei |

### setGasMoneyWithdrawal

```solidity
function setGasMoneyWithdrawal(uint256 limit, uint256 period) public
```

Set gas money withdrawal parameters. Requires the `BOT_ADMIN_ROLE`.

#### Parameters

| Name   | Type    | Description                                       |
| ------ | ------- | ------------------------------------------------- |
| limit  | uint256 | Amount of gas money to withdraw in wei per period |
| period | uint256 | Duration of the withdrawal period in seconds      |

### initiateCancellation

```solidity
function initiateCancellation(uint256 tokenId, uint256 refundAmount, uint16 cancellationReason, uint16 cancellationReasonVersion) external
```

### acceptCancellation

```solidity
function acceptCancellation(uint256 tokenId, uint256 refundAmount) external
```

### rejectCancellation

```solidity
function rejectCancellation(uint256 tokenId, uint16 rejectionReason, uint16 rejectionReasonVersion) external
```

### counterCancellation

```solidity
function counterCancellation(uint256 tokenId, uint256 refundAmount, uint16 counterReason, uint16 counterReasonVersion) external
```

### withdrawCancellation

```solidity
function withdrawCancellation(uint256 tokenId, uint16 reason, uint16 reasonVersion) external
```

Withdraws an active cancellation proposal. Only the initiator can withdraw.

#### Parameters

| Name          | Type    | Description                                                                   |
| ------------- | ------- | ----------------------------------------------------------------------------- |
| tokenId       | uint256 | The token id for which to withdraw the proposal                               |
| reason        | uint16  | The reason for withdrawing the proposal                                       |
| reasonVersion | uint16  | The version of the withdrawal reason from the Travel Token Messenger Protocol |

### finalizeCancellation

```solidity
function finalizeCancellation(uint256 tokenId, uint256 refundAmount) external
```

Finalizes a cancellation proposal. Only the supplier of the token can finalize.

#### Parameters

| Name         | Type    | Description                                                          |
| ------------ | ------- | -------------------------------------------------------------------- |
| tokenId      | uint256 | The token id for which to finalize the proposal                      |
| refundAmount | uint256 | The refund amount to check, this is to prevent front-running attacks |

## BookingToken

Booking Token contract represents a booking done on the Travel Token Messenger.

Suppliers can mint Booking Tokens and reserve them for a distributor address to
buy.

Booking Tokens can have zero price, meaning that the payment will be done
off-chain.

When a token is minted with a reservation, it can not be transferred until the
expiration timestamp is reached or the token is bought.

### VERSION_MAJOR

```solidity
uint16 VERSION_MAJOR
```

### VERSION_MINOR

```solidity
uint16 VERSION_MINOR
```

### VERSION_PATCH

```solidity
uint16 VERSION_PATCH
```

### version

```solidity
function version() external pure virtual returns (uint16 major, uint16 minor, uint16 patch)
```

Returns the semantic version of the contract.

- no version() func: Legacy version without Cancellation support
- v1.0.0: Version with Cancellation support

#### Return Values

| Name  | Type   | Description                                   |
| ----- | ------ | --------------------------------------------- |
| major | uint16 | Major version (breaking changes)              |
| minor | uint16 | Minor version (backwards-compatible features) |
| patch | uint16 | Patch version (backwards-compatible fixes)    |

### UPGRADER_ROLE

```solidity
bytes32 UPGRADER_ROLE
```

Upgrader role can upgrade the contract to a new implementation.

### MIN_EXPIRATION_ADMIN_ROLE

```solidity
bytes32 MIN_EXPIRATION_ADMIN_ROLE
```

This role can set the mininum allowed expiration timestamp difference.

### PAUSER_ROLE

```solidity
bytes32 PAUSER_ROLE
```

Pauser role can pause the contract, halting minting, buying, and
cancellation finalization.

### NATIVE_PAYMENT

```solidity
address NATIVE_PAYMENT
```

Tokens are directly transferred to the recipient.

_Special address for native payments._

### OFFCHAIN_PAYMENT

```solidity
address OFFCHAIN_PAYMENT
```

A third-party service is used to handle payments.

_Special address for offchain payments. The enum for this
is defined in the Travel Token Messenger Protocol's
ttm.types.<version>.IsoCurrency enum (currency.proto file)._

### BookingStatus

```solidity
enum BookingStatus {
    UNSPECIFIED,
    RESERVED,
    RESERVATION_EXPIRED,
    BOUGHT,
    CANCELLED
}
```

### TokenReservation

```solidity
struct TokenReservation {
  address reservedFor;
  address supplier;
  uint256 expirationTimestamp;
  uint256 price;
  contract IERC20 paymentToken;
  uint256 offchainPaymentCurrency;
  bool cancellable;
}
```

### BookingTokenStorage

```solidity
struct BookingTokenStorage {
  address _manager;
  uint256 _nextTokenId;
  uint256 _minExpirationTimestampDiff;
  mapping(uint256 => struct BookingToken.TokenReservation) _reservations;
  mapping(uint256 => enum BookingToken.BookingStatus) _bookingStatus;
}
```

### \_getBookingTokenStorage

```solidity
function _getBookingTokenStorage() internal pure returns (struct BookingToken.BookingTokenStorage $)
```

### TokenReserved

```solidity
event TokenReserved(uint256 tokenId, address reservedFor, address supplier, uint256 expirationTimestamp, uint256 price, contract IERC20 paymentToken, uint256 offchainPaymentCurrency, bool cancellable)
```

Event emitted when a token is reserved.

#### Parameters

| Name                    | Type            | Description           |
| ----------------------- | --------------- | --------------------- |
| tokenId                 | uint256         | token id              |
| reservedFor             | address         | reserved for address  |
| supplier                | address         | supplier address      |
| expirationTimestamp     | uint256         | expiration timestamp  |
| price                   | uint256         | price of the token    |
| paymentToken            | contract IERC20 | payment token address |
| offchainPaymentCurrency | uint256         |                       |
| cancellable             | bool            |                       |

### TokenBought

```solidity
event TokenBought(uint256 tokenId, address buyer)
```

Event emitted when a token is bought.

#### Parameters

| Name    | Type    | Description   |
| ------- | ------- | ------------- |
| tokenId | uint256 | token id      |
| buyer   | address | buyer address |

### TokenReservationExpired

```solidity
event TokenReservationExpired(uint256 tokenId)
```

Event emitted when a token is expired.

#### Parameters

| Name    | Type    | Description |
| ------- | ------- | ----------- |
| tokenId | uint256 | token id    |

### ManagerAddressUpdated

```solidity
event ManagerAddressUpdated(address oldManager, address newManager)
```

Emitted when the manager address is changed.

_This repoints the entire authorization oracle for this token - `isTTMAccount`
resolves through the manager - so the change is worth an explicit log._

### MinExpirationTimestampDiffUpdated

```solidity
event MinExpirationTimestampDiffUpdated(uint256 oldDiff, uint256 newDiff)
```

Emitted when the minimum expiration timestamp difference changes.

_This is a mint-time validation rule; changing it changes which mints succeed._

### ExpirationTimestampTooSoon

```solidity
error ExpirationTimestampTooSoon(uint256 expirationTimestamp, uint256 minExpirationTimestampDiff)
```

Error for expiration timestamp too soon. It must be at least
`_minExpirationTimestampDiff` seconds in the future.

### NotTTMAccount

```solidity
error NotTTMAccount(address account)
```

Address is not a TTM Account.

#### Parameters

| Name    | Type    | Description     |
| ------- | ------- | --------------- |
| account | address | account address |

### ReservationMismatch

```solidity
error ReservationMismatch(address reservedFor, address buyer)
```

ReservedFor and buyer mismatch.

#### Parameters

| Name        | Type    | Description          |
| ----------- | ------- | -------------------- |
| reservedFor | address | reserved for address |
| buyer       | address | buyer address        |

### ReservationExpired

```solidity
error ReservationExpired(uint256 tokenId, uint256 expirationTimestamp)
```

Reservation expired.

#### Parameters

| Name                | Type    | Description          |
| ------------------- | ------- | -------------------- |
| tokenId             | uint256 | token id             |
| expirationTimestamp | uint256 | expiration timestamp |

### IncorrectPrice

```solidity
error IncorrectPrice(uint256 price, uint256 reservationPrice)
```

Incorrect price.

#### Parameters

| Name             | Type    | Description        |
| ---------------- | ------- | ------------------ |
| price            | uint256 | price of the token |
| reservationPrice | uint256 | reservation price  |

### SupplierIsNotOwner

```solidity
error SupplierIsNotOwner(uint256 tokenId, address supplier)
```

Supplier is not the owner.

#### Parameters

| Name     | Type    | Description      |
| -------- | ------- | ---------------- |
| tokenId  | uint256 | token id         |
| supplier | address | supplier address |

### TokenIsReserved

```solidity
error TokenIsReserved(uint256 tokenId, address reservedFor)
```

Token is reserved and can not be transferred.

#### Parameters

| Name        | Type    | Description          |
| ----------- | ------- | -------------------- |
| tokenId     | uint256 | token id             |
| reservedFor | address | reserved for address |

### InvalidTokenStatus

```solidity
error InvalidTokenStatus(uint256 tokenId, enum BookingToken.BookingStatus status)
```

Invalid token status.

#### Parameters

| Name    | Type                            | Description |
| ------- | ------------------------------- | ----------- |
| tokenId | uint256                         | token id    |
| status  | enum BookingToken.BookingStatus | status      |

### UnexpectedOffchainPaymentCurrency

```solidity
error UnexpectedOffchainPaymentCurrency(uint256 offchainPaymentCurrency)
```

Unexpected offchain payment currency. Thrown when offchain payment currency is provided
but payment token is not address(1).

#### Parameters

| Name                    | Type    | Description               |
| ----------------------- | ------- | ------------------------- |
| offchainPaymentCurrency | uint256 | offchain payment currency |

### UnexpectedNativePayment

```solidity
error UnexpectedNativePayment(uint256 amount)
```

Error for when there is unexpected native payment.

#### Parameters

| Name   | Type    | Description           |
| ------ | ------- | --------------------- |
| amount | uint256 | The unexpected amount |

### ZeroAddress

```solidity
error ZeroAddress()
```

A required address parameter was the zero address.

### PaymentTokenNotSupported

```solidity
error PaymentTokenNotSupported(address paymentToken)
```

The supplier has not declared this payment token as supported.

#### Parameters

| Name         | Type    | Description                             |
| ------------ | ------- | --------------------------------------- |
| paymentToken | address | The rejected payment token, or sentinel |

### onlyTTMAccount

```solidity
modifier onlyTTMAccount(address account)
```

Only TTMAccount modifier.

### constructor

```solidity
constructor() public
```

### initialize

```solidity
function initialize(address manager, address defaultAdmin, address upgrader) public
```

### \_authorizeUpgrade

```solidity
function _authorizeUpgrade(address newImplementation) internal virtual
```

Function to authorize an upgrade for UUPS proxy.

### pause

```solidity
function pause() external
```

Pauses minting, buying, and cancellation finalization.

_Pausing halts commerce (minting, buying, and cancellation
finalization), not custody: ERC-721 transfers are unaffected, so a
pending cancellation can still be auto-resolved by a transfer while
paused. This is deliberate._

### unpause

```solidity
function unpause() external
```

Resumes normal operation.

### safeMintWithReservation

```solidity
function safeMintWithReservation(address reservedFor, string uri, uint256 expirationTimestamp, uint256 price, contract IERC20 paymentToken, uint256 offchainPaymentCurrency, bool cancellable) public virtual
```

Mints a new token with a reservation for a specific address.

#### Parameters

| Name                    | Type            | Description                                                           |
| ----------------------- | --------------- | --------------------------------------------------------------------- |
| reservedFor             | address         | The TTM Account address that can buy the token                        |
| uri                     | string          | The URI of the token                                                  |
| expirationTimestamp     | uint256         | The expiration timestamp                                              |
| price                   | uint256         | The price of the token                                                |
| paymentToken            | contract IERC20 | The token used to pay for the reservation. If address(0) then native. |
| offchainPaymentCurrency | uint256         | The offchain payment currency                                         |
| cancellable             | bool            | The flag that represents whether the booking is cancellable           |

### \_reserve

```solidity
function _reserve(uint256 tokenId, address reservedFor, address supplier, uint256 expirationTimestamp, uint256 price, contract IERC20 paymentToken, uint256 offchainPaymentCurrency, bool cancellable) internal virtual
```

Reserve a token for a specific address with an expiration timestamp

### buyReservedToken

```solidity
function buyReservedToken(uint256 tokenId) public payable virtual
```

Buys a reserved token. The reservation must be for the message sender.

Also the message sender should set allowance for the payment token to this
contract to at least the reservation price. (only for ERC20 tokens)

For native coin, the message sender should send the exact amount.

Only TTM Accounts can call this function

#### Parameters

| Name    | Type    | Description  |
| ------- | ------- | ------------ |
| tokenId | uint256 | The token id |

### processPayment

```solidity
function processPayment(contract IERC20 paymentToken, uint256 paymentAmount, address recipient) internal virtual
```

### getBookingStatus

```solidity
function getBookingStatus(uint256 tokenId) public view virtual returns (enum BookingToken.BookingStatus)
```

Return booking status

#### Parameters

| Name    | Type    | Description  |
| ------- | ------- | ------------ |
| tokenId | uint256 | The token id |

#### Return Values

| Name | Type                            | Description        |
| ---- | ------------------------------- | ------------------ |
| [0]  | enum BookingToken.BookingStatus | The booking status |

### getReservationPrice

```solidity
function getReservationPrice(uint256 tokenId) public view virtual returns (uint256 price, contract IERC20 paymentToken)
```

Returns the token reservation price for a specific token.

#### Parameters

| Name    | Type    | Description  |
| ------- | ------- | ------------ |
| tokenId | uint256 | The token id |

### getReservationPaymentToken

```solidity
function getReservationPaymentToken(uint256 tokenId) external view returns (contract IERC20 paymentToken)
```

Retrieves the payment token for a given token.

#### Parameters

| Name    | Type    | Description                                    |
| ------- | ------- | ---------------------------------------------- |
| tokenId | uint256 | The token id to retrieve the payment token for |

#### Return Values

| Name         | Type            | Description       |
| ------------ | --------------- | ----------------- |
| paymentToken | contract IERC20 | The payment token |

### isCancellable

```solidity
function isCancellable(uint256 tokenId) public view virtual returns (bool)
```

Returns if the token is cancellable

#### Parameters

| Name    | Type    | Description  |
| ------- | ------- | ------------ |
| tokenId | uint256 | The token id |

### checkTransferable

```solidity
function checkTransferable(uint256 tokenId) internal virtual
```

Check if the token is transferable

### recordExpiration

```solidity
function recordExpiration(uint256 tokenId) public virtual
```

Record expiration status if the token is expired

#### Parameters

| Name    | Type    | Description  |
| ------- | ------- | ------------ |
| tokenId | uint256 | The token id |

### isTTMAccount

```solidity
function isTTMAccount(address account) public view virtual returns (bool)
```

Checks if an address is a TTM Account.

#### Parameters

| Name    | Type    | Description          |
| ------- | ------- | -------------------- |
| account | address | The address to check |

#### Return Values

| Name | Type | Description                          |
| ---- | ---- | ------------------------------------ |
| [0]  | bool | true if the address is a TTM Account |

### requireTTMAccount

```solidity
function requireTTMAccount(address account) internal view virtual
```

Checks if the address is a TTM Account and reverts if not.

#### Parameters

| Name    | Type    | Description          |
| ------- | ------- | -------------------- |
| account | address | The address to check |

### setManagerAddress

```solidity
function setManagerAddress(address manager) public virtual
```

Sets for the manager address.

#### Parameters

| Name    | Type    | Description                |
| ------- | ------- | -------------------------- |
| manager | address | The address of the manager |

### getManagerAddress

```solidity
function getManagerAddress() public view virtual returns (address)
```

Returns for the manager address.

### setMinExpirationTimestampDiff

```solidity
function setMinExpirationTimestampDiff(uint256 minExpirationTimestampDiff) public virtual
```

Sets minimum expiration timestamp difference in seconds.

#### Parameters

| Name                       | Type    | Description                                        |
| -------------------------- | ------- | -------------------------------------------------- |
| minExpirationTimestampDiff | uint256 | Minimum expiration timestamp difference in seconds |

### getMinExpirationTimestampDiff

```solidity
function getMinExpirationTimestampDiff() public view virtual returns (uint256)
```

Returns minimum expiration timestamp difference in seconds.

### initiateCancellation

```solidity
function initiateCancellation(uint256 tokenId, uint256 refundAmount, uint16 cancellationReason, uint16 cancellationReasonVersion) external virtual
```

### acceptCancellation

```solidity
function acceptCancellation(uint256 tokenId, uint256 refundAmount) external virtual
```

### counterCancellation

```solidity
function counterCancellation(uint256 tokenId, uint256 refundAmount, uint16 counterReason, uint16 counterReasonVersion) external virtual
```

### withdrawCancellation

```solidity
function withdrawCancellation(uint256 tokenId, uint16 withdrawalReason, uint16 withdrawalReasonVersion) external virtual
```

### rejectCancellation

```solidity
function rejectCancellation(uint256 tokenId, uint16 rejectionReason, uint16 rejectionReasonVersion) external virtual
```

### finalizeCancellation

```solidity
function finalizeCancellation(uint256 tokenId, uint256 checkRefundAmount) external payable virtual
```

### transferFrom

```solidity
function transferFrom(address from, address to, uint256 tokenId) public virtual
```

Override transferFrom to check if token is reserved. It reverts if
the token is reserved.

### \_update

```solidity
function _update(address to, uint256 tokenId, address auth) internal returns (address)
```

### \_increaseBalance

```solidity
function _increaseBalance(address account, uint128 value) internal
```

### tokenURI

```solidity
function tokenURI(uint256 tokenId) public view returns (string)
```

### supportsInterface

```solidity
function supportsInterface(bytes4 interfaceId) public view returns (bool)
```

## CancellationProposalStatus

```solidity
enum CancellationProposalStatus {
    NO_PROPOSAL,
    PENDING,
    REJECTED,
    WITHDRAWN,
    FINALIZED
}
```

## BookingTokenCancellable

### Proposal

```solidity
struct Proposal {
  uint256 refundAmount;
  address initialProposer;
  uint32 timesCountered;
  bool ownerAccepted;
  bool supplierAccepted;
  address currentProposer;
  uint32 timesRejected;
  enum CancellationProposalStatus status;
  uint16 cancellationReason;
  uint16 cancellationVersion;
  uint16 rejectionReason;
  uint16 rejectionVersion;
  uint16 counterReason;
  uint16 counterVersion;
  uint16 withdrawalReason;
  uint16 withdrawalVersion;
}
```

### BookingTokenCancellableStorage

```solidity
struct BookingTokenCancellableStorage {
  mapping(uint256 => struct BookingTokenCancellable.Proposal) _proposals;
}
```

### CancellationPending

```solidity
event CancellationPending(uint256 tokenId, address initialProposer, address currentProposer, uint256 refundAmount, bool ownerAccepted, bool supplierAccepted, uint32 timesCountered, uint32 timesRejected)
```

### CancellationReasons

```solidity
event CancellationReasons(uint256 tokenId, uint16 cancellationReason, uint16 cancellationReasonVersion, uint16 rejectionReason, uint16 rejectionVersion, uint16 counterReason, uint16 counterVersion, uint16 withdrawalReason, uint16 withdrawalVersion)
```

### CancellationWithdrawn

```solidity
event CancellationWithdrawn(uint256 tokenId, uint16 withdrawalReason, uint16 withdrawalVersion)
```

### CancellationRejected

```solidity
event CancellationRejected(uint256 tokenId, uint16 rejectionReason, uint16 rejectionVersion)
```

### CancellationFinalized

```solidity
event CancellationFinalized(uint256 tokenId)
```

### NotOwnerOrSupplier

```solidity
error NotOwnerOrSupplier()
```

### IncorrectRefundAmount

```solidity
error IncorrectRefundAmount(uint256 tokenId, uint256 existing, uint256 checked)
```

### InvalidCancellationProposalStatus

```solidity
error InvalidCancellationProposalStatus(uint256 tokenId, enum CancellationProposalStatus status)
```

### OnlySupplierCanFinalizeCancellation

```solidity
error OnlySupplierCanFinalizeCancellation(uint256 tokenId)
```

### OwnerNotAcceptedCancellation

```solidity
error OwnerNotAcceptedCancellation(uint256 tokenId)
```

### ProposerCanNotRejectCancellation

```solidity
error ProposerCanNotRejectCancellation(uint256 tokenId)
```

### OnlyCurrentProposerCanWithdrawCancellation

```solidity
error OnlyCurrentProposerCanWithdrawCancellation(uint256 tokenId)
```

### requireOwnerOrSupplier

```solidity
function requireOwnerOrSupplier(address owner, address supplier) internal view
```

### onlyOwnerOrSupplier

```solidity
modifier onlyOwnerOrSupplier(address owner, address supplier)
```

### \_getCancellationProposalStatusAndCurrentProposer

```solidity
function _getCancellationProposalStatusAndCurrentProposer(uint256 tokenId) internal view returns (enum CancellationProposalStatus status, address currentProposer)
```

### getCancellationProposal

```solidity
function getCancellationProposal(uint256 tokenId) external view returns (enum CancellationProposalStatus, uint256 refundAmount, address initialProposer, address currentProposer, bool ownerAccepted, bool supplierAccepted, uint32 timesCountered, uint32 timesRejected)
```

### getCancellationReasons

```solidity
function getCancellationReasons(uint256 tokenId) external view returns (uint16 cancellationReason, uint16 cancellationVersion, uint16 rejectionReason, uint16 rejectionVersion, uint16 counterReason, uint16 counterVersion, uint16 withdrawalReason, uint16 withdrawalVersion)
```

### \_initiateCancellation

```solidity
function _initiateCancellation(address owner, address supplier, uint256 tokenId, uint256 refundAmount, uint16 cancellationReason, uint16 cancellationReasonVersion) internal virtual
```

### \_acceptCancellation

```solidity
function _acceptCancellation(address owner, address supplier, uint256 tokenId, uint256 checkRefundAmount) internal virtual
```

Used by the owner or supplier to accept a cancellation proposal that
is initiated or countered by the other party

#### Parameters

| Name              | Type    | Description                                              |
| ----------------- | ------- | -------------------------------------------------------- |
| owner             | address | Owner of the token                                       |
| supplier          | address | Supplier of the token                                    |
| tokenId           | uint256 | Token ID                                                 |
| checkRefundAmount | uint256 | Refund amount to check against, to prevent front-running |

### \_counterCancellation

```solidity
function _counterCancellation(address owner, address supplier, uint256 tokenId, uint256 refundAmount, uint16 counterReason, uint16 counterVersion) internal virtual
```

### \_withdrawCancellation

```solidity
function _withdrawCancellation(address actor, address owner, address supplier, uint256 tokenId, uint16 withdrawalReason, uint16 withdrawalVersion) internal virtual
```

### \_rejectCancellation

```solidity
function _rejectCancellation(address actor, address owner, address supplier, uint256 tokenId, uint16 rejectionReason, uint16 rejectionReasonVersion) internal virtual
```

### \_finalizeCancellation

```solidity
function _finalizeCancellation(address supplier, uint256 tokenId, uint256 checkRefundAmount) internal virtual returns (uint256 refundAmount)
```

## BookingTokenOperator

Booking token operator contract is used by the {TTMAccount} contract to mint
and buy booking tokens.

We made this a library so that we can use it in the {TTMAccount} contract without
increasing the size of the contract.

### NATIVE_PAYMENT

```solidity
address NATIVE_PAYMENT
```

Tokens are directly transferred to the recipient.

_Special address for native payments._

### OFFCHAIN_PAYMENT

```solidity
address OFFCHAIN_PAYMENT
```

A third-party service is used to handle payments.

_Special address for offchain payments._

### UnexpectedPrice

```solidity
error UnexpectedPrice(uint256 tokenId, uint256 actualPrice, uint256 expectedPrice)
```

### UnexpectedPaymentToken

```solidity
error UnexpectedPaymentToken(uint256 tokenId, contract IERC20 actualPaymentToken, contract IERC20 expectedPaymentToken)
```

### mintBookingToken

```solidity
function mintBookingToken(address bookingToken, address reservedFor, string uri, uint256 expirationTimestamp, uint256 price, contract IERC20 paymentToken, uint256 offchainPaymentCurrency, bool cancellable) public
```

_Mints a booking token with offchain payment currency and cancellable support._

#### Parameters

| Name                    | Type            | Description                                                                   |
| ----------------------- | --------------- | ----------------------------------------------------------------------------- |
| bookingToken            | address         | booking token contract address                                                |
| reservedFor             | address         | address of the TTM Account that can buy the token (generally the distributor) |
| uri                     | string          | URI of the token                                                              |
| expirationTimestamp     | uint256         | expiration timestamp of the token in seconds                                  |
| price                   | uint256         | price of the token                                                            |
| paymentToken            | contract IERC20 | payment token address                                                         |
| offchainPaymentCurrency | uint256         | payment token address                                                         |
| cancellable             | bool            | cancellable flag                                                              |

### buyBookingToken

```solidity
function buyBookingToken(address bookingToken, uint256 tokenId, uint256 expectedPrice, contract IERC20 expectedPaymentToken) public
```

_Buys a booking token with the specified price and payment token in the
reservation._

#### Parameters

| Name                 | Type            | Description                    |
| -------------------- | --------------- | ------------------------------ |
| bookingToken         | address         | booking token contract address |
| tokenId              | uint256         | token id                       |
| expectedPrice        | uint256         |                                |
| expectedPaymentToken | contract IERC20 |                                |

### recordExpiration

```solidity
function recordExpiration(address bookingToken, uint256 tokenId) public
```

Record the expiration of a booking token.

#### Parameters

| Name         | Type    | Description                    |
| ------------ | ------- | ------------------------------ |
| bookingToken | address | booking token contract address |
| tokenId      | uint256 | token id                       |

### initiateCancellation

```solidity
function initiateCancellation(address bookingToken, uint256 tokenId, uint256 refundAmount, uint16 cancellationReason, uint16 cancellationReasonVersion) external
```

Initiates a cancellation proposal for a bought token.

#### Parameters

| Name                      | Type    | Description                    |
| ------------------------- | ------- | ------------------------------ |
| bookingToken              | address | booking token contract address |
| tokenId                   | uint256 | token id                       |
| refundAmount              | uint256 | proposed refund amount         |
| cancellationReason        | uint16  | cancellation reason            |
| cancellationReasonVersion | uint16  | cancellation reason version    |

### acceptCancellation

```solidity
function acceptCancellation(address bookingToken, uint256 tokenId, uint256 refundAmount) external
```

Sets accepted by the owner or supplier flag for a cancellation proposal for a bought token.

#### Parameters

| Name         | Type    | Description                                                          |
| ------------ | ------- | -------------------------------------------------------------------- |
| bookingToken | address |                                                                      |
| tokenId      | uint256 | The token id to accept the cancellation for                          |
| refundAmount | uint256 | The refund amount to check, this is to prevent front-running attacks |

### counterCancellation

```solidity
function counterCancellation(address bookingToken, uint256 tokenId, uint256 refundAmount, uint16 counterReason, uint16 counterReasonVersion) public
```

Counters a cancellation proposal.

#### Parameters

| Name                 | Type    | Description                    |
| -------------------- | ------- | ------------------------------ |
| bookingToken         | address | booking token contract address |
| tokenId              | uint256 | token id                       |
| refundAmount         | uint256 | proposed refund amount         |
| counterReason        | uint16  |                                |
| counterReasonVersion | uint16  |                                |

### withdrawCancellation

```solidity
function withdrawCancellation(address bookingToken, uint256 tokenId, uint16 reason, uint16 reasonVersion) public
```

Withdraws a cancellation proposal.

#### Parameters

| Name          | Type    | Description                                                                   |
| ------------- | ------- | ----------------------------------------------------------------------------- |
| bookingToken  | address | booking token contract address                                                |
| tokenId       | uint256 | token id for which to withdraw the proposal                                   |
| reason        | uint16  | The reason for withdrawing the proposal                                       |
| reasonVersion | uint16  | The version of the withdrawal reason from the Travel Token Messenger Protocol |

### rejectCancellation

```solidity
function rejectCancellation(address bookingToken, uint256 tokenId, uint16 rejectionReason, uint16 rejectionReasonVersion) external
```

Reject a cancellation proposal for a bought token.

#### Parameters

| Name                   | Type    | Description                                                                   |
| ---------------------- | ------- | ----------------------------------------------------------------------------- |
| bookingToken           | address | booking token contract address                                                |
| tokenId                | uint256 | The token id to reject the cancellation for                                   |
| rejectionReason        | uint16  | The reason for rejecting the cancellation                                     |
| rejectionReasonVersion | uint16  | Version of the rejection reason enum from the Travel Token Messenger Protocol |

### finalizeCancellation

```solidity
function finalizeCancellation(address bookingToken, uint256 tokenId, uint256 refundAmount) public
```

Finalizes a cancellation proposal by transferring the refund amount
to the Booking Token contract.

#### Parameters

| Name         | Type    | Description                                                          |
| ------------ | ------- | -------------------------------------------------------------------- |
| bookingToken | address | BookingToken contract address                                        |
| tokenId      | uint256 | The token id for which to finalize the proposal                      |
| refundAmount | uint256 | The refund amount to check, this is to prevent front-running attacks |

## IBookingToken

### safeMintWithReservation

```solidity
function safeMintWithReservation(address reservedFor, string uri, uint256 expirationTimestamp, uint256 price, contract IERC20 paymentToken, uint256 offchainPaymentCurrency, bool isCancellable) external
```

### buyReservedToken

```solidity
function buyReservedToken(uint256 tokenId) external payable
```

### getReservationPrice

```solidity
function getReservationPrice(uint256 tokenId) external view returns (uint256 price, contract IERC20 paymentToken)
```

### getReservationPaymentToken

```solidity
function getReservationPaymentToken(uint256 tokenId) external view returns (contract IERC20 paymentToken)
```

### recordExpiration

```solidity
function recordExpiration(uint256 tokenId) external
```

Record expiration status if the token is expired

#### Parameters

| Name    | Type    | Description                       |
| ------- | ------- | --------------------------------- |
| tokenId | uint256 | The token id to record as expired |

### initiateCancellation

```solidity
function initiateCancellation(uint256 tokenId, uint256 refundAmount, uint16 cancellationReason, uint16 cancellationReasonVersion) external
```

Initiates a cancellation for a bought token.

#### Parameters

| Name                      | Type    | Description                                   |
| ------------------------- | ------- | --------------------------------------------- |
| tokenId                   | uint256 | The token id to initiate the cancellation for |
| refundAmount              | uint256 | The proposed refund amount in wei             |
| cancellationReason        | uint16  | The reason for cancellation                   |
| cancellationReasonVersion | uint16  | The version of the cancellation reason        |

### acceptCancellation

```solidity
function acceptCancellation(uint256 tokenId, uint256 refundAmount) external
```

Sets accepted by the owner or supplier flag for a cancellation proposal for a bought token.

#### Parameters

| Name         | Type    | Description                                                          |
| ------------ | ------- | -------------------------------------------------------------------- |
| tokenId      | uint256 | The token id to accept the cancellation for                          |
| refundAmount | uint256 | The refund amount to check, this is to prevent front-running attacks |

### rejectCancellation

```solidity
function rejectCancellation(uint256 tokenId, uint16 rejectionReason, uint16 rejectionReasonVersion) external
```

Reject a cancellation proposal for a bought token.

#### Parameters

| Name                   | Type    | Description                                 |
| ---------------------- | ------- | ------------------------------------------- |
| tokenId                | uint256 | The token id to reject the cancellation for |
| rejectionReason        | uint16  | The reason for rejection                    |
| rejectionReasonVersion | uint16  | The version of the rejection reason         |

### counterCancellation

```solidity
function counterCancellation(uint256 tokenId, uint256 refundAmount, uint16 counterReason, uint16 counterReasonVersion) external
```

Counters a cancellation proposal with a new proposal.

#### Parameters

| Name                 | Type    | Description                                                          |
| -------------------- | ------- | -------------------------------------------------------------------- |
| tokenId              | uint256 | The token id to counter the cancellation for                         |
| refundAmount         | uint256 | The refund amount to check, this is to prevent front-running attacks |
| counterReason        | uint16  | The reason for the counter                                           |
| counterReasonVersion | uint16  | The version of the counter reason                                    |

### withdrawCancellation

```solidity
function withdrawCancellation(uint256 tokenId, uint16 withdrawalReason, uint16 withdrawalReasonVersion) external
```

Withdraws an active cancellation proposal. Only the current proposer of the proposal can withdraw.

#### Parameters

| Name                    | Type    | Description                                     |
| ----------------------- | ------- | ----------------------------------------------- |
| tokenId                 | uint256 | The token id for which to withdraw the proposal |
| withdrawalReason        | uint16  | The reason for withdrawing the proposal         |
| withdrawalReasonVersion | uint16  | The version of the withdrawal reason            |

### finalizeCancellation

```solidity
function finalizeCancellation(uint256 tokenId, uint256 refundAmount) external payable
```

Finalizes a cancellation proposal. Only the supplier of the token can finalize.

#### Parameters

| Name         | Type    | Description                                                          |
| ------------ | ------- | -------------------------------------------------------------------- |
| tokenId      | uint256 | The token id for which to finalize the proposal                      |
| refundAmount | uint256 | The refund amount to check, this is to prevent front-running attacks |

## ITTMAccountManager

### getAccountImplementation

```solidity
function getAccountImplementation() external view returns (address)
```

### isTTMAccount

```solidity
function isTTMAccount(address account) external view returns (bool)
```

### getRegisteredServiceNameByHash

```solidity
function getRegisteredServiceNameByHash(bytes32 serviceHash) external view returns (string serviceName)
```

_Reverts if the hash is not registered. This is the sole remaining
manager dependency of TTMAccount, used to validate {addService}._

## TTMAccountManager

This contract manages the creation of the Travel Token Messenger accounts by
deploying {ERC1967Proxy} proxies that point to the{TTMAccount} implementation
address.

Create TTM Account: Users who want to create an account should call
`createTTMAccount(address admin, address upgrader)` function with addresses of
the accounts admin and upgrader roles.

When the manager contract is paused, account creation is stopped.

Service Registry: {TTMAccountManager} also acts as a registry for the services
that {TTMAccount} contracts add as a supported or wanted service. Registry
works by hashing (keccak256) the service name (string) and creating a mapping
as keccak256(serviceName) => serviceName. And provides functions that
{TTMAccount} function uses to register services. The {TTMAccount} only keeps
the hashes (byte32) of the registered services.

### PAUSER_ROLE

```solidity
bytes32 PAUSER_ROLE
```

Pauser role can pause the contract. Currently this only affects the
creation of TTM Accounts. When paused, account creation is stopped.

### UPGRADER_ROLE

```solidity
bytes32 UPGRADER_ROLE
```

Upgrader role can upgrade the contract to a new implementation.

### VERSIONER_ROLE

```solidity
bytes32 VERSIONER_ROLE
```

Versioner role can set new {TTMAccount} implementation address. When a
new implementation address is set, it is used for the new {TTMAccount}
creations.

The old {TTMAccount} contracts are not affected by this. Owners of those
should do the upgrade manually by calling the `upgradeToAndCall(address)`
function on the account.

### SERVICE_REGISTRY_ADMIN_ROLE

```solidity
bytes32 SERVICE_REGISTRY_ADMIN_ROLE
```

Service registry admin role can add and remove services to the service
registry mapping. Implemented by {ServiceRegistry} contract.

### TTMAccountManagerStorage

```solidity
struct TTMAccountManagerStorage {
  address _latestAccountImplementation;
  address _bookingToken;
  struct EnumerableSet.AddressSet _ttmAccounts;
  mapping(address => address) _ttmAccountCreator;
}
```

### TTMAccountCreated

```solidity
event TTMAccountCreated(address account, address creator, address admin)
```

Emitted when a TTM Account is created.

_Carries creator and admin so indexers need no follow-up call per account._

#### Parameters

| Name    | Type    | Description                                  |
| ------- | ------- | -------------------------------------------- |
| account | address | The address of the newly created TTM Account |
| creator | address | The address that called {createTTMAccount}   |
| admin   | address | The admin address granted on the new account |

### TTMAccountImplementationUpdated

```solidity
event TTMAccountImplementationUpdated(address oldImplementation, address newImplementation)
```

TTM Account implementation address updated event.

#### Parameters

| Name              | Type    | Description                    |
| ----------------- | ------- | ------------------------------ |
| oldImplementation | address | The old implementation address |
| newImplementation | address | The new implementation address |

### BookingTokenAddressUpdated

```solidity
event BookingTokenAddressUpdated(address oldBookingToken, address newBookingToken)
```

Booking token address updated event.

#### Parameters

| Name            | Type    | Description                   |
| --------------- | ------- | ----------------------------- |
| oldBookingToken | address | The old booking token address |
| newBookingToken | address | The new booking token address |

### TTMAccountInvalidImplementation

```solidity
error TTMAccountInvalidImplementation(address implementation)
```

The implementation of the TTMAccount is invalid.

#### Parameters

| Name           | Type    | Description                                  |
| -------------- | ------- | -------------------------------------------- |
| implementation | address | The implementation address of the TTMAccount |

### TTMAccountInvalidAdmin

```solidity
error TTMAccountInvalidAdmin(address admin)
```

The admin address is invalid.

#### Parameters

| Name  | Type    | Description       |
| ----- | ------- | ----------------- |
| admin | address | The admin address |

### InvalidBookingTokenAddress

```solidity
error InvalidBookingTokenAddress(address bookingToken)
```

Invalid booking token address.

#### Parameters

| Name         | Type    | Description               |
| ------------ | ------- | ------------------------- |
| bookingToken | address | The booking token address |

### ZeroAddress

```solidity
error ZeroAddress()
```

A required address parameter was the zero address.

### constructor

```solidity
constructor() public
```

### initialize

```solidity
function initialize(address defaultAdmin, address pauser, address upgrader, address versioner) public
```

### pause

```solidity
function pause() public
```

Pauses the TTMAccountManager contract. Currently this only affects the
creation of TTMAccount. When paused, account creation is stopped.

### unpause

```solidity
function unpause() public
```

Unpauses the TTMAccountManager contract.

### \_authorizeUpgrade

```solidity
function _authorizeUpgrade(address newImplementation) internal
```

Authorization for the TTMAccountManager contract upgrade.

### createTTMAccount

```solidity
function createTTMAccount(address admin, address upgrader) external payable returns (address)
```

Creates a new TTMAccount.

This function is currently permissionless: any address may create an
account. See docs/decisions/2026-07-21-contract-design-decisions.md
(Decision 1) -- gating must be resolved before Base mainnet.

_Emits a {TTMAccountCreated} event._

### getTTMAccountCreator

```solidity
function getTTMAccountCreator(address account) public view returns (address)
```

Returns the given account's creator, or the zero address if the
address is not a TTM Account.

#### Parameters

| Name    | Type    | Description         |
| ------- | ------- | ------------------- |
| account | address | The account address |

### isTTMAccount

```solidity
function isTTMAccount(address account) public view returns (bool)
```

Returns whether the given address is a TTM Account created by this manager.

#### Parameters

| Name    | Type    | Description          |
| ------- | ------- | -------------------- |
| account | address | The address to check |

### getTTMAccountCount

```solidity
function getTTMAccountCount() public view returns (uint256)
```

Returns the number of TTM Accounts created by this manager.

### getTTMAccounts

```solidity
function getTTMAccounts() public view returns (address[])
```

Returns every TTM Account created by this manager.

_Unbounded. Prefer {getTTMAccountsSlice} against a public RPC once the
ecosystem grows past a few hundred accounts._

### getTTMAccountsSlice

```solidity
function getTTMAccountsSlice(uint256 offset, uint256 limit) public view returns (address[] accounts)
```

Returns a bounded window of TTM Accounts, for callers that cannot
afford an unbounded read.

Returns an empty array if `offset` is at or past the end. The window is
clamped to the end of the set, so a `limit` larger than the remainder is
not an error.

#### Parameters

| Name   | Type    | Description                          |
| ------ | ------- | ------------------------------------ |
| offset | uint256 | Index to start at                    |
| limit  | uint256 | Maximum number of accounts to return |

### getAccountImplementation

```solidity
function getAccountImplementation() public view returns (address)
```

Returns the TTMAccount implementation address.

### setAccountImplementation

```solidity
function setAccountImplementation(address newImplementation) public
```

Set a new TTMAccount implementation address.

#### Parameters

| Name              | Type    | Description                    |
| ----------------- | ------- | ------------------------------ |
| newImplementation | address | The new implementation address |

### \_setAccountImplementation

```solidity
function _setAccountImplementation(address newImplementation) internal
```

### getBookingTokenAddress

```solidity
function getBookingTokenAddress() public view returns (address)
```

Returns the booking token address.

### setBookingTokenAddress

```solidity
function setBookingTokenAddress(address token) public
```

Sets booking token address.

### \_setBookingTokenAddress

```solidity
function _setBookingTokenAddress(address token) internal
```

### registerService

```solidity
function registerService(string serviceName) public
```

Registers a given service name. TTM Accounts can only register services
if they are also registered in the service registry on the manager contract.

#### Parameters

| Name        | Type   | Description         |
| ----------- | ------ | ------------------- |
| serviceName | string | Name of the service |

### unregisterService

```solidity
function unregisterService(string serviceName) public
```

Unregisters a given service name. TTM Accounts will not be able to register
the service anymore.

#### Parameters

| Name        | Type   | Description         |
| ----------- | ------ | ------------------- |
| serviceName | string | Name of the service |

## TTMAccountManagerTest

### getVersion

```solidity
function getVersion() public pure returns (string)
```

## PartnerConfiguration

Partner Configuration is used by the {TTMAccount} contract to register
supported and wanted services by the partner.

### Service

Struct for storing supported service details for suppliers

```solidity
struct Service {
    bool _restrictedRate;
    string[] _capabilities;
}
```

### PaymentInfo

```solidity
struct PaymentInfo {
  struct EnumerableSet.AddressSet _supportedTokens;
}
```

### PartnerConfigurationStorage

```solidity
struct PartnerConfigurationStorage {
  struct EnumerableSet.Bytes32Set _servicesHashSet;
  mapping(bytes32 => struct PartnerConfiguration.Service) _supportedServices;
  struct PartnerConfiguration.PaymentInfo _paymentInfo;
  struct EnumerableSet.AddressSet _publicKeyAddressesSet;
  mapping(address => bytes) _publicKeys;
  struct EnumerableSet.Bytes32Set _wantedServicesHashSet;
}
```

### ServiceAlreadyExists

```solidity
error ServiceAlreadyExists(bytes32 serviceHash)
```

### ServiceDoesNotExist

```solidity
error ServiceDoesNotExist(bytes32 serviceHash)
```

### WantedServiceAlreadyExists

```solidity
error WantedServiceAlreadyExists(bytes32 serviceHash)
```

### WantedServiceDoesNotExist

```solidity
error WantedServiceDoesNotExist(bytes32 serviceHash)
```

### PaymentTokenAlreadyExists

```solidity
error PaymentTokenAlreadyExists(address token)
```

### PaymentTokenDoesNotExist

```solidity
error PaymentTokenDoesNotExist(address token)
```

### PublicKeyAlreadyExists

```solidity
error PublicKeyAlreadyExists(address pubKeyAddress)
```

### PublicKeyDoesNotExist

```solidity
error PublicKeyDoesNotExist(address pubKeyAddress)
```

### CapabilityDoesNotExist

```solidity
error CapabilityDoesNotExist(bytes32 serviceHash, string capability)
```

### PaymentTokenAdded

```solidity
event PaymentTokenAdded(address token)
```

### PaymentTokenRemoved

```solidity
event PaymentTokenRemoved(address token)
```

### PublicKeyAdded

```solidity
event PublicKeyAdded(address pubKeyAddress)
```

### PublicKeyRemoved

```solidity
event PublicKeyRemoved(address pubKeyAddress)
```

### \_\_PartnerConfiguration_init

```solidity
function __PartnerConfiguration_init() internal
```

### \_\_PartnerConfiguration_init_unchained

```solidity
function __PartnerConfiguration_init_unchained() internal
```

### \_addService

```solidity
function _addService(bytes32 serviceHash, string[] capabilities, bool restrictedRate) internal virtual
```

Adds a supported Service object for a given hash.

#### Parameters

| Name           | Type     | Description                                   |
| -------------- | -------- | --------------------------------------------- |
| serviceHash    | bytes32  | Hash of the service                           |
| capabilities   | string[] | Capabilities for the service                  |
| restrictedRate | bool     | If the service is restricted to pre-agreement |

### \_removeService

```solidity
function _removeService(bytes32 serviceHash) internal virtual
```

Removes a supported Service object for a given hash.

#### Parameters

| Name        | Type    | Description         |
| ----------- | ------- | ------------------- |
| serviceHash | bytes32 | Hash of the service |

### \_setServiceRestrictedRate

```solidity
function _setServiceRestrictedRate(bytes32 serviceHash, bool restrictedRate) internal virtual
```

Sets the Service restricted rate for a given hash.

#### Parameters

| Name           | Type    | Description         |
| -------------- | ------- | ------------------- |
| serviceHash    | bytes32 | Hash of the service |
| restrictedRate | bool    | Restricted rate     |

### \_setServiceCapabilities

```solidity
function _setServiceCapabilities(bytes32 serviceHash, string[] capabilities) internal virtual
```

Sets the Service capabilities for a given hash.

#### Parameters

| Name         | Type     | Description         |
| ------------ | -------- | ------------------- |
| serviceHash  | bytes32  | Hash of the service |
| capabilities | string[] | Capabilities        |

### \_addServiceCapability

```solidity
function _addServiceCapability(bytes32 serviceHash, string capability) internal virtual
```

Adds a capability to the service.

#### Parameters

| Name        | Type    | Description         |
| ----------- | ------- | ------------------- |
| serviceHash | bytes32 | Hash of the service |
| capability  | string  | Capability          |

### \_removeServiceCapability

```solidity
function _removeServiceCapability(bytes32 serviceHash, string capability) internal virtual
```

Removes a capability from the service.

#### Parameters

| Name        | Type    | Description         |
| ----------- | ------- | ------------------- |
| serviceHash | bytes32 | Hash of the service |
| capability  | string  | Capability          |

### getAllServiceHashes

```solidity
function getAllServiceHashes() public view returns (bytes32[] serviceHashes)
```

Returns all supported service hashes.

### getService

```solidity
function getService(bytes32 serviceHash) public view virtual returns (struct PartnerConfiguration.Service service)
```

Returns the Service object for a given hash. Service object contains fee and capabilities.

`serviceHash` is keccak256 hash of the pkg + service name as:

```text
           ┌────────────── pkg ─────────────┐ ┌───── service name ─────┐
keccak256("ttm.services.accommodation.v1alpha.AccommodationSearchService")
```

_These services are coming from the Travel Token Messenger Protocol's protobuf
definitions._

#### Parameters

| Name        | Type    | Description         |
| ----------- | ------- | ------------------- |
| serviceHash | bytes32 | Hash of the service |

### \_isServiceSupported

```solidity
function _isServiceSupported(bytes32 serviceHash) internal view returns (bool)
```

Checks if the service is supported.

#### Parameters

| Name        | Type    | Description         |
| ----------- | ------- | ------------------- |
| serviceHash | bytes32 | Hash of the service |

### getServiceRestrictedRate

```solidity
function getServiceRestrictedRate(bytes32 serviceHash) public view virtual returns (bool restrictedRate)
```

Returns the restricted rate for a given service hash.

#### Parameters

| Name        | Type    | Description         |
| ----------- | ------- | ------------------- |
| serviceHash | bytes32 | Hash of the service |

### getServiceCapabilities

```solidity
function getServiceCapabilities(bytes32 serviceHash) public view virtual returns (string[] capabilities)
```

Returns the capabilities for a given service hash.

#### Parameters

| Name        | Type    | Description         |
| ----------- | ------- | ------------------- |
| serviceHash | bytes32 | Hash of the service |

### \_addWantedService

```solidity
function _addWantedService(bytes32 serviceHash) internal virtual
```

Adds a wanted service hash to the wanted services set.

Reverts if the service already exists.

#### Parameters

| Name        | Type    | Description         |
| ----------- | ------- | ------------------- |
| serviceHash | bytes32 | Hash of the service |

### \_removeWantedService

```solidity
function _removeWantedService(bytes32 serviceHash) internal virtual
```

Removes a wanted service hash from the wanted services set.

Reverts if the service does not exist.

#### Parameters

| Name        | Type    | Description         |
| ----------- | ------- | ------------------- |
| serviceHash | bytes32 | Hash of the service |

### getWantedServiceHashes

```solidity
function getWantedServiceHashes() public view virtual returns (bytes32[] serviceHashes)
```

Returns all wanted service hashes.

#### Return Values

| Name          | Type      | Description           |
| ------------- | --------- | --------------------- |
| serviceHashes | bytes32[] | Wanted service hashes |

### \_addSupportedToken

```solidity
function _addSupportedToken(address _token) internal virtual
```

Adds a supported payment token.

#### Parameters

| Name    | Type    | Description                       |
| ------- | ------- | --------------------------------- |
| \_token | address | Payment token address to be added |

### \_removeSupportedToken

```solidity
function _removeSupportedToken(address _token) internal virtual
```

Removes a supported payment token.

#### Parameters

| Name    | Type    | Description                         |
| ------- | ------- | ----------------------------------- |
| \_token | address | Payment token address to be removed |

### getSupportedTokens

```solidity
function getSupportedTokens() public view virtual returns (address[] tokens)
```

Returns supported token addresses.

#### Return Values

| Name   | Type      | Description               |
| ------ | --------- | ------------------------- |
| tokens | address[] | Supported token addresses |

### isSupportedToken

```solidity
function isSupportedToken(address _token) public view virtual returns (bool supported)
```

Returns whether a payment token is declared as supported.

The two sentinel values are legitimate members of this set:
`address(0)` means native currency and `address(1)` means off-chain
payment, matching how `BookingToken` encodes payment mode.

#### Parameters

| Name    | Type    | Description                          |
| ------- | ------- | ------------------------------------ |
| \_token | address | Payment token address, or a sentinel |

#### Return Values

| Name      | Type | Description                   |
| --------- | ---- | ----------------------------- |
| supported | bool | Whether the token is declared |

### \_addPublicKey

```solidity
function _addPublicKey(address pubKeyAddress, bytes publicKeyData) internal virtual
```

Adds public key with an address. Reverts if the public key already
exists.

Beware: This functions does not check if the public key is actually for the
given address.

### \_removePublicKey

```solidity
function _removePublicKey(address pubKeyAddress) internal virtual
```

Removes the public key for a given address

Reverts if the public key does not exist

### getPublicKeysAddresses

```solidity
function getPublicKeysAddresses() public view virtual returns (address[] pubKeyAddresses)
```

Returns the addresses of all public keys. These can then be used to
retrieve the public keys the `getPublicKey(address)` function.

### getPublicKey

```solidity
function getPublicKey(address pubKeyAddress) public view virtual returns (bytes data)
```

Returns the public key for a given address.

Reverts if the public key does not exist

#### Parameters

| Name          | Type    | Description               |
| ------------- | ------- | ------------------------- |
| pubKeyAddress | address | Address of the public key |

## ServiceRegistry

Service registry is used by the {TTMAccountManager} contract to register
services by hashing (keccak256) the service name (string) and creating a mapping
as keccak256(serviceName) => serviceName.

### ServiceRegistryStorage

```solidity
struct ServiceRegistryStorage {
  struct EnumerableSet.Bytes32Set _servicesHashSet;
  mapping(bytes32 => string) _serviceNameByHash;
  mapping(string => bytes32) _hashByServiceName;
}
```

### ServiceRegistered

```solidity
event ServiceRegistered(bytes32 serviceHash, string serviceName)
```

Emitted when a service is registered.

_The hash is indexed for filtering; the name travels in the data section so
consumers can build a complete name-to-hash map from logs alone, with no
`eth_call`. This is the authoritative publication of that mapping - `TTMAccount`
emits hashes only._

### ServiceUnregistered

```solidity
event ServiceUnregistered(bytes32 serviceHash, string serviceName)
```

Emitted when a service is unregistered.

_Existing accounts can still resolve a deprecated name, so this is the only
signal that a service was retired. See {\_unregisterServiceName}._

### ServiceAlreadyRegistered

```solidity
error ServiceAlreadyRegistered(string serviceName)
```

### ServiceNotRegistered

```solidity
error ServiceNotRegistered()
```

### \_\_ServiceRegistry_init

```solidity
function __ServiceRegistry_init() internal
```

### \_\_ServiceRegistry_init_unchained

```solidity
function __ServiceRegistry_init_unchained() internal
```

### \_registerServiceName

```solidity
function _registerServiceName(string serviceName) internal virtual
```

Adds a new service by its name. This function calculates the hash of the
service name and adds it to the registry

{serviceName} is the pkg + service name as:

```text
 ┌────────────── pkg ─────────────┐ ┌───── service name ─────┐
"ttm.services.accommodation.v1alpha.AccommodationSearchService"
```

_These services are coming from the Travel Token Messenger Protocol's protobuf
definitions._

#### Parameters

| Name        | Type   | Description         |
| ----------- | ------ | ------------------- |
| serviceName | string | Name of the service |

### \_unregisterServiceName

```solidity
function _unregisterServiceName(string serviceName) internal virtual
```

Removes a service by its name. This function calculates the hash of the
service name and removes it from the registry.

#### Parameters

| Name        | Type   | Description         |
| ----------- | ------ | ------------------- |
| serviceName | string | Name of the service |

### getRegisteredServiceNameByHash

```solidity
function getRegisteredServiceNameByHash(bytes32 serviceHash) public view returns (string serviceName)
```

Returns the name of a registered service by its hash.

#### Parameters

| Name        | Type    | Description         |
| ----------- | ------- | ------------------- |
| serviceHash | bytes32 | Hash of the service |

### getServiceNameByHash

```solidity
function getServiceNameByHash(bytes32 serviceHash) public view returns (string serviceName)
```

Returns the name of a service by its hash. Even if the service is unregistered at the moment.

#### Parameters

| Name        | Type    | Description         |
| ----------- | ------- | ------------------- |
| serviceHash | bytes32 | Hash of the service |

### getRegisteredServiceHashByName

```solidity
function getRegisteredServiceHashByName(string serviceName) public view returns (bytes32 serviceHash)
```

Returns the hash of a service by its name.

#### Parameters

| Name        | Type   | Description         |
| ----------- | ------ | ------------------- |
| serviceName | string | Name of the service |

### getServiceHashByName

```solidity
function getServiceHashByName(string serviceName) public view returns (bytes32 serviceHash)
```

Returns the hash of a service by its name. Even if the service is unregistered at the moment.

#### Parameters

| Name        | Type   | Description         |
| ----------- | ------ | ------------------- |
| serviceName | string | Name of the service |

### getAllRegisteredServiceHashes

```solidity
function getAllRegisteredServiceHashes() public view returns (bytes32[] services)
```

Returns all registered service **hashes**.

#### Return Values

| Name     | Type      | Description                   |
| -------- | --------- | ----------------------------- |
| services | bytes32[] | All registered service hashes |

### getAllRegisteredServiceNames

```solidity
function getAllRegisteredServiceNames() public view returns (string[] services)
```

Returns all registered service **names**.

#### Return Values

| Name     | Type     | Description                  |
| -------- | -------- | ---------------------------- |
| services | string[] | All registered service names |

## Dummy

### getVersion

```solidity
function getVersion() public pure returns (string)
```

## MockSafe

Minimal stand-in for a Gnosis Safe: the two owner-set getters the
`roles handoff` preflight probes, over a `SafeProxy`-shaped storage layout.

_`singleton` is declared first so it occupies storage slot 0, exactly as
`SafeProxy` does — that is the slot the preflight reads to report which Safe
implementation the address is running. Used to exercise the custody-type
preflight in `tasks/lib/preflight.js` without pulling the Safe contracts into
this repo's dependency tree._

### singleton

```solidity
address singleton
```

_Slot 0, mirroring `SafeProxy.singleton`._

### constructor

```solidity
constructor(address singleton_, address[] owners_, uint256 threshold_) public
```

### getOwners

```solidity
function getOwners() external view returns (address[])
```

Mirrors `Safe.getOwners()`.

### getThreshold

```solidity
function getThreshold() external view returns (uint256)
```

Mirrors `Safe.getThreshold()`.

## NullUSD

### constructor

```solidity
constructor() public
```

## RejectsEther

A contract that cannot receive ETH: no `receive`, no `fallback`.

_Used to exercise the cancellation refund path against a counterparty whose
address rejects a plain transfer. See Decision 3 in
docs/decisions/2026-07-21-contract-design-decisions.md._

### ping

```solidity
function ping() external pure returns (bool)
```

Lets tests confirm the contract deployed and has code.
