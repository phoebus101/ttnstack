// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: lorawan-stack/api/contact_info.proto

package ttnpb // import "go.thethings.network/lorawan-stack/pkg/ttnpb"

import proto "github.com/gogo/protobuf/proto"
import golang_proto "github.com/golang/protobuf/proto"
import fmt "fmt"
import math "math"
import _ "github.com/gogo/protobuf/gogoproto"
import types "github.com/gogo/protobuf/types"
import _ "google.golang.org/genproto/googleapis/api/annotations"

import time "time"

import strconv "strconv"

import context "context"
import grpc "google.golang.org/grpc"

import github_com_gogo_protobuf_types "github.com/gogo/protobuf/types"

import strings "strings"
import reflect "reflect"

import io "io"

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = golang_proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf
var _ = time.Kitchen

// This is a compile-time assertion to ensure that this generated file
// is compatible with the proto package it is being compiled against.
// A compilation error at this line likely means your copy of the
// proto package needs to be updated.
const _ = proto.GoGoProtoPackageIsVersion2 // please upgrade the proto package

type ContactType int32

const (
	CONTACT_TYPE_OTHER     ContactType = 0
	CONTACT_TYPE_ABUSE     ContactType = 1
	CONTACT_TYPE_BILLING   ContactType = 2
	CONTACT_TYPE_TECHNICAL ContactType = 3
)

var ContactType_name = map[int32]string{
	0: "CONTACT_TYPE_OTHER",
	1: "CONTACT_TYPE_ABUSE",
	2: "CONTACT_TYPE_BILLING",
	3: "CONTACT_TYPE_TECHNICAL",
}
var ContactType_value = map[string]int32{
	"CONTACT_TYPE_OTHER":     0,
	"CONTACT_TYPE_ABUSE":     1,
	"CONTACT_TYPE_BILLING":   2,
	"CONTACT_TYPE_TECHNICAL": 3,
}

func (ContactType) EnumDescriptor() ([]byte, []int) {
	return fileDescriptor_contact_info_e7da2f3dec1c0cf3, []int{0}
}

type ContactMethod int32

const (
	CONTACT_METHOD_OTHER ContactMethod = 0
	CONTACT_METHOD_EMAIL ContactMethod = 1
	CONTACT_METHOD_PHONE ContactMethod = 2
)

var ContactMethod_name = map[int32]string{
	0: "CONTACT_METHOD_OTHER",
	1: "CONTACT_METHOD_EMAIL",
	2: "CONTACT_METHOD_PHONE",
}
var ContactMethod_value = map[string]int32{
	"CONTACT_METHOD_OTHER": 0,
	"CONTACT_METHOD_EMAIL": 1,
	"CONTACT_METHOD_PHONE": 2,
}

func (ContactMethod) EnumDescriptor() ([]byte, []int) {
	return fileDescriptor_contact_info_e7da2f3dec1c0cf3, []int{1}
}

type ContactInfo struct {
	ContactType          ContactType   `protobuf:"varint,1,opt,name=contact_type,json=contactType,proto3,enum=ttn.lorawan.v3.ContactType" json:"contact_type,omitempty"`
	ContactMethod        ContactMethod `protobuf:"varint,2,opt,name=contact_method,json=contactMethod,proto3,enum=ttn.lorawan.v3.ContactMethod" json:"contact_method,omitempty"`
	Value                string        `protobuf:"bytes,3,opt,name=value,proto3" json:"value,omitempty"`
	Public               bool          `protobuf:"varint,4,opt,name=public,proto3" json:"public,omitempty"`
	ValidatedAt          *time.Time    `protobuf:"bytes,5,opt,name=validated_at,json=validatedAt,stdtime" json:"validated_at,omitempty"`
	XXX_NoUnkeyedLiteral struct{}      `json:"-"`
	XXX_sizecache        int32         `json:"-"`
}

func (m *ContactInfo) Reset()      { *m = ContactInfo{} }
func (*ContactInfo) ProtoMessage() {}
func (*ContactInfo) Descriptor() ([]byte, []int) {
	return fileDescriptor_contact_info_e7da2f3dec1c0cf3, []int{0}
}
func (m *ContactInfo) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *ContactInfo) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_ContactInfo.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalTo(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (dst *ContactInfo) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ContactInfo.Merge(dst, src)
}
func (m *ContactInfo) XXX_Size() int {
	return m.Size()
}
func (m *ContactInfo) XXX_DiscardUnknown() {
	xxx_messageInfo_ContactInfo.DiscardUnknown(m)
}

var xxx_messageInfo_ContactInfo proto.InternalMessageInfo

func (m *ContactInfo) GetContactType() ContactType {
	if m != nil {
		return m.ContactType
	}
	return CONTACT_TYPE_OTHER
}

func (m *ContactInfo) GetContactMethod() ContactMethod {
	if m != nil {
		return m.ContactMethod
	}
	return CONTACT_METHOD_OTHER
}

func (m *ContactInfo) GetValue() string {
	if m != nil {
		return m.Value
	}
	return ""
}

func (m *ContactInfo) GetPublic() bool {
	if m != nil {
		return m.Public
	}
	return false
}

func (m *ContactInfo) GetValidatedAt() *time.Time {
	if m != nil {
		return m.ValidatedAt
	}
	return nil
}

type ContactInfoValidation struct {
	Token                string   `protobuf:"bytes,1,opt,name=token,proto3" json:"token,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *ContactInfoValidation) Reset()      { *m = ContactInfoValidation{} }
func (*ContactInfoValidation) ProtoMessage() {}
func (*ContactInfoValidation) Descriptor() ([]byte, []int) {
	return fileDescriptor_contact_info_e7da2f3dec1c0cf3, []int{1}
}
func (m *ContactInfoValidation) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *ContactInfoValidation) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_ContactInfoValidation.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalTo(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (dst *ContactInfoValidation) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ContactInfoValidation.Merge(dst, src)
}
func (m *ContactInfoValidation) XXX_Size() int {
	return m.Size()
}
func (m *ContactInfoValidation) XXX_DiscardUnknown() {
	xxx_messageInfo_ContactInfoValidation.DiscardUnknown(m)
}

var xxx_messageInfo_ContactInfoValidation proto.InternalMessageInfo

func (m *ContactInfoValidation) GetToken() string {
	if m != nil {
		return m.Token
	}
	return ""
}

func init() {
	proto.RegisterType((*ContactInfo)(nil), "ttn.lorawan.v3.ContactInfo")
	golang_proto.RegisterType((*ContactInfo)(nil), "ttn.lorawan.v3.ContactInfo")
	proto.RegisterType((*ContactInfoValidation)(nil), "ttn.lorawan.v3.ContactInfoValidation")
	golang_proto.RegisterType((*ContactInfoValidation)(nil), "ttn.lorawan.v3.ContactInfoValidation")
	proto.RegisterEnum("ttn.lorawan.v3.ContactType", ContactType_name, ContactType_value)
	golang_proto.RegisterEnum("ttn.lorawan.v3.ContactType", ContactType_name, ContactType_value)
	proto.RegisterEnum("ttn.lorawan.v3.ContactMethod", ContactMethod_name, ContactMethod_value)
	golang_proto.RegisterEnum("ttn.lorawan.v3.ContactMethod", ContactMethod_name, ContactMethod_value)
}
func (x ContactType) String() string {
	s, ok := ContactType_name[int32(x)]
	if ok {
		return s
	}
	return strconv.Itoa(int(x))
}
func (x ContactMethod) String() string {
	s, ok := ContactMethod_name[int32(x)]
	if ok {
		return s
	}
	return strconv.Itoa(int(x))
}
func (this *ContactInfo) Equal(that interface{}) bool {
	if that == nil {
		return this == nil
	}

	that1, ok := that.(*ContactInfo)
	if !ok {
		that2, ok := that.(ContactInfo)
		if ok {
			that1 = &that2
		} else {
			return false
		}
	}
	if that1 == nil {
		return this == nil
	} else if this == nil {
		return false
	}
	if this.ContactType != that1.ContactType {
		return false
	}
	if this.ContactMethod != that1.ContactMethod {
		return false
	}
	if this.Value != that1.Value {
		return false
	}
	if this.Public != that1.Public {
		return false
	}
	if that1.ValidatedAt == nil {
		if this.ValidatedAt != nil {
			return false
		}
	} else if !this.ValidatedAt.Equal(*that1.ValidatedAt) {
		return false
	}
	return true
}
func (this *ContactInfoValidation) Equal(that interface{}) bool {
	if that == nil {
		return this == nil
	}

	that1, ok := that.(*ContactInfoValidation)
	if !ok {
		that2, ok := that.(ContactInfoValidation)
		if ok {
			that1 = &that2
		} else {
			return false
		}
	}
	if that1 == nil {
		return this == nil
	} else if this == nil {
		return false
	}
	if this.Token != that1.Token {
		return false
	}
	return true
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConn

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion4

// Client API for ContactInfoRegistry service

type ContactInfoRegistryClient interface {
	Validate(ctx context.Context, in *ContactInfoValidation, opts ...grpc.CallOption) (*types.Empty, error)
}

type contactInfoRegistryClient struct {
	cc *grpc.ClientConn
}

func NewContactInfoRegistryClient(cc *grpc.ClientConn) ContactInfoRegistryClient {
	return &contactInfoRegistryClient{cc}
}

func (c *contactInfoRegistryClient) Validate(ctx context.Context, in *ContactInfoValidation, opts ...grpc.CallOption) (*types.Empty, error) {
	out := new(types.Empty)
	err := c.cc.Invoke(ctx, "/ttn.lorawan.v3.ContactInfoRegistry/Validate", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// Server API for ContactInfoRegistry service

type ContactInfoRegistryServer interface {
	Validate(context.Context, *ContactInfoValidation) (*types.Empty, error)
}

func RegisterContactInfoRegistryServer(s *grpc.Server, srv ContactInfoRegistryServer) {
	s.RegisterService(&_ContactInfoRegistry_serviceDesc, srv)
}

func _ContactInfoRegistry_Validate_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ContactInfoValidation)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ContactInfoRegistryServer).Validate(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/ttn.lorawan.v3.ContactInfoRegistry/Validate",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ContactInfoRegistryServer).Validate(ctx, req.(*ContactInfoValidation))
	}
	return interceptor(ctx, in, info, handler)
}

var _ContactInfoRegistry_serviceDesc = grpc.ServiceDesc{
	ServiceName: "ttn.lorawan.v3.ContactInfoRegistry",
	HandlerType: (*ContactInfoRegistryServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "Validate",
			Handler:    _ContactInfoRegistry_Validate_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "lorawan-stack/api/contact_info.proto",
}

func (m *ContactInfo) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalTo(dAtA)
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *ContactInfo) MarshalTo(dAtA []byte) (int, error) {
	var i int
	_ = i
	var l int
	_ = l
	if m.ContactType != 0 {
		dAtA[i] = 0x8
		i++
		i = encodeVarintContactInfo(dAtA, i, uint64(m.ContactType))
	}
	if m.ContactMethod != 0 {
		dAtA[i] = 0x10
		i++
		i = encodeVarintContactInfo(dAtA, i, uint64(m.ContactMethod))
	}
	if len(m.Value) > 0 {
		dAtA[i] = 0x1a
		i++
		i = encodeVarintContactInfo(dAtA, i, uint64(len(m.Value)))
		i += copy(dAtA[i:], m.Value)
	}
	if m.Public {
		dAtA[i] = 0x20
		i++
		if m.Public {
			dAtA[i] = 1
		} else {
			dAtA[i] = 0
		}
		i++
	}
	if m.ValidatedAt != nil {
		dAtA[i] = 0x2a
		i++
		i = encodeVarintContactInfo(dAtA, i, uint64(github_com_gogo_protobuf_types.SizeOfStdTime(*m.ValidatedAt)))
		n1, err := github_com_gogo_protobuf_types.StdTimeMarshalTo(*m.ValidatedAt, dAtA[i:])
		if err != nil {
			return 0, err
		}
		i += n1
	}
	return i, nil
}

func (m *ContactInfoValidation) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalTo(dAtA)
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *ContactInfoValidation) MarshalTo(dAtA []byte) (int, error) {
	var i int
	_ = i
	var l int
	_ = l
	if len(m.Token) > 0 {
		dAtA[i] = 0xa
		i++
		i = encodeVarintContactInfo(dAtA, i, uint64(len(m.Token)))
		i += copy(dAtA[i:], m.Token)
	}
	return i, nil
}

func encodeVarintContactInfo(dAtA []byte, offset int, v uint64) int {
	for v >= 1<<7 {
		dAtA[offset] = uint8(v&0x7f | 0x80)
		v >>= 7
		offset++
	}
	dAtA[offset] = uint8(v)
	return offset + 1
}
func NewPopulatedContactInfo(r randyContactInfo, easy bool) *ContactInfo {
	this := &ContactInfo{}
	this.ContactType = ContactType([]int32{0, 1, 2, 3}[r.Intn(4)])
	this.ContactMethod = ContactMethod([]int32{0, 1, 2}[r.Intn(3)])
	this.Value = randStringContactInfo(r)
	this.Public = bool(r.Intn(2) == 0)
	if r.Intn(10) != 0 {
		this.ValidatedAt = github_com_gogo_protobuf_types.NewPopulatedStdTime(r, easy)
	}
	if !easy && r.Intn(10) != 0 {
	}
	return this
}

func NewPopulatedContactInfoValidation(r randyContactInfo, easy bool) *ContactInfoValidation {
	this := &ContactInfoValidation{}
	this.Token = randStringContactInfo(r)
	if !easy && r.Intn(10) != 0 {
	}
	return this
}

type randyContactInfo interface {
	Float32() float32
	Float64() float64
	Int63() int64
	Int31() int32
	Uint32() uint32
	Intn(n int) int
}

func randUTF8RuneContactInfo(r randyContactInfo) rune {
	ru := r.Intn(62)
	if ru < 10 {
		return rune(ru + 48)
	} else if ru < 36 {
		return rune(ru + 55)
	}
	return rune(ru + 61)
}
func randStringContactInfo(r randyContactInfo) string {
	v1 := r.Intn(100)
	tmps := make([]rune, v1)
	for i := 0; i < v1; i++ {
		tmps[i] = randUTF8RuneContactInfo(r)
	}
	return string(tmps)
}
func randUnrecognizedContactInfo(r randyContactInfo, maxFieldNumber int) (dAtA []byte) {
	l := r.Intn(5)
	for i := 0; i < l; i++ {
		wire := r.Intn(4)
		if wire == 3 {
			wire = 5
		}
		fieldNumber := maxFieldNumber + r.Intn(100)
		dAtA = randFieldContactInfo(dAtA, r, fieldNumber, wire)
	}
	return dAtA
}
func randFieldContactInfo(dAtA []byte, r randyContactInfo, fieldNumber int, wire int) []byte {
	key := uint32(fieldNumber)<<3 | uint32(wire)
	switch wire {
	case 0:
		dAtA = encodeVarintPopulateContactInfo(dAtA, uint64(key))
		v2 := r.Int63()
		if r.Intn(2) == 0 {
			v2 *= -1
		}
		dAtA = encodeVarintPopulateContactInfo(dAtA, uint64(v2))
	case 1:
		dAtA = encodeVarintPopulateContactInfo(dAtA, uint64(key))
		dAtA = append(dAtA, byte(r.Intn(256)), byte(r.Intn(256)), byte(r.Intn(256)), byte(r.Intn(256)), byte(r.Intn(256)), byte(r.Intn(256)), byte(r.Intn(256)), byte(r.Intn(256)))
	case 2:
		dAtA = encodeVarintPopulateContactInfo(dAtA, uint64(key))
		ll := r.Intn(100)
		dAtA = encodeVarintPopulateContactInfo(dAtA, uint64(ll))
		for j := 0; j < ll; j++ {
			dAtA = append(dAtA, byte(r.Intn(256)))
		}
	default:
		dAtA = encodeVarintPopulateContactInfo(dAtA, uint64(key))
		dAtA = append(dAtA, byte(r.Intn(256)), byte(r.Intn(256)), byte(r.Intn(256)), byte(r.Intn(256)))
	}
	return dAtA
}
func encodeVarintPopulateContactInfo(dAtA []byte, v uint64) []byte {
	for v >= 1<<7 {
		dAtA = append(dAtA, uint8(v&0x7f|0x80))
		v >>= 7
	}
	dAtA = append(dAtA, uint8(v))
	return dAtA
}
func (m *ContactInfo) Size() (n int) {
	var l int
	_ = l
	if m.ContactType != 0 {
		n += 1 + sovContactInfo(uint64(m.ContactType))
	}
	if m.ContactMethod != 0 {
		n += 1 + sovContactInfo(uint64(m.ContactMethod))
	}
	l = len(m.Value)
	if l > 0 {
		n += 1 + l + sovContactInfo(uint64(l))
	}
	if m.Public {
		n += 2
	}
	if m.ValidatedAt != nil {
		l = github_com_gogo_protobuf_types.SizeOfStdTime(*m.ValidatedAt)
		n += 1 + l + sovContactInfo(uint64(l))
	}
	return n
}

func (m *ContactInfoValidation) Size() (n int) {
	var l int
	_ = l
	l = len(m.Token)
	if l > 0 {
		n += 1 + l + sovContactInfo(uint64(l))
	}
	return n
}

func sovContactInfo(x uint64) (n int) {
	for {
		n++
		x >>= 7
		if x == 0 {
			break
		}
	}
	return n
}
func sozContactInfo(x uint64) (n int) {
	return sovContactInfo((x << 1) ^ uint64((int64(x) >> 63)))
}
func (this *ContactInfo) String() string {
	if this == nil {
		return "nil"
	}
	s := strings.Join([]string{`&ContactInfo{`,
		`ContactType:` + fmt.Sprintf("%v", this.ContactType) + `,`,
		`ContactMethod:` + fmt.Sprintf("%v", this.ContactMethod) + `,`,
		`Value:` + fmt.Sprintf("%v", this.Value) + `,`,
		`Public:` + fmt.Sprintf("%v", this.Public) + `,`,
		`ValidatedAt:` + strings.Replace(fmt.Sprintf("%v", this.ValidatedAt), "Timestamp", "types.Timestamp", 1) + `,`,
		`}`,
	}, "")
	return s
}
func (this *ContactInfoValidation) String() string {
	if this == nil {
		return "nil"
	}
	s := strings.Join([]string{`&ContactInfoValidation{`,
		`Token:` + fmt.Sprintf("%v", this.Token) + `,`,
		`}`,
	}, "")
	return s
}
func valueToStringContactInfo(v interface{}) string {
	rv := reflect.ValueOf(v)
	if rv.IsNil() {
		return "nil"
	}
	pv := reflect.Indirect(rv).Interface()
	return fmt.Sprintf("*%v", pv)
}
func (m *ContactInfo) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowContactInfo
			}
			if iNdEx >= l {
				return io.ErrUnexpectedEOF
			}
			b := dAtA[iNdEx]
			iNdEx++
			wire |= (uint64(b) & 0x7F) << shift
			if b < 0x80 {
				break
			}
		}
		fieldNum := int32(wire >> 3)
		wireType := int(wire & 0x7)
		if wireType == 4 {
			return fmt.Errorf("proto: ContactInfo: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: ContactInfo: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field ContactType", wireType)
			}
			m.ContactType = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowContactInfo
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.ContactType |= (ContactType(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		case 2:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field ContactMethod", wireType)
			}
			m.ContactMethod = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowContactInfo
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.ContactMethod |= (ContactMethod(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		case 3:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Value", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowContactInfo
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				stringLen |= (uint64(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			intStringLen := int(stringLen)
			if intStringLen < 0 {
				return ErrInvalidLengthContactInfo
			}
			postIndex := iNdEx + intStringLen
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Value = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 4:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field Public", wireType)
			}
			var v int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowContactInfo
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				v |= (int(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			m.Public = bool(v != 0)
		case 5:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field ValidatedAt", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowContactInfo
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				msglen |= (int(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			if msglen < 0 {
				return ErrInvalidLengthContactInfo
			}
			postIndex := iNdEx + msglen
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if m.ValidatedAt == nil {
				m.ValidatedAt = new(time.Time)
			}
			if err := github_com_gogo_protobuf_types.StdTimeUnmarshal(m.ValidatedAt, dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipContactInfo(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if skippy < 0 {
				return ErrInvalidLengthContactInfo
			}
			if (iNdEx + skippy) > l {
				return io.ErrUnexpectedEOF
			}
			iNdEx += skippy
		}
	}

	if iNdEx > l {
		return io.ErrUnexpectedEOF
	}
	return nil
}
func (m *ContactInfoValidation) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowContactInfo
			}
			if iNdEx >= l {
				return io.ErrUnexpectedEOF
			}
			b := dAtA[iNdEx]
			iNdEx++
			wire |= (uint64(b) & 0x7F) << shift
			if b < 0x80 {
				break
			}
		}
		fieldNum := int32(wire >> 3)
		wireType := int(wire & 0x7)
		if wireType == 4 {
			return fmt.Errorf("proto: ContactInfoValidation: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: ContactInfoValidation: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Token", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowContactInfo
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				stringLen |= (uint64(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			intStringLen := int(stringLen)
			if intStringLen < 0 {
				return ErrInvalidLengthContactInfo
			}
			postIndex := iNdEx + intStringLen
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Token = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipContactInfo(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if skippy < 0 {
				return ErrInvalidLengthContactInfo
			}
			if (iNdEx + skippy) > l {
				return io.ErrUnexpectedEOF
			}
			iNdEx += skippy
		}
	}

	if iNdEx > l {
		return io.ErrUnexpectedEOF
	}
	return nil
}
func skipContactInfo(dAtA []byte) (n int, err error) {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return 0, ErrIntOverflowContactInfo
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
					return 0, ErrIntOverflowContactInfo
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
					return 0, ErrIntOverflowContactInfo
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
			iNdEx += length
			if length < 0 {
				return 0, ErrInvalidLengthContactInfo
			}
			return iNdEx, nil
		case 3:
			for {
				var innerWire uint64
				var start int = iNdEx
				for shift := uint(0); ; shift += 7 {
					if shift >= 64 {
						return 0, ErrIntOverflowContactInfo
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
				next, err := skipContactInfo(dAtA[start:])
				if err != nil {
					return 0, err
				}
				iNdEx = start + next
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
	ErrInvalidLengthContactInfo = fmt.Errorf("proto: negative length found during unmarshaling")
	ErrIntOverflowContactInfo   = fmt.Errorf("proto: integer overflow")
)

func init() {
	proto.RegisterFile("lorawan-stack/api/contact_info.proto", fileDescriptor_contact_info_e7da2f3dec1c0cf3)
}
func init() {
	golang_proto.RegisterFile("lorawan-stack/api/contact_info.proto", fileDescriptor_contact_info_e7da2f3dec1c0cf3)
}

var fileDescriptor_contact_info_e7da2f3dec1c0cf3 = []byte{
	// 629 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x74, 0x53, 0x3f, 0x4c, 0xdb, 0x4e,
	0x14, 0xbe, 0x0b, 0x7f, 0x04, 0x17, 0x40, 0x91, 0x7f, 0xfc, 0x22, 0x2b, 0xb4, 0x8f, 0x08, 0xb5,
	0x52, 0x84, 0x8a, 0x2d, 0x85, 0xbd, 0x52, 0x12, 0xac, 0x26, 0x52, 0x48, 0x90, 0xeb, 0x56, 0x6a,
	0x97, 0xc8, 0x31, 0xc6, 0xb1, 0x92, 0xf8, 0xac, 0xe4, 0x02, 0xca, 0x86, 0x3a, 0x54, 0x8c, 0x48,
	0x5d, 0x3a, 0x56, 0xed, 0xc2, 0xc8, 0xc8, 0xc8, 0xc8, 0x88, 0xd4, 0x85, 0xad, 0xf8, 0xdc, 0x81,
	0x91, 0x91, 0xb1, 0xc2, 0x76, 0xd2, 0xb8, 0x11, 0x9b, 0xbf, 0xfb, 0xde, 0xfb, 0xee, 0x7d, 0xdf,
	0xf3, 0x91, 0x17, 0x1d, 0xda, 0xd3, 0x8f, 0x74, 0x67, 0xab, 0xcf, 0x74, 0xa3, 0x2d, 0xeb, 0xae,
	0x2d, 0x1b, 0xd4, 0x61, 0xba, 0xc1, 0x1a, 0xb6, 0x73, 0x40, 0x25, 0xb7, 0x47, 0x19, 0x15, 0x56,
	0x18, 0x73, 0xa4, 0xa8, 0x52, 0x3a, 0xdc, 0xce, 0x6c, 0x59, 0x36, 0x6b, 0x0d, 0x9a, 0x92, 0x41,
	0xbb, 0xb2, 0x45, 0x2d, 0x2a, 0x07, 0x65, 0xcd, 0xc1, 0x41, 0x80, 0x02, 0x10, 0x7c, 0x85, 0xed,
	0x99, 0x67, 0x16, 0xa5, 0x56, 0xc7, 0x0c, 0xd4, 0x75, 0xc7, 0xa1, 0x4c, 0x67, 0x36, 0x75, 0xfa,
	0x11, 0xbb, 0x16, 0xb1, 0x63, 0x0d, 0xb3, 0xeb, 0xb2, 0x61, 0x44, 0xae, 0xff, 0x4b, 0x32, 0xbb,
	0x6b, 0xf6, 0x99, 0xde, 0x75, 0xc3, 0x82, 0x8d, 0xcf, 0x09, 0x92, 0x2c, 0x85, 0x13, 0x57, 0x9c,
	0x03, 0x2a, 0xbc, 0x26, 0x4b, 0x23, 0x03, 0x6c, 0xe8, 0x9a, 0x22, 0xce, 0xe2, 0xdc, 0x4a, 0x7e,
	0x4d, 0x8a, 0x3b, 0x90, 0xa2, 0x16, 0x6d, 0xe8, 0x9a, 0x6a, 0xd2, 0xf8, 0x0b, 0x84, 0x1d, 0xb2,
	0x32, 0xea, 0xef, 0x9a, 0xac, 0x45, 0xf7, 0xc5, 0x44, 0xa0, 0xf0, 0xfc, 0x09, 0x85, 0xdd, 0xa0,
	0x48, 0x5d, 0x36, 0x26, 0xa1, 0xb0, 0x4a, 0xe6, 0x0e, 0xf5, 0xce, 0xc0, 0x14, 0x67, 0xb2, 0x38,
	0xb7, 0xa8, 0x86, 0x40, 0x48, 0x93, 0x79, 0x77, 0xd0, 0xec, 0xd8, 0x86, 0x38, 0x9b, 0xc5, 0xb9,
	0x05, 0x35, 0x42, 0x42, 0x89, 0x2c, 0x1d, 0xea, 0x1d, 0x7b, 0x5f, 0x67, 0xe6, 0x7e, 0x43, 0x67,
	0xe2, 0x5c, 0x16, 0xe7, 0x92, 0xf9, 0x8c, 0x14, 0x7a, 0x97, 0x46, 0xde, 0x25, 0x6d, 0xe4, 0xbd,
	0x38, 0x7b, 0xfa, 0x6b, 0x1d, 0xab, 0xc9, 0x71, 0x57, 0x81, 0x6d, 0x6c, 0x91, 0xff, 0x27, 0x72,
	0x78, 0x1f, 0x32, 0x36, 0x75, 0x1e, 0x67, 0x61, 0xb4, 0x6d, 0x3a, 0x41, 0x14, 0x8b, 0x6a, 0x08,
	0x36, 0x87, 0xe3, 0xd8, 0x02, 0xdb, 0x69, 0x22, 0x94, 0xea, 0x35, 0xad, 0x50, 0xd2, 0x1a, 0xda,
	0x87, 0x3d, 0xa5, 0x51, 0xd7, 0xca, 0x8a, 0x9a, 0x42, 0x53, 0xe7, 0x85, 0xe2, 0xbb, 0xb7, 0x4a,
	0x0a, 0x0b, 0x22, 0x59, 0x8d, 0x9d, 0x17, 0x2b, 0xd5, 0x6a, 0xa5, 0xf6, 0x26, 0x95, 0x10, 0x32,
	0x24, 0x1d, 0x63, 0x34, 0xa5, 0x54, 0xae, 0x55, 0x4a, 0x85, 0x6a, 0x6a, 0x26, 0x33, 0x7b, 0xf2,
	0x03, 0xd0, 0xa6, 0x41, 0x96, 0x63, 0xe1, 0x4d, 0x8a, 0xed, 0x2a, 0x5a, 0xb9, 0xbe, 0x33, 0xbe,
	0x7e, 0x9a, 0x51, 0x76, 0x0b, 0x95, 0x6a, 0x7c, 0x80, 0x88, 0xd9, 0x2b, 0xd7, 0x6b, 0x4a, 0x2a,
	0x11, 0x5e, 0x92, 0x3f, 0xc6, 0xe4, 0xbf, 0x89, 0x3c, 0x54, 0xd3, 0xb2, 0xfb, 0xac, 0x37, 0x14,
	0x6c, 0xb2, 0x10, 0x65, 0x63, 0x0a, 0x2f, 0x9f, 0xd8, 0x69, 0x3c, 0xc0, 0x4c, 0x7a, 0x6a, 0x11,
	0xca, 0xe3, 0x1f, 0xba, 0x01, 0x9f, 0x7e, 0xfe, 0xfe, 0x92, 0x10, 0xf3, 0xe9, 0xd8, 0x93, 0x91,
	0x47, 0x4b, 0x29, 0x7e, 0xc7, 0x57, 0x1e, 0xe0, 0x6b, 0x0f, 0xf0, 0x8d, 0x07, 0xe8, 0xd6, 0x03,
	0x74, 0xe7, 0x01, 0xba, 0xf7, 0x00, 0x3d, 0x78, 0x80, 0x8f, 0x39, 0xe0, 0x13, 0x0e, 0xe8, 0x8c,
	0x03, 0x3e, 0xe7, 0x80, 0x2e, 0x38, 0xa0, 0x4b, 0x0e, 0xe8, 0x8a, 0x03, 0xbe, 0xe6, 0x80, 0x6f,
	0x38, 0xa0, 0x5b, 0x0e, 0xf8, 0x8e, 0x03, 0xba, 0xe7, 0x80, 0x1f, 0x38, 0xa0, 0x63, 0x1f, 0xd0,
	0x89, 0x0f, 0xf8, 0xd4, 0x07, 0xf4, 0xd5, 0x07, 0xfc, 0xcd, 0x07, 0x74, 0xe6, 0x03, 0x3a, 0xf7,
	0x01, 0x5f, 0xf8, 0x80, 0x2f, 0x7d, 0xc0, 0x1f, 0x5f, 0x59, 0x54, 0x62, 0x2d, 0x93, 0xb5, 0x6c,
	0xc7, 0xea, 0x4b, 0x8e, 0xc9, 0x8e, 0x68, 0xaf, 0x2d, 0xc7, 0xdf, 0xb9, 0xdb, 0xb6, 0x64, 0xc6,
	0x1c, 0xb7, 0xd9, 0x9c, 0x0f, 0x4c, 0x6d, 0xff, 0x09, 0x00, 0x00, 0xff, 0xff, 0xbe, 0x4b, 0xf6,
	0x12, 0x09, 0x04, 0x00, 0x00,
}