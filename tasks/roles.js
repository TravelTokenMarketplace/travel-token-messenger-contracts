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
