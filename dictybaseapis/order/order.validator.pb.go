// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: order.proto

package order // import "github.com/dictyBase/go-genproto/dictybaseapis/order"

import fmt "fmt"
import github_com_mwitkow_go_proto_validators "github.com/mwitkow/go-proto-validators"
import proto "github.com/golang/protobuf/proto"
import math "math"
import _ "github.com/golang/protobuf/ptypes/timestamp"
import _ "github.com/mwitkow/go-proto-validators"
import _ "google.golang.org/genproto/googleapis/api/annotations"

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

func (this *OrderId) Validate() error {
	if this.Id == "" {
		return github_com_mwitkow_go_proto_validators.FieldError("Id", fmt.Errorf(`value '%v' must not be an empty string`, this.Id))
	}
	return nil
}
func (this *Order) Validate() error {
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
func (this *Order_Data) Validate() error {
	if this.Attributes != nil {
		if err := github_com_mwitkow_go_proto_validators.CallValidatorIfExists(this.Attributes); err != nil {
			return github_com_mwitkow_go_proto_validators.FieldError("Attributes", err)
		}
	}
	return nil
}
func (this *OrderAttributes) Validate() error {
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
	return nil
}
func (this *NewOrder) Validate() error {
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
func (this *NewOrder_Data) Validate() error {
	if this.Attributes != nil {
		if err := github_com_mwitkow_go_proto_validators.CallValidatorIfExists(this.Attributes); err != nil {
			return github_com_mwitkow_go_proto_validators.FieldError("Attributes", err)
		}
	}
	return nil
}
func (this *NewOrderAttributes) Validate() error {
	if this.Courier == "" {
		return github_com_mwitkow_go_proto_validators.FieldError("Courier", fmt.Errorf(`value '%v' must not be an empty string`, this.Courier))
	}
	if this.CourierAccount == "" {
		return github_com_mwitkow_go_proto_validators.FieldError("CourierAccount", fmt.Errorf(`value '%v' must not be an empty string`, this.CourierAccount))
	}
	if this.Payment == "" {
		return github_com_mwitkow_go_proto_validators.FieldError("Payment", fmt.Errorf(`value '%v' must not be an empty string`, this.Payment))
	}
	if this.PurchaseOrderNum == "" {
		return github_com_mwitkow_go_proto_validators.FieldError("PurchaseOrderNum", fmt.Errorf(`value '%v' must not be an empty string`, this.PurchaseOrderNum))
	}
	if this.Consumer == "" {
		return github_com_mwitkow_go_proto_validators.FieldError("Consumer", fmt.Errorf(`value '%v' must not be an empty string`, this.Consumer))
	}
	if this.Payer == "" {
		return github_com_mwitkow_go_proto_validators.FieldError("Payer", fmt.Errorf(`value '%v' must not be an empty string`, this.Payer))
	}
	if this.Purchaser == "" {
		return github_com_mwitkow_go_proto_validators.FieldError("Purchaser", fmt.Errorf(`value '%v' must not be an empty string`, this.Purchaser))
	}
	for _, item := range this.Items {
		if item == "" {
			return github_com_mwitkow_go_proto_validators.FieldError("Items", fmt.Errorf(`value '%v' must not be an empty string`, item))
		}
	}
	return nil
}
func (this *ExistingOrder) Validate() error {
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
func (this *ExistingOrder_Data) Validate() error {
	if this.Attributes != nil {
		if err := github_com_mwitkow_go_proto_validators.CallValidatorIfExists(this.Attributes); err != nil {
			return github_com_mwitkow_go_proto_validators.FieldError("Attributes", err)
		}
	}
	return nil
}
func (this *ExistingOrderAttributes) Validate() error {
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
	if this.Purchaser == "" {
		return github_com_mwitkow_go_proto_validators.FieldError("Purchaser", fmt.Errorf(`value '%v' must not be an empty string`, this.Purchaser))
	}
	for _, item := range this.Items {
		if item == "" {
			return github_com_mwitkow_go_proto_validators.FieldError("Items", fmt.Errorf(`value '%v' must not be an empty string`, item))
		}
	}
	return nil
}
func (this *OrderUpdate) Validate() error {
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
func (this *OrderUpdate_Data) Validate() error {
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
func (this *OrderUpdateAttributes) Validate() error {
	return nil
}
func (this *OrderCollection) Validate() error {
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
func (this *OrderCollection_Data) Validate() error {
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
func (this *Meta) Validate() error {
	return nil
}
