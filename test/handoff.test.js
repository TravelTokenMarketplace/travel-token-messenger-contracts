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

    it("relays MIN_EXPIRATION_ADMIN_ROLE to the Safe when the deployer holds it", async function () {
        const { manager, bookingToken, deployer, safe, pauser } = await loadFixture(deployAllOnDeployerFixture);

        await bookingToken.grantRole(await bookingToken.MIN_EXPIRATION_ADMIN_ROLE(), deployer.address);

        await handoffRoles({
            manager,
            bookingToken,
            deployer,
            safe: safe.address,
            pauser: pauser.address,
            log: quiet,
        });

        expect(
            await has(bookingToken, "MIN_EXPIRATION_ADMIN_ROLE", safe.address),
            "bt MIN_EXPIRATION_ADMIN_ROLE -> safe",
        ).to.be.true;
        expect(
            await has(bookingToken, "MIN_EXPIRATION_ADMIN_ROLE", deployer.address),
            "bt MIN_EXPIRATION_ADMIN_ROLE deployer renounced",
        ).to.be.false;
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
