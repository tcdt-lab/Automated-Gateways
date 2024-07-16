package accessible_network_handler

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

func convertContractAccessibleDataType(input ethereum_handler.AccessibleNetworkManagementContractAccessibleNetwork) (data *data_types.AccessibleNetworkInfo) {
	data = &data_types.AccessibleNetworkInfo{
		AccessibleNetworkId: input.AccessibleNetworkId.String(),
		CompanyName:         input.CompanyName,
		IP:                  input.NetworkIP,
		ADDRESS:             input.NetworkAddress,
		NetworkName:         input.NetworkName,
	}
	return data
}
