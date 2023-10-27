package hlf_mediator

import (
	"testing"
)

func TestCreatePermittedNetwork(t *testing.T) {
	var permittedNetwork HlfPermittedNetworkMediator
	submitRes, err := permittedNetwork.CreatePermittedNetwork("TestPermittedNet1", "127.0.0.1", "localhost", "KooshaComp")
	if err != nil {
		t.Errorf("Creating Accessible Netwrok: %v", err)
	}
	println(submitRes.PermittedNetworkId)
}

func TestPermittedNetworkExists(t *testing.T) {
	var permittedNetwork HlfPermittedNetworkMediator

	res, err := permittedNetwork.PermittedNetworkExists("PermittedNetwork_localhost_7")
	if err != nil {
		t.Errorf(" Accessible Netwrok Exists : %v", err)
	}
	if res != true {
		t.Errorf(" Accessible Netwrok Exists : %v", res)
	}
	wrongRes, err := permittedNetwork.PermittedNetworkExists("PermittedNetwork_localhost_7888")
	if wrongRes != false {
		t.Errorf(" Accessible Netwrok Exists : %v", wrongRes)
	}

}

func TestGetPermittedNetwork(t *testing.T) {
	var permittedNetwork HlfPermittedNetworkMediator

	res, err := permittedNetwork.GetPermittedNetwork("PermittedNetwork_localhost_7")
	if err != nil {
		t.Errorf(" Accessible Netwrok Exists : %v", err)
	}
	if res.PermittedNetworkId != "PermittedNetwork_localhost_7" {
		t.Errorf(" Accessible Netwrok Exists : %v", res)
	}
	wrongRes, err := permittedNetwork.GetPermittedNetwork("PermittedNetwork_localhost_7888")
	if err == nil {
		t.Errorf(" Accessible Netwrok Exists : %v", wrongRes)
	}

}

func TestGetAllPermittedNetworksByAddress(t *testing.T) {
	var permittedNetwork HlfPermittedNetworkMediator

	res, err := permittedNetwork.GetAllPermittedNetworksByAddress("localhost")
	if err != nil {
		t.Errorf(" Accessible Netwrok Exists : %v", err)
	}
	if len(res) == 0 {
		t.Errorf(" Accessible Netwrok Exists : %v", res)
	}
}

func TestRemovePermittedNetwork(t *testing.T) {
	var permittedNetwork HlfPermittedNetworkMediator

	err := permittedNetwork.RemovePermittedNetwork("PermittedNetwork_localhost_6")
	if err != nil {
		t.Errorf(" Accessible Netwrok Exists : %v", err)
	}

}

func TestUpdatePermittedNetwork(t *testing.T) {
	var permittedNetwork HlfPermittedNetworkMediator

	err := permittedNetwork.UpdatePermittedNetwork("PermittedNetwork_localhost_7", "TestPermittedNetupdated", "127.0.0.1", "localhost", "KooshaComp")
	if err != nil {
		t.Errorf(" Error in Updating Permitted Netwrok : %v", err)
	}
}

func TestGetPermittedNetworkByIndexAndAddress(t *testing.T) {
	var permittedNetwork HlfPermittedNetworkMediator

	res, err := permittedNetwork.GetPermittedNetworkByIndexAndAddress("1", "900", "localhost")
	if err != nil {
		t.Errorf(" Error in Getting Permitted Netwrok : %v", err)
	}
	if res[1].PermittedNetworkId != "PermittedNetwork_localhost_8" {
		t.Errorf(" Error in Getting Permitted Netwrok : %v", res)
	}
}
