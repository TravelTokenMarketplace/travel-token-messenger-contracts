require("@nomicfoundation/hardhat-toolbox");
require("@openzeppelin/hardhat-upgrades");
require("hardhat-contract-sizer");
require("solidity-docgen");
require("hardhat-abi-exporter");

// Tasks
require("./tasks/manager");
require("./tasks/account");

/** @type import('hardhat/config').HardhatUserConfig */
module.exports = {
    solidity: {
        version: "0.8.24",
        settings: {
            optimizer: {
                enabled: true,
                runs: 1,
            },
            evmVersion: "paris",
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
        apiKey: {
            base_sepolia: vars.get("BASESCAN_API_KEY", "abc"),
            base: vars.get("BASESCAN_API_KEY", "abc"),
        },
        customChains: [
            {
                network: "base_sepolia",
                chainId: 84532,
                urls: {
                    apiURL: "https://api-sepolia.basescan.org/api",
                    browserURL: "https://sepolia.basescan.org",
                },
            },
            {
                network: "base",
                chainId: 8453,
                urls: {
                    apiURL: "https://api.basescan.org/api",
                    browserURL: "https://basescan.org",
                },
            },
        ],
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
