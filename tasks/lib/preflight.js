// Deploy-time custody checks for the `roles handoff` task, unit-tested in
// test/preflight.test.js. Kept separate from tasks/lib/handoff.js: that file is
// role logic against contract handles, this one is about who the counterparties
// actually are on chain.

const { Contract } = require("ethers");

// Safe's owner-set getters. Probing them proves the address is a live Safe and
// not merely some contract that happens to sit at a mistyped address.
const SAFE_ABI = ["function getOwners() view returns (address[])", "function getThreshold() view returns (uint256)"];

// --keep-deployer-as-default-admin leaves one EOA with permanent admin
// authority. That is a testnet recovery hatch; an allowlist (rather than a
// `base` denylist) keeps any future production network out by default.
const BREAK_GLASS_NETWORKS = ["hardhat", "localhost", "base_sepolia"];

function assertBreakGlassAllowed(networkName) {
    if (!BREAK_GLASS_NETWORKS.includes(networkName)) {
        throw new Error(
            `--keep-deployer-as-default-admin is refused on network "${networkName}". ` +
                `It leaves the deployer EOA holding DEFAULT_ADMIN_ROLE permanently, ` +
                `and is allowed only on: ${BREAK_GLASS_NETWORKS.join(", ")}.`,
        );
    }
}

// Verifies custody *type*, which role membership cannot express: the Safe must
// be a contract that answers the Safe interface, the hot pauser must be an EOA.
// Returns the Safe's owner set so the operator can eyeball it before confirming.
async function preflightCustody({ provider, safe, pauser }) {
    if ((await provider.getCode(safe)) === "0x") {
        throw new Error(
            `The Safe ${safe} has no contract code on this network. ` +
                `A mistyped-but-valid EOA would receive every administrative role, and once the ` +
                `deployer renounces, recovery would depend on that single key.`,
        );
    }

    if ((await provider.getCode(pauser)) !== "0x") {
        throw new Error(
            `The hot pauser ${pauser} is a contract. It must be an EOA that can sign ` +
                `pause() immediately, without a multisig round-trip.`,
        );
    }

    const safeContract = new Contract(safe, SAFE_ABI, provider);
    let owners;
    let threshold;
    try {
        owners = await safeContract.getOwners();
        threshold = await safeContract.getThreshold();
    } catch (error) {
        throw new Error(
            `The address ${safe} has code but does not look like a Safe — ` +
                `getOwners()/getThreshold() did not answer (${error.shortMessage ?? error.message}).`,
        );
    }

    return { owners: [...owners], threshold };
}

module.exports = { assertBreakGlassAllowed, preflightCustody, BREAK_GLASS_NETWORKS, SAFE_ABI };
