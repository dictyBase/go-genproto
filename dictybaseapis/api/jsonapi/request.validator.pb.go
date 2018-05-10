// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: request.proto

package jsonapi

import fmt "fmt"
import github_com_mwitkow_go_proto_validators "github.com/mwitkow/go-proto-validators"
import proto "github.com/golang/protobuf/proto"
import math "math"
import _ "github.com/mwitkow/go-proto-validators"

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

func (this *GetRequest) Validate() error {
	return nil
}
func (this *GetEmailRequest) Validate() error {
	return nil
}
func (this *GetRequestWithFields) Validate() error {
	return nil
}
func (this *RelationshipRequest) Validate() error {
	return nil
}
func (this *RelationshipRequestWithPagination) Validate() error {
	return nil
}
func (this *ListRequest) Validate() error {
	return nil
}
func (this *SimpleListRequest) Validate() error {
	return nil
}
func (this *DeleteRequest) Validate() error {
	return nil
}
func (this *IdRequest) Validate() error {
	return nil
}
func (this *IdResponse) Validate() error {
	return nil
}
func (this *HealthzIdRequest) Validate() error {
	if !(this.Id > 0) {
		return github_com_mwitkow_go_proto_validators.FieldError("Id", fmt.Errorf(`value '%v' must be greater than '0'`, this.Id))
	}
	if !(this.Id < 2) {
		return github_com_mwitkow_go_proto_validators.FieldError("Id", fmt.Errorf(`value '%v' must be less than '2'`, this.Id))
	}
	return nil
}
func (this *ExistResponse) Validate() error {
	return nil
}
