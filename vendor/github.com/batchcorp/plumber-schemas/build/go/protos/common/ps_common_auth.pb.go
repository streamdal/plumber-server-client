// Code generated by protoc-gen-go. DO NOT EDIT.
// source: ps_common_auth.proto

package common

import (
	fmt "fmt"
	proto "github.com/golang/protobuf/proto"
	math "math"
)

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// This is a compile-time assertion to ensure that this generated file
// is compatible with the proto package it is being compiled against.
// A compilation error at this line likely means your copy of the
// proto package needs to be updated.
const _ = proto.ProtoPackageIsVersion3 // please upgrade the proto package

type Foreman struct {
	AuthToken            string   `protobuf:"bytes,1,opt,name=auth_token,json=authToken,proto3" json:"auth_token,omitempty"`
	PlumberClusterId     string   `protobuf:"bytes,2,opt,name=plumber_cluster_id,json=plumberClusterId,proto3" json:"plumber_cluster_id,omitempty"`
	TeamId               string   `protobuf:"bytes,3,opt,name=team_id,json=teamId,proto3" json:"team_id,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *Foreman) Reset()         { *m = Foreman{} }
func (m *Foreman) String() string { return proto.CompactTextString(m) }
func (*Foreman) ProtoMessage()    {}
func (*Foreman) Descriptor() ([]byte, []int) {
	return fileDescriptor_69536745fee4073a, []int{0}
}

func (m *Foreman) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Foreman.Unmarshal(m, b)
}
func (m *Foreman) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Foreman.Marshal(b, m, deterministic)
}
func (m *Foreman) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Foreman.Merge(m, src)
}
func (m *Foreman) XXX_Size() int {
	return xxx_messageInfo_Foreman.Size(m)
}
func (m *Foreman) XXX_DiscardUnknown() {
	xxx_messageInfo_Foreman.DiscardUnknown(m)
}

var xxx_messageInfo_Foreman proto.InternalMessageInfo

func (m *Foreman) GetAuthToken() string {
	if m != nil {
		return m.AuthToken
	}
	return ""
}

func (m *Foreman) GetPlumberClusterId() string {
	if m != nil {
		return m.PlumberClusterId
	}
	return ""
}

func (m *Foreman) GetTeamId() string {
	if m != nil {
		return m.TeamId
	}
	return ""
}

type Auth struct {
	Token string `protobuf:"bytes,1,opt,name=token,proto3" json:"token,omitempty"`
	// Filled out by batch. Not intended to be used by plumber.
	XForeman             *Foreman `protobuf:"bytes,2,opt,name=_foreman,json=Foreman,proto3" json:"_foreman,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *Auth) Reset()         { *m = Auth{} }
func (m *Auth) String() string { return proto.CompactTextString(m) }
func (*Auth) ProtoMessage()    {}
func (*Auth) Descriptor() ([]byte, []int) {
	return fileDescriptor_69536745fee4073a, []int{1}
}

func (m *Auth) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Auth.Unmarshal(m, b)
}
func (m *Auth) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Auth.Marshal(b, m, deterministic)
}
func (m *Auth) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Auth.Merge(m, src)
}
func (m *Auth) XXX_Size() int {
	return xxx_messageInfo_Auth.Size(m)
}
func (m *Auth) XXX_DiscardUnknown() {
	xxx_messageInfo_Auth.DiscardUnknown(m)
}

var xxx_messageInfo_Auth proto.InternalMessageInfo

func (m *Auth) GetToken() string {
	if m != nil {
		return m.Token
	}
	return ""
}

func (m *Auth) GetXForeman() *Foreman {
	if m != nil {
		return m.XForeman
	}
	return nil
}

func init() {
	proto.RegisterType((*Foreman)(nil), "protos.common.Foreman")
	proto.RegisterType((*Auth)(nil), "protos.common.Auth")
}

func init() { proto.RegisterFile("ps_common_auth.proto", fileDescriptor_69536745fee4073a) }

var fileDescriptor_69536745fee4073a = []byte{
	// 228 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x54, 0x8e, 0xcf, 0x4b, 0x03, 0x31,
	0x10, 0x85, 0xa9, 0x3f, 0x5a, 0x3b, 0x22, 0x48, 0x28, 0xda, 0x8b, 0x20, 0x3d, 0x79, 0xd0, 0x0d,
	0xea, 0x51, 0x3c, 0xa8, 0x20, 0xf4, 0x24, 0x14, 0x4f, 0x5e, 0x42, 0x92, 0x8d, 0x9b, 0xc5, 0xcd,
	0x4e, 0x48, 0x26, 0xff, 0xbf, 0x24, 0x59, 0xc1, 0x9e, 0xc2, 0x9b, 0xf7, 0xf1, 0xf2, 0xc1, 0xca,
	0x47, 0xa1, 0xd1, 0x39, 0x1c, 0x85, 0x4c, 0x64, 0x1b, 0x1f, 0x90, 0x90, 0x9d, 0x95, 0x27, 0x36,
	0xb5, 0xd9, 0x20, 0x2c, 0xde, 0x31, 0x18, 0x27, 0x47, 0x76, 0x05, 0x90, 0x39, 0x41, 0xf8, 0x63,
	0xc6, 0xf5, 0xec, 0x7a, 0x76, 0xb3, 0xdc, 0x2d, 0xf3, 0xe5, 0x33, 0x1f, 0xd8, 0x2d, 0x30, 0x3f,
	0x24, 0xa7, 0x4c, 0x10, 0x7a, 0x48, 0x91, 0x4c, 0x10, 0x7d, 0xbb, 0x3e, 0x28, 0xd8, 0xf9, 0xd4,
	0xbc, 0xd5, 0x62, 0xdb, 0xb2, 0x4b, 0x58, 0x90, 0x91, 0x2e, 0x23, 0x87, 0x05, 0x99, 0xe7, 0xb8,
	0x6d, 0x37, 0x1f, 0x70, 0xf4, 0x92, 0xc8, 0xb2, 0x15, 0x1c, 0xff, 0xff, 0xa8, 0x06, 0x76, 0x0f,
	0x27, 0xe2, 0xbb, 0xfa, 0x94, 0xe9, 0xd3, 0x87, 0x8b, 0x66, 0x4f, 0xb8, 0x99, 0x6c, 0x77, 0x7f,
	0xda, 0xaf, 0xcf, 0x5f, 0x4f, 0x5d, 0x4f, 0x36, 0xa9, 0x4c, 0x70, 0x25, 0x49, 0x5b, 0x8d, 0xc1,
	0xf3, 0x49, 0xe9, 0x2e, 0x6a, 0x6b, 0x9c, 0x8c, 0x5c, 0xa5, 0x7e, 0x68, 0x79, 0x87, 0xbc, 0xee,
	0xf1, 0xba, 0xa7, 0xe6, 0x25, 0x3e, 0xfe, 0x06, 0x00, 0x00, 0xff, 0xff, 0x87, 0x00, 0x48, 0x8d,
	0x2e, 0x01, 0x00, 0x00,
}
