// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.27.1
// 	protoc        v3.19.1
// source: proto/StarWars.proto

package proto

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

type ConsultRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Message string `protobuf:"bytes,1,opt,name=message,proto3" json:"message,omitempty"`
}

func (x *ConsultRequest) Reset() {
	*x = ConsultRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_proto_StarWars_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ConsultRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ConsultRequest) ProtoMessage() {}

func (x *ConsultRequest) ProtoReflect() protoreflect.Message {
	mi := &file_proto_StarWars_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ConsultRequest.ProtoReflect.Descriptor instead.
func (*ConsultRequest) Descriptor() ([]byte, []int) {
	return file_proto_StarWars_proto_rawDescGZIP(), []int{0}
}

func (x *ConsultRequest) GetMessage() string {
	if x != nil {
		return x.Message
	}
	return ""
}

type ConsultReply struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Message string `protobuf:"bytes,1,opt,name=message,proto3" json:"message,omitempty"`
}

func (x *ConsultReply) Reset() {
	*x = ConsultReply{}
	if protoimpl.UnsafeEnabled {
		mi := &file_proto_StarWars_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ConsultReply) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ConsultReply) ProtoMessage() {}

func (x *ConsultReply) ProtoReflect() protoreflect.Message {
	mi := &file_proto_StarWars_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ConsultReply.ProtoReflect.Descriptor instead.
func (*ConsultReply) Descriptor() ([]byte, []int) {
	return file_proto_StarWars_proto_rawDescGZIP(), []int{1}
}

func (x *ConsultReply) GetMessage() string {
	if x != nil {
		return x.Message
	}
	return ""
}

type SendRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Message string `protobuf:"bytes,1,opt,name=message,proto3" json:"message,omitempty"`
}

func (x *SendRequest) Reset() {
	*x = SendRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_proto_StarWars_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *SendRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*SendRequest) ProtoMessage() {}

func (x *SendRequest) ProtoReflect() protoreflect.Message {
	mi := &file_proto_StarWars_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use SendRequest.ProtoReflect.Descriptor instead.
func (*SendRequest) Descriptor() ([]byte, []int) {
	return file_proto_StarWars_proto_rawDescGZIP(), []int{2}
}

func (x *SendRequest) GetMessage() string {
	if x != nil {
		return x.Message
	}
	return ""
}

type SendReply struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Message string `protobuf:"bytes,1,opt,name=message,proto3" json:"message,omitempty"`
}

func (x *SendReply) Reset() {
	*x = SendReply{}
	if protoimpl.UnsafeEnabled {
		mi := &file_proto_StarWars_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *SendReply) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*SendReply) ProtoMessage() {}

func (x *SendReply) ProtoReflect() protoreflect.Message {
	mi := &file_proto_StarWars_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use SendReply.ProtoReflect.Descriptor instead.
func (*SendReply) Descriptor() ([]byte, []int) {
	return file_proto_StarWars_proto_rawDescGZIP(), []int{3}
}

func (x *SendReply) GetMessage() string {
	if x != nil {
		return x.Message
	}
	return ""
}

var File_proto_StarWars_proto protoreflect.FileDescriptor

var file_proto_StarWars_proto_rawDesc = []byte{
	0x0a, 0x14, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2f, 0x53, 0x74, 0x61, 0x72, 0x57, 0x61, 0x72, 0x73,
	0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x04, 0x67, 0x72, 0x70, 0x63, 0x22, 0x2a, 0x0a, 0x0e,
	0x43, 0x6f, 0x6e, 0x73, 0x75, 0x6c, 0x74, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x18,
	0x0a, 0x07, 0x6d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52,
	0x07, 0x6d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x22, 0x28, 0x0a, 0x0c, 0x43, 0x6f, 0x6e, 0x73,
	0x75, 0x6c, 0x74, 0x52, 0x65, 0x70, 0x6c, 0x79, 0x12, 0x18, 0x0a, 0x07, 0x6d, 0x65, 0x73, 0x73,
	0x61, 0x67, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x07, 0x6d, 0x65, 0x73, 0x73, 0x61,
	0x67, 0x65, 0x22, 0x27, 0x0a, 0x0b, 0x53, 0x65, 0x6e, 0x64, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73,
	0x74, 0x12, 0x18, 0x0a, 0x07, 0x6d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x18, 0x01, 0x20, 0x01,
	0x28, 0x09, 0x52, 0x07, 0x6d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x22, 0x25, 0x0a, 0x09, 0x53,
	0x65, 0x6e, 0x64, 0x52, 0x65, 0x70, 0x6c, 0x79, 0x12, 0x18, 0x0a, 0x07, 0x6d, 0x65, 0x73, 0x73,
	0x61, 0x67, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x07, 0x6d, 0x65, 0x73, 0x73, 0x61,
	0x67, 0x65, 0x32, 0x87, 0x01, 0x0a, 0x0f, 0x53, 0x74, 0x61, 0x72, 0x57, 0x61, 0x72, 0x73, 0x53,
	0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x12, 0x3b, 0x0a, 0x0d, 0x43, 0x6f, 0x6e, 0x73, 0x75, 0x6c,
	0x74, 0x50, 0x6c, 0x61, 0x6e, 0x65, 0x74, 0x12, 0x14, 0x2e, 0x67, 0x72, 0x70, 0x63, 0x2e, 0x43,
	0x6f, 0x6e, 0x73, 0x75, 0x6c, 0x74, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x12, 0x2e,
	0x67, 0x72, 0x70, 0x63, 0x2e, 0x43, 0x6f, 0x6e, 0x73, 0x75, 0x6c, 0x74, 0x52, 0x65, 0x70, 0x6c,
	0x79, 0x22, 0x00, 0x12, 0x37, 0x0a, 0x0f, 0x53, 0x65, 0x6e, 0x64, 0x49, 0x6e, 0x66, 0x6f, 0x72,
	0x6d, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x12, 0x11, 0x2e, 0x67, 0x72, 0x70, 0x63, 0x2e, 0x53, 0x65,
	0x6e, 0x64, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x0f, 0x2e, 0x67, 0x72, 0x70, 0x63,
	0x2e, 0x53, 0x65, 0x6e, 0x64, 0x52, 0x65, 0x70, 0x6c, 0x79, 0x22, 0x00, 0x42, 0x26, 0x5a, 0x24,
	0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x66, 0x61, 0x62, 0x69, 0x75,
	0x73, 0x69, 0x6e, 0x66, 0x6f, 0x2f, 0x53, 0x74, 0x61, 0x72, 0x57, 0x61, 0x72, 0x73, 0x2f, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_proto_StarWars_proto_rawDescOnce sync.Once
	file_proto_StarWars_proto_rawDescData = file_proto_StarWars_proto_rawDesc
)

func file_proto_StarWars_proto_rawDescGZIP() []byte {
	file_proto_StarWars_proto_rawDescOnce.Do(func() {
		file_proto_StarWars_proto_rawDescData = protoimpl.X.CompressGZIP(file_proto_StarWars_proto_rawDescData)
	})
	return file_proto_StarWars_proto_rawDescData
}

var file_proto_StarWars_proto_msgTypes = make([]protoimpl.MessageInfo, 4)
var file_proto_StarWars_proto_goTypes = []interface{}{
	(*ConsultRequest)(nil), // 0: grpc.ConsultRequest
	(*ConsultReply)(nil),   // 1: grpc.ConsultReply
	(*SendRequest)(nil),    // 2: grpc.SendRequest
	(*SendReply)(nil),      // 3: grpc.SendReply
}
var file_proto_StarWars_proto_depIdxs = []int32{
	0, // 0: grpc.StarWarsService.ConsultPlanet:input_type -> grpc.ConsultRequest
	2, // 1: grpc.StarWarsService.SendInformation:input_type -> grpc.SendRequest
	1, // 2: grpc.StarWarsService.ConsultPlanet:output_type -> grpc.ConsultReply
	3, // 3: grpc.StarWarsService.SendInformation:output_type -> grpc.SendReply
	2, // [2:4] is the sub-list for method output_type
	0, // [0:2] is the sub-list for method input_type
	0, // [0:0] is the sub-list for extension type_name
	0, // [0:0] is the sub-list for extension extendee
	0, // [0:0] is the sub-list for field type_name
}

func init() { file_proto_StarWars_proto_init() }
func file_proto_StarWars_proto_init() {
	if File_proto_StarWars_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_proto_StarWars_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ConsultRequest); i {
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
		file_proto_StarWars_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ConsultReply); i {
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
		file_proto_StarWars_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*SendRequest); i {
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
		file_proto_StarWars_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*SendReply); i {
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
			RawDescriptor: file_proto_StarWars_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   4,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_proto_StarWars_proto_goTypes,
		DependencyIndexes: file_proto_StarWars_proto_depIdxs,
		MessageInfos:      file_proto_StarWars_proto_msgTypes,
	}.Build()
	File_proto_StarWars_proto = out.File
	file_proto_StarWars_proto_rawDesc = nil
	file_proto_StarWars_proto_goTypes = nil
	file_proto_StarWars_proto_depIdxs = nil
}
