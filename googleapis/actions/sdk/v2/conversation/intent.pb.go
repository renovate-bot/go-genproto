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
// source: google/actions/sdk/v2/conversation/intent.proto

package conversation

import (
	reflect "reflect"
	sync "sync"

	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
	structpb "google.golang.org/protobuf/types/known/structpb"
)

const (
	// Verify that this generated code is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(20 - protoimpl.MinVersion)
	// Verify that runtime/protoimpl is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(protoimpl.MaxVersion - 20)
)

// Represents an intent.
type Intent struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// Required. The name of the last matched intent.
	Name string `protobuf:"bytes,1,opt,name=name,proto3" json:"name,omitempty"`
	// Required. Represents parameters identified as part of intent matching.
	// This is a map of the name of the identified parameter to the value of the
	// parameter identified from user input. All parameters defined in
	// the matched intent that are identified will be surfaced here.
	Params map[string]*IntentParameterValue `protobuf:"bytes,2,rep,name=params,proto3" json:"params,omitempty" protobuf_key:"bytes,1,opt,name=key,proto3" protobuf_val:"bytes,2,opt,name=value,proto3"`
	// Optional. Typed or spoken input from the end user that matched this intent.
	// This will be populated when an intent is matched, based on the user input.
	Query string `protobuf:"bytes,3,opt,name=query,proto3" json:"query,omitempty"`
}

func (x *Intent) Reset() {
	*x = Intent{}
	if protoimpl.UnsafeEnabled {
		mi := &file_google_actions_sdk_v2_conversation_intent_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Intent) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Intent) ProtoMessage() {}

func (x *Intent) ProtoReflect() protoreflect.Message {
	mi := &file_google_actions_sdk_v2_conversation_intent_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Intent.ProtoReflect.Descriptor instead.
func (*Intent) Descriptor() ([]byte, []int) {
	return file_google_actions_sdk_v2_conversation_intent_proto_rawDescGZIP(), []int{0}
}

func (x *Intent) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

func (x *Intent) GetParams() map[string]*IntentParameterValue {
	if x != nil {
		return x.Params
	}
	return nil
}

func (x *Intent) GetQuery() string {
	if x != nil {
		return x.Query
	}
	return ""
}

// Represents a value for intent parameter.
type IntentParameterValue struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// Required. Original text value extracted from user utterance.
	Original string `protobuf:"bytes,1,opt,name=original,proto3" json:"original,omitempty"`
	// Required. Structured value for parameter extracted from user input.
	// This will only be populated if the parameter is defined in the matched
	// intent and the value of the parameter could be identified during intent
	// matching.
	Resolved *structpb.Value `protobuf:"bytes,2,opt,name=resolved,proto3" json:"resolved,omitempty"`
}

func (x *IntentParameterValue) Reset() {
	*x = IntentParameterValue{}
	if protoimpl.UnsafeEnabled {
		mi := &file_google_actions_sdk_v2_conversation_intent_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *IntentParameterValue) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*IntentParameterValue) ProtoMessage() {}

func (x *IntentParameterValue) ProtoReflect() protoreflect.Message {
	mi := &file_google_actions_sdk_v2_conversation_intent_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use IntentParameterValue.ProtoReflect.Descriptor instead.
func (*IntentParameterValue) Descriptor() ([]byte, []int) {
	return file_google_actions_sdk_v2_conversation_intent_proto_rawDescGZIP(), []int{1}
}

func (x *IntentParameterValue) GetOriginal() string {
	if x != nil {
		return x.Original
	}
	return ""
}

func (x *IntentParameterValue) GetResolved() *structpb.Value {
	if x != nil {
		return x.Resolved
	}
	return nil
}

var File_google_actions_sdk_v2_conversation_intent_proto protoreflect.FileDescriptor

var file_google_actions_sdk_v2_conversation_intent_proto_rawDesc = []byte{
	0x0a, 0x2f, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x61, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x73,
	0x2f, 0x73, 0x64, 0x6b, 0x2f, 0x76, 0x32, 0x2f, 0x63, 0x6f, 0x6e, 0x76, 0x65, 0x72, 0x73, 0x61,
	0x74, 0x69, 0x6f, 0x6e, 0x2f, 0x69, 0x6e, 0x74, 0x65, 0x6e, 0x74, 0x2e, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x12, 0x22, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x61, 0x63, 0x74, 0x69, 0x6f, 0x6e,
	0x73, 0x2e, 0x73, 0x64, 0x6b, 0x2e, 0x76, 0x32, 0x2e, 0x63, 0x6f, 0x6e, 0x76, 0x65, 0x72, 0x73,
	0x61, 0x74, 0x69, 0x6f, 0x6e, 0x1a, 0x1c, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2f, 0x73, 0x74, 0x72, 0x75, 0x63, 0x74, 0x2e, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x22, 0xf7, 0x01, 0x0a, 0x06, 0x49, 0x6e, 0x74, 0x65, 0x6e, 0x74, 0x12, 0x12,
	0x0a, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x6e, 0x61,
	0x6d, 0x65, 0x12, 0x4e, 0x0a, 0x06, 0x70, 0x61, 0x72, 0x61, 0x6d, 0x73, 0x18, 0x02, 0x20, 0x03,
	0x28, 0x0b, 0x32, 0x36, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x61, 0x63, 0x74, 0x69,
	0x6f, 0x6e, 0x73, 0x2e, 0x73, 0x64, 0x6b, 0x2e, 0x76, 0x32, 0x2e, 0x63, 0x6f, 0x6e, 0x76, 0x65,
	0x72, 0x73, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x2e, 0x49, 0x6e, 0x74, 0x65, 0x6e, 0x74, 0x2e, 0x50,
	0x61, 0x72, 0x61, 0x6d, 0x73, 0x45, 0x6e, 0x74, 0x72, 0x79, 0x52, 0x06, 0x70, 0x61, 0x72, 0x61,
	0x6d, 0x73, 0x12, 0x14, 0x0a, 0x05, 0x71, 0x75, 0x65, 0x72, 0x79, 0x18, 0x03, 0x20, 0x01, 0x28,
	0x09, 0x52, 0x05, 0x71, 0x75, 0x65, 0x72, 0x79, 0x1a, 0x73, 0x0a, 0x0b, 0x50, 0x61, 0x72, 0x61,
	0x6d, 0x73, 0x45, 0x6e, 0x74, 0x72, 0x79, 0x12, 0x10, 0x0a, 0x03, 0x6b, 0x65, 0x79, 0x18, 0x01,
	0x20, 0x01, 0x28, 0x09, 0x52, 0x03, 0x6b, 0x65, 0x79, 0x12, 0x4e, 0x0a, 0x05, 0x76, 0x61, 0x6c,
	0x75, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x38, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c,
	0x65, 0x2e, 0x61, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x2e, 0x73, 0x64, 0x6b, 0x2e, 0x76, 0x32,
	0x2e, 0x63, 0x6f, 0x6e, 0x76, 0x65, 0x72, 0x73, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x2e, 0x49, 0x6e,
	0x74, 0x65, 0x6e, 0x74, 0x50, 0x61, 0x72, 0x61, 0x6d, 0x65, 0x74, 0x65, 0x72, 0x56, 0x61, 0x6c,
	0x75, 0x65, 0x52, 0x05, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x3a, 0x02, 0x38, 0x01, 0x22, 0x66, 0x0a,
	0x14, 0x49, 0x6e, 0x74, 0x65, 0x6e, 0x74, 0x50, 0x61, 0x72, 0x61, 0x6d, 0x65, 0x74, 0x65, 0x72,
	0x56, 0x61, 0x6c, 0x75, 0x65, 0x12, 0x1a, 0x0a, 0x08, 0x6f, 0x72, 0x69, 0x67, 0x69, 0x6e, 0x61,
	0x6c, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x6f, 0x72, 0x69, 0x67, 0x69, 0x6e, 0x61,
	0x6c, 0x12, 0x32, 0x0a, 0x08, 0x72, 0x65, 0x73, 0x6f, 0x6c, 0x76, 0x65, 0x64, 0x18, 0x02, 0x20,
	0x01, 0x28, 0x0b, 0x32, 0x16, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x56, 0x61, 0x6c, 0x75, 0x65, 0x52, 0x08, 0x72, 0x65, 0x73,
	0x6f, 0x6c, 0x76, 0x65, 0x64, 0x42, 0x87, 0x01, 0x0a, 0x26, 0x63, 0x6f, 0x6d, 0x2e, 0x67, 0x6f,
	0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x61, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x2e, 0x73, 0x64, 0x6b,
	0x2e, 0x76, 0x32, 0x2e, 0x63, 0x6f, 0x6e, 0x76, 0x65, 0x72, 0x73, 0x61, 0x74, 0x69, 0x6f, 0x6e,
	0x42, 0x0b, 0x49, 0x6e, 0x74, 0x65, 0x6e, 0x74, 0x50, 0x72, 0x6f, 0x74, 0x6f, 0x50, 0x01, 0x5a,
	0x4e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x67, 0x6f, 0x6c, 0x61, 0x6e, 0x67, 0x2e, 0x6f,
	0x72, 0x67, 0x2f, 0x67, 0x65, 0x6e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2f, 0x67, 0x6f, 0x6f, 0x67,
	0x6c, 0x65, 0x61, 0x70, 0x69, 0x73, 0x2f, 0x61, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x2f, 0x73,
	0x64, 0x6b, 0x2f, 0x76, 0x32, 0x2f, 0x63, 0x6f, 0x6e, 0x76, 0x65, 0x72, 0x73, 0x61, 0x74, 0x69,
	0x6f, 0x6e, 0x3b, 0x63, 0x6f, 0x6e, 0x76, 0x65, 0x72, 0x73, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x62,
	0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_google_actions_sdk_v2_conversation_intent_proto_rawDescOnce sync.Once
	file_google_actions_sdk_v2_conversation_intent_proto_rawDescData = file_google_actions_sdk_v2_conversation_intent_proto_rawDesc
)

func file_google_actions_sdk_v2_conversation_intent_proto_rawDescGZIP() []byte {
	file_google_actions_sdk_v2_conversation_intent_proto_rawDescOnce.Do(func() {
		file_google_actions_sdk_v2_conversation_intent_proto_rawDescData = protoimpl.X.CompressGZIP(file_google_actions_sdk_v2_conversation_intent_proto_rawDescData)
	})
	return file_google_actions_sdk_v2_conversation_intent_proto_rawDescData
}

var file_google_actions_sdk_v2_conversation_intent_proto_msgTypes = make([]protoimpl.MessageInfo, 3)
var file_google_actions_sdk_v2_conversation_intent_proto_goTypes = []interface{}{
	(*Intent)(nil),               // 0: google.actions.sdk.v2.conversation.Intent
	(*IntentParameterValue)(nil), // 1: google.actions.sdk.v2.conversation.IntentParameterValue
	nil,                          // 2: google.actions.sdk.v2.conversation.Intent.ParamsEntry
	(*structpb.Value)(nil),       // 3: google.protobuf.Value
}
var file_google_actions_sdk_v2_conversation_intent_proto_depIdxs = []int32{
	2, // 0: google.actions.sdk.v2.conversation.Intent.params:type_name -> google.actions.sdk.v2.conversation.Intent.ParamsEntry
	3, // 1: google.actions.sdk.v2.conversation.IntentParameterValue.resolved:type_name -> google.protobuf.Value
	1, // 2: google.actions.sdk.v2.conversation.Intent.ParamsEntry.value:type_name -> google.actions.sdk.v2.conversation.IntentParameterValue
	3, // [3:3] is the sub-list for method output_type
	3, // [3:3] is the sub-list for method input_type
	3, // [3:3] is the sub-list for extension type_name
	3, // [3:3] is the sub-list for extension extendee
	0, // [0:3] is the sub-list for field type_name
}

func init() { file_google_actions_sdk_v2_conversation_intent_proto_init() }
func file_google_actions_sdk_v2_conversation_intent_proto_init() {
	if File_google_actions_sdk_v2_conversation_intent_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_google_actions_sdk_v2_conversation_intent_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Intent); i {
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
		file_google_actions_sdk_v2_conversation_intent_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*IntentParameterValue); i {
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
			RawDescriptor: file_google_actions_sdk_v2_conversation_intent_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   3,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_google_actions_sdk_v2_conversation_intent_proto_goTypes,
		DependencyIndexes: file_google_actions_sdk_v2_conversation_intent_proto_depIdxs,
		MessageInfos:      file_google_actions_sdk_v2_conversation_intent_proto_msgTypes,
	}.Build()
	File_google_actions_sdk_v2_conversation_intent_proto = out.File
	file_google_actions_sdk_v2_conversation_intent_proto_rawDesc = nil
	file_google_actions_sdk_v2_conversation_intent_proto_goTypes = nil
	file_google_actions_sdk_v2_conversation_intent_proto_depIdxs = nil
}
