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
	ApplicationMetadataMessage_SYNC_ACCOUNT                            ApplicationMetadataMessage_Type = 45
	ApplicationMetadataMessage_ACCEPT_CONTACT_REQUEST                  ApplicationMetadataMessage_Type = 46
	ApplicationMetadataMessage_RETRACT_CONTACT_REQUEST                 ApplicationMetadataMessage_Type = 47
	ApplicationMetadataMessage_COMMUNITY_REQUEST_TO_JOIN_RESPONSE      ApplicationMetadataMessage_Type = 48
	ApplicationMetadataMessage_SYNC_COMMUNITY_SETTINGS                 ApplicationMetadataMessage_Type = 49
	ApplicationMetadataMessage_REQUEST_CONTACT_VERIFICATION            ApplicationMetadataMessage_Type = 50
	ApplicationMetadataMessage_ACCEPT_CONTACT_VERIFICATION             ApplicationMetadataMessage_Type = 51
	ApplicationMetadataMessage_DECLINE_CONTACT_VERIFICATION            ApplicationMetadataMessage_Type = 52
	ApplicationMetadataMessage_SYNC_TRUSTED_USER                       ApplicationMetadataMessage_Type = 53
	ApplicationMetadataMessage_SYNC_VERIFICATION_REQUEST               ApplicationMetadataMessage_Type = 54
	ApplicationMetadataMessage_SYNC_CONTACT_REQUEST_DECISION           ApplicationMetadataMessage_Type = 56
	ApplicationMetadataMessage_COMMUNITY_REQUEST_TO_LEAVE              ApplicationMetadataMessage_Type = 57
	ApplicationMetadataMessage_SYNC_DELETE_FOR_ME_MESSAGE              ApplicationMetadataMessage_Type = 58
	ApplicationMetadataMessage_SYNC_SAVED_ADDRESS                      ApplicationMetadataMessage_Type = 59
	ApplicationMetadataMessage_COMMUNITY_CANCEL_REQUEST_TO_JOIN        ApplicationMetadataMessage_Type = 60
	ApplicationMetadataMessage_CANCEL_CONTACT_VERIFICATION             ApplicationMetadataMessage_Type = 61
	ApplicationMetadataMessage_SYNC_KEYPAIR                            ApplicationMetadataMessage_Type = 62
	ApplicationMetadataMessage_SYNC_SOCIAL_LINKS                       ApplicationMetadataMessage_Type = 63
	ApplicationMetadataMessage_SYNC_ENS_USERNAME_DETAIL                ApplicationMetadataMessage_Type = 64
	ApplicationMetadataMessage_SYNC_ACTIVITY_CENTER_NOTIFICATION       ApplicationMetadataMessage_Type = 65
	ApplicationMetadataMessage_SYNC_ACTIVITY_CENTER_NOTIFICATION_STATE ApplicationMetadataMessage_Type = 66
	ApplicationMetadataMessage_COMMUNITY_ADMIN_MESSAGE                 ApplicationMetadataMessage_Type = 67
	ApplicationMetadataMessage_COMMUNITY_EDIT_SHARED_ADDRESSES         ApplicationMetadataMessage_Type = 68
	ApplicationMetadataMessage_SYNC_ACCOUNT_CUSTOMIZATION_COLOR        ApplicationMetadataMessage_Type = 69
	ApplicationMetadataMessage_SYNC_ACCOUNTS_POSITIONS                 ApplicationMetadataMessage_Type = 70
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
	45: "SYNC_ACCOUNT",
	46: "ACCEPT_CONTACT_REQUEST",
	47: "RETRACT_CONTACT_REQUEST",
	48: "COMMUNITY_REQUEST_TO_JOIN_RESPONSE",
	49: "SYNC_COMMUNITY_SETTINGS",
	50: "REQUEST_CONTACT_VERIFICATION",
	51: "ACCEPT_CONTACT_VERIFICATION",
	52: "DECLINE_CONTACT_VERIFICATION",
	53: "SYNC_TRUSTED_USER",
	54: "SYNC_VERIFICATION_REQUEST",
	56: "SYNC_CONTACT_REQUEST_DECISION",
	57: "COMMUNITY_REQUEST_TO_LEAVE",
	58: "SYNC_DELETE_FOR_ME_MESSAGE",
	59: "SYNC_SAVED_ADDRESS",
	60: "COMMUNITY_CANCEL_REQUEST_TO_JOIN",
	61: "CANCEL_CONTACT_VERIFICATION",
	62: "SYNC_KEYPAIR",
	63: "SYNC_SOCIAL_LINKS",
	64: "SYNC_ENS_USERNAME_DETAIL",
	65: "SYNC_ACTIVITY_CENTER_NOTIFICATION",
	66: "SYNC_ACTIVITY_CENTER_NOTIFICATION_STATE",
	67: "COMMUNITY_ADMIN_MESSAGE",
	68: "COMMUNITY_EDIT_SHARED_ADDRESSES",
	69: "SYNC_ACCOUNT_CUSTOMIZATION_COLOR",
	70: "SYNC_ACCOUNTS_POSITIONS",
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
	"SYNC_ACCOUNT":                            45,
	"ACCEPT_CONTACT_REQUEST":                  46,
	"RETRACT_CONTACT_REQUEST":                 47,
	"COMMUNITY_REQUEST_TO_JOIN_RESPONSE":      48,
	"SYNC_COMMUNITY_SETTINGS":                 49,
	"REQUEST_CONTACT_VERIFICATION":            50,
	"ACCEPT_CONTACT_VERIFICATION":             51,
	"DECLINE_CONTACT_VERIFICATION":            52,
	"SYNC_TRUSTED_USER":                       53,
	"SYNC_VERIFICATION_REQUEST":               54,
	"SYNC_CONTACT_REQUEST_DECISION":           56,
	"COMMUNITY_REQUEST_TO_LEAVE":              57,
	"SYNC_DELETE_FOR_ME_MESSAGE":              58,
	"SYNC_SAVED_ADDRESS":                      59,
	"COMMUNITY_CANCEL_REQUEST_TO_JOIN":        60,
	"CANCEL_CONTACT_VERIFICATION":             61,
	"SYNC_KEYPAIR":                            62,
	"SYNC_SOCIAL_LINKS":                       63,
	"SYNC_ENS_USERNAME_DETAIL":                64,
	"SYNC_ACTIVITY_CENTER_NOTIFICATION":       65,
	"SYNC_ACTIVITY_CENTER_NOTIFICATION_STATE": 66,
	"COMMUNITY_ADMIN_MESSAGE":                 67,
	"COMMUNITY_EDIT_SHARED_ADDRESSES":         68,
	"SYNC_ACCOUNT_CUSTOMIZATION_COLOR":        69,
	"SYNC_ACCOUNTS_POSITIONS":                 70,
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
	// 1010 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x84, 0x56, 0x6d, 0x73, 0x13, 0x37,
	0x10, 0x6e, 0x20, 0x4d, 0x40, 0x79, 0xdb, 0x88, 0xbc, 0x38, 0xef, 0x89, 0x81, 0x10, 0xa0, 0x35,
	0x2d, 0xb4, 0x9d, 0xb6, 0x94, 0xb6, 0xb2, 0xb4, 0x89, 0x85, 0xef, 0xa4, 0x43, 0xd2, 0xb9, 0x63,
	0xbe, 0x68, 0x4c, 0x71, 0x99, 0xcc, 0x00, 0xf1, 0x10, 0xf3, 0x21, 0xbf, 0xa0, 0xbf, 0xb7, 0xff,
	0xa0, 0xa3, 0xb3, 0xef, 0xce, 0x49, 0x1c, 0xf2, 0x29, 0xf1, 0x3e, 0x8f, 0x56, 0xda, 0x67, 0x9f,
	0x5d, 0x9b, 0x54, 0x3b, 0xbd, 0xde, 0xfb, 0xe3, 0xbf, 0x3b, 0xfd, 0xe3, 0x93, 0x8f, 0xfe, 0x43,
	0xb7, 0xdf, 0x79, 0xdb, 0xe9, 0x77, 0xfc, 0x87, 0xee, 0xe9, 0x69, 0xe7, 0x5d, 0xb7, 0xd6, 0xfb,
	0x74, 0xd2, 0x3f, 0xa1, 0xb7, 0xb2, 0x3f, 0x6f, 0x3e, 0xff, 0x53, 0xfd, 0x77, 0x91, 0xac, 0xb3,
	0xf2, 0x40, 0x3c, 0xe4, 0xc7, 0x03, 0x3a, 0xdd, 0x24, 0xb7, 0x4f, 0x8f, 0xdf, 0x7d, 0xec, 0xf4,
	0x3f, 0x7f, 0xea, 0x56, 0x26, 0x76, 0x27, 0x0e, 0x66, 0x4d, 0x19, 0xa0, 0x15, 0x32, 0xdd, 0xeb,
	0x9c, 0xbd, 0x3f, 0xe9, 0xbc, 0xad, 0xdc, 0xc8, 0xb0, 0xfc, 0x23, 0x7d, 0x41, 0x26, 0xfb, 0x67,
	0xbd, 0x6e, 0xe5, 0xe6, 0xee, 0xc4, 0xc1, 0xfc, 0xd3, 0x87, 0xb5, 0xfc, 0xbe, 0xda, 0xd5, 0x77,
	0xd5, 0xdc, 0x59, 0xaf, 0x6b, 0xb2, 0x63, 0xd5, 0xff, 0x16, 0xc8, 0x64, 0xf8, 0x48, 0x67, 0xc8,
	0x74, 0xaa, 0x9a, 0x4a, 0xff, 0xa5, 0xe0, 0x2b, 0x0a, 0x64, 0x96, 0x37, 0x98, 0xf3, 0x31, 0x5a,
	0xcb, 0x8e, 0x10, 0x26, 0x28, 0x25, 0xf3, 0x5c, 0x2b, 0xc7, 0xb8, 0xf3, 0x69, 0x22, 0x98, 0x43,
	0xb8, 0x41, 0xb7, 0xc8, 0x5a, 0x8c, 0x71, 0x1d, 0x8d, 0x6d, 0xc8, 0x64, 0x18, 0x2e, 0x8e, 0xdc,
	0xa4, 0xcb, 0x64, 0x31, 0x61, 0xd2, 0x78, 0xa9, 0xac, 0x63, 0x51, 0xc4, 0x9c, 0xd4, 0x0a, 0x26,
	0x43, 0xd8, 0xb6, 0x15, 0x3f, 0x1f, 0xfe, 0x9a, 0xde, 0x25, 0x3b, 0x06, 0x5f, 0xa5, 0x68, 0x9d,
	0x67, 0x42, 0x18, 0xb4, 0xd6, 0x1f, 0x6a, 0xe3, 0x9d, 0x61, 0xca, 0x32, 0x9e, 0x91, 0xa6, 0xe8,
	0x23, 0xb2, 0xcf, 0x38, 0xc7, 0xc4, 0xf9, 0xeb, 0xb8, 0xd3, 0xf4, 0x31, 0x79, 0x20, 0x90, 0x47,
	0x52, 0xe1, 0xb5, 0xe4, 0x5b, 0x74, 0x95, 0xdc, 0xc9, 0x49, 0xa3, 0xc0, 0x6d, 0xba, 0x44, 0xc0,
	0xa2, 0x12, 0xe7, 0xa2, 0x84, 0xee, 0x90, 0x8d, 0x8b, 0xb9, 0x47, 0x09, 0x33, 0x41, 0x9a, 0x4b,
	0x45, 0xfa, 0xa1, 0x80, 0x30, 0x3b, 0x1e, 0x66, 0x9c, 0xeb, 0x54, 0x39, 0x98, 0xa3, 0x7b, 0x64,
	0xeb, 0x32, 0x9c, 0xa4, 0xf5, 0x48, 0x72, 0x1f, 0xfa, 0x02, 0xf3, 0x74, 0x9b, 0xac, 0xe7, 0xfd,
	0xe0, 0x5a, 0xa0, 0x67, 0xa2, 0x85, 0xc6, 0x49, 0x8b, 0x31, 0x2a, 0x07, 0x0b, 0xb4, 0x4a, 0xb6,
	0x93, 0xd4, 0x36, 0xbc, 0xd2, 0x4e, 0x1e, 0x4a, 0x3e, 0x48, 0x61, 0xf0, 0x48, 0x5a, 0x67, 0x06,
	0x92, 0x43, 0x50, 0xe8, 0xcb, 0x1c, 0x6f, 0xd0, 0x26, 0x5a, 0x59, 0x84, 0x45, 0xba, 0x41, 0x56,
	0x2f, 0x93, 0x5f, 0xa5, 0x68, 0xda, 0x40, 0xe9, 0x3d, 0xb2, 0x7b, 0x05, 0x58, 0xa6, 0xb8, 0x13,
	0xaa, 0x1e, 0x77, 0x5f, 0xa6, 0x1f, 0x2c, 0x85, 0x92, 0xc6, 0xc1, 0xc3, 0xe3, 0xcb, 0xc1, 0x82,
	0x18, 0xeb, 0x97, 0xd2, 0x1b, 0x1c, 0xea, 0xbc, 0x42, 0xd7, 0xc8, 0xf2, 0x91, 0xd1, 0x69, 0x92,
	0xc9, 0xe2, 0xa5, 0x6a, 0x49, 0x37, 0xa8, 0x6e, 0x95, 0x2e, 0x92, 0xb9, 0x41, 0x50, 0xa0, 0x72,
	0xd2, 0xb5, 0xa1, 0x12, 0xd8, 0x5c, 0xc7, 0x71, 0xaa, 0xa4, 0x6b, 0x7b, 0x81, 0x96, 0x1b, 0x99,
	0x64, 0xec, 0x35, 0x5a, 0x21, 0x4b, 0x25, 0x34, 0x92, 0x67, 0x3d, 0xbc, 0xba, 0x44, 0x8a, 0x6e,
	0x6b, 0xff, 0x52, 0x4b, 0x05, 0x1b, 0x74, 0x81, 0xcc, 0x24, 0x52, 0x15, 0xb6, 0xdf, 0x0c, 0xb3,
	0x83, 0x42, 0x96, 0xb3, 0xb3, 0x15, 0x5e, 0x62, 0x1d, 0x73, 0xa9, 0xcd, 0x47, 0x67, 0x3b, 0xd4,
	0x22, 0x30, 0xc2, 0x91, 0x79, 0xd9, 0x09, 0xa6, 0x1a, 0xe7, 0x99, 0xe1, 0xd5, 0xb0, 0x4b, 0xd7,
	0xc9, 0x0a, 0x53, 0x5a, 0xb5, 0x63, 0x9d, 0x5a, 0x1f, 0xa3, 0x33, 0x92, 0xfb, 0x3a, 0x73, 0xbc,
	0x01, 0x7b, 0xc5, 0x54, 0x65, 0x25, 0x1b, 0x8c, 0x75, 0x0b, 0x05, 0x54, 0x43, 0xd7, 0xca, 0xf0,
	0xf0, 0x2a, 0x1b, 0x04, 0x14, 0x70, 0x97, 0x12, 0x32, 0x55, 0x67, 0xbc, 0x99, 0x26, 0x70, 0xaf,
	0x70, 0x64, 0x50, 0xb6, 0x15, 0x2a, 0xe5, 0xa8, 0x1c, 0x9a, 0x01, 0xf5, 0x7e, 0xe1, 0xc8, 0x8b,
	0xf0, 0x60, 0x1a, 0x51, 0xc0, 0x7e, 0x70, 0xdc, 0x58, 0x8a, 0x90, 0x36, 0x96, 0xd6, 0xa2, 0x80,
	0x07, 0x99, 0x12, 0x81, 0x53, 0xd7, 0xba, 0x19, 0x33, 0xd3, 0x84, 0x03, 0xba, 0x42, 0xe8, 0xe0,
	0x85, 0x11, 0x32, 0xe3, 0x1b, 0xd2, 0x3a, 0x6d, 0xda, 0xf0, 0x30, 0xc8, 0x98, 0xc5, 0x2d, 0x3a,
	0x27, 0xd5, 0x11, 0x3c, 0xa2, 0xbb, 0x64, 0xb3, 0x6c, 0x04, 0x33, 0xbc, 0x21, 0x5b, 0xe8, 0x63,
	0x76, 0xa4, 0xd0, 0x45, 0x52, 0x35, 0xe1, 0x71, 0x68, 0x62, 0x76, 0x26, 0x31, 0xfa, 0x50, 0x46,
	0xe8, 0x13, 0xc9, 0x5d, 0x6a, 0x10, 0xbe, 0x29, 0xb2, 0xe5, 0x33, 0xf6, 0x6d, 0x26, 0xe6, 0x60,
	0x95, 0xe4, 0x73, 0x94, 0x3b, 0xb1, 0x16, 0x54, 0x33, 0xe8, 0xcc, 0x60, 0xb8, 0xce, 0x83, 0x4f,
	0xe8, 0x3e, 0xa9, 0x5e, 0xe9, 0x87, 0xd2, 0xae, 0xdf, 0x95, 0xd2, 0x17, 0xe4, 0x61, 0x29, 0x16,
	0xbe, 0x0f, 0xb5, 0xe4, 0x47, 0xf3, 0x1b, 0x5a, 0x68, 0x0a, 0xdb, 0xc3, 0xd3, 0xe0, 0x86, 0x0b,
	0xef, 0x3b, 0x47, 0x78, 0x16, 0x52, 0xe4, 0x3b, 0x68, 0x2c, 0xe3, 0x87, 0xc2, 0x13, 0xce, 0xa4,
	0xd6, 0xa1, 0xf0, 0xa9, 0x45, 0x03, 0x3f, 0x16, 0xad, 0x1e, 0x65, 0x17, 0xf5, 0xfd, 0x54, 0xb4,
	0xfa, 0x42, 0xe5, 0x5e, 0x20, 0x97, 0x36, 0x24, 0xfe, 0x79, 0xb0, 0x7c, 0xc6, 0x48, 0x10, 0x21,
	0x6b, 0x21, 0xfc, 0x12, 0xf0, 0x2c, 0xc5, 0xd0, 0xe2, 0x61, 0xdd, 0xc6, 0xa5, 0xd3, 0x7f, 0x2d,
	0x7a, 0x6e, 0x59, 0x0b, 0x45, 0xbe, 0x95, 0xe1, 0x79, 0x58, 0x23, 0x65, 0x5e, 0xce, 0x14, 0xc7,
	0xe8, 0xd2, 0xc4, 0xfd, 0x16, 0x94, 0x19, 0x62, 0x63, 0xeb, 0x7e, 0x51, 0x34, 0xbb, 0x89, 0xed,
	0xf0, 0x05, 0x04, 0xbf, 0x17, 0x4a, 0x58, 0xcd, 0x25, 0x8b, 0x7c, 0xb0, 0x8b, 0x85, 0x3f, 0xe8,
	0x26, 0xa9, 0x64, 0x61, 0x54, 0x36, 0x13, 0x47, 0xb1, 0x18, 0xbd, 0x40, 0xc7, 0x64, 0x04, 0x7f,
	0xd2, 0xfb, 0x64, 0x6f, 0xac, 0xa1, 0x47, 0xf7, 0x13, 0xb0, 0xb0, 0x45, 0xaf, 0xa5, 0xf9, 0x30,
	0xff, 0x08, 0xf5, 0x60, 0x8a, 0x11, 0x0f, 0x8b, 0x78, 0x64, 0x73, 0xf0, 0xf0, 0x15, 0x58, 0x82,
	0xd9, 0x0e, 0xb1, 0x0d, 0x66, 0x4a, 0x85, 0xd0, 0x82, 0x08, 0x1a, 0x8d, 0x3a, 0xd9, 0xf3, 0xd4,
	0x3a, 0x1d, 0xcb, 0xd7, 0xf9, 0xba, 0x88, 0xb4, 0x01, 0x2c, 0xcc, 0x37, 0x64, 0x59, 0x9f, 0x68,
	0x2b, 0x03, 0xc3, 0xc2, 0x61, 0x7d, 0xee, 0xf5, 0x4c, 0xed, 0xc9, 0xf3, 0xfc, 0x87, 0xc2, 0x9b,
	0xa9, 0xec, 0xbf, 0x67, 0xff, 0x07, 0x00, 0x00, 0xff, 0xff, 0xd4, 0x9b, 0xda, 0x33, 0xcf, 0x08,
	0x00, 0x00,
}
