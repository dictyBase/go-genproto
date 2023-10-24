// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: dictybase/user/user.proto

package user

import (
	fmt "fmt"
	math "math"
	proto "github.com/golang/protobuf/proto"
	_ "github.com/dictyBase/go-genproto/dictybaseapis/api/jsonapi"
	_ "google.golang.org/protobuf/types/known/fieldmaskpb"
	_ "google.golang.org/protobuf/types/known/timestamppb"
	_ "google.golang.org/protobuf/types/known/emptypb"
	_ "google.golang.org/protobuf/types/known/anypb"
	github_com_mwitkow_go_proto_validators "github.com/mwitkow/go-proto-validators"
)

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

func (this *UpdateUserRequest) Validate() error {
	if this.Data != nil {
		if err := github_com_mwitkow_go_proto_validators.CallValidatorIfExists(this.Data); err != nil {
			return github_com_mwitkow_go_proto_validators.FieldError("Data", err)
		}
	}
	if this.UpdateMask != nil {
		if err := github_com_mwitkow_go_proto_validators.CallValidatorIfExists(this.UpdateMask); err != nil {
			return github_com_mwitkow_go_proto_validators.FieldError("UpdateMask", err)
		}
	}
	return nil
}
func (this *UpdateUserRequest_Data) Validate() error {
	if this.Attributes != nil {
		if err := github_com_mwitkow_go_proto_validators.CallValidatorIfExists(this.Attributes); err != nil {
			return github_com_mwitkow_go_proto_validators.FieldError("Attributes", err)
		}
	}
	if this.Relationships != nil {
		if err := github_com_mwitkow_go_proto_validators.CallValidatorIfExists(this.Relationships); err != nil {
			return github_com_mwitkow_go_proto_validators.FieldError("Relationships", err)
		}
	}
	return nil
}
func (this *CreateUserRequest) Validate() error {
	if this.Data != nil {
		if err := github_com_mwitkow_go_proto_validators.CallValidatorIfExists(this.Data); err != nil {
			return github_com_mwitkow_go_proto_validators.FieldError("Data", err)
		}
	}
	return nil
}
func (this *CreateUserRequest_Data) Validate() error {
	if this.Attributes != nil {
		if err := github_com_mwitkow_go_proto_validators.CallValidatorIfExists(this.Attributes); err != nil {
			return github_com_mwitkow_go_proto_validators.FieldError("Attributes", err)
		}
	}
	if this.Relationships != nil {
		if err := github_com_mwitkow_go_proto_validators.CallValidatorIfExists(this.Relationships); err != nil {
			return github_com_mwitkow_go_proto_validators.FieldError("Relationships", err)
		}
	}
	return nil
}
func (this *User) Validate() error {
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
	for _, item := range this.Included {
		if item != nil {
			if err := github_com_mwitkow_go_proto_validators.CallValidatorIfExists(item); err != nil {
				return github_com_mwitkow_go_proto_validators.FieldError("Included", err)
			}
		}
	}
	return nil
}
func (this *UserCollection) Validate() error {
	for _, item := range this.Data {
		if item != nil {
			if err := github_com_mwitkow_go_proto_validators.CallValidatorIfExists(item); err != nil {
				return github_com_mwitkow_go_proto_validators.FieldError("Data", err)
			}
		}
	}
	if this.Links != nil {
		if err := github_com_mwitkow_go_proto_validators.CallValidatorIfExists(this.Links); err != nil {
			return github_com_mwitkow_go_proto_validators.FieldError("Links", err)
		}
	}
	if this.Meta != nil {
		if err := github_com_mwitkow_go_proto_validators.CallValidatorIfExists(this.Meta); err != nil {
			return github_com_mwitkow_go_proto_validators.FieldError("Meta", err)
		}
	}
	for _, item := range this.Included {
		if item != nil {
			if err := github_com_mwitkow_go_proto_validators.CallValidatorIfExists(item); err != nil {
				return github_com_mwitkow_go_proto_validators.FieldError("Included", err)
			}
		}
	}
	return nil
}
func (this *UserData) Validate() error {
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
	if this.Relationships != nil {
		if err := github_com_mwitkow_go_proto_validators.CallValidatorIfExists(this.Relationships); err != nil {
			return github_com_mwitkow_go_proto_validators.FieldError("Relationships", err)
		}
	}
	return nil
}
func (this *UserAttributes) Validate() error {
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
func (this *ExistingUserRelationships) Validate() error {
	if this.Roles != nil {
		if err := github_com_mwitkow_go_proto_validators.CallValidatorIfExists(this.Roles); err != nil {
			return github_com_mwitkow_go_proto_validators.FieldError("Roles", err)
		}
	}
	return nil
}
func (this *ExistingUserRelationships_Roles) Validate() error {
	if this.Links != nil {
		if err := github_com_mwitkow_go_proto_validators.CallValidatorIfExists(this.Links); err != nil {
			return github_com_mwitkow_go_proto_validators.FieldError("Links", err)
		}
	}
	for _, item := range this.Data {
		if item != nil {
			if err := github_com_mwitkow_go_proto_validators.CallValidatorIfExists(item); err != nil {
				return github_com_mwitkow_go_proto_validators.FieldError("Data", err)
			}
		}
	}
	return nil
}
func (this *NewUserRelationships) Validate() error {
	if this.Roles != nil {
		if err := github_com_mwitkow_go_proto_validators.CallValidatorIfExists(this.Roles); err != nil {
			return github_com_mwitkow_go_proto_validators.FieldError("Roles", err)
		}
	}
	return nil
}
func (this *NewUserRelationships_Roles) Validate() error {
	for _, item := range this.Data {
		if item != nil {
			if err := github_com_mwitkow_go_proto_validators.CallValidatorIfExists(item); err != nil {
				return github_com_mwitkow_go_proto_validators.FieldError("Data", err)
			}
		}
	}
	return nil
}
func (this *UpdateRoleRequest) Validate() error {
	if this.Data != nil {
		if err := github_com_mwitkow_go_proto_validators.CallValidatorIfExists(this.Data); err != nil {
			return github_com_mwitkow_go_proto_validators.FieldError("Data", err)
		}
	}
	if this.UpdateMask != nil {
		if err := github_com_mwitkow_go_proto_validators.CallValidatorIfExists(this.UpdateMask); err != nil {
			return github_com_mwitkow_go_proto_validators.FieldError("UpdateMask", err)
		}
	}
	return nil
}
func (this *UpdateRoleRequest_Data) Validate() error {
	if this.Attributes != nil {
		if err := github_com_mwitkow_go_proto_validators.CallValidatorIfExists(this.Attributes); err != nil {
			return github_com_mwitkow_go_proto_validators.FieldError("Attributes", err)
		}
	}
	if this.Relationships != nil {
		if err := github_com_mwitkow_go_proto_validators.CallValidatorIfExists(this.Relationships); err != nil {
			return github_com_mwitkow_go_proto_validators.FieldError("Relationships", err)
		}
	}
	return nil
}
func (this *CreateRoleRequest) Validate() error {
	if this.Data != nil {
		if err := github_com_mwitkow_go_proto_validators.CallValidatorIfExists(this.Data); err != nil {
			return github_com_mwitkow_go_proto_validators.FieldError("Data", err)
		}
	}
	return nil
}
func (this *CreateRoleRequest_Data) Validate() error {
	if this.Attributes != nil {
		if err := github_com_mwitkow_go_proto_validators.CallValidatorIfExists(this.Attributes); err != nil {
			return github_com_mwitkow_go_proto_validators.FieldError("Attributes", err)
		}
	}
	if this.Relationships != nil {
		if err := github_com_mwitkow_go_proto_validators.CallValidatorIfExists(this.Relationships); err != nil {
			return github_com_mwitkow_go_proto_validators.FieldError("Relationships", err)
		}
	}
	return nil
}
func (this *Role) Validate() error {
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
	for _, item := range this.Included {
		if item != nil {
			if err := github_com_mwitkow_go_proto_validators.CallValidatorIfExists(item); err != nil {
				return github_com_mwitkow_go_proto_validators.FieldError("Included", err)
			}
		}
	}
	return nil
}
func (this *RoleCollection) Validate() error {
	for _, item := range this.Data {
		if item != nil {
			if err := github_com_mwitkow_go_proto_validators.CallValidatorIfExists(item); err != nil {
				return github_com_mwitkow_go_proto_validators.FieldError("Data", err)
			}
		}
	}
	if this.Links != nil {
		if err := github_com_mwitkow_go_proto_validators.CallValidatorIfExists(this.Links); err != nil {
			return github_com_mwitkow_go_proto_validators.FieldError("Links", err)
		}
	}
	for _, item := range this.Included {
		if item != nil {
			if err := github_com_mwitkow_go_proto_validators.CallValidatorIfExists(item); err != nil {
				return github_com_mwitkow_go_proto_validators.FieldError("Included", err)
			}
		}
	}
	return nil
}
func (this *RoleData) Validate() error {
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
	if this.Relationships != nil {
		if err := github_com_mwitkow_go_proto_validators.CallValidatorIfExists(this.Relationships); err != nil {
			return github_com_mwitkow_go_proto_validators.FieldError("Relationships", err)
		}
	}
	return nil
}
func (this *RoleAttributes) Validate() error {
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
func (this *ExistingRoleRelationships) Validate() error {
	if this.Permissions != nil {
		if err := github_com_mwitkow_go_proto_validators.CallValidatorIfExists(this.Permissions); err != nil {
			return github_com_mwitkow_go_proto_validators.FieldError("Permissions", err)
		}
	}
	if this.Users != nil {
		if err := github_com_mwitkow_go_proto_validators.CallValidatorIfExists(this.Users); err != nil {
			return github_com_mwitkow_go_proto_validators.FieldError("Users", err)
		}
	}
	return nil
}
func (this *ExistingRoleRelationships_Permissions) Validate() error {
	if this.Links != nil {
		if err := github_com_mwitkow_go_proto_validators.CallValidatorIfExists(this.Links); err != nil {
			return github_com_mwitkow_go_proto_validators.FieldError("Links", err)
		}
	}
	for _, item := range this.Data {
		if item != nil {
			if err := github_com_mwitkow_go_proto_validators.CallValidatorIfExists(item); err != nil {
				return github_com_mwitkow_go_proto_validators.FieldError("Data", err)
			}
		}
	}
	return nil
}
func (this *ExistingRoleRelationships_Users) Validate() error {
	if this.Links != nil {
		if err := github_com_mwitkow_go_proto_validators.CallValidatorIfExists(this.Links); err != nil {
			return github_com_mwitkow_go_proto_validators.FieldError("Links", err)
		}
	}
	for _, item := range this.Data {
		if item != nil {
			if err := github_com_mwitkow_go_proto_validators.CallValidatorIfExists(item); err != nil {
				return github_com_mwitkow_go_proto_validators.FieldError("Data", err)
			}
		}
	}
	return nil
}
func (this *NewRoleRelationships) Validate() error {
	if this.Permissions != nil {
		if err := github_com_mwitkow_go_proto_validators.CallValidatorIfExists(this.Permissions); err != nil {
			return github_com_mwitkow_go_proto_validators.FieldError("Permissions", err)
		}
	}
	if this.Users != nil {
		if err := github_com_mwitkow_go_proto_validators.CallValidatorIfExists(this.Users); err != nil {
			return github_com_mwitkow_go_proto_validators.FieldError("Users", err)
		}
	}
	return nil
}
func (this *NewRoleRelationships_Permissions) Validate() error {
	for _, item := range this.Data {
		if item != nil {
			if err := github_com_mwitkow_go_proto_validators.CallValidatorIfExists(item); err != nil {
				return github_com_mwitkow_go_proto_validators.FieldError("Data", err)
			}
		}
	}
	return nil
}
func (this *NewRoleRelationships_Users) Validate() error {
	for _, item := range this.Data {
		if item != nil {
			if err := github_com_mwitkow_go_proto_validators.CallValidatorIfExists(item); err != nil {
				return github_com_mwitkow_go_proto_validators.FieldError("Data", err)
			}
		}
	}
	return nil
}
func (this *UpdatePermissionRequest) Validate() error {
	if this.Data != nil {
		if err := github_com_mwitkow_go_proto_validators.CallValidatorIfExists(this.Data); err != nil {
			return github_com_mwitkow_go_proto_validators.FieldError("Data", err)
		}
	}
	if this.UpdateMask != nil {
		if err := github_com_mwitkow_go_proto_validators.CallValidatorIfExists(this.UpdateMask); err != nil {
			return github_com_mwitkow_go_proto_validators.FieldError("UpdateMask", err)
		}
	}
	return nil
}
func (this *UpdatePermissionRequest_Data) Validate() error {
	if this.Attributes != nil {
		if err := github_com_mwitkow_go_proto_validators.CallValidatorIfExists(this.Attributes); err != nil {
			return github_com_mwitkow_go_proto_validators.FieldError("Attributes", err)
		}
	}
	return nil
}
func (this *CreatePermissionRequest) Validate() error {
	if this.Data != nil {
		if err := github_com_mwitkow_go_proto_validators.CallValidatorIfExists(this.Data); err != nil {
			return github_com_mwitkow_go_proto_validators.FieldError("Data", err)
		}
	}
	return nil
}
func (this *CreatePermissionRequest_Data) Validate() error {
	if this.Attributes != nil {
		if err := github_com_mwitkow_go_proto_validators.CallValidatorIfExists(this.Attributes); err != nil {
			return github_com_mwitkow_go_proto_validators.FieldError("Attributes", err)
		}
	}
	return nil
}
func (this *Permission) Validate() error {
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
func (this *PermissionCollection) Validate() error {
	for _, item := range this.Data {
		if item != nil {
			if err := github_com_mwitkow_go_proto_validators.CallValidatorIfExists(item); err != nil {
				return github_com_mwitkow_go_proto_validators.FieldError("Data", err)
			}
		}
	}
	if this.Links != nil {
		if err := github_com_mwitkow_go_proto_validators.CallValidatorIfExists(this.Links); err != nil {
			return github_com_mwitkow_go_proto_validators.FieldError("Links", err)
		}
	}
	return nil
}
func (this *PermissionData) Validate() error {
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
func (this *PermissionAttributes) Validate() error {
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
