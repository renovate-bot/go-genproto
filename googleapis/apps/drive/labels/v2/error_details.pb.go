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
// source: google/apps/drive/labels/v2/error_details.proto

package labels

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

// Possible reasons a field is invalid.
type InvalidArgument_FieldViolation_Reason int32

const (
	// Unknown reason.
	InvalidArgument_FieldViolation_REASON_UNSPECIFIED InvalidArgument_FieldViolation_Reason = 0
	// The referenced field is required.
	InvalidArgument_FieldViolation_FIELD_REQUIRED InvalidArgument_FieldViolation_Reason = 1
	// The referenced value was invalid.
	InvalidArgument_FieldViolation_INVALID_VALUE InvalidArgument_FieldViolation_Reason = 2
	// The specified numeric value is out of the allowed range.
	InvalidArgument_FieldViolation_VALUE_OUT_OF_RANGE InvalidArgument_FieldViolation_Reason = 3
	// The specified string value was too long.
	InvalidArgument_FieldViolation_STRING_VALUE_TOO_LONG InvalidArgument_FieldViolation_Reason = 4
	// The number of entries exceeded the maximum.
	InvalidArgument_FieldViolation_MAX_ENTRIES_EXCEEDED InvalidArgument_FieldViolation_Reason = 5
	// The specified field is not found in the Label.
	InvalidArgument_FieldViolation_FIELD_NOT_FOUND InvalidArgument_FieldViolation_Reason = 6
	// The specified choice is not found in the Field.
	InvalidArgument_FieldViolation_CHOICE_NOT_FOUND InvalidArgument_FieldViolation_Reason = 7
)

// Enum value maps for InvalidArgument_FieldViolation_Reason.
var (
	InvalidArgument_FieldViolation_Reason_name = map[int32]string{
		0: "REASON_UNSPECIFIED",
		1: "FIELD_REQUIRED",
		2: "INVALID_VALUE",
		3: "VALUE_OUT_OF_RANGE",
		4: "STRING_VALUE_TOO_LONG",
		5: "MAX_ENTRIES_EXCEEDED",
		6: "FIELD_NOT_FOUND",
		7: "CHOICE_NOT_FOUND",
	}
	InvalidArgument_FieldViolation_Reason_value = map[string]int32{
		"REASON_UNSPECIFIED":    0,
		"FIELD_REQUIRED":        1,
		"INVALID_VALUE":         2,
		"VALUE_OUT_OF_RANGE":    3,
		"STRING_VALUE_TOO_LONG": 4,
		"MAX_ENTRIES_EXCEEDED":  5,
		"FIELD_NOT_FOUND":       6,
		"CHOICE_NOT_FOUND":      7,
	}
)

func (x InvalidArgument_FieldViolation_Reason) Enum() *InvalidArgument_FieldViolation_Reason {
	p := new(InvalidArgument_FieldViolation_Reason)
	*p = x
	return p
}

func (x InvalidArgument_FieldViolation_Reason) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (InvalidArgument_FieldViolation_Reason) Descriptor() protoreflect.EnumDescriptor {
	return file_google_apps_drive_labels_v2_error_details_proto_enumTypes[0].Descriptor()
}

func (InvalidArgument_FieldViolation_Reason) Type() protoreflect.EnumType {
	return &file_google_apps_drive_labels_v2_error_details_proto_enumTypes[0]
}

func (x InvalidArgument_FieldViolation_Reason) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use InvalidArgument_FieldViolation_Reason.Descriptor instead.
func (InvalidArgument_FieldViolation_Reason) EnumDescriptor() ([]byte, []int) {
	return file_google_apps_drive_labels_v2_error_details_proto_rawDescGZIP(), []int{0, 0, 0}
}

// The possible reasons a the violation occurred.
type PreconditionFailure_Violation_Reason int32

const (
	// Unknown violation type.
	PreconditionFailure_Violation_REASON_UNSPECIFIED PreconditionFailure_Violation_Reason = 0
	// This Resource cannot be Disabled. Only Published resources can be
	// Disabled.
	PreconditionFailure_Violation_CANNOT_DISABLE PreconditionFailure_Violation_Reason = 1
	// This Resource cannot be Enabled. Only Disabled resources can be
	// Enabled.
	PreconditionFailure_Violation_CANNOT_ENABLE PreconditionFailure_Violation_Reason = 2
	// This Resource cannot be Published. Only Draft or Disabled resources
	// can be Published.
	PreconditionFailure_Violation_CANNOT_PUBLISH PreconditionFailure_Violation_Reason = 3
	// This Resource cannot be Unpublished. Once published, resources may
	// not be set in "Draft" state.
	PreconditionFailure_Violation_CANNOT_UNPUBLISH PreconditionFailure_Violation_Reason = 4
	// This Resource cannot be Deleted. Only Disabled resources
	// can be Deleted.
	PreconditionFailure_Violation_CANNOT_DELETE PreconditionFailure_Violation_Reason = 5
	// The request modified a range in a Field, but the new range does
	// not include the previous range. When this error happens, `field` points
	// at the Field where the violation occurred.
	PreconditionFailure_Violation_CANNOT_RESTRICT_RANGE PreconditionFailure_Violation_Reason = 6
	// The specified change cannot be made to published Resources.
	PreconditionFailure_Violation_CANNOT_CHANGE_PUBLISHED_FIELD PreconditionFailure_Violation_Reason = 7
	// The customer cannot create new labels because the maximum number
	// of labels for the customer has been reached.
	PreconditionFailure_Violation_CANNOT_CREATE_MORE_LABELS PreconditionFailure_Violation_Reason = 8
	// The Field type cannot be changed because the Field has been published.
	PreconditionFailure_Violation_CANNOT_CHANGE_PUBLISHED_FIELD_TYPE PreconditionFailure_Violation_Reason = 9
	// The Label component is locked and cannot be deleted
	PreconditionFailure_Violation_CANNOT_MODIFY_LOCKED_COMPONENT PreconditionFailure_Violation_Reason = 10
	// The Label cannot be enabled in the target application or applications.
	PreconditionFailure_Violation_UNSUPPORT_ENABLED_APP_SETTINGS PreconditionFailure_Violation_Reason = 11
)

// Enum value maps for PreconditionFailure_Violation_Reason.
var (
	PreconditionFailure_Violation_Reason_name = map[int32]string{
		0:  "REASON_UNSPECIFIED",
		1:  "CANNOT_DISABLE",
		2:  "CANNOT_ENABLE",
		3:  "CANNOT_PUBLISH",
		4:  "CANNOT_UNPUBLISH",
		5:  "CANNOT_DELETE",
		6:  "CANNOT_RESTRICT_RANGE",
		7:  "CANNOT_CHANGE_PUBLISHED_FIELD",
		8:  "CANNOT_CREATE_MORE_LABELS",
		9:  "CANNOT_CHANGE_PUBLISHED_FIELD_TYPE",
		10: "CANNOT_MODIFY_LOCKED_COMPONENT",
		11: "UNSUPPORT_ENABLED_APP_SETTINGS",
	}
	PreconditionFailure_Violation_Reason_value = map[string]int32{
		"REASON_UNSPECIFIED":                 0,
		"CANNOT_DISABLE":                     1,
		"CANNOT_ENABLE":                      2,
		"CANNOT_PUBLISH":                     3,
		"CANNOT_UNPUBLISH":                   4,
		"CANNOT_DELETE":                      5,
		"CANNOT_RESTRICT_RANGE":              6,
		"CANNOT_CHANGE_PUBLISHED_FIELD":      7,
		"CANNOT_CREATE_MORE_LABELS":          8,
		"CANNOT_CHANGE_PUBLISHED_FIELD_TYPE": 9,
		"CANNOT_MODIFY_LOCKED_COMPONENT":     10,
		"UNSUPPORT_ENABLED_APP_SETTINGS":     11,
	}
)

func (x PreconditionFailure_Violation_Reason) Enum() *PreconditionFailure_Violation_Reason {
	p := new(PreconditionFailure_Violation_Reason)
	*p = x
	return p
}

func (x PreconditionFailure_Violation_Reason) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (PreconditionFailure_Violation_Reason) Descriptor() protoreflect.EnumDescriptor {
	return file_google_apps_drive_labels_v2_error_details_proto_enumTypes[1].Descriptor()
}

func (PreconditionFailure_Violation_Reason) Type() protoreflect.EnumType {
	return &file_google_apps_drive_labels_v2_error_details_proto_enumTypes[1]
}

func (x PreconditionFailure_Violation_Reason) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use PreconditionFailure_Violation_Reason.Descriptor instead.
func (PreconditionFailure_Violation_Reason) EnumDescriptor() ([]byte, []int) {
	return file_google_apps_drive_labels_v2_error_details_proto_rawDescGZIP(), []int{1, 0, 0}
}

// Describes violations in a request to create or update a Label or its
// Fields.
type InvalidArgument struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// Describes all violations in a client request.
	FieldViolations []*InvalidArgument_FieldViolation `protobuf:"bytes,1,rep,name=field_violations,json=fieldViolations,proto3" json:"field_violations,omitempty"`
}

func (x *InvalidArgument) Reset() {
	*x = InvalidArgument{}
	if protoimpl.UnsafeEnabled {
		mi := &file_google_apps_drive_labels_v2_error_details_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *InvalidArgument) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*InvalidArgument) ProtoMessage() {}

func (x *InvalidArgument) ProtoReflect() protoreflect.Message {
	mi := &file_google_apps_drive_labels_v2_error_details_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use InvalidArgument.ProtoReflect.Descriptor instead.
func (*InvalidArgument) Descriptor() ([]byte, []int) {
	return file_google_apps_drive_labels_v2_error_details_proto_rawDescGZIP(), []int{0}
}

func (x *InvalidArgument) GetFieldViolations() []*InvalidArgument_FieldViolation {
	if x != nil {
		return x.FieldViolations
	}
	return nil
}

// Describes what preconditions have failed.
type PreconditionFailure struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// Describes all violations in a client request.
	Violation []*PreconditionFailure_Violation `protobuf:"bytes,1,rep,name=violation,proto3" json:"violation,omitempty"`
}

func (x *PreconditionFailure) Reset() {
	*x = PreconditionFailure{}
	if protoimpl.UnsafeEnabled {
		mi := &file_google_apps_drive_labels_v2_error_details_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *PreconditionFailure) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*PreconditionFailure) ProtoMessage() {}

func (x *PreconditionFailure) ProtoReflect() protoreflect.Message {
	mi := &file_google_apps_drive_labels_v2_error_details_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use PreconditionFailure.ProtoReflect.Descriptor instead.
func (*PreconditionFailure) Descriptor() ([]byte, []int) {
	return file_google_apps_drive_labels_v2_error_details_proto_rawDescGZIP(), []int{1}
}

func (x *PreconditionFailure) GetViolation() []*PreconditionFailure_Violation {
	if x != nil {
		return x.Violation
	}
	return nil
}

// Describes the Field in which the violation occurred.
type InvalidArgument_FieldViolation struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// The path to the field where this violation occurred. This path is
	// specified using `FieldMask` format:
	// https://developers.google.com/protocol-buffers/docs/reference/google.protobuf#fieldmask
	Field string `protobuf:"bytes,1,opt,name=field,proto3" json:"field,omitempty"`
	// The detailed reason for this FieldViolation.
	Reason InvalidArgument_FieldViolation_Reason `protobuf:"varint,2,opt,name=reason,proto3,enum=google.apps.drive.labels.v2.InvalidArgument_FieldViolation_Reason" json:"reason,omitempty"`
	// A message that describes the violation. This message is intended to
	// be shown to end users, and is localized into the requesting user's
	// preferred language.
	DisplayMessage string `protobuf:"bytes,3,opt,name=display_message,json=displayMessage,proto3" json:"display_message,omitempty"`
}

func (x *InvalidArgument_FieldViolation) Reset() {
	*x = InvalidArgument_FieldViolation{}
	if protoimpl.UnsafeEnabled {
		mi := &file_google_apps_drive_labels_v2_error_details_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *InvalidArgument_FieldViolation) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*InvalidArgument_FieldViolation) ProtoMessage() {}

func (x *InvalidArgument_FieldViolation) ProtoReflect() protoreflect.Message {
	mi := &file_google_apps_drive_labels_v2_error_details_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use InvalidArgument_FieldViolation.ProtoReflect.Descriptor instead.
func (*InvalidArgument_FieldViolation) Descriptor() ([]byte, []int) {
	return file_google_apps_drive_labels_v2_error_details_proto_rawDescGZIP(), []int{0, 0}
}

func (x *InvalidArgument_FieldViolation) GetField() string {
	if x != nil {
		return x.Field
	}
	return ""
}

func (x *InvalidArgument_FieldViolation) GetReason() InvalidArgument_FieldViolation_Reason {
	if x != nil {
		return x.Reason
	}
	return InvalidArgument_FieldViolation_REASON_UNSPECIFIED
}

func (x *InvalidArgument_FieldViolation) GetDisplayMessage() string {
	if x != nil {
		return x.DisplayMessage
	}
	return ""
}

// Specific failure reason.
type PreconditionFailure_Violation struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// The path to the field where this violation occurred. This path is
	// specified using `FieldMask` format:
	// https://developers.google.com/protocol-buffers/docs/reference/google.protobuf#fieldmask
	Field string `protobuf:"bytes,1,opt,name=field,proto3" json:"field,omitempty"`
	// The type of this violation.
	Reason PreconditionFailure_Violation_Reason `protobuf:"varint,2,opt,name=reason,proto3,enum=google.apps.drive.labels.v2.PreconditionFailure_Violation_Reason" json:"reason,omitempty"`
	// A message that describes the violation. This message is intended to
	// be shown to end users, and is localized into the requesting user's
	// preferred language.
	DisplayMessage string `protobuf:"bytes,3,opt,name=display_message,json=displayMessage,proto3" json:"display_message,omitempty"`
}

func (x *PreconditionFailure_Violation) Reset() {
	*x = PreconditionFailure_Violation{}
	if protoimpl.UnsafeEnabled {
		mi := &file_google_apps_drive_labels_v2_error_details_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *PreconditionFailure_Violation) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*PreconditionFailure_Violation) ProtoMessage() {}

func (x *PreconditionFailure_Violation) ProtoReflect() protoreflect.Message {
	mi := &file_google_apps_drive_labels_v2_error_details_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use PreconditionFailure_Violation.ProtoReflect.Descriptor instead.
func (*PreconditionFailure_Violation) Descriptor() ([]byte, []int) {
	return file_google_apps_drive_labels_v2_error_details_proto_rawDescGZIP(), []int{1, 0}
}

func (x *PreconditionFailure_Violation) GetField() string {
	if x != nil {
		return x.Field
	}
	return ""
}

func (x *PreconditionFailure_Violation) GetReason() PreconditionFailure_Violation_Reason {
	if x != nil {
		return x.Reason
	}
	return PreconditionFailure_Violation_REASON_UNSPECIFIED
}

func (x *PreconditionFailure_Violation) GetDisplayMessage() string {
	if x != nil {
		return x.DisplayMessage
	}
	return ""
}

var File_google_apps_drive_labels_v2_error_details_proto protoreflect.FileDescriptor

var file_google_apps_drive_labels_v2_error_details_proto_rawDesc = []byte{
	0x0a, 0x2f, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x61, 0x70, 0x70, 0x73, 0x2f, 0x64, 0x72,
	0x69, 0x76, 0x65, 0x2f, 0x6c, 0x61, 0x62, 0x65, 0x6c, 0x73, 0x2f, 0x76, 0x32, 0x2f, 0x65, 0x72,
	0x72, 0x6f, 0x72, 0x5f, 0x64, 0x65, 0x74, 0x61, 0x69, 0x6c, 0x73, 0x2e, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x12, 0x1b, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x61, 0x70, 0x70, 0x73, 0x2e, 0x64,
	0x72, 0x69, 0x76, 0x65, 0x2e, 0x6c, 0x61, 0x62, 0x65, 0x6c, 0x73, 0x2e, 0x76, 0x32, 0x22, 0xe9,
	0x03, 0x0a, 0x0f, 0x49, 0x6e, 0x76, 0x61, 0x6c, 0x69, 0x64, 0x41, 0x72, 0x67, 0x75, 0x6d, 0x65,
	0x6e, 0x74, 0x12, 0x66, 0x0a, 0x10, 0x66, 0x69, 0x65, 0x6c, 0x64, 0x5f, 0x76, 0x69, 0x6f, 0x6c,
	0x61, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x18, 0x01, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x3b, 0x2e, 0x67,
	0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x61, 0x70, 0x70, 0x73, 0x2e, 0x64, 0x72, 0x69, 0x76, 0x65,
	0x2e, 0x6c, 0x61, 0x62, 0x65, 0x6c, 0x73, 0x2e, 0x76, 0x32, 0x2e, 0x49, 0x6e, 0x76, 0x61, 0x6c,
	0x69, 0x64, 0x41, 0x72, 0x67, 0x75, 0x6d, 0x65, 0x6e, 0x74, 0x2e, 0x46, 0x69, 0x65, 0x6c, 0x64,
	0x56, 0x69, 0x6f, 0x6c, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x52, 0x0f, 0x66, 0x69, 0x65, 0x6c, 0x64,
	0x56, 0x69, 0x6f, 0x6c, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x1a, 0xed, 0x02, 0x0a, 0x0e, 0x46,
	0x69, 0x65, 0x6c, 0x64, 0x56, 0x69, 0x6f, 0x6c, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x12, 0x14, 0x0a,
	0x05, 0x66, 0x69, 0x65, 0x6c, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x66, 0x69,
	0x65, 0x6c, 0x64, 0x12, 0x5a, 0x0a, 0x06, 0x72, 0x65, 0x61, 0x73, 0x6f, 0x6e, 0x18, 0x02, 0x20,
	0x01, 0x28, 0x0e, 0x32, 0x42, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x61, 0x70, 0x70,
	0x73, 0x2e, 0x64, 0x72, 0x69, 0x76, 0x65, 0x2e, 0x6c, 0x61, 0x62, 0x65, 0x6c, 0x73, 0x2e, 0x76,
	0x32, 0x2e, 0x49, 0x6e, 0x76, 0x61, 0x6c, 0x69, 0x64, 0x41, 0x72, 0x67, 0x75, 0x6d, 0x65, 0x6e,
	0x74, 0x2e, 0x46, 0x69, 0x65, 0x6c, 0x64, 0x56, 0x69, 0x6f, 0x6c, 0x61, 0x74, 0x69, 0x6f, 0x6e,
	0x2e, 0x52, 0x65, 0x61, 0x73, 0x6f, 0x6e, 0x52, 0x06, 0x72, 0x65, 0x61, 0x73, 0x6f, 0x6e, 0x12,
	0x27, 0x0a, 0x0f, 0x64, 0x69, 0x73, 0x70, 0x6c, 0x61, 0x79, 0x5f, 0x6d, 0x65, 0x73, 0x73, 0x61,
	0x67, 0x65, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0e, 0x64, 0x69, 0x73, 0x70, 0x6c, 0x61,
	0x79, 0x4d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x22, 0xbf, 0x01, 0x0a, 0x06, 0x52, 0x65, 0x61,
	0x73, 0x6f, 0x6e, 0x12, 0x16, 0x0a, 0x12, 0x52, 0x45, 0x41, 0x53, 0x4f, 0x4e, 0x5f, 0x55, 0x4e,
	0x53, 0x50, 0x45, 0x43, 0x49, 0x46, 0x49, 0x45, 0x44, 0x10, 0x00, 0x12, 0x12, 0x0a, 0x0e, 0x46,
	0x49, 0x45, 0x4c, 0x44, 0x5f, 0x52, 0x45, 0x51, 0x55, 0x49, 0x52, 0x45, 0x44, 0x10, 0x01, 0x12,
	0x11, 0x0a, 0x0d, 0x49, 0x4e, 0x56, 0x41, 0x4c, 0x49, 0x44, 0x5f, 0x56, 0x41, 0x4c, 0x55, 0x45,
	0x10, 0x02, 0x12, 0x16, 0x0a, 0x12, 0x56, 0x41, 0x4c, 0x55, 0x45, 0x5f, 0x4f, 0x55, 0x54, 0x5f,
	0x4f, 0x46, 0x5f, 0x52, 0x41, 0x4e, 0x47, 0x45, 0x10, 0x03, 0x12, 0x19, 0x0a, 0x15, 0x53, 0x54,
	0x52, 0x49, 0x4e, 0x47, 0x5f, 0x56, 0x41, 0x4c, 0x55, 0x45, 0x5f, 0x54, 0x4f, 0x4f, 0x5f, 0x4c,
	0x4f, 0x4e, 0x47, 0x10, 0x04, 0x12, 0x18, 0x0a, 0x14, 0x4d, 0x41, 0x58, 0x5f, 0x45, 0x4e, 0x54,
	0x52, 0x49, 0x45, 0x53, 0x5f, 0x45, 0x58, 0x43, 0x45, 0x45, 0x44, 0x45, 0x44, 0x10, 0x05, 0x12,
	0x13, 0x0a, 0x0f, 0x46, 0x49, 0x45, 0x4c, 0x44, 0x5f, 0x4e, 0x4f, 0x54, 0x5f, 0x46, 0x4f, 0x55,
	0x4e, 0x44, 0x10, 0x06, 0x12, 0x14, 0x0a, 0x10, 0x43, 0x48, 0x4f, 0x49, 0x43, 0x45, 0x5f, 0x4e,
	0x4f, 0x54, 0x5f, 0x46, 0x4f, 0x55, 0x4e, 0x44, 0x10, 0x07, 0x22, 0xeb, 0x04, 0x0a, 0x13, 0x50,
	0x72, 0x65, 0x63, 0x6f, 0x6e, 0x64, 0x69, 0x74, 0x69, 0x6f, 0x6e, 0x46, 0x61, 0x69, 0x6c, 0x75,
	0x72, 0x65, 0x12, 0x58, 0x0a, 0x09, 0x76, 0x69, 0x6f, 0x6c, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x18,
	0x01, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x3a, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x61,
	0x70, 0x70, 0x73, 0x2e, 0x64, 0x72, 0x69, 0x76, 0x65, 0x2e, 0x6c, 0x61, 0x62, 0x65, 0x6c, 0x73,
	0x2e, 0x76, 0x32, 0x2e, 0x50, 0x72, 0x65, 0x63, 0x6f, 0x6e, 0x64, 0x69, 0x74, 0x69, 0x6f, 0x6e,
	0x46, 0x61, 0x69, 0x6c, 0x75, 0x72, 0x65, 0x2e, 0x56, 0x69, 0x6f, 0x6c, 0x61, 0x74, 0x69, 0x6f,
	0x6e, 0x52, 0x09, 0x76, 0x69, 0x6f, 0x6c, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x1a, 0xf9, 0x03, 0x0a,
	0x09, 0x56, 0x69, 0x6f, 0x6c, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x12, 0x14, 0x0a, 0x05, 0x66, 0x69,
	0x65, 0x6c, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x66, 0x69, 0x65, 0x6c, 0x64,
	0x12, 0x59, 0x0a, 0x06, 0x72, 0x65, 0x61, 0x73, 0x6f, 0x6e, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0e,
	0x32, 0x41, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x61, 0x70, 0x70, 0x73, 0x2e, 0x64,
	0x72, 0x69, 0x76, 0x65, 0x2e, 0x6c, 0x61, 0x62, 0x65, 0x6c, 0x73, 0x2e, 0x76, 0x32, 0x2e, 0x50,
	0x72, 0x65, 0x63, 0x6f, 0x6e, 0x64, 0x69, 0x74, 0x69, 0x6f, 0x6e, 0x46, 0x61, 0x69, 0x6c, 0x75,
	0x72, 0x65, 0x2e, 0x56, 0x69, 0x6f, 0x6c, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x2e, 0x52, 0x65, 0x61,
	0x73, 0x6f, 0x6e, 0x52, 0x06, 0x72, 0x65, 0x61, 0x73, 0x6f, 0x6e, 0x12, 0x27, 0x0a, 0x0f, 0x64,
	0x69, 0x73, 0x70, 0x6c, 0x61, 0x79, 0x5f, 0x6d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x18, 0x03,
	0x20, 0x01, 0x28, 0x09, 0x52, 0x0e, 0x64, 0x69, 0x73, 0x70, 0x6c, 0x61, 0x79, 0x4d, 0x65, 0x73,
	0x73, 0x61, 0x67, 0x65, 0x22, 0xd1, 0x02, 0x0a, 0x06, 0x52, 0x65, 0x61, 0x73, 0x6f, 0x6e, 0x12,
	0x16, 0x0a, 0x12, 0x52, 0x45, 0x41, 0x53, 0x4f, 0x4e, 0x5f, 0x55, 0x4e, 0x53, 0x50, 0x45, 0x43,
	0x49, 0x46, 0x49, 0x45, 0x44, 0x10, 0x00, 0x12, 0x12, 0x0a, 0x0e, 0x43, 0x41, 0x4e, 0x4e, 0x4f,
	0x54, 0x5f, 0x44, 0x49, 0x53, 0x41, 0x42, 0x4c, 0x45, 0x10, 0x01, 0x12, 0x11, 0x0a, 0x0d, 0x43,
	0x41, 0x4e, 0x4e, 0x4f, 0x54, 0x5f, 0x45, 0x4e, 0x41, 0x42, 0x4c, 0x45, 0x10, 0x02, 0x12, 0x12,
	0x0a, 0x0e, 0x43, 0x41, 0x4e, 0x4e, 0x4f, 0x54, 0x5f, 0x50, 0x55, 0x42, 0x4c, 0x49, 0x53, 0x48,
	0x10, 0x03, 0x12, 0x14, 0x0a, 0x10, 0x43, 0x41, 0x4e, 0x4e, 0x4f, 0x54, 0x5f, 0x55, 0x4e, 0x50,
	0x55, 0x42, 0x4c, 0x49, 0x53, 0x48, 0x10, 0x04, 0x12, 0x11, 0x0a, 0x0d, 0x43, 0x41, 0x4e, 0x4e,
	0x4f, 0x54, 0x5f, 0x44, 0x45, 0x4c, 0x45, 0x54, 0x45, 0x10, 0x05, 0x12, 0x19, 0x0a, 0x15, 0x43,
	0x41, 0x4e, 0x4e, 0x4f, 0x54, 0x5f, 0x52, 0x45, 0x53, 0x54, 0x52, 0x49, 0x43, 0x54, 0x5f, 0x52,
	0x41, 0x4e, 0x47, 0x45, 0x10, 0x06, 0x12, 0x21, 0x0a, 0x1d, 0x43, 0x41, 0x4e, 0x4e, 0x4f, 0x54,
	0x5f, 0x43, 0x48, 0x41, 0x4e, 0x47, 0x45, 0x5f, 0x50, 0x55, 0x42, 0x4c, 0x49, 0x53, 0x48, 0x45,
	0x44, 0x5f, 0x46, 0x49, 0x45, 0x4c, 0x44, 0x10, 0x07, 0x12, 0x1d, 0x0a, 0x19, 0x43, 0x41, 0x4e,
	0x4e, 0x4f, 0x54, 0x5f, 0x43, 0x52, 0x45, 0x41, 0x54, 0x45, 0x5f, 0x4d, 0x4f, 0x52, 0x45, 0x5f,
	0x4c, 0x41, 0x42, 0x45, 0x4c, 0x53, 0x10, 0x08, 0x12, 0x26, 0x0a, 0x22, 0x43, 0x41, 0x4e, 0x4e,
	0x4f, 0x54, 0x5f, 0x43, 0x48, 0x41, 0x4e, 0x47, 0x45, 0x5f, 0x50, 0x55, 0x42, 0x4c, 0x49, 0x53,
	0x48, 0x45, 0x44, 0x5f, 0x46, 0x49, 0x45, 0x4c, 0x44, 0x5f, 0x54, 0x59, 0x50, 0x45, 0x10, 0x09,
	0x12, 0x22, 0x0a, 0x1e, 0x43, 0x41, 0x4e, 0x4e, 0x4f, 0x54, 0x5f, 0x4d, 0x4f, 0x44, 0x49, 0x46,
	0x59, 0x5f, 0x4c, 0x4f, 0x43, 0x4b, 0x45, 0x44, 0x5f, 0x43, 0x4f, 0x4d, 0x50, 0x4f, 0x4e, 0x45,
	0x4e, 0x54, 0x10, 0x0a, 0x12, 0x22, 0x0a, 0x1e, 0x55, 0x4e, 0x53, 0x55, 0x50, 0x50, 0x4f, 0x52,
	0x54, 0x5f, 0x45, 0x4e, 0x41, 0x42, 0x4c, 0x45, 0x44, 0x5f, 0x41, 0x50, 0x50, 0x5f, 0x53, 0x45,
	0x54, 0x54, 0x49, 0x4e, 0x47, 0x53, 0x10, 0x0b, 0x42, 0x79, 0x0a, 0x1f, 0x63, 0x6f, 0x6d, 0x2e,
	0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x61, 0x70, 0x70, 0x73, 0x2e, 0x64, 0x72, 0x69, 0x76,
	0x65, 0x2e, 0x6c, 0x61, 0x62, 0x65, 0x6c, 0x73, 0x2e, 0x76, 0x32, 0x42, 0x11, 0x45, 0x72, 0x72,
	0x6f, 0x72, 0x44, 0x65, 0x74, 0x61, 0x69, 0x6c, 0x73, 0x50, 0x72, 0x6f, 0x74, 0x6f, 0x50, 0x01,
	0x5a, 0x41, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x67, 0x6f, 0x6c, 0x61, 0x6e, 0x67, 0x2e,
	0x6f, 0x72, 0x67, 0x2f, 0x67, 0x65, 0x6e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2f, 0x67, 0x6f, 0x6f,
	0x67, 0x6c, 0x65, 0x61, 0x70, 0x69, 0x73, 0x2f, 0x61, 0x70, 0x70, 0x73, 0x2f, 0x64, 0x72, 0x69,
	0x76, 0x65, 0x2f, 0x6c, 0x61, 0x62, 0x65, 0x6c, 0x73, 0x2f, 0x76, 0x32, 0x3b, 0x6c, 0x61, 0x62,
	0x65, 0x6c, 0x73, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_google_apps_drive_labels_v2_error_details_proto_rawDescOnce sync.Once
	file_google_apps_drive_labels_v2_error_details_proto_rawDescData = file_google_apps_drive_labels_v2_error_details_proto_rawDesc
)

func file_google_apps_drive_labels_v2_error_details_proto_rawDescGZIP() []byte {
	file_google_apps_drive_labels_v2_error_details_proto_rawDescOnce.Do(func() {
		file_google_apps_drive_labels_v2_error_details_proto_rawDescData = protoimpl.X.CompressGZIP(file_google_apps_drive_labels_v2_error_details_proto_rawDescData)
	})
	return file_google_apps_drive_labels_v2_error_details_proto_rawDescData
}

var file_google_apps_drive_labels_v2_error_details_proto_enumTypes = make([]protoimpl.EnumInfo, 2)
var file_google_apps_drive_labels_v2_error_details_proto_msgTypes = make([]protoimpl.MessageInfo, 4)
var file_google_apps_drive_labels_v2_error_details_proto_goTypes = []interface{}{
	(InvalidArgument_FieldViolation_Reason)(0), // 0: google.apps.drive.labels.v2.InvalidArgument.FieldViolation.Reason
	(PreconditionFailure_Violation_Reason)(0),  // 1: google.apps.drive.labels.v2.PreconditionFailure.Violation.Reason
	(*InvalidArgument)(nil),                    // 2: google.apps.drive.labels.v2.InvalidArgument
	(*PreconditionFailure)(nil),                // 3: google.apps.drive.labels.v2.PreconditionFailure
	(*InvalidArgument_FieldViolation)(nil),     // 4: google.apps.drive.labels.v2.InvalidArgument.FieldViolation
	(*PreconditionFailure_Violation)(nil),      // 5: google.apps.drive.labels.v2.PreconditionFailure.Violation
}
var file_google_apps_drive_labels_v2_error_details_proto_depIdxs = []int32{
	4, // 0: google.apps.drive.labels.v2.InvalidArgument.field_violations:type_name -> google.apps.drive.labels.v2.InvalidArgument.FieldViolation
	5, // 1: google.apps.drive.labels.v2.PreconditionFailure.violation:type_name -> google.apps.drive.labels.v2.PreconditionFailure.Violation
	0, // 2: google.apps.drive.labels.v2.InvalidArgument.FieldViolation.reason:type_name -> google.apps.drive.labels.v2.InvalidArgument.FieldViolation.Reason
	1, // 3: google.apps.drive.labels.v2.PreconditionFailure.Violation.reason:type_name -> google.apps.drive.labels.v2.PreconditionFailure.Violation.Reason
	4, // [4:4] is the sub-list for method output_type
	4, // [4:4] is the sub-list for method input_type
	4, // [4:4] is the sub-list for extension type_name
	4, // [4:4] is the sub-list for extension extendee
	0, // [0:4] is the sub-list for field type_name
}

func init() { file_google_apps_drive_labels_v2_error_details_proto_init() }
func file_google_apps_drive_labels_v2_error_details_proto_init() {
	if File_google_apps_drive_labels_v2_error_details_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_google_apps_drive_labels_v2_error_details_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*InvalidArgument); i {
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
		file_google_apps_drive_labels_v2_error_details_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*PreconditionFailure); i {
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
		file_google_apps_drive_labels_v2_error_details_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*InvalidArgument_FieldViolation); i {
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
		file_google_apps_drive_labels_v2_error_details_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*PreconditionFailure_Violation); i {
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
			RawDescriptor: file_google_apps_drive_labels_v2_error_details_proto_rawDesc,
			NumEnums:      2,
			NumMessages:   4,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_google_apps_drive_labels_v2_error_details_proto_goTypes,
		DependencyIndexes: file_google_apps_drive_labels_v2_error_details_proto_depIdxs,
		EnumInfos:         file_google_apps_drive_labels_v2_error_details_proto_enumTypes,
		MessageInfos:      file_google_apps_drive_labels_v2_error_details_proto_msgTypes,
	}.Build()
	File_google_apps_drive_labels_v2_error_details_proto = out.File
	file_google_apps_drive_labels_v2_error_details_proto_rawDesc = nil
	file_google_apps_drive_labels_v2_error_details_proto_goTypes = nil
	file_google_apps_drive_labels_v2_error_details_proto_depIdxs = nil
}
