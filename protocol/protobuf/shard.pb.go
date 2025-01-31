// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.31.0
// 	protoc        v4.24.4
// source: shard.proto

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

type Shard struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Cluster int32 `protobuf:"varint,1,opt,name=cluster,proto3" json:"cluster,omitempty"`
	Index   int32 `protobuf:"varint,2,opt,name=index,proto3" json:"index,omitempty"`
}

func (x *Shard) Reset() {
	*x = Shard{}
	if protoimpl.UnsafeEnabled {
		mi := &file_shard_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Shard) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Shard) ProtoMessage() {}

func (x *Shard) ProtoReflect() protoreflect.Message {
	mi := &file_shard_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Shard.ProtoReflect.Descriptor instead.
func (*Shard) Descriptor() ([]byte, []int) {
	return file_shard_proto_rawDescGZIP(), []int{0}
}

func (x *Shard) GetCluster() int32 {
	if x != nil {
		return x.Cluster
	}
	return 0
}

func (x *Shard) GetIndex() int32 {
	if x != nil {
		return x.Index
	}
	return 0
}

var File_shard_proto protoreflect.FileDescriptor

var file_shard_proto_rawDesc = []byte{
	0x0a, 0x0b, 0x73, 0x68, 0x61, 0x72, 0x64, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x08, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x22, 0x37, 0x0a, 0x05, 0x53, 0x68, 0x61, 0x72, 0x64,
	0x12, 0x18, 0x0a, 0x07, 0x63, 0x6c, 0x75, 0x73, 0x74, 0x65, 0x72, 0x18, 0x01, 0x20, 0x01, 0x28,
	0x05, 0x52, 0x07, 0x63, 0x6c, 0x75, 0x73, 0x74, 0x65, 0x72, 0x12, 0x14, 0x0a, 0x05, 0x69, 0x6e,
	0x64, 0x65, 0x78, 0x18, 0x02, 0x20, 0x01, 0x28, 0x05, 0x52, 0x05, 0x69, 0x6e, 0x64, 0x65, 0x78,
	0x42, 0x0d, 0x5a, 0x0b, 0x2e, 0x2f, 0x3b, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x62,
	0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_shard_proto_rawDescOnce sync.Once
	file_shard_proto_rawDescData = file_shard_proto_rawDesc
)

func file_shard_proto_rawDescGZIP() []byte {
	file_shard_proto_rawDescOnce.Do(func() {
		file_shard_proto_rawDescData = protoimpl.X.CompressGZIP(file_shard_proto_rawDescData)
	})
	return file_shard_proto_rawDescData
}

var file_shard_proto_msgTypes = make([]protoimpl.MessageInfo, 1)
var file_shard_proto_goTypes = []interface{}{
	(*Shard)(nil), // 0: protobuf.Shard
}
var file_shard_proto_depIdxs = []int32{
	0, // [0:0] is the sub-list for method output_type
	0, // [0:0] is the sub-list for method input_type
	0, // [0:0] is the sub-list for extension type_name
	0, // [0:0] is the sub-list for extension extendee
	0, // [0:0] is the sub-list for field type_name
}

func init() { file_shard_proto_init() }
func file_shard_proto_init() {
	if File_shard_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_shard_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Shard); i {
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
			RawDescriptor: file_shard_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   1,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_shard_proto_goTypes,
		DependencyIndexes: file_shard_proto_depIdxs,
		MessageInfos:      file_shard_proto_msgTypes,
	}.Build()
	File_shard_proto = out.File
	file_shard_proto_rawDesc = nil
	file_shard_proto_goTypes = nil
	file_shard_proto_depIdxs = nil
}
