# Pre-Deploy Hash-Native Rework Implementation Plan

> **For agentic workers:** REQUIRED SUB-SKILL: Use superpowers:subagent-driven-development (recommended) or superpowers:executing-plans to implement this plan task-by-task. Steps use checkbox (`- [ ]`) syntax for tracking.

**Goal:** Move account identity from a role to a registry, make `TTMAccount`'s service API `bytes32`-native, and fix the event surface — all before the first deployment freezes them.

**Architecture:** `TTMAccountManager` stores accounts in an `EnumerableSet` with no external mutator, replacing `TTMACCOUNT_ROLE`. `TTMAccount` speaks only `bytes32` service hashes and emits hash-only events; the readable names move to `ServiceRegistry`'s events, which become the seed for client-side name↔hash maps. The UI gains one shared resolver; the Go bot gets a migration document written from the merged ABI.

**Tech Stack:** Solidity 0.8.24 (optimizer `runs: 1000`, `evmVersion: cancun`), Hardhat 2.x, OpenZeppelin upgradeable 5.x, ethers v6 + chai for tests, React + viem/wagmi + vitest in `ui/`.

**Spec:** `docs/superpowers/specs/2026-07-21-hash-native-rework-design.md`

## Global Constraints

- **Baseline to preserve:** 134 tests passing. A test removed during a signature change is a test that was quietly load-bearing — adapt, never delete.
- **Baseline sizes (init, KiB), 22.5 gate:** `TTMAccountManager` 12.800 · `TTMAccount` 21.371 · `BookingToken` 21.552. Record measured sizes, never estimates.
- **Every symbol-adding or -removing change must run** `yarn hardhat clean && yarn compile && yarn docgen && yarn hardhat export-abi` — `yarn compile` alone does **not** reliably run docgen (the compile cache short-circuits it despite `docgen.runOnCompile: true`), so `docs/` silently goes stale and CI's clean-tree check fails.
- **After `export-abi`, run `(cd ui && yarn sync)`** — `ui/scripts/sync-contracts.ts` reads `abi/` to regenerate `ui/src/contracts/generated/abis.ts`. Skipping it leaves the UI on a stale ABI.
- **Everything under `docs/` must be prettier-formatted** — `yarn docgen` ends in `prettier --write docs/`, so an unformatted file there dirties the tree and fails both the `docs` job and `lint`.
- **Go bindings embed bytecode**, so even a comment change shifts them via Solidity's metadata hash. They are regenerated **once**, in Task 11, because `scripts/generate_go_abi.sh` does its own `rm -rf node_modules && yarn install`. `/` is ~97% full — route `TMPDIR`, `GOCACHE`, `GOMODCACHE`, `GOTMPDIR`, and `YARN_CACHE_FOLDER` to `/hgst` before running it.
- **`vars.get(NAME)` with no default throws at config load** (`HH1201`), breaking *every* Hardhat command. Never add one without a fallback.
- **The Go bindings under `go/` are knowingly stale from Task 1 until Task 11 regenerates them.** This is deliberate sequencing, not an oversight: `scripts/generate_go_abi.sh` does its own `rm -rf node_modules && yarn install`, so running it per task is wasteful. Verified safe: `.github/workflows/ci.yaml` triggers only on `pull_request` and pushes to `dev`/`main`, so no CI runs on this feature branch in that window. **Do not open a PR before Task 11 completes** — the `go-bindings` job asserts a clean tree after regeneration and would fail.
- **Verify comments against code before acting on them.** Stale comments have twice produced wrong conclusions in this repo. Treat every line number in the technical backlog as unverified — five of its claims were already found wrong.
- **Verified signer and fixture names** (read from `test/utils/fixtures.js` on 2026-07-21 — earlier drafts of this plan guessed wrong twice). Signers: `managerAdmin`, `managerPauser`, `managerUpgrader`, `managerVersioner`, `ttmAccountAdmin`, `ttmAccountUpgrader`, `ttmServiceAdmin`, `botOperator`, `depositor`, `withdrawer`, `btAdmin`, `btUpgrader`, `registryAdmin`, `otherAccount1`, `otherAccount2`, `otherAccount3`. Fixtures: `deployNullUSDFixture`, `deployTTMAccountManagerFixture`, `deployTTMAccountImplFixture`, `deployTTMAccountManagerWithTTMAccountImplFixture`, `deployTTMAccountWithDepositFixture`, `deployBookingTokenFixture`, `deployBookingTokenWithNullUSDFixture`, `deployCancellationSupportFixture`, `deployAndConfigureAllFixture`, `deployAndConfigureAllWithRegisteredServicesFixture`. Note `deployTTMAccountManagerWithTTMAccountImplFixture` never calls `setBookingTokenAddress`, so `createTTMAccount` reverts `InvalidBookingTokenAddress` under it — use `deployAndConfigureAllFixture` when you need to create an account. Never invent a signer or fixture.
- **Solidity style:** 4-space indent, NatSpec on all public/external members. **JS tests:** 4-space indent, `describe`/`it`, fixtures from `test/utils/fixtures`. **UI:** 2-space indent, prettier-enforced.

## Deliberately NOT in scope

The hash-native direction of Tasks 5–7 does **not** generalise to everything, and two nearby changes are rejected on purpose. Do not implement them "for consistency":

- **Do not make `capabilities` a `bytes32[]`** (`contracts/partner/PartnerConfiguration.sol:26-33`). It would let the `Service` struct pack and turn the O(n) `keccak256` comparison in `_removeServiceCapability` into a word compare. But capabilities are arbitrary partner-typed strings, rendered and edited directly in `ui/src/pages/tabs/ServicesTab.tsx:156,213`, and — unlike service names — they have **no registry**, so there is nothing to resolve a hash back against. Hashing them makes them permanently unreadable to save one storage slot. The presence of a registry is exactly what makes hashing safe for service names and unsafe here.
- **Do not resolve Decision 1 by adding an access-control gate to `createTTMAccount`.** Task 10 adds a test that pins the current permissionless behaviour on purpose. Changing it is a business decision that has not been made — see `docs/decisions/2026-07-21-contract-design-decisions.md`.
- **Do not build a `TTMLens` contract.** The technical backlog §3 proposes one; the spec withdraws it after measuring what it would actually buy the bot (one `eth_call`, once, at startup).

---

### Task 1: Manager identity becomes a registry

Replaces `TTMACCOUNT_ROLE` and the `TTMAccountInfo` struct with an `EnumerableSet`. This is the only storage-layout change in the plan and the single most time-critical item — it is free today and a migration after deploy.

**Files:**

- Modify: `contracts/manager/TTMAccountManager.sol` (storage struct `:97-117`, role constant `:87`, `_createTTMAccount` `:273-280`, `_setTTMAccountInfo` `:290-293`, `isTTMAccount` `:308-311`, `getTTMAccountCreator` `:296-302`)
- Modify: `tasks/manager.js:15,252`, `tasks/account.js:781`
- Modify: `ui/src/hooks/useMyAccounts.ts` (`useManagerAccounts`, `:20-48`)
- Delete: `ui/src/components/AccountValidityNotice.tsx`
- Test: `test/TTMAccountManager.test.js`

**Interfaces:**

- Consumes: nothing (first task).
- Produces: `isTTMAccount(address) → bool` (unchanged signature), `getTTMAccountCreator(address) → address` (unchanged signature), `getTTMAccountCount() → uint256`, `getTTMAccounts() → address[]`, `getTTMAccountsSlice(uint256 offset, uint256 limit) → address[]`. `TTMACCOUNT_ROLE` no longer exists. Later tasks must not reference it.

> **Naming note:** `getTTMAccounts()` and `getTTMAccountsSlice(...)` are deliberately **not** overloads of one name. The repo already has overloaded getters (`getServiceCapabilities`), and they force `contract["fn(uint256,uint256)"](...)` disambiguation in ethers v6 and awkward `functionName` handling in wagmi. Distinct names avoid that everywhere.

- [ ] **Step 1: Write the failing tests**

Add to `test/TTMAccountManager.test.js`, inside the top-level `describe("TTMAccountManager", ...)`:

```js
    describe("Account registry", function () {
        it("should record created accounts in the registry with their creator", async function () {
            await setupSigners();
            const { ttmAccountManager } = await loadFixture(deployTTMAccountManagerWithTTMAccountImplFixture);

            expect(await ttmAccountManager.getTTMAccountCount()).to.equal(0);
            expect(await ttmAccountManager.getTTMAccounts()).to.deep.equal([]);

            const tx = await ttmAccountManager
                .connect(signers.otherAccount1)
                .createTTMAccount(signers.otherAccount1.address, signers.ttmAccountUpgrader.address);
            const receipt = await tx.wait();
            const created = receipt.logs
                .map((l) => {
                    try {
                        return ttmAccountManager.interface.parseLog(l);
                    } catch {
                        return null;
                    }
                })
                .find((p) => p && p.name === "TTMAccountCreated");
            const account = created.args.account;

            expect(await ttmAccountManager.isTTMAccount(account)).to.be.true;
            expect(await ttmAccountManager.getTTMAccountCreator(account)).to.equal(signers.otherAccount1.address);
            expect(await ttmAccountManager.getTTMAccountCount()).to.equal(1);
            expect(await ttmAccountManager.getTTMAccounts()).to.deep.equal([account]);
        });

        it("should report unknown addresses as not being TTM Accounts", async function () {
            await setupSigners();
            const { ttmAccountManager } = await loadFixture(deployTTMAccountManagerWithTTMAccountImplFixture);

            expect(await ttmAccountManager.isTTMAccount(signers.otherAccount1.address)).to.be.false;
            expect(await ttmAccountManager.getTTMAccountCreator(signers.otherAccount1.address)).to.equal(
                ethers.ZeroAddress,
            );
        });

        it("should expose no external way to add an account to the registry", async function () {
            await setupSigners();
            const { ttmAccountManager } = await loadFixture(deployTTMAccountManagerWithTTMAccountImplFixture);

            // The factory path is the only writer. Assert structurally: no ABI
            // entry other than createTTMAccount can mutate account identity.
            const mutators = ttmAccountManager.interface.fragments.filter(
                (f) =>
                    f.type === "function" &&
                    f.stateMutability !== "view" &&
                    f.stateMutability !== "pure" &&
                    /ttmaccount/i.test(f.name) &&
                    f.name !== "createTTMAccount",
            );
            expect(mutators.map((f) => f.name)).to.deep.equal([]);

            // And TTMACCOUNT_ROLE is gone entirely, so it cannot be granted.
            expect(ttmAccountManager.interface.fragments.some((f) => f.name === "TTMACCOUNT_ROLE")).to.be.false;
        });

        it("should paginate the account list", async function () {
            await setupSigners();
            const { ttmAccountManager } = await loadFixture(deployTTMAccountManagerWithTTMAccountImplFixture);

            for (let i = 0; i < 3; i++) {
                await ttmAccountManager
                    .connect(signers.otherAccount1)
                    .createTTMAccount(signers.otherAccount1.address, signers.ttmAccountUpgrader.address);
            }
            const all = await ttmAccountManager.getTTMAccounts();
            expect(all.length).to.equal(3);

            expect(await ttmAccountManager.getTTMAccountsSlice(0, 2)).to.deep.equal([all[0], all[1]]);
            expect(await ttmAccountManager.getTTMAccountsSlice(2, 2)).to.deep.equal([all[2]]);
            expect(await ttmAccountManager.getTTMAccountsSlice(3, 1)).to.deep.equal([]);
            expect(await ttmAccountManager.getTTMAccountsSlice(0, 100)).to.deep.equal(all);
        });
    });
```

> If `deployTTMAccountManagerWithTTMAccountImplFixture` does not also wire the BookingToken address, `createTTMAccount` will revert with `InvalidBookingTokenAddress`. Check `test/utils/fixtures.js` first and use `deployAndConfigureAllFixture` instead if so — do not add a new fixture.

- [ ] **Step 2: Run the tests to verify they fail**

Run: `yarn hardhat test test/TTMAccountManager.test.js --grep "Account registry"`
Expected: FAIL — `ttmAccountManager.getTTMAccountCount is not a function`.

- [ ] **Step 3: Replace the storage struct**

In `contracts/manager/TTMAccountManager.sol`, add the import beside the existing OpenZeppelin imports:

```solidity
import { EnumerableSet } from "@openzeppelin/contracts/utils/structs/EnumerableSet.sol";
```

Add the `using` directive at the top of the contract body, beside any existing ones:

```solidity
    using EnumerableSet for EnumerableSet.AddressSet;
```

Delete the `TTMAccountInfo` struct (`:97-100`) entirely, and replace the storage struct's third member:

```solidity
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
```

Leave `TTMAccountManagerStorageLocation` (`:120-121`) **byte-identical** — the namespace string is unchanged, so the constant is too.

- [ ] **Step 4: Delete the role and rewrite the accessors**

Delete the `TTMACCOUNT_ROLE` constant (`:87`) and its NatSpec block, and delete `_setTTMAccountInfo` (`:290-293`).

Replace `getTTMAccountCreator` and `isTTMAccount` with:

```solidity
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
```

- [ ] **Step 5: Rewrite the factory's recording step**

In `_createTTMAccount`, replace the `_setTTMAccountInfo` call and the `_grantRole(TTMACCOUNT_ROLE, ...)` call (`:276-280`) with:

```solidity
        // Record the account and its creator. This is the only writer.
        TTMAccountManagerStorage storage $ = _getTTMAccountManagerStorage();
        $._ttmAccounts.add(address(ttmAccountProxy));
        $._ttmAccountCreator[address(ttmAccountProxy)] = msg.sender;
```

Leave the `emit TTMAccountCreated(...)` line alone — Task 2 widens it.

- [ ] **Step 6: Run the full suite**

Run: `yarn test`
Expected: the four new tests PASS. Some existing tests will FAIL where they reference `TTMACCOUNT_ROLE` — fix each by switching to `isTTMAccount` / `getTTMAccounts`. **Adapt, do not delete.** Expected total afterwards: 138 passing.

- [ ] **Step 7: Update the Hardhat tasks**

In `tasks/manager.js`, delete the obsolete commented-out `"TTMACCOUNT_ROLE"` entry and its trailing comment at `:15` (the performance concern it describes no longer applies). At `:252`, replace the `role:members` invocation with a direct call:

```js
    const accounts = await manager.getTTMAccounts();
    console.log(`TTM Accounts (${accounts.length}):`);
    for (const account of accounts) {
        console.log(`  ${account}  creator: ${await manager.getTTMAccountCreator(account)}`);
    }
```

In `tasks/account.js:781`, replace the `manager.TTMACCOUNT_ROLE()` lookup and its subsequent role check with `await manager.isTTMAccount(address)`. Read the surrounding block first — the variable is named `ttmAccountRole` and may be used more than once.

- [ ] **Step 8: Update the UI hook and delete the notice**

Replace `useManagerAccounts` in `ui/src/hooks/useMyAccounts.ts`:

```ts
/**
 * Lists every TTM Account by reading the manager's account registry. This
 * mirrors the `account find` CLI task and avoids eth_getLogs, which free-tier
 * RPCs reject for wide block ranges.
 */
export function useManagerAccounts() {
  const { manager, managerAbi } = useActiveContracts();
  const { activeChainId } = useActiveChain();
  const abi = managerAbi as Abi;

  const { data, isLoading } = useReadContract({
    chainId: activeChainId,
    address: manager,
    abi,
    functionName: "getTTMAccounts",
    query: { enabled: Boolean(manager) },
  });

  const accounts = uniqueAddresses((data as string[]) ?? []) as Address[];
  return { accounts, isLoading };
}
```

Remove the now-unused `ACCOUNT_ROLES, roleHash` import if nothing else in the file uses it — check first, `useAccountStats` may.

Delete `ui/src/components/AccountValidityNotice.tsx` and every import/usage of it:

```bash
rm ui/src/components/AccountValidityNotice.tsx
grep -rn "AccountValidityNotice" ui/src
```

Remove each hit. Also update `ui/src/hooks/useMyAccounts.test.ts` if it mocks `TTMACCOUNT_ROLE` or `getRoleMembers`.

- [ ] **Step 9: Regenerate artifacts**

```bash
yarn hardhat clean && yarn compile && yarn docgen && yarn hardhat export-abi
(cd ui && yarn sync)
```

- [ ] **Step 10: Verify everything is green**

```bash
yarn test
yarn lint
(cd ui && yarn test && yarn build)
git status --porcelain
```

Expected: 138 contract tests pass, lint clean, UI tests and build pass, and `git status --porcelain` shows only files you intend to commit — no unexpected `docs/` or `abi/` drift.

- [ ] **Step 11: Commit**

```bash
git add -A
git commit -m "refactor(manager)!: replace TTMACCOUNT_ROLE with an account registry

Account identity was modelled as an AccessControl role but never used as an
authorization gate - only enumerated. Being a role is what exposed it to
grantRole divergence from the isTTMAccount mapping. An EnumerableSet with a
single writer removes that class of bug structurally.

BREAKING CHANGE: TTMACCOUNT_ROLE is removed. Use isTTMAccount(address) or
getTTMAccounts() instead of getRoleMembers(TTMACCOUNT_ROLE)."
```

---

### Task 2: Registry and manager events

Moves service names onto the registry's events, so consumers can build a name↔hash map from logs alone, and widens `TTMAccountCreated`.

**Files:**

- Modify: `contracts/partner/ServiceRegistry.sol:44-45` (declarations), `:95` and `:117` (emits)
- Modify: `contracts/manager/TTMAccountManager.sol` (`TTMAccountCreated` declaration `:142`, emit in `_createTTMAccount`)
- Modify: `ui/src/lib/receipt.ts:7`
- Test: `test/ServiceRegistry.test.js`, `test/TTMAccountManager.test.js`

**Interfaces:**

- Consumes: Task 1's registry (the emit site sits in the block Task 1 rewrote).
- Produces: `event ServiceRegistered(bytes32 indexed serviceHash, string serviceName)`, `event ServiceUnregistered(bytes32 indexed serviceHash, string serviceName)`, `event TTMAccountCreated(address indexed account, address indexed creator, address indexed admin)`.

- [ ] **Step 1: Write the failing tests**

Add to `test/ServiceRegistry.test.js`:

```js
    it("should emit the service name in the data section and index the hash", async function () {
        await setupSigners();
        const { ttmAccountManager } = await loadFixture(deployTTMAccountManagerFixture);

        const name = "ttm.services.accommodation.v1alpha.AccommodationSearchService";
        const hash = ethers.keccak256(ethers.toUtf8Bytes(name));

        await expect(ttmAccountManager.connect(signers.registryAdmin).registerService(name))
            .to.emit(ttmAccountManager, "ServiceRegistered")
            .withArgs(hash, name);

        await expect(ttmAccountManager.connect(signers.registryAdmin).unregisterService(name))
            .to.emit(ttmAccountManager, "ServiceUnregistered")
            .withArgs(hash, name);
    });
```

Add to `test/TTMAccountManager.test.js`, inside the `describe("Account registry", ...)` block from Task 1:

```js
        it("should emit TTMAccountCreated with creator and admin", async function () {
            await setupSigners();
            const { ttmAccountManager } = await loadFixture(deployTTMAccountManagerWithTTMAccountImplFixture);

            await expect(
                ttmAccountManager
                    .connect(signers.otherAccount1)
                    .createTTMAccount(signers.otherAccount1.address, signers.ttmAccountUpgrader.address),
            )
                .to.emit(ttmAccountManager, "TTMAccountCreated")
                .withArgs(anyValue, signers.otherAccount1.address, signers.otherAccount1.address);
        });
```

Add the `anyValue` import at the top of `test/TTMAccountManager.test.js` if not already present:

```js
const { anyValue } = require("@nomicfoundation/hardhat-chai-matchers/withArgs");
```

> `signers.registryAdmin` must actually hold `SERVICE_REGISTRY_ADMIN_ROLE` in that fixture. Check `test/utils/fixtures.js`; the role starts with zero members after a bare deploy, so the fixture may grant it — if it does not, use whichever signer the existing registry tests use.

- [ ] **Step 2: Run the tests to verify they fail**

Run: `yarn hardhat test test/ServiceRegistry.test.js test/TTMAccountManager.test.js`
Expected: FAIL on the argument order — `ServiceRegistered` currently emits `(serviceName, serviceHash)`, and `TTMAccountCreated` takes one argument.

- [ ] **Step 3: Change the registry events**

In `contracts/partner/ServiceRegistry.sol`, replace the two declarations at `:44-45`:

```solidity
    /**
     * @notice Emitted when a service is registered.
     *
     * @dev The hash is indexed for filtering; the name travels in the data section so
     * consumers can build a complete name-to-hash map from logs alone, with no
     * `eth_call`. This is the authoritative publication of that mapping - `TTMAccount`
     * emits hashes only.
     */
    event ServiceRegistered(bytes32 indexed serviceHash, string serviceName);

    /**
     * @notice Emitted when a service is unregistered.
     *
     * @dev Existing accounts can still resolve a deprecated name, so this is the only
     * signal that a service was retired. See {_unregisterServiceName}.
     */
    event ServiceUnregistered(bytes32 indexed serviceHash, string serviceName);
```

Swap the argument order at both emit sites — in `_registerServiceName` (around `:95`) and `_unregisterServiceName` (`:117`):

```solidity
        emit ServiceRegistered(serviceHash, serviceName);
```

```solidity
        emit ServiceUnregistered(serviceHash, serviceName);
```

- [ ] **Step 4: Widen `TTMAccountCreated`**

In `contracts/manager/TTMAccountManager.sol`, replace the declaration at `:142`:

```solidity
    /**
     * @notice Emitted when a TTM Account is created.
     *
     * @dev Carries creator and admin so indexers need no follow-up call per account.
     */
    event TTMAccountCreated(address indexed account, address indexed creator, address indexed admin);
```

And the emit in `_createTTMAccount`:

```solidity
        emit TTMAccountCreated(address(ttmAccountProxy), msg.sender, admin);
```

- [ ] **Step 5: Run the tests to verify they pass**

Run: `yarn test`
Expected: PASS. Existing assertions on the old argument order will fail — fix them in place. Expected total: 140 passing.

- [ ] **Step 6: Update the UI receipt decoder**

`ui/src/lib/receipt.ts:7` decodes `TTMAccountCreated` and previously got only the address. Extend it to surface `creator` and `admin`. Read the file first; the change is to widen the decoded shape and whatever type it returns. Update `ui/src/lib/receipt.test.ts` to match.

Also update `ui/src/lib/activity/catalog.ts` — the `ServiceRegistered` / `ServiceUnregistered` entries (`:71-72`) read `a.serviceName`, which still works since the name remains a named arg, but confirm and adjust the `TTMAccountCreated` renderer if you want creator in the sentence.

- [ ] **Step 7: Regenerate, verify, and commit**

```bash
yarn hardhat clean && yarn compile && yarn docgen && yarn hardhat export-abi
(cd ui && yarn sync)
yarn test && yarn lint
(cd ui && yarn test && yarn build)
git add -A
git commit -m "feat(events)!: publish service names from the registry, widen TTMAccountCreated

ServiceRegistered/ServiceUnregistered now index the hash and carry the name in
the data section, making them the authoritative name-to-hash publication.
TTMAccountCreated carries creator and admin so indexers need no follow-up call.

BREAKING CHANGE: ServiceRegistered and ServiceUnregistered argument order is
swapped and the hash is now indexed. TTMAccountCreated has three arguments."
```

---

### Task 3: Events for BookingToken's silent setters

Two admin actions currently change consequential state and emit nothing.

**Files:**

- Modify: `contracts/booking-token/BookingToken.sol` (`setManagerAddress` `:684-687`, `setMinExpirationTimestampDiff` `:702-707`, plus the events block at `:188-212`)
- Test: `test/BookingToken.test.js`

> `contracts/booking-token/IBookingToken.sol` declares **no events** (verified 2026-07-21) — all `BookingToken` events live in the contract itself. No interface change is needed.

**Interfaces:**

- Consumes: nothing from earlier tasks.
- Produces: `event ManagerAddressUpdated(address indexed oldManager, address indexed newManager)`, `event MinExpirationTimestampDiffUpdated(uint256 oldDiff, uint256 newDiff)`.

- [ ] **Step 1: Write the failing tests**

Add to `test/BookingToken.test.js`:

```js
    it("should emit ManagerAddressUpdated when the manager is repointed", async function () {
        await setupSigners();
        const { bookingToken, ttmAccountManager } = await loadFixture(deployAndConfigureAllFixture);

        const oldManager = await bookingToken.getManagerAddress();
        const newManager = await ttmAccountManager.getAddress();

        await expect(bookingToken.connect(signers.btAdmin).setManagerAddress(newManager))
            .to.emit(bookingToken, "ManagerAddressUpdated")
            .withArgs(oldManager, newManager);
    });

    it("should emit MinExpirationTimestampDiffUpdated when the mint rule changes", async function () {
        await setupSigners();
        const { bookingToken } = await loadFixture(deployAndConfigureAllFixture);

        const oldDiff = await bookingToken.getMinExpirationTimestampDiff();

        await expect(bookingToken.connect(signers.btAdmin).setMinExpirationTimestampDiff(120))
            .to.emit(bookingToken, "MinExpirationTimestampDiffUpdated")
            .withArgs(oldDiff, 120);
    });
```

> Check the real names of the admin signer and the getter (`getMinExpirationTimestampDiff`) in the existing `BookingToken` tests before running — use whatever those tests already use rather than inventing names.

- [ ] **Step 2: Run the tests to verify they fail**

Run: `yarn hardhat test test/BookingToken.test.js --grep "Updated when"`
Expected: FAIL — the events do not exist.

- [ ] **Step 3: Declare the events**

In `contracts/booking-token/BookingToken.sol`, add to the events section:

```solidity
    /**
     * @notice Emitted when the manager address is changed.
     *
     * @dev This repoints the entire authorization oracle for this token - `isTTMAccount`
     * resolves through the manager - so the change is worth an explicit log.
     */
    event ManagerAddressUpdated(address indexed oldManager, address indexed newManager);

    /**
     * @notice Emitted when the minimum expiration timestamp difference changes.
     *
     * @dev This is a mint-time validation rule; changing it changes which mints succeed.
     */
    event MinExpirationTimestampDiffUpdated(uint256 oldDiff, uint256 newDiff);
```

- [ ] **Step 4: Emit them**

Add the emit to each setter, capturing the old value before the write. For `setManagerAddress`:

```solidity
        BookingTokenStorage storage $ = _getBookingTokenStorage();
        address oldManager = $._manager;
        $._manager = manager;
        emit ManagerAddressUpdated(oldManager, manager);
```

For `setMinExpirationTimestampDiff`:

```solidity
        BookingTokenStorage storage $ = _getBookingTokenStorage();
        uint256 oldDiff = $._minExpirationTimestampDiff;
        $._minExpirationTimestampDiff = minExpirationTimestampDiff;
        emit MinExpirationTimestampDiffUpdated(oldDiff, minExpirationTimestampDiff);
```

> Read both functions first — the storage field names above are inferred from the getters. Use the actual field names from `BookingTokenStorage`.

- [ ] **Step 5: Run the tests to verify they pass**

Run: `yarn test`
Expected: PASS, 142 passing.

- [ ] **Step 6: Regenerate, verify, and commit**

```bash
yarn hardhat clean && yarn compile && yarn docgen && yarn hardhat export-abi
(cd ui && yarn sync)
yarn test && yarn lint
git add -A
git commit -m "feat(booking-token): emit events for manager and expiration-rule changes"
```

---

### Task 4: UI service catalog resolver

Built **before** the contract changes, because Tasks 5–7 each fix their own UI call sites and need this to exist. It is a pure addition against the current ABI — nothing breaks, nothing is removed yet.

**Files:**

- Create: `ui/src/lib/serviceCatalog.ts`, `ui/src/lib/serviceCatalog.test.ts`
- Create or modify: `ui/src/hooks/useServiceCatalog.ts` (see Step 5)

**Interfaces:**

- Consumes: `getAllRegisteredServiceNames()` on the manager — already public, unchanged by this plan.
- Produces: `hashServiceName(name: string): \`0x${string}\``, `buildServiceCatalog(names: string[]): ServiceCatalog`, `serviceNameForHash(catalog: ServiceCatalog, hash: string): string | undefined`, `type ServiceCatalog = { nameByHash: Map<string, string>; hashByName: Map<string, \`0x${string}\`> }`, and hook `useServiceCatalog(): { catalog: ServiceCatalog; isLoading: boolean }`. Tasks 5, 6 and 7 all consume these.

- [ ] **Step 1: Write the failing test**

Create `ui/src/lib/serviceCatalog.test.ts`:

```ts
import { describe, expect, it } from "vitest";
import { keccak256, toBytes } from "viem";
import { buildServiceCatalog, hashServiceName, serviceNameForHash } from "./serviceCatalog";

describe("serviceCatalog", () => {
  const name = "ttm.services.accommodation.v1alpha.AccommodationSearchService";

  it("hashes a service name the same way the contracts do", () => {
    // Contracts use keccak256(abi.encodePacked(serviceName)), which for a lone
    // string argument is just the keccak of its UTF-8 bytes.
    expect(hashServiceName(name)).toBe(keccak256(toBytes(name)));
  });

  it("builds a bidirectional map", () => {
    const catalog = buildServiceCatalog([name]);
    const hash = hashServiceName(name);

    expect(catalog.nameByHash.get(hash.toLowerCase())).toBe(name);
    expect(catalog.hashByName.get(name)).toBe(hash);
  });

  it("resolves a hash regardless of its casing", () => {
    const catalog = buildServiceCatalog([name]);
    const hash = hashServiceName(name);

    expect(serviceNameForHash(catalog, hash.toUpperCase().replace("0X", "0x"))).toBe(name);
  });

  it("returns undefined for an unknown hash", () => {
    const catalog = buildServiceCatalog([name]);

    expect(serviceNameForHash(catalog, hashServiceName("ttm.services.nope.v1.NopeService"))).toBeUndefined();
  });

  it("returns empty maps for an empty catalog", () => {
    const catalog = buildServiceCatalog([]);

    expect(catalog.nameByHash.size).toBe(0);
    expect(catalog.hashByName.size).toBe(0);
  });
});
```

- [ ] **Step 2: Run it to verify it fails**

Run: `cd ui && yarn test serviceCatalog`
Expected: FAIL — the module does not exist.

- [ ] **Step 3: Implement the resolver**

Create `ui/src/lib/serviceCatalog.ts`:

```ts
import { keccak256, toBytes } from "viem";

/**
 * Service names and hashes are two views of one thing. Contracts store and emit
 * `bytes32` hashes; people read names. The registry is the only authority that
 * knows both, so we seed from it once and resolve locally thereafter — which is
 * why `TTMAccount` needs no name-resolution code and its reads cost no extra
 * cross-contract calls.
 */
export interface ServiceCatalog {
  /** Keyed by lowercase hash, because callers get hashes from several sources. */
  nameByHash: Map<string, string>;
  hashByName: Map<string, `0x${string}`>;
}

/**
 * Mirrors `keccak256(abi.encodePacked(serviceName))` in ServiceRegistry. For a
 * single string argument, `encodePacked` is just the raw UTF-8 bytes.
 */
export function hashServiceName(name: string): `0x${string}` {
  return keccak256(toBytes(name));
}

/** Builds the bidirectional map from the registry's list of registered names. */
export function buildServiceCatalog(names: string[]): ServiceCatalog {
  const nameByHash = new Map<string, string>();
  const hashByName = new Map<string, `0x${string}`>();

  for (const name of names) {
    const hash = hashServiceName(name);
    nameByHash.set(hash.toLowerCase(), name);
    hashByName.set(name, hash);
  }

  return { nameByHash, hashByName };
}

/** Resolves a hash to its registered name, or undefined if it is unknown. */
export function serviceNameForHash(catalog: ServiceCatalog, hash: string): string | undefined {
  return catalog.nameByHash.get(hash.toLowerCase());
}
```

- [ ] **Step 4: Run the test to verify it passes**

Run: `cd ui && yarn test serviceCatalog`
Expected: PASS, 5 tests.

- [ ] **Step 5: Add the hook**

Create `ui/src/hooks/useServiceCatalog.ts` — check how `ui/src/hooks/useMyAccounts.ts` imports `useActiveContracts`, `useActiveChain` and `useReadContract`, and match it exactly:

```ts
import { useMemo } from "react";
import { type Abi } from "viem";
import { useReadContract } from "wagmi";
import { useActiveContracts } from "./useActiveContracts";
import { useActiveChain } from "../wallet/activeChain";
import { buildServiceCatalog } from "../lib/serviceCatalog";

/**
 * Reads the registry's full service list once and derives the name↔hash map
 * locally. One eth_call replaces the per-hash resolution round-trips the
 * activity feed used to make.
 */
export function useServiceCatalog() {
  const { manager, managerAbi } = useActiveContracts();
  const { activeChainId } = useActiveChain();

  const { data, isLoading } = useReadContract({
    chainId: activeChainId,
    address: manager,
    abi: managerAbi as Abi,
    functionName: "getAllRegisteredServiceNames",
    query: { enabled: Boolean(manager) },
  });

  const catalog = useMemo(() => buildServiceCatalog((data as string[]) ?? []), [data]);

  return { catalog, isLoading };
}
```

- [ ] **Step 6: Verify and commit**

```bash
cd ui && yarn test && yarn build && yarn lint:format
```

Expected: all green — this task adds code and removes none.

```bash
git add -A
git commit -m "feat(ui): add the service catalog resolver

Seeds a bidirectional name/hash map from one getAllRegisteredServiceNames()
read. Groundwork for the hash-native contract API."
```

---

### Task 5: Hash-native service CRUD on TTMAccount

The core of the rework. Service write functions take `bytes32`, and the service events become hash-only — these cannot be separated, because a hash-native function has no name to emit without reintroducing the very staticcall this removes.

**Files:**

- Modify: `contracts/account/TTMAccount.sol` (events `:164-166,170-174`, `addService` `:439-447`, `removeService` `:460-463`, `removeAllServices` `:469-476`, `setServiceRestrictedRate`, `addServiceCapability`, `removeServiceCapability`, `setServiceCapabilities` — all in `:427-635`)
- Modify: `ui/src/pages/tabs/ServicesTab.tsx` (write call sites)
- Test: `test/TTMAccount.test.js`, `test/PartnerConfiguration.test.js`

**Interfaces:**

- Consumes: `ServiceRegistry`'s `getRegisteredServiceNameByHash(bytes32) → string` from `ITTMAccountManager` (used for the registration check in Step 4); Task 4's `useServiceCatalog` / `hashServiceName` for the UI step.
- Produces: `addService(bytes32 serviceHash, bool restrictedRate, string[] capabilities)`, `removeService(bytes32)`, `removeAllServices()`, `setServiceRestrictedRate(bytes32, bool)`, `addServiceCapability(bytes32, string)`, `removeServiceCapability(bytes32, string)`, `setServiceCapabilities(bytes32, string[])`. Events: `ServiceAdded(bytes32 indexed)`, `ServiceRemoved(bytes32 indexed)`, `ServiceRestrictedRateUpdated(bytes32 indexed, bool)`, `ServiceCapabilitiesUpdated(bytes32 indexed)`, `ServiceCapabilityAdded(bytes32 indexed, string)`, `ServiceCapabilityRemoved(bytes32 indexed, string)`.

> `contracts/manager/ITTMAccountManager.sol` **already declares** `getRegisteredServiceNameByHash(bytes32) → string` (verified 2026-07-21), so no interface change is needed here. Task 7 trims the now-unused declarations.

- [ ] **Step 1: Write the failing tests**

Add to `test/TTMAccount.test.js`:

```js
    describe("Hash-native services", function () {
        it("should add a service by hash and emit the hash", async function () {
            await setupSigners();
            const { ttmAccount, ttmAccountManager } = await loadFixture(deployAndConfigureAllFixture);

            const name = "ttm.services.accommodation.v1alpha.AccommodationSearchService";
            await ttmAccountManager.connect(signers.registryAdmin).registerService(name);
            const hash = ethers.keccak256(ethers.toUtf8Bytes(name));

            await expect(ttmAccount.connect(signers.ttmServiceAdmin).addService(hash, false, []))
                .to.emit(ttmAccount, "ServiceAdded")
                .withArgs(hash);

            expect(await ttmAccount.isServiceSupported(hash)).to.be.true;
            expect(await ttmAccount.getAllServiceHashes()).to.deep.equal([hash]);
        });

        it("should reject adding a service whose hash is not registered", async function () {
            await setupSigners();
            const { ttmAccount } = await loadFixture(deployAndConfigureAllFixture);

            const unregistered = ethers.keccak256(ethers.toUtf8Bytes("ttm.services.nope.v1.NopeService"));

            await expect(
                ttmAccount.connect(signers.ttmServiceAdmin).addService(unregistered, false, []),
            ).to.be.revertedWithCustomError(ttmAccount, "ServiceNotRegistered");
        });

        it("should remove every service without resolving names", async function () {
            await setupSigners();
            const { ttmAccount, ttmAccountManager } = await loadFixture(deployAndConfigureAllFixture);

            const names = [
                "ttm.services.accommodation.v1alpha.AccommodationSearchService",
                "ttm.services.activity.v2.ActivitySearchService",
            ];
            const hashes = [];
            for (const name of names) {
                await ttmAccountManager.connect(signers.registryAdmin).registerService(name);
                const hash = ethers.keccak256(ethers.toUtf8Bytes(name));
                hashes.push(hash);
                await ttmAccount.connect(signers.ttmServiceAdmin).addService(hash, false, []);
            }
            expect(await ttmAccount.getAllServiceHashes()).to.have.lengthOf(2);

            await ttmAccount.connect(signers.ttmServiceAdmin).removeAllServices();

            expect(await ttmAccount.getAllServiceHashes()).to.deep.equal([]);
            for (const hash of hashes) {
                expect(await ttmAccount.isServiceSupported(hash)).to.be.false;
            }
        });

        it("should manage capabilities by hash, keeping capability strings readable", async function () {
            await setupSigners();
            const { ttmAccount, ttmAccountManager } = await loadFixture(deployAndConfigureAllFixture);

            const name = "ttm.services.accommodation.v1alpha.AccommodationSearchService";
            await ttmAccountManager.connect(signers.registryAdmin).registerService(name);
            const hash = ethers.keccak256(ethers.toUtf8Bytes(name));
            await ttmAccount.connect(signers.ttmServiceAdmin).addService(hash, false, []);

            await expect(ttmAccount.connect(signers.ttmServiceAdmin).addServiceCapability(hash, "luggage"))
                .to.emit(ttmAccount, "ServiceCapabilityAdded")
                .withArgs(hash, "luggage");

            expect(await ttmAccount.getServiceCapabilities(hash)).to.deep.equal(["luggage"]);

            await expect(ttmAccount.connect(signers.ttmServiceAdmin).removeServiceCapability(hash, "luggage"))
                .to.emit(ttmAccount, "ServiceCapabilityRemoved")
                .withArgs(hash, "luggage");
        });
    });
```

- [ ] **Step 2: Run the tests to verify they fail**

Run: `yarn hardhat test test/TTMAccount.test.js --grep "Hash-native services"`
Expected: FAIL — `addService` currently takes a `string`, so ethers rejects a `bytes32` argument.

- [ ] **Step 3: Change the event declarations**

In `contracts/account/TTMAccount.sol`, replace the six service events at `:164-166,170-174`:

```solidity
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

    event ServiceRestrictedRateUpdated(bytes32 indexed serviceHash, bool restrictedRate);

    event ServiceCapabilitiesUpdated(bytes32 indexed serviceHash);
    event ServiceCapabilityAdded(bytes32 indexed serviceHash, string capability);
    event ServiceCapabilityRemoved(bytes32 indexed serviceHash, string capability);
```

Leave `WantedServiceAdded` / `WantedServiceRemoved` (`:167-168`) for Task 6.

- [ ] **Step 4: Rewrite the write functions**

Replace `addService`:

```solidity
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
```

Replace `removeService`, `removeAllServices`, and the capability/rate setters:

```solidity
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
```

> `removeAllServices` previously iterated `getSupportedServices()` and re-resolved each name back to a hash. Iterating `getAllServiceHashes()` directly removes two staticcalls per service. Note that `_removeService` mutates the underlying set while the memory array is a snapshot — that is correct and is why the array must be read before the loop.

- [ ] **Step 5: Run the tests to verify they pass**

Run: `yarn hardhat test test/TTMAccount.test.js --grep "Hash-native services"`
Expected: PASS.

- [ ] **Step 6: Fix every other caller**

Run: `yarn test`

Many existing tests pass service names as strings. Convert each to `ethers.keccak256(ethers.toUtf8Bytes(name))`. Add a helper at the top of `test/utils/fixtures.js` and export it:

```js
const serviceHash = (name) => ethers.keccak256(ethers.toUtf8Bytes(name));
```

Import and use it rather than repeating the expression. **Adapt every failing test; delete none.**

- [ ] **Step 7: Convert the UI's service writes to hashes**

`ui/src/pages/tabs/ServicesTab.tsx` currently passes service **names** to `addService`, `removeService`, `setServiceRestrictedRate`, `addServiceCapability` and `removeServiceCapability`. Each must now pass a hash.

Use the Task 4 resolver:

```ts
import { useServiceCatalog } from "../../hooks/useServiceCatalog";
import { hashServiceName } from "../../lib/serviceCatalog";
```

Inside the component, `const { catalog } = useServiceCatalog();`. For each write call site, replace the name argument with `catalog.hashByName.get(service.name) ?? hashServiceName(service.name)`.

The fallback matters: a service the account supports but the registry has since unregistered will be missing from `hashByName`, and hashing the name directly still yields the correct hash for a `remove` call. Extract it to one local helper rather than repeating the expression at five call sites:

```ts
const hashFor = (name: string) => catalog.hashByName.get(name) ?? hashServiceName(name);
```

**Capability strings are unchanged** — they stay plain text in both the ABI and the UI.

Read the file before editing; the call sites go through a `run(label, functionName, args)` wrapper (see `:224`, `:245`, `:259`), so the change is to the `args` array in each.

- [ ] **Step 8: Regenerate, verify, and commit**

```bash
yarn hardhat clean && yarn compile && yarn docgen && yarn hardhat export-abi
(cd ui && yarn sync)
yarn test && yarn lint
(cd ui && yarn test && yarn build)
git add -A
git commit -m "refactor(account)!: hash-native service CRUD and hash-only service events

BREAKING CHANGE: addService, removeService, setServiceRestrictedRate,
setServiceCapabilities, addServiceCapability and removeServiceCapability take
bytes32 service hashes. The six service events carry the hash, not the name."
```

Both the contract suite and the UI build must be green before this task is done.

---

### Task 6: Hash-native wanted services

The two events the technical backlog's §1.1 miscount missed.

**Files:**

- Modify: `contracts/account/TTMAccount.sol` (events `:167-168`, `addWantedServices`, `removeWantedServices`, `getWantedServices`)
- Modify: any UI wanted-service call sites found in Step 6
- Test: `test/TTMAccount.test.js`

**Interfaces:**

- Consumes: `_requireRegisteredService(bytes32)` from Task 5.
- Produces: `addWantedServices(bytes32[] serviceHashes)`, `removeWantedServices(bytes32[] serviceHashes)`, `event WantedServiceAdded(bytes32 indexed serviceHash)`, `event WantedServiceRemoved(bytes32 indexed serviceHash)`. `getWantedServices()` returning names is **deleted**; `getWantedServiceHashes()` already exists and remains.

- [ ] **Step 1: Write the failing test**

Add to `test/TTMAccount.test.js`, inside `describe("Hash-native services", ...)`:

```js
        it("should add and remove wanted services by hash", async function () {
            await setupSigners();
            const { ttmAccount, ttmAccountManager } = await loadFixture(deployAndConfigureAllFixture);

            const name = "ttm.services.activity.v2.ActivitySearchService";
            await ttmAccountManager.connect(signers.registryAdmin).registerService(name);
            const hash = ethers.keccak256(ethers.toUtf8Bytes(name));

            await expect(ttmAccount.connect(signers.ttmServiceAdmin).addWantedServices([hash]))
                .to.emit(ttmAccount, "WantedServiceAdded")
                .withArgs(hash);

            expect(await ttmAccount.getWantedServiceHashes()).to.deep.equal([hash]);

            await expect(ttmAccount.connect(signers.ttmServiceAdmin).removeWantedServices([hash]))
                .to.emit(ttmAccount, "WantedServiceRemoved")
                .withArgs(hash);

            expect(await ttmAccount.getWantedServiceHashes()).to.deep.equal([]);
        });
```

- [ ] **Step 2: Run it to verify it fails**

Run: `yarn hardhat test test/TTMAccount.test.js --grep "wanted services by hash"`
Expected: FAIL — `addWantedServices` takes `string[]`.

- [ ] **Step 3: Change the events**

```solidity
    event WantedServiceAdded(bytes32 indexed serviceHash);
    event WantedServiceRemoved(bytes32 indexed serviceHash);
```

- [ ] **Step 4: Rewrite the functions**

```solidity
    /**
     * @notice Declares services this account wants to consume from other partners.
     *
     * @dev Each hash must be registered in the manager's ServiceRegistry, for the same
     * reason as {addService}.
     *
     * @param serviceHashes Hashes of the service names to want
     */
    function addWantedServices(bytes32[] memory serviceHashes) public onlyRole(SERVICE_ADMIN_ROLE) {
        for (uint256 i = 0; i < serviceHashes.length; i++) {
            _requireRegisteredService(serviceHashes[i]);
            _addWantedService(serviceHashes[i]);
            emit WantedServiceAdded(serviceHashes[i]);
        }
    }

    /**
     * @notice Removes services from this account's wanted list.
     *
     * @param serviceHashes Hashes of the service names to stop wanting
     */
    function removeWantedServices(bytes32[] memory serviceHashes) public onlyRole(SERVICE_ADMIN_ROLE) {
        for (uint256 i = 0; i < serviceHashes.length; i++) {
            _removeWantedService(serviceHashes[i]);
            emit WantedServiceRemoved(serviceHashes[i]);
        }
    }
```

Delete `getWantedServices()` (the name-resolving variant). `getWantedServiceHashes()` already exists.

> Read the existing implementations first — the internal helper names (`_addWantedService`, `_removeWantedService`) are inferred from the `PartnerConfiguration` pattern. Use the actual ones.

- [ ] **Step 5: Run the full suite and fix callers**

Run: `yarn test`
Expected: PASS after converting the remaining string-based wanted-service tests to hashes using the `serviceHash` helper from Task 5.

> **State check after Task 5** (verified by review): `ui/src/lib/activity/catalog.test.ts:52-60` asserts the `serviceName`-keyed fallback arm of `serviceHashArg()` **using `WantedServiceAdded`** as its example event. Once you convert wanted services to `bytes32`, that arm has no real producer left. The test will still pass (the helper falls back), so it silently becomes a test of dead code rather than failing. Either repoint it at a genuine remaining producer or note it for Task 7, which collapses the fallback.

- [ ] **Step 6: Update any UI wanted-service call sites**

```bash
grep -rn "WantedService\|wantedService\|getWantedServices" ui/src
```

Convert each write to pass `bytes32[]` using the Task 4 resolver's `hashFor` pattern, and each read from `getWantedServices()` to `getWantedServiceHashes()` plus `serviceNameForHash` for display. **If the grep returns nothing, the UI does not surface wanted services — record that in the report and skip this step rather than inventing a call site.**

- [ ] **Step 7: Regenerate and commit**

```bash
yarn hardhat clean && yarn compile && yarn docgen && yarn hardhat export-abi
(cd ui && yarn sync)
yarn test && yarn lint
(cd ui && yarn test && yarn build)
git add -A
git commit -m "refactor(account)!: hash-native wanted services

BREAKING CHANGE: addWantedServices and removeWantedServices take bytes32[].
getWantedServices() is removed; use getWantedServiceHashes()."
```

---

### Task 7: Hash-native read surface and helper removal

Deletes the manager staticcalls from every read path and the ~200 lines of name-resolution scaffolding.

**Files:**

- Modify: `contracts/account/TTMAccount.sol` (`getRegisteredServiceHash`/`getServiceHash`/`getServiceName` `:527-551`, `getSupportedServices` `:553-565`, `isServiceSupported` `:576-578`, plus any remaining string-typed getters in `:427-635`)
- Modify: `contracts/manager/ITTMAccountManager.sol` (Step 5)
- Modify: `ui/src/hooks/useAccountActivity.ts:28-53`, `ui/src/lib/activity/catalog.ts:20-35,217-219`, `ui/src/pages/tabs/ServicesTab.tsx` (read call sites)
- Test: `test/TTMAccount.test.js`, `test/PartnerConfiguration.test.js`, `ui/src/lib/activity/catalog.test.ts`

**Interfaces:**

- Consumes: Task 5's hash-native writes; Task 4's `useServiceCatalog` and `serviceNameForHash`.
- Produces: `getSupportedServices() → (bytes32[] serviceHashes, Service[] services)`, `getSupportedServicesSlice(uint256 offset, uint256 limit) → (bytes32[], Service[])`, `isServiceSupported(bytes32) → bool`, `getServiceCapabilities(bytes32) → string[]`, `getServiceRestrictedRate(bytes32) → bool`. The `string`-typed overloads of all of these are **deleted**, as are the three private resolution helpers.

- [ ] **Step 1: Write the failing test**

```js
        it("should return supported services as hashes with no manager round-trip", async function () {
            await setupSigners();
            const { ttmAccount, ttmAccountManager } = await loadFixture(deployAndConfigureAllFixture);

            const names = [
                "ttm.services.accommodation.v1alpha.AccommodationSearchService",
                "ttm.services.activity.v2.ActivitySearchService",
            ];
            const hashes = [];
            for (const name of names) {
                await ttmAccountManager.connect(signers.registryAdmin).registerService(name);
                const hash = ethers.keccak256(ethers.toUtf8Bytes(name));
                hashes.push(hash);
                await ttmAccount.connect(signers.ttmServiceAdmin).addService(hash, true, ["luggage"]);
            }

            const [serviceHashes, services] = await ttmAccount.getSupportedServices();
            expect(serviceHashes).to.deep.equal(hashes);
            expect(services).to.have.lengthOf(2);
            expect(services[0][0]).to.be.true; // _restrictedRate

            const [pageHashes] = await ttmAccount.getSupportedServicesSlice(1, 5);
            expect(pageHashes).to.deep.equal([hashes[1]]);

            // The string-typed API is gone.
            expect(ttmAccount.interface.fragments.some((f) => f.name === "getServiceHash")).to.be.false;
        });
```

- [ ] **Step 2: Run it to verify it fails**

Run: `yarn hardhat test test/TTMAccount.test.js --grep "no manager round-trip"`
Expected: FAIL — `getSupportedServicesSlice` does not exist.

- [ ] **Step 3: Delete the resolution helpers**

Delete `getRegisteredServiceHash`, `getServiceHash`, and `getServiceName` (`:527-551`) entirely, along with the `SERVICES WITH RESOLVED NAMES` section header comment. `_requireRegisteredService` from Task 5 is the only remaining manager call.

- [ ] **Step 4: Rewrite the read surface**

```solidity
    /**
     * @notice Returns every supported service as a hash plus its stored record.
     *
     * @dev Reads no longer touch the manager. Resolve hashes to names client-side from
     * the registry's `ServiceRegistered` events or `getAllRegisteredServiceNames()`.
     * Unbounded - prefer {getSupportedServicesSlice} against a public RPC.
     */
    function getSupportedServices() public view returns (bytes32[] memory serviceHashes, Service[] memory services) {
        serviceHashes = getAllServiceHashes();
        services = new Service[](serviceHashes.length);

        for (uint256 i = 0; i < serviceHashes.length; i++) {
            services[i] = getService(serviceHashes[i]);
        }
    }

    /**
     * @notice Returns a bounded window of supported services.
     *
     * Returns empty arrays if `offset` is at or past the end; the window is clamped to
     * the end of the list, so an oversized `limit` is not an error.
     *
     * @param offset Index to start at
     * @param limit Maximum number of services to return
     */
    function getSupportedServicesSlice(
        uint256 offset,
        uint256 limit
    ) public view returns (bytes32[] memory serviceHashes, Service[] memory services) {
        bytes32[] memory allHashes = getAllServiceHashes();
        uint256 total = allHashes.length;
        if (offset >= total) {
            return (new bytes32[](0), new Service[](0));
        }

        // Clamp by subtraction, not by computing `offset + limit`: under checked
        // arithmetic that sum reverts for a large `limit`, which would contradict
        // the "an oversized limit is not an error" contract above. `offset < total`
        // here, so `total - offset` cannot underflow.
        uint256 remaining = total - offset;
        if (limit > remaining) {
            limit = remaining;
        }

        serviceHashes = new bytes32[](limit);
        services = new Service[](limit);
        for (uint256 i = 0; i < limit; i++) {
            serviceHashes[i] = allHashes[offset + i];
            services[i] = getService(allHashes[offset + i]);
        }
    }

    /**
     * @notice Checks whether a service is supported by this account.
     *
     * @param serviceHash Hash of the service name to check
     */
    function isServiceSupported(bytes32 serviceHash) public view returns (bool) {
        return _isServiceSupported(serviceHash);
    }
```

Delete the `string`-typed overloads of `getServiceCapabilities` and `getServiceRestrictedRate` (the `bytes32` versions already exist per the current ABI), and any other string-typed getter left in `:427-635`.

- [ ] **Step 5: Trim the now-unused manager interface declarations**

With the resolution helpers gone, `TTMAccount` calls only `getRegisteredServiceNameByHash` on the manager. Delete the three declarations nothing calls any more from `contracts/manager/ITTMAccountManager.sol`:

```solidity
// SPDX-License-Identifier: LGPL-3.0-or-later

pragma solidity 0.8.24;

interface ITTMAccountManager {
    function getAccountImplementation() external view returns (address);

    function isTTMAccount(address account) external view returns (bool);

    /// @dev Reverts if the hash is not registered. This is the sole remaining
    /// manager dependency of TTMAccount, used to validate {addService}.
    function getRegisteredServiceNameByHash(bytes32 serviceHash) external view returns (string memory serviceName);
}
```

The removed declarations (`getRegisteredServiceHashByName`, `getServiceHashByName`, `getServiceNameByHash`) remain **public on `TTMAccountManager` itself** — they are declared by `ServiceRegistry` and are what the UI and bot call to seed their catalogs. Only the interface's view of them narrows.

Run: `yarn compile`
Expected: compiles clean. A failure here means something still calls one of the three through the interface — find it before deleting.

- [ ] **Step 6: Run the full suite**

Run: `yarn test`
Expected: PASS after adapting the remaining string-based tests. Expected total: roughly 148 passing.

- [ ] **Step 7: Measure the size win**

Run: `yarn compile 2>&1 | grep -E "TTMAccount |BookingToken |TTMAccountManager "`
Record the numbers. `TTMAccount` should have dropped meaningfully from 21.371 KiB. **If it did not drop, stop and investigate before continuing** — the change did not do what it was supposed to.

- [ ] **Step 8: Rewrite the activity feed's hash resolution**

This is where the old per-hash round-trip finally goes away.

In `ui/src/hooks/useAccountActivity.ts`, delete the `serviceHashes` memo, the `useReadContracts` block that calls `getServiceNameByHash` per hash (confirmed still present at `:43-49` after Task 5), and the `nameByHash` memo. Line numbers shifted by about +4 in Task 5 — locate by content. The file's docblock (`:9-19`) now describes the transitional two-shape situation and must be rewritten too, not just the code. Replace with the Task 4 resolver. Note the event arg is now named `serviceHash`, not `serviceName`:

```ts
  const { catalog } = useServiceCatalog();

  const events = useMemo(
    () =>
      activity.events.map((e) => {
        const hash = e.args.serviceHash;
        const serviceLabel = typeof hash === "string" ? serviceNameForHash(catalog, hash) : undefined;
        const sentence = serviceLabel ? renderSentence(e.source, e.eventName, { ...e.args, serviceLabel }) : e.sentence;
        return { ...e, timestamp: timestamps.get(e.blockNumber), sentence };
      }),
    [activity.events, timestamps, catalog],
  );
```

In `ui/src/lib/activity/catalog.ts`:

- Delete the `SERVICE_HASH_EVENTS` export (now at `:229-232`) — nothing needs to special-case which events carry hashes any more. Its consumers are `useAccountActivity.ts:7,35` and `catalog.test.ts:6,92-94`.
- **Do NOT blindly delete `serviceHashArg(args)` (`:231-233`, `args.serviceHash ?? args.serviceName`) alongside it.** Task 5 added this helper and it has **three** consumers, not one: `catalog.ts:239` inside `serviceLabel()`, plus `useAccountActivity.ts:35` and `:68`. Deleting the `useReadContracts` batch removes two of the three; the `serviceLabel()` use must survive or be replaced.
- By this point Task 6 has converted wanted services to `bytes32`, so the `?? args.serviceName` fallback arm is **dead**. Collapse `serviceHashArg` to a plain `args.serviceHash` read, or inline it — but verify no remaining event still carries a `serviceName` arg first (`ServiceRegistered`/`ServiceUnregistered` on the *manager* still do, deliberately; they are a different code path, so check which source each consumer handles).
- Update `serviceLabel()` to read the hash arg rather than `args.serviceName`.
- Update the header comment. It currently explains that indexed `string` params arrive as a keccak hash. That is no longer why hashes appear — the contracts now emit `bytes32` deliberately, and names come from the registry's own events. Also note `catalog.ts:228`'s comment still says "Account events whose indexed `serviceName` is a hash", which is stale.

Update `ui/src/lib/activity/catalog.test.ts` to match — it is **11 tests** after Task 5, and `:92-94` holds the `SERVICE_HASH_EVENTS` assertions that must go with the export. Re-evaluate `:52-60` (see the Task 6 note): after Task 6 its example event no longer produces the shape it asserts.

- [ ] **Step 9: Convert the UI's service reads**

`ServicesTab.tsx` reads `getSupportedServices()`, which now returns hashes. Map each hash through `serviceNameForHash(catalog, hash)` for display, falling back to a shortened hash when the name is unknown (a service unregistered after the account added it). Keep the existing `groupServicesByPackage` / `parseServiceName` grouping from `ui/src/lib/serviceName.ts` — it takes names, so it operates on the resolved values.

- [ ] **Step 10: Verify the whole UI is clean**

```bash
cd ui && yarn test && yarn build && yarn lint:format
grep -rn "SERVICE_HASH_EVENTS\|getServiceNameByHash\|TTMACCOUNT_ROLE" src
```

Expected: tests and build pass; the grep returns nothing.

- [ ] **Step 11: Regenerate and commit**

```bash
yarn hardhat clean && yarn compile && yarn docgen && yarn hardhat export-abi
(cd ui && yarn sync)
yarn test && yarn lint
(cd ui && yarn test && yarn build)
git add -A
git commit -m "refactor(account)!: hash-native reads, drop name-resolution helpers

Reads no longer make a cross-contract staticcall per service. getSupportedServices
previously cost two manager staticcalls per supported service.

BREAKING CHANGE: getSupportedServices returns bytes32 hashes. isServiceSupported,
getServiceCapabilities and getServiceRestrictedRate take bytes32 only. The
string-typed overloads and getServiceHash/getServiceName are removed."
```

---

### Task 8: recordExpiration posture and supportsInterface

Two small independent corrections.

**Files:**

- Modify: `contracts/account/TTMAccount.sol` (`recordExpiration` `:374-376`, `supportsInterface` around `:383`)
- Test: `test/TTMAccount.test.js`

**Interfaces:**

- Consumes: nothing.
- Produces: `recordExpiration(uint256)` with no role gate; `supportsInterface(bytes4)` reporting `IERC721Receiver`.

- [ ] **Step 1: Write the failing tests**

```js
    it("should let anyone record an expiration", async function () {
        await setupSigners();
        const { ttmAccount } = await loadFixture(deployAndConfigureAllFixture);

        const fragment = ttmAccount.interface.getFunction("recordExpiration");
        expect(fragment).to.not.be.undefined;

        // A caller with no roles must not be rejected by access control. The call
        // still reverts on token state, which is the point: the only gate is
        // whether the reservation has genuinely expired.
        await expect(
            ttmAccount.connect(signers.otherAccount3).recordExpiration(999999),
        ).to.not.be.revertedWithCustomError(ttmAccount, "AccessControlUnauthorizedAccount");
    });

    it("should report IERC721Receiver support", async function () {
        await setupSigners();
        const { ttmAccount } = await loadFixture(deployAndConfigureAllFixture);

        // bytes4(keccak256("onERC721Received(address,address,uint256,bytes)"))
        expect(await ttmAccount.supportsInterface("0x150b7a02")).to.be.true;
    });
```

- [ ] **Step 2: Run to verify they fail**

Run: `yarn hardhat test test/TTMAccount.test.js --grep "record an expiration|IERC721Receiver"`
Expected: FAIL — the first on `AccessControlUnauthorizedAccount`, the second returning `false`.

- [ ] **Step 3: Drop the role gate**

```solidity
    /**
     * @notice Marks an expired reservation as expired on the BookingToken.
     *
     * @dev Deliberately permissionless. The underlying `BookingToken.recordExpiration`
     * is public and unrestricted, so a role gate here would protect nothing - it only
     * created the false impression that one was needed. The operation is objective
     * housekeeping: it succeeds only once `block.timestamp` has genuinely passed the
     * reservation's expiry, so there is nothing for an attacker to gain.
     *
     * @param tokenId The booking token to mark expired
     */
    function recordExpiration(uint256 tokenId) public {
        IBookingToken(getBookingTokenAddress()).recordExpiration(tokenId);
    }
```

> Preserve the existing body exactly; only the modifier and NatSpec change.

- [ ] **Step 4: Override supportsInterface**

```solidity
    /**
     * @inheritdoc IERC165
     *
     * @dev This contract implements {IERC721Receiver}, so it must say so - counterparties
     * that capability-detect before transferring an NFT would otherwise conclude it
     * cannot receive one.
     */
    function supportsInterface(
        bytes4 interfaceId
    ) public view virtual override(AccessControlEnumerableUpgradeable) returns (bool) {
        return interfaceId == type(IERC721Receiver).interfaceId || super.supportsInterface(interfaceId);
    }
```

> The `override(...)` list must name every base that declares `supportsInterface`. Check the inheritance list at the top of `TTMAccount.sol` and adjust; the compiler will tell you precisely which bases are required.

- [ ] **Step 5: Run, regenerate, commit**

```bash
yarn test
yarn hardhat clean && yarn compile && yarn docgen && yarn hardhat export-abi
(cd ui && yarn sync)
yarn lint
git add -A
git commit -m "fix(account): drop the misleading recordExpiration gate, report IERC721Receiver

BREAKING CHANGE: recordExpiration no longer requires BOOKING_OPERATOR_ROLE. The
gate protected nothing - BookingToken.recordExpiration is public and correctly so."
```

---

### Task 9: Collapse the cancellation wrappers

Six near-identical wrappers, each repeating the same three checks.

**Files:**

- Modify: `contracts/booking-token/BookingToken.sol:721-850`
- Test: `test/BookingToken.test.js` (existing coverage; no new tests — this is ABI-identical)

**Interfaces:**

- Consumes: nothing.
- Produces: private `_requireBoughtAndParties(uint256 tokenId) returns (address owner, address supplier)`. No ABI change.

- [ ] **Step 1: Confirm the baseline is green and record the size**

```bash
yarn test
yarn compile 2>&1 | grep "BookingToken "
```

Record both. This task must not change test results at all.

- [ ] **Step 2: Add the shared helper**

```solidity
    /**
     * @notice Requires that `tokenId` exists and is in the BOUGHT state, returning both
     * parties to the booking.
     *
     * @dev Extracted from the six cancellation wrappers, which each repeated this
     * sequence verbatim.
     *
     * @param tokenId The booking token
     * @return owner The current owner (the buyer)
     * @return supplier The supplier that minted the reservation
     */
    function _requireBoughtAndParties(uint256 tokenId) private view returns (address owner, address supplier) {
        owner = _requireOwned(tokenId);

        BookingTokenStorage storage $ = _getBookingTokenStorage();
        if ($._bookingStatus[tokenId] != BookingStatus.BOUGHT) {
            revert TokenIsNotBought(tokenId);
        }

        supplier = $._reservations[tokenId].supplier;
    }
```

> The storage field names and the revert's custom error are inferred from the wrappers. Read `:721-850` first and copy the **exact** error and field names each wrapper uses. If the six wrappers do not all revert with the same error, they are not as identical as the backlog claimed — in that case, extract only what genuinely matches and say so in the commit message.

- [ ] **Step 3: Rewrite each of the six wrappers**

Replace each wrapper's preamble with a single call. For example:

```solidity
    function initiateCancellation(
        uint256 tokenId,
        uint256 refundAmount,
        uint16 cancellationReason,
        uint16 cancellationReasonVersion
    ) external {
        (address owner, address supplier) = _requireBoughtAndParties(tokenId);
        _initiateCancellation(tokenId, owner, supplier, refundAmount, cancellationReason, cancellationReasonVersion);
    }
```

Apply the same shape to the other five. Keep each wrapper's existing delegation call and its argument order exactly as they are.

- [ ] **Step 4: Verify nothing changed behaviourally**

```bash
yarn test
```

Expected: identical pass count to Step 1. Any change in test results means the refactor altered behaviour — revert and redo.

- [ ] **Step 5: Add nonReentrant to finalizeCancellation**

Add the `nonReentrant` modifier to `finalizeCancellation`, matching `buyReservedToken`.

Run: `yarn test`
Expected: still green.

- [ ] **Step 6: Measure, regenerate, commit**

```bash
yarn compile 2>&1 | grep "BookingToken "
```

Record the new size against Step 1's.

```bash
yarn hardhat clean && yarn compile && yarn docgen && yarn hardhat export-abi
(cd ui && yarn sync)
yarn lint
git add -A
git commit -m "refactor(booking-token): collapse the six cancellation wrappers

ABI-identical. Also adds nonReentrant to finalizeCancellation for consistency
with buyReservedToken."
```

---

### Task 10: Remaining test coverage

The gaps the fee removal left, headed by the one that pins an accepted risk.

**Files:**

- Test: `test/TTMAccountManager.test.js`, `test/BookingToken.test.js`
- Create: `contracts/test/RejectsEther.sol`

**Interfaces:**

- Consumes: everything above.
- Produces: `contracts/test/RejectsEther.sol` — a contract with no `receive`/`fallback`, for the refund test.

- [ ] **Step 1: Write the permissionless-creation test**

This is the highest-value test in the plan. It pins the accepted risk from Decision 1: Camino enforced KYC at the chain level and that guarantee vanished in the move to Base **without any code changing**. The day Decision 1 is implemented, this test must fail loudly rather than the change passing silently.

Add to `test/TTMAccountManager.test.js`:

```js
        it("should let any address create a TTM Account (deliberately permissionless)", async function () {
            await setupSigners();
            const { ttmAccountManager } = await loadFixture(deployTTMAccountManagerWithTTMAccountImplFixture);

            // signers.otherAccount3 holds no roles at all.
            await expect(
                ttmAccountManager
                    .connect(signers.otherAccount3)
                    .createTTMAccount(signers.otherAccount3.address, signers.otherAccount3.address),
            ).to.not.be.reverted;

            // IF THIS TEST FAILS, READ THIS BEFORE "FIXING" IT.
            //
            // Account creation being open to anyone is a deliberate, documented
            // decision for testnet - see docs/decisions/2026-07-21-contract-design-decisions.md,
            // Decision 1. Camino enforced KYC at the chain level; moving to Base
            // removed that guarantee without any code change.
            //
            // A failure here means someone added an access control gate. If that
            // was intentional (Decision 1 was resolved), update this test to assert
            // the new gate. If it was not intentional, the gate is the bug.
        });
```

- [ ] **Step 2: Run it**

Run: `yarn hardhat test test/TTMAccountManager.test.js --grep "deliberately permissionless"`
Expected: PASS immediately. That is correct — it is a characterisation test pinning current behaviour, not a red-green cycle. Verify it is meaningful by temporarily adding `onlyRole(DEFAULT_ADMIN_ROLE)` to `createTTMAccount`, confirming the test fails, then reverting.

- [ ] **Step 3: Write the setter-event tests**

```js
        it("should emit BookingTokenAddressUpdated and TTMAccountImplementationUpdated", async function () {
            await setupSigners();
            const { ttmAccountManager, bookingToken, ttmAccountImpl } = await loadFixture(
                deployAndConfigureAllFixture,
            );

            const oldToken = await ttmAccountManager.getBookingTokenAddress();
            const newToken = await bookingToken.getAddress();
            await expect(ttmAccountManager.connect(signers.managerVersioner).setBookingTokenAddress(newToken))
                .to.emit(ttmAccountManager, "BookingTokenAddressUpdated")
                .withArgs(oldToken, newToken);

            const oldImpl = await ttmAccountManager.getAccountImplementation();
            const newImpl = await ttmAccountImpl.getAddress();
            await expect(ttmAccountManager.connect(signers.managerVersioner).setAccountImplementation(newImpl))
                .to.emit(ttmAccountManager, "TTMAccountImplementationUpdated")
                .withArgs(oldImpl, newImpl);
        });
```

> If setting the same address twice is rejected or is a no-op, deploy a second implementation in the test rather than reusing the current one.

- [ ] **Step 4: Add the ETH-rejecting test contract**

Create `contracts/test/RejectsEther.sol`:

```solidity
// SPDX-License-Identifier: LGPL-3.0-or-later
pragma solidity 0.8.24;

/**
 * @notice A contract that cannot receive ETH: no `receive`, no `fallback`.
 *
 * @dev Used to exercise the cancellation refund path against a counterparty whose
 * address rejects a plain transfer. See Decision 3 in
 * docs/decisions/2026-07-21-contract-design-decisions.md.
 */
contract RejectsEther {
    /// @notice Lets tests confirm the contract deployed and has code.
    function ping() external pure returns (bool) {
        return true;
    }
}
```

- [ ] **Step 5: Write the stuck-refund test**

Add to `test/BookingToken.test.js` a test that drives a booking to `BOUGHT` with a `RejectsEther` instance as the party receiving the refund, then finalizes cancellation and asserts the documented behaviour.

**Assert what actually happens, not what you think should happen.** Run it first and observe: either the call reverts (refund blocked) or the ETH is stranded. Whichever it is, write the assertion to match and add a comment referencing Decision 3. This test documents a known-open decision; it is not fixing it.

- [ ] **Step 6: Run the full suite**

Run: `yarn test`
Expected: all green, roughly 153 passing.

- [ ] **Step 7: Regenerate and commit**

```bash
yarn hardhat clean && yarn compile && yarn docgen && yarn hardhat export-abi
yarn lint
git add -A
git commit -m "test: backfill coverage the fee removal left behind

Pins createTTMAccount being permissionless, asserts the two setter events, and
documents the cancellation refund path against a contract that rejects ETH."
```

---

### Task 11: Regenerate Go bindings and record final sizes

Done once, at the end, because the generator rebuilds `node_modules` from scratch.

**Files:**

- Modify: everything under `go/contracts/`
- Modify: `docs/superpowers/specs/2026-07-21-hash-native-rework-design.md` (section F measured sizes)

**Interfaces:**

- Consumes: the final ABI from Tasks 1–10.
- Produces: regenerated Go bindings; the measured size table used by Task 12.

- [ ] **Step 1: Route caches off the full root filesystem**

`/` is ~97% full and `scripts/generate_go_abi.sh` does its own `rm -rf node_modules && yarn install`:

```bash
export TMPDIR=/hgst/tmp GOTMPDIR=/hgst/tmp
export GOCACHE=/hgst/tmp/gocache GOMODCACHE=/hgst/tmp/gomodcache
export YARN_CACHE_FOLDER=/hgst/tmp/yarn-cache
mkdir -p /hgst/tmp/gocache /hgst/tmp/gomodcache /hgst/tmp/yarn-cache
df -h / /hgst
```

- [ ] **Step 2: Regenerate**

```bash
./scripts/generate_go_abi.sh
```

- [ ] **Step 3: Check for orphaned packages**

```bash
ls go/contracts/
grep -n "ARTIFACTS" -A20 scripts/generate_go_abi.sh
```

Every directory under `go/contracts/` must appear in the generator's `ARTIFACTS` list. An orphan is invisible to both regeneration and CI's drift check — this previously left a stale `go/contracts/servicefeetoken/` that had to be deleted by hand. No contract is removed by this plan, so expect no orphans; if you find one, delete it and say so.

- [ ] **Step 4: Verify the bindings compile**

```bash
cd go && go build ./... && go vet ./...
```

- [ ] **Step 5: Record final sizes**

```bash
yarn hardhat clean && yarn compile 2>&1 | grep -E "TTMAccount |TTMAccountManager |BookingToken "
```

Update section F of the spec with a measured before/after table against the 22.5 KiB gate. Baseline (init KiB): `TTMAccountManager` 12.800, `TTMAccount` 21.371, `BookingToken` 21.552. **Report the real numbers even if a contract grew.** Then:

```bash
npx prettier --write docs/superpowers/specs/2026-07-21-hash-native-rework-design.md
```

- [ ] **Step 6: Full verification**

```bash
yarn test && yarn lint
(cd ui && yarn test && yarn build)
yarn hardhat clean && yarn compile && yarn docgen && yarn hardhat export-abi
(cd ui && yarn sync)
git status --porcelain
```

Expected: all green and `git status --porcelain` **empty** after committing. A non-empty tree here is exactly what CI's drift check catches.

- [ ] **Step 7: Commit**

```bash
git add -A
git commit -m "chore: regenerate Go bindings and record measured contract sizes"
```

---

### Task 12: Write the bot migration guide

The handoff for a fresh session in another repo. Written now, from the merged ABI, because a guide predicted from a spec drifts during implementation and a fresh session cannot tell a stale instruction from a live one.

**Files:**

- Create: `../BOT-MIGRATION.md` (the workspace parent folder — **not** inside any repo)

**Interfaces:**

- Consumes: the final ABI and Go bindings from Task 11.

- [ ] **Step 1: Confirm the destination and the convention**

```bash
ls /hgst/work/github.com/TravelTokenMarketplace/travel-token-messenger/
```

Expected: `REBRANDING.md`, `TODOS.md`, `CONTRACTS-NEXT.md` and the four repo directories. `BOT-MIGRATION.md` joins them. **It must never be committed to any repo** — it references the multi-repo layout.

There is deliberately **no** `APP-SERVICE-MIGRATION.md`. `travel-token-matrix-app-service/go.mod` requires the **bot** module, not the contracts module, and `go-ethereum` is only an indirect dependency pulled through the bot. It has zero contract call sites. Its whole involvement is a dependency bump, which is why that is a step *inside* the bot guide.

- [ ] **Step 2: Extract the exact signatures**

```bash
grep -o '"name":"[a-zA-Z]*Service[a-zA-Z]*"[^}]*' abi/contracts/account/TTMAccount.sol/TTMAccount.json | head -40
grep -n "func.*Ttmaccount.*Service" go/contracts/ttmaccount/*.go | head -30
```

Use the real generated Go names — do not guess Go binding capitalisation.

- [ ] **Step 3: Write the guide**

It must contain, with no placeholders:

1. **What changed and why**, in three sentences, plus a pointer to `docs/superpowers/specs/2026-07-21-hash-native-rework-design.md` in the contracts repo.
2. **The contracts Go module version to pin**, as an exact `go get` line.
3. **A call-site checklist**, each with the old code and the new code:
   - `pkg/ttm_accounts/ttm_accounts.go:189` — `IsServiceSupported` takes a hash. Hash the service name **once at startup** and cache it; this runs per inbound message, and hash-native is what removes the cross-contract staticcall from that hot path. Show the `crypto.Keccak256Hash([]byte(name))` call.
   - `internal/messaging/service_registry.go:38` — `GetSupportedServices` returns hashes. Seed a local resolver from the manager's `getAllRegisteredServiceNames()` and derive hashes locally. This is where the bot's `map[message.Type]rpc.Service` gets built, so the name is still needed — it just comes from the registry now.
   - `internal/eventlistener/subscriber/subscriber.go:127` — `WatchServiceAdded`'s event struct field is now the hash. Note that nothing outside the subscriber currently calls `SubscribeServiceAdded`, so this is plumbing to keep correct rather than a live break.
   - `tests/e2e/blockchain/client.go:163,191` — `AddService` and the account-creation helper take hashes; `TTMAccountCreated` now has three fields.
4. **Identity changes:** `TTMACCOUNT_ROLE` no longer exists. Anything enumerating it uses `getTTMAccounts()` / `isTTMAccount()`. `GetRoleMembers` calls for *other* roles (e.g. `pkg/ttm_accounts/ttm_accounts.go:138` for `MESSENGER_BOT_ROLE`) are unaffected — say so explicitly, so a reader does not over-apply the change.
5. **`recordExpiration` no longer needs `BOOKING_OPERATOR_ROLE`** (`pkg/ttm_accounts/ttm_accounts.go:304`). No code change required; note it so nobody re-adds a grant.
6. **A verification section:** `go build ./... && go vet ./... && go test ./...`, plus the e2e suite.
7. **A closing step:** bump `travel-token-matrix-app-service`'s `travel-token-messenger-bot/v13` requirement to the migrated bot version and confirm its build is green. **This is the app-service's entire involvement** — recording it here is what stops it being forgotten.

- [ ] **Step 4: Verify every claim in the guide**

For each file:line cited, open it and confirm the line still says what the guide claims. Line numbers in this plan were taken on 2026-07-21 and the bot repo may have moved.

- [ ] **Step 5: Confirm it is not tracked by git**

```bash
cd /hgst/work/github.com/TravelTokenMarketplace/travel-token-messenger/travel-token-messenger-contracts
git status --porcelain
```

Expected: empty. `BOT-MIGRATION.md` lives outside the repo, so it must not appear.

---

## Final verification

Before opening the PR:

```bash
yarn test && yarn lint
(cd ui && yarn test && yarn build && yarn lint:format)
(cd go && go build ./... && go vet ./...)
yarn hardhat clean && yarn compile && yarn docgen && yarn hardhat export-abi
(cd ui && yarn sync)
git status --porcelain   # must be empty
grep -rn "TTMACCOUNT_ROLE" contracts/ test/ tasks/ ui/src   # must be empty
```

Report the measured size table against the 22.5 KiB gate and the final test count against the 134 baseline. Do not claim completion without pasting the actual command output.
