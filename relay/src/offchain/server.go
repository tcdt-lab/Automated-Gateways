package offchain

import (
	"context"
	"log"
	pb "relay/scripts"
	onchain "relay/src/onchain"
)

type IOPserver struct {
	pb.UnimplementedIOPServer
}

func (s *IOPserver) Invoke(ctx context.Context, info *pb.MethodInfo) (*pb.Response, error) {
	log.Println(" INVOKE METHOD START")
	var invokeRes, error = onchain.InvokeMethod(info)
	if error != nil {
		log.Println(" INVOKE METHOD ENDED. Error : %s", error)
		return nil, error
	}
	log.Println(" INVOKE METHOD ENDED.")
	return invokeRes, nil
}

func (s *IOPserver) GetNetworkInfo(ctx context.Context, info *pb.OutsiderNetworkId) (*pb.OutsiderNetworkInfo, error) {
	log.Println(" GetNetworkInfo METHOD START")
	clientCon, gw, err := onchain.OpenConnection()
	if err != nil {
		log.Println(" GetNetworkInfo METHOD ENDED.Error in Chain connection Error : %s", err)
	}
	defer onchain.CloseConnection(clientCon, gw)
	var invokeRes, error = onchain.GetOutsiderNetworksInfo(info.GetNetworkId(), gw)
	if error != nil {
		log.Println(" GetNetworkInfo METHOD ENDED. Error : %s", error)
		return nil, error
	}
	log.Println(" Login METHOD ENDED.")
	netId := pb.OutsiderNetworkInfo{
		NetworkId: invokeRes,
	}
	return &netId, nil
}

func (s *IOPserver) QueryMethods(outsiderNetwork *pb.OutsiderNetworkId, stream pb.IOP_QueryMethodsServer) error {
	log.Println("QueryMethods IS INVOKED")
	var methodsList, error = onchain.QueryMethods(outsiderNetwork)
	if error != nil {
		return error
	}
	for _, method := range methodsList {

		if err := stream.Send(method); err != nil {
			return err
		}

	}
	return nil
}
