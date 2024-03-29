// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.26.0
// 	protoc        v3.19.4
// source: grpc_application/grpc_application.proto

package grpc_application

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

	Idx  int32  `protobuf:"varint,1,opt,name=idx,proto3" json:"idx,omitempty"`
	Name string `protobuf:"bytes,2,opt,name=name,proto3" json:"name,omitempty"`
}

func (x *DataSet) Reset() {
	*x = DataSet{}
	if protoimpl.UnsafeEnabled {
		mi := &file_grpc_application_grpc_application_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *DataSet) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*DataSet) ProtoMessage() {}

func (x *DataSet) ProtoReflect() protoreflect.Message {
	mi := &file_grpc_application_grpc_application_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
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
	return file_grpc_application_grpc_application_proto_rawDescGZIP(), []int{0}
}

func (x *DataSet) GetIdx() int32 {
	if x != nil {
		return x.Idx
	}
	return 0
}

func (x *DataSet) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

type DataRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Idx int32 `protobuf:"varint,1,opt,name=idx,proto3" json:"idx,omitempty"`
}

func (x *DataRequest) Reset() {
	*x = DataRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_grpc_application_grpc_application_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *DataRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*DataRequest) ProtoMessage() {}

func (x *DataRequest) ProtoReflect() protoreflect.Message {
	mi := &file_grpc_application_grpc_application_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use DataRequest.ProtoReflect.Descriptor instead.
func (*DataRequest) Descriptor() ([]byte, []int) {
	return file_grpc_application_grpc_application_proto_rawDescGZIP(), []int{1}
}

func (x *DataRequest) GetIdx() int32 {
	if x != nil {
		return x.Idx
	}
	return 0
}

type DataResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	DataObject *DataSet `protobuf:"bytes,1,opt,name=data_object,json=dataObject,proto3" json:"data_object,omitempty"`
}

func (x *DataResponse) Reset() {
	*x = DataResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_grpc_application_grpc_application_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *DataResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*DataResponse) ProtoMessage() {}

func (x *DataResponse) ProtoReflect() protoreflect.Message {
	mi := &file_grpc_application_grpc_application_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use DataResponse.ProtoReflect.Descriptor instead.
func (*DataResponse) Descriptor() ([]byte, []int) {
	return file_grpc_application_grpc_application_proto_rawDescGZIP(), []int{2}
}

func (x *DataResponse) GetDataObject() *DataSet {
	if x != nil {
		return x.DataObject
	}
	return nil
}

type DataListRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields
}

func (x *DataListRequest) Reset() {
	*x = DataListRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_grpc_application_grpc_application_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *DataListRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*DataListRequest) ProtoMessage() {}

func (x *DataListRequest) ProtoReflect() protoreflect.Message {
	mi := &file_grpc_application_grpc_application_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use DataListRequest.ProtoReflect.Descriptor instead.
func (*DataListRequest) Descriptor() ([]byte, []int) {
	return file_grpc_application_grpc_application_proto_rawDescGZIP(), []int{3}
}

type DataListResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	DataObejctList []*DataSet `protobuf:"bytes,1,rep,name=data_obejct_list,json=dataObejctList,proto3" json:"data_obejct_list,omitempty"`
}

func (x *DataListResponse) Reset() {
	*x = DataListResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_grpc_application_grpc_application_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *DataListResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*DataListResponse) ProtoMessage() {}

func (x *DataListResponse) ProtoReflect() protoreflect.Message {
	mi := &file_grpc_application_grpc_application_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use DataListResponse.ProtoReflect.Descriptor instead.
func (*DataListResponse) Descriptor() ([]byte, []int) {
	return file_grpc_application_grpc_application_proto_rawDescGZIP(), []int{4}
}

func (x *DataListResponse) GetDataObejctList() []*DataSet {
	if x != nil {
		return x.DataObejctList
	}
	return nil
}

var File_grpc_application_grpc_application_proto protoreflect.FileDescriptor

var file_grpc_application_grpc_application_proto_rawDesc = []byte{
	0x0a, 0x27, 0x67, 0x72, 0x70, 0x63, 0x5f, 0x61, 0x70, 0x70, 0x6c, 0x69, 0x63, 0x61, 0x74, 0x69,
	0x6f, 0x6e, 0x2f, 0x67, 0x72, 0x70, 0x63, 0x5f, 0x61, 0x70, 0x70, 0x6c, 0x69, 0x63, 0x61, 0x74,
	0x69, 0x6f, 0x6e, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x10, 0x67, 0x72, 0x70, 0x63, 0x5f,
	0x61, 0x70, 0x70, 0x6c, 0x69, 0x63, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x22, 0x2f, 0x0a, 0x07, 0x44,
	0x61, 0x74, 0x61, 0x53, 0x65, 0x74, 0x12, 0x10, 0x0a, 0x03, 0x69, 0x64, 0x78, 0x18, 0x01, 0x20,
	0x01, 0x28, 0x05, 0x52, 0x03, 0x69, 0x64, 0x78, 0x12, 0x12, 0x0a, 0x04, 0x6e, 0x61, 0x6d, 0x65,
	0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x22, 0x1f, 0x0a, 0x0b,
	0x44, 0x61, 0x74, 0x61, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x10, 0x0a, 0x03, 0x69,
	0x64, 0x78, 0x18, 0x01, 0x20, 0x01, 0x28, 0x05, 0x52, 0x03, 0x69, 0x64, 0x78, 0x22, 0x4a, 0x0a,
	0x0c, 0x44, 0x61, 0x74, 0x61, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x3a, 0x0a,
	0x0b, 0x64, 0x61, 0x74, 0x61, 0x5f, 0x6f, 0x62, 0x6a, 0x65, 0x63, 0x74, 0x18, 0x01, 0x20, 0x01,
	0x28, 0x0b, 0x32, 0x19, 0x2e, 0x67, 0x72, 0x70, 0x63, 0x5f, 0x61, 0x70, 0x70, 0x6c, 0x69, 0x63,
	0x61, 0x74, 0x69, 0x6f, 0x6e, 0x2e, 0x44, 0x61, 0x74, 0x61, 0x53, 0x65, 0x74, 0x52, 0x0a, 0x64,
	0x61, 0x74, 0x61, 0x4f, 0x62, 0x6a, 0x65, 0x63, 0x74, 0x22, 0x11, 0x0a, 0x0f, 0x44, 0x61, 0x74,
	0x61, 0x4c, 0x69, 0x73, 0x74, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x22, 0x57, 0x0a, 0x10,
	0x44, 0x61, 0x74, 0x61, 0x4c, 0x69, 0x73, 0x74, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65,
	0x12, 0x43, 0x0a, 0x10, 0x64, 0x61, 0x74, 0x61, 0x5f, 0x6f, 0x62, 0x65, 0x6a, 0x63, 0x74, 0x5f,
	0x6c, 0x69, 0x73, 0x74, 0x18, 0x01, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x19, 0x2e, 0x67, 0x72, 0x70,
	0x63, 0x5f, 0x61, 0x70, 0x70, 0x6c, 0x69, 0x63, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x2e, 0x44, 0x61,
	0x74, 0x61, 0x53, 0x65, 0x74, 0x52, 0x0e, 0x64, 0x61, 0x74, 0x61, 0x4f, 0x62, 0x65, 0x6a, 0x63,
	0x74, 0x4c, 0x69, 0x73, 0x74, 0x32, 0xa2, 0x01, 0x0a, 0x03, 0x41, 0x70, 0x70, 0x12, 0x48, 0x0a,
	0x07, 0x47, 0x65, 0x74, 0x44, 0x61, 0x74, 0x61, 0x12, 0x1d, 0x2e, 0x67, 0x72, 0x70, 0x63, 0x5f,
	0x61, 0x70, 0x70, 0x6c, 0x69, 0x63, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x2e, 0x44, 0x61, 0x74, 0x61,
	0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x1e, 0x2e, 0x67, 0x72, 0x70, 0x63, 0x5f, 0x61,
	0x70, 0x70, 0x6c, 0x69, 0x63, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x2e, 0x44, 0x61, 0x74, 0x61, 0x52,
	0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x51, 0x0a, 0x08, 0x4c, 0x69, 0x73, 0x74, 0x44,
	0x61, 0x74, 0x61, 0x12, 0x21, 0x2e, 0x67, 0x72, 0x70, 0x63, 0x5f, 0x61, 0x70, 0x70, 0x6c, 0x69,
	0x63, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x2e, 0x44, 0x61, 0x74, 0x61, 0x4c, 0x69, 0x73, 0x74, 0x52,
	0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x22, 0x2e, 0x67, 0x72, 0x70, 0x63, 0x5f, 0x61, 0x70,
	0x70, 0x6c, 0x69, 0x63, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x2e, 0x44, 0x61, 0x74, 0x61, 0x4c, 0x69,
	0x73, 0x74, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x42, 0x1a, 0x5a, 0x18, 0x67, 0x6f,
	0x2d, 0x67, 0x72, 0x70, 0x63, 0x2f, 0x67, 0x72, 0x70, 0x63, 0x5f, 0x61, 0x70, 0x70, 0x6c, 0x69,
	0x63, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_grpc_application_grpc_application_proto_rawDescOnce sync.Once
	file_grpc_application_grpc_application_proto_rawDescData = file_grpc_application_grpc_application_proto_rawDesc
)

func file_grpc_application_grpc_application_proto_rawDescGZIP() []byte {
	file_grpc_application_grpc_application_proto_rawDescOnce.Do(func() {
		file_grpc_application_grpc_application_proto_rawDescData = protoimpl.X.CompressGZIP(file_grpc_application_grpc_application_proto_rawDescData)
	})
	return file_grpc_application_grpc_application_proto_rawDescData
}

var file_grpc_application_grpc_application_proto_msgTypes = make([]protoimpl.MessageInfo, 5)
var file_grpc_application_grpc_application_proto_goTypes = []interface{}{
	(*DataSet)(nil),          // 0: grpc_application.DataSet
	(*DataRequest)(nil),      // 1: grpc_application.DataRequest
	(*DataResponse)(nil),     // 2: grpc_application.DataResponse
	(*DataListRequest)(nil),  // 3: grpc_application.DataListRequest
	(*DataListResponse)(nil), // 4: grpc_application.DataListResponse
}
var file_grpc_application_grpc_application_proto_depIdxs = []int32{
	0, // 0: grpc_application.DataResponse.data_object:type_name -> grpc_application.DataSet
	0, // 1: grpc_application.DataListResponse.data_obejct_list:type_name -> grpc_application.DataSet
	1, // 2: grpc_application.App.GetData:input_type -> grpc_application.DataRequest
	3, // 3: grpc_application.App.ListData:input_type -> grpc_application.DataListRequest
	2, // 4: grpc_application.App.GetData:output_type -> grpc_application.DataResponse
	4, // 5: grpc_application.App.ListData:output_type -> grpc_application.DataListResponse
	4, // [4:6] is the sub-list for method output_type
	2, // [2:4] is the sub-list for method input_type
	2, // [2:2] is the sub-list for extension type_name
	2, // [2:2] is the sub-list for extension extendee
	0, // [0:2] is the sub-list for field type_name
}

func init() { file_grpc_application_grpc_application_proto_init() }
func file_grpc_application_grpc_application_proto_init() {
	if File_grpc_application_grpc_application_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_grpc_application_grpc_application_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*DataSet); i {
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
		file_grpc_application_grpc_application_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*DataRequest); i {
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
		file_grpc_application_grpc_application_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*DataResponse); i {
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
		file_grpc_application_grpc_application_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*DataListRequest); i {
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
		file_grpc_application_grpc_application_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*DataListResponse); i {
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
			RawDescriptor: file_grpc_application_grpc_application_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   5,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_grpc_application_grpc_application_proto_goTypes,
		DependencyIndexes: file_grpc_application_grpc_application_proto_depIdxs,
		MessageInfos:      file_grpc_application_grpc_application_proto_msgTypes,
	}.Build()
	File_grpc_application_grpc_application_proto = out.File
	file_grpc_application_grpc_application_proto_rawDesc = nil
	file_grpc_application_grpc_application_proto_goTypes = nil
	file_grpc_application_grpc_application_proto_depIdxs = nil
}
