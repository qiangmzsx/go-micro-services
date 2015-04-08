// Code generated by protoc-gen-go.
// source: proto/user/user.proto
// DO NOT EDIT!

/*
Package user is a generated protocol buffer package.

It is generated from these files:
	proto/user/user.proto

It has these top-level messages:
	Req
	User
*/
package user

import proto "github.com/golang/protobuf/proto"

import (
	context "golang.org/x/net/context"
	grpc "google.golang.org/grpc"
)

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConn

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal

type Req struct {
	Token string `protobuf:"bytes,1,opt" json:"Token,omitempty"`
}

func (m *Req) Reset()         { *m = Req{} }
func (m *Req) String() string { return proto.CompactTextString(m) }
func (*Req) ProtoMessage()    {}

type User struct {
	ID        int32  `protobuf:"varint,1,opt" json:"ID,omitempty"`
	Email     string `protobuf:"bytes,2,opt" json:"Email,omitempty"`
	FirstName string `protobuf:"bytes,3,opt" json:"FirstName,omitempty"`
	LastName  string `protobuf:"bytes,4,opt" json:"LastName,omitempty"`
}

func (m *User) Reset()         { *m = User{} }
func (m *User) String() string { return proto.CompactTextString(m) }
func (*User) ProtoMessage()    {}

func init() {
}

// Client API for UserLookup service

type UserLookupClient interface {
	GetUser(ctx context.Context, in *Req, opts ...grpc.CallOption) (*User, error)
}

type userLookupClient struct {
	cc *grpc.ClientConn
}

func NewUserLookupClient(cc *grpc.ClientConn) UserLookupClient {
	return &userLookupClient{cc}
}

func (c *userLookupClient) GetUser(ctx context.Context, in *Req, opts ...grpc.CallOption) (*User, error) {
	out := new(User)
	err := grpc.Invoke(ctx, "/.UserLookup/GetUser", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// Server API for UserLookup service

type UserLookupServer interface {
	GetUser(context.Context, *Req) (*User, error)
}

func RegisterUserLookupServer(s *grpc.Server, srv UserLookupServer) {
	s.RegisterService(&_UserLookup_serviceDesc, srv)
}

func _UserLookup_GetUser_Handler(srv interface{}, ctx context.Context, buf []byte) (interface{}, error) {
	in := new(Req)
	if err := proto.Unmarshal(buf, in); err != nil {
		return nil, err
	}
	out, err := srv.(UserLookupServer).GetUser(ctx, in)
	if err != nil {
		return nil, err
	}
	return out, nil
}

var _UserLookup_serviceDesc = grpc.ServiceDesc{
	ServiceName: ".UserLookup",
	HandlerType: (*UserLookupServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "GetUser",
			Handler:    _UserLookup_GetUser_Handler,
		},
	},
	Streams: []grpc.StreamDesc{},
}
