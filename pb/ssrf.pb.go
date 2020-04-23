// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.21.0
// 	protoc        v3.11.2
// source: ssrf.proto

package pb

import (
	context "context"
	proto "github.com/golang/protobuf/proto"
	empty "github.com/golang/protobuf/ptypes/empty"
	_ "google.golang.org/genproto/googleapis/api/annotations"
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

type SSRFAssignRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Domain    string `protobuf:"bytes,1,opt,name=domain,proto3" json:"domain,omitempty"`
	Ttl       uint32 `protobuf:"varint,4,opt,name=ttl,proto3" json:"ttl,omitempty"`
	ReplaceOk bool   `protobuf:"varint,5,opt,name=replaceOk,proto3" json:"replaceOk,omitempty"`
}

func (x *SSRFAssignRequest) Reset() {
	*x = SSRFAssignRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_ssrf_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *SSRFAssignRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*SSRFAssignRequest) ProtoMessage() {}

func (x *SSRFAssignRequest) ProtoReflect() protoreflect.Message {
	mi := &file_ssrf_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use SSRFAssignRequest.ProtoReflect.Descriptor instead.
func (*SSRFAssignRequest) Descriptor() ([]byte, []int) {
	return file_ssrf_proto_rawDescGZIP(), []int{0}
}

func (x *SSRFAssignRequest) GetDomain() string {
	if x != nil {
		return x.Domain
	}
	return ""
}

func (x *SSRFAssignRequest) GetTtl() uint32 {
	if x != nil {
		return x.Ttl
	}
	return 0
}

func (x *SSRFAssignRequest) GetReplaceOk() bool {
	if x != nil {
		return x.ReplaceOk
	}
	return false
}

type SSRFReleaseRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Domain string `protobuf:"bytes,1,opt,name=domain,proto3" json:"domain,omitempty"`
}

func (x *SSRFReleaseRequest) Reset() {
	*x = SSRFReleaseRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_ssrf_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *SSRFReleaseRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*SSRFReleaseRequest) ProtoMessage() {}

func (x *SSRFReleaseRequest) ProtoReflect() protoreflect.Message {
	mi := &file_ssrf_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use SSRFReleaseRequest.ProtoReflect.Descriptor instead.
func (*SSRFReleaseRequest) Descriptor() ([]byte, []int) {
	return file_ssrf_proto_rawDescGZIP(), []int{1}
}

func (x *SSRFReleaseRequest) GetDomain() string {
	if x != nil {
		return x.Domain
	}
	return ""
}

var File_ssrf_proto protoreflect.FileDescriptor

var file_ssrf_proto_rawDesc = []byte{
	0x0a, 0x0a, 0x73, 0x73, 0x72, 0x66, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x02, 0x70, 0x62,
	0x1a, 0x1b, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75,
	0x66, 0x2f, 0x65, 0x6d, 0x70, 0x74, 0x79, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x1c, 0x67,
	0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x61, 0x70, 0x69, 0x2f, 0x61, 0x6e, 0x6e, 0x6f, 0x74, 0x61,
	0x74, 0x69, 0x6f, 0x6e, 0x73, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0x5b, 0x0a, 0x11, 0x53,
	0x53, 0x52, 0x46, 0x41, 0x73, 0x73, 0x69, 0x67, 0x6e, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74,
	0x12, 0x16, 0x0a, 0x06, 0x64, 0x6f, 0x6d, 0x61, 0x69, 0x6e, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09,
	0x52, 0x06, 0x64, 0x6f, 0x6d, 0x61, 0x69, 0x6e, 0x12, 0x10, 0x0a, 0x03, 0x74, 0x74, 0x6c, 0x18,
	0x04, 0x20, 0x01, 0x28, 0x0d, 0x52, 0x03, 0x74, 0x74, 0x6c, 0x12, 0x1c, 0x0a, 0x09, 0x72, 0x65,
	0x70, 0x6c, 0x61, 0x63, 0x65, 0x4f, 0x6b, 0x18, 0x05, 0x20, 0x01, 0x28, 0x08, 0x52, 0x09, 0x72,
	0x65, 0x70, 0x6c, 0x61, 0x63, 0x65, 0x4f, 0x6b, 0x22, 0x2c, 0x0a, 0x12, 0x53, 0x53, 0x52, 0x46,
	0x52, 0x65, 0x6c, 0x65, 0x61, 0x73, 0x65, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x16,
	0x0a, 0x06, 0x64, 0x6f, 0x6d, 0x61, 0x69, 0x6e, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x06,
	0x64, 0x6f, 0x6d, 0x61, 0x69, 0x6e, 0x32, 0xc2, 0x01, 0x0a, 0x13, 0x53, 0x53, 0x52, 0x46, 0x52,
	0x65, 0x67, 0x69, 0x73, 0x74, 0x72, 0x79, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x12, 0x53,
	0x0a, 0x06, 0x41, 0x73, 0x73, 0x69, 0x67, 0x6e, 0x12, 0x15, 0x2e, 0x70, 0x62, 0x2e, 0x53, 0x53,
	0x52, 0x46, 0x41, 0x73, 0x73, 0x69, 0x67, 0x6e, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a,
	0x16, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75,
	0x66, 0x2e, 0x45, 0x6d, 0x70, 0x74, 0x79, 0x22, 0x1a, 0x82, 0xd3, 0xe4, 0x93, 0x02, 0x14, 0x22,
	0x0f, 0x2f, 0x76, 0x31, 0x2f, 0x73, 0x73, 0x72, 0x66, 0x2f, 0x61, 0x73, 0x73, 0x69, 0x67, 0x6e,
	0x3a, 0x01, 0x2a, 0x12, 0x56, 0x0a, 0x07, 0x52, 0x65, 0x6c, 0x65, 0x61, 0x73, 0x65, 0x12, 0x16,
	0x2e, 0x70, 0x62, 0x2e, 0x53, 0x53, 0x52, 0x46, 0x52, 0x65, 0x6c, 0x65, 0x61, 0x73, 0x65, 0x52,
	0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x16, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x45, 0x6d, 0x70, 0x74, 0x79, 0x22, 0x1b,
	0x82, 0xd3, 0xe4, 0x93, 0x02, 0x15, 0x22, 0x10, 0x2f, 0x76, 0x31, 0x2f, 0x73, 0x73, 0x72, 0x66,
	0x2f, 0x72, 0x65, 0x6c, 0x65, 0x61, 0x73, 0x65, 0x3a, 0x01, 0x2a, 0x42, 0x06, 0x5a, 0x04, 0x2e,
	0x3b, 0x70, 0x62, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_ssrf_proto_rawDescOnce sync.Once
	file_ssrf_proto_rawDescData = file_ssrf_proto_rawDesc
)

func file_ssrf_proto_rawDescGZIP() []byte {
	file_ssrf_proto_rawDescOnce.Do(func() {
		file_ssrf_proto_rawDescData = protoimpl.X.CompressGZIP(file_ssrf_proto_rawDescData)
	})
	return file_ssrf_proto_rawDescData
}

var file_ssrf_proto_msgTypes = make([]protoimpl.MessageInfo, 2)
var file_ssrf_proto_goTypes = []interface{}{
	(*SSRFAssignRequest)(nil),  // 0: pb.SSRFAssignRequest
	(*SSRFReleaseRequest)(nil), // 1: pb.SSRFReleaseRequest
	(*empty.Empty)(nil),        // 2: google.protobuf.Empty
}
var file_ssrf_proto_depIdxs = []int32{
	0, // 0: pb.SSRFRegistryService.Assign:input_type -> pb.SSRFAssignRequest
	1, // 1: pb.SSRFRegistryService.Release:input_type -> pb.SSRFReleaseRequest
	2, // 2: pb.SSRFRegistryService.Assign:output_type -> google.protobuf.Empty
	2, // 3: pb.SSRFRegistryService.Release:output_type -> google.protobuf.Empty
	2, // [2:4] is the sub-list for method output_type
	0, // [0:2] is the sub-list for method input_type
	0, // [0:0] is the sub-list for extension type_name
	0, // [0:0] is the sub-list for extension extendee
	0, // [0:0] is the sub-list for field type_name
}

func init() { file_ssrf_proto_init() }
func file_ssrf_proto_init() {
	if File_ssrf_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_ssrf_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*SSRFAssignRequest); i {
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
		file_ssrf_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*SSRFReleaseRequest); i {
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
			RawDescriptor: file_ssrf_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   2,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_ssrf_proto_goTypes,
		DependencyIndexes: file_ssrf_proto_depIdxs,
		MessageInfos:      file_ssrf_proto_msgTypes,
	}.Build()
	File_ssrf_proto = out.File
	file_ssrf_proto_rawDesc = nil
	file_ssrf_proto_goTypes = nil
	file_ssrf_proto_depIdxs = nil
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConnInterface

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion6

// SSRFRegistryServiceClient is the client API for SSRFRegistryService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type SSRFRegistryServiceClient interface {
	Assign(ctx context.Context, in *SSRFAssignRequest, opts ...grpc.CallOption) (*empty.Empty, error)
	Release(ctx context.Context, in *SSRFReleaseRequest, opts ...grpc.CallOption) (*empty.Empty, error)
}

type sSRFRegistryServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewSSRFRegistryServiceClient(cc grpc.ClientConnInterface) SSRFRegistryServiceClient {
	return &sSRFRegistryServiceClient{cc}
}

func (c *sSRFRegistryServiceClient) Assign(ctx context.Context, in *SSRFAssignRequest, opts ...grpc.CallOption) (*empty.Empty, error) {
	out := new(empty.Empty)
	err := c.cc.Invoke(ctx, "/pb.SSRFRegistryService/Assign", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *sSRFRegistryServiceClient) Release(ctx context.Context, in *SSRFReleaseRequest, opts ...grpc.CallOption) (*empty.Empty, error) {
	out := new(empty.Empty)
	err := c.cc.Invoke(ctx, "/pb.SSRFRegistryService/Release", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// SSRFRegistryServiceServer is the server API for SSRFRegistryService service.
type SSRFRegistryServiceServer interface {
	Assign(context.Context, *SSRFAssignRequest) (*empty.Empty, error)
	Release(context.Context, *SSRFReleaseRequest) (*empty.Empty, error)
}

// UnimplementedSSRFRegistryServiceServer can be embedded to have forward compatible implementations.
type UnimplementedSSRFRegistryServiceServer struct {
}

func (*UnimplementedSSRFRegistryServiceServer) Assign(context.Context, *SSRFAssignRequest) (*empty.Empty, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Assign not implemented")
}
func (*UnimplementedSSRFRegistryServiceServer) Release(context.Context, *SSRFReleaseRequest) (*empty.Empty, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Release not implemented")
}

func RegisterSSRFRegistryServiceServer(s *grpc.Server, srv SSRFRegistryServiceServer) {
	s.RegisterService(&_SSRFRegistryService_serviceDesc, srv)
}

func _SSRFRegistryService_Assign_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(SSRFAssignRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(SSRFRegistryServiceServer).Assign(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/pb.SSRFRegistryService/Assign",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(SSRFRegistryServiceServer).Assign(ctx, req.(*SSRFAssignRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _SSRFRegistryService_Release_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(SSRFReleaseRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(SSRFRegistryServiceServer).Release(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/pb.SSRFRegistryService/Release",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(SSRFRegistryServiceServer).Release(ctx, req.(*SSRFReleaseRequest))
	}
	return interceptor(ctx, in, info, handler)
}

var _SSRFRegistryService_serviceDesc = grpc.ServiceDesc{
	ServiceName: "pb.SSRFRegistryService",
	HandlerType: (*SSRFRegistryServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "Assign",
			Handler:    _SSRFRegistryService_Assign_Handler,
		},
		{
			MethodName: "Release",
			Handler:    _SSRFRegistryService_Release_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "ssrf.proto",
}
