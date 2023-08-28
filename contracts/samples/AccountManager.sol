// SPDX-License-Identifier: GPL-3.0
pragma solidity ^0.8.12;

import "../interfaces/UserOperation.sol";
import "../interfaces/IAccountManager.sol";
import "../@openzeppelin/contracts/proxy/utils/Initializable.sol";
import "../core/NonceManager.sol";

contract AccountManager is IAccountManager, NonceManager, Initializable {

    address public masterWallet;

    event AccountManagerInitialized(address indexed masterWallet);

    mapping(address => bool) public authorizedUsers;

    constructor(address _masterWallet) {
        masterWallet = _masterWallet;
        _disableInitializers();
    }

    function authorizeUserOp(UserOperation calldata userOp) external view {
        require(authorizedUsers[userOp.sender], "User not authorized");
    }

    function addUserAuthorization(address user) external onlyMasterWallet returns (bool) {
        authorizedUsers[user] = true;
        return true;
    }

    function removeUserAuthorization(address user) external onlyMasterWallet returns (bool) {
        authorizedUsers[user] = false;
        return true;
    }

    modifier onlyMasterWallet() {
        require(msg.sender == masterWallet, "Only master wallet can call");
        _;
    }
    
    function initialize(address anOwner) public virtual initializer {
        _initialize(anOwner);
    }

    function _initialize(address anOwner) internal virtual {
        masterWallet = anOwner;
        emit AccountManagerInitialized(masterWallet);
    }

}
