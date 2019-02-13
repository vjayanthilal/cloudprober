// Code generated by protoc-gen-go. DO NOT EDIT.
// source: github.com/google/cloudprober/surfacers/file/proto/config.proto

package proto

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

type SurfacerConf struct {
	// Where to write the results. If left unset, file surfacer writes to the
	// standard output.
	FilePath *string `protobuf:"bytes,1,opt,name=file_path,json=filePath" json:"file_path,omitempty"`
	Prefix   *string `protobuf:"bytes,2,opt,name=prefix,def=cloudprober" json:"prefix,omitempty"`
	// Compress data before writing to the file.
	CompressionEnabled   *bool    `protobuf:"varint,3,opt,name=compression_enabled,json=compressionEnabled,def=0" json:"compression_enabled,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *SurfacerConf) Reset()         { *m = SurfacerConf{} }
func (m *SurfacerConf) String() string { return proto.CompactTextString(m) }
func (*SurfacerConf) ProtoMessage()    {}
func (*SurfacerConf) Descriptor() ([]byte, []int) {
	return fileDescriptor_config_da73a796d09f0ba3, []int{0}
}
func (m *SurfacerConf) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_SurfacerConf.Unmarshal(m, b)
}
func (m *SurfacerConf) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_SurfacerConf.Marshal(b, m, deterministic)
}
func (dst *SurfacerConf) XXX_Merge(src proto.Message) {
	xxx_messageInfo_SurfacerConf.Merge(dst, src)
}
func (m *SurfacerConf) XXX_Size() int {
	return xxx_messageInfo_SurfacerConf.Size(m)
}
func (m *SurfacerConf) XXX_DiscardUnknown() {
	xxx_messageInfo_SurfacerConf.DiscardUnknown(m)
}

var xxx_messageInfo_SurfacerConf proto.InternalMessageInfo

const Default_SurfacerConf_Prefix string = "cloudprober"
const Default_SurfacerConf_CompressionEnabled bool = false

func (m *SurfacerConf) GetFilePath() string {
	if m != nil && m.FilePath != nil {
		return *m.FilePath
	}
	return ""
}

func (m *SurfacerConf) GetPrefix() string {
	if m != nil && m.Prefix != nil {
		return *m.Prefix
	}
	return Default_SurfacerConf_Prefix
}

func (m *SurfacerConf) GetCompressionEnabled() bool {
	if m != nil && m.CompressionEnabled != nil {
		return *m.CompressionEnabled
	}
	return Default_SurfacerConf_CompressionEnabled
}

func init() {
	proto.RegisterType((*SurfacerConf)(nil), "cloudprober.surfacer.file.SurfacerConf")
}

func init() {
	proto.RegisterFile("github.com/google/cloudprober/surfacers/file/proto/config.proto", fileDescriptor_config_da73a796d09f0ba3)
}

var fileDescriptor_config_da73a796d09f0ba3 = []byte{
	// 193 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x4c, 0xcc, 0xb1, 0x4e, 0xc5, 0x20,
	0x14, 0xc6, 0xf1, 0xa0, 0xd1, 0xb4, 0xe8, 0x84, 0x4b, 0x8d, 0x4b, 0xa3, 0x4b, 0x27, 0xd8, 0x1c,
	0xba, 0x38, 0x18, 0x77, 0x53, 0x1f, 0xa0, 0xa1, 0xf4, 0x40, 0x49, 0x28, 0x87, 0x00, 0x4d, 0x7c,
	0x84, 0xfb, 0xd8, 0x37, 0xdc, 0xdb, 0x9b, 0x74, 0x3c, 0xdf, 0xff, 0x97, 0x43, 0xbf, 0x8c, 0xcd,
	0xcb, 0x36, 0x71, 0x85, 0xab, 0x30, 0x88, 0xc6, 0x81, 0x50, 0x0e, 0xb7, 0x39, 0x44, 0x9c, 0x20,
	0x8a, 0xb4, 0x45, 0x2d, 0x15, 0xc4, 0x24, 0xb4, 0x75, 0x20, 0x42, 0xc4, 0x8c, 0x42, 0xa1, 0xd7,
	0xd6, 0xf0, 0xcb, 0xc1, 0x5e, 0x0f, 0x9c, 0xdf, 0x38, 0x2f, 0xfa, 0xfd, 0x44, 0xe8, 0xf3, 0xdf,
	0xbe, 0x7c, 0xa3, 0xd7, 0xec, 0x8d, 0xd6, 0x25, 0x8c, 0x41, 0xe6, 0xa5, 0x21, 0x2d, 0xe9, 0xea,
	0xa1, 0x2a, 0xc3, 0xaf, 0xcc, 0x0b, 0xfb, 0xa0, 0x8f, 0x21, 0x82, 0xb6, 0xff, 0xcd, 0x5d, 0x29,
	0xfd, 0xd3, 0xe1, 0xf1, 0xb0, 0x27, 0xf6, 0x49, 0x5f, 0x14, 0xae, 0x21, 0x42, 0x4a, 0x16, 0xfd,
	0x08, 0x5e, 0x4e, 0x0e, 0xe6, 0xe6, 0xbe, 0x25, 0x5d, 0xd5, 0x3f, 0x68, 0xe9, 0x12, 0x0c, 0xec,
	0x20, 0x7e, 0xae, 0xe0, 0x1c, 0x00, 0x00, 0xff, 0xff, 0xef, 0xff, 0x1c, 0xe0, 0xe7, 0x00, 0x00,
	0x00,
}