package chaincode

import (
	"encoding/json"
	"fmt"
	"github.com/hyperledger/fabric-contract-api-go/contractapi"
	"strconv"
)

const IOP_HISTORY_PREFIX = "Iop_History"
const INSIDE_INITIATED_CALL_PREFIX = "Inside_initiated_call"
const OUTSIDE_INTIATED_CALL_PREFIX = "Outside_initiated_call"

type SmartContract struct {
	contractapi.Contract
}
type HistoryLog struct {
	Id            string `json:"id"`
	CallType      string `json:"callType"`
	CallerIP      string `json:"callerIP"`
	CallerAddress string `json:"callerAddress"`
	MethodName    string `json:"methodName"`
	InputArgs     string `json:"inputArgs"`
	OutputArgs    string `json:"outputArgs"`
	Date          string `json:"date"`
	Hour          string `json:"hour"`
	Day           string `json:"day"`
	Month         string `json:"month"`
	Year          string `json:"year"`
	Distinctive   string `json:"distinctive"`
}

func (smartContract *SmartContract) createKey(ctx contractapi.TransactionContextInterface, callType string, callerIP string, callerAddress string, methodName string, year string, month string, day string, hour string, distinctive string) string {
	logger(false, "createKey", "METHOD START...")
	compositeKey, err := ctx.GetStub().CreateCompositeKey(IOP_HISTORY_PREFIX, []string{callType, callerIP, callerAddress, methodName, year, month, day, hour, distinctive})
	if err != nil {
		logger(true, "createKey", err.Error())
		return ""
	}
	logger(false, "createKey", "METHOD END WITH compositeKey: "+compositeKey)
	return compositeKey
}
func (s *SmartContract) InsertInsideInitiatedCall(ctx contractapi.TransactionContextInterface, callerIP string, callerAddress string, methodName string, inputArgs string, outputArgs string) (*HistoryLog, error) {
	logger(false, "InsertInsideInitiatedCall", "METHOD START...")
	date, err := ctx.GetStub().GetTxTimestamp()
	dateString := date.AsTime().String()
	year := strconv.Itoa(date.AsTime().Year())
	month := strconv.Itoa(int(date.AsTime().Month()))
	day := strconv.Itoa(date.AsTime().Day())
	hour := strconv.Itoa(date.AsTime().Hour())
	distinctivePart := strconv.Itoa(date.AsTime().Nanosecond())

	logger(false, "InsertInsideInitiatedCall", "dateString: "+dateString)
	logger(false, "InsertInsideInitiatedCall", "year: "+year)
	logger(false, "InsertInsideInitiatedCall", "month: "+month)
	logger(false, "InsertInsideInitiatedCall", "day: "+day)
	logger(false, "InsertInsideInitiatedCall", "hour: "+hour)
	logger(false, "InsertInsideInitiatedCall", "distinctive: "+distinctivePart)
	if err != nil {
		logger(true, "InsertInsideInitiatedCall", err.Error())
		return nil, err
	}

	key := s.createKey(ctx, INSIDE_INITIATED_CALL_PREFIX, callerIP, callerAddress, methodName, year, month, day, hour, distinctivePart)
	historyLog := HistoryLog{
		Id:            key,
		CallType:      INSIDE_INITIATED_CALL_PREFIX,
		CallerIP:      callerIP,
		CallerAddress: callerAddress,
		MethodName:    methodName,
		InputArgs:     inputArgs,
		OutputArgs:    outputArgs,
		Date:          dateString,
		Hour:          hour,
		Day:           day,
		Month:         month,
		Year:          year,
		Distinctive:   distinctivePart,
	}
	historyLogAsBytes, err := json.Marshal(historyLog)
	if err != nil {
		logger(true, "InsertInsideInitiatedCall", err.Error())
		return nil, err
	}
	err = ctx.GetStub().PutState(key, historyLogAsBytes)
	if err != nil {
		logger(true, "InsertInsideInitiatedCall", err.Error())
		return nil, err
	}
	logger(false, "InsertInsideInitiatedCall", "METHOD END")
	return &historyLog, nil
}

func (s *SmartContract) InsertOutsideInitiatedCall(ctx contractapi.TransactionContextInterface, callerIP string, callerAddress string, methodName string, inputArgs string, outputArgs string) (*HistoryLog, error) {
	logger(false, "InsertOutsideInitiatedCall", "METHOD START...")
	date, err := ctx.GetStub().GetTxTimestamp()
	dateString := date.AsTime().String()
	year := strconv.Itoa(date.AsTime().Year())
	month := strconv.Itoa(int(date.AsTime().Month()))
	day := strconv.Itoa(date.AsTime().Day())
	hour := strconv.Itoa(date.AsTime().Hour())
	distinctivePart := strconv.Itoa(date.AsTime().Nanosecond())
	if err != nil {
		logger(true, "InsertOutsideInitiatedCall", err.Error())
		return nil, err
	}
	key := s.createKey(ctx, OUTSIDE_INTIATED_CALL_PREFIX, callerIP, callerAddress, methodName, year, month, day, hour, distinctivePart)
	historyLog := HistoryLog{
		Id:            key,
		CallType:      OUTSIDE_INTIATED_CALL_PREFIX,
		CallerIP:      callerIP,
		CallerAddress: callerAddress,
		MethodName:    methodName,
		InputArgs:     inputArgs,
		OutputArgs:    outputArgs,
		Date:          dateString,
		Year:          year,
		Month:         month,
		Day:           day,
		Hour:          hour,
		Distinctive:   distinctivePart,
	}
	historyLogAsBytes, err := json.Marshal(historyLog)
	if err != nil {
		logger(true, "InsertOutsideInitiatedCall", err.Error())
		return nil, err
	}
	err = ctx.GetStub().PutState(key, historyLogAsBytes)
	if err != nil {
		logger(true, "InsertOutsideInitiatedCall", err.Error())
		return nil, err
	}
	logger(false, "InsertOutsideInitiatedCall", "METHOD END")
	return &historyLog, nil
}

func (s *SmartContract) GetHistoryBasedOnCallType(ctx contractapi.TransactionContextInterface, callType string) ([]*HistoryLog, error) {
	logger(false, "GetHistoryLogsBAsedOnCallType", "METHOD START...")
	resultsIterator, err := ctx.GetStub().GetStateByPartialCompositeKey(IOP_HISTORY_PREFIX, []string{callType})
	if err != nil {
		logger(true, "GetHistoryLogsBAsedOnCallType", err.Error())
		return nil, err
	}
	defer resultsIterator.Close()
	var historyLogs []*HistoryLog
	for resultsIterator.HasNext() {
		queryResponse, err := resultsIterator.Next()
		if err != nil {
			logger(true, "GetHistoryLogsBAsedOnCallType", err.Error())
			return nil, err
		}
		historyLog := new(HistoryLog)
		err = json.Unmarshal(queryResponse.Value, historyLog)
		if err != nil {
			logger(true, "GetHistoryLogsBAsedOnCallType", err.Error())
			return nil, err
		}
		historyLogs = append(historyLogs, historyLog)
	}
	logger(false, "GetHistoryLogsBAsedOnCallType", "METHOD END")
	return historyLogs, nil
}

func (s *SmartContract) GetHistoryBasedOnCallerIP(ctx contractapi.TransactionContextInterface, callerType, callerIP string) ([]*HistoryLog, error) {
	logger(false, "GetHistoryLogsBasedOnCallerIP", "METHOD START...")
	resultsIterator, err := ctx.GetStub().GetStateByPartialCompositeKey(IOP_HISTORY_PREFIX, []string{callerType, callerIP})
	if err != nil {
		logger(true, "GetHistoryLogsBasedOnCallerIP", err.Error())
		return nil, err
	}
	defer resultsIterator.Close()
	var historyLogs []*HistoryLog
	for resultsIterator.HasNext() {
		queryResponse, err := resultsIterator.Next()
		if err != nil {
			logger(true, "GetHistoryLogsBasedOnCallerIP", err.Error())
			return nil, err
		}
		historyLog := new(HistoryLog)
		err = json.Unmarshal(queryResponse.Value, historyLog)
		if err != nil {
			logger(true, "GetHistoryLogsBasedOnCallerIP", err.Error())
			return nil, err
		}
		logger(false, "GetHistoryLogsBasedOnCallerIP", "History Log:"+(historyLog).CallerIP)
		historyLogs = append(historyLogs, historyLog)
	}
	logger(false, "GetHistoryLogsBasedOnCallerIP", "METHOD END")
	return historyLogs, nil
}

func (s *SmartContract) GetHistoryBasedOnCallerAddress(ctx contractapi.TransactionContextInterface, callerType string, callerIP string, callerAddress string) ([]*HistoryLog, error) {
	logger(false, "GetHistoryLogsBasedOnCallerName", "METHOD START...")
	resultsIterator, err := ctx.GetStub().GetStateByPartialCompositeKey(IOP_HISTORY_PREFIX, []string{callerType, callerIP, callerAddress})
	if err != nil {
		logger(true, "GetHistoryLogsBasedOnCallerName", err.Error())
		return nil, err
	}
	defer resultsIterator.Close()
	var historyLogs []*HistoryLog
	for resultsIterator.HasNext() {
		queryResponse, err := resultsIterator.Next()
		if err != nil {
			logger(true, "GetHistoryLogsBasedOnCallerName", err.Error())
			return nil, err
		}
		historyLog := new(HistoryLog)
		err = json.Unmarshal(queryResponse.Value, historyLog)
		if err != nil {
			logger(true, "GetHistoryLogsBasedOnCallerName", err.Error())
			return nil, err
		}
		historyLogs = append(historyLogs, historyLog)
	}
	logger(false, "GetHistoryLogsBasedOnCallerName", "METHOD END")
	return historyLogs, nil
}

func (s *SmartContract) GetHistoryBasedOnMethodName(ctx contractapi.TransactionContextInterface, callerType string, callerIP string, callerAddress string, methodName string) ([]*HistoryLog, error) {
	logger(false, "GetHistoryBasedOnMethodName", "METHOD START...")
	resultsIterator, err := ctx.GetStub().GetStateByPartialCompositeKey(IOP_HISTORY_PREFIX, []string{callerType, callerIP, callerAddress, methodName})
	if err != nil {
		logger(true, "GetHistoryBasedOnMethodName", err.Error())
		return nil, err
	}
	defer resultsIterator.Close()
	var historyLogs []*HistoryLog
	for resultsIterator.HasNext() {
		queryResponse, err := resultsIterator.Next()
		if err != nil {
			logger(true, "GetHistoryBasedOnMethodName", err.Error())
			return nil, err
		}
		historyLog := new(HistoryLog)
		err = json.Unmarshal(queryResponse.Value, historyLog)
		if err != nil {
			logger(true, "GetHistoryBasedOnMethodName", err.Error())
			return nil, err
		}
		historyLogs = append(historyLogs, historyLog)
	}
	logger(false, "GetHistoryBasedOnMethodName", "METHOD END")
	return historyLogs, nil
}

func (s *SmartContract) GetHistoryBasedOnDate(ctx contractapi.TransactionContextInterface, callerType string, callerIP string, callerAddress string, methodName string, year string, month string, day string) ([]*HistoryLog, error) {

	logger(false, "GetHistoryBasedOnDate", "METHOD START...")
	resultsIterator, err := ctx.GetStub().GetStateByPartialCompositeKey(IOP_HISTORY_PREFIX, []string{callerType, callerIP, callerAddress, methodName, year, month, day})
	if err != nil {
		logger(true, "GetHistoryBasedOnDate", err.Error())
		return nil, err
	}
	defer resultsIterator.Close()
	var historyLogs []*HistoryLog
	for resultsIterator.HasNext() {
		queryResponse, err := resultsIterator.Next()
		if err != nil {
			logger(true, "GetHistoryBasedOnDate", err.Error())
			return nil, err
		}
		historyLog := new(HistoryLog)
		err = json.Unmarshal(queryResponse.Value, historyLog)
		if err != nil {
			logger(true, "GetHistoryBasedOnDate", err.Error())
			return nil, err
		}
		historyLogs = append(historyLogs, historyLog)
	}
	logger(false, "GetHistoryBasedOnDate", "METHOD END")
	return historyLogs, nil
}
func (s *SmartContract) GetHistoryBasedOnYear(ctx contractapi.TransactionContextInterface, callerType string, callerIP string, callerAddress string, methodName string, year string) ([]*HistoryLog, error) {
	logger(false, "GetHistoryBasedOnYear", "METHOD START...")
	resultsIterator, err := ctx.GetStub().GetStateByPartialCompositeKey(IOP_HISTORY_PREFIX, []string{callerType, callerIP, callerAddress, methodName, year})
	if err != nil {
		logger(true, "GetHistoryBasedOnYear", err.Error())
		return nil, err
	}
	defer resultsIterator.Close()
	var historyLogs []*HistoryLog
	for resultsIterator.HasNext() {
		queryResponse, err := resultsIterator.Next()
		if err != nil {
			logger(true, "GetHistoryBasedOnYear", err.Error())
			return nil, err
		}
		historyLog := new(HistoryLog)
		err = json.Unmarshal(queryResponse.Value, historyLog)
		if err != nil {
			logger(true, "GetHistoryBasedOnYear", err.Error())
			return nil, err
		}
		historyLogs = append(historyLogs, historyLog)
	}
	logger(false, "GetHistoryBasedOnYear", "METHOD END")
	return historyLogs, nil
}

func (s *SmartContract) GetHistoryBasedOnMonth(ctx contractapi.TransactionContextInterface, callerType string, callerIP string, callerAddress string, methodName string, year string, month string) ([]*HistoryLog, error) {
	logger(false, "GetHistoryBasedOnMonth", "METHOD START...")
	resultsIterator, err := ctx.GetStub().GetStateByPartialCompositeKey(IOP_HISTORY_PREFIX, []string{callerType, callerIP, callerAddress, methodName, year, month})
	if err != nil {
		logger(true, "GetHistoryBasedOnMonth", err.Error())
		return nil, err
	}
	defer resultsIterator.Close()
	var historyLogs []*HistoryLog
	for resultsIterator.HasNext() {
		queryResponse, err := resultsIterator.Next()
		if err != nil {
			logger(true, "GetHistoryBasedOnMonth", err.Error())
			return nil, err
		}
		historyLog := new(HistoryLog)
		err = json.Unmarshal(queryResponse.Value, historyLog)
		if err != nil {
			logger(true, "GetHistoryBasedOnMonth", err.Error())
			return nil, err
		}
		historyLogs = append(historyLogs, historyLog)
	}
	logger(false, "GetHistoryBasedOnMonth", "METHOD END")
	return historyLogs, nil
}

func (s *SmartContract) GetHistoryBasedOnHour(ctx contractapi.TransactionContextInterface, callerType string, callerIP string, callerAddress string, methodName string, year string, month string, day string, hour string) ([]*HistoryLog, error) {
	logger(false, "GetHistoryByHour", "METHOD START...")
	resultsIterator, err := ctx.GetStub().GetStateByPartialCompositeKey(IOP_HISTORY_PREFIX, []string{callerType, callerIP, callerAddress, methodName, year, month, day, hour})
	if err != nil {
		logger(true, "GetHistoryByHour", err.Error())
		return nil, err
	}
	defer resultsIterator.Close()
	var historyLogs []*HistoryLog
	for resultsIterator.HasNext() {
		queryResponse, err := resultsIterator.Next()
		if err != nil {
			logger(true, "GetHistoryByHour", err.Error())
			return nil, err
		}
		historyLog := new(HistoryLog)
		err = json.Unmarshal(queryResponse.Value, historyLog)
		if err != nil {
			logger(true, "GetHistoryByHour", err.Error())
			return nil, err
		}
		historyLogs = append(historyLogs, historyLog)
	}
	logger(false, "GetHistoryByHour", "METHOD END")
	return historyLogs, nil
}

func (s *SmartContract) RemoveAHistoryLog(ctx contractapi.TransactionContextInterface, callerType string, callerIP string, callerAddress string, methodName string, year string, month string, day string, hour string, distinctivePart string) error {
	logger(false, "RemoveAHistoryLog", "METHOD START...")
	id := s.createKey(ctx, callerType, callerIP, callerAddress, methodName, year, month, day, hour, distinctivePart)
	err := ctx.GetStub().DelState(id)
	if err != nil {
		logger(true, "RemoveAHistoryLog", err.Error())
		return err
	}
	logger(false, "RemoveAHistoryLog", "METHOD END")
	return nil
}

func logger(isError bool, methodName string, txt string) {
	if isError {
		fmt.Println("***Error at  METHOD_NAME:", methodName, " TEXT: ", txt)
		return
	}
	fmt.Println("METHOD_NAME: ", methodName, " TEXT: ", txt)
}
