// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.32.0
// 	protoc        (unknown)
// source: node.proto

package proto

import (
	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
	reflect "reflect"
	sync "sync"
)

const (
	// Verify that this generated code is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(20 - protoimpl.MinVersion)
	// Verify that runtime/protoimpl is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(protoimpl.MaxVersion - 20)
)

type MetadataRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields
}

func (x *MetadataRequest) Reset() {
	*x = MetadataRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_node_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *MetadataRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*MetadataRequest) ProtoMessage() {}

func (x *MetadataRequest) ProtoReflect() protoreflect.Message {
	mi := &file_node_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use MetadataRequest.ProtoReflect.Descriptor instead.
func (*MetadataRequest) Descriptor() ([]byte, []int) {
	return file_node_proto_rawDescGZIP(), []int{0}
}

type MetadataReply struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id string `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
}

func (x *MetadataReply) Reset() {
	*x = MetadataReply{}
	if protoimpl.UnsafeEnabled {
		mi := &file_node_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *MetadataReply) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*MetadataReply) ProtoMessage() {}

func (x *MetadataReply) ProtoReflect() protoreflect.Message {
	mi := &file_node_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use MetadataReply.ProtoReflect.Descriptor instead.
func (*MetadataReply) Descriptor() ([]byte, []int) {
	return file_node_proto_rawDescGZIP(), []int{1}
}

func (x *MetadataReply) GetId() string {
	if x != nil {
		return x.Id
	}
	return ""
}

type VoteRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	SourceId     string `protobuf:"bytes,1,opt,name=source_id,json=sourceId,proto3" json:"source_id,omitempty"`
	TargeId      string `protobuf:"bytes,2,opt,name=targe_id,json=targeId,proto3" json:"targe_id,omitempty"`
	Term         int64  `protobuf:"varint,3,opt,name=term,proto3" json:"term,omitempty"`
	CommitOffset int64  `protobuf:"varint,4,opt,name=commit_offset,json=commitOffset,proto3" json:"commit_offset,omitempty"`
}

func (x *VoteRequest) Reset() {
	*x = VoteRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_node_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *VoteRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*VoteRequest) ProtoMessage() {}

func (x *VoteRequest) ProtoReflect() protoreflect.Message {
	mi := &file_node_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use VoteRequest.ProtoReflect.Descriptor instead.
func (*VoteRequest) Descriptor() ([]byte, []int) {
	return file_node_proto_rawDescGZIP(), []int{2}
}

func (x *VoteRequest) GetSourceId() string {
	if x != nil {
		return x.SourceId
	}
	return ""
}

func (x *VoteRequest) GetTargeId() string {
	if x != nil {
		return x.TargeId
	}
	return ""
}

func (x *VoteRequest) GetTerm() int64 {
	if x != nil {
		return x.Term
	}
	return 0
}

func (x *VoteRequest) GetCommitOffset() int64 {
	if x != nil {
		return x.CommitOffset
	}
	return 0
}

type VoteReply struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	SourceId string `protobuf:"bytes,1,opt,name=source_id,json=sourceId,proto3" json:"source_id,omitempty"`
	TargeId  string `protobuf:"bytes,2,opt,name=targe_id,json=targeId,proto3" json:"targe_id,omitempty"`
	Granted  bool   `protobuf:"varint,3,opt,name=granted,proto3" json:"granted,omitempty"`
}

func (x *VoteReply) Reset() {
	*x = VoteReply{}
	if protoimpl.UnsafeEnabled {
		mi := &file_node_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *VoteReply) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*VoteReply) ProtoMessage() {}

func (x *VoteReply) ProtoReflect() protoreflect.Message {
	mi := &file_node_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use VoteReply.ProtoReflect.Descriptor instead.
func (*VoteReply) Descriptor() ([]byte, []int) {
	return file_node_proto_rawDescGZIP(), []int{3}
}

func (x *VoteReply) GetSourceId() string {
	if x != nil {
		return x.SourceId
	}
	return ""
}

func (x *VoteReply) GetTargeId() string {
	if x != nil {
		return x.TargeId
	}
	return ""
}

func (x *VoteReply) GetGranted() bool {
	if x != nil {
		return x.Granted
	}
	return false
}

type LogRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	SourceId     string `protobuf:"bytes,1,opt,name=source_id,json=sourceId,proto3" json:"source_id,omitempty"`
	TargeId      string `protobuf:"bytes,2,opt,name=targe_id,json=targeId,proto3" json:"targe_id,omitempty"`
	CommitOffset int64  `protobuf:"varint,3,opt,name=commit_offset,json=commitOffset,proto3" json:"commit_offset,omitempty"`
	StartOffset  int64  `protobuf:"varint,4,opt,name=start_offset,json=startOffset,proto3" json:"start_offset,omitempty"`
	Entries      []byte `protobuf:"bytes,5,opt,name=entries,proto3" json:"entries,omitempty"`
}

func (x *LogRequest) Reset() {
	*x = LogRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_node_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *LogRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*LogRequest) ProtoMessage() {}

func (x *LogRequest) ProtoReflect() protoreflect.Message {
	mi := &file_node_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use LogRequest.ProtoReflect.Descriptor instead.
func (*LogRequest) Descriptor() ([]byte, []int) {
	return file_node_proto_rawDescGZIP(), []int{4}
}

func (x *LogRequest) GetSourceId() string {
	if x != nil {
		return x.SourceId
	}
	return ""
}

func (x *LogRequest) GetTargeId() string {
	if x != nil {
		return x.TargeId
	}
	return ""
}

func (x *LogRequest) GetCommitOffset() int64 {
	if x != nil {
		return x.CommitOffset
	}
	return 0
}

func (x *LogRequest) GetStartOffset() int64 {
	if x != nil {
		return x.StartOffset
	}
	return 0
}

func (x *LogRequest) GetEntries() []byte {
	if x != nil {
		return x.Entries
	}
	return nil
}

type LogReply struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	SourceId     string `protobuf:"bytes,1,opt,name=source_id,json=sourceId,proto3" json:"source_id,omitempty"`
	TargeId      string `protobuf:"bytes,2,opt,name=targe_id,json=targeId,proto3" json:"targe_id,omitempty"`
	CommitOffset int64  `protobuf:"varint,3,opt,name=commit_offset,json=commitOffset,proto3" json:"commit_offset,omitempty"`
	Success      bool   `protobuf:"varint,4,opt,name=success,proto3" json:"success,omitempty"`
}

func (x *LogReply) Reset() {
	*x = LogReply{}
	if protoimpl.UnsafeEnabled {
		mi := &file_node_proto_msgTypes[5]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *LogReply) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*LogReply) ProtoMessage() {}

func (x *LogReply) ProtoReflect() protoreflect.Message {
	mi := &file_node_proto_msgTypes[5]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use LogReply.ProtoReflect.Descriptor instead.
func (*LogReply) Descriptor() ([]byte, []int) {
	return file_node_proto_rawDescGZIP(), []int{5}
}

func (x *LogReply) GetSourceId() string {
	if x != nil {
		return x.SourceId
	}
	return ""
}

func (x *LogReply) GetTargeId() string {
	if x != nil {
		return x.TargeId
	}
	return ""
}

func (x *LogReply) GetCommitOffset() int64 {
	if x != nil {
		return x.CommitOffset
	}
	return 0
}

func (x *LogReply) GetSuccess() bool {
	if x != nil {
		return x.Success
	}
	return false
}

type ProxyPushRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	SourceId string `protobuf:"bytes,1,opt,name=source_id,json=sourceId,proto3" json:"source_id,omitempty"`
	TargeId  string `protobuf:"bytes,2,opt,name=targe_id,json=targeId,proto3" json:"targe_id,omitempty"`
	Push     []byte `protobuf:"bytes,3,opt,name=push,proto3" json:"push,omitempty"`
}

func (x *ProxyPushRequest) Reset() {
	*x = ProxyPushRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_node_proto_msgTypes[6]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ProxyPushRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ProxyPushRequest) ProtoMessage() {}

func (x *ProxyPushRequest) ProtoReflect() protoreflect.Message {
	mi := &file_node_proto_msgTypes[6]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ProxyPushRequest.ProtoReflect.Descriptor instead.
func (*ProxyPushRequest) Descriptor() ([]byte, []int) {
	return file_node_proto_rawDescGZIP(), []int{6}
}

func (x *ProxyPushRequest) GetSourceId() string {
	if x != nil {
		return x.SourceId
	}
	return ""
}

func (x *ProxyPushRequest) GetTargeId() string {
	if x != nil {
		return x.TargeId
	}
	return ""
}

func (x *ProxyPushRequest) GetPush() []byte {
	if x != nil {
		return x.Push
	}
	return nil
}

type ProxyPushReply struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	SourceId string `protobuf:"bytes,1,opt,name=source_id,json=sourceId,proto3" json:"source_id,omitempty"`
	TargeId  string `protobuf:"bytes,2,opt,name=targe_id,json=targeId,proto3" json:"targe_id,omitempty"`
	Ok       bool   `protobuf:"varint,3,opt,name=ok,proto3" json:"ok,omitempty"`
}

func (x *ProxyPushReply) Reset() {
	*x = ProxyPushReply{}
	if protoimpl.UnsafeEnabled {
		mi := &file_node_proto_msgTypes[7]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ProxyPushReply) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ProxyPushReply) ProtoMessage() {}

func (x *ProxyPushReply) ProtoReflect() protoreflect.Message {
	mi := &file_node_proto_msgTypes[7]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ProxyPushReply.ProtoReflect.Descriptor instead.
func (*ProxyPushReply) Descriptor() ([]byte, []int) {
	return file_node_proto_rawDescGZIP(), []int{7}
}

func (x *ProxyPushReply) GetSourceId() string {
	if x != nil {
		return x.SourceId
	}
	return ""
}

func (x *ProxyPushReply) GetTargeId() string {
	if x != nil {
		return x.TargeId
	}
	return ""
}

func (x *ProxyPushReply) GetOk() bool {
	if x != nil {
		return x.Ok
	}
	return false
}

var File_node_proto protoreflect.FileDescriptor

var file_node_proto_rawDesc = []byte{
	0x0a, 0x0a, 0x6e, 0x6f, 0x64, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x09, 0x73, 0x68,
	0x69, 0x70, 0x77, 0x72, 0x65, 0x63, 0x6b, 0x22, 0x11, 0x0a, 0x0f, 0x4d, 0x65, 0x74, 0x61, 0x64,
	0x61, 0x74, 0x61, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x22, 0x1f, 0x0a, 0x0d, 0x4d, 0x65,
	0x74, 0x61, 0x64, 0x61, 0x74, 0x61, 0x52, 0x65, 0x70, 0x6c, 0x79, 0x12, 0x0e, 0x0a, 0x02, 0x69,
	0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x02, 0x69, 0x64, 0x22, 0x7e, 0x0a, 0x0b, 0x56,
	0x6f, 0x74, 0x65, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x1b, 0x0a, 0x09, 0x73, 0x6f,
	0x75, 0x72, 0x63, 0x65, 0x5f, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x73,
	0x6f, 0x75, 0x72, 0x63, 0x65, 0x49, 0x64, 0x12, 0x19, 0x0a, 0x08, 0x74, 0x61, 0x72, 0x67, 0x65,
	0x5f, 0x69, 0x64, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x07, 0x74, 0x61, 0x72, 0x67, 0x65,
	0x49, 0x64, 0x12, 0x12, 0x0a, 0x04, 0x74, 0x65, 0x72, 0x6d, 0x18, 0x03, 0x20, 0x01, 0x28, 0x03,
	0x52, 0x04, 0x74, 0x65, 0x72, 0x6d, 0x12, 0x23, 0x0a, 0x0d, 0x63, 0x6f, 0x6d, 0x6d, 0x69, 0x74,
	0x5f, 0x6f, 0x66, 0x66, 0x73, 0x65, 0x74, 0x18, 0x04, 0x20, 0x01, 0x28, 0x03, 0x52, 0x0c, 0x63,
	0x6f, 0x6d, 0x6d, 0x69, 0x74, 0x4f, 0x66, 0x66, 0x73, 0x65, 0x74, 0x22, 0x5d, 0x0a, 0x09, 0x56,
	0x6f, 0x74, 0x65, 0x52, 0x65, 0x70, 0x6c, 0x79, 0x12, 0x1b, 0x0a, 0x09, 0x73, 0x6f, 0x75, 0x72,
	0x63, 0x65, 0x5f, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x73, 0x6f, 0x75,
	0x72, 0x63, 0x65, 0x49, 0x64, 0x12, 0x19, 0x0a, 0x08, 0x74, 0x61, 0x72, 0x67, 0x65, 0x5f, 0x69,
	0x64, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x07, 0x74, 0x61, 0x72, 0x67, 0x65, 0x49, 0x64,
	0x12, 0x18, 0x0a, 0x07, 0x67, 0x72, 0x61, 0x6e, 0x74, 0x65, 0x64, 0x18, 0x03, 0x20, 0x01, 0x28,
	0x08, 0x52, 0x07, 0x67, 0x72, 0x61, 0x6e, 0x74, 0x65, 0x64, 0x22, 0xa6, 0x01, 0x0a, 0x0a, 0x4c,
	0x6f, 0x67, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x1b, 0x0a, 0x09, 0x73, 0x6f, 0x75,
	0x72, 0x63, 0x65, 0x5f, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x73, 0x6f,
	0x75, 0x72, 0x63, 0x65, 0x49, 0x64, 0x12, 0x19, 0x0a, 0x08, 0x74, 0x61, 0x72, 0x67, 0x65, 0x5f,
	0x69, 0x64, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x07, 0x74, 0x61, 0x72, 0x67, 0x65, 0x49,
	0x64, 0x12, 0x23, 0x0a, 0x0d, 0x63, 0x6f, 0x6d, 0x6d, 0x69, 0x74, 0x5f, 0x6f, 0x66, 0x66, 0x73,
	0x65, 0x74, 0x18, 0x03, 0x20, 0x01, 0x28, 0x03, 0x52, 0x0c, 0x63, 0x6f, 0x6d, 0x6d, 0x69, 0x74,
	0x4f, 0x66, 0x66, 0x73, 0x65, 0x74, 0x12, 0x21, 0x0a, 0x0c, 0x73, 0x74, 0x61, 0x72, 0x74, 0x5f,
	0x6f, 0x66, 0x66, 0x73, 0x65, 0x74, 0x18, 0x04, 0x20, 0x01, 0x28, 0x03, 0x52, 0x0b, 0x73, 0x74,
	0x61, 0x72, 0x74, 0x4f, 0x66, 0x66, 0x73, 0x65, 0x74, 0x12, 0x18, 0x0a, 0x07, 0x65, 0x6e, 0x74,
	0x72, 0x69, 0x65, 0x73, 0x18, 0x05, 0x20, 0x01, 0x28, 0x0c, 0x52, 0x07, 0x65, 0x6e, 0x74, 0x72,
	0x69, 0x65, 0x73, 0x22, 0x81, 0x01, 0x0a, 0x08, 0x4c, 0x6f, 0x67, 0x52, 0x65, 0x70, 0x6c, 0x79,
	0x12, 0x1b, 0x0a, 0x09, 0x73, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x5f, 0x69, 0x64, 0x18, 0x01, 0x20,
	0x01, 0x28, 0x09, 0x52, 0x08, 0x73, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x49, 0x64, 0x12, 0x19, 0x0a,
	0x08, 0x74, 0x61, 0x72, 0x67, 0x65, 0x5f, 0x69, 0x64, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52,
	0x07, 0x74, 0x61, 0x72, 0x67, 0x65, 0x49, 0x64, 0x12, 0x23, 0x0a, 0x0d, 0x63, 0x6f, 0x6d, 0x6d,
	0x69, 0x74, 0x5f, 0x6f, 0x66, 0x66, 0x73, 0x65, 0x74, 0x18, 0x03, 0x20, 0x01, 0x28, 0x03, 0x52,
	0x0c, 0x63, 0x6f, 0x6d, 0x6d, 0x69, 0x74, 0x4f, 0x66, 0x66, 0x73, 0x65, 0x74, 0x12, 0x18, 0x0a,
	0x07, 0x73, 0x75, 0x63, 0x63, 0x65, 0x73, 0x73, 0x18, 0x04, 0x20, 0x01, 0x28, 0x08, 0x52, 0x07,
	0x73, 0x75, 0x63, 0x63, 0x65, 0x73, 0x73, 0x22, 0x5e, 0x0a, 0x10, 0x50, 0x72, 0x6f, 0x78, 0x79,
	0x50, 0x75, 0x73, 0x68, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x1b, 0x0a, 0x09, 0x73,
	0x6f, 0x75, 0x72, 0x63, 0x65, 0x5f, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08,
	0x73, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x49, 0x64, 0x12, 0x19, 0x0a, 0x08, 0x74, 0x61, 0x72, 0x67,
	0x65, 0x5f, 0x69, 0x64, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x07, 0x74, 0x61, 0x72, 0x67,
	0x65, 0x49, 0x64, 0x12, 0x12, 0x0a, 0x04, 0x70, 0x75, 0x73, 0x68, 0x18, 0x03, 0x20, 0x01, 0x28,
	0x0c, 0x52, 0x04, 0x70, 0x75, 0x73, 0x68, 0x22, 0x58, 0x0a, 0x0e, 0x50, 0x72, 0x6f, 0x78, 0x79,
	0x50, 0x75, 0x73, 0x68, 0x52, 0x65, 0x70, 0x6c, 0x79, 0x12, 0x1b, 0x0a, 0x09, 0x73, 0x6f, 0x75,
	0x72, 0x63, 0x65, 0x5f, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x73, 0x6f,
	0x75, 0x72, 0x63, 0x65, 0x49, 0x64, 0x12, 0x19, 0x0a, 0x08, 0x74, 0x61, 0x72, 0x67, 0x65, 0x5f,
	0x69, 0x64, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x07, 0x74, 0x61, 0x72, 0x67, 0x65, 0x49,
	0x64, 0x12, 0x0e, 0x0a, 0x02, 0x6f, 0x6b, 0x18, 0x03, 0x20, 0x01, 0x28, 0x08, 0x52, 0x02, 0x6f,
	0x6b, 0x32, 0x92, 0x02, 0x0a, 0x04, 0x4e, 0x6f, 0x64, 0x65, 0x12, 0x45, 0x0a, 0x0b, 0x47, 0x65,
	0x74, 0x4d, 0x65, 0x74, 0x61, 0x64, 0x61, 0x74, 0x61, 0x12, 0x1a, 0x2e, 0x73, 0x68, 0x69, 0x70,
	0x77, 0x72, 0x65, 0x63, 0x6b, 0x2e, 0x4d, 0x65, 0x74, 0x61, 0x64, 0x61, 0x74, 0x61, 0x52, 0x65,
	0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x18, 0x2e, 0x73, 0x68, 0x69, 0x70, 0x77, 0x72, 0x65, 0x63,
	0x6b, 0x2e, 0x4d, 0x65, 0x74, 0x61, 0x64, 0x61, 0x74, 0x61, 0x52, 0x65, 0x70, 0x6c, 0x79, 0x22,
	0x00, 0x12, 0x3d, 0x0a, 0x0b, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x56, 0x6f, 0x74, 0x65,
	0x12, 0x16, 0x2e, 0x73, 0x68, 0x69, 0x70, 0x77, 0x72, 0x65, 0x63, 0x6b, 0x2e, 0x56, 0x6f, 0x74,
	0x65, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x14, 0x2e, 0x73, 0x68, 0x69, 0x70, 0x77,
	0x72, 0x65, 0x63, 0x6b, 0x2e, 0x56, 0x6f, 0x74, 0x65, 0x52, 0x65, 0x70, 0x6c, 0x79, 0x22, 0x00,
	0x12, 0x3d, 0x0a, 0x0d, 0x41, 0x70, 0x70, 0x65, 0x6e, 0x64, 0x45, 0x6e, 0x74, 0x72, 0x69, 0x65,
	0x73, 0x12, 0x15, 0x2e, 0x73, 0x68, 0x69, 0x70, 0x77, 0x72, 0x65, 0x63, 0x6b, 0x2e, 0x4c, 0x6f,
	0x67, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x13, 0x2e, 0x73, 0x68, 0x69, 0x70, 0x77,
	0x72, 0x65, 0x63, 0x6b, 0x2e, 0x4c, 0x6f, 0x67, 0x52, 0x65, 0x70, 0x6c, 0x79, 0x22, 0x00, 0x12,
	0x45, 0x0a, 0x09, 0x50, 0x72, 0x6f, 0x78, 0x79, 0x50, 0x75, 0x73, 0x68, 0x12, 0x1b, 0x2e, 0x73,
	0x68, 0x69, 0x70, 0x77, 0x72, 0x65, 0x63, 0x6b, 0x2e, 0x50, 0x72, 0x6f, 0x78, 0x79, 0x50, 0x75,
	0x73, 0x68, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x19, 0x2e, 0x73, 0x68, 0x69, 0x70,
	0x77, 0x72, 0x65, 0x63, 0x6b, 0x2e, 0x50, 0x72, 0x6f, 0x78, 0x79, 0x50, 0x75, 0x73, 0x68, 0x52,
	0x65, 0x70, 0x6c, 0x79, 0x22, 0x00, 0x42, 0x6f, 0x0a, 0x0d, 0x63, 0x6f, 0x6d, 0x2e, 0x73, 0x68,
	0x69, 0x70, 0x77, 0x72, 0x65, 0x63, 0x6b, 0x42, 0x09, 0x4e, 0x6f, 0x64, 0x65, 0x50, 0x72, 0x6f,
	0x74, 0x6f, 0x50, 0x01, 0x5a, 0x0f, 0x73, 0x68, 0x69, 0x70, 0x77, 0x72, 0x65, 0x63, 0x6b, 0x2f,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0xa2, 0x02, 0x03, 0x53, 0x58, 0x58, 0xaa, 0x02, 0x09, 0x53, 0x68,
	0x69, 0x70, 0x77, 0x72, 0x65, 0x63, 0x6b, 0xca, 0x02, 0x09, 0x53, 0x68, 0x69, 0x70, 0x77, 0x72,
	0x65, 0x63, 0x6b, 0xe2, 0x02, 0x15, 0x53, 0x68, 0x69, 0x70, 0x77, 0x72, 0x65, 0x63, 0x6b, 0x5c,
	0x47, 0x50, 0x42, 0x4d, 0x65, 0x74, 0x61, 0x64, 0x61, 0x74, 0x61, 0xea, 0x02, 0x09, 0x53, 0x68,
	0x69, 0x70, 0x77, 0x72, 0x65, 0x63, 0x6b, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_node_proto_rawDescOnce sync.Once
	file_node_proto_rawDescData = file_node_proto_rawDesc
)

func file_node_proto_rawDescGZIP() []byte {
	file_node_proto_rawDescOnce.Do(func() {
		file_node_proto_rawDescData = protoimpl.X.CompressGZIP(file_node_proto_rawDescData)
	})
	return file_node_proto_rawDescData
}

var file_node_proto_msgTypes = make([]protoimpl.MessageInfo, 8)
var file_node_proto_goTypes = []interface{}{
	(*MetadataRequest)(nil),  // 0: shipwreck.MetadataRequest
	(*MetadataReply)(nil),    // 1: shipwreck.MetadataReply
	(*VoteRequest)(nil),      // 2: shipwreck.VoteRequest
	(*VoteReply)(nil),        // 3: shipwreck.VoteReply
	(*LogRequest)(nil),       // 4: shipwreck.LogRequest
	(*LogReply)(nil),         // 5: shipwreck.LogReply
	(*ProxyPushRequest)(nil), // 6: shipwreck.ProxyPushRequest
	(*ProxyPushReply)(nil),   // 7: shipwreck.ProxyPushReply
}
var file_node_proto_depIdxs = []int32{
	0, // 0: shipwreck.Node.GetMetadata:input_type -> shipwreck.MetadataRequest
	2, // 1: shipwreck.Node.RequestVote:input_type -> shipwreck.VoteRequest
	4, // 2: shipwreck.Node.AppendEntries:input_type -> shipwreck.LogRequest
	6, // 3: shipwreck.Node.ProxyPush:input_type -> shipwreck.ProxyPushRequest
	1, // 4: shipwreck.Node.GetMetadata:output_type -> shipwreck.MetadataReply
	3, // 5: shipwreck.Node.RequestVote:output_type -> shipwreck.VoteReply
	5, // 6: shipwreck.Node.AppendEntries:output_type -> shipwreck.LogReply
	7, // 7: shipwreck.Node.ProxyPush:output_type -> shipwreck.ProxyPushReply
	4, // [4:8] is the sub-list for method output_type
	0, // [0:4] is the sub-list for method input_type
	0, // [0:0] is the sub-list for extension type_name
	0, // [0:0] is the sub-list for extension extendee
	0, // [0:0] is the sub-list for field type_name
}

func init() { file_node_proto_init() }
func file_node_proto_init() {
	if File_node_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_node_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*MetadataRequest); i {
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
		file_node_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*MetadataReply); i {
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
		file_node_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*VoteRequest); i {
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
		file_node_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*VoteReply); i {
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
		file_node_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*LogRequest); i {
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
		file_node_proto_msgTypes[5].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*LogReply); i {
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
		file_node_proto_msgTypes[6].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ProxyPushRequest); i {
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
		file_node_proto_msgTypes[7].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ProxyPushReply); i {
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
			RawDescriptor: file_node_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   8,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_node_proto_goTypes,
		DependencyIndexes: file_node_proto_depIdxs,
		MessageInfos:      file_node_proto_msgTypes,
	}.Build()
	File_node_proto = out.File
	file_node_proto_rawDesc = nil
	file_node_proto_goTypes = nil
	file_node_proto_depIdxs = nil
}
