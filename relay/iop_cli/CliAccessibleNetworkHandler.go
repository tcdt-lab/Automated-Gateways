package iop_cli

import (
	"fmt"
	"github.com/tcdt-lab/Automated-Gateways/relay/mediator"
)

func ShowInsideLedgerAccessibleNetworkInfoOptions(fsm *StateMachine) {
	fmt.Println("Please Select One of the Following Options: ")
	fmt.Println("1. Create Accessible Network")
	fmt.Println("2. Update Accessible Network")
	fmt.Println("3. Remove Accessible Network")
	fmt.Println("4. Get Accessible Network Info By Id")
	fmt.Println("5. Get Accessible Network By Address")
	fmt.Println("6. Check If Accessible Network Exist")
	fmt.Println("7. Return To Previous Menu")

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
		err := fsm.doTransition(EVENT_ACCESSIBLE_NETWORK_EXISTS)
		if err != nil {
			fmt.Println(err)
		}

	case "7":
		err := fsm.doTransition(EVENT_RETURN_TO_PREVIOUS_MENU)
		if err != nil {
			fmt.Println(err)
		}
	}
}

func ShowCreateAccessibleNetworkOptions(fsm *StateMachine) {
	var networkName string
	var ip string
	var address string
	var companyName string
	var mediatorFactory mediator.MediatorFactory

	accessibleNetwrokMediator, err := mediatorFactory.CreateAccessibleNetworkMediator(mediator.HYPERLEDGER_FABRIC_NETWROK_TYPE)
	if err != nil {
		fmt.Println(err)
		err = fsm.doTransition(EVENT_RETURN_TO_PREVIOUS_MENU)
		return
	}

	fmt.Println("Please Insert the Required Information: ")
	fmt.Println("1. Netwrok Name")
	fmt.Scanln(&networkName)
	fmt.Println("2. IP")
	fmt.Scanln(&ip)
	fmt.Println("3. Address")
	fmt.Scanln(&address)
	fmt.Println("4. Company Name")
	fmt.Scanln(&companyName)

	res, err := accessibleNetwrokMediator.CreateAccessibleNetwork(networkName, ip, address, companyName)
	println("Result: ", res)
	if err != nil {
		fmt.Println(err)
		err = fsm.doTransition(EVENT_RETURN_TO_PREVIOUS_MENU)
		return
	}
	err = fsm.doTransition(EVENT_RETURN_TO_PREVIOUS_MENU)
	if err != nil {
		fmt.Println(err)
	}

}

func ShowUpdateAccessibleNetworkInfo(fsm *StateMachine) {
	var id string
	var networkName string
	var ip string
	var address string
	var companyName string
	var mediatorFactory mediator.MediatorFactory

	accessibleNetworkMediator, err := mediatorFactory.CreateAccessibleNetworkMediator(mediator.HYPERLEDGER_FABRIC_NETWROK_TYPE)
	if err != nil {
		fmt.Println(err)
		err = fsm.doTransition(EVENT_RETURN_TO_PREVIOUS_MENU)
		return
	}

	fmt.Println("Please Insert the Required Information: ")
	fmt.Println("1. Id")
	fmt.Scanln(&id)
	fmt.Println("2. Netwrok Name")
	fmt.Scanln(&networkName)
	fmt.Println("3. IP")
	fmt.Scanln(&ip)
	fmt.Println("4. Address")
	fmt.Scanln(&address)
	fmt.Println("5. Company Name")
	fmt.Scanln(&companyName)

	err = accessibleNetworkMediator.UpdateAccessibleNetwork(id, networkName, ip, address, companyName)
	if err != nil {
		fmt.Println(err)
		err = fsm.doTransition(EVENT_RETURN_TO_PREVIOUS_MENU)
		return
	}
	err = fsm.doTransition(EVENT_RETURN_TO_PREVIOUS_MENU)
	if err != nil {
		fmt.Println(err)
	}
	println("Result: ", "Updated Successfully")

}

func ShowRemoveAccessibleNetworkOptions(fsm *StateMachine) {
	var id string
	var mediatorFactory mediator.MediatorFactory
	accessibleNetworkMediator, err := mediatorFactory.CreateAccessibleNetworkMediator(mediator.HYPERLEDGER_FABRIC_NETWROK_TYPE)
	if err != nil {
		fmt.Println(err)
		err = fsm.doTransition(EVENT_RETURN_TO_PREVIOUS_MENU)
		return
	}

	fmt.Println("Please Insert the Required Information: ")
	fmt.Println("1. Id")
	fmt.Scanln(&id)

	err = accessibleNetworkMediator.RemoveAccessibleNetwork(id)
	if err != nil {
		fmt.Println(err)
		err = fsm.doTransition(EVENT_RETURN_TO_PREVIOUS_MENU)
		return
	}
	err = fsm.doTransition(EVENT_RETURN_TO_PREVIOUS_MENU)
	if err != nil {
		fmt.Println(err)
	}
	println("Result: ", "Removed Successfully")
}

func ShowGetAccessibleNetworkInfoByIdOptions(fsm *StateMachine) {
	var id string
	var mediatorFactory mediator.MediatorFactory
	accessibleNetworkMediator, err := mediatorFactory.CreateAccessibleNetworkMediator(mediator.HYPERLEDGER_FABRIC_NETWROK_TYPE)
	if err != nil {
		fmt.Println(err)
		err = fsm.doTransition(EVENT_RETURN_TO_PREVIOUS_MENU)
		return
	}

	fmt.Println("Please Insert the Required Information: ")
	fmt.Println("1. Id")
	fmt.Scanln(&id)

	res, err := accessibleNetworkMediator.GetAccessibleNetwork(id)
	if err != nil {
		fmt.Println(err)
		err = fsm.doTransition(EVENT_RETURN_TO_PREVIOUS_MENU)
		return
	}
	fmt.Println("Network Id: ", res.AccessibleNetworkId)
	fmt.Println("Network Name: ", res.NetworkName)
	fmt.Println("Network IP: ", res.IP)
	fmt.Println("Network Address: ", res.ADDRESS)
	fmt.Println("Company Name: ", res.CompanyName)
	fmt.Println("*******************************************************")

	err = fsm.doTransition(EVENT_RETURN_TO_PREVIOUS_MENU)
	if err != nil {
		fmt.Println(err)
	}
}

func ShowGetAllAccessibleNetworksByAddressOptions(fsm *StateMachine) {
	var address string
	var mediatorFactory mediator.MediatorFactory
	accessibleNetworkMediator, err := mediatorFactory.CreateAccessibleNetworkMediator(mediator.HYPERLEDGER_FABRIC_NETWROK_TYPE)
	if err != nil {
		fmt.Println(err)
		err = fsm.doTransition(EVENT_RETURN_TO_PREVIOUS_MENU)
		return
	}

	fmt.Println("Please Insert the Required Information: ")
	fmt.Println("1. Address")
	fmt.Scanln(&address)

	res, err := accessibleNetworkMediator.GetAllAccessibleNetworksByAddress(address)
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
	err = fsm.doTransition(EVENT_RETURN_TO_PREVIOUS_MENU)
	if err != nil {
		fmt.Println(err)
		err = fsm.doTransition(EVENT_RETURN_TO_PREVIOUS_MENU)
		return
	}
}

func ShowIfAccessibleNetworkExists(fsm *StateMachine) {
	var id string
	var mediatorFactory mediator.MediatorFactory
	accessibleNetworkMediator, err := mediatorFactory.CreateAccessibleNetworkMediator(mediator.HYPERLEDGER_FABRIC_NETWROK_TYPE)
	if err != nil {
		fmt.Println(err)
	}

	fmt.Println("Please Insert the Required Information: ")
	fmt.Println("1. Id")
	fmt.Scanln(&id)

	res, err := accessibleNetworkMediator.AccessibleNetworkExists(id)
	if err != nil {
		fmt.Println(err)
		err = fsm.doTransition(EVENT_RETURN_TO_PREVIOUS_MENU)
		return
	}
	fmt.Println("Result: ", res)
	err = fsm.doTransition(EVENT_RETURN_TO_PREVIOUS_MENU)
	if err != nil {
		fmt.Println(err)
		return
	}
}
