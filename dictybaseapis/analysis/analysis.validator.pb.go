// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: analysis.proto

package analysis

import (
	fmt "fmt"
	math "math"
	proto "github.com/golang/protobuf/proto"
	_ "google.golang.org/genproto/googleapis/api/annotations"
	_ "."
	_ "github.com/golang/protobuf/ptypes/empty"
	github_com_mwitkow_go_proto_validators "github.com/mwitkow/go-proto-validators"
)

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

func (this *BlastDbRequest) Validate() error {
	if this.TaxonId == "" {
		return github_com_mwitkow_go_proto_validators.FieldError("TaxonId", fmt.Errorf(`value '%v' must not be an empty string`, this.TaxonId))
	}
	if this.Name == "" {
		return github_com_mwitkow_go_proto_validators.FieldError("Name", fmt.Errorf(`value '%v' must not be an empty string`, this.Name))
	}
	if this.Title == "" {
		return github_com_mwitkow_go_proto_validators.FieldError("Title", fmt.Errorf(`value '%v' must not be an empty string`, this.Title))
	}
	if this.Seqtype == "" {
		return github_com_mwitkow_go_proto_validators.FieldError("Seqtype", fmt.Errorf(`value '%v' must not be an empty string`, this.Seqtype))
	}
	return nil
}
func (this *BlastDbParams) Validate() error {
	if this.TaxonId == "" {
		return github_com_mwitkow_go_proto_validators.FieldError("TaxonId", fmt.Errorf(`value '%v' must not be an empty string`, this.TaxonId))
	}
	if this.Name == "" {
		return github_com_mwitkow_go_proto_validators.FieldError("Name", fmt.Errorf(`value '%v' must not be an empty string`, this.Name))
	}
	if this.Title == "" {
		return github_com_mwitkow_go_proto_validators.FieldError("Title", fmt.Errorf(`value '%v' must not be an empty string`, this.Title))
	}
	if this.Seqtype == "" {
		return github_com_mwitkow_go_proto_validators.FieldError("Seqtype", fmt.Errorf(`value '%v' must not be an empty string`, this.Seqtype))
	}
	return nil
}
