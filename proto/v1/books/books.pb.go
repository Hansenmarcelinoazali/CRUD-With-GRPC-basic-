// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.26.0
// 	protoc        v3.12.4
// source: v1/books/books.proto

package books

import (
	context "context"
	_ "google.golang.org/genproto/googleapis/api/annotations"
	grpc "google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
	emptypb "google.golang.org/protobuf/types/known/emptypb"
	reflect "reflect"
	sync "sync"
)

const (
	// Verify that this generated code is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(20 - protoimpl.MinVersion)
	// Verify that runtime/protoimpl is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(protoimpl.MaxVersion - 20)
)

type Response struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Success bool     `protobuf:"varint,1,opt,name=success,proto3" json:"success,omitempty"`
	Code    string   `protobuf:"bytes,2,opt,name=code,proto3" json:"code,omitempty"`
	Desc    string   `protobuf:"bytes,3,opt,name=desc,proto3" json:"desc,omitempty"`
	Data    []*Books `protobuf:"bytes,4,rep,name=data,proto3" json:"data,omitempty"`
}

func (x *Response) Reset() {
	*x = Response{}
	if protoimpl.UnsafeEnabled {
		mi := &file_v1_books_books_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Response) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Response) ProtoMessage() {}

func (x *Response) ProtoReflect() protoreflect.Message {
	mi := &file_v1_books_books_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Response.ProtoReflect.Descriptor instead.
func (*Response) Descriptor() ([]byte, []int) {
	return file_v1_books_books_proto_rawDescGZIP(), []int{0}
}

func (x *Response) GetSuccess() bool {
	if x != nil {
		return x.Success
	}
	return false
}

func (x *Response) GetCode() string {
	if x != nil {
		return x.Code
	}
	return ""
}

func (x *Response) GetDesc() string {
	if x != nil {
		return x.Desc
	}
	return ""
}

func (x *Response) GetData() []*Books {
	if x != nil {
		return x.Data
	}
	return nil
}

type ResponseCreate struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// string inputdata = 1;
	Success bool   `protobuf:"varint,1,opt,name=success,proto3" json:"success,omitempty"`
	Code    string `protobuf:"bytes,2,opt,name=code,proto3" json:"code,omitempty"`
	Desc    string `protobuf:"bytes,3,opt,name=desc,proto3" json:"desc,omitempty"`
	Books   *Books `protobuf:"bytes,5,opt,name=books,proto3" json:"books,omitempty"`
}

func (x *ResponseCreate) Reset() {
	*x = ResponseCreate{}
	if protoimpl.UnsafeEnabled {
		mi := &file_v1_books_books_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ResponseCreate) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ResponseCreate) ProtoMessage() {}

func (x *ResponseCreate) ProtoReflect() protoreflect.Message {
	mi := &file_v1_books_books_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ResponseCreate.ProtoReflect.Descriptor instead.
func (*ResponseCreate) Descriptor() ([]byte, []int) {
	return file_v1_books_books_proto_rawDescGZIP(), []int{1}
}

func (x *ResponseCreate) GetSuccess() bool {
	if x != nil {
		return x.Success
	}
	return false
}

func (x *ResponseCreate) GetCode() string {
	if x != nil {
		return x.Code
	}
	return ""
}

func (x *ResponseCreate) GetDesc() string {
	if x != nil {
		return x.Desc
	}
	return ""
}

func (x *ResponseCreate) GetBooks() *Books {
	if x != nil {
		return x.Books
	}
	return nil
}

type ResponseUpdateData struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Success bool   `protobuf:"varint,1,opt,name=success,proto3" json:"success,omitempty"`
	Code    string `protobuf:"bytes,2,opt,name=code,proto3" json:"code,omitempty"`
	Desc    string `protobuf:"bytes,3,opt,name=desc,proto3" json:"desc,omitempty"`
	Books   *Books `protobuf:"bytes,4,opt,name=books,proto3" json:"books,omitempty"`
}

func (x *ResponseUpdateData) Reset() {
	*x = ResponseUpdateData{}
	if protoimpl.UnsafeEnabled {
		mi := &file_v1_books_books_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ResponseUpdateData) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ResponseUpdateData) ProtoMessage() {}

func (x *ResponseUpdateData) ProtoReflect() protoreflect.Message {
	mi := &file_v1_books_books_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ResponseUpdateData.ProtoReflect.Descriptor instead.
func (*ResponseUpdateData) Descriptor() ([]byte, []int) {
	return file_v1_books_books_proto_rawDescGZIP(), []int{2}
}

func (x *ResponseUpdateData) GetSuccess() bool {
	if x != nil {
		return x.Success
	}
	return false
}

func (x *ResponseUpdateData) GetCode() string {
	if x != nil {
		return x.Code
	}
	return ""
}

func (x *ResponseUpdateData) GetDesc() string {
	if x != nil {
		return x.Desc
	}
	return ""
}

func (x *ResponseUpdateData) GetBooks() *Books {
	if x != nil {
		return x.Books
	}
	return nil
}

type InputBooks struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Namebook string `protobuf:"bytes,1,opt,name=namebook,proto3" json:"namebook,omitempty"`
	Author   string `protobuf:"bytes,2,opt,name=author,proto3" json:"author,omitempty"`
}

func (x *InputBooks) Reset() {
	*x = InputBooks{}
	if protoimpl.UnsafeEnabled {
		mi := &file_v1_books_books_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *InputBooks) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*InputBooks) ProtoMessage() {}

func (x *InputBooks) ProtoReflect() protoreflect.Message {
	mi := &file_v1_books_books_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use InputBooks.ProtoReflect.Descriptor instead.
func (*InputBooks) Descriptor() ([]byte, []int) {
	return file_v1_books_books_proto_rawDescGZIP(), []int{3}
}

func (x *InputBooks) GetNamebook() string {
	if x != nil {
		return x.Namebook
	}
	return ""
}

func (x *InputBooks) GetAuthor() string {
	if x != nil {
		return x.Author
	}
	return ""
}

type Books struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// string no       =3;
	Namebook string `protobuf:"bytes,1,opt,name=namebook,proto3" json:"namebook,omitempty"`
	Author   string `protobuf:"bytes,2,opt,name=author,proto3" json:"author,omitempty"`
}

func (x *Books) Reset() {
	*x = Books{}
	if protoimpl.UnsafeEnabled {
		mi := &file_v1_books_books_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Books) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Books) ProtoMessage() {}

func (x *Books) ProtoReflect() protoreflect.Message {
	mi := &file_v1_books_books_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Books.ProtoReflect.Descriptor instead.
func (*Books) Descriptor() ([]byte, []int) {
	return file_v1_books_books_proto_rawDescGZIP(), []int{4}
}

func (x *Books) GetNamebook() string {
	if x != nil {
		return x.Namebook
	}
	return ""
}

func (x *Books) GetAuthor() string {
	if x != nil {
		return x.Author
	}
	return ""
}

type UpdateBookData struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	No       int32  `protobuf:"varint,3,opt,name=no,proto3" json:"no,omitempty"`
	Namebook string `protobuf:"bytes,1,opt,name=namebook,proto3" json:"namebook,omitempty"`
	Author   string `protobuf:"bytes,2,opt,name=author,proto3" json:"author,omitempty"`
}

func (x *UpdateBookData) Reset() {
	*x = UpdateBookData{}
	if protoimpl.UnsafeEnabled {
		mi := &file_v1_books_books_proto_msgTypes[5]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *UpdateBookData) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*UpdateBookData) ProtoMessage() {}

func (x *UpdateBookData) ProtoReflect() protoreflect.Message {
	mi := &file_v1_books_books_proto_msgTypes[5]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use UpdateBookData.ProtoReflect.Descriptor instead.
func (*UpdateBookData) Descriptor() ([]byte, []int) {
	return file_v1_books_books_proto_rawDescGZIP(), []int{5}
}

func (x *UpdateBookData) GetNo() int32 {
	if x != nil {
		return x.No
	}
	return 0
}

func (x *UpdateBookData) GetNamebook() string {
	if x != nil {
		return x.Namebook
	}
	return ""
}

func (x *UpdateBookData) GetAuthor() string {
	if x != nil {
		return x.Author
	}
	return ""
}

var File_v1_books_books_proto protoreflect.FileDescriptor

var file_v1_books_books_proto_rawDesc = []byte{
	0x0a, 0x14, 0x76, 0x31, 0x2f, 0x62, 0x6f, 0x6f, 0x6b, 0x73, 0x2f, 0x62, 0x6f, 0x6f, 0x6b, 0x73,
	0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x15, 0x61, 0x70, 0x69, 0x2e, 0x6c, 0x61, 0x74, 0x69,
	0x68, 0x61, 0x6e, 0x2e, 0x76, 0x31, 0x2e, 0x68, 0x65, 0x61, 0x6c, 0x74, 0x68, 0x1a, 0x20, 0x6c,
	0x69, 0x62, 0x2f, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x61, 0x70, 0x69, 0x2f, 0x61, 0x6e,
	0x6e, 0x6f, 0x74, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a,
	0x1b, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66,
	0x2f, 0x65, 0x6d, 0x70, 0x74, 0x79, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0x7e, 0x0a, 0x08,
	0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x18, 0x0a, 0x07, 0x73, 0x75, 0x63, 0x63,
	0x65, 0x73, 0x73, 0x18, 0x01, 0x20, 0x01, 0x28, 0x08, 0x52, 0x07, 0x73, 0x75, 0x63, 0x63, 0x65,
	0x73, 0x73, 0x12, 0x12, 0x0a, 0x04, 0x63, 0x6f, 0x64, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09,
	0x52, 0x04, 0x63, 0x6f, 0x64, 0x65, 0x12, 0x12, 0x0a, 0x04, 0x64, 0x65, 0x73, 0x63, 0x18, 0x03,
	0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x64, 0x65, 0x73, 0x63, 0x12, 0x30, 0x0a, 0x04, 0x64, 0x61,
	0x74, 0x61, 0x18, 0x04, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x1c, 0x2e, 0x61, 0x70, 0x69, 0x2e, 0x6c,
	0x61, 0x74, 0x69, 0x68, 0x61, 0x6e, 0x2e, 0x76, 0x31, 0x2e, 0x68, 0x65, 0x61, 0x6c, 0x74, 0x68,
	0x2e, 0x42, 0x6f, 0x6f, 0x6b, 0x73, 0x52, 0x04, 0x64, 0x61, 0x74, 0x61, 0x22, 0x86, 0x01, 0x0a,
	0x0e, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x12,
	0x18, 0x0a, 0x07, 0x73, 0x75, 0x63, 0x63, 0x65, 0x73, 0x73, 0x18, 0x01, 0x20, 0x01, 0x28, 0x08,
	0x52, 0x07, 0x73, 0x75, 0x63, 0x63, 0x65, 0x73, 0x73, 0x12, 0x12, 0x0a, 0x04, 0x63, 0x6f, 0x64,
	0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x63, 0x6f, 0x64, 0x65, 0x12, 0x12, 0x0a,
	0x04, 0x64, 0x65, 0x73, 0x63, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x64, 0x65, 0x73,
	0x63, 0x12, 0x32, 0x0a, 0x05, 0x62, 0x6f, 0x6f, 0x6b, 0x73, 0x18, 0x05, 0x20, 0x01, 0x28, 0x0b,
	0x32, 0x1c, 0x2e, 0x61, 0x70, 0x69, 0x2e, 0x6c, 0x61, 0x74, 0x69, 0x68, 0x61, 0x6e, 0x2e, 0x76,
	0x31, 0x2e, 0x68, 0x65, 0x61, 0x6c, 0x74, 0x68, 0x2e, 0x42, 0x6f, 0x6f, 0x6b, 0x73, 0x52, 0x05,
	0x62, 0x6f, 0x6f, 0x6b, 0x73, 0x22, 0x8a, 0x01, 0x0a, 0x12, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e,
	0x73, 0x65, 0x55, 0x70, 0x64, 0x61, 0x74, 0x65, 0x44, 0x61, 0x74, 0x61, 0x12, 0x18, 0x0a, 0x07,
	0x73, 0x75, 0x63, 0x63, 0x65, 0x73, 0x73, 0x18, 0x01, 0x20, 0x01, 0x28, 0x08, 0x52, 0x07, 0x73,
	0x75, 0x63, 0x63, 0x65, 0x73, 0x73, 0x12, 0x12, 0x0a, 0x04, 0x63, 0x6f, 0x64, 0x65, 0x18, 0x02,
	0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x63, 0x6f, 0x64, 0x65, 0x12, 0x12, 0x0a, 0x04, 0x64, 0x65,
	0x73, 0x63, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x64, 0x65, 0x73, 0x63, 0x12, 0x32,
	0x0a, 0x05, 0x62, 0x6f, 0x6f, 0x6b, 0x73, 0x18, 0x04, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1c, 0x2e,
	0x61, 0x70, 0x69, 0x2e, 0x6c, 0x61, 0x74, 0x69, 0x68, 0x61, 0x6e, 0x2e, 0x76, 0x31, 0x2e, 0x68,
	0x65, 0x61, 0x6c, 0x74, 0x68, 0x2e, 0x42, 0x6f, 0x6f, 0x6b, 0x73, 0x52, 0x05, 0x62, 0x6f, 0x6f,
	0x6b, 0x73, 0x22, 0x40, 0x0a, 0x0a, 0x49, 0x6e, 0x70, 0x75, 0x74, 0x42, 0x6f, 0x6f, 0x6b, 0x73,
	0x12, 0x1a, 0x0a, 0x08, 0x6e, 0x61, 0x6d, 0x65, 0x62, 0x6f, 0x6f, 0x6b, 0x18, 0x01, 0x20, 0x01,
	0x28, 0x09, 0x52, 0x08, 0x6e, 0x61, 0x6d, 0x65, 0x62, 0x6f, 0x6f, 0x6b, 0x12, 0x16, 0x0a, 0x06,
	0x61, 0x75, 0x74, 0x68, 0x6f, 0x72, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x06, 0x61, 0x75,
	0x74, 0x68, 0x6f, 0x72, 0x22, 0x3b, 0x0a, 0x05, 0x42, 0x6f, 0x6f, 0x6b, 0x73, 0x12, 0x1a, 0x0a,
	0x08, 0x6e, 0x61, 0x6d, 0x65, 0x62, 0x6f, 0x6f, 0x6b, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52,
	0x08, 0x6e, 0x61, 0x6d, 0x65, 0x62, 0x6f, 0x6f, 0x6b, 0x12, 0x16, 0x0a, 0x06, 0x61, 0x75, 0x74,
	0x68, 0x6f, 0x72, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x06, 0x61, 0x75, 0x74, 0x68, 0x6f,
	0x72, 0x22, 0x54, 0x0a, 0x0e, 0x55, 0x70, 0x64, 0x61, 0x74, 0x65, 0x42, 0x6f, 0x6f, 0x6b, 0x44,
	0x61, 0x74, 0x61, 0x12, 0x0e, 0x0a, 0x02, 0x6e, 0x6f, 0x18, 0x03, 0x20, 0x01, 0x28, 0x05, 0x52,
	0x02, 0x6e, 0x6f, 0x12, 0x1a, 0x0a, 0x08, 0x6e, 0x61, 0x6d, 0x65, 0x62, 0x6f, 0x6f, 0x6b, 0x18,
	0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x6e, 0x61, 0x6d, 0x65, 0x62, 0x6f, 0x6f, 0x6b, 0x12,
	0x16, 0x0a, 0x06, 0x61, 0x75, 0x74, 0x68, 0x6f, 0x72, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52,
	0x06, 0x61, 0x75, 0x74, 0x68, 0x6f, 0x72, 0x32, 0xbc, 0x03, 0x0a, 0x0b, 0x42, 0x6f, 0x6f, 0x6b,
	0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x12, 0x64, 0x0a, 0x03, 0x47, 0x65, 0x74, 0x12, 0x16,
	0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66,
	0x2e, 0x45, 0x6d, 0x70, 0x74, 0x79, 0x1a, 0x1f, 0x2e, 0x61, 0x70, 0x69, 0x2e, 0x6c, 0x61, 0x74,
	0x69, 0x68, 0x61, 0x6e, 0x2e, 0x76, 0x31, 0x2e, 0x68, 0x65, 0x61, 0x6c, 0x74, 0x68, 0x2e, 0x52,
	0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x24, 0x82, 0xd3, 0xe4, 0x93, 0x02, 0x1e, 0x12,
	0x1c, 0x2f, 0x61, 0x70, 0x69, 0x2f, 0x67, 0x6f, 0x2d, 0x67, 0x72, 0x70, 0x63, 0x2d, 0x6c, 0x65,
	0x61, 0x72, 0x6e, 0x69, 0x6e, 0x67, 0x2f, 0x68, 0x65, 0x61, 0x6c, 0x74, 0x68, 0x12, 0x69, 0x0a,
	0x06, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x12, 0x1c, 0x2e, 0x61, 0x70, 0x69, 0x2e, 0x6c, 0x61,
	0x74, 0x69, 0x68, 0x61, 0x6e, 0x2e, 0x76, 0x31, 0x2e, 0x68, 0x65, 0x61, 0x6c, 0x74, 0x68, 0x2e,
	0x42, 0x6f, 0x6f, 0x6b, 0x73, 0x1a, 0x25, 0x2e, 0x61, 0x70, 0x69, 0x2e, 0x6c, 0x61, 0x74, 0x69,
	0x68, 0x61, 0x6e, 0x2e, 0x76, 0x31, 0x2e, 0x68, 0x65, 0x61, 0x6c, 0x74, 0x68, 0x2e, 0x52, 0x65,
	0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x22, 0x1a, 0x82, 0xd3,
	0xe4, 0x93, 0x02, 0x14, 0x22, 0x0f, 0x2f, 0x61, 0x70, 0x69, 0x2f, 0x63, 0x72, 0x65, 0x61, 0x74,
	0x65, 0x62, 0x6f, 0x6f, 0x6b, 0x3a, 0x01, 0x2a, 0x12, 0x5b, 0x0a, 0x06, 0x47, 0x65, 0x74, 0x41,
	0x6c, 0x6c, 0x12, 0x16, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x62, 0x75, 0x66, 0x2e, 0x45, 0x6d, 0x70, 0x74, 0x79, 0x1a, 0x1f, 0x2e, 0x61, 0x70, 0x69,
	0x2e, 0x6c, 0x61, 0x74, 0x69, 0x68, 0x61, 0x6e, 0x2e, 0x76, 0x31, 0x2e, 0x68, 0x65, 0x61, 0x6c,
	0x74, 0x68, 0x2e, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x18, 0x82, 0xd3, 0xe4,
	0x93, 0x02, 0x12, 0x12, 0x10, 0x2f, 0x61, 0x70, 0x69, 0x2f, 0x67, 0x65, 0x74, 0x2f, 0x61, 0x6c,
	0x6c, 0x64, 0x61, 0x74, 0x61, 0x12, 0x7f, 0x0a, 0x0a, 0x55, 0x70, 0x64, 0x61, 0x74, 0x65, 0x44,
	0x61, 0x74, 0x61, 0x12, 0x25, 0x2e, 0x61, 0x70, 0x69, 0x2e, 0x6c, 0x61, 0x74, 0x69, 0x68, 0x61,
	0x6e, 0x2e, 0x76, 0x31, 0x2e, 0x68, 0x65, 0x61, 0x6c, 0x74, 0x68, 0x2e, 0x55, 0x70, 0x64, 0x61,
	0x74, 0x65, 0x42, 0x6f, 0x6f, 0x6b, 0x44, 0x61, 0x74, 0x61, 0x1a, 0x29, 0x2e, 0x61, 0x70, 0x69,
	0x2e, 0x6c, 0x61, 0x74, 0x69, 0x68, 0x61, 0x6e, 0x2e, 0x76, 0x31, 0x2e, 0x68, 0x65, 0x61, 0x6c,
	0x74, 0x68, 0x2e, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x55, 0x70, 0x64, 0x61, 0x74,
	0x65, 0x44, 0x61, 0x74, 0x61, 0x22, 0x1f, 0x82, 0xd3, 0xe4, 0x93, 0x02, 0x19, 0x1a, 0x14, 0x2f,
	0x61, 0x70, 0x69, 0x2f, 0x75, 0x70, 0x64, 0x61, 0x74, 0x65, 0x62, 0x6f, 0x6f, 0x6b, 0x73, 0x64,
	0x61, 0x74, 0x61, 0x3a, 0x01, 0x2a, 0x42, 0x18, 0x5a, 0x16, 0x6c, 0x61, 0x74, 0x69, 0x68, 0x61,
	0x6e, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2f, 0x76, 0x31, 0x2f, 0x62, 0x6f, 0x6f, 0x6b, 0x73,
	0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_v1_books_books_proto_rawDescOnce sync.Once
	file_v1_books_books_proto_rawDescData = file_v1_books_books_proto_rawDesc
)

func file_v1_books_books_proto_rawDescGZIP() []byte {
	file_v1_books_books_proto_rawDescOnce.Do(func() {
		file_v1_books_books_proto_rawDescData = protoimpl.X.CompressGZIP(file_v1_books_books_proto_rawDescData)
	})
	return file_v1_books_books_proto_rawDescData
}

var file_v1_books_books_proto_msgTypes = make([]protoimpl.MessageInfo, 6)
var file_v1_books_books_proto_goTypes = []interface{}{
	(*Response)(nil),           // 0: api.latihan.v1.health.Response
	(*ResponseCreate)(nil),     // 1: api.latihan.v1.health.ResponseCreate
	(*ResponseUpdateData)(nil), // 2: api.latihan.v1.health.ResponseUpdateData
	(*InputBooks)(nil),         // 3: api.latihan.v1.health.InputBooks
	(*Books)(nil),              // 4: api.latihan.v1.health.Books
	(*UpdateBookData)(nil),     // 5: api.latihan.v1.health.UpdateBookData
	(*emptypb.Empty)(nil),      // 6: google.protobuf.Empty
}
var file_v1_books_books_proto_depIdxs = []int32{
	4, // 0: api.latihan.v1.health.Response.data:type_name -> api.latihan.v1.health.Books
	4, // 1: api.latihan.v1.health.ResponseCreate.books:type_name -> api.latihan.v1.health.Books
	4, // 2: api.latihan.v1.health.ResponseUpdateData.books:type_name -> api.latihan.v1.health.Books
	6, // 3: api.latihan.v1.health.BookService.Get:input_type -> google.protobuf.Empty
	4, // 4: api.latihan.v1.health.BookService.Create:input_type -> api.latihan.v1.health.Books
	6, // 5: api.latihan.v1.health.BookService.GetAll:input_type -> google.protobuf.Empty
	5, // 6: api.latihan.v1.health.BookService.UpdateData:input_type -> api.latihan.v1.health.UpdateBookData
	0, // 7: api.latihan.v1.health.BookService.Get:output_type -> api.latihan.v1.health.Response
	1, // 8: api.latihan.v1.health.BookService.Create:output_type -> api.latihan.v1.health.ResponseCreate
	0, // 9: api.latihan.v1.health.BookService.GetAll:output_type -> api.latihan.v1.health.Response
	2, // 10: api.latihan.v1.health.BookService.UpdateData:output_type -> api.latihan.v1.health.ResponseUpdateData
	7, // [7:11] is the sub-list for method output_type
	3, // [3:7] is the sub-list for method input_type
	3, // [3:3] is the sub-list for extension type_name
	3, // [3:3] is the sub-list for extension extendee
	0, // [0:3] is the sub-list for field type_name
}

func init() { file_v1_books_books_proto_init() }
func file_v1_books_books_proto_init() {
	if File_v1_books_books_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_v1_books_books_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Response); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_v1_books_books_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ResponseCreate); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_v1_books_books_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ResponseUpdateData); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_v1_books_books_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*InputBooks); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_v1_books_books_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Books); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_v1_books_books_proto_msgTypes[5].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*UpdateBookData); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
	}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_v1_books_books_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   6,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_v1_books_books_proto_goTypes,
		DependencyIndexes: file_v1_books_books_proto_depIdxs,
		MessageInfos:      file_v1_books_books_proto_msgTypes,
	}.Build()
	File_v1_books_books_proto = out.File
	file_v1_books_books_proto_rawDesc = nil
	file_v1_books_books_proto_goTypes = nil
	file_v1_books_books_proto_depIdxs = nil
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConnInterface

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion6

// BookServiceClient is the client API for BookService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type BookServiceClient interface {
	Get(ctx context.Context, in *emptypb.Empty, opts ...grpc.CallOption) (*Response, error)
	Create(ctx context.Context, in *Books, opts ...grpc.CallOption) (*ResponseCreate, error)
	GetAll(ctx context.Context, in *emptypb.Empty, opts ...grpc.CallOption) (*Response, error)
	UpdateData(ctx context.Context, in *UpdateBookData, opts ...grpc.CallOption) (*ResponseUpdateData, error)
}

type bookServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewBookServiceClient(cc grpc.ClientConnInterface) BookServiceClient {
	return &bookServiceClient{cc}
}

func (c *bookServiceClient) Get(ctx context.Context, in *emptypb.Empty, opts ...grpc.CallOption) (*Response, error) {
	out := new(Response)
	err := c.cc.Invoke(ctx, "/api.latihan.v1.health.BookService/Get", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *bookServiceClient) Create(ctx context.Context, in *Books, opts ...grpc.CallOption) (*ResponseCreate, error) {
	out := new(ResponseCreate)
	err := c.cc.Invoke(ctx, "/api.latihan.v1.health.BookService/Create", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *bookServiceClient) GetAll(ctx context.Context, in *emptypb.Empty, opts ...grpc.CallOption) (*Response, error) {
	out := new(Response)
	err := c.cc.Invoke(ctx, "/api.latihan.v1.health.BookService/GetAll", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *bookServiceClient) UpdateData(ctx context.Context, in *UpdateBookData, opts ...grpc.CallOption) (*ResponseUpdateData, error) {
	out := new(ResponseUpdateData)
	err := c.cc.Invoke(ctx, "/api.latihan.v1.health.BookService/UpdateData", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// BookServiceServer is the server API for BookService service.
type BookServiceServer interface {
	Get(context.Context, *emptypb.Empty) (*Response, error)
	Create(context.Context, *Books) (*ResponseCreate, error)
	GetAll(context.Context, *emptypb.Empty) (*Response, error)
	UpdateData(context.Context, *UpdateBookData) (*ResponseUpdateData, error)
}

// UnimplementedBookServiceServer can be embedded to have forward compatible implementations.
type UnimplementedBookServiceServer struct {
}

func (*UnimplementedBookServiceServer) Get(context.Context, *emptypb.Empty) (*Response, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Get not implemented")
}
func (*UnimplementedBookServiceServer) Create(context.Context, *Books) (*ResponseCreate, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Create not implemented")
}
func (*UnimplementedBookServiceServer) GetAll(context.Context, *emptypb.Empty) (*Response, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetAll not implemented")
}
func (*UnimplementedBookServiceServer) UpdateData(context.Context, *UpdateBookData) (*ResponseUpdateData, error) {
	return nil, status.Errorf(codes.Unimplemented, "method UpdateData not implemented")
}

func RegisterBookServiceServer(s *grpc.Server, srv BookServiceServer) {
	s.RegisterService(&_BookService_serviceDesc, srv)
}

func _BookService_Get_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(emptypb.Empty)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(BookServiceServer).Get(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/api.latihan.v1.health.BookService/Get",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(BookServiceServer).Get(ctx, req.(*emptypb.Empty))
	}
	return interceptor(ctx, in, info, handler)
}

func _BookService_Create_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(Books)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(BookServiceServer).Create(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/api.latihan.v1.health.BookService/Create",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(BookServiceServer).Create(ctx, req.(*Books))
	}
	return interceptor(ctx, in, info, handler)
}

func _BookService_GetAll_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(emptypb.Empty)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(BookServiceServer).GetAll(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/api.latihan.v1.health.BookService/GetAll",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(BookServiceServer).GetAll(ctx, req.(*emptypb.Empty))
	}
	return interceptor(ctx, in, info, handler)
}

func _BookService_UpdateData_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(UpdateBookData)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(BookServiceServer).UpdateData(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/api.latihan.v1.health.BookService/UpdateData",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(BookServiceServer).UpdateData(ctx, req.(*UpdateBookData))
	}
	return interceptor(ctx, in, info, handler)
}

var _BookService_serviceDesc = grpc.ServiceDesc{
	ServiceName: "api.latihan.v1.health.BookService",
	HandlerType: (*BookServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "Get",
			Handler:    _BookService_Get_Handler,
		},
		{
			MethodName: "Create",
			Handler:    _BookService_Create_Handler,
		},
		{
			MethodName: "GetAll",
			Handler:    _BookService_GetAll_Handler,
		},
		{
			MethodName: "UpdateData",
			Handler:    _BookService_UpdateData_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "v1/books/books.proto",
}
