package interfaces

import (
	dataType "github.com/tcdt-lab/Automated-Gateways/relay/data_types"
)

type IPermittedMethodMediator interface {
	AddPermittedMethod(permittedNetworkId string, permittedNetworkName string, chaincode string, channel string, inputArgs string, outputArgs string) (*dataType.MethodInfo, error)
	PermittedMethodExists(id string) (bool, error)
	GetPermittedMethod(permittedMethodId string) (*dataType.MethodInfo, error)
	GetPermittedMethodsByIndexAndAddress(permittedNetworkId string, startIndex string, endIndex string) ([]*dataType.MethodInfo, error)
	GetPermittedMethodsByNetworkId(permittedNetworkId string) ([]*dataType.MethodInfo, error)
	RemovePermittedMethod(id string) error
	UpdatePermittedMethod(id string, name string, chaincode string, channel string, inputArgs string, outputArgs string) error
	InvokePermittedMethod(name string, chaincode string, channel string, inputArgs []string) (*string, error)
}
