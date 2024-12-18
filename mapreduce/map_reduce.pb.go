// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.35.2
// 	protoc        v5.29.1
// source: mapreduce/map_reduce.proto

package mapreduce

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

type DataSet struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Values []int32 `protobuf:"varint,1,rep,packed,name=values,proto3" json:"values,omitempty"`
}

func (x *DataSet) Reset() {
	*x = DataSet{}
	mi := &file_mapreduce_map_reduce_proto_msgTypes[0]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *DataSet) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*DataSet) ProtoMessage() {}

func (x *DataSet) ProtoReflect() protoreflect.Message {
	mi := &file_mapreduce_map_reduce_proto_msgTypes[0]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use DataSet.ProtoReflect.Descriptor instead.
func (*DataSet) Descriptor() ([]byte, []int) {
	return file_mapreduce_map_reduce_proto_rawDescGZIP(), []int{0}
}

func (x *DataSet) GetValues() []int32 {
	if x != nil {
		return x.Values
	}
	return nil
}

var File_mapreduce_map_reduce_proto protoreflect.FileDescriptor

var file_mapreduce_map_reduce_proto_rawDesc = []byte{
	0x0a, 0x1a, 0x6d, 0x61, 0x70, 0x72, 0x65, 0x64, 0x75, 0x63, 0x65, 0x2f, 0x6d, 0x61, 0x70, 0x5f,
	0x72, 0x65, 0x64, 0x75, 0x63, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x09, 0x6d, 0x61,
	0x70, 0x72, 0x65, 0x64, 0x75, 0x63, 0x65, 0x22, 0x21, 0x0a, 0x07, 0x44, 0x61, 0x74, 0x61, 0x53,
	0x65, 0x74, 0x12, 0x16, 0x0a, 0x06, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x73, 0x18, 0x01, 0x20, 0x03,
	0x28, 0x05, 0x52, 0x06, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x73, 0x32, 0x3f, 0x0a, 0x09, 0x4d, 0x61,
	0x70, 0x52, 0x65, 0x64, 0x75, 0x63, 0x65, 0x12, 0x32, 0x0a, 0x08, 0x53, 0x6f, 0x72, 0x74, 0x44,
	0x61, 0x74, 0x61, 0x12, 0x12, 0x2e, 0x6d, 0x61, 0x70, 0x72, 0x65, 0x64, 0x75, 0x63, 0x65, 0x2e,
	0x44, 0x61, 0x74, 0x61, 0x53, 0x65, 0x74, 0x1a, 0x12, 0x2e, 0x6d, 0x61, 0x70, 0x72, 0x65, 0x64,
	0x75, 0x63, 0x65, 0x2e, 0x44, 0x61, 0x74, 0x61, 0x53, 0x65, 0x74, 0x42, 0x15, 0x5a, 0x13, 0x4d,
	0x61, 0x70, 0x52, 0x65, 0x64, 0x75, 0x63, 0x65, 0x2f, 0x6d, 0x61, 0x70, 0x72, 0x65, 0x64, 0x75,
	0x63, 0x65, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_mapreduce_map_reduce_proto_rawDescOnce sync.Once
	file_mapreduce_map_reduce_proto_rawDescData = file_mapreduce_map_reduce_proto_rawDesc
)

func file_mapreduce_map_reduce_proto_rawDescGZIP() []byte {
	file_mapreduce_map_reduce_proto_rawDescOnce.Do(func() {
		file_mapreduce_map_reduce_proto_rawDescData = protoimpl.X.CompressGZIP(file_mapreduce_map_reduce_proto_rawDescData)
	})
	return file_mapreduce_map_reduce_proto_rawDescData
}

var file_mapreduce_map_reduce_proto_msgTypes = make([]protoimpl.MessageInfo, 1)
var file_mapreduce_map_reduce_proto_goTypes = []any{
	(*DataSet)(nil), // 0: mapreduce.DataSet
}
var file_mapreduce_map_reduce_proto_depIdxs = []int32{
	0, // 0: mapreduce.MapReduce.SortData:input_type -> mapreduce.DataSet
	0, // 1: mapreduce.MapReduce.SortData:output_type -> mapreduce.DataSet
	1, // [1:2] is the sub-list for method output_type
	0, // [0:1] is the sub-list for method input_type
	0, // [0:0] is the sub-list for extension type_name
	0, // [0:0] is the sub-list for extension extendee
	0, // [0:0] is the sub-list for field type_name
}

func init() { file_mapreduce_map_reduce_proto_init() }
func file_mapreduce_map_reduce_proto_init() {
	if File_mapreduce_map_reduce_proto != nil {
		return
	}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_mapreduce_map_reduce_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   1,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_mapreduce_map_reduce_proto_goTypes,
		DependencyIndexes: file_mapreduce_map_reduce_proto_depIdxs,
		MessageInfos:      file_mapreduce_map_reduce_proto_msgTypes,
	}.Build()
	File_mapreduce_map_reduce_proto = out.File
	file_mapreduce_map_reduce_proto_rawDesc = nil
	file_mapreduce_map_reduce_proto_goTypes = nil
	file_mapreduce_map_reduce_proto_depIdxs = nil
}
