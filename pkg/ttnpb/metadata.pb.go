// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: lorawan-stack/api/metadata.proto

package ttnpb

import (
	fmt "fmt"
	_ "github.com/TheThingsIndustries/protoc-gen-go-json/annotations"
	_ "github.com/envoyproxy/protoc-gen-validate/validate"
	_ "github.com/gogo/protobuf/gogoproto"
	proto "github.com/gogo/protobuf/proto"
	types "github.com/gogo/protobuf/types"
	golang_proto "github.com/golang/protobuf/proto"
	go_thethings_network_lorawan_stack_v3_pkg_types "go.thethings.network/lorawan-stack/v3/pkg/types"
	math "math"
	strconv "strconv"
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

type LocationSource int32

const (
	// The source of the location is not known or not set.
	SOURCE_UNKNOWN LocationSource = 0
	// The location is determined by GPS.
	SOURCE_GPS LocationSource = 1
	// The location is set in and updated from a registry.
	SOURCE_REGISTRY LocationSource = 3
	// The location is estimated with IP geolocation.
	SOURCE_IP_GEOLOCATION LocationSource = 4
	// The location is estimated with WiFi RSSI geolocation.
	SOURCE_WIFI_RSSI_GEOLOCATION LocationSource = 5
	// The location is estimated with BT/BLE RSSI geolocation.
	SOURCE_BT_RSSI_GEOLOCATION LocationSource = 6
	// The location is estimated with LoRa RSSI geolocation.
	SOURCE_LORA_RSSI_GEOLOCATION LocationSource = 7
	// The location is estimated with LoRa TDOA geolocation.
	SOURCE_LORA_TDOA_GEOLOCATION LocationSource = 8
	// The location is estimated by a combination of geolocation sources.
	SOURCE_COMBINED_GEOLOCATION LocationSource = 9
)

var LocationSource_name = map[int32]string{
	0: "SOURCE_UNKNOWN",
	1: "SOURCE_GPS",
	3: "SOURCE_REGISTRY",
	4: "SOURCE_IP_GEOLOCATION",
	5: "SOURCE_WIFI_RSSI_GEOLOCATION",
	6: "SOURCE_BT_RSSI_GEOLOCATION",
	7: "SOURCE_LORA_RSSI_GEOLOCATION",
	8: "SOURCE_LORA_TDOA_GEOLOCATION",
	9: "SOURCE_COMBINED_GEOLOCATION",
}

var LocationSource_value = map[string]int32{
	"SOURCE_UNKNOWN":               0,
	"SOURCE_GPS":                   1,
	"SOURCE_REGISTRY":              3,
	"SOURCE_IP_GEOLOCATION":        4,
	"SOURCE_WIFI_RSSI_GEOLOCATION": 5,
	"SOURCE_BT_RSSI_GEOLOCATION":   6,
	"SOURCE_LORA_RSSI_GEOLOCATION": 7,
	"SOURCE_LORA_TDOA_GEOLOCATION": 8,
	"SOURCE_COMBINED_GEOLOCATION":  9,
}

func (LocationSource) EnumDescriptor() ([]byte, []int) {
	return fileDescriptor_e1123b3e8fd87092, []int{0}
}

// Contains metadata for a received message. Each antenna that receives
// a message corresponds to one RxMetadata.
type RxMetadata struct {
	GatewayIds   *GatewayIdentifiers   `protobuf:"bytes,1,opt,name=gateway_ids,json=gatewayIds,proto3" json:"gateway_ids,omitempty"`
	PacketBroker *PacketBrokerMetadata `protobuf:"bytes,18,opt,name=packet_broker,json=packetBroker,proto3" json:"packet_broker,omitempty"`
	AntennaIndex uint32                `protobuf:"varint,2,opt,name=antenna_index,json=antennaIndex,proto3" json:"antenna_index,omitempty"`
	Time         *types.Timestamp      `protobuf:"bytes,3,opt,name=time,proto3" json:"time,omitempty"`
	// Gateway concentrator timestamp when the Rx finished (microseconds).
	Timestamp uint32 `protobuf:"varint,4,opt,name=timestamp,proto3" json:"timestamp,omitempty"`
	// Gateway's internal fine timestamp when the Rx finished (nanoseconds).
	FineTimestamp uint64 `protobuf:"varint,5,opt,name=fine_timestamp,json=fineTimestamp,proto3" json:"fine_timestamp,omitempty"`
	// Encrypted gateway's internal fine timestamp when the Rx finished (nanoseconds).
	EncryptedFineTimestamp      []byte `protobuf:"bytes,6,opt,name=encrypted_fine_timestamp,json=encryptedFineTimestamp,proto3" json:"encrypted_fine_timestamp,omitempty"`
	EncryptedFineTimestampKeyId string `protobuf:"bytes,7,opt,name=encrypted_fine_timestamp_key_id,json=encryptedFineTimestampKeyId,proto3" json:"encrypted_fine_timestamp_key_id,omitempty"`
	// Received signal strength indicator (dBm).
	// This value equals `channel_rssi`.
	Rssi float32 `protobuf:"fixed32,8,opt,name=rssi,proto3" json:"rssi,omitempty"`
	// Received signal strength indicator of the signal (dBm).
	SignalRssi *types.FloatValue `protobuf:"bytes,16,opt,name=signal_rssi,json=signalRssi,proto3" json:"signal_rssi,omitempty"`
	// Received signal strength indicator of the channel (dBm).
	ChannelRssi float32 `protobuf:"fixed32,9,opt,name=channel_rssi,json=channelRssi,proto3" json:"channel_rssi,omitempty"`
	// Standard deviation of the RSSI during preamble.
	RssiStandardDeviation float32 `protobuf:"fixed32,10,opt,name=rssi_standard_deviation,json=rssiStandardDeviation,proto3" json:"rssi_standard_deviation,omitempty"`
	// Signal-to-noise ratio (dB).
	Snr float32 `protobuf:"fixed32,11,opt,name=snr,proto3" json:"snr,omitempty"`
	// Frequency offset (Hz).
	FrequencyOffset int64 `protobuf:"varint,12,opt,name=frequency_offset,json=frequencyOffset,proto3" json:"frequency_offset,omitempty"`
	// Antenna location; injected by the Gateway Server.
	Location *Location `protobuf:"bytes,13,opt,name=location,proto3" json:"location,omitempty"`
	// Gateway downlink path constraint; injected by the Gateway Server.
	DownlinkPathConstraint DownlinkPathConstraint `protobuf:"varint,14,opt,name=downlink_path_constraint,json=downlinkPathConstraint,proto3,enum=ttn.lorawan.v3.DownlinkPathConstraint" json:"downlink_path_constraint,omitempty"`
	// Uplink token to be included in the Tx request in class A downlink; injected by gateway, Gateway Server or fNS.
	UplinkToken []byte `protobuf:"bytes,15,opt,name=uplink_token,json=uplinkToken,proto3" json:"uplink_token,omitempty"`
	// Index of the gateway channel that received the message.
	ChannelIndex uint32 `protobuf:"varint,17,opt,name=channel_index,json=channelIndex,proto3" json:"channel_index,omitempty"`
	// Hopping width; a number describing the number of steps of the LR-FHSS grid.
	HoppingWidth uint32 `protobuf:"varint,19,opt,name=hopping_width,json=hoppingWidth,proto3" json:"hopping_width,omitempty"`
	// Frequency drift in Hz between start and end of an LR-FHSS packet (signed).
	FrequencyDrift int32 `protobuf:"varint,20,opt,name=frequency_drift,json=frequencyDrift,proto3" json:"frequency_drift,omitempty"`
	// Advanced metadata fields
	// - can be used for advanced information or experimental features that are not yet formally defined in the API
	// - field names are written in snake_case
	Advanced             *types.Struct `protobuf:"bytes,99,opt,name=advanced,proto3" json:"advanced,omitempty"`
	XXX_NoUnkeyedLiteral struct{}      `json:"-"`
	XXX_sizecache        int32         `json:"-"`
}

func (m *RxMetadata) Reset()         { *m = RxMetadata{} }
func (m *RxMetadata) String() string { return proto.CompactTextString(m) }
func (*RxMetadata) ProtoMessage()    {}
func (*RxMetadata) Descriptor() ([]byte, []int) {
	return fileDescriptor_e1123b3e8fd87092, []int{0}
}
func (m *RxMetadata) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_RxMetadata.Unmarshal(m, b)
}
func (m *RxMetadata) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_RxMetadata.Marshal(b, m, deterministic)
}
func (m *RxMetadata) XXX_Merge(src proto.Message) {
	xxx_messageInfo_RxMetadata.Merge(m, src)
}
func (m *RxMetadata) XXX_Size() int {
	return xxx_messageInfo_RxMetadata.Size(m)
}
func (m *RxMetadata) XXX_DiscardUnknown() {
	xxx_messageInfo_RxMetadata.DiscardUnknown(m)
}

var xxx_messageInfo_RxMetadata proto.InternalMessageInfo

func (m *RxMetadata) GetGatewayIds() *GatewayIdentifiers {
	if m != nil {
		return m.GatewayIds
	}
	return nil
}

func (m *RxMetadata) GetPacketBroker() *PacketBrokerMetadata {
	if m != nil {
		return m.PacketBroker
	}
	return nil
}

func (m *RxMetadata) GetAntennaIndex() uint32 {
	if m != nil {
		return m.AntennaIndex
	}
	return 0
}

func (m *RxMetadata) GetTime() *types.Timestamp {
	if m != nil {
		return m.Time
	}
	return nil
}

func (m *RxMetadata) GetTimestamp() uint32 {
	if m != nil {
		return m.Timestamp
	}
	return 0
}

func (m *RxMetadata) GetFineTimestamp() uint64 {
	if m != nil {
		return m.FineTimestamp
	}
	return 0
}

func (m *RxMetadata) GetEncryptedFineTimestamp() []byte {
	if m != nil {
		return m.EncryptedFineTimestamp
	}
	return nil
}

func (m *RxMetadata) GetEncryptedFineTimestampKeyId() string {
	if m != nil {
		return m.EncryptedFineTimestampKeyId
	}
	return ""
}

func (m *RxMetadata) GetRssi() float32 {
	if m != nil {
		return m.Rssi
	}
	return 0
}

func (m *RxMetadata) GetSignalRssi() *types.FloatValue {
	if m != nil {
		return m.SignalRssi
	}
	return nil
}

func (m *RxMetadata) GetChannelRssi() float32 {
	if m != nil {
		return m.ChannelRssi
	}
	return 0
}

func (m *RxMetadata) GetRssiStandardDeviation() float32 {
	if m != nil {
		return m.RssiStandardDeviation
	}
	return 0
}

func (m *RxMetadata) GetSnr() float32 {
	if m != nil {
		return m.Snr
	}
	return 0
}

func (m *RxMetadata) GetFrequencyOffset() int64 {
	if m != nil {
		return m.FrequencyOffset
	}
	return 0
}

func (m *RxMetadata) GetLocation() *Location {
	if m != nil {
		return m.Location
	}
	return nil
}

func (m *RxMetadata) GetDownlinkPathConstraint() DownlinkPathConstraint {
	if m != nil {
		return m.DownlinkPathConstraint
	}
	return DOWNLINK_PATH_CONSTRAINT_NONE
}

func (m *RxMetadata) GetUplinkToken() []byte {
	if m != nil {
		return m.UplinkToken
	}
	return nil
}

func (m *RxMetadata) GetChannelIndex() uint32 {
	if m != nil {
		return m.ChannelIndex
	}
	return 0
}

func (m *RxMetadata) GetHoppingWidth() uint32 {
	if m != nil {
		return m.HoppingWidth
	}
	return 0
}

func (m *RxMetadata) GetFrequencyDrift() int32 {
	if m != nil {
		return m.FrequencyDrift
	}
	return 0
}

func (m *RxMetadata) GetAdvanced() *types.Struct {
	if m != nil {
		return m.Advanced
	}
	return nil
}

type Location struct {
	// The North–South position (degrees; -90 to +90), where 0 is the equator, North pole is positive, South pole is negative.
	Latitude float64 `protobuf:"fixed64,1,opt,name=latitude,proto3" json:"latitude,omitempty"`
	// The East-West position (degrees; -180 to +180), where 0 is the Prime Meridian (Greenwich), East is positive , West is negative.
	Longitude float64 `protobuf:"fixed64,2,opt,name=longitude,proto3" json:"longitude,omitempty"`
	// The altitude (meters), where 0 is the mean sea level.
	Altitude int32 `protobuf:"varint,3,opt,name=altitude,proto3" json:"altitude,omitempty"`
	// The accuracy of the location (meters).
	Accuracy int32 `protobuf:"varint,4,opt,name=accuracy,proto3" json:"accuracy,omitempty"`
	// Source of the location information.
	Source               LocationSource `protobuf:"varint,5,opt,name=source,proto3,enum=ttn.lorawan.v3.LocationSource" json:"source,omitempty"`
	XXX_NoUnkeyedLiteral struct{}       `json:"-"`
	XXX_sizecache        int32          `json:"-"`
}

func (m *Location) Reset()         { *m = Location{} }
func (m *Location) String() string { return proto.CompactTextString(m) }
func (*Location) ProtoMessage()    {}
func (*Location) Descriptor() ([]byte, []int) {
	return fileDescriptor_e1123b3e8fd87092, []int{1}
}
func (m *Location) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Location.Unmarshal(m, b)
}
func (m *Location) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Location.Marshal(b, m, deterministic)
}
func (m *Location) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Location.Merge(m, src)
}
func (m *Location) XXX_Size() int {
	return xxx_messageInfo_Location.Size(m)
}
func (m *Location) XXX_DiscardUnknown() {
	xxx_messageInfo_Location.DiscardUnknown(m)
}

var xxx_messageInfo_Location proto.InternalMessageInfo

func (m *Location) GetLatitude() float64 {
	if m != nil {
		return m.Latitude
	}
	return 0
}

func (m *Location) GetLongitude() float64 {
	if m != nil {
		return m.Longitude
	}
	return 0
}

func (m *Location) GetAltitude() int32 {
	if m != nil {
		return m.Altitude
	}
	return 0
}

func (m *Location) GetAccuracy() int32 {
	if m != nil {
		return m.Accuracy
	}
	return 0
}

func (m *Location) GetSource() LocationSource {
	if m != nil {
		return m.Source
	}
	return SOURCE_UNKNOWN
}

type PacketBrokerMetadata struct {
	// Message identifier generated by Packet Broker Router.
	MessageId string `protobuf:"bytes,1,opt,name=message_id,json=messageId,proto3" json:"message_id,omitempty"`
	// LoRa Alliance NetID of the Packet Broker Forwarder Member.
	ForwarderNetId go_thethings_network_lorawan_stack_v3_pkg_types.NetID `protobuf:"bytes,2,opt,name=forwarder_net_id,json=forwarderNetId,proto3,customtype=go.thethings.network/lorawan-stack/v3/pkg/types.NetID" json:"forwarder_net_id"`
	// Tenant ID managed by the Packet Broker Forwarder Member.
	ForwarderTenantId string `protobuf:"bytes,3,opt,name=forwarder_tenant_id,json=forwarderTenantId,proto3" json:"forwarder_tenant_id,omitempty"`
	// Forwarder Cluster ID of the Packet Broker Forwarder.
	ForwarderClusterId string `protobuf:"bytes,4,opt,name=forwarder_cluster_id,json=forwarderClusterId,proto3" json:"forwarder_cluster_id,omitempty"`
	// Forwarder gateway EUI.
	ForwarderGatewayEui *go_thethings_network_lorawan_stack_v3_pkg_types.EUI64 `protobuf:"bytes,9,opt,name=forwarder_gateway_eui,json=forwarderGatewayEui,proto3,customtype=go.thethings.network/lorawan-stack/v3/pkg/types.EUI64" json:"forwarder_gateway_eui,omitempty"`
	// Forwarder gateway ID.
	ForwarderGatewayId *types.StringValue `protobuf:"bytes,10,opt,name=forwarder_gateway_id,json=forwarderGatewayId,proto3" json:"forwarder_gateway_id,omitempty"`
	// LoRa Alliance NetID of the Packet Broker Home Network Member.
	HomeNetworkNetId go_thethings_network_lorawan_stack_v3_pkg_types.NetID `protobuf:"bytes,5,opt,name=home_network_net_id,json=homeNetworkNetId,proto3,customtype=go.thethings.network/lorawan-stack/v3/pkg/types.NetID" json:"home_network_net_id"`
	// Tenant ID managed by the Packet Broker Home Network Member.
	// This value is empty if it cannot be determined by the Packet Broker Router.
	HomeNetworkTenantId string `protobuf:"bytes,6,opt,name=home_network_tenant_id,json=homeNetworkTenantId,proto3" json:"home_network_tenant_id,omitempty"`
	// Home Network Cluster ID of the Packet Broker Home Network.
	HomeNetworkClusterId string `protobuf:"bytes,8,opt,name=home_network_cluster_id,json=homeNetworkClusterId,proto3" json:"home_network_cluster_id,omitempty"`
	// Hops that the message passed. Each Packet Broker Router service appends an entry.
	Hops                 []*PacketBrokerRouteHop `protobuf:"bytes,7,rep,name=hops,proto3" json:"hops,omitempty"`
	XXX_NoUnkeyedLiteral struct{}                `json:"-"`
	XXX_sizecache        int32                   `json:"-"`
}

func (m *PacketBrokerMetadata) Reset()         { *m = PacketBrokerMetadata{} }
func (m *PacketBrokerMetadata) String() string { return proto.CompactTextString(m) }
func (*PacketBrokerMetadata) ProtoMessage()    {}
func (*PacketBrokerMetadata) Descriptor() ([]byte, []int) {
	return fileDescriptor_e1123b3e8fd87092, []int{2}
}
func (m *PacketBrokerMetadata) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_PacketBrokerMetadata.Unmarshal(m, b)
}
func (m *PacketBrokerMetadata) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_PacketBrokerMetadata.Marshal(b, m, deterministic)
}
func (m *PacketBrokerMetadata) XXX_Merge(src proto.Message) {
	xxx_messageInfo_PacketBrokerMetadata.Merge(m, src)
}
func (m *PacketBrokerMetadata) XXX_Size() int {
	return xxx_messageInfo_PacketBrokerMetadata.Size(m)
}
func (m *PacketBrokerMetadata) XXX_DiscardUnknown() {
	xxx_messageInfo_PacketBrokerMetadata.DiscardUnknown(m)
}

var xxx_messageInfo_PacketBrokerMetadata proto.InternalMessageInfo

func (m *PacketBrokerMetadata) GetMessageId() string {
	if m != nil {
		return m.MessageId
	}
	return ""
}

func (m *PacketBrokerMetadata) GetForwarderTenantId() string {
	if m != nil {
		return m.ForwarderTenantId
	}
	return ""
}

func (m *PacketBrokerMetadata) GetForwarderClusterId() string {
	if m != nil {
		return m.ForwarderClusterId
	}
	return ""
}

func (m *PacketBrokerMetadata) GetForwarderGatewayId() *types.StringValue {
	if m != nil {
		return m.ForwarderGatewayId
	}
	return nil
}

func (m *PacketBrokerMetadata) GetHomeNetworkTenantId() string {
	if m != nil {
		return m.HomeNetworkTenantId
	}
	return ""
}

func (m *PacketBrokerMetadata) GetHomeNetworkClusterId() string {
	if m != nil {
		return m.HomeNetworkClusterId
	}
	return ""
}

func (m *PacketBrokerMetadata) GetHops() []*PacketBrokerRouteHop {
	if m != nil {
		return m.Hops
	}
	return nil
}

type PacketBrokerRouteHop struct {
	// Time when the service received the message.
	ReceivedAt *types.Timestamp `protobuf:"bytes,1,opt,name=received_at,json=receivedAt,proto3" json:"received_at,omitempty"`
	// Sender of the message, typically the authorized client identifier.
	SenderName string `protobuf:"bytes,2,opt,name=sender_name,json=senderName,proto3" json:"sender_name,omitempty"`
	// Sender IP address or host name.
	SenderAddress string `protobuf:"bytes,3,opt,name=sender_address,json=senderAddress,proto3" json:"sender_address,omitempty"`
	// Receiver of the message.
	ReceiverName string `protobuf:"bytes,4,opt,name=receiver_name,json=receiverName,proto3" json:"receiver_name,omitempty"`
	// Receiver agent.
	ReceiverAgent        string   `protobuf:"bytes,5,opt,name=receiver_agent,json=receiverAgent,proto3" json:"receiver_agent,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *PacketBrokerRouteHop) Reset()         { *m = PacketBrokerRouteHop{} }
func (m *PacketBrokerRouteHop) String() string { return proto.CompactTextString(m) }
func (*PacketBrokerRouteHop) ProtoMessage()    {}
func (*PacketBrokerRouteHop) Descriptor() ([]byte, []int) {
	return fileDescriptor_e1123b3e8fd87092, []int{3}
}
func (m *PacketBrokerRouteHop) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_PacketBrokerRouteHop.Unmarshal(m, b)
}
func (m *PacketBrokerRouteHop) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_PacketBrokerRouteHop.Marshal(b, m, deterministic)
}
func (m *PacketBrokerRouteHop) XXX_Merge(src proto.Message) {
	xxx_messageInfo_PacketBrokerRouteHop.Merge(m, src)
}
func (m *PacketBrokerRouteHop) XXX_Size() int {
	return xxx_messageInfo_PacketBrokerRouteHop.Size(m)
}
func (m *PacketBrokerRouteHop) XXX_DiscardUnknown() {
	xxx_messageInfo_PacketBrokerRouteHop.DiscardUnknown(m)
}

var xxx_messageInfo_PacketBrokerRouteHop proto.InternalMessageInfo

func (m *PacketBrokerRouteHop) GetReceivedAt() *types.Timestamp {
	if m != nil {
		return m.ReceivedAt
	}
	return nil
}

func (m *PacketBrokerRouteHop) GetSenderName() string {
	if m != nil {
		return m.SenderName
	}
	return ""
}

func (m *PacketBrokerRouteHop) GetSenderAddress() string {
	if m != nil {
		return m.SenderAddress
	}
	return ""
}

func (m *PacketBrokerRouteHop) GetReceiverName() string {
	if m != nil {
		return m.ReceiverName
	}
	return ""
}

func (m *PacketBrokerRouteHop) GetReceiverAgent() string {
	if m != nil {
		return m.ReceiverAgent
	}
	return ""
}

func init() {
	proto.RegisterEnum("ttn.lorawan.v3.LocationSource", LocationSource_name, LocationSource_value)
	golang_proto.RegisterEnum("ttn.lorawan.v3.LocationSource", LocationSource_name, LocationSource_value)
	proto.RegisterType((*RxMetadata)(nil), "ttn.lorawan.v3.RxMetadata")
	golang_proto.RegisterType((*RxMetadata)(nil), "ttn.lorawan.v3.RxMetadata")
	proto.RegisterType((*Location)(nil), "ttn.lorawan.v3.Location")
	golang_proto.RegisterType((*Location)(nil), "ttn.lorawan.v3.Location")
	proto.RegisterType((*PacketBrokerMetadata)(nil), "ttn.lorawan.v3.PacketBrokerMetadata")
	golang_proto.RegisterType((*PacketBrokerMetadata)(nil), "ttn.lorawan.v3.PacketBrokerMetadata")
	proto.RegisterType((*PacketBrokerRouteHop)(nil), "ttn.lorawan.v3.PacketBrokerRouteHop")
	golang_proto.RegisterType((*PacketBrokerRouteHop)(nil), "ttn.lorawan.v3.PacketBrokerRouteHop")
}

func init() { proto.RegisterFile("lorawan-stack/api/metadata.proto", fileDescriptor_e1123b3e8fd87092) }
func init() {
	golang_proto.RegisterFile("lorawan-stack/api/metadata.proto", fileDescriptor_e1123b3e8fd87092)
}

var fileDescriptor_e1123b3e8fd87092 = []byte{
	// 1387 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xac, 0x55, 0x4d, 0x6f, 0xdb, 0x36,
	0x18, 0x8e, 0xe2, 0x24, 0xb5, 0xe9, 0x8f, 0xba, 0x4c, 0xda, 0xa8, 0x69, 0xea, 0x78, 0xe9, 0x3e,
	0xdc, 0x02, 0x91, 0x87, 0xa4, 0x2d, 0x3a, 0x6c, 0x03, 0x6a, 0x27, 0x69, 0x26, 0xb4, 0xb5, 0x03,
	0xda, 0x6d, 0xb1, 0x5d, 0x04, 0x46, 0xa2, 0x65, 0xd5, 0x36, 0xa9, 0x49, 0x54, 0x5c, 0xdf, 0x86,
	0x9d, 0x86, 0x9d, 0xf6, 0x0f, 0x86, 0x61, 0xb7, 0xfd, 0x8a, 0x1e, 0x77, 0xdf, 0x6d, 0x87, 0x1e,
	0xda, 0xc3, 0x86, 0x9d, 0x7a, 0xee, 0x65, 0x83, 0x28, 0x4a, 0xfe, 0x4a, 0x37, 0x6c, 0x98, 0x4e,
	0xe4, 0xf3, 0x3c, 0xef, 0x43, 0xea, 0x7d, 0x5f, 0x92, 0xa0, 0xdc, 0x67, 0x1e, 0x1e, 0x62, 0xba,
	0xe3, 0x73, 0x6c, 0xf6, 0xaa, 0xd8, 0x75, 0xaa, 0x03, 0xc2, 0xb1, 0x85, 0x39, 0xd6, 0x5c, 0x8f,
	0x71, 0x06, 0x0b, 0x9c, 0x53, 0x4d, 0xaa, 0xb4, 0xd3, 0xbd, 0x8d, 0x9a, 0xed, 0xf0, 0x6e, 0x70,
	0xa2, 0x99, 0x6c, 0x50, 0x25, 0xf4, 0x94, 0x8d, 0x5c, 0x8f, 0x3d, 0x1b, 0x55, 0x85, 0xd8, 0xdc,
	0xb1, 0x09, 0xdd, 0x39, 0xc5, 0x7d, 0xc7, 0xc2, 0x9c, 0x54, 0xe7, 0x06, 0x91, 0xe5, 0xc6, 0xce,
	0x84, 0x85, 0xcd, 0x6c, 0x16, 0x05, 0x9f, 0x04, 0x1d, 0x31, 0x13, 0x13, 0x31, 0x92, 0xf2, 0xfd,
	0x09, 0x79, 0xbb, 0x4b, 0xda, 0x5d, 0x87, 0xda, 0xbe, 0x4e, 0xad, 0xc0, 0xe7, 0x9e, 0x43, 0xfc,
	0xc9, 0xa5, 0x6d, 0xb6, 0xf3, 0xd4, 0x67, 0xb4, 0x8a, 0x29, 0x65, 0x1c, 0x73, 0x87, 0x51, 0x5f,
	0x9a, 0x6c, 0xda, 0x8c, 0xd9, 0x7d, 0x32, 0x5e, 0xca, 0xe7, 0x5e, 0x60, 0x72, 0xc9, 0x6e, 0xcd,
	0xb2, 0xdc, 0x19, 0x10, 0x9f, 0xe3, 0x81, 0x2b, 0x05, 0xa5, 0x59, 0xc1, 0xd0, 0xc3, 0xae, 0x4b,
	0xbc, 0xd8, 0xfe, 0xea, 0x7c, 0x1e, 0x09, 0x0d, 0x06, 0x31, 0x7d, 0x6d, 0x9e, 0x76, 0x2c, 0x42,
	0xb9, 0xd3, 0x71, 0x12, 0x8f, 0xed, 0x1f, 0xd2, 0x00, 0xa0, 0x67, 0x0f, 0x65, 0xfa, 0xe1, 0x43,
	0x90, 0xb5, 0x31, 0x27, 0x43, 0x3c, 0x32, 0x1c, 0xcb, 0x57, 0x95, 0xb2, 0x52, 0xc9, 0xee, 0x6e,
	0x6b, 0xd3, 0xe5, 0xd0, 0x8e, 0x22, 0x89, 0x3e, 0x76, 0xab, 0xa7, 0xdf, 0xd4, 0x97, 0xbf, 0x55,
	0x16, 0x8b, 0x0a, 0x02, 0x76, 0xcc, 0xfa, 0x50, 0x07, 0x79, 0x17, 0x9b, 0x3d, 0xc2, 0x8d, 0x13,
	0x8f, 0xf5, 0x88, 0xa7, 0x42, 0x61, 0xf8, 0xee, 0xac, 0xe1, 0xb1, 0x10, 0xd5, 0x85, 0x26, 0xde,
	0x0b, 0xca, 0xb9, 0x13, 0x28, 0xbc, 0x06, 0xf2, 0x98, 0x72, 0x42, 0x29, 0x36, 0x1c, 0x6a, 0x91,
	0x67, 0xea, 0x62, 0x59, 0xa9, 0xe4, 0x51, 0x4e, 0x82, 0x7a, 0x88, 0x41, 0x0d, 0x2c, 0x85, 0x49,
	0x54, 0x53, 0x62, 0x99, 0x0d, 0x2d, 0x4a, 0xa0, 0x16, 0x27, 0x50, 0x6b, 0xc7, 0x19, 0x46, 0x42,
	0x07, 0x37, 0x41, 0x26, 0x49, 0xba, 0xba, 0x24, 0x0c, 0xc7, 0x00, 0x7c, 0x0f, 0x14, 0x3a, 0x0e,
	0x25, 0xc6, 0x58, 0xb2, 0x5c, 0x56, 0x2a, 0x4b, 0x28, 0x1f, 0xa2, 0x89, 0x15, 0xbc, 0x03, 0x54,
	0x42, 0x4d, 0x6f, 0xe4, 0x72, 0x62, 0x19, 0x33, 0x01, 0x2b, 0x65, 0xa5, 0x92, 0x43, 0x97, 0x12,
	0xfe, 0xde, 0x54, 0xe4, 0x01, 0xd8, 0x7a, 0x5b, 0xa4, 0xd1, 0x23, 0x61, 0x09, 0xd4, 0x73, 0x65,
	0xa5, 0x92, 0x41, 0x57, 0xce, 0x36, 0xb8, 0x4f, 0x46, 0xba, 0x05, 0x21, 0x58, 0xf2, 0x7c, 0xdf,
	0x51, 0xd3, 0x65, 0xa5, 0xb2, 0x88, 0xc4, 0x18, 0x7e, 0x02, 0xb2, 0xbe, 0x63, 0x53, 0xdc, 0x37,
	0x04, 0x55, 0x14, 0xf9, 0xb8, 0x32, 0x97, 0x8f, 0x7b, 0x7d, 0x86, 0xf9, 0x63, 0xdc, 0x0f, 0x08,
	0x02, 0x91, 0x1e, 0x85, 0xd1, 0xef, 0x80, 0x9c, 0xd9, 0xc5, 0x94, 0x12, 0x19, 0x9e, 0x11, 0xce,
	0x59, 0x89, 0x09, 0xc9, 0x6d, 0xb0, 0x1e, 0x52, 0x86, 0xcf, 0x31, 0xb5, 0xb0, 0x67, 0x19, 0x16,
	0x39, 0x75, 0x44, 0xf3, 0xab, 0x40, 0xa8, 0x2f, 0x86, 0x74, 0x4b, 0xb2, 0x07, 0x31, 0x09, 0x8b,
	0x20, 0xe5, 0x53, 0x4f, 0xcd, 0x0a, 0x4d, 0x38, 0x84, 0xd7, 0x41, 0xb1, 0xe3, 0x91, 0x2f, 0x03,
	0x42, 0xcd, 0x91, 0xc1, 0x3a, 0x1d, 0x9f, 0x70, 0x35, 0x57, 0x56, 0x2a, 0x29, 0x74, 0x3e, 0xc1,
	0x9b, 0x02, 0x86, 0x37, 0x41, 0xba, 0xcf, 0xcc, 0x68, 0x95, 0xbc, 0xf8, 0x25, 0x75, 0xb6, 0x93,
	0x1e, 0x48, 0x1e, 0x25, 0x4a, 0xf8, 0x14, 0xa8, 0x16, 0x1b, 0xd2, 0xbe, 0x43, 0x7b, 0x86, 0x8b,
	0x79, 0xd7, 0x30, 0x19, 0xf5, 0xb9, 0x87, 0x1d, 0xca, 0xd5, 0x42, 0x59, 0xa9, 0x14, 0x76, 0xdf,
	0x9f, 0x75, 0x39, 0x90, 0xfa, 0x63, 0xcc, 0xbb, 0xfb, 0x89, 0x5a, 0x34, 0xf9, 0xd7, 0xa2, 0xc9,
	0x2f, 0x59, 0x67, 0x2a, 0xc2, 0xcc, 0x05, 0xae, 0x58, 0x89, 0xb3, 0x1e, 0xa1, 0xea, 0x79, 0x51,
	0xff, 0x6c, 0x84, 0xb5, 0x43, 0x08, 0xee, 0x80, 0x7c, 0x9c, 0xdc, 0xa8, 0x91, 0x2f, 0x84, 0x7d,
	0x27, 0xbc, 0x6f, 0xa4, 0xd4, 0x3f, 0x15, 0x14, 0xe7, 0x3e, 0x6a, 0xe9, 0x6b, 0x20, 0xdf, 0x65,
	0xae, 0xeb, 0x50, 0xdb, 0x18, 0x3a, 0x16, 0xef, 0xaa, 0xab, 0x51, 0xdf, 0x4b, 0xf0, 0x49, 0x88,
	0xc1, 0x0f, 0xc0, 0x38, 0x57, 0x86, 0xe5, 0x39, 0x1d, 0xae, 0xae, 0x95, 0x95, 0xca, 0x32, 0x2a,
	0x24, 0xf0, 0x41, 0x88, 0xc2, 0x3d, 0x90, 0xc6, 0xd6, 0x29, 0xa6, 0x26, 0xb1, 0x54, 0x53, 0x64,
	0x70, 0x7d, 0xae, 0x29, 0x5a, 0xe2, 0x92, 0x42, 0x89, 0x70, 0xfb, 0xb5, 0x02, 0xd2, 0x71, 0x5e,
	0x43, 0x87, 0x3e, 0xe6, 0x0e, 0x0f, 0x2c, 0x22, 0xae, 0x07, 0xa5, 0xbe, 0xfe, 0xa6, 0xbe, 0x06,
	0xe1, 0xe5, 0x85, 0xf0, 0xfb, 0xea, 0xf1, 0xdd, 0xeb, 0x72, 0xf0, 0x1c, 0x25, 0x42, 0x78, 0x0b,
	0x64, 0xfa, 0x8c, 0xda, 0x51, 0xd4, 0xe2, 0x7c, 0x54, 0x27, 0x8e, 0xea, 0x3c, 0x47, 0x63, 0x25,
	0xdc, 0x00, 0x69, 0xdc, 0x97, 0x6b, 0xa5, 0xc4, 0xff, 0x24, 0x73, 0xc1, 0x99, 0x66, 0xe0, 0x61,
	0x73, 0x24, 0x4e, 0x6e, 0xc8, 0xc9, 0x39, 0xbc, 0x0b, 0x56, 0x7c, 0x16, 0x78, 0x26, 0x11, 0x07,
	0xb6, 0xb0, 0x5b, 0x7a, 0x5b, 0x97, 0xb4, 0x84, 0x6a, 0xa2, 0xae, 0x32, 0x6e, 0xfb, 0x97, 0x65,
	0xb0, 0x76, 0xd6, 0xa5, 0x04, 0xaf, 0x02, 0x30, 0x20, 0xbe, 0x8f, 0x6d, 0x12, 0x9e, 0x4e, 0x45,
	0x9c, 0xce, 0x8c, 0x44, 0x74, 0x0b, 0xda, 0xa0, 0xd8, 0x61, 0xde, 0x10, 0x7b, 0x16, 0xf1, 0x0c,
	0x4a, 0x78, 0x28, 0x0a, 0xff, 0x37, 0x57, 0xff, 0xf4, 0xe7, 0x17, 0x5b, 0x0b, 0xbf, 0xbe, 0xd8,
	0xba, 0x65, 0x33, 0x8d, 0x77, 0x09, 0x17, 0x0f, 0x8a, 0x46, 0x09, 0x1f, 0x32, 0xaf, 0x57, 0x9d,
	0xbe, 0xaa, 0x4f, 0xf7, 0xaa, 0x6e, 0xcf, 0xae, 0xf2, 0x91, 0x4b, 0x7c, 0xad, 0x41, 0xb8, 0x7e,
	0x80, 0x0a, 0x89, 0x6d, 0x38, 0xb7, 0xa0, 0x06, 0x56, 0xc7, 0x0b, 0x71, 0x42, 0x31, 0x15, 0x6b,
	0xa5, 0xc4, 0x86, 0x2e, 0x24, 0x54, 0x5b, 0x30, 0xba, 0x05, 0x3f, 0x04, 0x6b, 0x63, 0xbd, 0xd9,
	0x0f, 0x7c, 0x4e, 0xbc, 0x30, 0x60, 0x49, 0x04, 0xc0, 0x84, 0xdb, 0x8f, 0x28, 0xdd, 0x82, 0x03,
	0x70, 0x71, 0x1c, 0x11, 0x3f, 0x0a, 0x24, 0x88, 0x6e, 0x83, 0x5c, 0xfd, 0xa3, 0xff, 0xf2, 0x2f,
	0x87, 0x8f, 0xf4, 0xdb, 0x37, 0xd1, 0x78, 0xe7, 0xf2, 0x21, 0x39, 0x0c, 0x1c, 0xd8, 0x98, 0xdc,
	0xe0, 0xf8, 0x0d, 0x12, 0xb7, 0x49, 0x76, 0x77, 0xf3, 0xac, 0x2e, 0x75, 0xa8, 0x1d, 0xdd, 0x5d,
	0x70, 0xd6, 0x50, 0xb7, 0x60, 0x1f, 0xac, 0x76, 0xd9, 0x80, 0x18, 0x72, 0x57, 0x71, 0x31, 0x96,
	0xff, 0x8f, 0x62, 0x14, 0x43, 0xe7, 0x46, 0xa4, 0x8e, 0xca, 0xb1, 0x07, 0x2e, 0x4d, 0xad, 0x36,
	0xae, 0xc8, 0x8a, 0x48, 0xf0, 0xea, 0x44, 0x44, 0x52, 0x93, 0x5b, 0x60, 0x7d, 0x2a, 0x68, 0xa2,
	0x2c, 0x69, 0x11, 0xb5, 0x36, 0x11, 0x35, 0x2e, 0xcc, 0x1d, 0xb0, 0xd4, 0x65, 0xae, 0xaf, 0x9e,
	0x2b, 0xa7, 0xfe, 0xe9, 0x2d, 0x45, 0x2c, 0xe0, 0xe4, 0x33, 0xe6, 0x22, 0x11, 0xb1, 0xfd, 0x9b,
	0x32, 0xdd, 0xd5, 0x31, 0x0d, 0x3f, 0x06, 0x59, 0x8f, 0x98, 0xc4, 0x39, 0x25, 0x96, 0x81, 0xb9,
	0x7c, 0xf6, 0xff, 0xee, 0xf9, 0x04, 0xb1, 0xbc, 0xc6, 0xe1, 0x16, 0xc8, 0xfa, 0x84, 0x8a, 0x86,
	0xc7, 0x83, 0xe8, 0x78, 0x67, 0x10, 0x88, 0xa0, 0x06, 0x1e, 0x90, 0xf0, 0x1d, 0x95, 0x02, 0x6c,
	0x59, 0x1e, 0xf1, 0x7d, 0xd9, 0xa6, 0xf9, 0x08, 0xad, 0x45, 0x60, 0x78, 0xd3, 0x49, 0x57, 0xe9,
	0x14, 0xf5, 0x66, 0x2e, 0x06, 0x63, 0xaf, 0x44, 0x84, 0x6d, 0x42, 0xb9, 0xa8, 0x68, 0x06, 0x25,
	0xa1, 0xb5, 0x10, 0xbc, 0xf1, 0xfd, 0x22, 0x28, 0x4c, 0x1f, 0x72, 0x08, 0x41, 0xa1, 0xd5, 0x7c,
	0x84, 0xf6, 0x0f, 0x8d, 0x47, 0x8d, 0xfb, 0x8d, 0xe6, 0x93, 0x46, 0x71, 0x01, 0x16, 0x00, 0x90,
	0xd8, 0xd1, 0x71, 0xab, 0xa8, 0xc0, 0x55, 0x70, 0x5e, 0xce, 0xd1, 0xe1, 0x91, 0xde, 0x6a, 0xa3,
	0xcf, 0x8b, 0x29, 0x78, 0x19, 0x5c, 0x94, 0xa0, 0x7e, 0x6c, 0x1c, 0x1d, 0x36, 0x1f, 0x34, 0xf7,
	0x6b, 0x6d, 0xbd, 0xd9, 0x28, 0x2e, 0xc1, 0x32, 0xd8, 0x94, 0xd4, 0x13, 0xfd, 0x9e, 0x6e, 0xa0,
	0x56, 0x4b, 0x9f, 0x52, 0x2c, 0xc3, 0x12, 0xd8, 0x90, 0x8a, 0x7a, 0x7b, 0x9e, 0x5f, 0x99, 0x70,
	0x78, 0xd0, 0x44, 0xb5, 0x79, 0xc5, 0xb9, 0x59, 0x45, 0xfb, 0xa0, 0x59, 0x9b, 0x52, 0xa4, 0xe1,
	0x16, 0xb8, 0x22, 0x15, 0xfb, 0xcd, 0x87, 0x75, 0xbd, 0x71, 0x78, 0x30, 0x25, 0xc8, 0x6c, 0xc0,
	0x3f, 0x7e, 0xba, 0x0c, 0x54, 0xe5, 0xc6, 0x4a, 0x24, 0xfb, 0xe6, 0xc7, 0xd2, 0x42, 0xbd, 0xf6,
	0xfb, 0xcb, 0xd2, 0xc2, 0xeb, 0x97, 0x25, 0xe5, 0xbb, 0x57, 0xa5, 0x85, 0xe7, 0xaf, 0x4a, 0xca,
	0x17, 0xd5, 0x7f, 0x71, 0x18, 0x38, 0x75, 0x4f, 0x4e, 0x56, 0x44, 0x63, 0xec, 0xfd, 0x15, 0x00,
	0x00, 0xff, 0xff, 0xff, 0xae, 0x44, 0x32, 0xd0, 0x0b, 0x00, 0x00,
}

func (x LocationSource) String() string {
	s, ok := LocationSource_name[int32(x)]
	if ok {
		return s
	}
	return strconv.Itoa(int(x))
}
