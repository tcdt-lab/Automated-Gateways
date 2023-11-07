package interfaces

import (
	dataType "github.com/tcdt-lab/Automated-Gateways/relay/data_types"
)

type IPermittedMethodMediator interface {
	AddPermittedMethod(permittedNetworkId string, permittedMethodName string, chaincode string, channel string, inputArgs string, outputArgs string) (*dataType.MethodInfo, error)
	PermittedMethodExists(id string) (bool, error)
	GetPermittedMethod(permittedMethodId string) (*dataType.MethodInfo, error)
	GetPermittedMethodsByIndexAndNetworkId(startIndex string, endIndex string, permittedNetworkId string) ([]*dataType.MethodInfo, error)
	GetPermittedMethodsByNetworkId(permittedNetworkId string) ([]*dataType.MethodInfo, error)
	RemovePermittedMethod(id string) error
	UpdatePermittedMethod(id string, name string, chaincode string, channel string, inputArgs string, outputArgs string) error
	InvokePermittedMethod(name string, chaincode string, channel string, inputArgs []string) (*string, error)
}
