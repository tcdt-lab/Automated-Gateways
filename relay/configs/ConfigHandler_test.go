package configs

import (
	"fmt"
	"testing"
)

func TestCreateAccessibleNetwork(t *testing.T) {
	YamlAddress = "/home/koosha/Desktop/Automated-Gateways/relay/cmd/config.yaml"
	config, err := ReadConfigYAMLFile()
	if err != nil {
		t.Errorf("Creating Accessible Netwrok: %v", err)
	}
	fmt.Println(config)
}
