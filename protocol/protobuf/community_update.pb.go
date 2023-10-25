// Code generated by protoc-gen-go. DO NOT EDIT.
// source: community_update.proto

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

type CommunityEvent_EventType int32

const (
	CommunityEvent_UNKNOWN                                  CommunityEvent_EventType = 0
	CommunityEvent_COMMUNITY_EDIT                           CommunityEvent_EventType = 1
	CommunityEvent_COMMUNITY_MEMBER_TOKEN_PERMISSION_CHANGE CommunityEvent_EventType = 2
	CommunityEvent_COMMUNITY_MEMBER_TOKEN_PERMISSION_DELETE CommunityEvent_EventType = 3
	CommunityEvent_COMMUNITY_CATEGORY_CREATE                CommunityEvent_EventType = 4
	CommunityEvent_COMMUNITY_CATEGORY_DELETE                CommunityEvent_EventType = 5
	CommunityEvent_COMMUNITY_CATEGORY_EDIT                  CommunityEvent_EventType = 6
	CommunityEvent_COMMUNITY_CHANNEL_CREATE                 CommunityEvent_EventType = 7
	CommunityEvent_COMMUNITY_CHANNEL_DELETE                 CommunityEvent_EventType = 8
	CommunityEvent_COMMUNITY_CHANNEL_EDIT                   CommunityEvent_EventType = 9
	CommunityEvent_COMMUNITY_CATEGORY_REORDER               CommunityEvent_EventType = 10
	CommunityEvent_COMMUNITY_CHANNEL_REORDER                CommunityEvent_EventType = 11
	CommunityEvent_COMMUNITY_REQUEST_TO_JOIN_ACCEPT         CommunityEvent_EventType = 12
	CommunityEvent_COMMUNITY_REQUEST_TO_JOIN_REJECT         CommunityEvent_EventType = 13
	CommunityEvent_COMMUNITY_MEMBER_KICK                    CommunityEvent_EventType = 14
	CommunityEvent_COMMUNITY_MEMBER_BAN                     CommunityEvent_EventType = 15
	CommunityEvent_COMMUNITY_MEMBER_UNBAN                   CommunityEvent_EventType = 16
	CommunityEvent_COMMUNITY_TOKEN_ADD                      CommunityEvent_EventType = 17
)

var CommunityEvent_EventType_name = map[int32]string{
	0:  "UNKNOWN",
	1:  "COMMUNITY_EDIT",
	2:  "COMMUNITY_MEMBER_TOKEN_PERMISSION_CHANGE",
	3:  "COMMUNITY_MEMBER_TOKEN_PERMISSION_DELETE",
	4:  "COMMUNITY_CATEGORY_CREATE",
	5:  "COMMUNITY_CATEGORY_DELETE",
	6:  "COMMUNITY_CATEGORY_EDIT",
	7:  "COMMUNITY_CHANNEL_CREATE",
	8:  "COMMUNITY_CHANNEL_DELETE",
	9:  "COMMUNITY_CHANNEL_EDIT",
	10: "COMMUNITY_CATEGORY_REORDER",
	11: "COMMUNITY_CHANNEL_REORDER",
	12: "COMMUNITY_REQUEST_TO_JOIN_ACCEPT",
	13: "COMMUNITY_REQUEST_TO_JOIN_REJECT",
	14: "COMMUNITY_MEMBER_KICK",
	15: "COMMUNITY_MEMBER_BAN",
	16: "COMMUNITY_MEMBER_UNBAN",
	17: "COMMUNITY_TOKEN_ADD",
}

var CommunityEvent_EventType_value = map[string]int32{
	"UNKNOWN":        0,
	"COMMUNITY_EDIT": 1,
	"COMMUNITY_MEMBER_TOKEN_PERMISSION_CHANGE": 2,
	"COMMUNITY_MEMBER_TOKEN_PERMISSION_DELETE": 3,
	"COMMUNITY_CATEGORY_CREATE":                4,
	"COMMUNITY_CATEGORY_DELETE":                5,
	"COMMUNITY_CATEGORY_EDIT":                  6,
	"COMMUNITY_CHANNEL_CREATE":                 7,
	"COMMUNITY_CHANNEL_DELETE":                 8,
	"COMMUNITY_CHANNEL_EDIT":                   9,
	"COMMUNITY_CATEGORY_REORDER":               10,
	"COMMUNITY_CHANNEL_REORDER":                11,
	"COMMUNITY_REQUEST_TO_JOIN_ACCEPT":         12,
	"COMMUNITY_REQUEST_TO_JOIN_REJECT":         13,
	"COMMUNITY_MEMBER_KICK":                    14,
	"COMMUNITY_MEMBER_BAN":                     15,
	"COMMUNITY_MEMBER_UNBAN":                   16,
	"COMMUNITY_TOKEN_ADD":                      17,
}

func (x CommunityEvent_EventType) String() string {
	return proto.EnumName(CommunityEvent_EventType_name, int32(x))
}

func (CommunityEvent_EventType) EnumDescriptor() ([]byte, []int) {
	return fileDescriptor_52ed23dfc73918ab, []int{0, 0}
}

type CommunityEvent struct {
	CommunityEventClock    uint64                             `protobuf:"varint,1,opt,name=community_event_clock,json=communityEventClock,proto3" json:"community_event_clock,omitempty"`
	Type                   CommunityEvent_EventType           `protobuf:"varint,2,opt,name=type,proto3,enum=protobuf.CommunityEvent_EventType" json:"type,omitempty"`
	CommunityConfig        *CommunityConfig                   `protobuf:"bytes,3,opt,name=community_config,json=communityConfig,proto3" json:"community_config,omitempty"`
	TokenPermission        *CommunityTokenPermission          `protobuf:"bytes,4,opt,name=token_permission,json=tokenPermission,proto3" json:"token_permission,omitempty"`
	CategoryData           *CategoryData                      `protobuf:"bytes,5,opt,name=category_data,json=categoryData,proto3" json:"category_data,omitempty"`
	ChannelData            *ChannelData                       `protobuf:"bytes,6,opt,name=channel_data,json=channelData,proto3" json:"channel_data,omitempty"`
	MemberToAction         string                             `protobuf:"bytes,7,opt,name=member_to_action,json=memberToAction,proto3" json:"member_to_action,omitempty"`
	MembersAdded           map[string]*CommunityMember        `protobuf:"bytes,8,rep,name=membersAdded,proto3" json:"membersAdded,omitempty" protobuf_key:"bytes,1,opt,name=key,proto3" protobuf_val:"bytes,2,opt,name=value,proto3"`
	RejectedRequestsToJoin map[string]*CommunityRequestToJoin `protobuf:"bytes,9,rep,name=rejectedRequestsToJoin,proto3" json:"rejectedRequestsToJoin,omitempty" protobuf_key:"bytes,1,opt,name=key,proto3" protobuf_val:"bytes,2,opt,name=value,proto3"`
	AcceptedRequestsToJoin map[string]*CommunityRequestToJoin `protobuf:"bytes,10,rep,name=acceptedRequestsToJoin,proto3" json:"acceptedRequestsToJoin,omitempty" protobuf_key:"bytes,1,opt,name=key,proto3" protobuf_val:"bytes,2,opt,name=value,proto3"`
	TokenMetadata          *CommunityTokenMetadata            `protobuf:"bytes,11,opt,name=token_metadata,json=tokenMetadata,proto3" json:"token_metadata,omitempty"`
	XXX_NoUnkeyedLiteral   struct{}                           `json:"-"`
	XXX_unrecognized       []byte                             `json:"-"`
	XXX_sizecache          int32                              `json:"-"`
}

func (m *CommunityEvent) Reset()         { *m = CommunityEvent{} }
func (m *CommunityEvent) String() string { return proto.CompactTextString(m) }
func (*CommunityEvent) ProtoMessage()    {}
func (*CommunityEvent) Descriptor() ([]byte, []int) {
	return fileDescriptor_52ed23dfc73918ab, []int{0}
}

func (m *CommunityEvent) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_CommunityEvent.Unmarshal(m, b)
}
func (m *CommunityEvent) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_CommunityEvent.Marshal(b, m, deterministic)
}
func (m *CommunityEvent) XXX_Merge(src proto.Message) {
	xxx_messageInfo_CommunityEvent.Merge(m, src)
}
func (m *CommunityEvent) XXX_Size() int {
	return xxx_messageInfo_CommunityEvent.Size(m)
}
func (m *CommunityEvent) XXX_DiscardUnknown() {
	xxx_messageInfo_CommunityEvent.DiscardUnknown(m)
}

var xxx_messageInfo_CommunityEvent proto.InternalMessageInfo

func (m *CommunityEvent) GetCommunityEventClock() uint64 {
	if m != nil {
		return m.CommunityEventClock
	}
	return 0
}

func (m *CommunityEvent) GetType() CommunityEvent_EventType {
	if m != nil {
		return m.Type
	}
	return CommunityEvent_UNKNOWN
}

func (m *CommunityEvent) GetCommunityConfig() *CommunityConfig {
	if m != nil {
		return m.CommunityConfig
	}
	return nil
}

func (m *CommunityEvent) GetTokenPermission() *CommunityTokenPermission {
	if m != nil {
		return m.TokenPermission
	}
	return nil
}

func (m *CommunityEvent) GetCategoryData() *CategoryData {
	if m != nil {
		return m.CategoryData
	}
	return nil
}

func (m *CommunityEvent) GetChannelData() *ChannelData {
	if m != nil {
		return m.ChannelData
	}
	return nil
}

func (m *CommunityEvent) GetMemberToAction() string {
	if m != nil {
		return m.MemberToAction
	}
	return ""
}

func (m *CommunityEvent) GetMembersAdded() map[string]*CommunityMember {
	if m != nil {
		return m.MembersAdded
	}
	return nil
}

func (m *CommunityEvent) GetRejectedRequestsToJoin() map[string]*CommunityRequestToJoin {
	if m != nil {
		return m.RejectedRequestsToJoin
	}
	return nil
}

func (m *CommunityEvent) GetAcceptedRequestsToJoin() map[string]*CommunityRequestToJoin {
	if m != nil {
		return m.AcceptedRequestsToJoin
	}
	return nil
}

func (m *CommunityEvent) GetTokenMetadata() *CommunityTokenMetadata {
	if m != nil {
		return m.TokenMetadata
	}
	return nil
}

type CommunityConfig struct {
	Identity             *ChatIdentity           `protobuf:"bytes,1,opt,name=identity,proto3" json:"identity,omitempty"`
	Permissions          *CommunityPermissions   `protobuf:"bytes,2,opt,name=permissions,proto3" json:"permissions,omitempty"`
	AdminSettings        *CommunityAdminSettings `protobuf:"bytes,3,opt,name=admin_settings,json=adminSettings,proto3" json:"admin_settings,omitempty"`
	IntroMessage         string                  `protobuf:"bytes,4,opt,name=intro_message,json=introMessage,proto3" json:"intro_message,omitempty"`
	OutroMessage         string                  `protobuf:"bytes,5,opt,name=outro_message,json=outroMessage,proto3" json:"outro_message,omitempty"`
	Tags                 []string                `protobuf:"bytes,6,rep,name=tags,proto3" json:"tags,omitempty"`
	XXX_NoUnkeyedLiteral struct{}                `json:"-"`
	XXX_unrecognized     []byte                  `json:"-"`
	XXX_sizecache        int32                   `json:"-"`
}

func (m *CommunityConfig) Reset()         { *m = CommunityConfig{} }
func (m *CommunityConfig) String() string { return proto.CompactTextString(m) }
func (*CommunityConfig) ProtoMessage()    {}
func (*CommunityConfig) Descriptor() ([]byte, []int) {
	return fileDescriptor_52ed23dfc73918ab, []int{1}
}

func (m *CommunityConfig) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_CommunityConfig.Unmarshal(m, b)
}
func (m *CommunityConfig) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_CommunityConfig.Marshal(b, m, deterministic)
}
func (m *CommunityConfig) XXX_Merge(src proto.Message) {
	xxx_messageInfo_CommunityConfig.Merge(m, src)
}
func (m *CommunityConfig) XXX_Size() int {
	return xxx_messageInfo_CommunityConfig.Size(m)
}
func (m *CommunityConfig) XXX_DiscardUnknown() {
	xxx_messageInfo_CommunityConfig.DiscardUnknown(m)
}

var xxx_messageInfo_CommunityConfig proto.InternalMessageInfo

func (m *CommunityConfig) GetIdentity() *ChatIdentity {
	if m != nil {
		return m.Identity
	}
	return nil
}

func (m *CommunityConfig) GetPermissions() *CommunityPermissions {
	if m != nil {
		return m.Permissions
	}
	return nil
}

func (m *CommunityConfig) GetAdminSettings() *CommunityAdminSettings {
	if m != nil {
		return m.AdminSettings
	}
	return nil
}

func (m *CommunityConfig) GetIntroMessage() string {
	if m != nil {
		return m.IntroMessage
	}
	return ""
}

func (m *CommunityConfig) GetOutroMessage() string {
	if m != nil {
		return m.OutroMessage
	}
	return ""
}

func (m *CommunityConfig) GetTags() []string {
	if m != nil {
		return m.Tags
	}
	return nil
}

type CategoryData struct {
	CategoryId           string   `protobuf:"bytes,1,opt,name=category_id,json=categoryId,proto3" json:"category_id,omitempty"`
	Name                 string   `protobuf:"bytes,2,opt,name=name,proto3" json:"name,omitempty"`
	ChannelsIds          []string `protobuf:"bytes,3,rep,name=channels_ids,json=channelsIds,proto3" json:"channels_ids,omitempty"`
	Position             int32    `protobuf:"varint,4,opt,name=position,proto3" json:"position,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *CategoryData) Reset()         { *m = CategoryData{} }
func (m *CategoryData) String() string { return proto.CompactTextString(m) }
func (*CategoryData) ProtoMessage()    {}
func (*CategoryData) Descriptor() ([]byte, []int) {
	return fileDescriptor_52ed23dfc73918ab, []int{2}
}

func (m *CategoryData) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_CategoryData.Unmarshal(m, b)
}
func (m *CategoryData) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_CategoryData.Marshal(b, m, deterministic)
}
func (m *CategoryData) XXX_Merge(src proto.Message) {
	xxx_messageInfo_CategoryData.Merge(m, src)
}
func (m *CategoryData) XXX_Size() int {
	return xxx_messageInfo_CategoryData.Size(m)
}
func (m *CategoryData) XXX_DiscardUnknown() {
	xxx_messageInfo_CategoryData.DiscardUnknown(m)
}

var xxx_messageInfo_CategoryData proto.InternalMessageInfo

func (m *CategoryData) GetCategoryId() string {
	if m != nil {
		return m.CategoryId
	}
	return ""
}

func (m *CategoryData) GetName() string {
	if m != nil {
		return m.Name
	}
	return ""
}

func (m *CategoryData) GetChannelsIds() []string {
	if m != nil {
		return m.ChannelsIds
	}
	return nil
}

func (m *CategoryData) GetPosition() int32 {
	if m != nil {
		return m.Position
	}
	return 0
}

type ChannelData struct {
	CategoryId           string         `protobuf:"bytes,1,opt,name=category_id,json=categoryId,proto3" json:"category_id,omitempty"`
	ChannelId            string         `protobuf:"bytes,2,opt,name=channel_id,json=channelId,proto3" json:"channel_id,omitempty"`
	Position             int32          `protobuf:"varint,3,opt,name=position,proto3" json:"position,omitempty"`
	Channel              *CommunityChat `protobuf:"bytes,4,opt,name=channel,proto3" json:"channel,omitempty"`
	XXX_NoUnkeyedLiteral struct{}       `json:"-"`
	XXX_unrecognized     []byte         `json:"-"`
	XXX_sizecache        int32          `json:"-"`
}

func (m *ChannelData) Reset()         { *m = ChannelData{} }
func (m *ChannelData) String() string { return proto.CompactTextString(m) }
func (*ChannelData) ProtoMessage()    {}
func (*ChannelData) Descriptor() ([]byte, []int) {
	return fileDescriptor_52ed23dfc73918ab, []int{3}
}

func (m *ChannelData) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_ChannelData.Unmarshal(m, b)
}
func (m *ChannelData) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_ChannelData.Marshal(b, m, deterministic)
}
func (m *ChannelData) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ChannelData.Merge(m, src)
}
func (m *ChannelData) XXX_Size() int {
	return xxx_messageInfo_ChannelData.Size(m)
}
func (m *ChannelData) XXX_DiscardUnknown() {
	xxx_messageInfo_ChannelData.DiscardUnknown(m)
}

var xxx_messageInfo_ChannelData proto.InternalMessageInfo

func (m *ChannelData) GetCategoryId() string {
	if m != nil {
		return m.CategoryId
	}
	return ""
}

func (m *ChannelData) GetChannelId() string {
	if m != nil {
		return m.ChannelId
	}
	return ""
}

func (m *ChannelData) GetPosition() int32 {
	if m != nil {
		return m.Position
	}
	return 0
}

func (m *ChannelData) GetChannel() *CommunityChat {
	if m != nil {
		return m.Channel
	}
	return nil
}

type SignedCommunityEvent struct {
	// Signature of the payload field
	Signature []byte `protobuf:"bytes,1,opt,name=signature,proto3" json:"signature,omitempty"`
	// Marshaled CommunityEvent
	Payload              []byte   `protobuf:"bytes,2,opt,name=payload,proto3" json:"payload,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *SignedCommunityEvent) Reset()         { *m = SignedCommunityEvent{} }
func (m *SignedCommunityEvent) String() string { return proto.CompactTextString(m) }
func (*SignedCommunityEvent) ProtoMessage()    {}
func (*SignedCommunityEvent) Descriptor() ([]byte, []int) {
	return fileDescriptor_52ed23dfc73918ab, []int{4}
}

func (m *SignedCommunityEvent) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_SignedCommunityEvent.Unmarshal(m, b)
}
func (m *SignedCommunityEvent) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_SignedCommunityEvent.Marshal(b, m, deterministic)
}
func (m *SignedCommunityEvent) XXX_Merge(src proto.Message) {
	xxx_messageInfo_SignedCommunityEvent.Merge(m, src)
}
func (m *SignedCommunityEvent) XXX_Size() int {
	return xxx_messageInfo_SignedCommunityEvent.Size(m)
}
func (m *SignedCommunityEvent) XXX_DiscardUnknown() {
	xxx_messageInfo_SignedCommunityEvent.DiscardUnknown(m)
}

var xxx_messageInfo_SignedCommunityEvent proto.InternalMessageInfo

func (m *SignedCommunityEvent) GetSignature() []byte {
	if m != nil {
		return m.Signature
	}
	return nil
}

func (m *SignedCommunityEvent) GetPayload() []byte {
	if m != nil {
		return m.Payload
	}
	return nil
}

// CommunityEventsMessage is a message used to propagate information
// about community changes.
type CommunityEventsMessage struct {
	CommunityId []byte `protobuf:"bytes,1,opt,name=community_id,json=communityId,proto3" json:"community_id,omitempty"`
	// Events base CommunityDescription with owner signature on top of which events were generated
	EventsBaseCommunityDescription []byte `protobuf:"bytes,2,opt,name=events_base_community_description,json=eventsBaseCommunityDescription,proto3" json:"events_base_community_description,omitempty"`
	// A list of admins events for the channel in bytes
	// Deprecated: use signed_events instead.
	Events [][]byte `protobuf:"bytes,3,rep,name=events,proto3" json:"events,omitempty"` // Deprecated: Do not use.
	// A list of signed community events
	SignedEvents         []*SignedCommunityEvent `protobuf:"bytes,4,rep,name=signed_events,json=signedEvents,proto3" json:"signed_events,omitempty"`
	XXX_NoUnkeyedLiteral struct{}                `json:"-"`
	XXX_unrecognized     []byte                  `json:"-"`
	XXX_sizecache        int32                   `json:"-"`
}

func (m *CommunityEventsMessage) Reset()         { *m = CommunityEventsMessage{} }
func (m *CommunityEventsMessage) String() string { return proto.CompactTextString(m) }
func (*CommunityEventsMessage) ProtoMessage()    {}
func (*CommunityEventsMessage) Descriptor() ([]byte, []int) {
	return fileDescriptor_52ed23dfc73918ab, []int{5}
}

func (m *CommunityEventsMessage) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_CommunityEventsMessage.Unmarshal(m, b)
}
func (m *CommunityEventsMessage) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_CommunityEventsMessage.Marshal(b, m, deterministic)
}
func (m *CommunityEventsMessage) XXX_Merge(src proto.Message) {
	xxx_messageInfo_CommunityEventsMessage.Merge(m, src)
}
func (m *CommunityEventsMessage) XXX_Size() int {
	return xxx_messageInfo_CommunityEventsMessage.Size(m)
}
func (m *CommunityEventsMessage) XXX_DiscardUnknown() {
	xxx_messageInfo_CommunityEventsMessage.DiscardUnknown(m)
}

var xxx_messageInfo_CommunityEventsMessage proto.InternalMessageInfo

func (m *CommunityEventsMessage) GetCommunityId() []byte {
	if m != nil {
		return m.CommunityId
	}
	return nil
}

func (m *CommunityEventsMessage) GetEventsBaseCommunityDescription() []byte {
	if m != nil {
		return m.EventsBaseCommunityDescription
	}
	return nil
}

// Deprecated: Do not use.
func (m *CommunityEventsMessage) GetEvents() [][]byte {
	if m != nil {
		return m.Events
	}
	return nil
}

func (m *CommunityEventsMessage) GetSignedEvents() []*SignedCommunityEvent {
	if m != nil {
		return m.SignedEvents
	}
	return nil
}

type CommunityEventsMessageRejected struct {
	Msg                  *CommunityEventsMessage `protobuf:"bytes,1,opt,name=msg,proto3" json:"msg,omitempty"`
	XXX_NoUnkeyedLiteral struct{}                `json:"-"`
	XXX_unrecognized     []byte                  `json:"-"`
	XXX_sizecache        int32                   `json:"-"`
}

func (m *CommunityEventsMessageRejected) Reset()         { *m = CommunityEventsMessageRejected{} }
func (m *CommunityEventsMessageRejected) String() string { return proto.CompactTextString(m) }
func (*CommunityEventsMessageRejected) ProtoMessage()    {}
func (*CommunityEventsMessageRejected) Descriptor() ([]byte, []int) {
	return fileDescriptor_52ed23dfc73918ab, []int{6}
}

func (m *CommunityEventsMessageRejected) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_CommunityEventsMessageRejected.Unmarshal(m, b)
}
func (m *CommunityEventsMessageRejected) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_CommunityEventsMessageRejected.Marshal(b, m, deterministic)
}
func (m *CommunityEventsMessageRejected) XXX_Merge(src proto.Message) {
	xxx_messageInfo_CommunityEventsMessageRejected.Merge(m, src)
}
func (m *CommunityEventsMessageRejected) XXX_Size() int {
	return xxx_messageInfo_CommunityEventsMessageRejected.Size(m)
}
func (m *CommunityEventsMessageRejected) XXX_DiscardUnknown() {
	xxx_messageInfo_CommunityEventsMessageRejected.DiscardUnknown(m)
}

var xxx_messageInfo_CommunityEventsMessageRejected proto.InternalMessageInfo

func (m *CommunityEventsMessageRejected) GetMsg() *CommunityEventsMessage {
	if m != nil {
		return m.Msg
	}
	return nil
}

func init() {
	proto.RegisterEnum("protobuf.CommunityEvent_EventType", CommunityEvent_EventType_name, CommunityEvent_EventType_value)
	proto.RegisterType((*CommunityEvent)(nil), "protobuf.CommunityEvent")
	proto.RegisterMapType((map[string]*CommunityRequestToJoin)(nil), "protobuf.CommunityEvent.AcceptedRequestsToJoinEntry")
	proto.RegisterMapType((map[string]*CommunityMember)(nil), "protobuf.CommunityEvent.MembersAddedEntry")
	proto.RegisterMapType((map[string]*CommunityRequestToJoin)(nil), "protobuf.CommunityEvent.RejectedRequestsToJoinEntry")
	proto.RegisterType((*CommunityConfig)(nil), "protobuf.CommunityConfig")
	proto.RegisterType((*CategoryData)(nil), "protobuf.CategoryData")
	proto.RegisterType((*ChannelData)(nil), "protobuf.ChannelData")
	proto.RegisterType((*SignedCommunityEvent)(nil), "protobuf.SignedCommunityEvent")
	proto.RegisterType((*CommunityEventsMessage)(nil), "protobuf.CommunityEventsMessage")
	proto.RegisterType((*CommunityEventsMessageRejected)(nil), "protobuf.CommunityEventsMessageRejected")
}

func init() {
	proto.RegisterFile("community_update.proto", fileDescriptor_52ed23dfc73918ab)
}

var fileDescriptor_52ed23dfc73918ab = []byte{
	// 1095 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xac, 0x56, 0xcb, 0x6e, 0xdb, 0x46,
	0x14, 0xad, 0x2c, 0xf9, 0xa1, 0xab, 0x87, 0xe9, 0x71, 0x6c, 0xd3, 0x4a, 0xe2, 0x2a, 0x6a, 0x17,
	0x42, 0x51, 0x38, 0xa8, 0x5a, 0x04, 0x41, 0xb3, 0xa9, 0x4c, 0x0d, 0x1c, 0xda, 0x11, 0xe5, 0x8e,
	0x69, 0x14, 0xc9, 0x86, 0x18, 0x93, 0x13, 0x99, 0xb5, 0x45, 0xaa, 0x9a, 0x51, 0x00, 0x6d, 0xfb,
	0x05, 0xfd, 0x80, 0x7e, 0x43, 0xd1, 0xbf, 0xea, 0x6f, 0x14, 0x9c, 0x21, 0x45, 0xd2, 0xa6, 0x9c,
	0x2e, 0xba, 0x91, 0x38, 0xf7, 0x9c, 0x7b, 0xce, 0xdc, 0x79, 0xf0, 0x12, 0xf6, 0xdd, 0x70, 0x32,
	0x99, 0x07, 0xbe, 0x58, 0x38, 0xf3, 0xa9, 0x47, 0x05, 0x3b, 0x9e, 0xce, 0x42, 0x11, 0xa2, 0x2d,
	0xf9, 0x77, 0x3d, 0xff, 0xd8, 0xda, 0x75, 0x6f, 0xa8, 0x70, 0x7c, 0x8f, 0x05, 0xc2, 0x17, 0x0b,
	0x05, 0xb7, 0x76, 0x92, 0x34, 0x9f, 0x71, 0x15, 0xea, 0xfc, 0xd1, 0x80, 0xa6, 0x91, 0x88, 0xe1,
	0x4f, 0x2c, 0x10, 0xa8, 0x07, 0x7b, 0xa9, 0x3c, 0x8b, 0x42, 0x8e, 0x7b, 0x17, 0xba, 0xb7, 0x7a,
	0xa9, 0x5d, 0xea, 0x56, 0xc8, 0xae, 0x9b, 0xa3, 0x1b, 0x11, 0x84, 0x5e, 0x41, 0x45, 0x2c, 0xa6,
	0x4c, 0x5f, 0x6b, 0x97, 0xba, 0xcd, 0x5e, 0xe7, 0x38, 0x99, 0xc7, 0x71, 0x5e, 0xfb, 0x58, 0xfe,
	0xda, 0x8b, 0x29, 0x23, 0x92, 0x8f, 0x06, 0xa0, 0xa5, 0x5e, 0x6e, 0x18, 0x7c, 0xf4, 0xc7, 0x7a,
	0xb9, 0x5d, 0xea, 0xd6, 0x7a, 0x87, 0x05, 0x1a, 0x86, 0x24, 0x90, 0x6d, 0x37, 0x1f, 0x40, 0x43,
	0xd0, 0x44, 0x78, 0xcb, 0x02, 0x67, 0xca, 0x66, 0x13, 0x9f, 0x73, 0x3f, 0x0c, 0xf4, 0x8a, 0x54,
	0x29, 0x9a, 0x89, 0x1d, 0x51, 0x2f, 0x96, 0x4c, 0xb2, 0x2d, 0xf2, 0x01, 0xf4, 0x06, 0x1a, 0x2e,
	0x15, 0x6c, 0x1c, 0xce, 0x16, 0x8e, 0x47, 0x05, 0xd5, 0xd7, 0xa5, 0xd6, 0x7e, 0x46, 0x2b, 0x86,
	0x07, 0x54, 0x50, 0x52, 0x77, 0x33, 0x23, 0xf4, 0x1a, 0xea, 0xee, 0x0d, 0x0d, 0x02, 0x76, 0xa7,
	0x72, 0x37, 0x64, 0xee, 0x5e, 0x26, 0x57, 0xa1, 0x32, 0xb5, 0xe6, 0xa6, 0x03, 0xd4, 0x05, 0x6d,
	0xc2, 0x26, 0xd7, 0x6c, 0xe6, 0x88, 0xd0, 0xa1, 0xae, 0x88, 0xaa, 0xd8, 0x6c, 0x97, 0xba, 0x55,
	0xd2, 0x54, 0x71, 0x3b, 0xec, 0xcb, 0x28, 0xb2, 0xa0, 0xae, 0x22, 0xbc, 0xef, 0x79, 0xcc, 0xd3,
	0xb7, 0xda, 0xe5, 0x6e, 0xad, 0xf7, 0xcd, 0xca, 0x55, 0x1f, 0x66, 0xc8, 0x38, 0x10, 0xb3, 0x05,
	0xc9, 0xe5, 0xa3, 0x3b, 0xd8, 0x9f, 0xb1, 0x5f, 0x99, 0x2b, 0x98, 0x47, 0xd8, 0x6f, 0x73, 0xc6,
	0x05, 0xb7, 0xc3, 0xb3, 0xd0, 0x0f, 0xf4, 0xaa, 0x54, 0xfe, 0x61, 0xa5, 0x32, 0x29, 0x4c, 0x53,
	0x1e, 0x2b, 0x34, 0x23, 0x37, 0xea, 0xba, 0x6c, 0xfa, 0xd0, 0x0d, 0x3e, 0xe3, 0xd6, 0x2f, 0x4c,
	0x8b, 0xdd, 0x8a, 0x35, 0xd1, 0x29, 0x34, 0xd5, 0xd9, 0x98, 0x30, 0x41, 0xe5, 0x8e, 0xd4, 0xe4,
	0x8e, 0xb4, 0x57, 0x9d, 0x8c, 0x61, 0xcc, 0x23, 0x0d, 0x91, 0x1d, 0xb6, 0x3e, 0xc0, 0xce, 0x83,
	0x75, 0x44, 0x1a, 0x94, 0x6f, 0xd9, 0x42, 0xde, 0x8c, 0x2a, 0x89, 0x1e, 0xd1, 0x4b, 0x58, 0xff,
	0x44, 0xef, 0xe6, 0xea, 0x2a, 0x14, 0x1f, 0x63, 0x25, 0x43, 0x14, 0xef, 0xc7, 0xb5, 0xd7, 0xa5,
	0xd6, 0x2d, 0x3c, 0x7d, 0x64, 0x25, 0x0b, 0x5c, 0x5e, 0xe5, 0x5d, 0x8a, 0x8a, 0x89, 0x85, 0x94,
	0xce, 0x3d, 0xb3, 0x47, 0x16, 0xf2, 0xff, 0x35, 0xeb, 0xfc, 0x5d, 0x81, 0xea, 0xf2, 0xd2, 0xa3,
	0x1a, 0x6c, 0x5e, 0x59, 0xe7, 0xd6, 0xe8, 0x17, 0x4b, 0xfb, 0x02, 0x21, 0x68, 0x1a, 0xa3, 0xe1,
	0xf0, 0xca, 0x32, 0xed, 0xf7, 0x0e, 0x1e, 0x98, 0xb6, 0x56, 0x42, 0xdf, 0x42, 0x37, 0x8d, 0x0d,
	0xf1, 0xf0, 0x04, 0x13, 0xc7, 0x1e, 0x9d, 0x63, 0xcb, 0xb9, 0xc0, 0x64, 0x68, 0x5e, 0x5e, 0x9a,
	0x23, 0xcb, 0x31, 0xde, 0xf6, 0xad, 0x53, 0xac, 0xad, 0xfd, 0x37, 0xf6, 0x00, 0xbf, 0xc3, 0x36,
	0xd6, 0xca, 0xe8, 0x39, 0x1c, 0xa6, 0x6c, 0xa3, 0x6f, 0xe3, 0xd3, 0x11, 0x79, 0xef, 0x18, 0x04,
	0xf7, 0x6d, 0xac, 0x55, 0x56, 0xc0, 0x71, 0xf6, 0x3a, 0x7a, 0x0a, 0x07, 0x05, 0xb0, 0x9c, 0xf6,
	0x06, 0x7a, 0x06, 0x7a, 0x06, 0x7c, 0xdb, 0xb7, 0x2c, 0xfc, 0x2e, 0x51, 0xde, 0x2c, 0x46, 0x63,
	0xe1, 0x2d, 0xd4, 0x82, 0xfd, 0x87, 0xa8, 0xd4, 0xad, 0xa2, 0x23, 0x68, 0x15, 0x98, 0x12, 0x3c,
	0x22, 0x03, 0x4c, 0x34, 0xb8, 0x37, 0xe7, 0x38, 0x37, 0x81, 0x6b, 0xe8, 0x6b, 0x68, 0xa7, 0x30,
	0xc1, 0x3f, 0x5f, 0xe1, 0x4b, 0xdb, 0xb1, 0x47, 0xce, 0xd9, 0xc8, 0xb4, 0x9c, 0xbe, 0x61, 0xe0,
	0x0b, 0x5b, 0xab, 0x3f, 0xce, 0x22, 0xf8, 0x0c, 0x1b, 0xb6, 0xd6, 0x40, 0x87, 0xb0, 0xf7, 0x60,
	0xad, 0xcf, 0x4d, 0xe3, 0x5c, 0x6b, 0x22, 0x1d, 0x9e, 0x3c, 0x80, 0x4e, 0xfa, 0x96, 0xb6, 0x9d,
	0xaf, 0x2d, 0x46, 0xae, 0xac, 0x08, 0xd3, 0xd0, 0x01, 0xec, 0xa6, 0x98, 0xda, 0xb5, 0xfe, 0x60,
	0xa0, 0xed, 0x74, 0xfe, 0x5a, 0x83, 0xed, 0x7b, 0xaf, 0x7c, 0xd4, 0x83, 0xad, 0xa4, 0x97, 0xc9,
	0x93, 0x99, 0x7f, 0x1b, 0xdf, 0x50, 0x61, 0xc6, 0x28, 0x59, 0xf2, 0xd0, 0x4f, 0x50, 0x4b, 0xfb,
	0x01, 0x8f, 0x0f, 0xef, 0x51, 0xc1, 0xe1, 0x4d, 0x5f, 0xfd, 0x9c, 0x64, 0x53, 0xa2, 0x77, 0x07,
	0xf5, 0x26, 0x7e, 0xe0, 0x70, 0x26, 0x84, 0x1f, 0x8c, 0x79, 0xdc, 0x9b, 0x8a, 0x6e, 0x40, 0x3f,
	0x22, 0x5e, 0xc6, 0x3c, 0xd2, 0xa0, 0xd9, 0x21, 0xfa, 0x0a, 0x1a, 0x7e, 0x20, 0x66, 0xa1, 0x33,
	0x61, 0x9c, 0xd3, 0x31, 0x93, 0xdd, 0xa9, 0x4a, 0xea, 0x32, 0x38, 0x54, 0xb1, 0x88, 0x14, 0xce,
	0xb3, 0xa4, 0x75, 0x45, 0x92, 0xc1, 0x84, 0x84, 0xa0, 0x22, 0xe8, 0x98, 0xeb, 0x1b, 0xed, 0x72,
	0xb7, 0x4a, 0xe4, 0x73, 0xe7, 0xf7, 0x12, 0xd4, 0xb3, 0x1d, 0x09, 0x7d, 0x09, 0xb5, 0x65, 0x03,
	0xf3, 0xbd, 0xf8, 0x2a, 0x43, 0x12, 0x32, 0xbd, 0x48, 0x25, 0xa0, 0x13, 0x75, 0xa1, 0xab, 0x44,
	0x3e, 0xa3, 0x17, 0xcb, 0xc6, 0xc5, 0x1d, 0xdf, 0x8b, 0x4a, 0x8d, 0x1c, 0x92, 0x0e, 0xc5, 0x4d,
	0x8f, 0xa3, 0x16, 0x6c, 0x4d, 0x43, 0xee, 0x8b, 0xa4, 0xbf, 0xae, 0x93, 0xe5, 0xb8, 0xf3, 0x67,
	0x09, 0x6a, 0x99, 0xd6, 0xf6, 0xf9, 0x39, 0x3c, 0x07, 0x48, 0x1a, 0xa5, 0xef, 0xc5, 0x33, 0xa9,
	0xc6, 0x11, 0xd3, 0xcb, 0x79, 0x95, 0xf3, 0x5e, 0xe8, 0x3b, 0xd8, 0x8c, 0x89, 0x71, 0x9b, 0x3f,
	0x28, 0xfa, 0x58, 0xb8, 0xa1, 0x82, 0x24, 0xbc, 0x8e, 0x05, 0x4f, 0x2e, 0xfd, 0x71, 0xc0, 0xbc,
	0x7b, 0x1f, 0x3b, 0xcf, 0xa0, 0xca, 0xfd, 0x71, 0x40, 0xc5, 0x7c, 0xc6, 0xe4, 0x24, 0xeb, 0x24,
	0x0d, 0x20, 0x1d, 0x36, 0xa7, 0x74, 0x71, 0x17, 0x52, 0x35, 0xc1, 0x3a, 0x49, 0x86, 0x9d, 0x7f,
	0x4a, 0xb0, 0x9f, 0x97, 0xe2, 0xc9, 0x16, 0x45, 0x0b, 0xb9, 0xfc, 0xa6, 0x89, 0x4b, 0xaf, 0x93,
	0xda, 0x32, 0x66, 0x7a, 0xc8, 0x84, 0x17, 0xf2, 0xc3, 0x8a, 0x3b, 0xd7, 0x94, 0x33, 0x27, 0xa5,
	0x7b, 0x8c, 0xbb, 0x33, 0x7f, 0x2a, 0xab, 0x56, 0x8e, 0x47, 0x8a, 0x78, 0x42, 0x39, 0x5b, 0xfa,
	0x0d, 0x52, 0x16, 0x6a, 0xc1, 0x86, 0x62, 0xc8, 0x0d, 0xab, 0x9f, 0xac, 0xe9, 0x25, 0x12, 0x47,
	0x90, 0x01, 0x0d, 0x2e, 0x8b, 0x76, 0x62, 0x4a, 0x45, 0x36, 0xd8, 0xcc, 0x1d, 0x28, 0x5a, 0x13,
	0x52, 0x57, 0x49, 0xaa, 0xaa, 0x8e, 0x0d, 0x47, 0xc5, 0x85, 0x26, 0x1d, 0x0b, 0xf5, 0xa0, 0x3c,
	0xe1, 0xe3, 0xf8, 0x5e, 0xb6, 0x57, 0x75, 0xef, 0x65, 0x5a, 0x44, 0x3e, 0x69, 0x7c, 0xa8, 0x1d,
	0xbf, 0x7c, 0x93, 0x50, 0xaf, 0x37, 0xe4, 0xd3, 0xf7, 0xff, 0x06, 0x00, 0x00, 0xff, 0xff, 0x92,
	0x1f, 0x8f, 0xf4, 0xd9, 0x0a, 0x00, 0x00,
}
