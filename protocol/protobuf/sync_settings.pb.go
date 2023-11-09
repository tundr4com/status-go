// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.31.0
// 	protoc        v4.24.4
// source: sync_settings.proto

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

type SyncSetting_Type int32

const (
	SyncSetting_UNKNOWN                     SyncSetting_Type = 0
	SyncSetting_CURRENCY                    SyncSetting_Type = 1
	SyncSetting_GIF_RECENTS                 SyncSetting_Type = 2
	SyncSetting_GIF_FAVOURITES              SyncSetting_Type = 3
	SyncSetting_MESSAGES_FROM_CONTACTS_ONLY SyncSetting_Type = 4
	SyncSetting_PREFERRED_NAME              SyncSetting_Type = 5
	SyncSetting_PREVIEW_PRIVACY             SyncSetting_Type = 6
	SyncSetting_PROFILE_PICTURES_SHOW_TO    SyncSetting_Type = 7
	SyncSetting_PROFILE_PICTURES_VISIBILITY SyncSetting_Type = 8
	SyncSetting_SEND_STATUS_UPDATES         SyncSetting_Type = 9
	SyncSetting_STICKERS_PACKS_INSTALLED    SyncSetting_Type = 10
	SyncSetting_STICKERS_PACKS_PENDING      SyncSetting_Type = 11
	SyncSetting_STICKERS_RECENT_STICKERS    SyncSetting_Type = 12
	SyncSetting_DISPLAY_NAME                SyncSetting_Type = 13
	SyncSetting_BIO                         SyncSetting_Type = 14
	SyncSetting_MNEMONIC_REMOVED            SyncSetting_Type = 15
	SyncSetting_URL_UNFURLING_MODE          SyncSetting_Type = 18
)

// Enum value maps for SyncSetting_Type.
var (
	SyncSetting_Type_name = map[int32]string{
		0:  "UNKNOWN",
		1:  "CURRENCY",
		2:  "GIF_RECENTS",
		3:  "GIF_FAVOURITES",
		4:  "MESSAGES_FROM_CONTACTS_ONLY",
		5:  "PREFERRED_NAME",
		6:  "PREVIEW_PRIVACY",
		7:  "PROFILE_PICTURES_SHOW_TO",
		8:  "PROFILE_PICTURES_VISIBILITY",
		9:  "SEND_STATUS_UPDATES",
		10: "STICKERS_PACKS_INSTALLED",
		11: "STICKERS_PACKS_PENDING",
		12: "STICKERS_RECENT_STICKERS",
		13: "DISPLAY_NAME",
		14: "BIO",
		15: "MNEMONIC_REMOVED",
		18: "URL_UNFURLING_MODE",
	}
	SyncSetting_Type_value = map[string]int32{
		"UNKNOWN":                     0,
		"CURRENCY":                    1,
		"GIF_RECENTS":                 2,
		"GIF_FAVOURITES":              3,
		"MESSAGES_FROM_CONTACTS_ONLY": 4,
		"PREFERRED_NAME":              5,
		"PREVIEW_PRIVACY":             6,
		"PROFILE_PICTURES_SHOW_TO":    7,
		"PROFILE_PICTURES_VISIBILITY": 8,
		"SEND_STATUS_UPDATES":         9,
		"STICKERS_PACKS_INSTALLED":    10,
		"STICKERS_PACKS_PENDING":      11,
		"STICKERS_RECENT_STICKERS":    12,
		"DISPLAY_NAME":                13,
		"BIO":                         14,
		"MNEMONIC_REMOVED":            15,
		"URL_UNFURLING_MODE":          18,
	}
)

func (x SyncSetting_Type) Enum() *SyncSetting_Type {
	p := new(SyncSetting_Type)
	*p = x
	return p
}

func (x SyncSetting_Type) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (SyncSetting_Type) Descriptor() protoreflect.EnumDescriptor {
	return file_sync_settings_proto_enumTypes[0].Descriptor()
}

func (SyncSetting_Type) Type() protoreflect.EnumType {
	return &file_sync_settings_proto_enumTypes[0]
}

func (x SyncSetting_Type) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use SyncSetting_Type.Descriptor instead.
func (SyncSetting_Type) EnumDescriptor() ([]byte, []int) {
	return file_sync_settings_proto_rawDescGZIP(), []int{0, 0}
}

type SyncSetting struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Type  SyncSetting_Type `protobuf:"varint,1,opt,name=type,proto3,enum=protobuf.SyncSetting_Type" json:"type,omitempty"`
	Clock uint64           `protobuf:"varint,2,opt,name=clock,proto3" json:"clock,omitempty"`
	// Types that are assignable to Value:
	//
	//	*SyncSetting_ValueString
	//	*SyncSetting_ValueBytes
	//	*SyncSetting_ValueBool
	//	*SyncSetting_ValueInt64
	Value isSyncSetting_Value `protobuf_oneof:"value"`
}

func (x *SyncSetting) Reset() {
	*x = SyncSetting{}
	if protoimpl.UnsafeEnabled {
		mi := &file_sync_settings_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *SyncSetting) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*SyncSetting) ProtoMessage() {}

func (x *SyncSetting) ProtoReflect() protoreflect.Message {
	mi := &file_sync_settings_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use SyncSetting.ProtoReflect.Descriptor instead.
func (*SyncSetting) Descriptor() ([]byte, []int) {
	return file_sync_settings_proto_rawDescGZIP(), []int{0}
}

func (x *SyncSetting) GetType() SyncSetting_Type {
	if x != nil {
		return x.Type
	}
	return SyncSetting_UNKNOWN
}

func (x *SyncSetting) GetClock() uint64 {
	if x != nil {
		return x.Clock
	}
	return 0
}

func (m *SyncSetting) GetValue() isSyncSetting_Value {
	if m != nil {
		return m.Value
	}
	return nil
}

func (x *SyncSetting) GetValueString() string {
	if x, ok := x.GetValue().(*SyncSetting_ValueString); ok {
		return x.ValueString
	}
	return ""
}

func (x *SyncSetting) GetValueBytes() []byte {
	if x, ok := x.GetValue().(*SyncSetting_ValueBytes); ok {
		return x.ValueBytes
	}
	return nil
}

func (x *SyncSetting) GetValueBool() bool {
	if x, ok := x.GetValue().(*SyncSetting_ValueBool); ok {
		return x.ValueBool
	}
	return false
}

func (x *SyncSetting) GetValueInt64() int64 {
	if x, ok := x.GetValue().(*SyncSetting_ValueInt64); ok {
		return x.ValueInt64
	}
	return 0
}

type isSyncSetting_Value interface {
	isSyncSetting_Value()
}

type SyncSetting_ValueString struct {
	ValueString string `protobuf:"bytes,3,opt,name=value_string,json=valueString,proto3,oneof"`
}

type SyncSetting_ValueBytes struct {
	ValueBytes []byte `protobuf:"bytes,4,opt,name=value_bytes,json=valueBytes,proto3,oneof"`
}

type SyncSetting_ValueBool struct {
	ValueBool bool `protobuf:"varint,5,opt,name=value_bool,json=valueBool,proto3,oneof"`
}

type SyncSetting_ValueInt64 struct {
	ValueInt64 int64 `protobuf:"varint,6,opt,name=value_int64,json=valueInt64,proto3,oneof"`
}

func (*SyncSetting_ValueString) isSyncSetting_Value() {}

func (*SyncSetting_ValueBytes) isSyncSetting_Value() {}

func (*SyncSetting_ValueBool) isSyncSetting_Value() {}

func (*SyncSetting_ValueInt64) isSyncSetting_Value() {}

var File_sync_settings_proto protoreflect.FileDescriptor

var file_sync_settings_proto_rawDesc = []byte{
	0x0a, 0x13, 0x73, 0x79, 0x6e, 0x63, 0x5f, 0x73, 0x65, 0x74, 0x74, 0x69, 0x6e, 0x67, 0x73, 0x2e,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x08, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x22,
	0xaa, 0x05, 0x0a, 0x0b, 0x53, 0x79, 0x6e, 0x63, 0x53, 0x65, 0x74, 0x74, 0x69, 0x6e, 0x67, 0x12,
	0x2e, 0x0a, 0x04, 0x74, 0x79, 0x70, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0e, 0x32, 0x1a, 0x2e,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x53, 0x79, 0x6e, 0x63, 0x53, 0x65, 0x74,
	0x74, 0x69, 0x6e, 0x67, 0x2e, 0x54, 0x79, 0x70, 0x65, 0x52, 0x04, 0x74, 0x79, 0x70, 0x65, 0x12,
	0x14, 0x0a, 0x05, 0x63, 0x6c, 0x6f, 0x63, 0x6b, 0x18, 0x02, 0x20, 0x01, 0x28, 0x04, 0x52, 0x05,
	0x63, 0x6c, 0x6f, 0x63, 0x6b, 0x12, 0x23, 0x0a, 0x0c, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x5f, 0x73,
	0x74, 0x72, 0x69, 0x6e, 0x67, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x48, 0x00, 0x52, 0x0b, 0x76,
	0x61, 0x6c, 0x75, 0x65, 0x53, 0x74, 0x72, 0x69, 0x6e, 0x67, 0x12, 0x21, 0x0a, 0x0b, 0x76, 0x61,
	0x6c, 0x75, 0x65, 0x5f, 0x62, 0x79, 0x74, 0x65, 0x73, 0x18, 0x04, 0x20, 0x01, 0x28, 0x0c, 0x48,
	0x00, 0x52, 0x0a, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x42, 0x79, 0x74, 0x65, 0x73, 0x12, 0x1f, 0x0a,
	0x0a, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x5f, 0x62, 0x6f, 0x6f, 0x6c, 0x18, 0x05, 0x20, 0x01, 0x28,
	0x08, 0x48, 0x00, 0x52, 0x09, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x42, 0x6f, 0x6f, 0x6c, 0x12, 0x21,
	0x0a, 0x0b, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x5f, 0x69, 0x6e, 0x74, 0x36, 0x34, 0x18, 0x06, 0x20,
	0x01, 0x28, 0x03, 0x48, 0x00, 0x52, 0x0a, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x49, 0x6e, 0x74, 0x36,
	0x34, 0x22, 0xbf, 0x03, 0x0a, 0x04, 0x54, 0x79, 0x70, 0x65, 0x12, 0x0b, 0x0a, 0x07, 0x55, 0x4e,
	0x4b, 0x4e, 0x4f, 0x57, 0x4e, 0x10, 0x00, 0x12, 0x0c, 0x0a, 0x08, 0x43, 0x55, 0x52, 0x52, 0x45,
	0x4e, 0x43, 0x59, 0x10, 0x01, 0x12, 0x0f, 0x0a, 0x0b, 0x47, 0x49, 0x46, 0x5f, 0x52, 0x45, 0x43,
	0x45, 0x4e, 0x54, 0x53, 0x10, 0x02, 0x12, 0x12, 0x0a, 0x0e, 0x47, 0x49, 0x46, 0x5f, 0x46, 0x41,
	0x56, 0x4f, 0x55, 0x52, 0x49, 0x54, 0x45, 0x53, 0x10, 0x03, 0x12, 0x1f, 0x0a, 0x1b, 0x4d, 0x45,
	0x53, 0x53, 0x41, 0x47, 0x45, 0x53, 0x5f, 0x46, 0x52, 0x4f, 0x4d, 0x5f, 0x43, 0x4f, 0x4e, 0x54,
	0x41, 0x43, 0x54, 0x53, 0x5f, 0x4f, 0x4e, 0x4c, 0x59, 0x10, 0x04, 0x12, 0x12, 0x0a, 0x0e, 0x50,
	0x52, 0x45, 0x46, 0x45, 0x52, 0x52, 0x45, 0x44, 0x5f, 0x4e, 0x41, 0x4d, 0x45, 0x10, 0x05, 0x12,
	0x13, 0x0a, 0x0f, 0x50, 0x52, 0x45, 0x56, 0x49, 0x45, 0x57, 0x5f, 0x50, 0x52, 0x49, 0x56, 0x41,
	0x43, 0x59, 0x10, 0x06, 0x12, 0x1c, 0x0a, 0x18, 0x50, 0x52, 0x4f, 0x46, 0x49, 0x4c, 0x45, 0x5f,
	0x50, 0x49, 0x43, 0x54, 0x55, 0x52, 0x45, 0x53, 0x5f, 0x53, 0x48, 0x4f, 0x57, 0x5f, 0x54, 0x4f,
	0x10, 0x07, 0x12, 0x1f, 0x0a, 0x1b, 0x50, 0x52, 0x4f, 0x46, 0x49, 0x4c, 0x45, 0x5f, 0x50, 0x49,
	0x43, 0x54, 0x55, 0x52, 0x45, 0x53, 0x5f, 0x56, 0x49, 0x53, 0x49, 0x42, 0x49, 0x4c, 0x49, 0x54,
	0x59, 0x10, 0x08, 0x12, 0x17, 0x0a, 0x13, 0x53, 0x45, 0x4e, 0x44, 0x5f, 0x53, 0x54, 0x41, 0x54,
	0x55, 0x53, 0x5f, 0x55, 0x50, 0x44, 0x41, 0x54, 0x45, 0x53, 0x10, 0x09, 0x12, 0x1c, 0x0a, 0x18,
	0x53, 0x54, 0x49, 0x43, 0x4b, 0x45, 0x52, 0x53, 0x5f, 0x50, 0x41, 0x43, 0x4b, 0x53, 0x5f, 0x49,
	0x4e, 0x53, 0x54, 0x41, 0x4c, 0x4c, 0x45, 0x44, 0x10, 0x0a, 0x12, 0x1a, 0x0a, 0x16, 0x53, 0x54,
	0x49, 0x43, 0x4b, 0x45, 0x52, 0x53, 0x5f, 0x50, 0x41, 0x43, 0x4b, 0x53, 0x5f, 0x50, 0x45, 0x4e,
	0x44, 0x49, 0x4e, 0x47, 0x10, 0x0b, 0x12, 0x1c, 0x0a, 0x18, 0x53, 0x54, 0x49, 0x43, 0x4b, 0x45,
	0x52, 0x53, 0x5f, 0x52, 0x45, 0x43, 0x45, 0x4e, 0x54, 0x5f, 0x53, 0x54, 0x49, 0x43, 0x4b, 0x45,
	0x52, 0x53, 0x10, 0x0c, 0x12, 0x10, 0x0a, 0x0c, 0x44, 0x49, 0x53, 0x50, 0x4c, 0x41, 0x59, 0x5f,
	0x4e, 0x41, 0x4d, 0x45, 0x10, 0x0d, 0x12, 0x07, 0x0a, 0x03, 0x42, 0x49, 0x4f, 0x10, 0x0e, 0x12,
	0x14, 0x0a, 0x10, 0x4d, 0x4e, 0x45, 0x4d, 0x4f, 0x4e, 0x49, 0x43, 0x5f, 0x52, 0x45, 0x4d, 0x4f,
	0x56, 0x45, 0x44, 0x10, 0x0f, 0x12, 0x16, 0x0a, 0x12, 0x55, 0x52, 0x4c, 0x5f, 0x55, 0x4e, 0x46,
	0x55, 0x52, 0x4c, 0x49, 0x4e, 0x47, 0x5f, 0x4d, 0x4f, 0x44, 0x45, 0x10, 0x12, 0x22, 0x04, 0x08,
	0x10, 0x10, 0x10, 0x22, 0x04, 0x08, 0x11, 0x10, 0x11, 0x2a, 0x0d, 0x45, 0x4e, 0x53, 0x5f, 0x55,
	0x53, 0x45, 0x52, 0x4e, 0x41, 0x4d, 0x45, 0x53, 0x2a, 0x19, 0x49, 0x4e, 0x43, 0x4c, 0x55, 0x44,
	0x45, 0x5f, 0x57, 0x41, 0x54, 0x43, 0x48, 0x4f, 0x4e, 0x4c, 0x59, 0x5f, 0x41, 0x43, 0x43, 0x4f,
	0x55, 0x4e, 0x54, 0x42, 0x07, 0x0a, 0x05, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x42, 0x0d, 0x5a, 0x0b,
	0x2e, 0x2f, 0x3b, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x62, 0x06, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x33,
}

var (
	file_sync_settings_proto_rawDescOnce sync.Once
	file_sync_settings_proto_rawDescData = file_sync_settings_proto_rawDesc
)

func file_sync_settings_proto_rawDescGZIP() []byte {
	file_sync_settings_proto_rawDescOnce.Do(func() {
		file_sync_settings_proto_rawDescData = protoimpl.X.CompressGZIP(file_sync_settings_proto_rawDescData)
	})
	return file_sync_settings_proto_rawDescData
}

var file_sync_settings_proto_enumTypes = make([]protoimpl.EnumInfo, 1)
var file_sync_settings_proto_msgTypes = make([]protoimpl.MessageInfo, 1)
var file_sync_settings_proto_goTypes = []interface{}{
	(SyncSetting_Type)(0), // 0: protobuf.SyncSetting.Type
	(*SyncSetting)(nil),   // 1: protobuf.SyncSetting
}
var file_sync_settings_proto_depIdxs = []int32{
	0, // 0: protobuf.SyncSetting.type:type_name -> protobuf.SyncSetting.Type
	1, // [1:1] is the sub-list for method output_type
	1, // [1:1] is the sub-list for method input_type
	1, // [1:1] is the sub-list for extension type_name
	1, // [1:1] is the sub-list for extension extendee
	0, // [0:1] is the sub-list for field type_name
}

func init() { file_sync_settings_proto_init() }
func file_sync_settings_proto_init() {
	if File_sync_settings_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_sync_settings_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*SyncSetting); i {
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
	file_sync_settings_proto_msgTypes[0].OneofWrappers = []interface{}{
		(*SyncSetting_ValueString)(nil),
		(*SyncSetting_ValueBytes)(nil),
		(*SyncSetting_ValueBool)(nil),
		(*SyncSetting_ValueInt64)(nil),
	}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_sync_settings_proto_rawDesc,
			NumEnums:      1,
			NumMessages:   1,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_sync_settings_proto_goTypes,
		DependencyIndexes: file_sync_settings_proto_depIdxs,
		EnumInfos:         file_sync_settings_proto_enumTypes,
		MessageInfos:      file_sync_settings_proto_msgTypes,
	}.Build()
	File_sync_settings_proto = out.File
	file_sync_settings_proto_rawDesc = nil
	file_sync_settings_proto_goTypes = nil
	file_sync_settings_proto_depIdxs = nil
}
