// SPDX-License-Identifier: LGPL-3.0-or-later
//
// Travel Token Messenger Gas Money Manager

pragma solidity 0.8.24;

import { Initializable } from "@openzeppelin/contracts-upgradeable/proxy/utils/Initializable.sol";
import { Address } from "@openzeppelin/contracts/utils/Address.sol";

/**
 * @title GasMoneyManager
 * @notice GasMoneyManager manages gas money withdrawals for a {TTMAccount}.
 *
 * Gas money withdrawals are restricted to a withdrawal limit and period.
 */
abstract contract GasMoneyManager is Initializable {
    using Address for address payable;

    /***************************************************
     *                   STORAGE                       *
     ***************************************************/

    /**
     * @notice Per-account withdrawal accounting, packed into a single slot.
     */
    struct GasMoneyWithdrawalRecord {
        uint128 amount; // wei withdrawn in the current period
        uint64 periodStart; // unix timestamp of the current period start
    }

    /// @custom:storage-location erc7201:traveltoken.messenger.storage.GasMoneyManager
    struct GasMoneyStorage {
        mapping(address => GasMoneyWithdrawalRecord) _withdrawals;
        uint128 _withdrawalLimit;
        uint64 _withdrawalPeriod;
    }

    // keccak256(abi.encode(uint256(keccak256("traveltoken.messenger.storage.GasMoneyManager")) - 1)) & ~bytes32(uint256(0xff));
    bytes32 private constant GasMoneyStorageLocation =
        0xc7ae2c65fdae475b1bb2dd4079b252a7d893d0822b91a40dd25c763e583f7700;

    function _getGasMoneyStorage() private pure returns (GasMoneyStorage storage $) {
        assembly {
            $.slot := GasMoneyStorageLocation
        }
    }

    /***************************************************
     *                   EVENTS                        *
     ***************************************************/

    /**
     * @notice Gas money withdrawal event
     *
     * @param withdrawer the address of the withdrawer
     * @param amount the amount withdrawn
     */
    event GasMoneyWithdrawal(address indexed withdrawer, uint256 amount);

    /**
     * @notice Gas money withdrawal limit and period updated event
     *
     * @param limit the withdrawal limit for the period
     * @param period the withdrawal period in seconds
     */
    event GasMoneyWithdrawalUpdated(uint256 limit, uint256 period);

    /***************************************************
     *                   ERRORS                        *
     ***************************************************/

    error WithdrawalLimitExceeded(uint256 limit, uint256 amount);
    error WithdrawalLimitExceededForPeriod(uint256 limit, uint256 amount);
    error GasMoneyValueOutOfRange(uint256 limit, uint256 period);

    /***************************************************
     *               INITIALIZATION                    *
     ***************************************************/

    function __GasMoneyManager_init(uint256 withdrawalLimit, uint256 withdrawalPeriod) internal onlyInitializing {
        GasMoneyStorage storage $ = _getGasMoneyStorage();
        $._withdrawalLimit = _toUint128(withdrawalLimit, withdrawalLimit, withdrawalPeriod);
        $._withdrawalPeriod = _toUint64(withdrawalPeriod, withdrawalLimit, withdrawalPeriod);
    }

    /***************************************************
     *                   LOGIC                        *
     ***************************************************/

    /**
     * @notice Withdraws gas money.
     *
     * This functions is intended to be called by the bot to withdraw gas money.
     * Inheriting contract should restrict who can call this with a public
     * function.
     */
    function _withdrawGasMoney(uint256 amount) internal {
        GasMoneyStorage storage $ = _getGasMoneyStorage();

        uint256 limit = $._withdrawalLimit;

        // Ensure the withdrawal does not exceed the allowed limit
        if (amount > limit) {
            revert WithdrawalLimitExceeded(limit, amount);
        }

        GasMoneyWithdrawalRecord memory withdrawal = $._withdrawals[msg.sender];
        uint256 currentTime = block.timestamp;

        // Reset the withdrawn amount if a new period has started. If more time than
        // the withdrawal period has passed, it is allowed to withdraw the full amount.
        if (currentTime > uint256(withdrawal.periodStart) + $._withdrawalPeriod) {
            withdrawal.amount = 0;
            withdrawal.periodStart = _toUint64(currentTime, limit, $._withdrawalPeriod);
        }

        // Ensure the withdrawal does not exceed the allowed limit for the period
        if (uint256(withdrawal.amount) + amount > limit) {
            revert WithdrawalLimitExceededForPeriod(limit, amount);
        }

        // Update the withdrawn amount. Safe: the sum was just checked against
        // limit, which is itself a uint128.
        withdrawal.amount = uint128(uint256(withdrawal.amount) + amount);
        $._withdrawals[msg.sender] = withdrawal;

        // Transfer the gas money
        payable(msg.sender).sendValue(amount);

        emit GasMoneyWithdrawal(msg.sender, amount);
    }

    /**
     * @notice Sets the gas money withdrawal limit and period.
     *
     * @param limit the withdrawal limit for the period
     * @param period the withdrawal period in seconds
     */
    function _setGasMoneyWithdrawal(uint256 limit, uint256 period) internal {
        GasMoneyStorage storage $ = _getGasMoneyStorage();
        $._withdrawalLimit = _toUint128(limit, limit, period);
        $._withdrawalPeriod = _toUint64(period, limit, period);

        emit GasMoneyWithdrawalUpdated(limit, period);
    }

    /**
     * @notice Returns the gas money withdrawal restrictions.
     *
     * @return withdrawalLimit
     * @return withdrawalPeriod
     */
    function getGasMoneyWithdrawal() public view returns (uint256 withdrawalLimit, uint256 withdrawalPeriod) {
        GasMoneyStorage storage $ = _getGasMoneyStorage();
        return (uint256($._withdrawalLimit), uint256($._withdrawalPeriod));
    }

    /**
     * @notice Returns the gas money withdrawal details for an account.
     *
     * @param account address of the account
     * @return periodStart timestamp of the withdrawal period start
     * @return withdrawnAmount amount withdrawn within the period
     */
    function getGasMoneyWithdrawalForAccount(
        address account
    ) public view returns (uint256 periodStart, uint256 withdrawnAmount) {
        GasMoneyStorage storage $ = _getGasMoneyStorage();
        GasMoneyWithdrawalRecord memory withdrawal = $._withdrawals[account];
        return (uint256(withdrawal.periodStart), uint256(withdrawal.amount));
    }

    /***************************************************
     *                   HELPERS                       *
     ***************************************************/

    function _toUint128(uint256 value, uint256 limit, uint256 period) private pure returns (uint128) {
        if (value > type(uint128).max) {
            revert GasMoneyValueOutOfRange(limit, period);
        }
        return uint128(value);
    }

    function _toUint64(uint256 value, uint256 limit, uint256 period) private pure returns (uint64) {
        if (value > type(uint64).max) {
            revert GasMoneyValueOutOfRange(limit, period);
        }
        return uint64(value);
    }
}
