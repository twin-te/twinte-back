// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.32.0
// 	protoc        (unknown)
// source: announcement/v1/type.proto

package announcementv1

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

type AnnouncementTag int32

const (
	AnnouncementTag_ANNOUNCEMENT_TAG_UNSPECIFIED  AnnouncementTag = 0
	AnnouncementTag_ANNOUNCEMENT_TAG_INFORMATION  AnnouncementTag = 1
	AnnouncementTag_ANNOUNCEMENT_TAG_NOTIFICATION AnnouncementTag = 2
)

// Enum value maps for AnnouncementTag.
var (
	AnnouncementTag_name = map[int32]string{
		0: "ANNOUNCEMENT_TAG_UNSPECIFIED",
		1: "ANNOUNCEMENT_TAG_INFORMATION",
		2: "ANNOUNCEMENT_TAG_NOTIFICATION",
	}
	AnnouncementTag_value = map[string]int32{
		"ANNOUNCEMENT_TAG_UNSPECIFIED":  0,
		"ANNOUNCEMENT_TAG_INFORMATION":  1,
		"ANNOUNCEMENT_TAG_NOTIFICATION": 2,
	}
)

func (x AnnouncementTag) Enum() *AnnouncementTag {
	p := new(AnnouncementTag)
	*p = x
	return p
}

func (x AnnouncementTag) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (AnnouncementTag) Descriptor() protoreflect.EnumDescriptor {
	return file_announcement_v1_type_proto_enumTypes[0].Descriptor()
}

func (AnnouncementTag) Type() protoreflect.EnumType {
	return &file_announcement_v1_type_proto_enumTypes[0]
}

func (x AnnouncementTag) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use AnnouncementTag.Descriptor instead.
func (AnnouncementTag) EnumDescriptor() ([]byte, []int) {
	return file_announcement_v1_type_proto_rawDescGZIP(), []int{0}
}

type Announcement struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id          *sharedpb.UUID            `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
	Tags        []AnnouncementTag         `protobuf:"varint,2,rep,packed,name=tags,proto3,enum=announcement.v1.AnnouncementTag" json:"tags,omitempty"`
	Title       string                    `protobuf:"bytes,3,opt,name=title,proto3" json:"title,omitempty"`
	Content     string                    `protobuf:"bytes,4,opt,name=content,proto3" json:"content,omitempty"`
	PublishedAt *sharedpb.RFC3339DateTime `protobuf:"bytes,5,opt,name=published_at,json=publishedAt,proto3" json:"published_at,omitempty"`
}

func (x *Announcement) Reset() {
	*x = Announcement{}
	if protoimpl.UnsafeEnabled {
		mi := &file_announcement_v1_type_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Announcement) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Announcement) ProtoMessage() {}

func (x *Announcement) ProtoReflect() protoreflect.Message {
	mi := &file_announcement_v1_type_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Announcement.ProtoReflect.Descriptor instead.
func (*Announcement) Descriptor() ([]byte, []int) {
	return file_announcement_v1_type_proto_rawDescGZIP(), []int{0}
}

func (x *Announcement) GetId() *sharedpb.UUID {
	if x != nil {
		return x.Id
	}
	return nil
}

func (x *Announcement) GetTags() []AnnouncementTag {
	if x != nil {
		return x.Tags
	}
	return nil
}

func (x *Announcement) GetTitle() string {
	if x != nil {
		return x.Title
	}
	return ""
}

func (x *Announcement) GetContent() string {
	if x != nil {
		return x.Content
	}
	return ""
}

func (x *Announcement) GetPublishedAt() *sharedpb.RFC3339DateTime {
	if x != nil {
		return x.PublishedAt
	}
	return nil
}

var File_announcement_v1_type_proto protoreflect.FileDescriptor

var file_announcement_v1_type_proto_rawDesc = []byte{
	0x0a, 0x1a, 0x61, 0x6e, 0x6e, 0x6f, 0x75, 0x6e, 0x63, 0x65, 0x6d, 0x65, 0x6e, 0x74, 0x2f, 0x76,
	0x31, 0x2f, 0x74, 0x79, 0x70, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x0f, 0x61, 0x6e,
	0x6e, 0x6f, 0x75, 0x6e, 0x63, 0x65, 0x6d, 0x65, 0x6e, 0x74, 0x2e, 0x76, 0x31, 0x1a, 0x11, 0x73,
	0x68, 0x61, 0x72, 0x65, 0x64, 0x2f, 0x74, 0x79, 0x70, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x22, 0xce, 0x01, 0x0a, 0x0c, 0x41, 0x6e, 0x6e, 0x6f, 0x75, 0x6e, 0x63, 0x65, 0x6d, 0x65, 0x6e,
	0x74, 0x12, 0x1c, 0x0a, 0x02, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x0c, 0x2e,
	0x73, 0x68, 0x61, 0x72, 0x65, 0x64, 0x2e, 0x55, 0x55, 0x49, 0x44, 0x52, 0x02, 0x69, 0x64, 0x12,
	0x34, 0x0a, 0x04, 0x74, 0x61, 0x67, 0x73, 0x18, 0x02, 0x20, 0x03, 0x28, 0x0e, 0x32, 0x20, 0x2e,
	0x61, 0x6e, 0x6e, 0x6f, 0x75, 0x6e, 0x63, 0x65, 0x6d, 0x65, 0x6e, 0x74, 0x2e, 0x76, 0x31, 0x2e,
	0x41, 0x6e, 0x6e, 0x6f, 0x75, 0x6e, 0x63, 0x65, 0x6d, 0x65, 0x6e, 0x74, 0x54, 0x61, 0x67, 0x52,
	0x04, 0x74, 0x61, 0x67, 0x73, 0x12, 0x14, 0x0a, 0x05, 0x74, 0x69, 0x74, 0x6c, 0x65, 0x18, 0x03,
	0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x74, 0x69, 0x74, 0x6c, 0x65, 0x12, 0x18, 0x0a, 0x07, 0x63,
	0x6f, 0x6e, 0x74, 0x65, 0x6e, 0x74, 0x18, 0x04, 0x20, 0x01, 0x28, 0x09, 0x52, 0x07, 0x63, 0x6f,
	0x6e, 0x74, 0x65, 0x6e, 0x74, 0x12, 0x3a, 0x0a, 0x0c, 0x70, 0x75, 0x62, 0x6c, 0x69, 0x73, 0x68,
	0x65, 0x64, 0x5f, 0x61, 0x74, 0x18, 0x05, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x17, 0x2e, 0x73, 0x68,
	0x61, 0x72, 0x65, 0x64, 0x2e, 0x52, 0x46, 0x43, 0x33, 0x33, 0x33, 0x39, 0x44, 0x61, 0x74, 0x65,
	0x54, 0x69, 0x6d, 0x65, 0x52, 0x0b, 0x70, 0x75, 0x62, 0x6c, 0x69, 0x73, 0x68, 0x65, 0x64, 0x41,
	0x74, 0x2a, 0x78, 0x0a, 0x0f, 0x41, 0x6e, 0x6e, 0x6f, 0x75, 0x6e, 0x63, 0x65, 0x6d, 0x65, 0x6e,
	0x74, 0x54, 0x61, 0x67, 0x12, 0x20, 0x0a, 0x1c, 0x41, 0x4e, 0x4e, 0x4f, 0x55, 0x4e, 0x43, 0x45,
	0x4d, 0x45, 0x4e, 0x54, 0x5f, 0x54, 0x41, 0x47, 0x5f, 0x55, 0x4e, 0x53, 0x50, 0x45, 0x43, 0x49,
	0x46, 0x49, 0x45, 0x44, 0x10, 0x00, 0x12, 0x20, 0x0a, 0x1c, 0x41, 0x4e, 0x4e, 0x4f, 0x55, 0x4e,
	0x43, 0x45, 0x4d, 0x45, 0x4e, 0x54, 0x5f, 0x54, 0x41, 0x47, 0x5f, 0x49, 0x4e, 0x46, 0x4f, 0x52,
	0x4d, 0x41, 0x54, 0x49, 0x4f, 0x4e, 0x10, 0x01, 0x12, 0x21, 0x0a, 0x1d, 0x41, 0x4e, 0x4e, 0x4f,
	0x55, 0x4e, 0x43, 0x45, 0x4d, 0x45, 0x4e, 0x54, 0x5f, 0x54, 0x41, 0x47, 0x5f, 0x4e, 0x4f, 0x54,
	0x49, 0x46, 0x49, 0x43, 0x41, 0x54, 0x49, 0x4f, 0x4e, 0x10, 0x02, 0x42, 0x52, 0x5a, 0x50, 0x67,
	0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x74, 0x77, 0x69, 0x6e, 0x2d, 0x74,
	0x65, 0x2f, 0x74, 0x77, 0x69, 0x6e, 0x74, 0x65, 0x2d, 0x62, 0x61, 0x63, 0x6b, 0x2f, 0x68, 0x61,
	0x6e, 0x64, 0x6c, 0x65, 0x72, 0x2f, 0x61, 0x70, 0x69, 0x2f, 0x72, 0x70, 0x63, 0x67, 0x65, 0x6e,
	0x2f, 0x61, 0x6e, 0x6e, 0x6f, 0x75, 0x6e, 0x63, 0x65, 0x6d, 0x65, 0x6e, 0x74, 0x2f, 0x76, 0x31,
	0x3b, 0x61, 0x6e, 0x6e, 0x6f, 0x75, 0x6e, 0x63, 0x65, 0x6d, 0x65, 0x6e, 0x74, 0x76, 0x31, 0x62,
	0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_announcement_v1_type_proto_rawDescOnce sync.Once
	file_announcement_v1_type_proto_rawDescData = file_announcement_v1_type_proto_rawDesc
)

func file_announcement_v1_type_proto_rawDescGZIP() []byte {
	file_announcement_v1_type_proto_rawDescOnce.Do(func() {
		file_announcement_v1_type_proto_rawDescData = protoimpl.X.CompressGZIP(file_announcement_v1_type_proto_rawDescData)
	})
	return file_announcement_v1_type_proto_rawDescData
}

var file_announcement_v1_type_proto_enumTypes = make([]protoimpl.EnumInfo, 1)
var file_announcement_v1_type_proto_msgTypes = make([]protoimpl.MessageInfo, 1)
var file_announcement_v1_type_proto_goTypes = []interface{}{
	(AnnouncementTag)(0),             // 0: announcement.v1.AnnouncementTag
	(*Announcement)(nil),             // 1: announcement.v1.Announcement
	(*sharedpb.UUID)(nil),            // 2: shared.UUID
	(*sharedpb.RFC3339DateTime)(nil), // 3: shared.RFC3339DateTime
}
var file_announcement_v1_type_proto_depIdxs = []int32{
	2, // 0: announcement.v1.Announcement.id:type_name -> shared.UUID
	0, // 1: announcement.v1.Announcement.tags:type_name -> announcement.v1.AnnouncementTag
	3, // 2: announcement.v1.Announcement.published_at:type_name -> shared.RFC3339DateTime
	3, // [3:3] is the sub-list for method output_type
	3, // [3:3] is the sub-list for method input_type
	3, // [3:3] is the sub-list for extension type_name
	3, // [3:3] is the sub-list for extension extendee
	0, // [0:3] is the sub-list for field type_name
}

func init() { file_announcement_v1_type_proto_init() }
func file_announcement_v1_type_proto_init() {
	if File_announcement_v1_type_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_announcement_v1_type_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Announcement); i {
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
			RawDescriptor: file_announcement_v1_type_proto_rawDesc,
			NumEnums:      1,
			NumMessages:   1,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_announcement_v1_type_proto_goTypes,
		DependencyIndexes: file_announcement_v1_type_proto_depIdxs,
		EnumInfos:         file_announcement_v1_type_proto_enumTypes,
		MessageInfos:      file_announcement_v1_type_proto_msgTypes,
	}.Build()
	File_announcement_v1_type_proto = out.File
	file_announcement_v1_type_proto_rawDesc = nil
	file_announcement_v1_type_proto_goTypes = nil
	file_announcement_v1_type_proto_depIdxs = nil
}
