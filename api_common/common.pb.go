// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.34.2
// 	protoc        v5.28.1
// source: common.proto

package __

import (
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

type ProcessTxnRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	MsgID    string  `protobuf:"bytes,1,opt,name=msgID,proto3" json:"msgID,omitempty"`
	Sender   string  `protobuf:"bytes,2,opt,name=sender,proto3" json:"sender,omitempty"`
	Receiver string  `protobuf:"bytes,3,opt,name=receiver,proto3" json:"receiver,omitempty"`
	Amount   float32 `protobuf:"fixed32,4,opt,name=amount,proto3" json:"amount,omitempty"`
}

func (x *ProcessTxnRequest) Reset() {
	*x = ProcessTxnRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_common_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ProcessTxnRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ProcessTxnRequest) ProtoMessage() {}

func (x *ProcessTxnRequest) ProtoReflect() protoreflect.Message {
	mi := &file_common_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ProcessTxnRequest.ProtoReflect.Descriptor instead.
func (*ProcessTxnRequest) Descriptor() ([]byte, []int) {
	return file_common_proto_rawDescGZIP(), []int{0}
}

func (x *ProcessTxnRequest) GetMsgID() string {
	if x != nil {
		return x.MsgID
	}
	return ""
}

func (x *ProcessTxnRequest) GetSender() string {
	if x != nil {
		return x.Sender
	}
	return ""
}

func (x *ProcessTxnRequest) GetReceiver() string {
	if x != nil {
		return x.Receiver
	}
	return ""
}

func (x *ProcessTxnRequest) GetAmount() float32 {
	if x != nil {
		return x.Amount
	}
	return 0
}

type Ballot struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	TermNumber   int32 `protobuf:"varint,1,opt,name=termNumber,proto3" json:"termNumber,omitempty"`
	ServerNumber int32 `protobuf:"varint,2,opt,name=serverNumber,proto3" json:"serverNumber,omitempty"`
}

func (x *Ballot) Reset() {
	*x = Ballot{}
	if protoimpl.UnsafeEnabled {
		mi := &file_common_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Ballot) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Ballot) ProtoMessage() {}

func (x *Ballot) ProtoReflect() protoreflect.Message {
	mi := &file_common_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Ballot.ProtoReflect.Descriptor instead.
func (*Ballot) Descriptor() ([]byte, []int) {
	return file_common_proto_rawDescGZIP(), []int{1}
}

func (x *Ballot) GetTermNumber() int32 {
	if x != nil {
		return x.TermNumber
	}
	return 0
}

func (x *Ballot) GetServerNumber() int32 {
	if x != nil {
		return x.ServerNumber
	}
	return 0
}

type Prepare struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	BallotNum *Ballot `protobuf:"bytes,1,opt,name=ballotNum,proto3" json:"ballotNum,omitempty"`
}

func (x *Prepare) Reset() {
	*x = Prepare{}
	if protoimpl.UnsafeEnabled {
		mi := &file_common_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Prepare) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Prepare) ProtoMessage() {}

func (x *Prepare) ProtoReflect() protoreflect.Message {
	mi := &file_common_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Prepare.ProtoReflect.Descriptor instead.
func (*Prepare) Descriptor() ([]byte, []int) {
	return file_common_proto_rawDescGZIP(), []int{2}
}

func (x *Prepare) GetBallotNum() *Ballot {
	if x != nil {
		return x.BallotNum
	}
	return nil
}

type Promise struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	PromiseAck   bool                 `protobuf:"varint,1,opt,name=promiseAck,proto3" json:"promiseAck,omitempty"`
	ServerNumber int32                `protobuf:"varint,2,opt,name=serverNumber,proto3" json:"serverNumber,omitempty"`
	BallotNum    *Ballot              `protobuf:"bytes,3,opt,name=ballotNum,proto3" json:"ballotNum,omitempty"`
	AcceptNum    *Ballot              `protobuf:"bytes,4,opt,name=acceptNum,proto3" json:"acceptNum,omitempty"`
	AcceptVal    []*ProcessTxnRequest `protobuf:"bytes,5,rep,name=acceptVal,proto3" json:"acceptVal,omitempty"`
	LocalVal     []*ProcessTxnRequest `protobuf:"bytes,6,rep,name=localVal,proto3" json:"localVal,omitempty"`
}

func (x *Promise) Reset() {
	*x = Promise{}
	if protoimpl.UnsafeEnabled {
		mi := &file_common_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Promise) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Promise) ProtoMessage() {}

func (x *Promise) ProtoReflect() protoreflect.Message {
	mi := &file_common_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Promise.ProtoReflect.Descriptor instead.
func (*Promise) Descriptor() ([]byte, []int) {
	return file_common_proto_rawDescGZIP(), []int{3}
}

func (x *Promise) GetPromiseAck() bool {
	if x != nil {
		return x.PromiseAck
	}
	return false
}

func (x *Promise) GetServerNumber() int32 {
	if x != nil {
		return x.ServerNumber
	}
	return 0
}

func (x *Promise) GetBallotNum() *Ballot {
	if x != nil {
		return x.BallotNum
	}
	return nil
}

func (x *Promise) GetAcceptNum() *Ballot {
	if x != nil {
		return x.AcceptNum
	}
	return nil
}

func (x *Promise) GetAcceptVal() []*ProcessTxnRequest {
	if x != nil {
		return x.AcceptVal
	}
	return nil
}

func (x *Promise) GetLocalVal() []*ProcessTxnRequest {
	if x != nil {
		return x.LocalVal
	}
	return nil
}

type Accept struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	BallotNum       *Ballot              `protobuf:"bytes,1,opt,name=ballotNum,proto3" json:"ballotNum,omitempty"`
	AcceptVal       []*ProcessTxnRequest `protobuf:"bytes,2,rep,name=acceptVal,proto3" json:"acceptVal,omitempty"`
	ServerAddresses []string             `protobuf:"bytes,3,rep,name=serverAddresses,proto3" json:"serverAddresses,omitempty"`
}

func (x *Accept) Reset() {
	*x = Accept{}
	if protoimpl.UnsafeEnabled {
		mi := &file_common_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Accept) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Accept) ProtoMessage() {}

func (x *Accept) ProtoReflect() protoreflect.Message {
	mi := &file_common_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Accept.ProtoReflect.Descriptor instead.
func (*Accept) Descriptor() ([]byte, []int) {
	return file_common_proto_rawDescGZIP(), []int{4}
}

func (x *Accept) GetBallotNum() *Ballot {
	if x != nil {
		return x.BallotNum
	}
	return nil
}

func (x *Accept) GetAcceptVal() []*ProcessTxnRequest {
	if x != nil {
		return x.AcceptVal
	}
	return nil
}

func (x *Accept) GetServerAddresses() []string {
	if x != nil {
		return x.ServerAddresses
	}
	return nil
}

type Accepted struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	BallotNum    *Ballot              `protobuf:"bytes,1,opt,name=ballotNum,proto3" json:"ballotNum,omitempty"`
	AcceptVal    []*ProcessTxnRequest `protobuf:"bytes,2,rep,name=acceptVal,proto3" json:"acceptVal,omitempty"`
	ServerNumber int32                `protobuf:"varint,3,opt,name=serverNumber,proto3" json:"serverNumber,omitempty"`
}

func (x *Accepted) Reset() {
	*x = Accepted{}
	if protoimpl.UnsafeEnabled {
		mi := &file_common_proto_msgTypes[5]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Accepted) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Accepted) ProtoMessage() {}

func (x *Accepted) ProtoReflect() protoreflect.Message {
	mi := &file_common_proto_msgTypes[5]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Accepted.ProtoReflect.Descriptor instead.
func (*Accepted) Descriptor() ([]byte, []int) {
	return file_common_proto_rawDescGZIP(), []int{5}
}

func (x *Accepted) GetBallotNum() *Ballot {
	if x != nil {
		return x.BallotNum
	}
	return nil
}

func (x *Accepted) GetAcceptVal() []*ProcessTxnRequest {
	if x != nil {
		return x.AcceptVal
	}
	return nil
}

func (x *Accepted) GetServerNumber() int32 {
	if x != nil {
		return x.ServerNumber
	}
	return 0
}

type Commit struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	BallotNum       *Ballot              `protobuf:"bytes,1,opt,name=ballotNum,proto3" json:"ballotNum,omitempty"`
	AcceptVal       []*ProcessTxnRequest `protobuf:"bytes,2,rep,name=acceptVal,proto3" json:"acceptVal,omitempty"`
	ServerAddresses []string             `protobuf:"bytes,3,rep,name=serverAddresses,proto3" json:"serverAddresses,omitempty"`
}

func (x *Commit) Reset() {
	*x = Commit{}
	if protoimpl.UnsafeEnabled {
		mi := &file_common_proto_msgTypes[6]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Commit) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Commit) ProtoMessage() {}

func (x *Commit) ProtoReflect() protoreflect.Message {
	mi := &file_common_proto_msgTypes[6]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Commit.ProtoReflect.Descriptor instead.
func (*Commit) Descriptor() ([]byte, []int) {
	return file_common_proto_rawDescGZIP(), []int{6}
}

func (x *Commit) GetBallotNum() *Ballot {
	if x != nil {
		return x.BallotNum
	}
	return nil
}

func (x *Commit) GetAcceptVal() []*ProcessTxnRequest {
	if x != nil {
		return x.AcceptVal
	}
	return nil
}

func (x *Commit) GetServerAddresses() []string {
	if x != nil {
		return x.ServerAddresses
	}
	return nil
}

var File_common_proto protoreflect.FileDescriptor

var file_common_proto_rawDesc = []byte{
	0x0a, 0x0c, 0x63, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x06,
	0x63, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x1a, 0x1b, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2f, 0x65, 0x6d, 0x70, 0x74, 0x79, 0x2e, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x22, 0x75, 0x0a, 0x11, 0x50, 0x72, 0x6f, 0x63, 0x65, 0x73, 0x73, 0x54, 0x78,
	0x6e, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x14, 0x0a, 0x05, 0x6d, 0x73, 0x67, 0x49,
	0x44, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x6d, 0x73, 0x67, 0x49, 0x44, 0x12, 0x16,
	0x0a, 0x06, 0x73, 0x65, 0x6e, 0x64, 0x65, 0x72, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x06,
	0x73, 0x65, 0x6e, 0x64, 0x65, 0x72, 0x12, 0x1a, 0x0a, 0x08, 0x72, 0x65, 0x63, 0x65, 0x69, 0x76,
	0x65, 0x72, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x72, 0x65, 0x63, 0x65, 0x69, 0x76,
	0x65, 0x72, 0x12, 0x16, 0x0a, 0x06, 0x61, 0x6d, 0x6f, 0x75, 0x6e, 0x74, 0x18, 0x04, 0x20, 0x01,
	0x28, 0x02, 0x52, 0x06, 0x61, 0x6d, 0x6f, 0x75, 0x6e, 0x74, 0x22, 0x4c, 0x0a, 0x06, 0x42, 0x61,
	0x6c, 0x6c, 0x6f, 0x74, 0x12, 0x1e, 0x0a, 0x0a, 0x74, 0x65, 0x72, 0x6d, 0x4e, 0x75, 0x6d, 0x62,
	0x65, 0x72, 0x18, 0x01, 0x20, 0x01, 0x28, 0x05, 0x52, 0x0a, 0x74, 0x65, 0x72, 0x6d, 0x4e, 0x75,
	0x6d, 0x62, 0x65, 0x72, 0x12, 0x22, 0x0a, 0x0c, 0x73, 0x65, 0x72, 0x76, 0x65, 0x72, 0x4e, 0x75,
	0x6d, 0x62, 0x65, 0x72, 0x18, 0x02, 0x20, 0x01, 0x28, 0x05, 0x52, 0x0c, 0x73, 0x65, 0x72, 0x76,
	0x65, 0x72, 0x4e, 0x75, 0x6d, 0x62, 0x65, 0x72, 0x22, 0x37, 0x0a, 0x07, 0x50, 0x72, 0x65, 0x70,
	0x61, 0x72, 0x65, 0x12, 0x2c, 0x0a, 0x09, 0x62, 0x61, 0x6c, 0x6c, 0x6f, 0x74, 0x4e, 0x75, 0x6d,
	0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x0e, 0x2e, 0x63, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x2e,
	0x42, 0x61, 0x6c, 0x6c, 0x6f, 0x74, 0x52, 0x09, 0x62, 0x61, 0x6c, 0x6c, 0x6f, 0x74, 0x4e, 0x75,
	0x6d, 0x22, 0x99, 0x02, 0x0a, 0x07, 0x50, 0x72, 0x6f, 0x6d, 0x69, 0x73, 0x65, 0x12, 0x1e, 0x0a,
	0x0a, 0x70, 0x72, 0x6f, 0x6d, 0x69, 0x73, 0x65, 0x41, 0x63, 0x6b, 0x18, 0x01, 0x20, 0x01, 0x28,
	0x08, 0x52, 0x0a, 0x70, 0x72, 0x6f, 0x6d, 0x69, 0x73, 0x65, 0x41, 0x63, 0x6b, 0x12, 0x22, 0x0a,
	0x0c, 0x73, 0x65, 0x72, 0x76, 0x65, 0x72, 0x4e, 0x75, 0x6d, 0x62, 0x65, 0x72, 0x18, 0x02, 0x20,
	0x01, 0x28, 0x05, 0x52, 0x0c, 0x73, 0x65, 0x72, 0x76, 0x65, 0x72, 0x4e, 0x75, 0x6d, 0x62, 0x65,
	0x72, 0x12, 0x2c, 0x0a, 0x09, 0x62, 0x61, 0x6c, 0x6c, 0x6f, 0x74, 0x4e, 0x75, 0x6d, 0x18, 0x03,
	0x20, 0x01, 0x28, 0x0b, 0x32, 0x0e, 0x2e, 0x63, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x2e, 0x42, 0x61,
	0x6c, 0x6c, 0x6f, 0x74, 0x52, 0x09, 0x62, 0x61, 0x6c, 0x6c, 0x6f, 0x74, 0x4e, 0x75, 0x6d, 0x12,
	0x2c, 0x0a, 0x09, 0x61, 0x63, 0x63, 0x65, 0x70, 0x74, 0x4e, 0x75, 0x6d, 0x18, 0x04, 0x20, 0x01,
	0x28, 0x0b, 0x32, 0x0e, 0x2e, 0x63, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x2e, 0x42, 0x61, 0x6c, 0x6c,
	0x6f, 0x74, 0x52, 0x09, 0x61, 0x63, 0x63, 0x65, 0x70, 0x74, 0x4e, 0x75, 0x6d, 0x12, 0x37, 0x0a,
	0x09, 0x61, 0x63, 0x63, 0x65, 0x70, 0x74, 0x56, 0x61, 0x6c, 0x18, 0x05, 0x20, 0x03, 0x28, 0x0b,
	0x32, 0x19, 0x2e, 0x63, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x2e, 0x50, 0x72, 0x6f, 0x63, 0x65, 0x73,
	0x73, 0x54, 0x78, 0x6e, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x52, 0x09, 0x61, 0x63, 0x63,
	0x65, 0x70, 0x74, 0x56, 0x61, 0x6c, 0x12, 0x35, 0x0a, 0x08, 0x6c, 0x6f, 0x63, 0x61, 0x6c, 0x56,
	0x61, 0x6c, 0x18, 0x06, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x19, 0x2e, 0x63, 0x6f, 0x6d, 0x6d, 0x6f,
	0x6e, 0x2e, 0x50, 0x72, 0x6f, 0x63, 0x65, 0x73, 0x73, 0x54, 0x78, 0x6e, 0x52, 0x65, 0x71, 0x75,
	0x65, 0x73, 0x74, 0x52, 0x08, 0x6c, 0x6f, 0x63, 0x61, 0x6c, 0x56, 0x61, 0x6c, 0x22, 0x99, 0x01,
	0x0a, 0x06, 0x41, 0x63, 0x63, 0x65, 0x70, 0x74, 0x12, 0x2c, 0x0a, 0x09, 0x62, 0x61, 0x6c, 0x6c,
	0x6f, 0x74, 0x4e, 0x75, 0x6d, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x0e, 0x2e, 0x63, 0x6f,
	0x6d, 0x6d, 0x6f, 0x6e, 0x2e, 0x42, 0x61, 0x6c, 0x6c, 0x6f, 0x74, 0x52, 0x09, 0x62, 0x61, 0x6c,
	0x6c, 0x6f, 0x74, 0x4e, 0x75, 0x6d, 0x12, 0x37, 0x0a, 0x09, 0x61, 0x63, 0x63, 0x65, 0x70, 0x74,
	0x56, 0x61, 0x6c, 0x18, 0x02, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x19, 0x2e, 0x63, 0x6f, 0x6d, 0x6d,
	0x6f, 0x6e, 0x2e, 0x50, 0x72, 0x6f, 0x63, 0x65, 0x73, 0x73, 0x54, 0x78, 0x6e, 0x52, 0x65, 0x71,
	0x75, 0x65, 0x73, 0x74, 0x52, 0x09, 0x61, 0x63, 0x63, 0x65, 0x70, 0x74, 0x56, 0x61, 0x6c, 0x12,
	0x28, 0x0a, 0x0f, 0x73, 0x65, 0x72, 0x76, 0x65, 0x72, 0x41, 0x64, 0x64, 0x72, 0x65, 0x73, 0x73,
	0x65, 0x73, 0x18, 0x03, 0x20, 0x03, 0x28, 0x09, 0x52, 0x0f, 0x73, 0x65, 0x72, 0x76, 0x65, 0x72,
	0x41, 0x64, 0x64, 0x72, 0x65, 0x73, 0x73, 0x65, 0x73, 0x22, 0x95, 0x01, 0x0a, 0x08, 0x41, 0x63,
	0x63, 0x65, 0x70, 0x74, 0x65, 0x64, 0x12, 0x2c, 0x0a, 0x09, 0x62, 0x61, 0x6c, 0x6c, 0x6f, 0x74,
	0x4e, 0x75, 0x6d, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x0e, 0x2e, 0x63, 0x6f, 0x6d, 0x6d,
	0x6f, 0x6e, 0x2e, 0x42, 0x61, 0x6c, 0x6c, 0x6f, 0x74, 0x52, 0x09, 0x62, 0x61, 0x6c, 0x6c, 0x6f,
	0x74, 0x4e, 0x75, 0x6d, 0x12, 0x37, 0x0a, 0x09, 0x61, 0x63, 0x63, 0x65, 0x70, 0x74, 0x56, 0x61,
	0x6c, 0x18, 0x02, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x19, 0x2e, 0x63, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e,
	0x2e, 0x50, 0x72, 0x6f, 0x63, 0x65, 0x73, 0x73, 0x54, 0x78, 0x6e, 0x52, 0x65, 0x71, 0x75, 0x65,
	0x73, 0x74, 0x52, 0x09, 0x61, 0x63, 0x63, 0x65, 0x70, 0x74, 0x56, 0x61, 0x6c, 0x12, 0x22, 0x0a,
	0x0c, 0x73, 0x65, 0x72, 0x76, 0x65, 0x72, 0x4e, 0x75, 0x6d, 0x62, 0x65, 0x72, 0x18, 0x03, 0x20,
	0x01, 0x28, 0x05, 0x52, 0x0c, 0x73, 0x65, 0x72, 0x76, 0x65, 0x72, 0x4e, 0x75, 0x6d, 0x62, 0x65,
	0x72, 0x22, 0x99, 0x01, 0x0a, 0x06, 0x43, 0x6f, 0x6d, 0x6d, 0x69, 0x74, 0x12, 0x2c, 0x0a, 0x09,
	0x62, 0x61, 0x6c, 0x6c, 0x6f, 0x74, 0x4e, 0x75, 0x6d, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32,
	0x0e, 0x2e, 0x63, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x2e, 0x42, 0x61, 0x6c, 0x6c, 0x6f, 0x74, 0x52,
	0x09, 0x62, 0x61, 0x6c, 0x6c, 0x6f, 0x74, 0x4e, 0x75, 0x6d, 0x12, 0x37, 0x0a, 0x09, 0x61, 0x63,
	0x63, 0x65, 0x70, 0x74, 0x56, 0x61, 0x6c, 0x18, 0x02, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x19, 0x2e,
	0x63, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x2e, 0x50, 0x72, 0x6f, 0x63, 0x65, 0x73, 0x73, 0x54, 0x78,
	0x6e, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x52, 0x09, 0x61, 0x63, 0x63, 0x65, 0x70, 0x74,
	0x56, 0x61, 0x6c, 0x12, 0x28, 0x0a, 0x0f, 0x73, 0x65, 0x72, 0x76, 0x65, 0x72, 0x41, 0x64, 0x64,
	0x72, 0x65, 0x73, 0x73, 0x65, 0x73, 0x18, 0x03, 0x20, 0x03, 0x28, 0x09, 0x52, 0x0f, 0x73, 0x65,
	0x72, 0x76, 0x65, 0x72, 0x41, 0x64, 0x64, 0x72, 0x65, 0x73, 0x73, 0x65, 0x73, 0x32, 0xca, 0x02,
	0x0a, 0x05, 0x50, 0x61, 0x78, 0x6f, 0x73, 0x12, 0x3f, 0x0a, 0x0a, 0x50, 0x72, 0x6f, 0x63, 0x65,
	0x73, 0x73, 0x54, 0x78, 0x6e, 0x12, 0x19, 0x2e, 0x63, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x2e, 0x50,
	0x72, 0x6f, 0x63, 0x65, 0x73, 0x73, 0x54, 0x78, 0x6e, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74,
	0x1a, 0x16, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62,
	0x75, 0x66, 0x2e, 0x45, 0x6d, 0x70, 0x74, 0x79, 0x12, 0x32, 0x0a, 0x07, 0x50, 0x72, 0x65, 0x70,
	0x61, 0x72, 0x65, 0x12, 0x0f, 0x2e, 0x63, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x2e, 0x50, 0x72, 0x65,
	0x70, 0x61, 0x72, 0x65, 0x1a, 0x16, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x45, 0x6d, 0x70, 0x74, 0x79, 0x12, 0x32, 0x0a, 0x07,
	0x50, 0x72, 0x6f, 0x6d, 0x69, 0x73, 0x65, 0x12, 0x0f, 0x2e, 0x63, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e,
	0x2e, 0x50, 0x72, 0x6f, 0x6d, 0x69, 0x73, 0x65, 0x1a, 0x16, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c,
	0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x45, 0x6d, 0x70, 0x74, 0x79,
	0x12, 0x30, 0x0a, 0x06, 0x41, 0x63, 0x63, 0x65, 0x70, 0x74, 0x12, 0x0e, 0x2e, 0x63, 0x6f, 0x6d,
	0x6d, 0x6f, 0x6e, 0x2e, 0x41, 0x63, 0x63, 0x65, 0x70, 0x74, 0x1a, 0x16, 0x2e, 0x67, 0x6f, 0x6f,
	0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x45, 0x6d, 0x70,
	0x74, 0x79, 0x12, 0x34, 0x0a, 0x08, 0x41, 0x63, 0x63, 0x65, 0x70, 0x74, 0x65, 0x64, 0x12, 0x10,
	0x2e, 0x63, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x2e, 0x41, 0x63, 0x63, 0x65, 0x70, 0x74, 0x65, 0x64,
	0x1a, 0x16, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62,
	0x75, 0x66, 0x2e, 0x45, 0x6d, 0x70, 0x74, 0x79, 0x12, 0x30, 0x0a, 0x06, 0x43, 0x6f, 0x6d, 0x6d,
	0x69, 0x74, 0x12, 0x0e, 0x2e, 0x63, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x2e, 0x43, 0x6f, 0x6d, 0x6d,
	0x69, 0x74, 0x1a, 0x16, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x62, 0x75, 0x66, 0x2e, 0x45, 0x6d, 0x70, 0x74, 0x79, 0x42, 0x03, 0x5a, 0x01, 0x2f, 0x62,
	0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_common_proto_rawDescOnce sync.Once
	file_common_proto_rawDescData = file_common_proto_rawDesc
)

func file_common_proto_rawDescGZIP() []byte {
	file_common_proto_rawDescOnce.Do(func() {
		file_common_proto_rawDescData = protoimpl.X.CompressGZIP(file_common_proto_rawDescData)
	})
	return file_common_proto_rawDescData
}

var file_common_proto_msgTypes = make([]protoimpl.MessageInfo, 7)
var file_common_proto_goTypes = []any{
	(*ProcessTxnRequest)(nil), // 0: common.ProcessTxnRequest
	(*Ballot)(nil),            // 1: common.Ballot
	(*Prepare)(nil),           // 2: common.Prepare
	(*Promise)(nil),           // 3: common.Promise
	(*Accept)(nil),            // 4: common.Accept
	(*Accepted)(nil),          // 5: common.Accepted
	(*Commit)(nil),            // 6: common.Commit
	(*emptypb.Empty)(nil),     // 7: google.protobuf.Empty
}
var file_common_proto_depIdxs = []int32{
	1,  // 0: common.Prepare.ballotNum:type_name -> common.Ballot
	1,  // 1: common.Promise.ballotNum:type_name -> common.Ballot
	1,  // 2: common.Promise.acceptNum:type_name -> common.Ballot
	0,  // 3: common.Promise.acceptVal:type_name -> common.ProcessTxnRequest
	0,  // 4: common.Promise.localVal:type_name -> common.ProcessTxnRequest
	1,  // 5: common.Accept.ballotNum:type_name -> common.Ballot
	0,  // 6: common.Accept.acceptVal:type_name -> common.ProcessTxnRequest
	1,  // 7: common.Accepted.ballotNum:type_name -> common.Ballot
	0,  // 8: common.Accepted.acceptVal:type_name -> common.ProcessTxnRequest
	1,  // 9: common.Commit.ballotNum:type_name -> common.Ballot
	0,  // 10: common.Commit.acceptVal:type_name -> common.ProcessTxnRequest
	0,  // 11: common.Paxos.ProcessTxn:input_type -> common.ProcessTxnRequest
	2,  // 12: common.Paxos.Prepare:input_type -> common.Prepare
	3,  // 13: common.Paxos.Promise:input_type -> common.Promise
	4,  // 14: common.Paxos.Accept:input_type -> common.Accept
	5,  // 15: common.Paxos.Accepted:input_type -> common.Accepted
	6,  // 16: common.Paxos.Commit:input_type -> common.Commit
	7,  // 17: common.Paxos.ProcessTxn:output_type -> google.protobuf.Empty
	7,  // 18: common.Paxos.Prepare:output_type -> google.protobuf.Empty
	7,  // 19: common.Paxos.Promise:output_type -> google.protobuf.Empty
	7,  // 20: common.Paxos.Accept:output_type -> google.protobuf.Empty
	7,  // 21: common.Paxos.Accepted:output_type -> google.protobuf.Empty
	7,  // 22: common.Paxos.Commit:output_type -> google.protobuf.Empty
	17, // [17:23] is the sub-list for method output_type
	11, // [11:17] is the sub-list for method input_type
	11, // [11:11] is the sub-list for extension type_name
	11, // [11:11] is the sub-list for extension extendee
	0,  // [0:11] is the sub-list for field type_name
}

func init() { file_common_proto_init() }
func file_common_proto_init() {
	if File_common_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_common_proto_msgTypes[0].Exporter = func(v any, i int) any {
			switch v := v.(*ProcessTxnRequest); i {
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
		file_common_proto_msgTypes[1].Exporter = func(v any, i int) any {
			switch v := v.(*Ballot); i {
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
		file_common_proto_msgTypes[2].Exporter = func(v any, i int) any {
			switch v := v.(*Prepare); i {
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
		file_common_proto_msgTypes[3].Exporter = func(v any, i int) any {
			switch v := v.(*Promise); i {
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
		file_common_proto_msgTypes[4].Exporter = func(v any, i int) any {
			switch v := v.(*Accept); i {
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
		file_common_proto_msgTypes[5].Exporter = func(v any, i int) any {
			switch v := v.(*Accepted); i {
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
		file_common_proto_msgTypes[6].Exporter = func(v any, i int) any {
			switch v := v.(*Commit); i {
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
			RawDescriptor: file_common_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   7,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_common_proto_goTypes,
		DependencyIndexes: file_common_proto_depIdxs,
		MessageInfos:      file_common_proto_msgTypes,
	}.Build()
	File_common_proto = out.File
	file_common_proto_rawDesc = nil
	file_common_proto_goTypes = nil
	file_common_proto_depIdxs = nil
}
