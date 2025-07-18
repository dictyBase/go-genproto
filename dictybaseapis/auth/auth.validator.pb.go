// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: dictybase/auth/auth.proto

package auth

import (
	fmt "fmt"
	math "math"
	proto "github.com/golang/protobuf/proto"
	_ "github.com/dictyBase/go-genproto/dictybaseapis/user"
	_ "github.com/mwitkow/go-proto-validators"
	_ "google.golang.org/protobuf/types/known/emptypb"
	_ "github.com/dictyBase/go-genproto/dictybaseapis/identity"
	github_com_mwitkow_go_proto_validators "github.com/mwitkow/go-proto-validators"
)

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

func (this *Auth) Validate() error {
	if this.Identity != nil {
		if err := github_com_mwitkow_go_proto_validators.CallValidatorIfExists(this.Identity); err != nil {
			return github_com_mwitkow_go_proto_validators.FieldError("Identity", err)
		}
	}
	if this.User != nil {
		if err := github_com_mwitkow_go_proto_validators.CallValidatorIfExists(this.User); err != nil {
			return github_com_mwitkow_go_proto_validators.FieldError("User", err)
		}
	}
	return nil
}
func (this *NewLogin) Validate() error {
	if this.ClientId == "" {
		return github_com_mwitkow_go_proto_validators.FieldError("ClientId", fmt.Errorf(`value '%v' must not be an empty string`, this.ClientId))
	}
	if this.Scopes == "" {
		return github_com_mwitkow_go_proto_validators.FieldError("Scopes", fmt.Errorf(`value '%v' must not be an empty string`, this.Scopes))
	}
	if this.State == "" {
		return github_com_mwitkow_go_proto_validators.FieldError("State", fmt.Errorf(`value '%v' must not be an empty string`, this.State))
	}
	if this.RedirectUrl == "" {
		return github_com_mwitkow_go_proto_validators.FieldError("RedirectUrl", fmt.Errorf(`value '%v' must not be an empty string`, this.RedirectUrl))
	}
	if this.Code == "" {
		return github_com_mwitkow_go_proto_validators.FieldError("Code", fmt.Errorf(`value '%v' must not be an empty string`, this.Code))
	}
	if this.Provider == "" {
		return github_com_mwitkow_go_proto_validators.FieldError("Provider", fmt.Errorf(`value '%v' must not be an empty string`, this.Provider))
	}
	return nil
}
func (this *NewRelogin) Validate() error {
	if this.RefreshToken == "" {
		return github_com_mwitkow_go_proto_validators.FieldError("RefreshToken", fmt.Errorf(`value '%v' must not be an empty string`, this.RefreshToken))
	}
	return nil
}
func (this *NewToken) Validate() error {
	return nil
}
func (this *NewRefreshToken) Validate() error {
	return nil
}
func (this *Token) Validate() error {
	return nil
}
