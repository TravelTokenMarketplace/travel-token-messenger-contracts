# Pre-Deployment Hardening Implementation Plan

> **For agentic workers:** REQUIRED SUB-SKILL: Use superpowers:subagent-driven-development (recommended) or superpowers:executing-plans to implement this plan task-by-task. Steps use checkbox (`- [ ]`) syntax for tracking.

**Goal:** Prepare the Travel Token Messenger contracts for a fresh Base Sepolia deployment by fixing everything that requires no ABI break, no consumer coordination, and no external sign-off.

**Architecture:** Solidity 0.8.24 UUPS-upgradeable contracts built with Hardhat 2. Changes fall into five groups: compiler settings, ERC-7201 storage cleanup (irreversible once deployed), internal correctness fixes, additive behaviour (Pausable, `Deposit` event), and tooling/CI/docs. Every external function signature is preserved except one deliberate removal (`reinitializeV2`).

**Tech Stack:** Solidity 0.8.24, Hardhat 2.28.6, OpenZeppelin Contracts (Upgradeable) 5.1.0, ethers v6, chai, Hardhat Ignition, abigen, solidity-docgen, hardhat-abi-exporter.

**Spec:** `docs/superpowers/specs/2026-07-21-predeploy-hardening-design.md`

## Global Constraints

- **Never break an external ABI signature.** The bot, matrix-app-service, and UI consume `abi/` and `go/contracts/`. The only permitted removal is `BookingToken.reinitializeV2`.
- **Storage layout changes are irreversible after deployment.** Tasks 5 and 6 must land before deploy or never.
- Solidity version is pinned at `0.8.24`. Do not change it.
- Target chains: Base Sepolia (84532) and Base (8453). No other network config may be added.
- Never introduce a `require(string)`. This codebase uses custom errors exclusively.
- ERC-7201 storage-location constants must never be edited. If a namespace string changes, recompute the constant — do not hand-edit the hex.
- Run `yarn test` at the end of every task. **120 tests pass at the start of this plan.** Expected running totals:

  | After task | Count | Change |
  |---|---|---|
  | baseline | 120 | — |
  | 6 | 124 | +4 storage packing |
  | 8 | 127 | +3 zero-address |
  | 9 | 130 | +3 pausable |
  | 10 | 131 | +1 deposit |
  | 11 | 128 | −3 removed `reinitializeV2` tests |
  | 12 | 129 | +1 capability |

  If a count does not match, stop and find out why before continuing — a
  silently skipped or double-counted test is a real problem.
- Do NOT run `yarn install` unless a task explicitly says to. Disk is tight on this machine.
- Never use `git checkout`, `git stash`, or `git reset` to undo work — the repo has uncommitted docs on this branch.
- Commit after every task. Work happens on branch `docs/predeploy-hardening-spec` or a branch created from it.

### Disk-space environment (this machine)

`/` and `/home` are ~97% full. Before any Go-bindings work (Task 17), export:

```sh
export TMPDIR=/hgst/work/.ttm-scratch
export GOCACHE=/hgst/work/.ttm-scratch/gocache
export GOTMPDIR=/hgst/work/.ttm-scratch
export GOMODCACHE=/hgst/work/.ttm-scratch/gomod
```

### Artifact regeneration policy

`scripts/generate_go_abi.sh` does its own `rm -rf node_modules && yarn install`, which is slow and disk-heavy. Therefore:

- Contract-changing tasks regenerate **`abi/` and `docs/` only** (fast: `yarn compile` runs docgen automatically; `yarn hardhat export-abi` is quick).
- **Go bindings are regenerated once, in Task 17.**
- Intermediate commits will therefore have a stale `go/contracts/`. This is expected and is resolved by Task 17 before any PR is opened.

---

## File Structure

**Contracts modified**

| File | Responsibility | Tasks |
|---|---|---|
| `contracts/account/TTMAccount.sol` | Per-partner account. Loses `_unused` slot, gains `Deposit` emission, initializer wiring + validation | 3, 5, 7, 8, 10, 13 |
| `contracts/account/GasMoneyManager.sol` | Gas-money allowance. Storage repacked; external signatures preserved | 6, 13 |
| `contracts/account/ITTMAccount.sol` | Account interface. Comment cleanup only | 13 |
| `contracts/manager/TTMAccountManager.sol` | Factory + registry + roles. Initializer fixes, NatSpec rewrite | 7, 8, 13 |
| `contracts/partner/PartnerConfiguration.sol` | Per-account config. Dead error removed; capability removal now reverts | 3, 7, 12 |
| `contracts/partner/ServiceRegistry.sol` | Service name↔hash registry. Initializer wiring | 7 |
| `contracts/booking-token/BookingToken.sol` | ERC-721. Gains Pausable; symbol fixed; `reinitializeV2` removed | 3, 8, 9, 11, 13 |
| `contracts/booking-token/BookingTokenCancellable.sol` | Cancellation state machine. Becomes `abstract` | 3, 13 |
| `contracts/booking-token/BookingTokenOperator.sol` | Linked library. Unused import removed | 3 |
| `contracts/manager/test/TTMAccountManagerTest.sol` | Upgrade-test mock. Dead comment block removed | 3 |

**Deleted outright** (Task 4): `utils/cheques.js`, `utils/` (now empty), `examples/sign_primitive.js`, `scripts/status.js`, `scripts/reinitialize_booking_token.js`, `contracts/test/ServiceFeeToken.sol`, `test/ServiceFeeToken.test.js`, `go/contracts/servicefeetoken/`

**Tooling / config**

| File | Tasks |
|---|---|
| `package.json` | 1, 16 |
| `hardhat.config.js` | 1, 2 |
| `tasks/manager.js` | 13, 15 |
| `tasks/account.js` | 13, 14 |
| `ignition/modules/messenger.js` | 15 |
| `.github/workflows/ci.yaml` | 16 |
| `.prettierignore`, `eslint.config.js`, `.solhint.json` | 16 |
| `README.md`, `CLAUDE.md` | 13, 17 |

**Tests**

| File | Tasks |
|---|---|
| `test/GasMoneyManager.test.js` | 6, 13 |
| `test/TTMAccount.test.js` | 8, 10, 13 |
| `test/TTMAccountManager.test.js` | 8 |
| `test/BookingToken.test.js` | 8, 9, 11, 13 |
| `test/PartnerConfiguration.test.js` | 12 |
| `test/utils/fixtures.js` | 13 |
| `ui/src/pages/tabs/BookingTokenTab.test.tsx` | 11 |

---

## Task 1: Toolchain upgrade and Etherscan V2 verification

Basescan verification is currently broken — the configured endpoints are deprecated Etherscan V1 URLs. Etherscan V2 support landed in `hardhat-verify@2.0.14`; version `2.1.3` is the latest in the Hardhat 2 line and peers `hardhat: ^2.26.0`.

**Do NOT upgrade to `hardhat-verify@3.x`** — it peers `hardhat: ^3.8.0` and drags in Hardhat 3, a hard breaking change. A previous attempt failed this way.

**Files:**
- Modify: `package.json` (devDependencies)
- Modify: `hardhat.config.js:1-11` (add `vars` import), `:44-67` (etherscan block)

**Interfaces:**
- Consumes: nothing
- Produces: working `yarn hardhat verify`; config variable named `ETHERSCAN_API_KEY` (was `BASESCAN_API_KEY`)

- [ ] **Step 1: Confirm the endpoint really is dead**

```bash
curl -s --max-time 20 "https://api-sepolia.basescan.org/api?module=contract&action=getabi&address=0x0000000000000000000000000000000000000000&apikey=abc"
```

Expected: `{"status":"0","message":"NOTOK","result":"You are using a deprecated V1 endpoint, switch to Etherscan API V2 using https://docs.etherscan.io/v2-migration"}`

- [ ] **Step 2: Bump the two packages**

In `package.json` `devDependencies`, change:

```json
"hardhat": "^2.14.0",
```
to
```json
"hardhat": "^2.28.6",
```

and add (the package is currently pulled in transitively by `hardhat-toolbox`; pin it explicitly):

```json
"@nomicfoundation/hardhat-verify": "^2.1.3",
```

- [ ] **Step 3: Install**

This is one of the two tasks permitted to run install.

```bash
yarn install
```

Expected: completes without peer-dependency errors. Verify:

```bash
node -p "require('./node_modules/hardhat/package.json').version"
node -p "require('./node_modules/@nomicfoundation/hardhat-verify/package.json').version"
```

Expected: `2.28.6` (or later 2.x) and `2.1.3` (or later 2.1.x).

- [ ] **Step 4: Add the missing `vars` import**

`hardhat.config.js` uses `vars.get(...)` but never imports it. It currently works via an injected global; make it explicit. After line 5, add:

```javascript
const { vars } = require("hardhat/config");
```

- [ ] **Step 5: Replace the etherscan block**

`base` and `base-sepolia` are built into hardhat-verify, so `customChains` is no longer needed. Etherscan V2 uses one unified key. Replace `hardhat.config.js:44-67` entirely with:

```javascript
    etherscan: {
        apiKey: vars.get("ETHERSCAN_API_KEY"),
    },
```

Note the removed `"abc"` default: a missing key must now fail loudly rather than producing a confusing auth error.

- [ ] **Step 6: Verify config loads and compiles**

```bash
yarn hardhat vars set ETHERSCAN_API_KEY
yarn compile
```

Expected: compilation succeeds. If `ETHERSCAN_API_KEY` is unset, `yarn compile` will fail with a clear message naming the variable — that is the intended behaviour.

- [ ] **Step 7: Run the full suite**

```bash
yarn test
```

Expected: `120 passing`.

- [ ] **Step 8: Commit**

```bash
git add package.json yarn.lock hardhat.config.js
git commit -m "build: upgrade to Hardhat 2.28 and fix Etherscan V2 verification

Configured apiURLs were deprecated Etherscan V1 endpoints; verification
could not succeed. V2 support landed in hardhat-verify 2.0.14.

- hardhat 2.22.17 -> ^2.28.6, hardhat-verify -> ^2.1.3 (still Hardhat 2)
- drop customChains; base/base-sepolia are built in
- unify on ETHERSCAN_API_KEY, remove the 'abc' fallback
- import vars explicitly instead of relying on an injected global"
```

---

## Task 2: Compiler settings

`runs: 1` optimizes deployment cost. These are upgradeable implementations deployed once and called constantly, so runtime cost is what matters. The original justification — fighting the 24 KiB limit while fee support existed — no longer applies.

**Files:**
- Modify: `hardhat.config.js:15-21`

**Interfaces:**
- Consumes: Task 1's config
- Produces: final bytecode settings for deployment

- [ ] **Step 1: Record the baseline sizes**

```bash
yarn compile 2>&1 | tee /tmp/sizes-before.txt
```

Note the `TTMAccount`, `BookingToken`, and `TTMAccountManager` rows from the contract-sizer table.

- [ ] **Step 2: Change the settings**

In `hardhat.config.js`, replace lines 15-21:

```javascript
        settings: {
            optimizer: {
                enabled: true,
                runs: 1000,
            },
            evmVersion: "cancun",
        },
```

- [ ] **Step 3: Recompile and check headroom**

```bash
yarn compile 2>&1 | tee /tmp/sizes-after.txt
```

Expected: all contracts compile. **Check every contract against the 24.576 KiB EIP-170 limit.**

**Decision rule:** if any contract exceeds **22.5 KiB** (leaving under 2 KiB of headroom for the Pausable additions in Task 9), change `runs` to `500` and repeat this step. If still over 22.5 KiB at 500, stop and report — do not proceed to Task 9 without headroom.

- [ ] **Step 4: Run the full suite under the new settings**

```bash
yarn test
```

Expected: `120 passing`. A different `evmVersion` can surface latent issues; if anything fails, report it rather than working around it.

- [ ] **Step 5: Commit**

```bash
git add hardhat.config.js
git commit -m "build: optimize for runtime gas (runs 1->1000, paris->cancun)

runs=1 optimized deployment cost, which is backwards for upgradeable
implementations deployed once and called constantly by messenger bots.
The 24KiB pressure that justified it disappeared with fee removal.

Base Sepolia has supported Cancun since Dencun."
```

---

## Task 3: Remove dead Solidity symbols

Four custom errors have zero `revert` sites and two imports are unreferenced (verified: the identifier appears exactly once in each file — the import line itself). `BookingTokenCancellable` is also a concrete deployable contract that should be abstract.

**Files:**
- Modify: `contracts/account/TTMAccount.sol:7`, `:196`
- Modify: `contracts/partner/PartnerConfiguration.sol:81`
- Modify: `contracts/booking-token/BookingToken.sol:257`
- Modify: `contracts/booking-token/BookingTokenCancellable.sol:12`, `:96`
- Modify: `contracts/booking-token/BookingTokenOperator.sol:5`
- Modify: `contracts/manager/test/TTMAccountManagerTest.sol:18-20`

**Interfaces:**
- Consumes: Task 2's compiler settings
- Produces: smaller ABI surface; `BookingTokenCancellable` no longer deployable

- [ ] **Step 1: Re-verify each symbol is dead before deleting**

```bash
for e in PrefundNotSpentYet InvalidPublicKeyUseType InsufficientAllowance CancellationProposalExists; do
  echo -n "revert $e: "; grep -rn "revert $e" contracts/ --include='*.sol' | wc -l
done
echo -n "ERC1967Proxy in TTMAccount.sol: "; grep -c '\bERC1967Proxy\b' contracts/account/TTMAccount.sol
echo -n "CancellationProposalStatus in BookingTokenOperator.sol: "; grep -c '\bCancellationProposalStatus\b' contracts/booking-token/BookingTokenOperator.sol
```

Expected: all four revert counts `0`; both identifier counts `1`. **If any count differs, stop and report** — the symbol is in use and must not be deleted.

- [ ] **Step 2: Delete the four errors**

Remove these declarations together with their NatSpec comment blocks:

- `contracts/account/TTMAccount.sol:196` — `error PrefundNotSpentYet(uint256 withdrawableAmount, uint256 prefundLeft, uint256 amount);` (also delete the `@notice Error to revert with if the prefund is not spent yet` block above it)
- `contracts/partner/PartnerConfiguration.sol:81` — `error InvalidPublicKeyUseType(uint8 use);`
- `contracts/booking-token/BookingToken.sol:257` — `error InsufficientAllowance(address sender, IERC20 paymentToken, uint256 price, uint256 allowance);`
- `contracts/booking-token/BookingTokenCancellable.sol:96` — `error CancellationProposalExists(uint256 tokenId);`

- [ ] **Step 3: Delete the two unused imports**

`contracts/account/TTMAccount.sol:7` — remove:
```solidity
import { ERC1967Proxy } from "@openzeppelin/contracts/proxy/ERC1967/ERC1967Proxy.sol";
```

`contracts/booking-token/BookingTokenOperator.sol:5` — remove:
```solidity
import { CancellationProposalStatus } from "./BookingTokenCancellable.sol";
```

- [ ] **Step 4: Make `BookingTokenCancellable` abstract**

`contracts/booking-token/BookingTokenCancellable.sol:12`. It currently emits a stray deployable artifact and, unlike every sibling mixin, does not inherit `Initializable`.

```solidity
// FROM:
contract BookingTokenCancellable {

// TO:
abstract contract BookingTokenCancellable is Initializable {
```

Add the import at the top of the file if not already present:

```solidity
import { Initializable } from "@openzeppelin/contracts-upgradeable/proxy/utils/Initializable.sol";
```

- [ ] **Step 5: Delete the dead comment block in the test mock**

`contracts/manager/test/TTMAccountManagerTest.sol:18-20` — remove the commented-out `setTTMAccountInfo` body.

- [ ] **Step 6: Compile and test**

```bash
yarn compile && yarn test
```

Expected: compiles cleanly, `120 passing`. Confirm `BookingTokenCancellable` no longer appears as a deployable contract in the sizer output.

- [ ] **Step 7: Regenerate ABI and docs**

```bash
yarn hardhat export-abi
```

- [ ] **Step 8: Commit**

```bash
git add contracts/ abi/ docs/
git commit -m "refactor: remove dead errors, unused imports, make Cancellable abstract

Four custom errors had zero revert sites; two imports were unreferenced.
BookingTokenCancellable was a concrete deployable contract emitting a
stray artifact and did not inherit Initializable unlike its siblings."
```

---

## Task 4: Remove orphaned files

Cheques and service fees were deleted by the fee-removal refactor, but their helper files survived. `scripts/reinitialize_booking_token.js` is actively dangerous — run against a fresh deploy it would downgrade the symbol and burn the `reinitializer(2)` slot.

**Files:**
- Delete: `utils/cheques.js`, `utils/`, `examples/sign_primitive.js`, `scripts/status.js`, `scripts/reinitialize_booking_token.js`, `contracts/test/ServiceFeeToken.sol`, `test/ServiceFeeToken.test.js`, `go/contracts/servicefeetoken/`

**Interfaces:**
- Consumes: nothing
- Produces: `ServiceFeeToken` fully removed from contracts, tests, ABI, and Go bindings

- [ ] **Step 1: Re-verify nothing references them**

```bash
grep -rn "cheques" --include='*.js' --include='*.ts' --include='*.json' . \
  --exclude-dir=node_modules --exclude-dir=.git --exclude-dir=artifacts --exclude-dir=cache --exclude-dir=docs
grep -rn "ServiceFeeToken" test/ tasks/ scripts/ ignition/ contracts/ examples/
grep -rn 'getContractFactory("ServiceFeeToken")' . --exclude-dir=node_modules --exclude-dir=.git
```

Expected: no `require` of `utils/cheques`; the only `ServiceFeeToken` hits are the two tombstone **comments** at `tasks/manager.js:52` and `tasks/account.js:53`; no `getContractFactory` hit. **If anything else appears, stop and report.**

- [ ] **Step 2: Delete the files**

```bash
git rm utils/cheques.js
git rm examples/sign_primitive.js
git rm scripts/status.js
git rm scripts/reinitialize_booking_token.js
git rm contracts/test/ServiceFeeToken.sol
git rm test/ServiceFeeToken.test.js
git rm -r go/contracts/servicefeetoken/
rmdir utils 2>/dev/null || true
```

`go/contracts/servicefeetoken/` must be removed by hand: it is not in `scripts/generate_go_abi.sh`'s `ARTIFACTS` list, so regeneration will never delete it, and CI's `git status --porcelain` drift check cannot detect an orphaned committed file.

- [ ] **Step 3: Compile and test**

```bash
yarn compile && yarn test
```

Expected: `120 passing` — the deleted test file contained no tests (it was a single comment line).

- [ ] **Step 4: Regenerate ABI (drops the ServiceFeeToken export)**

```bash
yarn hardhat export-abi
git status --short abi/
```

Expected: `abi/contracts/test/ServiceFeeToken.sol/` is gone (`abiExporter.clear: true` handles this).

- [ ] **Step 5: Commit**

```bash
git add -A
git commit -m "chore: delete orphaned cheque and service-fee artifacts

All verified unreferenced. reinitialize_booking_token.js was actively
dangerous: against a fresh deploy it would downgrade the symbol and
consume the reinitializer(2) slot.

go/contracts/servicefeetoken/ removed by hand -- it is not in the
generator's ARTIFACTS list, so the CI drift check could never see it."
```

---

## Task 5: Remove the vestigial prefund storage slot

`TTMAccount.sol:123` reserves a slot with the comment "Not used, but do not remove." That was correct while upgrading a live proxy. A from-scratch deploy is the one moment the constraint lifts — after deployment the slot is frozen again permanently.

**Files:**
- Modify: `contracts/account/TTMAccount.sol:110-124`

**Interfaces:**
- Consumes: nothing
- Produces: `TTMAccountStorage` = `{ address _manager; address _bookingToken; }`

- [ ] **Step 1: Verify the field is never used**

```bash
grep -rn "_unused" contracts/ --include='*.sol'
```

Expected: exactly one hit — the declaration at `TTMAccount.sol:123`. **If there are more, stop and report.**

- [ ] **Step 2: Delete the field**

In `contracts/account/TTMAccount.sol`, the struct becomes:

```solidity
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
```

Do **not** touch `TTMAccountStorageLocation` on line 127 — the namespace string is unchanged, so the constant stays correct.

- [ ] **Step 3: Compile and test**

```bash
yarn compile && yarn test
```

Expected: `120 passing`.

- [ ] **Step 4: Regenerate and commit**

```bash
yarn hardhat export-abi
git add contracts/ abi/ docs/
git commit -m "refactor!: reclaim the vestigial prefund storage slot

TTMAccountStorage._unused held the prefund amount, removed with the fee
refactor. The 'do not remove' comment protected a live proxy's layout;
this fresh deployment is the only chance to reclaim it."
```

---

## Task 6: Pack GasMoneyStorage

Two `uint256` slots hold values that fit in `uint128` and `uint64`. Packing removes a cold SLOAD from every bot gas withdrawal — the highest-frequency call in the system.

**Critical:** every external signature stays `uint256`. Narrowing must not leak into the ABI.

**Files:**
- Modify: `contracts/account/GasMoneyManager.sol:24-29`, `:72-76`, `:89-119`, `:127-133`, `:141-144`, `:153-158`
- Test: `test/GasMoneyManager.test.js`

**Interfaces:**
- Consumes: nothing
- Produces: `getGasMoneyWithdrawal() returns (uint256, uint256)` and `getGasMoneyWithdrawalForAccount(address) returns (uint256, uint256)` — **unchanged signatures**

- [ ] **Step 1: Write the failing tests**

Append inside the `describe("GasMoneyManager")` block in `test/GasMoneyManager.test.js`:

```javascript
    describe("Storage packing", function () {
        it("should still return uint256 from the getters", async function () {
            const { ttmAccount } = await loadFixture(deployAndConfigureAllFixture);
            const [limit, period] = await ttmAccount.getGasMoneyWithdrawal();
            expect(limit).to.equal(ethers.parseEther("10"));
            expect(period).to.equal(86400n);

            const iface = ttmAccount.interface.getFunction("getGasMoneyWithdrawal");
            expect(iface.outputs.map((o) => o.type)).to.deep.equal(["uint256", "uint256"]);

            const accIface = ttmAccount.interface.getFunction("getGasMoneyWithdrawalForAccount");
            expect(accIface.outputs.map((o) => o.type)).to.deep.equal(["uint256", "uint256"]);
        });

        it("should accept values at the uint128/uint64 bounds", async function () {
            const { ttmAccount, signers } = await loadFixture(deployAndConfigureAllFixture);
            const maxLimit = 2n ** 128n - 1n;
            const maxPeriod = 2n ** 64n - 1n;

            await ttmAccount.connect(signers.ttmAccountAdmin).setGasMoneyWithdrawal(maxLimit, maxPeriod);

            const [limit, period] = await ttmAccount.getGasMoneyWithdrawal();
            expect(limit).to.equal(maxLimit);
            expect(period).to.equal(maxPeriod);
        });

        it("should revert rather than truncate when the limit overflows uint128", async function () {
            const { ttmAccount, signers } = await loadFixture(deployAndConfigureAllFixture);
            const tooBig = 2n ** 128n;

            await expect(
                ttmAccount.connect(signers.ttmAccountAdmin).setGasMoneyWithdrawal(tooBig, 86400n),
            ).to.be.revertedWithCustomError(ttmAccount, "GasMoneyValueOutOfRange");
        });

        it("should revert rather than truncate when the period overflows uint64", async function () {
            const { ttmAccount, signers } = await loadFixture(deployAndConfigureAllFixture);
            const tooBig = 2n ** 64n;

            await expect(
                ttmAccount.connect(signers.ttmAccountAdmin).setGasMoneyWithdrawal(ethers.parseEther("10"), tooBig),
            ).to.be.revertedWithCustomError(ttmAccount, "GasMoneyValueOutOfRange");
        });
    });
```

> If `signers.ttmAccountAdmin` is not the correct role holder for `setGasMoneyWithdrawal` in this fixture, read `test/utils/fixtures.js` and the existing `setGasMoneyWithdrawal` tests in this file and use whichever signer they use. Do not grant new roles to make the test pass.

- [ ] **Step 2: Run to verify failure**

```bash
yarn hardhat test test/GasMoneyManager.test.js
```

Expected: the two overflow tests FAIL (no `GasMoneyValueOutOfRange` error exists yet); the bounds test FAILS (current `uint256` storage accepts `2**128-1` but the test asserts against the new error surface only after Step 3 — it may pass now, which is fine).

- [ ] **Step 3: Replace the storage struct**

`contracts/account/GasMoneyManager.sol`, replacing lines 24-29:

```solidity
    /**
     * @notice Per-account withdrawal accounting, packed into a single slot.
     */
    struct GasMoneyWithdrawal {
        uint128 amount; // wei withdrawn in the current period
        uint64 periodStart; // unix timestamp of the current period start
    }

    /// @custom:storage-location erc7201:traveltoken.messenger.storage.GasMoneyManager
    struct GasMoneyStorage {
        mapping(address => GasMoneyWithdrawal) _withdrawals;
        uint128 _withdrawalLimit;
        uint64 _withdrawalPeriod;
    }
```

Do **not** touch `GasMoneyStorageLocation` on line 32 — the namespace string is unchanged.

- [ ] **Step 4: Add the range error and a checked cast helper**

After the existing errors (around line 66), add:

```solidity
    error GasMoneyValueOutOfRange(uint256 limit, uint256 period);
```

And add these private helpers near the bottom of the contract:

```solidity
    function _toUint128(uint256 value, uint256 limit, uint256 period) private pure returns (uint128) {
        if (value > type(uint128).max) {
            revert GasMoneyValueOutOfRange(limit, period);
        }
        return uint128(value);
    }

    function _toUint64(uint256 value, uint256 limit, uint256 period) private pure returns (uint64) {
        if (value > type(uint64).max) {
            revert GasMoneyValueOutOfRange(limit, period);
        }
        return uint64(value);
    }
```

- [ ] **Step 5: Update the initializer (signature unchanged)**

Replace lines 72-76:

```solidity
    function __GasMoneyManager_init(uint256 withdrawalLimit, uint256 withdrawalPeriod) internal onlyInitializing {
        GasMoneyStorage storage $ = _getGasMoneyStorage();
        $._withdrawalLimit = _toUint128(withdrawalLimit, withdrawalLimit, withdrawalPeriod);
        $._withdrawalPeriod = _toUint64(withdrawalPeriod, withdrawalLimit, withdrawalPeriod);
    }
```

- [ ] **Step 6: Update the withdrawal logic**

Replace the body of `_withdrawGasMoney` (lines 89-119):

```solidity
    function _withdrawGasMoney(uint256 amount) internal {
        GasMoneyStorage storage $ = _getGasMoneyStorage();

        uint256 limit = $._withdrawalLimit;

        // Ensure the withdrawal does not exceed the allowed limit
        if (amount > limit) {
            revert WithdrawalLimitExceeded(limit, amount);
        }

        GasMoneyWithdrawal memory withdrawal = $._withdrawals[msg.sender];
        uint256 currentTime = block.timestamp;

        // Reset the withdrawn amount if a new period has started. If more time than
        // the withdrawal period has passed, it is allowed to withdraw the full amount.
        if (currentTime > uint256(withdrawal.periodStart) + $._withdrawalPeriod) {
            withdrawal.amount = 0;
            withdrawal.periodStart = _toUint64(currentTime, limit, $._withdrawalPeriod);
        }

        // Ensure the withdrawal does not exceed the allowed limit for the period
        if (uint256(withdrawal.amount) + amount > limit) {
            revert WithdrawalLimitExceededForPeriod(limit, amount);
        }

        // Update the withdrawn amount. Safe: the sum was just checked against
        // limit, which is itself a uint128.
        withdrawal.amount = uint128(uint256(withdrawal.amount) + amount);
        $._withdrawals[msg.sender] = withdrawal;

        // Transfer the gas money
        payable(msg.sender).sendValue(amount);

        emit GasMoneyWithdrawal(msg.sender, amount);
    }
```

- [ ] **Step 7: Update the setter and both getters (signatures unchanged)**

Replace lines 127-133:

```solidity
    function _setGasMoneyWithdrawal(uint256 limit, uint256 period) internal {
        GasMoneyStorage storage $ = _getGasMoneyStorage();
        $._withdrawalLimit = _toUint128(limit, limit, period);
        $._withdrawalPeriod = _toUint64(period, limit, period);

        emit GasMoneyWithdrawalUpdated(limit, period);
    }
```

Replace lines 141-144:

```solidity
    function getGasMoneyWithdrawal() public view returns (uint256 withdrawalLimit, uint256 withdrawalPeriod) {
        GasMoneyStorage storage $ = _getGasMoneyStorage();
        return (uint256($._withdrawalLimit), uint256($._withdrawalPeriod));
    }
```

Replace lines 153-158:

```solidity
    function getGasMoneyWithdrawalForAccount(
        address account
    ) public view returns (uint256 periodStart, uint256 withdrawnAmount) {
        GasMoneyStorage storage $ = _getGasMoneyStorage();
        GasMoneyWithdrawal memory withdrawal = $._withdrawals[account];
        return (uint256(withdrawal.periodStart), uint256(withdrawal.amount));
    }
```

- [ ] **Step 8: Run the tests**

```bash
yarn compile && yarn hardhat test test/GasMoneyManager.test.js
```

Expected: all pass, including the four new ones.

- [ ] **Step 9: Run the full suite**

```bash
yarn test
```

Expected: `124 passing`.

- [ ] **Step 10: Regenerate and commit**

```bash
yarn hardhat export-abi
git add contracts/ test/ abi/ docs/
git commit -m "perf: pack GasMoneyStorage into fewer slots

Limit fits uint128 (>> total ETH supply), period fits uint64. Merging the
paired per-account mappings into one struct removes a cold SLOAD from
every bot gas withdrawal, the highest-frequency call in the system.

External signatures are unchanged: getters widen on read, writes use a
checked cast that reverts (GasMoneyValueOutOfRange) rather than
truncating -- a silently truncated limit would be a security bug."
```

---

## Task 7: Fix initializer wiring

`TTMAccountManager` uses `nonReentrant` but never calls `__ReentrancyGuard_init()`. Two `__init` functions are declared but never invoked — harmless today because they are empty, which is exactly the trap: the day someone adds a default it silently will not apply.

**Files:**
- Modify: `contracts/manager/TTMAccountManager.sol:189-203`
- Modify: `contracts/account/TTMAccount.sol:213-237`

**Interfaces:**
- Consumes: nothing
- Produces: correctly initialized base contracts

- [ ] **Step 1: Confirm the gap**

```bash
grep -n "__ReentrancyGuard_init\|nonReentrant" contracts/manager/TTMAccountManager.sol
grep -rn "__PartnerConfiguration_init\|__ServiceRegistry_init" contracts/ --include='*.sol'
```

Expected: `TTMAccountManager` uses `nonReentrant` (line ~243) with no `__ReentrancyGuard_init` call; `__PartnerConfiguration_init` and `__ServiceRegistry_init` appear only as declarations.

- [ ] **Step 2: Fix the manager's initializer**

In `contracts/manager/TTMAccountManager.sol`, the `initialize` body becomes:

```solidity
        __Pausable_init();
        __AccessControl_init();
        __UUPSUpgradeable_init();
        __ReentrancyGuard_init();
        __ServiceRegistry_init();
```

- [ ] **Step 3: Wire the account's config initializer**

In `contracts/account/TTMAccount.sol`, `initialize` currently calls `__AccessControl_init()`, `__UUPSUpgradeable_init()`, `__ReentrancyGuard_init()`. Add `__PartnerConfiguration_init()` immediately after `__ReentrancyGuard_init()`:

```solidity
        __AccessControl_init();
        __UUPSUpgradeable_init();
        __ReentrancyGuard_init();
        __PartnerConfiguration_init();
```

Leave the existing `__GasMoneyManager_init(withdrawalLimit, withdrawalPeriod)` call at the end of the function where it is.

- [ ] **Step 4: Compile and test**

```bash
yarn compile && yarn test
```

Expected: `124 passing`.

- [ ] **Step 5: Commit**

```bash
git add contracts/
git commit -m "fix: initialize ReentrancyGuard and the unchained base contracts

TTMAccountManager used nonReentrant without ever calling
__ReentrancyGuard_init, paying a cold-slot SSTORE on first use.

__PartnerConfiguration_init and __ServiceRegistry_init were declared but
never called. Empty today -- which is the trap, since the first default
added to either would silently not apply."
```

---

## Task 8: Zero-address validation in initializers

`AccessControl._grantRole(role, address(0))` silently succeeds and is permanently unusable — a deploy-time footgun, particularly given the Ignition module wires several roles from parameters.

**Files:**
- Modify: `contracts/manager/TTMAccountManager.sol` (initialize)
- Modify: `contracts/account/TTMAccount.sol` (initialize)
- Modify: `contracts/booking-token/BookingToken.sol` (initialize)
- Test: `test/TTMAccountManager.test.js`, `test/TTMAccount.test.js`, `test/BookingToken.test.js`

**Interfaces:**
- Consumes: Task 7's initializers
- Produces: `error ZeroAddress()` on all three contracts

- [ ] **Step 1: Write the failing test for the manager**

Add inside the top-level `describe("TTMAccountManager")` block in `test/TTMAccountManager.test.js`:

```javascript
    describe("Initializer validation", function () {
        it("should reject a zero address for any role parameter", async function () {
            const signers = await setupSigners();
            const Manager = await ethers.getContractFactory("TTMAccountManager");
            const zero = ethers.ZeroAddress;
            const ok = signers.managerAdmin.address;

            for (const args of [
                [zero, ok, ok, ok],
                [ok, zero, ok, ok],
                [ok, ok, zero, ok],
                [ok, ok, ok, zero],
            ]) {
                await expect(upgrades.deployProxy(Manager, args, { kind: "uups" })).to.be.revertedWithCustomError(
                    Manager,
                    "ZeroAddress",
                );
            }
        });
    });
```

> Ensure `upgrades` is imported at the top of the file: `const { ethers, upgrades } = require("hardhat");`. If the existing file uses a different deployment helper for the manager proxy, mirror whatever `deployTTMAccountManagerFixture` in `test/utils/fixtures.js` does.

- [ ] **Step 2: Run to verify failure**

```bash
yarn hardhat test test/TTMAccountManager.test.js
```

Expected: FAIL — no `ZeroAddress` error exists.

- [ ] **Step 3: Add the error and checks to the manager**

In `contracts/manager/TTMAccountManager.sol`, add near the other errors:

```solidity
    /**
     * @notice A required address parameter was the zero address.
     */
    error ZeroAddress();
```

And at the top of `initialize`, before any `_grantRole`:

```solidity
        if (
            defaultAdmin == address(0) || pauser == address(0) || upgrader == address(0) || versioner == address(0)
        ) {
            revert ZeroAddress();
        }
```

- [ ] **Step 4: Add the same to `TTMAccount`**

In `contracts/account/TTMAccount.sol`, add the `error ZeroAddress();` declaration alongside the other errors, and at the top of `initialize`:

```solidity
        if (
            manager == address(0) ||
            bookingToken == address(0) ||
            defaultAdmin == address(0) ||
            upgrader == address(0)
        ) {
            revert ZeroAddress();
        }
```

- [ ] **Step 5: Add the same to `BookingToken`**

In `contracts/booking-token/BookingToken.sol`, add `error ZeroAddress();` alongside the other errors, and at the top of `initialize`:

```solidity
        if (manager == address(0) || defaultAdmin == address(0) || upgrader == address(0)) {
            revert ZeroAddress();
        }
```

- [ ] **Step 6: Add matching tests for the other two contracts**

In `test/TTMAccount.test.js`, inside the top-level describe:

```javascript
    describe("Initializer validation", function () {
        it("should reject a zero address for any constructor parameter", async function () {
            const signers = await setupSigners();
            const { ttmAccountManager } = await loadFixture(deployTTMAccountManagerFixture);
            const zero = ethers.ZeroAddress;
            const ok = signers.ttmAccountAdmin.address;
            const mgr = await ttmAccountManager.getAddress();

            const Account = await ethers.getContractFactory("TTMAccount", {
                libraries: { BookingTokenOperator: await (await ethers.deployContract("BookingTokenOperator")).getAddress() },
            });

            for (const args of [
                [zero, ok, ok, ok],
                [mgr, zero, ok, ok],
                [mgr, ok, zero, ok],
                [mgr, ok, ok, zero],
            ]) {
                await expect(
                    upgrades.deployProxy(Account, args, { kind: "uups", unsafeAllow: ["external-library-linking"] }),
                ).to.be.revertedWithCustomError(Account, "ZeroAddress");
            }
        });
    });
```

> `TTMAccount` links the `BookingTokenOperator` library. If `test/utils/fixtures.js` already has a helper that deploys and links it, use that instead of the inline `deployContract` above.

In `test/BookingToken.test.js`, inside the top-level describe:

```javascript
    describe("Initializer validation", function () {
        it("should reject a zero address for any constructor parameter", async function () {
            const signers = await setupSigners();
            const BookingToken = await ethers.getContractFactory("BookingToken");
            const zero = ethers.ZeroAddress;
            const ok = signers.btAdmin.address;

            for (const args of [
                [zero, ok, ok],
                [ok, zero, ok],
                [ok, ok, zero],
            ]) {
                await expect(
                    upgrades.deployProxy(BookingToken, args, { kind: "uups" }),
                ).to.be.revertedWithCustomError(BookingToken, "ZeroAddress");
            }
        });
    });
```

- [ ] **Step 7: Run the full suite**

```bash
yarn compile && yarn test
```

Expected: `127 passing`.

- [ ] **Step 8: Regenerate and commit**

```bash
yarn hardhat export-abi
git add contracts/ test/ abi/ docs/
git commit -m "fix: reject zero addresses in all three initializers

AccessControl._grantRole(role, address(0)) silently succeeds and is
permanently unusable. A parameter typo in the Ignition file would have
bricked a role with no error at deploy time."
```

---

## Task 9: Pausable on BookingToken

`BookingToken` holds the system's value flow but has no pause lever; the only response to a discovered bug in the payment path is currently a full upgrade. Pausability today exists only on the manager and covers only account creation.

**Files:**
- Modify: `contracts/booking-token/BookingToken.sol` (inheritance, roles, initialize, three function modifiers, new pause/unpause)
- Test: `test/BookingToken.test.js`

**Interfaces:**
- Consumes: Task 8's `ZeroAddress` validation
- Produces: `PAUSER_ROLE`, `pause()`, `unpause()`, `paused()`

**ABI impact: additive only.** `whenNotPaused` changes no signature and behaviour is identical while unpaused.

- [ ] **Step 1: Write the failing tests**

Add inside the top-level `describe("BookingToken")` block in `test/BookingToken.test.js`:

```javascript
    describe("Pausable", function () {
        it("should not let a non-pauser pause", async function () {
            const { bookingToken, signers } = await loadFixture(deployBookingTokenFixture);
            await expect(bookingToken.connect(signers.otherAccount1).pause()).to.be.revertedWithCustomError(
                bookingToken,
                "AccessControlUnauthorizedAccount",
            );
        });

        it("should let the pauser pause and unpause", async function () {
            const { bookingToken, signers } = await loadFixture(deployBookingTokenFixture);
            await bookingToken.connect(signers.btAdmin).grantRole(
                await bookingToken.PAUSER_ROLE(),
                signers.btAdmin.address,
            );

            await bookingToken.connect(signers.btAdmin).pause();
            expect(await bookingToken.paused()).to.be.true;

            await bookingToken.connect(signers.btAdmin).unpause();
            expect(await bookingToken.paused()).to.be.false;
        });

        it("should block minting while paused", async function () {
            const { bookingToken, ttmAccount, signers } = await loadFixture(
                deployAndConfigureAllWithRegisteredServicesFixture,
            );
            await bookingToken.connect(signers.btAdmin).grantRole(
                await bookingToken.PAUSER_ROLE(),
                signers.btAdmin.address,
            );
            await bookingToken.connect(signers.btAdmin).pause();

            await expect(
                ttmAccount.connect(signers.chequeOperator).mintBookingToken(
                    await ttmAccount.getAddress(),
                    "https://example.com/token",
                    (await time.latest()) + 3600,
                    100n,
                    ethers.ZeroAddress,
                ),
            ).to.be.revertedWithCustomError(bookingToken, "EnforcedPause");
        });
    });
```

> **Two things to check against the real code before running this.**
> 1. The `mintBookingToken` argument list must match the actual signature —
>    read an existing successful mint call in `test/BookingToken.test.js` and
>    copy its argument shape exactly.
> 2. `signers.chequeOperator` is the signer the existing tests use as the bot
>    address (confirmed in `test/utils/fixtures.js`). **Task 13 renames it to
>    `botOperator`**, so if Task 13 has already run, use `signers.botOperator`
>    here instead.
>
> Import `time` from `@nomicfoundation/hardhat-toolbox/network-helpers` if it is
> not already imported in this file.

- [ ] **Step 2: Run to verify failure**

```bash
yarn hardhat test test/BookingToken.test.js
```

Expected: FAIL — `pause` is not a function.

- [ ] **Step 3: Add the import and inheritance**

In `contracts/booking-token/BookingToken.sol`, add the import alongside the other upgradeable imports:

```solidity
import { PausableUpgradeable } from "@openzeppelin/contracts-upgradeable/utils/PausableUpgradeable.sol";
```

Add `PausableUpgradeable` to the inheritance list (after `AccessControlUpgradeable`):

```solidity
contract BookingToken is
    Initializable,
    ERC721Upgradeable,
    ERC721EnumerableUpgradeable,
    ERC721URIStorageUpgradeable,
    AccessControlUpgradeable,
    PausableUpgradeable,
    ReentrancyGuardUpgradeable,
    UUPSUpgradeable,
    BookingTokenCancellable
{
```

- [ ] **Step 4: Add the role**

Alongside `UPGRADER_ROLE` and `MIN_EXPIRATION_ADMIN_ROLE`:

```solidity
    /**
     * @notice Pauser role can pause the contract, halting minting, buying, and
     * cancellation finalization.
     */
    bytes32 public constant PAUSER_ROLE = keccak256("PAUSER_ROLE");
```

- [ ] **Step 5: Initialize Pausable**

In `initialize`, add `__Pausable_init();` after `__AccessControl_init();`.

- [ ] **Step 6: Add pause and unpause**

Add near `_authorizeUpgrade`:

```solidity
    /**
     * @notice Pauses minting, buying, and cancellation finalization.
     */
    function pause() external onlyRole(PAUSER_ROLE) {
        _pause();
    }

    /**
     * @notice Resumes normal operation.
     */
    function unpause() external onlyRole(PAUSER_ROLE) {
        _unpause();
    }
```

- [ ] **Step 7: Gate the three value-flow functions**

Add `whenNotPaused` to each. Preserve the existing modifier order and append:

- `safeMintWithReservation` (~line 351)
- `buyReservedToken` (~line 449) — becomes `public payable virtual nonReentrant whenNotPaused onlyTTMAccount(msg.sender)`
- `finalizeCancellation` (~line 824)

- [ ] **Step 8: Run the tests**

```bash
yarn compile && yarn hardhat test test/BookingToken.test.js
```

Expected: all pass.

- [ ] **Step 9: Check the size budget**

```bash
yarn compile 2>&1 | grep -A3 "BookingToken"
```

Expected: `BookingToken` under 22.5 KiB. If it exceeds 24.576 KiB the build is broken — drop `optimizer.runs` to 500 in `hardhat.config.js` and recompile.

- [ ] **Step 10: Run the full suite and commit**

Expected: `130 passing`.

```bash
yarn test && yarn hardhat export-abi
git add contracts/ test/ abi/ docs/
git commit -m "feat: add Pausable to BookingToken

BookingToken holds the system's value flow but had no pause lever; the
only response to a payment-path bug was a full upgrade. Pausability
previously existed only on the manager, covering account creation.

Gates safeMintWithReservation, buyReservedToken, finalizeCancellation.
Additive: whenNotPaused changes no signature and behaviour is identical
while unpaused."
```

---

## Task 10: Emit the Deposit event

`event Deposit` is declared but never emitted, so every ETH transfer into an account — including the `msg.value` the manager forwards at creation — is invisible. The UI already has a tested renderer for it.

**Files:**
- Modify: `contracts/account/TTMAccount.sol:239`
- Test: `test/TTMAccount.test.js`

**Interfaces:**
- Consumes: nothing
- Produces: `Deposit(address indexed sender, uint256 amount)` actually emitted

- [ ] **Step 1: Write the failing tests**

Add inside the top-level `describe("TTMAccount")` block in `test/TTMAccount.test.js`:

```javascript
    describe("Deposit event", function () {
        it("should emit Deposit when receiving ETH", async function () {
            const { ttmAccount, signers } = await loadFixture(deployAndConfigureAllFixture);
            const amount = ethers.parseEther("1.5");

            await expect(
                signers.otherAccount1.sendTransaction({ to: await ttmAccount.getAddress(), value: amount }),
            )
                .to.emit(ttmAccount, "Deposit")
                .withArgs(signers.otherAccount1.address, amount);
        });
    });
```

- [ ] **Step 2: Run to verify failure**

```bash
yarn hardhat test test/TTMAccount.test.js
```

Expected: FAIL — the event is never emitted.

- [ ] **Step 3: Emit it**

`contracts/account/TTMAccount.sol:239`:

```solidity
    receive() external payable {
        emit Deposit(msg.sender, msg.value);
    }
```

- [ ] **Step 4: Run the tests**

```bash
yarn compile && yarn hardhat test test/TTMAccount.test.js
```

Expected: PASS.

- [ ] **Step 5: Run the full suite and commit**

Expected: `131 passing`.

```bash
yarn test
git add contracts/ test/
git commit -m "fix: emit Deposit from receive()

The event was declared but never emitted, making every incoming ETH
transfer invisible to the activity feed -- including the msg.value the
manager forwards at account creation.

ui/src/lib/activity/catalog.ts:174 already renders this event and reads
sender/amount, so no UI change is needed."
```

---

## Task 11: Fix the token symbol and remove reinitializeV2

`BToken` is the intended symbol; it was previously applied post-deploy by a migration script. With it set correctly at initialization, `reinitializeV2` has no purpose — it burns reinitializer version 2 and is a standing admin backdoor to rewrite the token's identity.

**Files:**
- Modify: `contracts/booking-token/BookingToken.sol:299`, delete `:314-330` (the REINITIALIZE section)
- Modify: `test/BookingToken.test.js:45`, delete `:52,56,64` assertions
- Modify: `ui/src/pages/tabs/BookingTokenTab.test.tsx:11,49`

**Interfaces:**
- Consumes: nothing
- Produces: `symbol()` returns `BToken`; `reinitializeV2` removed from the ABI

**This is the only subtractive ABI change in the plan.** Safe: the sole non-test caller was `scripts/reinitialize_booking_token.js`, deleted in Task 4.

- [ ] **Step 1: Update the symbol assertion test**

`test/BookingToken.test.js:45` — change:

```javascript
const currentSymbol = expect(await bookingToken.symbol()).to.be.equal("TRIP");
```
to
```javascript
expect(await bookingToken.symbol()).to.be.equal("BToken");
```

(The unused `currentSymbol` binding is dropped — `expect` returns an assertion, not a symbol.)

- [ ] **Step 2: Run to verify failure**

```bash
yarn hardhat test test/BookingToken.test.js
```

Expected: FAIL — `expected 'TRIP' to equal 'BToken'`.

- [ ] **Step 3: Set the symbol at initialization**

`contracts/booking-token/BookingToken.sol:299`:

```solidity
        __ERC721_init("BookingToken", "BToken");
```

- [ ] **Step 4: Delete the reinitializer**

Remove the entire `REINITIALIZE` section from `contracts/booking-token/BookingToken.sol` — the banner comment, the NatSpec block, and:

```solidity
    function reinitializeV2(
        string memory newName,
        string memory newSymbol
    ) public reinitializer(2) onlyRole(DEFAULT_ADMIN_ROLE) {
        __ERC721_init(newName, newSymbol);
    }
```

- [ ] **Step 5: Delete the three reinitializeV2 tests**

In `test/BookingToken.test.js`, remove the assertions at (originally) lines 52, 56, and 64 — the non-admin revert test, the happy-path rename, and the double-reinitialize revert. Delete the enclosing `it(...)` blocks entirely, and the `newName`/`newSymbol` constants if they become unused.

- [ ] **Step 6: Update the UI test**

`ui/src/pages/tabs/BookingTokenTab.test.tsx` — change `symbol: "TRIP"` (line 11) to `symbol: "BToken"` and `screen.getByText("TRIP")` (line 49) to `screen.getByText("BToken")`.

Do **not** run the UI test suite (it needs its own install, and disk is tight). The change is a literal string swap.

- [ ] **Step 7: Run the contract suite**

```bash
yarn compile && yarn test
```

Expected: `128 passing` (three fewer than after Task 10).

- [ ] **Step 8: Confirm reinitializeV2 is gone from the ABI**

```bash
yarn hardhat export-abi
grep -c "reinitializeV2" abi/contracts/booking-token/BookingToken.sol/BookingToken.json || echo "absent (expected)"
```

Expected: `absent (expected)`.

- [ ] **Step 9: Commit**

```bash
git add contracts/ test/ ui/ abi/ docs/
git commit -m "fix!: set BToken symbol at init, drop reinitializeV2

BToken is the intended symbol; it was previously applied post-deploy by
a migration script. With it correct at initialization, reinitializeV2
has no purpose -- it burned reinitializer version 2 and was a standing
admin backdoor to rewrite the token's identity.

Only subtractive ABI change in this branch. Its sole non-test caller was
scripts/reinitialize_booking_token.js, already deleted."
```

---

## Task 12: Capability removal must not silently succeed

`_removeServiceCapability` breaks on match but falls through and returns successfully when the capability is absent, while `TTMAccount` emits `ServiceCapabilityRemoved` unconditionally. Indexers record removals that never happened.

**Files:**
- Modify: `contracts/partner/PartnerConfiguration.sol:193-207`
- Test: `test/PartnerConfiguration.test.js`

**Interfaces:**
- Consumes: nothing
- Produces: `error CapabilityDoesNotExist(bytes32 serviceHash, string capability)`

- [ ] **Step 1: Write the failing test**

Add inside the top-level `describe("PartnerConfiguration")` block in `test/PartnerConfiguration.test.js`:

```javascript
    describe("Capability removal", function () {
        it("should revert when removing a capability that does not exist", async function () {
            const { ttmAccount, signers, serviceName } = await loadFixture(
                deployAndConfigureAllWithRegisteredServicesFixture,
            );

            await expect(
                ttmAccount.connect(signers.ttmServiceAdmin).removeServiceCapability(serviceName, "no-such-capability"),
            ).to.be.revertedWithCustomError(ttmAccount, "CapabilityDoesNotExist");
        });
    });
```

> Use whatever service name and admin signer the neighbouring capability tests in this file already use. If the fixture does not expose `serviceName`, read an existing `addServiceCapability` test and copy its setup.

- [ ] **Step 2: Run to verify failure**

```bash
yarn hardhat test test/PartnerConfiguration.test.js
```

Expected: FAIL — the call currently succeeds silently.

- [ ] **Step 3: Add the error**

In `contracts/partner/PartnerConfiguration.sol`, alongside the other errors:

```solidity
    error CapabilityDoesNotExist(bytes32 serviceHash, string capability);
```

- [ ] **Step 4: Revert when not found**

Replace the loop in `_removeServiceCapability`:

```solidity
        string[] storage capabilities = $._supportedServices[serviceHash]._capabilities;
        uint256 length = capabilities.length;
        for (uint256 i = 0; i < length; i++) {
            if (keccak256(abi.encodePacked(capabilities[i])) == keccak256(abi.encodePacked(capability))) {
                capabilities[i] = capabilities[length - 1];
                capabilities.pop();
                return;
            }
        }
        revert CapabilityDoesNotExist(serviceHash, capability);
```

This matches the revert-on-absent convention used by every other remove in this file.

- [ ] **Step 5: Run the tests**

```bash
yarn compile && yarn hardhat test test/PartnerConfiguration.test.js
```

Expected: PASS.

- [ ] **Step 6: Run the full suite and commit**

Expected: `129 passing`.

```bash
yarn test && yarn hardhat export-abi
git add contracts/ test/ abi/ docs/
git commit -m "fix: revert when removing a non-existent service capability

The loop fell through and returned successfully when the capability was
absent, while TTMAccount emitted ServiceCapabilityRemoved regardless --
so indexers recorded removals that never happened.

Matches the revert-on-absent convention of every other remove here."
```

---

## Task 13: Correct stale comments and NatSpec

Every item here describes behaviour the code does not have. Two of them have already caused reviewers to reach wrong conclusions.

**Files:**
- Modify: `contracts/manager/TTMAccountManager.sol:32-42`, `:232-237`
- Modify: `contracts/booking-token/BookingTokenCancellable.sol:42`
- Modify: `contracts/booking-token/BookingToken.sol:587,590`
- Modify: `contracts/account/TTMAccount.sol:34-37`, `:212`; `contracts/account/ITTMAccount.sol:5`
- Modify: `README.md:59-62`, `CLAUDE.md`
- Modify: `tasks/manager.js:52,265`, `tasks/account.js:53,1111`
- Modify: `test/utils/fixtures.js`, `test/GasMoneyManager.test.js`

**Interfaces:**
- Consumes: nothing
- Produces: no code change except a named constant in `BookingToken`

- [ ] **Step 1: Fix the storage-slot derivation comment**

`contracts/booking-token/BookingTokenCancellable.sol:42`. The comment claims the preimage is `...BookingTokenCancellableV2`, but the constant is the hash of `...BookingTokenCancellable` (no `V2`) — verified by recomputation. **The value is correct; only the comment is wrong.**

```solidity
    // keccak256(abi.encode(uint256(keccak256("traveltoken.messenger.storage.BookingTokenCancellable")) - 1)) & ~bytes32(uint256(0xff));
```

**Do not touch the hex constant on the next line.** Anyone "fixing" the constant to match the old comment would silently relocate the proposals mapping.

- [ ] **Step 2: Verify the constant still matches after the edit**

```bash
node -e '
const {keccak256, toUtf8Bytes, AbiCoder, toBeHex} = require("ethers");
const ns = "traveltoken.messenger.storage.BookingTokenCancellable";
const inner = BigInt(keccak256(toUtf8Bytes(ns))) - 1n;
const enc = AbiCoder.defaultAbiCoder().encode(["uint256"],[inner]);
console.log(toBeHex((BigInt(keccak256(enc)) & ~0xffn),32));
'
```

Expected: `0x1973af827d98d3fe49ea0c9179b69c8f267978d372df89173953530663b72200`, matching line 44.

- [ ] **Step 3: Replace the FIXMEs with a named constant**

`contracts/booking-token/BookingToken.sol:587,590`. The `FIXME` comments are stale — the reason **is** defined in the protocol:

```proto
// AUTOMATIC REJECTION REASON DURING TRANSFER OF THE TOKEN ON-CHAIN
REJECTION_REASON_TRANSFER_ON_CHAIN = 99;
```

Delete both `// FIXME: Define a reason in the Travel Token Messenger Protocol and update this` comments. Add near the other constants:

```solidity
    /**
     * @notice Protocol rejection reason emitted when a pending cancellation is
     * automatically resolved because the token was transferred on-chain.
     * Mirrors REJECTION_REASON_TRANSFER_ON_CHAIN in the messenger protocol.
     */
    uint16 private constant REJECTION_REASON_TRANSFER_ON_CHAIN = 99;
    uint16 private constant REJECTION_REASON_VERSION = 1;
```

Replace the two magic-number call sites with `REJECTION_REASON_TRANSFER_ON_CHAIN, REJECTION_REASON_VERSION`. **The emitted values are unchanged.**

- [ ] **Step 4: Rewrite the manager's contract header**

`contracts/manager/TTMAccountManager.sol:32-42` — delete the sentence about approving "the service fee token with the amount of prefund" and the entire "Developer Fee" paragraph. Neither prefund, developer wallet, fee basis points, nor cheques exist. The manager is now factory + account registry + service registry.

- [ ] **Step 5: Fix the createTTMAccount NatSpec**

`contracts/manager/TTMAccountManager.sol:232-237` — delete the KYC/KYB claim (that was chain-enforced on Camino and does not exist on Base) and the "Caller must approve the pre-fund amount" sentence. Replace with an explicit statement of current behaviour:

```
     * @notice Creates a new TTMAccount.
     *
     * This function is currently permissionless: any address may create an
     * account. See docs/decisions/2026-07-21-contract-design-decisions.md
     * (Decision 1) -- gating must be resolved before Base mainnet.
```

- [ ] **Step 6: Fix the gas-money limit NatSpec**

`contracts/account/TTMAccount.sol:34-37` currently reads as though the limit is account-wide. It is **per bot address** — each bot gets its own allowance. Reword to say so explicitly.

- [ ] **Step 7: Delete the prefund tombstones**

Remove `// `uint256 prefundAmount` is removed as it is no longer used in the contract @2025-08-28` from `contracts/account/TTMAccount.sol:212` and `contracts/account/ITTMAccount.sol:5`.

- [ ] **Step 8: Fix the tombstone comments in tasks**

Delete the four "removed as service fees are deprecated" / "Obsolete developer fee and prefund set tasks removed" comments at `tasks/manager.js:52`, `tasks/manager.js:265`, `tasks/account.js:53`, `tasks/account.js:1111`.

- [ ] **Step 9: Fix the test comments**

- `test/utils/fixtures.js:175` — `// Deposit CAM` → `// Deposit ETH`
- `test/utils/fixtures.js:186`, `:242` — `// Deposit service fee token` → `// Deposit NullUSD`
- `test/utils/fixtures.js:16-18`, `:38-40` — delete the unused `developerWallet`, `developerWalletAdmin`, `feeAdmin` signers from **both** the `getSigners()` destructuring array and the `signers` object literal.

  > These are positional in the `await ethers.getSigners()` destructure, so
  > removing them shifts every signer after them onto a different account.
  > Nothing depends on those indices, but **run the full suite immediately
  > after this edit** — a balance-sensitive test would fail here and nowhere
  > else.

- Rename the signer `chequeOperator` → `botOperator` everywhere. Find all sites first rather than trusting line numbers, since earlier tasks added new usages:

  ```bash
  grep -rn "chequeOperator" test/
  ```

  Expected sites include `test/utils/fixtures.js` (destructure + object literal), `test/TTMAccount.test.js:137`, and any test added in Task 9.
- `test/GasMoneyManager.test.js` — replace the 14 `CAM` comments with `ETH`, and delete the 5 `// Add more funds ... so we are not under the prefund spent` comments (the check no longer exists)

- [ ] **Step 10: Fix README and CLAUDE**

- `README.md:59-62` — remove "developer fee" from the `TTMAccountManager` description
- `CLAUDE.md` — remove the same claim; correct "`yarn compile` also runs `contract-sizer`" to mention docgen as well

- [ ] **Step 11: Compile, test, commit**

```bash
yarn compile && yarn test
yarn hardhat export-abi
git add -A
git commit -m "docs: correct comments describing removed behaviour

Every item here described behaviour the code does not have, and two had
already misled reviewers:

- BookingTokenCancellable slot comment claimed a 'V2' preimage; the
  constant is the non-V2 hash. Value correct, comment wrong -- anyone
  'fixing' the constant would relocate the proposals mapping.
- The BookingToken FIXMEs made correct code look unfinished. Reason 99
  is REJECTION_REASON_TRANSFER_ON_CHAIN in the protocol; now a named
  constant. Emitted values unchanged.

Also drops the manager's fee/prefund/cheque header, the Camino KYC claim
on createTTMAccount, prefund tombstones, and CAM units in test comments."
```

---

## Task 14: ETH denominations in the account CLI

Base's native coin is ETH, but `tasks/account.js` uses `CAM`/`nCAM`/`aCAM` units and a user-facing `--camAmount` flag. This is a breaking CLI change, which is acceptable — no external automation depends on these task names.

**Files:**
- Modify: `tasks/account.js:216`, `:226-228`, `:267`, `:274-304`, `:291`, `:295`, `:304`, `:329`, `:483`, `:928`, `:937`, `:957`, `:966`

**Interfaces:**
- Consumes: nothing
- Produces: `--amount` flag; `eth`/`gwei`/`wei` units

- [ ] **Step 1: Find every site**

```bash
grep -n "camAmount\|cam-amount\|aCAM\|nCAM\|CAM" tasks/account.js
```

- [ ] **Step 2: Rename the parameter**

Change the `--camAmount` task parameter to `--amount` (Hardhat exposes `taskArgs.amount`), and update every `taskArgs.camAmount` reference.

- [ ] **Step 3: Rename the units**

Replace the unit strings: `aCAM` → `wei`, `nCAM` → `gwei`, `CAM` → `eth`. Change the `--unit` default from `"aCAM"` to `"wei"`. Update all `console.log` output that prints `CAM` to print `ETH`.

Ensure the unit values still map correctly onto `ethers.parseUnits` / `ethers.formatUnits` — `wei`, `gwei`, and `ether` are the ethers-native names, so if the code passes the unit string straight through, use `"ether"` rather than `"eth"` at the call site.

- [ ] **Step 4: Verify the tasks still load and register**

```bash
yarn hardhat --help 2>&1 | grep -i "account"
yarn hardhat account --help 2>&1 | head -30
```

Expected: tasks list without error; no `camAmount` remains.

- [ ] **Step 5: Confirm no stragglers**

```bash
grep -n "camAmount\|aCAM\|nCAM\|CAM" tasks/account.js || echo "clean"
```

Expected: `clean`.

- [ ] **Step 6: Run the suite and commit**

```bash
yarn test
git add tasks/account.js
git commit -m "refactor(tasks)!: use ETH denominations instead of CAM

Base's native coin is ETH. Renames the --camAmount flag to --amount and
the CAM/nCAM/aCAM units to eth/gwei/wei.

Breaking CLI change; no external automation depends on these names."
```

---

## Task 15: Fix the deploy-path task and Ignition bugs

`services:register` is in the deploy critical path and crashes on a relative `--json` path. The error handler can `TypeError` and abort the 63-service loop mid-way. The Ignition admin parameter is commented out.

**Files:**
- Modify: `tasks/manager.js:78`, `:105`
- Modify: `ignition/modules/messenger.js:12-13`

**Interfaces:**
- Consumes: nothing
- Produces: `services:register` accepts relative paths; `managerAdmin` is overridable

- [ ] **Step 1: Reproduce the path bug**

```bash
yarn hardhat manager services:register --json ./services/00_initial.json --network hardhat 2>&1 | head -5
```

Expected: `MODULE_NOT_FOUND` — `require()` resolves relative paths against `tasks/`, not cwd.

- [ ] **Step 2: Fix path resolution**

`tasks/manager.js:105`. Add `const path = require("path");` at the top of the file if absent, then replace:

```javascript
const services = require(taskArgs.json);
```
with
```javascript
const services = require(path.resolve(process.cwd(), taskArgs.json));
```

- [ ] **Step 3: Fix the error handler**

`tasks/manager.js:78` — replace `if (error.data.data && contract)` with:

```javascript
if (error.data?.data && contract) {
```

Without the optional chain, ethers v6 error shapes lacking `.data` throw a `TypeError` that masks the real revert reason and aborts the registration loop partway through.

- [ ] **Step 4: Verify the path fix**

```bash
yarn hardhat manager services:register --json ./services/00_initial.json --network hardhat 2>&1 | head -5
```

Expected: no `MODULE_NOT_FOUND`. It may still fail for lack of a deployed manager on the ephemeral `hardhat` network — that is fine and proves the path resolved.

- [ ] **Step 5: Make the Ignition admin overridable**

`ignition/modules/messenger.js:12-13` — uncomment the `managerAdmin` parameter so the admin can be set from the parameters file:

```javascript
const admin = m.getParameter("managerAdmin", m.getAccount(0));
```

Confirm the rest of the module uses this `admin` binding.

- [ ] **Step 6: Dry-run the module**

```bash
yarn hardhat ignition deploy ignition/modules/messenger.js --network hardhat
```

Expected: completes against the ephemeral network.

- [ ] **Step 7: Commit**

```bash
git add tasks/manager.js ignition/modules/messenger.js
git commit -m "fix(tasks): resolve --json relative to cwd; guard error.data

services:register is in the deploy critical path and crashed on relative
paths, because require() resolves against tasks/ rather than cwd.

error.data.data could TypeError on ethers v6 error shapes lacking .data,
masking the real revert and aborting the 63-service loop mid-way.

Also uncomments the Ignition managerAdmin parameter so the deploy admin
is overridable from the parameters file."
```

---

## Task 16: CI and lint tooling

`yarn lint` fails 100% of the time — it chains eslint and solhint, neither of which is installed or configured, while both `README.md` and `CLAUDE.md` document it. CI hides this by running prettier only. There is also no `abi/` drift check and no `push` trigger.

**Files:**
- Create: `eslint.config.js`, `.solhint.json`
- Modify: `package.json`, `.prettierignore`, `.github/workflows/ci.yaml`

**Interfaces:**
- Consumes: Task 1's toolchain
- Produces: working `yarn lint`; CI guards `abi/` drift

- [ ] **Step 1: Confirm the breakage**

```bash
yarn lint 2>&1 | tail -5
ls node_modules/.bin/ | grep -iE 'eslint|solhint' || echo "NEITHER INSTALLED"
```

Expected: failure, and `NEITHER INSTALLED`.

- [ ] **Step 2: Add the dependencies**

Add to `package.json` `devDependencies`:

```json
"eslint": "^9.17.0",
"solhint": "^5.0.5",
```

- [ ] **Step 3: Install**

The second and last permitted install.

```bash
yarn install
```

- [ ] **Step 4: Create the eslint config**

`eslint.config.js` (flat config, ESLint 9):

```javascript
const globals = require("globals");

module.exports = [
    {
        ignores: ["node_modules/**", "artifacts/**", "cache/**", "coverage/**", "abi/**", "docs/**", "ui/**", "go/**"],
    },
    {
        files: ["**/*.js"],
        languageOptions: {
            ecmaVersion: 2022,
            sourceType: "commonjs",
            globals: { ...globals.node, ...globals.mocha },
        },
        rules: {
            "no-unused-vars": ["error", { argsIgnorePattern: "^_" }],
            "no-undef": "error",
        },
    },
];
```

If `globals` is not already present, add `"globals": "^15.14.0"` to devDependencies and re-run `yarn install`.

- [ ] **Step 5: Create the solhint config**

`.solhint.json`:

```json
{
    "extends": "solhint:recommended",
    "rules": {
        "compiler-version": ["error", "0.8.24"],
        "func-visibility": ["warn", { "ignoreConstructors": true }],
        "no-inline-assembly": "off",
        "max-line-length": ["warn", 120],
        "not-rely-on-time": "off"
    }
}
```

`no-inline-assembly` is off because every ERC-7201 storage accessor uses `assembly { $.slot := ... }`. `not-rely-on-time` is off because the reservation-expiry design depends on `block.timestamp` by intent.

- [ ] **Step 6: Fix the solhint glob**

`package.json:16` — `solhint contracts/**/*.sol` relies on shell globbing that misses nested directories in some shells. Change to:

```json
"lint:solhint": "yarn solhint 'contracts/**/*.sol'",
```

- [ ] **Step 7: Run lint and fix what it finds**

```bash
yarn lint
```

Fix genuine errors. If a rule is noisy against this codebase's established style, turn it off in the config rather than churning the contracts — this task is about making `yarn lint` work, not restyling the repo.

- [ ] **Step 8: Add `.superpowers/` to prettierignore**

Append to `.prettierignore`:

```
/.superpowers/
```

Verify:

```bash
yarn lint:prettier
```

Expected: passes.

- [ ] **Step 9: Add the push trigger and the abi drift job**

In `.github/workflows/ci.yaml`, add a `push` trigger alongside `pull_request`:

```yaml
on:
    pull_request:
    push:
        branches: [dev, main]
```

Add an `abi` job mirroring the existing `docs` job — check out, setup-node, `yarn install --frozen-lockfile`, `yarn compile`, `yarn hardhat export-abi`, then:

```yaml
            - name: Check abi/ is up to date
              run: |
                  if [ -n "$(git status --porcelain abi/)" ]; then
                      echo "abi/ is out of date. Run 'yarn hardhat export-abi' and commit."
                      git --no-pager diff --stat abi/
                      exit 1
                  fi
```

- [ ] **Step 10: Enable the full lint in CI**

Change the lint step from `yarn lint:prettier` to `yarn lint`.

- [ ] **Step 11: Bump the action versions**

`actions/setup-node@v3` → `@v4` at lines 15, 33, 81, 115 (line 51 is already v4). `actions/setup-go@v4` → `@v5` at line 86.

- [ ] **Step 12: Validate the workflow parses**

```bash
python3 -c "import yaml,sys; yaml.safe_load(open('.github/workflows/ci.yaml')); print('valid YAML')"
```

Expected: `valid YAML`.

- [ ] **Step 13: Commit**

```bash
git add package.json yarn.lock eslint.config.js .solhint.json .prettierignore .github/workflows/ci.yaml
git commit -m "ci: make yarn lint work and guard abi/ drift

yarn lint chained eslint and solhint, neither installed nor configured,
so it failed 100% of the time while README and CLAUDE both documented
it. CI masked this by running prettier only.

Also adds an abi/ drift job (abi/ is the UI's sole ABI source and was
unguarded), a push trigger so merges are validated, and bumps
setup-node/setup-go."
```

---

## Task 17: Regenerate Go bindings and write the deployment runbook

Final task. Regenerates the artifact tree that intermediate commits left stale, and documents the post-deploy sequence — without which the deployment produces an inert system.

**Files:**
- Modify: `go/contracts/**` (regenerated)
- Modify: `README.md`
- Modify: `CLAUDE.md`

**Interfaces:**
- Consumes: every prior task
- Produces: `go/contracts/` in sync; a documented runbook

- [ ] **Step 1: Set the scratch environment**

Disk is ~97% full on `/` and `/home`.

```bash
export TMPDIR=/hgst/work/.ttm-scratch
export GOCACHE=/hgst/work/.ttm-scratch/gocache
export GOTMPDIR=/hgst/work/.ttm-scratch
export GOMODCACHE=/hgst/work/.ttm-scratch/gomod
mkdir -p /hgst/work/.ttm-scratch
```

- [ ] **Step 2: Regenerate the Go bindings**

Note this script does its own `rm -rf node_modules && yarn install`, so it is slow.

```bash
bash scripts/generate_go_abi.sh
```

Expected: all bindings regenerate. Confirm `go/contracts/servicefeetoken/` did **not** reappear (it was deleted in Task 4 and is not in `ARTIFACTS`).

- [ ] **Step 3: Confirm every artifact tree is in sync**

```bash
yarn compile && yarn hardhat export-abi
git status --porcelain abi/ docs/ go/
```

Expected: only the intended regeneration diffs; nothing unexpected.

- [ ] **Step 4: Write the deployment runbook**

Replace the Deploy section of `README.md` (currently a single command) with:

````markdown
### Deploy (Hardhat Ignition)

The Ignition module deploys the contracts but leaves the system **inert** —
`initialize` grants only DEFAULT_ADMIN/PAUSER/UPGRADER/VERSIONER, so no
services can be registered until roles are granted. Follow every step.

```bash
# 1. Configuration
yarn hardhat vars set BASE_SEPOLIA_DEPLOYER_PRIVATE_KEY
yarn hardhat vars set ETHERSCAN_API_KEY
export BASE_SEPOLIA_URL=<your-rpc-url>      # optional; a public default exists

# 2. Deploy
yarn hardhat ignition deploy ignition/modules/messenger.js \
  --network base_sepolia --parameters ignition/base_sepolia_parameters.json

# 3. Grant the service-registry admin role (0 members after deploy)
yarn hardhat manager role:grant --role SERVICE_REGISTRY_ADMIN_ROLE \
  --address <deployer> --network base_sepolia

# 4. Register the 63 canonical services (~11.5M gas total)
yarn hardhat manager services:register \
  --json ./services/00_initial.json --network base_sepolia

# 5. Optional: grant MIN_EXPIRATION_ADMIN_ROLE to change the 60s default
# 6. Grant PAUSER_ROLE on BookingToken to the operations key

# 7. Verify on Basescan
yarn hardhat ignition verify chain-84532
```

Marking the two `ERC1967Proxy` addresses as **proxies** on Basescan is a
separate manual step in the Basescan UI.

**8. Hand off admin roles.** Transfer `DEFAULT_ADMIN_ROLE`, `UPGRADER_ROLE`,
and `VERSIONER_ROLE` to the Safe, verify the Safe can act, and only **then**
renounce the deployer's roles. The manager is a singleton — renouncing before
verifying is unrecoverable.

**9. Commit `ignition/deployments/chain-84532/`.** The UI reads
`deployed_addresses.json` and filters enabled chains by whether contracts
exist, so until this is committed the UI has zero enabled chains.

**10. Fill in the address table** at the top of this README.

**11. Bump the contracts Go module** in the bot and matrix-app-service.
````

- [ ] **Step 5: Note the accepted risk**

Add to `README.md` under the Deploy section:

```markdown
> **Known limitation (Base Sepolia).** `createTTMAccount` is currently
> permissionless — anyone can create an account. On Camino this was prevented
> by chain-level KYC, which does not exist on Base. Accepted for testnet
> (blast radius is spam, not funds); must be resolved before Base mainnet.
> See `docs/decisions/2026-07-21-contract-design-decisions.md`, Decision 1.
```

- [ ] **Step 6: Update CLAUDE.md**

Correct the deploy line to point at the README runbook rather than implying the single Ignition command is sufficient.

- [ ] **Step 7: Full verification**

```bash
yarn compile && yarn test && yarn lint
git status --porcelain
```

Expected: tests pass, lint passes, working tree clean apart from intended changes.

- [ ] **Step 8: Commit**

```bash
git add go/ abi/ docs/ README.md CLAUDE.md
git commit -m "docs: regenerate Go bindings and add the deployment runbook

The Ignition module deploys contracts but leaves the system inert:
SERVICE_REGISTRY_ADMIN_ROLE has zero members and none of the 63 services
are registered. That sequence was undocumented; the README implied the
single deploy command was sufficient.

Also records permissionless account creation as a known Base Sepolia
limitation that must close before mainnet."
```

---

## Final Verification

After Task 17, confirm the whole branch:

- [ ] `yarn compile` — succeeds; no contract over 22.5 KiB
- [ ] `yarn test` — passes; **exactly 129** (120 baseline + 12 new − 3 removed `reinitializeV2` tests)
- [ ] `yarn lint` — passes
- [ ] `git status --porcelain` — clean
- [ ] `grep -rn "camino\|CAM\b\|cheque\|prefund\|ServiceFeeToken\|developer fee" contracts/ tasks/ scripts/ test/ --include='*.sol' --include='*.js' -i` — only the intentional historical reference in the Decision-1 NatSpec pointer
- [ ] `abi/`, `docs/`, and `go/contracts/` all regenerate to no diff
- [ ] ERC-7201 constants unchanged: re-run the slot verification from Task 13 Step 2 for all seven namespaces

## Out of Scope

Do not implement any of the following — they are deliberately deferred:

- Account-creation gating (`ACCOUNT_CREATOR_ROLE`), transfer semantics during pending cancellation, pull-payment refunds, payment-token allowlist enforcement, bot role splitting, Safe multisig membership → `docs/decisions/2026-07-21-contract-design-decisions.md`
- Event signature changes, identity source-of-truth collapse, hash-native API and lens contract, `recordExpiration` posture, service deprecation events, additional test coverage → `docs/decisions/2026-07-21-technical-backlog.md`
