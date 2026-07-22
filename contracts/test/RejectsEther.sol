// SPDX-License-Identifier: LGPL-3.0-or-later
pragma solidity 0.8.24;

/**
 * @notice A contract that cannot receive ETH: no `receive`, no `fallback`.
 *
 * @dev Used to exercise the cancellation refund path against a counterparty whose
 * address rejects a plain transfer. See Decision 3 in
 * docs/decisions/2026-07-21-contract-design-decisions.md.
 */
contract RejectsEther {
    /// @notice Lets tests confirm the contract deployed and has code.
    function ping() external pure returns (bool) {
        return true;
    }
}
