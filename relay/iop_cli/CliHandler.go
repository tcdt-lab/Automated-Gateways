package iop_cli

import (
	"fmt"
	mediator "github.com/tcdt-lab/Automated-Gateways/relay/mediator"
)

type State string
type Event string
type ActionFunction func(*StateMachine)

const (
	EVENT_INITIAL                               = "INITIAL_EVENT"
	EVENT_SELECTING_OUTSIDE_DATA                = "SELECTING_OUTSIDE_DATA_EVENT"
	EVENT_SELECTING_INVOKE_METHOD               = "SELECTING_INVOKE_METHOD_EVENT"
	EVENT_SELECTING_GET_ACCESSIBLE_NETWORK_INFO = "SELECTING_GET_ACCESSIBLE_NETWORK_INFO_EVENT"
	EVENT_SELECTING_GET_ACCESSIBLE_METHOD_LIST  = "SELECTING_GET_ACESSIBLE_METHOD_LIST_EVENT"
	EVENT_RETURN_TO_PREVIOUS_MENU               = "RETURN_TO_PREVIOUS_MENU_EVENT"
)

const (
	STATE_INITIAL                               = "INITIAL_STATE"
	STATE_SELECTING_OUTSIDE_DATA                = "SELECTING_OUTSIDE_DATA_STATE"
	STATE_SELECTING_INVOKE_METHOD               = "SELECTING_INVOKE_METHOD_STATE"
	STATE_SELECTING_GET_ACCESSIBLE_NETWORK_INFO = "SELECTING_GET_ACCESSIBLE_NETWORK_INFO_STATE"
	STATE_SELECTING_GET_ACCESSIBLE_METHOD_LIST  = "SELECTING_GET_ACESSIBLE_METHOD_LIST_STATE"
)

type Action struct {
	destination    State
	actionFunction ActionFunction
}

type Transition map[Event]Action
type StateMap map[State]Transition

type StateMachine struct {
	initial  State
	current  State
	stateMap StateMap
}

func (fsm *StateMachine) getCurrentState() State {
	if fsm.current == "" {
		return STATE_INITIAL
	}
	return fsm.current
}

func (fsm *StateMachine) doTransition(event Event) error {
	action := fsm.stateMap[fsm.getCurrentState()][event]
	fsm.current = action.destination
	if action.actionFunction != nil {
		action.actionFunction(fsm)
	}

	return nil
}

func (fsm *StateMachine) FsmCreator() *StateMachine {

	return &StateMachine{
		initial: STATE_INITIAL,
		stateMap: StateMap{
			STATE_INITIAL: {
				EVENT_SELECTING_OUTSIDE_DATA: {
					actionFunction: showOutsideDataOption,
					destination:    STATE_SELECTING_OUTSIDE_DATA,
				},
				EVENT_INITIAL: {
					actionFunction: showInitialOption,
					destination:    STATE_INITIAL,
				},
			},
			STATE_SELECTING_OUTSIDE_DATA: {
				EVENT_SELECTING_INVOKE_METHOD: {
					actionFunction: showInvokeMethodOption,
					destination:    STATE_SELECTING_INVOKE_METHOD,
				},
				EVENT_SELECTING_GET_ACCESSIBLE_NETWORK_INFO: {
					actionFunction: showGetAccessibleNetworkInfoOption,
					destination:    STATE_SELECTING_GET_ACCESSIBLE_NETWORK_INFO,
				},
				EVENT_SELECTING_GET_ACCESSIBLE_METHOD_LIST: {
					actionFunction: showGetAccessibleMethodListOption,
					destination:    STATE_SELECTING_GET_ACCESSIBLE_METHOD_LIST,
				},
				EVENT_RETURN_TO_PREVIOUS_MENU: {
					actionFunction: showInitialOption,
					destination:    STATE_INITIAL,
				},
			},
			STATE_SELECTING_INVOKE_METHOD: {
				EVENT_RETURN_TO_PREVIOUS_MENU: {
					actionFunction: showOutsideDataOption,
					destination:    STATE_SELECTING_OUTSIDE_DATA,
				},
			},
			STATE_SELECTING_GET_ACCESSIBLE_NETWORK_INFO: {
				EVENT_RETURN_TO_PREVIOUS_MENU: {
					actionFunction: showOutsideDataOption,
					destination:    STATE_SELECTING_OUTSIDE_DATA,
				},
			},
			STATE_SELECTING_GET_ACCESSIBLE_METHOD_LIST: {
				EVENT_RETURN_TO_PREVIOUS_MENU: {
					actionFunction: showOutsideDataOption,
					destination:    STATE_SELECTING_OUTSIDE_DATA,
				},
			},
		},
	}
}

func showInitialOption(fsm *StateMachine) {
	fmt.Println("Welcome To  Automated Gateways IOP Module! ")
	fmt.Println("Please Select One Of The Following Options: ")
	fmt.Println("1. Get Information From Outside Ledger")
	fmt.Println("2.  Get Information From Your Ledger")

	// var then variable name then variable type
	var answer string
	fmt.Scanln(&answer)
	switch answer {
	case "1":
		fmt.Println("You Selected Option 1")
		err := fsm.doTransition(EVENT_SELECTING_OUTSIDE_DATA)
		if err != nil {
			fmt.Println(err)
		}
	case "2":
		fmt.Println("You Selected Option 2")
	}

}

func showOutsideDataOption(fsm *StateMachine) {
	fmt.Println("Please Select One Of The Following Options: ")
	fmt.Println("1. Invoke Method")
	fmt.Println("2. Get Accessible Network Information")
	fmt.Println("3. Get Accessible Method List")
	fmt.Println("4. Return To Previous Menu")

	var answer string
	fmt.Scanln(&answer)
	switch answer {
	case "1":

		err := fsm.doTransition(EVENT_SELECTING_INVOKE_METHOD)
		if err != nil {
			fmt.Println(err)
		}
	case "2":

		err := fsm.doTransition(EVENT_SELECTING_GET_ACCESSIBLE_NETWORK_INFO)
		if err != nil {
			fmt.Println(err)
		}
	case "3":

		err := fsm.doTransition(EVENT_SELECTING_GET_ACCESSIBLE_METHOD_LIST)
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

func StartCli() {
	var fsm = new(StateMachine)
	fsm = fsm.FsmCreator()
	err := fsm.doTransition(EVENT_INITIAL)
	if err != nil {
		fmt.Println(err)
	}
}
