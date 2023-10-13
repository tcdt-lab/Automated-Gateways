package permitted_network_handler

import (
	"log"
	"testing"
)

func TestOpenCloseConnetion(t *testing.T) {
	var hlfPermittedNetwork HlfPermittedNetwork
	clientConn, gateway, err := hlfPermittedNetwork.OpenConnection()
	if err != nil {
		t.Errorf("Error opening connection: %v", err)
	}

	errConnetion := hlfPermittedNetwork.CloseConnection(clientConn, gateway)
	if errConnetion != nil {
		t.Errorf("Error closing connection: %v", errConnetion)
	}

}

func TestCreatePermittedNetwork(t *testing.T) {
	var hlfPermittedNetwork HlfPermittedNetwork
	clientConn, gateway, err := hlfPermittedNetwork.OpenConnection()
	if err != nil {
		t.Errorf("Error opening connection: %v", err)
	}
	defer hlfPermittedNetwork.CloseConnection(clientConn, gateway)

	submitRes, err := hlfPermittedNetwork.CreatePermittedNetwork(gateway, "TestPermittedNet1", "127.0.0.1", "localhost", "KooshaComp")
	if err != nil {
		t.Errorf("Creating Accessible Netwrok: %v", err)
	}
	println(submitRes.PermittedNetworkId)
}

func TestPermittedNetworkExists(t *testing.T) {
	var hlfPermittedNetwork HlfPermittedNetwork
	clientConn, gateway, err := hlfPermittedNetwork.OpenConnection()
	if err != nil {
		t.Errorf("Error opening connection: %v", err)
	}
	defer hlfPermittedNetwork.CloseConnection(clientConn, gateway)

	res, err := hlfPermittedNetwork.PermittedNetworkExists(gateway, "PermittedNetwork_localhost_7")
	if err != nil {
		t.Errorf(" Accessible Netwrok Exists : %v", err)
	}
	if res != true {
		t.Errorf(" Accessible Netwrok Exists : %v", res)
	}
	wrongRes, err := hlfPermittedNetwork.PermittedNetworkExists(gateway, "PermittedNetwork_localhost_1999")
	if wrongRes != false {
		t.Errorf(" Accessible Netwrok Exists does not work properly : %v", wrongRes)
	}

}

func TestRemovePermittedNetwork(t *testing.T) {
	var hlfPermittedNetwork HlfPermittedNetwork
	clientConn, gateway, err := hlfPermittedNetwork.OpenConnection()
	if err != nil {
		t.Errorf("Error opening connection: %v", err)
	}
	defer hlfPermittedNetwork.CloseConnection(clientConn, gateway)

	err = hlfPermittedNetwork.RemovePermittedNetwork(gateway, "PermittedNetwork_localhost_6")
	if err != nil {
		t.Errorf(" Accessible Netwrok Exists : %v", err)
	}

}

func TestGetPermittedNetwork(t *testing.T) {
	var hlfPermittedNetwork HlfPermittedNetwork
	clientConn, gateway, err := hlfPermittedNetwork.OpenConnection()
	if err != nil {
		t.Errorf("Error opening connection: %v", err)
	}
	defer hlfPermittedNetwork.CloseConnection(clientConn, gateway)

	res, err := hlfPermittedNetwork.GetPermittedNetwork(gateway, "PermittedNetwork_localhost_7")
	if err != nil {
		t.Errorf(" Error in Getting Permitted Netwrok : %v", err)
	}
	log.Println("Permitted Network:", res.PermittedNetworkId)

	res, err = hlfPermittedNetwork.GetPermittedNetwork(gateway, "PermittedNetwork_localhost_3777")
	if err == nil {
		t.Errorf(" Problem in Error detection : %v", err)
	}

}

func TestGetPermittedNetworksByAddress(t *testing.T) {
	var hlfPermittedNetwork HlfPermittedNetwork
	clientConn, gateway, err := hlfPermittedNetwork.OpenConnection()
	if err != nil {
		t.Errorf("Error opening connection: %v", err)
	}
	defer hlfPermittedNetwork.CloseConnection(clientConn, gateway)

	res, err := hlfPermittedNetwork.GetPermittedNetworksByAddress(gateway, "localhost")
	if err != nil {
		t.Errorf(" Error in Getting All permitted Netwroks : %v", err)
	}
	log.Println("All Permitted Networks:", res[0].PermittedNetworkId)

}

func TestUpdatePermittedNetwork(t *testing.T) {
	var hlfPermittedNetwork HlfPermittedNetwork
	clientConn, gateway, err := hlfPermittedNetwork.OpenConnection()
	if err != nil {
		t.Errorf("Error opening connection: %v", err)
	}
	defer hlfPermittedNetwork.CloseConnection(clientConn, gateway)

	err = hlfPermittedNetwork.UpdatePermittedNetwork(gateway, "PermittedNetwork_localhost_7", "TestPermittedNetupdated", "127.0.0.1", "localhost", "KooshaComp")
	if err != nil {
		t.Errorf(" Error in Updating Permitted Netwrok : %v", err)
	}
}
