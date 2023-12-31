package mediator

import (
	dataType "github.com/tcdt-lab/Automated-Gateways/relay/data_types"
	offchain "github.com/tcdt-lab/Automated-Gateways/relay/internal/offchain/client"
)

type IopMediator struct {
}

// gRPC server Talks to this part
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

func (iopMediator *IopMediator) ReturnPermittedMethodList(networkId string, networkType string) ([]*dataType.MethodInfo, error) {
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

// gRPC client talks to this part
func (iopMediator *IopMediator) GetSelfInformationOnOutsideNetworks(outsiderNetworkId string, address string) ([]*dataType.AccessibleNetworkInfo, error) {
	res, err := offchain.GetNetworkInformation(outsiderNetworkId, address)
	if err != nil {
		return nil, err
	}
	return res, nil
}

func (iopMediator *IopMediator) InvokeAccessibleMethod(outsiderNetworkId string, name string, chaincode string, channel string, inputArgs string, output string) (string, string) {
	result, err := offchain.InvokeAccessibleMethod(outsiderNetworkId, name, chaincode, channel, inputArgs, output)
	return result, err
}

func (iopMediator *IopMediator) GetAccessibleMethodsList(outsiderNetworkId string, networkId string) ([]*dataType.MethodInfo, error) {
	res, err := offchain.GetAccessibleMethodsList(outsiderNetworkId, networkId)
	if err != nil {
		return nil, err
	}
	return res, nil
}
