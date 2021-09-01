// Code generated by protoc-gen-go. DO NOT EDIT.
// source: application_metadata_message.proto

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

type ApplicationMetadataMessage_Type int32

const (
	ApplicationMetadataMessage_UNKNOWN                                 ApplicationMetadataMessage_Type = 0
	ApplicationMetadataMessage_CHAT_MESSAGE                            ApplicationMetadataMessage_Type = 1
	ApplicationMetadataMessage_CONTACT_UPDATE                          ApplicationMetadataMessage_Type = 2
	ApplicationMetadataMessage_MEMBERSHIP_UPDATE_MESSAGE               ApplicationMetadataMessage_Type = 3
	ApplicationMetadataMessage_PAIR_INSTALLATION                       ApplicationMetadataMessage_Type = 4
	ApplicationMetadataMessage_SYNC_INSTALLATION                       ApplicationMetadataMessage_Type = 5
	ApplicationMetadataMessage_REQUEST_ADDRESS_FOR_TRANSACTION         ApplicationMetadataMessage_Type = 6
	ApplicationMetadataMessage_ACCEPT_REQUEST_ADDRESS_FOR_TRANSACTION  ApplicationMetadataMessage_Type = 7
	ApplicationMetadataMessage_DECLINE_REQUEST_ADDRESS_FOR_TRANSACTION ApplicationMetadataMessage_Type = 8
	ApplicationMetadataMessage_REQUEST_TRANSACTION                     ApplicationMetadataMessage_Type = 9
	ApplicationMetadataMessage_SEND_TRANSACTION                        ApplicationMetadataMessage_Type = 10
	ApplicationMetadataMessage_DECLINE_REQUEST_TRANSACTION             ApplicationMetadataMessage_Type = 11
	ApplicationMetadataMessage_SYNC_INSTALLATION_CONTACT               ApplicationMetadataMessage_Type = 12
	ApplicationMetadataMessage_SYNC_INSTALLATION_ACCOUNT               ApplicationMetadataMessage_Type = 13
	ApplicationMetadataMessage_SYNC_INSTALLATION_PUBLIC_CHAT           ApplicationMetadataMessage_Type = 14
	ApplicationMetadataMessage_CONTACT_CODE_ADVERTISEMENT              ApplicationMetadataMessage_Type = 15
	ApplicationMetadataMessage_PUSH_NOTIFICATION_REGISTRATION          ApplicationMetadataMessage_Type = 16
	ApplicationMetadataMessage_PUSH_NOTIFICATION_REGISTRATION_RESPONSE ApplicationMetadataMessage_Type = 17
	ApplicationMetadataMessage_PUSH_NOTIFICATION_QUERY                 ApplicationMetadataMessage_Type = 18
	ApplicationMetadataMessage_PUSH_NOTIFICATION_QUERY_RESPONSE        ApplicationMetadataMessage_Type = 19
	ApplicationMetadataMessage_PUSH_NOTIFICATION_REQUEST               ApplicationMetadataMessage_Type = 20
	ApplicationMetadataMessage_PUSH_NOTIFICATION_RESPONSE              ApplicationMetadataMessage_Type = 21
	ApplicationMetadataMessage_EMOJI_REACTION                          ApplicationMetadataMessage_Type = 22
	ApplicationMetadataMessage_GROUP_CHAT_INVITATION                   ApplicationMetadataMessage_Type = 23
	ApplicationMetadataMessage_CHAT_IDENTITY                           ApplicationMetadataMessage_Type = 24
	ApplicationMetadataMessage_COMMUNITY_DESCRIPTION                   ApplicationMetadataMessage_Type = 25
	ApplicationMetadataMessage_COMMUNITY_INVITATION                    ApplicationMetadataMessage_Type = 26
	ApplicationMetadataMessage_COMMUNITY_REQUEST_TO_JOIN               ApplicationMetadataMessage_Type = 27
	ApplicationMetadataMessage_PIN_MESSAGE                             ApplicationMetadataMessage_Type = 28
	ApplicationMetadataMessage_EDIT_MESSAGE                            ApplicationMetadataMessage_Type = 29
	ApplicationMetadataMessage_STATUS_UPDATE                           ApplicationMetadataMessage_Type = 30
	ApplicationMetadataMessage_DELETE_MESSAGE                          ApplicationMetadataMessage_Type = 31
	ApplicationMetadataMessage_SYNC_INSTALLATION_COMMUNITY             ApplicationMetadataMessage_Type = 32
	ApplicationMetadataMessage_ANONYMOUS_METRIC_BATCH                  ApplicationMetadataMessage_Type = 33
)

var ApplicationMetadataMessage_Type_name = map[int32]string{
	0:  "UNKNOWN",
	1:  "CHAT_MESSAGE",
	2:  "CONTACT_UPDATE",
	3:  "MEMBERSHIP_UPDATE_MESSAGE",
	4:  "PAIR_INSTALLATION",
	5:  "SYNC_INSTALLATION",
	6:  "REQUEST_ADDRESS_FOR_TRANSACTION",
	7:  "ACCEPT_REQUEST_ADDRESS_FOR_TRANSACTION",
	8:  "DECLINE_REQUEST_ADDRESS_FOR_TRANSACTION",
	9:  "REQUEST_TRANSACTION",
	10: "SEND_TRANSACTION",
	11: "DECLINE_REQUEST_TRANSACTION",
	12: "SYNC_INSTALLATION_CONTACT",
	13: "SYNC_INSTALLATION_ACCOUNT",
	14: "SYNC_INSTALLATION_PUBLIC_CHAT",
	15: "CONTACT_CODE_ADVERTISEMENT",
	16: "PUSH_NOTIFICATION_REGISTRATION",
	17: "PUSH_NOTIFICATION_REGISTRATION_RESPONSE",
	18: "PUSH_NOTIFICATION_QUERY",
	19: "PUSH_NOTIFICATION_QUERY_RESPONSE",
	20: "PUSH_NOTIFICATION_REQUEST",
	21: "PUSH_NOTIFICATION_RESPONSE",
	22: "EMOJI_REACTION",
	23: "GROUP_CHAT_INVITATION",
	24: "CHAT_IDENTITY",
	25: "COMMUNITY_DESCRIPTION",
	26: "COMMUNITY_INVITATION",
	27: "COMMUNITY_REQUEST_TO_JOIN",
	28: "PIN_MESSAGE",
	29: "EDIT_MESSAGE",
	30: "STATUS_UPDATE",
	31: "DELETE_MESSAGE",
	32: "SYNC_INSTALLATION_COMMUNITY",
	33: "ANONYMOUS_METRIC_BATCH",
}

var ApplicationMetadataMessage_Type_value = map[string]int32{
	"UNKNOWN":                                 0,
	"CHAT_MESSAGE":                            1,
	"CONTACT_UPDATE":                          2,
	"MEMBERSHIP_UPDATE_MESSAGE":               3,
	"PAIR_INSTALLATION":                       4,
	"SYNC_INSTALLATION":                       5,
	"REQUEST_ADDRESS_FOR_TRANSACTION":         6,
	"ACCEPT_REQUEST_ADDRESS_FOR_TRANSACTION":  7,
	"DECLINE_REQUEST_ADDRESS_FOR_TRANSACTION": 8,
	"REQUEST_TRANSACTION":                     9,
	"SEND_TRANSACTION":                        10,
	"DECLINE_REQUEST_TRANSACTION":             11,
	"SYNC_INSTALLATION_CONTACT":               12,
	"SYNC_INSTALLATION_ACCOUNT":               13,
	"SYNC_INSTALLATION_PUBLIC_CHAT":           14,
	"CONTACT_CODE_ADVERTISEMENT":              15,
	"PUSH_NOTIFICATION_REGISTRATION":          16,
	"PUSH_NOTIFICATION_REGISTRATION_RESPONSE": 17,
	"PUSH_NOTIFICATION_QUERY":                 18,
	"PUSH_NOTIFICATION_QUERY_RESPONSE":        19,
	"PUSH_NOTIFICATION_REQUEST":               20,
	"PUSH_NOTIFICATION_RESPONSE":              21,
	"EMOJI_REACTION":                          22,
	"GROUP_CHAT_INVITATION":                   23,
	"CHAT_IDENTITY":                           24,
	"COMMUNITY_DESCRIPTION":                   25,
	"COMMUNITY_INVITATION":                    26,
	"COMMUNITY_REQUEST_TO_JOIN":               27,
	"PIN_MESSAGE":                             28,
	"EDIT_MESSAGE":                            29,
	"STATUS_UPDATE":                           30,
	"DELETE_MESSAGE":                          31,
	"SYNC_INSTALLATION_COMMUNITY":             32,
	"ANONYMOUS_METRIC_BATCH":                  33,
}

func (x ApplicationMetadataMessage_Type) String() string {
	return proto.EnumName(ApplicationMetadataMessage_Type_name, int32(x))
}

func (ApplicationMetadataMessage_Type) EnumDescriptor() ([]byte, []int) {
	return fileDescriptor_ad09a6406fcf24c7, []int{0, 0}
}

type ApplicationMetadataMessage struct {
	// Signature of the payload field
	Signature []byte `protobuf:"bytes,1,opt,name=signature,proto3" json:"signature,omitempty"`
	// This is the encoded protobuf of the application level message, i.e ChatMessage
	Payload []byte `protobuf:"bytes,2,opt,name=payload,proto3" json:"payload,omitempty"`
	// The type of protobuf message sent
	Type                 ApplicationMetadataMessage_Type `protobuf:"varint,3,opt,name=type,proto3,enum=protobuf.ApplicationMetadataMessage_Type" json:"type,omitempty"`
	XXX_NoUnkeyedLiteral struct{}                        `json:"-"`
	XXX_unrecognized     []byte                          `json:"-"`
	XXX_sizecache        int32                           `json:"-"`
}

func (m *ApplicationMetadataMessage) Reset()         { *m = ApplicationMetadataMessage{} }
func (m *ApplicationMetadataMessage) String() string { return proto.CompactTextString(m) }
func (*ApplicationMetadataMessage) ProtoMessage()    {}
func (*ApplicationMetadataMessage) Descriptor() ([]byte, []int) {
	return fileDescriptor_ad09a6406fcf24c7, []int{0}
}

func (m *ApplicationMetadataMessage) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_ApplicationMetadataMessage.Unmarshal(m, b)
}
func (m *ApplicationMetadataMessage) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_ApplicationMetadataMessage.Marshal(b, m, deterministic)
}
func (m *ApplicationMetadataMessage) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ApplicationMetadataMessage.Merge(m, src)
}
func (m *ApplicationMetadataMessage) XXX_Size() int {
	return xxx_messageInfo_ApplicationMetadataMessage.Size(m)
}
func (m *ApplicationMetadataMessage) XXX_DiscardUnknown() {
	xxx_messageInfo_ApplicationMetadataMessage.DiscardUnknown(m)
}

var xxx_messageInfo_ApplicationMetadataMessage proto.InternalMessageInfo

func (m *ApplicationMetadataMessage) GetSignature() []byte {
	if m != nil {
		return m.Signature
	}
	return nil
}

func (m *ApplicationMetadataMessage) GetPayload() []byte {
	if m != nil {
		return m.Payload
	}
	return nil
}

func (m *ApplicationMetadataMessage) GetType() ApplicationMetadataMessage_Type {
	if m != nil {
		return m.Type
	}
	return ApplicationMetadataMessage_UNKNOWN
}

func init() {
	proto.RegisterEnum("protobuf.ApplicationMetadataMessage_Type", ApplicationMetadataMessage_Type_name, ApplicationMetadataMessage_Type_value)
	proto.RegisterType((*ApplicationMetadataMessage)(nil), "protobuf.ApplicationMetadataMessage")
}

func init() {
	proto.RegisterFile("application_metadata_message.proto", fileDescriptor_ad09a6406fcf24c7)
}

var fileDescriptor_ad09a6406fcf24c7 = []byte{
	// 611 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x84, 0x53, 0x5d, 0x4f, 0x1b, 0x3b,
	0x10, 0xbd, 0x01, 0x2e, 0x81, 0x09, 0x1f, 0x66, 0xf8, 0x0a, 0xe1, 0x2b, 0xe4, 0x5e, 0xb5, 0xb4,
	0x95, 0x52, 0xa9, 0x7d, 0xac, 0xfa, 0xe0, 0x78, 0x07, 0x62, 0x9a, 0xb5, 0x17, 0xdb, 0x4b, 0x95,
	0xbe, 0x58, 0x4b, 0x49, 0x11, 0x12, 0x90, 0x08, 0xc2, 0x03, 0x3f, 0xb0, 0xbf, 0xa2, 0x7f, 0xa6,
	0xda, 0x4d, 0x42, 0x42, 0x09, 0xe5, 0x69, 0xd7, 0xe7, 0x9c, 0x99, 0xf1, 0x1c, 0xcf, 0x40, 0x25,
	0xe9, 0x74, 0x2e, 0x2f, 0xbe, 0x27, 0xdd, 0x8b, 0xf6, 0xb5, 0xbf, 0x6a, 0x75, 0x93, 0xb3, 0xa4,
	0x9b, 0xf8, 0xab, 0xd6, 0xed, 0x6d, 0x72, 0xde, 0xaa, 0x76, 0x6e, 0xda, 0xdd, 0x36, 0xce, 0x64,
	0x9f, 0xd3, 0xbb, 0x1f, 0x95, 0x5f, 0x33, 0x50, 0xe2, 0xc3, 0x80, 0xb0, 0xaf, 0x0f, 0x7b, 0x72,
	0xdc, 0x82, 0xd9, 0xdb, 0x8b, 0xf3, 0xeb, 0xa4, 0x7b, 0x77, 0xd3, 0x2a, 0xe6, 0xca, 0xb9, 0xfd,
	0x39, 0x33, 0x04, 0xb0, 0x08, 0xf9, 0x4e, 0x72, 0x7f, 0xd9, 0x4e, 0xce, 0x8a, 0x13, 0x19, 0x37,
	0x38, 0xe2, 0x67, 0x98, 0xea, 0xde, 0x77, 0x5a, 0xc5, 0xc9, 0x72, 0x6e, 0x7f, 0xe1, 0xc3, 0x9b,
	0xea, 0xa0, 0x5e, 0xf5, 0xf9, 0x5a, 0x55, 0x77, 0xdf, 0x69, 0x99, 0x2c, 0xac, 0xf2, 0x33, 0x0f,
	0x53, 0xe9, 0x11, 0x0b, 0x90, 0x8f, 0xd5, 0x17, 0xa5, 0xbf, 0x2a, 0xf6, 0x0f, 0x32, 0x98, 0x13,
	0x75, 0xee, 0x7c, 0x48, 0xd6, 0xf2, 0x43, 0x62, 0x39, 0x44, 0x58, 0x10, 0x5a, 0x39, 0x2e, 0x9c,
	0x8f, 0xa3, 0x80, 0x3b, 0x62, 0x13, 0xb8, 0x0d, 0x1b, 0x21, 0x85, 0x35, 0x32, 0xb6, 0x2e, 0xa3,
	0x3e, 0xfc, 0x10, 0x32, 0x89, 0xab, 0xb0, 0x14, 0x71, 0x69, 0xbc, 0x54, 0xd6, 0xf1, 0x46, 0x83,
	0x3b, 0xa9, 0x15, 0x9b, 0x4a, 0x61, 0xdb, 0x54, 0xe2, 0x31, 0xfc, 0x2f, 0xfe, 0x07, 0xbb, 0x86,
	0x8e, 0x63, 0xb2, 0xce, 0xf3, 0x20, 0x30, 0x64, 0xad, 0x3f, 0xd0, 0xc6, 0x3b, 0xc3, 0x95, 0xe5,
	0x22, 0x13, 0x4d, 0xe3, 0x5b, 0x78, 0xc5, 0x85, 0xa0, 0xc8, 0xf9, 0x97, 0xb4, 0x79, 0x7c, 0x07,
	0xaf, 0x03, 0x12, 0x0d, 0xa9, 0xe8, 0x45, 0xf1, 0x0c, 0xae, 0xc3, 0xf2, 0x40, 0x34, 0x4a, 0xcc,
	0xe2, 0x0a, 0x30, 0x4b, 0x2a, 0x78, 0x84, 0x02, 0xee, 0xc2, 0xe6, 0x9f, 0xb9, 0x47, 0x05, 0x85,
	0xd4, 0x9a, 0x27, 0x4d, 0xfa, 0xbe, 0x81, 0x6c, 0x6e, 0x3c, 0xcd, 0x85, 0xd0, 0xb1, 0x72, 0x6c,
	0x1e, 0xf7, 0x60, 0xfb, 0x29, 0x1d, 0xc5, 0xb5, 0x86, 0x14, 0x3e, 0x7d, 0x17, 0xb6, 0x80, 0x3b,
	0x50, 0x1a, 0xbc, 0x87, 0xd0, 0x01, 0x79, 0x1e, 0x9c, 0x90, 0x71, 0xd2, 0x52, 0x48, 0xca, 0xb1,
	0x45, 0xac, 0xc0, 0x4e, 0x14, 0xdb, 0xba, 0x57, 0xda, 0xc9, 0x03, 0x29, 0x7a, 0x29, 0x0c, 0x1d,
	0x4a, 0xeb, 0x4c, 0xcf, 0x72, 0x96, 0x3a, 0xf4, 0x77, 0x8d, 0x37, 0x64, 0x23, 0xad, 0x2c, 0xb1,
	0x25, 0xdc, 0x84, 0xf5, 0xa7, 0xe2, 0xe3, 0x98, 0x4c, 0x93, 0x21, 0xfe, 0x0f, 0xe5, 0x67, 0xc8,
	0x61, 0x8a, 0xe5, 0xb4, 0xeb, 0x71, 0xf5, 0x32, 0xff, 0xd8, 0x4a, 0xda, 0xd2, 0x38, 0xba, 0x1f,
	0xbe, 0x9a, 0x8e, 0x20, 0x85, 0xfa, 0x48, 0x7a, 0x43, 0x7d, 0x9f, 0xd7, 0x70, 0x03, 0x56, 0x0f,
	0x8d, 0x8e, 0xa3, 0xcc, 0x16, 0x2f, 0xd5, 0x89, 0x74, 0xbd, 0xee, 0xd6, 0x71, 0x09, 0xe6, 0x7b,
	0x60, 0x40, 0xca, 0x49, 0xd7, 0x64, 0xc5, 0x54, 0x2d, 0x74, 0x18, 0xc6, 0x4a, 0xba, 0xa6, 0x0f,
	0xc8, 0x0a, 0x23, 0xa3, 0x4c, 0xbd, 0x81, 0x45, 0x58, 0x19, 0x52, 0x23, 0x79, 0x4a, 0xe9, 0xad,
	0x87, 0xcc, 0xc3, 0x6b, 0x6b, 0x7f, 0xa4, 0xa5, 0x62, 0x9b, 0xb8, 0x08, 0x85, 0x48, 0xaa, 0x87,
	0xb1, 0xdf, 0x4a, 0x77, 0x87, 0x02, 0x39, 0xdc, 0x9d, 0xed, 0xf4, 0x26, 0xd6, 0x71, 0x17, 0xdb,
	0xc1, 0xea, 0xec, 0xa4, 0xbd, 0x04, 0xd4, 0xa0, 0x91, 0x7d, 0xd9, 0x4d, 0x87, 0x6a, 0xdc, 0xcc,
	0xf4, 0x4b, 0xb3, 0x32, 0x96, 0x60, 0x8d, 0x2b, 0xad, 0x9a, 0xa1, 0x8e, 0xad, 0x0f, 0xc9, 0x19,
	0x29, 0x7c, 0x8d, 0x3b, 0x51, 0x67, 0x7b, 0xb5, 0xf9, 0x6f, 0x85, 0xea, 0xfb, 0x4f, 0x83, 0xe5,
	0x3f, 0x9d, 0xce, 0xfe, 0x3e, 0xfe, 0x0e, 0x00, 0x00, 0xff, 0xff, 0x33, 0xbf, 0x72, 0xe8, 0xa3,
	0x04, 0x00, 0x00,
}
