// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.27.1
// 	protoc        v3.12.4
// source: executor.proto

package __

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

type ExecutionRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Content []byte `protobuf:"bytes,1,opt,name=content,proto3" json:"content,omitempty"`
}

func (x *ExecutionRequest) Reset() {
	*x = ExecutionRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_executor_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ExecutionRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ExecutionRequest) ProtoMessage() {}

func (x *ExecutionRequest) ProtoReflect() protoreflect.Message {
	mi := &file_executor_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ExecutionRequest.ProtoReflect.Descriptor instead.
func (*ExecutionRequest) Descriptor() ([]byte, []int) {
	return file_executor_proto_rawDescGZIP(), []int{0}
}

func (x *ExecutionRequest) GetContent() []byte {
	if x != nil {
		return x.Content
	}
	return nil
}

type ExecutionResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Output   string `protobuf:"bytes,1,opt,name=output,proto3" json:"output,omitempty"`
	ExecTime string `protobuf:"bytes,2,opt,name=execTime,proto3" json:"execTime,omitempty"`
	RamUsed  string `protobuf:"bytes,3,opt,name=ramUsed,proto3" json:"ramUsed,omitempty"`
	CpuStats string `protobuf:"bytes,4,opt,name=cpuStats,proto3" json:"cpuStats,omitempty"`
}

func (x *ExecutionResponse) Reset() {
	*x = ExecutionResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_executor_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ExecutionResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ExecutionResponse) ProtoMessage() {}

func (x *ExecutionResponse) ProtoReflect() protoreflect.Message {
	mi := &file_executor_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ExecutionResponse.ProtoReflect.Descriptor instead.
func (*ExecutionResponse) Descriptor() ([]byte, []int) {
	return file_executor_proto_rawDescGZIP(), []int{1}
}

func (x *ExecutionResponse) GetOutput() string {
	if x != nil {
		return x.Output
	}
	return ""
}

func (x *ExecutionResponse) GetExecTime() string {
	if x != nil {
		return x.ExecTime
	}
	return ""
}

func (x *ExecutionResponse) GetRamUsed() string {
	if x != nil {
		return x.RamUsed
	}
	return ""
}

func (x *ExecutionResponse) GetCpuStats() string {
	if x != nil {
		return x.CpuStats
	}
	return ""
}

var File_executor_proto protoreflect.FileDescriptor

var file_executor_proto_rawDesc = []byte{
	0x0a, 0x0e, 0x65, 0x78, 0x65, 0x63, 0x75, 0x74, 0x6f, 0x72, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x12, 0x05, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0x2c, 0x0a, 0x10, 0x45, 0x78, 0x65, 0x63, 0x75,
	0x74, 0x69, 0x6f, 0x6e, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x18, 0x0a, 0x07, 0x63,
	0x6f, 0x6e, 0x74, 0x65, 0x6e, 0x74, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0c, 0x52, 0x07, 0x63, 0x6f,
	0x6e, 0x74, 0x65, 0x6e, 0x74, 0x22, 0x7d, 0x0a, 0x11, 0x45, 0x78, 0x65, 0x63, 0x75, 0x74, 0x69,
	0x6f, 0x6e, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x16, 0x0a, 0x06, 0x6f, 0x75,
	0x74, 0x70, 0x75, 0x74, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x06, 0x6f, 0x75, 0x74, 0x70,
	0x75, 0x74, 0x12, 0x1a, 0x0a, 0x08, 0x65, 0x78, 0x65, 0x63, 0x54, 0x69, 0x6d, 0x65, 0x18, 0x02,
	0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x65, 0x78, 0x65, 0x63, 0x54, 0x69, 0x6d, 0x65, 0x12, 0x18,
	0x0a, 0x07, 0x72, 0x61, 0x6d, 0x55, 0x73, 0x65, 0x64, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52,
	0x07, 0x72, 0x61, 0x6d, 0x55, 0x73, 0x65, 0x64, 0x12, 0x1a, 0x0a, 0x08, 0x63, 0x70, 0x75, 0x53,
	0x74, 0x61, 0x74, 0x73, 0x18, 0x04, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x63, 0x70, 0x75, 0x53,
	0x74, 0x61, 0x74, 0x73, 0x32, 0x59, 0x0a, 0x0f, 0x45, 0x78, 0x65, 0x63, 0x75, 0x74, 0x6f, 0x72,
	0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x12, 0x46, 0x0a, 0x0f, 0x48, 0x61, 0x6e, 0x64, 0x6c,
	0x65, 0x45, 0x78, 0x65, 0x63, 0x75, 0x74, 0x69, 0x6f, 0x6e, 0x12, 0x17, 0x2e, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x2e, 0x45, 0x78, 0x65, 0x63, 0x75, 0x74, 0x69, 0x6f, 0x6e, 0x52, 0x65, 0x71, 0x75,
	0x65, 0x73, 0x74, 0x1a, 0x18, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2e, 0x45, 0x78, 0x65, 0x63,
	0x75, 0x74, 0x69, 0x6f, 0x6e, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x28, 0x01, 0x42,
	0x03, 0x5a, 0x01, 0x2e, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_executor_proto_rawDescOnce sync.Once
	file_executor_proto_rawDescData = file_executor_proto_rawDesc
)

func file_executor_proto_rawDescGZIP() []byte {
	file_executor_proto_rawDescOnce.Do(func() {
		file_executor_proto_rawDescData = protoimpl.X.CompressGZIP(file_executor_proto_rawDescData)
	})
	return file_executor_proto_rawDescData
}

var file_executor_proto_msgTypes = make([]protoimpl.MessageInfo, 2)
var file_executor_proto_goTypes = []interface{}{
	(*ExecutionRequest)(nil),  // 0: proto.ExecutionRequest
	(*ExecutionResponse)(nil), // 1: proto.ExecutionResponse
}
var file_executor_proto_depIdxs = []int32{
	0, // 0: proto.ExecutorService.HandleExecution:input_type -> proto.ExecutionRequest
	1, // 1: proto.ExecutorService.HandleExecution:output_type -> proto.ExecutionResponse
	1, // [1:2] is the sub-list for method output_type
	0, // [0:1] is the sub-list for method input_type
	0, // [0:0] is the sub-list for extension type_name
	0, // [0:0] is the sub-list for extension extendee
	0, // [0:0] is the sub-list for field type_name
}

func init() { file_executor_proto_init() }
func file_executor_proto_init() {
	if File_executor_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_executor_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ExecutionRequest); i {
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
		file_executor_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ExecutionResponse); i {
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
			RawDescriptor: file_executor_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   2,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_executor_proto_goTypes,
		DependencyIndexes: file_executor_proto_depIdxs,
		MessageInfos:      file_executor_proto_msgTypes,
	}.Build()
	File_executor_proto = out.File
	file_executor_proto_rawDesc = nil
	file_executor_proto_goTypes = nil
	file_executor_proto_depIdxs = nil
}
