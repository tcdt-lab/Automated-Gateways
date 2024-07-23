package permitted_network_handler

import (
	"fmt"
	"github.com/ethereum/go-ethereum/crypto"
	"github.com/tcdt-lab/Automated-Gateways/relay/configs"
	"log"
	"math/big"
	"testing"
	"time"
)

func TestEthPermittedNetworkHandler_CreatePermittedNetwork(t *testing.T) {
	configs.YamlAddress = "/home/koosha/Desktop/Automated-Gateways/relay/cmd/config.yaml"
	var ethPermittedNetwork EthPermittedNetworkHandler
	privateKey, err := crypto.HexToECDSA("c87509a1c067bbde78beb793e6fa76530b6382a4c0241e5e4a9ec0a0f44dc0d3")
	if err != nil {

		t.Errorf(err.Error())
	}

	res, err := ethPermittedNetwork.CreatePermittedNetwork(privateKey, "netNum_permitted_new2", "192.168.1.11", "localhost2", "KooshaCompNum3_permitted2", 300000, big.NewInt(0))

	fmt.Println(res.Hash().Hex())
	time.Sleep(30 * time.Second)
	receipt, err := ethPermittedNetwork.GetTransactionReceipt(res.Hash())
	if err != nil {
		log.Fatal(err)
		t.Errorf(err.Error())
	}
	fmt.Print(receipt.Status)
}

func TestEthPermittedNetworkHandler_GetAllPermittedNetworks(t *testing.T) {
	configs.YamlAddress = "/home/koosha/Desktop/Automated-Gateways/relay/cmd/config.yaml"
	var ethPermittedNetwork EthPermittedNetworkHandler
	privateKey, err := crypto.HexToECDSA("c87509a1c067bbde78beb793e6fa76530b6382a4c0241e5e4a9ec0a0f44dc0d3")
	if err != nil {
		t.Errorf("Error : %v", err)
	}
	res, err := ethPermittedNetwork.GetAllPermittedNetworks(privateKey)
	if err != nil {

		t.Errorf("Error getting all accessible networks: %v", err)
	}
	if res == nil {
		t.Errorf("Error getting all accessible networks: The List is empty")
	}

	for i := 0; i < len(res); i++ {
		fmt.Println(res[i])
	}
}

func TestEthPermittedNetworkHandler_IsPermittedNetworkExist(t *testing.T) {
	configs.YamlAddress = "/home/koosha/Desktop/Automated-Gateways/relay/cmd/config.yaml"
	var ethPermittedNetwork EthPermittedNetworkHandler
	privateKey, err := crypto.HexToECDSA("ae6ae8e5ccbfb04590405997ee2d52d2b330726137b875053c36d94e974d162f")
	if err != nil {
		t.Errorf("Error getting all accessible networks: %v", err)
	}
	res, err := ethPermittedNetwork.IsPermittedNetworkExist(privateKey, "2")
	if err != nil {

		t.Errorf("Error getting all accessible networks: %v", err)
	}
	if res == false {
		t.Errorf("Error getting checking if a permitted network existed")
	}

	res, err = ethPermittedNetwork.IsPermittedNetworkExist(privateKey, "222")
	if err != nil {

		t.Errorf("Error getting all accessible networks: %v", err)
	}
	if res == true {
		t.Errorf("Error getting all accessible networks: the Id should not exist")
	}
}

func TestEthPermittedNetworkHandler_GetPermittedNetwork(t *testing.T) {
	configs.YamlAddress = "/home/koosha/Desktop/Automated-Gateways/relay/cmd/config.yaml"
	var ethPermittedNetwork EthPermittedNetworkHandler
	privateKey, err := crypto.HexToECDSA("c87509a1c067bbde78beb793e6fa76530b6382a4c0241e5e4a9ec0a0f44dc0d3")
	if err != nil {
		log.Fatal(err)
	}
	res, err := ethPermittedNetwork.GetPermittedNetwork(privateKey, "1")
	if err != nil {

		t.Errorf("Error getting all accessible networks: %v", err)
	}
	if res == nil {
		t.Errorf("Error getting all accessible networks: The List is empty")
	}

	fmt.Println(res)
}

func TestEthPermittedNetworkHandler_UpdatePermittedNetwork(t *testing.T) {
	configs.YamlAddress = "/home/koosha/Desktop/Automated-Gateways/relay/cmd/config.yaml"
	var ethPermittedNetwork EthPermittedNetworkHandler
	privateKey, err := crypto.HexToECDSA("ae6ae8e5ccbfb04590405997ee2d52d2b330726137b875053c36d94e974d162f")
	if err != nil {
		t.Errorf("Error getting all accessible networks: %v", err)
	}
	res, err := ethPermittedNetwork.UpdatePermittedNetwork(privateKey, "1", "netNum_permitted_updated", "newIP_U", "newAddress_U", "newCompanyName_U", 300000, big.NewInt(0))
	if err != nil {
		t.Errorf("Error getting all accessible networks: %v", err)
	}
	fmt.Println(res.Hash().Hex())
}

func TestEthPermittedNetworkHandler_RemovePermittedNetwork(t *testing.T) {
	configs.YamlAddress = "/home/koosha/Desktop/Automated-Gateways/relay/cmd/config.yaml"
	var ethPErmittedNetwork EthPermittedNetworkHandler
	privateKey, err := crypto.HexToECDSA("ae6ae8e5ccbfb04590405997ee2d52d2b330726137b875053c36d94e974d162f")
	if err != nil {
		t.Errorf("Error getting all accessible networks: %v", err)
	}
	res, err := ethPErmittedNetwork.RemovePermittedNetwork(privateKey, "1", 300000, big.NewInt(0))
	if err != nil {
		t.Errorf("Error getting all accessible networks: %v", err)
	}
	fmt.Println(res.Hash().Hex())
}
