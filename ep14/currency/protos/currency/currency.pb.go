// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.28.1
// 	protoc        v3.21.5
// source: protos/currency.proto

package currency

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

// Currencies is an enum which represents the allowed currencies for the API
type Currencies int32

const (
	Currencies_EUR Currencies = 0
	Currencies_USD Currencies = 1
	Currencies_JPY Currencies = 2
	Currencies_TRY Currencies = 3
)

// Enum value maps for Currencies.
var (
	Currencies_name = map[int32]string{
		0: "EUR",
		1: "USD",
		2: "JPY",
		3: "TRY",
	}
	Currencies_value = map[string]int32{
		"EUR": 0,
		"USD": 1,
		"JPY": 2,
		"TRY": 3,
	}
)

func (x Currencies) Enum() *Currencies {
	p := new(Currencies)
	*p = x
	return p
}

func (x Currencies) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (Currencies) Descriptor() protoreflect.EnumDescriptor {
	return file_protos_currency_proto_enumTypes[0].Descriptor()
}

func (Currencies) Type() protoreflect.EnumType {
	return &file_protos_currency_proto_enumTypes[0]
}

func (x Currencies) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use Currencies.Descriptor instead.
func (Currencies) EnumDescriptor() ([]byte, []int) {
	return file_protos_currency_proto_rawDescGZIP(), []int{0}
}

// RateRequest defines the request for a GetRate call
type RateRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// Base is the base currency code for the rate
	Base Currencies `protobuf:"varint,1,opt,name=Base,proto3,enum=currency.Currencies" json:"Base,omitempty"`
	// Destination is the destination currency code for the rate
	Destination Currencies `protobuf:"varint,2,opt,name=Destination,proto3,enum=currency.Currencies" json:"Destination,omitempty"`
}

func (x *RateRequest) Reset() {
	*x = RateRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_protos_currency_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *RateRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*RateRequest) ProtoMessage() {}

func (x *RateRequest) ProtoReflect() protoreflect.Message {
	mi := &file_protos_currency_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use RateRequest.ProtoReflect.Descriptor instead.
func (*RateRequest) Descriptor() ([]byte, []int) {
	return file_protos_currency_proto_rawDescGZIP(), []int{0}
}

func (x *RateRequest) GetBase() Currencies {
	if x != nil {
		return x.Base
	}
	return Currencies_EUR
}

func (x *RateRequest) GetDestination() Currencies {
	if x != nil {
		return x.Destination
	}
	return Currencies_EUR
}

// RateResponse is the response from a GetRate call, it contains
// rate which is a floating point number and can be used to convert
// between the two currencies specified in the request
type RateResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Rate float32 `protobuf:"fixed32,1,opt,name=rate,proto3" json:"rate,omitempty"`
}

func (x *RateResponse) Reset() {
	*x = RateResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_protos_currency_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *RateResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*RateResponse) ProtoMessage() {}

func (x *RateResponse) ProtoReflect() protoreflect.Message {
	mi := &file_protos_currency_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use RateResponse.ProtoReflect.Descriptor instead.
func (*RateResponse) Descriptor() ([]byte, []int) {
	return file_protos_currency_proto_rawDescGZIP(), []int{1}
}

func (x *RateResponse) GetRate() float32 {
	if x != nil {
		return x.Rate
	}
	return 0
}

var File_protos_currency_proto protoreflect.FileDescriptor

var file_protos_currency_proto_rawDesc = []byte{
	0x0a, 0x15, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x73, 0x2f, 0x63, 0x75, 0x72, 0x72, 0x65, 0x6e, 0x63,
	0x79, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x08, 0x63, 0x75, 0x72, 0x72, 0x65, 0x6e, 0x63,
	0x79, 0x22, 0x6f, 0x0a, 0x0b, 0x52, 0x61, 0x74, 0x65, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74,
	0x12, 0x28, 0x0a, 0x04, 0x42, 0x61, 0x73, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0e, 0x32, 0x14,
	0x2e, 0x63, 0x75, 0x72, 0x72, 0x65, 0x6e, 0x63, 0x79, 0x2e, 0x43, 0x75, 0x72, 0x72, 0x65, 0x6e,
	0x63, 0x69, 0x65, 0x73, 0x52, 0x04, 0x42, 0x61, 0x73, 0x65, 0x12, 0x36, 0x0a, 0x0b, 0x44, 0x65,
	0x73, 0x74, 0x69, 0x6e, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0e, 0x32,
	0x14, 0x2e, 0x63, 0x75, 0x72, 0x72, 0x65, 0x6e, 0x63, 0x79, 0x2e, 0x43, 0x75, 0x72, 0x72, 0x65,
	0x6e, 0x63, 0x69, 0x65, 0x73, 0x52, 0x0b, 0x44, 0x65, 0x73, 0x74, 0x69, 0x6e, 0x61, 0x74, 0x69,
	0x6f, 0x6e, 0x22, 0x22, 0x0a, 0x0c, 0x52, 0x61, 0x74, 0x65, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e,
	0x73, 0x65, 0x12, 0x12, 0x0a, 0x04, 0x72, 0x61, 0x74, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x02,
	0x52, 0x04, 0x72, 0x61, 0x74, 0x65, 0x2a, 0x30, 0x0a, 0x0a, 0x43, 0x75, 0x72, 0x72, 0x65, 0x6e,
	0x63, 0x69, 0x65, 0x73, 0x12, 0x07, 0x0a, 0x03, 0x45, 0x55, 0x52, 0x10, 0x00, 0x12, 0x07, 0x0a,
	0x03, 0x55, 0x53, 0x44, 0x10, 0x01, 0x12, 0x07, 0x0a, 0x03, 0x4a, 0x50, 0x59, 0x10, 0x02, 0x12,
	0x07, 0x0a, 0x03, 0x54, 0x52, 0x59, 0x10, 0x03, 0x32, 0x44, 0x0a, 0x08, 0x43, 0x75, 0x72, 0x72,
	0x65, 0x6e, 0x63, 0x79, 0x12, 0x38, 0x0a, 0x07, 0x47, 0x65, 0x74, 0x52, 0x61, 0x74, 0x65, 0x12,
	0x15, 0x2e, 0x63, 0x75, 0x72, 0x72, 0x65, 0x6e, 0x63, 0x79, 0x2e, 0x52, 0x61, 0x74, 0x65, 0x52,
	0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x16, 0x2e, 0x63, 0x75, 0x72, 0x72, 0x65, 0x6e, 0x63,
	0x79, 0x2e, 0x52, 0x61, 0x74, 0x65, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x42, 0x0b,
	0x5a, 0x09, 0x2f, 0x63, 0x75, 0x72, 0x72, 0x65, 0x6e, 0x63, 0x79, 0x62, 0x06, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x33,
}

var (
	file_protos_currency_proto_rawDescOnce sync.Once
	file_protos_currency_proto_rawDescData = file_protos_currency_proto_rawDesc
)

func file_protos_currency_proto_rawDescGZIP() []byte {
	file_protos_currency_proto_rawDescOnce.Do(func() {
		file_protos_currency_proto_rawDescData = protoimpl.X.CompressGZIP(file_protos_currency_proto_rawDescData)
	})
	return file_protos_currency_proto_rawDescData
}

var file_protos_currency_proto_enumTypes = make([]protoimpl.EnumInfo, 1)
var file_protos_currency_proto_msgTypes = make([]protoimpl.MessageInfo, 2)
var file_protos_currency_proto_goTypes = []interface{}{
	(Currencies)(0),      // 0: currency.Currencies
	(*RateRequest)(nil),  // 1: currency.RateRequest
	(*RateResponse)(nil), // 2: currency.RateResponse
}
var file_protos_currency_proto_depIdxs = []int32{
	0, // 0: currency.RateRequest.Base:type_name -> currency.Currencies
	0, // 1: currency.RateRequest.Destination:type_name -> currency.Currencies
	1, // 2: currency.Currency.GetRate:input_type -> currency.RateRequest
	2, // 3: currency.Currency.GetRate:output_type -> currency.RateResponse
	3, // [3:4] is the sub-list for method output_type
	2, // [2:3] is the sub-list for method input_type
	2, // [2:2] is the sub-list for extension type_name
	2, // [2:2] is the sub-list for extension extendee
	0, // [0:2] is the sub-list for field type_name
}

func init() { file_protos_currency_proto_init() }
func file_protos_currency_proto_init() {
	if File_protos_currency_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_protos_currency_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*RateRequest); i {
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
		file_protos_currency_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*RateResponse); i {
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
			RawDescriptor: file_protos_currency_proto_rawDesc,
			NumEnums:      1,
			NumMessages:   2,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_protos_currency_proto_goTypes,
		DependencyIndexes: file_protos_currency_proto_depIdxs,
		EnumInfos:         file_protos_currency_proto_enumTypes,
		MessageInfos:      file_protos_currency_proto_msgTypes,
	}.Build()
	File_protos_currency_proto = out.File
	file_protos_currency_proto_rawDesc = nil
	file_protos_currency_proto_goTypes = nil
	file_protos_currency_proto_depIdxs = nil
}
