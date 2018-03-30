// Code generated by protoc-gen-go. DO NOT EDIT.
// source: pubsub.proto

/*
Package pubsub is a generated protocol buffer package.

It is generated from these files:
	pubsub.proto

It has these top-level messages:
	Reply
	UserReply
	IdRequest
*/
package pubsub

import proto "github.com/golang/protobuf/proto"
import fmt "fmt"
import math "math"
import google_rpc "google.golang.org/genproto/googleapis/rpc/status"
import dictybase_user "github.com/dictyBase/go-genproto/dictybaseapis/user"

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// This is a compile-time assertion to ensure that this generated file
// is compatible with the proto package it is being compiled against.
// A compilation error at this line likely means your copy of the
// proto package needs to be updated.
const _ = proto.ProtoPackageIsVersion2 // please upgrade the proto package

//
// Definition for communicating the existence of a resource
// and if there's any error during the communication
type Reply struct {
	// Flag to indicate the presence of resource
	Exist bool `protobuf:"varint,1,opt,name=exist" json:"exist,omitempty"`
	// Status error model, look here
	// for detail https://github.com/googleapis/googleapis/blob/master/google/rpc/status.proto
	Status *google_rpc.Status `protobuf:"bytes,2,opt,name=status" json:"status,omitempty"`
}

func (m *Reply) Reset()                    { *m = Reply{} }
func (m *Reply) String() string            { return proto.CompactTextString(m) }
func (*Reply) ProtoMessage()               {}
func (*Reply) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{0} }

func (m *Reply) GetExist() bool {
	if m != nil {
		return m.Exist
	}
	return false
}

func (m *Reply) GetStatus() *google_rpc.Status {
	if m != nil {
		return m.Status
	}
	return nil
}

//
// Definition for communicating various user definition
type UserReply struct {
	User   *dictybase_user.User           `protobuf:"bytes,1,opt,name=user" json:"user,omitempty"`
	Users  *dictybase_user.UserCollection `protobuf:"bytes,2,opt,name=users" json:"users,omitempty"`
	Status *google_rpc.Status             `protobuf:"bytes,3,opt,name=status" json:"status,omitempty"`
}

func (m *UserReply) Reset()                    { *m = UserReply{} }
func (m *UserReply) String() string            { return proto.CompactTextString(m) }
func (*UserReply) ProtoMessage()               {}
func (*UserReply) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{1} }

func (m *UserReply) GetUser() *dictybase_user.User {
	if m != nil {
		return m.User
	}
	return nil
}

func (m *UserReply) GetUsers() *dictybase_user.UserCollection {
	if m != nil {
		return m.Users
	}
	return nil
}

func (m *UserReply) GetStatus() *google_rpc.Status {
	if m != nil {
		return m.Status
	}
	return nil
}

// Defintion for communicating an identifer
type IdRequest struct {
	Id int64 `protobuf:"varint,1,opt,name=id" json:"id,omitempty"`
}

func (m *IdRequest) Reset()                    { *m = IdRequest{} }
func (m *IdRequest) String() string            { return proto.CompactTextString(m) }
func (*IdRequest) ProtoMessage()               {}
func (*IdRequest) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{2} }

func (m *IdRequest) GetId() int64 {
	if m != nil {
		return m.Id
	}
	return 0
}

func init() {
	proto.RegisterType((*Reply)(nil), "dictybase.pubsub.Reply")
	proto.RegisterType((*UserReply)(nil), "dictybase.pubsub.UserReply")
	proto.RegisterType((*IdRequest)(nil), "dictybase.pubsub.IdRequest")
}

func init() { proto.RegisterFile("pubsub.proto", fileDescriptor0) }

var fileDescriptor0 = []byte{
	// 298 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x84, 0x90, 0xcf, 0x4b, 0xc3, 0x30,
	0x1c, 0xc5, 0x49, 0xe7, 0xc6, 0x96, 0x89, 0x48, 0x28, 0x38, 0x27, 0xc8, 0xd8, 0xa9, 0x08, 0x26,
	0x30, 0xbd, 0xe9, 0xc5, 0xce, 0x4b, 0x6f, 0x25, 0xd3, 0x83, 0xde, 0x9a, 0x36, 0xc4, 0x40, 0x5d,
	0xb2, 0xfc, 0x00, 0xf7, 0x97, 0x78, 0xf7, 0xaf, 0xf4, 0x28, 0x4d, 0x74, 0x8a, 0x08, 0x5e, 0xf2,
	0x25, 0xf9, 0x7e, 0xf2, 0xde, 0xe3, 0xc1, 0x7d, 0xed, 0x99, 0xf5, 0x0c, 0x6b, 0xa3, 0x9c, 0x42,
	0x87, 0x8d, 0xac, 0xdd, 0x96, 0x55, 0x96, 0xe3, 0xf8, 0x3e, 0x3d, 0x12, 0x4a, 0x89, 0x96, 0x13,
	0xa3, 0x6b, 0x62, 0x5d, 0xe5, 0xbc, 0x8d, 0xe8, 0xf4, 0x78, 0x87, 0x12, 0x6f, 0xb9, 0x09, 0x47,
	0x5c, 0xcd, 0x0b, 0xd8, 0xa7, 0x5c, 0xb7, 0x5b, 0x94, 0xc2, 0x3e, 0x7f, 0x91, 0xd6, 0x4d, 0xc0,
	0x0c, 0x64, 0x43, 0x1a, 0x2f, 0xe8, 0x0c, 0x0e, 0xa2, 0xd2, 0x24, 0x99, 0x81, 0x6c, 0xbc, 0x40,
	0x38, 0x7a, 0x60, 0xa3, 0x6b, 0xbc, 0x0a, 0x1b, 0xfa, 0x49, 0xcc, 0x5f, 0x01, 0x1c, 0xdd, 0x5b,
	0x6e, 0xa2, 0x5e, 0x06, 0xf7, 0x3a, 0x9b, 0x20, 0x37, 0x5e, 0xa4, 0xf8, 0x3b, 0x6d, 0x70, 0x0f,
	0x60, 0x20, 0xd0, 0x25, 0xec, 0x77, 0xf3, 0xcb, 0xe2, 0xf4, 0x2f, 0x74, 0xa9, 0xda, 0x96, 0xd7,
	0x4e, 0xaa, 0x35, 0x8d, 0xf0, 0x8f, 0x64, 0xbd, 0x7f, 0x93, 0x9d, 0xc0, 0x51, 0xd1, 0x50, 0xbe,
	0xf1, 0xdc, 0x3a, 0x74, 0x00, 0x13, 0xd9, 0x84, 0x58, 0x3d, 0x9a, 0xc8, 0x26, 0xdf, 0xc0, 0x54,
	0x19, 0x81, 0x7f, 0xb7, 0x99, 0x8f, 0x4b, 0xcf, 0x56, 0x9e, 0x95, 0x5d, 0x4d, 0x25, 0x78, 0xbc,
	0x16, 0xd2, 0x3d, 0x79, 0x86, 0x6b, 0xf5, 0x4c, 0x02, 0x9b, 0x77, 0x75, 0x0a, 0x75, 0x2e, 0xf8,
	0x3a, 0x54, 0x49, 0x76, 0x0a, 0x95, 0x96, 0x96, 0x44, 0x95, 0xab, 0x38, 0xde, 0x01, 0x78, 0x4b,
	0x86, 0xb7, 0xc5, 0xf2, 0xee, 0xe1, 0xa6, 0x2c, 0xd8, 0x20, 0x7c, 0xb8, 0xf8, 0x08, 0x00, 0x00,
	0xff, 0xff, 0x28, 0x19, 0x5a, 0x39, 0xd1, 0x01, 0x00, 0x00,
}
