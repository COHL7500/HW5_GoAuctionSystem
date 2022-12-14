// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.28.1
// 	protoc        v3.21.12
// source: proto/Hw5_GoAuctionSystem.proto

package GoAuctionSystem

import (
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

type Acks int32

const (
	Acks_ACK_FAIL      Acks = 0
	Acks_ACK_SUCCESS   Acks = 1
	Acks_ACK_EXCEPTION Acks = 2
)

// Enum value maps for Acks.
var (
	Acks_name = map[int32]string{
		0: "ACK_FAIL",
		1: "ACK_SUCCESS",
		2: "ACK_EXCEPTION",
	}
	Acks_value = map[string]int32{
		"ACK_FAIL":      0,
		"ACK_SUCCESS":   1,
		"ACK_EXCEPTION": 2,
	}
)

func (x Acks) Enum() *Acks {
	p := new(Acks)
	*p = x
	return p
}

func (x Acks) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (Acks) Descriptor() protoreflect.EnumDescriptor {
	return file_proto_Hw5_GoAuctionSystem_proto_enumTypes[0].Descriptor()
}

func (Acks) Type() protoreflect.EnumType {
	return &file_proto_Hw5_GoAuctionSystem_proto_enumTypes[0]
}

func (x Acks) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use Acks.Descriptor instead.
func (Acks) EnumDescriptor() ([]byte, []int) {
	return file_proto_Hw5_GoAuctionSystem_proto_rawDescGZIP(), []int{0}
}

type BidPost struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id      int32 `protobuf:"varint,1,opt,name=id,proto3" json:"id,omitempty"`
	Amount  int32 `protobuf:"varint,2,opt,name=amount,proto3" json:"amount,omitempty"`
	Lamport int64 `protobuf:"varint,3,opt,name=lamport,proto3" json:"lamport,omitempty"`
}

func (x *BidPost) Reset() {
	*x = BidPost{}
	if protoimpl.UnsafeEnabled {
		mi := &file_proto_Hw5_GoAuctionSystem_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *BidPost) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*BidPost) ProtoMessage() {}

func (x *BidPost) ProtoReflect() protoreflect.Message {
	mi := &file_proto_Hw5_GoAuctionSystem_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use BidPost.ProtoReflect.Descriptor instead.
func (*BidPost) Descriptor() ([]byte, []int) {
	return file_proto_Hw5_GoAuctionSystem_proto_rawDescGZIP(), []int{0}
}

func (x *BidPost) GetId() int32 {
	if x != nil {
		return x.Id
	}
	return 0
}

func (x *BidPost) GetAmount() int32 {
	if x != nil {
		return x.Amount
	}
	return 0
}

func (x *BidPost) GetLamport() int64 {
	if x != nil {
		return x.Lamport
	}
	return 0
}

type Empty struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields
}

func (x *Empty) Reset() {
	*x = Empty{}
	if protoimpl.UnsafeEnabled {
		mi := &file_proto_Hw5_GoAuctionSystem_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Empty) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Empty) ProtoMessage() {}

func (x *Empty) ProtoReflect() protoreflect.Message {
	mi := &file_proto_Hw5_GoAuctionSystem_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Empty.ProtoReflect.Descriptor instead.
func (*Empty) Descriptor() ([]byte, []int) {
	return file_proto_Hw5_GoAuctionSystem_proto_rawDescGZIP(), []int{1}
}

type Outcome struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Amount int32 `protobuf:"varint,1,opt,name=amount,proto3" json:"amount,omitempty"`
	Over   bool  `protobuf:"varint,2,opt,name=over,proto3" json:"over,omitempty"`
}

func (x *Outcome) Reset() {
	*x = Outcome{}
	if protoimpl.UnsafeEnabled {
		mi := &file_proto_Hw5_GoAuctionSystem_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Outcome) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Outcome) ProtoMessage() {}

func (x *Outcome) ProtoReflect() protoreflect.Message {
	mi := &file_proto_Hw5_GoAuctionSystem_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Outcome.ProtoReflect.Descriptor instead.
func (*Outcome) Descriptor() ([]byte, []int) {
	return file_proto_Hw5_GoAuctionSystem_proto_rawDescGZIP(), []int{2}
}

func (x *Outcome) GetAmount() int32 {
	if x != nil {
		return x.Amount
	}
	return 0
}

func (x *Outcome) GetOver() bool {
	if x != nil {
		return x.Over
	}
	return false
}

type Ack struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Ack Acks `protobuf:"varint,1,opt,name=ack,proto3,enum=GoAuctionSystem.Acks" json:"ack,omitempty"`
}

func (x *Ack) Reset() {
	*x = Ack{}
	if protoimpl.UnsafeEnabled {
		mi := &file_proto_Hw5_GoAuctionSystem_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Ack) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Ack) ProtoMessage() {}

func (x *Ack) ProtoReflect() protoreflect.Message {
	mi := &file_proto_Hw5_GoAuctionSystem_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Ack.ProtoReflect.Descriptor instead.
func (*Ack) Descriptor() ([]byte, []int) {
	return file_proto_Hw5_GoAuctionSystem_proto_rawDescGZIP(), []int{3}
}

func (x *Ack) GetAck() Acks {
	if x != nil {
		return x.Ack
	}
	return Acks_ACK_FAIL
}

var File_proto_Hw5_GoAuctionSystem_proto protoreflect.FileDescriptor

var file_proto_Hw5_GoAuctionSystem_proto_rawDesc = []byte{
	0x0a, 0x1f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2f, 0x48, 0x77, 0x35, 0x5f, 0x47, 0x6f, 0x41, 0x75,
	0x63, 0x74, 0x69, 0x6f, 0x6e, 0x53, 0x79, 0x73, 0x74, 0x65, 0x6d, 0x2e, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x12, 0x0f, 0x47, 0x6f, 0x41, 0x75, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x53, 0x79, 0x73, 0x74,
	0x65, 0x6d, 0x22, 0x4b, 0x0a, 0x07, 0x42, 0x69, 0x64, 0x50, 0x6f, 0x73, 0x74, 0x12, 0x0e, 0x0a,
	0x02, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x05, 0x52, 0x02, 0x69, 0x64, 0x12, 0x16, 0x0a,
	0x06, 0x61, 0x6d, 0x6f, 0x75, 0x6e, 0x74, 0x18, 0x02, 0x20, 0x01, 0x28, 0x05, 0x52, 0x06, 0x61,
	0x6d, 0x6f, 0x75, 0x6e, 0x74, 0x12, 0x18, 0x0a, 0x07, 0x6c, 0x61, 0x6d, 0x70, 0x6f, 0x72, 0x74,
	0x18, 0x03, 0x20, 0x01, 0x28, 0x03, 0x52, 0x07, 0x6c, 0x61, 0x6d, 0x70, 0x6f, 0x72, 0x74, 0x22,
	0x07, 0x0a, 0x05, 0x45, 0x6d, 0x70, 0x74, 0x79, 0x22, 0x35, 0x0a, 0x07, 0x4f, 0x75, 0x74, 0x63,
	0x6f, 0x6d, 0x65, 0x12, 0x16, 0x0a, 0x06, 0x61, 0x6d, 0x6f, 0x75, 0x6e, 0x74, 0x18, 0x01, 0x20,
	0x01, 0x28, 0x05, 0x52, 0x06, 0x61, 0x6d, 0x6f, 0x75, 0x6e, 0x74, 0x12, 0x12, 0x0a, 0x04, 0x6f,
	0x76, 0x65, 0x72, 0x18, 0x02, 0x20, 0x01, 0x28, 0x08, 0x52, 0x04, 0x6f, 0x76, 0x65, 0x72, 0x22,
	0x2e, 0x0a, 0x03, 0x41, 0x63, 0x6b, 0x12, 0x27, 0x0a, 0x03, 0x61, 0x63, 0x6b, 0x18, 0x01, 0x20,
	0x01, 0x28, 0x0e, 0x32, 0x15, 0x2e, 0x47, 0x6f, 0x41, 0x75, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x53,
	0x79, 0x73, 0x74, 0x65, 0x6d, 0x2e, 0x41, 0x63, 0x6b, 0x73, 0x52, 0x03, 0x61, 0x63, 0x6b, 0x2a,
	0x38, 0x0a, 0x04, 0x41, 0x63, 0x6b, 0x73, 0x12, 0x0c, 0x0a, 0x08, 0x41, 0x43, 0x4b, 0x5f, 0x46,
	0x41, 0x49, 0x4c, 0x10, 0x00, 0x12, 0x0f, 0x0a, 0x0b, 0x41, 0x43, 0x4b, 0x5f, 0x53, 0x55, 0x43,
	0x43, 0x45, 0x53, 0x53, 0x10, 0x01, 0x12, 0x11, 0x0a, 0x0d, 0x41, 0x43, 0x4b, 0x5f, 0x45, 0x58,
	0x43, 0x45, 0x50, 0x54, 0x49, 0x4f, 0x4e, 0x10, 0x02, 0x32, 0x86, 0x01, 0x0a, 0x0d, 0x41, 0x75,
	0x63, 0x74, 0x69, 0x6f, 0x6e, 0x53, 0x79, 0x73, 0x74, 0x65, 0x6d, 0x12, 0x37, 0x0a, 0x03, 0x42,
	0x69, 0x64, 0x12, 0x18, 0x2e, 0x47, 0x6f, 0x41, 0x75, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x53, 0x79,
	0x73, 0x74, 0x65, 0x6d, 0x2e, 0x42, 0x69, 0x64, 0x50, 0x6f, 0x73, 0x74, 0x1a, 0x14, 0x2e, 0x47,
	0x6f, 0x41, 0x75, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x53, 0x79, 0x73, 0x74, 0x65, 0x6d, 0x2e, 0x41,
	0x63, 0x6b, 0x22, 0x00, 0x12, 0x3c, 0x0a, 0x06, 0x52, 0x65, 0x73, 0x75, 0x6c, 0x74, 0x12, 0x16,
	0x2e, 0x47, 0x6f, 0x41, 0x75, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x53, 0x79, 0x73, 0x74, 0x65, 0x6d,
	0x2e, 0x45, 0x6d, 0x70, 0x74, 0x79, 0x1a, 0x18, 0x2e, 0x47, 0x6f, 0x41, 0x75, 0x63, 0x74, 0x69,
	0x6f, 0x6e, 0x53, 0x79, 0x73, 0x74, 0x65, 0x6d, 0x2e, 0x4f, 0x75, 0x74, 0x63, 0x6f, 0x6d, 0x65,
	0x22, 0x00, 0x42, 0x46, 0x5a, 0x44, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d,
	0x2f, 0x48, 0x77, 0x35, 0x5f, 0x47, 0x6f, 0x41, 0x75, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x53, 0x79,
	0x73, 0x74, 0x65, 0x6d, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2f, 0x47, 0x6f, 0x41, 0x75, 0x63,
	0x74, 0x69, 0x6f, 0x6e, 0x53, 0x79, 0x73, 0x74, 0x65, 0x6d, 0x3b, 0x47, 0x6f, 0x41, 0x75, 0x63,
	0x74, 0x69, 0x6f, 0x6e, 0x53, 0x79, 0x73, 0x74, 0x65, 0x6d, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x33,
}

var (
	file_proto_Hw5_GoAuctionSystem_proto_rawDescOnce sync.Once
	file_proto_Hw5_GoAuctionSystem_proto_rawDescData = file_proto_Hw5_GoAuctionSystem_proto_rawDesc
)

func file_proto_Hw5_GoAuctionSystem_proto_rawDescGZIP() []byte {
	file_proto_Hw5_GoAuctionSystem_proto_rawDescOnce.Do(func() {
		file_proto_Hw5_GoAuctionSystem_proto_rawDescData = protoimpl.X.CompressGZIP(file_proto_Hw5_GoAuctionSystem_proto_rawDescData)
	})
	return file_proto_Hw5_GoAuctionSystem_proto_rawDescData
}

var file_proto_Hw5_GoAuctionSystem_proto_enumTypes = make([]protoimpl.EnumInfo, 1)
var file_proto_Hw5_GoAuctionSystem_proto_msgTypes = make([]protoimpl.MessageInfo, 4)
var file_proto_Hw5_GoAuctionSystem_proto_goTypes = []interface{}{
	(Acks)(0),       // 0: GoAuctionSystem.Acks
	(*BidPost)(nil), // 1: GoAuctionSystem.BidPost
	(*Empty)(nil),   // 2: GoAuctionSystem.Empty
	(*Outcome)(nil), // 3: GoAuctionSystem.Outcome
	(*Ack)(nil),     // 4: GoAuctionSystem.Ack
}
var file_proto_Hw5_GoAuctionSystem_proto_depIdxs = []int32{
	0, // 0: GoAuctionSystem.Ack.ack:type_name -> GoAuctionSystem.Acks
	1, // 1: GoAuctionSystem.AuctionSystem.Bid:input_type -> GoAuctionSystem.BidPost
	2, // 2: GoAuctionSystem.AuctionSystem.Result:input_type -> GoAuctionSystem.Empty
	4, // 3: GoAuctionSystem.AuctionSystem.Bid:output_type -> GoAuctionSystem.Ack
	3, // 4: GoAuctionSystem.AuctionSystem.Result:output_type -> GoAuctionSystem.Outcome
	3, // [3:5] is the sub-list for method output_type
	1, // [1:3] is the sub-list for method input_type
	1, // [1:1] is the sub-list for extension type_name
	1, // [1:1] is the sub-list for extension extendee
	0, // [0:1] is the sub-list for field type_name
}

func init() { file_proto_Hw5_GoAuctionSystem_proto_init() }
func file_proto_Hw5_GoAuctionSystem_proto_init() {
	if File_proto_Hw5_GoAuctionSystem_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_proto_Hw5_GoAuctionSystem_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*BidPost); i {
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
		file_proto_Hw5_GoAuctionSystem_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Empty); i {
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
		file_proto_Hw5_GoAuctionSystem_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Outcome); i {
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
		file_proto_Hw5_GoAuctionSystem_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Ack); i {
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
			RawDescriptor: file_proto_Hw5_GoAuctionSystem_proto_rawDesc,
			NumEnums:      1,
			NumMessages:   4,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_proto_Hw5_GoAuctionSystem_proto_goTypes,
		DependencyIndexes: file_proto_Hw5_GoAuctionSystem_proto_depIdxs,
		EnumInfos:         file_proto_Hw5_GoAuctionSystem_proto_enumTypes,
		MessageInfos:      file_proto_Hw5_GoAuctionSystem_proto_msgTypes,
	}.Build()
	File_proto_Hw5_GoAuctionSystem_proto = out.File
	file_proto_Hw5_GoAuctionSystem_proto_rawDesc = nil
	file_proto_Hw5_GoAuctionSystem_proto_goTypes = nil
	file_proto_Hw5_GoAuctionSystem_proto_depIdxs = nil
}
