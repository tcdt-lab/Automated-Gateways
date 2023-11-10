package configs

import (
	"os"
)
import "gopkg.in/yaml.v2"

const (
	HYPERLEDGER_FABRIC_NETWROK_TYPE = "HyperledgerFabricNetwork"
)

var YamlAddress string

type Configuration struct {
	Ca       CA       `yaml:"ca"`
	Server   Server   `yaml:"server"`
	Client   []Client `yaml:"client"`
	Platform Platform `yaml:"platform"`
}

type CA struct {
	CertPath string `yaml:"cert_path"`
	KeyPath  string `yaml:"key_path"`
}
type Server struct {
	Port     string `yaml:"port"`
	CertPath string `yaml:"cert_path"`
	KeyPath  string `yaml:"key_path"`
}

type Client struct {
	AccessibleNetworkId   string `yaml:"accessible_network_id"`
	AccessibleNetworkName string `yaml:"accessible_network_name"`
	CertPath              string `yaml:"cert_path"`
	KeyPath               string `yaml:"key_path"`
	CaCertPath            string `yaml:"ca_cert_path"`
	Ip                    string `yaml:"ip"`
	Address               string `yaml:"address"`
	CompanyName           string `yaml:"company_name"`
	Port                  string `yaml:"port"`
}
type Platform struct {
	Hyperledger Hyperledger `yaml:"hyperledger"`
}

type Hyperledger struct {
	Name                           string `yaml:"name"`
	MspID                          string `yaml:"msp_id"`
	CertPath                       string `yaml:"cert_path"`
	KeyPath                        string `yaml:"key_path"`
	TlsCertPath                    string `yaml:"tls_cert_path"`
	PeerEndpoint                   string `yaml:"peer_endpoint"`
	GatewayPeer                    string `yaml:"gateway_peer"`
	AccessibleNetworkChannelName   string `yaml:"accessible_network_channel_name"`
	AccessibleNetworkChaincodeName string `yaml:"accessible_network_chaincode_name"`
	PermittedNetworkChannelName    string `yaml:"permitted_network_channel_name"`
	PermittedNetworkChaincodeName  string `yaml:"permitted_network_chaincode_name"`
	PermittedMethodChaincodeName   string `yaml:"permitted_method_chaincode_name"`
	PermittedMethodChannelName     string `yaml:"permitted_method_channel_name"`
}

func ReadConfigYAMLFile() (Configuration, error) {
	// Open our yaml File

	var configuration Configuration
	yamlFile, err := os.ReadFile(YamlAddress)
	if err != nil {

		return configuration, err
	}

	err = yaml.Unmarshal(yamlFile, &configuration)
	if err != nil {
		return configuration, err
	}
	return configuration, nil
}

func GetClientConfigByClientAccessibleId(clientAccessibleId string, configuration Configuration) Client {
	for _, client := range configuration.Client {
		if client.AccessibleNetworkId == clientAccessibleId {
			return client
		}
	}
	return Client{}
}
