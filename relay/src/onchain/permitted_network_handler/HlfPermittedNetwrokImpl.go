package permitted_network_handler

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

type PermittedNetworkInfo struct {
	NetworkName        string `json:"NetworkName"`
	IP                 string `json:"IP"`
	ADDRESS            string `json:"ADDRESS"`
	CompanyName        string `json:"CompanyName"`
	PermittedNetworkId string `json:"PermittedNetworkId"`
}

var channelName = "mychannel"
var chaincodeName = "permitted_net"

type HlfPermittedNetwork struct {
}

func (hlfPermittedNetwork *HlfPermittedNetwork) OpenConnection() (*grpc.ClientConn, *client.Gateway, error) {
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

func (hlfPermittedNetwork *HlfPermittedNetwork) CloseConnection(clientConnection *grpc.ClientConn, gw *client.Gateway) error {
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
func (hlfPermittedNetwork *HlfPermittedNetwork) CreatePermittedNetwork(gateway *client.Gateway, networkName string, ip string, address string, companyName string) (*PermittedNetworkInfo, error) {
	log.Printf("Creating Permitted Network: %s\n", networkName)
	methodName := "CreatePermittedNetwork"

	if ccname := os.Getenv("CHAINCODE_NAME"); ccname != "" {
		chaincodeName = ccname
	}

	if cname := os.Getenv("CHANNEL_NAME"); cname != "" {
		channelName = cname
	}
	network := gateway.GetNetwork(channelName)
	contract := network.GetContract(chaincodeName)

	submitRes, err := contract.SubmitTransaction(methodName, networkName, ip, address, companyName)
	var permittedNetworkInfo PermittedNetworkInfo
	err = json.Unmarshal(submitRes, &permittedNetworkInfo)
	if err != nil {
		log.Printf("failed to submit transaction: %s", err)
		return nil, err
	}

	log.Printf("Result:%s\n", string(submitRes))
	return &permittedNetworkInfo, nil
}

func (hlfPermittedNetwork *HlfPermittedNetwork) PermittedNetworkExists(gateway *client.Gateway, id string) (bool, error) {
	log.Printf("PermittedNetworkExists: %s\n", id)
	methodName := "PermittedNetworkExists"

	if ccname := os.Getenv("CHAINCODE_NAME"); ccname != "" {
		chaincodeName = ccname
	}

	if cname := os.Getenv("CHANNEL_NAME"); cname != "" {
		channelName = cname
	}
	network := gateway.GetNetwork(channelName)
	contract := network.GetContract(chaincodeName)

	queryResponse, err := contract.EvaluateTransaction(methodName, id)
	var resBoolean bool
	err = json.Unmarshal(queryResponse, &resBoolean)
	if err != nil {
		log.Printf("failed to submit transaction: %s", err)
		return false, err
	}

	log.Printf("Result:%s\n", string(queryResponse))
	return resBoolean, nil
}

func (hlfPermittedNetwork *HlfPermittedNetwork) RemovePermittedNetwork(gateway *client.Gateway, id string) error {
	log.Printf("Removing PermittedNetwork: %s\n", id)
	methodName := "RemovePermittedNetwork"

	if ccname := os.Getenv("CHAINCODE_NAME"); ccname != "" {
		chaincodeName = ccname
	}

	if cname := os.Getenv("CHANNEL_NAME"); cname != "" {
		channelName = cname
	}
	network := gateway.GetNetwork(channelName)
	contract := network.GetContract(chaincodeName)

	submitRes, err := contract.SubmitTransaction(methodName, id)
	if err != nil {
		log.Printf("failed to submit remove transaction: %s", err)
		return err
	}

	log.Printf("Result:%s\n", string(submitRes))
	return nil
}

func (hlfPermittedNetwork *HlfPermittedNetwork) UpdatePermittedNetwork(gateway *client.Gateway, id string, networkName string, ip string, address string, companyName string) error {
	log.Printf("Update PermittedNetwork: %s\n", id)
	methodName := "UpdatePermittedNetwork"

	if ccname := os.Getenv("CHAINCODE_NAME"); ccname != "" {
		chaincodeName = ccname
	}

	if cname := os.Getenv("CHANNEL_NAME"); cname != "" {
		channelName = cname
	}
	network := gateway.GetNetwork(channelName)
	contract := network.GetContract(chaincodeName)

	submitRes, err := contract.SubmitTransaction(methodName, id, networkName, ip, address, companyName)
	if err != nil {
		log.Printf("failed to submit update transaction: %s", err)
		return err
	}

	log.Printf("Result:%s\n", string(submitRes))
	return nil
}
func (hlfPermittedNetwork *HlfPermittedNetwork) GetPermittedNetwork(gateway *client.Gateway, id string) (*PermittedNetworkInfo, error) {
	log.Printf("Get PermittedNetwork: %s\n", id)
	methodName := "GetPermittedNetwork"

	if ccname := os.Getenv("CHAINCODE_NAME"); ccname != "" {
		chaincodeName = ccname
	}

	if cname := os.Getenv("CHANNEL_NAME"); cname != "" {
		channelName = cname
	}
	network := gateway.GetNetwork(channelName)
	contract := network.GetContract(chaincodeName)

	queryResponse, err := contract.EvaluateTransaction(methodName, id)
	var permittedNetworkInfo PermittedNetworkInfo
	err = json.Unmarshal(queryResponse, &permittedNetworkInfo)
	if err != nil {
		log.Printf("failed to submit transaction: %s", err)
		return nil, err
	}

	log.Printf("Result:%s\n", string(queryResponse))
	return &permittedNetworkInfo, nil
}

func (hlfPermittedNetwork *HlfPermittedNetwork) GetPermittedNetworksByAddress(gateway *client.Gateway, address string) ([]*PermittedNetworkInfo, error) {
	log.Printf("Get PermittedNetworks: \n")
	methodName := "GetPermittedNetworks"

	if ccname := os.Getenv("CHAINCODE_NAME"); ccname != "" {
		chaincodeName = ccname
	}

	if cname := os.Getenv("CHANNEL_NAME"); cname != "" {
		channelName = cname
	}
	network := gateway.GetNetwork(channelName)
	contract := network.GetContract(chaincodeName)

	queryResponse, err := contract.EvaluateTransaction(methodName, address)
	var results []*PermittedNetworkInfo
	err = json.Unmarshal(queryResponse, &results)
	if err != nil {
		log.Printf("failed to submit transaction: %s", err)
		return nil, err
	}

	log.Printf("Result:%s\n", string(queryResponse))
	return results, nil
}

func (hlfPermittedNetwork *HlfPermittedNetwork) GetPermittedNetworkByIndexAndAddress(gateway *client.Gateway, startIndex string, endIndex string, address string) ([]*PermittedNetworkInfo, error) {
	log.Printf("Get PermittedNetworks: \n")
	methodName := "GetPermittedNetworksByIndex"

	if ccname := os.Getenv("CHAINCODE_NAME"); ccname != "" {
		chaincodeName = ccname
	}

	if cname := os.Getenv("CHANNEL_NAME"); cname != "" {
		channelName = cname
	}
	network := gateway.GetNetwork(channelName)
	contract := network.GetContract(chaincodeName)

	queryResponse, err := contract.EvaluateTransaction(methodName, startIndex, endIndex, address)
	if err != nil {
		log.Printf("failed to submit transaction: %s", err)
		return nil, err
	}
	var results []*PermittedNetworkInfo
	err = json.Unmarshal(queryResponse, &results)
	if err != nil {
		log.Printf("failed to submit transaction: %s", err)
		return nil, err
	}

	log.Printf("Result:%s\n", string(queryResponse))
	return results, nil
}
