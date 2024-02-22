package starter

import (
	"github.com/tcdt-lab/Automated-Gateways/relay/internal/offchain/server"
	cli "github.com/tcdt-lab/Automated-Gateways/relay/iop_cli"
)

func StartIopServer() {
	server.StartServer()

}

func StartIopCLI() {
	cli.StartCli()
}
