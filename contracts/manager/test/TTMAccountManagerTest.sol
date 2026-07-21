// SPDX-License-Identifier: LGPL-3.0-or-later
//
// Travel Token Messenger Account Manager V2 for Testing Upgrades

/**
 * TESTING ONLY - NOT FOR PRODUCTION
 */

pragma solidity 0.8.24;

import { TTMAccountManager } from "../TTMAccountManager.sol";

contract TTMAccountManagerTest is TTMAccountManager {
    function getVersion() public pure returns (string memory) {
        return "TESTING";
    }
}
