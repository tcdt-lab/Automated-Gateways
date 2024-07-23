const { expect } = require("chai");
const { ethers } = require("hardhat");

describe("permitted_method_management contract", function () {
    it("Deployment should create a new permitted method", async function () {
        const [owner] = await ethers.getSigners();
    
     
        const TestContract = await ethers.getContractFactory("PermittedMethodManagementContract");
        
        const contract = await TestContract.deploy();
        const permittedNet = await contract.createPermittedMethod("methodName1","chaincodeName1", "inputArgs1","outputArgs1" );
        const permittedNet2 = await contract.createPermittedMethod("methodName2","chaincodeName2", "inputArgs2","outputArgs2" );
        const permittedNet3 = await contract.createPermittedMethod("methodName3","chaincodeName3", "inputArgs3","outputArgs3" );


        const numberOfPermits = await contract.getPermittedMethodsNumber();
        expect(numberOfPermits).to.equal(3);
    });

    it("Should return the method details (all methods/based on Id)", async function () {
        const [owner] = await ethers.getSigners();
    
     
        const TestContract = await ethers.getContractFactory("PermittedMethodManagementContract");
        
        const contract = await TestContract.deploy();
        const permittedNet = await contract.createPermittedMethod("methodName1","chaincodeName1", "inputArgs1","outputArgs1" );
        const permittedNet2 = await contract.createPermittedMethod("methodName2","chaincodeName2", "inputArgs2","outputArgs2" );
        const permittedNet3 = await contract.createPermittedMethod("methodName3","chaincodeName3", "inputArgs3","outputArgs3" );

        const permittedMethod = await contract.getPermittedMethod(1);
        expect(permittedMethod.name).to.equal("methodName1");

        const permittedMethods = await contract.getAllPermittedMethods();
        expect(permittedMethods.length).to.equal(3);
        expect(permittedMethods[0].name).to.equal("methodName1");
        expect(permittedMethods[1].name).to.equal("methodName2");
        expect(permittedMethods[2].name).to.equal("methodName3");
    });

    it("Should update the method details", async function () {
        const [owner] = await ethers.getSigners();
    
     
        const TestContract = await ethers.getContractFactory("PermittedMethodManagementContract");
        
        const contract = await TestContract.deploy();
        const permittedNet = await contract.createPermittedMethod("methodName1","chaincodeName1", "inputArgs1","outputArgs1" );
        const permittedNet2 = await contract.createPermittedMethod("methodName2","chaincodeName2", "inputArgs2","outputArgs2" );
        const permittedNet3 = await contract.createPermittedMethod("methodName3","chaincodeName3", "inputArgs3","outputArgs3" );

        const updatedPermittedNet = await contract.updatePermittedMethod(1,"methodName2_updated","chaincodeName2_updated","inputArgs2_updated","outputArgs2_update");
        const updatedPermittedMethod = await contract.getPermittedMethod(1);
        expect(updatedPermittedMethod.name).to.equal("methodName2_updated");
    });

    it("Should delete the method", async function () {
        const [owner] = await ethers.getSigners();
    
     
        const TestContract = await ethers.getContractFactory("PermittedMethodManagementContract");
        
        const contract = await TestContract.deploy();
        const permittedNet = await contract.createPermittedMethod("methodName1","chaincodeName1", "inputArgs1","outputArgs1" );
        const permittedNet2 = await contract.createPermittedMethod("methodName2","chaincodeName2", "inputArgs2","outputArgs2" );
        const permittedNet3 = await contract.createPermittedMethod("methodName3","chaincodeName3", "inputArgs3","outputArgs3" );

        const deletedPermittedNet = await contract.deletePermittedMethod(1);
        const numberOfPermits = await contract.getPermittedMethodsNumber();
        expect(numberOfPermits).to.equal(2);

        const permittedMethods = await contract.getAllPermittedMethods();
        expect(permittedMethods.length).to.equal(2);
    });

    it("should test the existance", async function () {
        const [owner] = await ethers.getSigners();
    
     
        const TestContract = await ethers.getContractFactory("PermittedMethodManagementContract");
        
        const contract = await TestContract.deploy();
        const permittedNet = await contract.createPermittedMethod("methodName1","chaincodeName1", "inputArgs1","outputArgs1" );
        const permittedNet2 = await contract.createPermittedMethod("methodName2","chaincodeName2", "inputArgs2","outputArgs2" );
        const permittedNet3 = await contract.createPermittedMethod("methodName3","chaincodeName3", "inputArgs3","outputArgs3" );

        const existance = await contract.isPermittedMethodExist(1);
        expect(existance).to.equal(true);
    });
});
