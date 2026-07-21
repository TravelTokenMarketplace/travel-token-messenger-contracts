# Pre-Deployment Hardening — Design

**Date:** 2026-07-21
**Status:** Approved, pending implementation plan
**Target:** Fresh Base Sepolia (84532) deployment
**Source review:** `docs/superpowers/specs/2026-07-21-pre-deployment-review.md`

## Purpose

Prepare the contracts for a from-scratch Base Sepolia deployment by fixing
everything that can be fixed **without breaking consumers or requiring
decisions from outside the engineering team**.

Deferred work is tracked in two companion documents and is explicitly *not*
part of this spec:

- `docs/decisions/2026-07-21-contract-design-decisions.md` — items needing
  business or management sign-off.
- `docs/decisions/2026-07-21-technical-backlog.md` — engineering-only items
  deferred to a later spec.

## Scope boundary

**In scope:** changes that are internal, additive, or invisible to the bot,
matrix-app-service, and UI.

**Out of scope:** ABI-breaking signature changes, changes to business flows
(who may create an account, how refunds settle, what tokens are accepted),
and anything requiring cross-repo coordination beyond a Go-bindings version
bump.

**Accepted consumer impact:** `abi/`, `docs/`, and `go/contracts/` are
regenerated. The bot and matrix-app-service pin the contracts Go module and
will need a version bump to pick up the new bindings. No consumer *code*
changes are required — no symbol they use is removed or altered. This was
confirmed acceptable: the bot has not cut a release yet.

## Goals

1. The deployed bytecode is optimized for the way these contracts are actually
   used (called constantly, deployed once).
2. No storage slot is wasted or reserved for a concept that no longer exists.
3. `yarn hardhat ignition deploy` followed by a documented runbook produces a
   *working* system, not an inert one.
4. Contract verification works on Basescan.
5. No dead code, and no comment that describes behaviour the code does not have.

## Non-goals

- Closing the permissionless-account-creation gap (see Accepted Risks).
- Any change to event signatures.
- Any change to the service-name identity scheme.

---

## 1. Compiler settings

`hardhat.config.js`

| Setting | From | To | Rationale |
|---|---|---|---|
| `optimizer.runs` | `1` | `1000` | `runs: 1` optimizes *deployment* cost. These are upgradeable implementations deployed once and called constantly by messenger bots, so runtime cost is what matters. The original justification — fighting the 24 KiB limit while fee support existed — no longer applies. |
| `evmVersion` | `paris` | `cancun` | Base Sepolia has supported Cancun since Dencun. Paris forgoes transient storage and `MCOPY` for no benefit. |

**Verification:** re-run `yarn compile` and record `hardhat-contract-sizer`
output. Measured headroom at `runs: 1000` was 3.3 KiB before this spec's
additions; `Pausable` and the new role will consume some of it. **If any
contract exceeds 22.5 KiB (2 KiB headroom under the 24.576 KiB EIP-170 limit),
drop to `runs: 500` and re-measure.** Full test suite must pass at the final
setting.

## 2. Storage layout

Irreversible after deployment. Zero consumer impact — these are internal
namespaced-storage details.

### 2.1 Remove the vestigial prefund slot

`contracts/account/TTMAccount.sol:120-123`

```solidity
// DELETE:
/**
 * @dev Prefund amount
 */
uint256 _unused; // Not used, but do not remove. Previously used to store the prefund amount.
```

`TTMAccountStorage` becomes `{ address _manager; address _bookingToken; }`.
The "do not remove" comment was correct while upgrading a live proxy; a
from-scratch deploy is the moment that constraint lifts. Verified never read
or written.

### 2.2 Pack `GasMoneyStorage`

`contracts/account/GasMoneyManager.sol:26-31`

```solidity
// FROM:
uint256 _withdrawalLimit;
uint256 _withdrawalPeriod;
mapping(address => uint256) _withdrawalPeriodStart;
mapping(address => uint256) _withdrawnAmount;

// TO:
uint128 _withdrawalLimit;    // wei; 3.4e38 >> total ETH supply
uint64  _withdrawalPeriod;   // seconds; 5.8e11 years
mapping(address => GasMoneyWithdrawal) _withdrawals;

struct GasMoneyWithdrawal {
    uint128 amount;
    uint64  periodStart;
}
```

Removes one cold SLOAD from every bot gas withdrawal — the highest-frequency
call in the system — and collapses the paired mappings from two cold slots to
one.

**Critical constraint: every external signature stays `uint256`.** The narrowed
types are an internal storage detail only. Three call sites would otherwise
leak the change into the ABI and break consumers:

| Location | Signature | Requirement |
|---|---|---|
| `GasMoneyManager.sol:141` | `getGasMoneyWithdrawal() returns (uint256, uint256)` | Widen on read |
| `GasMoneyManager.sol:157` | `getGasMoneyWithdrawalForAccount(address) returns (uint256, uint256)` | Widen on read |
| `GasMoneyManager.sol:129` | `setGasMoneyWithdrawal(uint256 limit, uint256 period)` | Narrow on write, checked |
| `GasMoneyManager.sol:72` | `__GasMoneyManager_init(uint256, uint256)` | Narrow on write, checked |

Rules:

- **Reads** widen implicitly (`uint128` → `uint256` is safe and free).
- **Writes** use a checked narrowing cast that **reverts on overflow rather
  than truncating**. A silently truncated withdrawal limit would be a security
  bug, not a display bug.
- The `WithdrawalLimitExceeded` / `WithdrawalLimitExceededForPeriod` errors keep
  `uint256` parameters.

Net effect: the ABI is byte-identical, and the gas saving is entirely internal.

## 3. Correctness fixes

All internal. No signature changes.

### 3.1 Missing reentrancy-guard initializer

`contracts/manager/TTMAccountManager.sol:189-203` — `initialize` does not call
`__ReentrancyGuard_init()` despite `createTTMAccount` using `nonReentrant`.
Functional under OZ v5 (uninitialized `_status == 0 != ENTERED`) but pays a
cold-slot SSTORE on first use. `TTMAccount.initialize:221` already does this
correctly; this is consistency.

### 3.2 Wire the unchained initializers

`__PartnerConfiguration_init()` (`PartnerConfiguration.sol:99`) and
`__ServiceRegistry_init()` (`ServiceRegistry.sol:58`) are declared but never
called. They are empty today, which is exactly the trap: the day someone adds
a default, it silently will not apply.

**Keep the OZ `_init` / `_init_unchained` convention** — it is correct and
deliberate. Call them from `TTMAccount.initialize` and
`TTMAccountManager.initialize` respectively, matching how
`__GasMoneyManager_init` is already wired.

### 3.3 Zero-address validation in initializers

`TTMAccountManager.initialize`, `TTMAccount.initialize`,
`BookingToken.initialize` — none validate their role parameters.
`AccessControl._grantRole(role, address(0))` silently succeeds and is
permanently unusable. Add checks mirroring the existing
`TTMAccountInvalidAdmin` pattern in `createTTMAccount`.

### 3.4 Make `BookingTokenCancellable` abstract

`contracts/booking-token/BookingTokenCancellable.sol:12` — currently a
concrete, deployable contract that emits a stray 0.665 KiB artifact and does
not inherit `Initializable` unlike every sibling mixin.

```solidity
abstract contract BookingTokenCancellable is Initializable {
```

## 4. Pausable on BookingToken

`BookingToken` holds the system's value flow but has no pause lever; today the
only response to a discovered bug in the payment path is a full upgrade.

Add `PausableUpgradeable` to the inheritance list, a `PAUSER_ROLE`,
`__Pausable_init()` in `initialize`, and role-gated `pause()` / `unpause()`.
Apply `whenNotPaused` to the three value-flow functions:

- `safeMintWithReservation` (`:351`)
- `buyReservedToken` (`:449`)
- `finalizeCancellation` (`:824`)

**Consumer impact: none while unpaused.** `whenNotPaused` changes no
signature, and behaviour is identical in the unpaused state. Only `pause()`,
`unpause()`, `paused()`, and the `Paused`/`Unpaused` events are new, all
additive. OZ v5 `PausableUpgradeable` uses its own ERC-7201 namespace, so
there is no collision with `BookingTokenStorage`.

`PAUSER_ROLE` is granted to a hot operations key at deploy (see §8), not to
the Safe — incident response should not require a signing threshold.

## 5. Additive behaviour changes

### 5.1 Emit `Deposit`

`contracts/account/TTMAccount.sol:239`

```solidity
// FROM: receive() external payable {}
receive() external payable {
    emit Deposit(msg.sender, msg.value);
}
```

`event Deposit(address indexed sender, uint256 amount)` is declared
(`TTMAccount.sol:148`) but never emitted, so every ETH transfer into an
account — including the `msg.value` that `TTMAccountManager` forwards at
creation — is invisible to the activity feed.

`ui/src/lib/activity/catalog.ts:174` already renders this event and reads
`a.sender` / `a.amount`, matching the declaration exactly, with a passing test
at `catalog.test.ts:42`. Emitting it makes existing UI code work with no UI
change.

### 5.2 Token symbol

`contracts/booking-token/BookingToken.sol:299`

```solidity
__ERC721_init("BookingToken", "BToken");   // was "TRIP"
```

`BToken` is the intended symbol; it was previously applied post-deploy by a
migration script rather than set at initialization.

**Also delete `reinitializeV2`** (`BookingToken.sol:324`). It exists solely to
rename the token, burns reinitializer version 2, and is a standing admin
backdoor to rewrite the token's identity. With the correct symbol set at
`initialize`, it has no purpose.

**Consumer updates required:**

- `test/BookingToken.test.js:45` and `ui/src/pages/tabs/BookingTokenTab.test.tsx:11,49`
  assert `"TRIP"` — update to `"BToken"`.
- `test/BookingToken.test.js:52,56,64` exercise `reinitializeV2` (access
  control, happy path, and the double-reinitialize revert). **Delete these
  three assertions** along with the function.

Removing `reinitializeV2` deletes a function from the ABI. This is the only
subtractive ABI change in this spec. It is safe: the sole non-test caller was
`scripts/reinitialize_booking_token.js`, which is itself deleted in §6, and no
bot, app-service, or UI code references it.

### 5.3 Capability removal must not silently succeed

`contracts/partner/PartnerConfiguration.sol:193-207` —
`_removeServiceCapability` breaks on match but falls through and returns
successfully when the capability is absent. `TTMAccount.sol:520` then emits
`ServiceCapabilityRemoved` unconditionally, so indexers record removals that
never happened.

Track a `found` flag and `revert CapabilityDoesNotExist(serviceHash,
capability)`, matching the revert-on-absent convention used by every other
remove in the file (`:134-137`, `:321-325`, `:438-441`).

This is a behaviour change: callers removing a non-existent capability now
revert instead of silently succeeding. It is a bug fix, and no consumer is
known to depend on the silent-success path.

## 6. Dead code removal

All entries verified unreferenced during the review.

### Files

```
utils/cheques.js                        # EIP-712 helpers for MessengerCheque; zero requires
examples/sign_primitive.js              # calls TTMAccount.verifyCheque(), which no longer exists
scripts/status.js                       # calls 3 deleted manager getters; superseded by `yarn hardhat manager status`
scripts/reinitialize_booking_token.js   # legacy Camino migration; would corrupt a fresh deploy
contracts/test/ServiceFeeToken.sol      # self-declared "deprecated and removed"
test/ServiceFeeToken.test.js            # single comment line
go/contracts/servicefeetoken/           # 24-byte orphan — MUST be deleted by hand
```

`go/contracts/servicefeetoken/ServiceFeeToken.go` is not in
`generate_go_abi.sh`'s `ARTIFACTS` list, so regeneration will never remove it,
and CI's `git status --porcelain` drift check cannot detect an orphaned
committed file.

Deleting `utils/cheques.js` empties `utils/`; remove the directory.

### Solidity symbols

```
TTMAccount.sol:7                import { ERC1967Proxy }                  (unused)
TTMAccount.sol:196              error PrefundNotSpentYet(...)            (0 revert sites)
PartnerConfiguration.sol:81     error InvalidPublicKeyUseType(uint8)     (0 revert sites)
BookingToken.sol:257            error InsufficientAllowance(...)         (0 revert sites)
BookingTokenCancellable.sol:96  error CancellationProposalExists(...)    (0 revert sites)
BookingTokenOperator.sol:5      import { CancellationProposalStatus }    (unused)
TTMAccountManagerTest.sol:18-20 commented-out setTTMAccountInfo block
```

`ITTMAccountManager.sol:14` (`getRegisteredServiceNameByHash`) is **retained**.
It is unexercised through the interface, but `ITTMAccountManager` is a
published integration surface and keeping it is defensible.

### Dead test helpers

```
test/utils/fixtures.js:16-18, 38-40   signers developerWallet, developerWalletAdmin, feeAdmin
```

Rename the `chequeOperator` signer (`fixtures.js:19,41`,
`test/TTMAccount.test.js:137`) to `botOperator` — it is used once, as a
generic bot address.

## 7. Documentation and comments

Every item below describes behaviour the code does not have.

```
TTMAccountManager.sol:32-42     "approve the service fee token / prefund" + "Developer Fee" paragraph
TTMAccountManager.sol:232-237   KYC/KYB claim (Camino-only) + prefund-approval claim
BookingTokenCancellable.sol:42  derivation comment says "...CancellableV2"; the constant is the
                                hash of "...BookingTokenCancellable" (no V2) — verified by recomputation
TTMAccount.sol:212              "@2025-08-28 prefundAmount removed" tombstone
ITTMAccount.sol:5               same tombstone
TTMAccount.sol:34-37            gas-money limit described as account-wide; it is per-bot (see backlog)
README.md:59-62                 "developer fee" in the TTMAccountManager description
CLAUDE.md                       same claim; also says `yarn compile` runs contract-sizer but omits docgen
tasks/manager.js:52, :265       tombstone comments
tasks/account.js:53, :1111      tombstone comments
test/** (26 sites)              base64 token URI decoding to "Camino Messenger BookingToken Test"
test/GasMoneyManager.test.js    14 "CAM" comments; 5 "prefund spent" comments for a removed check
test/utils/fixtures.js:175      "// Deposit CAM" → ETH
test/utils/fixtures.js:186,242  "// Deposit service fee token" → NullUSD
```

The `BookingTokenCancellableV2` comment is the highest-priority item here: the
value is correct, so anyone who "fixes" the constant to match the comment
silently relocates the proposals mapping.

### Stale `FIXME`s at `BookingToken.sol:587,590`

```solidity
// FIXME: Define a reason in the Travel Token Messenger Protocol and update this
_rejectCancellation(owner, supplier, tokenId, 99, 1);
```

**Delete both comments; leave the code unchanged.** The reason has since been
defined in the protocol:

```proto
// AUTOMATIC REJECTION REASON DURING TRANSFER OF THE TOKEN ON-CHAIN
REJECTION_REASON_TRANSFER_ON_CHAIN = 99;
```

The `99, 1` values are correct and intentional. The comments are leftovers from
before the protocol entry existed, and they currently make correct code look
unfinished — an earlier draft of this review wrongly flagged the values as a
defect on the strength of these comments alone.

Consider replacing the magic numbers with a named constant
(`REJECTION_REASON_TRANSFER_ON_CHAIN = 99`) so the link to the protocol enum is
visible at the call site.

### CLI denominations

`tasks/account.js` uses `CAM`/`nCAM`/`aCAM` units and a user-facing
`--camAmount` flag (`:216, :226-228, :267, :291, :295, :304, :329, :483, :928,
:937, :957, :966`). Base's native coin is ETH. Rename to `--amount` with
`ETH`/`gwei`/`wei` units. This is a breaking CLI change, which is acceptable —
no external automation depends on these task names.

## 8. Deployment

### 8.1 Verification (currently broken)

`hardhat.config.js:54,61` point at deprecated Etherscan V1 endpoints. Confirmed
live:

```
{"status":"0","message":"NOTOK",
 "result":"You are using a deprecated V1 endpoint, switch to Etherscan API V2"}
```

Etherscan V2 support landed in `@nomicfoundation/hardhat-verify@2.0.14`, which
peers `hardhat: ^2.24.1` — **entirely within the Hardhat 2 line.** A previous
upgrade attempt failed because `hardhat-verify@3.x` peers `hardhat: ^3.8.0`
and drags in Hardhat 3, which is a hard breaking change (ESM-only config,
rewritten plugin API, `vars` replaced, matching majors required for
Ignition/toolbox/verify).

| Package | From | To |
|---|---|---|
| `hardhat` | 2.22.17 | 2.28.6 |
| `@nomicfoundation/hardhat-verify` | 2.0.12 | 2.1.3 |

`hardhat-verify@2.1.3` peers `hardhat: ^2.26.0` and requires Node ≥ 20 (CI is
already on 20). After upgrading, **remove the `customChains` block entirely** —
`base` and `base-sepolia` are built in — and collapse `etherscan.apiKey` to a
single unified V2 key.

Hardhat 3 is explicitly out of scope. Revisit after testnet is stable.

### 8.2 Configuration variables

Rename `BASESCAN_API_KEY` → `ETHERSCAN_API_KEY` (V2 uses one unified Etherscan
key across chains) and **remove the `"abc"` default** at `hardhat.config.js:46-47`
so a missing key fails loudly. Keep Hardhat `vars` rather than `.env` — it is
already wired and keeps secrets out of the repo. `BASE_SEPOLIA_URL` is not a
secret and may be a plain environment variable.

### 8.3 Task fixes

- `tasks/manager.js:105` — `require(taskArgs.json)` resolves relative paths
  against `tasks/`, not cwd, so `--json ./services/00_initial.json` crashes
  with `MODULE_NOT_FOUND`. Use `path.resolve(process.cwd(), taskArgs.json)`.
  This is in the deploy critical path.
- `tasks/manager.js:78` — `error.data.data` throws a `TypeError` on ethers v6
  error shapes lacking `.data`, masking the real error and aborting the
  63-service registration loop mid-way. Use `error.data?.data`.

### 8.4 Ignition module

`ignition/modules/messenger.js:12-13` — uncomment the `managerAdmin` parameter
so the admin is overridable from the parameters file. Currently `admin` is
hard-wired to `m.getAccount(0)` and cannot be overridden.

### 8.5 Runbook

The system is **inert after deployment** — `SERVICE_REGISTRY_ADMIN_ROLE` has
zero members and zero of 63 services are registered, because `initialize`
grants only DEFAULT_ADMIN/PAUSER/UPGRADER/VERSIONER. This is accepted as-is;
this spec documents rather than automates it.

Add to `README.md`:

1. Set `ETHERSCAN_API_KEY` and `BASE_SEPOLIA_URL`.
2. `yarn hardhat ignition deploy ignition/modules/messenger.js --network base_sepolia --parameters ignition/base_sepolia_parameters.json`
3. `yarn hardhat manager role:grant --role SERVICE_REGISTRY_ADMIN_ROLE --address <deployer> --network base_sepolia`
4. `yarn hardhat manager services:register --json $PWD/services/00_initial.json --network base_sepolia`
   — **absolute path required.** 63 transactions, ~182k gas each ≈ 11.5M total.
5. Grant `MIN_EXPIRATION_ADMIN_ROLE` if the 60s default needs changing.
6. Grant `PAUSER_ROLE` on `BookingToken` to the operations key.
7. `yarn hardhat ignition verify chain-84532`. Marking the two `ERC1967Proxy`
   addresses as *proxies* on Basescan is a separate manual UI step.
8. Transfer `DEFAULT_ADMIN_ROLE`, `UPGRADER_ROLE`, `VERSIONER_ROLE` to the Safe;
   verify the Safe can act; **then** renounce the deployer's roles.
9. Commit `ignition/deployments/chain-84532/`. The UI reads
   `deployed_addresses.json` via `ui/scripts/sync-contracts.ts:35-41` and
   `ui/src/config/chains.ts:31` filters enabled chains by `hasContracts(id)` —
   until this lands the UI has zero enabled chains. Pushing to `dev` also
   triggers `deploy-ui.yml`.
10. Fill in the address table at `README.md:14-16`.
11. Bump the contracts Go module in the bot and matrix-app-service.

Step 8's ordering matters: renouncing before verifying the Safe works is
unrecoverable on a singleton manager.

## 9. CI

`.github/workflows/ci.yaml`

- **Add an `abi/` drift job.** `abi/` is committed and is the UI's sole ABI
  source, but nothing runs `export-abi` and diffs. Mirror the existing `docs`
  job. (Currently in sync — this guards it.)
- **Add a `push` trigger.** CI runs only on `pull_request`, so merge commits on
  `dev`/`main` are never validated.
- **Fix the `lint` script.** `package.json:13` chains prettier → eslint →
  solhint, but neither eslint nor solhint is installed and neither has a config
  file, so `yarn lint` fails 100% of the time while `README.md:43` and
  `CLAUDE.md:30` both document it. Either add the dependencies and configs, or
  reduce `lint` to prettier and correct both docs. **Decision: add
  `eslint` + `solhint` with minimal configs**, since the docs advertise them and
  Solidity linting has real value pre-deploy.
- Bump `actions/setup-node@v3` → v4 (lines 15, 33, 81, 115 — line 51 is already
  v4) and `actions/setup-go@v4` → v5.
- Add `/.superpowers/` to `.prettierignore` so local `yarn lint:prettier` is not
  permanently red.

## 10. Testing

Existing suite: 120 passing. It must stay green at the final compiler settings.

New tests required:

| Area | Test |
|---|---|
| §2.2 | Gas-money packing: limit/period round-trip through `getGasMoneyWithdrawal()` and `getGasMoneyWithdrawalForAccount()` still returns `uint256`; values at the `uint128`/`uint64` bounds survive; **overflow reverts rather than truncating** in both `setGasMoneyWithdrawal` and `__GasMoneyManager_init` |
| §4 | `pause()` blocks mint/buy/finalize; `unpause()` restores; non-`PAUSER_ROLE` cannot pause |
| §5.1 | `receive()` emits `Deposit(sender, amount)`; account funding at creation emits it too |
| §5.2 | `symbol()` is `BToken` at initialization, with no post-deploy step; the three `reinitializeV2` assertions are removed |
| §5.3 | Removing an absent capability reverts `CapabilityDoesNotExist`; no event emitted |
| §3.3 | Each `initialize` reverts on a zero address for every role parameter |

Update `test/BookingToken.test.js:45` and
`ui/src/pages/tabs/BookingTokenTab.test.tsx:11,49` for the symbol change.

## 11. Accepted risks

### Permissionless account creation on Base Sepolia

`TTMAccountManager.createTTMAccount` (`:243`) is `external payable
nonReentrant whenNotPaused` with **no role gate**. Its NatSpec claims callers
must be KYC/KYB verified — that was **chain-enforced on Camino** and does not
exist on Base. The Camino design also charged a prefund, which made spam
expensive; that was removed with the fee refactor.

Consequently, on Base anyone can create a `TTMAccount` for gas, and since
`BookingToken`'s authorization model is `onlyTTMAccount`, mint arbitrary
booking tokens, reserve them against real partner accounts, and pollute the
ecosystem NFT and the UI activity feed.

**Accepted for Base Sepolia.** The blast radius is spam and feed noise, not
loss of funds. The misleading NatSpec is corrected in §7 so the gap is not
mistaken for a guarantee.

**This must be closed before Base mainnet.** The gating model is the first
item in the decision document.

### Deferred storage-layout change

Collapsing the dual source of truth for "is a TTM account" (the
`_ttmAccountInfo.isTTMAccount` bool vs `TTMACCOUNT_ROLE`) is a storage-layout
change and is deferred to the backlog. Cost of deferring: one wasted slot per
account, and the two views can diverge if an admin grants `TTMACCOUNT_ROLE`
manually. Not dangerous — `BookingToken` authorizes off the mapping, which has
no external setter — but it means the UI and the contract can disagree about
which accounts exist.

## 12. Regeneration

After any contract change:

```sh
yarn compile && yarn docgen && yarn hardhat export-abi && bash scripts/generate_go_abi.sh
```

The `docs` and `go-bindings` CI jobs both assert `git status --porcelain` is
empty. `go/contracts/servicefeetoken/` must be removed by hand — the generator
will not do it.
