// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.26.0
// 	protoc        v3.21.2
// source: v1/health/health.proto

package health

import (
	context "context"
	grpc "google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
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

type Response struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Success bool   `protobuf:"varint,1,opt,name=success,proto3" json:"success,omitempty"`
	Code    string `protobuf:"bytes,2,opt,name=code,proto3" json:"code,omitempty"`
	Desc    string `protobuf:"bytes,3,opt,name=desc,proto3" json:"desc,omitempty"`
}

func (x *Response) Reset() {
	*x = Response{}
	if protoimpl.UnsafeEnabled {
		mi := &file_v1_health_health_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Response) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Response) ProtoMessage() {}

func (x *Response) ProtoReflect() protoreflect.Message {
	mi := &file_v1_health_health_proto_msgTypes[0]
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
	return file_v1_health_health_proto_rawDescGZIP(), []int{0}
}

func (x *Response) GetSuccess() bool {
	if x != nil {
		return x.Success
	}
	return false
}

func (x *Response) GetCode() string {
	if x != nil {
		return x.Code
	}
	return ""
}

func (x *Response) GetDesc() string {
	if x != nil {
		return x.Desc
	}
	return ""
}

var File_v1_health_health_proto protoreflect.FileDescriptor

var file_v1_health_health_proto_rawDesc = []byte{
	0x0a, 0x16, 0x76, 0x31, 0x2f, 0x68, 0x65, 0x61, 0x6c, 0x74, 0x68, 0x2f, 0x68, 0x65, 0x61, 0x6c,
	0x74, 0x68, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x14, 0x61, 0x70, 0x69, 0x2e, 0x67, 0x6f,
	0x67, 0x72, 0x70, 0x63, 0x2e, 0x76, 0x31, 0x2e, 0x68, 0x65, 0x61, 0x6c, 0x74, 0x68, 0x1a, 0x1b,
	0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2f,
	0x65, 0x6d, 0x70, 0x74, 0x79, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0x4c, 0x0a, 0x08, 0x52,
	0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x18, 0x0a, 0x07, 0x73, 0x75, 0x63, 0x63, 0x65,
	0x73, 0x73, 0x18, 0x01, 0x20, 0x01, 0x28, 0x08, 0x52, 0x07, 0x73, 0x75, 0x63, 0x63, 0x65, 0x73,
	0x73, 0x12, 0x12, 0x0a, 0x04, 0x63, 0x6f, 0x64, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52,
	0x04, 0x63, 0x6f, 0x64, 0x65, 0x12, 0x12, 0x0a, 0x04, 0x64, 0x65, 0x73, 0x63, 0x18, 0x03, 0x20,
	0x01, 0x28, 0x09, 0x52, 0x04, 0x64, 0x65, 0x73, 0x63, 0x32, 0x53, 0x0a, 0x0d, 0x48, 0x65, 0x61,
	0x6c, 0x74, 0x68, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x12, 0x42, 0x0a, 0x06, 0x53, 0x74,
	0x61, 0x74, 0x75, 0x73, 0x12, 0x16, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x45, 0x6d, 0x70, 0x74, 0x79, 0x1a, 0x1e, 0x2e, 0x61,
	0x70, 0x69, 0x2e, 0x67, 0x6f, 0x67, 0x72, 0x70, 0x63, 0x2e, 0x76, 0x31, 0x2e, 0x68, 0x65, 0x61,
	0x6c, 0x74, 0x68, 0x2e, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x00, 0x42, 0x37,
	0x5a, 0x35, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x61, 0x6a, 0x69,
	0x6b, 0x61, 0x6d, 0x61, 0x6c, 0x75, 0x64, 0x69, 0x6e, 0x2f, 0x67, 0x6f, 0x2d, 0x67, 0x72, 0x70,
	0x63, 0x5f, 0x62, 0x61, 0x73, 0x69, 0x63, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2f, 0x76, 0x31,
	0x2f, 0x68, 0x65, 0x61, 0x6c, 0x74, 0x68, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_v1_health_health_proto_rawDescOnce sync.Once
	file_v1_health_health_proto_rawDescData = file_v1_health_health_proto_rawDesc
)

func file_v1_health_health_proto_rawDescGZIP() []byte {
	file_v1_health_health_proto_rawDescOnce.Do(func() {
		file_v1_health_health_proto_rawDescData = protoimpl.X.CompressGZIP(file_v1_health_health_proto_rawDescData)
	})
	return file_v1_health_health_proto_rawDescData
}

var file_v1_health_health_proto_msgTypes = make([]protoimpl.MessageInfo, 1)
var file_v1_health_health_proto_goTypes = []interface{}{
	(*Response)(nil),      // 0: api.gogrpc.v1.health.Response
	(*emptypb.Empty)(nil), // 1: google.protobuf.Empty
}
var file_v1_health_health_proto_depIdxs = []int32{
	1, // 0: api.gogrpc.v1.health.HealthService.Status:input_type -> google.protobuf.Empty
	0, // 1: api.gogrpc.v1.health.HealthService.Status:output_type -> api.gogrpc.v1.health.Response
	1, // [1:2] is the sub-list for method output_type
	0, // [0:1] is the sub-list for method input_type
	0, // [0:0] is the sub-list for extension type_name
	0, // [0:0] is the sub-list for extension extendee
	0, // [0:0] is the sub-list for field type_name
}

func init() { file_v1_health_health_proto_init() }
func file_v1_health_health_proto_init() {
	if File_v1_health_health_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_v1_health_health_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
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
	}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_v1_health_health_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   1,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_v1_health_health_proto_goTypes,
		DependencyIndexes: file_v1_health_health_proto_depIdxs,
		MessageInfos:      file_v1_health_health_proto_msgTypes,
	}.Build()
	File_v1_health_health_proto = out.File
	file_v1_health_health_proto_rawDesc = nil
	file_v1_health_health_proto_goTypes = nil
	file_v1_health_health_proto_depIdxs = nil
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConnInterface

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion6

// HealthServiceClient is the client API for HealthService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type HealthServiceClient interface {
	Status(ctx context.Context, in *emptypb.Empty, opts ...grpc.CallOption) (*Response, error)
}

type healthServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewHealthServiceClient(cc grpc.ClientConnInterface) HealthServiceClient {
	return &healthServiceClient{cc}
}

func (c *healthServiceClient) Status(ctx context.Context, in *emptypb.Empty, opts ...grpc.CallOption) (*Response, error) {
	out := new(Response)
	err := c.cc.Invoke(ctx, "/api.gogrpc.v1.health.HealthService/Status", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// HealthServiceServer is the server API for HealthService service.
type HealthServiceServer interface {
	Status(context.Context, *emptypb.Empty) (*Response, error)
}

// UnimplementedHealthServiceServer can be embedded to have forward compatible implementations.
type UnimplementedHealthServiceServer struct {
}

func (*UnimplementedHealthServiceServer) Status(context.Context, *emptypb.Empty) (*Response, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Status not implemented")
}

func RegisterHealthServiceServer(s *grpc.Server, srv HealthServiceServer) {
	s.RegisterService(&_HealthService_serviceDesc, srv)
}

func _HealthService_Status_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(emptypb.Empty)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(HealthServiceServer).Status(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/api.gogrpc.v1.health.HealthService/Status",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(HealthServiceServer).Status(ctx, req.(*emptypb.Empty))
	}
	return interceptor(ctx, in, info, handler)
}

var _HealthService_serviceDesc = grpc.ServiceDesc{
	ServiceName: "api.gogrpc.v1.health.HealthService",
	HandlerType: (*HealthServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "Status",
			Handler:    _HealthService_Status_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "v1/health/health.proto",
}
