// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.28.1
// 	protoc        v3.21.7
// source: envoy/service/ratelimit/v2/rls.proto

package ratelimitv2

import (
	context "context"
	_ "github.com/cncf/xds/go/udpa/annotations"
	core "github.com/emissary-ingress/emissary/v3/pkg/api/envoy/api/v2/core"
	ratelimit "github.com/emissary-ingress/emissary/v3/pkg/api/envoy/api/v2/ratelimit"
	grpc "google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
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

type RateLimitResponse_Code int32

const (
	// The response code is not known.
	RateLimitResponse_UNKNOWN RateLimitResponse_Code = 0
	// The response code to notify that the number of requests are under limit.
	RateLimitResponse_OK RateLimitResponse_Code = 1
	// The response code to notify that the number of requests are over limit.
	RateLimitResponse_OVER_LIMIT RateLimitResponse_Code = 2
)

// Enum value maps for RateLimitResponse_Code.
var (
	RateLimitResponse_Code_name = map[int32]string{
		0: "UNKNOWN",
		1: "OK",
		2: "OVER_LIMIT",
	}
	RateLimitResponse_Code_value = map[string]int32{
		"UNKNOWN":    0,
		"OK":         1,
		"OVER_LIMIT": 2,
	}
)

func (x RateLimitResponse_Code) Enum() *RateLimitResponse_Code {
	p := new(RateLimitResponse_Code)
	*p = x
	return p
}

func (x RateLimitResponse_Code) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (RateLimitResponse_Code) Descriptor() protoreflect.EnumDescriptor {
	return file_envoy_service_ratelimit_v2_rls_proto_enumTypes[0].Descriptor()
}

func (RateLimitResponse_Code) Type() protoreflect.EnumType {
	return &file_envoy_service_ratelimit_v2_rls_proto_enumTypes[0]
}

func (x RateLimitResponse_Code) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use RateLimitResponse_Code.Descriptor instead.
func (RateLimitResponse_Code) EnumDescriptor() ([]byte, []int) {
	return file_envoy_service_ratelimit_v2_rls_proto_rawDescGZIP(), []int{1, 0}
}

type RateLimitResponse_RateLimit_Unit int32

const (
	// The time unit is not known.
	RateLimitResponse_RateLimit_UNKNOWN RateLimitResponse_RateLimit_Unit = 0
	// The time unit representing a second.
	RateLimitResponse_RateLimit_SECOND RateLimitResponse_RateLimit_Unit = 1
	// The time unit representing a minute.
	RateLimitResponse_RateLimit_MINUTE RateLimitResponse_RateLimit_Unit = 2
	// The time unit representing an hour.
	RateLimitResponse_RateLimit_HOUR RateLimitResponse_RateLimit_Unit = 3
	// The time unit representing a day.
	RateLimitResponse_RateLimit_DAY RateLimitResponse_RateLimit_Unit = 4
)

// Enum value maps for RateLimitResponse_RateLimit_Unit.
var (
	RateLimitResponse_RateLimit_Unit_name = map[int32]string{
		0: "UNKNOWN",
		1: "SECOND",
		2: "MINUTE",
		3: "HOUR",
		4: "DAY",
	}
	RateLimitResponse_RateLimit_Unit_value = map[string]int32{
		"UNKNOWN": 0,
		"SECOND":  1,
		"MINUTE":  2,
		"HOUR":    3,
		"DAY":     4,
	}
)

func (x RateLimitResponse_RateLimit_Unit) Enum() *RateLimitResponse_RateLimit_Unit {
	p := new(RateLimitResponse_RateLimit_Unit)
	*p = x
	return p
}

func (x RateLimitResponse_RateLimit_Unit) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (RateLimitResponse_RateLimit_Unit) Descriptor() protoreflect.EnumDescriptor {
	return file_envoy_service_ratelimit_v2_rls_proto_enumTypes[1].Descriptor()
}

func (RateLimitResponse_RateLimit_Unit) Type() protoreflect.EnumType {
	return &file_envoy_service_ratelimit_v2_rls_proto_enumTypes[1]
}

func (x RateLimitResponse_RateLimit_Unit) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use RateLimitResponse_RateLimit_Unit.Descriptor instead.
func (RateLimitResponse_RateLimit_Unit) EnumDescriptor() ([]byte, []int) {
	return file_envoy_service_ratelimit_v2_rls_proto_rawDescGZIP(), []int{1, 0, 0}
}

// Main message for a rate limit request. The rate limit service is designed to be fully generic
// in the sense that it can operate on arbitrary hierarchical key/value pairs. The loaded
// configuration will parse the request and find the most specific limit to apply. In addition,
// a RateLimitRequest can contain multiple "descriptors" to limit on. When multiple descriptors
// are provided, the server will limit on *ALL* of them and return an OVER_LIMIT response if any
// of them are over limit. This enables more complex application level rate limiting scenarios
// if desired.
type RateLimitRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// All rate limit requests must specify a domain. This enables the configuration to be per
	// application without fear of overlap. E.g., "envoy".
	Domain string `protobuf:"bytes,1,opt,name=domain,proto3" json:"domain,omitempty"`
	// All rate limit requests must specify at least one RateLimitDescriptor. Each descriptor is
	// processed by the service (see below). If any of the descriptors are over limit, the entire
	// request is considered to be over limit.
	Descriptors []*ratelimit.RateLimitDescriptor `protobuf:"bytes,2,rep,name=descriptors,proto3" json:"descriptors,omitempty"`
	// Rate limit requests can optionally specify the number of hits a request adds to the matched
	// limit. If the value is not set in the message, a request increases the matched limit by 1.
	HitsAddend uint32 `protobuf:"varint,3,opt,name=hits_addend,json=hitsAddend,proto3" json:"hits_addend,omitempty"`
}

func (x *RateLimitRequest) Reset() {
	*x = RateLimitRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_envoy_service_ratelimit_v2_rls_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *RateLimitRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*RateLimitRequest) ProtoMessage() {}

func (x *RateLimitRequest) ProtoReflect() protoreflect.Message {
	mi := &file_envoy_service_ratelimit_v2_rls_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use RateLimitRequest.ProtoReflect.Descriptor instead.
func (*RateLimitRequest) Descriptor() ([]byte, []int) {
	return file_envoy_service_ratelimit_v2_rls_proto_rawDescGZIP(), []int{0}
}

func (x *RateLimitRequest) GetDomain() string {
	if x != nil {
		return x.Domain
	}
	return ""
}

func (x *RateLimitRequest) GetDescriptors() []*ratelimit.RateLimitDescriptor {
	if x != nil {
		return x.Descriptors
	}
	return nil
}

func (x *RateLimitRequest) GetHitsAddend() uint32 {
	if x != nil {
		return x.HitsAddend
	}
	return 0
}

// A response from a ShouldRateLimit call.
type RateLimitResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// The overall response code which takes into account all of the descriptors that were passed
	// in the RateLimitRequest message.
	OverallCode RateLimitResponse_Code `protobuf:"varint,1,opt,name=overall_code,json=overallCode,proto3,enum=envoy.service.ratelimit.v2.RateLimitResponse_Code" json:"overall_code,omitempty"`
	// A list of DescriptorStatus messages which matches the length of the descriptor list passed
	// in the RateLimitRequest. This can be used by the caller to determine which individual
	// descriptors failed and/or what the currently configured limits are for all of them.
	Statuses []*RateLimitResponse_DescriptorStatus `protobuf:"bytes,2,rep,name=statuses,proto3" json:"statuses,omitempty"`
	// A list of headers to add to the response
	Headers []*core.HeaderValue `protobuf:"bytes,3,rep,name=headers,proto3" json:"headers,omitempty"`
	// A list of headers to add to the request when forwarded
	RequestHeadersToAdd []*core.HeaderValue `protobuf:"bytes,4,rep,name=request_headers_to_add,json=requestHeadersToAdd,proto3" json:"request_headers_to_add,omitempty"`
}

func (x *RateLimitResponse) Reset() {
	*x = RateLimitResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_envoy_service_ratelimit_v2_rls_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *RateLimitResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*RateLimitResponse) ProtoMessage() {}

func (x *RateLimitResponse) ProtoReflect() protoreflect.Message {
	mi := &file_envoy_service_ratelimit_v2_rls_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use RateLimitResponse.ProtoReflect.Descriptor instead.
func (*RateLimitResponse) Descriptor() ([]byte, []int) {
	return file_envoy_service_ratelimit_v2_rls_proto_rawDescGZIP(), []int{1}
}

func (x *RateLimitResponse) GetOverallCode() RateLimitResponse_Code {
	if x != nil {
		return x.OverallCode
	}
	return RateLimitResponse_UNKNOWN
}

func (x *RateLimitResponse) GetStatuses() []*RateLimitResponse_DescriptorStatus {
	if x != nil {
		return x.Statuses
	}
	return nil
}

func (x *RateLimitResponse) GetHeaders() []*core.HeaderValue {
	if x != nil {
		return x.Headers
	}
	return nil
}

func (x *RateLimitResponse) GetRequestHeadersToAdd() []*core.HeaderValue {
	if x != nil {
		return x.RequestHeadersToAdd
	}
	return nil
}

// Defines an actual rate limit in terms of requests per unit of time and the unit itself.
type RateLimitResponse_RateLimit struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// A name or description of this limit.
	Name string `protobuf:"bytes,3,opt,name=name,proto3" json:"name,omitempty"`
	// The number of requests per unit of time.
	RequestsPerUnit uint32 `protobuf:"varint,1,opt,name=requests_per_unit,json=requestsPerUnit,proto3" json:"requests_per_unit,omitempty"`
	// The unit of time.
	Unit RateLimitResponse_RateLimit_Unit `protobuf:"varint,2,opt,name=unit,proto3,enum=envoy.service.ratelimit.v2.RateLimitResponse_RateLimit_Unit" json:"unit,omitempty"`
}

func (x *RateLimitResponse_RateLimit) Reset() {
	*x = RateLimitResponse_RateLimit{}
	if protoimpl.UnsafeEnabled {
		mi := &file_envoy_service_ratelimit_v2_rls_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *RateLimitResponse_RateLimit) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*RateLimitResponse_RateLimit) ProtoMessage() {}

func (x *RateLimitResponse_RateLimit) ProtoReflect() protoreflect.Message {
	mi := &file_envoy_service_ratelimit_v2_rls_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use RateLimitResponse_RateLimit.ProtoReflect.Descriptor instead.
func (*RateLimitResponse_RateLimit) Descriptor() ([]byte, []int) {
	return file_envoy_service_ratelimit_v2_rls_proto_rawDescGZIP(), []int{1, 0}
}

func (x *RateLimitResponse_RateLimit) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

func (x *RateLimitResponse_RateLimit) GetRequestsPerUnit() uint32 {
	if x != nil {
		return x.RequestsPerUnit
	}
	return 0
}

func (x *RateLimitResponse_RateLimit) GetUnit() RateLimitResponse_RateLimit_Unit {
	if x != nil {
		return x.Unit
	}
	return RateLimitResponse_RateLimit_UNKNOWN
}

type RateLimitResponse_DescriptorStatus struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// The response code for an individual descriptor.
	Code RateLimitResponse_Code `protobuf:"varint,1,opt,name=code,proto3,enum=envoy.service.ratelimit.v2.RateLimitResponse_Code" json:"code,omitempty"`
	// The current limit as configured by the server. Useful for debugging, etc.
	CurrentLimit *RateLimitResponse_RateLimit `protobuf:"bytes,2,opt,name=current_limit,json=currentLimit,proto3" json:"current_limit,omitempty"`
	// The limit remaining in the current time unit.
	LimitRemaining uint32 `protobuf:"varint,3,opt,name=limit_remaining,json=limitRemaining,proto3" json:"limit_remaining,omitempty"`
}

func (x *RateLimitResponse_DescriptorStatus) Reset() {
	*x = RateLimitResponse_DescriptorStatus{}
	if protoimpl.UnsafeEnabled {
		mi := &file_envoy_service_ratelimit_v2_rls_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *RateLimitResponse_DescriptorStatus) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*RateLimitResponse_DescriptorStatus) ProtoMessage() {}

func (x *RateLimitResponse_DescriptorStatus) ProtoReflect() protoreflect.Message {
	mi := &file_envoy_service_ratelimit_v2_rls_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use RateLimitResponse_DescriptorStatus.ProtoReflect.Descriptor instead.
func (*RateLimitResponse_DescriptorStatus) Descriptor() ([]byte, []int) {
	return file_envoy_service_ratelimit_v2_rls_proto_rawDescGZIP(), []int{1, 1}
}

func (x *RateLimitResponse_DescriptorStatus) GetCode() RateLimitResponse_Code {
	if x != nil {
		return x.Code
	}
	return RateLimitResponse_UNKNOWN
}

func (x *RateLimitResponse_DescriptorStatus) GetCurrentLimit() *RateLimitResponse_RateLimit {
	if x != nil {
		return x.CurrentLimit
	}
	return nil
}

func (x *RateLimitResponse_DescriptorStatus) GetLimitRemaining() uint32 {
	if x != nil {
		return x.LimitRemaining
	}
	return 0
}

var File_envoy_service_ratelimit_v2_rls_proto protoreflect.FileDescriptor

var file_envoy_service_ratelimit_v2_rls_proto_rawDesc = []byte{
	0x0a, 0x24, 0x65, 0x6e, 0x76, 0x6f, 0x79, 0x2f, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x2f,
	0x72, 0x61, 0x74, 0x65, 0x6c, 0x69, 0x6d, 0x69, 0x74, 0x2f, 0x76, 0x32, 0x2f, 0x72, 0x6c, 0x73,
	0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x1a, 0x65, 0x6e, 0x76, 0x6f, 0x79, 0x2e, 0x73, 0x65,
	0x72, 0x76, 0x69, 0x63, 0x65, 0x2e, 0x72, 0x61, 0x74, 0x65, 0x6c, 0x69, 0x6d, 0x69, 0x74, 0x2e,
	0x76, 0x32, 0x1a, 0x1c, 0x65, 0x6e, 0x76, 0x6f, 0x79, 0x2f, 0x61, 0x70, 0x69, 0x2f, 0x76, 0x32,
	0x2f, 0x63, 0x6f, 0x72, 0x65, 0x2f, 0x62, 0x61, 0x73, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x1a, 0x26, 0x65, 0x6e, 0x76, 0x6f, 0x79, 0x2f, 0x61, 0x70, 0x69, 0x2f, 0x76, 0x32, 0x2f, 0x72,
	0x61, 0x74, 0x65, 0x6c, 0x69, 0x6d, 0x69, 0x74, 0x2f, 0x72, 0x61, 0x74, 0x65, 0x6c, 0x69, 0x6d,
	0x69, 0x74, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x1e, 0x75, 0x64, 0x70, 0x61, 0x2f, 0x61,
	0x6e, 0x6e, 0x6f, 0x74, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x2f, 0x6d, 0x69, 0x67, 0x72, 0x61,
	0x74, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x1d, 0x75, 0x64, 0x70, 0x61, 0x2f, 0x61,
	0x6e, 0x6e, 0x6f, 0x74, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x2f, 0x73, 0x74, 0x61, 0x74, 0x75,
	0x73, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0x9a, 0x01, 0x0a, 0x10, 0x52, 0x61, 0x74, 0x65,
	0x4c, 0x69, 0x6d, 0x69, 0x74, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x16, 0x0a, 0x06,
	0x64, 0x6f, 0x6d, 0x61, 0x69, 0x6e, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x06, 0x64, 0x6f,
	0x6d, 0x61, 0x69, 0x6e, 0x12, 0x4d, 0x0a, 0x0b, 0x64, 0x65, 0x73, 0x63, 0x72, 0x69, 0x70, 0x74,
	0x6f, 0x72, 0x73, 0x18, 0x02, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x2b, 0x2e, 0x65, 0x6e, 0x76, 0x6f,
	0x79, 0x2e, 0x61, 0x70, 0x69, 0x2e, 0x76, 0x32, 0x2e, 0x72, 0x61, 0x74, 0x65, 0x6c, 0x69, 0x6d,
	0x69, 0x74, 0x2e, 0x52, 0x61, 0x74, 0x65, 0x4c, 0x69, 0x6d, 0x69, 0x74, 0x44, 0x65, 0x73, 0x63,
	0x72, 0x69, 0x70, 0x74, 0x6f, 0x72, 0x52, 0x0b, 0x64, 0x65, 0x73, 0x63, 0x72, 0x69, 0x70, 0x74,
	0x6f, 0x72, 0x73, 0x12, 0x1f, 0x0a, 0x0b, 0x68, 0x69, 0x74, 0x73, 0x5f, 0x61, 0x64, 0x64, 0x65,
	0x6e, 0x64, 0x18, 0x03, 0x20, 0x01, 0x28, 0x0d, 0x52, 0x0a, 0x68, 0x69, 0x74, 0x73, 0x41, 0x64,
	0x64, 0x65, 0x6e, 0x64, 0x22, 0xe7, 0x06, 0x0a, 0x11, 0x52, 0x61, 0x74, 0x65, 0x4c, 0x69, 0x6d,
	0x69, 0x74, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x55, 0x0a, 0x0c, 0x6f, 0x76,
	0x65, 0x72, 0x61, 0x6c, 0x6c, 0x5f, 0x63, 0x6f, 0x64, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0e,
	0x32, 0x32, 0x2e, 0x65, 0x6e, 0x76, 0x6f, 0x79, 0x2e, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65,
	0x2e, 0x72, 0x61, 0x74, 0x65, 0x6c, 0x69, 0x6d, 0x69, 0x74, 0x2e, 0x76, 0x32, 0x2e, 0x52, 0x61,
	0x74, 0x65, 0x4c, 0x69, 0x6d, 0x69, 0x74, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x2e,
	0x43, 0x6f, 0x64, 0x65, 0x52, 0x0b, 0x6f, 0x76, 0x65, 0x72, 0x61, 0x6c, 0x6c, 0x43, 0x6f, 0x64,
	0x65, 0x12, 0x5a, 0x0a, 0x08, 0x73, 0x74, 0x61, 0x74, 0x75, 0x73, 0x65, 0x73, 0x18, 0x02, 0x20,
	0x03, 0x28, 0x0b, 0x32, 0x3e, 0x2e, 0x65, 0x6e, 0x76, 0x6f, 0x79, 0x2e, 0x73, 0x65, 0x72, 0x76,
	0x69, 0x63, 0x65, 0x2e, 0x72, 0x61, 0x74, 0x65, 0x6c, 0x69, 0x6d, 0x69, 0x74, 0x2e, 0x76, 0x32,
	0x2e, 0x52, 0x61, 0x74, 0x65, 0x4c, 0x69, 0x6d, 0x69, 0x74, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e,
	0x73, 0x65, 0x2e, 0x44, 0x65, 0x73, 0x63, 0x72, 0x69, 0x70, 0x74, 0x6f, 0x72, 0x53, 0x74, 0x61,
	0x74, 0x75, 0x73, 0x52, 0x08, 0x73, 0x74, 0x61, 0x74, 0x75, 0x73, 0x65, 0x73, 0x12, 0x59, 0x0a,
	0x07, 0x68, 0x65, 0x61, 0x64, 0x65, 0x72, 0x73, 0x18, 0x03, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x1e,
	0x2e, 0x65, 0x6e, 0x76, 0x6f, 0x79, 0x2e, 0x61, 0x70, 0x69, 0x2e, 0x76, 0x32, 0x2e, 0x63, 0x6f,
	0x72, 0x65, 0x2e, 0x48, 0x65, 0x61, 0x64, 0x65, 0x72, 0x56, 0x61, 0x6c, 0x75, 0x65, 0x42, 0x1f,
	0xf2, 0x98, 0xfe, 0x8f, 0x05, 0x19, 0x0a, 0x17, 0x72, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65,
	0x5f, 0x68, 0x65, 0x61, 0x64, 0x65, 0x72, 0x73, 0x5f, 0x74, 0x6f, 0x5f, 0x61, 0x64, 0x64, 0x52,
	0x07, 0x68, 0x65, 0x61, 0x64, 0x65, 0x72, 0x73, 0x12, 0x53, 0x0a, 0x16, 0x72, 0x65, 0x71, 0x75,
	0x65, 0x73, 0x74, 0x5f, 0x68, 0x65, 0x61, 0x64, 0x65, 0x72, 0x73, 0x5f, 0x74, 0x6f, 0x5f, 0x61,
	0x64, 0x64, 0x18, 0x04, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x1e, 0x2e, 0x65, 0x6e, 0x76, 0x6f, 0x79,
	0x2e, 0x61, 0x70, 0x69, 0x2e, 0x76, 0x32, 0x2e, 0x63, 0x6f, 0x72, 0x65, 0x2e, 0x48, 0x65, 0x61,
	0x64, 0x65, 0x72, 0x56, 0x61, 0x6c, 0x75, 0x65, 0x52, 0x13, 0x72, 0x65, 0x71, 0x75, 0x65, 0x73,
	0x74, 0x48, 0x65, 0x61, 0x64, 0x65, 0x72, 0x73, 0x54, 0x6f, 0x41, 0x64, 0x64, 0x1a, 0xdd, 0x01,
	0x0a, 0x09, 0x52, 0x61, 0x74, 0x65, 0x4c, 0x69, 0x6d, 0x69, 0x74, 0x12, 0x12, 0x0a, 0x04, 0x6e,
	0x61, 0x6d, 0x65, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x12,
	0x2a, 0x0a, 0x11, 0x72, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x73, 0x5f, 0x70, 0x65, 0x72, 0x5f,
	0x75, 0x6e, 0x69, 0x74, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0d, 0x52, 0x0f, 0x72, 0x65, 0x71, 0x75,
	0x65, 0x73, 0x74, 0x73, 0x50, 0x65, 0x72, 0x55, 0x6e, 0x69, 0x74, 0x12, 0x50, 0x0a, 0x04, 0x75,
	0x6e, 0x69, 0x74, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0e, 0x32, 0x3c, 0x2e, 0x65, 0x6e, 0x76, 0x6f,
	0x79, 0x2e, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x2e, 0x72, 0x61, 0x74, 0x65, 0x6c, 0x69,
	0x6d, 0x69, 0x74, 0x2e, 0x76, 0x32, 0x2e, 0x52, 0x61, 0x74, 0x65, 0x4c, 0x69, 0x6d, 0x69, 0x74,
	0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x2e, 0x52, 0x61, 0x74, 0x65, 0x4c, 0x69, 0x6d,
	0x69, 0x74, 0x2e, 0x55, 0x6e, 0x69, 0x74, 0x52, 0x04, 0x75, 0x6e, 0x69, 0x74, 0x22, 0x3e, 0x0a,
	0x04, 0x55, 0x6e, 0x69, 0x74, 0x12, 0x0b, 0x0a, 0x07, 0x55, 0x4e, 0x4b, 0x4e, 0x4f, 0x57, 0x4e,
	0x10, 0x00, 0x12, 0x0a, 0x0a, 0x06, 0x53, 0x45, 0x43, 0x4f, 0x4e, 0x44, 0x10, 0x01, 0x12, 0x0a,
	0x0a, 0x06, 0x4d, 0x49, 0x4e, 0x55, 0x54, 0x45, 0x10, 0x02, 0x12, 0x08, 0x0a, 0x04, 0x48, 0x4f,
	0x55, 0x52, 0x10, 0x03, 0x12, 0x07, 0x0a, 0x03, 0x44, 0x41, 0x59, 0x10, 0x04, 0x1a, 0xe1, 0x01,
	0x0a, 0x10, 0x44, 0x65, 0x73, 0x63, 0x72, 0x69, 0x70, 0x74, 0x6f, 0x72, 0x53, 0x74, 0x61, 0x74,
	0x75, 0x73, 0x12, 0x46, 0x0a, 0x04, 0x63, 0x6f, 0x64, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0e,
	0x32, 0x32, 0x2e, 0x65, 0x6e, 0x76, 0x6f, 0x79, 0x2e, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65,
	0x2e, 0x72, 0x61, 0x74, 0x65, 0x6c, 0x69, 0x6d, 0x69, 0x74, 0x2e, 0x76, 0x32, 0x2e, 0x52, 0x61,
	0x74, 0x65, 0x4c, 0x69, 0x6d, 0x69, 0x74, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x2e,
	0x43, 0x6f, 0x64, 0x65, 0x52, 0x04, 0x63, 0x6f, 0x64, 0x65, 0x12, 0x5c, 0x0a, 0x0d, 0x63, 0x75,
	0x72, 0x72, 0x65, 0x6e, 0x74, 0x5f, 0x6c, 0x69, 0x6d, 0x69, 0x74, 0x18, 0x02, 0x20, 0x01, 0x28,
	0x0b, 0x32, 0x37, 0x2e, 0x65, 0x6e, 0x76, 0x6f, 0x79, 0x2e, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63,
	0x65, 0x2e, 0x72, 0x61, 0x74, 0x65, 0x6c, 0x69, 0x6d, 0x69, 0x74, 0x2e, 0x76, 0x32, 0x2e, 0x52,
	0x61, 0x74, 0x65, 0x4c, 0x69, 0x6d, 0x69, 0x74, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65,
	0x2e, 0x52, 0x61, 0x74, 0x65, 0x4c, 0x69, 0x6d, 0x69, 0x74, 0x52, 0x0c, 0x63, 0x75, 0x72, 0x72,
	0x65, 0x6e, 0x74, 0x4c, 0x69, 0x6d, 0x69, 0x74, 0x12, 0x27, 0x0a, 0x0f, 0x6c, 0x69, 0x6d, 0x69,
	0x74, 0x5f, 0x72, 0x65, 0x6d, 0x61, 0x69, 0x6e, 0x69, 0x6e, 0x67, 0x18, 0x03, 0x20, 0x01, 0x28,
	0x0d, 0x52, 0x0e, 0x6c, 0x69, 0x6d, 0x69, 0x74, 0x52, 0x65, 0x6d, 0x61, 0x69, 0x6e, 0x69, 0x6e,
	0x67, 0x22, 0x2b, 0x0a, 0x04, 0x43, 0x6f, 0x64, 0x65, 0x12, 0x0b, 0x0a, 0x07, 0x55, 0x4e, 0x4b,
	0x4e, 0x4f, 0x57, 0x4e, 0x10, 0x00, 0x12, 0x06, 0x0a, 0x02, 0x4f, 0x4b, 0x10, 0x01, 0x12, 0x0e,
	0x0a, 0x0a, 0x4f, 0x56, 0x45, 0x52, 0x5f, 0x4c, 0x49, 0x4d, 0x49, 0x54, 0x10, 0x02, 0x32, 0x84,
	0x01, 0x0a, 0x10, 0x52, 0x61, 0x74, 0x65, 0x4c, 0x69, 0x6d, 0x69, 0x74, 0x53, 0x65, 0x72, 0x76,
	0x69, 0x63, 0x65, 0x12, 0x70, 0x0a, 0x0f, 0x53, 0x68, 0x6f, 0x75, 0x6c, 0x64, 0x52, 0x61, 0x74,
	0x65, 0x4c, 0x69, 0x6d, 0x69, 0x74, 0x12, 0x2c, 0x2e, 0x65, 0x6e, 0x76, 0x6f, 0x79, 0x2e, 0x73,
	0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x2e, 0x72, 0x61, 0x74, 0x65, 0x6c, 0x69, 0x6d, 0x69, 0x74,
	0x2e, 0x76, 0x32, 0x2e, 0x52, 0x61, 0x74, 0x65, 0x4c, 0x69, 0x6d, 0x69, 0x74, 0x52, 0x65, 0x71,
	0x75, 0x65, 0x73, 0x74, 0x1a, 0x2d, 0x2e, 0x65, 0x6e, 0x76, 0x6f, 0x79, 0x2e, 0x73, 0x65, 0x72,
	0x76, 0x69, 0x63, 0x65, 0x2e, 0x72, 0x61, 0x74, 0x65, 0x6c, 0x69, 0x6d, 0x69, 0x74, 0x2e, 0x76,
	0x32, 0x2e, 0x52, 0x61, 0x74, 0x65, 0x4c, 0x69, 0x6d, 0x69, 0x74, 0x52, 0x65, 0x73, 0x70, 0x6f,
	0x6e, 0x73, 0x65, 0x22, 0x00, 0x42, 0x90, 0x01, 0x0a, 0x28, 0x69, 0x6f, 0x2e, 0x65, 0x6e, 0x76,
	0x6f, 0x79, 0x70, 0x72, 0x6f, 0x78, 0x79, 0x2e, 0x65, 0x6e, 0x76, 0x6f, 0x79, 0x2e, 0x73, 0x65,
	0x72, 0x76, 0x69, 0x63, 0x65, 0x2e, 0x72, 0x61, 0x74, 0x65, 0x6c, 0x69, 0x6d, 0x69, 0x74, 0x2e,
	0x76, 0x32, 0x42, 0x08, 0x52, 0x6c, 0x73, 0x50, 0x72, 0x6f, 0x74, 0x6f, 0x50, 0x01, 0x5a, 0x4d,
	0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x65, 0x6e, 0x76, 0x6f, 0x79,
	0x70, 0x72, 0x6f, 0x78, 0x79, 0x2f, 0x67, 0x6f, 0x2d, 0x63, 0x6f, 0x6e, 0x74, 0x72, 0x6f, 0x6c,
	0x2d, 0x70, 0x6c, 0x61, 0x6e, 0x65, 0x2f, 0x65, 0x6e, 0x76, 0x6f, 0x79, 0x2f, 0x73, 0x65, 0x72,
	0x76, 0x69, 0x63, 0x65, 0x2f, 0x72, 0x61, 0x74, 0x65, 0x6c, 0x69, 0x6d, 0x69, 0x74, 0x2f, 0x76,
	0x32, 0x3b, 0x72, 0x61, 0x74, 0x65, 0x6c, 0x69, 0x6d, 0x69, 0x74, 0x76, 0x32, 0x88, 0x01, 0x01,
	0xba, 0x80, 0xc8, 0xd1, 0x06, 0x02, 0x10, 0x01, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_envoy_service_ratelimit_v2_rls_proto_rawDescOnce sync.Once
	file_envoy_service_ratelimit_v2_rls_proto_rawDescData = file_envoy_service_ratelimit_v2_rls_proto_rawDesc
)

func file_envoy_service_ratelimit_v2_rls_proto_rawDescGZIP() []byte {
	file_envoy_service_ratelimit_v2_rls_proto_rawDescOnce.Do(func() {
		file_envoy_service_ratelimit_v2_rls_proto_rawDescData = protoimpl.X.CompressGZIP(file_envoy_service_ratelimit_v2_rls_proto_rawDescData)
	})
	return file_envoy_service_ratelimit_v2_rls_proto_rawDescData
}

var file_envoy_service_ratelimit_v2_rls_proto_enumTypes = make([]protoimpl.EnumInfo, 2)
var file_envoy_service_ratelimit_v2_rls_proto_msgTypes = make([]protoimpl.MessageInfo, 4)
var file_envoy_service_ratelimit_v2_rls_proto_goTypes = []interface{}{
	(RateLimitResponse_Code)(0),                // 0: envoy.service.ratelimit.v2.RateLimitResponse.Code
	(RateLimitResponse_RateLimit_Unit)(0),      // 1: envoy.service.ratelimit.v2.RateLimitResponse.RateLimit.Unit
	(*RateLimitRequest)(nil),                   // 2: envoy.service.ratelimit.v2.RateLimitRequest
	(*RateLimitResponse)(nil),                  // 3: envoy.service.ratelimit.v2.RateLimitResponse
	(*RateLimitResponse_RateLimit)(nil),        // 4: envoy.service.ratelimit.v2.RateLimitResponse.RateLimit
	(*RateLimitResponse_DescriptorStatus)(nil), // 5: envoy.service.ratelimit.v2.RateLimitResponse.DescriptorStatus
	(*ratelimit.RateLimitDescriptor)(nil),      // 6: envoy.api.v2.ratelimit.RateLimitDescriptor
	(*core.HeaderValue)(nil),                   // 7: envoy.api.v2.core.HeaderValue
}
var file_envoy_service_ratelimit_v2_rls_proto_depIdxs = []int32{
	6, // 0: envoy.service.ratelimit.v2.RateLimitRequest.descriptors:type_name -> envoy.api.v2.ratelimit.RateLimitDescriptor
	0, // 1: envoy.service.ratelimit.v2.RateLimitResponse.overall_code:type_name -> envoy.service.ratelimit.v2.RateLimitResponse.Code
	5, // 2: envoy.service.ratelimit.v2.RateLimitResponse.statuses:type_name -> envoy.service.ratelimit.v2.RateLimitResponse.DescriptorStatus
	7, // 3: envoy.service.ratelimit.v2.RateLimitResponse.headers:type_name -> envoy.api.v2.core.HeaderValue
	7, // 4: envoy.service.ratelimit.v2.RateLimitResponse.request_headers_to_add:type_name -> envoy.api.v2.core.HeaderValue
	1, // 5: envoy.service.ratelimit.v2.RateLimitResponse.RateLimit.unit:type_name -> envoy.service.ratelimit.v2.RateLimitResponse.RateLimit.Unit
	0, // 6: envoy.service.ratelimit.v2.RateLimitResponse.DescriptorStatus.code:type_name -> envoy.service.ratelimit.v2.RateLimitResponse.Code
	4, // 7: envoy.service.ratelimit.v2.RateLimitResponse.DescriptorStatus.current_limit:type_name -> envoy.service.ratelimit.v2.RateLimitResponse.RateLimit
	2, // 8: envoy.service.ratelimit.v2.RateLimitService.ShouldRateLimit:input_type -> envoy.service.ratelimit.v2.RateLimitRequest
	3, // 9: envoy.service.ratelimit.v2.RateLimitService.ShouldRateLimit:output_type -> envoy.service.ratelimit.v2.RateLimitResponse
	9, // [9:10] is the sub-list for method output_type
	8, // [8:9] is the sub-list for method input_type
	8, // [8:8] is the sub-list for extension type_name
	8, // [8:8] is the sub-list for extension extendee
	0, // [0:8] is the sub-list for field type_name
}

func init() { file_envoy_service_ratelimit_v2_rls_proto_init() }
func file_envoy_service_ratelimit_v2_rls_proto_init() {
	if File_envoy_service_ratelimit_v2_rls_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_envoy_service_ratelimit_v2_rls_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*RateLimitRequest); i {
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
		file_envoy_service_ratelimit_v2_rls_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*RateLimitResponse); i {
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
		file_envoy_service_ratelimit_v2_rls_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*RateLimitResponse_RateLimit); i {
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
		file_envoy_service_ratelimit_v2_rls_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*RateLimitResponse_DescriptorStatus); i {
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
			RawDescriptor: file_envoy_service_ratelimit_v2_rls_proto_rawDesc,
			NumEnums:      2,
			NumMessages:   4,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_envoy_service_ratelimit_v2_rls_proto_goTypes,
		DependencyIndexes: file_envoy_service_ratelimit_v2_rls_proto_depIdxs,
		EnumInfos:         file_envoy_service_ratelimit_v2_rls_proto_enumTypes,
		MessageInfos:      file_envoy_service_ratelimit_v2_rls_proto_msgTypes,
	}.Build()
	File_envoy_service_ratelimit_v2_rls_proto = out.File
	file_envoy_service_ratelimit_v2_rls_proto_rawDesc = nil
	file_envoy_service_ratelimit_v2_rls_proto_goTypes = nil
	file_envoy_service_ratelimit_v2_rls_proto_depIdxs = nil
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConnInterface

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion6

// RateLimitServiceClient is the client API for RateLimitService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type RateLimitServiceClient interface {
	// Determine whether rate limiting should take place.
	ShouldRateLimit(ctx context.Context, in *RateLimitRequest, opts ...grpc.CallOption) (*RateLimitResponse, error)
}

type rateLimitServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewRateLimitServiceClient(cc grpc.ClientConnInterface) RateLimitServiceClient {
	return &rateLimitServiceClient{cc}
}

func (c *rateLimitServiceClient) ShouldRateLimit(ctx context.Context, in *RateLimitRequest, opts ...grpc.CallOption) (*RateLimitResponse, error) {
	out := new(RateLimitResponse)
	err := c.cc.Invoke(ctx, "/envoy.service.ratelimit.v2.RateLimitService/ShouldRateLimit", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// RateLimitServiceServer is the server API for RateLimitService service.
type RateLimitServiceServer interface {
	// Determine whether rate limiting should take place.
	ShouldRateLimit(context.Context, *RateLimitRequest) (*RateLimitResponse, error)
}

// UnimplementedRateLimitServiceServer can be embedded to have forward compatible implementations.
type UnimplementedRateLimitServiceServer struct {
}

func (*UnimplementedRateLimitServiceServer) ShouldRateLimit(context.Context, *RateLimitRequest) (*RateLimitResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method ShouldRateLimit not implemented")
}

func RegisterRateLimitServiceServer(s *grpc.Server, srv RateLimitServiceServer) {
	s.RegisterService(&_RateLimitService_serviceDesc, srv)
}

func _RateLimitService_ShouldRateLimit_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(RateLimitRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(RateLimitServiceServer).ShouldRateLimit(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/envoy.service.ratelimit.v2.RateLimitService/ShouldRateLimit",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(RateLimitServiceServer).ShouldRateLimit(ctx, req.(*RateLimitRequest))
	}
	return interceptor(ctx, in, info, handler)
}

var _RateLimitService_serviceDesc = grpc.ServiceDesc{
	ServiceName: "envoy.service.ratelimit.v2.RateLimitService",
	HandlerType: (*RateLimitServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "ShouldRateLimit",
			Handler:    _RateLimitService_ShouldRateLimit_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "envoy/service/ratelimit/v2/rls.proto",
}
