# Pre-Deployment Review — travel-token-messenger-contracts

**Date:** 2026-07-21
**Branch:** `rebranding` @ `43cd7f9`
**Context:** Fresh from-scratch Base Sepolia (84532) deployment. Nothing in
production. Breaking changes, storage-layout changes, and ABI changes are all
explicitly permitted — this is the last cheap moment to make them.

## Method

Four independent reviewers swept the repo in parallel (dead-code, security +
upgradeability, architecture, tooling/deploy-readiness). Every load-bearing
claim below was then re-verified directly. Items marked **[verified]** were
confirmed by first-hand command output, not taken on a reviewer's word.

Baseline health: `yarn compile` succeeds, `yarn test` → **120 passing**,
`abi/` + `docs/` + `go/contracts/` all confirmed in sync with `contracts/`.
The tooling reviewer additionally ran a full Ignition deploy against a local
node, which is where the post-deploy findings come from.

### Conflicts between reviewers, resolved

Three reviewers disagreed on points that matter. Resolutions:

1. **`BookingTokenCancellable.sol:42` `V2` label.** The architecture reviewer
   read this as a *deliberate* namespace rotation and praised it. It is not.
   Three independent keccak computations (this review, the dead-code reviewer,
   the security reviewer) agree the constant `0x1973af82…` is the hash of
   `traveltoken.messenger.storage.BookingTokenCancellable` — **without** `V2`.
   The code is correct; the derivation comment is wrong. Treat as a stale
   comment, and a dangerous one.
2. **`ServiceFeeToken` liveness.** The tooling reviewer claimed it survives as
   a test mock. It does not. The only references in the repo are two tombstone
   *comments* (`tasks/manager.js:52`, `tasks/account.js:53`); there is no
   `getContractFactory("ServiceFeeToken")` anywhere and the test file is a
   single comment line. **[verified]** — fully dead.
3. **`__ReentrancyGuard_init()`.** Not a contradiction — the reviewers were
   describing different contracts. `TTMAccount.initialize` *does* call it
   (line 221). `TTMAccountManager.initialize` does **not**, despite using
   `nonReentrant`. **[verified]**

---

## 1. Deploy blockers

These either break the deploy or leave a non-functional system.

| # | Finding | Location |
|---|---------|----------|
| B1 | **Basescan verification is broken.** Configured `apiURL`s are deprecated Etherscan V1 endpoints; a live curl returns `"You are using a deprecated V1 endpoint"`. Installed `@nomicfoundation/hardhat-verify` is **2.0.12**, which predates V2 support — a URL edit alone will not fix it. Bump to ≥2.0.14 (2.1.x preferred), then drop the `customChains` block entirely (`base`/`base-sepolia` are built in). **[verified]** | `hardhat.config.js:54,61` |
| B2 | **`BASESCAN_API_KEY` and `BASE_SEPOLIA_URL` are unset.** The API key silently falls back to the literal string `"abc"`, producing a confusing auth error rather than "you forgot to set it". The URL falls back to a public RPC that will rate-limit during the 63 `registerService` transactions. | `hardhat.config.js:46-47` |
| B3 | **The deploy leaves an inert system.** After the Ignition module completes, `SERVICE_REGISTRY_ADMIN_ROLE` has **0 members** and **0 of 63** services are registered — `initialize` grants only DEFAULT_ADMIN/PAUSER/UPGRADER/VERSIONER. Nothing can be registered until the role is granted manually. Not automated, not documented anywhere. | `TTMAccountManager.sol:189-203` |
| B4 | **`scripts/reinitialize_booking_token.js` corrupts a fresh deploy.** It force-renames the symbol to `BToken`, but `initialize` now sets `("BookingToken", "TRIP")`. Running it downgrades TRIP → BToken and permanently consumes the `reinitializer(2)` slot. It is a legacy Camino migration. **Delete it.** | `scripts/reinitialize_booking_token.js:22-31` |
| B5 | **`services:register` silently requires an absolute path.** `require(taskArgs.json)` resolves relative paths against `tasks/`, not cwd — `--json ./services/00_initial.json` crashes with `MODULE_NOT_FOUND`. This sits directly in the deploy critical path. Fix with `path.resolve(process.cwd(), …)`. | `tasks/manager.js:105` |
| B6 | **`managerAdmin` is commented out in the Ignition module**, so `admin` is hard-wired to `m.getAccount(0)` and the parameters file *cannot* override it. The deployer EOA ends up holding DEFAULT_ADMIN + PAUSER + UPGRADER + VERSIONER on the manager and DEFAULT_ADMIN + UPGRADER on BookingToken. Decide this deliberately rather than by default. | `ignition/modules/messenger.js:12-13` |

---

## 2. Decisions frozen at deploy

Storage layout and event signatures are permanent once deployed. These must be
settled **before** the transaction, not after.

### 2.1 Storage layout

- **Delete the `_unused` slot.** `TTMAccount.sol:123` reserves a slot with the
  comment *"Not used, but do not remove."* That advice was correct while
  upgrading a live proxy; a from-scratch deploy is precisely when it lifts.
  Never read or written anywhere. **[verified]**
- **Pack `GasMoneyStorage`.** `_withdrawalLimit`/`_withdrawalPeriod`
  (`GasMoneyManager.sol:26-31`) are two full `uint256` slots; a wei limit fits
  `uint128` and a seconds period fits `uint64`. Packing removes a cold SLOAD
  from every bot gas withdrawal — the highest-frequency call in the system.
- **Collapse the dual source of truth for "is a TTM account."**
  `_createTTMAccount` writes both the `_ttmAccountInfo` mapping *and* grants
  `TTMACCOUNT_ROLE`. `BookingToken` authorizes off the mapping; the UI
  enumerates off the role (`ui/src/hooks/useMyAccounts.ts:42`). Because
  `TTMACCOUNT_ROLE`'s admin defaults to `DEFAULT_ADMIN_ROLE`, the manager admin
  can grant/revoke it directly and desynchronize the two views. Not spoofable
  (authorization is unaffected), but `ui/src/components/AccountValidityNotice.tsx`
  already exists purely to paper over this ambiguity.
- **Add `AccessControlDefaultAdminRules`** (2-step admin transfer) to the
  manager and BookingToken. Today a fat-fingered `grantRole` + `renounceRole`
  on the *singleton* manager permanently bricks account creation, service
  registration, and implementation versioning with no recovery path.
- **Add `Pausable` to `BookingToken`.** Pausability currently exists only on
  the manager and covers only account creation. The actual value flow
  (`buyReservedToken`, `finalizeCancellation`) has no pause lever at all; the
  only response to a discovered bug would be a full upgrade.

### 2.2 Event signatures

- **`Deposit` is declared but never emitted.** `TTMAccount.sol:239` is literally
  `receive() external payable {}`. Meanwhile `ui/src/lib/activity/catalog.ts:174`
  has a fully built renderer with a passing test at `catalog.test.ts:42`. Every
  ETH transfer into an account is invisible to the activity feed — and because
  `TTMAccountManager` forwards `msg.value` through this same silent `receive()`,
  account funding at creation is unobservable too. **[verified]** Fix:
  `receive() external payable { emit Deposit(msg.sender, msg.value); }`
- **`string indexed serviceName` makes the name unrecoverable.** Seven events
  (`TTMAccount.sol:167-177`) index a dynamic type, which stores only
  `keccak256(value)` in the topic and nothing in data. The UI already carries
  dedicated machinery (`SERVICE_HASH_EVENTS` + `useAccountActivity.ts:32-65`) to
  re-resolve the hash against the registry before it can render a sentence.
  Emit `bytes32 indexed serviceHash, string serviceName` instead.
- **Widen `TTMAccountCreated`** to include `creator` and `admin`; today an
  indexer must make a follow-up call per account.
- ~~**Replace the hardcoded `FIXME` reason code `99`.**~~ **Withdrawn — this
  finding was wrong.** `BookingToken.sol:587,590` auto-emit cancellation reason
  `99, 1`, and reason 99 *is* defined in the protocol as
  `REJECTION_REASON_TRANSFER_ON_CHAIN = 99` ("automatic rejection reason during
  transfer of the token on-chain"). The code is correct and the value is
  intentional; only the stale `FIXME` comments should be deleted. The reviewer
  inferred the enum gap without checking the protocol repo, and I repeated it.

### 2.3 Compiler settings

`optimizer.runs: 1` optimizes *deployment* cost for contracts that are deployed
once and called constantly — backwards for this system. The premise that
contracts are near the EIP-170 limit is false: measured headroom is **6.1 KiB**
at `runs: 1` and still **3.3 KiB** at `runs: 1000`. Recommend **`runs: 1000`**
(or 500 to leave margin for the Pausable/AccessControl additions above). Not
changeable after deploy without redeploying every implementation.

`evmVersion: "paris"` is safe but leaves gas on the table; Base Sepolia has
supported Cancun since Dencun. Optional, and requires a full test re-run.

---

## 3. Security findings

The repo has prior professional audits and the fee/cheque removal was
**mechanically clean** — no dangling fee math, no orphaned auth checks, no
half-migrated storage. Reentrancy protection was verified *not* weakened by
deleting `ChequeManager` (the guard was re-added directly to `TTMAccount` in
the same commit). Findings below are overwhelmingly pre-existing, not
refactor-induced.

### High

**H-1 — Cancellation refund can permanently stick.**
`BookingToken.sol:824-850` / `processPayment:486-512`. Once a token is `BOUGHT`,
`transferFrom` places no restriction on the recipient — only mint/buy enforce
`onlyTTMAccount`. `processPayment` pushes the refund with
`payable(recipient).sendValue(...)` **[verified]**. If the current holder is a
contract without a payable fallback (custody wrapper, multisig, vault) or the
ERC-20 path hits a blacklisting stablecoin, `finalizeCancellation` reverts every
time. The state transition and the payment are in the same call with no
try/catch, no pull-payment fallback, and no admin override — so the cancellation
is stuck permanently and the escrowed refund is trapped.
*Fix:* pull-payment for the refund leg, or a guarded `forceResolveCancellation`.
*Pre-existing (`69bc05a`), not from the fee removal.*

### Medium

- **M-1 — Approved-operator transfers are DoS'd.** `checkTransferable`
  (`BookingToken.sol:567-625`) authorizes auto-resolution on raw `msg.sender`.
  A marketplace or custody contract holding a valid `setApprovalForAll` cannot
  transfer while any proposal is `PENDING`, because `onlyOwnerOrSupplier`
  rejects it — even though the true owner authorized the transfer. Breaks
  composability with any marketplace built on `BookingToken`.
- **M-2 — Gas-money limit is per-bot, not per-account.**
  `GasMoneyManager.sol:89-119` keys the counter on `msg.sender`, while the
  NatSpec reads as an account-wide allowance. N bots ⇒ up to `N × 10 ETH` per
  24h can leave one account.
- **M-3 — The 3-role bot bundle maximizes single-key blast radius.**
  `addMessengerBot` grants `MESSENGER_BOT_ROLE` + `BOOKING_OPERATOR_ROLE` +
  `GAS_WITHDRAWER_ROLE` together. One compromised hot key — the realistic threat
  model for a hosted bot — yields minting, cancellation, and fund-withdrawal
  capability at once.

### Elevated from the reviewers' rating

**Permissionless `createTTMAccount` on Base.** `TTMAccountManager.sol:243` is
`external payable nonReentrant whenNotPaused` with **no role gate**
**[verified]**, and its NatSpec still claims *"it reverts if the caller is not
KYC or KYB verified."* That was **chain-enforced on Camino** and simply does not
exist on Base. Since `BookingToken`'s entire authorization model is
`onlyTTMAccount`, anyone can mint themselves an account for gas and then mint
arbitrary booking tokens with arbitrary URIs, reserve them against real partner
accounts, and pollute the ecosystem NFT and the UI activity feed.

This is a **behavioral regression introduced by the chain migration**, not by
the code — which is exactly why it is easy to miss. Recommend gating behind an
`ACCOUNT_CREATOR_ROLE`; partner onboarding is already a manual business process.

### Low

- **L-2 — The supported-token allowlist is never enforced.** `_supportedTokens`
  is only ever `.add`/`.remove`/`.values()` — never `.contains()` for a check,
  and `BookingToken` has **no reference to `PartnerConfiguration` at all**
  **[verified]**. Any ERC-20 can be used regardless of what the account
  declares. Either enforce it or document it as advisory metadata.
- **L-3 — `recordExpiration`'s role gate is cosmetic.**
  `BookingToken.recordExpiration` is `public` with no restriction, so
  `TTMAccount`'s `onlyRole(BOOKING_OPERATOR_ROLE)` wrapper protects nothing.
  Harmless (objective, time-gated state transition) but should be documented as
  deliberately permissionless.
- **L-4 — No zero-address checks in any `initialize()`.** Granting a role to
  `address(0)` silently succeeds and is permanently unusable — a deploy-time
  footgun given B6.
- **L-5 — `TTMAccountManager.initialize` never calls `__ReentrancyGuard_init()`**
  despite `createTTMAccount` using `nonReentrant` **[verified]**. Functional
  under OZ v5 (uninitialized `_status == 0 != ENTERED`) but pays a cold-slot
  SSTORE on first use.

---

## 4. Dead code — delete list

All entries verified unreferenced. Regenerate artifacts afterwards or CI fails.

### Files

```
utils/cheques.js                        # EIP-712 helpers for MessengerCheque; zero requires. Deletes utils/ entirely
examples/sign_primitive.js              # calls TTMAccount.verifyCheque(), which no longer exists
scripts/status.js                       # calls 3 deleted manager getters; superseded by `yarn hardhat manager status`
scripts/reinitialize_booking_token.js   # see B4 — actively harmful
contracts/test/ServiceFeeToken.sol      # self-declared "deprecated and removed"
test/ServiceFeeToken.test.js            # single comment line
go/contracts/servicefeetoken/           # 24-byte orphan — see note below
```

> `go/contracts/servicefeetoken/ServiceFeeToken.go` is **not** in
> `generate_go_abi.sh`'s `ARTIFACTS` list, so regeneration will never remove it,
> and CI's drift check (`git status --porcelain`) cannot detect an orphaned
> committed file. **Must be deleted by hand.** **[verified]**

### Solidity symbols

```
TTMAccount.sol:7             import { ERC1967Proxy }                    (unused)
TTMAccount.sol:123           uint256 _unused;                           (reclaim the slot)
TTMAccount.sol:196           error PrefundNotSpentYet(...)              (0 revert sites)
PartnerConfiguration.sol:81  error InvalidPublicKeyUseType(uint8)       (0 revert sites)
PartnerConfiguration.sol:99  __PartnerConfiguration_init()              (never called)
PartnerConfiguration.sol:101 __PartnerConfiguration_init_unchained()    (never called)
ServiceRegistry.sol:58       __ServiceRegistry_init()                   (never called)
ServiceRegistry.sol:60       __ServiceRegistry_init_unchained()         (never called)
BookingToken.sol:257         error InsufficientAllowance(...)           (0 revert sites)
BookingTokenCancellable.sol:96  error CancellationProposalExists(...)   (0 revert sites)
BookingTokenOperator.sol:5   import { CancellationProposalStatus }      (unused)
ITTMAccountManager.sol:14    getRegisteredServiceNameByHash             (interface entry only)
TTMAccountManagerTest.sol:18-20  commented-out setTTMAccountInfo block
```

All four errors confirmed with **0** `revert` sites; both imports confirmed by
identifier count of 1 (the import line itself). **[verified]**

### Correctness fixes surfaced by the sweep

- `BookingTokenCancellable.sol:12` — `contract` should be `abstract contract …
  is Initializable`. It is currently a stray deployable artifact.
- `PartnerConfiguration.sol:193-207` — `_removeServiceCapability` silently
  no-ops when the capability isn't found, yet `TTMAccount.sol:520` emits
  `ServiceCapabilityRemoved` unconditionally **[verified]**. Indexers record
  removals that never happened. Add a `found` flag and revert, matching the
  revert-on-absent convention used by every other remove in the file.

### Stale text

```
TTMAccountManager.sol:32-42    "approve the service fee token / prefund" + whole "Developer Fee" paragraph
TTMAccountManager.sol:232-237  KYC/KYB claim (Camino-only) + prefund-approval claim
BookingTokenCancellable.sol:42 derivation comment says "…CancellableV2"; actual preimage has no V2
TTMAccount.sol:212             "@2025-08-28 prefundAmount removed" tombstone
ITTMAccount.sol:5              same tombstone
README.md:59-62                "developer fee" in the TTMAccountManager description
CLAUDE.md                      same claim; also says `yarn compile` runs contract-sizer but omits docgen
test/**                        26 sites: base64 token URI decoding to "Camino Messenger BookingToken Test"
test/GasMoneyManager.test.js   14 "CAM" comments; 5 "prefund spent" comments for a check that no longer exists
test/utils/fixtures.js:16-18,38-40  dead signers developerWallet, developerWalletAdmin, feeAdmin
tasks/account.js               --camAmount flag + CAM/nCAM/aCAM units (user-facing CLI on a Base-only deploy)
```

---

## 5. Tooling and CI

- **`yarn lint` cannot work.** `package.json:13` chains prettier → eslint →
  solhint, but **neither eslint nor solhint is installed and neither has a
  config file** **[verified]**. Both `README.md:43` and `CLAUDE.md:30` tell you
  to run it. CI masks the failure by running `lint:prettier` only. Either add
  the deps + configs or reduce the script to prettier.
- **No `abi/` drift check in CI.** `abi/` is committed and is the UI's sole ABI
  source, but nothing runs `export-abi` and diffs. Currently in sync, but
  unguarded. Mirror the existing `docs` job.
- **CI only triggers on `pull_request`** — no `push` trigger, so merge commits
  on `dev`/`main` are never validated.
- Outdated actions: `setup-node@v3` (lines 15, 33, 81, 115 — inconsistent, line
  51 already uses v4), `setup-go@v4`. Node pinned to 20.
- `.superpowers/` missing from `.prettierignore`, so local `yarn lint:prettier`
  is permanently red.
- `tasks/manager.js:78` — `error.data.data` will `TypeError` on ethers v6 error
  shapes lacking `.data`, masking the real error and aborting the 63-service
  registration loop mid-way. Use `error.data?.data`.
- Staying on **Hardhat 2 is correct** for this deploy. Hardhat 3 is a hard
  breaking change across config, plugins, Ignition, and verify. Revisit after
  testnet is stable.

---

## 6. Deployment runbook

Currently the README's deploy section is one command and implies you are done.
The actual sequence:

1. Fix B1 (bump `hardhat-verify`, drop `customChains`) and B2 (set
   `BASESCAN_API_KEY`, `BASE_SEPOLIA_URL`).
2. Delete `scripts/reinitialize_booking_token.js` (B4) so nobody runs it.
3. Decide the admin (B6) — uncomment `managerAdmin` or accept deployer-as-admin.
4. `yarn hardhat ignition deploy ignition/modules/messenger.js --network base_sepolia --parameters ignition/base_sepolia_parameters.json`
   *(`base_sepolia_parameters.json` is currently `{}` — every parameter has a
   default, so `--parameters` is a no-op today.)*
5. `yarn hardhat manager role:grant --role SERVICE_REGISTRY_ADMIN_ROLE --address <deployer> --network base_sepolia`
6. `yarn hardhat manager services:register --json $PWD/services/00_initial.json --network base_sepolia`
   — **absolute path required** (B5). 63 transactions, ~182k gas each ≈ 11.5M total.
7. Grant `MIN_EXPIRATION_ADMIN_ROLE` on BookingToken if you need to change the
   60s default; `initialize` grants only DEFAULT_ADMIN + UPGRADER.
8. `yarn hardhat ignition verify chain-84532`. Marking the two `ERC1967Proxy`
   addresses as *proxies* on Basescan is a separate manual UI step.
9. **Commit `ignition/deployments/chain-84532/`.** The UI reads
   `deployed_addresses.json` via `ui/scripts/sync-contracts.ts:35-41`, and
   `ui/src/config/chains.ts:31` filters enabled chains by `hasContracts(id)` —
   until that JSON lands, the UI has **zero enabled chains**. Pushing to `dev`
   also triggers `deploy-ui.yml`.
10. Fill in the address table at `README.md:14-16`.
11. Propagate addresses to the bot and matrix-app-service (tracked in `TODOS.md`).

**Wiring that is already automated** (no action needed): `setAccountImplementation`
and `setBookingTokenAddress` both run inside the Ignition module and were
confirmed correct post-deploy.

---

## 7. Larger changes worth considering

These are real improvements but each carries meaningful churn. Listed so the
decision is explicit rather than defaulted.

- **Beacon proxy instead of per-account UUPS.** Today `_authorizeUpgrade`
  (`TTMAccount.sol:281-297`) *requires* `newImplementation ==
  manager.getAccountImplementation()`. So account holders already cannot choose
  their implementation or roll back — the only freedom per-account UUPS actually
  grants is the freedom to *not* upgrade, i.e. indefinite version drift against
  a shared manager/BookingToken ABI. You pay full UUPS cost (~2.5 KiB per
  account, N transactions per upgrade) for a liability. A beacon makes upgrades
  atomic; put it behind a `TimelockController` so partners get a public window,
  which is strictly more protection than today. Cost: one extra staticcall per
  call (noise on Base). **Biggest single call to make before deploy** — after
  deploy it means migrating every account.
- **Hash-native account API + a view/lens contract.** `TTMAccount.sol:427-635`
  is largely a string↔hash adapter; `getSupportedServices()` makes one external
  call *per service*. Moving name resolution to a stateless lens drops ~250
  lines and the per-read fan-out. ABI-breaking for the Go bot, app-service, and UI.
- **Service identity scheme.** Hashing the full versioned protobuf FQN means a
  `v4→v5` bump is an unrelated 32-byte value requiring every account to
  re-register, and deprecation is a dead end (`_unregisterServiceName` leaves
  the name mappings populated, so existing accounts advertise deprecated
  services forever with no event). Minimum viable fix: add a `deprecated` flag +
  `ServiceDeprecated(id, supersededBy)` event rather than silently mutating the set.
- **Test coverage gaps.** 120 passing is solid, but nothing tests: that
  `createTTMAccount` is permissionless (the most consequential property on
  Base), the role/mapping divergence, `Deposit` emission, `removeAllServices`,
  or the `_removeServiceCapability` no-op. Items 1–3 guard decisions being made
  right now.

---

## After any contract change

```sh
yarn compile && yarn docgen && yarn hardhat export-abi && bash scripts/generate_go_abi.sh
```

The `docs` and `go-bindings` CI jobs both assert `git status --porcelain` is
empty. Remember `go/contracts/servicefeetoken/` must be removed by hand.
