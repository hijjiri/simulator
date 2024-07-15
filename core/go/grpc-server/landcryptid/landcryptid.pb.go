// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.32.0-devel
// 	protoc        v3.12.4
// source: landcryptid.proto

package landcryptid

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

type CryptidInfo struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	LandType uint32        `protobuf:"varint,1,opt,name=land_type,json=landType,proto3" json:"land_type,omitempty"`
	Energy   int32         `protobuf:"varint,2,opt,name=energy,proto3" json:"energy,omitempty"`
	Lv       int32         `protobuf:"varint,3,opt,name=lv,proto3" json:"lv,omitempty"`
	Season   int32         `protobuf:"varint,4,opt,name=season,proto3" json:"season,omitempty"`
	LaboCe   int32         `protobuf:"varint,5,opt,name=labo_ce,json=laboCe,proto3" json:"labo_ce,omitempty"`
	Hp       *CryptidParam `protobuf:"bytes,6,opt,name=hp,proto3" json:"hp,omitempty"`
	Phy      *CryptidParam `protobuf:"bytes,7,opt,name=phy,proto3" json:"phy,omitempty"`
	Int      *CryptidParam `protobuf:"bytes,8,opt,name=int,proto3" json:"int,omitempty"`
	Agi      *CryptidParam `protobuf:"bytes,9,opt,name=agi,proto3" json:"agi,omitempty"`
	Recovery *CryptidParam `protobuf:"bytes,10,opt,name=recovery,proto3" json:"recovery,omitempty"`
}

func (x *CryptidInfo) Reset() {
	*x = CryptidInfo{}
	if protoimpl.UnsafeEnabled {
		mi := &file_landcryptid_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *CryptidInfo) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CryptidInfo) ProtoMessage() {}

func (x *CryptidInfo) ProtoReflect() protoreflect.Message {
	mi := &file_landcryptid_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CryptidInfo.ProtoReflect.Descriptor instead.
func (*CryptidInfo) Descriptor() ([]byte, []int) {
	return file_landcryptid_proto_rawDescGZIP(), []int{0}
}

func (x *CryptidInfo) GetLandType() uint32 {
	if x != nil {
		return x.LandType
	}
	return 0
}

func (x *CryptidInfo) GetEnergy() int32 {
	if x != nil {
		return x.Energy
	}
	return 0
}

func (x *CryptidInfo) GetLv() int32 {
	if x != nil {
		return x.Lv
	}
	return 0
}

func (x *CryptidInfo) GetSeason() int32 {
	if x != nil {
		return x.Season
	}
	return 0
}

func (x *CryptidInfo) GetLaboCe() int32 {
	if x != nil {
		return x.LaboCe
	}
	return 0
}

func (x *CryptidInfo) GetHp() *CryptidParam {
	if x != nil {
		return x.Hp
	}
	return nil
}

func (x *CryptidInfo) GetPhy() *CryptidParam {
	if x != nil {
		return x.Phy
	}
	return nil
}

func (x *CryptidInfo) GetInt() *CryptidParam {
	if x != nil {
		return x.Int
	}
	return nil
}

func (x *CryptidInfo) GetAgi() *CryptidParam {
	if x != nil {
		return x.Agi
	}
	return nil
}

func (x *CryptidInfo) GetRecovery() *CryptidParam {
	if x != nil {
		return x.Recovery
	}
	return nil
}

type CryptidParam struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	DivideRate100 int32 `protobuf:"varint,1,opt,name=divide_rate100,json=divideRate100,proto3" json:"divide_rate100,omitempty"`
	Ce            int32 `protobuf:"varint,2,opt,name=ce,proto3" json:"ce,omitempty"`
	Lv            int32 `protobuf:"varint,3,opt,name=lv,proto3" json:"lv,omitempty"`
	BuffRate1000  int32 `protobuf:"varint,4,opt,name=buff_rate1000,json=buffRate1000,proto3" json:"buff_rate1000,omitempty"`
}

func (x *CryptidParam) Reset() {
	*x = CryptidParam{}
	if protoimpl.UnsafeEnabled {
		mi := &file_landcryptid_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *CryptidParam) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CryptidParam) ProtoMessage() {}

func (x *CryptidParam) ProtoReflect() protoreflect.Message {
	mi := &file_landcryptid_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CryptidParam.ProtoReflect.Descriptor instead.
func (*CryptidParam) Descriptor() ([]byte, []int) {
	return file_landcryptid_proto_rawDescGZIP(), []int{1}
}

func (x *CryptidParam) GetDivideRate100() int32 {
	if x != nil {
		return x.DivideRate100
	}
	return 0
}

func (x *CryptidParam) GetCe() int32 {
	if x != nil {
		return x.Ce
	}
	return 0
}

func (x *CryptidParam) GetLv() int32 {
	if x != nil {
		return x.Lv
	}
	return 0
}

func (x *CryptidParam) GetBuffRate1000() int32 {
	if x != nil {
		return x.BuffRate1000
	}
	return 0
}

var File_landcryptid_proto protoreflect.FileDescriptor

var file_landcryptid_proto_rawDesc = []byte{
	0x0a, 0x11, 0x6c, 0x61, 0x6e, 0x64, 0x63, 0x72, 0x79, 0x70, 0x74, 0x69, 0x64, 0x2e, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x12, 0x0b, 0x6c, 0x61, 0x6e, 0x64, 0x63, 0x72, 0x79, 0x70, 0x74, 0x69, 0x64,
	0x22, 0xec, 0x02, 0x0a, 0x0b, 0x43, 0x72, 0x79, 0x70, 0x74, 0x69, 0x64, 0x49, 0x6e, 0x66, 0x6f,
	0x12, 0x1b, 0x0a, 0x09, 0x6c, 0x61, 0x6e, 0x64, 0x5f, 0x74, 0x79, 0x70, 0x65, 0x18, 0x01, 0x20,
	0x01, 0x28, 0x0d, 0x52, 0x08, 0x6c, 0x61, 0x6e, 0x64, 0x54, 0x79, 0x70, 0x65, 0x12, 0x16, 0x0a,
	0x06, 0x65, 0x6e, 0x65, 0x72, 0x67, 0x79, 0x18, 0x02, 0x20, 0x01, 0x28, 0x05, 0x52, 0x06, 0x65,
	0x6e, 0x65, 0x72, 0x67, 0x79, 0x12, 0x0e, 0x0a, 0x02, 0x6c, 0x76, 0x18, 0x03, 0x20, 0x01, 0x28,
	0x05, 0x52, 0x02, 0x6c, 0x76, 0x12, 0x16, 0x0a, 0x06, 0x73, 0x65, 0x61, 0x73, 0x6f, 0x6e, 0x18,
	0x04, 0x20, 0x01, 0x28, 0x05, 0x52, 0x06, 0x73, 0x65, 0x61, 0x73, 0x6f, 0x6e, 0x12, 0x17, 0x0a,
	0x07, 0x6c, 0x61, 0x62, 0x6f, 0x5f, 0x63, 0x65, 0x18, 0x05, 0x20, 0x01, 0x28, 0x05, 0x52, 0x06,
	0x6c, 0x61, 0x62, 0x6f, 0x43, 0x65, 0x12, 0x29, 0x0a, 0x02, 0x68, 0x70, 0x18, 0x06, 0x20, 0x01,
	0x28, 0x0b, 0x32, 0x19, 0x2e, 0x6c, 0x61, 0x6e, 0x64, 0x63, 0x72, 0x79, 0x70, 0x74, 0x69, 0x64,
	0x2e, 0x43, 0x72, 0x79, 0x70, 0x74, 0x69, 0x64, 0x50, 0x61, 0x72, 0x61, 0x6d, 0x52, 0x02, 0x68,
	0x70, 0x12, 0x2b, 0x0a, 0x03, 0x70, 0x68, 0x79, 0x18, 0x07, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x19,
	0x2e, 0x6c, 0x61, 0x6e, 0x64, 0x63, 0x72, 0x79, 0x70, 0x74, 0x69, 0x64, 0x2e, 0x43, 0x72, 0x79,
	0x70, 0x74, 0x69, 0x64, 0x50, 0x61, 0x72, 0x61, 0x6d, 0x52, 0x03, 0x70, 0x68, 0x79, 0x12, 0x2b,
	0x0a, 0x03, 0x69, 0x6e, 0x74, 0x18, 0x08, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x19, 0x2e, 0x6c, 0x61,
	0x6e, 0x64, 0x63, 0x72, 0x79, 0x70, 0x74, 0x69, 0x64, 0x2e, 0x43, 0x72, 0x79, 0x70, 0x74, 0x69,
	0x64, 0x50, 0x61, 0x72, 0x61, 0x6d, 0x52, 0x03, 0x69, 0x6e, 0x74, 0x12, 0x2b, 0x0a, 0x03, 0x61,
	0x67, 0x69, 0x18, 0x09, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x19, 0x2e, 0x6c, 0x61, 0x6e, 0x64, 0x63,
	0x72, 0x79, 0x70, 0x74, 0x69, 0x64, 0x2e, 0x43, 0x72, 0x79, 0x70, 0x74, 0x69, 0x64, 0x50, 0x61,
	0x72, 0x61, 0x6d, 0x52, 0x03, 0x61, 0x67, 0x69, 0x12, 0x35, 0x0a, 0x08, 0x72, 0x65, 0x63, 0x6f,
	0x76, 0x65, 0x72, 0x79, 0x18, 0x0a, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x19, 0x2e, 0x6c, 0x61, 0x6e,
	0x64, 0x63, 0x72, 0x79, 0x70, 0x74, 0x69, 0x64, 0x2e, 0x43, 0x72, 0x79, 0x70, 0x74, 0x69, 0x64,
	0x50, 0x61, 0x72, 0x61, 0x6d, 0x52, 0x08, 0x72, 0x65, 0x63, 0x6f, 0x76, 0x65, 0x72, 0x79, 0x22,
	0x7a, 0x0a, 0x0c, 0x43, 0x72, 0x79, 0x70, 0x74, 0x69, 0x64, 0x50, 0x61, 0x72, 0x61, 0x6d, 0x12,
	0x25, 0x0a, 0x0e, 0x64, 0x69, 0x76, 0x69, 0x64, 0x65, 0x5f, 0x72, 0x61, 0x74, 0x65, 0x31, 0x30,
	0x30, 0x18, 0x01, 0x20, 0x01, 0x28, 0x05, 0x52, 0x0d, 0x64, 0x69, 0x76, 0x69, 0x64, 0x65, 0x52,
	0x61, 0x74, 0x65, 0x31, 0x30, 0x30, 0x12, 0x0e, 0x0a, 0x02, 0x63, 0x65, 0x18, 0x02, 0x20, 0x01,
	0x28, 0x05, 0x52, 0x02, 0x63, 0x65, 0x12, 0x0e, 0x0a, 0x02, 0x6c, 0x76, 0x18, 0x03, 0x20, 0x01,
	0x28, 0x05, 0x52, 0x02, 0x6c, 0x76, 0x12, 0x23, 0x0a, 0x0d, 0x62, 0x75, 0x66, 0x66, 0x5f, 0x72,
	0x61, 0x74, 0x65, 0x31, 0x30, 0x30, 0x30, 0x18, 0x04, 0x20, 0x01, 0x28, 0x05, 0x52, 0x0c, 0x62,
	0x75, 0x66, 0x66, 0x52, 0x61, 0x74, 0x65, 0x31, 0x30, 0x30, 0x30, 0x42, 0x3b, 0x5a, 0x39, 0x67,
	0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x68, 0x69, 0x6a, 0x6a, 0x69, 0x72,
	0x69, 0x2f, 0x73, 0x69, 0x6d, 0x75, 0x6c, 0x61, 0x74, 0x6f, 0x72, 0x2f, 0x63, 0x6f, 0x72, 0x65,
	0x2f, 0x67, 0x72, 0x70, 0x63, 0x2d, 0x73, 0x65, 0x72, 0x76, 0x65, 0x72, 0x2f, 0x6c, 0x61, 0x6e,
	0x64, 0x63, 0x72, 0x79, 0x70, 0x74, 0x69, 0x64, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_landcryptid_proto_rawDescOnce sync.Once
	file_landcryptid_proto_rawDescData = file_landcryptid_proto_rawDesc
)

func file_landcryptid_proto_rawDescGZIP() []byte {
	file_landcryptid_proto_rawDescOnce.Do(func() {
		file_landcryptid_proto_rawDescData = protoimpl.X.CompressGZIP(file_landcryptid_proto_rawDescData)
	})
	return file_landcryptid_proto_rawDescData
}

var file_landcryptid_proto_msgTypes = make([]protoimpl.MessageInfo, 2)
var file_landcryptid_proto_goTypes = []interface{}{
	(*CryptidInfo)(nil),  // 0: landcryptid.CryptidInfo
	(*CryptidParam)(nil), // 1: landcryptid.CryptidParam
}
var file_landcryptid_proto_depIdxs = []int32{
	1, // 0: landcryptid.CryptidInfo.hp:type_name -> landcryptid.CryptidParam
	1, // 1: landcryptid.CryptidInfo.phy:type_name -> landcryptid.CryptidParam
	1, // 2: landcryptid.CryptidInfo.int:type_name -> landcryptid.CryptidParam
	1, // 3: landcryptid.CryptidInfo.agi:type_name -> landcryptid.CryptidParam
	1, // 4: landcryptid.CryptidInfo.recovery:type_name -> landcryptid.CryptidParam
	5, // [5:5] is the sub-list for method output_type
	5, // [5:5] is the sub-list for method input_type
	5, // [5:5] is the sub-list for extension type_name
	5, // [5:5] is the sub-list for extension extendee
	0, // [0:5] is the sub-list for field type_name
}

func init() { file_landcryptid_proto_init() }
func file_landcryptid_proto_init() {
	if File_landcryptid_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_landcryptid_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*CryptidInfo); i {
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
		file_landcryptid_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*CryptidParam); i {
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
			RawDescriptor: file_landcryptid_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   2,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_landcryptid_proto_goTypes,
		DependencyIndexes: file_landcryptid_proto_depIdxs,
		MessageInfos:      file_landcryptid_proto_msgTypes,
	}.Build()
	File_landcryptid_proto = out.File
	file_landcryptid_proto_rawDesc = nil
	file_landcryptid_proto_goTypes = nil
	file_landcryptid_proto_depIdxs = nil
}
