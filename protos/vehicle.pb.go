// protos/v1/user/user.proto

// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.26.0
// 	protoc        v3.15.8
// source: protos/vehicle.proto

package vehicle

import (
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

type VehicleMessage struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	VehicleId           string `protobuf:"bytes,1,opt,name=vehicle_id,json=vehicleId,proto3" json:"vehicle_id,omitempty"`
	VehicleName         string `protobuf:"bytes,2,opt,name=vehicle_name,json=vehicleName,proto3" json:"vehicle_name,omitempty"`
	VehicleNumber       string `protobuf:"bytes,3,opt,name=vehicle_number,json=vehicleNumber,proto3" json:"vehicle_number,omitempty"`
	VehicleVinNumber    string `protobuf:"bytes,4,opt,name=vehicle_vin_number,json=vehicleVinNumber,proto3" json:"vehicle_vin_number,omitempty"`
	VehicleSerialNumber string `protobuf:"bytes,5,opt,name=vehicle_serial_number,json=vehicleSerialNumber,proto3" json:"vehicle_serial_number,omitempty"`
}

func (x *VehicleMessage) Reset() {
	*x = VehicleMessage{}
	if protoimpl.UnsafeEnabled {
		mi := &file_protos_vehicle_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *VehicleMessage) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*VehicleMessage) ProtoMessage() {}

func (x *VehicleMessage) ProtoReflect() protoreflect.Message {
	mi := &file_protos_vehicle_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use VehicleMessage.ProtoReflect.Descriptor instead.
func (*VehicleMessage) Descriptor() ([]byte, []int) {
	return file_protos_vehicle_proto_rawDescGZIP(), []int{0}
}

func (x *VehicleMessage) GetVehicleId() string {
	if x != nil {
		return x.VehicleId
	}
	return ""
}

func (x *VehicleMessage) GetVehicleName() string {
	if x != nil {
		return x.VehicleName
	}
	return ""
}

func (x *VehicleMessage) GetVehicleNumber() string {
	if x != nil {
		return x.VehicleNumber
	}
	return ""
}

func (x *VehicleMessage) GetVehicleVinNumber() string {
	if x != nil {
		return x.VehicleVinNumber
	}
	return ""
}

func (x *VehicleMessage) GetVehicleSerialNumber() string {
	if x != nil {
		return x.VehicleSerialNumber
	}
	return ""
}

type GetVehicleRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	VehicleId string `protobuf:"bytes,1,opt,name=vehicle_id,json=vehicleId,proto3" json:"vehicle_id,omitempty"`
}

func (x *GetVehicleRequest) Reset() {
	*x = GetVehicleRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_protos_vehicle_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetVehicleRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetVehicleRequest) ProtoMessage() {}

func (x *GetVehicleRequest) ProtoReflect() protoreflect.Message {
	mi := &file_protos_vehicle_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetVehicleRequest.ProtoReflect.Descriptor instead.
func (*GetVehicleRequest) Descriptor() ([]byte, []int) {
	return file_protos_vehicle_proto_rawDescGZIP(), []int{1}
}

func (x *GetVehicleRequest) GetVehicleId() string {
	if x != nil {
		return x.VehicleId
	}
	return ""
}

type GetVehicleResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	VehicleMessage *VehicleMessage `protobuf:"bytes,1,opt,name=vehicle_message,json=vehicleMessage,proto3" json:"vehicle_message,omitempty"`
}

func (x *GetVehicleResponse) Reset() {
	*x = GetVehicleResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_protos_vehicle_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetVehicleResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetVehicleResponse) ProtoMessage() {}

func (x *GetVehicleResponse) ProtoReflect() protoreflect.Message {
	mi := &file_protos_vehicle_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetVehicleResponse.ProtoReflect.Descriptor instead.
func (*GetVehicleResponse) Descriptor() ([]byte, []int) {
	return file_protos_vehicle_proto_rawDescGZIP(), []int{2}
}

func (x *GetVehicleResponse) GetVehicleMessage() *VehicleMessage {
	if x != nil {
		return x.VehicleMessage
	}
	return nil
}

type ListVehiclesRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields
}

func (x *ListVehiclesRequest) Reset() {
	*x = ListVehiclesRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_protos_vehicle_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ListVehiclesRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ListVehiclesRequest) ProtoMessage() {}

func (x *ListVehiclesRequest) ProtoReflect() protoreflect.Message {
	mi := &file_protos_vehicle_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ListVehiclesRequest.ProtoReflect.Descriptor instead.
func (*ListVehiclesRequest) Descriptor() ([]byte, []int) {
	return file_protos_vehicle_proto_rawDescGZIP(), []int{3}
}

type ListVehiclesResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	VehicleMessages []*VehicleMessage `protobuf:"bytes,1,rep,name=vehicle_messages,json=vehicleMessages,proto3" json:"vehicle_messages,omitempty"`
}

func (x *ListVehiclesResponse) Reset() {
	*x = ListVehiclesResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_protos_vehicle_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ListVehiclesResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ListVehiclesResponse) ProtoMessage() {}

func (x *ListVehiclesResponse) ProtoReflect() protoreflect.Message {
	mi := &file_protos_vehicle_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ListVehiclesResponse.ProtoReflect.Descriptor instead.
func (*ListVehiclesResponse) Descriptor() ([]byte, []int) {
	return file_protos_vehicle_proto_rawDescGZIP(), []int{4}
}

func (x *ListVehiclesResponse) GetVehicleMessages() []*VehicleMessage {
	if x != nil {
		return x.VehicleMessages
	}
	return nil
}

var File_protos_vehicle_proto protoreflect.FileDescriptor

var file_protos_vehicle_proto_rawDesc = []byte{
	0x0a, 0x14, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x73, 0x2f, 0x76, 0x65, 0x68, 0x69, 0x63, 0x6c, 0x65,
	0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x0e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x73, 0x2e, 0x76,
	0x65, 0x68, 0x69, 0x63, 0x6c, 0x65, 0x22, 0xdb, 0x01, 0x0a, 0x0e, 0x56, 0x65, 0x68, 0x69, 0x63,
	0x6c, 0x65, 0x4d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x12, 0x1d, 0x0a, 0x0a, 0x76, 0x65, 0x68,
	0x69, 0x63, 0x6c, 0x65, 0x5f, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x09, 0x76,
	0x65, 0x68, 0x69, 0x63, 0x6c, 0x65, 0x49, 0x64, 0x12, 0x21, 0x0a, 0x0c, 0x76, 0x65, 0x68, 0x69,
	0x63, 0x6c, 0x65, 0x5f, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0b,
	0x76, 0x65, 0x68, 0x69, 0x63, 0x6c, 0x65, 0x4e, 0x61, 0x6d, 0x65, 0x12, 0x25, 0x0a, 0x0e, 0x76,
	0x65, 0x68, 0x69, 0x63, 0x6c, 0x65, 0x5f, 0x6e, 0x75, 0x6d, 0x62, 0x65, 0x72, 0x18, 0x03, 0x20,
	0x01, 0x28, 0x09, 0x52, 0x0d, 0x76, 0x65, 0x68, 0x69, 0x63, 0x6c, 0x65, 0x4e, 0x75, 0x6d, 0x62,
	0x65, 0x72, 0x12, 0x2c, 0x0a, 0x12, 0x76, 0x65, 0x68, 0x69, 0x63, 0x6c, 0x65, 0x5f, 0x76, 0x69,
	0x6e, 0x5f, 0x6e, 0x75, 0x6d, 0x62, 0x65, 0x72, 0x18, 0x04, 0x20, 0x01, 0x28, 0x09, 0x52, 0x10,
	0x76, 0x65, 0x68, 0x69, 0x63, 0x6c, 0x65, 0x56, 0x69, 0x6e, 0x4e, 0x75, 0x6d, 0x62, 0x65, 0x72,
	0x12, 0x32, 0x0a, 0x15, 0x76, 0x65, 0x68, 0x69, 0x63, 0x6c, 0x65, 0x5f, 0x73, 0x65, 0x72, 0x69,
	0x61, 0x6c, 0x5f, 0x6e, 0x75, 0x6d, 0x62, 0x65, 0x72, 0x18, 0x05, 0x20, 0x01, 0x28, 0x09, 0x52,
	0x13, 0x76, 0x65, 0x68, 0x69, 0x63, 0x6c, 0x65, 0x53, 0x65, 0x72, 0x69, 0x61, 0x6c, 0x4e, 0x75,
	0x6d, 0x62, 0x65, 0x72, 0x22, 0x32, 0x0a, 0x11, 0x47, 0x65, 0x74, 0x56, 0x65, 0x68, 0x69, 0x63,
	0x6c, 0x65, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x1d, 0x0a, 0x0a, 0x76, 0x65, 0x68,
	0x69, 0x63, 0x6c, 0x65, 0x5f, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x09, 0x76,
	0x65, 0x68, 0x69, 0x63, 0x6c, 0x65, 0x49, 0x64, 0x22, 0x5d, 0x0a, 0x12, 0x47, 0x65, 0x74, 0x56,
	0x65, 0x68, 0x69, 0x63, 0x6c, 0x65, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x47,
	0x0a, 0x0f, 0x76, 0x65, 0x68, 0x69, 0x63, 0x6c, 0x65, 0x5f, 0x6d, 0x65, 0x73, 0x73, 0x61, 0x67,
	0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1e, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x73,
	0x2e, 0x76, 0x65, 0x68, 0x69, 0x63, 0x6c, 0x65, 0x2e, 0x56, 0x65, 0x68, 0x69, 0x63, 0x6c, 0x65,
	0x4d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x52, 0x0e, 0x76, 0x65, 0x68, 0x69, 0x63, 0x6c, 0x65,
	0x4d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x22, 0x15, 0x0a, 0x13, 0x4c, 0x69, 0x73, 0x74, 0x56,
	0x65, 0x68, 0x69, 0x63, 0x6c, 0x65, 0x73, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x22, 0x61,
	0x0a, 0x14, 0x4c, 0x69, 0x73, 0x74, 0x56, 0x65, 0x68, 0x69, 0x63, 0x6c, 0x65, 0x73, 0x52, 0x65,
	0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x49, 0x0a, 0x10, 0x76, 0x65, 0x68, 0x69, 0x63, 0x6c,
	0x65, 0x5f, 0x6d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x73, 0x18, 0x01, 0x20, 0x03, 0x28, 0x0b,
	0x32, 0x1e, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x73, 0x2e, 0x76, 0x65, 0x68, 0x69, 0x63, 0x6c,
	0x65, 0x2e, 0x56, 0x65, 0x68, 0x69, 0x63, 0x6c, 0x65, 0x4d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65,
	0x52, 0x0f, 0x76, 0x65, 0x68, 0x69, 0x63, 0x6c, 0x65, 0x4d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65,
	0x73, 0x32, 0xb9, 0x01, 0x0a, 0x07, 0x56, 0x65, 0x68, 0x69, 0x63, 0x6c, 0x65, 0x12, 0x53, 0x0a,
	0x0a, 0x47, 0x65, 0x74, 0x56, 0x65, 0x68, 0x69, 0x63, 0x6c, 0x65, 0x12, 0x21, 0x2e, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x73, 0x2e, 0x76, 0x65, 0x68, 0x69, 0x63, 0x6c, 0x65, 0x2e, 0x47, 0x65, 0x74,
	0x56, 0x65, 0x68, 0x69, 0x63, 0x6c, 0x65, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x22,
	0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x73, 0x2e, 0x76, 0x65, 0x68, 0x69, 0x63, 0x6c, 0x65, 0x2e,
	0x47, 0x65, 0x74, 0x56, 0x65, 0x68, 0x69, 0x63, 0x6c, 0x65, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e,
	0x73, 0x65, 0x12, 0x59, 0x0a, 0x0c, 0x4c, 0x69, 0x73, 0x74, 0x56, 0x65, 0x68, 0x69, 0x63, 0x6c,
	0x65, 0x73, 0x12, 0x23, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x73, 0x2e, 0x76, 0x65, 0x68, 0x69,
	0x63, 0x6c, 0x65, 0x2e, 0x4c, 0x69, 0x73, 0x74, 0x56, 0x65, 0x68, 0x69, 0x63, 0x6c, 0x65, 0x73,
	0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x24, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x73,
	0x2e, 0x76, 0x65, 0x68, 0x69, 0x63, 0x6c, 0x65, 0x2e, 0x4c, 0x69, 0x73, 0x74, 0x56, 0x65, 0x68,
	0x69, 0x63, 0x6c, 0x65, 0x73, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x42, 0x10, 0x5a,
	0x0e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x73, 0x2f, 0x76, 0x65, 0x68, 0x69, 0x63, 0x6c, 0x65, 0x62,
	0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_protos_vehicle_proto_rawDescOnce sync.Once
	file_protos_vehicle_proto_rawDescData = file_protos_vehicle_proto_rawDesc
)

func file_protos_vehicle_proto_rawDescGZIP() []byte {
	file_protos_vehicle_proto_rawDescOnce.Do(func() {
		file_protos_vehicle_proto_rawDescData = protoimpl.X.CompressGZIP(file_protos_vehicle_proto_rawDescData)
	})
	return file_protos_vehicle_proto_rawDescData
}

var file_protos_vehicle_proto_msgTypes = make([]protoimpl.MessageInfo, 5)
var file_protos_vehicle_proto_goTypes = []interface{}{
	(*VehicleMessage)(nil),       // 0: protos.vehicle.VehicleMessage
	(*GetVehicleRequest)(nil),    // 1: protos.vehicle.GetVehicleRequest
	(*GetVehicleResponse)(nil),   // 2: protos.vehicle.GetVehicleResponse
	(*ListVehiclesRequest)(nil),  // 3: protos.vehicle.ListVehiclesRequest
	(*ListVehiclesResponse)(nil), // 4: protos.vehicle.ListVehiclesResponse
}
var file_protos_vehicle_proto_depIdxs = []int32{
	0, // 0: protos.vehicle.GetVehicleResponse.vehicle_message:type_name -> protos.vehicle.VehicleMessage
	0, // 1: protos.vehicle.ListVehiclesResponse.vehicle_messages:type_name -> protos.vehicle.VehicleMessage
	1, // 2: protos.vehicle.Vehicle.GetVehicle:input_type -> protos.vehicle.GetVehicleRequest
	3, // 3: protos.vehicle.Vehicle.ListVehicles:input_type -> protos.vehicle.ListVehiclesRequest
	2, // 4: protos.vehicle.Vehicle.GetVehicle:output_type -> protos.vehicle.GetVehicleResponse
	4, // 5: protos.vehicle.Vehicle.ListVehicles:output_type -> protos.vehicle.ListVehiclesResponse
	4, // [4:6] is the sub-list for method output_type
	2, // [2:4] is the sub-list for method input_type
	2, // [2:2] is the sub-list for extension type_name
	2, // [2:2] is the sub-list for extension extendee
	0, // [0:2] is the sub-list for field type_name
}

func init() { file_protos_vehicle_proto_init() }
func file_protos_vehicle_proto_init() {
	if File_protos_vehicle_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_protos_vehicle_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*VehicleMessage); i {
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
		file_protos_vehicle_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetVehicleRequest); i {
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
		file_protos_vehicle_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetVehicleResponse); i {
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
		file_protos_vehicle_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ListVehiclesRequest); i {
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
		file_protos_vehicle_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ListVehiclesResponse); i {
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
			RawDescriptor: file_protos_vehicle_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   5,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_protos_vehicle_proto_goTypes,
		DependencyIndexes: file_protos_vehicle_proto_depIdxs,
		MessageInfos:      file_protos_vehicle_proto_msgTypes,
	}.Build()
	File_protos_vehicle_proto = out.File
	file_protos_vehicle_proto_rawDesc = nil
	file_protos_vehicle_proto_goTypes = nil
	file_protos_vehicle_proto_depIdxs = nil
}