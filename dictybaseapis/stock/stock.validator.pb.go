// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: dictybase/stock/stock.proto

package stock

import (
	fmt "fmt"
	math "math"
	proto "github.com/golang/protobuf/proto"
	_ "google.golang.org/protobuf/types/known/emptypb"
	_ "google.golang.org/protobuf/types/known/timestamppb"
	_ "github.com/dictyBase/go-genproto/dictybaseapis/api/upload"
	_ "github.com/mwitkow/go-proto-validators"
	regexp "regexp"
	github_com_mwitkow_go_proto_validators "github.com/mwitkow/go-proto-validators"
)

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

func (this *StockId) Validate() error {
	if this.Id == "" {
		return github_com_mwitkow_go_proto_validators.FieldError("Id", fmt.Errorf(`value '%v' must not be an empty string`, this.Id))
	}
	return nil
}

var _regex_StockIdList_Id = regexp.MustCompile(`^DB(P|S)[0-9]{5,}$`)

func (this *StockIdList) Validate() error {
	for _, item := range this.Id {
		if !_regex_StockIdList_Id.MatchString(item) {
			return github_com_mwitkow_go_proto_validators.FieldError("Id", fmt.Errorf(`value '%v' must be a string conforming to regex "^DB(P|S)[0-9]{5,}$"`, item))
		}
	}
	return nil
}
func (this *Strain) Validate() error {
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
func (this *Strain_Data) Validate() error {
	if this.Attributes != nil {
		if err := github_com_mwitkow_go_proto_validators.CallValidatorIfExists(this.Attributes); err != nil {
			return github_com_mwitkow_go_proto_validators.FieldError("Attributes", err)
		}
	}
	return nil
}
func (this *Plasmid) Validate() error {
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
func (this *Plasmid_Data) Validate() error {
	if this.Attributes != nil {
		if err := github_com_mwitkow_go_proto_validators.CallValidatorIfExists(this.Attributes); err != nil {
			return github_com_mwitkow_go_proto_validators.FieldError("Attributes", err)
		}
	}
	return nil
}
func (this *StrainAttributes) Validate() error {
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
	if this.CreatedBy == "" {
		return github_com_mwitkow_go_proto_validators.FieldError("CreatedBy", fmt.Errorf(`value '%v' must not be an empty string`, this.CreatedBy))
	}
	if this.UpdatedBy == "" {
		return github_com_mwitkow_go_proto_validators.FieldError("UpdatedBy", fmt.Errorf(`value '%v' must not be an empty string`, this.UpdatedBy))
	}
	if this.Label == "" {
		return github_com_mwitkow_go_proto_validators.FieldError("Label", fmt.Errorf(`value '%v' must not be an empty string`, this.Label))
	}
	if this.Species == "" {
		return github_com_mwitkow_go_proto_validators.FieldError("Species", fmt.Errorf(`value '%v' must not be an empty string`, this.Species))
	}
	return nil
}
func (this *PlasmidAttributes) Validate() error {
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
	if this.CreatedBy == "" {
		return github_com_mwitkow_go_proto_validators.FieldError("CreatedBy", fmt.Errorf(`value '%v' must not be an empty string`, this.CreatedBy))
	}
	if this.UpdatedBy == "" {
		return github_com_mwitkow_go_proto_validators.FieldError("UpdatedBy", fmt.Errorf(`value '%v' must not be an empty string`, this.UpdatedBy))
	}
	if this.Name == "" {
		return github_com_mwitkow_go_proto_validators.FieldError("Name", fmt.Errorf(`value '%v' must not be an empty string`, this.Name))
	}
	return nil
}
func (this *NewStrain) Validate() error {
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
func (this *NewStrain_Data) Validate() error {
	if this.Attributes != nil {
		if err := github_com_mwitkow_go_proto_validators.CallValidatorIfExists(this.Attributes); err != nil {
			return github_com_mwitkow_go_proto_validators.FieldError("Attributes", err)
		}
	}
	return nil
}
func (this *NewPlasmid) Validate() error {
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
func (this *NewPlasmid_Data) Validate() error {
	if this.Attributes != nil {
		if err := github_com_mwitkow_go_proto_validators.CallValidatorIfExists(this.Attributes); err != nil {
			return github_com_mwitkow_go_proto_validators.FieldError("Attributes", err)
		}
	}
	return nil
}
func (this *NewStrainAttributes) Validate() error {
	if this.CreatedBy == "" {
		return github_com_mwitkow_go_proto_validators.FieldError("CreatedBy", fmt.Errorf(`value '%v' must not be an empty string`, this.CreatedBy))
	}
	if this.UpdatedBy == "" {
		return github_com_mwitkow_go_proto_validators.FieldError("UpdatedBy", fmt.Errorf(`value '%v' must not be an empty string`, this.UpdatedBy))
	}
	if this.Depositor == "" {
		return github_com_mwitkow_go_proto_validators.FieldError("Depositor", fmt.Errorf(`value '%v' must not be an empty string`, this.Depositor))
	}
	if this.Label == "" {
		return github_com_mwitkow_go_proto_validators.FieldError("Label", fmt.Errorf(`value '%v' must not be an empty string`, this.Label))
	}
	if this.Species == "" {
		return github_com_mwitkow_go_proto_validators.FieldError("Species", fmt.Errorf(`value '%v' must not be an empty string`, this.Species))
	}
	return nil
}
func (this *NewPlasmidAttributes) Validate() error {
	if this.CreatedBy == "" {
		return github_com_mwitkow_go_proto_validators.FieldError("CreatedBy", fmt.Errorf(`value '%v' must not be an empty string`, this.CreatedBy))
	}
	if this.UpdatedBy == "" {
		return github_com_mwitkow_go_proto_validators.FieldError("UpdatedBy", fmt.Errorf(`value '%v' must not be an empty string`, this.UpdatedBy))
	}
	if this.Depositor == "" {
		return github_com_mwitkow_go_proto_validators.FieldError("Depositor", fmt.Errorf(`value '%v' must not be an empty string`, this.Depositor))
	}
	if this.Name == "" {
		return github_com_mwitkow_go_proto_validators.FieldError("Name", fmt.Errorf(`value '%v' must not be an empty string`, this.Name))
	}
	return nil
}
func (this *ExistingStrain) Validate() error {
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
func (this *ExistingStrain_Data) Validate() error {
	if this.Attributes != nil {
		if err := github_com_mwitkow_go_proto_validators.CallValidatorIfExists(this.Attributes); err != nil {
			return github_com_mwitkow_go_proto_validators.FieldError("Attributes", err)
		}
	}
	return nil
}
func (this *ExistingPlasmid) Validate() error {
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
func (this *ExistingPlasmid_Data) Validate() error {
	if this.Attributes != nil {
		if err := github_com_mwitkow_go_proto_validators.CallValidatorIfExists(this.Attributes); err != nil {
			return github_com_mwitkow_go_proto_validators.FieldError("Attributes", err)
		}
	}
	return nil
}
func (this *ExistingStrainAttributes) Validate() error {
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
	if this.CreatedBy == "" {
		return github_com_mwitkow_go_proto_validators.FieldError("CreatedBy", fmt.Errorf(`value '%v' must not be an empty string`, this.CreatedBy))
	}
	if this.UpdatedBy == "" {
		return github_com_mwitkow_go_proto_validators.FieldError("UpdatedBy", fmt.Errorf(`value '%v' must not be an empty string`, this.UpdatedBy))
	}
	if this.Label == "" {
		return github_com_mwitkow_go_proto_validators.FieldError("Label", fmt.Errorf(`value '%v' must not be an empty string`, this.Label))
	}
	if this.Species == "" {
		return github_com_mwitkow_go_proto_validators.FieldError("Species", fmt.Errorf(`value '%v' must not be an empty string`, this.Species))
	}
	return nil
}
func (this *ExistingPlasmidAttributes) Validate() error {
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
	if this.CreatedBy == "" {
		return github_com_mwitkow_go_proto_validators.FieldError("CreatedBy", fmt.Errorf(`value '%v' must not be an empty string`, this.CreatedBy))
	}
	if this.Name == "" {
		return github_com_mwitkow_go_proto_validators.FieldError("Name", fmt.Errorf(`value '%v' must not be an empty string`, this.Name))
	}
	return nil
}
func (this *StrainUpdate) Validate() error {
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
func (this *StrainUpdate_Data) Validate() error {
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
func (this *PlasmidUpdate) Validate() error {
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
func (this *PlasmidUpdate_Data) Validate() error {
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
func (this *StrainUpdateAttributes) Validate() error {
	if this.UpdatedBy == "" {
		return github_com_mwitkow_go_proto_validators.FieldError("UpdatedBy", fmt.Errorf(`value '%v' must not be an empty string`, this.UpdatedBy))
	}
	return nil
}
func (this *PlasmidUpdateAttributes) Validate() error {
	if this.UpdatedBy == "" {
		return github_com_mwitkow_go_proto_validators.FieldError("UpdatedBy", fmt.Errorf(`value '%v' must not be an empty string`, this.UpdatedBy))
	}
	return nil
}
func (this *StrainList) Validate() error {
	for _, item := range this.Data {
		if item != nil {
			if err := github_com_mwitkow_go_proto_validators.CallValidatorIfExists(item); err != nil {
				return github_com_mwitkow_go_proto_validators.FieldError("Data", err)
			}
		}
	}
	return nil
}
func (this *StrainList_Data) Validate() error {
	if this.Attributes != nil {
		if err := github_com_mwitkow_go_proto_validators.CallValidatorIfExists(this.Attributes); err != nil {
			return github_com_mwitkow_go_proto_validators.FieldError("Attributes", err)
		}
	}
	return nil
}
func (this *StrainCollection) Validate() error {
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
func (this *StrainCollection_Data) Validate() error {
	if this.Attributes != nil {
		if err := github_com_mwitkow_go_proto_validators.CallValidatorIfExists(this.Attributes); err != nil {
			return github_com_mwitkow_go_proto_validators.FieldError("Attributes", err)
		}
	}
	return nil
}
func (this *PlasmidCollection) Validate() error {
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
func (this *PlasmidCollection_Data) Validate() error {
	if this.Attributes != nil {
		if err := github_com_mwitkow_go_proto_validators.CallValidatorIfExists(this.Attributes); err != nil {
			return github_com_mwitkow_go_proto_validators.FieldError("Attributes", err)
		}
	}
	return nil
}
func (this *StockParameters) Validate() error {
	return nil
}
func (this *Meta) Validate() error {
	return nil
}
