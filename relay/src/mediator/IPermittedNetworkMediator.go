package mediator

import permittedHandler "relay/src/onchain/permitted_network_handler"

type IPermittedNetworkMediator interface {
	CreatePermittedNetwork(networkName string, ip string, address string, companyName string) (*permittedHandler.PermittedNetworkInfo, error)
	RemovePermittedNetwork(id string) error
	UpdatePermittedNetwork(id string, networkName string, ip string, address string, companyName string) error
	GetPermittedNetwork(id string) (*permittedHandler.PermittedNetworkInfo, error)
	GetPermittedNetworkByIndexAndAddress(startIndex string, endIndex string, address string) ([]*permittedHandler.PermittedNetworkInfo, error)
	GetAllPermittedNetworksByAddress(address string) ([]*permittedHandler.PermittedNetworkInfo, error)
	PermittedNetworkExists(id string) (bool, error)
}
