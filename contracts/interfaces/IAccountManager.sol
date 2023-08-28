// SPDX-License-Identifier: GPL-3.0
pragma solidity ^0.8.12;

import "./UserOperation.sol";

interface IAccountManager {

function authorizeUserOp(UserOperation calldata userOp) external view;

function addUserAuthorization(address user) external returns (bool);

function removeUserAuthorization(address user) external returns (bool);

}