package main

import (
	"github.com/tcdt-lab/Automated-Gateways/relay/configs"
)

func main() {
	configs.YamlAddress = "/home/koosha/Desktop/Automated-Gateways/relay/cmd/config.yaml"
	startIop()
}
