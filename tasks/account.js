require("@nomicfoundation/hardhat-toolbox");
const { types } = require("hardhat/config");

const ACCOUNT_SCOPE = scope("account", "TTM Account Tasks");

// TODO: Get private key from .env or hardhat vars

const ACC_ROLES = [
    "DEFAULT_ADMIN_ROLE",
    "UPGRADER_ROLE",
    "BOT_ADMIN_ROLE",
    "MESSENGER_BOT_ROLE",
    "GAS_WITHDRAWER_ROLE",
    "WITHDRAWER_ROLE",
    "BOOKING_OPERATOR_ROLE",
    "SERVICE_ADMIN_ROLE",
];

function bold(text) {
    const boldCode = "\x1b[1m";
    const resetCode = "\x1b[0m";
    return `${boldCode}${text}${resetCode}`;
}

function getAddressesForNetwork(hre) {
    let addresses;

    if (hre.network.name === "localhost") {
        console.log("Running on localhost");
        addresses = require("../ignition/deployments/chain-31337/deployed_addresses.json");
    } else if (hre.network.name === "base_sepolia") {
        console.log("Running on base_sepolia");
        addresses = require("../ignition/deployments/chain-84532/deployed_addresses.json");
    } else if (hre.network.name === "base") {
        console.log("Running on base");
        addresses = require("../ignition/deployments/chain-8453/deployed_addresses.json");
    } else {
        throw new Error(`Unsupported network: ${hre.network.name}`);
    }

    return addresses;
}

async function getManager(hre) {
    const addresses = getAddressesForNetwork(hre);
    return await ethers.getContractAt("TTMAccountManager", addresses["TravelTokenMessengerModule#ManagerProxy"]);
}

async function getTTMAccount(ttmAccountAddress) {
    return await ethers.getContractAt("TTMAccount", ttmAccountAddress);
}

async function handleRoles(taskArgs, hre, action) {
    const ttmAccount = await getTTMAccount(taskArgs.ttmAccount);
    console.log("TTMAccount:", taskArgs.ttmAccount);

    try {
        const signer = new ethers.Wallet(taskArgs.privateKey, ethers.provider);

        console.log(
            `${action === "grantRole" ? "Granting" : "Revoking"} role ${taskArgs.role} for address ${taskArgs.address}...`,
        );

        const role = await ttmAccount.connect(signer)[taskArgs.role]();
        const tx = await ttmAccount.connect(signer)[action](role, taskArgs.address);
        const txReceipt = await tx.wait();
        console.log("Tx:", txReceipt.hash);
    } catch (error) {
        handleTransactionError(error, ttmAccount);
    }
}

function handleTransactionError(error, contract) {
    console.error("❌ Transaction failed!");

    if (error.data?.message) {
        console.error(`Reason: ${error.data.message}`);
    } else if (error.data && contract) {
        const decodedError = contract.interface.parseError(error.data);
        console.error("Message:", error.message);
        console.error(`Reason: ${decodedError?.name} (${decodedError?.args})`);
    } else if (error.message?.includes("[taskArgs.role] is not a function")) {
        console.error("Reason: TTMAccount does not have this role.");
    } else if (error.message) {
        console.error("Message:", error.message);
    } else {
        // General error logging
        console.error("An unexpected error occurred.");
        console.error("Error:", error);
    }
}

async function getImplementationAddressForProxy(proxyAddress) {
    // Implementation slot for ERC1967Proxy
    const IMPLEMENTATION_SLOT = "0x360894a13ba1a3210667c828492db98dca3e2076cc3735a920a3ca505d382bbc";

    // Read the implementation slot from the proxy
    const paddedAddress = await ethers.provider.getStorage(proxyAddress, IMPLEMENTATION_SLOT);

    // Convert the result to an address
    const address = ethers.getAddress(ethers.dataSlice(paddedAddress, 12));

    return address;
}

ACCOUNT_SCOPE.task("role:grant", "Grant role")
    .addOptionalParam(
        "privateKey",
        "Private key to use, default: TTMACCOUNT_PK env variable",
        process.env.TTMACCOUNT_PK,
    )
    .addParam("role", "Role to grant. Ex: SERVICE_ADMIN_ROLE")
    .addParam("address", "Address to grant role to")
    .addOptionalParam(
        "ttmAccount",
        "TTMAccount address, default: TTMACCOUNT_ADDRESS env variable",
        process.env.TTMACCOUNT_ADDRESS,
    )
    .setAction(async (taskArgs, hre) => {
        await handleRoles(taskArgs, hre, "grantRole");
    });

ACCOUNT_SCOPE.task("role:revoke", "Revoke role")
    .addOptionalParam(
        "privateKey",
        "Private key to use, default: TTMACCOUNT_PK env variable",
        process.env.TTMACCOUNT_PK,
    )
    .addParam("role", "Role to grant. Ex: SERVICE_ADMIN_ROLE")
    .addParam("address", "Address to revoke role to")
    .addOptionalParam(
        "ttmAccount",
        "TTMAccount address, default: TTMACCOUNT_ADDRESS env variable",
        process.env.TTMACCOUNT_ADDRESS,
    )
    .setAction(async (taskArgs, hre) => {
        await handleRoles(taskArgs, hre, "revokeRole");
    });

ACCOUNT_SCOPE.task("role:has", "Check if address has role")
    .addParam("role", "Role to check. Ex: SERVICE_ADMIN_ROLE")
    .addParam("address", "Address to check")
    .addOptionalParam(
        "ttmAccount",
        "TTMAccount address, default: TTMACCOUNT_ADDRESS env variable",
        process.env.TTMACCOUNT_ADDRESS,
    )
    .setAction(async (taskArgs, hre) => {
        const ttmAccount = await getTTMAccount(taskArgs.ttmAccount);
        console.log("TTMAccount:", taskArgs.ttmAccount);

        try {
            console.log("Running on", hre.network.name);
            const role = await ttmAccount[taskArgs.role]();
            const hasRole = await ttmAccount.hasRole(role, taskArgs.address);

            console.log(`Address ${taskArgs.address} ${hasRole ? "has" : "does not have"} role ${taskArgs.role}`);
            console.log(`${hasRole ? "🟢" : "🔴"}`, hasRole);
        } catch (error) {
            handleTransactionError(error, ttmAccount);
        }
    });

ACCOUNT_SCOPE.task("role:members", "List role members")
    .addParam("role", "Role to list. Ex: SERVICE_ADMIN_ROLE")
    .addOptionalParam(
        "ttmAccount",
        "TTMAccount address, default: TTMACCOUNT_ADDRESS env variable",
        process.env.TTMACCOUNT_ADDRESS,
    )
    .setAction(async (taskArgs, hre) => {
        const ttmAccount = await getTTMAccount(taskArgs.ttmAccount);
        console.log("TTMAccount:", taskArgs.ttmAccount);
        console.log("Role:", taskArgs.role);

        try {
            const role = await ttmAccount[taskArgs.role]();
            const count = await ttmAccount.getRoleMemberCount(role);
            console.log("Total Members:", count);

            // Iterate over the members of the role
            const members = [];
            for (let i = 0; i < count; i++) {
                const member = await ttmAccount.getRoleMember(role, i);
                members.push(member);
            }
            console.log(members);
        } catch (error) {
            handleTransactionError(error, ttmAccount);
        }
    });

ACCOUNT_SCOPE.task("role:all", "List all roles and their members")
    .addOptionalParam(
        "ttmAccount",
        "TTMAccount address, default: TTMACCOUNT_ADDRESS env variable",
        process.env.TTMACCOUNT_ADDRESS,
    )
    .setAction(async (taskArgs, hre) => {
        for (const role of ACC_ROLES) {
            console.log(`🛡️  ${bold(role)}`);
            console.log(`${bold("=".repeat(53))}`);
            await hre.run({ scope: "account", task: "role:members" }, { role, ttmAccount: taskArgs.ttmAccount });
            console.log();
        }
    });

ACCOUNT_SCOPE.task("create", "Create TTMAccount")
    .addOptionalParam(
        "privateKey",
        "Private key to use, default: TTMACCOUNT_PK env variable",
        process.env.TTMACCOUNT_PK,
    )
    .addOptionalParam("amount", "Amount of ETH to send to TTMAccount", "0")
    .setAction(async (taskArgs, hre) => {
        const manager = await getManager(hre);

        try {
            // Get signer from private key
            const signer = new ethers.Wallet(taskArgs.privateKey, ethers.provider);
            console.log("Signer:", signer.address);

            // Create TTMAccount
            const parsedAmount = ethers.parseEther(taskArgs.amount);
            const formattedAmount = ethers.formatEther(parsedAmount);
            console.log(`Creating TTMAccount... (Sending ${formattedAmount} ETH to the new TTMAccount)`);
            const tx = await manager
                .connect(signer)
                .createTTMAccount(signer.address, signer.address, { value: parsedAmount });

            const receipt = await tx.wait();
            console.log("Tx:", receipt.hash);

            // Parse event to get the TTMAccount address (this is the UUPS proxy address)
            const event = receipt.logs.find((log) => {
                try {
                    return manager.interface.parseLog(log).name === "TTMAccountCreated";
                } catch (e) {
                    return false;
                }
            });

            const parsedEvent = manager.interface.parseLog(event);
            const ttmAccountAddress = parsedEvent.args.account;

            console.log("TTMAccount Address:", ttmAccountAddress);
        } catch (error) {
            handleTransactionError(error, manager);
        }
    });

ACCOUNT_SCOPE.task("withdraw", "Withdraw funds from TTMAccount")
    .addOptionalParam(
        "privateKey",
        "Private key to use, default: TTMACCOUNT_PK env variable",
        process.env.TTMACCOUNT_PK,
    )
    .addOptionalParam(
        "ttmAccount",
        "TTMAccount address, default: TTMACCOUNT_ADDRESS env variable",
        process.env.TTMACCOUNT_ADDRESS,
    )
    .addParam("recipient", "Recipient address")
    .addParam("amount", "Amount to withdraw")
    .addOptionalParam("unit", "Unit of amount (eth/gwei/wei)", "wei")
    .setAction(async (taskArgs, hre) => {
        const ttmAccount = await getTTMAccount(taskArgs.ttmAccount);
        console.log("TTMAccount:", taskArgs.ttmAccount);

        try {
            // Validate unit
            const validUnits = ["eth", "gwei", "wei"];
            if (!validUnits.includes(taskArgs.unit.toLowerCase())) {
                throw new Error(`Invalid unit. Must be one of: ${validUnits.join(", ")}`);
            }

            // Convert amount to wei based on unit
            let amountInWei;
            let hrUnit;
            const unit = taskArgs.unit.toLowerCase();

            switch (unit) {
                case "eth":
                    amountInWei = ethers.parseEther(taskArgs.amount.toString());
                    hrUnit = "ETH";
                    break;
                case "gwei":
                    amountInWei = ethers.parseUnits(taskArgs.amount.toString(), "gwei");
                    hrUnit = "GWEI";
                    break;
                case "wei":
                    amountInWei = taskArgs.amount.toString();
                    hrUnit = "WEI";
                    break;
            }

            console.log("Running on", hre.network.name);
            console.log("💸 Withdrawing funds...");
            console.log("From         :", taskArgs.ttmAccount);
            console.log("To           :", taskArgs.recipient);
            console.log("Amount       :", `${taskArgs.amount} ${hrUnit}`);
            console.log("Amount (wei) :", amountInWei.toString());

            const signer = new ethers.Wallet(taskArgs.privateKey, ethers.provider);
            const tx = await ttmAccount.connect(signer).withdraw(taskArgs.recipient, amountInWei);
            const receipt = await tx.wait();
            console.log("Tx:", receipt.hash);
        } catch (error) {
            handleTransactionError(error, ttmAccount);
        }
    });

ACCOUNT_SCOPE.task("bot:add", "Add bot to the TTMAccount")
    .addOptionalParam(
        "privateKey",
        "Private key to use, default: TTMACCOUNT_PK env variable",
        process.env.TTMACCOUNT_PK,
    )
    .addOptionalParam(
        "ttmAccount",
        "TTMAccount address, default: TTMACCOUNT_ADDRESS env variable",
        process.env.TTMACCOUNT_ADDRESS,
    )
    .addParam("bot", "Bot address")
    .addOptionalParam(
        "gasMoney",
        "Gas money in ETH. This amount will be transferred from the TTMAccount to the bot address (Ex: 1 or 0.1)",
        "0",
        types.string,
    )
    .setAction(async (taskArgs, hre) => {
        const ttmAccount = await getTTMAccount(taskArgs.ttmAccount);
        console.log("TTMAccount:", taskArgs.ttmAccount);
        console.log("Bot:", taskArgs.bot);
        console.log(
            "Gas:",
            taskArgs.gasMoney,
            "(This amount will be transferred from the TTMAccount to the bot address)",
        );

        try {
            const signer = new ethers.Wallet(taskArgs.privateKey, ethers.provider);
            console.log("Adding bot to TTMAccount...");
            console.log("Signer:", signer.address);

            const gasMoney = ethers.parseEther(taskArgs.gasMoney);

            const tx = await ttmAccount.connect(signer).addMessengerBot(taskArgs.bot, gasMoney);
            const receipt = await tx.wait();
            console.log("Tx:", receipt.hash);
        } catch (error) {
            handleTransactionError(error, ttmAccount);
        }
    });

ACCOUNT_SCOPE.task("bot:remove", "Remove bot from the TTMAccount")
    .addOptionalParam(
        "privateKey",
        "Private key to use, default: TTMACCOUNT_PK env variable",
        process.env.TTMACCOUNT_PK,
    )
    .addOptionalParam(
        "ttmAccount",
        "TTMAccount address, default: TTMACCOUNT_ADDRESS env variable",
        process.env.TTMACCOUNT_ADDRESS,
    )
    .addParam("bot", "Bot address")
    .setAction(async (taskArgs, hre) => {
        const ttmAccount = await getTTMAccount(taskArgs.ttmAccount);
        console.log("TTMAccount:", taskArgs.ttmAccount);
        console.log("Bot:", taskArgs.bot);

        try {
            const signer = new ethers.Wallet(taskArgs.privateKey, ethers.provider);
            console.log("Removing bot from TTMAccount...");
            console.log("Signer:", signer.address);

            const tx = await ttmAccount.connect(signer).removeMessengerBot(taskArgs.bot);
            const receipt = await tx.wait();
            console.log("Tx:", receipt.hash);
        } catch (error) {
            handleTransactionError(error, ttmAccount);
        }
    });

ACCOUNT_SCOPE.task("payment-token:add", "Add payment token to TTMAccount")
    .addOptionalParam(
        "privateKey",
        "Private key to use, default: TTMACCOUNT_PK env variable",
        process.env.TTMACCOUNT_PK,
    )
    .addOptionalParam(
        "ttmAccount",
        "TTMAccount address, default: TTMACCOUNT_ADDRESS env variable",
        process.env.TTMACCOUNT_ADDRESS,
    )
    .addParam("paymentToken", "Payment token address")
    .setAction(async (taskArgs, hre) => {
        const ttmAccount = await getTTMAccount(taskArgs.ttmAccount);
        console.log("TTMAccount:", taskArgs.ttmAccount);
        console.log("Payment Token:", taskArgs.paymentToken);

        try {
            const signer = new ethers.Wallet(taskArgs.privateKey, ethers.provider);
            console.log("Adding payment token to TTMAccount...");
            console.log("Signer:", signer.address);

            const tx = await ttmAccount.connect(signer).addSupportedToken(taskArgs.paymentToken);
            const receipt = await tx.wait();
            console.log("Tx:", receipt.hash);
        } catch (error) {
            handleTransactionError(error, ttmAccount);
        }
    });

ACCOUNT_SCOPE.task("payment-token:remove", "Remove payment token from TTMAccount")
    .addOptionalParam(
        "privateKey",
        "Private key to use, default: TTMACCOUNT_PK env variable",
        process.env.TTMACCOUNT_PK,
    )
    .addOptionalParam(
        "ttmAccount",
        "TTMAccount address, default: TTMACCOUNT_ADDRESS env variable",
        process.env.TTMACCOUNT_ADDRESS,
    )
    .addParam("paymentToken", "Payment token address")
    .setAction(async (taskArgs, hre) => {
        const ttmAccount = await getTTMAccount(taskArgs.ttmAccount);
        console.log("TTMAccount:", taskArgs.ttmAccount);
        console.log("Payment Token:", taskArgs.paymentToken);

        try {
            const signer = new ethers.Wallet(taskArgs.privateKey, ethers.provider);
            console.log("Removing payment token from TTMAccount...");
            console.log("Signer:", signer.address);

            const tx = await ttmAccount.connect(signer).removeSupportedToken(taskArgs.paymentToken);
            const receipt = await tx.wait();
            console.log("Tx:", receipt.hash);
        } catch (error) {
            handleTransactionError(error, ttmAccount);
        }
    });

ACCOUNT_SCOPE.task("payment-token:list", "List supported payment tokens from TTMAccount")
    .addOptionalParam(
        "ttmAccount",
        "TTMAccount address, default: TTMACCOUNT_ADDRESS env variable",
        process.env.TTMACCOUNT_ADDRESS,
    )
    .setAction(async (taskArgs, hre) => {
        console.log("TTMAccount:", taskArgs.ttmAccount, "\n");

        const ttmAccount = await getTTMAccount(taskArgs.ttmAccount);

        const supportedTokens = await ttmAccount.getSupportedTokens();
        console.log("💵 Supported payment tokens:");
        console.log(supportedTokens);

        try {
            const offChainSupported = await ttmAccount.offChainPaymentSupported();
            console.log(`🔗 Off-chain payment supported: ${offChainSupported ? "✅" : "❌"}`);
        } catch (e) {
            console.log("Failed to fetch off-chain payment support info.");
        }
    });

ACCOUNT_SCOPE.task("bot:list", "List all bots from TTMAccount")
    .addOptionalParam(
        "ttmAccount",
        "TTMAccount address, default: TTMACCOUNT_ADDRESS env variable",
        process.env.TTMACCOUNT_ADDRESS,
    )
    .setAction(async (taskArgs, hre) => {
        console.log("TTMAccount:", taskArgs.ttmAccount, "\n");

        const ttmAccount = await getTTMAccount(taskArgs.ttmAccount);
        try {
            const [limit, period] = await ttmAccount.getGasMoneyWithdrawal();
            console.log(`Gas limit: ${ethers.formatEther(limit)} ETH per ${period} seconds.`);
        } catch (e) {
            console.log("Failed to fetch gas limit settings.");
        }

        console.log("\n📢 A bot is an address that has been granted some special roles on the TTMAccount.");

        const role1 = "MESSENGER_BOT_ROLE";
        console.log("\n🤖", role1, "(Can represent the TTMAccount / interact on behalf of it)");
        console.log("======================================================");
        await hre.run({ scope: "account", task: "role:members" }, { role: role1, ttmAccount: taskArgs.ttmAccount });

        const role2 = "BOOKING_OPERATOR_ROLE";
        console.log("\n🤖", role2, "(Can mint and buy Booking Tokens for the TTMAccount)");
        console.log("======================================================");
        await hre.run({ scope: "account", task: "role:members" }, { role: role2, ttmAccount: taskArgs.ttmAccount });

        const role3 = "GAS_WITHDRAWER_ROLE";
        console.log("\n🤖", role3, "(Can withdraw gas from the TTMAccount)");
        console.log("======================================================");
        await hre.run({ scope: "account", task: "role:members" }, { role: role3, ttmAccount: taskArgs.ttmAccount });
    });

ACCOUNT_SCOPE.task("wanted:add", "Add wanted service to TTMAccount")
    .addOptionalParam(
        "privateKey",
        "Private key to use, default: TTMACCOUNT_PK env variable",
        process.env.TTMACCOUNT_PK,
    )
    .addOptionalParam(
        "ttmAccount",
        "TTMAccount address, default: TTMACCOUNT_ADDRESS env variable",
        process.env.TTMACCOUNT_ADDRESS,
    )
    .addParam("serviceName", "Name of service to add")
    .setAction(async (taskArgs, hre) => {
        const ttmAccount = await getTTMAccount(taskArgs.ttmAccount);
        console.log("TTMAccount:", taskArgs.ttmAccount);
        console.log("Service Name:", taskArgs.serviceName);

        try {
            const signer = new ethers.Wallet(taskArgs.privateKey, ethers.provider);

            console.log("Adding service to TTMAccount...");
            console.log("Signer:", signer.address);

            const serviceHash = ethers.keccak256(ethers.toUtf8Bytes(taskArgs.serviceName));
            const tx = await ttmAccount.connect(signer).addWantedServices([serviceHash]);
            const receipt = await tx.wait();
            console.log("Tx:", receipt.hash);
        } catch (error) {
            handleTransactionError(error, ttmAccount);
        }
    });

ACCOUNT_SCOPE.task("wanted:remove", "Remove wanted service from TTMAccount")
    .addOptionalParam(
        "privateKey",
        "Private key to use, default: TTMACCOUNT_PK env variable",
        process.env.TTMACCOUNT_PK,
    )
    .addOptionalParam(
        "ttmAccount",
        "TTMAccount address, default: TTMACCOUNT_ADDRESS env variable",
        process.env.TTMACCOUNT_ADDRESS,
    )
    .addParam("serviceName", "Name of service to remove")
    .setAction(async (taskArgs, hre) => {
        const ttmAccount = await getTTMAccount(taskArgs.ttmAccount);
        console.log("TTMAccount:", taskArgs.ttmAccount);
        console.log("Service Name:", taskArgs.serviceName);

        try {
            const signer = new ethers.Wallet(taskArgs.privateKey, ethers.provider);

            console.log("Removing service from TTMAccount...");
            console.log("Signer:", signer.address);

            const serviceHash = ethers.keccak256(ethers.toUtf8Bytes(taskArgs.serviceName));
            const tx = await ttmAccount.connect(signer).removeWantedServices([serviceHash]);
            const receipt = await tx.wait();
            console.log("Tx:", receipt.hash);
        } catch (error) {
            handleTransactionError(error, ttmAccount);
        }
    });

ACCOUNT_SCOPE.task("wanted:list", "List all wanted service from TTMAccount")
    .addOptionalParam(
        "ttmAccount",
        "TTMAccount address, default: TTMACCOUNT_ADDRESS env variable",
        process.env.TTMACCOUNT_ADDRESS,
    )
    .setAction(async (taskArgs, hre) => {
        const ttmAccount = await getTTMAccount(taskArgs.ttmAccount);
        console.log("TTMAccount:", taskArgs.ttmAccount);

        try {
            console.log("Listing all wanted services from TTMAccount...");

            const manager = await ethers.getContractAt("TTMAccountManager", await ttmAccount.getManagerAddress());
            const wantedServiceHashes = await ttmAccount.getWantedServiceHashes();
            console.log("Wanted Services:");
            for (const serviceHash of wantedServiceHashes) {
                // getServiceNameByHash resolves even a service that was later unregistered
                // from the manager; an empty result just means the hash was never registered.
                const serviceName = await manager.getServiceNameByHash(serviceHash);
                console.log(serviceName ? `📦 ${serviceName} (${serviceHash})` : `📦 ${serviceHash}`);
            }
        } catch (error) {
            handleTransactionError(error, ttmAccount);
        }
    });

ACCOUNT_SCOPE.task("service:add", "Add supported service to TTMAccount")
    .addOptionalParam(
        "privateKey",
        "Private key to use, default: TTMACCOUNT_PK env variable",
        process.env.TTMACCOUNT_PK,
    )
    .addOptionalParam(
        "ttmAccount",
        "TTMAccount address, default: TTMACCOUNT_ADDRESS env variable",
        process.env.TTMACCOUNT_ADDRESS,
    )
    .addParam("serviceName", "Name of service to add")
    .addParam("restrictedRate", "Restricted rate of the service", false, types.boolean)
    .addOptionalParam("capabilities", "Capabilities of the service, comma separated (optional)")
    .setAction(async (taskArgs, hre) => {
        const ttmAccount = await getTTMAccount(taskArgs.ttmAccount);
        console.log("TTMAccount:", taskArgs.ttmAccount);
        console.log("Service Name:", taskArgs.serviceName);
        console.log("Restricted Rate:", taskArgs.restrictedRate);
        console.log("Capabilities:", taskArgs.capabilities);

        try {
            const signer = new ethers.Wallet(taskArgs.privateKey, ethers.provider);

            const capabilities = taskArgs.capabilities ? taskArgs.capabilities.split(",") : [];
            const serviceHash = ethers.keccak256(ethers.toUtf8Bytes(taskArgs.serviceName));

            console.log("Adding service to TTMAccount...");
            console.log("Signer:", signer.address);

            const tx = await ttmAccount.connect(signer).addService(serviceHash, taskArgs.restrictedRate, capabilities);
            const receipt = await tx.wait();
            console.log("Tx:", receipt.hash);
        } catch (error) {
            handleTransactionError(error, ttmAccount);
        }
    });

ACCOUNT_SCOPE.task("service:remove", "Remove wanted service from TTMAccount")
    .addOptionalParam(
        "privateKey",
        "Private key to use, default: TTMACCOUNT_PK env variable",
        process.env.TTMACCOUNT_PK,
    )
    .addOptionalParam(
        "ttmAccount",
        "TTMAccount address, default: TTMACCOUNT_ADDRESS env variable",
        process.env.TTMACCOUNT_ADDRESS,
    )
    .addParam("serviceName", "Name of service to remove")
    .setAction(async (taskArgs, hre) => {
        const ttmAccount = await getTTMAccount(taskArgs.ttmAccount);
        console.log("TTMAccount:", taskArgs.ttmAccount);
        console.log("Service Name:", taskArgs.serviceName);

        try {
            const signer = new ethers.Wallet(taskArgs.privateKey, ethers.provider);

            console.log("Removing service from TTMAccount...");
            console.log("Signer:", signer.address);

            const serviceHash = ethers.keccak256(ethers.toUtf8Bytes(taskArgs.serviceName));
            const tx = await ttmAccount.connect(signer).removeService(serviceHash);
            const receipt = await tx.wait();
            console.log("Tx:", receipt.hash);
        } catch (error) {
            handleTransactionError(error, ttmAccount);
        }
    });

ACCOUNT_SCOPE.task("service:list", "List supported services from TTMAccount")
    .addOptionalParam(
        "ttmAccount",
        "TTMAccount address, default: TTMACCOUNT_ADDRESS env variable",
        process.env.TTMACCOUNT_ADDRESS,
    )
    .setAction(async (taskArgs, hre) => {
        const ttmAccount = await getTTMAccount(taskArgs.ttmAccount);
        console.log("TTMAccount:", taskArgs.ttmAccount);

        try {
            console.log("Listing all supported services from TTMAccount...");

            // getSupportedServices() returns hashes, not names — resolve each via the
            // manager, the same as `wanted:list`. getServiceNameByHash resolves even a
            // service that was later unregistered from the manager; an empty result
            // just means the hash was never registered.
            const manager = await ethers.getContractAt("TTMAccountManager", await ttmAccount.getManagerAddress());
            const supportedServices = await ttmAccount.getSupportedServices();
            const serviceHashes = supportedServices[0];
            const serviceDetails = supportedServices[1];
            if (serviceHashes.length > 0) {
                console.log("Supported Services:");
                for (let i = 0; i < serviceHashes.length; i++) {
                    const serviceName = await manager.getServiceNameByHash(serviceHashes[i]);
                    console.log(`📦 ${serviceName ? `${serviceName} (${serviceHashes[i]})` : serviceHashes[i]}`);
                    const restrictedRate =
                        serviceDetails[i]._restrictedRate !== undefined
                            ? serviceDetails[i]._restrictedRate
                            : serviceDetails[i][0];
                    const capabilities =
                        serviceDetails[i]._capabilities !== undefined
                            ? serviceDetails[i]._capabilities
                            : serviceDetails[i][1];
                    console.log(`\t🔒 Restricted Rate: ${restrictedRate} ${restrictedRate ? "✅" : "❌"}`);

                    for (let j = 0; j < capabilities.length; j++) {
                        console.log(`\t🔧 ${capabilities[j]}`);
                    }
                }
            } else {
                console.log("🛑 TTM Account does not have any supported services!");
            }
        } catch (error) {
            handleTransactionError(error, ttmAccount);
        }
    });

ACCOUNT_SCOPE.task("upgrade", "Upgrade TTMAccount to latest implementation")
    .addOptionalParam(
        "privateKey",
        "Private key to use, default: TTMACCOUNT_PK env variable",
        process.env.TTMACCOUNT_PK,
    )
    .addOptionalParam(
        "ttmAccount",
        "TTMAccount address, default: TTMACCOUNT_ADDRESS env variable",
        process.env.TTMACCOUNT_ADDRESS,
    )
    .setAction(async (taskArgs, hre) => {
        const ttmAccount = await getTTMAccount(taskArgs.ttmAccount);
        console.log("TTMAccount:", taskArgs.ttmAccount);

        // Get new implementation
        const manager = await getManager(hre);
        const currentImplOnManager = await manager.getAccountImplementation();
        console.log("Implementation on the Manager  :", currentImplOnManager);

        // Get implementation on TTMAccount
        const implementation = await getImplementationAddressForProxy(await ttmAccount.getAddress());
        console.log("Implementation on the TTMAccount:", implementation);

        if (implementation === currentImplOnManager) {
            console.log("✅ TTMAccount is using the latest implementation!");
            console.log("No need for upgrade!");
            return;
        } else {
            console.log("⏫ There is an upgrade available for the TTMAccount! Starting upgrade...");
        }

        try {
            const signer = new ethers.Wallet(taskArgs.privateKey, ethers.provider);
            console.log("Upgrading TTMAccount implementation...");
            console.log("Signer:", signer.address);
            const tx = await ttmAccount.connect(signer).upgradeToAndCall(currentImplOnManager, "0x");
            const receipt = await tx.wait();
            console.log("Tx:", receipt.hash);
        } catch (error) {
            handleTransactionError(error, ttmAccount);
        }
    });

ACCOUNT_SCOPE.task("upgrade:check", "Check if TTMAccount is upgradable")
    .addOptionalParam(
        "ttmAccount",
        "TTMAccount address, default: TTMACCOUNT_ADDRESS env variable",
        process.env.TTMACCOUNT_ADDRESS,
    )
    .setAction(async (taskArgs, hre) => {
        const manager = await getManager(hre);
        console.log("Manager:", await manager.getAddress());
        const ttmAccount = await getTTMAccount(taskArgs.ttmAccount);
        console.log("TTMAccount:", taskArgs.ttmAccount);

        console.log(`📜 ${bold("Implementation:")}`);
        const implementation = await getImplementationAddressForProxy(await ttmAccount.getAddress());

        const currentImplOnManager = await manager.getAccountImplementation();

        // Create visual diff of the addresses using ANSI red color code
        const addressDiff = Array.from(implementation)
            .map((char, i) => {
                return char === currentImplOnManager[i] ? char : `\x1b[31m${char}\x1b[0m`;
            })
            .join("");

        console.log(`${bold("   - Active:")}`, addressDiff);
        console.log(`${bold("   - Latest:")}`, currentImplOnManager);

        if (implementation !== currentImplOnManager) {
            console.log("⏫ TTMAccount needs an upgrade!");
        } else {
            console.log("✅ TTMAccount is using the latest implementation!");
        }
    });

ACCOUNT_SCOPE.task("find", "Scan all TTM Accounts for roles of a given address")
    .addParam("address", "Address to search for")
    .setAction(async (taskArgs, hre) => {
        const manager = await getManager(hre);
        const ttmAccounts = await manager.getTTMAccounts();
        const count = ttmAccounts.length;

        console.log(`🔍 Searching for address: ${bold(taskArgs.address)}`);
        console.log(`📡 Found ${bold(count)} TTM Accounts. Starting scan...\n`);

        const spinners = ["⠋", "⠙", "⠹", "⠸", "⠼", "⠴", "⠦", "⠧", "⠇", "⠏"];
        let spinnerIdx = 0;

        const formatProgress = (current, total, address) => {
            const percentage = Math.floor((Number(current) / Number(total)) * 100);
            const spinner = spinners[spinnerIdx % spinners.length];
            spinnerIdx++;
            return `\r${spinner} [${current}/${total}] ${percentage}% | Scanning: ${address.slice(0, 10)}...${address.slice(-8)}`;
        };

        let foundCount = 0;
        const findings = [];

        for (let i = 0; i < count; i++) {
            const ttmAccountAddress = ttmAccounts[i];

            // Update progress on every account
            process.stdout.write(formatProgress(i + 1, count, ttmAccountAddress));

            const ttmAccount = await getTTMAccount(ttmAccountAddress);

            const roleChecks = await Promise.all(
                ACC_ROLES.map(async (roleName) => {
                    try {
                        const roleHash = await ttmAccount[roleName]();
                        const hasRole = await ttmAccount.hasRole(roleHash, taskArgs.address);
                        return hasRole ? roleName : null;
                    } catch (e) {
                        if (e?.message?.includes("is not a function")) {
                            return null;
                        }
                        throw new Error(
                            `Role scan failed for TTMAccount ${ttmAccountAddress}, role ${roleName}: ${e?.message || e}`,
                        );
                    }
                }),
            );

            const rolesHeld = roleChecks.filter((r) => r !== null);

            if (rolesHeld.length > 0) {
                foundCount++;
                findings.push({ address: ttmAccountAddress, roles: rolesHeld });
                // If we found something, clear the current progress line and print it
                process.stdout.write("\r" + " ".repeat(80) + "\r"); // Clear line
                console.log(`✨ ${bold("Found match:")} ${bold(ttmAccountAddress)}`);
                rolesHeld.forEach((r) => console.log(`   └─ ${r}`));
                console.log();
            }
        }

        // Final clear of the progress line
        process.stdout.write("\r" + " ".repeat(80) + "\r");

        if (foundCount === 0) {
            console.log("❌ No TTM Accounts found where this address holds a role.");
        } else {
            console.log(`✅ Search complete. Found matches in ${bold(foundCount)} TTM Account(s).`);
        }
    });

ACCOUNT_SCOPE.task("withdraw:erc20", "Withdraw ERC20 tokens from TTMAccount")
    .addOptionalParam(
        "privateKey",
        "Private key to use, default: TTMACCOUNT_PK env variable",
        process.env.TTMACCOUNT_PK,
    )
    .addOptionalParam(
        "ttmAccount",
        "TTMAccount address, default: TTMACCOUNT_ADDRESS env variable",
        process.env.TTMACCOUNT_ADDRESS,
    )
    .addParam("token", "ERC20 token address")
    .addParam("recipient", "Recipient address")
    .addParam("amount", "Amount to withdraw in human readable units")
    .setAction(async (taskArgs, hre) => {
        const ttmAccount = await getTTMAccount(taskArgs.ttmAccount);
        console.log("TTMAccount:", taskArgs.ttmAccount);
        console.log("Token:", taskArgs.token);
        console.log("Recipient:", taskArgs.recipient);

        try {
            const tokenContract = await ethers.getContractAt(
                ["function decimals() view returns (uint8)"],
                taskArgs.token,
            );
            const decimals = await tokenContract.decimals();
            const amountWei = ethers.parseUnits(taskArgs.amount, decimals);
            console.log(`Withdrawing ${taskArgs.amount} tokens (wei: ${amountWei.toString()})...`);

            const signer = new ethers.Wallet(taskArgs.privateKey, ethers.provider);
            const tx = await ttmAccount.connect(signer).transferERC20(taskArgs.token, taskArgs.recipient, amountWei);
            const receipt = await tx.wait();
            console.log("Tx:", receipt.hash);
        } catch (error) {
            handleTransactionError(error, ttmAccount);
        }
    });

ACCOUNT_SCOPE.task("withdraw:erc721", "Withdraw ERC721 tokens from TTMAccount")
    .addOptionalParam(
        "privateKey",
        "Private key to use, default: TTMACCOUNT_PK env variable",
        process.env.TTMACCOUNT_PK,
    )
    .addOptionalParam(
        "ttmAccount",
        "TTMAccount address, default: TTMACCOUNT_ADDRESS env variable",
        process.env.TTMACCOUNT_ADDRESS,
    )
    .addParam("token", "ERC721 token address")
    .addParam("recipient", "Recipient address")
    .addParam("tokenId", "Token ID to withdraw")
    .setAction(async (taskArgs, hre) => {
        const ttmAccount = await getTTMAccount(taskArgs.ttmAccount);
        console.log("TTMAccount:", taskArgs.ttmAccount);
        console.log("Token:", taskArgs.token);
        console.log("Recipient:", taskArgs.recipient);
        console.log("Token ID:", taskArgs.tokenId);

        try {
            const signer = new ethers.Wallet(taskArgs.privateKey, ethers.provider);
            const tx = await ttmAccount
                .connect(signer)
                .transferERC721(taskArgs.token, taskArgs.recipient, taskArgs.tokenId);
            const receipt = await tx.wait();
            console.log("Tx:", receipt.hash);
        } catch (error) {
            handleTransactionError(error, ttmAccount);
        }
    });

ACCOUNT_SCOPE.task("bot:withdraw-gas", "Withdraw gas money for a bot from TTMAccount")
    .addParam("privateKey", "Private key of the bot")
    .addOptionalParam(
        "ttmAccount",
        "TTMAccount address, default: TTMACCOUNT_ADDRESS env variable",
        process.env.TTMACCOUNT_ADDRESS,
    )
    .addParam("amount", "Amount to withdraw in ETH (e.g. 0.5)")
    .setAction(async (taskArgs, hre) => {
        const ttmAccount = await getTTMAccount(taskArgs.ttmAccount);
        console.log("TTMAccount:", taskArgs.ttmAccount);

        try {
            const amountWei = ethers.parseEther(taskArgs.amount);
            const signer = new ethers.Wallet(taskArgs.privateKey, ethers.provider);
            console.log(`Bot address: ${signer.address}`);
            console.log(`Withdrawing ${taskArgs.amount} ETH (wei: ${amountWei.toString()})...`);
            const tx = await ttmAccount.connect(signer).withdrawGasMoney(amountWei);
            const receipt = await tx.wait();
            console.log("Tx:", receipt.hash);
        } catch (error) {
            handleTransactionError(error, ttmAccount);
        }
    });

ACCOUNT_SCOPE.task("bot:set-gas-limit", "Set gas money withdrawal limit and period for bots")
    .addOptionalParam(
        "privateKey",
        "Private key to use, default: TTMACCOUNT_PK env variable",
        process.env.TTMACCOUNT_PK,
    )
    .addOptionalParam(
        "ttmAccount",
        "TTMAccount address, default: TTMACCOUNT_ADDRESS env variable",
        process.env.TTMACCOUNT_ADDRESS,
    )
    .addParam("limit", "Withdrawal limit in ETH (e.g. 10)")
    .addParam("period", "Withdrawal period in seconds (e.g. 86400 for 24h)")
    .setAction(async (taskArgs, hre) => {
        const ttmAccount = await getTTMAccount(taskArgs.ttmAccount);
        console.log("TTMAccount:", taskArgs.ttmAccount);

        try {
            const limitWei = ethers.parseEther(taskArgs.limit);
            console.log(
                `Setting gas limit to ${taskArgs.limit} ETH (wei: ${limitWei.toString()}) per ${taskArgs.period} seconds...`,
            );
            const signer = new ethers.Wallet(taskArgs.privateKey, ethers.provider);
            const tx = await ttmAccount.connect(signer).setGasMoneyWithdrawal(limitWei, taskArgs.period);
            const receipt = await tx.wait();
            console.log("Tx:", receipt.hash);
        } catch (error) {
            handleTransactionError(error, ttmAccount);
        }
    });

ACCOUNT_SCOPE.task("pubkey:add", "Add public key with address for off-chain encryption")
    .addOptionalParam(
        "privateKey",
        "Private key to use, default: TTMACCOUNT_PK env variable",
        process.env.TTMACCOUNT_PK,
    )
    .addOptionalParam(
        "ttmAccount",
        "TTMAccount address, default: TTMACCOUNT_ADDRESS env variable",
        process.env.TTMACCOUNT_ADDRESS,
    )
    .addParam("pubkeyAddress", "Address of the public key")
    .addParam("pubkeyData", "Public key data in hex format (must start with 0x)")
    .setAction(async (taskArgs, hre) => {
        const ttmAccount = await getTTMAccount(taskArgs.ttmAccount);
        console.log("TTMAccount:", taskArgs.ttmAccount);

        try {
            const signer = new ethers.Wallet(taskArgs.privateKey, ethers.provider);
            console.log(`Adding public key for address ${taskArgs.pubkeyAddress}...`);
            const tx = await ttmAccount.connect(signer).addPublicKey(taskArgs.pubkeyAddress, taskArgs.pubkeyData);
            const receipt = await tx.wait();
            console.log("Tx:", receipt.hash);
        } catch (error) {
            handleTransactionError(error, ttmAccount);
        }
    });

ACCOUNT_SCOPE.task("pubkey:remove", "Remove public key by address")
    .addOptionalParam(
        "privateKey",
        "Private key to use, default: TTMACCOUNT_PK env variable",
        process.env.TTMACCOUNT_PK,
    )
    .addOptionalParam(
        "ttmAccount",
        "TTMAccount address, default: TTMACCOUNT_ADDRESS env variable",
        process.env.TTMACCOUNT_ADDRESS,
    )
    .addParam("pubkeyAddress", "Address of the public key to remove")
    .setAction(async (taskArgs, hre) => {
        const ttmAccount = await getTTMAccount(taskArgs.ttmAccount);
        console.log("TTMAccount:", taskArgs.ttmAccount);

        try {
            const signer = new ethers.Wallet(taskArgs.privateKey, ethers.provider);
            console.log(`Removing public key for address ${taskArgs.pubkeyAddress}...`);
            const tx = await ttmAccount.connect(signer).removePublicKey(taskArgs.pubkeyAddress);
            const receipt = await tx.wait();
            console.log("Tx:", receipt.hash);
        } catch (error) {
            handleTransactionError(error, ttmAccount);
        }
    });

ACCOUNT_SCOPE.task("pubkey:list", "List all public keys registered on TTMAccount")
    .addOptionalParam(
        "ttmAccount",
        "TTMAccount address, default: TTMACCOUNT_ADDRESS env variable",
        process.env.TTMACCOUNT_ADDRESS,
    )
    .setAction(async (taskArgs, hre) => {
        const ttmAccount = await getTTMAccount(taskArgs.ttmAccount);
        console.log("TTMAccount:", taskArgs.ttmAccount);

        try {
            const addresses = await ttmAccount.getPublicKeysAddresses();
            console.log("🔑 Registered Public Keys:");
            for (const addr of addresses) {
                const pubkeyData = await ttmAccount.getPublicKey(addr);
                console.log(`Address: ${addr}`);
                console.log(`Data   : ${pubkeyData}\n`);
            }
            if (addresses.length === 0) {
                console.log("No public keys registered.");
            }
        } catch (error) {
            handleTransactionError(error, ttmAccount);
        }
    });

ACCOUNT_SCOPE.task("payment:set-offchain", "Set if off-chain payment is supported by TTMAccount")
    .addOptionalParam(
        "privateKey",
        "Private key to use, default: TTMACCOUNT_PK env variable",
        process.env.TTMACCOUNT_PK,
    )
    .addOptionalParam(
        "ttmAccount",
        "TTMAccount address, default: TTMACCOUNT_ADDRESS env variable",
        process.env.TTMACCOUNT_ADDRESS,
    )
    .addParam("supported", "true if supported, false otherwise", null, types.boolean)
    .setAction(async (taskArgs, hre) => {
        const ttmAccount = await getTTMAccount(taskArgs.ttmAccount);
        console.log("TTMAccount:", taskArgs.ttmAccount);

        try {
            const signer = new ethers.Wallet(taskArgs.privateKey, ethers.provider);
            console.log(`Setting off-chain payment support to: ${taskArgs.supported}`);
            const tx = await ttmAccount.connect(signer).setOffChainPaymentSupported(taskArgs.supported);
            const receipt = await tx.wait();
            console.log("Tx:", receipt.hash);
        } catch (error) {
            handleTransactionError(error, ttmAccount);
        }
    });

ACCOUNT_SCOPE.task("service:remove-all", "Remove all supported services from TTMAccount")
    .addOptionalParam(
        "privateKey",
        "Private key to use, default: TTMACCOUNT_PK env variable",
        process.env.TTMACCOUNT_PK,
    )
    .addOptionalParam(
        "ttmAccount",
        "TTMAccount address, default: TTMACCOUNT_ADDRESS env variable",
        process.env.TTMACCOUNT_ADDRESS,
    )
    .setAction(async (taskArgs, hre) => {
        const ttmAccount = await getTTMAccount(taskArgs.ttmAccount);
        console.log("TTMAccount:", taskArgs.ttmAccount);

        try {
            const signer = new ethers.Wallet(taskArgs.privateKey, ethers.provider);
            console.log("Removing all supported services from TTMAccount...");
            const tx = await ttmAccount.connect(signer).removeAllServices();
            const receipt = await tx.wait();
            console.log("Tx:", receipt.hash);
        } catch (error) {
            handleTransactionError(error, ttmAccount);
        }
    });

ACCOUNT_SCOPE.task("service:set-restricted-rate", "Set the restricted rate property of a supported service")
    .addOptionalParam(
        "privateKey",
        "Private key to use, default: TTMACCOUNT_PK env variable",
        process.env.TTMACCOUNT_PK,
    )
    .addOptionalParam(
        "ttmAccount",
        "TTMAccount address, default: TTMACCOUNT_ADDRESS env variable",
        process.env.TTMACCOUNT_ADDRESS,
    )
    .addParam("serviceName", "Name of the service")
    .addParam("restrictedRate", "Restricted rate status (true/false)", null, types.boolean)
    .setAction(async (taskArgs, hre) => {
        const ttmAccount = await getTTMAccount(taskArgs.ttmAccount);
        console.log("TTMAccount:", taskArgs.ttmAccount);

        try {
            const signer = new ethers.Wallet(taskArgs.privateKey, ethers.provider);
            console.log(`Setting restricted rate of service ${taskArgs.serviceName} to ${taskArgs.restrictedRate}...`);
            const serviceHash = ethers.keccak256(ethers.toUtf8Bytes(taskArgs.serviceName));
            const tx = await ttmAccount.connect(signer).setServiceRestrictedRate(serviceHash, taskArgs.restrictedRate);
            const receipt = await tx.wait();
            console.log("Tx:", receipt.hash);
        } catch (error) {
            handleTransactionError(error, ttmAccount);
        }
    });

ACCOUNT_SCOPE.task("service:set-capabilities", "Set all capabilities of a supported service")
    .addOptionalParam(
        "privateKey",
        "Private key to use, default: TTMACCOUNT_PK env variable",
        process.env.TTMACCOUNT_PK,
    )
    .addOptionalParam(
        "ttmAccount",
        "TTMAccount address, default: TTMACCOUNT_ADDRESS env variable",
        process.env.TTMACCOUNT_ADDRESS,
    )
    .addParam("serviceName", "Name of the service")
    .addParam("capabilities", "Comma-separated capabilities")
    .setAction(async (taskArgs, hre) => {
        const ttmAccount = await getTTMAccount(taskArgs.ttmAccount);
        console.log("TTMAccount:", taskArgs.ttmAccount);

        try {
            const capabilities = taskArgs.capabilities ? taskArgs.capabilities.split(",") : [];
            const signer = new ethers.Wallet(taskArgs.privateKey, ethers.provider);
            console.log(`Setting capabilities of service ${taskArgs.serviceName} to:`, capabilities);
            const serviceHash = ethers.keccak256(ethers.toUtf8Bytes(taskArgs.serviceName));
            const tx = await ttmAccount.connect(signer).setServiceCapabilities(serviceHash, capabilities);
            const receipt = await tx.wait();
            console.log("Tx:", receipt.hash);
        } catch (error) {
            handleTransactionError(error, ttmAccount);
        }
    });

module.exports = {};
