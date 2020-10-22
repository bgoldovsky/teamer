// Code generated by protoc-gen-go. DO NOT EDIT.
// source: duties.proto

package v1

import (
	context "context"
	fmt "fmt"
	proto "github.com/golang/protobuf/proto"
	empty "github.com/golang/protobuf/ptypes/empty"
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

type GetDutiesRequest struct {
	TeamId               int64    `protobuf:"varint,1,opt,name=team_id,json=teamId,proto3" json:"team_id,omitempty"`
	Count                int64    `protobuf:"varint,2,opt,name=count,proto3" json:"count,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *GetDutiesRequest) Reset()         { *m = GetDutiesRequest{} }
func (m *GetDutiesRequest) String() string { return proto.CompactTextString(m) }
func (*GetDutiesRequest) ProtoMessage()    {}
func (*GetDutiesRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_971c6befef230320, []int{2}
}

func (m *GetDutiesRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_GetDutiesRequest.Unmarshal(m, b)
}
func (m *GetDutiesRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_GetDutiesRequest.Marshal(b, m, deterministic)
}
func (m *GetDutiesRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_GetDutiesRequest.Merge(m, src)
}
func (m *GetDutiesRequest) XXX_Size() int {
	return xxx_messageInfo_GetDutiesRequest.Size(m)
}
func (m *GetDutiesRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_GetDutiesRequest.DiscardUnknown(m)
}

var xxx_messageInfo_GetDutiesRequest proto.InternalMessageInfo

func (m *GetDutiesRequest) GetTeamId() int64 {
	if m != nil {
		return m.TeamId
	}
	return 0
}

func (m *GetDutiesRequest) GetCount() int64 {
	if m != nil {
		return m.Count
	}
	return 0
}

type GetDutiesReply struct {
	Duties               []*Duty  `protobuf:"bytes,1,rep,name=duties,proto3" json:"duties,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *GetDutiesReply) Reset()         { *m = GetDutiesReply{} }
func (m *GetDutiesReply) String() string { return proto.CompactTextString(m) }
func (*GetDutiesReply) ProtoMessage()    {}
func (*GetDutiesReply) Descriptor() ([]byte, []int) {
	return fileDescriptor_971c6befef230320, []int{3}
}

func (m *GetDutiesReply) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_GetDutiesReply.Unmarshal(m, b)
}
func (m *GetDutiesReply) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_GetDutiesReply.Marshal(b, m, deterministic)
}
func (m *GetDutiesReply) XXX_Merge(src proto.Message) {
	xxx_messageInfo_GetDutiesReply.Merge(m, src)
}
func (m *GetDutiesReply) XXX_Size() int {
	return xxx_messageInfo_GetDutiesReply.Size(m)
}
func (m *GetDutiesReply) XXX_DiscardUnknown() {
	xxx_messageInfo_GetDutiesReply.DiscardUnknown(m)
}

var xxx_messageInfo_GetDutiesReply proto.InternalMessageInfo

func (m *GetDutiesReply) GetDuties() []*Duty {
	if m != nil {
		return m.Duties
	}
	return nil
}

type AssignRequest struct {
	TeamId               int64    `protobuf:"varint,1,opt,name=team_id,json=teamId,proto3" json:"team_id,omitempty"`
	PersonId             int64    `protobuf:"varint,2,opt,name=person_id,json=personId,proto3" json:"person_id,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *AssignRequest) Reset()         { *m = AssignRequest{} }
func (m *AssignRequest) String() string { return proto.CompactTextString(m) }
func (*AssignRequest) ProtoMessage()    {}
func (*AssignRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_971c6befef230320, []int{4}
}

func (m *AssignRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_AssignRequest.Unmarshal(m, b)
}
func (m *AssignRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_AssignRequest.Marshal(b, m, deterministic)
}
func (m *AssignRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_AssignRequest.Merge(m, src)
}
func (m *AssignRequest) XXX_Size() int {
	return xxx_messageInfo_AssignRequest.Size(m)
}
func (m *AssignRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_AssignRequest.DiscardUnknown(m)
}

var xxx_messageInfo_AssignRequest proto.InternalMessageInfo

func (m *AssignRequest) GetTeamId() int64 {
	if m != nil {
		return m.TeamId
	}
	return 0
}

func (m *AssignRequest) GetPersonId() int64 {
	if m != nil {
		return m.PersonId
	}
	return 0
}

type SwapRequest struct {
	TeamId               int64    `protobuf:"varint,1,opt,name=team_id,json=teamId,proto3" json:"team_id,omitempty"`
	FirstPersonId        int64    `protobuf:"varint,2,opt,name=first_person_id,json=firstPersonId,proto3" json:"first_person_id,omitempty"`
	SecondPersonId       int64    `protobuf:"varint,3,opt,name=second_person_id,json=secondPersonId,proto3" json:"second_person_id,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *SwapRequest) Reset()         { *m = SwapRequest{} }
func (m *SwapRequest) String() string { return proto.CompactTextString(m) }
func (*SwapRequest) ProtoMessage()    {}
func (*SwapRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_971c6befef230320, []int{5}
}

func (m *SwapRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_SwapRequest.Unmarshal(m, b)
}
func (m *SwapRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_SwapRequest.Marshal(b, m, deterministic)
}
func (m *SwapRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_SwapRequest.Merge(m, src)
}
func (m *SwapRequest) XXX_Size() int {
	return xxx_messageInfo_SwapRequest.Size(m)
}
func (m *SwapRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_SwapRequest.DiscardUnknown(m)
}

var xxx_messageInfo_SwapRequest proto.InternalMessageInfo

func (m *SwapRequest) GetTeamId() int64 {
	if m != nil {
		return m.TeamId
	}
	return 0
}

func (m *SwapRequest) GetFirstPersonId() int64 {
	if m != nil {
		return m.FirstPersonId
	}
	return 0
}

func (m *SwapRequest) GetSecondPersonId() int64 {
	if m != nil {
		return m.SecondPersonId
	}
	return 0
}

type Duty struct {
	TeamId               int64    `protobuf:"varint,1,opt,name=team_id,json=teamId,proto3" json:"team_id,omitempty"`
	PersonId             int64    `protobuf:"varint,2,opt,name=person_id,json=personId,proto3" json:"person_id,omitempty"`
	FirstName            string   `protobuf:"bytes,3,opt,name=first_name,json=firstName,proto3" json:"first_name,omitempty"`
	LastName             string   `protobuf:"bytes,4,opt,name=last_name,json=lastName,proto3" json:"last_name,omitempty"`
	Slack                string   `protobuf:"bytes,5,opt,name=slack,proto3" json:"slack,omitempty"`
	DutyOrder            int64    `protobuf:"varint,6,opt,name=duty_order,json=dutyOrder,proto3" json:"duty_order,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *Duty) Reset()         { *m = Duty{} }
func (m *Duty) String() string { return proto.CompactTextString(m) }
func (*Duty) ProtoMessage()    {}
func (*Duty) Descriptor() ([]byte, []int) {
	return fileDescriptor_971c6befef230320, []int{6}
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

func (m *Duty) GetDutyOrder() int64 {
	if m != nil {
		return m.DutyOrder
	}
	return 0
}

func init() {
	proto.RegisterType((*GetCurrentDutyRequest)(nil), "api.GetCurrentDutyRequest")
	proto.RegisterType((*GetCurrentDutyReply)(nil), "api.GetCurrentDutyReply")
	proto.RegisterType((*GetDutiesRequest)(nil), "api.GetDutiesRequest")
	proto.RegisterType((*GetDutiesReply)(nil), "api.GetDutiesReply")
	proto.RegisterType((*AssignRequest)(nil), "api.AssignRequest")
	proto.RegisterType((*SwapRequest)(nil), "api.SwapRequest")
	proto.RegisterType((*Duty)(nil), "api.Duty")
}

func init() { proto.RegisterFile("duties.proto", fileDescriptor_971c6befef230320) }

var fileDescriptor_971c6befef230320 = []byte{
	// 443 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x9c, 0x53, 0xcf, 0x6f, 0xd3, 0x30,
	0x14, 0x5e, 0x68, 0x1b, 0x96, 0x37, 0x36, 0x2a, 0x6f, 0x83, 0x28, 0x53, 0xc5, 0xf0, 0x01, 0xed,
	0x94, 0x40, 0x87, 0x90, 0x38, 0x0e, 0x98, 0xc6, 0x2e, 0x80, 0xca, 0x8d, 0x4b, 0xe5, 0x35, 0x6f,
	0x95, 0x45, 0xea, 0x18, 0xdb, 0x01, 0xf2, 0x47, 0xf1, 0x07, 0x72, 0x43, 0x7e, 0x5e, 0x51, 0x5a,
	0x15, 0x55, 0xe2, 0xe8, 0xef, 0xc7, 0xfb, 0x6c, 0x7d, 0xcf, 0xf0, 0xa0, 0x6c, 0x9c, 0x44, 0x9b,
	0x6b, 0x53, 0xbb, 0x9a, 0xf5, 0x84, 0x96, 0xd9, 0xc9, 0xbc, 0xae, 0xe7, 0x15, 0x16, 0x04, 0xdd,
	0x34, 0xb7, 0x05, 0x2e, 0xb4, 0x6b, 0x83, 0x82, 0x3f, 0x87, 0xe3, 0x2b, 0x74, 0x6f, 0x1b, 0x63,
	0x50, 0xb9, 0x77, 0x8d, 0x6b, 0x27, 0xf8, 0xad, 0x41, 0xeb, 0xd8, 0x63, 0xb8, 0xef, 0x50, 0x2c,
	0xa6, 0xb2, 0x4c, 0xa3, 0xd3, 0xe8, 0xac, 0x37, 0x89, 0xfd, 0xf1, 0xba, 0xe4, 0x2f, 0xe1, 0x70,
	0xdd, 0xa1, 0xab, 0x96, 0x8d, 0xa0, 0x5f, 0x36, 0xae, 0x25, 0xf1, 0xde, 0x38, 0xc9, 0x85, 0x96,
	0x39, 0xb1, 0x04, 0xf3, 0x0b, 0x18, 0x5e, 0xa1, 0x97, 0x4b, 0xb4, 0xdb, 0x22, 0xd8, 0x11, 0x0c,
	0x66, 0x75, 0xa3, 0x5c, 0x7a, 0x8f, 0xe0, 0x70, 0xe0, 0xe7, 0x70, 0xd0, 0x19, 0xe1, 0x33, 0x9f,
	0x42, 0x1c, 0x9e, 0x9b, 0x46, 0xa7, 0xbd, 0xd5, 0xd4, 0x3b, 0x82, 0x5f, 0xc2, 0xfe, 0x85, 0xb5,
	0x72, 0xae, 0xb6, 0x86, 0x9e, 0x40, 0xa2, 0xd1, 0xd8, 0x5a, 0x79, 0x2a, 0x04, 0xef, 0x06, 0xe0,
	0xba, 0xe4, 0x3f, 0x61, 0xef, 0xf3, 0x0f, 0xa1, 0xb7, 0x0e, 0x79, 0x06, 0x0f, 0x6f, 0xa5, 0xb1,
	0x6e, 0xba, 0x3e, 0x6a, 0x9f, 0xe0, 0x4f, 0x77, 0xf3, 0xd8, 0x19, 0x0c, 0x2d, 0xce, 0x6a, 0x55,
	0x76, 0x84, 0x3d, 0x12, 0x1e, 0x04, 0x7c, 0xa9, 0xe4, 0xbf, 0x22, 0xe8, 0xfb, 0x17, 0xfd, 0xdf,
	0xc5, 0xd9, 0x08, 0x20, 0x5c, 0x48, 0x89, 0x05, 0x52, 0x44, 0x32, 0x49, 0x08, 0xf9, 0x20, 0x16,
	0xe8, 0xbd, 0x95, 0x58, 0xb2, 0x7d, 0x62, 0x77, 0x3d, 0x40, 0xe4, 0x11, 0x0c, 0x6c, 0x25, 0x66,
	0x5f, 0xd3, 0x01, 0x11, 0xe1, 0xe0, 0x27, 0xfa, 0x46, 0xa7, 0xb5, 0x29, 0xd1, 0xa4, 0x31, 0xe5,
	0x25, 0x1e, 0xf9, 0xe8, 0x81, 0xf1, 0xef, 0x08, 0xe2, 0xd0, 0x11, 0x7b, 0x4f, 0x85, 0x75, 0x36,
	0x85, 0x65, 0x54, 0xd0, 0xc6, 0x85, 0xcb, 0xd2, 0x8d, 0x9c, 0xae, 0x5a, 0xbe, 0xc3, 0x5e, 0x43,
	0xf2, 0xb7, 0x7a, 0x76, 0xbc, 0x14, 0xae, 0x6c, 0x53, 0x76, 0xb8, 0x0e, 0x07, 0xeb, 0x2b, 0x88,
	0xc3, 0x02, 0x30, 0x46, 0x82, 0x95, 0x6d, 0xc8, 0x1e, 0xe5, 0xe1, 0x73, 0xe4, 0xcb, 0xcf, 0x91,
	0x5f, 0xfa, 0xcf, 0xc1, 0x77, 0xd8, 0x18, 0xfa, 0xbe, 0x71, 0x36, 0x24, 0x57, 0xa7, 0xfc, 0x7f,
	0x7b, 0xde, 0x3c, 0xf9, 0x32, 0x92, 0xca, 0xa1, 0x51, 0xa2, 0x2a, 0xe6, 0xa8, 0xd0, 0x08, 0x87,
	0x65, 0x31, 0xab, 0x24, 0x2a, 0x67, 0x8b, 0xef, 0x2f, 0x6e, 0x62, 0xb2, 0x9c, 0xff, 0x09, 0x00,
	0x00, 0xff, 0xff, 0x15, 0x30, 0xea, 0x50, 0xa6, 0x03, 0x00, 0x00,
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
	GetDuties(ctx context.Context, in *GetDutiesRequest, opts ...grpc.CallOption) (*GetDutiesReply, error)
	Assign(ctx context.Context, in *AssignRequest, opts ...grpc.CallOption) (*empty.Empty, error)
	Swap(ctx context.Context, in *SwapRequest, opts ...grpc.CallOption) (*empty.Empty, error)
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

func (c *dutiesClient) GetDuties(ctx context.Context, in *GetDutiesRequest, opts ...grpc.CallOption) (*GetDutiesReply, error) {
	out := new(GetDutiesReply)
	err := c.cc.Invoke(ctx, "/api.Duties/GetDuties", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *dutiesClient) Assign(ctx context.Context, in *AssignRequest, opts ...grpc.CallOption) (*empty.Empty, error) {
	out := new(empty.Empty)
	err := c.cc.Invoke(ctx, "/api.Duties/Assign", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *dutiesClient) Swap(ctx context.Context, in *SwapRequest, opts ...grpc.CallOption) (*empty.Empty, error) {
	out := new(empty.Empty)
	err := c.cc.Invoke(ctx, "/api.Duties/Swap", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// DutiesServer is the server API for Duties service.
type DutiesServer interface {
	GetCurrentDuty(context.Context, *GetCurrentDutyRequest) (*GetCurrentDutyReply, error)
	GetDuties(context.Context, *GetDutiesRequest) (*GetDutiesReply, error)
	Assign(context.Context, *AssignRequest) (*empty.Empty, error)
	Swap(context.Context, *SwapRequest) (*empty.Empty, error)
}

// UnimplementedDutiesServer can be embedded to have forward compatible implementations.
type UnimplementedDutiesServer struct {
}

func (*UnimplementedDutiesServer) GetCurrentDuty(ctx context.Context, req *GetCurrentDutyRequest) (*GetCurrentDutyReply, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetCurrentDuty not implemented")
}
func (*UnimplementedDutiesServer) GetDuties(ctx context.Context, req *GetDutiesRequest) (*GetDutiesReply, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetDuties not implemented")
}
func (*UnimplementedDutiesServer) Assign(ctx context.Context, req *AssignRequest) (*empty.Empty, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Assign not implemented")
}
func (*UnimplementedDutiesServer) Swap(ctx context.Context, req *SwapRequest) (*empty.Empty, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Swap not implemented")
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

func _Duties_GetDuties_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetDutiesRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(DutiesServer).GetDuties(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/api.Duties/GetDuties",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(DutiesServer).GetDuties(ctx, req.(*GetDutiesRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Duties_Assign_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(AssignRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(DutiesServer).Assign(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/api.Duties/Assign",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(DutiesServer).Assign(ctx, req.(*AssignRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Duties_Swap_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(SwapRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(DutiesServer).Swap(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/api.Duties/Swap",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(DutiesServer).Swap(ctx, req.(*SwapRequest))
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
		{
			MethodName: "GetDuties",
			Handler:    _Duties_GetDuties_Handler,
		},
		{
			MethodName: "Assign",
			Handler:    _Duties_Assign_Handler,
		},
		{
			MethodName: "Swap",
			Handler:    _Duties_Swap_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "duties.proto",
}
