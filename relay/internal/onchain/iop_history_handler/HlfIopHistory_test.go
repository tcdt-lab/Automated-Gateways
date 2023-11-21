package iop_history_handler

import (
	"github.com/tcdt-lab/Automated-Gateways/relay/configs"
	"github.com/tcdt-lab/Automated-Gateways/relay/data_types"
	"testing"
)

func TestOpenCloseConnetion(t *testing.T) {
	configs.YamlAddress = "/home/koosha/Desktop/Automated-Gateways/relay/cmd/config.yaml"
	var hlfIopHistory HlfIopHistory
	clientConn, gateway, err := hlfIopHistory.OpenConnection()
	if err != nil {
		t.Errorf("Error opening connection: %v", err)
	}

	errConnetion := hlfIopHistory.CloseConnection(clientConn, gateway)
	if errConnetion != nil {
		t.Errorf("Error closing connection: %v", errConnetion)
	}

}

func TestInsertIopHistory(t *testing.T) {
	configs.YamlAddress = "/home/koosha/Desktop/Automated-Gateways/relay/cmd/config.yaml"
	var hlfIopHistory HlfIopHistory
	clientConn, gateway, err := hlfIopHistory.OpenConnection()
	if err != nil {
		t.Errorf("Error opening connection: %v", err)
	}
	defer hlfIopHistory.CloseConnection(clientConn, gateway)

	submitRes, err := hlfIopHistory.InsertHistoryLog(gateway, data_types.HISTORY_TYPE_INSIDE_INITIATED, "127.0.0.1", "localhost", "sssss", "	{'23','99'}", "{'111'}")
	if err != nil {
		t.Errorf("Get HistoryNetwrok: %v", err)
	}
	println(submitRes.Id)
}

func TestGetHistoryBasedOnCallType(t *testing.T) {
	var hlfIopHistory HlfIopHistory
	configs.YamlAddress = "/home/koosha/Desktop/Automated-Gateways/relay/cmd/config.yaml"
	clientConn, gateway, err := hlfIopHistory.OpenConnection()
	if err != nil {
		t.Errorf("Error opening connection: %v", err)
	}
	defer hlfIopHistory.CloseConnection(clientConn, gateway)

	res, err := hlfIopHistory.GetHistoryBasedOnCallType(gateway, data_types.HISTORY_TYPE_INSIDE_INITIATED)
	if err != nil {
		t.Errorf(" Get History Exists : %v", err)
	}
	if res == nil {
		t.Errorf(" Get History Exists : %v", res)
	}

	for _, log := range res {
		println(log.Id)
		println(log.CallType)
		println(log.InputArgs)
		println(log.OutputArgs)
		println(log.CallerIP)
		println(log.CallerAddress)
		println(log.MethodName)

	}
}

func TestGetHistoryBasedOnCallerIP(t *testing.T) {
	var hlfIopHistory HlfIopHistory
	configs.YamlAddress = "/home/koosha/Desktop/Automated-Gateways/relay/cmd/config.yaml"
	clientConn, gateway, err := hlfIopHistory.OpenConnection()
	if err != nil {
		t.Errorf("Error opening connection: %v", err)
	}
	defer hlfIopHistory.CloseConnection(clientConn, gateway)

	res, err := hlfIopHistory.GetHistoryBasedOnCallerIP(gateway, data_types.HISTORY_TYPE_INSIDE_INITIATED, "127.0.0.1")
	if err != nil {
		t.Errorf(" Get History : %v", err)
	}
	if res == nil {
		t.Errorf(" Get History Exists : %v", res)
	}
	for _, log := range res {
		println(log.Id)
		println(log.CallType)
		println(log.InputArgs)
		println(log.OutputArgs)
		println(log.CallerIP)
		println(log.CallerAddress)
	}
}

func TestGetHistoryBasedOnCallerAddress(t *testing.T) {
	var hlfIopHistory HlfIopHistory
	configs.YamlAddress = "/home/koosha/Desktop/Automated-Gateways/relay/cmd/config.yaml"
	clientConn, gateway, err := hlfIopHistory.OpenConnection()
	if err != nil {
		t.Errorf("Error opening connection: %v", err)
	}
	defer hlfIopHistory.CloseConnection(clientConn, gateway)

	res, err := hlfIopHistory.GetHistoryBasedOnCallerAddress(gateway, data_types.HISTORY_TYPE_INSIDE_INITIATED, "127.0.0.1", "localhost")
	if err != nil {
		t.Errorf(" AGet History : %v", err)
	}
	if res == nil {
		t.Errorf(" Get History Exists : %v", res)
	}
	for _, log := range res {
		println(log.Id)
		println(log.CallType)
		println(log.InputArgs)
		println(log.OutputArgs)
		println(log.CallerIP)
		println(log.CallerAddress)
	}
}

func TestGetHistoryBasedOnMethodName(t *testing.T) {
	var hlfIopHistory HlfIopHistory
	configs.YamlAddress = "/home/koosha/Desktop/Automated-Gateways/relay/cmd/config.yaml"
	clientConn, gateway, err := hlfIopHistory.OpenConnection()
	if err != nil {
		t.Errorf("Error opening connection: %v", err)
	}
	defer hlfIopHistory.CloseConnection(clientConn, gateway)

	res, err := hlfIopHistory.GetHistoryBasedOnMethodName(gateway, data_types.HISTORY_TYPE_INSIDE_INITIATED, "127.0.0.1", "localhost", "sssss")
	if err != nil {
		t.Errorf(" Get History : %v", err)
	}
	if res == nil {
		t.Errorf(" Get History Exists : %v", res)
	}
	for _, log := range res {
		println(log.Id)
		println(log.CallType)
		println(log.InputArgs)
		println(log.OutputArgs)
		println(log.CallerIP)
		println(log.CallerAddress)
	}
}

func TestGetHistoryBasedOnDate(t *testing.T) {
	var hlfIopHistory HlfIopHistory
	configs.YamlAddress = "/home/koosha/Desktop/Automated-Gateways/relay/cmd/config.yaml"
	clientConn, gateway, err := hlfIopHistory.OpenConnection()
	if err != nil {
		t.Errorf("Error opening connection: %v", err)
	}
	defer hlfIopHistory.CloseConnection(clientConn, gateway)

	res, err := hlfIopHistory.GetHistoryBasedOnDate(gateway, data_types.HISTORY_TYPE_INSIDE_INITIATED, "127.0.0.1", "localhost", "sssss", "2023", "11", "19")

	if err != nil {
		t.Errorf(" Get History : %v", err)
	}
	if res == nil {
		t.Errorf(" Get History Exists : %v", res)
	}
	for _, log := range res {
		println(log.Id)
		println(log.CallType)
		println(log.InputArgs)
		println(log.OutputArgs)
		println(log.CallerIP)
		println(log.CallerAddress)
	}
}

func TestGetHistoryBasedOnYear(t *testing.T) {
	var hlfIopHistory HlfIopHistory
	configs.YamlAddress = "/home/koosha/Desktop/Automated-Gateways/relay/cmd/config.yaml"
	clientConn, gateway, err := hlfIopHistory.OpenConnection()
	if err != nil {
		t.Errorf("Error opening connection: %v", err)
	}
	defer hlfIopHistory.CloseConnection(clientConn, gateway)

	res, err := hlfIopHistory.GetHistoryBasedOnYear(gateway, data_types.HISTORY_TYPE_INSIDE_INITIATED, "127.0.0.1", "localhost", "sssss", "2023")

	if err != nil {
		t.Errorf(" Get History : %v", err)
	}
	if res == nil {
		t.Errorf(" Get History Exists : %v", res)
	}
	for _, log := range res {
		println(log.Id)
		println(log.CallType)
		println(log.InputArgs)
		println(log.OutputArgs)
		println(log.CallerIP)
		println(log.CallerAddress)
	}
}

func TestGetHistoryBasedOnMoth(t *testing.T) {
	var hlfIopHistory HlfIopHistory
	configs.YamlAddress = "/home/koosha/Desktop/Automated-Gateways/relay/cmd/config.yaml"
	clientConn, gateway, err := hlfIopHistory.OpenConnection()
	if err != nil {
		t.Errorf("Error opening connection: %v", err)
	}
	defer hlfIopHistory.CloseConnection(clientConn, gateway)

	res, err := hlfIopHistory.GetHistoryBasedOnMonth(gateway, data_types.HISTORY_TYPE_INSIDE_INITIATED, "127.0.0.1", "localhost", "sssss", "2023", "11")

	if err != nil {
		t.Errorf(" Get History : %v", err)
	}
	if res == nil {
		t.Errorf(" Get History Not Exists : %v", res)
	}
	for _, log := range res {
		println(log.Id)
		println(log.CallType)
		println(log.InputArgs)
		println(log.OutputArgs)
		println(log.CallerIP)
		println(log.CallerAddress)
	}
}

func TestGetHistoryBasedOnHour(t *testing.T) {
	var hlfIopHistory HlfIopHistory
	configs.YamlAddress = "/home/koosha/Desktop/Automated-Gateways/relay/cmd/config.yaml"
	clientConn, gateway, err := hlfIopHistory.OpenConnection()
	if err != nil {
		t.Errorf("Error opening connection: %v", err)
	}
	defer hlfIopHistory.CloseConnection(clientConn, gateway)

	res, err := hlfIopHistory.GetHistoryBasedOnHour(gateway, data_types.HISTORY_TYPE_INSIDE_INITIATED, "127.0.0.1", "localhost", "sssss", "2023", "11", "19", "1")

	if err != nil {
		t.Errorf(" Get History : %v", err)
	}
	if res == nil {
		t.Errorf(" Get History Not Exists : %v", res)
	}
	for _, log := range res {
		println(log.Id)
		println(log.CallType)
		println(log.InputArgs)
		println(log.OutputArgs)
		println(log.CallerIP)
		println(log.CallerAddress)
	}
}

func TestRemoveHistoryLog(t *testing.T) {
	var hlfIopHistory HlfIopHistory
	configs.YamlAddress = "/home/koosha/Desktop/Automated-Gateways/relay/cmd/config.yaml"
	clientConn, gateway, err := hlfIopHistory.OpenConnection()
	if err != nil {
		t.Errorf("Error opening connection: %v", err)
	}
	defer hlfIopHistory.CloseConnection(clientConn, gateway)

	err = hlfIopHistory.RemoveHistoryLog(gateway, data_types.HISTORY_TYPE_INSIDE_INITIATED, "127.0.0.1", "localhost", "KooshaMethod", "2023", "11", "19", "1", "998853374")
	if err != nil {
		t.Errorf(" Get History : %v", err)
	}

}
