package main

import (
	"data-exposure-contract/chaincode"
	"github.com/hyperledger/fabric-contract-api-go/contractapi"
	"log"
)

func main() {
	assetChaincode, err := contractapi.NewChaincode(&chaincode.SmartContract{})
	if err != nil {
		log.Panicf("Error creating data exposure chaincode: %v", err)
	}

	if err := assetChaincode.Start(); err != nil {
		log.Panicf("Error starting data exposure chaincode: %v", err)
	}
}
