// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: lorawan-stack/api/events.proto

package ttnpb

import (
	context "context"
	fmt "fmt"
	_ "github.com/envoyproxy/protoc-gen-validate/validate"
	_ "github.com/gogo/protobuf/gogoproto"
	proto "github.com/gogo/protobuf/proto"
	types "github.com/gogo/protobuf/types"
	golang_proto "github.com/golang/protobuf/proto"
	_ "google.golang.org/genproto/googleapis/api/annotations"
	grpc "google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
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

type Event struct {
	// Name of the event. This can be used to find the (localized) event description.
	Name string `protobuf:"bytes,1,opt,name=name,proto3" json:"name,omitempty"`
	// Time at which the event was triggered.
	Time *types.Timestamp `protobuf:"bytes,2,opt,name=time,proto3" json:"time,omitempty"`
	// Identifiers of the entity (or entities) involved.
	Identifiers []*EntityIdentifiers `protobuf:"bytes,3,rep,name=identifiers,proto3" json:"identifiers,omitempty"`
	// Optional data attached to the event.
	Data *types.Any `protobuf:"bytes,4,opt,name=data,proto3" json:"data,omitempty"`
	// Correlation IDs can be used to find related events and actions such as API calls.
	CorrelationIds []string `protobuf:"bytes,5,rep,name=correlation_ids,json=correlationIds,proto3" json:"correlation_ids,omitempty"`
	// The origin of the event. Typically the hostname of the server that created it.
	Origin string `protobuf:"bytes,6,opt,name=origin,proto3" json:"origin,omitempty"`
	// Event context, internal use only.
	Context map[string][]byte `protobuf:"bytes,7,rep,name=context,proto3" json:"context,omitempty" protobuf_key:"bytes,1,opt,name=key,proto3" protobuf_val:"bytes,2,opt,name=value,proto3"`
	// The event will be visible to a caller that has any of these rights.
	Visibility *Rights `protobuf:"bytes,8,opt,name=visibility,proto3" json:"visibility,omitempty"`
	// Details on the authentication provided by the caller that triggered this event.
	Authentication *Event_Authentication `protobuf:"bytes,9,opt,name=authentication,proto3" json:"authentication,omitempty"`
	// The IP address of the caller that triggered this event.
	RemoteIp string `protobuf:"bytes,10,opt,name=remote_ip,json=remoteIp,proto3" json:"remote_ip,omitempty"`
	// The IP address of the caller that triggered this event.
	UserAgent string `protobuf:"bytes,11,opt,name=user_agent,json=userAgent,proto3" json:"user_agent,omitempty"`
	// The unique identifier of the event, assigned on creation.
	UniqueId             string   `protobuf:"bytes,12,opt,name=unique_id,json=uniqueId,proto3" json:"unique_id,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *Event) Reset()         { *m = Event{} }
func (m *Event) String() string { return proto.CompactTextString(m) }
func (*Event) ProtoMessage()    {}
func (*Event) Descriptor() ([]byte, []int) {
	return fileDescriptor_4fd8551d68f51e44, []int{0}
}
func (m *Event) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Event.Unmarshal(m, b)
}
func (m *Event) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Event.Marshal(b, m, deterministic)
}
func (m *Event) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Event.Merge(m, src)
}
func (m *Event) XXX_Size() int {
	return xxx_messageInfo_Event.Size(m)
}
func (m *Event) XXX_DiscardUnknown() {
	xxx_messageInfo_Event.DiscardUnknown(m)
}

var xxx_messageInfo_Event proto.InternalMessageInfo

func (m *Event) GetName() string {
	if m != nil {
		return m.Name
	}
	return ""
}

func (m *Event) GetTime() *types.Timestamp {
	if m != nil {
		return m.Time
	}
	return nil
}

func (m *Event) GetIdentifiers() []*EntityIdentifiers {
	if m != nil {
		return m.Identifiers
	}
	return nil
}

func (m *Event) GetData() *types.Any {
	if m != nil {
		return m.Data
	}
	return nil
}

func (m *Event) GetCorrelationIds() []string {
	if m != nil {
		return m.CorrelationIds
	}
	return nil
}

func (m *Event) GetOrigin() string {
	if m != nil {
		return m.Origin
	}
	return ""
}

func (m *Event) GetContext() map[string][]byte {
	if m != nil {
		return m.Context
	}
	return nil
}

func (m *Event) GetVisibility() *Rights {
	if m != nil {
		return m.Visibility
	}
	return nil
}

func (m *Event) GetAuthentication() *Event_Authentication {
	if m != nil {
		return m.Authentication
	}
	return nil
}

func (m *Event) GetRemoteIp() string {
	if m != nil {
		return m.RemoteIp
	}
	return ""
}

func (m *Event) GetUserAgent() string {
	if m != nil {
		return m.UserAgent
	}
	return ""
}

func (m *Event) GetUniqueId() string {
	if m != nil {
		return m.UniqueId
	}
	return ""
}

type Event_Authentication struct {
	// The type of authentication that was used. This is typically a bearer token.
	Type string `protobuf:"bytes,1,opt,name=type,proto3" json:"type,omitempty"`
	// The type of token that was used. Common types are APIKey, AccessToken and SessionToken.
	TokenType string `protobuf:"bytes,2,opt,name=token_type,json=tokenType,proto3" json:"token_type,omitempty"`
	// The ID of the token that was used.
	TokenId              string   `protobuf:"bytes,3,opt,name=token_id,json=tokenId,proto3" json:"token_id,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *Event_Authentication) Reset()         { *m = Event_Authentication{} }
func (m *Event_Authentication) String() string { return proto.CompactTextString(m) }
func (*Event_Authentication) ProtoMessage()    {}
func (*Event_Authentication) Descriptor() ([]byte, []int) {
	return fileDescriptor_4fd8551d68f51e44, []int{0, 1}
}
func (m *Event_Authentication) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Event_Authentication.Unmarshal(m, b)
}
func (m *Event_Authentication) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Event_Authentication.Marshal(b, m, deterministic)
}
func (m *Event_Authentication) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Event_Authentication.Merge(m, src)
}
func (m *Event_Authentication) XXX_Size() int {
	return xxx_messageInfo_Event_Authentication.Size(m)
}
func (m *Event_Authentication) XXX_DiscardUnknown() {
	xxx_messageInfo_Event_Authentication.DiscardUnknown(m)
}

var xxx_messageInfo_Event_Authentication proto.InternalMessageInfo

func (m *Event_Authentication) GetType() string {
	if m != nil {
		return m.Type
	}
	return ""
}

func (m *Event_Authentication) GetTokenType() string {
	if m != nil {
		return m.TokenType
	}
	return ""
}

func (m *Event_Authentication) GetTokenId() string {
	if m != nil {
		return m.TokenId
	}
	return ""
}

type StreamEventsRequest struct {
	Identifiers []*EntityIdentifiers `protobuf:"bytes,1,rep,name=identifiers,proto3" json:"identifiers,omitempty"`
	// If greater than zero, this will return historical events, up to this maximum when the stream starts.
	// If used in combination with "after", the limit that is reached first, is used.
	// The availability of historical events depends on server support and retention policy.
	Tail uint32 `protobuf:"varint,2,opt,name=tail,proto3" json:"tail,omitempty"`
	// If not empty, this will return historical events after the given time when the stream starts.
	// If used in combination with "tail", the limit that is reached first, is used.
	// The availability of historical events depends on server support and retention policy.
	After *types.Timestamp `protobuf:"bytes,3,opt,name=after,proto3" json:"after,omitempty"`
	// If provided, this will filter events, so that only events with the given names are returned.
	// Names can be provided as either exact event names (e.g. 'gs.up.receive'),
	// or as regular expressions (e.g. '/^gs\..+/').
	Names                []string `protobuf:"bytes,4,rep,name=names,proto3" json:"names,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *StreamEventsRequest) Reset()         { *m = StreamEventsRequest{} }
func (m *StreamEventsRequest) String() string { return proto.CompactTextString(m) }
func (*StreamEventsRequest) ProtoMessage()    {}
func (*StreamEventsRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_4fd8551d68f51e44, []int{1}
}
func (m *StreamEventsRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_StreamEventsRequest.Unmarshal(m, b)
}
func (m *StreamEventsRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_StreamEventsRequest.Marshal(b, m, deterministic)
}
func (m *StreamEventsRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_StreamEventsRequest.Merge(m, src)
}
func (m *StreamEventsRequest) XXX_Size() int {
	return xxx_messageInfo_StreamEventsRequest.Size(m)
}
func (m *StreamEventsRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_StreamEventsRequest.DiscardUnknown(m)
}

var xxx_messageInfo_StreamEventsRequest proto.InternalMessageInfo

func (m *StreamEventsRequest) GetIdentifiers() []*EntityIdentifiers {
	if m != nil {
		return m.Identifiers
	}
	return nil
}

func (m *StreamEventsRequest) GetTail() uint32 {
	if m != nil {
		return m.Tail
	}
	return 0
}

func (m *StreamEventsRequest) GetAfter() *types.Timestamp {
	if m != nil {
		return m.After
	}
	return nil
}

func (m *StreamEventsRequest) GetNames() []string {
	if m != nil {
		return m.Names
	}
	return nil
}

type FindRelatedEventsRequest struct {
	CorrelationId        string   `protobuf:"bytes,1,opt,name=correlation_id,json=correlationId,proto3" json:"correlation_id,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *FindRelatedEventsRequest) Reset()         { *m = FindRelatedEventsRequest{} }
func (m *FindRelatedEventsRequest) String() string { return proto.CompactTextString(m) }
func (*FindRelatedEventsRequest) ProtoMessage()    {}
func (*FindRelatedEventsRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_4fd8551d68f51e44, []int{2}
}
func (m *FindRelatedEventsRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_FindRelatedEventsRequest.Unmarshal(m, b)
}
func (m *FindRelatedEventsRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_FindRelatedEventsRequest.Marshal(b, m, deterministic)
}
func (m *FindRelatedEventsRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_FindRelatedEventsRequest.Merge(m, src)
}
func (m *FindRelatedEventsRequest) XXX_Size() int {
	return xxx_messageInfo_FindRelatedEventsRequest.Size(m)
}
func (m *FindRelatedEventsRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_FindRelatedEventsRequest.DiscardUnknown(m)
}

var xxx_messageInfo_FindRelatedEventsRequest proto.InternalMessageInfo

func (m *FindRelatedEventsRequest) GetCorrelationId() string {
	if m != nil {
		return m.CorrelationId
	}
	return ""
}

type FindRelatedEventsResponse struct {
	Events               []*Event `protobuf:"bytes,1,rep,name=events,proto3" json:"events,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *FindRelatedEventsResponse) Reset()         { *m = FindRelatedEventsResponse{} }
func (m *FindRelatedEventsResponse) String() string { return proto.CompactTextString(m) }
func (*FindRelatedEventsResponse) ProtoMessage()    {}
func (*FindRelatedEventsResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_4fd8551d68f51e44, []int{3}
}
func (m *FindRelatedEventsResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_FindRelatedEventsResponse.Unmarshal(m, b)
}
func (m *FindRelatedEventsResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_FindRelatedEventsResponse.Marshal(b, m, deterministic)
}
func (m *FindRelatedEventsResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_FindRelatedEventsResponse.Merge(m, src)
}
func (m *FindRelatedEventsResponse) XXX_Size() int {
	return xxx_messageInfo_FindRelatedEventsResponse.Size(m)
}
func (m *FindRelatedEventsResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_FindRelatedEventsResponse.DiscardUnknown(m)
}

var xxx_messageInfo_FindRelatedEventsResponse proto.InternalMessageInfo

func (m *FindRelatedEventsResponse) GetEvents() []*Event {
	if m != nil {
		return m.Events
	}
	return nil
}

func init() {
	proto.RegisterType((*Event)(nil), "ttn.lorawan.v3.Event")
	golang_proto.RegisterType((*Event)(nil), "ttn.lorawan.v3.Event")
	proto.RegisterMapType((map[string][]byte)(nil), "ttn.lorawan.v3.Event.ContextEntry")
	golang_proto.RegisterMapType((map[string][]byte)(nil), "ttn.lorawan.v3.Event.ContextEntry")
	proto.RegisterType((*Event_Authentication)(nil), "ttn.lorawan.v3.Event.Authentication")
	golang_proto.RegisterType((*Event_Authentication)(nil), "ttn.lorawan.v3.Event.Authentication")
	proto.RegisterType((*StreamEventsRequest)(nil), "ttn.lorawan.v3.StreamEventsRequest")
	golang_proto.RegisterType((*StreamEventsRequest)(nil), "ttn.lorawan.v3.StreamEventsRequest")
	proto.RegisterType((*FindRelatedEventsRequest)(nil), "ttn.lorawan.v3.FindRelatedEventsRequest")
	golang_proto.RegisterType((*FindRelatedEventsRequest)(nil), "ttn.lorawan.v3.FindRelatedEventsRequest")
	proto.RegisterType((*FindRelatedEventsResponse)(nil), "ttn.lorawan.v3.FindRelatedEventsResponse")
	golang_proto.RegisterType((*FindRelatedEventsResponse)(nil), "ttn.lorawan.v3.FindRelatedEventsResponse")
}

func init() { proto.RegisterFile("lorawan-stack/api/events.proto", fileDescriptor_4fd8551d68f51e44) }
func init() {
	golang_proto.RegisterFile("lorawan-stack/api/events.proto", fileDescriptor_4fd8551d68f51e44)
}

var fileDescriptor_4fd8551d68f51e44 = []byte{
	// 810 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x9c, 0x54, 0x4f, 0x8f, 0x1b, 0x35,
	0x14, 0xef, 0xe4, 0xdf, 0x26, 0x4e, 0x9a, 0x16, 0x53, 0x8a, 0x77, 0x80, 0xb0, 0x4c, 0x39, 0x04,
	0xa4, 0xcc, 0xac, 0x76, 0x05, 0xaa, 0x56, 0x5c, 0x92, 0x6a, 0x91, 0x82, 0x7a, 0x32, 0x3d, 0xf5,
	0x40, 0xe4, 0x64, 0xbc, 0x13, 0x2b, 0x89, 0x3d, 0xf5, 0xbc, 0x49, 0x3b, 0xe2, 0xc6, 0x27, 0x40,
	0xf0, 0x4d, 0x38, 0x70, 0xe6, 0x03, 0x70, 0xe3, 0x13, 0xd0, 0x72, 0xe0, 0xc8, 0x79, 0x4f, 0xc8,
	0xf6, 0xa4, 0x24, 0xd9, 0x80, 0x50, 0x6f, 0x7e, 0x7e, 0xbf, 0xf7, 0xde, 0xef, 0xfd, 0xe6, 0x37,
	0x46, 0xbd, 0xa5, 0xd2, 0xec, 0x39, 0x93, 0x83, 0x0c, 0xd8, 0x6c, 0x11, 0xb1, 0x54, 0x44, 0x7c,
	0xcd, 0x25, 0x64, 0x61, 0xaa, 0x15, 0x28, 0xdc, 0x05, 0x90, 0x61, 0x89, 0x09, 0xd7, 0xe7, 0xfe,
	0x30, 0x11, 0x30, 0xcf, 0xa7, 0xe1, 0x4c, 0xad, 0x22, 0x2e, 0xd7, 0xaa, 0x48, 0xb5, 0x7a, 0x51,
	0x44, 0x16, 0x3c, 0x1b, 0x24, 0x5c, 0x0e, 0xd6, 0x6c, 0x29, 0x62, 0x06, 0x3c, 0xba, 0x71, 0x70,
	0x2d, 0xfd, 0xc1, 0x56, 0x8b, 0x44, 0x25, 0xca, 0x15, 0x4f, 0xf3, 0x2b, 0x1b, 0xd9, 0xc0, 0x9e,
	0x4a, 0xf8, 0xfb, 0x89, 0x52, 0xc9, 0x92, 0x5b, 0x6a, 0x4c, 0x4a, 0x05, 0x0c, 0x84, 0x92, 0x25,
	0x3f, 0xff, 0xb8, 0xcc, 0xbe, 0xee, 0xc1, 0x64, 0x51, 0xa6, 0x3e, 0xdc, 0x4f, 0x81, 0x58, 0xf1,
	0x0c, 0xd8, 0x2a, 0x2d, 0x01, 0x0f, 0x6e, 0xee, 0x2e, 0x62, 0x2e, 0x41, 0x5c, 0x09, 0xae, 0x37,
	0x03, 0x0e, 0x08, 0xa4, 0x45, 0x32, 0xdf, 0x08, 0x14, 0xfc, 0x5a, 0x47, 0xf5, 0x4b, 0xa3, 0x18,
	0xc6, 0xa8, 0x26, 0xd9, 0x8a, 0x13, 0xef, 0xc4, 0xeb, 0xb7, 0xa8, 0x3d, 0xe3, 0x87, 0xa8, 0x66,
	0xa6, 0x92, 0xca, 0x89, 0xd7, 0x6f, 0x9f, 0xf9, 0xa1, 0xa3, 0x14, 0x6e, 0x28, 0x85, 0x4f, 0x36,
	0x94, 0x46, 0xcd, 0xeb, 0x51, 0xfd, 0x27, 0xaf, 0xd2, 0xf4, 0xa8, 0xad, 0xc0, 0x8f, 0x50, 0x7b,
	0x8b, 0x0c, 0xa9, 0x9e, 0x54, 0xfb, 0xed, 0xb3, 0x8f, 0xc2, 0xdd, 0xcf, 0x11, 0x5e, 0x4a, 0x10,
	0x50, 0x8c, 0xff, 0x01, 0xd2, 0xed, 0x2a, 0xdc, 0x47, 0xb5, 0x98, 0x01, 0x23, 0x35, 0x3b, 0xfe,
	0xde, 0x8d, 0xf1, 0x43, 0x59, 0x50, 0x8b, 0xc0, 0x9f, 0xa1, 0x3b, 0x33, 0xa5, 0x35, 0x5f, 0x5a,
	0x75, 0x27, 0x22, 0xce, 0x48, 0xfd, 0xa4, 0xda, 0x6f, 0x8d, 0x3a, 0xd7, 0xa3, 0xd6, 0x0f, 0x5e,
	0x23, 0xa8, 0xe9, 0x0a, 0x89, 0x69, 0x77, 0x0b, 0x34, 0x8e, 0x33, 0x7c, 0x1f, 0x35, 0x94, 0x16,
	0x89, 0x90, 0xa4, 0x61, 0xb7, 0x2e, 0x23, 0xfc, 0x05, 0x3a, 0x9a, 0x29, 0x09, 0xfc, 0x05, 0x90,
	0x23, 0xcb, 0x3c, 0xb8, 0xc1, 0xdc, 0x68, 0x16, 0x3e, 0x72, 0xa0, 0x4b, 0x09, 0xba, 0xa0, 0x9b,
	0x12, 0xfc, 0x39, 0x42, 0x6b, 0x91, 0x89, 0xa9, 0x58, 0x0a, 0x28, 0x48, 0xd3, 0x92, 0xbf, 0xbf,
	0xdf, 0x80, 0xda, 0xaf, 0x40, 0xb7, 0x90, 0xf8, 0x31, 0xea, 0xb2, 0x1c, 0xe6, 0x66, 0xff, 0x99,
	0xa5, 0x48, 0x5a, 0xb6, 0xf6, 0xe3, 0xc3, 0xc3, 0x87, 0x3b, 0x58, 0xba, 0x57, 0x8b, 0xdf, 0x43,
	0x2d, 0xcd, 0x57, 0x0a, 0xf8, 0x44, 0xa4, 0x04, 0xd9, 0xf5, 0x9a, 0xee, 0x62, 0x9c, 0xe2, 0x0f,
	0x10, 0xca, 0x33, 0xae, 0x27, 0x2c, 0xe1, 0x12, 0x48, 0xdb, 0x66, 0x5b, 0xe6, 0x66, 0x68, 0x2e,
	0x4c, 0x6d, 0x2e, 0xc5, 0xb3, 0x9c, 0x4f, 0x44, 0x4c, 0x3a, 0xae, 0xd6, 0x5d, 0x8c, 0x63, 0xff,
	0x02, 0x75, 0xb6, 0xf7, 0xc6, 0x77, 0x51, 0x75, 0xc1, 0x8b, 0xd2, 0x37, 0xe6, 0x88, 0xef, 0xa1,
	0xfa, 0x9a, 0x2d, 0x73, 0xe7, 0x9b, 0x0e, 0x75, 0xc1, 0x45, 0xe5, 0xa1, 0xe7, 0x7f, 0x83, 0xba,
	0xbb, 0xb4, 0x8d, 0xed, 0xa0, 0x48, 0x5f, 0xdb, 0xce, 0x9c, 0x0d, 0x3b, 0x50, 0x0b, 0x2e, 0x27,
	0x36, 0x53, 0x71, 0xec, 0xec, 0xcd, 0x13, 0x93, 0x3e, 0x46, 0x4d, 0x97, 0x16, 0x31, 0xa9, 0xda,
	0xe4, 0x91, 0x8d, 0xc7, 0x71, 0xf0, 0xb3, 0x87, 0xde, 0xfe, 0x1a, 0x34, 0x67, 0x2b, 0xab, 0x51,
	0x46, 0xf9, 0xb3, 0x9c, 0x67, 0xb0, 0x6f, 0x47, 0xef, 0x8d, 0xec, 0x68, 0xa8, 0x32, 0xb1, 0xb4,
	0x84, 0x6e, 0x53, 0x7b, 0xc6, 0xa7, 0xa8, 0xce, 0xae, 0x80, 0x6b, 0x4b, 0xe4, 0x3f, 0x7f, 0x11,
	0xea, 0x80, 0x46, 0x1c, 0xf3, 0x6f, 0x65, 0xa4, 0x66, 0x0c, 0x4a, 0x5d, 0x10, 0x3c, 0x46, 0xe4,
	0x4b, 0x21, 0x63, 0x6a, 0xcc, 0xc9, 0xe3, 0x5d, 0xf2, 0xa7, 0xa8, 0xbb, 0x6b, 0x6e, 0x27, 0xd6,
	0xa8, 0x75, 0x3d, 0x6a, 0xe8, 0xda, 0x5d, 0x8f, 0xc4, 0xf4, 0xf6, 0x8e, 0xb1, 0x83, 0xaf, 0xd0,
	0xf1, 0x81, 0x6e, 0x59, 0xaa, 0x64, 0xc6, 0xf1, 0x00, 0x35, 0xdc, 0x1b, 0x59, 0xca, 0xf0, 0xce,
	0x41, 0x7b, 0xd1, 0x12, 0x74, 0xf6, 0xbb, 0x87, 0x1a, 0xae, 0x03, 0x7e, 0x8a, 0x1a, 0x4e, 0x5c,
	0xfc, 0x60, 0xbf, 0xe6, 0x80, 0xe8, 0xfe, 0xe1, 0xc6, 0x01, 0xfe, 0xee, 0xb7, 0x3f, 0x7e, 0xac,
	0x74, 0x82, 0xa3, 0xf2, 0xa9, 0xbe, 0xf0, 0x3e, 0x3d, 0xf5, 0xf0, 0xb7, 0xa8, 0xbd, 0x45, 0x19,
	0xf7, 0xf7, 0x6b, 0xff, 0x4d, 0x1d, 0xff, 0x93, 0xff, 0x81, 0x74, 0x9b, 0x07, 0xef, 0xda, 0xc9,
	0x6f, 0xe1, 0x3b, 0xe5, 0xe4, 0x48, 0x3b, 0xd8, 0x68, 0xf8, 0xe7, 0xcb, 0xde, 0xad, 0xbf, 0x5e,
	0xf6, 0xbc, 0xef, 0x5f, 0xf5, 0x6e, 0xfd, 0xf2, 0xaa, 0xe7, 0x3d, 0x8d, 0x12, 0x15, 0xc2, 0x9c,
	0xc3, 0x5c, 0xc8, 0x24, 0x0b, 0x25, 0x87, 0xe7, 0x4a, 0x2f, 0xa2, 0xdd, 0xc7, 0x74, 0x7d, 0x1e,
	0xa5, 0x8b, 0x24, 0x02, 0x90, 0xe9, 0x74, 0xda, 0xb0, 0x5f, 0xfc, 0xfc, 0xef, 0x00, 0x00, 0x00,
	0xff, 0xff, 0x9e, 0x42, 0xc9, 0xa5, 0x92, 0x06, 0x00, 0x00,
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConn

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion4

// EventsClient is the client API for Events service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type EventsClient interface {
	// Stream live events, optionally with a tail of historical events (depending on server support and retention policy).
	// Events may arrive out-of-order.
	Stream(ctx context.Context, in *StreamEventsRequest, opts ...grpc.CallOption) (Events_StreamClient, error)
	FindRelated(ctx context.Context, in *FindRelatedEventsRequest, opts ...grpc.CallOption) (*FindRelatedEventsResponse, error)
}

type eventsClient struct {
	cc *grpc.ClientConn
}

func NewEventsClient(cc *grpc.ClientConn) EventsClient {
	return &eventsClient{cc}
}

func (c *eventsClient) Stream(ctx context.Context, in *StreamEventsRequest, opts ...grpc.CallOption) (Events_StreamClient, error) {
	stream, err := c.cc.NewStream(ctx, &_Events_serviceDesc.Streams[0], "/ttn.lorawan.v3.Events/Stream", opts...)
	if err != nil {
		return nil, err
	}
	x := &eventsStreamClient{stream}
	if err := x.ClientStream.SendMsg(in); err != nil {
		return nil, err
	}
	if err := x.ClientStream.CloseSend(); err != nil {
		return nil, err
	}
	return x, nil
}

type Events_StreamClient interface {
	Recv() (*Event, error)
	grpc.ClientStream
}

type eventsStreamClient struct {
	grpc.ClientStream
}

func (x *eventsStreamClient) Recv() (*Event, error) {
	m := new(Event)
	if err := x.ClientStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

func (c *eventsClient) FindRelated(ctx context.Context, in *FindRelatedEventsRequest, opts ...grpc.CallOption) (*FindRelatedEventsResponse, error) {
	out := new(FindRelatedEventsResponse)
	err := c.cc.Invoke(ctx, "/ttn.lorawan.v3.Events/FindRelated", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// EventsServer is the server API for Events service.
type EventsServer interface {
	// Stream live events, optionally with a tail of historical events (depending on server support and retention policy).
	// Events may arrive out-of-order.
	Stream(*StreamEventsRequest, Events_StreamServer) error
	FindRelated(context.Context, *FindRelatedEventsRequest) (*FindRelatedEventsResponse, error)
}

// UnimplementedEventsServer can be embedded to have forward compatible implementations.
type UnimplementedEventsServer struct {
}

func (*UnimplementedEventsServer) Stream(req *StreamEventsRequest, srv Events_StreamServer) error {
	return status.Errorf(codes.Unimplemented, "method Stream not implemented")
}
func (*UnimplementedEventsServer) FindRelated(ctx context.Context, req *FindRelatedEventsRequest) (*FindRelatedEventsResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method FindRelated not implemented")
}

func RegisterEventsServer(s *grpc.Server, srv EventsServer) {
	s.RegisterService(&_Events_serviceDesc, srv)
}

func _Events_Stream_Handler(srv interface{}, stream grpc.ServerStream) error {
	m := new(StreamEventsRequest)
	if err := stream.RecvMsg(m); err != nil {
		return err
	}
	return srv.(EventsServer).Stream(m, &eventsStreamServer{stream})
}

type Events_StreamServer interface {
	Send(*Event) error
	grpc.ServerStream
}

type eventsStreamServer struct {
	grpc.ServerStream
}

func (x *eventsStreamServer) Send(m *Event) error {
	return x.ServerStream.SendMsg(m)
}

func _Events_FindRelated_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(FindRelatedEventsRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(EventsServer).FindRelated(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/ttn.lorawan.v3.Events/FindRelated",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(EventsServer).FindRelated(ctx, req.(*FindRelatedEventsRequest))
	}
	return interceptor(ctx, in, info, handler)
}

var _Events_serviceDesc = grpc.ServiceDesc{
	ServiceName: "ttn.lorawan.v3.Events",
	HandlerType: (*EventsServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "FindRelated",
			Handler:    _Events_FindRelated_Handler,
		},
	},
	Streams: []grpc.StreamDesc{
		{
			StreamName:    "Stream",
			Handler:       _Events_Stream_Handler,
			ServerStreams: true,
		},
	},
	Metadata: "lorawan-stack/api/events.proto",
}
