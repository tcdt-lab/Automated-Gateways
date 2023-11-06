package iop_cli

import (
	"fmt"
	"github.com/tcdt-lab/Automated-Gateways/relay/mediator"
)

func ShowInsideLedgerPermittedMethodInfoOptions(fsm *StateMachine) {
	fmt.Println("Please Select One of the Following Options: ")
	fmt.Println("1. Create Permitted Method")
	fmt.Println("2. Update Permitted Method")
	fmt.Println("3. Remove Permitted Method")
	fmt.Println("4. Get Permitted Method Info By Id")
	fmt.Println("5. Get Permitted Method By NetworkId")
	fmt.Println("6. Get Permitted Method By Network ID and Index")
	fmt.Println("7. Check If Permitted Method Exist")
	fmt.Println("8. Invoke  Permitted Method")
	fmt.Println("9. Return To Previous Menu")

	var answer string
	fmt.Scanln(&answer)
	switch answer {
	case "1":
		err := fsm.doTransition(EVENT_CREATE_PERMITTED_METHOD)
		if err != nil {
			fmt.Println(err)
			return
		}
	case "2":
		err := fsm.doTransition(EVENT_UPDATE_PERMITTED_METHOD)
		if err != nil {
			fmt.Println(err)
			return
		}
	case "3":
		err := fsm.doTransition(EVENT_REMOVE_PERMITTED_METHOD)
		if err != nil {
			fmt.Println(err)
			return
		}
	case "4":
		err := fsm.doTransition(EVENT_GET_PERMITTED_METHOD_BY_ID)
		if err != nil {
			fmt.Println(err)
			return
		}
	case "5":
		err := fsm.doTransition(EVENT_GET_ALL_PERMITTED_METHODS_BY_ADDRESS)
		if err != nil {
			fmt.Println(err)
			return
		}
	case "6":
		err := fsm.doTransition(EVENT_GET_ALL_PERMITTED_METHODS_BY_ID_AND_INDEX)
		if err != nil {
			fmt.Println(err)
			return
		}
	case "7":
		err := fsm.doTransition(EVENT_PERMITTED_METHOD_EXISTS)
		if err != nil {
			fmt.Println(err)
			return
		}
	case "8":
		err := fsm.doTransition(EVENT_PERMITTED_METHOD_INVOKE)
		if err != nil {
			fmt.Println(err)
			return
		}
	case "9":
		err := fsm.doTransition(EVENT_RETURN_TO_PREVIOUS_MENU)
		if err != nil {
			fmt.Println(err)
			return
		}
	}
}

func ShowCreatePermittedMethodOptions(fsm *StateMachine) {
	var permittedNetworkId string
	var permittedNetworkName string
	var chaincode string
	var channel string
	var inputArgs string
	var outputArgs string
	var mediatorFactory mediator.MediatorFactory
	permittedMethod, err := mediatorFactory.CreatePermittedMethodsMediator(mediator.HYPERLEDGER_FABRIC_NETWROK_TYPE)
	if err != nil {
		fmt.Println(err)
		err = fsm.doTransition(EVENT_RETURN_TO_PREVIOUS_MENU)
		return
	}

	fmt.Println("Please Enter Permitted Network Id:")
	fmt.Scanln(&permittedNetworkId)
	fmt.Println("Please Enter Permitted Network Name:")
	fmt.Scanln(&permittedNetworkName)
	fmt.Println("Please Enter Chaincode Name:")
	fmt.Scanln(&chaincode)
	fmt.Println("Please Enter Channel Name:")
	fmt.Scanln(&channel)
	fmt.Println("Please Enter Input Arguments:")
	fmt.Scanln(&inputArgs)
	fmt.Println("Please Enter Output Arguments:")
	fmt.Scanln(&outputArgs)

	res, err := permittedMethod.AddPermittedMethod(permittedNetworkId, permittedNetworkName, chaincode, channel, inputArgs, outputArgs)
	if err != nil {
		fmt.Println(err)
		err = fsm.doTransition(EVENT_RETURN_TO_PREVIOUS_MENU)
		return
	}
	println("Result: ", res)
}

func ShowUpdatePermittedMethodOptions(fsm *StateMachine) {
	var id string
	var permittedNetworkName string
	var chaincode string
	var channel string
	var inputArgs string
	var outputArgs string
	var mediatorFactory mediator.MediatorFactory
	permittedMethod, err := mediatorFactory.CreatePermittedMethodsMediator(mediator.HYPERLEDGER_FABRIC_NETWROK_TYPE)
	if err != nil {
		fmt.Println(err)
		err = fsm.doTransition(EVENT_RETURN_TO_PREVIOUS_MENU)
		return
	}

	fmt.Println("Please Enter Permitted Method Id:")
	fmt.Scanln(&id)
	fmt.Println("Please Enter Permitted Network Name:")
	fmt.Scanln(&permittedNetworkName)
	fmt.Println("Please Enter Chaincode Name:")
	fmt.Scanln(&chaincode)
	fmt.Println("Please Enter Channel Name:")
	fmt.Scanln(&channel)
	fmt.Println("Please Enter Input Arguments:")
	fmt.Scanln(&inputArgs)
	fmt.Println("Please Enter Output Arguments:")
	fmt.Scanln(&outputArgs)

	err = permittedMethod.UpdatePermittedMethod(id, permittedNetworkName, chaincode, channel, inputArgs, outputArgs)
	if err != nil {
		fmt.Println(err)
		err = fsm.doTransition(EVENT_RETURN_TO_PREVIOUS_MENU)
		return
	}
	println("Result: ", "Permitted Method Updated Successfully")
}

func ShowRemovePermittedMethodOptions(fsm *StateMachine) {
	var id string
	var mediatorFactory mediator.MediatorFactory
	permittedMethod, err := mediatorFactory.CreatePermittedMethodsMediator(mediator.HYPERLEDGER_FABRIC_NETWROK_TYPE)
	if err != nil {
		fmt.Println(err)
		err = fsm.doTransition(EVENT_RETURN_TO_PREVIOUS_MENU)
		return
	}

	fmt.Println("Please Enter Permitted Method Id:")
	fmt.Scanln(&id)

	err = permittedMethod.RemovePermittedMethod(id)
	if err != nil {
		fmt.Println(err)
		err = fsm.doTransition(EVENT_RETURN_TO_PREVIOUS_MENU)
		return
	}
	println("Result: ", "Permitted Method Removed Successfully")
}

func ShowGetPermittedMethodByIdOptions(fsm *StateMachine) {
	var id string
	var mediatorFactory mediator.MediatorFactory
	permittedMethod, err := mediatorFactory.CreatePermittedMethodsMediator(mediator.HYPERLEDGER_FABRIC_NETWROK_TYPE)
	if err != nil {
		fmt.Println(err)
		err = fsm.doTransition(EVENT_RETURN_TO_PREVIOUS_MENU)
		return
	}

	fmt.Println("Please Enter Permitted Method Id:")
	fmt.Scanln(&id)

	res, err := permittedMethod.GetPermittedMethod(id)
	if err != nil {
		fmt.Println(err)
		err = fsm.doTransition(EVENT_RETURN_TO_PREVIOUS_MENU)
		return
	}

	fmt.Println("Permitted Method Id: ", res.PermittedMethodId)
	fmt.Println("Permitted Network Name: ", res.Name)
	fmt.Println("Chaincode Name: ", res.Chaincode)
	fmt.Println("Channel Name: ", res.Channel)
	fmt.Println("Input Arguments: ", res.InputArgs)
	fmt.Println("Output Arguments: ", res.OutputArgs)
}

func ShowGetPermittedMethodsByNetworkIdOptions(fsm *StateMachine) {
	var permittedNetworkId string
	var mediatorFactory mediator.MediatorFactory
	permittedMethod, err := mediatorFactory.CreatePermittedMethodsMediator(mediator.HYPERLEDGER_FABRIC_NETWROK_TYPE)
	if err != nil {
		fmt.Println(err)
		err = fsm.doTransition(EVENT_RETURN_TO_PREVIOUS_MENU)
		return
	}

	fmt.Println("Please Enter Permitted Network Id:")
	fmt.Scanln(&permittedNetworkId)

	res, err := permittedMethod.GetPermittedMethodsByNetworkId(permittedNetworkId)
	if err != nil {
		fmt.Println(err)
		err = fsm.doTransition(EVENT_RETURN_TO_PREVIOUS_MENU)
		return
	}

	for _, method := range res {
		fmt.Println("Permitted Method Id: ", method.PermittedMethodId)
		fmt.Println("Permitted Network Name: ", method.Name)
		fmt.Println("Chaincode Name: ", method.Chaincode)
		fmt.Println("Channel Name: ", method.Channel)
		fmt.Println("Input Arguments: ", method.InputArgs)
		fmt.Println("Output Arguments: ", method.OutputArgs)
	}
}

func ShowGetPermittedMethodsByNetworkIdAndIndexOptions(fsm *StateMachine) {
	var permittedNetworkId string
	var startIndex string
	var endIndex string
	var mediatorFactory mediator.MediatorFactory
	permittedMethod, err := mediatorFactory.CreatePermittedMethodsMediator(mediator.HYPERLEDGER_FABRIC_NETWROK_TYPE)
	if err != nil {
		fmt.Println(err)
		err = fsm.doTransition(EVENT_RETURN_TO_PREVIOUS_MENU)
		return
	}

	fmt.Println("Please Enter Permitted Network Id:")
	fmt.Scanln(&permittedNetworkId)
	fmt.Println("Please Enter Start Index:")
	fmt.Scanln(&startIndex)
	fmt.Println("Please Enter End Index:")
	fmt.Scanln(&endIndex)

	res, err := permittedMethod.GetPermittedMethodsByIndexAndNetworkId(permittedNetworkId, startIndex, endIndex)
	if err != nil {
		fmt.Println(err)
		err = fsm.doTransition(EVENT_RETURN_TO_PREVIOUS_MENU)
		return
	}

	for _, method := range res {
		fmt.Println("Permitted Method Id: ", method.PermittedMethodId)
		fmt.Println("Permitted Network Name: ", method.Name)
		fmt.Println("Chaincode Name: ", method.Chaincode)
		fmt.Println("Channel Name: ", method.Channel)
		fmt.Println("Input Arguments: ", method.InputArgs)
		fmt.Println("Output Arguments: ", method.OutputArgs)
	}
}

func ShowPermittedMethodExistsOptions(fsm *StateMachine) {
	var id string
	var mediatorFactory mediator.MediatorFactory
	permittedMethod, err := mediatorFactory.CreatePermittedMethodsMediator(mediator.HYPERLEDGER_FABRIC_NETWROK_TYPE)
	if err != nil {
		fmt.Println(err)
		err = fsm.doTransition(EVENT_RETURN_TO_PREVIOUS_MENU)
		return
	}

	fmt.Println("Please Enter Permitted Method Id:")
	fmt.Scanln(&id)

	res, err := permittedMethod.PermittedMethodExists(id)
	if err != nil {
		fmt.Println(err)
		err = fsm.doTransition(EVENT_RETURN_TO_PREVIOUS_MENU)
		return
	}

	fmt.Println("Result: ", res)
}

func ShowInvokePermittedMethodOptions(fsm *StateMachine) {
	var name string
	var chaincode string
	var channel string
	var inputArgs []string
	var mediatorFactory mediator.MediatorFactory
	permittedMethod, err := mediatorFactory.CreatePermittedMethodsMediator(mediator.HYPERLEDGER_FABRIC_NETWROK_TYPE)
	if err != nil {
		fmt.Println(err)
		return
	}

	fmt.Println("Please Enter Permitted Method Name:")
	fmt.Scanln(&name)
	fmt.Println("Please Enter Chaincode Name:")
	fmt.Scanln(&chaincode)
	fmt.Println("Please Enter Channel Name:")
	fmt.Scanln(&channel)
	fmt.Println("Please Enter Input Arguments:")
	fmt.Scanln(&inputArgs)

	res, err := permittedMethod.InvokePermittedMethod(name, chaincode, channel, inputArgs)
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println("Result: ", res)
}
