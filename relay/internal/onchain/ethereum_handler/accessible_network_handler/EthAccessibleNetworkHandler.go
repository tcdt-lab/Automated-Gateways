package accessible_network_handler

import (
	"context"
	"crypto/ecdsa"
	"fmt"
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/crypto"
	eth "github.com/ethereum/go-ethereum/ethclient"
	"github.com/tcdt-lab/Automated-Gateways/relay/configs"
	dataTypes "github.com/tcdt-lab/Automated-Gateways/relay/data_types"
	"github.com/tcdt-lab/Automated-Gateways/relay/internal/onchain/ethereum_handler"
	"log"
	"math/big"
)

type EthAccessibleNetworkHandler struct {
}

var accountPrivateKey *ecdsa.PrivateKey

func (ethAccessibleNetworkHandler *EthAccessibleNetworkHandler) CreateAccessibleNetwork(accountPrivateKey *ecdsa.PrivateKey, networkName string, ip string, address string, companyName string) (*dataTypes.AccessibleNetworkInfo, error) {
	cnf, err := configs.ReadConfigYAMLFile()
	if err != nil {
		log.Fatal(err)
	}
	var ethereumURL = cnf.Platform.Ethereum.Url
	var contractAddress = cnf.Platform.Ethereum.AccessibleNetworkSmartContractAddress

	conn, err := eth.Dial(ethereumURL)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("we have a connection")

	chainID, err := conn.ChainID(context.Background())

	publicKey := accountPrivateKey.Public()
	publicKeyECDSA, ok := publicKey.(*ecdsa.PublicKey)
	if !ok {
		log.Fatal("cannot assert type: publicKey is not of type *ecdsa.PublicKey")
	}

	fromAddress := crypto.PubkeyToAddress(*publicKeyECDSA)
	nonce, err := conn.PendingNonceAt(context.Background(), fromAddress)
	if err != nil {
		log.Fatal(err)
	}

	gasPrice, err := conn.SuggestGasPrice(context.Background())
	if err != nil {
		log.Fatal(err)
	}

	auth, err := bind.NewKeyedTransactorWithChainID(accountPrivateKey, chainID)
	auth.Nonce = big.NewInt(int64(nonce))
	auth.Value = big.NewInt(0)     // in wei
	auth.GasLimit = uint64(300000) // in units
	auth.GasPrice = gasPrice
	permittedManager, err := ethereum_handler.NewAccessibleNetworkManagement(common.HexToAddress(contractAddress), conn)
	if err != nil {
		log.Fatalf("Failed to instantiate Storage contract: %v", err)
	}

	createRes, err := permittedManager.CreateAccessibleNetwork(auth, networkName, ip, address, companyName)

	if err != nil {
		log.Fatalf("Failed to create Permitted Network: %v", err)
	}
	fmt.Println("Permitted Network created: ", createRes.Hash())
	return nil, nil
}
