// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.31.0
// 	protoc        v4.23.4
// source: hero.proto

package hero

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

type HeroTypeMaster struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	HeroType       uint32   `protobuf:"varint,1,opt,name=hero_type,json=heroType,proto3" json:"hero_type,omitempty"`
	Asset          bool     `protobuf:"varint,2,opt,name=asset,proto3" json:"asset,omitempty"`
	Rarity         int32    `protobuf:"varint,3,opt,name=rarity,proto3" json:"rarity,omitempty"`
	Supply         int32    `protobuf:"varint,4,opt,name=supply,proto3" json:"supply,omitempty"`
	Param          *Param   `protobuf:"bytes,5,opt,name=param,proto3" json:"param,omitempty"`
	Trend          *Param   `protobuf:"bytes,6,opt,name=trend,proto3" json:"trend,omitempty"`
	Active         uint32   `protobuf:"varint,7,opt,name=active,proto3" json:"active,omitempty"` // deprecated
	Passive        uint32   `protobuf:"varint,8,opt,name=passive,proto3" json:"passive,omitempty"`
	Name           string   `protobuf:"bytes,9,opt,name=name,proto3" json:"name,omitempty"`
	Attribute      []string `protobuf:"bytes,16,rep,name=attribute,proto3" json:"attribute,omitempty"` // deprecated
	AttributeTypes []uint32 `protobuf:"varint,17,rep,packed,name=attribute_types,json=attributeTypes,proto3" json:"attribute_types,omitempty"`
	PassiveV2      []uint32 `protobuf:"varint,18,rep,packed,name=passive_v2,json=passiveV2,proto3" json:"passive_v2,omitempty"`
}

func (x *HeroTypeMaster) Reset() {
	*x = HeroTypeMaster{}
	if protoimpl.UnsafeEnabled {
		mi := &file_hero_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *HeroTypeMaster) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*HeroTypeMaster) ProtoMessage() {}

func (x *HeroTypeMaster) ProtoReflect() protoreflect.Message {
	mi := &file_hero_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use HeroTypeMaster.ProtoReflect.Descriptor instead.
func (*HeroTypeMaster) Descriptor() ([]byte, []int) {
	return file_hero_proto_rawDescGZIP(), []int{0}
}

func (x *HeroTypeMaster) GetHeroType() uint32 {
	if x != nil {
		return x.HeroType
	}
	return 0
}

func (x *HeroTypeMaster) GetAsset() bool {
	if x != nil {
		return x.Asset
	}
	return false
}

func (x *HeroTypeMaster) GetRarity() int32 {
	if x != nil {
		return x.Rarity
	}
	return 0
}

func (x *HeroTypeMaster) GetSupply() int32 {
	if x != nil {
		return x.Supply
	}
	return 0
}

func (x *HeroTypeMaster) GetParam() *Param {
	if x != nil {
		return x.Param
	}
	return nil
}

func (x *HeroTypeMaster) GetTrend() *Param {
	if x != nil {
		return x.Trend
	}
	return nil
}

func (x *HeroTypeMaster) GetActive() uint32 {
	if x != nil {
		return x.Active
	}
	return 0
}

func (x *HeroTypeMaster) GetPassive() uint32 {
	if x != nil {
		return x.Passive
	}
	return 0
}

func (x *HeroTypeMaster) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

func (x *HeroTypeMaster) GetAttribute() []string {
	if x != nil {
		return x.Attribute
	}
	return nil
}

func (x *HeroTypeMaster) GetAttributeTypes() []uint32 {
	if x != nil {
		return x.AttributeTypes
	}
	return nil
}

func (x *HeroTypeMaster) GetPassiveV2() []uint32 {
	if x != nil {
		return x.PassiveV2
	}
	return nil
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
		mi := &file_hero_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Param) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Param) ProtoMessage() {}

func (x *Param) ProtoReflect() protoreflect.Message {
	mi := &file_hero_proto_msgTypes[1]
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
	return file_hero_proto_rawDescGZIP(), []int{1}
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
		mi := &file_hero_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Empty) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Empty) ProtoMessage() {}

func (x *Empty) ProtoReflect() protoreflect.Message {
	mi := &file_hero_proto_msgTypes[2]
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
	return file_hero_proto_rawDescGZIP(), []int{2}
}

var File_hero_proto protoreflect.FileDescriptor

var file_hero_proto_rawDesc = []byte{
	0x0a, 0x0a, 0x68, 0x65, 0x72, 0x6f, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x04, 0x68, 0x65,
	0x72, 0x6f, 0x22, 0xe5, 0x02, 0x0a, 0x0e, 0x48, 0x65, 0x72, 0x6f, 0x54, 0x79, 0x70, 0x65, 0x4d,
	0x61, 0x73, 0x74, 0x65, 0x72, 0x12, 0x1b, 0x0a, 0x09, 0x68, 0x65, 0x72, 0x6f, 0x5f, 0x74, 0x79,
	0x70, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0d, 0x52, 0x08, 0x68, 0x65, 0x72, 0x6f, 0x54, 0x79,
	0x70, 0x65, 0x12, 0x14, 0x0a, 0x05, 0x61, 0x73, 0x73, 0x65, 0x74, 0x18, 0x02, 0x20, 0x01, 0x28,
	0x08, 0x52, 0x05, 0x61, 0x73, 0x73, 0x65, 0x74, 0x12, 0x16, 0x0a, 0x06, 0x72, 0x61, 0x72, 0x69,
	0x74, 0x79, 0x18, 0x03, 0x20, 0x01, 0x28, 0x05, 0x52, 0x06, 0x72, 0x61, 0x72, 0x69, 0x74, 0x79,
	0x12, 0x16, 0x0a, 0x06, 0x73, 0x75, 0x70, 0x70, 0x6c, 0x79, 0x18, 0x04, 0x20, 0x01, 0x28, 0x05,
	0x52, 0x06, 0x73, 0x75, 0x70, 0x70, 0x6c, 0x79, 0x12, 0x21, 0x0a, 0x05, 0x70, 0x61, 0x72, 0x61,
	0x6d, 0x18, 0x05, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x0b, 0x2e, 0x68, 0x65, 0x72, 0x6f, 0x2e, 0x50,
	0x61, 0x72, 0x61, 0x6d, 0x52, 0x05, 0x70, 0x61, 0x72, 0x61, 0x6d, 0x12, 0x21, 0x0a, 0x05, 0x74,
	0x72, 0x65, 0x6e, 0x64, 0x18, 0x06, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x0b, 0x2e, 0x68, 0x65, 0x72,
	0x6f, 0x2e, 0x50, 0x61, 0x72, 0x61, 0x6d, 0x52, 0x05, 0x74, 0x72, 0x65, 0x6e, 0x64, 0x12, 0x16,
	0x0a, 0x06, 0x61, 0x63, 0x74, 0x69, 0x76, 0x65, 0x18, 0x07, 0x20, 0x01, 0x28, 0x0d, 0x52, 0x06,
	0x61, 0x63, 0x74, 0x69, 0x76, 0x65, 0x12, 0x18, 0x0a, 0x07, 0x70, 0x61, 0x73, 0x73, 0x69, 0x76,
	0x65, 0x18, 0x08, 0x20, 0x01, 0x28, 0x0d, 0x52, 0x07, 0x70, 0x61, 0x73, 0x73, 0x69, 0x76, 0x65,
	0x12, 0x12, 0x0a, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x09, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04,
	0x6e, 0x61, 0x6d, 0x65, 0x12, 0x1c, 0x0a, 0x09, 0x61, 0x74, 0x74, 0x72, 0x69, 0x62, 0x75, 0x74,
	0x65, 0x18, 0x10, 0x20, 0x03, 0x28, 0x09, 0x52, 0x09, 0x61, 0x74, 0x74, 0x72, 0x69, 0x62, 0x75,
	0x74, 0x65, 0x12, 0x27, 0x0a, 0x0f, 0x61, 0x74, 0x74, 0x72, 0x69, 0x62, 0x75, 0x74, 0x65, 0x5f,
	0x74, 0x79, 0x70, 0x65, 0x73, 0x18, 0x11, 0x20, 0x03, 0x28, 0x0d, 0x52, 0x0e, 0x61, 0x74, 0x74,
	0x72, 0x69, 0x62, 0x75, 0x74, 0x65, 0x54, 0x79, 0x70, 0x65, 0x73, 0x12, 0x1d, 0x0a, 0x0a, 0x70,
	0x61, 0x73, 0x73, 0x69, 0x76, 0x65, 0x5f, 0x76, 0x32, 0x18, 0x12, 0x20, 0x03, 0x28, 0x0d, 0x52,
	0x09, 0x70, 0x61, 0x73, 0x73, 0x69, 0x76, 0x65, 0x56, 0x32, 0x22, 0x4d, 0x0a, 0x05, 0x50, 0x61,
	0x72, 0x61, 0x6d, 0x12, 0x0e, 0x0a, 0x02, 0x68, 0x70, 0x18, 0x01, 0x20, 0x01, 0x28, 0x05, 0x52,
	0x02, 0x68, 0x70, 0x12, 0x10, 0x0a, 0x03, 0x70, 0x68, 0x79, 0x18, 0x02, 0x20, 0x01, 0x28, 0x05,
	0x52, 0x03, 0x70, 0x68, 0x79, 0x12, 0x10, 0x0a, 0x03, 0x69, 0x6e, 0x74, 0x18, 0x03, 0x20, 0x01,
	0x28, 0x05, 0x52, 0x03, 0x69, 0x6e, 0x74, 0x12, 0x10, 0x0a, 0x03, 0x61, 0x67, 0x69, 0x18, 0x04,
	0x20, 0x01, 0x28, 0x05, 0x52, 0x03, 0x61, 0x67, 0x69, 0x22, 0x07, 0x0a, 0x05, 0x45, 0x6d, 0x70,
	0x74, 0x79, 0x32, 0x48, 0x0a, 0x0b, 0x48, 0x65, 0x72, 0x6f, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63,
	0x65, 0x12, 0x39, 0x0a, 0x12, 0x53, 0x61, 0x76, 0x65, 0x49, 0x70, 0x66, 0x73, 0x53, 0x6b, 0x69,
	0x6c, 0x6c, 0x53, 0x74, 0x6f, 0x63, 0x6b, 0x12, 0x0b, 0x2e, 0x68, 0x65, 0x72, 0x6f, 0x2e, 0x45,
	0x6d, 0x70, 0x74, 0x79, 0x1a, 0x14, 0x2e, 0x68, 0x65, 0x72, 0x6f, 0x2e, 0x48, 0x65, 0x72, 0x6f,
	0x54, 0x79, 0x70, 0x65, 0x4d, 0x61, 0x73, 0x74, 0x65, 0x72, 0x22, 0x00, 0x42, 0x1e, 0x5a, 0x1c,
	0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x68, 0x69, 0x6a, 0x6a, 0x69,
	0x72, 0x69, 0x2f, 0x63, 0x6f, 0x72, 0x65, 0x2f, 0x68, 0x65, 0x72, 0x6f, 0x62, 0x06, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_hero_proto_rawDescOnce sync.Once
	file_hero_proto_rawDescData = file_hero_proto_rawDesc
)

func file_hero_proto_rawDescGZIP() []byte {
	file_hero_proto_rawDescOnce.Do(func() {
		file_hero_proto_rawDescData = protoimpl.X.CompressGZIP(file_hero_proto_rawDescData)
	})
	return file_hero_proto_rawDescData
}

var file_hero_proto_msgTypes = make([]protoimpl.MessageInfo, 3)
var file_hero_proto_goTypes = []interface{}{
	(*HeroTypeMaster)(nil), // 0: hero.HeroTypeMaster
	(*Param)(nil),          // 1: hero.Param
	(*Empty)(nil),          // 2: hero.Empty
}
var file_hero_proto_depIdxs = []int32{
	1, // 0: hero.HeroTypeMaster.param:type_name -> hero.Param
	1, // 1: hero.HeroTypeMaster.trend:type_name -> hero.Param
	2, // 2: hero.HeroService.SaveIpfsSkillStock:input_type -> hero.Empty
	0, // 3: hero.HeroService.SaveIpfsSkillStock:output_type -> hero.HeroTypeMaster
	3, // [3:4] is the sub-list for method output_type
	2, // [2:3] is the sub-list for method input_type
	2, // [2:2] is the sub-list for extension type_name
	2, // [2:2] is the sub-list for extension extendee
	0, // [0:2] is the sub-list for field type_name
}

func init() { file_hero_proto_init() }
func file_hero_proto_init() {
	if File_hero_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_hero_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*HeroTypeMaster); i {
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
		file_hero_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
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
		file_hero_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
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
			RawDescriptor: file_hero_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   3,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_hero_proto_goTypes,
		DependencyIndexes: file_hero_proto_depIdxs,
		MessageInfos:      file_hero_proto_msgTypes,
	}.Build()
	File_hero_proto = out.File
	file_hero_proto_rawDesc = nil
	file_hero_proto_goTypes = nil
	file_hero_proto_depIdxs = nil
}
