// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: dictybase/identity/identity.proto

package identity

import (
	fmt "fmt"
	math "math"
	proto "github.com/golang/protobuf/proto"
	_ "github.com/dictyBase/go-genproto/dictybaseapis/api/jsonapi"
	_ "google.golang.org/protobuf/types/known/emptypb"
	_ "google.golang.org/protobuf/types/known/timestamppb"
	github_com_mwitkow_go_proto_validators "github.com/mwitkow/go-proto-validators"
)

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

func (this *IdentityAttributes) Validate() error {
	if this.CreatedAt != nil {
		if err := github_com_mwitkow_go_proto_validators.CallValidatorIfExists(this.CreatedAt); err != nil {
			return github_com_mwitkow_go_proto_validators.FieldError("CreatedAt", err)
		}
	}
	if this.UpdatedAt != nil {
		if err := github_com_mwitkow_go_proto_validators.CallValidatorIfExists(this.UpdatedAt); err != nil {
			return github_com_mwitkow_go_proto_validators.FieldError("UpdatedAt", err)
		}
	}
	return nil
}
func (this *IdentityData) Validate() error {
	if this.Attributes != nil {
		if err := github_com_mwitkow_go_proto_validators.CallValidatorIfExists(this.Attributes); err != nil {
			return github_com_mwitkow_go_proto_validators.FieldError("Attributes", err)
		}
	}
	if this.Links != nil {
		if err := github_com_mwitkow_go_proto_validators.CallValidatorIfExists(this.Links); err != nil {
			return github_com_mwitkow_go_proto_validators.FieldError("Links", err)
		}
	}
	return nil
}
func (this *Identity) Validate() error {
	if this.Data != nil {
		if err := github_com_mwitkow_go_proto_validators.CallValidatorIfExists(this.Data); err != nil {
			return github_com_mwitkow_go_proto_validators.FieldError("Data", err)
		}
	}
	if this.Links != nil {
		if err := github_com_mwitkow_go_proto_validators.CallValidatorIfExists(this.Links); err != nil {
			return github_com_mwitkow_go_proto_validators.FieldError("Links", err)
		}
	}
	return nil
}
func (this *IdentityProviderReq) Validate() error {
	return nil
}
func (this *CreateIdentityReq) Validate() error {
	if this.Data != nil {
		if err := github_com_mwitkow_go_proto_validators.CallValidatorIfExists(this.Data); err != nil {
			return github_com_mwitkow_go_proto_validators.FieldError("Data", err)
		}
	}
	return nil
}
func (this *CreateIdentityReq_Data) Validate() error {
	if this.Attributes != nil {
		if err := github_com_mwitkow_go_proto_validators.CallValidatorIfExists(this.Attributes); err != nil {
			return github_com_mwitkow_go_proto_validators.FieldError("Attributes", err)
		}
	}
	return nil
}
func (this *NewIdentityAttributes) Validate() error {
	return nil
}
