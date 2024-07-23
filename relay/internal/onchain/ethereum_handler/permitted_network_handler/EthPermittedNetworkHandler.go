package permitted_network_handler

import (
	"context"
	"crypto/ecdsa"
	"errors"
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/crypto"
	eth "github.com/ethereum/go-ethereum/ethclient"
	"github.com/tcdt-lab/Automated-Gateways/relay/configs"
	"github.com/tcdt-lab/Automated-Gateways/relay/data_types"
	"github.com/tcdt-lab/Automated-Gateways/relay/internal/onchain/ethereum_handler"
	"log"
	"math/big"
)

type EthPermittedNetworkHandler struct {
}

func (ethPermittedNetworkHandler *EthPermittedNetworkHandler) Connect(cnf configs.Configuration, err error) (*eth.Client, error) {
	var ethereumURL = cnf.Platform.Ethereum.Url
	conn, err := eth.Dial(ethereumURL)
	if err != nil {
		log.Fatal(err)
		return nil, err
	}
	return conn, nil
}

func (ethPermittedNetworkHandler *EthPermittedNetworkHandler) CreatePermittedNetwork(privateKey *ecdsa.PrivateKey, networkName string, networkIP string, networkAddress string, companyName string, gasLimit uint64, value *big.Int) (*types.Transaction, error) {
	cnf, err := configs.ReadConfigYAMLFile()

	if err != nil {
		log.Fatal(err)
		return nil, err
	}
	var contractAddress = cnf.Platform.Ethereum.PermittedNetworkSmartContractAddress
	conn, err_connect := ethPermittedNetworkHandler.Connect(cnf, err)
	if err_connect != nil {
		return nil, err_connect
	}

	chainID, err := conn.ChainID(context.Background())

	publicKey := privateKey.Public()
	publicKeyECDSA, ok := publicKey.(*ecdsa.PublicKey)
	if !ok {
		log.Fatal("cannot assert type: publicKey is not of type *ecdsa.PublicKey")
		return nil, errors.New("cannot assert type: publicKey is not of type *ecdsa.PublicKey")
	}

	fromAddress := crypto.PubkeyToAddress(*publicKeyECDSA)
	nonce, err := conn.PendingNonceAt(context.Background(), fromAddress)
	if err != nil {
		log.Fatal(err)
		return nil, err
	}

	gasPrice, err := conn.SuggestGasPrice(context.Background())
	if err != nil {
		log.Fatal(err)
		return nil, err
	}

	transOpts, err := bind.NewKeyedTransactorWithChainID(privateKey, chainID)
	transOpts.Nonce = big.NewInt(int64(nonce))
	transOpts.Value = value       // in wei
	transOpts.GasLimit = gasLimit // in units
	transOpts.GasPrice = gasPrice
	accessibleManager, err := ethereum_handler.NewPermittedNetworkManagement(common.HexToAddress(contractAddress), conn)
	if err != nil {
		log.Fatalf("Failed to instantiate Storage contract: %v", err)
		return nil, err
	}

	transaction, err := accessibleManager.CreatePermittedNetwork(transOpts, networkName, networkIP, networkAddress, companyName)

	if err != nil {
		log.Fatalf("Failed to create Permitted Network: %v", err)
		return nil, err
	}

	return transaction, nil
}

func (ethPermittedNetworkHandler *EthPermittedNetworkHandler) GetTransactionReceipt(hash common.Hash) (*types.Receipt, error) {
	cnf, err := configs.ReadConfigYAMLFile()
	conn, err_connect := ethPermittedNetworkHandler.Connect(cnf, err)
	if err_connect != nil {
		return nil, err_connect
	}
	receipt, err := conn.TransactionReceipt(context.Background(), hash)
	if err != nil {
		log.Fatal(err)
		return nil, err
	}
	return receipt, nil
}

func (ethPermittedNetworkHandler *EthPermittedNetworkHandler) GetAllPermittedNetworks(accountPrivateKey *ecdsa.PrivateKey) ([]*data_types.PermittedNetworkInfo, error) {
	var permittedNetworkInfoList []*data_types.PermittedNetworkInfo
	cnf, err := configs.ReadConfigYAMLFile()

	if err != nil {
		log.Fatal(err)
		return nil, err
	}
	var contractAddress = cnf.Platform.Ethereum.PermittedNetworkSmartContractAddress
	conn, err_connect := ethPermittedNetworkHandler.Connect(cnf, err)
	if err_connect != nil {
		return nil, err_connect
	}

	publicKey := accountPrivateKey.Public()
	publicKeyECDSA, ok := publicKey.(*ecdsa.PublicKey)
	if !ok {
		log.Fatal("cannot assert type: publicKey is not of type *ecdsa.PublicKey")
		return nil, errors.New("cannot assert type: publicKey is not of type *ecdsa.PublicKey")
	}

	fromAddress := crypto.PubkeyToAddress(*publicKeyECDSA)
	CallOpts := &bind.CallOpts{
		Pending: true,
		From:    fromAddress,
		Context: context.Background(),
	}
	permittedManager, err := ethereum_handler.NewPermittedNetworkManagement(common.HexToAddress(contractAddress), conn)
	res, err := permittedManager.GetAllPermittedNetworks(CallOpts)
	for _, network := range res {
		networkInfo := convertContractPermittedNetworkDataType(network)
		if err != nil {
			log.Printf("failed to unmarshal network info: %s", err)
			return nil, err
		}
		permittedNetworkInfoList = append(permittedNetworkInfoList, networkInfo)
	}
	conn.Close()
	return permittedNetworkInfoList, nil

}

func (ethPermittedNetworkHandler *EthPermittedNetworkHandler) IsPermittedNetworkExist(accountPrivateKey *ecdsa.PrivateKey, networkId string) (bool, error) {
	cnf, err := configs.ReadConfigYAMLFile()
	if err != nil {
		log.Fatal(err)
		return false, err
	}
	var contractAddress = cnf.Platform.Ethereum.PermittedNetworkSmartContractAddress
	conn, err_connect := ethPermittedNetworkHandler.Connect(cnf, err)
	if err_connect != nil {
		return false, err_connect
	}

	publicKey := accountPrivateKey.Public()
	publicKeyECDSA, ok := publicKey.(*ecdsa.PublicKey)
	if !ok {
		log.Fatal("cannot assert type: publicKey is not of type *ecdsa.PublicKey")
		return false, errors.New("cannot assert type: publicKey is not of type *ecdsa.PublicKey")
	}

	fromAddress := crypto.PubkeyToAddress(*publicKeyECDSA)
	CallOpts := &bind.CallOpts{
		Pending: true,
		From:    fromAddress,
		Context: context.Background(),
	}
	permittedManager, err := ethereum_handler.NewPermittedNetworkManagement(common.HexToAddress(contractAddress), conn)
	intNetId, err := stringToBigIntConvertor(networkId)
	if err != nil {
		log.Fatal(err)
		return false, err

	}
	res, err := permittedManager.PermittedNetworkExists(CallOpts, intNetId)
	if err != nil {
		log.Printf("failed to unmarshal network info: %s", err)
		return false, err
	}
	conn.Close()
	return res, nil
}

func (ethPermittedNetworkHandler *EthPermittedNetworkHandler) RemovePermittedNetwork(privateKey *ecdsa.PrivateKey, networkId string, gasLimit uint64, value *big.Int) (*types.Transaction, error) {
	cnf, err := configs.ReadConfigYAMLFile()

	if err != nil {
		log.Fatal(err)
		return nil, err
	}
	var contractAddress = cnf.Platform.Ethereum.PermittedNetworkSmartContractAddress
	conn, err_connect := ethPermittedNetworkHandler.Connect(cnf, err)
	if err_connect != nil {
		return nil, err_connect
	}

	chainID, err := conn.ChainID(context.Background())

	publicKey := privateKey.Public()
	publicKeyECDSA, ok := publicKey.(*ecdsa.PublicKey)
	if !ok {
		log.Fatal("cannot assert type: publicKey is not of type *ecdsa.PublicKey")
		return nil, errors.New("cannot assert type: publicKey is not of type *ecdsa.PublicKey")
	}

	fromAddress := crypto.PubkeyToAddress(*publicKeyECDSA)
	nonce, err := conn.PendingNonceAt(context.Background(), fromAddress)
	if err != nil {
		log.Fatal(err)
		return nil, err
	}

	gasPrice, err := conn.SuggestGasPrice(context.Background())
	if err != nil {
		log.Fatal(err)
		return nil, err
	}

	transOpts, err := bind.NewKeyedTransactorWithChainID(privateKey, chainID)
	transOpts.Nonce = big.NewInt(int64(nonce))
	transOpts.GasLimit = gasLimit // in units
	transOpts.Value = value
	transOpts.GasPrice = gasPrice
	permittedManager, err := ethereum_handler.NewPermittedNetworkManagement(common.HexToAddress(contractAddress), conn)
	if err != nil {
		log.Fatalf("Failed to instantiate Storage contract: %v", err)
		return nil, err
	}

	intNetId, err := stringToBigIntConvertor(networkId)
	if err != nil {
		log.Fatal(err)
		return nil, err
	}
	transaction, err := permittedManager.DeletePermittedNetwork(transOpts, intNetId)
	if err != nil {
		log.Fatalf("Failed to remove Permitted Network: %v", err)
		return nil, err
	}
	return transaction, nil
}

func (ethPermittedNetworkHandler *EthPermittedNetworkHandler) UpdatePermittedNetwork(privateKey *ecdsa.PrivateKey, networkId string, networkName string, networkIP string, networkAddress string, companyName string, gasLimit uint64, value *big.Int) (*types.Transaction, error) {
	cnf, err := configs.ReadConfigYAMLFile()

	if err != nil {
		log.Fatal(err)
		return nil, err
	}
	var contractAddress = cnf.Platform.Ethereum.PermittedNetworkSmartContractAddress
	conn, err_connect := ethPermittedNetworkHandler.Connect(cnf, err)
	if err_connect != nil {
		return nil, err_connect
	}

	chainID, err := conn.ChainID(context.Background())

	publicKey := privateKey.Public()
	publicKeyECDSA, ok := publicKey.(*ecdsa.PublicKey)
	if !ok {
		log.Fatal("cannot assert type: publicKey is not of type *ecdsa.PublicKey")
		return nil, errors.New("cannot assert type: publicKey is not of type *ecdsa.PublicKey")
	}

	fromAddress := crypto.PubkeyToAddress(*publicKeyECDSA)
	nonce, err := conn.PendingNonceAt(context.Background(), fromAddress)
	if err != nil {
		log.Fatal(err)
		return nil, err
	}

	gasPrice, err := conn.SuggestGasPrice(context.Background())
	if err != nil {
		log.Fatal(err)
		return nil, err
	}

	transOpts, err := bind.NewKeyedTransactorWithChainID(privateKey, chainID)
	transOpts.Nonce = big.NewInt(int64(nonce))
	transOpts.GasLimit = gasLimit // in units
	transOpts.GasPrice = gasPrice
	transOpts.Value = value
	permittedManager, err := ethereum_handler.NewPermittedNetworkManagement(common.HexToAddress(contractAddress), conn)
	if err != nil {
		log.Fatalf("Failed to instantiate Storage contract: %v", err)
		return nil, err
	}

	intNetId, err := stringToBigIntConvertor(networkId)
	if err != nil {
		log.Fatal(err)
		return nil, err
	}

	transaction, err := permittedManager.UpdatePermittedNetwork(transOpts, intNetId, networkName, networkIP, networkAddress, companyName)
	if err != nil {
		log.Fatalf("Failed to update Permitted Network: %v", err)
		return nil, err
	}
	return transaction, nil
}

func (ethPermittedNetworkHandler *EthPermittedNetworkHandler) GetPermittedNetwork(accountPrivateKey *ecdsa.PrivateKey, networkId string) (*data_types.PermittedNetworkInfo, error) {
	var permittedNetworkInfo *data_types.PermittedNetworkInfo
	cnf, err := configs.ReadConfigYAMLFile()

	if err != nil {
		log.Fatal(err)
		return nil, err
	}
	var contractAddress = cnf.Platform.Ethereum.PermittedNetworkSmartContractAddress
	conn, err_connect := ethPermittedNetworkHandler.Connect(cnf, err)
	if err_connect != nil {
		return nil, err_connect
	}

	publicKey := accountPrivateKey.Public()
	publicKeyECDSA, ok := publicKey.(*ecdsa.PublicKey)
	if !ok {
		log.Fatal("cannot assert type: publicKey is not of type *ecdsa.PublicKey")
		return nil, errors.New("cannot assert type: publicKey is not of type *ecdsa.PublicKey")
	}

	fromAddress := crypto.PubkeyToAddress(*publicKeyECDSA)
	CallOpts := &bind.CallOpts{
		Pending: true,
		From:    fromAddress,
		Context: context.Background(),
	}
	permittedManager, err := ethereum_handler.NewPermittedNetworkManagement(common.HexToAddress(contractAddress), conn)
	intNetId, err := stringToBigIntConvertor(networkId)
	if err != nil {
		log.Fatal(err)
		return nil, err

	}
	res, err := permittedManager.GetPermittedNetwork(CallOpts, intNetId)
	if err != nil {
		log.Printf("failed to unmarshal network info: %s", err)
		return nil, err
	}
	permittedNetworkInfo = convertContractPermittedNetworkDataType(res)
	conn.Close()

	if permittedNetworkInfo.PermittedNetworkId == "0" {
		return nil, errors.New("The network does not exist")
	}
	return permittedNetworkInfo, nil
}
