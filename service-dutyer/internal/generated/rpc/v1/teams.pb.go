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

type GetTeamRequest struct {
	Id                   int64    `protobuf:"varint,1,opt,name=id,proto3" json:"id,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *GetTeamRequest) Reset()         { *m = GetTeamRequest{} }
func (m *GetTeamRequest) String() string { return proto.CompactTextString(m) }
func (*GetTeamRequest) ProtoMessage()    {}
func (*GetTeamRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_f63f4a1b2b4dddb4, []int{7}
}

func (m *GetTeamRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_GetTeamRequest.Unmarshal(m, b)
}
func (m *GetTeamRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_GetTeamRequest.Marshal(b, m, deterministic)
}
func (m *GetTeamRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_GetTeamRequest.Merge(m, src)
}
func (m *GetTeamRequest) XXX_Size() int {
	return xxx_messageInfo_GetTeamRequest.Size(m)
}
func (m *GetTeamRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_GetTeamRequest.DiscardUnknown(m)
}

var xxx_messageInfo_GetTeamRequest proto.InternalMessageInfo

func (m *GetTeamRequest) GetId() int64 {
	if m != nil {
		return m.Id
	}
	return 0
}

type GetTeamReply struct {
	Team                 *Team    `protobuf:"bytes,1,opt,name=team,proto3" json:"team,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *GetTeamReply) Reset()         { *m = GetTeamReply{} }
func (m *GetTeamReply) String() string { return proto.CompactTextString(m) }
func (*GetTeamReply) ProtoMessage()    {}
func (*GetTeamReply) Descriptor() ([]byte, []int) {
	return fileDescriptor_f63f4a1b2b4dddb4, []int{8}
}

func (m *GetTeamReply) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_GetTeamReply.Unmarshal(m, b)
}
func (m *GetTeamReply) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_GetTeamReply.Marshal(b, m, deterministic)
}
func (m *GetTeamReply) XXX_Merge(src proto.Message) {
	xxx_messageInfo_GetTeamReply.Merge(m, src)
}
func (m *GetTeamReply) XXX_Size() int {
	return xxx_messageInfo_GetTeamReply.Size(m)
}
func (m *GetTeamReply) XXX_DiscardUnknown() {
	xxx_messageInfo_GetTeamReply.DiscardUnknown(m)
}

var xxx_messageInfo_GetTeamReply proto.InternalMessageInfo

func (m *GetTeamReply) GetTeam() *Team {
	if m != nil {
		return m.Team
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
	return fileDescriptor_f63f4a1b2b4dddb4, []int{9}
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
	proto.RegisterType((*GetTeamRequest)(nil), "api.GetTeamRequest")
	proto.RegisterType((*GetTeamReply)(nil), "api.GetTeamReply")
	proto.RegisterType((*Team)(nil), "api.Team")
}

func init() { proto.RegisterFile("teams.proto", fileDescriptor_f63f4a1b2b4dddb4) }

var fileDescriptor_f63f4a1b2b4dddb4 = []byte{
	// 538 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xb4, 0x54, 0x4d, 0x6f, 0x13, 0x3d,
	0x10, 0xee, 0x66, 0xf3, 0xf1, 0x76, 0x92, 0xa6, 0x6f, 0x4c, 0x15, 0x2d, 0x5b, 0x41, 0xa3, 0xe5,
	0x40, 0x2e, 0x6c, 0x20, 0x41, 0x1c, 0x91, 0x40, 0x50, 0xee, 0xab, 0x70, 0x41, 0x5c, 0xdc, 0x78,
	0x12, 0x59, 0xec, 0xc6, 0xc6, 0xeb, 0x54, 0xca, 0x5f, 0xe0, 0xc8, 0xbf, 0xe0, 0xcf, 0xf0, 0x9b,
	0x90, 0x3f, 0x36, 0x9f, 0x6a, 0x7b, 0x81, 0x9b, 0x67, 0xfc, 0x3c, 0x9e, 0x99, 0x67, 0x66, 0x0c,
	0x6d, 0x8d, 0xb4, 0x28, 0x53, 0xa9, 0x84, 0x16, 0x24, 0xa4, 0x92, 0xc7, 0x97, 0x0b, 0x21, 0x16,
	0x39, 0x8e, 0xac, 0xeb, 0x66, 0x35, 0x1f, 0x61, 0x21, 0xf5, 0xda, 0x21, 0xe2, 0xab, 0xc3, 0x4b,
	0xcd, 0x0b, 0x2c, 0x35, 0x2d, 0xa4, 0x03, 0x24, 0x5f, 0xa1, 0xfb, 0x8e, 0xb1, 0x29, 0xd2, 0x22,
	0xc3, 0xef, 0x2b, 0x2c, 0x35, 0x21, 0x50, 0x5f, 0xd2, 0x02, 0xa3, 0xda, 0x20, 0x18, 0x9e, 0x66,
	0xf6, 0x4c, 0x06, 0xd0, 0x66, 0x58, 0xce, 0x14, 0x97, 0x9a, 0x8b, 0x65, 0x14, 0xda, 0xab, 0x5d,
	0x17, 0xb9, 0x80, 0x46, 0x99, 0xd3, 0xd9, 0xb7, 0xa8, 0x6e, 0xef, 0x9c, 0x91, 0x3c, 0x85, 0xce,
	0xe6, 0x75, 0x99, 0xaf, 0x49, 0x17, 0x6a, 0x9c, 0x45, 0xc1, 0x20, 0x18, 0x86, 0x59, 0x8d, 0xb3,
	0x44, 0x40, 0xef, 0xb3, 0x64, 0x54, 0xe3, 0x6e, 0x02, 0x07, 0xa0, 0xbf, 0x9a, 0xd0, 0x33, 0xe8,
	0x65, 0x58, 0x88, 0xdb, 0xfb, 0x02, 0x26, 0x3f, 0x03, 0x38, 0xff, 0x84, 0xda, 0x40, 0xca, 0x0a,
	0xf3, 0x1c, 0x9a, 0x73, 0x9e, 0x6b, 0x54, 0x16, 0xd7, 0x1e, 0x9f, 0xa7, 0x54, 0xf2, 0xd4, 0x40,
	0xae, 0xad, 0x3b, 0xf3, 0xd7, 0x26, 0x6e, 0xce, 0x0b, 0xae, 0x6d, 0x4e, 0x61, 0xe6, 0x0c, 0xd2,
	0x87, 0xa6, 0x98, 0xcf, 0x4b, 0xd4, 0x36, 0x9d, 0x30, 0xf3, 0x96, 0x41, 0x0b, 0xc5, 0x50, 0x45,
	0x0d, 0x97, 0xa5, 0x35, 0x4c, 0xc5, 0xa5, 0x50, 0x3a, 0x6a, 0xba, 0x8a, 0xcd, 0x39, 0xf9, 0x11,
	0x00, 0x6c, 0xc3, 0x91, 0xff, 0x21, 0xe4, 0xac, 0x8c, 0x82, 0x41, 0x38, 0x0c, 0x33, 0x73, 0x24,
	0x6f, 0xe0, 0xbf, 0x0f, 0x54, 0xe3, 0xb5, 0x12, 0x85, 0x95, 0xaa, 0x3d, 0x8e, 0x53, 0xd7, 0xfd,
	0xb4, 0xea, 0x7e, 0x3a, 0xad, 0xba, 0x9f, 0x6d, 0xb0, 0x64, 0x0c, 0x4d, 0x73, 0x9e, 0x0a, 0x9b,
	0xf1, 0xfd, 0x2c, 0x8f, 0x4c, 0x5e, 0xc2, 0xd9, 0x56, 0x20, 0xd3, 0xd8, 0x2b, 0x68, 0xd8, 0xc1,
	0xb4, 0x09, 0xb5, 0xc7, 0xa7, 0x1b, 0x75, 0x32, 0xe7, 0x4f, 0x06, 0xd0, 0xf5, 0x8c, 0xbb, 0x54,
	0x7f, 0x01, 0x9d, 0x0d, 0xc2, 0x3c, 0xf9, 0x04, 0xea, 0x86, 0xea, 0xf5, 0xde, 0x79, 0xd1, 0xba,
	0x93, 0xdf, 0x01, 0xd4, 0x8d, 0xf9, 0x2f, 0xc7, 0x85, 0xbc, 0x86, 0xd6, 0x4c, 0x21, 0xd5, 0xc8,
	0xa2, 0xce, 0x83, 0xe2, 0x54, 0x50, 0xc3, 0x5a, 0xd9, 0xa9, 0x66, 0xd1, 0xd9, 0xc3, 0x2c, 0x0f,
	0x1d, 0xff, 0xaa, 0x41, 0xc3, 0x2a, 0x4a, 0x26, 0xd0, 0xf2, 0x4a, 0x90, 0x47, 0xb6, 0xec, 0x7d,
	0xe5, 0xe2, 0xde, 0xbe, 0x53, 0xe6, 0xeb, 0xe4, 0xc4, 0xb4, 0xbf, 0x6a, 0x09, 0xb9, 0xd8, 0x05,
	0x54, 0x23, 0x1c, 0x93, 0x03, 0xaf, 0xe3, 0x4d, 0xa0, 0xe5, 0x57, 0xd4, 0x07, 0xdb, 0xff, 0x0e,
	0x7c, 0xb0, 0xdd, 0x2d, 0x4e, 0x4e, 0xc8, 0x5b, 0x80, 0xed, 0xde, 0x92, 0xbe, 0x85, 0x1c, 0x2d,
	0x72, 0xdc, 0x3f, 0x2a, 0xfb, 0xa3, 0xf9, 0x9a, 0x1c, 0x7f, 0xbb, 0x86, 0x9e, 0x7f, 0xb4, 0x97,
	0x77, 0xf3, 0xdf, 0x5f, 0x7e, 0x79, 0xcc, 0x97, 0x1a, 0xd5, 0x92, 0xe6, 0xa3, 0x05, 0x2e, 0x51,
	0x19, 0x05, 0x47, 0x4a, 0xce, 0x46, 0xb7, 0xaf, 0x6e, 0x9a, 0x16, 0x3e, 0xf9, 0x13, 0x00, 0x00,
	0xff, 0xff, 0xb9, 0x93, 0x15, 0x4a, 0x2b, 0x05, 0x00, 0x00,
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
	GetTeam(ctx context.Context, in *GetTeamRequest, opts ...grpc.CallOption) (*GetTeamReply, error)
	GetTeams(ctx context.Context, in *GetTeamsRequest, opts ...grpc.CallOption) (*GetTeamsReply, error)
	AddTeam(ctx context.Context, in *AddTeamRequest, opts ...grpc.CallOption) (*AddTeamReply, error)
	UpdateTeam(ctx context.Context, in *UpdateTeamRequest, opts ...grpc.CallOption) (*empty.Empty, error)
	RemoveTeam(ctx context.Context, in *RemoveTeamRequest, opts ...grpc.CallOption) (*empty.Empty, error)
}

type teamsClient struct {
	cc grpc.ClientConnInterface
}

func NewTeamsClient(cc grpc.ClientConnInterface) TeamsClient {
	return &teamsClient{cc}
}

func (c *teamsClient) GetTeam(ctx context.Context, in *GetTeamRequest, opts ...grpc.CallOption) (*GetTeamReply, error) {
	out := new(GetTeamReply)
	err := c.cc.Invoke(ctx, "/api.Teams/GetTeam", in, out, opts...)
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

// TeamsServer is the server API for Teams service.
type TeamsServer interface {
	GetTeam(context.Context, *GetTeamRequest) (*GetTeamReply, error)
	GetTeams(context.Context, *GetTeamsRequest) (*GetTeamsReply, error)
	AddTeam(context.Context, *AddTeamRequest) (*AddTeamReply, error)
	UpdateTeam(context.Context, *UpdateTeamRequest) (*empty.Empty, error)
	RemoveTeam(context.Context, *RemoveTeamRequest) (*empty.Empty, error)
}

// UnimplementedTeamsServer can be embedded to have forward compatible implementations.
type UnimplementedTeamsServer struct {
}

func (*UnimplementedTeamsServer) GetTeam(ctx context.Context, req *GetTeamRequest) (*GetTeamReply, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetTeam not implemented")
}
func (*UnimplementedTeamsServer) GetTeams(ctx context.Context, req *GetTeamsRequest) (*GetTeamsReply, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetTeams not implemented")
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

func RegisterTeamsServer(s *grpc.Server, srv TeamsServer) {
	s.RegisterService(&_Teams_serviceDesc, srv)
}

func _Teams_GetTeam_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetTeamRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(TeamsServer).GetTeam(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/api.Teams/GetTeam",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(TeamsServer).GetTeam(ctx, req.(*GetTeamRequest))
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

var _Teams_serviceDesc = grpc.ServiceDesc{
	ServiceName: "api.Teams",
	HandlerType: (*TeamsServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "GetTeam",
			Handler:    _Teams_GetTeam_Handler,
		},
		{
			MethodName: "GetTeams",
			Handler:    _Teams_GetTeams_Handler,
		},
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
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "teams.proto",
}