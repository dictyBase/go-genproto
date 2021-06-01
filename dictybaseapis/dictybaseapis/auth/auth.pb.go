// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.26.0
// 	protoc        v3.17.1
// source: dictybase/auth/auth.proto

package auth

import (
	identity "github.com/dictyBase/go-genproto/dictybaseapis/identity"
	user "github.com/dictyBase/go-genproto/dictybaseapis/user"
	_ "github.com/mwitkow/go-proto-validators"
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

// Definition of an individual auth response
type Auth struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// JSON Web Token (JWT)
	Token string `protobuf:"bytes,1,opt,name=token,proto3" json:"token,omitempty"`
	// Refresh token
	RefreshToken string `protobuf:"bytes,2,opt,name=refresh_token,json=refreshToken,proto3" json:"refresh_token,omitempty"`
	// Identity of user
	Identity *identity.Identity `protobuf:"bytes,3,opt,name=identity,proto3" json:"identity,omitempty"`
	// User API data
	User *user.User `protobuf:"bytes,4,opt,name=user,proto3" json:"user,omitempty"`
}

func (x *Auth) Reset() {
	*x = Auth{}
	if protoimpl.UnsafeEnabled {
		mi := &file_dictybase_auth_auth_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Auth) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Auth) ProtoMessage() {}

func (x *Auth) ProtoReflect() protoreflect.Message {
	mi := &file_dictybase_auth_auth_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Auth.ProtoReflect.Descriptor instead.
func (*Auth) Descriptor() ([]byte, []int) {
	return file_dictybase_auth_auth_proto_rawDescGZIP(), []int{0}
}

func (x *Auth) GetToken() string {
	if x != nil {
		return x.Token
	}
	return ""
}

func (x *Auth) GetRefreshToken() string {
	if x != nil {
		return x.RefreshToken
	}
	return ""
}

func (x *Auth) GetIdentity() *identity.Identity {
	if x != nil {
		return x.Identity
	}
	return nil
}

func (x *Auth) GetUser() *user.User {
	if x != nil {
		return x.User
	}
	return nil
}

type NewLogin struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// Client ID received during application registration from every provider
	ClientId string `protobuf:"bytes,1,opt,name=client_id,json=clientId,proto3" json:"client_id,omitempty"`
	// Scope of the application
	Scopes string `protobuf:"bytes,2,opt,name=scopes,proto3" json:"scopes,omitempty"`
	// An unguessable random string. It is used to protect against cross-site request
	// forgery attacks. It is passed to the provider during first login.
	State string `protobuf:"bytes,3,opt,name=state,proto3" json:"state,omitempty"`
	// The URL in the application where users will be sent after authorization,
	// generally provided during the registration of the application.
	RedirectUrl string `protobuf:"bytes,4,opt,name=redirect_url,json=redirectUrl,proto3" json:"redirect_url,omitempty"`
	// The code that is received as response from the first login
	Code string `protobuf:"bytes,5,opt,name=code,proto3" json:"code,omitempty"`
	// Third party oAuth provider
	Provider string `protobuf:"bytes,6,opt,name=provider,proto3" json:"provider,omitempty"`
}

func (x *NewLogin) Reset() {
	*x = NewLogin{}
	if protoimpl.UnsafeEnabled {
		mi := &file_dictybase_auth_auth_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *NewLogin) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*NewLogin) ProtoMessage() {}

func (x *NewLogin) ProtoReflect() protoreflect.Message {
	mi := &file_dictybase_auth_auth_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use NewLogin.ProtoReflect.Descriptor instead.
func (*NewLogin) Descriptor() ([]byte, []int) {
	return file_dictybase_auth_auth_proto_rawDescGZIP(), []int{1}
}

func (x *NewLogin) GetClientId() string {
	if x != nil {
		return x.ClientId
	}
	return ""
}

func (x *NewLogin) GetScopes() string {
	if x != nil {
		return x.Scopes
	}
	return ""
}

func (x *NewLogin) GetState() string {
	if x != nil {
		return x.State
	}
	return ""
}

func (x *NewLogin) GetRedirectUrl() string {
	if x != nil {
		return x.RedirectUrl
	}
	return ""
}

func (x *NewLogin) GetCode() string {
	if x != nil {
		return x.Code
	}
	return ""
}

func (x *NewLogin) GetProvider() string {
	if x != nil {
		return x.Provider
	}
	return ""
}

type NewRelogin struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// Refresh token
	RefreshToken string `protobuf:"bytes,1,opt,name=refresh_token,json=refreshToken,proto3" json:"refresh_token,omitempty"`
}

func (x *NewRelogin) Reset() {
	*x = NewRelogin{}
	if protoimpl.UnsafeEnabled {
		mi := &file_dictybase_auth_auth_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *NewRelogin) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*NewRelogin) ProtoMessage() {}

func (x *NewRelogin) ProtoReflect() protoreflect.Message {
	mi := &file_dictybase_auth_auth_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use NewRelogin.ProtoReflect.Descriptor instead.
func (*NewRelogin) Descriptor() ([]byte, []int) {
	return file_dictybase_auth_auth_proto_rawDescGZIP(), []int{2}
}

func (x *NewRelogin) GetRefreshToken() string {
	if x != nil {
		return x.RefreshToken
	}
	return ""
}

type NewToken struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// JSON Web Token (JWT)
	Token string `protobuf:"bytes,1,opt,name=token,proto3" json:"token,omitempty"`
	// Refresh token (unique ID)
	RefreshToken string `protobuf:"bytes,2,opt,name=refresh_token,json=refreshToken,proto3" json:"refresh_token,omitempty"`
}

func (x *NewToken) Reset() {
	*x = NewToken{}
	if protoimpl.UnsafeEnabled {
		mi := &file_dictybase_auth_auth_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *NewToken) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*NewToken) ProtoMessage() {}

func (x *NewToken) ProtoReflect() protoreflect.Message {
	mi := &file_dictybase_auth_auth_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use NewToken.ProtoReflect.Descriptor instead.
func (*NewToken) Descriptor() ([]byte, []int) {
	return file_dictybase_auth_auth_proto_rawDescGZIP(), []int{3}
}

func (x *NewToken) GetToken() string {
	if x != nil {
		return x.Token
	}
	return ""
}

func (x *NewToken) GetRefreshToken() string {
	if x != nil {
		return x.RefreshToken
	}
	return ""
}

type NewRefreshToken struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// Refresh token (unique ID)
	RefreshToken string `protobuf:"bytes,1,opt,name=refresh_token,json=refreshToken,proto3" json:"refresh_token,omitempty"`
}

func (x *NewRefreshToken) Reset() {
	*x = NewRefreshToken{}
	if protoimpl.UnsafeEnabled {
		mi := &file_dictybase_auth_auth_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *NewRefreshToken) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*NewRefreshToken) ProtoMessage() {}

func (x *NewRefreshToken) ProtoReflect() protoreflect.Message {
	mi := &file_dictybase_auth_auth_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use NewRefreshToken.ProtoReflect.Descriptor instead.
func (*NewRefreshToken) Descriptor() ([]byte, []int) {
	return file_dictybase_auth_auth_proto_rawDescGZIP(), []int{4}
}

func (x *NewRefreshToken) GetRefreshToken() string {
	if x != nil {
		return x.RefreshToken
	}
	return ""
}

type Token struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// JSON Web Token (JWT)
	Token string `protobuf:"bytes,1,opt,name=token,proto3" json:"token,omitempty"`
	// Refresh token (unique ID)
	RefreshToken string `protobuf:"bytes,2,opt,name=refresh_token,json=refreshToken,proto3" json:"refresh_token,omitempty"`
}

func (x *Token) Reset() {
	*x = Token{}
	if protoimpl.UnsafeEnabled {
		mi := &file_dictybase_auth_auth_proto_msgTypes[5]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Token) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Token) ProtoMessage() {}

func (x *Token) ProtoReflect() protoreflect.Message {
	mi := &file_dictybase_auth_auth_proto_msgTypes[5]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Token.ProtoReflect.Descriptor instead.
func (*Token) Descriptor() ([]byte, []int) {
	return file_dictybase_auth_auth_proto_rawDescGZIP(), []int{5}
}

func (x *Token) GetToken() string {
	if x != nil {
		return x.Token
	}
	return ""
}

func (x *Token) GetRefreshToken() string {
	if x != nil {
		return x.RefreshToken
	}
	return ""
}

var File_dictybase_auth_auth_proto protoreflect.FileDescriptor

var file_dictybase_auth_auth_proto_rawDesc = []byte{
	0x0a, 0x19, 0x64, 0x69, 0x63, 0x74, 0x79, 0x62, 0x61, 0x73, 0x65, 0x2f, 0x61, 0x75, 0x74, 0x68,
	0x2f, 0x61, 0x75, 0x74, 0x68, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x0e, 0x64, 0x69, 0x63,
	0x74, 0x79, 0x62, 0x61, 0x73, 0x65, 0x2e, 0x61, 0x75, 0x74, 0x68, 0x1a, 0x1b, 0x67, 0x6f, 0x6f,
	0x67, 0x6c, 0x65, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2f, 0x65, 0x6d, 0x70,
	0x74, 0x79, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x36, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62,
	0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x6d, 0x77, 0x69, 0x74, 0x6b, 0x6f, 0x77, 0x2f, 0x67, 0x6f, 0x2d,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2d, 0x76, 0x61, 0x6c, 0x69, 0x64, 0x61, 0x74, 0x6f, 0x72, 0x73,
	0x2f, 0x76, 0x61, 0x6c, 0x69, 0x64, 0x61, 0x74, 0x6f, 0x72, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x1a, 0x21, 0x64, 0x69, 0x63, 0x74, 0x79, 0x62, 0x61, 0x73, 0x65, 0x2f, 0x69, 0x64, 0x65, 0x6e,
	0x74, 0x69, 0x74, 0x79, 0x2f, 0x69, 0x64, 0x65, 0x6e, 0x74, 0x69, 0x74, 0x79, 0x2e, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x1a, 0x19, 0x64, 0x69, 0x63, 0x74, 0x79, 0x62, 0x61, 0x73, 0x65, 0x2f, 0x75,
	0x73, 0x65, 0x72, 0x2f, 0x75, 0x73, 0x65, 0x72, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0xa5,
	0x01, 0x0a, 0x04, 0x41, 0x75, 0x74, 0x68, 0x12, 0x14, 0x0a, 0x05, 0x74, 0x6f, 0x6b, 0x65, 0x6e,
	0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x74, 0x6f, 0x6b, 0x65, 0x6e, 0x12, 0x23, 0x0a,
	0x0d, 0x72, 0x65, 0x66, 0x72, 0x65, 0x73, 0x68, 0x5f, 0x74, 0x6f, 0x6b, 0x65, 0x6e, 0x18, 0x02,
	0x20, 0x01, 0x28, 0x09, 0x52, 0x0c, 0x72, 0x65, 0x66, 0x72, 0x65, 0x73, 0x68, 0x54, 0x6f, 0x6b,
	0x65, 0x6e, 0x12, 0x38, 0x0a, 0x08, 0x69, 0x64, 0x65, 0x6e, 0x74, 0x69, 0x74, 0x79, 0x18, 0x03,
	0x20, 0x01, 0x28, 0x0b, 0x32, 0x1c, 0x2e, 0x64, 0x69, 0x63, 0x74, 0x79, 0x62, 0x61, 0x73, 0x65,
	0x2e, 0x69, 0x64, 0x65, 0x6e, 0x74, 0x69, 0x74, 0x79, 0x2e, 0x49, 0x64, 0x65, 0x6e, 0x74, 0x69,
	0x74, 0x79, 0x52, 0x08, 0x69, 0x64, 0x65, 0x6e, 0x74, 0x69, 0x74, 0x79, 0x12, 0x28, 0x0a, 0x04,
	0x75, 0x73, 0x65, 0x72, 0x18, 0x04, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x14, 0x2e, 0x64, 0x69, 0x63,
	0x74, 0x79, 0x62, 0x61, 0x73, 0x65, 0x2e, 0x75, 0x73, 0x65, 0x72, 0x2e, 0x55, 0x73, 0x65, 0x72,
	0x52, 0x04, 0x75, 0x73, 0x65, 0x72, 0x22, 0xd8, 0x01, 0x0a, 0x08, 0x4e, 0x65, 0x77, 0x4c, 0x6f,
	0x67, 0x69, 0x6e, 0x12, 0x23, 0x0a, 0x09, 0x63, 0x6c, 0x69, 0x65, 0x6e, 0x74, 0x5f, 0x69, 0x64,
	0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x42, 0x06, 0xe2, 0xdf, 0x1f, 0x02, 0x58, 0x01, 0x52, 0x08,
	0x63, 0x6c, 0x69, 0x65, 0x6e, 0x74, 0x49, 0x64, 0x12, 0x1e, 0x0a, 0x06, 0x73, 0x63, 0x6f, 0x70,
	0x65, 0x73, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x42, 0x06, 0xe2, 0xdf, 0x1f, 0x02, 0x58, 0x01,
	0x52, 0x06, 0x73, 0x63, 0x6f, 0x70, 0x65, 0x73, 0x12, 0x1c, 0x0a, 0x05, 0x73, 0x74, 0x61, 0x74,
	0x65, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x42, 0x06, 0xe2, 0xdf, 0x1f, 0x02, 0x58, 0x01, 0x52,
	0x05, 0x73, 0x74, 0x61, 0x74, 0x65, 0x12, 0x29, 0x0a, 0x0c, 0x72, 0x65, 0x64, 0x69, 0x72, 0x65,
	0x63, 0x74, 0x5f, 0x75, 0x72, 0x6c, 0x18, 0x04, 0x20, 0x01, 0x28, 0x09, 0x42, 0x06, 0xe2, 0xdf,
	0x1f, 0x02, 0x58, 0x01, 0x52, 0x0b, 0x72, 0x65, 0x64, 0x69, 0x72, 0x65, 0x63, 0x74, 0x55, 0x72,
	0x6c, 0x12, 0x1a, 0x0a, 0x04, 0x63, 0x6f, 0x64, 0x65, 0x18, 0x05, 0x20, 0x01, 0x28, 0x09, 0x42,
	0x06, 0xe2, 0xdf, 0x1f, 0x02, 0x58, 0x01, 0x52, 0x04, 0x63, 0x6f, 0x64, 0x65, 0x12, 0x22, 0x0a,
	0x08, 0x70, 0x72, 0x6f, 0x76, 0x69, 0x64, 0x65, 0x72, 0x18, 0x06, 0x20, 0x01, 0x28, 0x09, 0x42,
	0x06, 0xe2, 0xdf, 0x1f, 0x02, 0x58, 0x01, 0x52, 0x08, 0x70, 0x72, 0x6f, 0x76, 0x69, 0x64, 0x65,
	0x72, 0x22, 0x39, 0x0a, 0x0a, 0x4e, 0x65, 0x77, 0x52, 0x65, 0x6c, 0x6f, 0x67, 0x69, 0x6e, 0x12,
	0x2b, 0x0a, 0x0d, 0x72, 0x65, 0x66, 0x72, 0x65, 0x73, 0x68, 0x5f, 0x74, 0x6f, 0x6b, 0x65, 0x6e,
	0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x42, 0x06, 0xe2, 0xdf, 0x1f, 0x02, 0x58, 0x01, 0x52, 0x0c,
	0x72, 0x65, 0x66, 0x72, 0x65, 0x73, 0x68, 0x54, 0x6f, 0x6b, 0x65, 0x6e, 0x22, 0x45, 0x0a, 0x08,
	0x4e, 0x65, 0x77, 0x54, 0x6f, 0x6b, 0x65, 0x6e, 0x12, 0x14, 0x0a, 0x05, 0x74, 0x6f, 0x6b, 0x65,
	0x6e, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x74, 0x6f, 0x6b, 0x65, 0x6e, 0x12, 0x23,
	0x0a, 0x0d, 0x72, 0x65, 0x66, 0x72, 0x65, 0x73, 0x68, 0x5f, 0x74, 0x6f, 0x6b, 0x65, 0x6e, 0x18,
	0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0c, 0x72, 0x65, 0x66, 0x72, 0x65, 0x73, 0x68, 0x54, 0x6f,
	0x6b, 0x65, 0x6e, 0x22, 0x36, 0x0a, 0x0f, 0x4e, 0x65, 0x77, 0x52, 0x65, 0x66, 0x72, 0x65, 0x73,
	0x68, 0x54, 0x6f, 0x6b, 0x65, 0x6e, 0x12, 0x23, 0x0a, 0x0d, 0x72, 0x65, 0x66, 0x72, 0x65, 0x73,
	0x68, 0x5f, 0x74, 0x6f, 0x6b, 0x65, 0x6e, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0c, 0x72,
	0x65, 0x66, 0x72, 0x65, 0x73, 0x68, 0x54, 0x6f, 0x6b, 0x65, 0x6e, 0x22, 0x42, 0x0a, 0x05, 0x54,
	0x6f, 0x6b, 0x65, 0x6e, 0x12, 0x14, 0x0a, 0x05, 0x74, 0x6f, 0x6b, 0x65, 0x6e, 0x18, 0x01, 0x20,
	0x01, 0x28, 0x09, 0x52, 0x05, 0x74, 0x6f, 0x6b, 0x65, 0x6e, 0x12, 0x23, 0x0a, 0x0d, 0x72, 0x65,
	0x66, 0x72, 0x65, 0x73, 0x68, 0x5f, 0x74, 0x6f, 0x6b, 0x65, 0x6e, 0x18, 0x02, 0x20, 0x01, 0x28,
	0x09, 0x52, 0x0c, 0x72, 0x65, 0x66, 0x72, 0x65, 0x73, 0x68, 0x54, 0x6f, 0x6b, 0x65, 0x6e, 0x32,
	0x92, 0x02, 0x0a, 0x0b, 0x41, 0x75, 0x74, 0x68, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x12,
	0x39, 0x0a, 0x05, 0x4c, 0x6f, 0x67, 0x69, 0x6e, 0x12, 0x18, 0x2e, 0x64, 0x69, 0x63, 0x74, 0x79,
	0x62, 0x61, 0x73, 0x65, 0x2e, 0x61, 0x75, 0x74, 0x68, 0x2e, 0x4e, 0x65, 0x77, 0x4c, 0x6f, 0x67,
	0x69, 0x6e, 0x1a, 0x14, 0x2e, 0x64, 0x69, 0x63, 0x74, 0x79, 0x62, 0x61, 0x73, 0x65, 0x2e, 0x61,
	0x75, 0x74, 0x68, 0x2e, 0x41, 0x75, 0x74, 0x68, 0x22, 0x00, 0x12, 0x3d, 0x0a, 0x07, 0x52, 0x65,
	0x6c, 0x6f, 0x67, 0x69, 0x6e, 0x12, 0x1a, 0x2e, 0x64, 0x69, 0x63, 0x74, 0x79, 0x62, 0x61, 0x73,
	0x65, 0x2e, 0x61, 0x75, 0x74, 0x68, 0x2e, 0x4e, 0x65, 0x77, 0x52, 0x65, 0x6c, 0x6f, 0x67, 0x69,
	0x6e, 0x1a, 0x14, 0x2e, 0x64, 0x69, 0x63, 0x74, 0x79, 0x62, 0x61, 0x73, 0x65, 0x2e, 0x61, 0x75,
	0x74, 0x68, 0x2e, 0x41, 0x75, 0x74, 0x68, 0x22, 0x00, 0x12, 0x44, 0x0a, 0x0f, 0x47, 0x65, 0x74,
	0x52, 0x65, 0x66, 0x72, 0x65, 0x73, 0x68, 0x54, 0x6f, 0x6b, 0x65, 0x6e, 0x12, 0x18, 0x2e, 0x64,
	0x69, 0x63, 0x74, 0x79, 0x62, 0x61, 0x73, 0x65, 0x2e, 0x61, 0x75, 0x74, 0x68, 0x2e, 0x4e, 0x65,
	0x77, 0x54, 0x6f, 0x6b, 0x65, 0x6e, 0x1a, 0x15, 0x2e, 0x64, 0x69, 0x63, 0x74, 0x79, 0x62, 0x61,
	0x73, 0x65, 0x2e, 0x61, 0x75, 0x74, 0x68, 0x2e, 0x54, 0x6f, 0x6b, 0x65, 0x6e, 0x22, 0x00, 0x12,
	0x43, 0x0a, 0x06, 0x4c, 0x6f, 0x67, 0x6f, 0x75, 0x74, 0x12, 0x1f, 0x2e, 0x64, 0x69, 0x63, 0x74,
	0x79, 0x62, 0x61, 0x73, 0x65, 0x2e, 0x61, 0x75, 0x74, 0x68, 0x2e, 0x4e, 0x65, 0x77, 0x52, 0x65,
	0x66, 0x72, 0x65, 0x73, 0x68, 0x54, 0x6f, 0x6b, 0x65, 0x6e, 0x1a, 0x16, 0x2e, 0x67, 0x6f, 0x6f,
	0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x45, 0x6d, 0x70,
	0x74, 0x79, 0x22, 0x00, 0x42, 0x69, 0x0a, 0x12, 0x6f, 0x72, 0x67, 0x2e, 0x64, 0x69, 0x63, 0x74,
	0x79, 0x62, 0x61, 0x73, 0x65, 0x2e, 0x61, 0x75, 0x74, 0x68, 0x42, 0x09, 0x41, 0x75, 0x74, 0x68,
	0x50, 0x72, 0x6f, 0x74, 0x6f, 0x50, 0x01, 0x5a, 0x38, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e,
	0x63, 0x6f, 0x6d, 0x2f, 0x64, 0x69, 0x63, 0x74, 0x79, 0x42, 0x61, 0x73, 0x65, 0x2f, 0x67, 0x6f,
	0x2d, 0x67, 0x65, 0x6e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2f, 0x64, 0x69, 0x63, 0x74, 0x79, 0x62,
	0x61, 0x73, 0x65, 0x61, 0x70, 0x69, 0x73, 0x2f, 0x61, 0x75, 0x74, 0x68, 0x3b, 0x61, 0x75, 0x74,
	0x68, 0xf8, 0x01, 0x01, 0xa2, 0x02, 0x08, 0x44, 0x49, 0x43, 0x54, 0x59, 0x41, 0x50, 0x49, 0x62,
	0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_dictybase_auth_auth_proto_rawDescOnce sync.Once
	file_dictybase_auth_auth_proto_rawDescData = file_dictybase_auth_auth_proto_rawDesc
)

func file_dictybase_auth_auth_proto_rawDescGZIP() []byte {
	file_dictybase_auth_auth_proto_rawDescOnce.Do(func() {
		file_dictybase_auth_auth_proto_rawDescData = protoimpl.X.CompressGZIP(file_dictybase_auth_auth_proto_rawDescData)
	})
	return file_dictybase_auth_auth_proto_rawDescData
}

var file_dictybase_auth_auth_proto_msgTypes = make([]protoimpl.MessageInfo, 6)
var file_dictybase_auth_auth_proto_goTypes = []interface{}{
	(*Auth)(nil),              // 0: dictybase.auth.Auth
	(*NewLogin)(nil),          // 1: dictybase.auth.NewLogin
	(*NewRelogin)(nil),        // 2: dictybase.auth.NewRelogin
	(*NewToken)(nil),          // 3: dictybase.auth.NewToken
	(*NewRefreshToken)(nil),   // 4: dictybase.auth.NewRefreshToken
	(*Token)(nil),             // 5: dictybase.auth.Token
	(*identity.Identity)(nil), // 6: dictybase.identity.Identity
	(*user.User)(nil),         // 7: dictybase.user.User
	(*emptypb.Empty)(nil),     // 8: google.protobuf.Empty
}
var file_dictybase_auth_auth_proto_depIdxs = []int32{
	6, // 0: dictybase.auth.Auth.identity:type_name -> dictybase.identity.Identity
	7, // 1: dictybase.auth.Auth.user:type_name -> dictybase.user.User
	1, // 2: dictybase.auth.AuthService.Login:input_type -> dictybase.auth.NewLogin
	2, // 3: dictybase.auth.AuthService.Relogin:input_type -> dictybase.auth.NewRelogin
	3, // 4: dictybase.auth.AuthService.GetRefreshToken:input_type -> dictybase.auth.NewToken
	4, // 5: dictybase.auth.AuthService.Logout:input_type -> dictybase.auth.NewRefreshToken
	0, // 6: dictybase.auth.AuthService.Login:output_type -> dictybase.auth.Auth
	0, // 7: dictybase.auth.AuthService.Relogin:output_type -> dictybase.auth.Auth
	5, // 8: dictybase.auth.AuthService.GetRefreshToken:output_type -> dictybase.auth.Token
	8, // 9: dictybase.auth.AuthService.Logout:output_type -> google.protobuf.Empty
	6, // [6:10] is the sub-list for method output_type
	2, // [2:6] is the sub-list for method input_type
	2, // [2:2] is the sub-list for extension type_name
	2, // [2:2] is the sub-list for extension extendee
	0, // [0:2] is the sub-list for field type_name
}

func init() { file_dictybase_auth_auth_proto_init() }
func file_dictybase_auth_auth_proto_init() {
	if File_dictybase_auth_auth_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_dictybase_auth_auth_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Auth); i {
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
		file_dictybase_auth_auth_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*NewLogin); i {
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
		file_dictybase_auth_auth_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*NewRelogin); i {
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
		file_dictybase_auth_auth_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*NewToken); i {
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
		file_dictybase_auth_auth_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*NewRefreshToken); i {
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
		file_dictybase_auth_auth_proto_msgTypes[5].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Token); i {
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
			RawDescriptor: file_dictybase_auth_auth_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   6,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_dictybase_auth_auth_proto_goTypes,
		DependencyIndexes: file_dictybase_auth_auth_proto_depIdxs,
		MessageInfos:      file_dictybase_auth_auth_proto_msgTypes,
	}.Build()
	File_dictybase_auth_auth_proto = out.File
	file_dictybase_auth_auth_proto_rawDesc = nil
	file_dictybase_auth_auth_proto_goTypes = nil
	file_dictybase_auth_auth_proto_depIdxs = nil
}
