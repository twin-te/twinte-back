// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.32.0
// 	protoc        (unknown)
// source: nativeapp/v1/service.proto

package nativeappv1

import (
	v1 "github.com/twin-te/twinte-back/handler/api/rpcgen/schoolcalendar/v1"
	sharedpb "github.com/twin-te/twinte-back/handler/api/rpcgen/sharedpb"
	v11 "github.com/twin-te/twinte-back/handler/api/rpcgen/timetable/v1"
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

type GetForWidgetRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Date *sharedpb.RFC3339FullDate `protobuf:"bytes,1,opt,name=date,proto3" json:"date,omitempty"`
}

func (x *GetForWidgetRequest) Reset() {
	*x = GetForWidgetRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_nativeapp_v1_service_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetForWidgetRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetForWidgetRequest) ProtoMessage() {}

func (x *GetForWidgetRequest) ProtoReflect() protoreflect.Message {
	mi := &file_nativeapp_v1_service_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetForWidgetRequest.ProtoReflect.Descriptor instead.
func (*GetForWidgetRequest) Descriptor() ([]byte, []int) {
	return file_nativeapp_v1_service_proto_rawDescGZIP(), []int{0}
}

func (x *GetForWidgetRequest) GetDate() *sharedpb.RFC3339FullDate {
	if x != nil {
		return x.Date
	}
	return nil
}

// If the user is not authenticated, the length of registered_courses is 0.
type GetForWidgetResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Events            []*v1.Event             `protobuf:"bytes,1,rep,name=events,proto3" json:"events,omitempty"`
	CurrentModule     v1.Module               `protobuf:"varint,2,opt,name=current_module,json=currentModule,proto3,enum=schoolcalendar.v1.Module" json:"current_module,omitempty"`
	RegisteredCourses []*v11.RegisteredCourse `protobuf:"bytes,3,rep,name=registered_courses,json=registeredCourses,proto3" json:"registered_courses,omitempty"` // in this academic year
}

func (x *GetForWidgetResponse) Reset() {
	*x = GetForWidgetResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_nativeapp_v1_service_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetForWidgetResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetForWidgetResponse) ProtoMessage() {}

func (x *GetForWidgetResponse) ProtoReflect() protoreflect.Message {
	mi := &file_nativeapp_v1_service_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetForWidgetResponse.ProtoReflect.Descriptor instead.
func (*GetForWidgetResponse) Descriptor() ([]byte, []int) {
	return file_nativeapp_v1_service_proto_rawDescGZIP(), []int{1}
}

func (x *GetForWidgetResponse) GetEvents() []*v1.Event {
	if x != nil {
		return x.Events
	}
	return nil
}

func (x *GetForWidgetResponse) GetCurrentModule() v1.Module {
	if x != nil {
		return x.CurrentModule
	}
	return v1.Module(0)
}

func (x *GetForWidgetResponse) GetRegisteredCourses() []*v11.RegisteredCourse {
	if x != nil {
		return x.RegisteredCourses
	}
	return nil
}

var File_nativeapp_v1_service_proto protoreflect.FileDescriptor

var file_nativeapp_v1_service_proto_rawDesc = []byte{
	0x0a, 0x1a, 0x6e, 0x61, 0x74, 0x69, 0x76, 0x65, 0x61, 0x70, 0x70, 0x2f, 0x76, 0x31, 0x2f, 0x73,
	0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x0c, 0x6e, 0x61,
	0x74, 0x69, 0x76, 0x65, 0x61, 0x70, 0x70, 0x2e, 0x76, 0x31, 0x1a, 0x1c, 0x73, 0x63, 0x68, 0x6f,
	0x6f, 0x6c, 0x63, 0x61, 0x6c, 0x65, 0x6e, 0x64, 0x61, 0x72, 0x2f, 0x76, 0x31, 0x2f, 0x74, 0x79,
	0x70, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x11, 0x73, 0x68, 0x61, 0x72, 0x65, 0x64,
	0x2f, 0x74, 0x79, 0x70, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x17, 0x74, 0x69, 0x6d,
	0x65, 0x74, 0x61, 0x62, 0x6c, 0x65, 0x2f, 0x76, 0x31, 0x2f, 0x74, 0x79, 0x70, 0x65, 0x2e, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x22, 0x42, 0x0a, 0x13, 0x47, 0x65, 0x74, 0x46, 0x6f, 0x72, 0x57, 0x69,
	0x64, 0x67, 0x65, 0x74, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x2b, 0x0a, 0x04, 0x64,
	0x61, 0x74, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x17, 0x2e, 0x73, 0x68, 0x61, 0x72,
	0x65, 0x64, 0x2e, 0x52, 0x46, 0x43, 0x33, 0x33, 0x33, 0x39, 0x46, 0x75, 0x6c, 0x6c, 0x44, 0x61,
	0x74, 0x65, 0x52, 0x04, 0x64, 0x61, 0x74, 0x65, 0x22, 0xd9, 0x01, 0x0a, 0x14, 0x47, 0x65, 0x74,
	0x46, 0x6f, 0x72, 0x57, 0x69, 0x64, 0x67, 0x65, 0x74, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73,
	0x65, 0x12, 0x30, 0x0a, 0x06, 0x65, 0x76, 0x65, 0x6e, 0x74, 0x73, 0x18, 0x01, 0x20, 0x03, 0x28,
	0x0b, 0x32, 0x18, 0x2e, 0x73, 0x63, 0x68, 0x6f, 0x6f, 0x6c, 0x63, 0x61, 0x6c, 0x65, 0x6e, 0x64,
	0x61, 0x72, 0x2e, 0x76, 0x31, 0x2e, 0x45, 0x76, 0x65, 0x6e, 0x74, 0x52, 0x06, 0x65, 0x76, 0x65,
	0x6e, 0x74, 0x73, 0x12, 0x40, 0x0a, 0x0e, 0x63, 0x75, 0x72, 0x72, 0x65, 0x6e, 0x74, 0x5f, 0x6d,
	0x6f, 0x64, 0x75, 0x6c, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0e, 0x32, 0x19, 0x2e, 0x73, 0x63,
	0x68, 0x6f, 0x6f, 0x6c, 0x63, 0x61, 0x6c, 0x65, 0x6e, 0x64, 0x61, 0x72, 0x2e, 0x76, 0x31, 0x2e,
	0x4d, 0x6f, 0x64, 0x75, 0x6c, 0x65, 0x52, 0x0d, 0x63, 0x75, 0x72, 0x72, 0x65, 0x6e, 0x74, 0x4d,
	0x6f, 0x64, 0x75, 0x6c, 0x65, 0x12, 0x4d, 0x0a, 0x12, 0x72, 0x65, 0x67, 0x69, 0x73, 0x74, 0x65,
	0x72, 0x65, 0x64, 0x5f, 0x63, 0x6f, 0x75, 0x72, 0x73, 0x65, 0x73, 0x18, 0x03, 0x20, 0x03, 0x28,
	0x0b, 0x32, 0x1e, 0x2e, 0x74, 0x69, 0x6d, 0x65, 0x74, 0x61, 0x62, 0x6c, 0x65, 0x2e, 0x76, 0x31,
	0x2e, 0x52, 0x65, 0x67, 0x69, 0x73, 0x74, 0x65, 0x72, 0x65, 0x64, 0x43, 0x6f, 0x75, 0x72, 0x73,
	0x65, 0x52, 0x11, 0x72, 0x65, 0x67, 0x69, 0x73, 0x74, 0x65, 0x72, 0x65, 0x64, 0x43, 0x6f, 0x75,
	0x72, 0x73, 0x65, 0x73, 0x32, 0x6e, 0x0a, 0x10, 0x4e, 0x61, 0x74, 0x69, 0x76, 0x65, 0x41, 0x70,
	0x70, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x12, 0x5a, 0x0a, 0x0c, 0x47, 0x65, 0x74, 0x46,
	0x6f, 0x72, 0x57, 0x69, 0x64, 0x67, 0x65, 0x74, 0x12, 0x21, 0x2e, 0x6e, 0x61, 0x74, 0x69, 0x76,
	0x65, 0x61, 0x70, 0x70, 0x2e, 0x76, 0x31, 0x2e, 0x47, 0x65, 0x74, 0x46, 0x6f, 0x72, 0x57, 0x69,
	0x64, 0x67, 0x65, 0x74, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x22, 0x2e, 0x6e, 0x61,
	0x74, 0x69, 0x76, 0x65, 0x61, 0x70, 0x70, 0x2e, 0x76, 0x31, 0x2e, 0x47, 0x65, 0x74, 0x46, 0x6f,
	0x72, 0x57, 0x69, 0x64, 0x67, 0x65, 0x74, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22,
	0x03, 0x90, 0x02, 0x01, 0x42, 0x4c, 0x5a, 0x4a, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63,
	0x6f, 0x6d, 0x2f, 0x74, 0x77, 0x69, 0x6e, 0x2d, 0x74, 0x65, 0x2f, 0x74, 0x77, 0x69, 0x6e, 0x74,
	0x65, 0x2d, 0x62, 0x61, 0x63, 0x6b, 0x2f, 0x68, 0x61, 0x6e, 0x64, 0x6c, 0x65, 0x72, 0x2f, 0x61,
	0x70, 0x69, 0x2f, 0x72, 0x70, 0x63, 0x67, 0x65, 0x6e, 0x2f, 0x6e, 0x61, 0x74, 0x69, 0x76, 0x65,
	0x61, 0x70, 0x70, 0x2f, 0x76, 0x31, 0x3b, 0x6e, 0x61, 0x74, 0x69, 0x76, 0x65, 0x61, 0x70, 0x70,
	0x76, 0x31, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_nativeapp_v1_service_proto_rawDescOnce sync.Once
	file_nativeapp_v1_service_proto_rawDescData = file_nativeapp_v1_service_proto_rawDesc
)

func file_nativeapp_v1_service_proto_rawDescGZIP() []byte {
	file_nativeapp_v1_service_proto_rawDescOnce.Do(func() {
		file_nativeapp_v1_service_proto_rawDescData = protoimpl.X.CompressGZIP(file_nativeapp_v1_service_proto_rawDescData)
	})
	return file_nativeapp_v1_service_proto_rawDescData
}

var file_nativeapp_v1_service_proto_msgTypes = make([]protoimpl.MessageInfo, 2)
var file_nativeapp_v1_service_proto_goTypes = []interface{}{
	(*GetForWidgetRequest)(nil),      // 0: nativeapp.v1.GetForWidgetRequest
	(*GetForWidgetResponse)(nil),     // 1: nativeapp.v1.GetForWidgetResponse
	(*sharedpb.RFC3339FullDate)(nil), // 2: shared.RFC3339FullDate
	(*v1.Event)(nil),                 // 3: schoolcalendar.v1.Event
	(v1.Module)(0),                   // 4: schoolcalendar.v1.Module
	(*v11.RegisteredCourse)(nil),     // 5: timetable.v1.RegisteredCourse
}
var file_nativeapp_v1_service_proto_depIdxs = []int32{
	2, // 0: nativeapp.v1.GetForWidgetRequest.date:type_name -> shared.RFC3339FullDate
	3, // 1: nativeapp.v1.GetForWidgetResponse.events:type_name -> schoolcalendar.v1.Event
	4, // 2: nativeapp.v1.GetForWidgetResponse.current_module:type_name -> schoolcalendar.v1.Module
	5, // 3: nativeapp.v1.GetForWidgetResponse.registered_courses:type_name -> timetable.v1.RegisteredCourse
	0, // 4: nativeapp.v1.NativeAppService.GetForWidget:input_type -> nativeapp.v1.GetForWidgetRequest
	1, // 5: nativeapp.v1.NativeAppService.GetForWidget:output_type -> nativeapp.v1.GetForWidgetResponse
	5, // [5:6] is the sub-list for method output_type
	4, // [4:5] is the sub-list for method input_type
	4, // [4:4] is the sub-list for extension type_name
	4, // [4:4] is the sub-list for extension extendee
	0, // [0:4] is the sub-list for field type_name
}

func init() { file_nativeapp_v1_service_proto_init() }
func file_nativeapp_v1_service_proto_init() {
	if File_nativeapp_v1_service_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_nativeapp_v1_service_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetForWidgetRequest); i {
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
		file_nativeapp_v1_service_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetForWidgetResponse); i {
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
			RawDescriptor: file_nativeapp_v1_service_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   2,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_nativeapp_v1_service_proto_goTypes,
		DependencyIndexes: file_nativeapp_v1_service_proto_depIdxs,
		MessageInfos:      file_nativeapp_v1_service_proto_msgTypes,
	}.Build()
	File_nativeapp_v1_service_proto = out.File
	file_nativeapp_v1_service_proto_rawDesc = nil
	file_nativeapp_v1_service_proto_goTypes = nil
	file_nativeapp_v1_service_proto_depIdxs = nil
}
