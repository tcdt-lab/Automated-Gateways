// scripts/index.js
async function main () {
    // Our code will go here
    const address = '0xC8c03647d39a96f02f6Ce8999bc22493C290e734';
    const Box = await ethers.getContractFactory('PermittedNetworkManagementContract');
    const box = await Box.attach(address);
    const value = await box.getAllPermittedNetworks();
    console.log(value);
  }
  
  main()
    .then(() => process.exit(0))
    .catch(error => {
      console.error(error);
      process.exit(1);
    });