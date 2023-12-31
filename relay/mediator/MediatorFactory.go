package mediator

import (
	"github.com/tcdt-lab/Automated-Gateways/relay/mediator/hlf_mediator"
	interfaces "github.com/tcdt-lab/Automated-Gateways/relay/mediator/interfaces"
)

const (
	HYPERLEDGER_FABRIC_NETWROK_TYPE = "HyperledgerFabricNetwork"
)

type MediatorFactory struct {
}

func (mediatorFactory *MediatorFactory) CreatePermittedNetworkMediator(networkType string) (interfaces.IPermittedNetworkMediator, error) {
	if networkType == HYPERLEDGER_FABRIC_NETWROK_TYPE {
		return &hlf_mediator.HlfPermittedNetworkMediator{}, nil
	}
	return nil, nil
}

func (mediatorFactory *MediatorFactory) CreateAccessibleNetworkMediator(networkType string) (interfaces.IAccessibleNetworkMediator, error) {
	if networkType == HYPERLEDGER_FABRIC_NETWROK_TYPE {
		return &hlf_mediator.HlfAccessibleNetworkMediator{}, nil
	}
	return nil, nil
}

func (mediatorFactory *MediatorFactory) CreatePermittedMethodsMediator(networkType string) (interfaces.IPermittedMethodMediator, error) {
	if networkType == HYPERLEDGER_FABRIC_NETWROK_TYPE {
		return &hlf_mediator.HlfPermittedMethodsMediator{}, nil
	}
	return nil, nil
}

func (mediatorFactory *MediatorFactory) CreateIopHistoryMediator(networkType string) (interfaces.IIopHistoryMediator, error) {
	if networkType == HYPERLEDGER_FABRIC_NETWROK_TYPE {
		return &hlf_mediator.HlfIopHistoryMediator{}, nil
	}
	return nil, nil
}
