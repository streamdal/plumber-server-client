// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.28.1
// 	protoc        v3.21.6
// source: ps_args_redis_streams.proto

package args

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

type OffsetStart int32

const (
	OffsetStart_LATEST OffsetStart = 0
	OffsetStart_OLDEST OffsetStart = 1
)

// Enum value maps for OffsetStart.
var (
	OffsetStart_name = map[int32]string{
		0: "LATEST",
		1: "OLDEST",
	}
	OffsetStart_value = map[string]int32{
		"LATEST": 0,
		"OLDEST": 1,
	}
)

func (x OffsetStart) Enum() *OffsetStart {
	p := new(OffsetStart)
	*p = x
	return p
}

func (x OffsetStart) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (OffsetStart) Descriptor() protoreflect.EnumDescriptor {
	return file_ps_args_redis_streams_proto_enumTypes[0].Descriptor()
}

func (OffsetStart) Type() protoreflect.EnumType {
	return &file_ps_args_redis_streams_proto_enumTypes[0]
}

func (x OffsetStart) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use OffsetStart.Descriptor instead.
func (OffsetStart) EnumDescriptor() ([]byte, []int) {
	return file_ps_args_redis_streams_proto_rawDescGZIP(), []int{0}
}

type RedisStreamsConn struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// @gotags: kong:"help='Address of redis server',default=localhost:6379,required,env='PLUMBER_RELAY_REDIS_STREAMS_ADDRESS'"
	Address string `protobuf:"bytes,1,opt,name=address,proto3" json:"address,omitempty" kong:"help='Address of redis server',default=localhost:6379,required,env='PLUMBER_RELAY_REDIS_STREAMS_ADDRESS'"`
	// @gotags: kong:"help='Username (redis >= v6.0.0)',env='PLUMBER_RELAY_REDIS_STREAMS_USERNAME'"
	Username string `protobuf:"bytes,2,opt,name=username,proto3" json:"username,omitempty" kong:"help='Username (redis >= v6.0.0)',env='PLUMBER_RELAY_REDIS_STREAMS_USERNAME'"`
	// @gotags: kong:"help='Password (redis >= v6.0.0)',env='PLUMBER_RELAY_REDIS_STREAMS_PASSWORD'"
	Password string `protobuf:"bytes,3,opt,name=password,proto3" json:"password,omitempty" kong:"help='Password (redis >= v6.0.0)',env='PLUMBER_RELAY_REDIS_STREAMS_PASSWORD'"`
	// @gotags: kong:"help='Database (0-16)',env='PLUMBER_RELAY_REDIS_PUBSUB_DATABASE'"
	Database uint32 `protobuf:"varint,4,opt,name=database,proto3" json:"database,omitempty" kong:"help='Database (0-16)',env='PLUMBER_RELAY_REDIS_PUBSUB_DATABASE'"`
}

func (x *RedisStreamsConn) Reset() {
	*x = RedisStreamsConn{}
	if protoimpl.UnsafeEnabled {
		mi := &file_ps_args_redis_streams_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *RedisStreamsConn) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*RedisStreamsConn) ProtoMessage() {}

func (x *RedisStreamsConn) ProtoReflect() protoreflect.Message {
	mi := &file_ps_args_redis_streams_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use RedisStreamsConn.ProtoReflect.Descriptor instead.
func (*RedisStreamsConn) Descriptor() ([]byte, []int) {
	return file_ps_args_redis_streams_proto_rawDescGZIP(), []int{0}
}

func (x *RedisStreamsConn) GetAddress() string {
	if x != nil {
		return x.Address
	}
	return ""
}

func (x *RedisStreamsConn) GetUsername() string {
	if x != nil {
		return x.Username
	}
	return ""
}

func (x *RedisStreamsConn) GetPassword() string {
	if x != nil {
		return x.Password
	}
	return ""
}

func (x *RedisStreamsConn) GetDatabase() uint32 {
	if x != nil {
		return x.Database
	}
	return 0
}

type CreateConsumerConfig struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// @gotags: kong:"help='Create the streams if creating a new consumer group',env='PLUMBER_RELAY_REDIS_STREAMS_CREATE_STREAMS'"
	CreateStreams bool `protobuf:"varint,1,opt,name=create_streams,json=createStreams,proto3" json:"create_streams,omitempty" kong:"help='Create the streams if creating a new consumer group',env='PLUMBER_RELAY_REDIS_STREAMS_CREATE_STREAMS'"`
	// @gotags: kong:"help='Recreate this consumer group if it does not exist',env='PLUMBER_RELAY_REDIS_STREAMS_RECREATE_CONSUMER_GROUP'"
	RecreateConsumerGroup bool `protobuf:"varint,2,opt,name=recreate_consumer_group,json=recreateConsumerGroup,proto3" json:"recreate_consumer_group,omitempty" kong:"help='Recreate this consumer group if it does not exist',env='PLUMBER_RELAY_REDIS_STREAMS_RECREATE_CONSUMER_GROUP'"`
	// @gotags: kong:"help='What offset to start reading at (options: latest oldest)',default=latest,required,env='PLUMBER_RELAY_REDIS_STREAMS_START_ID',type=pbenum,pbenum_lowercase"
	OffsetStart OffsetStart `protobuf:"varint,3,opt,name=offset_start,json=offsetStart,proto3,enum=protos.args.OffsetStart" json:"offset_start,omitempty" kong:"help='What offset to start reading at (options: latest oldest)',default=latest,required,env='PLUMBER_RELAY_REDIS_STREAMS_START_ID',type=pbenum,pbenum_lowercase"`
}

func (x *CreateConsumerConfig) Reset() {
	*x = CreateConsumerConfig{}
	if protoimpl.UnsafeEnabled {
		mi := &file_ps_args_redis_streams_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *CreateConsumerConfig) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CreateConsumerConfig) ProtoMessage() {}

func (x *CreateConsumerConfig) ProtoReflect() protoreflect.Message {
	mi := &file_ps_args_redis_streams_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CreateConsumerConfig.ProtoReflect.Descriptor instead.
func (*CreateConsumerConfig) Descriptor() ([]byte, []int) {
	return file_ps_args_redis_streams_proto_rawDescGZIP(), []int{1}
}

func (x *CreateConsumerConfig) GetCreateStreams() bool {
	if x != nil {
		return x.CreateStreams
	}
	return false
}

func (x *CreateConsumerConfig) GetRecreateConsumerGroup() bool {
	if x != nil {
		return x.RecreateConsumerGroup
	}
	return false
}

func (x *CreateConsumerConfig) GetOffsetStart() OffsetStart {
	if x != nil {
		return x.OffsetStart
	}
	return OffsetStart_LATEST
}

type RedisStreamsReadArgs struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// @gotags: kong:"help='Streams to read from',required,env='PLUMBER_RELAY_REDIS_STREAMS_STREAMS'"
	Streams []string `protobuf:"bytes,1,rep,name=streams,proto3" json:"streams,omitempty" kong:"help='Streams to read from',required,env='PLUMBER_RELAY_REDIS_STREAMS_STREAMS'"`
	// @gotags: kong:"help='Consumer group name',env='PLUMBER_RELAY_REDIS_STREAMS_CONSUMER_GROUP',default=plumber"
	ConsumerGroup string `protobuf:"bytes,2,opt,name=consumer_group,json=consumerGroup,proto3" json:"consumer_group,omitempty" kong:"help='Consumer group name',env='PLUMBER_RELAY_REDIS_STREAMS_CONSUMER_GROUP',default=plumber"`
	// @gotags: kong:"help='Consumer name',env='PLUMBER_RELAY_REDIS_STREAMS_CONSUMER_NAME',default=plumber-consumer-1"
	ConsumerName string `protobuf:"bytes,3,opt,name=consumer_name,json=consumerName,proto3" json:"consumer_name,omitempty" kong:"help='Consumer name',env='PLUMBER_RELAY_REDIS_STREAMS_CONSUMER_NAME',default=plumber-consumer-1"`
	// @gotags: kong:"help='Number of records to read from stream(s) per read',env='PLUMBER_RELAY_REDIS_STREAMS_COUNT',default=10"
	Count uint32 `protobuf:"varint,4,opt,name=count,proto3" json:"count,omitempty" kong:"help='Number of records to read from stream(s) per read',env='PLUMBER_RELAY_REDIS_STREAMS_COUNT',default=10"`
	// @gotags: kong:"embed"
	CreateConsumerConfig *CreateConsumerConfig `protobuf:"bytes,5,opt,name=create_consumer_config,json=createConsumerConfig,proto3" json:"create_consumer_config,omitempty" kong:"embed"`
}

func (x *RedisStreamsReadArgs) Reset() {
	*x = RedisStreamsReadArgs{}
	if protoimpl.UnsafeEnabled {
		mi := &file_ps_args_redis_streams_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *RedisStreamsReadArgs) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*RedisStreamsReadArgs) ProtoMessage() {}

func (x *RedisStreamsReadArgs) ProtoReflect() protoreflect.Message {
	mi := &file_ps_args_redis_streams_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use RedisStreamsReadArgs.ProtoReflect.Descriptor instead.
func (*RedisStreamsReadArgs) Descriptor() ([]byte, []int) {
	return file_ps_args_redis_streams_proto_rawDescGZIP(), []int{2}
}

func (x *RedisStreamsReadArgs) GetStreams() []string {
	if x != nil {
		return x.Streams
	}
	return nil
}

func (x *RedisStreamsReadArgs) GetConsumerGroup() string {
	if x != nil {
		return x.ConsumerGroup
	}
	return ""
}

func (x *RedisStreamsReadArgs) GetConsumerName() string {
	if x != nil {
		return x.ConsumerName
	}
	return ""
}

func (x *RedisStreamsReadArgs) GetCount() uint32 {
	if x != nil {
		return x.Count
	}
	return 0
}

func (x *RedisStreamsReadArgs) GetCreateConsumerConfig() *CreateConsumerConfig {
	if x != nil {
		return x.CreateConsumerConfig
	}
	return nil
}

type RedisStreamsWriteArgs struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// @gotags: kong:"help='What redis ID to use for input data (* = auto-generate)',default='*'"
	WriteId string `protobuf:"bytes,1,opt,name=write_id,json=writeId,proto3" json:"write_id,omitempty" kong:"help='What redis ID to use for input data (* = auto-generate)',default='*'"`
	// @gotags: kong:"help='Streams to write to'"
	Streams []string `protobuf:"bytes,2,rep,name=streams,proto3" json:"streams,omitempty" kong:"help='Streams to write to'"`
	// @gotags: kong:"help='Key name to write input data to'"
	Key string `protobuf:"bytes,3,opt,name=key,proto3" json:"key,omitempty" kong:"help='Key name to write input data to'"`
}

func (x *RedisStreamsWriteArgs) Reset() {
	*x = RedisStreamsWriteArgs{}
	if protoimpl.UnsafeEnabled {
		mi := &file_ps_args_redis_streams_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *RedisStreamsWriteArgs) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*RedisStreamsWriteArgs) ProtoMessage() {}

func (x *RedisStreamsWriteArgs) ProtoReflect() protoreflect.Message {
	mi := &file_ps_args_redis_streams_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use RedisStreamsWriteArgs.ProtoReflect.Descriptor instead.
func (*RedisStreamsWriteArgs) Descriptor() ([]byte, []int) {
	return file_ps_args_redis_streams_proto_rawDescGZIP(), []int{3}
}

func (x *RedisStreamsWriteArgs) GetWriteId() string {
	if x != nil {
		return x.WriteId
	}
	return ""
}

func (x *RedisStreamsWriteArgs) GetStreams() []string {
	if x != nil {
		return x.Streams
	}
	return nil
}

func (x *RedisStreamsWriteArgs) GetKey() string {
	if x != nil {
		return x.Key
	}
	return ""
}

var File_ps_args_redis_streams_proto protoreflect.FileDescriptor

var file_ps_args_redis_streams_proto_rawDesc = []byte{
	0x0a, 0x1b, 0x70, 0x73, 0x5f, 0x61, 0x72, 0x67, 0x73, 0x5f, 0x72, 0x65, 0x64, 0x69, 0x73, 0x5f,
	0x73, 0x74, 0x72, 0x65, 0x61, 0x6d, 0x73, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x0b, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x73, 0x2e, 0x61, 0x72, 0x67, 0x73, 0x22, 0x80, 0x01, 0x0a, 0x10, 0x52,
	0x65, 0x64, 0x69, 0x73, 0x53, 0x74, 0x72, 0x65, 0x61, 0x6d, 0x73, 0x43, 0x6f, 0x6e, 0x6e, 0x12,
	0x18, 0x0a, 0x07, 0x61, 0x64, 0x64, 0x72, 0x65, 0x73, 0x73, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09,
	0x52, 0x07, 0x61, 0x64, 0x64, 0x72, 0x65, 0x73, 0x73, 0x12, 0x1a, 0x0a, 0x08, 0x75, 0x73, 0x65,
	0x72, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x75, 0x73, 0x65,
	0x72, 0x6e, 0x61, 0x6d, 0x65, 0x12, 0x1a, 0x0a, 0x08, 0x70, 0x61, 0x73, 0x73, 0x77, 0x6f, 0x72,
	0x64, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x70, 0x61, 0x73, 0x73, 0x77, 0x6f, 0x72,
	0x64, 0x12, 0x1a, 0x0a, 0x08, 0x64, 0x61, 0x74, 0x61, 0x62, 0x61, 0x73, 0x65, 0x18, 0x04, 0x20,
	0x01, 0x28, 0x0d, 0x52, 0x08, 0x64, 0x61, 0x74, 0x61, 0x62, 0x61, 0x73, 0x65, 0x22, 0xb2, 0x01,
	0x0a, 0x14, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x43, 0x6f, 0x6e, 0x73, 0x75, 0x6d, 0x65, 0x72,
	0x43, 0x6f, 0x6e, 0x66, 0x69, 0x67, 0x12, 0x25, 0x0a, 0x0e, 0x63, 0x72, 0x65, 0x61, 0x74, 0x65,
	0x5f, 0x73, 0x74, 0x72, 0x65, 0x61, 0x6d, 0x73, 0x18, 0x01, 0x20, 0x01, 0x28, 0x08, 0x52, 0x0d,
	0x63, 0x72, 0x65, 0x61, 0x74, 0x65, 0x53, 0x74, 0x72, 0x65, 0x61, 0x6d, 0x73, 0x12, 0x36, 0x0a,
	0x17, 0x72, 0x65, 0x63, 0x72, 0x65, 0x61, 0x74, 0x65, 0x5f, 0x63, 0x6f, 0x6e, 0x73, 0x75, 0x6d,
	0x65, 0x72, 0x5f, 0x67, 0x72, 0x6f, 0x75, 0x70, 0x18, 0x02, 0x20, 0x01, 0x28, 0x08, 0x52, 0x15,
	0x72, 0x65, 0x63, 0x72, 0x65, 0x61, 0x74, 0x65, 0x43, 0x6f, 0x6e, 0x73, 0x75, 0x6d, 0x65, 0x72,
	0x47, 0x72, 0x6f, 0x75, 0x70, 0x12, 0x3b, 0x0a, 0x0c, 0x6f, 0x66, 0x66, 0x73, 0x65, 0x74, 0x5f,
	0x73, 0x74, 0x61, 0x72, 0x74, 0x18, 0x03, 0x20, 0x01, 0x28, 0x0e, 0x32, 0x18, 0x2e, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x73, 0x2e, 0x61, 0x72, 0x67, 0x73, 0x2e, 0x4f, 0x66, 0x66, 0x73, 0x65, 0x74,
	0x53, 0x74, 0x61, 0x72, 0x74, 0x52, 0x0b, 0x6f, 0x66, 0x66, 0x73, 0x65, 0x74, 0x53, 0x74, 0x61,
	0x72, 0x74, 0x22, 0xeb, 0x01, 0x0a, 0x14, 0x52, 0x65, 0x64, 0x69, 0x73, 0x53, 0x74, 0x72, 0x65,
	0x61, 0x6d, 0x73, 0x52, 0x65, 0x61, 0x64, 0x41, 0x72, 0x67, 0x73, 0x12, 0x18, 0x0a, 0x07, 0x73,
	0x74, 0x72, 0x65, 0x61, 0x6d, 0x73, 0x18, 0x01, 0x20, 0x03, 0x28, 0x09, 0x52, 0x07, 0x73, 0x74,
	0x72, 0x65, 0x61, 0x6d, 0x73, 0x12, 0x25, 0x0a, 0x0e, 0x63, 0x6f, 0x6e, 0x73, 0x75, 0x6d, 0x65,
	0x72, 0x5f, 0x67, 0x72, 0x6f, 0x75, 0x70, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0d, 0x63,
	0x6f, 0x6e, 0x73, 0x75, 0x6d, 0x65, 0x72, 0x47, 0x72, 0x6f, 0x75, 0x70, 0x12, 0x23, 0x0a, 0x0d,
	0x63, 0x6f, 0x6e, 0x73, 0x75, 0x6d, 0x65, 0x72, 0x5f, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x03, 0x20,
	0x01, 0x28, 0x09, 0x52, 0x0c, 0x63, 0x6f, 0x6e, 0x73, 0x75, 0x6d, 0x65, 0x72, 0x4e, 0x61, 0x6d,
	0x65, 0x12, 0x14, 0x0a, 0x05, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x18, 0x04, 0x20, 0x01, 0x28, 0x0d,
	0x52, 0x05, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x12, 0x57, 0x0a, 0x16, 0x63, 0x72, 0x65, 0x61, 0x74,
	0x65, 0x5f, 0x63, 0x6f, 0x6e, 0x73, 0x75, 0x6d, 0x65, 0x72, 0x5f, 0x63, 0x6f, 0x6e, 0x66, 0x69,
	0x67, 0x18, 0x05, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x21, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x73,
	0x2e, 0x61, 0x72, 0x67, 0x73, 0x2e, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x43, 0x6f, 0x6e, 0x73,
	0x75, 0x6d, 0x65, 0x72, 0x43, 0x6f, 0x6e, 0x66, 0x69, 0x67, 0x52, 0x14, 0x63, 0x72, 0x65, 0x61,
	0x74, 0x65, 0x43, 0x6f, 0x6e, 0x73, 0x75, 0x6d, 0x65, 0x72, 0x43, 0x6f, 0x6e, 0x66, 0x69, 0x67,
	0x22, 0x5e, 0x0a, 0x15, 0x52, 0x65, 0x64, 0x69, 0x73, 0x53, 0x74, 0x72, 0x65, 0x61, 0x6d, 0x73,
	0x57, 0x72, 0x69, 0x74, 0x65, 0x41, 0x72, 0x67, 0x73, 0x12, 0x19, 0x0a, 0x08, 0x77, 0x72, 0x69,
	0x74, 0x65, 0x5f, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x07, 0x77, 0x72, 0x69,
	0x74, 0x65, 0x49, 0x64, 0x12, 0x18, 0x0a, 0x07, 0x73, 0x74, 0x72, 0x65, 0x61, 0x6d, 0x73, 0x18,
	0x02, 0x20, 0x03, 0x28, 0x09, 0x52, 0x07, 0x73, 0x74, 0x72, 0x65, 0x61, 0x6d, 0x73, 0x12, 0x10,
	0x0a, 0x03, 0x6b, 0x65, 0x79, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x03, 0x6b, 0x65, 0x79,
	0x2a, 0x25, 0x0a, 0x0b, 0x4f, 0x66, 0x66, 0x73, 0x65, 0x74, 0x53, 0x74, 0x61, 0x72, 0x74, 0x12,
	0x0a, 0x0a, 0x06, 0x4c, 0x41, 0x54, 0x45, 0x53, 0x54, 0x10, 0x00, 0x12, 0x0a, 0x0a, 0x06, 0x4f,
	0x4c, 0x44, 0x45, 0x53, 0x54, 0x10, 0x01, 0x42, 0x3b, 0x5a, 0x39, 0x67, 0x69, 0x74, 0x68, 0x75,
	0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x62, 0x61, 0x74, 0x63, 0x68, 0x63, 0x6f, 0x72, 0x70, 0x2f,
	0x70, 0x6c, 0x75, 0x6d, 0x62, 0x65, 0x72, 0x2d, 0x73, 0x63, 0x68, 0x65, 0x6d, 0x61, 0x73, 0x2f,
	0x62, 0x75, 0x69, 0x6c, 0x64, 0x2f, 0x67, 0x6f, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x73, 0x2f,
	0x61, 0x72, 0x67, 0x73, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_ps_args_redis_streams_proto_rawDescOnce sync.Once
	file_ps_args_redis_streams_proto_rawDescData = file_ps_args_redis_streams_proto_rawDesc
)

func file_ps_args_redis_streams_proto_rawDescGZIP() []byte {
	file_ps_args_redis_streams_proto_rawDescOnce.Do(func() {
		file_ps_args_redis_streams_proto_rawDescData = protoimpl.X.CompressGZIP(file_ps_args_redis_streams_proto_rawDescData)
	})
	return file_ps_args_redis_streams_proto_rawDescData
}

var file_ps_args_redis_streams_proto_enumTypes = make([]protoimpl.EnumInfo, 1)
var file_ps_args_redis_streams_proto_msgTypes = make([]protoimpl.MessageInfo, 4)
var file_ps_args_redis_streams_proto_goTypes = []interface{}{
	(OffsetStart)(0),              // 0: protos.args.OffsetStart
	(*RedisStreamsConn)(nil),      // 1: protos.args.RedisStreamsConn
	(*CreateConsumerConfig)(nil),  // 2: protos.args.CreateConsumerConfig
	(*RedisStreamsReadArgs)(nil),  // 3: protos.args.RedisStreamsReadArgs
	(*RedisStreamsWriteArgs)(nil), // 4: protos.args.RedisStreamsWriteArgs
}
var file_ps_args_redis_streams_proto_depIdxs = []int32{
	0, // 0: protos.args.CreateConsumerConfig.offset_start:type_name -> protos.args.OffsetStart
	2, // 1: protos.args.RedisStreamsReadArgs.create_consumer_config:type_name -> protos.args.CreateConsumerConfig
	2, // [2:2] is the sub-list for method output_type
	2, // [2:2] is the sub-list for method input_type
	2, // [2:2] is the sub-list for extension type_name
	2, // [2:2] is the sub-list for extension extendee
	0, // [0:2] is the sub-list for field type_name
}

func init() { file_ps_args_redis_streams_proto_init() }
func file_ps_args_redis_streams_proto_init() {
	if File_ps_args_redis_streams_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_ps_args_redis_streams_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*RedisStreamsConn); i {
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
		file_ps_args_redis_streams_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*CreateConsumerConfig); i {
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
		file_ps_args_redis_streams_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*RedisStreamsReadArgs); i {
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
		file_ps_args_redis_streams_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*RedisStreamsWriteArgs); i {
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
			RawDescriptor: file_ps_args_redis_streams_proto_rawDesc,
			NumEnums:      1,
			NumMessages:   4,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_ps_args_redis_streams_proto_goTypes,
		DependencyIndexes: file_ps_args_redis_streams_proto_depIdxs,
		EnumInfos:         file_ps_args_redis_streams_proto_enumTypes,
		MessageInfos:      file_ps_args_redis_streams_proto_msgTypes,
	}.Build()
	File_ps_args_redis_streams_proto = out.File
	file_ps_args_redis_streams_proto_rawDesc = nil
	file_ps_args_redis_streams_proto_goTypes = nil
	file_ps_args_redis_streams_proto_depIdxs = nil
}
