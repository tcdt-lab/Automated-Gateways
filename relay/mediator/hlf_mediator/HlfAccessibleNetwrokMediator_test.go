package hlf_mediator

import (
	"fmt"
	"testing"
)

func TestCreateAccessibleNetwork(t *testing.T) {
	var hlfAccessibleNetwork HlfAccessibleNetworkMediator
	submitRes, err := hlfAccessibleNetwork.CreateAccessibleNetwork("TestAccessibleNet1", "127.0.0.1", "localhost", "KooshaComp")
	if err != nil {
		t.Errorf("Creating Accessible Netwrok: %v", err)
	}
	println(submitRes.AccessibleNetworkId)
}

func TestAccessibleNetworkExists(t *testing.T) {
	var hlfAccessibleNetwork HlfAccessibleNetworkMediator

	res, err := hlfAccessibleNetwork.AccessibleNetworkExists("AccessibleNetwork_localhost_7")
	if err != nil {
		t.Errorf(" Accessible Netwrok Exists : %v", err)
	}
	if res != true {
		t.Errorf(" Accessible Netwrok Exists : %v", res)
	}
	wrongRes, err := hlfAccessibleNetwork.AccessibleNetworkExists("AccessibleNetwork_localhost_7888")
	if wrongRes != false {
		t.Errorf(" Accessible Netwrok Exists : %v", wrongRes)
	}

}

func TestGetAccessibleNetwork(t *testing.T) {
	var hlfAccessibleNetwork HlfAccessibleNetworkMediator

	res, err := hlfAccessibleNetwork.GetAccessibleNetwork("AccessibleNetwork_localhost_7")
	if err != nil {
		t.Errorf(" Accessible Netwrok Exists : %v", err)
	}
	if res.AccessibleNetworkId != "AccessibleNetwork_localhost_7" {
		t.Errorf(" Accessible Netwrok Exists : %v", res)
	}
	wrongRes, err := hlfAccessibleNetwork.GetAccessibleNetwork("AccessibleNetwork_localhost_7888")
	if wrongRes != nil {
		t.Errorf(" Accessible Netwrok Exists : %v", wrongRes)
	}

}

func TestGetAllAccessibleNetworksByAddress(t *testing.T) {
	var hlfAccessibleNetwork HlfAccessibleNetworkMediator

	res, err := hlfAccessibleNetwork.GetAllAccessibleNetworksByAddress("localhost")
	if err != nil {
		t.Errorf(" Accessible Netwrok Exists : %v", err)
	}
	if len(res) == 0 {
		t.Errorf(" Accessible Netwrok Exists : %v", res)
	}
	wrongRes, err := hlfAccessibleNetwork.GetAllAccessibleNetworksByAddress("localhost2")
	if err != nil {
		t.Errorf(" Accessible Netwrok Exists : %v", wrongRes)
	}
}

func TestRemoveAccessibleNetwork(t *testing.T) {
	var hlfAccessibleNetwork HlfAccessibleNetworkMediator

	err := hlfAccessibleNetwork.RemoveAccessibleNetwork("AccessibleNetwork_localhost_7")
	if err != nil {
		t.Errorf(" Accessible Netwrok Exists : %v", err)
	}
}

func TestUpdateAccessibleNetwork(t *testing.T) {
	var hlfAccessibleNetwork HlfAccessibleNetworkMediator

	err := hlfAccessibleNetwork.UpdateAccessibleNetwork("AccessibleNetwork_localhost_7", "NetworkName3", "127.0.0.1", "localhost", "KooshaComp3")
	if err != nil {
		t.Errorf(" Accessible Netwrok Exists : %v", err)
	}
	err = hlfAccessibleNetwork.UpdateAccessibleNetwork("AccessibleNetwork_localhost_7324", "NetworkName3", "127.0.0.1", "localhost", "KooshaComp3")
	if err == nil {
		fmt.Printf(" Cannot Detect fake ID : %v", err)
	}
}
