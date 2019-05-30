// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: envoy/config/filter/http/ip_tagging/v2/ip_tagging.proto

package v2

import (
	fmt "fmt"
	core "github.com/datawire/ambassador/go/apis/envoy/api/v2/core"
	_ "github.com/envoyproxy/protoc-gen-validate/validate"
	proto "github.com/gogo/protobuf/proto"
	io "io"
	math "math"
)

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// This is a compile-time assertion to ensure that this generated file
// is compatible with the proto package it is being compiled against.
// A compilation error at this line likely means your copy of the
// proto package needs to be updated.
const _ = proto.GoGoProtoPackageIsVersion2 // please upgrade the proto package

// The type of requests the filter should apply to. The supported types
// are internal, external or both. The
// :ref:`x-forwarded-for<config_http_conn_man_headers_x-forwarded-for_internal_origin>` header is
// used to determine if a request is internal and will result in
// :ref:`x-envoy-internal<config_http_conn_man_headers_x-envoy-internal>`
// being set. The filter defaults to both, and it will apply to all request types.
type IPTagging_RequestType int32

const (
	// Both external and internal requests will be tagged. This is the default value.
	IPTagging_BOTH IPTagging_RequestType = 0
	// Only internal requests will be tagged.
	IPTagging_INTERNAL IPTagging_RequestType = 1
	// Only external requests will be tagged.
	IPTagging_EXTERNAL IPTagging_RequestType = 2
)

var IPTagging_RequestType_name = map[int32]string{
	0: "BOTH",
	1: "INTERNAL",
	2: "EXTERNAL",
}

var IPTagging_RequestType_value = map[string]int32{
	"BOTH":     0,
	"INTERNAL": 1,
	"EXTERNAL": 2,
}

func (x IPTagging_RequestType) String() string {
	return proto.EnumName(IPTagging_RequestType_name, int32(x))
}

func (IPTagging_RequestType) EnumDescriptor() ([]byte, []int) {
	return fileDescriptor_f4275d0b367744d2, []int{0, 0}
}

type IPTagging struct {
	// The type of request the filter should apply to.
	RequestType IPTagging_RequestType `protobuf:"varint,1,opt,name=request_type,json=requestType,proto3,enum=envoy.config.filter.http.ip_tagging.v2.IPTagging_RequestType" json:"request_type,omitempty"`
	// [#comment:TODO(ccaraman): Extend functionality to load IP tags from file system.
	// Tracked by issue https://github.com/envoyproxy/envoy/issues/2695]
	// The set of IP tags for the filter.
	IpTags               []*IPTagging_IPTag `protobuf:"bytes,4,rep,name=ip_tags,json=ipTags,proto3" json:"ip_tags,omitempty"`
	XXX_NoUnkeyedLiteral struct{}           `json:"-"`
	XXX_unrecognized     []byte             `json:"-"`
	XXX_sizecache        int32              `json:"-"`
}

func (m *IPTagging) Reset()         { *m = IPTagging{} }
func (m *IPTagging) String() string { return proto.CompactTextString(m) }
func (*IPTagging) ProtoMessage()    {}
func (*IPTagging) Descriptor() ([]byte, []int) {
	return fileDescriptor_f4275d0b367744d2, []int{0}
}
func (m *IPTagging) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *IPTagging) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_IPTagging.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalTo(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *IPTagging) XXX_Merge(src proto.Message) {
	xxx_messageInfo_IPTagging.Merge(m, src)
}
func (m *IPTagging) XXX_Size() int {
	return m.Size()
}
func (m *IPTagging) XXX_DiscardUnknown() {
	xxx_messageInfo_IPTagging.DiscardUnknown(m)
}

var xxx_messageInfo_IPTagging proto.InternalMessageInfo

func (m *IPTagging) GetRequestType() IPTagging_RequestType {
	if m != nil {
		return m.RequestType
	}
	return IPTagging_BOTH
}

func (m *IPTagging) GetIpTags() []*IPTagging_IPTag {
	if m != nil {
		return m.IpTags
	}
	return nil
}

// Supplies the IP tag name and the IP address subnets.
type IPTagging_IPTag struct {
	// Specifies the IP tag name to apply.
	IpTagName string `protobuf:"bytes,1,opt,name=ip_tag_name,json=ipTagName,proto3" json:"ip_tag_name,omitempty"`
	// A list of IP address subnets that will be tagged with
	// ip_tag_name. Both IPv4 and IPv6 are supported.
	IpList               []*core.CidrRange `protobuf:"bytes,2,rep,name=ip_list,json=ipList,proto3" json:"ip_list,omitempty"`
	XXX_NoUnkeyedLiteral struct{}          `json:"-"`
	XXX_unrecognized     []byte            `json:"-"`
	XXX_sizecache        int32             `json:"-"`
}

func (m *IPTagging_IPTag) Reset()         { *m = IPTagging_IPTag{} }
func (m *IPTagging_IPTag) String() string { return proto.CompactTextString(m) }
func (*IPTagging_IPTag) ProtoMessage()    {}
func (*IPTagging_IPTag) Descriptor() ([]byte, []int) {
	return fileDescriptor_f4275d0b367744d2, []int{0, 0}
}
func (m *IPTagging_IPTag) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *IPTagging_IPTag) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_IPTagging_IPTag.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalTo(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *IPTagging_IPTag) XXX_Merge(src proto.Message) {
	xxx_messageInfo_IPTagging_IPTag.Merge(m, src)
}
func (m *IPTagging_IPTag) XXX_Size() int {
	return m.Size()
}
func (m *IPTagging_IPTag) XXX_DiscardUnknown() {
	xxx_messageInfo_IPTagging_IPTag.DiscardUnknown(m)
}

var xxx_messageInfo_IPTagging_IPTag proto.InternalMessageInfo

func (m *IPTagging_IPTag) GetIpTagName() string {
	if m != nil {
		return m.IpTagName
	}
	return ""
}

func (m *IPTagging_IPTag) GetIpList() []*core.CidrRange {
	if m != nil {
		return m.IpList
	}
	return nil
}

func init() {
	proto.RegisterEnum("envoy.config.filter.http.ip_tagging.v2.IPTagging_RequestType", IPTagging_RequestType_name, IPTagging_RequestType_value)
	proto.RegisterType((*IPTagging)(nil), "envoy.config.filter.http.ip_tagging.v2.IPTagging")
	proto.RegisterType((*IPTagging_IPTag)(nil), "envoy.config.filter.http.ip_tagging.v2.IPTagging.IPTag")
}

func init() {
	proto.RegisterFile("envoy/config/filter/http/ip_tagging/v2/ip_tagging.proto", fileDescriptor_f4275d0b367744d2)
}

var fileDescriptor_f4275d0b367744d2 = []byte{
	// 373 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x94, 0x51, 0x3d, 0x4b, 0xfb, 0x40,
	0x18, 0xff, 0x5f, 0xfa, 0xf2, 0x6f, 0x2f, 0xa5, 0x94, 0x2c, 0x96, 0x22, 0xb1, 0x74, 0x90, 0x4e,
	0x17, 0x48, 0x95, 0x4e, 0x0e, 0x46, 0x0a, 0x16, 0x4a, 0x2d, 0x21, 0x43, 0x11, 0xb1, 0x9c, 0xcd,
	0x35, 0x1e, 0xa4, 0xc9, 0x79, 0x39, 0x83, 0x59, 0xfd, 0x08, 0x7e, 0x1c, 0x27, 0x47, 0x47, 0x3f,
	0x82, 0x74, 0x13, 0xfc, 0x10, 0x92, 0xbb, 0xa8, 0x1d, 0xeb, 0x76, 0x0f, 0xbf, 0xe7, 0xf7, 0x76,
	0x0f, 0x1c, 0x92, 0x28, 0x8d, 0x33, 0x6b, 0x19, 0x47, 0x2b, 0x1a, 0x58, 0x2b, 0x1a, 0x0a, 0xc2,
	0xad, 0x5b, 0x21, 0x98, 0x45, 0xd9, 0x42, 0xe0, 0x20, 0xa0, 0x51, 0x60, 0xa5, 0xf6, 0xd6, 0x84,
	0x18, 0x8f, 0x45, 0x6c, 0x1c, 0x4a, 0x22, 0x52, 0x44, 0xa4, 0x88, 0x28, 0x27, 0xa2, 0xad, 0xd5,
	0xd4, 0xee, 0x1c, 0x28, 0x03, 0xcc, 0x68, 0x2e, 0xb3, 0x8c, 0x39, 0xb1, 0xb0, 0xef, 0x73, 0x92,
	0x24, 0x4a, 0xa8, 0xb3, 0x97, 0xe2, 0x90, 0xfa, 0x58, 0x10, 0xeb, 0xfb, 0xa1, 0x80, 0xde, 0xa7,
	0x06, 0xeb, 0xe3, 0x99, 0xa7, 0xa4, 0x8c, 0x10, 0x36, 0x38, 0xb9, 0xbb, 0x27, 0x89, 0x58, 0x88,
	0x8c, 0x91, 0x36, 0xe8, 0x82, 0x7e, 0xd3, 0x3e, 0x41, 0xbb, 0xc5, 0x40, 0x3f, 0x42, 0xc8, 0x55,
	0x2a, 0x5e, 0xc6, 0x88, 0x03, 0x9f, 0x3f, 0x5e, 0x4a, 0x95, 0x47, 0xa0, 0xb5, 0x80, 0xab, 0xf3,
	0x5f, 0xc0, 0xb8, 0x82, 0xff, 0x15, 0x3f, 0x69, 0x97, 0xbb, 0xa5, 0xbe, 0x6e, 0x0f, 0xff, 0x6e,
	0x24, 0x5f, 0x85, 0xc5, 0x13, 0xd0, 0x6a, 0xc0, 0xad, 0x52, 0xe6, 0xe1, 0x20, 0xe9, 0x5c, 0xc3,
	0x8a, 0x04, 0x0d, 0x13, 0xea, 0x8a, 0xbd, 0x88, 0xf0, 0x5a, 0x75, 0xaa, 0xbb, 0x75, 0xb9, 0x35,
	0xc5, 0x6b, 0x62, 0x1c, 0xcb, 0x18, 0x21, 0x4d, 0x44, 0x5b, 0x93, 0x31, 0xf6, 0x8b, 0x18, 0x98,
	0xd1, 0xdc, 0x2c, 0xff, 0x4e, 0x74, 0x46, 0x7d, 0xee, 0xe2, 0x28, 0x20, 0xb9, 0xfe, 0x84, 0x26,
	0xa2, 0x37, 0x80, 0xfa, 0x56, 0x4b, 0xa3, 0x06, 0xcb, 0xce, 0x85, 0x77, 0xde, 0xfa, 0x67, 0x34,
	0x60, 0x6d, 0x3c, 0xf5, 0x46, 0xee, 0xf4, 0x74, 0xd2, 0x02, 0xf9, 0x34, 0x9a, 0x17, 0x93, 0xe6,
	0xcc, 0x5f, 0x37, 0x26, 0x78, 0xdb, 0x98, 0xe0, 0x7d, 0x63, 0x02, 0x78, 0x44, 0x63, 0x65, 0xc5,
	0x78, 0xfc, 0x90, 0xed, 0x58, 0xde, 0x69, 0x8e, 0x59, 0xd1, 0x7e, 0x96, 0x9f, 0x70, 0x06, 0x2e,
	0xb5, 0xd4, 0xbe, 0xa9, 0xca, 0x7b, 0x0e, 0xbe, 0x02, 0x00, 0x00, 0xff, 0xff, 0x02, 0x51, 0x7f,
	0x1e, 0x6c, 0x02, 0x00, 0x00,
}

func (m *IPTagging) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalTo(dAtA)
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *IPTagging) MarshalTo(dAtA []byte) (int, error) {
	var i int
	_ = i
	var l int
	_ = l
	if m.RequestType != 0 {
		dAtA[i] = 0x8
		i++
		i = encodeVarintIpTagging(dAtA, i, uint64(m.RequestType))
	}
	if len(m.IpTags) > 0 {
		for _, msg := range m.IpTags {
			dAtA[i] = 0x22
			i++
			i = encodeVarintIpTagging(dAtA, i, uint64(msg.Size()))
			n, err := msg.MarshalTo(dAtA[i:])
			if err != nil {
				return 0, err
			}
			i += n
		}
	}
	if m.XXX_unrecognized != nil {
		i += copy(dAtA[i:], m.XXX_unrecognized)
	}
	return i, nil
}

func (m *IPTagging_IPTag) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalTo(dAtA)
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *IPTagging_IPTag) MarshalTo(dAtA []byte) (int, error) {
	var i int
	_ = i
	var l int
	_ = l
	if len(m.IpTagName) > 0 {
		dAtA[i] = 0xa
		i++
		i = encodeVarintIpTagging(dAtA, i, uint64(len(m.IpTagName)))
		i += copy(dAtA[i:], m.IpTagName)
	}
	if len(m.IpList) > 0 {
		for _, msg := range m.IpList {
			dAtA[i] = 0x12
			i++
			i = encodeVarintIpTagging(dAtA, i, uint64(msg.Size()))
			n, err := msg.MarshalTo(dAtA[i:])
			if err != nil {
				return 0, err
			}
			i += n
		}
	}
	if m.XXX_unrecognized != nil {
		i += copy(dAtA[i:], m.XXX_unrecognized)
	}
	return i, nil
}

func encodeVarintIpTagging(dAtA []byte, offset int, v uint64) int {
	for v >= 1<<7 {
		dAtA[offset] = uint8(v&0x7f | 0x80)
		v >>= 7
		offset++
	}
	dAtA[offset] = uint8(v)
	return offset + 1
}
func (m *IPTagging) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	if m.RequestType != 0 {
		n += 1 + sovIpTagging(uint64(m.RequestType))
	}
	if len(m.IpTags) > 0 {
		for _, e := range m.IpTags {
			l = e.Size()
			n += 1 + l + sovIpTagging(uint64(l))
		}
	}
	if m.XXX_unrecognized != nil {
		n += len(m.XXX_unrecognized)
	}
	return n
}

func (m *IPTagging_IPTag) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	l = len(m.IpTagName)
	if l > 0 {
		n += 1 + l + sovIpTagging(uint64(l))
	}
	if len(m.IpList) > 0 {
		for _, e := range m.IpList {
			l = e.Size()
			n += 1 + l + sovIpTagging(uint64(l))
		}
	}
	if m.XXX_unrecognized != nil {
		n += len(m.XXX_unrecognized)
	}
	return n
}

func sovIpTagging(x uint64) (n int) {
	for {
		n++
		x >>= 7
		if x == 0 {
			break
		}
	}
	return n
}
func sozIpTagging(x uint64) (n int) {
	return sovIpTagging(uint64((x << 1) ^ uint64((int64(x) >> 63))))
}
func (m *IPTagging) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowIpTagging
			}
			if iNdEx >= l {
				return io.ErrUnexpectedEOF
			}
			b := dAtA[iNdEx]
			iNdEx++
			wire |= uint64(b&0x7F) << shift
			if b < 0x80 {
				break
			}
		}
		fieldNum := int32(wire >> 3)
		wireType := int(wire & 0x7)
		if wireType == 4 {
			return fmt.Errorf("proto: IPTagging: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: IPTagging: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field RequestType", wireType)
			}
			m.RequestType = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowIpTagging
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.RequestType |= IPTagging_RequestType(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		case 4:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field IpTags", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowIpTagging
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				msglen |= int(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			if msglen < 0 {
				return ErrInvalidLengthIpTagging
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthIpTagging
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.IpTags = append(m.IpTags, &IPTagging_IPTag{})
			if err := m.IpTags[len(m.IpTags)-1].Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipIpTagging(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if skippy < 0 {
				return ErrInvalidLengthIpTagging
			}
			if (iNdEx + skippy) < 0 {
				return ErrInvalidLengthIpTagging
			}
			if (iNdEx + skippy) > l {
				return io.ErrUnexpectedEOF
			}
			m.XXX_unrecognized = append(m.XXX_unrecognized, dAtA[iNdEx:iNdEx+skippy]...)
			iNdEx += skippy
		}
	}

	if iNdEx > l {
		return io.ErrUnexpectedEOF
	}
	return nil
}
func (m *IPTagging_IPTag) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowIpTagging
			}
			if iNdEx >= l {
				return io.ErrUnexpectedEOF
			}
			b := dAtA[iNdEx]
			iNdEx++
			wire |= uint64(b&0x7F) << shift
			if b < 0x80 {
				break
			}
		}
		fieldNum := int32(wire >> 3)
		wireType := int(wire & 0x7)
		if wireType == 4 {
			return fmt.Errorf("proto: IPTag: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: IPTag: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field IpTagName", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowIpTagging
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				stringLen |= uint64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			intStringLen := int(stringLen)
			if intStringLen < 0 {
				return ErrInvalidLengthIpTagging
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthIpTagging
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.IpTagName = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 2:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field IpList", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowIpTagging
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				msglen |= int(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			if msglen < 0 {
				return ErrInvalidLengthIpTagging
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthIpTagging
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.IpList = append(m.IpList, &core.CidrRange{})
			if err := m.IpList[len(m.IpList)-1].Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipIpTagging(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if skippy < 0 {
				return ErrInvalidLengthIpTagging
			}
			if (iNdEx + skippy) < 0 {
				return ErrInvalidLengthIpTagging
			}
			if (iNdEx + skippy) > l {
				return io.ErrUnexpectedEOF
			}
			m.XXX_unrecognized = append(m.XXX_unrecognized, dAtA[iNdEx:iNdEx+skippy]...)
			iNdEx += skippy
		}
	}

	if iNdEx > l {
		return io.ErrUnexpectedEOF
	}
	return nil
}
func skipIpTagging(dAtA []byte) (n int, err error) {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return 0, ErrIntOverflowIpTagging
			}
			if iNdEx >= l {
				return 0, io.ErrUnexpectedEOF
			}
			b := dAtA[iNdEx]
			iNdEx++
			wire |= (uint64(b) & 0x7F) << shift
			if b < 0x80 {
				break
			}
		}
		wireType := int(wire & 0x7)
		switch wireType {
		case 0:
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return 0, ErrIntOverflowIpTagging
				}
				if iNdEx >= l {
					return 0, io.ErrUnexpectedEOF
				}
				iNdEx++
				if dAtA[iNdEx-1] < 0x80 {
					break
				}
			}
			return iNdEx, nil
		case 1:
			iNdEx += 8
			return iNdEx, nil
		case 2:
			var length int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return 0, ErrIntOverflowIpTagging
				}
				if iNdEx >= l {
					return 0, io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				length |= (int(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			if length < 0 {
				return 0, ErrInvalidLengthIpTagging
			}
			iNdEx += length
			if iNdEx < 0 {
				return 0, ErrInvalidLengthIpTagging
			}
			return iNdEx, nil
		case 3:
			for {
				var innerWire uint64
				var start int = iNdEx
				for shift := uint(0); ; shift += 7 {
					if shift >= 64 {
						return 0, ErrIntOverflowIpTagging
					}
					if iNdEx >= l {
						return 0, io.ErrUnexpectedEOF
					}
					b := dAtA[iNdEx]
					iNdEx++
					innerWire |= (uint64(b) & 0x7F) << shift
					if b < 0x80 {
						break
					}
				}
				innerWireType := int(innerWire & 0x7)
				if innerWireType == 4 {
					break
				}
				next, err := skipIpTagging(dAtA[start:])
				if err != nil {
					return 0, err
				}
				iNdEx = start + next
				if iNdEx < 0 {
					return 0, ErrInvalidLengthIpTagging
				}
			}
			return iNdEx, nil
		case 4:
			return iNdEx, nil
		case 5:
			iNdEx += 4
			return iNdEx, nil
		default:
			return 0, fmt.Errorf("proto: illegal wireType %d", wireType)
		}
	}
	panic("unreachable")
}

var (
	ErrInvalidLengthIpTagging = fmt.Errorf("proto: negative length found during unmarshaling")
	ErrIntOverflowIpTagging   = fmt.Errorf("proto: integer overflow")
)
