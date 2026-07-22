/**
 * @dev Fixtures
 */
const { loadFixture } = require("@nomicfoundation/hardhat-toolbox/network-helpers");
const { ethers, upgrades } = require("hardhat");

/**
 * @dev Returns the keccak256 hash of a service name, matching the on-chain
 * `keccak256(abi.encodePacked(serviceName))` computation used by ServiceRegistry
 * and PartnerConfiguration.
 */
const serviceHash = (name) => ethers.keccak256(ethers.toUtf8Bytes(name));

async function setupSigners() {
    const [
        managerAdmin,
        managerPauser,
        managerUpgrader,
        managerVersioner,
        ttmAccountAdmin,
        ttmAccountUpgrader,
        ttmServiceAdmin,
        botOperator,
        depositor,
        withdrawer,
        btAdmin,
        btUpgrader,
        registryAdmin,
        otherAccount1,
        otherAccount2,
        otherAccount3,
        otherAccount4,
    ] = await ethers.getSigners();

    signers = {
        managerAdmin,
        managerPauser,
        managerUpgrader,
        managerVersioner,
        ttmAccountAdmin,
        ttmAccountUpgrader,
        ttmServiceAdmin,
        botOperator,
        depositor,
        withdrawer,
        btAdmin,
        btUpgrader,
        registryAdmin,
        otherAccount1,
        otherAccount2,
        otherAccount3,
        otherAccount4,
    };
}

// Deploy NullUSD
async function deployNullUSDFixture() {
    await setupSigners();

    const NullUSD = await ethers.getContractFactory("NullUSD");
    const nullUSD = await NullUSD.deploy();

    const nullUSDDecimals = await nullUSD.decimals();

    return { nullUSD, nullUSDDecimals };
}

async function deployTTMAccountManagerFixture() {
    // Set up signers
    await setupSigners();

    const TTMAccountManager = await ethers.getContractFactory("TTMAccountManager");
    const ttmAccountManager = await upgrades.deployProxy(
        TTMAccountManager,
        [
            signers.managerAdmin.address,
            signers.managerPauser.address,
            signers.managerUpgrader.address,
            signers.managerVersioner.address,
        ],
        { kind: "uups" },
    );
    return { ttmAccountManager };
}

async function deployTTMAccountImplFixture() {
    const BookingTokenOperator = await ethers.getContractFactory("BookingTokenOperator");
    const bookingTokenOperator = await BookingTokenOperator.deploy();
    const TTMAccount = await ethers.getContractFactory("TTMAccount", {
        libraries: { BookingTokenOperator: await bookingTokenOperator.getAddress() },
    });
    const ttmAccountImpl = await TTMAccount.deploy();
    await ttmAccountImpl.waitForDeployment();

    return { ttmAccountImpl };
}

async function deployTTMAccountManagerWithTTMAccountImplFixture() {
    // Set up signers
    await setupSigners();

    const { ttmAccountManager } = await loadFixture(deployTTMAccountManagerFixture);
    const { ttmAccountImpl } = await loadFixture(deployTTMAccountImplFixture);

    const ttmAccountImplAddress = await ttmAccountImpl.getAddress();

    await ttmAccountManager.grantRole(await ttmAccountManager.VERSIONER_ROLE(), signers.managerVersioner.address);
    await ttmAccountManager.connect(signers.managerVersioner).setAccountImplementation(ttmAccountImplAddress);

    return { ttmAccountManager, ttmAccountImplAddress };
}

async function deployAndConfigureAllFixture() {
    // Set up signers
    await setupSigners();

    const { ttmAccountManager, ttmAccountImplAddress } = await loadFixture(
        deployTTMAccountManagerWithTTMAccountImplFixture,
    );

    const { nullUSD, nullUSDDecimals } = await loadFixture(deployNullUSDFixture);

    // Deploy BookingToken

    const BookingToken = await ethers.getContractFactory("BookingToken");
    const bookingToken = await upgrades.deployProxy(
        BookingToken,
        [await ttmAccountManager.getAddress(), signers.btAdmin.address, signers.btUpgrader.address],
        { kind: "uups" },
    );

    // Set BookingToken address on the manager
    await ttmAccountManager.connect(signers.managerVersioner).setBookingTokenAddress(bookingToken.getAddress());

    const tx = await ttmAccountManager.createTTMAccount(
        signers.ttmAccountAdmin.address,
        signers.ttmAccountUpgrader.address,
        { value: ethers.parseEther("100") },
    );

    const receipt = await tx.wait();

    // Parse event to get the TTMAccount address (this is the UUPS proxy address)
    const event = receipt.logs.find((log) => {
        try {
            return ttmAccountManager.interface.parseLog(log).name === "TTMAccountCreated";
        } catch (e) {
            return false;
        }
    });

    const parsedEvent = ttmAccountManager.interface.parseLog(event);
    const ttmAccountAddress = parsedEvent.args.account;

    // Get the TTMAccount instance at the address
    const ttmAccount = await ethers.getContractAt("TTMAccount", ttmAccountAddress);

    return {
        ttmAccountManager,
        ttmAccount,
        bookingToken,
        nullUSD,
        nullUSDDecimals,
    };
}

async function deployTTMAccountWithDepositFixture() {
    // Set up signers
    await setupSigners();

    const { ttmAccountManager, ttmAccount, bookingToken, nullUSD, nullUSDDecimals } =
        await loadFixture(deployAndConfigureAllFixture);

    // Grant withdrawer role
    const WITHDRAWER_ROLE = await ttmAccount.WITHDRAWER_ROLE();
    await ttmAccount.connect(signers.ttmAccountAdmin).grantRole(WITHDRAWER_ROLE, signers.withdrawer.address);

    // Deposit ETH
    const depositAmount = ethers.parseEther("1");

    const depositTx = {
        to: ttmAccount.getAddress(),
        value: depositAmount,
    };

    const txResponse = await signers.depositor.sendTransaction(depositTx);
    await txResponse.wait();

    // Deposit NullUSD
    await nullUSD.transfer(ttmAccount.getAddress(), ethers.parseUnits("1", nullUSDDecimals));

    return {
        ttmAccountManager,
        ttmAccount,
        bookingToken,
        nullUSD,
        nullUSDDecimals,
    };
}

async function deployBookingTokenFixture() {
    // Set up signers
    await setupSigners();

    const { ttmAccountManager, ttmAccount, bookingToken, nullUSD, nullUSDDecimals } = await loadFixture(
        deployTTMAccountWithDepositFixture,
    );

    // Supplier TTMAccount with deposit
    const supplierTTMAccount = ttmAccount;

    // Create distributor TTMAccount
    const tx = await ttmAccountManager.createTTMAccount(
        signers.ttmAccountAdmin.address,
        signers.ttmAccountUpgrader.address,
        { value: ethers.parseEther("100") },
    );

    const receipt = await tx.wait();

    // Parse event to get the TTMAccount address (this is the UUPS proxy address)
    const event = receipt.logs.find((log) => {
        try {
            return ttmAccountManager.interface.parseLog(log).name === "TTMAccountCreated";
        } catch (e) {
            return false;
        }
    });

    const parsedEvent = ttmAccountManager.interface.parseLog(event);
    const distributorTTMAccountAddress = parsedEvent.args.account;

    // Get the TTMAccount instance at the address
    const distributorTTMAccount = await ethers.getContractAt("TTMAccount", distributorTTMAccountAddress);

    // Deposit funds to distributor TTMAccount
    const depositAmount = ethers.parseEther("1");
    const depositTx = {
        to: distributorTTMAccount.getAddress(),
        value: depositAmount,
    };
    const txResponse = await signers.depositor.sendTransaction(depositTx);
    await txResponse.wait();

    // Deposit NullUSD
    await nullUSD.transfer(distributorTTMAccount.getAddress(), ethers.parseUnits("1", nullUSDDecimals));

    return {
        ttmAccountManager,
        supplierTTMAccount,
        distributorTTMAccount,
        bookingToken,
        nullUSD,
        nullUSDDecimals,
    };
}

async function deployBookingTokenWithNullUSDFixture() {
    // Set up signers
    await setupSigners();

    const { ttmAccountManager, supplierTTMAccount, distributorTTMAccount, bookingToken, nullUSD, nullUSDDecimals } =
        await loadFixture(deployBookingTokenFixture);

    // Fund NullUSD to the TTM accounts
    const fundAmount = ethers.parseEther("1000");
    await nullUSD.transfer(await supplierTTMAccount.getAddress(), fundAmount);
    await nullUSD.transfer(await distributorTTMAccount.getAddress(), fundAmount);

    return {
        ttmAccountManager,
        supplierTTMAccount,
        distributorTTMAccount,
        bookingToken,
        nullUSD,
        nullUSDDecimals,
    };
}

async function deployCancellationSupportFixture() {
    // Set up signers
    await setupSigners();

    const { ttmAccountManager, supplierTTMAccount, distributorTTMAccount, bookingToken, nullUSD, nullUSDDecimals } =
        await loadFixture(deployBookingTokenWithNullUSDFixture);

    // Set accounts
    const otherBookingOperator = signers.otherAccount1;
    const supplierBookingOperator = signers.otherAccount2;
    const distributorBookingOperator = signers.otherAccount3;

    // Set BOOKING_OPERATOR_ROLE
    // Supplier
    const BOOKING_OPERATOR_ROLE = await supplierTTMAccount.BOOKING_OPERATOR_ROLE();
    await supplierTTMAccount
        .connect(signers.ttmAccountAdmin)
        .grantRole(BOOKING_OPERATOR_ROLE, supplierBookingOperator.address);

    // Distributor
    await distributorTTMAccount
        .connect(signers.ttmAccountAdmin)
        .grantRole(BOOKING_OPERATOR_ROLE, distributorBookingOperator.address);

    // Mint BOOKING TOKEN with NATIVE PAYMENT -----------------------------------------------------------------------

    const tokenURI = "data:application/json;base64,eyJuYW1lIjoiQ2FtaW5vIE1lc3NlbmdlciBCb29raW5nVG9rZW4gVGVzdCJ9Cg==";
    const expirationTimestamp = Math.floor(Date.now() / 1000) + 120;
    const price = ethers.parseEther("0.05");

    await supplierTTMAccount.connect(supplierBookingOperator).mintBookingToken(
        distributorTTMAccount.getAddress(), // Reserved for
        tokenURI, // URI
        expirationTimestamp, // Expiration of the reservation
        price, // Price of token in wei
        ethers.ZeroAddress, // paymentToken: zero address, means native coin
        0, // offchain payment currency, zero means unset
        true, // cancellable
    );

    // Token with ID 0 minted with native payment
    const tokenWithNativePayment = 0n;

    // Buy the token
    await distributorTTMAccount
        .connect(distributorBookingOperator)
        .buyBookingToken(tokenWithNativePayment, price, ethers.ZeroAddress);

    // Mint BOOKING TOKEN with NULLUSD PAYMENT------------------------------------------------------------------------

    const tokenURI2 = "data:application/json;base64,eyJuYW1lIjoiQ2FtaW5vIE1lc3NlbmdlciBCb29raW5nVG9rZW4gVGVzdCJ9Cg==";
    const expirationTimestamp2 = Math.floor(Date.now() / 1000) + 120;
    const price2 = ethers.parseEther("99.95");

    await supplierTTMAccount.connect(supplierBookingOperator).mintBookingToken(
        distributorTTMAccount.getAddress(), // Reserved for
        tokenURI2, // URI
        expirationTimestamp2, // Expiration of the reservation
        price2, // Price of token in wei
        nullUSD.getAddress(), // paymentToken
        0, // offchain payment currency, zero means unset
        true, // cancellable
    );

    // Token with ID 1 minted with NullUSD payment
    const tokenWithNullUSDPayment = 1n;

    // Buy the token
    await distributorTTMAccount
        .connect(distributorBookingOperator)
        .buyBookingToken(tokenWithNullUSDPayment, price2, await nullUSD.getAddress());

    // Mint BOOKING TOKEN without buying -----------------------------------------------------------------------------

    const tokenURI3 = "data:application/json;base64,eyJuYW1lIjoiQ2FtaW5vIE1lc3NlbmdlciBCb29raW5nVG9rZW4gVGVzdCJ9Cg==";
    const expirationTimestamp3 = Math.floor(Date.now() / 1000) + 600;
    const price3 = ethers.parseEther("0.95");

    await supplierTTMAccount.connect(supplierBookingOperator).mintBookingToken(
        distributorTTMAccount.getAddress(), // Reserved for
        tokenURI3, // URI
        expirationTimestamp3, // Expiration of the reservation
        price3, // Price of token in wei
        ethers.ZeroAddress, // paymentToken
        0, // offchain payment currency, zero means unset
        true, // cancellable
    );

    // Token with ID 2 minted without buying
    const tokenWithoutBuying = 2n;

    // Mint BOOKING TOKEN with passed expiration ---------------------------------------------------------------------

    const tokenURI4 = "data:application/json;base64,eyJuYW1lIjoiQ2FtaW5vIE1lc3NlbmdlciBCb29raW5nVG9rZW4gVGVzdCJ9Cg==";

    // get block time from the chain
    const block = await ethers.provider.getBlock("latest");

    const expirationTimestamp4 = block.timestamp + 70; // min expiration time diff is 60
    const price4 = ethers.parseEther("0.95");

    await supplierTTMAccount.connect(supplierBookingOperator).mintBookingToken(
        distributorTTMAccount.getAddress(), // Reserved for
        tokenURI4, // URI
        expirationTimestamp4, // Expiration of the reservation
        price4, // Price of token in wei
        ethers.ZeroAddress, // paymentToken
        0, // offchain payment currency, zero means unset
        true, // cancellable
    );

    // Advance time to after the expiration
    await network.provider.send("evm_increaseTime", [expirationTimestamp4 - block.timestamp + 10]);
    await network.provider.send("evm_mine");

    // Token with ID 3 minted with passed expiration
    const tokenWithPassedExpiration = 3n;

    // Mint BOOKING TOKEN with off chain payment ---------------------------------------------------------------------

    const tokenURI5 = "data:application/json;base64,eyJuYW1lIjoiQ2FtaW5vIE1lc3NlbmdlciBCb29raW5nVG9rZW4gVGVzdCJ9Cg==";
    const expirationTimestamp5 = (await ethers.provider.getBlock("latest")).timestamp + 120;
    const price5 = ethers.parseEther("0.95");
    const offChainPaymentToken = ethers.getAddress("0x0000000000000000000000000000000000000001");
    const offChainPaymentCurrency = 123;

    await supplierTTMAccount.connect(supplierBookingOperator).mintBookingToken(
        distributorTTMAccount.getAddress(), // Reserved for
        tokenURI5, // URI
        expirationTimestamp5, // Expiration of the reservation
        price5, // Price of token in wei
        offChainPaymentToken, // paymentToken
        offChainPaymentCurrency, // offchain payment currency
        true, // cancellable
    );

    // Token with ID 4 minted with off chain payment
    const tokenWithOffChainPayment = 4n;

    // Buy the token
    await distributorTTMAccount
        .connect(distributorBookingOperator)
        .buyBookingToken(tokenWithOffChainPayment, price5, offChainPaymentToken);

    /// OTHER TTM ACCOUNT ///

    // We also need another TTM Account to test for fail cases
    // Create other TTMAccount
    const tx = await ttmAccountManager.createTTMAccount(
        signers.ttmAccountAdmin.address,
        signers.ttmAccountUpgrader.address,
        { value: ethers.parseEther("100") },
    );

    const receipt = await tx.wait();

    // Parse event to get the TTMAccount address (this is the UUPS proxy address)
    const event = receipt.logs.find((log) => {
        try {
            return ttmAccountManager.interface.parseLog(log).name === "TTMAccountCreated";
        } catch (e) {
            return false;
        }
    });

    const parsedEvent = ttmAccountManager.interface.parseLog(event);
    const otherTTMAccountAddress = parsedEvent.args.account;

    // Get the TTMAccount instance at the address
    const otherTTMAccount = await ethers.getContractAt("TTMAccount", otherTTMAccountAddress);

    // Deposit funds to the TTMAccount
    const depositAmount = ethers.parseEther("5");
    const depositTx = {
        to: otherTTMAccount.getAddress(),
        value: depositAmount,
    };
    const txResponse = await signers.depositor.sendTransaction(depositTx);
    await txResponse.wait();

    // Distributor
    await otherTTMAccount
        .connect(signers.ttmAccountAdmin)
        .grantRole(BOOKING_OPERATOR_ROLE, otherBookingOperator.address);

    return {
        supplierTTMAccount,
        distributorTTMAccount,
        bookingToken,
        nullUSD,
        tokenWithNativePayment,
        tokenWithNullUSDPayment,
        supplierBookingOperator,
        distributorBookingOperator,
        otherTTMAccount,
        otherBookingOperator,
        tokenWithoutBuying,
        tokenWithPassedExpiration,
        offChainPaymentToken,
        offChainPaymentCurrency,
        tokenWithOffChainPayment,
    };
}

async function deployAndConfigureAllWithRegisteredServicesFixture() {
    // Set up signers
    await setupSigners();

    const { ttmAccountManager, ttmAccount } = await loadFixture(deployAndConfigureAllFixture);

    // Grant SERVICE_REGISTRY_ADMIN_ROLE
    const SERVICE_REGISTRY_ADMIN_ROLE = await ttmAccountManager.SERVICE_REGISTRY_ADMIN_ROLE();
    await ttmAccountManager
        .connect(signers.managerAdmin)
        .grantRole(SERVICE_REGISTRY_ADMIN_ROLE, signers.registryAdmin.address);

    // Services to register
    const serviceName1 = "ttm.service.accommodation.v1.AccommodationSearchService";
    const serviceHash1 = ethers.keccak256(ethers.toUtf8Bytes(serviceName1));

    const serviceName2 = "ttm.service.accommodation.v2.AccommodationSearchService";
    const serviceHash2 = ethers.keccak256(ethers.toUtf8Bytes(serviceName2));

    const serviceName3 = "ttm.service.accommodation.v3.AccommodationSearchService";
    const serviceHash3 = ethers.keccak256(ethers.toUtf8Bytes(serviceName3));

    const serviceName4 = "ttm.service.accommodation.v4.AccommodationSearchService";
    const serviceHash4 = ethers.keccak256(ethers.toUtf8Bytes(serviceName4));

    const serviceName5 = "ttm.service.accommodation.v5.AccommodationSearchService";
    const serviceHash5 = ethers.keccak256(ethers.toUtf8Bytes(serviceName5));

    const serviceName6 = "ttm.service.accommodation.v6.AccommodationSearchService";
    const serviceHash6 = ethers.keccak256(ethers.toUtf8Bytes(serviceName6));

    const services = {
        serviceName1,
        serviceHash1,
        serviceName2,
        serviceHash2,
        serviceName3,
        serviceHash3,
        serviceName4,
        serviceHash4,
        serviceName5,
        serviceHash5,
        serviceName6,
        serviceHash6,
    };

    // Register services
    await ttmAccountManager.connect(signers.registryAdmin).registerService(serviceName1);
    await ttmAccountManager.connect(signers.registryAdmin).registerService(serviceName2);
    await ttmAccountManager.connect(signers.registryAdmin).registerService(serviceName3);
    await ttmAccountManager.connect(signers.registryAdmin).registerService(serviceName4);
    await ttmAccountManager.connect(signers.registryAdmin).registerService(serviceName5);
    await ttmAccountManager.connect(signers.registryAdmin).registerService(serviceName6);

    // Get the SERVICE_ADMIN_ROLE
    const SERVICE_ADMIN_ROLE = await ttmAccount.SERVICE_ADMIN_ROLE();

    // Grant SERVICE_ADMIN_ROLE to otherAccount1
    await ttmAccount.connect(signers.ttmAccountAdmin).grantRole(SERVICE_ADMIN_ROLE, signers.ttmServiceAdmin.address);

    return { ttmAccountManager, ttmAccount, services };
}

module.exports = {
    setupSigners,
    serviceHash,
    deployTTMAccountManagerFixture,
    deployTTMAccountImplFixture,
    deployTTMAccountManagerWithTTMAccountImplFixture,
    deployAndConfigureAllFixture,
    deployTTMAccountWithDepositFixture,
    deployBookingTokenFixture,
    deployAndConfigureAllWithRegisteredServicesFixture,
    deployBookingTokenWithNullUSDFixture,
    deployCancellationSupportFixture,
    deployNullUSDFixture,
};
