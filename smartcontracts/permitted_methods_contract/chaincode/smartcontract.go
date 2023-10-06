package chaincode

import (
	"encoding/json"
	"fmt"
	"github.com/hyperledger/fabric-contract-api-go/contractapi"
	"strconv"
)

var lastPermittedMethodId = "lastPermittedMethodId"

const PermittedMethodIdPrefix = "permittedMethod"

type SmartContract struct {
	contractapi.Contract
}

type PermittedMethodInfo struct {
	Id         string `json:"Id"`
	Name       string `json:"Name"`
	Chaincode  string `json:"Chaincode"`
	Channel    string `json:"Channel"`
	InputArgs  string `json:"inputArgs"`
	OutputType string `json:"outputArgs"`
}

func (s *SmartContract) InitLedger(ctx contractapi.TransactionContextInterface) error {
	logger(false, "InitLedger", "METHOD START...")
	ctx.GetStub().PutState(
		lastPermittedMethodId,
		[]byte("0"),
	)
	logger(false, "InitLedger", "METHOD END")
	return nil
}

func permittedMethodInfoGenerator(id string, name string, chaincode string, channel string, inputArgs string, outputArgs string) PermittedMethodInfo {
	return PermittedMethodInfo{
		Id:         id,
		Name:       name,
		Chaincode:  chaincode,
		Channel:    channel,
		InputArgs:  inputArgs,
		OutputType: outputArgs,
	}
}
func (s *SmartContract) GeneratePermittedMethodId(ctx contractapi.TransactionContextInterface, permittedNetworkId string) string {
	logger(false, "GeneratePermittedMethodId", "METHOD START...")
	lastID, err := ctx.GetStub().GetState(lastPermittedMethodId)
	if err != nil {
		logger(true, "GeneratePermittedMethodId", err.Error())
		return ""
	}
	lastIDStr := string(lastID)
	newLastIdStr := fmt.Sprintf("%d", lastIDStr)
	ctx.GetStub().PutState(lastPermittedMethodId, []byte(newLastIdStr))
	newPermittedID := PermittedMethodIdPrefix + permittedNetworkId + "_" + newLastIdStr
	logger(false, "GeneratePermittedMethodId", "METHOD END WITH new Permitted ID: "+newPermittedID+"...")
	return newPermittedID
}

func (s *SmartContract) AddPermittedMethod(ctx contractapi.TransactionContextInterface, permittedNetworkId string, name string, chaincode string, channel string, inputArgs string, outputArgs string) error {
	logger(false, "AddPermittedMethod", "METHOD START...")
	permittedMethodId := s.GeneratePermittedMethodId(ctx, permittedNetworkId)
	permittedMethodInfo := permittedMethodInfoGenerator(permittedMethodId, name, chaincode, channel, inputArgs, outputArgs)
	permittedMethodInfoAsBytes, err := json.Marshal(permittedMethodInfo)
	if err != nil {
		logger(true, "AddPermittedMethod", err.Error())
		return err
	}
	err = ctx.GetStub().PutState(permittedMethodId, permittedMethodInfoAsBytes)
	if err != nil {
		logger(true, "AddPermittedMethod", err.Error())
		return err
	}
	logger(false, "AddPermittedMethod", "METHOD END...")
	return nil
}

func (s *SmartContract) GetPermittedMethod(ctx contractapi.TransactionContextInterface, permittedMethodId string) (*PermittedMethodInfo, error) {
	logger(false, "GetPermittedMethod", "METHOD START with Id "+permittedMethodId+" as input ...")
	permittedMethodInfoAsBytes, err := ctx.GetStub().GetState(permittedMethodId)
	if err != nil {
		logger(true, "GetPermittedMethod", "METHOD END with error: "+err.Error())
		return nil, err
	}
	permittedMethodInfo := new(PermittedMethodInfo)
	err = json.Unmarshal(permittedMethodInfoAsBytes, permittedMethodInfo)
	if err != nil {
		logger(true, "GetPermittedMethod", "METHOD END with error: "+err.Error())
		return nil, err
	}
	logger(false, "GetPermittedMethod", "METHOD END with permittedMethodInfo: "+string(permittedMethodInfoAsBytes)+"...")
	return permittedMethodInfo, nil
}

func (s *SmartContract) GetPermittedMethods(ctx contractapi.TransactionContextInterface, permittedNetworkId string, startIndex string, endIndex string) ([]*PermittedMethodInfo, error) {
	logger(false, "GetPermittedMethods", "METHOD START with startIndex "+startIndex+" and endIndex "+endIndex+" as input ...")
	startKey := PermittedMethodIdPrefix + permittedNetworkId + "_" + startIndex
	endKey := PermittedMethodIdPrefix + permittedNetworkId + "_" + endIndex
	resultsIterator, err := ctx.GetStub().GetStateByRange(startKey, endKey)
	if err != nil {
		logger(true, "GetPermittedMethods", "METHOD END with error: "+err.Error())
		return nil, err
	}
	defer resultsIterator.Close()
	results := []*PermittedMethodInfo{}
	for resultsIterator.HasNext() {
		queryResponse, err := resultsIterator.Next()
		if err != nil {
			logger(true, "GetPermittedMethods", "METHOD END with error in loop : "+err.Error())
			return nil, err
		}
		permittedMethodInfo := new(PermittedMethodInfo)
		_ = json.Unmarshal(queryResponse.Value, permittedMethodInfo)
		results = append(results, permittedMethodInfo)
	}
	logger(false, "GetPermittedMethods", "METHOD END with successful results")
	return results, nil
}

func (s *SmartContract) PermittedMethodExists(ctx contractapi.TransactionContextInterface, id string) (bool, error) {
	logger(false, "PermittedMethodExists", "METHOD START with id "+id+" as input ...")
	permittedMethodInfoAsBytes, err := ctx.GetStub().GetState(id)
	if err != nil {
		logger(true, "PermittedMethodExists", "METHOD END with error: "+err.Error())
		return false, fmt.Errorf("Failed to read from world state. %s", err.Error())
	}
	logger(false, "PermittedMethodExists", "METHOD END with existed permittedMethodInfo: "+string(permittedMethodInfoAsBytes)+" and the result is : "+strconv.FormatBool(permittedMethodInfoAsBytes != nil)+" ...")
	return permittedMethodInfoAsBytes != nil, nil
}

func (s *SmartContract) RemovePermittedMethod(ctx contractapi.TransactionContextInterface, id string) error {
	logger(false, "RemovePermittedMethod", "METHOD START with id "+id+" as input ...")
	exists, err := s.PermittedMethodExists(ctx, id)
	if err != nil {
		logger(true, "RemovePermittedMethod", "METHOD END with error: "+err.Error())
		return err
	}
	if !exists {
		logger(true, "RemovePermittedMethod", "METHOD END with error: "+id+" does not exist...")
		return fmt.Errorf("%s does not exist", id)
	}
	err = ctx.GetStub().DelState(id)
	if err != nil {
		logger(true, "RemovePermittedMethod", "METHOD END with error: "+err.Error())
		return err
	}
	logger(false, "RemovePermittedMethod", "METHOD END with success...")
	return nil
}

func (s *SmartContract) UpdatePermittedMethod(ctx contractapi.TransactionContextInterface, id string, name string, chaincode string, channel string, inputArgs string, outputArgs string) error {
	logger(false, "UpdatePermittedMethod", "METHOD START with id "+id+" as input ...")
	exists, err := s.PermittedMethodExists(ctx, id)
	if err != nil {
		logger(true, "UpdatePermittedMethod", "METHOD END with error: "+err.Error())
		return err
	}
	if !exists {
		logger(true, "UpdatePermittedMethod", "METHOD END with error: "+id+" does not exist...")
		return fmt.Errorf("%s does not exist", id)
	}
	permittedMethodInfo := permittedMethodInfoGenerator(id, name, chaincode, channel, inputArgs, outputArgs)
	permittedMethodInfoAsBytes, err := json.Marshal(permittedMethodInfo)
	if err != nil {
		logger(true, "UpdatePermittedMethod", "METHOD END with error: "+err.Error())
		return err
	}
	err = ctx.GetStub().PutState(id, permittedMethodInfoAsBytes)
	if err != nil {
		logger(true, "UpdatePermittedMethod", "METHOD END with error: "+err.Error())
		return err
	}
	logger(false, "UpdatePermittedMethod", "METHOD END with success...")
	return nil
}

func logger(isError bool, methodName string, txt string) {
	if isError {
		fmt.Println("***Error at  METHOD_NAME:", methodName, " TEXT: ", txt)
		return
	}
	fmt.Println("METHOD_NAME: ", methodName, " TEXT: ", txt)
}

func (s *SmartContract) GetPermittedMethodsByNetworkId(ctx contractapi.TransactionContextInterface, permittedNetworkId string) ([]*PermittedMethodInfo, error) {
	logger(false, "GetPermittedMethodsByNetworkId", "METHOD START with permittedNetworkId "+permittedNetworkId+" as input ...")
	startKey := PermittedMethodIdPrefix + permittedNetworkId + "_0"
	endKey := PermittedMethodIdPrefix + permittedNetworkId + "_9999999999999999999"
	resultsIterator, err := ctx.GetStub().GetStateByRange(startKey, endKey)
	if err != nil {
		logger(true, "GetPermittedMethodsByNetworkId", "METHOD END with error: "+err.Error())
		return nil, err
	}
	defer resultsIterator.Close()
	results := []*PermittedMethodInfo{}
	for resultsIterator.HasNext() {
		queryResponse, err := resultsIterator.Next()
		if err != nil {
			logger(true, "GetPermittedMethodsByNetworkId", "METHOD END with error in loop : "+err.Error())
			return nil, err
		}
		permittedMethodInfo := new(PermittedMethodInfo)
		_ = json.Unmarshal(queryResponse.Value, permittedMethodInfo)
		results = append(results, permittedMethodInfo)
	}
	logger(false, "GetPermittedMethodsByNetworkId", "METHOD END with successful results")
	return results, nil
}
