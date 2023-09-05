// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.28.1
// 	protoc        v3.12.4
// source: iop.proto

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
		mi := &file_iop_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *OutsiderNetwork) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*OutsiderNetwork) ProtoMessage() {}

func (x *OutsiderNetwork) ProtoReflect() protoreflect.Message {
	mi := &file_iop_proto_msgTypes[0]
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
	return file_iop_proto_rawDescGZIP(), []int{0}
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
		mi := &file_iop_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Response) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Response) ProtoMessage() {}

func (x *Response) ProtoReflect() protoreflect.Message {
	mi := &file_iop_proto_msgTypes[1]
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
	return file_iop_proto_rawDescGZIP(), []int{1}
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

type IOPInfo struct {
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
	UserId    string `protobuf:"bytes,8,opt,name=user_id,json=userId,proto3" json:"user_id,omitempty"`
	UserToken string `protobuf:"bytes,9,opt,name=user_token,json=userToken,proto3" json:"user_token,omitempty"`
}

func (x *IOPInfo) Reset() {
	*x = IOPInfo{}
	if protoimpl.UnsafeEnabled {
		mi := &file_iop_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *IOPInfo) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*IOPInfo) ProtoMessage() {}

func (x *IOPInfo) ProtoReflect() protoreflect.Message {
	mi := &file_iop_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use IOPInfo.ProtoReflect.Descriptor instead.
func (*IOPInfo) Descriptor() ([]byte, []int) {
	return file_iop_proto_rawDescGZIP(), []int{2}
}

func (x *IOPInfo) GetMethodName() string {
	if x != nil {
		return x.MethodName
	}
	return ""
}

func (x *IOPInfo) GetMethodType() string {
	if x != nil {
		return x.MethodType
	}
	return ""
}

func (x *IOPInfo) GetMethodInput() string {
	if x != nil {
		return x.MethodInput
	}
	return ""
}

func (x *IOPInfo) GetMethodOutput() string {
	if x != nil {
		return x.MethodOutput
	}
	return ""
}

func (x *IOPInfo) GetMethod_Id() string {
	if x != nil {
		return x.Method_Id
	}
	return ""
}

func (x *IOPInfo) GetChannelName() string {
	if x != nil {
		return x.ChannelName
	}
	return ""
}

func (x *IOPInfo) GetChaincodeName() string {
	if x != nil {
		return x.ChaincodeName
	}
	return ""
}

func (x *IOPInfo) GetUserId() string {
	if x != nil {
		return x.UserId
	}
	return ""
}

func (x *IOPInfo) GetUserToken() string {
	if x != nil {
		return x.UserToken
	}
	return ""
}

var File_iop_proto protoreflect.FileDescriptor

var file_iop_proto_rawDesc = []byte{
	0x0a, 0x09, 0x69, 0x6f, 0x70, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0x55, 0x0a, 0x0f, 0x4f,
	0x75, 0x74, 0x73, 0x69, 0x64, 0x65, 0x72, 0x4e, 0x65, 0x74, 0x77, 0x6f, 0x72, 0x6b, 0x12, 0x1d,
	0x0a, 0x0a, 0x6e, 0x65, 0x74, 0x77, 0x6f, 0x72, 0x6b, 0x5f, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01,
	0x28, 0x09, 0x52, 0x09, 0x6e, 0x65, 0x74, 0x77, 0x6f, 0x72, 0x6b, 0x49, 0x64, 0x12, 0x23, 0x0a,
	0x0d, 0x6e, 0x65, 0x74, 0x77, 0x6f, 0x72, 0x6b, 0x5f, 0x74, 0x6f, 0x6b, 0x65, 0x6e, 0x18, 0x02,
	0x20, 0x01, 0x28, 0x09, 0x52, 0x0c, 0x6e, 0x65, 0x74, 0x77, 0x6f, 0x72, 0x6b, 0x54, 0x6f, 0x6b,
	0x65, 0x6e, 0x22, 0x45, 0x0a, 0x08, 0x72, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x16,
	0x0a, 0x06, 0x73, 0x74, 0x61, 0x74, 0x75, 0x73, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x06,
	0x73, 0x74, 0x61, 0x74, 0x75, 0x73, 0x12, 0x21, 0x0a, 0x0c, 0x72, 0x65, 0x73, 0x70, 0x6f, 0x6e,
	0x73, 0x65, 0x5f, 0x73, 0x74, 0x72, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0b, 0x72, 0x65,
	0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x53, 0x74, 0x72, 0x22, 0xb0, 0x02, 0x0a, 0x07, 0x49, 0x4f,
	0x50, 0x49, 0x6e, 0x66, 0x6f, 0x12, 0x1f, 0x0a, 0x0b, 0x6d, 0x65, 0x74, 0x68, 0x6f, 0x64, 0x5f,
	0x6e, 0x61, 0x6d, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0a, 0x6d, 0x65, 0x74, 0x68,
	0x6f, 0x64, 0x4e, 0x61, 0x6d, 0x65, 0x12, 0x1f, 0x0a, 0x0b, 0x6d, 0x65, 0x74, 0x68, 0x6f, 0x64,
	0x5f, 0x74, 0x79, 0x70, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0a, 0x6d, 0x65, 0x74,
	0x68, 0x6f, 0x64, 0x54, 0x79, 0x70, 0x65, 0x12, 0x21, 0x0a, 0x0c, 0x6d, 0x65, 0x74, 0x68, 0x6f,
	0x64, 0x5f, 0x69, 0x6e, 0x70, 0x75, 0x74, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0b, 0x6d,
	0x65, 0x74, 0x68, 0x6f, 0x64, 0x49, 0x6e, 0x70, 0x75, 0x74, 0x12, 0x23, 0x0a, 0x0d, 0x6d, 0x65,
	0x74, 0x68, 0x6f, 0x64, 0x5f, 0x6f, 0x75, 0x74, 0x70, 0x75, 0x74, 0x18, 0x04, 0x20, 0x01, 0x28,
	0x09, 0x52, 0x0c, 0x6d, 0x65, 0x74, 0x68, 0x6f, 0x64, 0x4f, 0x75, 0x74, 0x70, 0x75, 0x74, 0x12,
	0x1b, 0x0a, 0x09, 0x6d, 0x65, 0x74, 0x68, 0x6f, 0x64, 0x5f, 0x49, 0x64, 0x18, 0x05, 0x20, 0x01,
	0x28, 0x09, 0x52, 0x08, 0x6d, 0x65, 0x74, 0x68, 0x6f, 0x64, 0x49, 0x64, 0x12, 0x20, 0x0a, 0x0b,
	0x63, 0x68, 0x61, 0x6e, 0x6e, 0x65, 0x6c, 0x4e, 0x61, 0x6d, 0x65, 0x18, 0x06, 0x20, 0x01, 0x28,
	0x09, 0x52, 0x0b, 0x63, 0x68, 0x61, 0x6e, 0x6e, 0x65, 0x6c, 0x4e, 0x61, 0x6d, 0x65, 0x12, 0x24,
	0x0a, 0x0d, 0x63, 0x68, 0x61, 0x69, 0x6e, 0x63, 0x6f, 0x64, 0x65, 0x4e, 0x61, 0x6d, 0x65, 0x18,
	0x07, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0d, 0x63, 0x68, 0x61, 0x69, 0x6e, 0x63, 0x6f, 0x64, 0x65,
	0x4e, 0x61, 0x6d, 0x65, 0x12, 0x17, 0x0a, 0x07, 0x75, 0x73, 0x65, 0x72, 0x5f, 0x69, 0x64, 0x18,
	0x08, 0x20, 0x01, 0x28, 0x09, 0x52, 0x06, 0x75, 0x73, 0x65, 0x72, 0x49, 0x64, 0x12, 0x1d, 0x0a,
	0x0a, 0x75, 0x73, 0x65, 0x72, 0x5f, 0x74, 0x6f, 0x6b, 0x65, 0x6e, 0x18, 0x09, 0x20, 0x01, 0x28,
	0x09, 0x52, 0x09, 0x75, 0x73, 0x65, 0x72, 0x54, 0x6f, 0x6b, 0x65, 0x6e, 0x32, 0x52, 0x0a, 0x03,
	0x49, 0x4f, 0x50, 0x12, 0x1d, 0x0a, 0x06, 0x69, 0x6e, 0x76, 0x6f, 0x6b, 0x65, 0x12, 0x08, 0x2e,
	0x49, 0x4f, 0x50, 0x49, 0x6e, 0x66, 0x6f, 0x1a, 0x09, 0x2e, 0x72, 0x65, 0x73, 0x70, 0x6f, 0x6e,
	0x73, 0x65, 0x12, 0x2c, 0x0a, 0x0c, 0x71, 0x75, 0x65, 0x72, 0x79, 0x4d, 0x65, 0x74, 0x68, 0x6f,
	0x64, 0x73, 0x12, 0x10, 0x2e, 0x4f, 0x75, 0x74, 0x73, 0x69, 0x64, 0x65, 0x72, 0x4e, 0x65, 0x74,
	0x77, 0x6f, 0x72, 0x6b, 0x1a, 0x08, 0x2e, 0x49, 0x4f, 0x50, 0x49, 0x6e, 0x66, 0x6f, 0x30, 0x01,
	0x42, 0x0f, 0x5a, 0x0d, 0x72, 0x65, 0x6c, 0x61, 0x79, 0x2f, 0x73, 0x63, 0x72, 0x69, 0x70, 0x74,
	0x73, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_iop_proto_rawDescOnce sync.Once
	file_iop_proto_rawDescData = file_iop_proto_rawDesc
)

func file_iop_proto_rawDescGZIP() []byte {
	file_iop_proto_rawDescOnce.Do(func() {
		file_iop_proto_rawDescData = protoimpl.X.CompressGZIP(file_iop_proto_rawDescData)
	})
	return file_iop_proto_rawDescData
}

var file_iop_proto_msgTypes = make([]protoimpl.MessageInfo, 3)
var file_iop_proto_goTypes = []interface{}{
	(*OutsiderNetwork)(nil), // 0: OutsiderNetwork
	(*Response)(nil),        // 1: response
	(*IOPInfo)(nil),         // 2: IOPInfo
}
var file_iop_proto_depIdxs = []int32{
	2, // 0: IOP.invoke:input_type -> IOPInfo
	0, // 1: IOP.queryMethods:input_type -> OutsiderNetwork
	1, // 2: IOP.invoke:output_type -> response
	2, // 3: IOP.queryMethods:output_type -> IOPInfo
	2, // [2:4] is the sub-list for method output_type
	0, // [0:2] is the sub-list for method input_type
	0, // [0:0] is the sub-list for extension type_name
	0, // [0:0] is the sub-list for extension extendee
	0, // [0:0] is the sub-list for field type_name
}

func init() { file_iop_proto_init() }
func file_iop_proto_init() {
	if File_iop_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_iop_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
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
		file_iop_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
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
		file_iop_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*IOPInfo); i {
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
			RawDescriptor: file_iop_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   3,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_iop_proto_goTypes,
		DependencyIndexes: file_iop_proto_depIdxs,
		MessageInfos:      file_iop_proto_msgTypes,
	}.Build()
	File_iop_proto = out.File
	file_iop_proto_rawDesc = nil
	file_iop_proto_goTypes = nil
	file_iop_proto_depIdxs = nil
}
