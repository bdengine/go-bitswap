// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: message.proto

package bitswap_message_pb

import (
	fmt "fmt"
	_ "github.com/gogo/protobuf/gogoproto"
	proto "github.com/gogo/protobuf/proto"
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
const _ = proto.GoGoProtoPackageIsVersion3 // please upgrade the proto package

type Message_BlockPresenceType int32

const (
	Message_Have     Message_BlockPresenceType = 0
	Message_DontHave Message_BlockPresenceType = 1
)

var Message_BlockPresenceType_name = map[int32]string{
	0: "Have",
	1: "DontHave",
}

var Message_BlockPresenceType_value = map[string]int32{
	"Have":     0,
	"DontHave": 1,
}

func (x Message_BlockPresenceType) String() string {
	return proto.EnumName(Message_BlockPresenceType_name, int32(x))
}

func (Message_BlockPresenceType) EnumDescriptor() ([]byte, []int) {
	return fileDescriptor_33c57e4bae7b9afd, []int{0, 0}
}

type Message_Wantlist_WantType int32

const (
	Message_Wantlist_Block Message_Wantlist_WantType = 0
	Message_Wantlist_Have  Message_Wantlist_WantType = 1
)

var Message_Wantlist_WantType_name = map[int32]string{
	0: "Block",
	1: "Have",
}

var Message_Wantlist_WantType_value = map[string]int32{
	"Block": 0,
	"Have":  1,
}

func (x Message_Wantlist_WantType) String() string {
	return proto.EnumName(Message_Wantlist_WantType_name, int32(x))
}

func (Message_Wantlist_WantType) EnumDescriptor() ([]byte, []int) {
	return fileDescriptor_33c57e4bae7b9afd, []int{0, 0, 0}
}

type Message struct {
	Wantlist             Message_Wantlist        `protobuf:"bytes,1,opt,name=wantlist,proto3" json:"wantlist"`
	Blocks               [][]byte                `protobuf:"bytes,2,rep,name=blocks,proto3" json:"blocks,omitempty"`
	Payload              []Message_Block         `protobuf:"bytes,3,rep,name=payload,proto3" json:"payload"`
	BlockPresences       []Message_BlockPresence `protobuf:"bytes,4,rep,name=blockPresences,proto3" json:"blockPresences"`
	PendingBytes         int32                   `protobuf:"varint,5,opt,name=pendingBytes,proto3" json:"pendingBytes,omitempty"`
	BackupLoad           []*Message_BackupLoad   `protobuf:"bytes,6,rep,name=backupLoad,proto3" json:"backupLoad,omitempty"`
	DeleteLoad           []*Message_DeleteLoad   `protobuf:"bytes,7,rep,name=deleteLoad,proto3" json:"deleteLoad,omitempty"`
	XXX_NoUnkeyedLiteral struct{}                `json:"-"`
	XXX_unrecognized     []byte                  `json:"-"`
	XXX_sizecache        int32                   `json:"-"`
}

func (m *Message) Reset()         { *m = Message{} }
func (m *Message) String() string { return proto.CompactTextString(m) }
func (*Message) ProtoMessage()    {}
func (*Message) Descriptor() ([]byte, []int) {
	return fileDescriptor_33c57e4bae7b9afd, []int{0}
}
func (m *Message) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Message.Unmarshal(m, b)
}
func (m *Message) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Message.Marshal(b, m, deterministic)
}
func (m *Message) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Message.Merge(m, src)
}
func (m *Message) XXX_Size() int {
	return xxx_messageInfo_Message.Size(m)
}
func (m *Message) XXX_DiscardUnknown() {
	xxx_messageInfo_Message.DiscardUnknown(m)
}

var xxx_messageInfo_Message proto.InternalMessageInfo

func (m *Message) GetWantlist() Message_Wantlist {
	if m != nil {
		return m.Wantlist
	}
	return Message_Wantlist{}
}

func (m *Message) GetBlocks() [][]byte {
	if m != nil {
		return m.Blocks
	}
	return nil
}

func (m *Message) GetPayload() []Message_Block {
	if m != nil {
		return m.Payload
	}
	return nil
}

func (m *Message) GetBlockPresences() []Message_BlockPresence {
	if m != nil {
		return m.BlockPresences
	}
	return nil
}

func (m *Message) GetPendingBytes() int32 {
	if m != nil {
		return m.PendingBytes
	}
	return 0
}

func (m *Message) GetBackupLoad() []*Message_BackupLoad {
	if m != nil {
		return m.BackupLoad
	}
	return nil
}

func (m *Message) GetDeleteLoad() []*Message_DeleteLoad {
	if m != nil {
		return m.DeleteLoad
	}
	return nil
}

type Message_Wantlist struct {
	Entries              []Message_Wantlist_Entry `protobuf:"bytes,1,rep,name=entries,proto3" json:"entries"`
	Full                 bool                     `protobuf:"varint,2,opt,name=full,proto3" json:"full,omitempty"`
	XXX_NoUnkeyedLiteral struct{}                 `json:"-"`
	XXX_unrecognized     []byte                   `json:"-"`
	XXX_sizecache        int32                    `json:"-"`
}

func (m *Message_Wantlist) Reset()         { *m = Message_Wantlist{} }
func (m *Message_Wantlist) String() string { return proto.CompactTextString(m) }
func (*Message_Wantlist) ProtoMessage()    {}
func (*Message_Wantlist) Descriptor() ([]byte, []int) {
	return fileDescriptor_33c57e4bae7b9afd, []int{0, 0}
}
func (m *Message_Wantlist) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Message_Wantlist.Unmarshal(m, b)
}
func (m *Message_Wantlist) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Message_Wantlist.Marshal(b, m, deterministic)
}
func (m *Message_Wantlist) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Message_Wantlist.Merge(m, src)
}
func (m *Message_Wantlist) XXX_Size() int {
	return xxx_messageInfo_Message_Wantlist.Size(m)
}
func (m *Message_Wantlist) XXX_DiscardUnknown() {
	xxx_messageInfo_Message_Wantlist.DiscardUnknown(m)
}

var xxx_messageInfo_Message_Wantlist proto.InternalMessageInfo

func (m *Message_Wantlist) GetEntries() []Message_Wantlist_Entry {
	if m != nil {
		return m.Entries
	}
	return nil
}

func (m *Message_Wantlist) GetFull() bool {
	if m != nil {
		return m.Full
	}
	return false
}

type Message_Wantlist_Entry struct {
	Block                Cid                       `protobuf:"bytes,1,opt,name=block,proto3,customtype=Cid" json:"block"`
	Priority             int32                     `protobuf:"varint,2,opt,name=priority,proto3" json:"priority,omitempty"`
	Cancel               bool                      `protobuf:"varint,3,opt,name=cancel,proto3" json:"cancel,omitempty"`
	WantType             Message_Wantlist_WantType `protobuf:"varint,4,opt,name=wantType,proto3,enum=bitswap.message.pb.Message_Wantlist_WantType" json:"wantType,omitempty"`
	SendDontHave         bool                      `protobuf:"varint,5,opt,name=sendDontHave,proto3" json:"sendDontHave,omitempty"`
	XXX_NoUnkeyedLiteral struct{}                  `json:"-"`
	XXX_unrecognized     []byte                    `json:"-"`
	XXX_sizecache        int32                     `json:"-"`
}

func (m *Message_Wantlist_Entry) Reset()         { *m = Message_Wantlist_Entry{} }
func (m *Message_Wantlist_Entry) String() string { return proto.CompactTextString(m) }
func (*Message_Wantlist_Entry) ProtoMessage()    {}
func (*Message_Wantlist_Entry) Descriptor() ([]byte, []int) {
	return fileDescriptor_33c57e4bae7b9afd, []int{0, 0, 0}
}
func (m *Message_Wantlist_Entry) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Message_Wantlist_Entry.Unmarshal(m, b)
}
func (m *Message_Wantlist_Entry) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Message_Wantlist_Entry.Marshal(b, m, deterministic)
}
func (m *Message_Wantlist_Entry) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Message_Wantlist_Entry.Merge(m, src)
}
func (m *Message_Wantlist_Entry) XXX_Size() int {
	return xxx_messageInfo_Message_Wantlist_Entry.Size(m)
}
func (m *Message_Wantlist_Entry) XXX_DiscardUnknown() {
	xxx_messageInfo_Message_Wantlist_Entry.DiscardUnknown(m)
}

var xxx_messageInfo_Message_Wantlist_Entry proto.InternalMessageInfo

func (m *Message_Wantlist_Entry) GetPriority() int32 {
	if m != nil {
		return m.Priority
	}
	return 0
}

func (m *Message_Wantlist_Entry) GetCancel() bool {
	if m != nil {
		return m.Cancel
	}
	return false
}

func (m *Message_Wantlist_Entry) GetWantType() Message_Wantlist_WantType {
	if m != nil {
		return m.WantType
	}
	return Message_Wantlist_Block
}

func (m *Message_Wantlist_Entry) GetSendDontHave() bool {
	if m != nil {
		return m.SendDontHave
	}
	return false
}

type Message_Block struct {
	Prefix               []byte   `protobuf:"bytes,1,opt,name=prefix,proto3" json:"prefix,omitempty"`
	Data                 []byte   `protobuf:"bytes,2,opt,name=data,proto3" json:"data,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *Message_Block) Reset()         { *m = Message_Block{} }
func (m *Message_Block) String() string { return proto.CompactTextString(m) }
func (*Message_Block) ProtoMessage()    {}
func (*Message_Block) Descriptor() ([]byte, []int) {
	return fileDescriptor_33c57e4bae7b9afd, []int{0, 1}
}
func (m *Message_Block) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Message_Block.Unmarshal(m, b)
}
func (m *Message_Block) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Message_Block.Marshal(b, m, deterministic)
}
func (m *Message_Block) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Message_Block.Merge(m, src)
}
func (m *Message_Block) XXX_Size() int {
	return xxx_messageInfo_Message_Block.Size(m)
}
func (m *Message_Block) XXX_DiscardUnknown() {
	xxx_messageInfo_Message_Block.DiscardUnknown(m)
}

var xxx_messageInfo_Message_Block proto.InternalMessageInfo

func (m *Message_Block) GetPrefix() []byte {
	if m != nil {
		return m.Prefix
	}
	return nil
}

func (m *Message_Block) GetData() []byte {
	if m != nil {
		return m.Data
	}
	return nil
}

type Message_BlockPresence struct {
	Cid                  Cid                       `protobuf:"bytes,1,opt,name=cid,proto3,customtype=Cid" json:"cid"`
	Type                 Message_BlockPresenceType `protobuf:"varint,2,opt,name=type,proto3,enum=bitswap.message.pb.Message_BlockPresenceType" json:"type,omitempty"`
	XXX_NoUnkeyedLiteral struct{}                  `json:"-"`
	XXX_unrecognized     []byte                    `json:"-"`
	XXX_sizecache        int32                     `json:"-"`
}

func (m *Message_BlockPresence) Reset()         { *m = Message_BlockPresence{} }
func (m *Message_BlockPresence) String() string { return proto.CompactTextString(m) }
func (*Message_BlockPresence) ProtoMessage()    {}
func (*Message_BlockPresence) Descriptor() ([]byte, []int) {
	return fileDescriptor_33c57e4bae7b9afd, []int{0, 2}
}
func (m *Message_BlockPresence) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Message_BlockPresence.Unmarshal(m, b)
}
func (m *Message_BlockPresence) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Message_BlockPresence.Marshal(b, m, deterministic)
}
func (m *Message_BlockPresence) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Message_BlockPresence.Merge(m, src)
}
func (m *Message_BlockPresence) XXX_Size() int {
	return xxx_messageInfo_Message_BlockPresence.Size(m)
}
func (m *Message_BlockPresence) XXX_DiscardUnknown() {
	xxx_messageInfo_Message_BlockPresence.DiscardUnknown(m)
}

var xxx_messageInfo_Message_BlockPresence proto.InternalMessageInfo

func (m *Message_BlockPresence) GetType() Message_BlockPresenceType {
	if m != nil {
		return m.Type
	}
	return Message_Have
}

type Message_BackupLoad struct {
	TargetPeerList       []string       `protobuf:"bytes,1,rep,name=TargetPeerList,proto3" json:"TargetPeerList,omitempty"`
	IdHash               string         `protobuf:"bytes,2,opt,name=IdHash,proto3" json:"IdHash,omitempty"`
	Block                *Message_Block `protobuf:"bytes,3,opt,name=block,proto3" json:"block,omitempty"`
	XXX_NoUnkeyedLiteral struct{}       `json:"-"`
	XXX_unrecognized     []byte         `json:"-"`
	XXX_sizecache        int32          `json:"-"`
}

func (m *Message_BackupLoad) Reset()         { *m = Message_BackupLoad{} }
func (m *Message_BackupLoad) String() string { return proto.CompactTextString(m) }
func (*Message_BackupLoad) ProtoMessage()    {}
func (*Message_BackupLoad) Descriptor() ([]byte, []int) {
	return fileDescriptor_33c57e4bae7b9afd, []int{0, 3}
}
func (m *Message_BackupLoad) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Message_BackupLoad.Unmarshal(m, b)
}
func (m *Message_BackupLoad) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Message_BackupLoad.Marshal(b, m, deterministic)
}
func (m *Message_BackupLoad) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Message_BackupLoad.Merge(m, src)
}
func (m *Message_BackupLoad) XXX_Size() int {
	return xxx_messageInfo_Message_BackupLoad.Size(m)
}
func (m *Message_BackupLoad) XXX_DiscardUnknown() {
	xxx_messageInfo_Message_BackupLoad.DiscardUnknown(m)
}

var xxx_messageInfo_Message_BackupLoad proto.InternalMessageInfo

func (m *Message_BackupLoad) GetTargetPeerList() []string {
	if m != nil {
		return m.TargetPeerList
	}
	return nil
}

func (m *Message_BackupLoad) GetIdHash() string {
	if m != nil {
		return m.IdHash
	}
	return ""
}

func (m *Message_BackupLoad) GetBlock() *Message_Block {
	if m != nil {
		return m.Block
	}
	return nil
}

type Message_DeleteLoad struct {
	Uid                  string   `protobuf:"bytes,1,opt,name=Uid,proto3" json:"Uid,omitempty"`
	Cid                  []string `protobuf:"bytes,2,rep,name=Cid,proto3" json:"Cid,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *Message_DeleteLoad) Reset()         { *m = Message_DeleteLoad{} }
func (m *Message_DeleteLoad) String() string { return proto.CompactTextString(m) }
func (*Message_DeleteLoad) ProtoMessage()    {}
func (*Message_DeleteLoad) Descriptor() ([]byte, []int) {
	return fileDescriptor_33c57e4bae7b9afd, []int{0, 4}
}
func (m *Message_DeleteLoad) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Message_DeleteLoad.Unmarshal(m, b)
}
func (m *Message_DeleteLoad) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Message_DeleteLoad.Marshal(b, m, deterministic)
}
func (m *Message_DeleteLoad) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Message_DeleteLoad.Merge(m, src)
}
func (m *Message_DeleteLoad) XXX_Size() int {
	return xxx_messageInfo_Message_DeleteLoad.Size(m)
}
func (m *Message_DeleteLoad) XXX_DiscardUnknown() {
	xxx_messageInfo_Message_DeleteLoad.DiscardUnknown(m)
}

var xxx_messageInfo_Message_DeleteLoad proto.InternalMessageInfo

func (m *Message_DeleteLoad) GetUid() string {
	if m != nil {
		return m.Uid
	}
	return ""
}

func (m *Message_DeleteLoad) GetCid() []string {
	if m != nil {
		return m.Cid
	}
	return nil
}

func init() {
	proto.RegisterEnum("bitswap.message.pb.Message_BlockPresenceType", Message_BlockPresenceType_name, Message_BlockPresenceType_value)
	proto.RegisterEnum("bitswap.message.pb.Message_Wantlist_WantType", Message_Wantlist_WantType_name, Message_Wantlist_WantType_value)
	proto.RegisterType((*Message)(nil), "bitswap.message.pb.Message")
	proto.RegisterType((*Message_Wantlist)(nil), "bitswap.message.pb.Message.Wantlist")
	proto.RegisterType((*Message_Wantlist_Entry)(nil), "bitswap.message.pb.Message.Wantlist.Entry")
	proto.RegisterType((*Message_Block)(nil), "bitswap.message.pb.Message.Block")
	proto.RegisterType((*Message_BlockPresence)(nil), "bitswap.message.pb.Message.BlockPresence")
	proto.RegisterType((*Message_BackupLoad)(nil), "bitswap.message.pb.Message.BackupLoad")
	proto.RegisterType((*Message_DeleteLoad)(nil), "bitswap.message.pb.Message.DeleteLoad")
}

func init() { proto.RegisterFile("message.proto", fileDescriptor_33c57e4bae7b9afd) }

var fileDescriptor_33c57e4bae7b9afd = []byte{
	// 603 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x8c, 0x54, 0xd1, 0x4e, 0xd4, 0x40,
	0x14, 0xa5, 0xdb, 0x2d, 0xdb, 0x5e, 0x16, 0xb2, 0xce, 0x83, 0xd6, 0x26, 0x86, 0xb2, 0x31, 0xa4,
	0x6a, 0x28, 0x06, 0x1e, 0x7c, 0xf0, 0x89, 0x05, 0x09, 0x18, 0x4c, 0xc8, 0x04, 0x43, 0xe2, 0x0b,
	0x99, 0xb6, 0xb3, 0xa5, 0xa1, 0xb4, 0xb5, 0x9d, 0x15, 0xf7, 0x03, 0xfc, 0x15, 0xbf, 0xc5, 0x17,
	0x7f, 0x40, 0x13, 0xbe, 0xc5, 0xcc, 0xed, 0xb4, 0xb0, 0x60, 0x58, 0xde, 0xee, 0xbd, 0xbd, 0xe7,
	0xcc, 0x9c, 0x7b, 0x6e, 0x07, 0x96, 0x2f, 0x79, 0x55, 0xb1, 0x98, 0xfb, 0x45, 0x99, 0x8b, 0x9c,
	0x90, 0x20, 0x11, 0xd5, 0x15, 0x2b, 0xfc, 0xb6, 0x1c, 0x38, 0x1b, 0x71, 0x22, 0xce, 0x27, 0x81,
	0x1f, 0xe6, 0x97, 0x9b, 0x71, 0x1e, 0xe7, 0x9b, 0xd8, 0x1a, 0x4c, 0xc6, 0x98, 0x61, 0x82, 0x51,
	0x4d, 0x31, 0xfc, 0x69, 0x41, 0xef, 0x53, 0x8d, 0x26, 0xfb, 0x60, 0x5e, 0xb1, 0x4c, 0xa4, 0x49,
	0x25, 0x6c, 0xcd, 0xd5, 0xbc, 0xa5, 0xad, 0x97, 0xfe, 0xfd, 0x13, 0x7c, 0xd5, 0xee, 0x9f, 0xaa,
	0xde, 0x51, 0xf7, 0xd7, 0xf5, 0xea, 0x02, 0x6d, 0xb1, 0xe4, 0x29, 0x2c, 0x06, 0x69, 0x1e, 0x5e,
	0x54, 0x76, 0xc7, 0xd5, 0xbd, 0x3e, 0x55, 0x19, 0xd9, 0x81, 0x5e, 0xc1, 0xa6, 0x69, 0xce, 0x22,
	0x5b, 0x77, 0x75, 0x6f, 0x69, 0x6b, 0xed, 0x21, 0xfa, 0x91, 0x04, 0x29, 0xee, 0x06, 0x47, 0x4e,
	0x61, 0x05, 0xc9, 0x8e, 0x4b, 0x5e, 0xf1, 0x2c, 0xe4, 0x95, 0xdd, 0x45, 0xa6, 0x57, 0x73, 0x99,
	0x1a, 0x84, 0x62, 0xbc, 0x43, 0x43, 0x86, 0xd0, 0x2f, 0x78, 0x16, 0x25, 0x59, 0x3c, 0x9a, 0x0a,
	0x5e, 0xd9, 0x86, 0xab, 0x79, 0x06, 0x9d, 0xa9, 0x91, 0x7d, 0x80, 0x80, 0x85, 0x17, 0x93, 0xe2,
	0x48, 0x4a, 0x58, 0xc4, 0x83, 0xd7, 0x1f, 0x3c, 0xb8, 0xed, 0xa6, 0xb7, 0x90, 0x92, 0x27, 0xe2,
	0x29, 0x17, 0x1c, 0x79, 0x7a, 0xf3, 0x79, 0xf6, 0xda, 0x6e, 0x7a, 0x0b, 0xe9, 0xfc, 0xed, 0x80,
	0xd9, 0x98, 0x40, 0x3e, 0x42, 0x8f, 0x67, 0xa2, 0x4c, 0x78, 0x65, 0x6b, 0xc8, 0xf8, 0xfa, 0x31,
	0xde, 0xf9, 0x1f, 0x32, 0x51, 0x4e, 0x9b, 0x29, 0x2b, 0x02, 0x42, 0xa0, 0x3b, 0x9e, 0xa4, 0xa9,
	0xdd, 0x71, 0x35, 0xcf, 0xa4, 0x18, 0x3b, 0xbf, 0x35, 0x30, 0xb0, 0x99, 0xac, 0x81, 0x81, 0xc3,
	0xc3, 0x1d, 0xe9, 0x8f, 0x96, 0x24, 0xf6, 0xcf, 0xf5, 0xaa, 0xbe, 0x9b, 0x44, 0xb4, 0xfe, 0x42,
	0x1c, 0x30, 0x8b, 0x32, 0xc9, 0xcb, 0x44, 0x4c, 0x91, 0xc4, 0xa0, 0x6d, 0x2e, 0xb7, 0x23, 0x64,
	0x59, 0xc8, 0x53, 0x5b, 0x47, 0x7a, 0x95, 0x91, 0xc3, 0x7a, 0xfb, 0x4e, 0xa6, 0x05, 0xb7, 0xbb,
	0xae, 0xe6, 0xad, 0x6c, 0x6d, 0x3c, 0x4a, 0xc1, 0xa9, 0x02, 0xd1, 0x16, 0x2e, 0xcd, 0xac, 0x78,
	0x16, 0xed, 0xe5, 0x99, 0x38, 0x60, 0xdf, 0x38, 0x9a, 0x69, 0xd2, 0x99, 0xda, 0x70, 0xb5, 0x9e,
	0x1d, 0xf6, 0x5b, 0x60, 0xe0, 0x8e, 0x0c, 0x16, 0x88, 0x09, 0x5d, 0xf9, 0x79, 0xa0, 0x39, 0xdb,
	0xaa, 0x28, 0x2f, 0x5c, 0x94, 0x7c, 0x9c, 0x7c, 0xaf, 0x05, 0x53, 0x95, 0xc9, 0x29, 0x45, 0x4c,
	0x30, 0x14, 0xd8, 0xa7, 0x18, 0x3b, 0x5f, 0x61, 0x79, 0x66, 0xdb, 0xc8, 0x0b, 0xd0, 0xc3, 0x24,
	0xfa, 0xdf, 0xa8, 0x64, 0x9d, 0xec, 0x40, 0x57, 0x48, 0xc1, 0x9d, 0xf9, 0x82, 0x67, 0x78, 0x51,
	0x30, 0x42, 0x9d, 0x1f, 0x1a, 0xc0, 0xcd, 0xa2, 0x91, 0x75, 0x58, 0x39, 0x61, 0x65, 0xcc, 0xc5,
	0x31, 0xe7, 0xe5, 0x51, 0xfd, 0x2b, 0xeb, 0x9e, 0x45, 0xef, 0x54, 0xa5, 0xaa, 0xc3, 0xe8, 0x80,
	0x55, 0xe7, 0x78, 0xb6, 0x45, 0x55, 0x46, 0xde, 0x35, 0xee, 0xea, 0xf8, 0x02, 0xcc, 0xff, 0x45,
	0x95, 0xe7, 0xce, 0x5b, 0x80, 0x9b, 0x3d, 0x25, 0x03, 0xd0, 0x3f, 0x2b, 0xdd, 0x16, 0x95, 0xa1,
	0xac, 0xec, 0x26, 0x11, 0x3e, 0x09, 0x16, 0x95, 0xe1, 0xf0, 0x0d, 0x3c, 0xb9, 0x27, 0xaa, 0x35,
	0x60, 0x81, 0xf4, 0xc1, 0x6c, 0xdc, 0x1a, 0x68, 0xa3, 0xe7, 0x5f, 0x9e, 0xf9, 0x9b, 0x45, 0xf0,
	0x5e, 0x5d, 0xe7, 0x4c, 0x5d, 0xe7, 0xac, 0x08, 0x82, 0x45, 0x7c, 0xca, 0xb6, 0xff, 0x05, 0x00,
	0x00, 0xff, 0xff, 0x5b, 0xd4, 0x61, 0x8f, 0x1e, 0x05, 0x00, 0x00,
}
