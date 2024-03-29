// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.31.0
// 	protoc        (unknown)
// source: dictybase/api/upload/file.proto

package upload

import (
	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
	reflect "reflect"
	sync "sync"
)

const (
	// Verify that this generated code is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(20 - protoimpl.MinVersion)
	// Verify that runtime/protoimpl is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(protoimpl.MaxVersion - 20)
)

type FileUploadResponse_Status int32

const (
	// a new file
	FileUploadResponse_CREATED FileUploadResponse_Status = 0
	// an existing file
	FileUploadResponse_UPDATED FileUploadResponse_Status = 1
)

// Enum value maps for FileUploadResponse_Status.
var (
	FileUploadResponse_Status_name = map[int32]string{
		0: "CREATED",
		1: "UPDATED",
	}
	FileUploadResponse_Status_value = map[string]int32{
		"CREATED": 0,
		"UPDATED": 1,
	}
)

func (x FileUploadResponse_Status) Enum() *FileUploadResponse_Status {
	p := new(FileUploadResponse_Status)
	*p = x
	return p
}

func (x FileUploadResponse_Status) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (FileUploadResponse_Status) Descriptor() protoreflect.EnumDescriptor {
	return file_dictybase_api_upload_file_proto_enumTypes[0].Descriptor()
}

func (FileUploadResponse_Status) Type() protoreflect.EnumType {
	return &file_dictybase_api_upload_file_proto_enumTypes[0]
}

func (x FileUploadResponse_Status) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use FileUploadResponse_Status.Descriptor instead.
func (FileUploadResponse_Status) EnumDescriptor() ([]byte, []int) {
	return file_dictybase_api_upload_file_proto_rawDescGZIP(), []int{1, 0}
}

// FileUploadRequest defines requests parameters for uploading
// a file through protocol buffer. The file is expected to be uploaded
// grpc client side streaming
type FileUploadRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// name of the file
	Name string `protobuf:"bytes,1,opt,name=name,proto3" json:"name,omitempty"`
	// content of the file, expected to be streamed in chunk
	Content []byte `protobuf:"bytes,2,opt,name=content,proto3" json:"content,omitempty"`
}

func (x *FileUploadRequest) Reset() {
	*x = FileUploadRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_dictybase_api_upload_file_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *FileUploadRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*FileUploadRequest) ProtoMessage() {}

func (x *FileUploadRequest) ProtoReflect() protoreflect.Message {
	mi := &file_dictybase_api_upload_file_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use FileUploadRequest.ProtoReflect.Descriptor instead.
func (*FileUploadRequest) Descriptor() ([]byte, []int) {
	return file_dictybase_api_upload_file_proto_rawDescGZIP(), []int{0}
}

func (x *FileUploadRequest) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

func (x *FileUploadRequest) GetContent() []byte {
	if x != nil {
		return x.Content
	}
	return nil
}

// FileUploadResponse defines response received after completion of
// file upload
type FileUploadResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// Status gives the status after file upload
	Status FileUploadResponse_Status `protobuf:"varint,1,opt,name=status,proto3,enum=dictybase.api.upload.FileUploadResponse_Status" json:"status,omitempty"`
	// message send by the server after file upload
	Msg string `protobuf:"bytes,2,opt,name=msg,proto3" json:"msg,omitempty"`
}

func (x *FileUploadResponse) Reset() {
	*x = FileUploadResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_dictybase_api_upload_file_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *FileUploadResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*FileUploadResponse) ProtoMessage() {}

func (x *FileUploadResponse) ProtoReflect() protoreflect.Message {
	mi := &file_dictybase_api_upload_file_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use FileUploadResponse.ProtoReflect.Descriptor instead.
func (*FileUploadResponse) Descriptor() ([]byte, []int) {
	return file_dictybase_api_upload_file_proto_rawDescGZIP(), []int{1}
}

func (x *FileUploadResponse) GetStatus() FileUploadResponse_Status {
	if x != nil {
		return x.Status
	}
	return FileUploadResponse_CREATED
}

func (x *FileUploadResponse) GetMsg() string {
	if x != nil {
		return x.Msg
	}
	return ""
}

var File_dictybase_api_upload_file_proto protoreflect.FileDescriptor

var file_dictybase_api_upload_file_proto_rawDesc = []byte{
	0x0a, 0x1f, 0x64, 0x69, 0x63, 0x74, 0x79, 0x62, 0x61, 0x73, 0x65, 0x2f, 0x61, 0x70, 0x69, 0x2f,
	0x75, 0x70, 0x6c, 0x6f, 0x61, 0x64, 0x2f, 0x66, 0x69, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x12, 0x14, 0x64, 0x69, 0x63, 0x74, 0x79, 0x62, 0x61, 0x73, 0x65, 0x2e, 0x61, 0x70, 0x69,
	0x2e, 0x75, 0x70, 0x6c, 0x6f, 0x61, 0x64, 0x22, 0x41, 0x0a, 0x11, 0x46, 0x69, 0x6c, 0x65, 0x55,
	0x70, 0x6c, 0x6f, 0x61, 0x64, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x12, 0x0a, 0x04,
	0x6e, 0x61, 0x6d, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x6e, 0x61, 0x6d, 0x65,
	0x12, 0x18, 0x0a, 0x07, 0x63, 0x6f, 0x6e, 0x74, 0x65, 0x6e, 0x74, 0x18, 0x02, 0x20, 0x01, 0x28,
	0x0c, 0x52, 0x07, 0x63, 0x6f, 0x6e, 0x74, 0x65, 0x6e, 0x74, 0x22, 0x93, 0x01, 0x0a, 0x12, 0x46,
	0x69, 0x6c, 0x65, 0x55, 0x70, 0x6c, 0x6f, 0x61, 0x64, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73,
	0x65, 0x12, 0x47, 0x0a, 0x06, 0x73, 0x74, 0x61, 0x74, 0x75, 0x73, 0x18, 0x01, 0x20, 0x01, 0x28,
	0x0e, 0x32, 0x2f, 0x2e, 0x64, 0x69, 0x63, 0x74, 0x79, 0x62, 0x61, 0x73, 0x65, 0x2e, 0x61, 0x70,
	0x69, 0x2e, 0x75, 0x70, 0x6c, 0x6f, 0x61, 0x64, 0x2e, 0x46, 0x69, 0x6c, 0x65, 0x55, 0x70, 0x6c,
	0x6f, 0x61, 0x64, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x2e, 0x53, 0x74, 0x61, 0x74,
	0x75, 0x73, 0x52, 0x06, 0x73, 0x74, 0x61, 0x74, 0x75, 0x73, 0x12, 0x10, 0x0a, 0x03, 0x6d, 0x73,
	0x67, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x03, 0x6d, 0x73, 0x67, 0x22, 0x22, 0x0a, 0x06,
	0x53, 0x74, 0x61, 0x74, 0x75, 0x73, 0x12, 0x0b, 0x0a, 0x07, 0x43, 0x52, 0x45, 0x41, 0x54, 0x45,
	0x44, 0x10, 0x00, 0x12, 0x0b, 0x0a, 0x07, 0x55, 0x50, 0x44, 0x41, 0x54, 0x45, 0x44, 0x10, 0x01,
	0x42, 0x7d, 0x0a, 0x18, 0x6f, 0x72, 0x67, 0x2e, 0x64, 0x69, 0x63, 0x74, 0x79, 0x62, 0x61, 0x73,
	0x65, 0x2e, 0x61, 0x70, 0x69, 0x2e, 0x75, 0x70, 0x6c, 0x6f, 0x61, 0x64, 0x42, 0x0f, 0x46, 0x69,
	0x6c, 0x65, 0x55, 0x70, 0x6c, 0x6f, 0x61, 0x64, 0x50, 0x72, 0x6f, 0x74, 0x6f, 0x50, 0x01, 0x5a,
	0x40, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x64, 0x69, 0x63, 0x74,
	0x79, 0x42, 0x61, 0x73, 0x65, 0x2f, 0x67, 0x6f, 0x2d, 0x67, 0x65, 0x6e, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x2f, 0x64, 0x69, 0x63, 0x74, 0x79, 0x62, 0x61, 0x73, 0x65, 0x61, 0x70, 0x69, 0x73, 0x2f,
	0x61, 0x70, 0x69, 0x2f, 0x75, 0x70, 0x6c, 0x6f, 0x61, 0x64, 0x3b, 0x75, 0x70, 0x6c, 0x6f, 0x61,
	0x64, 0xf8, 0x01, 0x01, 0xa2, 0x02, 0x08, 0x44, 0x49, 0x43, 0x54, 0x59, 0x41, 0x50, 0x49, 0x62,
	0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_dictybase_api_upload_file_proto_rawDescOnce sync.Once
	file_dictybase_api_upload_file_proto_rawDescData = file_dictybase_api_upload_file_proto_rawDesc
)

func file_dictybase_api_upload_file_proto_rawDescGZIP() []byte {
	file_dictybase_api_upload_file_proto_rawDescOnce.Do(func() {
		file_dictybase_api_upload_file_proto_rawDescData = protoimpl.X.CompressGZIP(file_dictybase_api_upload_file_proto_rawDescData)
	})
	return file_dictybase_api_upload_file_proto_rawDescData
}

var file_dictybase_api_upload_file_proto_enumTypes = make([]protoimpl.EnumInfo, 1)
var file_dictybase_api_upload_file_proto_msgTypes = make([]protoimpl.MessageInfo, 2)
var file_dictybase_api_upload_file_proto_goTypes = []interface{}{
	(FileUploadResponse_Status)(0), // 0: dictybase.api.upload.FileUploadResponse.Status
	(*FileUploadRequest)(nil),      // 1: dictybase.api.upload.FileUploadRequest
	(*FileUploadResponse)(nil),     // 2: dictybase.api.upload.FileUploadResponse
}
var file_dictybase_api_upload_file_proto_depIdxs = []int32{
	0, // 0: dictybase.api.upload.FileUploadResponse.status:type_name -> dictybase.api.upload.FileUploadResponse.Status
	1, // [1:1] is the sub-list for method output_type
	1, // [1:1] is the sub-list for method input_type
	1, // [1:1] is the sub-list for extension type_name
	1, // [1:1] is the sub-list for extension extendee
	0, // [0:1] is the sub-list for field type_name
}

func init() { file_dictybase_api_upload_file_proto_init() }
func file_dictybase_api_upload_file_proto_init() {
	if File_dictybase_api_upload_file_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_dictybase_api_upload_file_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*FileUploadRequest); i {
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
		file_dictybase_api_upload_file_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*FileUploadResponse); i {
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
			RawDescriptor: file_dictybase_api_upload_file_proto_rawDesc,
			NumEnums:      1,
			NumMessages:   2,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_dictybase_api_upload_file_proto_goTypes,
		DependencyIndexes: file_dictybase_api_upload_file_proto_depIdxs,
		EnumInfos:         file_dictybase_api_upload_file_proto_enumTypes,
		MessageInfos:      file_dictybase_api_upload_file_proto_msgTypes,
	}.Build()
	File_dictybase_api_upload_file_proto = out.File
	file_dictybase_api_upload_file_proto_rawDesc = nil
	file_dictybase_api_upload_file_proto_goTypes = nil
	file_dictybase_api_upload_file_proto_depIdxs = nil
}
