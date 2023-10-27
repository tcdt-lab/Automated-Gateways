package hlf_mediator

import (
	dataTypes "github.com/tcdt-lab/Automated-Gateways/tree/main/relay/src/data_types"
	permittedHandler "github.com/tcdt-lab/Automated-Gateways/tree/main/relay/src/onchain/permitted_method_handler"
)

type HlfPermittedMethodsMediator struct {
}

func (hlfPermittedMethodsMediator *HlfPermittedMethodsMediator) AddPermittedMethod(permittedNetworkId string, permittedNetworkName string, chaincode string, channel string, inputArgs string, outputArgs string) (*dataTypes.MethodInfo, error) {
	var permittedMethods permittedHandler.HlfPermittedMethod
	connection, gateway, err := permittedMethods.OpenConnection()
	if err != nil {
		return nil, err
	}
	defer permittedMethods.CloseConnection(connection, gateway)
	return permittedMethods.AddPermittedMethod(gateway, permittedNetworkId, permittedNetworkName, chaincode, channel, inputArgs, outputArgs)

}

func (hlfPermittedMethodsMediator *HlfPermittedMethodsMediator) RemovePermittedMethod(id string) error {
	var permittedMethods permittedHandler.HlfPermittedMethod
	connection, gateway, err := permittedMethods.OpenConnection()
	if err != nil {
		return err
	}
	defer permittedMethods.CloseConnection(connection, gateway)
	return permittedMethods.RemovePermittedMethod(gateway, id)
}
func (hlfPermittedMethodsMediator *HlfPermittedMethodsMediator) UpdatePermittedMethod(id string, permittedNetworkName string, chaincode string, channel string, inputArgs string, outputArgs string) error {
	var permittedMethods permittedHandler.HlfPermittedMethod
	connection, gateway, err := permittedMethods.OpenConnection()
	if err != nil {
		return err
	}
	defer permittedMethods.CloseConnection(connection, gateway)
	return permittedMethods.UpdatePermittedMethod(gateway, id, permittedNetworkName, chaincode, channel, inputArgs, outputArgs)

}

func (hlfPermittedMethodsMediator *HlfPermittedMethodsMediator) GetPermittedMethodsByNetworkId(permittedNetworkId string) ([]*dataTypes.MethodInfo, error) {
	var permittedMethods permittedHandler.HlfPermittedMethod
	connection, gateway, err := permittedMethods.OpenConnection()
	if err != nil {
		return nil, err
	}
	defer permittedMethods.CloseConnection(connection, gateway)
	return permittedMethods.GetPermittedMethodsByNetworkId(gateway, permittedNetworkId)
}

func (hlfPermittedMethodsMediator *HlfPermittedMethodsMediator) GetPermittedMethod(id string) (*dataTypes.MethodInfo, error) {
	var permittedMethods permittedHandler.HlfPermittedMethod
	connection, gateway, err := permittedMethods.OpenConnection()
	if err != nil {
		return nil, err
	}
	defer permittedMethods.CloseConnection(connection, gateway)
	return permittedMethods.GetPermittedMethod(gateway, id)
}

func (hlfPermittedMethodsMediator *HlfPermittedMethodsMediator) GetPermittedMethodsByIndexAndAddress(startIndex string, endIndex string, address string) ([]*dataTypes.MethodInfo, error) {
	var permittedMethods permittedHandler.HlfPermittedMethod
	connection, gateway, err := permittedMethods.OpenConnection()
	if err != nil {
		return nil, err
	}
	defer permittedMethods.CloseConnection(connection, gateway)
	return permittedMethods.GetPermittedMethodsByIndex(gateway, startIndex, endIndex, address)
}

func (hlfPermittedMethodsMediator *HlfPermittedMethodsMediator) InvokePermittedMethod(name string, chaincode string, channel string, inputArgs []string) (*string, error) {
	var permittedMethods permittedHandler.HlfPermittedMethod
	connection, gateway, err := permittedMethods.OpenConnection()
	if err != nil {
		return nil, err
	}
	defer permittedMethods.CloseConnection(connection, gateway)
	return permittedMethods.InvokePermittedMethod(gateway, name, chaincode, channel, inputArgs)
}

func (hlfPermittedMethodsMediator *HlfPermittedMethodsMediator) PermittedMethodExists(id string) (bool, error) {
	var permittedMethods permittedHandler.HlfPermittedMethod
	connection, gateway, err := permittedMethods.OpenConnection()
	if err != nil {
		return false, err
	}
	defer permittedMethods.CloseConnection(connection, gateway)
	return permittedMethods.PermittedMethodExists(gateway, id)
}
