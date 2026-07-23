require("@nomicfoundation/hardhat-toolbox");
require("@openzeppelin/hardhat-upgrades");
require("hardhat-contract-sizer");
require("solidity-docgen");
require("hardhat-abi-exporter");
const { vars } = require("hardhat/config");

// Tasks
require("./tasks/manager");
require("./tasks/account");
require("./tasks/roles");

/** @type import('hardhat/config').HardhatUserConfig */
module.exports = {
    solidity: {
        version: "0.8.24",
        settings: {
            optimizer: {
                enabled: true,
                runs: 1000,
            },
            evmVersion: "cancun",
        },
    },
    contractSizer: {
        runOnCompile: true,
    },
    ignition: {
        requiredConfirmations: 1,
    },
    networks: {
        localhost: {
            url: "http://127.0.0.1:8545",
        },
        base_sepolia: {
            url: vars.get("BASE_SEPOLIA_URL", "https://base-sepolia.drpc.org"),
            accounts: vars.has("BASE_SEPOLIA_DEPLOYER_PRIVATE_KEY")
                ? [vars.get("BASE_SEPOLIA_DEPLOYER_PRIVATE_KEY")]
                : [],
        },
        base: {
            url: vars.get("BASE_URL", "https://base.drpc.org"),
            accounts: vars.has("BASE_DEPLOYER_PRIVATE_KEY") ? [vars.get("BASE_DEPLOYER_PRIVATE_KEY")] : [],
        },
    },
    etherscan: {
        apiKey: vars.get("ETHERSCAN_API_KEY", ""),
    },
    docgen: {
        path: "./docs",
        pages: "single",
        runOnCompile: true,
    },
    abiExporter: {
        path: "./abi",
        runOnCompile: false,
        format: "json",
        clear: true,
    },
};
