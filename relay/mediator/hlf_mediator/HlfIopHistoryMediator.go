package hlf_mediator

import (
	datatypes "github.com/tcdt-lab/Automated-Gateways/relay/data_types"
)
import iopHistoryOnchain "github.com/tcdt-lab/Automated-Gateways/relay/internal/onchain/iop_history_handler"

type HlfIopHistoryMediator struct {
}

var hlfIopHistory = iopHistoryOnchain.HlfIopHistory{}

func (iopHistory *HlfIopHistoryMediator) InsertHistoryLog(callType string, callerIP string, callerAddress string, methodName string, inputArgs string, outputArgs string) (*datatypes.HistoryLog, error) {
	conn, gw, err := hlfIopHistory.OpenConnection()
	if err != nil {
		return nil, err
	}
	defer hlfIopHistory.CloseConnection(conn, gw)
	return hlfIopHistory.InsertHistoryLog(gw, callType, callerIP, callerAddress, methodName, inputArgs, outputArgs)
}

func (iopHistory *HlfIopHistoryMediator) GetHistoryBasedOnCallType(callType string) ([]*datatypes.HistoryLog, error) {
	conn, gw, err := hlfIopHistory.OpenConnection()
	if err != nil {
		return nil, err
	}
	defer hlfIopHistory.CloseConnection(conn, gw)
	return hlfIopHistory.GetHistoryBasedOnCallType(gw, callType)
}

func (iopHistory *HlfIopHistoryMediator) GetHistoryBasedOnCallerIP(callType string, callerIP string) ([]*datatypes.HistoryLog, error) {
	conn, gw, err := hlfIopHistory.OpenConnection()
	if err != nil {
		return nil, err
	}
	defer hlfIopHistory.CloseConnection(conn, gw)
	return hlfIopHistory.GetHistoryBasedOnCallerIP(gw, callType, callerIP)
}

func (iopHistory *HlfIopHistoryMediator) GetHistoryBasedOnCallerAddress(callType string, callerIP string, callerAddress string) ([]*datatypes.HistoryLog, error) {
	conn, gw, err := hlfIopHistory.OpenConnection()
	if err != nil {
		return nil, err
	}
	defer hlfIopHistory.CloseConnection(conn, gw)
	return hlfIopHistory.GetHistoryBasedOnCallerAddress(gw, callType, callerIP, callerAddress)
}

func (iopHistory *HlfIopHistoryMediator) GetHistoryBasedOnMethodName(callType string, callerIP string, callerAddress string, methodName string) ([]*datatypes.HistoryLog, error) {
	conn, gw, err := hlfIopHistory.OpenConnection()
	if err != nil {
		return nil, err
	}
	defer hlfIopHistory.CloseConnection(conn, gw)
	return hlfIopHistory.GetHistoryBasedOnMethodName(gw, callType, callerIP, callerAddress, methodName)
}

func (iopHistory *HlfIopHistoryMediator) GetHistoryBasedOnDate(callerType string, callerIP string, callerAddress string, methodName string, year string, month string, day string) ([]*datatypes.HistoryLog, error) {
	conn, gw, err := hlfIopHistory.OpenConnection()
	if err != nil {
		return nil, err
	}
	defer hlfIopHistory.CloseConnection(conn, gw)
	return hlfIopHistory.GetHistoryBasedOnDate(gw, callerType, callerIP, callerAddress, methodName, year, month, day)
}

func (iopHistory *HlfIopHistoryMediator) GetHistoryBasedOnYear(callerType string, callerIP string, callerAddress string, methodName string, year string) ([]datatypes.HistoryLog, error) {
	conn, gw, err := hlfIopHistory.OpenConnection()
	if err != nil {
		return nil, err
	}
	defer hlfIopHistory.CloseConnection(conn, gw)
	return hlfIopHistory.GetHistoryBasedOnYear(gw, callerType, callerIP, callerAddress, methodName, year)
}

func (iopHistory *HlfIopHistoryMediator) GetHistoryBasedOnMonth(callerType string, callerIP string, callerAddress string, methodName string, year string, month string) ([]datatypes.HistoryLog, error) {
	conn, gw, err := hlfIopHistory.OpenConnection()
	if err != nil {
		return nil, err
	}
	defer hlfIopHistory.CloseConnection(conn, gw)
	return hlfIopHistory.GetHistoryBasedOnMonth(gw, callerType, callerIP, callerAddress, methodName, year, month)
}

func (iopHistory *HlfIopHistoryMediator) GetHistoryBasedOnHour(callerType string, callerIP string, callerAddress string, methodName string, year string, month string, day string, hour string) ([]datatypes.HistoryLog, error) {
	conn, gw, err := hlfIopHistory.OpenConnection()
	if err != nil {
		return nil, err
	}
	defer hlfIopHistory.CloseConnection(conn, gw)
	return hlfIopHistory.GetHistoryBasedOnHour(gw, callerType, callerIP, callerAddress, methodName, year, month, day, hour)
}

func (iopHistory *HlfIopHistoryMediator) RemoveHistoryLog(callerType string, callerIP string, callerAddress string, methodName string, year string, month string, day string, hour string, distinctive string) error {
	conn, gw, err := hlfIopHistory.OpenConnection()
	if err != nil {
		return err
	}
	defer hlfIopHistory.CloseConnection(conn, gw)
	return hlfIopHistory.RemoveHistoryLog(gw, callerType, callerIP, callerAddress, methodName, year, month, day, hour, distinctive)
}
