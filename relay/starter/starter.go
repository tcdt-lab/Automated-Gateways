package starter

import (
	"github.com/tcdt-lab/Automated-Gateways/relay/internal/offchain/server"
	cli "github.com/tcdt-lab/Automated-Gateways/relay/iop_cli"
	"github.com/tcdt-lab/Automated-Gateways/relay/logger"
)

func init() {
	logger.Init("app.log")
}
func StartIopServer() {
	logger.Info("Starting Iop Server", "starter", "StartIopServer", "nil")
	server.StartServer()

}

func StartIopCLI() {
	logger.Info("Starting Iop CLI", "starter", "StartIopCLI", "nil")

	cli.StartCli()
}
