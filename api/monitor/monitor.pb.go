// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.31.0
// 	protoc        v4.22.2
// source: api/proto/monitor.proto

package monitor

import (
	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
	emptypb "google.golang.org/protobuf/types/known/emptypb"
	reflect "reflect"
	sync "sync"
)

const (
	// Verify that this generated code is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(20 - protoimpl.MinVersion)
	// Verify that runtime/protoimpl is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(protoimpl.MaxVersion - 20)
)

// FileServerInfo the file server basic info
type MonitorMessage struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	FileInfo *FileInfo `protobuf:"bytes,1,opt,name=file_info,json=fileInfo,proto3" json:"file_info,omitempty"`
	// Action the action of file change
	Action int32 `protobuf:"varint,2,opt,name=action,proto3" json:"action,omitempty"`
	// BaseUrl the base url of file server
	BaseUrl string `protobuf:"bytes,3,opt,name=base_url,json=baseUrl,proto3" json:"base_url,omitempty"`
}

func (x *MonitorMessage) Reset() {
	*x = MonitorMessage{}
	if protoimpl.UnsafeEnabled {
		mi := &file_api_proto_monitor_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *MonitorMessage) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*MonitorMessage) ProtoMessage() {}

func (x *MonitorMessage) ProtoReflect() protoreflect.Message {
	mi := &file_api_proto_monitor_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use MonitorMessage.ProtoReflect.Descriptor instead.
func (*MonitorMessage) Descriptor() ([]byte, []int) {
	return file_api_proto_monitor_proto_rawDescGZIP(), []int{0}
}

func (x *MonitorMessage) GetFileInfo() *FileInfo {
	if x != nil {
		return x.FileInfo
	}
	return nil
}

func (x *MonitorMessage) GetAction() int32 {
	if x != nil {
		return x.Action
	}
	return 0
}

func (x *MonitorMessage) GetBaseUrl() string {
	if x != nil {
		return x.BaseUrl
	}
	return ""
}

// FileInfo the basic file info description
type FileInfo struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// Path the file path
	Path string `protobuf:"bytes,1,opt,name=path,proto3" json:"path,omitempty"`
	// IsDir is a dir the path
	IsDir int32 `protobuf:"varint,2,opt,name=is_dir,json=isDir,proto3" json:"is_dir,omitempty"`
	// Size the size of path for bytes
	Size int64 `protobuf:"varint,3,opt,name=size,proto3" json:"size,omitempty"`
	// Hash calculate the path hash value, if the path is a file
	Hash string `protobuf:"bytes,4,opt,name=hash,proto3" json:"hash,omitempty"`
	// HashValues the hash value of the entire file and first chunk and some checkpoints
	HashValues []*HashValue `protobuf:"bytes,5,rep,name=hash_values,json=hashValues,proto3" json:"hash_values,omitempty"`
	// CTime creation time, unix sec
	CTime int64 `protobuf:"varint,6,opt,name=c_time,json=cTime,proto3" json:"c_time,omitempty"`
	// ATime last access time, unix sec
	ATime int64 `protobuf:"varint,7,opt,name=a_time,json=aTime,proto3" json:"a_time,omitempty"`
	// MTime last modify time, unix sec
	MTime int64 `protobuf:"varint,8,opt,name=m_time,json=mTime,proto3" json:"m_time,omitempty"`
	// LinkTo link to the real file
	LinkTo string `protobuf:"bytes,9,opt,name=link_to,json=linkTo,proto3" json:"link_to,omitempty"`
}

func (x *FileInfo) Reset() {
	*x = FileInfo{}
	if protoimpl.UnsafeEnabled {
		mi := &file_api_proto_monitor_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *FileInfo) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*FileInfo) ProtoMessage() {}

func (x *FileInfo) ProtoReflect() protoreflect.Message {
	mi := &file_api_proto_monitor_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use FileInfo.ProtoReflect.Descriptor instead.
func (*FileInfo) Descriptor() ([]byte, []int) {
	return file_api_proto_monitor_proto_rawDescGZIP(), []int{1}
}

func (x *FileInfo) GetPath() string {
	if x != nil {
		return x.Path
	}
	return ""
}

func (x *FileInfo) GetIsDir() int32 {
	if x != nil {
		return x.IsDir
	}
	return 0
}

func (x *FileInfo) GetSize() int64 {
	if x != nil {
		return x.Size
	}
	return 0
}

func (x *FileInfo) GetHash() string {
	if x != nil {
		return x.Hash
	}
	return ""
}

func (x *FileInfo) GetHashValues() []*HashValue {
	if x != nil {
		return x.HashValues
	}
	return nil
}

func (x *FileInfo) GetCTime() int64 {
	if x != nil {
		return x.CTime
	}
	return 0
}

func (x *FileInfo) GetATime() int64 {
	if x != nil {
		return x.ATime
	}
	return 0
}

func (x *FileInfo) GetMTime() int64 {
	if x != nil {
		return x.MTime
	}
	return 0
}

func (x *FileInfo) GetLinkTo() string {
	if x != nil {
		return x.LinkTo
	}
	return ""
}

// HashValue the file hash info
type HashValue struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// Offset the file data to calculate the hash value from zero to offset
	Offset int64 `protobuf:"varint,1,opt,name=offset,proto3" json:"offset,omitempty"`
	// Hash the file checkpoint hash value
	Hash string `protobuf:"bytes,2,opt,name=hash,proto3" json:"hash,omitempty"`
}

func (x *HashValue) Reset() {
	*x = HashValue{}
	if protoimpl.UnsafeEnabled {
		mi := &file_api_proto_monitor_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *HashValue) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*HashValue) ProtoMessage() {}

func (x *HashValue) ProtoReflect() protoreflect.Message {
	mi := &file_api_proto_monitor_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use HashValue.ProtoReflect.Descriptor instead.
func (*HashValue) Descriptor() ([]byte, []int) {
	return file_api_proto_monitor_proto_rawDescGZIP(), []int{2}
}

func (x *HashValue) GetOffset() int64 {
	if x != nil {
		return x.Offset
	}
	return 0
}

func (x *HashValue) GetHash() string {
	if x != nil {
		return x.Hash
	}
	return ""
}

var File_api_proto_monitor_proto protoreflect.FileDescriptor

var file_api_proto_monitor_proto_rawDesc = []byte{
	0x0a, 0x17, 0x61, 0x70, 0x69, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2f, 0x6d, 0x6f, 0x6e, 0x69,
	0x74, 0x6f, 0x72, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x07, 0x6d, 0x6f, 0x6e, 0x69, 0x74,
	0x6f, 0x72, 0x1a, 0x1b, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x62, 0x75, 0x66, 0x2f, 0x65, 0x6d, 0x70, 0x74, 0x79, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22,
	0x73, 0x0a, 0x0e, 0x4d, 0x6f, 0x6e, 0x69, 0x74, 0x6f, 0x72, 0x4d, 0x65, 0x73, 0x73, 0x61, 0x67,
	0x65, 0x12, 0x2e, 0x0a, 0x09, 0x66, 0x69, 0x6c, 0x65, 0x5f, 0x69, 0x6e, 0x66, 0x6f, 0x18, 0x01,
	0x20, 0x01, 0x28, 0x0b, 0x32, 0x11, 0x2e, 0x6d, 0x6f, 0x6e, 0x69, 0x74, 0x6f, 0x72, 0x2e, 0x46,
	0x69, 0x6c, 0x65, 0x49, 0x6e, 0x66, 0x6f, 0x52, 0x08, 0x66, 0x69, 0x6c, 0x65, 0x49, 0x6e, 0x66,
	0x6f, 0x12, 0x16, 0x0a, 0x06, 0x61, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x18, 0x02, 0x20, 0x01, 0x28,
	0x05, 0x52, 0x06, 0x61, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x12, 0x19, 0x0a, 0x08, 0x62, 0x61, 0x73,
	0x65, 0x5f, 0x75, 0x72, 0x6c, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x07, 0x62, 0x61, 0x73,
	0x65, 0x55, 0x72, 0x6c, 0x22, 0xf0, 0x01, 0x0a, 0x08, 0x46, 0x69, 0x6c, 0x65, 0x49, 0x6e, 0x66,
	0x6f, 0x12, 0x12, 0x0a, 0x04, 0x70, 0x61, 0x74, 0x68, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52,
	0x04, 0x70, 0x61, 0x74, 0x68, 0x12, 0x15, 0x0a, 0x06, 0x69, 0x73, 0x5f, 0x64, 0x69, 0x72, 0x18,
	0x02, 0x20, 0x01, 0x28, 0x05, 0x52, 0x05, 0x69, 0x73, 0x44, 0x69, 0x72, 0x12, 0x12, 0x0a, 0x04,
	0x73, 0x69, 0x7a, 0x65, 0x18, 0x03, 0x20, 0x01, 0x28, 0x03, 0x52, 0x04, 0x73, 0x69, 0x7a, 0x65,
	0x12, 0x12, 0x0a, 0x04, 0x68, 0x61, 0x73, 0x68, 0x18, 0x04, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04,
	0x68, 0x61, 0x73, 0x68, 0x12, 0x33, 0x0a, 0x0b, 0x68, 0x61, 0x73, 0x68, 0x5f, 0x76, 0x61, 0x6c,
	0x75, 0x65, 0x73, 0x18, 0x05, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x12, 0x2e, 0x6d, 0x6f, 0x6e, 0x69,
	0x74, 0x6f, 0x72, 0x2e, 0x48, 0x61, 0x73, 0x68, 0x56, 0x61, 0x6c, 0x75, 0x65, 0x52, 0x0a, 0x68,
	0x61, 0x73, 0x68, 0x56, 0x61, 0x6c, 0x75, 0x65, 0x73, 0x12, 0x15, 0x0a, 0x06, 0x63, 0x5f, 0x74,
	0x69, 0x6d, 0x65, 0x18, 0x06, 0x20, 0x01, 0x28, 0x03, 0x52, 0x05, 0x63, 0x54, 0x69, 0x6d, 0x65,
	0x12, 0x15, 0x0a, 0x06, 0x61, 0x5f, 0x74, 0x69, 0x6d, 0x65, 0x18, 0x07, 0x20, 0x01, 0x28, 0x03,
	0x52, 0x05, 0x61, 0x54, 0x69, 0x6d, 0x65, 0x12, 0x15, 0x0a, 0x06, 0x6d, 0x5f, 0x74, 0x69, 0x6d,
	0x65, 0x18, 0x08, 0x20, 0x01, 0x28, 0x03, 0x52, 0x05, 0x6d, 0x54, 0x69, 0x6d, 0x65, 0x12, 0x17,
	0x0a, 0x07, 0x6c, 0x69, 0x6e, 0x6b, 0x5f, 0x74, 0x6f, 0x18, 0x09, 0x20, 0x01, 0x28, 0x09, 0x52,
	0x06, 0x6c, 0x69, 0x6e, 0x6b, 0x54, 0x6f, 0x22, 0x37, 0x0a, 0x09, 0x48, 0x61, 0x73, 0x68, 0x56,
	0x61, 0x6c, 0x75, 0x65, 0x12, 0x16, 0x0a, 0x06, 0x6f, 0x66, 0x66, 0x73, 0x65, 0x74, 0x18, 0x01,
	0x20, 0x01, 0x28, 0x03, 0x52, 0x06, 0x6f, 0x66, 0x66, 0x73, 0x65, 0x74, 0x12, 0x12, 0x0a, 0x04,
	0x68, 0x61, 0x73, 0x68, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x68, 0x61, 0x73, 0x68,
	0x32, 0x50, 0x0a, 0x0e, 0x4d, 0x6f, 0x6e, 0x69, 0x74, 0x6f, 0x72, 0x53, 0x65, 0x72, 0x76, 0x69,
	0x63, 0x65, 0x12, 0x3e, 0x0a, 0x07, 0x4d, 0x6f, 0x6e, 0x69, 0x74, 0x6f, 0x72, 0x12, 0x16, 0x2e,
	0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e,
	0x45, 0x6d, 0x70, 0x74, 0x79, 0x1a, 0x17, 0x2e, 0x6d, 0x6f, 0x6e, 0x69, 0x74, 0x6f, 0x72, 0x2e,
	0x4d, 0x6f, 0x6e, 0x69, 0x74, 0x6f, 0x72, 0x4d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x22, 0x00,
	0x30, 0x01, 0x42, 0x24, 0x5a, 0x22, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d,
	0x2f, 0x6e, 0x6f, 0x2d, 0x73, 0x72, 0x63, 0x2f, 0x67, 0x6f, 0x66, 0x73, 0x2f, 0x61, 0x70, 0x69,
	0x2f, 0x6d, 0x6f, 0x6e, 0x69, 0x74, 0x6f, 0x72, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_api_proto_monitor_proto_rawDescOnce sync.Once
	file_api_proto_monitor_proto_rawDescData = file_api_proto_monitor_proto_rawDesc
)

func file_api_proto_monitor_proto_rawDescGZIP() []byte {
	file_api_proto_monitor_proto_rawDescOnce.Do(func() {
		file_api_proto_monitor_proto_rawDescData = protoimpl.X.CompressGZIP(file_api_proto_monitor_proto_rawDescData)
	})
	return file_api_proto_monitor_proto_rawDescData
}

var file_api_proto_monitor_proto_msgTypes = make([]protoimpl.MessageInfo, 3)
var file_api_proto_monitor_proto_goTypes = []interface{}{
	(*MonitorMessage)(nil), // 0: monitor.MonitorMessage
	(*FileInfo)(nil),       // 1: monitor.FileInfo
	(*HashValue)(nil),      // 2: monitor.HashValue
	(*emptypb.Empty)(nil),  // 3: google.protobuf.Empty
}
var file_api_proto_monitor_proto_depIdxs = []int32{
	1, // 0: monitor.MonitorMessage.file_info:type_name -> monitor.FileInfo
	2, // 1: monitor.FileInfo.hash_values:type_name -> monitor.HashValue
	3, // 2: monitor.MonitorService.Monitor:input_type -> google.protobuf.Empty
	0, // 3: monitor.MonitorService.Monitor:output_type -> monitor.MonitorMessage
	3, // [3:4] is the sub-list for method output_type
	2, // [2:3] is the sub-list for method input_type
	2, // [2:2] is the sub-list for extension type_name
	2, // [2:2] is the sub-list for extension extendee
	0, // [0:2] is the sub-list for field type_name
}

func init() { file_api_proto_monitor_proto_init() }
func file_api_proto_monitor_proto_init() {
	if File_api_proto_monitor_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_api_proto_monitor_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*MonitorMessage); i {
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
		file_api_proto_monitor_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*FileInfo); i {
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
		file_api_proto_monitor_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*HashValue); i {
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
			RawDescriptor: file_api_proto_monitor_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   3,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_api_proto_monitor_proto_goTypes,
		DependencyIndexes: file_api_proto_monitor_proto_depIdxs,
		MessageInfos:      file_api_proto_monitor_proto_msgTypes,
	}.Build()
	File_api_proto_monitor_proto = out.File
	file_api_proto_monitor_proto_rawDesc = nil
	file_api_proto_monitor_proto_goTypes = nil
	file_api_proto_monitor_proto_depIdxs = nil
}
