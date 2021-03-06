// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.26.0
// 	protoc        v3.19.4
// source: pollutant.proto

package aqi

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

type Pollutant int32

const (
	Pollutant_UNKNOWN   Pollutant = 0
	Pollutant_AQI       Pollutant = 1
	Pollutant_O3_1H     Pollutant = 10
	Pollutant_O3_8H     Pollutant = 11
	Pollutant_PM2_5_1H  Pollutant = 20
	Pollutant_PM2_5_24H Pollutant = 21
	Pollutant_PM10_1H   Pollutant = 30
	Pollutant_PM10_24H  Pollutant = 31
	Pollutant_SO2_1H    Pollutant = 40
	Pollutant_SO2_24H   Pollutant = 41
	Pollutant_NO2_1H    Pollutant = 50
	Pollutant_NO2_24H   Pollutant = 51
	Pollutant_CO_1H     Pollutant = 60
	Pollutant_CO_8H     Pollutant = 61
	Pollutant_CO_24H    Pollutant = 62
)

// Enum value maps for Pollutant.
var (
	Pollutant_name = map[int32]string{
		0:  "UNKNOWN",
		1:  "AQI",
		10: "O3_1H",
		11: "O3_8H",
		20: "PM2_5_1H",
		21: "PM2_5_24H",
		30: "PM10_1H",
		31: "PM10_24H",
		40: "SO2_1H",
		41: "SO2_24H",
		50: "NO2_1H",
		51: "NO2_24H",
		60: "CO_1H",
		61: "CO_8H",
		62: "CO_24H",
	}
	Pollutant_value = map[string]int32{
		"UNKNOWN":   0,
		"AQI":       1,
		"O3_1H":     10,
		"O3_8H":     11,
		"PM2_5_1H":  20,
		"PM2_5_24H": 21,
		"PM10_1H":   30,
		"PM10_24H":  31,
		"SO2_1H":    40,
		"SO2_24H":   41,
		"NO2_1H":    50,
		"NO2_24H":   51,
		"CO_1H":     60,
		"CO_8H":     61,
		"CO_24H":    62,
	}
)

func (x Pollutant) Enum() *Pollutant {
	p := new(Pollutant)
	*p = x
	return p
}

func (x Pollutant) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (Pollutant) Descriptor() protoreflect.EnumDescriptor {
	return file_pollutant_proto_enumTypes[0].Descriptor()
}

func (Pollutant) Type() protoreflect.EnumType {
	return &file_pollutant_proto_enumTypes[0]
}

func (x Pollutant) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use Pollutant.Descriptor instead.
func (Pollutant) EnumDescriptor() ([]byte, []int) {
	return file_pollutant_proto_rawDescGZIP(), []int{0}
}

var File_pollutant_proto protoreflect.FileDescriptor

var file_pollutant_proto_rawDesc = []byte{
	0x0a, 0x0f, 0x70, 0x6f, 0x6c, 0x6c, 0x75, 0x74, 0x61, 0x6e, 0x74, 0x2e, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x12, 0x03, 0x61, 0x71, 0x69, 0x2a, 0xc3, 0x01, 0x0a, 0x09, 0x70, 0x6f, 0x6c, 0x6c, 0x75,
	0x74, 0x61, 0x6e, 0x74, 0x12, 0x0b, 0x0a, 0x07, 0x55, 0x4e, 0x4b, 0x4e, 0x4f, 0x57, 0x4e, 0x10,
	0x00, 0x12, 0x07, 0x0a, 0x03, 0x41, 0x51, 0x49, 0x10, 0x01, 0x12, 0x09, 0x0a, 0x05, 0x4f, 0x33,
	0x5f, 0x31, 0x48, 0x10, 0x0a, 0x12, 0x09, 0x0a, 0x05, 0x4f, 0x33, 0x5f, 0x38, 0x48, 0x10, 0x0b,
	0x12, 0x0c, 0x0a, 0x08, 0x50, 0x4d, 0x32, 0x5f, 0x35, 0x5f, 0x31, 0x48, 0x10, 0x14, 0x12, 0x0d,
	0x0a, 0x09, 0x50, 0x4d, 0x32, 0x5f, 0x35, 0x5f, 0x32, 0x34, 0x48, 0x10, 0x15, 0x12, 0x0b, 0x0a,
	0x07, 0x50, 0x4d, 0x31, 0x30, 0x5f, 0x31, 0x48, 0x10, 0x1e, 0x12, 0x0c, 0x0a, 0x08, 0x50, 0x4d,
	0x31, 0x30, 0x5f, 0x32, 0x34, 0x48, 0x10, 0x1f, 0x12, 0x0a, 0x0a, 0x06, 0x53, 0x4f, 0x32, 0x5f,
	0x31, 0x48, 0x10, 0x28, 0x12, 0x0b, 0x0a, 0x07, 0x53, 0x4f, 0x32, 0x5f, 0x32, 0x34, 0x48, 0x10,
	0x29, 0x12, 0x0a, 0x0a, 0x06, 0x4e, 0x4f, 0x32, 0x5f, 0x31, 0x48, 0x10, 0x32, 0x12, 0x0b, 0x0a,
	0x07, 0x4e, 0x4f, 0x32, 0x5f, 0x32, 0x34, 0x48, 0x10, 0x33, 0x12, 0x09, 0x0a, 0x05, 0x43, 0x4f,
	0x5f, 0x31, 0x48, 0x10, 0x3c, 0x12, 0x09, 0x0a, 0x05, 0x43, 0x4f, 0x5f, 0x38, 0x48, 0x10, 0x3d,
	0x12, 0x0a, 0x0a, 0x06, 0x43, 0x4f, 0x5f, 0x32, 0x34, 0x48, 0x10, 0x3e, 0x42, 0x08, 0x5a, 0x06,
	0x2e, 0x2f, 0x3b, 0x61, 0x71, 0x69, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_pollutant_proto_rawDescOnce sync.Once
	file_pollutant_proto_rawDescData = file_pollutant_proto_rawDesc
)

func file_pollutant_proto_rawDescGZIP() []byte {
	file_pollutant_proto_rawDescOnce.Do(func() {
		file_pollutant_proto_rawDescData = protoimpl.X.CompressGZIP(file_pollutant_proto_rawDescData)
	})
	return file_pollutant_proto_rawDescData
}

var file_pollutant_proto_enumTypes = make([]protoimpl.EnumInfo, 1)
var file_pollutant_proto_goTypes = []interface{}{
	(Pollutant)(0), // 0: aqi.pollutant
}
var file_pollutant_proto_depIdxs = []int32{
	0, // [0:0] is the sub-list for method output_type
	0, // [0:0] is the sub-list for method input_type
	0, // [0:0] is the sub-list for extension type_name
	0, // [0:0] is the sub-list for extension extendee
	0, // [0:0] is the sub-list for field type_name
}

func init() { file_pollutant_proto_init() }
func file_pollutant_proto_init() {
	if File_pollutant_proto != nil {
		return
	}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_pollutant_proto_rawDesc,
			NumEnums:      1,
			NumMessages:   0,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_pollutant_proto_goTypes,
		DependencyIndexes: file_pollutant_proto_depIdxs,
		EnumInfos:         file_pollutant_proto_enumTypes,
	}.Build()
	File_pollutant_proto = out.File
	file_pollutant_proto_rawDesc = nil
	file_pollutant_proto_goTypes = nil
	file_pollutant_proto_depIdxs = nil
}
