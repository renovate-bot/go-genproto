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
// source: google/apps/drive/labels/v2beta/label_lock.proto

package labels

import (
	reflect "reflect"
	sync "sync"

	_ "google.golang.org/genproto/googleapis/api/annotations"
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

// A description of a LabelLock's state.
type LabelLock_State int32

const (
	// Unknown state.
	LabelLock_STATE_UNSPECIFIED LabelLock_State = 0
	// The LabelLock is active and is being enforced by the server.
	LabelLock_ACTIVE LabelLock_State = 1
	// The LabelLock is being deleted.  The LabelLock will continue to be
	// enforced by the server until it has been fully removed.
	LabelLock_DELETING LabelLock_State = 2
)

// Enum value maps for LabelLock_State.
var (
	LabelLock_State_name = map[int32]string{
		0: "STATE_UNSPECIFIED",
		1: "ACTIVE",
		2: "DELETING",
	}
	LabelLock_State_value = map[string]int32{
		"STATE_UNSPECIFIED": 0,
		"ACTIVE":            1,
		"DELETING":          2,
	}
)

func (x LabelLock_State) Enum() *LabelLock_State {
	p := new(LabelLock_State)
	*p = x
	return p
}

func (x LabelLock_State) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (LabelLock_State) Descriptor() protoreflect.EnumDescriptor {
	return file_google_apps_drive_labels_v2beta_label_lock_proto_enumTypes[0].Descriptor()
}

func (LabelLock_State) Type() protoreflect.EnumType {
	return &file_google_apps_drive_labels_v2beta_label_lock_proto_enumTypes[0]
}

func (x LabelLock_State) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use LabelLock_State.Descriptor instead.
func (LabelLock_State) EnumDescriptor() ([]byte, []int) {
	return file_google_apps_drive_labels_v2beta_label_lock_proto_rawDescGZIP(), []int{0, 0}
}

// A Lock that can be applied to a Label, Field, or Choice.
type LabelLock struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// Output only. Resource name of this LabelLock.
	Name string `protobuf:"bytes,1,opt,name=name,proto3" json:"name,omitempty"`
	// The ID of the Field that should be locked.  Empty if the whole
	// Label should be locked.
	FieldId string `protobuf:"bytes,2,opt,name=field_id,json=fieldId,proto3" json:"field_id,omitempty"`
	// The ID of the Selection Field Choice that should be locked.  If present,
	// `field_id` must also be present.
	ChoiceId string `protobuf:"bytes,3,opt,name=choice_id,json=choiceId,proto3" json:"choice_id,omitempty"`
	// Output only. The time this LabelLock was created.
	CreateTime *timestamppb.Timestamp `protobuf:"bytes,4,opt,name=create_time,json=createTime,proto3" json:"create_time,omitempty"`
	// Output only. The user whose credentials were used to create the LabelLock.
	// This will not be present if no user was responsible for creating the
	// LabelLock.
	Creator *UserInfo `protobuf:"bytes,5,opt,name=creator,proto3" json:"creator,omitempty"`
	// Output only. A timestamp indicating when this LabelLock was scheduled for
	// deletion. This will be present only if this LabelLock is in the DELETING
	// state.
	DeleteTime *timestamppb.Timestamp `protobuf:"bytes,6,opt,name=delete_time,json=deleteTime,proto3" json:"delete_time,omitempty"`
	// Output only. The user's capabilities on this LabelLock.
	Capabilities *LabelLock_Capabilities `protobuf:"bytes,8,opt,name=capabilities,proto3" json:"capabilities,omitempty"`
	// Output only. This LabelLock's state.
	State LabelLock_State `protobuf:"varint,9,opt,name=state,proto3,enum=google.apps.drive.labels.v2beta.LabelLock_State" json:"state,omitempty"`
}

func (x *LabelLock) Reset() {
	*x = LabelLock{}
	if protoimpl.UnsafeEnabled {
		mi := &file_google_apps_drive_labels_v2beta_label_lock_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *LabelLock) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*LabelLock) ProtoMessage() {}

func (x *LabelLock) ProtoReflect() protoreflect.Message {
	mi := &file_google_apps_drive_labels_v2beta_label_lock_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use LabelLock.ProtoReflect.Descriptor instead.
func (*LabelLock) Descriptor() ([]byte, []int) {
	return file_google_apps_drive_labels_v2beta_label_lock_proto_rawDescGZIP(), []int{0}
}

func (x *LabelLock) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

func (x *LabelLock) GetFieldId() string {
	if x != nil {
		return x.FieldId
	}
	return ""
}

func (x *LabelLock) GetChoiceId() string {
	if x != nil {
		return x.ChoiceId
	}
	return ""
}

func (x *LabelLock) GetCreateTime() *timestamppb.Timestamp {
	if x != nil {
		return x.CreateTime
	}
	return nil
}

func (x *LabelLock) GetCreator() *UserInfo {
	if x != nil {
		return x.Creator
	}
	return nil
}

func (x *LabelLock) GetDeleteTime() *timestamppb.Timestamp {
	if x != nil {
		return x.DeleteTime
	}
	return nil
}

func (x *LabelLock) GetCapabilities() *LabelLock_Capabilities {
	if x != nil {
		return x.Capabilities
	}
	return nil
}

func (x *LabelLock) GetState() LabelLock_State {
	if x != nil {
		return x.State
	}
	return LabelLock_STATE_UNSPECIFIED
}

// A description of a user's capabilities on a LabelLock.
type LabelLock_Capabilities struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// True if the user is authorized to view the policy.
	CanViewPolicy bool `protobuf:"varint,1,opt,name=can_view_policy,json=canViewPolicy,proto3" json:"can_view_policy,omitempty"`
}

func (x *LabelLock_Capabilities) Reset() {
	*x = LabelLock_Capabilities{}
	if protoimpl.UnsafeEnabled {
		mi := &file_google_apps_drive_labels_v2beta_label_lock_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *LabelLock_Capabilities) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*LabelLock_Capabilities) ProtoMessage() {}

func (x *LabelLock_Capabilities) ProtoReflect() protoreflect.Message {
	mi := &file_google_apps_drive_labels_v2beta_label_lock_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use LabelLock_Capabilities.ProtoReflect.Descriptor instead.
func (*LabelLock_Capabilities) Descriptor() ([]byte, []int) {
	return file_google_apps_drive_labels_v2beta_label_lock_proto_rawDescGZIP(), []int{0, 0}
}

func (x *LabelLock_Capabilities) GetCanViewPolicy() bool {
	if x != nil {
		return x.CanViewPolicy
	}
	return false
}

var File_google_apps_drive_labels_v2beta_label_lock_proto protoreflect.FileDescriptor

var file_google_apps_drive_labels_v2beta_label_lock_proto_rawDesc = []byte{
	0x0a, 0x30, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x61, 0x70, 0x70, 0x73, 0x2f, 0x64, 0x72,
	0x69, 0x76, 0x65, 0x2f, 0x6c, 0x61, 0x62, 0x65, 0x6c, 0x73, 0x2f, 0x76, 0x32, 0x62, 0x65, 0x74,
	0x61, 0x2f, 0x6c, 0x61, 0x62, 0x65, 0x6c, 0x5f, 0x6c, 0x6f, 0x63, 0x6b, 0x2e, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x12, 0x1f, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x61, 0x70, 0x70, 0x73, 0x2e,
	0x64, 0x72, 0x69, 0x76, 0x65, 0x2e, 0x6c, 0x61, 0x62, 0x65, 0x6c, 0x73, 0x2e, 0x76, 0x32, 0x62,
	0x65, 0x74, 0x61, 0x1a, 0x1f, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x61, 0x70, 0x69, 0x2f,
	0x66, 0x69, 0x65, 0x6c, 0x64, 0x5f, 0x62, 0x65, 0x68, 0x61, 0x76, 0x69, 0x6f, 0x72, 0x2e, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x19, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x61, 0x70, 0x69,
	0x2f, 0x72, 0x65, 0x73, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a,
	0x2c, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x61, 0x70, 0x70, 0x73, 0x2f, 0x64, 0x72, 0x69,
	0x76, 0x65, 0x2f, 0x6c, 0x61, 0x62, 0x65, 0x6c, 0x73, 0x2f, 0x76, 0x32, 0x62, 0x65, 0x74, 0x61,
	0x2f, 0x63, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x1f, 0x67,
	0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2f, 0x74,
	0x69, 0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d, 0x70, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0x93,
	0x05, 0x0a, 0x09, 0x4c, 0x61, 0x62, 0x65, 0x6c, 0x4c, 0x6f, 0x63, 0x6b, 0x12, 0x17, 0x0a, 0x04,
	0x6e, 0x61, 0x6d, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x42, 0x03, 0xe0, 0x41, 0x03, 0x52,
	0x04, 0x6e, 0x61, 0x6d, 0x65, 0x12, 0x19, 0x0a, 0x08, 0x66, 0x69, 0x65, 0x6c, 0x64, 0x5f, 0x69,
	0x64, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x07, 0x66, 0x69, 0x65, 0x6c, 0x64, 0x49, 0x64,
	0x12, 0x1b, 0x0a, 0x09, 0x63, 0x68, 0x6f, 0x69, 0x63, 0x65, 0x5f, 0x69, 0x64, 0x18, 0x03, 0x20,
	0x01, 0x28, 0x09, 0x52, 0x08, 0x63, 0x68, 0x6f, 0x69, 0x63, 0x65, 0x49, 0x64, 0x12, 0x40, 0x0a,
	0x0b, 0x63, 0x72, 0x65, 0x61, 0x74, 0x65, 0x5f, 0x74, 0x69, 0x6d, 0x65, 0x18, 0x04, 0x20, 0x01,
	0x28, 0x0b, 0x32, 0x1a, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x62, 0x75, 0x66, 0x2e, 0x54, 0x69, 0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d, 0x70, 0x42, 0x03,
	0xe0, 0x41, 0x03, 0x52, 0x0a, 0x63, 0x72, 0x65, 0x61, 0x74, 0x65, 0x54, 0x69, 0x6d, 0x65, 0x12,
	0x48, 0x0a, 0x07, 0x63, 0x72, 0x65, 0x61, 0x74, 0x6f, 0x72, 0x18, 0x05, 0x20, 0x01, 0x28, 0x0b,
	0x32, 0x29, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x61, 0x70, 0x70, 0x73, 0x2e, 0x64,
	0x72, 0x69, 0x76, 0x65, 0x2e, 0x6c, 0x61, 0x62, 0x65, 0x6c, 0x73, 0x2e, 0x76, 0x32, 0x62, 0x65,
	0x74, 0x61, 0x2e, 0x55, 0x73, 0x65, 0x72, 0x49, 0x6e, 0x66, 0x6f, 0x42, 0x03, 0xe0, 0x41, 0x03,
	0x52, 0x07, 0x63, 0x72, 0x65, 0x61, 0x74, 0x6f, 0x72, 0x12, 0x40, 0x0a, 0x0b, 0x64, 0x65, 0x6c,
	0x65, 0x74, 0x65, 0x5f, 0x74, 0x69, 0x6d, 0x65, 0x18, 0x06, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1a,
	0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66,
	0x2e, 0x54, 0x69, 0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d, 0x70, 0x42, 0x03, 0xe0, 0x41, 0x03, 0x52,
	0x0a, 0x64, 0x65, 0x6c, 0x65, 0x74, 0x65, 0x54, 0x69, 0x6d, 0x65, 0x12, 0x60, 0x0a, 0x0c, 0x63,
	0x61, 0x70, 0x61, 0x62, 0x69, 0x6c, 0x69, 0x74, 0x69, 0x65, 0x73, 0x18, 0x08, 0x20, 0x01, 0x28,
	0x0b, 0x32, 0x37, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x61, 0x70, 0x70, 0x73, 0x2e,
	0x64, 0x72, 0x69, 0x76, 0x65, 0x2e, 0x6c, 0x61, 0x62, 0x65, 0x6c, 0x73, 0x2e, 0x76, 0x32, 0x62,
	0x65, 0x74, 0x61, 0x2e, 0x4c, 0x61, 0x62, 0x65, 0x6c, 0x4c, 0x6f, 0x63, 0x6b, 0x2e, 0x43, 0x61,
	0x70, 0x61, 0x62, 0x69, 0x6c, 0x69, 0x74, 0x69, 0x65, 0x73, 0x42, 0x03, 0xe0, 0x41, 0x03, 0x52,
	0x0c, 0x63, 0x61, 0x70, 0x61, 0x62, 0x69, 0x6c, 0x69, 0x74, 0x69, 0x65, 0x73, 0x12, 0x4b, 0x0a,
	0x05, 0x73, 0x74, 0x61, 0x74, 0x65, 0x18, 0x09, 0x20, 0x01, 0x28, 0x0e, 0x32, 0x30, 0x2e, 0x67,
	0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x61, 0x70, 0x70, 0x73, 0x2e, 0x64, 0x72, 0x69, 0x76, 0x65,
	0x2e, 0x6c, 0x61, 0x62, 0x65, 0x6c, 0x73, 0x2e, 0x76, 0x32, 0x62, 0x65, 0x74, 0x61, 0x2e, 0x4c,
	0x61, 0x62, 0x65, 0x6c, 0x4c, 0x6f, 0x63, 0x6b, 0x2e, 0x53, 0x74, 0x61, 0x74, 0x65, 0x42, 0x03,
	0xe0, 0x41, 0x03, 0x52, 0x05, 0x73, 0x74, 0x61, 0x74, 0x65, 0x1a, 0x36, 0x0a, 0x0c, 0x43, 0x61,
	0x70, 0x61, 0x62, 0x69, 0x6c, 0x69, 0x74, 0x69, 0x65, 0x73, 0x12, 0x26, 0x0a, 0x0f, 0x63, 0x61,
	0x6e, 0x5f, 0x76, 0x69, 0x65, 0x77, 0x5f, 0x70, 0x6f, 0x6c, 0x69, 0x63, 0x79, 0x18, 0x01, 0x20,
	0x01, 0x28, 0x08, 0x52, 0x0d, 0x63, 0x61, 0x6e, 0x56, 0x69, 0x65, 0x77, 0x50, 0x6f, 0x6c, 0x69,
	0x63, 0x79, 0x22, 0x38, 0x0a, 0x05, 0x53, 0x74, 0x61, 0x74, 0x65, 0x12, 0x15, 0x0a, 0x11, 0x53,
	0x54, 0x41, 0x54, 0x45, 0x5f, 0x55, 0x4e, 0x53, 0x50, 0x45, 0x43, 0x49, 0x46, 0x49, 0x45, 0x44,
	0x10, 0x00, 0x12, 0x0a, 0x0a, 0x06, 0x41, 0x43, 0x54, 0x49, 0x56, 0x45, 0x10, 0x01, 0x12, 0x0c,
	0x0a, 0x08, 0x44, 0x45, 0x4c, 0x45, 0x54, 0x49, 0x4e, 0x47, 0x10, 0x02, 0x3a, 0x46, 0xea, 0x41,
	0x43, 0x0a, 0x24, 0x64, 0x72, 0x69, 0x76, 0x65, 0x6c, 0x61, 0x62, 0x65, 0x6c, 0x73, 0x2e, 0x67,
	0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x61, 0x70, 0x69, 0x73, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x4c, 0x61,
	0x62, 0x65, 0x6c, 0x4c, 0x6f, 0x63, 0x6b, 0x12, 0x1b, 0x6c, 0x61, 0x62, 0x65, 0x6c, 0x73, 0x2f,
	0x7b, 0x6c, 0x61, 0x62, 0x65, 0x6c, 0x7d, 0x2f, 0x6c, 0x6f, 0x63, 0x6b, 0x73, 0x2f, 0x7b, 0x6c,
	0x6f, 0x63, 0x6b, 0x7d, 0x42, 0x85, 0x01, 0x0a, 0x23, 0x63, 0x6f, 0x6d, 0x2e, 0x67, 0x6f, 0x6f,
	0x67, 0x6c, 0x65, 0x2e, 0x61, 0x70, 0x70, 0x73, 0x2e, 0x64, 0x72, 0x69, 0x76, 0x65, 0x2e, 0x6c,
	0x61, 0x62, 0x65, 0x6c, 0x73, 0x2e, 0x76, 0x32, 0x62, 0x65, 0x74, 0x61, 0x42, 0x0e, 0x4c, 0x61,
	0x62, 0x65, 0x6c, 0x4c, 0x6f, 0x63, 0x6b, 0x50, 0x72, 0x6f, 0x74, 0x6f, 0x50, 0x01, 0x5a, 0x45,
	0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x67, 0x6f, 0x6c, 0x61, 0x6e, 0x67, 0x2e, 0x6f, 0x72,
	0x67, 0x2f, 0x67, 0x65, 0x6e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2f, 0x67, 0x6f, 0x6f, 0x67, 0x6c,
	0x65, 0x61, 0x70, 0x69, 0x73, 0x2f, 0x61, 0x70, 0x70, 0x73, 0x2f, 0x64, 0x72, 0x69, 0x76, 0x65,
	0x2f, 0x6c, 0x61, 0x62, 0x65, 0x6c, 0x73, 0x2f, 0x76, 0x32, 0x62, 0x65, 0x74, 0x61, 0x3b, 0x6c,
	0x61, 0x62, 0x65, 0x6c, 0x73, 0xa2, 0x02, 0x04, 0x44, 0x4c, 0x42, 0x4c, 0x62, 0x06, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_google_apps_drive_labels_v2beta_label_lock_proto_rawDescOnce sync.Once
	file_google_apps_drive_labels_v2beta_label_lock_proto_rawDescData = file_google_apps_drive_labels_v2beta_label_lock_proto_rawDesc
)

func file_google_apps_drive_labels_v2beta_label_lock_proto_rawDescGZIP() []byte {
	file_google_apps_drive_labels_v2beta_label_lock_proto_rawDescOnce.Do(func() {
		file_google_apps_drive_labels_v2beta_label_lock_proto_rawDescData = protoimpl.X.CompressGZIP(file_google_apps_drive_labels_v2beta_label_lock_proto_rawDescData)
	})
	return file_google_apps_drive_labels_v2beta_label_lock_proto_rawDescData
}

var file_google_apps_drive_labels_v2beta_label_lock_proto_enumTypes = make([]protoimpl.EnumInfo, 1)
var file_google_apps_drive_labels_v2beta_label_lock_proto_msgTypes = make([]protoimpl.MessageInfo, 2)
var file_google_apps_drive_labels_v2beta_label_lock_proto_goTypes = []interface{}{
	(LabelLock_State)(0),           // 0: google.apps.drive.labels.v2beta.LabelLock.State
	(*LabelLock)(nil),              // 1: google.apps.drive.labels.v2beta.LabelLock
	(*LabelLock_Capabilities)(nil), // 2: google.apps.drive.labels.v2beta.LabelLock.Capabilities
	(*timestamppb.Timestamp)(nil),  // 3: google.protobuf.Timestamp
	(*UserInfo)(nil),               // 4: google.apps.drive.labels.v2beta.UserInfo
}
var file_google_apps_drive_labels_v2beta_label_lock_proto_depIdxs = []int32{
	3, // 0: google.apps.drive.labels.v2beta.LabelLock.create_time:type_name -> google.protobuf.Timestamp
	4, // 1: google.apps.drive.labels.v2beta.LabelLock.creator:type_name -> google.apps.drive.labels.v2beta.UserInfo
	3, // 2: google.apps.drive.labels.v2beta.LabelLock.delete_time:type_name -> google.protobuf.Timestamp
	2, // 3: google.apps.drive.labels.v2beta.LabelLock.capabilities:type_name -> google.apps.drive.labels.v2beta.LabelLock.Capabilities
	0, // 4: google.apps.drive.labels.v2beta.LabelLock.state:type_name -> google.apps.drive.labels.v2beta.LabelLock.State
	5, // [5:5] is the sub-list for method output_type
	5, // [5:5] is the sub-list for method input_type
	5, // [5:5] is the sub-list for extension type_name
	5, // [5:5] is the sub-list for extension extendee
	0, // [0:5] is the sub-list for field type_name
}

func init() { file_google_apps_drive_labels_v2beta_label_lock_proto_init() }
func file_google_apps_drive_labels_v2beta_label_lock_proto_init() {
	if File_google_apps_drive_labels_v2beta_label_lock_proto != nil {
		return
	}
	file_google_apps_drive_labels_v2beta_common_proto_init()
	if !protoimpl.UnsafeEnabled {
		file_google_apps_drive_labels_v2beta_label_lock_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*LabelLock); i {
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
		file_google_apps_drive_labels_v2beta_label_lock_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*LabelLock_Capabilities); i {
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
			RawDescriptor: file_google_apps_drive_labels_v2beta_label_lock_proto_rawDesc,
			NumEnums:      1,
			NumMessages:   2,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_google_apps_drive_labels_v2beta_label_lock_proto_goTypes,
		DependencyIndexes: file_google_apps_drive_labels_v2beta_label_lock_proto_depIdxs,
		EnumInfos:         file_google_apps_drive_labels_v2beta_label_lock_proto_enumTypes,
		MessageInfos:      file_google_apps_drive_labels_v2beta_label_lock_proto_msgTypes,
	}.Build()
	File_google_apps_drive_labels_v2beta_label_lock_proto = out.File
	file_google_apps_drive_labels_v2beta_label_lock_proto_rawDesc = nil
	file_google_apps_drive_labels_v2beta_label_lock_proto_goTypes = nil
	file_google_apps_drive_labels_v2beta_label_lock_proto_depIdxs = nil
}
