// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.25.0
// 	protoc        v3.13.0
// source: prod.proto

package services

import (
	context "context"
	proto "github.com/golang/protobuf/proto"
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

type ProdRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	ProdId int32 `protobuf:"varint,1,opt,name=prodId,proto3" json:"prodId,omitempty"`
}

func (x *ProdRequest) Reset() {
	*x = ProdRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_prod_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ProdRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ProdRequest) ProtoMessage() {}

func (x *ProdRequest) ProtoReflect() protoreflect.Message {
	mi := &file_prod_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ProdRequest.ProtoReflect.Descriptor instead.
func (*ProdRequest) Descriptor() ([]byte, []int) {
	return file_prod_proto_rawDescGZIP(), []int{0}
}

func (x *ProdRequest) GetProdId() int32 {
	if x != nil {
		return x.ProdId
	}
	return 0
}

type ProdResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	ProdStock int32 `protobuf:"varint,1,opt,name=prodStock,proto3" json:"prodStock,omitempty"`
}

func (x *ProdResponse) Reset() {
	*x = ProdResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_prod_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ProdResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ProdResponse) ProtoMessage() {}

func (x *ProdResponse) ProtoReflect() protoreflect.Message {
	mi := &file_prod_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ProdResponse.ProtoReflect.Descriptor instead.
func (*ProdResponse) Descriptor() ([]byte, []int) {
	return file_prod_proto_rawDescGZIP(), []int{1}
}

func (x *ProdResponse) GetProdStock() int32 {
	if x != nil {
		return x.ProdStock
	}
	return 0
}

type QuerySize struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Size int32 `protobuf:"varint,1,opt,name=size,proto3" json:"size,omitempty"` //页码尺寸
}

func (x *QuerySize) Reset() {
	*x = QuerySize{}
	if protoimpl.UnsafeEnabled {
		mi := &file_prod_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *QuerySize) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*QuerySize) ProtoMessage() {}

func (x *QuerySize) ProtoReflect() protoreflect.Message {
	mi := &file_prod_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use QuerySize.ProtoReflect.Descriptor instead.
func (*QuerySize) Descriptor() ([]byte, []int) {
	return file_prod_proto_rawDescGZIP(), []int{2}
}

func (x *QuerySize) GetSize() int32 {
	if x != nil {
		return x.Size
	}
	return 0
}

type ProdRespList struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	ProdList []*ProdResponse `protobuf:"bytes,1,rep,name=prodList,proto3" json:"prodList,omitempty"` //修饰符
}

func (x *ProdRespList) Reset() {
	*x = ProdRespList{}
	if protoimpl.UnsafeEnabled {
		mi := &file_prod_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ProdRespList) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ProdRespList) ProtoMessage() {}

func (x *ProdRespList) ProtoReflect() protoreflect.Message {
	mi := &file_prod_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ProdRespList.ProtoReflect.Descriptor instead.
func (*ProdRespList) Descriptor() ([]byte, []int) {
	return file_prod_proto_rawDescGZIP(), []int{3}
}

func (x *ProdRespList) GetProdList() []*ProdResponse {
	if x != nil {
		return x.ProdList
	}
	return nil
}

var File_prod_proto protoreflect.FileDescriptor

var file_prod_proto_rawDesc = []byte{
	0x0a, 0x0a, 0x70, 0x72, 0x6f, 0x64, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x08, 0x73, 0x65,
	0x72, 0x76, 0x69, 0x63, 0x65, 0x73, 0x1a, 0x1c, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x61,
	0x70, 0x69, 0x2f, 0x61, 0x6e, 0x6e, 0x6f, 0x74, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x2e, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x22, 0x25, 0x0a, 0x0b, 0x50, 0x72, 0x6f, 0x64, 0x52, 0x65, 0x71, 0x75,
	0x65, 0x73, 0x74, 0x12, 0x16, 0x0a, 0x06, 0x70, 0x72, 0x6f, 0x64, 0x49, 0x64, 0x18, 0x01, 0x20,
	0x01, 0x28, 0x05, 0x52, 0x06, 0x70, 0x72, 0x6f, 0x64, 0x49, 0x64, 0x22, 0x2c, 0x0a, 0x0c, 0x50,
	0x72, 0x6f, 0x64, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x1c, 0x0a, 0x09, 0x70,
	0x72, 0x6f, 0x64, 0x53, 0x74, 0x6f, 0x63, 0x6b, 0x18, 0x01, 0x20, 0x01, 0x28, 0x05, 0x52, 0x09,
	0x70, 0x72, 0x6f, 0x64, 0x53, 0x74, 0x6f, 0x63, 0x6b, 0x22, 0x1f, 0x0a, 0x09, 0x51, 0x75, 0x65,
	0x72, 0x79, 0x53, 0x69, 0x7a, 0x65, 0x12, 0x12, 0x0a, 0x04, 0x73, 0x69, 0x7a, 0x65, 0x18, 0x01,
	0x20, 0x01, 0x28, 0x05, 0x52, 0x04, 0x73, 0x69, 0x7a, 0x65, 0x22, 0x42, 0x0a, 0x0c, 0x50, 0x72,
	0x6f, 0x64, 0x52, 0x65, 0x73, 0x70, 0x4c, 0x69, 0x73, 0x74, 0x12, 0x32, 0x0a, 0x08, 0x70, 0x72,
	0x6f, 0x64, 0x4c, 0x69, 0x73, 0x74, 0x18, 0x01, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x16, 0x2e, 0x73,
	0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x73, 0x2e, 0x50, 0x72, 0x6f, 0x64, 0x52, 0x65, 0x73, 0x70,
	0x6f, 0x6e, 0x73, 0x65, 0x52, 0x08, 0x70, 0x72, 0x6f, 0x64, 0x4c, 0x69, 0x73, 0x74, 0x32, 0xa7,
	0x01, 0x0a, 0x0b, 0x50, 0x72, 0x6f, 0x64, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x12, 0x58,
	0x0a, 0x0c, 0x47, 0x65, 0x74, 0x50, 0x72, 0x6f, 0x64, 0x53, 0x74, 0x6f, 0x63, 0x6b, 0x12, 0x15,
	0x2e, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x73, 0x2e, 0x50, 0x72, 0x6f, 0x64, 0x52, 0x65,
	0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x16, 0x2e, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x73,
	0x2e, 0x50, 0x72, 0x6f, 0x64, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x19, 0x82,
	0xd3, 0xe4, 0x93, 0x02, 0x13, 0x12, 0x11, 0x2f, 0x76, 0x31, 0x2f, 0x70, 0x72, 0x6f, 0x64, 0x2f,
	0x7b, 0x70, 0x72, 0x6f, 0x64, 0x49, 0x64, 0x7d, 0x12, 0x3e, 0x0a, 0x0d, 0x47, 0x65, 0x74, 0x50,
	0x72, 0x6f, 0x64, 0x53, 0x74, 0x6f, 0x63, 0x6b, 0x73, 0x12, 0x13, 0x2e, 0x73, 0x65, 0x72, 0x76,
	0x69, 0x63, 0x65, 0x73, 0x2e, 0x51, 0x75, 0x65, 0x72, 0x79, 0x53, 0x69, 0x7a, 0x65, 0x1a, 0x16,
	0x2e, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x73, 0x2e, 0x50, 0x72, 0x6f, 0x64, 0x52, 0x65,
	0x73, 0x70, 0x4c, 0x69, 0x73, 0x74, 0x22, 0x00, 0x42, 0x0c, 0x5a, 0x0a, 0x2e, 0x2e, 0x2f, 0x70,
	0x62, 0x66, 0x69, 0x6c, 0x65, 0x73, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_prod_proto_rawDescOnce sync.Once
	file_prod_proto_rawDescData = file_prod_proto_rawDesc
)

func file_prod_proto_rawDescGZIP() []byte {
	file_prod_proto_rawDescOnce.Do(func() {
		file_prod_proto_rawDescData = protoimpl.X.CompressGZIP(file_prod_proto_rawDescData)
	})
	return file_prod_proto_rawDescData
}

var file_prod_proto_msgTypes = make([]protoimpl.MessageInfo, 4)
var file_prod_proto_goTypes = []interface{}{
	(*ProdRequest)(nil),  // 0: services.ProdRequest
	(*ProdResponse)(nil), // 1: services.ProdResponse
	(*QuerySize)(nil),    // 2: services.QuerySize
	(*ProdRespList)(nil), // 3: services.ProdRespList
}
var file_prod_proto_depIdxs = []int32{
	1, // 0: services.ProdRespList.prodList:type_name -> services.ProdResponse
	0, // 1: services.ProdService.GetProdStock:input_type -> services.ProdRequest
	2, // 2: services.ProdService.GetProdStocks:input_type -> services.QuerySize
	1, // 3: services.ProdService.GetProdStock:output_type -> services.ProdResponse
	3, // 4: services.ProdService.GetProdStocks:output_type -> services.ProdRespList
	3, // [3:5] is the sub-list for method output_type
	1, // [1:3] is the sub-list for method input_type
	1, // [1:1] is the sub-list for extension type_name
	1, // [1:1] is the sub-list for extension extendee
	0, // [0:1] is the sub-list for field type_name
}

func init() { file_prod_proto_init() }
func file_prod_proto_init() {
	if File_prod_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_prod_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ProdRequest); i {
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
		file_prod_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ProdResponse); i {
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
		file_prod_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*QuerySize); i {
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
		file_prod_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ProdRespList); i {
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
			RawDescriptor: file_prod_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   4,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_prod_proto_goTypes,
		DependencyIndexes: file_prod_proto_depIdxs,
		MessageInfos:      file_prod_proto_msgTypes,
	}.Build()
	File_prod_proto = out.File
	file_prod_proto_rawDesc = nil
	file_prod_proto_goTypes = nil
	file_prod_proto_depIdxs = nil
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConnInterface

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion6

// ProdServiceClient is the client API for ProdService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type ProdServiceClient interface {
	GetProdStock(ctx context.Context, in *ProdRequest, opts ...grpc.CallOption) (*ProdResponse, error)
	GetProdStocks(ctx context.Context, in *QuerySize, opts ...grpc.CallOption) (*ProdRespList, error)
}

type prodServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewProdServiceClient(cc grpc.ClientConnInterface) ProdServiceClient {
	return &prodServiceClient{cc}
}

func (c *prodServiceClient) GetProdStock(ctx context.Context, in *ProdRequest, opts ...grpc.CallOption) (*ProdResponse, error) {
	out := new(ProdResponse)
	err := c.cc.Invoke(ctx, "/services.ProdService/GetProdStock", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *prodServiceClient) GetProdStocks(ctx context.Context, in *QuerySize, opts ...grpc.CallOption) (*ProdRespList, error) {
	out := new(ProdRespList)
	err := c.cc.Invoke(ctx, "/services.ProdService/GetProdStocks", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// ProdServiceServer is the server API for ProdService service.
type ProdServiceServer interface {
	GetProdStock(context.Context, *ProdRequest) (*ProdResponse, error)
	GetProdStocks(context.Context, *QuerySize) (*ProdRespList, error)
}

// UnimplementedProdServiceServer can be embedded to have forward compatible implementations.
type UnimplementedProdServiceServer struct {
}

func (*UnimplementedProdServiceServer) GetProdStock(context.Context, *ProdRequest) (*ProdResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetProdStock not implemented")
}
func (*UnimplementedProdServiceServer) GetProdStocks(context.Context, *QuerySize) (*ProdRespList, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetProdStocks not implemented")
}

func RegisterProdServiceServer(s *grpc.Server, srv ProdServiceServer) {
	s.RegisterService(&_ProdService_serviceDesc, srv)
}

func _ProdService_GetProdStock_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ProdRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ProdServiceServer).GetProdStock(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/services.ProdService/GetProdStock",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ProdServiceServer).GetProdStock(ctx, req.(*ProdRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _ProdService_GetProdStocks_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(QuerySize)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ProdServiceServer).GetProdStocks(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/services.ProdService/GetProdStocks",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ProdServiceServer).GetProdStocks(ctx, req.(*QuerySize))
	}
	return interceptor(ctx, in, info, handler)
}

var _ProdService_serviceDesc = grpc.ServiceDesc{
	ServiceName: "services.ProdService",
	HandlerType: (*ProdServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "GetProdStock",
			Handler:    _ProdService_GetProdStock_Handler,
		},
		{
			MethodName: "GetProdStocks",
			Handler:    _ProdService_GetProdStocks_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "prod.proto",
}
