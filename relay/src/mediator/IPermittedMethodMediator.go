package mediator

import (
	permittedHandler "relay/src/onchain/permitted_method_handler"
)

type IPermittedMethodMediator interface {
	AddPermittedMethod(permittedNetworkId string, permittedNetworkName string, chaincode string, channel string, inputArgs string, outputArgs string) (*permittedHandler.PermittedMethodInfo, error)
	PermittedMethodExists(id string) (bool, error)
	GetPermittedMethod(permittedMethodId string) (*permittedHandler.PermittedMethodInfo, error)
	GetPermittedMethodsByIndexAndAddress(permittedNetworkId string, startIndex string, endIndex string) ([]*permittedHandler.PermittedMethodInfo, error)
	GetPermittedMethodsByNetworkId(permittedNetworkId string) ([]*permittedHandler.PermittedMethodInfo, error)
	RemovePermittedMethod(id string) error
	UpdatePermittedMethod(id string, name string, chaincode string, channel string, inputArgs string, outputArgs string) error
	InvokePermittedMethod(name string, chaincode string, channel string, inputArgs []string) (*string, error)
}
