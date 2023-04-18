// protoc --go_out=. --go_opt=paths=source_relative \
//	--go-grpc_out=. --go-grpc_opt=paths=source_relative \
//	bypass.proto

// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.25.0-devel
// 	protoc        v3.14.0
// source: hosts.proto

package proto

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

type LookupRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Network string `protobuf:"bytes,1,opt,name=network,proto3" json:"network,omitempty"`
	Host    string `protobuf:"bytes,2,opt,name=host,proto3" json:"host,omitempty"`
}

func (x *LookupRequest) Reset() {
	*x = LookupRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_hosts_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *LookupRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*LookupRequest) ProtoMessage() {}

func (x *LookupRequest) ProtoReflect() protoreflect.Message {
	mi := &file_hosts_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use LookupRequest.ProtoReflect.Descriptor instead.
func (*LookupRequest) Descriptor() ([]byte, []int) {
	return file_hosts_proto_rawDescGZIP(), []int{0}
}

func (x *LookupRequest) GetNetwork() string {
	if x != nil {
		return x.Network
	}
	return ""
}

func (x *LookupRequest) GetHost() string {
	if x != nil {
		return x.Host
	}
	return ""
}

type LookupReply struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Ips []string `protobuf:"bytes,1,rep,name=ips,proto3" json:"ips,omitempty"`
	Ok  bool     `protobuf:"varint,2,opt,name=ok,proto3" json:"ok,omitempty"`
}

func (x *LookupReply) Reset() {
	*x = LookupReply{}
	if protoimpl.UnsafeEnabled {
		mi := &file_hosts_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *LookupReply) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*LookupReply) ProtoMessage() {}

func (x *LookupReply) ProtoReflect() protoreflect.Message {
	mi := &file_hosts_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use LookupReply.ProtoReflect.Descriptor instead.
func (*LookupReply) Descriptor() ([]byte, []int) {
	return file_hosts_proto_rawDescGZIP(), []int{1}
}

func (x *LookupReply) GetIps() []string {
	if x != nil {
		return x.Ips
	}
	return nil
}

func (x *LookupReply) GetOk() bool {
	if x != nil {
		return x.Ok
	}
	return false
}

var File_hosts_proto protoreflect.FileDescriptor

var file_hosts_proto_rawDesc = []byte{
	0x0a, 0x0b, 0x68, 0x6f, 0x73, 0x74, 0x73, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x05, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x22, 0x3d, 0x0a, 0x0d, 0x4c, 0x6f, 0x6f, 0x6b, 0x75, 0x70, 0x52, 0x65,
	0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x18, 0x0a, 0x07, 0x6e, 0x65, 0x74, 0x77, 0x6f, 0x72, 0x6b,
	0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x07, 0x6e, 0x65, 0x74, 0x77, 0x6f, 0x72, 0x6b, 0x12,
	0x12, 0x0a, 0x04, 0x68, 0x6f, 0x73, 0x74, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x68,
	0x6f, 0x73, 0x74, 0x22, 0x2f, 0x0a, 0x0b, 0x4c, 0x6f, 0x6f, 0x6b, 0x75, 0x70, 0x52, 0x65, 0x70,
	0x6c, 0x79, 0x12, 0x10, 0x0a, 0x03, 0x69, 0x70, 0x73, 0x18, 0x01, 0x20, 0x03, 0x28, 0x09, 0x52,
	0x03, 0x69, 0x70, 0x73, 0x12, 0x0e, 0x0a, 0x02, 0x6f, 0x6b, 0x18, 0x02, 0x20, 0x01, 0x28, 0x08,
	0x52, 0x02, 0x6f, 0x6b, 0x32, 0x40, 0x0a, 0x0a, 0x48, 0x6f, 0x73, 0x74, 0x4d, 0x61, 0x70, 0x70,
	0x65, 0x72, 0x12, 0x32, 0x0a, 0x06, 0x4c, 0x6f, 0x6f, 0x6b, 0x75, 0x70, 0x12, 0x14, 0x2e, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x2e, 0x4c, 0x6f, 0x6f, 0x6b, 0x75, 0x70, 0x52, 0x65, 0x71, 0x75, 0x65,
	0x73, 0x74, 0x1a, 0x12, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2e, 0x4c, 0x6f, 0x6f, 0x6b, 0x75,
	0x70, 0x52, 0x65, 0x70, 0x6c, 0x79, 0x42, 0x27, 0x5a, 0x25, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62,
	0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x67, 0x6f, 0x2d, 0x67, 0x6f, 0x73, 0x74, 0x2f, 0x70, 0x6c, 0x75,
	0x67, 0x69, 0x6e, 0x2f, 0x68, 0x6f, 0x73, 0x74, 0x73, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62,
	0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_hosts_proto_rawDescOnce sync.Once
	file_hosts_proto_rawDescData = file_hosts_proto_rawDesc
)

func file_hosts_proto_rawDescGZIP() []byte {
	file_hosts_proto_rawDescOnce.Do(func() {
		file_hosts_proto_rawDescData = protoimpl.X.CompressGZIP(file_hosts_proto_rawDescData)
	})
	return file_hosts_proto_rawDescData
}

var file_hosts_proto_msgTypes = make([]protoimpl.MessageInfo, 2)
var file_hosts_proto_goTypes = []interface{}{
	(*LookupRequest)(nil), // 0: proto.LookupRequest
	(*LookupReply)(nil),   // 1: proto.LookupReply
}
var file_hosts_proto_depIdxs = []int32{
	0, // 0: proto.HostMapper.Lookup:input_type -> proto.LookupRequest
	1, // 1: proto.HostMapper.Lookup:output_type -> proto.LookupReply
	1, // [1:2] is the sub-list for method output_type
	0, // [0:1] is the sub-list for method input_type
	0, // [0:0] is the sub-list for extension type_name
	0, // [0:0] is the sub-list for extension extendee
	0, // [0:0] is the sub-list for field type_name
}

func init() { file_hosts_proto_init() }
func file_hosts_proto_init() {
	if File_hosts_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_hosts_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*LookupRequest); i {
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
		file_hosts_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*LookupReply); i {
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
			RawDescriptor: file_hosts_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   2,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_hosts_proto_goTypes,
		DependencyIndexes: file_hosts_proto_depIdxs,
		MessageInfos:      file_hosts_proto_msgTypes,
	}.Build()
	File_hosts_proto = out.File
	file_hosts_proto_rawDesc = nil
	file_hosts_proto_goTypes = nil
	file_hosts_proto_depIdxs = nil
}
