// SPDX-License-Identifier: LGPL-3.0-or-later

pragma solidity 0.8.24;

interface ITTMAccount {
    function initialize(address manager, address bookingToken, address owner, address upgrader) external;

    /**
     * @notice Whether a payment token is declared as supported by this account.
     *
     * Payment mode is encoded as an address, matching BookingToken:
     * `address(0)` is native currency, `address(1)` is off-chain payment, and
     * any other value is an ERC-20 address. All three are declared through the
     * same allowlist.
     */
    function isSupportedToken(address _token) external view returns (bool);
}
