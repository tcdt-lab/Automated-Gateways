package accessible_network_handler

import (
	"fmt"
	"github.com/ethereum/go-ethereum/crypto"
	"github.com/tcdt-lab/Automated-Gateways/relay/configs"
	"github.com/tcdt-lab/Automated-Gateways/relay/data_types"
	"log"
	"math/big"
	"testing"
	"time"
)

func TestEthAccessibleNetworkHandler_CreateAccessibleNetwork(t *testing.T) {
	configs.YamlAddress = "/home/koosha/Desktop/Automated-Gateways/relay/cmd/config.yaml"
	var ethAccessibleNetwork EthAccessibleNetworkHandler
	privateKey, err := crypto.HexToECDSA("c87509a1c067bbde78beb793e6fa76530b6382a4c0241e5e4a9ec0a0f44dc0d3")
	if err != nil {
		log.Fatal(err)
		t.Errorf(err.Error())
	}

	res, err := ethAccessibleNetwork.CreateAccessibleNetwork(privateKey, "netNum3", "192.168.1.11", "localhost", "KooshaCompNum3", 300000, big.NewInt(0))

	fmt.Print(res.Hash().Hex())
	time.Sleep(30 * time.Second)
	receipt, err := ethAccessibleNetwork.GetTransactionReceipt(res.Hash())
	if err != nil {

		t.Errorf(err.Error())
	}
	if receipt.Status != 1 {
		t.Errorf("Error transaction is failed. Status: %v", receipt.Status)
	}
	fmt.Println(receipt.Status)
}

func TestEthAccessibleNetworkHandler_GetAllAccessibleNetworks(t *testing.T) {
	configs.YamlAddress = "/home/koosha/Desktop/Automated-Gateways/relay/cmd/config.yaml"
	var ethAccessibleNetwork EthAccessibleNetworkHandler
	privateKey, err := crypto.HexToECDSA("c87509a1c067bbde78beb793e6fa76530b6382a4c0241e5e4a9ec0a0f44dc0d3")
	if err != nil {
		log.Fatal(err)
	}
	res, err := ethAccessibleNetwork.GetAllAccessibleNetworks(privateKey)
	if err != nil {
		log.Fatal(err)
		t.Errorf("Error getting all accessible networks: %v", err)
	}
	if res == nil {
		t.Errorf("Error getting all accessible networks: The List is empty")
	}

	for i := 0; i < len(res); i++ {
		fmt.Println(res[i])
	}

}

func TestEthAccessibleNetworkHandler_IsAccessibleNetworkExist(t *testing.T) {
	configs.YamlAddress = "/home/koosha/Desktop/Automated-Gateways/relay/cmd/config.yaml"
	var ethAccessibleNetwork EthAccessibleNetworkHandler
	privateKey, err := crypto.HexToECDSA("ae6ae8e5ccbfb04590405997ee2d52d2b330726137b875053c36d94e974d162f")
	if err != nil {
		log.Fatal(err)
	}
	res, err := ethAccessibleNetwork.AccessibleNetworkExists(privateKey, "2")
	if err != nil {

		t.Errorf("Error getting accessible network: %v", err)
	}
	if !res {
		t.Errorf("Error getting a false result while the id does  exist")
	}

	res, err = ethAccessibleNetwork.AccessibleNetworkExists(privateKey, "10000")
	if err != nil {

		t.Errorf("Error getting accessible network: %v", err)
	}
	if res {
		t.Errorf("Error getting a true result while the id does not exist")
	}
}

func TestEthAccessibleNetworkHandler_UpdateAccessibleNetwork(t *testing.T) {
	configs.YamlAddress = "/home/koosha/Desktop/Automated-Gateways/relay/cmd/config.yaml"
	var ethAccessibleNetwork EthAccessibleNetworkHandler
	privateKey, err := crypto.HexToECDSA("ae6ae8e5ccbfb04590405997ee2d52d2b330726137b875053c36d94e974d162f")
	if err != nil {
		log.Fatal(err)
	}
	accessibleNetInfo := data_types.AccessibleNetworkInfo{
		AccessibleNetworkId: "1",
		NetworkName:         "TestAccessibleNet1",
		IP:                  "11",
		ADDRESS:             "11asf",
		CompanyName:         "KooshaComUpdated",
	}
	res, err := ethAccessibleNetwork.UpdateAccessibleNetwork(privateKey, 300000, big.NewInt(0), &accessibleNetInfo)
	if err != nil {
		log.Fatal(err)
		t.Errorf("Error updating accessible network: %v", err)
	}
	fmt.Print(res)

	resAll, err := ethAccessibleNetwork.GetAllAccessibleNetworks(privateKey)
	if err != nil {
		log.Fatal(err)
		t.Errorf("Error getting all accessible networks: %v", err)
	}
	fmt.Print(resAll[0].CompanyName)

}

func TestEthAccessibleNetworkHandler_GetAccessibleNetworksNumber(t *testing.T) {
	configs.YamlAddress = "/home/koosha/Desktop/Automated-Gateways/relay/cmd/config.yaml"
	var ethAccessibleNetwork EthAccessibleNetworkHandler
	privateKey, err := crypto.HexToECDSA("8f2a55949038a9610f50fb23b5883af3b4ecb3c3bb792cbcefbd1542c692be63")
	if err != nil {
		log.Fatal(err)
	}
	res, err := ethAccessibleNetwork.GetAccessibleNetworksNumber(privateKey)
	if err != nil {
		log.Fatal(err)
		t.Errorf("Error getting accessible network number: %v", err)
	}
	fmt.Print(res)
}

func TestEthAccessibleNetworkHandler_GetAccessibleNetworkById(t *testing.T) {
	configs.YamlAddress = "/home/koosha/Desktop/Automated-Gateways/relay/cmd/config.yaml"
	var ethAccessibleNetwork EthAccessibleNetworkHandler
	privateKey, err := crypto.HexToECDSA("8f2a55949038a9610f50fb23b5883af3b4ecb3c3bb792cbcefbd1542c692be63")
	if err != nil {
		log.Fatal(err)
	}
	res, err := ethAccessibleNetwork.GetAccessibleNetwork(privateKey, "2")
	if err != nil {
		log.Fatal(err)
		t.Errorf("Error getting accessible network by id: %v", err)
	}
	fmt.Print(res)
}

func TestEthAccessibleNetworkHandler_DeleteAccessibleNetwork(t *testing.T) {
	configs.YamlAddress = "/home/koosha/Desktop/Automated-Gateways/relay/cmd/config.yaml"
	var ethAccessibleNetwork EthAccessibleNetworkHandler
	privateKey, err := crypto.HexToECDSA("8f2a55949038a9610f50fb23b5883af3b4ecb3c3bb792cbcefbd1542c692be63")
	if err != nil {
		log.Fatal(err)
	}
	res, err := ethAccessibleNetwork.DeleteAccessibleNetwork(privateKey, "1", 300000, big.NewInt(0))
	if err != nil {
		log.Fatal(err)
		t.Errorf("Error deleting accessible network: %v", err)
	}
	fmt.Print(res)
}
