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
// 	protoc-gen-go v1.25.0
// 	protoc        v3.13.0
// source: google/cloud/pubsublite/v1/topic_stats.proto

package pubsublite

import (
	context "context"
	reflect "reflect"
	sync "sync"

	proto "github.com/golang/protobuf/proto"
	_ "google.golang.org/genproto/googleapis/api/annotations"
	grpc "google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
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

// This is a compile-time assertion that a sufficiently up-to-date version
// of the legacy proto package is being used.
const _ = proto.ProtoPackageIsVersion4

// Compute statistics about a range of messages in a given topic and partition.
type ComputeMessageStatsRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// Required. The topic for which we should compute message stats.
	Topic string `protobuf:"bytes,1,opt,name=topic,proto3" json:"topic,omitempty"`
	// Required. The partition for which we should compute message stats.
	Partition int64 `protobuf:"varint,2,opt,name=partition,proto3" json:"partition,omitempty"`
	// The inclusive start of the range.
	StartCursor *Cursor `protobuf:"bytes,3,opt,name=start_cursor,json=startCursor,proto3" json:"start_cursor,omitempty"`
	// The exclusive end of the range. The range is empty if end_cursor <=
	// start_cursor. Specifying a start_cursor before the first message and an
	// end_cursor after the last message will retrieve all messages.
	EndCursor *Cursor `protobuf:"bytes,4,opt,name=end_cursor,json=endCursor,proto3" json:"end_cursor,omitempty"`
}

func (x *ComputeMessageStatsRequest) Reset() {
	*x = ComputeMessageStatsRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_google_cloud_pubsublite_v1_topic_stats_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ComputeMessageStatsRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ComputeMessageStatsRequest) ProtoMessage() {}

func (x *ComputeMessageStatsRequest) ProtoReflect() protoreflect.Message {
	mi := &file_google_cloud_pubsublite_v1_topic_stats_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ComputeMessageStatsRequest.ProtoReflect.Descriptor instead.
func (*ComputeMessageStatsRequest) Descriptor() ([]byte, []int) {
	return file_google_cloud_pubsublite_v1_topic_stats_proto_rawDescGZIP(), []int{0}
}

func (x *ComputeMessageStatsRequest) GetTopic() string {
	if x != nil {
		return x.Topic
	}
	return ""
}

func (x *ComputeMessageStatsRequest) GetPartition() int64 {
	if x != nil {
		return x.Partition
	}
	return 0
}

func (x *ComputeMessageStatsRequest) GetStartCursor() *Cursor {
	if x != nil {
		return x.StartCursor
	}
	return nil
}

func (x *ComputeMessageStatsRequest) GetEndCursor() *Cursor {
	if x != nil {
		return x.EndCursor
	}
	return nil
}

// Response containing stats for messages in the requested topic and partition.
type ComputeMessageStatsResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// The count of messages.
	MessageCount int64 `protobuf:"varint,1,opt,name=message_count,json=messageCount,proto3" json:"message_count,omitempty"`
	// The number of quota bytes accounted to these messages.
	MessageBytes int64 `protobuf:"varint,2,opt,name=message_bytes,json=messageBytes,proto3" json:"message_bytes,omitempty"`
	// The minimum publish timestamp across these messages. Note that publish
	// timestamps within a partition are not guaranteed to be non-decreasing. The
	// timestamp will be unset if there are no messages.
	MinimumPublishTime *timestamppb.Timestamp `protobuf:"bytes,3,opt,name=minimum_publish_time,json=minimumPublishTime,proto3" json:"minimum_publish_time,omitempty"`
	// The minimum event timestamp across these messages. For the purposes of this
	// computation, if a message does not have an event time, we use the publish
	// time. The timestamp will be unset if there are no messages.
	MinimumEventTime *timestamppb.Timestamp `protobuf:"bytes,4,opt,name=minimum_event_time,json=minimumEventTime,proto3" json:"minimum_event_time,omitempty"`
}

func (x *ComputeMessageStatsResponse) Reset() {
	*x = ComputeMessageStatsResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_google_cloud_pubsublite_v1_topic_stats_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ComputeMessageStatsResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ComputeMessageStatsResponse) ProtoMessage() {}

func (x *ComputeMessageStatsResponse) ProtoReflect() protoreflect.Message {
	mi := &file_google_cloud_pubsublite_v1_topic_stats_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ComputeMessageStatsResponse.ProtoReflect.Descriptor instead.
func (*ComputeMessageStatsResponse) Descriptor() ([]byte, []int) {
	return file_google_cloud_pubsublite_v1_topic_stats_proto_rawDescGZIP(), []int{1}
}

func (x *ComputeMessageStatsResponse) GetMessageCount() int64 {
	if x != nil {
		return x.MessageCount
	}
	return 0
}

func (x *ComputeMessageStatsResponse) GetMessageBytes() int64 {
	if x != nil {
		return x.MessageBytes
	}
	return 0
}

func (x *ComputeMessageStatsResponse) GetMinimumPublishTime() *timestamppb.Timestamp {
	if x != nil {
		return x.MinimumPublishTime
	}
	return nil
}

func (x *ComputeMessageStatsResponse) GetMinimumEventTime() *timestamppb.Timestamp {
	if x != nil {
		return x.MinimumEventTime
	}
	return nil
}

// Compute the current head cursor for a partition.
type ComputeHeadCursorRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// Required. The topic for which we should compute the head cursor.
	Topic string `protobuf:"bytes,1,opt,name=topic,proto3" json:"topic,omitempty"`
	// Required. The partition for which we should compute the head cursor.
	Partition int64 `protobuf:"varint,2,opt,name=partition,proto3" json:"partition,omitempty"`
}

func (x *ComputeHeadCursorRequest) Reset() {
	*x = ComputeHeadCursorRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_google_cloud_pubsublite_v1_topic_stats_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ComputeHeadCursorRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ComputeHeadCursorRequest) ProtoMessage() {}

func (x *ComputeHeadCursorRequest) ProtoReflect() protoreflect.Message {
	mi := &file_google_cloud_pubsublite_v1_topic_stats_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ComputeHeadCursorRequest.ProtoReflect.Descriptor instead.
func (*ComputeHeadCursorRequest) Descriptor() ([]byte, []int) {
	return file_google_cloud_pubsublite_v1_topic_stats_proto_rawDescGZIP(), []int{2}
}

func (x *ComputeHeadCursorRequest) GetTopic() string {
	if x != nil {
		return x.Topic
	}
	return ""
}

func (x *ComputeHeadCursorRequest) GetPartition() int64 {
	if x != nil {
		return x.Partition
	}
	return 0
}

// Response containing the head cursor for the requested topic and partition.
type ComputeHeadCursorResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// The head cursor.
	HeadCursor *Cursor `protobuf:"bytes,1,opt,name=head_cursor,json=headCursor,proto3" json:"head_cursor,omitempty"`
}

func (x *ComputeHeadCursorResponse) Reset() {
	*x = ComputeHeadCursorResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_google_cloud_pubsublite_v1_topic_stats_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ComputeHeadCursorResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ComputeHeadCursorResponse) ProtoMessage() {}

func (x *ComputeHeadCursorResponse) ProtoReflect() protoreflect.Message {
	mi := &file_google_cloud_pubsublite_v1_topic_stats_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ComputeHeadCursorResponse.ProtoReflect.Descriptor instead.
func (*ComputeHeadCursorResponse) Descriptor() ([]byte, []int) {
	return file_google_cloud_pubsublite_v1_topic_stats_proto_rawDescGZIP(), []int{3}
}

func (x *ComputeHeadCursorResponse) GetHeadCursor() *Cursor {
	if x != nil {
		return x.HeadCursor
	}
	return nil
}

var File_google_cloud_pubsublite_v1_topic_stats_proto protoreflect.FileDescriptor

var file_google_cloud_pubsublite_v1_topic_stats_proto_rawDesc = []byte{
	0x0a, 0x2c, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x63, 0x6c, 0x6f, 0x75, 0x64, 0x2f, 0x70,
	0x75, 0x62, 0x73, 0x75, 0x62, 0x6c, 0x69, 0x74, 0x65, 0x2f, 0x76, 0x31, 0x2f, 0x74, 0x6f, 0x70,
	0x69, 0x63, 0x5f, 0x73, 0x74, 0x61, 0x74, 0x73, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x1a,
	0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x63, 0x6c, 0x6f, 0x75, 0x64, 0x2e, 0x70, 0x75, 0x62,
	0x73, 0x75, 0x62, 0x6c, 0x69, 0x74, 0x65, 0x2e, 0x76, 0x31, 0x1a, 0x1c, 0x67, 0x6f, 0x6f, 0x67,
	0x6c, 0x65, 0x2f, 0x61, 0x70, 0x69, 0x2f, 0x61, 0x6e, 0x6e, 0x6f, 0x74, 0x61, 0x74, 0x69, 0x6f,
	0x6e, 0x73, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x1f, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65,
	0x2f, 0x61, 0x70, 0x69, 0x2f, 0x66, 0x69, 0x65, 0x6c, 0x64, 0x5f, 0x62, 0x65, 0x68, 0x61, 0x76,
	0x69, 0x6f, 0x72, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x19, 0x67, 0x6f, 0x6f, 0x67, 0x6c,
	0x65, 0x2f, 0x61, 0x70, 0x69, 0x2f, 0x72, 0x65, 0x73, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x2e, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x27, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x63, 0x6c, 0x6f,
	0x75, 0x64, 0x2f, 0x70, 0x75, 0x62, 0x73, 0x75, 0x62, 0x6c, 0x69, 0x74, 0x65, 0x2f, 0x76, 0x31,
	0x2f, 0x63, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x1f, 0x67,
	0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2f, 0x74,
	0x69, 0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d, 0x70, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x17,
	0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x61, 0x70, 0x69, 0x2f, 0x63, 0x6c, 0x69, 0x65, 0x6e,
	0x74, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0x88, 0x02, 0x0a, 0x1a, 0x43, 0x6f, 0x6d, 0x70,
	0x75, 0x74, 0x65, 0x4d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x53, 0x74, 0x61, 0x74, 0x73, 0x52,
	0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x3d, 0x0a, 0x05, 0x74, 0x6f, 0x70, 0x69, 0x63, 0x18,
	0x01, 0x20, 0x01, 0x28, 0x09, 0x42, 0x27, 0xe0, 0x41, 0x02, 0xfa, 0x41, 0x21, 0x0a, 0x1f, 0x70,
	0x75, 0x62, 0x73, 0x75, 0x62, 0x6c, 0x69, 0x74, 0x65, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65,
	0x61, 0x70, 0x69, 0x73, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x54, 0x6f, 0x70, 0x69, 0x63, 0x52, 0x05,
	0x74, 0x6f, 0x70, 0x69, 0x63, 0x12, 0x21, 0x0a, 0x09, 0x70, 0x61, 0x72, 0x74, 0x69, 0x74, 0x69,
	0x6f, 0x6e, 0x18, 0x02, 0x20, 0x01, 0x28, 0x03, 0x42, 0x03, 0xe0, 0x41, 0x02, 0x52, 0x09, 0x70,
	0x61, 0x72, 0x74, 0x69, 0x74, 0x69, 0x6f, 0x6e, 0x12, 0x45, 0x0a, 0x0c, 0x73, 0x74, 0x61, 0x72,
	0x74, 0x5f, 0x63, 0x75, 0x72, 0x73, 0x6f, 0x72, 0x18, 0x03, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x22,
	0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x63, 0x6c, 0x6f, 0x75, 0x64, 0x2e, 0x70, 0x75,
	0x62, 0x73, 0x75, 0x62, 0x6c, 0x69, 0x74, 0x65, 0x2e, 0x76, 0x31, 0x2e, 0x43, 0x75, 0x72, 0x73,
	0x6f, 0x72, 0x52, 0x0b, 0x73, 0x74, 0x61, 0x72, 0x74, 0x43, 0x75, 0x72, 0x73, 0x6f, 0x72, 0x12,
	0x41, 0x0a, 0x0a, 0x65, 0x6e, 0x64, 0x5f, 0x63, 0x75, 0x72, 0x73, 0x6f, 0x72, 0x18, 0x04, 0x20,
	0x01, 0x28, 0x0b, 0x32, 0x22, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x63, 0x6c, 0x6f,
	0x75, 0x64, 0x2e, 0x70, 0x75, 0x62, 0x73, 0x75, 0x62, 0x6c, 0x69, 0x74, 0x65, 0x2e, 0x76, 0x31,
	0x2e, 0x43, 0x75, 0x72, 0x73, 0x6f, 0x72, 0x52, 0x09, 0x65, 0x6e, 0x64, 0x43, 0x75, 0x72, 0x73,
	0x6f, 0x72, 0x22, 0xff, 0x01, 0x0a, 0x1b, 0x43, 0x6f, 0x6d, 0x70, 0x75, 0x74, 0x65, 0x4d, 0x65,
	0x73, 0x73, 0x61, 0x67, 0x65, 0x53, 0x74, 0x61, 0x74, 0x73, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e,
	0x73, 0x65, 0x12, 0x23, 0x0a, 0x0d, 0x6d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x5f, 0x63, 0x6f,
	0x75, 0x6e, 0x74, 0x18, 0x01, 0x20, 0x01, 0x28, 0x03, 0x52, 0x0c, 0x6d, 0x65, 0x73, 0x73, 0x61,
	0x67, 0x65, 0x43, 0x6f, 0x75, 0x6e, 0x74, 0x12, 0x23, 0x0a, 0x0d, 0x6d, 0x65, 0x73, 0x73, 0x61,
	0x67, 0x65, 0x5f, 0x62, 0x79, 0x74, 0x65, 0x73, 0x18, 0x02, 0x20, 0x01, 0x28, 0x03, 0x52, 0x0c,
	0x6d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x42, 0x79, 0x74, 0x65, 0x73, 0x12, 0x4c, 0x0a, 0x14,
	0x6d, 0x69, 0x6e, 0x69, 0x6d, 0x75, 0x6d, 0x5f, 0x70, 0x75, 0x62, 0x6c, 0x69, 0x73, 0x68, 0x5f,
	0x74, 0x69, 0x6d, 0x65, 0x18, 0x03, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1a, 0x2e, 0x67, 0x6f, 0x6f,
	0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x54, 0x69, 0x6d,
	0x65, 0x73, 0x74, 0x61, 0x6d, 0x70, 0x52, 0x12, 0x6d, 0x69, 0x6e, 0x69, 0x6d, 0x75, 0x6d, 0x50,
	0x75, 0x62, 0x6c, 0x69, 0x73, 0x68, 0x54, 0x69, 0x6d, 0x65, 0x12, 0x48, 0x0a, 0x12, 0x6d, 0x69,
	0x6e, 0x69, 0x6d, 0x75, 0x6d, 0x5f, 0x65, 0x76, 0x65, 0x6e, 0x74, 0x5f, 0x74, 0x69, 0x6d, 0x65,
	0x18, 0x04, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1a, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x54, 0x69, 0x6d, 0x65, 0x73, 0x74, 0x61,
	0x6d, 0x70, 0x52, 0x10, 0x6d, 0x69, 0x6e, 0x69, 0x6d, 0x75, 0x6d, 0x45, 0x76, 0x65, 0x6e, 0x74,
	0x54, 0x69, 0x6d, 0x65, 0x22, 0x7c, 0x0a, 0x18, 0x43, 0x6f, 0x6d, 0x70, 0x75, 0x74, 0x65, 0x48,
	0x65, 0x61, 0x64, 0x43, 0x75, 0x72, 0x73, 0x6f, 0x72, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74,
	0x12, 0x3d, 0x0a, 0x05, 0x74, 0x6f, 0x70, 0x69, 0x63, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x42,
	0x27, 0xe0, 0x41, 0x02, 0xfa, 0x41, 0x21, 0x0a, 0x1f, 0x70, 0x75, 0x62, 0x73, 0x75, 0x62, 0x6c,
	0x69, 0x74, 0x65, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x61, 0x70, 0x69, 0x73, 0x2e, 0x63,
	0x6f, 0x6d, 0x2f, 0x54, 0x6f, 0x70, 0x69, 0x63, 0x52, 0x05, 0x74, 0x6f, 0x70, 0x69, 0x63, 0x12,
	0x21, 0x0a, 0x09, 0x70, 0x61, 0x72, 0x74, 0x69, 0x74, 0x69, 0x6f, 0x6e, 0x18, 0x02, 0x20, 0x01,
	0x28, 0x03, 0x42, 0x03, 0xe0, 0x41, 0x02, 0x52, 0x09, 0x70, 0x61, 0x72, 0x74, 0x69, 0x74, 0x69,
	0x6f, 0x6e, 0x22, 0x60, 0x0a, 0x19, 0x43, 0x6f, 0x6d, 0x70, 0x75, 0x74, 0x65, 0x48, 0x65, 0x61,
	0x64, 0x43, 0x75, 0x72, 0x73, 0x6f, 0x72, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12,
	0x43, 0x0a, 0x0b, 0x68, 0x65, 0x61, 0x64, 0x5f, 0x63, 0x75, 0x72, 0x73, 0x6f, 0x72, 0x18, 0x01,
	0x20, 0x01, 0x28, 0x0b, 0x32, 0x22, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x63, 0x6c,
	0x6f, 0x75, 0x64, 0x2e, 0x70, 0x75, 0x62, 0x73, 0x75, 0x62, 0x6c, 0x69, 0x74, 0x65, 0x2e, 0x76,
	0x31, 0x2e, 0x43, 0x75, 0x72, 0x73, 0x6f, 0x72, 0x52, 0x0a, 0x68, 0x65, 0x61, 0x64, 0x43, 0x75,
	0x72, 0x73, 0x6f, 0x72, 0x32, 0x9a, 0x04, 0x0a, 0x11, 0x54, 0x6f, 0x70, 0x69, 0x63, 0x53, 0x74,
	0x61, 0x74, 0x73, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x12, 0xdd, 0x01, 0x0a, 0x13, 0x43,
	0x6f, 0x6d, 0x70, 0x75, 0x74, 0x65, 0x4d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x53, 0x74, 0x61,
	0x74, 0x73, 0x12, 0x36, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x63, 0x6c, 0x6f, 0x75,
	0x64, 0x2e, 0x70, 0x75, 0x62, 0x73, 0x75, 0x62, 0x6c, 0x69, 0x74, 0x65, 0x2e, 0x76, 0x31, 0x2e,
	0x43, 0x6f, 0x6d, 0x70, 0x75, 0x74, 0x65, 0x4d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x53, 0x74,
	0x61, 0x74, 0x73, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x37, 0x2e, 0x67, 0x6f, 0x6f,
	0x67, 0x6c, 0x65, 0x2e, 0x63, 0x6c, 0x6f, 0x75, 0x64, 0x2e, 0x70, 0x75, 0x62, 0x73, 0x75, 0x62,
	0x6c, 0x69, 0x74, 0x65, 0x2e, 0x76, 0x31, 0x2e, 0x43, 0x6f, 0x6d, 0x70, 0x75, 0x74, 0x65, 0x4d,
	0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x53, 0x74, 0x61, 0x74, 0x73, 0x52, 0x65, 0x73, 0x70, 0x6f,
	0x6e, 0x73, 0x65, 0x22, 0x55, 0x82, 0xd3, 0xe4, 0x93, 0x02, 0x4f, 0x22, 0x4a, 0x2f, 0x76, 0x31,
	0x2f, 0x74, 0x6f, 0x70, 0x69, 0x63, 0x53, 0x74, 0x61, 0x74, 0x73, 0x2f, 0x7b, 0x74, 0x6f, 0x70,
	0x69, 0x63, 0x3d, 0x70, 0x72, 0x6f, 0x6a, 0x65, 0x63, 0x74, 0x73, 0x2f, 0x2a, 0x2f, 0x6c, 0x6f,
	0x63, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x2f, 0x2a, 0x2f, 0x74, 0x6f, 0x70, 0x69, 0x63, 0x73,
	0x2f, 0x2a, 0x7d, 0x3a, 0x63, 0x6f, 0x6d, 0x70, 0x75, 0x74, 0x65, 0x4d, 0x65, 0x73, 0x73, 0x61,
	0x67, 0x65, 0x53, 0x74, 0x61, 0x74, 0x73, 0x3a, 0x01, 0x2a, 0x12, 0xd5, 0x01, 0x0a, 0x11, 0x43,
	0x6f, 0x6d, 0x70, 0x75, 0x74, 0x65, 0x48, 0x65, 0x61, 0x64, 0x43, 0x75, 0x72, 0x73, 0x6f, 0x72,
	0x12, 0x34, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x63, 0x6c, 0x6f, 0x75, 0x64, 0x2e,
	0x70, 0x75, 0x62, 0x73, 0x75, 0x62, 0x6c, 0x69, 0x74, 0x65, 0x2e, 0x76, 0x31, 0x2e, 0x43, 0x6f,
	0x6d, 0x70, 0x75, 0x74, 0x65, 0x48, 0x65, 0x61, 0x64, 0x43, 0x75, 0x72, 0x73, 0x6f, 0x72, 0x52,
	0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x35, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e,
	0x63, 0x6c, 0x6f, 0x75, 0x64, 0x2e, 0x70, 0x75, 0x62, 0x73, 0x75, 0x62, 0x6c, 0x69, 0x74, 0x65,
	0x2e, 0x76, 0x31, 0x2e, 0x43, 0x6f, 0x6d, 0x70, 0x75, 0x74, 0x65, 0x48, 0x65, 0x61, 0x64, 0x43,
	0x75, 0x72, 0x73, 0x6f, 0x72, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x53, 0x82,
	0xd3, 0xe4, 0x93, 0x02, 0x4d, 0x22, 0x48, 0x2f, 0x76, 0x31, 0x2f, 0x74, 0x6f, 0x70, 0x69, 0x63,
	0x53, 0x74, 0x61, 0x74, 0x73, 0x2f, 0x7b, 0x74, 0x6f, 0x70, 0x69, 0x63, 0x3d, 0x70, 0x72, 0x6f,
	0x6a, 0x65, 0x63, 0x74, 0x73, 0x2f, 0x2a, 0x2f, 0x6c, 0x6f, 0x63, 0x61, 0x74, 0x69, 0x6f, 0x6e,
	0x73, 0x2f, 0x2a, 0x2f, 0x74, 0x6f, 0x70, 0x69, 0x63, 0x73, 0x2f, 0x2a, 0x7d, 0x3a, 0x63, 0x6f,
	0x6d, 0x70, 0x75, 0x74, 0x65, 0x48, 0x65, 0x61, 0x64, 0x43, 0x75, 0x72, 0x73, 0x6f, 0x72, 0x3a,
	0x01, 0x2a, 0x1a, 0x4d, 0xca, 0x41, 0x19, 0x70, 0x75, 0x62, 0x73, 0x75, 0x62, 0x6c, 0x69, 0x74,
	0x65, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x61, 0x70, 0x69, 0x73, 0x2e, 0x63, 0x6f, 0x6d,
	0xd2, 0x41, 0x2e, 0x68, 0x74, 0x74, 0x70, 0x73, 0x3a, 0x2f, 0x2f, 0x77, 0x77, 0x77, 0x2e, 0x67,
	0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x61, 0x70, 0x69, 0x73, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x61, 0x75,
	0x74, 0x68, 0x2f, 0x63, 0x6c, 0x6f, 0x75, 0x64, 0x2d, 0x70, 0x6c, 0x61, 0x74, 0x66, 0x6f, 0x72,
	0x6d, 0x42, 0xd6, 0x01, 0x0a, 0x21, 0x63, 0x6f, 0x6d, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65,
	0x2e, 0x63, 0x6c, 0x6f, 0x75, 0x64, 0x2e, 0x70, 0x75, 0x62, 0x73, 0x75, 0x62, 0x6c, 0x69, 0x74,
	0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x42, 0x0f, 0x54, 0x6f, 0x70, 0x69, 0x63, 0x53, 0x74,
	0x61, 0x74, 0x73, 0x50, 0x72, 0x6f, 0x74, 0x6f, 0x50, 0x01, 0x5a, 0x44, 0x67, 0x6f, 0x6f, 0x67,
	0x6c, 0x65, 0x2e, 0x67, 0x6f, 0x6c, 0x61, 0x6e, 0x67, 0x2e, 0x6f, 0x72, 0x67, 0x2f, 0x67, 0x65,
	0x6e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2f, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x61, 0x70, 0x69,
	0x73, 0x2f, 0x63, 0x6c, 0x6f, 0x75, 0x64, 0x2f, 0x70, 0x75, 0x62, 0x73, 0x75, 0x62, 0x6c, 0x69,
	0x74, 0x65, 0x2f, 0x76, 0x31, 0x3b, 0x70, 0x75, 0x62, 0x73, 0x75, 0x62, 0x6c, 0x69, 0x74, 0x65,
	0xaa, 0x02, 0x1a, 0x47, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x43, 0x6c, 0x6f, 0x75, 0x64, 0x2e,
	0x50, 0x75, 0x62, 0x53, 0x75, 0x62, 0x4c, 0x69, 0x74, 0x65, 0x2e, 0x56, 0x31, 0xca, 0x02, 0x1a,
	0x47, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x5c, 0x43, 0x6c, 0x6f, 0x75, 0x64, 0x5c, 0x50, 0x75, 0x62,
	0x53, 0x75, 0x62, 0x4c, 0x69, 0x74, 0x65, 0x5c, 0x56, 0x31, 0xea, 0x02, 0x1d, 0x47, 0x6f, 0x6f,
	0x67, 0x6c, 0x65, 0x3a, 0x3a, 0x43, 0x6c, 0x6f, 0x75, 0x64, 0x3a, 0x3a, 0x50, 0x75, 0x62, 0x53,
	0x75, 0x62, 0x4c, 0x69, 0x74, 0x65, 0x3a, 0x3a, 0x56, 0x31, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x33,
}

var (
	file_google_cloud_pubsublite_v1_topic_stats_proto_rawDescOnce sync.Once
	file_google_cloud_pubsublite_v1_topic_stats_proto_rawDescData = file_google_cloud_pubsublite_v1_topic_stats_proto_rawDesc
)

func file_google_cloud_pubsublite_v1_topic_stats_proto_rawDescGZIP() []byte {
	file_google_cloud_pubsublite_v1_topic_stats_proto_rawDescOnce.Do(func() {
		file_google_cloud_pubsublite_v1_topic_stats_proto_rawDescData = protoimpl.X.CompressGZIP(file_google_cloud_pubsublite_v1_topic_stats_proto_rawDescData)
	})
	return file_google_cloud_pubsublite_v1_topic_stats_proto_rawDescData
}

var file_google_cloud_pubsublite_v1_topic_stats_proto_msgTypes = make([]protoimpl.MessageInfo, 4)
var file_google_cloud_pubsublite_v1_topic_stats_proto_goTypes = []interface{}{
	(*ComputeMessageStatsRequest)(nil),  // 0: google.cloud.pubsublite.v1.ComputeMessageStatsRequest
	(*ComputeMessageStatsResponse)(nil), // 1: google.cloud.pubsublite.v1.ComputeMessageStatsResponse
	(*ComputeHeadCursorRequest)(nil),    // 2: google.cloud.pubsublite.v1.ComputeHeadCursorRequest
	(*ComputeHeadCursorResponse)(nil),   // 3: google.cloud.pubsublite.v1.ComputeHeadCursorResponse
	(*Cursor)(nil),                      // 4: google.cloud.pubsublite.v1.Cursor
	(*timestamppb.Timestamp)(nil),       // 5: google.protobuf.Timestamp
}
var file_google_cloud_pubsublite_v1_topic_stats_proto_depIdxs = []int32{
	4, // 0: google.cloud.pubsublite.v1.ComputeMessageStatsRequest.start_cursor:type_name -> google.cloud.pubsublite.v1.Cursor
	4, // 1: google.cloud.pubsublite.v1.ComputeMessageStatsRequest.end_cursor:type_name -> google.cloud.pubsublite.v1.Cursor
	5, // 2: google.cloud.pubsublite.v1.ComputeMessageStatsResponse.minimum_publish_time:type_name -> google.protobuf.Timestamp
	5, // 3: google.cloud.pubsublite.v1.ComputeMessageStatsResponse.minimum_event_time:type_name -> google.protobuf.Timestamp
	4, // 4: google.cloud.pubsublite.v1.ComputeHeadCursorResponse.head_cursor:type_name -> google.cloud.pubsublite.v1.Cursor
	0, // 5: google.cloud.pubsublite.v1.TopicStatsService.ComputeMessageStats:input_type -> google.cloud.pubsublite.v1.ComputeMessageStatsRequest
	2, // 6: google.cloud.pubsublite.v1.TopicStatsService.ComputeHeadCursor:input_type -> google.cloud.pubsublite.v1.ComputeHeadCursorRequest
	1, // 7: google.cloud.pubsublite.v1.TopicStatsService.ComputeMessageStats:output_type -> google.cloud.pubsublite.v1.ComputeMessageStatsResponse
	3, // 8: google.cloud.pubsublite.v1.TopicStatsService.ComputeHeadCursor:output_type -> google.cloud.pubsublite.v1.ComputeHeadCursorResponse
	7, // [7:9] is the sub-list for method output_type
	5, // [5:7] is the sub-list for method input_type
	5, // [5:5] is the sub-list for extension type_name
	5, // [5:5] is the sub-list for extension extendee
	0, // [0:5] is the sub-list for field type_name
}

func init() { file_google_cloud_pubsublite_v1_topic_stats_proto_init() }
func file_google_cloud_pubsublite_v1_topic_stats_proto_init() {
	if File_google_cloud_pubsublite_v1_topic_stats_proto != nil {
		return
	}
	file_google_cloud_pubsublite_v1_common_proto_init()
	if !protoimpl.UnsafeEnabled {
		file_google_cloud_pubsublite_v1_topic_stats_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ComputeMessageStatsRequest); i {
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
		file_google_cloud_pubsublite_v1_topic_stats_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ComputeMessageStatsResponse); i {
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
		file_google_cloud_pubsublite_v1_topic_stats_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ComputeHeadCursorRequest); i {
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
		file_google_cloud_pubsublite_v1_topic_stats_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ComputeHeadCursorResponse); i {
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
			RawDescriptor: file_google_cloud_pubsublite_v1_topic_stats_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   4,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_google_cloud_pubsublite_v1_topic_stats_proto_goTypes,
		DependencyIndexes: file_google_cloud_pubsublite_v1_topic_stats_proto_depIdxs,
		MessageInfos:      file_google_cloud_pubsublite_v1_topic_stats_proto_msgTypes,
	}.Build()
	File_google_cloud_pubsublite_v1_topic_stats_proto = out.File
	file_google_cloud_pubsublite_v1_topic_stats_proto_rawDesc = nil
	file_google_cloud_pubsublite_v1_topic_stats_proto_goTypes = nil
	file_google_cloud_pubsublite_v1_topic_stats_proto_depIdxs = nil
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConnInterface

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion6

// TopicStatsServiceClient is the client API for TopicStatsService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type TopicStatsServiceClient interface {
	// Compute statistics about a range of messages in a given topic and
	// partition.
	ComputeMessageStats(ctx context.Context, in *ComputeMessageStatsRequest, opts ...grpc.CallOption) (*ComputeMessageStatsResponse, error)
	// Compute the head cursor for the partition.
	// The head cursor???s offset is guaranteed to be before or equal to all
	// messages which have not yet been acknowledged to be published, and
	// greater than the offset of any message whose publish has already
	// been acknowledged. It is 0 if there have never been messages on the
	// partition.
	ComputeHeadCursor(ctx context.Context, in *ComputeHeadCursorRequest, opts ...grpc.CallOption) (*ComputeHeadCursorResponse, error)
}

type topicStatsServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewTopicStatsServiceClient(cc grpc.ClientConnInterface) TopicStatsServiceClient {
	return &topicStatsServiceClient{cc}
}

func (c *topicStatsServiceClient) ComputeMessageStats(ctx context.Context, in *ComputeMessageStatsRequest, opts ...grpc.CallOption) (*ComputeMessageStatsResponse, error) {
	out := new(ComputeMessageStatsResponse)
	err := c.cc.Invoke(ctx, "/google.cloud.pubsublite.v1.TopicStatsService/ComputeMessageStats", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *topicStatsServiceClient) ComputeHeadCursor(ctx context.Context, in *ComputeHeadCursorRequest, opts ...grpc.CallOption) (*ComputeHeadCursorResponse, error) {
	out := new(ComputeHeadCursorResponse)
	err := c.cc.Invoke(ctx, "/google.cloud.pubsublite.v1.TopicStatsService/ComputeHeadCursor", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// TopicStatsServiceServer is the server API for TopicStatsService service.
type TopicStatsServiceServer interface {
	// Compute statistics about a range of messages in a given topic and
	// partition.
	ComputeMessageStats(context.Context, *ComputeMessageStatsRequest) (*ComputeMessageStatsResponse, error)
	// Compute the head cursor for the partition.
	// The head cursor???s offset is guaranteed to be before or equal to all
	// messages which have not yet been acknowledged to be published, and
	// greater than the offset of any message whose publish has already
	// been acknowledged. It is 0 if there have never been messages on the
	// partition.
	ComputeHeadCursor(context.Context, *ComputeHeadCursorRequest) (*ComputeHeadCursorResponse, error)
}

// UnimplementedTopicStatsServiceServer can be embedded to have forward compatible implementations.
type UnimplementedTopicStatsServiceServer struct {
}

func (*UnimplementedTopicStatsServiceServer) ComputeMessageStats(context.Context, *ComputeMessageStatsRequest) (*ComputeMessageStatsResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method ComputeMessageStats not implemented")
}
func (*UnimplementedTopicStatsServiceServer) ComputeHeadCursor(context.Context, *ComputeHeadCursorRequest) (*ComputeHeadCursorResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method ComputeHeadCursor not implemented")
}

func RegisterTopicStatsServiceServer(s *grpc.Server, srv TopicStatsServiceServer) {
	s.RegisterService(&_TopicStatsService_serviceDesc, srv)
}

func _TopicStatsService_ComputeMessageStats_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ComputeMessageStatsRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(TopicStatsServiceServer).ComputeMessageStats(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/google.cloud.pubsublite.v1.TopicStatsService/ComputeMessageStats",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(TopicStatsServiceServer).ComputeMessageStats(ctx, req.(*ComputeMessageStatsRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _TopicStatsService_ComputeHeadCursor_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ComputeHeadCursorRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(TopicStatsServiceServer).ComputeHeadCursor(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/google.cloud.pubsublite.v1.TopicStatsService/ComputeHeadCursor",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(TopicStatsServiceServer).ComputeHeadCursor(ctx, req.(*ComputeHeadCursorRequest))
	}
	return interceptor(ctx, in, info, handler)
}

var _TopicStatsService_serviceDesc = grpc.ServiceDesc{
	ServiceName: "google.cloud.pubsublite.v1.TopicStatsService",
	HandlerType: (*TopicStatsServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "ComputeMessageStats",
			Handler:    _TopicStatsService_ComputeMessageStats_Handler,
		},
		{
			MethodName: "ComputeHeadCursor",
			Handler:    _TopicStatsService_ComputeHeadCursor_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "google/cloud/pubsublite/v1/topic_stats.proto",
}
