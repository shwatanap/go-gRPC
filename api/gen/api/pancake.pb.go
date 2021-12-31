// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.27.1
// 	protoc        v3.19.1
// source: pancake.proto

package api

import (
	context "context"
	grpc "google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
	timestamppb "google.golang.org/protobuf/types/known/timestamppb"
	reflect "reflect"
	sync "sync"
)

const (
	// Verify that this generated code is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(20 - protoimpl.MinVersion)
	// Verify that runtime/protoimpl is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(protoimpl.MaxVersion - 20)
)

//メニュー表
type Pancake_Menu int32

const (
	Pancake_UNKNOWN           Pancake_Menu = 0
	Pancake_CLASSIC           Pancake_Menu = 1
	Pancake_BANANA_AND_WHIP   Pancake_Menu = 2
	Pancake_BACON_AND_CHEESE  Pancake_Menu = 3
	Pancake_MIX_BERRY         Pancake_Menu = 4
	Pancake_BAKED_MARSHMALLOW Pancake_Menu = 5
	Pancake_SPICY_CURRY       Pancake_Menu = 6
)

// Enum value maps for Pancake_Menu.
var (
	Pancake_Menu_name = map[int32]string{
		0: "UNKNOWN",
		1: "CLASSIC",
		2: "BANANA_AND_WHIP",
		3: "BACON_AND_CHEESE",
		4: "MIX_BERRY",
		5: "BAKED_MARSHMALLOW",
		6: "SPICY_CURRY",
	}
	Pancake_Menu_value = map[string]int32{
		"UNKNOWN":           0,
		"CLASSIC":           1,
		"BANANA_AND_WHIP":   2,
		"BACON_AND_CHEESE":  3,
		"MIX_BERRY":         4,
		"BAKED_MARSHMALLOW": 5,
		"SPICY_CURRY":       6,
	}
)

func (x Pancake_Menu) Enum() *Pancake_Menu {
	p := new(Pancake_Menu)
	*p = x
	return p
}

func (x Pancake_Menu) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (Pancake_Menu) Descriptor() protoreflect.EnumDescriptor {
	return file_pancake_proto_enumTypes[0].Descriptor()
}

func (Pancake_Menu) Type() protoreflect.EnumType {
	return &file_pancake_proto_enumTypes[0]
}

func (x Pancake_Menu) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use Pancake_Menu.Descriptor instead.
func (Pancake_Menu) EnumDescriptor() ([]byte, []int) {
	return file_pancake_proto_rawDescGZIP(), []int{0, 0}
}

//Pancakeは一枚一枚の焼かれたパンケーキを表します。
type Pancake struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	//シェフの名前
	ChefName string `protobuf:"bytes,1,opt,name=chef_name,json=chefName,proto3" json:"chef_name,omitempty"`
	//メニュー
	Menu Pancake_Menu `protobuf:"varint,2,opt,name=menu,proto3,enum=pancake.baker.Pancake_Menu" json:"menu,omitempty"`
	//焼き具合を表すスコアです(0-0.9)
	TechnicalScore float32 `protobuf:"fixed32,3,opt,name=technical_score,json=technicalScore,proto3" json:"technical_score,omitempty"`
	//焼いた日時
	CreateTime *timestamppb.Timestamp `protobuf:"bytes,15,opt,name=create_time,json=createTime,proto3" json:"create_time,omitempty"`
}

func (x *Pancake) Reset() {
	*x = Pancake{}
	if protoimpl.UnsafeEnabled {
		mi := &file_pancake_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Pancake) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Pancake) ProtoMessage() {}

func (x *Pancake) ProtoReflect() protoreflect.Message {
	mi := &file_pancake_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Pancake.ProtoReflect.Descriptor instead.
func (*Pancake) Descriptor() ([]byte, []int) {
	return file_pancake_proto_rawDescGZIP(), []int{0}
}

func (x *Pancake) GetChefName() string {
	if x != nil {
		return x.ChefName
	}
	return ""
}

func (x *Pancake) GetMenu() Pancake_Menu {
	if x != nil {
		return x.Menu
	}
	return Pancake_UNKNOWN
}

func (x *Pancake) GetTechnicalScore() float32 {
	if x != nil {
		return x.TechnicalScore
	}
	return 0
}

func (x *Pancake) GetCreateTime() *timestamppb.Timestamp {
	if x != nil {
		return x.CreateTime
	}
	return nil
}

//Reportはどのくらいパンケーキを焼いたかについての報告書を表します
type Report struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	BakeCounts []*Report_BakeCount `protobuf:"bytes,1,rep,name=bake_counts,json=bakeCounts,proto3" json:"bake_counts,omitempty"`
}

func (x *Report) Reset() {
	*x = Report{}
	if protoimpl.UnsafeEnabled {
		mi := &file_pancake_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Report) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Report) ProtoMessage() {}

func (x *Report) ProtoReflect() protoreflect.Message {
	mi := &file_pancake_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Report.ProtoReflect.Descriptor instead.
func (*Report) Descriptor() ([]byte, []int) {
	return file_pancake_proto_rawDescGZIP(), []int{1}
}

func (x *Report) GetBakeCounts() []*Report_BakeCount {
	if x != nil {
		return x.BakeCounts
	}
	return nil
}

type BakeRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Menu Pancake_Menu `protobuf:"varint,1,opt,name=menu,proto3,enum=pancake.baker.Pancake_Menu" json:"menu,omitempty"`
}

func (x *BakeRequest) Reset() {
	*x = BakeRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_pancake_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *BakeRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*BakeRequest) ProtoMessage() {}

func (x *BakeRequest) ProtoReflect() protoreflect.Message {
	mi := &file_pancake_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use BakeRequest.ProtoReflect.Descriptor instead.
func (*BakeRequest) Descriptor() ([]byte, []int) {
	return file_pancake_proto_rawDescGZIP(), []int{2}
}

func (x *BakeRequest) GetMenu() Pancake_Menu {
	if x != nil {
		return x.Menu
	}
	return Pancake_UNKNOWN
}

type BakeResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Pancake *Pancake `protobuf:"bytes,1,opt,name=pancake,proto3" json:"pancake,omitempty"`
}

func (x *BakeResponse) Reset() {
	*x = BakeResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_pancake_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *BakeResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*BakeResponse) ProtoMessage() {}

func (x *BakeResponse) ProtoReflect() protoreflect.Message {
	mi := &file_pancake_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use BakeResponse.ProtoReflect.Descriptor instead.
func (*BakeResponse) Descriptor() ([]byte, []int) {
	return file_pancake_proto_rawDescGZIP(), []int{3}
}

func (x *BakeResponse) GetPancake() *Pancake {
	if x != nil {
		return x.Pancake
	}
	return nil
}

type ReportRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields
}

func (x *ReportRequest) Reset() {
	*x = ReportRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_pancake_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ReportRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ReportRequest) ProtoMessage() {}

func (x *ReportRequest) ProtoReflect() protoreflect.Message {
	mi := &file_pancake_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ReportRequest.ProtoReflect.Descriptor instead.
func (*ReportRequest) Descriptor() ([]byte, []int) {
	return file_pancake_proto_rawDescGZIP(), []int{4}
}

type ReportResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Report *Report `protobuf:"bytes,1,opt,name=report,proto3" json:"report,omitempty"`
}

func (x *ReportResponse) Reset() {
	*x = ReportResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_pancake_proto_msgTypes[5]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ReportResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ReportResponse) ProtoMessage() {}

func (x *ReportResponse) ProtoReflect() protoreflect.Message {
	mi := &file_pancake_proto_msgTypes[5]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ReportResponse.ProtoReflect.Descriptor instead.
func (*ReportResponse) Descriptor() ([]byte, []int) {
	return file_pancake_proto_rawDescGZIP(), []int{5}
}

func (x *ReportResponse) GetReport() *Report {
	if x != nil {
		return x.Report
	}
	return nil
}

type Report_BakeCount struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Menu  Pancake_Menu `protobuf:"varint,1,opt,name=menu,proto3,enum=pancake.baker.Pancake_Menu" json:"menu,omitempty"`
	Count int32        `protobuf:"varint,2,opt,name=count,proto3" json:"count,omitempty"`
}

func (x *Report_BakeCount) Reset() {
	*x = Report_BakeCount{}
	if protoimpl.UnsafeEnabled {
		mi := &file_pancake_proto_msgTypes[6]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Report_BakeCount) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Report_BakeCount) ProtoMessage() {}

func (x *Report_BakeCount) ProtoReflect() protoreflect.Message {
	mi := &file_pancake_proto_msgTypes[6]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Report_BakeCount.ProtoReflect.Descriptor instead.
func (*Report_BakeCount) Descriptor() ([]byte, []int) {
	return file_pancake_proto_rawDescGZIP(), []int{1, 0}
}

func (x *Report_BakeCount) GetMenu() Pancake_Menu {
	if x != nil {
		return x.Menu
	}
	return Pancake_UNKNOWN
}

func (x *Report_BakeCount) GetCount() int32 {
	if x != nil {
		return x.Count
	}
	return 0
}

var File_pancake_proto protoreflect.FileDescriptor

var file_pancake_proto_rawDesc = []byte{
	0x0a, 0x0d, 0x70, 0x61, 0x6e, 0x63, 0x61, 0x6b, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12,
	0x0d, 0x70, 0x61, 0x6e, 0x63, 0x61, 0x6b, 0x65, 0x2e, 0x62, 0x61, 0x6b, 0x65, 0x72, 0x1a, 0x1f,
	0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2f,
	0x74, 0x69, 0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d, 0x70, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22,
	0xc2, 0x02, 0x0a, 0x07, 0x50, 0x61, 0x6e, 0x63, 0x61, 0x6b, 0x65, 0x12, 0x1b, 0x0a, 0x09, 0x63,
	0x68, 0x65, 0x66, 0x5f, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08,
	0x63, 0x68, 0x65, 0x66, 0x4e, 0x61, 0x6d, 0x65, 0x12, 0x2f, 0x0a, 0x04, 0x6d, 0x65, 0x6e, 0x75,
	0x18, 0x02, 0x20, 0x01, 0x28, 0x0e, 0x32, 0x1b, 0x2e, 0x70, 0x61, 0x6e, 0x63, 0x61, 0x6b, 0x65,
	0x2e, 0x62, 0x61, 0x6b, 0x65, 0x72, 0x2e, 0x50, 0x61, 0x6e, 0x63, 0x61, 0x6b, 0x65, 0x2e, 0x4d,
	0x65, 0x6e, 0x75, 0x52, 0x04, 0x6d, 0x65, 0x6e, 0x75, 0x12, 0x27, 0x0a, 0x0f, 0x74, 0x65, 0x63,
	0x68, 0x6e, 0x69, 0x63, 0x61, 0x6c, 0x5f, 0x73, 0x63, 0x6f, 0x72, 0x65, 0x18, 0x03, 0x20, 0x01,
	0x28, 0x02, 0x52, 0x0e, 0x74, 0x65, 0x63, 0x68, 0x6e, 0x69, 0x63, 0x61, 0x6c, 0x53, 0x63, 0x6f,
	0x72, 0x65, 0x12, 0x3b, 0x0a, 0x0b, 0x63, 0x72, 0x65, 0x61, 0x74, 0x65, 0x5f, 0x74, 0x69, 0x6d,
	0x65, 0x18, 0x0f, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1a, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65,
	0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x54, 0x69, 0x6d, 0x65, 0x73, 0x74,
	0x61, 0x6d, 0x70, 0x52, 0x0a, 0x63, 0x72, 0x65, 0x61, 0x74, 0x65, 0x54, 0x69, 0x6d, 0x65, 0x22,
	0x82, 0x01, 0x0a, 0x04, 0x4d, 0x65, 0x6e, 0x75, 0x12, 0x0b, 0x0a, 0x07, 0x55, 0x4e, 0x4b, 0x4e,
	0x4f, 0x57, 0x4e, 0x10, 0x00, 0x12, 0x0b, 0x0a, 0x07, 0x43, 0x4c, 0x41, 0x53, 0x53, 0x49, 0x43,
	0x10, 0x01, 0x12, 0x13, 0x0a, 0x0f, 0x42, 0x41, 0x4e, 0x41, 0x4e, 0x41, 0x5f, 0x41, 0x4e, 0x44,
	0x5f, 0x57, 0x48, 0x49, 0x50, 0x10, 0x02, 0x12, 0x14, 0x0a, 0x10, 0x42, 0x41, 0x43, 0x4f, 0x4e,
	0x5f, 0x41, 0x4e, 0x44, 0x5f, 0x43, 0x48, 0x45, 0x45, 0x53, 0x45, 0x10, 0x03, 0x12, 0x0d, 0x0a,
	0x09, 0x4d, 0x49, 0x58, 0x5f, 0x42, 0x45, 0x52, 0x52, 0x59, 0x10, 0x04, 0x12, 0x15, 0x0a, 0x11,
	0x42, 0x41, 0x4b, 0x45, 0x44, 0x5f, 0x4d, 0x41, 0x52, 0x53, 0x48, 0x4d, 0x41, 0x4c, 0x4c, 0x4f,
	0x57, 0x10, 0x05, 0x12, 0x0f, 0x0a, 0x0b, 0x53, 0x50, 0x49, 0x43, 0x59, 0x5f, 0x43, 0x55, 0x52,
	0x52, 0x59, 0x10, 0x06, 0x22, 0x9e, 0x01, 0x0a, 0x06, 0x52, 0x65, 0x70, 0x6f, 0x72, 0x74, 0x12,
	0x40, 0x0a, 0x0b, 0x62, 0x61, 0x6b, 0x65, 0x5f, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x73, 0x18, 0x01,
	0x20, 0x03, 0x28, 0x0b, 0x32, 0x1f, 0x2e, 0x70, 0x61, 0x6e, 0x63, 0x61, 0x6b, 0x65, 0x2e, 0x62,
	0x61, 0x6b, 0x65, 0x72, 0x2e, 0x52, 0x65, 0x70, 0x6f, 0x72, 0x74, 0x2e, 0x42, 0x61, 0x6b, 0x65,
	0x43, 0x6f, 0x75, 0x6e, 0x74, 0x52, 0x0a, 0x62, 0x61, 0x6b, 0x65, 0x43, 0x6f, 0x75, 0x6e, 0x74,
	0x73, 0x1a, 0x52, 0x0a, 0x09, 0x42, 0x61, 0x6b, 0x65, 0x43, 0x6f, 0x75, 0x6e, 0x74, 0x12, 0x2f,
	0x0a, 0x04, 0x6d, 0x65, 0x6e, 0x75, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0e, 0x32, 0x1b, 0x2e, 0x70,
	0x61, 0x6e, 0x63, 0x61, 0x6b, 0x65, 0x2e, 0x62, 0x61, 0x6b, 0x65, 0x72, 0x2e, 0x50, 0x61, 0x6e,
	0x63, 0x61, 0x6b, 0x65, 0x2e, 0x4d, 0x65, 0x6e, 0x75, 0x52, 0x04, 0x6d, 0x65, 0x6e, 0x75, 0x12,
	0x14, 0x0a, 0x05, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x18, 0x02, 0x20, 0x01, 0x28, 0x05, 0x52, 0x05,
	0x63, 0x6f, 0x75, 0x6e, 0x74, 0x22, 0x3e, 0x0a, 0x0b, 0x42, 0x61, 0x6b, 0x65, 0x52, 0x65, 0x71,
	0x75, 0x65, 0x73, 0x74, 0x12, 0x2f, 0x0a, 0x04, 0x6d, 0x65, 0x6e, 0x75, 0x18, 0x01, 0x20, 0x01,
	0x28, 0x0e, 0x32, 0x1b, 0x2e, 0x70, 0x61, 0x6e, 0x63, 0x61, 0x6b, 0x65, 0x2e, 0x62, 0x61, 0x6b,
	0x65, 0x72, 0x2e, 0x50, 0x61, 0x6e, 0x63, 0x61, 0x6b, 0x65, 0x2e, 0x4d, 0x65, 0x6e, 0x75, 0x52,
	0x04, 0x6d, 0x65, 0x6e, 0x75, 0x22, 0x40, 0x0a, 0x0c, 0x42, 0x61, 0x6b, 0x65, 0x52, 0x65, 0x73,
	0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x30, 0x0a, 0x07, 0x70, 0x61, 0x6e, 0x63, 0x61, 0x6b, 0x65,
	0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x16, 0x2e, 0x70, 0x61, 0x6e, 0x63, 0x61, 0x6b, 0x65,
	0x2e, 0x62, 0x61, 0x6b, 0x65, 0x72, 0x2e, 0x50, 0x61, 0x6e, 0x63, 0x61, 0x6b, 0x65, 0x52, 0x07,
	0x70, 0x61, 0x6e, 0x63, 0x61, 0x6b, 0x65, 0x22, 0x0f, 0x0a, 0x0d, 0x52, 0x65, 0x70, 0x6f, 0x72,
	0x74, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x22, 0x3f, 0x0a, 0x0e, 0x52, 0x65, 0x70, 0x6f,
	0x72, 0x74, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x2d, 0x0a, 0x06, 0x72, 0x65,
	0x70, 0x6f, 0x72, 0x74, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x15, 0x2e, 0x70, 0x61, 0x6e,
	0x63, 0x61, 0x6b, 0x65, 0x2e, 0x62, 0x61, 0x6b, 0x65, 0x72, 0x2e, 0x52, 0x65, 0x70, 0x6f, 0x72,
	0x74, 0x52, 0x06, 0x72, 0x65, 0x70, 0x6f, 0x72, 0x74, 0x32, 0xa1, 0x01, 0x0a, 0x13, 0x50, 0x61,
	0x6e, 0x63, 0x61, 0x6b, 0x65, 0x42, 0x61, 0x6b, 0x65, 0x72, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63,
	0x65, 0x12, 0x41, 0x0a, 0x04, 0x42, 0x61, 0x6b, 0x65, 0x12, 0x1a, 0x2e, 0x70, 0x61, 0x6e, 0x63,
	0x61, 0x6b, 0x65, 0x2e, 0x62, 0x61, 0x6b, 0x65, 0x72, 0x2e, 0x42, 0x61, 0x6b, 0x65, 0x52, 0x65,
	0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x1b, 0x2e, 0x70, 0x61, 0x6e, 0x63, 0x61, 0x6b, 0x65, 0x2e,
	0x62, 0x61, 0x6b, 0x65, 0x72, 0x2e, 0x42, 0x61, 0x6b, 0x65, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e,
	0x73, 0x65, 0x22, 0x00, 0x12, 0x47, 0x0a, 0x06, 0x52, 0x65, 0x70, 0x6f, 0x72, 0x74, 0x12, 0x1c,
	0x2e, 0x70, 0x61, 0x6e, 0x63, 0x61, 0x6b, 0x65, 0x2e, 0x62, 0x61, 0x6b, 0x65, 0x72, 0x2e, 0x52,
	0x65, 0x70, 0x6f, 0x72, 0x74, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x1d, 0x2e, 0x70,
	0x61, 0x6e, 0x63, 0x61, 0x6b, 0x65, 0x2e, 0x62, 0x61, 0x6b, 0x65, 0x72, 0x2e, 0x52, 0x65, 0x70,
	0x6f, 0x72, 0x74, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x00, 0x42, 0x09, 0x5a,
	0x07, 0x67, 0x65, 0x6e, 0x2f, 0x61, 0x70, 0x69, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_pancake_proto_rawDescOnce sync.Once
	file_pancake_proto_rawDescData = file_pancake_proto_rawDesc
)

func file_pancake_proto_rawDescGZIP() []byte {
	file_pancake_proto_rawDescOnce.Do(func() {
		file_pancake_proto_rawDescData = protoimpl.X.CompressGZIP(file_pancake_proto_rawDescData)
	})
	return file_pancake_proto_rawDescData
}

var file_pancake_proto_enumTypes = make([]protoimpl.EnumInfo, 1)
var file_pancake_proto_msgTypes = make([]protoimpl.MessageInfo, 7)
var file_pancake_proto_goTypes = []interface{}{
	(Pancake_Menu)(0),             // 0: pancake.baker.Pancake.Menu
	(*Pancake)(nil),               // 1: pancake.baker.Pancake
	(*Report)(nil),                // 2: pancake.baker.Report
	(*BakeRequest)(nil),           // 3: pancake.baker.BakeRequest
	(*BakeResponse)(nil),          // 4: pancake.baker.BakeResponse
	(*ReportRequest)(nil),         // 5: pancake.baker.ReportRequest
	(*ReportResponse)(nil),        // 6: pancake.baker.ReportResponse
	(*Report_BakeCount)(nil),      // 7: pancake.baker.Report.BakeCount
	(*timestamppb.Timestamp)(nil), // 8: google.protobuf.Timestamp
}
var file_pancake_proto_depIdxs = []int32{
	0, // 0: pancake.baker.Pancake.menu:type_name -> pancake.baker.Pancake.Menu
	8, // 1: pancake.baker.Pancake.create_time:type_name -> google.protobuf.Timestamp
	7, // 2: pancake.baker.Report.bake_counts:type_name -> pancake.baker.Report.BakeCount
	0, // 3: pancake.baker.BakeRequest.menu:type_name -> pancake.baker.Pancake.Menu
	1, // 4: pancake.baker.BakeResponse.pancake:type_name -> pancake.baker.Pancake
	2, // 5: pancake.baker.ReportResponse.report:type_name -> pancake.baker.Report
	0, // 6: pancake.baker.Report.BakeCount.menu:type_name -> pancake.baker.Pancake.Menu
	3, // 7: pancake.baker.PancakeBakerService.Bake:input_type -> pancake.baker.BakeRequest
	5, // 8: pancake.baker.PancakeBakerService.Report:input_type -> pancake.baker.ReportRequest
	4, // 9: pancake.baker.PancakeBakerService.Bake:output_type -> pancake.baker.BakeResponse
	6, // 10: pancake.baker.PancakeBakerService.Report:output_type -> pancake.baker.ReportResponse
	9, // [9:11] is the sub-list for method output_type
	7, // [7:9] is the sub-list for method input_type
	7, // [7:7] is the sub-list for extension type_name
	7, // [7:7] is the sub-list for extension extendee
	0, // [0:7] is the sub-list for field type_name
}

func init() { file_pancake_proto_init() }
func file_pancake_proto_init() {
	if File_pancake_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_pancake_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Pancake); i {
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
		file_pancake_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Report); i {
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
		file_pancake_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*BakeRequest); i {
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
		file_pancake_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*BakeResponse); i {
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
		file_pancake_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ReportRequest); i {
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
		file_pancake_proto_msgTypes[5].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ReportResponse); i {
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
		file_pancake_proto_msgTypes[6].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Report_BakeCount); i {
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
			RawDescriptor: file_pancake_proto_rawDesc,
			NumEnums:      1,
			NumMessages:   7,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_pancake_proto_goTypes,
		DependencyIndexes: file_pancake_proto_depIdxs,
		EnumInfos:         file_pancake_proto_enumTypes,
		MessageInfos:      file_pancake_proto_msgTypes,
	}.Build()
	File_pancake_proto = out.File
	file_pancake_proto_rawDesc = nil
	file_pancake_proto_goTypes = nil
	file_pancake_proto_depIdxs = nil
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConnInterface

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion6

// PancakeBakerServiceClient is the client API for PancakeBakerService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type PancakeBakerServiceClient interface {
	// Bakeは指定されたメニューのパンケーキを焼く関数です
	Bake(ctx context.Context, in *BakeRequest, opts ...grpc.CallOption) (*BakeResponse, error)
	// Reportはメニューごとに焼いたパンケーキの数を返します
	Report(ctx context.Context, in *ReportRequest, opts ...grpc.CallOption) (*ReportResponse, error)
}

type pancakeBakerServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewPancakeBakerServiceClient(cc grpc.ClientConnInterface) PancakeBakerServiceClient {
	return &pancakeBakerServiceClient{cc}
}

func (c *pancakeBakerServiceClient) Bake(ctx context.Context, in *BakeRequest, opts ...grpc.CallOption) (*BakeResponse, error) {
	out := new(BakeResponse)
	err := c.cc.Invoke(ctx, "/pancake.baker.PancakeBakerService/Bake", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *pancakeBakerServiceClient) Report(ctx context.Context, in *ReportRequest, opts ...grpc.CallOption) (*ReportResponse, error) {
	out := new(ReportResponse)
	err := c.cc.Invoke(ctx, "/pancake.baker.PancakeBakerService/Report", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// PancakeBakerServiceServer is the server API for PancakeBakerService service.
type PancakeBakerServiceServer interface {
	// Bakeは指定されたメニューのパンケーキを焼く関数です
	Bake(context.Context, *BakeRequest) (*BakeResponse, error)
	// Reportはメニューごとに焼いたパンケーキの数を返します
	Report(context.Context, *ReportRequest) (*ReportResponse, error)
}

// UnimplementedPancakeBakerServiceServer can be embedded to have forward compatible implementations.
type UnimplementedPancakeBakerServiceServer struct {
}

func (*UnimplementedPancakeBakerServiceServer) Bake(context.Context, *BakeRequest) (*BakeResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Bake not implemented")
}
func (*UnimplementedPancakeBakerServiceServer) Report(context.Context, *ReportRequest) (*ReportResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Report not implemented")
}

func RegisterPancakeBakerServiceServer(s *grpc.Server, srv PancakeBakerServiceServer) {
	s.RegisterService(&_PancakeBakerService_serviceDesc, srv)
}

func _PancakeBakerService_Bake_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(BakeRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(PancakeBakerServiceServer).Bake(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/pancake.baker.PancakeBakerService/Bake",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(PancakeBakerServiceServer).Bake(ctx, req.(*BakeRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _PancakeBakerService_Report_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ReportRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(PancakeBakerServiceServer).Report(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/pancake.baker.PancakeBakerService/Report",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(PancakeBakerServiceServer).Report(ctx, req.(*ReportRequest))
	}
	return interceptor(ctx, in, info, handler)
}

var _PancakeBakerService_serviceDesc = grpc.ServiceDesc{
	ServiceName: "pancake.baker.PancakeBakerService",
	HandlerType: (*PancakeBakerServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "Bake",
			Handler:    _PancakeBakerService_Bake_Handler,
		},
		{
			MethodName: "Report",
			Handler:    _PancakeBakerService_Report_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "pancake.proto",
}
