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
// source: google/actions/sdk/v2/interactionmodel/prompt/static_prompt.proto

package prompt

import (
	reflect "reflect"
	sync "sync"

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

// Represents a list of prompt candidates, one of which will be selected as the
// prompt to be shown in the response to the user.
// **This message is localizable.**
type StaticPrompt struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// The list of candidate prompts to be sent to the client. Each prompt has a
	// selector to determine when it can be used. The first selector that matches
	// a request will be sent and the rest will be ignored.
	Candidates []*StaticPrompt_StaticPromptCandidate `protobuf:"bytes,1,rep,name=candidates,proto3" json:"candidates,omitempty"`
}

func (x *StaticPrompt) Reset() {
	*x = StaticPrompt{}
	if protoimpl.UnsafeEnabled {
		mi := &file_google_actions_sdk_v2_interactionmodel_prompt_static_prompt_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *StaticPrompt) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*StaticPrompt) ProtoMessage() {}

func (x *StaticPrompt) ProtoReflect() protoreflect.Message {
	mi := &file_google_actions_sdk_v2_interactionmodel_prompt_static_prompt_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use StaticPrompt.ProtoReflect.Descriptor instead.
func (*StaticPrompt) Descriptor() ([]byte, []int) {
	return file_google_actions_sdk_v2_interactionmodel_prompt_static_prompt_proto_rawDescGZIP(), []int{0}
}

func (x *StaticPrompt) GetCandidates() []*StaticPrompt_StaticPromptCandidate {
	if x != nil {
		return x.Candidates
	}
	return nil
}

// Represents a static prompt candidate.
type StaticPrompt_StaticPromptCandidate struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// Optional. The criteria for whether this prompt matches a request. If the selector
	// is empty, this prompt will always be triggered.
	Selector *StaticPrompt_Selector `protobuf:"bytes,1,opt,name=selector,proto3" json:"selector,omitempty"`
	// The prompt response associated with the selector.
	PromptResponse *StaticPrompt_StaticPromptCandidate_StaticPromptResponse `protobuf:"bytes,2,opt,name=prompt_response,json=promptResponse,proto3" json:"prompt_response,omitempty"`
}

func (x *StaticPrompt_StaticPromptCandidate) Reset() {
	*x = StaticPrompt_StaticPromptCandidate{}
	if protoimpl.UnsafeEnabled {
		mi := &file_google_actions_sdk_v2_interactionmodel_prompt_static_prompt_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *StaticPrompt_StaticPromptCandidate) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*StaticPrompt_StaticPromptCandidate) ProtoMessage() {}

func (x *StaticPrompt_StaticPromptCandidate) ProtoReflect() protoreflect.Message {
	mi := &file_google_actions_sdk_v2_interactionmodel_prompt_static_prompt_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use StaticPrompt_StaticPromptCandidate.ProtoReflect.Descriptor instead.
func (*StaticPrompt_StaticPromptCandidate) Descriptor() ([]byte, []int) {
	return file_google_actions_sdk_v2_interactionmodel_prompt_static_prompt_proto_rawDescGZIP(), []int{0, 0}
}

func (x *StaticPrompt_StaticPromptCandidate) GetSelector() *StaticPrompt_Selector {
	if x != nil {
		return x.Selector
	}
	return nil
}

func (x *StaticPrompt_StaticPromptCandidate) GetPromptResponse() *StaticPrompt_StaticPromptCandidate_StaticPromptResponse {
	if x != nil {
		return x.PromptResponse
	}
	return nil
}

// Defines the criteria for whether a prompt matches a request.
type StaticPrompt_Selector struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// The set of required surface capabilities.
	SurfaceCapabilities *SurfaceCapabilities `protobuf:"bytes,1,opt,name=surface_capabilities,json=surfaceCapabilities,proto3" json:"surface_capabilities,omitempty"`
}

func (x *StaticPrompt_Selector) Reset() {
	*x = StaticPrompt_Selector{}
	if protoimpl.UnsafeEnabled {
		mi := &file_google_actions_sdk_v2_interactionmodel_prompt_static_prompt_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *StaticPrompt_Selector) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*StaticPrompt_Selector) ProtoMessage() {}

func (x *StaticPrompt_Selector) ProtoReflect() protoreflect.Message {
	mi := &file_google_actions_sdk_v2_interactionmodel_prompt_static_prompt_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use StaticPrompt_Selector.ProtoReflect.Descriptor instead.
func (*StaticPrompt_Selector) Descriptor() ([]byte, []int) {
	return file_google_actions_sdk_v2_interactionmodel_prompt_static_prompt_proto_rawDescGZIP(), []int{0, 1}
}

func (x *StaticPrompt_Selector) GetSurfaceCapabilities() *SurfaceCapabilities {
	if x != nil {
		return x.SurfaceCapabilities
	}
	return nil
}

// Represents structured responses to send to the user, such as text,
// speech, cards, canvas data, suggestion chips, etc.
type StaticPrompt_StaticPromptCandidate_StaticPromptResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// Optional. The first voice and text-only response.
	FirstSimple *StaticSimplePrompt `protobuf:"bytes,2,opt,name=first_simple,json=firstSimple,proto3" json:"first_simple,omitempty"`
	// Optional. A content like a card, list or media to display to the user.
	Content *StaticContentPrompt `protobuf:"bytes,3,opt,name=content,proto3" json:"content,omitempty"`
	// Optional. The last voice and text-only response.
	LastSimple *StaticSimplePrompt `protobuf:"bytes,4,opt,name=last_simple,json=lastSimple,proto3" json:"last_simple,omitempty"`
	// Optional. Suggestions to be displayed to the user which will always
	// appear at the end of the response. If the `append` field in the
	// containing prompt is `true` the titles defined in this field will be
	// added to titles defined in any previously defined suggestions prompts
	// and duplicate values will be removed.
	Suggestions []*Suggestion `protobuf:"bytes,5,rep,name=suggestions,proto3" json:"suggestions,omitempty"`
	// Optional. An additional suggestion chip that can link out to the associated app
	// or site.
	// The chip will be rendered with the title "Open <name>". Max 20 chars.
	Link *StaticLinkPrompt `protobuf:"bytes,6,opt,name=link,proto3" json:"link,omitempty"`
	// Optional. Mode for how this messages should be merged with previously defined
	// messages.
	// `true` will clear all previously defined messages (first and last
	// simple, content, suggestions link and canvas) and add messages defined
	// in this prompt. `false` will add messages defined in this prompt to
	// messages defined in previous responses. Setting this field to `false`
	// will also enable appending to some fields inside Simple prompts, the
	// Suggestions prompt and the Canvas prompt (part of the Content prompt).
	// The Content and Link messages will always be overwritten if defined in
	// the prompt. Default value is `false`.
	Override bool `protobuf:"varint,7,opt,name=override,proto3" json:"override,omitempty"`
	// A response to be used for interactive canvas experience.
	Canvas *StaticCanvasPrompt `protobuf:"bytes,8,opt,name=canvas,proto3" json:"canvas,omitempty"`
}

func (x *StaticPrompt_StaticPromptCandidate_StaticPromptResponse) Reset() {
	*x = StaticPrompt_StaticPromptCandidate_StaticPromptResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_google_actions_sdk_v2_interactionmodel_prompt_static_prompt_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *StaticPrompt_StaticPromptCandidate_StaticPromptResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*StaticPrompt_StaticPromptCandidate_StaticPromptResponse) ProtoMessage() {}

func (x *StaticPrompt_StaticPromptCandidate_StaticPromptResponse) ProtoReflect() protoreflect.Message {
	mi := &file_google_actions_sdk_v2_interactionmodel_prompt_static_prompt_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use StaticPrompt_StaticPromptCandidate_StaticPromptResponse.ProtoReflect.Descriptor instead.
func (*StaticPrompt_StaticPromptCandidate_StaticPromptResponse) Descriptor() ([]byte, []int) {
	return file_google_actions_sdk_v2_interactionmodel_prompt_static_prompt_proto_rawDescGZIP(), []int{0, 0, 0}
}

func (x *StaticPrompt_StaticPromptCandidate_StaticPromptResponse) GetFirstSimple() *StaticSimplePrompt {
	if x != nil {
		return x.FirstSimple
	}
	return nil
}

func (x *StaticPrompt_StaticPromptCandidate_StaticPromptResponse) GetContent() *StaticContentPrompt {
	if x != nil {
		return x.Content
	}
	return nil
}

func (x *StaticPrompt_StaticPromptCandidate_StaticPromptResponse) GetLastSimple() *StaticSimplePrompt {
	if x != nil {
		return x.LastSimple
	}
	return nil
}

func (x *StaticPrompt_StaticPromptCandidate_StaticPromptResponse) GetSuggestions() []*Suggestion {
	if x != nil {
		return x.Suggestions
	}
	return nil
}

func (x *StaticPrompt_StaticPromptCandidate_StaticPromptResponse) GetLink() *StaticLinkPrompt {
	if x != nil {
		return x.Link
	}
	return nil
}

func (x *StaticPrompt_StaticPromptCandidate_StaticPromptResponse) GetOverride() bool {
	if x != nil {
		return x.Override
	}
	return false
}

func (x *StaticPrompt_StaticPromptCandidate_StaticPromptResponse) GetCanvas() *StaticCanvasPrompt {
	if x != nil {
		return x.Canvas
	}
	return nil
}

var File_google_actions_sdk_v2_interactionmodel_prompt_static_prompt_proto protoreflect.FileDescriptor

var file_google_actions_sdk_v2_interactionmodel_prompt_static_prompt_proto_rawDesc = []byte{
	0x0a, 0x41, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x61, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x73,
	0x2f, 0x73, 0x64, 0x6b, 0x2f, 0x76, 0x32, 0x2f, 0x69, 0x6e, 0x74, 0x65, 0x72, 0x61, 0x63, 0x74,
	0x69, 0x6f, 0x6e, 0x6d, 0x6f, 0x64, 0x65, 0x6c, 0x2f, 0x70, 0x72, 0x6f, 0x6d, 0x70, 0x74, 0x2f,
	0x73, 0x74, 0x61, 0x74, 0x69, 0x63, 0x5f, 0x70, 0x72, 0x6f, 0x6d, 0x70, 0x74, 0x2e, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x12, 0x2d, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x61, 0x63, 0x74, 0x69,
	0x6f, 0x6e, 0x73, 0x2e, 0x73, 0x64, 0x6b, 0x2e, 0x76, 0x32, 0x2e, 0x69, 0x6e, 0x74, 0x65, 0x72,
	0x61, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x6d, 0x6f, 0x64, 0x65, 0x6c, 0x2e, 0x70, 0x72, 0x6f, 0x6d,
	0x70, 0x74, 0x1a, 0x50, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x61, 0x63, 0x74, 0x69, 0x6f,
	0x6e, 0x73, 0x2f, 0x73, 0x64, 0x6b, 0x2f, 0x76, 0x32, 0x2f, 0x69, 0x6e, 0x74, 0x65, 0x72, 0x61,
	0x63, 0x74, 0x69, 0x6f, 0x6e, 0x6d, 0x6f, 0x64, 0x65, 0x6c, 0x2f, 0x70, 0x72, 0x6f, 0x6d, 0x70,
	0x74, 0x2f, 0x63, 0x6f, 0x6e, 0x74, 0x65, 0x6e, 0x74, 0x2f, 0x73, 0x74, 0x61, 0x74, 0x69, 0x63,
	0x5f, 0x63, 0x61, 0x6e, 0x76, 0x61, 0x73, 0x5f, 0x70, 0x72, 0x6f, 0x6d, 0x70, 0x74, 0x2e, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x51, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x61, 0x63, 0x74,
	0x69, 0x6f, 0x6e, 0x73, 0x2f, 0x73, 0x64, 0x6b, 0x2f, 0x76, 0x32, 0x2f, 0x69, 0x6e, 0x74, 0x65,
	0x72, 0x61, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x6d, 0x6f, 0x64, 0x65, 0x6c, 0x2f, 0x70, 0x72, 0x6f,
	0x6d, 0x70, 0x74, 0x2f, 0x63, 0x6f, 0x6e, 0x74, 0x65, 0x6e, 0x74, 0x2f, 0x73, 0x74, 0x61, 0x74,
	0x69, 0x63, 0x5f, 0x63, 0x6f, 0x6e, 0x74, 0x65, 0x6e, 0x74, 0x5f, 0x70, 0x72, 0x6f, 0x6d, 0x70,
	0x74, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x4e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f,
	0x61, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x2f, 0x73, 0x64, 0x6b, 0x2f, 0x76, 0x32, 0x2f, 0x69,
	0x6e, 0x74, 0x65, 0x72, 0x61, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x6d, 0x6f, 0x64, 0x65, 0x6c, 0x2f,
	0x70, 0x72, 0x6f, 0x6d, 0x70, 0x74, 0x2f, 0x63, 0x6f, 0x6e, 0x74, 0x65, 0x6e, 0x74, 0x2f, 0x73,
	0x74, 0x61, 0x74, 0x69, 0x63, 0x5f, 0x6c, 0x69, 0x6e, 0x6b, 0x5f, 0x70, 0x72, 0x6f, 0x6d, 0x70,
	0x74, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x48, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f,
	0x61, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x2f, 0x73, 0x64, 0x6b, 0x2f, 0x76, 0x32, 0x2f, 0x69,
	0x6e, 0x74, 0x65, 0x72, 0x61, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x6d, 0x6f, 0x64, 0x65, 0x6c, 0x2f,
	0x70, 0x72, 0x6f, 0x6d, 0x70, 0x74, 0x2f, 0x73, 0x74, 0x61, 0x74, 0x69, 0x63, 0x5f, 0x73, 0x69,
	0x6d, 0x70, 0x6c, 0x65, 0x5f, 0x70, 0x72, 0x6f, 0x6d, 0x70, 0x74, 0x2e, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x1a, 0x3e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x61, 0x63, 0x74, 0x69, 0x6f, 0x6e,
	0x73, 0x2f, 0x73, 0x64, 0x6b, 0x2f, 0x76, 0x32, 0x2f, 0x69, 0x6e, 0x74, 0x65, 0x72, 0x61, 0x63,
	0x74, 0x69, 0x6f, 0x6e, 0x6d, 0x6f, 0x64, 0x65, 0x6c, 0x2f, 0x70, 0x72, 0x6f, 0x6d, 0x70, 0x74,
	0x2f, 0x73, 0x75, 0x67, 0x67, 0x65, 0x73, 0x74, 0x69, 0x6f, 0x6e, 0x2e, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x1a, 0x48, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x61, 0x63, 0x74, 0x69, 0x6f, 0x6e,
	0x73, 0x2f, 0x73, 0x64, 0x6b, 0x2f, 0x76, 0x32, 0x2f, 0x69, 0x6e, 0x74, 0x65, 0x72, 0x61, 0x63,
	0x74, 0x69, 0x6f, 0x6e, 0x6d, 0x6f, 0x64, 0x65, 0x6c, 0x2f, 0x70, 0x72, 0x6f, 0x6d, 0x70, 0x74,
	0x2f, 0x73, 0x75, 0x72, 0x66, 0x61, 0x63, 0x65, 0x5f, 0x63, 0x61, 0x70, 0x61, 0x62, 0x69, 0x6c,
	0x69, 0x74, 0x69, 0x65, 0x73, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x1f, 0x67, 0x6f, 0x6f,
	0x67, 0x6c, 0x65, 0x2f, 0x61, 0x70, 0x69, 0x2f, 0x66, 0x69, 0x65, 0x6c, 0x64, 0x5f, 0x62, 0x65,
	0x68, 0x61, 0x76, 0x69, 0x6f, 0x72, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0xa0, 0x09, 0x0a,
	0x0c, 0x53, 0x74, 0x61, 0x74, 0x69, 0x63, 0x50, 0x72, 0x6f, 0x6d, 0x70, 0x74, 0x12, 0x71, 0x0a,
	0x0a, 0x63, 0x61, 0x6e, 0x64, 0x69, 0x64, 0x61, 0x74, 0x65, 0x73, 0x18, 0x01, 0x20, 0x03, 0x28,
	0x0b, 0x32, 0x51, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x61, 0x63, 0x74, 0x69, 0x6f,
	0x6e, 0x73, 0x2e, 0x73, 0x64, 0x6b, 0x2e, 0x76, 0x32, 0x2e, 0x69, 0x6e, 0x74, 0x65, 0x72, 0x61,
	0x63, 0x74, 0x69, 0x6f, 0x6e, 0x6d, 0x6f, 0x64, 0x65, 0x6c, 0x2e, 0x70, 0x72, 0x6f, 0x6d, 0x70,
	0x74, 0x2e, 0x53, 0x74, 0x61, 0x74, 0x69, 0x63, 0x50, 0x72, 0x6f, 0x6d, 0x70, 0x74, 0x2e, 0x53,
	0x74, 0x61, 0x74, 0x69, 0x63, 0x50, 0x72, 0x6f, 0x6d, 0x70, 0x74, 0x43, 0x61, 0x6e, 0x64, 0x69,
	0x64, 0x61, 0x74, 0x65, 0x52, 0x0a, 0x63, 0x61, 0x6e, 0x64, 0x69, 0x64, 0x61, 0x74, 0x65, 0x73,
	0x1a, 0x98, 0x07, 0x0a, 0x15, 0x53, 0x74, 0x61, 0x74, 0x69, 0x63, 0x50, 0x72, 0x6f, 0x6d, 0x70,
	0x74, 0x43, 0x61, 0x6e, 0x64, 0x69, 0x64, 0x61, 0x74, 0x65, 0x12, 0x65, 0x0a, 0x08, 0x73, 0x65,
	0x6c, 0x65, 0x63, 0x74, 0x6f, 0x72, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x44, 0x2e, 0x67,
	0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x61, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x2e, 0x73, 0x64,
	0x6b, 0x2e, 0x76, 0x32, 0x2e, 0x69, 0x6e, 0x74, 0x65, 0x72, 0x61, 0x63, 0x74, 0x69, 0x6f, 0x6e,
	0x6d, 0x6f, 0x64, 0x65, 0x6c, 0x2e, 0x70, 0x72, 0x6f, 0x6d, 0x70, 0x74, 0x2e, 0x53, 0x74, 0x61,
	0x74, 0x69, 0x63, 0x50, 0x72, 0x6f, 0x6d, 0x70, 0x74, 0x2e, 0x53, 0x65, 0x6c, 0x65, 0x63, 0x74,
	0x6f, 0x72, 0x42, 0x03, 0xe0, 0x41, 0x01, 0x52, 0x08, 0x73, 0x65, 0x6c, 0x65, 0x63, 0x74, 0x6f,
	0x72, 0x12, 0x8f, 0x01, 0x0a, 0x0f, 0x70, 0x72, 0x6f, 0x6d, 0x70, 0x74, 0x5f, 0x72, 0x65, 0x73,
	0x70, 0x6f, 0x6e, 0x73, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x66, 0x2e, 0x67, 0x6f,
	0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x61, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x2e, 0x73, 0x64, 0x6b,
	0x2e, 0x76, 0x32, 0x2e, 0x69, 0x6e, 0x74, 0x65, 0x72, 0x61, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x6d,
	0x6f, 0x64, 0x65, 0x6c, 0x2e, 0x70, 0x72, 0x6f, 0x6d, 0x70, 0x74, 0x2e, 0x53, 0x74, 0x61, 0x74,
	0x69, 0x63, 0x50, 0x72, 0x6f, 0x6d, 0x70, 0x74, 0x2e, 0x53, 0x74, 0x61, 0x74, 0x69, 0x63, 0x50,
	0x72, 0x6f, 0x6d, 0x70, 0x74, 0x43, 0x61, 0x6e, 0x64, 0x69, 0x64, 0x61, 0x74, 0x65, 0x2e, 0x53,
	0x74, 0x61, 0x74, 0x69, 0x63, 0x50, 0x72, 0x6f, 0x6d, 0x70, 0x74, 0x52, 0x65, 0x73, 0x70, 0x6f,
	0x6e, 0x73, 0x65, 0x52, 0x0e, 0x70, 0x72, 0x6f, 0x6d, 0x70, 0x74, 0x52, 0x65, 0x73, 0x70, 0x6f,
	0x6e, 0x73, 0x65, 0x1a, 0x85, 0x05, 0x0a, 0x14, 0x53, 0x74, 0x61, 0x74, 0x69, 0x63, 0x50, 0x72,
	0x6f, 0x6d, 0x70, 0x74, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x69, 0x0a, 0x0c,
	0x66, 0x69, 0x72, 0x73, 0x74, 0x5f, 0x73, 0x69, 0x6d, 0x70, 0x6c, 0x65, 0x18, 0x02, 0x20, 0x01,
	0x28, 0x0b, 0x32, 0x41, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x61, 0x63, 0x74, 0x69,
	0x6f, 0x6e, 0x73, 0x2e, 0x73, 0x64, 0x6b, 0x2e, 0x76, 0x32, 0x2e, 0x69, 0x6e, 0x74, 0x65, 0x72,
	0x61, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x6d, 0x6f, 0x64, 0x65, 0x6c, 0x2e, 0x70, 0x72, 0x6f, 0x6d,
	0x70, 0x74, 0x2e, 0x53, 0x74, 0x61, 0x74, 0x69, 0x63, 0x53, 0x69, 0x6d, 0x70, 0x6c, 0x65, 0x50,
	0x72, 0x6f, 0x6d, 0x70, 0x74, 0x42, 0x03, 0xe0, 0x41, 0x01, 0x52, 0x0b, 0x66, 0x69, 0x72, 0x73,
	0x74, 0x53, 0x69, 0x6d, 0x70, 0x6c, 0x65, 0x12, 0x61, 0x0a, 0x07, 0x63, 0x6f, 0x6e, 0x74, 0x65,
	0x6e, 0x74, 0x18, 0x03, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x42, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c,
	0x65, 0x2e, 0x61, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x2e, 0x73, 0x64, 0x6b, 0x2e, 0x76, 0x32,
	0x2e, 0x69, 0x6e, 0x74, 0x65, 0x72, 0x61, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x6d, 0x6f, 0x64, 0x65,
	0x6c, 0x2e, 0x70, 0x72, 0x6f, 0x6d, 0x70, 0x74, 0x2e, 0x53, 0x74, 0x61, 0x74, 0x69, 0x63, 0x43,
	0x6f, 0x6e, 0x74, 0x65, 0x6e, 0x74, 0x50, 0x72, 0x6f, 0x6d, 0x70, 0x74, 0x42, 0x03, 0xe0, 0x41,
	0x01, 0x52, 0x07, 0x63, 0x6f, 0x6e, 0x74, 0x65, 0x6e, 0x74, 0x12, 0x67, 0x0a, 0x0b, 0x6c, 0x61,
	0x73, 0x74, 0x5f, 0x73, 0x69, 0x6d, 0x70, 0x6c, 0x65, 0x18, 0x04, 0x20, 0x01, 0x28, 0x0b, 0x32,
	0x41, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x61, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x73,
	0x2e, 0x73, 0x64, 0x6b, 0x2e, 0x76, 0x32, 0x2e, 0x69, 0x6e, 0x74, 0x65, 0x72, 0x61, 0x63, 0x74,
	0x69, 0x6f, 0x6e, 0x6d, 0x6f, 0x64, 0x65, 0x6c, 0x2e, 0x70, 0x72, 0x6f, 0x6d, 0x70, 0x74, 0x2e,
	0x53, 0x74, 0x61, 0x74, 0x69, 0x63, 0x53, 0x69, 0x6d, 0x70, 0x6c, 0x65, 0x50, 0x72, 0x6f, 0x6d,
	0x70, 0x74, 0x42, 0x03, 0xe0, 0x41, 0x01, 0x52, 0x0a, 0x6c, 0x61, 0x73, 0x74, 0x53, 0x69, 0x6d,
	0x70, 0x6c, 0x65, 0x12, 0x60, 0x0a, 0x0b, 0x73, 0x75, 0x67, 0x67, 0x65, 0x73, 0x74, 0x69, 0x6f,
	0x6e, 0x73, 0x18, 0x05, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x39, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c,
	0x65, 0x2e, 0x61, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x2e, 0x73, 0x64, 0x6b, 0x2e, 0x76, 0x32,
	0x2e, 0x69, 0x6e, 0x74, 0x65, 0x72, 0x61, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x6d, 0x6f, 0x64, 0x65,
	0x6c, 0x2e, 0x70, 0x72, 0x6f, 0x6d, 0x70, 0x74, 0x2e, 0x53, 0x75, 0x67, 0x67, 0x65, 0x73, 0x74,
	0x69, 0x6f, 0x6e, 0x42, 0x03, 0xe0, 0x41, 0x01, 0x52, 0x0b, 0x73, 0x75, 0x67, 0x67, 0x65, 0x73,
	0x74, 0x69, 0x6f, 0x6e, 0x73, 0x12, 0x58, 0x0a, 0x04, 0x6c, 0x69, 0x6e, 0x6b, 0x18, 0x06, 0x20,
	0x01, 0x28, 0x0b, 0x32, 0x3f, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x61, 0x63, 0x74,
	0x69, 0x6f, 0x6e, 0x73, 0x2e, 0x73, 0x64, 0x6b, 0x2e, 0x76, 0x32, 0x2e, 0x69, 0x6e, 0x74, 0x65,
	0x72, 0x61, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x6d, 0x6f, 0x64, 0x65, 0x6c, 0x2e, 0x70, 0x72, 0x6f,
	0x6d, 0x70, 0x74, 0x2e, 0x53, 0x74, 0x61, 0x74, 0x69, 0x63, 0x4c, 0x69, 0x6e, 0x6b, 0x50, 0x72,
	0x6f, 0x6d, 0x70, 0x74, 0x42, 0x03, 0xe0, 0x41, 0x01, 0x52, 0x04, 0x6c, 0x69, 0x6e, 0x6b, 0x12,
	0x1f, 0x0a, 0x08, 0x6f, 0x76, 0x65, 0x72, 0x72, 0x69, 0x64, 0x65, 0x18, 0x07, 0x20, 0x01, 0x28,
	0x08, 0x42, 0x03, 0xe0, 0x41, 0x01, 0x52, 0x08, 0x6f, 0x76, 0x65, 0x72, 0x72, 0x69, 0x64, 0x65,
	0x12, 0x59, 0x0a, 0x06, 0x63, 0x61, 0x6e, 0x76, 0x61, 0x73, 0x18, 0x08, 0x20, 0x01, 0x28, 0x0b,
	0x32, 0x41, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x61, 0x63, 0x74, 0x69, 0x6f, 0x6e,
	0x73, 0x2e, 0x73, 0x64, 0x6b, 0x2e, 0x76, 0x32, 0x2e, 0x69, 0x6e, 0x74, 0x65, 0x72, 0x61, 0x63,
	0x74, 0x69, 0x6f, 0x6e, 0x6d, 0x6f, 0x64, 0x65, 0x6c, 0x2e, 0x70, 0x72, 0x6f, 0x6d, 0x70, 0x74,
	0x2e, 0x53, 0x74, 0x61, 0x74, 0x69, 0x63, 0x43, 0x61, 0x6e, 0x76, 0x61, 0x73, 0x50, 0x72, 0x6f,
	0x6d, 0x70, 0x74, 0x52, 0x06, 0x63, 0x61, 0x6e, 0x76, 0x61, 0x73, 0x1a, 0x81, 0x01, 0x0a, 0x08,
	0x53, 0x65, 0x6c, 0x65, 0x63, 0x74, 0x6f, 0x72, 0x12, 0x75, 0x0a, 0x14, 0x73, 0x75, 0x72, 0x66,
	0x61, 0x63, 0x65, 0x5f, 0x63, 0x61, 0x70, 0x61, 0x62, 0x69, 0x6c, 0x69, 0x74, 0x69, 0x65, 0x73,
	0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x42, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e,
	0x61, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x2e, 0x73, 0x64, 0x6b, 0x2e, 0x76, 0x32, 0x2e, 0x69,
	0x6e, 0x74, 0x65, 0x72, 0x61, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x6d, 0x6f, 0x64, 0x65, 0x6c, 0x2e,
	0x70, 0x72, 0x6f, 0x6d, 0x70, 0x74, 0x2e, 0x53, 0x75, 0x72, 0x66, 0x61, 0x63, 0x65, 0x43, 0x61,
	0x70, 0x61, 0x62, 0x69, 0x6c, 0x69, 0x74, 0x69, 0x65, 0x73, 0x52, 0x13, 0x73, 0x75, 0x72, 0x66,
	0x61, 0x63, 0x65, 0x43, 0x61, 0x70, 0x61, 0x62, 0x69, 0x6c, 0x69, 0x74, 0x69, 0x65, 0x73, 0x42,
	0x9d, 0x01, 0x0a, 0x31, 0x63, 0x6f, 0x6d, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x61,
	0x63, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x2e, 0x73, 0x64, 0x6b, 0x2e, 0x76, 0x32, 0x2e, 0x69, 0x6e,
	0x74, 0x65, 0x72, 0x61, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x6d, 0x6f, 0x64, 0x65, 0x6c, 0x2e, 0x70,
	0x72, 0x6f, 0x6d, 0x70, 0x74, 0x42, 0x11, 0x53, 0x74, 0x61, 0x74, 0x69, 0x63, 0x50, 0x72, 0x6f,
	0x6d, 0x70, 0x74, 0x50, 0x72, 0x6f, 0x74, 0x6f, 0x50, 0x01, 0x5a, 0x53, 0x67, 0x6f, 0x6f, 0x67,
	0x6c, 0x65, 0x2e, 0x67, 0x6f, 0x6c, 0x61, 0x6e, 0x67, 0x2e, 0x6f, 0x72, 0x67, 0x2f, 0x67, 0x65,
	0x6e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2f, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x61, 0x70, 0x69,
	0x73, 0x2f, 0x61, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x2f, 0x73, 0x64, 0x6b, 0x2f, 0x76, 0x32,
	0x2f, 0x69, 0x6e, 0x74, 0x65, 0x72, 0x61, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x6d, 0x6f, 0x64, 0x65,
	0x6c, 0x2f, 0x70, 0x72, 0x6f, 0x6d, 0x70, 0x74, 0x3b, 0x70, 0x72, 0x6f, 0x6d, 0x70, 0x74, 0x62,
	0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_google_actions_sdk_v2_interactionmodel_prompt_static_prompt_proto_rawDescOnce sync.Once
	file_google_actions_sdk_v2_interactionmodel_prompt_static_prompt_proto_rawDescData = file_google_actions_sdk_v2_interactionmodel_prompt_static_prompt_proto_rawDesc
)

func file_google_actions_sdk_v2_interactionmodel_prompt_static_prompt_proto_rawDescGZIP() []byte {
	file_google_actions_sdk_v2_interactionmodel_prompt_static_prompt_proto_rawDescOnce.Do(func() {
		file_google_actions_sdk_v2_interactionmodel_prompt_static_prompt_proto_rawDescData = protoimpl.X.CompressGZIP(file_google_actions_sdk_v2_interactionmodel_prompt_static_prompt_proto_rawDescData)
	})
	return file_google_actions_sdk_v2_interactionmodel_prompt_static_prompt_proto_rawDescData
}

var file_google_actions_sdk_v2_interactionmodel_prompt_static_prompt_proto_msgTypes = make([]protoimpl.MessageInfo, 4)
var file_google_actions_sdk_v2_interactionmodel_prompt_static_prompt_proto_goTypes = []interface{}{
	(*StaticPrompt)(nil),                                            // 0: google.actions.sdk.v2.interactionmodel.prompt.StaticPrompt
	(*StaticPrompt_StaticPromptCandidate)(nil),                      // 1: google.actions.sdk.v2.interactionmodel.prompt.StaticPrompt.StaticPromptCandidate
	(*StaticPrompt_Selector)(nil),                                   // 2: google.actions.sdk.v2.interactionmodel.prompt.StaticPrompt.Selector
	(*StaticPrompt_StaticPromptCandidate_StaticPromptResponse)(nil), // 3: google.actions.sdk.v2.interactionmodel.prompt.StaticPrompt.StaticPromptCandidate.StaticPromptResponse
	(*SurfaceCapabilities)(nil),                                     // 4: google.actions.sdk.v2.interactionmodel.prompt.SurfaceCapabilities
	(*StaticSimplePrompt)(nil),                                      // 5: google.actions.sdk.v2.interactionmodel.prompt.StaticSimplePrompt
	(*StaticContentPrompt)(nil),                                     // 6: google.actions.sdk.v2.interactionmodel.prompt.StaticContentPrompt
	(*Suggestion)(nil),                                              // 7: google.actions.sdk.v2.interactionmodel.prompt.Suggestion
	(*StaticLinkPrompt)(nil),                                        // 8: google.actions.sdk.v2.interactionmodel.prompt.StaticLinkPrompt
	(*StaticCanvasPrompt)(nil),                                      // 9: google.actions.sdk.v2.interactionmodel.prompt.StaticCanvasPrompt
}
var file_google_actions_sdk_v2_interactionmodel_prompt_static_prompt_proto_depIdxs = []int32{
	1,  // 0: google.actions.sdk.v2.interactionmodel.prompt.StaticPrompt.candidates:type_name -> google.actions.sdk.v2.interactionmodel.prompt.StaticPrompt.StaticPromptCandidate
	2,  // 1: google.actions.sdk.v2.interactionmodel.prompt.StaticPrompt.StaticPromptCandidate.selector:type_name -> google.actions.sdk.v2.interactionmodel.prompt.StaticPrompt.Selector
	3,  // 2: google.actions.sdk.v2.interactionmodel.prompt.StaticPrompt.StaticPromptCandidate.prompt_response:type_name -> google.actions.sdk.v2.interactionmodel.prompt.StaticPrompt.StaticPromptCandidate.StaticPromptResponse
	4,  // 3: google.actions.sdk.v2.interactionmodel.prompt.StaticPrompt.Selector.surface_capabilities:type_name -> google.actions.sdk.v2.interactionmodel.prompt.SurfaceCapabilities
	5,  // 4: google.actions.sdk.v2.interactionmodel.prompt.StaticPrompt.StaticPromptCandidate.StaticPromptResponse.first_simple:type_name -> google.actions.sdk.v2.interactionmodel.prompt.StaticSimplePrompt
	6,  // 5: google.actions.sdk.v2.interactionmodel.prompt.StaticPrompt.StaticPromptCandidate.StaticPromptResponse.content:type_name -> google.actions.sdk.v2.interactionmodel.prompt.StaticContentPrompt
	5,  // 6: google.actions.sdk.v2.interactionmodel.prompt.StaticPrompt.StaticPromptCandidate.StaticPromptResponse.last_simple:type_name -> google.actions.sdk.v2.interactionmodel.prompt.StaticSimplePrompt
	7,  // 7: google.actions.sdk.v2.interactionmodel.prompt.StaticPrompt.StaticPromptCandidate.StaticPromptResponse.suggestions:type_name -> google.actions.sdk.v2.interactionmodel.prompt.Suggestion
	8,  // 8: google.actions.sdk.v2.interactionmodel.prompt.StaticPrompt.StaticPromptCandidate.StaticPromptResponse.link:type_name -> google.actions.sdk.v2.interactionmodel.prompt.StaticLinkPrompt
	9,  // 9: google.actions.sdk.v2.interactionmodel.prompt.StaticPrompt.StaticPromptCandidate.StaticPromptResponse.canvas:type_name -> google.actions.sdk.v2.interactionmodel.prompt.StaticCanvasPrompt
	10, // [10:10] is the sub-list for method output_type
	10, // [10:10] is the sub-list for method input_type
	10, // [10:10] is the sub-list for extension type_name
	10, // [10:10] is the sub-list for extension extendee
	0,  // [0:10] is the sub-list for field type_name
}

func init() { file_google_actions_sdk_v2_interactionmodel_prompt_static_prompt_proto_init() }
func file_google_actions_sdk_v2_interactionmodel_prompt_static_prompt_proto_init() {
	if File_google_actions_sdk_v2_interactionmodel_prompt_static_prompt_proto != nil {
		return
	}
	file_google_actions_sdk_v2_interactionmodel_prompt_content_static_canvas_prompt_proto_init()
	file_google_actions_sdk_v2_interactionmodel_prompt_content_static_content_prompt_proto_init()
	file_google_actions_sdk_v2_interactionmodel_prompt_content_static_link_prompt_proto_init()
	file_google_actions_sdk_v2_interactionmodel_prompt_static_simple_prompt_proto_init()
	file_google_actions_sdk_v2_interactionmodel_prompt_suggestion_proto_init()
	file_google_actions_sdk_v2_interactionmodel_prompt_surface_capabilities_proto_init()
	if !protoimpl.UnsafeEnabled {
		file_google_actions_sdk_v2_interactionmodel_prompt_static_prompt_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*StaticPrompt); i {
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
		file_google_actions_sdk_v2_interactionmodel_prompt_static_prompt_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*StaticPrompt_StaticPromptCandidate); i {
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
		file_google_actions_sdk_v2_interactionmodel_prompt_static_prompt_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*StaticPrompt_Selector); i {
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
		file_google_actions_sdk_v2_interactionmodel_prompt_static_prompt_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*StaticPrompt_StaticPromptCandidate_StaticPromptResponse); i {
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
			RawDescriptor: file_google_actions_sdk_v2_interactionmodel_prompt_static_prompt_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   4,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_google_actions_sdk_v2_interactionmodel_prompt_static_prompt_proto_goTypes,
		DependencyIndexes: file_google_actions_sdk_v2_interactionmodel_prompt_static_prompt_proto_depIdxs,
		MessageInfos:      file_google_actions_sdk_v2_interactionmodel_prompt_static_prompt_proto_msgTypes,
	}.Build()
	File_google_actions_sdk_v2_interactionmodel_prompt_static_prompt_proto = out.File
	file_google_actions_sdk_v2_interactionmodel_prompt_static_prompt_proto_rawDesc = nil
	file_google_actions_sdk_v2_interactionmodel_prompt_static_prompt_proto_goTypes = nil
	file_google_actions_sdk_v2_interactionmodel_prompt_static_prompt_proto_depIdxs = nil
}
