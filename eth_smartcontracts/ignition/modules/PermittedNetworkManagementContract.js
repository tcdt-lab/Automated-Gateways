const { buildModule } = require("@nomicfoundation/hardhat-ignition/modules");

const PermittedNetworkManagementContractModule = buildModule("PermittedNetworkManagementContractsModule", (m) => {
  const permittedNet = m.contract("PermittedNetworkManagementContract");

  return { permittedNet };
});

module.exports = PermittedNetworkManagementContractModule;