// Code generated by protoc-gen-go. DO NOT EDIT.
// source: edge.proto

package gon2n

import (
	context "context"
	fmt "fmt"
	proto "github.com/golang/protobuf/proto"
	grpc "google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
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

type EdgeManagerCreateArgs struct {
	AllowP2P             bool     `protobuf:"varint,1,opt,name=AllowP2P,proto3" json:"AllowP2P,omitempty"`
	AllowRouting         bool     `protobuf:"varint,2,opt,name=AllowRouting,proto3" json:"AllowRouting,omitempty"`
	CommunityName        string   `protobuf:"bytes,3,opt,name=CommunityName,proto3" json:"CommunityName,omitempty"`
	DisablePMTUDiscovery bool     `protobuf:"varint,4,opt,name=DisablePMTUDiscovery,proto3" json:"DisablePMTUDiscovery,omitempty"`
	DisableMulticast     bool     `protobuf:"varint,5,opt,name=DisableMulticast,proto3" json:"DisableMulticast,omitempty"`
	DynamicIPMode        bool     `protobuf:"varint,6,opt,name=DynamicIPMode,proto3" json:"DynamicIPMode,omitempty"`
	EncryptionKey        string   `protobuf:"bytes,7,opt,name=EncryptionKey,proto3" json:"EncryptionKey,omitempty"`
	LocalPort            int64    `protobuf:"varint,8,opt,name=LocalPort,proto3" json:"LocalPort,omitempty"`
	ManagementPort       int64    `protobuf:"varint,9,opt,name=ManagementPort,proto3" json:"ManagementPort,omitempty"`
	RegisterInterval     int64    `protobuf:"varint,10,opt,name=RegisterInterval,proto3" json:"RegisterInterval,omitempty"`
	RegisterTTL          int64    `protobuf:"varint,11,opt,name=RegisterTTL,proto3" json:"RegisterTTL,omitempty"`
	SupernodeHostPort    string   `protobuf:"bytes,12,opt,name=SupernodeHostPort,proto3" json:"SupernodeHostPort,omitempty"`
	TypeOfService        int64    `protobuf:"varint,13,opt,name=TypeOfService,proto3" json:"TypeOfService,omitempty"`
	EncryptionMethod     int64    `protobuf:"varint,14,opt,name=EncryptionMethod,proto3" json:"EncryptionMethod,omitempty"`
	DeviceName           string   `protobuf:"bytes,15,opt,name=DeviceName,proto3" json:"DeviceName,omitempty"`
	AddressMode          string   `protobuf:"bytes,16,opt,name=AddressMode,proto3" json:"AddressMode,omitempty"`
	DeviceIP             string   `protobuf:"bytes,17,opt,name=DeviceIP,proto3" json:"DeviceIP,omitempty"`
	DeviceNetmask        string   `protobuf:"bytes,18,opt,name=DeviceNetmask,proto3" json:"DeviceNetmask,omitempty"`
	DeviceMACAddress     string   `protobuf:"bytes,19,opt,name=DeviceMACAddress,proto3" json:"DeviceMACAddress,omitempty"`
	MTU                  int64    `protobuf:"varint,20,opt,name=MTU,proto3" json:"MTU,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *EdgeManagerCreateArgs) Reset()         { *m = EdgeManagerCreateArgs{} }
func (m *EdgeManagerCreateArgs) String() string { return proto.CompactTextString(m) }
func (*EdgeManagerCreateArgs) ProtoMessage()    {}
func (*EdgeManagerCreateArgs) Descriptor() ([]byte, []int) {
	return fileDescriptor_cab1176173a95651, []int{0}
}

func (m *EdgeManagerCreateArgs) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_EdgeManagerCreateArgs.Unmarshal(m, b)
}
func (m *EdgeManagerCreateArgs) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_EdgeManagerCreateArgs.Marshal(b, m, deterministic)
}
func (m *EdgeManagerCreateArgs) XXX_Merge(src proto.Message) {
	xxx_messageInfo_EdgeManagerCreateArgs.Merge(m, src)
}
func (m *EdgeManagerCreateArgs) XXX_Size() int {
	return xxx_messageInfo_EdgeManagerCreateArgs.Size(m)
}
func (m *EdgeManagerCreateArgs) XXX_DiscardUnknown() {
	xxx_messageInfo_EdgeManagerCreateArgs.DiscardUnknown(m)
}

var xxx_messageInfo_EdgeManagerCreateArgs proto.InternalMessageInfo

func (m *EdgeManagerCreateArgs) GetAllowP2P() bool {
	if m != nil {
		return m.AllowP2P
	}
	return false
}

func (m *EdgeManagerCreateArgs) GetAllowRouting() bool {
	if m != nil {
		return m.AllowRouting
	}
	return false
}

func (m *EdgeManagerCreateArgs) GetCommunityName() string {
	if m != nil {
		return m.CommunityName
	}
	return ""
}

func (m *EdgeManagerCreateArgs) GetDisablePMTUDiscovery() bool {
	if m != nil {
		return m.DisablePMTUDiscovery
	}
	return false
}

func (m *EdgeManagerCreateArgs) GetDisableMulticast() bool {
	if m != nil {
		return m.DisableMulticast
	}
	return false
}

func (m *EdgeManagerCreateArgs) GetDynamicIPMode() bool {
	if m != nil {
		return m.DynamicIPMode
	}
	return false
}

func (m *EdgeManagerCreateArgs) GetEncryptionKey() string {
	if m != nil {
		return m.EncryptionKey
	}
	return ""
}

func (m *EdgeManagerCreateArgs) GetLocalPort() int64 {
	if m != nil {
		return m.LocalPort
	}
	return 0
}

func (m *EdgeManagerCreateArgs) GetManagementPort() int64 {
	if m != nil {
		return m.ManagementPort
	}
	return 0
}

func (m *EdgeManagerCreateArgs) GetRegisterInterval() int64 {
	if m != nil {
		return m.RegisterInterval
	}
	return 0
}

func (m *EdgeManagerCreateArgs) GetRegisterTTL() int64 {
	if m != nil {
		return m.RegisterTTL
	}
	return 0
}

func (m *EdgeManagerCreateArgs) GetSupernodeHostPort() string {
	if m != nil {
		return m.SupernodeHostPort
	}
	return ""
}

func (m *EdgeManagerCreateArgs) GetTypeOfService() int64 {
	if m != nil {
		return m.TypeOfService
	}
	return 0
}

func (m *EdgeManagerCreateArgs) GetEncryptionMethod() int64 {
	if m != nil {
		return m.EncryptionMethod
	}
	return 0
}

func (m *EdgeManagerCreateArgs) GetDeviceName() string {
	if m != nil {
		return m.DeviceName
	}
	return ""
}

func (m *EdgeManagerCreateArgs) GetAddressMode() string {
	if m != nil {
		return m.AddressMode
	}
	return ""
}

func (m *EdgeManagerCreateArgs) GetDeviceIP() string {
	if m != nil {
		return m.DeviceIP
	}
	return ""
}

func (m *EdgeManagerCreateArgs) GetDeviceNetmask() string {
	if m != nil {
		return m.DeviceNetmask
	}
	return ""
}

func (m *EdgeManagerCreateArgs) GetDeviceMACAddress() string {
	if m != nil {
		return m.DeviceMACAddress
	}
	return ""
}

func (m *EdgeManagerCreateArgs) GetMTU() int64 {
	if m != nil {
		return m.MTU
	}
	return 0
}

type EdgeManagerListArgs struct {
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *EdgeManagerListArgs) Reset()         { *m = EdgeManagerListArgs{} }
func (m *EdgeManagerListArgs) String() string { return proto.CompactTextString(m) }
func (*EdgeManagerListArgs) ProtoMessage()    {}
func (*EdgeManagerListArgs) Descriptor() ([]byte, []int) {
	return fileDescriptor_cab1176173a95651, []int{1}
}

func (m *EdgeManagerListArgs) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_EdgeManagerListArgs.Unmarshal(m, b)
}
func (m *EdgeManagerListArgs) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_EdgeManagerListArgs.Marshal(b, m, deterministic)
}
func (m *EdgeManagerListArgs) XXX_Merge(src proto.Message) {
	xxx_messageInfo_EdgeManagerListArgs.Merge(m, src)
}
func (m *EdgeManagerListArgs) XXX_Size() int {
	return xxx_messageInfo_EdgeManagerListArgs.Size(m)
}
func (m *EdgeManagerListArgs) XXX_DiscardUnknown() {
	xxx_messageInfo_EdgeManagerListArgs.DiscardUnknown(m)
}

var xxx_messageInfo_EdgeManagerListArgs proto.InternalMessageInfo

type EdgeManagerDeleteArgs struct {
	Id                   string   `protobuf:"bytes,1,opt,name=Id,proto3" json:"Id,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *EdgeManagerDeleteArgs) Reset()         { *m = EdgeManagerDeleteArgs{} }
func (m *EdgeManagerDeleteArgs) String() string { return proto.CompactTextString(m) }
func (*EdgeManagerDeleteArgs) ProtoMessage()    {}
func (*EdgeManagerDeleteArgs) Descriptor() ([]byte, []int) {
	return fileDescriptor_cab1176173a95651, []int{2}
}

func (m *EdgeManagerDeleteArgs) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_EdgeManagerDeleteArgs.Unmarshal(m, b)
}
func (m *EdgeManagerDeleteArgs) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_EdgeManagerDeleteArgs.Marshal(b, m, deterministic)
}
func (m *EdgeManagerDeleteArgs) XXX_Merge(src proto.Message) {
	xxx_messageInfo_EdgeManagerDeleteArgs.Merge(m, src)
}
func (m *EdgeManagerDeleteArgs) XXX_Size() int {
	return xxx_messageInfo_EdgeManagerDeleteArgs.Size(m)
}
func (m *EdgeManagerDeleteArgs) XXX_DiscardUnknown() {
	xxx_messageInfo_EdgeManagerDeleteArgs.DiscardUnknown(m)
}

var xxx_messageInfo_EdgeManagerDeleteArgs proto.InternalMessageInfo

func (m *EdgeManagerDeleteArgs) GetId() string {
	if m != nil {
		return m.Id
	}
	return ""
}

type EdgeManagerCreateReply struct {
	Id                   string   `protobuf:"bytes,1,opt,name=Id,proto3" json:"Id,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *EdgeManagerCreateReply) Reset()         { *m = EdgeManagerCreateReply{} }
func (m *EdgeManagerCreateReply) String() string { return proto.CompactTextString(m) }
func (*EdgeManagerCreateReply) ProtoMessage()    {}
func (*EdgeManagerCreateReply) Descriptor() ([]byte, []int) {
	return fileDescriptor_cab1176173a95651, []int{3}
}

func (m *EdgeManagerCreateReply) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_EdgeManagerCreateReply.Unmarshal(m, b)
}
func (m *EdgeManagerCreateReply) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_EdgeManagerCreateReply.Marshal(b, m, deterministic)
}
func (m *EdgeManagerCreateReply) XXX_Merge(src proto.Message) {
	xxx_messageInfo_EdgeManagerCreateReply.Merge(m, src)
}
func (m *EdgeManagerCreateReply) XXX_Size() int {
	return xxx_messageInfo_EdgeManagerCreateReply.Size(m)
}
func (m *EdgeManagerCreateReply) XXX_DiscardUnknown() {
	xxx_messageInfo_EdgeManagerCreateReply.DiscardUnknown(m)
}

var xxx_messageInfo_EdgeManagerCreateReply proto.InternalMessageInfo

func (m *EdgeManagerCreateReply) GetId() string {
	if m != nil {
		return m.Id
	}
	return ""
}

type EdgeManagerListReply struct {
	EdgesManaged         []*EdgeManaged `protobuf:"bytes,1,rep,name=EdgesManaged,proto3" json:"EdgesManaged,omitempty"`
	XXX_NoUnkeyedLiteral struct{}       `json:"-"`
	XXX_unrecognized     []byte         `json:"-"`
	XXX_sizecache        int32          `json:"-"`
}

func (m *EdgeManagerListReply) Reset()         { *m = EdgeManagerListReply{} }
func (m *EdgeManagerListReply) String() string { return proto.CompactTextString(m) }
func (*EdgeManagerListReply) ProtoMessage()    {}
func (*EdgeManagerListReply) Descriptor() ([]byte, []int) {
	return fileDescriptor_cab1176173a95651, []int{4}
}

func (m *EdgeManagerListReply) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_EdgeManagerListReply.Unmarshal(m, b)
}
func (m *EdgeManagerListReply) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_EdgeManagerListReply.Marshal(b, m, deterministic)
}
func (m *EdgeManagerListReply) XXX_Merge(src proto.Message) {
	xxx_messageInfo_EdgeManagerListReply.Merge(m, src)
}
func (m *EdgeManagerListReply) XXX_Size() int {
	return xxx_messageInfo_EdgeManagerListReply.Size(m)
}
func (m *EdgeManagerListReply) XXX_DiscardUnknown() {
	xxx_messageInfo_EdgeManagerListReply.DiscardUnknown(m)
}

var xxx_messageInfo_EdgeManagerListReply proto.InternalMessageInfo

func (m *EdgeManagerListReply) GetEdgesManaged() []*EdgeManaged {
	if m != nil {
		return m.EdgesManaged
	}
	return nil
}

type EdgeManaged struct {
	Id                   string   `protobuf:"bytes,1,opt,name=Id,proto3" json:"Id,omitempty"`
	AllowP2P             bool     `protobuf:"varint,2,opt,name=AllowP2P,proto3" json:"AllowP2P,omitempty"`
	AllowRouting         bool     `protobuf:"varint,3,opt,name=AllowRouting,proto3" json:"AllowRouting,omitempty"`
	CommunityName        string   `protobuf:"bytes,4,opt,name=CommunityName,proto3" json:"CommunityName,omitempty"`
	DisablePMTUDiscovery bool     `protobuf:"varint,5,opt,name=DisablePMTUDiscovery,proto3" json:"DisablePMTUDiscovery,omitempty"`
	DisableMulticast     bool     `protobuf:"varint,6,opt,name=DisableMulticast,proto3" json:"DisableMulticast,omitempty"`
	DynamicIPMode        bool     `protobuf:"varint,7,opt,name=DynamicIPMode,proto3" json:"DynamicIPMode,omitempty"`
	LocalPort            int64    `protobuf:"varint,8,opt,name=LocalPort,proto3" json:"LocalPort,omitempty"`
	ManagementPort       int64    `protobuf:"varint,9,opt,name=ManagementPort,proto3" json:"ManagementPort,omitempty"`
	RegisterInterval     int64    `protobuf:"varint,10,opt,name=RegisterInterval,proto3" json:"RegisterInterval,omitempty"`
	RegisterTTL          int64    `protobuf:"varint,11,opt,name=RegisterTTL,proto3" json:"RegisterTTL,omitempty"`
	SupernodeHostPort    string   `protobuf:"bytes,12,opt,name=SupernodeHostPort,proto3" json:"SupernodeHostPort,omitempty"`
	TypeOfService        int64    `protobuf:"varint,13,opt,name=TypeOfService,proto3" json:"TypeOfService,omitempty"`
	EncryptionMethod     int64    `protobuf:"varint,14,opt,name=EncryptionMethod,proto3" json:"EncryptionMethod,omitempty"`
	DeviceName           string   `protobuf:"bytes,15,opt,name=DeviceName,proto3" json:"DeviceName,omitempty"`
	AddressMode          string   `protobuf:"bytes,16,opt,name=AddressMode,proto3" json:"AddressMode,omitempty"`
	DeviceIP             string   `protobuf:"bytes,17,opt,name=DeviceIP,proto3" json:"DeviceIP,omitempty"`
	DeviceNetmask        string   `protobuf:"bytes,18,opt,name=DeviceNetmask,proto3" json:"DeviceNetmask,omitempty"`
	DeviceMACAddress     string   `protobuf:"bytes,19,opt,name=DeviceMACAddress,proto3" json:"DeviceMACAddress,omitempty"`
	MTU                  int64    `protobuf:"varint,20,opt,name=MTU,proto3" json:"MTU,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *EdgeManaged) Reset()         { *m = EdgeManaged{} }
func (m *EdgeManaged) String() string { return proto.CompactTextString(m) }
func (*EdgeManaged) ProtoMessage()    {}
func (*EdgeManaged) Descriptor() ([]byte, []int) {
	return fileDescriptor_cab1176173a95651, []int{5}
}

func (m *EdgeManaged) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_EdgeManaged.Unmarshal(m, b)
}
func (m *EdgeManaged) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_EdgeManaged.Marshal(b, m, deterministic)
}
func (m *EdgeManaged) XXX_Merge(src proto.Message) {
	xxx_messageInfo_EdgeManaged.Merge(m, src)
}
func (m *EdgeManaged) XXX_Size() int {
	return xxx_messageInfo_EdgeManaged.Size(m)
}
func (m *EdgeManaged) XXX_DiscardUnknown() {
	xxx_messageInfo_EdgeManaged.DiscardUnknown(m)
}

var xxx_messageInfo_EdgeManaged proto.InternalMessageInfo

func (m *EdgeManaged) GetId() string {
	if m != nil {
		return m.Id
	}
	return ""
}

func (m *EdgeManaged) GetAllowP2P() bool {
	if m != nil {
		return m.AllowP2P
	}
	return false
}

func (m *EdgeManaged) GetAllowRouting() bool {
	if m != nil {
		return m.AllowRouting
	}
	return false
}

func (m *EdgeManaged) GetCommunityName() string {
	if m != nil {
		return m.CommunityName
	}
	return ""
}

func (m *EdgeManaged) GetDisablePMTUDiscovery() bool {
	if m != nil {
		return m.DisablePMTUDiscovery
	}
	return false
}

func (m *EdgeManaged) GetDisableMulticast() bool {
	if m != nil {
		return m.DisableMulticast
	}
	return false
}

func (m *EdgeManaged) GetDynamicIPMode() bool {
	if m != nil {
		return m.DynamicIPMode
	}
	return false
}

func (m *EdgeManaged) GetLocalPort() int64 {
	if m != nil {
		return m.LocalPort
	}
	return 0
}

func (m *EdgeManaged) GetManagementPort() int64 {
	if m != nil {
		return m.ManagementPort
	}
	return 0
}

func (m *EdgeManaged) GetRegisterInterval() int64 {
	if m != nil {
		return m.RegisterInterval
	}
	return 0
}

func (m *EdgeManaged) GetRegisterTTL() int64 {
	if m != nil {
		return m.RegisterTTL
	}
	return 0
}

func (m *EdgeManaged) GetSupernodeHostPort() string {
	if m != nil {
		return m.SupernodeHostPort
	}
	return ""
}

func (m *EdgeManaged) GetTypeOfService() int64 {
	if m != nil {
		return m.TypeOfService
	}
	return 0
}

func (m *EdgeManaged) GetEncryptionMethod() int64 {
	if m != nil {
		return m.EncryptionMethod
	}
	return 0
}

func (m *EdgeManaged) GetDeviceName() string {
	if m != nil {
		return m.DeviceName
	}
	return ""
}

func (m *EdgeManaged) GetAddressMode() string {
	if m != nil {
		return m.AddressMode
	}
	return ""
}

func (m *EdgeManaged) GetDeviceIP() string {
	if m != nil {
		return m.DeviceIP
	}
	return ""
}

func (m *EdgeManaged) GetDeviceNetmask() string {
	if m != nil {
		return m.DeviceNetmask
	}
	return ""
}

func (m *EdgeManaged) GetDeviceMACAddress() string {
	if m != nil {
		return m.DeviceMACAddress
	}
	return ""
}

func (m *EdgeManaged) GetMTU() int64 {
	if m != nil {
		return m.MTU
	}
	return 0
}

type EdgeManagerDeleteReply struct {
	Id                   string   `protobuf:"bytes,1,opt,name=Id,proto3" json:"Id,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *EdgeManagerDeleteReply) Reset()         { *m = EdgeManagerDeleteReply{} }
func (m *EdgeManagerDeleteReply) String() string { return proto.CompactTextString(m) }
func (*EdgeManagerDeleteReply) ProtoMessage()    {}
func (*EdgeManagerDeleteReply) Descriptor() ([]byte, []int) {
	return fileDescriptor_cab1176173a95651, []int{6}
}

func (m *EdgeManagerDeleteReply) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_EdgeManagerDeleteReply.Unmarshal(m, b)
}
func (m *EdgeManagerDeleteReply) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_EdgeManagerDeleteReply.Marshal(b, m, deterministic)
}
func (m *EdgeManagerDeleteReply) XXX_Merge(src proto.Message) {
	xxx_messageInfo_EdgeManagerDeleteReply.Merge(m, src)
}
func (m *EdgeManagerDeleteReply) XXX_Size() int {
	return xxx_messageInfo_EdgeManagerDeleteReply.Size(m)
}
func (m *EdgeManagerDeleteReply) XXX_DiscardUnknown() {
	xxx_messageInfo_EdgeManagerDeleteReply.DiscardUnknown(m)
}

var xxx_messageInfo_EdgeManagerDeleteReply proto.InternalMessageInfo

func (m *EdgeManagerDeleteReply) GetId() string {
	if m != nil {
		return m.Id
	}
	return ""
}

func init() {
	proto.RegisterType((*EdgeManagerCreateArgs)(nil), "gon2n.EdgeManagerCreateArgs")
	proto.RegisterType((*EdgeManagerListArgs)(nil), "gon2n.EdgeManagerListArgs")
	proto.RegisterType((*EdgeManagerDeleteArgs)(nil), "gon2n.EdgeManagerDeleteArgs")
	proto.RegisterType((*EdgeManagerCreateReply)(nil), "gon2n.EdgeManagerCreateReply")
	proto.RegisterType((*EdgeManagerListReply)(nil), "gon2n.EdgeManagerListReply")
	proto.RegisterType((*EdgeManaged)(nil), "gon2n.EdgeManaged")
	proto.RegisterType((*EdgeManagerDeleteReply)(nil), "gon2n.EdgeManagerDeleteReply")
}

func init() { proto.RegisterFile("edge.proto", fileDescriptor_cab1176173a95651) }

var fileDescriptor_cab1176173a95651 = []byte{
	// 603 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xec, 0x95, 0x5f, 0x4f, 0xdb, 0x3c,
	0x14, 0xc6, 0x69, 0x4b, 0x0b, 0x3d, 0x05, 0x5e, 0x30, 0xf0, 0xca, 0xea, 0xd8, 0x54, 0x45, 0xd3,
	0x56, 0x4d, 0x13, 0x17, 0x9d, 0xb4, 0xfb, 0x8a, 0x22, 0x56, 0x8d, 0xb0, 0x28, 0x94, 0x0f, 0x10,
	0x92, 0xb3, 0x2c, 0x5a, 0x12, 0x57, 0x8e, 0xcb, 0x94, 0x4f, 0x3b, 0x69, 0x5f, 0x64, 0x93, 0x8f,
	0x29, 0x4d, 0x9a, 0x4c, 0x83, 0xdb, 0x69, 0x77, 0xc9, 0x73, 0x1e, 0x9f, 0x3f, 0xb1, 0x7f, 0x31,
	0x00, 0x06, 0x21, 0x9e, 0xce, 0xa5, 0x50, 0x82, 0xb5, 0x43, 0x91, 0x8e, 0x52, 0xeb, 0x67, 0x1b,
	0x8e, 0xcf, 0x83, 0x10, 0x6d, 0x2f, 0xf5, 0x42, 0x94, 0x67, 0x12, 0x3d, 0x85, 0x63, 0x19, 0x66,
	0xac, 0x0f, 0xdb, 0xe3, 0x38, 0x16, 0xdf, 0x9c, 0x91, 0xc3, 0x1b, 0x83, 0xc6, 0x70, 0xdb, 0x7d,
	0x78, 0x67, 0x16, 0xec, 0xd0, 0xb3, 0x2b, 0x16, 0x2a, 0x4a, 0x43, 0xde, 0xa4, 0x78, 0x49, 0x63,
	0x2f, 0x61, 0xf7, 0x4c, 0x24, 0xc9, 0x22, 0x8d, 0x54, 0x7e, 0xe5, 0x25, 0xc8, 0x5b, 0x83, 0xc6,
	0xb0, 0xeb, 0x96, 0x45, 0x36, 0x82, 0xa3, 0x49, 0x94, 0x79, 0xb7, 0x31, 0x3a, 0xf6, 0xec, 0x66,
	0x12, 0x65, 0xbe, 0xb8, 0x43, 0x99, 0xf3, 0x4d, 0xca, 0x58, 0x1b, 0x63, 0x6f, 0x60, 0xff, 0x5e,
	0xb7, 0x17, 0xb1, 0x8a, 0x7c, 0x2f, 0x53, 0xbc, 0x4d, 0xfe, 0x8a, 0xae, 0xbb, 0x98, 0xe4, 0xa9,
	0x97, 0x44, 0xfe, 0xd4, 0xb1, 0x45, 0x80, 0xbc, 0x43, 0xc6, 0xb2, 0xa8, 0x5d, 0xe7, 0xa9, 0x2f,
	0xf3, 0xb9, 0x8a, 0x44, 0xfa, 0x11, 0x73, 0xbe, 0x65, 0x7a, 0x2d, 0x89, 0xec, 0x04, 0xba, 0x97,
	0xc2, 0xf7, 0x62, 0x47, 0x48, 0xc5, 0xb7, 0x07, 0x8d, 0x61, 0xcb, 0x5d, 0x09, 0xec, 0x15, 0xec,
	0x99, 0x8f, 0x98, 0x60, 0xaa, 0xc8, 0xd2, 0x25, 0xcb, 0x9a, 0xaa, 0xbb, 0x77, 0x31, 0x8c, 0x32,
	0x85, 0x72, 0x9a, 0x2a, 0x94, 0x77, 0x5e, 0xcc, 0x81, 0x9c, 0x15, 0x9d, 0x0d, 0xa0, 0xb7, 0xd4,
	0x66, 0xb3, 0x4b, 0xde, 0x23, 0x5b, 0x51, 0x62, 0x6f, 0xe1, 0xe0, 0x7a, 0x31, 0x47, 0x99, 0x8a,
	0x00, 0x3f, 0x88, 0xcc, 0x14, 0xde, 0xa1, 0xee, 0xab, 0x01, 0x3d, 0xe7, 0x2c, 0x9f, 0xe3, 0xa7,
	0xcf, 0xd7, 0x28, 0xef, 0x22, 0x1f, 0xf9, 0x2e, 0x65, 0x2c, 0x8b, 0xba, 0xc3, 0xd5, 0xe0, 0x36,
	0xaa, 0x2f, 0x22, 0xe0, 0x7b, 0xa6, 0xc3, 0x75, 0x9d, 0xbd, 0x00, 0x98, 0xa0, 0x5e, 0x45, 0x5b,
	0xfc, 0x1f, 0x15, 0x2e, 0x28, 0x7a, 0x82, 0x71, 0x10, 0x48, 0xcc, 0x32, 0xfa, 0xfa, 0xfb, 0x64,
	0x28, 0x4a, 0xfa, 0x9c, 0x19, 0xff, 0xd4, 0xe1, 0x07, 0x14, 0x7e, 0x78, 0xa7, 0xdd, 0x33, 0xb9,
	0x50, 0x25, 0x5e, 0xf6, 0x95, 0x33, 0xb3, 0x2f, 0x25, 0x91, 0xce, 0x03, 0x09, 0xf6, 0xf8, 0xec,
	0x3e, 0x33, 0x3f, 0x24, 0x63, 0x45, 0x67, 0xfb, 0xd0, 0xb2, 0x67, 0x37, 0xfc, 0x88, 0xc6, 0xd1,
	0x8f, 0xd6, 0x31, 0x1c, 0x16, 0x00, 0xb8, 0x8c, 0x32, 0xa5, 0x8f, 0xbf, 0xf5, 0xba, 0xc4, 0xc5,
	0x04, 0x63, 0xbc, 0xe7, 0x62, 0x0f, 0x9a, 0xd3, 0x80, 0x88, 0xe8, 0xba, 0xcd, 0x69, 0x60, 0x0d,
	0xe1, 0xff, 0x0a, 0x40, 0x2e, 0xce, 0xe3, 0xbc, 0xe2, 0xbc, 0x82, 0xa3, 0xb5, 0x4a, 0xc6, 0xf7,
	0x1e, 0x76, 0xb4, 0x9e, 0x99, 0x80, 0x5e, 0xd1, 0x1a, 0xf6, 0x46, 0xec, 0x94, 0x08, 0x3d, 0x5d,
	0x2d, 0x09, 0xdc, 0x92, 0xcf, 0xfa, 0xde, 0x86, 0x5e, 0x21, 0xba, 0x5e, 0xaf, 0x44, 0x70, 0xf3,
	0x0f, 0x04, 0xb7, 0x1e, 0x43, 0xf0, 0xe6, 0x53, 0x08, 0x6e, 0x3f, 0x91, 0xe0, 0xce, 0x63, 0x09,
	0xde, 0xaa, 0x23, 0xf8, 0x1f, 0x9b, 0x7f, 0x23, 0x9b, 0x65, 0xb6, 0x0c, 0x84, 0xb5, 0x6c, 0x8d,
	0x7e, 0x34, 0x8a, 0x2c, 0x48, 0x76, 0x01, 0x1d, 0x83, 0x22, 0x3b, 0xa9, 0x70, 0x54, 0xb8, 0xe5,
	0xfa, 0xcf, 0x7f, 0x17, 0xa5, 0x32, 0xd6, 0x06, 0x1b, 0xc3, 0xa6, 0x26, 0x95, 0xf5, 0xab, 0xc6,
	0xe5, 0xbf, 0xa2, 0xff, 0xac, 0x3e, 0xb6, 0x4c, 0x71, 0x01, 0x1d, 0xd3, 0x7a, 0x5d, 0x2f, 0xab,
	0x3f, 0x4b, 0x5d, 0x2f, 0x85, 0x91, 0xad, 0x8d, 0xdb, 0x0e, 0x5d, 0xdd, 0xef, 0x7e, 0x05, 0x00,
	0x00, 0xff, 0xff, 0x3e, 0x2b, 0x40, 0x3c, 0xc8, 0x07, 0x00, 0x00,
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConn

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion4

// EdgeManagerClient is the client API for EdgeManager service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type EdgeManagerClient interface {
	Create(ctx context.Context, in *EdgeManagerCreateArgs, opts ...grpc.CallOption) (*EdgeManagerCreateReply, error)
	List(ctx context.Context, in *EdgeManagerListArgs, opts ...grpc.CallOption) (*EdgeManagerListReply, error)
	Delete(ctx context.Context, in *EdgeManagerDeleteArgs, opts ...grpc.CallOption) (*EdgeManagerDeleteReply, error)
}

type edgeManagerClient struct {
	cc *grpc.ClientConn
}

func NewEdgeManagerClient(cc *grpc.ClientConn) EdgeManagerClient {
	return &edgeManagerClient{cc}
}

func (c *edgeManagerClient) Create(ctx context.Context, in *EdgeManagerCreateArgs, opts ...grpc.CallOption) (*EdgeManagerCreateReply, error) {
	out := new(EdgeManagerCreateReply)
	err := c.cc.Invoke(ctx, "/gon2n.EdgeManager/Create", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *edgeManagerClient) List(ctx context.Context, in *EdgeManagerListArgs, opts ...grpc.CallOption) (*EdgeManagerListReply, error) {
	out := new(EdgeManagerListReply)
	err := c.cc.Invoke(ctx, "/gon2n.EdgeManager/List", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *edgeManagerClient) Delete(ctx context.Context, in *EdgeManagerDeleteArgs, opts ...grpc.CallOption) (*EdgeManagerDeleteReply, error) {
	out := new(EdgeManagerDeleteReply)
	err := c.cc.Invoke(ctx, "/gon2n.EdgeManager/Delete", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// EdgeManagerServer is the server API for EdgeManager service.
type EdgeManagerServer interface {
	Create(context.Context, *EdgeManagerCreateArgs) (*EdgeManagerCreateReply, error)
	List(context.Context, *EdgeManagerListArgs) (*EdgeManagerListReply, error)
	Delete(context.Context, *EdgeManagerDeleteArgs) (*EdgeManagerDeleteReply, error)
}

// UnimplementedEdgeManagerServer can be embedded to have forward compatible implementations.
type UnimplementedEdgeManagerServer struct {
}

func (*UnimplementedEdgeManagerServer) Create(ctx context.Context, req *EdgeManagerCreateArgs) (*EdgeManagerCreateReply, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Create not implemented")
}
func (*UnimplementedEdgeManagerServer) List(ctx context.Context, req *EdgeManagerListArgs) (*EdgeManagerListReply, error) {
	return nil, status.Errorf(codes.Unimplemented, "method List not implemented")
}
func (*UnimplementedEdgeManagerServer) Delete(ctx context.Context, req *EdgeManagerDeleteArgs) (*EdgeManagerDeleteReply, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Delete not implemented")
}

func RegisterEdgeManagerServer(s *grpc.Server, srv EdgeManagerServer) {
	s.RegisterService(&_EdgeManager_serviceDesc, srv)
}

func _EdgeManager_Create_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(EdgeManagerCreateArgs)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(EdgeManagerServer).Create(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/gon2n.EdgeManager/Create",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(EdgeManagerServer).Create(ctx, req.(*EdgeManagerCreateArgs))
	}
	return interceptor(ctx, in, info, handler)
}

func _EdgeManager_List_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(EdgeManagerListArgs)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(EdgeManagerServer).List(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/gon2n.EdgeManager/List",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(EdgeManagerServer).List(ctx, req.(*EdgeManagerListArgs))
	}
	return interceptor(ctx, in, info, handler)
}

func _EdgeManager_Delete_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(EdgeManagerDeleteArgs)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(EdgeManagerServer).Delete(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/gon2n.EdgeManager/Delete",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(EdgeManagerServer).Delete(ctx, req.(*EdgeManagerDeleteArgs))
	}
	return interceptor(ctx, in, info, handler)
}

var _EdgeManager_serviceDesc = grpc.ServiceDesc{
	ServiceName: "gon2n.EdgeManager",
	HandlerType: (*EdgeManagerServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "Create",
			Handler:    _EdgeManager_Create_Handler,
		},
		{
			MethodName: "List",
			Handler:    _EdgeManager_List_Handler,
		},
		{
			MethodName: "Delete",
			Handler:    _EdgeManager_Delete_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "edge.proto",
}
