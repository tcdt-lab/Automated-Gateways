package permitted_method_handler

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

type EthPermittedMethodHandler struct {
}

func (ethPermittedMethodHandler *EthPermittedMethodHandler) Connect(cnf configs.Configuration, err error) (*eth.Client, error) {
	var ethereumURL = cnf.Platform.Ethereum.Url
	conn, err := eth.Dial(ethereumURL)
	if err != nil {
		log.Fatal(err)
		return nil, err
	}
	return conn, nil
}

func (ethPermittedMethodHandler *EthPermittedMethodHandler) CreatePermittedMethod(privateKey *ecdsa.PrivateKey, name string, chaincode string, inputArgs string, outputArgs string, gasLimit uint64, value *big.Int) (*types.Transaction, error) {

	cnf, err := configs.ReadConfigYAMLFile()
	if err != nil {
		log.Fatal(err)
		return nil, err
	}
	var contractAddress = cnf.Platform.Ethereum.PermittedMethodSmartContractAddress
	conn, err_connect := ethPermittedMethodHandler.Connect(cnf, err)
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
	methodManager, err := ethereum_handler.NewPermittedMethodManagement(common.HexToAddress(contractAddress), conn)
	if err != nil {
		log.Fatalf("Failed to instantiate Storage contract: %v", err)
		return nil, err
	}

	transaction, err := methodManager.CreatePermittedMethod(transOpts, name, chaincode, inputArgs, outputArgs)

	if err != nil {
		log.Fatalf("Failed to create Permitted Network: %v", err)
		return nil, err
	}

	return transaction, nil
}

func (ethPermittedMethodHandler *EthPermittedMethodHandler) GetAllPermittedMethods(privateKey *ecdsa.PrivateKey) ([]*data_types.MethodInfo, error) {

	var permittedMethodInfoList []*data_types.MethodInfo
	cnf, err := configs.ReadConfigYAMLFile()

	if err != nil {
		log.Fatal(err)
		return nil, err
	}
	var contractAddress = cnf.Platform.Ethereum.PermittedMethodSmartContractAddress
	conn, err_connect := ethPermittedMethodHandler.Connect(cnf, err)
	if err_connect != nil {
		return nil, err_connect
	}

	publicKey := privateKey.Public()
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
	permittedManager, err := ethereum_handler.NewPermittedMethodManagement(common.HexToAddress(contractAddress), conn)
	res, err := permittedManager.GetAllPermittedMethods(CallOpts)
	for _, network := range res {
		networkInfo := convertContractPermittedMethodDataType(network)
		if err != nil {
			log.Printf("failed to unmarshal method info: %s", err)
			return nil, err
		}
		permittedMethodInfoList = append(permittedMethodInfoList, networkInfo)
	}
	conn.Close()
	return permittedMethodInfoList, nil
}

func (ethPermittedMethodHandler *EthPermittedMethodHandler) GetPermittedMethodById(privateKey *ecdsa.PrivateKey, id string) (*data_types.MethodInfo, error) {

	cnf, err := configs.ReadConfigYAMLFile()
	if err != nil {
		log.Fatal(err)
		return nil, err
	}
	var contractAddress = cnf.Platform.Ethereum.PermittedMethodSmartContractAddress
	conn, err_connect := ethPermittedMethodHandler.Connect(cnf, err)
	if err_connect != nil {
		return nil, err_connect
	}

	publicKey := privateKey.Public()
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
	permittedManager, err := ethereum_handler.NewPermittedMethodManagement(common.HexToAddress(contractAddress), conn)
	intId, err := stringToBigIntConvertor(id)

	if err != nil {
		log.Fatalf("Failed to convert accessibleNetworkId to big.Int")
		return nil, errors.New("Failed to convert accessibleNetworkId to big.Int")
	}

	res, err := permittedManager.GetPermittedMethodById(CallOpts, intId)
	if err != nil {
		log.Printf("failed to unmarshal method info: %s", err)
		return nil, err
	}

	networkInfo := convertContractPermittedMethodDataType(res)
	conn.Close()
	return networkInfo, nil

}

func (ethPermittedMethodHandler *EthPermittedMethodHandler) DeletePermittedMethod(privateKey *ecdsa.PrivateKey, id string, gasLimit uint64, value *big.Int) (*types.Transaction, error) {

	cnf, err := configs.ReadConfigYAMLFile()
	if err != nil {
		log.Fatal(err)
		return nil, err
	}
	var contractAddress = cnf.Platform.Ethereum.PermittedMethodSmartContractAddress
	conn, err_connect := ethPermittedMethodHandler.Connect(cnf, err)
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
	methodManager, err := ethereum_handler.NewPermittedMethodManagement(common.HexToAddress(contractAddress), conn)
	intId, err := stringToBigIntConvertor(id)

	if err != nil {
		log.Fatalf("Failed to convert accessibleNetworkId to big.Int")
		return nil, errors.New("Failed to convert accessibleNetworkId to big.Int")

	}

	transaction, err := methodManager.DeletePermittedMethod(transOpts, intId)
	if err != nil {
		log.Fatalf("Failed to delete Permitted Network: %v", err)
		return nil, err

	}

	return transaction, nil
}

func (ethPermittedMethodHandler *EthPermittedMethodHandler) UpdatePermittedMethod(privateKey *ecdsa.PrivateKey, id string, name string, chaincode string, inputArgs string, outputArgs string, gasLimit uint64, value *big.Int) (*types.Transaction, error) {

	cnf, err := configs.ReadConfigYAMLFile()
	if err != nil {
		log.Fatal(err)
		return nil, err
	}
	var contractAddress = cnf.Platform.Ethereum.PermittedMethodSmartContractAddress
	conn, err_connect := ethPermittedMethodHandler.Connect(cnf, err)
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
	methodManager, err := ethereum_handler.NewPermittedMethodManagement(common.HexToAddress(contractAddress), conn)
	intId, err := stringToBigIntConvertor(id)

	if err != nil {
		log.Fatalf("Failed to convert accessibleNetworkId to big.Int")
		return nil, errors.New("Failed to convert accessibleNetworkId to big.Int")
	}

	transaction, err := methodManager.UpdatePermittedMethod(transOpts, intId, name, chaincode, inputArgs, outputArgs)
	if err != nil {
		log.Fatalf("Failed to update Permitted Network: %v", err)
		return nil, err
	}

	return transaction, nil

}

func (ethPermittedMethodHandler *EthPermittedMethodHandler) GetTransactionReceipt(hash common.Hash) (*types.Receipt, error) {
	cnf, err := configs.ReadConfigYAMLFile()
	conn, err_connect := ethPermittedMethodHandler.Connect(cnf, err)
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

func (ethPermittedMethodHandler *EthPermittedMethodHandler) PermittedMethodExist(privateKey *ecdsa.PrivateKey, id string) (bool, error) {
	cnf, err := configs.ReadConfigYAMLFile()
	if err != nil {
		log.Fatal(err)
		return false, err
	}
	var contractAddress = cnf.Platform.Ethereum.PermittedMethodSmartContractAddress
	conn, err_connect := ethPermittedMethodHandler.Connect(cnf, err)
	if err_connect != nil {
		return false, err_connect
	}

	publicKey := privateKey.Public()
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
	permittedManager, err := ethereum_handler.NewPermittedMethodManagement(common.HexToAddress(contractAddress), conn)
	intId, err := stringToBigIntConvertor(id)

	if err != nil {
		log.Fatalf("Failed to convert accessibleNetworkId to big.Int")
		return false, errors.New("Failed to convert accessibleNetworkId to big.Int")
	}

	res, err := permittedManager.IsPermittedMethodExist(CallOpts, intId)
	if err != nil {
		log.Printf("failed to unmarshal method info: %s", err)
		return false, err
	}

	conn.Close()
	return res, nil
}
