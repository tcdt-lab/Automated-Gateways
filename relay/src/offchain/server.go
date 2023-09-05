package offchain

import (
	"context"
	"log"
	pb "relay/scripts"
	onchain "relay/src/onchain"
)

type IOPserver struct {
	pb.UnimplementedIOPServer
	methodInfoList []*pb.MethodInfo
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

func (s *IOPserver) Login(ctx context.Context, info *pb.OutsiderNetworkLoginInfo) (*pb.OutsiderNetwork, error) {
	log.Println(" Login METHOD START")
	var invokeRes, error = onchain.CheckNetworkInfo(info)
	if error != nil {
		log.Println(" Login METHOD ENDED. Error : %s", error)
		return nil, error
	}
	log.Println(" Login METHOD ENDED.")
	return invokeRes, nil
}

func (s *IOPserver) QueryMethods(outsiderNetwork *pb.OutsiderNetwork, stream pb.IOP_QueryMethodsServer) error {
	log.Println("QueryMethods IS INVOKED")
	var methodsList, error = onchain.QueryMethods(outsiderNetwork)
	if error != nil {
		return errorcd
	}
	for _, method := range methodsList {

		if err := stream.Send(method); err != nil {
			return err
		}

	}
	return nil
}
