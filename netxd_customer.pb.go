// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.28.1
// 	protoc        v3.12.4
// source: netxd_customer/netxd_customer.proto

package netxd_customer

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

type Details struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	CustomerId int64  `protobuf:"varint,1,opt,name=CustomerId,proto3" json:"CustomerId,omitempty"`
	Firstname  string `protobuf:"bytes,2,opt,name=Firstname,proto3" json:"Firstname,omitempty"`
	Lastname   string `protobuf:"bytes,3,opt,name=Lastname,proto3" json:"Lastname,omitempty"`
	BankId     int64  `protobuf:"varint,4,opt,name=BankId,proto3" json:"BankId,omitempty"`
	Balance    int64  `protobuf:"varint,5,opt,name=Balance,proto3" json:"Balance,omitempty"`
	CreatedAt  string `protobuf:"bytes,6,opt,name=CreatedAt,proto3" json:"CreatedAt,omitempty"`
	UpdatesAt  string `protobuf:"bytes,7,opt,name=UpdatesAt,proto3" json:"UpdatesAt,omitempty"`
	IsActive   bool   `protobuf:"varint,8,opt,name=IsActive,proto3" json:"IsActive,omitempty"`
}

func (x *Details) Reset() {
	*x = Details{}
	if protoimpl.UnsafeEnabled {
		mi := &file_netxd_customer_netxd_customer_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Details) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Details) ProtoMessage() {}

func (x *Details) ProtoReflect() protoreflect.Message {
	mi := &file_netxd_customer_netxd_customer_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Details.ProtoReflect.Descriptor instead.
func (*Details) Descriptor() ([]byte, []int) {
	return file_netxd_customer_netxd_customer_proto_rawDescGZIP(), []int{0}
}

func (x *Details) GetCustomerId() int64 {
	if x != nil {
		return x.CustomerId
	}
	return 0
}

func (x *Details) GetFirstname() string {
	if x != nil {
		return x.Firstname
	}
	return ""
}

func (x *Details) GetLastname() string {
	if x != nil {
		return x.Lastname
	}
	return ""
}

func (x *Details) GetBankId() int64 {
	if x != nil {
		return x.BankId
	}
	return 0
}

func (x *Details) GetBalance() int64 {
	if x != nil {
		return x.Balance
	}
	return 0
}

func (x *Details) GetCreatedAt() string {
	if x != nil {
		return x.CreatedAt
	}
	return ""
}

func (x *Details) GetUpdatesAt() string {
	if x != nil {
		return x.UpdatesAt
	}
	return ""
}

func (x *Details) GetIsActive() bool {
	if x != nil {
		return x.IsActive
	}
	return false
}

type DetailResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	CustomerId int64  `protobuf:"varint,1,opt,name=CustomerId,proto3" json:"CustomerId,omitempty"`
	CreatedAt  string `protobuf:"bytes,2,opt,name=CreatedAt,proto3" json:"CreatedAt,omitempty"`
}

func (x *DetailResponse) Reset() {
	*x = DetailResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_netxd_customer_netxd_customer_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *DetailResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*DetailResponse) ProtoMessage() {}

func (x *DetailResponse) ProtoReflect() protoreflect.Message {
	mi := &file_netxd_customer_netxd_customer_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use DetailResponse.ProtoReflect.Descriptor instead.
func (*DetailResponse) Descriptor() ([]byte, []int) {
	return file_netxd_customer_netxd_customer_proto_rawDescGZIP(), []int{1}
}

func (x *DetailResponse) GetCustomerId() int64 {
	if x != nil {
		return x.CustomerId
	}
	return 0
}

func (x *DetailResponse) GetCreatedAt() string {
	if x != nil {
		return x.CreatedAt
	}
	return ""
}

type Request struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	TransactionId int64 `protobuf:"varint,1,opt,name=Transaction_id,json=TransactionId,proto3" json:"Transaction_id,omitempty"`
	FromAccount   int64 `protobuf:"varint,2,opt,name=From_account,json=FromAccount,proto3" json:"From_account,omitempty"`
	ToAccount     int64 `protobuf:"varint,3,opt,name=To_account,json=ToAccount,proto3" json:"To_account,omitempty"`
	Amount        int64 `protobuf:"varint,4,opt,name=Amount,proto3" json:"Amount,omitempty"`
}

func (x *Request) Reset() {
	*x = Request{}
	if protoimpl.UnsafeEnabled {
		mi := &file_netxd_customer_netxd_customer_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Request) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Request) ProtoMessage() {}

func (x *Request) ProtoReflect() protoreflect.Message {
	mi := &file_netxd_customer_netxd_customer_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Request.ProtoReflect.Descriptor instead.
func (*Request) Descriptor() ([]byte, []int) {
	return file_netxd_customer_netxd_customer_proto_rawDescGZIP(), []int{2}
}

func (x *Request) GetTransactionId() int64 {
	if x != nil {
		return x.TransactionId
	}
	return 0
}

func (x *Request) GetFromAccount() int64 {
	if x != nil {
		return x.FromAccount
	}
	return 0
}

func (x *Request) GetToAccount() int64 {
	if x != nil {
		return x.ToAccount
	}
	return 0
}

func (x *Request) GetAmount() int64 {
	if x != nil {
		return x.Amount
	}
	return 0
}

type Response struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Message string `protobuf:"bytes,1,opt,name=message,proto3" json:"message,omitempty"`
}

func (x *Response) Reset() {
	*x = Response{}
	if protoimpl.UnsafeEnabled {
		mi := &file_netxd_customer_netxd_customer_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Response) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Response) ProtoMessage() {}

func (x *Response) ProtoReflect() protoreflect.Message {
	mi := &file_netxd_customer_netxd_customer_proto_msgTypes[3]
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
	return file_netxd_customer_netxd_customer_proto_rawDescGZIP(), []int{3}
}

func (x *Response) GetMessage() string {
	if x != nil {
		return x.Message
	}
	return ""
}

var File_netxd_customer_netxd_customer_proto protoreflect.FileDescriptor

var file_netxd_customer_netxd_customer_proto_rawDesc = []byte{
	0x0a, 0x23, 0x6e, 0x65, 0x74, 0x78, 0x64, 0x5f, 0x63, 0x75, 0x73, 0x74, 0x6f, 0x6d, 0x65, 0x72,
	0x2f, 0x6e, 0x65, 0x74, 0x78, 0x64, 0x5f, 0x63, 0x75, 0x73, 0x74, 0x6f, 0x6d, 0x65, 0x72, 0x2e,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x0e, 0x6e, 0x65, 0x74, 0x78, 0x64, 0x5f, 0x63, 0x75, 0x73,
	0x74, 0x6f, 0x6d, 0x65, 0x72, 0x22, 0xed, 0x01, 0x0a, 0x07, 0x44, 0x65, 0x74, 0x61, 0x69, 0x6c,
	0x73, 0x12, 0x1e, 0x0a, 0x0a, 0x43, 0x75, 0x73, 0x74, 0x6f, 0x6d, 0x65, 0x72, 0x49, 0x64, 0x18,
	0x01, 0x20, 0x01, 0x28, 0x03, 0x52, 0x0a, 0x43, 0x75, 0x73, 0x74, 0x6f, 0x6d, 0x65, 0x72, 0x49,
	0x64, 0x12, 0x1c, 0x0a, 0x09, 0x46, 0x69, 0x72, 0x73, 0x74, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x02,
	0x20, 0x01, 0x28, 0x09, 0x52, 0x09, 0x46, 0x69, 0x72, 0x73, 0x74, 0x6e, 0x61, 0x6d, 0x65, 0x12,
	0x1a, 0x0a, 0x08, 0x4c, 0x61, 0x73, 0x74, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x03, 0x20, 0x01, 0x28,
	0x09, 0x52, 0x08, 0x4c, 0x61, 0x73, 0x74, 0x6e, 0x61, 0x6d, 0x65, 0x12, 0x16, 0x0a, 0x06, 0x42,
	0x61, 0x6e, 0x6b, 0x49, 0x64, 0x18, 0x04, 0x20, 0x01, 0x28, 0x03, 0x52, 0x06, 0x42, 0x61, 0x6e,
	0x6b, 0x49, 0x64, 0x12, 0x18, 0x0a, 0x07, 0x42, 0x61, 0x6c, 0x61, 0x6e, 0x63, 0x65, 0x18, 0x05,
	0x20, 0x01, 0x28, 0x03, 0x52, 0x07, 0x42, 0x61, 0x6c, 0x61, 0x6e, 0x63, 0x65, 0x12, 0x1c, 0x0a,
	0x09, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x64, 0x41, 0x74, 0x18, 0x06, 0x20, 0x01, 0x28, 0x09,
	0x52, 0x09, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x64, 0x41, 0x74, 0x12, 0x1c, 0x0a, 0x09, 0x55,
	0x70, 0x64, 0x61, 0x74, 0x65, 0x73, 0x41, 0x74, 0x18, 0x07, 0x20, 0x01, 0x28, 0x09, 0x52, 0x09,
	0x55, 0x70, 0x64, 0x61, 0x74, 0x65, 0x73, 0x41, 0x74, 0x12, 0x1a, 0x0a, 0x08, 0x49, 0x73, 0x41,
	0x63, 0x74, 0x69, 0x76, 0x65, 0x18, 0x08, 0x20, 0x01, 0x28, 0x08, 0x52, 0x08, 0x49, 0x73, 0x41,
	0x63, 0x74, 0x69, 0x76, 0x65, 0x22, 0x4e, 0x0a, 0x0e, 0x44, 0x65, 0x74, 0x61, 0x69, 0x6c, 0x52,
	0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x1e, 0x0a, 0x0a, 0x43, 0x75, 0x73, 0x74, 0x6f,
	0x6d, 0x65, 0x72, 0x49, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x03, 0x52, 0x0a, 0x43, 0x75, 0x73,
	0x74, 0x6f, 0x6d, 0x65, 0x72, 0x49, 0x64, 0x12, 0x1c, 0x0a, 0x09, 0x43, 0x72, 0x65, 0x61, 0x74,
	0x65, 0x64, 0x41, 0x74, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x09, 0x43, 0x72, 0x65, 0x61,
	0x74, 0x65, 0x64, 0x41, 0x74, 0x22, 0x8a, 0x01, 0x0a, 0x07, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73,
	0x74, 0x12, 0x25, 0x0a, 0x0e, 0x54, 0x72, 0x61, 0x6e, 0x73, 0x61, 0x63, 0x74, 0x69, 0x6f, 0x6e,
	0x5f, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x03, 0x52, 0x0d, 0x54, 0x72, 0x61, 0x6e, 0x73,
	0x61, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x49, 0x64, 0x12, 0x21, 0x0a, 0x0c, 0x46, 0x72, 0x6f, 0x6d,
	0x5f, 0x61, 0x63, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x18, 0x02, 0x20, 0x01, 0x28, 0x03, 0x52, 0x0b,
	0x46, 0x72, 0x6f, 0x6d, 0x41, 0x63, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x12, 0x1d, 0x0a, 0x0a, 0x54,
	0x6f, 0x5f, 0x61, 0x63, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x18, 0x03, 0x20, 0x01, 0x28, 0x03, 0x52,
	0x09, 0x54, 0x6f, 0x41, 0x63, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x12, 0x16, 0x0a, 0x06, 0x41, 0x6d,
	0x6f, 0x75, 0x6e, 0x74, 0x18, 0x04, 0x20, 0x01, 0x28, 0x03, 0x52, 0x06, 0x41, 0x6d, 0x6f, 0x75,
	0x6e, 0x74, 0x22, 0x24, 0x0a, 0x08, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x18,
	0x0a, 0x07, 0x6d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52,
	0x07, 0x6d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x32, 0x9b, 0x01, 0x0a, 0x0f, 0x43, 0x75, 0x73,
	0x74, 0x6f, 0x6d, 0x65, 0x72, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x12, 0x49, 0x0a, 0x0e,
	0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x43, 0x75, 0x73, 0x74, 0x6f, 0x6d, 0x65, 0x72, 0x12, 0x17,
	0x2e, 0x6e, 0x65, 0x74, 0x78, 0x64, 0x5f, 0x63, 0x75, 0x73, 0x74, 0x6f, 0x6d, 0x65, 0x72, 0x2e,
	0x44, 0x65, 0x74, 0x61, 0x69, 0x6c, 0x73, 0x1a, 0x1e, 0x2e, 0x6e, 0x65, 0x74, 0x78, 0x64, 0x5f,
	0x63, 0x75, 0x73, 0x74, 0x6f, 0x6d, 0x65, 0x72, 0x2e, 0x44, 0x65, 0x74, 0x61, 0x69, 0x6c, 0x52,
	0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x3d, 0x0a, 0x08, 0x54, 0x72, 0x61, 0x6e, 0x73,
	0x66, 0x65, 0x72, 0x12, 0x17, 0x2e, 0x6e, 0x65, 0x74, 0x78, 0x64, 0x5f, 0x63, 0x75, 0x73, 0x74,
	0x6f, 0x6d, 0x65, 0x72, 0x2e, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x18, 0x2e, 0x6e,
	0x65, 0x74, 0x78, 0x64, 0x5f, 0x63, 0x75, 0x73, 0x74, 0x6f, 0x6d, 0x65, 0x72, 0x2e, 0x52, 0x65,
	0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x42, 0x2d, 0x5a, 0x2b, 0x68, 0x74, 0x74, 0x70, 0x73, 0x3a,
	0x2f, 0x2f, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x54, 0x68, 0x61,
	0x73, 0x68, 0x6d, 0x69, 0x30, 0x33, 0x2f, 0x6e, 0x65, 0x74, 0x78, 0x64, 0x5f, 0x63, 0x75, 0x73,
	0x74, 0x6f, 0x6d, 0x65, 0x72, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_netxd_customer_netxd_customer_proto_rawDescOnce sync.Once
	file_netxd_customer_netxd_customer_proto_rawDescData = file_netxd_customer_netxd_customer_proto_rawDesc
)

func file_netxd_customer_netxd_customer_proto_rawDescGZIP() []byte {
	file_netxd_customer_netxd_customer_proto_rawDescOnce.Do(func() {
		file_netxd_customer_netxd_customer_proto_rawDescData = protoimpl.X.CompressGZIP(file_netxd_customer_netxd_customer_proto_rawDescData)
	})
	return file_netxd_customer_netxd_customer_proto_rawDescData
}

var file_netxd_customer_netxd_customer_proto_msgTypes = make([]protoimpl.MessageInfo, 4)
var file_netxd_customer_netxd_customer_proto_goTypes = []interface{}{
	(*Details)(nil),        // 0: netxd_customer.Details
	(*DetailResponse)(nil), // 1: netxd_customer.DetailResponse
	(*Request)(nil),        // 2: netxd_customer.Request
	(*Response)(nil),       // 3: netxd_customer.Response
}
var file_netxd_customer_netxd_customer_proto_depIdxs = []int32{
	0, // 0: netxd_customer.CustomerService.CreateCustomer:input_type -> netxd_customer.Details
	2, // 1: netxd_customer.CustomerService.Transfer:input_type -> netxd_customer.Request
	1, // 2: netxd_customer.CustomerService.CreateCustomer:output_type -> netxd_customer.DetailResponse
	3, // 3: netxd_customer.CustomerService.Transfer:output_type -> netxd_customer.Response
	2, // [2:4] is the sub-list for method output_type
	0, // [0:2] is the sub-list for method input_type
	0, // [0:0] is the sub-list for extension type_name
	0, // [0:0] is the sub-list for extension extendee
	0, // [0:0] is the sub-list for field type_name
}

func init() { file_netxd_customer_netxd_customer_proto_init() }
func file_netxd_customer_netxd_customer_proto_init() {
	if File_netxd_customer_netxd_customer_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_netxd_customer_netxd_customer_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Details); i {
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
		file_netxd_customer_netxd_customer_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*DetailResponse); i {
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
		file_netxd_customer_netxd_customer_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Request); i {
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
		file_netxd_customer_netxd_customer_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
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
	}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_netxd_customer_netxd_customer_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   4,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_netxd_customer_netxd_customer_proto_goTypes,
		DependencyIndexes: file_netxd_customer_netxd_customer_proto_depIdxs,
		MessageInfos:      file_netxd_customer_netxd_customer_proto_msgTypes,
	}.Build()
	File_netxd_customer_netxd_customer_proto = out.File
	file_netxd_customer_netxd_customer_proto_rawDesc = nil
	file_netxd_customer_netxd_customer_proto_goTypes = nil
	file_netxd_customer_netxd_customer_proto_depIdxs = nil
}
