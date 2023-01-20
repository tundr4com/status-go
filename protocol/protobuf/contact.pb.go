// Code generated by protoc-gen-go. DO NOT EDIT.
// source: contact.proto

package protobuf

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

type ContactUpdate struct {
	Clock                           uint64                   `protobuf:"varint,1,opt,name=clock,proto3" json:"clock,omitempty"`
	EnsName                         string                   `protobuf:"bytes,2,opt,name=ens_name,json=ensName,proto3" json:"ens_name,omitempty"`
	ProfileImage                    string                   `protobuf:"bytes,3,opt,name=profile_image,json=profileImage,proto3" json:"profile_image,omitempty"`
	DisplayName                     string                   `protobuf:"bytes,4,opt,name=display_name,json=displayName,proto3" json:"display_name,omitempty"`
	ContactRequestClock             uint64                   `protobuf:"varint,5,opt,name=contact_request_clock,json=contactRequestClock,proto3" json:"contact_request_clock,omitempty"`
	SentContactRequestSignature     *ContactRequestSignature `protobuf:"bytes,14,opt,name=sent_contact_request_signature,json=sentContactRequestSignature,proto3" json:"sent_contact_request_signature,omitempty"`
	ReceivedContactRequestSignature *ContactRequestSignature `protobuf:"bytes,15,opt,name=received_contact_request_signature,json=receivedContactRequestSignature,proto3" json:"received_contact_request_signature,omitempty"`
	XXX_NoUnkeyedLiteral            struct{}                 `json:"-"`
	XXX_unrecognized                []byte                   `json:"-"`
	XXX_sizecache                   int32                    `json:"-"`
}

func (m *ContactUpdate) Reset()         { *m = ContactUpdate{} }
func (m *ContactUpdate) String() string { return proto.CompactTextString(m) }
func (*ContactUpdate) ProtoMessage()    {}
func (*ContactUpdate) Descriptor() ([]byte, []int) {
	return fileDescriptor_a5036fff2565fb15, []int{0}
}

func (m *ContactUpdate) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_ContactUpdate.Unmarshal(m, b)
}
func (m *ContactUpdate) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_ContactUpdate.Marshal(b, m, deterministic)
}
func (m *ContactUpdate) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ContactUpdate.Merge(m, src)
}
func (m *ContactUpdate) XXX_Size() int {
	return xxx_messageInfo_ContactUpdate.Size(m)
}
func (m *ContactUpdate) XXX_DiscardUnknown() {
	xxx_messageInfo_ContactUpdate.DiscardUnknown(m)
}

var xxx_messageInfo_ContactUpdate proto.InternalMessageInfo

func (m *ContactUpdate) GetClock() uint64 {
	if m != nil {
		return m.Clock
	}
	return 0
}

func (m *ContactUpdate) GetEnsName() string {
	if m != nil {
		return m.EnsName
	}
	return ""
}

func (m *ContactUpdate) GetProfileImage() string {
	if m != nil {
		return m.ProfileImage
	}
	return ""
}

func (m *ContactUpdate) GetDisplayName() string {
	if m != nil {
		return m.DisplayName
	}
	return ""
}

func (m *ContactUpdate) GetContactRequestClock() uint64 {
	if m != nil {
		return m.ContactRequestClock
	}
	return 0
}

func (m *ContactUpdate) GetSentContactRequestSignature() *ContactRequestSignature {
	if m != nil {
		return m.SentContactRequestSignature
	}
	return nil
}

func (m *ContactUpdate) GetReceivedContactRequestSignature() *ContactRequestSignature {
	if m != nil {
		return m.ReceivedContactRequestSignature
	}
	return nil
}

type AcceptContactRequest struct {
	Id                   string   `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
	Clock                uint64   `protobuf:"varint,2,opt,name=clock,proto3" json:"clock,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *AcceptContactRequest) Reset()         { *m = AcceptContactRequest{} }
func (m *AcceptContactRequest) String() string { return proto.CompactTextString(m) }
func (*AcceptContactRequest) ProtoMessage()    {}
func (*AcceptContactRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_a5036fff2565fb15, []int{1}
}

func (m *AcceptContactRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_AcceptContactRequest.Unmarshal(m, b)
}
func (m *AcceptContactRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_AcceptContactRequest.Marshal(b, m, deterministic)
}
func (m *AcceptContactRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_AcceptContactRequest.Merge(m, src)
}
func (m *AcceptContactRequest) XXX_Size() int {
	return xxx_messageInfo_AcceptContactRequest.Size(m)
}
func (m *AcceptContactRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_AcceptContactRequest.DiscardUnknown(m)
}

var xxx_messageInfo_AcceptContactRequest proto.InternalMessageInfo

func (m *AcceptContactRequest) GetId() string {
	if m != nil {
		return m.Id
	}
	return ""
}

func (m *AcceptContactRequest) GetClock() uint64 {
	if m != nil {
		return m.Clock
	}
	return 0
}

type RetractContactRequest struct {
	Id                   string   `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
	Clock                uint64   `protobuf:"varint,2,opt,name=clock,proto3" json:"clock,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *RetractContactRequest) Reset()         { *m = RetractContactRequest{} }
func (m *RetractContactRequest) String() string { return proto.CompactTextString(m) }
func (*RetractContactRequest) ProtoMessage()    {}
func (*RetractContactRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_a5036fff2565fb15, []int{2}
}

func (m *RetractContactRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_RetractContactRequest.Unmarshal(m, b)
}
func (m *RetractContactRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_RetractContactRequest.Marshal(b, m, deterministic)
}
func (m *RetractContactRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_RetractContactRequest.Merge(m, src)
}
func (m *RetractContactRequest) XXX_Size() int {
	return xxx_messageInfo_RetractContactRequest.Size(m)
}
func (m *RetractContactRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_RetractContactRequest.DiscardUnknown(m)
}

var xxx_messageInfo_RetractContactRequest proto.InternalMessageInfo

func (m *RetractContactRequest) GetId() string {
	if m != nil {
		return m.Id
	}
	return ""
}

func (m *RetractContactRequest) GetClock() uint64 {
	if m != nil {
		return m.Clock
	}
	return 0
}

func init() {
	proto.RegisterType((*ContactUpdate)(nil), "protobuf.ContactUpdate")
	proto.RegisterType((*AcceptContactRequest)(nil), "protobuf.AcceptContactRequest")
	proto.RegisterType((*RetractContactRequest)(nil), "protobuf.RetractContactRequest")
}

func init() {
	proto.RegisterFile("contact.proto", fileDescriptor_a5036fff2565fb15)
}

var fileDescriptor_a5036fff2565fb15 = []byte{
	// 311 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x9c, 0x51, 0x31, 0x4f, 0xf3, 0x30,
	0x14, 0x54, 0xd2, 0xf6, 0xfb, 0xda, 0xd7, 0xa6, 0x48, 0xa6, 0x95, 0x02, 0x48, 0xd0, 0x86, 0xa5,
	0x53, 0x90, 0xca, 0x08, 0x0c, 0xd0, 0x89, 0x85, 0xc1, 0x88, 0x85, 0xc5, 0x72, 0x9d, 0xd7, 0x62,
	0xd1, 0x38, 0xc1, 0x76, 0x91, 0xf8, 0x1f, 0xfc, 0x60, 0x54, 0x3b, 0x11, 0x50, 0xa9, 0x08, 0x31,
	0x25, 0xbe, 0xbb, 0x77, 0xef, 0xec, 0x83, 0x48, 0x14, 0xca, 0x72, 0x61, 0xd3, 0x52, 0x17, 0xb6,
	0x20, 0x6d, 0xf7, 0x99, 0xaf, 0x17, 0x87, 0x44, 0x3c, 0x71, 0xcb, 0x72, 0x34, 0x86, 0x2f, 0xd1,
	0xb3, 0xc9, 0x7b, 0x03, 0xa2, 0x99, 0xd7, 0x3f, 0x94, 0x19, 0xb7, 0x48, 0x06, 0xd0, 0x12, 0xab,
	0x42, 0x3c, 0xc7, 0xc1, 0x28, 0x98, 0x34, 0xa9, 0x3f, 0x90, 0x03, 0x68, 0xa3, 0x32, 0x4c, 0xf1,
	0x1c, 0xe3, 0x70, 0x14, 0x4c, 0x3a, 0xf4, 0x3f, 0x2a, 0x73, 0xc7, 0x73, 0x24, 0xa7, 0x10, 0x95,
	0xba, 0x58, 0xc8, 0x15, 0x32, 0x99, 0xf3, 0x25, 0xc6, 0x0d, 0xc7, 0xf7, 0x2a, 0xf0, 0x76, 0x83,
	0x91, 0x31, 0xf4, 0x32, 0x69, 0xca, 0x15, 0x7f, 0xf3, 0x1e, 0x4d, 0xa7, 0xe9, 0x56, 0x98, 0xf3,
	0x99, 0xc2, 0xb0, 0x4a, 0xce, 0x34, 0xbe, 0xac, 0xd1, 0x58, 0xe6, 0x83, 0xb4, 0x5c, 0x90, 0xfd,
	0x8a, 0xa4, 0x9e, 0x9b, 0xb9, 0x58, 0x0b, 0x38, 0x36, 0xa8, 0x2c, 0xdb, 0x1e, 0x34, 0x72, 0xa9,
	0xb8, 0x5d, 0x6b, 0x8c, 0xfb, 0xa3, 0x60, 0xd2, 0x9d, 0x8e, 0xd3, 0xfa, 0x15, 0xd2, 0xd9, 0x37,
	0x9b, 0xfb, 0x5a, 0x48, 0x8f, 0x36, 0x46, 0x3b, 0x48, 0xa2, 0x20, 0xd1, 0x28, 0x50, 0xbe, 0x62,
	0xf6, 0xc3, 0xae, 0xbd, 0xdf, 0xee, 0x3a, 0xa9, 0xcd, 0x76, 0x08, 0x92, 0x4b, 0x18, 0x5c, 0x0b,
	0x81, 0xe5, 0x56, 0x20, 0xd2, 0x87, 0x50, 0x66, 0xae, 0x99, 0x0e, 0x0d, 0x65, 0xf6, 0x59, 0x56,
	0xf8, 0xa5, 0xac, 0xe4, 0x0a, 0x86, 0x14, 0xad, 0xe6, 0xe2, 0x4f, 0xe3, 0x37, 0xd1, 0x63, 0x37,
	0x3d, 0xbb, 0xa8, 0x2f, 0x31, 0xff, 0xe7, 0xfe, 0xce, 0x3f, 0x02, 0x00, 0x00, 0xff, 0xff, 0x8b,
	0xda, 0xbb, 0xda, 0x58, 0x02, 0x00, 0x00,
}
