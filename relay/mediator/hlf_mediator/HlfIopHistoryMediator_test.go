package hlf_mediator

import (
	"github.com/tcdt-lab/Automated-Gateways/relay/configs"
	"github.com/tcdt-lab/Automated-Gateways/relay/data_types"
	"testing"
)

func TestHlfIopHistoryMediator_InsertHistoryLog(t *testing.T) {
	configs.YamlAddress = "/home/koosha/Desktop/Automated-Gateways/relay/cmd/config.yaml"
	var hlfHistoryMediator HlfIopHistoryMediator
	res, err := hlfHistoryMediator.InsertHistoryLog(data_types.HISTORY_TYPE_INSIDE_INITIATED, "127.0.0.1", "localhost", "sssss", "	{'23','99'}", "{'111'}")

	if err != nil {
		t.Errorf("Error opening connection: %v", err)
	}
	if res == nil {
		t.Errorf(" Get History Exists : %v", res)
	}

	println(res.Id)
	println(res.CallType)
	println(res.CallerIP)
	println(res.CallerAddress)
	println(res.Distinctive)

}

func TestGetHistoryBasedOnCallType(t *testing.T) {
	var hlfHistoryMediator HlfIopHistoryMediator
	configs.YamlAddress = "/home/koosha/Desktop/Automated-Gateways/relay/cmd/config.yaml"
	clientConn, gateway, err := hlfIopHistory.OpenConnection()
	if err != nil {
		t.Errorf("Error opening connection: %v", err)
	}
	defer hlfIopHistory.CloseConnection(clientConn, gateway)

	res, err := hlfHistoryMediator.GetHistoryBasedOnCallType(data_types.HISTORY_TYPE_INSIDE_INITIATED)
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
	var hlfHistoryMediator HlfIopHistoryMediator
	configs.YamlAddress = "/home/koosha/Desktop/Automated-Gateways/relay/cmd/config.yaml"
	clientConn, gateway, err := hlfIopHistory.OpenConnection()
	if err != nil {
		t.Errorf("Error opening connection: %v", err)
	}
	defer hlfIopHistory.CloseConnection(clientConn, gateway)

	res, err := hlfHistoryMediator.GetHistoryBasedOnCallerIP(data_types.HISTORY_TYPE_INSIDE_INITIATED, "127.0.0.1")
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
	var hlfHistoryMediator HlfIopHistoryMediator
	configs.YamlAddress = "/home/koosha/Desktop/Automated-Gateways/relay/cmd/config.yaml"
	clientConn, gateway, err := hlfIopHistory.OpenConnection()
	if err != nil {
		t.Errorf("Error opening connection: %v", err)
	}
	defer hlfIopHistory.CloseConnection(clientConn, gateway)
	res, err := hlfHistoryMediator.GetHistoryBasedOnCallerAddress(data_types.HISTORY_TYPE_INSIDE_INITIATED, "127.0.0.1", "localhost")
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

	var hlfHistoryMediator HlfIopHistoryMediator
	configs.YamlAddress = "/home/koosha/Desktop/Automated-Gateways/relay/cmd/config.yaml"
	clientConn, gateway, err := hlfIopHistory.OpenConnection()
	if err != nil {
		t.Errorf("Error opening connection: %v", err)
	}
	defer hlfIopHistory.CloseConnection(clientConn, gateway)
	res, err := hlfHistoryMediator.GetHistoryBasedOnMethodName(data_types.HISTORY_TYPE_INSIDE_INITIATED, "127.0.0.1", "localhost", "sssss")
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
	var hlfHistoryMediator HlfIopHistoryMediator
	configs.YamlAddress = "/home/koosha/Desktop/Automated-Gateways/relay/cmd/config.yaml"
	clientConn, gateway, err := hlfIopHistory.OpenConnection()
	if err != nil {
		t.Errorf("Error opening connection: %v", err)
	}
	defer hlfIopHistory.CloseConnection(clientConn, gateway)
	res, err := hlfHistoryMediator.GetHistoryBasedOnDate(data_types.HISTORY_TYPE_INSIDE_INITIATED, "127.0.0.1", "localhost", "sssss", "2023", "11", "19")
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
	var hlfHistoryMediator HlfIopHistoryMediator
	configs.YamlAddress = "/home/koosha/Desktop/Automated-Gateways/relay/cmd/config.yaml"
	clientConn, gateway, err := hlfIopHistory.OpenConnection()
	if err != nil {
		t.Errorf("Error opening connection: %v", err)
	}
	defer hlfIopHistory.CloseConnection(clientConn, gateway)
	res, err := hlfHistoryMediator.GetHistoryBasedOnYear(data_types.HISTORY_TYPE_INSIDE_INITIATED, "127.0.0.1", "localhost", "sssss", "2023")
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

func TestGetHistoryBasedOnMonth(t *testing.T) {
	var hlfHistoryMediator HlfIopHistoryMediator
	configs.YamlAddress = "/home/koosha/Desktop/Automated-Gateways/relay/cmd/config.yaml"
	clientConn, gateway, err := hlfIopHistory.OpenConnection()
	if err != nil {
		t.Errorf("Error opening connection: %v", err)
	}
	defer hlfIopHistory.CloseConnection(clientConn, gateway)
	res, err := hlfHistoryMediator.GetHistoryBasedOnMonth(data_types.HISTORY_TYPE_INSIDE_INITIATED, "127.0.0.1", "localhost", "sssss", "2023", "11")
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

func TestGetHistoryBasedOnHour(t *testing.T) {

	var hlfHistoryMediator HlfIopHistoryMediator
	configs.YamlAddress = "/home/koosha/Desktop/Automated-Gateways/relay/cmd/config.yaml"

	res, err := hlfHistoryMediator.GetHistoryBasedOnHour(data_types.HISTORY_TYPE_INSIDE_INITIATED, "127.0.0.1", "localhost", "sssss", "2023", "11", "19", "1")

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

func TestRemoveHistoryLog(t *testing.T) {

	var hlfHistoryMediator HlfIopHistoryMediator
	configs.YamlAddress = "/home/koosha/Desktop/Automated-Gateways/relay/cmd/config.yaml"

	err := hlfHistoryMediator.RemoveHistoryLog(data_types.HISTORY_TYPE_INSIDE_INITIATED, "127.0.0.1", "localhost", "sssss", "2023", "11", "19", "1", "998853374")

	if err != nil {
		t.Errorf(" Get History : %v", err)
	}

}
