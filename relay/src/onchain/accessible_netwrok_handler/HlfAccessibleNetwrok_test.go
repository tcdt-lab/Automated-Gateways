package accessible_netwrok_handler

import (
	"fmt"
	"testing"
)

func TestOpenCloseConnetion(t *testing.T) {
	var hlfAccessibleNetwork HlfAccessibleNetwork
	clientConn, gateway, err := hlfAccessibleNetwork.OpenConnection()
	if err != nil {
		t.Errorf("Error opening connection: %v", err)
	}

	errConnetion := hlfAccessibleNetwork.CloseConnection(clientConn, gateway)
	if errConnetion != nil {
		t.Errorf("Error closing connection: %v", errConnetion)
	}

}

func TestCreateAccessibleNetwork(t *testing.T) {
	var hlfAccessibleNetwork HlfAccessibleNetwork
	clientConn, gateway, err := hlfAccessibleNetwork.OpenConnection()
	if err != nil {
		t.Errorf("Error opening connection: %v", err)
	}
	defer hlfAccessibleNetwork.CloseConnection(clientConn, gateway)

	submitRes, err := hlfAccessibleNetwork.CreateAccessibleNetwork(gateway, "TestAccessibleNet1", "127.0.0.1", "localhost", "KooshaComp")
	if err != nil {
		t.Errorf("Creating Accessible Netwrok: %v", err)
	}
	println(submitRes.AccessibleNetworkId)
}
func TestAccessibleNetworkExists(t *testing.T) {
	var hlfAccessibleNetwork HlfAccessibleNetwork
	clientConn, gateway, err := hlfAccessibleNetwork.OpenConnection()
	if err != nil {
		t.Errorf("Error opening connection: %v", err)
	}
	defer hlfAccessibleNetwork.CloseConnection(clientConn, gateway)

	res, err := hlfAccessibleNetwork.AccessibleNetworkExists(gateway, "AccessibleNetwork1")
	if err != nil {
		t.Errorf(" Accessible Netwrok Exists : %v", err)
	}
	if res != true {
		t.Errorf(" Accessible Netwrok Exists : %v", res)
	}
	wrongRes, err := hlfAccessibleNetwork.AccessibleNetworkExists(gateway, "AccessibleNetwork1999")
	if wrongRes != false {
		t.Errorf(" Accessible Netwrok Exists : %v", wrongRes)
	}

	fmt.Println("Test Res:", res)

}

func TestRemoveAccessibleNetwork(t *testing.T) {
	var hlfAccessibleNetwork HlfAccessibleNetwork
	clientConn, gateway, err := hlfAccessibleNetwork.OpenConnection()
	if err != nil {
		t.Errorf("Error opening connection: %v", err)
	}
	defer hlfAccessibleNetwork.CloseConnection(clientConn, gateway)

	err = hlfAccessibleNetwork.RemoveAccessibleNetwork(gateway, "AccessibleNetwork2")
	if err != nil {
		t.Errorf(" Accessible Netwrok Exists : %v", err)
	}

	err = hlfAccessibleNetwork.RemoveAccessibleNetwork(gateway, "AccessibleNetwork1000")
	if err == nil {
		fmt.Printf(" Accessible Netwrok Exists : %v", err)
	}

}

func TestUpdateAccessibleNetwork(t *testing.T) {
	var hlfAccessibleNetwork HlfAccessibleNetwork
	clientConn, gateway, err := hlfAccessibleNetwork.OpenConnection()
	if err != nil {
		t.Errorf("Error opening connection: %v", err)
	}
	defer hlfAccessibleNetwork.CloseConnection(clientConn, gateway)

	err = hlfAccessibleNetwork.UpdateAccessibleNetwork(gateway, "AccessibleNetwork3", "NetworkName3", "127.0.0.1", "localhost", "KooshaComp3")
	if err != nil {
		t.Errorf(" Accessible Netwrok Exists : %v", err)
	}
	err = hlfAccessibleNetwork.UpdateAccessibleNetwork(gateway, "AccessibleNetwork3000", "NetworkName3", "127.0.0.1", "localhost", "KooshaComp3")
	if err == nil {
		fmt.Printf(" Cannot Detect fake ID : %v", err)
	}
}

func TestGetAccessibleNetwork(t *testing.T) {
	var hlfAccessibleNetwork HlfAccessibleNetwork
	clientConn, gateway, err := hlfAccessibleNetwork.OpenConnection()
	if err != nil {
		t.Errorf("Error opening connection: %v", err)
	}
	defer hlfAccessibleNetwork.CloseConnection(clientConn, gateway)

	res, err := hlfAccessibleNetwork.GetAccessibleNetwork(gateway, "AccessibleNetwork3")
	if err != nil {
		t.Errorf("Error while Getting  Accessible Netwrok   : %v", err)
	}
	if res.AccessibleNetworkId != "AccessibleNetwork3" {
		t.Errorf("Error in getting data : %v", res)
	}
	wrongRes, err := hlfAccessibleNetwork.GetAccessibleNetwork(gateway, "AccessibleNetwork1999")
	if wrongRes != nil {
		t.Errorf(
			" Cannot get data based on Id : %v", wrongRes)
	}
	fmt.Println("Test Res:", res)
}

func TestQueryAllAccessibleNetworks(t *testing.T) {
	var hlfAccessibleNetwork HlfAccessibleNetwork
	clientConn, gateway, err := hlfAccessibleNetwork.OpenConnection()
	if err != nil {
		t.Errorf("Error opening connection: %v", err)
	}
	defer hlfAccessibleNetwork.CloseConnection(clientConn, gateway)

	res, err := hlfAccessibleNetwork.QueryAllAccessibleNetworks(gateway)
	if err != nil {
		t.Errorf("Error while Getting  Accessible Netwrok   : %v", err)
	}

	fmt.Println("Test Res:", res)
}
