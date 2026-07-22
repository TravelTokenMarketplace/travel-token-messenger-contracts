/**
 * @dev TTMAccount tests
 */
const { loadFixture } = require("@nomicfoundation/hardhat-toolbox/network-helpers");
const { expect } = require("chai");
const { ethers, upgrades, network } = require("hardhat");

const {
    setupSigners,
    serviceHash,
    deployTTMAccountManagerFixture,
    deployTTMAccountImplFixture,
    deployTTMAccountManagerWithTTMAccountImplFixture,
    deployAndConfigureAllFixture,
    deployAndConfigureAllWithRegisteredServicesFixture,
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

    describe("ERC165", function () {
        it("should report IERC721Receiver support without breaking pre-existing interfaces", async function () {
            await setupSigners();
            const { ttmAccount } = await loadFixture(deployAndConfigureAllFixture);

            const IERC165_INTERFACE_ID = "0x01ffc9a7";
            // bytes4(keccak256("onERC721Received(address,address,uint256,bytes)"))
            const IERC721_RECEIVER_INTERFACE_ID = "0x150b7a02";
            const IACCESSCONTROL_INTERFACE_ID = "0x7965db0b";
            const IACCESSCONTROLENUMERABLE_INTERFACE_ID = "0x5a05180f";

            // Pre-existing interfaces must still be reported - an override that forgets
            // `super.supportsInterface` would silently break these.
            expect(await ttmAccount.supportsInterface(IERC165_INTERFACE_ID)).to.be.true;
            expect(await ttmAccount.supportsInterface(IACCESSCONTROL_INTERFACE_ID)).to.be.true;
            expect(await ttmAccount.supportsInterface(IACCESSCONTROLENUMERABLE_INTERFACE_ID)).to.be.true;

            // The new interface this task adds.
            expect(await ttmAccount.supportsInterface(IERC721_RECEIVER_INTERFACE_ID)).to.be.true;

            // A bogus id must not be reported as supported - guards against an override
            // that returns true unconditionally instead of checking interfaceId.
            expect(await ttmAccount.supportsInterface("0xaaaaaaaa")).to.be.false;
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
            // GAS_WITHDRAWER_ROLE is no longer granted by addMessengerBot (Decision 5);
            // it must be granted explicitly by DEFAULT_ADMIN_ROLE.
            expect(await ttmAccount.hasRole(await ttmAccount.GAS_WITHDRAWER_ROLE(), bot.address)).to.be.false;
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

        it("should not grant GAS_WITHDRAWER_ROLE when adding a bot", async function () {
            const { ttmAccount } = await loadFixture(deployAndConfigureAllFixture);
            const bot = signers.otherAccount1;

            await ttmAccount.connect(signers.ttmAccountAdmin).addMessengerBot(bot.address, 0n);

            expect(await ttmAccount.hasRole(await ttmAccount.MESSENGER_BOT_ROLE(), bot.address)).to.be.true;
            expect(await ttmAccount.hasRole(await ttmAccount.BOOKING_OPERATOR_ROLE(), bot.address)).to.be.true;
            expect(await ttmAccount.hasRole(await ttmAccount.GAS_WITHDRAWER_ROLE(), bot.address)).to.be.false;
        });

        it("should let the default admin grant GAS_WITHDRAWER_ROLE explicitly", async function () {
            const { ttmAccount } = await loadFixture(deployAndConfigureAllFixture);
            const bot = signers.otherAccount1;
            const GAS_WITHDRAWER_ROLE = await ttmAccount.GAS_WITHDRAWER_ROLE();

            await ttmAccount.connect(signers.ttmAccountAdmin).addMessengerBot(bot.address, 0n);
            await ttmAccount.connect(signers.ttmAccountAdmin).grantRole(GAS_WITHDRAWER_ROLE, bot.address);

            expect(await ttmAccount.hasRole(GAS_WITHDRAWER_ROLE, bot.address)).to.be.true;

            // Removal still fully de-authorizes a bot granted the role later.
            await ttmAccount.connect(signers.ttmAccountAdmin).removeMessengerBot(bot.address);
            expect(await ttmAccount.hasRole(GAS_WITHDRAWER_ROLE, bot.address)).to.be.false;
        });

        it("should default the gas allowance to 0.01 ETH per 24 hours", async function () {
            const { ttmAccount } = await loadFixture(deployAndConfigureAllFixture);

            const [limit, period] = await ttmAccount.getGasMoneyWithdrawal();
            expect(limit).to.equal(ethers.parseEther("0.01"));
            expect(period).to.equal(24n * 60n * 60n);
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

    describe("recordExpiration", function () {
        it("should let an account with no roles genuinely expire a reservation once it has passed", async function () {
            await setupSigners();
            const { supplierTTMAccount, distributorTTMAccount, bookingToken } = await loadFixture(
                deployBookingTokenWithNullUSDFixture,
            );

            const tokenURI =
                "data:application/json;base64,eyJuYW1lIjoiQ2FtaW5vIE1lc3NlbmdlciBCb29raW5nVG9rZW4gVGVzdCJ9Cg==";
            const expirationTimestamp = Math.floor(Date.now() / 1000) + 120;
            const price = ethers.parseEther("0.05");

            // Grant BOOKING_OPERATOR_ROLE to mint the reservation, as usual.
            const BOOKING_OPERATOR_ROLE = await supplierTTMAccount.BOOKING_OPERATOR_ROLE();
            await expect(
                supplierTTMAccount
                    .connect(signers.ttmAccountAdmin)
                    .grantRole(BOOKING_OPERATOR_ROLE, signers.btAdmin.address),
            ).to.not.reverted;

            await supplierTTMAccount.connect(signers.btAdmin).mintBookingToken(
                distributorTTMAccount.getAddress(),
                tokenURI,
                expirationTimestamp,
                price,
                ethers.ZeroAddress, // native coin
                0,
                false,
            );

            // The caller below holds no role at all on this TTMAccount - not
            // BOOKING_OPERATOR_ROLE, not anything else.
            expect(await supplierTTMAccount.hasRole(BOOKING_OPERATOR_ROLE, signers.otherAccount3.address)).to.be.false;

            expect(await bookingToken.getBookingStatus(0n)).to.equal(1); // Reserved

            // Advance past the expiration timestamp.
            await network.provider.send("evm_setNextBlockTimestamp", [expirationTimestamp + 1]);
            await network.provider.send("evm_mine");

            // An unprivileged caller reaches the real BookingToken logic and genuinely
            // mutates state - this is not a mere "does not revert" check. If the role
            // gate were still in place, this call would revert with
            // AccessControlUnauthorizedAccount instead of emitting and updating status.
            await expect(supplierTTMAccount.connect(signers.otherAccount3).recordExpiration(0n))
                .to.emit(bookingToken, "TokenReservationExpired")
                .withArgs(0n);

            expect(await bookingToken.getBookingStatus(0n)).to.equal(2); // Expired
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

    describe("Hash-native services", function () {
        it("should add a service by hash and emit the hash", async function () {
            await setupSigners();
            const { ttmAccount, ttmAccountManager } = await loadFixture(
                deployAndConfigureAllWithRegisteredServicesFixture,
            );

            const name = "ttm.services.accommodation.v1alpha.AccommodationSearchService";
            await ttmAccountManager.connect(signers.registryAdmin).registerService(name);
            // Independently computed (not read back from the contract): this pins the hash to
            // the exact keccak256 the registry itself uses, so a subtly wrong storage key would
            // still be caught even if the contract were merely self-consistent.
            const hash = ethers.keccak256(ethers.toUtf8Bytes(name));

            await expect(ttmAccount.connect(signers.ttmServiceAdmin).addService(hash, false, []))
                .to.emit(ttmAccount, "ServiceAdded")
                .withArgs(hash);

            expect(await ttmAccount.getAllServiceHashes()).to.deep.equal([hash]);
            expect(await ttmAccount.getService(hash)).to.deep.equal([false, []]);
        });

        it("should reject adding a service whose hash is not registered", async function () {
            await setupSigners();
            const { ttmAccount } = await loadFixture(deployAndConfigureAllWithRegisteredServicesFixture);

            // ttmServiceAdmin genuinely holds SERVICE_ADMIN_ROLE here (granted by the fixture),
            // so this must revert on the registry lookup, not on access control.
            const unregistered = ethers.keccak256(ethers.toUtf8Bytes("ttm.services.nope.v1.NopeService"));

            await expect(
                ttmAccount.connect(signers.ttmServiceAdmin).addService(unregistered, false, []),
            ).to.be.revertedWithCustomError(ttmAccount, "ServiceNotRegistered");
        });

        it("should remove every service without resolving names", async function () {
            await setupSigners();
            const { ttmAccount, ttmAccountManager } = await loadFixture(
                deployAndConfigureAllWithRegisteredServicesFixture,
            );

            const names = [
                "ttm.services.accommodation.v1alpha.AccommodationSearchService",
                "ttm.services.activity.v2.ActivitySearchService",
            ];
            const hashes = [];
            for (const name of names) {
                await ttmAccountManager.connect(signers.registryAdmin).registerService(name);
                const hash = serviceHash(name);
                hashes.push(hash);
                await ttmAccount.connect(signers.ttmServiceAdmin).addService(hash, false, []);
            }
            expect(await ttmAccount.getAllServiceHashes()).to.deep.equal(hashes);

            await expect(ttmAccount.connect(signers.ttmServiceAdmin).removeAllServices())
                .to.emit(ttmAccount, "ServiceRemoved")
                .withArgs(hashes[0])
                .to.emit(ttmAccount, "ServiceRemoved")
                .withArgs(hashes[1]);

            expect(await ttmAccount.getAllServiceHashes()).to.deep.equal([]);
            for (const hash of hashes) {
                await expect(ttmAccount.getService(hash)).to.be.revertedWithCustomError(
                    ttmAccount,
                    "ServiceDoesNotExist",
                );
            }
        });

        it("should manage capabilities by hash, keeping capability strings readable", async function () {
            await setupSigners();
            const { ttmAccount, ttmAccountManager } = await loadFixture(
                deployAndConfigureAllWithRegisteredServicesFixture,
            );

            const name = "ttm.services.accommodation.v1alpha.AccommodationSearchService";
            await ttmAccountManager.connect(signers.registryAdmin).registerService(name);
            const hash = serviceHash(name);
            await ttmAccount.connect(signers.ttmServiceAdmin).addService(hash, false, []);

            await expect(ttmAccount.connect(signers.ttmServiceAdmin).addServiceCapability(hash, "luggage"))
                .to.emit(ttmAccount, "ServiceCapabilityAdded")
                .withArgs(hash, "luggage");

            expect(await ttmAccount["getServiceCapabilities(bytes32)"](hash)).to.deep.equal(["luggage"]);

            await expect(ttmAccount.connect(signers.ttmServiceAdmin).removeServiceCapability(hash, "luggage"))
                .to.emit(ttmAccount, "ServiceCapabilityRemoved")
                .withArgs(hash, "luggage");

            expect(await ttmAccount["getServiceCapabilities(bytes32)"](hash)).to.deep.equal([]);
        });

        it("should add and remove wanted services by hash", async function () {
            await setupSigners();
            const { ttmAccount, ttmAccountManager } = await loadFixture(
                deployAndConfigureAllWithRegisteredServicesFixture,
            );

            const name = "ttm.services.activity.v2.ActivitySearchService";
            await ttmAccountManager.connect(signers.registryAdmin).registerService(name);
            // Independently computed, same reasoning as the ServiceAdded test above: pins the
            // hash to the registry's own keccak256 rather than trusting a value read back from
            // the contract, so a wrong storage key would still be caught.
            const hash = ethers.keccak256(ethers.toUtf8Bytes(name));

            await expect(ttmAccount.connect(signers.ttmServiceAdmin).addWantedServices([hash]))
                .to.emit(ttmAccount, "WantedServiceAdded")
                .withArgs(hash);

            expect(await ttmAccount.getWantedServiceHashes()).to.deep.equal([hash]);

            await expect(ttmAccount.connect(signers.ttmServiceAdmin).removeWantedServices([hash]))
                .to.emit(ttmAccount, "WantedServiceRemoved")
                .withArgs(hash);

            expect(await ttmAccount.getWantedServiceHashes()).to.deep.equal([]);
        });

        it("should reject adding a wanted service whose hash is not registered", async function () {
            await setupSigners();
            const { ttmAccount } = await loadFixture(deployAndConfigureAllWithRegisteredServicesFixture);

            // ttmServiceAdmin genuinely holds SERVICE_ADMIN_ROLE here (granted by the fixture),
            // so this must revert on the registry lookup, not on access control.
            const unregistered = ethers.keccak256(ethers.toUtf8Bytes("ttm.services.nope.v1.NopeService"));

            await expect(
                ttmAccount.connect(signers.ttmServiceAdmin).addWantedServices([unregistered]),
            ).to.be.revertedWithCustomError(ttmAccount, "ServiceNotRegistered");
        });

        it("should return supported services as hashes with no manager round-trip", async function () {
            await setupSigners();
            // deployAndConfigureAllFixture alone does not grant ttmServiceAdmin the
            // SERVICE_ADMIN_ROLE; only the WithRegisteredServices variant does.
            const { ttmAccount, ttmAccountManager } = await loadFixture(
                deployAndConfigureAllWithRegisteredServicesFixture,
            );

            const names = [
                "ttm.services.accommodation.v1alpha.AccommodationSearchService",
                "ttm.services.activity.v2.ActivitySearchService",
            ];
            const hashes = [];
            for (const name of names) {
                await ttmAccountManager.connect(signers.registryAdmin).registerService(name);
                const hash = serviceHash(name);
                hashes.push(hash);
                await ttmAccount.connect(signers.ttmServiceAdmin).addService(hash, true, ["luggage"]);
            }

            const [serviceHashes, services] = await ttmAccount.getSupportedServices();
            expect(serviceHashes).to.deep.equal(hashes);
            expect(services).to.have.lengthOf(2);
            expect(services[0][0]).to.be.true; // _restrictedRate

            const [pageHashes] = await ttmAccount.getSupportedServicesSlice(1, 5);
            expect(pageHashes).to.deep.equal([hashes[1]]);

            // The string-typed API is gone.
            expect(ttmAccount.interface.fragments.some((f) => f.name === "getServiceHash")).to.be.false;
        });

        describe("getSupportedServicesSlice pagination", function () {
            async function threeServicesFixture() {
                await setupSigners();
                const { ttmAccount, ttmAccountManager } = await loadFixture(
                    deployAndConfigureAllWithRegisteredServicesFixture,
                );

                const names = [
                    "ttm.services.accommodation.v1alpha.AccommodationSearchService",
                    "ttm.services.activity.v2.ActivitySearchService",
                    "ttm.services.transport.v1.TransportSearchService",
                ];
                const hashes = [];
                for (const name of names) {
                    await ttmAccountManager.connect(signers.registryAdmin).registerService(name);
                    const hash = serviceHash(name);
                    hashes.push(hash);
                    await ttmAccount.connect(signers.ttmServiceAdmin).addService(hash, false, []);
                }
                return { ttmAccount, hashes };
            }

            it("returns an empty window when offset is at or past the end", async function () {
                const { ttmAccount, hashes } = await threeServicesFixture();

                const [atEnd] = await ttmAccount.getSupportedServicesSlice(hashes.length, 5);
                expect(atEnd).to.deep.equal([]);

                const [pastEnd] = await ttmAccount.getSupportedServicesSlice(hashes.length + 10, 5);
                expect(pastEnd).to.deep.equal([]);
            });

            it("clamps the window when limit exceeds the remainder, without reverting", async function () {
                const { ttmAccount, hashes } = await threeServicesFixture();

                // offset = 1, limit far larger than the 2 remaining entries: must clamp to
                // what's left, not revert and not return garbage beyond the array bounds.
                const [pageHashes, pageServices] = await ttmAccount.getSupportedServicesSlice(1, 1000);
                expect(pageHashes).to.deep.equal(hashes.slice(1));
                expect(pageServices).to.have.lengthOf(2);
            });

            it("returns an empty window for limit == 0", async function () {
                const { ttmAccount } = await threeServicesFixture();

                const [pageHashes, pageServices] = await ttmAccount.getSupportedServicesSlice(0, 0);
                expect(pageHashes).to.deep.equal([]);
                expect(pageServices).to.deep.equal([]);
            });

            it("does not revert for an oversized limit at a non-zero offset", async function () {
                const { ttmAccount, hashes } = await threeServicesFixture();

                // A naive `offset + limit` bound check reverts under Solidity 0.8 checked
                // arithmetic when that sum overflows uint256; clamping by subtraction
                // (total - offset) must not. Offset must be non-zero here: at offset 0,
                // `0 + limit` never overflows regardless of how large `limit` is, so that
                // case can't distinguish the additive form from the subtractive one. This
                // pins the exact overflow shape Task 1 shipped.
                const maxUint256 = 2n ** 256n - 1n;
                const [pageHashes] = await ttmAccount.getSupportedServicesSlice(1, maxUint256);
                expect(pageHashes).to.deep.equal(hashes.slice(1));
            });
        });
    });
});
