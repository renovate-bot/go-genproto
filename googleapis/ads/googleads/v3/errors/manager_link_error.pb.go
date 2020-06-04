// Copyright 2020 Google LLC
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
// 	protoc-gen-go v1.22.0
// 	protoc        v3.12.3
// source: google/ads/googleads/v3/errors/manager_link_error.proto

package errors

import (
	reflect "reflect"
	sync "sync"

	proto "github.com/golang/protobuf/proto"
	_ "google.golang.org/genproto/googleapis/api/annotations"
	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
)

const (
	// Verify that this generated code is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(20 - protoimpl.MinVersion)
	// Verify that runtime/protoimpl is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(protoimpl.MaxVersion - 20)
)

// This is a compile-time assertion that a sufficiently up-to-date version
// of the legacy proto package is being used.
const _ = proto.ProtoPackageIsVersion4

// Enum describing possible ManagerLink errors.
type ManagerLinkErrorEnum_ManagerLinkError int32

const (
	// Enum unspecified.
	ManagerLinkErrorEnum_UNSPECIFIED ManagerLinkErrorEnum_ManagerLinkError = 0
	// The received error code is not known in this version.
	ManagerLinkErrorEnum_UNKNOWN ManagerLinkErrorEnum_ManagerLinkError = 1
	// The manager and client have incompatible account types.
	ManagerLinkErrorEnum_ACCOUNTS_NOT_COMPATIBLE_FOR_LINKING ManagerLinkErrorEnum_ManagerLinkError = 2
	// Client is already linked to too many managers.
	ManagerLinkErrorEnum_TOO_MANY_MANAGERS ManagerLinkErrorEnum_ManagerLinkError = 3
	// Manager has too many pending invitations.
	ManagerLinkErrorEnum_TOO_MANY_INVITES ManagerLinkErrorEnum_ManagerLinkError = 4
	// Client is already invited by this manager.
	ManagerLinkErrorEnum_ALREADY_INVITED_BY_THIS_MANAGER ManagerLinkErrorEnum_ManagerLinkError = 5
	// The client is already managed by this manager.
	ManagerLinkErrorEnum_ALREADY_MANAGED_BY_THIS_MANAGER ManagerLinkErrorEnum_ManagerLinkError = 6
	// Client is already managed in hierarchy.
	ManagerLinkErrorEnum_ALREADY_MANAGED_IN_HIERARCHY ManagerLinkErrorEnum_ManagerLinkError = 7
	// Manger and sub-manager to be linked have duplicate client.
	ManagerLinkErrorEnum_DUPLICATE_CHILD_FOUND ManagerLinkErrorEnum_ManagerLinkError = 8
	// Client has no active user that can access the client account.
	ManagerLinkErrorEnum_CLIENT_HAS_NO_ADMIN_USER ManagerLinkErrorEnum_ManagerLinkError = 9
	// Adding this link would exceed the maximum hierarchy depth.
	ManagerLinkErrorEnum_MAX_DEPTH_EXCEEDED ManagerLinkErrorEnum_ManagerLinkError = 10
	// Adding this link will create a cycle.
	ManagerLinkErrorEnum_CYCLE_NOT_ALLOWED ManagerLinkErrorEnum_ManagerLinkError = 11
	// Manager account has the maximum number of linked clients.
	ManagerLinkErrorEnum_TOO_MANY_ACCOUNTS ManagerLinkErrorEnum_ManagerLinkError = 12
	// Parent manager account has the maximum number of linked clients.
	ManagerLinkErrorEnum_TOO_MANY_ACCOUNTS_AT_MANAGER ManagerLinkErrorEnum_ManagerLinkError = 13
	// The account is not authorized owner.
	ManagerLinkErrorEnum_NON_OWNER_USER_CANNOT_MODIFY_LINK ManagerLinkErrorEnum_ManagerLinkError = 14
	// Your manager account is suspended, and you are no longer allowed to link
	// to clients.
	ManagerLinkErrorEnum_SUSPENDED_ACCOUNT_CANNOT_ADD_CLIENTS ManagerLinkErrorEnum_ManagerLinkError = 15
	// You are not allowed to move a client to a manager that is not under your
	// current hierarchy.
	ManagerLinkErrorEnum_CLIENT_OUTSIDE_TREE ManagerLinkErrorEnum_ManagerLinkError = 16
	// The changed status for mutate link is invalid.
	ManagerLinkErrorEnum_INVALID_STATUS_CHANGE ManagerLinkErrorEnum_ManagerLinkError = 17
	// The change for mutate link is invalid.
	ManagerLinkErrorEnum_INVALID_CHANGE ManagerLinkErrorEnum_ManagerLinkError = 18
)

// Enum value maps for ManagerLinkErrorEnum_ManagerLinkError.
var (
	ManagerLinkErrorEnum_ManagerLinkError_name = map[int32]string{
		0:  "UNSPECIFIED",
		1:  "UNKNOWN",
		2:  "ACCOUNTS_NOT_COMPATIBLE_FOR_LINKING",
		3:  "TOO_MANY_MANAGERS",
		4:  "TOO_MANY_INVITES",
		5:  "ALREADY_INVITED_BY_THIS_MANAGER",
		6:  "ALREADY_MANAGED_BY_THIS_MANAGER",
		7:  "ALREADY_MANAGED_IN_HIERARCHY",
		8:  "DUPLICATE_CHILD_FOUND",
		9:  "CLIENT_HAS_NO_ADMIN_USER",
		10: "MAX_DEPTH_EXCEEDED",
		11: "CYCLE_NOT_ALLOWED",
		12: "TOO_MANY_ACCOUNTS",
		13: "TOO_MANY_ACCOUNTS_AT_MANAGER",
		14: "NON_OWNER_USER_CANNOT_MODIFY_LINK",
		15: "SUSPENDED_ACCOUNT_CANNOT_ADD_CLIENTS",
		16: "CLIENT_OUTSIDE_TREE",
		17: "INVALID_STATUS_CHANGE",
		18: "INVALID_CHANGE",
	}
	ManagerLinkErrorEnum_ManagerLinkError_value = map[string]int32{
		"UNSPECIFIED":                          0,
		"UNKNOWN":                              1,
		"ACCOUNTS_NOT_COMPATIBLE_FOR_LINKING":  2,
		"TOO_MANY_MANAGERS":                    3,
		"TOO_MANY_INVITES":                     4,
		"ALREADY_INVITED_BY_THIS_MANAGER":      5,
		"ALREADY_MANAGED_BY_THIS_MANAGER":      6,
		"ALREADY_MANAGED_IN_HIERARCHY":         7,
		"DUPLICATE_CHILD_FOUND":                8,
		"CLIENT_HAS_NO_ADMIN_USER":             9,
		"MAX_DEPTH_EXCEEDED":                   10,
		"CYCLE_NOT_ALLOWED":                    11,
		"TOO_MANY_ACCOUNTS":                    12,
		"TOO_MANY_ACCOUNTS_AT_MANAGER":         13,
		"NON_OWNER_USER_CANNOT_MODIFY_LINK":    14,
		"SUSPENDED_ACCOUNT_CANNOT_ADD_CLIENTS": 15,
		"CLIENT_OUTSIDE_TREE":                  16,
		"INVALID_STATUS_CHANGE":                17,
		"INVALID_CHANGE":                       18,
	}
)

func (x ManagerLinkErrorEnum_ManagerLinkError) Enum() *ManagerLinkErrorEnum_ManagerLinkError {
	p := new(ManagerLinkErrorEnum_ManagerLinkError)
	*p = x
	return p
}

func (x ManagerLinkErrorEnum_ManagerLinkError) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (ManagerLinkErrorEnum_ManagerLinkError) Descriptor() protoreflect.EnumDescriptor {
	return file_google_ads_googleads_v3_errors_manager_link_error_proto_enumTypes[0].Descriptor()
}

func (ManagerLinkErrorEnum_ManagerLinkError) Type() protoreflect.EnumType {
	return &file_google_ads_googleads_v3_errors_manager_link_error_proto_enumTypes[0]
}

func (x ManagerLinkErrorEnum_ManagerLinkError) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use ManagerLinkErrorEnum_ManagerLinkError.Descriptor instead.
func (ManagerLinkErrorEnum_ManagerLinkError) EnumDescriptor() ([]byte, []int) {
	return file_google_ads_googleads_v3_errors_manager_link_error_proto_rawDescGZIP(), []int{0, 0}
}

// Container for enum describing possible ManagerLink errors.
type ManagerLinkErrorEnum struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields
}

func (x *ManagerLinkErrorEnum) Reset() {
	*x = ManagerLinkErrorEnum{}
	if protoimpl.UnsafeEnabled {
		mi := &file_google_ads_googleads_v3_errors_manager_link_error_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ManagerLinkErrorEnum) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ManagerLinkErrorEnum) ProtoMessage() {}

func (x *ManagerLinkErrorEnum) ProtoReflect() protoreflect.Message {
	mi := &file_google_ads_googleads_v3_errors_manager_link_error_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ManagerLinkErrorEnum.ProtoReflect.Descriptor instead.
func (*ManagerLinkErrorEnum) Descriptor() ([]byte, []int) {
	return file_google_ads_googleads_v3_errors_manager_link_error_proto_rawDescGZIP(), []int{0}
}

var File_google_ads_googleads_v3_errors_manager_link_error_proto protoreflect.FileDescriptor

var file_google_ads_googleads_v3_errors_manager_link_error_proto_rawDesc = []byte{
	0x0a, 0x37, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x61, 0x64, 0x73, 0x2f, 0x67, 0x6f, 0x6f,
	0x67, 0x6c, 0x65, 0x61, 0x64, 0x73, 0x2f, 0x76, 0x33, 0x2f, 0x65, 0x72, 0x72, 0x6f, 0x72, 0x73,
	0x2f, 0x6d, 0x61, 0x6e, 0x61, 0x67, 0x65, 0x72, 0x5f, 0x6c, 0x69, 0x6e, 0x6b, 0x5f, 0x65, 0x72,
	0x72, 0x6f, 0x72, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x1e, 0x67, 0x6f, 0x6f, 0x67, 0x6c,
	0x65, 0x2e, 0x61, 0x64, 0x73, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x61, 0x64, 0x73, 0x2e,
	0x76, 0x33, 0x2e, 0x65, 0x72, 0x72, 0x6f, 0x72, 0x73, 0x1a, 0x1c, 0x67, 0x6f, 0x6f, 0x67, 0x6c,
	0x65, 0x2f, 0x61, 0x70, 0x69, 0x2f, 0x61, 0x6e, 0x6e, 0x6f, 0x74, 0x61, 0x74, 0x69, 0x6f, 0x6e,
	0x73, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0xc5, 0x04, 0x0a, 0x14, 0x4d, 0x61, 0x6e, 0x61,
	0x67, 0x65, 0x72, 0x4c, 0x69, 0x6e, 0x6b, 0x45, 0x72, 0x72, 0x6f, 0x72, 0x45, 0x6e, 0x75, 0x6d,
	0x22, 0xac, 0x04, 0x0a, 0x10, 0x4d, 0x61, 0x6e, 0x61, 0x67, 0x65, 0x72, 0x4c, 0x69, 0x6e, 0x6b,
	0x45, 0x72, 0x72, 0x6f, 0x72, 0x12, 0x0f, 0x0a, 0x0b, 0x55, 0x4e, 0x53, 0x50, 0x45, 0x43, 0x49,
	0x46, 0x49, 0x45, 0x44, 0x10, 0x00, 0x12, 0x0b, 0x0a, 0x07, 0x55, 0x4e, 0x4b, 0x4e, 0x4f, 0x57,
	0x4e, 0x10, 0x01, 0x12, 0x27, 0x0a, 0x23, 0x41, 0x43, 0x43, 0x4f, 0x55, 0x4e, 0x54, 0x53, 0x5f,
	0x4e, 0x4f, 0x54, 0x5f, 0x43, 0x4f, 0x4d, 0x50, 0x41, 0x54, 0x49, 0x42, 0x4c, 0x45, 0x5f, 0x46,
	0x4f, 0x52, 0x5f, 0x4c, 0x49, 0x4e, 0x4b, 0x49, 0x4e, 0x47, 0x10, 0x02, 0x12, 0x15, 0x0a, 0x11,
	0x54, 0x4f, 0x4f, 0x5f, 0x4d, 0x41, 0x4e, 0x59, 0x5f, 0x4d, 0x41, 0x4e, 0x41, 0x47, 0x45, 0x52,
	0x53, 0x10, 0x03, 0x12, 0x14, 0x0a, 0x10, 0x54, 0x4f, 0x4f, 0x5f, 0x4d, 0x41, 0x4e, 0x59, 0x5f,
	0x49, 0x4e, 0x56, 0x49, 0x54, 0x45, 0x53, 0x10, 0x04, 0x12, 0x23, 0x0a, 0x1f, 0x41, 0x4c, 0x52,
	0x45, 0x41, 0x44, 0x59, 0x5f, 0x49, 0x4e, 0x56, 0x49, 0x54, 0x45, 0x44, 0x5f, 0x42, 0x59, 0x5f,
	0x54, 0x48, 0x49, 0x53, 0x5f, 0x4d, 0x41, 0x4e, 0x41, 0x47, 0x45, 0x52, 0x10, 0x05, 0x12, 0x23,
	0x0a, 0x1f, 0x41, 0x4c, 0x52, 0x45, 0x41, 0x44, 0x59, 0x5f, 0x4d, 0x41, 0x4e, 0x41, 0x47, 0x45,
	0x44, 0x5f, 0x42, 0x59, 0x5f, 0x54, 0x48, 0x49, 0x53, 0x5f, 0x4d, 0x41, 0x4e, 0x41, 0x47, 0x45,
	0x52, 0x10, 0x06, 0x12, 0x20, 0x0a, 0x1c, 0x41, 0x4c, 0x52, 0x45, 0x41, 0x44, 0x59, 0x5f, 0x4d,
	0x41, 0x4e, 0x41, 0x47, 0x45, 0x44, 0x5f, 0x49, 0x4e, 0x5f, 0x48, 0x49, 0x45, 0x52, 0x41, 0x52,
	0x43, 0x48, 0x59, 0x10, 0x07, 0x12, 0x19, 0x0a, 0x15, 0x44, 0x55, 0x50, 0x4c, 0x49, 0x43, 0x41,
	0x54, 0x45, 0x5f, 0x43, 0x48, 0x49, 0x4c, 0x44, 0x5f, 0x46, 0x4f, 0x55, 0x4e, 0x44, 0x10, 0x08,
	0x12, 0x1c, 0x0a, 0x18, 0x43, 0x4c, 0x49, 0x45, 0x4e, 0x54, 0x5f, 0x48, 0x41, 0x53, 0x5f, 0x4e,
	0x4f, 0x5f, 0x41, 0x44, 0x4d, 0x49, 0x4e, 0x5f, 0x55, 0x53, 0x45, 0x52, 0x10, 0x09, 0x12, 0x16,
	0x0a, 0x12, 0x4d, 0x41, 0x58, 0x5f, 0x44, 0x45, 0x50, 0x54, 0x48, 0x5f, 0x45, 0x58, 0x43, 0x45,
	0x45, 0x44, 0x45, 0x44, 0x10, 0x0a, 0x12, 0x15, 0x0a, 0x11, 0x43, 0x59, 0x43, 0x4c, 0x45, 0x5f,
	0x4e, 0x4f, 0x54, 0x5f, 0x41, 0x4c, 0x4c, 0x4f, 0x57, 0x45, 0x44, 0x10, 0x0b, 0x12, 0x15, 0x0a,
	0x11, 0x54, 0x4f, 0x4f, 0x5f, 0x4d, 0x41, 0x4e, 0x59, 0x5f, 0x41, 0x43, 0x43, 0x4f, 0x55, 0x4e,
	0x54, 0x53, 0x10, 0x0c, 0x12, 0x20, 0x0a, 0x1c, 0x54, 0x4f, 0x4f, 0x5f, 0x4d, 0x41, 0x4e, 0x59,
	0x5f, 0x41, 0x43, 0x43, 0x4f, 0x55, 0x4e, 0x54, 0x53, 0x5f, 0x41, 0x54, 0x5f, 0x4d, 0x41, 0x4e,
	0x41, 0x47, 0x45, 0x52, 0x10, 0x0d, 0x12, 0x25, 0x0a, 0x21, 0x4e, 0x4f, 0x4e, 0x5f, 0x4f, 0x57,
	0x4e, 0x45, 0x52, 0x5f, 0x55, 0x53, 0x45, 0x52, 0x5f, 0x43, 0x41, 0x4e, 0x4e, 0x4f, 0x54, 0x5f,
	0x4d, 0x4f, 0x44, 0x49, 0x46, 0x59, 0x5f, 0x4c, 0x49, 0x4e, 0x4b, 0x10, 0x0e, 0x12, 0x28, 0x0a,
	0x24, 0x53, 0x55, 0x53, 0x50, 0x45, 0x4e, 0x44, 0x45, 0x44, 0x5f, 0x41, 0x43, 0x43, 0x4f, 0x55,
	0x4e, 0x54, 0x5f, 0x43, 0x41, 0x4e, 0x4e, 0x4f, 0x54, 0x5f, 0x41, 0x44, 0x44, 0x5f, 0x43, 0x4c,
	0x49, 0x45, 0x4e, 0x54, 0x53, 0x10, 0x0f, 0x12, 0x17, 0x0a, 0x13, 0x43, 0x4c, 0x49, 0x45, 0x4e,
	0x54, 0x5f, 0x4f, 0x55, 0x54, 0x53, 0x49, 0x44, 0x45, 0x5f, 0x54, 0x52, 0x45, 0x45, 0x10, 0x10,
	0x12, 0x19, 0x0a, 0x15, 0x49, 0x4e, 0x56, 0x41, 0x4c, 0x49, 0x44, 0x5f, 0x53, 0x54, 0x41, 0x54,
	0x55, 0x53, 0x5f, 0x43, 0x48, 0x41, 0x4e, 0x47, 0x45, 0x10, 0x11, 0x12, 0x12, 0x0a, 0x0e, 0x49,
	0x4e, 0x56, 0x41, 0x4c, 0x49, 0x44, 0x5f, 0x43, 0x48, 0x41, 0x4e, 0x47, 0x45, 0x10, 0x12, 0x42,
	0xf0, 0x01, 0x0a, 0x22, 0x63, 0x6f, 0x6d, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x61,
	0x64, 0x73, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x61, 0x64, 0x73, 0x2e, 0x76, 0x33, 0x2e,
	0x65, 0x72, 0x72, 0x6f, 0x72, 0x73, 0x42, 0x15, 0x4d, 0x61, 0x6e, 0x61, 0x67, 0x65, 0x72, 0x4c,
	0x69, 0x6e, 0x6b, 0x45, 0x72, 0x72, 0x6f, 0x72, 0x50, 0x72, 0x6f, 0x74, 0x6f, 0x50, 0x01, 0x5a,
	0x44, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x67, 0x6f, 0x6c, 0x61, 0x6e, 0x67, 0x2e, 0x6f,
	0x72, 0x67, 0x2f, 0x67, 0x65, 0x6e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2f, 0x67, 0x6f, 0x6f, 0x67,
	0x6c, 0x65, 0x61, 0x70, 0x69, 0x73, 0x2f, 0x61, 0x64, 0x73, 0x2f, 0x67, 0x6f, 0x6f, 0x67, 0x6c,
	0x65, 0x61, 0x64, 0x73, 0x2f, 0x76, 0x33, 0x2f, 0x65, 0x72, 0x72, 0x6f, 0x72, 0x73, 0x3b, 0x65,
	0x72, 0x72, 0x6f, 0x72, 0x73, 0xa2, 0x02, 0x03, 0x47, 0x41, 0x41, 0xaa, 0x02, 0x1e, 0x47, 0x6f,
	0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x41, 0x64, 0x73, 0x2e, 0x47, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x41,
	0x64, 0x73, 0x2e, 0x56, 0x33, 0x2e, 0x45, 0x72, 0x72, 0x6f, 0x72, 0x73, 0xca, 0x02, 0x1e, 0x47,
	0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x5c, 0x41, 0x64, 0x73, 0x5c, 0x47, 0x6f, 0x6f, 0x67, 0x6c, 0x65,
	0x41, 0x64, 0x73, 0x5c, 0x56, 0x33, 0x5c, 0x45, 0x72, 0x72, 0x6f, 0x72, 0x73, 0xea, 0x02, 0x22,
	0x47, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x3a, 0x3a, 0x41, 0x64, 0x73, 0x3a, 0x3a, 0x47, 0x6f, 0x6f,
	0x67, 0x6c, 0x65, 0x41, 0x64, 0x73, 0x3a, 0x3a, 0x56, 0x33, 0x3a, 0x3a, 0x45, 0x72, 0x72, 0x6f,
	0x72, 0x73, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_google_ads_googleads_v3_errors_manager_link_error_proto_rawDescOnce sync.Once
	file_google_ads_googleads_v3_errors_manager_link_error_proto_rawDescData = file_google_ads_googleads_v3_errors_manager_link_error_proto_rawDesc
)

func file_google_ads_googleads_v3_errors_manager_link_error_proto_rawDescGZIP() []byte {
	file_google_ads_googleads_v3_errors_manager_link_error_proto_rawDescOnce.Do(func() {
		file_google_ads_googleads_v3_errors_manager_link_error_proto_rawDescData = protoimpl.X.CompressGZIP(file_google_ads_googleads_v3_errors_manager_link_error_proto_rawDescData)
	})
	return file_google_ads_googleads_v3_errors_manager_link_error_proto_rawDescData
}

var file_google_ads_googleads_v3_errors_manager_link_error_proto_enumTypes = make([]protoimpl.EnumInfo, 1)
var file_google_ads_googleads_v3_errors_manager_link_error_proto_msgTypes = make([]protoimpl.MessageInfo, 1)
var file_google_ads_googleads_v3_errors_manager_link_error_proto_goTypes = []interface{}{
	(ManagerLinkErrorEnum_ManagerLinkError)(0), // 0: google.ads.googleads.v3.errors.ManagerLinkErrorEnum.ManagerLinkError
	(*ManagerLinkErrorEnum)(nil),               // 1: google.ads.googleads.v3.errors.ManagerLinkErrorEnum
}
var file_google_ads_googleads_v3_errors_manager_link_error_proto_depIdxs = []int32{
	0, // [0:0] is the sub-list for method output_type
	0, // [0:0] is the sub-list for method input_type
	0, // [0:0] is the sub-list for extension type_name
	0, // [0:0] is the sub-list for extension extendee
	0, // [0:0] is the sub-list for field type_name
}

func init() { file_google_ads_googleads_v3_errors_manager_link_error_proto_init() }
func file_google_ads_googleads_v3_errors_manager_link_error_proto_init() {
	if File_google_ads_googleads_v3_errors_manager_link_error_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_google_ads_googleads_v3_errors_manager_link_error_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ManagerLinkErrorEnum); i {
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
			RawDescriptor: file_google_ads_googleads_v3_errors_manager_link_error_proto_rawDesc,
			NumEnums:      1,
			NumMessages:   1,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_google_ads_googleads_v3_errors_manager_link_error_proto_goTypes,
		DependencyIndexes: file_google_ads_googleads_v3_errors_manager_link_error_proto_depIdxs,
		EnumInfos:         file_google_ads_googleads_v3_errors_manager_link_error_proto_enumTypes,
		MessageInfos:      file_google_ads_googleads_v3_errors_manager_link_error_proto_msgTypes,
	}.Build()
	File_google_ads_googleads_v3_errors_manager_link_error_proto = out.File
	file_google_ads_googleads_v3_errors_manager_link_error_proto_rawDesc = nil
	file_google_ads_googleads_v3_errors_manager_link_error_proto_goTypes = nil
	file_google_ads_googleads_v3_errors_manager_link_error_proto_depIdxs = nil
}
