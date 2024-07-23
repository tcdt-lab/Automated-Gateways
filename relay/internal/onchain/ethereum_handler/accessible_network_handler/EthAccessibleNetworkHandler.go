package accessible_network_handler

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

type EthAccessibleNetworkHandler struct {
}

var accountPrivateKey *ecdsa.PrivateKey

func (ethAccessibleNetworkHandler *EthAccessibleNetworkHandler) CreateAccessibleNetwork(accountPrivateKey *ecdsa.PrivateKey, networkName string, ip string, address string, companyName string, gasLimit uint64, value *big.Int) (*types.Transaction, error) {
	cnf, err := configs.ReadConfigYAMLFile()

	if err != nil {
		log.Fatal(err)
		return nil, err
	}
	var contractAddress = cnf.Platform.Ethereum.AccessibleNetworkSmartContractAddress
	conn, err_connect := ethAccessibleNetworkHandler.Connect(cnf, err)
	if err_connect != nil {
		return nil, err_connect
	}

	chainID, err := conn.ChainID(context.Background())

	publicKey := accountPrivateKey.Public()
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

	transOpts, err := bind.NewKeyedTransactorWithChainID(accountPrivateKey, chainID)
	transOpts.Nonce = big.NewInt(int64(nonce))
	transOpts.Value = value       // in wei
	transOpts.GasLimit = gasLimit // in units
	transOpts.GasPrice = gasPrice
	accessibleManager, err := ethereum_handler.NewAccessibleNetworkManagement(common.HexToAddress(contractAddress), conn)
	if err != nil {
		log.Fatalf("Failed to instantiate Storage contract: %v", err)
		return nil, err
	}

	transaction, err := accessibleManager.CreateAccessibleNetwork(transOpts, networkName, ip, address, companyName)

	if err != nil {
		log.Fatalf("Failed to create Permitted Network: %v", err)
		return nil, err
	}

	return transaction, nil
}

func (ethAccessibleNetworkHandler *EthAccessibleNetworkHandler) GetAllAccessibleNetworks(accountPrivateKey *ecdsa.PrivateKey) ([]*data_types.AccessibleNetworkInfo, error) {

	var accessibleNetworkInfoList []*data_types.AccessibleNetworkInfo
	cnf, err := configs.ReadConfigYAMLFile()

	if err != nil {
		log.Fatal(err)
		return nil, err
	}
	var contractAddress = cnf.Platform.Ethereum.AccessibleNetworkSmartContractAddress
	conn, err_connect := ethAccessibleNetworkHandler.Connect(cnf, err)
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
	accessibleManager, err := ethereum_handler.NewAccessibleNetworkManagement(common.HexToAddress(contractAddress), conn)
	res, err := accessibleManager.GetAllAccessibleNetworks(CallOpts)
	for _, network := range res {
		networkInfo := convertContractAccessibleDataType(network)
		if err != nil {
			log.Printf("failed to unmarshal network info: %s", err)
			return nil, err
		}
		accessibleNetworkInfoList = append(accessibleNetworkInfoList, networkInfo)
	}
	conn.Close()
	return accessibleNetworkInfoList, nil
}

func (ethAccessibleNetworkHandler *EthAccessibleNetworkHandler) AccessibleNetworkExists(accountPrivateKey *ecdsa.PrivateKey, accessibleNetworkId string) (bool, error) {
	cnf, err := configs.ReadConfigYAMLFile()

	if err != nil {
		log.Fatal(err)
		return false, err
	}
	var contractAddress = cnf.Platform.Ethereum.AccessibleNetworkSmartContractAddress
	conn, err_connect := ethAccessibleNetworkHandler.Connect(cnf, err)
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
	accessibleManager, err := ethereum_handler.NewAccessibleNetworkManagement(common.HexToAddress(contractAddress), conn)
	n := new(big.Int)
	accessibleNetIdInt, isSuccess := n.SetString(accessibleNetworkId, 10)
	if !isSuccess {
		log.Fatalf("Failed to convert accessibleNetworkId to big.Int")
	}

	res, err := accessibleManager.IsAccessibleNetworkExist(CallOpts, accessibleNetIdInt)
	if err != nil {
		log.Printf("failed to check network existence: %s", err)
		return false, err
	}
	conn.Close()
	return res, nil
}

func (ethAccessibleNetworkHandler *EthAccessibleNetworkHandler) UpdateAccessibleNetwork(accountPrivateKey *ecdsa.PrivateKey, gasLimit uint64, value *big.Int, accessibleNetworkInfo *data_types.AccessibleNetworkInfo) (*types.Transaction, error) {

	cnf, err := configs.ReadConfigYAMLFile()

	if err != nil {
		log.Fatal(err)
	}
	var contractAddress = cnf.Platform.Ethereum.AccessibleNetworkSmartContractAddress
	var ethereumURL = cnf.Platform.Ethereum.Url
	conn, err := eth.Dial(ethereumURL)
	if err != nil {
		log.Fatal(err)
		return nil, err
	}

	chainID, err := conn.ChainID(context.Background())

	publicKey := accountPrivateKey.Public()
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

	transOpts, err := bind.NewKeyedTransactorWithChainID(accountPrivateKey, chainID)
	transOpts.Nonce = big.NewInt(int64(nonce))
	transOpts.Value = value       // in wei
	transOpts.GasLimit = gasLimit // in units
	transOpts.GasPrice = gasPrice
	accessibleManager, err := ethereum_handler.NewAccessibleNetworkManagement(common.HexToAddress(contractAddress), conn)
	if err != nil {
		log.Fatalf("Failed to instantiate Storage contract: %v", err)
		return nil, err
	}

	accessibleId, err := stringToBigIntConvertor(accessibleNetworkInfo.AccessibleNetworkId)
	if err != nil {
		return nil, err
	}

	transaction, err := accessibleManager.UpdateAccessibleNetwork(transOpts, accessibleId, accessibleNetworkInfo.NetworkName,
		accessibleNetworkInfo.IP, accessibleNetworkInfo.ADDRESS, accessibleNetworkInfo.CompanyName)

	if err != nil {
		log.Fatalf("Failed to create Permitted Network: %v", err)
		return nil, err
	}
	conn.Close()
	return transaction, nil

}

func (ethAccessibleNetworkHandler *EthAccessibleNetworkHandler) GetAccessibleNetworksNumber(accountPrivateKey *ecdsa.PrivateKey) (*big.Int, error) {
	cnf, err := configs.ReadConfigYAMLFile()

	if err != nil {
		log.Fatal(err)
		return nil, err
	}
	var contractAddress = cnf.Platform.Ethereum.AccessibleNetworkSmartContractAddress
	conn, err_connect := ethAccessibleNetworkHandler.Connect(cnf, err)
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
	accessibleManager, err := ethereum_handler.NewAccessibleNetworkManagement(common.HexToAddress(contractAddress), conn)
	res, err := accessibleManager.GetAccessibleNetworksNumber(CallOpts)
	if err != nil {
		log.Printf("failed to check network existence: %s", err)
		return nil, err
	}
	conn.Close()
	return res, nil
}

func (ethAccessibleNetworkHandler *EthAccessibleNetworkHandler) GetAccessibleNetwork(accountPrivateKey *ecdsa.PrivateKey, accessibleNetworkId string) (*data_types.AccessibleNetworkInfo, error) {
	cnf, err := configs.ReadConfigYAMLFile()

	if err != nil {
		log.Fatal(err)
		return nil, err
	}
	var contractAddress = cnf.Platform.Ethereum.AccessibleNetworkSmartContractAddress
	conn, err_connect := ethAccessibleNetworkHandler.Connect(cnf, err)
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
	accessibleManager, err := ethereum_handler.NewAccessibleNetworkManagement(common.HexToAddress(contractAddress), conn)
	n := new(big.Int)
	accessibleNetIdInt, isSuccess := n.SetString(accessibleNetworkId, 10)
	if !isSuccess {
		log.Fatalf("Failed to convert accessibleNetworkId to big.Int")
	}

	res, err := accessibleManager.GetAccessibleNetwork(CallOpts, accessibleNetIdInt)
	if err != nil {
		log.Printf("failed to check network existence: %s", err)
		return nil, err
	}

	conn.Close()
	return convertContractAccessibleDataType(res), nil
}

func (ethAccessibleNetworkHandler *EthAccessibleNetworkHandler) DeleteAccessibleNetwork(accountPrivateKey *ecdsa.PrivateKey, accessibleNetworkId string, gasLimit uint64, value *big.Int) (*types.Transaction, error) {
	cnf, err := configs.ReadConfigYAMLFile()

	if err != nil {
		log.Fatal(err)
		return nil, err
	}
	var contractAddress = cnf.Platform.Ethereum.AccessibleNetworkSmartContractAddress
	conn, err_connect := ethAccessibleNetworkHandler.Connect(cnf, err)
	if err_connect != nil {
		return nil, err_connect
	}

	chainID, err := conn.ChainID(context.Background())

	publicKey := accountPrivateKey.Public()
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

	transOpts, err := bind.NewKeyedTransactorWithChainID(accountPrivateKey, chainID)
	transOpts.Nonce = big.NewInt(int64(nonce))
	transOpts.Value = value       // in wei
	transOpts.GasLimit = gasLimit // in units
	transOpts.GasPrice = gasPrice
	accessibleManager, err := ethereum_handler.NewAccessibleNetworkManagement(common.HexToAddress(contractAddress), conn)
	if err != nil {
		log.Fatalf("Failed to instantiate Storage contract: %v", err)
		return nil, err
	}

	n := new(big.Int)
	accessibleNetIdInt, isSuccess := n.SetString(accessibleNetworkId, 10)
	if !isSuccess {
		log.Fatalf("Failed to convert accessibleNetworkId to big.Int")
	}

	transaction, err := accessibleManager.DeleteAccessibleNetwork(transOpts, accessibleNetIdInt)
	if err != nil {
		log.Fatalf("Failed to create Permitted Network: %v", err)
		return nil, err
	}
	conn.Close()
	return transaction, nil
}

func (ethAccessibleNetworkHandler *EthAccessibleNetworkHandler) Connect(cnf configs.Configuration, err error) (*eth.Client, error) {
	var ethereumURL = cnf.Platform.Ethereum.Url
	conn, err := eth.Dial(ethereumURL)
	if err != nil {
		log.Fatal(err)
		return nil, err
	}
	return conn, nil
}

func (ethAccessibleNetworkHandler *EthAccessibleNetworkHandler) GetTransactionReceipt(txHash common.Hash) (*types.Receipt, error) {
	cnf, err := configs.ReadConfigYAMLFile()

	if err != nil {
		log.Fatal(err)
		return nil, err
	}

	conn, err_connect := ethAccessibleNetworkHandler.Connect(cnf, err)
	if err_connect != nil {
		return nil, err_connect
	}
	result, err := conn.TransactionReceipt(context.Background(), txHash)
	if err != nil {
		log.Fatal(err)
		return nil, err
	}
	return result, nil
}
