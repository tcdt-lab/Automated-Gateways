package permitted_network_handler

import (
	"errors"
	"github.com/tcdt-lab/Automated-Gateways/relay/data_types"
	"github.com/tcdt-lab/Automated-Gateways/relay/internal/onchain/ethereum_handler"
	"log"
	"math/big"
)

func stringToBigIntConvertor(number string) (*big.Int, error) {
	n := new(big.Int)
	numberInt, isSuccess := n.SetString(number, 10)
	if !isSuccess {
		log.Fatalf("Failed to convert accessibleNetworkId to big.Int")
		return nil, errors.New("Failed to convert accessibleNetworkId to big.Int")
	}
	return numberInt, nil
}

func convertContractPermittedNetworkDataType(input ethereum_handler.PermittedNetworkManagementContractPermittedNetwork) (data *data_types.PermittedNetworkInfo) {
	data = &data_types.PermittedNetworkInfo{
		PermittedNetworkId: input.PermittedNetworkId.String(),
		CompanyName:        input.CompanyName,
		IP:                 input.NetworkIP,
		ADDRESS:            input.NetworkAddress,
		NetworkName:        input.NetworkName,
	}
	return data
}
