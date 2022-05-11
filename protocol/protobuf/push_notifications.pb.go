// Code generated by protoc-gen-go. DO NOT EDIT.
// source: push_notifications.proto

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

type PushNotificationRegistration_TokenType int32

const (
	PushNotificationRegistration_UNKNOWN_TOKEN_TYPE PushNotificationRegistration_TokenType = 0
	PushNotificationRegistration_APN_TOKEN          PushNotificationRegistration_TokenType = 1
	PushNotificationRegistration_FIREBASE_TOKEN     PushNotificationRegistration_TokenType = 2
)

var PushNotificationRegistration_TokenType_name = map[int32]string{
	0: "UNKNOWN_TOKEN_TYPE",
	1: "APN_TOKEN",
	2: "FIREBASE_TOKEN",
}

var PushNotificationRegistration_TokenType_value = map[string]int32{
	"UNKNOWN_TOKEN_TYPE": 0,
	"APN_TOKEN":          1,
	"FIREBASE_TOKEN":     2,
}

func (x PushNotificationRegistration_TokenType) String() string {
	return proto.EnumName(PushNotificationRegistration_TokenType_name, int32(x))
}

func (PushNotificationRegistration_TokenType) EnumDescriptor() ([]byte, []int) {
	return fileDescriptor_200acd86044eaa5d, []int{0, 0}
}

type PushNotificationRegistrationResponse_ErrorType int32

const (
	PushNotificationRegistrationResponse_UNKNOWN_ERROR_TYPE     PushNotificationRegistrationResponse_ErrorType = 0
	PushNotificationRegistrationResponse_MALFORMED_MESSAGE      PushNotificationRegistrationResponse_ErrorType = 1
	PushNotificationRegistrationResponse_VERSION_MISMATCH       PushNotificationRegistrationResponse_ErrorType = 2
	PushNotificationRegistrationResponse_UNSUPPORTED_TOKEN_TYPE PushNotificationRegistrationResponse_ErrorType = 3
	PushNotificationRegistrationResponse_INTERNAL_ERROR         PushNotificationRegistrationResponse_ErrorType = 4
)

var PushNotificationRegistrationResponse_ErrorType_name = map[int32]string{
	0: "UNKNOWN_ERROR_TYPE",
	1: "MALFORMED_MESSAGE",
	2: "VERSION_MISMATCH",
	3: "UNSUPPORTED_TOKEN_TYPE",
	4: "INTERNAL_ERROR",
}

var PushNotificationRegistrationResponse_ErrorType_value = map[string]int32{
	"UNKNOWN_ERROR_TYPE":     0,
	"MALFORMED_MESSAGE":      1,
	"VERSION_MISMATCH":       2,
	"UNSUPPORTED_TOKEN_TYPE": 3,
	"INTERNAL_ERROR":         4,
}

func (x PushNotificationRegistrationResponse_ErrorType) String() string {
	return proto.EnumName(PushNotificationRegistrationResponse_ErrorType_name, int32(x))
}

func (PushNotificationRegistrationResponse_ErrorType) EnumDescriptor() ([]byte, []int) {
	return fileDescriptor_200acd86044eaa5d, []int{1, 0}
}

type PushNotification_PushNotificationType int32

const (
	PushNotification_UNKNOWN_PUSH_NOTIFICATION_TYPE PushNotification_PushNotificationType = 0
	PushNotification_MESSAGE                        PushNotification_PushNotificationType = 1
	PushNotification_MENTION                        PushNotification_PushNotificationType = 2
	PushNotification_REQUEST_TO_JOIN_COMMUNITY      PushNotification_PushNotificationType = 3
)

var PushNotification_PushNotificationType_name = map[int32]string{
	0: "UNKNOWN_PUSH_NOTIFICATION_TYPE",
	1: "MESSAGE",
	2: "MENTION",
	3: "REQUEST_TO_JOIN_COMMUNITY",
}

var PushNotification_PushNotificationType_value = map[string]int32{
	"UNKNOWN_PUSH_NOTIFICATION_TYPE": 0,
	"MESSAGE":                        1,
	"MENTION":                        2,
	"REQUEST_TO_JOIN_COMMUNITY":      3,
}

func (x PushNotification_PushNotificationType) String() string {
	return proto.EnumName(PushNotification_PushNotificationType_name, int32(x))
}

func (PushNotification_PushNotificationType) EnumDescriptor() ([]byte, []int) {
	return fileDescriptor_200acd86044eaa5d, []int{6, 0}
}

type PushNotificationReport_ErrorType int32

const (
	PushNotificationReport_UNKNOWN_ERROR_TYPE PushNotificationReport_ErrorType = 0
	PushNotificationReport_WRONG_TOKEN        PushNotificationReport_ErrorType = 1
	PushNotificationReport_INTERNAL_ERROR     PushNotificationReport_ErrorType = 2
	PushNotificationReport_NOT_REGISTERED     PushNotificationReport_ErrorType = 3
)

var PushNotificationReport_ErrorType_name = map[int32]string{
	0: "UNKNOWN_ERROR_TYPE",
	1: "WRONG_TOKEN",
	2: "INTERNAL_ERROR",
	3: "NOT_REGISTERED",
}

var PushNotificationReport_ErrorType_value = map[string]int32{
	"UNKNOWN_ERROR_TYPE": 0,
	"WRONG_TOKEN":        1,
	"INTERNAL_ERROR":     2,
	"NOT_REGISTERED":     3,
}

func (x PushNotificationReport_ErrorType) String() string {
	return proto.EnumName(PushNotificationReport_ErrorType_name, int32(x))
}

func (PushNotificationReport_ErrorType) EnumDescriptor() ([]byte, []int) {
	return fileDescriptor_200acd86044eaa5d, []int{8, 0}
}

type PushNotificationRegistration struct {
	TokenType               PushNotificationRegistration_TokenType `protobuf:"varint,1,opt,name=token_type,json=tokenType,proto3,enum=protobuf.PushNotificationRegistration_TokenType" json:"token_type,omitempty"`
	DeviceToken             string                                 `protobuf:"bytes,2,opt,name=device_token,json=deviceToken,proto3" json:"device_token,omitempty"`
	InstallationId          string                                 `protobuf:"bytes,3,opt,name=installation_id,json=installationId,proto3" json:"installation_id,omitempty"`
	AccessToken             string                                 `protobuf:"bytes,4,opt,name=access_token,json=accessToken,proto3" json:"access_token,omitempty"`
	Enabled                 bool                                   `protobuf:"varint,5,opt,name=enabled,proto3" json:"enabled,omitempty"`
	Version                 uint64                                 `protobuf:"varint,6,opt,name=version,proto3" json:"version,omitempty"`
	AllowedKeyList          [][]byte                               `protobuf:"bytes,7,rep,name=allowed_key_list,json=allowedKeyList,proto3" json:"allowed_key_list,omitempty"`
	BlockedChatList         [][]byte                               `protobuf:"bytes,8,rep,name=blocked_chat_list,json=blockedChatList,proto3" json:"blocked_chat_list,omitempty"`
	Unregister              bool                                   `protobuf:"varint,9,opt,name=unregister,proto3" json:"unregister,omitempty"`
	Grant                   []byte                                 `protobuf:"bytes,10,opt,name=grant,proto3" json:"grant,omitempty"`
	AllowFromContactsOnly   bool                                   `protobuf:"varint,11,opt,name=allow_from_contacts_only,json=allowFromContactsOnly,proto3" json:"allow_from_contacts_only,omitempty"`
	ApnTopic                string                                 `protobuf:"bytes,12,opt,name=apn_topic,json=apnTopic,proto3" json:"apn_topic,omitempty"`
	BlockMentions           bool                                   `protobuf:"varint,13,opt,name=block_mentions,json=blockMentions,proto3" json:"block_mentions,omitempty"`
	AllowedMentionsChatList [][]byte                               `protobuf:"bytes,14,rep,name=allowed_mentions_chat_list,json=allowedMentionsChatList,proto3" json:"allowed_mentions_chat_list,omitempty"`
	XXX_NoUnkeyedLiteral    struct{}                               `json:"-"`
	XXX_unrecognized        []byte                                 `json:"-"`
	XXX_sizecache           int32                                  `json:"-"`
}

func (m *PushNotificationRegistration) Reset()         { *m = PushNotificationRegistration{} }
func (m *PushNotificationRegistration) String() string { return proto.CompactTextString(m) }
func (*PushNotificationRegistration) ProtoMessage()    {}
func (*PushNotificationRegistration) Descriptor() ([]byte, []int) {
	return fileDescriptor_200acd86044eaa5d, []int{0}
}

func (m *PushNotificationRegistration) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_PushNotificationRegistration.Unmarshal(m, b)
}
func (m *PushNotificationRegistration) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_PushNotificationRegistration.Marshal(b, m, deterministic)
}
func (m *PushNotificationRegistration) XXX_Merge(src proto.Message) {
	xxx_messageInfo_PushNotificationRegistration.Merge(m, src)
}
func (m *PushNotificationRegistration) XXX_Size() int {
	return xxx_messageInfo_PushNotificationRegistration.Size(m)
}
func (m *PushNotificationRegistration) XXX_DiscardUnknown() {
	xxx_messageInfo_PushNotificationRegistration.DiscardUnknown(m)
}

var xxx_messageInfo_PushNotificationRegistration proto.InternalMessageInfo

func (m *PushNotificationRegistration) GetTokenType() PushNotificationRegistration_TokenType {
	if m != nil {
		return m.TokenType
	}
	return PushNotificationRegistration_UNKNOWN_TOKEN_TYPE
}

func (m *PushNotificationRegistration) GetDeviceToken() string {
	if m != nil {
		return m.DeviceToken
	}
	return ""
}

func (m *PushNotificationRegistration) GetInstallationId() string {
	if m != nil {
		return m.InstallationId
	}
	return ""
}

func (m *PushNotificationRegistration) GetAccessToken() string {
	if m != nil {
		return m.AccessToken
	}
	return ""
}

func (m *PushNotificationRegistration) GetEnabled() bool {
	if m != nil {
		return m.Enabled
	}
	return false
}

func (m *PushNotificationRegistration) GetVersion() uint64 {
	if m != nil {
		return m.Version
	}
	return 0
}

func (m *PushNotificationRegistration) GetAllowedKeyList() [][]byte {
	if m != nil {
		return m.AllowedKeyList
	}
	return nil
}

func (m *PushNotificationRegistration) GetBlockedChatList() [][]byte {
	if m != nil {
		return m.BlockedChatList
	}
	return nil
}

func (m *PushNotificationRegistration) GetUnregister() bool {
	if m != nil {
		return m.Unregister
	}
	return false
}

func (m *PushNotificationRegistration) GetGrant() []byte {
	if m != nil {
		return m.Grant
	}
	return nil
}

func (m *PushNotificationRegistration) GetAllowFromContactsOnly() bool {
	if m != nil {
		return m.AllowFromContactsOnly
	}
	return false
}

func (m *PushNotificationRegistration) GetApnTopic() string {
	if m != nil {
		return m.ApnTopic
	}
	return ""
}

func (m *PushNotificationRegistration) GetBlockMentions() bool {
	if m != nil {
		return m.BlockMentions
	}
	return false
}

func (m *PushNotificationRegistration) GetAllowedMentionsChatList() [][]byte {
	if m != nil {
		return m.AllowedMentionsChatList
	}
	return nil
}

type PushNotificationRegistrationResponse struct {
	Success              bool                                           `protobuf:"varint,1,opt,name=success,proto3" json:"success,omitempty"`
	Error                PushNotificationRegistrationResponse_ErrorType `protobuf:"varint,2,opt,name=error,proto3,enum=protobuf.PushNotificationRegistrationResponse_ErrorType" json:"error,omitempty"`
	RequestId            []byte                                         `protobuf:"bytes,3,opt,name=request_id,json=requestId,proto3" json:"request_id,omitempty"`
	XXX_NoUnkeyedLiteral struct{}                                       `json:"-"`
	XXX_unrecognized     []byte                                         `json:"-"`
	XXX_sizecache        int32                                          `json:"-"`
}

func (m *PushNotificationRegistrationResponse) Reset()         { *m = PushNotificationRegistrationResponse{} }
func (m *PushNotificationRegistrationResponse) String() string { return proto.CompactTextString(m) }
func (*PushNotificationRegistrationResponse) ProtoMessage()    {}
func (*PushNotificationRegistrationResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_200acd86044eaa5d, []int{1}
}

func (m *PushNotificationRegistrationResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_PushNotificationRegistrationResponse.Unmarshal(m, b)
}
func (m *PushNotificationRegistrationResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_PushNotificationRegistrationResponse.Marshal(b, m, deterministic)
}
func (m *PushNotificationRegistrationResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_PushNotificationRegistrationResponse.Merge(m, src)
}
func (m *PushNotificationRegistrationResponse) XXX_Size() int {
	return xxx_messageInfo_PushNotificationRegistrationResponse.Size(m)
}
func (m *PushNotificationRegistrationResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_PushNotificationRegistrationResponse.DiscardUnknown(m)
}

var xxx_messageInfo_PushNotificationRegistrationResponse proto.InternalMessageInfo

func (m *PushNotificationRegistrationResponse) GetSuccess() bool {
	if m != nil {
		return m.Success
	}
	return false
}

func (m *PushNotificationRegistrationResponse) GetError() PushNotificationRegistrationResponse_ErrorType {
	if m != nil {
		return m.Error
	}
	return PushNotificationRegistrationResponse_UNKNOWN_ERROR_TYPE
}

func (m *PushNotificationRegistrationResponse) GetRequestId() []byte {
	if m != nil {
		return m.RequestId
	}
	return nil
}

type ContactCodeAdvertisement struct {
	PushNotificationInfo []*PushNotificationQueryInfo `protobuf:"bytes,1,rep,name=push_notification_info,json=pushNotificationInfo,proto3" json:"push_notification_info,omitempty"`
	ChatIdentity         *ChatIdentity                `protobuf:"bytes,2,opt,name=chat_identity,json=chatIdentity,proto3" json:"chat_identity,omitempty"`
	XXX_NoUnkeyedLiteral struct{}                     `json:"-"`
	XXX_unrecognized     []byte                       `json:"-"`
	XXX_sizecache        int32                        `json:"-"`
}

func (m *ContactCodeAdvertisement) Reset()         { *m = ContactCodeAdvertisement{} }
func (m *ContactCodeAdvertisement) String() string { return proto.CompactTextString(m) }
func (*ContactCodeAdvertisement) ProtoMessage()    {}
func (*ContactCodeAdvertisement) Descriptor() ([]byte, []int) {
	return fileDescriptor_200acd86044eaa5d, []int{2}
}

func (m *ContactCodeAdvertisement) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_ContactCodeAdvertisement.Unmarshal(m, b)
}
func (m *ContactCodeAdvertisement) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_ContactCodeAdvertisement.Marshal(b, m, deterministic)
}
func (m *ContactCodeAdvertisement) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ContactCodeAdvertisement.Merge(m, src)
}
func (m *ContactCodeAdvertisement) XXX_Size() int {
	return xxx_messageInfo_ContactCodeAdvertisement.Size(m)
}
func (m *ContactCodeAdvertisement) XXX_DiscardUnknown() {
	xxx_messageInfo_ContactCodeAdvertisement.DiscardUnknown(m)
}

var xxx_messageInfo_ContactCodeAdvertisement proto.InternalMessageInfo

func (m *ContactCodeAdvertisement) GetPushNotificationInfo() []*PushNotificationQueryInfo {
	if m != nil {
		return m.PushNotificationInfo
	}
	return nil
}

func (m *ContactCodeAdvertisement) GetChatIdentity() *ChatIdentity {
	if m != nil {
		return m.ChatIdentity
	}
	return nil
}

type PushNotificationQuery struct {
	PublicKeys           [][]byte `protobuf:"bytes,1,rep,name=public_keys,json=publicKeys,proto3" json:"public_keys,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *PushNotificationQuery) Reset()         { *m = PushNotificationQuery{} }
func (m *PushNotificationQuery) String() string { return proto.CompactTextString(m) }
func (*PushNotificationQuery) ProtoMessage()    {}
func (*PushNotificationQuery) Descriptor() ([]byte, []int) {
	return fileDescriptor_200acd86044eaa5d, []int{3}
}

func (m *PushNotificationQuery) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_PushNotificationQuery.Unmarshal(m, b)
}
func (m *PushNotificationQuery) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_PushNotificationQuery.Marshal(b, m, deterministic)
}
func (m *PushNotificationQuery) XXX_Merge(src proto.Message) {
	xxx_messageInfo_PushNotificationQuery.Merge(m, src)
}
func (m *PushNotificationQuery) XXX_Size() int {
	return xxx_messageInfo_PushNotificationQuery.Size(m)
}
func (m *PushNotificationQuery) XXX_DiscardUnknown() {
	xxx_messageInfo_PushNotificationQuery.DiscardUnknown(m)
}

var xxx_messageInfo_PushNotificationQuery proto.InternalMessageInfo

func (m *PushNotificationQuery) GetPublicKeys() [][]byte {
	if m != nil {
		return m.PublicKeys
	}
	return nil
}

type PushNotificationQueryInfo struct {
	AccessToken          string   `protobuf:"bytes,1,opt,name=access_token,json=accessToken,proto3" json:"access_token,omitempty"`
	InstallationId       string   `protobuf:"bytes,2,opt,name=installation_id,json=installationId,proto3" json:"installation_id,omitempty"`
	PublicKey            []byte   `protobuf:"bytes,3,opt,name=public_key,json=publicKey,proto3" json:"public_key,omitempty"`
	AllowedKeyList       [][]byte `protobuf:"bytes,4,rep,name=allowed_key_list,json=allowedKeyList,proto3" json:"allowed_key_list,omitempty"`
	Grant                []byte   `protobuf:"bytes,5,opt,name=grant,proto3" json:"grant,omitempty"`
	Version              uint64   `protobuf:"varint,6,opt,name=version,proto3" json:"version,omitempty"`
	ServerPublicKey      []byte   `protobuf:"bytes,7,opt,name=server_public_key,json=serverPublicKey,proto3" json:"server_public_key,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *PushNotificationQueryInfo) Reset()         { *m = PushNotificationQueryInfo{} }
func (m *PushNotificationQueryInfo) String() string { return proto.CompactTextString(m) }
func (*PushNotificationQueryInfo) ProtoMessage()    {}
func (*PushNotificationQueryInfo) Descriptor() ([]byte, []int) {
	return fileDescriptor_200acd86044eaa5d, []int{4}
}

func (m *PushNotificationQueryInfo) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_PushNotificationQueryInfo.Unmarshal(m, b)
}
func (m *PushNotificationQueryInfo) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_PushNotificationQueryInfo.Marshal(b, m, deterministic)
}
func (m *PushNotificationQueryInfo) XXX_Merge(src proto.Message) {
	xxx_messageInfo_PushNotificationQueryInfo.Merge(m, src)
}
func (m *PushNotificationQueryInfo) XXX_Size() int {
	return xxx_messageInfo_PushNotificationQueryInfo.Size(m)
}
func (m *PushNotificationQueryInfo) XXX_DiscardUnknown() {
	xxx_messageInfo_PushNotificationQueryInfo.DiscardUnknown(m)
}

var xxx_messageInfo_PushNotificationQueryInfo proto.InternalMessageInfo

func (m *PushNotificationQueryInfo) GetAccessToken() string {
	if m != nil {
		return m.AccessToken
	}
	return ""
}

func (m *PushNotificationQueryInfo) GetInstallationId() string {
	if m != nil {
		return m.InstallationId
	}
	return ""
}

func (m *PushNotificationQueryInfo) GetPublicKey() []byte {
	if m != nil {
		return m.PublicKey
	}
	return nil
}

func (m *PushNotificationQueryInfo) GetAllowedKeyList() [][]byte {
	if m != nil {
		return m.AllowedKeyList
	}
	return nil
}

func (m *PushNotificationQueryInfo) GetGrant() []byte {
	if m != nil {
		return m.Grant
	}
	return nil
}

func (m *PushNotificationQueryInfo) GetVersion() uint64 {
	if m != nil {
		return m.Version
	}
	return 0
}

func (m *PushNotificationQueryInfo) GetServerPublicKey() []byte {
	if m != nil {
		return m.ServerPublicKey
	}
	return nil
}

type PushNotificationQueryResponse struct {
	Info                 []*PushNotificationQueryInfo `protobuf:"bytes,1,rep,name=info,proto3" json:"info,omitempty"`
	MessageId            []byte                       `protobuf:"bytes,2,opt,name=message_id,json=messageId,proto3" json:"message_id,omitempty"`
	Success              bool                         `protobuf:"varint,3,opt,name=success,proto3" json:"success,omitempty"`
	XXX_NoUnkeyedLiteral struct{}                     `json:"-"`
	XXX_unrecognized     []byte                       `json:"-"`
	XXX_sizecache        int32                        `json:"-"`
}

func (m *PushNotificationQueryResponse) Reset()         { *m = PushNotificationQueryResponse{} }
func (m *PushNotificationQueryResponse) String() string { return proto.CompactTextString(m) }
func (*PushNotificationQueryResponse) ProtoMessage()    {}
func (*PushNotificationQueryResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_200acd86044eaa5d, []int{5}
}

func (m *PushNotificationQueryResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_PushNotificationQueryResponse.Unmarshal(m, b)
}
func (m *PushNotificationQueryResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_PushNotificationQueryResponse.Marshal(b, m, deterministic)
}
func (m *PushNotificationQueryResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_PushNotificationQueryResponse.Merge(m, src)
}
func (m *PushNotificationQueryResponse) XXX_Size() int {
	return xxx_messageInfo_PushNotificationQueryResponse.Size(m)
}
func (m *PushNotificationQueryResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_PushNotificationQueryResponse.DiscardUnknown(m)
}

var xxx_messageInfo_PushNotificationQueryResponse proto.InternalMessageInfo

func (m *PushNotificationQueryResponse) GetInfo() []*PushNotificationQueryInfo {
	if m != nil {
		return m.Info
	}
	return nil
}

func (m *PushNotificationQueryResponse) GetMessageId() []byte {
	if m != nil {
		return m.MessageId
	}
	return nil
}

func (m *PushNotificationQueryResponse) GetSuccess() bool {
	if m != nil {
		return m.Success
	}
	return false
}

type PushNotification struct {
	AccessToken          string                                `protobuf:"bytes,1,opt,name=access_token,json=accessToken,proto3" json:"access_token,omitempty"`
	ChatId               []byte                                `protobuf:"bytes,2,opt,name=chat_id,json=chatId,proto3" json:"chat_id,omitempty"`
	PublicKey            []byte                                `protobuf:"bytes,3,opt,name=public_key,json=publicKey,proto3" json:"public_key,omitempty"`
	InstallationId       string                                `protobuf:"bytes,4,opt,name=installation_id,json=installationId,proto3" json:"installation_id,omitempty"`
	Message              []byte                                `protobuf:"bytes,5,opt,name=message,proto3" json:"message,omitempty"`
	Type                 PushNotification_PushNotificationType `protobuf:"varint,6,opt,name=type,proto3,enum=protobuf.PushNotification_PushNotificationType" json:"type,omitempty"`
	Author               []byte                                `protobuf:"bytes,7,opt,name=author,proto3" json:"author,omitempty"`
	XXX_NoUnkeyedLiteral struct{}                              `json:"-"`
	XXX_unrecognized     []byte                                `json:"-"`
	XXX_sizecache        int32                                 `json:"-"`
}

func (m *PushNotification) Reset()         { *m = PushNotification{} }
func (m *PushNotification) String() string { return proto.CompactTextString(m) }
func (*PushNotification) ProtoMessage()    {}
func (*PushNotification) Descriptor() ([]byte, []int) {
	return fileDescriptor_200acd86044eaa5d, []int{6}
}

func (m *PushNotification) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_PushNotification.Unmarshal(m, b)
}
func (m *PushNotification) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_PushNotification.Marshal(b, m, deterministic)
}
func (m *PushNotification) XXX_Merge(src proto.Message) {
	xxx_messageInfo_PushNotification.Merge(m, src)
}
func (m *PushNotification) XXX_Size() int {
	return xxx_messageInfo_PushNotification.Size(m)
}
func (m *PushNotification) XXX_DiscardUnknown() {
	xxx_messageInfo_PushNotification.DiscardUnknown(m)
}

var xxx_messageInfo_PushNotification proto.InternalMessageInfo

func (m *PushNotification) GetAccessToken() string {
	if m != nil {
		return m.AccessToken
	}
	return ""
}

func (m *PushNotification) GetChatId() []byte {
	if m != nil {
		return m.ChatId
	}
	return nil
}

func (m *PushNotification) GetPublicKey() []byte {
	if m != nil {
		return m.PublicKey
	}
	return nil
}

func (m *PushNotification) GetInstallationId() string {
	if m != nil {
		return m.InstallationId
	}
	return ""
}

func (m *PushNotification) GetMessage() []byte {
	if m != nil {
		return m.Message
	}
	return nil
}

func (m *PushNotification) GetType() PushNotification_PushNotificationType {
	if m != nil {
		return m.Type
	}
	return PushNotification_UNKNOWN_PUSH_NOTIFICATION_TYPE
}

func (m *PushNotification) GetAuthor() []byte {
	if m != nil {
		return m.Author
	}
	return nil
}

type PushNotificationRequest struct {
	Requests             []*PushNotification `protobuf:"bytes,1,rep,name=requests,proto3" json:"requests,omitempty"`
	MessageId            []byte              `protobuf:"bytes,2,opt,name=message_id,json=messageId,proto3" json:"message_id,omitempty"`
	XXX_NoUnkeyedLiteral struct{}            `json:"-"`
	XXX_unrecognized     []byte              `json:"-"`
	XXX_sizecache        int32               `json:"-"`
}

func (m *PushNotificationRequest) Reset()         { *m = PushNotificationRequest{} }
func (m *PushNotificationRequest) String() string { return proto.CompactTextString(m) }
func (*PushNotificationRequest) ProtoMessage()    {}
func (*PushNotificationRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_200acd86044eaa5d, []int{7}
}

func (m *PushNotificationRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_PushNotificationRequest.Unmarshal(m, b)
}
func (m *PushNotificationRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_PushNotificationRequest.Marshal(b, m, deterministic)
}
func (m *PushNotificationRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_PushNotificationRequest.Merge(m, src)
}
func (m *PushNotificationRequest) XXX_Size() int {
	return xxx_messageInfo_PushNotificationRequest.Size(m)
}
func (m *PushNotificationRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_PushNotificationRequest.DiscardUnknown(m)
}

var xxx_messageInfo_PushNotificationRequest proto.InternalMessageInfo

func (m *PushNotificationRequest) GetRequests() []*PushNotification {
	if m != nil {
		return m.Requests
	}
	return nil
}

func (m *PushNotificationRequest) GetMessageId() []byte {
	if m != nil {
		return m.MessageId
	}
	return nil
}

type PushNotificationReport struct {
	Success              bool                             `protobuf:"varint,1,opt,name=success,proto3" json:"success,omitempty"`
	Error                PushNotificationReport_ErrorType `protobuf:"varint,2,opt,name=error,proto3,enum=protobuf.PushNotificationReport_ErrorType" json:"error,omitempty"`
	PublicKey            []byte                           `protobuf:"bytes,3,opt,name=public_key,json=publicKey,proto3" json:"public_key,omitempty"`
	InstallationId       string                           `protobuf:"bytes,4,opt,name=installation_id,json=installationId,proto3" json:"installation_id,omitempty"`
	XXX_NoUnkeyedLiteral struct{}                         `json:"-"`
	XXX_unrecognized     []byte                           `json:"-"`
	XXX_sizecache        int32                            `json:"-"`
}

func (m *PushNotificationReport) Reset()         { *m = PushNotificationReport{} }
func (m *PushNotificationReport) String() string { return proto.CompactTextString(m) }
func (*PushNotificationReport) ProtoMessage()    {}
func (*PushNotificationReport) Descriptor() ([]byte, []int) {
	return fileDescriptor_200acd86044eaa5d, []int{8}
}

func (m *PushNotificationReport) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_PushNotificationReport.Unmarshal(m, b)
}
func (m *PushNotificationReport) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_PushNotificationReport.Marshal(b, m, deterministic)
}
func (m *PushNotificationReport) XXX_Merge(src proto.Message) {
	xxx_messageInfo_PushNotificationReport.Merge(m, src)
}
func (m *PushNotificationReport) XXX_Size() int {
	return xxx_messageInfo_PushNotificationReport.Size(m)
}
func (m *PushNotificationReport) XXX_DiscardUnknown() {
	xxx_messageInfo_PushNotificationReport.DiscardUnknown(m)
}

var xxx_messageInfo_PushNotificationReport proto.InternalMessageInfo

func (m *PushNotificationReport) GetSuccess() bool {
	if m != nil {
		return m.Success
	}
	return false
}

func (m *PushNotificationReport) GetError() PushNotificationReport_ErrorType {
	if m != nil {
		return m.Error
	}
	return PushNotificationReport_UNKNOWN_ERROR_TYPE
}

func (m *PushNotificationReport) GetPublicKey() []byte {
	if m != nil {
		return m.PublicKey
	}
	return nil
}

func (m *PushNotificationReport) GetInstallationId() string {
	if m != nil {
		return m.InstallationId
	}
	return ""
}

type PushNotificationResponse struct {
	MessageId            []byte                    `protobuf:"bytes,1,opt,name=message_id,json=messageId,proto3" json:"message_id,omitempty"`
	Reports              []*PushNotificationReport `protobuf:"bytes,2,rep,name=reports,proto3" json:"reports,omitempty"`
	XXX_NoUnkeyedLiteral struct{}                  `json:"-"`
	XXX_unrecognized     []byte                    `json:"-"`
	XXX_sizecache        int32                     `json:"-"`
}

func (m *PushNotificationResponse) Reset()         { *m = PushNotificationResponse{} }
func (m *PushNotificationResponse) String() string { return proto.CompactTextString(m) }
func (*PushNotificationResponse) ProtoMessage()    {}
func (*PushNotificationResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_200acd86044eaa5d, []int{9}
}

func (m *PushNotificationResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_PushNotificationResponse.Unmarshal(m, b)
}
func (m *PushNotificationResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_PushNotificationResponse.Marshal(b, m, deterministic)
}
func (m *PushNotificationResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_PushNotificationResponse.Merge(m, src)
}
func (m *PushNotificationResponse) XXX_Size() int {
	return xxx_messageInfo_PushNotificationResponse.Size(m)
}
func (m *PushNotificationResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_PushNotificationResponse.DiscardUnknown(m)
}

var xxx_messageInfo_PushNotificationResponse proto.InternalMessageInfo

func (m *PushNotificationResponse) GetMessageId() []byte {
	if m != nil {
		return m.MessageId
	}
	return nil
}

func (m *PushNotificationResponse) GetReports() []*PushNotificationReport {
	if m != nil {
		return m.Reports
	}
	return nil
}

func init() {
	proto.RegisterEnum("protobuf.PushNotificationRegistration_TokenType", PushNotificationRegistration_TokenType_name, PushNotificationRegistration_TokenType_value)
	proto.RegisterEnum("protobuf.PushNotificationRegistrationResponse_ErrorType", PushNotificationRegistrationResponse_ErrorType_name, PushNotificationRegistrationResponse_ErrorType_value)
	proto.RegisterEnum("protobuf.PushNotification_PushNotificationType", PushNotification_PushNotificationType_name, PushNotification_PushNotificationType_value)
	proto.RegisterEnum("protobuf.PushNotificationReport_ErrorType", PushNotificationReport_ErrorType_name, PushNotificationReport_ErrorType_value)
	proto.RegisterType((*PushNotificationRegistration)(nil), "protobuf.PushNotificationRegistration")
	proto.RegisterType((*PushNotificationRegistrationResponse)(nil), "protobuf.PushNotificationRegistrationResponse")
	proto.RegisterType((*ContactCodeAdvertisement)(nil), "protobuf.ContactCodeAdvertisement")
	proto.RegisterType((*PushNotificationQuery)(nil), "protobuf.PushNotificationQuery")
	proto.RegisterType((*PushNotificationQueryInfo)(nil), "protobuf.PushNotificationQueryInfo")
	proto.RegisterType((*PushNotificationQueryResponse)(nil), "protobuf.PushNotificationQueryResponse")
	proto.RegisterType((*PushNotification)(nil), "protobuf.PushNotification")
	proto.RegisterType((*PushNotificationRequest)(nil), "protobuf.PushNotificationRequest")
	proto.RegisterType((*PushNotificationReport)(nil), "protobuf.PushNotificationReport")
	proto.RegisterType((*PushNotificationResponse)(nil), "protobuf.PushNotificationResponse")
}

func init() {
	proto.RegisterFile("push_notifications.proto", fileDescriptor_200acd86044eaa5d)
}

var fileDescriptor_200acd86044eaa5d = []byte{
	// 1072 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xac, 0x55, 0xcd, 0x6f, 0xe3, 0x44,
	0x14, 0xc7, 0x49, 0xda, 0x24, 0x2f, 0x69, 0x9a, 0x1d, 0xba, 0x59, 0x6f, 0xa1, 0x4b, 0x30, 0x20,
	0xa2, 0x1e, 0xb2, 0xa8, 0x48, 0xec, 0x8a, 0x5e, 0xc8, 0xa6, 0x6e, 0xd7, 0xb4, 0xb1, 0xb3, 0x13,
	0x87, 0x55, 0x11, 0xd2, 0xc8, 0xb5, 0xa7, 0xad, 0xd5, 0xd4, 0x36, 0x9e, 0x49, 0x51, 0x6e, 0x88,
	0x33, 0x17, 0xae, 0x9c, 0xf8, 0x1b, 0xf6, 0x2f, 0x44, 0x1e, 0x8f, 0x53, 0xb7, 0x49, 0x3f, 0x90,
	0x38, 0xc5, 0xef, 0xf7, 0x3e, 0xe6, 0x7d, 0xfe, 0x02, 0x6a, 0x34, 0x65, 0xe7, 0x24, 0x08, 0xb9,
	0x7f, 0xea, 0xbb, 0x0e, 0xf7, 0xc3, 0x80, 0x75, 0xa3, 0x38, 0xe4, 0x21, 0xaa, 0x88, 0x9f, 0x93,
	0xe9, 0xe9, 0xe6, 0xc7, 0xee, 0xb9, 0xc3, 0x89, 0xef, 0xd1, 0x80, 0xfb, 0x7c, 0x96, 0xaa, 0xb5,
	0x7f, 0x56, 0xe0, 0xd3, 0xe1, 0x94, 0x9d, 0x9b, 0x39, 0x57, 0x4c, 0xcf, 0x7c, 0xc6, 0x63, 0xf1,
	0x8d, 0x2c, 0x00, 0x1e, 0x5e, 0xd0, 0x80, 0xf0, 0x59, 0x44, 0x55, 0xa5, 0xad, 0x74, 0x1a, 0x3b,
	0xdf, 0x74, 0xb3, 0xa0, 0xdd, 0xfb, 0x7c, 0xbb, 0x76, 0xe2, 0x68, 0xcf, 0x22, 0x8a, 0xab, 0x3c,
	0xfb, 0x44, 0x9f, 0x43, 0xdd, 0xa3, 0x57, 0xbe, 0x4b, 0x89, 0xc0, 0xd4, 0x42, 0x5b, 0xe9, 0x54,
	0x71, 0x2d, 0xc5, 0x84, 0x07, 0xfa, 0x1a, 0xd6, 0xfd, 0x80, 0x71, 0x67, 0x32, 0x11, 0x71, 0x88,
	0xef, 0xa9, 0x45, 0x61, 0xd5, 0xc8, 0xc3, 0x86, 0x97, 0xc4, 0x72, 0x5c, 0x97, 0x32, 0x26, 0x63,
	0x95, 0xd2, 0x58, 0x29, 0x96, 0xc6, 0x52, 0xa1, 0x4c, 0x03, 0xe7, 0x64, 0x42, 0x3d, 0x75, 0xa5,
	0xad, 0x74, 0x2a, 0x38, 0x13, 0x13, 0xcd, 0x15, 0x8d, 0x99, 0x1f, 0x06, 0xea, 0x6a, 0x5b, 0xe9,
	0x94, 0x70, 0x26, 0xa2, 0x0e, 0x34, 0x9d, 0xc9, 0x24, 0xfc, 0x8d, 0x7a, 0xe4, 0x82, 0xce, 0xc8,
	0xc4, 0x67, 0x5c, 0x2d, 0xb7, 0x8b, 0x9d, 0x3a, 0x6e, 0x48, 0xfc, 0x90, 0xce, 0x8e, 0x7c, 0xc6,
	0xd1, 0x36, 0x3c, 0x39, 0x99, 0x84, 0xee, 0x05, 0xf5, 0x88, 0xe8, 0xae, 0x30, 0xad, 0x08, 0xd3,
	0x75, 0xa9, 0xe8, 0x9f, 0x3b, 0x5c, 0xd8, 0xbe, 0x00, 0x98, 0x06, 0xb1, 0xe8, 0x0f, 0x8d, 0xd5,
	0xaa, 0x48, 0x26, 0x87, 0xa0, 0x0d, 0x58, 0x39, 0x8b, 0x9d, 0x80, 0xab, 0xd0, 0x56, 0x3a, 0x75,
	0x9c, 0x0a, 0xe8, 0x15, 0xa8, 0xe2, 0x4d, 0x72, 0x1a, 0x87, 0x97, 0xc4, 0x0d, 0x03, 0xee, 0xb8,
	0x9c, 0x91, 0x30, 0x98, 0xcc, 0xd4, 0x9a, 0x88, 0xf1, 0x54, 0xe8, 0xf7, 0xe3, 0xf0, 0xb2, 0x2f,
	0xb5, 0x56, 0x30, 0x99, 0xa1, 0x4f, 0xa0, 0xea, 0x44, 0x01, 0xe1, 0x61, 0xe4, 0xbb, 0x6a, 0x5d,
	0x34, 0xa6, 0xe2, 0x44, 0x81, 0x9d, 0xc8, 0xe8, 0x2b, 0x68, 0x88, 0xf4, 0xc8, 0x65, 0xb2, 0x0d,
	0x61, 0xc0, 0xd4, 0x35, 0x11, 0x6b, 0x4d, 0xa0, 0x03, 0x09, 0xa2, 0x5d, 0xd8, 0xcc, 0x1a, 0x91,
	0x19, 0xe6, 0xea, 0x6c, 0x88, 0x3a, 0x9f, 0x49, 0x8b, 0xcc, 0x29, 0xab, 0x57, 0xdb, 0x87, 0xea,
	0x7c, 0x01, 0x50, 0x0b, 0xd0, 0xd8, 0x3c, 0x34, 0xad, 0xf7, 0x26, 0xb1, 0xad, 0x43, 0xdd, 0x24,
	0xf6, 0xf1, 0x50, 0x6f, 0x7e, 0x84, 0xd6, 0xa0, 0xda, 0x1b, 0x4a, 0xac, 0xa9, 0x20, 0x04, 0x8d,
	0x7d, 0x03, 0xeb, 0x6f, 0x7a, 0x23, 0x5d, 0x62, 0x05, 0xed, 0x43, 0x01, 0xbe, 0xbc, 0x6f, 0xcd,
	0x30, 0x65, 0x51, 0x18, 0x30, 0x9a, 0x0c, 0x94, 0x4d, 0xc5, 0xe8, 0xc5, 0x9e, 0x56, 0x70, 0x26,
	0x22, 0x13, 0x56, 0x68, 0x1c, 0x87, 0xb1, 0x58, 0xb6, 0xc6, 0xce, 0xeb, 0xc7, 0xed, 0x6f, 0x16,
	0xb8, 0xab, 0x27, 0xbe, 0x62, 0x8f, 0xd3, 0x30, 0x68, 0x0b, 0x20, 0xa6, 0xbf, 0x4e, 0x29, 0xe3,
	0xd9, 0x6e, 0xd6, 0x71, 0x55, 0x22, 0x86, 0xa7, 0xfd, 0xae, 0x40, 0x75, 0xee, 0x93, 0x2f, 0x5d,
	0xc7, 0xd8, 0xc2, 0x59, 0xe9, 0x4f, 0xe1, 0xc9, 0xa0, 0x77, 0xb4, 0x6f, 0xe1, 0x81, 0xbe, 0x47,
	0x06, 0xfa, 0x68, 0xd4, 0x3b, 0xd0, 0x9b, 0x0a, 0xda, 0x80, 0xe6, 0x4f, 0x3a, 0x1e, 0x19, 0x96,
	0x49, 0x06, 0xc6, 0x68, 0xd0, 0xb3, 0xfb, 0x6f, 0x9b, 0x05, 0xb4, 0x09, 0xad, 0xb1, 0x39, 0x1a,
	0x0f, 0x87, 0x16, 0xb6, 0xf5, 0xbd, 0x7c, 0x0f, 0x8b, 0x49, 0xd3, 0x0c, 0xd3, 0xd6, 0xb1, 0xd9,
	0x3b, 0x4a, 0x5f, 0x68, 0x96, 0xb4, 0x0f, 0x0a, 0xa8, 0x72, 0x1d, 0xfa, 0xa1, 0x47, 0x7b, 0xde,
	0x15, 0x8d, 0xb9, 0xcf, 0x68, 0x32, 0x46, 0x74, 0x0c, 0xad, 0x05, 0xbe, 0x20, 0x7e, 0x70, 0x1a,
	0xaa, 0x4a, 0xbb, 0xd8, 0xa9, 0xed, 0x7c, 0x71, 0x77, 0x7f, 0xde, 0x4d, 0x69, 0x3c, 0x33, 0x82,
	0xd3, 0x10, 0x6f, 0x44, 0xb7, 0x54, 0x09, 0x8a, 0x76, 0x61, 0xed, 0x06, 0xcd, 0x88, 0x8e, 0xd7,
	0x76, 0x5a, 0xd7, 0x11, 0x93, 0xfd, 0x30, 0xa4, 0x16, 0xd7, 0xdd, 0x9c, 0xa4, 0xbd, 0x86, 0xa7,
	0x4b, 0xdf, 0x43, 0x9f, 0x41, 0x2d, 0x9a, 0x9e, 0x4c, 0x7c, 0x37, 0xb9, 0x47, 0x26, 0xb2, 0xac,
	0x63, 0x48, 0xa1, 0x43, 0x3a, 0x63, 0xda, 0x9f, 0x05, 0x78, 0x7e, 0x67, 0xaa, 0x0b, 0x34, 0xa1,
	0x2c, 0xd2, 0xc4, 0x12, 0xca, 0x29, 0x2c, 0xa5, 0x9c, 0x2d, 0x80, 0xeb, 0x54, 0xb2, 0xd1, 0xcf,
	0x33, 0x59, 0x4a, 0x1d, 0xa5, 0xa5, 0xd4, 0x31, 0x3f, 0xf7, 0x95, 0xfc, 0xb9, 0xdf, 0x4d, 0x4a,
	0xdb, 0xf0, 0x84, 0xd1, 0xf8, 0x8a, 0xc6, 0x24, 0xf7, 0x7e, 0x59, 0xf8, 0xae, 0xa7, 0x8a, 0x61,
	0x96, 0x85, 0xf6, 0x97, 0x02, 0x5b, 0x4b, 0xdb, 0x31, 0xbf, 0x95, 0x57, 0x50, 0xfa, 0xaf, 0x03,
	0x17, 0x0e, 0x49, 0xfd, 0x97, 0x94, 0x31, 0xe7, 0x8c, 0x66, 0x3d, 0xaa, 0xe3, 0xaa, 0x44, 0x0c,
	0x2f, 0x7f, 0x83, 0xc5, 0x1b, 0x37, 0xa8, 0xfd, 0x51, 0x84, 0xe6, 0xed, 0xe0, 0x8f, 0x99, 0xcc,
	0x33, 0x28, 0xcb, 0x8d, 0x92, 0xaf, 0xad, 0xa6, 0x3b, 0xf3, 0xd0, 0x24, 0x96, 0x4c, 0xb4, 0xb4,
	0x74, 0xa2, 0x2a, 0x94, 0x65, 0xfe, 0x72, 0x14, 0x99, 0x88, 0xfa, 0x50, 0x12, 0xff, 0x7a, 0xab,
	0x82, 0x35, 0x5e, 0xde, 0xdd, 0xa4, 0x05, 0x40, 0x90, 0x85, 0x70, 0x46, 0x2d, 0x58, 0x75, 0xa6,
	0xfc, 0x3c, 0x8c, 0xe5, 0xb0, 0xa4, 0xa4, 0x31, 0xd8, 0x58, 0xe6, 0x85, 0x34, 0x78, 0x91, 0xd1,
	0xc5, 0x70, 0x3c, 0x7a, 0x4b, 0x4c, 0xcb, 0x36, 0xf6, 0x8d, 0x7e, 0xcf, 0x4e, 0x18, 0x41, 0x52,
	0x47, 0x0d, 0xca, 0xd7, 0x84, 0x21, 0x04, 0x33, 0x51, 0x37, 0x0b, 0x68, 0x0b, 0x9e, 0x63, 0xfd,
	0xdd, 0x58, 0x1f, 0xd9, 0xc4, 0xb6, 0xc8, 0x8f, 0x96, 0x61, 0x92, 0xbe, 0x35, 0x18, 0x8c, 0x4d,
	0xc3, 0x3e, 0x6e, 0x16, 0xb5, 0x08, 0x9e, 0x2d, 0x32, 0x9e, 0xa0, 0x2d, 0xf4, 0x1d, 0x54, 0x24,
	0x83, 0x31, 0xb9, 0x15, 0x9b, 0xf7, 0xd0, 0xe4, 0xdc, 0xf6, 0x81, 0x85, 0xd0, 0xfe, 0x2e, 0x40,
	0x6b, 0xf1, 0xc9, 0x28, 0x8c, 0xf9, 0x3d, 0x7c, 0xfd, 0xc3, 0x4d, 0xbe, 0xde, 0xbe, 0x8f, 0xaf,
	0x93, 0x50, 0x4b, 0x19, 0xfa, 0xff, 0x58, 0x0e, 0xed, 0x97, 0xc7, 0x30, 0xf9, 0x3a, 0xd4, 0xde,
	0x63, 0xcb, 0x3c, 0xc8, 0xff, 0x8d, 0xdd, 0x62, 0xe4, 0x42, 0x82, 0x99, 0x96, 0x4d, 0xb0, 0x7e,
	0x60, 0x8c, 0x6c, 0x1d, 0xeb, 0x7b, 0xcd, 0xa2, 0x36, 0x05, 0x75, 0xb1, 0x20, 0x79, 0xa1, 0x37,
	0xfb, 0xaa, 0xdc, 0x3e, 0xb4, 0xef, 0xa1, 0x1c, 0x8b, 0xda, 0x99, 0x5a, 0x10, 0xd3, 0x6a, 0x3f,
	0xd4, 0x24, 0x9c, 0x39, 0xbc, 0x59, 0xfb, 0xb9, 0xd6, 0x7d, 0xb9, 0x9b, 0x99, 0x9f, 0xac, 0x8a,
	0xaf, 0x6f, 0xff, 0x0d, 0x00, 0x00, 0xff, 0xff, 0x1d, 0x15, 0x64, 0x94, 0x45, 0x0a, 0x00, 0x00,
}
