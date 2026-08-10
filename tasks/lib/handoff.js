// Pure role-handoff logic, unit-tested in test/handoff.test.js. The Hardhat
// task in tasks/roles.js is a thin wrapper that resolves contracts + signer.

const { getAddress, ZeroAddress } = require("ethers");

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
const BOOKINGTOKEN_RENOUNCE_ORDER = ["UPGRADER_ROLE", "PAUSER_ROLE", "MIN_EXPIRATION_ADMIN_ROLE", "DEFAULT_ADMIN_ROLE"];

// The hot pauser is a standing online key: it must be able to pause, and nothing
// else. These are every role on each contract other than PAUSER_ROLE — holding
// one would leave the pauser silently privileged after the deployer renounces.
const MANAGER_ADMIN_ROLES = MANAGER_SAFE_ROLES.filter((role) => role !== "PAUSER_ROLE");
const BOOKINGTOKEN_ADMIN_ROLES = ["DEFAULT_ADMIN_ROLE", "UPGRADER_ROLE", "MIN_EXPIRATION_ADMIN_ROLE"];

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

function normalizePrincipal(label, value) {
    let address;
    try {
        address = getAddress(value);
    } catch {
        throw new Error(`Invalid ${label} address: ${value}`);
    }
    if (address === ZeroAddress) {
        throw new Error(`The ${label} must not be the zero address.`);
    }
    return address;
}

// The manager is a singleton: if two principals collide, the verify gate checks
// an address that is about to be stripped by the renounce loop and passes. With
// safe === deployer that strips the last DEFAULT_ADMIN_ROLE and bricks it; with
// pauser === deployer it leaves no hot pauser. Reject before the first tx.
function validatePrincipals({ deployer, safe, pauser }) {
    const principals = {
        deployer: normalizePrincipal("deployer", deployer),
        safe: normalizePrincipal("safe", safe),
        pauser: normalizePrincipal("pauser", pauser),
    };

    const seen = new Map();
    for (const [label, address] of Object.entries(principals)) {
        const clash = seen.get(address);
        if (clash) {
            throw new Error(
                `The deployer, Safe and hot pauser must be three distinct addresses — ` +
                    `${clash} and ${label} are both ${address}.`,
            );
        }
        seen.set(address, label);
    }
    return principals;
}

// Catches a typo'd or reused hot-pauser key: on a fresh deploy the pauser is a
// new EOA holding nothing, so this is a no-op — but if it does hold admin
// authority, the deployer renounce would lock that in unnoticed.
async function assertPauserIsNotPrivileged(manager, bookingToken, pauser) {
    const held = [];
    for (const role of MANAGER_ADMIN_ROLES) {
        if (await manager.hasRole(await manager[role](), pauser)) held.push(`manager.${role}`);
    }
    for (const role of BOOKINGTOKEN_ADMIN_ROLES) {
        if (await bookingToken.hasRole(await bookingToken[role](), pauser)) held.push(`bookingToken.${role}`);
    }
    if (held.length > 0) {
        throw new Error(
            `The hot pauser ${pauser} already holds an administrative role (${held.join(", ")}). ` +
                `It must hold PAUSER_ROLE only — check the address, or revoke those roles first.`,
        );
    }
}

async function handoffRoles({
    manager,
    bookingToken,
    deployer,
    safe: safeInput,
    pauser: pauserInput,
    keepDeployerAsDefaultAdmin = false,
    log = console.log,
}) {
    // 0. PREFLIGHT — no transaction has been sent yet, so throwing here is free.
    const { safe, pauser } = validatePrincipals({
        deployer: deployer.address,
        safe: safeInput,
        pauser: pauserInput,
    });

    await assertPauserIsNotPrivileged(manager, bookingToken, pauser);

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
    if (deployerHasMinExp) {
        await assertHolds(bookingToken, "MIN_EXPIRATION_ADMIN_ROLE", safe, "bookingToken", missing);
    }
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

    // 4. FINAL SUMMARY — the only operator-facing confirmation, so it must cover
    // the optional role too whenever it was actually part of the handoff.
    const bookingTokenSummaryRoles = deployerHasMinExp
        ? [...BOOKINGTOKEN_SAFE_ROLES, "MIN_EXPIRATION_ADMIN_ROLE"]
        : BOOKINGTOKEN_SAFE_ROLES;

    return {
        manager: {
            safe: await membership(manager, MANAGER_SAFE_ROLES, safe),
            pauser: await membership(manager, ["PAUSER_ROLE"], pauser),
            deployer: await membership(manager, MANAGER_SAFE_ROLES, deployer.address),
        },
        bookingToken: {
            safe: await membership(bookingToken, bookingTokenSummaryRoles, safe),
            pauser: await membership(bookingToken, ["PAUSER_ROLE"], pauser),
            deployer: await membership(bookingToken, bookingTokenSummaryRoles, deployer.address),
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
