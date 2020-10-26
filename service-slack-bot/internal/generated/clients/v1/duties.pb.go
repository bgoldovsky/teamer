// Code generated by protoc-gen-go. DO NOT EDIT.
// source: duties.proto

package v1

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

type GetCurrentDutyRequest struct {
	TeamId               int64    `protobuf:"varint,1,opt,name=team_id,json=teamId,proto3" json:"team_id,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *GetCurrentDutyRequest) Reset()         { *m = GetCurrentDutyRequest{} }
func (m *GetCurrentDutyRequest) String() string { return proto.CompactTextString(m) }
func (*GetCurrentDutyRequest) ProtoMessage()    {}
func (*GetCurrentDutyRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_971c6befef230320, []int{0}
}

func (m *GetCurrentDutyRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_GetCurrentDutyRequest.Unmarshal(m, b)
}
func (m *GetCurrentDutyRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_GetCurrentDutyRequest.Marshal(b, m, deterministic)
}
func (m *GetCurrentDutyRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_GetCurrentDutyRequest.Merge(m, src)
}
func (m *GetCurrentDutyRequest) XXX_Size() int {
	return xxx_messageInfo_GetCurrentDutyRequest.Size(m)
}
func (m *GetCurrentDutyRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_GetCurrentDutyRequest.DiscardUnknown(m)
}

var xxx_messageInfo_GetCurrentDutyRequest proto.InternalMessageInfo

func (m *GetCurrentDutyRequest) GetTeamId() int64 {
	if m != nil {
		return m.TeamId
	}
	return 0
}

type GetCurrentDutyReply struct {
	Duty                 *Duty    `protobuf:"bytes,1,opt,name=duty,proto3" json:"duty,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *GetCurrentDutyReply) Reset()         { *m = GetCurrentDutyReply{} }
func (m *GetCurrentDutyReply) String() string { return proto.CompactTextString(m) }
func (*GetCurrentDutyReply) ProtoMessage()    {}
func (*GetCurrentDutyReply) Descriptor() ([]byte, []int) {
	return fileDescriptor_971c6befef230320, []int{1}
}

func (m *GetCurrentDutyReply) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_GetCurrentDutyReply.Unmarshal(m, b)
}
func (m *GetCurrentDutyReply) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_GetCurrentDutyReply.Marshal(b, m, deterministic)
}
func (m *GetCurrentDutyReply) XXX_Merge(src proto.Message) {
	xxx_messageInfo_GetCurrentDutyReply.Merge(m, src)
}
func (m *GetCurrentDutyReply) XXX_Size() int {
	return xxx_messageInfo_GetCurrentDutyReply.Size(m)
}
func (m *GetCurrentDutyReply) XXX_DiscardUnknown() {
	xxx_messageInfo_GetCurrentDutyReply.DiscardUnknown(m)
}

var xxx_messageInfo_GetCurrentDutyReply proto.InternalMessageInfo

func (m *GetCurrentDutyReply) GetDuty() *Duty {
	if m != nil {
		return m.Duty
	}
	return nil
}

type Duty struct {
	TeamId               int64    `protobuf:"varint,1,opt,name=team_id,json=teamId,proto3" json:"team_id,omitempty"`
	PersonId             int64    `protobuf:"varint,2,opt,name=person_id,json=personId,proto3" json:"person_id,omitempty"`
	FirstName            string   `protobuf:"bytes,3,opt,name=first_name,json=firstName,proto3" json:"first_name,omitempty"`
	LastName             string   `protobuf:"bytes,4,opt,name=last_name,json=lastName,proto3" json:"last_name,omitempty"`
	Slack                string   `protobuf:"bytes,5,opt,name=slack,proto3" json:"slack,omitempty"`
	Channel              string   `protobuf:"bytes,6,opt,name=channel,proto3" json:"channel,omitempty"`
	DutyOrder            int64    `protobuf:"varint,7,opt,name=duty_order,json=dutyOrder,proto3" json:"duty_order,omitempty"`
	Month                int64    `protobuf:"varint,8,opt,name=month,proto3" json:"month,omitempty"`
	Day                  int64    `protobuf:"varint,9,opt,name=day,proto3" json:"day,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *Duty) Reset()         { *m = Duty{} }
func (m *Duty) String() string { return proto.CompactTextString(m) }
func (*Duty) ProtoMessage()    {}
func (*Duty) Descriptor() ([]byte, []int) {
	return fileDescriptor_971c6befef230320, []int{2}
}

func (m *Duty) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Duty.Unmarshal(m, b)
}
func (m *Duty) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Duty.Marshal(b, m, deterministic)
}
func (m *Duty) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Duty.Merge(m, src)
}
func (m *Duty) XXX_Size() int {
	return xxx_messageInfo_Duty.Size(m)
}
func (m *Duty) XXX_DiscardUnknown() {
	xxx_messageInfo_Duty.DiscardUnknown(m)
}

var xxx_messageInfo_Duty proto.InternalMessageInfo

func (m *Duty) GetTeamId() int64 {
	if m != nil {
		return m.TeamId
	}
	return 0
}

func (m *Duty) GetPersonId() int64 {
	if m != nil {
		return m.PersonId
	}
	return 0
}

func (m *Duty) GetFirstName() string {
	if m != nil {
		return m.FirstName
	}
	return ""
}

func (m *Duty) GetLastName() string {
	if m != nil {
		return m.LastName
	}
	return ""
}

func (m *Duty) GetSlack() string {
	if m != nil {
		return m.Slack
	}
	return ""
}

func (m *Duty) GetChannel() string {
	if m != nil {
		return m.Channel
	}
	return ""
}

func (m *Duty) GetDutyOrder() int64 {
	if m != nil {
		return m.DutyOrder
	}
	return 0
}

func (m *Duty) GetMonth() int64 {
	if m != nil {
		return m.Month
	}
	return 0
}

func (m *Duty) GetDay() int64 {
	if m != nil {
		return m.Day
	}
	return 0
}

func init() {
	proto.RegisterType((*GetCurrentDutyRequest)(nil), "api.GetCurrentDutyRequest")
	proto.RegisterType((*GetCurrentDutyReply)(nil), "api.GetCurrentDutyReply")
	proto.RegisterType((*Duty)(nil), "api.Duty")
}

func init() { proto.RegisterFile("duties.proto", fileDescriptor_971c6befef230320) }

var fileDescriptor_971c6befef230320 = []byte{
	// 308 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x74, 0x91, 0xc1, 0x4a, 0xc3, 0x40,
	0x10, 0x86, 0x8d, 0x69, 0xd3, 0x66, 0x14, 0x91, 0x55, 0x71, 0xa9, 0x14, 0x4b, 0x4e, 0x3d, 0xb5,
	0x5a, 0x7d, 0x02, 0x2d, 0x68, 0x2f, 0x0a, 0x39, 0x7a, 0x29, 0x6b, 0x77, 0xb4, 0x8b, 0x9b, 0x4d,
	0xdc, 0x4c, 0x84, 0x3c, 0xb5, 0xaf, 0x20, 0x3b, 0xa5, 0x07, 0xa5, 0xde, 0xf2, 0x7f, 0x5f, 0x66,
	0x32, 0xfc, 0x81, 0x43, 0xdd, 0x90, 0xc1, 0x7a, 0x52, 0xf9, 0x92, 0x4a, 0x11, 0xab, 0xca, 0x64,
	0x57, 0x70, 0xf6, 0x80, 0x74, 0xdf, 0x78, 0x8f, 0x8e, 0xe6, 0x0d, 0xb5, 0x39, 0x7e, 0x36, 0x58,
	0x93, 0x38, 0x87, 0x1e, 0xa1, 0x2a, 0x96, 0x46, 0xcb, 0x68, 0x14, 0x8d, 0xe3, 0x3c, 0x09, 0x71,
	0xa1, 0xb3, 0x5b, 0x38, 0xf9, 0x3b, 0x51, 0xd9, 0x56, 0x0c, 0xa1, 0xa3, 0x1b, 0x6a, 0xf9, 0xe5,
	0x83, 0x59, 0x3a, 0x51, 0x95, 0x99, 0xb0, 0x65, 0x9c, 0x7d, 0x47, 0xd0, 0x09, 0xf1, 0xdf, 0xbd,
	0xe2, 0x02, 0xd2, 0x0a, 0x7d, 0x5d, 0xba, 0xa0, 0xf6, 0x59, 0xf5, 0x37, 0x60, 0xa1, 0xc5, 0x10,
	0xe0, 0xcd, 0xf8, 0x9a, 0x96, 0x4e, 0x15, 0x28, 0xe3, 0x51, 0x34, 0x4e, 0xf3, 0x94, 0xc9, 0x93,
	0x2a, 0x30, 0xcc, 0x5a, 0xb5, 0xb5, 0x1d, 0xb6, 0xfd, 0x00, 0x58, 0x9e, 0x42, 0xb7, 0xb6, 0x6a,
	0xf5, 0x21, 0xbb, 0x2c, 0x36, 0x41, 0x48, 0xe8, 0xad, 0xd6, 0xca, 0x39, 0xb4, 0x32, 0x61, 0xbe,
	0x8d, 0xe1, 0x5b, 0xe1, 0xe4, 0x65, 0xe9, 0x35, 0x7a, 0xd9, 0xe3, 0x4b, 0xd2, 0x40, 0x9e, 0x03,
	0x08, 0xeb, 0x8a, 0xd2, 0xd1, 0x5a, 0xf6, 0xd9, 0x6c, 0x82, 0x38, 0x86, 0x58, 0xab, 0x56, 0xa6,
	0xcc, 0xc2, 0xe3, 0x2c, 0x87, 0x64, 0xce, 0x75, 0x8b, 0x47, 0x38, 0xfa, 0xdd, 0x98, 0x18, 0x70,
	0x3d, 0x3b, 0x8b, 0x1f, 0xc8, 0x9d, 0xae, 0xb2, 0x6d, 0xb6, 0x77, 0x77, 0xf9, 0x32, 0x34, 0x8e,
	0xd0, 0x3b, 0x65, 0xa7, 0xef, 0xe8, 0xd0, 0x2b, 0x42, 0x3d, 0x5d, 0x59, 0x83, 0x8e, 0xea, 0xe9,
	0xd7, 0xf5, 0x6b, 0xc2, 0xbf, 0xf6, 0xe6, 0x27, 0x00, 0x00, 0xff, 0xff, 0x12, 0xcb, 0xef, 0x86,
	0xea, 0x01, 0x00, 0x00,
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConnInterface

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion6

// DutiesClient is the client API for Duties service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type DutiesClient interface {
	GetCurrentDuty(ctx context.Context, in *GetCurrentDutyRequest, opts ...grpc.CallOption) (*GetCurrentDutyReply, error)
}

type dutiesClient struct {
	cc grpc.ClientConnInterface
}

func NewDutiesClient(cc grpc.ClientConnInterface) DutiesClient {
	return &dutiesClient{cc}
}

func (c *dutiesClient) GetCurrentDuty(ctx context.Context, in *GetCurrentDutyRequest, opts ...grpc.CallOption) (*GetCurrentDutyReply, error) {
	out := new(GetCurrentDutyReply)
	err := c.cc.Invoke(ctx, "/api.Duties/GetCurrentDuty", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// DutiesServer is the server API for Duties service.
type DutiesServer interface {
	GetCurrentDuty(context.Context, *GetCurrentDutyRequest) (*GetCurrentDutyReply, error)
}

// UnimplementedDutiesServer can be embedded to have forward compatible implementations.
type UnimplementedDutiesServer struct {
}

func (*UnimplementedDutiesServer) GetCurrentDuty(ctx context.Context, req *GetCurrentDutyRequest) (*GetCurrentDutyReply, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetCurrentDuty not implemented")
}

func RegisterDutiesServer(s *grpc.Server, srv DutiesServer) {
	s.RegisterService(&_Duties_serviceDesc, srv)
}

func _Duties_GetCurrentDuty_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetCurrentDutyRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(DutiesServer).GetCurrentDuty(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/api.Duties/GetCurrentDuty",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(DutiesServer).GetCurrentDuty(ctx, req.(*GetCurrentDutyRequest))
	}
	return interceptor(ctx, in, info, handler)
}

var _Duties_serviceDesc = grpc.ServiceDesc{
	ServiceName: "api.Duties",
	HandlerType: (*DutiesServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "GetCurrentDuty",
			Handler:    _Duties_GetCurrentDuty_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "duties.proto",
}
