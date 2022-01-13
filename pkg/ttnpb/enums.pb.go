// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: lorawan-stack/api/enums.proto

package ttnpb

import (
	fmt "fmt"
	_ "github.com/TheThingsIndustries/protoc-gen-go-json/annotations"
	_ "github.com/gogo/protobuf/gogoproto"
	proto "github.com/gogo/protobuf/proto"
	golang_proto "github.com/golang/protobuf/proto"
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

type DownlinkPathConstraint int32

const (
	// Indicates that the gateway can be selected for downlink without constraints by the Network Server.
	DOWNLINK_PATH_CONSTRAINT_NONE DownlinkPathConstraint = 0
	// Indicates that the gateway can be selected for downlink only if no other or better gateway can be selected.
	DOWNLINK_PATH_CONSTRAINT_PREFER_OTHER DownlinkPathConstraint = 1
	// Indicates that this gateway will never be selected for downlink, even if that results in no available downlink path.
	DOWNLINK_PATH_CONSTRAINT_NEVER DownlinkPathConstraint = 2
)

var DownlinkPathConstraint_name = map[int32]string{
	0: "DOWNLINK_PATH_CONSTRAINT_NONE",
	1: "DOWNLINK_PATH_CONSTRAINT_PREFER_OTHER",
	2: "DOWNLINK_PATH_CONSTRAINT_NEVER",
}

var DownlinkPathConstraint_value = map[string]int32{
	"DOWNLINK_PATH_CONSTRAINT_NONE":         0,
	"DOWNLINK_PATH_CONSTRAINT_PREFER_OTHER": 1,
	"DOWNLINK_PATH_CONSTRAINT_NEVER":        2,
}

func (DownlinkPathConstraint) EnumDescriptor() ([]byte, []int) {
	return fileDescriptor_e36318a1e2f407cb, []int{0}
}

// State enum defines states that an entity can be in.
type State int32

const (
	// Denotes that the entity has been requested and is pending review by an admin.
	STATE_REQUESTED State = 0
	// Denotes that the entity has been reviewed and approved by an admin.
	STATE_APPROVED State = 1
	// Denotes that the entity has been reviewed and rejected by an admin.
	STATE_REJECTED State = 2
	// Denotes that the entity has been flagged and is pending review by an admin.
	STATE_FLAGGED State = 3
	// Denotes that the entity has been reviewed and suspended by an admin.
	STATE_SUSPENDED State = 4
)

var State_name = map[int32]string{
	0: "STATE_REQUESTED",
	1: "STATE_APPROVED",
	2: "STATE_REJECTED",
	3: "STATE_FLAGGED",
	4: "STATE_SUSPENDED",
}

var State_value = map[string]int32{
	"STATE_REQUESTED": 0,
	"STATE_APPROVED":  1,
	"STATE_REJECTED":  2,
	"STATE_FLAGGED":   3,
	"STATE_SUSPENDED": 4,
}

func (State) EnumDescriptor() ([]byte, []int) {
	return fileDescriptor_e36318a1e2f407cb, []int{1}
}

type ClusterRole int32

const (
	ClusterRole_NONE                         ClusterRole = 0
	ClusterRole_ENTITY_REGISTRY              ClusterRole = 1
	ClusterRole_ACCESS                       ClusterRole = 2
	ClusterRole_GATEWAY_SERVER               ClusterRole = 3
	ClusterRole_NETWORK_SERVER               ClusterRole = 4
	ClusterRole_APPLICATION_SERVER           ClusterRole = 5
	ClusterRole_JOIN_SERVER                  ClusterRole = 6
	ClusterRole_CRYPTO_SERVER                ClusterRole = 7
	ClusterRole_DEVICE_TEMPLATE_CONVERTER    ClusterRole = 8
	ClusterRole_DEVICE_CLAIMING_SERVER       ClusterRole = 9
	ClusterRole_GATEWAY_CONFIGURATION_SERVER ClusterRole = 10
	ClusterRole_QR_CODE_GENERATOR            ClusterRole = 11
	ClusterRole_PACKET_BROKER_AGENT          ClusterRole = 12
	ClusterRole_DEVICE_REPOSITORY            ClusterRole = 13
)

var ClusterRole_name = map[int32]string{
	0:  "NONE",
	1:  "ENTITY_REGISTRY",
	2:  "ACCESS",
	3:  "GATEWAY_SERVER",
	4:  "NETWORK_SERVER",
	5:  "APPLICATION_SERVER",
	6:  "JOIN_SERVER",
	7:  "CRYPTO_SERVER",
	8:  "DEVICE_TEMPLATE_CONVERTER",
	9:  "DEVICE_CLAIMING_SERVER",
	10: "GATEWAY_CONFIGURATION_SERVER",
	11: "QR_CODE_GENERATOR",
	12: "PACKET_BROKER_AGENT",
	13: "DEVICE_REPOSITORY",
}

var ClusterRole_value = map[string]int32{
	"NONE":                         0,
	"ENTITY_REGISTRY":              1,
	"ACCESS":                       2,
	"GATEWAY_SERVER":               3,
	"NETWORK_SERVER":               4,
	"APPLICATION_SERVER":           5,
	"JOIN_SERVER":                  6,
	"CRYPTO_SERVER":                7,
	"DEVICE_TEMPLATE_CONVERTER":    8,
	"DEVICE_CLAIMING_SERVER":       9,
	"GATEWAY_CONFIGURATION_SERVER": 10,
	"QR_CODE_GENERATOR":            11,
	"PACKET_BROKER_AGENT":          12,
	"DEVICE_REPOSITORY":            13,
}

func (ClusterRole) EnumDescriptor() ([]byte, []int) {
	return fileDescriptor_e36318a1e2f407cb, []int{2}
}

func init() {
	proto.RegisterEnum("ttn.lorawan.v3.DownlinkPathConstraint", DownlinkPathConstraint_name, DownlinkPathConstraint_value)
	golang_proto.RegisterEnum("ttn.lorawan.v3.DownlinkPathConstraint", DownlinkPathConstraint_name, DownlinkPathConstraint_value)
	proto.RegisterEnum("ttn.lorawan.v3.State", State_name, State_value)
	golang_proto.RegisterEnum("ttn.lorawan.v3.State", State_name, State_value)
	proto.RegisterEnum("ttn.lorawan.v3.ClusterRole", ClusterRole_name, ClusterRole_value)
	golang_proto.RegisterEnum("ttn.lorawan.v3.ClusterRole", ClusterRole_name, ClusterRole_value)
}

func init() { proto.RegisterFile("lorawan-stack/api/enums.proto", fileDescriptor_e36318a1e2f407cb) }
func init() {
	golang_proto.RegisterFile("lorawan-stack/api/enums.proto", fileDescriptor_e36318a1e2f407cb)
}

var fileDescriptor_e36318a1e2f407cb = []byte{
	// 599 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x74, 0x93, 0xcf, 0x52, 0xdb, 0x3a,
	0x14, 0xc6, 0xe3, 0xf0, 0xe7, 0x82, 0xb8, 0x80, 0x11, 0x73, 0xb9, 0x90, 0x81, 0xcc, 0xbd, 0x9d,
	0x76, 0xd1, 0xcc, 0x24, 0x5e, 0xb0, 0xee, 0x42, 0xc8, 0x07, 0x63, 0x12, 0x24, 0x23, 0x8b, 0x30,
	0xe9, 0xc6, 0xe3, 0x50, 0xd7, 0x49, 0x13, 0xa4, 0x4c, 0xac, 0xc0, 0xb2, 0xdb, 0xbe, 0x47, 0x77,
	0xac, 0xfa, 0x08, 0x3c, 0x0b, 0xdd, 0x74, 0x58, 0xf5, 0x11, 0x3a, 0x76, 0x92, 0x52, 0x16, 0xec,
	0x74, 0x7e, 0xe7, 0x3b, 0x47, 0xdf, 0xd1, 0x1c, 0xa1, 0x83, 0xa1, 0x1e, 0xc7, 0xb7, 0xb1, 0xaa,
	0x67, 0x26, 0xbe, 0x1a, 0x38, 0xf1, 0xa8, 0xef, 0x24, 0x6a, 0x72, 0x9d, 0x35, 0x46, 0x63, 0x6d,
	0x34, 0xde, 0x30, 0x46, 0x35, 0x66, 0x92, 0xc6, 0xcd, 0x61, 0xa5, 0x9e, 0xf6, 0x4d, 0x6f, 0xd2,
	0x6d, 0x5c, 0xe9, 0x6b, 0x27, 0xd5, 0xa9, 0x76, 0x0a, 0x59, 0x77, 0xf2, 0xb1, 0x88, 0x8a, 0xa0,
	0x38, 0x4d, 0xcb, 0x2b, 0xf4, 0x0f, 0xb9, 0xec, 0x25, 0xb2, 0xd7, 0x57, 0x69, 0xe6, 0xab, 0x0f,
	0x93, 0xcc, 0x8c, 0xfb, 0x49, 0x36, 0xad, 0xbe, 0xaa, 0xa7, 0x89, 0xaa, 0xa7, 0xba, 0xfe, 0x29,
	0xd3, 0xca, 0x89, 0x95, 0xd2, 0x26, 0x36, 0x7d, 0xad, 0x66, 0x1e, 0x6a, 0xdf, 0x2c, 0xb4, 0xe3,
	0xea, 0x5b, 0x35, 0xec, 0xab, 0x41, 0x10, 0x9b, 0x1e, 0xd5, 0x2a, 0x33, 0xe3, 0xb8, 0xaf, 0x0c,
	0xfe, 0x1f, 0x1d, 0xb8, 0xfc, 0x92, 0xb5, 0x7c, 0xd6, 0x8c, 0x02, 0x22, 0x4f, 0x22, 0xca, 0x59,
	0x28, 0x05, 0xf1, 0x99, 0x8c, 0x18, 0x67, 0x60, 0x97, 0xf0, 0x5b, 0xf4, 0xe6, 0x45, 0x49, 0x20,
	0xe0, 0x18, 0x44, 0xc4, 0xe5, 0x09, 0x08, 0xdb, 0xc2, 0xaf, 0x50, 0xf5, 0xe5, 0x6e, 0xd0, 0x06,
	0x61, 0x97, 0x2b, 0xaf, 0x1f, 0xef, 0xf6, 0xf6, 0x77, 0xad, 0xda, 0xee, 0x4b, 0xca, 0x2f, 0x5f,
	0xab, 0xa5, 0xda, 0x67, 0xb4, 0x14, 0x9a, 0xd8, 0x24, 0x78, 0x1b, 0x6d, 0x86, 0x92, 0x48, 0x88,
	0x04, 0x9c, 0x5f, 0x40, 0x28, 0xc1, 0xb5, 0x4b, 0x18, 0xa3, 0x8d, 0x29, 0x24, 0x41, 0x20, 0x78,
	0x1b, 0x5c, 0xdb, 0x7a, 0x62, 0x02, 0x4e, 0x81, 0xe6, 0xba, 0x32, 0xde, 0x42, 0xeb, 0x53, 0x76,
	0xdc, 0x22, 0x9e, 0x07, 0xae, 0xbd, 0xf0, 0xd4, 0x2f, 0xbc, 0x08, 0x03, 0x60, 0x2e, 0xb8, 0xf6,
	0x62, 0x65, 0xeb, 0xf1, 0x6e, 0x6f, 0x75, 0xd7, 0xaa, 0x2d, 0x15, 0xa9, 0xc2, 0xc0, 0x7d, 0x19,
	0xad, 0xd1, 0xe1, 0x24, 0x33, 0xc9, 0x58, 0xe8, 0x61, 0x82, 0x57, 0xd0, 0xe2, 0xec, 0x3d, 0xb6,
	0xd1, 0x26, 0x30, 0xe9, 0xcb, 0x4e, 0x24, 0xc0, 0xf3, 0x43, 0x29, 0x3a, 0xb6, 0x85, 0x11, 0x5a,
	0x26, 0x94, 0x42, 0x18, 0xda, 0xe5, 0xdc, 0x89, 0x47, 0x24, 0x5c, 0x92, 0x4e, 0x14, 0x82, 0xc8,
	0xa7, 0x5e, 0xc8, 0x19, 0x03, 0x79, 0xc9, 0x45, 0x73, 0xce, 0x16, 0xf1, 0x0e, 0xc2, 0x24, 0x08,
	0x5a, 0x3e, 0x25, 0xd2, 0xe7, 0x6c, 0xce, 0x97, 0xf0, 0x26, 0x5a, 0x3b, 0xe5, 0xfe, 0x6f, 0xb0,
	0x9c, 0x8f, 0x41, 0x45, 0x27, 0x90, 0x7c, 0x8e, 0xfe, 0xc2, 0x07, 0x68, 0xcf, 0x85, 0xb6, 0x4f,
	0x21, 0x92, 0x70, 0x16, 0xb4, 0xf2, 0x81, 0x28, 0x67, 0x6d, 0x10, 0x12, 0x84, 0xbd, 0x82, 0x2b,
	0x68, 0x67, 0x96, 0xa6, 0x2d, 0xe2, 0x9f, 0xf9, 0xcc, 0x9b, 0x97, 0xae, 0xe2, 0xff, 0xd0, 0xfe,
	0xdc, 0x1e, 0xe5, 0xec, 0xd8, 0xf7, 0x2e, 0xc4, 0x33, 0x03, 0x08, 0xff, 0x83, 0xb6, 0xce, 0x45,
	0x44, 0xb9, 0x0b, 0x91, 0x07, 0x0c, 0x04, 0x91, 0x5c, 0xd8, 0x6b, 0xf8, 0x5f, 0xb4, 0x1d, 0x10,
	0xda, 0x04, 0x19, 0x1d, 0x09, 0xde, 0x04, 0x11, 0x11, 0x0f, 0x98, 0xb4, 0xff, 0xce, 0xf5, 0xb3,
	0xdb, 0x04, 0x04, 0x3c, 0xf4, 0x25, 0x17, 0x1d, 0x7b, 0xfd, 0xe8, 0xdd, 0x8f, 0x87, 0x6a, 0xe9,
	0xe7, 0x43, 0xd5, 0xba, 0xff, 0x5e, 0xb5, 0xde, 0x3b, 0xa9, 0x6e, 0x98, 0x5e, 0x62, 0x8a, 0x0d,
	0x6e, 0xa8, 0xc4, 0xdc, 0xea, 0xf1, 0xc0, 0x79, 0xfe, 0x79, 0x6e, 0x0e, 0x9d, 0xd1, 0x20, 0x75,
	0x8c, 0x51, 0xa3, 0x6e, 0x77, 0xb9, 0x58, 0xde, 0xc3, 0x5f, 0x01, 0x00, 0x00, 0xff, 0xff, 0x76,
	0xd3, 0x61, 0x3a, 0x61, 0x03, 0x00, 0x00,
}

func (x DownlinkPathConstraint) String() string {
	s, ok := DownlinkPathConstraint_name[int32(x)]
	if ok {
		return s
	}
	return strconv.Itoa(int(x))
}
func (x State) String() string {
	s, ok := State_name[int32(x)]
	if ok {
		return s
	}
	return strconv.Itoa(int(x))
}
func (x ClusterRole) String() string {
	s, ok := ClusterRole_name[int32(x)]
	if ok {
		return s
	}
	return strconv.Itoa(int(x))
}
