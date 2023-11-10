package main

import (
	"github.com/tcdt-lab/Automated-Gateways/relay/configs"
	server "github.com/tcdt-lab/Automated-Gateways/relay/internal/offchain/server"
	cli "github.com/tcdt-lab/Automated-Gateways/relay/iop_cli"
)

func main() {
	configs.YamlAddress = "/home/koosha/Desktop/Automated-Gateways/relay/cmd/config.yaml"
	go server.StartServer()
	cli.StartCli()
}
