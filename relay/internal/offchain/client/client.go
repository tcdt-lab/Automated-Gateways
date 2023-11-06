package client

import (
	"context"
	"crypto/tls"
	"crypto/x509"
	"flag"
	"github.com/tcdt-lab/Automated-Gateways/relay/data_types"
	"github.com/tcdt-lab/Automated-Gateways/relay/internal/scripts"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials"
	"io"
	"log"
	"os"
	"time"
)

var (
	caFile             = flag.String("ca_file", "", "The file containing the CA root cert file")
	serverAddr         = flag.String("addr", "192.168.100.125:50051", "The server address in the format of host:port")
	serverHostOverride = flag.String("server_host_override", "x.test.example.com", "The server name used to verify the hostname returned by the TLS handshake")
)

func openConnection() (*grpc.ClientConn, error) {

	flag.Parse()
	caCert, err := os.ReadFile("../relay/certs/ca-cert.pem")
	if err != nil {
		log.Fatal(caCert)
	}

	certPool := x509.NewCertPool()
	if ok := certPool.AppendCertsFromPEM(caCert); !ok {
		log.Fatal(err)
	}

	clientCert, err := tls.LoadX509KeyPair("../relay/certs/client-cert.pem", "../relay/certs/client-key.pem")
	if err != nil {
		log.Fatal(err)
	}

	config := &tls.Config{
		Certificates: []tls.Certificate{clientCert},
		RootCAs:      certPool,
	}

	tlsCredential := credentials.NewTLS(config)

	var opts []grpc.DialOption
	opts = append(opts, grpc.WithTransportCredentials(tlsCredential))
	return grpc.Dial(*serverAddr, opts...)
}

func closeConnection(conn *grpc.ClientConn) error {
	return conn.Close()
}

func InvokeAccessibleMethod(accessibleMethodName string, chaincodeName string, channelName,
	accessibleMethodInput string, accessibleMethodOutput string) (string, string) {
	conn, err := openConnection()
	if err != nil {
		log.Fatalf("openConnection failed: %v", err)
	}
	log.Println("invokeMethod method is invoked")
	defer closeConnection(conn)
	client := scripts.NewIOPClient(conn)
	methodInfo := scripts.MethodInfo{
		Name:          accessibleMethodName,
		ChaincodeName: chaincodeName,
		ChannelName:   channelName,
		MethodInput:   accessibleMethodInput,
		MethodOutput:  accessibleMethodOutput,
	}
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()
	var res, errInvoke = client.InvokePermittedMethod(ctx, &methodInfo)
	if errInvoke != nil {
		log.Fatalf("invokeMethod failed: %v", err)
	}

	log.Println(res)
	return res.Response, res.Error
}

func GetAccessibleMethodsList(accessibleNetworkId string) ([]*data_types.MethodInfo, error) {
	conn, err := openConnection()
	if err != nil {
		log.Fatalf("openConnection failed: %v", err)
		return nil, err
	}
	log.Println("getMethodsList method is invoked")
	defer closeConnection(conn)
	client := scripts.NewIOPClient(conn)

	networkId := scripts.PermittedNetworkId{
		NetworkId: accessibleNetworkId,
	}
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()
	var network, errGetPermittedList = client.GetPermittedMethodsList(ctx, &networkId)
	if errGetPermittedList != nil {
		log.Printf("getMethodsList failed: %v", errGetPermittedList)
		return nil, errGetPermittedList
	}
	var methodInfos []*data_types.MethodInfo

	for {
		accessibleMethod, err := network.Recv()
		if err == io.EOF {
			break
		}
		if err != nil {
			log.Printf("client.ListNetworks failed: %v", err)
			return nil, err
		}

		methodInfos = append(methodInfos, &data_types.MethodInfo{
			Name:       accessibleMethod.Name,
			Chaincode:  accessibleMethod.ChaincodeName,
			Channel:    accessibleMethod.ChannelName,
			InputArgs:  accessibleMethod.MethodInput,
			OutputArgs: accessibleMethod.MethodOutput,
		})
	}
	println("client Res")
	println(methodInfos[0].Name)
	return methodInfos, nil
}

func GetNetworkInformation(selfAddress string) ([]*data_types.AccessibleNetworkInfo, error) {
	conn, err := openConnection()
	if err != nil {
		log.Fatalf("openConnection failed: %v", err)
		return nil, err
	}
	log.Println("GetNetworkInformation method is invoked")
	defer closeConnection(conn)
	client := scripts.NewIOPClient(conn)

	log.Println("getInfo method is invoked")
	loginInfo := scripts.PermittedNetworkAddress{

		Address: selfAddress,
	}
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()
	var network, errInovke = client.GetPermittedNetworkInfo(ctx, &loginInfo)
	if errInovke != nil {
		log.Fatalf("getInfo failed: %v", errInovke)
		return nil, errInovke
	}
	var netInfos []*data_types.AccessibleNetworkInfo
	for {
		permittedNets, err := network.Recv()
		if err == io.EOF {
			break
		}
		if err != nil {
			log.Fatalf("client.ListNetworks failed: %v", err)
		}

		netInfos = append(netInfos, &data_types.AccessibleNetworkInfo{
			NetworkName:         permittedNets.NetworkName,
			IP:                  permittedNets.Network_IP,
			ADDRESS:             permittedNets.NetworkAddress,
			AccessibleNetworkId: permittedNets.NetworkId,
			CompanyName:         permittedNets.CompanyName,
		})

	}
	println("client Res3")
	log.Println(netInfos[0].NetworkName, netInfos[0].AccessibleNetworkId)
	return netInfos, nil
}
