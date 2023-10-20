package hlf_mediator

import "testing"

func TestHlfAddPermittedNetwork(t *testing.T) {
	var hlfPermittedMethod HlfPermittedMethodsMediator
	submitRes, err := hlfPermittedMethod.AddPermittedMethod("PermittedNetwork3", "PermittedNetwork3", "PermittedMethod1", "mychannel", "add", "add")
	if err != nil {
		t.Errorf("Creating Accessible Netwrok: %v", err)
	}
	println(submitRes.PermittedMethodId)
}

func TestHlfRemovePermittedMethod(t *testing.T) {
	var hlfPermittedMethod HlfPermittedMethodsMediator
	err := hlfPermittedMethod.RemovePermittedMethod("permittedMethod_PermittedNetwork3_8")
	if err != nil {
		t.Errorf("Creating Accessible Netwrok: %v", err)
	}
}

func TestHlfUpdatePermittedMethod(t *testing.T) {
	var hlfPermittedMethod HlfPermittedMethodsMediator
	err := hlfPermittedMethod.UpdatePermittedMethod("permittedMethod_PermittedNetwork3_9", "PermittedNetwork2", "PermittedMethod1", "mychannel", "add", "add")
	if err != nil {
		t.Errorf("Creating Accessible Netwrok: %v", err)
	}
}

func TestHlfGetPermittedMethodsByNetworkId(t *testing.T) {
	var hlfPermittedMethod HlfPermittedMethodsMediator
	res, err := hlfPermittedMethod.GetPermittedMethodsByNetworkId("PermittedNetwork3")
	if err != nil {
		t.Errorf("Creating Accessible Netwrok: %v", err)
	}
	if len(res) == 0 {
		t.Errorf("Creating Accessible Netwrok: %v", res)
	}
}

func TestHlfGetPermittedMethod(t *testing.T) {
	var hlfPermittedMethod HlfPermittedMethodsMediator
	res, err := hlfPermittedMethod.GetPermittedMethod("permittedMethod_PermittedNetwork2_1")
	if err != nil {
		t.Errorf("Creating Accessible Netwrok: %v", err)
	}
	if res.PermittedMethodId != "permittedMethod_PermittedNetwork2_1" {
		t.Errorf("Creating Accessible Netwrok: %v", res)
	}
}

func TestHlfInvokePermittedMethod(t *testing.T) {
	var hlfPermittedMethod HlfPermittedMethodsMediator
	srtArray := []string{"9", "2"}
	res, err := hlfPermittedMethod.InvokePermittedMethod("AddTwoNumbers", "addition", "mychannel", srtArray)
	if err != nil {
		t.Errorf(" Invoke addition Method error : %v", err)
	}
	println(*res)
	if *res != "11" {
		t.Errorf(" Invoke addition Method error : %v", res)
	}

}

func TestHlfGetPermittedMethodsByIndex(t *testing.T) {
	var hlfPermittedMethod HlfPermittedMethodsMediator
	res, err := hlfPermittedMethod.GetPermittedMethodsByIndexAndAddress("PermittedNetwork2", "0", "10")
	if err != nil {
		t.Errorf("Creating Accessible Netwrok: %v", err)
	}
	if len(res) == 0 {
		t.Errorf("Creating Accessible Netwrok: %v", res)
	}
}

func TestHlfPermittedMethodExists(t *testing.T) {
	var hlfPermittedMethod HlfPermittedMethodsMediator
	res, err := hlfPermittedMethod.PermittedMethodExists("permittedMethod_PermittedNetwork2_1")
	if err != nil {
		t.Errorf("Creating Accessible Netwrok: %v", err)
	}
	if res != true {
		t.Errorf("Creating Accessible Netwrok: %v", res)
	}
	wrongRes, err := hlfPermittedMethod.PermittedMethodExists("permittedMethod_PermittedNetwork2_7888")
	if wrongRes != false {
		t.Errorf("Creating Accessible Netwrok: %v", wrongRes)
	}
}
