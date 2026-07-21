# Technical Backlog — Deferred Engineering Work

**Date:** 2026-07-21
**Audience:** Engineering only
**Status:** Deferred — seed material for a follow-up spec

## Scope

Work deferred from the 2026-07-21 pre-deployment review that requires **no
business or management decision**. Every item here is an engineering call that
can be scheduled whenever capacity allows.

These items are **not blocked on the decision document**
(`2026-07-21-contract-design-decisions.md`). Any of them could form a Spec 2
immediately after the Base Sepolia deployment lands.

The main cost of deferral is that consumers — the bot, matrix-app-service, and
UI — must be updated in step with items 1 and 3.

---

## 1. Event signature rework

**Impact:** breaking for log consumers · **Effort:** S in contracts, M in consumers

### 1.1 Service events lose the service name

`contracts/account/TTMAccount.sol:167-177` — seven events declare
`string indexed serviceName`.

Indexing a dynamic type stores only `keccak256(value)` in the topic and
**nothing in the data section**. The service name is not recoverable from the
log. The UI already carries machinery to work around this: `SERVICE_HASH_EVENTS`
in `ui/src/lib/activity/catalog.ts` plus `ui/src/hooks/useAccountActivity.ts:32-65`,
which re-resolves the hash against the registry before it can render a sentence.

That workaround happens to function because the topic hash is
`keccak256(abi.encodePacked(name))`, numerically equal to the registry's
`serviceHash` — a coincidence of both using the same hashing, not a designed
invariant.

**Change:**

```solidity
event ServiceAdded(bytes32 indexed serviceHash, string serviceName);
```

Indexed for filtering, plain for reading, no resolution round-trip. Same for
the other six.

### 1.2 `TTMAccountCreated` omits creator and admin

`contracts/manager/TTMAccountManager.sol:142` emits only the account address,
so an indexer must make a follow-up `getTTMAccountCreator` call per account.
`ui/src/lib/receipt.ts:7` decodes this event and gets only the address.

**Change:** `TTMAccountCreated(address indexed account, address indexed creator, address indexed admin)`.

### 1.3 Silent state changes

Add events for two admin actions that currently emit nothing:

- `BookingToken.setManagerAddress` (`:684-687`) — repoints the entire
  authorization oracle for the token.
- `BookingToken.setMinExpirationTimestampDiff` (`:702-707`) — changes a
  mint-time validation rule.

### Migration note

These are UUPS proxies, so event signatures **can** change in an upgrade —
only storage layout is irreversible. The cost of deferring is a log-format
discontinuity: the activity feed will index the old format from launch, and
after this change the indexer must understand both formats indefinitely.

Doing 1.1 and 1.2 together in one upgrade keeps that to a single boundary.

---

## 2. Identity source of truth

**Impact:** storage layout · **Effort:** S · **Deferring has a real cost**

`TTMAccountManager._createTTMAccount` records account identity **twice**:

- `_ttmAccountInfo[addr] = {isTTMAccount: true, creator: msg.sender}` (`:273`)
- `grantRole(TTMACCOUNT_ROLE, addr)` (`:276`)

They are consumed by different systems:

- `BookingToken.isTTMAccount` → `ITTMAccountManager.isTTMAccount` → the **mapping**
- `ui/src/hooks/useMyAccounts.ts:21,34,42` → `getRoleMembers(TTMACCOUNT_ROLE)` → the **role**

`TTMACCOUNT_ROLE`'s admin defaults to `DEFAULT_ADMIN_ROLE`, so the manager
admin can `grantRole`/`revokeRole` at any time and the two views diverge — the
UI would show an account `BookingToken` rejects, or the reverse.
`ui/src/components/AccountValidityNotice.tsx` exists purely to paper over this.

**Not a security hole.** `BookingToken` authorizes off the mapping, which has
no external setter besides the factory path. It is a data-integrity problem.

**Change:** make the role authoritative — delete the `isTTMAccount` bool,
implement `isTTMAccount(a)` as `hasRole(TTMACCOUNT_ROLE, a)`, keep `creator` in
the mapping. Enumeration stays free via `AccessControlEnumerable`, so the UI
needs no change.

**Caveat that makes this more than a one-liner:** you must also prevent manual
granting, or the divergence simply moves. That means overriding
`grantRole`/`revokeRole` to revert for `TTMACCOUNT_ROLE`, or setting its role
admin to a role nobody holds.

**Cost of deferring:** the layout is frozen at deploy, so this becomes a
migration rather than an edit. Living with it costs one wasted slot per account
plus the divergence risk. Survivable, but this is the cheapest it will ever be.

---

## 3. Hash-native account API and a lens contract

**Impact:** breaking ABI · **Effort:** M · **Payoff:** large

`contracts/account/TTMAccount.sol:427-635` is ~200 lines of the same
three-line shape:

```solidity
function setServiceRestrictedRate(string memory serviceName, bool restrictedRate) public onlyRole(SERVICE_ADMIN_ROLE) {
    _setServiceRestrictedRate(getServiceHash(serviceName), restrictedRate);
    emit ServiceRestrictedRateUpdated(serviceName, restrictedRate);
}
```

Every call resolves a name to a hash via an **external staticcall to the
manager** (`:527-544`). Worse, `getSupportedServices()` (`:553-565`) makes
**one external call per service** — an account supporting 20 services costs 40
cross-contract staticcalls for a single view.

The string API is a presentation concern living inside the account contract.

**Change:** make the account's public surface hash-native
(`addService(bytes32 serviceHash, ...)`) and move name resolution into a
stateless `TTMLens` view contract that the UI and Go bot call. The account
stores and speaks `bytes32`; the lens does the joins and batch resolution.

**Benefits:** drops ~250 lines and roughly 2 KiB from `TTMAccount`, removes the
per-read fan-out, and removes the account's runtime dependency on the manager
for reads.

**Cost:** ABI-breaking for the Go bot, matrix-app-service, and UI
simultaneously. This is the item most worth batching with §1.

---

## 4. `recordExpiration` access-control posture

**Impact:** none · **Effort:** XS

`BookingToken.recordExpiration` (`:632-656`) is `public` with no restriction,
while `TTMAccount.recordExpiration` (`:374-376`) wraps it behind
`onlyRole(BOOKING_OPERATOR_ROLE)`. The role gate protects nothing — anyone can
call the underlying function directly.

This is fine: the function only marks a reservation expired once
`block.timestamp` has genuinely passed. It is objective housekeeping and
permissionless is correct.

**Change:** make the intent explicit — drop the misleading role gate on the
`TTMAccount` wrapper and add a NatSpec line stating the operation is
deliberately permissionless.

---

## 5. Service lifecycle events

**Impact:** additive · **Effort:** S

`ServiceRegistry._unregisterServiceName` (`:100-115`) removes the hash from the
registered set but **deliberately leaves both name mappings populated** so
existing accounts can still resolve deprecated names.

Effect: no _new_ account can add the service, existing accounts advertise it
indefinitely, and nothing tells the ecosystem it was deprecated. A bot querying
`getSupportedServices()` on a partner cannot distinguish a live service from a
deprecated one.

A full versioned-identity scheme (numeric service IDs, `supersededBy` chains)
was considered and **rejected** — partners are expected to retire old services
promptly as the protocol moves, so the maintenance burden outweighs the
benefit.

**Minimum useful change:** emit a `ServiceDeprecated(bytes32 indexed
serviceHash, string serviceName)` event on unregister, so consumers can learn
about lifecycle transitions from logs instead of diffing
`services/00_initial.json`.

---

## 6. Bot role separation

**Impact:** breaking for bot setup · **Effort:** S
**Related:** decision document, Decision 5

`TTMAccount.addMessengerBot` (`:699-712`) grants three roles at once:
`MESSENGER_BOT_ROLE`, `BOOKING_OPERATOR_ROLE`, `GAS_WITHDRAWER_ROLE`.

Once the operational policy is set (Decision 5), the engineering change is to
make gas withdrawal separately grantable so message-relaying bots cannot move
funds.

**Blocked on Decision 5** for the policy; the implementation itself is
straightforward.

---

## 7. Test coverage gaps

**Impact:** none · **Effort:** S–M

The suite is solid — 120 passing, with `test/BookingToken.test.js` covering the
cancellation state machine thoroughly, and `test/PartnerConfiguration.test.js`
properly _rewritten_ rather than gutted during the fee removal.

But the fee removal deleted ~284 lines of manager tests and **replaced none of
them**, and the behaviours that the removal made load-bearing were never
backfilled:

| #   | Missing test                                                     | Why it matters                                                                                                                         |
| --- | ---------------------------------------------------------------- | -------------------------------------------------------------------------------------------------------------------------------------- |
| 1   | `createTTMAccount` access control — in either direction          | The most consequential behavioural property of the manager on Base is untested. A test here would have caught the KYC-gate regression. |
| 2   | Role/mapping divergence (§2)                                     | No test grants `TTMACCOUNT_ROLE` directly and checks whether `BookingToken` accepts the address.                                       |
| 3   | `removeAllServices` (`TTMAccount.sol:464-471`)                   | The multi-service iteration path is unexercised.                                                                                       |
| 4   | `BookingTokenAddressUpdated` / `TTMAccountImplementationUpdated` | Emitted but never asserted.                                                                                                            |
| 5   | Cancellation refund to a contract that cannot receive ETH        | Would have caught the stuck-refund bug (Decision 3).                                                                                   |

Note on the general lesson: several of these test the _absence_ of a
restriction. Such tests do not catch bugs you already know about — they pin a
property so that a future refactor, or an environment change like the Camino →
Base migration, fails loudly instead of silently.

---

## 8. Minor code quality

**Impact:** none · **Effort:** S each

- **`supportsInterface` not overridden in `TTMAccount`** (`:383`). The contract
  implements `IERC721Receiver` (`:54`) but reports `false` for
  `type(IERC721Receiver).interfaceId`. Counterparties doing capability
  detection before transferring an NFT will conclude it cannot receive one.
- **Cancellation wrapper duplication.** `BookingToken.sol:721-850` is six
  near-identical wrappers, each doing `_requireOwned`, checking
  `BookingStatus.BOUGHT`, reading `supplier`, then delegating. One internal
  `_requireBoughtAndParties(tokenId) returns (address owner, address supplier)`
  collapses them to ~4 lines each, saving ~90 lines and roughly 1 KiB. ABI-identical.
- **`Service` struct cannot pack** (`PartnerConfiguration.sol:26-33`) — a
  `bool` sits in a full slot beside a `string[]`. Only fixable if capabilities
  become `bytes32[]`, which would also turn the O(n) `keccak256` comparison in
  `_removeServiceCapability` (`:201`) into a word compare. Layout + ABI change.
- **Unbounded reads.** `getRoleMembers(TTMACCOUNT_ROLE)` (used by
  `ui/src/hooks/useMyAccounts.ts:42`), `TTMAccount.getSupportedServices()`, and
  `ServiceRegistry.getAllRegisteredServiceNames()` all return unbounded arrays
  from a single `eth_call`. Not DoS vectors — no state-changing function
  iterates them — but they will time out on a public RPC at scale. Add
  paginated variants; additive and non-breaking.
- **`nonReentrant` on `finalizeCancellation`** for consistency with
  `buyReservedToken`.

---

## Suggested sequencing

1. **Spec 2 — consumer-breaking batch:** §1 (events), §3 (hash-native + lens),
   §2 (identity). Batch these so the bot, app-service, and UI absorb one
   coordinated change instead of three.
2. **Spec 3 — independent cleanup:** §4, §5, §7, §8. No consumer coordination
   required; can land piecemeal at any time.
3. **§6** follows Decision 5.
