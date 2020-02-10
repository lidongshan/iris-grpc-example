// Code generated by protoc-gen-go. DO NOT EDIT.
// source: servers.proto

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

type Id struct {
	Id                   int32    `protobuf:"varint,1,opt,name=id,proto3" json:"id,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *Id) Reset()         { *m = Id{} }
func (m *Id) String() string { return proto.CompactTextString(m) }
func (*Id) ProtoMessage()    {}
func (*Id) Descriptor() ([]byte, []int) {
	return fileDescriptor_a6eaaceba9c5e837, []int{0}
}

func (m *Id) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Id.Unmarshal(m, b)
}
func (m *Id) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Id.Marshal(b, m, deterministic)
}
func (m *Id) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Id.Merge(m, src)
}
func (m *Id) XXX_Size() int {
	return xxx_messageInfo_Id.Size(m)
}
func (m *Id) XXX_DiscardUnknown() {
	xxx_messageInfo_Id.DiscardUnknown(m)
}

var xxx_messageInfo_Id proto.InternalMessageInfo

func (m *Id) GetId() int32 {
	if m != nil {
		return m.Id
	}
	return 0
}

type Name struct {
	Name                 string   `protobuf:"bytes,1,opt,name=name,proto3" json:"name,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *Name) Reset()         { *m = Name{} }
func (m *Name) String() string { return proto.CompactTextString(m) }
func (*Name) ProtoMessage()    {}
func (*Name) Descriptor() ([]byte, []int) {
	return fileDescriptor_a6eaaceba9c5e837, []int{1}
}

func (m *Name) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Name.Unmarshal(m, b)
}
func (m *Name) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Name.Marshal(b, m, deterministic)
}
func (m *Name) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Name.Merge(m, src)
}
func (m *Name) XXX_Size() int {
	return xxx_messageInfo_Name.Size(m)
}
func (m *Name) XXX_DiscardUnknown() {
	xxx_messageInfo_Name.DiscardUnknown(m)
}

var xxx_messageInfo_Name proto.InternalMessageInfo

func (m *Name) GetName() string {
	if m != nil {
		return m.Name
	}
	return ""
}

type Age struct {
	Age                  int32    `protobuf:"varint,1,opt,name=age,proto3" json:"age,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *Age) Reset()         { *m = Age{} }
func (m *Age) String() string { return proto.CompactTextString(m) }
func (*Age) ProtoMessage()    {}
func (*Age) Descriptor() ([]byte, []int) {
	return fileDescriptor_a6eaaceba9c5e837, []int{2}
}

func (m *Age) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Age.Unmarshal(m, b)
}
func (m *Age) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Age.Marshal(b, m, deterministic)
}
func (m *Age) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Age.Merge(m, src)
}
func (m *Age) XXX_Size() int {
	return xxx_messageInfo_Age.Size(m)
}
func (m *Age) XXX_DiscardUnknown() {
	xxx_messageInfo_Age.DiscardUnknown(m)
}

var xxx_messageInfo_Age proto.InternalMessageInfo

func (m *Age) GetAge() int32 {
	if m != nil {
		return m.Age
	}
	return 0
}

// 用户变量
type User struct {
	Id                   int32    `protobuf:"varint,1,opt,name=id,proto3" json:"id,omitempty"`
	Name                 string   `protobuf:"bytes,2,opt,name=name,proto3" json:"name,omitempty"`
	Age                  int32    `protobuf:"varint,3,opt,name=age,proto3" json:"age,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *User) Reset()         { *m = User{} }
func (m *User) String() string { return proto.CompactTextString(m) }
func (*User) ProtoMessage()    {}
func (*User) Descriptor() ([]byte, []int) {
	return fileDescriptor_a6eaaceba9c5e837, []int{3}
}

func (m *User) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_User.Unmarshal(m, b)
}
func (m *User) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_User.Marshal(b, m, deterministic)
}
func (m *User) XXX_Merge(src proto.Message) {
	xxx_messageInfo_User.Merge(m, src)
}
func (m *User) XXX_Size() int {
	return xxx_messageInfo_User.Size(m)
}
func (m *User) XXX_DiscardUnknown() {
	xxx_messageInfo_User.DiscardUnknown(m)
}

var xxx_messageInfo_User proto.InternalMessageInfo

func (m *User) GetId() int32 {
	if m != nil {
		return m.Id
	}
	return 0
}

func (m *User) GetName() string {
	if m != nil {
		return m.Name
	}
	return ""
}

func (m *User) GetAge() int32 {
	if m != nil {
		return m.Age
	}
	return 0
}

// 用户参数
type UserParams struct {
	Name                 *Name    `protobuf:"bytes,1,opt,name=name,proto3" json:"name,omitempty"`
	Age                  *Age     `protobuf:"bytes,2,opt,name=age,proto3" json:"age,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *UserParams) Reset()         { *m = UserParams{} }
func (m *UserParams) String() string { return proto.CompactTextString(m) }
func (*UserParams) ProtoMessage()    {}
func (*UserParams) Descriptor() ([]byte, []int) {
	return fileDescriptor_a6eaaceba9c5e837, []int{4}
}

func (m *UserParams) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_UserParams.Unmarshal(m, b)
}
func (m *UserParams) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_UserParams.Marshal(b, m, deterministic)
}
func (m *UserParams) XXX_Merge(src proto.Message) {
	xxx_messageInfo_UserParams.Merge(m, src)
}
func (m *UserParams) XXX_Size() int {
	return xxx_messageInfo_UserParams.Size(m)
}
func (m *UserParams) XXX_DiscardUnknown() {
	xxx_messageInfo_UserParams.DiscardUnknown(m)
}

var xxx_messageInfo_UserParams proto.InternalMessageInfo

func (m *UserParams) GetName() *Name {
	if m != nil {
		return m.Name
	}
	return nil
}

func (m *UserParams) GetAge() *Age {
	if m != nil {
		return m.Age
	}
	return nil
}

func init() {
	proto.RegisterType((*Id)(nil), "proto.Id")
	proto.RegisterType((*Name)(nil), "proto.Name")
	proto.RegisterType((*Age)(nil), "proto.Age")
	proto.RegisterType((*User)(nil), "proto.User")
	proto.RegisterType((*UserParams)(nil), "proto.UserParams")
}

func init() { proto.RegisterFile("servers.proto", fileDescriptor_a6eaaceba9c5e837) }

var fileDescriptor_a6eaaceba9c5e837 = []byte{
	// 223 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x5c, 0x8e, 0xc1, 0x4a, 0xc4, 0x30,
	0x10, 0x86, 0x77, 0xd3, 0xae, 0x6c, 0xa7, 0x54, 0x74, 0x10, 0x94, 0x22, 0x28, 0xc1, 0x83, 0x78,
	0xe8, 0xa1, 0x5e, 0xbd, 0xf4, 0x58, 0x04, 0x91, 0x16, 0x1f, 0x60, 0x6c, 0xc6, 0xda, 0x43, 0x5b,
	0x49, 0xa4, 0xcf, 0x2f, 0x49, 0xa3, 0x29, 0x9e, 0x32, 0xc9, 0x9f, 0xef, 0x9b, 0x1f, 0x32, 0xc3,
	0x7a, 0x61, 0x6d, 0x8a, 0x2f, 0x3d, 0x7f, 0xcf, 0x78, 0x70, 0x87, 0xbc, 0x00, 0x51, 0x2b, 0x3c,
	0x05, 0x31, 0xa8, 0xab, 0xfd, 0xed, 0xfe, 0xfe, 0xd0, 0x88, 0x41, 0xc9, 0x1c, 0xe2, 0x17, 0x1a,
	0x19, 0x11, 0xe2, 0x89, 0x46, 0x76, 0x49, 0xd2, 0xb8, 0x59, 0x5e, 0x42, 0x54, 0xf5, 0x8c, 0x67,
	0x10, 0x51, 0xcf, 0x9e, 0xb1, 0xa3, 0x7c, 0x82, 0xf8, 0xcd, 0xb0, 0xfe, 0x2f, 0xfb, 0x93, 0x88,
	0x20, 0xf9, 0xa5, 0xa3, 0x40, 0x3f, 0x03, 0x58, 0xfa, 0x95, 0x34, 0x8d, 0x06, 0x6f, 0x36, 0x8b,
	0xd3, 0x32, 0x5d, 0x3b, 0x17, 0xb6, 0x93, 0x17, 0x5c, 0xaf, 0x02, 0xe1, 0x72, 0xf0, 0x79, 0xd5,
	0xb3, 0x93, 0x95, 0x04, 0x59, 0xcb, 0x7a, 0x19, 0x3a, 0x6e, 0x99, 0x74, 0xf7, 0x89, 0x0f, 0x70,
	0x6c, 0x69, 0x61, 0xd7, 0xef, 0xdc, 0xff, 0x0e, 0xeb, 0xf2, 0xc4, 0x3f, 0xd5, 0x4a, 0xee, 0xf0,
	0x0e, 0x8e, 0x36, 0xaa, 0xa7, 0x8f, 0x19, 0x43, 0x90, 0xa7, 0x1b, 0x4c, 0xee, 0xde, 0x4f, 0xdc,
	0xed, 0xf1, 0x27, 0x00, 0x00, 0xff, 0xff, 0x7f, 0x92, 0x59, 0x17, 0x57, 0x01, 0x00, 0x00,
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConnInterface

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion6

// ServiceSearchClient is the client API for ServiceSearch service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type ServiceSearchClient interface {
	SaveUser(ctx context.Context, in *UserParams, opts ...grpc.CallOption) (*Id, error)
	UserInfo(ctx context.Context, in *Id, opts ...grpc.CallOption) (*User, error)
}

type serviceSearchClient struct {
	cc grpc.ClientConnInterface
}

func NewServiceSearchClient(cc grpc.ClientConnInterface) ServiceSearchClient {
	return &serviceSearchClient{cc}
}

func (c *serviceSearchClient) SaveUser(ctx context.Context, in *UserParams, opts ...grpc.CallOption) (*Id, error) {
	out := new(Id)
	err := c.cc.Invoke(ctx, "/proto.ServiceSearch/SaveUser", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *serviceSearchClient) UserInfo(ctx context.Context, in *Id, opts ...grpc.CallOption) (*User, error) {
	out := new(User)
	err := c.cc.Invoke(ctx, "/proto.ServiceSearch/UserInfo", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// ServiceSearchServer is the server API for ServiceSearch service.
type ServiceSearchServer interface {
	SaveUser(context.Context, *UserParams) (*Id, error)
	UserInfo(context.Context, *Id) (*User, error)
}

// UnimplementedServiceSearchServer can be embedded to have forward compatible implementations.
type UnimplementedServiceSearchServer struct {
}

func (*UnimplementedServiceSearchServer) SaveUser(ctx context.Context, req *UserParams) (*Id, error) {
	return nil, status.Errorf(codes.Unimplemented, "method SaveUser not implemented")
}
func (*UnimplementedServiceSearchServer) UserInfo(ctx context.Context, req *Id) (*User, error) {
	return nil, status.Errorf(codes.Unimplemented, "method UserInfo not implemented")
}

func RegisterServiceSearchServer(s *grpc.Server, srv ServiceSearchServer) {
	s.RegisterService(&_ServiceSearch_serviceDesc, srv)
}

func _ServiceSearch_SaveUser_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(UserParams)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ServiceSearchServer).SaveUser(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/proto.ServiceSearch/SaveUser",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ServiceSearchServer).SaveUser(ctx, req.(*UserParams))
	}
	return interceptor(ctx, in, info, handler)
}

func _ServiceSearch_UserInfo_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(Id)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ServiceSearchServer).UserInfo(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/proto.ServiceSearch/UserInfo",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ServiceSearchServer).UserInfo(ctx, req.(*Id))
	}
	return interceptor(ctx, in, info, handler)
}

var _ServiceSearch_serviceDesc = grpc.ServiceDesc{
	ServiceName: "proto.ServiceSearch",
	HandlerType: (*ServiceSearchServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "SaveUser",
			Handler:    _ServiceSearch_SaveUser_Handler,
		},
		{
			MethodName: "UserInfo",
			Handler:    _ServiceSearch_UserInfo_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "servers.proto",
}
