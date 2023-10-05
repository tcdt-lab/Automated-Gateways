package main

import (
	"github.com/hyperledger/fabric-contract-api-go/contractapi"
	"log"
	"permitted-network-management-contarct/chaincode"
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
