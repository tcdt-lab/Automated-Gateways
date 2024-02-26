package main

import (
	"github.com/tcdt-lab/Automated-Gateways/relay/configs"
	"github.com/tcdt-lab/Automated-Gateways/relay/starter"
	"sync"
)

func main() {
	var wg sync.WaitGroup
	configs.YamlAddress = "/home/koosha/Desktop/Automated-Gateways/relay/cmd/config.yaml"
	wg.Add(1)
	go func() {

		defer wg.Done()
		starter.StartIopServer()
	}()
	wg.Add(1)
	go func() {

		defer wg.Done()
		starter.StartIopCLI()
	}()

	wg.Wait()

}
