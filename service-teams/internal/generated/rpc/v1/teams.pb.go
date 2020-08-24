// Code generated by protoc-gen-go. DO NOT EDIT.
// source: teams.proto

package v1

import (
	context "context"
	fmt "fmt"
	proto "github.com/golang/protobuf/proto"
	empty "github.com/golang/protobuf/ptypes/empty"
	timestamp "github.com/golang/protobuf/ptypes/timestamp"
	wrappers "github.com/golang/protobuf/ptypes/wrappers"
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
	Created              *timestamp.Timestamp `protobuf:"bytes,5,opt,name=created,proto3" json:"created,omitempty"`
	Updated              *timestamp.Timestamp `protobuf:"bytes,6,opt,name=updated,proto3" json:"updated,omitempty"`
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
	return fileDescriptor_f63f4a1b2b4dddb4, []int{8}
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
	Id                   int64                `protobuf:"varint,1,opt,name=id,proto3" json:"id,omitempty"`
	Name                 string               `protobuf:"bytes,2,opt,name=name,proto3" json:"name,omitempty"`
	Description          string               `protobuf:"bytes,3,opt,name=description,proto3" json:"description,omitempty"`
	Slack                string               `protobuf:"bytes,4,opt,name=slack,proto3" json:"slack,omitempty"`
	Created              *timestamp.Timestamp `protobuf:"bytes,5,opt,name=created,proto3" json:"created,omitempty"`
	Updated              *timestamp.Timestamp `protobuf:"bytes,6,opt,name=updated,proto3" json:"updated,omitempty"`
	Members              []*Member            `protobuf:"bytes,7,rep,name=members,proto3" json:"members,omitempty"`
	XXX_NoUnkeyedLiteral struct{}             `json:"-"`
	XXX_unrecognized     []byte               `json:"-"`
	XXX_sizecache        int32                `json:"-"`
}

func (m *GetTeamReply) Reset()         { *m = GetTeamReply{} }
func (m *GetTeamReply) String() string { return proto.CompactTextString(m) }
func (*GetTeamReply) ProtoMessage()    {}
func (*GetTeamReply) Descriptor() ([]byte, []int) {
	return fileDescriptor_f63f4a1b2b4dddb4, []int{9}
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

func (m *GetTeamReply) GetId() int64 {
	if m != nil {
		return m.Id
	}
	return 0
}

func (m *GetTeamReply) GetName() string {
	if m != nil {
		return m.Name
	}
	return ""
}

func (m *GetTeamReply) GetDescription() string {
	if m != nil {
		return m.Description
	}
	return ""
}

func (m *GetTeamReply) GetSlack() string {
	if m != nil {
		return m.Slack
	}
	return ""
}

func (m *GetTeamReply) GetCreated() *timestamp.Timestamp {
	if m != nil {
		return m.Created
	}
	return nil
}

func (m *GetTeamReply) GetUpdated() *timestamp.Timestamp {
	if m != nil {
		return m.Updated
	}
	return nil
}

func (m *GetTeamReply) GetMembers() []*Member {
	if m != nil {
		return m.Members
	}
	return nil
}

type Member struct {
	Id                   int64                 `protobuf:"varint,1,opt,name=id,proto3" json:"id,omitempty"`
	FirstName            string                `protobuf:"bytes,2,opt,name=first_name,json=firstName,proto3" json:"first_name,omitempty"`
	MiddleName           *wrappers.StringValue `protobuf:"bytes,3,opt,name=middle_name,json=middleName,proto3" json:"middle_name,omitempty"`
	LastName             string                `protobuf:"bytes,4,opt,name=last_name,json=lastName,proto3" json:"last_name,omitempty"`
	Birthday             *timestamp.Timestamp  `protobuf:"bytes,5,opt,name=birthday,proto3" json:"birthday,omitempty"`
	Email                *wrappers.StringValue `protobuf:"bytes,6,opt,name=email,proto3" json:"email,omitempty"`
	Phone                *wrappers.StringValue `protobuf:"bytes,7,opt,name=phone,proto3" json:"phone,omitempty"`
	Slack                string                `protobuf:"bytes,8,opt,name=slack,proto3" json:"slack,omitempty"`
	Role                 int64                 `protobuf:"varint,9,opt,name=role,proto3" json:"role,omitempty"`
	DutyOrder            int64                 `protobuf:"varint,10,opt,name=duty_order,json=dutyOrder,proto3" json:"duty_order,omitempty"`
	IsActive             bool                  `protobuf:"varint,11,opt,name=isActive,proto3" json:"isActive,omitempty"`
	XXX_NoUnkeyedLiteral struct{}              `json:"-"`
	XXX_unrecognized     []byte                `json:"-"`
	XXX_sizecache        int32                 `json:"-"`
}

func (m *Member) Reset()         { *m = Member{} }
func (m *Member) String() string { return proto.CompactTextString(m) }
func (*Member) ProtoMessage()    {}
func (*Member) Descriptor() ([]byte, []int) {
	return fileDescriptor_f63f4a1b2b4dddb4, []int{10}
}

func (m *Member) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Member.Unmarshal(m, b)
}
func (m *Member) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Member.Marshal(b, m, deterministic)
}
func (m *Member) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Member.Merge(m, src)
}
func (m *Member) XXX_Size() int {
	return xxx_messageInfo_Member.Size(m)
}
func (m *Member) XXX_DiscardUnknown() {
	xxx_messageInfo_Member.DiscardUnknown(m)
}

var xxx_messageInfo_Member proto.InternalMessageInfo

func (m *Member) GetId() int64 {
	if m != nil {
		return m.Id
	}
	return 0
}

func (m *Member) GetFirstName() string {
	if m != nil {
		return m.FirstName
	}
	return ""
}

func (m *Member) GetMiddleName() *wrappers.StringValue {
	if m != nil {
		return m.MiddleName
	}
	return nil
}

func (m *Member) GetLastName() string {
	if m != nil {
		return m.LastName
	}
	return ""
}

func (m *Member) GetBirthday() *timestamp.Timestamp {
	if m != nil {
		return m.Birthday
	}
	return nil
}

func (m *Member) GetEmail() *wrappers.StringValue {
	if m != nil {
		return m.Email
	}
	return nil
}

func (m *Member) GetPhone() *wrappers.StringValue {
	if m != nil {
		return m.Phone
	}
	return nil
}

func (m *Member) GetSlack() string {
	if m != nil {
		return m.Slack
	}
	return ""
}

func (m *Member) GetRole() int64 {
	if m != nil {
		return m.Role
	}
	return 0
}

func (m *Member) GetDutyOrder() int64 {
	if m != nil {
		return m.DutyOrder
	}
	return 0
}

func (m *Member) GetIsActive() bool {
	if m != nil {
		return m.IsActive
	}
	return false
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
	proto.RegisterType((*GetTeamRequest)(nil), "api.GetTeamRequest")
	proto.RegisterType((*GetTeamReply)(nil), "api.GetTeamReply")
	proto.RegisterType((*Member)(nil), "api.Member")
}

func init() { proto.RegisterFile("teams.proto", fileDescriptor_f63f4a1b2b4dddb4) }

var fileDescriptor_f63f4a1b2b4dddb4 = []byte{
	// 719 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xd4, 0x55, 0xcf, 0x6e, 0xd3, 0x4c,
	0x10, 0xaf, 0xe3, 0xfc, 0x1d, 0xf7, 0x6b, 0xbf, 0xee, 0x57, 0x55, 0xfe, 0x52, 0x68, 0x23, 0x23,
	0x44, 0x4e, 0x09, 0xa4, 0xa8, 0x37, 0x90, 0x8a, 0xa0, 0x9c, 0x00, 0xc9, 0x14, 0x0e, 0x08, 0xa9,
	0xda, 0xc4, 0x93, 0x74, 0x85, 0x9d, 0x35, 0xeb, 0x4d, 0x51, 0x5e, 0x81, 0x23, 0x27, 0x5e, 0x81,
	0x97, 0xe1, 0x8d, 0x90, 0xd0, 0xce, 0xda, 0x49, 0x9a, 0x40, 0xe8, 0x01, 0x0e, 0xdc, 0x76, 0x7e,
	0xf3, 0x1b, 0xef, 0xcc, 0x6f, 0x66, 0xc7, 0xe0, 0x69, 0xe4, 0x49, 0xd6, 0x49, 0x95, 0xd4, 0x92,
	0xb9, 0x3c, 0x15, 0xcd, 0xfd, 0x91, 0x94, 0xa3, 0x18, 0xbb, 0x04, 0xf5, 0x27, 0xc3, 0x2e, 0x26,
	0xa9, 0x9e, 0x5a, 0x46, 0xf3, 0x70, 0xd9, 0xa9, 0x45, 0x82, 0x99, 0xe6, 0x49, 0x9a, 0x13, 0x0e,
	0x96, 0x09, 0x1f, 0x14, 0x4f, 0x53, 0x54, 0xf9, 0x15, 0xc1, 0x5b, 0xd8, 0x3a, 0x89, 0xa2, 0x33,
	0xe4, 0x49, 0x88, 0xef, 0x27, 0x98, 0x69, 0xc6, 0xa0, 0x3c, 0xe6, 0x09, 0xfa, 0xa5, 0x96, 0xd3,
	0x6e, 0x84, 0x74, 0x66, 0x2d, 0xf0, 0x22, 0xcc, 0x06, 0x4a, 0xa4, 0x5a, 0xc8, 0xb1, 0xef, 0x92,
	0x6b, 0x11, 0x62, 0xbb, 0x50, 0xc9, 0x62, 0x3e, 0x78, 0xe7, 0x97, 0xc9, 0x67, 0x8d, 0xe0, 0x00,
	0x36, 0x67, 0x5f, 0x4f, 0xe3, 0x29, 0xdb, 0x82, 0x92, 0x88, 0x7c, 0xa7, 0xe5, 0xb4, 0xdd, 0xb0,
	0x24, 0xa2, 0x40, 0xc2, 0xce, 0xab, 0x34, 0xe2, 0x1a, 0x17, 0x13, 0x58, 0x22, 0xfd, 0xd6, 0x84,
	0x6e, 0xc1, 0x4e, 0x88, 0x89, 0xbc, 0x5c, 0x77, 0x61, 0xf0, 0xc9, 0x81, 0xed, 0xa7, 0xa8, 0x0d,
	0x25, 0x2b, 0x38, 0x77, 0xa0, 0x3a, 0x14, 0xb1, 0x46, 0x45, 0x3c, 0xaf, 0xb7, 0xdd, 0xe1, 0xa9,
	0xe8, 0x18, 0xca, 0x29, 0xc1, 0x61, 0xee, 0x36, 0xf7, 0xc6, 0x22, 0x11, 0x9a, 0x72, 0x72, 0x43,
	0x6b, 0xb0, 0x3d, 0xa8, 0xca, 0xe1, 0x30, 0x43, 0x4d, 0xe9, 0xb8, 0x61, 0x6e, 0x19, 0xb6, 0x54,
	0x11, 0x2a, 0xbf, 0x62, 0xb3, 0x24, 0xc3, 0x54, 0x9c, 0x49, 0xa5, 0xfd, 0xaa, 0xad, 0xd8, 0x9c,
	0x83, 0x8f, 0x0e, 0xc0, 0xfc, 0x3a, 0xf6, 0x2f, 0xb8, 0x22, 0xca, 0x7c, 0xa7, 0xe5, 0xb6, 0xdd,
	0xd0, 0x1c, 0xd9, 0x31, 0xd4, 0x1f, 0x73, 0x8d, 0xa7, 0x4a, 0x26, 0x24, 0x95, 0xd7, 0x6b, 0x76,
	0x6c, 0xf3, 0x3b, 0x45, 0xf3, 0x3b, 0x67, 0xc5, 0x74, 0x84, 0x33, 0x2e, 0xeb, 0x41, 0xd5, 0x9c,
	0xcf, 0x24, 0x65, 0xbc, 0x3e, 0x2a, 0x67, 0x06, 0x77, 0xe1, 0x9f, 0xb9, 0x40, 0xa6, 0xb1, 0x87,
	0x50, 0xa1, 0xc1, 0xa5, 0x84, 0xbc, 0x5e, 0x63, 0xa6, 0x4e, 0x68, 0xf1, 0xe0, 0xab, 0x03, 0x65,
	0x63, 0xff, 0xc9, 0xee, 0xb2, 0xfb, 0x50, 0x1b, 0x28, 0xe4, 0x1a, 0x23, 0xd2, 0x73, 0x7d, 0x2d,
	0x05, 0xd5, 0x44, 0x4d, 0x68, 0x08, 0x23, 0x12, 0xfc, 0x17, 0x51, 0x39, 0x35, 0x68, 0xc1, 0x56,
	0x2e, 0xc1, 0xcf, 0xc6, 0xe8, 0x9b, 0x03, 0x9b, 0x33, 0xca, 0x0f, 0xa6, 0xff, 0x6f, 0x2d, 0x9d,
	0xdd, 0x86, 0x5a, 0x82, 0x49, 0x1f, 0x55, 0xe6, 0xd7, 0xa8, 0xdd, 0x1e, 0xb5, 0xfb, 0x19, 0x61,
	0x61, 0xe1, 0x0b, 0x3e, 0xbb, 0x50, 0xb5, 0xd8, 0x4a, 0xe5, 0x37, 0x01, 0x86, 0x42, 0x65, 0xfa,
	0x7c, 0xa1, 0xfe, 0x06, 0x21, 0xcf, 0x8d, 0x08, 0x0f, 0xc0, 0x4b, 0x44, 0x14, 0xc5, 0x68, 0xfd,
	0x76, 0x2e, 0x6f, 0xac, 0xa4, 0xf6, 0x52, 0x2b, 0x31, 0x1e, 0xbd, 0xe6, 0xf1, 0x04, 0x43, 0xb0,
	0x01, 0x14, 0xbe, 0x0f, 0x8d, 0x98, 0x17, 0x1f, 0xb7, 0x2a, 0xd5, 0x0d, 0x40, 0xce, 0x63, 0xa8,
	0xf7, 0x85, 0xd2, 0x17, 0x11, 0x9f, 0x5e, 0x43, 0xa9, 0x19, 0x97, 0xf5, 0xa0, 0x82, 0x09, 0x17,
	0x71, 0x2e, 0xd4, 0xfa, 0x6c, 0x2c, 0xd5, 0xc4, 0xa4, 0x17, 0x72, 0x8c, 0x7e, 0xed, 0x3a, 0x31,
	0x44, 0x9d, 0xb7, 0xb7, 0xbe, 0xd8, 0x5e, 0x06, 0x65, 0x25, 0x63, 0xf4, 0x1b, 0x24, 0x21, 0x9d,
	0x8d, 0x88, 0xd1, 0x44, 0x4f, 0xcf, 0xed, 0x02, 0x01, 0xf2, 0x34, 0x0c, 0xf2, 0x82, 0x96, 0x48,
	0x13, 0xea, 0x22, 0x3b, 0x19, 0x68, 0x71, 0x89, 0xbe, 0xd7, 0x72, 0xda, 0xf5, 0x70, 0x66, 0xf7,
	0xbe, 0x94, 0xa0, 0x42, 0xaf, 0x97, 0x1d, 0x41, 0x2d, 0xdf, 0xd0, 0xec, 0x3f, 0xea, 0xe2, 0xd5,
	0xbf, 0x41, 0x73, 0xe7, 0x2a, 0x98, 0xc6, 0xd3, 0x60, 0x83, 0x3d, 0x04, 0x98, 0xaf, 0x6d, 0xb6,
	0x47, 0x94, 0x95, 0x3d, 0xde, 0xdc, 0x5b, 0x29, 0xf7, 0x89, 0xf9, 0x73, 0xd9, 0xf8, 0xf9, 0x16,
	0xce, 0xe3, 0x57, 0xd6, 0xf2, 0x9a, 0xf8, 0x63, 0xa8, 0x17, 0xeb, 0x87, 0xed, 0x52, 0xf4, 0xd2,
	0xba, 0x6e, 0xb2, 0x25, 0xd4, 0xe6, 0x7d, 0x04, 0xb5, 0x1c, 0xca, 0x8b, 0xbd, 0xfa, 0x82, 0xf3,
	0x62, 0x17, 0xdf, 0x6c, 0xb0, 0xf1, 0x68, 0xff, 0xcd, 0xff, 0x62, 0xac, 0x51, 0x8d, 0x79, 0xdc,
	0x1d, 0xe1, 0x18, 0x95, 0x79, 0x03, 0x5d, 0x95, 0x0e, 0xba, 0x97, 0xf7, 0xfa, 0x55, 0xca, 0xed,
	0xe8, 0x7b, 0x00, 0x00, 0x00, 0xff, 0xff, 0x30, 0x17, 0xbf, 0x54, 0xb7, 0x07, 0x00, 0x00,
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
	GetTeam(ctx context.Context, in *GetTeamRequest, opts ...grpc.CallOption) (*GetTeamReply, error)
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

func (c *teamsClient) GetTeam(ctx context.Context, in *GetTeamRequest, opts ...grpc.CallOption) (*GetTeamReply, error) {
	out := new(GetTeamReply)
	err := c.cc.Invoke(ctx, "/api.Teams/GetTeam", in, out, opts...)
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
	GetTeam(context.Context, *GetTeamRequest) (*GetTeamReply, error)
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
func (*UnimplementedTeamsServer) GetTeam(ctx context.Context, req *GetTeamRequest) (*GetTeamReply, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetTeam not implemented")
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
		{
			MethodName: "GetTeam",
			Handler:    _Teams_GetTeam_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "teams.proto",
}
