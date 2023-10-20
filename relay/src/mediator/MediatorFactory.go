package mediator

import hlfMediator "relay/src/mediator/hlf_mediator"

const (
	HYPERLEDGER_FABRIC_NETWROK_TYPE = "HyperledgerFabricNetwork"
)

type MediatorFactory struct {
}

func (mediatorFactory *MediatorFactory) CreatePermittedNetworkMediator(networkType string) (IPermittedNetworkMediator, error) {
	if networkType == HYPERLEDGER_FABRIC_NETWROK_TYPE {
		return &hlfMediator.HlfPermittedNetworkMediator{}, nil
	}
	return nil, nil
}

func (mediatorFactory *MediatorFactory) CreateAccessibleNetworkMediator(networkType string) (IAccessibleNetworkMediator, error) {
	if networkType == HYPERLEDGER_FABRIC_NETWROK_TYPE {
		return &hlfMediator.HlfAccessibleNetworkMediator{}, nil
	}
	return nil, nil
}

func (mediatorFactory *MediatorFactory) CreatePermittedMethodsMediator(networkType string) (IPermittedMethodMediator, error) {
	if networkType == HYPERLEDGER_FABRIC_NETWROK_TYPE {
		return &hlfMediator.HlfPermittedMethodsMediator{}, nil
	}
	return nil, nil
}
