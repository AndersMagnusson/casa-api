// Code generated by protoc-gen-go. DO NOT EDIT.
// source: casa/protos/casa.proto

package communication

import proto "github.com/golang/protobuf/proto"
import fmt "fmt"
import math "math"
import timestamp "github.com/golang/protobuf/ptypes/timestamp"

import (
	context "golang.org/x/net/context"
	grpc "google.golang.org/grpc"
)

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// This is a compile-time assertion to ensure that this generated file
// is compatible with the proto package it is being compiled against.
// A compilation error at this line likely means your copy of the
// proto package needs to be updated.
const _ = proto.ProtoPackageIsVersion2 // please upgrade the proto package

type LoginRequest struct {
	Password             string   `protobuf:"bytes,1,opt,name=password,proto3" json:"password,omitempty"`
	Name                 string   `protobuf:"bytes,2,opt,name=name,proto3" json:"name,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *LoginRequest) Reset()         { *m = LoginRequest{} }
func (m *LoginRequest) String() string { return proto.CompactTextString(m) }
func (*LoginRequest) ProtoMessage()    {}
func (*LoginRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_casa_ce0b1d6a277a7944, []int{0}
}
func (m *LoginRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_LoginRequest.Unmarshal(m, b)
}
func (m *LoginRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_LoginRequest.Marshal(b, m, deterministic)
}
func (dst *LoginRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_LoginRequest.Merge(dst, src)
}
func (m *LoginRequest) XXX_Size() int {
	return xxx_messageInfo_LoginRequest.Size(m)
}
func (m *LoginRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_LoginRequest.DiscardUnknown(m)
}

var xxx_messageInfo_LoginRequest proto.InternalMessageInfo

func (m *LoginRequest) GetPassword() string {
	if m != nil {
		return m.Password
	}
	return ""
}

func (m *LoginRequest) GetName() string {
	if m != nil {
		return m.Name
	}
	return ""
}

type LoginResponse struct {
	Token                string   `protobuf:"bytes,1,opt,name=token,proto3" json:"token,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *LoginResponse) Reset()         { *m = LoginResponse{} }
func (m *LoginResponse) String() string { return proto.CompactTextString(m) }
func (*LoginResponse) ProtoMessage()    {}
func (*LoginResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_casa_ce0b1d6a277a7944, []int{1}
}
func (m *LoginResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_LoginResponse.Unmarshal(m, b)
}
func (m *LoginResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_LoginResponse.Marshal(b, m, deterministic)
}
func (dst *LoginResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_LoginResponse.Merge(dst, src)
}
func (m *LoginResponse) XXX_Size() int {
	return xxx_messageInfo_LoginResponse.Size(m)
}
func (m *LoginResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_LoginResponse.DiscardUnknown(m)
}

var xxx_messageInfo_LoginResponse proto.InternalMessageInfo

func (m *LoginResponse) GetToken() string {
	if m != nil {
		return m.Token
	}
	return ""
}

type LogoutRequest struct {
	Token                string   `protobuf:"bytes,1,opt,name=token,proto3" json:"token,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *LogoutRequest) Reset()         { *m = LogoutRequest{} }
func (m *LogoutRequest) String() string { return proto.CompactTextString(m) }
func (*LogoutRequest) ProtoMessage()    {}
func (*LogoutRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_casa_ce0b1d6a277a7944, []int{2}
}
func (m *LogoutRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_LogoutRequest.Unmarshal(m, b)
}
func (m *LogoutRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_LogoutRequest.Marshal(b, m, deterministic)
}
func (dst *LogoutRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_LogoutRequest.Merge(dst, src)
}
func (m *LogoutRequest) XXX_Size() int {
	return xxx_messageInfo_LogoutRequest.Size(m)
}
func (m *LogoutRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_LogoutRequest.DiscardUnknown(m)
}

var xxx_messageInfo_LogoutRequest proto.InternalMessageInfo

func (m *LogoutRequest) GetToken() string {
	if m != nil {
		return m.Token
	}
	return ""
}

type LogoutResponse struct {
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *LogoutResponse) Reset()         { *m = LogoutResponse{} }
func (m *LogoutResponse) String() string { return proto.CompactTextString(m) }
func (*LogoutResponse) ProtoMessage()    {}
func (*LogoutResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_casa_ce0b1d6a277a7944, []int{3}
}
func (m *LogoutResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_LogoutResponse.Unmarshal(m, b)
}
func (m *LogoutResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_LogoutResponse.Marshal(b, m, deterministic)
}
func (dst *LogoutResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_LogoutResponse.Merge(dst, src)
}
func (m *LogoutResponse) XXX_Size() int {
	return xxx_messageInfo_LogoutResponse.Size(m)
}
func (m *LogoutResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_LogoutResponse.DiscardUnknown(m)
}

var xxx_messageInfo_LogoutResponse proto.InternalMessageInfo

type StreamRequest struct {
	Id                   string   `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
	Message              string   `protobuf:"bytes,2,opt,name=message,proto3" json:"message,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *StreamRequest) Reset()         { *m = StreamRequest{} }
func (m *StreamRequest) String() string { return proto.CompactTextString(m) }
func (*StreamRequest) ProtoMessage()    {}
func (*StreamRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_casa_ce0b1d6a277a7944, []int{4}
}
func (m *StreamRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_StreamRequest.Unmarshal(m, b)
}
func (m *StreamRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_StreamRequest.Marshal(b, m, deterministic)
}
func (dst *StreamRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_StreamRequest.Merge(dst, src)
}
func (m *StreamRequest) XXX_Size() int {
	return xxx_messageInfo_StreamRequest.Size(m)
}
func (m *StreamRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_StreamRequest.DiscardUnknown(m)
}

var xxx_messageInfo_StreamRequest proto.InternalMessageInfo

func (m *StreamRequest) GetId() string {
	if m != nil {
		return m.Id
	}
	return ""
}

func (m *StreamRequest) GetMessage() string {
	if m != nil {
		return m.Message
	}
	return ""
}

type StreamResponse struct {
	Timestamp *timestamp.Timestamp `protobuf:"bytes,1,opt,name=timestamp,proto3" json:"timestamp,omitempty"`
	// Types that are valid to be assigned to Event:
	//	*StreamResponse_ClientAlarms
	//	*StreamResponse_ClientAlerts
	//	*StreamResponse_ClientDevices
	//	*StreamResponse_ClientToggle
	Event                isStreamResponse_Event `protobuf_oneof:"event"`
	XXX_NoUnkeyedLiteral struct{}               `json:"-"`
	XXX_unrecognized     []byte                 `json:"-"`
	XXX_sizecache        int32                  `json:"-"`
}

func (m *StreamResponse) Reset()         { *m = StreamResponse{} }
func (m *StreamResponse) String() string { return proto.CompactTextString(m) }
func (*StreamResponse) ProtoMessage()    {}
func (*StreamResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_casa_ce0b1d6a277a7944, []int{5}
}
func (m *StreamResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_StreamResponse.Unmarshal(m, b)
}
func (m *StreamResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_StreamResponse.Marshal(b, m, deterministic)
}
func (dst *StreamResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_StreamResponse.Merge(dst, src)
}
func (m *StreamResponse) XXX_Size() int {
	return xxx_messageInfo_StreamResponse.Size(m)
}
func (m *StreamResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_StreamResponse.DiscardUnknown(m)
}

var xxx_messageInfo_StreamResponse proto.InternalMessageInfo

func (m *StreamResponse) GetTimestamp() *timestamp.Timestamp {
	if m != nil {
		return m.Timestamp
	}
	return nil
}

type isStreamResponse_Event interface {
	isStreamResponse_Event()
}

type StreamResponse_ClientAlarms struct {
	ClientAlarms *StreamResponse_Alarms `protobuf:"bytes,2,opt,name=client_alarms,json=clientAlarms,proto3,oneof"`
}

type StreamResponse_ClientAlerts struct {
	ClientAlerts *StreamResponse_Alerts `protobuf:"bytes,3,opt,name=client_alerts,json=clientAlerts,proto3,oneof"`
}

type StreamResponse_ClientDevices struct {
	ClientDevices *StreamResponse_Devices `protobuf:"bytes,4,opt,name=client_devices,json=clientDevices,proto3,oneof"`
}

type StreamResponse_ClientToggle struct {
	ClientToggle *StreamResponse_Toggle `protobuf:"bytes,5,opt,name=client_toggle,json=clientToggle,proto3,oneof"`
}

func (*StreamResponse_ClientAlarms) isStreamResponse_Event() {}

func (*StreamResponse_ClientAlerts) isStreamResponse_Event() {}

func (*StreamResponse_ClientDevices) isStreamResponse_Event() {}

func (*StreamResponse_ClientToggle) isStreamResponse_Event() {}

func (m *StreamResponse) GetEvent() isStreamResponse_Event {
	if m != nil {
		return m.Event
	}
	return nil
}

func (m *StreamResponse) GetClientAlarms() *StreamResponse_Alarms {
	if x, ok := m.GetEvent().(*StreamResponse_ClientAlarms); ok {
		return x.ClientAlarms
	}
	return nil
}

func (m *StreamResponse) GetClientAlerts() *StreamResponse_Alerts {
	if x, ok := m.GetEvent().(*StreamResponse_ClientAlerts); ok {
		return x.ClientAlerts
	}
	return nil
}

func (m *StreamResponse) GetClientDevices() *StreamResponse_Devices {
	if x, ok := m.GetEvent().(*StreamResponse_ClientDevices); ok {
		return x.ClientDevices
	}
	return nil
}

func (m *StreamResponse) GetClientToggle() *StreamResponse_Toggle {
	if x, ok := m.GetEvent().(*StreamResponse_ClientToggle); ok {
		return x.ClientToggle
	}
	return nil
}

// XXX_OneofFuncs is for the internal use of the proto package.
func (*StreamResponse) XXX_OneofFuncs() (func(msg proto.Message, b *proto.Buffer) error, func(msg proto.Message, tag, wire int, b *proto.Buffer) (bool, error), func(msg proto.Message) (n int), []interface{}) {
	return _StreamResponse_OneofMarshaler, _StreamResponse_OneofUnmarshaler, _StreamResponse_OneofSizer, []interface{}{
		(*StreamResponse_ClientAlarms)(nil),
		(*StreamResponse_ClientAlerts)(nil),
		(*StreamResponse_ClientDevices)(nil),
		(*StreamResponse_ClientToggle)(nil),
	}
}

func _StreamResponse_OneofMarshaler(msg proto.Message, b *proto.Buffer) error {
	m := msg.(*StreamResponse)
	// event
	switch x := m.Event.(type) {
	case *StreamResponse_ClientAlarms:
		b.EncodeVarint(2<<3 | proto.WireBytes)
		if err := b.EncodeMessage(x.ClientAlarms); err != nil {
			return err
		}
	case *StreamResponse_ClientAlerts:
		b.EncodeVarint(3<<3 | proto.WireBytes)
		if err := b.EncodeMessage(x.ClientAlerts); err != nil {
			return err
		}
	case *StreamResponse_ClientDevices:
		b.EncodeVarint(4<<3 | proto.WireBytes)
		if err := b.EncodeMessage(x.ClientDevices); err != nil {
			return err
		}
	case *StreamResponse_ClientToggle:
		b.EncodeVarint(5<<3 | proto.WireBytes)
		if err := b.EncodeMessage(x.ClientToggle); err != nil {
			return err
		}
	case nil:
	default:
		return fmt.Errorf("StreamResponse.Event has unexpected type %T", x)
	}
	return nil
}

func _StreamResponse_OneofUnmarshaler(msg proto.Message, tag, wire int, b *proto.Buffer) (bool, error) {
	m := msg.(*StreamResponse)
	switch tag {
	case 2: // event.client_alarms
		if wire != proto.WireBytes {
			return true, proto.ErrInternalBadWireType
		}
		msg := new(StreamResponse_Alarms)
		err := b.DecodeMessage(msg)
		m.Event = &StreamResponse_ClientAlarms{msg}
		return true, err
	case 3: // event.client_alerts
		if wire != proto.WireBytes {
			return true, proto.ErrInternalBadWireType
		}
		msg := new(StreamResponse_Alerts)
		err := b.DecodeMessage(msg)
		m.Event = &StreamResponse_ClientAlerts{msg}
		return true, err
	case 4: // event.client_devices
		if wire != proto.WireBytes {
			return true, proto.ErrInternalBadWireType
		}
		msg := new(StreamResponse_Devices)
		err := b.DecodeMessage(msg)
		m.Event = &StreamResponse_ClientDevices{msg}
		return true, err
	case 5: // event.client_toggle
		if wire != proto.WireBytes {
			return true, proto.ErrInternalBadWireType
		}
		msg := new(StreamResponse_Toggle)
		err := b.DecodeMessage(msg)
		m.Event = &StreamResponse_ClientToggle{msg}
		return true, err
	default:
		return false, nil
	}
}

func _StreamResponse_OneofSizer(msg proto.Message) (n int) {
	m := msg.(*StreamResponse)
	// event
	switch x := m.Event.(type) {
	case *StreamResponse_ClientAlarms:
		s := proto.Size(x.ClientAlarms)
		n += 1 // tag and wire
		n += proto.SizeVarint(uint64(s))
		n += s
	case *StreamResponse_ClientAlerts:
		s := proto.Size(x.ClientAlerts)
		n += 1 // tag and wire
		n += proto.SizeVarint(uint64(s))
		n += s
	case *StreamResponse_ClientDevices:
		s := proto.Size(x.ClientDevices)
		n += 1 // tag and wire
		n += proto.SizeVarint(uint64(s))
		n += s
	case *StreamResponse_ClientToggle:
		s := proto.Size(x.ClientToggle)
		n += 1 // tag and wire
		n += proto.SizeVarint(uint64(s))
		n += s
	case nil:
	default:
		panic(fmt.Sprintf("proto: unexpected type %T in oneof", x))
	}
	return n
}

type StreamResponse_Login struct {
	Name                 string   `protobuf:"bytes,1,opt,name=name,proto3" json:"name,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *StreamResponse_Login) Reset()         { *m = StreamResponse_Login{} }
func (m *StreamResponse_Login) String() string { return proto.CompactTextString(m) }
func (*StreamResponse_Login) ProtoMessage()    {}
func (*StreamResponse_Login) Descriptor() ([]byte, []int) {
	return fileDescriptor_casa_ce0b1d6a277a7944, []int{5, 0}
}
func (m *StreamResponse_Login) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_StreamResponse_Login.Unmarshal(m, b)
}
func (m *StreamResponse_Login) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_StreamResponse_Login.Marshal(b, m, deterministic)
}
func (dst *StreamResponse_Login) XXX_Merge(src proto.Message) {
	xxx_messageInfo_StreamResponse_Login.Merge(dst, src)
}
func (m *StreamResponse_Login) XXX_Size() int {
	return xxx_messageInfo_StreamResponse_Login.Size(m)
}
func (m *StreamResponse_Login) XXX_DiscardUnknown() {
	xxx_messageInfo_StreamResponse_Login.DiscardUnknown(m)
}

var xxx_messageInfo_StreamResponse_Login proto.InternalMessageInfo

func (m *StreamResponse_Login) GetName() string {
	if m != nil {
		return m.Name
	}
	return ""
}

type StreamResponse_Logout struct {
	Name                 string   `protobuf:"bytes,1,opt,name=name,proto3" json:"name,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *StreamResponse_Logout) Reset()         { *m = StreamResponse_Logout{} }
func (m *StreamResponse_Logout) String() string { return proto.CompactTextString(m) }
func (*StreamResponse_Logout) ProtoMessage()    {}
func (*StreamResponse_Logout) Descriptor() ([]byte, []int) {
	return fileDescriptor_casa_ce0b1d6a277a7944, []int{5, 1}
}
func (m *StreamResponse_Logout) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_StreamResponse_Logout.Unmarshal(m, b)
}
func (m *StreamResponse_Logout) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_StreamResponse_Logout.Marshal(b, m, deterministic)
}
func (dst *StreamResponse_Logout) XXX_Merge(src proto.Message) {
	xxx_messageInfo_StreamResponse_Logout.Merge(dst, src)
}
func (m *StreamResponse_Logout) XXX_Size() int {
	return xxx_messageInfo_StreamResponse_Logout.Size(m)
}
func (m *StreamResponse_Logout) XXX_DiscardUnknown() {
	xxx_messageInfo_StreamResponse_Logout.DiscardUnknown(m)
}

var xxx_messageInfo_StreamResponse_Logout proto.InternalMessageInfo

func (m *StreamResponse_Logout) GetName() string {
	if m != nil {
		return m.Name
	}
	return ""
}

type StreamResponse_Alarms struct {
	Id                   string   `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
	Method               string   `protobuf:"bytes,2,opt,name=method,proto3" json:"method,omitempty"`
	Identifier           string   `protobuf:"bytes,3,opt,name=identifier,proto3" json:"identifier,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *StreamResponse_Alarms) Reset()         { *m = StreamResponse_Alarms{} }
func (m *StreamResponse_Alarms) String() string { return proto.CompactTextString(m) }
func (*StreamResponse_Alarms) ProtoMessage()    {}
func (*StreamResponse_Alarms) Descriptor() ([]byte, []int) {
	return fileDescriptor_casa_ce0b1d6a277a7944, []int{5, 2}
}
func (m *StreamResponse_Alarms) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_StreamResponse_Alarms.Unmarshal(m, b)
}
func (m *StreamResponse_Alarms) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_StreamResponse_Alarms.Marshal(b, m, deterministic)
}
func (dst *StreamResponse_Alarms) XXX_Merge(src proto.Message) {
	xxx_messageInfo_StreamResponse_Alarms.Merge(dst, src)
}
func (m *StreamResponse_Alarms) XXX_Size() int {
	return xxx_messageInfo_StreamResponse_Alarms.Size(m)
}
func (m *StreamResponse_Alarms) XXX_DiscardUnknown() {
	xxx_messageInfo_StreamResponse_Alarms.DiscardUnknown(m)
}

var xxx_messageInfo_StreamResponse_Alarms proto.InternalMessageInfo

func (m *StreamResponse_Alarms) GetId() string {
	if m != nil {
		return m.Id
	}
	return ""
}

func (m *StreamResponse_Alarms) GetMethod() string {
	if m != nil {
		return m.Method
	}
	return ""
}

func (m *StreamResponse_Alarms) GetIdentifier() string {
	if m != nil {
		return m.Identifier
	}
	return ""
}

type StreamResponse_Alerts struct {
	Method               string   `protobuf:"bytes,1,opt,name=method,proto3" json:"method,omitempty"`
	Identifier           string   `protobuf:"bytes,2,opt,name=identifier,proto3" json:"identifier,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *StreamResponse_Alerts) Reset()         { *m = StreamResponse_Alerts{} }
func (m *StreamResponse_Alerts) String() string { return proto.CompactTextString(m) }
func (*StreamResponse_Alerts) ProtoMessage()    {}
func (*StreamResponse_Alerts) Descriptor() ([]byte, []int) {
	return fileDescriptor_casa_ce0b1d6a277a7944, []int{5, 3}
}
func (m *StreamResponse_Alerts) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_StreamResponse_Alerts.Unmarshal(m, b)
}
func (m *StreamResponse_Alerts) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_StreamResponse_Alerts.Marshal(b, m, deterministic)
}
func (dst *StreamResponse_Alerts) XXX_Merge(src proto.Message) {
	xxx_messageInfo_StreamResponse_Alerts.Merge(dst, src)
}
func (m *StreamResponse_Alerts) XXX_Size() int {
	return xxx_messageInfo_StreamResponse_Alerts.Size(m)
}
func (m *StreamResponse_Alerts) XXX_DiscardUnknown() {
	xxx_messageInfo_StreamResponse_Alerts.DiscardUnknown(m)
}

var xxx_messageInfo_StreamResponse_Alerts proto.InternalMessageInfo

func (m *StreamResponse_Alerts) GetMethod() string {
	if m != nil {
		return m.Method
	}
	return ""
}

func (m *StreamResponse_Alerts) GetIdentifier() string {
	if m != nil {
		return m.Identifier
	}
	return ""
}

type StreamResponse_Devices struct {
	Method               string   `protobuf:"bytes,1,opt,name=method,proto3" json:"method,omitempty"`
	Identifier           string   `protobuf:"bytes,2,opt,name=identifier,proto3" json:"identifier,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *StreamResponse_Devices) Reset()         { *m = StreamResponse_Devices{} }
func (m *StreamResponse_Devices) String() string { return proto.CompactTextString(m) }
func (*StreamResponse_Devices) ProtoMessage()    {}
func (*StreamResponse_Devices) Descriptor() ([]byte, []int) {
	return fileDescriptor_casa_ce0b1d6a277a7944, []int{5, 4}
}
func (m *StreamResponse_Devices) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_StreamResponse_Devices.Unmarshal(m, b)
}
func (m *StreamResponse_Devices) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_StreamResponse_Devices.Marshal(b, m, deterministic)
}
func (dst *StreamResponse_Devices) XXX_Merge(src proto.Message) {
	xxx_messageInfo_StreamResponse_Devices.Merge(dst, src)
}
func (m *StreamResponse_Devices) XXX_Size() int {
	return xxx_messageInfo_StreamResponse_Devices.Size(m)
}
func (m *StreamResponse_Devices) XXX_DiscardUnknown() {
	xxx_messageInfo_StreamResponse_Devices.DiscardUnknown(m)
}

var xxx_messageInfo_StreamResponse_Devices proto.InternalMessageInfo

func (m *StreamResponse_Devices) GetMethod() string {
	if m != nil {
		return m.Method
	}
	return ""
}

func (m *StreamResponse_Devices) GetIdentifier() string {
	if m != nil {
		return m.Identifier
	}
	return ""
}

type StreamResponse_Toggle struct {
	On                   bool     `protobuf:"varint,1,opt,name=on,proto3" json:"on,omitempty"`
	AlarmId              string   `protobuf:"bytes,2,opt,name=alarmId,proto3" json:"alarmId,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *StreamResponse_Toggle) Reset()         { *m = StreamResponse_Toggle{} }
func (m *StreamResponse_Toggle) String() string { return proto.CompactTextString(m) }
func (*StreamResponse_Toggle) ProtoMessage()    {}
func (*StreamResponse_Toggle) Descriptor() ([]byte, []int) {
	return fileDescriptor_casa_ce0b1d6a277a7944, []int{5, 5}
}
func (m *StreamResponse_Toggle) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_StreamResponse_Toggle.Unmarshal(m, b)
}
func (m *StreamResponse_Toggle) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_StreamResponse_Toggle.Marshal(b, m, deterministic)
}
func (dst *StreamResponse_Toggle) XXX_Merge(src proto.Message) {
	xxx_messageInfo_StreamResponse_Toggle.Merge(dst, src)
}
func (m *StreamResponse_Toggle) XXX_Size() int {
	return xxx_messageInfo_StreamResponse_Toggle.Size(m)
}
func (m *StreamResponse_Toggle) XXX_DiscardUnknown() {
	xxx_messageInfo_StreamResponse_Toggle.DiscardUnknown(m)
}

var xxx_messageInfo_StreamResponse_Toggle proto.InternalMessageInfo

func (m *StreamResponse_Toggle) GetOn() bool {
	if m != nil {
		return m.On
	}
	return false
}

func (m *StreamResponse_Toggle) GetAlarmId() string {
	if m != nil {
		return m.AlarmId
	}
	return ""
}

func init() {
	proto.RegisterType((*LoginRequest)(nil), "communication.LoginRequest")
	proto.RegisterType((*LoginResponse)(nil), "communication.LoginResponse")
	proto.RegisterType((*LogoutRequest)(nil), "communication.LogoutRequest")
	proto.RegisterType((*LogoutResponse)(nil), "communication.LogoutResponse")
	proto.RegisterType((*StreamRequest)(nil), "communication.StreamRequest")
	proto.RegisterType((*StreamResponse)(nil), "communication.StreamResponse")
	proto.RegisterType((*StreamResponse_Login)(nil), "communication.StreamResponse.Login")
	proto.RegisterType((*StreamResponse_Logout)(nil), "communication.StreamResponse.Logout")
	proto.RegisterType((*StreamResponse_Alarms)(nil), "communication.StreamResponse.Alarms")
	proto.RegisterType((*StreamResponse_Alerts)(nil), "communication.StreamResponse.Alerts")
	proto.RegisterType((*StreamResponse_Devices)(nil), "communication.StreamResponse.Devices")
	proto.RegisterType((*StreamResponse_Toggle)(nil), "communication.StreamResponse.Toggle")
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConn

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion4

// CasaClient is the client API for Casa service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type CasaClient interface {
	Login(ctx context.Context, in *LoginRequest, opts ...grpc.CallOption) (*LoginResponse, error)
	Logout(ctx context.Context, in *LogoutRequest, opts ...grpc.CallOption) (*LogoutResponse, error)
	Stream(ctx context.Context, opts ...grpc.CallOption) (Casa_StreamClient, error)
}

type casaClient struct {
	cc *grpc.ClientConn
}

func NewCasaClient(cc *grpc.ClientConn) CasaClient {
	return &casaClient{cc}
}

func (c *casaClient) Login(ctx context.Context, in *LoginRequest, opts ...grpc.CallOption) (*LoginResponse, error) {
	out := new(LoginResponse)
	err := c.cc.Invoke(ctx, "/communication.Casa/Login", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *casaClient) Logout(ctx context.Context, in *LogoutRequest, opts ...grpc.CallOption) (*LogoutResponse, error) {
	out := new(LogoutResponse)
	err := c.cc.Invoke(ctx, "/communication.Casa/Logout", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *casaClient) Stream(ctx context.Context, opts ...grpc.CallOption) (Casa_StreamClient, error) {
	stream, err := c.cc.NewStream(ctx, &_Casa_serviceDesc.Streams[0], "/communication.Casa/Stream", opts...)
	if err != nil {
		return nil, err
	}
	x := &casaStreamClient{stream}
	return x, nil
}

type Casa_StreamClient interface {
	Send(*StreamRequest) error
	Recv() (*StreamResponse, error)
	grpc.ClientStream
}

type casaStreamClient struct {
	grpc.ClientStream
}

func (x *casaStreamClient) Send(m *StreamRequest) error {
	return x.ClientStream.SendMsg(m)
}

func (x *casaStreamClient) Recv() (*StreamResponse, error) {
	m := new(StreamResponse)
	if err := x.ClientStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

// CasaServer is the server API for Casa service.
type CasaServer interface {
	Login(context.Context, *LoginRequest) (*LoginResponse, error)
	Logout(context.Context, *LogoutRequest) (*LogoutResponse, error)
	Stream(Casa_StreamServer) error
}

func RegisterCasaServer(s *grpc.Server, srv CasaServer) {
	s.RegisterService(&_Casa_serviceDesc, srv)
}

func _Casa_Login_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(LoginRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(CasaServer).Login(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/communication.Casa/Login",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(CasaServer).Login(ctx, req.(*LoginRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Casa_Logout_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(LogoutRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(CasaServer).Logout(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/communication.Casa/Logout",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(CasaServer).Logout(ctx, req.(*LogoutRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Casa_Stream_Handler(srv interface{}, stream grpc.ServerStream) error {
	return srv.(CasaServer).Stream(&casaStreamServer{stream})
}

type Casa_StreamServer interface {
	Send(*StreamResponse) error
	Recv() (*StreamRequest, error)
	grpc.ServerStream
}

type casaStreamServer struct {
	grpc.ServerStream
}

func (x *casaStreamServer) Send(m *StreamResponse) error {
	return x.ServerStream.SendMsg(m)
}

func (x *casaStreamServer) Recv() (*StreamRequest, error) {
	m := new(StreamRequest)
	if err := x.ServerStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

var _Casa_serviceDesc = grpc.ServiceDesc{
	ServiceName: "communication.Casa",
	HandlerType: (*CasaServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "Login",
			Handler:    _Casa_Login_Handler,
		},
		{
			MethodName: "Logout",
			Handler:    _Casa_Logout_Handler,
		},
	},
	Streams: []grpc.StreamDesc{
		{
			StreamName:    "Stream",
			Handler:       _Casa_Stream_Handler,
			ServerStreams: true,
			ClientStreams: true,
		},
	},
	Metadata: "casa/protos/casa.proto",
}

func init() { proto.RegisterFile("casa/protos/casa.proto", fileDescriptor_casa_ce0b1d6a277a7944) }

var fileDescriptor_casa_ce0b1d6a277a7944 = []byte{
	// 500 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x9c, 0x93, 0xcd, 0x6e, 0xd3, 0x40,
	0x10, 0xc7, 0x63, 0x13, 0x3b, 0xcd, 0xb4, 0x89, 0xd0, 0x0a, 0x55, 0x96, 0x5b, 0x3e, 0x64, 0x51,
	0xa9, 0x27, 0x07, 0x85, 0x0b, 0x5c, 0x10, 0x85, 0x4a, 0x14, 0x15, 0x21, 0x64, 0x7a, 0x47, 0x5b,
	0x7b, 0x6a, 0x56, 0xc4, 0xbb, 0xc1, 0xbb, 0x29, 0xaf, 0xc9, 0x33, 0xf0, 0x24, 0xc8, 0xfb, 0x61,
	0xc7, 0x4d, 0xa8, 0x22, 0x6e, 0x9e, 0xdd, 0xff, 0xfc, 0x76, 0xfc, 0x9f, 0x19, 0x38, 0xcc, 0xa9,
	0xa4, 0xb3, 0x65, 0x2d, 0x94, 0x90, 0xb3, 0xe6, 0x3b, 0xd5, 0xdf, 0x64, 0x92, 0x8b, 0xaa, 0x5a,
	0x71, 0x96, 0x53, 0xc5, 0x04, 0x8f, 0x9f, 0x96, 0x42, 0x94, 0x0b, 0x34, 0xc2, 0xeb, 0xd5, 0xcd,
	0x4c, 0xb1, 0x0a, 0xa5, 0xa2, 0xd5, 0xd2, 0xe8, 0x93, 0x37, 0x70, 0xf0, 0x49, 0x94, 0x8c, 0x67,
	0xf8, 0x73, 0x85, 0x52, 0x91, 0x18, 0xf6, 0x96, 0x54, 0xca, 0x5f, 0xa2, 0x2e, 0x22, 0xef, 0x99,
	0x77, 0x3a, 0xce, 0xda, 0x98, 0x10, 0x18, 0x72, 0x5a, 0x61, 0xe4, 0xeb, 0x73, 0xfd, 0x9d, 0x9c,
	0xc0, 0xc4, 0xe6, 0xcb, 0xa5, 0xe0, 0x12, 0xc9, 0x23, 0x08, 0x94, 0xf8, 0x81, 0xdc, 0x66, 0x9b,
	0xc0, 0xca, 0xc4, 0x4a, 0xb9, 0x77, 0xb6, 0xcb, 0x1e, 0xc2, 0xd4, 0xc9, 0x0c, 0x2e, 0x79, 0x0d,
	0x93, 0xaf, 0xaa, 0x46, 0x5a, 0xb9, 0xc4, 0x29, 0xf8, 0xcc, 0x95, 0xe6, 0xb3, 0x82, 0x44, 0x30,
	0xaa, 0x50, 0x4a, 0x5a, 0xba, 0xba, 0x5c, 0x98, 0xfc, 0x0e, 0x60, 0xea, 0x72, 0x6d, 0x71, 0xaf,
	0x60, 0xdc, 0x1a, 0xa0, 0x19, 0xfb, 0xf3, 0x38, 0x35, 0x16, 0xa5, 0xce, 0xa2, 0xf4, 0xca, 0x29,
	0xb2, 0x4e, 0x4c, 0x2e, 0x61, 0x92, 0x2f, 0x18, 0x72, 0xf5, 0x8d, 0x2e, 0x68, 0x5d, 0x49, 0xfd,
	0xd8, 0xfe, 0xfc, 0x79, 0xda, 0xf3, 0x3b, 0xed, 0xbf, 0x97, 0x9e, 0x69, 0xed, 0xc5, 0x20, 0x3b,
	0x30, 0xc9, 0x26, 0xee, 0xc1, 0xb0, 0x56, 0x32, 0x7a, 0xb0, 0x1b, 0xac, 0xd1, 0xae, 0xc3, 0x9a,
	0x98, 0x7c, 0x86, 0xa9, 0x85, 0x15, 0x78, 0xcb, 0x72, 0x94, 0xd1, 0x50, 0xd3, 0x4e, 0xee, 0xa7,
	0x9d, 0x1b, 0xf1, 0xc5, 0x20, 0xb3, 0xb5, 0xd8, 0x83, 0xb5, 0xe2, 0x94, 0x28, 0xcb, 0x05, 0x46,
	0xc1, 0x2e, 0xc5, 0x5d, 0x69, 0x6d, 0x57, 0x9c, 0x89, 0xe3, 0x23, 0x08, 0xf4, 0x78, 0xb4, 0xb3,
	0xe3, 0x75, 0xb3, 0x13, 0x1f, 0x43, 0x68, 0xba, 0xbd, 0xf5, 0xf6, 0x0b, 0x84, 0xd6, 0xae, 0xbb,
	0x2d, 0x3f, 0x84, 0xb0, 0x42, 0xf5, 0x5d, 0x14, 0xb6, 0xe3, 0x36, 0x22, 0x4f, 0x00, 0x58, 0x81,
	0x5c, 0xb1, 0x1b, 0x86, 0xb5, 0xf6, 0x74, 0x9c, 0xad, 0x9d, 0xc4, 0x6f, 0x1b, 0xa2, 0xf6, 0xac,
	0x23, 0x78, 0xf7, 0x10, 0xfc, 0x0d, 0xc2, 0x19, 0x8c, 0x9c, 0x4d, 0xff, 0x8b, 0x98, 0x43, 0x68,
	0xbc, 0x69, 0x7e, 0x4b, 0x98, 0xf9, 0xdf, 0xcb, 0x7c, 0xc1, 0x9b, 0x49, 0xd6, 0xb3, 0xf5, 0xd1,
	0xfd, 0x97, 0x0b, 0xdf, 0x8d, 0x20, 0xc0, 0x5b, 0xe4, 0x6a, 0xfe, 0xc7, 0x83, 0xe1, 0x7b, 0x2a,
	0x29, 0x39, 0x77, 0xbe, 0x1e, 0xdd, 0x69, 0xcb, 0xfa, 0x32, 0xc7, 0xc7, 0xdb, 0x2f, 0xed, 0x6a,
	0x0d, 0xc8, 0x87, 0xb6, 0x01, 0x5b, 0x94, 0xdd, 0xb2, 0xc6, 0x8f, 0xff, 0x71, 0xdb, 0x82, 0x2e,
	0x21, 0x34, 0xf3, 0xb0, 0x01, 0xea, 0x2d, 0xef, 0x06, 0xa8, 0x3f, 0x44, 0xc9, 0xe0, 0xd4, 0x7b,
	0xe1, 0x5d, 0x87, 0x7a, 0x13, 0x5f, 0xfe, 0x0d, 0x00, 0x00, 0xff, 0xff, 0xf4, 0x1b, 0x3f, 0xa8,
	0xe3, 0x04, 0x00, 0x00,
}