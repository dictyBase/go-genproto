// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.26.0
// 	protoc        v3.17.1
// source: dictybase/analysis/analysis.proto

package analysis

import (
	_ "github.com/mwitkow/go-proto-validators"
	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
	emptypb "google.golang.org/protobuf/types/known/emptypb"
	reflect "reflect"
	sync "sync"
)

const (
	// Verify that this generated code is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(20 - protoimpl.MinVersion)
	// Verify that runtime/protoimpl is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(protoimpl.MaxVersion - 20)
)

// Various parameters required for creating a blast database
type BlastDbRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// sequence needed for database, provided in chunks
	SeqChunk []byte `protobuf:"bytes,1,opt,name=seq_chunk,json=seqChunk,proto3" json:"seq_chunk,omitempty"`
	// ncbi taxonid of the organism
	TaxonId string `protobuf:"bytes,2,opt,name=taxon_id,json=taxonId,proto3" json:"taxon_id,omitempty"`
	// name of the database
	Name string `protobuf:"bytes,3,opt,name=name,proto3" json:"name,omitempty"`
	// human readable description of the database
	Title string `protobuf:"bytes,4,opt,name=title,proto3" json:"title,omitempty"`
	// Type of sequence, nucleotide or protein
	Seqtype string `protobuf:"bytes,5,opt,name=seqtype,proto3" json:"seqtype,omitempty"`
}

func (x *BlastDbRequest) Reset() {
	*x = BlastDbRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_dictybase_analysis_analysis_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *BlastDbRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*BlastDbRequest) ProtoMessage() {}

func (x *BlastDbRequest) ProtoReflect() protoreflect.Message {
	mi := &file_dictybase_analysis_analysis_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use BlastDbRequest.ProtoReflect.Descriptor instead.
func (*BlastDbRequest) Descriptor() ([]byte, []int) {
	return file_dictybase_analysis_analysis_proto_rawDescGZIP(), []int{0}
}

func (x *BlastDbRequest) GetSeqChunk() []byte {
	if x != nil {
		return x.SeqChunk
	}
	return nil
}

func (x *BlastDbRequest) GetTaxonId() string {
	if x != nil {
		return x.TaxonId
	}
	return ""
}

func (x *BlastDbRequest) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

func (x *BlastDbRequest) GetTitle() string {
	if x != nil {
		return x.Title
	}
	return ""
}

func (x *BlastDbRequest) GetSeqtype() string {
	if x != nil {
		return x.Seqtype
	}
	return ""
}

// Similar to the BlastDbParams except the sequence chunks
type BlastDbParams struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// name of the sequence file
	FileName string `protobuf:"bytes,1,opt,name=file_name,json=fileName,proto3" json:"file_name,omitempty"`
	// ncbi taxonid of the organism
	TaxonId string `protobuf:"bytes,2,opt,name=taxon_id,json=taxonId,proto3" json:"taxon_id,omitempty"`
	// name of the database
	Name string `protobuf:"bytes,3,opt,name=name,proto3" json:"name,omitempty"`
	// human readable description of the database
	Title string `protobuf:"bytes,4,opt,name=title,proto3" json:"title,omitempty"`
	// Type of sequence, nucleotide or protein
	Seqtype string `protobuf:"bytes,6,opt,name=seqtype,proto3" json:"seqtype,omitempty"`
}

func (x *BlastDbParams) Reset() {
	*x = BlastDbParams{}
	if protoimpl.UnsafeEnabled {
		mi := &file_dictybase_analysis_analysis_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *BlastDbParams) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*BlastDbParams) ProtoMessage() {}

func (x *BlastDbParams) ProtoReflect() protoreflect.Message {
	mi := &file_dictybase_analysis_analysis_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use BlastDbParams.ProtoReflect.Descriptor instead.
func (*BlastDbParams) Descriptor() ([]byte, []int) {
	return file_dictybase_analysis_analysis_proto_rawDescGZIP(), []int{1}
}

func (x *BlastDbParams) GetFileName() string {
	if x != nil {
		return x.FileName
	}
	return ""
}

func (x *BlastDbParams) GetTaxonId() string {
	if x != nil {
		return x.TaxonId
	}
	return ""
}

func (x *BlastDbParams) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

func (x *BlastDbParams) GetTitle() string {
	if x != nil {
		return x.Title
	}
	return ""
}

func (x *BlastDbParams) GetSeqtype() string {
	if x != nil {
		return x.Seqtype
	}
	return ""
}

var File_dictybase_analysis_analysis_proto protoreflect.FileDescriptor

var file_dictybase_analysis_analysis_proto_rawDesc = []byte{
	0x0a, 0x21, 0x64, 0x69, 0x63, 0x74, 0x79, 0x62, 0x61, 0x73, 0x65, 0x2f, 0x61, 0x6e, 0x61, 0x6c,
	0x79, 0x73, 0x69, 0x73, 0x2f, 0x61, 0x6e, 0x61, 0x6c, 0x79, 0x73, 0x69, 0x73, 0x2e, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x12, 0x12, 0x64, 0x69, 0x63, 0x74, 0x79, 0x62, 0x61, 0x73, 0x65, 0x2e, 0x61,
	0x6e, 0x61, 0x6c, 0x79, 0x73, 0x69, 0x73, 0x1a, 0x1b, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2f, 0x65, 0x6d, 0x70, 0x74, 0x79, 0x2e, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x36, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d,
	0x2f, 0x6d, 0x77, 0x69, 0x74, 0x6b, 0x6f, 0x77, 0x2f, 0x67, 0x6f, 0x2d, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x2d, 0x76, 0x61, 0x6c, 0x69, 0x64, 0x61, 0x74, 0x6f, 0x72, 0x73, 0x2f, 0x76, 0x61, 0x6c,
	0x69, 0x64, 0x61, 0x74, 0x6f, 0x72, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0xac, 0x01, 0x0a,
	0x0e, 0x42, 0x6c, 0x61, 0x73, 0x74, 0x44, 0x62, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12,
	0x1b, 0x0a, 0x09, 0x73, 0x65, 0x71, 0x5f, 0x63, 0x68, 0x75, 0x6e, 0x6b, 0x18, 0x01, 0x20, 0x01,
	0x28, 0x0c, 0x52, 0x08, 0x73, 0x65, 0x71, 0x43, 0x68, 0x75, 0x6e, 0x6b, 0x12, 0x21, 0x0a, 0x08,
	0x74, 0x61, 0x78, 0x6f, 0x6e, 0x5f, 0x69, 0x64, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x42, 0x06,
	0xe2, 0xdf, 0x1f, 0x02, 0x58, 0x01, 0x52, 0x07, 0x74, 0x61, 0x78, 0x6f, 0x6e, 0x49, 0x64, 0x12,
	0x1a, 0x0a, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x42, 0x06, 0xe2,
	0xdf, 0x1f, 0x02, 0x58, 0x01, 0x52, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x12, 0x1c, 0x0a, 0x05, 0x74,
	0x69, 0x74, 0x6c, 0x65, 0x18, 0x04, 0x20, 0x01, 0x28, 0x09, 0x42, 0x06, 0xe2, 0xdf, 0x1f, 0x02,
	0x58, 0x01, 0x52, 0x05, 0x74, 0x69, 0x74, 0x6c, 0x65, 0x12, 0x20, 0x0a, 0x07, 0x73, 0x65, 0x71,
	0x74, 0x79, 0x70, 0x65, 0x18, 0x05, 0x20, 0x01, 0x28, 0x09, 0x42, 0x06, 0xe2, 0xdf, 0x1f, 0x02,
	0x58, 0x01, 0x52, 0x07, 0x73, 0x65, 0x71, 0x74, 0x79, 0x70, 0x65, 0x22, 0xab, 0x01, 0x0a, 0x0d,
	0x42, 0x6c, 0x61, 0x73, 0x74, 0x44, 0x62, 0x50, 0x61, 0x72, 0x61, 0x6d, 0x73, 0x12, 0x1b, 0x0a,
	0x09, 0x66, 0x69, 0x6c, 0x65, 0x5f, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09,
	0x52, 0x08, 0x66, 0x69, 0x6c, 0x65, 0x4e, 0x61, 0x6d, 0x65, 0x12, 0x21, 0x0a, 0x08, 0x74, 0x61,
	0x78, 0x6f, 0x6e, 0x5f, 0x69, 0x64, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x42, 0x06, 0xe2, 0xdf,
	0x1f, 0x02, 0x58, 0x01, 0x52, 0x07, 0x74, 0x61, 0x78, 0x6f, 0x6e, 0x49, 0x64, 0x12, 0x1a, 0x0a,
	0x04, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x42, 0x06, 0xe2, 0xdf, 0x1f,
	0x02, 0x58, 0x01, 0x52, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x12, 0x1c, 0x0a, 0x05, 0x74, 0x69, 0x74,
	0x6c, 0x65, 0x18, 0x04, 0x20, 0x01, 0x28, 0x09, 0x42, 0x06, 0xe2, 0xdf, 0x1f, 0x02, 0x58, 0x01,
	0x52, 0x05, 0x74, 0x69, 0x74, 0x6c, 0x65, 0x12, 0x20, 0x0a, 0x07, 0x73, 0x65, 0x71, 0x74, 0x79,
	0x70, 0x65, 0x18, 0x06, 0x20, 0x01, 0x28, 0x09, 0x42, 0x06, 0xe2, 0xdf, 0x1f, 0x02, 0x58, 0x01,
	0x52, 0x07, 0x73, 0x65, 0x71, 0x74, 0x79, 0x70, 0x65, 0x32, 0x62, 0x0a, 0x0f, 0x41, 0x6e, 0x61,
	0x6c, 0x79, 0x73, 0x69, 0x73, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x12, 0x4f, 0x0a, 0x0d,
	0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x42, 0x6c, 0x61, 0x73, 0x74, 0x44, 0x62, 0x12, 0x22, 0x2e,
	0x64, 0x69, 0x63, 0x74, 0x79, 0x62, 0x61, 0x73, 0x65, 0x2e, 0x61, 0x6e, 0x61, 0x6c, 0x79, 0x73,
	0x69, 0x73, 0x2e, 0x42, 0x6c, 0x61, 0x73, 0x74, 0x44, 0x62, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73,
	0x74, 0x1a, 0x16, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x62, 0x75, 0x66, 0x2e, 0x45, 0x6d, 0x70, 0x74, 0x79, 0x22, 0x00, 0x28, 0x01, 0x42, 0x79, 0x0a,
	0x16, 0x6f, 0x72, 0x67, 0x2e, 0x64, 0x69, 0x63, 0x74, 0x79, 0x62, 0x61, 0x73, 0x65, 0x2e, 0x61,
	0x6e, 0x61, 0x6c, 0x79, 0x73, 0x69, 0x73, 0x42, 0x0d, 0x41, 0x6e, 0x61, 0x6c, 0x79, 0x73, 0x69,
	0x73, 0x50, 0x72, 0x6f, 0x74, 0x6f, 0x50, 0x01, 0x5a, 0x40, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62,
	0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x64, 0x69, 0x63, 0x74, 0x79, 0x42, 0x61, 0x73, 0x65, 0x2f, 0x67,
	0x6f, 0x2d, 0x67, 0x65, 0x6e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2f, 0x64, 0x69, 0x63, 0x74, 0x79,
	0x62, 0x61, 0x73, 0x65, 0x61, 0x70, 0x69, 0x73, 0x2f, 0x61, 0x6e, 0x61, 0x6c, 0x79, 0x73, 0x69,
	0x73, 0x3b, 0x61, 0x6e, 0x61, 0x6c, 0x79, 0x73, 0x69, 0x73, 0xf8, 0x01, 0x01, 0xa2, 0x02, 0x08,
	0x44, 0x49, 0x43, 0x54, 0x59, 0x41, 0x50, 0x49, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_dictybase_analysis_analysis_proto_rawDescOnce sync.Once
	file_dictybase_analysis_analysis_proto_rawDescData = file_dictybase_analysis_analysis_proto_rawDesc
)

func file_dictybase_analysis_analysis_proto_rawDescGZIP() []byte {
	file_dictybase_analysis_analysis_proto_rawDescOnce.Do(func() {
		file_dictybase_analysis_analysis_proto_rawDescData = protoimpl.X.CompressGZIP(file_dictybase_analysis_analysis_proto_rawDescData)
	})
	return file_dictybase_analysis_analysis_proto_rawDescData
}

var file_dictybase_analysis_analysis_proto_msgTypes = make([]protoimpl.MessageInfo, 2)
var file_dictybase_analysis_analysis_proto_goTypes = []interface{}{
	(*BlastDbRequest)(nil), // 0: dictybase.analysis.BlastDbRequest
	(*BlastDbParams)(nil),  // 1: dictybase.analysis.BlastDbParams
	(*emptypb.Empty)(nil),  // 2: google.protobuf.Empty
}
var file_dictybase_analysis_analysis_proto_depIdxs = []int32{
	0, // 0: dictybase.analysis.AnalysisService.CreateBlastDb:input_type -> dictybase.analysis.BlastDbRequest
	2, // 1: dictybase.analysis.AnalysisService.CreateBlastDb:output_type -> google.protobuf.Empty
	1, // [1:2] is the sub-list for method output_type
	0, // [0:1] is the sub-list for method input_type
	0, // [0:0] is the sub-list for extension type_name
	0, // [0:0] is the sub-list for extension extendee
	0, // [0:0] is the sub-list for field type_name
}

func init() { file_dictybase_analysis_analysis_proto_init() }
func file_dictybase_analysis_analysis_proto_init() {
	if File_dictybase_analysis_analysis_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_dictybase_analysis_analysis_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*BlastDbRequest); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_dictybase_analysis_analysis_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*BlastDbParams); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
	}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_dictybase_analysis_analysis_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   2,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_dictybase_analysis_analysis_proto_goTypes,
		DependencyIndexes: file_dictybase_analysis_analysis_proto_depIdxs,
		MessageInfos:      file_dictybase_analysis_analysis_proto_msgTypes,
	}.Build()
	File_dictybase_analysis_analysis_proto = out.File
	file_dictybase_analysis_analysis_proto_rawDesc = nil
	file_dictybase_analysis_analysis_proto_goTypes = nil
	file_dictybase_analysis_analysis_proto_depIdxs = nil
}
