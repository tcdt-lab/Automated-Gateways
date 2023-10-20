package offchain

import (
	pb "relay/scripts"
	"relay/src/mediator"
)

type IOPserver struct {
	pb.UnimplementedIOPServer
}

func (s *IOPserver) GetPermittedNetworkInfo(address *pb.PermittedNetworkAddress, stream pb.IOP_GetPermittedNetworkInfoServer) error {
	var iopMediator mediator.IopMediator
	res, err := iopMediator.GetPermittedNetwork(address.Address, mediator.HYPERLEDGER_FABRIC_NETWROK_TYPE)
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
