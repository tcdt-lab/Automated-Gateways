package interfaces

import (
	dataType "relay/src/data_types"
)

type IPermittedMethodMediator interface {
	AddPermittedMethod(permittedNetworkId string, permittedNetworkName string, chaincode string, channel string, inputArgs string, outputArgs string) (*dataType.PermittedMethodInfo, error)
	PermittedMethodExists(id string) (bool, error)
	GetPermittedMethod(permittedMethodId string) (*dataType.PermittedMethodInfo, error)
	GetPermittedMethodsByIndexAndAddress(permittedNetworkId string, startIndex string, endIndex string) ([]*dataType.PermittedMethodInfo, error)
	GetPermittedMethodsByNetworkId(permittedNetworkId string) ([]*dataType.PermittedMethodInfo, error)
	RemovePermittedMethod(id string) error
	UpdatePermittedMethod(id string, name string, chaincode string, channel string, inputArgs string, outputArgs string) error
	InvokePermittedMethod(name string, chaincode string, channel string, inputArgs []string) (*string, error)
}
