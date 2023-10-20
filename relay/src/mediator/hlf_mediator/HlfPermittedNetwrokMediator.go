package hlf_mediator

import permittedHandler "relay/src/onchain/permitted_network_handler"

type HlfPermittedNetworkMediator struct {
}

func (hlfPermittedNetworkMediator *HlfPermittedNetworkMediator) CreatePermittedNetwork(networkName string, ip string, address string, companyName string) (*permittedHandler.PermittedNetworkInfo, error) {
	var permittedNetwork permittedHandler.HlfPermittedNetwork
	connection, gateway, err := permittedNetwork.OpenConnection()
	if err != nil {
		return nil, err
	}
	defer permittedNetwork.CloseConnection(connection, gateway)
	return permittedNetwork.CreatePermittedNetwork(gateway, networkName, ip, address, companyName)

}
func (hlfPermittedNetworkMediator *HlfPermittedNetworkMediator) RemovePermittedNetwork(id string) error {
	var permittedNetwork permittedHandler.HlfPermittedNetwork
	connection, gateway, err := permittedNetwork.OpenConnection()
	if err != nil {
		return err
	}
	defer permittedNetwork.CloseConnection(connection, gateway)
	return permittedNetwork.RemovePermittedNetwork(gateway, id)
}
func (hlfPermittedNetworkMediator *HlfPermittedNetworkMediator) UpdatePermittedNetwork(id string, networkName string, ip string, address string, companyName string) error {
	var permittedNetwork permittedHandler.HlfPermittedNetwork
	connection, gateway, err := permittedNetwork.OpenConnection()
	if err != nil {
		return err
	}
	defer permittedNetwork.CloseConnection(connection, gateway)
	return permittedNetwork.UpdatePermittedNetwork(gateway, id, networkName, ip, address, companyName)
}
func (hlfPermittedNetworkMediator *HlfPermittedNetworkMediator) GetPermittedNetwork(id string) (*permittedHandler.PermittedNetworkInfo, error) {
	var permittedNetwork permittedHandler.HlfPermittedNetwork
	connection, gateway, err := permittedNetwork.OpenConnection()
	if err != nil {
		return nil, err
	}
	defer permittedNetwork.CloseConnection(connection, gateway)
	return permittedNetwork.GetPermittedNetwork(gateway, id)
}
func (hlfPermittedNetworkMediator *HlfPermittedNetworkMediator) GetPermittedNetworkByIndexAndAddress(startIndex string, endIndex string, address string) ([]*permittedHandler.PermittedNetworkInfo, error) {
	var permittedNetwork permittedHandler.HlfPermittedNetwork
	connection, gateway, err := permittedNetwork.OpenConnection()
	if err != nil {
		return nil, err
	}
	defer permittedNetwork.CloseConnection(connection, gateway)
	return permittedNetwork.GetPermittedNetworkByIndexAndAddress(gateway, startIndex, endIndex, address)
}
func (hlfPermittedNetworkMediator *HlfPermittedNetworkMediator) GetAllPermittedNetworksByAddress(address string) ([]*permittedHandler.PermittedNetworkInfo, error) {
	var permittedNetwork permittedHandler.HlfPermittedNetwork
	connection, gateway, err := permittedNetwork.OpenConnection()
	if err != nil {
		return nil, err
	}
	defer permittedNetwork.CloseConnection(connection, gateway)
	return permittedNetwork.GetPermittedNetworksByAddress(gateway, address)
}
func (hlfPermittedNetworkMediator *HlfPermittedNetworkMediator) PermittedNetworkExists(id string) (bool, error) {
	var permittedNetwork permittedHandler.HlfPermittedNetwork
	connection, gateway, err := permittedNetwork.OpenConnection()
	if err != nil {
		return false, err
	}
	defer permittedNetwork.CloseConnection(connection, gateway)
	return permittedNetwork.PermittedNetworkExists(gateway, id)
}