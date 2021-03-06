// Code generated by protoc-gen-go. DO NOT EDIT.
// source: runtime/service/proto/runtime.proto

package go_micro_runtime

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

type Service struct {
	// name of the service
	Name string `protobuf:"bytes,1,opt,name=name,proto3" json:"name,omitempty"`
	// version of the service
	Version string `protobuf:"bytes,2,opt,name=version,proto3" json:"version,omitempty"`
	// git url of the source
	Source string `protobuf:"bytes,3,opt,name=source,proto3" json:"source,omitempty"`
	// service metadata
	Metadata             map[string]string `protobuf:"bytes,4,rep,name=metadata,proto3" json:"metadata,omitempty" protobuf_key:"bytes,1,opt,name=key,proto3" protobuf_val:"bytes,2,opt,name=value,proto3"`
	XXX_NoUnkeyedLiteral struct{}          `json:"-"`
	XXX_unrecognized     []byte            `json:"-"`
	XXX_sizecache        int32             `json:"-"`
}

func (m *Service) Reset()         { *m = Service{} }
func (m *Service) String() string { return proto.CompactTextString(m) }
func (*Service) ProtoMessage()    {}
func (*Service) Descriptor() ([]byte, []int) {
	return fileDescriptor_2434d8152598889b, []int{0}
}

func (m *Service) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Service.Unmarshal(m, b)
}
func (m *Service) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Service.Marshal(b, m, deterministic)
}
func (m *Service) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Service.Merge(m, src)
}
func (m *Service) XXX_Size() int {
	return xxx_messageInfo_Service.Size(m)
}
func (m *Service) XXX_DiscardUnknown() {
	xxx_messageInfo_Service.DiscardUnknown(m)
}

var xxx_messageInfo_Service proto.InternalMessageInfo

func (m *Service) GetName() string {
	if m != nil {
		return m.Name
	}
	return ""
}

func (m *Service) GetVersion() string {
	if m != nil {
		return m.Version
	}
	return ""
}

func (m *Service) GetSource() string {
	if m != nil {
		return m.Source
	}
	return ""
}

func (m *Service) GetMetadata() map[string]string {
	if m != nil {
		return m.Metadata
	}
	return nil
}

type Event struct {
	Type                 string   `protobuf:"bytes,1,opt,name=type,proto3" json:"type,omitempty"`
	Timestamp            int64    `protobuf:"varint,2,opt,name=timestamp,proto3" json:"timestamp,omitempty"`
	Service              string   `protobuf:"bytes,3,opt,name=service,proto3" json:"service,omitempty"`
	Version              string   `protobuf:"bytes,4,opt,name=version,proto3" json:"version,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *Event) Reset()         { *m = Event{} }
func (m *Event) String() string { return proto.CompactTextString(m) }
func (*Event) ProtoMessage()    {}
func (*Event) Descriptor() ([]byte, []int) {
	return fileDescriptor_2434d8152598889b, []int{1}
}

func (m *Event) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Event.Unmarshal(m, b)
}
func (m *Event) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Event.Marshal(b, m, deterministic)
}
func (m *Event) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Event.Merge(m, src)
}
func (m *Event) XXX_Size() int {
	return xxx_messageInfo_Event.Size(m)
}
func (m *Event) XXX_DiscardUnknown() {
	xxx_messageInfo_Event.DiscardUnknown(m)
}

var xxx_messageInfo_Event proto.InternalMessageInfo

func (m *Event) GetType() string {
	if m != nil {
		return m.Type
	}
	return ""
}

func (m *Event) GetTimestamp() int64 {
	if m != nil {
		return m.Timestamp
	}
	return 0
}

func (m *Event) GetService() string {
	if m != nil {
		return m.Service
	}
	return ""
}

func (m *Event) GetVersion() string {
	if m != nil {
		return m.Version
	}
	return ""
}

type CreateOptions struct {
	// command to pass in
	Command []string `protobuf:"bytes,1,rep,name=command,proto3" json:"command,omitempty"`
	// environment to pass in
	Env []string `protobuf:"bytes,2,rep,name=env,proto3" json:"env,omitempty"`
	// output to send to
	Output               string   `protobuf:"bytes,3,opt,name=output,proto3" json:"output,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *CreateOptions) Reset()         { *m = CreateOptions{} }
func (m *CreateOptions) String() string { return proto.CompactTextString(m) }
func (*CreateOptions) ProtoMessage()    {}
func (*CreateOptions) Descriptor() ([]byte, []int) {
	return fileDescriptor_2434d8152598889b, []int{2}
}

func (m *CreateOptions) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_CreateOptions.Unmarshal(m, b)
}
func (m *CreateOptions) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_CreateOptions.Marshal(b, m, deterministic)
}
func (m *CreateOptions) XXX_Merge(src proto.Message) {
	xxx_messageInfo_CreateOptions.Merge(m, src)
}
func (m *CreateOptions) XXX_Size() int {
	return xxx_messageInfo_CreateOptions.Size(m)
}
func (m *CreateOptions) XXX_DiscardUnknown() {
	xxx_messageInfo_CreateOptions.DiscardUnknown(m)
}

var xxx_messageInfo_CreateOptions proto.InternalMessageInfo

func (m *CreateOptions) GetCommand() []string {
	if m != nil {
		return m.Command
	}
	return nil
}

func (m *CreateOptions) GetEnv() []string {
	if m != nil {
		return m.Env
	}
	return nil
}

func (m *CreateOptions) GetOutput() string {
	if m != nil {
		return m.Output
	}
	return ""
}

type CreateRequest struct {
	Service              *Service       `protobuf:"bytes,1,opt,name=service,proto3" json:"service,omitempty"`
	Options              *CreateOptions `protobuf:"bytes,2,opt,name=options,proto3" json:"options,omitempty"`
	XXX_NoUnkeyedLiteral struct{}       `json:"-"`
	XXX_unrecognized     []byte         `json:"-"`
	XXX_sizecache        int32          `json:"-"`
}

func (m *CreateRequest) Reset()         { *m = CreateRequest{} }
func (m *CreateRequest) String() string { return proto.CompactTextString(m) }
func (*CreateRequest) ProtoMessage()    {}
func (*CreateRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_2434d8152598889b, []int{3}
}

func (m *CreateRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_CreateRequest.Unmarshal(m, b)
}
func (m *CreateRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_CreateRequest.Marshal(b, m, deterministic)
}
func (m *CreateRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_CreateRequest.Merge(m, src)
}
func (m *CreateRequest) XXX_Size() int {
	return xxx_messageInfo_CreateRequest.Size(m)
}
func (m *CreateRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_CreateRequest.DiscardUnknown(m)
}

var xxx_messageInfo_CreateRequest proto.InternalMessageInfo

func (m *CreateRequest) GetService() *Service {
	if m != nil {
		return m.Service
	}
	return nil
}

func (m *CreateRequest) GetOptions() *CreateOptions {
	if m != nil {
		return m.Options
	}
	return nil
}

type CreateResponse struct {
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *CreateResponse) Reset()         { *m = CreateResponse{} }
func (m *CreateResponse) String() string { return proto.CompactTextString(m) }
func (*CreateResponse) ProtoMessage()    {}
func (*CreateResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_2434d8152598889b, []int{4}
}

func (m *CreateResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_CreateResponse.Unmarshal(m, b)
}
func (m *CreateResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_CreateResponse.Marshal(b, m, deterministic)
}
func (m *CreateResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_CreateResponse.Merge(m, src)
}
func (m *CreateResponse) XXX_Size() int {
	return xxx_messageInfo_CreateResponse.Size(m)
}
func (m *CreateResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_CreateResponse.DiscardUnknown(m)
}

var xxx_messageInfo_CreateResponse proto.InternalMessageInfo

type ReadOptions struct {
	// service name
	Service string `protobuf:"bytes,1,opt,name=service,proto3" json:"service,omitempty"`
	// version of the service
	Version string `protobuf:"bytes,2,opt,name=version,proto3" json:"version,omitempty"`
	// type of service
	Type                 string   `protobuf:"bytes,3,opt,name=type,proto3" json:"type,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *ReadOptions) Reset()         { *m = ReadOptions{} }
func (m *ReadOptions) String() string { return proto.CompactTextString(m) }
func (*ReadOptions) ProtoMessage()    {}
func (*ReadOptions) Descriptor() ([]byte, []int) {
	return fileDescriptor_2434d8152598889b, []int{5}
}

func (m *ReadOptions) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_ReadOptions.Unmarshal(m, b)
}
func (m *ReadOptions) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_ReadOptions.Marshal(b, m, deterministic)
}
func (m *ReadOptions) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ReadOptions.Merge(m, src)
}
func (m *ReadOptions) XXX_Size() int {
	return xxx_messageInfo_ReadOptions.Size(m)
}
func (m *ReadOptions) XXX_DiscardUnknown() {
	xxx_messageInfo_ReadOptions.DiscardUnknown(m)
}

var xxx_messageInfo_ReadOptions proto.InternalMessageInfo

func (m *ReadOptions) GetService() string {
	if m != nil {
		return m.Service
	}
	return ""
}

func (m *ReadOptions) GetVersion() string {
	if m != nil {
		return m.Version
	}
	return ""
}

func (m *ReadOptions) GetType() string {
	if m != nil {
		return m.Type
	}
	return ""
}

type ReadRequest struct {
	Options              *ReadOptions `protobuf:"bytes,1,opt,name=options,proto3" json:"options,omitempty"`
	XXX_NoUnkeyedLiteral struct{}     `json:"-"`
	XXX_unrecognized     []byte       `json:"-"`
	XXX_sizecache        int32        `json:"-"`
}

func (m *ReadRequest) Reset()         { *m = ReadRequest{} }
func (m *ReadRequest) String() string { return proto.CompactTextString(m) }
func (*ReadRequest) ProtoMessage()    {}
func (*ReadRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_2434d8152598889b, []int{6}
}

func (m *ReadRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_ReadRequest.Unmarshal(m, b)
}
func (m *ReadRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_ReadRequest.Marshal(b, m, deterministic)
}
func (m *ReadRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ReadRequest.Merge(m, src)
}
func (m *ReadRequest) XXX_Size() int {
	return xxx_messageInfo_ReadRequest.Size(m)
}
func (m *ReadRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_ReadRequest.DiscardUnknown(m)
}

var xxx_messageInfo_ReadRequest proto.InternalMessageInfo

func (m *ReadRequest) GetOptions() *ReadOptions {
	if m != nil {
		return m.Options
	}
	return nil
}

type ReadResponse struct {
	Services             []*Service `protobuf:"bytes,1,rep,name=services,proto3" json:"services,omitempty"`
	XXX_NoUnkeyedLiteral struct{}   `json:"-"`
	XXX_unrecognized     []byte     `json:"-"`
	XXX_sizecache        int32      `json:"-"`
}

func (m *ReadResponse) Reset()         { *m = ReadResponse{} }
func (m *ReadResponse) String() string { return proto.CompactTextString(m) }
func (*ReadResponse) ProtoMessage()    {}
func (*ReadResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_2434d8152598889b, []int{7}
}

func (m *ReadResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_ReadResponse.Unmarshal(m, b)
}
func (m *ReadResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_ReadResponse.Marshal(b, m, deterministic)
}
func (m *ReadResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ReadResponse.Merge(m, src)
}
func (m *ReadResponse) XXX_Size() int {
	return xxx_messageInfo_ReadResponse.Size(m)
}
func (m *ReadResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_ReadResponse.DiscardUnknown(m)
}

var xxx_messageInfo_ReadResponse proto.InternalMessageInfo

func (m *ReadResponse) GetServices() []*Service {
	if m != nil {
		return m.Services
	}
	return nil
}

type DeleteRequest struct {
	Service              *Service `protobuf:"bytes,1,opt,name=service,proto3" json:"service,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *DeleteRequest) Reset()         { *m = DeleteRequest{} }
func (m *DeleteRequest) String() string { return proto.CompactTextString(m) }
func (*DeleteRequest) ProtoMessage()    {}
func (*DeleteRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_2434d8152598889b, []int{8}
}

func (m *DeleteRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_DeleteRequest.Unmarshal(m, b)
}
func (m *DeleteRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_DeleteRequest.Marshal(b, m, deterministic)
}
func (m *DeleteRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_DeleteRequest.Merge(m, src)
}
func (m *DeleteRequest) XXX_Size() int {
	return xxx_messageInfo_DeleteRequest.Size(m)
}
func (m *DeleteRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_DeleteRequest.DiscardUnknown(m)
}

var xxx_messageInfo_DeleteRequest proto.InternalMessageInfo

func (m *DeleteRequest) GetService() *Service {
	if m != nil {
		return m.Service
	}
	return nil
}

type DeleteResponse struct {
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *DeleteResponse) Reset()         { *m = DeleteResponse{} }
func (m *DeleteResponse) String() string { return proto.CompactTextString(m) }
func (*DeleteResponse) ProtoMessage()    {}
func (*DeleteResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_2434d8152598889b, []int{9}
}

func (m *DeleteResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_DeleteResponse.Unmarshal(m, b)
}
func (m *DeleteResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_DeleteResponse.Marshal(b, m, deterministic)
}
func (m *DeleteResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_DeleteResponse.Merge(m, src)
}
func (m *DeleteResponse) XXX_Size() int {
	return xxx_messageInfo_DeleteResponse.Size(m)
}
func (m *DeleteResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_DeleteResponse.DiscardUnknown(m)
}

var xxx_messageInfo_DeleteResponse proto.InternalMessageInfo

type UpdateRequest struct {
	Service              *Service `protobuf:"bytes,1,opt,name=service,proto3" json:"service,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *UpdateRequest) Reset()         { *m = UpdateRequest{} }
func (m *UpdateRequest) String() string { return proto.CompactTextString(m) }
func (*UpdateRequest) ProtoMessage()    {}
func (*UpdateRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_2434d8152598889b, []int{10}
}

func (m *UpdateRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_UpdateRequest.Unmarshal(m, b)
}
func (m *UpdateRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_UpdateRequest.Marshal(b, m, deterministic)
}
func (m *UpdateRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_UpdateRequest.Merge(m, src)
}
func (m *UpdateRequest) XXX_Size() int {
	return xxx_messageInfo_UpdateRequest.Size(m)
}
func (m *UpdateRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_UpdateRequest.DiscardUnknown(m)
}

var xxx_messageInfo_UpdateRequest proto.InternalMessageInfo

func (m *UpdateRequest) GetService() *Service {
	if m != nil {
		return m.Service
	}
	return nil
}

type UpdateResponse struct {
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *UpdateResponse) Reset()         { *m = UpdateResponse{} }
func (m *UpdateResponse) String() string { return proto.CompactTextString(m) }
func (*UpdateResponse) ProtoMessage()    {}
func (*UpdateResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_2434d8152598889b, []int{11}
}

func (m *UpdateResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_UpdateResponse.Unmarshal(m, b)
}
func (m *UpdateResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_UpdateResponse.Marshal(b, m, deterministic)
}
func (m *UpdateResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_UpdateResponse.Merge(m, src)
}
func (m *UpdateResponse) XXX_Size() int {
	return xxx_messageInfo_UpdateResponse.Size(m)
}
func (m *UpdateResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_UpdateResponse.DiscardUnknown(m)
}

var xxx_messageInfo_UpdateResponse proto.InternalMessageInfo

type ListRequest struct {
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *ListRequest) Reset()         { *m = ListRequest{} }
func (m *ListRequest) String() string { return proto.CompactTextString(m) }
func (*ListRequest) ProtoMessage()    {}
func (*ListRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_2434d8152598889b, []int{12}
}

func (m *ListRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_ListRequest.Unmarshal(m, b)
}
func (m *ListRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_ListRequest.Marshal(b, m, deterministic)
}
func (m *ListRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ListRequest.Merge(m, src)
}
func (m *ListRequest) XXX_Size() int {
	return xxx_messageInfo_ListRequest.Size(m)
}
func (m *ListRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_ListRequest.DiscardUnknown(m)
}

var xxx_messageInfo_ListRequest proto.InternalMessageInfo

type ListResponse struct {
	Services             []*Service `protobuf:"bytes,1,rep,name=services,proto3" json:"services,omitempty"`
	XXX_NoUnkeyedLiteral struct{}   `json:"-"`
	XXX_unrecognized     []byte     `json:"-"`
	XXX_sizecache        int32      `json:"-"`
}

func (m *ListResponse) Reset()         { *m = ListResponse{} }
func (m *ListResponse) String() string { return proto.CompactTextString(m) }
func (*ListResponse) ProtoMessage()    {}
func (*ListResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_2434d8152598889b, []int{13}
}

func (m *ListResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_ListResponse.Unmarshal(m, b)
}
func (m *ListResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_ListResponse.Marshal(b, m, deterministic)
}
func (m *ListResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ListResponse.Merge(m, src)
}
func (m *ListResponse) XXX_Size() int {
	return xxx_messageInfo_ListResponse.Size(m)
}
func (m *ListResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_ListResponse.DiscardUnknown(m)
}

var xxx_messageInfo_ListResponse proto.InternalMessageInfo

func (m *ListResponse) GetServices() []*Service {
	if m != nil {
		return m.Services
	}
	return nil
}

func init() {
	proto.RegisterType((*Service)(nil), "go.micro.runtime.Service")
	proto.RegisterMapType((map[string]string)(nil), "go.micro.runtime.Service.MetadataEntry")
	proto.RegisterType((*Event)(nil), "go.micro.runtime.Event")
	proto.RegisterType((*CreateOptions)(nil), "go.micro.runtime.CreateOptions")
	proto.RegisterType((*CreateRequest)(nil), "go.micro.runtime.CreateRequest")
	proto.RegisterType((*CreateResponse)(nil), "go.micro.runtime.CreateResponse")
	proto.RegisterType((*ReadOptions)(nil), "go.micro.runtime.ReadOptions")
	proto.RegisterType((*ReadRequest)(nil), "go.micro.runtime.ReadRequest")
	proto.RegisterType((*ReadResponse)(nil), "go.micro.runtime.ReadResponse")
	proto.RegisterType((*DeleteRequest)(nil), "go.micro.runtime.DeleteRequest")
	proto.RegisterType((*DeleteResponse)(nil), "go.micro.runtime.DeleteResponse")
	proto.RegisterType((*UpdateRequest)(nil), "go.micro.runtime.UpdateRequest")
	proto.RegisterType((*UpdateResponse)(nil), "go.micro.runtime.UpdateResponse")
	proto.RegisterType((*ListRequest)(nil), "go.micro.runtime.ListRequest")
	proto.RegisterType((*ListResponse)(nil), "go.micro.runtime.ListResponse")
}

func init() {
	proto.RegisterFile("runtime/service/proto/runtime.proto", fileDescriptor_2434d8152598889b)
}

var fileDescriptor_2434d8152598889b = []byte{
	// 521 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xa4, 0x54, 0x4b, 0x6f, 0xd3, 0x40,
	0x10, 0xae, 0xeb, 0x34, 0x69, 0xc7, 0x18, 0x45, 0x2b, 0x84, 0x4c, 0xc5, 0x23, 0x32, 0x07, 0x7a,
	0x72, 0xa4, 0x54, 0x88, 0xd7, 0xb1, 0x09, 0x5c, 0x88, 0x90, 0x5c, 0xf5, 0x07, 0x2c, 0xc9, 0x08,
	0x59, 0xd4, 0xbb, 0xc6, 0xbb, 0xb6, 0x94, 0x13, 0x57, 0xfe, 0x1e, 0xff, 0x08, 0xed, 0x2b, 0xb6,
	0x53, 0x9b, 0x4b, 0x6e, 0x3b, 0xb3, 0x33, 0x9f, 0xbf, 0xc7, 0xca, 0xf0, 0xba, 0xac, 0x98, 0xcc,
	0x72, 0x9c, 0x0b, 0x2c, 0xeb, 0x6c, 0x83, 0xf3, 0xa2, 0xe4, 0x92, 0xcf, 0x6d, 0x37, 0xd1, 0x15,
	0x99, 0xfe, 0xe0, 0x49, 0x9e, 0x6d, 0x4a, 0x9e, 0xd8, 0x7e, 0xfc, 0xd7, 0x83, 0xc9, 0xad, 0xd9,
	0x20, 0x04, 0x46, 0x8c, 0xe6, 0x18, 0x79, 0x33, 0xef, 0xea, 0x22, 0xd5, 0x67, 0x12, 0xc1, 0xa4,
	0xc6, 0x52, 0x64, 0x9c, 0x45, 0xa7, 0xba, 0xed, 0x4a, 0xf2, 0x14, 0xc6, 0x82, 0x57, 0xe5, 0x06,
	0x23, 0x5f, 0x5f, 0xd8, 0x8a, 0xdc, 0xc0, 0x79, 0x8e, 0x92, 0x6e, 0xa9, 0xa4, 0xd1, 0x68, 0xe6,
	0x5f, 0x05, 0x8b, 0x37, 0xc9, 0xe1, 0x67, 0x13, 0xfb, 0xc9, 0x64, 0x6d, 0x27, 0x57, 0x4c, 0x96,
	0xbb, 0x74, 0xbf, 0x78, 0xf9, 0x09, 0xc2, 0xce, 0x15, 0x99, 0x82, 0xff, 0x13, 0x77, 0x96, 0x9a,
	0x3a, 0x92, 0x27, 0x70, 0x56, 0xd3, 0xfb, 0x0a, 0x2d, 0x2f, 0x53, 0x7c, 0x3c, 0x7d, 0xef, 0xc5,
	0x39, 0x9c, 0xad, 0x6a, 0x64, 0x52, 0x09, 0x92, 0xbb, 0x62, 0x2f, 0x48, 0x9d, 0xc9, 0x73, 0xb8,
	0x50, 0x0c, 0x84, 0xa4, 0x79, 0xa1, 0x57, 0xfd, 0xb4, 0x69, 0x28, 0xb9, 0xd6, 0x3f, 0xab, 0xca,
	0x95, 0x6d, 0x23, 0x46, 0x1d, 0x23, 0xe2, 0x5b, 0x08, 0x6f, 0x4a, 0xa4, 0x12, 0xbf, 0x15, 0x32,
	0xe3, 0x4c, 0xa8, 0xd1, 0x0d, 0xcf, 0x73, 0xca, 0xb6, 0x91, 0x37, 0xf3, 0xd5, 0xa8, 0x2d, 0x95,
	0x0a, 0x64, 0x75, 0x74, 0xaa, 0xbb, 0xea, 0xa8, 0x5c, 0xe4, 0x95, 0x2c, 0x2a, 0xe9, 0x5c, 0x34,
	0x55, 0xfc, 0xdb, 0x81, 0xa6, 0xf8, 0xab, 0x42, 0x21, 0xc9, 0x75, 0xc3, 0x4c, 0xc9, 0x09, 0x16,
	0xcf, 0x06, 0x5d, 0x6d, 0x48, 0x7f, 0x80, 0x09, 0x37, 0xa4, 0xb4, 0xd4, 0x60, 0xf1, 0xea, 0xe1,
	0x52, 0x87, 0x7b, 0xea, 0xe6, 0xe3, 0x29, 0x3c, 0x76, 0x04, 0x44, 0xc1, 0x99, 0xc0, 0xf8, 0x0e,
	0x82, 0x14, 0xe9, 0xb6, 0xa5, 0xb2, 0x4d, 0xa8, 0xdf, 0xaa, 0x83, 0x37, 0xe3, 0x02, 0xf1, 0x9b,
	0x40, 0xe2, 0xcf, 0x06, 0xd6, 0xe9, 0x7c, 0xd7, 0x50, 0x36, 0x3a, 0x5f, 0x3c, 0xa4, 0xdc, 0xa2,
	0xd1, 0x10, 0x5e, 0xc1, 0x23, 0x83, 0x63, 0xe8, 0x92, 0xb7, 0x70, 0x6e, 0x09, 0x09, 0x1d, 0xc3,
	0x7f, 0x1d, 0xdb, 0x8f, 0xc6, 0x4b, 0x08, 0x97, 0x78, 0x8f, 0xc7, 0x19, 0xaf, 0xdc, 0x73, 0x28,
	0xd6, 0xbd, 0x25, 0x84, 0x77, 0xc5, 0x96, 0x1e, 0x8f, 0xeb, 0x50, 0x2c, 0x6e, 0x08, 0xc1, 0xd7,
	0x4c, 0x48, 0x8b, 0xaa, 0x5c, 0x30, 0xe5, 0x51, 0x2e, 0x2c, 0xfe, 0xf8, 0x30, 0x49, 0xcd, 0x2d,
	0x59, 0xc3, 0xd8, 0xbc, 0x04, 0x32, 0xf8, 0x7a, 0xec, 0xd7, 0x2f, 0x67, 0xc3, 0x03, 0x96, 0xee,
	0x09, 0xf9, 0x02, 0x23, 0x95, 0x13, 0x19, 0xc8, 0xd5, 0x41, 0xbd, 0x1c, 0xba, 0xde, 0x03, 0xad,
	0x61, 0x6c, 0x3c, 0xee, 0xe3, 0xd5, 0xc9, 0xb0, 0x8f, 0xd7, 0x41, 0x3c, 0x1a, 0xce, 0x58, 0xdb,
	0x07, 0xd7, 0x89, 0xae, 0x0f, 0xee, 0x20, 0x15, 0x2d, 0x53, 0x05, 0xd1, 0x27, 0xb3, 0x95, 0x57,
	0x9f, 0xcc, 0x76, 0x7e, 0xf1, 0xc9, 0xf7, 0xb1, 0xfe, 0x75, 0x5f, 0xff, 0x0b, 0x00, 0x00, 0xff,
	0xff, 0x43, 0x9c, 0x97, 0x62, 0xe1, 0x05, 0x00, 0x00,
}
