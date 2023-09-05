package onchain

import (
	"errors"
	"fmt"
	"relay/scripts"
)

func InvokeMethod(methodInfo *scripts.MethodInfo) (*scripts.Response, error) {
	fmt.Println("The method %s from chaincode %s and channel %s has been invoked by Network Id :%s", methodInfo.GetMethodName(), methodInfo.GetChaincodeName(), methodInfo.GetChannelName(), methodInfo.GetNetwork().GetNetworkId())
	var mockResult = scripts.Response{Status: "200", ResponseStr: "The method has been invoked"}
	return &mockResult, nil
}
func CheckNetworkInfo(info *scripts.OutsiderNetworkLoginInfo) (*scripts.OutsiderNetwork, error) {
	fmt.Println("The network %s credentialCheck", info.GetNetworkId())
	if info.NetworkId == "mockID" && info.Network_IP == "mockIP" && info.NetworkPassword == "mockPassword" && info.NetworkUsername == "mockUsername" {
		fmt.Println("The network %s has been logged in", info.GetNetworkId())
		return &scripts.OutsiderNetwork{NetworkId: "mockID", NetworkToken: "mockToken"}, nil
	}

	return nil, errors.New(fmt.Sprintln("The network %s credentialCheck failed", info.GetNetworkId()))

}

func QueryMethods(network *scripts.OutsiderNetwork) ([]*scripts.MethodInfo, error) {
	fmt.Println("The methods from network %s has been queried", network.GetNetworkId())
	var mockNetwork = scripts.OutsiderNetwork{NetworkId: "mockID", NetworkToken: "mockToken"}
	var mockResult = scripts.MethodInfo{MethodName: "mockMethod", ChaincodeName: "mockChaincode", ChannelName: "mockChannel", Method_Id: "mockID", MethodInput: "mockInput", MethodOutput: "mockOutput", MethodType: "mockType", Network: &mockNetwork}
	var mockResult2 = scripts.MethodInfo{MethodName: "mockMethod2", ChaincodeName: "mockChaincode2", ChannelName: "mockChannel2", Method_Id: "mockID2", MethodInput: "mockInput", MethodOutput: "mockOutput", MethodType: "mockType", Network: &mockNetwork}
	var mockResult3 = scripts.MethodInfo{MethodName: "mockMethod3", ChaincodeName: "mockChaincode3", ChannelName: "mockChannel3", Method_Id: "mockID3", MethodInput: "mockInput", MethodOutput: "mockOutput", MethodType: "mockType", Network: &mockNetwork}
	mockList := []*scripts.MethodInfo{&mockResult, &mockResult2, &mockResult3}
	return mockList, nil
}
