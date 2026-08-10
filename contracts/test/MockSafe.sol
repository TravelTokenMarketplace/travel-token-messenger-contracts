// SPDX-License-Identifier: LGPL-3.0-or-later
pragma solidity 0.8.24;

/**
 * @notice Minimal stand-in for a Gnosis Safe: the two owner-set getters the
 * `roles handoff` preflight probes, over a `SafeProxy`-shaped storage layout.
 *
 * @dev `singleton` is declared first so it occupies storage slot 0, exactly as
 * `SafeProxy` does — that is the slot the preflight reads to report which Safe
 * implementation the address is running. Used to exercise the custody-type
 * preflight in `tasks/lib/preflight.js` without pulling the Safe contracts into
 * this repo's dependency tree.
 */
contract MockSafe {
    /// @dev Slot 0, mirroring `SafeProxy.singleton`.
    address internal singleton;

    address[] private _owners;
    uint256 private _threshold;

    constructor(address singleton_, address[] memory owners_, uint256 threshold_) {
        singleton = singleton_;
        _owners = owners_;
        _threshold = threshold_;
    }

    /// @notice Mirrors `Safe.getOwners()`.
    function getOwners() external view returns (address[] memory) {
        return _owners;
    }

    /// @notice Mirrors `Safe.getThreshold()`.
    function getThreshold() external view returns (uint256) {
        return _threshold;
    }
}
