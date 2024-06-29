package main

import (
	"github.com/tcdt-lab/Automated-Gateways/relay/configs"
	"github.com/tcdt-lab/Automated-Gateways/relay/mediator"
	"github.com/tcdt-lab/Automated-Gateways/relay/starter"
	"sync"
	"time"
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

	wg.Add(1)
	go func() {
		time.Sleep(time.Second * 20)
		defer wg.Done()
		for true {
			time.Sleep(time.Second * 5)
			var iopMediator mediator.IopMediator
			res, err := iopMediator.InvokeAccessibleMethod("AccessibleNetwork_koosha.com_1", "AddTwoNumbers", "addition_1", "mychannel", "[\"12\",\"14\"]", "int")
			if err != "nil" {
				println(err)
			}
			println("Res in CODE *****************")
			println(res)
		}

	}()

	wg.Wait()

}
