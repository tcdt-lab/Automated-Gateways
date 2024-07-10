package accessible_network_handler

import (
	"fmt"
	"github.com/ethereum/go-ethereum/crypto"
	"github.com/tcdt-lab/Automated-Gateways/relay/configs"
	"log"
	"testing"
)

func TestEthAccessibleNetworkHandler_CreateAccessibleNetwork(t *testing.T) {
	configs.YamlAddress = "/home/koosha/Desktop/Automated-Gateways/relay/cmd/config.yaml"
	var ethAccessibleNetwork EthAccessibleNetworkHandler
	privateKey, err := crypto.HexToECDSA("ae6ae8e5ccbfb04590405997ee2d52d2b330726137b875053c36d94e974d162f")
	if err != nil {
		log.Fatal(err)
	}

	res, err := ethAccessibleNetwork.CreateAccessibleNetwork(privateKey, "TestAccessibleNet1", "192.168.1.1", "localhost", "KooshaComp")
	fmt.Print(res)
}
