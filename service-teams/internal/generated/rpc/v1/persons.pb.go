// Code generated by protoc-gen-go. DO NOT EDIT.
// source: persons.proto

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

type Role int32

const (
	Role_NONE      Role = 0
	Role_BACK_END  Role = 1
	Role_FRONT_END Role = 2
	Role_MOBILE    Role = 3
	Role_QA        Role = 4
)

var Role_name = map[int32]string{
	0: "NONE",
	1: "BACK_END",
	2: "FRONT_END",
	3: "MOBILE",
	4: "QA",
}

var Role_value = map[string]int32{
	"NONE":      0,
	"BACK_END":  1,
	"FRONT_END": 2,
	"MOBILE":    3,
	"QA":        4,
}

func (x Role) String() string {
	return proto.EnumName(Role_name, int32(x))
}

func (Role) EnumDescriptor() ([]byte, []int) {
	return fileDescriptor_e7a3cab4d409bfb2, []int{0}
}

type AddPersonRequest struct {
	TeamId               int64                 `protobuf:"varint,1,opt,name=team_id,json=teamId,proto3" json:"team_id,omitempty"`
	FirstName            string                `protobuf:"bytes,2,opt,name=first_name,json=firstName,proto3" json:"first_name,omitempty"`
	MiddleName           *wrappers.StringValue `protobuf:"bytes,3,opt,name=middle_name,json=middleName,proto3" json:"middle_name,omitempty"`
	LastName             string                `protobuf:"bytes,4,opt,name=last_name,json=lastName,proto3" json:"last_name,omitempty"`
	Birthday             *timestamp.Timestamp  `protobuf:"bytes,5,opt,name=birthday,proto3" json:"birthday,omitempty"`
	Email                *wrappers.StringValue `protobuf:"bytes,6,opt,name=email,proto3" json:"email,omitempty"`
	Phone                *wrappers.StringValue `protobuf:"bytes,7,opt,name=phone,proto3" json:"phone,omitempty"`
	Slack                string                `protobuf:"bytes,8,opt,name=slack,proto3" json:"slack,omitempty"`
	Role                 Role                  `protobuf:"varint,9,opt,name=role,proto3,enum=api.Role" json:"role,omitempty"`
	IsActive             bool                  `protobuf:"varint,10,opt,name=isActive,proto3" json:"isActive,omitempty"`
	XXX_NoUnkeyedLiteral struct{}              `json:"-"`
	XXX_unrecognized     []byte                `json:"-"`
	XXX_sizecache        int32                 `json:"-"`
}

func (m *AddPersonRequest) Reset()         { *m = AddPersonRequest{} }
func (m *AddPersonRequest) String() string { return proto.CompactTextString(m) }
func (*AddPersonRequest) ProtoMessage()    {}
func (*AddPersonRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_e7a3cab4d409bfb2, []int{0}
}

func (m *AddPersonRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_AddPersonRequest.Unmarshal(m, b)
}
func (m *AddPersonRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_AddPersonRequest.Marshal(b, m, deterministic)
}
func (m *AddPersonRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_AddPersonRequest.Merge(m, src)
}
func (m *AddPersonRequest) XXX_Size() int {
	return xxx_messageInfo_AddPersonRequest.Size(m)
}
func (m *AddPersonRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_AddPersonRequest.DiscardUnknown(m)
}

var xxx_messageInfo_AddPersonRequest proto.InternalMessageInfo

func (m *AddPersonRequest) GetTeamId() int64 {
	if m != nil {
		return m.TeamId
	}
	return 0
}

func (m *AddPersonRequest) GetFirstName() string {
	if m != nil {
		return m.FirstName
	}
	return ""
}

func (m *AddPersonRequest) GetMiddleName() *wrappers.StringValue {
	if m != nil {
		return m.MiddleName
	}
	return nil
}

func (m *AddPersonRequest) GetLastName() string {
	if m != nil {
		return m.LastName
	}
	return ""
}

func (m *AddPersonRequest) GetBirthday() *timestamp.Timestamp {
	if m != nil {
		return m.Birthday
	}
	return nil
}

func (m *AddPersonRequest) GetEmail() *wrappers.StringValue {
	if m != nil {
		return m.Email
	}
	return nil
}

func (m *AddPersonRequest) GetPhone() *wrappers.StringValue {
	if m != nil {
		return m.Phone
	}
	return nil
}

func (m *AddPersonRequest) GetSlack() string {
	if m != nil {
		return m.Slack
	}
	return ""
}

func (m *AddPersonRequest) GetRole() Role {
	if m != nil {
		return m.Role
	}
	return Role_NONE
}

func (m *AddPersonRequest) GetIsActive() bool {
	if m != nil {
		return m.IsActive
	}
	return false
}

type AddPersonReply struct {
	Id                   int64    `protobuf:"varint,1,opt,name=id,proto3" json:"id,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *AddPersonReply) Reset()         { *m = AddPersonReply{} }
func (m *AddPersonReply) String() string { return proto.CompactTextString(m) }
func (*AddPersonReply) ProtoMessage()    {}
func (*AddPersonReply) Descriptor() ([]byte, []int) {
	return fileDescriptor_e7a3cab4d409bfb2, []int{1}
}

func (m *AddPersonReply) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_AddPersonReply.Unmarshal(m, b)
}
func (m *AddPersonReply) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_AddPersonReply.Marshal(b, m, deterministic)
}
func (m *AddPersonReply) XXX_Merge(src proto.Message) {
	xxx_messageInfo_AddPersonReply.Merge(m, src)
}
func (m *AddPersonReply) XXX_Size() int {
	return xxx_messageInfo_AddPersonReply.Size(m)
}
func (m *AddPersonReply) XXX_DiscardUnknown() {
	xxx_messageInfo_AddPersonReply.DiscardUnknown(m)
}

var xxx_messageInfo_AddPersonReply proto.InternalMessageInfo

func (m *AddPersonReply) GetId() int64 {
	if m != nil {
		return m.Id
	}
	return 0
}

type UpdatePersonRequest struct {
	Id                   int64                 `protobuf:"varint,1,opt,name=id,proto3" json:"id,omitempty"`
	TeamId               int64                 `protobuf:"varint,2,opt,name=team_id,json=teamId,proto3" json:"team_id,omitempty"`
	FirstName            string                `protobuf:"bytes,3,opt,name=first_name,json=firstName,proto3" json:"first_name,omitempty"`
	MiddleName           *wrappers.StringValue `protobuf:"bytes,4,opt,name=middle_name,json=middleName,proto3" json:"middle_name,omitempty"`
	LastName             string                `protobuf:"bytes,5,opt,name=last_name,json=lastName,proto3" json:"last_name,omitempty"`
	Birthday             *timestamp.Timestamp  `protobuf:"bytes,6,opt,name=birthday,proto3" json:"birthday,omitempty"`
	Email                *wrappers.StringValue `protobuf:"bytes,7,opt,name=email,proto3" json:"email,omitempty"`
	Phone                *wrappers.StringValue `protobuf:"bytes,8,opt,name=phone,proto3" json:"phone,omitempty"`
	Slack                string                `protobuf:"bytes,9,opt,name=slack,proto3" json:"slack,omitempty"`
	Role                 Role                  `protobuf:"varint,10,opt,name=role,proto3,enum=api.Role" json:"role,omitempty"`
	DutyOrder            int64                 `protobuf:"varint,11,opt,name=duty_order,json=dutyOrder,proto3" json:"duty_order,omitempty"`
	IsActive             bool                  `protobuf:"varint,12,opt,name=isActive,proto3" json:"isActive,omitempty"`
	XXX_NoUnkeyedLiteral struct{}              `json:"-"`
	XXX_unrecognized     []byte                `json:"-"`
	XXX_sizecache        int32                 `json:"-"`
}

func (m *UpdatePersonRequest) Reset()         { *m = UpdatePersonRequest{} }
func (m *UpdatePersonRequest) String() string { return proto.CompactTextString(m) }
func (*UpdatePersonRequest) ProtoMessage()    {}
func (*UpdatePersonRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_e7a3cab4d409bfb2, []int{2}
}

func (m *UpdatePersonRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_UpdatePersonRequest.Unmarshal(m, b)
}
func (m *UpdatePersonRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_UpdatePersonRequest.Marshal(b, m, deterministic)
}
func (m *UpdatePersonRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_UpdatePersonRequest.Merge(m, src)
}
func (m *UpdatePersonRequest) XXX_Size() int {
	return xxx_messageInfo_UpdatePersonRequest.Size(m)
}
func (m *UpdatePersonRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_UpdatePersonRequest.DiscardUnknown(m)
}

var xxx_messageInfo_UpdatePersonRequest proto.InternalMessageInfo

func (m *UpdatePersonRequest) GetId() int64 {
	if m != nil {
		return m.Id
	}
	return 0
}

func (m *UpdatePersonRequest) GetTeamId() int64 {
	if m != nil {
		return m.TeamId
	}
	return 0
}

func (m *UpdatePersonRequest) GetFirstName() string {
	if m != nil {
		return m.FirstName
	}
	return ""
}

func (m *UpdatePersonRequest) GetMiddleName() *wrappers.StringValue {
	if m != nil {
		return m.MiddleName
	}
	return nil
}

func (m *UpdatePersonRequest) GetLastName() string {
	if m != nil {
		return m.LastName
	}
	return ""
}

func (m *UpdatePersonRequest) GetBirthday() *timestamp.Timestamp {
	if m != nil {
		return m.Birthday
	}
	return nil
}

func (m *UpdatePersonRequest) GetEmail() *wrappers.StringValue {
	if m != nil {
		return m.Email
	}
	return nil
}

func (m *UpdatePersonRequest) GetPhone() *wrappers.StringValue {
	if m != nil {
		return m.Phone
	}
	return nil
}

func (m *UpdatePersonRequest) GetSlack() string {
	if m != nil {
		return m.Slack
	}
	return ""
}

func (m *UpdatePersonRequest) GetRole() Role {
	if m != nil {
		return m.Role
	}
	return Role_NONE
}

func (m *UpdatePersonRequest) GetDutyOrder() int64 {
	if m != nil {
		return m.DutyOrder
	}
	return 0
}

func (m *UpdatePersonRequest) GetIsActive() bool {
	if m != nil {
		return m.IsActive
	}
	return false
}

type RemovePersonRequest struct {
	Id                   int64    `protobuf:"varint,1,opt,name=id,proto3" json:"id,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *RemovePersonRequest) Reset()         { *m = RemovePersonRequest{} }
func (m *RemovePersonRequest) String() string { return proto.CompactTextString(m) }
func (*RemovePersonRequest) ProtoMessage()    {}
func (*RemovePersonRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_e7a3cab4d409bfb2, []int{3}
}

func (m *RemovePersonRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_RemovePersonRequest.Unmarshal(m, b)
}
func (m *RemovePersonRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_RemovePersonRequest.Marshal(b, m, deterministic)
}
func (m *RemovePersonRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_RemovePersonRequest.Merge(m, src)
}
func (m *RemovePersonRequest) XXX_Size() int {
	return xxx_messageInfo_RemovePersonRequest.Size(m)
}
func (m *RemovePersonRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_RemovePersonRequest.DiscardUnknown(m)
}

var xxx_messageInfo_RemovePersonRequest proto.InternalMessageInfo

func (m *RemovePersonRequest) GetId() int64 {
	if m != nil {
		return m.Id
	}
	return 0
}

type GetPersonsRequest struct {
	Filter               *PersonFilter `protobuf:"bytes,1,opt,name=filter,proto3" json:"filter,omitempty"`
	Limit                int64         `protobuf:"varint,2,opt,name=limit,proto3" json:"limit,omitempty"`
	Offset               int64         `protobuf:"varint,3,opt,name=offset,proto3" json:"offset,omitempty"`
	Order                string        `protobuf:"bytes,4,opt,name=order,proto3" json:"order,omitempty"`
	Sort                 string        `protobuf:"bytes,5,opt,name=sort,proto3" json:"sort,omitempty"`
	XXX_NoUnkeyedLiteral struct{}      `json:"-"`
	XXX_unrecognized     []byte        `json:"-"`
	XXX_sizecache        int32         `json:"-"`
}

func (m *GetPersonsRequest) Reset()         { *m = GetPersonsRequest{} }
func (m *GetPersonsRequest) String() string { return proto.CompactTextString(m) }
func (*GetPersonsRequest) ProtoMessage()    {}
func (*GetPersonsRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_e7a3cab4d409bfb2, []int{4}
}

func (m *GetPersonsRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_GetPersonsRequest.Unmarshal(m, b)
}
func (m *GetPersonsRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_GetPersonsRequest.Marshal(b, m, deterministic)
}
func (m *GetPersonsRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_GetPersonsRequest.Merge(m, src)
}
func (m *GetPersonsRequest) XXX_Size() int {
	return xxx_messageInfo_GetPersonsRequest.Size(m)
}
func (m *GetPersonsRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_GetPersonsRequest.DiscardUnknown(m)
}

var xxx_messageInfo_GetPersonsRequest proto.InternalMessageInfo

func (m *GetPersonsRequest) GetFilter() *PersonFilter {
	if m != nil {
		return m.Filter
	}
	return nil
}

func (m *GetPersonsRequest) GetLimit() int64 {
	if m != nil {
		return m.Limit
	}
	return 0
}

func (m *GetPersonsRequest) GetOffset() int64 {
	if m != nil {
		return m.Offset
	}
	return 0
}

func (m *GetPersonsRequest) GetOrder() string {
	if m != nil {
		return m.Order
	}
	return ""
}

func (m *GetPersonsRequest) GetSort() string {
	if m != nil {
		return m.Sort
	}
	return ""
}

type PersonFilter struct {
	PersonIds            []int64              `protobuf:"varint,1,rep,packed,name=person_ids,json=personIds,proto3" json:"person_ids,omitempty"`
	TeamIds              []int64              `protobuf:"varint,2,rep,packed,name=team_ids,json=teamIds,proto3" json:"team_ids,omitempty"`
	DateFrom             *timestamp.Timestamp `protobuf:"bytes,3,opt,name=DateFrom,proto3" json:"DateFrom,omitempty"`
	DateTo               *timestamp.Timestamp `protobuf:"bytes,4,opt,name=DateTo,proto3" json:"DateTo,omitempty"`
	XXX_NoUnkeyedLiteral struct{}             `json:"-"`
	XXX_unrecognized     []byte               `json:"-"`
	XXX_sizecache        int32                `json:"-"`
}

func (m *PersonFilter) Reset()         { *m = PersonFilter{} }
func (m *PersonFilter) String() string { return proto.CompactTextString(m) }
func (*PersonFilter) ProtoMessage()    {}
func (*PersonFilter) Descriptor() ([]byte, []int) {
	return fileDescriptor_e7a3cab4d409bfb2, []int{5}
}

func (m *PersonFilter) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_PersonFilter.Unmarshal(m, b)
}
func (m *PersonFilter) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_PersonFilter.Marshal(b, m, deterministic)
}
func (m *PersonFilter) XXX_Merge(src proto.Message) {
	xxx_messageInfo_PersonFilter.Merge(m, src)
}
func (m *PersonFilter) XXX_Size() int {
	return xxx_messageInfo_PersonFilter.Size(m)
}
func (m *PersonFilter) XXX_DiscardUnknown() {
	xxx_messageInfo_PersonFilter.DiscardUnknown(m)
}

var xxx_messageInfo_PersonFilter proto.InternalMessageInfo

func (m *PersonFilter) GetPersonIds() []int64 {
	if m != nil {
		return m.PersonIds
	}
	return nil
}

func (m *PersonFilter) GetTeamIds() []int64 {
	if m != nil {
		return m.TeamIds
	}
	return nil
}

func (m *PersonFilter) GetDateFrom() *timestamp.Timestamp {
	if m != nil {
		return m.DateFrom
	}
	return nil
}

func (m *PersonFilter) GetDateTo() *timestamp.Timestamp {
	if m != nil {
		return m.DateTo
	}
	return nil
}

type GetPersonsReply struct {
	Persons              []*Person `protobuf:"bytes,1,rep,name=persons,proto3" json:"persons,omitempty"`
	XXX_NoUnkeyedLiteral struct{}  `json:"-"`
	XXX_unrecognized     []byte    `json:"-"`
	XXX_sizecache        int32     `json:"-"`
}

func (m *GetPersonsReply) Reset()         { *m = GetPersonsReply{} }
func (m *GetPersonsReply) String() string { return proto.CompactTextString(m) }
func (*GetPersonsReply) ProtoMessage()    {}
func (*GetPersonsReply) Descriptor() ([]byte, []int) {
	return fileDescriptor_e7a3cab4d409bfb2, []int{6}
}

func (m *GetPersonsReply) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_GetPersonsReply.Unmarshal(m, b)
}
func (m *GetPersonsReply) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_GetPersonsReply.Marshal(b, m, deterministic)
}
func (m *GetPersonsReply) XXX_Merge(src proto.Message) {
	xxx_messageInfo_GetPersonsReply.Merge(m, src)
}
func (m *GetPersonsReply) XXX_Size() int {
	return xxx_messageInfo_GetPersonsReply.Size(m)
}
func (m *GetPersonsReply) XXX_DiscardUnknown() {
	xxx_messageInfo_GetPersonsReply.DiscardUnknown(m)
}

var xxx_messageInfo_GetPersonsReply proto.InternalMessageInfo

func (m *GetPersonsReply) GetPersons() []*Person {
	if m != nil {
		return m.Persons
	}
	return nil
}

type Person struct {
	Id                   int64                 `protobuf:"varint,1,opt,name=id,proto3" json:"id,omitempty"`
	TeamId               int64                 `protobuf:"varint,2,opt,name=team_id,json=teamId,proto3" json:"team_id,omitempty"`
	FirstName            string                `protobuf:"bytes,3,opt,name=first_name,json=firstName,proto3" json:"first_name,omitempty"`
	MiddleName           *wrappers.StringValue `protobuf:"bytes,4,opt,name=middle_name,json=middleName,proto3" json:"middle_name,omitempty"`
	LastName             string                `protobuf:"bytes,5,opt,name=last_name,json=lastName,proto3" json:"last_name,omitempty"`
	Birthday             *timestamp.Timestamp  `protobuf:"bytes,6,opt,name=birthday,proto3" json:"birthday,omitempty"`
	Email                *wrappers.StringValue `protobuf:"bytes,7,opt,name=email,proto3" json:"email,omitempty"`
	Phone                *wrappers.StringValue `protobuf:"bytes,8,opt,name=phone,proto3" json:"phone,omitempty"`
	Slack                string                `protobuf:"bytes,9,opt,name=slack,proto3" json:"slack,omitempty"`
	Role                 Role                  `protobuf:"varint,10,opt,name=role,proto3,enum=api.Role" json:"role,omitempty"`
	DutyOrder            int64                 `protobuf:"varint,11,opt,name=duty_order,json=dutyOrder,proto3" json:"duty_order,omitempty"`
	IsActive             bool                  `protobuf:"varint,12,opt,name=isActive,proto3" json:"isActive,omitempty"`
	Created              *timestamp.Timestamp  `protobuf:"bytes,13,opt,name=created,proto3" json:"created,omitempty"`
	Updated              *timestamp.Timestamp  `protobuf:"bytes,14,opt,name=updated,proto3" json:"updated,omitempty"`
	XXX_NoUnkeyedLiteral struct{}              `json:"-"`
	XXX_unrecognized     []byte                `json:"-"`
	XXX_sizecache        int32                 `json:"-"`
}

func (m *Person) Reset()         { *m = Person{} }
func (m *Person) String() string { return proto.CompactTextString(m) }
func (*Person) ProtoMessage()    {}
func (*Person) Descriptor() ([]byte, []int) {
	return fileDescriptor_e7a3cab4d409bfb2, []int{7}
}

func (m *Person) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Person.Unmarshal(m, b)
}
func (m *Person) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Person.Marshal(b, m, deterministic)
}
func (m *Person) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Person.Merge(m, src)
}
func (m *Person) XXX_Size() int {
	return xxx_messageInfo_Person.Size(m)
}
func (m *Person) XXX_DiscardUnknown() {
	xxx_messageInfo_Person.DiscardUnknown(m)
}

var xxx_messageInfo_Person proto.InternalMessageInfo

func (m *Person) GetId() int64 {
	if m != nil {
		return m.Id
	}
	return 0
}

func (m *Person) GetTeamId() int64 {
	if m != nil {
		return m.TeamId
	}
	return 0
}

func (m *Person) GetFirstName() string {
	if m != nil {
		return m.FirstName
	}
	return ""
}

func (m *Person) GetMiddleName() *wrappers.StringValue {
	if m != nil {
		return m.MiddleName
	}
	return nil
}

func (m *Person) GetLastName() string {
	if m != nil {
		return m.LastName
	}
	return ""
}

func (m *Person) GetBirthday() *timestamp.Timestamp {
	if m != nil {
		return m.Birthday
	}
	return nil
}

func (m *Person) GetEmail() *wrappers.StringValue {
	if m != nil {
		return m.Email
	}
	return nil
}

func (m *Person) GetPhone() *wrappers.StringValue {
	if m != nil {
		return m.Phone
	}
	return nil
}

func (m *Person) GetSlack() string {
	if m != nil {
		return m.Slack
	}
	return ""
}

func (m *Person) GetRole() Role {
	if m != nil {
		return m.Role
	}
	return Role_NONE
}

func (m *Person) GetDutyOrder() int64 {
	if m != nil {
		return m.DutyOrder
	}
	return 0
}

func (m *Person) GetIsActive() bool {
	if m != nil {
		return m.IsActive
	}
	return false
}

func (m *Person) GetCreated() *timestamp.Timestamp {
	if m != nil {
		return m.Created
	}
	return nil
}

func (m *Person) GetUpdated() *timestamp.Timestamp {
	if m != nil {
		return m.Updated
	}
	return nil
}

func init() {
	proto.RegisterEnum("api.Role", Role_name, Role_value)
	proto.RegisterType((*AddPersonRequest)(nil), "api.AddPersonRequest")
	proto.RegisterType((*AddPersonReply)(nil), "api.AddPersonReply")
	proto.RegisterType((*UpdatePersonRequest)(nil), "api.UpdatePersonRequest")
	proto.RegisterType((*RemovePersonRequest)(nil), "api.RemovePersonRequest")
	proto.RegisterType((*GetPersonsRequest)(nil), "api.GetPersonsRequest")
	proto.RegisterType((*PersonFilter)(nil), "api.PersonFilter")
	proto.RegisterType((*GetPersonsReply)(nil), "api.GetPersonsReply")
	proto.RegisterType((*Person)(nil), "api.Person")
}

func init() { proto.RegisterFile("persons.proto", fileDescriptor_e7a3cab4d409bfb2) }

var fileDescriptor_e7a3cab4d409bfb2 = []byte{
	// 784 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xec, 0x56, 0xef, 0x6e, 0xda, 0x48,
	0x10, 0x8f, 0xff, 0x60, 0xec, 0x81, 0x70, 0x64, 0x93, 0xcb, 0x39, 0xe4, 0x72, 0x87, 0x2c, 0x45,
	0xe2, 0xee, 0x03, 0xe8, 0xb8, 0xd3, 0xe9, 0x4e, 0x6a, 0x3f, 0x40, 0x43, 0xaa, 0xa8, 0x2d, 0xb4,
	0x6e, 0xda, 0x0f, 0xfd, 0x82, 0x36, 0x78, 0x21, 0xab, 0xda, 0xd8, 0x5d, 0x2f, 0xa9, 0x78, 0x85,
	0x3e, 0x40, 0xfb, 0x0a, 0x7d, 0x80, 0xbe, 0x44, 0xdf, 0xaa, 0xda, 0x5d, 0x43, 0x09, 0x20, 0x12,
	0xa5, 0x9f, 0x2a, 0xf5, 0x9b, 0x67, 0xe6, 0x37, 0xb3, 0xb3, 0xf3, 0x9b, 0x99, 0x35, 0x6c, 0x27,
	0x84, 0xa5, 0xf1, 0x38, 0xad, 0x27, 0x2c, 0xe6, 0x31, 0x32, 0x70, 0x42, 0x2b, 0x87, 0xa3, 0x38,
	0x1e, 0x85, 0xa4, 0x21, 0x55, 0x17, 0x93, 0x61, 0x83, 0x44, 0x09, 0x9f, 0x2a, 0x44, 0xe5, 0xf7,
	0x65, 0x23, 0xa7, 0x11, 0x49, 0x39, 0x8e, 0x92, 0x0c, 0xf0, 0xdb, 0x32, 0xe0, 0x2d, 0xc3, 0x89,
	0x38, 0x45, 0xd9, 0xbd, 0x0f, 0x06, 0x94, 0x5b, 0x41, 0xf0, 0x54, 0x9e, 0xeb, 0x93, 0x37, 0x13,
	0x92, 0x72, 0xf4, 0x0b, 0xe4, 0x39, 0xc1, 0x51, 0x9f, 0x06, 0xae, 0x56, 0xd5, 0x6a, 0x86, 0x6f,
	0x09, 0xf1, 0x2c, 0x40, 0x47, 0x00, 0x43, 0xca, 0x52, 0xde, 0x1f, 0xe3, 0x88, 0xb8, 0x7a, 0x55,
	0xab, 0x39, 0xbe, 0x23, 0x35, 0x5d, 0x1c, 0x11, 0x74, 0x1f, 0x0a, 0x11, 0x0d, 0x82, 0x90, 0x28,
	0xbb, 0x51, 0xd5, 0x6a, 0x85, 0xe6, 0xaf, 0x75, 0x95, 0x42, 0x7d, 0x96, 0x42, 0xfd, 0x39, 0x67,
	0x74, 0x3c, 0x7a, 0x89, 0xc3, 0x09, 0xf1, 0x41, 0x39, 0x48, 0xf7, 0x43, 0x70, 0x42, 0x3c, 0x0b,
	0x6e, 0xca, 0xe0, 0xb6, 0x50, 0x48, 0xe3, 0xbf, 0x60, 0x5f, 0x50, 0xc6, 0x2f, 0x03, 0x3c, 0x75,
	0x73, 0x32, 0x70, 0x65, 0x25, 0xf0, 0xf9, 0xec, 0xf2, 0xfe, 0x1c, 0x8b, 0x9a, 0x90, 0x23, 0x11,
	0xa6, 0xa1, 0x6b, 0xdd, 0x22, 0x1b, 0x05, 0x15, 0x3e, 0xc9, 0x65, 0x3c, 0x26, 0x6e, 0xfe, 0x36,
	0x3e, 0x12, 0x8a, 0xf6, 0x20, 0x97, 0x86, 0x78, 0xf0, 0xda, 0xb5, 0x65, 0xe2, 0x4a, 0x40, 0x47,
	0x60, 0xb2, 0x38, 0x24, 0xae, 0x53, 0xd5, 0x6a, 0xa5, 0xa6, 0x53, 0xc7, 0x09, 0xad, 0xfb, 0x71,
	0x48, 0x7c, 0xa9, 0x46, 0x15, 0xb0, 0x69, 0xda, 0x1a, 0x70, 0x7a, 0x45, 0x5c, 0xa8, 0x6a, 0x35,
	0xdb, 0x9f, 0xcb, 0x5e, 0x15, 0x4a, 0x0b, 0xc4, 0x24, 0xe1, 0x14, 0x95, 0x40, 0x9f, 0x33, 0xa2,
	0xd3, 0xc0, 0xfb, 0x6c, 0xc0, 0xee, 0x8b, 0x24, 0xc0, 0x9c, 0x5c, 0xa7, 0x6f, 0x09, 0xb7, 0x48,
	0xa7, 0xbe, 0x81, 0x4e, 0xe3, 0x06, 0x3a, 0xcd, 0x6f, 0xa1, 0x33, 0xb7, 0x81, 0x4e, 0xeb, 0x2e,
	0x74, 0xe6, 0xef, 0x40, 0xa7, 0x7d, 0x07, 0x3a, 0x9d, 0x75, 0x74, 0xc2, 0x7a, 0x3a, 0x8f, 0x00,
	0x82, 0x09, 0x9f, 0xf6, 0x63, 0x16, 0x10, 0xe6, 0x16, 0x64, 0xad, 0x1d, 0xa1, 0xe9, 0x09, 0xc5,
	0x35, 0xb6, 0x8b, 0x4b, 0x6c, 0x1f, 0xc3, 0xae, 0x4f, 0xa2, 0xf8, 0x6a, 0x33, 0x95, 0xde, 0x7b,
	0x0d, 0x76, 0x1e, 0x12, 0xae, 0x40, 0xe9, 0x0c, 0xf5, 0x07, 0x58, 0x43, 0x1a, 0x72, 0xc2, 0x24,
	0xb2, 0xd0, 0xdc, 0x91, 0x89, 0x29, 0xd0, 0xa9, 0x34, 0xf8, 0x19, 0x40, 0xdc, 0x2b, 0xa4, 0x11,
	0xe5, 0x59, 0x27, 0x28, 0x01, 0xed, 0x83, 0x15, 0x0f, 0x87, 0x29, 0xe1, 0xb2, 0x09, 0x0c, 0x3f,
	0x93, 0x04, 0x5a, 0xdd, 0x45, 0x4d, 0xa3, 0x12, 0x10, 0x02, 0x33, 0x8d, 0x19, 0xcf, 0x38, 0x95,
	0xdf, 0xde, 0x27, 0x0d, 0x8a, 0x8b, 0x07, 0x8a, 0x5a, 0xa8, 0x65, 0xd6, 0xa7, 0x41, 0xea, 0x6a,
	0x55, 0x43, 0xd4, 0x42, 0x69, 0xce, 0x82, 0x14, 0x1d, 0x80, 0x9d, 0xf5, 0x64, 0xea, 0xea, 0xd2,
	0x98, 0x57, 0x4d, 0x99, 0x8a, 0xd6, 0x38, 0xc1, 0x9c, 0x9c, 0xb2, 0x38, 0xca, 0x56, 0xc8, 0xc6,
	0xd6, 0x98, 0x61, 0x51, 0x13, 0x2c, 0xf1, 0x7d, 0x1e, 0x67, 0x9d, 0xba, 0xc9, 0x2b, 0x43, 0x7a,
	0xff, 0xc1, 0x4f, 0x8b, 0xe5, 0x14, 0x53, 0x76, 0x0c, 0xf9, 0x6c, 0x0b, 0xcb, 0xac, 0x0b, 0xcd,
	0xc2, 0x42, 0x35, 0xfd, 0x99, 0xcd, 0xfb, 0x68, 0x82, 0xa5, 0x74, 0x3f, 0xe6, 0xed, 0xfb, 0x9b,
	0x37, 0xf4, 0x0f, 0xe4, 0x07, 0x8c, 0x60, 0x4e, 0x02, 0x77, 0xfb, 0xc6, 0x72, 0xcc, 0xa0, 0xc2,
	0x6b, 0x22, 0x17, 0x6e, 0xe0, 0x96, 0x6e, 0xf6, 0xca, 0xa0, 0x7f, 0xb6, 0xc0, 0x14, 0x49, 0x23,
	0x1b, 0xcc, 0x6e, 0xaf, 0xdb, 0x29, 0x6f, 0xa1, 0x22, 0xd8, 0xed, 0xd6, 0x83, 0x47, 0xfd, 0x4e,
	0xf7, 0xa4, 0xac, 0xa1, 0x6d, 0x70, 0x4e, 0xfd, 0x5e, 0xf7, 0x5c, 0x8a, 0x3a, 0x02, 0xb0, 0x9e,
	0xf4, 0xda, 0x67, 0x8f, 0x3b, 0x65, 0x03, 0x59, 0xa0, 0x3f, 0x6b, 0x95, 0xcd, 0xe6, 0x3b, 0x1d,
	0xf2, 0x59, 0x97, 0xa2, 0xff, 0xc1, 0x99, 0x3f, 0x0c, 0xe8, 0x67, 0x59, 0x93, 0xe5, 0x17, 0xbc,
	0xb2, 0xbb, 0xac, 0x4e, 0xc2, 0xa9, 0xb7, 0x85, 0xda, 0x50, 0x5c, 0x7c, 0x30, 0x90, 0x2b, 0x61,
	0x6b, 0xde, 0x90, 0xca, 0xfe, 0xca, 0xc5, 0x3a, 0xe2, 0xb7, 0x43, 0xc5, 0x58, 0xdc, 0x54, 0x59,
	0x8c, 0x35, 0xcb, 0x6b, 0x43, 0x8c, 0x7b, 0x00, 0x5f, 0xc7, 0x0e, 0xed, 0xcb, 0x08, 0x2b, 0x6b,
	0xad, 0xb2, 0xb7, 0xa2, 0x97, 0xb7, 0x68, 0x1f, 0xbe, 0x3a, 0xa0, 0x63, 0x4e, 0xd8, 0x18, 0x87,
	0x8d, 0x11, 0x19, 0x13, 0x26, 0xaa, 0xdc, 0x60, 0xc9, 0xa0, 0x71, 0xf5, 0xd7, 0x85, 0x25, 0x0f,
	0xfb, 0xfb, 0x4b, 0x00, 0x00, 0x00, 0xff, 0xff, 0x4d, 0xae, 0xe1, 0xb1, 0x4b, 0x09, 0x00, 0x00,
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConnInterface

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion6

// PersonsClient is the client API for Persons service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type PersonsClient interface {
	AddPerson(ctx context.Context, in *AddPersonRequest, opts ...grpc.CallOption) (*AddPersonReply, error)
	UpdatePerson(ctx context.Context, in *UpdatePersonRequest, opts ...grpc.CallOption) (*empty.Empty, error)
	RemovePerson(ctx context.Context, in *RemovePersonRequest, opts ...grpc.CallOption) (*empty.Empty, error)
	GetPersons(ctx context.Context, in *GetPersonsRequest, opts ...grpc.CallOption) (*GetPersonsReply, error)
}

type personsClient struct {
	cc grpc.ClientConnInterface
}

func NewPersonsClient(cc grpc.ClientConnInterface) PersonsClient {
	return &personsClient{cc}
}

func (c *personsClient) AddPerson(ctx context.Context, in *AddPersonRequest, opts ...grpc.CallOption) (*AddPersonReply, error) {
	out := new(AddPersonReply)
	err := c.cc.Invoke(ctx, "/api.Persons/AddPerson", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *personsClient) UpdatePerson(ctx context.Context, in *UpdatePersonRequest, opts ...grpc.CallOption) (*empty.Empty, error) {
	out := new(empty.Empty)
	err := c.cc.Invoke(ctx, "/api.Persons/UpdatePerson", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *personsClient) RemovePerson(ctx context.Context, in *RemovePersonRequest, opts ...grpc.CallOption) (*empty.Empty, error) {
	out := new(empty.Empty)
	err := c.cc.Invoke(ctx, "/api.Persons/RemovePerson", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *personsClient) GetPersons(ctx context.Context, in *GetPersonsRequest, opts ...grpc.CallOption) (*GetPersonsReply, error) {
	out := new(GetPersonsReply)
	err := c.cc.Invoke(ctx, "/api.Persons/GetPersons", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// PersonsServer is the server API for Persons service.
type PersonsServer interface {
	AddPerson(context.Context, *AddPersonRequest) (*AddPersonReply, error)
	UpdatePerson(context.Context, *UpdatePersonRequest) (*empty.Empty, error)
	RemovePerson(context.Context, *RemovePersonRequest) (*empty.Empty, error)
	GetPersons(context.Context, *GetPersonsRequest) (*GetPersonsReply, error)
}

// UnimplementedPersonsServer can be embedded to have forward compatible implementations.
type UnimplementedPersonsServer struct {
}

func (*UnimplementedPersonsServer) AddPerson(ctx context.Context, req *AddPersonRequest) (*AddPersonReply, error) {
	return nil, status.Errorf(codes.Unimplemented, "method AddPerson not implemented")
}
func (*UnimplementedPersonsServer) UpdatePerson(ctx context.Context, req *UpdatePersonRequest) (*empty.Empty, error) {
	return nil, status.Errorf(codes.Unimplemented, "method UpdatePerson not implemented")
}
func (*UnimplementedPersonsServer) RemovePerson(ctx context.Context, req *RemovePersonRequest) (*empty.Empty, error) {
	return nil, status.Errorf(codes.Unimplemented, "method RemovePerson not implemented")
}
func (*UnimplementedPersonsServer) GetPersons(ctx context.Context, req *GetPersonsRequest) (*GetPersonsReply, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetPersons not implemented")
}

func RegisterPersonsServer(s *grpc.Server, srv PersonsServer) {
	s.RegisterService(&_Persons_serviceDesc, srv)
}

func _Persons_AddPerson_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(AddPersonRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(PersonsServer).AddPerson(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/api.Persons/AddPerson",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(PersonsServer).AddPerson(ctx, req.(*AddPersonRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Persons_UpdatePerson_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(UpdatePersonRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(PersonsServer).UpdatePerson(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/api.Persons/UpdatePerson",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(PersonsServer).UpdatePerson(ctx, req.(*UpdatePersonRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Persons_RemovePerson_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(RemovePersonRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(PersonsServer).RemovePerson(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/api.Persons/RemovePerson",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(PersonsServer).RemovePerson(ctx, req.(*RemovePersonRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Persons_GetPersons_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetPersonsRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(PersonsServer).GetPersons(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/api.Persons/GetPersons",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(PersonsServer).GetPersons(ctx, req.(*GetPersonsRequest))
	}
	return interceptor(ctx, in, info, handler)
}

var _Persons_serviceDesc = grpc.ServiceDesc{
	ServiceName: "api.Persons",
	HandlerType: (*PersonsServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "AddPerson",
			Handler:    _Persons_AddPerson_Handler,
		},
		{
			MethodName: "UpdatePerson",
			Handler:    _Persons_UpdatePerson_Handler,
		},
		{
			MethodName: "RemovePerson",
			Handler:    _Persons_RemovePerson_Handler,
		},
		{
			MethodName: "GetPersons",
			Handler:    _Persons_GetPersons_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "persons.proto",
}
