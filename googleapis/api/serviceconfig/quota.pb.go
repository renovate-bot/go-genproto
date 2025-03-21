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
// source: google/api/quota.proto

package serviceconfig

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

// Quota configuration helps to achieve fairness and budgeting in service
// usage.
//
// The metric based quota configuration works this way:
//   - The service configuration defines a set of metrics.
//   - For API calls, the quota.metric_rules maps methods to metrics with
//     corresponding costs.
//   - The quota.limits defines limits on the metrics, which will be used for
//     quota checks at runtime.
//
// An example quota configuration in yaml format:
//
//	  quota:
//	    limits:
//
//	    - name: apiWriteQpsPerProject
//	      metric: library.googleapis.com/write_calls
//	      unit: "1/min/{project}"  # rate limit for consumer projects
//	      values:
//	        STANDARD: 10000
//
//
//	    (The metric rules bind all methods to the read_calls metric,
//	     except for the UpdateBook and DeleteBook methods. These two methods
//	     are mapped to the write_calls metric, with the UpdateBook method
//	     consuming at twice rate as the DeleteBook method.)
//	    metric_rules:
//	    - selector: "*"
//	      metric_costs:
//	        library.googleapis.com/read_calls: 1
//	    - selector: google.example.library.v1.LibraryService.UpdateBook
//	      metric_costs:
//	        library.googleapis.com/write_calls: 2
//	    - selector: google.example.library.v1.LibraryService.DeleteBook
//	      metric_costs:
//	        library.googleapis.com/write_calls: 1
//
//	Corresponding Metric definition:
//
//	    metrics:
//	    - name: library.googleapis.com/read_calls
//	      display_name: Read requests
//	      metric_kind: DELTA
//	      value_type: INT64
//
//	    - name: library.googleapis.com/write_calls
//	      display_name: Write requests
//	      metric_kind: DELTA
//	      value_type: INT64
type Quota struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// List of QuotaLimit definitions for the service.
	Limits []*QuotaLimit `protobuf:"bytes,3,rep,name=limits,proto3" json:"limits,omitempty"`
	// List of MetricRule definitions, each one mapping a selected method to one
	// or more metrics.
	MetricRules []*MetricRule `protobuf:"bytes,4,rep,name=metric_rules,json=metricRules,proto3" json:"metric_rules,omitempty"`
}

func (x *Quota) Reset() {
	*x = Quota{}
	if protoimpl.UnsafeEnabled {
		mi := &file_google_api_quota_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Quota) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Quota) ProtoMessage() {}

func (x *Quota) ProtoReflect() protoreflect.Message {
	mi := &file_google_api_quota_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Quota.ProtoReflect.Descriptor instead.
func (*Quota) Descriptor() ([]byte, []int) {
	return file_google_api_quota_proto_rawDescGZIP(), []int{0}
}

func (x *Quota) GetLimits() []*QuotaLimit {
	if x != nil {
		return x.Limits
	}
	return nil
}

func (x *Quota) GetMetricRules() []*MetricRule {
	if x != nil {
		return x.MetricRules
	}
	return nil
}

// Bind API methods to metrics. Binding a method to a metric causes that
// metric's configured quota behaviors to apply to the method call.
type MetricRule struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// Selects the methods to which this rule applies.
	//
	// Refer to [selector][google.api.DocumentationRule.selector] for syntax
	// details.
	Selector string `protobuf:"bytes,1,opt,name=selector,proto3" json:"selector,omitempty"`
	// Metrics to update when the selected methods are called, and the associated
	// cost applied to each metric.
	//
	// The key of the map is the metric name, and the values are the amount
	// increased for the metric against which the quota limits are defined.
	// The value must not be negative.
	MetricCosts map[string]int64 `protobuf:"bytes,2,rep,name=metric_costs,json=metricCosts,proto3" json:"metric_costs,omitempty" protobuf_key:"bytes,1,opt,name=key,proto3" protobuf_val:"varint,2,opt,name=value,proto3"`
}

func (x *MetricRule) Reset() {
	*x = MetricRule{}
	if protoimpl.UnsafeEnabled {
		mi := &file_google_api_quota_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *MetricRule) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*MetricRule) ProtoMessage() {}

func (x *MetricRule) ProtoReflect() protoreflect.Message {
	mi := &file_google_api_quota_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use MetricRule.ProtoReflect.Descriptor instead.
func (*MetricRule) Descriptor() ([]byte, []int) {
	return file_google_api_quota_proto_rawDescGZIP(), []int{1}
}

func (x *MetricRule) GetSelector() string {
	if x != nil {
		return x.Selector
	}
	return ""
}

func (x *MetricRule) GetMetricCosts() map[string]int64 {
	if x != nil {
		return x.MetricCosts
	}
	return nil
}

// `QuotaLimit` defines a specific limit that applies over a specified duration
// for a limit type. There can be at most one limit for a duration and limit
// type combination defined within a `QuotaGroup`.
type QuotaLimit struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// Name of the quota limit.
	//
	// The name must be provided, and it must be unique within the service. The
	// name can only include alphanumeric characters as well as '-'.
	//
	// The maximum length of the limit name is 64 characters.
	Name string `protobuf:"bytes,6,opt,name=name,proto3" json:"name,omitempty"`
	// Optional. User-visible, extended description for this quota limit.
	// Should be used only when more context is needed to understand this limit
	// than provided by the limit's display name (see: `display_name`).
	Description string `protobuf:"bytes,2,opt,name=description,proto3" json:"description,omitempty"`
	// Default number of tokens that can be consumed during the specified
	// duration. This is the number of tokens assigned when a client
	// application developer activates the service for his/her project.
	//
	// Specifying a value of 0 will block all requests. This can be used if you
	// are provisioning quota to selected consumers and blocking others.
	// Similarly, a value of -1 will indicate an unlimited quota. No other
	// negative values are allowed.
	//
	// Used by group-based quotas only.
	DefaultLimit int64 `protobuf:"varint,3,opt,name=default_limit,json=defaultLimit,proto3" json:"default_limit,omitempty"`
	// Maximum number of tokens that can be consumed during the specified
	// duration. Client application developers can override the default limit up
	// to this maximum. If specified, this value cannot be set to a value less
	// than the default limit. If not specified, it is set to the default limit.
	//
	// To allow clients to apply overrides with no upper bound, set this to -1,
	// indicating unlimited maximum quota.
	//
	// Used by group-based quotas only.
	MaxLimit int64 `protobuf:"varint,4,opt,name=max_limit,json=maxLimit,proto3" json:"max_limit,omitempty"`
	// Free tier value displayed in the Developers Console for this limit.
	// The free tier is the number of tokens that will be subtracted from the
	// billed amount when billing is enabled.
	// This field can only be set on a limit with duration "1d", in a billable
	// group; it is invalid on any other limit. If this field is not set, it
	// defaults to 0, indicating that there is no free tier for this service.
	//
	// Used by group-based quotas only.
	FreeTier int64 `protobuf:"varint,7,opt,name=free_tier,json=freeTier,proto3" json:"free_tier,omitempty"`
	// Duration of this limit in textual notation. Must be "100s" or "1d".
	//
	// Used by group-based quotas only.
	Duration string `protobuf:"bytes,5,opt,name=duration,proto3" json:"duration,omitempty"`
	// The name of the metric this quota limit applies to. The quota limits with
	// the same metric will be checked together during runtime. The metric must be
	// defined within the service config.
	Metric string `protobuf:"bytes,8,opt,name=metric,proto3" json:"metric,omitempty"`
	// Specify the unit of the quota limit. It uses the same syntax as
	// [MetricDescriptor.unit][google.api.MetricDescriptor.unit]. The supported
	// unit kinds are determined by the quota backend system.
	//
	// Here are some examples:
	// * "1/min/{project}" for quota per minute per project.
	//
	// Note: the order of unit components is insignificant.
	// The "1" at the beginning is required to follow the metric unit syntax.
	Unit string `protobuf:"bytes,9,opt,name=unit,proto3" json:"unit,omitempty"`
	// Tiered limit values. You must specify this as a key:value pair, with an
	// integer value that is the maximum number of requests allowed for the
	// specified unit. Currently only STANDARD is supported.
	Values map[string]int64 `protobuf:"bytes,10,rep,name=values,proto3" json:"values,omitempty" protobuf_key:"bytes,1,opt,name=key,proto3" protobuf_val:"varint,2,opt,name=value,proto3"`
	// User-visible display name for this limit.
	// Optional. If not set, the UI will provide a default display name based on
	// the quota configuration. This field can be used to override the default
	// display name generated from the configuration.
	DisplayName string `protobuf:"bytes,12,opt,name=display_name,json=displayName,proto3" json:"display_name,omitempty"`
}

func (x *QuotaLimit) Reset() {
	*x = QuotaLimit{}
	if protoimpl.UnsafeEnabled {
		mi := &file_google_api_quota_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *QuotaLimit) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*QuotaLimit) ProtoMessage() {}

func (x *QuotaLimit) ProtoReflect() protoreflect.Message {
	mi := &file_google_api_quota_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use QuotaLimit.ProtoReflect.Descriptor instead.
func (*QuotaLimit) Descriptor() ([]byte, []int) {
	return file_google_api_quota_proto_rawDescGZIP(), []int{2}
}

func (x *QuotaLimit) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

func (x *QuotaLimit) GetDescription() string {
	if x != nil {
		return x.Description
	}
	return ""
}

func (x *QuotaLimit) GetDefaultLimit() int64 {
	if x != nil {
		return x.DefaultLimit
	}
	return 0
}

func (x *QuotaLimit) GetMaxLimit() int64 {
	if x != nil {
		return x.MaxLimit
	}
	return 0
}

func (x *QuotaLimit) GetFreeTier() int64 {
	if x != nil {
		return x.FreeTier
	}
	return 0
}

func (x *QuotaLimit) GetDuration() string {
	if x != nil {
		return x.Duration
	}
	return ""
}

func (x *QuotaLimit) GetMetric() string {
	if x != nil {
		return x.Metric
	}
	return ""
}

func (x *QuotaLimit) GetUnit() string {
	if x != nil {
		return x.Unit
	}
	return ""
}

func (x *QuotaLimit) GetValues() map[string]int64 {
	if x != nil {
		return x.Values
	}
	return nil
}

func (x *QuotaLimit) GetDisplayName() string {
	if x != nil {
		return x.DisplayName
	}
	return ""
}

var File_google_api_quota_proto protoreflect.FileDescriptor

var file_google_api_quota_proto_rawDesc = []byte{
	0x0a, 0x16, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x61, 0x70, 0x69, 0x2f, 0x71, 0x75, 0x6f,
	0x74, 0x61, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x0a, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65,
	0x2e, 0x61, 0x70, 0x69, 0x22, 0x72, 0x0a, 0x05, 0x51, 0x75, 0x6f, 0x74, 0x61, 0x12, 0x2e, 0x0a,
	0x06, 0x6c, 0x69, 0x6d, 0x69, 0x74, 0x73, 0x18, 0x03, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x16, 0x2e,
	0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x61, 0x70, 0x69, 0x2e, 0x51, 0x75, 0x6f, 0x74, 0x61,
	0x4c, 0x69, 0x6d, 0x69, 0x74, 0x52, 0x06, 0x6c, 0x69, 0x6d, 0x69, 0x74, 0x73, 0x12, 0x39, 0x0a,
	0x0c, 0x6d, 0x65, 0x74, 0x72, 0x69, 0x63, 0x5f, 0x72, 0x75, 0x6c, 0x65, 0x73, 0x18, 0x04, 0x20,
	0x03, 0x28, 0x0b, 0x32, 0x16, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x61, 0x70, 0x69,
	0x2e, 0x4d, 0x65, 0x74, 0x72, 0x69, 0x63, 0x52, 0x75, 0x6c, 0x65, 0x52, 0x0b, 0x6d, 0x65, 0x74,
	0x72, 0x69, 0x63, 0x52, 0x75, 0x6c, 0x65, 0x73, 0x22, 0xb4, 0x01, 0x0a, 0x0a, 0x4d, 0x65, 0x74,
	0x72, 0x69, 0x63, 0x52, 0x75, 0x6c, 0x65, 0x12, 0x1a, 0x0a, 0x08, 0x73, 0x65, 0x6c, 0x65, 0x63,
	0x74, 0x6f, 0x72, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x73, 0x65, 0x6c, 0x65, 0x63,
	0x74, 0x6f, 0x72, 0x12, 0x4a, 0x0a, 0x0c, 0x6d, 0x65, 0x74, 0x72, 0x69, 0x63, 0x5f, 0x63, 0x6f,
	0x73, 0x74, 0x73, 0x18, 0x02, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x27, 0x2e, 0x67, 0x6f, 0x6f, 0x67,
	0x6c, 0x65, 0x2e, 0x61, 0x70, 0x69, 0x2e, 0x4d, 0x65, 0x74, 0x72, 0x69, 0x63, 0x52, 0x75, 0x6c,
	0x65, 0x2e, 0x4d, 0x65, 0x74, 0x72, 0x69, 0x63, 0x43, 0x6f, 0x73, 0x74, 0x73, 0x45, 0x6e, 0x74,
	0x72, 0x79, 0x52, 0x0b, 0x6d, 0x65, 0x74, 0x72, 0x69, 0x63, 0x43, 0x6f, 0x73, 0x74, 0x73, 0x1a,
	0x3e, 0x0a, 0x10, 0x4d, 0x65, 0x74, 0x72, 0x69, 0x63, 0x43, 0x6f, 0x73, 0x74, 0x73, 0x45, 0x6e,
	0x74, 0x72, 0x79, 0x12, 0x10, 0x0a, 0x03, 0x6b, 0x65, 0x79, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09,
	0x52, 0x03, 0x6b, 0x65, 0x79, 0x12, 0x14, 0x0a, 0x05, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x18, 0x02,
	0x20, 0x01, 0x28, 0x03, 0x52, 0x05, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x3a, 0x02, 0x38, 0x01, 0x22,
	0x83, 0x03, 0x0a, 0x0a, 0x51, 0x75, 0x6f, 0x74, 0x61, 0x4c, 0x69, 0x6d, 0x69, 0x74, 0x12, 0x12,
	0x0a, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x06, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x6e, 0x61,
	0x6d, 0x65, 0x12, 0x20, 0x0a, 0x0b, 0x64, 0x65, 0x73, 0x63, 0x72, 0x69, 0x70, 0x74, 0x69, 0x6f,
	0x6e, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0b, 0x64, 0x65, 0x73, 0x63, 0x72, 0x69, 0x70,
	0x74, 0x69, 0x6f, 0x6e, 0x12, 0x23, 0x0a, 0x0d, 0x64, 0x65, 0x66, 0x61, 0x75, 0x6c, 0x74, 0x5f,
	0x6c, 0x69, 0x6d, 0x69, 0x74, 0x18, 0x03, 0x20, 0x01, 0x28, 0x03, 0x52, 0x0c, 0x64, 0x65, 0x66,
	0x61, 0x75, 0x6c, 0x74, 0x4c, 0x69, 0x6d, 0x69, 0x74, 0x12, 0x1b, 0x0a, 0x09, 0x6d, 0x61, 0x78,
	0x5f, 0x6c, 0x69, 0x6d, 0x69, 0x74, 0x18, 0x04, 0x20, 0x01, 0x28, 0x03, 0x52, 0x08, 0x6d, 0x61,
	0x78, 0x4c, 0x69, 0x6d, 0x69, 0x74, 0x12, 0x1b, 0x0a, 0x09, 0x66, 0x72, 0x65, 0x65, 0x5f, 0x74,
	0x69, 0x65, 0x72, 0x18, 0x07, 0x20, 0x01, 0x28, 0x03, 0x52, 0x08, 0x66, 0x72, 0x65, 0x65, 0x54,
	0x69, 0x65, 0x72, 0x12, 0x1a, 0x0a, 0x08, 0x64, 0x75, 0x72, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x18,
	0x05, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x64, 0x75, 0x72, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x12,
	0x16, 0x0a, 0x06, 0x6d, 0x65, 0x74, 0x72, 0x69, 0x63, 0x18, 0x08, 0x20, 0x01, 0x28, 0x09, 0x52,
	0x06, 0x6d, 0x65, 0x74, 0x72, 0x69, 0x63, 0x12, 0x12, 0x0a, 0x04, 0x75, 0x6e, 0x69, 0x74, 0x18,
	0x09, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x75, 0x6e, 0x69, 0x74, 0x12, 0x3a, 0x0a, 0x06, 0x76,
	0x61, 0x6c, 0x75, 0x65, 0x73, 0x18, 0x0a, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x22, 0x2e, 0x67, 0x6f,
	0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x61, 0x70, 0x69, 0x2e, 0x51, 0x75, 0x6f, 0x74, 0x61, 0x4c, 0x69,
	0x6d, 0x69, 0x74, 0x2e, 0x56, 0x61, 0x6c, 0x75, 0x65, 0x73, 0x45, 0x6e, 0x74, 0x72, 0x79, 0x52,
	0x06, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x73, 0x12, 0x21, 0x0a, 0x0c, 0x64, 0x69, 0x73, 0x70, 0x6c,
	0x61, 0x79, 0x5f, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x0c, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0b, 0x64,
	0x69, 0x73, 0x70, 0x6c, 0x61, 0x79, 0x4e, 0x61, 0x6d, 0x65, 0x1a, 0x39, 0x0a, 0x0b, 0x56, 0x61,
	0x6c, 0x75, 0x65, 0x73, 0x45, 0x6e, 0x74, 0x72, 0x79, 0x12, 0x10, 0x0a, 0x03, 0x6b, 0x65, 0x79,
	0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x03, 0x6b, 0x65, 0x79, 0x12, 0x14, 0x0a, 0x05, 0x76,
	0x61, 0x6c, 0x75, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x03, 0x52, 0x05, 0x76, 0x61, 0x6c, 0x75,
	0x65, 0x3a, 0x02, 0x38, 0x01, 0x42, 0x6c, 0x0a, 0x0e, 0x63, 0x6f, 0x6d, 0x2e, 0x67, 0x6f, 0x6f,
	0x67, 0x6c, 0x65, 0x2e, 0x61, 0x70, 0x69, 0x42, 0x0a, 0x51, 0x75, 0x6f, 0x74, 0x61, 0x50, 0x72,
	0x6f, 0x74, 0x6f, 0x50, 0x01, 0x5a, 0x45, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x67, 0x6f,
	0x6c, 0x61, 0x6e, 0x67, 0x2e, 0x6f, 0x72, 0x67, 0x2f, 0x67, 0x65, 0x6e, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x2f, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x61, 0x70, 0x69, 0x73, 0x2f, 0x61, 0x70, 0x69,
	0x2f, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x63, 0x6f, 0x6e, 0x66, 0x69, 0x67, 0x3b, 0x73,
	0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x63, 0x6f, 0x6e, 0x66, 0x69, 0x67, 0xa2, 0x02, 0x04, 0x47,
	0x41, 0x50, 0x49, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_google_api_quota_proto_rawDescOnce sync.Once
	file_google_api_quota_proto_rawDescData = file_google_api_quota_proto_rawDesc
)

func file_google_api_quota_proto_rawDescGZIP() []byte {
	file_google_api_quota_proto_rawDescOnce.Do(func() {
		file_google_api_quota_proto_rawDescData = protoimpl.X.CompressGZIP(file_google_api_quota_proto_rawDescData)
	})
	return file_google_api_quota_proto_rawDescData
}

var file_google_api_quota_proto_msgTypes = make([]protoimpl.MessageInfo, 5)
var file_google_api_quota_proto_goTypes = []interface{}{
	(*Quota)(nil),      // 0: google.api.Quota
	(*MetricRule)(nil), // 1: google.api.MetricRule
	(*QuotaLimit)(nil), // 2: google.api.QuotaLimit
	nil,                // 3: google.api.MetricRule.MetricCostsEntry
	nil,                // 4: google.api.QuotaLimit.ValuesEntry
}
var file_google_api_quota_proto_depIdxs = []int32{
	2, // 0: google.api.Quota.limits:type_name -> google.api.QuotaLimit
	1, // 1: google.api.Quota.metric_rules:type_name -> google.api.MetricRule
	3, // 2: google.api.MetricRule.metric_costs:type_name -> google.api.MetricRule.MetricCostsEntry
	4, // 3: google.api.QuotaLimit.values:type_name -> google.api.QuotaLimit.ValuesEntry
	4, // [4:4] is the sub-list for method output_type
	4, // [4:4] is the sub-list for method input_type
	4, // [4:4] is the sub-list for extension type_name
	4, // [4:4] is the sub-list for extension extendee
	0, // [0:4] is the sub-list for field type_name
}

func init() { file_google_api_quota_proto_init() }
func file_google_api_quota_proto_init() {
	if File_google_api_quota_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_google_api_quota_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Quota); i {
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
		file_google_api_quota_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*MetricRule); i {
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
		file_google_api_quota_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*QuotaLimit); i {
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
			RawDescriptor: file_google_api_quota_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   5,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_google_api_quota_proto_goTypes,
		DependencyIndexes: file_google_api_quota_proto_depIdxs,
		MessageInfos:      file_google_api_quota_proto_msgTypes,
	}.Build()
	File_google_api_quota_proto = out.File
	file_google_api_quota_proto_rawDesc = nil
	file_google_api_quota_proto_goTypes = nil
	file_google_api_quota_proto_depIdxs = nil
}
