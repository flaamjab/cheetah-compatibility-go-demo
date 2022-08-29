// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.28.1
// 	protoc        v3.12.4
// source: system.compatibility.external.proto

package compatibility_grpc

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

type CheckCompatibilityResponse_Status int32

const (
	//*
	//Версия клиента поддерживается, в обновлении нет необходимости
	CheckCompatibilityResponse_Supported CheckCompatibilityResponse_Status = 0
	//*
	//Планируется окончания поддержки - за 4-24 часа до окончания
	CheckCompatibilityResponse_PlanningUnsupportedAfterSomeHours CheckCompatibilityResponse_Status = 1
	//*
	//Планируется окончания поддержки - за 0-4 часа до окончания
	CheckCompatibilityResponse_PlanningUnsupportedSoon CheckCompatibilityResponse_Status = 2
	//*
	//Данная версия клиента не поддерживается, если в это время клиент в битве - необходимо обновить клиент после
	//окончания битвы, если клиент только что запущен или в процессе запуска - необходимо принудительное обновление
	CheckCompatibilityResponse_Unsupported CheckCompatibilityResponse_Status = 3
)

// Enum value maps for CheckCompatibilityResponse_Status.
var (
	CheckCompatibilityResponse_Status_name = map[int32]string{
		0: "Supported",
		1: "PlanningUnsupportedAfterSomeHours",
		2: "PlanningUnsupportedSoon",
		3: "Unsupported",
	}
	CheckCompatibilityResponse_Status_value = map[string]int32{
		"Supported":                         0,
		"PlanningUnsupportedAfterSomeHours": 1,
		"PlanningUnsupportedSoon":           2,
		"Unsupported":                       3,
	}
)

func (x CheckCompatibilityResponse_Status) Enum() *CheckCompatibilityResponse_Status {
	p := new(CheckCompatibilityResponse_Status)
	*p = x
	return p
}

func (x CheckCompatibilityResponse_Status) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (CheckCompatibilityResponse_Status) Descriptor() protoreflect.EnumDescriptor {
	return file_system_compatibility_external_proto_enumTypes[0].Descriptor()
}

func (CheckCompatibilityResponse_Status) Type() protoreflect.EnumType {
	return &file_system_compatibility_external_proto_enumTypes[0]
}

func (x CheckCompatibilityResponse_Status) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use CheckCompatibilityResponse_Status.Descriptor instead.
func (CheckCompatibilityResponse_Status) EnumDescriptor() ([]byte, []int) {
	return file_system_compatibility_external_proto_rawDescGZIP(), []int{1, 0}
}

type CheckCompatibilityRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Version string `protobuf:"bytes,1,opt,name=version,proto3" json:"version,omitempty"`
}

func (x *CheckCompatibilityRequest) Reset() {
	*x = CheckCompatibilityRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_system_compatibility_external_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *CheckCompatibilityRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CheckCompatibilityRequest) ProtoMessage() {}

func (x *CheckCompatibilityRequest) ProtoReflect() protoreflect.Message {
	mi := &file_system_compatibility_external_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CheckCompatibilityRequest.ProtoReflect.Descriptor instead.
func (*CheckCompatibilityRequest) Descriptor() ([]byte, []int) {
	return file_system_compatibility_external_proto_rawDescGZIP(), []int{0}
}

func (x *CheckCompatibilityRequest) GetVersion() string {
	if x != nil {
		return x.Version
	}
	return ""
}

type CheckCompatibilityResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Status CheckCompatibilityResponse_Status `protobuf:"varint,1,opt,name=status,proto3,enum=cheetah.system.compatibility.CheckCompatibilityResponse_Status" json:"status,omitempty"`
}

func (x *CheckCompatibilityResponse) Reset() {
	*x = CheckCompatibilityResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_system_compatibility_external_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *CheckCompatibilityResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CheckCompatibilityResponse) ProtoMessage() {}

func (x *CheckCompatibilityResponse) ProtoReflect() protoreflect.Message {
	mi := &file_system_compatibility_external_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CheckCompatibilityResponse.ProtoReflect.Descriptor instead.
func (*CheckCompatibilityResponse) Descriptor() ([]byte, []int) {
	return file_system_compatibility_external_proto_rawDescGZIP(), []int{1}
}

func (x *CheckCompatibilityResponse) GetStatus() CheckCompatibilityResponse_Status {
	if x != nil {
		return x.Status
	}
	return CheckCompatibilityResponse_Supported
}

var File_system_compatibility_external_proto protoreflect.FileDescriptor

var file_system_compatibility_external_proto_rawDesc = []byte{
	0x0a, 0x23, 0x73, 0x79, 0x73, 0x74, 0x65, 0x6d, 0x2e, 0x63, 0x6f, 0x6d, 0x70, 0x61, 0x74, 0x69,
	0x62, 0x69, 0x6c, 0x69, 0x74, 0x79, 0x2e, 0x65, 0x78, 0x74, 0x65, 0x72, 0x6e, 0x61, 0x6c, 0x2e,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x1c, 0x63, 0x68, 0x65, 0x65, 0x74, 0x61, 0x68, 0x2e, 0x73,
	0x79, 0x73, 0x74, 0x65, 0x6d, 0x2e, 0x63, 0x6f, 0x6d, 0x70, 0x61, 0x74, 0x69, 0x62, 0x69, 0x6c,
	0x69, 0x74, 0x79, 0x22, 0x35, 0x0a, 0x19, 0x43, 0x68, 0x65, 0x63, 0x6b, 0x43, 0x6f, 0x6d, 0x70,
	0x61, 0x74, 0x69, 0x62, 0x69, 0x6c, 0x69, 0x74, 0x79, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74,
	0x12, 0x18, 0x0a, 0x07, 0x76, 0x65, 0x72, 0x73, 0x69, 0x6f, 0x6e, 0x18, 0x01, 0x20, 0x01, 0x28,
	0x09, 0x52, 0x07, 0x76, 0x65, 0x72, 0x73, 0x69, 0x6f, 0x6e, 0x22, 0xe3, 0x01, 0x0a, 0x1a, 0x43,
	0x68, 0x65, 0x63, 0x6b, 0x43, 0x6f, 0x6d, 0x70, 0x61, 0x74, 0x69, 0x62, 0x69, 0x6c, 0x69, 0x74,
	0x79, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x57, 0x0a, 0x06, 0x73, 0x74, 0x61,
	0x74, 0x75, 0x73, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0e, 0x32, 0x3f, 0x2e, 0x63, 0x68, 0x65, 0x65,
	0x74, 0x61, 0x68, 0x2e, 0x73, 0x79, 0x73, 0x74, 0x65, 0x6d, 0x2e, 0x63, 0x6f, 0x6d, 0x70, 0x61,
	0x74, 0x69, 0x62, 0x69, 0x6c, 0x69, 0x74, 0x79, 0x2e, 0x43, 0x68, 0x65, 0x63, 0x6b, 0x43, 0x6f,
	0x6d, 0x70, 0x61, 0x74, 0x69, 0x62, 0x69, 0x6c, 0x69, 0x74, 0x79, 0x52, 0x65, 0x73, 0x70, 0x6f,
	0x6e, 0x73, 0x65, 0x2e, 0x53, 0x74, 0x61, 0x74, 0x75, 0x73, 0x52, 0x06, 0x73, 0x74, 0x61, 0x74,
	0x75, 0x73, 0x22, 0x6c, 0x0a, 0x06, 0x53, 0x74, 0x61, 0x74, 0x75, 0x73, 0x12, 0x0d, 0x0a, 0x09,
	0x53, 0x75, 0x70, 0x70, 0x6f, 0x72, 0x74, 0x65, 0x64, 0x10, 0x00, 0x12, 0x25, 0x0a, 0x21, 0x50,
	0x6c, 0x61, 0x6e, 0x6e, 0x69, 0x6e, 0x67, 0x55, 0x6e, 0x73, 0x75, 0x70, 0x70, 0x6f, 0x72, 0x74,
	0x65, 0x64, 0x41, 0x66, 0x74, 0x65, 0x72, 0x53, 0x6f, 0x6d, 0x65, 0x48, 0x6f, 0x75, 0x72, 0x73,
	0x10, 0x01, 0x12, 0x1b, 0x0a, 0x17, 0x50, 0x6c, 0x61, 0x6e, 0x6e, 0x69, 0x6e, 0x67, 0x55, 0x6e,
	0x73, 0x75, 0x70, 0x70, 0x6f, 0x72, 0x74, 0x65, 0x64, 0x53, 0x6f, 0x6f, 0x6e, 0x10, 0x02, 0x12,
	0x0f, 0x0a, 0x0b, 0x55, 0x6e, 0x73, 0x75, 0x70, 0x70, 0x6f, 0x72, 0x74, 0x65, 0x64, 0x10, 0x03,
	0x32, 0xa0, 0x01, 0x0a, 0x14, 0x43, 0x6f, 0x6d, 0x70, 0x61, 0x74, 0x69, 0x62, 0x69, 0x6c, 0x69,
	0x74, 0x79, 0x43, 0x68, 0x65, 0x63, 0x6b, 0x65, 0x72, 0x12, 0x87, 0x01, 0x0a, 0x12, 0x43, 0x68,
	0x65, 0x63, 0x6b, 0x43, 0x6f, 0x6d, 0x70, 0x61, 0x74, 0x69, 0x62, 0x69, 0x6c, 0x69, 0x74, 0x79,
	0x12, 0x37, 0x2e, 0x63, 0x68, 0x65, 0x65, 0x74, 0x61, 0x68, 0x2e, 0x73, 0x79, 0x73, 0x74, 0x65,
	0x6d, 0x2e, 0x63, 0x6f, 0x6d, 0x70, 0x61, 0x74, 0x69, 0x62, 0x69, 0x6c, 0x69, 0x74, 0x79, 0x2e,
	0x43, 0x68, 0x65, 0x63, 0x6b, 0x43, 0x6f, 0x6d, 0x70, 0x61, 0x74, 0x69, 0x62, 0x69, 0x6c, 0x69,
	0x74, 0x79, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x38, 0x2e, 0x63, 0x68, 0x65, 0x65,
	0x74, 0x61, 0x68, 0x2e, 0x73, 0x79, 0x73, 0x74, 0x65, 0x6d, 0x2e, 0x63, 0x6f, 0x6d, 0x70, 0x61,
	0x74, 0x69, 0x62, 0x69, 0x6c, 0x69, 0x74, 0x79, 0x2e, 0x43, 0x68, 0x65, 0x63, 0x6b, 0x43, 0x6f,
	0x6d, 0x70, 0x61, 0x74, 0x69, 0x62, 0x69, 0x6c, 0x69, 0x74, 0x79, 0x52, 0x65, 0x73, 0x70, 0x6f,
	0x6e, 0x73, 0x65, 0x42, 0x14, 0x5a, 0x12, 0x63, 0x6f, 0x6d, 0x70, 0x61, 0x74, 0x69, 0x62, 0x69,
	0x6c, 0x69, 0x74, 0x79, 0x2e, 0x67, 0x72, 0x70, 0x63, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x33,
}

var (
	file_system_compatibility_external_proto_rawDescOnce sync.Once
	file_system_compatibility_external_proto_rawDescData = file_system_compatibility_external_proto_rawDesc
)

func file_system_compatibility_external_proto_rawDescGZIP() []byte {
	file_system_compatibility_external_proto_rawDescOnce.Do(func() {
		file_system_compatibility_external_proto_rawDescData = protoimpl.X.CompressGZIP(file_system_compatibility_external_proto_rawDescData)
	})
	return file_system_compatibility_external_proto_rawDescData
}

var file_system_compatibility_external_proto_enumTypes = make([]protoimpl.EnumInfo, 1)
var file_system_compatibility_external_proto_msgTypes = make([]protoimpl.MessageInfo, 2)
var file_system_compatibility_external_proto_goTypes = []interface{}{
	(CheckCompatibilityResponse_Status)(0), // 0: cheetah.system.compatibility.CheckCompatibilityResponse.Status
	(*CheckCompatibilityRequest)(nil),      // 1: cheetah.system.compatibility.CheckCompatibilityRequest
	(*CheckCompatibilityResponse)(nil),     // 2: cheetah.system.compatibility.CheckCompatibilityResponse
}
var file_system_compatibility_external_proto_depIdxs = []int32{
	0, // 0: cheetah.system.compatibility.CheckCompatibilityResponse.status:type_name -> cheetah.system.compatibility.CheckCompatibilityResponse.Status
	1, // 1: cheetah.system.compatibility.CompatibilityChecker.CheckCompatibility:input_type -> cheetah.system.compatibility.CheckCompatibilityRequest
	2, // 2: cheetah.system.compatibility.CompatibilityChecker.CheckCompatibility:output_type -> cheetah.system.compatibility.CheckCompatibilityResponse
	2, // [2:3] is the sub-list for method output_type
	1, // [1:2] is the sub-list for method input_type
	1, // [1:1] is the sub-list for extension type_name
	1, // [1:1] is the sub-list for extension extendee
	0, // [0:1] is the sub-list for field type_name
}

func init() { file_system_compatibility_external_proto_init() }
func file_system_compatibility_external_proto_init() {
	if File_system_compatibility_external_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_system_compatibility_external_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*CheckCompatibilityRequest); i {
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
		file_system_compatibility_external_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*CheckCompatibilityResponse); i {
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
			RawDescriptor: file_system_compatibility_external_proto_rawDesc,
			NumEnums:      1,
			NumMessages:   2,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_system_compatibility_external_proto_goTypes,
		DependencyIndexes: file_system_compatibility_external_proto_depIdxs,
		EnumInfos:         file_system_compatibility_external_proto_enumTypes,
		MessageInfos:      file_system_compatibility_external_proto_msgTypes,
	}.Build()
	File_system_compatibility_external_proto = out.File
	file_system_compatibility_external_proto_rawDesc = nil
	file_system_compatibility_external_proto_goTypes = nil
	file_system_compatibility_external_proto_depIdxs = nil
}