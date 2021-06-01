// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: dictybase/publication/publication.proto

package publication

import (
	fmt "fmt"
	math "math"
	proto "github.com/golang/protobuf/proto"
	_ "google.golang.org/protobuf/types/known/timestamppb"
	_ "google.golang.org/protobuf/types/known/emptypb"
	_ "github.com/mwitkow/go-proto-validators"
	github_com_mwitkow_go_proto_validators "github.com/mwitkow/go-proto-validators"
)

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

func (this *PublicationId) Validate() error {
	if this.Id == "" {
		return github_com_mwitkow_go_proto_validators.FieldError("Id", fmt.Errorf(`value '%v' must not be an empty string`, this.Id))
	}
	return nil
}
func (this *Publication) Validate() error {
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
func (this *Publication_Data) Validate() error {
	if this.Attributes != nil {
		if err := github_com_mwitkow_go_proto_validators.CallValidatorIfExists(this.Attributes); err != nil {
			return github_com_mwitkow_go_proto_validators.FieldError("Attributes", err)
		}
	}
	return nil
}
func (this *PublicationAttributes) Validate() error {
	if this.PubDate != nil {
		if err := github_com_mwitkow_go_proto_validators.CallValidatorIfExists(this.PubDate); err != nil {
			return github_com_mwitkow_go_proto_validators.FieldError("PubDate", err)
		}
	}
	for _, item := range this.Authors {
		if item != nil {
			if err := github_com_mwitkow_go_proto_validators.CallValidatorIfExists(item); err != nil {
				return github_com_mwitkow_go_proto_validators.FieldError("Authors", err)
			}
		}
	}
	return nil
}
func (this *NewPublication) Validate() error {
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
func (this *NewPublication_Data) Validate() error {
	if this.Attributes != nil {
		if err := github_com_mwitkow_go_proto_validators.CallValidatorIfExists(this.Attributes); err != nil {
			return github_com_mwitkow_go_proto_validators.FieldError("Attributes", err)
		}
	}
	return nil
}
func (this *PublicationUpdate) Validate() error {
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
func (this *PublicationUpdate_Data) Validate() error {
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
func (this *PublicationUpdateAttributes) Validate() error {
	if this.PubDate != nil {
		if err := github_com_mwitkow_go_proto_validators.CallValidatorIfExists(this.PubDate); err != nil {
			return github_com_mwitkow_go_proto_validators.FieldError("PubDate", err)
		}
	}
	for _, item := range this.Authors {
		if item != nil {
			if err := github_com_mwitkow_go_proto_validators.CallValidatorIfExists(item); err != nil {
				return github_com_mwitkow_go_proto_validators.FieldError("Authors", err)
			}
		}
	}
	return nil
}
func (this *PublicationCollection) Validate() error {
	if len(this.Data) < 1 {
		return github_com_mwitkow_go_proto_validators.FieldError("Data", fmt.Errorf(`value '%v' must contain at least 1 elements`, this.Data))
	}
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
func (this *PublicationCollection_Data) Validate() error {
	if this.Attributes != nil {
		if err := github_com_mwitkow_go_proto_validators.CallValidatorIfExists(this.Attributes); err != nil {
			return github_com_mwitkow_go_proto_validators.FieldError("Attributes", err)
		}
	}
	return nil
}
func (this *ListPublicationParameters) Validate() error {
	return nil
}
func (this *Author) Validate() error {
	return nil
}
func (this *Meta) Validate() error {
	return nil
}