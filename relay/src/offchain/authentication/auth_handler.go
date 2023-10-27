package authentication

import (
	"context"
	"log"

	permittedHAndler "github.com/tcdt-lab/Automated-Gateways/relay/src/onchain/permitted_network_handler"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/credentials"
	"google.golang.org/grpc/peer"
	"google.golang.org/grpc/status"
)

func TlsAuth(ctx context.Context) (newCtx context.Context, err error) {
	p, ok := peer.FromContext(ctx)
	if !ok {
		return ctx, status.Error(codes.Unauthenticated, "no peer found")
	}

	tlsAuth, ok := p.AuthInfo.(credentials.TLSInfo)
	if !ok {
		return ctx, status.Error(codes.Unauthenticated, "unexpected peer transport credentials")
	}

	if len(tlsAuth.State.VerifiedChains) == 0 || len(tlsAuth.State.VerifiedChains[0]) == 0 {
		return ctx, status.Error(codes.Unauthenticated, "could not verify peer certificate")
	}
	commonName := tlsAuth.State.VerifiedChains[0][0].Subject.CommonName
	isPermitted, err := isHostPermitted(commonName)
	if err != nil {
		return ctx, status.Error(codes.Unauthenticated, "could not verify peer certificate")
	}
	if isPermitted != nil && *isPermitted == false {
		return ctx, status.Error(codes.Unauthenticated, "invalid subject common name")
	}

	return ctx, nil
}

func isHostPermitted(hostName string) (*bool, error) {
	var hlfPermittedMethod permittedHAndler.HlfPermittedNetwork
	var boolValue bool
	clientConn, gateway, err := hlfPermittedMethod.OpenConnection()
	if err != nil {
		log.Println("Error opening connection: %v", err)
	}
	defer hlfPermittedMethod.CloseConnection(clientConn, gateway)

	res, err := hlfPermittedMethod.GetPermittedNetworksByAddress(gateway, hostName)
	if err != nil {
		log.Println(" Error in Getting All permitted Netwroks : %v", err)
		return nil, err
	}
	for _, permittedNetwork := range res {
		if permittedNetwork.ADDRESS == hostName {
			boolValue = true
			return &boolValue, nil
		}
	}
	boolValue = false
	return &boolValue, nil

}
