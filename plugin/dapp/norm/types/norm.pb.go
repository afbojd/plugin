// Code generated by protoc-gen-go. DO NOT EDIT.
// source: norm.proto

package types

import (
	fmt "fmt"
	proto "github.com/golang/protobuf/proto"
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
const _ = proto.ProtoPackageIsVersion2 // please upgrade the proto package

type Norm struct {
	NormId               []byte   `protobuf:"bytes,1,opt,name=normId,proto3" json:"normId,omitempty"`
	CreateTime           int64    `protobuf:"varint,2,opt,name=createTime,proto3" json:"createTime,omitempty"`
	Key                  string   `protobuf:"bytes,3,opt,name=key,proto3" json:"key,omitempty"`
	Value                []byte   `protobuf:"bytes,4,opt,name=value,proto3" json:"value,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *Norm) Reset()         { *m = Norm{} }
func (m *Norm) String() string { return proto.CompactTextString(m) }
func (*Norm) ProtoMessage()    {}
func (*Norm) Descriptor() ([]byte, []int) {
	return fileDescriptor_4c505000707f0293, []int{0}
}

func (m *Norm) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Norm.Unmarshal(m, b)
}
func (m *Norm) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Norm.Marshal(b, m, deterministic)
}
func (m *Norm) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Norm.Merge(m, src)
}
func (m *Norm) XXX_Size() int {
	return xxx_messageInfo_Norm.Size(m)
}
func (m *Norm) XXX_DiscardUnknown() {
	xxx_messageInfo_Norm.DiscardUnknown(m)
}

var xxx_messageInfo_Norm proto.InternalMessageInfo

func (m *Norm) GetNormId() []byte {
	if m != nil {
		return m.NormId
	}
	return nil
}

func (m *Norm) GetCreateTime() int64 {
	if m != nil {
		return m.CreateTime
	}
	return 0
}

func (m *Norm) GetKey() string {
	if m != nil {
		return m.Key
	}
	return ""
}

func (m *Norm) GetValue() []byte {
	if m != nil {
		return m.Value
	}
	return nil
}

type NormAction struct {
	// Types that are valid to be assigned to Value:
	//	*NormAction_Nput
	Value                isNormAction_Value `protobuf_oneof:"value"`
	Ty                   int32              `protobuf:"varint,5,opt,name=ty,proto3" json:"ty,omitempty"`
	XXX_NoUnkeyedLiteral struct{}           `json:"-"`
	XXX_unrecognized     []byte             `json:"-"`
	XXX_sizecache        int32              `json:"-"`
}

func (m *NormAction) Reset()         { *m = NormAction{} }
func (m *NormAction) String() string { return proto.CompactTextString(m) }
func (*NormAction) ProtoMessage()    {}
func (*NormAction) Descriptor() ([]byte, []int) {
	return fileDescriptor_4c505000707f0293, []int{1}
}

func (m *NormAction) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_NormAction.Unmarshal(m, b)
}
func (m *NormAction) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_NormAction.Marshal(b, m, deterministic)
}
func (m *NormAction) XXX_Merge(src proto.Message) {
	xxx_messageInfo_NormAction.Merge(m, src)
}
func (m *NormAction) XXX_Size() int {
	return xxx_messageInfo_NormAction.Size(m)
}
func (m *NormAction) XXX_DiscardUnknown() {
	xxx_messageInfo_NormAction.DiscardUnknown(m)
}

var xxx_messageInfo_NormAction proto.InternalMessageInfo

type isNormAction_Value interface {
	isNormAction_Value()
}

type NormAction_Nput struct {
	Nput *NormPut `protobuf:"bytes,1,opt,name=nput,proto3,oneof"`
}

func (*NormAction_Nput) isNormAction_Value() {}

func (m *NormAction) GetValue() isNormAction_Value {
	if m != nil {
		return m.Value
	}
	return nil
}

func (m *NormAction) GetNput() *NormPut {
	if x, ok := m.GetValue().(*NormAction_Nput); ok {
		return x.Nput
	}
	return nil
}

func (m *NormAction) GetTy() int32 {
	if m != nil {
		return m.Ty
	}
	return 0
}

// XXX_OneofFuncs is for the internal use of the proto package.
func (*NormAction) XXX_OneofFuncs() (func(msg proto.Message, b *proto.Buffer) error, func(msg proto.Message, tag, wire int, b *proto.Buffer) (bool, error), func(msg proto.Message) (n int), []interface{}) {
	return _NormAction_OneofMarshaler, _NormAction_OneofUnmarshaler, _NormAction_OneofSizer, []interface{}{
		(*NormAction_Nput)(nil),
	}
}

func _NormAction_OneofMarshaler(msg proto.Message, b *proto.Buffer) error {
	m := msg.(*NormAction)
	// value
	switch x := m.Value.(type) {
	case *NormAction_Nput:
		b.EncodeVarint(1<<3 | proto.WireBytes)
		if err := b.EncodeMessage(x.Nput); err != nil {
			return err
		}
	case nil:
	default:
		return fmt.Errorf("NormAction.Value has unexpected type %T", x)
	}
	return nil
}

func _NormAction_OneofUnmarshaler(msg proto.Message, tag, wire int, b *proto.Buffer) (bool, error) {
	m := msg.(*NormAction)
	switch tag {
	case 1: // value.nput
		if wire != proto.WireBytes {
			return true, proto.ErrInternalBadWireType
		}
		msg := new(NormPut)
		err := b.DecodeMessage(msg)
		m.Value = &NormAction_Nput{msg}
		return true, err
	default:
		return false, nil
	}
}

func _NormAction_OneofSizer(msg proto.Message) (n int) {
	m := msg.(*NormAction)
	// value
	switch x := m.Value.(type) {
	case *NormAction_Nput:
		s := proto.Size(x.Nput)
		n += 1 // tag and wire
		n += proto.SizeVarint(uint64(s))
		n += s
	case nil:
	default:
		panic(fmt.Sprintf("proto: unexpected type %T in oneof", x))
	}
	return n
}

type NormPut struct {
	Key                  string   `protobuf:"bytes,1,opt,name=key,proto3" json:"key,omitempty"`
	Value                []byte   `protobuf:"bytes,2,opt,name=value,proto3" json:"value,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *NormPut) Reset()         { *m = NormPut{} }
func (m *NormPut) String() string { return proto.CompactTextString(m) }
func (*NormPut) ProtoMessage()    {}
func (*NormPut) Descriptor() ([]byte, []int) {
	return fileDescriptor_4c505000707f0293, []int{2}
}

func (m *NormPut) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_NormPut.Unmarshal(m, b)
}
func (m *NormPut) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_NormPut.Marshal(b, m, deterministic)
}
func (m *NormPut) XXX_Merge(src proto.Message) {
	xxx_messageInfo_NormPut.Merge(m, src)
}
func (m *NormPut) XXX_Size() int {
	return xxx_messageInfo_NormPut.Size(m)
}
func (m *NormPut) XXX_DiscardUnknown() {
	xxx_messageInfo_NormPut.DiscardUnknown(m)
}

var xxx_messageInfo_NormPut proto.InternalMessageInfo

func (m *NormPut) GetKey() string {
	if m != nil {
		return m.Key
	}
	return ""
}

func (m *NormPut) GetValue() []byte {
	if m != nil {
		return m.Value
	}
	return nil
}

type NormGetKey struct {
	Key                  string   `protobuf:"bytes,1,opt,name=key,proto3" json:"key,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *NormGetKey) Reset()         { *m = NormGetKey{} }
func (m *NormGetKey) String() string { return proto.CompactTextString(m) }
func (*NormGetKey) ProtoMessage()    {}
func (*NormGetKey) Descriptor() ([]byte, []int) {
	return fileDescriptor_4c505000707f0293, []int{3}
}

func (m *NormGetKey) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_NormGetKey.Unmarshal(m, b)
}
func (m *NormGetKey) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_NormGetKey.Marshal(b, m, deterministic)
}
func (m *NormGetKey) XXX_Merge(src proto.Message) {
	xxx_messageInfo_NormGetKey.Merge(m, src)
}
func (m *NormGetKey) XXX_Size() int {
	return xxx_messageInfo_NormGetKey.Size(m)
}
func (m *NormGetKey) XXX_DiscardUnknown() {
	xxx_messageInfo_NormGetKey.DiscardUnknown(m)
}

var xxx_messageInfo_NormGetKey proto.InternalMessageInfo

func (m *NormGetKey) GetKey() string {
	if m != nil {
		return m.Key
	}
	return ""
}

func init() {
	proto.RegisterType((*Norm)(nil), "types.Norm")
	proto.RegisterType((*NormAction)(nil), "types.NormAction")
	proto.RegisterType((*NormPut)(nil), "types.NormPut")
	proto.RegisterType((*NormGetKey)(nil), "types.NormGetKey")
}

func init() { proto.RegisterFile("norm.proto", fileDescriptor_4c505000707f0293) }

var fileDescriptor_4c505000707f0293 = []byte{
	// 211 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x6c, 0x90, 0xc1, 0x4a, 0xc6, 0x30,
	0x10, 0x84, 0x4d, 0xda, 0xfc, 0xc5, 0x55, 0x8a, 0x04, 0x91, 0x9c, 0x4a, 0x28, 0x1e, 0x72, 0x2a,
	0xa8, 0x4f, 0xa0, 0x17, 0x95, 0x82, 0x48, 0xf0, 0x05, 0x6a, 0x5d, 0xa1, 0x68, 0x9b, 0x12, 0xb7,
	0x42, 0xde, 0x5e, 0x92, 0x56, 0x28, 0xe8, 0x2d, 0x93, 0x99, 0x9d, 0x6f, 0x59, 0x80, 0xc9, 0xf9,
	0xb1, 0x99, 0xbd, 0x23, 0x27, 0x05, 0x85, 0x19, 0xbf, 0xea, 0x77, 0xc8, 0x9f, 0x9c, 0x1f, 0xe5,
	0x05, 0x1c, 0xa2, 0xf9, 0xf8, 0xa6, 0x98, 0x66, 0xe6, 0xd4, 0x6e, 0x4a, 0x56, 0x00, 0xbd, 0xc7,
	0x8e, 0xf0, 0x65, 0x18, 0x51, 0x71, 0xcd, 0x4c, 0x66, 0x77, 0x3f, 0xf2, 0x0c, 0xb2, 0x0f, 0x0c,
	0x2a, 0xd3, 0xcc, 0x1c, 0xdb, 0xf8, 0x94, 0xe7, 0x20, 0xbe, 0xbb, 0xcf, 0x05, 0x55, 0x9e, 0x8a,
	0x56, 0x51, 0xb7, 0x00, 0x91, 0x73, 0xdb, 0xd3, 0xe0, 0x26, 0x79, 0x09, 0xf9, 0x34, 0x2f, 0x94,
	0x58, 0x27, 0xd7, 0x65, 0x93, 0x76, 0x69, 0x62, 0xe0, 0x79, 0xa1, 0x87, 0x23, 0x9b, 0x5c, 0x59,
	0x02, 0xa7, 0xa0, 0x84, 0x66, 0x46, 0x58, 0x4e, 0xe1, 0xae, 0xd8, 0x9a, 0xeb, 0x2b, 0x28, 0xb6,
	0xec, 0x2f, 0x9f, 0xfd, 0xc3, 0xe7, 0x7b, 0x7e, 0xb5, 0xf2, 0xef, 0x91, 0x5a, 0x0c, 0x7f, 0xa7,
	0x5e, 0x0f, 0xe9, 0x2a, 0x37, 0x3f, 0x01, 0x00, 0x00, 0xff, 0xff, 0x5b, 0x77, 0x91, 0xf3, 0x23,
	0x01, 0x00, 0x00,
}
