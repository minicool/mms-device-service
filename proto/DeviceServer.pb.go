// Code generated by protoc-gen-go. DO NOT EDIT.
// source: DeviceServer.proto

/*
Package RFIDDevice is a generated protocol buffer package.

c++

It is generated from these files:
	DeviceServer.proto

It has these top-level messages:
	ErrorMessage
	DeviceRegeditRequest
	DeviceRegeditResponse
	RFIDprint_Data
	RFIDprintWriteDataStream_Request
	RFIDprintWriteDataStream_Response
*/
package RFIDDevice

import proto "github.com/golang/protobuf/proto"
import fmt "fmt"
import math "math"

import (
	context "golang.org/x/net/context"
	grpc "google.golang.org/grpc"
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

// Error Message
// 错误消息返回
type ERRORCODE int32

const (
	ERRORCODE_ERRORCODE_NONE             ERRORCODE = 0
	ERRORCODE_EXCUTION_ERROR             ERRORCODE = 1001
	ERRORCODE_ACCESS_FAILURE             ERRORCODE = 1002
	ERRORCODE_SYSTEM_BUSY                ERRORCODE = 1003
	ERRORCODE_FAILED_CONNECT_RFIDCARD    ERRORCODE = 1004
	ERRORCODE_FAILED_CONNECT_RFIDCARDER  ERRORCODE = 1005
	ERRORCODE_FAILED_CONNECT_RFIDPRINTER ERRORCODE = 1006
)

var ERRORCODE_name = map[int32]string{
	0:    "ERRORCODE_NONE",
	1001: "EXCUTION_ERROR",
	1002: "ACCESS_FAILURE",
	1003: "SYSTEM_BUSY",
	1004: "FAILED_CONNECT_RFIDCARD",
	1005: "FAILED_CONNECT_RFIDCARDER",
	1006: "FAILED_CONNECT_RFIDPRINTER",
}
var ERRORCODE_value = map[string]int32{
	"ERRORCODE_NONE":             0,
	"EXCUTION_ERROR":             1001,
	"ACCESS_FAILURE":             1002,
	"SYSTEM_BUSY":                1003,
	"FAILED_CONNECT_RFIDCARD":    1004,
	"FAILED_CONNECT_RFIDCARDER":  1005,
	"FAILED_CONNECT_RFIDPRINTER": 1006,
}

func (x ERRORCODE) String() string {
	return proto.EnumName(ERRORCODE_name, int32(x))
}
func (ERRORCODE) EnumDescriptor() ([]byte, []int) { return fileDescriptor0, []int{0} }

// DeviceRegedit
// 硬件系统注册
type DEVICETYPE int32

const (
	DEVICETYPE_TYPE_NONE DEVICETYPE = 0
	DEVICETYPE_PRINT     DEVICETYPE = 1
	DEVICETYPE_CARD      DEVICETYPE = 2
	DEVICETYPE_READ      DEVICETYPE = 3
)

var DEVICETYPE_name = map[int32]string{
	0: "TYPE_NONE",
	1: "PRINT",
	2: "CARD",
	3: "READ",
}
var DEVICETYPE_value = map[string]int32{
	"TYPE_NONE": 0,
	"PRINT":     1,
	"CARD":      2,
	"READ":      3,
}

func (x DEVICETYPE) String() string {
	return proto.EnumName(DEVICETYPE_name, int32(x))
}
func (DEVICETYPE) EnumDescriptor() ([]byte, []int) { return fileDescriptor0, []int{1} }

type ErrorMessage struct {
	ErrorCode ERRORCODE `protobuf:"varint,1,opt,name=errorCode,enum=RFIDDevice.ERRORCODE" json:"errorCode,omitempty"`
	ErrorMsg  string    `protobuf:"bytes,2,opt,name=errorMsg" json:"errorMsg,omitempty"`
}

func (m *ErrorMessage) Reset()                    { *m = ErrorMessage{} }
func (m *ErrorMessage) String() string            { return proto.CompactTextString(m) }
func (*ErrorMessage) ProtoMessage()               {}
func (*ErrorMessage) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{0} }

func (m *ErrorMessage) GetErrorCode() ERRORCODE {
	if m != nil {
		return m.ErrorCode
	}
	return ERRORCODE_ERRORCODE_NONE
}

func (m *ErrorMessage) GetErrorMsg() string {
	if m != nil {
		return m.ErrorMsg
	}
	return ""
}

type DeviceRegeditRequest struct {
	DeviceName string     `protobuf:"bytes,1,opt,name=device_name,json=deviceName" json:"device_name,omitempty"`
	DeviceType DEVICETYPE `protobuf:"varint,2,opt,name=device_type,json=deviceType,enum=RFIDDevice.DEVICETYPE" json:"device_type,omitempty"`
	IpAdd      string     `protobuf:"bytes,3,opt,name=ip_add,json=ipAdd" json:"ip_add,omitempty"`
	IpPort     int32      `protobuf:"varint,4,opt,name=ip_port,json=ipPort" json:"ip_port,omitempty"`
	MacAdd     string     `protobuf:"bytes,5,opt,name=mac_add,json=macAdd" json:"mac_add,omitempty"`
}

func (m *DeviceRegeditRequest) Reset()                    { *m = DeviceRegeditRequest{} }
func (m *DeviceRegeditRequest) String() string            { return proto.CompactTextString(m) }
func (*DeviceRegeditRequest) ProtoMessage()               {}
func (*DeviceRegeditRequest) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{1} }

func (m *DeviceRegeditRequest) GetDeviceName() string {
	if m != nil {
		return m.DeviceName
	}
	return ""
}

func (m *DeviceRegeditRequest) GetDeviceType() DEVICETYPE {
	if m != nil {
		return m.DeviceType
	}
	return DEVICETYPE_TYPE_NONE
}

func (m *DeviceRegeditRequest) GetIpAdd() string {
	if m != nil {
		return m.IpAdd
	}
	return ""
}

func (m *DeviceRegeditRequest) GetIpPort() int32 {
	if m != nil {
		return m.IpPort
	}
	return 0
}

func (m *DeviceRegeditRequest) GetMacAdd() string {
	if m != nil {
		return m.MacAdd
	}
	return ""
}

type DeviceRegeditResponse struct {
	Success       bool          `protobuf:"varint,1,opt,name=success" json:"success,omitempty"`
	DeviceRegedid uint64        `protobuf:"varint,2,opt,name=device_regedid,json=deviceRegedid" json:"device_regedid,omitempty"`
	Errormsg      *ErrorMessage `protobuf:"bytes,3,opt,name=errormsg" json:"errormsg,omitempty"`
}

func (m *DeviceRegeditResponse) Reset()                    { *m = DeviceRegeditResponse{} }
func (m *DeviceRegeditResponse) String() string            { return proto.CompactTextString(m) }
func (*DeviceRegeditResponse) ProtoMessage()               {}
func (*DeviceRegeditResponse) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{2} }

func (m *DeviceRegeditResponse) GetSuccess() bool {
	if m != nil {
		return m.Success
	}
	return false
}

func (m *DeviceRegeditResponse) GetDeviceRegedid() uint64 {
	if m != nil {
		return m.DeviceRegedid
	}
	return 0
}

func (m *DeviceRegeditResponse) GetErrormsg() *ErrorMessage {
	if m != nil {
		return m.Errormsg
	}
	return nil
}

// RFIDprintStream
// 物资管理系统向RFID打印机发送数据
type RFIDprint_Data struct {
	AssetsId      string `protobuf:"bytes,1,opt,name=assetsId" json:"assetsId,omitempty"`
	CompanyName   string `protobuf:"bytes,2,opt,name=companyName" json:"companyName,omitempty"`
	DevpementNmae string `protobuf:"bytes,3,opt,name=devpementNmae" json:"devpementNmae,omitempty"`
	AssetsName    string `protobuf:"bytes,4,opt,name=assetsName" json:"assetsName,omitempty"`
	AssetsType    string `protobuf:"bytes,5,opt,name=assetsType" json:"assetsType,omitempty"`
	AssetsModel   string `protobuf:"bytes,6,opt,name=assetsModel" json:"assetsModel,omitempty"`
	Url           string `protobuf:"bytes,7,opt,name=url" json:"url,omitempty"`
	Date          string `protobuf:"bytes,8,opt,name=date" json:"date,omitempty"`
}

func (m *RFIDprint_Data) Reset()                    { *m = RFIDprint_Data{} }
func (m *RFIDprint_Data) String() string            { return proto.CompactTextString(m) }
func (*RFIDprint_Data) ProtoMessage()               {}
func (*RFIDprint_Data) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{3} }

func (m *RFIDprint_Data) GetAssetsId() string {
	if m != nil {
		return m.AssetsId
	}
	return ""
}

func (m *RFIDprint_Data) GetCompanyName() string {
	if m != nil {
		return m.CompanyName
	}
	return ""
}

func (m *RFIDprint_Data) GetDevpementNmae() string {
	if m != nil {
		return m.DevpementNmae
	}
	return ""
}

func (m *RFIDprint_Data) GetAssetsName() string {
	if m != nil {
		return m.AssetsName
	}
	return ""
}

func (m *RFIDprint_Data) GetAssetsType() string {
	if m != nil {
		return m.AssetsType
	}
	return ""
}

func (m *RFIDprint_Data) GetAssetsModel() string {
	if m != nil {
		return m.AssetsModel
	}
	return ""
}

func (m *RFIDprint_Data) GetUrl() string {
	if m != nil {
		return m.Url
	}
	return ""
}

func (m *RFIDprint_Data) GetDate() string {
	if m != nil {
		return m.Date
	}
	return ""
}

type RFIDprintWriteDataStream_Request struct {
	DeviceRegedid int64 `protobuf:"varint,1,opt,name=device_regedid,json=deviceRegedid" json:"device_regedid,omitempty"`
}

func (m *RFIDprintWriteDataStream_Request) Reset()         { *m = RFIDprintWriteDataStream_Request{} }
func (m *RFIDprintWriteDataStream_Request) String() string { return proto.CompactTextString(m) }
func (*RFIDprintWriteDataStream_Request) ProtoMessage()    {}
func (*RFIDprintWriteDataStream_Request) Descriptor() ([]byte, []int) {
	return fileDescriptor0, []int{4}
}

func (m *RFIDprintWriteDataStream_Request) GetDeviceRegedid() int64 {
	if m != nil {
		return m.DeviceRegedid
	}
	return 0
}

type RFIDprintWriteDataStream_Response struct {
	Success  bool            `protobuf:"varint,1,opt,name=success" json:"success,omitempty"`
	Data     *RFIDprint_Data `protobuf:"bytes,2,opt,name=data" json:"data,omitempty"`
	Errormsg *ErrorMessage   `protobuf:"bytes,3,opt,name=errormsg" json:"errormsg,omitempty"`
}

func (m *RFIDprintWriteDataStream_Response) Reset()         { *m = RFIDprintWriteDataStream_Response{} }
func (m *RFIDprintWriteDataStream_Response) String() string { return proto.CompactTextString(m) }
func (*RFIDprintWriteDataStream_Response) ProtoMessage()    {}
func (*RFIDprintWriteDataStream_Response) Descriptor() ([]byte, []int) {
	return fileDescriptor0, []int{5}
}

func (m *RFIDprintWriteDataStream_Response) GetSuccess() bool {
	if m != nil {
		return m.Success
	}
	return false
}

func (m *RFIDprintWriteDataStream_Response) GetData() *RFIDprint_Data {
	if m != nil {
		return m.Data
	}
	return nil
}

func (m *RFIDprintWriteDataStream_Response) GetErrormsg() *ErrorMessage {
	if m != nil {
		return m.Errormsg
	}
	return nil
}

func init() {
	proto.RegisterType((*ErrorMessage)(nil), "RFIDDevice.ErrorMessage")
	proto.RegisterType((*DeviceRegeditRequest)(nil), "RFIDDevice.DeviceRegeditRequest")
	proto.RegisterType((*DeviceRegeditResponse)(nil), "RFIDDevice.DeviceRegeditResponse")
	proto.RegisterType((*RFIDprint_Data)(nil), "RFIDDevice.RFIDprint_Data")
	proto.RegisterType((*RFIDprintWriteDataStream_Request)(nil), "RFIDDevice.RFIDprint_writeDataStream_Request")
	proto.RegisterType((*RFIDprintWriteDataStream_Response)(nil), "RFIDDevice.RFIDprint_writeDataStream_Response")
	proto.RegisterEnum("RFIDDevice.ERRORCODE", ERRORCODE_name, ERRORCODE_value)
	proto.RegisterEnum("RFIDDevice.DEVICETYPE", DEVICETYPE_name, DEVICETYPE_value)
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConn

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion4

// Client API for RFIDDeviceServer service

type RFIDDeviceServerClient interface {
	// DeviceRegedit
	DeviceRegedit(ctx context.Context, in *DeviceRegeditRequest, opts ...grpc.CallOption) (*DeviceRegeditResponse, error)
	RFIDprintWriteDataStream(ctx context.Context, in *RFIDprintWriteDataStream_Request, opts ...grpc.CallOption) (RFIDDeviceServer_RFIDprintWriteDataStreamClient, error)
}

type rFIDDeviceServerClient struct {
	cc *grpc.ClientConn
}

func NewRFIDDeviceServerClient(cc *grpc.ClientConn) RFIDDeviceServerClient {
	return &rFIDDeviceServerClient{cc}
}

func (c *rFIDDeviceServerClient) DeviceRegedit(ctx context.Context, in *DeviceRegeditRequest, opts ...grpc.CallOption) (*DeviceRegeditResponse, error) {
	out := new(DeviceRegeditResponse)
	err := grpc.Invoke(ctx, "/RFIDDevice.RFIDDeviceServer/DeviceRegedit", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *rFIDDeviceServerClient) RFIDprintWriteDataStream(ctx context.Context, in *RFIDprintWriteDataStream_Request, opts ...grpc.CallOption) (RFIDDeviceServer_RFIDprintWriteDataStreamClient, error) {
	stream, err := grpc.NewClientStream(ctx, &_RFIDDeviceServer_serviceDesc.Streams[0], c.cc, "/RFIDDevice.RFIDDeviceServer/RFIDprint_writeData_stream", opts...)
	if err != nil {
		return nil, err
	}
	x := &rFIDDeviceServerRFIDprintWriteDataStreamClient{stream}
	if err := x.ClientStream.SendMsg(in); err != nil {
		return nil, err
	}
	if err := x.ClientStream.CloseSend(); err != nil {
		return nil, err
	}
	return x, nil
}

type RFIDDeviceServer_RFIDprintWriteDataStreamClient interface {
	Recv() (*RFIDprintWriteDataStream_Response, error)
	grpc.ClientStream
}

type rFIDDeviceServerRFIDprintWriteDataStreamClient struct {
	grpc.ClientStream
}

func (x *rFIDDeviceServerRFIDprintWriteDataStreamClient) Recv() (*RFIDprintWriteDataStream_Response, error) {
	m := new(RFIDprintWriteDataStream_Response)
	if err := x.ClientStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

// Server API for RFIDDeviceServer service

type RFIDDeviceServerServer interface {
	// DeviceRegedit
	DeviceRegedit(context.Context, *DeviceRegeditRequest) (*DeviceRegeditResponse, error)
	RFIDprintWriteDataStream(*RFIDprintWriteDataStream_Request, RFIDDeviceServer_RFIDprintWriteDataStreamServer) error
}

func RegisterRFIDDeviceServerServer(s *grpc.Server, srv RFIDDeviceServerServer) {
	s.RegisterService(&_RFIDDeviceServer_serviceDesc, srv)
}

func _RFIDDeviceServer_DeviceRegedit_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(DeviceRegeditRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(RFIDDeviceServerServer).DeviceRegedit(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/RFIDDevice.RFIDDeviceServer/DeviceRegedit",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(RFIDDeviceServerServer).DeviceRegedit(ctx, req.(*DeviceRegeditRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _RFIDDeviceServer_RFIDprintWriteDataStream_Handler(srv interface{}, stream grpc.ServerStream) error {
	m := new(RFIDprintWriteDataStream_Request)
	if err := stream.RecvMsg(m); err != nil {
		return err
	}
	return srv.(RFIDDeviceServerServer).RFIDprintWriteDataStream(m, &rFIDDeviceServerRFIDprintWriteDataStreamServer{stream})
}

type RFIDDeviceServer_RFIDprintWriteDataStreamServer interface {
	Send(*RFIDprintWriteDataStream_Response) error
	grpc.ServerStream
}

type rFIDDeviceServerRFIDprintWriteDataStreamServer struct {
	grpc.ServerStream
}

func (x *rFIDDeviceServerRFIDprintWriteDataStreamServer) Send(m *RFIDprintWriteDataStream_Response) error {
	return x.ServerStream.SendMsg(m)
}

var _RFIDDeviceServer_serviceDesc = grpc.ServiceDesc{
	ServiceName: "RFIDDevice.RFIDDeviceServer",
	HandlerType: (*RFIDDeviceServerServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "DeviceRegedit",
			Handler:    _RFIDDeviceServer_DeviceRegedit_Handler,
		},
	},
	Streams: []grpc.StreamDesc{
		{
			StreamName:    "RFIDprint_writeData_stream",
			Handler:       _RFIDDeviceServer_RFIDprintWriteDataStream_Handler,
			ServerStreams: true,
		},
	},
	Metadata: "DeviceServer.proto",
}

func init() { proto.RegisterFile("DeviceServer.proto", fileDescriptor0) }

var fileDescriptor0 = []byte{
	// 724 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x9c, 0x54, 0xdd, 0x6e, 0xd3, 0x48,
	0x14, 0xae, 0x9b, 0x1f, 0x27, 0x27, 0xdb, 0xc8, 0x9a, 0x6d, 0xb7, 0xde, 0x68, 0xd5, 0xa6, 0xd6,
	0xae, 0x54, 0x55, 0x5a, 0x0b, 0xa5, 0x08, 0x24, 0xee, 0x52, 0x7b, 0x2a, 0x05, 0x11, 0x27, 0x9a,
	0xa4, 0x15, 0xe5, 0xc6, 0x32, 0x9e, 0x51, 0xb0, 0x54, 0xc7, 0xc6, 0xe3, 0x16, 0xc2, 0x0d, 0x6f,
	0xc0, 0x43, 0x20, 0x5e, 0x02, 0xde, 0x08, 0x0a, 0x3c, 0x01, 0x17, 0x68, 0xc6, 0xf9, 0x71, 0x4a,
	0x5a, 0x10, 0x57, 0xf1, 0xf9, 0xce, 0x99, 0x33, 0x5f, 0xbe, 0xef, 0xcc, 0x01, 0x64, 0xb3, 0xcb,
	0xc0, 0x67, 0x03, 0x96, 0x5c, 0xb2, 0xc4, 0x8c, 0x93, 0x28, 0x8d, 0x10, 0x90, 0xe3, 0x8e, 0x9d,
	0xe1, 0x86, 0x0b, 0x7f, 0xe0, 0x24, 0x89, 0x92, 0x2e, 0xe3, 0xdc, 0x1b, 0x31, 0x74, 0x08, 0x55,
	0x26, 0x62, 0x2b, 0xa2, 0x4c, 0x57, 0x9a, 0xca, 0x7e, 0xbd, 0xb5, 0x65, 0x2e, 0xea, 0x4d, 0x4c,
	0x48, 0x8f, 0x58, 0x3d, 0x1b, 0x93, 0x45, 0x1d, 0x6a, 0x40, 0x45, 0x06, 0x5d, 0x3e, 0xd2, 0xd7,
	0x9b, 0xca, 0x7e, 0x95, 0xcc, 0x63, 0xe3, 0xbd, 0x02, 0x9b, 0xd9, 0x59, 0xc2, 0x46, 0x8c, 0x06,
	0x29, 0x61, 0xcf, 0x2f, 0x18, 0x4f, 0xd1, 0x2e, 0xd4, 0xa8, 0xc4, 0xdd, 0xb1, 0x17, 0x66, 0x77,
	0x55, 0x09, 0x64, 0x90, 0xe3, 0x85, 0x0c, 0xdd, 0x9f, 0x17, 0xa4, 0x93, 0x98, 0xc9, 0xc6, 0xf5,
	0xd6, 0x5f, 0x79, 0x32, 0x36, 0x3e, 0xed, 0x58, 0x78, 0x78, 0xd6, 0xc7, 0xb3, 0x83, 0xc3, 0x49,
	0xcc, 0xd0, 0x16, 0x94, 0x83, 0xd8, 0xf5, 0x28, 0xd5, 0x0b, 0xb2, 0x69, 0x29, 0x88, 0xdb, 0x94,
	0xa2, 0x6d, 0x50, 0x83, 0xd8, 0x8d, 0xa3, 0x24, 0xd5, 0x8b, 0x4d, 0x65, 0xbf, 0x44, 0xca, 0x41,
	0xdc, 0x8f, 0x92, 0x54, 0x24, 0x42, 0xcf, 0x97, 0x07, 0x4a, 0xf2, 0x40, 0x39, 0xf4, 0xfc, 0x36,
	0xa5, 0xc6, 0x1b, 0x05, 0xb6, 0xae, 0x71, 0xe7, 0x71, 0x34, 0xe6, 0x0c, 0xe9, 0xa0, 0xf2, 0x0b,
	0xdf, 0x67, 0x9c, 0x4b, 0xe2, 0x15, 0x32, 0x0b, 0xd1, 0x7f, 0x50, 0x9f, 0xb2, 0x4e, 0xe4, 0x19,
	0x2a, 0x89, 0x17, 0xc9, 0x06, 0xcd, 0x35, 0xa2, 0xe8, 0xee, 0x54, 0xb2, 0x90, 0x8f, 0x24, 0xcb,
	0x5a, 0x4b, 0x5f, 0x92, 0x39, 0xe7, 0x09, 0x99, 0x57, 0x1a, 0xdf, 0x14, 0xa8, 0x8b, 0xaa, 0x38,
	0x09, 0xc6, 0xa9, 0x6b, 0x7b, 0xa9, 0x27, 0xb4, 0xf7, 0x38, 0x67, 0x29, 0xef, 0xd0, 0xa9, 0x86,
	0xf3, 0x18, 0x35, 0xa1, 0xe6, 0x47, 0x61, 0xec, 0x8d, 0x27, 0x42, 0xd0, 0xa9, 0x35, 0x79, 0x08,
	0xfd, 0x0b, 0x82, 0x57, 0xcc, 0x42, 0x36, 0x4e, 0x9d, 0xd0, 0x63, 0x53, 0xc5, 0x96, 0x41, 0xb4,
	0x03, 0x90, 0xf5, 0x94, 0x6d, 0x8a, 0x99, 0x53, 0x0b, 0x64, 0x91, 0x17, 0xf2, 0x4f, 0x35, 0xcc,
	0x21, 0x82, 0x47, 0x16, 0x75, 0x23, 0xca, 0xce, 0xf5, 0x72, 0xc6, 0x23, 0x07, 0x21, 0x0d, 0x0a,
	0x17, 0xc9, 0xb9, 0xae, 0xca, 0x8c, 0xf8, 0x44, 0x08, 0x8a, 0xd4, 0x4b, 0x99, 0x5e, 0x91, 0x90,
	0xfc, 0x36, 0x1e, 0xc2, 0xde, 0xe2, 0xdf, 0xbf, 0x48, 0x82, 0x94, 0x09, 0x09, 0x06, 0x69, 0xc2,
	0xbc, 0xd0, 0x9d, 0xcd, 0xd5, 0x8f, 0x06, 0x08, 0x59, 0x0a, 0xd7, 0x0c, 0x30, 0xde, 0x29, 0x60,
	0xdc, 0xd6, 0xec, 0xa7, 0x46, 0x9b, 0x92, 0xa0, 0x27, 0x55, 0xad, 0xb5, 0x1a, 0x79, 0xf7, 0x96,
	0x2d, 0x92, 0xe4, 0xbd, 0xdf, 0x73, 0xfc, 0xe0, 0x83, 0x02, 0xd5, 0xf9, 0x9b, 0x43, 0x08, 0xea,
	0xf3, 0xc0, 0x75, 0x7a, 0x0e, 0xd6, 0xd6, 0xd0, 0x9f, 0x50, 0xc7, 0x8f, 0xad, 0x93, 0x61, 0xa7,
	0xe7, 0xb8, 0x32, 0xa9, 0x7d, 0x54, 0x05, 0xd8, 0xb6, 0x2c, 0x3c, 0x18, 0xb8, 0xc7, 0xed, 0xce,
	0xa3, 0x13, 0x82, 0xb5, 0x4f, 0x2a, 0xd2, 0xa0, 0x36, 0x38, 0x1b, 0x0c, 0x71, 0xd7, 0x3d, 0x3a,
	0x19, 0x9c, 0x69, 0x57, 0x2a, 0xfa, 0x07, 0xb6, 0x45, 0x1e, 0xdb, 0xae, 0xd5, 0x73, 0x1c, 0x6c,
	0x0d, 0x5d, 0xc1, 0xc8, 0x6a, 0x13, 0x5b, 0xfb, 0xac, 0xa2, 0x1d, 0xf8, 0xfb, 0x86, 0x2c, 0x26,
	0xda, 0x17, 0x15, 0xed, 0x42, 0x63, 0x45, 0xbe, 0x4f, 0x3a, 0xce, 0x10, 0x13, 0xed, 0xab, 0x7a,
	0xf0, 0x00, 0x60, 0xf1, 0x44, 0xd1, 0x06, 0x54, 0xc5, 0xef, 0x8c, 0x77, 0x15, 0x4a, 0xb2, 0x54,
	0x53, 0x50, 0x05, 0x8a, 0xf2, 0xce, 0x75, 0xf1, 0x45, 0x70, 0xdb, 0xd6, 0x0a, 0xad, 0x2b, 0x05,
	0xb4, 0x85, 0x3c, 0xd9, 0xfe, 0x42, 0xa7, 0xb0, 0xb1, 0xf4, 0x1e, 0x51, 0x73, 0x69, 0x1d, 0xac,
	0x58, 0x33, 0x8d, 0xbd, 0x5b, 0x2a, 0x32, 0x8f, 0x8d, 0x35, 0xf4, 0x1a, 0x1a, 0x2b, 0x66, 0xc1,
	0xe5, 0x72, 0x18, 0xd0, 0xff, 0xab, 0xbd, 0xbd, 0x61, 0x00, 0x1b, 0xe6, 0xaf, 0x96, 0xcf, 0xae,
	0xbf, 0xa3, 0x1c, 0xdd, 0x83, 0x4d, 0x3f, 0x0a, 0xcd, 0x67, 0x2f, 0xa9, 0xf9, 0xca, 0xe7, 0x13,
	0x6e, 0x66, 0xc3, 0x7a, 0x94, 0x5b, 0xd5, 0x7d, 0xe5, 0x49, 0x2e, 0x7a, 0xbb, 0x5e, 0xb0, 0xf1,
	0xe9, 0xd3, 0xb2, 0xdc, 0xe8, 0x87, 0xdf, 0x03, 0x00, 0x00, 0xff, 0xff, 0xe0, 0xcd, 0x33, 0xfc,
	0xe7, 0x05, 0x00, 0x00,
}
