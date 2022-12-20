// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.28.1
// 	protoc        v3.21.7
// source: envoy/config/filter/network/local_rate_limit/v2alpha/local_rate_limit.proto

package v2alpha

import (
	_ "github.com/cncf/xds/go/udpa/annotations"
	core "github.com/emissary-ingress/emissary/v3/pkg/api/envoy/api/v2/core"
	_type "github.com/emissary-ingress/emissary/v3/pkg/api/envoy/type"
	_ "github.com/envoyproxy/protoc-gen-validate/validate"
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

type LocalRateLimit struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// The prefix to use when emitting :ref:`statistics
	// <config_network_filters_local_rate_limit_stats>`.
	StatPrefix string `protobuf:"bytes,1,opt,name=stat_prefix,json=statPrefix,proto3" json:"stat_prefix,omitempty"`
	// The token bucket configuration to use for rate limiting connections that are processed by the
	// filter's filter chain. Each incoming connection processed by the filter consumes a single
	// token. If the token is available, the connection will be allowed. If no tokens are available,
	// the connection will be immediately closed.
	//
	// .. note::
	//   In the current implementation each filter and filter chain has an independent rate limit.
	//
	// .. note::
	//   In the current implementation the token bucket's :ref:`fill_interval
	//   <envoy_api_field_type.TokenBucket.fill_interval>` must be >= 50ms to avoid too aggressive
	//   refills.
	TokenBucket *_type.TokenBucket `protobuf:"bytes,2,opt,name=token_bucket,json=tokenBucket,proto3" json:"token_bucket,omitempty"`
	// Runtime flag that controls whether the filter is enabled or not. If not specified, defaults
	// to enabled.
	RuntimeEnabled *core.RuntimeFeatureFlag `protobuf:"bytes,3,opt,name=runtime_enabled,json=runtimeEnabled,proto3" json:"runtime_enabled,omitempty"`
}

func (x *LocalRateLimit) Reset() {
	*x = LocalRateLimit{}
	if protoimpl.UnsafeEnabled {
		mi := &file_envoy_config_filter_network_local_rate_limit_v2alpha_local_rate_limit_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *LocalRateLimit) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*LocalRateLimit) ProtoMessage() {}

func (x *LocalRateLimit) ProtoReflect() protoreflect.Message {
	mi := &file_envoy_config_filter_network_local_rate_limit_v2alpha_local_rate_limit_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use LocalRateLimit.ProtoReflect.Descriptor instead.
func (*LocalRateLimit) Descriptor() ([]byte, []int) {
	return file_envoy_config_filter_network_local_rate_limit_v2alpha_local_rate_limit_proto_rawDescGZIP(), []int{0}
}

func (x *LocalRateLimit) GetStatPrefix() string {
	if x != nil {
		return x.StatPrefix
	}
	return ""
}

func (x *LocalRateLimit) GetTokenBucket() *_type.TokenBucket {
	if x != nil {
		return x.TokenBucket
	}
	return nil
}

func (x *LocalRateLimit) GetRuntimeEnabled() *core.RuntimeFeatureFlag {
	if x != nil {
		return x.RuntimeEnabled
	}
	return nil
}

var File_envoy_config_filter_network_local_rate_limit_v2alpha_local_rate_limit_proto protoreflect.FileDescriptor

var file_envoy_config_filter_network_local_rate_limit_v2alpha_local_rate_limit_proto_rawDesc = []byte{
	0x0a, 0x4b, 0x65, 0x6e, 0x76, 0x6f, 0x79, 0x2f, 0x63, 0x6f, 0x6e, 0x66, 0x69, 0x67, 0x2f, 0x66,
	0x69, 0x6c, 0x74, 0x65, 0x72, 0x2f, 0x6e, 0x65, 0x74, 0x77, 0x6f, 0x72, 0x6b, 0x2f, 0x6c, 0x6f,
	0x63, 0x61, 0x6c, 0x5f, 0x72, 0x61, 0x74, 0x65, 0x5f, 0x6c, 0x69, 0x6d, 0x69, 0x74, 0x2f, 0x76,
	0x32, 0x61, 0x6c, 0x70, 0x68, 0x61, 0x2f, 0x6c, 0x6f, 0x63, 0x61, 0x6c, 0x5f, 0x72, 0x61, 0x74,
	0x65, 0x5f, 0x6c, 0x69, 0x6d, 0x69, 0x74, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x34, 0x65,
	0x6e, 0x76, 0x6f, 0x79, 0x2e, 0x63, 0x6f, 0x6e, 0x66, 0x69, 0x67, 0x2e, 0x66, 0x69, 0x6c, 0x74,
	0x65, 0x72, 0x2e, 0x6e, 0x65, 0x74, 0x77, 0x6f, 0x72, 0x6b, 0x2e, 0x6c, 0x6f, 0x63, 0x61, 0x6c,
	0x5f, 0x72, 0x61, 0x74, 0x65, 0x5f, 0x6c, 0x69, 0x6d, 0x69, 0x74, 0x2e, 0x76, 0x32, 0x61, 0x6c,
	0x70, 0x68, 0x61, 0x1a, 0x1c, 0x65, 0x6e, 0x76, 0x6f, 0x79, 0x2f, 0x61, 0x70, 0x69, 0x2f, 0x76,
	0x32, 0x2f, 0x63, 0x6f, 0x72, 0x65, 0x2f, 0x62, 0x61, 0x73, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x1a, 0x1d, 0x65, 0x6e, 0x76, 0x6f, 0x79, 0x2f, 0x74, 0x79, 0x70, 0x65, 0x2f, 0x74, 0x6f,
	0x6b, 0x65, 0x6e, 0x5f, 0x62, 0x75, 0x63, 0x6b, 0x65, 0x74, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x1a, 0x1e, 0x75, 0x64, 0x70, 0x61, 0x2f, 0x61, 0x6e, 0x6e, 0x6f, 0x74, 0x61, 0x74, 0x69, 0x6f,
	0x6e, 0x73, 0x2f, 0x6d, 0x69, 0x67, 0x72, 0x61, 0x74, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x1a, 0x1d, 0x75, 0x64, 0x70, 0x61, 0x2f, 0x61, 0x6e, 0x6e, 0x6f, 0x74, 0x61, 0x74, 0x69, 0x6f,
	0x6e, 0x73, 0x2f, 0x73, 0x74, 0x61, 0x74, 0x75, 0x73, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a,
	0x17, 0x76, 0x61, 0x6c, 0x69, 0x64, 0x61, 0x74, 0x65, 0x2f, 0x76, 0x61, 0x6c, 0x69, 0x64, 0x61,
	0x74, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0xd0, 0x01, 0x0a, 0x0e, 0x4c, 0x6f, 0x63,
	0x61, 0x6c, 0x52, 0x61, 0x74, 0x65, 0x4c, 0x69, 0x6d, 0x69, 0x74, 0x12, 0x28, 0x0a, 0x0b, 0x73,
	0x74, 0x61, 0x74, 0x5f, 0x70, 0x72, 0x65, 0x66, 0x69, 0x78, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09,
	0x42, 0x07, 0xfa, 0x42, 0x04, 0x72, 0x02, 0x20, 0x01, 0x52, 0x0a, 0x73, 0x74, 0x61, 0x74, 0x50,
	0x72, 0x65, 0x66, 0x69, 0x78, 0x12, 0x44, 0x0a, 0x0c, 0x74, 0x6f, 0x6b, 0x65, 0x6e, 0x5f, 0x62,
	0x75, 0x63, 0x6b, 0x65, 0x74, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x17, 0x2e, 0x65, 0x6e,
	0x76, 0x6f, 0x79, 0x2e, 0x74, 0x79, 0x70, 0x65, 0x2e, 0x54, 0x6f, 0x6b, 0x65, 0x6e, 0x42, 0x75,
	0x63, 0x6b, 0x65, 0x74, 0x42, 0x08, 0xfa, 0x42, 0x05, 0x8a, 0x01, 0x02, 0x10, 0x01, 0x52, 0x0b,
	0x74, 0x6f, 0x6b, 0x65, 0x6e, 0x42, 0x75, 0x63, 0x6b, 0x65, 0x74, 0x12, 0x4e, 0x0a, 0x0f, 0x72,
	0x75, 0x6e, 0x74, 0x69, 0x6d, 0x65, 0x5f, 0x65, 0x6e, 0x61, 0x62, 0x6c, 0x65, 0x64, 0x18, 0x03,
	0x20, 0x01, 0x28, 0x0b, 0x32, 0x25, 0x2e, 0x65, 0x6e, 0x76, 0x6f, 0x79, 0x2e, 0x61, 0x70, 0x69,
	0x2e, 0x76, 0x32, 0x2e, 0x63, 0x6f, 0x72, 0x65, 0x2e, 0x52, 0x75, 0x6e, 0x74, 0x69, 0x6d, 0x65,
	0x46, 0x65, 0x61, 0x74, 0x75, 0x72, 0x65, 0x46, 0x6c, 0x61, 0x67, 0x52, 0x0e, 0x72, 0x75, 0x6e,
	0x74, 0x69, 0x6d, 0x65, 0x45, 0x6e, 0x61, 0x62, 0x6c, 0x65, 0x64, 0x42, 0xfb, 0x01, 0x0a, 0x42,
	0x69, 0x6f, 0x2e, 0x65, 0x6e, 0x76, 0x6f, 0x79, 0x70, 0x72, 0x6f, 0x78, 0x79, 0x2e, 0x65, 0x6e,
	0x76, 0x6f, 0x79, 0x2e, 0x63, 0x6f, 0x6e, 0x66, 0x69, 0x67, 0x2e, 0x66, 0x69, 0x6c, 0x74, 0x65,
	0x72, 0x2e, 0x6e, 0x65, 0x74, 0x77, 0x6f, 0x72, 0x6b, 0x2e, 0x6c, 0x6f, 0x63, 0x61, 0x6c, 0x5f,
	0x72, 0x61, 0x74, 0x65, 0x5f, 0x6c, 0x69, 0x6d, 0x69, 0x74, 0x2e, 0x76, 0x32, 0x61, 0x6c, 0x70,
	0x68, 0x61, 0x42, 0x13, 0x4c, 0x6f, 0x63, 0x61, 0x6c, 0x52, 0x61, 0x74, 0x65, 0x4c, 0x69, 0x6d,
	0x69, 0x74, 0x50, 0x72, 0x6f, 0x74, 0x6f, 0x50, 0x01, 0x5a, 0x5b, 0x67, 0x69, 0x74, 0x68, 0x75,
	0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x65, 0x6e, 0x76, 0x6f, 0x79, 0x70, 0x72, 0x6f, 0x78, 0x79,
	0x2f, 0x67, 0x6f, 0x2d, 0x63, 0x6f, 0x6e, 0x74, 0x72, 0x6f, 0x6c, 0x2d, 0x70, 0x6c, 0x61, 0x6e,
	0x65, 0x2f, 0x65, 0x6e, 0x76, 0x6f, 0x79, 0x2f, 0x63, 0x6f, 0x6e, 0x66, 0x69, 0x67, 0x2f, 0x66,
	0x69, 0x6c, 0x74, 0x65, 0x72, 0x2f, 0x6e, 0x65, 0x74, 0x77, 0x6f, 0x72, 0x6b, 0x2f, 0x6c, 0x6f,
	0x63, 0x61, 0x6c, 0x5f, 0x72, 0x61, 0x74, 0x65, 0x5f, 0x6c, 0x69, 0x6d, 0x69, 0x74, 0x2f, 0x76,
	0x32, 0x61, 0x6c, 0x70, 0x68, 0x61, 0xf2, 0x98, 0xfe, 0x8f, 0x05, 0x35, 0x12, 0x33, 0x65, 0x6e,
	0x76, 0x6f, 0x79, 0x2e, 0x65, 0x78, 0x74, 0x65, 0x6e, 0x73, 0x69, 0x6f, 0x6e, 0x73, 0x2e, 0x66,
	0x69, 0x6c, 0x74, 0x65, 0x72, 0x73, 0x2e, 0x6e, 0x65, 0x74, 0x77, 0x6f, 0x72, 0x6b, 0x2e, 0x6c,
	0x6f, 0x63, 0x61, 0x6c, 0x5f, 0x72, 0x61, 0x74, 0x65, 0x6c, 0x69, 0x6d, 0x69, 0x74, 0x2e, 0x76,
	0x33, 0xba, 0x80, 0xc8, 0xd1, 0x06, 0x02, 0x10, 0x01, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x33,
}

var (
	file_envoy_config_filter_network_local_rate_limit_v2alpha_local_rate_limit_proto_rawDescOnce sync.Once
	file_envoy_config_filter_network_local_rate_limit_v2alpha_local_rate_limit_proto_rawDescData = file_envoy_config_filter_network_local_rate_limit_v2alpha_local_rate_limit_proto_rawDesc
)

func file_envoy_config_filter_network_local_rate_limit_v2alpha_local_rate_limit_proto_rawDescGZIP() []byte {
	file_envoy_config_filter_network_local_rate_limit_v2alpha_local_rate_limit_proto_rawDescOnce.Do(func() {
		file_envoy_config_filter_network_local_rate_limit_v2alpha_local_rate_limit_proto_rawDescData = protoimpl.X.CompressGZIP(file_envoy_config_filter_network_local_rate_limit_v2alpha_local_rate_limit_proto_rawDescData)
	})
	return file_envoy_config_filter_network_local_rate_limit_v2alpha_local_rate_limit_proto_rawDescData
}

var file_envoy_config_filter_network_local_rate_limit_v2alpha_local_rate_limit_proto_msgTypes = make([]protoimpl.MessageInfo, 1)
var file_envoy_config_filter_network_local_rate_limit_v2alpha_local_rate_limit_proto_goTypes = []interface{}{
	(*LocalRateLimit)(nil),          // 0: envoy.config.filter.network.local_rate_limit.v2alpha.LocalRateLimit
	(*_type.TokenBucket)(nil),       // 1: envoy.type.TokenBucket
	(*core.RuntimeFeatureFlag)(nil), // 2: envoy.api.v2.core.RuntimeFeatureFlag
}
var file_envoy_config_filter_network_local_rate_limit_v2alpha_local_rate_limit_proto_depIdxs = []int32{
	1, // 0: envoy.config.filter.network.local_rate_limit.v2alpha.LocalRateLimit.token_bucket:type_name -> envoy.type.TokenBucket
	2, // 1: envoy.config.filter.network.local_rate_limit.v2alpha.LocalRateLimit.runtime_enabled:type_name -> envoy.api.v2.core.RuntimeFeatureFlag
	2, // [2:2] is the sub-list for method output_type
	2, // [2:2] is the sub-list for method input_type
	2, // [2:2] is the sub-list for extension type_name
	2, // [2:2] is the sub-list for extension extendee
	0, // [0:2] is the sub-list for field type_name
}

func init() { file_envoy_config_filter_network_local_rate_limit_v2alpha_local_rate_limit_proto_init() }
func file_envoy_config_filter_network_local_rate_limit_v2alpha_local_rate_limit_proto_init() {
	if File_envoy_config_filter_network_local_rate_limit_v2alpha_local_rate_limit_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_envoy_config_filter_network_local_rate_limit_v2alpha_local_rate_limit_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*LocalRateLimit); i {
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
			RawDescriptor: file_envoy_config_filter_network_local_rate_limit_v2alpha_local_rate_limit_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   1,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_envoy_config_filter_network_local_rate_limit_v2alpha_local_rate_limit_proto_goTypes,
		DependencyIndexes: file_envoy_config_filter_network_local_rate_limit_v2alpha_local_rate_limit_proto_depIdxs,
		MessageInfos:      file_envoy_config_filter_network_local_rate_limit_v2alpha_local_rate_limit_proto_msgTypes,
	}.Build()
	File_envoy_config_filter_network_local_rate_limit_v2alpha_local_rate_limit_proto = out.File
	file_envoy_config_filter_network_local_rate_limit_v2alpha_local_rate_limit_proto_rawDesc = nil
	file_envoy_config_filter_network_local_rate_limit_v2alpha_local_rate_limit_proto_goTypes = nil
	file_envoy_config_filter_network_local_rate_limit_v2alpha_local_rate_limit_proto_depIdxs = nil
}
