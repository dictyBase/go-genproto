// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: annotation.proto

package annotation // import "github.com/dictyBase/go-genproto/dictybaseapis/annotation"

import fmt "fmt"
import github_com_mwitkow_go_proto_validators "github.com/mwitkow/go-proto-validators"
import proto "github.com/golang/protobuf/proto"
import math "math"
import _ "github.com/golang/protobuf/ptypes/empty"
import _ "github.com/golang/protobuf/ptypes/timestamp"
import _ "github.com/mwitkow/go-proto-validators"
import _ "google.golang.org/genproto/googleapis/api/annotations"

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

func (this *AnnotationId) Validate() error {
	if this.Id == "" {
		return github_com_mwitkow_go_proto_validators.FieldError("Id", fmt.Errorf(`value '%v' must not be an empty string`, this.Id))
	}
	return nil
}
func (this *EntryAnnotationRequest) Validate() error {
	if this.Tag == "" {
		return github_com_mwitkow_go_proto_validators.FieldError("Tag", fmt.Errorf(`value '%v' must not be an empty string`, this.Tag))
	}
	if this.EntryId == "" {
		return github_com_mwitkow_go_proto_validators.FieldError("EntryId", fmt.Errorf(`value '%v' must not be an empty string`, this.EntryId))
	}
	if this.Ontology == "" {
		return github_com_mwitkow_go_proto_validators.FieldError("Ontology", fmt.Errorf(`value '%v' must not be an empty string`, this.Ontology))
	}
	return nil
}
func (this *TaggedAnnotationCollection) Validate() error {
	for _, item := range this.Data {
		if item != nil {
			if err := github_com_mwitkow_go_proto_validators.CallValidatorIfExists(item); err != nil {
				return github_com_mwitkow_go_proto_validators.FieldError("Data", err)
			}
		}
	}
	if nil == this.Meta {
		return github_com_mwitkow_go_proto_validators.FieldError("Meta", fmt.Errorf("message must exist"))
	}
	if this.Meta != nil {
		if err := github_com_mwitkow_go_proto_validators.CallValidatorIfExists(this.Meta); err != nil {
			return github_com_mwitkow_go_proto_validators.FieldError("Meta", err)
		}
	}
	return nil
}
func (this *TaggedAnnotationCollection_Data) Validate() error {
	if this.Attributes != nil {
		if err := github_com_mwitkow_go_proto_validators.CallValidatorIfExists(this.Attributes); err != nil {
			return github_com_mwitkow_go_proto_validators.FieldError("Attributes", err)
		}
	}
	return nil
}
func (this *TaggedAnnotation) Validate() error {
	if nil == this.Data {
		return github_com_mwitkow_go_proto_validators.FieldError("Data", fmt.Errorf("message must exist"))
	}
	if this.Data != nil {
		if err := github_com_mwitkow_go_proto_validators.CallValidatorIfExists(this.Data); err != nil {
			return github_com_mwitkow_go_proto_validators.FieldError("Data", err)
		}
	}
	return nil
}
func (this *TaggedAnnotation_Data) Validate() error {
	if this.Attributes != nil {
		if err := github_com_mwitkow_go_proto_validators.CallValidatorIfExists(this.Attributes); err != nil {
			return github_com_mwitkow_go_proto_validators.FieldError("Attributes", err)
		}
	}
	return nil
}
func (this *TaggedAnnotationAttributes) Validate() error {
	if this.Value == "" {
		return github_com_mwitkow_go_proto_validators.FieldError("Value", fmt.Errorf(`value '%v' must not be an empty string`, this.Value))
	}
	if this.EditableValue == "" {
		return github_com_mwitkow_go_proto_validators.FieldError("EditableValue", fmt.Errorf(`value '%v' must not be an empty string`, this.EditableValue))
	}
	if this.CreatedBy == "" {
		return github_com_mwitkow_go_proto_validators.FieldError("CreatedBy", fmt.Errorf(`value '%v' must not be an empty string`, this.CreatedBy))
	}
	if nil == this.CreatedAt {
		return github_com_mwitkow_go_proto_validators.FieldError("CreatedAt", fmt.Errorf("message must exist"))
	}
	if this.CreatedAt != nil {
		if err := github_com_mwitkow_go_proto_validators.CallValidatorIfExists(this.CreatedAt); err != nil {
			return github_com_mwitkow_go_proto_validators.FieldError("CreatedAt", err)
		}
	}
	if this.Tag == "" {
		return github_com_mwitkow_go_proto_validators.FieldError("Tag", fmt.Errorf(`value '%v' must not be an empty string`, this.Tag))
	}
	if !(this.Version > 0) {
		return github_com_mwitkow_go_proto_validators.FieldError("Version", fmt.Errorf(`value '%v' must be greater than '0'`, this.Version))
	}
	if this.EntryId == "" {
		return github_com_mwitkow_go_proto_validators.FieldError("EntryId", fmt.Errorf(`value '%v' must not be an empty string`, this.EntryId))
	}
	if this.Ontology == "" {
		return github_com_mwitkow_go_proto_validators.FieldError("Ontology", fmt.Errorf(`value '%v' must not be an empty string`, this.Ontology))
	}
	return nil
}
func (this *NewTaggedAnnotation) Validate() error {
	if nil == this.Data {
		return github_com_mwitkow_go_proto_validators.FieldError("Data", fmt.Errorf("message must exist"))
	}
	if this.Data != nil {
		if err := github_com_mwitkow_go_proto_validators.CallValidatorIfExists(this.Data); err != nil {
			return github_com_mwitkow_go_proto_validators.FieldError("Data", err)
		}
	}
	return nil
}
func (this *NewTaggedAnnotation_Data) Validate() error {
	if this.Attributes != nil {
		if err := github_com_mwitkow_go_proto_validators.CallValidatorIfExists(this.Attributes); err != nil {
			return github_com_mwitkow_go_proto_validators.FieldError("Attributes", err)
		}
	}
	return nil
}
func (this *NewTaggedAnnotationAttributes) Validate() error {
	if this.Value == "" {
		return github_com_mwitkow_go_proto_validators.FieldError("Value", fmt.Errorf(`value '%v' must not be an empty string`, this.Value))
	}
	if this.EditableValue == "" {
		return github_com_mwitkow_go_proto_validators.FieldError("EditableValue", fmt.Errorf(`value '%v' must not be an empty string`, this.EditableValue))
	}
	if this.CreatedBy == "" {
		return github_com_mwitkow_go_proto_validators.FieldError("CreatedBy", fmt.Errorf(`value '%v' must not be an empty string`, this.CreatedBy))
	}
	if this.Tag == "" {
		return github_com_mwitkow_go_proto_validators.FieldError("Tag", fmt.Errorf(`value '%v' must not be an empty string`, this.Tag))
	}
	if this.EntryId == "" {
		return github_com_mwitkow_go_proto_validators.FieldError("EntryId", fmt.Errorf(`value '%v' must not be an empty string`, this.EntryId))
	}
	if this.Ontology == "" {
		return github_com_mwitkow_go_proto_validators.FieldError("Ontology", fmt.Errorf(`value '%v' must not be an empty string`, this.Ontology))
	}
	return nil
}
func (this *TaggedAnnotationUpdate) Validate() error {
	if nil == this.Data {
		return github_com_mwitkow_go_proto_validators.FieldError("Data", fmt.Errorf("message must exist"))
	}
	if this.Data != nil {
		if err := github_com_mwitkow_go_proto_validators.CallValidatorIfExists(this.Data); err != nil {
			return github_com_mwitkow_go_proto_validators.FieldError("Data", err)
		}
	}
	return nil
}
func (this *TaggedAnnotationUpdate_Data) Validate() error {
	if this.Id == "" {
		return github_com_mwitkow_go_proto_validators.FieldError("Id", fmt.Errorf(`value '%v' must not be an empty string`, this.Id))
	}
	if this.Attributes != nil {
		if err := github_com_mwitkow_go_proto_validators.CallValidatorIfExists(this.Attributes); err != nil {
			return github_com_mwitkow_go_proto_validators.FieldError("Attributes", err)
		}
	}
	return nil
}
func (this *TaggedAnnotationUpdateAttributes) Validate() error {
	if this.Value == "" {
		return github_com_mwitkow_go_proto_validators.FieldError("Value", fmt.Errorf(`value '%v' must not be an empty string`, this.Value))
	}
	if this.EditableValue == "" {
		return github_com_mwitkow_go_proto_validators.FieldError("EditableValue", fmt.Errorf(`value '%v' must not be an empty string`, this.EditableValue))
	}
	if this.CreatedBy == "" {
		return github_com_mwitkow_go_proto_validators.FieldError("CreatedBy", fmt.Errorf(`value '%v' must not be an empty string`, this.CreatedBy))
	}
	return nil
}
func (this *ListParameters) Validate() error {
	return nil
}
func (this *Meta) Validate() error {
	return nil
}
