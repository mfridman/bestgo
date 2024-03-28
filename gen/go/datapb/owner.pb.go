// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.33.0
// 	protoc        (unknown)
// source: datapb/owner.proto

package datapb

import (
	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
	timestamppb "google.golang.org/protobuf/types/known/timestamppb"
	reflect "reflect"
	sync "sync"
)

const (
	// Verify that this generated code is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(20 - protoimpl.MinVersion)
	// Verify that runtime/protoimpl is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(protoimpl.MaxVersion - 20)
)

type Owner struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	OwnerLogin           string                 `protobuf:"bytes,1,opt,name=owner_login,json=ownerLogin,proto3" json:"owner_login,omitempty"`
	OwnerId              int64                  `protobuf:"varint,2,opt,name=owner_id,json=ownerId,proto3" json:"owner_id,omitempty"`
	OwnerNodeId          string                 `protobuf:"bytes,3,opt,name=owner_node_id,json=ownerNodeId,proto3" json:"owner_node_id,omitempty"`
	OwnerGravatarId      string                 `protobuf:"bytes,4,opt,name=owner_gravatar_id,json=ownerGravatarId,proto3" json:"owner_gravatar_id,omitempty"`
	OwnerType            RepoOwnerType          `protobuf:"varint,5,opt,name=owner_type,json=ownerType,proto3,enum=datapb.RepoOwnerType" json:"owner_type,omitempty"`
	OwnerSiteAdmin       bool                   `protobuf:"varint,6,opt,name=owner_site_admin,json=ownerSiteAdmin,proto3" json:"owner_site_admin,omitempty"`
	OwnerName            string                 `protobuf:"bytes,7,opt,name=owner_name,json=ownerName,proto3" json:"owner_name,omitempty"`
	OwnerCompany         string                 `protobuf:"bytes,8,opt,name=owner_company,json=ownerCompany,proto3" json:"owner_company,omitempty"`
	OwnerBlog            string                 `protobuf:"bytes,9,opt,name=owner_blog,json=ownerBlog,proto3" json:"owner_blog,omitempty"`
	OwnerLocation        string                 `protobuf:"bytes,10,opt,name=owner_location,json=ownerLocation,proto3" json:"owner_location,omitempty"`
	OwnerEmail           string                 `protobuf:"bytes,11,opt,name=owner_email,json=ownerEmail,proto3" json:"owner_email,omitempty"`
	OwnerHireable        bool                   `protobuf:"varint,12,opt,name=owner_hireable,json=ownerHireable,proto3" json:"owner_hireable,omitempty"`
	OwnerBio             string                 `protobuf:"bytes,13,opt,name=owner_bio,json=ownerBio,proto3" json:"owner_bio,omitempty"`
	OwnerTwitterUsername string                 `protobuf:"bytes,14,opt,name=owner_twitter_username,json=ownerTwitterUsername,proto3" json:"owner_twitter_username,omitempty"`
	OwnerPublicRepos     int32                  `protobuf:"varint,15,opt,name=owner_public_repos,json=ownerPublicRepos,proto3" json:"owner_public_repos,omitempty"`
	OwnerPublicGists     int32                  `protobuf:"varint,16,opt,name=owner_public_gists,json=ownerPublicGists,proto3" json:"owner_public_gists,omitempty"`
	OwnerFollowers       int32                  `protobuf:"varint,17,opt,name=owner_followers,json=ownerFollowers,proto3" json:"owner_followers,omitempty"`
	OwnerFollowing       int32                  `protobuf:"varint,18,opt,name=owner_following,json=ownerFollowing,proto3" json:"owner_following,omitempty"`
	OwnerCreatedAt       *timestamppb.Timestamp `protobuf:"bytes,19,opt,name=owner_created_at,json=ownerCreatedAt,proto3" json:"owner_created_at,omitempty"`
	OwnerUpdatedAt       *timestamppb.Timestamp `protobuf:"bytes,20,opt,name=owner_updated_at,json=ownerUpdatedAt,proto3" json:"owner_updated_at,omitempty"`
}

func (x *Owner) Reset() {
	*x = Owner{}
	if protoimpl.UnsafeEnabled {
		mi := &file_datapb_owner_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Owner) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Owner) ProtoMessage() {}

func (x *Owner) ProtoReflect() protoreflect.Message {
	mi := &file_datapb_owner_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Owner.ProtoReflect.Descriptor instead.
func (*Owner) Descriptor() ([]byte, []int) {
	return file_datapb_owner_proto_rawDescGZIP(), []int{0}
}

func (x *Owner) GetOwnerLogin() string {
	if x != nil {
		return x.OwnerLogin
	}
	return ""
}

func (x *Owner) GetOwnerId() int64 {
	if x != nil {
		return x.OwnerId
	}
	return 0
}

func (x *Owner) GetOwnerNodeId() string {
	if x != nil {
		return x.OwnerNodeId
	}
	return ""
}

func (x *Owner) GetOwnerGravatarId() string {
	if x != nil {
		return x.OwnerGravatarId
	}
	return ""
}

func (x *Owner) GetOwnerType() RepoOwnerType {
	if x != nil {
		return x.OwnerType
	}
	return RepoOwnerType_REPO_OWNER_TYPE_UNSPECIFIED
}

func (x *Owner) GetOwnerSiteAdmin() bool {
	if x != nil {
		return x.OwnerSiteAdmin
	}
	return false
}

func (x *Owner) GetOwnerName() string {
	if x != nil {
		return x.OwnerName
	}
	return ""
}

func (x *Owner) GetOwnerCompany() string {
	if x != nil {
		return x.OwnerCompany
	}
	return ""
}

func (x *Owner) GetOwnerBlog() string {
	if x != nil {
		return x.OwnerBlog
	}
	return ""
}

func (x *Owner) GetOwnerLocation() string {
	if x != nil {
		return x.OwnerLocation
	}
	return ""
}

func (x *Owner) GetOwnerEmail() string {
	if x != nil {
		return x.OwnerEmail
	}
	return ""
}

func (x *Owner) GetOwnerHireable() bool {
	if x != nil {
		return x.OwnerHireable
	}
	return false
}

func (x *Owner) GetOwnerBio() string {
	if x != nil {
		return x.OwnerBio
	}
	return ""
}

func (x *Owner) GetOwnerTwitterUsername() string {
	if x != nil {
		return x.OwnerTwitterUsername
	}
	return ""
}

func (x *Owner) GetOwnerPublicRepos() int32 {
	if x != nil {
		return x.OwnerPublicRepos
	}
	return 0
}

func (x *Owner) GetOwnerPublicGists() int32 {
	if x != nil {
		return x.OwnerPublicGists
	}
	return 0
}

func (x *Owner) GetOwnerFollowers() int32 {
	if x != nil {
		return x.OwnerFollowers
	}
	return 0
}

func (x *Owner) GetOwnerFollowing() int32 {
	if x != nil {
		return x.OwnerFollowing
	}
	return 0
}

func (x *Owner) GetOwnerCreatedAt() *timestamppb.Timestamp {
	if x != nil {
		return x.OwnerCreatedAt
	}
	return nil
}

func (x *Owner) GetOwnerUpdatedAt() *timestamppb.Timestamp {
	if x != nil {
		return x.OwnerUpdatedAt
	}
	return nil
}

var File_datapb_owner_proto protoreflect.FileDescriptor

var file_datapb_owner_proto_rawDesc = []byte{
	0x0a, 0x12, 0x64, 0x61, 0x74, 0x61, 0x70, 0x62, 0x2f, 0x6f, 0x77, 0x6e, 0x65, 0x72, 0x2e, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x12, 0x06, 0x64, 0x61, 0x74, 0x61, 0x70, 0x62, 0x1a, 0x12, 0x64, 0x61,
	0x74, 0x61, 0x70, 0x62, 0x2f, 0x65, 0x6e, 0x75, 0x6d, 0x73, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x1a, 0x1f, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75,
	0x66, 0x2f, 0x74, 0x69, 0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d, 0x70, 0x2e, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x22, 0xd2, 0x06, 0x0a, 0x05, 0x4f, 0x77, 0x6e, 0x65, 0x72, 0x12, 0x1f, 0x0a, 0x0b, 0x6f,
	0x77, 0x6e, 0x65, 0x72, 0x5f, 0x6c, 0x6f, 0x67, 0x69, 0x6e, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09,
	0x52, 0x0a, 0x6f, 0x77, 0x6e, 0x65, 0x72, 0x4c, 0x6f, 0x67, 0x69, 0x6e, 0x12, 0x19, 0x0a, 0x08,
	0x6f, 0x77, 0x6e, 0x65, 0x72, 0x5f, 0x69, 0x64, 0x18, 0x02, 0x20, 0x01, 0x28, 0x03, 0x52, 0x07,
	0x6f, 0x77, 0x6e, 0x65, 0x72, 0x49, 0x64, 0x12, 0x22, 0x0a, 0x0d, 0x6f, 0x77, 0x6e, 0x65, 0x72,
	0x5f, 0x6e, 0x6f, 0x64, 0x65, 0x5f, 0x69, 0x64, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0b,
	0x6f, 0x77, 0x6e, 0x65, 0x72, 0x4e, 0x6f, 0x64, 0x65, 0x49, 0x64, 0x12, 0x2a, 0x0a, 0x11, 0x6f,
	0x77, 0x6e, 0x65, 0x72, 0x5f, 0x67, 0x72, 0x61, 0x76, 0x61, 0x74, 0x61, 0x72, 0x5f, 0x69, 0x64,
	0x18, 0x04, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0f, 0x6f, 0x77, 0x6e, 0x65, 0x72, 0x47, 0x72, 0x61,
	0x76, 0x61, 0x74, 0x61, 0x72, 0x49, 0x64, 0x12, 0x34, 0x0a, 0x0a, 0x6f, 0x77, 0x6e, 0x65, 0x72,
	0x5f, 0x74, 0x79, 0x70, 0x65, 0x18, 0x05, 0x20, 0x01, 0x28, 0x0e, 0x32, 0x15, 0x2e, 0x64, 0x61,
	0x74, 0x61, 0x70, 0x62, 0x2e, 0x52, 0x65, 0x70, 0x6f, 0x4f, 0x77, 0x6e, 0x65, 0x72, 0x54, 0x79,
	0x70, 0x65, 0x52, 0x09, 0x6f, 0x77, 0x6e, 0x65, 0x72, 0x54, 0x79, 0x70, 0x65, 0x12, 0x28, 0x0a,
	0x10, 0x6f, 0x77, 0x6e, 0x65, 0x72, 0x5f, 0x73, 0x69, 0x74, 0x65, 0x5f, 0x61, 0x64, 0x6d, 0x69,
	0x6e, 0x18, 0x06, 0x20, 0x01, 0x28, 0x08, 0x52, 0x0e, 0x6f, 0x77, 0x6e, 0x65, 0x72, 0x53, 0x69,
	0x74, 0x65, 0x41, 0x64, 0x6d, 0x69, 0x6e, 0x12, 0x1d, 0x0a, 0x0a, 0x6f, 0x77, 0x6e, 0x65, 0x72,
	0x5f, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x07, 0x20, 0x01, 0x28, 0x09, 0x52, 0x09, 0x6f, 0x77, 0x6e,
	0x65, 0x72, 0x4e, 0x61, 0x6d, 0x65, 0x12, 0x23, 0x0a, 0x0d, 0x6f, 0x77, 0x6e, 0x65, 0x72, 0x5f,
	0x63, 0x6f, 0x6d, 0x70, 0x61, 0x6e, 0x79, 0x18, 0x08, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0c, 0x6f,
	0x77, 0x6e, 0x65, 0x72, 0x43, 0x6f, 0x6d, 0x70, 0x61, 0x6e, 0x79, 0x12, 0x1d, 0x0a, 0x0a, 0x6f,
	0x77, 0x6e, 0x65, 0x72, 0x5f, 0x62, 0x6c, 0x6f, 0x67, 0x18, 0x09, 0x20, 0x01, 0x28, 0x09, 0x52,
	0x09, 0x6f, 0x77, 0x6e, 0x65, 0x72, 0x42, 0x6c, 0x6f, 0x67, 0x12, 0x25, 0x0a, 0x0e, 0x6f, 0x77,
	0x6e, 0x65, 0x72, 0x5f, 0x6c, 0x6f, 0x63, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x18, 0x0a, 0x20, 0x01,
	0x28, 0x09, 0x52, 0x0d, 0x6f, 0x77, 0x6e, 0x65, 0x72, 0x4c, 0x6f, 0x63, 0x61, 0x74, 0x69, 0x6f,
	0x6e, 0x12, 0x1f, 0x0a, 0x0b, 0x6f, 0x77, 0x6e, 0x65, 0x72, 0x5f, 0x65, 0x6d, 0x61, 0x69, 0x6c,
	0x18, 0x0b, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0a, 0x6f, 0x77, 0x6e, 0x65, 0x72, 0x45, 0x6d, 0x61,
	0x69, 0x6c, 0x12, 0x25, 0x0a, 0x0e, 0x6f, 0x77, 0x6e, 0x65, 0x72, 0x5f, 0x68, 0x69, 0x72, 0x65,
	0x61, 0x62, 0x6c, 0x65, 0x18, 0x0c, 0x20, 0x01, 0x28, 0x08, 0x52, 0x0d, 0x6f, 0x77, 0x6e, 0x65,
	0x72, 0x48, 0x69, 0x72, 0x65, 0x61, 0x62, 0x6c, 0x65, 0x12, 0x1b, 0x0a, 0x09, 0x6f, 0x77, 0x6e,
	0x65, 0x72, 0x5f, 0x62, 0x69, 0x6f, 0x18, 0x0d, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x6f, 0x77,
	0x6e, 0x65, 0x72, 0x42, 0x69, 0x6f, 0x12, 0x34, 0x0a, 0x16, 0x6f, 0x77, 0x6e, 0x65, 0x72, 0x5f,
	0x74, 0x77, 0x69, 0x74, 0x74, 0x65, 0x72, 0x5f, 0x75, 0x73, 0x65, 0x72, 0x6e, 0x61, 0x6d, 0x65,
	0x18, 0x0e, 0x20, 0x01, 0x28, 0x09, 0x52, 0x14, 0x6f, 0x77, 0x6e, 0x65, 0x72, 0x54, 0x77, 0x69,
	0x74, 0x74, 0x65, 0x72, 0x55, 0x73, 0x65, 0x72, 0x6e, 0x61, 0x6d, 0x65, 0x12, 0x2c, 0x0a, 0x12,
	0x6f, 0x77, 0x6e, 0x65, 0x72, 0x5f, 0x70, 0x75, 0x62, 0x6c, 0x69, 0x63, 0x5f, 0x72, 0x65, 0x70,
	0x6f, 0x73, 0x18, 0x0f, 0x20, 0x01, 0x28, 0x05, 0x52, 0x10, 0x6f, 0x77, 0x6e, 0x65, 0x72, 0x50,
	0x75, 0x62, 0x6c, 0x69, 0x63, 0x52, 0x65, 0x70, 0x6f, 0x73, 0x12, 0x2c, 0x0a, 0x12, 0x6f, 0x77,
	0x6e, 0x65, 0x72, 0x5f, 0x70, 0x75, 0x62, 0x6c, 0x69, 0x63, 0x5f, 0x67, 0x69, 0x73, 0x74, 0x73,
	0x18, 0x10, 0x20, 0x01, 0x28, 0x05, 0x52, 0x10, 0x6f, 0x77, 0x6e, 0x65, 0x72, 0x50, 0x75, 0x62,
	0x6c, 0x69, 0x63, 0x47, 0x69, 0x73, 0x74, 0x73, 0x12, 0x27, 0x0a, 0x0f, 0x6f, 0x77, 0x6e, 0x65,
	0x72, 0x5f, 0x66, 0x6f, 0x6c, 0x6c, 0x6f, 0x77, 0x65, 0x72, 0x73, 0x18, 0x11, 0x20, 0x01, 0x28,
	0x05, 0x52, 0x0e, 0x6f, 0x77, 0x6e, 0x65, 0x72, 0x46, 0x6f, 0x6c, 0x6c, 0x6f, 0x77, 0x65, 0x72,
	0x73, 0x12, 0x27, 0x0a, 0x0f, 0x6f, 0x77, 0x6e, 0x65, 0x72, 0x5f, 0x66, 0x6f, 0x6c, 0x6c, 0x6f,
	0x77, 0x69, 0x6e, 0x67, 0x18, 0x12, 0x20, 0x01, 0x28, 0x05, 0x52, 0x0e, 0x6f, 0x77, 0x6e, 0x65,
	0x72, 0x46, 0x6f, 0x6c, 0x6c, 0x6f, 0x77, 0x69, 0x6e, 0x67, 0x12, 0x44, 0x0a, 0x10, 0x6f, 0x77,
	0x6e, 0x65, 0x72, 0x5f, 0x63, 0x72, 0x65, 0x61, 0x74, 0x65, 0x64, 0x5f, 0x61, 0x74, 0x18, 0x13,
	0x20, 0x01, 0x28, 0x0b, 0x32, 0x1a, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x54, 0x69, 0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d, 0x70,
	0x52, 0x0e, 0x6f, 0x77, 0x6e, 0x65, 0x72, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x64, 0x41, 0x74,
	0x12, 0x44, 0x0a, 0x10, 0x6f, 0x77, 0x6e, 0x65, 0x72, 0x5f, 0x75, 0x70, 0x64, 0x61, 0x74, 0x65,
	0x64, 0x5f, 0x61, 0x74, 0x18, 0x14, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1a, 0x2e, 0x67, 0x6f, 0x6f,
	0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x54, 0x69, 0x6d,
	0x65, 0x73, 0x74, 0x61, 0x6d, 0x70, 0x52, 0x0e, 0x6f, 0x77, 0x6e, 0x65, 0x72, 0x55, 0x70, 0x64,
	0x61, 0x74, 0x65, 0x64, 0x41, 0x74, 0x42, 0x7a, 0x0a, 0x0a, 0x63, 0x6f, 0x6d, 0x2e, 0x64, 0x61,
	0x74, 0x61, 0x70, 0x62, 0x42, 0x0a, 0x4f, 0x77, 0x6e, 0x65, 0x72, 0x50, 0x72, 0x6f, 0x74, 0x6f,
	0x50, 0x01, 0x5a, 0x28, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x6d,
	0x66, 0x72, 0x69, 0x64, 0x6d, 0x61, 0x6e, 0x2f, 0x62, 0x65, 0x73, 0x74, 0x67, 0x6f, 0x2f, 0x67,
	0x65, 0x6e, 0x2f, 0x67, 0x6f, 0x2f, 0x64, 0x61, 0x74, 0x61, 0x70, 0x62, 0xa2, 0x02, 0x03, 0x44,
	0x58, 0x58, 0xaa, 0x02, 0x06, 0x44, 0x61, 0x74, 0x61, 0x70, 0x62, 0xca, 0x02, 0x06, 0x44, 0x61,
	0x74, 0x61, 0x70, 0x62, 0xe2, 0x02, 0x12, 0x44, 0x61, 0x74, 0x61, 0x70, 0x62, 0x5c, 0x47, 0x50,
	0x42, 0x4d, 0x65, 0x74, 0x61, 0x64, 0x61, 0x74, 0x61, 0xea, 0x02, 0x06, 0x44, 0x61, 0x74, 0x61,
	0x70, 0x62, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_datapb_owner_proto_rawDescOnce sync.Once
	file_datapb_owner_proto_rawDescData = file_datapb_owner_proto_rawDesc
)

func file_datapb_owner_proto_rawDescGZIP() []byte {
	file_datapb_owner_proto_rawDescOnce.Do(func() {
		file_datapb_owner_proto_rawDescData = protoimpl.X.CompressGZIP(file_datapb_owner_proto_rawDescData)
	})
	return file_datapb_owner_proto_rawDescData
}

var file_datapb_owner_proto_msgTypes = make([]protoimpl.MessageInfo, 1)
var file_datapb_owner_proto_goTypes = []interface{}{
	(*Owner)(nil),                 // 0: datapb.Owner
	(RepoOwnerType)(0),            // 1: datapb.RepoOwnerType
	(*timestamppb.Timestamp)(nil), // 2: google.protobuf.Timestamp
}
var file_datapb_owner_proto_depIdxs = []int32{
	1, // 0: datapb.Owner.owner_type:type_name -> datapb.RepoOwnerType
	2, // 1: datapb.Owner.owner_created_at:type_name -> google.protobuf.Timestamp
	2, // 2: datapb.Owner.owner_updated_at:type_name -> google.protobuf.Timestamp
	3, // [3:3] is the sub-list for method output_type
	3, // [3:3] is the sub-list for method input_type
	3, // [3:3] is the sub-list for extension type_name
	3, // [3:3] is the sub-list for extension extendee
	0, // [0:3] is the sub-list for field type_name
}

func init() { file_datapb_owner_proto_init() }
func file_datapb_owner_proto_init() {
	if File_datapb_owner_proto != nil {
		return
	}
	file_datapb_enums_proto_init()
	if !protoimpl.UnsafeEnabled {
		file_datapb_owner_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Owner); i {
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
			RawDescriptor: file_datapb_owner_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   1,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_datapb_owner_proto_goTypes,
		DependencyIndexes: file_datapb_owner_proto_depIdxs,
		MessageInfos:      file_datapb_owner_proto_msgTypes,
	}.Build()
	File_datapb_owner_proto = out.File
	file_datapb_owner_proto_rawDesc = nil
	file_datapb_owner_proto_goTypes = nil
	file_datapb_owner_proto_depIdxs = nil
}
