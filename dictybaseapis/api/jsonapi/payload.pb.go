// Code generated by protoc-gen-go. DO NOT EDIT.
// source: payload.proto

/*
Package jsonapi is a generated protocol buffer package.

It is generated from these files:
	payload.proto
	request.proto

It has these top-level messages:
	Data
	DataCollection
	Links
	PaginationLinks
	Pagination
	Meta
	GetRequest
	RelationshipRequest
	ListRequest
	DeleteRequest
*/
package jsonapi

import proto "github.com/golang/protobuf/proto"
import fmt "fmt"
import math "math"

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// This is a compile-time assertion to ensure that this generated file
// is compatible with the proto package it is being compiled against.
// A compilation error at this line likely means your copy of the
// proto package needs to be updated.
const _ = proto.ProtoPackageIsVersion2 // please upgrade the proto package

// A [resource identifier object](http://jsonapi.org/format/#document-resource-identifier-objects).
type Data struct {
	// The resource name.
	Type string `protobuf:"bytes,1,opt,name=type" json:"type,omitempty"`
	// Unique id.
	Id int64 `protobuf:"varint,2,opt,name=id" json:"id,omitempty"`
}

func (m *Data) Reset()                    { *m = Data{} }
func (m *Data) String() string            { return proto.CompactTextString(m) }
func (*Data) ProtoMessage()               {}
func (*Data) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{0} }

func (m *Data) GetType() string {
	if m != nil {
		return m.Type
	}
	return ""
}

func (m *Data) GetId() int64 {
	if m != nil {
		return m.Id
	}
	return 0
}

// Definition for resource identifier collection objects.
type DataCollection struct {
	Id   int64   `protobuf:"varint,2,opt,name=id" json:"id,omitempty"`
	Data []*Data `protobuf:"bytes,1,rep,name=data" json:"data,omitempty"`
}

func (m *DataCollection) Reset()                    { *m = DataCollection{} }
func (m *DataCollection) String() string            { return proto.CompactTextString(m) }
func (*DataCollection) ProtoMessage()               {}
func (*DataCollection) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{1} }

func (m *DataCollection) GetId() int64 {
	if m != nil {
		return m.Id
	}
	return 0
}

func (m *DataCollection) GetData() []*Data {
	if m != nil {
		return m.Data
	}
	return nil
}

// A container for http links.
type Links struct {
	// A http link. It points to the resource itself.
	Self string `protobuf:"bytes,1,opt,name=self" json:"self,omitempty"`
	// A http link. It points to a related resource.
	Related string `protobuf:"bytes,2,opt,name=related" json:"related,omitempty"`
}

func (m *Links) Reset()                    { *m = Links{} }
func (m *Links) String() string            { return proto.CompactTextString(m) }
func (*Links) ProtoMessage()               {}
func (*Links) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{2} }

func (m *Links) GetSelf() string {
	if m != nil {
		return m.Self
	}
	return ""
}

func (m *Links) GetRelated() string {
	if m != nil {
		return m.Related
	}
	return ""
}

// A container for pagination links.
type PaginationLinks struct {
	// A http link to the resource itself.
	Self string `protobuf:"bytes,1,opt,name=self" json:"self,omitempty"`
	// A http link to the next page of data.
	Next string `protobuf:"bytes,2,opt,name=next" json:"next,omitempty"`
	// A http link to the previous page of data.
	Prev string `protobuf:"bytes,3,opt,name=prev" json:"prev,omitempty"`
	// A http link to the last page of data.
	Last string `protobuf:"bytes,4,opt,name=last" json:"last,omitempty"`
	// A http link to the first page of data.
	First string `protobuf:"bytes,5,opt,name=first" json:"first,omitempty"`
}

func (m *PaginationLinks) Reset()                    { *m = PaginationLinks{} }
func (m *PaginationLinks) String() string            { return proto.CompactTextString(m) }
func (*PaginationLinks) ProtoMessage()               {}
func (*PaginationLinks) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{3} }

func (m *PaginationLinks) GetSelf() string {
	if m != nil {
		return m.Self
	}
	return ""
}

func (m *PaginationLinks) GetNext() string {
	if m != nil {
		return m.Next
	}
	return ""
}

func (m *PaginationLinks) GetPrev() string {
	if m != nil {
		return m.Prev
	}
	return ""
}

func (m *PaginationLinks) GetLast() string {
	if m != nil {
		return m.Last
	}
	return ""
}

func (m *PaginationLinks) GetFirst() string {
	if m != nil {
		return m.First
	}
	return ""
}

// A container for various pagination properties
type Pagination struct {
	// Total number of entries, regardless of pages.
	Records int64 `protobuf:"varint,1,opt,name=records" json:"records,omitempty"`
	// Total number of pages.
	Total int64 `protobuf:"varint,2,opt,name=total" json:"total,omitempty"`
	// Number of entries per page.
	Size int64 `protobuf:"varint,3,opt,name=size" json:"size,omitempty"`
	// Current page number.
	Number int64 `protobuf:"varint,4,opt,name=number" json:"number,omitempty"`
}

func (m *Pagination) Reset()                    { *m = Pagination{} }
func (m *Pagination) String() string            { return proto.CompactTextString(m) }
func (*Pagination) ProtoMessage()               {}
func (*Pagination) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{4} }

func (m *Pagination) GetRecords() int64 {
	if m != nil {
		return m.Records
	}
	return 0
}

func (m *Pagination) GetTotal() int64 {
	if m != nil {
		return m.Total
	}
	return 0
}

func (m *Pagination) GetSize() int64 {
	if m != nil {
		return m.Size
	}
	return 0
}

func (m *Pagination) GetNumber() int64 {
	if m != nil {
		return m.Number
	}
	return 0
}

// Top level meta container.
type Meta struct {
	Pagination *Pagination `protobuf:"bytes,1,opt,name=pagination" json:"pagination,omitempty"`
}

func (m *Meta) Reset()                    { *m = Meta{} }
func (m *Meta) String() string            { return proto.CompactTextString(m) }
func (*Meta) ProtoMessage()               {}
func (*Meta) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{5} }

func (m *Meta) GetPagination() *Pagination {
	if m != nil {
		return m.Pagination
	}
	return nil
}

func init() {
	proto.RegisterType((*Data)(nil), "dictybase.api.jsonapi.Data")
	proto.RegisterType((*DataCollection)(nil), "dictybase.api.jsonapi.DataCollection")
	proto.RegisterType((*Links)(nil), "dictybase.api.jsonapi.Links")
	proto.RegisterType((*PaginationLinks)(nil), "dictybase.api.jsonapi.PaginationLinks")
	proto.RegisterType((*Pagination)(nil), "dictybase.api.jsonapi.Pagination")
	proto.RegisterType((*Meta)(nil), "dictybase.api.jsonapi.Meta")
}

func init() { proto.RegisterFile("payload.proto", fileDescriptor0) }

var fileDescriptor0 = []byte{
	// 386 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x74, 0x92, 0x4f, 0x6b, 0xdb, 0x40,
	0x10, 0xc5, 0xd1, 0x1f, 0xbb, 0xf5, 0xb8, 0x75, 0x61, 0x69, 0x8b, 0x4a, 0x2f, 0xae, 0x4e, 0xa6,
	0x50, 0x09, 0x5c, 0x7a, 0xea, 0xc9, 0xb2, 0x2f, 0x86, 0x16, 0xd4, 0xa5, 0x97, 0xe4, 0x36, 0x92,
	0xd6, 0xf2, 0x26, 0xb2, 0x76, 0xd1, 0xae, 0x93, 0x38, 0x90, 0x2f, 0x93, 0x4f, 0x99, 0x63, 0xd8,
	0x95, 0x6c, 0x87, 0x60, 0x9f, 0xf6, 0xcd, 0x63, 0xde, 0xfc, 0x86, 0x65, 0xe0, 0xbd, 0xc4, 0x5d,
	0x25, 0xb0, 0x88, 0x64, 0x23, 0xb4, 0x20, 0x9f, 0x0a, 0x9e, 0xeb, 0x5d, 0x86, 0x8a, 0x45, 0x28,
	0x79, 0x74, 0xa5, 0x44, 0x8d, 0x92, 0x87, 0xdf, 0xc1, 0x5f, 0xa0, 0x46, 0x42, 0xc0, 0xd7, 0x3b,
	0xc9, 0x02, 0x67, 0xec, 0x4c, 0x06, 0xd4, 0x6a, 0x32, 0x02, 0x97, 0x17, 0x81, 0x3b, 0x76, 0x26,
	0x1e, 0x75, 0x79, 0x11, 0xfe, 0x83, 0x91, 0xe9, 0x9d, 0x8b, 0xaa, 0x62, 0xb9, 0xe6, 0xa2, 0x7e,
	0xdd, 0x41, 0x62, 0xf0, 0x0b, 0xd4, 0x18, 0x38, 0x63, 0x6f, 0x32, 0x9c, 0x7e, 0x8d, 0x4e, 0x32,
	0x23, 0x33, 0x84, 0xda, 0xc6, 0xf0, 0x17, 0xf4, 0xfe, 0xf0, 0xfa, 0x5a, 0x19, 0xbe, 0x62, 0xd5,
	0x6a, 0xcf, 0x37, 0x9a, 0x04, 0xf0, 0xa6, 0x61, 0x15, 0x6a, 0xd6, 0x22, 0x06, 0x74, 0x5f, 0x86,
	0xb7, 0xf0, 0x21, 0xc5, 0x92, 0xd7, 0x68, 0xb6, 0x38, 0x3f, 0x80, 0x80, 0x5f, 0xb3, 0x3b, 0xdd,
	0xa5, 0xad, 0x36, 0x9e, 0x6c, 0xd8, 0x4d, 0xe0, 0xb5, 0x9e, 0xd1, 0xc6, 0xab, 0x50, 0xe9, 0xc0,
	0x6f, 0x3d, 0xa3, 0xc9, 0x47, 0xe8, 0xad, 0x78, 0xa3, 0x74, 0xd0, 0xb3, 0x66, 0x5b, 0x84, 0x6b,
	0x80, 0x23, 0xb8, 0x5d, 0x30, 0x17, 0x4d, 0xa1, 0x2c, 0xd6, 0xa3, 0xfb, 0xd2, 0xa4, 0xb5, 0xd0,
	0x58, 0x75, 0x7f, 0xd3, 0x16, 0x76, 0x47, 0x7e, 0xcf, 0x2c, 0xdb, 0xa3, 0x56, 0x93, 0xcf, 0xd0,
	0xaf, 0xb7, 0x9b, 0x8c, 0x35, 0x96, 0xee, 0xd1, 0xae, 0x0a, 0x97, 0xe0, 0xff, 0x65, 0x1a, 0xc9,
	0x0c, 0x40, 0x1e, 0x88, 0x16, 0x33, 0x9c, 0x7e, 0x3b, 0xf3, 0xb1, 0xc7, 0xd5, 0xe8, 0x8b, 0x50,
	0xf2, 0x00, 0x5f, 0x44, 0x53, 0x9e, 0xce, 0x24, 0xef, 0xd2, 0xf6, 0x4c, 0x52, 0x73, 0x25, 0xa9,
	0x73, 0x99, 0x94, 0x5c, 0xaf, 0xb7, 0x59, 0x94, 0x8b, 0x4d, 0x6c, 0x13, 0x09, 0x2a, 0x16, 0x97,
	0xe2, 0x47, 0xc9, 0x6a, 0x7b, 0x49, 0xf1, 0x61, 0x0e, 0x4a, 0xae, 0x62, 0x94, 0x3c, 0xee, 0x66,
	0xfd, 0xee, 0xde, 0x27, 0xc7, 0x79, 0x74, 0xdf, 0x2e, 0x96, 0xf3, 0xff, 0x17, 0xb3, 0x74, 0x99,
	0xf5, 0x6d, 0xec, 0xe7, 0x73, 0x00, 0x00, 0x00, 0xff, 0xff, 0xf4, 0xe9, 0x16, 0x73, 0x91, 0x02,
	0x00, 0x00,
}
