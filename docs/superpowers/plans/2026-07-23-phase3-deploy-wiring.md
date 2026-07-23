# Phase 3 Deploy-Wiring Implementation Plan

> **For agentic workers:** REQUIRED SUB-SKILL: Use superpowers:subagent-driven-development (recommended) or superpowers:executing-plans to implement this plan task-by-task. Steps use checkbox (`- [ ]`) syntax for tracking.

**Goal:** Add a verifiable `roles:handoff` Hardhat task that moves all privileged roles from the deployer to a Safe (with `PAUSER` also on a dedicated hot key), and document it in the deploy runbook — no Solidity change, no deploy.

**Architecture:** A pure logic module (`tasks/lib/handoff.js`) does grant → verify-gate → renounce against already-instantiated contract objects, so it is unit-testable against a local deployment. A thin Hardhat task (`tasks/roles.js`) resolves the deployed proxies and deployer signer for the network and calls it. The verify gate refuses to renounce anything until the Safe provably holds every target role — the guard against bricking a singleton.

**Tech Stack:** Hardhat 2, ethers v6 (via `@nomicfoundation/hardhat-toolbox`), `@openzeppelin/hardhat-upgrades` (UUPS), Chai, Mocha.

## Global Constraints

- Solidity is untouched this phase — do **not** edit any `contracts/**` file, and therefore `abi/`, `docs/` (docgen), and `go/` bindings must not change.
- Contracts use OpenZeppelin `AccessControlEnumerable`; `renounceRole(role, account)` requires `account == msg.sender` (self-renounce only).
- Role getters are public constants: `contract.DEFAULT_ADMIN_ROLE()`, `contract.PAUSER_ROLE()`, etc. `DEFAULT_ADMIN_ROLE` is `bytes32(0)`.
- Deployed proxy keys in `deployed_addresses.json`: `TravelTokenMessengerModule#ManagerProxy`, `TravelTokenMessengerModule#BookingTokenProxy`.
- Chain dirs: `localhost`→`chain-31337`, `base_sepolia`→`chain-84532`, `base`→`chain-8453`.
- Manager roles: `DEFAULT_ADMIN_ROLE`, `PAUSER_ROLE`, `UPGRADER_ROLE`, `VERSIONER_ROLE`, `SERVICE_REGISTRY_ADMIN_ROLE`. BookingToken roles: `DEFAULT_ADMIN_ROLE`, `UPGRADER_ROLE`, `PAUSER_ROLE`, `MIN_EXPIRATION_ADMIN_ROLE`.
- Steady-state topology (from the spec): Safe holds every manager role in that list and BookingToken `DEFAULT_ADMIN`/`UPGRADER`/`PAUSER`; the hot pauser holds `PAUSER_ROLE` on both; deployer holds nothing afterward (or only `DEFAULT_ADMIN_ROLE` when `--keep-deployer-as-default-admin`).
- `MIN_EXPIRATION_ADMIN_ROLE` is optional: grant to the Safe and renounce from the deployer **only if the deployer holds it**.
- Lint: `yarn lint` (Prettier + ESLint) must pass on new `.js`; anything under `docs/` must be prettier-formatted.
- Follow existing task style in `tasks/manager.js`: `scope(...)`, `console.log` progress, `await tx.wait()`.

---

### Task 1: Handoff core logic + tests

**Files:**
- Create: `tasks/lib/handoff.js`
- Test: `test/handoff.test.js`

**Interfaces:**
- Consumes: nothing (leaf module). Uses ethers Contract objects and a Signer passed in by the caller.
- Produces:
  - `MANAGER_SAFE_ROLES: string[]`, `BOOKINGTOKEN_SAFE_ROLES: string[]` — role-name arrays granted to the Safe.
  - `async handoffRoles({ manager, bookingToken, deployer, safe, pauser, keepDeployerAsDefaultAdmin = false, log = console.log }) -> { manager: {...}, bookingToken: {...} }` — performs grant → verify → renounce and returns the final membership summary. Throws `Error` (message begins `Verify gate failed:`) **before any renounce** if the Safe/pauser do not hold every target role. `manager`/`bookingToken` are ethers Contract instances; `deployer`/`safe`/`pauser` — `deployer` is a Signer, `safe` and `pauser` are address strings.

- [ ] **Step 1: Write the failing test**

Create `test/handoff.test.js`:

```javascript
const { loadFixture } = require("@nomicfoundation/hardhat-toolbox/network-helpers");
const { expect } = require("chai");
const { ethers, upgrades } = require("hardhat");

const { handoffRoles, MANAGER_SAFE_ROLES, BOOKINGTOKEN_SAFE_ROLES } = require("../tasks/lib/handoff");

// Deploy manager + BookingToken with EVERY role on a single deployer signer,
// matching Approach H (parameters.json is `{}`, so all roles land on account 0).
async function deployAllOnDeployerFixture() {
    const [deployer, safe, pauser, other] = await ethers.getSigners();

    const TTMAccountManager = await ethers.getContractFactory("TTMAccountManager");
    const manager = await upgrades.deployProxy(
        TTMAccountManager,
        [deployer.address, deployer.address, deployer.address, deployer.address],
        { kind: "uups" },
    );
    await manager.waitForDeployment();

    const BookingToken = await ethers.getContractFactory("BookingToken");
    const bookingToken = await upgrades.deployProxy(
        BookingToken,
        [await manager.getAddress(), deployer.address, deployer.address],
        { kind: "uups" },
    );
    await bookingToken.waitForDeployment();

    // The deployer also picks up SERVICE_REGISTRY_ADMIN_ROLE during setup.
    await manager.grantRole(await manager.SERVICE_REGISTRY_ADMIN_ROLE(), deployer.address);

    return { manager, bookingToken, deployer, safe, pauser, other };
}

async function has(contract, roleName, address) {
    return contract.hasRole(await contract[roleName](), address);
}

describe("handoffRoles", function () {
    const quiet = () => {};

    it("grants the full topology to the Safe and hot pauser, and de-privileges the deployer", async function () {
        const { manager, bookingToken, deployer, safe, pauser } = await loadFixture(deployAllOnDeployerFixture);

        await handoffRoles({
            manager,
            bookingToken,
            deployer,
            safe: safe.address,
            pauser: pauser.address,
            log: quiet,
        });

        for (const role of MANAGER_SAFE_ROLES) {
            expect(await has(manager, role, safe.address), `manager ${role} -> safe`).to.be.true;
            expect(await has(manager, role, deployer.address), `manager ${role} deployer renounced`).to.be.false;
        }
        for (const role of BOOKINGTOKEN_SAFE_ROLES) {
            expect(await has(bookingToken, role, safe.address), `bt ${role} -> safe`).to.be.true;
            expect(await has(bookingToken, role, deployer.address), `bt ${role} deployer renounced`).to.be.false;
        }
        expect(await has(manager, "PAUSER_ROLE", pauser.address), "manager pauser hot key").to.be.true;
        expect(await has(bookingToken, "PAUSER_ROLE", pauser.address), "bt pauser hot key").to.be.true;
    });

    it("keeps the deployer as DEFAULT_ADMIN when the flag is set, renouncing everything else", async function () {
        const { manager, bookingToken, deployer, safe, pauser } = await loadFixture(deployAllOnDeployerFixture);

        await handoffRoles({
            manager,
            bookingToken,
            deployer,
            safe: safe.address,
            pauser: pauser.address,
            keepDeployerAsDefaultAdmin: true,
            log: quiet,
        });

        expect(await has(manager, "DEFAULT_ADMIN_ROLE", deployer.address)).to.be.true;
        expect(await has(bookingToken, "DEFAULT_ADMIN_ROLE", deployer.address)).to.be.true;
        expect(await has(manager, "UPGRADER_ROLE", deployer.address), "upgrader still renounced").to.be.false;
        expect(await has(manager, "VERSIONER_ROLE", deployer.address), "versioner still renounced").to.be.false;
        expect(await has(manager, "DEFAULT_ADMIN_ROLE", safe.address), "safe still admin").to.be.true;
    });

    // Wrap a contract so hasRole(role, safe) reports false for ONE role, forcing
    // the verify gate to fail. A plain Proxy is used (not reassigning a method on
    // the ethers instance) because ethers v6 method access goes through its own
    // proxy; `connect`/role-getters still delegate to the real contract, so the
    // grant/renounce transactions run for real.
    function withMissingSafeRole(contract, roleHash, safeAddr) {
        return new Proxy(contract, {
            get(target, prop) {
                if (prop === "hasRole") {
                    return async (role, addr) =>
                        role === roleHash && addr === safeAddr ? false : target.hasRole(role, addr);
                }
                const value = target[prop];
                return typeof value === "function" ? value.bind(target) : value;
            },
        });
    }

    it("aborts before any renounce if the Safe does not end up holding a role", async function () {
        const { manager, bookingToken, deployer, safe, pauser } = await loadFixture(deployAllOnDeployerFixture);

        const versioner = await manager.VERSIONER_ROLE();

        let error;
        try {
            await handoffRoles({
                manager: withMissingSafeRole(manager, versioner, safe.address),
                bookingToken,
                deployer,
                safe: safe.address,
                pauser: pauser.address,
                log: quiet,
            });
        } catch (e) {
            error = e;
        }
        expect(error, "handoffRoles should have thrown").to.exist;
        expect(error.message).to.match(/Verify gate failed/);

        // Deployer must still hold its roles on the REAL contract — nothing was renounced.
        expect(await has(manager, "DEFAULT_ADMIN_ROLE", deployer.address)).to.be.true;
        expect(await has(manager, "UPGRADER_ROLE", deployer.address)).to.be.true;
    });

    it("is idempotent — a second run is a no-op and leaves the topology unchanged", async function () {
        const { manager, bookingToken, deployer, safe, pauser } = await loadFixture(deployAllOnDeployerFixture);

        const args = { manager, bookingToken, deployer, safe: safe.address, pauser: pauser.address, log: quiet };
        await handoffRoles(args);
        await handoffRoles(args); // must not throw

        expect(await has(manager, "DEFAULT_ADMIN_ROLE", safe.address)).to.be.true;
        expect(await has(manager, "DEFAULT_ADMIN_ROLE", deployer.address)).to.be.false;
    });
});
```

- [ ] **Step 2: Run the test to verify it fails**

Run: `yarn hardhat test test/handoff.test.js`
Expected: FAIL — `Cannot find module '../tasks/lib/handoff'`.

- [ ] **Step 3: Write the minimal implementation**

Create `tasks/lib/handoff.js`:

```javascript
// Pure role-handoff logic, unit-tested in test/handoff.test.js. The Hardhat
// task in tasks/roles.js is a thin wrapper that resolves contracts + signer.

// Roles the Safe must end up holding.
const MANAGER_SAFE_ROLES = [
    "DEFAULT_ADMIN_ROLE",
    "UPGRADER_ROLE",
    "VERSIONER_ROLE",
    "PAUSER_ROLE",
    "SERVICE_REGISTRY_ADMIN_ROLE",
];
const BOOKINGTOKEN_SAFE_ROLES = ["DEFAULT_ADMIN_ROLE", "UPGRADER_ROLE", "PAUSER_ROLE"];

// Deployer renounce order — DEFAULT_ADMIN_ROLE is intentionally LAST so an
// aborted run never strands a contract without an admin.
const MANAGER_RENOUNCE_ORDER = [
    "VERSIONER_ROLE",
    "UPGRADER_ROLE",
    "PAUSER_ROLE",
    "SERVICE_REGISTRY_ADMIN_ROLE",
    "DEFAULT_ADMIN_ROLE",
];
const BOOKINGTOKEN_RENOUNCE_ORDER = ["UPGRADER_ROLE", "PAUSER_ROLE", "DEFAULT_ADMIN_ROLE"];

async function grantIfMissing(contract, roleName, address, deployer, log) {
    const role = await contract[roleName]();
    if (await contract.hasRole(role, address)) {
        log(`  ${roleName} already held by ${address}`);
        return;
    }
    const tx = await contract.connect(deployer).grantRole(role, address);
    await tx.wait();
    log(`  granted ${roleName} -> ${address}`);
}

async function renounceIfHeld(contract, roleName, deployer, log) {
    const role = await contract[roleName]();
    if (!(await contract.hasRole(role, deployer.address))) return;
    const tx = await contract.connect(deployer).renounceRole(role, deployer.address);
    await tx.wait();
    log(`  renounced ${roleName} from deployer`);
}

async function assertHolds(contract, roleName, address, label, missing) {
    const role = await contract[roleName]();
    if (!(await contract.hasRole(role, address))) {
        missing.push(`${label}.${roleName} not held by ${address}`);
    }
}

async function membership(contract, roleNames, address) {
    const out = {};
    for (const roleName of roleNames) {
        out[roleName] = await contract.hasRole(await contract[roleName](), address);
    }
    return out;
}

async function handoffRoles({
    manager,
    bookingToken,
    deployer,
    safe,
    pauser,
    keepDeployerAsDefaultAdmin = false,
    log = console.log,
}) {
    // 1. GRANT
    log("Granting roles to the Safe and hot pauser...");
    for (const role of MANAGER_SAFE_ROLES) {
        await grantIfMissing(manager, role, safe, deployer, log);
    }
    await grantIfMissing(manager, "PAUSER_ROLE", pauser, deployer, log);
    for (const role of BOOKINGTOKEN_SAFE_ROLES) {
        await grantIfMissing(bookingToken, role, safe, deployer, log);
    }
    await grantIfMissing(bookingToken, "PAUSER_ROLE", pauser, deployer, log);

    // MIN_EXPIRATION_ADMIN_ROLE is optional — only relayed if the deployer holds it.
    const minExp = await bookingToken.MIN_EXPIRATION_ADMIN_ROLE();
    const deployerHasMinExp = await bookingToken.hasRole(minExp, deployer.address);
    if (deployerHasMinExp) {
        await grantIfMissing(bookingToken, "MIN_EXPIRATION_ADMIN_ROLE", safe, deployer, log);
    }

    // 2. VERIFY GATE — abort before any renounce if the Safe/pauser fall short.
    const missing = [];
    for (const role of MANAGER_SAFE_ROLES) {
        await assertHolds(manager, role, safe, "manager", missing);
    }
    await assertHolds(manager, "PAUSER_ROLE", pauser, "manager", missing);
    for (const role of BOOKINGTOKEN_SAFE_ROLES) {
        await assertHolds(bookingToken, role, safe, "bookingToken", missing);
    }
    await assertHolds(bookingToken, "PAUSER_ROLE", pauser, "bookingToken", missing);
    if (missing.length > 0) {
        throw new Error(`Verify gate failed: ${missing.join("; ")}. Deployer roles left untouched.`);
    }
    log("Verify gate passed — Safe and hot pauser hold every target role.");

    // 3. RENOUNCE deployer roles, DEFAULT_ADMIN_ROLE last.
    log("Renouncing deployer roles...");
    for (const role of MANAGER_RENOUNCE_ORDER) {
        if (role === "DEFAULT_ADMIN_ROLE" && keepDeployerAsDefaultAdmin) continue;
        await renounceIfHeld(manager, role, deployer, log);
    }
    for (const role of BOOKINGTOKEN_RENOUNCE_ORDER) {
        if (role === "DEFAULT_ADMIN_ROLE" && keepDeployerAsDefaultAdmin) continue;
        await renounceIfHeld(bookingToken, role, deployer, log);
    }
    if (deployerHasMinExp) {
        await renounceIfHeld(bookingToken, "MIN_EXPIRATION_ADMIN_ROLE", deployer, log);
    }

    // 4. FINAL SUMMARY
    return {
        manager: {
            safe: await membership(manager, MANAGER_SAFE_ROLES, safe),
            pauser: await membership(manager, ["PAUSER_ROLE"], pauser),
            deployer: await membership(manager, MANAGER_SAFE_ROLES, deployer.address),
        },
        bookingToken: {
            safe: await membership(bookingToken, BOOKINGTOKEN_SAFE_ROLES, safe),
            pauser: await membership(bookingToken, ["PAUSER_ROLE"], pauser),
            deployer: await membership(bookingToken, BOOKINGTOKEN_SAFE_ROLES, deployer.address),
        },
    };
}

module.exports = {
    handoffRoles,
    MANAGER_SAFE_ROLES,
    BOOKINGTOKEN_SAFE_ROLES,
    MANAGER_RENOUNCE_ORDER,
    BOOKINGTOKEN_RENOUNCE_ORDER,
};
```

- [ ] **Step 4: Run the tests to verify they pass**

Run: `yarn hardhat test test/handoff.test.js`
Expected: PASS — 4 passing.

- [ ] **Step 5: Lint**

Run: `yarn lint`
Expected: 0 errors (pre-existing solhint warnings are unrelated to `.js`).

- [ ] **Step 6: Commit**

```bash
git add tasks/lib/handoff.js test/handoff.test.js
git commit -m "feat: role-handoff core logic with verify gate (phase 3)"
```

---

### Task 2: `roles:handoff` Hardhat task wrapper

**Files:**
- Create: `tasks/roles.js`
- Modify: `hardhat.config.js` (add `require("./tasks/roles");` beside the other task requires, ~line 10)

**Interfaces:**
- Consumes: `handoffRoles` from `tasks/lib/handoff.js` (Task 1).
- Produces: a Hardhat task `roles handoff` (scope `roles`, task `handoff`) with params `--safe <address>`, `--pauser <address>`, and flag `--keep-deployer-as-default-admin`. No programmatic export.

- [ ] **Step 1: Create the task wrapper**

Create `tasks/roles.js`:

```javascript
require("@nomicfoundation/hardhat-toolbox");
const { handoffRoles } = require("./lib/handoff");

const ROLES_SCOPE = scope("roles", "Privileged-role administration");

// Mirrors tasks/manager.js:getAddressesForNetwork — kept local so this task
// has no dependency on that file's internals.
const CHAIN_DIR = {
    localhost: "chain-31337",
    base_sepolia: "chain-84532",
    base: "chain-8453",
};

function resolveContracts(hre) {
    const dir = CHAIN_DIR[hre.network.name];
    if (!dir) throw new Error(`Unsupported network: ${hre.network.name}`);
    const addresses = require(`../ignition/deployments/${dir}/deployed_addresses.json`);
    return {
        managerAddress: addresses["TravelTokenMessengerModule#ManagerProxy"],
        bookingTokenAddress: addresses["TravelTokenMessengerModule#BookingTokenProxy"],
    };
}

ROLES_SCOPE.task("handoff", "Hand privileged roles to a Safe (+ hot pauser) and de-privilege the deployer")
    .addParam("safe", "Safe address that receives DEFAULT_ADMIN/UPGRADER/VERSIONER/PAUSER/SERVICE_REGISTRY_ADMIN")
    .addParam("pauser", "Dedicated hot pauser EOA that receives PAUSER_ROLE on both contracts")
    .addFlag(
        "keepDeployerAsDefaultAdmin",
        "TESTNET ONLY: skip renouncing the deployer's DEFAULT_ADMIN_ROLE (break-glass recovery admin)",
    )
    .setAction(async (taskArgs, hre) => {
        const { ethers } = hre;
        const [deployer] = await ethers.getSigners();
        const { managerAddress, bookingTokenAddress } = resolveContracts(hre);

        const manager = await ethers.getContractAt("TTMAccountManager", managerAddress);
        const bookingToken = await ethers.getContractAt("BookingToken", bookingTokenAddress);

        console.log(`Network:  ${hre.network.name}`);
        console.log(`Deployer: ${deployer.address}`);
        console.log(`Safe:     ${taskArgs.safe}`);
        console.log(`Pauser:   ${taskArgs.pauser}`);
        if (taskArgs.keepDeployerAsDefaultAdmin) {
            console.log(
                "\n⚠️  --keep-deployer-as-default-admin is set: the deployer keeps DEFAULT_ADMIN_ROLE.\n" +
                    "    This is a TESTNET recovery hatch. Do NOT use it on Base mainnet.\n",
            );
        }

        const summary = await handoffRoles({
            manager,
            bookingToken,
            deployer,
            safe: taskArgs.safe,
            pauser: taskArgs.pauser,
            keepDeployerAsDefaultAdmin: taskArgs.keepDeployerAsDefaultAdmin,
        });

        console.log("\nFinal membership:");
        console.log(JSON.stringify(summary, null, 2));
    });
```

- [ ] **Step 2: Register the task**

In `hardhat.config.js`, add the require next to the existing ones:

```javascript
require("./tasks/manager");
require("./tasks/account");
require("./tasks/roles");
```

- [ ] **Step 3: Verify the task is registered**

Run: `yarn hardhat roles handoff --help`
Expected: usage output listing `--safe`, `--pauser`, and `--keep-deployer-as-default-admin`. (It should NOT error with "unrecognized task".)

- [ ] **Step 4: Smoke-test against a throwaway local deployment (optional but recommended)**

Run:
```bash
yarn hardhat node &          # in one shell
yarn hardhat ignition deploy ignition/modules/messenger.js --network localhost
yarn hardhat roles handoff --network localhost \
  --safe 0x70997970C51812dc3A010C7d01b50e0d17dc79C8 \
  --pauser 0x3C44CdDdB6a900fa2b585dd299e03d12FA4293BC
```
Expected: grant lines, "Verify gate passed", renounce lines, and a final-membership JSON showing the Safe address holding the admin roles and the deployer holding none. Kill the node afterward.

- [ ] **Step 5: Lint**

Run: `yarn lint`
Expected: 0 errors.

- [ ] **Step 6: Commit**

```bash
git add tasks/roles.js hardhat.config.js
git commit -m "feat: roles:handoff task wiring the Safe custody deploy step (phase 3)"
```

---

### Task 3: Deploy runbook documentation

**Files:**
- Modify: `README.md` (the `### Deploy (Hardhat Ignition)` section)

**Interfaces:**
- Consumes: the `roles handoff` task (Task 2). No code.

- [ ] **Step 1: Add the Safe/pauser prerequisite**

In `README.md`, before the numbered deploy steps in `### Deploy (Hardhat Ignition)`, add:

```markdown
**Prerequisite — create custody keys.** Before deploying:

- Create the Safe on Base Sepolia via the Safe web app: `SafeL2` singleton
  (`0x29fcB43b46531BcA003ddC8FCB67FFE91900C762`) with the
  `CompatibilityFallbackHandler` (`0xfd0732Dc9E303f09fCEf3a7388Ad10A83459Ec99`),
  the owner set you control, and your chosen threshold. Record the Safe address.
- Provision a dedicated hot **pauser** EOA, separate from the Safe owner keys.

`base_sepolia_parameters.json` stays `{}` — every role deploys onto the deployer
key and is handed off in step 8. This keeps `managerVersioner` on the deployer
through the Ignition run (the module's `setAccountImplementation` /
`setBookingTokenAddress` calls need `VERSIONER_ROLE` as account 0).
```

- [ ] **Step 2: Remove the now-automated manual step 6**

Delete the line `# 6. Grant PAUSER_ROLE on BookingToken to the operations key` from the fenced command block — the handoff task grants BookingToken `PAUSER_ROLE` to both the Safe and the hot pauser.

- [ ] **Step 3: Replace the step-8 prose with the task**

Replace the `**8. Hand off admin roles.**` paragraph with:

```markdown
**8. Hand off privileged roles to the Safe.**

    yarn hardhat roles handoff --network base_sepolia \
      --safe <safe-address> --pauser <hot-pauser-address>

The task grants the Safe `DEFAULT_ADMIN_ROLE`, `UPGRADER_ROLE`, `VERSIONER_ROLE`,
`PAUSER_ROLE` and `SERVICE_REGISTRY_ADMIN_ROLE` on the manager and
`DEFAULT_ADMIN_ROLE`/`UPGRADER_ROLE`/`PAUSER_ROLE` on BookingToken; grants the
hot pauser `PAUSER_ROLE` on both; **verifies the Safe holds every role before it
renounces anything** (the manager is a singleton — an incomplete grant must not
strand it without an admin); then renounces the deployer's roles,
`DEFAULT_ADMIN_ROLE` last. It is idempotent and safe to re-run.

Pass `--keep-deployer-as-default-admin` to keep the deployer as a break-glass
recovery admin. **Testnet only** — do not use it on Base mainnet.
```

- [ ] **Step 4: Verify formatting**

Run: `npx prettier --check README.md`
Expected: no output / "All matched files use Prettier code style". If it fails, run `npx prettier --write README.md`.

- [ ] **Step 5: Confirm no generated artifacts changed**

Run: `git status --porcelain`
Expected: only `README.md` (and earlier tasks' files if uncommitted) — **no** changes under `abi/`, `docs/`, `go/`, or `contracts/`.

- [ ] **Step 6: Commit**

```bash
git add README.md
git commit -m "docs: deploy runbook for roles:handoff and Safe custody (phase 3)"
```

---

## Self-Review

**Spec coverage:**
- Role topology table → `MANAGER_SAFE_ROLES` / `BOOKINGTOKEN_SAFE_ROLES` + pauser grants (Task 1). ✓
- Option B (PAUSER on Safe + hot key) → `PAUSER_ROLE` granted to both (Task 1). ✓
- Verify gate before renounce → step 2 of `handoffRoles`, test "aborts before any renounce" (Task 1). ✓
- `DEFAULT_ADMIN` renounced last → `MANAGER_RENOUNCE_ORDER` / `BOOKINGTOKEN_RENOUNCE_ORDER` ordering. ✓
- `--keep-deployer-as-default-admin` (testnet-only, both contracts, skips only DEFAULT_ADMIN) → flag + test (Tasks 1, 2). ✓
- `SERVICE_REGISTRY_ADMIN_ROLE` folded into handoff → in `MANAGER_SAFE_ROLES` + renounce order. ✓
- `MIN_EXPIRATION_ADMIN_ROLE` conditional → guarded by `deployerHasMinExp` (Task 1). ✓
- Approach H / parameters `{}` / trap #7 → README prerequisite note (Task 3); no parameters file edit. ✓
- Idempotency → `grantIfMissing` / `renounceIfHeld` + test (Task 1). ✓
- No Solidity/ABI/docgen/bindings change → Global Constraints + Task 3 step 5 check. ✓
- Local-only tests (no deploy) → fixture uses `upgrades.deployProxy` on the in-process network. ✓

**Placeholder scan:** none — every step has concrete code or an exact command.

**Type consistency:** `handoffRoles(...)` signature, `MANAGER_SAFE_ROLES` / `BOOKINGTOKEN_SAFE_ROLES` exports, and the `deployed_addresses.json` keys are used identically across Tasks 1–3.
