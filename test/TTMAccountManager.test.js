/**
 * @dev TTMAccountManager tests
 */
const { loadFixture } = require("@nomicfoundation/hardhat-toolbox/network-helpers");
const { expect } = require("chai");
const { ethers, upgrades } = require("hardhat");

const { setCode } = require("@nomicfoundation/hardhat-network-helpers");

// Fixtures
const {
    setupSigners,
    deployNullUSDFixture,
    deployTTMAccountManagerFixture,
    deployTTMAccountImplFixture,
    deployTTMAccountManagerWithTTMAccountImplFixture,
    deployAndConfigureAllFixture,
} = require("./utils/fixtures");

describe("TTMAccountManager", function () {
    describe("Main", function () {
        it("should deploy correctly with the right state", async function () {
            const { ttmAccountManager } = await loadFixture(deployTTMAccountManagerFixture);

            const DEFAULT_ADMIN_ROLE = await ttmAccountManager.DEFAULT_ADMIN_ROLE();
            const PAUSER_ROLE = await ttmAccountManager.PAUSER_ROLE();
            const UPGRADER_ROLE = await ttmAccountManager.UPGRADER_ROLE();
            const VERSIONER_ROLE = await ttmAccountManager.VERSIONER_ROLE();

            // Check roles
            expect(await ttmAccountManager.hasRole(DEFAULT_ADMIN_ROLE, signers.managerAdmin.address)).to.be.true;
            expect(await ttmAccountManager.hasRole(PAUSER_ROLE, signers.managerPauser.address)).to.be.true;
            expect(await ttmAccountManager.hasRole(UPGRADER_ROLE, signers.managerUpgrader.address)).to.be.true;
            expect(await ttmAccountManager.hasRole(VERSIONER_ROLE, signers.managerVersioner.address)).to.be.true;

            // Check state
            expect(await ttmAccountManager.paused()).to.be.false;
        });

        it("should get role counts correctly", async function () {
            // Set up signers
            await setupSigners();

            const { ttmAccountManager } = await loadFixture(deployTTMAccountManagerFixture);

            const DEFAULT_ADMIN_ROLE = await ttmAccountManager.DEFAULT_ADMIN_ROLE();
            const PAUSER_ROLE = await ttmAccountManager.PAUSER_ROLE();
            const UPGRADER_ROLE = await ttmAccountManager.UPGRADER_ROLE();
            const VERSIONER_ROLE = await ttmAccountManager.VERSIONER_ROLE();

            expect(await ttmAccountManager.getRoleMemberCount(DEFAULT_ADMIN_ROLE)).to.be.equal(1);
            expect(await ttmAccountManager.getRoleMemberCount(PAUSER_ROLE)).to.be.equal(1);
            expect(await ttmAccountManager.getRoleMemberCount(UPGRADER_ROLE)).to.be.equal(1);
            expect(await ttmAccountManager.getRoleMemberCount(VERSIONER_ROLE)).to.be.equal(1);
        });

        it("should set and get booking token addr correctly", async function () {
            // Set up signers
            await setupSigners();

            const { ttmAccountManager } = await loadFixture(deployAndConfigureAllFixture);

            // Get booking token address
            expect(await ttmAccountManager.getBookingTokenAddress()).to.be.not.reverted;

            // Try to set booking token addr with non-auth address
            await expect(
                ttmAccountManager.connect(signers.otherAccount3).setBookingTokenAddress(ethers.ZeroAddress),
            ).to.be.revertedWithCustomError(ttmAccountManager, "AccessControlUnauthorizedAccount");

            // Try to set booking token to invalid addresses
            await expect(ttmAccountManager.connect(signers.managerVersioner).setBookingTokenAddress(ethers.ZeroAddress))
                .to.be.revertedWithCustomError(ttmAccountManager, "InvalidBookingTokenAddress")
                .withArgs(ethers.ZeroAddress);

            await expect(
                ttmAccountManager
                    .connect(signers.managerVersioner)
                    .setBookingTokenAddress(signers.otherAccount1.address),
            )
                .to.be.revertedWithCustomError(ttmAccountManager, "InvalidBookingTokenAddress")
                .withArgs(signers.otherAccount1.address);
        });
    });

    describe("Upgrades", function () {
        it("should upgrade correctly", async function () {
            // Set up signers
            await setupSigners();

            const { ttmAccountManager } = await loadFixture(deployTTMAccountManagerFixture);

            const TTMAccountManagerTest = await ethers.getContractFactory(
                "TTMAccountManagerTest",
                signers.managerUpgrader,
            );
            const ttmAccountManagerTest = await upgrades.upgradeProxy(ttmAccountManager, TTMAccountManagerTest);

            await expect(await ttmAccountManagerTest.getVersion()).to.be.equal("TESTING");
        });

        it("should not upgrade if the caller does not have the upgrader role", async function () {
            // Set up signers
            await setupSigners();

            const { ttmAccountManager } = await loadFixture(deployTTMAccountManagerFixture);

            const TTMAccountManagerTest = await ethers.getContractFactory("TTMAccountManagerTest", signers.managerPauser);

            await expect(upgrades.upgradeProxy(ttmAccountManager, TTMAccountManagerTest)).to.be.revertedWithCustomError(
                TTMAccountManagerTest,
                "AccessControlUnauthorizedAccount",
            );
        });
    });

    describe("TTMAccount Implementation", function () {
        it("should set TTMAccount implementation correctly", async function () {
            // Set up signers
            await setupSigners();

            const { ttmAccountManager } = await loadFixture(deployTTMAccountManagerFixture);
            const { ttmAccountImpl } = await loadFixture(deployTTMAccountImplFixture);

            const ttmAccountImplAddress = await ttmAccountImpl.getAddress();

            await expect(
                await ttmAccountManager.connect(signers.managerVersioner).setAccountImplementation(ttmAccountImplAddress),
            )
                .to.emit(ttmAccountManager, "TTMAccountImplementationUpdated")
                .withArgs(ethers.ZeroAddress, ttmAccountImplAddress);
        });

        it("should get TTMAccount implementation correctly", async function () {
            // Set up signers
            await setupSigners();

            const { ttmAccountManager, ttmAccountImplAddress } = await loadFixture(
                deployTTMAccountManagerWithTTMAccountImplFixture,
            );

            await expect(await ttmAccountManager.getAccountImplementation()).to.be.equal(ttmAccountImplAddress);
        });

        it("should revert if the implementation is zero code length address", async function () {
            // Set up signers
            await setupSigners();

            const { ttmAccountManager, ttmAccountImplAddress } = await loadFixture(
                deployTTMAccountManagerWithTTMAccountImplFixture,
            );

            await expect(
                ttmAccountManager.connect(signers.managerVersioner).setAccountImplementation(ethers.ZeroAddress),
            )
                .to.be.revertedWithCustomError(ttmAccountManager, "TTMAccountInvalidImplementation")
                .withArgs(ethers.ZeroAddress);
        });

        it("should revert if the caller does not have the versioner role", async function () {
            // Set up signers
            await setupSigners();

            const { ttmAccountManager, ttmAccountImplAddress } = await loadFixture(
                deployTTMAccountManagerWithTTMAccountImplFixture,
            );

            await expect(
                ttmAccountManager.connect(signers.otherAccount1).setAccountImplementation(ttmAccountImplAddress),
            ).to.be.revertedWithCustomError(ttmAccountManager, "AccessControlUnauthorizedAccount");
        });
    });

    describe("Pausable", function () {
        it("should pause and unpause the contract", async function () {
            // Set up signers
            await setupSigners();

            const { ttmAccountManager } = await loadFixture(deployTTMAccountManagerFixture);

            await ttmAccountManager.connect(signers.managerPauser).pause();
            await expect(await ttmAccountManager.paused()).to.be.true;

            await ttmAccountManager.connect(signers.managerPauser).unpause();
            await expect(await ttmAccountManager.paused()).to.be.false;
        });

        it("should not allow non-pauser to pause", async function () {
            // Set up signers
            await setupSigners();

            const { ttmAccountManager } = await loadFixture(deployTTMAccountManagerFixture);

            await expect(ttmAccountManager.connect(signers.otherAccount1).pause()).to.be.revertedWithCustomError(
                ttmAccountManager,
                "AccessControlUnauthorizedAccount",
            );

            await expect(ttmAccountManager.connect(signers.otherAccount1).unpause()).to.be.revertedWithCustomError(
                ttmAccountManager,
                "AccessControlUnauthorizedAccount",
            );
        });
    });

    describe("TTMAccount", function () {
        it("should create TTMAccount correctly", async function () {
            // Set up signers
            await setupSigners();

            const { ttmAccountManager, ttmAccount } = await loadFixture(deployAndConfigureAllFixture);

            const ttmAccountManagerAddress = await ttmAccountManager.getAddress();
            const ttmAccountAddress = await ttmAccount.getAddress();

            expect(await ttmAccountManager.isTTMAccount(ttmAccountAddress)).to.be.true;
            expect(await ttmAccountManager.isTTMAccount(signers.otherAccount1.address)).to.be.false;
            expect(await ttmAccountManager.isTTMAccount(ethers.ZeroAddress)).to.be.false;
            expect(await ttmAccount.getManagerAddress()).to.be.equal(ttmAccountManagerAddress);

            // Check balance for native deposit
            expect(await ethers.provider.getBalance(ttmAccountAddress)).to.be.equal(ethers.parseEther("100"));
        });

        it("should revert creating for zero code btoken and impl addr", async function () {
            // Set up signers
            await setupSigners();

            const { ttmAccountManager, bookingToken } = await loadFixture(deployAndConfigureAllFixture);

            // Get account impl
            const TTMAccountImplAddr = await ttmAccountManager.getAccountImplementation();

            // Set booking token code to zero
            await network.provider.send("hardhat_setCode", [await bookingToken.getAddress(), "0x"]);

            // Create TTMAccount
            await expect(
                ttmAccountManager.createTTMAccount(signers.ttmAccountAdmin.address, signers.ttmAccountUpgrader.address, {
                    value: ethers.parseEther("100"),
                }),
            ).to.be.revertedWithCustomError(ttmAccountManager, "InvalidBookingTokenAddress");

            // Set acct impl code to zero
            await network.provider.send("hardhat_setCode", [TTMAccountImplAddr, "0x"]);

            // Create TTMAccount
            await expect(
                ttmAccountManager.createTTMAccount(signers.ttmAccountAdmin.address, signers.ttmAccountUpgrader.address, {
                    value: ethers.parseEther("100"),
                }),
            ).to.be.revertedWithCustomError(ttmAccountManager, "TTMAccountInvalidImplementation");
        });

        it("should fail if admin is zero address", async function () {
            // Set up signers
            await setupSigners();

            const { ttmAccountManager } = await loadFixture(deployAndConfigureAllFixture);

            await expect(
                ttmAccountManager.createTTMAccount(ethers.ZeroAddress, signers.ttmAccountUpgrader, {
                    value: ethers.parseEther("100"),
                }),
            ).to.be.revertedWithCustomError(ttmAccountManager, "TTMAccountInvalidAdmin");
        });

        it("should fail if the manager is paused", async function () {
            // Set up signers
            await setupSigners();

            const { ttmAccountManager } = await loadFixture(deployAndConfigureAllFixture);

            await ttmAccountManager.connect(signers.managerPauser).pause();
            await expect(
                ttmAccountManager.createTTMAccount(signers.ttmAccountAdmin.address, signers.ttmAccountUpgrader, {
                    value: ethers.parseEther("100"),
                }),
            ).to.be.revertedWithCustomError(ttmAccountManager, "EnforcedPause");
        });

        it("should not fail for any msg.value", async function () {
            // Set up signers
            await setupSigners();

            const { ttmAccountManager } = await loadFixture(deployAndConfigureAllFixture);

            // Zero msg.value
            await expect(
                ttmAccountManager.createTTMAccount(signers.ttmAccountAdmin.address, signers.ttmAccountUpgrader, {
                    value: 0n,
                }),
            ).to.be.not.reverted;

            const nonZeroMsgValue = ethers.parseEther("300");

            // Non-zero msg.value
            const nonZeroMsgValueTx = await ttmAccountManager.createTTMAccount(
                signers.ttmAccountAdmin.address,
                signers.ttmAccountUpgrader,
                {
                    value: nonZeroMsgValue,
                },
            );

            await expect(nonZeroMsgValueTx).to.be.not.reverted;

            const receipt = await nonZeroMsgValueTx.wait();

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

            // Check balances
            await expect(nonZeroMsgValueTx).to.changeEtherBalances(
                [signers.managerAdmin, ttmAccountAddress],
                [-nonZeroMsgValue, nonZeroMsgValue],
            );
        });

        it("should set and get correct account creator", async function () {
            // Set up signers
            await setupSigners();

            const { ttmAccountManager } = await loadFixture(deployAndConfigureAllFixture);

            // Create distributor TTMAccount
            // This is called with managerAdmin as the signer
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
            const newTTMAccountAddress = parsedEvent.args.account;

            expect(await ttmAccountManager.getTTMAccountCreator(newTTMAccountAddress)).to.be.equal(
                signers.managerAdmin.address,
            );
        });
    });
});
