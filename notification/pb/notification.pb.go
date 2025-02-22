// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.35.1
// 	protoc        v5.28.2
// source: pb/notification.proto

package pb

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

type PhoneNotification struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	PhoneNo string `protobuf:"bytes,1,opt,name=PhoneNo,proto3" json:"PhoneNo,omitempty"`
	Text    string `protobuf:"bytes,2,opt,name=text,proto3" json:"text,omitempty"`
}

func (x *PhoneNotification) Reset() {
	*x = PhoneNotification{}
	mi := &file_pb_notification_proto_msgTypes[0]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *PhoneNotification) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*PhoneNotification) ProtoMessage() {}

func (x *PhoneNotification) ProtoReflect() protoreflect.Message {
	mi := &file_pb_notification_proto_msgTypes[0]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use PhoneNotification.ProtoReflect.Descriptor instead.
func (*PhoneNotification) Descriptor() ([]byte, []int) {
	return file_pb_notification_proto_rawDescGZIP(), []int{0}
}

func (x *PhoneNotification) GetPhoneNo() string {
	if x != nil {
		return x.PhoneNo
	}
	return ""
}

func (x *PhoneNotification) GetText() string {
	if x != nil {
		return x.Text
	}
	return ""
}

type EmailNotification struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Email string `protobuf:"bytes,1,opt,name=Email,proto3" json:"Email,omitempty"`
	Text  string `protobuf:"bytes,2,opt,name=text,proto3" json:"text,omitempty"`
}

func (x *EmailNotification) Reset() {
	*x = EmailNotification{}
	mi := &file_pb_notification_proto_msgTypes[1]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *EmailNotification) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*EmailNotification) ProtoMessage() {}

func (x *EmailNotification) ProtoReflect() protoreflect.Message {
	mi := &file_pb_notification_proto_msgTypes[1]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use EmailNotification.ProtoReflect.Descriptor instead.
func (*EmailNotification) Descriptor() ([]byte, []int) {
	return file_pb_notification_proto_rawDescGZIP(), []int{1}
}

func (x *EmailNotification) GetEmail() string {
	if x != nil {
		return x.Email
	}
	return ""
}

func (x *EmailNotification) GetText() string {
	if x != nil {
		return x.Text
	}
	return ""
}

type NotificationStatus struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Message string `protobuf:"bytes,1,opt,name=message,proto3" json:"message,omitempty"`
}

func (x *NotificationStatus) Reset() {
	*x = NotificationStatus{}
	mi := &file_pb_notification_proto_msgTypes[2]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *NotificationStatus) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*NotificationStatus) ProtoMessage() {}

func (x *NotificationStatus) ProtoReflect() protoreflect.Message {
	mi := &file_pb_notification_proto_msgTypes[2]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use NotificationStatus.ProtoReflect.Descriptor instead.
func (*NotificationStatus) Descriptor() ([]byte, []int) {
	return file_pb_notification_proto_rawDescGZIP(), []int{2}
}

func (x *NotificationStatus) GetMessage() string {
	if x != nil {
		return x.Message
	}
	return ""
}

var File_pb_notification_proto protoreflect.FileDescriptor

var file_pb_notification_proto_rawDesc = []byte{
	0x0a, 0x15, 0x70, 0x62, 0x2f, 0x6e, 0x6f, 0x74, 0x69, 0x66, 0x69, 0x63, 0x61, 0x74, 0x69, 0x6f,
	0x6e, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x0c, 0x6e, 0x6f, 0x74, 0x69, 0x66, 0x69, 0x63,
	0x61, 0x74, 0x69, 0x6f, 0x6e, 0x22, 0x41, 0x0a, 0x11, 0x50, 0x68, 0x6f, 0x6e, 0x65, 0x4e, 0x6f,
	0x74, 0x69, 0x66, 0x69, 0x63, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x12, 0x18, 0x0a, 0x07, 0x50, 0x68,
	0x6f, 0x6e, 0x65, 0x4e, 0x6f, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x07, 0x50, 0x68, 0x6f,
	0x6e, 0x65, 0x4e, 0x6f, 0x12, 0x12, 0x0a, 0x04, 0x74, 0x65, 0x78, 0x74, 0x18, 0x02, 0x20, 0x01,
	0x28, 0x09, 0x52, 0x04, 0x74, 0x65, 0x78, 0x74, 0x22, 0x3d, 0x0a, 0x11, 0x45, 0x6d, 0x61, 0x69,
	0x6c, 0x4e, 0x6f, 0x74, 0x69, 0x66, 0x69, 0x63, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x12, 0x14, 0x0a,
	0x05, 0x45, 0x6d, 0x61, 0x69, 0x6c, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x45, 0x6d,
	0x61, 0x69, 0x6c, 0x12, 0x12, 0x0a, 0x04, 0x74, 0x65, 0x78, 0x74, 0x18, 0x02, 0x20, 0x01, 0x28,
	0x09, 0x52, 0x04, 0x74, 0x65, 0x78, 0x74, 0x22, 0x2e, 0x0a, 0x12, 0x4e, 0x6f, 0x74, 0x69, 0x66,
	0x69, 0x63, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x53, 0x74, 0x61, 0x74, 0x75, 0x73, 0x12, 0x18, 0x0a,
	0x07, 0x6d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x07,
	0x6d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x32, 0x9f, 0x02, 0x0a, 0x0c, 0x4e, 0x6f, 0x74, 0x69,
	0x66, 0x69, 0x63, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x12, 0x5c, 0x0a, 0x15, 0x53, 0x65, 0x6e, 0x64,
	0x45, 0x6d, 0x61, 0x69, 0x6c, 0x4e, 0x6f, 0x74, 0x69, 0x66, 0x69, 0x63, 0x61, 0x74, 0x69, 0x6f,
	0x6e, 0x12, 0x1f, 0x2e, 0x6e, 0x6f, 0x74, 0x69, 0x66, 0x69, 0x63, 0x61, 0x74, 0x69, 0x6f, 0x6e,
	0x2e, 0x45, 0x6d, 0x61, 0x69, 0x6c, 0x4e, 0x6f, 0x74, 0x69, 0x66, 0x69, 0x63, 0x61, 0x74, 0x69,
	0x6f, 0x6e, 0x1a, 0x20, 0x2e, 0x6e, 0x6f, 0x74, 0x69, 0x66, 0x69, 0x63, 0x61, 0x74, 0x69, 0x6f,
	0x6e, 0x2e, 0x4e, 0x6f, 0x74, 0x69, 0x66, 0x69, 0x63, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x53, 0x74,
	0x61, 0x74, 0x75, 0x73, 0x22, 0x00, 0x12, 0x5c, 0x0a, 0x15, 0x53, 0x65, 0x6e, 0x64, 0x50, 0x68,
	0x6f, 0x6e, 0x65, 0x4e, 0x6f, 0x74, 0x69, 0x66, 0x69, 0x63, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x12,
	0x1f, 0x2e, 0x6e, 0x6f, 0x74, 0x69, 0x66, 0x69, 0x63, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x2e, 0x50,
	0x68, 0x6f, 0x6e, 0x65, 0x4e, 0x6f, 0x74, 0x69, 0x66, 0x69, 0x63, 0x61, 0x74, 0x69, 0x6f, 0x6e,
	0x1a, 0x20, 0x2e, 0x6e, 0x6f, 0x74, 0x69, 0x66, 0x69, 0x63, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x2e,
	0x4e, 0x6f, 0x74, 0x69, 0x66, 0x69, 0x63, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x53, 0x74, 0x61, 0x74,
	0x75, 0x73, 0x22, 0x00, 0x12, 0x53, 0x0a, 0x0a, 0x53, 0x65, 0x6e, 0x64, 0x53, 0x74, 0x61, 0x74,
	0x75, 0x73, 0x12, 0x1f, 0x2e, 0x6e, 0x6f, 0x74, 0x69, 0x66, 0x69, 0x63, 0x61, 0x74, 0x69, 0x6f,
	0x6e, 0x2e, 0x45, 0x6d, 0x61, 0x69, 0x6c, 0x4e, 0x6f, 0x74, 0x69, 0x66, 0x69, 0x63, 0x61, 0x74,
	0x69, 0x6f, 0x6e, 0x1a, 0x20, 0x2e, 0x6e, 0x6f, 0x74, 0x69, 0x66, 0x69, 0x63, 0x61, 0x74, 0x69,
	0x6f, 0x6e, 0x2e, 0x4e, 0x6f, 0x74, 0x69, 0x66, 0x69, 0x63, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x53,
	0x74, 0x61, 0x74, 0x75, 0x73, 0x22, 0x00, 0x28, 0x01, 0x42, 0x2a, 0x5a, 0x28, 0x67, 0x69, 0x74,
	0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x6f, 0x72, 0x64, 0x65, 0x72, 0x5f, 0x73, 0x65,
	0x72, 0x76, 0x69, 0x63, 0x65, 0x2f, 0x6e, 0x6f, 0x74, 0x69, 0x66, 0x69, 0x63, 0x61, 0x74, 0x69,
	0x6f, 0x6e, 0x2f, 0x70, 0x62, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_pb_notification_proto_rawDescOnce sync.Once
	file_pb_notification_proto_rawDescData = file_pb_notification_proto_rawDesc
)

func file_pb_notification_proto_rawDescGZIP() []byte {
	file_pb_notification_proto_rawDescOnce.Do(func() {
		file_pb_notification_proto_rawDescData = protoimpl.X.CompressGZIP(file_pb_notification_proto_rawDescData)
	})
	return file_pb_notification_proto_rawDescData
}

var file_pb_notification_proto_msgTypes = make([]protoimpl.MessageInfo, 3)
var file_pb_notification_proto_goTypes = []any{
	(*PhoneNotification)(nil),  // 0: notification.PhoneNotification
	(*EmailNotification)(nil),  // 1: notification.EmailNotification
	(*NotificationStatus)(nil), // 2: notification.NotificationStatus
}
var file_pb_notification_proto_depIdxs = []int32{
	1, // 0: notification.Notification.SendEmailNotification:input_type -> notification.EmailNotification
	0, // 1: notification.Notification.SendPhoneNotification:input_type -> notification.PhoneNotification
	1, // 2: notification.Notification.SendStatus:input_type -> notification.EmailNotification
	2, // 3: notification.Notification.SendEmailNotification:output_type -> notification.NotificationStatus
	2, // 4: notification.Notification.SendPhoneNotification:output_type -> notification.NotificationStatus
	2, // 5: notification.Notification.SendStatus:output_type -> notification.NotificationStatus
	3, // [3:6] is the sub-list for method output_type
	0, // [0:3] is the sub-list for method input_type
	0, // [0:0] is the sub-list for extension type_name
	0, // [0:0] is the sub-list for extension extendee
	0, // [0:0] is the sub-list for field type_name
}

func init() { file_pb_notification_proto_init() }
func file_pb_notification_proto_init() {
	if File_pb_notification_proto != nil {
		return
	}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_pb_notification_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   3,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_pb_notification_proto_goTypes,
		DependencyIndexes: file_pb_notification_proto_depIdxs,
		MessageInfos:      file_pb_notification_proto_msgTypes,
	}.Build()
	File_pb_notification_proto = out.File
	file_pb_notification_proto_rawDesc = nil
	file_pb_notification_proto_goTypes = nil
	file_pb_notification_proto_depIdxs = nil
}
