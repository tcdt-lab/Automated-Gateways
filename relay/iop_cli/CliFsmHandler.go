package iop_cli

import (
	"fmt"
)

type State string
type Event string
type ActionFunction func(*StateMachine)

const (
	EVENT_INITIAL                 = "INITIAL_EVENT"
	EVENT_SELECTING_OUTSIDE_DATA  = "SELECTING_OUTSIDE_DATA_EVENT"
	EVENT_SELECTING_INSIDE_DATA   = "SELECTING_INSIDE_DATA_EVENT"
	EVENT_RETURN_TO_PREVIOUS_MENU = "RETURN_TO_PREVIOUS_MENU_EVENT"

	//*********************** OUTSIDE DATA EVENTS *************************
	EVENT_SELECTING_INVOKE_METHOD               = "SELECTING_INVOKE_METHOD_EVENT"
	EVENT_SELECTING_GET_ACCESSIBLE_NETWORK_INFO = "SELECTING_GET_ACCESSIBLE_NETWORK_INFO_EVENT"
	EVENT_SELECTING_GET_ACCESSIBLE_METHOD_LIST  = "SELECTING_GET_ACCESSIBLE_METHOD_LIST_EVENT"

	//*********************** INSIDE DATA EVENTS *************************
	EVENT_SELECTING_ACCESSIBLE_NETWORK_OPTIONS   = "SELECTING_ACCESSIBLE_NETWORK_OPTIONS_EVENT"
	EVENT_CREATE_ACCESSIBLE_NETWORK              = "CREATE_ACCESSIBLE_NETWORK_EVENT"
	EVENT_REMOVE_ACCESSIBLE_NETWORK              = "REMOVE_ACCESSIBLE_NETWORK_EVENT"
	EVENT_UPDATE_ACCESSIBLE_NETWORK              = "UPDATE_ACCESSIBLE_NETWORK_EVENT"
	EVENT_GET_ACCESSIBLE_NETWORK_BY_ID           = "GET_ACCESSIBLE_NETWORK_BY_ID_EVENT"
	EVENT_GET_ALL_ACCESSIBLE_NETWORKS_BY_ADDRESS = "GET_ALL_ACCESSIBLE_NETWORKS_BY_ADDRESS_EVENT"
	EVENT_ACCESSIBLE_NETWORK_EXISTS              = "ACCESSIBLE_NETWORK_EXISTS_EVENT"

	EVENT_SELECTING_PERMITTED_NETWORK_OPTIONS   = "SELECTING_PERMITTED_NETWORK_OPTIONS_EVENT"
	EVENT_CREATE_PERMITTED_NETWORK              = "CREATE_PERMITTED_NETWORK_EVENT"
	EVENT_REMOVE_PERMITTED_NETWORK              = "REMOVE_PERMITTED_NETWORK_EVENT"
	EVENT_UPDATE_PERMITTED_NETWORK              = "UPDATE_PERMITTED_NETWORK_EVENT"
	EVENT_GET_PERMITTED_NETWORK                 = "GET_PERMITTED_NETWORK_EVENT"
	EVENT_GET_ALL_PERMITTED_NETWORKS_BY_ADDRESS = "GET_ALL_PERMITTED_NETWORKS_BY_ADDRESS_EVENT"
	EVENT_PERMITTED_NETWORK_EXISTS              = "PERMITTED_NETWORK_EXISTS_EVENT"

	EVENT_SELECTING_PERMITTED_METHOD_OPTIONS   = "SELECTING_PERMITTED_METHOD_OPTIONS_EVENT"
	EVENT_CREATE_PERMITTED_METHOD              = "CREATE_PERMITTED_METHOD_EVENT"
	EVENT_REMOVE_PERMITTED_METHOD              = "REMOVE_PERMITTED_METHOD_EVENT"
	EVENT_UPDATE_PERMITTED_METHOD              = "UPDATE_PERMITTED_METHOD_EVENT"
	EVENT_GET_PERMITTED_METHOD                 = "GET_PERMITTED_METHOD_EVENT"
	EVENT_GET_ALL_PERMITTED_METHODS_BY_ADDRESS = "GET_ALL_PERMITTED_METHODS_BY_ADDRESS_EVENT"
	EVENT_PERMITTED_METHOD_EXISTS              = "PERMITTED_METHOD_EXISTS_EVENT"
)

const (
	STATE_INITIAL                = "INITIAL_STATE"
	STATE_SELECTING_INSIDE_DATA  = "SELECTING_INSIDE_DATA_STATE"
	STATE_SELECTING_OUTSIDE_DATA = "SELECTING_OUTSIDE_DATA_STATE"

	//*********************** Outside DATA STATES *************************
	STATE_SELECTING_INVOKE_METHOD               = "SELECTING_INVOKE_METHOD_STATE"
	STATE_SELECTING_GET_ACCESSIBLE_NETWORK_INFO = "SELECTING_GET_ACCESSIBLE_NETWORK_INFO_STATE"
	STATE_SELECTING_GET_ACCESSIBLE_METHOD_LIST  = "SELECTING_GET_ACESSIBLE_METHOD_LIST_STATE"

	//*********************** Inside DATA STATES *************************
	STATE_SELECTING_ACCESSIBLE_NETWORK_OPTIONS        = "SELECTING_ACCESSIBLE_NETWORK_OPTIONS_STATE"
	STATE_SELECTING_GET_ACCESSIBLE_NETWORK_BY_ADDRESS = "SELECTING_GET_ACCESSIBLE_NETWORK_BY_ADDRESS_STATE"
	STATE_SELECTING_GET_ALL_ACCESSIBLE_NETWORKS       = "SELECTING_GET_ALL_ACCESSIBLE_NETWORKS_STATE"
	STATE_SELECTING_CREATE_ACCESSIBLE_NETWORK         = "SELECTING_CREATE_ACCESSIBLE_NETWORK_STATE"
	STATE_SELECTING_REMOVE_ACCESSIBLE_NETWORK         = "SELECTING_REMOVE_ACCESSIBLE_NETWORK_STATE"
	STATE_SELECTING_UPDATE_ACCESSIBLE_NETWORK         = "SELECTING_UPDATE_ACCESSIBLE_NETWORK_STATE"
	STATE_SELECTING_ACCESSIBLE_NETWORK_EXISTS         = "SELECTING_ACCESSIBLE_NETWORK_EXISTS_STATE"

	STATE_SELECTING_PERMITTED_NETWORK_OPTIONS = "SELECTING_PERMITTED_NETWORK_OPTIONS_STATE"
	STATE_SELECTING_PERMITTED_METHOD_OPTIONS  = "SELECTING_PERMITTED_METHOD_OPTIONS_STATE"
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
				EVENT_SELECTING_INSIDE_DATA: {
					actionFunction: showInsideDataOption,
					destination:    STATE_SELECTING_INSIDE_DATA,
				},
				EVENT_INITIAL: {
					actionFunction: showInitialOption,
					destination:    STATE_INITIAL,
				},
			},

			//*********************** OUTSIDE DATA STATES *************************
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

			//*********************** IOP STATES *************************
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

			//*********************** INSIDE DATA STATES *************************
			STATE_SELECTING_INSIDE_DATA: {
				EVENT_SELECTING_ACCESSIBLE_NETWORK_OPTIONS: {
					actionFunction: ShowInsideLedgerAccessibleNetworkInfoOptions,
					destination:    STATE_SELECTING_ACCESSIBLE_NETWORK_OPTIONS,
				},
				EVENT_SELECTING_PERMITTED_NETWORK_OPTIONS: {
					actionFunction: ShowInsideLedgerPermittedNetworkInfoOptions,
					destination:    STATE_SELECTING_PERMITTED_NETWORK_OPTIONS,
				},
				EVENT_SELECTING_PERMITTED_METHOD_OPTIONS: {
					actionFunction: ShowInsideLedgerPermittedMethodInfoOptions,
					destination:    STATE_SELECTING_PERMITTED_METHOD_OPTIONS,
				},
				EVENT_RETURN_TO_PREVIOUS_MENU: {
					actionFunction: showInitialOption,
					destination:    STATE_INITIAL,
				},
			},

			//*********************** Accessible Network STATES *************************
			STATE_SELECTING_ACCESSIBLE_NETWORK_OPTIONS: {
				EVENT_CREATE_ACCESSIBLE_NETWORK: {
					actionFunction: ShowCreateAccessibleNetworkOptions,
					destination:    STATE_SELECTING_CREATE_ACCESSIBLE_NETWORK,
				},
				EVENT_REMOVE_ACCESSIBLE_NETWORK: {
					actionFunction: ShowRemoveAccessibleNetworkOptions,
					destination:    STATE_SELECTING_REMOVE_ACCESSIBLE_NETWORK,
				},
				EVENT_UPDATE_ACCESSIBLE_NETWORK: {
					actionFunction: ShowUpdateAccessibleNetworkInfo,
					destination:    STATE_SELECTING_UPDATE_ACCESSIBLE_NETWORK,
				},
				EVENT_GET_ACCESSIBLE_NETWORK_BY_ID: {
					actionFunction: ShowGetAccessibleNetworkInfoByIdOptions,
					destination:    STATE_SELECTING_GET_ACCESSIBLE_NETWORK_BY_ADDRESS,
				},
				EVENT_GET_ALL_ACCESSIBLE_NETWORKS_BY_ADDRESS: {
					actionFunction: ShowGetAllAccessibleNetworksByAddressOptions,
					destination:    STATE_SELECTING_GET_ALL_ACCESSIBLE_NETWORKS,
				},
				EVENT_ACCESSIBLE_NETWORK_EXISTS: {
					actionFunction: ShowIfAccessibleNetworkExists,
					destination:    STATE_SELECTING_ACCESSIBLE_NETWORK_EXISTS,
				},
				EVENT_RETURN_TO_PREVIOUS_MENU: {
					actionFunction: showInsideDataOption,
					destination:    STATE_SELECTING_INSIDE_DATA,
				},
			},
			STATE_SELECTING_CREATE_ACCESSIBLE_NETWORK: {
				EVENT_RETURN_TO_PREVIOUS_MENU: {
					actionFunction: ShowInsideLedgerAccessibleNetworkInfoOptions,
					destination:    STATE_SELECTING_ACCESSIBLE_NETWORK_OPTIONS,
				},
			},
			STATE_SELECTING_REMOVE_ACCESSIBLE_NETWORK: {
				EVENT_RETURN_TO_PREVIOUS_MENU: {
					actionFunction: ShowInsideLedgerAccessibleNetworkInfoOptions,
					destination:    STATE_SELECTING_ACCESSIBLE_NETWORK_OPTIONS,
				},
			},
			STATE_SELECTING_UPDATE_ACCESSIBLE_NETWORK: {
				EVENT_RETURN_TO_PREVIOUS_MENU: {
					actionFunction: ShowInsideLedgerAccessibleNetworkInfoOptions,
					destination:    STATE_SELECTING_ACCESSIBLE_NETWORK_OPTIONS,
				},
			},
			STATE_SELECTING_GET_ACCESSIBLE_NETWORK_BY_ADDRESS: {
				EVENT_RETURN_TO_PREVIOUS_MENU: {
					actionFunction: ShowInsideLedgerAccessibleNetworkInfoOptions,
					destination:    STATE_SELECTING_ACCESSIBLE_NETWORK_OPTIONS,
				},
			},
			STATE_SELECTING_GET_ALL_ACCESSIBLE_NETWORKS: {
				EVENT_RETURN_TO_PREVIOUS_MENU: {
					actionFunction: ShowInsideLedgerAccessibleNetworkInfoOptions,
					destination:    STATE_SELECTING_ACCESSIBLE_NETWORK_OPTIONS,
				},
			},
			STATE_SELECTING_ACCESSIBLE_NETWORK_EXISTS: {
				EVENT_RETURN_TO_PREVIOUS_MENU: {
					actionFunction: ShowInsideLedgerAccessibleNetworkInfoOptions,
					destination:    STATE_SELECTING_ACCESSIBLE_NETWORK_OPTIONS,
				},
			},
			//*********************** Permitted Network STATES *************************
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
		err := fsm.doTransition(EVENT_SELECTING_OUTSIDE_DATA)
		if err != nil {
			fmt.Println(err)
		}
	case "2":
		err := fsm.doTransition(EVENT_SELECTING_INSIDE_DATA)
		if err != nil {
			fmt.Println(err)
		}

	}

}

func showOutsideDataOption(fsm *StateMachine) {
	fmt.Println("Please Select One Of The Following Options: ")

	fmt.Println("1. Get Accessible Network Information")
	fmt.Println("2. Get Accessible Method List")
	fmt.Println("3. Invoke Method")
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

func showInsideDataOption(fsm *StateMachine) {
	fmt.Println("Please Select One Of The Following Options: ")
	fmt.Println("1. Accessible Network Handler")
	fmt.Println("2. Permitted Network Handler")
	fmt.Println("3. Permitted Method Handler")
	fmt.Println("4. Return To Previous Menu")

	var answer string
	fmt.Scanln(&answer)
	switch answer {
	case "1":
		err := fsm.doTransition(EVENT_SELECTING_ACCESSIBLE_NETWORK_OPTIONS)
		if err != nil {
			fmt.Println(err)
		}
	case "2":
		err := fsm.doTransition(EVENT_SELECTING_PERMITTED_NETWORK_OPTIONS)
		if err != nil {
			fmt.Println(err)
		}
	case "3":
		err := fsm.doTransition(EVENT_SELECTING_PERMITTED_METHOD_OPTIONS)
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

func StartCli() {
	var fsm = new(StateMachine)
	fsm = fsm.FsmCreator()
	err := fsm.doTransition(EVENT_INITIAL)
	if err != nil {
		fmt.Println(err)
	}
}
