// Code generated by protoc-gen-go. DO NOT EDIT.
// source: common.proto

package messages

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
const _ = proto.ProtoPackageIsVersion3 // please upgrade the proto package

type PubResponse struct {
	Pub                  string   `protobuf:"bytes,1,opt,name=pub,proto3" json:"pub,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *PubResponse) Reset()         { *m = PubResponse{} }
func (m *PubResponse) String() string { return proto.CompactTextString(m) }
func (*PubResponse) ProtoMessage()    {}
func (*PubResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_555bd8c177793206, []int{0}
}

func (m *PubResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_PubResponse.Unmarshal(m, b)
}
func (m *PubResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_PubResponse.Marshal(b, m, deterministic)
}
func (m *PubResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_PubResponse.Merge(m, src)
}
func (m *PubResponse) XXX_Size() int {
	return xxx_messageInfo_PubResponse.Size(m)
}
func (m *PubResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_PubResponse.DiscardUnknown(m)
}

var xxx_messageInfo_PubResponse proto.InternalMessageInfo

func (m *PubResponse) GetPub() string {
	if m != nil {
		return m.Pub
	}
	return ""
}

type RootFingerprintRequest struct {
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *RootFingerprintRequest) Reset()         { *m = RootFingerprintRequest{} }
func (m *RootFingerprintRequest) String() string { return proto.CompactTextString(m) }
func (*RootFingerprintRequest) ProtoMessage()    {}
func (*RootFingerprintRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_555bd8c177793206, []int{1}
}

func (m *RootFingerprintRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_RootFingerprintRequest.Unmarshal(m, b)
}
func (m *RootFingerprintRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_RootFingerprintRequest.Marshal(b, m, deterministic)
}
func (m *RootFingerprintRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_RootFingerprintRequest.Merge(m, src)
}
func (m *RootFingerprintRequest) XXX_Size() int {
	return xxx_messageInfo_RootFingerprintRequest.Size(m)
}
func (m *RootFingerprintRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_RootFingerprintRequest.DiscardUnknown(m)
}

var xxx_messageInfo_RootFingerprintRequest proto.InternalMessageInfo

type RootFingerprintResponse struct {
	Fingerprint          []byte   `protobuf:"bytes,1,opt,name=fingerprint,proto3" json:"fingerprint,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *RootFingerprintResponse) Reset()         { *m = RootFingerprintResponse{} }
func (m *RootFingerprintResponse) String() string { return proto.CompactTextString(m) }
func (*RootFingerprintResponse) ProtoMessage()    {}
func (*RootFingerprintResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_555bd8c177793206, []int{2}
}

func (m *RootFingerprintResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_RootFingerprintResponse.Unmarshal(m, b)
}
func (m *RootFingerprintResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_RootFingerprintResponse.Marshal(b, m, deterministic)
}
func (m *RootFingerprintResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_RootFingerprintResponse.Merge(m, src)
}
func (m *RootFingerprintResponse) XXX_Size() int {
	return xxx_messageInfo_RootFingerprintResponse.Size(m)
}
func (m *RootFingerprintResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_RootFingerprintResponse.DiscardUnknown(m)
}

var xxx_messageInfo_RootFingerprintResponse proto.InternalMessageInfo

func (m *RootFingerprintResponse) GetFingerprint() []byte {
	if m != nil {
		return m.Fingerprint
	}
	return nil
}

// See https://github.com/bitcoin/bips/blob/master/bip-0032.mediawiki.
// version field dropped as it will set dynamically based on the context (xpub, ypub, etc.).
type XPub struct {
	Depth                []byte   `protobuf:"bytes,1,opt,name=depth,proto3" json:"depth,omitempty"`
	ParentFingerprint    []byte   `protobuf:"bytes,2,opt,name=parent_fingerprint,json=parentFingerprint,proto3" json:"parent_fingerprint,omitempty"`
	ChildNum             uint32   `protobuf:"varint,3,opt,name=child_num,json=childNum,proto3" json:"child_num,omitempty"`
	ChainCode            []byte   `protobuf:"bytes,4,opt,name=chain_code,json=chainCode,proto3" json:"chain_code,omitempty"`
	PublicKey            []byte   `protobuf:"bytes,5,opt,name=public_key,json=publicKey,proto3" json:"public_key,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *XPub) Reset()         { *m = XPub{} }
func (m *XPub) String() string { return proto.CompactTextString(m) }
func (*XPub) ProtoMessage()    {}
func (*XPub) Descriptor() ([]byte, []int) {
	return fileDescriptor_555bd8c177793206, []int{3}
}

func (m *XPub) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_XPub.Unmarshal(m, b)
}
func (m *XPub) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_XPub.Marshal(b, m, deterministic)
}
func (m *XPub) XXX_Merge(src proto.Message) {
	xxx_messageInfo_XPub.Merge(m, src)
}
func (m *XPub) XXX_Size() int {
	return xxx_messageInfo_XPub.Size(m)
}
func (m *XPub) XXX_DiscardUnknown() {
	xxx_messageInfo_XPub.DiscardUnknown(m)
}

var xxx_messageInfo_XPub proto.InternalMessageInfo

func (m *XPub) GetDepth() []byte {
	if m != nil {
		return m.Depth
	}
	return nil
}

func (m *XPub) GetParentFingerprint() []byte {
	if m != nil {
		return m.ParentFingerprint
	}
	return nil
}

func (m *XPub) GetChildNum() uint32 {
	if m != nil {
		return m.ChildNum
	}
	return 0
}

func (m *XPub) GetChainCode() []byte {
	if m != nil {
		return m.ChainCode
	}
	return nil
}

func (m *XPub) GetPublicKey() []byte {
	if m != nil {
		return m.PublicKey
	}
	return nil
}

// This message exists for use in oneof or repeated fields, where one can't inline `repeated uint32` due to protobuf rules.
type Keypath struct {
	Keypath              []uint32 `protobuf:"varint,1,rep,packed,name=keypath,proto3" json:"keypath,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *Keypath) Reset()         { *m = Keypath{} }
func (m *Keypath) String() string { return proto.CompactTextString(m) }
func (*Keypath) ProtoMessage()    {}
func (*Keypath) Descriptor() ([]byte, []int) {
	return fileDescriptor_555bd8c177793206, []int{4}
}

func (m *Keypath) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Keypath.Unmarshal(m, b)
}
func (m *Keypath) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Keypath.Marshal(b, m, deterministic)
}
func (m *Keypath) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Keypath.Merge(m, src)
}
func (m *Keypath) XXX_Size() int {
	return xxx_messageInfo_Keypath.Size(m)
}
func (m *Keypath) XXX_DiscardUnknown() {
	xxx_messageInfo_Keypath.DiscardUnknown(m)
}

var xxx_messageInfo_Keypath proto.InternalMessageInfo

func (m *Keypath) GetKeypath() []uint32 {
	if m != nil {
		return m.Keypath
	}
	return nil
}

func init() {
	proto.RegisterType((*PubResponse)(nil), "shiftcrypto.bitbox02.PubResponse")
	proto.RegisterType((*RootFingerprintRequest)(nil), "shiftcrypto.bitbox02.RootFingerprintRequest")
	proto.RegisterType((*RootFingerprintResponse)(nil), "shiftcrypto.bitbox02.RootFingerprintResponse")
	proto.RegisterType((*XPub)(nil), "shiftcrypto.bitbox02.XPub")
	proto.RegisterType((*Keypath)(nil), "shiftcrypto.bitbox02.Keypath")
}

func init() { proto.RegisterFile("common.proto", fileDescriptor_555bd8c177793206) }

var fileDescriptor_555bd8c177793206 = []byte{
	// 267 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x64, 0x90, 0x41, 0x4b, 0xc3, 0x30,
	0x14, 0x80, 0xa9, 0xdb, 0x9c, 0x7d, 0xdb, 0x40, 0xc3, 0xd0, 0x80, 0x88, 0xa5, 0x5e, 0x7a, 0x71,
	0x88, 0x1e, 0x3d, 0x0a, 0x5e, 0x06, 0x32, 0x72, 0xf2, 0x56, 0x9a, 0xf4, 0xcd, 0x86, 0xad, 0x79,
	0xb1, 0x4d, 0xc0, 0xfe, 0x21, 0x7f, 0xa7, 0x2c, 0x9d, 0x58, 0xf0, 0xf6, 0xde, 0xf7, 0x3d, 0xbe,
	0x40, 0x60, 0xae, 0xa8, 0xae, 0xc9, 0xac, 0x6c, 0x43, 0x8e, 0xd8, 0xb2, 0xad, 0xf4, 0xd6, 0xa9,
	0xa6, 0xb3, 0x8e, 0x56, 0x52, 0x3b, 0x49, 0x5f, 0x0f, 0x8f, 0xe9, 0x2d, 0xcc, 0x36, 0x5e, 0x0a,
	0x6c, 0x2d, 0x99, 0x16, 0xd9, 0x39, 0x8c, 0xac, 0x97, 0x3c, 0x4a, 0xa2, 0x2c, 0x16, 0x87, 0x31,
	0xe5, 0x70, 0x29, 0x88, 0xdc, 0xab, 0x36, 0x1f, 0xd8, 0xd8, 0x46, 0x1b, 0x27, 0xf0, 0xd3, 0x63,
	0xeb, 0xd2, 0x67, 0xb8, 0xfa, 0x67, 0x8e, 0x99, 0x04, 0x66, 0xdb, 0x3f, 0x1c, 0x72, 0x73, 0x31,
	0x44, 0xe9, 0x77, 0x04, 0xe3, 0xf7, 0x8d, 0x97, 0x6c, 0x09, 0x93, 0x12, 0xad, 0xab, 0x8e, 0x47,
	0xfd, 0xc2, 0xee, 0x81, 0xd9, 0xa2, 0x41, 0xe3, 0xf2, 0x61, 0xe7, 0x24, 0x9c, 0x5c, 0xf4, 0x66,
	0xf0, 0x2e, 0xbb, 0x86, 0x58, 0x55, 0x7a, 0x5f, 0xe6, 0xc6, 0xd7, 0x7c, 0x94, 0x44, 0xd9, 0x42,
	0x9c, 0x05, 0xf0, 0xe6, 0x6b, 0x76, 0x03, 0xa0, 0xaa, 0x42, 0x9b, 0x5c, 0x51, 0x89, 0x7c, 0x1c,
	0x1a, 0x71, 0x20, 0x2f, 0x54, 0xe2, 0x41, 0x5b, 0x2f, 0xf7, 0x5a, 0xe5, 0x3b, 0xec, 0xf8, 0xa4,
	0xd7, 0x3d, 0x59, 0x63, 0x97, 0xde, 0xc1, 0x74, 0x8d, 0x9d, 0x2d, 0x5c, 0xc5, 0x38, 0x4c, 0x77,
	0xfd, 0xc8, 0xa3, 0x64, 0x94, 0x2d, 0xc4, 0xef, 0x2a, 0x4f, 0xc3, 0x17, 0x3f, 0xfd, 0x04, 0x00,
	0x00, 0xff, 0xff, 0x34, 0x35, 0x29, 0x2b, 0x72, 0x01, 0x00, 0x00,
}
