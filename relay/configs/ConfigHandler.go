package configs

import "os"

const (
	HYPERLEDGER_FABRIC_NETWROK_TYPE = "HyperledgerFabricNetwork"
)

type Configuration struct {
	CA_CERT_ADDRESS     string `json:"ca_cert_address"`
	CA_KEY_ADDRESS      string `json:"ca_key_address"`
	SERVER_CERT_ADDRESS string `json:"server_cert_address"`
	SERVER_KEY_ADDRESS  string `json:"server_key_address"`
	SERVER_PORT         string `json:"server_port"`
	CLIENT_CERT_ADDRESS string `json:"client_cert_address"`
	CLIENT_KEY_ADDRESS  string `json:"client_key_address"`

	//******* SUPPORTED NETWORKS ********
	hyperledegerFabric HyperledgerFabricChain `json:"hyperledger_fabric"`
}

type HyperledgerFabricChain struct {
	ChainName          string `json:"chain_name"`
	CHAIN_CERT_ADDRESS string `json:"chain_cert_address"`
	CHAIN_KEY_ADDRESS  string `json:"chain_key_address"`
	CHAIN_CA_ADDRESS   string `json:"chain_ca_address"`
	MSP_ID             string `json:"msp_id"`
	PEER_ENDPOINT      string `json:"peer_endpoint"`
	GATEWAY_PEER       string `json:"gateway_peer"`
}

func (Configuration *Configuration) readConfigYAMLFile(yamlLocation string) error {
	// Open our yaml File
	yamlFile, err := os.Open(yamlLocation)
	if err != nil {
		return err
	}
	defer func(yamlFile *os.File) {
		err := yamlFile.Close()
		if err != nil {

		}
	}(yamlFile)

	return nil
}
