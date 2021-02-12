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
// 	protoc-gen-go v1.25.0-devel
// 	protoc        v3.13.0
// source: google/cloud/aiplatform/v1/schema/predict/prediction/image_object_detection.proto

package prediction

import (
	reflect "reflect"
	sync "sync"

	proto "github.com/golang/protobuf/proto"
	_ "google.golang.org/genproto/googleapis/api/annotations"
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

// This is a compile-time assertion that a sufficiently up-to-date version
// of the legacy proto package is being used.
const _ = proto.ProtoPackageIsVersion4

// Prediction output format for Image Object Detection.
type ImageObjectDetectionPredictionResult struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// The resource IDs of the AnnotationSpecs that had been identified, ordered
	// by the confidence score descendingly.
	Ids []int64 `protobuf:"varint,1,rep,packed,name=ids,proto3" json:"ids,omitempty"`
	// The display names of the AnnotationSpecs that had been identified, order
	// matches the IDs.
	DisplayNames []string `protobuf:"bytes,2,rep,name=display_names,json=displayNames,proto3" json:"display_names,omitempty"`
	// The Model's confidences in correctness of the predicted IDs, higher value
	// means higher confidence. Order matches the Ids.
	Confidences []float32 `protobuf:"fixed32,3,rep,packed,name=confidences,proto3" json:"confidences,omitempty"`
	// Bounding boxes, i.e. the rectangles over the image, that pinpoint
	// the found AnnotationSpecs. Given in order that matches the IDs. Each
	// bounding box is an array of 4 numbers `xMin`, `xMax`, `yMin`, and
	// `yMax`, which represent the extremal coordinates of the box. They are
	// relative to the image size, and the point 0,0 is in the top left
	// of the image.
	Bboxes []*structpb.ListValue `protobuf:"bytes,4,rep,name=bboxes,proto3" json:"bboxes,omitempty"`
}

func (x *ImageObjectDetectionPredictionResult) Reset() {
	*x = ImageObjectDetectionPredictionResult{}
	if protoimpl.UnsafeEnabled {
		mi := &file_google_cloud_aiplatform_v1_schema_predict_prediction_image_object_detection_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ImageObjectDetectionPredictionResult) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ImageObjectDetectionPredictionResult) ProtoMessage() {}

func (x *ImageObjectDetectionPredictionResult) ProtoReflect() protoreflect.Message {
	mi := &file_google_cloud_aiplatform_v1_schema_predict_prediction_image_object_detection_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ImageObjectDetectionPredictionResult.ProtoReflect.Descriptor instead.
func (*ImageObjectDetectionPredictionResult) Descriptor() ([]byte, []int) {
	return file_google_cloud_aiplatform_v1_schema_predict_prediction_image_object_detection_proto_rawDescGZIP(), []int{0}
}

func (x *ImageObjectDetectionPredictionResult) GetIds() []int64 {
	if x != nil {
		return x.Ids
	}
	return nil
}

func (x *ImageObjectDetectionPredictionResult) GetDisplayNames() []string {
	if x != nil {
		return x.DisplayNames
	}
	return nil
}

func (x *ImageObjectDetectionPredictionResult) GetConfidences() []float32 {
	if x != nil {
		return x.Confidences
	}
	return nil
}

func (x *ImageObjectDetectionPredictionResult) GetBboxes() []*structpb.ListValue {
	if x != nil {
		return x.Bboxes
	}
	return nil
}

var File_google_cloud_aiplatform_v1_schema_predict_prediction_image_object_detection_proto protoreflect.FileDescriptor

var file_google_cloud_aiplatform_v1_schema_predict_prediction_image_object_detection_proto_rawDesc = []byte{
	0x0a, 0x51, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x63, 0x6c, 0x6f, 0x75, 0x64, 0x2f, 0x61,
	0x69, 0x70, 0x6c, 0x61, 0x74, 0x66, 0x6f, 0x72, 0x6d, 0x2f, 0x76, 0x31, 0x2f, 0x73, 0x63, 0x68,
	0x65, 0x6d, 0x61, 0x2f, 0x70, 0x72, 0x65, 0x64, 0x69, 0x63, 0x74, 0x2f, 0x70, 0x72, 0x65, 0x64,
	0x69, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x2f, 0x69, 0x6d, 0x61, 0x67, 0x65, 0x5f, 0x6f, 0x62, 0x6a,
	0x65, 0x63, 0x74, 0x5f, 0x64, 0x65, 0x74, 0x65, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x2e, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x12, 0x34, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x63, 0x6c, 0x6f, 0x75,
	0x64, 0x2e, 0x61, 0x69, 0x70, 0x6c, 0x61, 0x74, 0x66, 0x6f, 0x72, 0x6d, 0x2e, 0x76, 0x31, 0x2e,
	0x73, 0x63, 0x68, 0x65, 0x6d, 0x61, 0x2e, 0x70, 0x72, 0x65, 0x64, 0x69, 0x63, 0x74, 0x2e, 0x70,
	0x72, 0x65, 0x64, 0x69, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x1a, 0x1c, 0x67, 0x6f, 0x6f, 0x67, 0x6c,
	0x65, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2f, 0x73, 0x74, 0x72, 0x75, 0x63,
	0x74, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x1c, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f,
	0x61, 0x70, 0x69, 0x2f, 0x61, 0x6e, 0x6e, 0x6f, 0x74, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x2e,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0xb3, 0x01, 0x0a, 0x24, 0x49, 0x6d, 0x61, 0x67, 0x65, 0x4f,
	0x62, 0x6a, 0x65, 0x63, 0x74, 0x44, 0x65, 0x74, 0x65, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x50, 0x72,
	0x65, 0x64, 0x69, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x52, 0x65, 0x73, 0x75, 0x6c, 0x74, 0x12, 0x10,
	0x0a, 0x03, 0x69, 0x64, 0x73, 0x18, 0x01, 0x20, 0x03, 0x28, 0x03, 0x52, 0x03, 0x69, 0x64, 0x73,
	0x12, 0x23, 0x0a, 0x0d, 0x64, 0x69, 0x73, 0x70, 0x6c, 0x61, 0x79, 0x5f, 0x6e, 0x61, 0x6d, 0x65,
	0x73, 0x18, 0x02, 0x20, 0x03, 0x28, 0x09, 0x52, 0x0c, 0x64, 0x69, 0x73, 0x70, 0x6c, 0x61, 0x79,
	0x4e, 0x61, 0x6d, 0x65, 0x73, 0x12, 0x20, 0x0a, 0x0b, 0x63, 0x6f, 0x6e, 0x66, 0x69, 0x64, 0x65,
	0x6e, 0x63, 0x65, 0x73, 0x18, 0x03, 0x20, 0x03, 0x28, 0x02, 0x52, 0x0b, 0x63, 0x6f, 0x6e, 0x66,
	0x69, 0x64, 0x65, 0x6e, 0x63, 0x65, 0x73, 0x12, 0x32, 0x0a, 0x06, 0x62, 0x62, 0x6f, 0x78, 0x65,
	0x73, 0x18, 0x04, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x1a, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65,
	0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x4c, 0x69, 0x73, 0x74, 0x56, 0x61,
	0x6c, 0x75, 0x65, 0x52, 0x06, 0x62, 0x62, 0x6f, 0x78, 0x65, 0x73, 0x42, 0xc7, 0x01, 0x0a, 0x38,
	0x63, 0x6f, 0x6d, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x63, 0x6c, 0x6f, 0x75, 0x64,
	0x2e, 0x61, 0x69, 0x70, 0x6c, 0x61, 0x74, 0x66, 0x6f, 0x72, 0x6d, 0x2e, 0x76, 0x31, 0x2e, 0x73,
	0x63, 0x68, 0x65, 0x6d, 0x61, 0x2e, 0x70, 0x72, 0x65, 0x64, 0x69, 0x63, 0x74, 0x2e, 0x70, 0x72,
	0x65, 0x64, 0x69, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x42, 0x29, 0x49, 0x6d, 0x61, 0x67, 0x65, 0x4f,
	0x62, 0x6a, 0x65, 0x63, 0x74, 0x44, 0x65, 0x74, 0x65, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x50, 0x72,
	0x65, 0x64, 0x69, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x52, 0x65, 0x73, 0x75, 0x6c, 0x74, 0x50, 0x72,
	0x6f, 0x74, 0x6f, 0x50, 0x01, 0x5a, 0x5e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x67, 0x6f,
	0x6c, 0x61, 0x6e, 0x67, 0x2e, 0x6f, 0x72, 0x67, 0x2f, 0x67, 0x65, 0x6e, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x2f, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x61, 0x70, 0x69, 0x73, 0x2f, 0x63, 0x6c, 0x6f,
	0x75, 0x64, 0x2f, 0x61, 0x69, 0x70, 0x6c, 0x61, 0x74, 0x66, 0x6f, 0x72, 0x6d, 0x2f, 0x76, 0x31,
	0x2f, 0x73, 0x63, 0x68, 0x65, 0x6d, 0x61, 0x2f, 0x70, 0x72, 0x65, 0x64, 0x69, 0x63, 0x74, 0x2f,
	0x70, 0x72, 0x65, 0x64, 0x69, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x3b, 0x70, 0x72, 0x65, 0x64, 0x69,
	0x63, 0x74, 0x69, 0x6f, 0x6e, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_google_cloud_aiplatform_v1_schema_predict_prediction_image_object_detection_proto_rawDescOnce sync.Once
	file_google_cloud_aiplatform_v1_schema_predict_prediction_image_object_detection_proto_rawDescData = file_google_cloud_aiplatform_v1_schema_predict_prediction_image_object_detection_proto_rawDesc
)

func file_google_cloud_aiplatform_v1_schema_predict_prediction_image_object_detection_proto_rawDescGZIP() []byte {
	file_google_cloud_aiplatform_v1_schema_predict_prediction_image_object_detection_proto_rawDescOnce.Do(func() {
		file_google_cloud_aiplatform_v1_schema_predict_prediction_image_object_detection_proto_rawDescData = protoimpl.X.CompressGZIP(file_google_cloud_aiplatform_v1_schema_predict_prediction_image_object_detection_proto_rawDescData)
	})
	return file_google_cloud_aiplatform_v1_schema_predict_prediction_image_object_detection_proto_rawDescData
}

var file_google_cloud_aiplatform_v1_schema_predict_prediction_image_object_detection_proto_msgTypes = make([]protoimpl.MessageInfo, 1)
var file_google_cloud_aiplatform_v1_schema_predict_prediction_image_object_detection_proto_goTypes = []interface{}{
	(*ImageObjectDetectionPredictionResult)(nil), // 0: google.cloud.aiplatform.v1.schema.predict.prediction.ImageObjectDetectionPredictionResult
	(*structpb.ListValue)(nil),                   // 1: google.protobuf.ListValue
}
var file_google_cloud_aiplatform_v1_schema_predict_prediction_image_object_detection_proto_depIdxs = []int32{
	1, // 0: google.cloud.aiplatform.v1.schema.predict.prediction.ImageObjectDetectionPredictionResult.bboxes:type_name -> google.protobuf.ListValue
	1, // [1:1] is the sub-list for method output_type
	1, // [1:1] is the sub-list for method input_type
	1, // [1:1] is the sub-list for extension type_name
	1, // [1:1] is the sub-list for extension extendee
	0, // [0:1] is the sub-list for field type_name
}

func init() {
	file_google_cloud_aiplatform_v1_schema_predict_prediction_image_object_detection_proto_init()
}
func file_google_cloud_aiplatform_v1_schema_predict_prediction_image_object_detection_proto_init() {
	if File_google_cloud_aiplatform_v1_schema_predict_prediction_image_object_detection_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_google_cloud_aiplatform_v1_schema_predict_prediction_image_object_detection_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ImageObjectDetectionPredictionResult); i {
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
			RawDescriptor: file_google_cloud_aiplatform_v1_schema_predict_prediction_image_object_detection_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   1,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_google_cloud_aiplatform_v1_schema_predict_prediction_image_object_detection_proto_goTypes,
		DependencyIndexes: file_google_cloud_aiplatform_v1_schema_predict_prediction_image_object_detection_proto_depIdxs,
		MessageInfos:      file_google_cloud_aiplatform_v1_schema_predict_prediction_image_object_detection_proto_msgTypes,
	}.Build()
	File_google_cloud_aiplatform_v1_schema_predict_prediction_image_object_detection_proto = out.File
	file_google_cloud_aiplatform_v1_schema_predict_prediction_image_object_detection_proto_rawDesc = nil
	file_google_cloud_aiplatform_v1_schema_predict_prediction_image_object_detection_proto_goTypes = nil
	file_google_cloud_aiplatform_v1_schema_predict_prediction_image_object_detection_proto_depIdxs = nil
}