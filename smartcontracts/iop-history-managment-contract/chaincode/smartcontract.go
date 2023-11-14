package chaincode

import "github.com/hyperledger/fabric-contract-api-go/contractapi"

const PermittedMethodIdPrefix = "Iop_History"
const insideInitiatedCallPrefix = "Inside_initiated_call"

type SmartContract struct {
	contractapi.Contract
}
