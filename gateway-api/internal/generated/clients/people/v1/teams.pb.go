// Code generated by protoc-gen-go. DO NOT EDIT.
// source: teams.proto

package v1

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

type AddTeamRequest struct {
	Name                 string   `protobuf:"bytes,2,opt,name=name,proto3" json:"name,omitempty"`
	Description          string   `protobuf:"bytes,3,opt,name=description,proto3" json:"description,omitempty"`
	Slack                string   `protobuf:"bytes,4,opt,name=slack,proto3" json:"slack,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *AddTeamRequest) Reset()         { *m = AddTeamRequest{} }
func (m *AddTeamRequest) String() string { return proto.CompactTextString(m) }
func (*AddTeamRequest) ProtoMessage()    {}
func (*AddTeamRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_f63f4a1b2b4dddb4, []int{0}
}

func (m *AddTeamRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_AddTeamRequest.Unmarshal(m, b)
}
func (m *AddTeamRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_AddTeamRequest.Marshal(b, m, deterministic)
}
func (m *AddTeamRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_AddTeamRequest.Merge(m, src)
}
func (m *AddTeamRequest) XXX_Size() int {
	return xxx_messageInfo_AddTeamRequest.Size(m)
}
func (m *AddTeamRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_AddTeamRequest.DiscardUnknown(m)
}

var xxx_messageInfo_AddTeamRequest proto.InternalMessageInfo

func (m *AddTeamRequest) GetName() string {
	if m != nil {
		return m.Name
	}
	return ""
}

func (m *AddTeamRequest) GetDescription() string {
	if m != nil {
		return m.Description
	}
	return ""
}

func (m *AddTeamRequest) GetSlack() string {
	if m != nil {
		return m.Slack
	}
	return ""
}

type AddTeamReply struct {
	Id                   int64    `protobuf:"varint,1,opt,name=id,proto3" json:"id,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *AddTeamReply) Reset()         { *m = AddTeamReply{} }
func (m *AddTeamReply) String() string { return proto.CompactTextString(m) }
func (*AddTeamReply) ProtoMessage()    {}
func (*AddTeamReply) Descriptor() ([]byte, []int) {
	return fileDescriptor_f63f4a1b2b4dddb4, []int{1}
}

func (m *AddTeamReply) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_AddTeamReply.Unmarshal(m, b)
}
func (m *AddTeamReply) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_AddTeamReply.Marshal(b, m, deterministic)
}
func (m *AddTeamReply) XXX_Merge(src proto.Message) {
	xxx_messageInfo_AddTeamReply.Merge(m, src)
}
func (m *AddTeamReply) XXX_Size() int {
	return xxx_messageInfo_AddTeamReply.Size(m)
}
func (m *AddTeamReply) XXX_DiscardUnknown() {
	xxx_messageInfo_AddTeamReply.DiscardUnknown(m)
}

var xxx_messageInfo_AddTeamReply proto.InternalMessageInfo

func (m *AddTeamReply) GetId() int64 {
	if m != nil {
		return m.Id
	}
	return 0
}

type UpdateTeamRequest struct {
	Id                   int64    `protobuf:"varint,1,opt,name=id,proto3" json:"id,omitempty"`
	Name                 string   `protobuf:"bytes,2,opt,name=name,proto3" json:"name,omitempty"`
	Description          string   `protobuf:"bytes,3,opt,name=description,proto3" json:"description,omitempty"`
	Slack                string   `protobuf:"bytes,4,opt,name=slack,proto3" json:"slack,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *UpdateTeamRequest) Reset()         { *m = UpdateTeamRequest{} }
func (m *UpdateTeamRequest) String() string { return proto.CompactTextString(m) }
func (*UpdateTeamRequest) ProtoMessage()    {}
func (*UpdateTeamRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_f63f4a1b2b4dddb4, []int{2}
}

func (m *UpdateTeamRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_UpdateTeamRequest.Unmarshal(m, b)
}
func (m *UpdateTeamRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_UpdateTeamRequest.Marshal(b, m, deterministic)
}
func (m *UpdateTeamRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_UpdateTeamRequest.Merge(m, src)
}
func (m *UpdateTeamRequest) XXX_Size() int {
	return xxx_messageInfo_UpdateTeamRequest.Size(m)
}
func (m *UpdateTeamRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_UpdateTeamRequest.DiscardUnknown(m)
}

var xxx_messageInfo_UpdateTeamRequest proto.InternalMessageInfo

func (m *UpdateTeamRequest) GetId() int64 {
	if m != nil {
		return m.Id
	}
	return 0
}

func (m *UpdateTeamRequest) GetName() string {
	if m != nil {
		return m.Name
	}
	return ""
}

func (m *UpdateTeamRequest) GetDescription() string {
	if m != nil {
		return m.Description
	}
	return ""
}

func (m *UpdateTeamRequest) GetSlack() string {
	if m != nil {
		return m.Slack
	}
	return ""
}

type RemoveTeamRequest struct {
	Id                   int64    `protobuf:"varint,1,opt,name=id,proto3" json:"id,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *RemoveTeamRequest) Reset()         { *m = RemoveTeamRequest{} }
func (m *RemoveTeamRequest) String() string { return proto.CompactTextString(m) }
func (*RemoveTeamRequest) ProtoMessage()    {}
func (*RemoveTeamRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_f63f4a1b2b4dddb4, []int{3}
}

func (m *RemoveTeamRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_RemoveTeamRequest.Unmarshal(m, b)
}
func (m *RemoveTeamRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_RemoveTeamRequest.Marshal(b, m, deterministic)
}
func (m *RemoveTeamRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_RemoveTeamRequest.Merge(m, src)
}
func (m *RemoveTeamRequest) XXX_Size() int {
	return xxx_messageInfo_RemoveTeamRequest.Size(m)
}
func (m *RemoveTeamRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_RemoveTeamRequest.DiscardUnknown(m)
}

var xxx_messageInfo_RemoveTeamRequest proto.InternalMessageInfo

func (m *RemoveTeamRequest) GetId() int64 {
	if m != nil {
		return m.Id
	}
	return 0
}

type GetTeamsRequest struct {
	Filter               *TeamFilter `protobuf:"bytes,1,opt,name=filter,proto3" json:"filter,omitempty"`
	Limit                int64       `protobuf:"varint,3,opt,name=limit,proto3" json:"limit,omitempty"`
	Offset               int64       `protobuf:"varint,4,opt,name=offset,proto3" json:"offset,omitempty"`
	Order                string      `protobuf:"bytes,5,opt,name=order,proto3" json:"order,omitempty"`
	Sort                 string      `protobuf:"bytes,6,opt,name=sort,proto3" json:"sort,omitempty"`
	XXX_NoUnkeyedLiteral struct{}    `json:"-"`
	XXX_unrecognized     []byte      `json:"-"`
	XXX_sizecache        int32       `json:"-"`
}

func (m *GetTeamsRequest) Reset()         { *m = GetTeamsRequest{} }
func (m *GetTeamsRequest) String() string { return proto.CompactTextString(m) }
func (*GetTeamsRequest) ProtoMessage()    {}
func (*GetTeamsRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_f63f4a1b2b4dddb4, []int{4}
}

func (m *GetTeamsRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_GetTeamsRequest.Unmarshal(m, b)
}
func (m *GetTeamsRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_GetTeamsRequest.Marshal(b, m, deterministic)
}
func (m *GetTeamsRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_GetTeamsRequest.Merge(m, src)
}
func (m *GetTeamsRequest) XXX_Size() int {
	return xxx_messageInfo_GetTeamsRequest.Size(m)
}
func (m *GetTeamsRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_GetTeamsRequest.DiscardUnknown(m)
}

var xxx_messageInfo_GetTeamsRequest proto.InternalMessageInfo

func (m *GetTeamsRequest) GetFilter() *TeamFilter {
	if m != nil {
		return m.Filter
	}
	return nil
}

func (m *GetTeamsRequest) GetLimit() int64 {
	if m != nil {
		return m.Limit
	}
	return 0
}

func (m *GetTeamsRequest) GetOffset() int64 {
	if m != nil {
		return m.Offset
	}
	return 0
}

func (m *GetTeamsRequest) GetOrder() string {
	if m != nil {
		return m.Order
	}
	return ""
}

func (m *GetTeamsRequest) GetSort() string {
	if m != nil {
		return m.Sort
	}
	return ""
}

type TeamFilter struct {
	Ids                  []int64              `protobuf:"varint,1,rep,packed,name=ids,proto3" json:"ids,omitempty"`
	DateFrom             *timestamp.Timestamp `protobuf:"bytes,2,opt,name=DateFrom,proto3" json:"DateFrom,omitempty"`
	DateTo               *timestamp.Timestamp `protobuf:"bytes,3,opt,name=DateTo,proto3" json:"DateTo,omitempty"`
	XXX_NoUnkeyedLiteral struct{}             `json:"-"`
	XXX_unrecognized     []byte               `json:"-"`
	XXX_sizecache        int32                `json:"-"`
}

func (m *TeamFilter) Reset()         { *m = TeamFilter{} }
func (m *TeamFilter) String() string { return proto.CompactTextString(m) }
func (*TeamFilter) ProtoMessage()    {}
func (*TeamFilter) Descriptor() ([]byte, []int) {
	return fileDescriptor_f63f4a1b2b4dddb4, []int{5}
}

func (m *TeamFilter) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_TeamFilter.Unmarshal(m, b)
}
func (m *TeamFilter) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_TeamFilter.Marshal(b, m, deterministic)
}
func (m *TeamFilter) XXX_Merge(src proto.Message) {
	xxx_messageInfo_TeamFilter.Merge(m, src)
}
func (m *TeamFilter) XXX_Size() int {
	return xxx_messageInfo_TeamFilter.Size(m)
}
func (m *TeamFilter) XXX_DiscardUnknown() {
	xxx_messageInfo_TeamFilter.DiscardUnknown(m)
}

var xxx_messageInfo_TeamFilter proto.InternalMessageInfo

func (m *TeamFilter) GetIds() []int64 {
	if m != nil {
		return m.Ids
	}
	return nil
}

func (m *TeamFilter) GetDateFrom() *timestamp.Timestamp {
	if m != nil {
		return m.DateFrom
	}
	return nil
}

func (m *TeamFilter) GetDateTo() *timestamp.Timestamp {
	if m != nil {
		return m.DateTo
	}
	return nil
}

type GetTeamsReply struct {
	Teams                []*Team  `protobuf:"bytes,1,rep,name=teams,proto3" json:"teams,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *GetTeamsReply) Reset()         { *m = GetTeamsReply{} }
func (m *GetTeamsReply) String() string { return proto.CompactTextString(m) }
func (*GetTeamsReply) ProtoMessage()    {}
func (*GetTeamsReply) Descriptor() ([]byte, []int) {
	return fileDescriptor_f63f4a1b2b4dddb4, []int{6}
}

func (m *GetTeamsReply) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_GetTeamsReply.Unmarshal(m, b)
}
func (m *GetTeamsReply) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_GetTeamsReply.Marshal(b, m, deterministic)
}
func (m *GetTeamsReply) XXX_Merge(src proto.Message) {
	xxx_messageInfo_GetTeamsReply.Merge(m, src)
}
func (m *GetTeamsReply) XXX_Size() int {
	return xxx_messageInfo_GetTeamsReply.Size(m)
}
func (m *GetTeamsReply) XXX_DiscardUnknown() {
	xxx_messageInfo_GetTeamsReply.DiscardUnknown(m)
}

var xxx_messageInfo_GetTeamsReply proto.InternalMessageInfo

func (m *GetTeamsReply) GetTeams() []*Team {
	if m != nil {
		return m.Teams
	}
	return nil
}

type Team struct {
	Id                   int64                `protobuf:"varint,1,opt,name=id,proto3" json:"id,omitempty"`
	Name                 string               `protobuf:"bytes,2,opt,name=name,proto3" json:"name,omitempty"`
	Description          string               `protobuf:"bytes,3,opt,name=description,proto3" json:"description,omitempty"`
	Slack                string               `protobuf:"bytes,4,opt,name=slack,proto3" json:"slack,omitempty"`
	Created              *timestamp.Timestamp `protobuf:"bytes,12,opt,name=created,proto3" json:"created,omitempty"`
	Updated              *timestamp.Timestamp `protobuf:"bytes,13,opt,name=updated,proto3" json:"updated,omitempty"`
	XXX_NoUnkeyedLiteral struct{}             `json:"-"`
	XXX_unrecognized     []byte               `json:"-"`
	XXX_sizecache        int32                `json:"-"`
}

func (m *Team) Reset()         { *m = Team{} }
func (m *Team) String() string { return proto.CompactTextString(m) }
func (*Team) ProtoMessage()    {}
func (*Team) Descriptor() ([]byte, []int) {
	return fileDescriptor_f63f4a1b2b4dddb4, []int{7}
}

func (m *Team) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Team.Unmarshal(m, b)
}
func (m *Team) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Team.Marshal(b, m, deterministic)
}
func (m *Team) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Team.Merge(m, src)
}
func (m *Team) XXX_Size() int {
	return xxx_messageInfo_Team.Size(m)
}
func (m *Team) XXX_DiscardUnknown() {
	xxx_messageInfo_Team.DiscardUnknown(m)
}

var xxx_messageInfo_Team proto.InternalMessageInfo

func (m *Team) GetId() int64 {
	if m != nil {
		return m.Id
	}
	return 0
}

func (m *Team) GetName() string {
	if m != nil {
		return m.Name
	}
	return ""
}

func (m *Team) GetDescription() string {
	if m != nil {
		return m.Description
	}
	return ""
}

func (m *Team) GetSlack() string {
	if m != nil {
		return m.Slack
	}
	return ""
}

func (m *Team) GetCreated() *timestamp.Timestamp {
	if m != nil {
		return m.Created
	}
	return nil
}

func (m *Team) GetUpdated() *timestamp.Timestamp {
	if m != nil {
		return m.Updated
	}
	return nil
}

func init() {
	proto.RegisterType((*AddTeamRequest)(nil), "api.AddTeamRequest")
	proto.RegisterType((*AddTeamReply)(nil), "api.AddTeamReply")
	proto.RegisterType((*UpdateTeamRequest)(nil), "api.UpdateTeamRequest")
	proto.RegisterType((*RemoveTeamRequest)(nil), "api.RemoveTeamRequest")
	proto.RegisterType((*GetTeamsRequest)(nil), "api.GetTeamsRequest")
	proto.RegisterType((*TeamFilter)(nil), "api.TeamFilter")
	proto.RegisterType((*GetTeamsReply)(nil), "api.GetTeamsReply")
	proto.RegisterType((*Team)(nil), "api.Team")
}

func init() { proto.RegisterFile("teams.proto", fileDescriptor_f63f4a1b2b4dddb4) }

var fileDescriptor_f63f4a1b2b4dddb4 = []byte{
	// 505 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xb4, 0x53, 0xcd, 0x6e, 0xd3, 0x40,
	0x10, 0xc6, 0x71, 0x92, 0xb6, 0x93, 0xfe, 0x90, 0xa5, 0x8a, 0x2c, 0x23, 0xd1, 0xc8, 0x20, 0xc8,
	0xc9, 0x86, 0x14, 0xf5, 0x88, 0x04, 0x82, 0x72, 0xb7, 0xc2, 0x05, 0x71, 0xd9, 0x66, 0x27, 0xd1,
	0x0a, 0xdb, 0xbb, 0xec, 0x6e, 0x2a, 0xe5, 0x15, 0x38, 0xf2, 0x60, 0x3c, 0x0d, 0x0f, 0x80, 0x76,
	0xec, 0x24, 0x6d, 0x22, 0xa5, 0x17, 0xb8, 0xcd, 0xcf, 0x37, 0x3b, 0xdf, 0xcc, 0x7c, 0x0b, 0x3d,
	0x87, 0xbc, 0xb4, 0xa9, 0x36, 0xca, 0x29, 0x16, 0x72, 0x2d, 0xe3, 0xa7, 0x73, 0xa5, 0xe6, 0x05,
	0x66, 0x14, 0xba, 0x59, 0xcc, 0x32, 0x2c, 0xb5, 0x5b, 0xd6, 0x88, 0xf8, 0x62, 0x3b, 0xe9, 0x64,
	0x89, 0xd6, 0xf1, 0x52, 0xd7, 0x80, 0xe4, 0x1b, 0x9c, 0xbe, 0x17, 0x62, 0x82, 0xbc, 0xcc, 0xf1,
	0xc7, 0x02, 0xad, 0x63, 0x0c, 0xda, 0x15, 0x2f, 0x31, 0x6a, 0x0d, 0x83, 0xd1, 0x51, 0x4e, 0x36,
	0x1b, 0x42, 0x4f, 0xa0, 0x9d, 0x1a, 0xa9, 0x9d, 0x54, 0x55, 0x14, 0x52, 0xea, 0x6e, 0x88, 0x9d,
	0x43, 0xc7, 0x16, 0x7c, 0xfa, 0x3d, 0x6a, 0x53, 0xae, 0x76, 0x92, 0x67, 0x70, 0xbc, 0x7e, 0x5d,
	0x17, 0x4b, 0x76, 0x0a, 0x2d, 0x29, 0xa2, 0x60, 0x18, 0x8c, 0xc2, 0xbc, 0x25, 0x45, 0xa2, 0xa0,
	0xff, 0x45, 0x0b, 0xee, 0xf0, 0x2e, 0x81, 0x2d, 0xd0, 0x3f, 0x25, 0xf4, 0x1c, 0xfa, 0x39, 0x96,
	0xea, 0x76, 0x5f, 0xc3, 0xe4, 0x57, 0x00, 0x67, 0x9f, 0xd1, 0x79, 0x88, 0x5d, 0x61, 0x5e, 0x41,
	0x77, 0x26, 0x0b, 0x87, 0x86, 0x70, 0xbd, 0xf1, 0x59, 0xca, 0xb5, 0x4c, 0x3d, 0xe4, 0x9a, 0xc2,
	0x79, 0x93, 0xf6, 0x7d, 0x0b, 0x59, 0x4a, 0x47, 0x9c, 0xc2, 0xbc, 0x76, 0xd8, 0x00, 0xba, 0x6a,
	0x36, 0xb3, 0xe8, 0x88, 0x4e, 0x98, 0x37, 0x9e, 0x47, 0x2b, 0x23, 0xd0, 0x44, 0x9d, 0x9a, 0x25,
	0x39, 0x7e, 0x62, 0xab, 0x8c, 0x8b, 0xba, 0xf5, 0xc4, 0xde, 0x4e, 0x7e, 0x06, 0x00, 0x9b, 0x76,
	0xec, 0x31, 0x84, 0x52, 0xd8, 0x28, 0x18, 0x86, 0xa3, 0x30, 0xf7, 0x26, 0xbb, 0x82, 0xc3, 0x8f,
	0xdc, 0xe1, 0xb5, 0x51, 0x25, 0xad, 0xaa, 0x37, 0x8e, 0xd3, 0xfa, 0xfa, 0xe9, 0xea, 0xfa, 0xe9,
	0x64, 0x75, 0xfd, 0x7c, 0x8d, 0x65, 0x63, 0xe8, 0x7a, 0x7b, 0xa2, 0x88, 0xf1, 0xfe, 0xaa, 0x06,
	0x99, 0xbc, 0x86, 0x93, 0xcd, 0x82, 0xfc, 0x61, 0x2f, 0xa0, 0x43, 0xc2, 0x24, 0x42, 0xbd, 0xf1,
	0xd1, 0x7a, 0x3b, 0x79, 0x1d, 0x4f, 0x7e, 0x07, 0xd0, 0xf6, 0xfe, 0xff, 0xbc, 0x2e, 0x7b, 0x0b,
	0x07, 0x53, 0x83, 0xdc, 0xa1, 0x88, 0x8e, 0x1f, 0x9c, 0x65, 0x05, 0xf5, 0x55, 0x0b, 0x12, 0xa1,
	0x88, 0x4e, 0x1e, 0xae, 0x6a, 0xa0, 0xe3, 0x3f, 0x01, 0x74, 0x68, 0x01, 0xec, 0x12, 0x0e, 0x1a,
	0x91, 0xb3, 0x27, 0x34, 0xf7, 0xfd, 0x0f, 0x15, 0xf7, 0xef, 0x07, 0x75, 0xb1, 0x4c, 0x1e, 0xb1,
	0x77, 0x00, 0x1b, 0xe5, 0xb3, 0x01, 0x41, 0x76, 0xbe, 0x42, 0x3c, 0xd8, 0x61, 0xf2, 0xc9, 0x7f,
	0xee, 0xba, 0x7e, 0x23, 0xe4, 0xa6, 0x7e, 0x47, 0xd9, 0x7b, 0xea, 0xaf, 0xe0, 0x70, 0x75, 0x41,
	0x76, 0x4e, 0xd5, 0x5b, 0x8a, 0x8f, 0xd9, 0x56, 0x94, 0x78, 0x7f, 0x78, 0xf9, 0xf5, 0x85, 0xac,
	0x1c, 0x9a, 0x8a, 0x17, 0xd9, 0x1c, 0x2b, 0x34, 0x7e, 0x19, 0xd9, 0xb4, 0x90, 0x58, 0x39, 0x9b,
	0x69, 0x54, 0xba, 0xc0, 0xec, 0xf6, 0xcd, 0x4d, 0x97, 0x3a, 0x5e, 0xfe, 0x0d, 0x00, 0x00, 0xff,
	0xff, 0x42, 0xb4, 0x26, 0x3e, 0xb0, 0x04, 0x00, 0x00,
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConnInterface

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion6

// TeamsClient is the client API for Teams service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type TeamsClient interface {
	AddTeam(ctx context.Context, in *AddTeamRequest, opts ...grpc.CallOption) (*AddTeamReply, error)
	UpdateTeam(ctx context.Context, in *UpdateTeamRequest, opts ...grpc.CallOption) (*empty.Empty, error)
	RemoveTeam(ctx context.Context, in *RemoveTeamRequest, opts ...grpc.CallOption) (*empty.Empty, error)
	GetTeams(ctx context.Context, in *GetTeamsRequest, opts ...grpc.CallOption) (*GetTeamsReply, error)
}

type teamsClient struct {
	cc grpc.ClientConnInterface
}

func NewTeamsClient(cc grpc.ClientConnInterface) TeamsClient {
	return &teamsClient{cc}
}

func (c *teamsClient) AddTeam(ctx context.Context, in *AddTeamRequest, opts ...grpc.CallOption) (*AddTeamReply, error) {
	out := new(AddTeamReply)
	err := c.cc.Invoke(ctx, "/api.Teams/AddTeam", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *teamsClient) UpdateTeam(ctx context.Context, in *UpdateTeamRequest, opts ...grpc.CallOption) (*empty.Empty, error) {
	out := new(empty.Empty)
	err := c.cc.Invoke(ctx, "/api.Teams/UpdateTeam", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *teamsClient) RemoveTeam(ctx context.Context, in *RemoveTeamRequest, opts ...grpc.CallOption) (*empty.Empty, error) {
	out := new(empty.Empty)
	err := c.cc.Invoke(ctx, "/api.Teams/RemoveTeam", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *teamsClient) GetTeams(ctx context.Context, in *GetTeamsRequest, opts ...grpc.CallOption) (*GetTeamsReply, error) {
	out := new(GetTeamsReply)
	err := c.cc.Invoke(ctx, "/api.Teams/GetTeams", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// TeamsServer is the server API for Teams service.
type TeamsServer interface {
	AddTeam(context.Context, *AddTeamRequest) (*AddTeamReply, error)
	UpdateTeam(context.Context, *UpdateTeamRequest) (*empty.Empty, error)
	RemoveTeam(context.Context, *RemoveTeamRequest) (*empty.Empty, error)
	GetTeams(context.Context, *GetTeamsRequest) (*GetTeamsReply, error)
}

// UnimplementedTeamsServer can be embedded to have forward compatible implementations.
type UnimplementedTeamsServer struct {
}

func (*UnimplementedTeamsServer) AddTeam(ctx context.Context, req *AddTeamRequest) (*AddTeamReply, error) {
	return nil, status.Errorf(codes.Unimplemented, "method AddTeam not implemented")
}
func (*UnimplementedTeamsServer) UpdateTeam(ctx context.Context, req *UpdateTeamRequest) (*empty.Empty, error) {
	return nil, status.Errorf(codes.Unimplemented, "method UpdateTeam not implemented")
}
func (*UnimplementedTeamsServer) RemoveTeam(ctx context.Context, req *RemoveTeamRequest) (*empty.Empty, error) {
	return nil, status.Errorf(codes.Unimplemented, "method RemoveTeam not implemented")
}
func (*UnimplementedTeamsServer) GetTeams(ctx context.Context, req *GetTeamsRequest) (*GetTeamsReply, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetTeams not implemented")
}

func RegisterTeamsServer(s *grpc.Server, srv TeamsServer) {
	s.RegisterService(&_Teams_serviceDesc, srv)
}

func _Teams_AddTeam_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(AddTeamRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(TeamsServer).AddTeam(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/api.Teams/AddTeam",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(TeamsServer).AddTeam(ctx, req.(*AddTeamRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Teams_UpdateTeam_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(UpdateTeamRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(TeamsServer).UpdateTeam(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/api.Teams/UpdateTeam",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(TeamsServer).UpdateTeam(ctx, req.(*UpdateTeamRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Teams_RemoveTeam_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(RemoveTeamRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(TeamsServer).RemoveTeam(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/api.Teams/RemoveTeam",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(TeamsServer).RemoveTeam(ctx, req.(*RemoveTeamRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Teams_GetTeams_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetTeamsRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(TeamsServer).GetTeams(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/api.Teams/GetTeams",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(TeamsServer).GetTeams(ctx, req.(*GetTeamsRequest))
	}
	return interceptor(ctx, in, info, handler)
}

var _Teams_serviceDesc = grpc.ServiceDesc{
	ServiceName: "api.Teams",
	HandlerType: (*TeamsServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "AddTeam",
			Handler:    _Teams_AddTeam_Handler,
		},
		{
			MethodName: "UpdateTeam",
			Handler:    _Teams_UpdateTeam_Handler,
		},
		{
			MethodName: "RemoveTeam",
			Handler:    _Teams_RemoveTeam_Handler,
		},
		{
			MethodName: "GetTeams",
			Handler:    _Teams_GetTeams_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "teams.proto",
}
