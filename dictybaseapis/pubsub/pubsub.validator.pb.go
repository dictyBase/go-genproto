// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: pubsub.proto

/*
Package pubsub is a generated protocol buffer package.

It is generated from these files:
	pubsub.proto

It has these top-level messages:
	Reply
	IdentityReply
	IdentityReq
	UserReply
	IdRequest
*/
package pubsub

import github_com_mwitkow_go_proto_validators "github.com/mwitkow/go-proto-validators"
import proto "github.com/golang/protobuf/proto"
import fmt "fmt"
import math "math"
import _ "google.golang.org/genproto/googleapis/rpc/status"
import _ "github.com/dictyBase/go-genproto/dictybaseapis/user"
import _ "github.com/dictyBase/go-genproto/dictybaseapis/identity"

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

func (this *Reply) Validate() error {
	if this.Status != nil {
		if err := github_com_mwitkow_go_proto_validators.CallValidatorIfExists(this.Status); err != nil {
			return github_com_mwitkow_go_proto_validators.FieldError("Status", err)
		}
	}
	return nil
}
func (this *IdentityReply) Validate() error {
	if this.Identity != nil {
		if err := github_com_mwitkow_go_proto_validators.CallValidatorIfExists(this.Identity); err != nil {
			return github_com_mwitkow_go_proto_validators.FieldError("Identity", err)
		}
	}
	if this.Status != nil {
		if err := github_com_mwitkow_go_proto_validators.CallValidatorIfExists(this.Status); err != nil {
			return github_com_mwitkow_go_proto_validators.FieldError("Status", err)
		}
	}
	return nil
}
func (this *IdentityReq) Validate() error {
	return nil
}
func (this *UserReply) Validate() error {
	if this.User != nil {
		if err := github_com_mwitkow_go_proto_validators.CallValidatorIfExists(this.User); err != nil {
			return github_com_mwitkow_go_proto_validators.FieldError("User", err)
		}
	}
	if this.Users != nil {
		if err := github_com_mwitkow_go_proto_validators.CallValidatorIfExists(this.Users); err != nil {
			return github_com_mwitkow_go_proto_validators.FieldError("Users", err)
		}
	}
	if this.Status != nil {
		if err := github_com_mwitkow_go_proto_validators.CallValidatorIfExists(this.Status); err != nil {
			return github_com_mwitkow_go_proto_validators.FieldError("Status", err)
		}
	}
	return nil
}
func (this *IdRequest) Validate() error {
	return nil
}
