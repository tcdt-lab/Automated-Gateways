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
	AccessibleNetworkId string `yaml:"accessible_network_id"`
	CertPath            string `yaml:"cert_path"`
	KeyPath             string `yaml:"key_path"`
	CaCertPath          string `yaml:"ca_cert_path"`
}
type Platform struct {
	Hyperledger Hyperledger `yaml:"hyperledger"`
}
type Hyperledger struct {
	Name         string `yaml:"name"`
	MspID        string `yaml:"msp_id"`
	CertPath     string `yaml:"cert_path"`
	KeyPath      string `yaml:"key_path"`
	TlsCertPath  string `yaml:"tls_cert_path"`
	PeerEndpoint string `yaml:"peer_endpoint"`
	GatewayPeer  string `yaml:"gateway_peer"`
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
