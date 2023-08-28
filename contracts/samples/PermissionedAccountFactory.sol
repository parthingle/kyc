// SPDX-License-Identifier: GPL-3.0
pragma solidity ^0.8.12;

import "../@openzeppelin/contracts/utils/Create2.sol";
import "../@openzeppelin/contracts/proxy/ERC1967/ERC1967Proxy.sol";

import "./PermissionedAccount.sol";

/**
 * A sample factory contract for PermissionedAccount
 * A UserOperations "initCode" holds the address of the factory, and a method call (to createAccount, in this sample factory).
 * The factory's createAccount returns the target account address even if it is already installed.
 * This way, the entryPoint.getSenderAddress() can be called either before or after the account is created.
 */
contract PermissionedAccountFactory {
    PermissionedAccount public immutable accountImplementation;
    IAccountManager public immutable _accountManager;

    constructor(IEntryPoint _entryPoint, IAccountManager accountManager) {
        accountImplementation = new PermissionedAccount(_entryPoint, accountManager);
        _accountManager = accountManager;
    }

    /**
     * create an account, and return its address.
     * returns the address even if the account is already deployed.
     * Note that during UserOperation execution, this method is called only if the account is not deployed.
     * This method returns an existing account address so that entryPoint.getSenderAddress() would work even after account creation
     */
    function createAccount(address owner,uint256 salt) private onlyAccountManager returns (PermissionedAccount ret) {
        address addr = getAddress(owner, salt);
        uint codeSize = addr.code.length;
        if (codeSize > 0) {
            return PermissionedAccount(payable(addr));
        }
        ret = PermissionedAccount(payable(new ERC1967Proxy{salt : bytes32(salt)}(
                address(accountImplementation),
                abi.encodeCall(PermissionedAccount.initialize, (owner))
            )));
    }

    /**
     * calculate the counterfactual address of this account as it would be returned by createAccount()
     */
    function getAddress(address owner,uint256 salt) private onlyAccountManager view returns (address) {
        return Create2.computeAddress(bytes32(salt), keccak256(abi.encodePacked(
                type(ERC1967Proxy).creationCode,
                abi.encode(
                    address(accountImplementation),
                    abi.encodeCall(PermissionedAccount.initialize, (owner))
                )
            )));
    }

    modifier onlyAccountManager() {
        require(msg.sender == address(_accountManager), "only account manager");
        _;
    }
}
