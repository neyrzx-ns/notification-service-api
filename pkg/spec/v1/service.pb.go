// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.36.6
// 	protoc        v6.31.1
// source: spec/v1/service.proto

package v1

import (
	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
	reflect "reflect"
	sync "sync"
	unsafe "unsafe"
)

const (
	// Verify that this generated code is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(20 - protoimpl.MinVersion)
	// Verify that runtime/protoimpl is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(protoimpl.MaxVersion - 20)
)

type NotificationTemplate struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	Id            string                 `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
	Slug          string                 `protobuf:"bytes,2,opt,name=slug,proto3" json:"slug,omitempty"`
	Name          string                 `protobuf:"bytes,3,opt,name=name,proto3" json:"name,omitempty"`
	Content       string                 `protobuf:"bytes,4,opt,name=content,proto3" json:"content,omitempty"`
	Params        map[string]string      `protobuf:"bytes,5,rep,name=params,proto3" json:"params,omitempty" protobuf_key:"bytes,1,opt,name=key" protobuf_val:"bytes,2,opt,name=value"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *NotificationTemplate) Reset() {
	*x = NotificationTemplate{}
	mi := &file_spec_v1_service_proto_msgTypes[0]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *NotificationTemplate) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*NotificationTemplate) ProtoMessage() {}

func (x *NotificationTemplate) ProtoReflect() protoreflect.Message {
	mi := &file_spec_v1_service_proto_msgTypes[0]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use NotificationTemplate.ProtoReflect.Descriptor instead.
func (*NotificationTemplate) Descriptor() ([]byte, []int) {
	return file_spec_v1_service_proto_rawDescGZIP(), []int{0}
}

func (x *NotificationTemplate) GetId() string {
	if x != nil {
		return x.Id
	}
	return ""
}

func (x *NotificationTemplate) GetSlug() string {
	if x != nil {
		return x.Slug
	}
	return ""
}

func (x *NotificationTemplate) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

func (x *NotificationTemplate) GetContent() string {
	if x != nil {
		return x.Content
	}
	return ""
}

func (x *NotificationTemplate) GetParams() map[string]string {
	if x != nil {
		return x.Params
	}
	return nil
}

type CreateNotificationTemplateRequest struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	Name          string                 `protobuf:"bytes,1,opt,name=name,proto3" json:"name,omitempty"`
	Content       string                 `protobuf:"bytes,2,opt,name=content,proto3" json:"content,omitempty"`
	Params        map[string]string      `protobuf:"bytes,3,rep,name=params,proto3" json:"params,omitempty" protobuf_key:"bytes,1,opt,name=key" protobuf_val:"bytes,2,opt,name=value"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *CreateNotificationTemplateRequest) Reset() {
	*x = CreateNotificationTemplateRequest{}
	mi := &file_spec_v1_service_proto_msgTypes[1]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *CreateNotificationTemplateRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CreateNotificationTemplateRequest) ProtoMessage() {}

func (x *CreateNotificationTemplateRequest) ProtoReflect() protoreflect.Message {
	mi := &file_spec_v1_service_proto_msgTypes[1]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CreateNotificationTemplateRequest.ProtoReflect.Descriptor instead.
func (*CreateNotificationTemplateRequest) Descriptor() ([]byte, []int) {
	return file_spec_v1_service_proto_rawDescGZIP(), []int{1}
}

func (x *CreateNotificationTemplateRequest) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

func (x *CreateNotificationTemplateRequest) GetContent() string {
	if x != nil {
		return x.Content
	}
	return ""
}

func (x *CreateNotificationTemplateRequest) GetParams() map[string]string {
	if x != nil {
		return x.Params
	}
	return nil
}

type CreateNotificationTemplateReplay struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	Template      *NotificationTemplate  `protobuf:"bytes,1,opt,name=template,proto3" json:"template,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *CreateNotificationTemplateReplay) Reset() {
	*x = CreateNotificationTemplateReplay{}
	mi := &file_spec_v1_service_proto_msgTypes[2]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *CreateNotificationTemplateReplay) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CreateNotificationTemplateReplay) ProtoMessage() {}

func (x *CreateNotificationTemplateReplay) ProtoReflect() protoreflect.Message {
	mi := &file_spec_v1_service_proto_msgTypes[2]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CreateNotificationTemplateReplay.ProtoReflect.Descriptor instead.
func (*CreateNotificationTemplateReplay) Descriptor() ([]byte, []int) {
	return file_spec_v1_service_proto_rawDescGZIP(), []int{2}
}

func (x *CreateNotificationTemplateReplay) GetTemplate() *NotificationTemplate {
	if x != nil {
		return x.Template
	}
	return nil
}

var File_spec_v1_service_proto protoreflect.FileDescriptor

const file_spec_v1_service_proto_rawDesc = "" +
	"\n" +
	"\x15spec/v1/service.proto\x12\x1ens.notification.service.api.v1\"\xfd\x01\n" +
	"\x14NotificationTemplate\x12\x0e\n" +
	"\x02id\x18\x01 \x01(\tR\x02id\x12\x12\n" +
	"\x04slug\x18\x02 \x01(\tR\x04slug\x12\x12\n" +
	"\x04name\x18\x03 \x01(\tR\x04name\x12\x18\n" +
	"\acontent\x18\x04 \x01(\tR\acontent\x12X\n" +
	"\x06params\x18\x05 \x03(\v2@.ns.notification.service.api.v1.NotificationTemplate.ParamsEntryR\x06params\x1a9\n" +
	"\vParamsEntry\x12\x10\n" +
	"\x03key\x18\x01 \x01(\tR\x03key\x12\x14\n" +
	"\x05value\x18\x02 \x01(\tR\x05value:\x028\x01\"\xf3\x01\n" +
	"!CreateNotificationTemplateRequest\x12\x12\n" +
	"\x04name\x18\x01 \x01(\tR\x04name\x12\x18\n" +
	"\acontent\x18\x02 \x01(\tR\acontent\x12e\n" +
	"\x06params\x18\x03 \x03(\v2M.ns.notification.service.api.v1.CreateNotificationTemplateRequest.ParamsEntryR\x06params\x1a9\n" +
	"\vParamsEntry\x12\x10\n" +
	"\x03key\x18\x01 \x01(\tR\x03key\x12\x14\n" +
	"\x05value\x18\x02 \x01(\tR\x05value:\x028\x01\"t\n" +
	" CreateNotificationTemplateReplay\x12P\n" +
	"\btemplate\x18\x01 \x01(\v24.ns.notification.service.api.v1.NotificationTemplateR\btemplate2\xbb\x01\n" +
	"\x13NotificationService\x12\xa3\x01\n" +
	"\x1aCreateNotificationTemplate\x12A.ns.notification.service.api.v1.CreateNotificationTemplateRequest\x1a@.ns.notification.service.api.v1.CreateNotificationTemplateReplay\"\x00B2Z0github.com/neryzx-ns/notification-service-api/v1b\x06proto3"

var (
	file_spec_v1_service_proto_rawDescOnce sync.Once
	file_spec_v1_service_proto_rawDescData []byte
)

func file_spec_v1_service_proto_rawDescGZIP() []byte {
	file_spec_v1_service_proto_rawDescOnce.Do(func() {
		file_spec_v1_service_proto_rawDescData = protoimpl.X.CompressGZIP(unsafe.Slice(unsafe.StringData(file_spec_v1_service_proto_rawDesc), len(file_spec_v1_service_proto_rawDesc)))
	})
	return file_spec_v1_service_proto_rawDescData
}

var file_spec_v1_service_proto_msgTypes = make([]protoimpl.MessageInfo, 5)
var file_spec_v1_service_proto_goTypes = []any{
	(*NotificationTemplate)(nil),              // 0: ns.notification.service.api.v1.NotificationTemplate
	(*CreateNotificationTemplateRequest)(nil), // 1: ns.notification.service.api.v1.CreateNotificationTemplateRequest
	(*CreateNotificationTemplateReplay)(nil),  // 2: ns.notification.service.api.v1.CreateNotificationTemplateReplay
	nil,                                       // 3: ns.notification.service.api.v1.NotificationTemplate.ParamsEntry
	nil,                                       // 4: ns.notification.service.api.v1.CreateNotificationTemplateRequest.ParamsEntry
}
var file_spec_v1_service_proto_depIdxs = []int32{
	3, // 0: ns.notification.service.api.v1.NotificationTemplate.params:type_name -> ns.notification.service.api.v1.NotificationTemplate.ParamsEntry
	4, // 1: ns.notification.service.api.v1.CreateNotificationTemplateRequest.params:type_name -> ns.notification.service.api.v1.CreateNotificationTemplateRequest.ParamsEntry
	0, // 2: ns.notification.service.api.v1.CreateNotificationTemplateReplay.template:type_name -> ns.notification.service.api.v1.NotificationTemplate
	1, // 3: ns.notification.service.api.v1.NotificationService.CreateNotificationTemplate:input_type -> ns.notification.service.api.v1.CreateNotificationTemplateRequest
	2, // 4: ns.notification.service.api.v1.NotificationService.CreateNotificationTemplate:output_type -> ns.notification.service.api.v1.CreateNotificationTemplateReplay
	4, // [4:5] is the sub-list for method output_type
	3, // [3:4] is the sub-list for method input_type
	3, // [3:3] is the sub-list for extension type_name
	3, // [3:3] is the sub-list for extension extendee
	0, // [0:3] is the sub-list for field type_name
}

func init() { file_spec_v1_service_proto_init() }
func file_spec_v1_service_proto_init() {
	if File_spec_v1_service_proto != nil {
		return
	}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: unsafe.Slice(unsafe.StringData(file_spec_v1_service_proto_rawDesc), len(file_spec_v1_service_proto_rawDesc)),
			NumEnums:      0,
			NumMessages:   5,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_spec_v1_service_proto_goTypes,
		DependencyIndexes: file_spec_v1_service_proto_depIdxs,
		MessageInfos:      file_spec_v1_service_proto_msgTypes,
	}.Build()
	File_spec_v1_service_proto = out.File
	file_spec_v1_service_proto_goTypes = nil
	file_spec_v1_service_proto_depIdxs = nil
}
