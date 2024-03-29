// Code generated by protoc-gen-go. DO NOT EDIT.
// source: gbooth.proto

package proto

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

type Request struct {
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *Request) Reset()         { *m = Request{} }
func (m *Request) String() string { return proto.CompactTextString(m) }
func (*Request) ProtoMessage()    {}
func (*Request) Descriptor() ([]byte, []int) {
	return fileDescriptor_ac730dea810b8175, []int{0}
}

func (m *Request) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Request.Unmarshal(m, b)
}
func (m *Request) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Request.Marshal(b, m, deterministic)
}
func (m *Request) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Request.Merge(m, src)
}
func (m *Request) XXX_Size() int {
	return xxx_messageInfo_Request.Size(m)
}
func (m *Request) XXX_DiscardUnknown() {
	xxx_messageInfo_Request.DiscardUnknown(m)
}

var xxx_messageInfo_Request proto.InternalMessageInfo

type ImageResponse struct {
	Name                 string   `protobuf:"bytes,1,opt,name=name,proto3" json:"name,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *ImageResponse) Reset()         { *m = ImageResponse{} }
func (m *ImageResponse) String() string { return proto.CompactTextString(m) }
func (*ImageResponse) ProtoMessage()    {}
func (*ImageResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_ac730dea810b8175, []int{1}
}

func (m *ImageResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_ImageResponse.Unmarshal(m, b)
}
func (m *ImageResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_ImageResponse.Marshal(b, m, deterministic)
}
func (m *ImageResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ImageResponse.Merge(m, src)
}
func (m *ImageResponse) XXX_Size() int {
	return xxx_messageInfo_ImageResponse.Size(m)
}
func (m *ImageResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_ImageResponse.DiscardUnknown(m)
}

var xxx_messageInfo_ImageResponse proto.InternalMessageInfo

func (m *ImageResponse) GetName() string {
	if m != nil {
		return m.Name
	}
	return ""
}

type DateResponse struct {
	Date                 string   `protobuf:"bytes,1,opt,name=date,proto3" json:"date,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *DateResponse) Reset()         { *m = DateResponse{} }
func (m *DateResponse) String() string { return proto.CompactTextString(m) }
func (*DateResponse) ProtoMessage()    {}
func (*DateResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_ac730dea810b8175, []int{2}
}

func (m *DateResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_DateResponse.Unmarshal(m, b)
}
func (m *DateResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_DateResponse.Marshal(b, m, deterministic)
}
func (m *DateResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_DateResponse.Merge(m, src)
}
func (m *DateResponse) XXX_Size() int {
	return xxx_messageInfo_DateResponse.Size(m)
}
func (m *DateResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_DateResponse.DiscardUnknown(m)
}

var xxx_messageInfo_DateResponse proto.InternalMessageInfo

func (m *DateResponse) GetDate() string {
	if m != nil {
		return m.Date
	}
	return ""
}

type Image struct {
	Image                []byte   `protobuf:"bytes,1,opt,name=image,proto3" json:"image,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *Image) Reset()         { *m = Image{} }
func (m *Image) String() string { return proto.CompactTextString(m) }
func (*Image) ProtoMessage()    {}
func (*Image) Descriptor() ([]byte, []int) {
	return fileDescriptor_ac730dea810b8175, []int{3}
}

func (m *Image) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Image.Unmarshal(m, b)
}
func (m *Image) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Image.Marshal(b, m, deterministic)
}
func (m *Image) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Image.Merge(m, src)
}
func (m *Image) XXX_Size() int {
	return xxx_messageInfo_Image.Size(m)
}
func (m *Image) XXX_DiscardUnknown() {
	xxx_messageInfo_Image.DiscardUnknown(m)
}

var xxx_messageInfo_Image proto.InternalMessageInfo

func (m *Image) GetImage() []byte {
	if m != nil {
		return m.Image
	}
	return nil
}

type TextRequest struct {
	Text                 string   `protobuf:"bytes,1,opt,name=text,proto3" json:"text,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *TextRequest) Reset()         { *m = TextRequest{} }
func (m *TextRequest) String() string { return proto.CompactTextString(m) }
func (*TextRequest) ProtoMessage()    {}
func (*TextRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_ac730dea810b8175, []int{4}
}

func (m *TextRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_TextRequest.Unmarshal(m, b)
}
func (m *TextRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_TextRequest.Marshal(b, m, deterministic)
}
func (m *TextRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_TextRequest.Merge(m, src)
}
func (m *TextRequest) XXX_Size() int {
	return xxx_messageInfo_TextRequest.Size(m)
}
func (m *TextRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_TextRequest.DiscardUnknown(m)
}

var xxx_messageInfo_TextRequest proto.InternalMessageInfo

func (m *TextRequest) GetText() string {
	if m != nil {
		return m.Text
	}
	return ""
}

type UIResponse struct {
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *UIResponse) Reset()         { *m = UIResponse{} }
func (m *UIResponse) String() string { return proto.CompactTextString(m) }
func (*UIResponse) ProtoMessage()    {}
func (*UIResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_ac730dea810b8175, []int{5}
}

func (m *UIResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_UIResponse.Unmarshal(m, b)
}
func (m *UIResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_UIResponse.Marshal(b, m, deterministic)
}
func (m *UIResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_UIResponse.Merge(m, src)
}
func (m *UIResponse) XXX_Size() int {
	return xxx_messageInfo_UIResponse.Size(m)
}
func (m *UIResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_UIResponse.DiscardUnknown(m)
}

var xxx_messageInfo_UIResponse proto.InternalMessageInfo

type TriggerRequest struct {
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *TriggerRequest) Reset()         { *m = TriggerRequest{} }
func (m *TriggerRequest) String() string { return proto.CompactTextString(m) }
func (*TriggerRequest) ProtoMessage()    {}
func (*TriggerRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_ac730dea810b8175, []int{6}
}

func (m *TriggerRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_TriggerRequest.Unmarshal(m, b)
}
func (m *TriggerRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_TriggerRequest.Marshal(b, m, deterministic)
}
func (m *TriggerRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_TriggerRequest.Merge(m, src)
}
func (m *TriggerRequest) XXX_Size() int {
	return xxx_messageInfo_TriggerRequest.Size(m)
}
func (m *TriggerRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_TriggerRequest.DiscardUnknown(m)
}

var xxx_messageInfo_TriggerRequest proto.InternalMessageInfo

type CoreResponse struct {
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *CoreResponse) Reset()         { *m = CoreResponse{} }
func (m *CoreResponse) String() string { return proto.CompactTextString(m) }
func (*CoreResponse) ProtoMessage()    {}
func (*CoreResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_ac730dea810b8175, []int{7}
}

func (m *CoreResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_CoreResponse.Unmarshal(m, b)
}
func (m *CoreResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_CoreResponse.Marshal(b, m, deterministic)
}
func (m *CoreResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_CoreResponse.Merge(m, src)
}
func (m *CoreResponse) XXX_Size() int {
	return xxx_messageInfo_CoreResponse.Size(m)
}
func (m *CoreResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_CoreResponse.DiscardUnknown(m)
}

var xxx_messageInfo_CoreResponse proto.InternalMessageInfo

func init() {
	proto.RegisterType((*Request)(nil), "proto.Request")
	proto.RegisterType((*ImageResponse)(nil), "proto.ImageResponse")
	proto.RegisterType((*DateResponse)(nil), "proto.DateResponse")
	proto.RegisterType((*Image)(nil), "proto.Image")
	proto.RegisterType((*TextRequest)(nil), "proto.TextRequest")
	proto.RegisterType((*UIResponse)(nil), "proto.UIResponse")
	proto.RegisterType((*TriggerRequest)(nil), "proto.TriggerRequest")
	proto.RegisterType((*CoreResponse)(nil), "proto.CoreResponse")
}

func init() { proto.RegisterFile("gbooth.proto", fileDescriptor_ac730dea810b8175) }

var fileDescriptor_ac730dea810b8175 = []byte{
	// 296 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x8c, 0x50, 0x51, 0x4b, 0xb4, 0x40,
	0x14, 0xd5, 0x0f, 0xfd, 0x64, 0x6f, 0x93, 0xd4, 0x6d, 0x83, 0x58, 0x08, 0x6a, 0x7a, 0xe9, 0x69,
	0x2b, 0x7b, 0xe8, 0x31, 0x68, 0x37, 0x42, 0xe8, 0xc9, 0xdc, 0x1f, 0x30, 0xdb, 0x5e, 0x4c, 0x16,
	0x1d, 0x1b, 0x47, 0xd8, 0xbf, 0xd2, 0xbf, 0x8d, 0x71, 0x47, 0x51, 0x22, 0xe8, 0xc9, 0x73, 0x9c,
	0x7b, 0xce, 0xbd, 0xe7, 0x00, 0xcb, 0xd6, 0x52, 0xea, 0x8f, 0x79, 0xa5, 0xa4, 0x96, 0xe8, 0xb7,
	0x1f, 0x3e, 0x81, 0x20, 0xa1, 0xcf, 0x86, 0x6a, 0xcd, 0xaf, 0xe0, 0x30, 0x2e, 0x44, 0x46, 0x09,
	0xd5, 0x95, 0x2c, 0x6b, 0x42, 0x04, 0xaf, 0x14, 0x05, 0x9d, 0xb9, 0x17, 0xee, 0xf5, 0x24, 0x69,
	0x31, 0xe7, 0xc0, 0x96, 0x42, 0x8f, 0x66, 0x36, 0x42, 0xf7, 0x33, 0x06, 0xf3, 0x73, 0xf0, 0x5b,
	0x23, 0x9c, 0x82, 0x9f, 0x1b, 0xd0, 0xbe, 0xb2, 0x64, 0x4f, 0xf8, 0x25, 0x1c, 0xa4, 0xb4, 0xd3,
	0x76, 0xad, 0x71, 0xd0, 0xb4, 0xd3, 0x9d, 0x83, 0xc1, 0x9c, 0x01, 0xac, 0xe2, 0x6e, 0x07, 0x3f,
	0x82, 0x30, 0x55, 0x79, 0x96, 0x91, 0xea, 0x4e, 0x0d, 0x81, 0x2d, 0xa4, 0xea, 0xaf, 0x88, 0xb6,
	0xe0, 0x2d, 0xdf, 0x5e, 0x13, 0xbc, 0x83, 0x60, 0x21, 0x2a, 0xdd, 0x28, 0xc2, 0x70, 0x9f, 0x73,
	0x6e, 0x25, 0xb3, 0xa9, 0xe5, 0xa3, 0x88, 0xdc, 0xc1, 0x5b, 0x08, 0x5e, 0x48, 0x9b, 0x4c, 0x3f,
	0x24, 0x27, 0x96, 0x0f, 0x03, 0x73, 0x27, 0xfa, 0x72, 0xe1, 0xdf, 0x2a, 0xc6, 0x1b, 0x80, 0x27,
	0xf1, 0xbe, 0xcd, 0x94, 0x6c, 0xca, 0x0d, 0xb2, 0xa1, 0xfd, 0xec, 0xd8, 0xb2, 0x41, 0x08, 0xb3,
	0xc9, 0x4f, 0xf3, 0x82, 0x14, 0xa2, 0x7d, 0x1d, 0xb4, 0xf0, 0xab, 0xe2, 0x59, 0x29, 0xf9, 0x77,
	0x45, 0xf4, 0x08, 0x9e, 0x29, 0x06, 0x1f, 0x20, 0xb0, 0x95, 0xe1, 0x69, 0xa7, 0x1d, 0x55, 0xd8,
	0x87, 0x1b, 0xf6, 0xc8, 0x9d, 0xf5, 0xff, 0xf6, 0xef, 0xfd, 0x77, 0x00, 0x00, 0x00, 0xff, 0xff,
	0xaf, 0x79, 0x44, 0xfd, 0x2d, 0x02, 0x00, 0x00,
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConn

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion4

// DSLRClient is the client API for DSLR service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type DSLRClient interface {
	Capture(ctx context.Context, in *Request, opts ...grpc.CallOption) (*ImageResponse, error)
	GetDate(ctx context.Context, in *Request, opts ...grpc.CallOption) (*DateResponse, error)
}

type dSLRClient struct {
	cc *grpc.ClientConn
}

func NewDSLRClient(cc *grpc.ClientConn) DSLRClient {
	return &dSLRClient{cc}
}

func (c *dSLRClient) Capture(ctx context.Context, in *Request, opts ...grpc.CallOption) (*ImageResponse, error) {
	out := new(ImageResponse)
	err := c.cc.Invoke(ctx, "/proto.DSLR/Capture", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *dSLRClient) GetDate(ctx context.Context, in *Request, opts ...grpc.CallOption) (*DateResponse, error) {
	out := new(DateResponse)
	err := c.cc.Invoke(ctx, "/proto.DSLR/GetDate", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// DSLRServer is the server API for DSLR service.
type DSLRServer interface {
	Capture(context.Context, *Request) (*ImageResponse, error)
	GetDate(context.Context, *Request) (*DateResponse, error)
}

// UnimplementedDSLRServer can be embedded to have forward compatible implementations.
type UnimplementedDSLRServer struct {
}

func (*UnimplementedDSLRServer) Capture(ctx context.Context, req *Request) (*ImageResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Capture not implemented")
}
func (*UnimplementedDSLRServer) GetDate(ctx context.Context, req *Request) (*DateResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetDate not implemented")
}

func RegisterDSLRServer(s *grpc.Server, srv DSLRServer) {
	s.RegisterService(&_DSLR_serviceDesc, srv)
}

func _DSLR_Capture_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(Request)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(DSLRServer).Capture(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/proto.DSLR/Capture",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(DSLRServer).Capture(ctx, req.(*Request))
	}
	return interceptor(ctx, in, info, handler)
}

func _DSLR_GetDate_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(Request)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(DSLRServer).GetDate(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/proto.DSLR/GetDate",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(DSLRServer).GetDate(ctx, req.(*Request))
	}
	return interceptor(ctx, in, info, handler)
}

var _DSLR_serviceDesc = grpc.ServiceDesc{
	ServiceName: "proto.DSLR",
	HandlerType: (*DSLRServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "Capture",
			Handler:    _DSLR_Capture_Handler,
		},
		{
			MethodName: "GetDate",
			Handler:    _DSLR_GetDate_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "gbooth.proto",
}

// UIClient is the client API for UI service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type UIClient interface {
	Background(ctx context.Context, in *Image, opts ...grpc.CallOption) (*UIResponse, error)
	Timer(ctx context.Context, in *TextRequest, opts ...grpc.CallOption) (*UIResponse, error)
	Error(ctx context.Context, in *TextRequest, opts ...grpc.CallOption) (*UIResponse, error)
}

type uIClient struct {
	cc *grpc.ClientConn
}

func NewUIClient(cc *grpc.ClientConn) UIClient {
	return &uIClient{cc}
}

func (c *uIClient) Background(ctx context.Context, in *Image, opts ...grpc.CallOption) (*UIResponse, error) {
	out := new(UIResponse)
	err := c.cc.Invoke(ctx, "/proto.UI/Background", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *uIClient) Timer(ctx context.Context, in *TextRequest, opts ...grpc.CallOption) (*UIResponse, error) {
	out := new(UIResponse)
	err := c.cc.Invoke(ctx, "/proto.UI/Timer", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *uIClient) Error(ctx context.Context, in *TextRequest, opts ...grpc.CallOption) (*UIResponse, error) {
	out := new(UIResponse)
	err := c.cc.Invoke(ctx, "/proto.UI/Error", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// UIServer is the server API for UI service.
type UIServer interface {
	Background(context.Context, *Image) (*UIResponse, error)
	Timer(context.Context, *TextRequest) (*UIResponse, error)
	Error(context.Context, *TextRequest) (*UIResponse, error)
}

// UnimplementedUIServer can be embedded to have forward compatible implementations.
type UnimplementedUIServer struct {
}

func (*UnimplementedUIServer) Background(ctx context.Context, req *Image) (*UIResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Background not implemented")
}
func (*UnimplementedUIServer) Timer(ctx context.Context, req *TextRequest) (*UIResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Timer not implemented")
}
func (*UnimplementedUIServer) Error(ctx context.Context, req *TextRequest) (*UIResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Error not implemented")
}

func RegisterUIServer(s *grpc.Server, srv UIServer) {
	s.RegisterService(&_UI_serviceDesc, srv)
}

func _UI_Background_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(Image)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(UIServer).Background(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/proto.UI/Background",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(UIServer).Background(ctx, req.(*Image))
	}
	return interceptor(ctx, in, info, handler)
}

func _UI_Timer_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(TextRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(UIServer).Timer(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/proto.UI/Timer",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(UIServer).Timer(ctx, req.(*TextRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _UI_Error_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(TextRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(UIServer).Error(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/proto.UI/Error",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(UIServer).Error(ctx, req.(*TextRequest))
	}
	return interceptor(ctx, in, info, handler)
}

var _UI_serviceDesc = grpc.ServiceDesc{
	ServiceName: "proto.UI",
	HandlerType: (*UIServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "Background",
			Handler:    _UI_Background_Handler,
		},
		{
			MethodName: "Timer",
			Handler:    _UI_Timer_Handler,
		},
		{
			MethodName: "Error",
			Handler:    _UI_Error_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "gbooth.proto",
}

// CoreClient is the client API for Core service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type CoreClient interface {
	Trigger(ctx context.Context, in *TriggerRequest, opts ...grpc.CallOption) (*CoreResponse, error)
}

type coreClient struct {
	cc *grpc.ClientConn
}

func NewCoreClient(cc *grpc.ClientConn) CoreClient {
	return &coreClient{cc}
}

func (c *coreClient) Trigger(ctx context.Context, in *TriggerRequest, opts ...grpc.CallOption) (*CoreResponse, error) {
	out := new(CoreResponse)
	err := c.cc.Invoke(ctx, "/proto.Core/Trigger", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// CoreServer is the server API for Core service.
type CoreServer interface {
	Trigger(context.Context, *TriggerRequest) (*CoreResponse, error)
}

// UnimplementedCoreServer can be embedded to have forward compatible implementations.
type UnimplementedCoreServer struct {
}

func (*UnimplementedCoreServer) Trigger(ctx context.Context, req *TriggerRequest) (*CoreResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Trigger not implemented")
}

func RegisterCoreServer(s *grpc.Server, srv CoreServer) {
	s.RegisterService(&_Core_serviceDesc, srv)
}

func _Core_Trigger_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(TriggerRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(CoreServer).Trigger(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/proto.Core/Trigger",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(CoreServer).Trigger(ctx, req.(*TriggerRequest))
	}
	return interceptor(ctx, in, info, handler)
}

var _Core_serviceDesc = grpc.ServiceDesc{
	ServiceName: "proto.Core",
	HandlerType: (*CoreServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "Trigger",
			Handler:    _Core_Trigger_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "gbooth.proto",
}
