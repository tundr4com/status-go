// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.26.0
// 	protoc        v3.21.12
// source: waku_lightpush.proto

package pb

import (
	pb "github.com/waku-org/go-waku/waku/v2/protocol/pb"
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

type PushRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	PubsubTopic string          `protobuf:"bytes,1,opt,name=pubsub_topic,json=pubsubTopic,proto3" json:"pubsub_topic,omitempty"`
	Message     *pb.WakuMessage `protobuf:"bytes,2,opt,name=message,proto3" json:"message,omitempty"`
}

func (x *PushRequest) Reset() {
	*x = PushRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_waku_lightpush_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *PushRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*PushRequest) ProtoMessage() {}

func (x *PushRequest) ProtoReflect() protoreflect.Message {
	mi := &file_waku_lightpush_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use PushRequest.ProtoReflect.Descriptor instead.
func (*PushRequest) Descriptor() ([]byte, []int) {
	return file_waku_lightpush_proto_rawDescGZIP(), []int{0}
}

func (x *PushRequest) GetPubsubTopic() string {
	if x != nil {
		return x.PubsubTopic
	}
	return ""
}

func (x *PushRequest) GetMessage() *pb.WakuMessage {
	if x != nil {
		return x.Message
	}
	return nil
}

type PushResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	IsSuccess bool `protobuf:"varint,1,opt,name=is_success,json=isSuccess,proto3" json:"is_success,omitempty"`
	// Error messages, etc
	Info string `protobuf:"bytes,2,opt,name=info,proto3" json:"info,omitempty"`
}

func (x *PushResponse) Reset() {
	*x = PushResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_waku_lightpush_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *PushResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*PushResponse) ProtoMessage() {}

func (x *PushResponse) ProtoReflect() protoreflect.Message {
	mi := &file_waku_lightpush_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use PushResponse.ProtoReflect.Descriptor instead.
func (*PushResponse) Descriptor() ([]byte, []int) {
	return file_waku_lightpush_proto_rawDescGZIP(), []int{1}
}

func (x *PushResponse) GetIsSuccess() bool {
	if x != nil {
		return x.IsSuccess
	}
	return false
}

func (x *PushResponse) GetInfo() string {
	if x != nil {
		return x.Info
	}
	return ""
}

type PushRPC struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	RequestId string        `protobuf:"bytes,1,opt,name=request_id,json=requestId,proto3" json:"request_id,omitempty"`
	Query     *PushRequest  `protobuf:"bytes,2,opt,name=query,proto3" json:"query,omitempty"`
	Response  *PushResponse `protobuf:"bytes,3,opt,name=response,proto3" json:"response,omitempty"`
}

func (x *PushRPC) Reset() {
	*x = PushRPC{}
	if protoimpl.UnsafeEnabled {
		mi := &file_waku_lightpush_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *PushRPC) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*PushRPC) ProtoMessage() {}

func (x *PushRPC) ProtoReflect() protoreflect.Message {
	mi := &file_waku_lightpush_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use PushRPC.ProtoReflect.Descriptor instead.
func (*PushRPC) Descriptor() ([]byte, []int) {
	return file_waku_lightpush_proto_rawDescGZIP(), []int{2}
}

func (x *PushRPC) GetRequestId() string {
	if x != nil {
		return x.RequestId
	}
	return ""
}

func (x *PushRPC) GetQuery() *PushRequest {
	if x != nil {
		return x.Query
	}
	return nil
}

func (x *PushRPC) GetResponse() *PushResponse {
	if x != nil {
		return x.Response
	}
	return nil
}

var File_waku_lightpush_proto protoreflect.FileDescriptor

var file_waku_lightpush_proto_rawDesc = []byte{
	0x0a, 0x14, 0x77, 0x61, 0x6b, 0x75, 0x5f, 0x6c, 0x69, 0x67, 0x68, 0x74, 0x70, 0x75, 0x73, 0x68,
	0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x02, 0x70, 0x62, 0x1a, 0x12, 0x77, 0x61, 0x6b, 0x75,
	0x5f, 0x6d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0x5b,
	0x0a, 0x0b, 0x50, 0x75, 0x73, 0x68, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x21, 0x0a,
	0x0c, 0x70, 0x75, 0x62, 0x73, 0x75, 0x62, 0x5f, 0x74, 0x6f, 0x70, 0x69, 0x63, 0x18, 0x01, 0x20,
	0x01, 0x28, 0x09, 0x52, 0x0b, 0x70, 0x75, 0x62, 0x73, 0x75, 0x62, 0x54, 0x6f, 0x70, 0x69, 0x63,
	0x12, 0x29, 0x0a, 0x07, 0x6d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28,
	0x0b, 0x32, 0x0f, 0x2e, 0x70, 0x62, 0x2e, 0x57, 0x61, 0x6b, 0x75, 0x4d, 0x65, 0x73, 0x73, 0x61,
	0x67, 0x65, 0x52, 0x07, 0x6d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x22, 0x41, 0x0a, 0x0c, 0x50,
	0x75, 0x73, 0x68, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x1d, 0x0a, 0x0a, 0x69,
	0x73, 0x5f, 0x73, 0x75, 0x63, 0x63, 0x65, 0x73, 0x73, 0x18, 0x01, 0x20, 0x01, 0x28, 0x08, 0x52,
	0x09, 0x69, 0x73, 0x53, 0x75, 0x63, 0x63, 0x65, 0x73, 0x73, 0x12, 0x12, 0x0a, 0x04, 0x69, 0x6e,
	0x66, 0x6f, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x69, 0x6e, 0x66, 0x6f, 0x22, 0x7d,
	0x0a, 0x07, 0x50, 0x75, 0x73, 0x68, 0x52, 0x50, 0x43, 0x12, 0x1d, 0x0a, 0x0a, 0x72, 0x65, 0x71,
	0x75, 0x65, 0x73, 0x74, 0x5f, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x09, 0x72,
	0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x49, 0x64, 0x12, 0x25, 0x0a, 0x05, 0x71, 0x75, 0x65, 0x72,
	0x79, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x0f, 0x2e, 0x70, 0x62, 0x2e, 0x50, 0x75, 0x73,
	0x68, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x52, 0x05, 0x71, 0x75, 0x65, 0x72, 0x79, 0x12,
	0x2c, 0x0a, 0x08, 0x72, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x18, 0x03, 0x20, 0x01, 0x28,
	0x0b, 0x32, 0x10, 0x2e, 0x70, 0x62, 0x2e, 0x50, 0x75, 0x73, 0x68, 0x52, 0x65, 0x73, 0x70, 0x6f,
	0x6e, 0x73, 0x65, 0x52, 0x08, 0x72, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x62, 0x06, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_waku_lightpush_proto_rawDescOnce sync.Once
	file_waku_lightpush_proto_rawDescData = file_waku_lightpush_proto_rawDesc
)

func file_waku_lightpush_proto_rawDescGZIP() []byte {
	file_waku_lightpush_proto_rawDescOnce.Do(func() {
		file_waku_lightpush_proto_rawDescData = protoimpl.X.CompressGZIP(file_waku_lightpush_proto_rawDescData)
	})
	return file_waku_lightpush_proto_rawDescData
}

var file_waku_lightpush_proto_msgTypes = make([]protoimpl.MessageInfo, 3)
var file_waku_lightpush_proto_goTypes = []interface{}{
	(*PushRequest)(nil),    // 0: pb.PushRequest
	(*PushResponse)(nil),   // 1: pb.PushResponse
	(*PushRPC)(nil),        // 2: pb.PushRPC
	(*pb.WakuMessage)(nil), // 3: pb.WakuMessage
}
var file_waku_lightpush_proto_depIdxs = []int32{
	3, // 0: pb.PushRequest.message:type_name -> pb.WakuMessage
	0, // 1: pb.PushRPC.query:type_name -> pb.PushRequest
	1, // 2: pb.PushRPC.response:type_name -> pb.PushResponse
	3, // [3:3] is the sub-list for method output_type
	3, // [3:3] is the sub-list for method input_type
	3, // [3:3] is the sub-list for extension type_name
	3, // [3:3] is the sub-list for extension extendee
	0, // [0:3] is the sub-list for field type_name
}

func init() { file_waku_lightpush_proto_init() }
func file_waku_lightpush_proto_init() {
	if File_waku_lightpush_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_waku_lightpush_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*PushRequest); i {
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
		file_waku_lightpush_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*PushResponse); i {
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
		file_waku_lightpush_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*PushRPC); i {
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
			RawDescriptor: file_waku_lightpush_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   3,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_waku_lightpush_proto_goTypes,
		DependencyIndexes: file_waku_lightpush_proto_depIdxs,
		MessageInfos:      file_waku_lightpush_proto_msgTypes,
	}.Build()
	File_waku_lightpush_proto = out.File
	file_waku_lightpush_proto_rawDesc = nil
	file_waku_lightpush_proto_goTypes = nil
	file_waku_lightpush_proto_depIdxs = nil
}
