// Code generated by protoc-gen-go. DO NOT EDIT.
// source: chat.proto

/*
Package chat is a generated protocol buffer package.

It is generated from these files:
	chat.proto

It has these top-level messages:
	Bundle
	BundleContainer
	OneToOnePayload
	ContactUpdatePayload
	DirectMessageRPC
	DirectMessageProtocol
	RPCMessage
	ProtocolMessage
*/
package chat

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

type BundleContainer struct {
	Bundle              *Bundle `protobuf:"bytes,1,opt,name=bundle" json:"bundle,omitempty"`
	PrivateSignedPreKey []byte  `protobuf:"bytes,2,opt,name=private_signed_pre_key,json=privateSignedPreKey,proto3" json:"private_signed_pre_key,omitempty"`
}

func (m *BundleContainer) Reset()                    { *m = BundleContainer{} }
func (m *BundleContainer) String() string            { return proto.CompactTextString(m) }
func (*BundleContainer) ProtoMessage()               {}
func (*BundleContainer) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{1} }

func (m *BundleContainer) GetBundle() *Bundle {
	if m != nil {
		return m.Bundle
	}
	return nil
}

func (m *BundleContainer) GetPrivateSignedPreKey() []byte {
	if m != nil {
		return m.PrivateSignedPreKey
	}
	return nil
}

// What is sent through the wire
type OneToOnePayload struct {
	Content     string `protobuf:"bytes,1,opt,name=content" json:"content,omitempty"`
	ContentType string `protobuf:"bytes,2,opt,name=content_type,json=contentType" json:"content_type,omitempty"`
	MessageType string `protobuf:"bytes,3,opt,name=message_type,json=messageType" json:"message_type,omitempty"`
	ClockValue  int64  `protobuf:"varint,4,opt,name=clock_value,json=clockValue" json:"clock_value,omitempty"`
}

func (m *OneToOnePayload) Reset()                    { *m = OneToOnePayload{} }
func (m *OneToOnePayload) String() string            { return proto.CompactTextString(m) }
func (*OneToOnePayload) ProtoMessage()               {}
func (*OneToOnePayload) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{2} }

func (m *OneToOnePayload) GetContent() string {
	if m != nil {
		return m.Content
	}
	return ""
}

func (m *OneToOnePayload) GetContentType() string {
	if m != nil {
		return m.ContentType
	}
	return ""
}

func (m *OneToOnePayload) GetMessageType() string {
	if m != nil {
		return m.MessageType
	}
	return ""
}

func (m *OneToOnePayload) GetClockValue() int64 {
	if m != nil {
		return m.ClockValue
	}
	return 0
}

type ContactUpdatePayload struct {
	Name         string `protobuf:"bytes,1,opt,name=name" json:"name,omitempty"`
	ProfileImage string `protobuf:"bytes,2,opt,name=profile_image,json=profileImage" json:"profile_image,omitempty"`
	Address      string `protobuf:"bytes,3,opt,name=address" json:"address,omitempty"`
	FcmToken     string `protobuf:"bytes,4,opt,name=fcm_token,json=fcmToken" json:"fcm_token,omitempty"`
}

func (m *ContactUpdatePayload) Reset()                    { *m = ContactUpdatePayload{} }
func (m *ContactUpdatePayload) String() string            { return proto.CompactTextString(m) }
func (*ContactUpdatePayload) ProtoMessage()               {}
func (*ContactUpdatePayload) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{3} }

func (m *ContactUpdatePayload) GetName() string {
	if m != nil {
		return m.Name
	}
	return ""
}

func (m *ContactUpdatePayload) GetProfileImage() string {
	if m != nil {
		return m.ProfileImage
	}
	return ""
}

func (m *ContactUpdatePayload) GetAddress() string {
	if m != nil {
		return m.Address
	}
	return ""
}

func (m *ContactUpdatePayload) GetFcmToken() string {
	if m != nil {
		return m.FcmToken
	}
	return ""
}

// Sent to the RPC
type DirectMessageRPC struct {
	Dst []byte `protobuf:"bytes,1,opt,name=dst,proto3" json:"dst,omitempty"`
	// Types that are valid to be assigned to MessageType:
	//	*DirectMessageRPC_OneToOnePayload
	//	*DirectMessageRPC_ContactUpdatePayload
	MessageType isDirectMessageRPC_MessageType `protobuf_oneof:"message_type"`
	// Eventually removed if we persist bundles in status-go
	BundleId string `protobuf:"bytes,1000,opt,name=bundle_id,json=bundleId" json:"bundle_id,omitempty"`
	SymKeyId string `protobuf:"bytes,1001,opt,name=sym_key_id,json=symKeyId" json:"sym_key_id,omitempty"`
}

func (m *DirectMessageRPC) Reset()                    { *m = DirectMessageRPC{} }
func (m *DirectMessageRPC) String() string            { return proto.CompactTextString(m) }
func (*DirectMessageRPC) ProtoMessage()               {}
func (*DirectMessageRPC) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{4} }

type isDirectMessageRPC_MessageType interface{ isDirectMessageRPC_MessageType() }

type DirectMessageRPC_OneToOnePayload struct {
	OneToOnePayload *OneToOnePayload `protobuf:"bytes,101,opt,name=one_to_one_payload,json=oneToOnePayload,oneof"`
}
type DirectMessageRPC_ContactUpdatePayload struct {
	ContactUpdatePayload *ContactUpdatePayload `protobuf:"bytes,102,opt,name=contact_update_payload,json=contactUpdatePayload,oneof"`
}

func (*DirectMessageRPC_OneToOnePayload) isDirectMessageRPC_MessageType()      {}
func (*DirectMessageRPC_ContactUpdatePayload) isDirectMessageRPC_MessageType() {}

func (m *DirectMessageRPC) GetMessageType() isDirectMessageRPC_MessageType {
	if m != nil {
		return m.MessageType
	}
	return nil
}

func (m *DirectMessageRPC) GetDst() []byte {
	if m != nil {
		return m.Dst
	}
	return nil
}

func (m *DirectMessageRPC) GetOneToOnePayload() *OneToOnePayload {
	if x, ok := m.GetMessageType().(*DirectMessageRPC_OneToOnePayload); ok {
		return x.OneToOnePayload
	}
	return nil
}

func (m *DirectMessageRPC) GetContactUpdatePayload() *ContactUpdatePayload {
	if x, ok := m.GetMessageType().(*DirectMessageRPC_ContactUpdatePayload); ok {
		return x.ContactUpdatePayload
	}
	return nil
}

func (m *DirectMessageRPC) GetBundleId() string {
	if m != nil {
		return m.BundleId
	}
	return ""
}

func (m *DirectMessageRPC) GetSymKeyId() string {
	if m != nil {
		return m.SymKeyId
	}
	return ""
}

// XXX_OneofFuncs is for the internal use of the proto package.
func (*DirectMessageRPC) XXX_OneofFuncs() (func(msg proto.Message, b *proto.Buffer) error, func(msg proto.Message, tag, wire int, b *proto.Buffer) (bool, error), func(msg proto.Message) (n int), []interface{}) {
	return _DirectMessageRPC_OneofMarshaler, _DirectMessageRPC_OneofUnmarshaler, _DirectMessageRPC_OneofSizer, []interface{}{
		(*DirectMessageRPC_OneToOnePayload)(nil),
		(*DirectMessageRPC_ContactUpdatePayload)(nil),
	}
}

func _DirectMessageRPC_OneofMarshaler(msg proto.Message, b *proto.Buffer) error {
	m := msg.(*DirectMessageRPC)
	// message_type
	switch x := m.MessageType.(type) {
	case *DirectMessageRPC_OneToOnePayload:
		b.EncodeVarint(101<<3 | proto.WireBytes)
		if err := b.EncodeMessage(x.OneToOnePayload); err != nil {
			return err
		}
	case *DirectMessageRPC_ContactUpdatePayload:
		b.EncodeVarint(102<<3 | proto.WireBytes)
		if err := b.EncodeMessage(x.ContactUpdatePayload); err != nil {
			return err
		}
	case nil:
	default:
		return fmt.Errorf("DirectMessageRPC.MessageType has unexpected type %T", x)
	}
	return nil
}

func _DirectMessageRPC_OneofUnmarshaler(msg proto.Message, tag, wire int, b *proto.Buffer) (bool, error) {
	m := msg.(*DirectMessageRPC)
	switch tag {
	case 101: // message_type.one_to_one_payload
		if wire != proto.WireBytes {
			return true, proto.ErrInternalBadWireType
		}
		msg := new(OneToOnePayload)
		err := b.DecodeMessage(msg)
		m.MessageType = &DirectMessageRPC_OneToOnePayload{msg}
		return true, err
	case 102: // message_type.contact_update_payload
		if wire != proto.WireBytes {
			return true, proto.ErrInternalBadWireType
		}
		msg := new(ContactUpdatePayload)
		err := b.DecodeMessage(msg)
		m.MessageType = &DirectMessageRPC_ContactUpdatePayload{msg}
		return true, err
	default:
		return false, nil
	}
}

func _DirectMessageRPC_OneofSizer(msg proto.Message) (n int) {
	m := msg.(*DirectMessageRPC)
	// message_type
	switch x := m.MessageType.(type) {
	case *DirectMessageRPC_OneToOnePayload:
		s := proto.Size(x.OneToOnePayload)
		n += proto.SizeVarint(101<<3 | proto.WireBytes)
		n += proto.SizeVarint(uint64(s))
		n += s
	case *DirectMessageRPC_ContactUpdatePayload:
		s := proto.Size(x.ContactUpdatePayload)
		n += proto.SizeVarint(102<<3 | proto.WireBytes)
		n += proto.SizeVarint(uint64(s))
		n += s
	case nil:
	default:
		panic(fmt.Sprintf("proto: unexpected type %T in oneof", x))
	}
	return n
}

// Sent among peers
type DirectMessageProtocol struct {
	// Which ephemeral key has been used
	EphemeralKey []byte `protobuf:"bytes,1,opt,name=ephemeral_key,json=ephemeralKey,proto3" json:"ephemeral_key,omitempty"`
	// Bundle pre key in case of x3dh
	BundlePreKey []byte `protobuf:"bytes,2,opt,name=bundle_pre_key,json=bundlePreKey,proto3" json:"bundle_pre_key,omitempty"`
	// Types that are valid to be assigned to MessageType:
	//	*DirectMessageProtocol_OneToOnePayload
	//	*DirectMessageProtocol_ContactUpdatePayload
	MessageType isDirectMessageProtocol_MessageType `protobuf_oneof:"message_type"`
}

func (m *DirectMessageProtocol) Reset()                    { *m = DirectMessageProtocol{} }
func (m *DirectMessageProtocol) String() string            { return proto.CompactTextString(m) }
func (*DirectMessageProtocol) ProtoMessage()               {}
func (*DirectMessageProtocol) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{5} }

type isDirectMessageProtocol_MessageType interface{ isDirectMessageProtocol_MessageType() }

type DirectMessageProtocol_OneToOnePayload struct {
	OneToOnePayload *OneToOnePayload `protobuf:"bytes,101,opt,name=one_to_one_payload,json=oneToOnePayload,oneof"`
}
type DirectMessageProtocol_ContactUpdatePayload struct {
	ContactUpdatePayload *ContactUpdatePayload `protobuf:"bytes,102,opt,name=contact_update_payload,json=contactUpdatePayload,oneof"`
}

func (*DirectMessageProtocol_OneToOnePayload) isDirectMessageProtocol_MessageType()      {}
func (*DirectMessageProtocol_ContactUpdatePayload) isDirectMessageProtocol_MessageType() {}

func (m *DirectMessageProtocol) GetMessageType() isDirectMessageProtocol_MessageType {
	if m != nil {
		return m.MessageType
	}
	return nil
}

func (m *DirectMessageProtocol) GetEphemeralKey() []byte {
	if m != nil {
		return m.EphemeralKey
	}
	return nil
}

func (m *DirectMessageProtocol) GetBundlePreKey() []byte {
	if m != nil {
		return m.BundlePreKey
	}
	return nil
}

func (m *DirectMessageProtocol) GetOneToOnePayload() *OneToOnePayload {
	if x, ok := m.GetMessageType().(*DirectMessageProtocol_OneToOnePayload); ok {
		return x.OneToOnePayload
	}
	return nil
}

func (m *DirectMessageProtocol) GetContactUpdatePayload() *ContactUpdatePayload {
	if x, ok := m.GetMessageType().(*DirectMessageProtocol_ContactUpdatePayload); ok {
		return x.ContactUpdatePayload
	}
	return nil
}

// XXX_OneofFuncs is for the internal use of the proto package.
func (*DirectMessageProtocol) XXX_OneofFuncs() (func(msg proto.Message, b *proto.Buffer) error, func(msg proto.Message, tag, wire int, b *proto.Buffer) (bool, error), func(msg proto.Message) (n int), []interface{}) {
	return _DirectMessageProtocol_OneofMarshaler, _DirectMessageProtocol_OneofUnmarshaler, _DirectMessageProtocol_OneofSizer, []interface{}{
		(*DirectMessageProtocol_OneToOnePayload)(nil),
		(*DirectMessageProtocol_ContactUpdatePayload)(nil),
	}
}

func _DirectMessageProtocol_OneofMarshaler(msg proto.Message, b *proto.Buffer) error {
	m := msg.(*DirectMessageProtocol)
	// message_type
	switch x := m.MessageType.(type) {
	case *DirectMessageProtocol_OneToOnePayload:
		b.EncodeVarint(101<<3 | proto.WireBytes)
		if err := b.EncodeMessage(x.OneToOnePayload); err != nil {
			return err
		}
	case *DirectMessageProtocol_ContactUpdatePayload:
		b.EncodeVarint(102<<3 | proto.WireBytes)
		if err := b.EncodeMessage(x.ContactUpdatePayload); err != nil {
			return err
		}
	case nil:
	default:
		return fmt.Errorf("DirectMessageProtocol.MessageType has unexpected type %T", x)
	}
	return nil
}

func _DirectMessageProtocol_OneofUnmarshaler(msg proto.Message, tag, wire int, b *proto.Buffer) (bool, error) {
	m := msg.(*DirectMessageProtocol)
	switch tag {
	case 101: // message_type.one_to_one_payload
		if wire != proto.WireBytes {
			return true, proto.ErrInternalBadWireType
		}
		msg := new(OneToOnePayload)
		err := b.DecodeMessage(msg)
		m.MessageType = &DirectMessageProtocol_OneToOnePayload{msg}
		return true, err
	case 102: // message_type.contact_update_payload
		if wire != proto.WireBytes {
			return true, proto.ErrInternalBadWireType
		}
		msg := new(ContactUpdatePayload)
		err := b.DecodeMessage(msg)
		m.MessageType = &DirectMessageProtocol_ContactUpdatePayload{msg}
		return true, err
	default:
		return false, nil
	}
}

func _DirectMessageProtocol_OneofSizer(msg proto.Message) (n int) {
	m := msg.(*DirectMessageProtocol)
	// message_type
	switch x := m.MessageType.(type) {
	case *DirectMessageProtocol_OneToOnePayload:
		s := proto.Size(x.OneToOnePayload)
		n += proto.SizeVarint(101<<3 | proto.WireBytes)
		n += proto.SizeVarint(uint64(s))
		n += s
	case *DirectMessageProtocol_ContactUpdatePayload:
		s := proto.Size(x.ContactUpdatePayload)
		n += proto.SizeVarint(102<<3 | proto.WireBytes)
		n += proto.SizeVarint(uint64(s))
		n += s
	case nil:
	default:
		panic(fmt.Sprintf("proto: unexpected type %T in oneof", x))
	}
	return n
}

// Message sent to RPC
type RPCMessage struct {
	// Types that are valid to be assigned to MessageType:
	//	*RPCMessage_DirectMessage
	MessageType isRPCMessage_MessageType `protobuf_oneof:"message_type"`
}

func (m *RPCMessage) Reset()                    { *m = RPCMessage{} }
func (m *RPCMessage) String() string            { return proto.CompactTextString(m) }
func (*RPCMessage) ProtoMessage()               {}
func (*RPCMessage) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{6} }

type isRPCMessage_MessageType interface{ isRPCMessage_MessageType() }

type RPCMessage_DirectMessage struct {
	DirectMessage *DirectMessageRPC `protobuf:"bytes,1,opt,name=direct_message,json=directMessage,oneof"`
}

func (*RPCMessage_DirectMessage) isRPCMessage_MessageType() {}

func (m *RPCMessage) GetMessageType() isRPCMessage_MessageType {
	if m != nil {
		return m.MessageType
	}
	return nil
}

func (m *RPCMessage) GetDirectMessage() *DirectMessageRPC {
	if x, ok := m.GetMessageType().(*RPCMessage_DirectMessage); ok {
		return x.DirectMessage
	}
	return nil
}

// XXX_OneofFuncs is for the internal use of the proto package.
func (*RPCMessage) XXX_OneofFuncs() (func(msg proto.Message, b *proto.Buffer) error, func(msg proto.Message, tag, wire int, b *proto.Buffer) (bool, error), func(msg proto.Message) (n int), []interface{}) {
	return _RPCMessage_OneofMarshaler, _RPCMessage_OneofUnmarshaler, _RPCMessage_OneofSizer, []interface{}{
		(*RPCMessage_DirectMessage)(nil),
	}
}

func _RPCMessage_OneofMarshaler(msg proto.Message, b *proto.Buffer) error {
	m := msg.(*RPCMessage)
	// message_type
	switch x := m.MessageType.(type) {
	case *RPCMessage_DirectMessage:
		b.EncodeVarint(1<<3 | proto.WireBytes)
		if err := b.EncodeMessage(x.DirectMessage); err != nil {
			return err
		}
	case nil:
	default:
		return fmt.Errorf("RPCMessage.MessageType has unexpected type %T", x)
	}
	return nil
}

func _RPCMessage_OneofUnmarshaler(msg proto.Message, tag, wire int, b *proto.Buffer) (bool, error) {
	m := msg.(*RPCMessage)
	switch tag {
	case 1: // message_type.direct_message
		if wire != proto.WireBytes {
			return true, proto.ErrInternalBadWireType
		}
		msg := new(DirectMessageRPC)
		err := b.DecodeMessage(msg)
		m.MessageType = &RPCMessage_DirectMessage{msg}
		return true, err
	default:
		return false, nil
	}
}

func _RPCMessage_OneofSizer(msg proto.Message) (n int) {
	m := msg.(*RPCMessage)
	// message_type
	switch x := m.MessageType.(type) {
	case *RPCMessage_DirectMessage:
		s := proto.Size(x.DirectMessage)
		n += proto.SizeVarint(1<<3 | proto.WireBytes)
		n += proto.SizeVarint(uint64(s))
		n += s
	case nil:
	default:
		panic(fmt.Sprintf("proto: unexpected type %T in oneof", x))
	}
	return n
}

// Message sent among peers
type ProtocolMessage struct {
	// An optional bundle is exchanged with each message
	Bundle *Bundle `protobuf:"bytes,1,opt,name=bundle" json:"bundle,omitempty"`
	// Types that are valid to be assigned to MessageType:
	//	*ProtocolMessage_DirectMessage
	MessageType isProtocolMessage_MessageType `protobuf_oneof:"message_type"`
}

func (m *ProtocolMessage) Reset()                    { *m = ProtocolMessage{} }
func (m *ProtocolMessage) String() string            { return proto.CompactTextString(m) }
func (*ProtocolMessage) ProtoMessage()               {}
func (*ProtocolMessage) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{7} }

type isProtocolMessage_MessageType interface{ isProtocolMessage_MessageType() }

type ProtocolMessage_DirectMessage struct {
	DirectMessage *DirectMessageProtocol `protobuf:"bytes,101,opt,name=direct_message,json=directMessage,oneof"`
}

func (*ProtocolMessage_DirectMessage) isProtocolMessage_MessageType() {}

func (m *ProtocolMessage) GetMessageType() isProtocolMessage_MessageType {
	if m != nil {
		return m.MessageType
	}
	return nil
}

func (m *ProtocolMessage) GetBundle() *Bundle {
	if m != nil {
		return m.Bundle
	}
	return nil
}

func (m *ProtocolMessage) GetDirectMessage() *DirectMessageProtocol {
	if x, ok := m.GetMessageType().(*ProtocolMessage_DirectMessage); ok {
		return x.DirectMessage
	}
	return nil
}

// XXX_OneofFuncs is for the internal use of the proto package.
func (*ProtocolMessage) XXX_OneofFuncs() (func(msg proto.Message, b *proto.Buffer) error, func(msg proto.Message, tag, wire int, b *proto.Buffer) (bool, error), func(msg proto.Message) (n int), []interface{}) {
	return _ProtocolMessage_OneofMarshaler, _ProtocolMessage_OneofUnmarshaler, _ProtocolMessage_OneofSizer, []interface{}{
		(*ProtocolMessage_DirectMessage)(nil),
	}
}

func _ProtocolMessage_OneofMarshaler(msg proto.Message, b *proto.Buffer) error {
	m := msg.(*ProtocolMessage)
	// message_type
	switch x := m.MessageType.(type) {
	case *ProtocolMessage_DirectMessage:
		b.EncodeVarint(101<<3 | proto.WireBytes)
		if err := b.EncodeMessage(x.DirectMessage); err != nil {
			return err
		}
	case nil:
	default:
		return fmt.Errorf("ProtocolMessage.MessageType has unexpected type %T", x)
	}
	return nil
}

func _ProtocolMessage_OneofUnmarshaler(msg proto.Message, tag, wire int, b *proto.Buffer) (bool, error) {
	m := msg.(*ProtocolMessage)
	switch tag {
	case 101: // message_type.direct_message
		if wire != proto.WireBytes {
			return true, proto.ErrInternalBadWireType
		}
		msg := new(DirectMessageProtocol)
		err := b.DecodeMessage(msg)
		m.MessageType = &ProtocolMessage_DirectMessage{msg}
		return true, err
	default:
		return false, nil
	}
}

func _ProtocolMessage_OneofSizer(msg proto.Message) (n int) {
	m := msg.(*ProtocolMessage)
	// message_type
	switch x := m.MessageType.(type) {
	case *ProtocolMessage_DirectMessage:
		s := proto.Size(x.DirectMessage)
		n += proto.SizeVarint(101<<3 | proto.WireBytes)
		n += proto.SizeVarint(uint64(s))
		n += s
	case nil:
	default:
		panic(fmt.Sprintf("proto: unexpected type %T in oneof", x))
	}
	return n
}

func init() {
	proto.RegisterType((*Bundle)(nil), "chat.Bundle")
	proto.RegisterType((*BundleContainer)(nil), "chat.BundleContainer")
	proto.RegisterType((*OneToOnePayload)(nil), "chat.OneToOnePayload")
	proto.RegisterType((*ContactUpdatePayload)(nil), "chat.ContactUpdatePayload")
	proto.RegisterType((*DirectMessageRPC)(nil), "chat.DirectMessageRPC")
	proto.RegisterType((*DirectMessageProtocol)(nil), "chat.DirectMessageProtocol")
	proto.RegisterType((*RPCMessage)(nil), "chat.RPCMessage")
	proto.RegisterType((*ProtocolMessage)(nil), "chat.ProtocolMessage")
}

func init() { proto.RegisterFile("chat.proto", fileDescriptor0) }

var fileDescriptor0 = []byte{
	// 564 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xcc, 0x54, 0xcd, 0x4e, 0xdb, 0x40,
	0x10, 0xc6, 0x04, 0x01, 0x1e, 0x4c, 0x82, 0xb6, 0x80, 0x2c, 0xa0, 0x2a, 0x75, 0x39, 0x70, 0xe2,
	0x50, 0x1e, 0xa0, 0x12, 0x70, 0x20, 0x42, 0x15, 0xd1, 0x36, 0xed, 0xad, 0x5a, 0x2d, 0xde, 0x49,
	0x62, 0xc5, 0xde, 0xb5, 0xec, 0x0d, 0x92, 0x5f, 0xa0, 0x97, 0xb6, 0x97, 0x3e, 0x65, 0xfb, 0x16,
	0xd5, 0xfe, 0x38, 0x85, 0x24, 0x52, 0x7b, 0xec, 0xc9, 0x3b, 0xdf, 0x7e, 0x9a, 0xf9, 0xe6, 0x9b,
	0x59, 0x03, 0xa4, 0x13, 0xae, 0x2f, 0xca, 0x4a, 0x69, 0x45, 0x36, 0xcc, 0x39, 0x99, 0xc0, 0xe6,
	0xd5, 0x4c, 0x8a, 0x1c, 0xc9, 0x11, 0x6c, 0x67, 0x02, 0xa5, 0xce, 0x74, 0x13, 0x07, 0xa7, 0xc1,
	0x79, 0x44, 0xe7, 0x31, 0x39, 0x83, 0x6e, 0x9d, 0x8d, 0x25, 0x0a, 0x56, 0x56, 0xc8, 0xa6, 0xd8,
	0xc4, 0xeb, 0x96, 0x11, 0x39, 0x74, 0x50, 0xe1, 0x1d, 0x36, 0xe4, 0x04, 0x42, 0x13, 0x73, 0x3d,
	0xab, 0x30, 0xee, 0x58, 0xc2, 0x1f, 0x20, 0xc9, 0xa1, 0xe7, 0x2a, 0x5d, 0x2b, 0xa9, 0x79, 0x26,
	0xb1, 0x22, 0x67, 0xb0, 0xf9, 0x60, 0x21, 0x5b, 0x70, 0xe7, 0x6d, 0x74, 0x61, 0xf5, 0x39, 0x1a,
	0xf5, 0x77, 0xe4, 0x12, 0x0e, 0xcb, 0x2a, 0x7b, 0xe4, 0x1a, 0xd9, 0x4a, 0x11, 0x2f, 0xfc, 0xed,
	0x87, 0x27, 0x5a, 0x92, 0x1f, 0x01, 0xf4, 0xee, 0x25, 0x0e, 0xd5, 0xbd, 0xc4, 0x01, 0x6f, 0x72,
	0xc5, 0x05, 0x89, 0x61, 0x2b, 0x55, 0x52, 0xa3, 0xd4, 0xb6, 0x5e, 0x48, 0xdb, 0x90, 0xbc, 0x86,
	0xc8, 0x1f, 0x99, 0x6e, 0x4a, 0xb4, 0x89, 0x43, 0xba, 0xe3, 0xb1, 0x61, 0x53, 0xa2, 0xa1, 0x14,
	0x58, 0xd7, 0x7c, 0x8c, 0x8e, 0xd2, 0x71, 0x14, 0x8f, 0x59, 0xca, 0x2b, 0xd8, 0x49, 0x73, 0x95,
	0x4e, 0xd9, 0x23, 0xcf, 0x67, 0x18, 0x6f, 0x9c, 0x06, 0xe7, 0x1d, 0x0a, 0x16, 0xfa, 0x64, 0x90,
	0xe4, 0x4b, 0x00, 0xfb, 0xb6, 0xfb, 0x54, 0x7f, 0x2c, 0x05, 0xd7, 0x73, 0x65, 0x04, 0x36, 0x24,
	0x2f, 0xd0, 0xcb, 0xb2, 0x67, 0xf2, 0x06, 0x76, 0xcb, 0x4a, 0x8d, 0xb2, 0x1c, 0x59, 0x56, 0xf0,
	0x71, 0x2b, 0x2a, 0xf2, 0x60, 0xdf, 0x60, 0xa6, 0x25, 0x2e, 0x44, 0x85, 0x75, 0xed, 0x05, 0xb5,
	0x21, 0x39, 0x86, 0x70, 0x94, 0x16, 0x4c, 0xab, 0x29, 0x4a, 0x2b, 0x25, 0xa4, 0xdb, 0xa3, 0xb4,
	0x18, 0x9a, 0x38, 0xf9, 0xba, 0x0e, 0x7b, 0x37, 0x59, 0x85, 0xa9, 0x7e, 0xef, 0xf4, 0xd3, 0xc1,
	0x35, 0xd9, 0x83, 0x8e, 0xa8, 0xb5, 0x9f, 0xbd, 0x39, 0x92, 0x1b, 0x20, 0x4a, 0x22, 0xd3, 0x8a,
	0x99, 0x4f, 0xe9, 0xc4, 0xc6, 0x68, 0x67, 0x75, 0xe0, 0x66, 0xb5, 0xe0, 0xf1, 0xed, 0x1a, 0xed,
	0xa9, 0x05, 0xdb, 0x29, 0x1c, 0xa6, 0xae, 0x69, 0x36, 0xb3, 0x5d, 0xcf, 0x33, 0x8d, 0x6c, 0xa6,
	0x23, 0x97, 0x69, 0x95, 0x31, 0xb7, 0x6b, 0x74, 0x3f, 0x5d, 0x65, 0xd8, 0x09, 0x84, 0x6e, 0x3b,
	0x58, 0x26, 0xe2, 0x9f, 0x5b, 0xae, 0x3d, 0x87, 0xf4, 0x05, 0x79, 0x09, 0x50, 0x37, 0x85, 0x59,
	0x11, 0x73, 0xfd, 0xcb, 0x5f, 0xd7, 0x4d, 0x71, 0x87, 0x4d, 0x5f, 0x5c, 0x75, 0x9f, 0x8f, 0x32,
	0xf9, 0xb6, 0x0e, 0x07, 0xcf, 0xdc, 0x18, 0x98, 0x07, 0x92, 0xaa, 0xdc, 0xcc, 0x00, 0xcb, 0x09,
	0x16, 0x58, 0xf1, 0xdc, 0x6e, 0x9c, 0x33, 0x27, 0x9a, 0x83, 0x66, 0xed, 0xcf, 0xa0, 0xeb, 0xb5,
	0x2c, 0x3c, 0x0e, 0x87, 0xfa, 0xc7, 0xf1, 0xdf, 0x7a, 0xb9, 0x64, 0xc7, 0x67, 0x00, 0x3a, 0xb8,
	0xf6, 0x56, 0x90, 0x77, 0xd0, 0x15, 0xd6, 0x1b, 0xe6, 0x49, 0xfe, 0xad, 0x1e, 0xba, 0x4a, 0x8b,
	0x5b, 0x74, 0xbb, 0x46, 0x77, 0xc5, 0x53, 0x6c, 0x29, 0xfd, 0xf7, 0x00, 0x7a, 0xad, 0xc1, 0x6d,
	0x91, 0x7f, 0xfb, 0x11, 0xdc, 0x2c, 0x49, 0x71, 0xf6, 0x1d, 0xaf, 0x90, 0xd2, 0x56, 0xf8, 0xab,
	0x9e, 0x87, 0x4d, 0xfb, 0x3b, 0xbc, 0xfc, 0x1d, 0x00, 0x00, 0xff, 0xff, 0x07, 0xc1, 0xac, 0x31,
	0x1c, 0x05, 0x00, 0x00,
}
