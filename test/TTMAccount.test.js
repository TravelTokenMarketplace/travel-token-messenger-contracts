/**
 * @dev TTMAccount tests
 */
const { loadFixture } = require("@nomicfoundation/hardhat-toolbox/network-helpers");
const { expect } = require("chai");
const { ethers, upgrades } = require("hardhat");

const {
    setupSigners,
    deployTTMAccountManagerFixture,
    deployTTMAccountImplFixture,
    deployTTMAccountManagerWithTTMAccountImplFixture,
    deployAndConfigureAllFixture,
    deployTTMAccountWithDepositFixture,
    deployBookingTokenWithNullUSDFixture,
} = require("./utils/fixtures");

describe("TTMAccount", function () {
    describe("Upgrade", function () {
        it("should upgrade to new implementation address", async function () {
            const { ttmAccountManager, ttmAccount } = await loadFixture(deployAndConfigureAllFixture);

            // Old implementation
            const oldImplementationAddress = await ttmAccountManager.getAccountImplementation();

            // Create a new implementation for TTMAccount
            const BookingTokenOperator = await ethers.getContractFactory("BookingTokenOperator");
            const bookingTokenOperator = await BookingTokenOperator.deploy();
            const TTMAccountImplV2 = await ethers.getContractFactory("TTMAccount", {
                libraries: { BookingTokenOperator: await bookingTokenOperator.getAddress() },
            });
            const ttmAccountImplV2 = await TTMAccountImplV2.deploy();
            await ttmAccountImplV2.waitForDeployment();
            const newImplementationAddress = await ttmAccountImplV2.getAddress();

            // Set new implementation on the manager
            await ttmAccountManager
                .connect(signers.managerVersioner)
                .setAccountImplementation(newImplementationAddress);
            await expect(await ttmAccountManager.getAccountImplementation()).to.be.equal(newImplementationAddress);

            // Upgrade the account
            await expect(
                ttmAccount.connect(signers.ttmAccountUpgrader).upgradeToAndCall(newImplementationAddress, "0x"),
            )
                .to.emit(ttmAccount, "TTMAccountUpgraded")
                .withArgs(oldImplementationAddress, newImplementationAddress);
        });

        it("should revert upgrade if implementation address does not match", async function () {
            const { ttmAccountManager, ttmAccount } = await loadFixture(deployAndConfigureAllFixture);

            // Old implementation
            const oldImplementationAddress = await ttmAccountManager.getAccountImplementation();

            // Create a new implementation for TTMAccount
            const BookingTokenOperator = await ethers.getContractFactory("BookingTokenOperator");
            const bookingTokenOperator = await BookingTokenOperator.deploy();
            const TTMAccountImplV2 = await ethers.getContractFactory("TTMAccount", {
                libraries: { BookingTokenOperator: await bookingTokenOperator.getAddress() },
            });
            const ttmAccountImplV2 = await TTMAccountImplV2.deploy();
            await ttmAccountImplV2.waitForDeployment();
            const newImplementationAddress = await ttmAccountImplV2.getAddress();

            // SKIP: DO NOT set new implementation on the manager here

            // Try to upgrade the account
            await expect(
                ttmAccount.connect(signers.ttmAccountUpgrader).upgradeToAndCall(newImplementationAddress, "0x"),
            )
                .to.be.revertedWithCustomError(ttmAccount, "TTMAccountImplementationMismatch")
                .withArgs(oldImplementationAddress, newImplementationAddress);
        });

        it("should revert upgrade if address is not uups upgradeable", async function () {
            const { ttmAccountManager, ttmAccount } = await loadFixture(deployAndConfigureAllFixture);

            // Old implementation
            const oldImplementationAddress = await ttmAccountManager.getAccountImplementation();

            // Create a new implementation for TTMAccount
            const dummyAccountImpl = await ethers.getContractFactory("Dummy");
            const dummyAccountImplV2 = await dummyAccountImpl.deploy();
            await dummyAccountImplV2.waitForDeployment();
            const newImplementationAddress = await dummyAccountImplV2.getAddress();

            // Set new implementation on the manager
            await ttmAccountManager
                .connect(signers.managerVersioner)
                .setAccountImplementation(newImplementationAddress);
            await expect(await ttmAccountManager.getAccountImplementation()).to.be.equal(newImplementationAddress);

            // Upgrade the account
            await expect(
                ttmAccount.connect(signers.ttmAccountUpgrader).upgradeToAndCall(newImplementationAddress, "0x"),
            )
                .to.be.revertedWithCustomError(ttmAccount, "ERC1967InvalidImplementation")
                .withArgs(newImplementationAddress);
        });

        it("should revert upgrade if address is same with the current one", async function () {
            const { ttmAccountManager, ttmAccount } = await loadFixture(deployAndConfigureAllFixture);

            // Old implementation
            const oldImplementationAddress = await ttmAccountManager.getAccountImplementation();

            // Upgrade the account
            await expect(
                ttmAccount.connect(signers.ttmAccountUpgrader).upgradeToAndCall(oldImplementationAddress, "0x"),
            )
                .to.be.revertedWithCustomError(ttmAccount, "TTMAccountNoUpgradeNeeded")
                .withArgs(oldImplementationAddress, oldImplementationAddress);
        });

        it("should revert upgrade if caller is not authorized", async function () {
            const { ttmAccountManager, ttmAccount } = await loadFixture(deployAndConfigureAllFixture);

            // Old implementation
            const oldImplementationAddress = await ttmAccountManager.getAccountImplementation();

            const UPGRADER_ROLE = await ttmAccount.UPGRADER_ROLE();
            const unauthorizedCaller = signers.otherAccount1;

            // Try to upgrade with unauthorized caller
            await expect(ttmAccount.connect(unauthorizedCaller).upgradeToAndCall(oldImplementationAddress, "0x"))
                .to.be.revertedWithCustomError(ttmAccount, "AccessControlUnauthorizedAccount")
                .withArgs(unauthorizedCaller.address, UPGRADER_ROLE);
        });
    });

    describe("Registering Bots", function () {
        it("should register bots correctly", async function () {
            const { ttmAccountManager, ttmAccount } = await loadFixture(deployAndConfigureAllFixture);

            const MESSENGER_BOT_ROLE = await ttmAccount.MESSENGER_BOT_ROLE();
            const botAddr = signers.botOperator.address;

            // Grant MESSENGER_BOT_ROLE
            await expect(ttmAccount.connect(signers.ttmAccountAdmin).grantRole(MESSENGER_BOT_ROLE, botAddr))
                .to.emit(ttmAccount, "RoleGranted")
                .withArgs(MESSENGER_BOT_ROLE, botAddr, signers.ttmAccountAdmin.address);

            await expect(await ttmAccount.isBotAllowed(botAddr)).to.be.true;
        });
    });

    describe("Deposit", function () {
        it("should allow anyone to send funds", async function () {
            const { ttmAccountManager, ttmAccount } = await loadFixture(deployAndConfigureAllFixture);

            const anyone = signers.otherAccount1;

            const anyoneInitialBalance = await ethers.provider.getBalance(anyone.address);
            const ttmAccountInitialBalance = await ethers.provider.getBalance(ttmAccount.getAddress());

            const depositAmount = ethers.parseEther("1");

            // Sender
            const depositTx = {
                to: ttmAccount.getAddress(),
                value: depositAmount,
            };

            await expect(await anyone.sendTransaction(depositTx)).to.not.be.reverted;

            // Check balances
            // Sender balance should be lower than the difference between their initial balance and the deposit
            expect(await ethers.provider.getBalance(anyone.address)).to.be.lt(anyoneInitialBalance - depositAmount);

            // TTMAccount balance should be equal to the sum of the initial balance and the deposit
            expect(await ethers.provider.getBalance(ttmAccount.getAddress())).to.be.equal(
                ttmAccountInitialBalance + depositAmount,
            );
        });
    });

    describe("Withdraw", function () {
        it("should allow withdrawer role to withdraw", async function () {
            const { ttmAccount } = await loadFixture(deployTTMAccountWithDepositFixture);

            const withdrawer = signers.withdrawer;
            const withdrawAmount = ethers.parseEther("0.5");

            // Withdraw
            const withdrawTx = ttmAccount.connect(withdrawer).withdraw(withdrawer.address, withdrawAmount);
            await expect(withdrawTx).to.changeEtherBalances(
                [ttmAccount, withdrawer],
                [-withdrawAmount, withdrawAmount],
            );
            await expect(withdrawTx).to.emit(ttmAccount, "Withdraw").withArgs(withdrawer.address, withdrawAmount);
        });

        it("should revert if not withdrawer role", async function () {
            const { ttmAccount } = await loadFixture(deployTTMAccountWithDepositFixture);

            const withdrawer = signers.otherAccount1;
            const withdrawAmount = ethers.parseEther("0.5");

            const WITHDRAWER_ROLE = await ttmAccount.WITHDRAWER_ROLE();

            // Withdraw
            const withdrawTx = ttmAccount.connect(withdrawer).withdraw(withdrawer.address, withdrawAmount);
            await expect(withdrawTx)
                .to.be.revertedWithCustomError(ttmAccount, "AccessControlUnauthorizedAccount")
                .withArgs(withdrawer.address, WITHDRAWER_ROLE);
        });

        it("should withdraw all amount", async function () {
            const { ttmAccount } = await loadFixture(deployTTMAccountWithDepositFixture);

            const withdrawer = signers.withdrawer;
            // Withdraw 1 ETH from initial deposit of 1 ETH
            const withdrawAmount = ethers.parseEther("1");

            // Try withdraw
            const withdrawTx = ttmAccount.connect(withdrawer).withdraw(withdrawer.address, withdrawAmount);
            await expect(withdrawTx).to.be.not.reverted;

            // Check balances
            await expect(withdrawTx).to.changeEtherBalances(
                [ttmAccount, withdrawer],
                [-withdrawAmount, withdrawAmount],
            );

            // Full balance is withdrawable

            // Withdraw all remaining 100 ETH
            const withdrawAmount2 = ethers.parseEther("100");

            // Try withdraw
            const withdrawTx2 = ttmAccount.connect(withdrawer).withdraw(withdrawer.address, withdrawAmount2);
            await expect(withdrawTx2).to.be.emit(ttmAccount, "Withdraw").withArgs(withdrawer.address, withdrawAmount2);

            // Check balances
            await expect(withdrawTx2).to.changeEtherBalances(
                [ttmAccount, withdrawer],
                [-withdrawAmount2, withdrawAmount2],
            );
        });

        it("should revert if withdraw recipient is zero address", async function () {
            const { ttmAccount } = await loadFixture(deployTTMAccountWithDepositFixture);

            const withdrawer = signers.withdrawer;
            const withdrawAmount = ethers.parseEther("0.5");

            // Withdraw
            const withdrawTx = ttmAccount.connect(withdrawer).withdraw(ethers.ZeroAddress, withdrawAmount);
            await expect(withdrawTx).to.be.revertedWithCustomError(ttmAccount, "TransferToZeroAddress");
        });
    });

    describe("Enumerable", function () {
        it("should get role counts correctly", async function () {
            const { ttmAccountManager, ttmAccount } = await loadFixture(deployAndConfigureAllFixture);

            const DEFAULT_ADMIN_ROLE = await ttmAccount.DEFAULT_ADMIN_ROLE();
            const UPGRADER_ROLE = await ttmAccount.UPGRADER_ROLE();
            const BOOKING_OPERATOR_ROLE = await ttmAccount.BOOKING_OPERATOR_ROLE();

            expect(await ttmAccount.getRoleMemberCount(DEFAULT_ADMIN_ROLE)).to.be.equal(1);
            expect(await ttmAccount.getRoleMemberCount(UPGRADER_ROLE)).to.be.equal(1);

            // Booking operator role is not granted by default
            expect(await ttmAccount.getRoleMemberCount(BOOKING_OPERATOR_ROLE)).to.be.equal(0);

            // Grant booking operator role
            await expect(
                ttmAccount
                    .connect(signers.ttmAccountAdmin)
                    .grantRole(await ttmAccount.BOOKING_OPERATOR_ROLE(), signers.otherAccount1.address),
            ).to.not.reverted;

            // Booking operator role is granted, count should be 1
            expect(await ttmAccount.getRoleMemberCount(BOOKING_OPERATOR_ROLE)).to.be.equal(1);
        });
    });

    describe("Messenger Bot", function () {
        it("should add messenger bot correctly", async function () {
            const { ttmAccount } = await loadFixture(deployAndConfigureAllFixture);

            const bot = signers.otherAccount1;

            // Register bot
            await expect(ttmAccount.connect(signers.ttmAccountAdmin).addMessengerBot(bot.address, 0n))
                .to.emit(ttmAccount, "MessengerBotAdded")
                .withArgs(bot.address);

            // Check if bot is allowed
            expect(await ttmAccount.isBotAllowed(bot.address)).to.be.true;

            // Check roles
            expect(await ttmAccount.hasRole(await ttmAccount.MESSENGER_BOT_ROLE(), bot.address)).to.be.true;
            expect(await ttmAccount.hasRole(await ttmAccount.BOOKING_OPERATOR_ROLE(), bot.address)).to.be.true;
            expect(await ttmAccount.hasRole(await ttmAccount.GAS_WITHDRAWER_ROLE(), bot.address)).to.be.true;
        });

        it("should remove messenger bot correctly", async function () {
            const { ttmAccount } = await loadFixture(deployAndConfigureAllFixture);

            const bot = signers.otherAccount1;

            // Register bot
            await expect(ttmAccount.connect(signers.ttmAccountAdmin).addMessengerBot(bot.address, 0n))
                .to.emit(ttmAccount, "MessengerBotAdded")
                .withArgs(bot.address);

            // Remove bot
            await expect(ttmAccount.connect(signers.ttmAccountAdmin).removeMessengerBot(bot.address))
                .to.emit(ttmAccount, "MessengerBotRemoved")
                .withArgs(bot.address);
        });

        it("should add messenger bot with gas money withdrawal correctly", async function () {
            const { ttmAccount } = await loadFixture(deployTTMAccountWithDepositFixture);

            const bot = signers.otherAccount1;

            const withdrawAmount = ethers.parseEther("25"); // Withdraw 25 ETH to the bot.

            // Register bot
            const withdrawTx = ttmAccount
                .connect(signers.ttmAccountAdmin)
                ["addMessengerBot(address,uint256)"](bot.address, withdrawAmount);

            await expect(withdrawTx).to.changeEtherBalances([ttmAccount, bot], [-withdrawAmount, withdrawAmount]);
            await expect(withdrawTx).to.emit(ttmAccount, "MessengerBotAdded").withArgs(bot.address);
        });

        it("should revert addMessengerBot if caller is not authorized", async function () {
            const { ttmAccount } = await loadFixture(deployAndConfigureAllFixture);

            const BOT_ADMIN_ROLE = await ttmAccount.BOT_ADMIN_ROLE();
            const unauthorizedCaller = signers.otherAccount1;
            const bot = signers.otherAccount2;

            // Try to add messenger bot with unauthorized caller
            await expect(ttmAccount.connect(unauthorizedCaller).addMessengerBot(bot.address, 0n))
                .to.be.revertedWithCustomError(ttmAccount, "AccessControlUnauthorizedAccount")
                .withArgs(unauthorizedCaller.address, BOT_ADMIN_ROLE);
        });

        it("should revert removeMessengerBot if caller is not authorized", async function () {
            const { ttmAccount } = await loadFixture(deployAndConfigureAllFixture);

            const BOT_ADMIN_ROLE = await ttmAccount.BOT_ADMIN_ROLE();
            const unauthorizedCaller = signers.otherAccount1;
            const bot = signers.otherAccount2;

            // Try to remove messenger bot with unauthorized caller
            await expect(ttmAccount.connect(unauthorizedCaller).removeMessengerBot(bot.address))
                .to.be.revertedWithCustomError(ttmAccount, "AccessControlUnauthorizedAccount")
                .withArgs(unauthorizedCaller.address, BOT_ADMIN_ROLE);
        });

        it("should revert addMessengerBot if bot address is zero", async function () {
            const { ttmAccount } = await loadFixture(deployAndConfigureAllFixture);

            // Try to add messenger bot with zero address
            await expect(
                ttmAccount.connect(signers.ttmAccountAdmin).addMessengerBot(ethers.ZeroAddress, 0n),
            ).to.be.revertedWithCustomError(ttmAccount, "TransferToZeroAddress");
        });
    });

    describe("Transfer ERC20 & ERC721", function () {
        it("should transfer ERC20 correctly", async function () {
            const { ttmAccountManager, supplierTTMAccount, distributorTTMAccount, bookingToken, nullUSD } =
                await loadFixture(deployBookingTokenWithNullUSDFixture);

            // Supplier and distributor TTM accounts has 10k NullUSD from the fixture
            const amount = ethers.parseEther("100");

            // Try to send to zero address
            await expect(
                supplierTTMAccount
                    .connect(signers.withdrawer)
                    .transferERC20(await nullUSD.getAddress(), ethers.ZeroAddress, amount),
            ).to.be.revertedWithCustomError(supplierTTMAccount, "TransferToZeroAddress");

            // Try to send with non-auth address
            await expect(
                supplierTTMAccount
                    .connect(signers.otherAccount1)
                    .transferERC20(await nullUSD.getAddress(), signers.otherAccount2.address, amount),
            ).to.be.revertedWithCustomError(supplierTTMAccount, "AccessControlUnauthorizedAccount");

            // Transfer
            await expect(
                await supplierTTMAccount
                    .connect(signers.withdrawer)
                    .transferERC20(await nullUSD.getAddress(), signers.otherAccount1.address, amount),
            ).to.changeTokenBalances(
                nullUSD,
                [await supplierTTMAccount.getAddress(), signers.otherAccount1],
                [-amount, amount],
            );

            // Check balance
            expect(await nullUSD.balanceOf(signers.otherAccount1.address)).to.be.equal(amount);

            // Get remaining balance of NullUSD
            const supplierNullUSDBalance = await nullUSD.balanceOf(await supplierTTMAccount.getAddress());

            // Try to transfer all remaining balance
            await expect(
                supplierTTMAccount
                    .connect(signers.withdrawer)
                    .transferERC20(await nullUSD.getAddress(), signers.otherAccount1.address, supplierNullUSDBalance),
            ).to.changeTokenBalances(
                nullUSD,
                [await supplierTTMAccount.getAddress(), signers.otherAccount1],
                [-supplierNullUSDBalance, supplierNullUSDBalance],
            );

            // Check balances
            expect(await nullUSD.balanceOf(signers.otherAccount1.address)).to.be.equal(supplierNullUSDBalance + amount);
            expect(await nullUSD.balanceOf(await supplierTTMAccount.getAddress())).to.be.equal(0n);
        });

        it("should transfer ERC721 correctly after it expires", async function () {
            const { ttmAccountManager, supplierTTMAccount, distributorTTMAccount, bookingToken, nullUSD } =
                await loadFixture(deployBookingTokenWithNullUSDFixture);

            const tokenURI =
                "data:application/json;base64,eyJuYW1lIjoiQ2FtaW5vIE1lc3NlbmdlciBCb29raW5nVG9rZW4gVGVzdCJ9Cg==";

            const expirationTimestamp = Math.floor(Date.now() / 1000) + 120;

            const price = ethers.parseEther("0.05");

            // Grant BOOKING_OPERATOR_ROLE
            const BOOKING_OPERATOR_ROLE = await supplierTTMAccount.BOOKING_OPERATOR_ROLE();
            await expect(
                supplierTTMAccount
                    .connect(signers.ttmAccountAdmin)
                    .grantRole(BOOKING_OPERATOR_ROLE, signers.btAdmin.address),
            ).to.not.reverted;

            await expect(
                await supplierTTMAccount.connect(signers.btAdmin).mintBookingToken(
                    distributorTTMAccount.getAddress(), // set reservedFor address to distributor TTMAccount
                    tokenURI, // tokenURI
                    expirationTimestamp, // expiration
                    price, // price
                    ethers.ZeroAddress, // zero address
                    0, // off chain payment currency
                    true,
                ),
            )
                .to.be.emit(bookingToken, "TokenReserved")
                .withArgs(
                    0n,
                    distributorTTMAccount.getAddress(),
                    supplierTTMAccount.getAddress(),
                    expirationTimestamp,
                    price,
                    ethers.ZeroAddress, // zero address
                    0n, // off chain payment currency
                    true, // cancellable
                );

            // Advance time to after the expiration
            await network.provider.send("evm_setNextBlockTimestamp", [expirationTimestamp + 1]);
            await network.provider.send("evm_mine");

            // Try to transfer the token with the supplier TTMAccount
            await expect(
                supplierTTMAccount
                    .connect(signers.withdrawer)
                    .transferERC721(await bookingToken.getAddress(), signers.otherAccount1.address, 0n),
            )
                .to.emit(bookingToken, "Transfer")
                .withArgs(await supplierTTMAccount.getAddress(), signers.otherAccount1.address, 0n);

            // Check token ownership
            expect(await bookingToken.ownerOf(0n)).to.equal(await signers.otherAccount1.getAddress());

            // Try to transfer the token to zero address
            await expect(
                supplierTTMAccount
                    .connect(signers.withdrawer)
                    .transferERC721(await bookingToken.getAddress(), ethers.ZeroAddress, 0n),
            ).to.be.revertedWithCustomError(supplierTTMAccount, "TransferToZeroAddress");

            // Try to transfer the token with non-auth address
            await expect(
                supplierTTMAccount
                    .connect(signers.otherAccount1)
                    .transferERC721(await bookingToken.getAddress(), signers.otherAccount2.address, 0n),
            ).to.be.revertedWithCustomError(supplierTTMAccount, "AccessControlUnauthorizedAccount");
        });
    });

    describe("Cancellation Functions", function () {
        it("should revert initiateCancellation if caller is not authorized", async function () {
            const { ttmAccount } = await loadFixture(deployAndConfigureAllFixture);

            const BOOKING_OPERATOR_ROLE = await ttmAccount.BOOKING_OPERATOR_ROLE();
            const unauthorizedCaller = signers.otherAccount1;

            // Try to initiate cancellation with unauthorized caller
            await expect(
                ttmAccount.connect(unauthorizedCaller).initiateCancellation(0n, ethers.parseEther("0.05"), 1, 1),
            )
                .to.be.revertedWithCustomError(ttmAccount, "AccessControlUnauthorizedAccount")
                .withArgs(unauthorizedCaller.address, BOOKING_OPERATOR_ROLE);
        });
        it("should revert acceptCancellation if caller is not authorized", async function () {
            const { ttmAccount } = await loadFixture(deployAndConfigureAllFixture);

            const BOOKING_OPERATOR_ROLE = await ttmAccount.BOOKING_OPERATOR_ROLE();
            const unauthorizedCaller = signers.otherAccount1;

            // Try to accept cancellation with unauthorized caller
            await expect(ttmAccount.connect(unauthorizedCaller).acceptCancellation(0n, ethers.parseEther("0.05")))
                .to.be.revertedWithCustomError(ttmAccount, "AccessControlUnauthorizedAccount")
                .withArgs(unauthorizedCaller.address, BOOKING_OPERATOR_ROLE);
        });

        it("should revert rejectCancellation if caller is not authorized", async function () {
            const { ttmAccount } = await loadFixture(deployAndConfigureAllFixture);

            const BOOKING_OPERATOR_ROLE = await ttmAccount.BOOKING_OPERATOR_ROLE();
            const unauthorizedCaller = signers.otherAccount1;

            // Try to reject cancellation with unauthorized caller
            await expect(ttmAccount.connect(unauthorizedCaller).rejectCancellation(0n, 1, 1))
                .to.be.revertedWithCustomError(ttmAccount, "AccessControlUnauthorizedAccount")
                .withArgs(unauthorizedCaller.address, BOOKING_OPERATOR_ROLE);
        });

        it("should revert counterCancellation if caller is not authorized", async function () {
            const { ttmAccount } = await loadFixture(deployAndConfigureAllFixture);

            const BOOKING_OPERATOR_ROLE = await ttmAccount.BOOKING_OPERATOR_ROLE();
            const unauthorizedCaller = signers.otherAccount1;

            // Try to counter cancellation with unauthorized caller
            await expect(
                ttmAccount.connect(unauthorizedCaller).counterCancellation(0n, ethers.parseEther("0.03"), 1, 1),
            )
                .to.be.revertedWithCustomError(ttmAccount, "AccessControlUnauthorizedAccount")
                .withArgs(unauthorizedCaller.address, BOOKING_OPERATOR_ROLE);
        });

        it("should revert withdrawCancellation if caller is not authorized", async function () {
            const { ttmAccount } = await loadFixture(deployAndConfigureAllFixture);

            const BOOKING_OPERATOR_ROLE = await ttmAccount.BOOKING_OPERATOR_ROLE();
            const unauthorizedCaller = signers.otherAccount1;

            // Try to withdraw cancellation with unauthorized caller
            await expect(ttmAccount.connect(unauthorizedCaller).withdrawCancellation(0n, 1, 1))
                .to.be.revertedWithCustomError(ttmAccount, "AccessControlUnauthorizedAccount")
                .withArgs(unauthorizedCaller.address, BOOKING_OPERATOR_ROLE);
        });

        it("should revert finalizeCancellation if caller is not authorized", async function () {
            const { ttmAccount } = await loadFixture(deployAndConfigureAllFixture);

            const BOOKING_OPERATOR_ROLE = await ttmAccount.BOOKING_OPERATOR_ROLE();
            const unauthorizedCaller = signers.otherAccount1;

            // Try to finalize cancellation with unauthorized caller
            await expect(ttmAccount.connect(unauthorizedCaller).finalizeCancellation(0n, ethers.parseEther("0.05")))
                .to.be.revertedWithCustomError(ttmAccount, "AccessControlUnauthorizedAccount")
                .withArgs(unauthorizedCaller.address, BOOKING_OPERATOR_ROLE);
        });
    });

    describe("Initializer validation", function () {
        it("should reject a zero address for any initializer parameter", async function () {
            await setupSigners();
            const { ttmAccountManager } = await loadFixture(deployTTMAccountManagerFixture);
            const zero = ethers.ZeroAddress;
            const ok = signers.ttmAccountAdmin.address;
            const mgr = await ttmAccountManager.getAddress();

            const bookingTokenOperator = await ethers.deployContract("BookingTokenOperator");
            const Account = await ethers.getContractFactory("TTMAccount", {
                libraries: { BookingTokenOperator: await bookingTokenOperator.getAddress() },
            });

            for (const args of [
                [zero, ok, ok, ok],
                [mgr, zero, ok, ok],
                [mgr, ok, zero, ok],
                [mgr, ok, ok, zero],
            ]) {
                await expect(
                    upgrades.deployProxy(Account, args, { kind: "uups", unsafeAllow: ["external-library-linking"] }),
                ).to.be.revertedWithCustomError(Account, "ZeroAddress");
            }
        });
    });

    describe("Deposit event", function () {
        it("should emit Deposit when receiving ETH", async function () {
            const { ttmAccount } = await loadFixture(deployAndConfigureAllFixture);
            const amount = ethers.parseEther("1.5");

            await expect(signers.otherAccount1.sendTransaction({ to: await ttmAccount.getAddress(), value: amount }))
                .to.emit(ttmAccount, "Deposit")
                .withArgs(signers.otherAccount1.address, amount);
        });
    });
});
