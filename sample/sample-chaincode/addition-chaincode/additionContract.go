package main

import (
	"addition-chaincode/chaincode"
	"github.com/hyperledger/fabric-contract-api-go/contractapi"
	"log"
)

func main() {
	assetChaincode, err := contractapi.NewChaincode(&chaincode.SmartContract{})
	if err != nil {
		log.Panicf("Error creating managment chaincode: %v", err)
	}

	if err := assetChaincode.Start(); err != nil {
		log.Panicf("Error starting managment chaincode: %v", err)
	}
}
