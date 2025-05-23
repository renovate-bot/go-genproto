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
// source: google/apps/script/type/script_manifest.proto

package _type

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

// Authorization header sent in add-on HTTP requests
type HttpAuthorizationHeader int32

const (
	// Default value, equivalent to `SYSTEM_ID_TOKEN`
	HttpAuthorizationHeader_HTTP_AUTHORIZATION_HEADER_UNSPECIFIED HttpAuthorizationHeader = 0
	// Send an ID token for the project-specific Google Workspace add-ons system
	// service account (default)
	HttpAuthorizationHeader_SYSTEM_ID_TOKEN HttpAuthorizationHeader = 1
	// Send an ID token for the end user
	HttpAuthorizationHeader_USER_ID_TOKEN HttpAuthorizationHeader = 2
	// Do not send an Authentication header
	HttpAuthorizationHeader_NONE HttpAuthorizationHeader = 3
)

// Enum value maps for HttpAuthorizationHeader.
var (
	HttpAuthorizationHeader_name = map[int32]string{
		0: "HTTP_AUTHORIZATION_HEADER_UNSPECIFIED",
		1: "SYSTEM_ID_TOKEN",
		2: "USER_ID_TOKEN",
		3: "NONE",
	}
	HttpAuthorizationHeader_value = map[string]int32{
		"HTTP_AUTHORIZATION_HEADER_UNSPECIFIED": 0,
		"SYSTEM_ID_TOKEN":                       1,
		"USER_ID_TOKEN":                         2,
		"NONE":                                  3,
	}
)

func (x HttpAuthorizationHeader) Enum() *HttpAuthorizationHeader {
	p := new(HttpAuthorizationHeader)
	*p = x
	return p
}

func (x HttpAuthorizationHeader) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (HttpAuthorizationHeader) Descriptor() protoreflect.EnumDescriptor {
	return file_google_apps_script_type_script_manifest_proto_enumTypes[0].Descriptor()
}

func (HttpAuthorizationHeader) Type() protoreflect.EnumType {
	return &file_google_apps_script_type_script_manifest_proto_enumTypes[0]
}

func (x HttpAuthorizationHeader) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use HttpAuthorizationHeader.Descriptor instead.
func (HttpAuthorizationHeader) EnumDescriptor() ([]byte, []int) {
	return file_google_apps_script_type_script_manifest_proto_rawDescGZIP(), []int{0}
}

// Add-on configuration that is shared across all add-on host applications.
type CommonAddOnManifest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// Required. The display name of the add-on.
	Name string `protobuf:"bytes,1,opt,name=name,proto3" json:"name,omitempty"`
	// Required. The URL for the logo image shown in the add-on toolbar.
	LogoUrl string `protobuf:"bytes,2,opt,name=logo_url,json=logoUrl,proto3" json:"logo_url,omitempty"`
	// Common layout properties for the add-on cards.
	LayoutProperties *LayoutProperties `protobuf:"bytes,3,opt,name=layout_properties,json=layoutProperties,proto3" json:"layout_properties,omitempty"`
	// The widgets used in the add-on. If this field is not specified,
	// it indicates that default set is used.
	AddOnWidgetSet *AddOnWidgetSet `protobuf:"bytes,4,opt,name=add_on_widget_set,json=addOnWidgetSet,proto3" json:"add_on_widget_set,omitempty"`
	// Whether to pass locale information from host app.
	UseLocaleFromApp bool `protobuf:"varint,5,opt,name=use_locale_from_app,json=useLocaleFromApp,proto3" json:"use_locale_from_app,omitempty"`
	// Defines an endpoint that will be executed in any context, in
	// any host. Any cards generated by this function will always be available to
	// the user, but may be eclipsed by contextual content when this add-on
	// declares more targeted triggers.
	HomepageTrigger *HomepageExtensionPoint `protobuf:"bytes,6,opt,name=homepage_trigger,json=homepageTrigger,proto3" json:"homepage_trigger,omitempty"`
	// Defines a list of extension points in the universal action menu which
	// serves as a setting menu for the add-on. The extension point can be
	// link URL to open or an endpoint to execute as a form
	// submission.
	UniversalActions []*UniversalActionExtensionPoint `protobuf:"bytes,7,rep,name=universal_actions,json=universalActions,proto3" json:"universal_actions,omitempty"`
	// An OpenLink action
	// can only use a URL with an HTTPS, MAILTO or TEL scheme.  For HTTPS links,
	// the URL must also
	// [match](/gmail/add-ons/concepts/manifests#whitelisting_urls) one of the
	// prefixes specified in this whitelist. If the prefix omits the scheme, HTTPS
	// is assumed.  Notice that HTTP links are automatically rewritten to HTTPS
	// links.
	OpenLinkUrlPrefixes *structpb.ListValue `protobuf:"bytes,8,opt,name=open_link_url_prefixes,json=openLinkUrlPrefixes,proto3" json:"open_link_url_prefixes,omitempty"`
}

func (x *CommonAddOnManifest) Reset() {
	*x = CommonAddOnManifest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_google_apps_script_type_script_manifest_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *CommonAddOnManifest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CommonAddOnManifest) ProtoMessage() {}

func (x *CommonAddOnManifest) ProtoReflect() protoreflect.Message {
	mi := &file_google_apps_script_type_script_manifest_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CommonAddOnManifest.ProtoReflect.Descriptor instead.
func (*CommonAddOnManifest) Descriptor() ([]byte, []int) {
	return file_google_apps_script_type_script_manifest_proto_rawDescGZIP(), []int{0}
}

func (x *CommonAddOnManifest) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

func (x *CommonAddOnManifest) GetLogoUrl() string {
	if x != nil {
		return x.LogoUrl
	}
	return ""
}

func (x *CommonAddOnManifest) GetLayoutProperties() *LayoutProperties {
	if x != nil {
		return x.LayoutProperties
	}
	return nil
}

func (x *CommonAddOnManifest) GetAddOnWidgetSet() *AddOnWidgetSet {
	if x != nil {
		return x.AddOnWidgetSet
	}
	return nil
}

func (x *CommonAddOnManifest) GetUseLocaleFromApp() bool {
	if x != nil {
		return x.UseLocaleFromApp
	}
	return false
}

func (x *CommonAddOnManifest) GetHomepageTrigger() *HomepageExtensionPoint {
	if x != nil {
		return x.HomepageTrigger
	}
	return nil
}

func (x *CommonAddOnManifest) GetUniversalActions() []*UniversalActionExtensionPoint {
	if x != nil {
		return x.UniversalActions
	}
	return nil
}

func (x *CommonAddOnManifest) GetOpenLinkUrlPrefixes() *structpb.ListValue {
	if x != nil {
		return x.OpenLinkUrlPrefixes
	}
	return nil
}

// Card layout properties shared across all add-on host applications.
type LayoutProperties struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// The primary color of the add-on. It sets the color of toolbar. If no
	// primary color is set explicitly, the default value provided by the
	// framework is used.
	PrimaryColor string `protobuf:"bytes,1,opt,name=primary_color,json=primaryColor,proto3" json:"primary_color,omitempty"`
	// The secondary color of the add-on. It sets the color of buttons.
	// If primary color is set but no secondary color is set, the
	// secondary color is the same as the primary color. If neither primary
	// color nor secondary color is set, the default value provided by the
	// framework is used.
	SecondaryColor string `protobuf:"bytes,2,opt,name=secondary_color,json=secondaryColor,proto3" json:"secondary_color,omitempty"`
}

func (x *LayoutProperties) Reset() {
	*x = LayoutProperties{}
	if protoimpl.UnsafeEnabled {
		mi := &file_google_apps_script_type_script_manifest_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *LayoutProperties) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*LayoutProperties) ProtoMessage() {}

func (x *LayoutProperties) ProtoReflect() protoreflect.Message {
	mi := &file_google_apps_script_type_script_manifest_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use LayoutProperties.ProtoReflect.Descriptor instead.
func (*LayoutProperties) Descriptor() ([]byte, []int) {
	return file_google_apps_script_type_script_manifest_proto_rawDescGZIP(), []int{1}
}

func (x *LayoutProperties) GetPrimaryColor() string {
	if x != nil {
		return x.PrimaryColor
	}
	return ""
}

func (x *LayoutProperties) GetSecondaryColor() string {
	if x != nil {
		return x.SecondaryColor
	}
	return ""
}

// Options for sending requests to add-on HTTP endpoints
type HttpOptions struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// Configuration for the token sent in the HTTP Authorization header
	AuthorizationHeader HttpAuthorizationHeader `protobuf:"varint,1,opt,name=authorization_header,json=authorizationHeader,proto3,enum=google.apps.script.type.HttpAuthorizationHeader" json:"authorization_header,omitempty"`
}

func (x *HttpOptions) Reset() {
	*x = HttpOptions{}
	if protoimpl.UnsafeEnabled {
		mi := &file_google_apps_script_type_script_manifest_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *HttpOptions) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*HttpOptions) ProtoMessage() {}

func (x *HttpOptions) ProtoReflect() protoreflect.Message {
	mi := &file_google_apps_script_type_script_manifest_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use HttpOptions.ProtoReflect.Descriptor instead.
func (*HttpOptions) Descriptor() ([]byte, []int) {
	return file_google_apps_script_type_script_manifest_proto_rawDescGZIP(), []int{2}
}

func (x *HttpOptions) GetAuthorizationHeader() HttpAuthorizationHeader {
	if x != nil {
		return x.AuthorizationHeader
	}
	return HttpAuthorizationHeader_HTTP_AUTHORIZATION_HEADER_UNSPECIFIED
}

var File_google_apps_script_type_script_manifest_proto protoreflect.FileDescriptor

var file_google_apps_script_type_script_manifest_proto_rawDesc = []byte{
	0x0a, 0x2d, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x61, 0x70, 0x70, 0x73, 0x2f, 0x73, 0x63,
	0x72, 0x69, 0x70, 0x74, 0x2f, 0x74, 0x79, 0x70, 0x65, 0x2f, 0x73, 0x63, 0x72, 0x69, 0x70, 0x74,
	0x5f, 0x6d, 0x61, 0x6e, 0x69, 0x66, 0x65, 0x73, 0x74, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12,
	0x17, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x61, 0x70, 0x70, 0x73, 0x2e, 0x73, 0x63, 0x72,
	0x69, 0x70, 0x74, 0x2e, 0x74, 0x79, 0x70, 0x65, 0x1a, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65,
	0x2f, 0x61, 0x70, 0x70, 0x73, 0x2f, 0x73, 0x63, 0x72, 0x69, 0x70, 0x74, 0x2f, 0x74, 0x79, 0x70,
	0x65, 0x2f, 0x61, 0x64, 0x64, 0x6f, 0x6e, 0x5f, 0x77, 0x69, 0x64, 0x67, 0x65, 0x74, 0x5f, 0x73,
	0x65, 0x74, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x2d, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65,
	0x2f, 0x61, 0x70, 0x70, 0x73, 0x2f, 0x73, 0x63, 0x72, 0x69, 0x70, 0x74, 0x2f, 0x74, 0x79, 0x70,
	0x65, 0x2f, 0x65, 0x78, 0x74, 0x65, 0x6e, 0x73, 0x69, 0x6f, 0x6e, 0x5f, 0x70, 0x6f, 0x69, 0x6e,
	0x74, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x1c, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2f, 0x73, 0x74, 0x72, 0x75, 0x63, 0x74, 0x2e,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0xb1, 0x04, 0x0a, 0x13, 0x43, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e,
	0x41, 0x64, 0x64, 0x4f, 0x6e, 0x4d, 0x61, 0x6e, 0x69, 0x66, 0x65, 0x73, 0x74, 0x12, 0x12, 0x0a,
	0x04, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x6e, 0x61, 0x6d,
	0x65, 0x12, 0x19, 0x0a, 0x08, 0x6c, 0x6f, 0x67, 0x6f, 0x5f, 0x75, 0x72, 0x6c, 0x18, 0x02, 0x20,
	0x01, 0x28, 0x09, 0x52, 0x07, 0x6c, 0x6f, 0x67, 0x6f, 0x55, 0x72, 0x6c, 0x12, 0x56, 0x0a, 0x11,
	0x6c, 0x61, 0x79, 0x6f, 0x75, 0x74, 0x5f, 0x70, 0x72, 0x6f, 0x70, 0x65, 0x72, 0x74, 0x69, 0x65,
	0x73, 0x18, 0x03, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x29, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65,
	0x2e, 0x61, 0x70, 0x70, 0x73, 0x2e, 0x73, 0x63, 0x72, 0x69, 0x70, 0x74, 0x2e, 0x74, 0x79, 0x70,
	0x65, 0x2e, 0x4c, 0x61, 0x79, 0x6f, 0x75, 0x74, 0x50, 0x72, 0x6f, 0x70, 0x65, 0x72, 0x74, 0x69,
	0x65, 0x73, 0x52, 0x10, 0x6c, 0x61, 0x79, 0x6f, 0x75, 0x74, 0x50, 0x72, 0x6f, 0x70, 0x65, 0x72,
	0x74, 0x69, 0x65, 0x73, 0x12, 0x52, 0x0a, 0x11, 0x61, 0x64, 0x64, 0x5f, 0x6f, 0x6e, 0x5f, 0x77,
	0x69, 0x64, 0x67, 0x65, 0x74, 0x5f, 0x73, 0x65, 0x74, 0x18, 0x04, 0x20, 0x01, 0x28, 0x0b, 0x32,
	0x27, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x61, 0x70, 0x70, 0x73, 0x2e, 0x73, 0x63,
	0x72, 0x69, 0x70, 0x74, 0x2e, 0x74, 0x79, 0x70, 0x65, 0x2e, 0x41, 0x64, 0x64, 0x4f, 0x6e, 0x57,
	0x69, 0x64, 0x67, 0x65, 0x74, 0x53, 0x65, 0x74, 0x52, 0x0e, 0x61, 0x64, 0x64, 0x4f, 0x6e, 0x57,
	0x69, 0x64, 0x67, 0x65, 0x74, 0x53, 0x65, 0x74, 0x12, 0x2d, 0x0a, 0x13, 0x75, 0x73, 0x65, 0x5f,
	0x6c, 0x6f, 0x63, 0x61, 0x6c, 0x65, 0x5f, 0x66, 0x72, 0x6f, 0x6d, 0x5f, 0x61, 0x70, 0x70, 0x18,
	0x05, 0x20, 0x01, 0x28, 0x08, 0x52, 0x10, 0x75, 0x73, 0x65, 0x4c, 0x6f, 0x63, 0x61, 0x6c, 0x65,
	0x46, 0x72, 0x6f, 0x6d, 0x41, 0x70, 0x70, 0x12, 0x5a, 0x0a, 0x10, 0x68, 0x6f, 0x6d, 0x65, 0x70,
	0x61, 0x67, 0x65, 0x5f, 0x74, 0x72, 0x69, 0x67, 0x67, 0x65, 0x72, 0x18, 0x06, 0x20, 0x01, 0x28,
	0x0b, 0x32, 0x2f, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x61, 0x70, 0x70, 0x73, 0x2e,
	0x73, 0x63, 0x72, 0x69, 0x70, 0x74, 0x2e, 0x74, 0x79, 0x70, 0x65, 0x2e, 0x48, 0x6f, 0x6d, 0x65,
	0x70, 0x61, 0x67, 0x65, 0x45, 0x78, 0x74, 0x65, 0x6e, 0x73, 0x69, 0x6f, 0x6e, 0x50, 0x6f, 0x69,
	0x6e, 0x74, 0x52, 0x0f, 0x68, 0x6f, 0x6d, 0x65, 0x70, 0x61, 0x67, 0x65, 0x54, 0x72, 0x69, 0x67,
	0x67, 0x65, 0x72, 0x12, 0x63, 0x0a, 0x11, 0x75, 0x6e, 0x69, 0x76, 0x65, 0x72, 0x73, 0x61, 0x6c,
	0x5f, 0x61, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x18, 0x07, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x36,
	0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x61, 0x70, 0x70, 0x73, 0x2e, 0x73, 0x63, 0x72,
	0x69, 0x70, 0x74, 0x2e, 0x74, 0x79, 0x70, 0x65, 0x2e, 0x55, 0x6e, 0x69, 0x76, 0x65, 0x72, 0x73,
	0x61, 0x6c, 0x41, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x45, 0x78, 0x74, 0x65, 0x6e, 0x73, 0x69, 0x6f,
	0x6e, 0x50, 0x6f, 0x69, 0x6e, 0x74, 0x52, 0x10, 0x75, 0x6e, 0x69, 0x76, 0x65, 0x72, 0x73, 0x61,
	0x6c, 0x41, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x12, 0x4f, 0x0a, 0x16, 0x6f, 0x70, 0x65, 0x6e,
	0x5f, 0x6c, 0x69, 0x6e, 0x6b, 0x5f, 0x75, 0x72, 0x6c, 0x5f, 0x70, 0x72, 0x65, 0x66, 0x69, 0x78,
	0x65, 0x73, 0x18, 0x08, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1a, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c,
	0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x4c, 0x69, 0x73, 0x74, 0x56,
	0x61, 0x6c, 0x75, 0x65, 0x52, 0x13, 0x6f, 0x70, 0x65, 0x6e, 0x4c, 0x69, 0x6e, 0x6b, 0x55, 0x72,
	0x6c, 0x50, 0x72, 0x65, 0x66, 0x69, 0x78, 0x65, 0x73, 0x22, 0x60, 0x0a, 0x10, 0x4c, 0x61, 0x79,
	0x6f, 0x75, 0x74, 0x50, 0x72, 0x6f, 0x70, 0x65, 0x72, 0x74, 0x69, 0x65, 0x73, 0x12, 0x23, 0x0a,
	0x0d, 0x70, 0x72, 0x69, 0x6d, 0x61, 0x72, 0x79, 0x5f, 0x63, 0x6f, 0x6c, 0x6f, 0x72, 0x18, 0x01,
	0x20, 0x01, 0x28, 0x09, 0x52, 0x0c, 0x70, 0x72, 0x69, 0x6d, 0x61, 0x72, 0x79, 0x43, 0x6f, 0x6c,
	0x6f, 0x72, 0x12, 0x27, 0x0a, 0x0f, 0x73, 0x65, 0x63, 0x6f, 0x6e, 0x64, 0x61, 0x72, 0x79, 0x5f,
	0x63, 0x6f, 0x6c, 0x6f, 0x72, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0e, 0x73, 0x65, 0x63,
	0x6f, 0x6e, 0x64, 0x61, 0x72, 0x79, 0x43, 0x6f, 0x6c, 0x6f, 0x72, 0x22, 0x72, 0x0a, 0x0b, 0x48,
	0x74, 0x74, 0x70, 0x4f, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x12, 0x63, 0x0a, 0x14, 0x61, 0x75,
	0x74, 0x68, 0x6f, 0x72, 0x69, 0x7a, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x5f, 0x68, 0x65, 0x61, 0x64,
	0x65, 0x72, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0e, 0x32, 0x30, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c,
	0x65, 0x2e, 0x61, 0x70, 0x70, 0x73, 0x2e, 0x73, 0x63, 0x72, 0x69, 0x70, 0x74, 0x2e, 0x74, 0x79,
	0x70, 0x65, 0x2e, 0x48, 0x74, 0x74, 0x70, 0x41, 0x75, 0x74, 0x68, 0x6f, 0x72, 0x69, 0x7a, 0x61,
	0x74, 0x69, 0x6f, 0x6e, 0x48, 0x65, 0x61, 0x64, 0x65, 0x72, 0x52, 0x13, 0x61, 0x75, 0x74, 0x68,
	0x6f, 0x72, 0x69, 0x7a, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x48, 0x65, 0x61, 0x64, 0x65, 0x72, 0x2a,
	0x76, 0x0a, 0x17, 0x48, 0x74, 0x74, 0x70, 0x41, 0x75, 0x74, 0x68, 0x6f, 0x72, 0x69, 0x7a, 0x61,
	0x74, 0x69, 0x6f, 0x6e, 0x48, 0x65, 0x61, 0x64, 0x65, 0x72, 0x12, 0x29, 0x0a, 0x25, 0x48, 0x54,
	0x54, 0x50, 0x5f, 0x41, 0x55, 0x54, 0x48, 0x4f, 0x52, 0x49, 0x5a, 0x41, 0x54, 0x49, 0x4f, 0x4e,
	0x5f, 0x48, 0x45, 0x41, 0x44, 0x45, 0x52, 0x5f, 0x55, 0x4e, 0x53, 0x50, 0x45, 0x43, 0x49, 0x46,
	0x49, 0x45, 0x44, 0x10, 0x00, 0x12, 0x13, 0x0a, 0x0f, 0x53, 0x59, 0x53, 0x54, 0x45, 0x4d, 0x5f,
	0x49, 0x44, 0x5f, 0x54, 0x4f, 0x4b, 0x45, 0x4e, 0x10, 0x01, 0x12, 0x11, 0x0a, 0x0d, 0x55, 0x53,
	0x45, 0x52, 0x5f, 0x49, 0x44, 0x5f, 0x54, 0x4f, 0x4b, 0x45, 0x4e, 0x10, 0x02, 0x12, 0x08, 0x0a,
	0x04, 0x4e, 0x4f, 0x4e, 0x45, 0x10, 0x03, 0x42, 0xa8, 0x01, 0x0a, 0x1b, 0x63, 0x6f, 0x6d, 0x2e,
	0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x61, 0x70, 0x70, 0x73, 0x2e, 0x73, 0x63, 0x72, 0x69,
	0x70, 0x74, 0x2e, 0x74, 0x79, 0x70, 0x65, 0x50, 0x01, 0x5a, 0x36, 0x67, 0x6f, 0x6f, 0x67, 0x6c,
	0x65, 0x2e, 0x67, 0x6f, 0x6c, 0x61, 0x6e, 0x67, 0x2e, 0x6f, 0x72, 0x67, 0x2f, 0x67, 0x65, 0x6e,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2f, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x61, 0x70, 0x69, 0x73,
	0x2f, 0x61, 0x70, 0x70, 0x73, 0x2f, 0x73, 0x63, 0x72, 0x69, 0x70, 0x74, 0x2f, 0x74, 0x79, 0x70,
	0x65, 0xaa, 0x02, 0x17, 0x47, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x41, 0x70, 0x70, 0x73, 0x2e,
	0x53, 0x63, 0x72, 0x69, 0x70, 0x74, 0x2e, 0x54, 0x79, 0x70, 0x65, 0xca, 0x02, 0x17, 0x47, 0x6f,
	0x6f, 0x67, 0x6c, 0x65, 0x5c, 0x41, 0x70, 0x70, 0x73, 0x5c, 0x53, 0x63, 0x72, 0x69, 0x70, 0x74,
	0x5c, 0x54, 0x79, 0x70, 0x65, 0xea, 0x02, 0x1a, 0x47, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x3a, 0x3a,
	0x41, 0x70, 0x70, 0x73, 0x3a, 0x3a, 0x53, 0x63, 0x72, 0x69, 0x70, 0x74, 0x3a, 0x3a, 0x54, 0x79,
	0x70, 0x65, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_google_apps_script_type_script_manifest_proto_rawDescOnce sync.Once
	file_google_apps_script_type_script_manifest_proto_rawDescData = file_google_apps_script_type_script_manifest_proto_rawDesc
)

func file_google_apps_script_type_script_manifest_proto_rawDescGZIP() []byte {
	file_google_apps_script_type_script_manifest_proto_rawDescOnce.Do(func() {
		file_google_apps_script_type_script_manifest_proto_rawDescData = protoimpl.X.CompressGZIP(file_google_apps_script_type_script_manifest_proto_rawDescData)
	})
	return file_google_apps_script_type_script_manifest_proto_rawDescData
}

var file_google_apps_script_type_script_manifest_proto_enumTypes = make([]protoimpl.EnumInfo, 1)
var file_google_apps_script_type_script_manifest_proto_msgTypes = make([]protoimpl.MessageInfo, 3)
var file_google_apps_script_type_script_manifest_proto_goTypes = []interface{}{
	(HttpAuthorizationHeader)(0),          // 0: google.apps.script.type.HttpAuthorizationHeader
	(*CommonAddOnManifest)(nil),           // 1: google.apps.script.type.CommonAddOnManifest
	(*LayoutProperties)(nil),              // 2: google.apps.script.type.LayoutProperties
	(*HttpOptions)(nil),                   // 3: google.apps.script.type.HttpOptions
	(*AddOnWidgetSet)(nil),                // 4: google.apps.script.type.AddOnWidgetSet
	(*HomepageExtensionPoint)(nil),        // 5: google.apps.script.type.HomepageExtensionPoint
	(*UniversalActionExtensionPoint)(nil), // 6: google.apps.script.type.UniversalActionExtensionPoint
	(*structpb.ListValue)(nil),            // 7: google.protobuf.ListValue
}
var file_google_apps_script_type_script_manifest_proto_depIdxs = []int32{
	2, // 0: google.apps.script.type.CommonAddOnManifest.layout_properties:type_name -> google.apps.script.type.LayoutProperties
	4, // 1: google.apps.script.type.CommonAddOnManifest.add_on_widget_set:type_name -> google.apps.script.type.AddOnWidgetSet
	5, // 2: google.apps.script.type.CommonAddOnManifest.homepage_trigger:type_name -> google.apps.script.type.HomepageExtensionPoint
	6, // 3: google.apps.script.type.CommonAddOnManifest.universal_actions:type_name -> google.apps.script.type.UniversalActionExtensionPoint
	7, // 4: google.apps.script.type.CommonAddOnManifest.open_link_url_prefixes:type_name -> google.protobuf.ListValue
	0, // 5: google.apps.script.type.HttpOptions.authorization_header:type_name -> google.apps.script.type.HttpAuthorizationHeader
	6, // [6:6] is the sub-list for method output_type
	6, // [6:6] is the sub-list for method input_type
	6, // [6:6] is the sub-list for extension type_name
	6, // [6:6] is the sub-list for extension extendee
	0, // [0:6] is the sub-list for field type_name
}

func init() { file_google_apps_script_type_script_manifest_proto_init() }
func file_google_apps_script_type_script_manifest_proto_init() {
	if File_google_apps_script_type_script_manifest_proto != nil {
		return
	}
	file_google_apps_script_type_addon_widget_set_proto_init()
	file_google_apps_script_type_extension_point_proto_init()
	if !protoimpl.UnsafeEnabled {
		file_google_apps_script_type_script_manifest_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*CommonAddOnManifest); i {
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
		file_google_apps_script_type_script_manifest_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*LayoutProperties); i {
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
		file_google_apps_script_type_script_manifest_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*HttpOptions); i {
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
			RawDescriptor: file_google_apps_script_type_script_manifest_proto_rawDesc,
			NumEnums:      1,
			NumMessages:   3,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_google_apps_script_type_script_manifest_proto_goTypes,
		DependencyIndexes: file_google_apps_script_type_script_manifest_proto_depIdxs,
		EnumInfos:         file_google_apps_script_type_script_manifest_proto_enumTypes,
		MessageInfos:      file_google_apps_script_type_script_manifest_proto_msgTypes,
	}.Build()
	File_google_apps_script_type_script_manifest_proto = out.File
	file_google_apps_script_type_script_manifest_proto_rawDesc = nil
	file_google_apps_script_type_script_manifest_proto_goTypes = nil
	file_google_apps_script_type_script_manifest_proto_depIdxs = nil
}
