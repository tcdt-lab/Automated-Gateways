package iop_cli

import (
	"fmt"
	"github.com/tcdt-lab/Automated-Gateways/relay/mediator"
)

func showInvokeMethodOption(fsm *StateMachine) {
	var name string
	var chaincode string
	var channel string
	var inputArgs string
	var output string
	var iopMediator mediator.IopMediator
	fmt.Println("Please Insert the Required Information: ")
	fmt.Println("1. Method Name")
	fmt.Scanln(&name)
	fmt.Println("2. Chaincode")
	fmt.Scanln(&chaincode)
	fmt.Println("3. channel")
	fmt.Scanln(&channel)
	fmt.Println("4. inputArgs")
	fmt.Scanln(&inputArgs)
	fmt.Println("5. output")
	fmt.Scanln(&output)

	iopMediator.InvokeAccessibleMethod(name, chaincode, channel, inputArgs, output)

	fsm.doTransition(EVENT_RETURN_TO_PREVIOUS_MENU)
}

func showGetAccessibleNetworkInfoOption(fsm *StateMachine) {
	var address string
	var iopMediator mediator.IopMediator
	fmt.Println("Please Insert the Required Information: ")
	fmt.Println("1. Address")
	fmt.Scanln(&address)

	res, err := iopMediator.GetAccessibleNetworkInfo(address)
	if err != nil {
		fmt.Println(err)
	}
	for _, network := range res {

		fmt.Println("Network Id: ", network.AccessibleNetworkId)
		fmt.Println("Network Name: ", network.NetworkName)
		fmt.Println("Network IP: ", network.IP)
		fmt.Println("Network Address: ", network.ADDRESS)
		fmt.Println("Company Name: ", network.CompanyName)
		fmt.Println("*******************************************************")
	}
	fsm.doTransition(EVENT_RETURN_TO_PREVIOUS_MENU)
}

func showGetAccessibleMethodListOption(fsm *StateMachine) {
	var networkId string
	var iopMediator mediator.IopMediator
	fmt.Println("Please Insert the Required Information: ")
	fmt.Println("1. Network Id")
	fmt.Scanln(&networkId)
	res, err := iopMediator.GetAccessibleMethodsList(networkId)

	if err != nil {
		fmt.Println(err)
	}

	for _, method := range res {
		fmt.Println("Method Id: ", method.PermittedMethodId)
		fmt.Println("Method Name: ", method.Name)
		fmt.Println("Chaincode Name: ", method.Chaincode)
		fmt.Println("Channel Name: ", method.Channel)
		fmt.Println("Input Args: ", method.InputArgs)
		fmt.Println("Output Args: ", method.OutputArgs)
		fmt.Println("*******************************************************")
	}

	fsm.doTransition(EVENT_RETURN_TO_PREVIOUS_MENU)
}
