// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.25.0-devel
// 	protoc        v3.13.0
// source: remoteio.proto

package remoteio

import (
	context "context"
	grpc "google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
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

type PinModeMessage_PortType int32

const (
	PinModeMessage_DIGITAL_IN  PinModeMessage_PortType = 0
	PinModeMessage_DIGITAL_OUT PinModeMessage_PortType = 1
	PinModeMessage_ANALOG_IN   PinModeMessage_PortType = 2
	PinModeMessage_ANALOG_OUT  PinModeMessage_PortType = 3
)

// Enum value maps for PinModeMessage_PortType.
var (
	PinModeMessage_PortType_name = map[int32]string{
		0: "DIGITAL_IN",
		1: "DIGITAL_OUT",
		2: "ANALOG_IN",
		3: "ANALOG_OUT",
	}
	PinModeMessage_PortType_value = map[string]int32{
		"DIGITAL_IN":  0,
		"DIGITAL_OUT": 1,
		"ANALOG_IN":   2,
		"ANALOG_OUT":  3,
	}
)

func (x PinModeMessage_PortType) Enum() *PinModeMessage_PortType {
	p := new(PinModeMessage_PortType)
	*p = x
	return p
}

func (x PinModeMessage_PortType) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (PinModeMessage_PortType) Descriptor() protoreflect.EnumDescriptor {
	return file_remoteio_proto_enumTypes[0].Descriptor()
}

func (PinModeMessage_PortType) Type() protoreflect.EnumType {
	return &file_remoteio_proto_enumTypes[0]
}

func (x PinModeMessage_PortType) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use PinModeMessage_PortType.Descriptor instead.
func (PinModeMessage_PortType) EnumDescriptor() ([]byte, []int) {
	return file_remoteio_proto_rawDescGZIP(), []int{2, 0}
}

type DigitalState struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Pin   uint32 `protobuf:"varint,1,opt,name=pin,proto3" json:"pin,omitempty"`
	State bool   `protobuf:"varint,2,opt,name=state,proto3" json:"state,omitempty"`
}

func (x *DigitalState) Reset() {
	*x = DigitalState{}
	if protoimpl.UnsafeEnabled {
		mi := &file_remoteio_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *DigitalState) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*DigitalState) ProtoMessage() {}

func (x *DigitalState) ProtoReflect() protoreflect.Message {
	mi := &file_remoteio_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use DigitalState.ProtoReflect.Descriptor instead.
func (*DigitalState) Descriptor() ([]byte, []int) {
	return file_remoteio_proto_rawDescGZIP(), []int{0}
}

func (x *DigitalState) GetPin() uint32 {
	if x != nil {
		return x.Pin
	}
	return 0
}

func (x *DigitalState) GetState() bool {
	if x != nil {
		return x.State
	}
	return false
}

type AnalogState struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Pin   uint32 `protobuf:"varint,1,opt,name=pin,proto3" json:"pin,omitempty"`
	Value uint32 `protobuf:"varint,2,opt,name=value,proto3" json:"value,omitempty"`
}

func (x *AnalogState) Reset() {
	*x = AnalogState{}
	if protoimpl.UnsafeEnabled {
		mi := &file_remoteio_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *AnalogState) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*AnalogState) ProtoMessage() {}

func (x *AnalogState) ProtoReflect() protoreflect.Message {
	mi := &file_remoteio_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use AnalogState.ProtoReflect.Descriptor instead.
func (*AnalogState) Descriptor() ([]byte, []int) {
	return file_remoteio_proto_rawDescGZIP(), []int{1}
}

func (x *AnalogState) GetPin() uint32 {
	if x != nil {
		return x.Pin
	}
	return 0
}

func (x *AnalogState) GetValue() uint32 {
	if x != nil {
		return x.Value
	}
	return 0
}

type PinModeMessage struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Pin  uint32                  `protobuf:"varint,1,opt,name=pin,proto3" json:"pin,omitempty"`
	Mode PinModeMessage_PortType `protobuf:"varint,2,opt,name=mode,proto3,enum=remoteio.PinModeMessage_PortType" json:"mode,omitempty"`
}

func (x *PinModeMessage) Reset() {
	*x = PinModeMessage{}
	if protoimpl.UnsafeEnabled {
		mi := &file_remoteio_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *PinModeMessage) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*PinModeMessage) ProtoMessage() {}

func (x *PinModeMessage) ProtoReflect() protoreflect.Message {
	mi := &file_remoteio_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use PinModeMessage.ProtoReflect.Descriptor instead.
func (*PinModeMessage) Descriptor() ([]byte, []int) {
	return file_remoteio_proto_rawDescGZIP(), []int{2}
}

func (x *PinModeMessage) GetPin() uint32 {
	if x != nil {
		return x.Pin
	}
	return 0
}

func (x *PinModeMessage) GetMode() PinModeMessage_PortType {
	if x != nil {
		return x.Mode
	}
	return PinModeMessage_DIGITAL_IN
}

type SPIMessage struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Speed uint32   `protobuf:"varint,1,opt,name=speed,proto3" json:"speed,omitempty"`
	Cs    uint32   `protobuf:"varint,2,opt,name=cs,proto3" json:"cs,omitempty"`
	Mode  uint32   `protobuf:"varint,3,opt,name=mode,proto3" json:"mode,omitempty"`
	Bytes []uint32 `protobuf:"varint,4,rep,packed,name=bytes,proto3" json:"bytes,omitempty"`
}

func (x *SPIMessage) Reset() {
	*x = SPIMessage{}
	if protoimpl.UnsafeEnabled {
		mi := &file_remoteio_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *SPIMessage) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*SPIMessage) ProtoMessage() {}

func (x *SPIMessage) ProtoReflect() protoreflect.Message {
	mi := &file_remoteio_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use SPIMessage.ProtoReflect.Descriptor instead.
func (*SPIMessage) Descriptor() ([]byte, []int) {
	return file_remoteio_proto_rawDescGZIP(), []int{3}
}

func (x *SPIMessage) GetSpeed() uint32 {
	if x != nil {
		return x.Speed
	}
	return 0
}

func (x *SPIMessage) GetCs() uint32 {
	if x != nil {
		return x.Cs
	}
	return 0
}

func (x *SPIMessage) GetMode() uint32 {
	if x != nil {
		return x.Mode
	}
	return 0
}

func (x *SPIMessage) GetBytes() []uint32 {
	if x != nil {
		return x.Bytes
	}
	return nil
}

var File_remoteio_proto protoreflect.FileDescriptor

var file_remoteio_proto_rawDesc = []byte{
	0x0a, 0x0e, 0x72, 0x65, 0x6d, 0x6f, 0x74, 0x65, 0x69, 0x6f, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x12, 0x08, 0x72, 0x65, 0x6d, 0x6f, 0x74, 0x65, 0x69, 0x6f, 0x22, 0x36, 0x0a, 0x0c, 0x44, 0x69,
	0x67, 0x69, 0x74, 0x61, 0x6c, 0x53, 0x74, 0x61, 0x74, 0x65, 0x12, 0x10, 0x0a, 0x03, 0x70, 0x69,
	0x6e, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0d, 0x52, 0x03, 0x70, 0x69, 0x6e, 0x12, 0x14, 0x0a, 0x05,
	0x73, 0x74, 0x61, 0x74, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x08, 0x52, 0x05, 0x73, 0x74, 0x61,
	0x74, 0x65, 0x22, 0x35, 0x0a, 0x0b, 0x41, 0x6e, 0x61, 0x6c, 0x6f, 0x67, 0x53, 0x74, 0x61, 0x74,
	0x65, 0x12, 0x10, 0x0a, 0x03, 0x70, 0x69, 0x6e, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0d, 0x52, 0x03,
	0x70, 0x69, 0x6e, 0x12, 0x14, 0x0a, 0x05, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x18, 0x02, 0x20, 0x01,
	0x28, 0x0d, 0x52, 0x05, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x22, 0xa5, 0x01, 0x0a, 0x0e, 0x50, 0x69,
	0x6e, 0x4d, 0x6f, 0x64, 0x65, 0x4d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x12, 0x10, 0x0a, 0x03,
	0x70, 0x69, 0x6e, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0d, 0x52, 0x03, 0x70, 0x69, 0x6e, 0x12, 0x35,
	0x0a, 0x04, 0x6d, 0x6f, 0x64, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0e, 0x32, 0x21, 0x2e, 0x72,
	0x65, 0x6d, 0x6f, 0x74, 0x65, 0x69, 0x6f, 0x2e, 0x50, 0x69, 0x6e, 0x4d, 0x6f, 0x64, 0x65, 0x4d,
	0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x2e, 0x50, 0x6f, 0x72, 0x74, 0x54, 0x79, 0x70, 0x65, 0x52,
	0x04, 0x6d, 0x6f, 0x64, 0x65, 0x22, 0x4a, 0x0a, 0x08, 0x50, 0x6f, 0x72, 0x74, 0x54, 0x79, 0x70,
	0x65, 0x12, 0x0e, 0x0a, 0x0a, 0x44, 0x49, 0x47, 0x49, 0x54, 0x41, 0x4c, 0x5f, 0x49, 0x4e, 0x10,
	0x00, 0x12, 0x0f, 0x0a, 0x0b, 0x44, 0x49, 0x47, 0x49, 0x54, 0x41, 0x4c, 0x5f, 0x4f, 0x55, 0x54,
	0x10, 0x01, 0x12, 0x0d, 0x0a, 0x09, 0x41, 0x4e, 0x41, 0x4c, 0x4f, 0x47, 0x5f, 0x49, 0x4e, 0x10,
	0x02, 0x12, 0x0e, 0x0a, 0x0a, 0x41, 0x4e, 0x41, 0x4c, 0x4f, 0x47, 0x5f, 0x4f, 0x55, 0x54, 0x10,
	0x03, 0x22, 0x5c, 0x0a, 0x0a, 0x53, 0x50, 0x49, 0x4d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x12,
	0x14, 0x0a, 0x05, 0x73, 0x70, 0x65, 0x65, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0d, 0x52, 0x05,
	0x73, 0x70, 0x65, 0x65, 0x64, 0x12, 0x0e, 0x0a, 0x02, 0x63, 0x73, 0x18, 0x02, 0x20, 0x01, 0x28,
	0x0d, 0x52, 0x02, 0x63, 0x73, 0x12, 0x12, 0x0a, 0x04, 0x6d, 0x6f, 0x64, 0x65, 0x18, 0x03, 0x20,
	0x01, 0x28, 0x0d, 0x52, 0x04, 0x6d, 0x6f, 0x64, 0x65, 0x12, 0x14, 0x0a, 0x05, 0x62, 0x79, 0x74,
	0x65, 0x73, 0x18, 0x04, 0x20, 0x03, 0x28, 0x0d, 0x52, 0x05, 0x62, 0x79, 0x74, 0x65, 0x73, 0x32,
	0xbe, 0x03, 0x0a, 0x08, 0x52, 0x65, 0x6d, 0x6f, 0x74, 0x65, 0x49, 0x4f, 0x12, 0x3f, 0x0a, 0x07,
	0x70, 0x69, 0x6e, 0x4d, 0x6f, 0x64, 0x65, 0x12, 0x18, 0x2e, 0x72, 0x65, 0x6d, 0x6f, 0x74, 0x65,
	0x69, 0x6f, 0x2e, 0x50, 0x69, 0x6e, 0x4d, 0x6f, 0x64, 0x65, 0x4d, 0x65, 0x73, 0x73, 0x61, 0x67,
	0x65, 0x1a, 0x18, 0x2e, 0x72, 0x65, 0x6d, 0x6f, 0x74, 0x65, 0x69, 0x6f, 0x2e, 0x50, 0x69, 0x6e,
	0x4d, 0x6f, 0x64, 0x65, 0x4d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x22, 0x00, 0x12, 0x3f, 0x0a,
	0x0b, 0x64, 0x69, 0x67, 0x69, 0x74, 0x61, 0x6c, 0x52, 0x65, 0x61, 0x64, 0x12, 0x16, 0x2e, 0x72,
	0x65, 0x6d, 0x6f, 0x74, 0x65, 0x69, 0x6f, 0x2e, 0x44, 0x69, 0x67, 0x69, 0x74, 0x61, 0x6c, 0x53,
	0x74, 0x61, 0x74, 0x65, 0x1a, 0x16, 0x2e, 0x72, 0x65, 0x6d, 0x6f, 0x74, 0x65, 0x69, 0x6f, 0x2e,
	0x44, 0x69, 0x67, 0x69, 0x74, 0x61, 0x6c, 0x53, 0x74, 0x61, 0x74, 0x65, 0x22, 0x00, 0x12, 0x40,
	0x0a, 0x0c, 0x64, 0x69, 0x67, 0x69, 0x74, 0x61, 0x6c, 0x57, 0x72, 0x69, 0x74, 0x65, 0x12, 0x16,
	0x2e, 0x72, 0x65, 0x6d, 0x6f, 0x74, 0x65, 0x69, 0x6f, 0x2e, 0x44, 0x69, 0x67, 0x69, 0x74, 0x61,
	0x6c, 0x53, 0x74, 0x61, 0x74, 0x65, 0x1a, 0x16, 0x2e, 0x72, 0x65, 0x6d, 0x6f, 0x74, 0x65, 0x69,
	0x6f, 0x2e, 0x44, 0x69, 0x67, 0x69, 0x74, 0x61, 0x6c, 0x53, 0x74, 0x61, 0x74, 0x65, 0x22, 0x00,
	0x12, 0x3c, 0x0a, 0x0a, 0x61, 0x6e, 0x61, 0x6c, 0x6f, 0x67, 0x52, 0x65, 0x61, 0x64, 0x12, 0x15,
	0x2e, 0x72, 0x65, 0x6d, 0x6f, 0x74, 0x65, 0x69, 0x6f, 0x2e, 0x41, 0x6e, 0x61, 0x6c, 0x6f, 0x67,
	0x53, 0x74, 0x61, 0x74, 0x65, 0x1a, 0x15, 0x2e, 0x72, 0x65, 0x6d, 0x6f, 0x74, 0x65, 0x69, 0x6f,
	0x2e, 0x41, 0x6e, 0x61, 0x6c, 0x6f, 0x67, 0x53, 0x74, 0x61, 0x74, 0x65, 0x22, 0x00, 0x12, 0x3d,
	0x0a, 0x0b, 0x61, 0x6e, 0x61, 0x6c, 0x6f, 0x67, 0x57, 0x72, 0x69, 0x74, 0x65, 0x12, 0x15, 0x2e,
	0x72, 0x65, 0x6d, 0x6f, 0x74, 0x65, 0x69, 0x6f, 0x2e, 0x41, 0x6e, 0x61, 0x6c, 0x6f, 0x67, 0x53,
	0x74, 0x61, 0x74, 0x65, 0x1a, 0x15, 0x2e, 0x72, 0x65, 0x6d, 0x6f, 0x74, 0x65, 0x69, 0x6f, 0x2e,
	0x41, 0x6e, 0x61, 0x6c, 0x6f, 0x67, 0x53, 0x74, 0x61, 0x74, 0x65, 0x22, 0x00, 0x12, 0x37, 0x0a,
	0x07, 0x73, 0x70, 0x69, 0x52, 0x65, 0x61, 0x64, 0x12, 0x14, 0x2e, 0x72, 0x65, 0x6d, 0x6f, 0x74,
	0x65, 0x69, 0x6f, 0x2e, 0x53, 0x50, 0x49, 0x4d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x1a, 0x14,
	0x2e, 0x72, 0x65, 0x6d, 0x6f, 0x74, 0x65, 0x69, 0x6f, 0x2e, 0x53, 0x50, 0x49, 0x4d, 0x65, 0x73,
	0x73, 0x61, 0x67, 0x65, 0x22, 0x00, 0x12, 0x38, 0x0a, 0x08, 0x73, 0x70, 0x69, 0x57, 0x72, 0x69,
	0x74, 0x65, 0x12, 0x14, 0x2e, 0x72, 0x65, 0x6d, 0x6f, 0x74, 0x65, 0x69, 0x6f, 0x2e, 0x53, 0x50,
	0x49, 0x4d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x1a, 0x14, 0x2e, 0x72, 0x65, 0x6d, 0x6f, 0x74,
	0x65, 0x69, 0x6f, 0x2e, 0x53, 0x50, 0x49, 0x4d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x22, 0x00,
	0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_remoteio_proto_rawDescOnce sync.Once
	file_remoteio_proto_rawDescData = file_remoteio_proto_rawDesc
)

func file_remoteio_proto_rawDescGZIP() []byte {
	file_remoteio_proto_rawDescOnce.Do(func() {
		file_remoteio_proto_rawDescData = protoimpl.X.CompressGZIP(file_remoteio_proto_rawDescData)
	})
	return file_remoteio_proto_rawDescData
}

var file_remoteio_proto_enumTypes = make([]protoimpl.EnumInfo, 1)
var file_remoteio_proto_msgTypes = make([]protoimpl.MessageInfo, 4)
var file_remoteio_proto_goTypes = []interface{}{
	(PinModeMessage_PortType)(0), // 0: remoteio.PinModeMessage.PortType
	(*DigitalState)(nil),         // 1: remoteio.DigitalState
	(*AnalogState)(nil),          // 2: remoteio.AnalogState
	(*PinModeMessage)(nil),       // 3: remoteio.PinModeMessage
	(*SPIMessage)(nil),           // 4: remoteio.SPIMessage
}
var file_remoteio_proto_depIdxs = []int32{
	0, // 0: remoteio.PinModeMessage.mode:type_name -> remoteio.PinModeMessage.PortType
	3, // 1: remoteio.RemoteIO.pinMode:input_type -> remoteio.PinModeMessage
	1, // 2: remoteio.RemoteIO.digitalRead:input_type -> remoteio.DigitalState
	1, // 3: remoteio.RemoteIO.digitalWrite:input_type -> remoteio.DigitalState
	2, // 4: remoteio.RemoteIO.analogRead:input_type -> remoteio.AnalogState
	2, // 5: remoteio.RemoteIO.analogWrite:input_type -> remoteio.AnalogState
	4, // 6: remoteio.RemoteIO.spiRead:input_type -> remoteio.SPIMessage
	4, // 7: remoteio.RemoteIO.spiWrite:input_type -> remoteio.SPIMessage
	3, // 8: remoteio.RemoteIO.pinMode:output_type -> remoteio.PinModeMessage
	1, // 9: remoteio.RemoteIO.digitalRead:output_type -> remoteio.DigitalState
	1, // 10: remoteio.RemoteIO.digitalWrite:output_type -> remoteio.DigitalState
	2, // 11: remoteio.RemoteIO.analogRead:output_type -> remoteio.AnalogState
	2, // 12: remoteio.RemoteIO.analogWrite:output_type -> remoteio.AnalogState
	4, // 13: remoteio.RemoteIO.spiRead:output_type -> remoteio.SPIMessage
	4, // 14: remoteio.RemoteIO.spiWrite:output_type -> remoteio.SPIMessage
	8, // [8:15] is the sub-list for method output_type
	1, // [1:8] is the sub-list for method input_type
	1, // [1:1] is the sub-list for extension type_name
	1, // [1:1] is the sub-list for extension extendee
	0, // [0:1] is the sub-list for field type_name
}

func init() { file_remoteio_proto_init() }
func file_remoteio_proto_init() {
	if File_remoteio_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_remoteio_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*DigitalState); i {
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
		file_remoteio_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*AnalogState); i {
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
		file_remoteio_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*PinModeMessage); i {
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
		file_remoteio_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*SPIMessage); i {
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
			RawDescriptor: file_remoteio_proto_rawDesc,
			NumEnums:      1,
			NumMessages:   4,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_remoteio_proto_goTypes,
		DependencyIndexes: file_remoteio_proto_depIdxs,
		EnumInfos:         file_remoteio_proto_enumTypes,
		MessageInfos:      file_remoteio_proto_msgTypes,
	}.Build()
	File_remoteio_proto = out.File
	file_remoteio_proto_rawDesc = nil
	file_remoteio_proto_goTypes = nil
	file_remoteio_proto_depIdxs = nil
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConnInterface

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion6

// RemoteIOClient is the client API for RemoteIO service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type RemoteIOClient interface {
	PinMode(ctx context.Context, in *PinModeMessage, opts ...grpc.CallOption) (*PinModeMessage, error)
	DigitalRead(ctx context.Context, in *DigitalState, opts ...grpc.CallOption) (*DigitalState, error)
	DigitalWrite(ctx context.Context, in *DigitalState, opts ...grpc.CallOption) (*DigitalState, error)
	AnalogRead(ctx context.Context, in *AnalogState, opts ...grpc.CallOption) (*AnalogState, error)
	AnalogWrite(ctx context.Context, in *AnalogState, opts ...grpc.CallOption) (*AnalogState, error)
	SpiRead(ctx context.Context, in *SPIMessage, opts ...grpc.CallOption) (*SPIMessage, error)
	SpiWrite(ctx context.Context, in *SPIMessage, opts ...grpc.CallOption) (*SPIMessage, error)
}

type remoteIOClient struct {
	cc grpc.ClientConnInterface
}

func NewRemoteIOClient(cc grpc.ClientConnInterface) RemoteIOClient {
	return &remoteIOClient{cc}
}

func (c *remoteIOClient) PinMode(ctx context.Context, in *PinModeMessage, opts ...grpc.CallOption) (*PinModeMessage, error) {
	out := new(PinModeMessage)
	err := c.cc.Invoke(ctx, "/remoteio.RemoteIO/pinMode", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *remoteIOClient) DigitalRead(ctx context.Context, in *DigitalState, opts ...grpc.CallOption) (*DigitalState, error) {
	out := new(DigitalState)
	err := c.cc.Invoke(ctx, "/remoteio.RemoteIO/digitalRead", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *remoteIOClient) DigitalWrite(ctx context.Context, in *DigitalState, opts ...grpc.CallOption) (*DigitalState, error) {
	out := new(DigitalState)
	err := c.cc.Invoke(ctx, "/remoteio.RemoteIO/digitalWrite", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *remoteIOClient) AnalogRead(ctx context.Context, in *AnalogState, opts ...grpc.CallOption) (*AnalogState, error) {
	out := new(AnalogState)
	err := c.cc.Invoke(ctx, "/remoteio.RemoteIO/analogRead", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *remoteIOClient) AnalogWrite(ctx context.Context, in *AnalogState, opts ...grpc.CallOption) (*AnalogState, error) {
	out := new(AnalogState)
	err := c.cc.Invoke(ctx, "/remoteio.RemoteIO/analogWrite", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *remoteIOClient) SpiRead(ctx context.Context, in *SPIMessage, opts ...grpc.CallOption) (*SPIMessage, error) {
	out := new(SPIMessage)
	err := c.cc.Invoke(ctx, "/remoteio.RemoteIO/spiRead", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *remoteIOClient) SpiWrite(ctx context.Context, in *SPIMessage, opts ...grpc.CallOption) (*SPIMessage, error) {
	out := new(SPIMessage)
	err := c.cc.Invoke(ctx, "/remoteio.RemoteIO/spiWrite", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// RemoteIOServer is the server API for RemoteIO service.
type RemoteIOServer interface {
	PinMode(context.Context, *PinModeMessage) (*PinModeMessage, error)
	DigitalRead(context.Context, *DigitalState) (*DigitalState, error)
	DigitalWrite(context.Context, *DigitalState) (*DigitalState, error)
	AnalogRead(context.Context, *AnalogState) (*AnalogState, error)
	AnalogWrite(context.Context, *AnalogState) (*AnalogState, error)
	SpiRead(context.Context, *SPIMessage) (*SPIMessage, error)
	SpiWrite(context.Context, *SPIMessage) (*SPIMessage, error)
}

// UnimplementedRemoteIOServer can be embedded to have forward compatible implementations.
type UnimplementedRemoteIOServer struct {
}

func (*UnimplementedRemoteIOServer) PinMode(context.Context, *PinModeMessage) (*PinModeMessage, error) {
	return nil, status.Errorf(codes.Unimplemented, "method PinMode not implemented")
}
func (*UnimplementedRemoteIOServer) DigitalRead(context.Context, *DigitalState) (*DigitalState, error) {
	return nil, status.Errorf(codes.Unimplemented, "method DigitalRead not implemented")
}
func (*UnimplementedRemoteIOServer) DigitalWrite(context.Context, *DigitalState) (*DigitalState, error) {
	return nil, status.Errorf(codes.Unimplemented, "method DigitalWrite not implemented")
}
func (*UnimplementedRemoteIOServer) AnalogRead(context.Context, *AnalogState) (*AnalogState, error) {
	return nil, status.Errorf(codes.Unimplemented, "method AnalogRead not implemented")
}
func (*UnimplementedRemoteIOServer) AnalogWrite(context.Context, *AnalogState) (*AnalogState, error) {
	return nil, status.Errorf(codes.Unimplemented, "method AnalogWrite not implemented")
}
func (*UnimplementedRemoteIOServer) SpiRead(context.Context, *SPIMessage) (*SPIMessage, error) {
	return nil, status.Errorf(codes.Unimplemented, "method SpiRead not implemented")
}
func (*UnimplementedRemoteIOServer) SpiWrite(context.Context, *SPIMessage) (*SPIMessage, error) {
	return nil, status.Errorf(codes.Unimplemented, "method SpiWrite not implemented")
}

func RegisterRemoteIOServer(s *grpc.Server, srv RemoteIOServer) {
	s.RegisterService(&_RemoteIO_serviceDesc, srv)
}

func _RemoteIO_PinMode_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(PinModeMessage)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(RemoteIOServer).PinMode(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/remoteio.RemoteIO/PinMode",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(RemoteIOServer).PinMode(ctx, req.(*PinModeMessage))
	}
	return interceptor(ctx, in, info, handler)
}

func _RemoteIO_DigitalRead_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(DigitalState)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(RemoteIOServer).DigitalRead(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/remoteio.RemoteIO/DigitalRead",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(RemoteIOServer).DigitalRead(ctx, req.(*DigitalState))
	}
	return interceptor(ctx, in, info, handler)
}

func _RemoteIO_DigitalWrite_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(DigitalState)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(RemoteIOServer).DigitalWrite(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/remoteio.RemoteIO/DigitalWrite",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(RemoteIOServer).DigitalWrite(ctx, req.(*DigitalState))
	}
	return interceptor(ctx, in, info, handler)
}

func _RemoteIO_AnalogRead_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(AnalogState)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(RemoteIOServer).AnalogRead(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/remoteio.RemoteIO/AnalogRead",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(RemoteIOServer).AnalogRead(ctx, req.(*AnalogState))
	}
	return interceptor(ctx, in, info, handler)
}

func _RemoteIO_AnalogWrite_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(AnalogState)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(RemoteIOServer).AnalogWrite(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/remoteio.RemoteIO/AnalogWrite",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(RemoteIOServer).AnalogWrite(ctx, req.(*AnalogState))
	}
	return interceptor(ctx, in, info, handler)
}

func _RemoteIO_SpiRead_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(SPIMessage)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(RemoteIOServer).SpiRead(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/remoteio.RemoteIO/SpiRead",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(RemoteIOServer).SpiRead(ctx, req.(*SPIMessage))
	}
	return interceptor(ctx, in, info, handler)
}

func _RemoteIO_SpiWrite_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(SPIMessage)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(RemoteIOServer).SpiWrite(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/remoteio.RemoteIO/SpiWrite",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(RemoteIOServer).SpiWrite(ctx, req.(*SPIMessage))
	}
	return interceptor(ctx, in, info, handler)
}

var _RemoteIO_serviceDesc = grpc.ServiceDesc{
	ServiceName: "remoteio.RemoteIO",
	HandlerType: (*RemoteIOServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "pinMode",
			Handler:    _RemoteIO_PinMode_Handler,
		},
		{
			MethodName: "digitalRead",
			Handler:    _RemoteIO_DigitalRead_Handler,
		},
		{
			MethodName: "digitalWrite",
			Handler:    _RemoteIO_DigitalWrite_Handler,
		},
		{
			MethodName: "analogRead",
			Handler:    _RemoteIO_AnalogRead_Handler,
		},
		{
			MethodName: "analogWrite",
			Handler:    _RemoteIO_AnalogWrite_Handler,
		},
		{
			MethodName: "spiRead",
			Handler:    _RemoteIO_SpiRead_Handler,
		},
		{
			MethodName: "spiWrite",
			Handler:    _RemoteIO_SpiWrite_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "remoteio.proto",
}
