// Code generated by protoc-gen-gogo.
// source: combos/neither/types.proto
// DO NOT EDIT!

/*
Package types is a generated protocol buffer package.

It is generated from these files:
	combos/neither/types.proto

It has these top-level messages:
	KnownTypes
*/
package types

import proto "github.com/gogo/protobuf/proto"
import fmt "fmt"
import math "math"
import _ "github.com/gogo/protobuf/gogoproto"
import google_protobuf1 "github.com/gogo/protobuf/types"
import google_protobuf2 "github.com/gogo/protobuf/types"
import google_protobuf3 "github.com/gogo/protobuf/types"

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// This is a compile-time assertion to ensure that this generated file
// is compatible with the proto package it is being compiled against.
// A compilation error at this line likely means your copy of the
// proto package needs to be updated.
const _ = proto.GoGoProtoPackageIsVersion2 // please upgrade the proto package

type KnownTypes struct {
	// google.protobuf.Any an = 14;
	Dur *google_protobuf1.Duration `protobuf:"bytes,1,opt,name=dur" json:"dur,omitempty"`
	// google.protobuf.Struct st = 12;
	Ts    *google_protobuf2.Timestamp   `protobuf:"bytes,2,opt,name=ts" json:"ts,omitempty"`
	Dbl   *google_protobuf3.DoubleValue `protobuf:"bytes,3,opt,name=dbl" json:"dbl,omitempty"`
	Flt   *google_protobuf3.FloatValue  `protobuf:"bytes,4,opt,name=flt" json:"flt,omitempty"`
	I64   *google_protobuf3.Int64Value  `protobuf:"bytes,5,opt,name=i64" json:"i64,omitempty"`
	U64   *google_protobuf3.UInt64Value `protobuf:"bytes,6,opt,name=u64" json:"u64,omitempty"`
	I32   *google_protobuf3.Int32Value  `protobuf:"bytes,7,opt,name=i32" json:"i32,omitempty"`
	U32   *google_protobuf3.UInt32Value `protobuf:"bytes,8,opt,name=u32" json:"u32,omitempty"`
	Bool  *google_protobuf3.BoolValue   `protobuf:"bytes,9,opt,name=bool" json:"bool,omitempty"`
	Str   *google_protobuf3.StringValue `protobuf:"bytes,10,opt,name=str" json:"str,omitempty"`
	Bytes *google_protobuf3.BytesValue  `protobuf:"bytes,11,opt,name=bytes" json:"bytes,omitempty"`
}

func (m *KnownTypes) Reset()                    { *m = KnownTypes{} }
func (m *KnownTypes) String() string            { return proto.CompactTextString(m) }
func (*KnownTypes) ProtoMessage()               {}
func (*KnownTypes) Descriptor() ([]byte, []int) { return fileDescriptorTypes, []int{0} }

func (m *KnownTypes) GetDur() *google_protobuf1.Duration {
	if m != nil {
		return m.Dur
	}
	return nil
}

func (m *KnownTypes) GetTs() *google_protobuf2.Timestamp {
	if m != nil {
		return m.Ts
	}
	return nil
}

func (m *KnownTypes) GetDbl() *google_protobuf3.DoubleValue {
	if m != nil {
		return m.Dbl
	}
	return nil
}

func (m *KnownTypes) GetFlt() *google_protobuf3.FloatValue {
	if m != nil {
		return m.Flt
	}
	return nil
}

func (m *KnownTypes) GetI64() *google_protobuf3.Int64Value {
	if m != nil {
		return m.I64
	}
	return nil
}

func (m *KnownTypes) GetU64() *google_protobuf3.UInt64Value {
	if m != nil {
		return m.U64
	}
	return nil
}

func (m *KnownTypes) GetI32() *google_protobuf3.Int32Value {
	if m != nil {
		return m.I32
	}
	return nil
}

func (m *KnownTypes) GetU32() *google_protobuf3.UInt32Value {
	if m != nil {
		return m.U32
	}
	return nil
}

func (m *KnownTypes) GetBool() *google_protobuf3.BoolValue {
	if m != nil {
		return m.Bool
	}
	return nil
}

func (m *KnownTypes) GetStr() *google_protobuf3.StringValue {
	if m != nil {
		return m.Str
	}
	return nil
}

func (m *KnownTypes) GetBytes() *google_protobuf3.BytesValue {
	if m != nil {
		return m.Bytes
	}
	return nil
}

func init() {
	proto.RegisterType((*KnownTypes)(nil), "types.KnownTypes")
}
func (this *KnownTypes) Equal(that interface{}) bool {
	if that == nil {
		if this == nil {
			return true
		}
		return false
	}

	that1, ok := that.(*KnownTypes)
	if !ok {
		that2, ok := that.(KnownTypes)
		if ok {
			that1 = &that2
		} else {
			return false
		}
	}
	if that1 == nil {
		if this == nil {
			return true
		}
		return false
	} else if this == nil {
		return false
	}
	if !this.Dur.Equal(that1.Dur) {
		return false
	}
	if !this.Ts.Equal(that1.Ts) {
		return false
	}
	if !this.Dbl.Equal(that1.Dbl) {
		return false
	}
	if !this.Flt.Equal(that1.Flt) {
		return false
	}
	if !this.I64.Equal(that1.I64) {
		return false
	}
	if !this.U64.Equal(that1.U64) {
		return false
	}
	if !this.I32.Equal(that1.I32) {
		return false
	}
	if !this.U32.Equal(that1.U32) {
		return false
	}
	if !this.Bool.Equal(that1.Bool) {
		return false
	}
	if !this.Str.Equal(that1.Str) {
		return false
	}
	if !this.Bytes.Equal(that1.Bytes) {
		return false
	}
	return true
}
func NewPopulatedKnownTypes(r randyTypes, easy bool) *KnownTypes {
	this := &KnownTypes{}
	if r.Intn(10) != 0 {
		this.Dur = google_protobuf1.NewPopulatedDuration(r, easy)
	}
	if r.Intn(10) != 0 {
		this.Ts = google_protobuf2.NewPopulatedTimestamp(r, easy)
	}
	if r.Intn(10) != 0 {
		this.Dbl = google_protobuf3.NewPopulatedDoubleValue(r, easy)
	}
	if r.Intn(10) != 0 {
		this.Flt = google_protobuf3.NewPopulatedFloatValue(r, easy)
	}
	if r.Intn(10) != 0 {
		this.I64 = google_protobuf3.NewPopulatedInt64Value(r, easy)
	}
	if r.Intn(10) != 0 {
		this.U64 = google_protobuf3.NewPopulatedUInt64Value(r, easy)
	}
	if r.Intn(10) != 0 {
		this.I32 = google_protobuf3.NewPopulatedInt32Value(r, easy)
	}
	if r.Intn(10) != 0 {
		this.U32 = google_protobuf3.NewPopulatedUInt32Value(r, easy)
	}
	if r.Intn(10) != 0 {
		this.Bool = google_protobuf3.NewPopulatedBoolValue(r, easy)
	}
	if r.Intn(10) != 0 {
		this.Str = google_protobuf3.NewPopulatedStringValue(r, easy)
	}
	if r.Intn(10) != 0 {
		this.Bytes = google_protobuf3.NewPopulatedBytesValue(r, easy)
	}
	if !easy && r.Intn(10) != 0 {
	}
	return this
}

type randyTypes interface {
	Float32() float32
	Float64() float64
	Int63() int64
	Int31() int32
	Uint32() uint32
	Intn(n int) int
}

func randUTF8RuneTypes(r randyTypes) rune {
	ru := r.Intn(62)
	if ru < 10 {
		return rune(ru + 48)
	} else if ru < 36 {
		return rune(ru + 55)
	}
	return rune(ru + 61)
}
func randStringTypes(r randyTypes) string {
	v1 := r.Intn(100)
	tmps := make([]rune, v1)
	for i := 0; i < v1; i++ {
		tmps[i] = randUTF8RuneTypes(r)
	}
	return string(tmps)
}
func randUnrecognizedTypes(r randyTypes, maxFieldNumber int) (data []byte) {
	l := r.Intn(5)
	for i := 0; i < l; i++ {
		wire := r.Intn(4)
		if wire == 3 {
			wire = 5
		}
		fieldNumber := maxFieldNumber + r.Intn(100)
		data = randFieldTypes(data, r, fieldNumber, wire)
	}
	return data
}
func randFieldTypes(data []byte, r randyTypes, fieldNumber int, wire int) []byte {
	key := uint32(fieldNumber)<<3 | uint32(wire)
	switch wire {
	case 0:
		data = encodeVarintPopulateTypes(data, uint64(key))
		v2 := r.Int63()
		if r.Intn(2) == 0 {
			v2 *= -1
		}
		data = encodeVarintPopulateTypes(data, uint64(v2))
	case 1:
		data = encodeVarintPopulateTypes(data, uint64(key))
		data = append(data, byte(r.Intn(256)), byte(r.Intn(256)), byte(r.Intn(256)), byte(r.Intn(256)), byte(r.Intn(256)), byte(r.Intn(256)), byte(r.Intn(256)), byte(r.Intn(256)))
	case 2:
		data = encodeVarintPopulateTypes(data, uint64(key))
		ll := r.Intn(100)
		data = encodeVarintPopulateTypes(data, uint64(ll))
		for j := 0; j < ll; j++ {
			data = append(data, byte(r.Intn(256)))
		}
	default:
		data = encodeVarintPopulateTypes(data, uint64(key))
		data = append(data, byte(r.Intn(256)), byte(r.Intn(256)), byte(r.Intn(256)), byte(r.Intn(256)))
	}
	return data
}
func encodeVarintPopulateTypes(data []byte, v uint64) []byte {
	for v >= 1<<7 {
		data = append(data, uint8(uint64(v)&0x7f|0x80))
		v >>= 7
	}
	data = append(data, uint8(v))
	return data
}
func (m *KnownTypes) Size() (n int) {
	var l int
	_ = l
	if m.Dur != nil {
		l = m.Dur.Size()
		n += 1 + l + sovTypes(uint64(l))
	}
	if m.Ts != nil {
		l = m.Ts.Size()
		n += 1 + l + sovTypes(uint64(l))
	}
	if m.Dbl != nil {
		l = m.Dbl.Size()
		n += 1 + l + sovTypes(uint64(l))
	}
	if m.Flt != nil {
		l = m.Flt.Size()
		n += 1 + l + sovTypes(uint64(l))
	}
	if m.I64 != nil {
		l = m.I64.Size()
		n += 1 + l + sovTypes(uint64(l))
	}
	if m.U64 != nil {
		l = m.U64.Size()
		n += 1 + l + sovTypes(uint64(l))
	}
	if m.I32 != nil {
		l = m.I32.Size()
		n += 1 + l + sovTypes(uint64(l))
	}
	if m.U32 != nil {
		l = m.U32.Size()
		n += 1 + l + sovTypes(uint64(l))
	}
	if m.Bool != nil {
		l = m.Bool.Size()
		n += 1 + l + sovTypes(uint64(l))
	}
	if m.Str != nil {
		l = m.Str.Size()
		n += 1 + l + sovTypes(uint64(l))
	}
	if m.Bytes != nil {
		l = m.Bytes.Size()
		n += 1 + l + sovTypes(uint64(l))
	}
	return n
}

func sovTypes(x uint64) (n int) {
	for {
		n++
		x >>= 7
		if x == 0 {
			break
		}
	}
	return n
}
func sozTypes(x uint64) (n int) {
	return sovTypes(uint64((x << 1) ^ uint64((int64(x) >> 63))))
}

func init() { proto.RegisterFile("combos/neither/types.proto", fileDescriptorTypes) }

var fileDescriptorTypes = []byte{
	// 384 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x09, 0x6e, 0x88, 0x02, 0xff, 0x7c, 0x90, 0x41, 0xce, 0x93, 0x40,
	0x18, 0x86, 0xa1, 0xb4, 0x55, 0xa7, 0x3b, 0x56, 0x23, 0x9a, 0xd1, 0x18, 0x17, 0x46, 0x53, 0x88,
	0x40, 0x38, 0x40, 0x63, 0x4c, 0x8c, 0xbb, 0x5a, 0xdd, 0x33, 0xed, 0x94, 0x92, 0x0c, 0x7c, 0x64,
	0xe6, 0x9b, 0x34, 0xdd, 0x79, 0x09, 0xef, 0xe0, 0x11, 0x5c, 0xba, 0x74, 0xe9, 0x11, 0x14, 0x2f,
	0xd1, 0xa5, 0x61, 0xa0, 0xfa, 0xe7, 0x6f, 0xf8, 0x77, 0x7c, 0x79, 0x9f, 0xf7, 0xe1, 0x05, 0x12,
	0x6c, 0xa1, 0xe2, 0xa0, 0xa3, 0x5a, 0x94, 0x78, 0x10, 0x2a, 0xc2, 0x53, 0x23, 0x74, 0xd8, 0x28,
	0x40, 0xf0, 0x67, 0xf6, 0x08, 0x96, 0x45, 0x89, 0x07, 0xc3, 0xc3, 0x2d, 0x54, 0x51, 0x01, 0x05,
	0x44, 0x36, 0xe5, 0x66, 0x6f, 0x2f, 0x7b, 0xd8, 0xa7, 0xbe, 0x15, 0xb0, 0x02, 0xa0, 0x90, 0xe2,
	0x3f, 0xb5, 0x33, 0x2a, 0xc7, 0x12, 0xea, 0x21, 0x7f, 0x72, 0x3b, 0xc7, 0xb2, 0x12, 0x1a, 0xf3,
	0xaa, 0x19, 0x13, 0x1c, 0x55, 0xde, 0x34, 0x42, 0x0d, 0xb3, 0x9e, 0x7d, 0x99, 0x12, 0xf2, 0xbe,
	0x86, 0x63, 0xbd, 0xe9, 0xe6, 0xf9, 0xaf, 0x88, 0xb7, 0x33, 0x8a, 0xba, 0x4f, 0xdd, 0x17, 0x8b,
	0xf8, 0x61, 0xd8, 0x97, 0xc3, 0x4b, 0x39, 0x7c, 0x33, 0xbc, 0x7d, 0xdd, 0x51, 0xfe, 0x4b, 0x32,
	0x41, 0x4d, 0x27, 0x96, 0x0d, 0xae, 0xd8, 0xcd, 0x65, 0xc9, 0x7a, 0x82, 0xda, 0x0f, 0x89, 0xb7,
	0xe3, 0x92, 0x7a, 0x16, 0x7e, 0x7c, 0x2d, 0x06, 0xc3, 0xa5, 0xf8, 0x94, 0x4b, 0x23, 0xd6, 0x1d,
	0xe8, 0x2f, 0x89, 0xb7, 0x97, 0x48, 0xa7, 0x96, 0x7f, 0x74, 0xc5, 0xbf, 0x95, 0x90, 0xe3, 0x80,
	0xef, 0x25, 0x76, 0x78, 0x99, 0xa5, 0x74, 0x36, 0x82, 0xbf, 0xab, 0x31, 0x4b, 0x07, 0xbc, 0xcc,
	0xd2, 0x6e, 0x8d, 0xc9, 0x52, 0x3a, 0x1f, 0x59, 0xf3, 0xf1, 0x26, 0x6f, 0xb2, 0xd4, 0xea, 0x93,
	0x98, 0xde, 0x1b, 0xd7, 0x27, 0xf1, 0x45, 0x9f, 0xc4, 0x56, 0x9f, 0xc4, 0xf4, 0xfe, 0x1d, 0xfa,
	0x7f, 0xbc, 0xb1, 0xfc, 0x94, 0x03, 0x48, 0xfa, 0x60, 0xe4, 0x57, 0xae, 0x00, 0x64, 0x8f, 0x5b,
	0xae, 0xf3, 0x6b, 0x54, 0x94, 0x8c, 0xf8, 0x3f, 0xa0, 0x2a, 0xeb, 0x62, 0xf0, 0x6b, 0x54, 0xfe,
	0x6b, 0x32, 0xe3, 0x27, 0x14, 0x9a, 0x2e, 0x46, 0x3e, 0x60, 0xd5, 0xa5, 0x7d, 0xa1, 0x27, 0x57,
	0xcf, 0xcf, 0xbf, 0x99, 0xfb, 0xb5, 0x65, 0xee, 0xb7, 0x96, 0xb9, 0xdf, 0x5b, 0xe6, 0xfe, 0x68,
	0x99, 0xf3, 0xb3, 0x65, 0xce, 0xaf, 0x96, 0xb9, 0xe7, 0x96, 0x39, 0x9f, 0xff, 0x30, 0x87, 0xcf,
	0xad, 0x21, 0xf9, 0x1b, 0x00, 0x00, 0xff, 0xff, 0x22, 0x75, 0x1d, 0x1f, 0xf9, 0x02, 0x00, 0x00,
}
