package onchain

import (
	"bytes"
	"crypto/x509"
	"encoding/json"
	"fmt"
	"github.com/hyperledger/fabric-gateway/pkg/client"
	"github.com/hyperledger/fabric-gateway/pkg/identity"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials"
	"os"
	"relay/scripts"
	"time"
)

const (
	mspID        = "Org1MSP"
	certPath     = "../relay/src/onchain/chain_certs/user_cert.pem"
	keyPath      = "../relay/src/onchain/chain_certs/user_key.pem"
	tlsCertPath  = "../relay/src/onchain/chain_certs/ca.crt"
	peerEndpoint = "localhost:7051"
	gatewayPeer  = "peer0.org1.example.com"
)

var now = time.Now()
var assetId = fmt.Sprintf("asset%d", now.Unix()*1e3+int64(now.Nanosecond())/1e6)

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
	certificate, err := loadCertificate(tlsCertPath)
	if err != nil {
		panic(err)
	}

	certPool := x509.NewCertPool()
	certPool.AddCert(certificate)
	transportCredentials := credentials.NewClientTLSFromCert(certPool, gatewayPeer)

	connection, err := grpc.Dial(peerEndpoint, grpc.WithTransportCredentials(transportCredentials))
	if err != nil {
		panic(fmt.Errorf("failed to create gRPC connection: %w", err))
	}

	return connection
}
func newSign() identity.Sign {

	privateKeyPEM, err := os.ReadFile(keyPath)

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
	certificate, err := loadCertificate(certPath)
	if err != nil {
		panic(err)
	}

	id, err := identity.NewX509Identity(mspID, certificate)
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
	gw.Close()
	clientConnection.Close()
	return nil
}

func GetPermittedNetworksInfo(address string, gw *client.Gateway) (string, error) {
	chaincodeName := "mng"
	if ccname := os.Getenv("CHAINCODE_NAME"); ccname != "" {
		chaincodeName = ccname
	}

	channelName := "mychannel"
	if cname := os.Getenv("CHANNEL_NAME"); cname != "" {
		channelName = cname
	}
	network := gw.GetNetwork(channelName)
	contract := network.GetContract(chaincodeName)
	methodName := "GetPermittedNetworkByAddress"
	fmt.Println("\n--> Evaluate Transaction: GetAllAssets, function returns all the current assets on the ledger")

	evaluateResult, err := contract.EvaluateTransaction(methodName, address)
	if err != nil {
		panic(fmt.Errorf("failed to evaluate transaction: %w", err))
	}
	result := formatJSON(evaluateResult)

	fmt.Printf("*** Result:%s\n", result)
	return result, nil
}
func InvokeMethod(methodInfo *scripts.MethodInfo, gw *client.Gateway) (*scripts.Response, error) {
	chaincodeName := methodInfo.GetChaincodeName()
	if ccname := os.Getenv("CHAINCODE_NAME"); ccname != "" {
		chaincodeName = ccname
	}

	channelName := methodInfo.GetChannelName()
	if cname := os.Getenv("CHANNEL_NAME"); cname != "" {
		channelName = cname
	}
	network := gw.GetNetwork(channelName)
	contract := network.GetContract(chaincodeName)
	methodName := methodInfo.GetMethodName()
	fmt.Println("\n--> Evaluate Transaction: GetAllAssets, function returns all the current assets on the ledger")

	evaluateResult, err := contract.EvaluateTransaction(methodName, address)
	if err != nil {
		panic(fmt.Errorf("failed to evaluate transaction: %w", err))
	}
	result := formatJSON(evaluateResult)

	fmt.Printf("*** Result:%s\n", result)
	return result, nil
}

func QueryMethods(network *scripts.PermittedNetworkId) ([]*scripts.MethodInfo, error) {
	fmt.Println("The methods from network %s has been queried", network.GetNetworkId())

	return nil, nil
}
func formatJSON(data []byte) string {
	var prettyJSON bytes.Buffer
	if err := json.Indent(&prettyJSON, data, "", "  "); err != nil {
		panic(fmt.Errorf("failed to parse JSON: %w", err))
	}
	return prettyJSON.String()
}

func formatMethodArgs(jsonArgStr string) []string {
	var jsonArgs map[string]interface{}
	json.Unmarshal([]byte(jsonArgStr), &jsonArgs)
	for _, v := range jsonArgs {
		fmt.Println(v)
	}
}
