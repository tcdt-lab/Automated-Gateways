package offchain

import (
	"context"
	"flag"
	"fmt"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
	"io"
	"log"
	pb "relay/scripts"
	"time"
)

var (
	tls                = flag.Bool("tls", false, "Connection uses TLS if true, else plain TCP")
	caFile             = flag.String("ca_file", "", "The file containing the CA root cert file")
	serverAddr         = flag.String("addr", "localhost:50051", "The server address in the format of host:port")
	serverHostOverride = flag.String("server_host_override", "x.test.example.com", "The server name used to verify the hostname returned by the TLS handshake")
)

func main() {
	flag.Parse()
	var opts []grpc.DialOption
	opts = append(opts, grpc.WithTransportCredentials(insecure.NewCredentials()))
	conn, err := grpc.Dial(*serverAddr, opts...)
	if err != nil {
		log.Fatalf("fail to dial: %v", err)
	}
	defer conn.Close()

	client := pb.NewIOPClient(conn)
	doLogin(client)
	queryMethods(client)
	doInvoke(client)
}

func doInvoke(client pb.IOPClient) {
	log.Println("client side doInvoke method is invoked")
	var methodInfo = &pb.MethodInfo{
		MethodName:    "mockMethodccccccccccc",
		ChaincodeName: "mockChaincode",
		ChannelName:   "mockChannel",
		Method_Id:     "mockID",
		MethodInput:   "mockInput",
		MethodOutput:  "mockOutput",
		MethodType:    "mockType",
		Network: &pb.OutsiderNetwork{
			NetworkId:    "mockID",
			NetworkToken: "mockToken",
		},
	}
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()
	var response, err = client.Invoke(ctx, methodInfo)
	if err != nil {
		log.Fatalf("Invoke failed: %v", err)
	}
	fmt.Print(response)
}

func doLogin(client pb.IOPClient) {
	log.Println("Login method is invoked")
	var loginInfo = &pb.OutsiderNetworkLoginInfo{
		NetworkUsername: "mockUsername",
		NetworkId:       "mockID",
		NetworkPassword: "mockPassword",
		Network_IP:      "mockIP",
	}
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()
	var network, err = client.Login(ctx, loginInfo)
	if err != nil {
		log.Fatalf("Login failed: %v", err)
	}
	log.Printf("Login succeeded: %v", network)
}

func queryMethods(client pb.IOPClient) {
	log.Println("queryMethods method is invoked")
	var outsiderNetwork = &pb.OutsiderNetwork{
		NetworkId:    "mockID",
		NetworkToken: "mockToken",
	}
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()
	var stream, err = client.QueryMethods(ctx, outsiderNetwork)
	if err != nil {
		log.Fatalf("QueryMethods failed: %v", err)
	}
	for {
		method, err := stream.Recv()
		if err == io.EOF {
			break
		}
		if err != nil {
			log.Fatalf("client.QueryMethods failed: %v", err)
		}
		log.Printf("Method: name: %q", method.MethodName)
	}
	log.Printf("QueryMethod succeeded: %v", stream)
}
