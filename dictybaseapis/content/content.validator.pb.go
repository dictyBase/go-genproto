// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: content.proto

/*
Package content is a generated protocol buffer package.

It is generated from these files:
	content.proto

It has these top-level messages:
	Content
	ContentData
	ContentAttributes
	ContentRequest
	ContentIdRequest
	NewContentAttributes
	StoreContentRequest
	ExistingContentAttributes
	UpdateContentRequest
*/
package content

import regexp "regexp"
import fmt "fmt"
import github_com_mwitkow_go_proto_validators "github.com/mwitkow/go-proto-validators"
import proto "github.com/golang/protobuf/proto"
import math "math"
import _ "google.golang.org/genproto/googleapis/api/annotations"
import _ "google.golang.org/genproto/protobuf/field_mask"
import _ "github.com/golang/protobuf/ptypes/timestamp"
import _ "github.com/golang/protobuf/ptypes/empty"
import _ "github.com/golang/protobuf/ptypes/any"
import _ "github.com/dictyBase/go-genproto/dictybaseapis/api/jsonapi"
import _ "github.com/dictyBase/go-genproto/dictybaseapis/api/jsonapi"
import _ "github.com/mwitkow/go-proto-validators"

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

func (this *Content) Validate() error {
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
func (this *ContentData) Validate() error {
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
func (this *ContentAttributes) Validate() error {
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
func (this *ContentRequest) Validate() error {
	return nil
}
func (this *ContentIdRequest) Validate() error {
	return nil
}

var _regex_NewContentAttributes_Name = regexp.MustCompile("[a-z,A-z,0-9]+")

func (this *NewContentAttributes) Validate() error {
	if !_regex_NewContentAttributes_Name.MatchString(this.Name) {
		return github_com_mwitkow_go_proto_validators.FieldError("Name", fmt.Errorf(`value '%v' must be a string conforming to regex "[a-z,A-z,0-9]+"`, this.Name))
	}
	if !(this.CreatedBy > 0) {
		return github_com_mwitkow_go_proto_validators.FieldError("CreatedBy", fmt.Errorf(`value '%v' must be greater than '0'`, this.CreatedBy))
	}
	if this.Content == "" {
		return github_com_mwitkow_go_proto_validators.FieldError("Content", fmt.Errorf(`value '%v' must not be an empty string`, this.Content))
	}
	if this.Namespace == "" {
		return github_com_mwitkow_go_proto_validators.FieldError("Namespace", fmt.Errorf(`value '%v' must not be an empty string`, this.Namespace))
	}
	return nil
}
func (this *StoreContentRequest) Validate() error {
	if this.Data != nil {
		if err := github_com_mwitkow_go_proto_validators.CallValidatorIfExists(this.Data); err != nil {
			return github_com_mwitkow_go_proto_validators.FieldError("Data", err)
		}
	}
	return nil
}
func (this *StoreContentRequest_Data) Validate() error {
	if this.Attributes != nil {
		if err := github_com_mwitkow_go_proto_validators.CallValidatorIfExists(this.Attributes); err != nil {
			return github_com_mwitkow_go_proto_validators.FieldError("Attributes", err)
		}
	}
	return nil
}
func (this *ExistingContentAttributes) Validate() error {
	if !(this.UpdatedBy > 0) {
		return github_com_mwitkow_go_proto_validators.FieldError("UpdatedBy", fmt.Errorf(`value '%v' must be greater than '0'`, this.UpdatedBy))
	}
	if this.Content == "" {
		return github_com_mwitkow_go_proto_validators.FieldError("Content", fmt.Errorf(`value '%v' must not be an empty string`, this.Content))
	}
	return nil
}
func (this *UpdateContentRequest) Validate() error {
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
func (this *UpdateContentRequest_Data) Validate() error {
	if this.Attributes != nil {
		if err := github_com_mwitkow_go_proto_validators.CallValidatorIfExists(this.Attributes); err != nil {
			return github_com_mwitkow_go_proto_validators.FieldError("Attributes", err)
		}
	}
	return nil
}