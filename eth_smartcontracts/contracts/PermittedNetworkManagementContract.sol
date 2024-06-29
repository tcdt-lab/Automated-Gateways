// SPDX-License-Identifier: UNLICENSED
pragma solidity ^0.8.19;


contract PermittedNetworkManagementContract {
   
    struct PermittedNetwork {
        string networkName;
        string networkIP;
        string networkAddress;
        string companyName;
        uint permittedNetworkId;
        
    }
    mapping(uint => PermittedNetwork) public permittedNetworks;
    uint256[] public permittedNetworkIndexs;
    uint16 public  existingAccessibleNetworks = 0;

    function generatePermittedNetworkId() public view returns (uint) {
        return permittedNetworkIndexs.length +1;
    }

    function createPermittedNetwork(string calldata networkName, string calldata networkIP,string calldata networkAddress, string calldata companyName) public returns (PermittedNetwork memory) {

        uint permittedNetworkId = generatePermittedNetworkId();
        PermittedNetwork memory permittedNetwork = PermittedNetwork(networkName, networkIP, networkAddress, companyName,permittedNetworkId) ;
        permittedNetworks[permittedNetworkId] = permittedNetwork;
        permittedNetworkIndexs.push(permittedNetworkId);
        existingAccessibleNetworks = existingAccessibleNetworks + 1;
        return permittedNetwork;
    }

    function getPermittedNetworksNumber() public view returns (uint) {
        return permittedNetworkIndexs.length;
    }

    function getPermittedNetwork(uint permittedNetworkId) public view returns (PermittedNetwork memory) {
        return permittedNetworks[permittedNetworkId];
    }

    function getAllPermittedNetworks() public view returns (PermittedNetwork[] memory) {
        PermittedNetwork[] memory permittedNetworksArray = new PermittedNetwork[](permittedNetworkIndexs.length);
        for (uint i = 0; i < permittedNetworkIndexs.length; i++) {

            permittedNetworksArray[i] = permittedNetworks[permittedNetworkIndexs[i]];
        }
        return permittedNetworksArray;
    }

    function updatePermittedNetwork(uint permittedNetworkId, string calldata networkName, string calldata networkIP,string calldata networkAddress, string calldata companyName) public
     returns (PermittedNetwork memory) {
        PermittedNetwork storage permittedNetwork = permittedNetworks[permittedNetworkId];
        permittedNetwork.networkName = networkName;
        permittedNetwork.networkIP = networkIP;
        permittedNetwork.networkAddress =networkAddress;
        permittedNetwork.companyName = companyName;
        return permittedNetwork;
    }

    function deletePermittedNetwork(uint permittedNetworkId) public {
        delete permittedNetworks[permittedNetworkId];
        for (uint i = 0; i < permittedNetworkIndexs.length; i++) {
            if (permittedNetworkIndexs[i] == permittedNetworkId) {
                delete permittedNetworkIndexs[i];
            }
        }
    }

    function permittedNetworkExists(uint permittedNetworkId) public view returns (bool) {
        return permittedNetworks[permittedNetworkId].permittedNetworkId != 0;
    }

    function getPermittedNetworkIndexArray() public view returns (uint256[] memory) {
        return permittedNetworkIndexs;
    }

    function getAlltheKeysInMapping() public view returns (uint[] memory) {
        uint[] memory keys = new uint[](permittedNetworkIndexs.length);
        for (uint i = 0; i < permittedNetworkIndexs.length; i++) {
            keys[i] = permittedNetworks[permittedNetworkIndexs[i]].permittedNetworkId;
        }
        return keys;
    }
}