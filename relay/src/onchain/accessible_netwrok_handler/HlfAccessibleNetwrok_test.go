package accessible_netwrok_handler

import (
	"fmt"
	"testing"
)

func TestOpenCloseConnetion(t *testing.T) {
	clientConn, gateway, err := OpenConnection()
	if err != nil {
		t.Errorf("Error opening connection: %v", err)
	}

	errConnetion := CloseConnection(clientConn, gateway)
	if errConnetion != nil {
		t.Errorf("Error closing connection: %v", errConnetion)
	}

}

func TestCreateAccessibleNetwork(t *testing.T) {
	clientConn, gateway, err := OpenConnection()
	if err != nil {
		t.Errorf("Error opening connection: %v", err)
	}
	defer CloseConnection(clientConn, gateway)
	var hlfAccessibleNetwork HlfAccessibleNetwork
	submitRes, err := hlfAccessibleNetwork.CreateAccessibleNetwork(gateway, "TestAccessibleNet1", "127.0.0.1", "localhost", "KooshaComp")
	if err != nil {
		t.Errorf("Creating Accessible Netwrok: %v", err)
	}
	println(submitRes.AccessibleNetworkId)
}
func TestAccessibleNetworkExists(t *testing.T) {
	clientConn, gateway, err := OpenConnection()
	if err != nil {
		t.Errorf("Error opening connection: %v", err)
	}
	defer CloseConnection(clientConn, gateway)
	var hlfAccessibleNetwork HlfAccessibleNetwork

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
	clientConn, gateway, err := OpenConnection()
	if err != nil {
		t.Errorf("Error opening connection: %v", err)
	}
	defer CloseConnection(clientConn, gateway)
	var hlfAccessibleNetwork HlfAccessibleNetwork

	err = hlfAccessibleNetwork.RemoveAccessibleNetwork(gateway, "AccessibleNetwork2")
	if err != nil {
		t.Errorf(" Accessible Netwrok Exists : %v", err)
	}

	err = hlfAccessibleNetwork.RemoveAccessibleNetwork(gateway, "AccessibleNetwork1000")
	if err == nil {
		fmt.Printf(" Accessible Netwrok Exists : %v", err)
	}

}
