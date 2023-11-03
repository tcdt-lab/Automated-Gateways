package main

import (
	server "github.com/tcdt-lab/Automated-Gateways/relay/internal/offchain/server"
	cli "github.com/tcdt-lab/Automated-Gateways/relay/iop_cli"
)

//var (
//	tls        = flag.Bool("tls", false, "Connection uses TLS if true, else plain TCP")
//	certFile   = flag.String("cert_file", "", "The TLS cert file")
//	keyFile    = flag.String("key_file", "", "The TLS key file")
//	jsonDBFile = flag.String("json_db_file", "", "A json file containing a list of features")
//	port       = flag.Int("port", 50051, "The server port")
//)

func main() {

	//onchain.CloseConnection()
	go server.StartServer()
	cli.StartCli()
	//offchain.GetAccessibleMethodsList("PermittedNetwork_localhost_7")

	//offchain.InvokeAccessibleMethod("AddTwoNumbers", "addition", "mychannel", `["9","5"]`, "{'string'}")
	//offchain.GetNetworkInformation("localhost")
}
func startOnchainProcess() {

}

//func startServer() {
//	log.Println("gRPC server is starting...")
//	flag.Parse()
//	lis, err := net.Listen("tcp", fmt.Sprintf("localhost:%d", *port))
//	if err != nil {
//		log.Fatalf("failed to listen: %v", err)
//	}
//
//	caPem, err := os.ReadFile("../relay/certs/ca-cert.pem")
//	if err != nil {
//		log.Fatal(err)
//	}
//
//	// create cert pool and append ca's cert
//	certPool := x509.NewCertPool()
//	if !certPool.AppendCertsFromPEM(caPem) {
//		log.Fatal(err)
//	}
//
//	// read server cert & key
//	serverCert, err := tlsInterface.LoadX509KeyPair("../relay/certs/server-cert.pem", "../relay/certs/server-key.pem")
//	if err != nil {
//		log.Fatal(err)
//	}
//
//	// configuration of the certificate what we want to
//	conf := &tlsInterface.Config{
//		Certificates: []tlsInterface.Certificate{serverCert},
//		ClientAuth:   tlsInterface.RequireAndVerifyClientCert,
//		ClientCAs:    certPool,
//	}
//
//	tlsCredentials := credentials.NewTLS(conf)
//
//	var opts []grpc.ServerOption
//
//	opts = []grpc.ServerOption{grpc.Creds(tlsCredentials),
//		grpc.StreamInterceptor(grpc_auth.StreamServerInterceptor(auth.TlsAuth)),
//		grpc.UnaryInterceptor(grpc_auth.UnaryServerInterceptor(auth.TlsAuth))}
//
//	grpcServer := grpc.NewServer(opts...)
//	s := offchain.IOPserver{}
//	scripts.RegisterIOPServer(grpcServer, &s)
//	grpcServer.Serve(lis)
//}
