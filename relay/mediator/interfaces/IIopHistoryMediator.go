package interfaces

import (
	datatypes "github.com/tcdt-lab/Automated-Gateways/relay/data_types"
)

type IIopHistoryMediator interface {
	InsertHistoryLog(callType string, callerIP string, callerAddress string, methodName string, inputArgs string, outputArgs string) (*datatypes.HistoryLog, error)
	GetHistoryBasedOnCallType(callType string) ([]*datatypes.HistoryLog, error)
	GetHistoryBasedOnCallerIP(callType string, callerIP string) ([]*datatypes.HistoryLog, error)
	GetHistoryBasedOnCallerAddress(callType string, callerIP string, callerAddress string) ([]*datatypes.HistoryLog, error)
	GetHistoryBasedOnMethodName(callType string, callerIP string, callerAddress string, methodName string) ([]*datatypes.HistoryLog, error)
	GetHistoryBasedOnDate(callerType string, callerIP string, callerAddress string, methodName string, year string, month string, day string) ([]*datatypes.HistoryLog, error)
	GetHistoryBasedOnYear(callerType string, callerIP string, callerAddress string, methodName string, year string) ([]datatypes.HistoryLog, error)
	GetHistoryBasedOnMonth(callerType string, callerIP string, callerAddress string, methodName string, year string, month string) ([]datatypes.HistoryLog, error)
	GetHistoryBasedOnHour(callerType string, callerIP string, callerAddress string, methodName string, year string, month string, day string, hour string) ([]datatypes.HistoryLog, error)
	RemoveHistoryLog(callerType string, callerIP string, callerAddress string, methodName string, year string, month string, day string, hour string, distinctive string) error
}
