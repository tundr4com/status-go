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
	ApplicationMetadataMessage_SYNC_PAIR_INSTALLATION                  ApplicationMetadataMessage_Type = 4
	ApplicationMetadataMessage_DEPRECATED_SYNC_INSTALLATION            ApplicationMetadataMessage_Type = 5 // Deprecated: Do not use.
	ApplicationMetadataMessage_REQUEST_ADDRESS_FOR_TRANSACTION         ApplicationMetadataMessage_Type = 6
	ApplicationMetadataMessage_ACCEPT_REQUEST_ADDRESS_FOR_TRANSACTION  ApplicationMetadataMessage_Type = 7
	ApplicationMetadataMessage_DECLINE_REQUEST_ADDRESS_FOR_TRANSACTION ApplicationMetadataMessage_Type = 8
	ApplicationMetadataMessage_REQUEST_TRANSACTION                     ApplicationMetadataMessage_Type = 9
	ApplicationMetadataMessage_SEND_TRANSACTION                        ApplicationMetadataMessage_Type = 10
	ApplicationMetadataMessage_DECLINE_REQUEST_TRANSACTION             ApplicationMetadataMessage_Type = 11
	ApplicationMetadataMessage_SYNC_INSTALLATION_CONTACT_V2            ApplicationMetadataMessage_Type = 12
	ApplicationMetadataMessage_SYNC_INSTALLATION_ACCOUNT               ApplicationMetadataMessage_Type = 13
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
	ApplicationMetadataMessage_COMMUNITY_INVITATION                    ApplicationMetadataMessage_Type = 26 // Deprecated: Do not use.
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
	ApplicationMetadataMessage_COMMUNITY_MESSAGE_ARCHIVE_MAGNETLINK    ApplicationMetadataMessage_Type = 43
	ApplicationMetadataMessage_SYNC_PROFILE_PICTURES                   ApplicationMetadataMessage_Type = 44
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
	ApplicationMetadataMessage_SYNC_ACTIVITY_CENTER_NOTIFICATIONS      ApplicationMetadataMessage_Type = 65
	ApplicationMetadataMessage_SYNC_ACTIVITY_CENTER_NOTIFICATION_STATE ApplicationMetadataMessage_Type = 66
	ApplicationMetadataMessage_COMMUNITY_EVENTS_MESSAGE                ApplicationMetadataMessage_Type = 67
	ApplicationMetadataMessage_COMMUNITY_EDIT_SHARED_ADDRESSES         ApplicationMetadataMessage_Type = 68
	ApplicationMetadataMessage_SYNC_ACCOUNT_CUSTOMIZATION_COLOR        ApplicationMetadataMessage_Type = 69
	ApplicationMetadataMessage_SYNC_ACCOUNTS_POSITIONS                 ApplicationMetadataMessage_Type = 70
	ApplicationMetadataMessage_COMMUNITY_EVENTS_MESSAGE_REJECTED       ApplicationMetadataMessage_Type = 71
	ApplicationMetadataMessage_COMMUNITY_PRIVILEGED_USER_SYNC_MESSAGE  ApplicationMetadataMessage_Type = 72
	ApplicationMetadataMessage_COMMUNITY_SHARD_KEY                     ApplicationMetadataMessage_Type = 73
	ApplicationMetadataMessage_SYNC_CHAT                               ApplicationMetadataMessage_Type = 74
)

var ApplicationMetadataMessage_Type_name = map[int32]string{
	0:  "UNKNOWN",
	1:  "CHAT_MESSAGE",
	2:  "CONTACT_UPDATE",
	3:  "MEMBERSHIP_UPDATE_MESSAGE",
	4:  "SYNC_PAIR_INSTALLATION",
	5:  "DEPRECATED_SYNC_INSTALLATION",
	6:  "REQUEST_ADDRESS_FOR_TRANSACTION",
	7:  "ACCEPT_REQUEST_ADDRESS_FOR_TRANSACTION",
	8:  "DECLINE_REQUEST_ADDRESS_FOR_TRANSACTION",
	9:  "REQUEST_TRANSACTION",
	10: "SEND_TRANSACTION",
	11: "DECLINE_REQUEST_TRANSACTION",
	12: "SYNC_INSTALLATION_CONTACT_V2",
	13: "SYNC_INSTALLATION_ACCOUNT",
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
	43: "COMMUNITY_MESSAGE_ARCHIVE_MAGNETLINK",
	44: "SYNC_PROFILE_PICTURES",
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
	65: "SYNC_ACTIVITY_CENTER_NOTIFICATIONS",
	66: "SYNC_ACTIVITY_CENTER_NOTIFICATION_STATE",
	67: "COMMUNITY_EVENTS_MESSAGE",
	68: "COMMUNITY_EDIT_SHARED_ADDRESSES",
	69: "SYNC_ACCOUNT_CUSTOMIZATION_COLOR",
	70: "SYNC_ACCOUNTS_POSITIONS",
	71: "COMMUNITY_EVENTS_MESSAGE_REJECTED",
	72: "COMMUNITY_PRIVILEGED_USER_SYNC_MESSAGE",
	73: "COMMUNITY_SHARD_KEY",
	74: "SYNC_CHAT",
}

var ApplicationMetadataMessage_Type_value = map[string]int32{
	"UNKNOWN":                                 0,
	"CHAT_MESSAGE":                            1,
	"CONTACT_UPDATE":                          2,
	"MEMBERSHIP_UPDATE_MESSAGE":               3,
	"SYNC_PAIR_INSTALLATION":                  4,
	"DEPRECATED_SYNC_INSTALLATION":            5,
	"REQUEST_ADDRESS_FOR_TRANSACTION":         6,
	"ACCEPT_REQUEST_ADDRESS_FOR_TRANSACTION":  7,
	"DECLINE_REQUEST_ADDRESS_FOR_TRANSACTION": 8,
	"REQUEST_TRANSACTION":                     9,
	"SEND_TRANSACTION":                        10,
	"DECLINE_REQUEST_TRANSACTION":             11,
	"SYNC_INSTALLATION_CONTACT_V2":            12,
	"SYNC_INSTALLATION_ACCOUNT":               13,
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
	"COMMUNITY_MESSAGE_ARCHIVE_MAGNETLINK":    43,
	"SYNC_PROFILE_PICTURES":                   44,
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
	"SYNC_ACTIVITY_CENTER_NOTIFICATIONS":      65,
	"SYNC_ACTIVITY_CENTER_NOTIFICATION_STATE": 66,
	"COMMUNITY_EVENTS_MESSAGE":                67,
	"COMMUNITY_EDIT_SHARED_ADDRESSES":         68,
	"SYNC_ACCOUNT_CUSTOMIZATION_COLOR":        69,
	"SYNC_ACCOUNTS_POSITIONS":                 70,
	"COMMUNITY_EVENTS_MESSAGE_REJECTED":       71,
	"COMMUNITY_PRIVILEGED_USER_SYNC_MESSAGE":  72,
	"COMMUNITY_SHARD_KEY":                     73,
	"SYNC_CHAT":                               74,
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
	// 1088 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x84, 0x56, 0xeb, 0x52, 0x5b, 0x37,
	0x17, 0xfd, 0x48, 0xf8, 0xb8, 0x88, 0x4b, 0x84, 0xc2, 0xc5, 0x80, 0xb9, 0x39, 0x84, 0x10, 0xd2,
	0x3a, 0x2d, 0x69, 0x3b, 0x6d, 0xd3, 0xb4, 0x95, 0xa5, 0x8d, 0x2d, 0x7c, 0x8e, 0x74, 0x22, 0xe9,
	0xb8, 0xe3, 0xfc, 0xd1, 0x38, 0x8d, 0x9b, 0x61, 0x26, 0x09, 0x9e, 0xe0, 0xfc, 0xe0, 0x3d, 0xfa,
	0x14, 0x7d, 0xca, 0x8e, 0xce, 0xd5, 0x80, 0x29, 0xbf, 0xe0, 0x68, 0x2f, 0x6d, 0x69, 0xaf, 0xbd,
	0xf6, 0x92, 0x51, 0xad, 0x37, 0x18, 0x7c, 0x38, 0xfb, 0xb3, 0x37, 0x3c, 0x3b, 0xff, 0xe4, 0x3e,
	0xf6, 0x87, 0xbd, 0x77, 0xbd, 0x61, 0xcf, 0x7d, 0xec, 0x5f, 0x5c, 0xf4, 0xde, 0xf7, 0xeb, 0x83,
	0xcf, 0xe7, 0xc3, 0x73, 0x32, 0x93, 0xfc, 0x79, 0xfb, 0xe5, 0xaf, 0xda, 0x3f, 0x04, 0x6d, 0xd0,
	0x72, 0x43, 0x98, 0xe1, 0xc3, 0x14, 0x4e, 0xaa, 0x68, 0xf6, 0xe2, 0xec, 0xfd, 0xa7, 0xde, 0xf0,
	0xcb, 0xe7, 0x7e, 0x65, 0x62, 0x77, 0xe2, 0x70, 0x5e, 0x97, 0x0b, 0xa4, 0x82, 0xa6, 0x07, 0xbd,
	0xcb, 0x0f, 0xe7, 0xbd, 0x77, 0x95, 0x7b, 0x49, 0x2c, 0xff, 0x24, 0xaf, 0xd0, 0xe4, 0xf0, 0x72,
	0xd0, 0xaf, 0xdc, 0xdf, 0x9d, 0x38, 0x5c, 0x3c, 0x7e, 0x5a, 0xcf, 0xcf, 0xab, 0xdf, 0x7e, 0x56,
	0xdd, 0x5e, 0x0e, 0xfa, 0x3a, 0xd9, 0x56, 0xfb, 0x7b, 0x09, 0x4d, 0xfa, 0x4f, 0x32, 0x87, 0xa6,
	0x63, 0xd9, 0x96, 0xea, 0x0f, 0x89, 0xff, 0x47, 0x30, 0x9a, 0x67, 0x2d, 0x6a, 0x5d, 0x08, 0xc6,
	0xd0, 0x26, 0xe0, 0x09, 0x42, 0xd0, 0x22, 0x53, 0xd2, 0x52, 0x66, 0x5d, 0x1c, 0x71, 0x6a, 0x01,
	0xdf, 0x23, 0x5b, 0x68, 0x3d, 0x84, 0xb0, 0x01, 0xda, 0xb4, 0x44, 0x94, 0x2d, 0x17, 0x5b, 0xee,
	0x93, 0x0d, 0xb4, 0x6a, 0xba, 0x92, 0xb9, 0x88, 0x0a, 0xed, 0x84, 0x34, 0x96, 0x06, 0x01, 0xb5,
	0x42, 0x49, 0x3c, 0x49, 0xf6, 0x51, 0x95, 0x43, 0xa4, 0x81, 0x51, 0x0b, 0xdc, 0x25, 0xb0, 0x2b,
	0x88, 0xff, 0x6f, 0xdc, 0x9b, 0x99, 0x20, 0x8f, 0xd0, 0x8e, 0x86, 0xd7, 0x31, 0x18, 0xeb, 0x28,
	0xe7, 0x1a, 0x8c, 0x71, 0x27, 0x4a, 0x3b, 0xab, 0xa9, 0x34, 0x94, 0x25, 0xc0, 0x29, 0x72, 0x84,
	0x0e, 0x28, 0x63, 0x10, 0x59, 0x77, 0x17, 0x76, 0x9a, 0x3c, 0x43, 0x4f, 0x38, 0xb0, 0x40, 0x48,
	0xb8, 0x13, 0x3c, 0x43, 0xd6, 0xd0, 0xc3, 0x1c, 0x34, 0x1a, 0x98, 0x25, 0xcb, 0x08, 0x1b, 0x90,
	0xfc, 0xca, 0x2a, 0x22, 0x3b, 0x68, 0xf3, 0x7a, 0xee, 0x51, 0xc0, 0x1c, 0xd9, 0x45, 0xd5, 0x1b,
	0x85, 0xba, 0x9c, 0xd4, 0xce, 0x31, 0x9e, 0xf7, 0x84, 0xde, 0x44, 0x50, 0xc6, 0x54, 0x2c, 0x2d,
	0x5e, 0x20, 0xdb, 0x68, 0x23, 0x87, 0x33, 0xc5, 0xc1, 0x51, 0xde, 0x01, 0x6d, 0x85, 0x81, 0x10,
	0xa4, 0xc5, 0x0f, 0x48, 0x0d, 0x6d, 0x47, 0xb1, 0x69, 0x39, 0xa9, 0xac, 0x38, 0x11, 0x2c, 0xdd,
	0xae, 0xa1, 0x29, 0x8c, 0xd5, 0x29, 0xad, 0xd8, 0x33, 0xf0, 0xdf, 0x18, 0xa7, 0xc1, 0x44, 0x4a,
	0x1a, 0xc0, 0x4b, 0x64, 0x13, 0xad, 0xdd, 0x04, 0xbf, 0x8e, 0x41, 0x77, 0x31, 0x21, 0xfb, 0x68,
	0xf7, 0x96, 0x60, 0x99, 0xe2, 0xa1, 0x2f, 0x69, 0xdc, 0x79, 0x09, 0x3f, 0x78, 0xd9, 0x97, 0x34,
	0x2e, 0x9c, 0x6d, 0x5f, 0xf1, 0xb2, 0x83, 0x50, 0x9d, 0x0a, 0xa7, 0x21, 0xe3, 0x71, 0x95, 0xac,
	0xa3, 0x95, 0xa6, 0x56, 0x71, 0xe4, 0x12, 0x89, 0x0a, 0xd9, 0x11, 0x36, 0xad, 0x6e, 0x8d, 0x2c,
	0xa1, 0x85, 0x74, 0x91, 0x83, 0xb4, 0xc2, 0x76, 0x71, 0xc5, 0xa3, 0x99, 0x0a, 0xc3, 0x58, 0x0a,
	0xdb, 0x75, 0x1c, 0x0c, 0xd3, 0x22, 0x4a, 0xd0, 0xeb, 0xa4, 0x8a, 0x96, 0xcb, 0xd0, 0x48, 0x9e,
	0x8d, 0x44, 0x7c, 0x5b, 0x68, 0xbd, 0x8c, 0x16, 0x1d, 0x55, 0xee, 0x54, 0x09, 0x89, 0x37, 0xc9,
	0x03, 0x34, 0x17, 0x09, 0x59, 0xc8, 0xbd, 0xea, 0x67, 0x06, 0xb8, 0x28, 0x67, 0x66, 0xcb, 0xdf,
	0xc6, 0x58, 0x6a, 0x63, 0x93, 0x8f, 0xcc, 0xb6, 0xaf, 0x87, 0x43, 0x00, 0x23, 0x73, 0xb2, 0xe3,
	0x85, 0x33, 0x4e, 0x17, 0xd9, 0xd1, 0x78, 0xd7, 0x0f, 0x12, 0x95, 0x4a, 0x76, 0x43, 0x15, 0x1b,
	0x17, 0x82, 0xd5, 0x82, 0xb9, 0x06, 0xb5, 0xac, 0x85, 0xf7, 0xc8, 0x0a, 0x5a, 0x4a, 0x36, 0x27,
	0x65, 0x6b, 0x08, 0x55, 0x07, 0x38, 0xae, 0xf9, 0xce, 0x95, 0xcb, 0xd9, 0x51, 0xc6, 0x93, 0xc8,
	0xf1, 0x23, 0x82, 0xd0, 0x54, 0x83, 0xb2, 0x76, 0x1c, 0xe1, 0xfd, 0x42, 0x72, 0x9e, 0xdd, 0x8e,
	0xaf, 0x94, 0x81, 0xb4, 0xa0, 0x53, 0xe8, 0x63, 0xb2, 0x87, 0xb6, 0xc6, 0x86, 0xd3, 0x89, 0x03,
	0x8e, 0x0f, 0xbc, 0xea, 0xc6, 0x42, 0xb8, 0x30, 0xa1, 0x30, 0x06, 0x38, 0x7e, 0x92, 0x30, 0xe1,
	0x31, 0x0d, 0xa5, 0xda, 0x21, 0xd5, 0x6d, 0x7c, 0x48, 0x56, 0x11, 0x49, 0x6f, 0x18, 0x00, 0xd5,
	0xae, 0x25, 0x8c, 0x55, 0xba, 0x8b, 0x9f, 0x7a, 0x1a, 0x93, 0x75, 0x03, 0xd6, 0x0a, 0xd9, 0xc4,
	0x47, 0xe4, 0x10, 0xed, 0x97, 0x8d, 0xc8, 0x6a, 0x71, 0x54, 0xb3, 0x96, 0xe8, 0x80, 0x0b, 0x69,
	0x53, 0x82, 0x0d, 0x84, 0x6c, 0xe3, 0x67, 0xbe, 0xd7, 0xa9, 0xe3, 0x68, 0x75, 0x22, 0x02, 0x70,
	0x91, 0x60, 0x36, 0xd6, 0x60, 0xf0, 0x57, 0x45, 0xda, 0x7c, 0x9a, 0xbe, 0x4e, 0x58, 0x4d, 0x7d,
	0x23, 0x1f, 0xaa, 0x5c, 0x96, 0x75, 0x4f, 0x9f, 0x06, 0xab, 0xd3, 0x49, 0xbb, 0x1a, 0x7c, 0x4e,
	0x0e, 0x50, 0xed, 0x56, 0x61, 0x94, 0xda, 0xfd, 0xa6, 0xec, 0x41, 0x01, 0xce, 0x6a, 0x32, 0xf8,
	0x5b, 0x6f, 0x06, 0xf9, 0xd6, 0xc2, 0x02, 0x40, 0x17, 0x33, 0x80, 0x8f, 0xbd, 0x2c, 0xae, 0xdd,
	0xef, 0x0a, 0xe0, 0x85, 0x4f, 0x91, 0x1b, 0xce, 0x58, 0xc4, 0x77, 0x85, 0x38, 0xac, 0x8e, 0x8d,
	0xf7, 0xd9, 0xd8, 0x80, 0xc6, 0xdf, 0x17, 0x3d, 0x1f, 0x45, 0x17, 0xf5, 0xfd, 0x50, 0xf4, 0xfc,
	0x5a, 0xe5, 0x8e, 0x03, 0x13, 0xc6, 0x27, 0xfe, 0x31, 0x75, 0xa2, 0x31, 0x14, 0x04, 0x40, 0x3b,
	0x80, 0x7f, 0xf2, 0xf1, 0x24, 0x45, 0xa6, 0x75, 0xef, 0xad, 0x61, 0x29, 0xf9, 0x9f, 0x8b, 0xe6,
	0x1b, 0xda, 0x01, 0x9e, 0x5b, 0x30, 0x7e, 0xe9, 0x3d, 0xa5, 0xcc, 0xcb, 0xa8, 0x64, 0x10, 0xdc,
	0x18, 0xbd, 0x5f, 0x3c, 0x33, 0x59, 0x6c, 0x6c, 0xdd, 0xaf, 0x8a, 0x66, 0xb7, 0xa1, 0xeb, 0x1f,
	0x1f, 0xfc, 0x6b, 0xc1, 0x84, 0x51, 0x4c, 0xd0, 0xc0, 0x79, 0xbd, 0x18, 0xfc, 0x1b, 0xa9, 0xa2,
	0x4a, 0xb2, 0x0c, 0xd2, 0x24, 0xe4, 0x48, 0x1a, 0x82, 0xe3, 0x60, 0xa9, 0x08, 0xf0, 0xef, 0xbe,
	0xd1, 0x63, 0x95, 0x3d, 0x6a, 0x56, 0x06, 0x53, 0xef, 0xa9, 0x77, 0xe2, 0x9c, 0x77, 0x02, 0xc0,
	0x0d, 0x7f, 0x64, 0x59, 0x22, 0x74, 0x40, 0x5a, 0x53, 0x10, 0xc3, 0xfc, 0x8b, 0x37, 0x12, 0xf5,
	0x76, 0x62, 0x5a, 0x54, 0x97, 0x1c, 0x81, 0xc1, 0xdc, 0xb3, 0x34, 0xaa, 0x65, 0xc7, 0x62, 0x63,
	0x55, 0x28, 0xde, 0xe4, 0xce, 0x11, 0x28, 0x8d, 0xa1, 0x90, 0x5f, 0x86, 0x32, 0x2e, 0x52, 0x46,
	0xa4, 0x57, 0x3e, 0x21, 0x8f, 0xd1, 0xde, 0x6d, 0xb7, 0x70, 0x1a, 0x4e, 0x81, 0xf9, 0xd9, 0x6e,
	0xfa, 0xb7, 0xb5, 0x84, 0x45, 0x5a, 0x74, 0x44, 0x00, 0xcd, 0x4c, 0x48, 0xe9, 0xab, 0x9d, 0x5f,
	0xbd, 0xe5, 0x9f, 0xcb, 0x11, 0xa5, 0xb7, 0xa8, 0xe6, 0x9e, 0x7f, 0x2c, 0xc8, 0x02, 0x9a, 0x2d,
	0xbc, 0x08, 0x9f, 0xd6, 0x26, 0x67, 0x16, 0xf1, 0xe2, 0xd1, 0xd6, 0x4d, 0xcb, 0x8b, 0xe2, 0x46,
	0x20, 0x52, 0x58, 0x63, 0xe1, 0xcd, 0x5c, 0xfd, 0xf9, 0xcb, 0xfc, 0xb7, 0xcc, 0xdb, 0xa9, 0xe4,
	0xbf, 0x17, 0xff, 0x06, 0x00, 0x00, 0xff, 0xff, 0xf8, 0x6d, 0x24, 0xd2, 0x72, 0x09, 0x00, 0x00,
}
