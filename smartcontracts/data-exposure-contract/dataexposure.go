package main

import (
	"data-exposure-contract/chaincode"
	"github.com/hyperledger/fabric-contract-api-go/contractapi"
	"log"
)

type MethodArg struct {
	ArgName string `json:"ArgName"`
	ArgType string `json:"ArgType"`
}
type Items struct {
	item string `json:"item"`
}

func main() {

	assetChaincode, err := contractapi.NewChaincode(&chaincode.SmartContract{})
	if err != nil {
		log.Panicf("Error creating data exposure chaincode: %v", err)
	}

	if err := assetChaincode.Start(); err != nil {
		log.Panicf("Error starting data exposure chaincode: %v", err)
	}
}
