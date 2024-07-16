const { expect } = require("chai");
const { ethers } = require("hardhat");

describe("accessible_network_management contract", function () {
  it("Deployment should create a new accessible network", async function () {
    const [owner] = await ethers.getSigners();
    const TestContract = await ethers.getContractFactory("AccessibleNetworkManagementContract");
    
    const contract = await TestContract.deploy();


    const permittedNet = await contract.createAccessibleNetwork("networkAccessibleName1","networkIP1","networkAddress1","networkCompany1");
    const permittedNet2 = await contract.createAccessibleNetwork("networkAccessibleName2","networkIP2","networkAddress2","networkCompany2");
    const permittedNet3 = await contract.createAccessibleNetwork("networkAccessibleName3","networkIP3","networkAddress3","networkCompany3");

   
    const numberOfPermits = await contract.getAccessibleNetworksNumber();
   
    expect(numberOfPermits).to.equal(3);
  });

  it("Should return the network details (all netwroks/based on Id)", async function () {
    const [owner] = await ethers.getSigners();
    const TestContract = await ethers.getContractFactory("AccessibleNetworkManagementContract");
    
    const contract = await TestContract.deploy();

    const AccessibleNet = await contract.createAccessibleNetwork("networkAccessibleName1","networkIP1","networkAddress1","networkCompany1");
    const AccessibleNet2 = await contract.createAccessibleNetwork("networkAccessibleName2","networkIP2","networkAddress2","networkCompany2");
    const AccessibleNet3 = await contract.createAccessibleNetwork("networkAccessibleName3","networkIP3","networkAddress3","networkCompany3");
    const AccessibleNet4 = await contract.createAccessibleNetwork("networkAccessibleName3","networkIP3","networkAddress3","networkCompany3");
    const AccessibleNet5 = await contract.createAccessibleNetwork("networkAccessibleName3","networkIP3","networkAddress3","networkCompany3");
    const AccessibleNet6 = await contract.createAccessibleNetwork("networkAccessibleName3","networkIP3","networkAddress3","networkCompany3");

    const network = await contract.getAccessibleNetwork(1);
    const network2 = await contract.getAccessibleNetwork(2);
    const network3 = await contract.getAccessibleNetwork(3);

    

    expect(network.networkName).to.equal("networkAccessibleName1");
    expect(network2.networkName).to.equal("networkAccessibleName2");
    expect(network3.networkName).to.equal("networkAccessibleName3");

    const networks = await contract.getAllAccessibleNetworks();
    console.log("************TEST GET ALL ************");
    console.log(networks);
    console.log("*************************************");
    expect(networks.length).to.equal(6);
    expect(networks[0].networkName).to.equal("networkAccessibleName1");

  });

  it("Should update the network details", async function () {
    const [owner] = await ethers.getSigners();
    const TestContract = await ethers.getContractFactory("AccessibleNetworkManagementContract");
    
    const contract = await TestContract.deploy();

    const AccessibleNet = await contract.createAccessibleNetwork("networkAccessibleName1","networkIP1","networkAddress1","networkCompany1");
    const AccessibleNet2 = await contract.createAccessibleNetwork("networkAccessibleName2","networkIP2","networkAddress2","networkCompany2");
    const AccessibleNet3 = await contract.createAccessibleNetwork("networkAccessibleName3","networkIP3","networkAddress3","networkCompany3");

    const updateAccessibleNet= await contract.updateAccessibleNetwork(1,"networkAccessibleName2_updated","networkIP2_updated","networkAddress2_updated","networkCompany2_update");
    const updatedAccessibleNet= await contract.getAccessibleNetwork(1);

    expect(updatedAccessibleNet.networkName).to.equal("networkAccessibleName2_updated");
  });
    
  it("Should delete the network details", async function () {
    const [owner] = await ethers.getSigners();
    const TestContract = await ethers.getContractFactory("AccessibleNetworkManagementContract");
    
    const contract = await TestContract.deploy();

    const AccessibleNet = await contract.createAccessibleNetwork("networkAccessibleName1","networkIP1","networkAddress1","networkCompany1");
   
    const AccessibleNet2 = await contract.createAccessibleNetwork("networkAccessibleName2","networkIP2","networkAddress2","networkCompany2");
    const AccessibleNet3 = await contract.createAccessibleNetwork("networkAccessibleName3","networkIP3","networkAddress3","networkCompany3");

    const deleteAccessibleNet= await contract.deleteAccessibleNetwork(1);
    const getAllAccessibleNetworks = await contract.getAllAccessibleNetworks();
    console.log(getAllAccessibleNetworks);

    const numberOfPermits = await contract.getAccessibleNetworksNumber();
    expect(numberOfPermits).to.equal(2);
  });
});