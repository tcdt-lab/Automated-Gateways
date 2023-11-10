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

	iopMediator.InvokeAccessibleMethod(fsm.selectedAccessibleNetworkId, name, chaincode, channel, inputArgs, output)

	fsm.doTransition(EVENT_RETURN_TO_PREVIOUS_MENU)
}

func showGetAccessibleNetworkInfoOption(fsm *StateMachine) {
	var address string
	var iopMediator mediator.IopMediator
	fmt.Println("Please Insert the Required Information: ")
	fmt.Println("1. Your Address")
	fmt.Scanln(&address)

	res, err := iopMediator.GetSelfInformationOnOutsideNetworks(fsm.selectedAccessibleNetworkId, address)
	if err != nil {
		fmt.Println(err)
	}
	for _, network := range res {

		fmt.Println("Your Registered Network Id: ", network.AccessibleNetworkId)
		fmt.Println("Your Registered Network Name: ", network.NetworkName)
		fmt.Println("Your Registered Network IP: ", network.IP)
		fmt.Println("Network Registered Address: ", network.ADDRESS)
		fmt.Println("Company Registered Name: ", network.CompanyName)
		fmt.Println("*******************************************************")
	}
	fsm.doTransition(EVENT_RETURN_TO_PREVIOUS_MENU)
}

func showGetAccessibleMethodListOption(fsm *StateMachine) {
	var networkId string
	var iopMediator mediator.IopMediator
	fmt.Println("Please Insert the Required Information: ")
	fmt.Println("1. Your Network Id")
	fmt.Scanln(&networkId)
	res, err := iopMediator.GetAccessibleMethodsList(fsm.selectedAccessibleNetworkId, networkId)

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

func showOutsideDataOption(fsm *StateMachine) {
	fmt.Println("Please Select One Of The Following Options: ")
	fmt.Println("1. Get Your Network Information From Outside Ledger")
	fmt.Println("2. Get Your Accessible Method List From Outside Ledger")
	fmt.Println("3. Invoke Method On Outside Ledger")
	fmt.Println("4. Return To Previous Menu")
	var answer string
	fmt.Scanln(&answer)

	switch answer {
	case "1":
		err := fsm.doTransition(EVENT_SELECTING_GET_ACCESSIBLE_NETWORK_INFO)
		if err != nil {
			fmt.Println(err)
		}
	case "2":
		err := fsm.doTransition(EVENT_SELECTING_GET_ACCESSIBLE_METHOD_LIST)
		if err != nil {
			fmt.Println(err)
		}
	case "3":
		err := fsm.doTransition(EVENT_SELECTING_INVOKE_METHOD)
		if err != nil {
			fmt.Println(err)
		}
	case "4":
		err := fsm.doTransition(EVENT_RETURN_TO_PREVIOUS_MENU)
		if err != nil {
			fmt.Println(err)
		}
	}
}
