package mediator

import dataType "relay/src/data_types"

type IopMediator struct {
}

func (iopMediator *IopMediator) ReturnPermittedNetworkInfo(address string, networkType string) ([]*dataType.PermittedNetworkInfo, error) {

	var mediatorFactory MediatorFactory
	permittedNetworkMediator, err := mediatorFactory.CreatePermittedNetworkMediator(networkType)
	if err != nil {
		return nil, err
	}
	permittedNetwork, getAddErr := permittedNetworkMediator.GetAllPermittedNetworksByAddress(address)
	if getAddErr != nil {
		return nil, getAddErr
	}
	return permittedNetwork, nil
}

func (iopMediator *IopMediator) ReturnPermittedMethodList(networkId string, networkType string) ([]*dataType.PermittedMethodInfo, error) {
	var mediatorFactory MediatorFactory
	permittedMethodMediator, err := mediatorFactory.CreatePermittedMethodsMediator(networkType)
	if err != nil {
		return nil, err
	}
	permittedMethod, getAddErr := permittedMethodMediator.GetPermittedMethodsByNetworkId(networkId)
	if getAddErr != nil {
		return nil, getAddErr
	}
	return permittedMethod, nil
}

func (iopMediator *IopMediator) InvokePermittedMethod(name string, chaincode string, channel string, inputArgs []string, networkType string) (*string, error) {
	var mediatorFactory MediatorFactory
	permittedMethodMediator, err := mediatorFactory.CreatePermittedMethodsMediator(networkType)
	if err != nil {
		return nil, err
	}
	result, getAddErr := permittedMethodMediator.InvokePermittedMethod(name, chaincode, channel, inputArgs)
	if getAddErr != nil {
		return nil, getAddErr
	}
	return result, nil
}
