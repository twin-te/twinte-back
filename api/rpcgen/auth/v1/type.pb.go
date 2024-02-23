// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.32.0
// 	protoc        (unknown)
// source: auth/v1/type.proto

package authv1

import (
	sharedpb "github.com/twin-te/twinte-back/api/rpcgen/sharedpb"
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

type Provider int32

const (
	Provider_PROVIDER_UNSPECIFIED Provider = 0
	Provider_PROVIDER_GOOGLE      Provider = 1
	Provider_PROVIDER_TWITTER     Provider = 2
	Provider_PROVIDER_APPLE       Provider = 3
)

// Enum value maps for Provider.
var (
	Provider_name = map[int32]string{
		0: "PROVIDER_UNSPECIFIED",
		1: "PROVIDER_GOOGLE",
		2: "PROVIDER_TWITTER",
		3: "PROVIDER_APPLE",
	}
	Provider_value = map[string]int32{
		"PROVIDER_UNSPECIFIED": 0,
		"PROVIDER_GOOGLE":      1,
		"PROVIDER_TWITTER":     2,
		"PROVIDER_APPLE":       3,
	}
)

func (x Provider) Enum() *Provider {
	p := new(Provider)
	*p = x
	return p
}

func (x Provider) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (Provider) Descriptor() protoreflect.EnumDescriptor {
	return file_auth_v1_type_proto_enumTypes[0].Descriptor()
}

func (Provider) Type() protoreflect.EnumType {
	return &file_auth_v1_type_proto_enumTypes[0]
}

func (x Provider) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use Provider.Descriptor instead.
func (Provider) EnumDescriptor() ([]byte, []int) {
	return file_auth_v1_type_proto_rawDescGZIP(), []int{0}
}

type UserAuthentication struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Provider Provider `protobuf:"varint,1,opt,name=provider,proto3,enum=auth.v1.Provider" json:"provider,omitempty"`
	SocialId string   `protobuf:"bytes,2,opt,name=social_id,json=socialId,proto3" json:"social_id,omitempty"`
}

func (x *UserAuthentication) Reset() {
	*x = UserAuthentication{}
	if protoimpl.UnsafeEnabled {
		mi := &file_auth_v1_type_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *UserAuthentication) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*UserAuthentication) ProtoMessage() {}

func (x *UserAuthentication) ProtoReflect() protoreflect.Message {
	mi := &file_auth_v1_type_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use UserAuthentication.ProtoReflect.Descriptor instead.
func (*UserAuthentication) Descriptor() ([]byte, []int) {
	return file_auth_v1_type_proto_rawDescGZIP(), []int{0}
}

func (x *UserAuthentication) GetProvider() Provider {
	if x != nil {
		return x.Provider
	}
	return Provider_PROVIDER_UNSPECIFIED
}

func (x *UserAuthentication) GetSocialId() string {
	if x != nil {
		return x.SocialId
	}
	return ""
}

type User struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id              *sharedpb.UUID            `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
	Authentications []*UserAuthentication     `protobuf:"bytes,2,rep,name=authentications,proto3" json:"authentications,omitempty"`
	CreatedAt       *sharedpb.RFC3339DateTime `protobuf:"bytes,3,opt,name=created_at,json=createdAt,proto3" json:"created_at,omitempty"`
}

func (x *User) Reset() {
	*x = User{}
	if protoimpl.UnsafeEnabled {
		mi := &file_auth_v1_type_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *User) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*User) ProtoMessage() {}

func (x *User) ProtoReflect() protoreflect.Message {
	mi := &file_auth_v1_type_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use User.ProtoReflect.Descriptor instead.
func (*User) Descriptor() ([]byte, []int) {
	return file_auth_v1_type_proto_rawDescGZIP(), []int{1}
}

func (x *User) GetId() *sharedpb.UUID {
	if x != nil {
		return x.Id
	}
	return nil
}

func (x *User) GetAuthentications() []*UserAuthentication {
	if x != nil {
		return x.Authentications
	}
	return nil
}

func (x *User) GetCreatedAt() *sharedpb.RFC3339DateTime {
	if x != nil {
		return x.CreatedAt
	}
	return nil
}

var File_auth_v1_type_proto protoreflect.FileDescriptor

var file_auth_v1_type_proto_rawDesc = []byte{
	0x0a, 0x12, 0x61, 0x75, 0x74, 0x68, 0x2f, 0x76, 0x31, 0x2f, 0x74, 0x79, 0x70, 0x65, 0x2e, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x12, 0x07, 0x61, 0x75, 0x74, 0x68, 0x2e, 0x76, 0x31, 0x1a, 0x11, 0x73,
	0x68, 0x61, 0x72, 0x65, 0x64, 0x2f, 0x74, 0x79, 0x70, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x22, 0x60, 0x0a, 0x12, 0x55, 0x73, 0x65, 0x72, 0x41, 0x75, 0x74, 0x68, 0x65, 0x6e, 0x74, 0x69,
	0x63, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x12, 0x2d, 0x0a, 0x08, 0x70, 0x72, 0x6f, 0x76, 0x69, 0x64,
	0x65, 0x72, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0e, 0x32, 0x11, 0x2e, 0x61, 0x75, 0x74, 0x68, 0x2e,
	0x76, 0x31, 0x2e, 0x50, 0x72, 0x6f, 0x76, 0x69, 0x64, 0x65, 0x72, 0x52, 0x08, 0x70, 0x72, 0x6f,
	0x76, 0x69, 0x64, 0x65, 0x72, 0x12, 0x1b, 0x0a, 0x09, 0x73, 0x6f, 0x63, 0x69, 0x61, 0x6c, 0x5f,
	0x69, 0x64, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x73, 0x6f, 0x63, 0x69, 0x61, 0x6c,
	0x49, 0x64, 0x22, 0xa3, 0x01, 0x0a, 0x04, 0x55, 0x73, 0x65, 0x72, 0x12, 0x1c, 0x0a, 0x02, 0x69,
	0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x0c, 0x2e, 0x73, 0x68, 0x61, 0x72, 0x65, 0x64,
	0x2e, 0x55, 0x55, 0x49, 0x44, 0x52, 0x02, 0x69, 0x64, 0x12, 0x45, 0x0a, 0x0f, 0x61, 0x75, 0x74,
	0x68, 0x65, 0x6e, 0x74, 0x69, 0x63, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x18, 0x02, 0x20, 0x03,
	0x28, 0x0b, 0x32, 0x1b, 0x2e, 0x61, 0x75, 0x74, 0x68, 0x2e, 0x76, 0x31, 0x2e, 0x55, 0x73, 0x65,
	0x72, 0x41, 0x75, 0x74, 0x68, 0x65, 0x6e, 0x74, 0x69, 0x63, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x52,
	0x0f, 0x61, 0x75, 0x74, 0x68, 0x65, 0x6e, 0x74, 0x69, 0x63, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x73,
	0x12, 0x36, 0x0a, 0x0a, 0x63, 0x72, 0x65, 0x61, 0x74, 0x65, 0x64, 0x5f, 0x61, 0x74, 0x18, 0x03,
	0x20, 0x01, 0x28, 0x0b, 0x32, 0x17, 0x2e, 0x73, 0x68, 0x61, 0x72, 0x65, 0x64, 0x2e, 0x52, 0x46,
	0x43, 0x33, 0x33, 0x33, 0x39, 0x44, 0x61, 0x74, 0x65, 0x54, 0x69, 0x6d, 0x65, 0x52, 0x09, 0x63,
	0x72, 0x65, 0x61, 0x74, 0x65, 0x64, 0x41, 0x74, 0x2a, 0x63, 0x0a, 0x08, 0x50, 0x72, 0x6f, 0x76,
	0x69, 0x64, 0x65, 0x72, 0x12, 0x18, 0x0a, 0x14, 0x50, 0x52, 0x4f, 0x56, 0x49, 0x44, 0x45, 0x52,
	0x5f, 0x55, 0x4e, 0x53, 0x50, 0x45, 0x43, 0x49, 0x46, 0x49, 0x45, 0x44, 0x10, 0x00, 0x12, 0x13,
	0x0a, 0x0f, 0x50, 0x52, 0x4f, 0x56, 0x49, 0x44, 0x45, 0x52, 0x5f, 0x47, 0x4f, 0x4f, 0x47, 0x4c,
	0x45, 0x10, 0x01, 0x12, 0x14, 0x0a, 0x10, 0x50, 0x52, 0x4f, 0x56, 0x49, 0x44, 0x45, 0x52, 0x5f,
	0x54, 0x57, 0x49, 0x54, 0x54, 0x45, 0x52, 0x10, 0x02, 0x12, 0x12, 0x0a, 0x0e, 0x50, 0x52, 0x4f,
	0x56, 0x49, 0x44, 0x45, 0x52, 0x5f, 0x41, 0x50, 0x50, 0x4c, 0x45, 0x10, 0x03, 0x42, 0x3a, 0x5a,
	0x38, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x74, 0x77, 0x69, 0x6e,
	0x2d, 0x74, 0x65, 0x2f, 0x74, 0x77, 0x69, 0x6e, 0x74, 0x65, 0x2d, 0x62, 0x61, 0x63, 0x6b, 0x2f,
	0x61, 0x70, 0x69, 0x2f, 0x72, 0x70, 0x63, 0x67, 0x65, 0x6e, 0x2f, 0x61, 0x75, 0x74, 0x68, 0x2f,
	0x76, 0x31, 0x3b, 0x61, 0x75, 0x74, 0x68, 0x76, 0x31, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x33,
}

var (
	file_auth_v1_type_proto_rawDescOnce sync.Once
	file_auth_v1_type_proto_rawDescData = file_auth_v1_type_proto_rawDesc
)

func file_auth_v1_type_proto_rawDescGZIP() []byte {
	file_auth_v1_type_proto_rawDescOnce.Do(func() {
		file_auth_v1_type_proto_rawDescData = protoimpl.X.CompressGZIP(file_auth_v1_type_proto_rawDescData)
	})
	return file_auth_v1_type_proto_rawDescData
}

var file_auth_v1_type_proto_enumTypes = make([]protoimpl.EnumInfo, 1)
var file_auth_v1_type_proto_msgTypes = make([]protoimpl.MessageInfo, 2)
var file_auth_v1_type_proto_goTypes = []interface{}{
	(Provider)(0),                    // 0: auth.v1.Provider
	(*UserAuthentication)(nil),       // 1: auth.v1.UserAuthentication
	(*User)(nil),                     // 2: auth.v1.User
	(*sharedpb.UUID)(nil),            // 3: shared.UUID
	(*sharedpb.RFC3339DateTime)(nil), // 4: shared.RFC3339DateTime
}
var file_auth_v1_type_proto_depIdxs = []int32{
	0, // 0: auth.v1.UserAuthentication.provider:type_name -> auth.v1.Provider
	3, // 1: auth.v1.User.id:type_name -> shared.UUID
	1, // 2: auth.v1.User.authentications:type_name -> auth.v1.UserAuthentication
	4, // 3: auth.v1.User.created_at:type_name -> shared.RFC3339DateTime
	4, // [4:4] is the sub-list for method output_type
	4, // [4:4] is the sub-list for method input_type
	4, // [4:4] is the sub-list for extension type_name
	4, // [4:4] is the sub-list for extension extendee
	0, // [0:4] is the sub-list for field type_name
}

func init() { file_auth_v1_type_proto_init() }
func file_auth_v1_type_proto_init() {
	if File_auth_v1_type_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_auth_v1_type_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*UserAuthentication); i {
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
		file_auth_v1_type_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*User); i {
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
			RawDescriptor: file_auth_v1_type_proto_rawDesc,
			NumEnums:      1,
			NumMessages:   2,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_auth_v1_type_proto_goTypes,
		DependencyIndexes: file_auth_v1_type_proto_depIdxs,
		EnumInfos:         file_auth_v1_type_proto_enumTypes,
		MessageInfos:      file_auth_v1_type_proto_msgTypes,
	}.Build()
	File_auth_v1_type_proto = out.File
	file_auth_v1_type_proto_rawDesc = nil
	file_auth_v1_type_proto_goTypes = nil
	file_auth_v1_type_proto_depIdxs = nil
}
