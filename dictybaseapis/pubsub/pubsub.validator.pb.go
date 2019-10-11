// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: pubsub.proto

package pubsub

import (
	fmt "fmt"
	math "math"
	proto "github.com/golang/protobuf/proto"
	_ "github.com/dictyBase/go-genproto/dictybaseapis/user"
	_ "github.com/dictyBase/go-genproto/dictybaseapis/identity"
	_ "google.golang.org/genproto/googleapis/rpc/status"
	github_com_mwitkow_go_proto_validators "github.com/mwitkow/go-proto-validators"
)

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
