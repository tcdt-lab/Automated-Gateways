package iop_cli

import (
	"fmt"
	"github.com/tcdt-lab/Automated-Gateways/relay/configs"
	"strconv"
	"time"
)

type State string
type Event string
type ActionFunction func(*StateMachine)

const (
	EVENT_INITIAL                        = "INITIAL_EVENT"
	EVENT_SELECTING_OUTSIDE_DATA         = "SELECTING_OUTSIDE_DATA_EVENT"
	EVENT_SELECTING_OUTSIDE_DATA_OPTIONS = "SELECTING_OUTSIDE_DATA_OPTIONS_EVENT"
	EVENT_SELECTING_INSIDE_DATA_OPTIONS  = "SELECTING_INSIDE_DATA_EVENT"
	EVENT_RETURN_TO_PREVIOUS_MENU        = "RETURN_TO_PREVIOUS_MENU_EVENT"

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

	EVENT_SELECTING_PERMITTED_NETWORK_OPTIONS        = "SELECTING_PERMITTED_NETWORK_OPTIONS_EVENT"
	EVENT_CREATE_PERMITTED_NETWORK                   = "CREATE_PERMITTED_NETWORK_EVENT"
	EVENT_REMOVE_PERMITTED_NETWORK                   = "REMOVE_PERMITTED_NETWORK_EVENT"
	EVENT_UPDATE_PERMITTED_NETWORK                   = "UPDATE_PERMITTED_NETWORK_EVENT"
	EVENT_GET_PERMITTED_NETWORK_BY_ID                = "GET_PERMITTED_NETWORK_EVENT"
	EVENT_GET_ALL_PERMITTED_NETWORKS_BY_ADDRESS      = "GET_ALL_PERMITTED_NETWORKS_BY_ADDRESS_EVENT"
	EVENT_GET_PERMITTED_NETWORK_BY_INDEX_AND_ADDRESS = "GET_PERMITTED_NETWORK_BY_INDEX_AND_ADDRESS_EVENT"
	EVENT_PERMITTED_NETWORK_EXISTS                   = "PERMITTED_NETWORK_EXISTS_EVENT"

	EVENT_SELECTING_PERMITTED_METHOD_OPTIONS        = "SELECTING_PERMITTED_METHOD_OPTIONS_EVENT"
	EVENT_CREATE_PERMITTED_METHOD                   = "CREATE_PERMITTED_METHOD_EVENT"
	EVENT_REMOVE_PERMITTED_METHOD                   = "REMOVE_PERMITTED_METHOD_EVENT"
	EVENT_UPDATE_PERMITTED_METHOD                   = "UPDATE_PERMITTED_METHOD_EVENT"
	EVENT_GET_PERMITTED_METHOD_BY_ID                = "GET_PERMITTED_METHOD_BY_ID_EVENT"
	EVENT_GET_ALL_PERMITTED_METHODS_BY_ADDRESS      = "GET_ALL_PERMITTED_METHODS_BY_ADDRESS_EVENT"
	EVENT_GET_ALL_PERMITTED_METHODS_BY_ID_AND_INDEX = "GET_ALL_PERMITTED_METHODS_BY_ID_AND_INDEX_EVENT"
	EVENT_PERMITTED_METHOD_EXISTS                   = "PERMITTED_METHOD_EXISTS_EVENT"
	EVENT_PERMITTED_METHOD_INVOKE                   = "PERMITTED_METHOD_INVOKE_EVENT"
)

const (
	STATE_INITIAL                           = "INITIAL_STATE"
	STATE_SELECTING_INSIDE_NETWORK_OPTIONS  = "SELECTING_INSIDE_DATA_STATE"
	STATE_SELECTING_ACCESSIBLE_NETWORK      = "SELECTING_ACCESSIBLE_NETWORK_STATE"
	STATE_SELECTING_OUTSIDE_NETWORK_OPTIONS = "SELECTING_OUTSIDE_DATA_OPTIONS_STATE"
	//*********************** OUTSIDE DATA STATES ***********************************
	STATE_SELECTING_INVOKE_METHOD               = "SELECTING_INVOKE_METHOD_STATE"
	STATE_SELECTING_GET_ACCESSIBLE_NETWORK_INFO = "SELECTING_GET_ACCESSIBLE_NETWORK_INFO_STATE"
	STATE_SELECTING_GET_ACCESSIBLE_METHOD_LIST  = "SELECTING_GET_ACESSIBLE_METHOD_LIST_STATE"

	//*********************** INSIDE DATA STATES ************************************

	//*********************** ACCESSIBLE NETWORK STATES **********************
	STATE_SELECTING_ACCESSIBLE_NETWORK_OPTIONS        = "SELECTING_ACCESSIBLE_NETWORK_OPTIONS_STATE"
	STATE_SELECTING_GET_ACCESSIBLE_NETWORK_BY_ADDRESS = "SELECTING_GET_ACCESSIBLE_NETWORK_BY_ADDRESS_STATE"
	STATE_SELECTING_GET_ALL_ACCESSIBLE_NETWORKS       = "SELECTING_GET_ALL_ACCESSIBLE_NETWORKS_STATE"
	STATE_SELECTING_CREATE_ACCESSIBLE_NETWORK         = "SELECTING_CREATE_ACCESSIBLE_NETWORK_STATE"
	STATE_SELECTING_REMOVE_ACCESSIBLE_NETWORK         = "SELECTING_REMOVE_ACCESSIBLE_NETWORK_STATE"
	STATE_SELECTING_UPDATE_ACCESSIBLE_NETWORK         = "SELECTING_UPDATE_ACCESSIBLE_NETWORK_STATE"
	STATE_SELECTING_ACCESSIBLE_NETWORK_EXISTS         = "SELECTING_ACCESSIBLE_NETWORK_EXISTS_STATE"
	//*********************** PERMITTED NETWORK STATES **********************
	STATE_SELECTING_PERMITTED_NETWORK_OPTIONS                  = "SELECTING_PERMITTED_NETWORK_OPTIONS_STATE"
	STATE_SELECTING_CREATE_PERMITTED_NETWORK                   = "SELECTING_CREATE_PERMITTED_NETWORK_STATE"
	STATE_SELECTING_REMOVE_PERMITTED_NETWORK                   = "SELECTING_REMOVE_PERMITTED_NETWORK_STATE"
	STATE_SELECTING_UPDATE_PERMITTED_NETWORK                   = "SELECTING_UPDATE_PERMITTED_NETWORK_STATE"
	STATE_SELECTING_GET_PERMITTED_NETWORK                      = "SELECTING_GET_PERMITTED_NETWORK_STATE"
	STATE_SELECTING_GET_ALL_PERMITTED_NETWORKS_BY_ADDRESS      = "SELECTING_GET_ALL_PERMITTED_NETWORKS_BY_ADDRESS_STATE"
	STATE_SELECTING_GET_PERMITTED_NETWORK_BY_INDEX_AND_ADDRESS = "SELECTING_GET_PERMITTED_NETWORK_BY_INDEX_AND_ADDRESS_STATE"
	STATE_SELECTING_PERMITTED_NETWORK_EXISTS                   = "SELECTING_PERMITTED_NETWORK_EXISTS_STATE"
	//*********************** PERMITTED METHODS STATES **********************
	STATE_SELECTING_PERMITTED_METHOD_OPTIONS                  = "SELECTING_PERMITTED_METHOD_OPTIONS_STATE"
	STATE_SELECTING_CREATE_PERMITTED_METHOD                   = "SELECTING_CREATE_PERMITTED_METHOD_STATE"
	STATE_SELECTING_REMOVE_PERMITTED_METHOD                   = "SELECTING_REMOVE_PERMITTED_METHOD_STATE"
	STATE_SELECTING_UPDATE_PERMITTED_METHOD                   = "SELECTING_UPDATE_PERMITTED_METHOD_STATE"
	STATE_SELECTING_GET_PERMITTED_METHOD_BY_ID                = "SELECTING_GET_PERMITTED_METHOD_BY_ID_STATE"
	STATE_SELECTING_GET_ALL_PERMITTED_METHODS_BY_ADDRESS      = "SELECTING_GET_ALL_PERMITTED_METHODS_BY_ADDRESS_STATE"
	STATE_SELECTING_GET_ALL_PERMITTED_METHODS_BY_ID_AND_INDEX = "SELECTING_GET_ALL_PERMITTED_METHODS_BY_ID_AND_INDEX_STATE"
	STATE_SELECTING_PERMITTED_METHOD_EXISTS                   = "SELECTING_PERMITTED_METHOD_EXISTS_STATE"
	STATE_SELECTING_INVOKE_PERMITTED_METHOD_BY_ID             = "STATE_SELECTING_INVOKE_PERMITTED_METHOD_BY_ID_STATE"
)

type Action struct {
	destination    State
	actionFunction ActionFunction
}

type Transition map[Event]Action
type StateMap map[State]Transition

type StateMachine struct {
	initial                     State
	current                     State
	stateMap                    StateMap
	selectedAccessibleNetworkId string
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
					actionFunction: showSelectAccessibleNetworksOption,
					destination:    STATE_SELECTING_ACCESSIBLE_NETWORK,
				},
				EVENT_SELECTING_INSIDE_DATA_OPTIONS: {
					actionFunction: showInsideDataOption,
					destination:    STATE_SELECTING_INSIDE_NETWORK_OPTIONS,
				},
				EVENT_INITIAL: {
					actionFunction: showInitialOption,
					destination:    STATE_INITIAL,
				},
			},
			//*********************** OUTSIDE DATA STATES *************************
			STATE_SELECTING_ACCESSIBLE_NETWORK: {
				EVENT_SELECTING_OUTSIDE_DATA_OPTIONS: {
					actionFunction: showOutsideDataOption,
					destination:    STATE_SELECTING_OUTSIDE_NETWORK_OPTIONS,
				},
				EVENT_RETURN_TO_PREVIOUS_MENU: {
					actionFunction: showInitialOption,
					destination:    STATE_INITIAL,
				},
			},
			STATE_SELECTING_OUTSIDE_NETWORK_OPTIONS: {
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
					actionFunction: showSelectAccessibleNetworksOption,
					destination:    STATE_SELECTING_ACCESSIBLE_NETWORK,
				},
			},
			//*********************** IOP STATES *************************
			STATE_SELECTING_INVOKE_METHOD: {
				EVENT_RETURN_TO_PREVIOUS_MENU: {
					actionFunction: showOutsideDataOption,
<<<<<<< HEAD
					destination:    STATE_SELECTING_ACCESSIBLE_NETWORK,
=======
					destination:    STATE_SELECTING_OUTSIDE_DATA_OPTIONS,
>>>>>>> 3c2eb0d... Minor changes have been applied
				},
			},
			STATE_SELECTING_GET_ACCESSIBLE_NETWORK_INFO: {
				EVENT_RETURN_TO_PREVIOUS_MENU: {
					actionFunction: showOutsideDataOption,
<<<<<<< HEAD
					destination:    STATE_SELECTING_ACCESSIBLE_NETWORK,
=======
					destination:    STATE_SELECTING_OUTSIDE_DATA_OPTIONS,
>>>>>>> 3c2eb0d... Minor changes have been applied
				},
			},
			STATE_SELECTING_GET_ACCESSIBLE_METHOD_LIST: {
				EVENT_RETURN_TO_PREVIOUS_MENU: {
					actionFunction: showOutsideDataOption,
<<<<<<< HEAD
					destination:    STATE_SELECTING_ACCESSIBLE_NETWORK,
=======
					destination:    STATE_SELECTING_OUTSIDE_DATA_OPTIONS,
>>>>>>> 3c2eb0d... Minor changes have been applied
				},
			},
			//*********************** INSIDE DATA STATES *************************
			STATE_SELECTING_INSIDE_NETWORK_OPTIONS: {
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
					destination:    STATE_SELECTING_INSIDE_NETWORK_OPTIONS,
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
			STATE_SELECTING_PERMITTED_NETWORK_OPTIONS: {
				EVENT_CREATE_PERMITTED_NETWORK: {
					actionFunction: ShowCreatePermittedNetworkOptions,
					destination:    STATE_SELECTING_CREATE_PERMITTED_NETWORK,
				},
				EVENT_REMOVE_PERMITTED_NETWORK: {
					actionFunction: ShowRemovePermittedNetworkOptions,
					destination:    STATE_SELECTING_REMOVE_PERMITTED_NETWORK,
				},
				EVENT_UPDATE_PERMITTED_NETWORK: {
					actionFunction: ShowUpdatePermittedNetworkOptions,
					destination:    STATE_SELECTING_UPDATE_PERMITTED_NETWORK,
				},
				EVENT_GET_PERMITTED_NETWORK_BY_ID: {
					actionFunction: ShowGetPermittedNetworkByIdOptions,
					destination:    STATE_SELECTING_GET_PERMITTED_NETWORK,
				},
				EVENT_GET_ALL_PERMITTED_NETWORKS_BY_ADDRESS: {
					actionFunction: ShowGetAllPermittedNetworksByAddressOptions,
					destination:    STATE_SELECTING_GET_ALL_PERMITTED_NETWORKS_BY_ADDRESS,
				},
				EVENT_GET_PERMITTED_NETWORK_BY_INDEX_AND_ADDRESS: {
					actionFunction: ShowGetPermittedNetworkByIndexAndAddressOptions,
					destination:    STATE_SELECTING_GET_PERMITTED_NETWORK_BY_INDEX_AND_ADDRESS,
				},
				EVENT_PERMITTED_NETWORK_EXISTS: {
					actionFunction: ShowPermittedNetworkExistsOptions,
					destination:    STATE_SELECTING_PERMITTED_NETWORK_EXISTS,
				},
				EVENT_RETURN_TO_PREVIOUS_MENU: {
					actionFunction: showInsideDataOption,
					destination:    STATE_SELECTING_INSIDE_NETWORK_OPTIONS,
				},
			},
			STATE_SELECTING_CREATE_PERMITTED_NETWORK: {
				EVENT_RETURN_TO_PREVIOUS_MENU: {
					actionFunction: ShowInsideLedgerPermittedNetworkInfoOptions,
					destination:    STATE_SELECTING_PERMITTED_NETWORK_OPTIONS,
				},
			},
			STATE_SELECTING_REMOVE_PERMITTED_NETWORK: {
				EVENT_RETURN_TO_PREVIOUS_MENU: {
					actionFunction: ShowInsideLedgerPermittedNetworkInfoOptions,
					destination:    STATE_SELECTING_PERMITTED_NETWORK_OPTIONS,
				},
			},
			STATE_SELECTING_UPDATE_PERMITTED_NETWORK: {
				EVENT_RETURN_TO_PREVIOUS_MENU: {
					actionFunction: ShowInsideLedgerPermittedNetworkInfoOptions,
					destination:    STATE_SELECTING_PERMITTED_NETWORK_OPTIONS,
				},
			},
			STATE_SELECTING_GET_PERMITTED_NETWORK: {
				EVENT_RETURN_TO_PREVIOUS_MENU: {
					actionFunction: ShowInsideLedgerPermittedNetworkInfoOptions,
					destination:    STATE_SELECTING_PERMITTED_NETWORK_OPTIONS,
				},
			},
			STATE_SELECTING_GET_ALL_PERMITTED_NETWORKS_BY_ADDRESS: {
				EVENT_RETURN_TO_PREVIOUS_MENU: {
					actionFunction: ShowInsideLedgerPermittedNetworkInfoOptions,
					destination:    STATE_SELECTING_PERMITTED_NETWORK_OPTIONS,
				},
			},
			STATE_SELECTING_GET_PERMITTED_NETWORK_BY_INDEX_AND_ADDRESS: {
				EVENT_RETURN_TO_PREVIOUS_MENU: {
					actionFunction: ShowInsideLedgerPermittedNetworkInfoOptions,
					destination:    STATE_SELECTING_PERMITTED_NETWORK_OPTIONS,
				},
			},
			STATE_SELECTING_PERMITTED_NETWORK_EXISTS: {
				EVENT_RETURN_TO_PREVIOUS_MENU: {
					actionFunction: ShowInsideLedgerPermittedNetworkInfoOptions,
					destination:    STATE_SELECTING_PERMITTED_NETWORK_OPTIONS,
				},
			},
			//*********************** Permitted Methods STATES *************************
			STATE_SELECTING_PERMITTED_METHOD_OPTIONS: {
				EVENT_CREATE_PERMITTED_METHOD: {
					actionFunction: ShowCreatePermittedMethodOptions,
					destination:    STATE_SELECTING_CREATE_PERMITTED_METHOD,
				},
				EVENT_REMOVE_PERMITTED_METHOD: {
					actionFunction: ShowRemovePermittedMethodOptions,
					destination:    STATE_SELECTING_REMOVE_PERMITTED_METHOD,
				},
				EVENT_UPDATE_PERMITTED_METHOD: {
					actionFunction: ShowUpdatePermittedMethodOptions,
					destination:    STATE_SELECTING_UPDATE_PERMITTED_METHOD,
				},
				EVENT_GET_PERMITTED_METHOD_BY_ID: {
					actionFunction: ShowGetPermittedMethodByIdOptions,
					destination:    STATE_SELECTING_GET_PERMITTED_METHOD_BY_ID,
				},
				EVENT_GET_ALL_PERMITTED_METHODS_BY_ADDRESS: {
					actionFunction: ShowGetPermittedMethodsByNetworkIdOptions,
					destination:    STATE_SELECTING_GET_ALL_PERMITTED_METHODS_BY_ADDRESS,
				},
				EVENT_GET_ALL_PERMITTED_METHODS_BY_ID_AND_INDEX: {
					actionFunction: ShowGetPermittedMethodsByNetworkIdAndIndexOptions,
					destination:    STATE_SELECTING_GET_ALL_PERMITTED_METHODS_BY_ID_AND_INDEX,
				},
				EVENT_PERMITTED_METHOD_EXISTS: {
					actionFunction: ShowPermittedMethodExistsOptions,
					destination:    STATE_SELECTING_PERMITTED_METHOD_EXISTS,
				},
				EVENT_PERMITTED_METHOD_INVOKE: {
					actionFunction: ShowInvokePermittedMethodOptions,
					destination:    STATE_SELECTING_INVOKE_PERMITTED_METHOD_BY_ID,
				},
				EVENT_RETURN_TO_PREVIOUS_MENU: {
					actionFunction: showInsideDataOption,
					destination:    STATE_SELECTING_INSIDE_NETWORK_OPTIONS,
				},
			},
			STATE_SELECTING_CREATE_PERMITTED_METHOD: {
				EVENT_RETURN_TO_PREVIOUS_MENU: {
					actionFunction: ShowInsideLedgerPermittedMethodInfoOptions,
					destination:    STATE_SELECTING_PERMITTED_METHOD_OPTIONS,
				},
			},
			STATE_SELECTING_REMOVE_PERMITTED_METHOD: {
				EVENT_RETURN_TO_PREVIOUS_MENU: {
					actionFunction: ShowInsideLedgerPermittedMethodInfoOptions,
					destination:    STATE_SELECTING_PERMITTED_METHOD_OPTIONS,
				},
			},
			STATE_SELECTING_UPDATE_PERMITTED_METHOD: {
				EVENT_RETURN_TO_PREVIOUS_MENU: {
					actionFunction: ShowInsideLedgerPermittedMethodInfoOptions,
					destination:    STATE_SELECTING_PERMITTED_METHOD_OPTIONS,
				},
			},
			STATE_SELECTING_GET_PERMITTED_METHOD_BY_ID: {
				EVENT_RETURN_TO_PREVIOUS_MENU: {
					actionFunction: ShowInsideLedgerPermittedMethodInfoOptions,
					destination:    STATE_SELECTING_PERMITTED_METHOD_OPTIONS,
				},
			},
			STATE_SELECTING_GET_ALL_PERMITTED_METHODS_BY_ADDRESS: {
				EVENT_RETURN_TO_PREVIOUS_MENU: {
					actionFunction: ShowInsideLedgerPermittedMethodInfoOptions,
					destination:    STATE_SELECTING_PERMITTED_METHOD_OPTIONS,
				},
			},
			STATE_SELECTING_GET_ALL_PERMITTED_METHODS_BY_ID_AND_INDEX: {
				EVENT_RETURN_TO_PREVIOUS_MENU: {
					actionFunction: ShowInsideLedgerPermittedMethodInfoOptions,
					destination:    STATE_SELECTING_PERMITTED_METHOD_OPTIONS,
				},
			},
			STATE_SELECTING_PERMITTED_METHOD_EXISTS: {
				EVENT_RETURN_TO_PREVIOUS_MENU: {
					actionFunction: ShowInsideLedgerPermittedMethodInfoOptions,
					destination:    STATE_SELECTING_PERMITTED_METHOD_OPTIONS,
				},
			},
			STATE_SELECTING_INVOKE_PERMITTED_METHOD_BY_ID: {
				EVENT_RETURN_TO_PREVIOUS_MENU: {
					actionFunction: ShowInsideLedgerPermittedMethodInfoOptions,
					destination:    STATE_SELECTING_PERMITTED_METHOD_OPTIONS,
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
		err := fsm.doTransition(EVENT_SELECTING_OUTSIDE_DATA)
		if err != nil {
			fmt.Println(err)
		}
	case "2":
		err := fsm.doTransition(EVENT_SELECTING_INSIDE_DATA_OPTIONS)
		if err != nil {
			fmt.Println(err)
		}

	}

}

func showSelectAccessibleNetworksOption(fsm *StateMachine) {
	configuration, err := configs.ReadConfigYAMLFile()
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println("Please Select the Accessible Network You Want to Connect: ")
	fmt.Println("0. Return to Previous Menu")
	counter := 1
	for _, network := range configuration.Client {
		fmt.Println(strconv.Itoa(counter)+".Network Name: ", network.AccessibleNetworkName)
	}
	var index string
	fmt.Scanln(&index)
	if index == "0" {
		fsm.doTransition(EVENT_RETURN_TO_PREVIOUS_MENU)
		return
	}
	intIndex, err := strconv.Atoi(index)
	intIndex -= 1
	fsm.selectedAccessibleNetworkId = configuration.Client[intIndex].AccessibleNetworkId
	fsm.doTransition(EVENT_SELECTING_OUTSIDE_DATA_OPTIONS)

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
	time.Sleep(time.Second / 4)
	var fsm = new(StateMachine)
	fsm = fsm.FsmCreator()
	err := fsm.doTransition(EVENT_INITIAL)
	if err != nil {
		fmt.Println(err)
	}
}
