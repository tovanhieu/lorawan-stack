// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: lorawan-stack/api/oauth.proto

package ttnpb

import (
	fmt "fmt"
	_ "github.com/envoyproxy/protoc-gen-validate/validate"
	_ "github.com/gogo/protobuf/gogoproto"
	proto "github.com/gogo/protobuf/proto"
	types "github.com/gogo/protobuf/types"
	golang_proto "github.com/golang/protobuf/proto"
	math "math"
)

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = golang_proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// This is a compile-time assertion to ensure that this generated file
// is compatible with the proto package it is being compiled against.
// A compilation error at this line likely means your copy of the
// proto package needs to be updated.
const _ = proto.GoGoProtoPackageIsVersion3 // please upgrade the proto package

type OAuthClientAuthorizationIdentifiers struct {
	UserIds              *UserIdentifiers   `protobuf:"bytes,1,opt,name=user_ids,json=userIds,proto3" json:"user_ids,omitempty"`
	ClientIds            *ClientIdentifiers `protobuf:"bytes,2,opt,name=client_ids,json=clientIds,proto3" json:"client_ids,omitempty"`
	XXX_NoUnkeyedLiteral struct{}           `json:"-"`
	XXX_sizecache        int32              `json:"-"`
}

func (m *OAuthClientAuthorizationIdentifiers) Reset()         { *m = OAuthClientAuthorizationIdentifiers{} }
func (m *OAuthClientAuthorizationIdentifiers) String() string { return proto.CompactTextString(m) }
func (*OAuthClientAuthorizationIdentifiers) ProtoMessage()    {}
func (*OAuthClientAuthorizationIdentifiers) Descriptor() ([]byte, []int) {
	return fileDescriptor_1454904971eaa7d7, []int{0}
}
func (m *OAuthClientAuthorizationIdentifiers) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_OAuthClientAuthorizationIdentifiers.Unmarshal(m, b)
}
func (m *OAuthClientAuthorizationIdentifiers) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_OAuthClientAuthorizationIdentifiers.Marshal(b, m, deterministic)
}
func (m *OAuthClientAuthorizationIdentifiers) XXX_Merge(src proto.Message) {
	xxx_messageInfo_OAuthClientAuthorizationIdentifiers.Merge(m, src)
}
func (m *OAuthClientAuthorizationIdentifiers) XXX_Size() int {
	return xxx_messageInfo_OAuthClientAuthorizationIdentifiers.Size(m)
}
func (m *OAuthClientAuthorizationIdentifiers) XXX_DiscardUnknown() {
	xxx_messageInfo_OAuthClientAuthorizationIdentifiers.DiscardUnknown(m)
}

var xxx_messageInfo_OAuthClientAuthorizationIdentifiers proto.InternalMessageInfo

func (m *OAuthClientAuthorizationIdentifiers) GetUserIds() *UserIdentifiers {
	if m != nil {
		return m.UserIds
	}
	return nil
}

func (m *OAuthClientAuthorizationIdentifiers) GetClientIds() *ClientIdentifiers {
	if m != nil {
		return m.ClientIds
	}
	return nil
}

type OAuthClientAuthorization struct {
	UserIds              *UserIdentifiers   `protobuf:"bytes,1,opt,name=user_ids,json=userIds,proto3" json:"user_ids,omitempty"`
	ClientIds            *ClientIdentifiers `protobuf:"bytes,2,opt,name=client_ids,json=clientIds,proto3" json:"client_ids,omitempty"`
	Rights               []Right            `protobuf:"varint,3,rep,packed,name=rights,proto3,enum=ttn.lorawan.v3.Right" json:"rights,omitempty"`
	CreatedAt            *types.Timestamp   `protobuf:"bytes,4,opt,name=created_at,json=createdAt,proto3" json:"created_at,omitempty"`
	UpdatedAt            *types.Timestamp   `protobuf:"bytes,5,opt,name=updated_at,json=updatedAt,proto3" json:"updated_at,omitempty"`
	XXX_NoUnkeyedLiteral struct{}           `json:"-"`
	XXX_sizecache        int32              `json:"-"`
}

func (m *OAuthClientAuthorization) Reset()         { *m = OAuthClientAuthorization{} }
func (m *OAuthClientAuthorization) String() string { return proto.CompactTextString(m) }
func (*OAuthClientAuthorization) ProtoMessage()    {}
func (*OAuthClientAuthorization) Descriptor() ([]byte, []int) {
	return fileDescriptor_1454904971eaa7d7, []int{1}
}
func (m *OAuthClientAuthorization) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_OAuthClientAuthorization.Unmarshal(m, b)
}
func (m *OAuthClientAuthorization) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_OAuthClientAuthorization.Marshal(b, m, deterministic)
}
func (m *OAuthClientAuthorization) XXX_Merge(src proto.Message) {
	xxx_messageInfo_OAuthClientAuthorization.Merge(m, src)
}
func (m *OAuthClientAuthorization) XXX_Size() int {
	return xxx_messageInfo_OAuthClientAuthorization.Size(m)
}
func (m *OAuthClientAuthorization) XXX_DiscardUnknown() {
	xxx_messageInfo_OAuthClientAuthorization.DiscardUnknown(m)
}

var xxx_messageInfo_OAuthClientAuthorization proto.InternalMessageInfo

func (m *OAuthClientAuthorization) GetUserIds() *UserIdentifiers {
	if m != nil {
		return m.UserIds
	}
	return nil
}

func (m *OAuthClientAuthorization) GetClientIds() *ClientIdentifiers {
	if m != nil {
		return m.ClientIds
	}
	return nil
}

func (m *OAuthClientAuthorization) GetRights() []Right {
	if m != nil {
		return m.Rights
	}
	return nil
}

func (m *OAuthClientAuthorization) GetCreatedAt() *types.Timestamp {
	if m != nil {
		return m.CreatedAt
	}
	return nil
}

func (m *OAuthClientAuthorization) GetUpdatedAt() *types.Timestamp {
	if m != nil {
		return m.UpdatedAt
	}
	return nil
}

type OAuthClientAuthorizations struct {
	Authorizations       []*OAuthClientAuthorization `protobuf:"bytes,1,rep,name=authorizations,proto3" json:"authorizations,omitempty"`
	XXX_NoUnkeyedLiteral struct{}                    `json:"-"`
	XXX_sizecache        int32                       `json:"-"`
}

func (m *OAuthClientAuthorizations) Reset()         { *m = OAuthClientAuthorizations{} }
func (m *OAuthClientAuthorizations) String() string { return proto.CompactTextString(m) }
func (*OAuthClientAuthorizations) ProtoMessage()    {}
func (*OAuthClientAuthorizations) Descriptor() ([]byte, []int) {
	return fileDescriptor_1454904971eaa7d7, []int{2}
}
func (m *OAuthClientAuthorizations) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_OAuthClientAuthorizations.Unmarshal(m, b)
}
func (m *OAuthClientAuthorizations) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_OAuthClientAuthorizations.Marshal(b, m, deterministic)
}
func (m *OAuthClientAuthorizations) XXX_Merge(src proto.Message) {
	xxx_messageInfo_OAuthClientAuthorizations.Merge(m, src)
}
func (m *OAuthClientAuthorizations) XXX_Size() int {
	return xxx_messageInfo_OAuthClientAuthorizations.Size(m)
}
func (m *OAuthClientAuthorizations) XXX_DiscardUnknown() {
	xxx_messageInfo_OAuthClientAuthorizations.DiscardUnknown(m)
}

var xxx_messageInfo_OAuthClientAuthorizations proto.InternalMessageInfo

func (m *OAuthClientAuthorizations) GetAuthorizations() []*OAuthClientAuthorization {
	if m != nil {
		return m.Authorizations
	}
	return nil
}

type ListOAuthClientAuthorizationsRequest struct {
	UserIds *UserIdentifiers `protobuf:"bytes,1,opt,name=user_ids,json=userIds,proto3" json:"user_ids,omitempty"`
	// Order the results by this field path (must be present in the field mask).
	// Default ordering is by ID. Prepend with a minus (-) to reverse the order.
	Order string `protobuf:"bytes,2,opt,name=order,proto3" json:"order,omitempty"`
	// Limit the number of results per page.
	Limit uint32 `protobuf:"varint,3,opt,name=limit,proto3" json:"limit,omitempty"`
	// Page number for pagination. 0 is interpreted as 1.
	Page                 uint32   `protobuf:"varint,4,opt,name=page,proto3" json:"page,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *ListOAuthClientAuthorizationsRequest) Reset()         { *m = ListOAuthClientAuthorizationsRequest{} }
func (m *ListOAuthClientAuthorizationsRequest) String() string { return proto.CompactTextString(m) }
func (*ListOAuthClientAuthorizationsRequest) ProtoMessage()    {}
func (*ListOAuthClientAuthorizationsRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_1454904971eaa7d7, []int{3}
}
func (m *ListOAuthClientAuthorizationsRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_ListOAuthClientAuthorizationsRequest.Unmarshal(m, b)
}
func (m *ListOAuthClientAuthorizationsRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_ListOAuthClientAuthorizationsRequest.Marshal(b, m, deterministic)
}
func (m *ListOAuthClientAuthorizationsRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ListOAuthClientAuthorizationsRequest.Merge(m, src)
}
func (m *ListOAuthClientAuthorizationsRequest) XXX_Size() int {
	return xxx_messageInfo_ListOAuthClientAuthorizationsRequest.Size(m)
}
func (m *ListOAuthClientAuthorizationsRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_ListOAuthClientAuthorizationsRequest.DiscardUnknown(m)
}

var xxx_messageInfo_ListOAuthClientAuthorizationsRequest proto.InternalMessageInfo

func (m *ListOAuthClientAuthorizationsRequest) GetUserIds() *UserIdentifiers {
	if m != nil {
		return m.UserIds
	}
	return nil
}

func (m *ListOAuthClientAuthorizationsRequest) GetOrder() string {
	if m != nil {
		return m.Order
	}
	return ""
}

func (m *ListOAuthClientAuthorizationsRequest) GetLimit() uint32 {
	if m != nil {
		return m.Limit
	}
	return 0
}

func (m *ListOAuthClientAuthorizationsRequest) GetPage() uint32 {
	if m != nil {
		return m.Page
	}
	return 0
}

type OAuthAuthorizationCode struct {
	UserIds              *UserIdentifiers   `protobuf:"bytes,1,opt,name=user_ids,json=userIds,proto3" json:"user_ids,omitempty"`
	UserSessionId        string             `protobuf:"bytes,9,opt,name=user_session_id,json=userSessionId,proto3" json:"user_session_id,omitempty"`
	ClientIds            *ClientIdentifiers `protobuf:"bytes,2,opt,name=client_ids,json=clientIds,proto3" json:"client_ids,omitempty"`
	Rights               []Right            `protobuf:"varint,3,rep,packed,name=rights,proto3,enum=ttn.lorawan.v3.Right" json:"rights,omitempty"`
	Code                 string             `protobuf:"bytes,4,opt,name=code,proto3" json:"code,omitempty"`
	RedirectUri          string             `protobuf:"bytes,5,opt,name=redirect_uri,json=redirectUri,proto3" json:"redirect_uri,omitempty"`
	State                string             `protobuf:"bytes,6,opt,name=state,proto3" json:"state,omitempty"`
	CreatedAt            *types.Timestamp   `protobuf:"bytes,7,opt,name=created_at,json=createdAt,proto3" json:"created_at,omitempty"`
	ExpiresAt            *types.Timestamp   `protobuf:"bytes,8,opt,name=expires_at,json=expiresAt,proto3" json:"expires_at,omitempty"`
	XXX_NoUnkeyedLiteral struct{}           `json:"-"`
	XXX_sizecache        int32              `json:"-"`
}

func (m *OAuthAuthorizationCode) Reset()         { *m = OAuthAuthorizationCode{} }
func (m *OAuthAuthorizationCode) String() string { return proto.CompactTextString(m) }
func (*OAuthAuthorizationCode) ProtoMessage()    {}
func (*OAuthAuthorizationCode) Descriptor() ([]byte, []int) {
	return fileDescriptor_1454904971eaa7d7, []int{4}
}
func (m *OAuthAuthorizationCode) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_OAuthAuthorizationCode.Unmarshal(m, b)
}
func (m *OAuthAuthorizationCode) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_OAuthAuthorizationCode.Marshal(b, m, deterministic)
}
func (m *OAuthAuthorizationCode) XXX_Merge(src proto.Message) {
	xxx_messageInfo_OAuthAuthorizationCode.Merge(m, src)
}
func (m *OAuthAuthorizationCode) XXX_Size() int {
	return xxx_messageInfo_OAuthAuthorizationCode.Size(m)
}
func (m *OAuthAuthorizationCode) XXX_DiscardUnknown() {
	xxx_messageInfo_OAuthAuthorizationCode.DiscardUnknown(m)
}

var xxx_messageInfo_OAuthAuthorizationCode proto.InternalMessageInfo

func (m *OAuthAuthorizationCode) GetUserIds() *UserIdentifiers {
	if m != nil {
		return m.UserIds
	}
	return nil
}

func (m *OAuthAuthorizationCode) GetUserSessionId() string {
	if m != nil {
		return m.UserSessionId
	}
	return ""
}

func (m *OAuthAuthorizationCode) GetClientIds() *ClientIdentifiers {
	if m != nil {
		return m.ClientIds
	}
	return nil
}

func (m *OAuthAuthorizationCode) GetRights() []Right {
	if m != nil {
		return m.Rights
	}
	return nil
}

func (m *OAuthAuthorizationCode) GetCode() string {
	if m != nil {
		return m.Code
	}
	return ""
}

func (m *OAuthAuthorizationCode) GetRedirectUri() string {
	if m != nil {
		return m.RedirectUri
	}
	return ""
}

func (m *OAuthAuthorizationCode) GetState() string {
	if m != nil {
		return m.State
	}
	return ""
}

func (m *OAuthAuthorizationCode) GetCreatedAt() *types.Timestamp {
	if m != nil {
		return m.CreatedAt
	}
	return nil
}

func (m *OAuthAuthorizationCode) GetExpiresAt() *types.Timestamp {
	if m != nil {
		return m.ExpiresAt
	}
	return nil
}

type OAuthAccessTokenIdentifiers struct {
	UserIds              *UserIdentifiers   `protobuf:"bytes,1,opt,name=user_ids,json=userIds,proto3" json:"user_ids,omitempty"`
	ClientIds            *ClientIdentifiers `protobuf:"bytes,2,opt,name=client_ids,json=clientIds,proto3" json:"client_ids,omitempty"`
	Id                   string             `protobuf:"bytes,3,opt,name=id,proto3" json:"id,omitempty"`
	XXX_NoUnkeyedLiteral struct{}           `json:"-"`
	XXX_sizecache        int32              `json:"-"`
}

func (m *OAuthAccessTokenIdentifiers) Reset()         { *m = OAuthAccessTokenIdentifiers{} }
func (m *OAuthAccessTokenIdentifiers) String() string { return proto.CompactTextString(m) }
func (*OAuthAccessTokenIdentifiers) ProtoMessage()    {}
func (*OAuthAccessTokenIdentifiers) Descriptor() ([]byte, []int) {
	return fileDescriptor_1454904971eaa7d7, []int{5}
}
func (m *OAuthAccessTokenIdentifiers) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_OAuthAccessTokenIdentifiers.Unmarshal(m, b)
}
func (m *OAuthAccessTokenIdentifiers) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_OAuthAccessTokenIdentifiers.Marshal(b, m, deterministic)
}
func (m *OAuthAccessTokenIdentifiers) XXX_Merge(src proto.Message) {
	xxx_messageInfo_OAuthAccessTokenIdentifiers.Merge(m, src)
}
func (m *OAuthAccessTokenIdentifiers) XXX_Size() int {
	return xxx_messageInfo_OAuthAccessTokenIdentifiers.Size(m)
}
func (m *OAuthAccessTokenIdentifiers) XXX_DiscardUnknown() {
	xxx_messageInfo_OAuthAccessTokenIdentifiers.DiscardUnknown(m)
}

var xxx_messageInfo_OAuthAccessTokenIdentifiers proto.InternalMessageInfo

func (m *OAuthAccessTokenIdentifiers) GetUserIds() *UserIdentifiers {
	if m != nil {
		return m.UserIds
	}
	return nil
}

func (m *OAuthAccessTokenIdentifiers) GetClientIds() *ClientIdentifiers {
	if m != nil {
		return m.ClientIds
	}
	return nil
}

func (m *OAuthAccessTokenIdentifiers) GetId() string {
	if m != nil {
		return m.Id
	}
	return ""
}

type OAuthAccessToken struct {
	UserIds              *UserIdentifiers   `protobuf:"bytes,1,opt,name=user_ids,json=userIds,proto3" json:"user_ids,omitempty"`
	UserSessionId        string             `protobuf:"bytes,9,opt,name=user_session_id,json=userSessionId,proto3" json:"user_session_id,omitempty"`
	ClientIds            *ClientIdentifiers `protobuf:"bytes,2,opt,name=client_ids,json=clientIds,proto3" json:"client_ids,omitempty"`
	Id                   string             `protobuf:"bytes,3,opt,name=id,proto3" json:"id,omitempty"`
	AccessToken          string             `protobuf:"bytes,4,opt,name=access_token,json=accessToken,proto3" json:"access_token,omitempty"`
	RefreshToken         string             `protobuf:"bytes,5,opt,name=refresh_token,json=refreshToken,proto3" json:"refresh_token,omitempty"`
	Rights               []Right            `protobuf:"varint,6,rep,packed,name=rights,proto3,enum=ttn.lorawan.v3.Right" json:"rights,omitempty"`
	CreatedAt            *types.Timestamp   `protobuf:"bytes,7,opt,name=created_at,json=createdAt,proto3" json:"created_at,omitempty"`
	ExpiresAt            *types.Timestamp   `protobuf:"bytes,8,opt,name=expires_at,json=expiresAt,proto3" json:"expires_at,omitempty"`
	XXX_NoUnkeyedLiteral struct{}           `json:"-"`
	XXX_sizecache        int32              `json:"-"`
}

func (m *OAuthAccessToken) Reset()         { *m = OAuthAccessToken{} }
func (m *OAuthAccessToken) String() string { return proto.CompactTextString(m) }
func (*OAuthAccessToken) ProtoMessage()    {}
func (*OAuthAccessToken) Descriptor() ([]byte, []int) {
	return fileDescriptor_1454904971eaa7d7, []int{6}
}
func (m *OAuthAccessToken) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_OAuthAccessToken.Unmarshal(m, b)
}
func (m *OAuthAccessToken) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_OAuthAccessToken.Marshal(b, m, deterministic)
}
func (m *OAuthAccessToken) XXX_Merge(src proto.Message) {
	xxx_messageInfo_OAuthAccessToken.Merge(m, src)
}
func (m *OAuthAccessToken) XXX_Size() int {
	return xxx_messageInfo_OAuthAccessToken.Size(m)
}
func (m *OAuthAccessToken) XXX_DiscardUnknown() {
	xxx_messageInfo_OAuthAccessToken.DiscardUnknown(m)
}

var xxx_messageInfo_OAuthAccessToken proto.InternalMessageInfo

func (m *OAuthAccessToken) GetUserIds() *UserIdentifiers {
	if m != nil {
		return m.UserIds
	}
	return nil
}

func (m *OAuthAccessToken) GetUserSessionId() string {
	if m != nil {
		return m.UserSessionId
	}
	return ""
}

func (m *OAuthAccessToken) GetClientIds() *ClientIdentifiers {
	if m != nil {
		return m.ClientIds
	}
	return nil
}

func (m *OAuthAccessToken) GetId() string {
	if m != nil {
		return m.Id
	}
	return ""
}

func (m *OAuthAccessToken) GetAccessToken() string {
	if m != nil {
		return m.AccessToken
	}
	return ""
}

func (m *OAuthAccessToken) GetRefreshToken() string {
	if m != nil {
		return m.RefreshToken
	}
	return ""
}

func (m *OAuthAccessToken) GetRights() []Right {
	if m != nil {
		return m.Rights
	}
	return nil
}

func (m *OAuthAccessToken) GetCreatedAt() *types.Timestamp {
	if m != nil {
		return m.CreatedAt
	}
	return nil
}

func (m *OAuthAccessToken) GetExpiresAt() *types.Timestamp {
	if m != nil {
		return m.ExpiresAt
	}
	return nil
}

type OAuthAccessTokens struct {
	Tokens               []*OAuthAccessToken `protobuf:"bytes,1,rep,name=tokens,proto3" json:"tokens,omitempty"`
	XXX_NoUnkeyedLiteral struct{}            `json:"-"`
	XXX_sizecache        int32               `json:"-"`
}

func (m *OAuthAccessTokens) Reset()         { *m = OAuthAccessTokens{} }
func (m *OAuthAccessTokens) String() string { return proto.CompactTextString(m) }
func (*OAuthAccessTokens) ProtoMessage()    {}
func (*OAuthAccessTokens) Descriptor() ([]byte, []int) {
	return fileDescriptor_1454904971eaa7d7, []int{7}
}
func (m *OAuthAccessTokens) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_OAuthAccessTokens.Unmarshal(m, b)
}
func (m *OAuthAccessTokens) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_OAuthAccessTokens.Marshal(b, m, deterministic)
}
func (m *OAuthAccessTokens) XXX_Merge(src proto.Message) {
	xxx_messageInfo_OAuthAccessTokens.Merge(m, src)
}
func (m *OAuthAccessTokens) XXX_Size() int {
	return xxx_messageInfo_OAuthAccessTokens.Size(m)
}
func (m *OAuthAccessTokens) XXX_DiscardUnknown() {
	xxx_messageInfo_OAuthAccessTokens.DiscardUnknown(m)
}

var xxx_messageInfo_OAuthAccessTokens proto.InternalMessageInfo

func (m *OAuthAccessTokens) GetTokens() []*OAuthAccessToken {
	if m != nil {
		return m.Tokens
	}
	return nil
}

type ListOAuthAccessTokensRequest struct {
	UserIds   *UserIdentifiers   `protobuf:"bytes,1,opt,name=user_ids,json=userIds,proto3" json:"user_ids,omitempty"`
	ClientIds *ClientIdentifiers `protobuf:"bytes,2,opt,name=client_ids,json=clientIds,proto3" json:"client_ids,omitempty"`
	// Order the results by this field path (must be present in the field mask).
	// Default ordering is by ID. Prepend with a minus (-) to reverse the order.
	Order string `protobuf:"bytes,3,opt,name=order,proto3" json:"order,omitempty"`
	// Limit the number of results per page.
	Limit uint32 `protobuf:"varint,4,opt,name=limit,proto3" json:"limit,omitempty"`
	// Page number for pagination. 0 is interpreted as 1.
	Page                 uint32   `protobuf:"varint,5,opt,name=page,proto3" json:"page,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *ListOAuthAccessTokensRequest) Reset()         { *m = ListOAuthAccessTokensRequest{} }
func (m *ListOAuthAccessTokensRequest) String() string { return proto.CompactTextString(m) }
func (*ListOAuthAccessTokensRequest) ProtoMessage()    {}
func (*ListOAuthAccessTokensRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_1454904971eaa7d7, []int{8}
}
func (m *ListOAuthAccessTokensRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_ListOAuthAccessTokensRequest.Unmarshal(m, b)
}
func (m *ListOAuthAccessTokensRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_ListOAuthAccessTokensRequest.Marshal(b, m, deterministic)
}
func (m *ListOAuthAccessTokensRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ListOAuthAccessTokensRequest.Merge(m, src)
}
func (m *ListOAuthAccessTokensRequest) XXX_Size() int {
	return xxx_messageInfo_ListOAuthAccessTokensRequest.Size(m)
}
func (m *ListOAuthAccessTokensRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_ListOAuthAccessTokensRequest.DiscardUnknown(m)
}

var xxx_messageInfo_ListOAuthAccessTokensRequest proto.InternalMessageInfo

func (m *ListOAuthAccessTokensRequest) GetUserIds() *UserIdentifiers {
	if m != nil {
		return m.UserIds
	}
	return nil
}

func (m *ListOAuthAccessTokensRequest) GetClientIds() *ClientIdentifiers {
	if m != nil {
		return m.ClientIds
	}
	return nil
}

func (m *ListOAuthAccessTokensRequest) GetOrder() string {
	if m != nil {
		return m.Order
	}
	return ""
}

func (m *ListOAuthAccessTokensRequest) GetLimit() uint32 {
	if m != nil {
		return m.Limit
	}
	return 0
}

func (m *ListOAuthAccessTokensRequest) GetPage() uint32 {
	if m != nil {
		return m.Page
	}
	return 0
}

func init() {
	proto.RegisterType((*OAuthClientAuthorizationIdentifiers)(nil), "ttn.lorawan.v3.OAuthClientAuthorizationIdentifiers")
	golang_proto.RegisterType((*OAuthClientAuthorizationIdentifiers)(nil), "ttn.lorawan.v3.OAuthClientAuthorizationIdentifiers")
	proto.RegisterType((*OAuthClientAuthorization)(nil), "ttn.lorawan.v3.OAuthClientAuthorization")
	golang_proto.RegisterType((*OAuthClientAuthorization)(nil), "ttn.lorawan.v3.OAuthClientAuthorization")
	proto.RegisterType((*OAuthClientAuthorizations)(nil), "ttn.lorawan.v3.OAuthClientAuthorizations")
	golang_proto.RegisterType((*OAuthClientAuthorizations)(nil), "ttn.lorawan.v3.OAuthClientAuthorizations")
	proto.RegisterType((*ListOAuthClientAuthorizationsRequest)(nil), "ttn.lorawan.v3.ListOAuthClientAuthorizationsRequest")
	golang_proto.RegisterType((*ListOAuthClientAuthorizationsRequest)(nil), "ttn.lorawan.v3.ListOAuthClientAuthorizationsRequest")
	proto.RegisterType((*OAuthAuthorizationCode)(nil), "ttn.lorawan.v3.OAuthAuthorizationCode")
	golang_proto.RegisterType((*OAuthAuthorizationCode)(nil), "ttn.lorawan.v3.OAuthAuthorizationCode")
	proto.RegisterType((*OAuthAccessTokenIdentifiers)(nil), "ttn.lorawan.v3.OAuthAccessTokenIdentifiers")
	golang_proto.RegisterType((*OAuthAccessTokenIdentifiers)(nil), "ttn.lorawan.v3.OAuthAccessTokenIdentifiers")
	proto.RegisterType((*OAuthAccessToken)(nil), "ttn.lorawan.v3.OAuthAccessToken")
	golang_proto.RegisterType((*OAuthAccessToken)(nil), "ttn.lorawan.v3.OAuthAccessToken")
	proto.RegisterType((*OAuthAccessTokens)(nil), "ttn.lorawan.v3.OAuthAccessTokens")
	golang_proto.RegisterType((*OAuthAccessTokens)(nil), "ttn.lorawan.v3.OAuthAccessTokens")
	proto.RegisterType((*ListOAuthAccessTokensRequest)(nil), "ttn.lorawan.v3.ListOAuthAccessTokensRequest")
	golang_proto.RegisterType((*ListOAuthAccessTokensRequest)(nil), "ttn.lorawan.v3.ListOAuthAccessTokensRequest")
}

func init() { proto.RegisterFile("lorawan-stack/api/oauth.proto", fileDescriptor_1454904971eaa7d7) }
func init() {
	golang_proto.RegisterFile("lorawan-stack/api/oauth.proto", fileDescriptor_1454904971eaa7d7)
}

var fileDescriptor_1454904971eaa7d7 = []byte{
	// 774 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xdc, 0x56, 0xcd, 0x6e, 0xdb, 0x46,
	0x10, 0x36, 0x45, 0xfd, 0x58, 0x2b, 0x4b, 0x75, 0x89, 0xb6, 0x60, 0xed, 0x5a, 0x96, 0xe9, 0x1e,
	0x84, 0x16, 0x22, 0x01, 0x1b, 0x28, 0xda, 0x5b, 0x25, 0xf7, 0xe2, 0xa2, 0x45, 0x8b, 0x8d, 0x7d,
	0xc9, 0x45, 0x58, 0x91, 0x6b, 0x6a, 0x21, 0x89, 0xcb, 0xec, 0x0e, 0x65, 0x3b, 0x8f, 0x90, 0x93,
	0x1f, 0x20, 0xef, 0x91, 0x63, 0xf2, 0x0a, 0xc9, 0x25, 0x39, 0xdb, 0x17, 0x1f, 0x73, 0xf6, 0x29,
	0xd0, 0x92, 0xb2, 0x24, 0xda, 0x8a, 0xe1, 0x40, 0x48, 0x8c, 0xdc, 0x76, 0x77, 0xbe, 0xf9, 0x76,
	0xbf, 0x99, 0x8f, 0x03, 0xa2, 0x8d, 0x3e, 0x17, 0xe4, 0x98, 0x04, 0x0d, 0x09, 0xc4, 0xed, 0x39,
	0x24, 0x64, 0x0e, 0x27, 0x11, 0x74, 0xed, 0x50, 0x70, 0xe0, 0x46, 0x05, 0x20, 0xb0, 0x13, 0x88,
	0x3d, 0xdc, 0x5d, 0x6b, 0xfa, 0x0c, 0xba, 0x51, 0xc7, 0x76, 0xf9, 0xc0, 0xa1, 0xc1, 0x90, 0x9f,
	0x86, 0x82, 0x9f, 0x9c, 0x3a, 0x0a, 0xec, 0x36, 0x7c, 0x1a, 0x34, 0x86, 0xa4, 0xcf, 0x3c, 0x02,
	0xd4, 0xb9, 0xb1, 0x88, 0x29, 0xd7, 0x1a, 0x53, 0x14, 0x3e, 0xf7, 0x79, 0x9c, 0xdc, 0x89, 0x8e,
	0xd4, 0x4e, 0x6d, 0xd4, 0x2a, 0x81, 0x6f, 0xfa, 0x9c, 0xfb, 0x7d, 0x3a, 0x41, 0x01, 0x1b, 0x50,
	0x09, 0x64, 0x10, 0x26, 0x80, 0xed, 0x9b, 0x0a, 0x98, 0x47, 0x03, 0x60, 0x47, 0x8c, 0x0a, 0x99,
	0x80, 0xaa, 0x37, 0x41, 0x82, 0xf9, 0x5d, 0x48, 0xe2, 0xd6, 0x0b, 0x0d, 0x6d, 0xff, 0xd7, 0x8c,
	0xa0, 0xbb, 0xd7, 0x67, 0x34, 0x80, 0xd1, 0x8a, 0x0b, 0xf6, 0x94, 0x00, 0xe3, 0xc1, 0xfe, 0x84,
	0xcd, 0xf8, 0x0b, 0x2d, 0x47, 0x92, 0x8a, 0x36, 0xf3, 0xa4, 0xa9, 0xd5, 0xb4, 0x7a, 0x69, 0x67,
	0xd3, 0x9e, 0x2d, 0x91, 0x7d, 0x28, 0xa9, 0x98, 0x4a, 0x69, 0x2d, 0x5f, 0xb5, 0x72, 0xcf, 0xb4,
	0xcc, 0xaa, 0x86, 0x0b, 0x91, 0x0a, 0x49, 0xe3, 0x6f, 0x84, 0x5c, 0x75, 0x8f, 0xe2, 0xc9, 0x28,
	0x9e, 0xad, 0x34, 0x4f, 0xfc, 0x92, 0xdb, 0x99, 0x8a, 0x6e, 0x12, 0x94, 0xd6, 0xeb, 0x0c, 0x32,
	0xe7, 0xbd, 0xfc, 0xe1, 0x3d, 0xd7, 0x68, 0xa0, 0x7c, 0x5c, 0x78, 0x53, 0xaf, 0xe9, 0xf5, 0xca,
	0xce, 0xf7, 0x69, 0x1e, 0x3c, 0x8a, 0xe2, 0x04, 0x64, 0xfc, 0x81, 0x90, 0x2b, 0x28, 0x01, 0xea,
	0xb5, 0x09, 0x98, 0x59, 0x75, 0xf5, 0x9a, 0x1d, 0x5b, 0xc2, 0x1e, 0x5b, 0xc2, 0x3e, 0x18, 0x5b,
	0x02, 0x17, 0x13, 0x74, 0x13, 0x46, 0xa9, 0x51, 0xe8, 0x8d, 0x53, 0x73, 0x77, 0xa7, 0x26, 0xe8,
	0x26, 0x58, 0x03, 0xf4, 0xe3, 0xbc, 0x92, 0x4a, 0xe3, 0x7f, 0x54, 0x21, 0x33, 0x27, 0xa6, 0x56,
	0xd3, 0xeb, 0xa5, 0x9d, 0x7a, 0x5a, 0xc9, 0x3c, 0x0a, 0x9c, 0xca, 0xb7, 0xde, 0x69, 0xe8, 0xe7,
	0x7f, 0x98, 0x84, 0xb9, 0x77, 0x62, 0xfa, 0x24, 0xa2, 0x12, 0x16, 0xd4, 0xce, 0xdf, 0x50, 0x8e,
	0x0b, 0x8f, 0x0a, 0xd5, 0xc9, 0x62, 0xab, 0x76, 0xd5, 0xda, 0x10, 0xeb, 0x78, 0x09, 0x4f, 0x55,
	0x1a, 0x97, 0x1a, 0x53, 0x9b, 0x18, 0x6e, 0x54, 0x51, 0xae, 0xcf, 0x06, 0x0c, 0x4c, 0xbd, 0xa6,
	0xd5, 0xcb, 0x8a, 0xf9, 0x17, 0xdd, 0xbc, 0x2c, 0xe0, 0xf8, 0xd8, 0x30, 0x50, 0x36, 0x24, 0x3e,
	0x55, 0x5d, 0x2a, 0x63, 0xb5, 0xb6, 0xde, 0xea, 0xe8, 0x07, 0x25, 0x6b, 0x46, 0xd0, 0x1e, 0xf7,
	0xe8, 0x82, 0xc4, 0x38, 0xe8, 0x1b, 0xc5, 0x22, 0xa9, 0x94, 0x8c, 0x07, 0x6d, 0xe6, 0x99, 0x45,
	0x25, 0xab, 0x70, 0xd5, 0xca, 0x8a, 0x8c, 0xf9, 0x27, 0x2e, 0x8f, 0xe2, 0x8f, 0xe2, 0xf0, 0xbe,
	0xf7, 0x25, 0xcd, 0x6c, 0xa0, 0xac, 0xcb, 0xbd, 0xb8, 0x40, 0x45, 0xac, 0xd6, 0xc6, 0xaf, 0x68,
	0x45, 0x50, 0x8f, 0x09, 0xea, 0x42, 0x3b, 0x12, 0x4c, 0xf9, 0xb4, 0xa8, 0x6e, 0x13, 0xfa, 0x99,
	0xa6, 0xe1, 0xd2, 0x38, 0x7a, 0x28, 0x98, 0xf1, 0x1d, 0xca, 0x49, 0x20, 0x40, 0xcd, 0xbc, 0x62,
	0x88, 0x37, 0xa9, 0x6f, 0xa4, 0x70, 0xcf, 0x6f, 0x84, 0x9e, 0x84, 0x4c, 0x50, 0x39, 0x4a, 0x5d,
	0xbe, 0x3b, 0x35, 0x41, 0x37, 0xc1, 0x7a, 0xa9, 0xa1, 0xf5, 0xb8, 0xb3, 0xae, 0x4b, 0xa5, 0x3c,
	0xe0, 0x3d, 0xfa, 0xb0, 0x27, 0xa5, 0x51, 0x41, 0x19, 0xe6, 0x29, 0xf3, 0x16, 0x71, 0x86, 0x79,
	0xd6, 0x1b, 0x1d, 0xad, 0xa6, 0x15, 0x7c, 0x0d, 0xae, 0x4c, 0xe9, 0x34, 0xb6, 0xd0, 0x0a, 0x51,
	0x0a, 0xdb, 0x30, 0x92, 0x98, 0xd8, 0xaf, 0x44, 0xa6, 0x54, 0x6f, 0xa3, 0xb2, 0xa0, 0x47, 0x82,
	0xca, 0x6e, 0x82, 0x51, 0x36, 0xc4, 0x2b, 0xc9, 0x61, 0x0c, 0x9a, 0xb8, 0x3d, 0x7f, 0xff, 0xd1,
	0xfd, 0xb9, 0x6c, 0xf9, 0x2f, 0xfa, 0x36, 0xdd, 0x53, 0x69, 0xfc, 0x8e, 0xf2, 0x4a, 0xd6, 0x78,
	0x54, 0xd7, 0x6e, 0x1d, 0xd5, 0x53, 0x29, 0x38, 0xc1, 0x5b, 0xcf, 0x33, 0xe8, 0xa7, 0xeb, 0xd1,
	0x3c, 0xcd, 0xb9, 0xd8, 0x91, 0xbc, 0xc8, 0xf6, 0x5f, 0x8f, 0x77, 0xfd, 0x13, 0xc7, 0x7b, 0xf6,
	0xe3, 0xe3, 0x3d, 0x37, 0x19, 0xef, 0xad, 0xe6, 0xe5, 0x79, 0x75, 0xe9, 0xfd, 0x79, 0x55, 0x3b,
	0xbb, 0xa8, 0x2e, 0xbd, 0xba, 0xa8, 0x6a, 0x8f, 0x1d, 0x9f, 0xdb, 0xd0, 0xa5, 0xd0, 0x65, 0x81,
	0x2f, 0xed, 0x80, 0xc2, 0x31, 0x17, 0x3d, 0x67, 0xf6, 0xef, 0x6b, 0xb8, 0xeb, 0x84, 0x3d, 0xdf,
	0x01, 0x08, 0xc2, 0x4e, 0x27, 0xaf, 0xfa, 0xb9, 0xfb, 0x21, 0x00, 0x00, 0xff, 0xff, 0x9e, 0x50,
	0xa5, 0xa8, 0x89, 0x0a, 0x00, 0x00,
}
