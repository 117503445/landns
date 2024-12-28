// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.36.1
// 	protoc        v5.29.2
// source: pkg/grpcgen/dhcp-manager.proto

package grpcgen

import (
	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
	emptypb "google.golang.org/protobuf/types/known/emptypb"
	reflect "reflect"
	sync "sync"
)

const (
	// Verify that this generated code is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(20 - protoimpl.MinVersion)
	// Verify that runtime/protoimpl is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(protoimpl.MaxVersion - 20)
)

type Lease struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	Mac           string                 `protobuf:"bytes,1,opt,name=mac,proto3" json:"mac,omitempty"`
	Ip            string                 `protobuf:"bytes,2,opt,name=ip,proto3" json:"ip,omitempty"`
	Hostname      string                 `protobuf:"bytes,3,opt,name=hostname,proto3" json:"hostname,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *Lease) Reset() {
	*x = Lease{}
	mi := &file_pkg_grpcgen_dhcp_manager_proto_msgTypes[0]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *Lease) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Lease) ProtoMessage() {}

func (x *Lease) ProtoReflect() protoreflect.Message {
	mi := &file_pkg_grpcgen_dhcp_manager_proto_msgTypes[0]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Lease.ProtoReflect.Descriptor instead.
func (*Lease) Descriptor() ([]byte, []int) {
	return file_pkg_grpcgen_dhcp_manager_proto_rawDescGZIP(), []int{0}
}

func (x *Lease) GetMac() string {
	if x != nil {
		return x.Mac
	}
	return ""
}

func (x *Lease) GetIp() string {
	if x != nil {
		return x.Ip
	}
	return ""
}

func (x *Lease) GetHostname() string {
	if x != nil {
		return x.Hostname
	}
	return ""
}

type GetLeasesResponse struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	Leases        []*Lease               `protobuf:"bytes,1,rep,name=leases,proto3" json:"leases,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *GetLeasesResponse) Reset() {
	*x = GetLeasesResponse{}
	mi := &file_pkg_grpcgen_dhcp_manager_proto_msgTypes[1]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *GetLeasesResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetLeasesResponse) ProtoMessage() {}

func (x *GetLeasesResponse) ProtoReflect() protoreflect.Message {
	mi := &file_pkg_grpcgen_dhcp_manager_proto_msgTypes[1]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetLeasesResponse.ProtoReflect.Descriptor instead.
func (*GetLeasesResponse) Descriptor() ([]byte, []int) {
	return file_pkg_grpcgen_dhcp_manager_proto_rawDescGZIP(), []int{1}
}

func (x *GetLeasesResponse) GetLeases() []*Lease {
	if x != nil {
		return x.Leases
	}
	return nil
}

var File_pkg_grpcgen_dhcp_manager_proto protoreflect.FileDescriptor

var file_pkg_grpcgen_dhcp_manager_proto_rawDesc = []byte{
	0x0a, 0x1e, 0x70, 0x6b, 0x67, 0x2f, 0x67, 0x72, 0x70, 0x63, 0x67, 0x65, 0x6e, 0x2f, 0x64, 0x68,
	0x63, 0x70, 0x2d, 0x6d, 0x61, 0x6e, 0x61, 0x67, 0x65, 0x72, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x12, 0x0c, 0x64, 0x68, 0x63, 0x70, 0x5f, 0x6d, 0x61, 0x6e, 0x61, 0x67, 0x65, 0x72, 0x1a, 0x1b,
	0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2f,
	0x65, 0x6d, 0x70, 0x74, 0x79, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0x45, 0x0a, 0x05, 0x4c,
	0x65, 0x61, 0x73, 0x65, 0x12, 0x10, 0x0a, 0x03, 0x6d, 0x61, 0x63, 0x18, 0x01, 0x20, 0x01, 0x28,
	0x09, 0x52, 0x03, 0x6d, 0x61, 0x63, 0x12, 0x0e, 0x0a, 0x02, 0x69, 0x70, 0x18, 0x02, 0x20, 0x01,
	0x28, 0x09, 0x52, 0x02, 0x69, 0x70, 0x12, 0x1a, 0x0a, 0x08, 0x68, 0x6f, 0x73, 0x74, 0x6e, 0x61,
	0x6d, 0x65, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x68, 0x6f, 0x73, 0x74, 0x6e, 0x61,
	0x6d, 0x65, 0x22, 0x40, 0x0a, 0x11, 0x47, 0x65, 0x74, 0x4c, 0x65, 0x61, 0x73, 0x65, 0x73, 0x52,
	0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x2b, 0x0a, 0x06, 0x6c, 0x65, 0x61, 0x73, 0x65,
	0x73, 0x18, 0x01, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x13, 0x2e, 0x64, 0x68, 0x63, 0x70, 0x5f, 0x6d,
	0x61, 0x6e, 0x61, 0x67, 0x65, 0x72, 0x2e, 0x4c, 0x65, 0x61, 0x73, 0x65, 0x52, 0x06, 0x6c, 0x65,
	0x61, 0x73, 0x65, 0x73, 0x32, 0x53, 0x0a, 0x0b, 0x44, 0x48, 0x43, 0x50, 0x4d, 0x61, 0x6e, 0x61,
	0x67, 0x65, 0x72, 0x12, 0x44, 0x0a, 0x09, 0x47, 0x65, 0x74, 0x4c, 0x65, 0x61, 0x73, 0x65, 0x73,
	0x12, 0x16, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62,
	0x75, 0x66, 0x2e, 0x45, 0x6d, 0x70, 0x74, 0x79, 0x1a, 0x1f, 0x2e, 0x64, 0x68, 0x63, 0x70, 0x5f,
	0x6d, 0x61, 0x6e, 0x61, 0x67, 0x65, 0x72, 0x2e, 0x47, 0x65, 0x74, 0x4c, 0x65, 0x61, 0x73, 0x65,
	0x73, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x42, 0x0d, 0x5a, 0x0b, 0x70, 0x6b, 0x67,
	0x2f, 0x67, 0x72, 0x70, 0x63, 0x67, 0x65, 0x6e, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_pkg_grpcgen_dhcp_manager_proto_rawDescOnce sync.Once
	file_pkg_grpcgen_dhcp_manager_proto_rawDescData = file_pkg_grpcgen_dhcp_manager_proto_rawDesc
)

func file_pkg_grpcgen_dhcp_manager_proto_rawDescGZIP() []byte {
	file_pkg_grpcgen_dhcp_manager_proto_rawDescOnce.Do(func() {
		file_pkg_grpcgen_dhcp_manager_proto_rawDescData = protoimpl.X.CompressGZIP(file_pkg_grpcgen_dhcp_manager_proto_rawDescData)
	})
	return file_pkg_grpcgen_dhcp_manager_proto_rawDescData
}

var file_pkg_grpcgen_dhcp_manager_proto_msgTypes = make([]protoimpl.MessageInfo, 2)
var file_pkg_grpcgen_dhcp_manager_proto_goTypes = []any{
	(*Lease)(nil),             // 0: dhcp_manager.Lease
	(*GetLeasesResponse)(nil), // 1: dhcp_manager.GetLeasesResponse
	(*emptypb.Empty)(nil),     // 2: google.protobuf.Empty
}
var file_pkg_grpcgen_dhcp_manager_proto_depIdxs = []int32{
	0, // 0: dhcp_manager.GetLeasesResponse.leases:type_name -> dhcp_manager.Lease
	2, // 1: dhcp_manager.DHCPManager.GetLeases:input_type -> google.protobuf.Empty
	1, // 2: dhcp_manager.DHCPManager.GetLeases:output_type -> dhcp_manager.GetLeasesResponse
	2, // [2:3] is the sub-list for method output_type
	1, // [1:2] is the sub-list for method input_type
	1, // [1:1] is the sub-list for extension type_name
	1, // [1:1] is the sub-list for extension extendee
	0, // [0:1] is the sub-list for field type_name
}

func init() { file_pkg_grpcgen_dhcp_manager_proto_init() }
func file_pkg_grpcgen_dhcp_manager_proto_init() {
	if File_pkg_grpcgen_dhcp_manager_proto != nil {
		return
	}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_pkg_grpcgen_dhcp_manager_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   2,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_pkg_grpcgen_dhcp_manager_proto_goTypes,
		DependencyIndexes: file_pkg_grpcgen_dhcp_manager_proto_depIdxs,
		MessageInfos:      file_pkg_grpcgen_dhcp_manager_proto_msgTypes,
	}.Build()
	File_pkg_grpcgen_dhcp_manager_proto = out.File
	file_pkg_grpcgen_dhcp_manager_proto_rawDesc = nil
	file_pkg_grpcgen_dhcp_manager_proto_goTypes = nil
	file_pkg_grpcgen_dhcp_manager_proto_depIdxs = nil
}
