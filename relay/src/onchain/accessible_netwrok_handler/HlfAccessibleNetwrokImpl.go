package accessible_netwrok_handler

import (
	"bytes"
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

type AccessibleNetworkInfo struct {
	NetworkName         string `json:"NetworkName"`
	IP                  string `json:"IP"`
	ADDRESS             string `json:"ADDRESS"`
	CompanyName         string `json:"CompanyName"`
	AccessibleNetworkId string `json:"AccessibleNetworkId"`
}

var channelName = "mychannel"
var chaincodeName = "accessible_net"

type HlfAccessibleNetwork struct {
}

func OpenConnection() (*grpc.ClientConn, *client.Gateway, error) {
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

func CloseConnection(clientConnection *grpc.ClientConn, gw *client.Gateway) error {
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

func (hlfAccessibleNetwork *HlfAccessibleNetwork) CreateAccessibleNetwork(gw *client.Gateway, networkName string, ip string, address string, companyName string) (*AccessibleNetworkInfo, error) {
	log.Printf("Creating Accessible Network: %s\n", networkName)
	methodName := "CreateAccessibleNetwork"

	if ccname := os.Getenv("CHAINCODE_NAME"); ccname != "" {
		chaincodeName = ccname
	}

	if cname := os.Getenv("CHANNEL_NAME"); cname != "" {
		channelName = cname
	}
	network := gw.GetNetwork(channelName)
	contract := network.GetContract(chaincodeName)

	submitRes, err := contract.SubmitTransaction(methodName, networkName, ip, address, companyName)
	var accessibleNetworkInfo AccessibleNetworkInfo
	err = json.Unmarshal(submitRes, &accessibleNetworkInfo)
	if err != nil {
		log.Printf("failed to submit transaction: %s", err)
		return nil, err
	}

	log.Printf("Result:%s\n", string(submitRes))
	return &accessibleNetworkInfo, nil
}

func (hlfAccessibleNetwork *HlfAccessibleNetwork) AccessibleNetworkExists(gw *client.Gateway, id string) (bool, error) {
	log.Printf("Check If AccessibleNetwork Exists: %s\n", id)
	methodName := "AccessibleNetworkExists"

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

	log.Printf("Result:%s\n", string(evalRes))
	return resBoolean, nil
}

func (hlfAccessibleNetwork *HlfAccessibleNetwork) RemoveAccessibleNetwork(gw *client.Gateway, id string) error {
	log.Printf("Removing AccessibleNetwork: %s\n", id)
	methodName := "RemoveAccessibleNetwork"

	if ccname := os.Getenv("CHAINCODE_NAME"); ccname != "" {
		chaincodeName = ccname
	}

	if cname := os.Getenv("CHANNEL_NAME"); cname != "" {
		channelName = cname
	}
	network := gw.GetNetwork(channelName)
	contract := network.GetContract(chaincodeName)

	submitRes, err := contract.SubmitTransaction(methodName, id)
	if err != nil {
		log.Printf("failed to submit remove transaction: %s", err)
		return err
	}

	log.Printf("Result:%s\n", string(submitRes))
	return nil

}

func (hlfAccessibleNetwork *HlfAccessibleNetwork) UpdateAccessibleNetwork(gw *client.Gateway, id string, networkName string, ip string, address string, companyName string) error {
	panic("implement me")
}

func (hlfAccessibleNetwork *HlfAccessibleNetwork) GetAccessibleNetwork(gw *client.Gateway, id string) (*AccessibleNetworkInfo, error) {
	panic("implement me")
}

func (hlfAccessibleNetwork *HlfAccessibleNetwork) QueryAllAccessibleNetworks(gw *client.Gateway) ([]*AccessibleNetworkInfo, error) {
	panic("implement me")
}

func formatJSON(data []byte) string {
	var prettyJSON bytes.Buffer
	if err := json.Indent(&prettyJSON, data, "", "  "); err != nil {
		fmt.Errorf("failed to parse JSON: %w", err)
	}
	return prettyJSON.String()
}
