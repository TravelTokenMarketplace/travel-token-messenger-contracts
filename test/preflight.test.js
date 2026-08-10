const { loadFixture } = require("@nomicfoundation/hardhat-toolbox/network-helpers");
const { expect } = require("chai");
const { ethers } = require("hardhat");

const { assertBreakGlassAllowed, preflightCustody } = require("../tasks/lib/preflight");

// A real Safe-like contract, a real non-Safe contract, and real EOAs — the
// checks under test are all `provider.getCode` / real `eth_call`s.
// The SafeL2 singleton recorded for Base Sepolia in the deploy runbook.
const SAFE_L2_SINGLETON = "0x29fcB43b46531BcA003ddC8FCB67FFE91900C762";

async function custodyFixture() {
    const [deployer, eoa] = await ethers.getSigners();

    const MockSafe = await ethers.getContractFactory("MockSafe");
    const mockSafe = await MockSafe.deploy(SAFE_L2_SINGLETON, [deployer.address, eoa.address], 1);
    await mockSafe.waitForDeployment();

    const Dummy = await ethers.getContractFactory("Dummy");
    const notASafe = await Dummy.deploy();
    await notASafe.waitForDeployment();

    return {
        eoa,
        otherEoa: deployer,
        safeAddress: await mockSafe.getAddress(),
        notASafeAddress: await notASafe.getAddress(),
    };
}

describe("assertBreakGlassAllowed", function () {
    it("throws on Base mainnet", function () {
        expect(() => assertBreakGlassAllowed("base")).to.throw(/base/i);
    });

    it("throws on any network outside the testnet allowlist", function () {
        expect(() => assertBreakGlassAllowed("mainnet")).to.throw(/keep-deployer-as-default-admin/i);
    });

    it("allows Base Sepolia and local networks", function () {
        expect(() => assertBreakGlassAllowed("base_sepolia")).to.not.throw();
        expect(() => assertBreakGlassAllowed("localhost")).to.not.throw();
        expect(() => assertBreakGlassAllowed("hardhat")).to.not.throw();
    });
});

describe("preflightCustody", function () {
    it("returns the Safe's owners and threshold for a real Safe", async function () {
        const { safeAddress, eoa, otherEoa } = await loadFixture(custodyFixture);

        const info = await preflightCustody({
            provider: ethers.provider,
            safe: safeAddress,
            pauser: eoa.address,
        });

        expect(info.owners).to.deep.equal([otherEoa.address, eoa.address]);
        expect(info.threshold).to.equal(1n);
    });

    it("reports the singleton the Safe proxy runs, for the operator to confirm", async function () {
        const { safeAddress, eoa } = await loadFixture(custodyFixture);

        const info = await preflightCustody({
            provider: ethers.provider,
            safe: safeAddress,
            pauser: eoa.address,
        });

        expect(info.singleton).to.equal(SAFE_L2_SINGLETON);
    });

    it("reports a null singleton when slot 0 does not hold an address", async function () {
        const { notASafeAddress } = await loadFixture(custodyFixture);

        // Dummy has no storage at all, so slot 0 reads back as zero. Proving the
        // preflight reports that rather than printing a bogus 0x00…0 "singleton".
        const { readSafeSingleton } = require("../tasks/lib/preflight");
        expect(await readSafeSingleton(ethers.provider, notASafeAddress)).to.be.null;
    });

    it("rejects a Safe address with no code — a mistyped EOA would receive every admin role", async function () {
        const { eoa, otherEoa } = await loadFixture(custodyFixture);

        await expect(
            preflightCustody({ provider: ethers.provider, safe: otherEoa.address, pauser: eoa.address }),
        ).to.be.rejectedWith(/safe .* has no contract code/is);
    });

    it("rejects a Safe address whose code does not answer the Safe interface", async function () {
        const { notASafeAddress, eoa } = await loadFixture(custodyFixture);

        await expect(
            preflightCustody({ provider: ethers.provider, safe: notASafeAddress, pauser: eoa.address }),
        ).to.be.rejectedWith(/does not look like a Safe/i);
    });

    it("rejects a hot pauser that is a contract", async function () {
        const { safeAddress, notASafeAddress } = await loadFixture(custodyFixture);

        await expect(
            preflightCustody({ provider: ethers.provider, safe: safeAddress, pauser: notASafeAddress }),
        ).to.be.rejectedWith(/pauser .* is a contract/is);
    });
});
