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
	ApplicationMetadataMessage_SYNC_CHAT_REMOVED                       ApplicationMetadataMessage_Type = 34
	ApplicationMetadataMessage_SYNC_CHAT_MESSAGES_READ                 ApplicationMetadataMessage_Type = 35
	ApplicationMetadataMessage_BACKUP                                  ApplicationMetadataMessage_Type = 36
	ApplicationMetadataMessage_SYNC_ACTIVITY_CENTER_READ               ApplicationMetadataMessage_Type = 37
	ApplicationMetadataMessage_SYNC_ACTIVITY_CENTER_ACCEPTED           ApplicationMetadataMessage_Type = 38
	ApplicationMetadataMessage_SYNC_ACTIVITY_CENTER_DISMISSED          ApplicationMetadataMessage_Type = 39
	ApplicationMetadataMessage_SYNC_CLEAR_HISTORY                      ApplicationMetadataMessage_Type = 41
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
	34: "SYNC_CHAT_REMOVED",
	35: "SYNC_CHAT_MESSAGES_READ",
	36: "BACKUP",
	37: "SYNC_ACTIVITY_CENTER_READ",
	38: "SYNC_ACTIVITY_CENTER_ACCEPTED",
	39: "SYNC_ACTIVITY_CENTER_DISMISSED",
	41: "SYNC_CLEAR_HISTORY",
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
	"SYNC_CHAT_REMOVED":                       34,
	"SYNC_CHAT_MESSAGES_READ":                 35,
	"BACKUP":                                  36,
	"SYNC_ACTIVITY_CENTER_READ":               37,
	"SYNC_ACTIVITY_CENTER_ACCEPTED":           38,
	"SYNC_ACTIVITY_CENTER_DISMISSED":          39,
	"SYNC_CLEAR_HISTORY":                      41,
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
	// 692 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x84, 0x54, 0x4d, 0x53, 0x23, 0x37,
	0x10, 0x8d, 0x77, 0x09, 0x2c, 0x6d, 0x60, 0x85, 0x96, 0x0f, 0x63, 0x16, 0xf0, 0x1a, 0xc2, 0x47,
	0x52, 0xe5, 0x54, 0x25, 0xc7, 0x54, 0x0e, 0xb2, 0xd4, 0x60, 0x81, 0x47, 0x1a, 0x24, 0x8d, 0x53,
	0xce, 0x45, 0x35, 0x04, 0x87, 0xa2, 0x0a, 0xb0, 0x0b, 0xcc, 0x81, 0x73, 0x7e, 0x45, 0xfe, 0x6d,
	0x4a, 0x33, 0xfe, 0x02, 0x4c, 0x38, 0xd9, 0xea, 0xf7, 0xd4, 0xad, 0x7e, 0xfd, 0x7a, 0xa0, 0x9a,
	0xf6, 0x7a, 0x37, 0xd7, 0x7f, 0xa5, 0xfd, 0xeb, 0xee, 0x9d, 0xbf, 0xed, 0xf4, 0xd3, 0xcb, 0xb4,
	0x9f, 0xfa, 0xdb, 0xce, 0xc3, 0x43, 0x7a, 0xd5, 0xa9, 0xf5, 0xee, 0xbb, 0xfd, 0x2e, 0xfd, 0x94,
	0xfd, 0x5c, 0x3c, 0xfe, 0x5d, 0xfd, 0x17, 0xa0, 0xcc, 0xc6, 0x17, 0xa2, 0x01, 0x3f, 0xca, 0xe9,
	0xf4, 0x2b, 0xcc, 0x3f, 0x5c, 0x5f, 0xdd, 0xa5, 0xfd, 0xc7, 0xfb, 0x4e, 0xa9, 0x50, 0x29, 0x1c,
	0x2e, 0x98, 0x71, 0x80, 0x96, 0x60, 0xae, 0x97, 0x3e, 0xdd, 0x74, 0xd3, 0xcb, 0xd2, 0x87, 0x0c,
	0x1b, 0x1e, 0xe9, 0xef, 0x30, 0xd3, 0x7f, 0xea, 0x75, 0x4a, 0x1f, 0x2b, 0x85, 0xc3, 0xa5, 0x5f,
	0x8e, 0x6a, 0xc3, 0x7a, 0xb5, 0xb7, 0x6b, 0xd5, 0xdc, 0x53, 0xaf, 0x63, 0xb2, 0x6b, 0xd5, 0x7f,
	0xe6, 0x61, 0x26, 0x1c, 0x69, 0x11, 0xe6, 0x12, 0x75, 0xa6, 0xf4, 0x1f, 0x8a, 0x7c, 0x47, 0x09,
	0x2c, 0xf0, 0x06, 0x73, 0x3e, 0x42, 0x6b, 0xd9, 0x09, 0x92, 0x02, 0xa5, 0xb0, 0xc4, 0xb5, 0x72,
	0x8c, 0x3b, 0x9f, 0xc4, 0x82, 0x39, 0x24, 0x1f, 0xe8, 0x16, 0x6c, 0x44, 0x18, 0xd5, 0xd1, 0xd8,
	0x86, 0x8c, 0x07, 0xe1, 0xd1, 0x95, 0x8f, 0x74, 0x15, 0x96, 0x63, 0x26, 0x8d, 0x97, 0xca, 0x3a,
	0xd6, 0x6c, 0x32, 0x27, 0xb5, 0x22, 0x33, 0x21, 0x6c, 0xdb, 0x8a, 0x3f, 0x0f, 0x7f, 0x4f, 0x77,
	0x61, 0xc7, 0xe0, 0x79, 0x82, 0xd6, 0x79, 0x26, 0x84, 0x41, 0x6b, 0xfd, 0xb1, 0x36, 0xde, 0x19,
	0xa6, 0x2c, 0xe3, 0x19, 0x69, 0x96, 0xfe, 0x08, 0xfb, 0x8c, 0x73, 0x8c, 0x9d, 0x7f, 0x8f, 0x3b,
	0x47, 0x7f, 0x82, 0x03, 0x81, 0xbc, 0x29, 0x15, 0xbe, 0x4b, 0xfe, 0x44, 0xd7, 0xe1, 0xcb, 0x90,
	0x34, 0x09, 0xcc, 0xd3, 0x15, 0x20, 0x16, 0x95, 0x78, 0x16, 0x05, 0xba, 0x03, 0x9b, 0x2f, 0x73,
	0x4f, 0x12, 0x8a, 0x41, 0x9a, 0x57, 0x4d, 0xfa, 0x81, 0x80, 0x64, 0x61, 0x3a, 0xcc, 0x38, 0xd7,
	0x89, 0x72, 0x64, 0x91, 0x7e, 0x83, 0xad, 0xd7, 0x70, 0x9c, 0xd4, 0x9b, 0x92, 0xfb, 0x30, 0x17,
	0xb2, 0x44, 0xb7, 0xa1, 0x3c, 0x9c, 0x07, 0xd7, 0x02, 0x3d, 0x13, 0x2d, 0x34, 0x4e, 0x5a, 0x8c,
	0x50, 0x39, 0xf2, 0x99, 0x56, 0x61, 0x3b, 0x4e, 0x6c, 0xc3, 0x2b, 0xed, 0xe4, 0xb1, 0xe4, 0x79,
	0x0a, 0x83, 0x27, 0xd2, 0x3a, 0x93, 0x4b, 0x4e, 0x82, 0x42, 0xff, 0xcf, 0xf1, 0x06, 0x6d, 0xac,
	0x95, 0x45, 0xb2, 0x4c, 0x37, 0x61, 0xfd, 0x35, 0xf9, 0x3c, 0x41, 0xd3, 0x26, 0x94, 0xee, 0x41,
	0xe5, 0x0d, 0x70, 0x9c, 0xe2, 0x4b, 0xe8, 0x7a, 0x5a, 0xbd, 0x4c, 0x3f, 0xb2, 0x12, 0x5a, 0x9a,
	0x06, 0x0f, 0xae, 0xaf, 0x06, 0x0b, 0x62, 0xa4, 0x4f, 0xa5, 0x37, 0x38, 0xd0, 0x79, 0x8d, 0x6e,
	0xc0, 0xea, 0x89, 0xd1, 0x49, 0x9c, 0xc9, 0xe2, 0xa5, 0x6a, 0x49, 0x97, 0x77, 0xb7, 0x4e, 0x97,
	0x61, 0x31, 0x0f, 0x0a, 0x54, 0x4e, 0xba, 0x36, 0x29, 0x05, 0x36, 0xd7, 0x51, 0x94, 0x28, 0xe9,
	0xda, 0x5e, 0xa0, 0xe5, 0x46, 0xc6, 0x19, 0x7b, 0x83, 0x96, 0x60, 0x65, 0x0c, 0x4d, 0xe4, 0x29,
	0x87, 0x57, 0x8f, 0x91, 0xd1, 0xb4, 0xb5, 0x3f, 0xd5, 0x52, 0x91, 0x4d, 0xfa, 0x19, 0x8a, 0xb1,
	0x54, 0x23, 0xdb, 0x7f, 0x0d, 0xbb, 0x83, 0x42, 0x8e, 0x77, 0x67, 0x2b, 0xbc, 0xc4, 0x3a, 0xe6,
	0x12, 0x3b, 0x5c, 0x9d, 0xed, 0xd0, 0x8b, 0xc0, 0x26, 0x4e, 0xec, 0xcb, 0x4e, 0x30, 0xd5, 0x34,
	0xcf, 0x0c, 0x4a, 0x93, 0x0a, 0x2d, 0xc3, 0x1a, 0x53, 0x5a, 0xb5, 0x23, 0x9d, 0x58, 0x1f, 0xa1,
	0x33, 0x92, 0xfb, 0x3a, 0x73, 0xbc, 0x41, 0xbe, 0x8d, 0xb6, 0x2a, 0x6b, 0xd9, 0x60, 0xa4, 0x5b,
	0x28, 0x48, 0x35, 0x4c, 0x6d, 0x1c, 0x1e, 0x94, 0xb2, 0x41, 0x40, 0x41, 0x76, 0x29, 0xc0, 0x6c,
	0x9d, 0xf1, 0xb3, 0x24, 0x26, 0x7b, 0x23, 0x47, 0x06, 0x65, 0x5b, 0xa1, 0x53, 0x8e, 0xca, 0xa1,
	0xc9, 0xa9, 0x3f, 0x8c, 0x1c, 0xf9, 0x12, 0xce, 0xb7, 0x11, 0x05, 0xd9, 0x0f, 0x8e, 0x9b, 0x4a,
	0x11, 0xd2, 0x46, 0xd2, 0x5a, 0x14, 0xe4, 0x80, 0xae, 0x01, 0xcd, 0x9f, 0xd3, 0x44, 0x66, 0x7c,
	0x43, 0x5a, 0xa7, 0x4d, 0x9b, 0x1c, 0xd5, 0x17, 0xff, 0x2c, 0xd6, 0x7e, 0xfe, 0x6d, 0xf8, 0xe9,
	0xba, 0x98, 0xcd, 0xfe, 0xfd, 0xfa, 0x5f, 0x00, 0x00, 0x00, 0xff, 0xff, 0xa0, 0x61, 0x34, 0xbb,
	0x61, 0x05, 0x00, 0x00,
}
