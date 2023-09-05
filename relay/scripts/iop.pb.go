// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.28.1
// 	protoc        v3.12.4
// source: scripts/iop.proto

package scripts

import (
	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
	reflect "reflect"
	sync "sync"
)

const (
	// Verify that this generated code is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(20 - protoimpl.MinVersion)
	// Verify that runtime/protoimpl is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(protoimpl.MaxVersion - 20)
)

type OutsiderNetworkLoginInfo struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	NetworkUsername string `protobuf:"bytes,1,opt,name=network_username,json=networkUsername,proto3" json:"network_username,omitempty"`
	NetworkId       string `protobuf:"bytes,2,opt,name=network_id,json=networkId,proto3" json:"network_id,omitempty"`
	NetworkPassword string `protobuf:"bytes,3,opt,name=network_password,json=networkPassword,proto3" json:"network_password,omitempty"`
	Network_IP      string `protobuf:"bytes,4,opt,name=network_IP,json=networkIP,proto3" json:"network_IP,omitempty"`
}

func (x *OutsiderNetworkLoginInfo) Reset() {
	*x = OutsiderNetworkLoginInfo{}
	if protoimpl.UnsafeEnabled {
		mi := &file_scripts_iop_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *OutsiderNetworkLoginInfo) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*OutsiderNetworkLoginInfo) ProtoMessage() {}

func (x *OutsiderNetworkLoginInfo) ProtoReflect() protoreflect.Message {
	mi := &file_scripts_iop_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use OutsiderNetworkLoginInfo.ProtoReflect.Descriptor instead.
func (*OutsiderNetworkLoginInfo) Descriptor() ([]byte, []int) {
	return file_scripts_iop_proto_rawDescGZIP(), []int{0}
}

func (x *OutsiderNetworkLoginInfo) GetNetworkUsername() string {
	if x != nil {
		return x.NetworkUsername
	}
	return ""
}

func (x *OutsiderNetworkLoginInfo) GetNetworkId() string {
	if x != nil {
		return x.NetworkId
	}
	return ""
}

func (x *OutsiderNetworkLoginInfo) GetNetworkPassword() string {
	if x != nil {
		return x.NetworkPassword
	}
	return ""
}

func (x *OutsiderNetworkLoginInfo) GetNetwork_IP() string {
	if x != nil {
		return x.Network_IP
	}
	return ""
}

type OutsiderNetwork struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	NetworkId    string `protobuf:"bytes,1,opt,name=network_id,json=networkId,proto3" json:"network_id,omitempty"`
	NetworkToken string `protobuf:"bytes,2,opt,name=network_token,json=networkToken,proto3" json:"network_token,omitempty"`
}

func (x *OutsiderNetwork) Reset() {
	*x = OutsiderNetwork{}
	if protoimpl.UnsafeEnabled {
		mi := &file_scripts_iop_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *OutsiderNetwork) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*OutsiderNetwork) ProtoMessage() {}

func (x *OutsiderNetwork) ProtoReflect() protoreflect.Message {
	mi := &file_scripts_iop_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use OutsiderNetwork.ProtoReflect.Descriptor instead.
func (*OutsiderNetwork) Descriptor() ([]byte, []int) {
	return file_scripts_iop_proto_rawDescGZIP(), []int{1}
}

func (x *OutsiderNetwork) GetNetworkId() string {
	if x != nil {
		return x.NetworkId
	}
	return ""
}

func (x *OutsiderNetwork) GetNetworkToken() string {
	if x != nil {
		return x.NetworkToken
	}
	return ""
}

type Response struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Status      string `protobuf:"bytes,1,opt,name=status,proto3" json:"status,omitempty"`
	ResponseStr string `protobuf:"bytes,2,opt,name=response_str,json=responseStr,proto3" json:"response_str,omitempty"`
}

func (x *Response) Reset() {
	*x = Response{}
	if protoimpl.UnsafeEnabled {
		mi := &file_scripts_iop_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Response) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Response) ProtoMessage() {}

func (x *Response) ProtoReflect() protoreflect.Message {
	mi := &file_scripts_iop_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Response.ProtoReflect.Descriptor instead.
func (*Response) Descriptor() ([]byte, []int) {
	return file_scripts_iop_proto_rawDescGZIP(), []int{2}
}

func (x *Response) GetStatus() string {
	if x != nil {
		return x.Status
	}
	return ""
}

func (x *Response) GetResponseStr() string {
	if x != nil {
		return x.ResponseStr
	}
	return ""
}

type MethodInfo struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// Exposed Method Info
	MethodName    string `protobuf:"bytes,1,opt,name=method_name,json=methodName,proto3" json:"method_name,omitempty"`
	MethodType    string `protobuf:"bytes,2,opt,name=method_type,json=methodType,proto3" json:"method_type,omitempty"`
	MethodInput   string `protobuf:"bytes,3,opt,name=method_input,json=methodInput,proto3" json:"method_input,omitempty"`
	MethodOutput  string `protobuf:"bytes,4,opt,name=method_output,json=methodOutput,proto3" json:"method_output,omitempty"`
	Method_Id     string `protobuf:"bytes,5,opt,name=method_Id,json=methodId,proto3" json:"method_Id,omitempty"`
	ChannelName   string `protobuf:"bytes,6,opt,name=channelName,proto3" json:"channelName,omitempty"`
	ChaincodeName string `protobuf:"bytes,7,opt,name=chaincodeName,proto3" json:"chaincodeName,omitempty"`
	// User Info
	Network *OutsiderNetwork `protobuf:"bytes,8,opt,name=network,proto3" json:"network,omitempty"`
}

func (x *MethodInfo) Reset() {
	*x = MethodInfo{}
	if protoimpl.UnsafeEnabled {
		mi := &file_scripts_iop_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *MethodInfo) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*MethodInfo) ProtoMessage() {}

func (x *MethodInfo) ProtoReflect() protoreflect.Message {
	mi := &file_scripts_iop_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use MethodInfo.ProtoReflect.Descriptor instead.
func (*MethodInfo) Descriptor() ([]byte, []int) {
	return file_scripts_iop_proto_rawDescGZIP(), []int{3}
}

func (x *MethodInfo) GetMethodName() string {
	if x != nil {
		return x.MethodName
	}
	return ""
}

func (x *MethodInfo) GetMethodType() string {
	if x != nil {
		return x.MethodType
	}
	return ""
}

func (x *MethodInfo) GetMethodInput() string {
	if x != nil {
		return x.MethodInput
	}
	return ""
}

func (x *MethodInfo) GetMethodOutput() string {
	if x != nil {
		return x.MethodOutput
	}
	return ""
}

func (x *MethodInfo) GetMethod_Id() string {
	if x != nil {
		return x.Method_Id
	}
	return ""
}

func (x *MethodInfo) GetChannelName() string {
	if x != nil {
		return x.ChannelName
	}
	return ""
}

func (x *MethodInfo) GetChaincodeName() string {
	if x != nil {
		return x.ChaincodeName
	}
	return ""
}

func (x *MethodInfo) GetNetwork() *OutsiderNetwork {
	if x != nil {
		return x.Network
	}
	return nil
}

var File_scripts_iop_proto protoreflect.FileDescriptor

var file_scripts_iop_proto_rawDesc = []byte{
	0x0a, 0x11, 0x73, 0x63, 0x72, 0x69, 0x70, 0x74, 0x73, 0x2f, 0x69, 0x6f, 0x70, 0x2e, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x22, 0xae, 0x01, 0x0a, 0x18, 0x4f, 0x75, 0x74, 0x73, 0x69, 0x64, 0x65, 0x72,
	0x4e, 0x65, 0x74, 0x77, 0x6f, 0x72, 0x6b, 0x4c, 0x6f, 0x67, 0x69, 0x6e, 0x49, 0x6e, 0x66, 0x6f,
	0x12, 0x29, 0x0a, 0x10, 0x6e, 0x65, 0x74, 0x77, 0x6f, 0x72, 0x6b, 0x5f, 0x75, 0x73, 0x65, 0x72,
	0x6e, 0x61, 0x6d, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0f, 0x6e, 0x65, 0x74, 0x77,
	0x6f, 0x72, 0x6b, 0x55, 0x73, 0x65, 0x72, 0x6e, 0x61, 0x6d, 0x65, 0x12, 0x1d, 0x0a, 0x0a, 0x6e,
	0x65, 0x74, 0x77, 0x6f, 0x72, 0x6b, 0x5f, 0x69, 0x64, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52,
	0x09, 0x6e, 0x65, 0x74, 0x77, 0x6f, 0x72, 0x6b, 0x49, 0x64, 0x12, 0x29, 0x0a, 0x10, 0x6e, 0x65,
	0x74, 0x77, 0x6f, 0x72, 0x6b, 0x5f, 0x70, 0x61, 0x73, 0x73, 0x77, 0x6f, 0x72, 0x64, 0x18, 0x03,
	0x20, 0x01, 0x28, 0x09, 0x52, 0x0f, 0x6e, 0x65, 0x74, 0x77, 0x6f, 0x72, 0x6b, 0x50, 0x61, 0x73,
	0x73, 0x77, 0x6f, 0x72, 0x64, 0x12, 0x1d, 0x0a, 0x0a, 0x6e, 0x65, 0x74, 0x77, 0x6f, 0x72, 0x6b,
	0x5f, 0x49, 0x50, 0x18, 0x04, 0x20, 0x01, 0x28, 0x09, 0x52, 0x09, 0x6e, 0x65, 0x74, 0x77, 0x6f,
	0x72, 0x6b, 0x49, 0x50, 0x22, 0x55, 0x0a, 0x0f, 0x4f, 0x75, 0x74, 0x73, 0x69, 0x64, 0x65, 0x72,
	0x4e, 0x65, 0x74, 0x77, 0x6f, 0x72, 0x6b, 0x12, 0x1d, 0x0a, 0x0a, 0x6e, 0x65, 0x74, 0x77, 0x6f,
	0x72, 0x6b, 0x5f, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x09, 0x6e, 0x65, 0x74,
	0x77, 0x6f, 0x72, 0x6b, 0x49, 0x64, 0x12, 0x23, 0x0a, 0x0d, 0x6e, 0x65, 0x74, 0x77, 0x6f, 0x72,
	0x6b, 0x5f, 0x74, 0x6f, 0x6b, 0x65, 0x6e, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0c, 0x6e,
	0x65, 0x74, 0x77, 0x6f, 0x72, 0x6b, 0x54, 0x6f, 0x6b, 0x65, 0x6e, 0x22, 0x45, 0x0a, 0x08, 0x72,
	0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x16, 0x0a, 0x06, 0x73, 0x74, 0x61, 0x74, 0x75,
	0x73, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x06, 0x73, 0x74, 0x61, 0x74, 0x75, 0x73, 0x12,
	0x21, 0x0a, 0x0c, 0x72, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x5f, 0x73, 0x74, 0x72, 0x18,
	0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0b, 0x72, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x53,
	0x74, 0x72, 0x22, 0xa7, 0x02, 0x0a, 0x0a, 0x4d, 0x65, 0x74, 0x68, 0x6f, 0x64, 0x49, 0x6e, 0x66,
	0x6f, 0x12, 0x1f, 0x0a, 0x0b, 0x6d, 0x65, 0x74, 0x68, 0x6f, 0x64, 0x5f, 0x6e, 0x61, 0x6d, 0x65,
	0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0a, 0x6d, 0x65, 0x74, 0x68, 0x6f, 0x64, 0x4e, 0x61,
	0x6d, 0x65, 0x12, 0x1f, 0x0a, 0x0b, 0x6d, 0x65, 0x74, 0x68, 0x6f, 0x64, 0x5f, 0x74, 0x79, 0x70,
	0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0a, 0x6d, 0x65, 0x74, 0x68, 0x6f, 0x64, 0x54,
	0x79, 0x70, 0x65, 0x12, 0x21, 0x0a, 0x0c, 0x6d, 0x65, 0x74, 0x68, 0x6f, 0x64, 0x5f, 0x69, 0x6e,
	0x70, 0x75, 0x74, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0b, 0x6d, 0x65, 0x74, 0x68, 0x6f,
	0x64, 0x49, 0x6e, 0x70, 0x75, 0x74, 0x12, 0x23, 0x0a, 0x0d, 0x6d, 0x65, 0x74, 0x68, 0x6f, 0x64,
	0x5f, 0x6f, 0x75, 0x74, 0x70, 0x75, 0x74, 0x18, 0x04, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0c, 0x6d,
	0x65, 0x74, 0x68, 0x6f, 0x64, 0x4f, 0x75, 0x74, 0x70, 0x75, 0x74, 0x12, 0x1b, 0x0a, 0x09, 0x6d,
	0x65, 0x74, 0x68, 0x6f, 0x64, 0x5f, 0x49, 0x64, 0x18, 0x05, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08,
	0x6d, 0x65, 0x74, 0x68, 0x6f, 0x64, 0x49, 0x64, 0x12, 0x20, 0x0a, 0x0b, 0x63, 0x68, 0x61, 0x6e,
	0x6e, 0x65, 0x6c, 0x4e, 0x61, 0x6d, 0x65, 0x18, 0x06, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0b, 0x63,
	0x68, 0x61, 0x6e, 0x6e, 0x65, 0x6c, 0x4e, 0x61, 0x6d, 0x65, 0x12, 0x24, 0x0a, 0x0d, 0x63, 0x68,
	0x61, 0x69, 0x6e, 0x63, 0x6f, 0x64, 0x65, 0x4e, 0x61, 0x6d, 0x65, 0x18, 0x07, 0x20, 0x01, 0x28,
	0x09, 0x52, 0x0d, 0x63, 0x68, 0x61, 0x69, 0x6e, 0x63, 0x6f, 0x64, 0x65, 0x4e, 0x61, 0x6d, 0x65,
	0x12, 0x2a, 0x0a, 0x07, 0x6e, 0x65, 0x74, 0x77, 0x6f, 0x72, 0x6b, 0x18, 0x08, 0x20, 0x01, 0x28,
	0x0b, 0x32, 0x10, 0x2e, 0x4f, 0x75, 0x74, 0x73, 0x69, 0x64, 0x65, 0x72, 0x4e, 0x65, 0x74, 0x77,
	0x6f, 0x72, 0x6b, 0x52, 0x07, 0x6e, 0x65, 0x74, 0x77, 0x6f, 0x72, 0x6b, 0x32, 0x8e, 0x01, 0x0a,
	0x03, 0x49, 0x4f, 0x50, 0x12, 0x34, 0x0a, 0x05, 0x6c, 0x6f, 0x67, 0x69, 0x6e, 0x12, 0x19, 0x2e,
	0x4f, 0x75, 0x74, 0x73, 0x69, 0x64, 0x65, 0x72, 0x4e, 0x65, 0x74, 0x77, 0x6f, 0x72, 0x6b, 0x4c,
	0x6f, 0x67, 0x69, 0x6e, 0x49, 0x6e, 0x66, 0x6f, 0x1a, 0x10, 0x2e, 0x4f, 0x75, 0x74, 0x73, 0x69,
	0x64, 0x65, 0x72, 0x4e, 0x65, 0x74, 0x77, 0x6f, 0x72, 0x6b, 0x12, 0x20, 0x0a, 0x06, 0x69, 0x6e,
	0x76, 0x6f, 0x6b, 0x65, 0x12, 0x0b, 0x2e, 0x4d, 0x65, 0x74, 0x68, 0x6f, 0x64, 0x49, 0x6e, 0x66,
	0x6f, 0x1a, 0x09, 0x2e, 0x72, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x2f, 0x0a, 0x0c,
	0x71, 0x75, 0x65, 0x72, 0x79, 0x4d, 0x65, 0x74, 0x68, 0x6f, 0x64, 0x73, 0x12, 0x10, 0x2e, 0x4f,
	0x75, 0x74, 0x73, 0x69, 0x64, 0x65, 0x72, 0x4e, 0x65, 0x74, 0x77, 0x6f, 0x72, 0x6b, 0x1a, 0x0b,
	0x2e, 0x4d, 0x65, 0x74, 0x68, 0x6f, 0x64, 0x49, 0x6e, 0x66, 0x6f, 0x30, 0x01, 0x42, 0x0f, 0x5a,
	0x0d, 0x72, 0x65, 0x6c, 0x61, 0x79, 0x2f, 0x73, 0x63, 0x72, 0x69, 0x70, 0x74, 0x73, 0x62, 0x06,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_scripts_iop_proto_rawDescOnce sync.Once
	file_scripts_iop_proto_rawDescData = file_scripts_iop_proto_rawDesc
)

func file_scripts_iop_proto_rawDescGZIP() []byte {
	file_scripts_iop_proto_rawDescOnce.Do(func() {
		file_scripts_iop_proto_rawDescData = protoimpl.X.CompressGZIP(file_scripts_iop_proto_rawDescData)
	})
	return file_scripts_iop_proto_rawDescData
}

var file_scripts_iop_proto_msgTypes = make([]protoimpl.MessageInfo, 4)
var file_scripts_iop_proto_goTypes = []interface{}{
	(*OutsiderNetworkLoginInfo)(nil), // 0: OutsiderNetworkLoginInfo
	(*OutsiderNetwork)(nil),          // 1: OutsiderNetwork
	(*Response)(nil),                 // 2: response
	(*MethodInfo)(nil),               // 3: MethodInfo
}
var file_scripts_iop_proto_depIdxs = []int32{
	1, // 0: MethodInfo.network:type_name -> OutsiderNetwork
	0, // 1: IOP.login:input_type -> OutsiderNetworkLoginInfo
	3, // 2: IOP.invoke:input_type -> MethodInfo
	1, // 3: IOP.queryMethods:input_type -> OutsiderNetwork
	1, // 4: IOP.login:output_type -> OutsiderNetwork
	2, // 5: IOP.invoke:output_type -> response
	3, // 6: IOP.queryMethods:output_type -> MethodInfo
	4, // [4:7] is the sub-list for method output_type
	1, // [1:4] is the sub-list for method input_type
	1, // [1:1] is the sub-list for extension type_name
	1, // [1:1] is the sub-list for extension extendee
	0, // [0:1] is the sub-list for field type_name
}

func init() { file_scripts_iop_proto_init() }
func file_scripts_iop_proto_init() {
	if File_scripts_iop_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_scripts_iop_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*OutsiderNetworkLoginInfo); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_scripts_iop_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*OutsiderNetwork); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_scripts_iop_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Response); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_scripts_iop_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*MethodInfo); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
	}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_scripts_iop_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   4,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_scripts_iop_proto_goTypes,
		DependencyIndexes: file_scripts_iop_proto_depIdxs,
		MessageInfos:      file_scripts_iop_proto_msgTypes,
	}.Build()
	File_scripts_iop_proto = out.File
	file_scripts_iop_proto_rawDesc = nil
	file_scripts_iop_proto_goTypes = nil
	file_scripts_iop_proto_depIdxs = nil
}
