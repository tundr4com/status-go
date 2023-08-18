// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.29.1
// 	protoc        v3.20.3
// source: command.proto

package protobuf

import (
	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
	reflect "reflect"
	sync "sync"
)

const (
	// Verify that this generated code is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(20 - protoimpl.MinVersion)
	// Verify that runtime/protoimpl is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(protoimpl.MaxVersion - 20)
)

type RequestAddressForTransaction struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Clock    uint64 `protobuf:"varint,1,opt,name=clock,proto3" json:"clock,omitempty"`
	Value    string `protobuf:"bytes,2,opt,name=value,proto3" json:"value,omitempty"`
	Contract string `protobuf:"bytes,3,opt,name=contract,proto3" json:"contract,omitempty"`
	ChatId   string `protobuf:"bytes,4,opt,name=chat_id,json=chatId,proto3" json:"chat_id,omitempty"`
}

func (x *RequestAddressForTransaction) Reset() {
	*x = RequestAddressForTransaction{}
	if protoimpl.UnsafeEnabled {
		mi := &file_command_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *RequestAddressForTransaction) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*RequestAddressForTransaction) ProtoMessage() {}

func (x *RequestAddressForTransaction) ProtoReflect() protoreflect.Message {
	mi := &file_command_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use RequestAddressForTransaction.ProtoReflect.Descriptor instead.
func (*RequestAddressForTransaction) Descriptor() ([]byte, []int) {
	return file_command_proto_rawDescGZIP(), []int{0}
}

func (x *RequestAddressForTransaction) GetClock() uint64 {
	if x != nil {
		return x.Clock
	}
	return 0
}

func (x *RequestAddressForTransaction) GetValue() string {
	if x != nil {
		return x.Value
	}
	return ""
}

func (x *RequestAddressForTransaction) GetContract() string {
	if x != nil {
		return x.Contract
	}
	return ""
}

func (x *RequestAddressForTransaction) GetChatId() string {
	if x != nil {
		return x.ChatId
	}
	return ""
}

type AcceptRequestAddressForTransaction struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Clock   uint64 `protobuf:"varint,1,opt,name=clock,proto3" json:"clock,omitempty"`
	Id      string `protobuf:"bytes,2,opt,name=id,proto3" json:"id,omitempty"`
	Address string `protobuf:"bytes,3,opt,name=address,proto3" json:"address,omitempty"`
	ChatId  string `protobuf:"bytes,4,opt,name=chat_id,json=chatId,proto3" json:"chat_id,omitempty"`
}

func (x *AcceptRequestAddressForTransaction) Reset() {
	*x = AcceptRequestAddressForTransaction{}
	if protoimpl.UnsafeEnabled {
		mi := &file_command_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *AcceptRequestAddressForTransaction) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*AcceptRequestAddressForTransaction) ProtoMessage() {}

func (x *AcceptRequestAddressForTransaction) ProtoReflect() protoreflect.Message {
	mi := &file_command_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use AcceptRequestAddressForTransaction.ProtoReflect.Descriptor instead.
func (*AcceptRequestAddressForTransaction) Descriptor() ([]byte, []int) {
	return file_command_proto_rawDescGZIP(), []int{1}
}

func (x *AcceptRequestAddressForTransaction) GetClock() uint64 {
	if x != nil {
		return x.Clock
	}
	return 0
}

func (x *AcceptRequestAddressForTransaction) GetId() string {
	if x != nil {
		return x.Id
	}
	return ""
}

func (x *AcceptRequestAddressForTransaction) GetAddress() string {
	if x != nil {
		return x.Address
	}
	return ""
}

func (x *AcceptRequestAddressForTransaction) GetChatId() string {
	if x != nil {
		return x.ChatId
	}
	return ""
}

type DeclineRequestAddressForTransaction struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Clock  uint64 `protobuf:"varint,1,opt,name=clock,proto3" json:"clock,omitempty"`
	Id     string `protobuf:"bytes,2,opt,name=id,proto3" json:"id,omitempty"`
	ChatId string `protobuf:"bytes,3,opt,name=chat_id,json=chatId,proto3" json:"chat_id,omitempty"`
}

func (x *DeclineRequestAddressForTransaction) Reset() {
	*x = DeclineRequestAddressForTransaction{}
	if protoimpl.UnsafeEnabled {
		mi := &file_command_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *DeclineRequestAddressForTransaction) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*DeclineRequestAddressForTransaction) ProtoMessage() {}

func (x *DeclineRequestAddressForTransaction) ProtoReflect() protoreflect.Message {
	mi := &file_command_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use DeclineRequestAddressForTransaction.ProtoReflect.Descriptor instead.
func (*DeclineRequestAddressForTransaction) Descriptor() ([]byte, []int) {
	return file_command_proto_rawDescGZIP(), []int{2}
}

func (x *DeclineRequestAddressForTransaction) GetClock() uint64 {
	if x != nil {
		return x.Clock
	}
	return 0
}

func (x *DeclineRequestAddressForTransaction) GetId() string {
	if x != nil {
		return x.Id
	}
	return ""
}

func (x *DeclineRequestAddressForTransaction) GetChatId() string {
	if x != nil {
		return x.ChatId
	}
	return ""
}

type DeclineRequestTransaction struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Clock  uint64 `protobuf:"varint,1,opt,name=clock,proto3" json:"clock,omitempty"`
	Id     string `protobuf:"bytes,2,opt,name=id,proto3" json:"id,omitempty"`
	ChatId string `protobuf:"bytes,3,opt,name=chat_id,json=chatId,proto3" json:"chat_id,omitempty"`
}

func (x *DeclineRequestTransaction) Reset() {
	*x = DeclineRequestTransaction{}
	if protoimpl.UnsafeEnabled {
		mi := &file_command_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *DeclineRequestTransaction) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*DeclineRequestTransaction) ProtoMessage() {}

func (x *DeclineRequestTransaction) ProtoReflect() protoreflect.Message {
	mi := &file_command_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use DeclineRequestTransaction.ProtoReflect.Descriptor instead.
func (*DeclineRequestTransaction) Descriptor() ([]byte, []int) {
	return file_command_proto_rawDescGZIP(), []int{3}
}

func (x *DeclineRequestTransaction) GetClock() uint64 {
	if x != nil {
		return x.Clock
	}
	return 0
}

func (x *DeclineRequestTransaction) GetId() string {
	if x != nil {
		return x.Id
	}
	return ""
}

func (x *DeclineRequestTransaction) GetChatId() string {
	if x != nil {
		return x.ChatId
	}
	return ""
}

type RequestTransaction struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Clock    uint64 `protobuf:"varint,1,opt,name=clock,proto3" json:"clock,omitempty"`
	Address  string `protobuf:"bytes,2,opt,name=address,proto3" json:"address,omitempty"`
	Value    string `protobuf:"bytes,3,opt,name=value,proto3" json:"value,omitempty"`
	Contract string `protobuf:"bytes,4,opt,name=contract,proto3" json:"contract,omitempty"`
	ChatId   string `protobuf:"bytes,5,opt,name=chat_id,json=chatId,proto3" json:"chat_id,omitempty"`
}

func (x *RequestTransaction) Reset() {
	*x = RequestTransaction{}
	if protoimpl.UnsafeEnabled {
		mi := &file_command_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *RequestTransaction) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*RequestTransaction) ProtoMessage() {}

func (x *RequestTransaction) ProtoReflect() protoreflect.Message {
	mi := &file_command_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use RequestTransaction.ProtoReflect.Descriptor instead.
func (*RequestTransaction) Descriptor() ([]byte, []int) {
	return file_command_proto_rawDescGZIP(), []int{4}
}

func (x *RequestTransaction) GetClock() uint64 {
	if x != nil {
		return x.Clock
	}
	return 0
}

func (x *RequestTransaction) GetAddress() string {
	if x != nil {
		return x.Address
	}
	return ""
}

func (x *RequestTransaction) GetValue() string {
	if x != nil {
		return x.Value
	}
	return ""
}

func (x *RequestTransaction) GetContract() string {
	if x != nil {
		return x.Contract
	}
	return ""
}

func (x *RequestTransaction) GetChatId() string {
	if x != nil {
		return x.ChatId
	}
	return ""
}

type SendTransaction struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Clock           uint64 `protobuf:"varint,1,opt,name=clock,proto3" json:"clock,omitempty"`
	Id              string `protobuf:"bytes,2,opt,name=id,proto3" json:"id,omitempty"`
	TransactionHash string `protobuf:"bytes,3,opt,name=transaction_hash,json=transactionHash,proto3" json:"transaction_hash,omitempty"`
	Signature       []byte `protobuf:"bytes,4,opt,name=signature,proto3" json:"signature,omitempty"`
	ChatId          string `protobuf:"bytes,5,opt,name=chat_id,json=chatId,proto3" json:"chat_id,omitempty"`
}

func (x *SendTransaction) Reset() {
	*x = SendTransaction{}
	if protoimpl.UnsafeEnabled {
		mi := &file_command_proto_msgTypes[5]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *SendTransaction) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*SendTransaction) ProtoMessage() {}

func (x *SendTransaction) ProtoReflect() protoreflect.Message {
	mi := &file_command_proto_msgTypes[5]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use SendTransaction.ProtoReflect.Descriptor instead.
func (*SendTransaction) Descriptor() ([]byte, []int) {
	return file_command_proto_rawDescGZIP(), []int{5}
}

func (x *SendTransaction) GetClock() uint64 {
	if x != nil {
		return x.Clock
	}
	return 0
}

func (x *SendTransaction) GetId() string {
	if x != nil {
		return x.Id
	}
	return ""
}

func (x *SendTransaction) GetTransactionHash() string {
	if x != nil {
		return x.TransactionHash
	}
	return ""
}

func (x *SendTransaction) GetSignature() []byte {
	if x != nil {
		return x.Signature
	}
	return nil
}

func (x *SendTransaction) GetChatId() string {
	if x != nil {
		return x.ChatId
	}
	return ""
}

var File_command_proto protoreflect.FileDescriptor

var file_command_proto_rawDesc = []byte{
	0x0a, 0x0d, 0x63, 0x6f, 0x6d, 0x6d, 0x61, 0x6e, 0x64, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12,
	0x08, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x22, 0x7f, 0x0a, 0x1c, 0x52, 0x65, 0x71,
	0x75, 0x65, 0x73, 0x74, 0x41, 0x64, 0x64, 0x72, 0x65, 0x73, 0x73, 0x46, 0x6f, 0x72, 0x54, 0x72,
	0x61, 0x6e, 0x73, 0x61, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x12, 0x14, 0x0a, 0x05, 0x63, 0x6c, 0x6f,
	0x63, 0x6b, 0x18, 0x01, 0x20, 0x01, 0x28, 0x04, 0x52, 0x05, 0x63, 0x6c, 0x6f, 0x63, 0x6b, 0x12,
	0x14, 0x0a, 0x05, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05,
	0x76, 0x61, 0x6c, 0x75, 0x65, 0x12, 0x1a, 0x0a, 0x08, 0x63, 0x6f, 0x6e, 0x74, 0x72, 0x61, 0x63,
	0x74, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x63, 0x6f, 0x6e, 0x74, 0x72, 0x61, 0x63,
	0x74, 0x12, 0x17, 0x0a, 0x07, 0x63, 0x68, 0x61, 0x74, 0x5f, 0x69, 0x64, 0x18, 0x04, 0x20, 0x01,
	0x28, 0x09, 0x52, 0x06, 0x63, 0x68, 0x61, 0x74, 0x49, 0x64, 0x22, 0x7d, 0x0a, 0x22, 0x41, 0x63,
	0x63, 0x65, 0x70, 0x74, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x41, 0x64, 0x64, 0x72, 0x65,
	0x73, 0x73, 0x46, 0x6f, 0x72, 0x54, 0x72, 0x61, 0x6e, 0x73, 0x61, 0x63, 0x74, 0x69, 0x6f, 0x6e,
	0x12, 0x14, 0x0a, 0x05, 0x63, 0x6c, 0x6f, 0x63, 0x6b, 0x18, 0x01, 0x20, 0x01, 0x28, 0x04, 0x52,
	0x05, 0x63, 0x6c, 0x6f, 0x63, 0x6b, 0x12, 0x0e, 0x0a, 0x02, 0x69, 0x64, 0x18, 0x02, 0x20, 0x01,
	0x28, 0x09, 0x52, 0x02, 0x69, 0x64, 0x12, 0x18, 0x0a, 0x07, 0x61, 0x64, 0x64, 0x72, 0x65, 0x73,
	0x73, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x07, 0x61, 0x64, 0x64, 0x72, 0x65, 0x73, 0x73,
	0x12, 0x17, 0x0a, 0x07, 0x63, 0x68, 0x61, 0x74, 0x5f, 0x69, 0x64, 0x18, 0x04, 0x20, 0x01, 0x28,
	0x09, 0x52, 0x06, 0x63, 0x68, 0x61, 0x74, 0x49, 0x64, 0x22, 0x64, 0x0a, 0x23, 0x44, 0x65, 0x63,
	0x6c, 0x69, 0x6e, 0x65, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x41, 0x64, 0x64, 0x72, 0x65,
	0x73, 0x73, 0x46, 0x6f, 0x72, 0x54, 0x72, 0x61, 0x6e, 0x73, 0x61, 0x63, 0x74, 0x69, 0x6f, 0x6e,
	0x12, 0x14, 0x0a, 0x05, 0x63, 0x6c, 0x6f, 0x63, 0x6b, 0x18, 0x01, 0x20, 0x01, 0x28, 0x04, 0x52,
	0x05, 0x63, 0x6c, 0x6f, 0x63, 0x6b, 0x12, 0x0e, 0x0a, 0x02, 0x69, 0x64, 0x18, 0x02, 0x20, 0x01,
	0x28, 0x09, 0x52, 0x02, 0x69, 0x64, 0x12, 0x17, 0x0a, 0x07, 0x63, 0x68, 0x61, 0x74, 0x5f, 0x69,
	0x64, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x06, 0x63, 0x68, 0x61, 0x74, 0x49, 0x64, 0x22,
	0x5a, 0x0a, 0x19, 0x44, 0x65, 0x63, 0x6c, 0x69, 0x6e, 0x65, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73,
	0x74, 0x54, 0x72, 0x61, 0x6e, 0x73, 0x61, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x12, 0x14, 0x0a, 0x05,
	0x63, 0x6c, 0x6f, 0x63, 0x6b, 0x18, 0x01, 0x20, 0x01, 0x28, 0x04, 0x52, 0x05, 0x63, 0x6c, 0x6f,
	0x63, 0x6b, 0x12, 0x0e, 0x0a, 0x02, 0x69, 0x64, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x02,
	0x69, 0x64, 0x12, 0x17, 0x0a, 0x07, 0x63, 0x68, 0x61, 0x74, 0x5f, 0x69, 0x64, 0x18, 0x03, 0x20,
	0x01, 0x28, 0x09, 0x52, 0x06, 0x63, 0x68, 0x61, 0x74, 0x49, 0x64, 0x22, 0x8f, 0x01, 0x0a, 0x12,
	0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x54, 0x72, 0x61, 0x6e, 0x73, 0x61, 0x63, 0x74, 0x69,
	0x6f, 0x6e, 0x12, 0x14, 0x0a, 0x05, 0x63, 0x6c, 0x6f, 0x63, 0x6b, 0x18, 0x01, 0x20, 0x01, 0x28,
	0x04, 0x52, 0x05, 0x63, 0x6c, 0x6f, 0x63, 0x6b, 0x12, 0x18, 0x0a, 0x07, 0x61, 0x64, 0x64, 0x72,
	0x65, 0x73, 0x73, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x07, 0x61, 0x64, 0x64, 0x72, 0x65,
	0x73, 0x73, 0x12, 0x14, 0x0a, 0x05, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x18, 0x03, 0x20, 0x01, 0x28,
	0x09, 0x52, 0x05, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x12, 0x1a, 0x0a, 0x08, 0x63, 0x6f, 0x6e, 0x74,
	0x72, 0x61, 0x63, 0x74, 0x18, 0x04, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x63, 0x6f, 0x6e, 0x74,
	0x72, 0x61, 0x63, 0x74, 0x12, 0x17, 0x0a, 0x07, 0x63, 0x68, 0x61, 0x74, 0x5f, 0x69, 0x64, 0x18,
	0x05, 0x20, 0x01, 0x28, 0x09, 0x52, 0x06, 0x63, 0x68, 0x61, 0x74, 0x49, 0x64, 0x22, 0x99, 0x01,
	0x0a, 0x0f, 0x53, 0x65, 0x6e, 0x64, 0x54, 0x72, 0x61, 0x6e, 0x73, 0x61, 0x63, 0x74, 0x69, 0x6f,
	0x6e, 0x12, 0x14, 0x0a, 0x05, 0x63, 0x6c, 0x6f, 0x63, 0x6b, 0x18, 0x01, 0x20, 0x01, 0x28, 0x04,
	0x52, 0x05, 0x63, 0x6c, 0x6f, 0x63, 0x6b, 0x12, 0x0e, 0x0a, 0x02, 0x69, 0x64, 0x18, 0x02, 0x20,
	0x01, 0x28, 0x09, 0x52, 0x02, 0x69, 0x64, 0x12, 0x29, 0x0a, 0x10, 0x74, 0x72, 0x61, 0x6e, 0x73,
	0x61, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x5f, 0x68, 0x61, 0x73, 0x68, 0x18, 0x03, 0x20, 0x01, 0x28,
	0x09, 0x52, 0x0f, 0x74, 0x72, 0x61, 0x6e, 0x73, 0x61, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x48, 0x61,
	0x73, 0x68, 0x12, 0x1c, 0x0a, 0x09, 0x73, 0x69, 0x67, 0x6e, 0x61, 0x74, 0x75, 0x72, 0x65, 0x18,
	0x04, 0x20, 0x01, 0x28, 0x0c, 0x52, 0x09, 0x73, 0x69, 0x67, 0x6e, 0x61, 0x74, 0x75, 0x72, 0x65,
	0x12, 0x17, 0x0a, 0x07, 0x63, 0x68, 0x61, 0x74, 0x5f, 0x69, 0x64, 0x18, 0x05, 0x20, 0x01, 0x28,
	0x09, 0x52, 0x06, 0x63, 0x68, 0x61, 0x74, 0x49, 0x64, 0x42, 0x0d, 0x5a, 0x0b, 0x2e, 0x2f, 0x3b,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_command_proto_rawDescOnce sync.Once
	file_command_proto_rawDescData = file_command_proto_rawDesc
)

func file_command_proto_rawDescGZIP() []byte {
	file_command_proto_rawDescOnce.Do(func() {
		file_command_proto_rawDescData = protoimpl.X.CompressGZIP(file_command_proto_rawDescData)
	})
	return file_command_proto_rawDescData
}

var file_command_proto_msgTypes = make([]protoimpl.MessageInfo, 6)
var file_command_proto_goTypes = []interface{}{
	(*RequestAddressForTransaction)(nil),        // 0: protobuf.RequestAddressForTransaction
	(*AcceptRequestAddressForTransaction)(nil),  // 1: protobuf.AcceptRequestAddressForTransaction
	(*DeclineRequestAddressForTransaction)(nil), // 2: protobuf.DeclineRequestAddressForTransaction
	(*DeclineRequestTransaction)(nil),           // 3: protobuf.DeclineRequestTransaction
	(*RequestTransaction)(nil),                  // 4: protobuf.RequestTransaction
	(*SendTransaction)(nil),                     // 5: protobuf.SendTransaction
}
var file_command_proto_depIdxs = []int32{
	0, // [0:0] is the sub-list for method output_type
	0, // [0:0] is the sub-list for method input_type
	0, // [0:0] is the sub-list for extension type_name
	0, // [0:0] is the sub-list for extension extendee
	0, // [0:0] is the sub-list for field type_name
}

func init() { file_command_proto_init() }
func file_command_proto_init() {
	if File_command_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_command_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*RequestAddressForTransaction); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_command_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*AcceptRequestAddressForTransaction); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_command_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*DeclineRequestAddressForTransaction); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_command_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*DeclineRequestTransaction); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_command_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*RequestTransaction); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_command_proto_msgTypes[5].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*SendTransaction); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
	}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_command_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   6,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_command_proto_goTypes,
		DependencyIndexes: file_command_proto_depIdxs,
		MessageInfos:      file_command_proto_msgTypes,
	}.Build()
	File_command_proto = out.File
	file_command_proto_rawDesc = nil
	file_command_proto_goTypes = nil
	file_command_proto_depIdxs = nil
}
