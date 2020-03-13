// Code generated by protoc-gen-go. DO NOT EDIT.
// source: proto/post.proto

package postpb

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

type Post struct {
	Id                   string   `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
	ForumId              string   `protobuf:"bytes,2,opt,name=forum_id,json=forumId,proto3" json:"forum_id,omitempty"`
	AuthorId             string   `protobuf:"bytes,3,opt,name=author_id,json=authorId,proto3" json:"author_id,omitempty"`
	Title                string   `protobuf:"bytes,4,opt,name=title,proto3" json:"title,omitempty"`
	Content              string   `protobuf:"bytes,5,opt,name=content,proto3" json:"content,omitempty"`
	Timestamp            string   `protobuf:"bytes,6,opt,name=timestamp,proto3" json:"timestamp,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *Post) Reset()         { *m = Post{} }
func (m *Post) String() string { return proto.CompactTextString(m) }
func (*Post) ProtoMessage()    {}
func (*Post) Descriptor() ([]byte, []int) {
	return fileDescriptor_5e57bba97c118b83, []int{0}
}

func (m *Post) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Post.Unmarshal(m, b)
}
func (m *Post) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Post.Marshal(b, m, deterministic)
}
func (m *Post) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Post.Merge(m, src)
}
func (m *Post) XXX_Size() int {
	return xxx_messageInfo_Post.Size(m)
}
func (m *Post) XXX_DiscardUnknown() {
	xxx_messageInfo_Post.DiscardUnknown(m)
}

var xxx_messageInfo_Post proto.InternalMessageInfo

func (m *Post) GetId() string {
	if m != nil {
		return m.Id
	}
	return ""
}

func (m *Post) GetForumId() string {
	if m != nil {
		return m.ForumId
	}
	return ""
}

func (m *Post) GetAuthorId() string {
	if m != nil {
		return m.AuthorId
	}
	return ""
}

func (m *Post) GetTitle() string {
	if m != nil {
		return m.Title
	}
	return ""
}

func (m *Post) GetContent() string {
	if m != nil {
		return m.Content
	}
	return ""
}

func (m *Post) GetTimestamp() string {
	if m != nil {
		return m.Timestamp
	}
	return ""
}

type CreatePostReq struct {
	Post                 *Post    `protobuf:"bytes,1,opt,name=post,proto3" json:"post,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *CreatePostReq) Reset()         { *m = CreatePostReq{} }
func (m *CreatePostReq) String() string { return proto.CompactTextString(m) }
func (*CreatePostReq) ProtoMessage()    {}
func (*CreatePostReq) Descriptor() ([]byte, []int) {
	return fileDescriptor_5e57bba97c118b83, []int{1}
}

func (m *CreatePostReq) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_CreatePostReq.Unmarshal(m, b)
}
func (m *CreatePostReq) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_CreatePostReq.Marshal(b, m, deterministic)
}
func (m *CreatePostReq) XXX_Merge(src proto.Message) {
	xxx_messageInfo_CreatePostReq.Merge(m, src)
}
func (m *CreatePostReq) XXX_Size() int {
	return xxx_messageInfo_CreatePostReq.Size(m)
}
func (m *CreatePostReq) XXX_DiscardUnknown() {
	xxx_messageInfo_CreatePostReq.DiscardUnknown(m)
}

var xxx_messageInfo_CreatePostReq proto.InternalMessageInfo

func (m *CreatePostReq) GetPost() *Post {
	if m != nil {
		return m.Post
	}
	return nil
}

type CreatePostRes struct {
	Post                 *Post    `protobuf:"bytes,1,opt,name=post,proto3" json:"post,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *CreatePostRes) Reset()         { *m = CreatePostRes{} }
func (m *CreatePostRes) String() string { return proto.CompactTextString(m) }
func (*CreatePostRes) ProtoMessage()    {}
func (*CreatePostRes) Descriptor() ([]byte, []int) {
	return fileDescriptor_5e57bba97c118b83, []int{2}
}

func (m *CreatePostRes) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_CreatePostRes.Unmarshal(m, b)
}
func (m *CreatePostRes) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_CreatePostRes.Marshal(b, m, deterministic)
}
func (m *CreatePostRes) XXX_Merge(src proto.Message) {
	xxx_messageInfo_CreatePostRes.Merge(m, src)
}
func (m *CreatePostRes) XXX_Size() int {
	return xxx_messageInfo_CreatePostRes.Size(m)
}
func (m *CreatePostRes) XXX_DiscardUnknown() {
	xxx_messageInfo_CreatePostRes.DiscardUnknown(m)
}

var xxx_messageInfo_CreatePostRes proto.InternalMessageInfo

func (m *CreatePostRes) GetPost() *Post {
	if m != nil {
		return m.Post
	}
	return nil
}

type ReadPostReq struct {
	Id                   string   `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *ReadPostReq) Reset()         { *m = ReadPostReq{} }
func (m *ReadPostReq) String() string { return proto.CompactTextString(m) }
func (*ReadPostReq) ProtoMessage()    {}
func (*ReadPostReq) Descriptor() ([]byte, []int) {
	return fileDescriptor_5e57bba97c118b83, []int{3}
}

func (m *ReadPostReq) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_ReadPostReq.Unmarshal(m, b)
}
func (m *ReadPostReq) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_ReadPostReq.Marshal(b, m, deterministic)
}
func (m *ReadPostReq) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ReadPostReq.Merge(m, src)
}
func (m *ReadPostReq) XXX_Size() int {
	return xxx_messageInfo_ReadPostReq.Size(m)
}
func (m *ReadPostReq) XXX_DiscardUnknown() {
	xxx_messageInfo_ReadPostReq.DiscardUnknown(m)
}

var xxx_messageInfo_ReadPostReq proto.InternalMessageInfo

func (m *ReadPostReq) GetId() string {
	if m != nil {
		return m.Id
	}
	return ""
}

type ReadPostRes struct {
	Post                 *Post    `protobuf:"bytes,1,opt,name=post,proto3" json:"post,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *ReadPostRes) Reset()         { *m = ReadPostRes{} }
func (m *ReadPostRes) String() string { return proto.CompactTextString(m) }
func (*ReadPostRes) ProtoMessage()    {}
func (*ReadPostRes) Descriptor() ([]byte, []int) {
	return fileDescriptor_5e57bba97c118b83, []int{4}
}

func (m *ReadPostRes) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_ReadPostRes.Unmarshal(m, b)
}
func (m *ReadPostRes) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_ReadPostRes.Marshal(b, m, deterministic)
}
func (m *ReadPostRes) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ReadPostRes.Merge(m, src)
}
func (m *ReadPostRes) XXX_Size() int {
	return xxx_messageInfo_ReadPostRes.Size(m)
}
func (m *ReadPostRes) XXX_DiscardUnknown() {
	xxx_messageInfo_ReadPostRes.DiscardUnknown(m)
}

var xxx_messageInfo_ReadPostRes proto.InternalMessageInfo

func (m *ReadPostRes) GetPost() *Post {
	if m != nil {
		return m.Post
	}
	return nil
}

type UpdatePostReq struct {
	Post                 *Post    `protobuf:"bytes,1,opt,name=post,proto3" json:"post,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *UpdatePostReq) Reset()         { *m = UpdatePostReq{} }
func (m *UpdatePostReq) String() string { return proto.CompactTextString(m) }
func (*UpdatePostReq) ProtoMessage()    {}
func (*UpdatePostReq) Descriptor() ([]byte, []int) {
	return fileDescriptor_5e57bba97c118b83, []int{5}
}

func (m *UpdatePostReq) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_UpdatePostReq.Unmarshal(m, b)
}
func (m *UpdatePostReq) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_UpdatePostReq.Marshal(b, m, deterministic)
}
func (m *UpdatePostReq) XXX_Merge(src proto.Message) {
	xxx_messageInfo_UpdatePostReq.Merge(m, src)
}
func (m *UpdatePostReq) XXX_Size() int {
	return xxx_messageInfo_UpdatePostReq.Size(m)
}
func (m *UpdatePostReq) XXX_DiscardUnknown() {
	xxx_messageInfo_UpdatePostReq.DiscardUnknown(m)
}

var xxx_messageInfo_UpdatePostReq proto.InternalMessageInfo

func (m *UpdatePostReq) GetPost() *Post {
	if m != nil {
		return m.Post
	}
	return nil
}

type UpdatePostRes struct {
	Post                 *Post    `protobuf:"bytes,1,opt,name=post,proto3" json:"post,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *UpdatePostRes) Reset()         { *m = UpdatePostRes{} }
func (m *UpdatePostRes) String() string { return proto.CompactTextString(m) }
func (*UpdatePostRes) ProtoMessage()    {}
func (*UpdatePostRes) Descriptor() ([]byte, []int) {
	return fileDescriptor_5e57bba97c118b83, []int{6}
}

func (m *UpdatePostRes) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_UpdatePostRes.Unmarshal(m, b)
}
func (m *UpdatePostRes) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_UpdatePostRes.Marshal(b, m, deterministic)
}
func (m *UpdatePostRes) XXX_Merge(src proto.Message) {
	xxx_messageInfo_UpdatePostRes.Merge(m, src)
}
func (m *UpdatePostRes) XXX_Size() int {
	return xxx_messageInfo_UpdatePostRes.Size(m)
}
func (m *UpdatePostRes) XXX_DiscardUnknown() {
	xxx_messageInfo_UpdatePostRes.DiscardUnknown(m)
}

var xxx_messageInfo_UpdatePostRes proto.InternalMessageInfo

func (m *UpdatePostRes) GetPost() *Post {
	if m != nil {
		return m.Post
	}
	return nil
}

type DeletePostReq struct {
	Id                   string   `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *DeletePostReq) Reset()         { *m = DeletePostReq{} }
func (m *DeletePostReq) String() string { return proto.CompactTextString(m) }
func (*DeletePostReq) ProtoMessage()    {}
func (*DeletePostReq) Descriptor() ([]byte, []int) {
	return fileDescriptor_5e57bba97c118b83, []int{7}
}

func (m *DeletePostReq) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_DeletePostReq.Unmarshal(m, b)
}
func (m *DeletePostReq) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_DeletePostReq.Marshal(b, m, deterministic)
}
func (m *DeletePostReq) XXX_Merge(src proto.Message) {
	xxx_messageInfo_DeletePostReq.Merge(m, src)
}
func (m *DeletePostReq) XXX_Size() int {
	return xxx_messageInfo_DeletePostReq.Size(m)
}
func (m *DeletePostReq) XXX_DiscardUnknown() {
	xxx_messageInfo_DeletePostReq.DiscardUnknown(m)
}

var xxx_messageInfo_DeletePostReq proto.InternalMessageInfo

func (m *DeletePostReq) GetId() string {
	if m != nil {
		return m.Id
	}
	return ""
}

type DeletePostRes struct {
	Success              bool     `protobuf:"varint,1,opt,name=success,proto3" json:"success,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *DeletePostRes) Reset()         { *m = DeletePostRes{} }
func (m *DeletePostRes) String() string { return proto.CompactTextString(m) }
func (*DeletePostRes) ProtoMessage()    {}
func (*DeletePostRes) Descriptor() ([]byte, []int) {
	return fileDescriptor_5e57bba97c118b83, []int{8}
}

func (m *DeletePostRes) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_DeletePostRes.Unmarshal(m, b)
}
func (m *DeletePostRes) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_DeletePostRes.Marshal(b, m, deterministic)
}
func (m *DeletePostRes) XXX_Merge(src proto.Message) {
	xxx_messageInfo_DeletePostRes.Merge(m, src)
}
func (m *DeletePostRes) XXX_Size() int {
	return xxx_messageInfo_DeletePostRes.Size(m)
}
func (m *DeletePostRes) XXX_DiscardUnknown() {
	xxx_messageInfo_DeletePostRes.DiscardUnknown(m)
}

var xxx_messageInfo_DeletePostRes proto.InternalMessageInfo

func (m *DeletePostRes) GetSuccess() bool {
	if m != nil {
		return m.Success
	}
	return false
}

type ListPostReq struct {
	ForumId              string   `protobuf:"bytes,1,opt,name=forum_id,json=forumId,proto3" json:"forum_id,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *ListPostReq) Reset()         { *m = ListPostReq{} }
func (m *ListPostReq) String() string { return proto.CompactTextString(m) }
func (*ListPostReq) ProtoMessage()    {}
func (*ListPostReq) Descriptor() ([]byte, []int) {
	return fileDescriptor_5e57bba97c118b83, []int{9}
}

func (m *ListPostReq) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_ListPostReq.Unmarshal(m, b)
}
func (m *ListPostReq) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_ListPostReq.Marshal(b, m, deterministic)
}
func (m *ListPostReq) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ListPostReq.Merge(m, src)
}
func (m *ListPostReq) XXX_Size() int {
	return xxx_messageInfo_ListPostReq.Size(m)
}
func (m *ListPostReq) XXX_DiscardUnknown() {
	xxx_messageInfo_ListPostReq.DiscardUnknown(m)
}

var xxx_messageInfo_ListPostReq proto.InternalMessageInfo

func (m *ListPostReq) GetForumId() string {
	if m != nil {
		return m.ForumId
	}
	return ""
}

type ListPostRes struct {
	Post                 *Post    `protobuf:"bytes,1,opt,name=post,proto3" json:"post,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *ListPostRes) Reset()         { *m = ListPostRes{} }
func (m *ListPostRes) String() string { return proto.CompactTextString(m) }
func (*ListPostRes) ProtoMessage()    {}
func (*ListPostRes) Descriptor() ([]byte, []int) {
	return fileDescriptor_5e57bba97c118b83, []int{10}
}

func (m *ListPostRes) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_ListPostRes.Unmarshal(m, b)
}
func (m *ListPostRes) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_ListPostRes.Marshal(b, m, deterministic)
}
func (m *ListPostRes) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ListPostRes.Merge(m, src)
}
func (m *ListPostRes) XXX_Size() int {
	return xxx_messageInfo_ListPostRes.Size(m)
}
func (m *ListPostRes) XXX_DiscardUnknown() {
	xxx_messageInfo_ListPostRes.DiscardUnknown(m)
}

var xxx_messageInfo_ListPostRes proto.InternalMessageInfo

func (m *ListPostRes) GetPost() *Post {
	if m != nil {
		return m.Post
	}
	return nil
}

func init() {
	proto.RegisterType((*Post)(nil), "post.Post")
	proto.RegisterType((*CreatePostReq)(nil), "post.CreatePostReq")
	proto.RegisterType((*CreatePostRes)(nil), "post.CreatePostRes")
	proto.RegisterType((*ReadPostReq)(nil), "post.ReadPostReq")
	proto.RegisterType((*ReadPostRes)(nil), "post.ReadPostRes")
	proto.RegisterType((*UpdatePostReq)(nil), "post.UpdatePostReq")
	proto.RegisterType((*UpdatePostRes)(nil), "post.UpdatePostRes")
	proto.RegisterType((*DeletePostReq)(nil), "post.DeletePostReq")
	proto.RegisterType((*DeletePostRes)(nil), "post.DeletePostRes")
	proto.RegisterType((*ListPostReq)(nil), "post.ListPostReq")
	proto.RegisterType((*ListPostRes)(nil), "post.ListPostRes")
}

func init() {
	proto.RegisterFile("proto/post.proto", fileDescriptor_5e57bba97c118b83)
}

var fileDescriptor_5e57bba97c118b83 = []byte{
	// 362 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x94, 0x53, 0xcd, 0x4a, 0xf3, 0x40,
	0x14, 0x25, 0xf9, 0xfa, 0x93, 0xdc, 0xd0, 0x0f, 0x1d, 0x5d, 0xc4, 0xf8, 0x4b, 0x56, 0x75, 0x61,
	0x5b, 0x5a, 0xf0, 0x01, 0xd4, 0x4d, 0xc1, 0x85, 0x44, 0xdc, 0xb8, 0x91, 0x34, 0x33, 0xe2, 0x40,
	0xdb, 0x89, 0xb9, 0x53, 0x9f, 0xa6, 0x0f, 0x2b, 0x73, 0xe3, 0x90, 0xa4, 0x2d, 0x54, 0x77, 0xb9,
	0xe7, 0xdc, 0x9f, 0x33, 0xe7, 0x10, 0x38, 0xc8, 0x0b, 0xa5, 0xd5, 0x30, 0x57, 0xa8, 0x07, 0xf4,
	0xc9, 0x5a, 0xe6, 0x3b, 0x5e, 0x3b, 0xd0, 0x7a, 0x52, 0xa8, 0xd9, 0x7f, 0x70, 0x25, 0x0f, 0x9d,
	0x2b, 0xa7, 0xef, 0x27, 0xae, 0xe4, 0xec, 0x04, 0xbc, 0x77, 0x55, 0xac, 0x16, 0x6f, 0x92, 0x87,
	0x2e, 0xa1, 0x5d, 0xaa, 0xa7, 0x9c, 0x9d, 0x82, 0x9f, 0xae, 0xf4, 0x87, 0x2a, 0x0c, 0xf7, 0x8f,
	0x38, 0xaf, 0x04, 0xa6, 0x9c, 0x1d, 0x43, 0x5b, 0x4b, 0x3d, 0x17, 0x61, 0x8b, 0x88, 0xb2, 0x60,
	0x21, 0x74, 0x33, 0xb5, 0xd4, 0x62, 0xa9, 0xc3, 0x76, 0xb9, 0xec, 0xa7, 0x64, 0x67, 0xe0, 0x6b,
	0xb9, 0x10, 0xa8, 0xd3, 0x45, 0x1e, 0x76, 0x88, 0xab, 0x80, 0x78, 0x08, 0xbd, 0xfb, 0x42, 0xa4,
	0x5a, 0x18, 0x8d, 0x89, 0xf8, 0x64, 0x17, 0x40, 0xba, 0x49, 0x68, 0x30, 0x86, 0x01, 0x3d, 0x88,
	0xc8, 0xf2, 0x3d, 0x1b, 0x03, 0xb8, 0x77, 0xe0, 0x1c, 0x82, 0x44, 0xa4, 0xdc, 0xee, 0xdf, 0xb0,
	0x21, 0xbe, 0xa9, 0xd3, 0xf8, 0x9b, 0xf3, 0x2f, 0x39, 0xff, 0x9b, 0xde, 0xfa, 0xc0, 0xfe, 0x0b,
	0x97, 0xd0, 0x7b, 0x10, 0x73, 0x51, 0x5d, 0xd8, 0x54, 0x7c, 0xdd, 0x6c, 0x40, 0xe3, 0x3d, 0xae,
	0xb2, 0x4c, 0x20, 0x52, 0x97, 0x97, 0xd8, 0x32, 0xee, 0x43, 0xf0, 0x28, 0x51, 0xdb, 0x4d, 0xf5,
	0xc8, 0x9d, 0x46, 0xe4, 0xc6, 0x86, 0xaa, 0x73, 0xaf, 0xc8, 0xf1, 0xda, 0x85, 0xc0, 0x94, 0xcf,
	0xa2, 0xf8, 0x92, 0x99, 0x60, 0xb7, 0x00, 0x55, 0x2a, 0xec, 0xa8, 0xec, 0x6f, 0x04, 0x1b, 0xed,
	0x00, 0x91, 0x8d, 0xc0, 0xb3, 0xee, 0xb3, 0xc3, 0xb2, 0xa1, 0x16, 0x56, 0xb4, 0x05, 0xa1, 0xb9,
	0x54, 0xf9, 0x69, 0x2f, 0x35, 0x22, 0x89, 0x76, 0x80, 0x34, 0x57, 0xb9, 0x66, 0xe7, 0x1a, 0x46,
	0x47, 0x3b, 0x40, 0x64, 0x13, 0xf0, 0xad, 0x31, 0x68, 0x25, 0xd6, 0x3c, 0x8d, 0xb6, 0x20, 0x1c,
	0x39, 0x77, 0xde, 0x6b, 0xc7, 0xa0, 0xf9, 0x6c, 0xd6, 0xa1, 0x7f, 0x71, 0xf2, 0x1d, 0x00, 0x00,
	0xff, 0xff, 0xcb, 0x6d, 0xc3, 0x99, 0x9f, 0x03, 0x00, 0x00,
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConnInterface

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion6

// PostServiceClient is the client API for PostService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type PostServiceClient interface {
	CreatePost(ctx context.Context, in *CreatePostReq, opts ...grpc.CallOption) (*CreatePostRes, error)
	ReadPost(ctx context.Context, in *ReadPostReq, opts ...grpc.CallOption) (*ReadPostRes, error)
	UpdatePost(ctx context.Context, in *UpdatePostReq, opts ...grpc.CallOption) (*UpdatePostRes, error)
	DeletePost(ctx context.Context, in *DeletePostReq, opts ...grpc.CallOption) (*DeletePostRes, error)
	ListPosts(ctx context.Context, in *ListPostReq, opts ...grpc.CallOption) (PostService_ListPostsClient, error)
}

type postServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewPostServiceClient(cc grpc.ClientConnInterface) PostServiceClient {
	return &postServiceClient{cc}
}

func (c *postServiceClient) CreatePost(ctx context.Context, in *CreatePostReq, opts ...grpc.CallOption) (*CreatePostRes, error) {
	out := new(CreatePostRes)
	err := c.cc.Invoke(ctx, "/post.PostService/CreatePost", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *postServiceClient) ReadPost(ctx context.Context, in *ReadPostReq, opts ...grpc.CallOption) (*ReadPostRes, error) {
	out := new(ReadPostRes)
	err := c.cc.Invoke(ctx, "/post.PostService/ReadPost", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *postServiceClient) UpdatePost(ctx context.Context, in *UpdatePostReq, opts ...grpc.CallOption) (*UpdatePostRes, error) {
	out := new(UpdatePostRes)
	err := c.cc.Invoke(ctx, "/post.PostService/UpdatePost", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *postServiceClient) DeletePost(ctx context.Context, in *DeletePostReq, opts ...grpc.CallOption) (*DeletePostRes, error) {
	out := new(DeletePostRes)
	err := c.cc.Invoke(ctx, "/post.PostService/DeletePost", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *postServiceClient) ListPosts(ctx context.Context, in *ListPostReq, opts ...grpc.CallOption) (PostService_ListPostsClient, error) {
	stream, err := c.cc.NewStream(ctx, &_PostService_serviceDesc.Streams[0], "/post.PostService/ListPosts", opts...)
	if err != nil {
		return nil, err
	}
	x := &postServiceListPostsClient{stream}
	if err := x.ClientStream.SendMsg(in); err != nil {
		return nil, err
	}
	if err := x.ClientStream.CloseSend(); err != nil {
		return nil, err
	}
	return x, nil
}

type PostService_ListPostsClient interface {
	Recv() (*ListPostRes, error)
	grpc.ClientStream
}

type postServiceListPostsClient struct {
	grpc.ClientStream
}

func (x *postServiceListPostsClient) Recv() (*ListPostRes, error) {
	m := new(ListPostRes)
	if err := x.ClientStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

// PostServiceServer is the server API for PostService service.
type PostServiceServer interface {
	CreatePost(context.Context, *CreatePostReq) (*CreatePostRes, error)
	ReadPost(context.Context, *ReadPostReq) (*ReadPostRes, error)
	UpdatePost(context.Context, *UpdatePostReq) (*UpdatePostRes, error)
	DeletePost(context.Context, *DeletePostReq) (*DeletePostRes, error)
	ListPosts(*ListPostReq, PostService_ListPostsServer) error
}

// UnimplementedPostServiceServer can be embedded to have forward compatible implementations.
type UnimplementedPostServiceServer struct {
}

func (*UnimplementedPostServiceServer) CreatePost(ctx context.Context, req *CreatePostReq) (*CreatePostRes, error) {
	return nil, status.Errorf(codes.Unimplemented, "method CreatePost not implemented")
}
func (*UnimplementedPostServiceServer) ReadPost(ctx context.Context, req *ReadPostReq) (*ReadPostRes, error) {
	return nil, status.Errorf(codes.Unimplemented, "method ReadPost not implemented")
}
func (*UnimplementedPostServiceServer) UpdatePost(ctx context.Context, req *UpdatePostReq) (*UpdatePostRes, error) {
	return nil, status.Errorf(codes.Unimplemented, "method UpdatePost not implemented")
}
func (*UnimplementedPostServiceServer) DeletePost(ctx context.Context, req *DeletePostReq) (*DeletePostRes, error) {
	return nil, status.Errorf(codes.Unimplemented, "method DeletePost not implemented")
}
func (*UnimplementedPostServiceServer) ListPosts(req *ListPostReq, srv PostService_ListPostsServer) error {
	return status.Errorf(codes.Unimplemented, "method ListPosts not implemented")
}

func RegisterPostServiceServer(s *grpc.Server, srv PostServiceServer) {
	s.RegisterService(&_PostService_serviceDesc, srv)
}

func _PostService_CreatePost_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(CreatePostReq)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(PostServiceServer).CreatePost(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/post.PostService/CreatePost",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(PostServiceServer).CreatePost(ctx, req.(*CreatePostReq))
	}
	return interceptor(ctx, in, info, handler)
}

func _PostService_ReadPost_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ReadPostReq)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(PostServiceServer).ReadPost(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/post.PostService/ReadPost",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(PostServiceServer).ReadPost(ctx, req.(*ReadPostReq))
	}
	return interceptor(ctx, in, info, handler)
}

func _PostService_UpdatePost_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(UpdatePostReq)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(PostServiceServer).UpdatePost(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/post.PostService/UpdatePost",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(PostServiceServer).UpdatePost(ctx, req.(*UpdatePostReq))
	}
	return interceptor(ctx, in, info, handler)
}

func _PostService_DeletePost_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(DeletePostReq)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(PostServiceServer).DeletePost(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/post.PostService/DeletePost",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(PostServiceServer).DeletePost(ctx, req.(*DeletePostReq))
	}
	return interceptor(ctx, in, info, handler)
}

func _PostService_ListPosts_Handler(srv interface{}, stream grpc.ServerStream) error {
	m := new(ListPostReq)
	if err := stream.RecvMsg(m); err != nil {
		return err
	}
	return srv.(PostServiceServer).ListPosts(m, &postServiceListPostsServer{stream})
}

type PostService_ListPostsServer interface {
	Send(*ListPostRes) error
	grpc.ServerStream
}

type postServiceListPostsServer struct {
	grpc.ServerStream
}

func (x *postServiceListPostsServer) Send(m *ListPostRes) error {
	return x.ServerStream.SendMsg(m)
}

var _PostService_serviceDesc = grpc.ServiceDesc{
	ServiceName: "post.PostService",
	HandlerType: (*PostServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "CreatePost",
			Handler:    _PostService_CreatePost_Handler,
		},
		{
			MethodName: "ReadPost",
			Handler:    _PostService_ReadPost_Handler,
		},
		{
			MethodName: "UpdatePost",
			Handler:    _PostService_UpdatePost_Handler,
		},
		{
			MethodName: "DeletePost",
			Handler:    _PostService_DeletePost_Handler,
		},
	},
	Streams: []grpc.StreamDesc{
		{
			StreamName:    "ListPosts",
			Handler:       _PostService_ListPosts_Handler,
			ServerStreams: true,
		},
	},
	Metadata: "proto/post.proto",
}
