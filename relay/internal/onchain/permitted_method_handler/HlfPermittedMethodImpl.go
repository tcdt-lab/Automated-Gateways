package permitted_method_handler

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

var channelName string
var chaincodeName string
var cnf configs.Configuration

type HlfPermittedMethod struct {
}

func (hlfPermittedMethod *HlfPermittedMethod) OpenConnection() (*grpc.ClientConn, *client.Gateway, error) {
	configYaml, errYaml := configs.ReadConfigYAMLFile()
	cnf = configYaml
	if errYaml != nil {
		return nil, nil, errYaml
	}
	chaincodeName = cnf.Platform.Hyperledger.PermittedMethodChaincodeName
	channelName = cnf.Platform.Hyperledger.PermittedNetworkChannelName
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

func (hlfPermittedMethod *HlfPermittedMethod) AddPermittedMethod(gw *client.Gateway, permittedNetworkId string, permittedMethodName string, chaincode string, channel string, inputArgs string, outputArgs string) (*datatypes.MethodInfo, error) {

	log.Printf("add method %s to  permitted Network with Id: %s\n", permittedMethodName, permittedNetworkId)
	methodName := "AddPermittedMethod"

	if ccname := os.Getenv("CHAINCODE_NAME"); ccname != "" {
		chaincodeName = ccname
	}

	if cname := os.Getenv("CHANNEL_NAME"); cname != "" {
		channelName = cname
	}
	network := gw.GetNetwork(channelName)
	contract := network.GetContract(chaincodeName)

	submitRes, err := contract.SubmitTransaction(methodName, permittedNetworkId, permittedMethodName, chaincode, channel, inputArgs, outputArgs)
	var permittedMethodInfo datatypes.MethodInfo
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

func (hlfPermittedMethod *HlfPermittedMethod) GetPermittedMethod(gw *client.Gateway, permittedMethodId string) (*datatypes.MethodInfo, error) {
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
	var permittedMethodInfo datatypes.MethodInfo
	err = json.Unmarshal(evalRes, &permittedMethodInfo)
	if err != nil {
		log.Printf("failed to submit transaction: %s", err)
		return nil, err
	}

	log.Printf("Result:%s\n", string(evalRes))
	return &permittedMethodInfo, nil
}

func (hlfPermittedMethod HlfPermittedMethod) GetPermittedMethodsByIndex(gw *client.Gateway, startIndex string, endIndex string, permittedNetworkId string) ([]*datatypes.MethodInfo, error) {
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

	var permittedMethodInfo []*datatypes.MethodInfo
	err = json.Unmarshal(evalRes, &permittedMethodInfo)
	if err != nil {
		log.Printf("failed to submit transaction: %s", err)
		return nil, err
	}

	log.Printf("Result:%s\n", string(evalRes))
	return permittedMethodInfo, nil
}
func (hlfPermittedMethod *HlfPermittedMethod) GetPermittedMethodsByNetworkId(gw *client.Gateway, permittedNetworkId string) ([]*datatypes.MethodInfo, error) {
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

	var permittedMethodInfo []*datatypes.MethodInfo
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
func (hlfPermittedMethod *HlfPermittedMethod) InvokePermittedMethod(gw *client.Gateway, name string, chaincode string, channel string, inputArgs []string) (*string, error) {

	methodName := name

	if ccname := os.Getenv("CHAINCODE_NAME"); ccname != "" {
		chaincodeName = ccname
	}

	if cname := os.Getenv("CHANNEL_NAME"); cname != "" {
		channelName = cname
	}
	network := gw.GetNetwork(channel)
	contract := network.GetContract(chaincode)
	submitRes, err := contract.SubmitTransaction(methodName, inputArgs...)
	if err != nil {
		log.Printf("failed to submit transaction: %s", err)
		return nil, err
	}
	strRes := string(submitRes)
	log.Printf("Result:%s\n", string(strRes))
	return &strRes, nil
}
