package main

import (
	tlsInterface "crypto/tls"
	"crypto/x509"
	"flag"
	"fmt"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials"
	"log"
	"net"
	"os"
	"relay/scripts"
	"relay/src/offchain"
)

var (
	tls        = flag.Bool("tls", false, "Connection uses TLS if true, else plain TCP")
	certFile   = flag.String("cert_file", "", "The TLS cert file")
	keyFile    = flag.String("key_file", "", "The TLS key file")
	jsonDBFile = flag.String("json_db_file", "", "A json file containing a list of features")
	port       = flag.Int("port", 50051, "The server port")
)

func main() {

	startServer()
	//offchain.RunClient()
}

func startServer() {
	log.Println("gRPC server is starting...")
	flag.Parse()
	lis, err := net.Listen("tcp", fmt.Sprintf("localhost:%d", *port))
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

	//create tls certificate
	tlsCredentials := credentials.NewTLS(conf)
	//
	var opts []grpc.ServerOption
	//if *tls {
	//	if *certFile == "" {
	//		//*certFile = data.Path("x509/server_cert.pem")
	//	}
	//	if *keyFile == "" {
	//		//*keyFile = data.Path("x509/server_key.pem")
	//	}
	//	creds, err := credentials.NewServerTLSFromFile(*certFile, *keyFile)
	//	if err != nil {
	//		log.Fatalf("Failed to generate credentials: %v", err)
	//	}
	opts = []grpc.ServerOption{grpc.Creds(tlsCredentials)}

	grpcServer := grpc.NewServer(opts...)
	s := offchain.IOPserver{}
	scripts.RegisterIOPServer(grpcServer, &s)
	grpcServer.Serve(lis)
}