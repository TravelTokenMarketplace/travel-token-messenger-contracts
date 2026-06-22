const { ethers } = require("hardhat");

async function main() {
    const [deployer] = await ethers.getSigners();
    console.log(`Using address: ${deployer.address}`);

    const chainId = await deployer.provider.getNetwork().then((n) => n.chainId);
    console.log(`Connected to chain: ${chainId}`);

    const addresses = require(`../ignition/deployments/chain-${chainId}/deployed_addresses.json`);
    const bookingTokenProxyAddress = addresses["CaminoMessengerModule#BookingTokenProxy"];
    console.log(`BookingTokenProxy address: ${bookingTokenProxyAddress}`);

    const bookingToken = await ethers.getContractAt("BookingToken", bookingTokenProxyAddress);

    // Read current details
    const currentName = await bookingToken.name();
    const currentSymbol = await bookingToken.symbol();
    console.log(`Current Token Name: ${currentName}`);
    console.log(`Current Token Symbol: ${currentSymbol}`);

    if (currentSymbol === "BToken") {
        console.log("Token symbol is already BToken. No action needed.");
        return;
    }

    const newName = currentName || "BookingToken";
    const newSymbol = "BToken";

    console.log(`Reinitializing BookingToken to name: "${newName}" and symbol: "${newSymbol}"...`);
    const tx = await bookingToken.reinitializeV2(newName, newSymbol);
    console.log(`Transaction sent: ${tx.hash}`);

    console.log("Waiting for confirmation...");
    await tx.wait();
    console.log("Transaction confirmed!");

    // Verify change
    const updatedName = await bookingToken.name();
    const updatedSymbol = await bookingToken.symbol();
    console.log(`Updated Token Name: ${updatedName}`);
    console.log(`Updated Token Symbol: ${updatedSymbol}`);
}

main()
    .then(() => process.exit(0))
    .catch((error) => {
        console.error(error);
        process.exit(1);
    });
