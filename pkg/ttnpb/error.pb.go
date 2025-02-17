// Copyright © 2019 The Things Network Foundation, The Things Industries B.V.
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
// 	protoc-gen-go v1.30.0
// 	protoc        v4.22.2
// source: ttn/lorawan/v3/error.proto

package ttnpb

import (
	_ "github.com/TheThingsIndustries/protoc-gen-go-flags/annotations"
	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
	anypb "google.golang.org/protobuf/types/known/anypb"
	structpb "google.golang.org/protobuf/types/known/structpb"
	reflect "reflect"
	sync "sync"
)

const (
	// Verify that this generated code is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(20 - protoimpl.MinVersion)
	// Verify that runtime/protoimpl is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(protoimpl.MaxVersion - 20)
)

// Error details that are communicated over gRPC (and HTTP) APIs.
// The messages (for translation) are stored as "error:<namespace>:<name>".
type ErrorDetails struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// Namespace of the error (typically the package name in The Things Stack).
	Namespace string `protobuf:"bytes,1,opt,name=namespace,proto3" json:"namespace,omitempty"`
	// Name of the error.
	Name string `protobuf:"bytes,2,opt,name=name,proto3" json:"name,omitempty"`
	// The default (fallback) message format that should be used for the error.
	// This is also used if the client does not have a translation for the error.
	MessageFormat string `protobuf:"bytes,3,opt,name=message_format,json=messageFormat,proto3" json:"message_format,omitempty"`
	// Attributes that should be filled into the message format. Any extra attributes
	// can be displayed as error details.
	Attributes *structpb.Struct `protobuf:"bytes,4,opt,name=attributes,proto3" json:"attributes,omitempty"`
	// The correlation ID of the error can be used to correlate the error to stack
	// traces the network may (or may not) store about recent errors.
	CorrelationId string `protobuf:"bytes,5,opt,name=correlation_id,json=correlationId,proto3" json:"correlation_id,omitempty"`
	// The error that caused this error.
	Cause *ErrorDetails `protobuf:"bytes,6,opt,name=cause,proto3" json:"cause,omitempty"`
	// The status code of the error.
	Code uint32 `protobuf:"varint,7,opt,name=code,proto3" json:"code,omitempty"`
	// The details of the error.
	Details []*anypb.Any `protobuf:"bytes,8,rep,name=details,proto3" json:"details,omitempty"`
}

func (x *ErrorDetails) Reset() {
	*x = ErrorDetails{}
	if protoimpl.UnsafeEnabled {
		mi := &file_ttn_lorawan_v3_error_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ErrorDetails) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ErrorDetails) ProtoMessage() {}

func (x *ErrorDetails) ProtoReflect() protoreflect.Message {
	mi := &file_ttn_lorawan_v3_error_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ErrorDetails.ProtoReflect.Descriptor instead.
func (*ErrorDetails) Descriptor() ([]byte, []int) {
	return file_ttn_lorawan_v3_error_proto_rawDescGZIP(), []int{0}
}

func (x *ErrorDetails) GetNamespace() string {
	if x != nil {
		return x.Namespace
	}
	return ""
}

func (x *ErrorDetails) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

func (x *ErrorDetails) GetMessageFormat() string {
	if x != nil {
		return x.MessageFormat
	}
	return ""
}

func (x *ErrorDetails) GetAttributes() *structpb.Struct {
	if x != nil {
		return x.Attributes
	}
	return nil
}

func (x *ErrorDetails) GetCorrelationId() string {
	if x != nil {
		return x.CorrelationId
	}
	return ""
}

func (x *ErrorDetails) GetCause() *ErrorDetails {
	if x != nil {
		return x.Cause
	}
	return nil
}

func (x *ErrorDetails) GetCode() uint32 {
	if x != nil {
		return x.Code
	}
	return 0
}

func (x *ErrorDetails) GetDetails() []*anypb.Any {
	if x != nil {
		return x.Details
	}
	return nil
}

var File_ttn_lorawan_v3_error_proto protoreflect.FileDescriptor

var file_ttn_lorawan_v3_error_proto_rawDesc = []byte{
	0x0a, 0x1a, 0x74, 0x74, 0x6e, 0x2f, 0x6c, 0x6f, 0x72, 0x61, 0x77, 0x61, 0x6e, 0x2f, 0x76, 0x33,
	0x2f, 0x65, 0x72, 0x72, 0x6f, 0x72, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x0e, 0x74, 0x74,
	0x6e, 0x2e, 0x6c, 0x6f, 0x72, 0x61, 0x77, 0x61, 0x6e, 0x2e, 0x76, 0x33, 0x1a, 0x19, 0x67, 0x6f,
	0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2f, 0x61, 0x6e,
	0x79, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x1c, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2f, 0x73, 0x74, 0x72, 0x75, 0x63, 0x74, 0x2e,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x21, 0x74, 0x68, 0x65, 0x74, 0x68, 0x69, 0x6e, 0x67, 0x73,
	0x2f, 0x66, 0x6c, 0x61, 0x67, 0x73, 0x2f, 0x61, 0x6e, 0x6e, 0x6f, 0x74, 0x61, 0x74, 0x69, 0x6f,
	0x6e, 0x73, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0xd1, 0x02, 0x0a, 0x0c, 0x45, 0x72, 0x72,
	0x6f, 0x72, 0x44, 0x65, 0x74, 0x61, 0x69, 0x6c, 0x73, 0x12, 0x1c, 0x0a, 0x09, 0x6e, 0x61, 0x6d,
	0x65, 0x73, 0x70, 0x61, 0x63, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x09, 0x6e, 0x61,
	0x6d, 0x65, 0x73, 0x70, 0x61, 0x63, 0x65, 0x12, 0x12, 0x0a, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x18,
	0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x12, 0x25, 0x0a, 0x0e, 0x6d,
	0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x5f, 0x66, 0x6f, 0x72, 0x6d, 0x61, 0x74, 0x18, 0x03, 0x20,
	0x01, 0x28, 0x09, 0x52, 0x0d, 0x6d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x46, 0x6f, 0x72, 0x6d,
	0x61, 0x74, 0x12, 0x37, 0x0a, 0x0a, 0x61, 0x74, 0x74, 0x72, 0x69, 0x62, 0x75, 0x74, 0x65, 0x73,
	0x18, 0x04, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x17, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x53, 0x74, 0x72, 0x75, 0x63, 0x74, 0x52,
	0x0a, 0x61, 0x74, 0x74, 0x72, 0x69, 0x62, 0x75, 0x74, 0x65, 0x73, 0x12, 0x25, 0x0a, 0x0e, 0x63,
	0x6f, 0x72, 0x72, 0x65, 0x6c, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x5f, 0x69, 0x64, 0x18, 0x05, 0x20,
	0x01, 0x28, 0x09, 0x52, 0x0d, 0x63, 0x6f, 0x72, 0x72, 0x65, 0x6c, 0x61, 0x74, 0x69, 0x6f, 0x6e,
	0x49, 0x64, 0x12, 0x3a, 0x0a, 0x05, 0x63, 0x61, 0x75, 0x73, 0x65, 0x18, 0x06, 0x20, 0x01, 0x28,
	0x0b, 0x32, 0x1c, 0x2e, 0x74, 0x74, 0x6e, 0x2e, 0x6c, 0x6f, 0x72, 0x61, 0x77, 0x61, 0x6e, 0x2e,
	0x76, 0x33, 0x2e, 0x45, 0x72, 0x72, 0x6f, 0x72, 0x44, 0x65, 0x74, 0x61, 0x69, 0x6c, 0x73, 0x42,
	0x06, 0xf2, 0xaa, 0x19, 0x02, 0x08, 0x00, 0x52, 0x05, 0x63, 0x61, 0x75, 0x73, 0x65, 0x12, 0x12,
	0x0a, 0x04, 0x63, 0x6f, 0x64, 0x65, 0x18, 0x07, 0x20, 0x01, 0x28, 0x0d, 0x52, 0x04, 0x63, 0x6f,
	0x64, 0x65, 0x12, 0x2e, 0x0a, 0x07, 0x64, 0x65, 0x74, 0x61, 0x69, 0x6c, 0x73, 0x18, 0x08, 0x20,
	0x03, 0x28, 0x0b, 0x32, 0x14, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x41, 0x6e, 0x79, 0x52, 0x07, 0x64, 0x65, 0x74, 0x61, 0x69,
	0x6c, 0x73, 0x3a, 0x08, 0xf2, 0xaa, 0x19, 0x04, 0x08, 0x01, 0x10, 0x00, 0x42, 0x31, 0x5a, 0x2f,
	0x67, 0x6f, 0x2e, 0x74, 0x68, 0x65, 0x74, 0x68, 0x69, 0x6e, 0x67, 0x73, 0x2e, 0x6e, 0x65, 0x74,
	0x77, 0x6f, 0x72, 0x6b, 0x2f, 0x6c, 0x6f, 0x72, 0x61, 0x77, 0x61, 0x6e, 0x2d, 0x73, 0x74, 0x61,
	0x63, 0x6b, 0x2f, 0x76, 0x33, 0x2f, 0x70, 0x6b, 0x67, 0x2f, 0x74, 0x74, 0x6e, 0x70, 0x62, 0x62,
	0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_ttn_lorawan_v3_error_proto_rawDescOnce sync.Once
	file_ttn_lorawan_v3_error_proto_rawDescData = file_ttn_lorawan_v3_error_proto_rawDesc
)

func file_ttn_lorawan_v3_error_proto_rawDescGZIP() []byte {
	file_ttn_lorawan_v3_error_proto_rawDescOnce.Do(func() {
		file_ttn_lorawan_v3_error_proto_rawDescData = protoimpl.X.CompressGZIP(file_ttn_lorawan_v3_error_proto_rawDescData)
	})
	return file_ttn_lorawan_v3_error_proto_rawDescData
}

var file_ttn_lorawan_v3_error_proto_msgTypes = make([]protoimpl.MessageInfo, 1)
var file_ttn_lorawan_v3_error_proto_goTypes = []interface{}{
	(*ErrorDetails)(nil),    // 0: ttn.lorawan.v3.ErrorDetails
	(*structpb.Struct)(nil), // 1: google.protobuf.Struct
	(*anypb.Any)(nil),       // 2: google.protobuf.Any
}
var file_ttn_lorawan_v3_error_proto_depIdxs = []int32{
	1, // 0: ttn.lorawan.v3.ErrorDetails.attributes:type_name -> google.protobuf.Struct
	0, // 1: ttn.lorawan.v3.ErrorDetails.cause:type_name -> ttn.lorawan.v3.ErrorDetails
	2, // 2: ttn.lorawan.v3.ErrorDetails.details:type_name -> google.protobuf.Any
	3, // [3:3] is the sub-list for method output_type
	3, // [3:3] is the sub-list for method input_type
	3, // [3:3] is the sub-list for extension type_name
	3, // [3:3] is the sub-list for extension extendee
	0, // [0:3] is the sub-list for field type_name
}

func init() { file_ttn_lorawan_v3_error_proto_init() }
func file_ttn_lorawan_v3_error_proto_init() {
	if File_ttn_lorawan_v3_error_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_ttn_lorawan_v3_error_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ErrorDetails); i {
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
			RawDescriptor: file_ttn_lorawan_v3_error_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   1,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_ttn_lorawan_v3_error_proto_goTypes,
		DependencyIndexes: file_ttn_lorawan_v3_error_proto_depIdxs,
		MessageInfos:      file_ttn_lorawan_v3_error_proto_msgTypes,
	}.Build()
	File_ttn_lorawan_v3_error_proto = out.File
	file_ttn_lorawan_v3_error_proto_rawDesc = nil
	file_ttn_lorawan_v3_error_proto_goTypes = nil
	file_ttn_lorawan_v3_error_proto_depIdxs = nil
}
