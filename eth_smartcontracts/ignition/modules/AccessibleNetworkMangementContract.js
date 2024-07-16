const { buildModule } = require("@nomicfoundation/hardhat-ignition/modules");

const AccessibleNetworkManagementContractModule = buildModule("AccessibleNetworkManagementContractsModule", (m) => {
  const token = m.contract("AccessibleNetworkManagementContract");

  return { token };
});

module.exports = AccessibleNetworkManagementContractModule;