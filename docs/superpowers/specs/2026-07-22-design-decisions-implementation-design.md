# Design Decisions Implementation — Design

**Date:** 2026-07-22
**Status:** Approved
**Supersedes discussion in:** `docs/decisions/2026-07-21-contract-design-decisions.md`

## Why this document exists

`docs/decisions/2026-07-21-contract-design-decisions.md` raised six questions
that needed answers from outside engineering. Management answered all six. This
spec turns those answers into an implementation.

Nothing is deployed yet — `ignition/deployments/` does not exist and
`ui/src/contracts/generated/addresses.ts` exports an empty `ADDRESSES` map.
Storage layout and ABI are therefore free to change. That window closes at the
Base Sepolia deployment, which happens after this work lands.

## Decisions as taken

| #   | Decision                              | Outcome                                                              |
| --- | ------------------------------------- | -------------------------------------------------------------------- |
| 1   | Who may create an account             | **No change.** Creation stays open and free.                          |
| 2   | Transfers during pending cancellation | **Option B.** Allow transfer, resolve the proposal, fix operator auth. |
| 3   | Cancellation when owner is not a TTM Account | **Gate uniformly.** Reject on all six entry points.            |
| 4   | Payment-token allowlist               | **Enforce on-chain**, as one uniform allowlist.                        |
| 5   | Bot key policy                        | **Drop `GAS_WITHDRAWER_ROLE` from bot defaults**; allowance 10 → 0.01 ETH. |
| 6   | Admin key custody                     | **Safe multisig**, for our roles and available to partners.            |

Decision 1 requires no work and is not discussed further.

---

## Decision 2 — Fix operator authorization on transfer

### What is wrong

`checkTransferable` (`contracts/booking-token/BookingToken.sol:607-643`)
auto-closes a `PENDING` cancellation proposal before letting a transfer
through, using the reserved protocol code
`REJECTION_REASON_TRANSFER_ON_CHAIN = 99`. That is designed behaviour, not an
improvisation.

It only works when the owner or supplier submits the transaction directly. Two
places assume `msg.sender` is a party to the booking:

1. `_rejectCancellation` and `_withdrawCancellation` both carry
   `onlyOwnerOrSupplier(owner, supplier)`
   (`contracts/booking-token/BookingTokenCancellable.sol:418,388`), which
   compares against `msg.sender`.
2. The branch in `checkTransferable` that chooses between reject and withdraw
   tests `msg.sender != currentProposer`.

When an approved marketplace or custody contract initiates the transfer,
`msg.sender` is the operator, both assumptions fail, and the transfer reverts
instead of closing the proposal.

### The fix

Authorize against the true owner rather than the transaction sender, using
`ownerOf` / `getApproved` / `isApprovedForAll` on the transfer path.

The reject-versus-withdraw choice keys off the **owner**, not `msg.sender`. A
transfer is initiated by the owner or by someone the owner approved, so:

- owner **is** the current proposer → `_withdrawCancellation` (they are
  abandoning their own proposal)
- owner **is not** the current proposer → `_rejectCancellation` (the
  counterparty's proposal is being rejected)

This produces identical behaviour to today for direct owner-submitted
transfers, and fixes the operator case rather than changing the semantics.

Scope the change to the transfer path. The six external cancellation entry
points keep authorizing against `msg.sender`, which is correct for them — they
are direct calls by a party, not delegated transfers.

### Consequence accepted

A booking token transferred to a non-TTM-Account address cannot be cancelled.
Management accepted this explicitly. Decision 3 makes it fail loudly rather
than silently.

---

## Decision 3 — Cancellation requires a TTM Account owner

### What is wrong

All six cancellation entry points
(`contracts/booking-token/BookingToken.sol:802-858`) are
`onlyTTMAccount(msg.sender)` and route through `_requireBoughtAndParties` →
`requireOwnerOrSupplier`. Separately, `ownerAccepted` is only ever set true
when `msg.sender == owner`
(`contracts/booking-token/BookingTokenCancellable.sol:226,299,352`).

When the owner is an ordinary wallet, the **supplier** can still call
`initiateCancellation` — it satisfies both gates. That writes a `PENDING`
proposal with `ownerAccepted = false` which can never become true, because the
wallet owner can never call `acceptCancellation`. `finalizeCancellation` then
reverts with `OwnerNotAcceptedCancellation` forever.

Nothing is stuck financially; the whole call reverts atomically. What is broken
is that a supplier can open a negotiation that is dead on arrival, with no
signal until finalize.

### The fix

One check in `_requireBoughtAndParties`
(`contracts/booking-token/BookingToken.sol:787`), which is the shared helper
for all six entry points:

```solidity
requireTTMAccount(owner);
```

Applied uniformly. No escape-hatch exemption for `withdrawCancellation` or
`rejectCancellation`, for reasons recorded below.

### Why uniform, and why the transfer path is unaffected

An earlier draft proposed exempting `withdraw` and `reject` so a supplier could
always close a dead proposal. Two findings retired that idea:

- **`reject` could not have helped.** The proposer cannot reject
  (`ProposerCanNotRejectCancellation`,
  `contracts/booking-token/BookingTokenCancellable.sol:432`). With a wallet
  owner the proposer is necessarily the supplier, so the only party who could
  reject is the owner — who cannot call anything. The exemption unlocks nobody.
- **The state is unreachable.** A proposal opened while the owner was a TTM
  Account is closed by `checkTransferable` during any subsequent transfer; a
  proposal against an already-wallet-held token is what this gate blocks. With
  Decision 2 fixed, there is no third path in.

`_requireBoughtAndParties` is used **only** by the six external entry points.
`checkTransferable` computes `owner` and `supplier` itself and calls
`_rejectCancellation` / `_withdrawCancellation` directly, bypassing the helper.
The new gate therefore cannot block the transfer-time auto-close. This is the
one way the change could have been dangerous, and it structurally is not.

### Required test

The unreachability argument above is load-bearing, so it gets pinned rather
than asserted:

- **Transfer always closes a pending proposal** — including when an approved
  operator initiates the transfer, which is the case Decision 2 repairs. This
  test is what makes the uniform gate safe.

### Partner-facing rule

> A booking token is cancellable only while a TTM Account holds it.
> Transferring it out cancels any negotiation in progress, and cancellability
> returns if the token comes back to a TTM Account.

---

## Decision 4 — Enforce the payment allowlist, as one uniform list

### Why, restated

The original document justified this as protection against "a worthless token
crafted to look like a real one." That attacker is not in the picture:
`safeMintWithReservation` is `onlyTTMAccount(msg.sender)` and **the supplier is
`msg.sender`**, setting its own `price` and `paymentToken`. Nobody else can
price a supplier's booking.

The real beneficiary is different and better. Minting is performed by a bot
holding `BOOKING_OPERATOR_ROLE` — a hot key on a hosted service. Enforcing the
allowlist turns partner configuration into a **bound on what a compromised or
misconfigured bot can do**, set by `SERVICE_ADMIN_ROLE`. That is the same theme
as Decision 5, and it is the justification that survives scrutiny.

### One allowlist, not three mechanisms

`BookingToken` already encodes payment mode as an address
(`contracts/booking-token/BookingToken.sol:103,111`):

- `NATIVE_PAYMENT = address(0)`
- `OFFCHAIN_PAYMENT = address(1)`
- anything else — an ERC-20 address

Configuration will use the same encoding. `_supportedTokens` becomes the single
declaration of what a partner accepts, with `address(0)` and `address(1)` as
legitimate members.

Consequently **`_supportsOffChainPayment` is deleted**, along with
`setOffChainPaymentSupported` and `offChainPaymentSupported`
(`contracts/partner/PartnerConfiguration.sol:36,394,403`;
`contracts/account/TTMAccount.sol:672`). Keeping it would mean two encodings
for one concept, and enforcement code that branches three ways to translate
between them.

### Plumbing this requires

The one-line check below does not work against the code as it stands. Three
pieces are missing:

1. **`isSupportedToken(address) → bool` does not exist.** Only
   `getSupportedTokens()` does, returning the whole array
   (`contracts/partner/PartnerConfiguration.sol:384`). Add a membership check
   wrapping `EnumerableSet.contains`. Looping the array from `BookingToken`
   instead would make mint gas scale with the size of a partner's allowlist,
   which is unacceptable.
2. **`ITTMAccount` is a stub.** It declares `initialize` and nothing else
   (`contracts/account/ITTMAccount.sol:6`). `isSupportedToken` must be declared
   on it.
3. **`BookingToken` cannot reach the account.** It imports `ITTMAccountManager`
   but not `ITTMAccount`. Add the import.

With those in place, enforcement at mint is one line — `msg.sender` is the
supplier's account:

```solidity
if (!ITTMAccount(msg.sender).isSupportedToken(address(paymentToken))) {
    revert PaymentTokenNotSupported(paymentToken);
}
```

### Enforce at mint only

A reservation is a standing offer. Removing a token from configuration must not
retroactively break an outstanding reservation a distributor is about to
accept, so `buyReservedToken` does not re-check. The token was valid when the
supplier offered it.

### Discoverability is an interface concern

Partners never see `address(0)` or `address(1)`. Both
`ui/src/pages/tabs/PaymentTokensTab.tsx` and `tasks/account.js` present "Native
ETH" and "Off-chain payment" as named options that map to the sentinels. The
sentinel values are documented on the interface for integrators.

### Risks carried into implementation

- **`EnumerableSet.AddressSet` must accept `address(0)`.** It tracks membership
  by stored position rather than by the value being non-zero, so this should
  hold — but it is load-bearing enough to get an explicit test rather than an
  assurance.
- **`address(0)` in `getSupportedTokens()` can be misread as "unset"** by a
  naive integrator. Mitigated by documenting the sentinels on the interface and
  labelling them in UI and CLI.

### Blast radius

`PartnerConfiguration` (add membership check, delete the off-chain bool),
`ITTMAccount` (declare it), `BookingToken` (import and enforce), `TTMAccount`
(delete the off-chain setter), tests,
`ui/src/pages/tabs/PaymentTokensTab.tsx`, `tasks/account.js`.

The bot reads neither surface — its `offChainPayment*` references are all
`offChainPaymentCurrency`, the mint parameter, not the configuration flag.

---

## Decision 5 — Bot key policy

### Changes

1. **Drop `_grantRole(GAS_WITHDRAWER_ROLE, bot)` from `addMessengerBot`**
   (`contracts/account/TTMAccount.sol:732`). Registering a bot grants
   `MESSENGER_BOT_ROLE` and `BOOKING_OPERATOR_ROLE` only.
2. **Keep the revoke in `removeMessengerBot`**
   (`contracts/account/TTMAccount.sol:746`) — revoking an unheld role is a
   no-op, and keeping it means removal still fully de-authorizes a bot that was
   granted the role later.
3. **Default allowance `10 ether` → `0.01 ether`**
   (`contracts/account/TTMAccount.sol:257`), period unchanged at 24 hours.

### No new function is needed

`TTMAccount` contains no `_setRoleAdmin` calls, so every role admins under
`DEFAULT_ADMIN_ROLE`. A partner grants gas access through the inherited public
`grantRole(GAS_WITHDRAWER_ROLE, bot)`.

### The separation this creates

A delegated `BOT_ADMIN_ROLE` holder can onboard and remove bots but **cannot**
give them access to funds — that requires `DEFAULT_ADMIN_ROLE`. This is
intended: fund access becomes a deliberate act at a higher privilege level than
routine bot operations.

### Note on initial funding

`addMessengerBot(bot, gasMoney)` still sends `gasMoney` to the bot immediately
via `sendValue`. That path is independent of `GAS_WITHDRAWER_ROLE` and is
unchanged — it remains gated on `BOT_ADMIN_ROLE`.

### Confirmed no impact on our bot

`grep -rli gasmoney` over `travel-token-messenger-bot` returns nothing. The bot
has never withdrawn gas money.

---

## Decision 6 — Safe multisig

### Two separate scopes

The original document's Decision 6 concerned **our** protocol roles. Management
also wants Safe support for **partner** accounts. These are different problems.

### Partner accounts: creation already works, operating does not

`ui/src/pages/CreateAccount.tsx:18-45` takes the admin as a free-text field
defaulting to the connected address, and passes it to `createTTMAccount`. A
partner can wire a Safe as admin today by pasting its address.

But `ui/src/wallet/wagmi.ts:29` configures `connectors: [injected()]` and
nothing else. RainbowKit is in `package.json` but its connector set is not
wired up. The moment a partner's admin is a Safe, they can no longer administer
their own account through the UI — add bots, register services, set payment
tokens, none of it — because an injected EOA does not hold the role.

**Scope for now: add WalletConnect and wagmi's `safe()` connector.** This opens
two paths at once. Over WalletConnect the partner connects their Safe as a
wallet and our transactions arrive as proposals for signers to approve. With
the `safe()` connector the UI becomes loadable inside the Safe interface as a
Safe App.

This UI is developer-facing, so Safe support is not strictly required here. It
is being done now anyway: it is small, it lets the flow be exercised on testnet
before it matters, and it builds familiarity with the Safe libraries ahead of
partner-facing work.

**Deferred:** guided Safe creation inside `CreateAccount.tsx` — collecting
owners and threshold, deploying through Safe's canonical `SafeProxyFactory`,
pre-filling the result as the admin. Recorded for later, not built now. Safe's
canonical deployment addresses on Base Sepolia are unverified and must be
confirmed before that work starts.

**Rejected:** an on-chain `createTTMAccountWithSafe(...)` on the manager. It
buys one transaction at the cost of a permanent dependency from our core
contract on a third party's per-chain deployment addresses.

**Not enforced:** no contract-level check that an admin is a multisig.
`code.length > 0` proves only "contract", not "Safe", is wrong for EIP-7702
wallets, and goes stale when the admin changes.

### Our own roles

`DEFAULT_ADMIN`, `UPGRADER` and `VERSIONER` move to a Safe. `PAUSER` stays on a
hot operations key so incident response does not need a signing threshold.

This is deployment wiring — no contract change. Signers, threshold, and the
pauser escalation path are settled as part of phase 3.

---

## Phasing

Three phases, in order. Phase 2 depends on phase 1's ABI regeneration; phase 3
is independent of both.

### Phase 1 — Contracts

Decisions 2, 3, 4, 5 plus tests. Then `yarn compile` and
`yarn hardhat export-abi` so `abi/` — and therefore the UI — stays in sync.

Test coverage this phase must add:

- Transfer always closes a pending proposal, including operator-initiated
  transfers (Decision 3's safety argument depends on this).
- Operator-initiated transfer succeeds rather than reverting (Decision 2).
- Cancellation entry points revert when the owner is not a TTM Account
  (Decision 3).
- Mint rejects an undeclared payment token; accepts a declared ERC-20,
  `address(0)`, and `address(1)` (Decision 4).
- Mint reverts for a supplier that has declared nothing at all (Decision 4 —
  this is the onboarding-order change the bot must handle).
- `isSupportedToken` agrees with `getSupportedTokens` across add and remove,
  including for `address(0)` and `address(1)` (Decision 4).
- `addMessengerBot` does not grant `GAS_WITHDRAWER_ROLE`; `DEFAULT_ADMIN` can
  grant it afterwards; `removeMessengerBot` still revokes it (Decision 5).
- Default allowance is 0.01 ETH per 24 hours (Decision 5).

### Phase 2 — UI

- WalletConnect and `safe()` connectors in `ui/src/wallet/wagmi.ts` (Decision 6).
- Grant/revoke `GAS_WITHDRAWER_ROLE` affordance (Decision 5).
- `PaymentTokensTab` presents native and off-chain as named options over the
  sentinel addresses, and drops the off-chain boolean control (Decision 4).
- `tasks/account.js` gets the same treatment for CLI onboarding.

### Phase 3 — Deployment wiring

Safe for our privileged roles, signers and threshold agreed, pauser on a
separate hot key. Base Sepolia deployment follows once all three phases are
complete. There is no schedule pressure on the deployment.

## Cross-repo follow-up

`BOT-MIGRATION.md` in the workspace parent folder must be updated once phase 1
lands. Three items:

- **Decision 4 changes onboarding order.** Supplier mints revert until payment
  configuration exists. A partner who has declared nothing cannot trade.
- **Decision 3 adds a new revert.** Cancellation calls against a token held by
  a non-TTM-Account address now fail immediately. The bot should surface this
  rather than retry.
- **Decision 5 is confirmed no-impact.** Recorded so the question is not
  reopened.
