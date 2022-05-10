// Code generated by protoc-gen-go. DO NOT EDIT.
// source: membership_update_message.proto

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

type MembershipUpdateEvent_EventType int32

const (
	MembershipUpdateEvent_UNKNOWN        MembershipUpdateEvent_EventType = 0
	MembershipUpdateEvent_CHAT_CREATED   MembershipUpdateEvent_EventType = 1
	MembershipUpdateEvent_NAME_CHANGED   MembershipUpdateEvent_EventType = 2
	MembershipUpdateEvent_MEMBERS_ADDED  MembershipUpdateEvent_EventType = 3
	MembershipUpdateEvent_MEMBER_JOINED  MembershipUpdateEvent_EventType = 4
	MembershipUpdateEvent_MEMBER_REMOVED MembershipUpdateEvent_EventType = 5
	MembershipUpdateEvent_ADMINS_ADDED   MembershipUpdateEvent_EventType = 6
	MembershipUpdateEvent_ADMIN_REMOVED  MembershipUpdateEvent_EventType = 7
)

var MembershipUpdateEvent_EventType_name = map[int32]string{
	0: "UNKNOWN",
	1: "CHAT_CREATED",
	2: "NAME_CHANGED",
	3: "MEMBERS_ADDED",
	4: "MEMBER_JOINED",
	5: "MEMBER_REMOVED",
	6: "ADMINS_ADDED",
	7: "ADMIN_REMOVED",
}

var MembershipUpdateEvent_EventType_value = map[string]int32{
	"UNKNOWN":        0,
	"CHAT_CREATED":   1,
	"NAME_CHANGED":   2,
	"MEMBERS_ADDED":  3,
	"MEMBER_JOINED":  4,
	"MEMBER_REMOVED": 5,
	"ADMINS_ADDED":   6,
	"ADMIN_REMOVED":  7,
}

func (x MembershipUpdateEvent_EventType) String() string {
	return proto.EnumName(MembershipUpdateEvent_EventType_name, int32(x))
}

func (MembershipUpdateEvent_EventType) EnumDescriptor() ([]byte, []int) {
	return fileDescriptor_8d37dd0dc857a6be, []int{0, 0}
}

type MembershipUpdateEvent struct {
	// Lamport timestamp of the event
	Clock uint64 `protobuf:"varint,1,opt,name=clock,proto3" json:"clock,omitempty"`
	// List of public keys of objects of the action
	Members []string `protobuf:"bytes,2,rep,name=members,proto3" json:"members,omitempty"`
	// Name of the chat for the CHAT_CREATED/NAME_CHANGED event types
	Name string `protobuf:"bytes,3,opt,name=name,proto3" json:"name,omitempty"`
	// The type of the event
	Type                 MembershipUpdateEvent_EventType `protobuf:"varint,4,opt,name=type,proto3,enum=protobuf.MembershipUpdateEvent_EventType" json:"type,omitempty"`
	XXX_NoUnkeyedLiteral struct{}                        `json:"-"`
	XXX_unrecognized     []byte                          `json:"-"`
	XXX_sizecache        int32                           `json:"-"`
}

func (m *MembershipUpdateEvent) Reset()         { *m = MembershipUpdateEvent{} }
func (m *MembershipUpdateEvent) String() string { return proto.CompactTextString(m) }
func (*MembershipUpdateEvent) ProtoMessage()    {}
func (*MembershipUpdateEvent) Descriptor() ([]byte, []int) {
	return fileDescriptor_8d37dd0dc857a6be, []int{0}
}

func (m *MembershipUpdateEvent) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_MembershipUpdateEvent.Unmarshal(m, b)
}
func (m *MembershipUpdateEvent) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_MembershipUpdateEvent.Marshal(b, m, deterministic)
}
func (m *MembershipUpdateEvent) XXX_Merge(src proto.Message) {
	xxx_messageInfo_MembershipUpdateEvent.Merge(m, src)
}
func (m *MembershipUpdateEvent) XXX_Size() int {
	return xxx_messageInfo_MembershipUpdateEvent.Size(m)
}
func (m *MembershipUpdateEvent) XXX_DiscardUnknown() {
	xxx_messageInfo_MembershipUpdateEvent.DiscardUnknown(m)
}

var xxx_messageInfo_MembershipUpdateEvent proto.InternalMessageInfo

func (m *MembershipUpdateEvent) GetClock() uint64 {
	if m != nil {
		return m.Clock
	}
	return 0
}

func (m *MembershipUpdateEvent) GetMembers() []string {
	if m != nil {
		return m.Members
	}
	return nil
}

func (m *MembershipUpdateEvent) GetName() string {
	if m != nil {
		return m.Name
	}
	return ""
}

func (m *MembershipUpdateEvent) GetType() MembershipUpdateEvent_EventType {
	if m != nil {
		return m.Type
	}
	return MembershipUpdateEvent_UNKNOWN
}

// MembershipUpdateMessage is a message used to propagate information
// about group membership changes.
// For more information, see https://github.com/status-im/specs/blob/master/status-group-chats-spec.md.
type MembershipUpdateMessage struct {
	// The chat id of the private group chat
	ChatId string `protobuf:"bytes,1,opt,name=chat_id,json=chatId,proto3" json:"chat_id,omitempty"`
	// A list of events for this group chat, first x bytes are the signature, then is a
	// protobuf encoded MembershipUpdateEvent
	Events [][]byte `protobuf:"bytes,2,rep,name=events,proto3" json:"events,omitempty"`
	// An optional chat message
	//
	// Types that are valid to be assigned to ChatEntity:
	//	*MembershipUpdateMessage_Message
	//	*MembershipUpdateMessage_EmojiReaction
	ChatEntity           isMembershipUpdateMessage_ChatEntity `protobuf_oneof:"chat_entity"`
	XXX_NoUnkeyedLiteral struct{}                             `json:"-"`
	XXX_unrecognized     []byte                               `json:"-"`
	XXX_sizecache        int32                                `json:"-"`
}

func (m *MembershipUpdateMessage) Reset()         { *m = MembershipUpdateMessage{} }
func (m *MembershipUpdateMessage) String() string { return proto.CompactTextString(m) }
func (*MembershipUpdateMessage) ProtoMessage()    {}
func (*MembershipUpdateMessage) Descriptor() ([]byte, []int) {
	return fileDescriptor_8d37dd0dc857a6be, []int{1}
}

func (m *MembershipUpdateMessage) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_MembershipUpdateMessage.Unmarshal(m, b)
}
func (m *MembershipUpdateMessage) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_MembershipUpdateMessage.Marshal(b, m, deterministic)
}
func (m *MembershipUpdateMessage) XXX_Merge(src proto.Message) {
	xxx_messageInfo_MembershipUpdateMessage.Merge(m, src)
}
func (m *MembershipUpdateMessage) XXX_Size() int {
	return xxx_messageInfo_MembershipUpdateMessage.Size(m)
}
func (m *MembershipUpdateMessage) XXX_DiscardUnknown() {
	xxx_messageInfo_MembershipUpdateMessage.DiscardUnknown(m)
}

var xxx_messageInfo_MembershipUpdateMessage proto.InternalMessageInfo

func (m *MembershipUpdateMessage) GetChatId() string {
	if m != nil {
		return m.ChatId
	}
	return ""
}

func (m *MembershipUpdateMessage) GetEvents() [][]byte {
	if m != nil {
		return m.Events
	}
	return nil
}

type isMembershipUpdateMessage_ChatEntity interface {
	isMembershipUpdateMessage_ChatEntity()
}

type MembershipUpdateMessage_Message struct {
	Message *ChatMessage `protobuf:"bytes,3,opt,name=message,proto3,oneof"`
}

type MembershipUpdateMessage_EmojiReaction struct {
	EmojiReaction *EmojiReaction `protobuf:"bytes,4,opt,name=emoji_reaction,json=emojiReaction,proto3,oneof"`
}

func (*MembershipUpdateMessage_Message) isMembershipUpdateMessage_ChatEntity() {}

func (*MembershipUpdateMessage_EmojiReaction) isMembershipUpdateMessage_ChatEntity() {}

func (m *MembershipUpdateMessage) GetChatEntity() isMembershipUpdateMessage_ChatEntity {
	if m != nil {
		return m.ChatEntity
	}
	return nil
}

func (m *MembershipUpdateMessage) GetMessage() *ChatMessage {
	if x, ok := m.GetChatEntity().(*MembershipUpdateMessage_Message); ok {
		return x.Message
	}
	return nil
}

func (m *MembershipUpdateMessage) GetEmojiReaction() *EmojiReaction {
	if x, ok := m.GetChatEntity().(*MembershipUpdateMessage_EmojiReaction); ok {
		return x.EmojiReaction
	}
	return nil
}

// XXX_OneofWrappers is for the internal use of the proto package.
func (*MembershipUpdateMessage) XXX_OneofWrappers() []interface{} {
	return []interface{}{
		(*MembershipUpdateMessage_Message)(nil),
		(*MembershipUpdateMessage_EmojiReaction)(nil),
	}
}

func init() {
	proto.RegisterEnum("protobuf.MembershipUpdateEvent_EventType", MembershipUpdateEvent_EventType_name, MembershipUpdateEvent_EventType_value)
	proto.RegisterType((*MembershipUpdateEvent)(nil), "protobuf.MembershipUpdateEvent")
	proto.RegisterType((*MembershipUpdateMessage)(nil), "protobuf.MembershipUpdateMessage")
}

func init() { proto.RegisterFile("membership_update_message.proto", fileDescriptor_8d37dd0dc857a6be) }

var fileDescriptor_8d37dd0dc857a6be = []byte{
	// 405 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x6c, 0x51, 0x5d, 0x8f, 0x93, 0x40,
	0x14, 0x2d, 0x2d, 0x5b, 0xe4, 0xb2, 0x34, 0x78, 0xb3, 0x6b, 0xc9, 0xbe, 0x48, 0xfa, 0x84, 0x2f,
	0x18, 0xeb, 0xa3, 0x31, 0x91, 0x32, 0x13, 0xa9, 0x06, 0x9a, 0x8c, 0x5d, 0x4d, 0x7c, 0x21, 0x94,
	0x8e, 0x16, 0x95, 0x8f, 0xb4, 0xb3, 0x26, 0xfd, 0x2d, 0xfe, 0x11, 0x7f, 0x85, 0xbf, 0xc9, 0x30,
	0x40, 0xeb, 0x9a, 0x7d, 0x81, 0x39, 0xe7, 0xce, 0x39, 0x73, 0xef, 0xb9, 0xf0, 0xb4, 0xe0, 0xc5,
	0x86, 0xef, 0x0f, 0xbb, 0xbc, 0x4e, 0xee, 0xea, 0x6d, 0x2a, 0x78, 0x52, 0xf0, 0xc3, 0x21, 0xfd,
	0xca, 0xbd, 0x7a, 0x5f, 0x89, 0x0a, 0x1f, 0xc9, 0xdf, 0xe6, 0xee, 0xcb, 0x0d, 0x66, 0xbb, 0x54,
	0xdc, 0xaf, 0xde, 0x5c, 0xf1, 0xa2, 0xfa, 0x96, 0x27, 0x7b, 0x9e, 0x66, 0x22, 0xaf, 0xca, 0x96,
	0x9d, 0xfd, 0x1e, 0xc2, 0x75, 0x74, 0xf2, 0xbd, 0x95, 0xb6, 0xf4, 0x27, 0x2f, 0x05, 0x5e, 0xc1,
	0x45, 0xf6, 0xa3, 0xca, 0xbe, 0xdb, 0x8a, 0xa3, 0xb8, 0x2a, 0x6b, 0x01, 0xda, 0xa0, 0x75, 0x6d,
	0xd8, 0x43, 0x67, 0xe4, 0xea, 0xac, 0x87, 0x88, 0xa0, 0x96, 0x69, 0xc1, 0xed, 0x91, 0xa3, 0xb8,
	0x3a, 0x93, 0x67, 0x7c, 0x0d, 0xaa, 0x38, 0xd6, 0xdc, 0x56, 0x1d, 0xc5, 0x9d, 0xcc, 0x9f, 0x79,
	0x7d, 0x83, 0xde, 0x83, 0x4f, 0x7a, 0xf2, 0xbb, 0x3e, 0xd6, 0x9c, 0x49, 0xd9, 0xec, 0x97, 0x02,
	0xfa, 0x89, 0x43, 0x03, 0xb4, 0xdb, 0xf8, 0x7d, 0xbc, 0xfa, 0x14, 0x5b, 0x03, 0xb4, 0xe0, 0x32,
	0x08, 0xfd, 0x75, 0x12, 0x30, 0xea, 0xaf, 0x29, 0xb1, 0x94, 0x86, 0x89, 0xfd, 0x88, 0x26, 0x41,
	0xe8, 0xc7, 0x6f, 0x29, 0xb1, 0x86, 0xf8, 0x18, 0xcc, 0x88, 0x46, 0x0b, 0xca, 0x3e, 0x24, 0x3e,
	0x21, 0x94, 0x58, 0xa3, 0x33, 0x95, 0xbc, 0x5b, 0x2d, 0x63, 0x4a, 0x2c, 0x15, 0x11, 0x26, 0x1d,
	0xc5, 0x68, 0xb4, 0xfa, 0x48, 0x89, 0x75, 0xd1, 0x78, 0xf9, 0x24, 0x5a, 0xc6, 0xbd, 0x70, 0xdc,
	0x08, 0x25, 0x73, 0xba, 0xa4, 0xcd, 0xfe, 0x28, 0x30, 0xfd, 0x7f, 0x8e, 0xa8, 0x8d, 0x1c, 0xa7,
	0xa0, 0xc9, 0x15, 0xe4, 0x5b, 0x19, 0x9f, 0xce, 0xc6, 0x0d, 0x5c, 0x6e, 0xf1, 0x09, 0x8c, 0x79,
	0x33, 0x51, 0x1b, 0xdf, 0x25, 0xeb, 0x10, 0xbe, 0x68, 0x72, 0x95, 0x5a, 0x19, 0xa0, 0x31, 0xbf,
	0x3e, 0x87, 0x15, 0xec, 0x52, 0xd1, 0x19, 0x87, 0x03, 0xd6, 0xdf, 0xc3, 0x37, 0x30, 0xb9, 0xbf,
	0x52, 0x19, 0xb3, 0x31, 0x9f, 0x9e, 0x95, 0xb4, 0xa9, 0xb3, 0xae, 0x1c, 0x0e, 0x98, 0xc9, 0xff,
	0x25, 0x16, 0x26, 0x18, 0xb2, 0x4b, 0x5e, 0x8a, 0x5c, 0x1c, 0x17, 0xe6, 0x67, 0xc3, 0x7b, 0xfe,
	0xaa, 0x17, 0x6f, 0xc6, 0xf2, 0xf4, 0xf2, 0x6f, 0x00, 0x00, 0x00, 0xff, 0xff, 0xc2, 0x2a, 0x26,
	0x02, 0x78, 0x02, 0x00, 0x00,
}
