package mediator

import (
	hlfMediator "relay/src/mediator/hlf_mediator"
	"relay/src/mediator/interfaces"
)

const (
	HYPERLEDGER_FABRIC_NETWROK_TYPE = "HyperledgerFabricNetwork"
)

type MediatorFactory struct {
}

func (mediatorFactory *MediatorFactory) CreatePermittedNetworkMediator(networkType string) (interfaces.IPermittedNetworkMediator, error) {
	if networkType == HYPERLEDGER_FABRIC_NETWROK_TYPE {
		return &hlfMediator.HlfPermittedNetworkMediator{}, nil
	}
	return nil, nil
}

func (mediatorFactory *MediatorFactory) CreateAccessibleNetworkMediator(networkType string) (interfaces.IAccessibleNetworkMediator, error) {
	if networkType == HYPERLEDGER_FABRIC_NETWROK_TYPE {
		return &hlfMediator.HlfAccessibleNetworkMediator{}, nil
	}
	return nil, nil
}

func (mediatorFactory *MediatorFactory) CreatePermittedMethodsMediator(networkType string) (interfaces.IPermittedMethodMediator, error) {
	if networkType == HYPERLEDGER_FABRIC_NETWROK_TYPE {
		return &hlfMediator.HlfPermittedMethodsMediator{}, nil
	}
	return nil, nil
}
