package permitted_method_handler

import (
	"fmt"
	"github.com/ethereum/go-ethereum/crypto"
	"github.com/tcdt-lab/Automated-Gateways/relay/configs"
	"log"
	"math/big"
	"testing"
	"time"
)

func TestEthPermittedMethodHandler_CreatePermittedMethod(t *testing.T) {
	configs.YamlAddress = "/home/koosha/Desktop/Automated-Gateways/relay/cmd/config.yaml"
	var ethPermittedMethod EthPermittedMethodHandler
	privateKey, err := crypto.HexToECDSA("c87509a1c067bbde78beb793e6fa76530b6382a4c0241e5e4a9ec0a0f44dc0d3")
	if err != nil {

		t.Errorf(err.Error())
	}

	res, err := ethPermittedMethod.CreatePermittedMethod(privateKey, "methodNAme34", "chaincodeNAme34", "inputArgs34", "OutputArgs34", 300000, big.NewInt(0))

	fmt.Println(res.Hash().Hex())
	time.Sleep(20 * time.Second)
	receipt, err := ethPermittedMethod.GetTransactionReceipt(res.Hash())
	if err != nil {
		log.Fatal(err)
		t.Errorf(err.Error())
	}
	fmt.Print(receipt.Status)
}

func TestEthPermittedMethodHandler_GetAllPermittedMethods(t *testing.T) {
	configs.YamlAddress = "/home/koosha/Desktop/Automated-Gateways/relay/cmd/config.yaml"
	var ethPermittedMethod EthPermittedMethodHandler
	privateKey, err := crypto.HexToECDSA("c87509a1c067bbde78beb793e6fa76530b6382a4c0241e5e4a9ec0a0f44dc0d3")
	if err != nil {
		t.Errorf("Error : %v", err)
	}
	res, err := ethPermittedMethod.GetAllPermittedMethods(privateKey)
	if err != nil {

		t.Errorf("Error getting all permitted Method: %v", err)
	}
	if res == nil {
		t.Errorf("Error getting all  permitted Method: The List is empty")
	}

	for i := 0; i < len(res); i++ {
		fmt.Println(res[i])
	}
}

func TestEthPermittedMethodHandler_IsPermittedMethodExist(t *testing.T) {
	configs.YamlAddress = "/home/koosha/Desktop/Automated-Gateways/relay/cmd/config.yaml"
	var ethPermittedMethod EthPermittedMethodHandler
	privateKey, err := crypto.HexToECDSA("c87509a1c067bbde78beb793e6fa76530b6382a4c0241e5e4a9ec0a0f44dc0d3")
	if err != nil {
		t.Errorf("Error : %v", err)
	}
	res, err := ethPermittedMethod.PermittedMethodExist(privateKey, "2")
	if err != nil {
		t.Errorf("Error getting an existed permitted: %v", err)
	}
	if res == false {
		t.Errorf("Error getting an existed permitted Method networks")
	}

	res, err = ethPermittedMethod.PermittedMethodExist(privateKey, "222")
	if err != nil {
		t.Errorf("Error getting an existed permitted: %v", err)
	}
	if res == true {
		t.Errorf("Error getting an existed permitted Method networks")
	}
}

func TestEthPermittedMethodHandler_GetPermittedMethod(t *testing.T) {
	configs.YamlAddress = "/home/koosha/Desktop/Automated-Gateways/relay/cmd/config.yaml"
	var ethPermittedMethod EthPermittedMethodHandler
	privateKey, err := crypto.HexToECDSA("c87509a1c067bbde78beb793e6fa76530b6382a4c0241e5e4a9ec0a0f44dc0d3")
	if err != nil {
		t.Errorf("Error : %v", err)
	}
	res, err := ethPermittedMethod.GetPermittedMethodById(privateKey, "2")
	if err != nil {
		t.Errorf("Error getting an existed permitted: %v", err)
	}

	fmt.Println(res)
}

func TestEthPermittedMethodHandler_UpdatePermittedMethod(t *testing.T) {
	configs.YamlAddress = "/home/koosha/Desktop/Automated-Gateways/relay/cmd/config.yaml"
	var ethPermittedMethod EthPermittedMethodHandler
	privateKey, err := crypto.HexToECDSA("c87509a1c067bbde78beb793e6fa76530b6382a4c0241e5e4a9ec0a0f44dc0d3")
	if err != nil {
		t.Errorf("Error : %v", err)
	}
	res, err := ethPermittedMethod.UpdatePermittedMethod(privateKey, "2", "methodNAme34", "chaincodeNAme34_updated", "inputArgs34_updated", "OutputArgs34_updated", 300000, big.NewInt(0))
	if err != nil {
		t.Errorf("Error getting an existed permitted: %v", err)
	}
	fmt.Println(res.Hash().Hex())
	time.Sleep(20 * time.Second)
	receipt, err := ethPermittedMethod.GetTransactionReceipt(res.Hash())
	if err != nil {

		t.Errorf(err.Error())
	}
	if receipt.Status != 1 {
		t.Errorf("Error transaction is failed. Status: %v", receipt.Status)
	}
	fmt.Println(receipt.Status)
}

func TestEthPermittedMethodHandler_DeletePermittedMethod(t *testing.T) {
	configs.YamlAddress = "/home/koosha/Desktop/Automated-Gateways/relay/cmd/config.yaml"
	var ethPermittedMethod EthPermittedMethodHandler
	privateKey, err := crypto.HexToECDSA("c87509a1c067bbde78beb793e6fa76530b6382a4c0241e5e4a9ec0a0f44dc0d3")
	if err != nil {
		t.Errorf("Error : %v", err)
	}
	res, err := ethPermittedMethod.DeletePermittedMethod(privateKey, "2", 300000, big.NewInt(0))
	if err != nil {
		t.Errorf("Error getting an existed permitted: %v", err)
	}
	fmt.Println(res.Hash().Hex())
	time.Sleep(20 * time.Second)
	receipt, err := ethPermittedMethod.GetTransactionReceipt(res.Hash())
	if err != nil {

		t.Errorf(err.Error())
	}
	if receipt.Status != 1 {
		t.Errorf("Error transaction is failed. Status: %v", receipt.Status)
	}
	fmt.Println(receipt.Status)
}
