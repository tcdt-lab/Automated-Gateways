package permitted_method_handler

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

func convertContractPermittedMethodDataType(input ethereum_handler.PermittedMethodManagementContractPermittedMethod) (data *data_types.MethodInfo) {
	data = &data_types.MethodInfo{
		PermittedMethodId: input.PermittedMethodId.String(),
		Name:              input.Name,
		Chaincode:         input.Chaincode,
		InputArgs:         input.InputArgs,
		OutputArgs:        input.OutputArgs,
	}
	return data
}
