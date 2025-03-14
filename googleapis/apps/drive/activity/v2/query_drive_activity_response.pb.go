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
// source: google/apps/drive/activity/v2/query_drive_activity_response.proto

package activity

import (
	reflect "reflect"
	sync "sync"

	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
	timestamppb "google.golang.org/protobuf/types/known/timestamppb"
)

const (
	// Verify that this generated code is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(20 - protoimpl.MinVersion)
	// Verify that runtime/protoimpl is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(protoimpl.MaxVersion - 20)
)

// Response message for querying Drive activity.
type QueryDriveActivityResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// List of activity requested.
	Activities []*DriveActivity `protobuf:"bytes,1,rep,name=activities,proto3" json:"activities,omitempty"`
	// Token to retrieve the next page of results, or
	// empty if there are no more results in the list.
	NextPageToken string `protobuf:"bytes,2,opt,name=next_page_token,json=nextPageToken,proto3" json:"next_page_token,omitempty"`
}

func (x *QueryDriveActivityResponse) Reset() {
	*x = QueryDriveActivityResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_google_apps_drive_activity_v2_query_drive_activity_response_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *QueryDriveActivityResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*QueryDriveActivityResponse) ProtoMessage() {}

func (x *QueryDriveActivityResponse) ProtoReflect() protoreflect.Message {
	mi := &file_google_apps_drive_activity_v2_query_drive_activity_response_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use QueryDriveActivityResponse.ProtoReflect.Descriptor instead.
func (*QueryDriveActivityResponse) Descriptor() ([]byte, []int) {
	return file_google_apps_drive_activity_v2_query_drive_activity_response_proto_rawDescGZIP(), []int{0}
}

func (x *QueryDriveActivityResponse) GetActivities() []*DriveActivity {
	if x != nil {
		return x.Activities
	}
	return nil
}

func (x *QueryDriveActivityResponse) GetNextPageToken() string {
	if x != nil {
		return x.NextPageToken
	}
	return ""
}

// A single Drive activity comprising one or more Actions by one or more
// Actors on one or more Targets. Some Action groupings occur spontaneously,
// such as moving an item into a shared folder triggering a permission change.
// Other groupings of related Actions, such as multiple Actors editing one item
// or moving multiple files into a new folder, are controlled by the selection
// of a ConsolidationStrategy in the QueryDriveActivityRequest.
type DriveActivity struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// Key information about the primary action for this activity. This is either
	// representative, or the most important, of all actions in the activity,
	// according to the ConsolidationStrategy in the request.
	PrimaryActionDetail *ActionDetail `protobuf:"bytes,2,opt,name=primary_action_detail,json=primaryActionDetail,proto3" json:"primary_action_detail,omitempty"`
	// All actor(s) responsible for the activity.
	Actors []*Actor `protobuf:"bytes,3,rep,name=actors,proto3" json:"actors,omitempty"`
	// Details on all actions in this activity.
	Actions []*Action `protobuf:"bytes,4,rep,name=actions,proto3" json:"actions,omitempty"`
	// All Google Drive objects this activity is about (e.g. file, folder, drive).
	// This represents the state of the target immediately after the actions
	// occurred.
	Targets []*Target `protobuf:"bytes,5,rep,name=targets,proto3" json:"targets,omitempty"`
	// The period of time when this activity occurred.
	//
	// Types that are assignable to Time:
	//
	//	*DriveActivity_Timestamp
	//	*DriveActivity_TimeRange
	Time isDriveActivity_Time `protobuf_oneof:"time"`
}

func (x *DriveActivity) Reset() {
	*x = DriveActivity{}
	if protoimpl.UnsafeEnabled {
		mi := &file_google_apps_drive_activity_v2_query_drive_activity_response_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *DriveActivity) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*DriveActivity) ProtoMessage() {}

func (x *DriveActivity) ProtoReflect() protoreflect.Message {
	mi := &file_google_apps_drive_activity_v2_query_drive_activity_response_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use DriveActivity.ProtoReflect.Descriptor instead.
func (*DriveActivity) Descriptor() ([]byte, []int) {
	return file_google_apps_drive_activity_v2_query_drive_activity_response_proto_rawDescGZIP(), []int{1}
}

func (x *DriveActivity) GetPrimaryActionDetail() *ActionDetail {
	if x != nil {
		return x.PrimaryActionDetail
	}
	return nil
}

func (x *DriveActivity) GetActors() []*Actor {
	if x != nil {
		return x.Actors
	}
	return nil
}

func (x *DriveActivity) GetActions() []*Action {
	if x != nil {
		return x.Actions
	}
	return nil
}

func (x *DriveActivity) GetTargets() []*Target {
	if x != nil {
		return x.Targets
	}
	return nil
}

func (m *DriveActivity) GetTime() isDriveActivity_Time {
	if m != nil {
		return m.Time
	}
	return nil
}

func (x *DriveActivity) GetTimestamp() *timestamppb.Timestamp {
	if x, ok := x.GetTime().(*DriveActivity_Timestamp); ok {
		return x.Timestamp
	}
	return nil
}

func (x *DriveActivity) GetTimeRange() *TimeRange {
	if x, ok := x.GetTime().(*DriveActivity_TimeRange); ok {
		return x.TimeRange
	}
	return nil
}

type isDriveActivity_Time interface {
	isDriveActivity_Time()
}

type DriveActivity_Timestamp struct {
	// The activity occurred at this specific time.
	Timestamp *timestamppb.Timestamp `protobuf:"bytes,6,opt,name=timestamp,proto3,oneof"`
}

type DriveActivity_TimeRange struct {
	// The activity occurred over this time range.
	TimeRange *TimeRange `protobuf:"bytes,7,opt,name=time_range,json=timeRange,proto3,oneof"`
}

func (*DriveActivity_Timestamp) isDriveActivity_Time() {}

func (*DriveActivity_TimeRange) isDriveActivity_Time() {}

var File_google_apps_drive_activity_v2_query_drive_activity_response_proto protoreflect.FileDescriptor

var file_google_apps_drive_activity_v2_query_drive_activity_response_proto_rawDesc = []byte{
	0x0a, 0x41, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x61, 0x70, 0x70, 0x73, 0x2f, 0x64, 0x72,
	0x69, 0x76, 0x65, 0x2f, 0x61, 0x63, 0x74, 0x69, 0x76, 0x69, 0x74, 0x79, 0x2f, 0x76, 0x32, 0x2f,
	0x71, 0x75, 0x65, 0x72, 0x79, 0x5f, 0x64, 0x72, 0x69, 0x76, 0x65, 0x5f, 0x61, 0x63, 0x74, 0x69,
	0x76, 0x69, 0x74, 0x79, 0x5f, 0x72, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x2e, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x12, 0x1d, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x61, 0x70, 0x70, 0x73,
	0x2e, 0x64, 0x72, 0x69, 0x76, 0x65, 0x2e, 0x61, 0x63, 0x74, 0x69, 0x76, 0x69, 0x74, 0x79, 0x2e,
	0x76, 0x32, 0x1a, 0x2a, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x61, 0x70, 0x70, 0x73, 0x2f,
	0x64, 0x72, 0x69, 0x76, 0x65, 0x2f, 0x61, 0x63, 0x74, 0x69, 0x76, 0x69, 0x74, 0x79, 0x2f, 0x76,
	0x32, 0x2f, 0x61, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x29,
	0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x61, 0x70, 0x70, 0x73, 0x2f, 0x64, 0x72, 0x69, 0x76,
	0x65, 0x2f, 0x61, 0x63, 0x74, 0x69, 0x76, 0x69, 0x74, 0x79, 0x2f, 0x76, 0x32, 0x2f, 0x61, 0x63,
	0x74, 0x6f, 0x72, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x2a, 0x67, 0x6f, 0x6f, 0x67, 0x6c,
	0x65, 0x2f, 0x61, 0x70, 0x70, 0x73, 0x2f, 0x64, 0x72, 0x69, 0x76, 0x65, 0x2f, 0x61, 0x63, 0x74,
	0x69, 0x76, 0x69, 0x74, 0x79, 0x2f, 0x76, 0x32, 0x2f, 0x63, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x2e,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x2a, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x61, 0x70,
	0x70, 0x73, 0x2f, 0x64, 0x72, 0x69, 0x76, 0x65, 0x2f, 0x61, 0x63, 0x74, 0x69, 0x76, 0x69, 0x74,
	0x79, 0x2f, 0x76, 0x32, 0x2f, 0x74, 0x61, 0x72, 0x67, 0x65, 0x74, 0x2e, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x1a, 0x1f, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62,
	0x75, 0x66, 0x2f, 0x74, 0x69, 0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d, 0x70, 0x2e, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x22, 0x92, 0x01, 0x0a, 0x1a, 0x51, 0x75, 0x65, 0x72, 0x79, 0x44, 0x72, 0x69, 0x76,
	0x65, 0x41, 0x63, 0x74, 0x69, 0x76, 0x69, 0x74, 0x79, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73,
	0x65, 0x12, 0x4c, 0x0a, 0x0a, 0x61, 0x63, 0x74, 0x69, 0x76, 0x69, 0x74, 0x69, 0x65, 0x73, 0x18,
	0x01, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x2c, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x61,
	0x70, 0x70, 0x73, 0x2e, 0x64, 0x72, 0x69, 0x76, 0x65, 0x2e, 0x61, 0x63, 0x74, 0x69, 0x76, 0x69,
	0x74, 0x79, 0x2e, 0x76, 0x32, 0x2e, 0x44, 0x72, 0x69, 0x76, 0x65, 0x41, 0x63, 0x74, 0x69, 0x76,
	0x69, 0x74, 0x79, 0x52, 0x0a, 0x61, 0x63, 0x74, 0x69, 0x76, 0x69, 0x74, 0x69, 0x65, 0x73, 0x12,
	0x26, 0x0a, 0x0f, 0x6e, 0x65, 0x78, 0x74, 0x5f, 0x70, 0x61, 0x67, 0x65, 0x5f, 0x74, 0x6f, 0x6b,
	0x65, 0x6e, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0d, 0x6e, 0x65, 0x78, 0x74, 0x50, 0x61,
	0x67, 0x65, 0x54, 0x6f, 0x6b, 0x65, 0x6e, 0x22, 0xbf, 0x03, 0x0a, 0x0d, 0x44, 0x72, 0x69, 0x76,
	0x65, 0x41, 0x63, 0x74, 0x69, 0x76, 0x69, 0x74, 0x79, 0x12, 0x5f, 0x0a, 0x15, 0x70, 0x72, 0x69,
	0x6d, 0x61, 0x72, 0x79, 0x5f, 0x61, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x5f, 0x64, 0x65, 0x74, 0x61,
	0x69, 0x6c, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x2b, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c,
	0x65, 0x2e, 0x61, 0x70, 0x70, 0x73, 0x2e, 0x64, 0x72, 0x69, 0x76, 0x65, 0x2e, 0x61, 0x63, 0x74,
	0x69, 0x76, 0x69, 0x74, 0x79, 0x2e, 0x76, 0x32, 0x2e, 0x41, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x44,
	0x65, 0x74, 0x61, 0x69, 0x6c, 0x52, 0x13, 0x70, 0x72, 0x69, 0x6d, 0x61, 0x72, 0x79, 0x41, 0x63,
	0x74, 0x69, 0x6f, 0x6e, 0x44, 0x65, 0x74, 0x61, 0x69, 0x6c, 0x12, 0x3c, 0x0a, 0x06, 0x61, 0x63,
	0x74, 0x6f, 0x72, 0x73, 0x18, 0x03, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x24, 0x2e, 0x67, 0x6f, 0x6f,
	0x67, 0x6c, 0x65, 0x2e, 0x61, 0x70, 0x70, 0x73, 0x2e, 0x64, 0x72, 0x69, 0x76, 0x65, 0x2e, 0x61,
	0x63, 0x74, 0x69, 0x76, 0x69, 0x74, 0x79, 0x2e, 0x76, 0x32, 0x2e, 0x41, 0x63, 0x74, 0x6f, 0x72,
	0x52, 0x06, 0x61, 0x63, 0x74, 0x6f, 0x72, 0x73, 0x12, 0x3f, 0x0a, 0x07, 0x61, 0x63, 0x74, 0x69,
	0x6f, 0x6e, 0x73, 0x18, 0x04, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x25, 0x2e, 0x67, 0x6f, 0x6f, 0x67,
	0x6c, 0x65, 0x2e, 0x61, 0x70, 0x70, 0x73, 0x2e, 0x64, 0x72, 0x69, 0x76, 0x65, 0x2e, 0x61, 0x63,
	0x74, 0x69, 0x76, 0x69, 0x74, 0x79, 0x2e, 0x76, 0x32, 0x2e, 0x41, 0x63, 0x74, 0x69, 0x6f, 0x6e,
	0x52, 0x07, 0x61, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x12, 0x3f, 0x0a, 0x07, 0x74, 0x61, 0x72,
	0x67, 0x65, 0x74, 0x73, 0x18, 0x05, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x25, 0x2e, 0x67, 0x6f, 0x6f,
	0x67, 0x6c, 0x65, 0x2e, 0x61, 0x70, 0x70, 0x73, 0x2e, 0x64, 0x72, 0x69, 0x76, 0x65, 0x2e, 0x61,
	0x63, 0x74, 0x69, 0x76, 0x69, 0x74, 0x79, 0x2e, 0x76, 0x32, 0x2e, 0x54, 0x61, 0x72, 0x67, 0x65,
	0x74, 0x52, 0x07, 0x74, 0x61, 0x72, 0x67, 0x65, 0x74, 0x73, 0x12, 0x3a, 0x0a, 0x09, 0x74, 0x69,
	0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d, 0x70, 0x18, 0x06, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1a, 0x2e,
	0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e,
	0x54, 0x69, 0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d, 0x70, 0x48, 0x00, 0x52, 0x09, 0x74, 0x69, 0x6d,
	0x65, 0x73, 0x74, 0x61, 0x6d, 0x70, 0x12, 0x49, 0x0a, 0x0a, 0x74, 0x69, 0x6d, 0x65, 0x5f, 0x72,
	0x61, 0x6e, 0x67, 0x65, 0x18, 0x07, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x28, 0x2e, 0x67, 0x6f, 0x6f,
	0x67, 0x6c, 0x65, 0x2e, 0x61, 0x70, 0x70, 0x73, 0x2e, 0x64, 0x72, 0x69, 0x76, 0x65, 0x2e, 0x61,
	0x63, 0x74, 0x69, 0x76, 0x69, 0x74, 0x79, 0x2e, 0x76, 0x32, 0x2e, 0x54, 0x69, 0x6d, 0x65, 0x52,
	0x61, 0x6e, 0x67, 0x65, 0x48, 0x00, 0x52, 0x09, 0x74, 0x69, 0x6d, 0x65, 0x52, 0x61, 0x6e, 0x67,
	0x65, 0x42, 0x06, 0x0a, 0x04, 0x74, 0x69, 0x6d, 0x65, 0x42, 0xd4, 0x01, 0x0a, 0x21, 0x63, 0x6f,
	0x6d, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x61, 0x70, 0x70, 0x73, 0x2e, 0x64, 0x72,
	0x69, 0x76, 0x65, 0x2e, 0x61, 0x63, 0x74, 0x69, 0x76, 0x69, 0x74, 0x79, 0x2e, 0x76, 0x32, 0x42,
	0x1f, 0x51, 0x75, 0x65, 0x72, 0x79, 0x44, 0x72, 0x69, 0x76, 0x65, 0x41, 0x63, 0x74, 0x69, 0x76,
	0x69, 0x74, 0x79, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x50, 0x72, 0x6f, 0x74, 0x6f,
	0x50, 0x01, 0x5a, 0x45, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x67, 0x6f, 0x6c, 0x61, 0x6e,
	0x67, 0x2e, 0x6f, 0x72, 0x67, 0x2f, 0x67, 0x65, 0x6e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2f, 0x67,
	0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x61, 0x70, 0x69, 0x73, 0x2f, 0x61, 0x70, 0x70, 0x73, 0x2f, 0x64,
	0x72, 0x69, 0x76, 0x65, 0x2f, 0x61, 0x63, 0x74, 0x69, 0x76, 0x69, 0x74, 0x79, 0x2f, 0x76, 0x32,
	0x3b, 0x61, 0x63, 0x74, 0x69, 0x76, 0x69, 0x74, 0x79, 0xa2, 0x02, 0x04, 0x47, 0x41, 0x44, 0x41,
	0xaa, 0x02, 0x1d, 0x47, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x41, 0x70, 0x70, 0x73, 0x2e, 0x44,
	0x72, 0x69, 0x76, 0x65, 0x2e, 0x41, 0x63, 0x74, 0x69, 0x76, 0x69, 0x74, 0x79, 0x2e, 0x56, 0x32,
	0xca, 0x02, 0x1d, 0x47, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x5c, 0x41, 0x70, 0x70, 0x73, 0x5c, 0x44,
	0x72, 0x69, 0x76, 0x65, 0x5c, 0x41, 0x63, 0x74, 0x69, 0x76, 0x69, 0x74, 0x79, 0x5c, 0x56, 0x32,
	0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_google_apps_drive_activity_v2_query_drive_activity_response_proto_rawDescOnce sync.Once
	file_google_apps_drive_activity_v2_query_drive_activity_response_proto_rawDescData = file_google_apps_drive_activity_v2_query_drive_activity_response_proto_rawDesc
)

func file_google_apps_drive_activity_v2_query_drive_activity_response_proto_rawDescGZIP() []byte {
	file_google_apps_drive_activity_v2_query_drive_activity_response_proto_rawDescOnce.Do(func() {
		file_google_apps_drive_activity_v2_query_drive_activity_response_proto_rawDescData = protoimpl.X.CompressGZIP(file_google_apps_drive_activity_v2_query_drive_activity_response_proto_rawDescData)
	})
	return file_google_apps_drive_activity_v2_query_drive_activity_response_proto_rawDescData
}

var file_google_apps_drive_activity_v2_query_drive_activity_response_proto_msgTypes = make([]protoimpl.MessageInfo, 2)
var file_google_apps_drive_activity_v2_query_drive_activity_response_proto_goTypes = []interface{}{
	(*QueryDriveActivityResponse)(nil), // 0: google.apps.drive.activity.v2.QueryDriveActivityResponse
	(*DriveActivity)(nil),              // 1: google.apps.drive.activity.v2.DriveActivity
	(*ActionDetail)(nil),               // 2: google.apps.drive.activity.v2.ActionDetail
	(*Actor)(nil),                      // 3: google.apps.drive.activity.v2.Actor
	(*Action)(nil),                     // 4: google.apps.drive.activity.v2.Action
	(*Target)(nil),                     // 5: google.apps.drive.activity.v2.Target
	(*timestamppb.Timestamp)(nil),      // 6: google.protobuf.Timestamp
	(*TimeRange)(nil),                  // 7: google.apps.drive.activity.v2.TimeRange
}
var file_google_apps_drive_activity_v2_query_drive_activity_response_proto_depIdxs = []int32{
	1, // 0: google.apps.drive.activity.v2.QueryDriveActivityResponse.activities:type_name -> google.apps.drive.activity.v2.DriveActivity
	2, // 1: google.apps.drive.activity.v2.DriveActivity.primary_action_detail:type_name -> google.apps.drive.activity.v2.ActionDetail
	3, // 2: google.apps.drive.activity.v2.DriveActivity.actors:type_name -> google.apps.drive.activity.v2.Actor
	4, // 3: google.apps.drive.activity.v2.DriveActivity.actions:type_name -> google.apps.drive.activity.v2.Action
	5, // 4: google.apps.drive.activity.v2.DriveActivity.targets:type_name -> google.apps.drive.activity.v2.Target
	6, // 5: google.apps.drive.activity.v2.DriveActivity.timestamp:type_name -> google.protobuf.Timestamp
	7, // 6: google.apps.drive.activity.v2.DriveActivity.time_range:type_name -> google.apps.drive.activity.v2.TimeRange
	7, // [7:7] is the sub-list for method output_type
	7, // [7:7] is the sub-list for method input_type
	7, // [7:7] is the sub-list for extension type_name
	7, // [7:7] is the sub-list for extension extendee
	0, // [0:7] is the sub-list for field type_name
}

func init() { file_google_apps_drive_activity_v2_query_drive_activity_response_proto_init() }
func file_google_apps_drive_activity_v2_query_drive_activity_response_proto_init() {
	if File_google_apps_drive_activity_v2_query_drive_activity_response_proto != nil {
		return
	}
	file_google_apps_drive_activity_v2_action_proto_init()
	file_google_apps_drive_activity_v2_actor_proto_init()
	file_google_apps_drive_activity_v2_common_proto_init()
	file_google_apps_drive_activity_v2_target_proto_init()
	if !protoimpl.UnsafeEnabled {
		file_google_apps_drive_activity_v2_query_drive_activity_response_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*QueryDriveActivityResponse); i {
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
		file_google_apps_drive_activity_v2_query_drive_activity_response_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*DriveActivity); i {
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
	file_google_apps_drive_activity_v2_query_drive_activity_response_proto_msgTypes[1].OneofWrappers = []interface{}{
		(*DriveActivity_Timestamp)(nil),
		(*DriveActivity_TimeRange)(nil),
	}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_google_apps_drive_activity_v2_query_drive_activity_response_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   2,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_google_apps_drive_activity_v2_query_drive_activity_response_proto_goTypes,
		DependencyIndexes: file_google_apps_drive_activity_v2_query_drive_activity_response_proto_depIdxs,
		MessageInfos:      file_google_apps_drive_activity_v2_query_drive_activity_response_proto_msgTypes,
	}.Build()
	File_google_apps_drive_activity_v2_query_drive_activity_response_proto = out.File
	file_google_apps_drive_activity_v2_query_drive_activity_response_proto_rawDesc = nil
	file_google_apps_drive_activity_v2_query_drive_activity_response_proto_goTypes = nil
	file_google_apps_drive_activity_v2_query_drive_activity_response_proto_depIdxs = nil
}
