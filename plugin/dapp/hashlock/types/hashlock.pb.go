// Code generated by protoc-gen-go. DO NOT EDIT.
// source: hashlock.proto

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

type Hashlock struct {
	HashlockId           []byte   `protobuf:"bytes,1,opt,name=hashlockId,proto3" json:"hashlockId,omitempty"`
	Status               int32    `protobuf:"varint,2,opt,name=status,proto3" json:"status,omitempty"`
	CreateTime           int64    `protobuf:"varint,3,opt,name=CreateTime,proto3" json:"CreateTime,omitempty"`
	ToAddress            string   `protobuf:"bytes,4,opt,name=toAddress,proto3" json:"toAddress,omitempty"`
	ReturnAddress        string   `protobuf:"bytes,5,opt,name=returnAddress,proto3" json:"returnAddress,omitempty"`
	Amount               int64    `protobuf:"varint,6,opt,name=amount,proto3" json:"amount,omitempty"`
	Frozentime           int64    `protobuf:"varint,7,opt,name=frozentime,proto3" json:"frozentime,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *Hashlock) Reset()         { *m = Hashlock{} }
func (m *Hashlock) String() string { return proto.CompactTextString(m) }
func (*Hashlock) ProtoMessage()    {}
func (*Hashlock) Descriptor() ([]byte, []int) {
	return fileDescriptor_acb83e90536b5ff8, []int{0}
}

func (m *Hashlock) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Hashlock.Unmarshal(m, b)
}
func (m *Hashlock) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Hashlock.Marshal(b, m, deterministic)
}
func (m *Hashlock) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Hashlock.Merge(m, src)
}
func (m *Hashlock) XXX_Size() int {
	return xxx_messageInfo_Hashlock.Size(m)
}
func (m *Hashlock) XXX_DiscardUnknown() {
	xxx_messageInfo_Hashlock.DiscardUnknown(m)
}

var xxx_messageInfo_Hashlock proto.InternalMessageInfo

func (m *Hashlock) GetHashlockId() []byte {
	if m != nil {
		return m.HashlockId
	}
	return nil
}

func (m *Hashlock) GetStatus() int32 {
	if m != nil {
		return m.Status
	}
	return 0
}

func (m *Hashlock) GetCreateTime() int64 {
	if m != nil {
		return m.CreateTime
	}
	return 0
}

func (m *Hashlock) GetToAddress() string {
	if m != nil {
		return m.ToAddress
	}
	return ""
}

func (m *Hashlock) GetReturnAddress() string {
	if m != nil {
		return m.ReturnAddress
	}
	return ""
}

func (m *Hashlock) GetAmount() int64 {
	if m != nil {
		return m.Amount
	}
	return 0
}

func (m *Hashlock) GetFrozentime() int64 {
	if m != nil {
		return m.Frozentime
	}
	return 0
}

type HashlockLock struct {
	Amount               int64    `protobuf:"varint,1,opt,name=amount,proto3" json:"amount,omitempty"`
	Time                 int64    `protobuf:"varint,2,opt,name=time,proto3" json:"time,omitempty"`
	Hash                 []byte   `protobuf:"bytes,3,opt,name=hash,proto3" json:"hash,omitempty"`
	ToAddress            string   `protobuf:"bytes,4,opt,name=toAddress,proto3" json:"toAddress,omitempty"`
	ReturnAddress        string   `protobuf:"bytes,5,opt,name=returnAddress,proto3" json:"returnAddress,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *HashlockLock) Reset()         { *m = HashlockLock{} }
func (m *HashlockLock) String() string { return proto.CompactTextString(m) }
func (*HashlockLock) ProtoMessage()    {}
func (*HashlockLock) Descriptor() ([]byte, []int) {
	return fileDescriptor_acb83e90536b5ff8, []int{1}
}

func (m *HashlockLock) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_HashlockLock.Unmarshal(m, b)
}
func (m *HashlockLock) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_HashlockLock.Marshal(b, m, deterministic)
}
func (m *HashlockLock) XXX_Merge(src proto.Message) {
	xxx_messageInfo_HashlockLock.Merge(m, src)
}
func (m *HashlockLock) XXX_Size() int {
	return xxx_messageInfo_HashlockLock.Size(m)
}
func (m *HashlockLock) XXX_DiscardUnknown() {
	xxx_messageInfo_HashlockLock.DiscardUnknown(m)
}

var xxx_messageInfo_HashlockLock proto.InternalMessageInfo

func (m *HashlockLock) GetAmount() int64 {
	if m != nil {
		return m.Amount
	}
	return 0
}

func (m *HashlockLock) GetTime() int64 {
	if m != nil {
		return m.Time
	}
	return 0
}

func (m *HashlockLock) GetHash() []byte {
	if m != nil {
		return m.Hash
	}
	return nil
}

func (m *HashlockLock) GetToAddress() string {
	if m != nil {
		return m.ToAddress
	}
	return ""
}

func (m *HashlockLock) GetReturnAddress() string {
	if m != nil {
		return m.ReturnAddress
	}
	return ""
}

type HashlockSend struct {
	Secret               []byte   `protobuf:"bytes,1,opt,name=secret,proto3" json:"secret,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *HashlockSend) Reset()         { *m = HashlockSend{} }
func (m *HashlockSend) String() string { return proto.CompactTextString(m) }
func (*HashlockSend) ProtoMessage()    {}
func (*HashlockSend) Descriptor() ([]byte, []int) {
	return fileDescriptor_acb83e90536b5ff8, []int{2}
}

func (m *HashlockSend) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_HashlockSend.Unmarshal(m, b)
}
func (m *HashlockSend) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_HashlockSend.Marshal(b, m, deterministic)
}
func (m *HashlockSend) XXX_Merge(src proto.Message) {
	xxx_messageInfo_HashlockSend.Merge(m, src)
}
func (m *HashlockSend) XXX_Size() int {
	return xxx_messageInfo_HashlockSend.Size(m)
}
func (m *HashlockSend) XXX_DiscardUnknown() {
	xxx_messageInfo_HashlockSend.DiscardUnknown(m)
}

var xxx_messageInfo_HashlockSend proto.InternalMessageInfo

func (m *HashlockSend) GetSecret() []byte {
	if m != nil {
		return m.Secret
	}
	return nil
}

type Hashlockquery struct {
	Time                 int64    `protobuf:"varint,1,opt,name=time,proto3" json:"time,omitempty"`
	Status               int32    `protobuf:"varint,2,opt,name=status,proto3" json:"status,omitempty"`
	Amount               int64    `protobuf:"varint,3,opt,name=amount,proto3" json:"amount,omitempty"`
	CreateTime           int64    `protobuf:"varint,4,opt,name=createTime,proto3" json:"createTime,omitempty"`
	CurrentTime          int64    `protobuf:"varint,5,opt,name=currentTime,proto3" json:"currentTime,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *Hashlockquery) Reset()         { *m = Hashlockquery{} }
func (m *Hashlockquery) String() string { return proto.CompactTextString(m) }
func (*Hashlockquery) ProtoMessage()    {}
func (*Hashlockquery) Descriptor() ([]byte, []int) {
	return fileDescriptor_acb83e90536b5ff8, []int{3}
}

func (m *Hashlockquery) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Hashlockquery.Unmarshal(m, b)
}
func (m *Hashlockquery) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Hashlockquery.Marshal(b, m, deterministic)
}
func (m *Hashlockquery) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Hashlockquery.Merge(m, src)
}
func (m *Hashlockquery) XXX_Size() int {
	return xxx_messageInfo_Hashlockquery.Size(m)
}
func (m *Hashlockquery) XXX_DiscardUnknown() {
	xxx_messageInfo_Hashlockquery.DiscardUnknown(m)
}

var xxx_messageInfo_Hashlockquery proto.InternalMessageInfo

func (m *Hashlockquery) GetTime() int64 {
	if m != nil {
		return m.Time
	}
	return 0
}

func (m *Hashlockquery) GetStatus() int32 {
	if m != nil {
		return m.Status
	}
	return 0
}

func (m *Hashlockquery) GetAmount() int64 {
	if m != nil {
		return m.Amount
	}
	return 0
}

func (m *Hashlockquery) GetCreateTime() int64 {
	if m != nil {
		return m.CreateTime
	}
	return 0
}

func (m *Hashlockquery) GetCurrentTime() int64 {
	if m != nil {
		return m.CurrentTime
	}
	return 0
}

type HashRecv struct {
	HashlockId           []byte         `protobuf:"bytes,1,opt,name=HashlockId,proto3" json:"HashlockId,omitempty"`
	Information          *Hashlockquery `protobuf:"bytes,2,opt,name=Information,proto3" json:"Information,omitempty"`
	XXX_NoUnkeyedLiteral struct{}       `json:"-"`
	XXX_unrecognized     []byte         `json:"-"`
	XXX_sizecache        int32          `json:"-"`
}

func (m *HashRecv) Reset()         { *m = HashRecv{} }
func (m *HashRecv) String() string { return proto.CompactTextString(m) }
func (*HashRecv) ProtoMessage()    {}
func (*HashRecv) Descriptor() ([]byte, []int) {
	return fileDescriptor_acb83e90536b5ff8, []int{4}
}

func (m *HashRecv) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_HashRecv.Unmarshal(m, b)
}
func (m *HashRecv) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_HashRecv.Marshal(b, m, deterministic)
}
func (m *HashRecv) XXX_Merge(src proto.Message) {
	xxx_messageInfo_HashRecv.Merge(m, src)
}
func (m *HashRecv) XXX_Size() int {
	return xxx_messageInfo_HashRecv.Size(m)
}
func (m *HashRecv) XXX_DiscardUnknown() {
	xxx_messageInfo_HashRecv.DiscardUnknown(m)
}

var xxx_messageInfo_HashRecv proto.InternalMessageInfo

func (m *HashRecv) GetHashlockId() []byte {
	if m != nil {
		return m.HashlockId
	}
	return nil
}

func (m *HashRecv) GetInformation() *Hashlockquery {
	if m != nil {
		return m.Information
	}
	return nil
}

type HashlockUnlock struct {
	Secret               []byte   `protobuf:"bytes,1,opt,name=secret,proto3" json:"secret,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *HashlockUnlock) Reset()         { *m = HashlockUnlock{} }
func (m *HashlockUnlock) String() string { return proto.CompactTextString(m) }
func (*HashlockUnlock) ProtoMessage()    {}
func (*HashlockUnlock) Descriptor() ([]byte, []int) {
	return fileDescriptor_acb83e90536b5ff8, []int{5}
}

func (m *HashlockUnlock) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_HashlockUnlock.Unmarshal(m, b)
}
func (m *HashlockUnlock) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_HashlockUnlock.Marshal(b, m, deterministic)
}
func (m *HashlockUnlock) XXX_Merge(src proto.Message) {
	xxx_messageInfo_HashlockUnlock.Merge(m, src)
}
func (m *HashlockUnlock) XXX_Size() int {
	return xxx_messageInfo_HashlockUnlock.Size(m)
}
func (m *HashlockUnlock) XXX_DiscardUnknown() {
	xxx_messageInfo_HashlockUnlock.DiscardUnknown(m)
}

var xxx_messageInfo_HashlockUnlock proto.InternalMessageInfo

func (m *HashlockUnlock) GetSecret() []byte {
	if m != nil {
		return m.Secret
	}
	return nil
}

// message for hashlock
type HashlockAction struct {
	// Types that are valid to be assigned to Value:
	//	*HashlockAction_Hlock
	//	*HashlockAction_Hsend
	//	*HashlockAction_Hunlock
	Value                isHashlockAction_Value `protobuf_oneof:"value"`
	Ty                   int32                  `protobuf:"varint,4,opt,name=ty,proto3" json:"ty,omitempty"`
	XXX_NoUnkeyedLiteral struct{}               `json:"-"`
	XXX_unrecognized     []byte                 `json:"-"`
	XXX_sizecache        int32                  `json:"-"`
}

func (m *HashlockAction) Reset()         { *m = HashlockAction{} }
func (m *HashlockAction) String() string { return proto.CompactTextString(m) }
func (*HashlockAction) ProtoMessage()    {}
func (*HashlockAction) Descriptor() ([]byte, []int) {
	return fileDescriptor_acb83e90536b5ff8, []int{6}
}

func (m *HashlockAction) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_HashlockAction.Unmarshal(m, b)
}
func (m *HashlockAction) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_HashlockAction.Marshal(b, m, deterministic)
}
func (m *HashlockAction) XXX_Merge(src proto.Message) {
	xxx_messageInfo_HashlockAction.Merge(m, src)
}
func (m *HashlockAction) XXX_Size() int {
	return xxx_messageInfo_HashlockAction.Size(m)
}
func (m *HashlockAction) XXX_DiscardUnknown() {
	xxx_messageInfo_HashlockAction.DiscardUnknown(m)
}

var xxx_messageInfo_HashlockAction proto.InternalMessageInfo

type isHashlockAction_Value interface {
	isHashlockAction_Value()
}

type HashlockAction_Hlock struct {
	Hlock *HashlockLock `protobuf:"bytes,1,opt,name=hlock,proto3,oneof"`
}

type HashlockAction_Hsend struct {
	Hsend *HashlockSend `protobuf:"bytes,2,opt,name=hsend,proto3,oneof"`
}

type HashlockAction_Hunlock struct {
	Hunlock *HashlockUnlock `protobuf:"bytes,3,opt,name=hunlock,proto3,oneof"`
}

func (*HashlockAction_Hlock) isHashlockAction_Value() {}

func (*HashlockAction_Hsend) isHashlockAction_Value() {}

func (*HashlockAction_Hunlock) isHashlockAction_Value() {}

func (m *HashlockAction) GetValue() isHashlockAction_Value {
	if m != nil {
		return m.Value
	}
	return nil
}

func (m *HashlockAction) GetHlock() *HashlockLock {
	if x, ok := m.GetValue().(*HashlockAction_Hlock); ok {
		return x.Hlock
	}
	return nil
}

func (m *HashlockAction) GetHsend() *HashlockSend {
	if x, ok := m.GetValue().(*HashlockAction_Hsend); ok {
		return x.Hsend
	}
	return nil
}

func (m *HashlockAction) GetHunlock() *HashlockUnlock {
	if x, ok := m.GetValue().(*HashlockAction_Hunlock); ok {
		return x.Hunlock
	}
	return nil
}

func (m *HashlockAction) GetTy() int32 {
	if m != nil {
		return m.Ty
	}
	return 0
}

// XXX_OneofFuncs is for the internal use of the proto package.
func (*HashlockAction) XXX_OneofFuncs() (func(msg proto.Message, b *proto.Buffer) error, func(msg proto.Message, tag, wire int, b *proto.Buffer) (bool, error), func(msg proto.Message) (n int), []interface{}) {
	return _HashlockAction_OneofMarshaler, _HashlockAction_OneofUnmarshaler, _HashlockAction_OneofSizer, []interface{}{
		(*HashlockAction_Hlock)(nil),
		(*HashlockAction_Hsend)(nil),
		(*HashlockAction_Hunlock)(nil),
	}
}

func _HashlockAction_OneofMarshaler(msg proto.Message, b *proto.Buffer) error {
	m := msg.(*HashlockAction)
	// value
	switch x := m.Value.(type) {
	case *HashlockAction_Hlock:
		b.EncodeVarint(1<<3 | proto.WireBytes)
		if err := b.EncodeMessage(x.Hlock); err != nil {
			return err
		}
	case *HashlockAction_Hsend:
		b.EncodeVarint(2<<3 | proto.WireBytes)
		if err := b.EncodeMessage(x.Hsend); err != nil {
			return err
		}
	case *HashlockAction_Hunlock:
		b.EncodeVarint(3<<3 | proto.WireBytes)
		if err := b.EncodeMessage(x.Hunlock); err != nil {
			return err
		}
	case nil:
	default:
		return fmt.Errorf("HashlockAction.Value has unexpected type %T", x)
	}
	return nil
}

func _HashlockAction_OneofUnmarshaler(msg proto.Message, tag, wire int, b *proto.Buffer) (bool, error) {
	m := msg.(*HashlockAction)
	switch tag {
	case 1: // value.hlock
		if wire != proto.WireBytes {
			return true, proto.ErrInternalBadWireType
		}
		msg := new(HashlockLock)
		err := b.DecodeMessage(msg)
		m.Value = &HashlockAction_Hlock{msg}
		return true, err
	case 2: // value.hsend
		if wire != proto.WireBytes {
			return true, proto.ErrInternalBadWireType
		}
		msg := new(HashlockSend)
		err := b.DecodeMessage(msg)
		m.Value = &HashlockAction_Hsend{msg}
		return true, err
	case 3: // value.hunlock
		if wire != proto.WireBytes {
			return true, proto.ErrInternalBadWireType
		}
		msg := new(HashlockUnlock)
		err := b.DecodeMessage(msg)
		m.Value = &HashlockAction_Hunlock{msg}
		return true, err
	default:
		return false, nil
	}
}

func _HashlockAction_OneofSizer(msg proto.Message) (n int) {
	m := msg.(*HashlockAction)
	// value
	switch x := m.Value.(type) {
	case *HashlockAction_Hlock:
		s := proto.Size(x.Hlock)
		n += 1 // tag and wire
		n += proto.SizeVarint(uint64(s))
		n += s
	case *HashlockAction_Hsend:
		s := proto.Size(x.Hsend)
		n += 1 // tag and wire
		n += proto.SizeVarint(uint64(s))
		n += s
	case *HashlockAction_Hunlock:
		s := proto.Size(x.Hunlock)
		n += 1 // tag and wire
		n += proto.SizeVarint(uint64(s))
		n += s
	case nil:
	default:
		panic(fmt.Sprintf("proto: unexpected type %T in oneof", x))
	}
	return n
}

func init() {
	proto.RegisterType((*Hashlock)(nil), "types.Hashlock")
	proto.RegisterType((*HashlockLock)(nil), "types.HashlockLock")
	proto.RegisterType((*HashlockSend)(nil), "types.HashlockSend")
	proto.RegisterType((*Hashlockquery)(nil), "types.Hashlockquery")
	proto.RegisterType((*HashRecv)(nil), "types.HashRecv")
	proto.RegisterType((*HashlockUnlock)(nil), "types.HashlockUnlock")
	proto.RegisterType((*HashlockAction)(nil), "types.HashlockAction")
}

func init() { proto.RegisterFile("hashlock.proto", fileDescriptor_acb83e90536b5ff8) }

var fileDescriptor_acb83e90536b5ff8 = []byte{
	// 421 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xac, 0x93, 0xcd, 0x6a, 0xdb, 0x40,
	0x10, 0xc7, 0xbd, 0xb2, 0x65, 0xd7, 0x23, 0xdb, 0x87, 0xed, 0x07, 0x3a, 0x94, 0x22, 0x44, 0x29,
	0x82, 0x82, 0xa1, 0x2e, 0xf4, 0xee, 0xf6, 0x22, 0x43, 0x4f, 0x9b, 0xe4, 0x01, 0x64, 0x69, 0x8d,
	0x4d, 0xec, 0x5d, 0x67, 0xb5, 0x32, 0x28, 0x8f, 0x11, 0xc8, 0xab, 0xe4, 0x71, 0xf2, 0x2c, 0x61,
	0x47, 0x52, 0x56, 0x56, 0x92, 0x5b, 0x6e, 0x9a, 0xbf, 0x7e, 0x9a, 0x8f, 0xff, 0x8c, 0x60, 0xb6,
	0x4d, 0xf2, 0xed, 0x5e, 0xa6, 0xd7, 0xf3, 0xa3, 0x92, 0x5a, 0x52, 0x57, 0x97, 0x47, 0x9e, 0x87,
	0x8f, 0x04, 0x3e, 0xc4, 0xf5, 0x1b, 0xfa, 0x0d, 0xa0, 0xa1, 0x56, 0x99, 0x4f, 0x02, 0x12, 0x4d,
	0x58, 0x4b, 0xa1, 0x5f, 0x60, 0x98, 0xeb, 0x44, 0x17, 0xb9, 0xef, 0x04, 0x24, 0x72, 0x59, 0x1d,
	0x99, 0xef, 0xfe, 0x29, 0x9e, 0x68, 0x7e, 0xb9, 0x3b, 0x70, 0xbf, 0x1f, 0x90, 0xa8, 0xcf, 0x5a,
	0x0a, 0xfd, 0x0a, 0x63, 0x2d, 0x97, 0x59, 0xa6, 0x78, 0x9e, 0xfb, 0x83, 0x80, 0x44, 0x63, 0x66,
	0x05, 0xfa, 0x1d, 0xa6, 0x8a, 0xeb, 0x42, 0x89, 0x86, 0x70, 0x91, 0x38, 0x17, 0x4d, 0xed, 0xe4,
	0x20, 0x0b, 0xa1, 0xfd, 0x21, 0xe6, 0xaf, 0x23, 0x53, 0x7b, 0xa3, 0xe4, 0x2d, 0x17, 0xda, 0xd4,
	0x1e, 0x55, 0xb5, 0xad, 0x12, 0xde, 0x11, 0x98, 0x34, 0x03, 0xfe, 0x37, 0x43, 0xda, 0x44, 0xe4,
	0x2c, 0x11, 0x85, 0x01, 0xa6, 0x70, 0x50, 0xc5, 0x67, 0xa3, 0x99, 0xf1, 0x71, 0xa4, 0x09, 0xc3,
	0xe7, 0xf7, 0x18, 0x26, 0xfc, 0x61, 0x7b, 0xba, 0xe0, 0xa2, 0x32, 0x96, 0xa7, 0x8a, 0xeb, 0xda,
	0xf4, 0x3a, 0x0a, 0xef, 0x09, 0x4c, 0x1b, 0xf0, 0xa6, 0xe0, 0xaa, 0x7c, 0xee, 0x92, 0xb4, 0xba,
	0x7c, 0x6b, 0x2d, 0x76, 0xd2, 0x7e, 0xd7, 0xb2, 0xd4, 0xae, 0x6b, 0x50, 0x59, 0x66, 0x15, 0x1a,
	0x80, 0x97, 0x16, 0x4a, 0x71, 0xa1, 0x11, 0x70, 0x11, 0x68, 0x4b, 0xe1, 0xba, 0x3a, 0x1a, 0xc6,
	0xd3, 0x93, 0xc9, 0x16, 0xbf, 0x38, 0x1a, 0xab, 0xd0, 0x3f, 0xe0, 0xad, 0xc4, 0x46, 0xaa, 0x43,
	0xa2, 0x77, 0x52, 0x60, 0x8b, 0xde, 0xe2, 0xd3, 0x1c, 0xcf, 0x6f, 0x7e, 0x36, 0x1c, 0x6b, 0x83,
	0x61, 0x04, 0xb3, 0xe6, 0xed, 0x95, 0xd8, 0xd7, 0x9b, 0x7b, 0xd5, 0xa5, 0x07, 0x62, 0xd1, 0x65,
	0x6a, 0x3e, 0xa6, 0x3f, 0xc1, 0xc5, 0x10, 0x49, 0x6f, 0xf1, 0xb1, 0x53, 0xce, 0x1c, 0x42, 0xdc,
	0x63, 0x15, 0x83, 0x70, 0xce, 0x45, 0x56, 0xf7, 0xd6, 0x85, 0xcd, 0x86, 0x10, 0x36, 0x0c, 0xfd,
	0x05, 0xa3, 0x6d, 0x81, 0xfd, 0xa0, 0xab, 0xde, 0xe2, 0x73, 0x07, 0xaf, 0x9a, 0x8d, 0x7b, 0xac,
	0xe1, 0xe8, 0x0c, 0x1c, 0x5d, 0xa2, 0xcf, 0x2e, 0x73, 0x74, 0xf9, 0x77, 0x04, 0xee, 0x29, 0xd9,
	0x17, 0x7c, 0x3d, 0xc4, 0x5f, 0xf1, 0xf7, 0x53, 0x00, 0x00, 0x00, 0xff, 0xff, 0x38, 0xf6, 0x91,
	0xc1, 0x9c, 0x03, 0x00, 0x00,
}
