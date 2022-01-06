// Copyright 2021 Google LLC
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
// 	protoc        v3.12.2
// source: google/cloud/eventarc/publishing/v1/publisher.proto

package publisher

import (
	context "context"
	reflect "reflect"
	sync "sync"

	_ "google.golang.org/genproto/googleapis/api/annotations"
	grpc "google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
	anypb "google.golang.org/protobuf/types/known/anypb"
)

const (
	// Verify that this generated code is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(20 - protoimpl.MinVersion)
	// Verify that runtime/protoimpl is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(protoimpl.MaxVersion - 20)
)

// The request message for the PublishChannelConnectionEvents method.
type PublishChannelConnectionEventsRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// The channel_connection that the events are published from. For example:
	// `projects/{partner_project_id}/locations/{location}/channelConnections/{channel_connection_id}`.
	ChannelConnection string `protobuf:"bytes,1,opt,name=channel_connection,json=channelConnection,proto3" json:"channel_connection,omitempty"`
	// The CloudEvents v1.0 events to publish. No other types are allowed.
	Events []*anypb.Any `protobuf:"bytes,2,rep,name=events,proto3" json:"events,omitempty"`
}

func (x *PublishChannelConnectionEventsRequest) Reset() {
	*x = PublishChannelConnectionEventsRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_google_cloud_eventarc_publishing_v1_publisher_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *PublishChannelConnectionEventsRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*PublishChannelConnectionEventsRequest) ProtoMessage() {}

func (x *PublishChannelConnectionEventsRequest) ProtoReflect() protoreflect.Message {
	mi := &file_google_cloud_eventarc_publishing_v1_publisher_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use PublishChannelConnectionEventsRequest.ProtoReflect.Descriptor instead.
func (*PublishChannelConnectionEventsRequest) Descriptor() ([]byte, []int) {
	return file_google_cloud_eventarc_publishing_v1_publisher_proto_rawDescGZIP(), []int{0}
}

func (x *PublishChannelConnectionEventsRequest) GetChannelConnection() string {
	if x != nil {
		return x.ChannelConnection
	}
	return ""
}

func (x *PublishChannelConnectionEventsRequest) GetEvents() []*anypb.Any {
	if x != nil {
		return x.Events
	}
	return nil
}

// The response message for the PublishChannelConnectionEvents method.
type PublishChannelConnectionEventsResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields
}

func (x *PublishChannelConnectionEventsResponse) Reset() {
	*x = PublishChannelConnectionEventsResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_google_cloud_eventarc_publishing_v1_publisher_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *PublishChannelConnectionEventsResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*PublishChannelConnectionEventsResponse) ProtoMessage() {}

func (x *PublishChannelConnectionEventsResponse) ProtoReflect() protoreflect.Message {
	mi := &file_google_cloud_eventarc_publishing_v1_publisher_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use PublishChannelConnectionEventsResponse.ProtoReflect.Descriptor instead.
func (*PublishChannelConnectionEventsResponse) Descriptor() ([]byte, []int) {
	return file_google_cloud_eventarc_publishing_v1_publisher_proto_rawDescGZIP(), []int{1}
}

var File_google_cloud_eventarc_publishing_v1_publisher_proto protoreflect.FileDescriptor

var file_google_cloud_eventarc_publishing_v1_publisher_proto_rawDesc = []byte{
	0x0a, 0x33, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x63, 0x6c, 0x6f, 0x75, 0x64, 0x2f, 0x65,
	0x76, 0x65, 0x6e, 0x74, 0x61, 0x72, 0x63, 0x2f, 0x70, 0x75, 0x62, 0x6c, 0x69, 0x73, 0x68, 0x69,
	0x6e, 0x67, 0x2f, 0x76, 0x31, 0x2f, 0x70, 0x75, 0x62, 0x6c, 0x69, 0x73, 0x68, 0x65, 0x72, 0x2e,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x23, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x63, 0x6c,
	0x6f, 0x75, 0x64, 0x2e, 0x65, 0x76, 0x65, 0x6e, 0x74, 0x61, 0x72, 0x63, 0x2e, 0x70, 0x75, 0x62,
	0x6c, 0x69, 0x73, 0x68, 0x69, 0x6e, 0x67, 0x2e, 0x76, 0x31, 0x1a, 0x1c, 0x67, 0x6f, 0x6f, 0x67,
	0x6c, 0x65, 0x2f, 0x61, 0x70, 0x69, 0x2f, 0x61, 0x6e, 0x6e, 0x6f, 0x74, 0x61, 0x74, 0x69, 0x6f,
	0x6e, 0x73, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x19, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65,
	0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2f, 0x61, 0x6e, 0x79, 0x2e, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x1a, 0x17, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x61, 0x70, 0x69, 0x2f,
	0x63, 0x6c, 0x69, 0x65, 0x6e, 0x74, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0x84, 0x01, 0x0a,
	0x25, 0x50, 0x75, 0x62, 0x6c, 0x69, 0x73, 0x68, 0x43, 0x68, 0x61, 0x6e, 0x6e, 0x65, 0x6c, 0x43,
	0x6f, 0x6e, 0x6e, 0x65, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x45, 0x76, 0x65, 0x6e, 0x74, 0x73, 0x52,
	0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x2d, 0x0a, 0x12, 0x63, 0x68, 0x61, 0x6e, 0x6e, 0x65,
	0x6c, 0x5f, 0x63, 0x6f, 0x6e, 0x6e, 0x65, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x18, 0x01, 0x20, 0x01,
	0x28, 0x09, 0x52, 0x11, 0x63, 0x68, 0x61, 0x6e, 0x6e, 0x65, 0x6c, 0x43, 0x6f, 0x6e, 0x6e, 0x65,
	0x63, 0x74, 0x69, 0x6f, 0x6e, 0x12, 0x2c, 0x0a, 0x06, 0x65, 0x76, 0x65, 0x6e, 0x74, 0x73, 0x18,
	0x02, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x14, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x41, 0x6e, 0x79, 0x52, 0x06, 0x65, 0x76, 0x65,
	0x6e, 0x74, 0x73, 0x22, 0x28, 0x0a, 0x26, 0x50, 0x75, 0x62, 0x6c, 0x69, 0x73, 0x68, 0x43, 0x68,
	0x61, 0x6e, 0x6e, 0x65, 0x6c, 0x43, 0x6f, 0x6e, 0x6e, 0x65, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x45,
	0x76, 0x65, 0x6e, 0x74, 0x73, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x32, 0xfd, 0x02,
	0x0a, 0x09, 0x50, 0x75, 0x62, 0x6c, 0x69, 0x73, 0x68, 0x65, 0x72, 0x12, 0x98, 0x02, 0x0a, 0x1e,
	0x50, 0x75, 0x62, 0x6c, 0x69, 0x73, 0x68, 0x43, 0x68, 0x61, 0x6e, 0x6e, 0x65, 0x6c, 0x43, 0x6f,
	0x6e, 0x6e, 0x65, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x45, 0x76, 0x65, 0x6e, 0x74, 0x73, 0x12, 0x4a,
	0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x63, 0x6c, 0x6f, 0x75, 0x64, 0x2e, 0x65, 0x76,
	0x65, 0x6e, 0x74, 0x61, 0x72, 0x63, 0x2e, 0x70, 0x75, 0x62, 0x6c, 0x69, 0x73, 0x68, 0x69, 0x6e,
	0x67, 0x2e, 0x76, 0x31, 0x2e, 0x50, 0x75, 0x62, 0x6c, 0x69, 0x73, 0x68, 0x43, 0x68, 0x61, 0x6e,
	0x6e, 0x65, 0x6c, 0x43, 0x6f, 0x6e, 0x6e, 0x65, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x45, 0x76, 0x65,
	0x6e, 0x74, 0x73, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x4b, 0x2e, 0x67, 0x6f, 0x6f,
	0x67, 0x6c, 0x65, 0x2e, 0x63, 0x6c, 0x6f, 0x75, 0x64, 0x2e, 0x65, 0x76, 0x65, 0x6e, 0x74, 0x61,
	0x72, 0x63, 0x2e, 0x70, 0x75, 0x62, 0x6c, 0x69, 0x73, 0x68, 0x69, 0x6e, 0x67, 0x2e, 0x76, 0x31,
	0x2e, 0x50, 0x75, 0x62, 0x6c, 0x69, 0x73, 0x68, 0x43, 0x68, 0x61, 0x6e, 0x6e, 0x65, 0x6c, 0x43,
	0x6f, 0x6e, 0x6e, 0x65, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x45, 0x76, 0x65, 0x6e, 0x74, 0x73, 0x52,
	0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x5d, 0x82, 0xd3, 0xe4, 0x93, 0x02, 0x57, 0x22,
	0x52, 0x2f, 0x76, 0x31, 0x2f, 0x7b, 0x63, 0x68, 0x61, 0x6e, 0x6e, 0x65, 0x6c, 0x5f, 0x63, 0x6f,
	0x6e, 0x6e, 0x65, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x3d, 0x70, 0x72, 0x6f, 0x6a, 0x65, 0x63, 0x74,
	0x73, 0x2f, 0x2a, 0x2f, 0x6c, 0x6f, 0x63, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x2f, 0x2a, 0x2f,
	0x63, 0x68, 0x61, 0x6e, 0x6e, 0x65, 0x6c, 0x43, 0x6f, 0x6e, 0x6e, 0x65, 0x63, 0x74, 0x69, 0x6f,
	0x6e, 0x73, 0x2f, 0x2a, 0x7d, 0x3a, 0x70, 0x75, 0x62, 0x6c, 0x69, 0x73, 0x68, 0x45, 0x76, 0x65,
	0x6e, 0x74, 0x73, 0x3a, 0x01, 0x2a, 0x1a, 0x55, 0xca, 0x41, 0x21, 0x65, 0x76, 0x65, 0x6e, 0x74,
	0x61, 0x72, 0x63, 0x70, 0x75, 0x62, 0x6c, 0x69, 0x73, 0x68, 0x69, 0x6e, 0x67, 0x2e, 0x67, 0x6f,
	0x6f, 0x67, 0x6c, 0x65, 0x61, 0x70, 0x69, 0x73, 0x2e, 0x63, 0x6f, 0x6d, 0xd2, 0x41, 0x2e, 0x68,
	0x74, 0x74, 0x70, 0x73, 0x3a, 0x2f, 0x2f, 0x77, 0x77, 0x77, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c,
	0x65, 0x61, 0x70, 0x69, 0x73, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x61, 0x75, 0x74, 0x68, 0x2f, 0x63,
	0x6c, 0x6f, 0x75, 0x64, 0x2d, 0x70, 0x6c, 0x61, 0x74, 0x66, 0x6f, 0x72, 0x6d, 0x42, 0xff, 0x01,
	0x0a, 0x27, 0x63, 0x6f, 0x6d, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x63, 0x6c, 0x6f,
	0x75, 0x64, 0x2e, 0x65, 0x76, 0x65, 0x6e, 0x74, 0x61, 0x72, 0x63, 0x2e, 0x70, 0x75, 0x62, 0x6c,
	0x69, 0x73, 0x68, 0x69, 0x6e, 0x67, 0x2e, 0x76, 0x31, 0x42, 0x0e, 0x50, 0x75, 0x62, 0x6c, 0x69,
	0x73, 0x68, 0x65, 0x72, 0x50, 0x72, 0x6f, 0x74, 0x6f, 0x50, 0x01, 0x5a, 0x4c, 0x67, 0x6f, 0x6f,
	0x67, 0x6c, 0x65, 0x2e, 0x67, 0x6f, 0x6c, 0x61, 0x6e, 0x67, 0x2e, 0x6f, 0x72, 0x67, 0x2f, 0x67,
	0x65, 0x6e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2f, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x61, 0x70,
	0x69, 0x73, 0x2f, 0x63, 0x6c, 0x6f, 0x75, 0x64, 0x2f, 0x65, 0x76, 0x65, 0x6e, 0x74, 0x61, 0x72,
	0x63, 0x2f, 0x70, 0x75, 0x62, 0x6c, 0x69, 0x73, 0x68, 0x69, 0x6e, 0x67, 0x2f, 0x76, 0x31, 0x3b,
	0x70, 0x75, 0x62, 0x6c, 0x69, 0x73, 0x68, 0x65, 0x72, 0xaa, 0x02, 0x23, 0x47, 0x6f, 0x6f, 0x67,
	0x6c, 0x65, 0x2e, 0x43, 0x6c, 0x6f, 0x75, 0x64, 0x2e, 0x45, 0x76, 0x65, 0x6e, 0x74, 0x61, 0x72,
	0x63, 0x2e, 0x50, 0x75, 0x62, 0x6c, 0x69, 0x73, 0x68, 0x69, 0x6e, 0x67, 0x2e, 0x56, 0x31, 0xca,
	0x02, 0x23, 0x47, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x5c, 0x43, 0x6c, 0x6f, 0x75, 0x64, 0x5c, 0x45,
	0x76, 0x65, 0x6e, 0x74, 0x61, 0x72, 0x63, 0x5c, 0x50, 0x75, 0x62, 0x6c, 0x69, 0x73, 0x68, 0x69,
	0x6e, 0x67, 0x5c, 0x56, 0x31, 0xea, 0x02, 0x27, 0x47, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x3a, 0x3a,
	0x43, 0x6c, 0x6f, 0x75, 0x64, 0x3a, 0x3a, 0x45, 0x76, 0x65, 0x6e, 0x74, 0x61, 0x72, 0x63, 0x3a,
	0x3a, 0x50, 0x75, 0x62, 0x6c, 0x69, 0x73, 0x68, 0x69, 0x6e, 0x67, 0x3a, 0x3a, 0x56, 0x31, 0x62,
	0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_google_cloud_eventarc_publishing_v1_publisher_proto_rawDescOnce sync.Once
	file_google_cloud_eventarc_publishing_v1_publisher_proto_rawDescData = file_google_cloud_eventarc_publishing_v1_publisher_proto_rawDesc
)

func file_google_cloud_eventarc_publishing_v1_publisher_proto_rawDescGZIP() []byte {
	file_google_cloud_eventarc_publishing_v1_publisher_proto_rawDescOnce.Do(func() {
		file_google_cloud_eventarc_publishing_v1_publisher_proto_rawDescData = protoimpl.X.CompressGZIP(file_google_cloud_eventarc_publishing_v1_publisher_proto_rawDescData)
	})
	return file_google_cloud_eventarc_publishing_v1_publisher_proto_rawDescData
}

var file_google_cloud_eventarc_publishing_v1_publisher_proto_msgTypes = make([]protoimpl.MessageInfo, 2)
var file_google_cloud_eventarc_publishing_v1_publisher_proto_goTypes = []interface{}{
	(*PublishChannelConnectionEventsRequest)(nil),  // 0: google.cloud.eventarc.publishing.v1.PublishChannelConnectionEventsRequest
	(*PublishChannelConnectionEventsResponse)(nil), // 1: google.cloud.eventarc.publishing.v1.PublishChannelConnectionEventsResponse
	(*anypb.Any)(nil), // 2: google.protobuf.Any
}
var file_google_cloud_eventarc_publishing_v1_publisher_proto_depIdxs = []int32{
	2, // 0: google.cloud.eventarc.publishing.v1.PublishChannelConnectionEventsRequest.events:type_name -> google.protobuf.Any
	0, // 1: google.cloud.eventarc.publishing.v1.Publisher.PublishChannelConnectionEvents:input_type -> google.cloud.eventarc.publishing.v1.PublishChannelConnectionEventsRequest
	1, // 2: google.cloud.eventarc.publishing.v1.Publisher.PublishChannelConnectionEvents:output_type -> google.cloud.eventarc.publishing.v1.PublishChannelConnectionEventsResponse
	2, // [2:3] is the sub-list for method output_type
	1, // [1:2] is the sub-list for method input_type
	1, // [1:1] is the sub-list for extension type_name
	1, // [1:1] is the sub-list for extension extendee
	0, // [0:1] is the sub-list for field type_name
}

func init() { file_google_cloud_eventarc_publishing_v1_publisher_proto_init() }
func file_google_cloud_eventarc_publishing_v1_publisher_proto_init() {
	if File_google_cloud_eventarc_publishing_v1_publisher_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_google_cloud_eventarc_publishing_v1_publisher_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*PublishChannelConnectionEventsRequest); i {
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
		file_google_cloud_eventarc_publishing_v1_publisher_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*PublishChannelConnectionEventsResponse); i {
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
			RawDescriptor: file_google_cloud_eventarc_publishing_v1_publisher_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   2,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_google_cloud_eventarc_publishing_v1_publisher_proto_goTypes,
		DependencyIndexes: file_google_cloud_eventarc_publishing_v1_publisher_proto_depIdxs,
		MessageInfos:      file_google_cloud_eventarc_publishing_v1_publisher_proto_msgTypes,
	}.Build()
	File_google_cloud_eventarc_publishing_v1_publisher_proto = out.File
	file_google_cloud_eventarc_publishing_v1_publisher_proto_rawDesc = nil
	file_google_cloud_eventarc_publishing_v1_publisher_proto_goTypes = nil
	file_google_cloud_eventarc_publishing_v1_publisher_proto_depIdxs = nil
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConnInterface

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion6

// PublisherClient is the client API for Publisher service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type PublisherClient interface {
	// Publish events to a ChannelConnection in a partner's project.
	PublishChannelConnectionEvents(ctx context.Context, in *PublishChannelConnectionEventsRequest, opts ...grpc.CallOption) (*PublishChannelConnectionEventsResponse, error)
}

type publisherClient struct {
	cc grpc.ClientConnInterface
}

func NewPublisherClient(cc grpc.ClientConnInterface) PublisherClient {
	return &publisherClient{cc}
}

func (c *publisherClient) PublishChannelConnectionEvents(ctx context.Context, in *PublishChannelConnectionEventsRequest, opts ...grpc.CallOption) (*PublishChannelConnectionEventsResponse, error) {
	out := new(PublishChannelConnectionEventsResponse)
	err := c.cc.Invoke(ctx, "/google.cloud.eventarc.publishing.v1.Publisher/PublishChannelConnectionEvents", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// PublisherServer is the server API for Publisher service.
type PublisherServer interface {
	// Publish events to a ChannelConnection in a partner's project.
	PublishChannelConnectionEvents(context.Context, *PublishChannelConnectionEventsRequest) (*PublishChannelConnectionEventsResponse, error)
}

// UnimplementedPublisherServer can be embedded to have forward compatible implementations.
type UnimplementedPublisherServer struct {
}

func (*UnimplementedPublisherServer) PublishChannelConnectionEvents(context.Context, *PublishChannelConnectionEventsRequest) (*PublishChannelConnectionEventsResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method PublishChannelConnectionEvents not implemented")
}

func RegisterPublisherServer(s *grpc.Server, srv PublisherServer) {
	s.RegisterService(&_Publisher_serviceDesc, srv)
}

func _Publisher_PublishChannelConnectionEvents_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(PublishChannelConnectionEventsRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(PublisherServer).PublishChannelConnectionEvents(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/google.cloud.eventarc.publishing.v1.Publisher/PublishChannelConnectionEvents",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(PublisherServer).PublishChannelConnectionEvents(ctx, req.(*PublishChannelConnectionEventsRequest))
	}
	return interceptor(ctx, in, info, handler)
}

var _Publisher_serviceDesc = grpc.ServiceDesc{
	ServiceName: "google.cloud.eventarc.publishing.v1.Publisher",
	HandlerType: (*PublisherServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "PublishChannelConnectionEvents",
			Handler:    _Publisher_PublishChannelConnectionEvents_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "google/cloud/eventarc/publishing/v1/publisher.proto",
}