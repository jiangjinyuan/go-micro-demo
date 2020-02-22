// Code generated by protoc-gen-go. DO NOT EDIT.
// source: proto/poolHashrate/poolHashrate.proto

package btccom_explorer_srv_poolHashrate

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

type PoolHashrate struct {
	CreatedTime          string   `protobuf:"bytes,1,opt,name=createdTime,proto3" json:"createdTime,omitempty"`
	PoolName             string   `protobuf:"bytes,2,opt,name=pool_name,json=poolName,proto3" json:"pool_name,omitempty"`
	PoolId               int32    `protobuf:"varint,3,opt,name=pool_id,json=poolId,proto3" json:"pool_id,omitempty"`
	Hashrate             float32  `protobuf:"fixed32,4,opt,name=hashrate,proto3" json:"hashrate,omitempty"`
	UpdatedTime          string   `protobuf:"bytes,5,opt,name=updatedTime,proto3" json:"updatedTime,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *PoolHashrate) Reset()         { *m = PoolHashrate{} }
func (m *PoolHashrate) String() string { return proto.CompactTextString(m) }
func (*PoolHashrate) ProtoMessage()    {}
func (*PoolHashrate) Descriptor() ([]byte, []int) {
	return fileDescriptor_55f9957aac84fd73, []int{0}
}

func (m *PoolHashrate) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_PoolHashrate.Unmarshal(m, b)
}
func (m *PoolHashrate) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_PoolHashrate.Marshal(b, m, deterministic)
}
func (m *PoolHashrate) XXX_Merge(src proto.Message) {
	xxx_messageInfo_PoolHashrate.Merge(m, src)
}
func (m *PoolHashrate) XXX_Size() int {
	return xxx_messageInfo_PoolHashrate.Size(m)
}
func (m *PoolHashrate) XXX_DiscardUnknown() {
	xxx_messageInfo_PoolHashrate.DiscardUnknown(m)
}

var xxx_messageInfo_PoolHashrate proto.InternalMessageInfo

func (m *PoolHashrate) GetCreatedTime() string {
	if m != nil {
		return m.CreatedTime
	}
	return ""
}

func (m *PoolHashrate) GetPoolName() string {
	if m != nil {
		return m.PoolName
	}
	return ""
}

func (m *PoolHashrate) GetPoolId() int32 {
	if m != nil {
		return m.PoolId
	}
	return 0
}

func (m *PoolHashrate) GetHashrate() float32 {
	if m != nil {
		return m.Hashrate
	}
	return 0
}

func (m *PoolHashrate) GetUpdatedTime() string {
	if m != nil {
		return m.UpdatedTime
	}
	return ""
}

type Request struct {
	PoolID               string    `protobuf:"bytes,1,opt,name=poolID,proto3" json:"poolID,omitempty"`
	//PoolName             string    `protobuf:"bytes,1,opt,name=poolName,proto3" json:"poolName,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *Request) Reset()         { *m = Request{} }
func (m *Request) String() string { return proto.CompactTextString(m) }
func (*Request) ProtoMessage()    {}
func (*Request) Descriptor() ([]byte, []int) {
	return fileDescriptor_55f9957aac84fd73, []int{1}
}

func (m *Request) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Request.Unmarshal(m, b)
}
func (m *Request) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Request.Marshal(b, m, deterministic)
}
func (m *Request) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Request.Merge(m, src)
}
func (m *Request) XXX_Size() int {
	return xxx_messageInfo_Request.Size(m)
}
func (m *Request) XXX_DiscardUnknown() {
	xxx_messageInfo_Request.DiscardUnknown(m)
}

var xxx_messageInfo_Request proto.InternalMessageInfo

func (m *Request) GetPoolID() string {
	if m != nil {
		return m.PoolID
	}
	return ""
}

/*func (m *Request) GetPoolName() string {
	if m != nil {
		return m.PoolName
	}
	return ""
}*/

type Response struct {
	Success              bool          `protobuf:"varint,1,opt,name=success,proto3" json:"success,omitempty"`
	Error                *Error        `protobuf:"bytes,2,opt,name=error,proto3" json:"error,omitempty"`
	PoolHashrate         *PoolHashrate `protobuf:"bytes,3,opt,name=poolHashrate,proto3" json:"poolHashrate,omitempty"`
	XXX_NoUnkeyedLiteral struct{}      `json:"-"`
	XXX_unrecognized     []byte        `json:"-"`
	XXX_sizecache        int32         `json:"-"`
}

func (m *Response) Reset()         { *m = Response{} }
func (m *Response) String() string { return proto.CompactTextString(m) }
func (*Response) ProtoMessage()    {}
func (*Response) Descriptor() ([]byte, []int) {
	return fileDescriptor_55f9957aac84fd73, []int{2}
}

func (m *Response) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Response.Unmarshal(m, b)
}
func (m *Response) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Response.Marshal(b, m, deterministic)
}
func (m *Response) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Response.Merge(m, src)
}
func (m *Response) XXX_Size() int {
	return xxx_messageInfo_Response.Size(m)
}
func (m *Response) XXX_DiscardUnknown() {
	xxx_messageInfo_Response.DiscardUnknown(m)
}

var xxx_messageInfo_Response proto.InternalMessageInfo

func (m *Response) GetSuccess() bool {
	if m != nil {
		return m.Success
	}
	return false
}

func (m *Response) GetError() *Error {
	if m != nil {
		return m.Error
	}
	return nil
}

func (m *Response) GetPoolHashrate() *PoolHashrate {
	if m != nil {
		return m.PoolHashrate
	}
	return nil
}

type Error struct {
	Code                 int32    `protobuf:"varint,1,opt,name=code,proto3" json:"code,omitempty"`
	Detail               string   `protobuf:"bytes,2,opt,name=detail,proto3" json:"detail,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *Error) Reset()         { *m = Error{} }
func (m *Error) String() string { return proto.CompactTextString(m) }
func (*Error) ProtoMessage()    {}
func (*Error) Descriptor() ([]byte, []int) {
	return fileDescriptor_55f9957aac84fd73, []int{3}
}

func (m *Error) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Error.Unmarshal(m, b)
}
func (m *Error) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Error.Marshal(b, m, deterministic)
}
func (m *Error) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Error.Merge(m, src)
}
func (m *Error) XXX_Size() int {
	return xxx_messageInfo_Error.Size(m)
}
func (m *Error) XXX_DiscardUnknown() {
	xxx_messageInfo_Error.DiscardUnknown(m)
}

var xxx_messageInfo_Error proto.InternalMessageInfo

func (m *Error) GetCode() int32 {
	if m != nil {
		return m.Code
	}
	return 0
}

func (m *Error) GetDetail() string {
	if m != nil {
		return m.Detail
	}
	return ""
}

func init() {
	proto.RegisterType((*PoolHashrate)(nil), "btccom.explorer.srv.poolHashrate.poolHashrate")
	proto.RegisterType((*Request)(nil), "btccom.explorer.srv.poolHashrate.Request")
	proto.RegisterType((*Response)(nil), "btccom.explorer.srv.poolHashrate.Response")
	proto.RegisterType((*Error)(nil), "btccom.explorer.srv.poolHashrate.Error")
}

func init() {
	proto.RegisterFile("proto/poolHashrate/poolHashrate.proto", fileDescriptor_55f9957aac84fd73)
}

var fileDescriptor_55f9957aac84fd73 = []byte{
	// 322 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x8c, 0x52, 0x4d, 0x4b, 0x03, 0x31,
	0x14, 0x74, 0x6b, 0xb7, 0xdd, 0xbe, 0x16, 0x84, 0x1c, 0x74, 0xa9, 0x97, 0x65, 0x41, 0xac, 0x1e,
	0x22, 0xb4, 0xe7, 0xde, 0x14, 0xf5, 0x22, 0x12, 0xbc, 0x4b, 0x9a, 0x3c, 0x68, 0xa5, 0xdb, 0xac,
	0x49, 0x2a, 0xe2, 0xef, 0xf1, 0x3f, 0xf8, 0xf7, 0x24, 0xd9, 0xb4, 0x64, 0xbd, 0xd4, 0xdb, 0xce,
	0xbc, 0x8f, 0x99, 0x9d, 0x17, 0xb8, 0xa8, 0xb5, 0xb2, 0xea, 0xa6, 0x56, 0x6a, 0xfd, 0xc0, 0xcd,
	0x52, 0x73, 0x8b, 0x2d, 0x40, 0x7d, 0x9d, 0x14, 0x0b, 0x2b, 0x84, 0xaa, 0x28, 0x7e, 0xd6, 0x6b,
	0xa5, 0x51, 0x53, 0xa3, 0x3f, 0x68, 0xdc, 0x57, 0x7e, 0x27, 0x30, 0x8a, 0x09, 0x52, 0xc0, 0x50,
	0x68, 0xe4, 0x16, 0xe5, 0xcb, 0xaa, 0xc2, 0x3c, 0x29, 0x92, 0xc9, 0x80, 0xc5, 0x14, 0x39, 0x87,
	0x81, 0x9b, 0x78, 0xdd, 0xf0, 0x0a, 0xf3, 0x8e, 0xaf, 0x67, 0x8e, 0x78, 0xe2, 0x15, 0x92, 0x33,
	0xe8, 0xfb, 0xe2, 0x4a, 0xe6, 0xc7, 0x45, 0x32, 0x49, 0x59, 0xcf, 0xc1, 0x47, 0x49, 0xc6, 0x90,
	0x2d, 0x83, 0x46, 0xde, 0x2d, 0x92, 0x49, 0x87, 0xed, 0xb1, 0xd3, 0xdc, 0xd6, 0x72, 0xaf, 0x99,
	0x36, 0x9a, 0x11, 0x55, 0xce, 0xa1, 0xcf, 0xf0, 0x7d, 0x8b, 0xc6, 0x92, 0x53, 0x68, 0x56, 0xde,
	0x06, 0x6f, 0x01, 0x39, 0x81, 0x9d, 0x8b, 0xbf, 0xae, 0xca, 0x9f, 0x04, 0x32, 0x86, 0xa6, 0x56,
	0x1b, 0x83, 0x24, 0x87, 0xbe, 0xd9, 0x0a, 0x81, 0xc6, 0xf8, 0x0d, 0x19, 0xdb, 0x41, 0x32, 0x87,
	0x14, 0xb5, 0x56, 0xda, 0xcf, 0x0f, 0xa7, 0x97, 0xf4, 0x50, 0x7c, 0xf4, 0xce, 0xb5, 0xb3, 0x66,
	0x8a, 0xb0, 0x76, 0x94, 0x3e, 0x80, 0xe1, 0x94, 0x1e, 0xde, 0x12, 0x03, 0xd6, 0xda, 0x51, 0xce,
	0x20, 0xf5, 0x1a, 0x84, 0x40, 0x57, 0x28, 0xd9, 0x1c, 0x24, 0x65, 0xfe, 0xdb, 0x45, 0x21, 0xd1,
	0xf2, 0xd5, 0x3a, 0xfc, 0x70, 0x40, 0xd3, 0x2f, 0x18, 0x3d, 0xc7, 0x37, 0x7d, 0x83, 0x93, 0x7b,
	0xb4, 0x2d, 0xea, 0xea, 0xb0, 0xab, 0x10, 0xf8, 0xf8, 0xfa, 0x3f, 0xad, 0x4d, 0xb6, 0xe5, 0xd1,
	0xa2, 0xe7, 0x5f, 0xde, 0xec, 0x37, 0x00, 0x00, 0xff, 0xff, 0x49, 0x0e, 0xf2, 0xe8, 0xa2, 0x02,
	0x00, 0x00,
}
