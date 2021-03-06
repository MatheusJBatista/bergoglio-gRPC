// Code generated by protoc-gen-go. DO NOT EDIT.
// source: proto/code.proto

package bergogliov1

import (
	context "context"
	fmt "fmt"
	proto "github.com/golang/protobuf/proto"
	grpc "google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
	math "math"
)

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// This is a compile-time assertion to ensure that this generated file
// is compatible with the proto package it is being compiled against.
// A compilation error at this line likely means your copy of the
// proto package needs to be updated.
const _ = proto.ProtoPackageIsVersion3 // please upgrade the proto package

type Entry struct {
	Name                 string   `protobuf:"bytes,1,opt,name=name,proto3" json:"name,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *Entry) Reset()         { *m = Entry{} }
func (m *Entry) String() string { return proto.CompactTextString(m) }
func (*Entry) ProtoMessage()    {}
func (*Entry) Descriptor() ([]byte, []int) {
	return fileDescriptor_4153a92670452c72, []int{0}
}

func (m *Entry) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Entry.Unmarshal(m, b)
}
func (m *Entry) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Entry.Marshal(b, m, deterministic)
}
func (m *Entry) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Entry.Merge(m, src)
}
func (m *Entry) XXX_Size() int {
	return xxx_messageInfo_Entry.Size(m)
}
func (m *Entry) XXX_DiscardUnknown() {
	xxx_messageInfo_Entry.DiscardUnknown(m)
}

var xxx_messageInfo_Entry proto.InternalMessageInfo

func (m *Entry) GetName() string {
	if m != nil {
		return m.Name
	}
	return ""
}

type End struct {
	Code                 string   `protobuf:"bytes,1,opt,name=code,proto3" json:"code,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *End) Reset()         { *m = End{} }
func (m *End) String() string { return proto.CompactTextString(m) }
func (*End) ProtoMessage()    {}
func (*End) Descriptor() ([]byte, []int) {
	return fileDescriptor_4153a92670452c72, []int{1}
}

func (m *End) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_End.Unmarshal(m, b)
}
func (m *End) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_End.Marshal(b, m, deterministic)
}
func (m *End) XXX_Merge(src proto.Message) {
	xxx_messageInfo_End.Merge(m, src)
}
func (m *End) XXX_Size() int {
	return xxx_messageInfo_End.Size(m)
}
func (m *End) XXX_DiscardUnknown() {
	xxx_messageInfo_End.DiscardUnknown(m)
}

var xxx_messageInfo_End proto.InternalMessageInfo

func (m *End) GetCode() string {
	if m != nil {
		return m.Code
	}
	return ""
}

type ManyCodesRequest struct {
	SerialNumberInit     int32    `protobuf:"varint,1,opt,name=serialNumberInit,proto3" json:"serialNumberInit,omitempty"`
	SerialNumberFinal    int32    `protobuf:"varint,2,opt,name=serialNumberFinal,proto3" json:"serialNumberFinal,omitempty"`
	QuantityPerSerie     int32    `protobuf:"varint,3,opt,name=quantityPerSerie,proto3" json:"quantityPerSerie,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *ManyCodesRequest) Reset()         { *m = ManyCodesRequest{} }
func (m *ManyCodesRequest) String() string { return proto.CompactTextString(m) }
func (*ManyCodesRequest) ProtoMessage()    {}
func (*ManyCodesRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_4153a92670452c72, []int{2}
}

func (m *ManyCodesRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_ManyCodesRequest.Unmarshal(m, b)
}
func (m *ManyCodesRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_ManyCodesRequest.Marshal(b, m, deterministic)
}
func (m *ManyCodesRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ManyCodesRequest.Merge(m, src)
}
func (m *ManyCodesRequest) XXX_Size() int {
	return xxx_messageInfo_ManyCodesRequest.Size(m)
}
func (m *ManyCodesRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_ManyCodesRequest.DiscardUnknown(m)
}

var xxx_messageInfo_ManyCodesRequest proto.InternalMessageInfo

func (m *ManyCodesRequest) GetSerialNumberInit() int32 {
	if m != nil {
		return m.SerialNumberInit
	}
	return 0
}

func (m *ManyCodesRequest) GetSerialNumberFinal() int32 {
	if m != nil {
		return m.SerialNumberFinal
	}
	return 0
}

func (m *ManyCodesRequest) GetQuantityPerSerie() int32 {
	if m != nil {
		return m.QuantityPerSerie
	}
	return 0
}

type ManyCodesResponse struct {
	Codes                []string `protobuf:"bytes,1,rep,name=codes,proto3" json:"codes,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *ManyCodesResponse) Reset()         { *m = ManyCodesResponse{} }
func (m *ManyCodesResponse) String() string { return proto.CompactTextString(m) }
func (*ManyCodesResponse) ProtoMessage()    {}
func (*ManyCodesResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_4153a92670452c72, []int{3}
}

func (m *ManyCodesResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_ManyCodesResponse.Unmarshal(m, b)
}
func (m *ManyCodesResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_ManyCodesResponse.Marshal(b, m, deterministic)
}
func (m *ManyCodesResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ManyCodesResponse.Merge(m, src)
}
func (m *ManyCodesResponse) XXX_Size() int {
	return xxx_messageInfo_ManyCodesResponse.Size(m)
}
func (m *ManyCodesResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_ManyCodesResponse.DiscardUnknown(m)
}

var xxx_messageInfo_ManyCodesResponse proto.InternalMessageInfo

func (m *ManyCodesResponse) GetCodes() []string {
	if m != nil {
		return m.Codes
	}
	return nil
}

func init() {
	proto.RegisterType((*Entry)(nil), "bergogliov1.Entry")
	proto.RegisterType((*End)(nil), "bergogliov1.End")
	proto.RegisterType((*ManyCodesRequest)(nil), "bergogliov1.ManyCodesRequest")
	proto.RegisterType((*ManyCodesResponse)(nil), "bergogliov1.ManyCodesResponse")
}

func init() { proto.RegisterFile("proto/code.proto", fileDescriptor_4153a92670452c72) }

var fileDescriptor_4153a92670452c72 = []byte{
	// 293 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xa4, 0x52, 0x4d, 0x4b, 0xc3, 0x40,
	0x14, 0xec, 0x5a, 0x2b, 0xf4, 0xe9, 0x21, 0x7d, 0x88, 0xc4, 0x8a, 0x52, 0xf6, 0x54, 0x45, 0xea,
	0xd7, 0x1f, 0x90, 0x4a, 0x2d, 0x1e, 0xb4, 0x92, 0x1e, 0x7a, 0xde, 0x9a, 0x47, 0x5d, 0x48, 0x77,
	0xdb, 0xdd, 0x8d, 0x90, 0x7f, 0xe2, 0x3f, 0xf5, 0x2a, 0xbb, 0xa1, 0x12, 0x53, 0x3c, 0x79, 0x9b,
	0xcc, 0xbc, 0x99, 0x0c, 0xc3, 0x42, 0xb4, 0x32, 0xda, 0xe9, 0xab, 0x37, 0x9d, 0xd2, 0x20, 0x40,
	0xdc, 0x9f, 0x93, 0x59, 0xe8, 0x45, 0x26, 0xf5, 0xc7, 0x0d, 0x3f, 0x81, 0xd6, 0x48, 0x39, 0x53,
	0x20, 0xc2, 0xae, 0x12, 0x4b, 0x8a, 0x59, 0x8f, 0xf5, 0xdb, 0x49, 0xc0, 0xfc, 0x18, 0x9a, 0x23,
	0x95, 0x7a, 0xc9, 0xdb, 0x37, 0x92, 0xc7, 0xfc, 0x93, 0x41, 0xf4, 0x2c, 0x54, 0xf1, 0xa0, 0x53,
	0xb2, 0x09, 0xad, 0x73, 0xb2, 0x0e, 0x2f, 0x20, 0xb2, 0x64, 0xa4, 0xc8, 0x5e, 0xf2, 0xe5, 0x9c,
	0xcc, 0x93, 0x92, 0x2e, 0x98, 0x5a, 0xc9, 0x16, 0x8f, 0x97, 0xd0, 0xa9, 0x72, 0x8f, 0x52, 0x89,
	0x2c, 0xde, 0x09, 0xc7, 0xdb, 0x82, 0x4f, 0x5e, 0xe7, 0x42, 0x39, 0xe9, 0x8a, 0x57, 0x32, 0x53,
	0x32, 0x92, 0xe2, 0x66, 0x99, 0x5c, 0xe7, 0xf9, 0x39, 0x74, 0x2a, 0xcd, 0xec, 0x4a, 0x2b, 0x4b,
	0x78, 0x08, 0x2d, 0xdf, 0xdb, 0xc6, 0xac, 0xd7, 0xec, 0xb7, 0x93, 0xf2, 0xe3, 0xf6, 0x8b, 0x41,
	0x7b, 0xb8, 0x59, 0x03, 0xef, 0xe1, 0x68, 0x4c, 0x6e, 0xa2, 0xc8, 0x5b, 0x67, 0xd2, 0xbd, 0x0f,
	0x33, 0xb2, 0x56, 0xaa, 0x05, 0xe2, 0xa0, 0xb2, 0xd9, 0x20, 0x0c, 0xd6, 0x8d, 0x6a, 0x5c, 0xca,
	0x1b, 0x38, 0x81, 0x83, 0x31, 0xb9, 0x9f, 0xbf, 0xe3, 0xe9, 0xaf, 0x9b, 0xfa, 0x5e, 0xdd, 0xb3,
	0xbf, 0xe4, 0xb2, 0x34, 0x6f, 0xe0, 0x0c, 0xb0, 0x1a, 0x38, 0x75, 0x86, 0xc4, 0xf2, 0xdf, 0xb1,
	0xd7, 0x6c, 0xbe, 0x17, 0xde, 0xc2, 0xdd, 0x77, 0x00, 0x00, 0x00, 0xff, 0xff, 0xf5, 0x83, 0x71,
	0x2b, 0x1f, 0x02, 0x00, 0x00,
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConn

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion4

// BergoglioClient is the client API for Bergoglio service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type BergoglioClient interface {
	GetOneCodeWithBlessing(ctx context.Context, in *Entry, opts ...grpc.CallOption) (*End, error)
	GetManyCodes(ctx context.Context, in *ManyCodesRequest, opts ...grpc.CallOption) (*ManyCodesResponse, error)
	GetManyCodesStream(ctx context.Context, in *ManyCodesRequest, opts ...grpc.CallOption) (Bergoglio_GetManyCodesStreamClient, error)
}

type bergoglioClient struct {
	cc *grpc.ClientConn
}

func NewBergoglioClient(cc *grpc.ClientConn) BergoglioClient {
	return &bergoglioClient{cc}
}

func (c *bergoglioClient) GetOneCodeWithBlessing(ctx context.Context, in *Entry, opts ...grpc.CallOption) (*End, error) {
	out := new(End)
	err := c.cc.Invoke(ctx, "/bergogliov1.Bergoglio/GetOneCodeWithBlessing", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *bergoglioClient) GetManyCodes(ctx context.Context, in *ManyCodesRequest, opts ...grpc.CallOption) (*ManyCodesResponse, error) {
	out := new(ManyCodesResponse)
	err := c.cc.Invoke(ctx, "/bergogliov1.Bergoglio/GetManyCodes", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *bergoglioClient) GetManyCodesStream(ctx context.Context, in *ManyCodesRequest, opts ...grpc.CallOption) (Bergoglio_GetManyCodesStreamClient, error) {
	stream, err := c.cc.NewStream(ctx, &_Bergoglio_serviceDesc.Streams[0], "/bergogliov1.Bergoglio/GetManyCodesStream", opts...)
	if err != nil {
		return nil, err
	}
	x := &bergoglioGetManyCodesStreamClient{stream}
	if err := x.ClientStream.SendMsg(in); err != nil {
		return nil, err
	}
	if err := x.ClientStream.CloseSend(); err != nil {
		return nil, err
	}
	return x, nil
}

type Bergoglio_GetManyCodesStreamClient interface {
	Recv() (*ManyCodesResponse, error)
	grpc.ClientStream
}

type bergoglioGetManyCodesStreamClient struct {
	grpc.ClientStream
}

func (x *bergoglioGetManyCodesStreamClient) Recv() (*ManyCodesResponse, error) {
	m := new(ManyCodesResponse)
	if err := x.ClientStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

// BergoglioServer is the server API for Bergoglio service.
type BergoglioServer interface {
	GetOneCodeWithBlessing(context.Context, *Entry) (*End, error)
	GetManyCodes(context.Context, *ManyCodesRequest) (*ManyCodesResponse, error)
	GetManyCodesStream(*ManyCodesRequest, Bergoglio_GetManyCodesStreamServer) error
}

// UnimplementedBergoglioServer can be embedded to have forward compatible implementations.
type UnimplementedBergoglioServer struct {
}

func (*UnimplementedBergoglioServer) GetOneCodeWithBlessing(ctx context.Context, req *Entry) (*End, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetOneCodeWithBlessing not implemented")
}
func (*UnimplementedBergoglioServer) GetManyCodes(ctx context.Context, req *ManyCodesRequest) (*ManyCodesResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetManyCodes not implemented")
}
func (*UnimplementedBergoglioServer) GetManyCodesStream(req *ManyCodesRequest, srv Bergoglio_GetManyCodesStreamServer) error {
	return status.Errorf(codes.Unimplemented, "method GetManyCodesStream not implemented")
}

func RegisterBergoglioServer(s *grpc.Server, srv BergoglioServer) {
	s.RegisterService(&_Bergoglio_serviceDesc, srv)
}

func _Bergoglio_GetOneCodeWithBlessing_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(Entry)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(BergoglioServer).GetOneCodeWithBlessing(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/bergogliov1.Bergoglio/GetOneCodeWithBlessing",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(BergoglioServer).GetOneCodeWithBlessing(ctx, req.(*Entry))
	}
	return interceptor(ctx, in, info, handler)
}

func _Bergoglio_GetManyCodes_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ManyCodesRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(BergoglioServer).GetManyCodes(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/bergogliov1.Bergoglio/GetManyCodes",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(BergoglioServer).GetManyCodes(ctx, req.(*ManyCodesRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Bergoglio_GetManyCodesStream_Handler(srv interface{}, stream grpc.ServerStream) error {
	m := new(ManyCodesRequest)
	if err := stream.RecvMsg(m); err != nil {
		return err
	}
	return srv.(BergoglioServer).GetManyCodesStream(m, &bergoglioGetManyCodesStreamServer{stream})
}

type Bergoglio_GetManyCodesStreamServer interface {
	Send(*ManyCodesResponse) error
	grpc.ServerStream
}

type bergoglioGetManyCodesStreamServer struct {
	grpc.ServerStream
}

func (x *bergoglioGetManyCodesStreamServer) Send(m *ManyCodesResponse) error {
	return x.ServerStream.SendMsg(m)
}

var _Bergoglio_serviceDesc = grpc.ServiceDesc{
	ServiceName: "bergogliov1.Bergoglio",
	HandlerType: (*BergoglioServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "GetOneCodeWithBlessing",
			Handler:    _Bergoglio_GetOneCodeWithBlessing_Handler,
		},
		{
			MethodName: "GetManyCodes",
			Handler:    _Bergoglio_GetManyCodes_Handler,
		},
	},
	Streams: []grpc.StreamDesc{
		{
			StreamName:    "GetManyCodesStream",
			Handler:       _Bergoglio_GetManyCodesStream_Handler,
			ServerStreams: true,
		},
	},
	Metadata: "proto/code.proto",
}
