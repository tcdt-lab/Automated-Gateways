// SPDX-License-Identifier: UNLICENSED
pragma solidity ^0.8.19;

contract PermittedMethodManagementContract {
   struct PermittedMethod{
        uint permittedMethodId;
        string name;
        string chaincode;
        string inputArgs;
        string outputArgs;
    }


    mapping(uint => PermittedMethod) public permittedMethods;
    uint256[] public permittedMethodsIndexs;
    uint16 public  existingPermittedMethods = 0;

    function generatePermittedMethodId() public view returns (uint) {
        return permittedMethodsIndexs.length +1;
    }

    function createPermittedMethod(string calldata name,string calldata chaincode,string calldata inputArgs,string calldata outputArgs ) public returns (PermittedMethod memory) {

        uint permittedMethodId = generatePermittedMethodId();
        PermittedMethod memory permittedMethod = PermittedMethod(permittedMethodId,name,chaincode,inputArgs,outputArgs) ;
        permittedMethods[permittedMethodId] = permittedMethod;
        permittedMethodsIndexs.push(permittedMethodId);
        existingPermittedMethods = existingPermittedMethods + 1;
        return permittedMethod;
    }

    function getPermittedMethodsNumber() public view returns (uint) {
        return existingPermittedMethods;
    }

    function getPermittedMethod(uint permittedMethodId) public view returns (PermittedMethod memory) {
        return permittedMethods[permittedMethodId];
    }

    function getAllPermittedMethods() public view returns (PermittedMethod[] memory) {
        PermittedMethod[] memory permittedMethodsArray = new PermittedMethod[](existingPermittedMethods);
        uint16 counter = 0;
        for (uint i = 0; i < permittedMethodsIndexs.length; i++) {

            
            uint permittedMethodId = permittedMethodsIndexs[i];
            if (permittedMethodId == 0) {
                continue;
            }
            permittedMethodsArray[counter] = permittedMethods[permittedMethodId];
            counter = counter + 1;
            
        }
        return permittedMethodsArray;

    }

    function getPermittedMethodById(uint id) public view returns (PermittedMethod memory) {
        return permittedMethods[id];
    }

    function updatePermittedMethod(uint permittedMethodId, string calldata name,string calldata chaincode,string calldata inputArgs,string calldata outputArgs ) public returns (PermittedMethod memory) {
        PermittedMethod storage permittedMethod = permittedMethods[permittedMethodId];
        permittedMethod.outputArgs = outputArgs;
        permittedMethod.inputArgs = inputArgs;
        permittedMethod.chaincode = chaincode;
        permittedMethod.name = name;
        return permittedMethod;
    }


    function deletePermittedMethod(uint permittedMethodId) public {
        delete permittedMethods[permittedMethodId];
        for (uint i = 0; i < permittedMethodsIndexs.length; i++) {
            if (permittedMethodsIndexs[i] == permittedMethodId) {
                delete permittedMethodsIndexs[i];
            }
        }
        existingPermittedMethods = existingPermittedMethods - 1;
    }

     function isPermittedMethodExist(uint accessibleNetworkId) public view returns (bool) {
        if (permittedMethods[accessibleNetworkId].permittedMethodId == 0) {
            return false;
        }
        return true;
    }

}