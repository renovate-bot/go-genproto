// Copyright 2025 Google LLC
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.26.0
// 	protoc        v4.24.4
// source: google/type/month.proto

package month

import (
	reflect "reflect"
	sync "sync"

	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
)

const (
	// Verify that this generated code is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(20 - protoimpl.MinVersion)
	// Verify that runtime/protoimpl is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(protoimpl.MaxVersion - 20)
)

// Represents a month in the Gregorian calendar.
type Month int32

const (
	// The unspecified month.
	Month_MONTH_UNSPECIFIED Month = 0
	// The month of January.
	Month_JANUARY Month = 1
	// The month of February.
	Month_FEBRUARY Month = 2
	// The month of March.
	Month_MARCH Month = 3
	// The month of April.
	Month_APRIL Month = 4
	// The month of May.
	Month_MAY Month = 5
	// The month of June.
	Month_JUNE Month = 6
	// The month of July.
	Month_JULY Month = 7
	// The month of August.
	Month_AUGUST Month = 8
	// The month of September.
	Month_SEPTEMBER Month = 9
	// The month of October.
	Month_OCTOBER Month = 10
	// The month of November.
	Month_NOVEMBER Month = 11
	// The month of December.
	Month_DECEMBER Month = 12
)

// Enum value maps for Month.
var (
	Month_name = map[int32]string{
		0:  "MONTH_UNSPECIFIED",
		1:  "JANUARY",
		2:  "FEBRUARY",
		3:  "MARCH",
		4:  "APRIL",
		5:  "MAY",
		6:  "JUNE",
		7:  "JULY",
		8:  "AUGUST",
		9:  "SEPTEMBER",
		10: "OCTOBER",
		11: "NOVEMBER",
		12: "DECEMBER",
	}
	Month_value = map[string]int32{
		"MONTH_UNSPECIFIED": 0,
		"JANUARY":           1,
		"FEBRUARY":          2,
		"MARCH":             3,
		"APRIL":             4,
		"MAY":               5,
		"JUNE":              6,
		"JULY":              7,
		"AUGUST":            8,
		"SEPTEMBER":         9,
		"OCTOBER":           10,
		"NOVEMBER":          11,
		"DECEMBER":          12,
	}
)

func (x Month) Enum() *Month {
	p := new(Month)
	*p = x
	return p
}

func (x Month) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (Month) Descriptor() protoreflect.EnumDescriptor {
	return file_google_type_month_proto_enumTypes[0].Descriptor()
}

func (Month) Type() protoreflect.EnumType {
	return &file_google_type_month_proto_enumTypes[0]
}

func (x Month) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use Month.Descriptor instead.
func (Month) EnumDescriptor() ([]byte, []int) {
	return file_google_type_month_proto_rawDescGZIP(), []int{0}
}

var File_google_type_month_proto protoreflect.FileDescriptor

var file_google_type_month_proto_rawDesc = []byte{
	0x0a, 0x17, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x74, 0x79, 0x70, 0x65, 0x2f, 0x6d, 0x6f,
	0x6e, 0x74, 0x68, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x0b, 0x67, 0x6f, 0x6f, 0x67, 0x6c,
	0x65, 0x2e, 0x74, 0x79, 0x70, 0x65, 0x2a, 0xb0, 0x01, 0x0a, 0x05, 0x4d, 0x6f, 0x6e, 0x74, 0x68,
	0x12, 0x15, 0x0a, 0x11, 0x4d, 0x4f, 0x4e, 0x54, 0x48, 0x5f, 0x55, 0x4e, 0x53, 0x50, 0x45, 0x43,
	0x49, 0x46, 0x49, 0x45, 0x44, 0x10, 0x00, 0x12, 0x0b, 0x0a, 0x07, 0x4a, 0x41, 0x4e, 0x55, 0x41,
	0x52, 0x59, 0x10, 0x01, 0x12, 0x0c, 0x0a, 0x08, 0x46, 0x45, 0x42, 0x52, 0x55, 0x41, 0x52, 0x59,
	0x10, 0x02, 0x12, 0x09, 0x0a, 0x05, 0x4d, 0x41, 0x52, 0x43, 0x48, 0x10, 0x03, 0x12, 0x09, 0x0a,
	0x05, 0x41, 0x50, 0x52, 0x49, 0x4c, 0x10, 0x04, 0x12, 0x07, 0x0a, 0x03, 0x4d, 0x41, 0x59, 0x10,
	0x05, 0x12, 0x08, 0x0a, 0x04, 0x4a, 0x55, 0x4e, 0x45, 0x10, 0x06, 0x12, 0x08, 0x0a, 0x04, 0x4a,
	0x55, 0x4c, 0x59, 0x10, 0x07, 0x12, 0x0a, 0x0a, 0x06, 0x41, 0x55, 0x47, 0x55, 0x53, 0x54, 0x10,
	0x08, 0x12, 0x0d, 0x0a, 0x09, 0x53, 0x45, 0x50, 0x54, 0x45, 0x4d, 0x42, 0x45, 0x52, 0x10, 0x09,
	0x12, 0x0b, 0x0a, 0x07, 0x4f, 0x43, 0x54, 0x4f, 0x42, 0x45, 0x52, 0x10, 0x0a, 0x12, 0x0c, 0x0a,
	0x08, 0x4e, 0x4f, 0x56, 0x45, 0x4d, 0x42, 0x45, 0x52, 0x10, 0x0b, 0x12, 0x0c, 0x0a, 0x08, 0x44,
	0x45, 0x43, 0x45, 0x4d, 0x42, 0x45, 0x52, 0x10, 0x0c, 0x42, 0x5d, 0x0a, 0x0f, 0x63, 0x6f, 0x6d,
	0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x74, 0x79, 0x70, 0x65, 0x42, 0x0a, 0x4d, 0x6f,
	0x6e, 0x74, 0x68, 0x50, 0x72, 0x6f, 0x74, 0x6f, 0x50, 0x01, 0x5a, 0x36, 0x67, 0x6f, 0x6f, 0x67,
	0x6c, 0x65, 0x2e, 0x67, 0x6f, 0x6c, 0x61, 0x6e, 0x67, 0x2e, 0x6f, 0x72, 0x67, 0x2f, 0x67, 0x65,
	0x6e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2f, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x61, 0x70, 0x69,
	0x73, 0x2f, 0x74, 0x79, 0x70, 0x65, 0x2f, 0x6d, 0x6f, 0x6e, 0x74, 0x68, 0x3b, 0x6d, 0x6f, 0x6e,
	0x74, 0x68, 0xa2, 0x02, 0x03, 0x47, 0x54, 0x50, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_google_type_month_proto_rawDescOnce sync.Once
	file_google_type_month_proto_rawDescData = file_google_type_month_proto_rawDesc
)

func file_google_type_month_proto_rawDescGZIP() []byte {
	file_google_type_month_proto_rawDescOnce.Do(func() {
		file_google_type_month_proto_rawDescData = protoimpl.X.CompressGZIP(file_google_type_month_proto_rawDescData)
	})
	return file_google_type_month_proto_rawDescData
}

var file_google_type_month_proto_enumTypes = make([]protoimpl.EnumInfo, 1)
var file_google_type_month_proto_goTypes = []interface{}{
	(Month)(0), // 0: google.type.Month
}
var file_google_type_month_proto_depIdxs = []int32{
	0, // [0:0] is the sub-list for method output_type
	0, // [0:0] is the sub-list for method input_type
	0, // [0:0] is the sub-list for extension type_name
	0, // [0:0] is the sub-list for extension extendee
	0, // [0:0] is the sub-list for field type_name
}

func init() { file_google_type_month_proto_init() }
func file_google_type_month_proto_init() {
	if File_google_type_month_proto != nil {
		return
	}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_google_type_month_proto_rawDesc,
			NumEnums:      1,
			NumMessages:   0,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_google_type_month_proto_goTypes,
		DependencyIndexes: file_google_type_month_proto_depIdxs,
		EnumInfos:         file_google_type_month_proto_enumTypes,
	}.Build()
	File_google_type_month_proto = out.File
	file_google_type_month_proto_rawDesc = nil
	file_google_type_month_proto_goTypes = nil
	file_google_type_month_proto_depIdxs = nil
}
