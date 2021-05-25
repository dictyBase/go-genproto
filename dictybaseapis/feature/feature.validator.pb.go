// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: feature.proto

package feature

import (
	fmt "fmt"
	math "math"
	proto "github.com/golang/protobuf/proto"
	_ "github.com/mwitkow/go-proto-validators"
	_ "google.golang.org/genproto/protobuf/field_mask"
	_ "github.com/golang/protobuf/ptypes/timestamp"
	_ "github.com/golang/protobuf/ptypes/empty"
	github_com_mwitkow_go_proto_validators "github.com/mwitkow/go-proto-validators"
)

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

func (this *FeatureId) Validate() error {
	return nil
}
func (this *FeatureRelationFilter) Validate() error {
	return nil
}
func (this *ReferenceFeatureFilter) Validate() error {
	return nil
}
func (this *LocatedFeatureFilter) Validate() error {
	return nil
}
func (this *NewFeature) Validate() error {
	if this.Data != nil {
		if err := github_com_mwitkow_go_proto_validators.CallValidatorIfExists(this.Data); err != nil {
			return github_com_mwitkow_go_proto_validators.FieldError("Data", err)
		}
	}
	return nil
}
func (this *NewFeature_Data) Validate() error {
	if this.Type == "" {
		return github_com_mwitkow_go_proto_validators.FieldError("Type", fmt.Errorf(`value '%v' must not be an empty string`, this.Type))
	}
	if this.Attributes != nil {
		if err := github_com_mwitkow_go_proto_validators.CallValidatorIfExists(this.Attributes); err != nil {
			return github_com_mwitkow_go_proto_validators.FieldError("Attributes", err)
		}
	}
	return nil
}
func (this *NewFeatureAttributes) Validate() error {
	if !(this.CreatedBy > 0) {
		return github_com_mwitkow_go_proto_validators.FieldError("CreatedBy", fmt.Errorf(`value '%v' must be greater than '0'`, this.CreatedBy))
	}
	if !(this.UpdatedBy > 0) {
		return github_com_mwitkow_go_proto_validators.FieldError("UpdatedBy", fmt.Errorf(`value '%v' must be greater than '0'`, this.UpdatedBy))
	}
	for _, item := range this.Dbxrefs {
		if item != nil {
			if err := github_com_mwitkow_go_proto_validators.CallValidatorIfExists(item); err != nil {
				return github_com_mwitkow_go_proto_validators.FieldError("Dbxrefs", err)
			}
		}
	}
	if this.Organism != nil {
		if err := github_com_mwitkow_go_proto_validators.CallValidatorIfExists(this.Organism); err != nil {
			return github_com_mwitkow_go_proto_validators.FieldError("Organism", err)
		}
	}
	for _, item := range this.Parents {
		if item != nil {
			if err := github_com_mwitkow_go_proto_validators.CallValidatorIfExists(item); err != nil {
				return github_com_mwitkow_go_proto_validators.FieldError("Parents", err)
			}
		}
	}
	for _, item := range this.Children {
		if item != nil {
			if err := github_com_mwitkow_go_proto_validators.CallValidatorIfExists(item); err != nil {
				return github_com_mwitkow_go_proto_validators.FieldError("Children", err)
			}
		}
	}
	for _, item := range this.Location {
		if item != nil {
			if err := github_com_mwitkow_go_proto_validators.CallValidatorIfExists(item); err != nil {
				return github_com_mwitkow_go_proto_validators.FieldError("Location", err)
			}
		}
	}
	for _, item := range this.Publications {
		if item != nil {
			if err := github_com_mwitkow_go_proto_validators.CallValidatorIfExists(item); err != nil {
				return github_com_mwitkow_go_proto_validators.FieldError("Publications", err)
			}
		}
	}
	return nil
}
func (this *PaginatedFeatureCollection) Validate() error {
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
func (this *FeatureCollection) Validate() error {
	for _, item := range this.Data {
		if item != nil {
			if err := github_com_mwitkow_go_proto_validators.CallValidatorIfExists(item); err != nil {
				return github_com_mwitkow_go_proto_validators.FieldError("Data", err)
			}
		}
	}
	return nil
}
func (this *Feature) Validate() error {
	if this.Data != nil {
		if err := github_com_mwitkow_go_proto_validators.CallValidatorIfExists(this.Data); err != nil {
			return github_com_mwitkow_go_proto_validators.FieldError("Data", err)
		}
	}
	return nil
}
func (this *FeatureData) Validate() error {
	if this.Attributes != nil {
		if err := github_com_mwitkow_go_proto_validators.CallValidatorIfExists(this.Attributes); err != nil {
			return github_com_mwitkow_go_proto_validators.FieldError("Attributes", err)
		}
	}
	return nil
}
func (this *FeatureAttributes) Validate() error {
	if !(this.CreatedBy > 0) {
		return github_com_mwitkow_go_proto_validators.FieldError("CreatedBy", fmt.Errorf(`value '%v' must be greater than '0'`, this.CreatedBy))
	}
	if !(this.UpdatedBy > 0) {
		return github_com_mwitkow_go_proto_validators.FieldError("UpdatedBy", fmt.Errorf(`value '%v' must be greater than '0'`, this.UpdatedBy))
	}
	if nil == this.CreatedAt {
		return github_com_mwitkow_go_proto_validators.FieldError("CreatedAt", fmt.Errorf("message must exist"))
	}
	if this.CreatedAt != nil {
		if err := github_com_mwitkow_go_proto_validators.CallValidatorIfExists(this.CreatedAt); err != nil {
			return github_com_mwitkow_go_proto_validators.FieldError("CreatedAt", err)
		}
	}
	if nil == this.UpdatedAt {
		return github_com_mwitkow_go_proto_validators.FieldError("UpdatedAt", fmt.Errorf("message must exist"))
	}
	if this.UpdatedAt != nil {
		if err := github_com_mwitkow_go_proto_validators.CallValidatorIfExists(this.UpdatedAt); err != nil {
			return github_com_mwitkow_go_proto_validators.FieldError("UpdatedAt", err)
		}
	}
	if this.Organism != nil {
		if err := github_com_mwitkow_go_proto_validators.CallValidatorIfExists(this.Organism); err != nil {
			return github_com_mwitkow_go_proto_validators.FieldError("Organism", err)
		}
	}
	for _, item := range this.Parents {
		if item != nil {
			if err := github_com_mwitkow_go_proto_validators.CallValidatorIfExists(item); err != nil {
				return github_com_mwitkow_go_proto_validators.FieldError("Parents", err)
			}
		}
	}
	for _, item := range this.Children {
		if item != nil {
			if err := github_com_mwitkow_go_proto_validators.CallValidatorIfExists(item); err != nil {
				return github_com_mwitkow_go_proto_validators.FieldError("Children", err)
			}
		}
	}
	for _, item := range this.Location {
		if item != nil {
			if err := github_com_mwitkow_go_proto_validators.CallValidatorIfExists(item); err != nil {
				return github_com_mwitkow_go_proto_validators.FieldError("Location", err)
			}
		}
	}
	for _, item := range this.Publications {
		if item != nil {
			if err := github_com_mwitkow_go_proto_validators.CallValidatorIfExists(item); err != nil {
				return github_com_mwitkow_go_proto_validators.FieldError("Publications", err)
			}
		}
	}
	if !(this.Version > 0) {
		return github_com_mwitkow_go_proto_validators.FieldError("Version", fmt.Errorf(`value '%v' must be greater than '0'`, this.Version))
	}
	for _, item := range this.Dbxrefs {
		if item != nil {
			if err := github_com_mwitkow_go_proto_validators.CallValidatorIfExists(item); err != nil {
				return github_com_mwitkow_go_proto_validators.FieldError("Dbxrefs", err)
			}
		}
	}
	return nil
}
func (this *FeatureUpdate) Validate() error {
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
func (this *FeatureUpdate_Data) Validate() error {
	if this.Type == "" {
		return github_com_mwitkow_go_proto_validators.FieldError("Type", fmt.Errorf(`value '%v' must not be an empty string`, this.Type))
	}
	if this.Attributes != nil {
		if err := github_com_mwitkow_go_proto_validators.CallValidatorIfExists(this.Attributes); err != nil {
			return github_com_mwitkow_go_proto_validators.FieldError("Attributes", err)
		}
	}
	if this.Id == "" {
		return github_com_mwitkow_go_proto_validators.FieldError("Id", fmt.Errorf(`value '%v' must not be an empty string`, this.Id))
	}
	return nil
}
func (this *FeatureUpdateAttributes) Validate() error {
	if !(this.UpdatedBy > 0) {
		return github_com_mwitkow_go_proto_validators.FieldError("UpdatedBy", fmt.Errorf(`value '%v' must be greater than '0'`, this.UpdatedBy))
	}
	for _, item := range this.Dbxrefs {
		if item != nil {
			if err := github_com_mwitkow_go_proto_validators.CallValidatorIfExists(item); err != nil {
				return github_com_mwitkow_go_proto_validators.FieldError("Dbxrefs", err)
			}
		}
	}
	if this.Organism != nil {
		if err := github_com_mwitkow_go_proto_validators.CallValidatorIfExists(this.Organism); err != nil {
			return github_com_mwitkow_go_proto_validators.FieldError("Organism", err)
		}
	}
	for _, item := range this.Parents {
		if item != nil {
			if err := github_com_mwitkow_go_proto_validators.CallValidatorIfExists(item); err != nil {
				return github_com_mwitkow_go_proto_validators.FieldError("Parents", err)
			}
		}
	}
	for _, item := range this.Children {
		if item != nil {
			if err := github_com_mwitkow_go_proto_validators.CallValidatorIfExists(item); err != nil {
				return github_com_mwitkow_go_proto_validators.FieldError("Children", err)
			}
		}
	}
	for _, item := range this.Location {
		if item != nil {
			if err := github_com_mwitkow_go_proto_validators.CallValidatorIfExists(item); err != nil {
				return github_com_mwitkow_go_proto_validators.FieldError("Location", err)
			}
		}
	}
	for _, item := range this.Publications {
		if item != nil {
			if err := github_com_mwitkow_go_proto_validators.CallValidatorIfExists(item); err != nil {
				return github_com_mwitkow_go_proto_validators.FieldError("Publications", err)
			}
		}
	}
	return nil
}
func (this *FeatureSynonym) Validate() error {
	return nil
}
func (this *FeatureHistory) Validate() error {
	return nil
}
func (this *Publication) Validate() error {
	if this.Source == "" {
		return github_com_mwitkow_go_proto_validators.FieldError("Source", fmt.Errorf(`value '%v' must not be an empty string`, this.Source))
	}
	return nil
}
func (this *FeaturePublication) Validate() error {
	if this.Publication != nil {
		if err := github_com_mwitkow_go_proto_validators.CallValidatorIfExists(this.Publication); err != nil {
			return github_com_mwitkow_go_proto_validators.FieldError("Publication", err)
		}
	}
	return nil
}
func (this *FeatureConnection) Validate() error {
	if this.Relationship != nil {
		if err := github_com_mwitkow_go_proto_validators.CallValidatorIfExists(this.Relationship); err != nil {
			return github_com_mwitkow_go_proto_validators.FieldError("Relationship", err)
		}
	}
	return nil
}
func (this *FeatureRelationship) Validate() error {
	return nil
}
func (this *FeatureDbxref) Validate() error {
	if this.Dbxref != nil {
		if err := github_com_mwitkow_go_proto_validators.CallValidatorIfExists(this.Dbxref); err != nil {
			return github_com_mwitkow_go_proto_validators.FieldError("Dbxref", err)
		}
	}
	return nil
}
func (this *Dbxref) Validate() error {
	if this.DbxrefId == "" {
		return github_com_mwitkow_go_proto_validators.FieldError("DbxrefId", fmt.Errorf(`value '%v' must not be an empty string`, this.DbxrefId))
	}
	if this.Database == "" {
		return github_com_mwitkow_go_proto_validators.FieldError("Database", fmt.Errorf(`value '%v' must not be an empty string`, this.Database))
	}
	return nil
}
func (this *NewFeatureLocation) Validate() error {
	if this.Location != nil {
		if err := github_com_mwitkow_go_proto_validators.CallValidatorIfExists(this.Location); err != nil {
			return github_com_mwitkow_go_proto_validators.FieldError("Location", err)
		}
	}
	return nil
}
func (this *FeatureLocation) Validate() error {
	return nil
}
func (this *Organism) Validate() error {
	return nil
}
func (this *ListParameters) Validate() error {
	return nil
}
func (this *Meta) Validate() error {
	return nil
}
