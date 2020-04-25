// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.21.0
// 	protoc        v3.11.2
// source: dnsregistry.proto

package pb

import (
	context "context"
	proto "github.com/golang/protobuf/proto"
	empty "github.com/golang/protobuf/ptypes/empty"
	grpc "google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
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

// This is a compile-time assertion that a sufficiently up-to-date version
// of the legacy proto package is being used.
const _ = proto.ProtoPackageIsVersion4

type DnsAssignRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Domain    string `protobuf:"bytes,1,opt,name=domain,proto3" json:"domain,omitempty"`
	Ip4       []byte `protobuf:"bytes,2,opt,name=ip4,proto3" json:"ip4,omitempty"`
	Ip6       []byte `protobuf:"bytes,3,opt,name=ip6,proto3" json:"ip6,omitempty"`
	Ttl       uint32 `protobuf:"varint,4,opt,name=ttl,proto3" json:"ttl,omitempty"`
	ReplaceOk bool   `protobuf:"varint,5,opt,name=replaceOk,proto3" json:"replaceOk,omitempty"`
}

func (x *DnsAssignRequest) Reset() {
	*x = DnsAssignRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_dnsregistry_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *DnsAssignRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*DnsAssignRequest) ProtoMessage() {}

func (x *DnsAssignRequest) ProtoReflect() protoreflect.Message {
	mi := &file_dnsregistry_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use DnsAssignRequest.ProtoReflect.Descriptor instead.
func (*DnsAssignRequest) Descriptor() ([]byte, []int) {
	return file_dnsregistry_proto_rawDescGZIP(), []int{0}
}

func (x *DnsAssignRequest) GetDomain() string {
	if x != nil {
		return x.Domain
	}
	return ""
}

func (x *DnsAssignRequest) GetIp4() []byte {
	if x != nil {
		return x.Ip4
	}
	return nil
}

func (x *DnsAssignRequest) GetIp6() []byte {
	if x != nil {
		return x.Ip6
	}
	return nil
}

func (x *DnsAssignRequest) GetTtl() uint32 {
	if x != nil {
		return x.Ttl
	}
	return 0
}

func (x *DnsAssignRequest) GetReplaceOk() bool {
	if x != nil {
		return x.ReplaceOk
	}
	return false
}

type DnsReleaseRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Domain string `protobuf:"bytes,1,opt,name=domain,proto3" json:"domain,omitempty"`
}

func (x *DnsReleaseRequest) Reset() {
	*x = DnsReleaseRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_dnsregistry_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *DnsReleaseRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*DnsReleaseRequest) ProtoMessage() {}

func (x *DnsReleaseRequest) ProtoReflect() protoreflect.Message {
	mi := &file_dnsregistry_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use DnsReleaseRequest.ProtoReflect.Descriptor instead.
func (*DnsReleaseRequest) Descriptor() ([]byte, []int) {
	return file_dnsregistry_proto_rawDescGZIP(), []int{1}
}

func (x *DnsReleaseRequest) GetDomain() string {
	if x != nil {
		return x.Domain
	}
	return ""
}

var File_dnsregistry_proto protoreflect.FileDescriptor

var file_dnsregistry_proto_rawDesc = []byte{
	0x0a, 0x11, 0x64, 0x6e, 0x73, 0x72, 0x65, 0x67, 0x69, 0x73, 0x74, 0x72, 0x79, 0x2e, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x12, 0x02, 0x70, 0x62, 0x1a, 0x1b, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2f, 0x65, 0x6d, 0x70, 0x74, 0x79, 0x2e, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x22, 0x7e, 0x0a, 0x10, 0x44, 0x6e, 0x73, 0x41, 0x73, 0x73, 0x69, 0x67,
	0x6e, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x16, 0x0a, 0x06, 0x64, 0x6f, 0x6d, 0x61,
	0x69, 0x6e, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x06, 0x64, 0x6f, 0x6d, 0x61, 0x69, 0x6e,
	0x12, 0x10, 0x0a, 0x03, 0x69, 0x70, 0x34, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0c, 0x52, 0x03, 0x69,
	0x70, 0x34, 0x12, 0x10, 0x0a, 0x03, 0x69, 0x70, 0x36, 0x18, 0x03, 0x20, 0x01, 0x28, 0x0c, 0x52,
	0x03, 0x69, 0x70, 0x36, 0x12, 0x10, 0x0a, 0x03, 0x74, 0x74, 0x6c, 0x18, 0x04, 0x20, 0x01, 0x28,
	0x0d, 0x52, 0x03, 0x74, 0x74, 0x6c, 0x12, 0x1c, 0x0a, 0x09, 0x72, 0x65, 0x70, 0x6c, 0x61, 0x63,
	0x65, 0x4f, 0x6b, 0x18, 0x05, 0x20, 0x01, 0x28, 0x08, 0x52, 0x09, 0x72, 0x65, 0x70, 0x6c, 0x61,
	0x63, 0x65, 0x4f, 0x6b, 0x22, 0x2b, 0x0a, 0x11, 0x44, 0x6e, 0x73, 0x52, 0x65, 0x6c, 0x65, 0x61,
	0x73, 0x65, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x16, 0x0a, 0x06, 0x64, 0x6f, 0x6d,
	0x61, 0x69, 0x6e, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x06, 0x64, 0x6f, 0x6d, 0x61, 0x69,
	0x6e, 0x32, 0x86, 0x01, 0x0a, 0x12, 0x44, 0x6e, 0x73, 0x52, 0x65, 0x67, 0x69, 0x73, 0x74, 0x72,
	0x79, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x12, 0x36, 0x0a, 0x06, 0x41, 0x73, 0x73, 0x69,
	0x67, 0x6e, 0x12, 0x14, 0x2e, 0x70, 0x62, 0x2e, 0x44, 0x6e, 0x73, 0x41, 0x73, 0x73, 0x69, 0x67,
	0x6e, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x16, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c,
	0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x45, 0x6d, 0x70, 0x74, 0x79,
	0x12, 0x38, 0x0a, 0x07, 0x52, 0x65, 0x6c, 0x65, 0x61, 0x73, 0x65, 0x12, 0x15, 0x2e, 0x70, 0x62,
	0x2e, 0x44, 0x6e, 0x73, 0x52, 0x65, 0x6c, 0x65, 0x61, 0x73, 0x65, 0x52, 0x65, 0x71, 0x75, 0x65,
	0x73, 0x74, 0x1a, 0x16, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x62, 0x75, 0x66, 0x2e, 0x45, 0x6d, 0x70, 0x74, 0x79, 0x42, 0x06, 0x5a, 0x04, 0x2e, 0x3b,
	0x70, 0x62, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_dnsregistry_proto_rawDescOnce sync.Once
	file_dnsregistry_proto_rawDescData = file_dnsregistry_proto_rawDesc
)

func file_dnsregistry_proto_rawDescGZIP() []byte {
	file_dnsregistry_proto_rawDescOnce.Do(func() {
		file_dnsregistry_proto_rawDescData = protoimpl.X.CompressGZIP(file_dnsregistry_proto_rawDescData)
	})
	return file_dnsregistry_proto_rawDescData
}

var file_dnsregistry_proto_msgTypes = make([]protoimpl.MessageInfo, 2)
var file_dnsregistry_proto_goTypes = []interface{}{
	(*DnsAssignRequest)(nil),  // 0: pb.DnsAssignRequest
	(*DnsReleaseRequest)(nil), // 1: pb.DnsReleaseRequest
	(*empty.Empty)(nil),       // 2: google.protobuf.Empty
}
var file_dnsregistry_proto_depIdxs = []int32{
	0, // 0: pb.DnsRegistryService.Assign:input_type -> pb.DnsAssignRequest
	1, // 1: pb.DnsRegistryService.Release:input_type -> pb.DnsReleaseRequest
	2, // 2: pb.DnsRegistryService.Assign:output_type -> google.protobuf.Empty
	2, // 3: pb.DnsRegistryService.Release:output_type -> google.protobuf.Empty
	2, // [2:4] is the sub-list for method output_type
	0, // [0:2] is the sub-list for method input_type
	0, // [0:0] is the sub-list for extension type_name
	0, // [0:0] is the sub-list for extension extendee
	0, // [0:0] is the sub-list for field type_name
}

func init() { file_dnsregistry_proto_init() }
func file_dnsregistry_proto_init() {
	if File_dnsregistry_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_dnsregistry_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*DnsAssignRequest); i {
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
		file_dnsregistry_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*DnsReleaseRequest); i {
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
			RawDescriptor: file_dnsregistry_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   2,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_dnsregistry_proto_goTypes,
		DependencyIndexes: file_dnsregistry_proto_depIdxs,
		MessageInfos:      file_dnsregistry_proto_msgTypes,
	}.Build()
	File_dnsregistry_proto = out.File
	file_dnsregistry_proto_rawDesc = nil
	file_dnsregistry_proto_goTypes = nil
	file_dnsregistry_proto_depIdxs = nil
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConnInterface

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion6

// DnsRegistryServiceClient is the client API for DnsRegistryService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type DnsRegistryServiceClient interface {
	Assign(ctx context.Context, in *DnsAssignRequest, opts ...grpc.CallOption) (*empty.Empty, error)
	Release(ctx context.Context, in *DnsReleaseRequest, opts ...grpc.CallOption) (*empty.Empty, error)
}

type dnsRegistryServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewDnsRegistryServiceClient(cc grpc.ClientConnInterface) DnsRegistryServiceClient {
	return &dnsRegistryServiceClient{cc}
}

func (c *dnsRegistryServiceClient) Assign(ctx context.Context, in *DnsAssignRequest, opts ...grpc.CallOption) (*empty.Empty, error) {
	out := new(empty.Empty)
	err := c.cc.Invoke(ctx, "/pb.DnsRegistryService/Assign", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *dnsRegistryServiceClient) Release(ctx context.Context, in *DnsReleaseRequest, opts ...grpc.CallOption) (*empty.Empty, error) {
	out := new(empty.Empty)
	err := c.cc.Invoke(ctx, "/pb.DnsRegistryService/Release", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// DnsRegistryServiceServer is the server API for DnsRegistryService service.
type DnsRegistryServiceServer interface {
	Assign(context.Context, *DnsAssignRequest) (*empty.Empty, error)
	Release(context.Context, *DnsReleaseRequest) (*empty.Empty, error)
}

// UnimplementedDnsRegistryServiceServer can be embedded to have forward compatible implementations.
type UnimplementedDnsRegistryServiceServer struct {
}

func (*UnimplementedDnsRegistryServiceServer) Assign(context.Context, *DnsAssignRequest) (*empty.Empty, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Assign not implemented")
}
func (*UnimplementedDnsRegistryServiceServer) Release(context.Context, *DnsReleaseRequest) (*empty.Empty, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Release not implemented")
}

func RegisterDnsRegistryServiceServer(s *grpc.Server, srv DnsRegistryServiceServer) {
	s.RegisterService(&_DnsRegistryService_serviceDesc, srv)
}

func _DnsRegistryService_Assign_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(DnsAssignRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(DnsRegistryServiceServer).Assign(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/pb.DnsRegistryService/Assign",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(DnsRegistryServiceServer).Assign(ctx, req.(*DnsAssignRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _DnsRegistryService_Release_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(DnsReleaseRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(DnsRegistryServiceServer).Release(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/pb.DnsRegistryService/Release",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(DnsRegistryServiceServer).Release(ctx, req.(*DnsReleaseRequest))
	}
	return interceptor(ctx, in, info, handler)
}

var _DnsRegistryService_serviceDesc = grpc.ServiceDesc{
	ServiceName: "pb.DnsRegistryService",
	HandlerType: (*DnsRegistryServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "Assign",
			Handler:    _DnsRegistryService_Assign_Handler,
		},
		{
			MethodName: "Release",
			Handler:    _DnsRegistryService_Release_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "dnsregistry.proto",
}