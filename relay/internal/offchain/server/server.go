package server

import (
	"context"
	tlsInterface "crypto/tls"
	"crypto/x509"
	"encoding/json"
	"flag"
	"fmt"
	grpc_auth "github.com/grpc-ecosystem/go-grpc-middleware/auth"
	auth "github.com/tcdt-lab/Automated-Gateways/relay/internal/offchain/authentication"
	"github.com/tcdt-lab/Automated-Gateways/relay/internal/scripts"
	mediator2 "github.com/tcdt-lab/Automated-Gateways/relay/mediator"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials"
	"log"
	"net"
	"os"
)

type IOPserver struct {
	scripts.UnimplementedIOPServer
}

var (
	certFile   = flag.String("cert_file", "", "The TLS cert file")
	keyFile    = flag.String("key_file", "", "The TLS key file")
	jsonDBFile = flag.String("json_db_file", "", "A json file containing a list of features")
	port       = flag.Int("port", 50051, "The server port")
)

func (s *IOPserver) GetPermittedNetworkInfo(address *scripts.PermittedNetworkAddress, stream scripts.IOP_GetPermittedNetworkInfoServer) error {
	var iopMediator mediator2.IopMediator
	res, err := iopMediator.ReturnPermittedNetworkInfo(address.Address, mediator2.HYPERLEDGER_FABRIC_NETWROK_TYPE)
	if err != nil {
		return err
	}
	for _, network := range res {
		permittedNetworkInfo := &scripts.PermittedNetworkInfo{
			NetworkId:      network.PermittedNetworkId,
			NetworkName:    network.NetworkName,
			Network_IP:     network.IP,
			NetworkAddress: network.ADDRESS,
			CompanyName:    network.CompanyName,
		}
		stream.Send(permittedNetworkInfo)
	}
	return nil
}

func (s *IOPserver) InvokePermittedMethod(ctx context.Context, info *scripts.MethodInfo) (*scripts.MethodResponse, error) {
	var iopMediator mediator2.IopMediator
	var response *scripts.MethodResponse
	response = &scripts.MethodResponse{}
	inputArgs, errStr := convertStringArrayToArrayOfStrings(info.MethodInput)
	if errStr != nil {
		response.Response = "error"
		response.Error = errStr.Error()
		return response, errStr
	}
	res, err := iopMediator.InvokePermittedMethod(info.Name, info.ChaincodeName, info.ChannelName, inputArgs, mediator2.HYPERLEDGER_FABRIC_NETWROK_TYPE)
	if err != nil {
		response.Response = "error"
		response.Error = err.Error()
		return response, err
	}
	response.Response = *res
	response.Error = ""
	return response, nil
}

func (s *IOPserver) GetPermittedMethodsList(networkId *scripts.PermittedNetworkId, stream scripts.IOP_GetPermittedMethodsListServer) error {
	var iopMediator mediator2.IopMediator
	res, err := iopMediator.ReturnPermittedMethodList(networkId.NetworkId, mediator2.HYPERLEDGER_FABRIC_NETWROK_TYPE)
	if err != nil {
		return err
	}

	for _, method := range res {
		permittedMethodInfo := &scripts.MethodInfo{
			Name:          method.Name,
			ChaincodeName: method.Chaincode,
			ChannelName:   method.Channel,
			MethodInput:   method.InputArgs,
			MethodOutput:  method.OutputArgs,
		}
		stream.Send(permittedMethodInfo)
	}
	return nil
}
func convertStringArrayToArrayOfStrings(arrayStr string) ([]string, error) {

	var res []string
	err := json.Unmarshal([]byte(arrayStr), &res)
	if err != nil {
		return nil, err
	}

	return res, nil
}

func StartServer() {

	log.Println("gRPC server is starting...")
	flag.Parse()
	lis, err := net.Listen("tcp", fmt.Sprintf("0.0.0.0:%d", *port))
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}

	caPem, err := os.ReadFile("../relay/certs/ca-cert.pem")
	if err != nil {
		log.Fatal(err)
	}

	// create cert pool and append ca's cert
	certPool := x509.NewCertPool()
	if !certPool.AppendCertsFromPEM(caPem) {
		log.Fatal(err)
	}

	// read server cert & key
	serverCert, err := tlsInterface.LoadX509KeyPair("../relay/certs/server-cert.pem", "../relay/certs/server-key.pem")
	if err != nil {
		log.Fatal(err)
	}

	// configuration of the certificate what we want to
	conf := &tlsInterface.Config{
		Certificates: []tlsInterface.Certificate{serverCert},
		ClientAuth:   tlsInterface.RequireAndVerifyClientCert,
		ClientCAs:    certPool,
	}

	tlsCredentials := credentials.NewTLS(conf)

	var opts []grpc.ServerOption

	opts = []grpc.ServerOption{grpc.Creds(tlsCredentials),
		grpc.StreamInterceptor(grpc_auth.StreamServerInterceptor(auth.TlsAuth)),
		grpc.UnaryInterceptor(grpc_auth.UnaryServerInterceptor(auth.TlsAuth))}

	grpcServer := grpc.NewServer(opts...)
	s := IOPserver{}
	scripts.RegisterIOPServer(grpcServer, &s)
	grpcServer.Serve(lis)
}
