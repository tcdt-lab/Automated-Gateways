package main

import (
	"github.com/tcdt-lab/Automated-Gateways/relay/configs"
	"github.com/tcdt-lab/Automated-Gateways/relay/mediator"
	"github.com/tcdt-lab/Automated-Gateways/relay/starter"
	"sync"
)

func main() {
	var wg sync.WaitGroup
	configs.YamlAddress = "/home/koosha/Desktop/Automated-Gateways/relay/cmd/config.yaml"

	go func() {
		wg.Add(1)
		defer wg.Done()
		starter.StartIopServer()
	}()

	go func() {
		wg.Add(1)
		defer wg.Done()
		starter.StartIopCLI()
	}()

	var mediatorFactory mediator.MediatorFactory

	permittedNetwork, err := mediatorFactory.CreatePermittedNetworkMediator(mediator.HYPERLEDGER_FABRIC_NETWROK_TYPE)
	if err != nil {
		panic(err)
	}
	permittedNt, err := permittedNetwork.GetAllPermittedNetworksByAddress("localhost")

	for _, v := range permittedNt {
		println(v.NetworkName)
	}
	wg.Wait()

}
