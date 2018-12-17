// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: stock.proto

package stock // import "github.com/dictyBase/go-genproto/dictybaseapis/stock"

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

func (this *StockId) Validate() error {
	if this.Id == "" {
		return github_com_mwitkow_go_proto_validators.FieldError("Id", fmt.Errorf(`value '%v' must not be an empty string`, this.Id))
	}
	return nil
}
func (this *Stock) Validate() error {
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
func (this *Stock_Data) Validate() error {
	if this.Attributes != nil {
		if err := github_com_mwitkow_go_proto_validators.CallValidatorIfExists(this.Attributes); err != nil {
			return github_com_mwitkow_go_proto_validators.FieldError("Attributes", err)
		}
	}
	return nil
}
func (this *StockAttributes) Validate() error {
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
	if this.Summary == "" {
		return github_com_mwitkow_go_proto_validators.FieldError("Summary", fmt.Errorf(`value '%v' must not be an empty string`, this.Summary))
	}
	if this.EditableSummary == "" {
		return github_com_mwitkow_go_proto_validators.FieldError("EditableSummary", fmt.Errorf(`value '%v' must not be an empty string`, this.EditableSummary))
	}
	if this.StrainProperties != nil {
		if err := github_com_mwitkow_go_proto_validators.CallValidatorIfExists(this.StrainProperties); err != nil {
			return github_com_mwitkow_go_proto_validators.FieldError("StrainProperties", err)
		}
	}
	if this.PlasmidProperties != nil {
		if err := github_com_mwitkow_go_proto_validators.CallValidatorIfExists(this.PlasmidProperties); err != nil {
			return github_com_mwitkow_go_proto_validators.FieldError("PlasmidProperties", err)
		}
	}
	return nil
}
func (this *StrainProperties) Validate() error {
	if this.SystematicName == "" {
		return github_com_mwitkow_go_proto_validators.FieldError("SystematicName", fmt.Errorf(`value '%v' must not be an empty string`, this.SystematicName))
	}
	if this.Descriptor_ == "" {
		return github_com_mwitkow_go_proto_validators.FieldError("Descriptor_", fmt.Errorf(`value '%v' must not be an empty string`, this.Descriptor_))
	}
	return nil
}
func (this *PlasmidProperties) Validate() error {
	return nil
}
func (this *NewStock) Validate() error {
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
func (this *NewStock_Data) Validate() error {
	if this.Attributes != nil {
		if err := github_com_mwitkow_go_proto_validators.CallValidatorIfExists(this.Attributes); err != nil {
			return github_com_mwitkow_go_proto_validators.FieldError("Attributes", err)
		}
	}
	return nil
}
func (this *NewStockAttributes) Validate() error {
	if this.CreatedBy == "" {
		return github_com_mwitkow_go_proto_validators.FieldError("CreatedBy", fmt.Errorf(`value '%v' must not be an empty string`, this.CreatedBy))
	}
	if this.UpdatedBy == "" {
		return github_com_mwitkow_go_proto_validators.FieldError("UpdatedBy", fmt.Errorf(`value '%v' must not be an empty string`, this.UpdatedBy))
	}
	if this.Summary == "" {
		return github_com_mwitkow_go_proto_validators.FieldError("Summary", fmt.Errorf(`value '%v' must not be an empty string`, this.Summary))
	}
	if this.EditableSummary == "" {
		return github_com_mwitkow_go_proto_validators.FieldError("EditableSummary", fmt.Errorf(`value '%v' must not be an empty string`, this.EditableSummary))
	}
	if this.Depositor == "" {
		return github_com_mwitkow_go_proto_validators.FieldError("Depositor", fmt.Errorf(`value '%v' must not be an empty string`, this.Depositor))
	}
	if this.StrainProperties != nil {
		if err := github_com_mwitkow_go_proto_validators.CallValidatorIfExists(this.StrainProperties); err != nil {
			return github_com_mwitkow_go_proto_validators.FieldError("StrainProperties", err)
		}
	}
	if this.PlasmidProperties != nil {
		if err := github_com_mwitkow_go_proto_validators.CallValidatorIfExists(this.PlasmidProperties); err != nil {
			return github_com_mwitkow_go_proto_validators.FieldError("PlasmidProperties", err)
		}
	}
	return nil
}
func (this *StockUpdate) Validate() error {
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
func (this *StockUpdate_Data) Validate() error {
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
func (this *StockUpdateAttributes) Validate() error {
	if this.UpdatedBy == "" {
		return github_com_mwitkow_go_proto_validators.FieldError("UpdatedBy", fmt.Errorf(`value '%v' must not be an empty string`, this.UpdatedBy))
	}
	if this.StrainProperties != nil {
		if err := github_com_mwitkow_go_proto_validators.CallValidatorIfExists(this.StrainProperties); err != nil {
			return github_com_mwitkow_go_proto_validators.FieldError("StrainProperties", err)
		}
	}
	if this.PlasmidProperties != nil {
		if err := github_com_mwitkow_go_proto_validators.CallValidatorIfExists(this.PlasmidProperties); err != nil {
			return github_com_mwitkow_go_proto_validators.FieldError("PlasmidProperties", err)
		}
	}
	return nil
}
func (this *StrainUpdateProperties) Validate() error {
	return nil
}
func (this *StockCollection) Validate() error {
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
func (this *StockCollection_Data) Validate() error {
	if this.Attributes != nil {
		if err := github_com_mwitkow_go_proto_validators.CallValidatorIfExists(this.Attributes); err != nil {
			return github_com_mwitkow_go_proto_validators.FieldError("Attributes", err)
		}
	}
	return nil
}
func (this *ListParameters) Validate() error {
	return nil
}
func (this *StrainListParameters) Validate() error {
	return nil
}
func (this *PlasmidListParameters) Validate() error {
	return nil
}
func (this *Meta) Validate() error {
	return nil
}
