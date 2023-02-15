// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: bertydirectory.proto

package directorytypes

import (
	_ "berty.tech/berty/v2/go/pkg/protocoltypes"
	context "context"
	fmt "fmt"
	_ "github.com/gogo/protobuf/gogoproto"
	proto "github.com/gogo/protobuf/proto"
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
const _ = proto.GoGoProtoPackageIsVersion3 // please upgrade the proto package

type Register struct {
}

func (m *Register) Reset()         { *m = Register{} }
func (m *Register) String() string { return proto.CompactTextString(m) }
func (*Register) ProtoMessage()    {}
func (*Register) Descriptor() ([]byte, []int) {
	return fileDescriptor_dd73011b15fb0caf, []int{0}
}
func (m *Register) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Register.Unmarshal(m, b)
}
func (m *Register) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Register.Marshal(b, m, deterministic)
}
func (m *Register) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Register.Merge(m, src)
}
func (m *Register) XXX_Size() int {
	return xxx_messageInfo_Register.Size(m)
}
func (m *Register) XXX_DiscardUnknown() {
	xxx_messageInfo_Register.DiscardUnknown(m)
}

var xxx_messageInfo_Register proto.InternalMessageInfo

type Register_Request struct {
	VerifiedCredential      []byte `protobuf:"bytes,1,opt,name=verified_credential,json=verifiedCredential,proto3" json:"verified_credential,omitempty"`
	ExpirationDate          int64  `protobuf:"varint,2,opt,name=expiration_date,json=expirationDate,proto3" json:"expiration_date,omitempty"`
	LockedUntilDate         int64  `protobuf:"varint,3,opt,name=locked_until_date,json=lockedUntilDate,proto3" json:"locked_until_date,omitempty"`
	AccountURI              string `protobuf:"bytes,4,opt,name=account_uri,json=accountUri,proto3" json:"account_uri,omitempty"`
	OverwriteExistingRecord bool   `protobuf:"varint,5,opt,name=overwrite_existing_record,json=overwriteExistingRecord,proto3" json:"overwrite_existing_record,omitempty"`
}

func (m *Register_Request) Reset()         { *m = Register_Request{} }
func (m *Register_Request) String() string { return proto.CompactTextString(m) }
func (*Register_Request) ProtoMessage()    {}
func (*Register_Request) Descriptor() ([]byte, []int) {
	return fileDescriptor_dd73011b15fb0caf, []int{0, 0}
}
func (m *Register_Request) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Register_Request.Unmarshal(m, b)
}
func (m *Register_Request) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Register_Request.Marshal(b, m, deterministic)
}
func (m *Register_Request) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Register_Request.Merge(m, src)
}
func (m *Register_Request) XXX_Size() int {
	return xxx_messageInfo_Register_Request.Size(m)
}
func (m *Register_Request) XXX_DiscardUnknown() {
	xxx_messageInfo_Register_Request.DiscardUnknown(m)
}

var xxx_messageInfo_Register_Request proto.InternalMessageInfo

func (m *Register_Request) GetVerifiedCredential() []byte {
	if m != nil {
		return m.VerifiedCredential
	}
	return nil
}

func (m *Register_Request) GetExpirationDate() int64 {
	if m != nil {
		return m.ExpirationDate
	}
	return 0
}

func (m *Register_Request) GetLockedUntilDate() int64 {
	if m != nil {
		return m.LockedUntilDate
	}
	return 0
}

func (m *Register_Request) GetAccountURI() string {
	if m != nil {
		return m.AccountURI
	}
	return ""
}

func (m *Register_Request) GetOverwriteExistingRecord() bool {
	if m != nil {
		return m.OverwriteExistingRecord
	}
	return false
}

type Register_Reply struct {
	DirectoryRecordToken string `protobuf:"bytes,1,opt,name=directory_record_token,json=directoryRecordToken,proto3" json:"directory_record_token,omitempty"`
	DirectoryIdentifier  string `protobuf:"bytes,2,opt,name=directory_identifier,json=directoryIdentifier,proto3" json:"directory_identifier,omitempty"`
	ExpirationDate       int64  `protobuf:"varint,3,opt,name=expiration_date,json=expirationDate,proto3" json:"expiration_date,omitempty"`
	UnregisterToken      string `protobuf:"bytes,4,opt,name=unregister_token,json=unregisterToken,proto3" json:"unregister_token,omitempty"`
}

func (m *Register_Reply) Reset()         { *m = Register_Reply{} }
func (m *Register_Reply) String() string { return proto.CompactTextString(m) }
func (*Register_Reply) ProtoMessage()    {}
func (*Register_Reply) Descriptor() ([]byte, []int) {
	return fileDescriptor_dd73011b15fb0caf, []int{0, 1}
}
func (m *Register_Reply) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Register_Reply.Unmarshal(m, b)
}
func (m *Register_Reply) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Register_Reply.Marshal(b, m, deterministic)
}
func (m *Register_Reply) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Register_Reply.Merge(m, src)
}
func (m *Register_Reply) XXX_Size() int {
	return xxx_messageInfo_Register_Reply.Size(m)
}
func (m *Register_Reply) XXX_DiscardUnknown() {
	xxx_messageInfo_Register_Reply.DiscardUnknown(m)
}

var xxx_messageInfo_Register_Reply proto.InternalMessageInfo

func (m *Register_Reply) GetDirectoryRecordToken() string {
	if m != nil {
		return m.DirectoryRecordToken
	}
	return ""
}

func (m *Register_Reply) GetDirectoryIdentifier() string {
	if m != nil {
		return m.DirectoryIdentifier
	}
	return ""
}

func (m *Register_Reply) GetExpirationDate() int64 {
	if m != nil {
		return m.ExpirationDate
	}
	return 0
}

func (m *Register_Reply) GetUnregisterToken() string {
	if m != nil {
		return m.UnregisterToken
	}
	return ""
}

type Query struct {
}

func (m *Query) Reset()         { *m = Query{} }
func (m *Query) String() string { return proto.CompactTextString(m) }
func (*Query) ProtoMessage()    {}
func (*Query) Descriptor() ([]byte, []int) {
	return fileDescriptor_dd73011b15fb0caf, []int{1}
}
func (m *Query) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Query.Unmarshal(m, b)
}
func (m *Query) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Query.Marshal(b, m, deterministic)
}
func (m *Query) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Query.Merge(m, src)
}
func (m *Query) XXX_Size() int {
	return xxx_messageInfo_Query.Size(m)
}
func (m *Query) XXX_DiscardUnknown() {
	xxx_messageInfo_Query.DiscardUnknown(m)
}

var xxx_messageInfo_Query proto.InternalMessageInfo

type Query_Request struct {
	DirectoryIdentifiers []string `protobuf:"bytes,1,rep,name=directory_identifiers,json=directoryIdentifiers,proto3" json:"directory_identifiers,omitempty"`
}

func (m *Query_Request) Reset()         { *m = Query_Request{} }
func (m *Query_Request) String() string { return proto.CompactTextString(m) }
func (*Query_Request) ProtoMessage()    {}
func (*Query_Request) Descriptor() ([]byte, []int) {
	return fileDescriptor_dd73011b15fb0caf, []int{1, 0}
}
func (m *Query_Request) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Query_Request.Unmarshal(m, b)
}
func (m *Query_Request) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Query_Request.Marshal(b, m, deterministic)
}
func (m *Query_Request) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Query_Request.Merge(m, src)
}
func (m *Query_Request) XXX_Size() int {
	return xxx_messageInfo_Query_Request.Size(m)
}
func (m *Query_Request) XXX_DiscardUnknown() {
	xxx_messageInfo_Query_Request.DiscardUnknown(m)
}

var xxx_messageInfo_Query_Request proto.InternalMessageInfo

func (m *Query_Request) GetDirectoryIdentifiers() []string {
	if m != nil {
		return m.DirectoryIdentifiers
	}
	return nil
}

type Query_Reply struct {
	DirectoryIdentifier string `protobuf:"bytes,1,opt,name=directory_identifier,json=directoryIdentifier,proto3" json:"directory_identifier,omitempty"`
	ExpiresAt           int64  `protobuf:"varint,2,opt,name=expires_at,json=expiresAt,proto3" json:"expires_at,omitempty"`
	AccountURI          string `protobuf:"bytes,3,opt,name=account_uri,json=accountUri,proto3" json:"account_uri,omitempty"`
	VerifiedCredential  []byte `protobuf:"bytes,4,opt,name=verified_credential,json=verifiedCredential,proto3" json:"verified_credential,omitempty"`
}

func (m *Query_Reply) Reset()         { *m = Query_Reply{} }
func (m *Query_Reply) String() string { return proto.CompactTextString(m) }
func (*Query_Reply) ProtoMessage()    {}
func (*Query_Reply) Descriptor() ([]byte, []int) {
	return fileDescriptor_dd73011b15fb0caf, []int{1, 1}
}
func (m *Query_Reply) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Query_Reply.Unmarshal(m, b)
}
func (m *Query_Reply) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Query_Reply.Marshal(b, m, deterministic)
}
func (m *Query_Reply) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Query_Reply.Merge(m, src)
}
func (m *Query_Reply) XXX_Size() int {
	return xxx_messageInfo_Query_Reply.Size(m)
}
func (m *Query_Reply) XXX_DiscardUnknown() {
	xxx_messageInfo_Query_Reply.DiscardUnknown(m)
}

var xxx_messageInfo_Query_Reply proto.InternalMessageInfo

func (m *Query_Reply) GetDirectoryIdentifier() string {
	if m != nil {
		return m.DirectoryIdentifier
	}
	return ""
}

func (m *Query_Reply) GetExpiresAt() int64 {
	if m != nil {
		return m.ExpiresAt
	}
	return 0
}

func (m *Query_Reply) GetAccountURI() string {
	if m != nil {
		return m.AccountURI
	}
	return ""
}

func (m *Query_Reply) GetVerifiedCredential() []byte {
	if m != nil {
		return m.VerifiedCredential
	}
	return nil
}

type Unregister struct {
}

func (m *Unregister) Reset()         { *m = Unregister{} }
func (m *Unregister) String() string { return proto.CompactTextString(m) }
func (*Unregister) ProtoMessage()    {}
func (*Unregister) Descriptor() ([]byte, []int) {
	return fileDescriptor_dd73011b15fb0caf, []int{2}
}
func (m *Unregister) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Unregister.Unmarshal(m, b)
}
func (m *Unregister) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Unregister.Marshal(b, m, deterministic)
}
func (m *Unregister) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Unregister.Merge(m, src)
}
func (m *Unregister) XXX_Size() int {
	return xxx_messageInfo_Unregister.Size(m)
}
func (m *Unregister) XXX_DiscardUnknown() {
	xxx_messageInfo_Unregister.DiscardUnknown(m)
}

var xxx_messageInfo_Unregister proto.InternalMessageInfo

type Unregister_Request struct {
	DirectoryIdentifier  string `protobuf:"bytes,1,opt,name=directory_identifier,json=directoryIdentifier,proto3" json:"directory_identifier,omitempty"`
	DirectoryRecordToken string `protobuf:"bytes,2,opt,name=directory_record_token,json=directoryRecordToken,proto3" json:"directory_record_token,omitempty"`
	UnregisterToken      string `protobuf:"bytes,3,opt,name=unregister_token,json=unregisterToken,proto3" json:"unregister_token,omitempty"`
}

func (m *Unregister_Request) Reset()         { *m = Unregister_Request{} }
func (m *Unregister_Request) String() string { return proto.CompactTextString(m) }
func (*Unregister_Request) ProtoMessage()    {}
func (*Unregister_Request) Descriptor() ([]byte, []int) {
	return fileDescriptor_dd73011b15fb0caf, []int{2, 0}
}
func (m *Unregister_Request) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Unregister_Request.Unmarshal(m, b)
}
func (m *Unregister_Request) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Unregister_Request.Marshal(b, m, deterministic)
}
func (m *Unregister_Request) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Unregister_Request.Merge(m, src)
}
func (m *Unregister_Request) XXX_Size() int {
	return xxx_messageInfo_Unregister_Request.Size(m)
}
func (m *Unregister_Request) XXX_DiscardUnknown() {
	xxx_messageInfo_Unregister_Request.DiscardUnknown(m)
}

var xxx_messageInfo_Unregister_Request proto.InternalMessageInfo

func (m *Unregister_Request) GetDirectoryIdentifier() string {
	if m != nil {
		return m.DirectoryIdentifier
	}
	return ""
}

func (m *Unregister_Request) GetDirectoryRecordToken() string {
	if m != nil {
		return m.DirectoryRecordToken
	}
	return ""
}

func (m *Unregister_Request) GetUnregisterToken() string {
	if m != nil {
		return m.UnregisterToken
	}
	return ""
}

type Unregister_Reply struct {
}

func (m *Unregister_Reply) Reset()         { *m = Unregister_Reply{} }
func (m *Unregister_Reply) String() string { return proto.CompactTextString(m) }
func (*Unregister_Reply) ProtoMessage()    {}
func (*Unregister_Reply) Descriptor() ([]byte, []int) {
	return fileDescriptor_dd73011b15fb0caf, []int{2, 1}
}
func (m *Unregister_Reply) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Unregister_Reply.Unmarshal(m, b)
}
func (m *Unregister_Reply) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Unregister_Reply.Marshal(b, m, deterministic)
}
func (m *Unregister_Reply) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Unregister_Reply.Merge(m, src)
}
func (m *Unregister_Reply) XXX_Size() int {
	return xxx_messageInfo_Unregister_Reply.Size(m)
}
func (m *Unregister_Reply) XXX_DiscardUnknown() {
	xxx_messageInfo_Unregister_Reply.DiscardUnknown(m)
}

var xxx_messageInfo_Unregister_Reply proto.InternalMessageInfo

type Record struct {
	DirectoryIdentifier  string `protobuf:"bytes,1,opt,name=directory_identifier,json=directoryIdentifier,proto3" json:"directory_identifier,omitempty" gorm:"index;primaryKey;autoIncrement:false"`
	DirectoryRecordToken string `protobuf:"bytes,2,opt,name=directory_record_token,json=directoryRecordToken,proto3" json:"directory_record_token,omitempty"`
	ExpiresAt            int64  `protobuf:"varint,3,opt,name=expires_at,json=expiresAt,proto3" json:"expires_at,omitempty"`
	LockedUntil          int64  `protobuf:"varint,4,opt,name=locked_until,json=lockedUntil,proto3" json:"locked_until,omitempty"`
	UnregisterToken      string `protobuf:"bytes,5,opt,name=unregister_token,json=unregisterToken,proto3" json:"unregister_token,omitempty"`
	AccountURI           string `protobuf:"bytes,6,opt,name=account_uri,json=accountUri,proto3" json:"account_uri,omitempty"`
	VerifiedCredential   []byte `protobuf:"bytes,7,opt,name=verified_credential,json=verifiedCredential,proto3" json:"verified_credential,omitempty"`
}

func (m *Record) Reset()         { *m = Record{} }
func (m *Record) String() string { return proto.CompactTextString(m) }
func (*Record) ProtoMessage()    {}
func (*Record) Descriptor() ([]byte, []int) {
	return fileDescriptor_dd73011b15fb0caf, []int{3}
}
func (m *Record) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Record.Unmarshal(m, b)
}
func (m *Record) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Record.Marshal(b, m, deterministic)
}
func (m *Record) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Record.Merge(m, src)
}
func (m *Record) XXX_Size() int {
	return xxx_messageInfo_Record.Size(m)
}
func (m *Record) XXX_DiscardUnknown() {
	xxx_messageInfo_Record.DiscardUnknown(m)
}

var xxx_messageInfo_Record proto.InternalMessageInfo

func (m *Record) GetDirectoryIdentifier() string {
	if m != nil {
		return m.DirectoryIdentifier
	}
	return ""
}

func (m *Record) GetDirectoryRecordToken() string {
	if m != nil {
		return m.DirectoryRecordToken
	}
	return ""
}

func (m *Record) GetExpiresAt() int64 {
	if m != nil {
		return m.ExpiresAt
	}
	return 0
}

func (m *Record) GetLockedUntil() int64 {
	if m != nil {
		return m.LockedUntil
	}
	return 0
}

func (m *Record) GetUnregisterToken() string {
	if m != nil {
		return m.UnregisterToken
	}
	return ""
}

func (m *Record) GetAccountURI() string {
	if m != nil {
		return m.AccountURI
	}
	return ""
}

func (m *Record) GetVerifiedCredential() []byte {
	if m != nil {
		return m.VerifiedCredential
	}
	return nil
}

func init() {
	proto.RegisterType((*Register)(nil), "berty.directory.v1.Register")
	proto.RegisterType((*Register_Request)(nil), "berty.directory.v1.Register.Request")
	proto.RegisterType((*Register_Reply)(nil), "berty.directory.v1.Register.Reply")
	proto.RegisterType((*Query)(nil), "berty.directory.v1.Query")
	proto.RegisterType((*Query_Request)(nil), "berty.directory.v1.Query.Request")
	proto.RegisterType((*Query_Reply)(nil), "berty.directory.v1.Query.Reply")
	proto.RegisterType((*Unregister)(nil), "berty.directory.v1.Unregister")
	proto.RegisterType((*Unregister_Request)(nil), "berty.directory.v1.Unregister.Request")
	proto.RegisterType((*Unregister_Reply)(nil), "berty.directory.v1.Unregister.Reply")
	proto.RegisterType((*Record)(nil), "berty.directory.v1.Record")
}

func init() { proto.RegisterFile("bertydirectory.proto", fileDescriptor_dd73011b15fb0caf) }

var fileDescriptor_dd73011b15fb0caf = []byte{
	// 657 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xac, 0x55, 0x41, 0x4f, 0xd4, 0x4e,
	0x1c, 0xa5, 0x5b, 0x16, 0xd8, 0x81, 0x00, 0xff, 0x81, 0xbf, 0xae, 0x4d, 0x84, 0xa5, 0x21, 0xb8,
	0x6a, 0xb2, 0x15, 0xe1, 0x04, 0x89, 0x09, 0x88, 0x07, 0x62, 0x3c, 0x38, 0xb2, 0x17, 0x2e, 0x4d,
	0x69, 0x7f, 0xd4, 0x09, 0xa5, 0x53, 0xa7, 0xd3, 0x75, 0xfb, 0x01, 0xbc, 0x9b, 0x78, 0xf6, 0x63,
	0x78, 0xf2, 0xea, 0x07, 0xe0, 0xe8, 0x89, 0x83, 0xfb, 0x09, 0xf4, 0xec, 0xc1, 0xec, 0xb4, 0xdb,
	0x2e, 0x6b, 0xbb, 0x9b, 0x10, 0x4f, 0x6d, 0xde, 0x7b, 0x33, 0x99, 0xdf, 0x7b, 0x6f, 0x5a, 0xb4,
	0x7a, 0x06, 0x5c, 0xc4, 0x0e, 0xe5, 0x60, 0x0b, 0xc6, 0xe3, 0x56, 0xc0, 0x99, 0x60, 0x18, 0x4b,
	0xb4, 0x95, 0xc3, 0x9d, 0x6d, 0x6d, 0xd5, 0x65, 0x2e, 0x93, 0xb4, 0xd1, 0x7f, 0x4b, 0x94, 0xda,
	0x8a, 0x7c, 0xd8, 0xcc, 0x13, 0x71, 0x00, 0x61, 0x02, 0xea, 0x57, 0x2a, 0x9a, 0x23, 0xe0, 0xd2,
	0x50, 0x00, 0xd7, 0x7e, 0x2b, 0x68, 0x96, 0xc0, 0xbb, 0x08, 0x42, 0x81, 0x0d, 0xb4, 0xd2, 0x01,
	0x4e, 0xcf, 0x29, 0x38, 0xa6, 0xcd, 0xc1, 0x01, 0x5f, 0x50, 0xcb, 0xab, 0x2b, 0x0d, 0xa5, 0xb9,
	0x40, 0xf0, 0x80, 0x7a, 0x9e, 0x31, 0xf8, 0x01, 0x5a, 0x82, 0x6e, 0x40, 0xb9, 0x25, 0x28, 0xf3,
	0x4d, 0xc7, 0x12, 0x50, 0xaf, 0x34, 0x94, 0xa6, 0x4a, 0x16, 0x73, 0xf8, 0xc8, 0x12, 0x80, 0x1f,
	0xa1, 0xff, 0x3c, 0x66, 0x5f, 0x80, 0x63, 0x46, 0xbe, 0xa0, 0x5e, 0x22, 0x55, 0xa5, 0x74, 0x29,
	0x21, 0xda, 0x7d, 0x5c, 0x6a, 0x0d, 0x34, 0x6f, 0xd9, 0x36, 0x8b, 0x7c, 0x61, 0x46, 0x9c, 0xd6,
	0xa7, 0x1b, 0x4a, 0xb3, 0x76, 0xb8, 0xf8, 0xe3, 0x7a, 0x1d, 0x1d, 0x24, 0x70, 0x9b, 0x1c, 0x13,
	0x94, 0x4a, 0xda, 0x9c, 0xe2, 0x3d, 0x74, 0x8f, 0x75, 0x80, 0xbf, 0xe7, 0x54, 0x80, 0x09, 0x5d,
	0x1a, 0x0a, 0xea, 0xbb, 0x26, 0x07, 0x9b, 0x71, 0xa7, 0x5e, 0x6d, 0x28, 0xcd, 0x39, 0x72, 0x37,
	0x13, 0xbc, 0x48, 0x79, 0x22, 0x69, 0xed, 0x9b, 0x82, 0xaa, 0x04, 0x02, 0x2f, 0xc6, 0xbb, 0xe8,
	0x4e, 0x66, 0x68, 0xba, 0xd8, 0x14, 0xec, 0x02, 0x7c, 0x39, 0x7f, 0x8d, 0xac, 0x66, 0x6c, 0xb2,
	0xf4, 0xa4, 0xcf, 0xe1, 0x6d, 0x94, 0xe3, 0x26, 0x95, 0xbe, 0x9c, 0x53, 0xe0, 0xd2, 0x86, 0x1a,
	0x59, 0xc9, 0xb8, 0xe3, 0x8c, 0x2a, 0x32, 0x4d, 0x2d, 0x34, 0xed, 0x21, 0x5a, 0x8e, 0x7c, 0x9e,
	0x06, 0x95, 0x9e, 0x45, 0xba, 0x41, 0x96, 0x72, 0x5c, 0x1e, 0x43, 0xff, 0x50, 0x41, 0xd5, 0xd7,
	0x11, 0xf0, 0x58, 0x7b, 0x96, 0xc7, 0xb9, 0x83, 0xfe, 0x2f, 0x3a, 0x5b, 0x58, 0x57, 0x1a, 0xea,
	0x8d, 0x81, 0xf2, 0xc3, 0x85, 0xda, 0xd7, 0xcc, 0x90, 0xb2, 0xd1, 0x94, 0xf2, 0xd1, 0xee, 0x23,
	0x24, 0x67, 0x80, 0xd0, 0xb4, 0x44, 0x5a, 0x85, 0x5a, 0x8a, 0x1c, 0x88, 0xd1, 0x64, 0xd5, 0x89,
	0xc9, 0x96, 0x14, 0x72, 0xba, 0xac, 0x90, 0xfa, 0x17, 0x05, 0xa1, 0x76, 0xe6, 0x8d, 0xf6, 0x79,
	0xa8, 0xdc, 0xb7, 0x18, 0xa7, 0xbc, 0x12, 0x95, 0x31, 0x95, 0x28, 0x8a, 0x4d, 0x2d, 0x8c, 0x4d,
	0x9b, 0x4d, 0xbd, 0xd6, 0x7f, 0x56, 0xd0, 0x4c, 0xb2, 0x07, 0x3e, 0x1b, 0x77, 0xce, 0x43, 0xe3,
	0xd7, 0xf5, 0xfa, 0x63, 0x97, 0xf1, 0xcb, 0x3d, 0x9d, 0xfa, 0x0e, 0x74, 0xf7, 0x03, 0x4e, 0x2f,
	0x2d, 0x1e, 0xbf, 0x84, 0x78, 0xdf, 0x8a, 0x04, 0x3b, 0xf6, 0x6d, 0x0e, 0x97, 0xe0, 0x8b, 0xbd,
	0x73, 0xcb, 0x0b, 0x41, 0xff, 0x97, 0x83, 0xdd, 0x4c, 0x57, 0x1d, 0x4d, 0x77, 0x03, 0x2d, 0x0c,
	0xdf, 0x71, 0x99, 0x92, 0x4a, 0xe6, 0x87, 0xae, 0x77, 0xa1, 0x35, 0xd5, 0x42, 0x6b, 0x46, 0xbb,
	0x32, 0x73, 0xdb, 0xae, 0xcc, 0x96, 0x75, 0xe5, 0xe9, 0xa7, 0x0a, 0x5a, 0x3e, 0x1a, 0xcc, 0xf9,
	0x06, 0x78, 0x87, 0xda, 0x80, 0x4f, 0xf2, 0x4f, 0x23, 0xde, 0x6c, 0xfd, 0xfd, 0x9d, 0x6d, 0x0d,
	0xd8, 0x56, 0xda, 0x2b, 0x4d, 0x9f, 0xa0, 0xea, 0x5f, 0xa5, 0x57, 0xe9, 0xed, 0xc4, 0x1b, 0x45,
	0x62, 0x49, 0x65, 0xfb, 0xad, 0x8f, 0x93, 0x04, 0x5e, 0xfc, 0x44, 0xc1, 0xa7, 0xc3, 0x25, 0xc7,
	0x5b, 0x45, 0x0b, 0x72, 0x3e, 0xdb, 0x78, 0x73, 0xa2, 0x2e, 0xf0, 0xe2, 0xc3, 0xdd, 0x8f, 0xbd,
	0xb5, 0xa9, 0xab, 0xde, 0xda, 0xd4, 0xf7, 0xde, 0xda, 0xd4, 0xe9, 0x56, 0xb2, 0x44, 0x80, 0xfd,
	0xd6, 0x90, 0xaf, 0x86, 0xcb, 0x8c, 0xe0, 0xc2, 0x35, 0xb2, 0x4d, 0xe4, 0x8f, 0xe5, 0x6c, 0x46,
	0xfe, 0x59, 0x76, 0xfe, 0x04, 0x00, 0x00, 0xff, 0xff, 0x39, 0x59, 0x58, 0x7e, 0xb0, 0x06, 0x00,
	0x00,
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConn

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion4

// DirectoryServiceClient is the client API for DirectoryService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type DirectoryServiceClient interface {
	Register(ctx context.Context, in *Register_Request, opts ...grpc.CallOption) (*Register_Reply, error)
	Query(ctx context.Context, in *Query_Request, opts ...grpc.CallOption) (DirectoryService_QueryClient, error)
	Unregister(ctx context.Context, in *Unregister_Request, opts ...grpc.CallOption) (*Unregister_Reply, error)
}

type directoryServiceClient struct {
	cc *grpc.ClientConn
}

func NewDirectoryServiceClient(cc *grpc.ClientConn) DirectoryServiceClient {
	return &directoryServiceClient{cc}
}

func (c *directoryServiceClient) Register(ctx context.Context, in *Register_Request, opts ...grpc.CallOption) (*Register_Reply, error) {
	out := new(Register_Reply)
	err := c.cc.Invoke(ctx, "/berty.directory.v1.DirectoryService/Register", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *directoryServiceClient) Query(ctx context.Context, in *Query_Request, opts ...grpc.CallOption) (DirectoryService_QueryClient, error) {
	stream, err := c.cc.NewStream(ctx, &_DirectoryService_serviceDesc.Streams[0], "/berty.directory.v1.DirectoryService/Query", opts...)
	if err != nil {
		return nil, err
	}
	x := &directoryServiceQueryClient{stream}
	if err := x.ClientStream.SendMsg(in); err != nil {
		return nil, err
	}
	if err := x.ClientStream.CloseSend(); err != nil {
		return nil, err
	}
	return x, nil
}

type DirectoryService_QueryClient interface {
	Recv() (*Query_Reply, error)
	grpc.ClientStream
}

type directoryServiceQueryClient struct {
	grpc.ClientStream
}

func (x *directoryServiceQueryClient) Recv() (*Query_Reply, error) {
	m := new(Query_Reply)
	if err := x.ClientStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

func (c *directoryServiceClient) Unregister(ctx context.Context, in *Unregister_Request, opts ...grpc.CallOption) (*Unregister_Reply, error) {
	out := new(Unregister_Reply)
	err := c.cc.Invoke(ctx, "/berty.directory.v1.DirectoryService/Unregister", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// DirectoryServiceServer is the server API for DirectoryService service.
type DirectoryServiceServer interface {
	Register(context.Context, *Register_Request) (*Register_Reply, error)
	Query(*Query_Request, DirectoryService_QueryServer) error
	Unregister(context.Context, *Unregister_Request) (*Unregister_Reply, error)
}

// UnimplementedDirectoryServiceServer can be embedded to have forward compatible implementations.
type UnimplementedDirectoryServiceServer struct {
}

func (*UnimplementedDirectoryServiceServer) Register(ctx context.Context, req *Register_Request) (*Register_Reply, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Register not implemented")
}
func (*UnimplementedDirectoryServiceServer) Query(req *Query_Request, srv DirectoryService_QueryServer) error {
	return status.Errorf(codes.Unimplemented, "method Query not implemented")
}
func (*UnimplementedDirectoryServiceServer) Unregister(ctx context.Context, req *Unregister_Request) (*Unregister_Reply, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Unregister not implemented")
}

func RegisterDirectoryServiceServer(s *grpc.Server, srv DirectoryServiceServer) {
	s.RegisterService(&_DirectoryService_serviceDesc, srv)
}

func _DirectoryService_Register_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(Register_Request)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(DirectoryServiceServer).Register(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/berty.directory.v1.DirectoryService/Register",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(DirectoryServiceServer).Register(ctx, req.(*Register_Request))
	}
	return interceptor(ctx, in, info, handler)
}

func _DirectoryService_Query_Handler(srv interface{}, stream grpc.ServerStream) error {
	m := new(Query_Request)
	if err := stream.RecvMsg(m); err != nil {
		return err
	}
	return srv.(DirectoryServiceServer).Query(m, &directoryServiceQueryServer{stream})
}

type DirectoryService_QueryServer interface {
	Send(*Query_Reply) error
	grpc.ServerStream
}

type directoryServiceQueryServer struct {
	grpc.ServerStream
}

func (x *directoryServiceQueryServer) Send(m *Query_Reply) error {
	return x.ServerStream.SendMsg(m)
}

func _DirectoryService_Unregister_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(Unregister_Request)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(DirectoryServiceServer).Unregister(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/berty.directory.v1.DirectoryService/Unregister",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(DirectoryServiceServer).Unregister(ctx, req.(*Unregister_Request))
	}
	return interceptor(ctx, in, info, handler)
}

var _DirectoryService_serviceDesc = grpc.ServiceDesc{
	ServiceName: "berty.directory.v1.DirectoryService",
	HandlerType: (*DirectoryServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "Register",
			Handler:    _DirectoryService_Register_Handler,
		},
		{
			MethodName: "Unregister",
			Handler:    _DirectoryService_Unregister_Handler,
		},
	},
	Streams: []grpc.StreamDesc{
		{
			StreamName:    "Query",
			Handler:       _DirectoryService_Query_Handler,
			ServerStreams: true,
		},
	},
	Metadata: "bertydirectory.proto",
}