/** @type import('hardhat/config').HardhatUserConfig */

require("@nomicfoundation/hardhat-toolbox");


// Ensure your configuration variables are set before executing the script
const { vars } = require("hardhat/config");



//Besu account
const BESU_PRIVATE_KEY = vars.get("BESU_ACCOUNT_1");

module.exports = {
  solidity: "0.8.19",
  networks: {
    besu: {
      url: "http://localhost:8545/",
      accounts: [BESU_PRIVATE_KEY],
    },
  },
  paths: {
    sources: "./contracts",
    tests: "./test",
    cache: "./cache",
    artifacts: "./artifacts"
  },
};

