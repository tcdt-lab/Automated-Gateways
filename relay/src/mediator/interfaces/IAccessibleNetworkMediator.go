package interfaces

import (
	dataTypes "github.com/tcdt-lab/Automated-Gateways/relay/src/data_types"
)

type IAccessibleNetworkMediator interface {
	CreateAccessibleNetwork(networkName string, ip string, address string, companyName string) (*dataTypes.AccessibleNetworkInfo, error)
	RemoveAccessibleNetwork(id string) error
	UpdateAccessibleNetwork(id string, networkName string, ip string, address string, companyName string) error
	GetAccessibleNetwork(id string) (*dataTypes.AccessibleNetworkInfo, error)
	GetAllAccessibleNetworksByAddress(address string) ([]*dataTypes.AccessibleNetworkInfo, error)
	AccessibleNetworkExists(id string) (bool, error)
}
