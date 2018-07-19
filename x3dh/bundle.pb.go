// Code generated by protoc-gen-go. DO NOT EDIT.
// source: x3dh/bundle.proto

/*
Package x3dh is a generated protocol buffer package.

It is generated from these files:
	x3dh/bundle.proto

It has these top-level messages:
	Bundle
*/
package x3dh

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

type Bundle struct {
	Identity     []byte `protobuf:"bytes,1,opt,name=identity,proto3" json:"identity,omitempty"`
	SignedPreKey []byte `protobuf:"bytes,2,opt,name=signed_pre_key,json=signedPreKey,proto3" json:"signed_pre_key,omitempty"`
	Signature    []byte `protobuf:"bytes,3,opt,name=signature,proto3" json:"signature,omitempty"`
}

func (m *Bundle) Reset()                    { *m = Bundle{} }
func (m *Bundle) String() string            { return proto.CompactTextString(m) }
func (*Bundle) ProtoMessage()               {}
func (*Bundle) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{0} }

func (m *Bundle) GetIdentity() []byte {
	if m != nil {
		return m.Identity
	}
	return nil
}

func (m *Bundle) GetSignedPreKey() []byte {
	if m != nil {
		return m.SignedPreKey
	}
	return nil
}

func (m *Bundle) GetSignature() []byte {
	if m != nil {
		return m.Signature
	}
	return nil
}

func init() {
	proto.RegisterType((*Bundle)(nil), "x3dh.Bundle")
}

func init() { proto.RegisterFile("x3dh/bundle.proto", fileDescriptor0) }

var fileDescriptor0 = []byte{
	// 127 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xe2, 0x12, 0xac, 0x30, 0x4e, 0xc9,
	0xd0, 0x4f, 0x2a, 0xcd, 0x4b, 0xc9, 0x49, 0xd5, 0x2b, 0x28, 0xca, 0x2f, 0xc9, 0x17, 0x62, 0x01,
	0x09, 0x29, 0x65, 0x70, 0xb1, 0x39, 0x81, 0x45, 0x85, 0xa4, 0xb8, 0x38, 0x32, 0x53, 0x52, 0xf3,
	0x4a, 0x32, 0x4b, 0x2a, 0x25, 0x18, 0x15, 0x18, 0x35, 0x78, 0x82, 0xe0, 0x7c, 0x21, 0x15, 0x2e,
	0xbe, 0xe2, 0xcc, 0xf4, 0xbc, 0xd4, 0x94, 0xf8, 0x82, 0xa2, 0xd4, 0xf8, 0xec, 0xd4, 0x4a, 0x09,
	0x26, 0xb0, 0x0a, 0x1e, 0x88, 0x68, 0x40, 0x51, 0xaa, 0x77, 0x6a, 0xa5, 0x90, 0x0c, 0x17, 0x27,
	0x88, 0x9f, 0x58, 0x52, 0x5a, 0x94, 0x2a, 0xc1, 0x0c, 0x56, 0x80, 0x10, 0x48, 0x62, 0x03, 0x5b,
	0x6b, 0x0c, 0x08, 0x00, 0x00, 0xff, 0xff, 0xba, 0x31, 0xf8, 0x55, 0x8b, 0x00, 0x00, 0x00,
}