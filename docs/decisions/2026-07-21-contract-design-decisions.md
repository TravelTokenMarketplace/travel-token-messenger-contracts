# Contract Design Decisions — For Discussion

**Date:** 2026-07-21
**Audience:** Engineering + product/management
**Status:** Open — no decision taken

## Why this document exists

A full review of the contracts was completed on 2026-07-21 ahead of a fresh
Base Sepolia deployment. Most findings were fixed directly (see
`docs/superpowers/specs/2026-07-21-predeploy-hardening-design.md`).

The six items below were **not** fixed, because each one changes how the
business operates — who may join the network, how partner funds move, what
partners must configure before they can trade. They need a decision from
outside engineering.

Engineering-only deferred work is tracked separately in
`docs/decisions/2026-07-21-technical-backlog.md`.

## Context you need

- The ecosystem moved from the **Camino** chain to **Base**. Camino enforced
  KYC/KYB at the chain level; Base does not. Some of our security assumptions
  were inherited from Camino and silently stopped holding.
- A **"fee removal"** refactor deleted message fees, cheques, and the account
  prefund. The prefund also served a second purpose nobody wrote down: it made
  spam expensive.
- The upcoming deployment is **Base Sepolia (testnet)**. Nothing is in
  production and no real value is at risk yet. Decisions here are cheap now and
  expensive after Base mainnet.

---

## Decision 1 — Who may create an account?

**Priority: highest.** Blocks Base mainnet.

### Situation

`createTTMAccount` has no access control. Anyone can create a Travel Token
Messenger account for the price of gas. Because `BookingToken` trusts any
address that holds an account, that person can then mint booking tokens,
reserve them against real partners, and appear in the ecosystem activity feed.

On Camino this was safe: the chain itself refused to let unverified addresses
deploy contracts, and the prefund made bulk creation costly. Both protections
are gone. The code did not change — the ground underneath it did.

Testnet impact is spam and noise. Mainnet impact is a permanently open door
into the partner network.

### Options

**A. Gate behind an `ACCOUNT_CREATOR_ROLE`** _(engineering recommendation)_

We hold a role; only role-holders may call `createTTMAccount`.

Common misconception worth clearing up: **this does not mean we hold partner
keys or operate their accounts.** The function already takes the partner's
admin address as a parameter. We call it _with their address_, and they control
the account from the first block. We gate who may _trigger_ creation, not who
_owns_ the result.

- Cost: one modifier. Onboarding is already a manual business process, so no
  operational change.
- Trade-off: partners cannot self-serve. If self-service onboarding is a
  product goal, this blocks it.

**B. Refundable deposit**

Creation stays open, but requires a deposit, refundable under conditions we
define. This is the closest restoration of the Camino economics.

- Cost: escrow accounting, a claim path, and a policy decision on who may
  reclaim and when.
- Trade-off: preserves permissionless onboarding. Sets a price on entry, which
  may deter legitimate small partners as well as spammers.

**C. Open creation, gated minting**

Anyone may create an account, but `BookingToken` checks a separate "verified"
flag before allowing a mint.

- Cost: new admin surface, plus a policy for who flips the flag and when.
- Trade-off: decouples "has an account" from "can affect the ecosystem."
  Accounts become free to create but inert until approved — arguably the
  cleanest model, and the most work.

### Recommendation

**Option A** for now: cheapest, matches how onboarding already works, and easy
to relax later. If self-service onboarding is on the roadmap, say so in this
meeting — that pushes toward **C**, and it is much cheaper to build now than to
retrofit.

### If we defer

Testnet is fine. Mainnet is not. Retrofitting after partners are live means an
upgrade plus a migration decision about already-created accounts.

---

## Decision 2 — Can a booking token move while a cancellation is pending?

### Situation

Today the behaviour is inconsistent, and probably not what anyone intended.

When a token with a pending cancellation proposal is transferred, the contract
**auto-rejects the proposal** and lets the transfer through, emitting rejection
reason `99`. But that only works when the transfer is initiated by the owner or
the supplier directly. If the owner has approved a marketplace or custody
contract and _it_ initiates the transfer, the transaction **reverts**.

So: same intent, two different outcomes depending on who submits the
transaction. The inconsistency is accidental.

**The auto-rejection itself is deliberate, however.** The protocol defines
`REJECTION_REASON_TRANSFER_ON_CHAIN = 99` — "automatic rejection reason during
transfer of the token on-chain" — so "transferring the token cancels a pending
proposal" is designed behaviour with a reserved protocol code, not an
improvisation. That materially raises the bar for Option A below.

### Options

**A. Block transfers while a proposal is pending** _(engineering recommendation)_

An explicit revert. Consistent for owners, suppliers, and approved operators
alike.

- Simple, predictable, easy to explain to partners.
- Trade-off: a pending proposal temporarily freezes the asset. If a
  counterparty stops responding, the holder is stuck until it resolves — so
  this option needs a proposal timeout or an abandonment path.
- **Trade-off: this abandons a designed protocol behaviour.** Rejection reason
  99 exists specifically to signal auto-rejection on transfer. Choosing A means
  retiring that code path and the enum value with it, which is a protocol-level
  change, not just a contract one.

**B. Allow transfer, resolve the proposal, fix the operator case** _(engineering recommendation)_

Keep the designed behaviour and fix only the defect: authorize against the true
owner (`ownerOf` / `getApproved` / `isApprovedForAll`) rather than the
transaction sender, so marketplace and custody transfers stop failing.

- Preserves composability with marketplaces and custody providers.
- Keeps reason 99 working as the protocol intends.
- Smaller change: one authorization check, no protocol coordination.
- Trade-off: keeps a rule that surprises people — transferring an asset cancels
  a negotiation in progress. That surprise is at least _documented_ in the
  protocol enum.

### Recommendation

**Option B.** An earlier draft of this document recommended A on the mistaken
belief that reason 99 was an undefined placeholder. It is not — it is a
reserved protocol value with a clear meaning, so the auto-rejection is
intentional design and option B simply repairs the operator-authorization bug
underneath it.

The question for this meeting is still worth asking, because it is the thing
that would change the answer: _do we expect booking tokens to be traded on
marketplaces or held in custody contracts?_ If yes, B is clearly right. If they
only ever move between known partner accounts, A becomes defensible as a
simplification — but it costs a protocol change, so it should be a deliberate
product decision rather than a cleanup.

**If A: we need a timeout policy.** How long may a proposal pend before it can
be unilaterally withdrawn? That is a business answer, not a technical one.

---

## Decision 3 — How should cancellation refunds be paid?

### Situation

When a cancellation is finalized, the contract **pushes** the refund to
whoever currently holds the booking token, in the same transaction that marks
the booking cancelled.

If that push fails, the entire cancellation fails. Permanently. There is no
retry that helps and no administrative override.

The push fails when the holder is a contract that cannot receive ETH — a
custody wrapper, a vault, a multisig that did not implement the right function
— or when an ERC-20 refund hits a blacklisting stablecoin. The holder does not
need to be malicious; ordinary composition is enough.

Result: the booking can never be cancelled, and the refund stays locked in the
partner's account.

**Correction (2026-07-22, found during pre-deploy test backfill):** the
diagnosis above is wrong in a way that matters, and none of the options below
fix the actual problem.

`ownerAccepted` is only ever set at
`contracts/booking-token/BookingTokenCancellable.sol:226,299,352`, all gated
on `msg.sender == owner`, and every public entry point in
`contracts/booking-token/BookingToken.sol` is `onlyTTMAccount(msg.sender)`.
That means the real blocker is: **the current owner must be a registered TTM
Account.** `finalizeCancellation` reverts with `OwnerNotAcceptedCancellation`
one layer _before_ the refund push described above is ever reached — the push
is never attempted, so pull payments (either option below) would not change
this outcome at all. `finalizeCancellation` would revert exactly as it does
today.

The reachable condition is not "the holder is a contract that cannot receive
ETH." It is simply: **the booking token is transferred to a normal wallet.**
An ordinary, willing EOA holder triggers this — not just ETH-rejecting
contracts — because an EOA is never a registered TTM Account. This is a
supported flow, not an exotic edge case: `checkTransferable`
(`BookingToken.sol:646-653`) explicitly permits transferring a `BOUGHT` token
to any address, and this is a standard ERC-721 with no transfer allowlist.

No ETH is ever stranded — the whole call reverts atomically, so nothing about
this is a funds-safety issue. The defect is liveness: the booking becomes
permanently un-cancellable, with no retry, no admin override, no timeout, and
no expiry path (`recordExpiration` explicitly reverts on `BOUGHT`). The only
remedy today is a contract upgrade.

Therefore: whichever option is chosen from the list below, it will not
deliver the fix it promises. This decision needs to be re-framed around the
actual blocker (the TTM-Account-holder gate) before it is decided. The
behaviour described here is now pinned by a test in
`test/BookingToken.test.js`.

### Options

**A. Pull payment for refunds only** _(engineering recommendation)_

Mark the booking cancelled, credit the refund to the recipient's balance, and
let them withdraw when they choose.

- Fixes the stuck-funds path with the smallest change. Mint and buy payments
  are unaffected, because those recipients are always partner accounts we know
  can receive value.
- Trade-off: recipients gain an extra step. Needs a policy answer: **if the
  token changes hands between cancellation and claim, who is owed the refund —
  the holder at cancellation time, or the current holder?** Engineering
  recommends crediting at cancellation time; it is simpler and matches who
  actually agreed to the cancellation.

**B. Pull payment everywhere**

Every outbound payment goes through a withdrawable balance.

- Uniform and maximally safe.
- Trade-off: changes the normal, non-cancellation flow for every supplier —
  they would have to claim their sale proceeds. Worse day-to-day experience to
  fix a rare failure.

### Recommendation

**Option A.** B fixes a rare problem by degrading the common path.

### If we defer

The bug is real but narrow — it needs a counterparty holding the token in a
contract that cannot receive value. Low probability today, rising as soon as
anyone integrates custody or a marketplace.

---

## Decision 4 — Should the payment-token allowlist be enforced?

### Situation

Partners declare which tokens they accept, via `addSupportedToken`. Nothing
enforces it. The list is read by nobody in the payment path — a booking can be
priced in **any** ERC-20, including one the partner never approved, or a
worthless token crafted to look like a real one.

Partners reasonably assume declared configuration is enforced. It is
decoration.

### Options

**A. Enforce on-chain** _(engineering recommendation)_

Reject mints priced in a token the supplier has not declared.

- Makes declared configuration a real guarantee.
- Trade-off: **partners must configure payment tokens before they can trade.**
  A partner who has not added any token silently cannot mint. This is an
  onboarding-flow change and needs a good error surfaced in the UI.

**B. Document as advisory**

State clearly in the interface docs that the list is discovery metadata and
validation is the bot's responsibility.

- Zero contract risk, no onboarding change.
- Trade-off: keeps a trap for integrators who assume it is enforced.

### Recommendation

**Option A**, paired with an onboarding checklist item and a clear UI error.
The security value is modest; the value of configuration meaning what it says
is high.

The question for this meeting: **does requiring token configuration before
first trade break the intended onboarding experience?**

---

## Decision 5 — Bot key policy

### Situation

Registering a messenger bot grants one address three capabilities at once:
send protocol messages, operate bookings, and withdraw gas money (default:
10 ETH per 24 hours, **per bot** — adding a second bot adds a second
allowance).

Messenger bots run as hosted services with hot keys. That is the realistic
compromise scenario, and one leaked key currently yields all three
capabilities.

### Options

**A. Split the funds capability from the messaging capability**

Make gas withdrawal separately grantable so bots that only relay messages
cannot move funds.

**B. Reduce the default allowance and make it per-bot configurable**

Cap exposure without changing the role model.

**C. Accept and document**

Treat it as inherent to hot-key automation; rely on monitoring and the existing
bot-removal path.

### Recommendation

**B now, A when convenient.** B is a parameter change with immediate benefit.
The real question for this meeting is operational: **what is the maximum we are
willing to lose to a single compromised bot key in 24 hours?** The contract
default should follow that number, not the other way round.

Also needs an owner: who monitors for anomalous withdrawals, and who is on call
to remove a compromised bot?

---

## Decision 6 — Admin key custody

### Situation

The deployment currently puts every privileged role on a single deployer EOA:
manager admin, pauser, upgrader, implementation versioner, plus booking-token
admin and upgrader.

One key compromise means: pause the network, upgrade any contract to arbitrary
code, and change the implementation every future partner account deploys
against.

There is also no recovery path. The manager is a singleton — an accidental
admin transfer to a wrong address permanently disables account creation,
service registration, and versioning.

Note: OpenZeppelin's built-in two-step admin protection enforces exactly **one**
admin address, which is incompatible with the current multi-admin model. A
multisig resolves this properly — several people control it, but on-chain it is
a single address.

### Options

**A. Safe multisig for privileged roles** _(engineering recommendation)_

`DEFAULT_ADMIN`, `UPGRADER`, `VERSIONER` move to a Safe. `PAUSER` stays on a hot
operations key so incident response does not need a signing threshold.

- No contract changes. Purely deployment wiring.
- Needs decisions: **who are the signers, and what threshold?**

**B. Add a timelock on upgrades**

Upgrades queue publicly for N days before execution, so partners can observe a
pending implementation change.

- Complements the current design, where partners review a verified
  implementation and choose when to adopt it.
- Trade-off: emergency fixes are slower.

**C. Keep EOAs, manage by process**

No change; rely on discipline.

### Recommendation

**A for Base Sepolia** — and rehearse it on testnet, because mainnet is where
getting it wrong is unrecoverable. Revisit **B** before mainnet.

Concrete questions for this meeting:

1. Who are the Safe signers, and what is the threshold?
2. Who holds the pauser key, and what is the escalation path?
3. Do partners get a timelock guarantee on upgrades before mainnet?

---

## Summary

| #   | Decision                              | Blocks mainnet? | Engineering recommendation                      |
| --- | ------------------------------------- | --------------- | ----------------------------------------------- |
| 1   | Who may create an account             | **Yes**         | `ACCOUNT_CREATOR_ROLE`                          |
| 2   | Transfers during pending cancellation | No              | Keep auto-rejection; fix operator authorization |
| 3   | Refund payment model                  | No              | Pull payment, refund leg only                   |
| 4   | Enforce payment-token allowlist       | No              | Enforce, with onboarding change                 |
| 5   | Bot key policy                        | No              | Lower default, split roles later                |
| 6   | Admin key custody                     | **Yes**         | Safe multisig; timelock before mainnet          |

Decisions 1 and 6 should be settled before Base mainnet. The rest can follow
testnet experience — and testnet is a good place to learn which of them
actually matter in practice.
