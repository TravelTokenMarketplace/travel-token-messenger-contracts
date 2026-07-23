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
const BOOKINGTOKEN_RENOUNCE_ORDER = ["UPGRADER_ROLE", "PAUSER_ROLE", "MIN_EXPIRATION_ADMIN_ROLE", "DEFAULT_ADMIN_ROLE"];

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
