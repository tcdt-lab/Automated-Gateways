package chaincode

import (
	"encoding/json"
	"fmt"
	"github.com/hyperledger/fabric-contract-api-go/contractapi"
)

var LastOutsiderID = "LAST_OUTSIDER_ID"

type SmartContract struct {
	contractapi.Contract
}

type OutsiderAsset struct {
	NAME     string `json:"Name"`
	IP       string `json:"IP"`
	ID       string `json:"ID"`
	USERNAME string `json:"USERNAME"`
	PASSWORD string `json:"PASSWORD"`
}

func (s *SmartContract) CreateOutsiderAsset(ctx contractapi.TransactionContextInterface, id string, name string, username string, ip string, password string) error {
	exists, err := s.OutsiderExists(ctx, id)
	if err != nil {
		return err
	}
	if exists {
		return fmt.Errorf("the outsider netwrok %s already exists", id)
	}

	asset := OutsiderAsset{
		ID:       id,
		IP:       ip,
		USERNAME: username,
		PASSWORD: password,
		NAME:     name,
	}
	assetJSON, err := json.Marshal(asset)
	if err != nil {
		return err
	}

	ctx.GetStub().PutState(LastOutsiderID, []byte(id))
	return ctx.GetStub().PutState(id, assetJSON)
}

func (s *SmartContract) OutsiderExists(ctx contractapi.TransactionContextInterface, id string) (bool, error) {
	assetJSON, err := ctx.GetStub().GetState(id)
	if err != nil {
		return false, fmt.Errorf("failed to read from world state: %v", err)
	}
	return assetJSON != nil, nil
}

func (s *SmartContract) GetOutsiderAsset(ctx contractapi.TransactionContextInterface, id string) (*OutsiderAsset, error) {
	assetJSON, err := ctx.GetStub().GetState(id)
	if err != nil {
		return nil, fmt.Errorf("failed to read from world state: %v", err)
	}
	if assetJSON == nil {
		return nil, fmt.Errorf("the outsiderAsset %s does not exist", id)
	}

	var outsiderAsset OutsiderAsset
	err = json.Unmarshal(assetJSON, &outsiderAsset)
	if err != nil {
		return nil, err
	}

	return &outsiderAsset, nil
}



func (s *SmartContract) UpdateOutsiderAsset(ctx contractapi.TransactionContextInterface, id string, name string, username string, ip string, password string) error {
	exists, err := s.OutsiderExists(ctx, id)
	if err != nil {
		return err
	}
	if !exists {
		return fmt.Errorf("the asset %s does not exist", id)
	}

	asset := OutsiderAsset{
		ID:       id,
		IP:       ip,
		USERNAME: username,
		PASSWORD: password,
		NAME:     name,
	}
	assetJSON, err := json.Marshal(asset)
	if err != nil {
		return err
	}

	return ctx.GetStub().PutState(id, assetJSON)
}

func (s *SmartContract) DeleteOutsiderAsset(ctx contractapi.TransactionContextInterface, id string) error {
	exists, err := s.OutsiderExists(ctx, id)
	if err != nil {
		return err
	}
	if !exists {
		return fmt.Errorf("the asset %s does not exist", id)
	}

	return ctx.GetStub().DelState(id)
}

func (s *SmartContract) GetLastOutsiderID(ctx contractapi.TransactionContextInterface) (string, error) {
	lastID, err := ctx.GetStub().GetState(LastOutsiderID)
	if err != nil {
		return "", fmt.Errorf("failed to read from world state: %v", err)
	}
	if lastID == nil {
		return "", fmt.Errorf("the Last ID %s does not exist")
	}

	lastIDStr := string(lastID)
	return lastIDStr, nil
}
