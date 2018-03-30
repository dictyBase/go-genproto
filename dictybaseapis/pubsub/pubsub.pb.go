// Code generated by protoc-gen-go. DO NOT EDIT.
// source: pubsub.proto

/*
Package pubsub is a generated protocol buffer package.

It is generated from these files:
	pubsub.proto

It has these top-level messages:
	Reply
	UserReply
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

func init() {
	proto.RegisterType((*Reply)(nil), "dictybase.pubsub.Reply")
	proto.RegisterType((*UserReply)(nil), "dictybase.pubsub.UserReply")
}

func init() { proto.RegisterFile("pubsub.proto", fileDescriptor0) }

var fileDescriptor0 = []byte{
	// 279 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x84, 0x90, 0x3f, 0x4b, 0x03, 0x31,
	0x18, 0xc6, 0x49, 0xb5, 0xa5, 0xe6, 0x1c, 0x24, 0x1c, 0x58, 0x3b, 0x48, 0xe9, 0x74, 0x08, 0x26,
	0x50, 0xdd, 0x74, 0xf1, 0xea, 0x72, 0xdb, 0x71, 0xd5, 0x41, 0xb7, 0x4b, 0x0c, 0x31, 0x70, 0x36,
	0x31, 0x7f, 0xc0, 0x7e, 0x12, 0x77, 0x3f, 0xa5, 0xa3, 0x5c, 0x5e, 0xad, 0x22, 0x82, 0x4b, 0x5e,
	0x92, 0xf7, 0x97, 0xe7, 0x79, 0x78, 0xf0, 0xbe, 0x8d, 0xdc, 0x47, 0x4e, 0xad, 0x33, 0xc1, 0x90,
	0x83, 0x07, 0x2d, 0xc2, 0x86, 0xb7, 0x5e, 0x52, 0x78, 0x9f, 0x1e, 0x2a, 0x63, 0x54, 0x27, 0x99,
	0xb3, 0x82, 0xf9, 0xd0, 0x86, 0xe8, 0x01, 0x9d, 0x1e, 0x6d, 0x51, 0x16, 0xbd, 0x74, 0xe9, 0x80,
	0xd5, 0xbc, 0xc2, 0xc3, 0x46, 0xda, 0x6e, 0x43, 0x72, 0x3c, 0x94, 0x2f, 0xda, 0x87, 0x09, 0x9a,
	0xa1, 0x62, 0xdc, 0xc0, 0x85, 0x9c, 0xe0, 0x11, 0x28, 0x4d, 0x06, 0x33, 0x54, 0x64, 0x0b, 0x42,
	0xc1, 0x83, 0x3a, 0x2b, 0xe8, 0x2a, 0x6d, 0x9a, 0x4f, 0x62, 0xfe, 0x8a, 0xf0, 0xde, 0xad, 0x97,
	0x0e, 0xf4, 0x0a, 0xbc, 0xdb, 0xdb, 0x24, 0xb9, 0x6c, 0x91, 0xd3, 0xef, 0xb4, 0xc9, 0x3d, 0x81,
	0x89, 0x20, 0xe7, 0x78, 0xd8, 0xcf, 0x2f, 0x8b, 0xe3, 0xbf, 0xd0, 0xa5, 0xe9, 0x3a, 0x29, 0x82,
	0x36, 0xeb, 0x06, 0xe0, 0x1f, 0xc9, 0x76, 0xfe, 0x4b, 0x56, 0x3e, 0xe3, 0xdc, 0x38, 0x45, 0x7f,
	0x17, 0x56, 0x66, 0x75, 0xe4, 0xab, 0xc8, 0xeb, 0xbe, 0x89, 0x1a, 0xdd, 0x5f, 0x2a, 0x1d, 0x1e,
	0x23, 0xa7, 0xc2, 0x3c, 0xb1, 0xc4, 0x96, 0x7d, 0x63, 0xca, 0x9c, 0x2a, 0xb9, 0x4e, 0x6d, 0xb1,
	0xad, 0x42, 0x6b, 0xb5, 0x67, 0xa0, 0x72, 0x01, 0xe3, 0x1d, 0xa1, 0xb7, 0xc1, 0xf8, 0xba, 0x5a,
	0xde, 0xdc, 0x5d, 0xd5, 0x15, 0x1f, 0xa5, 0x0f, 0x67, 0x1f, 0x01, 0x00, 0x00, 0xff, 0xff, 0x69,
	0xbe, 0x36, 0x24, 0xb4, 0x01, 0x00, 0x00,
}
