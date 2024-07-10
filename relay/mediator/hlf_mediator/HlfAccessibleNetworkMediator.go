package hlf_mediator

import (
	dataTypes "github.com/tcdt-lab/Automated-Gateways/relay/data_types"
	acessibleNetworkHandler "github.com/tcdt-lab/Automated-Gateways/relay/internal/onchain/hlf_handler/accessible_network_handler"
)

type HlfAccessibleNetworkMediator struct {
}

var hlfAccessibleNetwork = acessibleNetworkHandler.HlfAccessibleNetwork{}

func (hlfAccessibleNetworkMediator *HlfAccessibleNetworkMediator) CreateAccessibleNetwork(networkName string, ip string, address string, companyName string) (*dataTypes.AccessibleNetworkInfo, error) {
	clientConnection, gateway, err := hlfAccessibleNetwork.OpenConnection()
	if err != nil {
		return nil, err
	}
	defer hlfAccessibleNetwork.CloseConnection(clientConnection, gateway)
	return hlfAccessibleNetwork.CreateAccessibleNetwork(gateway, networkName, ip, address, companyName)
}

func (hlfAccessibleNetworkMediator *HlfAccessibleNetworkMediator) RemoveAccessibleNetwork(id string) error {
	clientConnection, gateway, err := hlfAccessibleNetwork.OpenConnection()
	if err != nil {
		return err
	}
	defer hlfAccessibleNetwork.CloseConnection(clientConnection, gateway)
	return hlfAccessibleNetwork.RemoveAccessibleNetwork(gateway, id)
}
func (hlfAccessibleNetworkMediator *HlfAccessibleNetworkMediator) UpdateAccessibleNetwork(id string, networkName string, ip string, address string, companyName string) error {
	clientConnection, gateway, err := hlfAccessibleNetwork.OpenConnection()
	if err != nil {
		return err
	}
	defer hlfAccessibleNetwork.CloseConnection(clientConnection, gateway)
	return hlfAccessibleNetwork.UpdateAccessibleNetwork(gateway, id, networkName, ip, address, companyName)
}
func (hlfAccessibleNetworkMediator *HlfAccessibleNetworkMediator) GetAccessibleNetwork(id string) (*dataTypes.AccessibleNetworkInfo, error) {
	clientConnection, gateway, err := hlfAccessibleNetwork.OpenConnection()
	if err != nil {
		return nil, err
	}
	defer hlfAccessibleNetwork.CloseConnection(clientConnection, gateway)
	return hlfAccessibleNetwork.GetAccessibleNetwork(gateway, id)

}
func (hlfAccessibleNetworkMediator *HlfAccessibleNetworkMediator) GetAllAccessibleNetworksByAddress(address string) ([]*dataTypes.AccessibleNetworkInfo, error) {
	clientConnection, gateway, err := hlfAccessibleNetwork.OpenConnection()
	if err != nil {
		return nil, err
	}
	defer hlfAccessibleNetwork.CloseConnection(clientConnection, gateway)
	return hlfAccessibleNetwork.GetAllAccessibleNetworksByAddress(gateway, address)

}
func (hlfAccessibleNetworkMediator *HlfAccessibleNetworkMediator) AccessibleNetworkExists(id string) (bool, error) {
	clientConnection, gateway, err := hlfAccessibleNetwork.OpenConnection()
	if err != nil {
		return false, err
	}
	defer hlfAccessibleNetwork.CloseConnection(clientConnection, gateway)
	return hlfAccessibleNetwork.AccessibleNetworkExists(gateway, id)
}
