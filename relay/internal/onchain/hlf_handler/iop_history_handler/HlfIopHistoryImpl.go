package iop_history_handler

import (
	"crypto/x509"
	"encoding/json"
	"fmt"
	"github.com/hyperledger/fabric-gateway/pkg/client"
	"github.com/hyperledger/fabric-gateway/pkg/identity"
	"github.com/tcdt-lab/Automated-Gateways/relay/configs"
	datatypes "github.com/tcdt-lab/Automated-Gateways/relay/data_types"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials"
	"log"
	"os"
	"time"
)

type HlfIopHistory struct {
}

var cnf configs.Configuration
var chaincodeName string
var channelName string

func (hlfIopHistory *HlfIopHistory) OpenConnection() (*grpc.ClientConn, *client.Gateway, error) {

	configYaml, errYaml := configs.ReadConfigYAMLFile()
	cnf = configYaml
	if errYaml != nil {
		return nil, nil, errYaml
	}
	chaincodeName = cnf.Platform.Hyperledger.IopHistoryChaincodeName
	channelName = cnf.Platform.Hyperledger.IopHistoryChannelName
	clientConnection := newGrpcConnection()
	id := newIdentity()
	sign := newSign()
	var err error
	gw, err := client.Connect(
		id,
		client.WithSign(sign),
		client.WithClientConnection(clientConnection),
		// Default timeouts for different gRPC calls
		client.WithEvaluateTimeout(5*time.Second),
		client.WithEndorseTimeout(15*time.Second),
		client.WithSubmitTimeout(5*time.Second),
		client.WithCommitStatusTimeout(1*time.Minute),
	)
	fmt.Print(*gw)
	if err != nil {
		panic(err)
		return nil, nil, err
	}
	return clientConnection, gw, nil
}

func newGrpcConnection() *grpc.ClientConn {
	certificate, err := loadCertificate(cnf.Platform.Hyperledger.TlsCertPath)
	if err != nil {
		panic(err)
	}

	certPool := x509.NewCertPool()
	certPool.AddCert(certificate)
	transportCredentials := credentials.NewClientTLSFromCert(certPool, cnf.Platform.Hyperledger.GatewayPeer)

	connection, err := grpc.Dial(cnf.Platform.Hyperledger.PeerEndpoint, grpc.WithTransportCredentials(transportCredentials))
	if err != nil {
		panic(fmt.Errorf("failed to create gRPC connection: %w", err))
	}

	return connection
}
func newSign() identity.Sign {

	privateKeyPEM, err := os.ReadFile(cnf.Platform.Hyperledger.KeyPath)

	if err != nil {
		panic(fmt.Errorf("failed to read private key file: %w", err))
	}

	privateKey, err := identity.PrivateKeyFromPEM(privateKeyPEM)
	if err != nil {
		panic(err)
	}

	sign, err := identity.NewPrivateKeySign(privateKey)
	if err != nil {
		panic(err)
	}

	return sign
}

func newIdentity() *identity.X509Identity {
	certificate, err := loadCertificate(cnf.Platform.Hyperledger.CertPath)
	if err != nil {
		panic(err)
	}

	id, err := identity.NewX509Identity(cnf.Platform.Hyperledger.MspID, certificate)
	if err != nil {
		panic(err)
	}

	return id
}
func loadCertificate(filename string) (*x509.Certificate, error) {
	certificatePEM, err := os.ReadFile(filename)
	if err != nil {
		return nil, fmt.Errorf("failed to read certificate file: %w", err)
	}
	return identity.CertificateFromPEM(certificatePEM)
}

func (hlfIopHistory *HlfIopHistory) CloseConnection(clientConnection *grpc.ClientConn, gw *client.Gateway) error {
	err := gw.Close()
	if err != nil {
		return err
	}

	errConnection := clientConnection.Close()
	if errConnection != nil {
		return errConnection
	}
	return nil
}

func (hlfIopHistory *HlfIopHistory) InsertHistoryLog(gateway *client.Gateway, callType string, callerIP string, callerAddress string, methodName string, inputArgs string, outputArgs string) (*datatypes.HistoryLog, error) {
	log.Printf("Insert History Log:")
	var chaincodeMethodName string
	if callType == datatypes.HISTORY_TYPE_INSIDE_INITIATED {
		chaincodeMethodName = "InsertInsideInitiatedCall"
	} else {
		chaincodeMethodName = "InsertOutsideInitiatedCall"
	}

	if ccname := os.Getenv("CHAINCODE_NAME"); ccname != "" {
		chaincodeName = ccname
	}

	if cname := os.Getenv("CHANNEL_NAME"); cname != "" {
		channelName = cname
	}
	network := gateway.GetNetwork(channelName)
	contract := network.GetContract(chaincodeName)
	submitRes, err := contract.SubmitTransaction(chaincodeMethodName, callerIP, callerAddress, methodName, inputArgs, outputArgs)
	if err != nil {
		log.Printf("failed to submit transaction: %s", err)
		return nil, err
	}

	var historyLog datatypes.HistoryLog
	err = json.Unmarshal(submitRes, &historyLog)
	if err != nil {
		log.Printf("failed to submit transaction: %s", err)
		return nil, err
	}

	log.Printf("Result:%s\n", string(submitRes))
	return &historyLog, nil

}

func (hlfIopHistory *HlfIopHistory) GetHistoryBasedOnCallType(gateway *client.Gateway, callType string) ([]*datatypes.HistoryLog, error) {
	log.Printf("Get History Based On Call Type: %s\n", callType)
	chaincodeMethodName := "GetHistoryBasedOnCallType"

	if ccname := os.Getenv("CHAINCODE_NAME"); ccname != "" {
		chaincodeName = ccname
	}

	if cname := os.Getenv("CHANNEL_NAME"); cname != "" {
		channelName = cname
	}
	network := gateway.GetNetwork(channelName)
	contract := network.GetContract(chaincodeName)

	evalRes, err := contract.EvaluateTransaction(chaincodeMethodName, callType)
	var res []*datatypes.HistoryLog
	err = json.Unmarshal(evalRes, &res)
	if err != nil {
		log.Printf("failed to submit transaction: %s", err)
		return nil, err
	}
	return res, nil
}

func (hlfIopHistory *HlfIopHistory) GetHistoryBasedOnCallerIP(gateway *client.Gateway, callType string, callerIP string) ([]*datatypes.HistoryLog, error) {

	log.Printf("Get History Based On Caller IP: %s\n", callerIP)
	chaincodeMethodName := "GetHistoryBasedOnCallerIP"

	if ccname := os.Getenv("CHAINCODE_NAME"); ccname != "" {
		chaincodeName = ccname
	}

	if cname := os.Getenv("CHANNEL_NAME"); cname != "" {
		channelName = cname
	}
	network := gateway.GetNetwork(channelName)
	contract := network.GetContract(chaincodeName)

	evalRes, err := contract.EvaluateTransaction(chaincodeMethodName, callType, callerIP)
	var res []*datatypes.HistoryLog
	err = json.Unmarshal(evalRes, &res)
	if err != nil {
		log.Printf("failed to submit transaction: %s", err)
		return nil, err
	}
	return res, nil
}

func (hlfIopHistory *HlfIopHistory) GetHistoryBasedOnCallerAddress(gateway *client.Gateway, callType string, callerIP string, callerAddress string) ([]*datatypes.HistoryLog, error) {

	log.Printf("Get History Based On Caller Address: %s\n", callerAddress)
	chaincodeMethodName := "GetHistoryBasedOnCallerAddress"

	if ccname := os.Getenv("CHAINCODE_NAME"); ccname != "" {
		chaincodeName = ccname
	}

	if cname := os.Getenv("CHANNEL_NAME"); cname != "" {
		channelName = cname
	}
	network := gateway.GetNetwork(channelName)
	contract := network.GetContract(chaincodeName)

	evalRes, err := contract.EvaluateTransaction(chaincodeMethodName, callType, callerIP, callerAddress)
	var res []*datatypes.HistoryLog
	err = json.Unmarshal(evalRes, &res)
	if err != nil {
		log.Printf("failed to submit transaction: %s", err)
		return nil, err
	}
	return res, nil
}

func (hlfIopHistory *HlfIopHistory) GetHistoryBasedOnMethodName(gateway *client.Gateway, callType string, callerIP string, callerAddress string, methodName string) ([]*datatypes.HistoryLog, error) {

	log.Printf("Get History Based On Method Name: %s\n", methodName)
	chaincodeMethodName := "GetHistoryBasedOnMethodName"

	if ccname := os.Getenv("CHAINCODE_NAME"); ccname != "" {
		chaincodeName = ccname
	}

	if cname := os.Getenv("CHANNEL_NAME"); cname != "" {
		channelName = cname
	}
	network := gateway.GetNetwork(channelName)
	contract := network.GetContract(chaincodeName)

	evalRes, err := contract.EvaluateTransaction(chaincodeMethodName, callType, callerIP, callerAddress, methodName)
	var res []*datatypes.HistoryLog
	err = json.Unmarshal(evalRes, &res)
	if err != nil {
		log.Printf("failed to submit transaction: %s", err)
		return nil, err
	}
	return res, nil
}

func (hlfIopHistory *HlfIopHistory) GetHistoryBasedOnDate(gateway *client.Gateway, callerType string, callerIP string, callerAddress string, methodName string, year string, month string, day string) ([]*datatypes.HistoryLog, error) {

	log.Printf("Get History Based On Date: %s\n", year)
	chaincodeMethodName := "GetHistoryBasedOnDate"

	if ccname := os.Getenv("CHAINCODE_NAME"); ccname != "" {
		chaincodeName = ccname
	}

	if cname := os.Getenv("CHANNEL_NAME"); cname != "" {
		channelName = cname
	}
	network := gateway.GetNetwork(channelName)
	contract := network.GetContract(chaincodeName)

	evalRes, err := contract.EvaluateTransaction(chaincodeMethodName, callerType, callerIP, callerAddress, methodName, year, month, day)
	var res []*datatypes.HistoryLog
	err = json.Unmarshal(evalRes, &res)
	if err != nil {
		log.Printf("failed to submit transaction: %s", err)
		return nil, err
	}
	return res, nil
}

func (hlfIopHistory *HlfIopHistory) GetHistoryBasedOnYear(gateway *client.Gateway, callerType string, callerIP string, callerAddress string, methodName string, year string) ([]datatypes.HistoryLog, error) {

	log.Printf("Get History Based On Year: %s\n", year)
	chaincodeMethodName := "GetHistoryBasedOnYear"

	if ccname := os.Getenv("CHAINCODE_NAME"); ccname != "" {
		chaincodeName = ccname
	}

	if cname := os.Getenv("CHANNEL_NAME"); cname != "" {
		channelName = cname
	}
	network := gateway.GetNetwork(channelName)
	contract := network.GetContract(chaincodeName)

	evalRes, err := contract.EvaluateTransaction(chaincodeMethodName, callerType, callerIP, callerAddress, methodName, year)
	var res []datatypes.HistoryLog
	err = json.Unmarshal(evalRes, &res)
	if err != nil {
		log.Printf("failed to submit transaction: %s", err)
		return nil, err
	}
	return res, nil
}

func (hlfIopHistory *HlfIopHistory) GetHistoryBasedOnMonth(gateway *client.Gateway, callerType string, callerIP string, callerAddress string, methodName string, year string, month string) ([]datatypes.HistoryLog, error) {

	log.Printf("Get History Based On Month: %s\n", month)
	chaincodeMethodName := "GetHistoryBasedOnMonth"

	if ccname := os.Getenv("CHAINCODE_NAME"); ccname != "" {
		chaincodeName = ccname
	}

	if cname := os.Getenv("CHANNEL_NAME"); cname != "" {
		channelName = cname
	}
	network := gateway.GetNetwork(channelName)
	contract := network.GetContract(chaincodeName)

	evalRes, err := contract.EvaluateTransaction(chaincodeMethodName, callerType, callerIP, callerAddress, methodName, year, month)
	var res []datatypes.HistoryLog
	err = json.Unmarshal(evalRes, &res)
	if err != nil {
		log.Printf("failed to submit transaction: %s", err)
		return nil, err
	}
	return res, nil

}

func (hlfIopHistory *HlfIopHistory) GetHistoryBasedOnHour(gateway *client.Gateway, callerType string, callerIP string, callerAddress string, methodName string, year string, month string, day string, hour string) ([]datatypes.HistoryLog, error) {

	log.Printf("Get History Based On Hour: %s\n", hour)
	chaincodeMethodName := "GetHistoryBasedOnHour"

	if ccname := os.Getenv("CHAINCODE_NAME"); ccname != "" {
		chaincodeName = ccname
	}

	if cname := os.Getenv("CHANNEL_NAME"); cname != "" {
		channelName = cname
	}
	network := gateway.GetNetwork(channelName)
	contract := network.GetContract(chaincodeName)

	evalRes, err := contract.EvaluateTransaction(chaincodeMethodName, callerType, callerIP, callerAddress, methodName, year, month, day, hour)
	var res []datatypes.HistoryLog
	err = json.Unmarshal(evalRes, &res)
	if err != nil {
		log.Printf("failed to submit transaction: %s", err)
		return nil, err
	}
	return res, nil

}

func (hlfIopHistory *HlfIopHistory) RemoveHistoryLog(gateway *client.Gateway, callerType string, callerIP string, callerAddress string, methodName string, year string, month string, day string, hour string, distinctive string) error {

	log.Printf("Removing History Log: %s\n", distinctive)
	chaincodeMethodName := "RemoveAHistoryLog"

	if ccname := os.Getenv("CHAINCODE_NAME"); ccname != "" {
		chaincodeName = ccname
	}

	if cname := os.Getenv("CHANNEL_NAME"); cname != "" {
		channelName = cname
	}
	network := gateway.GetNetwork(channelName)
	contract := network.GetContract(chaincodeName)

	submitRes, err := contract.SubmitTransaction(chaincodeMethodName, callerType, callerIP, callerAddress, methodName, year, month, day, hour, distinctive)
	if err != nil {
		log.Printf("failed to submit remove transaction: %s", err)
		return err
	}

	log.Printf("Result:%s\n", string(submitRes))
	return nil
}
