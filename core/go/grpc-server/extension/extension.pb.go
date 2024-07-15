// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.32.0-devel
// 	protoc        v3.12.4
// source: extension.proto

package extension

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

type ExtensionData_Lock int32

const (
	ExtensionData_NONE        ExtensionData_Lock = 0
	ExtensionData_ADMIN       ExtensionData_Lock = 1
	ExtensionData_DECK        ExtensionData_Lock = 2
	ExtensionData_TRADE       ExtensionData_Lock = 3
	ExtensionData_CRAFT       ExtensionData_Lock = 4
	ExtensionData_OUTSOURCING ExtensionData_Lock = 5
)

// Enum value maps for ExtensionData_Lock.
var (
	ExtensionData_Lock_name = map[int32]string{
		0: "NONE",
		1: "ADMIN",
		2: "DECK",
		3: "TRADE",
		4: "CRAFT",
		5: "OUTSOURCING",
	}
	ExtensionData_Lock_value = map[string]int32{
		"NONE":        0,
		"ADMIN":       1,
		"DECK":        2,
		"TRADE":       3,
		"CRAFT":       4,
		"OUTSOURCING": 5,
	}
)

func (x ExtensionData_Lock) Enum() *ExtensionData_Lock {
	p := new(ExtensionData_Lock)
	*p = x
	return p
}

func (x ExtensionData_Lock) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (ExtensionData_Lock) Descriptor() protoreflect.EnumDescriptor {
	return file_extension_proto_enumTypes[0].Descriptor()
}

func (ExtensionData_Lock) Type() protoreflect.EnumType {
	return &file_extension_proto_enumTypes[0]
}

func (x ExtensionData_Lock) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use ExtensionData_Lock.Descriptor instead.
func (ExtensionData_Lock) EnumDescriptor() ([]byte, []int) {
	return file_extension_proto_rawDescGZIP(), []int{0, 0}
}

type ExtensionData struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	ExtensionId   uint32             `protobuf:"varint,1,opt,name=extension_id,json=extensionId,proto3" json:"extension_id,omitempty"`
	ExtensionType uint32             `protobuf:"varint,2,opt,name=extension_type,json=extensionType,proto3" json:"extension_type,omitempty"`
	Param         *Param             `protobuf:"bytes,3,opt,name=param,proto3" json:"param,omitempty"`
	Trend         *Param             `protobuf:"bytes,4,opt,name=trend,proto3" json:"trend,omitempty"`
	Active        uint32             `protobuf:"varint,5,opt,name=active,proto3" json:"active,omitempty"`
	Lock          ExtensionData_Lock `protobuf:"varint,6,opt,name=lock,proto3,enum=extension.ExtensionData_Lock" json:"lock,omitempty"`
	Title         []string           `protobuf:"bytes,7,rep,name=title,proto3" json:"title,omitempty"`
	LandType      uint32             `protobuf:"varint,8,opt,name=land_type,json=landType,proto3" json:"land_type,omitempty"`
	LockUntil     int64              `protobuf:"varint,9,opt,name=lock_until,json=lockUntil,proto3" json:"lock_until,omitempty"`
}

func (x *ExtensionData) Reset() {
	*x = ExtensionData{}
	if protoimpl.UnsafeEnabled {
		mi := &file_extension_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ExtensionData) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ExtensionData) ProtoMessage() {}

func (x *ExtensionData) ProtoReflect() protoreflect.Message {
	mi := &file_extension_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ExtensionData.ProtoReflect.Descriptor instead.
func (*ExtensionData) Descriptor() ([]byte, []int) {
	return file_extension_proto_rawDescGZIP(), []int{0}
}

func (x *ExtensionData) GetExtensionId() uint32 {
	if x != nil {
		return x.ExtensionId
	}
	return 0
}

func (x *ExtensionData) GetExtensionType() uint32 {
	if x != nil {
		return x.ExtensionType
	}
	return 0
}

func (x *ExtensionData) GetParam() *Param {
	if x != nil {
		return x.Param
	}
	return nil
}

func (x *ExtensionData) GetTrend() *Param {
	if x != nil {
		return x.Trend
	}
	return nil
}

func (x *ExtensionData) GetActive() uint32 {
	if x != nil {
		return x.Active
	}
	return 0
}

func (x *ExtensionData) GetLock() ExtensionData_Lock {
	if x != nil {
		return x.Lock
	}
	return ExtensionData_NONE
}

func (x *ExtensionData) GetTitle() []string {
	if x != nil {
		return x.Title
	}
	return nil
}

func (x *ExtensionData) GetLandType() uint32 {
	if x != nil {
		return x.LandType
	}
	return 0
}

func (x *ExtensionData) GetLockUntil() int64 {
	if x != nil {
		return x.LockUntil
	}
	return 0
}

type Param struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Hp  int32 `protobuf:"varint,1,opt,name=hp,proto3" json:"hp,omitempty"`
	Phy int32 `protobuf:"varint,2,opt,name=phy,proto3" json:"phy,omitempty"`
	Int int32 `protobuf:"varint,3,opt,name=int,proto3" json:"int,omitempty"`
	Agi int32 `protobuf:"varint,4,opt,name=agi,proto3" json:"agi,omitempty"`
}

func (x *Param) Reset() {
	*x = Param{}
	if protoimpl.UnsafeEnabled {
		mi := &file_extension_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Param) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Param) ProtoMessage() {}

func (x *Param) ProtoReflect() protoreflect.Message {
	mi := &file_extension_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Param.ProtoReflect.Descriptor instead.
func (*Param) Descriptor() ([]byte, []int) {
	return file_extension_proto_rawDescGZIP(), []int{1}
}

func (x *Param) GetHp() int32 {
	if x != nil {
		return x.Hp
	}
	return 0
}

func (x *Param) GetPhy() int32 {
	if x != nil {
		return x.Phy
	}
	return 0
}

func (x *Param) GetInt() int32 {
	if x != nil {
		return x.Int
	}
	return 0
}

func (x *Param) GetAgi() int32 {
	if x != nil {
		return x.Agi
	}
	return 0
}

type Empty struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields
}

func (x *Empty) Reset() {
	*x = Empty{}
	if protoimpl.UnsafeEnabled {
		mi := &file_extension_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Empty) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Empty) ProtoMessage() {}

func (x *Empty) ProtoReflect() protoreflect.Message {
	mi := &file_extension_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Empty.ProtoReflect.Descriptor instead.
func (*Empty) Descriptor() ([]byte, []int) {
	return file_extension_proto_rawDescGZIP(), []int{2}
}

var File_extension_proto protoreflect.FileDescriptor

var file_extension_proto_rawDesc = []byte{
	0x0a, 0x0f, 0x65, 0x78, 0x74, 0x65, 0x6e, 0x73, 0x69, 0x6f, 0x6e, 0x2e, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x12, 0x09, 0x65, 0x78, 0x74, 0x65, 0x6e, 0x73, 0x69, 0x6f, 0x6e, 0x22, 0x94, 0x03, 0x0a,
	0x0d, 0x45, 0x78, 0x74, 0x65, 0x6e, 0x73, 0x69, 0x6f, 0x6e, 0x44, 0x61, 0x74, 0x61, 0x12, 0x21,
	0x0a, 0x0c, 0x65, 0x78, 0x74, 0x65, 0x6e, 0x73, 0x69, 0x6f, 0x6e, 0x5f, 0x69, 0x64, 0x18, 0x01,
	0x20, 0x01, 0x28, 0x0d, 0x52, 0x0b, 0x65, 0x78, 0x74, 0x65, 0x6e, 0x73, 0x69, 0x6f, 0x6e, 0x49,
	0x64, 0x12, 0x25, 0x0a, 0x0e, 0x65, 0x78, 0x74, 0x65, 0x6e, 0x73, 0x69, 0x6f, 0x6e, 0x5f, 0x74,
	0x79, 0x70, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0d, 0x52, 0x0d, 0x65, 0x78, 0x74, 0x65, 0x6e,
	0x73, 0x69, 0x6f, 0x6e, 0x54, 0x79, 0x70, 0x65, 0x12, 0x26, 0x0a, 0x05, 0x70, 0x61, 0x72, 0x61,
	0x6d, 0x18, 0x03, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x10, 0x2e, 0x65, 0x78, 0x74, 0x65, 0x6e, 0x73,
	0x69, 0x6f, 0x6e, 0x2e, 0x50, 0x61, 0x72, 0x61, 0x6d, 0x52, 0x05, 0x70, 0x61, 0x72, 0x61, 0x6d,
	0x12, 0x26, 0x0a, 0x05, 0x74, 0x72, 0x65, 0x6e, 0x64, 0x18, 0x04, 0x20, 0x01, 0x28, 0x0b, 0x32,
	0x10, 0x2e, 0x65, 0x78, 0x74, 0x65, 0x6e, 0x73, 0x69, 0x6f, 0x6e, 0x2e, 0x50, 0x61, 0x72, 0x61,
	0x6d, 0x52, 0x05, 0x74, 0x72, 0x65, 0x6e, 0x64, 0x12, 0x16, 0x0a, 0x06, 0x61, 0x63, 0x74, 0x69,
	0x76, 0x65, 0x18, 0x05, 0x20, 0x01, 0x28, 0x0d, 0x52, 0x06, 0x61, 0x63, 0x74, 0x69, 0x76, 0x65,
	0x12, 0x31, 0x0a, 0x04, 0x6c, 0x6f, 0x63, 0x6b, 0x18, 0x06, 0x20, 0x01, 0x28, 0x0e, 0x32, 0x1d,
	0x2e, 0x65, 0x78, 0x74, 0x65, 0x6e, 0x73, 0x69, 0x6f, 0x6e, 0x2e, 0x45, 0x78, 0x74, 0x65, 0x6e,
	0x73, 0x69, 0x6f, 0x6e, 0x44, 0x61, 0x74, 0x61, 0x2e, 0x4c, 0x6f, 0x63, 0x6b, 0x52, 0x04, 0x6c,
	0x6f, 0x63, 0x6b, 0x12, 0x14, 0x0a, 0x05, 0x74, 0x69, 0x74, 0x6c, 0x65, 0x18, 0x07, 0x20, 0x03,
	0x28, 0x09, 0x52, 0x05, 0x74, 0x69, 0x74, 0x6c, 0x65, 0x12, 0x1b, 0x0a, 0x09, 0x6c, 0x61, 0x6e,
	0x64, 0x5f, 0x74, 0x79, 0x70, 0x65, 0x18, 0x08, 0x20, 0x01, 0x28, 0x0d, 0x52, 0x08, 0x6c, 0x61,
	0x6e, 0x64, 0x54, 0x79, 0x70, 0x65, 0x12, 0x1d, 0x0a, 0x0a, 0x6c, 0x6f, 0x63, 0x6b, 0x5f, 0x75,
	0x6e, 0x74, 0x69, 0x6c, 0x18, 0x09, 0x20, 0x01, 0x28, 0x03, 0x52, 0x09, 0x6c, 0x6f, 0x63, 0x6b,
	0x55, 0x6e, 0x74, 0x69, 0x6c, 0x22, 0x4c, 0x0a, 0x04, 0x4c, 0x6f, 0x63, 0x6b, 0x12, 0x08, 0x0a,
	0x04, 0x4e, 0x4f, 0x4e, 0x45, 0x10, 0x00, 0x12, 0x09, 0x0a, 0x05, 0x41, 0x44, 0x4d, 0x49, 0x4e,
	0x10, 0x01, 0x12, 0x08, 0x0a, 0x04, 0x44, 0x45, 0x43, 0x4b, 0x10, 0x02, 0x12, 0x09, 0x0a, 0x05,
	0x54, 0x52, 0x41, 0x44, 0x45, 0x10, 0x03, 0x12, 0x09, 0x0a, 0x05, 0x43, 0x52, 0x41, 0x46, 0x54,
	0x10, 0x04, 0x12, 0x0f, 0x0a, 0x0b, 0x4f, 0x55, 0x54, 0x53, 0x4f, 0x55, 0x52, 0x43, 0x49, 0x4e,
	0x47, 0x10, 0x05, 0x22, 0x4d, 0x0a, 0x05, 0x50, 0x61, 0x72, 0x61, 0x6d, 0x12, 0x0e, 0x0a, 0x02,
	0x68, 0x70, 0x18, 0x01, 0x20, 0x01, 0x28, 0x05, 0x52, 0x02, 0x68, 0x70, 0x12, 0x10, 0x0a, 0x03,
	0x70, 0x68, 0x79, 0x18, 0x02, 0x20, 0x01, 0x28, 0x05, 0x52, 0x03, 0x70, 0x68, 0x79, 0x12, 0x10,
	0x0a, 0x03, 0x69, 0x6e, 0x74, 0x18, 0x03, 0x20, 0x01, 0x28, 0x05, 0x52, 0x03, 0x69, 0x6e, 0x74,
	0x12, 0x10, 0x0a, 0x03, 0x61, 0x67, 0x69, 0x18, 0x04, 0x20, 0x01, 0x28, 0x05, 0x52, 0x03, 0x61,
	0x67, 0x69, 0x22, 0x07, 0x0a, 0x05, 0x45, 0x6d, 0x70, 0x74, 0x79, 0x42, 0x39, 0x5a, 0x37, 0x67,
	0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x68, 0x69, 0x6a, 0x6a, 0x69, 0x72,
	0x69, 0x2f, 0x73, 0x69, 0x6d, 0x75, 0x6c, 0x61, 0x74, 0x6f, 0x72, 0x2f, 0x63, 0x6f, 0x72, 0x65,
	0x2f, 0x67, 0x72, 0x70, 0x63, 0x2d, 0x73, 0x65, 0x72, 0x76, 0x65, 0x72, 0x2f, 0x65, 0x78, 0x74,
	0x65, 0x6e, 0x73, 0x69, 0x6f, 0x6e, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_extension_proto_rawDescOnce sync.Once
	file_extension_proto_rawDescData = file_extension_proto_rawDesc
)

func file_extension_proto_rawDescGZIP() []byte {
	file_extension_proto_rawDescOnce.Do(func() {
		file_extension_proto_rawDescData = protoimpl.X.CompressGZIP(file_extension_proto_rawDescData)
	})
	return file_extension_proto_rawDescData
}

var file_extension_proto_enumTypes = make([]protoimpl.EnumInfo, 1)
var file_extension_proto_msgTypes = make([]protoimpl.MessageInfo, 3)
var file_extension_proto_goTypes = []interface{}{
	(ExtensionData_Lock)(0), // 0: extension.ExtensionData.Lock
	(*ExtensionData)(nil),   // 1: extension.ExtensionData
	(*Param)(nil),           // 2: extension.Param
	(*Empty)(nil),           // 3: extension.Empty
}
var file_extension_proto_depIdxs = []int32{
	2, // 0: extension.ExtensionData.param:type_name -> extension.Param
	2, // 1: extension.ExtensionData.trend:type_name -> extension.Param
	0, // 2: extension.ExtensionData.lock:type_name -> extension.ExtensionData.Lock
	3, // [3:3] is the sub-list for method output_type
	3, // [3:3] is the sub-list for method input_type
	3, // [3:3] is the sub-list for extension type_name
	3, // [3:3] is the sub-list for extension extendee
	0, // [0:3] is the sub-list for field type_name
}

func init() { file_extension_proto_init() }
func file_extension_proto_init() {
	if File_extension_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_extension_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ExtensionData); i {
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
		file_extension_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Param); i {
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
		file_extension_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Empty); i {
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
			RawDescriptor: file_extension_proto_rawDesc,
			NumEnums:      1,
			NumMessages:   3,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_extension_proto_goTypes,
		DependencyIndexes: file_extension_proto_depIdxs,
		EnumInfos:         file_extension_proto_enumTypes,
		MessageInfos:      file_extension_proto_msgTypes,
	}.Build()
	File_extension_proto = out.File
	file_extension_proto_rawDesc = nil
	file_extension_proto_goTypes = nil
	file_extension_proto_depIdxs = nil
}
