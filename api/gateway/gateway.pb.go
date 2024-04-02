// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.33.0
// 	protoc        v5.26.0
// source: gateway/gateway.proto

package gateway

import (
	_ "github.com/grpc-ecosystem/grpc-gateway/v2/protoc-gen-openapiv2/options"
	_ "google.golang.org/genproto/googleapis/api/annotations"
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

type GroupCreateResp struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id uint64 `protobuf:"varint,1,opt,name=id,proto3" json:"id,omitempty"`
}

func (x *GroupCreateResp) Reset() {
	*x = GroupCreateResp{}
	if protoimpl.UnsafeEnabled {
		mi := &file_gateway_gateway_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GroupCreateResp) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GroupCreateResp) ProtoMessage() {}

func (x *GroupCreateResp) ProtoReflect() protoreflect.Message {
	mi := &file_gateway_gateway_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GroupCreateResp.ProtoReflect.Descriptor instead.
func (*GroupCreateResp) Descriptor() ([]byte, []int) {
	return file_gateway_gateway_proto_rawDescGZIP(), []int{0}
}

func (x *GroupCreateResp) GetId() uint64 {
	if x != nil {
		return x.Id
	}
	return 0
}

type GroupPutinReq struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	GroupId    uint64 `protobuf:"varint,2,opt,name=groupId,proto3" json:"groupId,omitempty"`
	ReqId      string `protobuf:"bytes,3,opt,name=reqId,proto3" json:"reqId,omitempty"`
	ReqMsg     string `protobuf:"bytes,4,opt,name=reqMsg,proto3" json:"reqMsg,omitempty"`
	ReqTime    int64  `protobuf:"varint,5,opt,name=reqTime,proto3" json:"reqTime,omitempty"`
	JoinSource int32  `protobuf:"varint,6,opt,name=joinSource,proto3" json:"joinSource,omitempty"`
	InviterUid string `protobuf:"bytes,7,opt,name=inviterUid,proto3" json:"inviterUid,omitempty"`
}

func (x *GroupPutinReq) Reset() {
	*x = GroupPutinReq{}
	if protoimpl.UnsafeEnabled {
		mi := &file_gateway_gateway_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GroupPutinReq) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GroupPutinReq) ProtoMessage() {}

func (x *GroupPutinReq) ProtoReflect() protoreflect.Message {
	mi := &file_gateway_gateway_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GroupPutinReq.ProtoReflect.Descriptor instead.
func (*GroupPutinReq) Descriptor() ([]byte, []int) {
	return file_gateway_gateway_proto_rawDescGZIP(), []int{1}
}

func (x *GroupPutinReq) GetGroupId() uint64 {
	if x != nil {
		return x.GroupId
	}
	return 0
}

func (x *GroupPutinReq) GetReqId() string {
	if x != nil {
		return x.ReqId
	}
	return ""
}

func (x *GroupPutinReq) GetReqMsg() string {
	if x != nil {
		return x.ReqMsg
	}
	return ""
}

func (x *GroupPutinReq) GetReqTime() int64 {
	if x != nil {
		return x.ReqTime
	}
	return 0
}

func (x *GroupPutinReq) GetJoinSource() int32 {
	if x != nil {
		return x.JoinSource
	}
	return 0
}

func (x *GroupPutinReq) GetInviterUid() string {
	if x != nil {
		return x.InviterUid
	}
	return ""
}

type GroupPutinResp struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	GroupId uint64 `protobuf:"varint,1,opt,name=groupId,proto3" json:"groupId,omitempty"`
}

func (x *GroupPutinResp) Reset() {
	*x = GroupPutinResp{}
	if protoimpl.UnsafeEnabled {
		mi := &file_gateway_gateway_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GroupPutinResp) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GroupPutinResp) ProtoMessage() {}

func (x *GroupPutinResp) ProtoReflect() protoreflect.Message {
	mi := &file_gateway_gateway_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GroupPutinResp.ProtoReflect.Descriptor instead.
func (*GroupPutinResp) Descriptor() ([]byte, []int) {
	return file_gateway_gateway_proto_rawDescGZIP(), []int{2}
}

func (x *GroupPutinResp) GetGroupId() uint64 {
	if x != nil {
		return x.GroupId
	}
	return 0
}

var File_gateway_gateway_proto protoreflect.FileDescriptor

var file_gateway_gateway_proto_rawDesc = []byte{
	0x0a, 0x15, 0x67, 0x61, 0x74, 0x65, 0x77, 0x61, 0x79, 0x2f, 0x67, 0x61, 0x74, 0x65, 0x77, 0x61,
	0x79, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x0b, 0x61, 0x70, 0x69, 0x2e, 0x67, 0x61, 0x74,
	0x65, 0x77, 0x61, 0x79, 0x1a, 0x1c, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x61, 0x70, 0x69,
	0x2f, 0x61, 0x6e, 0x6e, 0x6f, 0x74, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x2e, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x1a, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x63, 0x2d, 0x67, 0x65, 0x6e, 0x2d, 0x6f,
	0x70, 0x65, 0x6e, 0x61, 0x70, 0x69, 0x76, 0x32, 0x2f, 0x6f, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x73,
	0x2f, 0x61, 0x6e, 0x6e, 0x6f, 0x74, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x2e, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x22, 0x21, 0x0a, 0x0f, 0x47, 0x72, 0x6f, 0x75, 0x70, 0x43, 0x72, 0x65, 0x61, 0x74,
	0x65, 0x52, 0x65, 0x73, 0x70, 0x12, 0x0e, 0x0a, 0x02, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28,
	0x04, 0x52, 0x02, 0x69, 0x64, 0x22, 0xb1, 0x01, 0x0a, 0x0d, 0x47, 0x72, 0x6f, 0x75, 0x70, 0x50,
	0x75, 0x74, 0x69, 0x6e, 0x52, 0x65, 0x71, 0x12, 0x18, 0x0a, 0x07, 0x67, 0x72, 0x6f, 0x75, 0x70,
	0x49, 0x64, 0x18, 0x02, 0x20, 0x01, 0x28, 0x04, 0x52, 0x07, 0x67, 0x72, 0x6f, 0x75, 0x70, 0x49,
	0x64, 0x12, 0x14, 0x0a, 0x05, 0x72, 0x65, 0x71, 0x49, 0x64, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09,
	0x52, 0x05, 0x72, 0x65, 0x71, 0x49, 0x64, 0x12, 0x16, 0x0a, 0x06, 0x72, 0x65, 0x71, 0x4d, 0x73,
	0x67, 0x18, 0x04, 0x20, 0x01, 0x28, 0x09, 0x52, 0x06, 0x72, 0x65, 0x71, 0x4d, 0x73, 0x67, 0x12,
	0x18, 0x0a, 0x07, 0x72, 0x65, 0x71, 0x54, 0x69, 0x6d, 0x65, 0x18, 0x05, 0x20, 0x01, 0x28, 0x03,
	0x52, 0x07, 0x72, 0x65, 0x71, 0x54, 0x69, 0x6d, 0x65, 0x12, 0x1e, 0x0a, 0x0a, 0x6a, 0x6f, 0x69,
	0x6e, 0x53, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x18, 0x06, 0x20, 0x01, 0x28, 0x05, 0x52, 0x0a, 0x6a,
	0x6f, 0x69, 0x6e, 0x53, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x12, 0x1e, 0x0a, 0x0a, 0x69, 0x6e, 0x76,
	0x69, 0x74, 0x65, 0x72, 0x55, 0x69, 0x64, 0x18, 0x07, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0a, 0x69,
	0x6e, 0x76, 0x69, 0x74, 0x65, 0x72, 0x55, 0x69, 0x64, 0x22, 0x2a, 0x0a, 0x0e, 0x47, 0x72, 0x6f,
	0x75, 0x70, 0x50, 0x75, 0x74, 0x69, 0x6e, 0x52, 0x65, 0x73, 0x70, 0x12, 0x18, 0x0a, 0x07, 0x67,
	0x72, 0x6f, 0x75, 0x70, 0x49, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x04, 0x52, 0x07, 0x67, 0x72,
	0x6f, 0x75, 0x70, 0x49, 0x64, 0x32, 0x81, 0x01, 0x0a, 0x07, 0x47, 0x61, 0x74, 0x65, 0x77, 0x61,
	0x79, 0x12, 0x76, 0x0a, 0x0a, 0x47, 0x72, 0x6f, 0x75, 0x70, 0x50, 0x75, 0x74, 0x69, 0x6e, 0x12,
	0x1a, 0x2e, 0x61, 0x70, 0x69, 0x2e, 0x67, 0x61, 0x74, 0x65, 0x77, 0x61, 0x79, 0x2e, 0x47, 0x72,
	0x6f, 0x75, 0x70, 0x50, 0x75, 0x74, 0x69, 0x6e, 0x52, 0x65, 0x71, 0x1a, 0x1b, 0x2e, 0x61, 0x70,
	0x69, 0x2e, 0x67, 0x61, 0x74, 0x65, 0x77, 0x61, 0x79, 0x2e, 0x47, 0x72, 0x6f, 0x75, 0x70, 0x50,
	0x75, 0x74, 0x69, 0x6e, 0x52, 0x65, 0x73, 0x70, 0x22, 0x2f, 0x92, 0x41, 0x15, 0x0a, 0x05, 0x67,
	0x72, 0x6f, 0x75, 0x70, 0x12, 0x0c, 0xe7, 0x94, 0xb3, 0xe8, 0xaf, 0xb7, 0xe5, 0x85, 0xa5, 0xe7,
	0xbe, 0xa4, 0x82, 0xd3, 0xe4, 0x93, 0x02, 0x11, 0x3a, 0x01, 0x2a, 0x1a, 0x0c, 0x2f, 0x67, 0x72,
	0x6f, 0x75, 0x70, 0x2f, 0x70, 0x75, 0x74, 0x69, 0x6e, 0x42, 0xa4, 0x01, 0x92, 0x41, 0x73, 0x12,
	0x1c, 0x0a, 0x15, 0x6b, 0x72, 0x61, 0x74, 0x6f, 0x73, 0x2d, 0x69, 0x6d, 0x2e, 0x61, 0x70, 0x69,
	0x2e, 0x67, 0x61, 0x74, 0x65, 0x77, 0x61, 0x79, 0x32, 0x03, 0x31, 0x2e, 0x30, 0x2a, 0x02, 0x01,
	0x02, 0x5a, 0x4f, 0x0a, 0x4d, 0x0a, 0x09, 0x54, 0x6f, 0x6b, 0x65, 0x6e, 0x41, 0x75, 0x74, 0x68,
	0x12, 0x40, 0x08, 0x02, 0x12, 0x2b, 0xe7, 0x94, 0xa8, 0xe6, 0x88, 0xb7, 0xe8, 0xae, 0xa4, 0xe8,
	0xaf, 0x81, 0x2c, 0xe6, 0xa0, 0xbc, 0xe5, 0xbc, 0x8f, 0xe4, 0xb8, 0xba, 0x3a, 0x20, 0x42, 0x65,
	0x61, 0x72, 0x65, 0x72, 0x2b, 0xe7, 0xa9, 0xba, 0xe6, 0xa0, 0xbc, 0x2b, 0x74, 0x6f, 0x6b, 0x65,
	0x6e, 0x1a, 0x0d, 0x41, 0x75, 0x74, 0x68, 0x6f, 0x72, 0x69, 0x7a, 0x61, 0x74, 0x69, 0x6f, 0x6e,
	0x20, 0x02, 0x0a, 0x0b, 0x61, 0x70, 0x69, 0x2e, 0x67, 0x61, 0x74, 0x65, 0x77, 0x61, 0x79, 0x50,
	0x01, 0x5a, 0x1d, 0x6b, 0x72, 0x61, 0x74, 0x6f, 0x73, 0x2d, 0x69, 0x6d, 0x2f, 0x61, 0x70, 0x69,
	0x2f, 0x67, 0x61, 0x74, 0x65, 0x77, 0x61, 0x79, 0x3b, 0x67, 0x61, 0x74, 0x65, 0x77, 0x61, 0x79,
	0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_gateway_gateway_proto_rawDescOnce sync.Once
	file_gateway_gateway_proto_rawDescData = file_gateway_gateway_proto_rawDesc
)

func file_gateway_gateway_proto_rawDescGZIP() []byte {
	file_gateway_gateway_proto_rawDescOnce.Do(func() {
		file_gateway_gateway_proto_rawDescData = protoimpl.X.CompressGZIP(file_gateway_gateway_proto_rawDescData)
	})
	return file_gateway_gateway_proto_rawDescData
}

var file_gateway_gateway_proto_msgTypes = make([]protoimpl.MessageInfo, 3)
var file_gateway_gateway_proto_goTypes = []interface{}{
	(*GroupCreateResp)(nil), // 0: api.gateway.GroupCreateResp
	(*GroupPutinReq)(nil),   // 1: api.gateway.GroupPutinReq
	(*GroupPutinResp)(nil),  // 2: api.gateway.GroupPutinResp
}
var file_gateway_gateway_proto_depIdxs = []int32{
	1, // 0: api.gateway.Gateway.GroupPutin:input_type -> api.gateway.GroupPutinReq
	2, // 1: api.gateway.Gateway.GroupPutin:output_type -> api.gateway.GroupPutinResp
	1, // [1:2] is the sub-list for method output_type
	0, // [0:1] is the sub-list for method input_type
	0, // [0:0] is the sub-list for extension type_name
	0, // [0:0] is the sub-list for extension extendee
	0, // [0:0] is the sub-list for field type_name
}

func init() { file_gateway_gateway_proto_init() }
func file_gateway_gateway_proto_init() {
	if File_gateway_gateway_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_gateway_gateway_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GroupCreateResp); i {
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
		file_gateway_gateway_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GroupPutinReq); i {
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
		file_gateway_gateway_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GroupPutinResp); i {
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
			RawDescriptor: file_gateway_gateway_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   3,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_gateway_gateway_proto_goTypes,
		DependencyIndexes: file_gateway_gateway_proto_depIdxs,
		MessageInfos:      file_gateway_gateway_proto_msgTypes,
	}.Build()
	File_gateway_gateway_proto = out.File
	file_gateway_gateway_proto_rawDesc = nil
	file_gateway_gateway_proto_goTypes = nil
	file_gateway_gateway_proto_depIdxs = nil
}