package permitted_method_handler

import "testing"

func TestOpenCloseConnection(t *testing.T) {
	var hlfPermittedMethod HlfPermittedMethod
	clientConn, gateway, err := hlfPermittedMethod.OpenConnection()
	if err != nil {
		t.Errorf("Error opening connection: %v", err)
	}

	errConnetion := hlfPermittedMethod.CloseConnection(clientConn, gateway)
	if errConnetion != nil {
		t.Errorf("Error closing connection: %v", errConnetion)
	}

}

func TestAddPermittedMethod(t *testing.T) {
	var hlfPermittedMethod HlfPermittedMethod
	clientConn, gateway, err := hlfPermittedMethod.OpenConnection()
	if err != nil {
		t.Errorf("Error opening connection: %v", err)
	}
	defer hlfPermittedMethod.CloseConnection(clientConn, gateway)

	submitRes, err := hlfPermittedMethod.AddPermittedMethod(gateway, "PermittedNetwork3", "PermittedNetwork3", "PermittedMethod1", "mychannel", "add", "add")
	if err != nil {
		t.Errorf("Creating Accessible Netwrok: %v", err)
	}
	println(submitRes.PermittedMethodId)
}

func TestPermittedMethodExists(t *testing.T) {
	var hlfPermittedMethod HlfPermittedMethod
	clientConn, gateway, err := hlfPermittedMethod.OpenConnection()
	if err != nil {
		t.Errorf("Error opening connection: %v", err)
	}
	defer hlfPermittedMethod.CloseConnection(clientConn, gateway)

	res, err := hlfPermittedMethod.PermittedMethodExists(gateway, "permittedMethod_PermittedNetwork2_1")
	if err != nil {
		t.Errorf(" Accessible Netwrok Exists : %v", err)
	}
	if res != true {
		t.Errorf(" Accessible Netwrok Exists : %v", res)
	}
	wrongRes, err := hlfPermittedMethod.PermittedMethodExists(gateway, "permittedMethod_PermittedNetwork2_16666")
	if wrongRes != false {
		t.Errorf(" Accessible Netwrok Exists does not work properly : %v", wrongRes)
	}

}

func TestGetPermittedMethod(t *testing.T) {
	var hlfPermittedMethod HlfPermittedMethod
	clientConn, gateway, err := hlfPermittedMethod.OpenConnection()
	if err != nil {
		t.Errorf("Error opening connection: %v", err)
	}
	defer hlfPermittedMethod.CloseConnection(clientConn, gateway)

	res, err := hlfPermittedMethod.GetPermittedMethod(gateway, "permittedMethod_PermittedNetwork2_1")
	if err != nil {
		t.Errorf(" Get PErmitted Netwrok error : %v", err)
	}
	if res.PermittedMethodId != "permittedMethod_PermittedNetwork2_1" {
		t.Errorf(" t PErmitted Netwrok error : %v", res)
	}

}

func TestGetPermittedMethodsByIndex(t *testing.T) {
	var hlfPermittedMethod HlfPermittedMethod
	clientConn, gateway, err := hlfPermittedMethod.OpenConnection()
	if err != nil {
		t.Errorf("Error opening connection: %v", err)
	}
	defer hlfPermittedMethod.CloseConnection(clientConn, gateway)

	res, err := hlfPermittedMethod.GetPermittedMethodsByIndex(gateway, "PermittedNetwork2", "2", "4")
	if err != nil {
		t.Errorf(" Get PErmitted Netwrok error : %v", err)
	}
	if res[1].PermittedMethodId != "permittedMethod_PermittedNetwork2_3" {
		t.Errorf(" Get PErmitted Netwrok error : %v", res)
	}

}

func TestRemovePermittedMethod(t *testing.T) {
	var hlfPermittedMethod HlfPermittedMethod
	clientConn, gateway, err := hlfPermittedMethod.OpenConnection()
	if err != nil {
		t.Errorf("Error opening connection: %v", err)
	}
	defer hlfPermittedMethod.CloseConnection(clientConn, gateway)

	err = hlfPermittedMethod.RemovePermittedMethod(gateway, "permittedMethod_PermittedNetwork2_4")
	if err != nil {
		t.Errorf(" Remove Permitted Netwrok Error : %v", err)
	}

}

func TestGetPermittedMethodsByNetworkId(t *testing.T) {
	var hlfPermittedMethod HlfPermittedMethod
	clientConn, gateway, err := hlfPermittedMethod.OpenConnection()
	if err != nil {
		t.Errorf("Error opening connection: %v", err)
	}
	defer hlfPermittedMethod.CloseConnection(clientConn, gateway)

	res, err := hlfPermittedMethod.GetPermittedMethodsByNetworkId(gateway, "PermittedNetwork2")
	if err != nil {
		t.Errorf(" Get PErmitted Netwrok error : %v", err)
	}
	if res[0].PermittedMethodId != "permittedMethod_PermittedNetwork2_1" {
		t.Errorf(" Get PErmitted Netwrok error : %v", res)
	}

}

func TestUpdatePermittedMethod(t *testing.T) {
	var hlfPermittedMethod HlfPermittedMethod
	clientConn, gateway, err := hlfPermittedMethod.OpenConnection()
	if err != nil {
		t.Errorf("Error opening connection: %v", err)
	}
	defer hlfPermittedMethod.CloseConnection(clientConn, gateway)

	err = hlfPermittedMethod.UpdatePermittedMethod(gateway, "permittedMethod_PermittedNetwork2_1", "PermittedNetwork2__updated", "PermittedMethod1", "mychannel", "add", "add")
	if err != nil {
		t.Errorf(" Update PErmitted Netwrok error : %v", err)
	}

}
