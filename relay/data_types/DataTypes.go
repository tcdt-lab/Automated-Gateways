package data_types

type MethodInfo struct {
	PermittedMethodId string `json:"PermittedMethodId"`
	Name              string `json:"Name"`
	Chaincode         string `json:"Chaincode"`
	Channel           string `json:"Channel"`
	InputArgs         string `json:"inputArgs"`
	OutputArgs        string `json:"outputArgs"`
}

type PermittedNetworkInfo struct {
	NetworkName        string `json:"NetworkName"`
	IP                 string `json:"IP"`
	ADDRESS            string `json:"ADDRESS"`
	CompanyName        string `json:"CompanyName"`
	PermittedNetworkId string `json:"PermittedNetworkId"`
}

type AccessibleNetworkInfo struct {
	NetworkName         string `json:"NetworkName"`
	IP                  string `json:"IP"`
	ADDRESS             string `json:"ADDRESS"`
	CompanyName         string `json:"CompanyName"`
	AccessibleNetworkId string `json:"AccessibleNetworkId"`
}

const HISTORY_TYPE_INSIDE_INITIATED = "Inside_initiated_call"
const HISTORY_TYPE_OUTSIDE_INITIATED = "Outside_initiated_call"

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
