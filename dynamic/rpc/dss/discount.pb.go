// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.26.0
// 	protoc        v3.16.0
// source: discount.proto

package dss

import (
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

type PurchaseItem int32

const (
	PurchaseItem_FullWeekendPassPurchaseItem PurchaseItem = 0
	PurchaseItem_DanceOnlyPassPurchaseItem   PurchaseItem = 1
	PurchaseItem_MixAndMatchPurchaseItem     PurchaseItem = 2
	PurchaseItem_SoloJazzPurchaseItem        PurchaseItem = 3
	PurchaseItem_TeamCompetitionPurchaseItem PurchaseItem = 4
	PurchaseItem_TShirtPurchaseItem          PurchaseItem = 5
)

// Enum value maps for PurchaseItem.
var (
	PurchaseItem_name = map[int32]string{
		0: "FullWeekendPassPurchaseItem",
		1: "DanceOnlyPassPurchaseItem",
		2: "MixAndMatchPurchaseItem",
		3: "SoloJazzPurchaseItem",
		4: "TeamCompetitionPurchaseItem",
		5: "TShirtPurchaseItem",
	}
	PurchaseItem_value = map[string]int32{
		"FullWeekendPassPurchaseItem": 0,
		"DanceOnlyPassPurchaseItem":   1,
		"MixAndMatchPurchaseItem":     2,
		"SoloJazzPurchaseItem":        3,
		"TeamCompetitionPurchaseItem": 4,
		"TShirtPurchaseItem":          5,
	}
)

func (x PurchaseItem) Enum() *PurchaseItem {
	p := new(PurchaseItem)
	*p = x
	return p
}

func (x PurchaseItem) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (PurchaseItem) Descriptor() protoreflect.EnumDescriptor {
	return file_discount_proto_enumTypes[0].Descriptor()
}

func (PurchaseItem) Type() protoreflect.EnumType {
	return &file_discount_proto_enumTypes[0]
}

func (x PurchaseItem) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use PurchaseItem.Descriptor instead.
func (PurchaseItem) EnumDescriptor() ([]byte, []int) {
	return file_discount_proto_rawDescGZIP(), []int{0}
}

type DiscountAmount struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// Types that are assignable to Amount:
	//	*DiscountAmount_Dollar
	//	*DiscountAmount_Percent
	//	*DiscountAmount_SquareNotFound
	Amount isDiscountAmount_Amount `protobuf_oneof:"amount"`
}

func (x *DiscountAmount) Reset() {
	*x = DiscountAmount{}
	if protoimpl.UnsafeEnabled {
		mi := &file_discount_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *DiscountAmount) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*DiscountAmount) ProtoMessage() {}

func (x *DiscountAmount) ProtoReflect() protoreflect.Message {
	mi := &file_discount_proto_msgTypes[0]
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
	return file_discount_proto_rawDescGZIP(), []int{0}
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

func (x *DiscountAmount) GetSquareNotFound() *emptypb.Empty {
	if x, ok := x.GetAmount().(*DiscountAmount_SquareNotFound); ok {
		return x.SquareNotFound
	}
	return nil
}

type isDiscountAmount_Amount interface {
	isDiscountAmount_Amount()
}

type DiscountAmount_Dollar struct {
	Dollar int64 `protobuf:"varint,1,opt,name=dollar,proto3,oneof"`
}

type DiscountAmount_Percent struct {
	Percent string `protobuf:"bytes,2,opt,name=percent,proto3,oneof"`
}

type DiscountAmount_SquareNotFound struct {
	SquareNotFound *emptypb.Empty `protobuf:"bytes,3,opt,name=squareNotFound,proto3,oneof"`
}

func (*DiscountAmount_Dollar) isDiscountAmount_Amount() {}

func (*DiscountAmount_Percent) isDiscountAmount_Amount() {}

func (*DiscountAmount_SquareNotFound) isDiscountAmount_Amount() {}

type SingleDiscount struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Name      string          `protobuf:"bytes,1,opt,name=name,proto3" json:"name,omitempty"`
	Amount    *DiscountAmount `protobuf:"bytes,2,opt,name=amount,proto3" json:"amount,omitempty"`
	AppliedTo PurchaseItem    `protobuf:"varint,3,opt,name=applied_to,json=appliedTo,proto3,enum=dss.PurchaseItem" json:"applied_to,omitempty"`
}

func (x *SingleDiscount) Reset() {
	*x = SingleDiscount{}
	if protoimpl.UnsafeEnabled {
		mi := &file_discount_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *SingleDiscount) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*SingleDiscount) ProtoMessage() {}

func (x *SingleDiscount) ProtoReflect() protoreflect.Message {
	mi := &file_discount_proto_msgTypes[1]
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
	return file_discount_proto_rawDescGZIP(), []int{1}
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

func (x *SingleDiscount) GetAppliedTo() PurchaseItem {
	if x != nil {
		return x.AppliedTo
	}
	return PurchaseItem_FullWeekendPassPurchaseItem
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
		mi := &file_discount_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *DiscountBundle) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*DiscountBundle) ProtoMessage() {}

func (x *DiscountBundle) ProtoReflect() protoreflect.Message {
	mi := &file_discount_proto_msgTypes[2]
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
	return file_discount_proto_rawDescGZIP(), []int{2}
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
		mi := &file_discount_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *DiscountAddReq) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*DiscountAddReq) ProtoMessage() {}

func (x *DiscountAddReq) ProtoReflect() protoreflect.Message {
	mi := &file_discount_proto_msgTypes[3]
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
	return file_discount_proto_rawDescGZIP(), []int{3}
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
		mi := &file_discount_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *DiscountAddRes) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*DiscountAddRes) ProtoMessage() {}

func (x *DiscountAddRes) ProtoReflect() protoreflect.Message {
	mi := &file_discount_proto_msgTypes[4]
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
	return file_discount_proto_rawDescGZIP(), []int{4}
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
		mi := &file_discount_proto_msgTypes[5]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *DiscountGetReq) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*DiscountGetReq) ProtoMessage() {}

func (x *DiscountGetReq) ProtoReflect() protoreflect.Message {
	mi := &file_discount_proto_msgTypes[5]
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
	return file_discount_proto_rawDescGZIP(), []int{5}
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
		mi := &file_discount_proto_msgTypes[6]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *DiscountGetRes) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*DiscountGetRes) ProtoMessage() {}

func (x *DiscountGetRes) ProtoReflect() protoreflect.Message {
	mi := &file_discount_proto_msgTypes[6]
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
	return file_discount_proto_rawDescGZIP(), []int{6}
}

func (x *DiscountGetRes) GetBundle() *DiscountBundle {
	if x != nil {
		return x.Bundle
	}
	return nil
}

type DiscountListReq struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields
}

func (x *DiscountListReq) Reset() {
	*x = DiscountListReq{}
	if protoimpl.UnsafeEnabled {
		mi := &file_discount_proto_msgTypes[7]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *DiscountListReq) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*DiscountListReq) ProtoMessage() {}

func (x *DiscountListReq) ProtoReflect() protoreflect.Message {
	mi := &file_discount_proto_msgTypes[7]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use DiscountListReq.ProtoReflect.Descriptor instead.
func (*DiscountListReq) Descriptor() ([]byte, []int) {
	return file_discount_proto_rawDescGZIP(), []int{7}
}

type DiscountListRes struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Bundles []*DiscountBundle `protobuf:"bytes,1,rep,name=bundles,proto3" json:"bundles,omitempty"`
}

func (x *DiscountListRes) Reset() {
	*x = DiscountListRes{}
	if protoimpl.UnsafeEnabled {
		mi := &file_discount_proto_msgTypes[8]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *DiscountListRes) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*DiscountListRes) ProtoMessage() {}

func (x *DiscountListRes) ProtoReflect() protoreflect.Message {
	mi := &file_discount_proto_msgTypes[8]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use DiscountListRes.ProtoReflect.Descriptor instead.
func (*DiscountListRes) Descriptor() ([]byte, []int) {
	return file_discount_proto_rawDescGZIP(), []int{8}
}

func (x *DiscountListRes) GetBundles() []*DiscountBundle {
	if x != nil {
		return x.Bundles
	}
	return nil
}

type DiscountUpdateReq struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	OldCode string          `protobuf:"bytes,1,opt,name=oldCode,proto3" json:"oldCode,omitempty"`
	Bundle  *DiscountBundle `protobuf:"bytes,2,opt,name=bundle,proto3" json:"bundle,omitempty"`
}

func (x *DiscountUpdateReq) Reset() {
	*x = DiscountUpdateReq{}
	if protoimpl.UnsafeEnabled {
		mi := &file_discount_proto_msgTypes[9]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *DiscountUpdateReq) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*DiscountUpdateReq) ProtoMessage() {}

func (x *DiscountUpdateReq) ProtoReflect() protoreflect.Message {
	mi := &file_discount_proto_msgTypes[9]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use DiscountUpdateReq.ProtoReflect.Descriptor instead.
func (*DiscountUpdateReq) Descriptor() ([]byte, []int) {
	return file_discount_proto_rawDescGZIP(), []int{9}
}

func (x *DiscountUpdateReq) GetOldCode() string {
	if x != nil {
		return x.OldCode
	}
	return ""
}

func (x *DiscountUpdateReq) GetBundle() *DiscountBundle {
	if x != nil {
		return x.Bundle
	}
	return nil
}

type DiscountUpdateRes struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields
}

func (x *DiscountUpdateRes) Reset() {
	*x = DiscountUpdateRes{}
	if protoimpl.UnsafeEnabled {
		mi := &file_discount_proto_msgTypes[10]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *DiscountUpdateRes) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*DiscountUpdateRes) ProtoMessage() {}

func (x *DiscountUpdateRes) ProtoReflect() protoreflect.Message {
	mi := &file_discount_proto_msgTypes[10]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use DiscountUpdateRes.ProtoReflect.Descriptor instead.
func (*DiscountUpdateRes) Descriptor() ([]byte, []int) {
	return file_discount_proto_rawDescGZIP(), []int{10}
}

type DiscountDeleteReq struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Code string `protobuf:"bytes,1,opt,name=code,proto3" json:"code,omitempty"`
}

func (x *DiscountDeleteReq) Reset() {
	*x = DiscountDeleteReq{}
	if protoimpl.UnsafeEnabled {
		mi := &file_discount_proto_msgTypes[11]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *DiscountDeleteReq) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*DiscountDeleteReq) ProtoMessage() {}

func (x *DiscountDeleteReq) ProtoReflect() protoreflect.Message {
	mi := &file_discount_proto_msgTypes[11]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use DiscountDeleteReq.ProtoReflect.Descriptor instead.
func (*DiscountDeleteReq) Descriptor() ([]byte, []int) {
	return file_discount_proto_rawDescGZIP(), []int{11}
}

func (x *DiscountDeleteReq) GetCode() string {
	if x != nil {
		return x.Code
	}
	return ""
}

type DiscountDeleteRes struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields
}

func (x *DiscountDeleteRes) Reset() {
	*x = DiscountDeleteRes{}
	if protoimpl.UnsafeEnabled {
		mi := &file_discount_proto_msgTypes[12]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *DiscountDeleteRes) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*DiscountDeleteRes) ProtoMessage() {}

func (x *DiscountDeleteRes) ProtoReflect() protoreflect.Message {
	mi := &file_discount_proto_msgTypes[12]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use DiscountDeleteRes.ProtoReflect.Descriptor instead.
func (*DiscountDeleteRes) Descriptor() ([]byte, []int) {
	return file_discount_proto_rawDescGZIP(), []int{12}
}

var File_discount_proto protoreflect.FileDescriptor

var file_discount_proto_rawDesc = []byte{
	0x0a, 0x0e, 0x64, 0x69, 0x73, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x12, 0x03, 0x64, 0x73, 0x73, 0x1a, 0x1b, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2f, 0x65, 0x6d, 0x70, 0x74, 0x79, 0x2e, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x22, 0x92, 0x01, 0x0a, 0x0e, 0x44, 0x69, 0x73, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x41,
	0x6d, 0x6f, 0x75, 0x6e, 0x74, 0x12, 0x18, 0x0a, 0x06, 0x64, 0x6f, 0x6c, 0x6c, 0x61, 0x72, 0x18,
	0x01, 0x20, 0x01, 0x28, 0x03, 0x48, 0x00, 0x52, 0x06, 0x64, 0x6f, 0x6c, 0x6c, 0x61, 0x72, 0x12,
	0x1a, 0x0a, 0x07, 0x70, 0x65, 0x72, 0x63, 0x65, 0x6e, 0x74, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09,
	0x48, 0x00, 0x52, 0x07, 0x70, 0x65, 0x72, 0x63, 0x65, 0x6e, 0x74, 0x12, 0x40, 0x0a, 0x0e, 0x73,
	0x71, 0x75, 0x61, 0x72, 0x65, 0x4e, 0x6f, 0x74, 0x46, 0x6f, 0x75, 0x6e, 0x64, 0x18, 0x03, 0x20,
	0x01, 0x28, 0x0b, 0x32, 0x16, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x45, 0x6d, 0x70, 0x74, 0x79, 0x48, 0x00, 0x52, 0x0e, 0x73,
	0x71, 0x75, 0x61, 0x72, 0x65, 0x4e, 0x6f, 0x74, 0x46, 0x6f, 0x75, 0x6e, 0x64, 0x42, 0x08, 0x0a,
	0x06, 0x61, 0x6d, 0x6f, 0x75, 0x6e, 0x74, 0x22, 0x83, 0x01, 0x0a, 0x0e, 0x53, 0x69, 0x6e, 0x67,
	0x6c, 0x65, 0x44, 0x69, 0x73, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x12, 0x12, 0x0a, 0x04, 0x6e, 0x61,
	0x6d, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x12, 0x2b,
	0x0a, 0x06, 0x61, 0x6d, 0x6f, 0x75, 0x6e, 0x74, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x13,
	0x2e, 0x64, 0x73, 0x73, 0x2e, 0x44, 0x69, 0x73, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x41, 0x6d, 0x6f,
	0x75, 0x6e, 0x74, 0x52, 0x06, 0x61, 0x6d, 0x6f, 0x75, 0x6e, 0x74, 0x12, 0x30, 0x0a, 0x0a, 0x61,
	0x70, 0x70, 0x6c, 0x69, 0x65, 0x64, 0x5f, 0x74, 0x6f, 0x18, 0x03, 0x20, 0x01, 0x28, 0x0e, 0x32,
	0x11, 0x2e, 0x64, 0x73, 0x73, 0x2e, 0x50, 0x75, 0x72, 0x63, 0x68, 0x61, 0x73, 0x65, 0x49, 0x74,
	0x65, 0x6d, 0x52, 0x09, 0x61, 0x70, 0x70, 0x6c, 0x69, 0x65, 0x64, 0x54, 0x6f, 0x22, 0x57, 0x0a,
	0x0e, 0x44, 0x69, 0x73, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x42, 0x75, 0x6e, 0x64, 0x6c, 0x65, 0x12,
	0x12, 0x0a, 0x04, 0x63, 0x6f, 0x64, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x63,
	0x6f, 0x64, 0x65, 0x12, 0x31, 0x0a, 0x09, 0x64, 0x69, 0x73, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x73,
	0x18, 0x02, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x13, 0x2e, 0x64, 0x73, 0x73, 0x2e, 0x53, 0x69, 0x6e,
	0x67, 0x6c, 0x65, 0x44, 0x69, 0x73, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x52, 0x09, 0x64, 0x69, 0x73,
	0x63, 0x6f, 0x75, 0x6e, 0x74, 0x73, 0x22, 0x3d, 0x0a, 0x0e, 0x44, 0x69, 0x73, 0x63, 0x6f, 0x75,
	0x6e, 0x74, 0x41, 0x64, 0x64, 0x52, 0x65, 0x71, 0x12, 0x2b, 0x0a, 0x06, 0x62, 0x75, 0x6e, 0x64,
	0x6c, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x13, 0x2e, 0x64, 0x73, 0x73, 0x2e, 0x44,
	0x69, 0x73, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x42, 0x75, 0x6e, 0x64, 0x6c, 0x65, 0x52, 0x06, 0x62,
	0x75, 0x6e, 0x64, 0x6c, 0x65, 0x22, 0x10, 0x0a, 0x0e, 0x44, 0x69, 0x73, 0x63, 0x6f, 0x75, 0x6e,
	0x74, 0x41, 0x64, 0x64, 0x52, 0x65, 0x73, 0x22, 0x24, 0x0a, 0x0e, 0x44, 0x69, 0x73, 0x63, 0x6f,
	0x75, 0x6e, 0x74, 0x47, 0x65, 0x74, 0x52, 0x65, 0x71, 0x12, 0x12, 0x0a, 0x04, 0x63, 0x6f, 0x64,
	0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x63, 0x6f, 0x64, 0x65, 0x22, 0x3d, 0x0a,
	0x0e, 0x44, 0x69, 0x73, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x47, 0x65, 0x74, 0x52, 0x65, 0x73, 0x12,
	0x2b, 0x0a, 0x06, 0x62, 0x75, 0x6e, 0x64, 0x6c, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32,
	0x13, 0x2e, 0x64, 0x73, 0x73, 0x2e, 0x44, 0x69, 0x73, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x42, 0x75,
	0x6e, 0x64, 0x6c, 0x65, 0x52, 0x06, 0x62, 0x75, 0x6e, 0x64, 0x6c, 0x65, 0x22, 0x11, 0x0a, 0x0f,
	0x44, 0x69, 0x73, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x4c, 0x69, 0x73, 0x74, 0x52, 0x65, 0x71, 0x22,
	0x40, 0x0a, 0x0f, 0x44, 0x69, 0x73, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x4c, 0x69, 0x73, 0x74, 0x52,
	0x65, 0x73, 0x12, 0x2d, 0x0a, 0x07, 0x62, 0x75, 0x6e, 0x64, 0x6c, 0x65, 0x73, 0x18, 0x01, 0x20,
	0x03, 0x28, 0x0b, 0x32, 0x13, 0x2e, 0x64, 0x73, 0x73, 0x2e, 0x44, 0x69, 0x73, 0x63, 0x6f, 0x75,
	0x6e, 0x74, 0x42, 0x75, 0x6e, 0x64, 0x6c, 0x65, 0x52, 0x07, 0x62, 0x75, 0x6e, 0x64, 0x6c, 0x65,
	0x73, 0x22, 0x5a, 0x0a, 0x11, 0x44, 0x69, 0x73, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x55, 0x70, 0x64,
	0x61, 0x74, 0x65, 0x52, 0x65, 0x71, 0x12, 0x18, 0x0a, 0x07, 0x6f, 0x6c, 0x64, 0x43, 0x6f, 0x64,
	0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x07, 0x6f, 0x6c, 0x64, 0x43, 0x6f, 0x64, 0x65,
	0x12, 0x2b, 0x0a, 0x06, 0x62, 0x75, 0x6e, 0x64, 0x6c, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0b,
	0x32, 0x13, 0x2e, 0x64, 0x73, 0x73, 0x2e, 0x44, 0x69, 0x73, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x42,
	0x75, 0x6e, 0x64, 0x6c, 0x65, 0x52, 0x06, 0x62, 0x75, 0x6e, 0x64, 0x6c, 0x65, 0x22, 0x13, 0x0a,
	0x11, 0x44, 0x69, 0x73, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x55, 0x70, 0x64, 0x61, 0x74, 0x65, 0x52,
	0x65, 0x73, 0x22, 0x27, 0x0a, 0x11, 0x44, 0x69, 0x73, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x44, 0x65,
	0x6c, 0x65, 0x74, 0x65, 0x52, 0x65, 0x71, 0x12, 0x12, 0x0a, 0x04, 0x63, 0x6f, 0x64, 0x65, 0x18,
	0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x63, 0x6f, 0x64, 0x65, 0x22, 0x13, 0x0a, 0x11, 0x44,
	0x69, 0x73, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x44, 0x65, 0x6c, 0x65, 0x74, 0x65, 0x52, 0x65, 0x73,
	0x2a, 0xbe, 0x01, 0x0a, 0x0c, 0x50, 0x75, 0x72, 0x63, 0x68, 0x61, 0x73, 0x65, 0x49, 0x74, 0x65,
	0x6d, 0x12, 0x1f, 0x0a, 0x1b, 0x46, 0x75, 0x6c, 0x6c, 0x57, 0x65, 0x65, 0x6b, 0x65, 0x6e, 0x64,
	0x50, 0x61, 0x73, 0x73, 0x50, 0x75, 0x72, 0x63, 0x68, 0x61, 0x73, 0x65, 0x49, 0x74, 0x65, 0x6d,
	0x10, 0x00, 0x12, 0x1d, 0x0a, 0x19, 0x44, 0x61, 0x6e, 0x63, 0x65, 0x4f, 0x6e, 0x6c, 0x79, 0x50,
	0x61, 0x73, 0x73, 0x50, 0x75, 0x72, 0x63, 0x68, 0x61, 0x73, 0x65, 0x49, 0x74, 0x65, 0x6d, 0x10,
	0x01, 0x12, 0x1b, 0x0a, 0x17, 0x4d, 0x69, 0x78, 0x41, 0x6e, 0x64, 0x4d, 0x61, 0x74, 0x63, 0x68,
	0x50, 0x75, 0x72, 0x63, 0x68, 0x61, 0x73, 0x65, 0x49, 0x74, 0x65, 0x6d, 0x10, 0x02, 0x12, 0x18,
	0x0a, 0x14, 0x53, 0x6f, 0x6c, 0x6f, 0x4a, 0x61, 0x7a, 0x7a, 0x50, 0x75, 0x72, 0x63, 0x68, 0x61,
	0x73, 0x65, 0x49, 0x74, 0x65, 0x6d, 0x10, 0x03, 0x12, 0x1f, 0x0a, 0x1b, 0x54, 0x65, 0x61, 0x6d,
	0x43, 0x6f, 0x6d, 0x70, 0x65, 0x74, 0x69, 0x74, 0x69, 0x6f, 0x6e, 0x50, 0x75, 0x72, 0x63, 0x68,
	0x61, 0x73, 0x65, 0x49, 0x74, 0x65, 0x6d, 0x10, 0x04, 0x12, 0x16, 0x0a, 0x12, 0x54, 0x53, 0x68,
	0x69, 0x72, 0x74, 0x50, 0x75, 0x72, 0x63, 0x68, 0x61, 0x73, 0x65, 0x49, 0x74, 0x65, 0x6d, 0x10,
	0x05, 0x32, 0x94, 0x02, 0x0a, 0x08, 0x44, 0x69, 0x73, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x12, 0x2f,
	0x0a, 0x03, 0x41, 0x64, 0x64, 0x12, 0x13, 0x2e, 0x64, 0x73, 0x73, 0x2e, 0x44, 0x69, 0x73, 0x63,
	0x6f, 0x75, 0x6e, 0x74, 0x41, 0x64, 0x64, 0x52, 0x65, 0x71, 0x1a, 0x13, 0x2e, 0x64, 0x73, 0x73,
	0x2e, 0x44, 0x69, 0x73, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x41, 0x64, 0x64, 0x52, 0x65, 0x73, 0x12,
	0x2f, 0x0a, 0x03, 0x47, 0x65, 0x74, 0x12, 0x13, 0x2e, 0x64, 0x73, 0x73, 0x2e, 0x44, 0x69, 0x73,
	0x63, 0x6f, 0x75, 0x6e, 0x74, 0x47, 0x65, 0x74, 0x52, 0x65, 0x71, 0x1a, 0x13, 0x2e, 0x64, 0x73,
	0x73, 0x2e, 0x44, 0x69, 0x73, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x47, 0x65, 0x74, 0x52, 0x65, 0x73,
	0x12, 0x32, 0x0a, 0x04, 0x4c, 0x69, 0x73, 0x74, 0x12, 0x14, 0x2e, 0x64, 0x73, 0x73, 0x2e, 0x44,
	0x69, 0x73, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x4c, 0x69, 0x73, 0x74, 0x52, 0x65, 0x71, 0x1a, 0x14,
	0x2e, 0x64, 0x73, 0x73, 0x2e, 0x44, 0x69, 0x73, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x4c, 0x69, 0x73,
	0x74, 0x52, 0x65, 0x73, 0x12, 0x38, 0x0a, 0x06, 0x55, 0x70, 0x64, 0x61, 0x74, 0x65, 0x12, 0x16,
	0x2e, 0x64, 0x73, 0x73, 0x2e, 0x44, 0x69, 0x73, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x55, 0x70, 0x64,
	0x61, 0x74, 0x65, 0x52, 0x65, 0x71, 0x1a, 0x16, 0x2e, 0x64, 0x73, 0x73, 0x2e, 0x44, 0x69, 0x73,
	0x63, 0x6f, 0x75, 0x6e, 0x74, 0x55, 0x70, 0x64, 0x61, 0x74, 0x65, 0x52, 0x65, 0x73, 0x12, 0x38,
	0x0a, 0x06, 0x44, 0x65, 0x6c, 0x65, 0x74, 0x65, 0x12, 0x16, 0x2e, 0x64, 0x73, 0x73, 0x2e, 0x44,
	0x69, 0x73, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x44, 0x65, 0x6c, 0x65, 0x74, 0x65, 0x52, 0x65, 0x71,
	0x1a, 0x16, 0x2e, 0x64, 0x73, 0x73, 0x2e, 0x44, 0x69, 0x73, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x44,
	0x65, 0x6c, 0x65, 0x74, 0x65, 0x52, 0x65, 0x73, 0x42, 0x09, 0x5a, 0x07, 0x72, 0x70, 0x63, 0x2f,
	0x64, 0x73, 0x73, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_discount_proto_rawDescOnce sync.Once
	file_discount_proto_rawDescData = file_discount_proto_rawDesc
)

func file_discount_proto_rawDescGZIP() []byte {
	file_discount_proto_rawDescOnce.Do(func() {
		file_discount_proto_rawDescData = protoimpl.X.CompressGZIP(file_discount_proto_rawDescData)
	})
	return file_discount_proto_rawDescData
}

var file_discount_proto_enumTypes = make([]protoimpl.EnumInfo, 1)
var file_discount_proto_msgTypes = make([]protoimpl.MessageInfo, 13)
var file_discount_proto_goTypes = []interface{}{
	(PurchaseItem)(0),         // 0: dss.PurchaseItem
	(*DiscountAmount)(nil),    // 1: dss.DiscountAmount
	(*SingleDiscount)(nil),    // 2: dss.SingleDiscount
	(*DiscountBundle)(nil),    // 3: dss.DiscountBundle
	(*DiscountAddReq)(nil),    // 4: dss.DiscountAddReq
	(*DiscountAddRes)(nil),    // 5: dss.DiscountAddRes
	(*DiscountGetReq)(nil),    // 6: dss.DiscountGetReq
	(*DiscountGetRes)(nil),    // 7: dss.DiscountGetRes
	(*DiscountListReq)(nil),   // 8: dss.DiscountListReq
	(*DiscountListRes)(nil),   // 9: dss.DiscountListRes
	(*DiscountUpdateReq)(nil), // 10: dss.DiscountUpdateReq
	(*DiscountUpdateRes)(nil), // 11: dss.DiscountUpdateRes
	(*DiscountDeleteReq)(nil), // 12: dss.DiscountDeleteReq
	(*DiscountDeleteRes)(nil), // 13: dss.DiscountDeleteRes
	(*emptypb.Empty)(nil),     // 14: google.protobuf.Empty
}
var file_discount_proto_depIdxs = []int32{
	14, // 0: dss.DiscountAmount.squareNotFound:type_name -> google.protobuf.Empty
	1,  // 1: dss.SingleDiscount.amount:type_name -> dss.DiscountAmount
	0,  // 2: dss.SingleDiscount.applied_to:type_name -> dss.PurchaseItem
	2,  // 3: dss.DiscountBundle.discounts:type_name -> dss.SingleDiscount
	3,  // 4: dss.DiscountAddReq.bundle:type_name -> dss.DiscountBundle
	3,  // 5: dss.DiscountGetRes.bundle:type_name -> dss.DiscountBundle
	3,  // 6: dss.DiscountListRes.bundles:type_name -> dss.DiscountBundle
	3,  // 7: dss.DiscountUpdateReq.bundle:type_name -> dss.DiscountBundle
	4,  // 8: dss.Discount.Add:input_type -> dss.DiscountAddReq
	6,  // 9: dss.Discount.Get:input_type -> dss.DiscountGetReq
	8,  // 10: dss.Discount.List:input_type -> dss.DiscountListReq
	10, // 11: dss.Discount.Update:input_type -> dss.DiscountUpdateReq
	12, // 12: dss.Discount.Delete:input_type -> dss.DiscountDeleteReq
	5,  // 13: dss.Discount.Add:output_type -> dss.DiscountAddRes
	7,  // 14: dss.Discount.Get:output_type -> dss.DiscountGetRes
	9,  // 15: dss.Discount.List:output_type -> dss.DiscountListRes
	11, // 16: dss.Discount.Update:output_type -> dss.DiscountUpdateRes
	13, // 17: dss.Discount.Delete:output_type -> dss.DiscountDeleteRes
	13, // [13:18] is the sub-list for method output_type
	8,  // [8:13] is the sub-list for method input_type
	8,  // [8:8] is the sub-list for extension type_name
	8,  // [8:8] is the sub-list for extension extendee
	0,  // [0:8] is the sub-list for field type_name
}

func init() { file_discount_proto_init() }
func file_discount_proto_init() {
	if File_discount_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_discount_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
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
		file_discount_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
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
		file_discount_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
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
		file_discount_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
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
		file_discount_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
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
		file_discount_proto_msgTypes[5].Exporter = func(v interface{}, i int) interface{} {
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
		file_discount_proto_msgTypes[6].Exporter = func(v interface{}, i int) interface{} {
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
		file_discount_proto_msgTypes[7].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*DiscountListReq); i {
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
		file_discount_proto_msgTypes[8].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*DiscountListRes); i {
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
		file_discount_proto_msgTypes[9].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*DiscountUpdateReq); i {
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
		file_discount_proto_msgTypes[10].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*DiscountUpdateRes); i {
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
		file_discount_proto_msgTypes[11].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*DiscountDeleteReq); i {
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
		file_discount_proto_msgTypes[12].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*DiscountDeleteRes); i {
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
	file_discount_proto_msgTypes[0].OneofWrappers = []interface{}{
		(*DiscountAmount_Dollar)(nil),
		(*DiscountAmount_Percent)(nil),
		(*DiscountAmount_SquareNotFound)(nil),
	}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_discount_proto_rawDesc,
			NumEnums:      1,
			NumMessages:   13,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_discount_proto_goTypes,
		DependencyIndexes: file_discount_proto_depIdxs,
		EnumInfos:         file_discount_proto_enumTypes,
		MessageInfos:      file_discount_proto_msgTypes,
	}.Build()
	File_discount_proto = out.File
	file_discount_proto_rawDesc = nil
	file_discount_proto_goTypes = nil
	file_discount_proto_depIdxs = nil
}
