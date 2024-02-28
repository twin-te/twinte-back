// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.32.0
// 	protoc        (unknown)
// source: schoolcalendar/v1/type.proto

package schoolcalendarv1

import (
	sharedpb "github.com/twin-te/twinte-back/handler/api/rpcgen/sharedpb"
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

type Weekday int32

const (
	Weekday_WEEKDAY_UNSPECIFIED Weekday = 0
	Weekday_WEEKDAY_SUNDAY      Weekday = 1
	Weekday_WEEKDAY_MONDAY      Weekday = 2
	Weekday_WEEKDAY_TUESDAY     Weekday = 3
	Weekday_WEEKDAY_WEDNESDAY   Weekday = 4
	Weekday_WEEKDAY_THURSDAY    Weekday = 5
	Weekday_WEEKDAY_FRIDAY      Weekday = 6
	Weekday_WEEKDAY_SATURDAY    Weekday = 7
)

// Enum value maps for Weekday.
var (
	Weekday_name = map[int32]string{
		0: "WEEKDAY_UNSPECIFIED",
		1: "WEEKDAY_SUNDAY",
		2: "WEEKDAY_MONDAY",
		3: "WEEKDAY_TUESDAY",
		4: "WEEKDAY_WEDNESDAY",
		5: "WEEKDAY_THURSDAY",
		6: "WEEKDAY_FRIDAY",
		7: "WEEKDAY_SATURDAY",
	}
	Weekday_value = map[string]int32{
		"WEEKDAY_UNSPECIFIED": 0,
		"WEEKDAY_SUNDAY":      1,
		"WEEKDAY_MONDAY":      2,
		"WEEKDAY_TUESDAY":     3,
		"WEEKDAY_WEDNESDAY":   4,
		"WEEKDAY_THURSDAY":    5,
		"WEEKDAY_FRIDAY":      6,
		"WEEKDAY_SATURDAY":    7,
	}
)

func (x Weekday) Enum() *Weekday {
	p := new(Weekday)
	*p = x
	return p
}

func (x Weekday) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (Weekday) Descriptor() protoreflect.EnumDescriptor {
	return file_schoolcalendar_v1_type_proto_enumTypes[0].Descriptor()
}

func (Weekday) Type() protoreflect.EnumType {
	return &file_schoolcalendar_v1_type_proto_enumTypes[0]
}

func (x Weekday) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use Weekday.Descriptor instead.
func (Weekday) EnumDescriptor() ([]byte, []int) {
	return file_schoolcalendar_v1_type_proto_rawDescGZIP(), []int{0}
}

type EventType int32

const (
	EventType_EVENT_TYPE_UNSPECIFIED    EventType = 0
	EventType_EVENT_TYPE_HOLIDAY        EventType = 1
	EventType_EVENT_TYPE_PUBLIC_HOLIDAY EventType = 2
	EventType_EVENT_TYPE_EXAM           EventType = 3
	EventType_EVENT_TYPE_SUBSTITUTE_DAY EventType = 4
	EventType_EVENT_TYPE_OTHER          EventType = 5
)

// Enum value maps for EventType.
var (
	EventType_name = map[int32]string{
		0: "EVENT_TYPE_UNSPECIFIED",
		1: "EVENT_TYPE_HOLIDAY",
		2: "EVENT_TYPE_PUBLIC_HOLIDAY",
		3: "EVENT_TYPE_EXAM",
		4: "EVENT_TYPE_SUBSTITUTE_DAY",
		5: "EVENT_TYPE_OTHER",
	}
	EventType_value = map[string]int32{
		"EVENT_TYPE_UNSPECIFIED":    0,
		"EVENT_TYPE_HOLIDAY":        1,
		"EVENT_TYPE_PUBLIC_HOLIDAY": 2,
		"EVENT_TYPE_EXAM":           3,
		"EVENT_TYPE_SUBSTITUTE_DAY": 4,
		"EVENT_TYPE_OTHER":          5,
	}
)

func (x EventType) Enum() *EventType {
	p := new(EventType)
	*p = x
	return p
}

func (x EventType) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (EventType) Descriptor() protoreflect.EnumDescriptor {
	return file_schoolcalendar_v1_type_proto_enumTypes[1].Descriptor()
}

func (EventType) Type() protoreflect.EnumType {
	return &file_schoolcalendar_v1_type_proto_enumTypes[1]
}

func (x EventType) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use EventType.Descriptor instead.
func (EventType) EnumDescriptor() ([]byte, []int) {
	return file_schoolcalendar_v1_type_proto_rawDescGZIP(), []int{1}
}

type Module int32

const (
	Module_MODULE_UNSPECIFIED     Module = 0
	Module_MODULE_SPRING_A        Module = 1
	Module_MODULE_SPRING_B        Module = 2
	Module_MODULE_SPRING_C        Module = 3
	Module_MODULE_SUMMER_VACATION Module = 4
	Module_MODULE_FALL_A          Module = 5
	Module_MODULE_FALL_B          Module = 6
	Module_MODULE_WINTER_VACATION Module = 7
	Module_MODULE_FALL_C          Module = 8
	Module_MODULE_SPRING_VACATION Module = 9
)

// Enum value maps for Module.
var (
	Module_name = map[int32]string{
		0: "MODULE_UNSPECIFIED",
		1: "MODULE_SPRING_A",
		2: "MODULE_SPRING_B",
		3: "MODULE_SPRING_C",
		4: "MODULE_SUMMER_VACATION",
		5: "MODULE_FALL_A",
		6: "MODULE_FALL_B",
		7: "MODULE_WINTER_VACATION",
		8: "MODULE_FALL_C",
		9: "MODULE_SPRING_VACATION",
	}
	Module_value = map[string]int32{
		"MODULE_UNSPECIFIED":     0,
		"MODULE_SPRING_A":        1,
		"MODULE_SPRING_B":        2,
		"MODULE_SPRING_C":        3,
		"MODULE_SUMMER_VACATION": 4,
		"MODULE_FALL_A":          5,
		"MODULE_FALL_B":          6,
		"MODULE_WINTER_VACATION": 7,
		"MODULE_FALL_C":          8,
		"MODULE_SPRING_VACATION": 9,
	}
)

func (x Module) Enum() *Module {
	p := new(Module)
	*p = x
	return p
}

func (x Module) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (Module) Descriptor() protoreflect.EnumDescriptor {
	return file_schoolcalendar_v1_type_proto_enumTypes[2].Descriptor()
}

func (Module) Type() protoreflect.EnumType {
	return &file_schoolcalendar_v1_type_proto_enumTypes[2]
}

func (x Module) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use Module.Descriptor instead.
func (Module) EnumDescriptor() ([]byte, []int) {
	return file_schoolcalendar_v1_type_proto_rawDescGZIP(), []int{2}
}

type Event struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id          int32                     `protobuf:"varint,1,opt,name=id,proto3" json:"id,omitempty"`
	Type        EventType                 `protobuf:"varint,2,opt,name=type,proto3,enum=schoolcalendar.v1.EventType" json:"type,omitempty"`
	Date        *sharedpb.RFC3339FullDate `protobuf:"bytes,3,opt,name=date,proto3" json:"date,omitempty"`
	Description string                    `protobuf:"bytes,4,opt,name=description,proto3" json:"description,omitempty"`
	ChangeTo    *Weekday                  `protobuf:"varint,5,opt,name=change_to,json=changeTo,proto3,enum=schoolcalendar.v1.Weekday,oneof" json:"change_to,omitempty"`
}

func (x *Event) Reset() {
	*x = Event{}
	if protoimpl.UnsafeEnabled {
		mi := &file_schoolcalendar_v1_type_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Event) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Event) ProtoMessage() {}

func (x *Event) ProtoReflect() protoreflect.Message {
	mi := &file_schoolcalendar_v1_type_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Event.ProtoReflect.Descriptor instead.
func (*Event) Descriptor() ([]byte, []int) {
	return file_schoolcalendar_v1_type_proto_rawDescGZIP(), []int{0}
}

func (x *Event) GetId() int32 {
	if x != nil {
		return x.Id
	}
	return 0
}

func (x *Event) GetType() EventType {
	if x != nil {
		return x.Type
	}
	return EventType_EVENT_TYPE_UNSPECIFIED
}

func (x *Event) GetDate() *sharedpb.RFC3339FullDate {
	if x != nil {
		return x.Date
	}
	return nil
}

func (x *Event) GetDescription() string {
	if x != nil {
		return x.Description
	}
	return ""
}

func (x *Event) GetChangeTo() Weekday {
	if x != nil && x.ChangeTo != nil {
		return *x.ChangeTo
	}
	return Weekday_WEEKDAY_UNSPECIFIED
}

type ModuleDetail struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id     int32                     `protobuf:"varint,1,opt,name=id,proto3" json:"id,omitempty"`
	Year   *sharedpb.AcademicYear    `protobuf:"bytes,2,opt,name=year,proto3" json:"year,omitempty"`
	Module Module                    `protobuf:"varint,3,opt,name=module,proto3,enum=schoolcalendar.v1.Module" json:"module,omitempty"`
	Start  *sharedpb.RFC3339FullDate `protobuf:"bytes,4,opt,name=start,proto3" json:"start,omitempty"`
	End    *sharedpb.RFC3339FullDate `protobuf:"bytes,5,opt,name=end,proto3" json:"end,omitempty"`
}

func (x *ModuleDetail) Reset() {
	*x = ModuleDetail{}
	if protoimpl.UnsafeEnabled {
		mi := &file_schoolcalendar_v1_type_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ModuleDetail) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ModuleDetail) ProtoMessage() {}

func (x *ModuleDetail) ProtoReflect() protoreflect.Message {
	mi := &file_schoolcalendar_v1_type_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ModuleDetail.ProtoReflect.Descriptor instead.
func (*ModuleDetail) Descriptor() ([]byte, []int) {
	return file_schoolcalendar_v1_type_proto_rawDescGZIP(), []int{1}
}

func (x *ModuleDetail) GetId() int32 {
	if x != nil {
		return x.Id
	}
	return 0
}

func (x *ModuleDetail) GetYear() *sharedpb.AcademicYear {
	if x != nil {
		return x.Year
	}
	return nil
}

func (x *ModuleDetail) GetModule() Module {
	if x != nil {
		return x.Module
	}
	return Module_MODULE_UNSPECIFIED
}

func (x *ModuleDetail) GetStart() *sharedpb.RFC3339FullDate {
	if x != nil {
		return x.Start
	}
	return nil
}

func (x *ModuleDetail) GetEnd() *sharedpb.RFC3339FullDate {
	if x != nil {
		return x.End
	}
	return nil
}

var File_schoolcalendar_v1_type_proto protoreflect.FileDescriptor

var file_schoolcalendar_v1_type_proto_rawDesc = []byte{
	0x0a, 0x1c, 0x73, 0x63, 0x68, 0x6f, 0x6f, 0x6c, 0x63, 0x61, 0x6c, 0x65, 0x6e, 0x64, 0x61, 0x72,
	0x2f, 0x76, 0x31, 0x2f, 0x74, 0x79, 0x70, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x11,
	0x73, 0x63, 0x68, 0x6f, 0x6f, 0x6c, 0x63, 0x61, 0x6c, 0x65, 0x6e, 0x64, 0x61, 0x72, 0x2e, 0x76,
	0x31, 0x1a, 0x11, 0x73, 0x68, 0x61, 0x72, 0x65, 0x64, 0x2f, 0x74, 0x79, 0x70, 0x65, 0x2e, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x22, 0xe4, 0x01, 0x0a, 0x05, 0x45, 0x76, 0x65, 0x6e, 0x74, 0x12, 0x0e,
	0x0a, 0x02, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x05, 0x52, 0x02, 0x69, 0x64, 0x12, 0x30,
	0x0a, 0x04, 0x74, 0x79, 0x70, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0e, 0x32, 0x1c, 0x2e, 0x73,
	0x63, 0x68, 0x6f, 0x6f, 0x6c, 0x63, 0x61, 0x6c, 0x65, 0x6e, 0x64, 0x61, 0x72, 0x2e, 0x76, 0x31,
	0x2e, 0x45, 0x76, 0x65, 0x6e, 0x74, 0x54, 0x79, 0x70, 0x65, 0x52, 0x04, 0x74, 0x79, 0x70, 0x65,
	0x12, 0x2b, 0x0a, 0x04, 0x64, 0x61, 0x74, 0x65, 0x18, 0x03, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x17,
	0x2e, 0x73, 0x68, 0x61, 0x72, 0x65, 0x64, 0x2e, 0x52, 0x46, 0x43, 0x33, 0x33, 0x33, 0x39, 0x46,
	0x75, 0x6c, 0x6c, 0x44, 0x61, 0x74, 0x65, 0x52, 0x04, 0x64, 0x61, 0x74, 0x65, 0x12, 0x20, 0x0a,
	0x0b, 0x64, 0x65, 0x73, 0x63, 0x72, 0x69, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x18, 0x04, 0x20, 0x01,
	0x28, 0x09, 0x52, 0x0b, 0x64, 0x65, 0x73, 0x63, 0x72, 0x69, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x12,
	0x3c, 0x0a, 0x09, 0x63, 0x68, 0x61, 0x6e, 0x67, 0x65, 0x5f, 0x74, 0x6f, 0x18, 0x05, 0x20, 0x01,
	0x28, 0x0e, 0x32, 0x1a, 0x2e, 0x73, 0x63, 0x68, 0x6f, 0x6f, 0x6c, 0x63, 0x61, 0x6c, 0x65, 0x6e,
	0x64, 0x61, 0x72, 0x2e, 0x76, 0x31, 0x2e, 0x57, 0x65, 0x65, 0x6b, 0x64, 0x61, 0x79, 0x48, 0x00,
	0x52, 0x08, 0x63, 0x68, 0x61, 0x6e, 0x67, 0x65, 0x54, 0x6f, 0x88, 0x01, 0x01, 0x42, 0x0c, 0x0a,
	0x0a, 0x5f, 0x63, 0x68, 0x61, 0x6e, 0x67, 0x65, 0x5f, 0x74, 0x6f, 0x22, 0xd5, 0x01, 0x0a, 0x0c,
	0x4d, 0x6f, 0x64, 0x75, 0x6c, 0x65, 0x44, 0x65, 0x74, 0x61, 0x69, 0x6c, 0x12, 0x0e, 0x0a, 0x02,
	0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x05, 0x52, 0x02, 0x69, 0x64, 0x12, 0x28, 0x0a, 0x04,
	0x79, 0x65, 0x61, 0x72, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x14, 0x2e, 0x73, 0x68, 0x61,
	0x72, 0x65, 0x64, 0x2e, 0x41, 0x63, 0x61, 0x64, 0x65, 0x6d, 0x69, 0x63, 0x59, 0x65, 0x61, 0x72,
	0x52, 0x04, 0x79, 0x65, 0x61, 0x72, 0x12, 0x31, 0x0a, 0x06, 0x6d, 0x6f, 0x64, 0x75, 0x6c, 0x65,
	0x18, 0x03, 0x20, 0x01, 0x28, 0x0e, 0x32, 0x19, 0x2e, 0x73, 0x63, 0x68, 0x6f, 0x6f, 0x6c, 0x63,
	0x61, 0x6c, 0x65, 0x6e, 0x64, 0x61, 0x72, 0x2e, 0x76, 0x31, 0x2e, 0x4d, 0x6f, 0x64, 0x75, 0x6c,
	0x65, 0x52, 0x06, 0x6d, 0x6f, 0x64, 0x75, 0x6c, 0x65, 0x12, 0x2d, 0x0a, 0x05, 0x73, 0x74, 0x61,
	0x72, 0x74, 0x18, 0x04, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x17, 0x2e, 0x73, 0x68, 0x61, 0x72, 0x65,
	0x64, 0x2e, 0x52, 0x46, 0x43, 0x33, 0x33, 0x33, 0x39, 0x46, 0x75, 0x6c, 0x6c, 0x44, 0x61, 0x74,
	0x65, 0x52, 0x05, 0x73, 0x74, 0x61, 0x72, 0x74, 0x12, 0x29, 0x0a, 0x03, 0x65, 0x6e, 0x64, 0x18,
	0x05, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x17, 0x2e, 0x73, 0x68, 0x61, 0x72, 0x65, 0x64, 0x2e, 0x52,
	0x46, 0x43, 0x33, 0x33, 0x33, 0x39, 0x46, 0x75, 0x6c, 0x6c, 0x44, 0x61, 0x74, 0x65, 0x52, 0x03,
	0x65, 0x6e, 0x64, 0x2a, 0xb6, 0x01, 0x0a, 0x07, 0x57, 0x65, 0x65, 0x6b, 0x64, 0x61, 0x79, 0x12,
	0x17, 0x0a, 0x13, 0x57, 0x45, 0x45, 0x4b, 0x44, 0x41, 0x59, 0x5f, 0x55, 0x4e, 0x53, 0x50, 0x45,
	0x43, 0x49, 0x46, 0x49, 0x45, 0x44, 0x10, 0x00, 0x12, 0x12, 0x0a, 0x0e, 0x57, 0x45, 0x45, 0x4b,
	0x44, 0x41, 0x59, 0x5f, 0x53, 0x55, 0x4e, 0x44, 0x41, 0x59, 0x10, 0x01, 0x12, 0x12, 0x0a, 0x0e,
	0x57, 0x45, 0x45, 0x4b, 0x44, 0x41, 0x59, 0x5f, 0x4d, 0x4f, 0x4e, 0x44, 0x41, 0x59, 0x10, 0x02,
	0x12, 0x13, 0x0a, 0x0f, 0x57, 0x45, 0x45, 0x4b, 0x44, 0x41, 0x59, 0x5f, 0x54, 0x55, 0x45, 0x53,
	0x44, 0x41, 0x59, 0x10, 0x03, 0x12, 0x15, 0x0a, 0x11, 0x57, 0x45, 0x45, 0x4b, 0x44, 0x41, 0x59,
	0x5f, 0x57, 0x45, 0x44, 0x4e, 0x45, 0x53, 0x44, 0x41, 0x59, 0x10, 0x04, 0x12, 0x14, 0x0a, 0x10,
	0x57, 0x45, 0x45, 0x4b, 0x44, 0x41, 0x59, 0x5f, 0x54, 0x48, 0x55, 0x52, 0x53, 0x44, 0x41, 0x59,
	0x10, 0x05, 0x12, 0x12, 0x0a, 0x0e, 0x57, 0x45, 0x45, 0x4b, 0x44, 0x41, 0x59, 0x5f, 0x46, 0x52,
	0x49, 0x44, 0x41, 0x59, 0x10, 0x06, 0x12, 0x14, 0x0a, 0x10, 0x57, 0x45, 0x45, 0x4b, 0x44, 0x41,
	0x59, 0x5f, 0x53, 0x41, 0x54, 0x55, 0x52, 0x44, 0x41, 0x59, 0x10, 0x07, 0x2a, 0xa8, 0x01, 0x0a,
	0x09, 0x45, 0x76, 0x65, 0x6e, 0x74, 0x54, 0x79, 0x70, 0x65, 0x12, 0x1a, 0x0a, 0x16, 0x45, 0x56,
	0x45, 0x4e, 0x54, 0x5f, 0x54, 0x59, 0x50, 0x45, 0x5f, 0x55, 0x4e, 0x53, 0x50, 0x45, 0x43, 0x49,
	0x46, 0x49, 0x45, 0x44, 0x10, 0x00, 0x12, 0x16, 0x0a, 0x12, 0x45, 0x56, 0x45, 0x4e, 0x54, 0x5f,
	0x54, 0x59, 0x50, 0x45, 0x5f, 0x48, 0x4f, 0x4c, 0x49, 0x44, 0x41, 0x59, 0x10, 0x01, 0x12, 0x1d,
	0x0a, 0x19, 0x45, 0x56, 0x45, 0x4e, 0x54, 0x5f, 0x54, 0x59, 0x50, 0x45, 0x5f, 0x50, 0x55, 0x42,
	0x4c, 0x49, 0x43, 0x5f, 0x48, 0x4f, 0x4c, 0x49, 0x44, 0x41, 0x59, 0x10, 0x02, 0x12, 0x13, 0x0a,
	0x0f, 0x45, 0x56, 0x45, 0x4e, 0x54, 0x5f, 0x54, 0x59, 0x50, 0x45, 0x5f, 0x45, 0x58, 0x41, 0x4d,
	0x10, 0x03, 0x12, 0x1d, 0x0a, 0x19, 0x45, 0x56, 0x45, 0x4e, 0x54, 0x5f, 0x54, 0x59, 0x50, 0x45,
	0x5f, 0x53, 0x55, 0x42, 0x53, 0x54, 0x49, 0x54, 0x55, 0x54, 0x45, 0x5f, 0x44, 0x41, 0x59, 0x10,
	0x04, 0x12, 0x14, 0x0a, 0x10, 0x45, 0x56, 0x45, 0x4e, 0x54, 0x5f, 0x54, 0x59, 0x50, 0x45, 0x5f,
	0x4f, 0x54, 0x48, 0x45, 0x52, 0x10, 0x05, 0x2a, 0xec, 0x01, 0x0a, 0x06, 0x4d, 0x6f, 0x64, 0x75,
	0x6c, 0x65, 0x12, 0x16, 0x0a, 0x12, 0x4d, 0x4f, 0x44, 0x55, 0x4c, 0x45, 0x5f, 0x55, 0x4e, 0x53,
	0x50, 0x45, 0x43, 0x49, 0x46, 0x49, 0x45, 0x44, 0x10, 0x00, 0x12, 0x13, 0x0a, 0x0f, 0x4d, 0x4f,
	0x44, 0x55, 0x4c, 0x45, 0x5f, 0x53, 0x50, 0x52, 0x49, 0x4e, 0x47, 0x5f, 0x41, 0x10, 0x01, 0x12,
	0x13, 0x0a, 0x0f, 0x4d, 0x4f, 0x44, 0x55, 0x4c, 0x45, 0x5f, 0x53, 0x50, 0x52, 0x49, 0x4e, 0x47,
	0x5f, 0x42, 0x10, 0x02, 0x12, 0x13, 0x0a, 0x0f, 0x4d, 0x4f, 0x44, 0x55, 0x4c, 0x45, 0x5f, 0x53,
	0x50, 0x52, 0x49, 0x4e, 0x47, 0x5f, 0x43, 0x10, 0x03, 0x12, 0x1a, 0x0a, 0x16, 0x4d, 0x4f, 0x44,
	0x55, 0x4c, 0x45, 0x5f, 0x53, 0x55, 0x4d, 0x4d, 0x45, 0x52, 0x5f, 0x56, 0x41, 0x43, 0x41, 0x54,
	0x49, 0x4f, 0x4e, 0x10, 0x04, 0x12, 0x11, 0x0a, 0x0d, 0x4d, 0x4f, 0x44, 0x55, 0x4c, 0x45, 0x5f,
	0x46, 0x41, 0x4c, 0x4c, 0x5f, 0x41, 0x10, 0x05, 0x12, 0x11, 0x0a, 0x0d, 0x4d, 0x4f, 0x44, 0x55,
	0x4c, 0x45, 0x5f, 0x46, 0x41, 0x4c, 0x4c, 0x5f, 0x42, 0x10, 0x06, 0x12, 0x1a, 0x0a, 0x16, 0x4d,
	0x4f, 0x44, 0x55, 0x4c, 0x45, 0x5f, 0x57, 0x49, 0x4e, 0x54, 0x45, 0x52, 0x5f, 0x56, 0x41, 0x43,
	0x41, 0x54, 0x49, 0x4f, 0x4e, 0x10, 0x07, 0x12, 0x11, 0x0a, 0x0d, 0x4d, 0x4f, 0x44, 0x55, 0x4c,
	0x45, 0x5f, 0x46, 0x41, 0x4c, 0x4c, 0x5f, 0x43, 0x10, 0x08, 0x12, 0x1a, 0x0a, 0x16, 0x4d, 0x4f,
	0x44, 0x55, 0x4c, 0x45, 0x5f, 0x53, 0x50, 0x52, 0x49, 0x4e, 0x47, 0x5f, 0x56, 0x41, 0x43, 0x41,
	0x54, 0x49, 0x4f, 0x4e, 0x10, 0x09, 0x42, 0x56, 0x5a, 0x54, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62,
	0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x74, 0x77, 0x69, 0x6e, 0x2d, 0x74, 0x65, 0x2f, 0x74, 0x77, 0x69,
	0x6e, 0x74, 0x65, 0x2d, 0x62, 0x61, 0x63, 0x6b, 0x2f, 0x68, 0x61, 0x6e, 0x64, 0x6c, 0x65, 0x72,
	0x2f, 0x61, 0x70, 0x69, 0x2f, 0x72, 0x70, 0x63, 0x67, 0x65, 0x6e, 0x2f, 0x73, 0x63, 0x68, 0x6f,
	0x6f, 0x6c, 0x63, 0x61, 0x6c, 0x65, 0x6e, 0x64, 0x61, 0x72, 0x2f, 0x76, 0x31, 0x3b, 0x73, 0x63,
	0x68, 0x6f, 0x6f, 0x6c, 0x63, 0x61, 0x6c, 0x65, 0x6e, 0x64, 0x61, 0x72, 0x76, 0x31, 0x62, 0x06,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_schoolcalendar_v1_type_proto_rawDescOnce sync.Once
	file_schoolcalendar_v1_type_proto_rawDescData = file_schoolcalendar_v1_type_proto_rawDesc
)

func file_schoolcalendar_v1_type_proto_rawDescGZIP() []byte {
	file_schoolcalendar_v1_type_proto_rawDescOnce.Do(func() {
		file_schoolcalendar_v1_type_proto_rawDescData = protoimpl.X.CompressGZIP(file_schoolcalendar_v1_type_proto_rawDescData)
	})
	return file_schoolcalendar_v1_type_proto_rawDescData
}

var file_schoolcalendar_v1_type_proto_enumTypes = make([]protoimpl.EnumInfo, 3)
var file_schoolcalendar_v1_type_proto_msgTypes = make([]protoimpl.MessageInfo, 2)
var file_schoolcalendar_v1_type_proto_goTypes = []interface{}{
	(Weekday)(0),                     // 0: schoolcalendar.v1.Weekday
	(EventType)(0),                   // 1: schoolcalendar.v1.EventType
	(Module)(0),                      // 2: schoolcalendar.v1.Module
	(*Event)(nil),                    // 3: schoolcalendar.v1.Event
	(*ModuleDetail)(nil),             // 4: schoolcalendar.v1.ModuleDetail
	(*sharedpb.RFC3339FullDate)(nil), // 5: shared.RFC3339FullDate
	(*sharedpb.AcademicYear)(nil),    // 6: shared.AcademicYear
}
var file_schoolcalendar_v1_type_proto_depIdxs = []int32{
	1, // 0: schoolcalendar.v1.Event.type:type_name -> schoolcalendar.v1.EventType
	5, // 1: schoolcalendar.v1.Event.date:type_name -> shared.RFC3339FullDate
	0, // 2: schoolcalendar.v1.Event.change_to:type_name -> schoolcalendar.v1.Weekday
	6, // 3: schoolcalendar.v1.ModuleDetail.year:type_name -> shared.AcademicYear
	2, // 4: schoolcalendar.v1.ModuleDetail.module:type_name -> schoolcalendar.v1.Module
	5, // 5: schoolcalendar.v1.ModuleDetail.start:type_name -> shared.RFC3339FullDate
	5, // 6: schoolcalendar.v1.ModuleDetail.end:type_name -> shared.RFC3339FullDate
	7, // [7:7] is the sub-list for method output_type
	7, // [7:7] is the sub-list for method input_type
	7, // [7:7] is the sub-list for extension type_name
	7, // [7:7] is the sub-list for extension extendee
	0, // [0:7] is the sub-list for field type_name
}

func init() { file_schoolcalendar_v1_type_proto_init() }
func file_schoolcalendar_v1_type_proto_init() {
	if File_schoolcalendar_v1_type_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_schoolcalendar_v1_type_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Event); i {
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
		file_schoolcalendar_v1_type_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ModuleDetail); i {
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
	file_schoolcalendar_v1_type_proto_msgTypes[0].OneofWrappers = []interface{}{}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_schoolcalendar_v1_type_proto_rawDesc,
			NumEnums:      3,
			NumMessages:   2,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_schoolcalendar_v1_type_proto_goTypes,
		DependencyIndexes: file_schoolcalendar_v1_type_proto_depIdxs,
		EnumInfos:         file_schoolcalendar_v1_type_proto_enumTypes,
		MessageInfos:      file_schoolcalendar_v1_type_proto_msgTypes,
	}.Build()
	File_schoolcalendar_v1_type_proto = out.File
	file_schoolcalendar_v1_type_proto_rawDesc = nil
	file_schoolcalendar_v1_type_proto_goTypes = nil
	file_schoolcalendar_v1_type_proto_depIdxs = nil
}
