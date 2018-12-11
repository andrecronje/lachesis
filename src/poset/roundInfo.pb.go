// Code generated by protoc-gen-go. DO NOT EDIT.
// source: roundInfo.proto

package poset

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

type Trilean int32

const (
	Trilean_UNDEFINED Trilean = 0
	Trilean_TRUE      Trilean = 1
	Trilean_FALSE     Trilean = 2
)

var Trilean_name = map[int32]string{
	0: "UNDEFINED",
	1: "TRUE",
	2: "FALSE",
}

var Trilean_value = map[string]int32{
	"UNDEFINED": 0,
	"TRUE":      1,
	"FALSE":     2,
}

func (x Trilean) String() string {
	return proto.EnumName(Trilean_name, int32(x))
}

func (Trilean) EnumDescriptor() ([]byte, []int) {
	return fileDescriptor_06add9379b5a2b9a, []int{0}
}

type RoundEvent struct {
	Consensus            bool     `protobuf:"varint,1,opt,name=Consensus,proto3" json:"Consensus,omitempty"`
	Witness              bool     `protobuf:"varint,2,opt,name=Witness,proto3" json:"Witness,omitempty"`
	Famous               Trilean  `protobuf:"varint,3,opt,name=Famous,proto3,enum=poset.Trilean" json:"Famous,omitempty"`
	RoundReceived        int64    `protobuf:"varint,4,opt,name=RoundReceived,proto3" json:"RoundReceived,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *RoundEvent) Reset()         { *m = RoundEvent{} }
func (m *RoundEvent) String() string { return proto.CompactTextString(m) }
func (*RoundEvent) ProtoMessage()    {}
func (*RoundEvent) Descriptor() ([]byte, []int) {
	return fileDescriptor_06add9379b5a2b9a, []int{0}
}

func (m *RoundEvent) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_RoundEvent.Unmarshal(m, b)
}
func (m *RoundEvent) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_RoundEvent.Marshal(b, m, deterministic)
}
func (m *RoundEvent) XXX_Merge(src proto.Message) {
	xxx_messageInfo_RoundEvent.Merge(m, src)
}
func (m *RoundEvent) XXX_Size() int {
	return xxx_messageInfo_RoundEvent.Size(m)
}
func (m *RoundEvent) XXX_DiscardUnknown() {
	xxx_messageInfo_RoundEvent.DiscardUnknown(m)
}

var xxx_messageInfo_RoundEvent proto.InternalMessageInfo

func (m *RoundEvent) GetConsensus() bool {
	if m != nil {
		return m.Consensus
	}
	return false
}

func (m *RoundEvent) GetWitness() bool {
	if m != nil {
		return m.Witness
	}
	return false
}

func (m *RoundEvent) GetFamous() Trilean {
	if m != nil {
		return m.Famous
	}
	return Trilean_UNDEFINED
}

func (m *RoundEvent) GetRoundReceived() int64 {
	if m != nil {
		return m.RoundReceived
	}
	return 0
}

type RoundCreatedMessage struct {
	Events               map[string]*RoundEvent `protobuf:"bytes,1,rep,name=Events,proto3" json:"Events,omitempty" protobuf_key:"bytes,1,opt,name=key,proto3" protobuf_val:"bytes,2,opt,name=value,proto3"`
	Queued               bool                   `protobuf:"varint,2,opt,name=queued,proto3" json:"queued,omitempty"`
	XXX_NoUnkeyedLiteral struct{}               `json:"-"`
	XXX_unrecognized     []byte                 `json:"-"`
	XXX_sizecache        int32                  `json:"-"`
}

func (m *RoundCreatedMessage) Reset()         { *m = RoundCreatedMessage{} }
func (m *RoundCreatedMessage) String() string { return proto.CompactTextString(m) }
func (*RoundCreatedMessage) ProtoMessage()    {}
func (*RoundCreatedMessage) Descriptor() ([]byte, []int) {
	return fileDescriptor_06add9379b5a2b9a, []int{1}
}

func (m *RoundCreatedMessage) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_RoundCreatedMessage.Unmarshal(m, b)
}
func (m *RoundCreatedMessage) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_RoundCreatedMessage.Marshal(b, m, deterministic)
}
func (m *RoundCreatedMessage) XXX_Merge(src proto.Message) {
	xxx_messageInfo_RoundCreatedMessage.Merge(m, src)
}
func (m *RoundCreatedMessage) XXX_Size() int {
	return xxx_messageInfo_RoundCreatedMessage.Size(m)
}
func (m *RoundCreatedMessage) XXX_DiscardUnknown() {
	xxx_messageInfo_RoundCreatedMessage.DiscardUnknown(m)
}

var xxx_messageInfo_RoundCreatedMessage proto.InternalMessageInfo

func (m *RoundCreatedMessage) GetEvents() map[string]*RoundEvent {
	if m != nil {
		return m.Events
	}
	return nil
}

func (m *RoundCreatedMessage) GetQueued() bool {
	if m != nil {
		return m.Queued
	}
	return false
}

type RoundReceived struct {
	Rounds               []string `protobuf:"bytes,1,rep,name=Rounds,proto3" json:"Rounds,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *RoundReceived) Reset()         { *m = RoundReceived{} }
func (m *RoundReceived) String() string { return proto.CompactTextString(m) }
func (*RoundReceived) ProtoMessage()    {}
func (*RoundReceived) Descriptor() ([]byte, []int) {
	return fileDescriptor_06add9379b5a2b9a, []int{2}
}

func (m *RoundReceived) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_RoundReceived.Unmarshal(m, b)
}
func (m *RoundReceived) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_RoundReceived.Marshal(b, m, deterministic)
}
func (m *RoundReceived) XXX_Merge(src proto.Message) {
	xxx_messageInfo_RoundReceived.Merge(m, src)
}
func (m *RoundReceived) XXX_Size() int {
	return xxx_messageInfo_RoundReceived.Size(m)
}
func (m *RoundReceived) XXX_DiscardUnknown() {
	xxx_messageInfo_RoundReceived.DiscardUnknown(m)
}

var xxx_messageInfo_RoundReceived proto.InternalMessageInfo

func (m *RoundReceived) GetRounds() []string {
	if m != nil {
		return m.Rounds
	}
	return nil
}

func init() {
	proto.RegisterEnum("poset.Trilean", Trilean_name, Trilean_value)
	proto.RegisterType((*RoundEvent)(nil), "poset.RoundEvent")
	proto.RegisterType((*RoundCreatedMessage)(nil), "poset.RoundCreatedMessage")
	proto.RegisterMapType((map[string]*RoundEvent)(nil), "poset.RoundCreatedMessage.EventsEntry")
	proto.RegisterType((*RoundReceived)(nil), "poset.RoundReceived")
}

func init() { proto.RegisterFile("roundInfo.proto", fileDescriptor_06add9379b5a2b9a) }

var fileDescriptor_06add9379b5a2b9a = []byte{
	// 313 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x6c, 0x51, 0xdf, 0x4a, 0x3a, 0x41,
	0x14, 0xfe, 0x8d, 0xab, 0xab, 0x7b, 0x44, 0x7f, 0xdb, 0x09, 0x62, 0x89, 0x2e, 0x44, 0x42, 0x97,
	0xa0, 0xbd, 0xb0, 0x9b, 0xe8, 0x22, 0x08, 0x5d, 0x41, 0x30, 0x2f, 0x26, 0xa5, 0xeb, 0xad, 0x3d,
	0x85, 0x64, 0x33, 0xb6, 0x33, 0x23, 0xf8, 0x1a, 0xbd, 0x4e, 0x2f, 0x17, 0x3b, 0x4e, 0x64, 0xd1,
	0xdd, 0x9c, 0xef, 0xcf, 0x7c, 0xdf, 0x99, 0x81, 0xff, 0x85, 0x34, 0x22, 0x9f, 0x88, 0x27, 0x99,
	0xac, 0x0b, 0xa9, 0x25, 0xd6, 0xd6, 0x52, 0x91, 0xee, 0xbe, 0x33, 0x00, 0x5e, 0x52, 0xe9, 0x86,
	0x84, 0xc6, 0x13, 0x08, 0x86, 0x52, 0x28, 0x12, 0xca, 0xa8, 0x88, 0x75, 0x58, 0xdc, 0xe0, 0xdf,
	0x00, 0x46, 0x50, 0xbf, 0x5f, 0x6a, 0x41, 0x4a, 0x45, 0x15, 0xcb, 0x7d, 0x8d, 0xd8, 0x03, 0x7f,
	0x9c, 0xbd, 0x4a, 0xa3, 0x22, 0xaf, 0xc3, 0xe2, 0xf6, 0xa0, 0x9d, 0xd8, 0xeb, 0x93, 0x79, 0xb1,
	0x5c, 0x51, 0x26, 0xb8, 0x63, 0xf1, 0x14, 0x5a, 0x36, 0x8d, 0xd3, 0x23, 0x2d, 0x37, 0x94, 0x47,
	0xd5, 0x0e, 0x8b, 0x3d, 0xfe, 0x13, 0xec, 0x7e, 0x30, 0x38, 0xb4, 0xc8, 0xb0, 0xa0, 0x4c, 0x53,
	0x7e, 0x4b, 0x4a, 0x65, 0xcf, 0x84, 0xd7, 0xe0, 0xdb, 0x9a, 0x65, 0x35, 0x2f, 0x6e, 0x0e, 0x7a,
	0x2e, 0xe5, 0x0f, 0x6d, 0xb2, 0x13, 0xa6, 0x42, 0x17, 0x5b, 0xee, 0x5c, 0x78, 0x04, 0xfe, 0x9b,
	0x21, 0x43, 0xb9, 0xab, 0xef, 0xa6, 0xe3, 0x29, 0x34, 0xf7, 0xe4, 0x18, 0x82, 0xf7, 0x42, 0x5b,
	0xbb, 0x7e, 0xc0, 0xcb, 0x23, 0xf6, 0xa1, 0xb6, 0xc9, 0x56, 0x86, 0xac, 0xaf, 0x39, 0x38, 0xd8,
	0xcf, 0xb5, 0x4e, 0xbe, 0xe3, 0xaf, 0x2a, 0x97, 0xac, 0xdb, 0xff, 0xb5, 0x63, 0x19, 0x6b, 0x81,
	0x5d, 0xed, 0x80, 0xbb, 0xe9, 0xec, 0x1c, 0xea, 0xee, 0x7d, 0xb0, 0x05, 0xc1, 0x62, 0x36, 0x4a,
	0xc7, 0x93, 0x59, 0x3a, 0x0a, 0xff, 0x61, 0x03, 0xaa, 0x73, 0xbe, 0x48, 0x43, 0x86, 0x01, 0xd4,
	0xc6, 0x37, 0xd3, 0xbb, 0x34, 0xac, 0x3c, 0xf8, 0xf6, 0xe3, 0x2e, 0x3e, 0x03, 0x00, 0x00, 0xff,
	0xff, 0x3a, 0x3d, 0x34, 0xb9, 0xcb, 0x01, 0x00, 0x00,
}
