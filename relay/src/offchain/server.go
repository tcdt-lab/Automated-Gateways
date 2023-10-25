package offchain

import (
	"context"
	"encoding/json"
	pb "relay/scripts"
	"relay/src/mediator"
)

type IOPserver struct {
	pb.UnimplementedIOPServer
}

func (s *IOPserver) GetPermittedNetworkInfo(address *pb.PermittedNetworkAddress, stream pb.IOP_GetPermittedNetworkInfoServer) error {
	var iopMediator mediator.IopMediator
	res, err := iopMediator.ReturnPermittedNetworkInfo(address.Address, mediator.HYPERLEDGER_FABRIC_NETWROK_TYPE)
	if err != nil {
		return err
	}
	for _, network := range res {
		permittedNetworkInfo := &pb.PermittedNetworkInfo{
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

func (s *IOPserver) InvokePermittedMethod(ctx context.Context, info *pb.MethodInfo) (*pb.MethodResponse, error) {
	var iopMediator mediator.IopMediator
	var response *pb.MethodResponse
	response = &pb.MethodResponse{}
	inputArgs, errStr := convertStringArrayToArrayOfStrings(info.MethodInput)
	if errStr != nil {
		response.Response = "error"
		response.Error = errStr.Error()
		return response, errStr
	}
	res, err := iopMediator.InvokePermittedMethod(info.Name, info.ChaincodeName, info.ChannelName, inputArgs, mediator.HYPERLEDGER_FABRIC_NETWROK_TYPE)
	if err != nil {
		response.Response = "error"
		response.Error = err.Error()
		return response, err
	}
	response.Response = *res
	response.Error = ""
	return response, nil
}

func (s *IOPserver) GetPermittedMethodsList(networkId *pb.PermittedNetworkId, stream pb.IOP_GetPermittedMethodsListServer) error {
	var iopMediator mediator.IopMediator
	res, err := iopMediator.ReturnPermittedMethodList(networkId.NetworkId, mediator.HYPERLEDGER_FABRIC_NETWROK_TYPE)
	if err != nil {
		return err
	}

	for _, method := range res {
		permittedMethodInfo := &pb.MethodInfo{
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
