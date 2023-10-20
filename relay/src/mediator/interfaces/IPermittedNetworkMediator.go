package interfaces

import dataType "relay/src/data_types"

type IPermittedNetworkMediator interface {
	CreatePermittedNetwork(networkName string, ip string, address string, companyName string) (*dataType.PermittedNetworkInfo, error)
	RemovePermittedNetwork(id string) error
	UpdatePermittedNetwork(id string, networkName string, ip string, address string, companyName string) error
	GetPermittedNetwork(id string) (*dataType.PermittedNetworkInfo, error)
	GetPermittedNetworkByIndexAndAddress(startIndex string, endIndex string, address string) ([]*dataType.PermittedNetworkInfo, error)
	GetAllPermittedNetworksByAddress(address string) ([]*dataType.PermittedNetworkInfo, error)
	PermittedNetworkExists(id string) (bool, error)
}
