package iop_cli

import "fmt"

func ShowInsideLedgerAccessibleNetworkInfoOptions(fsm *StateMachine) {
	fmt.Println("Please Select One of the Following Options: ")
	fmt.Println("1. Create Accessible Network")
	fmt.Println("2. Update Accessible Network")
	fmt.Println("3. Remove Accessible Network")
	fmt.Println("4. Get Accessible Network Info By Id")
	fmt.Println("5. Get Accessible Network By Address")
	fmt.Println("6. Return To Previous Menu")

	var answer string
	fmt.Scanln(&answer)
	switch answer {
	case "1":
		err := fsm.doTransition(EVENT_CREATE_ACCESSIBLE_NETWORK)
		if err != nil {
			fmt.Println(err)
		}
	case "2":
		err := fsm.doTransition(EVENT_UPDATE_ACCESSIBLE_NETWORK)
		if err != nil {
			fmt.Println(err)
		}
	case "3":
		err := fsm.doTransition(EVENT_REMOVE_ACCESSIBLE_NETWORK)
		if err != nil {
			fmt.Println(err)
		}
	case "4":
		err := fsm.doTransition(EVENT_GET_ACCESSIBLE_NETWORK_BY_ID)
		if err != nil {
			fmt.Println(err)
		}
	case "5":
		err := fsm.doTransition(EVENT_GET_ALL_ACCESSIBLE_NETWORKS_BY_ADDRESS)
		if err != nil {
			fmt.Println(err)
		}

	case "6":
		err := fsm.doTransition(EVENT_RETURN_TO_PREVIOUS_MENU)
		if err != nil {
			fmt.Println(err)
		}
	}
}
