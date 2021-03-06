// Code generated by protoc-gen-go. DO NOT EDIT.
// source: authentication.proto

package authentication

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

type SignInMessageRequest struct {
	UserId               string   `protobuf:"bytes,1,opt,name=user_id,json=userId,proto3" json:"user_id,omitempty"`
	Password             string   `protobuf:"bytes,2,opt,name=password,proto3" json:"password,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *SignInMessageRequest) Reset()         { *m = SignInMessageRequest{} }
func (m *SignInMessageRequest) String() string { return proto.CompactTextString(m) }
func (*SignInMessageRequest) ProtoMessage()    {}
func (*SignInMessageRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_d0dbc99083440df2, []int{0}
}

func (m *SignInMessageRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_SignInMessageRequest.Unmarshal(m, b)
}
func (m *SignInMessageRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_SignInMessageRequest.Marshal(b, m, deterministic)
}
func (m *SignInMessageRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_SignInMessageRequest.Merge(m, src)
}
func (m *SignInMessageRequest) XXX_Size() int {
	return xxx_messageInfo_SignInMessageRequest.Size(m)
}
func (m *SignInMessageRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_SignInMessageRequest.DiscardUnknown(m)
}

var xxx_messageInfo_SignInMessageRequest proto.InternalMessageInfo

func (m *SignInMessageRequest) GetUserId() string {
	if m != nil {
		return m.UserId
	}
	return ""
}

func (m *SignInMessageRequest) GetPassword() string {
	if m != nil {
		return m.Password
	}
	return ""
}

type SignInMessageResponse struct {
	AccessToken          string   `protobuf:"bytes,1,opt,name=access_token,json=accessToken,proto3" json:"access_token,omitempty"`
	RefreshToken         string   `protobuf:"bytes,2,opt,name=refresh_token,json=refreshToken,proto3" json:"refresh_token,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *SignInMessageResponse) Reset()         { *m = SignInMessageResponse{} }
func (m *SignInMessageResponse) String() string { return proto.CompactTextString(m) }
func (*SignInMessageResponse) ProtoMessage()    {}
func (*SignInMessageResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_d0dbc99083440df2, []int{1}
}

func (m *SignInMessageResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_SignInMessageResponse.Unmarshal(m, b)
}
func (m *SignInMessageResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_SignInMessageResponse.Marshal(b, m, deterministic)
}
func (m *SignInMessageResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_SignInMessageResponse.Merge(m, src)
}
func (m *SignInMessageResponse) XXX_Size() int {
	return xxx_messageInfo_SignInMessageResponse.Size(m)
}
func (m *SignInMessageResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_SignInMessageResponse.DiscardUnknown(m)
}

var xxx_messageInfo_SignInMessageResponse proto.InternalMessageInfo

func (m *SignInMessageResponse) GetAccessToken() string {
	if m != nil {
		return m.AccessToken
	}
	return ""
}

func (m *SignInMessageResponse) GetRefreshToken() string {
	if m != nil {
		return m.RefreshToken
	}
	return ""
}

type User struct {
	Id                   string   `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
	Password             string   `protobuf:"bytes,2,opt,name=password,proto3" json:"password,omitempty"`
	DisplayName          string   `protobuf:"bytes,3,opt,name=display_name,json=displayName,proto3" json:"display_name,omitempty"`
	IconImagePath        string   `protobuf:"bytes,4,opt,name=icon_image_path,json=iconImagePath,proto3" json:"icon_image_path,omitempty"`
	BackgroundImagePath  string   `protobuf:"bytes,5,opt,name=background_image_path,json=backgroundImagePath,proto3" json:"background_image_path,omitempty"`
	Profile              string   `protobuf:"bytes,6,opt,name=profile,proto3" json:"profile,omitempty"`
	Email                string   `protobuf:"bytes,7,opt,name=email,proto3" json:"email,omitempty"`
	CreatedAt            int64    `protobuf:"varint,8,opt,name=created_at,json=createdAt,proto3" json:"created_at,omitempty"`
	UpdatedAt            int64    `protobuf:"varint,9,opt,name=updated_at,json=updatedAt,proto3" json:"updated_at,omitempty"`
	DeletedAt            int64    `protobuf:"varint,10,opt,name=deleted_at,json=deletedAt,proto3" json:"deleted_at,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *User) Reset()         { *m = User{} }
func (m *User) String() string { return proto.CompactTextString(m) }
func (*User) ProtoMessage()    {}
func (*User) Descriptor() ([]byte, []int) {
	return fileDescriptor_d0dbc99083440df2, []int{2}
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

func (m *User) GetId() string {
	if m != nil {
		return m.Id
	}
	return ""
}

func (m *User) GetPassword() string {
	if m != nil {
		return m.Password
	}
	return ""
}

func (m *User) GetDisplayName() string {
	if m != nil {
		return m.DisplayName
	}
	return ""
}

func (m *User) GetIconImagePath() string {
	if m != nil {
		return m.IconImagePath
	}
	return ""
}

func (m *User) GetBackgroundImagePath() string {
	if m != nil {
		return m.BackgroundImagePath
	}
	return ""
}

func (m *User) GetProfile() string {
	if m != nil {
		return m.Profile
	}
	return ""
}

func (m *User) GetEmail() string {
	if m != nil {
		return m.Email
	}
	return ""
}

func (m *User) GetCreatedAt() int64 {
	if m != nil {
		return m.CreatedAt
	}
	return 0
}

func (m *User) GetUpdatedAt() int64 {
	if m != nil {
		return m.UpdatedAt
	}
	return 0
}

func (m *User) GetDeletedAt() int64 {
	if m != nil {
		return m.DeletedAt
	}
	return 0
}

type CreateUserRequest struct {
	Password             string   `protobuf:"bytes,1,opt,name=password,proto3" json:"password,omitempty"`
	DisplayName          string   `protobuf:"bytes,2,opt,name=display_name,json=displayName,proto3" json:"display_name,omitempty"`
	IconImagePath        string   `protobuf:"bytes,3,opt,name=icon_image_path,json=iconImagePath,proto3" json:"icon_image_path,omitempty"`
	BackgroundImagePath  string   `protobuf:"bytes,4,opt,name=background_image_path,json=backgroundImagePath,proto3" json:"background_image_path,omitempty"`
	Profile              string   `protobuf:"bytes,5,opt,name=profile,proto3" json:"profile,omitempty"`
	Email                string   `protobuf:"bytes,6,opt,name=email,proto3" json:"email,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *CreateUserRequest) Reset()         { *m = CreateUserRequest{} }
func (m *CreateUserRequest) String() string { return proto.CompactTextString(m) }
func (*CreateUserRequest) ProtoMessage()    {}
func (*CreateUserRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_d0dbc99083440df2, []int{3}
}

func (m *CreateUserRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_CreateUserRequest.Unmarshal(m, b)
}
func (m *CreateUserRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_CreateUserRequest.Marshal(b, m, deterministic)
}
func (m *CreateUserRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_CreateUserRequest.Merge(m, src)
}
func (m *CreateUserRequest) XXX_Size() int {
	return xxx_messageInfo_CreateUserRequest.Size(m)
}
func (m *CreateUserRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_CreateUserRequest.DiscardUnknown(m)
}

var xxx_messageInfo_CreateUserRequest proto.InternalMessageInfo

func (m *CreateUserRequest) GetPassword() string {
	if m != nil {
		return m.Password
	}
	return ""
}

func (m *CreateUserRequest) GetDisplayName() string {
	if m != nil {
		return m.DisplayName
	}
	return ""
}

func (m *CreateUserRequest) GetIconImagePath() string {
	if m != nil {
		return m.IconImagePath
	}
	return ""
}

func (m *CreateUserRequest) GetBackgroundImagePath() string {
	if m != nil {
		return m.BackgroundImagePath
	}
	return ""
}

func (m *CreateUserRequest) GetProfile() string {
	if m != nil {
		return m.Profile
	}
	return ""
}

func (m *CreateUserRequest) GetEmail() string {
	if m != nil {
		return m.Email
	}
	return ""
}

type CreateUserResponse struct {
	User                 *User    `protobuf:"bytes,1,opt,name=user,proto3" json:"user,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *CreateUserResponse) Reset()         { *m = CreateUserResponse{} }
func (m *CreateUserResponse) String() string { return proto.CompactTextString(m) }
func (*CreateUserResponse) ProtoMessage()    {}
func (*CreateUserResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_d0dbc99083440df2, []int{4}
}

func (m *CreateUserResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_CreateUserResponse.Unmarshal(m, b)
}
func (m *CreateUserResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_CreateUserResponse.Marshal(b, m, deterministic)
}
func (m *CreateUserResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_CreateUserResponse.Merge(m, src)
}
func (m *CreateUserResponse) XXX_Size() int {
	return xxx_messageInfo_CreateUserResponse.Size(m)
}
func (m *CreateUserResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_CreateUserResponse.DiscardUnknown(m)
}

var xxx_messageInfo_CreateUserResponse proto.InternalMessageInfo

func (m *CreateUserResponse) GetUser() *User {
	if m != nil {
		return m.User
	}
	return nil
}

func init() {
	proto.RegisterType((*SignInMessageRequest)(nil), "SignInMessageRequest")
	proto.RegisterType((*SignInMessageResponse)(nil), "SignInMessageResponse")
	proto.RegisterType((*User)(nil), "User")
	proto.RegisterType((*CreateUserRequest)(nil), "CreateUserRequest")
	proto.RegisterType((*CreateUserResponse)(nil), "CreateUserResponse")
}

func init() { proto.RegisterFile("authentication.proto", fileDescriptor_d0dbc99083440df2) }

var fileDescriptor_d0dbc99083440df2 = []byte{
	// 432 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x94, 0x93, 0xcf, 0x8e, 0xd3, 0x30,
	0x10, 0xc6, 0x37, 0x69, 0x9b, 0x6e, 0x67, 0xff, 0x20, 0xbc, 0x2d, 0x98, 0x4a, 0x48, 0x4b, 0x90,
	0xd0, 0x9e, 0x82, 0x54, 0x0e, 0x88, 0x63, 0xc5, 0xa9, 0x42, 0x20, 0x54, 0xe0, 0x1c, 0x79, 0xe3,
	0xd9, 0xd6, 0xda, 0xc6, 0x0e, 0xb6, 0x23, 0xc4, 0x9d, 0xe7, 0xe2, 0x51, 0x78, 0x16, 0xe4, 0x3f,
	0xdd, 0x06, 0x76, 0x81, 0x72, 0x9c, 0xdf, 0xf7, 0x65, 0x6c, 0x7f, 0x99, 0x81, 0x31, 0x6b, 0xed,
	0x1a, 0xa5, 0x15, 0x15, 0xb3, 0x42, 0xc9, 0xa2, 0xd1, 0xca, 0xaa, 0xfc, 0x0d, 0x8c, 0x3f, 0x88,
	0x95, 0x5c, 0xc8, 0xb7, 0x68, 0x0c, 0x5b, 0xe1, 0x12, 0x3f, 0xb7, 0x68, 0x2c, 0x79, 0x08, 0xc3,
	0xd6, 0xa0, 0x2e, 0x05, 0xa7, 0xc9, 0x79, 0x72, 0x31, 0x5a, 0x66, 0xae, 0x5c, 0x70, 0x32, 0x85,
	0xc3, 0x86, 0x19, 0xf3, 0x45, 0x69, 0x4e, 0x53, 0xaf, 0xdc, 0xd4, 0x79, 0x09, 0x93, 0xdf, 0x9a,
	0x99, 0x46, 0x49, 0x83, 0xe4, 0x09, 0x1c, 0xb3, 0xaa, 0x42, 0x63, 0x4a, 0xab, 0xae, 0x51, 0xc6,
	0x96, 0x47, 0x81, 0x7d, 0x74, 0x88, 0x3c, 0x85, 0x13, 0x8d, 0x57, 0x1a, 0xcd, 0x3a, 0x7a, 0x42,
	0xf3, 0xe3, 0x08, 0xbd, 0x29, 0xff, 0x9e, 0x42, 0xff, 0x93, 0x41, 0x4d, 0x4e, 0x21, 0xbd, 0xb9,
	0x59, 0x2a, 0xfe, 0x7a, 0x2b, 0x77, 0x38, 0x17, 0xa6, 0xd9, 0xb0, 0xaf, 0xa5, 0x64, 0x35, 0xd2,
	0x5e, 0x38, 0x3c, 0xb2, 0x77, 0xac, 0x46, 0xf2, 0x0c, 0xee, 0x89, 0x4a, 0xc9, 0x52, 0xd4, 0x6c,
	0x85, 0x65, 0xc3, 0xec, 0x9a, 0xf6, 0xbd, 0xeb, 0xc4, 0xe1, 0x85, 0xa3, 0xef, 0x99, 0x5d, 0x93,
	0x19, 0x4c, 0x2e, 0x59, 0x75, 0xbd, 0xd2, 0xaa, 0x95, 0xbc, 0xeb, 0x1e, 0x78, 0xf7, 0xd9, 0x4e,
	0xdc, 0x7d, 0x43, 0x61, 0xd8, 0x68, 0x75, 0x25, 0x36, 0x48, 0x33, 0xef, 0xda, 0x96, 0x64, 0x0c,
	0x03, 0xac, 0x99, 0xd8, 0xd0, 0xa1, 0xe7, 0xa1, 0x20, 0x8f, 0x01, 0x2a, 0x8d, 0xcc, 0x22, 0x2f,
	0x99, 0xa5, 0x87, 0xe7, 0xc9, 0x45, 0x6f, 0x39, 0x8a, 0x64, 0x6e, 0x9d, 0xdc, 0x36, 0x7c, 0x2b,
	0x8f, 0x82, 0x1c, 0x49, 0x90, 0x39, 0x6e, 0x30, 0xca, 0x10, 0xe4, 0x48, 0xe6, 0x36, 0xff, 0x91,
	0xc0, 0xfd, 0xd7, 0xbe, 0x97, 0x8b, 0x71, 0xfb, 0xb3, 0xbb, 0xe9, 0x25, 0xff, 0x48, 0x2f, 0xdd,
	0x2b, 0xbd, 0xde, 0x7f, 0xa5, 0xd7, 0xdf, 0x2b, 0xbd, 0xc1, 0x1f, 0xd2, 0xcb, 0x3a, 0xe9, 0xe5,
	0xcf, 0x81, 0x74, 0xdf, 0x17, 0xe7, 0xef, 0x11, 0xf4, 0xdd, 0xf8, 0xfa, 0xc7, 0x1d, 0xcd, 0x06,
	0x85, 0x17, 0x3d, 0x9a, 0x7d, 0x4b, 0xe0, 0x74, 0xfe, 0xcb, 0x66, 0x90, 0x57, 0x90, 0x85, 0x31,
	0x26, 0x93, 0xe2, 0xae, 0xe5, 0x98, 0x3e, 0x28, 0xee, 0x1c, 0xf3, 0xfc, 0x80, 0xbc, 0x04, 0xd8,
	0x1d, 0x4f, 0x48, 0x71, 0x2b, 0xeb, 0xe9, 0x59, 0x71, 0xfb, 0x7e, 0xf9, 0xc1, 0x65, 0xe6, 0xd7,
	0xf1, 0xc5, 0xcf, 0x00, 0x00, 0x00, 0xff, 0xff, 0xad, 0x52, 0x31, 0xe8, 0xa6, 0x03, 0x00, 0x00,
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConn

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion4

// AuthenticationClient is the client API for Authentication service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type AuthenticationClient interface {
	SignIn(ctx context.Context, in *SignInMessageRequest, opts ...grpc.CallOption) (*SignInMessageResponse, error)
	CreateUser(ctx context.Context, in *CreateUserRequest, opts ...grpc.CallOption) (*CreateUserResponse, error)
}

type authenticationClient struct {
	cc *grpc.ClientConn
}

func NewAuthenticationClient(cc *grpc.ClientConn) AuthenticationClient {
	return &authenticationClient{cc}
}

func (c *authenticationClient) SignIn(ctx context.Context, in *SignInMessageRequest, opts ...grpc.CallOption) (*SignInMessageResponse, error) {
	out := new(SignInMessageResponse)
	err := c.cc.Invoke(ctx, "/Authentication/SignIn", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *authenticationClient) CreateUser(ctx context.Context, in *CreateUserRequest, opts ...grpc.CallOption) (*CreateUserResponse, error) {
	out := new(CreateUserResponse)
	err := c.cc.Invoke(ctx, "/Authentication/CreateUser", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// AuthenticationServer is the server API for Authentication service.
type AuthenticationServer interface {
	SignIn(context.Context, *SignInMessageRequest) (*SignInMessageResponse, error)
	CreateUser(context.Context, *CreateUserRequest) (*CreateUserResponse, error)
}

// UnimplementedAuthenticationServer can be embedded to have forward compatible implementations.
type UnimplementedAuthenticationServer struct {
}

func (*UnimplementedAuthenticationServer) SignIn(ctx context.Context, req *SignInMessageRequest) (*SignInMessageResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method SignIn not implemented")
}
func (*UnimplementedAuthenticationServer) CreateUser(ctx context.Context, req *CreateUserRequest) (*CreateUserResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method CreateUser not implemented")
}

func RegisterAuthenticationServer(s *grpc.Server, srv AuthenticationServer) {
	s.RegisterService(&_Authentication_serviceDesc, srv)
}

func _Authentication_SignIn_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(SignInMessageRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(AuthenticationServer).SignIn(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/Authentication/SignIn",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(AuthenticationServer).SignIn(ctx, req.(*SignInMessageRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Authentication_CreateUser_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(CreateUserRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(AuthenticationServer).CreateUser(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/Authentication/CreateUser",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(AuthenticationServer).CreateUser(ctx, req.(*CreateUserRequest))
	}
	return interceptor(ctx, in, info, handler)
}

var _Authentication_serviceDesc = grpc.ServiceDesc{
	ServiceName: "Authentication",
	HandlerType: (*AuthenticationServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "SignIn",
			Handler:    _Authentication_SignIn_Handler,
		},
		{
			MethodName: "CreateUser",
			Handler:    _Authentication_CreateUser_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "authentication.proto",
}
