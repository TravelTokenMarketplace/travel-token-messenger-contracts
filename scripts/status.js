async function main() {
    const [deployer] = await ethers.getSigners();

    const chainId = await deployer.provider.getNetwork().then((n) => n.chainId);

    const addresses = require(`../ignition/deployments/chain-${chainId}/deployed_addresses.json`);

    const managerImplementation = await ethers.getContractAt(
        "TTMAccountManager",
        addresses["TravelTokenMessengerModule#TTMAccountManager"],
    );

    const manager = await ethers.getContractAt(
        "TTMAccountManager",
        addresses["TravelTokenMessengerModule#ManagerProxy"],
    );

    const ttmAccount = await ethers.getContractAt("TTMAccount", addresses["TravelTokenMessengerModule#TTMAccount"]);

    const bookingTokenImplementation = await ethers.getContractAt(
        "BookingToken",
        addresses["TravelTokenMessengerModule#BookingToken"],
    );

    const bookingToken = await ethers.getContractAt(
        "BookingToken",
        addresses["TravelTokenMessengerModule#BookingTokenProxy"],
    );

    console.log("========================= MANAGER =========================");
    console.log(`Proxy: ${await manager.getAddress()}`);
    console.log(`Implementation: ${await managerImplementation.getAddress()}`);

    console.log();
    console.log("======================== TTM ACCOUNT ========================");
    console.log(`Implementation: ${await ttmAccount.getAddress()}`);

    console.log();
    console.log("====================== BOOKING TOKEN ======================");
    console.log(`Proxy: ${await bookingToken.getAddress()}`);
    console.log(`Implementation: ${await bookingTokenImplementation.getAddress()}`);

    console.log();
    console.log("====================== CONFIGURATION ======================");
    console.log(`TTM Account Impl: ${await manager.getAccountImplementation()}`);
    console.log(`Developer Wallet: ${await manager.getDeveloperWallet()}`);
    console.log(`Fee Basis Points: ${await manager.getDeveloperFeeBp()}`);
    console.log(`Prefund Amount: ${ethers.formatEther(await manager.getPrefundAmount())} CAM`);
}

main()
    .then(() => process.exit(0))
    .catch((error) => {
        console.error(error);
        process.exit(1);
    });
