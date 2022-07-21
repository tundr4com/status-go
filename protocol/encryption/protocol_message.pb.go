// Code generated by protoc-gen-go. DO NOT EDIT.
// source: protocol_message.proto

package encryption

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

type SignedPreKey struct {
	SignedPreKey         []byte   `protobuf:"bytes,1,opt,name=signed_pre_key,json=signedPreKey,proto3" json:"signed_pre_key,omitempty"`
	Version              uint32   `protobuf:"varint,2,opt,name=version,proto3" json:"version,omitempty"`
	ProtocolVersion      uint32   `protobuf:"varint,3,opt,name=protocol_version,json=protocolVersion,proto3" json:"protocol_version,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *SignedPreKey) Reset()         { *m = SignedPreKey{} }
func (m *SignedPreKey) String() string { return proto.CompactTextString(m) }
func (*SignedPreKey) ProtoMessage()    {}
func (*SignedPreKey) Descriptor() ([]byte, []int) {
	return fileDescriptor_4e37b52004a72e16, []int{0}
}

func (m *SignedPreKey) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_SignedPreKey.Unmarshal(m, b)
}
func (m *SignedPreKey) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_SignedPreKey.Marshal(b, m, deterministic)
}
func (m *SignedPreKey) XXX_Merge(src proto.Message) {
	xxx_messageInfo_SignedPreKey.Merge(m, src)
}
func (m *SignedPreKey) XXX_Size() int {
	return xxx_messageInfo_SignedPreKey.Size(m)
}
func (m *SignedPreKey) XXX_DiscardUnknown() {
	xxx_messageInfo_SignedPreKey.DiscardUnknown(m)
}

var xxx_messageInfo_SignedPreKey proto.InternalMessageInfo

func (m *SignedPreKey) GetSignedPreKey() []byte {
	if m != nil {
		return m.SignedPreKey
	}
	return nil
}

func (m *SignedPreKey) GetVersion() uint32 {
	if m != nil {
		return m.Version
	}
	return 0
}

func (m *SignedPreKey) GetProtocolVersion() uint32 {
	if m != nil {
		return m.ProtocolVersion
	}
	return 0
}

// X3DH prekey bundle
type Bundle struct {
	// Identity key
	Identity []byte `protobuf:"bytes,1,opt,name=identity,proto3" json:"identity,omitempty"`
	// Installation id
	SignedPreKeys map[string]*SignedPreKey `protobuf:"bytes,2,rep,name=signed_pre_keys,json=signedPreKeys,proto3" json:"signed_pre_keys,omitempty" protobuf_key:"bytes,1,opt,name=key,proto3" protobuf_val:"bytes,2,opt,name=value,proto3"`
	// Prekey signature
	Signature []byte `protobuf:"bytes,4,opt,name=signature,proto3" json:"signature,omitempty"`
	// When the bundle was created locally
	Timestamp            int64    `protobuf:"varint,5,opt,name=timestamp,proto3" json:"timestamp,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *Bundle) Reset()         { *m = Bundle{} }
func (m *Bundle) String() string { return proto.CompactTextString(m) }
func (*Bundle) ProtoMessage()    {}
func (*Bundle) Descriptor() ([]byte, []int) {
	return fileDescriptor_4e37b52004a72e16, []int{1}
}

func (m *Bundle) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Bundle.Unmarshal(m, b)
}
func (m *Bundle) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Bundle.Marshal(b, m, deterministic)
}
func (m *Bundle) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Bundle.Merge(m, src)
}
func (m *Bundle) XXX_Size() int {
	return xxx_messageInfo_Bundle.Size(m)
}
func (m *Bundle) XXX_DiscardUnknown() {
	xxx_messageInfo_Bundle.DiscardUnknown(m)
}

var xxx_messageInfo_Bundle proto.InternalMessageInfo

func (m *Bundle) GetIdentity() []byte {
	if m != nil {
		return m.Identity
	}
	return nil
}

func (m *Bundle) GetSignedPreKeys() map[string]*SignedPreKey {
	if m != nil {
		return m.SignedPreKeys
	}
	return nil
}

func (m *Bundle) GetSignature() []byte {
	if m != nil {
		return m.Signature
	}
	return nil
}

func (m *Bundle) GetTimestamp() int64 {
	if m != nil {
		return m.Timestamp
	}
	return 0
}

type BundleContainer struct {
	// X3DH prekey bundle
	Bundle *Bundle `protobuf:"bytes,1,opt,name=bundle,proto3" json:"bundle,omitempty"`
	// Private signed prekey
	PrivateSignedPreKey  []byte   `protobuf:"bytes,2,opt,name=private_signed_pre_key,json=privateSignedPreKey,proto3" json:"private_signed_pre_key,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *BundleContainer) Reset()         { *m = BundleContainer{} }
func (m *BundleContainer) String() string { return proto.CompactTextString(m) }
func (*BundleContainer) ProtoMessage()    {}
func (*BundleContainer) Descriptor() ([]byte, []int) {
	return fileDescriptor_4e37b52004a72e16, []int{2}
}

func (m *BundleContainer) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_BundleContainer.Unmarshal(m, b)
}
func (m *BundleContainer) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_BundleContainer.Marshal(b, m, deterministic)
}
func (m *BundleContainer) XXX_Merge(src proto.Message) {
	xxx_messageInfo_BundleContainer.Merge(m, src)
}
func (m *BundleContainer) XXX_Size() int {
	return xxx_messageInfo_BundleContainer.Size(m)
}
func (m *BundleContainer) XXX_DiscardUnknown() {
	xxx_messageInfo_BundleContainer.DiscardUnknown(m)
}

var xxx_messageInfo_BundleContainer proto.InternalMessageInfo

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

type DRHeader struct {
	// Current ratchet public key
	Key []byte `protobuf:"bytes,1,opt,name=key,proto3" json:"key,omitempty"`
	// Number of the message in the sending chain
	N uint32 `protobuf:"varint,2,opt,name=n,proto3" json:"n,omitempty"`
	// Length of the previous sending chain
	Pn uint32 `protobuf:"varint,3,opt,name=pn,proto3" json:"pn,omitempty"`
	// Bundle ID
	Id                   []byte   `protobuf:"bytes,4,opt,name=id,proto3" json:"id,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *DRHeader) Reset()         { *m = DRHeader{} }
func (m *DRHeader) String() string { return proto.CompactTextString(m) }
func (*DRHeader) ProtoMessage()    {}
func (*DRHeader) Descriptor() ([]byte, []int) {
	return fileDescriptor_4e37b52004a72e16, []int{3}
}

func (m *DRHeader) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_DRHeader.Unmarshal(m, b)
}
func (m *DRHeader) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_DRHeader.Marshal(b, m, deterministic)
}
func (m *DRHeader) XXX_Merge(src proto.Message) {
	xxx_messageInfo_DRHeader.Merge(m, src)
}
func (m *DRHeader) XXX_Size() int {
	return xxx_messageInfo_DRHeader.Size(m)
}
func (m *DRHeader) XXX_DiscardUnknown() {
	xxx_messageInfo_DRHeader.DiscardUnknown(m)
}

var xxx_messageInfo_DRHeader proto.InternalMessageInfo

func (m *DRHeader) GetKey() []byte {
	if m != nil {
		return m.Key
	}
	return nil
}

func (m *DRHeader) GetN() uint32 {
	if m != nil {
		return m.N
	}
	return 0
}

func (m *DRHeader) GetPn() uint32 {
	if m != nil {
		return m.Pn
	}
	return 0
}

func (m *DRHeader) GetId() []byte {
	if m != nil {
		return m.Id
	}
	return nil
}

type DHHeader struct {
	// Compressed ephemeral public key
	Key                  []byte   `protobuf:"bytes,1,opt,name=key,proto3" json:"key,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *DHHeader) Reset()         { *m = DHHeader{} }
func (m *DHHeader) String() string { return proto.CompactTextString(m) }
func (*DHHeader) ProtoMessage()    {}
func (*DHHeader) Descriptor() ([]byte, []int) {
	return fileDescriptor_4e37b52004a72e16, []int{4}
}

func (m *DHHeader) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_DHHeader.Unmarshal(m, b)
}
func (m *DHHeader) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_DHHeader.Marshal(b, m, deterministic)
}
func (m *DHHeader) XXX_Merge(src proto.Message) {
	xxx_messageInfo_DHHeader.Merge(m, src)
}
func (m *DHHeader) XXX_Size() int {
	return xxx_messageInfo_DHHeader.Size(m)
}
func (m *DHHeader) XXX_DiscardUnknown() {
	xxx_messageInfo_DHHeader.DiscardUnknown(m)
}

var xxx_messageInfo_DHHeader proto.InternalMessageInfo

func (m *DHHeader) GetKey() []byte {
	if m != nil {
		return m.Key
	}
	return nil
}

type X3DHHeader struct {
	// Ephemeral key used
	Key []byte `protobuf:"bytes,1,opt,name=key,proto3" json:"key,omitempty"`
	// Used bundle's signed prekey
	Id                   []byte   `protobuf:"bytes,4,opt,name=id,proto3" json:"id,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *X3DHHeader) Reset()         { *m = X3DHHeader{} }
func (m *X3DHHeader) String() string { return proto.CompactTextString(m) }
func (*X3DHHeader) ProtoMessage()    {}
func (*X3DHHeader) Descriptor() ([]byte, []int) {
	return fileDescriptor_4e37b52004a72e16, []int{5}
}

func (m *X3DHHeader) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_X3DHHeader.Unmarshal(m, b)
}
func (m *X3DHHeader) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_X3DHHeader.Marshal(b, m, deterministic)
}
func (m *X3DHHeader) XXX_Merge(src proto.Message) {
	xxx_messageInfo_X3DHHeader.Merge(m, src)
}
func (m *X3DHHeader) XXX_Size() int {
	return xxx_messageInfo_X3DHHeader.Size(m)
}
func (m *X3DHHeader) XXX_DiscardUnknown() {
	xxx_messageInfo_X3DHHeader.DiscardUnknown(m)
}

var xxx_messageInfo_X3DHHeader proto.InternalMessageInfo

func (m *X3DHHeader) GetKey() []byte {
	if m != nil {
		return m.Key
	}
	return nil
}

func (m *X3DHHeader) GetId() []byte {
	if m != nil {
		return m.Id
	}
	return nil
}

// Hash Ratchet Header
type HRHeader struct {
	// community key ID
	KeyId uint32 `protobuf:"varint,1,opt,name=key_id,json=keyId,proto3" json:"key_id,omitempty"`
	// Community message number for this key_id
	SeqNo uint32 `protobuf:"varint,2,opt,name=seq_no,json=seqNo,proto3" json:"seq_no,omitempty"`
	// Community ID
	GroupId              []byte   `protobuf:"bytes,3,opt,name=group_id,json=groupId,proto3" json:"group_id,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *HRHeader) Reset()         { *m = HRHeader{} }
func (m *HRHeader) String() string { return proto.CompactTextString(m) }
func (*HRHeader) ProtoMessage()    {}
func (*HRHeader) Descriptor() ([]byte, []int) {
	return fileDescriptor_4e37b52004a72e16, []int{6}
}

func (m *HRHeader) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_HRHeader.Unmarshal(m, b)
}
func (m *HRHeader) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_HRHeader.Marshal(b, m, deterministic)
}
func (m *HRHeader) XXX_Merge(src proto.Message) {
	xxx_messageInfo_HRHeader.Merge(m, src)
}
func (m *HRHeader) XXX_Size() int {
	return xxx_messageInfo_HRHeader.Size(m)
}
func (m *HRHeader) XXX_DiscardUnknown() {
	xxx_messageInfo_HRHeader.DiscardUnknown(m)
}

var xxx_messageInfo_HRHeader proto.InternalMessageInfo

func (m *HRHeader) GetKeyId() uint32 {
	if m != nil {
		return m.KeyId
	}
	return 0
}

func (m *HRHeader) GetSeqNo() uint32 {
	if m != nil {
		return m.SeqNo
	}
	return 0
}

func (m *HRHeader) GetGroupId() []byte {
	if m != nil {
		return m.GroupId
	}
	return nil
}

// Direct message value
type EncryptedMessageProtocol struct {
	X3DHHeader *X3DHHeader `protobuf:"bytes,1,opt,name=X3DH_header,json=X3DHHeader,proto3" json:"X3DH_header,omitempty"`
	DRHeader   *DRHeader   `protobuf:"bytes,2,opt,name=DR_header,json=DRHeader,proto3" json:"DR_header,omitempty"`
	DHHeader   *DHHeader   `protobuf:"bytes,101,opt,name=DH_header,json=DHHeader,proto3" json:"DH_header,omitempty"`
	HRHeader   *HRHeader   `protobuf:"bytes,102,opt,name=HR_header,json=HRHeader,proto3" json:"HR_header,omitempty"`
	// Encrypted payload
	Payload              []byte   `protobuf:"bytes,3,opt,name=payload,proto3" json:"payload,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *EncryptedMessageProtocol) Reset()         { *m = EncryptedMessageProtocol{} }
func (m *EncryptedMessageProtocol) String() string { return proto.CompactTextString(m) }
func (*EncryptedMessageProtocol) ProtoMessage()    {}
func (*EncryptedMessageProtocol) Descriptor() ([]byte, []int) {
	return fileDescriptor_4e37b52004a72e16, []int{7}
}

func (m *EncryptedMessageProtocol) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_EncryptedMessageProtocol.Unmarshal(m, b)
}
func (m *EncryptedMessageProtocol) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_EncryptedMessageProtocol.Marshal(b, m, deterministic)
}
func (m *EncryptedMessageProtocol) XXX_Merge(src proto.Message) {
	xxx_messageInfo_EncryptedMessageProtocol.Merge(m, src)
}
func (m *EncryptedMessageProtocol) XXX_Size() int {
	return xxx_messageInfo_EncryptedMessageProtocol.Size(m)
}
func (m *EncryptedMessageProtocol) XXX_DiscardUnknown() {
	xxx_messageInfo_EncryptedMessageProtocol.DiscardUnknown(m)
}

var xxx_messageInfo_EncryptedMessageProtocol proto.InternalMessageInfo

func (m *EncryptedMessageProtocol) GetX3DHHeader() *X3DHHeader {
	if m != nil {
		return m.X3DHHeader
	}
	return nil
}

func (m *EncryptedMessageProtocol) GetDRHeader() *DRHeader {
	if m != nil {
		return m.DRHeader
	}
	return nil
}

func (m *EncryptedMessageProtocol) GetDHHeader() *DHHeader {
	if m != nil {
		return m.DHHeader
	}
	return nil
}

func (m *EncryptedMessageProtocol) GetHRHeader() *HRHeader {
	if m != nil {
		return m.HRHeader
	}
	return nil
}

func (m *EncryptedMessageProtocol) GetPayload() []byte {
	if m != nil {
		return m.Payload
	}
	return nil
}

// Top-level protocol message
type ProtocolMessage struct {
	// The device id of the sender
	InstallationId string `protobuf:"bytes,2,opt,name=installation_id,json=installationId,proto3" json:"installation_id,omitempty"`
	// List of bundles
	Bundles []*Bundle `protobuf:"bytes,3,rep,name=bundles,proto3" json:"bundles,omitempty"`
	// One to one message, encrypted, indexed by installation_id
	// TODO map here is redundant in case of community messages
	EncryptedMessage map[string]*EncryptedMessageProtocol `protobuf:"bytes,101,rep,name=encrypted_message,json=encryptedMessage,proto3" json:"encrypted_message,omitempty" protobuf_key:"bytes,1,opt,name=key,proto3" protobuf_val:"bytes,2,opt,name=value,proto3"`
	// Public chats, not encrypted
	PublicMessage        []byte   `protobuf:"bytes,102,opt,name=public_message,json=publicMessage,proto3" json:"public_message,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *ProtocolMessage) Reset()         { *m = ProtocolMessage{} }
func (m *ProtocolMessage) String() string { return proto.CompactTextString(m) }
func (*ProtocolMessage) ProtoMessage()    {}
func (*ProtocolMessage) Descriptor() ([]byte, []int) {
	return fileDescriptor_4e37b52004a72e16, []int{8}
}

func (m *ProtocolMessage) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_ProtocolMessage.Unmarshal(m, b)
}
func (m *ProtocolMessage) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_ProtocolMessage.Marshal(b, m, deterministic)
}
func (m *ProtocolMessage) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ProtocolMessage.Merge(m, src)
}
func (m *ProtocolMessage) XXX_Size() int {
	return xxx_messageInfo_ProtocolMessage.Size(m)
}
func (m *ProtocolMessage) XXX_DiscardUnknown() {
	xxx_messageInfo_ProtocolMessage.DiscardUnknown(m)
}

var xxx_messageInfo_ProtocolMessage proto.InternalMessageInfo

func (m *ProtocolMessage) GetInstallationId() string {
	if m != nil {
		return m.InstallationId
	}
	return ""
}

func (m *ProtocolMessage) GetBundles() []*Bundle {
	if m != nil {
		return m.Bundles
	}
	return nil
}

func (m *ProtocolMessage) GetEncryptedMessage() map[string]*EncryptedMessageProtocol {
	if m != nil {
		return m.EncryptedMessage
	}
	return nil
}

func (m *ProtocolMessage) GetPublicMessage() []byte {
	if m != nil {
		return m.PublicMessage
	}
	return nil
}

func init() {
	proto.RegisterType((*SignedPreKey)(nil), "encryption.SignedPreKey")
	proto.RegisterType((*Bundle)(nil), "encryption.Bundle")
	proto.RegisterMapType((map[string]*SignedPreKey)(nil), "encryption.Bundle.SignedPreKeysEntry")
	proto.RegisterType((*BundleContainer)(nil), "encryption.BundleContainer")
	proto.RegisterType((*DRHeader)(nil), "encryption.DRHeader")
	proto.RegisterType((*DHHeader)(nil), "encryption.DHHeader")
	proto.RegisterType((*X3DHHeader)(nil), "encryption.X3DHHeader")
	proto.RegisterType((*HRHeader)(nil), "encryption.HRHeader")
	proto.RegisterType((*EncryptedMessageProtocol)(nil), "encryption.EncryptedMessageProtocol")
	proto.RegisterType((*ProtocolMessage)(nil), "encryption.ProtocolMessage")
	proto.RegisterMapType((map[string]*EncryptedMessageProtocol)(nil), "encryption.ProtocolMessage.EncryptedMessageEntry")
}

func init() { proto.RegisterFile("protocol_message.proto", fileDescriptor_4e37b52004a72e16) }

var fileDescriptor_4e37b52004a72e16 = []byte{
	// 631 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x74, 0x54, 0x5d, 0x4b, 0xdc, 0x4c,
	0x14, 0x26, 0x89, 0xfb, 0xe1, 0xd9, 0x8f, 0xec, 0x3b, 0x6f, 0x95, 0x54, 0xbc, 0x58, 0x82, 0xd2,
	0x6d, 0x29, 0x29, 0x6a, 0xa1, 0xc5, 0xde, 0x59, 0x85, 0x68, 0xb1, 0xc8, 0x08, 0xa5, 0x78, 0xd1,
	0x10, 0xcd, 0xd1, 0x0e, 0xc6, 0x49, 0xcc, 0x64, 0xa5, 0xf9, 0x03, 0xbd, 0xe9, 0x3f, 0xec, 0xaf,
	0x29, 0x99, 0x64, 0x92, 0x71, 0x5d, 0xef, 0x72, 0x9e, 0x39, 0x5f, 0xcf, 0x73, 0xce, 0x09, 0xac,
	0xa7, 0x59, 0x92, 0x27, 0x57, 0x49, 0x1c, 0xdc, 0xa1, 0x10, 0xe1, 0x0d, 0x7a, 0x12, 0x20, 0x80,
	0xfc, 0x2a, 0x2b, 0xd2, 0x9c, 0x25, 0xdc, 0x2d, 0x60, 0x78, 0xce, 0x6e, 0x38, 0x46, 0x67, 0x19,
	0x7e, 0xc1, 0x82, 0x6c, 0xc1, 0x58, 0x48, 0x3b, 0x48, 0x33, 0x0c, 0x6e, 0xb1, 0x70, 0x8c, 0xa9,
	0x31, 0x1b, 0xd2, 0xa1, 0xd0, 0xbd, 0x1c, 0xe8, 0x3d, 0x60, 0x26, 0x58, 0xc2, 0x1d, 0x73, 0x6a,
	0xcc, 0x46, 0x54, 0x99, 0xe4, 0x35, 0x4c, 0x9a, 0xaa, 0xca, 0xc5, 0x92, 0x2e, 0xb6, 0xc2, 0xbf,
	0x55, 0xb0, 0xfb, 0xc7, 0x84, 0xee, 0xc1, 0x9c, 0x47, 0x31, 0x92, 0x0d, 0xe8, 0xb3, 0x08, 0x79,
	0xce, 0x72, 0x55, 0xaf, 0xb1, 0xc9, 0x29, 0xd8, 0x8f, 0x3b, 0x12, 0x8e, 0x39, 0xb5, 0x66, 0x83,
	0xdd, 0x6d, 0xaf, 0xe5, 0xe1, 0x55, 0x89, 0x3c, 0x9d, 0x8b, 0x38, 0xe2, 0x79, 0x56, 0xd0, 0x91,
	0xde, 0xb9, 0x20, 0x9b, 0xb0, 0x5a, 0x02, 0x61, 0x3e, 0xcf, 0xd0, 0x59, 0x91, 0xb5, 0x5a, 0xa0,
	0x7c, 0xcd, 0xd9, 0x1d, 0x8a, 0x3c, 0xbc, 0x4b, 0x9d, 0xce, 0xd4, 0x98, 0x59, 0xb4, 0x05, 0x36,
	0x2e, 0x80, 0x3c, 0x2d, 0x40, 0x26, 0x60, 0x29, 0x9d, 0x56, 0x69, 0xf9, 0x49, 0x3c, 0xe8, 0x3c,
	0x84, 0xf1, 0x1c, 0xa5, 0x38, 0x83, 0x5d, 0x47, 0x6f, 0x54, 0x4f, 0x40, 0x2b, 0xb7, 0x7d, 0xf3,
	0xa3, 0xe1, 0xfe, 0x02, 0xbb, 0xe2, 0xf0, 0x39, 0xe1, 0x79, 0xc8, 0x38, 0x66, 0xe4, 0x0d, 0x74,
	0x2f, 0x25, 0x24, 0x73, 0x0f, 0x76, 0xc9, 0x53, 0xc2, 0xb4, 0xf6, 0x20, 0x7b, 0xe5, 0xb4, 0xd9,
	0x43, 0x98, 0x63, 0xb0, 0x30, 0x3f, 0x53, 0x72, 0xfc, 0xbf, 0x7e, 0xd5, 0xcb, 0x9f, 0xac, 0xf4,
	0xad, 0xc9, 0x8a, 0x7b, 0x02, 0xfd, 0x43, 0xea, 0x63, 0x18, 0x61, 0xa6, 0x73, 0x19, 0x56, 0x5c,
	0x86, 0x60, 0xa8, 0x21, 0x1b, 0x9c, 0x8c, 0xc1, 0x4c, 0xd5, 0x40, 0xcd, 0x54, 0xda, 0x2c, 0xaa,
	0x65, 0x34, 0x59, 0xe4, 0x6e, 0x42, 0xff, 0xd0, 0x7f, 0x2e, 0x97, 0xfb, 0x1e, 0xe0, 0xfb, 0xde,
	0xf3, 0xef, 0x8b, 0xd9, 0xea, 0xfe, 0xce, 0xa1, 0xef, 0xab, 0xfe, 0xd6, 0xa0, 0x7b, 0x8b, 0x45,
	0xc0, 0x22, 0x19, 0x36, 0xa2, 0x9d, 0x5b, 0x2c, 0x8e, 0xa3, 0x12, 0x16, 0x78, 0x1f, 0xf0, 0xa4,
	0xee, 0xb4, 0x23, 0xf0, 0xfe, 0x6b, 0x42, 0x5e, 0x42, 0xff, 0x26, 0x4b, 0xe6, 0x69, 0xe9, 0x6f,
	0xc9, 0xac, 0x3d, 0x69, 0x1f, 0x47, 0xee, 0x6f, 0x13, 0x9c, 0xa3, 0x4a, 0x4d, 0x8c, 0x4e, 0xab,
	0xf3, 0x38, 0xab, 0x17, 0x94, 0x7c, 0x80, 0x41, 0xd9, 0x67, 0xf0, 0x53, 0x16, 0xad, 0xd5, 0x5f,
	0xd7, 0xd5, 0x6f, 0x69, 0x50, 0x9d, 0xd2, 0x0e, 0xac, 0x1e, 0x52, 0x15, 0x56, 0x0d, 0xff, 0x85,
	0x1e, 0xa6, 0x74, 0xa6, 0xad, 0xe2, 0x65, 0x48, 0x53, 0x09, 0x97, 0x84, 0xf8, 0x4d, 0x88, 0x56,
	0xc5, 0x6f, 0xaa, 0x5c, 0x3f, 0x0d, 0xf1, 0x9b, 0x2a, 0x8d, 0x6e, 0x0e, 0xf4, 0xd2, 0xb0, 0x88,
	0x93, 0xb0, 0x11, 0xa2, 0x36, 0xdd, 0xbf, 0x26, 0xd8, 0x8a, 0x78, 0xad, 0x03, 0x79, 0x05, 0x36,
	0xe3, 0x22, 0x0f, 0xe3, 0x38, 0x2c, 0x13, 0x96, 0xf2, 0x99, 0x72, 0xbb, 0xc7, 0x3a, 0x7c, 0x1c,
	0x91, 0xb7, 0xd0, 0xab, 0xf6, 0x4f, 0x38, 0x96, 0xbc, 0xc9, 0x65, 0x2b, 0xaa, 0x5c, 0xc8, 0x0f,
	0xf8, 0x0f, 0x95, 0xe4, 0xea, 0x97, 0xe4, 0xa0, 0x8c, 0xdb, 0xd1, 0xe3, 0x16, 0xda, 0xf1, 0x16,
	0xe7, 0x54, 0xdd, 0xf5, 0x04, 0x17, 0x60, 0xb2, 0x0d, 0xe3, 0x74, 0x7e, 0x19, 0xb3, 0xab, 0x26,
	0xf9, 0xb5, 0xe4, 0x3a, 0xaa, 0xd0, 0xda, 0x6d, 0x83, 0xc1, 0xda, 0xd2, 0x8c, 0x4b, 0x0e, 0x79,
	0xff, 0xf1, 0x21, 0x6f, 0xe9, 0x5d, 0x3e, 0xb7, 0x3d, 0xda, 0x51, 0x1f, 0xd8, 0x17, 0x23, 0xef,
	0xdd, 0xa7, 0x36, 0xe8, 0xb2, 0x2b, 0x7f, 0x82, 0x7b, 0xff, 0x02, 0x00, 0x00, 0xff, 0xff, 0x33,
	0xbf, 0x9f, 0x4a, 0x9b, 0x05, 0x00, 0x00,
}
