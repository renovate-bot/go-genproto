// Copyright 2016 Google Inc.
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
// source: google/cloud/audit/audit_log.proto

package audit

import (
	reflect "reflect"
	sync "sync"

	proto "github.com/golang/protobuf/proto"
	any "github.com/golang/protobuf/ptypes/any"
	_struct "github.com/golang/protobuf/ptypes/struct"
	_ "google.golang.org/genproto/googleapis/api/annotations"
	status "google.golang.org/genproto/googleapis/rpc/status"
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

// Common audit log format for Google Cloud Platform API operations.
type AuditLog struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// The name of the API service performing the operation. For example,
	// `"datastore.googleapis.com"`.
	ServiceName string `protobuf:"bytes,7,opt,name=service_name,json=serviceName,proto3" json:"service_name,omitempty"`
	// The name of the service method or operation.
	// For API calls, this should be the name of the API method.
	// For example,
	//
	//     "google.datastore.v1.Datastore.RunQuery"
	//     "google.logging.v1.LoggingService.DeleteLog"
	MethodName string `protobuf:"bytes,8,opt,name=method_name,json=methodName,proto3" json:"method_name,omitempty"`
	// The resource or collection that is the target of the operation.
	// The name is a scheme-less URI, not including the API service name.
	// For example:
	//
	//     "shelves/SHELF_ID/books"
	//     "shelves/SHELF_ID/books/BOOK_ID"
	ResourceName string `protobuf:"bytes,11,opt,name=resource_name,json=resourceName,proto3" json:"resource_name,omitempty"`
	// The number of items returned from a List or Query API method,
	// if applicable.
	NumResponseItems int64 `protobuf:"varint,12,opt,name=num_response_items,json=numResponseItems,proto3" json:"num_response_items,omitempty"`
	// The status of the overall operation.
	Status *status.Status `protobuf:"bytes,2,opt,name=status,proto3" json:"status,omitempty"`
	// Authentication information.
	AuthenticationInfo *AuthenticationInfo `protobuf:"bytes,3,opt,name=authentication_info,json=authenticationInfo,proto3" json:"authentication_info,omitempty"`
	// Authorization information. If there are multiple
	// resources or permissions involved, then there is
	// one AuthorizationInfo element for each {resource, permission} tuple.
	AuthorizationInfo []*AuthorizationInfo `protobuf:"bytes,9,rep,name=authorization_info,json=authorizationInfo,proto3" json:"authorization_info,omitempty"`
	// Metadata about the operation.
	RequestMetadata *RequestMetadata `protobuf:"bytes,4,opt,name=request_metadata,json=requestMetadata,proto3" json:"request_metadata,omitempty"`
	// The operation request. This may not include all request parameters,
	// such as those that are too large, privacy-sensitive, or duplicated
	// elsewhere in the log record.
	// It should never include user-generated data, such as file contents.
	// When the JSON object represented here has a proto equivalent, the proto
	// name will be indicated in the `@type` property.
	Request *_struct.Struct `protobuf:"bytes,16,opt,name=request,proto3" json:"request,omitempty"`
	// The operation response. This may not include all response elements,
	// such as those that are too large, privacy-sensitive, or duplicated
	// elsewhere in the log record.
	// It should never include user-generated data, such as file contents.
	// When the JSON object represented here has a proto equivalent, the proto
	// name will be indicated in the `@type` property.
	Response *_struct.Struct `protobuf:"bytes,17,opt,name=response,proto3" json:"response,omitempty"`
	// Other service-specific data about the request, response, and other
	// activities.
	ServiceData *any.Any `protobuf:"bytes,15,opt,name=service_data,json=serviceData,proto3" json:"service_data,omitempty"`
}

func (x *AuditLog) Reset() {
	*x = AuditLog{}
	if protoimpl.UnsafeEnabled {
		mi := &file_google_cloud_audit_audit_log_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *AuditLog) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*AuditLog) ProtoMessage() {}

func (x *AuditLog) ProtoReflect() protoreflect.Message {
	mi := &file_google_cloud_audit_audit_log_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use AuditLog.ProtoReflect.Descriptor instead.
func (*AuditLog) Descriptor() ([]byte, []int) {
	return file_google_cloud_audit_audit_log_proto_rawDescGZIP(), []int{0}
}

func (x *AuditLog) GetServiceName() string {
	if x != nil {
		return x.ServiceName
	}
	return ""
}

func (x *AuditLog) GetMethodName() string {
	if x != nil {
		return x.MethodName
	}
	return ""
}

func (x *AuditLog) GetResourceName() string {
	if x != nil {
		return x.ResourceName
	}
	return ""
}

func (x *AuditLog) GetNumResponseItems() int64 {
	if x != nil {
		return x.NumResponseItems
	}
	return 0
}

func (x *AuditLog) GetStatus() *status.Status {
	if x != nil {
		return x.Status
	}
	return nil
}

func (x *AuditLog) GetAuthenticationInfo() *AuthenticationInfo {
	if x != nil {
		return x.AuthenticationInfo
	}
	return nil
}

func (x *AuditLog) GetAuthorizationInfo() []*AuthorizationInfo {
	if x != nil {
		return x.AuthorizationInfo
	}
	return nil
}

func (x *AuditLog) GetRequestMetadata() *RequestMetadata {
	if x != nil {
		return x.RequestMetadata
	}
	return nil
}

func (x *AuditLog) GetRequest() *_struct.Struct {
	if x != nil {
		return x.Request
	}
	return nil
}

func (x *AuditLog) GetResponse() *_struct.Struct {
	if x != nil {
		return x.Response
	}
	return nil
}

func (x *AuditLog) GetServiceData() *any.Any {
	if x != nil {
		return x.ServiceData
	}
	return nil
}

// Authentication information for the operation.
type AuthenticationInfo struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// The email address of the authenticated user making the request.
	PrincipalEmail string `protobuf:"bytes,1,opt,name=principal_email,json=principalEmail,proto3" json:"principal_email,omitempty"`
}

func (x *AuthenticationInfo) Reset() {
	*x = AuthenticationInfo{}
	if protoimpl.UnsafeEnabled {
		mi := &file_google_cloud_audit_audit_log_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *AuthenticationInfo) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*AuthenticationInfo) ProtoMessage() {}

func (x *AuthenticationInfo) ProtoReflect() protoreflect.Message {
	mi := &file_google_cloud_audit_audit_log_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use AuthenticationInfo.ProtoReflect.Descriptor instead.
func (*AuthenticationInfo) Descriptor() ([]byte, []int) {
	return file_google_cloud_audit_audit_log_proto_rawDescGZIP(), []int{1}
}

func (x *AuthenticationInfo) GetPrincipalEmail() string {
	if x != nil {
		return x.PrincipalEmail
	}
	return ""
}

// Authorization information for the operation.
type AuthorizationInfo struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// The resource being accessed, as a REST-style string. For example:
	//
	//     bigquery.googlapis.com/projects/PROJECTID/datasets/DATASETID
	Resource string `protobuf:"bytes,1,opt,name=resource,proto3" json:"resource,omitempty"`
	// The required IAM permission.
	Permission string `protobuf:"bytes,2,opt,name=permission,proto3" json:"permission,omitempty"`
	// Whether or not authorization for `resource` and `permission`
	// was granted.
	Granted bool `protobuf:"varint,3,opt,name=granted,proto3" json:"granted,omitempty"`
}

func (x *AuthorizationInfo) Reset() {
	*x = AuthorizationInfo{}
	if protoimpl.UnsafeEnabled {
		mi := &file_google_cloud_audit_audit_log_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *AuthorizationInfo) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*AuthorizationInfo) ProtoMessage() {}

func (x *AuthorizationInfo) ProtoReflect() protoreflect.Message {
	mi := &file_google_cloud_audit_audit_log_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use AuthorizationInfo.ProtoReflect.Descriptor instead.
func (*AuthorizationInfo) Descriptor() ([]byte, []int) {
	return file_google_cloud_audit_audit_log_proto_rawDescGZIP(), []int{2}
}

func (x *AuthorizationInfo) GetResource() string {
	if x != nil {
		return x.Resource
	}
	return ""
}

func (x *AuthorizationInfo) GetPermission() string {
	if x != nil {
		return x.Permission
	}
	return ""
}

func (x *AuthorizationInfo) GetGranted() bool {
	if x != nil {
		return x.Granted
	}
	return false
}

// Metadata about the request.
type RequestMetadata struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// The IP address of the caller.
	CallerIp string `protobuf:"bytes,1,opt,name=caller_ip,json=callerIp,proto3" json:"caller_ip,omitempty"`
	// The user agent of the caller.
	// This information is not authenticated and should be treated accordingly.
	// For example:
	//
	// +   `google-api-python-client/1.4.0`:
	//     The request was made by the Google API client for Python.
	// +   `Cloud SDK Command Line Tool apitools-client/1.0 gcloud/0.9.62`:
	//     The request was made by the Google Cloud SDK CLI (gcloud).
	// +   `AppEngine-Google; (+http://code.google.com/appengine; appid:
	// s~my-project`:
	//     The request was made from the `my-project` App Engine app.
	CallerSuppliedUserAgent string `protobuf:"bytes,2,opt,name=caller_supplied_user_agent,json=callerSuppliedUserAgent,proto3" json:"caller_supplied_user_agent,omitempty"`
}

func (x *RequestMetadata) Reset() {
	*x = RequestMetadata{}
	if protoimpl.UnsafeEnabled {
		mi := &file_google_cloud_audit_audit_log_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *RequestMetadata) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*RequestMetadata) ProtoMessage() {}

func (x *RequestMetadata) ProtoReflect() protoreflect.Message {
	mi := &file_google_cloud_audit_audit_log_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use RequestMetadata.ProtoReflect.Descriptor instead.
func (*RequestMetadata) Descriptor() ([]byte, []int) {
	return file_google_cloud_audit_audit_log_proto_rawDescGZIP(), []int{3}
}

func (x *RequestMetadata) GetCallerIp() string {
	if x != nil {
		return x.CallerIp
	}
	return ""
}

func (x *RequestMetadata) GetCallerSuppliedUserAgent() string {
	if x != nil {
		return x.CallerSuppliedUserAgent
	}
	return ""
}

var File_google_cloud_audit_audit_log_proto protoreflect.FileDescriptor

var file_google_cloud_audit_audit_log_proto_rawDesc = []byte{
	0x0a, 0x22, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x63, 0x6c, 0x6f, 0x75, 0x64, 0x2f, 0x61,
	0x75, 0x64, 0x69, 0x74, 0x2f, 0x61, 0x75, 0x64, 0x69, 0x74, 0x5f, 0x6c, 0x6f, 0x67, 0x2e, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x12, 0x12, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x63, 0x6c, 0x6f,
	0x75, 0x64, 0x2e, 0x61, 0x75, 0x64, 0x69, 0x74, 0x1a, 0x1c, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65,
	0x2f, 0x61, 0x70, 0x69, 0x2f, 0x61, 0x6e, 0x6e, 0x6f, 0x74, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x73,
	0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x19, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2f, 0x61, 0x6e, 0x79, 0x2e, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x1a, 0x1c, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62,
	0x75, 0x66, 0x2f, 0x73, 0x74, 0x72, 0x75, 0x63, 0x74, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a,
	0x17, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x72, 0x70, 0x63, 0x2f, 0x73, 0x74, 0x61, 0x74,
	0x75, 0x73, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0xed, 0x04, 0x0a, 0x08, 0x41, 0x75, 0x64,
	0x69, 0x74, 0x4c, 0x6f, 0x67, 0x12, 0x21, 0x0a, 0x0c, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65,
	0x5f, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x07, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0b, 0x73, 0x65, 0x72,
	0x76, 0x69, 0x63, 0x65, 0x4e, 0x61, 0x6d, 0x65, 0x12, 0x1f, 0x0a, 0x0b, 0x6d, 0x65, 0x74, 0x68,
	0x6f, 0x64, 0x5f, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x08, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0a, 0x6d,
	0x65, 0x74, 0x68, 0x6f, 0x64, 0x4e, 0x61, 0x6d, 0x65, 0x12, 0x23, 0x0a, 0x0d, 0x72, 0x65, 0x73,
	0x6f, 0x75, 0x72, 0x63, 0x65, 0x5f, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x0b, 0x20, 0x01, 0x28, 0x09,
	0x52, 0x0c, 0x72, 0x65, 0x73, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x4e, 0x61, 0x6d, 0x65, 0x12, 0x2c,
	0x0a, 0x12, 0x6e, 0x75, 0x6d, 0x5f, 0x72, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x5f, 0x69,
	0x74, 0x65, 0x6d, 0x73, 0x18, 0x0c, 0x20, 0x01, 0x28, 0x03, 0x52, 0x10, 0x6e, 0x75, 0x6d, 0x52,
	0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x49, 0x74, 0x65, 0x6d, 0x73, 0x12, 0x2a, 0x0a, 0x06,
	0x73, 0x74, 0x61, 0x74, 0x75, 0x73, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x12, 0x2e, 0x67,
	0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x72, 0x70, 0x63, 0x2e, 0x53, 0x74, 0x61, 0x74, 0x75, 0x73,
	0x52, 0x06, 0x73, 0x74, 0x61, 0x74, 0x75, 0x73, 0x12, 0x57, 0x0a, 0x13, 0x61, 0x75, 0x74, 0x68,
	0x65, 0x6e, 0x74, 0x69, 0x63, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x5f, 0x69, 0x6e, 0x66, 0x6f, 0x18,
	0x03, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x26, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x63,
	0x6c, 0x6f, 0x75, 0x64, 0x2e, 0x61, 0x75, 0x64, 0x69, 0x74, 0x2e, 0x41, 0x75, 0x74, 0x68, 0x65,
	0x6e, 0x74, 0x69, 0x63, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x49, 0x6e, 0x66, 0x6f, 0x52, 0x12, 0x61,
	0x75, 0x74, 0x68, 0x65, 0x6e, 0x74, 0x69, 0x63, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x49, 0x6e, 0x66,
	0x6f, 0x12, 0x54, 0x0a, 0x12, 0x61, 0x75, 0x74, 0x68, 0x6f, 0x72, 0x69, 0x7a, 0x61, 0x74, 0x69,
	0x6f, 0x6e, 0x5f, 0x69, 0x6e, 0x66, 0x6f, 0x18, 0x09, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x25, 0x2e,
	0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x63, 0x6c, 0x6f, 0x75, 0x64, 0x2e, 0x61, 0x75, 0x64,
	0x69, 0x74, 0x2e, 0x41, 0x75, 0x74, 0x68, 0x6f, 0x72, 0x69, 0x7a, 0x61, 0x74, 0x69, 0x6f, 0x6e,
	0x49, 0x6e, 0x66, 0x6f, 0x52, 0x11, 0x61, 0x75, 0x74, 0x68, 0x6f, 0x72, 0x69, 0x7a, 0x61, 0x74,
	0x69, 0x6f, 0x6e, 0x49, 0x6e, 0x66, 0x6f, 0x12, 0x4e, 0x0a, 0x10, 0x72, 0x65, 0x71, 0x75, 0x65,
	0x73, 0x74, 0x5f, 0x6d, 0x65, 0x74, 0x61, 0x64, 0x61, 0x74, 0x61, 0x18, 0x04, 0x20, 0x01, 0x28,
	0x0b, 0x32, 0x23, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x63, 0x6c, 0x6f, 0x75, 0x64,
	0x2e, 0x61, 0x75, 0x64, 0x69, 0x74, 0x2e, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x4d, 0x65,
	0x74, 0x61, 0x64, 0x61, 0x74, 0x61, 0x52, 0x0f, 0x72, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x4d,
	0x65, 0x74, 0x61, 0x64, 0x61, 0x74, 0x61, 0x12, 0x31, 0x0a, 0x07, 0x72, 0x65, 0x71, 0x75, 0x65,
	0x73, 0x74, 0x18, 0x10, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x17, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c,
	0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x53, 0x74, 0x72, 0x75, 0x63,
	0x74, 0x52, 0x07, 0x72, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x33, 0x0a, 0x08, 0x72, 0x65,
	0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x18, 0x11, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x17, 0x2e, 0x67,
	0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x53,
	0x74, 0x72, 0x75, 0x63, 0x74, 0x52, 0x08, 0x72, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12,
	0x37, 0x0a, 0x0c, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x5f, 0x64, 0x61, 0x74, 0x61, 0x18,
	0x0f, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x14, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x41, 0x6e, 0x79, 0x52, 0x0b, 0x73, 0x65, 0x72,
	0x76, 0x69, 0x63, 0x65, 0x44, 0x61, 0x74, 0x61, 0x22, 0x3d, 0x0a, 0x12, 0x41, 0x75, 0x74, 0x68,
	0x65, 0x6e, 0x74, 0x69, 0x63, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x49, 0x6e, 0x66, 0x6f, 0x12, 0x27,
	0x0a, 0x0f, 0x70, 0x72, 0x69, 0x6e, 0x63, 0x69, 0x70, 0x61, 0x6c, 0x5f, 0x65, 0x6d, 0x61, 0x69,
	0x6c, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0e, 0x70, 0x72, 0x69, 0x6e, 0x63, 0x69, 0x70,
	0x61, 0x6c, 0x45, 0x6d, 0x61, 0x69, 0x6c, 0x22, 0x69, 0x0a, 0x11, 0x41, 0x75, 0x74, 0x68, 0x6f,
	0x72, 0x69, 0x7a, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x49, 0x6e, 0x66, 0x6f, 0x12, 0x1a, 0x0a, 0x08,
	0x72, 0x65, 0x73, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08,
	0x72, 0x65, 0x73, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x12, 0x1e, 0x0a, 0x0a, 0x70, 0x65, 0x72, 0x6d,
	0x69, 0x73, 0x73, 0x69, 0x6f, 0x6e, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0a, 0x70, 0x65,
	0x72, 0x6d, 0x69, 0x73, 0x73, 0x69, 0x6f, 0x6e, 0x12, 0x18, 0x0a, 0x07, 0x67, 0x72, 0x61, 0x6e,
	0x74, 0x65, 0x64, 0x18, 0x03, 0x20, 0x01, 0x28, 0x08, 0x52, 0x07, 0x67, 0x72, 0x61, 0x6e, 0x74,
	0x65, 0x64, 0x22, 0x6b, 0x0a, 0x0f, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x4d, 0x65, 0x74,
	0x61, 0x64, 0x61, 0x74, 0x61, 0x12, 0x1b, 0x0a, 0x09, 0x63, 0x61, 0x6c, 0x6c, 0x65, 0x72, 0x5f,
	0x69, 0x70, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x63, 0x61, 0x6c, 0x6c, 0x65, 0x72,
	0x49, 0x70, 0x12, 0x3b, 0x0a, 0x1a, 0x63, 0x61, 0x6c, 0x6c, 0x65, 0x72, 0x5f, 0x73, 0x75, 0x70,
	0x70, 0x6c, 0x69, 0x65, 0x64, 0x5f, 0x75, 0x73, 0x65, 0x72, 0x5f, 0x61, 0x67, 0x65, 0x6e, 0x74,
	0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x17, 0x63, 0x61, 0x6c, 0x6c, 0x65, 0x72, 0x53, 0x75,
	0x70, 0x70, 0x6c, 0x69, 0x65, 0x64, 0x55, 0x73, 0x65, 0x72, 0x41, 0x67, 0x65, 0x6e, 0x74, 0x42,
	0x62, 0x0a, 0x16, 0x63, 0x6f, 0x6d, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x63, 0x6c,
	0x6f, 0x75, 0x64, 0x2e, 0x61, 0x75, 0x64, 0x69, 0x74, 0x42, 0x0d, 0x41, 0x75, 0x64, 0x69, 0x74,
	0x4c, 0x6f, 0x67, 0x50, 0x72, 0x6f, 0x74, 0x6f, 0x50, 0x01, 0x5a, 0x37, 0x67, 0x6f, 0x6f, 0x67,
	0x6c, 0x65, 0x2e, 0x67, 0x6f, 0x6c, 0x61, 0x6e, 0x67, 0x2e, 0x6f, 0x72, 0x67, 0x2f, 0x67, 0x65,
	0x6e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2f, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x61, 0x70, 0x69,
	0x73, 0x2f, 0x63, 0x6c, 0x6f, 0x75, 0x64, 0x2f, 0x61, 0x75, 0x64, 0x69, 0x74, 0x3b, 0x61, 0x75,
	0x64, 0x69, 0x74, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_google_cloud_audit_audit_log_proto_rawDescOnce sync.Once
	file_google_cloud_audit_audit_log_proto_rawDescData = file_google_cloud_audit_audit_log_proto_rawDesc
)

func file_google_cloud_audit_audit_log_proto_rawDescGZIP() []byte {
	file_google_cloud_audit_audit_log_proto_rawDescOnce.Do(func() {
		file_google_cloud_audit_audit_log_proto_rawDescData = protoimpl.X.CompressGZIP(file_google_cloud_audit_audit_log_proto_rawDescData)
	})
	return file_google_cloud_audit_audit_log_proto_rawDescData
}

var file_google_cloud_audit_audit_log_proto_msgTypes = make([]protoimpl.MessageInfo, 4)
var file_google_cloud_audit_audit_log_proto_goTypes = []interface{}{
	(*AuditLog)(nil),           // 0: google.cloud.audit.AuditLog
	(*AuthenticationInfo)(nil), // 1: google.cloud.audit.AuthenticationInfo
	(*AuthorizationInfo)(nil),  // 2: google.cloud.audit.AuthorizationInfo
	(*RequestMetadata)(nil),    // 3: google.cloud.audit.RequestMetadata
	(*status.Status)(nil),      // 4: google.rpc.Status
	(*_struct.Struct)(nil),     // 5: google.protobuf.Struct
	(*any.Any)(nil),            // 6: google.protobuf.Any
}
var file_google_cloud_audit_audit_log_proto_depIdxs = []int32{
	4, // 0: google.cloud.audit.AuditLog.status:type_name -> google.rpc.Status
	1, // 1: google.cloud.audit.AuditLog.authentication_info:type_name -> google.cloud.audit.AuthenticationInfo
	2, // 2: google.cloud.audit.AuditLog.authorization_info:type_name -> google.cloud.audit.AuthorizationInfo
	3, // 3: google.cloud.audit.AuditLog.request_metadata:type_name -> google.cloud.audit.RequestMetadata
	5, // 4: google.cloud.audit.AuditLog.request:type_name -> google.protobuf.Struct
	5, // 5: google.cloud.audit.AuditLog.response:type_name -> google.protobuf.Struct
	6, // 6: google.cloud.audit.AuditLog.service_data:type_name -> google.protobuf.Any
	7, // [7:7] is the sub-list for method output_type
	7, // [7:7] is the sub-list for method input_type
	7, // [7:7] is the sub-list for extension type_name
	7, // [7:7] is the sub-list for extension extendee
	0, // [0:7] is the sub-list for field type_name
}

func init() { file_google_cloud_audit_audit_log_proto_init() }
func file_google_cloud_audit_audit_log_proto_init() {
	if File_google_cloud_audit_audit_log_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_google_cloud_audit_audit_log_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*AuditLog); i {
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
		file_google_cloud_audit_audit_log_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*AuthenticationInfo); i {
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
		file_google_cloud_audit_audit_log_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*AuthorizationInfo); i {
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
		file_google_cloud_audit_audit_log_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*RequestMetadata); i {
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
			RawDescriptor: file_google_cloud_audit_audit_log_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   4,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_google_cloud_audit_audit_log_proto_goTypes,
		DependencyIndexes: file_google_cloud_audit_audit_log_proto_depIdxs,
		MessageInfos:      file_google_cloud_audit_audit_log_proto_msgTypes,
	}.Build()
	File_google_cloud_audit_audit_log_proto = out.File
	file_google_cloud_audit_audit_log_proto_rawDesc = nil
	file_google_cloud_audit_audit_log_proto_goTypes = nil
	file_google_cloud_audit_audit_log_proto_depIdxs = nil
}
