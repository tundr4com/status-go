// Code generated by protoc-gen-go. DO NOT EDIT.
// source: status_update.proto

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

type StatusUpdate_StatusType int32

const (
	StatusUpdate_UNKNOWN_STATUS_TYPE StatusUpdate_StatusType = 0
	StatusUpdate_AUTOMATIC           StatusUpdate_StatusType = 1
	StatusUpdate_DO_NOT_DISTURB      StatusUpdate_StatusType = 2
	StatusUpdate_ALWAYS_ONLINE       StatusUpdate_StatusType = 3
	StatusUpdate_INACTIVE            StatusUpdate_StatusType = 4
)

var StatusUpdate_StatusType_name = map[int32]string{
	0: "UNKNOWN_STATUS_TYPE",
	1: "AUTOMATIC",
	2: "DO_NOT_DISTURB",
	3: "ALWAYS_ONLINE",
	4: "INACTIVE",
}

var StatusUpdate_StatusType_value = map[string]int32{
	"UNKNOWN_STATUS_TYPE": 0,
	"AUTOMATIC":           1,
	"DO_NOT_DISTURB":      2,
	"ALWAYS_ONLINE":       3,
	"INACTIVE":            4,
}

func (x StatusUpdate_StatusType) String() string {
	return proto.EnumName(StatusUpdate_StatusType_name, int32(x))
}

func (StatusUpdate_StatusType) EnumDescriptor() ([]byte, []int) {
	return fileDescriptor_911acd91e62cd3d7, []int{0, 0}
}

// Specs:
//:AUTOMATIC
//To Send - "AUTOMATIC" status ping every 5 minutes
//Display - Online for up to 5 minutes from the last clock, after that Offline
//:ALWAYS_ONLINE
//To Send - "ALWAYS_ONLINE" status ping every 5 minutes
//Display - Online for up to 2 weeks from the last clock, after that Offline
//:INACTIVE
//To Send - A single "INACTIVE" status ping
//Display - Offline forever
//Note: Only send pings if the user interacted with the app in the last x minutes.
type StatusUpdate struct {
	Clock                uint64                  `protobuf:"varint,1,opt,name=clock,proto3" json:"clock,omitempty"`
	StatusType           StatusUpdate_StatusType `protobuf:"varint,2,opt,name=status_type,json=statusType,proto3,enum=protobuf.StatusUpdate_StatusType" json:"status_type,omitempty"`
	CustomText           string                  `protobuf:"bytes,3,opt,name=custom_text,json=customText,proto3" json:"custom_text,omitempty"`
	XXX_NoUnkeyedLiteral struct{}                `json:"-"`
	XXX_unrecognized     []byte                  `json:"-"`
	XXX_sizecache        int32                   `json:"-"`
}

func (m *StatusUpdate) Reset()         { *m = StatusUpdate{} }
func (m *StatusUpdate) String() string { return proto.CompactTextString(m) }
func (*StatusUpdate) ProtoMessage()    {}
func (*StatusUpdate) Descriptor() ([]byte, []int) {
	return fileDescriptor_911acd91e62cd3d7, []int{0}
}

func (m *StatusUpdate) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_StatusUpdate.Unmarshal(m, b)
}
func (m *StatusUpdate) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_StatusUpdate.Marshal(b, m, deterministic)
}
func (m *StatusUpdate) XXX_Merge(src proto.Message) {
	xxx_messageInfo_StatusUpdate.Merge(m, src)
}
func (m *StatusUpdate) XXX_Size() int {
	return xxx_messageInfo_StatusUpdate.Size(m)
}
func (m *StatusUpdate) XXX_DiscardUnknown() {
	xxx_messageInfo_StatusUpdate.DiscardUnknown(m)
}

var xxx_messageInfo_StatusUpdate proto.InternalMessageInfo

func (m *StatusUpdate) GetClock() uint64 {
	if m != nil {
		return m.Clock
	}
	return 0
}

func (m *StatusUpdate) GetStatusType() StatusUpdate_StatusType {
	if m != nil {
		return m.StatusType
	}
	return StatusUpdate_UNKNOWN_STATUS_TYPE
}

func (m *StatusUpdate) GetCustomText() string {
	if m != nil {
		return m.CustomText
	}
	return ""
}

func init() {
	proto.RegisterEnum("protobuf.StatusUpdate_StatusType", StatusUpdate_StatusType_name, StatusUpdate_StatusType_value)
	proto.RegisterType((*StatusUpdate)(nil), "protobuf.StatusUpdate")
}

func init() {
	proto.RegisterFile("status_update.proto", fileDescriptor_911acd91e62cd3d7)
}

var fileDescriptor_911acd91e62cd3d7 = []byte{
	// 253 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x4c, 0x8e, 0xc1, 0x4e, 0xc2, 0x40,
	0x10, 0x86, 0x5d, 0x40, 0x03, 0x53, 0x4a, 0xea, 0x60, 0x62, 0x6f, 0x56, 0x4e, 0x3d, 0xd5, 0x44,
	0x8f, 0x9e, 0xb6, 0xd0, 0x43, 0x23, 0x6e, 0x4d, 0x77, 0x56, 0x82, 0x97, 0x0d, 0xd4, 0x35, 0x31,
	0x6a, 0xda, 0xd8, 0x6d, 0x02, 0xef, 0xed, 0x03, 0x98, 0x14, 0x50, 0x4e, 0xf3, 0xff, 0x93, 0x6f,
	0xbe, 0x0c, 0x8c, 0x6b, 0xbb, 0xb2, 0x4d, 0xad, 0x9b, 0xea, 0x75, 0x65, 0x4d, 0x54, 0x7d, 0x97,
	0xb6, 0xc4, 0x7e, 0x3b, 0xd6, 0xcd, 0xdb, 0xe4, 0x87, 0xc1, 0x50, 0xb6, 0x84, 0x6a, 0x01, 0xbc,
	0x80, 0xd3, 0xe2, 0xb3, 0x2c, 0x3e, 0x7c, 0x16, 0xb0, 0xb0, 0x97, 0xef, 0x0a, 0xc6, 0xe0, 0xec,
	0x3d, 0x76, 0x5b, 0x19, 0xbf, 0x13, 0xb0, 0x70, 0x74, 0x7b, 0x1d, 0x1d, 0x34, 0xd1, 0xb1, 0x62,
	0x5f, 0x68, 0x5b, 0x99, 0x1c, 0xea, 0xbf, 0x8c, 0x57, 0xe0, 0x14, 0x4d, 0x6d, 0xcb, 0x2f, 0x6d,
	0xcd, 0xc6, 0xfa, 0xdd, 0x80, 0x85, 0x83, 0x1c, 0x76, 0x2b, 0x32, 0x1b, 0x3b, 0x79, 0x07, 0xf8,
	0x3f, 0xc5, 0x4b, 0x18, 0x2b, 0xf1, 0x20, 0xb2, 0x85, 0xd0, 0x92, 0x38, 0x29, 0xa9, 0x69, 0xf9,
	0x94, 0x78, 0x27, 0xe8, 0xc2, 0x80, 0x2b, 0xca, 0x1e, 0x39, 0xa5, 0x53, 0x8f, 0x21, 0xc2, 0x68,
	0x96, 0x69, 0x91, 0x91, 0x9e, 0xa5, 0x92, 0x54, 0x1e, 0x7b, 0x1d, 0x3c, 0x07, 0x97, 0xcf, 0x17,
	0x7c, 0x29, 0x75, 0x26, 0xe6, 0xa9, 0x48, 0xbc, 0x2e, 0x0e, 0xa1, 0x9f, 0x0a, 0x3e, 0xa5, 0xf4,
	0x39, 0xf1, 0x7a, 0xb1, 0xfb, 0xe2, 0x44, 0x37, 0xf7, 0x87, 0xf7, 0xd7, 0x67, 0x6d, 0xba, 0xfb,
	0x0d, 0x00, 0x00, 0xff, 0xff, 0xaa, 0xa1, 0x52, 0x1d, 0x2d, 0x01, 0x00, 0x00,
}
