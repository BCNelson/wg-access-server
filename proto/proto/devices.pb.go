// Code generated by protoc-gen-go. DO NOT EDIT.
// source: devices.proto

package proto

import (
	context "context"
	fmt "fmt"
	proto "github.com/golang/protobuf/proto"
	empty "github.com/golang/protobuf/ptypes/empty"
	timestamp "github.com/golang/protobuf/ptypes/timestamp"
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

type Device struct {
	Name                 string               `protobuf:"bytes,1,opt,name=name,proto3" json:"name,omitempty"`
	Owner                string               `protobuf:"bytes,2,opt,name=owner,proto3" json:"owner,omitempty"`
	PublicKey            string               `protobuf:"bytes,3,opt,name=public_key,json=publicKey,proto3" json:"public_key,omitempty"`
	Address              string               `protobuf:"bytes,4,opt,name=address,proto3" json:"address,omitempty"`
	CreatedAt            *timestamp.Timestamp `protobuf:"bytes,5,opt,name=created_at,json=createdAt,proto3" json:"created_at,omitempty"`
	Connected            bool                 `protobuf:"varint,6,opt,name=connected,proto3" json:"connected,omitempty"`
	LastHandshakeTime    *timestamp.Timestamp `protobuf:"bytes,7,opt,name=last_handshake_time,json=lastHandshakeTime,proto3" json:"last_handshake_time,omitempty"`
	ReceiveBytes         int64                `protobuf:"varint,8,opt,name=receive_bytes,json=receiveBytes,proto3" json:"receive_bytes,omitempty"`
	TransmitBytes        int64                `protobuf:"varint,9,opt,name=transmit_bytes,json=transmitBytes,proto3" json:"transmit_bytes,omitempty"`
	Endpoint             string               `protobuf:"bytes,10,opt,name=endpoint,proto3" json:"endpoint,omitempty"`
	OwnerName            string               `protobuf:"bytes,11,opt,name=owner_name,json=ownerName,proto3" json:"owner_name,omitempty"`
	OwnerEmail           string               `protobuf:"bytes,12,opt,name=owner_email,json=ownerEmail,proto3" json:"owner_email,omitempty"`
	OwnerProvider        string               `protobuf:"bytes,13,opt,name=owner_provider,json=ownerProvider,proto3" json:"owner_provider,omitempty"`
	XXX_NoUnkeyedLiteral struct{}             `json:"-"`
	XXX_unrecognized     []byte               `json:"-"`
	XXX_sizecache        int32                `json:"-"`
}

func (m *Device) Reset()         { *m = Device{} }
func (m *Device) String() string { return proto.CompactTextString(m) }
func (*Device) ProtoMessage()    {}
func (*Device) Descriptor() ([]byte, []int) {
	return fileDescriptor_6d27ec3f2c0e2043, []int{0}
}

func (m *Device) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Device.Unmarshal(m, b)
}
func (m *Device) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Device.Marshal(b, m, deterministic)
}
func (m *Device) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Device.Merge(m, src)
}
func (m *Device) XXX_Size() int {
	return xxx_messageInfo_Device.Size(m)
}
func (m *Device) XXX_DiscardUnknown() {
	xxx_messageInfo_Device.DiscardUnknown(m)
}

var xxx_messageInfo_Device proto.InternalMessageInfo

func (m *Device) GetName() string {
	if m != nil {
		return m.Name
	}
	return ""
}

func (m *Device) GetOwner() string {
	if m != nil {
		return m.Owner
	}
	return ""
}

func (m *Device) GetPublicKey() string {
	if m != nil {
		return m.PublicKey
	}
	return ""
}

func (m *Device) GetAddress() string {
	if m != nil {
		return m.Address
	}
	return ""
}

func (m *Device) GetCreatedAt() *timestamp.Timestamp {
	if m != nil {
		return m.CreatedAt
	}
	return nil
}

func (m *Device) GetConnected() bool {
	if m != nil {
		return m.Connected
	}
	return false
}

func (m *Device) GetLastHandshakeTime() *timestamp.Timestamp {
	if m != nil {
		return m.LastHandshakeTime
	}
	return nil
}

func (m *Device) GetReceiveBytes() int64 {
	if m != nil {
		return m.ReceiveBytes
	}
	return 0
}

func (m *Device) GetTransmitBytes() int64 {
	if m != nil {
		return m.TransmitBytes
	}
	return 0
}

func (m *Device) GetEndpoint() string {
	if m != nil {
		return m.Endpoint
	}
	return ""
}

func (m *Device) GetOwnerName() string {
	if m != nil {
		return m.OwnerName
	}
	return ""
}

func (m *Device) GetOwnerEmail() string {
	if m != nil {
		return m.OwnerEmail
	}
	return ""
}

func (m *Device) GetOwnerProvider() string {
	if m != nil {
		return m.OwnerProvider
	}
	return ""
}

type AddDeviceReq struct {
	Name                 string   `protobuf:"bytes,1,opt,name=name,proto3" json:"name,omitempty"`
	PublicKey            string   `protobuf:"bytes,2,opt,name=public_key,json=publicKey,proto3" json:"public_key,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *AddDeviceReq) Reset()         { *m = AddDeviceReq{} }
func (m *AddDeviceReq) String() string { return proto.CompactTextString(m) }
func (*AddDeviceReq) ProtoMessage()    {}
func (*AddDeviceReq) Descriptor() ([]byte, []int) {
	return fileDescriptor_6d27ec3f2c0e2043, []int{1}
}

func (m *AddDeviceReq) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_AddDeviceReq.Unmarshal(m, b)
}
func (m *AddDeviceReq) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_AddDeviceReq.Marshal(b, m, deterministic)
}
func (m *AddDeviceReq) XXX_Merge(src proto.Message) {
	xxx_messageInfo_AddDeviceReq.Merge(m, src)
}
func (m *AddDeviceReq) XXX_Size() int {
	return xxx_messageInfo_AddDeviceReq.Size(m)
}
func (m *AddDeviceReq) XXX_DiscardUnknown() {
	xxx_messageInfo_AddDeviceReq.DiscardUnknown(m)
}

var xxx_messageInfo_AddDeviceReq proto.InternalMessageInfo

func (m *AddDeviceReq) GetName() string {
	if m != nil {
		return m.Name
	}
	return ""
}

func (m *AddDeviceReq) GetPublicKey() string {
	if m != nil {
		return m.PublicKey
	}
	return ""
}

type ListDevicesReq struct {
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *ListDevicesReq) Reset()         { *m = ListDevicesReq{} }
func (m *ListDevicesReq) String() string { return proto.CompactTextString(m) }
func (*ListDevicesReq) ProtoMessage()    {}
func (*ListDevicesReq) Descriptor() ([]byte, []int) {
	return fileDescriptor_6d27ec3f2c0e2043, []int{2}
}

func (m *ListDevicesReq) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_ListDevicesReq.Unmarshal(m, b)
}
func (m *ListDevicesReq) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_ListDevicesReq.Marshal(b, m, deterministic)
}
func (m *ListDevicesReq) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ListDevicesReq.Merge(m, src)
}
func (m *ListDevicesReq) XXX_Size() int {
	return xxx_messageInfo_ListDevicesReq.Size(m)
}
func (m *ListDevicesReq) XXX_DiscardUnknown() {
	xxx_messageInfo_ListDevicesReq.DiscardUnknown(m)
}

var xxx_messageInfo_ListDevicesReq proto.InternalMessageInfo

type ListDevicesRes struct {
	Items                []*Device `protobuf:"bytes,1,rep,name=items,proto3" json:"items,omitempty"`
	XXX_NoUnkeyedLiteral struct{}  `json:"-"`
	XXX_unrecognized     []byte    `json:"-"`
	XXX_sizecache        int32     `json:"-"`
}

func (m *ListDevicesRes) Reset()         { *m = ListDevicesRes{} }
func (m *ListDevicesRes) String() string { return proto.CompactTextString(m) }
func (*ListDevicesRes) ProtoMessage()    {}
func (*ListDevicesRes) Descriptor() ([]byte, []int) {
	return fileDescriptor_6d27ec3f2c0e2043, []int{3}
}

func (m *ListDevicesRes) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_ListDevicesRes.Unmarshal(m, b)
}
func (m *ListDevicesRes) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_ListDevicesRes.Marshal(b, m, deterministic)
}
func (m *ListDevicesRes) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ListDevicesRes.Merge(m, src)
}
func (m *ListDevicesRes) XXX_Size() int {
	return xxx_messageInfo_ListDevicesRes.Size(m)
}
func (m *ListDevicesRes) XXX_DiscardUnknown() {
	xxx_messageInfo_ListDevicesRes.DiscardUnknown(m)
}

var xxx_messageInfo_ListDevicesRes proto.InternalMessageInfo

func (m *ListDevicesRes) GetItems() []*Device {
	if m != nil {
		return m.Items
	}
	return nil
}

type DeleteDeviceReq struct {
	Name                 string   `protobuf:"bytes,1,opt,name=name,proto3" json:"name,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *DeleteDeviceReq) Reset()         { *m = DeleteDeviceReq{} }
func (m *DeleteDeviceReq) String() string { return proto.CompactTextString(m) }
func (*DeleteDeviceReq) ProtoMessage()    {}
func (*DeleteDeviceReq) Descriptor() ([]byte, []int) {
	return fileDescriptor_6d27ec3f2c0e2043, []int{4}
}

func (m *DeleteDeviceReq) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_DeleteDeviceReq.Unmarshal(m, b)
}
func (m *DeleteDeviceReq) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_DeleteDeviceReq.Marshal(b, m, deterministic)
}
func (m *DeleteDeviceReq) XXX_Merge(src proto.Message) {
	xxx_messageInfo_DeleteDeviceReq.Merge(m, src)
}
func (m *DeleteDeviceReq) XXX_Size() int {
	return xxx_messageInfo_DeleteDeviceReq.Size(m)
}
func (m *DeleteDeviceReq) XXX_DiscardUnknown() {
	xxx_messageInfo_DeleteDeviceReq.DiscardUnknown(m)
}

var xxx_messageInfo_DeleteDeviceReq proto.InternalMessageInfo

func (m *DeleteDeviceReq) GetName() string {
	if m != nil {
		return m.Name
	}
	return ""
}

type ListAllDevicesReq struct {
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *ListAllDevicesReq) Reset()         { *m = ListAllDevicesReq{} }
func (m *ListAllDevicesReq) String() string { return proto.CompactTextString(m) }
func (*ListAllDevicesReq) ProtoMessage()    {}
func (*ListAllDevicesReq) Descriptor() ([]byte, []int) {
	return fileDescriptor_6d27ec3f2c0e2043, []int{5}
}

func (m *ListAllDevicesReq) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_ListAllDevicesReq.Unmarshal(m, b)
}
func (m *ListAllDevicesReq) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_ListAllDevicesReq.Marshal(b, m, deterministic)
}
func (m *ListAllDevicesReq) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ListAllDevicesReq.Merge(m, src)
}
func (m *ListAllDevicesReq) XXX_Size() int {
	return xxx_messageInfo_ListAllDevicesReq.Size(m)
}
func (m *ListAllDevicesReq) XXX_DiscardUnknown() {
	xxx_messageInfo_ListAllDevicesReq.DiscardUnknown(m)
}

var xxx_messageInfo_ListAllDevicesReq proto.InternalMessageInfo

type ListAllDevicesRes struct {
	Items                []*Device `protobuf:"bytes,1,rep,name=items,proto3" json:"items,omitempty"`
	XXX_NoUnkeyedLiteral struct{}  `json:"-"`
	XXX_unrecognized     []byte    `json:"-"`
	XXX_sizecache        int32     `json:"-"`
}

func (m *ListAllDevicesRes) Reset()         { *m = ListAllDevicesRes{} }
func (m *ListAllDevicesRes) String() string { return proto.CompactTextString(m) }
func (*ListAllDevicesRes) ProtoMessage()    {}
func (*ListAllDevicesRes) Descriptor() ([]byte, []int) {
	return fileDescriptor_6d27ec3f2c0e2043, []int{6}
}

func (m *ListAllDevicesRes) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_ListAllDevicesRes.Unmarshal(m, b)
}
func (m *ListAllDevicesRes) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_ListAllDevicesRes.Marshal(b, m, deterministic)
}
func (m *ListAllDevicesRes) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ListAllDevicesRes.Merge(m, src)
}
func (m *ListAllDevicesRes) XXX_Size() int {
	return xxx_messageInfo_ListAllDevicesRes.Size(m)
}
func (m *ListAllDevicesRes) XXX_DiscardUnknown() {
	xxx_messageInfo_ListAllDevicesRes.DiscardUnknown(m)
}

var xxx_messageInfo_ListAllDevicesRes proto.InternalMessageInfo

func (m *ListAllDevicesRes) GetItems() []*Device {
	if m != nil {
		return m.Items
	}
	return nil
}

func init() {
	proto.RegisterType((*Device)(nil), "proto.Device")
	proto.RegisterType((*AddDeviceReq)(nil), "proto.AddDeviceReq")
	proto.RegisterType((*ListDevicesReq)(nil), "proto.ListDevicesReq")
	proto.RegisterType((*ListDevicesRes)(nil), "proto.ListDevicesRes")
	proto.RegisterType((*DeleteDeviceReq)(nil), "proto.DeleteDeviceReq")
	proto.RegisterType((*ListAllDevicesReq)(nil), "proto.ListAllDevicesReq")
	proto.RegisterType((*ListAllDevicesRes)(nil), "proto.ListAllDevicesRes")
}

func init() { proto.RegisterFile("devices.proto", fileDescriptor_6d27ec3f2c0e2043) }

var fileDescriptor_6d27ec3f2c0e2043 = []byte{
	// 513 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x8c, 0x53, 0xdb, 0x6e, 0xd3, 0x40,
	0x10, 0x8d, 0x93, 0xe6, 0xe2, 0x49, 0x5c, 0xe8, 0x04, 0xaa, 0x95, 0x01, 0x35, 0x72, 0x55, 0x29,
	0x4f, 0xa9, 0x08, 0x42, 0x82, 0x07, 0x24, 0x82, 0x5a, 0x84, 0x00, 0x21, 0x64, 0xf1, 0x6e, 0x39,
	0xde, 0xa1, 0x5d, 0xd5, 0xb7, 0x7a, 0xb7, 0x41, 0xf9, 0x00, 0xbe, 0x93, 0x5f, 0x41, 0xbb, 0xeb,
	0xa4, 0x49, 0x5a, 0x2e, 0x4f, 0xf6, 0x9c, 0x73, 0xe6, 0x76, 0xc6, 0x06, 0x8f, 0xd3, 0x42, 0x24,
	0x24, 0x27, 0x65, 0x55, 0xa8, 0x02, 0xdb, 0xe6, 0xe1, 0x1f, 0x5d, 0x14, 0xc5, 0x45, 0x4a, 0xa7,
	0x26, 0x9a, 0xdf, 0x7c, 0x3f, 0x55, 0x22, 0x23, 0xa9, 0xe2, 0xac, 0xb4, 0x3a, 0xff, 0xc9, 0xae,
	0x80, 0xb2, 0x52, 0x2d, 0x2d, 0x19, 0xfc, 0x6a, 0x41, 0xe7, 0xcc, 0x94, 0x45, 0x84, 0xbd, 0x3c,
	0xce, 0x88, 0x39, 0x23, 0x67, 0xec, 0x86, 0xe6, 0x1d, 0x1f, 0x41, 0xbb, 0xf8, 0x91, 0x53, 0xc5,
	0x9a, 0x06, 0xb4, 0x01, 0x3e, 0x03, 0x28, 0x6f, 0xe6, 0xa9, 0x48, 0xa2, 0x2b, 0x5a, 0xb2, 0x96,
	0xa1, 0x5c, 0x8b, 0x7c, 0xa2, 0x25, 0x32, 0xe8, 0xc6, 0x9c, 0x57, 0x24, 0x25, 0xdb, 0x33, 0xdc,
	0x2a, 0xc4, 0xd7, 0x00, 0x49, 0x45, 0xb1, 0x22, 0x1e, 0xc5, 0x8a, 0xb5, 0x47, 0xce, 0xb8, 0x3f,
	0xf5, 0x27, 0x76, 0xbe, 0xc9, 0x6a, 0xbe, 0xc9, 0xb7, 0xd5, 0x02, 0xa1, 0x5b, 0xab, 0x67, 0x0a,
	0x9f, 0x82, 0x9b, 0x14, 0x79, 0x4e, 0x89, 0x22, 0xce, 0x3a, 0x23, 0x67, 0xdc, 0x0b, 0x6f, 0x01,
	0xfc, 0x08, 0xc3, 0x34, 0x96, 0x2a, 0xba, 0x8c, 0x73, 0x2e, 0x2f, 0xe3, 0x2b, 0x8a, 0xb4, 0x0b,
	0xac, 0xfb, 0xcf, 0x0e, 0x07, 0x3a, 0xed, 0xc3, 0x2a, 0x4b, 0xe3, 0x78, 0x0c, 0x5e, 0x45, 0x09,
	0x89, 0x05, 0x45, 0xf3, 0xa5, 0x22, 0xc9, 0x7a, 0x23, 0x67, 0xdc, 0x0a, 0x07, 0x35, 0xf8, 0x4e,
	0x63, 0x78, 0x02, 0xfb, 0xaa, 0x8a, 0x73, 0x99, 0x09, 0x55, 0xab, 0x5c, 0xa3, 0xf2, 0x56, 0xa8,
	0x95, 0xf9, 0xd0, 0xa3, 0x9c, 0x97, 0x85, 0xc8, 0x15, 0x03, 0xe3, 0xc5, 0x3a, 0xd6, 0x2e, 0x1a,
	0x3b, 0x23, 0xe3, 0x7a, 0xdf, 0xba, 0x68, 0x90, 0x2f, 0xda, 0xfa, 0x23, 0xe8, 0x5b, 0x9a, 0xb2,
	0x58, 0xa4, 0x6c, 0x60, 0x78, 0x9b, 0x71, 0xae, 0x11, 0x3d, 0x82, 0x15, 0x94, 0x55, 0xb1, 0x10,
	0x9c, 0x2a, 0xe6, 0x19, 0x8d, 0x67, 0xd0, 0xaf, 0x35, 0x18, 0xcc, 0x60, 0x30, 0xe3, 0xdc, 0xde,
	0x38, 0xa4, 0xeb, 0x7b, 0xcf, 0xbc, 0x7d, 0xd0, 0xe6, 0xce, 0x41, 0x83, 0x87, 0xb0, 0xff, 0x59,
	0x48, 0x65, 0x6b, 0xc8, 0x90, 0xae, 0x83, 0x97, 0x3b, 0x88, 0xc4, 0x63, 0x68, 0x0b, 0x45, 0x99,
	0x64, 0xce, 0xa8, 0x35, 0xee, 0x4f, 0x3d, 0x6b, 0xf6, 0xa4, 0xee, 0x6b, 0xb9, 0xe0, 0x04, 0x1e,
	0x9c, 0x51, 0x4a, 0x8a, 0xfe, 0x3a, 0x4e, 0x30, 0x84, 0x03, 0x5d, 0x7d, 0x96, 0xa6, 0x1b, 0x2d,
	0x5f, 0xdd, 0x05, 0xff, 0xaf, 0xeb, 0xf4, 0x67, 0x13, 0xba, 0x75, 0x0e, 0x3e, 0x07, 0x77, 0xed,
	0x06, 0x0e, 0x6b, 0xf9, 0xa6, 0x3f, 0xfe, 0x76, 0x8d, 0xa0, 0x81, 0x6f, 0xa0, 0xbf, 0xb1, 0x2b,
	0x3e, 0xae, 0xf9, 0x6d, 0x47, 0xfc, 0x7b, 0x61, 0x19, 0x34, 0xf0, 0x2d, 0x0c, 0x36, 0x77, 0xc6,
	0xc3, 0x75, 0xfd, 0x2d, 0x23, 0xfc, 0xc3, 0x3b, 0x5f, 0xe9, 0xb9, 0xfe, 0x4f, 0x83, 0x06, 0xbe,
	0xb7, 0x66, 0xdf, 0x6e, 0x8e, 0x6c, 0xa3, 0xd9, 0x96, 0x4b, 0xfe, 0x9f, 0x18, 0x19, 0x34, 0xe6,
	0x1d, 0x43, 0xbd, 0xf8, 0x1d, 0x00, 0x00, 0xff, 0xff, 0x89, 0x1c, 0x7b, 0x9d, 0x48, 0x04, 0x00,
	0x00,
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConnInterface

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion6

// DevicesClient is the client API for Devices service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type DevicesClient interface {
	AddDevice(ctx context.Context, in *AddDeviceReq, opts ...grpc.CallOption) (*Device, error)
	ListDevices(ctx context.Context, in *ListDevicesReq, opts ...grpc.CallOption) (*ListDevicesRes, error)
	DeleteDevice(ctx context.Context, in *DeleteDeviceReq, opts ...grpc.CallOption) (*empty.Empty, error)
	// admin only
	ListAllDevices(ctx context.Context, in *ListAllDevicesReq, opts ...grpc.CallOption) (*ListAllDevicesRes, error)
}

type devicesClient struct {
	cc grpc.ClientConnInterface
}

func NewDevicesClient(cc grpc.ClientConnInterface) DevicesClient {
	return &devicesClient{cc}
}

func (c *devicesClient) AddDevice(ctx context.Context, in *AddDeviceReq, opts ...grpc.CallOption) (*Device, error) {
	out := new(Device)
	err := c.cc.Invoke(ctx, "/proto.Devices/AddDevice", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *devicesClient) ListDevices(ctx context.Context, in *ListDevicesReq, opts ...grpc.CallOption) (*ListDevicesRes, error) {
	out := new(ListDevicesRes)
	err := c.cc.Invoke(ctx, "/proto.Devices/ListDevices", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *devicesClient) DeleteDevice(ctx context.Context, in *DeleteDeviceReq, opts ...grpc.CallOption) (*empty.Empty, error) {
	out := new(empty.Empty)
	err := c.cc.Invoke(ctx, "/proto.Devices/DeleteDevice", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *devicesClient) ListAllDevices(ctx context.Context, in *ListAllDevicesReq, opts ...grpc.CallOption) (*ListAllDevicesRes, error) {
	out := new(ListAllDevicesRes)
	err := c.cc.Invoke(ctx, "/proto.Devices/ListAllDevices", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// DevicesServer is the server API for Devices service.
type DevicesServer interface {
	AddDevice(context.Context, *AddDeviceReq) (*Device, error)
	ListDevices(context.Context, *ListDevicesReq) (*ListDevicesRes, error)
	DeleteDevice(context.Context, *DeleteDeviceReq) (*empty.Empty, error)
	// admin only
	ListAllDevices(context.Context, *ListAllDevicesReq) (*ListAllDevicesRes, error)
}

// UnimplementedDevicesServer can be embedded to have forward compatible implementations.
type UnimplementedDevicesServer struct {
}

func (*UnimplementedDevicesServer) AddDevice(ctx context.Context, req *AddDeviceReq) (*Device, error) {
	return nil, status.Errorf(codes.Unimplemented, "method AddDevice not implemented")
}
func (*UnimplementedDevicesServer) ListDevices(ctx context.Context, req *ListDevicesReq) (*ListDevicesRes, error) {
	return nil, status.Errorf(codes.Unimplemented, "method ListDevices not implemented")
}
func (*UnimplementedDevicesServer) DeleteDevice(ctx context.Context, req *DeleteDeviceReq) (*empty.Empty, error) {
	return nil, status.Errorf(codes.Unimplemented, "method DeleteDevice not implemented")
}
func (*UnimplementedDevicesServer) ListAllDevices(ctx context.Context, req *ListAllDevicesReq) (*ListAllDevicesRes, error) {
	return nil, status.Errorf(codes.Unimplemented, "method ListAllDevices not implemented")
}

func RegisterDevicesServer(s *grpc.Server, srv DevicesServer) {
	s.RegisterService(&_Devices_serviceDesc, srv)
}

func _Devices_AddDevice_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(AddDeviceReq)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(DevicesServer).AddDevice(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/proto.Devices/AddDevice",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(DevicesServer).AddDevice(ctx, req.(*AddDeviceReq))
	}
	return interceptor(ctx, in, info, handler)
}

func _Devices_ListDevices_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ListDevicesReq)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(DevicesServer).ListDevices(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/proto.Devices/ListDevices",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(DevicesServer).ListDevices(ctx, req.(*ListDevicesReq))
	}
	return interceptor(ctx, in, info, handler)
}

func _Devices_DeleteDevice_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(DeleteDeviceReq)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(DevicesServer).DeleteDevice(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/proto.Devices/DeleteDevice",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(DevicesServer).DeleteDevice(ctx, req.(*DeleteDeviceReq))
	}
	return interceptor(ctx, in, info, handler)
}

func _Devices_ListAllDevices_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ListAllDevicesReq)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(DevicesServer).ListAllDevices(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/proto.Devices/ListAllDevices",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(DevicesServer).ListAllDevices(ctx, req.(*ListAllDevicesReq))
	}
	return interceptor(ctx, in, info, handler)
}

var _Devices_serviceDesc = grpc.ServiceDesc{
	ServiceName: "proto.Devices",
	HandlerType: (*DevicesServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "AddDevice",
			Handler:    _Devices_AddDevice_Handler,
		},
		{
			MethodName: "ListDevices",
			Handler:    _Devices_ListDevices_Handler,
		},
		{
			MethodName: "DeleteDevice",
			Handler:    _Devices_DeleteDevice_Handler,
		},
		{
			MethodName: "ListAllDevices",
			Handler:    _Devices_ListAllDevices_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "devices.proto",
}
