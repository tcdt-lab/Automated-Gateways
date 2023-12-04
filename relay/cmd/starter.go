package main

import (
	"github.com/tcdt-lab/Automated-Gateways/relay/internal/offchain/server"
	cli "github.com/tcdt-lab/Automated-Gateways/relay/iop_cli"
)

func startIop() {
	go server.StartServer()
	cli.StartCli()
}
