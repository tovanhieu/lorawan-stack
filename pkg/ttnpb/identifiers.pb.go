// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: lorawan-stack/api/identifiers.proto

package ttnpb

import (
	fmt "fmt"
	_ "github.com/envoyproxy/protoc-gen-validate/validate"
	_ "github.com/gogo/protobuf/gogoproto"
	proto "github.com/gogo/protobuf/proto"
	golang_proto "github.com/golang/protobuf/proto"
	go_thethings_network_lorawan_stack_v3_pkg_types "go.thethings.network/lorawan-stack/v3/pkg/types"
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

type ApplicationIdentifiers struct {
	ApplicationId        string   `protobuf:"bytes,1,opt,name=application_id,json=applicationId,proto3" json:"application_id,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *ApplicationIdentifiers) Reset()         { *m = ApplicationIdentifiers{} }
func (m *ApplicationIdentifiers) String() string { return proto.CompactTextString(m) }
func (*ApplicationIdentifiers) ProtoMessage()    {}
func (*ApplicationIdentifiers) Descriptor() ([]byte, []int) {
	return fileDescriptor_6da1fbfdea4d7423, []int{0}
}
func (m *ApplicationIdentifiers) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_ApplicationIdentifiers.Unmarshal(m, b)
}
func (m *ApplicationIdentifiers) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_ApplicationIdentifiers.Marshal(b, m, deterministic)
}
func (m *ApplicationIdentifiers) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ApplicationIdentifiers.Merge(m, src)
}
func (m *ApplicationIdentifiers) XXX_Size() int {
	return xxx_messageInfo_ApplicationIdentifiers.Size(m)
}
func (m *ApplicationIdentifiers) XXX_DiscardUnknown() {
	xxx_messageInfo_ApplicationIdentifiers.DiscardUnknown(m)
}

var xxx_messageInfo_ApplicationIdentifiers proto.InternalMessageInfo

func (m *ApplicationIdentifiers) GetApplicationId() string {
	if m != nil {
		return m.ApplicationId
	}
	return ""
}

type ClientIdentifiers struct {
	ClientId             string   `protobuf:"bytes,1,opt,name=client_id,json=clientId,proto3" json:"client_id,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *ClientIdentifiers) Reset()         { *m = ClientIdentifiers{} }
func (m *ClientIdentifiers) String() string { return proto.CompactTextString(m) }
func (*ClientIdentifiers) ProtoMessage()    {}
func (*ClientIdentifiers) Descriptor() ([]byte, []int) {
	return fileDescriptor_6da1fbfdea4d7423, []int{1}
}
func (m *ClientIdentifiers) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_ClientIdentifiers.Unmarshal(m, b)
}
func (m *ClientIdentifiers) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_ClientIdentifiers.Marshal(b, m, deterministic)
}
func (m *ClientIdentifiers) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ClientIdentifiers.Merge(m, src)
}
func (m *ClientIdentifiers) XXX_Size() int {
	return xxx_messageInfo_ClientIdentifiers.Size(m)
}
func (m *ClientIdentifiers) XXX_DiscardUnknown() {
	xxx_messageInfo_ClientIdentifiers.DiscardUnknown(m)
}

var xxx_messageInfo_ClientIdentifiers proto.InternalMessageInfo

func (m *ClientIdentifiers) GetClientId() string {
	if m != nil {
		return m.ClientId
	}
	return ""
}

type EndDeviceIdentifiers struct {
	DeviceId       string                  `protobuf:"bytes,1,opt,name=device_id,json=deviceId,proto3" json:"device_id,omitempty"`
	ApplicationIds *ApplicationIdentifiers `protobuf:"bytes,2,opt,name=application_ids,json=applicationIds,proto3" json:"application_ids,omitempty"`
	// The LoRaWAN DevEUI.
	DevEui *go_thethings_network_lorawan_stack_v3_pkg_types.EUI64 `protobuf:"bytes,4,opt,name=dev_eui,json=devEui,proto3,customtype=go.thethings.network/lorawan-stack/v3/pkg/types.EUI64" json:"dev_eui,omitempty"`
	// The LoRaWAN JoinEUI (AppEUI until LoRaWAN 1.0.3 end devices).
	JoinEui *go_thethings_network_lorawan_stack_v3_pkg_types.EUI64 `protobuf:"bytes,5,opt,name=join_eui,json=joinEui,proto3,customtype=go.thethings.network/lorawan-stack/v3/pkg/types.EUI64" json:"join_eui,omitempty"`
	// The LoRaWAN DevAddr.
	DevAddr              *go_thethings_network_lorawan_stack_v3_pkg_types.DevAddr `protobuf:"bytes,6,opt,name=dev_addr,json=devAddr,proto3,customtype=go.thethings.network/lorawan-stack/v3/pkg/types.DevAddr" json:"dev_addr,omitempty"`
	XXX_NoUnkeyedLiteral struct{}                                                 `json:"-"`
	XXX_unrecognized     []byte                                                   `json:"-"`
	XXX_sizecache        int32                                                    `json:"-"`
}

func (m *EndDeviceIdentifiers) Reset()         { *m = EndDeviceIdentifiers{} }
func (m *EndDeviceIdentifiers) String() string { return proto.CompactTextString(m) }
func (*EndDeviceIdentifiers) ProtoMessage()    {}
func (*EndDeviceIdentifiers) Descriptor() ([]byte, []int) {
	return fileDescriptor_6da1fbfdea4d7423, []int{2}
}
func (m *EndDeviceIdentifiers) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_EndDeviceIdentifiers.Unmarshal(m, b)
}
func (m *EndDeviceIdentifiers) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_EndDeviceIdentifiers.Marshal(b, m, deterministic)
}
func (m *EndDeviceIdentifiers) XXX_Merge(src proto.Message) {
	xxx_messageInfo_EndDeviceIdentifiers.Merge(m, src)
}
func (m *EndDeviceIdentifiers) XXX_Size() int {
	return xxx_messageInfo_EndDeviceIdentifiers.Size(m)
}
func (m *EndDeviceIdentifiers) XXX_DiscardUnknown() {
	xxx_messageInfo_EndDeviceIdentifiers.DiscardUnknown(m)
}

var xxx_messageInfo_EndDeviceIdentifiers proto.InternalMessageInfo

func (m *EndDeviceIdentifiers) GetDeviceId() string {
	if m != nil {
		return m.DeviceId
	}
	return ""
}

func (m *EndDeviceIdentifiers) GetApplicationIds() *ApplicationIdentifiers {
	if m != nil {
		return m.ApplicationIds
	}
	return nil
}

type GatewayIdentifiers struct {
	GatewayId string `protobuf:"bytes,1,opt,name=gateway_id,json=gatewayId,proto3" json:"gateway_id,omitempty"`
	// Secondary identifier, which can only be used in specific requests.
	Eui                  *go_thethings_network_lorawan_stack_v3_pkg_types.EUI64 `protobuf:"bytes,2,opt,name=eui,proto3,customtype=go.thethings.network/lorawan-stack/v3/pkg/types.EUI64" json:"eui,omitempty"`
	XXX_NoUnkeyedLiteral struct{}                                               `json:"-"`
	XXX_unrecognized     []byte                                                 `json:"-"`
	XXX_sizecache        int32                                                  `json:"-"`
}

func (m *GatewayIdentifiers) Reset()         { *m = GatewayIdentifiers{} }
func (m *GatewayIdentifiers) String() string { return proto.CompactTextString(m) }
func (*GatewayIdentifiers) ProtoMessage()    {}
func (*GatewayIdentifiers) Descriptor() ([]byte, []int) {
	return fileDescriptor_6da1fbfdea4d7423, []int{3}
}
func (m *GatewayIdentifiers) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_GatewayIdentifiers.Unmarshal(m, b)
}
func (m *GatewayIdentifiers) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_GatewayIdentifiers.Marshal(b, m, deterministic)
}
func (m *GatewayIdentifiers) XXX_Merge(src proto.Message) {
	xxx_messageInfo_GatewayIdentifiers.Merge(m, src)
}
func (m *GatewayIdentifiers) XXX_Size() int {
	return xxx_messageInfo_GatewayIdentifiers.Size(m)
}
func (m *GatewayIdentifiers) XXX_DiscardUnknown() {
	xxx_messageInfo_GatewayIdentifiers.DiscardUnknown(m)
}

var xxx_messageInfo_GatewayIdentifiers proto.InternalMessageInfo

func (m *GatewayIdentifiers) GetGatewayId() string {
	if m != nil {
		return m.GatewayId
	}
	return ""
}

type OrganizationIdentifiers struct {
	// This ID shares namespace with user IDs.
	OrganizationId       string   `protobuf:"bytes,1,opt,name=organization_id,json=organizationId,proto3" json:"organization_id,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *OrganizationIdentifiers) Reset()         { *m = OrganizationIdentifiers{} }
func (m *OrganizationIdentifiers) String() string { return proto.CompactTextString(m) }
func (*OrganizationIdentifiers) ProtoMessage()    {}
func (*OrganizationIdentifiers) Descriptor() ([]byte, []int) {
	return fileDescriptor_6da1fbfdea4d7423, []int{4}
}
func (m *OrganizationIdentifiers) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_OrganizationIdentifiers.Unmarshal(m, b)
}
func (m *OrganizationIdentifiers) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_OrganizationIdentifiers.Marshal(b, m, deterministic)
}
func (m *OrganizationIdentifiers) XXX_Merge(src proto.Message) {
	xxx_messageInfo_OrganizationIdentifiers.Merge(m, src)
}
func (m *OrganizationIdentifiers) XXX_Size() int {
	return xxx_messageInfo_OrganizationIdentifiers.Size(m)
}
func (m *OrganizationIdentifiers) XXX_DiscardUnknown() {
	xxx_messageInfo_OrganizationIdentifiers.DiscardUnknown(m)
}

var xxx_messageInfo_OrganizationIdentifiers proto.InternalMessageInfo

func (m *OrganizationIdentifiers) GetOrganizationId() string {
	if m != nil {
		return m.OrganizationId
	}
	return ""
}

type UserIdentifiers struct {
	// This ID shares namespace with organization IDs.
	UserId string `protobuf:"bytes,1,opt,name=user_id,json=userId,proto3" json:"user_id,omitempty"`
	// Secondary identifier, which can only be used in specific requests.
	Email                string   `protobuf:"bytes,2,opt,name=email,proto3" json:"email,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *UserIdentifiers) Reset()         { *m = UserIdentifiers{} }
func (m *UserIdentifiers) String() string { return proto.CompactTextString(m) }
func (*UserIdentifiers) ProtoMessage()    {}
func (*UserIdentifiers) Descriptor() ([]byte, []int) {
	return fileDescriptor_6da1fbfdea4d7423, []int{5}
}
func (m *UserIdentifiers) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_UserIdentifiers.Unmarshal(m, b)
}
func (m *UserIdentifiers) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_UserIdentifiers.Marshal(b, m, deterministic)
}
func (m *UserIdentifiers) XXX_Merge(src proto.Message) {
	xxx_messageInfo_UserIdentifiers.Merge(m, src)
}
func (m *UserIdentifiers) XXX_Size() int {
	return xxx_messageInfo_UserIdentifiers.Size(m)
}
func (m *UserIdentifiers) XXX_DiscardUnknown() {
	xxx_messageInfo_UserIdentifiers.DiscardUnknown(m)
}

var xxx_messageInfo_UserIdentifiers proto.InternalMessageInfo

func (m *UserIdentifiers) GetUserId() string {
	if m != nil {
		return m.UserId
	}
	return ""
}

func (m *UserIdentifiers) GetEmail() string {
	if m != nil {
		return m.Email
	}
	return ""
}

// OrganizationOrUserIdentifiers contains either organization or user identifiers.
type OrganizationOrUserIdentifiers struct {
	// Types that are valid to be assigned to Ids:
	//	*OrganizationOrUserIdentifiers_OrganizationIds
	//	*OrganizationOrUserIdentifiers_UserIds
	Ids                  isOrganizationOrUserIdentifiers_Ids `protobuf_oneof:"ids"`
	XXX_NoUnkeyedLiteral struct{}                            `json:"-"`
	XXX_unrecognized     []byte                              `json:"-"`
	XXX_sizecache        int32                               `json:"-"`
}

func (m *OrganizationOrUserIdentifiers) Reset()         { *m = OrganizationOrUserIdentifiers{} }
func (m *OrganizationOrUserIdentifiers) String() string { return proto.CompactTextString(m) }
func (*OrganizationOrUserIdentifiers) ProtoMessage()    {}
func (*OrganizationOrUserIdentifiers) Descriptor() ([]byte, []int) {
	return fileDescriptor_6da1fbfdea4d7423, []int{6}
}
func (m *OrganizationOrUserIdentifiers) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_OrganizationOrUserIdentifiers.Unmarshal(m, b)
}
func (m *OrganizationOrUserIdentifiers) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_OrganizationOrUserIdentifiers.Marshal(b, m, deterministic)
}
func (m *OrganizationOrUserIdentifiers) XXX_Merge(src proto.Message) {
	xxx_messageInfo_OrganizationOrUserIdentifiers.Merge(m, src)
}
func (m *OrganizationOrUserIdentifiers) XXX_Size() int {
	return xxx_messageInfo_OrganizationOrUserIdentifiers.Size(m)
}
func (m *OrganizationOrUserIdentifiers) XXX_DiscardUnknown() {
	xxx_messageInfo_OrganizationOrUserIdentifiers.DiscardUnknown(m)
}

var xxx_messageInfo_OrganizationOrUserIdentifiers proto.InternalMessageInfo

type isOrganizationOrUserIdentifiers_Ids interface {
	isOrganizationOrUserIdentifiers_Ids()
}

type OrganizationOrUserIdentifiers_OrganizationIds struct {
	OrganizationIds *OrganizationIdentifiers `protobuf:"bytes,1,opt,name=organization_ids,json=organizationIds,proto3,oneof" json:"organization_ids,omitempty"`
}
type OrganizationOrUserIdentifiers_UserIds struct {
	UserIds *UserIdentifiers `protobuf:"bytes,2,opt,name=user_ids,json=userIds,proto3,oneof" json:"user_ids,omitempty"`
}

func (*OrganizationOrUserIdentifiers_OrganizationIds) isOrganizationOrUserIdentifiers_Ids() {}
func (*OrganizationOrUserIdentifiers_UserIds) isOrganizationOrUserIdentifiers_Ids()         {}

func (m *OrganizationOrUserIdentifiers) GetIds() isOrganizationOrUserIdentifiers_Ids {
	if m != nil {
		return m.Ids
	}
	return nil
}

func (m *OrganizationOrUserIdentifiers) GetOrganizationIds() *OrganizationIdentifiers {
	if x, ok := m.GetIds().(*OrganizationOrUserIdentifiers_OrganizationIds); ok {
		return x.OrganizationIds
	}
	return nil
}

func (m *OrganizationOrUserIdentifiers) GetUserIds() *UserIdentifiers {
	if x, ok := m.GetIds().(*OrganizationOrUserIdentifiers_UserIds); ok {
		return x.UserIds
	}
	return nil
}

// XXX_OneofWrappers is for the internal use of the proto package.
func (*OrganizationOrUserIdentifiers) XXX_OneofWrappers() []interface{} {
	return []interface{}{
		(*OrganizationOrUserIdentifiers_OrganizationIds)(nil),
		(*OrganizationOrUserIdentifiers_UserIds)(nil),
	}
}

// EntityIdentifiers contains one of the possible entity identifiers.
type EntityIdentifiers struct {
	// Types that are valid to be assigned to Ids:
	//	*EntityIdentifiers_ApplicationIds
	//	*EntityIdentifiers_ClientIds
	//	*EntityIdentifiers_DeviceIds
	//	*EntityIdentifiers_GatewayIds
	//	*EntityIdentifiers_OrganizationIds
	//	*EntityIdentifiers_UserIds
	Ids                  isEntityIdentifiers_Ids `protobuf_oneof:"ids"`
	XXX_NoUnkeyedLiteral struct{}                `json:"-"`
	XXX_unrecognized     []byte                  `json:"-"`
	XXX_sizecache        int32                   `json:"-"`
}

func (m *EntityIdentifiers) Reset()         { *m = EntityIdentifiers{} }
func (m *EntityIdentifiers) String() string { return proto.CompactTextString(m) }
func (*EntityIdentifiers) ProtoMessage()    {}
func (*EntityIdentifiers) Descriptor() ([]byte, []int) {
	return fileDescriptor_6da1fbfdea4d7423, []int{7}
}
func (m *EntityIdentifiers) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_EntityIdentifiers.Unmarshal(m, b)
}
func (m *EntityIdentifiers) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_EntityIdentifiers.Marshal(b, m, deterministic)
}
func (m *EntityIdentifiers) XXX_Merge(src proto.Message) {
	xxx_messageInfo_EntityIdentifiers.Merge(m, src)
}
func (m *EntityIdentifiers) XXX_Size() int {
	return xxx_messageInfo_EntityIdentifiers.Size(m)
}
func (m *EntityIdentifiers) XXX_DiscardUnknown() {
	xxx_messageInfo_EntityIdentifiers.DiscardUnknown(m)
}

var xxx_messageInfo_EntityIdentifiers proto.InternalMessageInfo

type isEntityIdentifiers_Ids interface {
	isEntityIdentifiers_Ids()
}

type EntityIdentifiers_ApplicationIds struct {
	ApplicationIds *ApplicationIdentifiers `protobuf:"bytes,1,opt,name=application_ids,json=applicationIds,proto3,oneof" json:"application_ids,omitempty"`
}
type EntityIdentifiers_ClientIds struct {
	ClientIds *ClientIdentifiers `protobuf:"bytes,2,opt,name=client_ids,json=clientIds,proto3,oneof" json:"client_ids,omitempty"`
}
type EntityIdentifiers_DeviceIds struct {
	DeviceIds *EndDeviceIdentifiers `protobuf:"bytes,3,opt,name=device_ids,json=deviceIds,proto3,oneof" json:"device_ids,omitempty"`
}
type EntityIdentifiers_GatewayIds struct {
	GatewayIds *GatewayIdentifiers `protobuf:"bytes,4,opt,name=gateway_ids,json=gatewayIds,proto3,oneof" json:"gateway_ids,omitempty"`
}
type EntityIdentifiers_OrganizationIds struct {
	OrganizationIds *OrganizationIdentifiers `protobuf:"bytes,5,opt,name=organization_ids,json=organizationIds,proto3,oneof" json:"organization_ids,omitempty"`
}
type EntityIdentifiers_UserIds struct {
	UserIds *UserIdentifiers `protobuf:"bytes,6,opt,name=user_ids,json=userIds,proto3,oneof" json:"user_ids,omitempty"`
}

func (*EntityIdentifiers_ApplicationIds) isEntityIdentifiers_Ids()  {}
func (*EntityIdentifiers_ClientIds) isEntityIdentifiers_Ids()       {}
func (*EntityIdentifiers_DeviceIds) isEntityIdentifiers_Ids()       {}
func (*EntityIdentifiers_GatewayIds) isEntityIdentifiers_Ids()      {}
func (*EntityIdentifiers_OrganizationIds) isEntityIdentifiers_Ids() {}
func (*EntityIdentifiers_UserIds) isEntityIdentifiers_Ids()         {}

func (m *EntityIdentifiers) GetIds() isEntityIdentifiers_Ids {
	if m != nil {
		return m.Ids
	}
	return nil
}

func (m *EntityIdentifiers) GetApplicationIds() *ApplicationIdentifiers {
	if x, ok := m.GetIds().(*EntityIdentifiers_ApplicationIds); ok {
		return x.ApplicationIds
	}
	return nil
}

func (m *EntityIdentifiers) GetClientIds() *ClientIdentifiers {
	if x, ok := m.GetIds().(*EntityIdentifiers_ClientIds); ok {
		return x.ClientIds
	}
	return nil
}

func (m *EntityIdentifiers) GetDeviceIds() *EndDeviceIdentifiers {
	if x, ok := m.GetIds().(*EntityIdentifiers_DeviceIds); ok {
		return x.DeviceIds
	}
	return nil
}

func (m *EntityIdentifiers) GetGatewayIds() *GatewayIdentifiers {
	if x, ok := m.GetIds().(*EntityIdentifiers_GatewayIds); ok {
		return x.GatewayIds
	}
	return nil
}

func (m *EntityIdentifiers) GetOrganizationIds() *OrganizationIdentifiers {
	if x, ok := m.GetIds().(*EntityIdentifiers_OrganizationIds); ok {
		return x.OrganizationIds
	}
	return nil
}

func (m *EntityIdentifiers) GetUserIds() *UserIdentifiers {
	if x, ok := m.GetIds().(*EntityIdentifiers_UserIds); ok {
		return x.UserIds
	}
	return nil
}

// XXX_OneofWrappers is for the internal use of the proto package.
func (*EntityIdentifiers) XXX_OneofWrappers() []interface{} {
	return []interface{}{
		(*EntityIdentifiers_ApplicationIds)(nil),
		(*EntityIdentifiers_ClientIds)(nil),
		(*EntityIdentifiers_DeviceIds)(nil),
		(*EntityIdentifiers_GatewayIds)(nil),
		(*EntityIdentifiers_OrganizationIds)(nil),
		(*EntityIdentifiers_UserIds)(nil),
	}
}

// Identifies an end device model with version information.
type EndDeviceVersionIdentifiers struct {
	BrandId              string   `protobuf:"bytes,1,opt,name=brand_id,json=brandId,proto3" json:"brand_id,omitempty"`
	ModelId              string   `protobuf:"bytes,2,opt,name=model_id,json=modelId,proto3" json:"model_id,omitempty"`
	HardwareVersion      string   `protobuf:"bytes,3,opt,name=hardware_version,json=hardwareVersion,proto3" json:"hardware_version,omitempty"`
	FirmwareVersion      string   `protobuf:"bytes,4,opt,name=firmware_version,json=firmwareVersion,proto3" json:"firmware_version,omitempty"`
	BandId               string   `protobuf:"bytes,5,opt,name=band_id,json=bandId,proto3" json:"band_id,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *EndDeviceVersionIdentifiers) Reset()         { *m = EndDeviceVersionIdentifiers{} }
func (m *EndDeviceVersionIdentifiers) String() string { return proto.CompactTextString(m) }
func (*EndDeviceVersionIdentifiers) ProtoMessage()    {}
func (*EndDeviceVersionIdentifiers) Descriptor() ([]byte, []int) {
	return fileDescriptor_6da1fbfdea4d7423, []int{8}
}
func (m *EndDeviceVersionIdentifiers) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_EndDeviceVersionIdentifiers.Unmarshal(m, b)
}
func (m *EndDeviceVersionIdentifiers) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_EndDeviceVersionIdentifiers.Marshal(b, m, deterministic)
}
func (m *EndDeviceVersionIdentifiers) XXX_Merge(src proto.Message) {
	xxx_messageInfo_EndDeviceVersionIdentifiers.Merge(m, src)
}
func (m *EndDeviceVersionIdentifiers) XXX_Size() int {
	return xxx_messageInfo_EndDeviceVersionIdentifiers.Size(m)
}
func (m *EndDeviceVersionIdentifiers) XXX_DiscardUnknown() {
	xxx_messageInfo_EndDeviceVersionIdentifiers.DiscardUnknown(m)
}

var xxx_messageInfo_EndDeviceVersionIdentifiers proto.InternalMessageInfo

func (m *EndDeviceVersionIdentifiers) GetBrandId() string {
	if m != nil {
		return m.BrandId
	}
	return ""
}

func (m *EndDeviceVersionIdentifiers) GetModelId() string {
	if m != nil {
		return m.ModelId
	}
	return ""
}

func (m *EndDeviceVersionIdentifiers) GetHardwareVersion() string {
	if m != nil {
		return m.HardwareVersion
	}
	return ""
}

func (m *EndDeviceVersionIdentifiers) GetFirmwareVersion() string {
	if m != nil {
		return m.FirmwareVersion
	}
	return ""
}

func (m *EndDeviceVersionIdentifiers) GetBandId() string {
	if m != nil {
		return m.BandId
	}
	return ""
}

// Identifies a Network Server.
type NetworkIdentifiers struct {
	// LoRa Alliance NetID.
	NetId *go_thethings_network_lorawan_stack_v3_pkg_types.NetID `protobuf:"bytes,1,opt,name=net_id,json=netId,proto3,customtype=go.thethings.network/lorawan-stack/v3/pkg/types.NetID" json:"net_id,omitempty"`
	// Optional tenant identifier for multi-tenant deployments.
	TenantId string `protobuf:"bytes,2,opt,name=tenant_id,json=tenantId,proto3" json:"tenant_id,omitempty"`
	// Cluster identifier of the Network Server.
	ClusterId string `protobuf:"bytes,3,opt,name=cluster_id,json=clusterId,proto3" json:"cluster_id,omitempty"`
	// Cluster address of the Network Server.
	ClusterAddress string `protobuf:"bytes,4,opt,name=cluster_address,json=clusterAddress,proto3" json:"cluster_address,omitempty"`
	// Optional tenant address for multi-tenant deployments.
	TenantAddress        string   `protobuf:"bytes,5,opt,name=tenant_address,json=tenantAddress,proto3" json:"tenant_address,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *NetworkIdentifiers) Reset()         { *m = NetworkIdentifiers{} }
func (m *NetworkIdentifiers) String() string { return proto.CompactTextString(m) }
func (*NetworkIdentifiers) ProtoMessage()    {}
func (*NetworkIdentifiers) Descriptor() ([]byte, []int) {
	return fileDescriptor_6da1fbfdea4d7423, []int{9}
}
func (m *NetworkIdentifiers) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_NetworkIdentifiers.Unmarshal(m, b)
}
func (m *NetworkIdentifiers) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_NetworkIdentifiers.Marshal(b, m, deterministic)
}
func (m *NetworkIdentifiers) XXX_Merge(src proto.Message) {
	xxx_messageInfo_NetworkIdentifiers.Merge(m, src)
}
func (m *NetworkIdentifiers) XXX_Size() int {
	return xxx_messageInfo_NetworkIdentifiers.Size(m)
}
func (m *NetworkIdentifiers) XXX_DiscardUnknown() {
	xxx_messageInfo_NetworkIdentifiers.DiscardUnknown(m)
}

var xxx_messageInfo_NetworkIdentifiers proto.InternalMessageInfo

func (m *NetworkIdentifiers) GetTenantId() string {
	if m != nil {
		return m.TenantId
	}
	return ""
}

func (m *NetworkIdentifiers) GetClusterId() string {
	if m != nil {
		return m.ClusterId
	}
	return ""
}

func (m *NetworkIdentifiers) GetClusterAddress() string {
	if m != nil {
		return m.ClusterAddress
	}
	return ""
}

func (m *NetworkIdentifiers) GetTenantAddress() string {
	if m != nil {
		return m.TenantAddress
	}
	return ""
}

func init() {
	proto.RegisterType((*ApplicationIdentifiers)(nil), "ttn.lorawan.v3.ApplicationIdentifiers")
	golang_proto.RegisterType((*ApplicationIdentifiers)(nil), "ttn.lorawan.v3.ApplicationIdentifiers")
	proto.RegisterType((*ClientIdentifiers)(nil), "ttn.lorawan.v3.ClientIdentifiers")
	golang_proto.RegisterType((*ClientIdentifiers)(nil), "ttn.lorawan.v3.ClientIdentifiers")
	proto.RegisterType((*EndDeviceIdentifiers)(nil), "ttn.lorawan.v3.EndDeviceIdentifiers")
	golang_proto.RegisterType((*EndDeviceIdentifiers)(nil), "ttn.lorawan.v3.EndDeviceIdentifiers")
	proto.RegisterType((*GatewayIdentifiers)(nil), "ttn.lorawan.v3.GatewayIdentifiers")
	golang_proto.RegisterType((*GatewayIdentifiers)(nil), "ttn.lorawan.v3.GatewayIdentifiers")
	proto.RegisterType((*OrganizationIdentifiers)(nil), "ttn.lorawan.v3.OrganizationIdentifiers")
	golang_proto.RegisterType((*OrganizationIdentifiers)(nil), "ttn.lorawan.v3.OrganizationIdentifiers")
	proto.RegisterType((*UserIdentifiers)(nil), "ttn.lorawan.v3.UserIdentifiers")
	golang_proto.RegisterType((*UserIdentifiers)(nil), "ttn.lorawan.v3.UserIdentifiers")
	proto.RegisterType((*OrganizationOrUserIdentifiers)(nil), "ttn.lorawan.v3.OrganizationOrUserIdentifiers")
	golang_proto.RegisterType((*OrganizationOrUserIdentifiers)(nil), "ttn.lorawan.v3.OrganizationOrUserIdentifiers")
	proto.RegisterType((*EntityIdentifiers)(nil), "ttn.lorawan.v3.EntityIdentifiers")
	golang_proto.RegisterType((*EntityIdentifiers)(nil), "ttn.lorawan.v3.EntityIdentifiers")
	proto.RegisterType((*EndDeviceVersionIdentifiers)(nil), "ttn.lorawan.v3.EndDeviceVersionIdentifiers")
	golang_proto.RegisterType((*EndDeviceVersionIdentifiers)(nil), "ttn.lorawan.v3.EndDeviceVersionIdentifiers")
	proto.RegisterType((*NetworkIdentifiers)(nil), "ttn.lorawan.v3.NetworkIdentifiers")
	golang_proto.RegisterType((*NetworkIdentifiers)(nil), "ttn.lorawan.v3.NetworkIdentifiers")
}

func init() {
	proto.RegisterFile("lorawan-stack/api/identifiers.proto", fileDescriptor_6da1fbfdea4d7423)
}
func init() {
	golang_proto.RegisterFile("lorawan-stack/api/identifiers.proto", fileDescriptor_6da1fbfdea4d7423)
}

var fileDescriptor_6da1fbfdea4d7423 = []byte{
	// 911 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xac, 0x96, 0xdf, 0x6e, 0x1b, 0x45,
	0x14, 0xc6, 0xbd, 0x76, 0xfc, 0xef, 0x84, 0xda, 0xe9, 0xaa, 0x02, 0xab, 0x88, 0xc6, 0x5d, 0x4c,
	0x13, 0x10, 0xf6, 0x52, 0x87, 0x82, 0x0a, 0x48, 0x6d, 0x96, 0x98, 0xda, 0x42, 0x4a, 0x61, 0xd5,
	0x56, 0x2a, 0x55, 0x1b, 0x8d, 0x3d, 0x93, 0xf5, 0x10, 0x67, 0xd6, 0x9a, 0x19, 0x6f, 0x48, 0x01,
	0x89, 0x6b, 0xde, 0x84, 0x07, 0xe0, 0x02, 0x71, 0x01, 0x8f, 0xc0, 0x25, 0xea, 0x45, 0x6e, 0x78,
	0x09, 0x94, 0x2b, 0x34, 0xb3, 0xbb, 0x8e, 0x77, 0x1d, 0x40, 0x76, 0x7a, 0xb7, 0xb3, 0x73, 0xce,
	0x6f, 0xe7, 0x7c, 0xc7, 0xdf, 0x19, 0xc3, 0x9b, 0x23, 0x9f, 0xa3, 0x23, 0xc4, 0x9a, 0x42, 0xa2,
	0xc1, 0x81, 0x8d, 0xc6, 0xd4, 0xa6, 0x98, 0x30, 0x49, 0xf7, 0x29, 0xe1, 0xa2, 0x35, 0xe6, 0xbe,
	0xf4, 0xcd, 0x8a, 0x94, 0xac, 0x15, 0x05, 0xb6, 0x82, 0xad, 0xab, 0xdb, 0x1e, 0x95, 0xc3, 0x49,
	0xbf, 0x35, 0xf0, 0x0f, 0x6d, 0xc2, 0x02, 0xff, 0x78, 0xcc, 0xfd, 0x6f, 0x8e, 0x6d, 0x1d, 0x3c,
	0x68, 0x7a, 0x84, 0x35, 0x03, 0x34, 0xa2, 0x18, 0x49, 0x62, 0xcf, 0x3d, 0x84, 0xc8, 0xab, 0xcd,
	0x19, 0x84, 0xe7, 0x7b, 0x7e, 0x98, 0xdc, 0x9f, 0xec, 0xeb, 0x95, 0x5e, 0xe8, 0xa7, 0x30, 0xdc,
	0x1a, 0xc2, 0xab, 0xdb, 0xe3, 0xf1, 0x88, 0x0e, 0x90, 0xa4, 0x3e, 0xeb, 0x9d, 0x9d, 0xd0, 0xdc,
	0x85, 0x0a, 0x3a, 0xdb, 0xd9, 0xa3, 0xb8, 0x66, 0xd4, 0x8d, 0xcd, 0xb2, 0xb3, 0x71, 0xea, 0x34,
	0xb8, 0x55, 0x6b, 0xb4, 0xaf, 0x3d, 0x7b, 0x82, 0x9a, 0xcf, 0xdf, 0x6b, 0xde, 0x7e, 0xba, 0x79,
	0xe7, 0xa3, 0x27, 0xcd, 0xa7, 0x77, 0xe2, 0xe5, 0xdb, 0xdf, 0xb6, 0xdf, 0xfd, 0xbe, 0xe1, 0x5e,
	0x42, 0xb3, 0x60, 0xeb, 0x31, 0x5c, 0xfe, 0x74, 0x44, 0x09, 0x93, 0xb3, 0x1f, 0xd9, 0x81, 0xf2,
	0x40, 0xbf, 0x5c, 0x82, 0x5f, 0x1a, 0x44, 0x38, 0xeb, 0x97, 0x1c, 0x5c, 0xe9, 0x30, 0xbc, 0x43,
	0x02, 0x3a, 0x20, 0x29, 0x3c, 0xd6, 0x2f, 0x97, 0xc1, 0xe3, 0x08, 0x67, 0x3e, 0x86, 0x6a, 0x52,
	0x09, 0x51, 0xcb, 0xd6, 0x8d, 0xcd, 0xd5, 0xf6, 0x8d, 0x56, 0xb2, 0x7f, 0xad, 0xf3, 0xa5, 0x74,
	0x4a, 0xa7, 0x4e, 0xfe, 0x47, 0x23, 0xbb, 0x66, 0xb8, 0x95, 0x84, 0x26, 0xc2, 0x74, 0xa1, 0x88,
	0x49, 0xb0, 0x47, 0x26, 0xb4, 0xb6, 0x52, 0x37, 0x36, 0x5f, 0x71, 0x6e, 0xbf, 0x38, 0x59, 0xbf,
	0xe5, 0xf9, 0x2d, 0x39, 0x24, 0x72, 0x48, 0x99, 0x27, 0x5a, 0x8c, 0xc8, 0x23, 0x9f, 0x1f, 0xd8,
	0xc9, 0xdf, 0x53, 0xb0, 0x65, 0x8f, 0x0f, 0x3c, 0x5b, 0x1e, 0x8f, 0x89, 0x68, 0x75, 0x1e, 0xf6,
	0x3e, 0x78, 0xdf, 0x2d, 0x60, 0x12, 0x74, 0x26, 0xd4, 0x7c, 0x00, 0xa5, 0xaf, 0x7d, 0xca, 0x34,
	0x34, 0x7f, 0x51, 0x68, 0x51, 0xa1, 0x14, 0xf5, 0x11, 0x28, 0x41, 0xf6, 0x10, 0xc6, 0xbc, 0x56,
	0xd0, 0xd4, 0x8f, 0x5f, 0x9c, 0xac, 0x7f, 0xb8, 0x28, 0x75, 0x87, 0x04, 0xdb, 0x18, 0x73, 0x57,
	0x95, 0xad, 0x1e, 0xac, 0x9f, 0x0c, 0x30, 0xef, 0x21, 0x49, 0x8e, 0xd0, 0xf1, 0x6c, 0xe7, 0x3e,
	0x03, 0xf0, 0xc2, 0xb7, 0x4b, 0xb4, 0xae, 0xec, 0xc5, 0x40, 0xf3, 0x73, 0xc8, 0x29, 0x1d, 0xb2,
	0x17, 0xd5, 0x41, 0x51, 0xac, 0x03, 0x78, 0xed, 0x3e, 0xf7, 0x10, 0xa3, 0xcf, 0xe7, 0xdc, 0xf2,
	0x05, 0x54, 0xfd, 0x99, 0xad, 0x25, 0x0e, 0x5d, 0xf1, 0x13, 0x68, 0x8b, 0x42, 0xf5, 0xa1, 0x20,
	0x7c, 0xf6, 0x23, 0x77, 0xa1, 0x38, 0x11, 0x84, 0x2f, 0x06, 0xbf, 0xa9, 0xe0, 0x85, 0x89, 0x46,
	0x99, 0x57, 0x20, 0x4f, 0x0e, 0x11, 0x1d, 0x69, 0x41, 0xca, 0x6e, 0xb8, 0xb0, 0x7e, 0x33, 0xe0,
	0x8d, 0xd9, 0xc2, 0xee, 0xf3, 0xf4, 0x97, 0x1f, 0xc0, 0x5a, 0xaa, 0x3c, 0xa1, 0x8f, 0xb0, 0xda,
	0xde, 0x48, 0x7b, 0xe0, 0x5f, 0x14, 0xea, 0x66, 0xdc, 0x6a, 0xb2, 0x42, 0x61, 0x7e, 0x02, 0xa5,
	0xa8, 0x9e, 0xd8, 0x51, 0xeb, 0x69, 0x5a, 0xea, 0x20, 0xdd, 0x8c, 0x5b, 0x0c, 0x4b, 0x11, 0x0e,
	0x40, 0x8e, 0x62, 0x61, 0xe6, 0xfe, 0x76, 0x0c, 0xeb, 0xcf, 0x1c, 0x5c, 0xee, 0x30, 0x49, 0x65,
	0xe2, 0x47, 0xf4, 0xe5, 0xbc, 0x71, 0x8d, 0x45, 0x8c, 0xdb, 0xcd, 0xcc, 0x19, 0xd6, 0x01, 0x98,
	0x0e, 0xac, 0xf8, 0xd0, 0xd7, 0xd3, 0xb4, 0xb9, 0x39, 0xd7, 0xcd, 0xb8, 0xe5, 0x78, 0x5a, 0x09,
	0xb3, 0x03, 0x30, 0x9d, 0x4a, 0xa2, 0x96, 0xd3, 0x8c, 0x46, 0x9a, 0x71, 0xde, 0x3c, 0x53, 0x98,
	0x78, 0x2a, 0x29, 0xcc, 0xea, 0x99, 0x45, 0x84, 0x9e, 0x1f, 0xab, 0x6d, 0x2b, 0xcd, 0x99, 0xf7,
	0x56, 0x37, 0xe3, 0xc2, 0xd4, 0x20, 0xe7, 0xb7, 0x36, 0xff, 0x52, 0x5b, 0x5b, 0xb8, 0x50, 0x6b,
	0x7f, 0xce, 0xc2, 0xeb, 0x53, 0x31, 0x1e, 0x11, 0x2e, 0x52, 0xce, 0xeb, 0x40, 0xa9, 0xcf, 0x11,
	0xc3, 0x67, 0xae, 0x78, 0xe7, 0xd4, 0xd9, 0xe0, 0x6f, 0xfd, 0xbf, 0xe5, 0xfe, 0x30, 0x0c, 0xb7,
	0xa8, 0x73, 0x7b, 0x58, 0x61, 0x0e, 0x7d, 0x4c, 0x46, 0x0a, 0x93, 0x5d, 0x1c, 0xa3, 0x73, 0x7b,
	0xd8, 0x6c, 0xc3, 0xda, 0x10, 0x71, 0x7c, 0x84, 0x38, 0xd9, 0x0b, 0xc2, 0xc3, 0xea, 0x0e, 0x97,
	0x9d, 0xe2, 0xa9, 0xb3, 0xc2, 0xb3, 0xb5, 0xba, 0x5b, 0x8d, 0x03, 0xa2, 0x62, 0x54, 0xce, 0x3e,
	0xe5, 0x87, 0x89, 0x9c, 0x95, 0x54, 0x4e, 0x1c, 0x10, 0xe7, 0xd4, 0xa1, 0xd8, 0x8f, 0x8a, 0xce,
	0x27, 0x43, 0x0b, 0x7d, 0x5d, 0x90, 0xf5, 0x6b, 0x16, 0xcc, 0xdd, 0x70, 0xbc, 0x25, 0x07, 0x55,
	0x81, 0x91, 0xe9, 0x75, 0xbb, 0xe4, 0x4c, 0xdc, 0x25, 0xb2, 0xb7, 0xe3, 0xe6, 0x19, 0x91, 0x3d,
	0x6c, 0xde, 0x83, 0xb2, 0x24, 0x0c, 0x85, 0x77, 0x78, 0x4a, 0xba, 0xeb, 0xff, 0x2d, 0xdd, 0x77,
	0xcf, 0x1a, 0x6e, 0x29, 0x4c, 0xee, 0x61, 0xf3, 0x86, 0xf2, 0xd6, 0x44, 0xc8, 0x70, 0xc2, 0x25,
	0x54, 0xbb, 0xab, 0xfc, 0xa3, 0xb7, 0x7a, 0xd8, 0xbc, 0x09, 0xd5, 0x38, 0x4e, 0x5d, 0x47, 0x44,
	0x88, 0x48, 0x2e, 0x75, 0xcf, 0xf2, 0x5c, 0xed, 0x87, 0xac, 0x5b, 0x89, 0x02, 0xb6, 0xc3, 0x7d,
	0xd3, 0x86, 0x4a, 0x74, 0xc6, 0x38, 0x23, 0x9f, 0xca, 0xb8, 0x14, 0xee, 0x47, 0x09, 0xce, 0xad,
	0xdf, 0xff, 0xba, 0x66, 0x7c, 0x65, 0x2f, 0x20, 0x8c, 0x64, 0xe3, 0x7e, 0xbf, 0xa0, 0xff, 0x55,
	0x6d, 0xfd, 0x13, 0x00, 0x00, 0xff, 0xff, 0x89, 0xe6, 0x23, 0x22, 0xfe, 0x09, 0x00, 0x00,
}
