// Copyright 2022 Google LLC
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
// 	protoc-gen-go v1.28.0
// 	protoc        v3.21.7
// source: google/cloud/discoveryengine/v1alpha/schema_service.proto

package discoveryengine

import (
	context "context"
	_ "google.golang.org/genproto/googleapis/api/annotations"
	grpc "google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
	emptypb "google.golang.org/protobuf/types/known/emptypb"
	reflect "reflect"
	sync "sync"
)

const (
	// Verify that this generated code is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(20 - protoimpl.MinVersion)
	// Verify that runtime/protoimpl is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(protoimpl.MaxVersion - 20)
)

// Request message for
// [SchemaService.GetSchema][google.cloud.discoveryengine.v1alpha.SchemaService.GetSchema]
// method.
type GetSchemaRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// Required. The full resource name of the schema, in the format of
	// `projects/{project}/locations/{location}/dataStores/{data_store}/schemas/{schema}`.
	Name string `protobuf:"bytes,1,opt,name=name,proto3" json:"name,omitempty"`
}

func (x *GetSchemaRequest) Reset() {
	*x = GetSchemaRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_google_cloud_discoveryengine_v1alpha_schema_service_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetSchemaRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetSchemaRequest) ProtoMessage() {}

func (x *GetSchemaRequest) ProtoReflect() protoreflect.Message {
	mi := &file_google_cloud_discoveryengine_v1alpha_schema_service_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetSchemaRequest.ProtoReflect.Descriptor instead.
func (*GetSchemaRequest) Descriptor() ([]byte, []int) {
	return file_google_cloud_discoveryengine_v1alpha_schema_service_proto_rawDescGZIP(), []int{0}
}

func (x *GetSchemaRequest) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

// Request message for
// [SchemaService.ListSchemas][google.cloud.discoveryengine.v1alpha.SchemaService.ListSchemas]
// method.
type ListSchemasRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// Required. The parent data store resource name, in the format of
	// `projects/{project}/locations/{location}/dataStores/{data_store}`.
	Parent string `protobuf:"bytes,1,opt,name=parent,proto3" json:"parent,omitempty"`
	// The maximum number of
	// [Schema][google.cloud.discoveryengine.v1alpha.Schema]s to return. The
	// service may return fewer than this value.
	//
	// If unspecified, at most 100
	// [Schema][google.cloud.discoveryengine.v1alpha.Schema]s will be returned.
	//
	// The maximum value is 1000; values above 1000 will be coerced to 1000.
	PageSize int32 `protobuf:"varint,2,opt,name=page_size,json=pageSize,proto3" json:"page_size,omitempty"`
	// A page token, received from a previous
	// [SchemaService.ListSchemas][google.cloud.discoveryengine.v1alpha.SchemaService.ListSchemas]
	// call. Provide this to retrieve the subsequent page.
	//
	// When paginating, all other parameters provided to
	// [SchemaService.ListSchemas][google.cloud.discoveryengine.v1alpha.SchemaService.ListSchemas]
	// must match the call that provided the page token.
	PageToken string `protobuf:"bytes,3,opt,name=page_token,json=pageToken,proto3" json:"page_token,omitempty"`
}

func (x *ListSchemasRequest) Reset() {
	*x = ListSchemasRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_google_cloud_discoveryengine_v1alpha_schema_service_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ListSchemasRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ListSchemasRequest) ProtoMessage() {}

func (x *ListSchemasRequest) ProtoReflect() protoreflect.Message {
	mi := &file_google_cloud_discoveryengine_v1alpha_schema_service_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ListSchemasRequest.ProtoReflect.Descriptor instead.
func (*ListSchemasRequest) Descriptor() ([]byte, []int) {
	return file_google_cloud_discoveryengine_v1alpha_schema_service_proto_rawDescGZIP(), []int{1}
}

func (x *ListSchemasRequest) GetParent() string {
	if x != nil {
		return x.Parent
	}
	return ""
}

func (x *ListSchemasRequest) GetPageSize() int32 {
	if x != nil {
		return x.PageSize
	}
	return 0
}

func (x *ListSchemasRequest) GetPageToken() string {
	if x != nil {
		return x.PageToken
	}
	return ""
}

// Response message for
// [SchemaService.ListSchemas][google.cloud.discoveryengine.v1alpha.SchemaService.ListSchemas]
// method.
type ListSchemasResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// The [Schema][google.cloud.discoveryengine.v1alpha.Schema]s.
	Schemas []*Schema `protobuf:"bytes,1,rep,name=schemas,proto3" json:"schemas,omitempty"`
	// A token that can be sent as
	// [ListSchemasRequest.page_token][google.cloud.discoveryengine.v1alpha.ListSchemasRequest.page_token]
	// to retrieve the next page. If this field is omitted, there are no
	// subsequent pages.
	NextPageToken string `protobuf:"bytes,2,opt,name=next_page_token,json=nextPageToken,proto3" json:"next_page_token,omitempty"`
}

func (x *ListSchemasResponse) Reset() {
	*x = ListSchemasResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_google_cloud_discoveryengine_v1alpha_schema_service_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ListSchemasResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ListSchemasResponse) ProtoMessage() {}

func (x *ListSchemasResponse) ProtoReflect() protoreflect.Message {
	mi := &file_google_cloud_discoveryengine_v1alpha_schema_service_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ListSchemasResponse.ProtoReflect.Descriptor instead.
func (*ListSchemasResponse) Descriptor() ([]byte, []int) {
	return file_google_cloud_discoveryengine_v1alpha_schema_service_proto_rawDescGZIP(), []int{2}
}

func (x *ListSchemasResponse) GetSchemas() []*Schema {
	if x != nil {
		return x.Schemas
	}
	return nil
}

func (x *ListSchemasResponse) GetNextPageToken() string {
	if x != nil {
		return x.NextPageToken
	}
	return ""
}

// Request message for
// [SchemaService.CreateSchema][google.cloud.discoveryengine.v1alpha.SchemaService.CreateSchema]
// method.
type CreateSchemaRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// Required. The parent data store resource name, in the format of
	// `projects/{project}/locations/{location}/dataStores/{data_store}`.
	Parent string `protobuf:"bytes,1,opt,name=parent,proto3" json:"parent,omitempty"`
	// Required. The [Schema][google.cloud.discoveryengine.v1alpha.Schema] to
	// create.
	Schema *Schema `protobuf:"bytes,2,opt,name=schema,proto3" json:"schema,omitempty"`
	// Required. The ID to use for the
	// [Schema][google.cloud.discoveryengine.v1alpha.Schema], which will become
	// the final component of the
	// [Schema.name][google.cloud.discoveryengine.v1alpha.Schema.name].
	//
	// This field should conform to
	// [RFC-1034](https://tools.ietf.org/html/rfc1034) standard with a length
	// limit of 63 characters.
	SchemaId string `protobuf:"bytes,3,opt,name=schema_id,json=schemaId,proto3" json:"schema_id,omitempty"`
}

func (x *CreateSchemaRequest) Reset() {
	*x = CreateSchemaRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_google_cloud_discoveryengine_v1alpha_schema_service_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *CreateSchemaRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CreateSchemaRequest) ProtoMessage() {}

func (x *CreateSchemaRequest) ProtoReflect() protoreflect.Message {
	mi := &file_google_cloud_discoveryengine_v1alpha_schema_service_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CreateSchemaRequest.ProtoReflect.Descriptor instead.
func (*CreateSchemaRequest) Descriptor() ([]byte, []int) {
	return file_google_cloud_discoveryengine_v1alpha_schema_service_proto_rawDescGZIP(), []int{3}
}

func (x *CreateSchemaRequest) GetParent() string {
	if x != nil {
		return x.Parent
	}
	return ""
}

func (x *CreateSchemaRequest) GetSchema() *Schema {
	if x != nil {
		return x.Schema
	}
	return nil
}

func (x *CreateSchemaRequest) GetSchemaId() string {
	if x != nil {
		return x.SchemaId
	}
	return ""
}

// Request message for
// [SchemaService.DeleteSchema][google.cloud.discoveryengine.v1alpha.SchemaService.DeleteSchema]
// method.
type DeleteSchemaRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// Required. The full resource name of the schema, in the format of
	// `projects/{project}/locations/{location}/dataStores/{data_store}/schemas/{schema}`.
	Name string `protobuf:"bytes,1,opt,name=name,proto3" json:"name,omitempty"`
}

func (x *DeleteSchemaRequest) Reset() {
	*x = DeleteSchemaRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_google_cloud_discoveryengine_v1alpha_schema_service_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *DeleteSchemaRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*DeleteSchemaRequest) ProtoMessage() {}

func (x *DeleteSchemaRequest) ProtoReflect() protoreflect.Message {
	mi := &file_google_cloud_discoveryengine_v1alpha_schema_service_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use DeleteSchemaRequest.ProtoReflect.Descriptor instead.
func (*DeleteSchemaRequest) Descriptor() ([]byte, []int) {
	return file_google_cloud_discoveryengine_v1alpha_schema_service_proto_rawDescGZIP(), []int{4}
}

func (x *DeleteSchemaRequest) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

var File_google_cloud_discoveryengine_v1alpha_schema_service_proto protoreflect.FileDescriptor

var file_google_cloud_discoveryengine_v1alpha_schema_service_proto_rawDesc = []byte{
	0x0a, 0x39, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x63, 0x6c, 0x6f, 0x75, 0x64, 0x2f, 0x64,
	0x69, 0x73, 0x63, 0x6f, 0x76, 0x65, 0x72, 0x79, 0x65, 0x6e, 0x67, 0x69, 0x6e, 0x65, 0x2f, 0x76,
	0x31, 0x61, 0x6c, 0x70, 0x68, 0x61, 0x2f, 0x73, 0x63, 0x68, 0x65, 0x6d, 0x61, 0x5f, 0x73, 0x65,
	0x72, 0x76, 0x69, 0x63, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x24, 0x67, 0x6f, 0x6f,
	0x67, 0x6c, 0x65, 0x2e, 0x63, 0x6c, 0x6f, 0x75, 0x64, 0x2e, 0x64, 0x69, 0x73, 0x63, 0x6f, 0x76,
	0x65, 0x72, 0x79, 0x65, 0x6e, 0x67, 0x69, 0x6e, 0x65, 0x2e, 0x76, 0x31, 0x61, 0x6c, 0x70, 0x68,
	0x61, 0x1a, 0x1c, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x61, 0x70, 0x69, 0x2f, 0x61, 0x6e,
	0x6e, 0x6f, 0x74, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a,
	0x17, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x61, 0x70, 0x69, 0x2f, 0x63, 0x6c, 0x69, 0x65,
	0x6e, 0x74, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x1f, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65,
	0x2f, 0x61, 0x70, 0x69, 0x2f, 0x66, 0x69, 0x65, 0x6c, 0x64, 0x5f, 0x62, 0x65, 0x68, 0x61, 0x76,
	0x69, 0x6f, 0x72, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x19, 0x67, 0x6f, 0x6f, 0x67, 0x6c,
	0x65, 0x2f, 0x61, 0x70, 0x69, 0x2f, 0x72, 0x65, 0x73, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x2e, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x31, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x63, 0x6c, 0x6f,
	0x75, 0x64, 0x2f, 0x64, 0x69, 0x73, 0x63, 0x6f, 0x76, 0x65, 0x72, 0x79, 0x65, 0x6e, 0x67, 0x69,
	0x6e, 0x65, 0x2f, 0x76, 0x31, 0x61, 0x6c, 0x70, 0x68, 0x61, 0x2f, 0x73, 0x63, 0x68, 0x65, 0x6d,
	0x61, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x1b, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2f, 0x65, 0x6d, 0x70, 0x74, 0x79, 0x2e, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x22, 0x55, 0x0a, 0x10, 0x47, 0x65, 0x74, 0x53, 0x63, 0x68, 0x65, 0x6d,
	0x61, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x41, 0x0a, 0x04, 0x6e, 0x61, 0x6d, 0x65,
	0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x42, 0x2d, 0xe0, 0x41, 0x02, 0xfa, 0x41, 0x27, 0x0a, 0x25,
	0x64, 0x69, 0x73, 0x63, 0x6f, 0x76, 0x65, 0x72, 0x79, 0x65, 0x6e, 0x67, 0x69, 0x6e, 0x65, 0x2e,
	0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x61, 0x70, 0x69, 0x73, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x53,
	0x63, 0x68, 0x65, 0x6d, 0x61, 0x52, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x22, 0x9a, 0x01, 0x0a, 0x12,
	0x4c, 0x69, 0x73, 0x74, 0x53, 0x63, 0x68, 0x65, 0x6d, 0x61, 0x73, 0x52, 0x65, 0x71, 0x75, 0x65,
	0x73, 0x74, 0x12, 0x48, 0x0a, 0x06, 0x70, 0x61, 0x72, 0x65, 0x6e, 0x74, 0x18, 0x01, 0x20, 0x01,
	0x28, 0x09, 0x42, 0x30, 0xe0, 0x41, 0x02, 0xfa, 0x41, 0x2a, 0x0a, 0x28, 0x64, 0x69, 0x73, 0x63,
	0x6f, 0x76, 0x65, 0x72, 0x79, 0x65, 0x6e, 0x67, 0x69, 0x6e, 0x65, 0x2e, 0x67, 0x6f, 0x6f, 0x67,
	0x6c, 0x65, 0x61, 0x70, 0x69, 0x73, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x44, 0x61, 0x74, 0x61, 0x53,
	0x74, 0x6f, 0x72, 0x65, 0x52, 0x06, 0x70, 0x61, 0x72, 0x65, 0x6e, 0x74, 0x12, 0x1b, 0x0a, 0x09,
	0x70, 0x61, 0x67, 0x65, 0x5f, 0x73, 0x69, 0x7a, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x05, 0x52,
	0x08, 0x70, 0x61, 0x67, 0x65, 0x53, 0x69, 0x7a, 0x65, 0x12, 0x1d, 0x0a, 0x0a, 0x70, 0x61, 0x67,
	0x65, 0x5f, 0x74, 0x6f, 0x6b, 0x65, 0x6e, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x09, 0x70,
	0x61, 0x67, 0x65, 0x54, 0x6f, 0x6b, 0x65, 0x6e, 0x22, 0x85, 0x01, 0x0a, 0x13, 0x4c, 0x69, 0x73,
	0x74, 0x53, 0x63, 0x68, 0x65, 0x6d, 0x61, 0x73, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65,
	0x12, 0x46, 0x0a, 0x07, 0x73, 0x63, 0x68, 0x65, 0x6d, 0x61, 0x73, 0x18, 0x01, 0x20, 0x03, 0x28,
	0x0b, 0x32, 0x2c, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x63, 0x6c, 0x6f, 0x75, 0x64,
	0x2e, 0x64, 0x69, 0x73, 0x63, 0x6f, 0x76, 0x65, 0x72, 0x79, 0x65, 0x6e, 0x67, 0x69, 0x6e, 0x65,
	0x2e, 0x76, 0x31, 0x61, 0x6c, 0x70, 0x68, 0x61, 0x2e, 0x53, 0x63, 0x68, 0x65, 0x6d, 0x61, 0x52,
	0x07, 0x73, 0x63, 0x68, 0x65, 0x6d, 0x61, 0x73, 0x12, 0x26, 0x0a, 0x0f, 0x6e, 0x65, 0x78, 0x74,
	0x5f, 0x70, 0x61, 0x67, 0x65, 0x5f, 0x74, 0x6f, 0x6b, 0x65, 0x6e, 0x18, 0x02, 0x20, 0x01, 0x28,
	0x09, 0x52, 0x0d, 0x6e, 0x65, 0x78, 0x74, 0x50, 0x61, 0x67, 0x65, 0x54, 0x6f, 0x6b, 0x65, 0x6e,
	0x22, 0xcc, 0x01, 0x0a, 0x13, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x53, 0x63, 0x68, 0x65, 0x6d,
	0x61, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x48, 0x0a, 0x06, 0x70, 0x61, 0x72, 0x65,
	0x6e, 0x74, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x42, 0x30, 0xe0, 0x41, 0x02, 0xfa, 0x41, 0x2a,
	0x0a, 0x28, 0x64, 0x69, 0x73, 0x63, 0x6f, 0x76, 0x65, 0x72, 0x79, 0x65, 0x6e, 0x67, 0x69, 0x6e,
	0x65, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x61, 0x70, 0x69, 0x73, 0x2e, 0x63, 0x6f, 0x6d,
	0x2f, 0x44, 0x61, 0x74, 0x61, 0x53, 0x74, 0x6f, 0x72, 0x65, 0x52, 0x06, 0x70, 0x61, 0x72, 0x65,
	0x6e, 0x74, 0x12, 0x49, 0x0a, 0x06, 0x73, 0x63, 0x68, 0x65, 0x6d, 0x61, 0x18, 0x02, 0x20, 0x01,
	0x28, 0x0b, 0x32, 0x2c, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x63, 0x6c, 0x6f, 0x75,
	0x64, 0x2e, 0x64, 0x69, 0x73, 0x63, 0x6f, 0x76, 0x65, 0x72, 0x79, 0x65, 0x6e, 0x67, 0x69, 0x6e,
	0x65, 0x2e, 0x76, 0x31, 0x61, 0x6c, 0x70, 0x68, 0x61, 0x2e, 0x53, 0x63, 0x68, 0x65, 0x6d, 0x61,
	0x42, 0x03, 0xe0, 0x41, 0x02, 0x52, 0x06, 0x73, 0x63, 0x68, 0x65, 0x6d, 0x61, 0x12, 0x20, 0x0a,
	0x09, 0x73, 0x63, 0x68, 0x65, 0x6d, 0x61, 0x5f, 0x69, 0x64, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09,
	0x42, 0x03, 0xe0, 0x41, 0x02, 0x52, 0x08, 0x73, 0x63, 0x68, 0x65, 0x6d, 0x61, 0x49, 0x64, 0x22,
	0x58, 0x0a, 0x13, 0x44, 0x65, 0x6c, 0x65, 0x74, 0x65, 0x53, 0x63, 0x68, 0x65, 0x6d, 0x61, 0x52,
	0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x41, 0x0a, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x01,
	0x20, 0x01, 0x28, 0x09, 0x42, 0x2d, 0xe0, 0x41, 0x02, 0xfa, 0x41, 0x27, 0x0a, 0x25, 0x64, 0x69,
	0x73, 0x63, 0x6f, 0x76, 0x65, 0x72, 0x79, 0x65, 0x6e, 0x67, 0x69, 0x6e, 0x65, 0x2e, 0x67, 0x6f,
	0x6f, 0x67, 0x6c, 0x65, 0x61, 0x70, 0x69, 0x73, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x53, 0x63, 0x68,
	0x65, 0x6d, 0x61, 0x52, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x32, 0x8f, 0x07, 0x0a, 0x0d, 0x53, 0x63,
	0x68, 0x65, 0x6d, 0x61, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x12, 0xbf, 0x01, 0x0a, 0x09,
	0x47, 0x65, 0x74, 0x53, 0x63, 0x68, 0x65, 0x6d, 0x61, 0x12, 0x36, 0x2e, 0x67, 0x6f, 0x6f, 0x67,
	0x6c, 0x65, 0x2e, 0x63, 0x6c, 0x6f, 0x75, 0x64, 0x2e, 0x64, 0x69, 0x73, 0x63, 0x6f, 0x76, 0x65,
	0x72, 0x79, 0x65, 0x6e, 0x67, 0x69, 0x6e, 0x65, 0x2e, 0x76, 0x31, 0x61, 0x6c, 0x70, 0x68, 0x61,
	0x2e, 0x47, 0x65, 0x74, 0x53, 0x63, 0x68, 0x65, 0x6d, 0x61, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73,
	0x74, 0x1a, 0x2c, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x63, 0x6c, 0x6f, 0x75, 0x64,
	0x2e, 0x64, 0x69, 0x73, 0x63, 0x6f, 0x76, 0x65, 0x72, 0x79, 0x65, 0x6e, 0x67, 0x69, 0x6e, 0x65,
	0x2e, 0x76, 0x31, 0x61, 0x6c, 0x70, 0x68, 0x61, 0x2e, 0x53, 0x63, 0x68, 0x65, 0x6d, 0x61, 0x22,
	0x4c, 0x82, 0xd3, 0xe4, 0x93, 0x02, 0x3f, 0x12, 0x3d, 0x2f, 0x76, 0x31, 0x61, 0x6c, 0x70, 0x68,
	0x61, 0x2f, 0x7b, 0x6e, 0x61, 0x6d, 0x65, 0x3d, 0x70, 0x72, 0x6f, 0x6a, 0x65, 0x63, 0x74, 0x73,
	0x2f, 0x2a, 0x2f, 0x6c, 0x6f, 0x63, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x2f, 0x2a, 0x2f, 0x64,
	0x61, 0x74, 0x61, 0x53, 0x74, 0x6f, 0x72, 0x65, 0x73, 0x2f, 0x2a, 0x2f, 0x73, 0x63, 0x68, 0x65,
	0x6d, 0x61, 0x73, 0x2f, 0x2a, 0x7d, 0xda, 0x41, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x12, 0xd2, 0x01,
	0x0a, 0x0b, 0x4c, 0x69, 0x73, 0x74, 0x53, 0x63, 0x68, 0x65, 0x6d, 0x61, 0x73, 0x12, 0x38, 0x2e,
	0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x63, 0x6c, 0x6f, 0x75, 0x64, 0x2e, 0x64, 0x69, 0x73,
	0x63, 0x6f, 0x76, 0x65, 0x72, 0x79, 0x65, 0x6e, 0x67, 0x69, 0x6e, 0x65, 0x2e, 0x76, 0x31, 0x61,
	0x6c, 0x70, 0x68, 0x61, 0x2e, 0x4c, 0x69, 0x73, 0x74, 0x53, 0x63, 0x68, 0x65, 0x6d, 0x61, 0x73,
	0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x39, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65,
	0x2e, 0x63, 0x6c, 0x6f, 0x75, 0x64, 0x2e, 0x64, 0x69, 0x73, 0x63, 0x6f, 0x76, 0x65, 0x72, 0x79,
	0x65, 0x6e, 0x67, 0x69, 0x6e, 0x65, 0x2e, 0x76, 0x31, 0x61, 0x6c, 0x70, 0x68, 0x61, 0x2e, 0x4c,
	0x69, 0x73, 0x74, 0x53, 0x63, 0x68, 0x65, 0x6d, 0x61, 0x73, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e,
	0x73, 0x65, 0x22, 0x4e, 0x82, 0xd3, 0xe4, 0x93, 0x02, 0x3f, 0x12, 0x3d, 0x2f, 0x76, 0x31, 0x61,
	0x6c, 0x70, 0x68, 0x61, 0x2f, 0x7b, 0x70, 0x61, 0x72, 0x65, 0x6e, 0x74, 0x3d, 0x70, 0x72, 0x6f,
	0x6a, 0x65, 0x63, 0x74, 0x73, 0x2f, 0x2a, 0x2f, 0x6c, 0x6f, 0x63, 0x61, 0x74, 0x69, 0x6f, 0x6e,
	0x73, 0x2f, 0x2a, 0x2f, 0x64, 0x61, 0x74, 0x61, 0x53, 0x74, 0x6f, 0x72, 0x65, 0x73, 0x2f, 0x2a,
	0x7d, 0x2f, 0x73, 0x63, 0x68, 0x65, 0x6d, 0x61, 0x73, 0xda, 0x41, 0x06, 0x70, 0x61, 0x72, 0x65,
	0x6e, 0x74, 0x12, 0xe0, 0x01, 0x0a, 0x0c, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x53, 0x63, 0x68,
	0x65, 0x6d, 0x61, 0x12, 0x39, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x63, 0x6c, 0x6f,
	0x75, 0x64, 0x2e, 0x64, 0x69, 0x73, 0x63, 0x6f, 0x76, 0x65, 0x72, 0x79, 0x65, 0x6e, 0x67, 0x69,
	0x6e, 0x65, 0x2e, 0x76, 0x31, 0x61, 0x6c, 0x70, 0x68, 0x61, 0x2e, 0x43, 0x72, 0x65, 0x61, 0x74,
	0x65, 0x53, 0x63, 0x68, 0x65, 0x6d, 0x61, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x2c,
	0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x63, 0x6c, 0x6f, 0x75, 0x64, 0x2e, 0x64, 0x69,
	0x73, 0x63, 0x6f, 0x76, 0x65, 0x72, 0x79, 0x65, 0x6e, 0x67, 0x69, 0x6e, 0x65, 0x2e, 0x76, 0x31,
	0x61, 0x6c, 0x70, 0x68, 0x61, 0x2e, 0x53, 0x63, 0x68, 0x65, 0x6d, 0x61, 0x22, 0x67, 0x82, 0xd3,
	0xe4, 0x93, 0x02, 0x47, 0x22, 0x3d, 0x2f, 0x76, 0x31, 0x61, 0x6c, 0x70, 0x68, 0x61, 0x2f, 0x7b,
	0x70, 0x61, 0x72, 0x65, 0x6e, 0x74, 0x3d, 0x70, 0x72, 0x6f, 0x6a, 0x65, 0x63, 0x74, 0x73, 0x2f,
	0x2a, 0x2f, 0x6c, 0x6f, 0x63, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x2f, 0x2a, 0x2f, 0x64, 0x61,
	0x74, 0x61, 0x53, 0x74, 0x6f, 0x72, 0x65, 0x73, 0x2f, 0x2a, 0x7d, 0x2f, 0x73, 0x63, 0x68, 0x65,
	0x6d, 0x61, 0x73, 0x3a, 0x06, 0x73, 0x63, 0x68, 0x65, 0x6d, 0x61, 0xda, 0x41, 0x17, 0x70, 0x61,
	0x72, 0x65, 0x6e, 0x74, 0x2c, 0x73, 0x63, 0x68, 0x65, 0x6d, 0x61, 0x2c, 0x73, 0x63, 0x68, 0x65,
	0x6d, 0x61, 0x5f, 0x69, 0x64, 0x12, 0xaf, 0x01, 0x0a, 0x0c, 0x44, 0x65, 0x6c, 0x65, 0x74, 0x65,
	0x53, 0x63, 0x68, 0x65, 0x6d, 0x61, 0x12, 0x39, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e,
	0x63, 0x6c, 0x6f, 0x75, 0x64, 0x2e, 0x64, 0x69, 0x73, 0x63, 0x6f, 0x76, 0x65, 0x72, 0x79, 0x65,
	0x6e, 0x67, 0x69, 0x6e, 0x65, 0x2e, 0x76, 0x31, 0x61, 0x6c, 0x70, 0x68, 0x61, 0x2e, 0x44, 0x65,
	0x6c, 0x65, 0x74, 0x65, 0x53, 0x63, 0x68, 0x65, 0x6d, 0x61, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73,
	0x74, 0x1a, 0x16, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x62, 0x75, 0x66, 0x2e, 0x45, 0x6d, 0x70, 0x74, 0x79, 0x22, 0x4c, 0x82, 0xd3, 0xe4, 0x93, 0x02,
	0x3f, 0x2a, 0x3d, 0x2f, 0x76, 0x31, 0x61, 0x6c, 0x70, 0x68, 0x61, 0x2f, 0x7b, 0x6e, 0x61, 0x6d,
	0x65, 0x3d, 0x70, 0x72, 0x6f, 0x6a, 0x65, 0x63, 0x74, 0x73, 0x2f, 0x2a, 0x2f, 0x6c, 0x6f, 0x63,
	0x61, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x2f, 0x2a, 0x2f, 0x64, 0x61, 0x74, 0x61, 0x53, 0x74, 0x6f,
	0x72, 0x65, 0x73, 0x2f, 0x2a, 0x2f, 0x73, 0x63, 0x68, 0x65, 0x6d, 0x61, 0x73, 0x2f, 0x2a, 0x7d,
	0xda, 0x41, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x1a, 0x52, 0xca, 0x41, 0x1e, 0x64, 0x69, 0x73, 0x63,
	0x6f, 0x76, 0x65, 0x72, 0x79, 0x65, 0x6e, 0x67, 0x69, 0x6e, 0x65, 0x2e, 0x67, 0x6f, 0x6f, 0x67,
	0x6c, 0x65, 0x61, 0x70, 0x69, 0x73, 0x2e, 0x63, 0x6f, 0x6d, 0xd2, 0x41, 0x2e, 0x68, 0x74, 0x74,
	0x70, 0x73, 0x3a, 0x2f, 0x2f, 0x77, 0x77, 0x77, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x61,
	0x70, 0x69, 0x73, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x61, 0x75, 0x74, 0x68, 0x2f, 0x63, 0x6c, 0x6f,
	0x75, 0x64, 0x2d, 0x70, 0x6c, 0x61, 0x74, 0x66, 0x6f, 0x72, 0x6d, 0x42, 0x9f, 0x02, 0x0a, 0x28,
	0x63, 0x6f, 0x6d, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x63, 0x6c, 0x6f, 0x75, 0x64,
	0x2e, 0x64, 0x69, 0x73, 0x63, 0x6f, 0x76, 0x65, 0x72, 0x79, 0x65, 0x6e, 0x67, 0x69, 0x6e, 0x65,
	0x2e, 0x76, 0x31, 0x61, 0x6c, 0x70, 0x68, 0x61, 0x42, 0x12, 0x53, 0x63, 0x68, 0x65, 0x6d, 0x61,
	0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x50, 0x72, 0x6f, 0x74, 0x6f, 0x50, 0x01, 0x5a, 0x53,
	0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x67, 0x6f, 0x6c, 0x61, 0x6e, 0x67, 0x2e, 0x6f, 0x72,
	0x67, 0x2f, 0x67, 0x65, 0x6e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2f, 0x67, 0x6f, 0x6f, 0x67, 0x6c,
	0x65, 0x61, 0x70, 0x69, 0x73, 0x2f, 0x63, 0x6c, 0x6f, 0x75, 0x64, 0x2f, 0x64, 0x69, 0x73, 0x63,
	0x6f, 0x76, 0x65, 0x72, 0x79, 0x65, 0x6e, 0x67, 0x69, 0x6e, 0x65, 0x2f, 0x76, 0x31, 0x61, 0x6c,
	0x70, 0x68, 0x61, 0x3b, 0x64, 0x69, 0x73, 0x63, 0x6f, 0x76, 0x65, 0x72, 0x79, 0x65, 0x6e, 0x67,
	0x69, 0x6e, 0x65, 0xa2, 0x02, 0x0f, 0x44, 0x49, 0x53, 0x43, 0x4f, 0x56, 0x45, 0x52, 0x59, 0x45,
	0x4e, 0x47, 0x49, 0x4e, 0x45, 0xaa, 0x02, 0x24, 0x47, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x43,
	0x6c, 0x6f, 0x75, 0x64, 0x2e, 0x44, 0x69, 0x73, 0x63, 0x6f, 0x76, 0x65, 0x72, 0x79, 0x45, 0x6e,
	0x67, 0x69, 0x6e, 0x65, 0x2e, 0x56, 0x31, 0x41, 0x6c, 0x70, 0x68, 0x61, 0xca, 0x02, 0x24, 0x47,
	0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x5c, 0x43, 0x6c, 0x6f, 0x75, 0x64, 0x5c, 0x44, 0x69, 0x73, 0x63,
	0x6f, 0x76, 0x65, 0x72, 0x79, 0x45, 0x6e, 0x67, 0x69, 0x6e, 0x65, 0x5c, 0x56, 0x31, 0x61, 0x6c,
	0x70, 0x68, 0x61, 0xea, 0x02, 0x27, 0x47, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x3a, 0x3a, 0x43, 0x6c,
	0x6f, 0x75, 0x64, 0x3a, 0x3a, 0x44, 0x69, 0x73, 0x63, 0x6f, 0x76, 0x65, 0x72, 0x79, 0x45, 0x6e,
	0x67, 0x69, 0x6e, 0x65, 0x3a, 0x3a, 0x56, 0x31, 0x61, 0x6c, 0x70, 0x68, 0x61, 0x62, 0x06, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_google_cloud_discoveryengine_v1alpha_schema_service_proto_rawDescOnce sync.Once
	file_google_cloud_discoveryengine_v1alpha_schema_service_proto_rawDescData = file_google_cloud_discoveryengine_v1alpha_schema_service_proto_rawDesc
)

func file_google_cloud_discoveryengine_v1alpha_schema_service_proto_rawDescGZIP() []byte {
	file_google_cloud_discoveryengine_v1alpha_schema_service_proto_rawDescOnce.Do(func() {
		file_google_cloud_discoveryengine_v1alpha_schema_service_proto_rawDescData = protoimpl.X.CompressGZIP(file_google_cloud_discoveryengine_v1alpha_schema_service_proto_rawDescData)
	})
	return file_google_cloud_discoveryengine_v1alpha_schema_service_proto_rawDescData
}

var file_google_cloud_discoveryengine_v1alpha_schema_service_proto_msgTypes = make([]protoimpl.MessageInfo, 5)
var file_google_cloud_discoveryengine_v1alpha_schema_service_proto_goTypes = []interface{}{
	(*GetSchemaRequest)(nil),    // 0: google.cloud.discoveryengine.v1alpha.GetSchemaRequest
	(*ListSchemasRequest)(nil),  // 1: google.cloud.discoveryengine.v1alpha.ListSchemasRequest
	(*ListSchemasResponse)(nil), // 2: google.cloud.discoveryengine.v1alpha.ListSchemasResponse
	(*CreateSchemaRequest)(nil), // 3: google.cloud.discoveryengine.v1alpha.CreateSchemaRequest
	(*DeleteSchemaRequest)(nil), // 4: google.cloud.discoveryengine.v1alpha.DeleteSchemaRequest
	(*Schema)(nil),              // 5: google.cloud.discoveryengine.v1alpha.Schema
	(*emptypb.Empty)(nil),       // 6: google.protobuf.Empty
}
var file_google_cloud_discoveryengine_v1alpha_schema_service_proto_depIdxs = []int32{
	5, // 0: google.cloud.discoveryengine.v1alpha.ListSchemasResponse.schemas:type_name -> google.cloud.discoveryengine.v1alpha.Schema
	5, // 1: google.cloud.discoveryengine.v1alpha.CreateSchemaRequest.schema:type_name -> google.cloud.discoveryengine.v1alpha.Schema
	0, // 2: google.cloud.discoveryengine.v1alpha.SchemaService.GetSchema:input_type -> google.cloud.discoveryengine.v1alpha.GetSchemaRequest
	1, // 3: google.cloud.discoveryengine.v1alpha.SchemaService.ListSchemas:input_type -> google.cloud.discoveryengine.v1alpha.ListSchemasRequest
	3, // 4: google.cloud.discoveryengine.v1alpha.SchemaService.CreateSchema:input_type -> google.cloud.discoveryengine.v1alpha.CreateSchemaRequest
	4, // 5: google.cloud.discoveryengine.v1alpha.SchemaService.DeleteSchema:input_type -> google.cloud.discoveryengine.v1alpha.DeleteSchemaRequest
	5, // 6: google.cloud.discoveryengine.v1alpha.SchemaService.GetSchema:output_type -> google.cloud.discoveryengine.v1alpha.Schema
	2, // 7: google.cloud.discoveryengine.v1alpha.SchemaService.ListSchemas:output_type -> google.cloud.discoveryengine.v1alpha.ListSchemasResponse
	5, // 8: google.cloud.discoveryengine.v1alpha.SchemaService.CreateSchema:output_type -> google.cloud.discoveryengine.v1alpha.Schema
	6, // 9: google.cloud.discoveryengine.v1alpha.SchemaService.DeleteSchema:output_type -> google.protobuf.Empty
	6, // [6:10] is the sub-list for method output_type
	2, // [2:6] is the sub-list for method input_type
	2, // [2:2] is the sub-list for extension type_name
	2, // [2:2] is the sub-list for extension extendee
	0, // [0:2] is the sub-list for field type_name
}

func init() { file_google_cloud_discoveryengine_v1alpha_schema_service_proto_init() }
func file_google_cloud_discoveryengine_v1alpha_schema_service_proto_init() {
	if File_google_cloud_discoveryengine_v1alpha_schema_service_proto != nil {
		return
	}
	file_google_cloud_discoveryengine_v1alpha_schema_proto_init()
	if !protoimpl.UnsafeEnabled {
		file_google_cloud_discoveryengine_v1alpha_schema_service_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetSchemaRequest); i {
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
		file_google_cloud_discoveryengine_v1alpha_schema_service_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ListSchemasRequest); i {
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
		file_google_cloud_discoveryengine_v1alpha_schema_service_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ListSchemasResponse); i {
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
		file_google_cloud_discoveryengine_v1alpha_schema_service_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*CreateSchemaRequest); i {
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
		file_google_cloud_discoveryengine_v1alpha_schema_service_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*DeleteSchemaRequest); i {
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
			RawDescriptor: file_google_cloud_discoveryengine_v1alpha_schema_service_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   5,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_google_cloud_discoveryengine_v1alpha_schema_service_proto_goTypes,
		DependencyIndexes: file_google_cloud_discoveryengine_v1alpha_schema_service_proto_depIdxs,
		MessageInfos:      file_google_cloud_discoveryengine_v1alpha_schema_service_proto_msgTypes,
	}.Build()
	File_google_cloud_discoveryengine_v1alpha_schema_service_proto = out.File
	file_google_cloud_discoveryengine_v1alpha_schema_service_proto_rawDesc = nil
	file_google_cloud_discoveryengine_v1alpha_schema_service_proto_goTypes = nil
	file_google_cloud_discoveryengine_v1alpha_schema_service_proto_depIdxs = nil
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConnInterface

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion6

// SchemaServiceClient is the client API for SchemaService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type SchemaServiceClient interface {
	// Gets a [Schema][google.cloud.discoveryengine.v1alpha.Schema].
	GetSchema(ctx context.Context, in *GetSchemaRequest, opts ...grpc.CallOption) (*Schema, error)
	// Gets a list of [Schema][google.cloud.discoveryengine.v1alpha.Schema]s.
	ListSchemas(ctx context.Context, in *ListSchemasRequest, opts ...grpc.CallOption) (*ListSchemasResponse, error)
	// Creates a [Schema][google.cloud.discoveryengine.v1alpha.Schema].
	CreateSchema(ctx context.Context, in *CreateSchemaRequest, opts ...grpc.CallOption) (*Schema, error)
	// Deletes a [Schema][google.cloud.discoveryengine.v1alpha.Schema].
	DeleteSchema(ctx context.Context, in *DeleteSchemaRequest, opts ...grpc.CallOption) (*emptypb.Empty, error)
}

type schemaServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewSchemaServiceClient(cc grpc.ClientConnInterface) SchemaServiceClient {
	return &schemaServiceClient{cc}
}

func (c *schemaServiceClient) GetSchema(ctx context.Context, in *GetSchemaRequest, opts ...grpc.CallOption) (*Schema, error) {
	out := new(Schema)
	err := c.cc.Invoke(ctx, "/google.cloud.discoveryengine.v1alpha.SchemaService/GetSchema", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *schemaServiceClient) ListSchemas(ctx context.Context, in *ListSchemasRequest, opts ...grpc.CallOption) (*ListSchemasResponse, error) {
	out := new(ListSchemasResponse)
	err := c.cc.Invoke(ctx, "/google.cloud.discoveryengine.v1alpha.SchemaService/ListSchemas", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *schemaServiceClient) CreateSchema(ctx context.Context, in *CreateSchemaRequest, opts ...grpc.CallOption) (*Schema, error) {
	out := new(Schema)
	err := c.cc.Invoke(ctx, "/google.cloud.discoveryengine.v1alpha.SchemaService/CreateSchema", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *schemaServiceClient) DeleteSchema(ctx context.Context, in *DeleteSchemaRequest, opts ...grpc.CallOption) (*emptypb.Empty, error) {
	out := new(emptypb.Empty)
	err := c.cc.Invoke(ctx, "/google.cloud.discoveryengine.v1alpha.SchemaService/DeleteSchema", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// SchemaServiceServer is the server API for SchemaService service.
type SchemaServiceServer interface {
	// Gets a [Schema][google.cloud.discoveryengine.v1alpha.Schema].
	GetSchema(context.Context, *GetSchemaRequest) (*Schema, error)
	// Gets a list of [Schema][google.cloud.discoveryengine.v1alpha.Schema]s.
	ListSchemas(context.Context, *ListSchemasRequest) (*ListSchemasResponse, error)
	// Creates a [Schema][google.cloud.discoveryengine.v1alpha.Schema].
	CreateSchema(context.Context, *CreateSchemaRequest) (*Schema, error)
	// Deletes a [Schema][google.cloud.discoveryengine.v1alpha.Schema].
	DeleteSchema(context.Context, *DeleteSchemaRequest) (*emptypb.Empty, error)
}

// UnimplementedSchemaServiceServer can be embedded to have forward compatible implementations.
type UnimplementedSchemaServiceServer struct {
}

func (*UnimplementedSchemaServiceServer) GetSchema(context.Context, *GetSchemaRequest) (*Schema, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetSchema not implemented")
}
func (*UnimplementedSchemaServiceServer) ListSchemas(context.Context, *ListSchemasRequest) (*ListSchemasResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method ListSchemas not implemented")
}
func (*UnimplementedSchemaServiceServer) CreateSchema(context.Context, *CreateSchemaRequest) (*Schema, error) {
	return nil, status.Errorf(codes.Unimplemented, "method CreateSchema not implemented")
}
func (*UnimplementedSchemaServiceServer) DeleteSchema(context.Context, *DeleteSchemaRequest) (*emptypb.Empty, error) {
	return nil, status.Errorf(codes.Unimplemented, "method DeleteSchema not implemented")
}

func RegisterSchemaServiceServer(s *grpc.Server, srv SchemaServiceServer) {
	s.RegisterService(&_SchemaService_serviceDesc, srv)
}

func _SchemaService_GetSchema_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetSchemaRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(SchemaServiceServer).GetSchema(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/google.cloud.discoveryengine.v1alpha.SchemaService/GetSchema",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(SchemaServiceServer).GetSchema(ctx, req.(*GetSchemaRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _SchemaService_ListSchemas_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ListSchemasRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(SchemaServiceServer).ListSchemas(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/google.cloud.discoveryengine.v1alpha.SchemaService/ListSchemas",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(SchemaServiceServer).ListSchemas(ctx, req.(*ListSchemasRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _SchemaService_CreateSchema_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(CreateSchemaRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(SchemaServiceServer).CreateSchema(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/google.cloud.discoveryengine.v1alpha.SchemaService/CreateSchema",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(SchemaServiceServer).CreateSchema(ctx, req.(*CreateSchemaRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _SchemaService_DeleteSchema_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(DeleteSchemaRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(SchemaServiceServer).DeleteSchema(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/google.cloud.discoveryengine.v1alpha.SchemaService/DeleteSchema",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(SchemaServiceServer).DeleteSchema(ctx, req.(*DeleteSchemaRequest))
	}
	return interceptor(ctx, in, info, handler)
}

var _SchemaService_serviceDesc = grpc.ServiceDesc{
	ServiceName: "google.cloud.discoveryengine.v1alpha.SchemaService",
	HandlerType: (*SchemaServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "GetSchema",
			Handler:    _SchemaService_GetSchema_Handler,
		},
		{
			MethodName: "ListSchemas",
			Handler:    _SchemaService_ListSchemas_Handler,
		},
		{
			MethodName: "CreateSchema",
			Handler:    _SchemaService_CreateSchema_Handler,
		},
		{
			MethodName: "DeleteSchema",
			Handler:    _SchemaService_DeleteSchema_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "google/cloud/discoveryengine/v1alpha/schema_service.proto",
}
