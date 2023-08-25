package chaincode

import (
	"encoding/json"
	"fmt"
	"github.com/hyperledger/fabric-contract-api-go/contractapi"
)

var methodIdConst = "METHOD_ID_"

var outsiderMethodIdConst = "OUTSIDER_METHOD_ID_"

type SmartContract struct {
	contractapi.Contract
}

type ExposedMethodAsset struct {
	ChannelName   string `json:"ChannelName"`
	ChaincodeName string `json:"ChaincodeName"`
	MethodName    string `json:"MethodName"`
	MethodID      string `json:"methodID"`
}

type MethodsOwnerAsset struct {
	OutsiderID string               `json:"OutsiderID"`
	Methods    []ExposedMethodAsset `json:"Methods"`
}

func (s *SmartContract) ExposeMethod(ctx contractapi.TransactionContextInterface, methodID string, methodName string, chaincodeName string, channelName string) error {
	exists, err := s.AssetExists(ctx, methodIdConst+methodID)
	if err != nil {
		return err
	}
	if exists {
		return fmt.Errorf("the method %s already exposed", methodName)
	}

	exposedMethodAsset := ExposedMethodAsset{

		MethodName:    methodName,
		MethodID:      methodID,
		ChannelName:   channelName,
		ChaincodeName: chaincodeName,
	}
	assetJSON, err := json.Marshal(exposedMethodAsset)
	if err != nil {
		return err
	}

	return ctx.GetStub().PutState(methodIdConst+methodID, assetJSON)
}

func (s *SmartContract) AssetExists(ctx contractapi.TransactionContextInterface, id string) (bool, error) {
	assetJSON, err := ctx.GetStub().GetState(id)
	if err != nil {
		return false, fmt.Errorf("failed to read from world state: %v", err)
	}

	return assetJSON != nil, nil
}

func (s *SmartContract) GetMethodAsset(ctx contractapi.TransactionContextInterface, id string) (*ExposedMethodAsset, error) {
	assetJSON, err := ctx.GetStub().GetState(id)
	if err != nil {
		return nil, fmt.Errorf("failed to read from world state: %v", err)
	}
	if assetJSON == nil {
		return nil, fmt.Errorf("the method with ID %s does not exist", id)
	}

	var asset ExposedMethodAsset
	err = json.Unmarshal(assetJSON, &asset)
	if err != nil {
		return nil, err
	}

	return &asset, nil
}

func (s *SmartContract) GetMethodsOwnerAsset(ctx contractapi.TransactionContextInterface, id string) (*MethodsOwnerAsset, error) {
	assetJSON, err := ctx.GetStub().GetState(id)
	if err != nil {
		return nil, fmt.Errorf("failed to read from world state: %v", err)
	}
	if assetJSON == nil {
		return nil, fmt.Errorf("the method with ID %s does not exist", id)
	}

	var asset MethodsOwnerAsset
	err = json.Unmarshal(assetJSON, &asset)
	if err != nil {
		return nil, err
	}

	return &asset, nil
}

func (s *SmartContract) AssignMethodToOutsider(ctx contractapi.TransactionContextInterface, methodID string, outsiderId string) error {
	exists, err := s.AssetExists(ctx, outsiderId)
	if err != nil {
		return err
	}
	if exists {
		assetMethod, _err := s.GetMethodAsset(ctx, methodID)

		if _err != nil {
			return _err
		}
		methodsOwnerAsset, __err := s.GetMethodsOwnerAsset(ctx, outsiderId)
		if __err != nil {
			return __err
		}

		methodsOwnerAsset.Methods = append(methodsOwnerAsset.Methods, *assetMethod)
		assetJSON, err := json.Marshal(methodsOwnerAsset)
		if err != nil {
			return err
		}
		return ctx.GetStub().PutState(outsiderMethodIdConst+outsiderId, assetJSON)

	}
	assetMethod, err := s.GetMethodAsset(ctx, methodID)

	var assetMethodsList []ExposedMethodAsset
	assetMethodsList = append(assetMethodsList, *assetMethod)
	outsiderMethodAsset := MethodsOwnerAsset{
		Methods:    assetMethodsList,
		OutsiderID: outsiderId,
	}
	assetJSON, err := json.Marshal(outsiderMethodAsset)
	if err != nil {
		return err
	}

	return ctx.GetStub().PutState(outsiderMethodIdConst+outsiderId, assetJSON)
}

func (s *SmartContract) removeAccessToMethod(ctx contractapi.TransactionContextInterface, methodID string, outsiderId string) error {
	methodsOwnerAsset, err := s.GetMethodsOwnerAsset(ctx, outsiderId)
	if err != nil {
		return err
	}

	for i := 0; i < (len(methodsOwnerAsset.Methods) - 1); {
		if methodsOwnerAsset.Methods[i].MethodID == methodID {
			methodsOwnerAsset.Methods = append(methodsOwnerAsset.Methods[:i], methodsOwnerAsset.Methods[i+1:]...)
			break
		}

	}

	assetJSON, err := json.Marshal(methodsOwnerAsset)
	if err != nil {
		return err
	}

	return ctx.GetStub().PutState(outsiderId, assetJSON)
}

func (s *SmartContract) removeOutsiderWholeAccess(ctx contractapi.TransactionContextInterface, outsiderId string) error {
	exists, err := s.AssetExists(ctx, outsiderId)
	if err != nil {
		return err
	}
	if !exists {
		return fmt.Errorf("the outsider ID  %s does not exist", outsiderId)
	}

	return ctx.GetStub().DelState(outsiderId)
}
