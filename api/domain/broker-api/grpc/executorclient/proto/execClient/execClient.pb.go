// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.27.1
// 	protoc        v3.12.4
// source: execClient.proto

package execClient

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
	Uid     string `protobuf:"bytes,2,opt,name=uid,proto3" json:"uid,omitempty"`
	Qid     string `protobuf:"bytes,3,opt,name=qid,proto3" json:"qid,omitempty"`
	Lang    string `protobuf:"bytes,4,opt,name=lang,proto3" json:"lang,omitempty"`
	FileExt string `protobuf:"bytes,5,opt,name=fileExt,proto3" json:"fileExt,omitempty"`
}

func (x *ExecutionRequest) Reset() {
	*x = ExecutionRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_execClient_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ExecutionRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ExecutionRequest) ProtoMessage() {}

func (x *ExecutionRequest) ProtoReflect() protoreflect.Message {
	mi := &file_execClient_proto_msgTypes[0]
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
	return file_execClient_proto_rawDescGZIP(), []int{0}
}

func (x *ExecutionRequest) GetContent() []byte {
	if x != nil {
		return x.Content
	}
	return nil
}

func (x *ExecutionRequest) GetUid() string {
	if x != nil {
		return x.Uid
	}
	return ""
}

func (x *ExecutionRequest) GetQid() string {
	if x != nil {
		return x.Qid
	}
	return ""
}

func (x *ExecutionRequest) GetLang() string {
	if x != nil {
		return x.Lang
	}
	return ""
}

func (x *ExecutionRequest) GetFileExt() string {
	if x != nil {
		return x.FileExt
	}
	return ""
}

type ExecutionResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Output         string `protobuf:"bytes,1,opt,name=output,proto3" json:"output,omitempty"`
	ExecTime       string `protobuf:"bytes,2,opt,name=execTime,proto3" json:"execTime,omitempty"`
	RamUsed        string `protobuf:"bytes,3,opt,name=ramUsed,proto3" json:"ramUsed,omitempty"`
	CpuStats       string `protobuf:"bytes,4,opt,name=cpuStats,proto3" json:"cpuStats,omitempty"`
	TotalRAM       string `protobuf:"bytes,5,opt,name=totalRAM,proto3" json:"totalRAM,omitempty"`
	PercetRAMUsage string `protobuf:"bytes,6,opt,name=percetRAMUsage,proto3" json:"percetRAMUsage,omitempty"`
	ContainerID    string `protobuf:"bytes,7,opt,name=containerID,proto3" json:"containerID,omitempty"`
	IsCorrect      bool   `protobuf:"varint,8,opt,name=isCorrect,proto3" json:"isCorrect,omitempty"`
	IsRunOnly      bool   `protobuf:"varint,9,opt,name=isRunOnly,proto3" json:"isRunOnly,omitempty"`
}

func (x *ExecutionResponse) Reset() {
	*x = ExecutionResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_execClient_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ExecutionResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ExecutionResponse) ProtoMessage() {}

func (x *ExecutionResponse) ProtoReflect() protoreflect.Message {
	mi := &file_execClient_proto_msgTypes[1]
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
	return file_execClient_proto_rawDescGZIP(), []int{1}
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

func (x *ExecutionResponse) GetTotalRAM() string {
	if x != nil {
		return x.TotalRAM
	}
	return ""
}

func (x *ExecutionResponse) GetPercetRAMUsage() string {
	if x != nil {
		return x.PercetRAMUsage
	}
	return ""
}

func (x *ExecutionResponse) GetContainerID() string {
	if x != nil {
		return x.ContainerID
	}
	return ""
}

func (x *ExecutionResponse) GetIsCorrect() bool {
	if x != nil {
		return x.IsCorrect
	}
	return false
}

func (x *ExecutionResponse) GetIsRunOnly() bool {
	if x != nil {
		return x.IsRunOnly
	}
	return false
}

var File_execClient_proto protoreflect.FileDescriptor

var file_execClient_proto_rawDesc = []byte{
	0x0a, 0x10, 0x65, 0x78, 0x65, 0x63, 0x43, 0x6c, 0x69, 0x65, 0x6e, 0x74, 0x2e, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x12, 0x05, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0x7e, 0x0a, 0x10, 0x45, 0x78, 0x65,
	0x63, 0x75, 0x74, 0x69, 0x6f, 0x6e, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x18, 0x0a,
	0x07, 0x63, 0x6f, 0x6e, 0x74, 0x65, 0x6e, 0x74, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0c, 0x52, 0x07,
	0x63, 0x6f, 0x6e, 0x74, 0x65, 0x6e, 0x74, 0x12, 0x10, 0x0a, 0x03, 0x75, 0x69, 0x64, 0x18, 0x02,
	0x20, 0x01, 0x28, 0x09, 0x52, 0x03, 0x75, 0x69, 0x64, 0x12, 0x10, 0x0a, 0x03, 0x71, 0x69, 0x64,
	0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x03, 0x71, 0x69, 0x64, 0x12, 0x12, 0x0a, 0x04, 0x6c,
	0x61, 0x6e, 0x67, 0x18, 0x04, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x6c, 0x61, 0x6e, 0x67, 0x12,
	0x18, 0x0a, 0x07, 0x66, 0x69, 0x6c, 0x65, 0x45, 0x78, 0x74, 0x18, 0x05, 0x20, 0x01, 0x28, 0x09,
	0x52, 0x07, 0x66, 0x69, 0x6c, 0x65, 0x45, 0x78, 0x74, 0x22, 0x9f, 0x02, 0x0a, 0x11, 0x45, 0x78,
	0x65, 0x63, 0x75, 0x74, 0x69, 0x6f, 0x6e, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12,
	0x16, 0x0a, 0x06, 0x6f, 0x75, 0x74, 0x70, 0x75, 0x74, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52,
	0x06, 0x6f, 0x75, 0x74, 0x70, 0x75, 0x74, 0x12, 0x1a, 0x0a, 0x08, 0x65, 0x78, 0x65, 0x63, 0x54,
	0x69, 0x6d, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x65, 0x78, 0x65, 0x63, 0x54,
	0x69, 0x6d, 0x65, 0x12, 0x18, 0x0a, 0x07, 0x72, 0x61, 0x6d, 0x55, 0x73, 0x65, 0x64, 0x18, 0x03,
	0x20, 0x01, 0x28, 0x09, 0x52, 0x07, 0x72, 0x61, 0x6d, 0x55, 0x73, 0x65, 0x64, 0x12, 0x1a, 0x0a,
	0x08, 0x63, 0x70, 0x75, 0x53, 0x74, 0x61, 0x74, 0x73, 0x18, 0x04, 0x20, 0x01, 0x28, 0x09, 0x52,
	0x08, 0x63, 0x70, 0x75, 0x53, 0x74, 0x61, 0x74, 0x73, 0x12, 0x1a, 0x0a, 0x08, 0x74, 0x6f, 0x74,
	0x61, 0x6c, 0x52, 0x41, 0x4d, 0x18, 0x05, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x74, 0x6f, 0x74,
	0x61, 0x6c, 0x52, 0x41, 0x4d, 0x12, 0x26, 0x0a, 0x0e, 0x70, 0x65, 0x72, 0x63, 0x65, 0x74, 0x52,
	0x41, 0x4d, 0x55, 0x73, 0x61, 0x67, 0x65, 0x18, 0x06, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0e, 0x70,
	0x65, 0x72, 0x63, 0x65, 0x74, 0x52, 0x41, 0x4d, 0x55, 0x73, 0x61, 0x67, 0x65, 0x12, 0x20, 0x0a,
	0x0b, 0x63, 0x6f, 0x6e, 0x74, 0x61, 0x69, 0x6e, 0x65, 0x72, 0x49, 0x44, 0x18, 0x07, 0x20, 0x01,
	0x28, 0x09, 0x52, 0x0b, 0x63, 0x6f, 0x6e, 0x74, 0x61, 0x69, 0x6e, 0x65, 0x72, 0x49, 0x44, 0x12,
	0x1c, 0x0a, 0x09, 0x69, 0x73, 0x43, 0x6f, 0x72, 0x72, 0x65, 0x63, 0x74, 0x18, 0x08, 0x20, 0x01,
	0x28, 0x08, 0x52, 0x09, 0x69, 0x73, 0x43, 0x6f, 0x72, 0x72, 0x65, 0x63, 0x74, 0x12, 0x1c, 0x0a,
	0x09, 0x69, 0x73, 0x52, 0x75, 0x6e, 0x4f, 0x6e, 0x6c, 0x79, 0x18, 0x09, 0x20, 0x01, 0x28, 0x08,
	0x52, 0x09, 0x69, 0x73, 0x52, 0x75, 0x6e, 0x4f, 0x6e, 0x6c, 0x79, 0x32, 0x59, 0x0a, 0x0f, 0x45,
	0x78, 0x65, 0x63, 0x75, 0x74, 0x6f, 0x72, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x12, 0x46,
	0x0a, 0x0f, 0x48, 0x61, 0x6e, 0x64, 0x6c, 0x65, 0x45, 0x78, 0x65, 0x63, 0x75, 0x74, 0x69, 0x6f,
	0x6e, 0x12, 0x17, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2e, 0x45, 0x78, 0x65, 0x63, 0x75, 0x74,
	0x69, 0x6f, 0x6e, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x18, 0x2e, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x2e, 0x45, 0x78, 0x65, 0x63, 0x75, 0x74, 0x69, 0x6f, 0x6e, 0x52, 0x65, 0x73, 0x70,
	0x6f, 0x6e, 0x73, 0x65, 0x28, 0x01, 0x42, 0x0e, 0x5a, 0x0c, 0x2e, 0x2f, 0x65, 0x78, 0x65, 0x63,
	0x43, 0x6c, 0x69, 0x65, 0x6e, 0x74, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_execClient_proto_rawDescOnce sync.Once
	file_execClient_proto_rawDescData = file_execClient_proto_rawDesc
)

func file_execClient_proto_rawDescGZIP() []byte {
	file_execClient_proto_rawDescOnce.Do(func() {
		file_execClient_proto_rawDescData = protoimpl.X.CompressGZIP(file_execClient_proto_rawDescData)
	})
	return file_execClient_proto_rawDescData
}

var file_execClient_proto_msgTypes = make([]protoimpl.MessageInfo, 2)
var file_execClient_proto_goTypes = []interface{}{
	(*ExecutionRequest)(nil),  // 0: proto.ExecutionRequest
	(*ExecutionResponse)(nil), // 1: proto.ExecutionResponse
}
var file_execClient_proto_depIdxs = []int32{
	0, // 0: proto.ExecutorService.HandleExecution:input_type -> proto.ExecutionRequest
	1, // 1: proto.ExecutorService.HandleExecution:output_type -> proto.ExecutionResponse
	1, // [1:2] is the sub-list for method output_type
	0, // [0:1] is the sub-list for method input_type
	0, // [0:0] is the sub-list for extension type_name
	0, // [0:0] is the sub-list for extension extendee
	0, // [0:0] is the sub-list for field type_name
}

func init() { file_execClient_proto_init() }
func file_execClient_proto_init() {
	if File_execClient_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_execClient_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
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
		file_execClient_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
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
			RawDescriptor: file_execClient_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   2,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_execClient_proto_goTypes,
		DependencyIndexes: file_execClient_proto_depIdxs,
		MessageInfos:      file_execClient_proto_msgTypes,
	}.Build()
	File_execClient_proto = out.File
	file_execClient_proto_rawDesc = nil
	file_execClient_proto_goTypes = nil
	file_execClient_proto_depIdxs = nil
}
