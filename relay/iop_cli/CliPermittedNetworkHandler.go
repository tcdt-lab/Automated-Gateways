package iop_cli

import (
	"fmt"
	mediator "github.com/tcdt-lab/Automated-Gateways/relay/mediator"
)

func ShowInsideLedgerPermittedNetworkInfoOptions(fsm *StateMachine) {
	fmt.Println("Please Select One of the Following Options: ")
	fmt.Println("1. Create Permitted Network")
	fmt.Println("2. Update Permitted Network")
	fmt.Println("3. Remove Permitted Network")
	fmt.Println("4. Get Permitted Network Info By Id")
	fmt.Println("5. Get Permitted Network By Address")
	fmt.Println("6. Get Permitted Network By Index and Address")
	fmt.Println("7. Check If Permitted Network Exist")
	fmt.Println("8. Return To Previous Menu")

	var answer string
	fmt.Scanln(&answer)
	switch answer {
	case "1":
		err := fsm.doTransition(EVENT_CREATE_PERMITTED_NETWORK)
		if err != nil {
			fmt.Println(err)
			return
		}
	case "2":
		err := fsm.doTransition(EVENT_UPDATE_PERMITTED_NETWORK)
		if err != nil {
			fmt.Println(err)
			return
		}
	case "3":
		err := fsm.doTransition(EVENT_REMOVE_PERMITTED_NETWORK)
		if err != nil {
			fmt.Println(err)
			return
		}
	case "4":
		err := fsm.doTransition(EVENT_GET_PERMITTED_NETWORK_BY_ID)
		if err != nil {
			fmt.Println(err)
			return
		}
	case "5":
		err := fsm.doTransition(EVENT_GET_ALL_PERMITTED_NETWORKS_BY_ADDRESS)
		if err != nil {
			fmt.Println(err)
			return
		}
	case "6":
		err := fsm.doTransition(EVENT_GET_PERMITTED_NETWORK_BY_INDEX_AND_ADDRESS)
		if err != nil {
			fmt.Println(err)
			return
		}
	case "7":
		err := fsm.doTransition(EVENT_PERMITTED_NETWORK_EXISTS)
		if err != nil {
			fmt.Println(err)
			return
		}
	case "8":
		err := fsm.doTransition(EVENT_RETURN_TO_PREVIOUS_MENU)
		if err != nil {
			fmt.Println(err)
			return
		}
	}
}

func ShowCreatePermittedNetworkOptions(fsm *StateMachine) {
	var networkName string
	var ip string
	var address string
	var companyName string
	var mediatorFactory mediator.MediatorFactory

	permittedNetwrokMediator, err := mediatorFactory.CreatePermittedNetworkMediator(mediator.HYPERLEDGER_FABRIC_NETWROK_TYPE)
	if err != nil {
		fmt.Println(err)
		err = fsm.doTransition(EVENT_RETURN_TO_PREVIOUS_MENU)
		return
	}

	fmt.Println("Please Enter Network Name: ")
	fmt.Scanln(&networkName)

	fmt.Println("Please Enter IP: ")
	fmt.Scanln(&ip)

	fmt.Println("Please Enter Address: ")
	fmt.Scanln(&address)

	fmt.Println("Please Enter Company Name: ")
	fmt.Scanln(&companyName)

	res, err := permittedNetwrokMediator.CreatePermittedNetwork(networkName, ip, address, companyName)

	println("Result: ", res)
	if err != nil {
		fmt.Println(err)
		return
	}
	err = fsm.doTransition(EVENT_RETURN_TO_PREVIOUS_MENU)
	if err != nil {
		fmt.Println(err)
	}
}

func ShowUpdatePermittedNetworkOptions(fsm *StateMachine) {
	var id string
	var networkName string
	var ip string
	var address string
	var companyName string
	var mediatorFactory mediator.MediatorFactory

	permittedNetwrokMediator, err := mediatorFactory.CreatePermittedNetworkMediator(mediator.HYPERLEDGER_FABRIC_NETWROK_TYPE)
	if err != nil {
		fmt.Println(err)
		err = fsm.doTransition(EVENT_RETURN_TO_PREVIOUS_MENU)
		return
	}

	fmt.Println("Please Enter Id: ")
	fmt.Scanln(&id)

	fmt.Println("Please Enter Network Name: ")
	fmt.Scanln(&networkName)

	fmt.Println("Please Enter IP: ")
	fmt.Scanln(&ip)

	fmt.Println("Please Enter Address: ")
	fmt.Scanln(&address)

	fmt.Println("Please Enter Company Name: ")
	fmt.Scanln(&companyName)

	err = permittedNetwrokMediator.UpdatePermittedNetwork(id, networkName, ip, address, companyName)
	if err != nil {
		fmt.Println(err)
		return
	}
	println("Result: ", "Permitted Network Updated Successfully")
	err = fsm.doTransition(EVENT_RETURN_TO_PREVIOUS_MENU)
	if err != nil {
		fmt.Println(err)
	}
}

func ShowRemovePermittedNetworkOptions(fsm *StateMachine) {
	var id string
	var mediatorFactory mediator.MediatorFactory

	permittedNetworkMediator, err := mediatorFactory.CreatePermittedNetworkMediator(mediator.HYPERLEDGER_FABRIC_NETWROK_TYPE)
	if err != nil {
		fmt.Println(err)
		err = fsm.doTransition(EVENT_RETURN_TO_PREVIOUS_MENU)
		return
	}

	fmt.Println("Please Enter Id: ")
	fmt.Scanln(&id)

	err = permittedNetworkMediator.RemovePermittedNetwork(id)
	if err != nil {
		fmt.Println(err)
		return
	}
	println("Result: ", "Permitted Network Removed Successfully")
	err = fsm.doTransition(EVENT_RETURN_TO_PREVIOUS_MENU)
	if err != nil {
		fmt.Println(err)
	}
}

func ShowGetPermittedNetworkByIdOptions(fsm *StateMachine) {
	var id string
	var mediatorFactory mediator.MediatorFactory

	permittedNetworkMediator, err := mediatorFactory.CreatePermittedNetworkMediator(mediator.HYPERLEDGER_FABRIC_NETWROK_TYPE)
	if err != nil {
		fmt.Println(err)
		err = fsm.doTransition(EVENT_RETURN_TO_PREVIOUS_MENU)
		return
	}

	fmt.Println("Please Enter Id: ")
	fmt.Scanln(&id)

	res, err := permittedNetworkMediator.GetPermittedNetwork(id)
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println(res)
	err = fsm.doTransition(EVENT_RETURN_TO_PREVIOUS_MENU)
	if err != nil {
		fmt.Println(err)
	}
}

func ShowGetPermittedNetworkByIndexAndAddressOptions(fsm *StateMachine) {
	var startIndex string
	var endIndex string
	var address string
	var mediatorFactory mediator.MediatorFactory

	permittedNetworkMediator, err := mediatorFactory.CreatePermittedNetworkMediator(mediator.HYPERLEDGER_FABRIC_NETWROK_TYPE)
	if err != nil {
		fmt.Println(err)
		err = fsm.doTransition(EVENT_RETURN_TO_PREVIOUS_MENU)
		return
	}

	fmt.Println("Please Enter Start Index: ")
	fmt.Scanln(&startIndex)

	fmt.Println("Please Enter End Index: ")
	fmt.Scanln(&endIndex)

	fmt.Println("Please Enter Address: ")
	fmt.Scanln(&address)

	res, err := permittedNetworkMediator.GetPermittedNetworkByIndexAndAddress(startIndex, endIndex, address)
	if err != nil {
		fmt.Println(err)
		return
	}
	for _, network := range res {

		fmt.Println("Network Id: ", network.PermittedNetworkId)
		fmt.Println("Network Name: ", network.NetworkName)
		fmt.Println("Network IP: ", network.IP)
		fmt.Println("Network Address: ", network.ADDRESS)
		fmt.Println("Company Name: ", network.CompanyName)
		fmt.Println("*******************************************************")
	}
	err = fsm.doTransition(EVENT_RETURN_TO_PREVIOUS_MENU)
	if err != nil {
		fmt.Println(err)
	}
}

func ShowGetAllPermittedNetworksByAddressOptions(fsm *StateMachine) {
	var address string
	var mediatorFactory mediator.MediatorFactory

	permittedNetworkMediator, err := mediatorFactory.CreatePermittedNetworkMediator(mediator.HYPERLEDGER_FABRIC_NETWROK_TYPE)
	if err != nil {
		fmt.Println(err)
		err = fsm.doTransition(EVENT_RETURN_TO_PREVIOUS_MENU)
		return
	}

	fmt.Println("Please Enter Address: ")
	fmt.Scanln(&address)

	res, err := permittedNetworkMediator.GetAllPermittedNetworksByAddress(address)
	if err != nil {
		fmt.Println(err)
		return
	}
	for _, network := range res {

		fmt.Println("Network Id: ", network.PermittedNetworkId)
		fmt.Println("Network Name: ", network.NetworkName)
		fmt.Println("Network IP: ", network.IP)
		fmt.Println("Network Address: ", network.ADDRESS)
		fmt.Println("Company Name: ", network.CompanyName)
		fmt.Println("*******************************************************")
	}
	err = fsm.doTransition(EVENT_RETURN_TO_PREVIOUS_MENU)
	if err != nil {
		fmt.Println(err)
	}
}

func ShowPermittedNetworkExistsOptions(fsm *StateMachine) {
	var id string
	var mediatorFactory mediator.MediatorFactory

	permittedNetworkMediator, err := mediatorFactory.CreatePermittedNetworkMediator(mediator.HYPERLEDGER_FABRIC_NETWROK_TYPE)
	if err != nil {
		fmt.Println(err)
		err = fsm.doTransition(EVENT_RETURN_TO_PREVIOUS_MENU)
		return
	}

	fmt.Println("Please Enter Id: ")
	fmt.Scanln(&id)

	res, err := permittedNetworkMediator.PermittedNetworkExists(id)
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println("Result: ", res)
	err = fsm.doTransition(EVENT_RETURN_TO_PREVIOUS_MENU)
	if err != nil {
		fmt.Println(err)
	}
}
