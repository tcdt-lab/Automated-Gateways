package data_types

type PermittedMethodInfo struct {
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
