// SPDX-License-Identifier: LGPL-3.0-or-later

pragma solidity 0.8.24;

interface ITTMAccount {
    function initialize(address manager, address bookingToken, address owner, address upgrader) external;
}
