const { buildModule } = require("@nomicfoundation/hardhat-ignition/modules");

const PermittedMethodManagementContractModule = buildModule("PermittedMethodManagementContractsModule", (m) => {
  const permittedMethod = m.contract("PermittedMethodManagementContract");

  return { permittedMethod };
});

module.exports = PermittedMethodManagementContractModule;