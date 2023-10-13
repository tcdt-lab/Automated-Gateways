package permitted_method_handler

import (
	"crypto/x509"
	"encoding/json"
	"fmt"
	"github.com/hyperledger/fabric-gateway/pkg/client"
	"github.com/hyperledger/fabric-gateway/pkg/identity"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials"
	"log"
	"os"
	hlfConfig "relay/src/onchain/configs"
	"time"
)

type PermittedMethodInfo struct {
	PermittedMethodId string `json:"PermittedMethodId"`
	Name              string `json:"Name"`
	Chaincode         string `json:"Chaincode"`
	Channel           string `json:"Channel"`
	InputArgs         string `json:"inputArgs"`
	OutputType        string `json:"outputArgs"`
}

var channelName = "mychannel"
var chaincodeName = "permitted_method"

type HlfPermittedMethod struct {
}

func (hlfPermittedMethod *HlfPermittedMethod) OpenConnection() (*grpc.ClientConn, *client.Gateway, error) {
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
	certificate, err := loadCertificate(hlfConfig.TlsCertPath)
	if err != nil {
		panic(err)
	}

	certPool := x509.NewCertPool()
	certPool.AddCert(certificate)
	transportCredentials := credentials.NewClientTLSFromCert(certPool, hlfConfig.GatewayPeer)

	connection, err := grpc.Dial(hlfConfig.PeerEndpoint, grpc.WithTransportCredentials(transportCredentials))
	if err != nil {
		panic(fmt.Errorf("failed to create gRPC connection: %w", err))
	}

	return connection
}
func newSign() identity.Sign {

	privateKeyPEM, err := os.ReadFile(hlfConfig.KeyPath)

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
	certificate, err := loadCertificate(hlfConfig.CertPath)
	if err != nil {
		panic(err)
	}

	id, err := identity.NewX509Identity(hlfConfig.MspID, certificate)
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

func (hlfPermittedMethod *HlfPermittedMethod) CloseConnection(clientConnection *grpc.ClientConn, gw *client.Gateway) error {
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

func (hlfPermittedMethod *HlfPermittedMethod) AddPermittedMethod(gw *client.Gateway, permittedNetworkId string, permittedNetworkName string, chaincode string, channel string, inputArgs string, outputArgs string) (*PermittedMethodInfo, error) {

	log.Printf("add method %s to  permitted Network with Id: %s\n", permittedNetworkName, permittedNetworkId)
	methodName := "AddPermittedMethod"

	if ccname := os.Getenv("CHAINCODE_NAME"); ccname != "" {
		chaincodeName = ccname
	}

	if cname := os.Getenv("CHANNEL_NAME"); cname != "" {
		channelName = cname
	}
	network := gw.GetNetwork(channelName)
	contract := network.GetContract(chaincodeName)

	submitRes, err := contract.SubmitTransaction(methodName, permittedNetworkId, permittedNetworkName, chaincode, channel, inputArgs, outputArgs)
	var permittedMethodInfo PermittedMethodInfo
	err = json.Unmarshal(submitRes, &permittedMethodInfo)
	if err != nil {
		log.Printf("failed to submit transaction: %s", err)
		return nil, err
	}

	log.Printf("Result:%s\n", string(submitRes))
	return &permittedMethodInfo, nil
}

func (hlfPermittedMethod *HlfPermittedMethod) PermittedMethodExists(gw *client.Gateway, id string) (bool, error) {
	log.Printf("PermittedMethodExists: %s\n", id)
	methodName := "PermittedMethodExists"

	if ccname := os.Getenv("CHAINCODE_NAME"); ccname != "" {
		chaincodeName = ccname
	}

	if cname := os.Getenv("CHANNEL_NAME"); cname != "" {
		channelName = cname
	}
	network := gw.GetNetwork(channelName)
	contract := network.GetContract(chaincodeName)

	evalRes, err := contract.EvaluateTransaction(methodName, id)
	var resBoolean bool
	err = json.Unmarshal(evalRes, &resBoolean)

	if err != nil {
		log.Printf("failed to submit transaction: %s", err)
		return false, err
	}

	log.Printf("Result:%t \n", resBoolean)
	return resBoolean, nil
}

func (hlfPermittedMethod *HlfPermittedMethod) GetPermittedMethod(gw *client.Gateway, permittedMethodId string) (*PermittedMethodInfo, error) {
	log.Printf("Get PermittedMethod: %s\n", permittedMethodId)
	methodName := "GetPermittedMethod"

	if ccname := os.Getenv("CHAINCODE_NAME"); ccname != "" {
		chaincodeName = ccname
	}

	if cname := os.Getenv("CHANNEL_NAME"); cname != "" {
		channelName = cname
	}
	network := gw.GetNetwork(channelName)
	contract := network.GetContract(chaincodeName)

	evalRes, err := contract.EvaluateTransaction(methodName, permittedMethodId)

	if evalRes == nil {
		return nil, fmt.Errorf("the accessible network %s does not exist", permittedMethodId)
	}
	log.Printf("The result of Get PermittedNetwork: %s\n", string(evalRes))
	var permittedMethodInfo PermittedMethodInfo
	err = json.Unmarshal(evalRes, &permittedMethodInfo)
	if err != nil {
		log.Printf("failed to submit transaction: %s", err)
		return nil, err
	}

	log.Printf("Result:%s\n", string(evalRes))
	return &permittedMethodInfo, nil
}

func (hlfPermittedMethod HlfPermittedMethod) GetPermittedMethodsByIndex(gw *client.Gateway, permittedNetworkId string, startIndex string, endIndex string) ([]*PermittedMethodInfo, error) {
	log.Printf("Get PermittedMethods by index: %s\n", permittedNetworkId)
	methodName := "GetPermittedMethodsByIndex"

	if ccname := os.Getenv("CHAINCODE_NAME"); ccname != "" {
		chaincodeName = ccname
	}

	if cname := os.Getenv("CHANNEL_NAME"); cname != "" {
		channelName = cname
	}
	network := gw.GetNetwork(channelName)
	contract := network.GetContract(chaincodeName)

	evalRes, err := contract.EvaluateTransaction(methodName, permittedNetworkId, startIndex, endIndex)
	if err != nil {
		log.Printf("failed to submit transaction: %s", err)
		return nil, err
	}

	var permittedMethodInfo []*PermittedMethodInfo
	err = json.Unmarshal(evalRes, &permittedMethodInfo)
	if err != nil {
		log.Printf("failed to submit transaction: %s", err)
		return nil, err
	}

	log.Printf("Result:%s\n", string(evalRes))
	return permittedMethodInfo, nil
}
func (hlfPermittedMethod *HlfPermittedMethod) GetPermittedMethodsByNetworkId(gw *client.Gateway, permittedNetworkId string) ([]*PermittedMethodInfo, error) {
	log.Printf("Get PermittedMethods: %s\n", permittedNetworkId)
	methodName := "GetPermittedMethodsByNetworkId"

	if ccname := os.Getenv("CHAINCODE_NAME"); ccname != "" {
		chaincodeName = ccname
	}

	if cname := os.Getenv("CHANNEL_NAME"); cname != "" {
		channelName = cname
	}
	network := gw.GetNetwork(channelName)
	contract := network.GetContract(chaincodeName)

	evalRes, err := contract.EvaluateTransaction(methodName, permittedNetworkId)
	if err != nil {
		log.Printf("failed to submit transaction: %s", err)
		return nil, err
	}

	var permittedMethodInfo []*PermittedMethodInfo
	err = json.Unmarshal(evalRes, &permittedMethodInfo)
	if err != nil {
		log.Printf("failed to submit transaction: %s", err)
		return nil, err
	}

	log.Printf("Result:%s\n", string(evalRes))
	return permittedMethodInfo, nil
}

func (hlfPermittedMethod *HlfPermittedMethod) RemovePermittedMethod(gw *client.Gateway, id string) error {
	log.Printf("Remove PermittedMethod: %s\n", id)
	methodName := "RemovePermittedMethod"

	if ccname := os.Getenv("CHAINCODE_NAME"); ccname != "" {
		chaincodeName = ccname
	}

	if cname := os.Getenv("CHANNEL_NAME"); cname != "" {
		channelName = cname
	}
	network := gw.GetNetwork(channelName)
	contract := network.GetContract(chaincodeName)

	evalRes, err := contract.SubmitTransaction(methodName, id)
	if err != nil {
		log.Printf("failed to submit transaction: %s", err)
		return err
	}

	log.Printf("Result:%s\n", string(evalRes))
	return nil
}
func (hlfPermittedMethod *HlfPermittedMethod) UpdatePermittedMethod(gw *client.Gateway, id string, name string, chaincode string, channel string, inputArgs string, outputArgs string) error {
	log.Printf("Update PermittedMethod: %s\n", id)
	methodName := "UpdatePermittedMethod"

	if ccname := os.Getenv("CHAINCODE_NAME"); ccname != "" {
		chaincodeName = ccname
	}

	if cname := os.Getenv("CHANNEL_NAME"); cname != "" {
		channelName = cname
	}
	network := gw.GetNetwork(channelName)
	contract := network.GetContract(chaincodeName)
	submitRes, err := contract.SubmitTransaction(methodName, id, name, chaincode, channel, inputArgs, outputArgs)
	if err != nil {
		log.Printf("failed to submit transaction: %s", err)
		return err
	}
	log.Printf("Result:%s\n", string(submitRes))
	return nil
}
