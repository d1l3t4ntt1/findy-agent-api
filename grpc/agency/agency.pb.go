// Copyright 2020 Harri @ OP Techlab.
//

// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.25.0-devel
// 	protoc        v3.13.0
// source: agency.proto

package agency

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

type Cmd_Type int32

const (
	Cmd_PING    Cmd_Type = 0
	Cmd_LOGGING Cmd_Type = 1
	Cmd_COUNT   Cmd_Type = 2
)

// Enum value maps for Cmd_Type.
var (
	Cmd_Type_name = map[int32]string{
		0: "PING",
		1: "LOGGING",
		2: "COUNT",
	}
	Cmd_Type_value = map[string]int32{
		"PING":    0,
		"LOGGING": 1,
		"COUNT":   2,
	}
)

func (x Cmd_Type) Enum() *Cmd_Type {
	p := new(Cmd_Type)
	*p = x
	return p
}

func (x Cmd_Type) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (Cmd_Type) Descriptor() protoreflect.EnumDescriptor {
	return file_agency_proto_enumTypes[0].Descriptor()
}

func (Cmd_Type) Type() protoreflect.EnumType {
	return &file_agency_proto_enumTypes[0]
}

func (x Cmd_Type) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use Cmd_Type.Descriptor instead.
func (Cmd_Type) EnumDescriptor() ([]byte, []int) {
	return file_agency_proto_rawDescGZIP(), []int{2, 0}
}

type AgencyStatus struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id   string `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
	Info string `protobuf:"bytes,2,opt,name=info,proto3" json:"info,omitempty"`
}

func (x *AgencyStatus) Reset() {
	*x = AgencyStatus{}
	if protoimpl.UnsafeEnabled {
		mi := &file_agency_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *AgencyStatus) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*AgencyStatus) ProtoMessage() {}

func (x *AgencyStatus) ProtoReflect() protoreflect.Message {
	mi := &file_agency_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use AgencyStatus.ProtoReflect.Descriptor instead.
func (*AgencyStatus) Descriptor() ([]byte, []int) {
	return file_agency_proto_rawDescGZIP(), []int{0}
}

func (x *AgencyStatus) GetId() string {
	if x != nil {
		return x.Id
	}
	return ""
}

func (x *AgencyStatus) GetInfo() string {
	if x != nil {
		return x.Info
	}
	return ""
}

type DataHook struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id string `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
}

func (x *DataHook) Reset() {
	*x = DataHook{}
	if protoimpl.UnsafeEnabled {
		mi := &file_agency_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *DataHook) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*DataHook) ProtoMessage() {}

func (x *DataHook) ProtoReflect() protoreflect.Message {
	mi := &file_agency_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use DataHook.ProtoReflect.Descriptor instead.
func (*DataHook) Descriptor() ([]byte, []int) {
	return file_agency_proto_rawDescGZIP(), []int{1}
}

func (x *DataHook) GetId() string {
	if x != nil {
		return x.Id
	}
	return ""
}

type Cmd struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Type Cmd_Type `protobuf:"varint,1,opt,name=type,proto3,enum=agency.Cmd_Type" json:"type,omitempty"`
	// Types that are assignable to Request:
	//	*Cmd_Logging
	Request isCmd_Request `protobuf_oneof:"Request"`
}

func (x *Cmd) Reset() {
	*x = Cmd{}
	if protoimpl.UnsafeEnabled {
		mi := &file_agency_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Cmd) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Cmd) ProtoMessage() {}

func (x *Cmd) ProtoReflect() protoreflect.Message {
	mi := &file_agency_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Cmd.ProtoReflect.Descriptor instead.
func (*Cmd) Descriptor() ([]byte, []int) {
	return file_agency_proto_rawDescGZIP(), []int{2}
}

func (x *Cmd) GetType() Cmd_Type {
	if x != nil {
		return x.Type
	}
	return Cmd_PING
}

func (m *Cmd) GetRequest() isCmd_Request {
	if m != nil {
		return m.Request
	}
	return nil
}

func (x *Cmd) GetLogging() string {
	if x, ok := x.GetRequest().(*Cmd_Logging); ok {
		return x.Logging
	}
	return ""
}

type isCmd_Request interface {
	isCmd_Request()
}

type Cmd_Logging struct {
	Logging string `protobuf:"bytes,2,opt,name=logging,proto3,oneof"`
}

func (*Cmd_Logging) isCmd_Request() {}

type CmdReturn struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Type Cmd_Type `protobuf:"varint,1,opt,name=type,proto3,enum=agency.Cmd_Type" json:"type,omitempty"`
	// Types that are assignable to Response:
	//	*CmdReturn_Ping
	//	*CmdReturn_Count
	Response isCmdReturn_Response `protobuf_oneof:"Response"`
}

func (x *CmdReturn) Reset() {
	*x = CmdReturn{}
	if protoimpl.UnsafeEnabled {
		mi := &file_agency_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *CmdReturn) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CmdReturn) ProtoMessage() {}

func (x *CmdReturn) ProtoReflect() protoreflect.Message {
	mi := &file_agency_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CmdReturn.ProtoReflect.Descriptor instead.
func (*CmdReturn) Descriptor() ([]byte, []int) {
	return file_agency_proto_rawDescGZIP(), []int{3}
}

func (x *CmdReturn) GetType() Cmd_Type {
	if x != nil {
		return x.Type
	}
	return Cmd_PING
}

func (m *CmdReturn) GetResponse() isCmdReturn_Response {
	if m != nil {
		return m.Response
	}
	return nil
}

func (x *CmdReturn) GetPing() string {
	if x, ok := x.GetResponse().(*CmdReturn_Ping); ok {
		return x.Ping
	}
	return ""
}

func (x *CmdReturn) GetCount() string {
	if x, ok := x.GetResponse().(*CmdReturn_Count); ok {
		return x.Count
	}
	return ""
}

type isCmdReturn_Response interface {
	isCmdReturn_Response()
}

type CmdReturn_Ping struct {
	Ping string `protobuf:"bytes,2,opt,name=ping,proto3,oneof"`
}

type CmdReturn_Count struct {
	Count string `protobuf:"bytes,3,opt,name=count,proto3,oneof"`
}

func (*CmdReturn_Ping) isCmdReturn_Response() {}

func (*CmdReturn_Count) isCmdReturn_Response() {}

var File_agency_proto protoreflect.FileDescriptor

var file_agency_proto_rawDesc = []byte{
	0x0a, 0x0c, 0x61, 0x67, 0x65, 0x6e, 0x63, 0x79, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x06,
	0x61, 0x67, 0x65, 0x6e, 0x63, 0x79, 0x22, 0x32, 0x0a, 0x0c, 0x41, 0x67, 0x65, 0x6e, 0x63, 0x79,
	0x53, 0x74, 0x61, 0x74, 0x75, 0x73, 0x12, 0x0e, 0x0a, 0x02, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01,
	0x28, 0x09, 0x52, 0x02, 0x69, 0x64, 0x12, 0x12, 0x0a, 0x04, 0x69, 0x6e, 0x66, 0x6f, 0x18, 0x02,
	0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x69, 0x6e, 0x66, 0x6f, 0x22, 0x1a, 0x0a, 0x08, 0x44, 0x61,
	0x74, 0x61, 0x48, 0x6f, 0x6f, 0x6b, 0x12, 0x0e, 0x0a, 0x02, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01,
	0x28, 0x09, 0x52, 0x02, 0x69, 0x64, 0x22, 0x7c, 0x0a, 0x03, 0x43, 0x6d, 0x64, 0x12, 0x24, 0x0a,
	0x04, 0x74, 0x79, 0x70, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0e, 0x32, 0x10, 0x2e, 0x61, 0x67,
	0x65, 0x6e, 0x63, 0x79, 0x2e, 0x43, 0x6d, 0x64, 0x2e, 0x54, 0x79, 0x70, 0x65, 0x52, 0x04, 0x74,
	0x79, 0x70, 0x65, 0x12, 0x1a, 0x0a, 0x07, 0x6c, 0x6f, 0x67, 0x67, 0x69, 0x6e, 0x67, 0x18, 0x02,
	0x20, 0x01, 0x28, 0x09, 0x48, 0x00, 0x52, 0x07, 0x6c, 0x6f, 0x67, 0x67, 0x69, 0x6e, 0x67, 0x22,
	0x28, 0x0a, 0x04, 0x54, 0x79, 0x70, 0x65, 0x12, 0x08, 0x0a, 0x04, 0x50, 0x49, 0x4e, 0x47, 0x10,
	0x00, 0x12, 0x0b, 0x0a, 0x07, 0x4c, 0x4f, 0x47, 0x47, 0x49, 0x4e, 0x47, 0x10, 0x01, 0x12, 0x09,
	0x0a, 0x05, 0x43, 0x4f, 0x55, 0x4e, 0x54, 0x10, 0x02, 0x42, 0x09, 0x0a, 0x07, 0x52, 0x65, 0x71,
	0x75, 0x65, 0x73, 0x74, 0x22, 0x6b, 0x0a, 0x09, 0x43, 0x6d, 0x64, 0x52, 0x65, 0x74, 0x75, 0x72,
	0x6e, 0x12, 0x24, 0x0a, 0x04, 0x74, 0x79, 0x70, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0e, 0x32,
	0x10, 0x2e, 0x61, 0x67, 0x65, 0x6e, 0x63, 0x79, 0x2e, 0x43, 0x6d, 0x64, 0x2e, 0x54, 0x79, 0x70,
	0x65, 0x52, 0x04, 0x74, 0x79, 0x70, 0x65, 0x12, 0x14, 0x0a, 0x04, 0x70, 0x69, 0x6e, 0x67, 0x18,
	0x02, 0x20, 0x01, 0x28, 0x09, 0x48, 0x00, 0x52, 0x04, 0x70, 0x69, 0x6e, 0x67, 0x12, 0x16, 0x0a,
	0x05, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x48, 0x00, 0x52, 0x05,
	0x63, 0x6f, 0x75, 0x6e, 0x74, 0x42, 0x0a, 0x0a, 0x08, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73,
	0x65, 0x32, 0x3f, 0x0a, 0x06, 0x41, 0x67, 0x65, 0x6e, 0x63, 0x79, 0x12, 0x35, 0x0a, 0x07, 0x50,
	0x53, 0x4d, 0x48, 0x6f, 0x6f, 0x6b, 0x12, 0x10, 0x2e, 0x61, 0x67, 0x65, 0x6e, 0x63, 0x79, 0x2e,
	0x44, 0x61, 0x74, 0x61, 0x48, 0x6f, 0x6f, 0x6b, 0x1a, 0x14, 0x2e, 0x61, 0x67, 0x65, 0x6e, 0x63,
	0x79, 0x2e, 0x41, 0x67, 0x65, 0x6e, 0x63, 0x79, 0x53, 0x74, 0x61, 0x74, 0x75, 0x73, 0x22, 0x00,
	0x30, 0x01, 0x32, 0x33, 0x0a, 0x06, 0x44, 0x65, 0x76, 0x4f, 0x70, 0x73, 0x12, 0x29, 0x0a, 0x05,
	0x45, 0x6e, 0x74, 0x65, 0x72, 0x12, 0x0b, 0x2e, 0x61, 0x67, 0x65, 0x6e, 0x63, 0x79, 0x2e, 0x43,
	0x6d, 0x64, 0x1a, 0x11, 0x2e, 0x61, 0x67, 0x65, 0x6e, 0x63, 0x79, 0x2e, 0x43, 0x6d, 0x64, 0x52,
	0x65, 0x74, 0x75, 0x72, 0x6e, 0x22, 0x00, 0x42, 0x36, 0x5a, 0x34, 0x67, 0x69, 0x74, 0x68, 0x75,
	0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x66, 0x69, 0x6e, 0x64, 0x79, 0x2d, 0x6e, 0x65, 0x74, 0x77,
	0x6f, 0x72, 0x6b, 0x2f, 0x66, 0x69, 0x6e, 0x64, 0x79, 0x2d, 0x61, 0x67, 0x65, 0x6e, 0x74, 0x2d,
	0x61, 0x70, 0x69, 0x2f, 0x67, 0x72, 0x70, 0x63, 0x2f, 0x61, 0x67, 0x65, 0x6e, 0x63, 0x79, 0x62,
	0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_agency_proto_rawDescOnce sync.Once
	file_agency_proto_rawDescData = file_agency_proto_rawDesc
)

func file_agency_proto_rawDescGZIP() []byte {
	file_agency_proto_rawDescOnce.Do(func() {
		file_agency_proto_rawDescData = protoimpl.X.CompressGZIP(file_agency_proto_rawDescData)
	})
	return file_agency_proto_rawDescData
}

var file_agency_proto_enumTypes = make([]protoimpl.EnumInfo, 1)
var file_agency_proto_msgTypes = make([]protoimpl.MessageInfo, 4)
var file_agency_proto_goTypes = []interface{}{
	(Cmd_Type)(0),        // 0: agency.Cmd.Type
	(*AgencyStatus)(nil), // 1: agency.AgencyStatus
	(*DataHook)(nil),     // 2: agency.DataHook
	(*Cmd)(nil),          // 3: agency.Cmd
	(*CmdReturn)(nil),    // 4: agency.CmdReturn
}
var file_agency_proto_depIdxs = []int32{
	0, // 0: agency.Cmd.type:type_name -> agency.Cmd.Type
	0, // 1: agency.CmdReturn.type:type_name -> agency.Cmd.Type
	2, // 2: agency.Agency.PSMHook:input_type -> agency.DataHook
	3, // 3: agency.DevOps.Enter:input_type -> agency.Cmd
	1, // 4: agency.Agency.PSMHook:output_type -> agency.AgencyStatus
	4, // 5: agency.DevOps.Enter:output_type -> agency.CmdReturn
	4, // [4:6] is the sub-list for method output_type
	2, // [2:4] is the sub-list for method input_type
	2, // [2:2] is the sub-list for extension type_name
	2, // [2:2] is the sub-list for extension extendee
	0, // [0:2] is the sub-list for field type_name
}

func init() { file_agency_proto_init() }
func file_agency_proto_init() {
	if File_agency_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_agency_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*AgencyStatus); i {
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
		file_agency_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*DataHook); i {
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
		file_agency_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Cmd); i {
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
		file_agency_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*CmdReturn); i {
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
	file_agency_proto_msgTypes[2].OneofWrappers = []interface{}{
		(*Cmd_Logging)(nil),
	}
	file_agency_proto_msgTypes[3].OneofWrappers = []interface{}{
		(*CmdReturn_Ping)(nil),
		(*CmdReturn_Count)(nil),
	}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_agency_proto_rawDesc,
			NumEnums:      1,
			NumMessages:   4,
			NumExtensions: 0,
			NumServices:   2,
		},
		GoTypes:           file_agency_proto_goTypes,
		DependencyIndexes: file_agency_proto_depIdxs,
		EnumInfos:         file_agency_proto_enumTypes,
		MessageInfos:      file_agency_proto_msgTypes,
	}.Build()
	File_agency_proto = out.File
	file_agency_proto_rawDesc = nil
	file_agency_proto_goTypes = nil
	file_agency_proto_depIdxs = nil
}
