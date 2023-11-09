// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.31.0
// 	protoc        v4.24.4
// source: contact.proto

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

type ContactRequestPropagatedState struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	LocalClock  uint64 `protobuf:"varint,1,opt,name=local_clock,json=localClock,proto3" json:"local_clock,omitempty"`
	LocalState  uint64 `protobuf:"varint,2,opt,name=local_state,json=localState,proto3" json:"local_state,omitempty"`
	RemoteClock uint64 `protobuf:"varint,3,opt,name=remote_clock,json=remoteClock,proto3" json:"remote_clock,omitempty"`
	RemoteState uint64 `protobuf:"varint,4,opt,name=remote_state,json=remoteState,proto3" json:"remote_state,omitempty"`
}

func (x *ContactRequestPropagatedState) Reset() {
	*x = ContactRequestPropagatedState{}
	if protoimpl.UnsafeEnabled {
		mi := &file_contact_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ContactRequestPropagatedState) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ContactRequestPropagatedState) ProtoMessage() {}

func (x *ContactRequestPropagatedState) ProtoReflect() protoreflect.Message {
	mi := &file_contact_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ContactRequestPropagatedState.ProtoReflect.Descriptor instead.
func (*ContactRequestPropagatedState) Descriptor() ([]byte, []int) {
	return file_contact_proto_rawDescGZIP(), []int{0}
}

func (x *ContactRequestPropagatedState) GetLocalClock() uint64 {
	if x != nil {
		return x.LocalClock
	}
	return 0
}

func (x *ContactRequestPropagatedState) GetLocalState() uint64 {
	if x != nil {
		return x.LocalState
	}
	return 0
}

func (x *ContactRequestPropagatedState) GetRemoteClock() uint64 {
	if x != nil {
		return x.RemoteClock
	}
	return 0
}

func (x *ContactRequestPropagatedState) GetRemoteState() uint64 {
	if x != nil {
		return x.RemoteState
	}
	return 0
}

type ContactUpdate struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Clock                         uint64                         `protobuf:"varint,1,opt,name=clock,proto3" json:"clock,omitempty"`
	EnsName                       string                         `protobuf:"bytes,2,opt,name=ens_name,json=ensName,proto3" json:"ens_name,omitempty"`
	ProfileImage                  string                         `protobuf:"bytes,3,opt,name=profile_image,json=profileImage,proto3" json:"profile_image,omitempty"`
	DisplayName                   string                         `protobuf:"bytes,4,opt,name=display_name,json=displayName,proto3" json:"display_name,omitempty"`
	ContactRequestClock           uint64                         `protobuf:"varint,5,opt,name=contact_request_clock,json=contactRequestClock,proto3" json:"contact_request_clock,omitempty"`
	ContactRequestPropagatedState *ContactRequestPropagatedState `protobuf:"bytes,6,opt,name=contact_request_propagated_state,json=contactRequestPropagatedState,proto3" json:"contact_request_propagated_state,omitempty"`
	PublicKey                     string                         `protobuf:"bytes,7,opt,name=public_key,json=publicKey,proto3" json:"public_key,omitempty"`
}

func (x *ContactUpdate) Reset() {
	*x = ContactUpdate{}
	if protoimpl.UnsafeEnabled {
		mi := &file_contact_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ContactUpdate) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ContactUpdate) ProtoMessage() {}

func (x *ContactUpdate) ProtoReflect() protoreflect.Message {
	mi := &file_contact_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ContactUpdate.ProtoReflect.Descriptor instead.
func (*ContactUpdate) Descriptor() ([]byte, []int) {
	return file_contact_proto_rawDescGZIP(), []int{1}
}

func (x *ContactUpdate) GetClock() uint64 {
	if x != nil {
		return x.Clock
	}
	return 0
}

func (x *ContactUpdate) GetEnsName() string {
	if x != nil {
		return x.EnsName
	}
	return ""
}

func (x *ContactUpdate) GetProfileImage() string {
	if x != nil {
		return x.ProfileImage
	}
	return ""
}

func (x *ContactUpdate) GetDisplayName() string {
	if x != nil {
		return x.DisplayName
	}
	return ""
}

func (x *ContactUpdate) GetContactRequestClock() uint64 {
	if x != nil {
		return x.ContactRequestClock
	}
	return 0
}

func (x *ContactUpdate) GetContactRequestPropagatedState() *ContactRequestPropagatedState {
	if x != nil {
		return x.ContactRequestPropagatedState
	}
	return nil
}

func (x *ContactUpdate) GetPublicKey() string {
	if x != nil {
		return x.PublicKey
	}
	return ""
}

type AcceptContactRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id    string `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
	Clock uint64 `protobuf:"varint,2,opt,name=clock,proto3" json:"clock,omitempty"`
}

func (x *AcceptContactRequest) Reset() {
	*x = AcceptContactRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_contact_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *AcceptContactRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*AcceptContactRequest) ProtoMessage() {}

func (x *AcceptContactRequest) ProtoReflect() protoreflect.Message {
	mi := &file_contact_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use AcceptContactRequest.ProtoReflect.Descriptor instead.
func (*AcceptContactRequest) Descriptor() ([]byte, []int) {
	return file_contact_proto_rawDescGZIP(), []int{2}
}

func (x *AcceptContactRequest) GetId() string {
	if x != nil {
		return x.Id
	}
	return ""
}

func (x *AcceptContactRequest) GetClock() uint64 {
	if x != nil {
		return x.Clock
	}
	return 0
}

type RetractContactRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id    string `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
	Clock uint64 `protobuf:"varint,2,opt,name=clock,proto3" json:"clock,omitempty"`
}

func (x *RetractContactRequest) Reset() {
	*x = RetractContactRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_contact_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *RetractContactRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*RetractContactRequest) ProtoMessage() {}

func (x *RetractContactRequest) ProtoReflect() protoreflect.Message {
	mi := &file_contact_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use RetractContactRequest.ProtoReflect.Descriptor instead.
func (*RetractContactRequest) Descriptor() ([]byte, []int) {
	return file_contact_proto_rawDescGZIP(), []int{3}
}

func (x *RetractContactRequest) GetId() string {
	if x != nil {
		return x.Id
	}
	return ""
}

func (x *RetractContactRequest) GetClock() uint64 {
	if x != nil {
		return x.Clock
	}
	return 0
}

var File_contact_proto protoreflect.FileDescriptor

var file_contact_proto_rawDesc = []byte{
	0x0a, 0x0d, 0x63, 0x6f, 0x6e, 0x74, 0x61, 0x63, 0x74, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12,
	0x08, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x22, 0xa7, 0x01, 0x0a, 0x1d, 0x43, 0x6f,
	0x6e, 0x74, 0x61, 0x63, 0x74, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x50, 0x72, 0x6f, 0x70,
	0x61, 0x67, 0x61, 0x74, 0x65, 0x64, 0x53, 0x74, 0x61, 0x74, 0x65, 0x12, 0x1f, 0x0a, 0x0b, 0x6c,
	0x6f, 0x63, 0x61, 0x6c, 0x5f, 0x63, 0x6c, 0x6f, 0x63, 0x6b, 0x18, 0x01, 0x20, 0x01, 0x28, 0x04,
	0x52, 0x0a, 0x6c, 0x6f, 0x63, 0x61, 0x6c, 0x43, 0x6c, 0x6f, 0x63, 0x6b, 0x12, 0x1f, 0x0a, 0x0b,
	0x6c, 0x6f, 0x63, 0x61, 0x6c, 0x5f, 0x73, 0x74, 0x61, 0x74, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28,
	0x04, 0x52, 0x0a, 0x6c, 0x6f, 0x63, 0x61, 0x6c, 0x53, 0x74, 0x61, 0x74, 0x65, 0x12, 0x21, 0x0a,
	0x0c, 0x72, 0x65, 0x6d, 0x6f, 0x74, 0x65, 0x5f, 0x63, 0x6c, 0x6f, 0x63, 0x6b, 0x18, 0x03, 0x20,
	0x01, 0x28, 0x04, 0x52, 0x0b, 0x72, 0x65, 0x6d, 0x6f, 0x74, 0x65, 0x43, 0x6c, 0x6f, 0x63, 0x6b,
	0x12, 0x21, 0x0a, 0x0c, 0x72, 0x65, 0x6d, 0x6f, 0x74, 0x65, 0x5f, 0x73, 0x74, 0x61, 0x74, 0x65,
	0x18, 0x04, 0x20, 0x01, 0x28, 0x04, 0x52, 0x0b, 0x72, 0x65, 0x6d, 0x6f, 0x74, 0x65, 0x53, 0x74,
	0x61, 0x74, 0x65, 0x22, 0xcd, 0x02, 0x0a, 0x0d, 0x43, 0x6f, 0x6e, 0x74, 0x61, 0x63, 0x74, 0x55,
	0x70, 0x64, 0x61, 0x74, 0x65, 0x12, 0x14, 0x0a, 0x05, 0x63, 0x6c, 0x6f, 0x63, 0x6b, 0x18, 0x01,
	0x20, 0x01, 0x28, 0x04, 0x52, 0x05, 0x63, 0x6c, 0x6f, 0x63, 0x6b, 0x12, 0x19, 0x0a, 0x08, 0x65,
	0x6e, 0x73, 0x5f, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x07, 0x65,
	0x6e, 0x73, 0x4e, 0x61, 0x6d, 0x65, 0x12, 0x23, 0x0a, 0x0d, 0x70, 0x72, 0x6f, 0x66, 0x69, 0x6c,
	0x65, 0x5f, 0x69, 0x6d, 0x61, 0x67, 0x65, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0c, 0x70,
	0x72, 0x6f, 0x66, 0x69, 0x6c, 0x65, 0x49, 0x6d, 0x61, 0x67, 0x65, 0x12, 0x21, 0x0a, 0x0c, 0x64,
	0x69, 0x73, 0x70, 0x6c, 0x61, 0x79, 0x5f, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x04, 0x20, 0x01, 0x28,
	0x09, 0x52, 0x0b, 0x64, 0x69, 0x73, 0x70, 0x6c, 0x61, 0x79, 0x4e, 0x61, 0x6d, 0x65, 0x12, 0x32,
	0x0a, 0x15, 0x63, 0x6f, 0x6e, 0x74, 0x61, 0x63, 0x74, 0x5f, 0x72, 0x65, 0x71, 0x75, 0x65, 0x73,
	0x74, 0x5f, 0x63, 0x6c, 0x6f, 0x63, 0x6b, 0x18, 0x05, 0x20, 0x01, 0x28, 0x04, 0x52, 0x13, 0x63,
	0x6f, 0x6e, 0x74, 0x61, 0x63, 0x74, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x43, 0x6c, 0x6f,
	0x63, 0x6b, 0x12, 0x70, 0x0a, 0x20, 0x63, 0x6f, 0x6e, 0x74, 0x61, 0x63, 0x74, 0x5f, 0x72, 0x65,
	0x71, 0x75, 0x65, 0x73, 0x74, 0x5f, 0x70, 0x72, 0x6f, 0x70, 0x61, 0x67, 0x61, 0x74, 0x65, 0x64,
	0x5f, 0x73, 0x74, 0x61, 0x74, 0x65, 0x18, 0x06, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x27, 0x2e, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x43, 0x6f, 0x6e, 0x74, 0x61, 0x63, 0x74, 0x52,
	0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x50, 0x72, 0x6f, 0x70, 0x61, 0x67, 0x61, 0x74, 0x65, 0x64,
	0x53, 0x74, 0x61, 0x74, 0x65, 0x52, 0x1d, 0x63, 0x6f, 0x6e, 0x74, 0x61, 0x63, 0x74, 0x52, 0x65,
	0x71, 0x75, 0x65, 0x73, 0x74, 0x50, 0x72, 0x6f, 0x70, 0x61, 0x67, 0x61, 0x74, 0x65, 0x64, 0x53,
	0x74, 0x61, 0x74, 0x65, 0x12, 0x1d, 0x0a, 0x0a, 0x70, 0x75, 0x62, 0x6c, 0x69, 0x63, 0x5f, 0x6b,
	0x65, 0x79, 0x18, 0x07, 0x20, 0x01, 0x28, 0x09, 0x52, 0x09, 0x70, 0x75, 0x62, 0x6c, 0x69, 0x63,
	0x4b, 0x65, 0x79, 0x22, 0x3c, 0x0a, 0x14, 0x41, 0x63, 0x63, 0x65, 0x70, 0x74, 0x43, 0x6f, 0x6e,
	0x74, 0x61, 0x63, 0x74, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x0e, 0x0a, 0x02, 0x69,
	0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x02, 0x69, 0x64, 0x12, 0x14, 0x0a, 0x05, 0x63,
	0x6c, 0x6f, 0x63, 0x6b, 0x18, 0x02, 0x20, 0x01, 0x28, 0x04, 0x52, 0x05, 0x63, 0x6c, 0x6f, 0x63,
	0x6b, 0x22, 0x3d, 0x0a, 0x15, 0x52, 0x65, 0x74, 0x72, 0x61, 0x63, 0x74, 0x43, 0x6f, 0x6e, 0x74,
	0x61, 0x63, 0x74, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x0e, 0x0a, 0x02, 0x69, 0x64,
	0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x02, 0x69, 0x64, 0x12, 0x14, 0x0a, 0x05, 0x63, 0x6c,
	0x6f, 0x63, 0x6b, 0x18, 0x02, 0x20, 0x01, 0x28, 0x04, 0x52, 0x05, 0x63, 0x6c, 0x6f, 0x63, 0x6b,
	0x42, 0x0d, 0x5a, 0x0b, 0x2e, 0x2f, 0x3b, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x62,
	0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_contact_proto_rawDescOnce sync.Once
	file_contact_proto_rawDescData = file_contact_proto_rawDesc
)

func file_contact_proto_rawDescGZIP() []byte {
	file_contact_proto_rawDescOnce.Do(func() {
		file_contact_proto_rawDescData = protoimpl.X.CompressGZIP(file_contact_proto_rawDescData)
	})
	return file_contact_proto_rawDescData
}

var file_contact_proto_msgTypes = make([]protoimpl.MessageInfo, 4)
var file_contact_proto_goTypes = []interface{}{
	(*ContactRequestPropagatedState)(nil), // 0: protobuf.ContactRequestPropagatedState
	(*ContactUpdate)(nil),                 // 1: protobuf.ContactUpdate
	(*AcceptContactRequest)(nil),          // 2: protobuf.AcceptContactRequest
	(*RetractContactRequest)(nil),         // 3: protobuf.RetractContactRequest
}
var file_contact_proto_depIdxs = []int32{
	0, // 0: protobuf.ContactUpdate.contact_request_propagated_state:type_name -> protobuf.ContactRequestPropagatedState
	1, // [1:1] is the sub-list for method output_type
	1, // [1:1] is the sub-list for method input_type
	1, // [1:1] is the sub-list for extension type_name
	1, // [1:1] is the sub-list for extension extendee
	0, // [0:1] is the sub-list for field type_name
}

func init() { file_contact_proto_init() }
func file_contact_proto_init() {
	if File_contact_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_contact_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ContactRequestPropagatedState); i {
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
		file_contact_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ContactUpdate); i {
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
		file_contact_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*AcceptContactRequest); i {
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
		file_contact_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*RetractContactRequest); i {
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
			RawDescriptor: file_contact_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   4,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_contact_proto_goTypes,
		DependencyIndexes: file_contact_proto_depIdxs,
		MessageInfos:      file_contact_proto_msgTypes,
	}.Build()
	File_contact_proto = out.File
	file_contact_proto_rawDesc = nil
	file_contact_proto_goTypes = nil
	file_contact_proto_depIdxs = nil
}
