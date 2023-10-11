package chaincode

import (
	"encoding/json"
	"fmt"
	"github.com/hyperledger/fabric-contract-api-go/contractapi"
	"strconv"
)

var lastAccessibleNetworkId = "lastAccessibleNetworkId"

const AccessibleNetworkIdPrefix = "AccessibleNetwork"

type SmartContract struct {
	contractapi.Contract
}

// AccessibleNetworkInfo Struct related to the other networks who are accessible from the current network
type AccessibleNetworkInfo struct {
	NetworkName         string `json:"NetworkName"`
	IP                  string `json:"IP"`
	ADDRESS             string `json:"ADDRESS"`
	CompanyName         string `json:"CompanyName"`
	AccessibleNetworkId string `json:"AccessibleNetworkId"`
}

func accessibleNetworkInfoGenerator(networkName string, ip string, address string, companyName string, accessibleNetworkId string) AccessibleNetworkInfo {
	return AccessibleNetworkInfo{
		NetworkName:         networkName,
		IP:                  ip,
		ADDRESS:             address,
		CompanyName:         companyName,
		AccessibleNetworkId: accessibleNetworkId,
	}
}

func (s *SmartContract) InitLedger(ctx contractapi.TransactionContextInterface) error {

	ctx.GetStub().PutState(
		lastAccessibleNetworkId,
		[]byte("0"),
	)

	return nil
}

func (s *SmartContract) GenerateAccessibleNetworkId(ctx contractapi.TransactionContextInterface) string {
	lastID, err := ctx.GetStub().GetState(lastAccessibleNetworkId)
	if err != nil {
		fmt.Errorf("the error occured when trying to generate Accessible Network Id %s ", err)
		return ""
	}
	lastIDStr := string(lastID)
	LastIdInt, err := strconv.Atoi(lastIDStr)
	newLastIdStr := strconv.Itoa(LastIdInt + 1)
	ctx.GetStub().PutState(lastAccessibleNetworkId, []byte(newLastIdStr))
	return AccessibleNetworkIdPrefix + newLastIdStr
}

func (s *SmartContract) CreateAccessibleNetwork(ctx contractapi.TransactionContextInterface, networkName string, ip string, address string, companyName string) (*AccessibleNetworkInfo, error) {
	accessibleNetworkId := s.GenerateAccessibleNetworkId(ctx)
	accessibleNetworkInfo := accessibleNetworkInfoGenerator(networkName, ip, address, companyName, accessibleNetworkId)
	accessibleNetworkInfoJSON, err := json.Marshal(accessibleNetworkInfo)
	if err != nil {
		return nil, err
	}

	fmt.Printf("Creation Result %s \n", accessibleNetworkInfoJSON)
	return &accessibleNetworkInfo, ctx.GetStub().PutState(accessibleNetworkId, accessibleNetworkInfoJSON)
}

func (s *SmartContract) AccessibleNetworkExists(ctx contractapi.TransactionContextInterface, id string) (bool, error) {
	assetJSON, err := ctx.GetStub().GetState(id)
	if err != nil {
		return false, fmt.Errorf("the error occured when trying to check if Accessible Network Exists %s ", err)
	}
	return assetJSON != nil, nil
}

func (s *SmartContract) RemoveAccessibleNetwork(ctx contractapi.TransactionContextInterface, id string) error {
	exists, err := s.AccessibleNetworkExists(ctx, id)
	if err != nil {
		return err
	}
	if !exists {
		return fmt.Errorf("the accessible network %s does not exist", id)
	}

	return ctx.GetStub().DelState(id)
}

func (s *SmartContract) UpdateAccessibleNetwork(ctx contractapi.TransactionContextInterface, id string, networkName string, ip string, address string, companyName string) error {
	exists, err := s.AccessibleNetworkExists(ctx, id)
	if err != nil {
		return err
	}
	if !exists {
		return fmt.Errorf("the accessible network %s does not exist", id)
	}

	accessibleNetworkInfo := accessibleNetworkInfoGenerator(networkName, ip, address, companyName, id)
	accessibleNetworkInfoJSON, err := json.Marshal(accessibleNetworkInfo)
	if err != nil {
		return err
	}

	return ctx.GetStub().PutState(id, accessibleNetworkInfoJSON)
}

func (s *SmartContract) GetAccessibleNetwork(ctx contractapi.TransactionContextInterface, id string) (*AccessibleNetworkInfo, error) {
	assetJSON, err := ctx.GetStub().GetState(id)
	if err != nil {
		return nil, fmt.Errorf("the error occured when trying to get Accessible Network %s ", err)
	}
	if assetJSON == nil {
		return nil, fmt.Errorf("the accessible network %s does not exist", id)
	}

	var accessibleNetworkInfo AccessibleNetworkInfo
	err = json.Unmarshal(assetJSON, &accessibleNetworkInfo)
	if err != nil {
		return nil, err
	}

	return &accessibleNetworkInfo, nil
}

func (s *SmartContract) QueryAllAccessibleNetworks(ctx contractapi.TransactionContextInterface) ([]*AccessibleNetworkInfo, error) {
	startKey := AccessibleNetworkIdPrefix + "0"
	endKey := AccessibleNetworkIdPrefix + "999"
	resultsIterator, err := ctx.GetStub().GetStateByRange(startKey, endKey)
	if err != nil {
		return nil, err
	}
	defer resultsIterator.Close()
	results := []*AccessibleNetworkInfo{}
	for resultsIterator.HasNext() {
		queryResponse, err := resultsIterator.Next()
		if err != nil {
			return nil, err
		}
		permittedNetworkInfo := new(AccessibleNetworkInfo)
		_ = json.Unmarshal(queryResponse.Value, permittedNetworkInfo)
		results = append(results, permittedNetworkInfo)
	}
	return results, nil
}
