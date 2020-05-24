// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.23.0
// 	protoc        v3.11.4
// source: rpc/dss/discount.proto

package dss

import (
	proto "github.com/golang/protobuf/proto"
	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
	reflect "reflect"
	sync "sync"
)

const (
	// Verify that this generated code is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(20 - protoimpl.MinVersion)
	// Verify that runtime/protoimpl is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(protoimpl.MaxVersion - 20)
)

// This is a compile-time assertion that a sufficiently up-to-date version
// of the legacy proto package is being used.
const _ = proto.ProtoPackageIsVersion4

type DiscountAmount struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// Types that are assignable to Amount:
	//	*DiscountAmount_Dollar
	//	*DiscountAmount_Percent
	Amount isDiscountAmount_Amount `protobuf_oneof:"amount"`
}

func (x *DiscountAmount) Reset() {
	*x = DiscountAmount{}
	if protoimpl.UnsafeEnabled {
		mi := &file_rpc_dss_discount_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *DiscountAmount) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*DiscountAmount) ProtoMessage() {}

func (x *DiscountAmount) ProtoReflect() protoreflect.Message {
	mi := &file_rpc_dss_discount_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use DiscountAmount.ProtoReflect.Descriptor instead.
func (*DiscountAmount) Descriptor() ([]byte, []int) {
	return file_rpc_dss_discount_proto_rawDescGZIP(), []int{0}
}

func (m *DiscountAmount) GetAmount() isDiscountAmount_Amount {
	if m != nil {
		return m.Amount
	}
	return nil
}

func (x *DiscountAmount) GetDollar() int64 {
	if x, ok := x.GetAmount().(*DiscountAmount_Dollar); ok {
		return x.Dollar
	}
	return 0
}

func (x *DiscountAmount) GetPercent() string {
	if x, ok := x.GetAmount().(*DiscountAmount_Percent); ok {
		return x.Percent
	}
	return ""
}

type isDiscountAmount_Amount interface {
	isDiscountAmount_Amount()
}

type DiscountAmount_Dollar struct {
	Dollar int64 `protobuf:"varint,8,opt,name=dollar,proto3,oneof"`
}

type DiscountAmount_Percent struct {
	Percent string `protobuf:"bytes,9,opt,name=percent,proto3,oneof"`
}

func (*DiscountAmount_Dollar) isDiscountAmount_Amount() {}

func (*DiscountAmount_Percent) isDiscountAmount_Amount() {}

type SingleDiscount struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Name   string          `protobuf:"bytes,1,opt,name=name,proto3" json:"name,omitempty"`
	Amount *DiscountAmount `protobuf:"bytes,2,opt,name=amount,proto3" json:"amount,omitempty"`
}

func (x *SingleDiscount) Reset() {
	*x = SingleDiscount{}
	if protoimpl.UnsafeEnabled {
		mi := &file_rpc_dss_discount_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *SingleDiscount) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*SingleDiscount) ProtoMessage() {}

func (x *SingleDiscount) ProtoReflect() protoreflect.Message {
	mi := &file_rpc_dss_discount_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use SingleDiscount.ProtoReflect.Descriptor instead.
func (*SingleDiscount) Descriptor() ([]byte, []int) {
	return file_rpc_dss_discount_proto_rawDescGZIP(), []int{1}
}

func (x *SingleDiscount) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

func (x *SingleDiscount) GetAmount() *DiscountAmount {
	if x != nil {
		return x.Amount
	}
	return nil
}

type DiscountBundle struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Code      string            `protobuf:"bytes,1,opt,name=code,proto3" json:"code,omitempty"`
	Discounts []*SingleDiscount `protobuf:"bytes,2,rep,name=discounts,proto3" json:"discounts,omitempty"`
}

func (x *DiscountBundle) Reset() {
	*x = DiscountBundle{}
	if protoimpl.UnsafeEnabled {
		mi := &file_rpc_dss_discount_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *DiscountBundle) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*DiscountBundle) ProtoMessage() {}

func (x *DiscountBundle) ProtoReflect() protoreflect.Message {
	mi := &file_rpc_dss_discount_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use DiscountBundle.ProtoReflect.Descriptor instead.
func (*DiscountBundle) Descriptor() ([]byte, []int) {
	return file_rpc_dss_discount_proto_rawDescGZIP(), []int{2}
}

func (x *DiscountBundle) GetCode() string {
	if x != nil {
		return x.Code
	}
	return ""
}

func (x *DiscountBundle) GetDiscounts() []*SingleDiscount {
	if x != nil {
		return x.Discounts
	}
	return nil
}

type DiscountAddReq struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Bundle *DiscountBundle `protobuf:"bytes,1,opt,name=bundle,proto3" json:"bundle,omitempty"`
}

func (x *DiscountAddReq) Reset() {
	*x = DiscountAddReq{}
	if protoimpl.UnsafeEnabled {
		mi := &file_rpc_dss_discount_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *DiscountAddReq) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*DiscountAddReq) ProtoMessage() {}

func (x *DiscountAddReq) ProtoReflect() protoreflect.Message {
	mi := &file_rpc_dss_discount_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use DiscountAddReq.ProtoReflect.Descriptor instead.
func (*DiscountAddReq) Descriptor() ([]byte, []int) {
	return file_rpc_dss_discount_proto_rawDescGZIP(), []int{3}
}

func (x *DiscountAddReq) GetBundle() *DiscountBundle {
	if x != nil {
		return x.Bundle
	}
	return nil
}

type DiscountAddRes struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields
}

func (x *DiscountAddRes) Reset() {
	*x = DiscountAddRes{}
	if protoimpl.UnsafeEnabled {
		mi := &file_rpc_dss_discount_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *DiscountAddRes) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*DiscountAddRes) ProtoMessage() {}

func (x *DiscountAddRes) ProtoReflect() protoreflect.Message {
	mi := &file_rpc_dss_discount_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use DiscountAddRes.ProtoReflect.Descriptor instead.
func (*DiscountAddRes) Descriptor() ([]byte, []int) {
	return file_rpc_dss_discount_proto_rawDescGZIP(), []int{4}
}

type DiscountGetReq struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Code string `protobuf:"bytes,1,opt,name=code,proto3" json:"code,omitempty"`
}

func (x *DiscountGetReq) Reset() {
	*x = DiscountGetReq{}
	if protoimpl.UnsafeEnabled {
		mi := &file_rpc_dss_discount_proto_msgTypes[5]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *DiscountGetReq) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*DiscountGetReq) ProtoMessage() {}

func (x *DiscountGetReq) ProtoReflect() protoreflect.Message {
	mi := &file_rpc_dss_discount_proto_msgTypes[5]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use DiscountGetReq.ProtoReflect.Descriptor instead.
func (*DiscountGetReq) Descriptor() ([]byte, []int) {
	return file_rpc_dss_discount_proto_rawDescGZIP(), []int{5}
}

func (x *DiscountGetReq) GetCode() string {
	if x != nil {
		return x.Code
	}
	return ""
}

type DiscountGetRes struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Bundle *DiscountBundle `protobuf:"bytes,1,opt,name=bundle,proto3" json:"bundle,omitempty"`
}

func (x *DiscountGetRes) Reset() {
	*x = DiscountGetRes{}
	if protoimpl.UnsafeEnabled {
		mi := &file_rpc_dss_discount_proto_msgTypes[6]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *DiscountGetRes) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*DiscountGetRes) ProtoMessage() {}

func (x *DiscountGetRes) ProtoReflect() protoreflect.Message {
	mi := &file_rpc_dss_discount_proto_msgTypes[6]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use DiscountGetRes.ProtoReflect.Descriptor instead.
func (*DiscountGetRes) Descriptor() ([]byte, []int) {
	return file_rpc_dss_discount_proto_rawDescGZIP(), []int{6}
}

func (x *DiscountGetRes) GetBundle() *DiscountBundle {
	if x != nil {
		return x.Bundle
	}
	return nil
}

var File_rpc_dss_discount_proto protoreflect.FileDescriptor

var file_rpc_dss_discount_proto_rawDesc = []byte{
	0x0a, 0x16, 0x72, 0x70, 0x63, 0x2f, 0x64, 0x73, 0x73, 0x2f, 0x64, 0x69, 0x73, 0x63, 0x6f, 0x75,
	0x6e, 0x74, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x03, 0x64, 0x73, 0x73, 0x22, 0x50, 0x0a,
	0x0e, 0x44, 0x69, 0x73, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x41, 0x6d, 0x6f, 0x75, 0x6e, 0x74, 0x12,
	0x18, 0x0a, 0x06, 0x64, 0x6f, 0x6c, 0x6c, 0x61, 0x72, 0x18, 0x08, 0x20, 0x01, 0x28, 0x03, 0x48,
	0x00, 0x52, 0x06, 0x64, 0x6f, 0x6c, 0x6c, 0x61, 0x72, 0x12, 0x1a, 0x0a, 0x07, 0x70, 0x65, 0x72,
	0x63, 0x65, 0x6e, 0x74, 0x18, 0x09, 0x20, 0x01, 0x28, 0x09, 0x48, 0x00, 0x52, 0x07, 0x70, 0x65,
	0x72, 0x63, 0x65, 0x6e, 0x74, 0x42, 0x08, 0x0a, 0x06, 0x61, 0x6d, 0x6f, 0x75, 0x6e, 0x74, 0x22,
	0x51, 0x0a, 0x0e, 0x53, 0x69, 0x6e, 0x67, 0x6c, 0x65, 0x44, 0x69, 0x73, 0x63, 0x6f, 0x75, 0x6e,
	0x74, 0x12, 0x12, 0x0a, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52,
	0x04, 0x6e, 0x61, 0x6d, 0x65, 0x12, 0x2b, 0x0a, 0x06, 0x61, 0x6d, 0x6f, 0x75, 0x6e, 0x74, 0x18,
	0x02, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x13, 0x2e, 0x64, 0x73, 0x73, 0x2e, 0x44, 0x69, 0x73, 0x63,
	0x6f, 0x75, 0x6e, 0x74, 0x41, 0x6d, 0x6f, 0x75, 0x6e, 0x74, 0x52, 0x06, 0x61, 0x6d, 0x6f, 0x75,
	0x6e, 0x74, 0x22, 0x57, 0x0a, 0x0e, 0x44, 0x69, 0x73, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x42, 0x75,
	0x6e, 0x64, 0x6c, 0x65, 0x12, 0x12, 0x0a, 0x04, 0x63, 0x6f, 0x64, 0x65, 0x18, 0x01, 0x20, 0x01,
	0x28, 0x09, 0x52, 0x04, 0x63, 0x6f, 0x64, 0x65, 0x12, 0x31, 0x0a, 0x09, 0x64, 0x69, 0x73, 0x63,
	0x6f, 0x75, 0x6e, 0x74, 0x73, 0x18, 0x02, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x13, 0x2e, 0x64, 0x73,
	0x73, 0x2e, 0x53, 0x69, 0x6e, 0x67, 0x6c, 0x65, 0x44, 0x69, 0x73, 0x63, 0x6f, 0x75, 0x6e, 0x74,
	0x52, 0x09, 0x64, 0x69, 0x73, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x73, 0x22, 0x3d, 0x0a, 0x0e, 0x44,
	0x69, 0x73, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x41, 0x64, 0x64, 0x52, 0x65, 0x71, 0x12, 0x2b, 0x0a,
	0x06, 0x62, 0x75, 0x6e, 0x64, 0x6c, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x13, 0x2e,
	0x64, 0x73, 0x73, 0x2e, 0x44, 0x69, 0x73, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x42, 0x75, 0x6e, 0x64,
	0x6c, 0x65, 0x52, 0x06, 0x62, 0x75, 0x6e, 0x64, 0x6c, 0x65, 0x22, 0x10, 0x0a, 0x0e, 0x44, 0x69,
	0x73, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x41, 0x64, 0x64, 0x52, 0x65, 0x73, 0x22, 0x24, 0x0a, 0x0e,
	0x44, 0x69, 0x73, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x47, 0x65, 0x74, 0x52, 0x65, 0x71, 0x12, 0x12,
	0x0a, 0x04, 0x63, 0x6f, 0x64, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x63, 0x6f,
	0x64, 0x65, 0x22, 0x3d, 0x0a, 0x0e, 0x44, 0x69, 0x73, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x47, 0x65,
	0x74, 0x52, 0x65, 0x73, 0x12, 0x2b, 0x0a, 0x06, 0x62, 0x75, 0x6e, 0x64, 0x6c, 0x65, 0x18, 0x01,
	0x20, 0x01, 0x28, 0x0b, 0x32, 0x13, 0x2e, 0x64, 0x73, 0x73, 0x2e, 0x44, 0x69, 0x73, 0x63, 0x6f,
	0x75, 0x6e, 0x74, 0x42, 0x75, 0x6e, 0x64, 0x6c, 0x65, 0x52, 0x06, 0x62, 0x75, 0x6e, 0x64, 0x6c,
	0x65, 0x32, 0x6c, 0x0a, 0x08, 0x44, 0x69, 0x73, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x12, 0x2f, 0x0a,
	0x03, 0x41, 0x64, 0x64, 0x12, 0x13, 0x2e, 0x64, 0x73, 0x73, 0x2e, 0x44, 0x69, 0x73, 0x63, 0x6f,
	0x75, 0x6e, 0x74, 0x41, 0x64, 0x64, 0x52, 0x65, 0x71, 0x1a, 0x13, 0x2e, 0x64, 0x73, 0x73, 0x2e,
	0x44, 0x69, 0x73, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x41, 0x64, 0x64, 0x52, 0x65, 0x73, 0x12, 0x2f,
	0x0a, 0x03, 0x47, 0x65, 0x74, 0x12, 0x13, 0x2e, 0x64, 0x73, 0x73, 0x2e, 0x44, 0x69, 0x73, 0x63,
	0x6f, 0x75, 0x6e, 0x74, 0x47, 0x65, 0x74, 0x52, 0x65, 0x71, 0x1a, 0x13, 0x2e, 0x64, 0x73, 0x73,
	0x2e, 0x44, 0x69, 0x73, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x47, 0x65, 0x74, 0x52, 0x65, 0x73, 0x42,
	0x09, 0x5a, 0x07, 0x72, 0x70, 0x63, 0x2f, 0x64, 0x73, 0x73, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x33,
}

var (
	file_rpc_dss_discount_proto_rawDescOnce sync.Once
	file_rpc_dss_discount_proto_rawDescData = file_rpc_dss_discount_proto_rawDesc
)

func file_rpc_dss_discount_proto_rawDescGZIP() []byte {
	file_rpc_dss_discount_proto_rawDescOnce.Do(func() {
		file_rpc_dss_discount_proto_rawDescData = protoimpl.X.CompressGZIP(file_rpc_dss_discount_proto_rawDescData)
	})
	return file_rpc_dss_discount_proto_rawDescData
}

var file_rpc_dss_discount_proto_msgTypes = make([]protoimpl.MessageInfo, 7)
var file_rpc_dss_discount_proto_goTypes = []interface{}{
	(*DiscountAmount)(nil), // 0: dss.DiscountAmount
	(*SingleDiscount)(nil), // 1: dss.SingleDiscount
	(*DiscountBundle)(nil), // 2: dss.DiscountBundle
	(*DiscountAddReq)(nil), // 3: dss.DiscountAddReq
	(*DiscountAddRes)(nil), // 4: dss.DiscountAddRes
	(*DiscountGetReq)(nil), // 5: dss.DiscountGetReq
	(*DiscountGetRes)(nil), // 6: dss.DiscountGetRes
}
var file_rpc_dss_discount_proto_depIdxs = []int32{
	0, // 0: dss.SingleDiscount.amount:type_name -> dss.DiscountAmount
	1, // 1: dss.DiscountBundle.discounts:type_name -> dss.SingleDiscount
	2, // 2: dss.DiscountAddReq.bundle:type_name -> dss.DiscountBundle
	2, // 3: dss.DiscountGetRes.bundle:type_name -> dss.DiscountBundle
	3, // 4: dss.Discount.Add:input_type -> dss.DiscountAddReq
	5, // 5: dss.Discount.Get:input_type -> dss.DiscountGetReq
	4, // 6: dss.Discount.Add:output_type -> dss.DiscountAddRes
	6, // 7: dss.Discount.Get:output_type -> dss.DiscountGetRes
	6, // [6:8] is the sub-list for method output_type
	4, // [4:6] is the sub-list for method input_type
	4, // [4:4] is the sub-list for extension type_name
	4, // [4:4] is the sub-list for extension extendee
	0, // [0:4] is the sub-list for field type_name
}

func init() { file_rpc_dss_discount_proto_init() }
func file_rpc_dss_discount_proto_init() {
	if File_rpc_dss_discount_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_rpc_dss_discount_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*DiscountAmount); i {
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
		file_rpc_dss_discount_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*SingleDiscount); i {
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
		file_rpc_dss_discount_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*DiscountBundle); i {
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
		file_rpc_dss_discount_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*DiscountAddReq); i {
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
		file_rpc_dss_discount_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*DiscountAddRes); i {
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
		file_rpc_dss_discount_proto_msgTypes[5].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*DiscountGetReq); i {
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
		file_rpc_dss_discount_proto_msgTypes[6].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*DiscountGetRes); i {
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
	file_rpc_dss_discount_proto_msgTypes[0].OneofWrappers = []interface{}{
		(*DiscountAmount_Dollar)(nil),
		(*DiscountAmount_Percent)(nil),
	}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_rpc_dss_discount_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   7,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_rpc_dss_discount_proto_goTypes,
		DependencyIndexes: file_rpc_dss_discount_proto_depIdxs,
		MessageInfos:      file_rpc_dss_discount_proto_msgTypes,
	}.Build()
	File_rpc_dss_discount_proto = out.File
	file_rpc_dss_discount_proto_rawDesc = nil
	file_rpc_dss_discount_proto_goTypes = nil
	file_rpc_dss_discount_proto_depIdxs = nil
}
