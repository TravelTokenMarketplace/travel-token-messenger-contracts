/**
 * @dev TTMAccount tests
 */
const { loadFixture } = require("@nomicfoundation/hardhat-toolbox/network-helpers");
const { expect } = require("chai");
const { ethers } = require("hardhat");

const {
    setupSigners,
    deployTTMAccountManagerFixture,
    deployTTMAccountImplFixture,
    deployTTMAccountManagerWithTTMAccountImplFixture,
    deployAndConfigureAllFixture,
    deployTTMAccountWithDepositFixture,
} = require("./utils/fixtures");

describe("GasMoneyManager", function () {
    describe("Main", function () {
        it("should initialize gas money manager correctly", async function () {
            const { ttmAccount } = await loadFixture(deployAndConfigureAllFixture);

            const expectedLimit = ethers.parseEther("10"); // 10 CAM
            const expectedPeriod = 24 * 60 * 60; // 24 hours

            expect(await ttmAccount.getGasMoneyWithdrawal()).to.be.deep.equal([expectedLimit, expectedPeriod]);
            //expect(await ttmAccount.getGasMoneyWithdrawalPeriod()).to.be.equal(expectedPeriod);
        });

        it("should set gas money limit and period correctly", async function () {
            const { ttmAccount } = await loadFixture(deployAndConfigureAllFixture);

            const expectedLimit = ethers.parseEther("10"); // 10 CAM
            const expectedPeriod = 24 * 60 * 60; // 24 hours

            expect(await ttmAccount.getGasMoneyWithdrawal()).to.be.deep.equal([expectedLimit, expectedPeriod]);
            //expect(await ttmAccount.getGasMoneyWithdrawalPeriod()).to.be.equal(expectedPeriod);

            const newLimit = ethers.parseEther("20"); // 20 CAM
            const newPeriod = 48 * 60 * 60; // 48 hours

            await expect(ttmAccount.connect(signers.ttmAccountAdmin).setGasMoneyWithdrawal(newLimit, newPeriod))
                .to.emit(ttmAccount, "GasMoneyWithdrawalUpdated")
                .withArgs(newLimit, newPeriod);

            // Try with non-auth address
            await expect(ttmAccount.connect(signers.otherAccount1).setGasMoneyWithdrawal(newLimit, newPeriod))
                .to.be.revertedWithCustomError(ttmAccount, "AccessControlUnauthorizedAccount")
                .withArgs(signers.otherAccount1.address, await ttmAccount.BOT_ADMIN_ROLE());

            // await expect(ttmAccount.connect(signers.ttmAccountAdmin).setGasMoneyWithdrawalPeriod(newPeriod))
            //     .to.emit(ttmAccount, "GasMoneyWithdrawalPeriodUpdated")
            //     .withArgs(newPeriod);

            expect(await ttmAccount.getGasMoneyWithdrawal()).to.be.deep.equal([newLimit, newPeriod]);
            //expect(await ttmAccount.getGasMoneyWithdrawalPeriod()).to.be.equal(newPeriod);
        });

        it("should withdraw gas money", async function () {
            const { ttmAccount } = await loadFixture(deployAndConfigureAllFixture);

            const withdrawer = signers.withdrawer;

            // Add more funds to the ttmAccount so we are not under the prefund spent
            const depositAmount = ethers.parseEther("100");

            const depositTx = {
                to: ttmAccount.getAddress(),
                value: depositAmount,
            };

            const txResponse = await signers.depositor.sendTransaction(depositTx);
            await txResponse.wait();

            // Register withdrawer as a bot
            await expect(ttmAccount.connect(signers.ttmAccountAdmin).addMessengerBot(withdrawer.address, 0n))
                .to.emit(ttmAccount, "MessengerBotAdded")
                .withArgs(withdrawer.address);

            // Withdraw
            const withdrawAmount = ethers.parseEther("1");

            const withdrawTx = ttmAccount.connect(withdrawer).withdrawGasMoney(withdrawAmount);
            await expect(withdrawTx).to.changeEtherBalances(
                [ttmAccount, withdrawer],
                [-withdrawAmount, withdrawAmount],
            );
            await expect(withdrawTx)
                .to.emit(ttmAccount, "GasMoneyWithdrawal")
                .withArgs(withdrawer.address, withdrawAmount);
        });

        it("should revert if not allowed to withdraw gas money", async function () {
            const { ttmAccount } = await loadFixture(deployAndConfigureAllFixture);

            const withdrawer = signers.withdrawer;

            // Add more funds to the ttmAccount so we are not under the prefund spent
            const depositAmount = ethers.parseEther("100");

            const depositTx = {
                to: ttmAccount.getAddress(),
                value: depositAmount,
            };

            const txResponse = await signers.depositor.sendTransaction(depositTx);
            await txResponse.wait();

            // Do not add withdrawer as a bot.

            // Withdraw
            const withdrawAmount = ethers.parseEther("1");

            await expect(ttmAccount.connect(withdrawer).withdrawGasMoney(withdrawAmount))
                .to.revertedWithCustomError(ttmAccount, "AccessControlUnauthorizedAccount")
                .withArgs(withdrawer.address, ttmAccount.GAS_WITHDRAWER_ROLE());
        });

        it("should revert if amount is over the limit", async function () {
            const { ttmAccount } = await loadFixture(deployAndConfigureAllFixture);

            const withdrawer = signers.withdrawer;

            // Add more funds to the ttmAccount so we are not under the prefund spent
            const depositAmount = ethers.parseEther("100");

            const depositTx = {
                to: ttmAccount.getAddress(),
                value: depositAmount,
            };

            const txResponse = await signers.depositor.sendTransaction(depositTx);
            await txResponse.wait();

            const expectedLimit = ethers.parseEther("10"); // 10 CAM

            // Register withdrawer as a bot
            await expect(ttmAccount.connect(signers.ttmAccountAdmin).addMessengerBot(withdrawer.address, 0n))
                .to.emit(ttmAccount, "MessengerBotAdded")
                .withArgs(withdrawer.address);

            // Withdraw
            const withdrawAmount = ethers.parseEther("11"); // 11 CAM, over the limit

            await expect(ttmAccount.connect(withdrawer).withdrawGasMoney(withdrawAmount))
                .to.revertedWithCustomError(ttmAccount, "WithdrawalLimitExceeded")
                .withArgs(expectedLimit, withdrawAmount);
        });

        it("should revert if amount is over the limit for the period", async function () {
            const { ttmAccount } = await loadFixture(deployAndConfigureAllFixture);

            const withdrawer = signers.withdrawer;

            // Add more funds to the ttmAccount so we are not under the prefund spent
            const depositAmount = ethers.parseEther("100");

            const depositTx = {
                to: ttmAccount.getAddress(),
                value: depositAmount,
            };

            const txResponse = await signers.depositor.sendTransaction(depositTx);
            await txResponse.wait();

            const expectedLimit = ethers.parseEther("10"); // 10 CAM

            // Register withdrawer as a bot
            await expect(ttmAccount.connect(signers.ttmAccountAdmin).addMessengerBot(withdrawer.address, 0n))
                .to.emit(ttmAccount, "MessengerBotAdded")
                .withArgs(withdrawer.address);

            // Withdraw
            const withdrawAmount = ethers.parseEther("1"); // Start with 1 CAM

            const withdrawTx1 = ttmAccount.connect(withdrawer).withdrawGasMoney(withdrawAmount);
            await expect(withdrawTx1).to.changeEtherBalances(
                [ttmAccount, withdrawer],
                [-withdrawAmount, withdrawAmount],
            );

            // Get block
            const block = await ethers.provider.getBlock("latest");

            await expect(withdrawTx1)
                .to.emit(ttmAccount, "GasMoneyWithdrawal")
                .withArgs(withdrawer.address, withdrawAmount);

            const withdrawAmount2 = ethers.parseEther("7"); // Withdraw 7 CAM, total 8

            const withdrawTx2 = ttmAccount.connect(withdrawer).withdrawGasMoney(withdrawAmount2);
            await expect(withdrawTx2).to.changeEtherBalances(
                [ttmAccount, withdrawer],
                [-withdrawAmount2, withdrawAmount2],
            );
            await expect(withdrawTx2)
                .to.emit(ttmAccount, "GasMoneyWithdrawal")
                .withArgs(withdrawer.address, withdrawAmount2);

            const withdrawAmount3 = ethers.parseEther("3"); // Withdraw 3 CAM, total 11, over the limit

            await expect(ttmAccount.connect(withdrawer).withdrawGasMoney(withdrawAmount3))
                .to.revertedWithCustomError(ttmAccount, "WithdrawalLimitExceededForPeriod")
                .withArgs(expectedLimit, withdrawAmount3);

            // Get withdrawal details for the withdrawer
            expect(await ttmAccount.getGasMoneyWithdrawalForAccount(withdrawer.address)).to.be.deep.equal([
                block.timestamp, // withdrawal start time (the first block that we withdrew)
                ethers.parseEther("8"), // We withdrawn 8 CAM
            ]);
        });

        it("should allow withdrawal after period resets", async function () {
            const { ttmAccount } = await loadFixture(deployAndConfigureAllFixture);

            const withdrawer = signers.withdrawer;

            // Add more funds to the ttmAccount so we are not under the prefund spent
            const depositAmount = ethers.parseEther("100");

            const depositTx = {
                to: ttmAccount.getAddress(),
                value: depositAmount,
            };

            const txResponse = await signers.depositor.sendTransaction(depositTx);
            await txResponse.wait();

            const expectedLimit = ethers.parseEther("10"); // 10 CAM

            // Register withdrawer as a bot
            await expect(ttmAccount.connect(signers.ttmAccountAdmin).addMessengerBot(withdrawer.address, 0n))
                .to.emit(ttmAccount, "MessengerBotAdded")
                .withArgs(withdrawer.address);

            // Withdraw
            const withdrawAmount = ethers.parseEther("10"); // withdraw all 10 CAM

            const withdrawTx1 = ttmAccount.connect(withdrawer).withdrawGasMoney(withdrawAmount);
            await expect(withdrawTx1).to.changeEtherBalances(
                [ttmAccount, withdrawer],
                [-withdrawAmount, withdrawAmount],
            );
            await expect(withdrawTx1)
                .to.emit(ttmAccount, "GasMoneyWithdrawal")
                .withArgs(withdrawer.address, withdrawAmount);

            const withdrawAmount2 = ethers.parseEther("3"); // Try to withdraw 3 CAM, over the limit

            await expect(ttmAccount.connect(withdrawer).withdrawGasMoney(withdrawAmount2))
                .to.revertedWithCustomError(ttmAccount, "WithdrawalLimitExceededForPeriod")
                .withArgs(expectedLimit, withdrawAmount2);

            // Advance time by 24 hours
            await network.provider.send("evm_increaseTime", [24 * 60 * 60]);
            await network.provider.send("evm_mine");

            // Withdraw again
            const withdrawAmount3 = ethers.parseEther("10"); // Try to withdraw the limit as the period has been reset

            const withdrawTx3 = ttmAccount.connect(withdrawer).withdrawGasMoney(withdrawAmount3);
            await expect(withdrawTx3).to.changeEtherBalances(
                [ttmAccount, withdrawer],
                [-withdrawAmount3, withdrawAmount3],
            );
            await expect(withdrawTx3)
                .to.emit(ttmAccount, "GasMoneyWithdrawal")
                .withArgs(withdrawer.address, withdrawAmount3);
        });
    });

    describe("Storage packing", function () {
        it("should still return uint256 from the getters", async function () {
            const { ttmAccount } = await loadFixture(deployAndConfigureAllFixture);
            const [limit, period] = await ttmAccount.getGasMoneyWithdrawal();
            expect(limit).to.equal(ethers.parseEther("10"));
            expect(period).to.equal(86400n);

            const iface = ttmAccount.interface.getFunction("getGasMoneyWithdrawal");
            expect(iface.outputs.map((o) => o.type)).to.deep.equal(["uint256", "uint256"]);

            const accIface = ttmAccount.interface.getFunction("getGasMoneyWithdrawalForAccount");
            expect(accIface.outputs.map((o) => o.type)).to.deep.equal(["uint256", "uint256"]);
        });

        it("should accept values at the uint128/uint64 bounds", async function () {
            const { ttmAccount } = await loadFixture(deployAndConfigureAllFixture);
            const maxLimit = 2n ** 128n - 1n;
            const maxPeriod = 2n ** 64n - 1n;

            await ttmAccount.connect(signers.ttmAccountAdmin).setGasMoneyWithdrawal(maxLimit, maxPeriod);

            const [limit, period] = await ttmAccount.getGasMoneyWithdrawal();
            expect(limit).to.equal(maxLimit);
            expect(period).to.equal(maxPeriod);
        });

        it("should revert rather than truncate when the limit overflows uint128", async function () {
            const { ttmAccount } = await loadFixture(deployAndConfigureAllFixture);
            const tooBig = 2n ** 128n;

            await expect(
                ttmAccount.connect(signers.ttmAccountAdmin).setGasMoneyWithdrawal(tooBig, 86400n),
            ).to.be.revertedWithCustomError(ttmAccount, "GasMoneyValueOutOfRange");
        });

        it("should revert rather than truncate when the period overflows uint64", async function () {
            const { ttmAccount } = await loadFixture(deployAndConfigureAllFixture);
            const tooBig = 2n ** 64n;

            await expect(
                ttmAccount.connect(signers.ttmAccountAdmin).setGasMoneyWithdrawal(ethers.parseEther("10"), tooBig),
            ).to.be.revertedWithCustomError(ttmAccount, "GasMoneyValueOutOfRange");
        });
    });
});
