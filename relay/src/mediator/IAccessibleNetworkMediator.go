package mediator

import (
	acessibleHandler "relay/src/onchain/accessible_netwrok_handler"
)

type IAccessibleNetworkMediator interface {
	CreateAccessibleNetwork(networkName string, ip string, address string, companyName string) (*acessibleHandler.AccessibleNetworkInfo, error)
	RemoveAccessibleNetwork(id string) error
	UpdateAccessibleNetwork(id string, networkName string, ip string, address string, companyName string) error
	GetAccessibleNetwork(id string) (*acessibleHandler.AccessibleNetworkInfo, error)
	GetAllAccessibleNetworksByAddress(address string) ([]*acessibleHandler.AccessibleNetworkInfo, error)
	AccessibleNetworkExists(id string) (bool, error)
}
