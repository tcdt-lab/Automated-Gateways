package mediator

import dataType "relay/src/data_types"

type IopMediator struct {
}

func (iopMediator *IopMediator) GetPermittedNetwork(address string, networkType string) ([]*dataType.PermittedNetworkInfo, error) {

	var mediotrFactory MediatorFactory
	permittedNetworkMediator, err := mediotrFactory.CreatePermittedNetworkMediator(networkType)
	if err != nil {
		return nil, err
	}
	permittedNetwork, getAddErr := permittedNetworkMediator.GetAllPermittedNetworksByAddress(address)
	if getAddErr != nil {
		return nil, getAddErr
	}
	return permittedNetwork, nil
}
