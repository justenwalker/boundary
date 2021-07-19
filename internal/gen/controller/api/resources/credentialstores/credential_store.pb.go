// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.26.0
// 	protoc        v3.15.8
// source: controller/api/resources/credentialstores/v1/credential_store.proto

package credentialstores

import (
	scopes "github.com/hashicorp/boundary/internal/gen/controller/api/resources/scopes"
	_ "github.com/hashicorp/boundary/internal/gen/controller/protooptions"
	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
	structpb "google.golang.org/protobuf/types/known/structpb"
	timestamppb "google.golang.org/protobuf/types/known/timestamppb"
	wrapperspb "google.golang.org/protobuf/types/known/wrapperspb"
	reflect "reflect"
	sync "sync"
)

const (
	// Verify that this generated code is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(20 - protoimpl.MinVersion)
	// Verify that runtime/protoimpl is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(protoimpl.MaxVersion - 20)
)

// CredentialStore contains all fields related to an Credential Store resource
type CredentialStore struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// Output only. The ID of the Credential Store.
	Id string `protobuf:"bytes,10,opt,name=id,proto3" json:"id,omitempty"`
	// The ID of the Scope of which this Credential Store is a part.
	ScopeId string `protobuf:"bytes,20,opt,name=scope_id,proto3" json:"scope_id,omitempty"`
	// Output only. Scope information for this Credential Store.
	Scope *scopes.ScopeInfo `protobuf:"bytes,30,opt,name=scope,proto3" json:"scope,omitempty"`
	// Optional name for identification purposes.
	Name *wrapperspb.StringValue `protobuf:"bytes,40,opt,name=name,proto3" json:"name,omitempty"`
	// Optional user-set description for identification purposes.
	Description *wrapperspb.StringValue `protobuf:"bytes,50,opt,name=description,proto3" json:"description,omitempty"`
	// Output only. The time this resource was created.
	CreatedTime *timestamppb.Timestamp `protobuf:"bytes,60,opt,name=created_time,proto3" json:"created_time,omitempty"`
	// Output only. The time this resource was last updated.
	UpdatedTime *timestamppb.Timestamp `protobuf:"bytes,70,opt,name=updated_time,proto3" json:"updated_time,omitempty"`
	// Version is used in mutation requests, after the initial creation, to ensure this resource has not changed.
	// The mutation will fail if the version does not match the latest known good version.
	Version uint32 `protobuf:"varint,80,opt,name=version,proto3" json:"version,omitempty"`
	// The Credential Store type.
	Type string `protobuf:"bytes,90,opt,name=type,proto3" json:"type,omitempty"`
	// The attributes that are applicable for the specific Credential Store type.
	Attributes *structpb.Struct `protobuf:"bytes,100,opt,name=attributes,proto3" json:"attributes,omitempty"`
	// Output only. The available actions on this resource for this user.
	AuthorizedActions []string `protobuf:"bytes,300,rep,name=authorized_actions,proto3" json:"authorized_actions,omitempty"`
	// Output only. The authorized actions for the scope's collections.
	AuthorizedCollectionActions map[string]*structpb.ListValue `protobuf:"bytes,310,rep,name=authorized_collection_actions,proto3" json:"authorized_collection_actions,omitempty" protobuf_key:"bytes,1,opt,name=key,proto3" protobuf_val:"bytes,2,opt,name=value,proto3"`
}

func (x *CredentialStore) Reset() {
	*x = CredentialStore{}
	if protoimpl.UnsafeEnabled {
		mi := &file_controller_api_resources_credentialstores_v1_credential_store_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *CredentialStore) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CredentialStore) ProtoMessage() {}

func (x *CredentialStore) ProtoReflect() protoreflect.Message {
	mi := &file_controller_api_resources_credentialstores_v1_credential_store_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CredentialStore.ProtoReflect.Descriptor instead.
func (*CredentialStore) Descriptor() ([]byte, []int) {
	return file_controller_api_resources_credentialstores_v1_credential_store_proto_rawDescGZIP(), []int{0}
}

func (x *CredentialStore) GetId() string {
	if x != nil {
		return x.Id
	}
	return ""
}

func (x *CredentialStore) GetScopeId() string {
	if x != nil {
		return x.ScopeId
	}
	return ""
}

func (x *CredentialStore) GetScope() *scopes.ScopeInfo {
	if x != nil {
		return x.Scope
	}
	return nil
}

func (x *CredentialStore) GetName() *wrapperspb.StringValue {
	if x != nil {
		return x.Name
	}
	return nil
}

func (x *CredentialStore) GetDescription() *wrapperspb.StringValue {
	if x != nil {
		return x.Description
	}
	return nil
}

func (x *CredentialStore) GetCreatedTime() *timestamppb.Timestamp {
	if x != nil {
		return x.CreatedTime
	}
	return nil
}

func (x *CredentialStore) GetUpdatedTime() *timestamppb.Timestamp {
	if x != nil {
		return x.UpdatedTime
	}
	return nil
}

func (x *CredentialStore) GetVersion() uint32 {
	if x != nil {
		return x.Version
	}
	return 0
}

func (x *CredentialStore) GetType() string {
	if x != nil {
		return x.Type
	}
	return ""
}

func (x *CredentialStore) GetAttributes() *structpb.Struct {
	if x != nil {
		return x.Attributes
	}
	return nil
}

func (x *CredentialStore) GetAuthorizedActions() []string {
	if x != nil {
		return x.AuthorizedActions
	}
	return nil
}

func (x *CredentialStore) GetAuthorizedCollectionActions() map[string]*structpb.ListValue {
	if x != nil {
		return x.AuthorizedCollectionActions
	}
	return nil
}

// The attributes of a vault typed Credential Store.
type VaultCredentialStoreAttributes struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// The complete url address of vault.
	Address *wrapperspb.StringValue `protobuf:"bytes,10,opt,name=address,proto3" json:"address,omitempty"`
	// The namespace of vault used by this store
	Namespace *wrapperspb.StringValue `protobuf:"bytes,20,opt,name=namespace,proto3" json:"namespace,omitempty"`
	// The PEM encoded CA cert to verify the vault server's ssl certificate
	CaCert *wrapperspb.StringValue `protobuf:"bytes,30,opt,name=ca_cert,proto3" json:"ca_cert,omitempty"`
	// The value to use as the SNI host when connecting to vault via TLS.
	TlsServerName *wrapperspb.StringValue `protobuf:"bytes,40,opt,name=tls_server_name,proto3" json:"tls_server_name,omitempty"`
	// When set to true verification of the TLS certificate is disabled.
	TlsSkipVerify *wrapperspb.BoolValue `protobuf:"bytes,50,opt,name=tls_skip_verify,proto3" json:"tls_skip_verify,omitempty"`
	// Input only. The current vault token used by this credential store for creating new credentials.
	Token *wrapperspb.StringValue `protobuf:"bytes,60,opt,name=token,proto3" json:"token,omitempty"`
	// Output only. The hmac value of the vault token used by this credential store.
	TokenHmac string `protobuf:"bytes,70,opt,name=token_hmac,proto3" json:"token_hmac,omitempty"`
	// Input only. A PEM encoded client certificate for vault with an
	// optional private key included in the bundle.  It is an error to include
	// the private key in this bundle as well as setting the certificate_key
	// field.
	ClientCertificate *wrapperspb.StringValue `protobuf:"bytes,80,opt,name=client_certificate,proto3" json:"client_certificate,omitempty"`
	// Input only. A client certificate private key.
	ClientCertificateKey *wrapperspb.StringValue `protobuf:"bytes,90,opt,name=client_certificate_key,proto3" json:"client_certificate_key,omitempty"`
	// Output only. The hmac value of the private key used by the credential store.
	ClientCertificateKeyHmac string `protobuf:"bytes,100,opt,name=client_certificate_key_hmac,proto3" json:"client_certificate_key_hmac,omitempty"`
}

func (x *VaultCredentialStoreAttributes) Reset() {
	*x = VaultCredentialStoreAttributes{}
	if protoimpl.UnsafeEnabled {
		mi := &file_controller_api_resources_credentialstores_v1_credential_store_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *VaultCredentialStoreAttributes) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*VaultCredentialStoreAttributes) ProtoMessage() {}

func (x *VaultCredentialStoreAttributes) ProtoReflect() protoreflect.Message {
	mi := &file_controller_api_resources_credentialstores_v1_credential_store_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use VaultCredentialStoreAttributes.ProtoReflect.Descriptor instead.
func (*VaultCredentialStoreAttributes) Descriptor() ([]byte, []int) {
	return file_controller_api_resources_credentialstores_v1_credential_store_proto_rawDescGZIP(), []int{1}
}

func (x *VaultCredentialStoreAttributes) GetAddress() *wrapperspb.StringValue {
	if x != nil {
		return x.Address
	}
	return nil
}

func (x *VaultCredentialStoreAttributes) GetNamespace() *wrapperspb.StringValue {
	if x != nil {
		return x.Namespace
	}
	return nil
}

func (x *VaultCredentialStoreAttributes) GetCaCert() *wrapperspb.StringValue {
	if x != nil {
		return x.CaCert
	}
	return nil
}

func (x *VaultCredentialStoreAttributes) GetTlsServerName() *wrapperspb.StringValue {
	if x != nil {
		return x.TlsServerName
	}
	return nil
}

func (x *VaultCredentialStoreAttributes) GetTlsSkipVerify() *wrapperspb.BoolValue {
	if x != nil {
		return x.TlsSkipVerify
	}
	return nil
}

func (x *VaultCredentialStoreAttributes) GetToken() *wrapperspb.StringValue {
	if x != nil {
		return x.Token
	}
	return nil
}

func (x *VaultCredentialStoreAttributes) GetTokenHmac() string {
	if x != nil {
		return x.TokenHmac
	}
	return ""
}

func (x *VaultCredentialStoreAttributes) GetClientCertificate() *wrapperspb.StringValue {
	if x != nil {
		return x.ClientCertificate
	}
	return nil
}

func (x *VaultCredentialStoreAttributes) GetClientCertificateKey() *wrapperspb.StringValue {
	if x != nil {
		return x.ClientCertificateKey
	}
	return nil
}

func (x *VaultCredentialStoreAttributes) GetClientCertificateKeyHmac() string {
	if x != nil {
		return x.ClientCertificateKeyHmac
	}
	return ""
}

var File_controller_api_resources_credentialstores_v1_credential_store_proto protoreflect.FileDescriptor

var file_controller_api_resources_credentialstores_v1_credential_store_proto_rawDesc = []byte{
	0x0a, 0x43, 0x63, 0x6f, 0x6e, 0x74, 0x72, 0x6f, 0x6c, 0x6c, 0x65, 0x72, 0x2f, 0x61, 0x70, 0x69,
	0x2f, 0x72, 0x65, 0x73, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x73, 0x2f, 0x63, 0x72, 0x65, 0x64, 0x65,
	0x6e, 0x74, 0x69, 0x61, 0x6c, 0x73, 0x74, 0x6f, 0x72, 0x65, 0x73, 0x2f, 0x76, 0x31, 0x2f, 0x63,
	0x72, 0x65, 0x64, 0x65, 0x6e, 0x74, 0x69, 0x61, 0x6c, 0x5f, 0x73, 0x74, 0x6f, 0x72, 0x65, 0x2e,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x2c, 0x63, 0x6f, 0x6e, 0x74, 0x72, 0x6f, 0x6c, 0x6c, 0x65,
	0x72, 0x2e, 0x61, 0x70, 0x69, 0x2e, 0x72, 0x65, 0x73, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x73, 0x2e,
	0x63, 0x72, 0x65, 0x64, 0x65, 0x6e, 0x74, 0x69, 0x61, 0x6c, 0x73, 0x74, 0x6f, 0x72, 0x65, 0x73,
	0x2e, 0x76, 0x31, 0x1a, 0x1f, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x62, 0x75, 0x66, 0x2f, 0x74, 0x69, 0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d, 0x70, 0x2e, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x1e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x62, 0x75, 0x66, 0x2f, 0x77, 0x72, 0x61, 0x70, 0x70, 0x65, 0x72, 0x73, 0x2e, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x1c, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x62, 0x75, 0x66, 0x2f, 0x73, 0x74, 0x72, 0x75, 0x63, 0x74, 0x2e, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x1a, 0x2e, 0x63, 0x6f, 0x6e, 0x74, 0x72, 0x6f, 0x6c, 0x6c, 0x65, 0x72, 0x2f, 0x61,
	0x70, 0x69, 0x2f, 0x72, 0x65, 0x73, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x73, 0x2f, 0x73, 0x63, 0x6f,
	0x70, 0x65, 0x73, 0x2f, 0x76, 0x31, 0x2f, 0x73, 0x63, 0x6f, 0x70, 0x65, 0x2e, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x1a, 0x2a, 0x63, 0x6f, 0x6e, 0x74, 0x72, 0x6f, 0x6c, 0x6c, 0x65, 0x72, 0x2f, 0x63,
	0x75, 0x73, 0x74, 0x6f, 0x6d, 0x5f, 0x6f, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x2f, 0x76, 0x31,
	0x2f, 0x6f, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0xe0,
	0x06, 0x0a, 0x0f, 0x43, 0x72, 0x65, 0x64, 0x65, 0x6e, 0x74, 0x69, 0x61, 0x6c, 0x53, 0x74, 0x6f,
	0x72, 0x65, 0x12, 0x0e, 0x0a, 0x02, 0x69, 0x64, 0x18, 0x0a, 0x20, 0x01, 0x28, 0x09, 0x52, 0x02,
	0x69, 0x64, 0x12, 0x1a, 0x0a, 0x08, 0x73, 0x63, 0x6f, 0x70, 0x65, 0x5f, 0x69, 0x64, 0x18, 0x14,
	0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x73, 0x63, 0x6f, 0x70, 0x65, 0x5f, 0x69, 0x64, 0x12, 0x43,
	0x0a, 0x05, 0x73, 0x63, 0x6f, 0x70, 0x65, 0x18, 0x1e, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x2d, 0x2e,
	0x63, 0x6f, 0x6e, 0x74, 0x72, 0x6f, 0x6c, 0x6c, 0x65, 0x72, 0x2e, 0x61, 0x70, 0x69, 0x2e, 0x72,
	0x65, 0x73, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x73, 0x2e, 0x73, 0x63, 0x6f, 0x70, 0x65, 0x73, 0x2e,
	0x76, 0x31, 0x2e, 0x53, 0x63, 0x6f, 0x70, 0x65, 0x49, 0x6e, 0x66, 0x6f, 0x52, 0x05, 0x73, 0x63,
	0x6f, 0x70, 0x65, 0x12, 0x46, 0x0a, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x28, 0x20, 0x01, 0x28,
	0x0b, 0x32, 0x1c, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x62, 0x75, 0x66, 0x2e, 0x53, 0x74, 0x72, 0x69, 0x6e, 0x67, 0x56, 0x61, 0x6c, 0x75, 0x65, 0x42,
	0x14, 0xa0, 0xda, 0x29, 0x01, 0xc2, 0xdd, 0x29, 0x0c, 0x0a, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x12,
	0x04, 0x4e, 0x61, 0x6d, 0x65, 0x52, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x12, 0x62, 0x0a, 0x0b, 0x64,
	0x65, 0x73, 0x63, 0x72, 0x69, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x18, 0x32, 0x20, 0x01, 0x28, 0x0b,
	0x32, 0x1c, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62,
	0x75, 0x66, 0x2e, 0x53, 0x74, 0x72, 0x69, 0x6e, 0x67, 0x56, 0x61, 0x6c, 0x75, 0x65, 0x42, 0x22,
	0xa0, 0xda, 0x29, 0x01, 0xc2, 0xdd, 0x29, 0x1a, 0x0a, 0x0b, 0x64, 0x65, 0x73, 0x63, 0x72, 0x69,
	0x70, 0x74, 0x69, 0x6f, 0x6e, 0x12, 0x0b, 0x44, 0x65, 0x73, 0x63, 0x72, 0x69, 0x70, 0x74, 0x69,
	0x6f, 0x6e, 0x52, 0x0b, 0x64, 0x65, 0x73, 0x63, 0x72, 0x69, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x12,
	0x3e, 0x0a, 0x0c, 0x63, 0x72, 0x65, 0x61, 0x74, 0x65, 0x64, 0x5f, 0x74, 0x69, 0x6d, 0x65, 0x18,
	0x3c, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1a, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x54, 0x69, 0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d,
	0x70, 0x52, 0x0c, 0x63, 0x72, 0x65, 0x61, 0x74, 0x65, 0x64, 0x5f, 0x74, 0x69, 0x6d, 0x65, 0x12,
	0x3e, 0x0a, 0x0c, 0x75, 0x70, 0x64, 0x61, 0x74, 0x65, 0x64, 0x5f, 0x74, 0x69, 0x6d, 0x65, 0x18,
	0x46, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1a, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x54, 0x69, 0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d,
	0x70, 0x52, 0x0c, 0x75, 0x70, 0x64, 0x61, 0x74, 0x65, 0x64, 0x5f, 0x74, 0x69, 0x6d, 0x65, 0x12,
	0x18, 0x0a, 0x07, 0x76, 0x65, 0x72, 0x73, 0x69, 0x6f, 0x6e, 0x18, 0x50, 0x20, 0x01, 0x28, 0x0d,
	0x52, 0x07, 0x76, 0x65, 0x72, 0x73, 0x69, 0x6f, 0x6e, 0x12, 0x12, 0x0a, 0x04, 0x74, 0x79, 0x70,
	0x65, 0x18, 0x5a, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x74, 0x79, 0x70, 0x65, 0x12, 0x3d, 0x0a,
	0x0a, 0x61, 0x74, 0x74, 0x72, 0x69, 0x62, 0x75, 0x74, 0x65, 0x73, 0x18, 0x64, 0x20, 0x01, 0x28,
	0x0b, 0x32, 0x17, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x62, 0x75, 0x66, 0x2e, 0x53, 0x74, 0x72, 0x75, 0x63, 0x74, 0x42, 0x04, 0xa0, 0xda, 0x29, 0x01,
	0x52, 0x0a, 0x61, 0x74, 0x74, 0x72, 0x69, 0x62, 0x75, 0x74, 0x65, 0x73, 0x12, 0x2f, 0x0a, 0x12,
	0x61, 0x75, 0x74, 0x68, 0x6f, 0x72, 0x69, 0x7a, 0x65, 0x64, 0x5f, 0x61, 0x63, 0x74, 0x69, 0x6f,
	0x6e, 0x73, 0x18, 0xac, 0x02, 0x20, 0x03, 0x28, 0x09, 0x52, 0x12, 0x61, 0x75, 0x74, 0x68, 0x6f,
	0x72, 0x69, 0x7a, 0x65, 0x64, 0x5f, 0x61, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x12, 0xa5, 0x01,
	0x0a, 0x1d, 0x61, 0x75, 0x74, 0x68, 0x6f, 0x72, 0x69, 0x7a, 0x65, 0x64, 0x5f, 0x63, 0x6f, 0x6c,
	0x6c, 0x65, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x5f, 0x61, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x18,
	0xb6, 0x02, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x5e, 0x2e, 0x63, 0x6f, 0x6e, 0x74, 0x72, 0x6f, 0x6c,
	0x6c, 0x65, 0x72, 0x2e, 0x61, 0x70, 0x69, 0x2e, 0x72, 0x65, 0x73, 0x6f, 0x75, 0x72, 0x63, 0x65,
	0x73, 0x2e, 0x63, 0x72, 0x65, 0x64, 0x65, 0x6e, 0x74, 0x69, 0x61, 0x6c, 0x73, 0x74, 0x6f, 0x72,
	0x65, 0x73, 0x2e, 0x76, 0x31, 0x2e, 0x43, 0x72, 0x65, 0x64, 0x65, 0x6e, 0x74, 0x69, 0x61, 0x6c,
	0x53, 0x74, 0x6f, 0x72, 0x65, 0x2e, 0x41, 0x75, 0x74, 0x68, 0x6f, 0x72, 0x69, 0x7a, 0x65, 0x64,
	0x43, 0x6f, 0x6c, 0x6c, 0x65, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x41, 0x63, 0x74, 0x69, 0x6f, 0x6e,
	0x73, 0x45, 0x6e, 0x74, 0x72, 0x79, 0x52, 0x1d, 0x61, 0x75, 0x74, 0x68, 0x6f, 0x72, 0x69, 0x7a,
	0x65, 0x64, 0x5f, 0x63, 0x6f, 0x6c, 0x6c, 0x65, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x5f, 0x61, 0x63,
	0x74, 0x69, 0x6f, 0x6e, 0x73, 0x1a, 0x6a, 0x0a, 0x20, 0x41, 0x75, 0x74, 0x68, 0x6f, 0x72, 0x69,
	0x7a, 0x65, 0x64, 0x43, 0x6f, 0x6c, 0x6c, 0x65, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x41, 0x63, 0x74,
	0x69, 0x6f, 0x6e, 0x73, 0x45, 0x6e, 0x74, 0x72, 0x79, 0x12, 0x10, 0x0a, 0x03, 0x6b, 0x65, 0x79,
	0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x03, 0x6b, 0x65, 0x79, 0x12, 0x30, 0x0a, 0x05, 0x76,
	0x61, 0x6c, 0x75, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1a, 0x2e, 0x67, 0x6f, 0x6f,
	0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x4c, 0x69, 0x73,
	0x74, 0x56, 0x61, 0x6c, 0x75, 0x65, 0x52, 0x05, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x3a, 0x02, 0x38,
	0x01, 0x22, 0x93, 0x08, 0x0a, 0x1e, 0x56, 0x61, 0x75, 0x6c, 0x74, 0x43, 0x72, 0x65, 0x64, 0x65,
	0x6e, 0x74, 0x69, 0x61, 0x6c, 0x53, 0x74, 0x6f, 0x72, 0x65, 0x41, 0x74, 0x74, 0x72, 0x69, 0x62,
	0x75, 0x74, 0x65, 0x73, 0x12, 0x62, 0x0a, 0x07, 0x61, 0x64, 0x64, 0x72, 0x65, 0x73, 0x73, 0x18,
	0x0a, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1c, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x53, 0x74, 0x72, 0x69, 0x6e, 0x67, 0x56, 0x61,
	0x6c, 0x75, 0x65, 0x42, 0x2a, 0xa0, 0xda, 0x29, 0x01, 0xc2, 0xdd, 0x29, 0x22, 0x0a, 0x12, 0x61,
	0x74, 0x74, 0x72, 0x69, 0x62, 0x75, 0x74, 0x65, 0x73, 0x2e, 0x61, 0x64, 0x64, 0x72, 0x65, 0x73,
	0x73, 0x12, 0x0c, 0x56, 0x61, 0x75, 0x6c, 0x74, 0x41, 0x64, 0x64, 0x72, 0x65, 0x73, 0x73, 0x52,
	0x07, 0x61, 0x64, 0x64, 0x72, 0x65, 0x73, 0x73, 0x12, 0x65, 0x0a, 0x09, 0x6e, 0x61, 0x6d, 0x65,
	0x73, 0x70, 0x61, 0x63, 0x65, 0x18, 0x14, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1c, 0x2e, 0x67, 0x6f,
	0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x53, 0x74,
	0x72, 0x69, 0x6e, 0x67, 0x56, 0x61, 0x6c, 0x75, 0x65, 0x42, 0x29, 0xa0, 0xda, 0x29, 0x01, 0xc2,
	0xdd, 0x29, 0x21, 0x0a, 0x14, 0x61, 0x74, 0x74, 0x72, 0x69, 0x62, 0x75, 0x74, 0x65, 0x73, 0x2e,
	0x6e, 0x61, 0x6d, 0x65, 0x73, 0x70, 0x61, 0x63, 0x65, 0x12, 0x09, 0x4e, 0x61, 0x6d, 0x65, 0x73,
	0x70, 0x61, 0x63, 0x65, 0x52, 0x09, 0x6e, 0x61, 0x6d, 0x65, 0x73, 0x70, 0x61, 0x63, 0x65, 0x12,
	0x5c, 0x0a, 0x07, 0x63, 0x61, 0x5f, 0x63, 0x65, 0x72, 0x74, 0x18, 0x1e, 0x20, 0x01, 0x28, 0x0b,
	0x32, 0x1c, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62,
	0x75, 0x66, 0x2e, 0x53, 0x74, 0x72, 0x69, 0x6e, 0x67, 0x56, 0x61, 0x6c, 0x75, 0x65, 0x42, 0x24,
	0xa0, 0xda, 0x29, 0x01, 0xc2, 0xdd, 0x29, 0x1c, 0x0a, 0x12, 0x61, 0x74, 0x74, 0x72, 0x69, 0x62,
	0x75, 0x74, 0x65, 0x73, 0x2e, 0x63, 0x61, 0x5f, 0x63, 0x65, 0x72, 0x74, 0x12, 0x06, 0x43, 0x61,
	0x43, 0x65, 0x72, 0x74, 0x52, 0x07, 0x63, 0x61, 0x5f, 0x63, 0x65, 0x72, 0x74, 0x12, 0x7b, 0x0a,
	0x0f, 0x74, 0x6c, 0x73, 0x5f, 0x73, 0x65, 0x72, 0x76, 0x65, 0x72, 0x5f, 0x6e, 0x61, 0x6d, 0x65,
	0x18, 0x28, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1c, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x53, 0x74, 0x72, 0x69, 0x6e, 0x67, 0x56,
	0x61, 0x6c, 0x75, 0x65, 0x42, 0x33, 0xa0, 0xda, 0x29, 0x01, 0xc2, 0xdd, 0x29, 0x2b, 0x0a, 0x1a,
	0x61, 0x74, 0x74, 0x72, 0x69, 0x62, 0x75, 0x74, 0x65, 0x73, 0x2e, 0x74, 0x6c, 0x73, 0x5f, 0x73,
	0x65, 0x72, 0x76, 0x65, 0x72, 0x5f, 0x6e, 0x61, 0x6d, 0x65, 0x12, 0x0d, 0x54, 0x6c, 0x73, 0x53,
	0x65, 0x72, 0x76, 0x65, 0x72, 0x4e, 0x61, 0x6d, 0x65, 0x52, 0x0f, 0x74, 0x6c, 0x73, 0x5f, 0x73,
	0x65, 0x72, 0x76, 0x65, 0x72, 0x5f, 0x6e, 0x61, 0x6d, 0x65, 0x12, 0x79, 0x0a, 0x0f, 0x74, 0x6c,
	0x73, 0x5f, 0x73, 0x6b, 0x69, 0x70, 0x5f, 0x76, 0x65, 0x72, 0x69, 0x66, 0x79, 0x18, 0x32, 0x20,
	0x01, 0x28, 0x0b, 0x32, 0x1a, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x42, 0x6f, 0x6f, 0x6c, 0x56, 0x61, 0x6c, 0x75, 0x65, 0x42,
	0x33, 0xa0, 0xda, 0x29, 0x01, 0xc2, 0xdd, 0x29, 0x2b, 0x0a, 0x1a, 0x61, 0x74, 0x74, 0x72, 0x69,
	0x62, 0x75, 0x74, 0x65, 0x73, 0x2e, 0x74, 0x6c, 0x73, 0x5f, 0x73, 0x6b, 0x69, 0x70, 0x5f, 0x76,
	0x65, 0x72, 0x69, 0x66, 0x79, 0x12, 0x0d, 0x54, 0x6c, 0x73, 0x53, 0x6b, 0x69, 0x70, 0x56, 0x65,
	0x72, 0x69, 0x66, 0x79, 0x52, 0x0f, 0x74, 0x6c, 0x73, 0x5f, 0x73, 0x6b, 0x69, 0x70, 0x5f, 0x76,
	0x65, 0x72, 0x69, 0x66, 0x79, 0x12, 0x55, 0x0a, 0x05, 0x74, 0x6f, 0x6b, 0x65, 0x6e, 0x18, 0x3c,
	0x20, 0x01, 0x28, 0x0b, 0x32, 0x1c, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x53, 0x74, 0x72, 0x69, 0x6e, 0x67, 0x56, 0x61, 0x6c,
	0x75, 0x65, 0x42, 0x21, 0xa0, 0xda, 0x29, 0x01, 0xc2, 0xdd, 0x29, 0x19, 0x0a, 0x10, 0x61, 0x74,
	0x74, 0x72, 0x69, 0x62, 0x75, 0x74, 0x65, 0x73, 0x2e, 0x74, 0x6f, 0x6b, 0x65, 0x6e, 0x12, 0x05,
	0x54, 0x6f, 0x6b, 0x65, 0x6e, 0x52, 0x05, 0x74, 0x6f, 0x6b, 0x65, 0x6e, 0x12, 0x1e, 0x0a, 0x0a,
	0x74, 0x6f, 0x6b, 0x65, 0x6e, 0x5f, 0x68, 0x6d, 0x61, 0x63, 0x18, 0x46, 0x20, 0x01, 0x28, 0x09,
	0x52, 0x0a, 0x74, 0x6f, 0x6b, 0x65, 0x6e, 0x5f, 0x68, 0x6d, 0x61, 0x63, 0x12, 0x82, 0x01, 0x0a,
	0x12, 0x63, 0x6c, 0x69, 0x65, 0x6e, 0x74, 0x5f, 0x63, 0x65, 0x72, 0x74, 0x69, 0x66, 0x69, 0x63,
	0x61, 0x74, 0x65, 0x18, 0x50, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1c, 0x2e, 0x67, 0x6f, 0x6f, 0x67,
	0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x53, 0x74, 0x72, 0x69,
	0x6e, 0x67, 0x56, 0x61, 0x6c, 0x75, 0x65, 0x42, 0x34, 0xa0, 0xda, 0x29, 0x01, 0xc2, 0xdd, 0x29,
	0x2c, 0x0a, 0x1d, 0x61, 0x74, 0x74, 0x72, 0x69, 0x62, 0x75, 0x74, 0x65, 0x73, 0x2e, 0x63, 0x6c,
	0x69, 0x65, 0x6e, 0x74, 0x5f, 0x63, 0x65, 0x72, 0x74, 0x69, 0x66, 0x69, 0x63, 0x61, 0x74, 0x65,
	0x12, 0x0b, 0x43, 0x65, 0x72, 0x74, 0x69, 0x66, 0x69, 0x63, 0x61, 0x74, 0x65, 0x52, 0x12, 0x63,
	0x6c, 0x69, 0x65, 0x6e, 0x74, 0x5f, 0x63, 0x65, 0x72, 0x74, 0x69, 0x66, 0x69, 0x63, 0x61, 0x74,
	0x65, 0x12, 0x91, 0x01, 0x0a, 0x16, 0x63, 0x6c, 0x69, 0x65, 0x6e, 0x74, 0x5f, 0x63, 0x65, 0x72,
	0x74, 0x69, 0x66, 0x69, 0x63, 0x61, 0x74, 0x65, 0x5f, 0x6b, 0x65, 0x79, 0x18, 0x5a, 0x20, 0x01,
	0x28, 0x0b, 0x32, 0x1c, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x62, 0x75, 0x66, 0x2e, 0x53, 0x74, 0x72, 0x69, 0x6e, 0x67, 0x56, 0x61, 0x6c, 0x75, 0x65,
	0x42, 0x3b, 0xa0, 0xda, 0x29, 0x01, 0xc2, 0xdd, 0x29, 0x33, 0x0a, 0x21, 0x61, 0x74, 0x74, 0x72,
	0x69, 0x62, 0x75, 0x74, 0x65, 0x73, 0x2e, 0x63, 0x6c, 0x69, 0x65, 0x6e, 0x74, 0x5f, 0x63, 0x65,
	0x72, 0x74, 0x69, 0x66, 0x69, 0x63, 0x61, 0x74, 0x65, 0x5f, 0x6b, 0x65, 0x79, 0x12, 0x0e, 0x43,
	0x65, 0x72, 0x74, 0x69, 0x66, 0x69, 0x63, 0x61, 0x74, 0x65, 0x4b, 0x65, 0x79, 0x52, 0x16, 0x63,
	0x6c, 0x69, 0x65, 0x6e, 0x74, 0x5f, 0x63, 0x65, 0x72, 0x74, 0x69, 0x66, 0x69, 0x63, 0x61, 0x74,
	0x65, 0x5f, 0x6b, 0x65, 0x79, 0x12, 0x40, 0x0a, 0x1b, 0x63, 0x6c, 0x69, 0x65, 0x6e, 0x74, 0x5f,
	0x63, 0x65, 0x72, 0x74, 0x69, 0x66, 0x69, 0x63, 0x61, 0x74, 0x65, 0x5f, 0x6b, 0x65, 0x79, 0x5f,
	0x68, 0x6d, 0x61, 0x63, 0x18, 0x64, 0x20, 0x01, 0x28, 0x09, 0x52, 0x1b, 0x63, 0x6c, 0x69, 0x65,
	0x6e, 0x74, 0x5f, 0x63, 0x65, 0x72, 0x74, 0x69, 0x66, 0x69, 0x63, 0x61, 0x74, 0x65, 0x5f, 0x6b,
	0x65, 0x79, 0x5f, 0x68, 0x6d, 0x61, 0x63, 0x42, 0x67, 0x5a, 0x65, 0x67, 0x69, 0x74, 0x68, 0x75,
	0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x68, 0x61, 0x73, 0x68, 0x69, 0x63, 0x6f, 0x72, 0x70, 0x2f,
	0x62, 0x6f, 0x75, 0x6e, 0x64, 0x61, 0x72, 0x79, 0x2f, 0x69, 0x6e, 0x74, 0x65, 0x72, 0x6e, 0x61,
	0x6c, 0x2f, 0x67, 0x65, 0x6e, 0x2f, 0x63, 0x6f, 0x6e, 0x74, 0x72, 0x6f, 0x6c, 0x6c, 0x65, 0x72,
	0x2f, 0x61, 0x70, 0x69, 0x2f, 0x72, 0x65, 0x73, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x73, 0x2f, 0x63,
	0x72, 0x65, 0x64, 0x65, 0x6e, 0x74, 0x69, 0x61, 0x6c, 0x73, 0x74, 0x6f, 0x72, 0x65, 0x73, 0x3b,
	0x63, 0x72, 0x65, 0x64, 0x65, 0x6e, 0x74, 0x69, 0x61, 0x6c, 0x73, 0x74, 0x6f, 0x72, 0x65, 0x73,
	0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_controller_api_resources_credentialstores_v1_credential_store_proto_rawDescOnce sync.Once
	file_controller_api_resources_credentialstores_v1_credential_store_proto_rawDescData = file_controller_api_resources_credentialstores_v1_credential_store_proto_rawDesc
)

func file_controller_api_resources_credentialstores_v1_credential_store_proto_rawDescGZIP() []byte {
	file_controller_api_resources_credentialstores_v1_credential_store_proto_rawDescOnce.Do(func() {
		file_controller_api_resources_credentialstores_v1_credential_store_proto_rawDescData = protoimpl.X.CompressGZIP(file_controller_api_resources_credentialstores_v1_credential_store_proto_rawDescData)
	})
	return file_controller_api_resources_credentialstores_v1_credential_store_proto_rawDescData
}

var file_controller_api_resources_credentialstores_v1_credential_store_proto_msgTypes = make([]protoimpl.MessageInfo, 3)
var file_controller_api_resources_credentialstores_v1_credential_store_proto_goTypes = []interface{}{
	(*CredentialStore)(nil),                // 0: controller.api.resources.credentialstores.v1.CredentialStore
	(*VaultCredentialStoreAttributes)(nil), // 1: controller.api.resources.credentialstores.v1.VaultCredentialStoreAttributes
	nil,                                    // 2: controller.api.resources.credentialstores.v1.CredentialStore.AuthorizedCollectionActionsEntry
	(*scopes.ScopeInfo)(nil),               // 3: controller.api.resources.scopes.v1.ScopeInfo
	(*wrapperspb.StringValue)(nil),         // 4: google.protobuf.StringValue
	(*timestamppb.Timestamp)(nil),          // 5: google.protobuf.Timestamp
	(*structpb.Struct)(nil),                // 6: google.protobuf.Struct
	(*wrapperspb.BoolValue)(nil),           // 7: google.protobuf.BoolValue
	(*structpb.ListValue)(nil),             // 8: google.protobuf.ListValue
}
var file_controller_api_resources_credentialstores_v1_credential_store_proto_depIdxs = []int32{
	3,  // 0: controller.api.resources.credentialstores.v1.CredentialStore.scope:type_name -> controller.api.resources.scopes.v1.ScopeInfo
	4,  // 1: controller.api.resources.credentialstores.v1.CredentialStore.name:type_name -> google.protobuf.StringValue
	4,  // 2: controller.api.resources.credentialstores.v1.CredentialStore.description:type_name -> google.protobuf.StringValue
	5,  // 3: controller.api.resources.credentialstores.v1.CredentialStore.created_time:type_name -> google.protobuf.Timestamp
	5,  // 4: controller.api.resources.credentialstores.v1.CredentialStore.updated_time:type_name -> google.protobuf.Timestamp
	6,  // 5: controller.api.resources.credentialstores.v1.CredentialStore.attributes:type_name -> google.protobuf.Struct
	2,  // 6: controller.api.resources.credentialstores.v1.CredentialStore.authorized_collection_actions:type_name -> controller.api.resources.credentialstores.v1.CredentialStore.AuthorizedCollectionActionsEntry
	4,  // 7: controller.api.resources.credentialstores.v1.VaultCredentialStoreAttributes.address:type_name -> google.protobuf.StringValue
	4,  // 8: controller.api.resources.credentialstores.v1.VaultCredentialStoreAttributes.namespace:type_name -> google.protobuf.StringValue
	4,  // 9: controller.api.resources.credentialstores.v1.VaultCredentialStoreAttributes.ca_cert:type_name -> google.protobuf.StringValue
	4,  // 10: controller.api.resources.credentialstores.v1.VaultCredentialStoreAttributes.tls_server_name:type_name -> google.protobuf.StringValue
	7,  // 11: controller.api.resources.credentialstores.v1.VaultCredentialStoreAttributes.tls_skip_verify:type_name -> google.protobuf.BoolValue
	4,  // 12: controller.api.resources.credentialstores.v1.VaultCredentialStoreAttributes.token:type_name -> google.protobuf.StringValue
	4,  // 13: controller.api.resources.credentialstores.v1.VaultCredentialStoreAttributes.client_certificate:type_name -> google.protobuf.StringValue
	4,  // 14: controller.api.resources.credentialstores.v1.VaultCredentialStoreAttributes.client_certificate_key:type_name -> google.protobuf.StringValue
	8,  // 15: controller.api.resources.credentialstores.v1.CredentialStore.AuthorizedCollectionActionsEntry.value:type_name -> google.protobuf.ListValue
	16, // [16:16] is the sub-list for method output_type
	16, // [16:16] is the sub-list for method input_type
	16, // [16:16] is the sub-list for extension type_name
	16, // [16:16] is the sub-list for extension extendee
	0,  // [0:16] is the sub-list for field type_name
}

func init() { file_controller_api_resources_credentialstores_v1_credential_store_proto_init() }
func file_controller_api_resources_credentialstores_v1_credential_store_proto_init() {
	if File_controller_api_resources_credentialstores_v1_credential_store_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_controller_api_resources_credentialstores_v1_credential_store_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*CredentialStore); i {
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
		file_controller_api_resources_credentialstores_v1_credential_store_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*VaultCredentialStoreAttributes); i {
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
			RawDescriptor: file_controller_api_resources_credentialstores_v1_credential_store_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   3,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_controller_api_resources_credentialstores_v1_credential_store_proto_goTypes,
		DependencyIndexes: file_controller_api_resources_credentialstores_v1_credential_store_proto_depIdxs,
		MessageInfos:      file_controller_api_resources_credentialstores_v1_credential_store_proto_msgTypes,
	}.Build()
	File_controller_api_resources_credentialstores_v1_credential_store_proto = out.File
	file_controller_api_resources_credentialstores_v1_credential_store_proto_rawDesc = nil
	file_controller_api_resources_credentialstores_v1_credential_store_proto_goTypes = nil
	file_controller_api_resources_credentialstores_v1_credential_store_proto_depIdxs = nil
}
