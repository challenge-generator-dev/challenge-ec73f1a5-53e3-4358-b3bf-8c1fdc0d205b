package proto

import (
	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
	constpref "google.golang.org/genproto/googleapis/type/date"
	durationpb "google.golang.org/protobuf/types/known/durationpb"
	timestamppb "google.golang.org/protobuf/types/known/timestamppb"
	reflect "reflect"
	sync "sync"
	unsafe "unsafe"
)

const (
	File_invoice_proto_name          = "invoice.proto"
	File_invoice_proto_rawDesc       = "\n\rinvoice.proto\x12\x05proto\x1a\x1fgoogle/protobuf/timestamp.proto" +
		"\xa0\x01\x03\xa8\x01\x03\xe8\x01\x03\x12\x8b\x01\n\x17CreateInvoiceRequest\x12\x12\n\x0bcustomer_id" +
		"\x18\x01 \x01(\tR\x0bcustomerId\x12\n\n\x06amount\x18\x02 \x01(\x01R\x06amount\x12\x31\n\x08due_date\x18\x03 " +
		"\x01(\x0b2\x1a.google.protobuf.TimestampR\x07dueDate\x12\x33\n\nissue_date\x18\x04 " +
		"\x01(\x0b2\x1a.google.protobuf.TimestampR\tissueDate\x12\x14\n\x0bdescription\x18\x05 " +
		"\x01(\tR\x0bdescription\x12\x8d\x01\n\x18CreateInvoiceResponse\x12-\n\x07invoice\x18\x01 " +
		"\x01(\x0b2\x14.proto.InvoiceResponseR\x07invoice\x12\x10\n\x05error\x18\x02 \x01(\tR\x05error" +
		"\x12\x80\x01\n\x14GetInvoiceRequest\x12\n\n\x02id\x18\x01 \x01(\tR\x02id\x12\x81\x01\n\x15GetInvoiceResponse" +
		"\x12-\n\x07invoice\x18\x01 \x01(\x0b2\x14.proto.InvoiceResponseR\x07invoice\x12\x10\n\x05error\x18\x02 " +
		"\x01(\tR\x05error\x12\x8a\x01\n\x16UpdateInvoiceRequest\x12\n\n\x02id\x18\x01 \x01(\tR\x02id\x12\n\n\x06amount" +
		"\x18\x02 \x01(\x01R\x06amount\x12\x10\n\x06status\x18\x03 \x01(\tR\x06status\x12\x8b\x01\n\x17UpdateInvoiceResponse" +
		"\x12-\n\x07invoice\x18\x01 \x01(\x0b2\x14.proto.InvoiceResponseR\x07invoice\x12\x10\n\x05error\x18\x02 " +
		"\x01(\tR\x05error\x12\x8a\x01\n\x0eInvoiceResponse\x12\n\n\x02id\x18\x01 \x01(\tR\x02id\x12\x14\n\x0bcustomer_id" +
		"\x18\x02 \x01(\tR\x0bcustomerId\x12\n\n\x06amount\x18\x03 \x01(\x01R\x06amount\x12\x10\n\x06status\x18\x04 " +
		"\x01(\tR\x06status\x12\x31\n\x08due_date\x18\x05 \x01(\x0b2\x1a.google.protobuf.TimestampR\x07dueDate" +
		"\x12\x33\n\nissue_date\x18\x06 \x01(\x0b2\x1a.google.protobuf.TimestampR\tissueDate\x12\x14\n\x0bdescription" +
		"\x18\x07 \x01(\tR\x0bdescription\x12\x80\x01\n\x11HealthCheckRequest\x12\x14\n\x07service\x18\x01 " +
		"\x01(\tR\x07service\x12\x8d\x01\n\x12HealthCheckResponse\x12-\n\x06status\x18\x01 " +
		"\x01(\x0e2\x1d.proto.HealthCheckResponse.ServingStatusR\x06status" +
		"\xe8\x01\x02\x12\n\n\x0eInvoiceService\x12\x46\n\x0eCreateInvoice\x12\x1a.proto.CreateInvoiceRequest" +
		"(\x0b2\x1b.proto.CreateInvoiceResponseR\x0eCreateInvoice\x12\x41\n\nGetInvoice\x12\x17.proto.GetInvoiceRequest" +
		"(\x0b2\x18.proto.GetInvoiceResponseR\nGetInvoice\x12\x43\n\x0cUpdateInvoice\x12\x19.proto.UpdateInvoiceRequest" +
		"(\x0b2\x1a.proto.UpdateInvoiceResponseR\x0cUpdateInvoice\x12\x1e\n\x05Health\x12\x14.proto.HealthCheckRequest" +
		"(\x0b2\x15.proto.HealthCheckResponseR\x05Health\x12\x1e\n\x05Watch\x12\x14.proto.HealthCheckRequest" +
		"\x12\x15.proto.HealthCheckResponseB\xca\x01\x1bGo package invoice-service/pkg/protob\x06proto3"
	File_invoice_proto_reflectDescFileInfo = []string{
		"invoice.proto",
	}
)

var File_invoice_proto protoreflect.FileDescriptor

func init() { protoimpl.RegisterFile(File_invoice_proto_name, File_invoice_proto_rawDesc, File_invoice_proto_reflectDescFileInfo) }

var File_invoice_proto_rawDescOnce sync.Once

func File_invoice_proto_rawDescGZIP() []byte {
	File_invoice_proto_rawDescOnce.Do(func() {
		File_invoice_proto_rawDesc = []byte{
			0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x8c, 0x54, 0xcd, 0x6e, 0xe3, 0x36,
			0x10, 0xc6, 0x5e, 0x40, 0x12, 0x8e, 0x9c, 0x84, 0x4a, 0x93, 0x4e, 0x92, 0x4b, 0x93, 0x52, 0x24,
			0x12, 0x65, 0x25, 0x24, 0x65, 0x25, 0xb2, 0x4a, 0x89, 0x52, 0x48, 0xa5, 0x28, 0x4b, 0xa5, 0x6c,
			0x28, 0x4a, 0x81, 0xd4, 0x6a, 0x28, 0x94, 0x0a, 0x52, 0x5c, 0x90, 0x6a, 0x90, 0x02, 0x52, 0x5c,
			0x90, 0x6a, 0x90, 0x02, 0x52, 0x5c, 0x90, 0x6a, 0x90, 0x02, 0x52, 0x5c, 0x90, 0x6a, 0x90, 0x02,
			0x52, 0x5c, 0x90, 0x6a, 0x90, 0x02, 0x52, 0x5c, 0x90, 0x6a, 0x90, 0x82, 0x50, 0x2d, 0x25, 0x4a,
			0x4a, 0x6d, 0xf1, 0x44, 0x60, 0x71, 0xe5, 0x97, 0xdc, 0x23, 0xb1, 0x89, 0x5c, 0x22, 0x24, 0x95,
			0x4a, 0x32, 0xb9, 0x40, 0x90, 0x90, 0x92, 0x4a, 0x32, 0xb9, 0x40, 0x90, 0x92, 0x4a, 0x32, 0xb9,
			0x40, 0x90, 0x92, 0x4a, 0x32, 0xb9, 0x40, 0x90, 0x92, 0x4a, 0x32, 0xb9, 0x40, 0x90, 0x92, 0x4a,
			0x32, 0xb9, 0x40, 0x90, 0x92, 0x4a, 0x32, 0xb9, 0x40, 0x90, 0x92, 0x4a, 0x32, 0xb9, 0x40, 0x90,
			0x92, 0x4a, 0x12, 0x23, 0x21, 0x25, 0x95, 0x64, 0x72, 0x81, 0x20, 0x25, 0x95, 0x64, 0x72, 0x81,
			0x20, 0x25, 0x95, 0x64, 0x72, 0x81, 0x20, 0x25, 0x95, 0x64, 0x72, 0x81, 0x20, 0x25, 0x95, 0x64,
			0x72, 0x81, 0x20, 0x25, 0x95, 0x64, 0x72, 0x81, 0x20, 0x25, 0x95, 0x64, 0x72, 0x81, 0x20, 0x25,
			0x95, 0x64, 0x72, 0x81, 0x20, 0x25, 0x95, 0x64, 0x72, 0x81, 0x20, 0x25, 0x95, 0x64, 0x72, 0x81,
			0x20, 0x25, 0x95, 0x64, 0x72, 0x81, 0x20, 0x25, 0x95, 0x64, 0x72, 0x81, 0x20, 0x25, 0x95, 0x64,
			0x72, 0x81, 0x20, 0x25, 0x95, 0x64, 0x72, 0x81, 0x20, 0x25, 0x95, 0x64, 0x72, 0x81, 0x20, 0x25,
			0x95, 0xa4, 0x48, 0x52, 0x52, 0x49, 0x93, 0x03, 0x04, 0x23, 0x23, 0x23, 0x23, 0x23, 0x23, 0x23,
			0x23, 0x23, 0x23, 0x23, 0x23, 0x23, 0x23, 0x23, 0x23, 0x23, 0x23, 0x23, 0x23, 0x23, 0x23, 0x23,
			0x23, 0x23, 0x23, 0x23, 0x23, 0x23, 0x23, 0x23, 0x23, 0x23, 0x23, 0x23, 0x23, 0x23, 0x23, 0x23,
			0x23, 0x23, 0x23, 0x23, 0x23, 0x23, 0x23, 0x23, 0x23, 0x23, 0x23, 0x23, 0x23, 0x23, 0x23, 0x23,
			0x23, 0x23, 0x23, 0x23, 0x23, 0x23, 0x23, 0x23, 0x23, 0x23, 0x23, 0x23, 0x23, 0x23, 0x23, 0x23,
			0x23, 0x23, 0x23, 0x23, 0x23, 0x23, 0x23, 0x23, 0x23, 0x23, 0x23, 0x23, 0x23, 0x23, 0x23, 0x23,
			0xa3, 0xfd, 0x0e, 0x6e, 0x80, 0x8c, 0x8c, 0x4e, 0x40, 0x22, 0x25, 0x0e, 0x2e, 0x46, 0x4a, 0x9c,
			0x80, 0x9c, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00,
			0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00,
			0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00,
			0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00,
			0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00,
			0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00,
			0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00,
			0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00,
			0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00,
		}
	})
	return File_invoice_proto_rawDesc
}

type CreateInvoiceRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields unknownFields

	CustomerID  string                 `protobuf:"bytes,1,opt,name=customer_id,json=customerId,proto3" json:"customer_id,omitempty"`
	Amount      float64                `protobuf:"fixed64,2,opt,name=amount,proto3" json:"amount,omitempty"`
	DueDate     *timestamppb.Timestamp `protobuf:"bytes,3,opt,name=due_date,json=dueDate,proto3" json:"due_date,omitempty"`
	IssueDate   *timestamppb.Timestamp `protobuf:"bytes,4,opt,name=issue_date,json=issueDate,proto3" json:"issue_date,omitempty"`
	Description string                 `protobuf:"bytes,5,opt,name=description,proto3" json:"description,omitempty"`
}

func (x *CreateInvoiceRequest) Reset() {
	*x = CreateInvoiceRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_invoice_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *CreateInvoiceRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CreateInvoiceRequest) ProtoMessage() {}

func (x *CreateInvoiceRequest) GetCustomerID() string {
	if x != nil {
		return x.CustomerID
	}
	return ""
}

func (x *CreateInvoiceRequest) GetAmount() float64 {
	if x != nil {
		return x.Amount
	}
	return 0
}

func (x *CreateInvoiceRequest) GetDueDate() *timestamppb.Timestamp {
	if x != nil {
		return x.DueDate
	}
	return nil
}

func (x *CreateInvoiceRequest) GetIssueDate() *timestamppb.Timestamp {
	if x != nil {
		return x.IssueDate
	}
	return nil
}

func (x *CreateInvoiceRequest) GetDescription() string {
	if x != nil {
		return x.Description
	}
	return ""
}

type InvoiceResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields unknownFields

	Id          string                  `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
	CustomerId  string                  `protobuf:"bytes,2,opt,name=customer_id,json=customerId,proto3" json:"customer_id,omitempty"`
	Amount      float64                 `protobuf:"fixed64,3,opt,name=amount,proto3" json:"amount,omitempty"`
	Status      string                  `protobuf:"bytes,4,opt,name=status,proto3" json:"status,omitempty"`
	DueDate     *timestamppb.Timestamp  `protobuf:"bytes,5,opt,name=due_date,json=dueDate,proto3" json:"due_date,omitempty"`
	IssueDate   *timestamppb.Timestamp  `protobuf:"bytes,6,opt,name=issue_date,json=issueDate,proto3" json:"issue_date,omitempty"`
	Description string                  `protobuf:"bytes,7,opt,name=description,proto3" json:"description,omitempty"`
}

func (x *InvoiceResponse) Reset() {
	*x = InvoiceResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_invoice_proto_msgTypes[5]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *InvoiceResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*InvoiceResponse) ProtoMessage() {}

func (x *InvoiceResponse) GetId() string {
	if x != nil {
		return x.Id
	}
	return ""
}

func (x *InvoiceResponse) GetCustomerId() string {
	if x != nil {
		return x.CustomerId
	}
	return ""
}

func (x *InvoiceResponse) GetAmount() float64 {
	if x != nil {
		return x.Amount
	}
	return 0
}

func (x *InvoiceResponse) GetStatus() string {
	if x != nil {
		return x.Status
	}
	return ""
}

func (x *InvoiceResponse) GetDueDate() *timestamppb.Timestamp {
	if x != nil {
		return x.DueDate
	}
	return nil
}

func (x *InvoiceResponse) GetIssueDate() *timestamppb.Timestamp {
	if x != nil {
		return x.IssueDate
	}
	return nil
}

func (x *InvoiceResponse) GetDescription() string {
	if x != nil {
		return x.Description
	}
	return ""
}

var file_invoice_proto_msgTypes = make([]protoimpl.MessageInfo, 8)
var file_invoice_proto_goTypes = []interface{}{
	(*CreateInvoiceRequest)(nil),
	(*CreateInvoiceResponse)(nil),
	(*GetInvoiceRequest)(nil),
	(*GetInvoiceResponse)(nil),
	(*UpdateInvoiceRequest)(nil),
	(*UpdateInvoiceResponse)(nil),
	(*InvoiceResponse)(nil),
	(*HealthCheckRequest)(nil),
	(*HealthCheckResponse)(nil),
}

var file_invoice_proto_rawDesc = protoimpl.RawDesc{
	TX:     rawDesc,
	GoTypes: file_invoice_proto_goTypes,
	MessageInfos: file_invoice_proto_msgTypes,
}

func init() { proto.RegisterFile(File_invoice_proto_name, file_invoice_proto_rawDesc) }

var File_invoice_proto protoreflect.FileDescriptor = file_invoice_proto_rawDesc