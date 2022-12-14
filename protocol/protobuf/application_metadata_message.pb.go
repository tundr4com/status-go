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
	ApplicationMetadataMessage_SYNC_BOOKMARK                           ApplicationMetadataMessage_Type = 40
	ApplicationMetadataMessage_SYNC_CLEAR_HISTORY                      ApplicationMetadataMessage_Type = 41
	ApplicationMetadataMessage_SYNC_SETTING                            ApplicationMetadataMessage_Type = 42
	ApplicationMetadataMessage_COMMUNITY_ARCHIVE_MAGNETLINK            ApplicationMetadataMessage_Type = 43
	ApplicationMetadataMessage_SYNC_PROFILE_PICTURE                    ApplicationMetadataMessage_Type = 44
	ApplicationMetadataMessage_SYNC_WALLET_ACCOUNT                     ApplicationMetadataMessage_Type = 45
	ApplicationMetadataMessage_ACCEPT_CONTACT_REQUEST                  ApplicationMetadataMessage_Type = 46
	ApplicationMetadataMessage_RETRACT_CONTACT_REQUEST                 ApplicationMetadataMessage_Type = 47
	ApplicationMetadataMessage_COMMUNITY_REQUEST_TO_JOIN_RESPONSE      ApplicationMetadataMessage_Type = 48
	ApplicationMetadataMessage_SYNC_COMMUNITY_SETTINGS                 ApplicationMetadataMessage_Type = 49
	ApplicationMetadataMessage_REQUEST_CONTACT_VERIFICATION            ApplicationMetadataMessage_Type = 50
	ApplicationMetadataMessage_ACCEPT_CONTACT_VERIFICATION             ApplicationMetadataMessage_Type = 51
	ApplicationMetadataMessage_DECLINE_CONTACT_VERIFICATION            ApplicationMetadataMessage_Type = 52
	ApplicationMetadataMessage_SYNC_TRUSTED_USER                       ApplicationMetadataMessage_Type = 53
	ApplicationMetadataMessage_SYNC_VERIFICATION_REQUEST               ApplicationMetadataMessage_Type = 54
	ApplicationMetadataMessage_CONTACT_VERIFICATION_TRUSTED            ApplicationMetadataMessage_Type = 55
	ApplicationMetadataMessage_SYNC_CONTACT_REQUEST_DECISION           ApplicationMetadataMessage_Type = 56
	ApplicationMetadataMessage_COMMUNITY_REQUEST_TO_LEAVE              ApplicationMetadataMessage_Type = 57
	ApplicationMetadataMessage_SYNC_DELETE_FOR_ME_MESSAGE              ApplicationMetadataMessage_Type = 58
	ApplicationMetadataMessage_SYNC_SAVED_ADDRESS                      ApplicationMetadataMessage_Type = 59
	ApplicationMetadataMessage_COMMUNITY_CANCEL_REQUEST_TO_JOIN        ApplicationMetadataMessage_Type = 60
	ApplicationMetadataMessage_CANCEL_CONTACT_VERIFICATION             ApplicationMetadataMessage_Type = 61
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
	40: "SYNC_BOOKMARK",
	41: "SYNC_CLEAR_HISTORY",
	42: "SYNC_SETTING",
	43: "COMMUNITY_ARCHIVE_MAGNETLINK",
	44: "SYNC_PROFILE_PICTURE",
	45: "SYNC_WALLET_ACCOUNT",
	46: "ACCEPT_CONTACT_REQUEST",
	47: "RETRACT_CONTACT_REQUEST",
	48: "COMMUNITY_REQUEST_TO_JOIN_RESPONSE",
	49: "SYNC_COMMUNITY_SETTINGS",
	50: "REQUEST_CONTACT_VERIFICATION",
	51: "ACCEPT_CONTACT_VERIFICATION",
	52: "DECLINE_CONTACT_VERIFICATION",
	53: "SYNC_TRUSTED_USER",
	54: "SYNC_VERIFICATION_REQUEST",
	55: "CONTACT_VERIFICATION_TRUSTED",
	56: "SYNC_CONTACT_REQUEST_DECISION",
	57: "COMMUNITY_REQUEST_TO_LEAVE",
	58: "SYNC_DELETE_FOR_ME_MESSAGE",
	59: "SYNC_SAVED_ADDRESS",
	60: "COMMUNITY_CANCEL_REQUEST_TO_JOIN",
	61: "CANCEL_CONTACT_VERIFICATION",
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
	"SYNC_BOOKMARK":                           40,
	"SYNC_CLEAR_HISTORY":                      41,
	"SYNC_SETTING":                            42,
	"COMMUNITY_ARCHIVE_MAGNETLINK":            43,
	"SYNC_PROFILE_PICTURE":                    44,
	"SYNC_WALLET_ACCOUNT":                     45,
	"ACCEPT_CONTACT_REQUEST":                  46,
	"RETRACT_CONTACT_REQUEST":                 47,
	"COMMUNITY_REQUEST_TO_JOIN_RESPONSE":      48,
	"SYNC_COMMUNITY_SETTINGS":                 49,
	"REQUEST_CONTACT_VERIFICATION":            50,
	"ACCEPT_CONTACT_VERIFICATION":             51,
	"DECLINE_CONTACT_VERIFICATION":            52,
	"SYNC_TRUSTED_USER":                       53,
	"SYNC_VERIFICATION_REQUEST":               54,
	"CONTACT_VERIFICATION_TRUSTED":            55,
	"SYNC_CONTACT_REQUEST_DECISION":           56,
	"COMMUNITY_REQUEST_TO_LEAVE":              57,
	"SYNC_DELETE_FOR_ME_MESSAGE":              58,
	"SYNC_SAVED_ADDRESS":                      59,
	"COMMUNITY_CANCEL_REQUEST_TO_JOIN":        60,
	"CANCEL_CONTACT_VERIFICATION":             61,
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
	// 914 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x84, 0x55, 0x5b, 0x77, 0x13, 0x37,
	0x10, 0x6e, 0x20, 0x4d, 0x40, 0xb9, 0x29, 0x22, 0x17, 0xe7, 0x6e, 0x0c, 0x0d, 0x01, 0x5a, 0xd3,
	0x42, 0xef, 0x94, 0x07, 0x59, 0x9a, 0xd8, 0xc2, 0xbb, 0xd2, 0x22, 0x69, 0xcd, 0x71, 0x5f, 0x74,
	0x96, 0xe2, 0x72, 0x72, 0x0e, 0x10, 0x1f, 0x62, 0x1e, 0xf2, 0x2f, 0xfa, 0x7b, 0xfb, 0xd4, 0xa3,
	0xbd, 0x3a, 0xb6, 0x53, 0x9e, 0x92, 0x9d, 0xf9, 0x34, 0xa3, 0xf9, 0xe6, 0xfb, 0x64, 0xd4, 0x48,
	0x86, 0xc3, 0xf7, 0x67, 0x7f, 0x25, 0xa3, 0xb3, 0xf3, 0x8f, 0xee, 0xc3, 0x60, 0x94, 0xbc, 0x4d,
	0x46, 0x89, 0xfb, 0x30, 0xb8, 0xb8, 0x48, 0xde, 0x0d, 0x9a, 0xc3, 0x4f, 0xe7, 0xa3, 0x73, 0x72,
	0x2b, 0xfd, 0xf3, 0xe6, 0xf3, 0xdf, 0x8d, 0x7f, 0xd6, 0xd0, 0x2e, 0xad, 0x0e, 0x84, 0x39, 0x3e,
	0xcc, 0xe0, 0x64, 0x1f, 0xdd, 0xbe, 0x38, 0x7b, 0xf7, 0x31, 0x19, 0x7d, 0xfe, 0x34, 0xa8, 0xcd,
	0xd5, 0xe7, 0x4e, 0x96, 0x75, 0x15, 0x20, 0x35, 0xb4, 0x38, 0x4c, 0x2e, 0xdf, 0x9f, 0x27, 0x6f,
	0x6b, 0x37, 0xd2, 0x5c, 0xf1, 0x49, 0x5e, 0xa0, 0xf9, 0xd1, 0xe5, 0x70, 0x50, 0xbb, 0x59, 0x9f,
	0x3b, 0x59, 0x7d, 0xfa, 0xb0, 0x59, 0xf4, 0x6b, 0x5e, 0xdf, 0xab, 0x69, 0x2f, 0x87, 0x03, 0x9d,
	0x1e, 0x6b, 0xfc, 0xbb, 0x82, 0xe6, 0xfd, 0x27, 0x59, 0x42, 0x8b, 0xb1, 0xec, 0x4a, 0xf5, 0x5a,
	0xe2, 0xaf, 0x08, 0x46, 0xcb, 0xac, 0x43, 0xad, 0x0b, 0xc1, 0x18, 0xda, 0x06, 0x3c, 0x47, 0x08,
	0x5a, 0x65, 0x4a, 0x5a, 0xca, 0xac, 0x8b, 0x23, 0x4e, 0x2d, 0xe0, 0x1b, 0xe4, 0x00, 0xed, 0x84,
	0x10, 0xb6, 0x40, 0x9b, 0x8e, 0x88, 0xf2, 0x70, 0x79, 0xe4, 0x26, 0xd9, 0x44, 0xeb, 0x11, 0x15,
	0xda, 0x09, 0x69, 0x2c, 0x0d, 0x02, 0x6a, 0x85, 0x92, 0x78, 0xde, 0x87, 0x4d, 0x5f, 0xb2, 0xab,
	0xe1, 0xaf, 0xc9, 0x3d, 0x74, 0xa4, 0xe1, 0x55, 0x0c, 0xc6, 0x3a, 0xca, 0xb9, 0x06, 0x63, 0xdc,
	0xa9, 0xd2, 0xce, 0x6a, 0x2a, 0x0d, 0x65, 0x29, 0x68, 0x81, 0x3c, 0x42, 0xc7, 0x94, 0x31, 0x88,
	0xac, 0xfb, 0x12, 0x76, 0x91, 0x3c, 0x46, 0x0f, 0x38, 0xb0, 0x40, 0x48, 0xf8, 0x22, 0xf8, 0x16,
	0xd9, 0x46, 0x77, 0x0a, 0xd0, 0x78, 0xe2, 0x36, 0xd9, 0x40, 0xd8, 0x80, 0xe4, 0x57, 0xa2, 0x88,
	0x1c, 0xa1, 0xbd, 0xc9, 0xda, 0xe3, 0x80, 0x25, 0x4f, 0xcd, 0xd4, 0x90, 0x2e, 0x27, 0x10, 0x2f,
	0xcf, 0x4e, 0x53, 0xc6, 0x54, 0x2c, 0x2d, 0x5e, 0x21, 0x77, 0xd1, 0xc1, 0x74, 0x3a, 0x8a, 0x5b,
	0x81, 0x60, 0xce, 0xef, 0x05, 0xaf, 0x92, 0x43, 0xb4, 0x5b, 0xec, 0x83, 0x29, 0x0e, 0x8e, 0xf2,
	0x1e, 0x68, 0x2b, 0x0c, 0x84, 0x20, 0x2d, 0x5e, 0x23, 0x0d, 0x74, 0x18, 0xc5, 0xa6, 0xe3, 0xa4,
	0xb2, 0xe2, 0x54, 0xb0, 0xac, 0x84, 0x86, 0xb6, 0x30, 0x56, 0x67, 0x94, 0x63, 0xcf, 0xd0, 0xff,
	0x63, 0x9c, 0x06, 0x13, 0x29, 0x69, 0x00, 0xaf, 0x93, 0x3d, 0xb4, 0x3d, 0x0d, 0x7e, 0x15, 0x83,
	0xee, 0x63, 0x42, 0xee, 0xa3, 0xfa, 0x35, 0xc9, 0xaa, 0xc4, 0x1d, 0x3f, 0xf5, 0xac, 0x7e, 0x29,
	0x7f, 0x78, 0xc3, 0x8f, 0x34, 0x2b, 0x9d, 0x1f, 0xdf, 0xf4, 0x12, 0x84, 0x50, 0xbd, 0x14, 0x4e,
	0x43, 0xce, 0xf3, 0x16, 0xd9, 0x41, 0x9b, 0x6d, 0xad, 0xe2, 0x28, 0xa5, 0xc5, 0x09, 0xd9, 0x13,
	0x36, 0x9b, 0x6e, 0x9b, 0xac, 0xa3, 0x95, 0x2c, 0xc8, 0x41, 0x5a, 0x61, 0xfb, 0xb8, 0xe6, 0xd1,
	0x4c, 0x85, 0x61, 0x2c, 0x85, 0xed, 0x3b, 0x0e, 0x86, 0x69, 0x11, 0xa5, 0xe8, 0x1d, 0x52, 0x43,
	0x1b, 0x55, 0x6a, 0xac, 0xce, 0xae, 0xbf, 0x75, 0x95, 0x29, 0xb7, 0xad, 0xdc, 0x4b, 0x25, 0x24,
	0xde, 0x23, 0x6b, 0x68, 0x29, 0x12, 0xb2, 0x94, 0xfd, 0xbe, 0xf7, 0x0e, 0x70, 0x51, 0x79, 0xe7,
	0xc0, 0xdf, 0xc4, 0x58, 0x6a, 0x63, 0x53, 0x58, 0xe7, 0xd0, 0xcf, 0xc2, 0x21, 0x80, 0x31, 0xbf,
	0x1c, 0x79, 0x51, 0xcd, 0xd2, 0x4c, 0xde, 0x1a, 0xd7, 0xc9, 0x2e, 0xda, 0xa2, 0x52, 0xc9, 0x7e,
	0xa8, 0x62, 0xe3, 0x42, 0xb0, 0x5a, 0x30, 0xd7, 0xa2, 0x96, 0x75, 0xf0, 0xdd, 0xd2, 0x55, 0xe9,
	0xc8, 0x1a, 0x42, 0xd5, 0x03, 0x8e, 0x1b, 0x7e, 0x6b, 0x55, 0x38, 0x6f, 0x65, 0x3c, 0x81, 0x1c,
	0xdf, 0x23, 0x08, 0x2d, 0xb4, 0x28, 0xeb, 0xc6, 0x11, 0xbe, 0x5f, 0x2a, 0xd2, 0x33, 0xdb, 0xf3,
	0x93, 0x32, 0x90, 0x16, 0x74, 0x06, 0xfd, 0xa6, 0x54, 0xe4, 0x64, 0x3a, 0x73, 0x23, 0x70, 0x7c,
	0xec, 0x15, 0x37, 0x13, 0xc2, 0x85, 0x09, 0x85, 0x31, 0xc0, 0xf1, 0x83, 0x94, 0x09, 0x8f, 0x69,
	0x29, 0xd5, 0x0d, 0xa9, 0xee, 0xe2, 0x13, 0xb2, 0x85, 0x48, 0x76, 0xc3, 0x00, 0xa8, 0x76, 0x1d,
	0x61, 0xac, 0xd2, 0x7d, 0xfc, 0xd0, 0xd3, 0x98, 0xc6, 0x0d, 0x58, 0x2b, 0x64, 0x1b, 0x3f, 0x22,
	0x75, 0xb4, 0x5f, 0x2d, 0x82, 0x6a, 0xd6, 0x11, 0x3d, 0x70, 0x21, 0x6d, 0x4b, 0xb0, 0x81, 0x90,
	0x5d, 0xfc, 0xd8, 0x2f, 0x31, 0x3d, 0x13, 0x69, 0x75, 0x2a, 0x02, 0x70, 0x91, 0x60, 0x36, 0xd6,
	0x80, 0xbf, 0xf5, 0xfe, 0x4e, 0x33, 0xaf, 0x69, 0x10, 0x80, 0x2d, 0xad, 0xf6, 0x5d, 0xca, 0x69,
	0xf6, 0xa2, 0x14, 0x76, 0x2a, 0x04, 0xd9, 0xf4, 0xe4, 0x69, 0xb0, 0x3a, 0xf3, 0xd8, 0xd5, 0xe4,
	0x13, 0x72, 0x8c, 0x1a, 0xd7, 0xca, 0xa2, 0x52, 0xed, 0xf7, 0xd5, 0x06, 0x4a, 0x70, 0x3e, 0x91,
	0xc1, 0x3f, 0xf8, 0x91, 0x8a, 0xa3, 0x45, 0x87, 0x1e, 0xe8, 0x52, 0xfd, 0xf8, 0xa9, 0x17, 0xc5,
	0xc4, 0xfd, 0xae, 0x00, 0x9e, 0xf9, 0x12, 0xc5, 0x53, 0x34, 0x13, 0xf1, 0x63, 0x29, 0x0d, 0xab,
	0x63, 0x63, 0x81, 0xbb, 0xd8, 0x80, 0xc6, 0x3f, 0x95, 0x1b, 0x1f, 0x47, 0x97, 0xf3, 0xfd, 0x9c,
	0xb1, 0x3d, 0x5d, 0xaf, 0xa8, 0x82, 0x7f, 0x29, 0x35, 0x31, 0xc1, 0x8d, 0xe3, 0xc0, 0x84, 0xf1,
	0xad, 0x7f, 0xcd, 0x5e, 0xa9, 0x19, 0x24, 0x05, 0x40, 0x7b, 0x80, 0x7f, 0xf3, 0xf9, 0xb4, 0x44,
	0xee, 0x05, 0xff, 0x2e, 0x87, 0x95, 0x25, 0x7e, 0x2f, 0xc5, 0x61, 0x68, 0x0f, 0x78, 0xf1, 0x7c,
	0xe3, 0xe7, 0xfe, 0xbd, 0xa9, 0xea, 0x32, 0x2a, 0x19, 0x04, 0x53, 0xd6, 0xfc, 0xc3, 0x73, 0x97,
	0xe7, 0x66, 0x32, 0xf3, 0xa2, 0xb5, 0xf2, 0xe7, 0x52, 0xf3, 0xc9, 0xf3, 0xe2, 0x17, 0xf3, 0xcd,
	0x42, 0xfa, 0xdf, 0xb3, 0xff, 0x02, 0x00, 0x00, 0xff, 0xff, 0xc3, 0x1d, 0x3f, 0x82, 0xd8, 0x07,
	0x00, 0x00,
}
