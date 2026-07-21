// SPDX-License-Identifier: LGPL-3.0-or-later

pragma solidity 0.8.24;

interface ITTMAccountManager {
    function getAccountImplementation() external view returns (address);

    function isTTMAccount(address account) external view returns (bool);

    /// @dev Reverts if the hash is not registered. This is the sole remaining
    /// manager dependency of TTMAccount, used to validate {addService}.
    function getRegisteredServiceNameByHash(bytes32 serviceHash) external view returns (string memory serviceName);
}
