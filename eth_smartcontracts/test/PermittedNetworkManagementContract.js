const { expect } = require("chai");
const { ethers } = require("hardhat");

describe("permitted_network_management contract", function () {
  it("Deployment should create a new permitted network", async function () {
    const [owner] = await ethers.getSigners();

 
    const TestContract = await ethers.getContractFactory("PermittedNetworkManagementContract");
    
    const contract = await TestContract.deploy();


    const permittedNet = await contract.createPermittedNetwork("networkName1","networkIP1","networkAddress1","networkCompany1");
    const permittedNet2 = await contract.createPermittedNetwork("networkName2","networkIP2","networkAddress2","networkCompany2");
    const permittedNet3 = await contract.createPermittedNetwork("networkName3","networkIP3","networkAddress3","networkCompany3");

   
    const numberOfPermits = await contract.getPermittedNetworksNumber();
   
    expect(numberOfPermits).to.equal(3);
  });
  it("Should return the network details (all netwroks/based on Id)", async function () {
    const [owner] = await ethers.getSigners();

 
    const TestContract = await ethers.getContractFactory("PermittedNetworkManagementContract");
    
    const contract = await TestContract.deploy();


    const permittedNet = await contract.createPermittedNetwork("networkName1","networkIP1","networkAddress1","networkCompany1");
    const permittedNet2 = await contract.createPermittedNetwork("networkName2","networkIP2","networkAddress2","networkCompany2");
    const permittedNet3 = await contract.createPermittedNetwork("networkName3","networkIP3","networkAddress3","networkCompany3");

    const permittedNetwork = await contract.getPermittedNetwork(1);
    console.log(permittedNetwork);
    expect(permittedNetwork.networkName).to.equal("networkName1");

    const permittedNetworks = await contract.getAllPermittedNetworks();
    console.log(permittedNetworks);
    expect(permittedNetworks.length).to.equal(3);
    expect(permittedNetworks[0].networkName).to.equal("networkName1");
  });

  it("Should update the network details", async function () {
    const [owner] = await ethers.getSigners();

 
    const TestContract = await ethers.getContractFactory("PermittedNetworkManagementContract");
    
    const contract = await TestContract.deploy();


    const permittedNet = await contract.createPermittedNetwork("networkName1","networkIP1","networkAddress1","networkCompany1");
    const permittedNet2 = await contract.createPermittedNetwork("networkName2","networkIP2","networkAddress2","networkCompany2");
    const permittedNet3 = await contract.createPermittedNetwork("networkName3","networkIP3","networkAddress3","networkCompany3");

    const updatedPermittedNet = await contract.updatePermittedNetwork(1,"networkName2_updated","networkIP2_updated","networkAddress2_updated","networkCompany2_update");
    const updatedPermittedNetwork = await contract.getPermittedNetwork(1);
    console.log(updatedPermittedNetwork);
    expect(updatedPermittedNetwork.networkName).to.equal("networkName2_updated");
  });

  it("Should delete the network details", async function () {
    const [owner] = await ethers.getSigners();

 
    const TestContract = await ethers.getContractFactory("PermittedNetworkManagementContract");
    
    const contract = await TestContract.deploy();


    const permittedNet = await contract.createPermittedNetwork("networkName1","networkIP1","networkAddress1","networkCompany1");
    const permittedNet2 = await contract.createPermittedNetwork("networkName2","networkIP2","networkAddress2","networkCompany2");
    const permittedNet3 = await contract.createPermittedNetwork("networkName3","networkIP3","networkAddress3","networkCompany3");
    const checkPermittedNetworkExists = await contract.permittedNetworkExists(1);
    expect(checkPermittedNetworkExists).to.equal(true);
    console.log("##############GET ARRAY INDEX BEFORE DELETE################");
    var  getPermittedNetworkIndexArray = await contract.getPermittedNetworkIndexArray();
    console.log(getPermittedNetworkIndexArray);
  

    const deletedPermittedNet = await contract.deletePermittedNetwork(1);
 console.log("#################GET ALL INDEXS AFTER DELETE#############");
     getPermittedNetworkIndexArray = await contract.getPermittedNetworkIndexArray();
    console.log(getPermittedNetworkIndexArray);
    console.log("##############GET ALL Permitted Nets AFTER DLETE################");
     allpermittedNets = await contract.getAllPermittedNetworks();
    console.log(allpermittedNets);
    console.log("##############################");

    console.log("##############################");
  
  });
});