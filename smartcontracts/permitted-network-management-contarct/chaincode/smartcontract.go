package chaincode

import (
	"encoding/json"
	"fmt"
	"github.com/hyperledger/fabric-contract-api-go/contractapi"
	"strconv"
)

var lastPermittedNetworkId = "lastPermittedNetworkId"

const PermittedNetworkIdPrefix = "PermittedNetwork"

type SmartContract struct {
	contractapi.Contract
}

type PermittedNetworkInfo struct {
	NetworkName        string `json:"NetworkName"`
	IP                 string `json:"IP"`
	ADDRESS            string `json:"ADDRESS"`
	CompanyName        string `json:"CompanyName"`
	PermittedNetworkId string `json:"PermittedNetworkId"`
}

func permittedNetworkInfoGenerator(networkName string, ip string, address string, companyName string, permittedNetworkId string) PermittedNetworkInfo {
	return PermittedNetworkInfo{
		NetworkName:        networkName,
		IP:                 ip,
		ADDRESS:            address,
		CompanyName:        companyName,
		PermittedNetworkId: permittedNetworkId,
	}
}
func (s *SmartContract) InitLedger(ctx contractapi.TransactionContextInterface) error {
	logger(false, "InitLedger", "METHOD START...")
	ctx.GetStub().PutState(
		lastPermittedNetworkId,
		[]byte("0"),
	)
	logger(false, "InitLedger", "METHOD END")
	return nil
}

func (s *SmartContract) GeneratePermittedNetworkId(ctx contractapi.TransactionContextInterface, address string) string {
	logger(false, "GeneratePermittedNetworkId", "METHOD START...")
	lastID, err := ctx.GetStub().GetState(lastPermittedNetworkId)
	if err != nil {
		logger(true, "GeneratePermittedNetworkId", err.Error())
		return ""
	}
	lastIDStr := string(lastID)
	LastIdInt, err := strconv.Atoi(lastIDStr)
	newLastIdStr := strconv.Itoa(LastIdInt + 1)
	ctx.GetStub().PutState(lastPermittedNetworkId, []byte(newLastIdStr))
	newPermittedID := PermittedNetworkIdPrefix + newLastIdStr
	logger(false, "GeneratePermittedNetworkId", "METHOD END WITH new Permitted ID: "+newPermittedID+"...")
	return newPermittedID
}

func (s *SmartContract) CreatePermittedNetwork(ctx contractapi.TransactionContextInterface, networkName string, ip string, address string, companyName string) (*PermittedNetworkInfo, error) {
	logger(false, "CreatePermittedNetwork", "METHOD START...")
	permittedNetworkId := s.GeneratePermittedNetworkId(ctx, address)
	permittedNetworkInfo := permittedNetworkInfoGenerator(networkName, ip, address, companyName, permittedNetworkId)
	permittedNetworkInfoAsBytes, err := json.Marshal(permittedNetworkInfo)
	if err != nil {
		logger(true, "CreatePermittedNetwork", err.Error())
		return nil, err
	}
	logger(false, "CreatePermittedNetwork", "METHOD END with permittedNetworkInfo: "+string(permittedNetworkInfoAsBytes)+"...")
	err = ctx.GetStub().PutState(permittedNetworkId, permittedNetworkInfoAsBytes)
	if err != nil {
		logger(true, "CreatePermittedNetwork", err.Error())
		return nil, err
	}
	return &permittedNetworkInfo, nil
}

func (s *SmartContract) GetPermittedNetwork(ctx contractapi.TransactionContextInterface, permittedNetworkId string) (*PermittedNetworkInfo, error) {
	logger(false, "GetPermittedNetwork", "METHOD START with Id "+permittedNetworkId+" as input ...")
	permittedNetworkInfoAsBytes, err := ctx.GetStub().GetState(permittedNetworkId)
	if err != nil {
		logger(true, "GetPermittedNetwork", "METHOD END with error: "+err.Error())
		return nil, fmt.Errorf("Failed to read from world state. %s", err.Error())
	}
	if permittedNetworkInfoAsBytes == nil {
		logger(true, "GetPermittedNetwork", "METHOD END with error: "+permittedNetworkId+" does not exist...")
		return nil, fmt.Errorf("%s does not exist", permittedNetworkId)
	}
	permittedNetworkInfo := new(PermittedNetworkInfo)
	_ = json.Unmarshal(permittedNetworkInfoAsBytes, permittedNetworkInfo)
	logger(false, "GetPermittedNetwork", "METHOD END with permittedNetworkInfo: "+string(permittedNetworkInfoAsBytes)+"...")
	return permittedNetworkInfo, nil
}

func (s *SmartContract) GetPermittedNetworksByIndex(ctx contractapi.TransactionContextInterface, startIndex string, endIndex string) ([]*PermittedNetworkInfo, error) {
	logger(false, "GetPermittedNetworks", "METHOD START with startIndex "+startIndex+" and endIndex "+endIndex+" as input ...")
	startKey := PermittedNetworkIdPrefix + startIndex
	endKey := PermittedNetworkIdPrefix + endIndex
	resultsIterator, err := ctx.GetStub().GetStateByRange(startKey, endKey)
	if err != nil {
		logger(true, "GetPermittedNetworks", "METHOD END with error: "+err.Error())
		return nil, err
	}
	defer resultsIterator.Close()
	results := []*PermittedNetworkInfo{}
	for resultsIterator.HasNext() {
		queryResponse, err := resultsIterator.Next()
		if err != nil {
			logger(true, "GetPermittedNetworks", "METHOD END with error in loop : "+err.Error())
			return nil, err
		}
		permittedNetworkInfo := new(PermittedNetworkInfo)
		_ = json.Unmarshal(queryResponse.Value, permittedNetworkInfo)
		results = append(results, permittedNetworkInfo)
	}
	logger(false, "GetPermittedNetworks", "METHOD END with successful results")
	return results, nil
}

func (s *SmartContract) PermittedNetworkExists(ctx contractapi.TransactionContextInterface, id string) (bool, error) {
	logger(false, "PermittedNetworkExists", "METHOD START with id "+id+" as input ...")
	permittedNetworkInfoAsBytes, err := ctx.GetStub().GetState(id)
	if err != nil {
		logger(true, "PermittedNetworkExists", "METHOD END with error: "+err.Error())
		return false, fmt.Errorf("Failed to read from world state. %s", err.Error())
	}
	logger(false, "PermittedNetworkExists", "METHOD END with existed permittedNetworkInfo: "+string(permittedNetworkInfoAsBytes)+" and the result is : "+strconv.FormatBool(permittedNetworkInfoAsBytes != nil)+" ...")
	return permittedNetworkInfoAsBytes != nil, nil
}

func (s *SmartContract) RemovePermittedNetwork(ctx contractapi.TransactionContextInterface, id string) error {
	logger(false, "RemovePermittedNetwork", "METHOD START with id "+id+" as input ...")
	exists, err := s.PermittedNetworkExists(ctx, id)
	if err != nil {
		logger(true, "RemovePermittedNetwork", "METHOD END with error: "+err.Error())
		return err
	}
	if !exists {
		logger(true, "RemovePermittedNetwork", "METHOD END with error: "+id+" does not exist...")
		return fmt.Errorf("the accessible network %s does not exist", id)
	}
	logger(false, "RemovePermittedNetwork", "METHOD END with success")
	return ctx.GetStub().DelState(id)
}

func (s *SmartContract) UpdatePermittedNetwork(ctx contractapi.TransactionContextInterface, id string, networkName string, ip string, address string, companyName string) error {
	logger(false, "UpdatePermittedNetwork", "METHOD START with id "+id+",networkName: "+networkName+" ,ip: "+ip+" ,address: "+address+",companyName: "+companyName+"   as input ...")
	exists, err := s.PermittedNetworkExists(ctx, id)
	if err != nil {
		logger(true, "UpdatePermittedNetwork", "METHOD END with error: "+err.Error())
		return err
	}
	if !exists {
		logger(true, "UpdatePermittedNetwork", "METHOD END with error: "+id+" does not exist...")
		return fmt.Errorf("the accessible network %s does not exist", id)
	}
	permittedNetworkInfo := permittedNetworkInfoGenerator(networkName, ip, address, companyName, id)
	permittedNetworkInfoJSON, err := json.Marshal(permittedNetworkInfo)
	if err != nil {
		return err
	}
	logger(false, "UpdatePermittedNetwork", "METHOD END with success")
	return ctx.GetStub().PutState(id, permittedNetworkInfoJSON)
}

func logger(isError bool, methodName string, txt string) {
	if isError {
		fmt.Println("***Error at  METHOD_NAME:", methodName, " TEXT: ", txt)
		return
	}
	fmt.Println("METHOD_NAME: ", methodName, " TEXT: ", txt)
}

func (s *SmartContract) GetPermittedNetworks(ctx contractapi.TransactionContextInterface) ([]*PermittedNetworkInfo, error) {
	logger(false, "GetPermittedNetworks", "METHOD STARTS ...")
	startKey := PermittedNetworkIdPrefix + "0"
	endKey := PermittedNetworkIdPrefix + "999999"
	resultsIterator, err := ctx.GetStub().GetStateByRange(startKey, endKey)
	if err != nil {
		logger(true, "GetPermittedNetworks", "METHOD END with error: "+err.Error())
		return nil, err
	}
	defer resultsIterator.Close()
	results := []*PermittedNetworkInfo{}
	for resultsIterator.HasNext() {
		queryResponse, err := resultsIterator.Next()
		if err != nil {
			logger(true, "GetPermittedNetworks", "METHOD END with error in loop : "+err.Error())
			return nil, err
		}
		permittedNetworkInfo := new(PermittedNetworkInfo)
		_ = json.Unmarshal(queryResponse.Value, permittedNetworkInfo)
		results = append(results, permittedNetworkInfo)
	}
	logger(false, "GetPermittedNetworks", "METHOD END with successful results")
	return results, nil
}
