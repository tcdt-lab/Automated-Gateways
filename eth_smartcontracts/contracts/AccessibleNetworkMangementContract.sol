// SPDX-License-Identifier: UNLICENSED
pragma solidity ^0.8.19;

contract AccessibleNetworkManagementContract{
    struct AccessibleNetwork {
        string networkName;
        string networkIP;
        string networkAddress;
        string companyName;
        uint accessibleNetworkId;
    }

    mapping(uint => AccessibleNetwork) public accessibleNetworks;
    uint256[] public accessibleNetworkIndexs;
    uint16 public  existingAccessibleNetworks = 0;

    function generateAccessibleNetworkId() public view returns (uint) {
        return accessibleNetworkIndexs.length +1;
    }

    function createAccessibleNetwork(string calldata networkName, string calldata networkIP,string calldata networkAddress, string calldata companyName) public returns (AccessibleNetwork memory) {

        uint accessibleNetworkId = generateAccessibleNetworkId();
        AccessibleNetwork memory accessibleNetwork = AccessibleNetwork(networkName, networkIP, networkAddress, companyName,accessibleNetworkId) ;
        accessibleNetworks[accessibleNetworkId] = accessibleNetwork;
        accessibleNetworkIndexs.push(accessibleNetworkId);
        existingAccessibleNetworks = existingAccessibleNetworks + 1;
        return accessibleNetwork;
    }

    function getAccessibleNetworksNumber() public view returns (uint) {
        return existingAccessibleNetworks;
    }

    function getAccessibleNetwork(uint accessibleNetworkId) public view returns (AccessibleNetwork memory) {
        return accessibleNetworks[accessibleNetworkId];
    }

    function getAllAccessibleNetworks() public view returns (AccessibleNetwork[] memory) {
        AccessibleNetwork[] memory accessibleNetworksArray = new AccessibleNetwork[](existingAccessibleNetworks);
        uint16 counter = 0;
        for (uint i = 0; i < accessibleNetworkIndexs.length; i++) {

            
            uint accessibleNetworkId = accessibleNetworkIndexs[i];
            if (accessibleNetworkId == 0) {
                continue;
            }
            accessibleNetworksArray[counter] = accessibleNetworks[accessibleNetworkId];
            counter = counter + 1;
            
        }
        return accessibleNetworksArray;

    }

    function updateAccessibleNetwork(uint accessibleNetworkId, string calldata networkName, string calldata networkIP,string calldata networkAddress, string calldata companyName) public returns (AccessibleNetwork memory) {
        AccessibleNetwork storage accessibleNetwork = accessibleNetworks[accessibleNetworkId];
        accessibleNetwork.networkName = networkName;
        accessibleNetwork.networkIP = networkIP;
        accessibleNetwork.networkAddress = networkAddress;
        accessibleNetwork.companyName = companyName;
        return accessibleNetwork;
    }

    function deleteAccessibleNetwork(uint accessibleNetworkId) public {
        delete accessibleNetworks[accessibleNetworkId];
        for (uint i = 0; i < accessibleNetworkIndexs.length; i++) {
            if (accessibleNetworkIndexs[i] == accessibleNetworkId) {
                delete accessibleNetworkIndexs[i];
            }
        }
        existingAccessibleNetworks = existingAccessibleNetworks - 1;
    }

    function isAccessibleNetworkExist(uint accessibleNetworkId) public view returns (bool) {
        if (accessibleNetworks[accessibleNetworkId].accessibleNetworkId == 0) {
            return false;
        }
        return true;
    }
}