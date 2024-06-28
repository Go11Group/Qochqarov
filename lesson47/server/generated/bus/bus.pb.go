// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.28.1
// 	protoc        v3.12.4
// source: bus.proto

package bus

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

type BusRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	BusNumber string `protobuf:"bytes,1,opt,name=busNumber,proto3" json:"busNumber,omitempty"`
}

func (x *BusRequest) Reset() {
	*x = BusRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_bus_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *BusRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*BusRequest) ProtoMessage() {}

func (x *BusRequest) ProtoReflect() protoreflect.Message {
	mi := &file_bus_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use BusRequest.ProtoReflect.Descriptor instead.
func (*BusRequest) Descriptor() ([]byte, []int) {
	return file_bus_proto_rawDescGZIP(), []int{0}
}

func (x *BusRequest) GetBusNumber() string {
	if x != nil {
		return x.BusNumber
	}
	return ""
}

type BusScheduleResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Schule []string `protobuf:"bytes,1,rep,name=Schule,proto3" json:"Schule,omitempty"`
}

func (x *BusScheduleResponse) Reset() {
	*x = BusScheduleResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_bus_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *BusScheduleResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*BusScheduleResponse) ProtoMessage() {}

func (x *BusScheduleResponse) ProtoReflect() protoreflect.Message {
	mi := &file_bus_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use BusScheduleResponse.ProtoReflect.Descriptor instead.
func (*BusScheduleResponse) Descriptor() ([]byte, []int) {
	return file_bus_proto_rawDescGZIP(), []int{1}
}

func (x *BusScheduleResponse) GetSchule() []string {
	if x != nil {
		return x.Schule
	}
	return nil
}

type BusLocationResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	BusNumber string `protobuf:"bytes,1,opt,name=BusNumber,proto3" json:"BusNumber,omitempty"`
	Locatsion string `protobuf:"bytes,2,opt,name=Locatsion,proto3" json:"Locatsion,omitempty"`
}

func (x *BusLocationResponse) Reset() {
	*x = BusLocationResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_bus_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *BusLocationResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*BusLocationResponse) ProtoMessage() {}

func (x *BusLocationResponse) ProtoReflect() protoreflect.Message {
	mi := &file_bus_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use BusLocationResponse.ProtoReflect.Descriptor instead.
func (*BusLocationResponse) Descriptor() ([]byte, []int) {
	return file_bus_proto_rawDescGZIP(), []int{2}
}

func (x *BusLocationResponse) GetBusNumber() string {
	if x != nil {
		return x.BusNumber
	}
	return ""
}

func (x *BusLocationResponse) GetLocatsion() string {
	if x != nil {
		return x.Locatsion
	}
	return ""
}

type TrafficJamReportRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Locatsion   string `protobuf:"bytes,1,opt,name=Locatsion,proto3" json:"Locatsion,omitempty"`
	Discription string `protobuf:"bytes,2,opt,name=Discription,proto3" json:"Discription,omitempty"`
}

func (x *TrafficJamReportRequest) Reset() {
	*x = TrafficJamReportRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_bus_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *TrafficJamReportRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*TrafficJamReportRequest) ProtoMessage() {}

func (x *TrafficJamReportRequest) ProtoReflect() protoreflect.Message {
	mi := &file_bus_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use TrafficJamReportRequest.ProtoReflect.Descriptor instead.
func (*TrafficJamReportRequest) Descriptor() ([]byte, []int) {
	return file_bus_proto_rawDescGZIP(), []int{3}
}

func (x *TrafficJamReportRequest) GetLocatsion() string {
	if x != nil {
		return x.Locatsion
	}
	return ""
}

func (x *TrafficJamReportRequest) GetDiscription() string {
	if x != nil {
		return x.Discription
	}
	return ""
}

type TrafficJamReportResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Condition string `protobuf:"bytes,1,opt,name=Condition,proto3" json:"Condition,omitempty"`
}

func (x *TrafficJamReportResponse) Reset() {
	*x = TrafficJamReportResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_bus_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *TrafficJamReportResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*TrafficJamReportResponse) ProtoMessage() {}

func (x *TrafficJamReportResponse) ProtoReflect() protoreflect.Message {
	mi := &file_bus_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use TrafficJamReportResponse.ProtoReflect.Descriptor instead.
func (*TrafficJamReportResponse) Descriptor() ([]byte, []int) {
	return file_bus_proto_rawDescGZIP(), []int{4}
}

func (x *TrafficJamReportResponse) GetCondition() string {
	if x != nil {
		return x.Condition
	}
	return ""
}

var File_bus_proto protoreflect.FileDescriptor

var file_bus_proto_rawDesc = []byte{
	0x0a, 0x09, 0x62, 0x75, 0x73, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x06, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x73, 0x22, 0x2a, 0x0a, 0x0a, 0x42, 0x75, 0x73, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73,
	0x74, 0x12, 0x1c, 0x0a, 0x09, 0x62, 0x75, 0x73, 0x4e, 0x75, 0x6d, 0x62, 0x65, 0x72, 0x18, 0x01,
	0x20, 0x01, 0x28, 0x09, 0x52, 0x09, 0x62, 0x75, 0x73, 0x4e, 0x75, 0x6d, 0x62, 0x65, 0x72, 0x22,
	0x2d, 0x0a, 0x13, 0x42, 0x75, 0x73, 0x53, 0x63, 0x68, 0x65, 0x64, 0x75, 0x6c, 0x65, 0x52, 0x65,
	0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x16, 0x0a, 0x06, 0x53, 0x63, 0x68, 0x75, 0x6c, 0x65,
	0x18, 0x01, 0x20, 0x03, 0x28, 0x09, 0x52, 0x06, 0x53, 0x63, 0x68, 0x75, 0x6c, 0x65, 0x22, 0x51,
	0x0a, 0x13, 0x42, 0x75, 0x73, 0x4c, 0x6f, 0x63, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x52, 0x65, 0x73,
	0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x1c, 0x0a, 0x09, 0x42, 0x75, 0x73, 0x4e, 0x75, 0x6d, 0x62,
	0x65, 0x72, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x09, 0x42, 0x75, 0x73, 0x4e, 0x75, 0x6d,
	0x62, 0x65, 0x72, 0x12, 0x1c, 0x0a, 0x09, 0x4c, 0x6f, 0x63, 0x61, 0x74, 0x73, 0x69, 0x6f, 0x6e,
	0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x09, 0x4c, 0x6f, 0x63, 0x61, 0x74, 0x73, 0x69, 0x6f,
	0x6e, 0x22, 0x59, 0x0a, 0x17, 0x54, 0x72, 0x61, 0x66, 0x66, 0x69, 0x63, 0x4a, 0x61, 0x6d, 0x52,
	0x65, 0x70, 0x6f, 0x72, 0x74, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x1c, 0x0a, 0x09,
	0x4c, 0x6f, 0x63, 0x61, 0x74, 0x73, 0x69, 0x6f, 0x6e, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52,
	0x09, 0x4c, 0x6f, 0x63, 0x61, 0x74, 0x73, 0x69, 0x6f, 0x6e, 0x12, 0x20, 0x0a, 0x0b, 0x44, 0x69,
	0x73, 0x63, 0x72, 0x69, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52,
	0x0b, 0x44, 0x69, 0x73, 0x63, 0x72, 0x69, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x22, 0x38, 0x0a, 0x18,
	0x54, 0x72, 0x61, 0x66, 0x66, 0x69, 0x63, 0x4a, 0x61, 0x6d, 0x52, 0x65, 0x70, 0x6f, 0x72, 0x74,
	0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x1c, 0x0a, 0x09, 0x43, 0x6f, 0x6e, 0x64,
	0x69, 0x74, 0x69, 0x6f, 0x6e, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x09, 0x43, 0x6f, 0x6e,
	0x64, 0x69, 0x74, 0x69, 0x6f, 0x6e, 0x32, 0xf1, 0x01, 0x0a, 0x10, 0x54, 0x72, 0x61, 0x6e, 0x73,
	0x70, 0x6f, 0x72, 0x74, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x12, 0x41, 0x0a, 0x0e, 0x47,
	0x65, 0x74, 0x42, 0x75, 0x73, 0x53, 0x63, 0x68, 0x65, 0x64, 0x75, 0x6c, 0x65, 0x12, 0x12, 0x2e,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x73, 0x2e, 0x42, 0x75, 0x73, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73,
	0x74, 0x1a, 0x1b, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x73, 0x2e, 0x42, 0x75, 0x73, 0x53, 0x63,
	0x68, 0x65, 0x64, 0x75, 0x6c, 0x65, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x43,
	0x0a, 0x10, 0x54, 0x72, 0x61, 0x63, 0x6b, 0x42, 0x75, 0x73, 0x4c, 0x6f, 0x63, 0x61, 0x74, 0x69,
	0x6f, 0x6e, 0x12, 0x12, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x73, 0x2e, 0x42, 0x75, 0x73, 0x52,
	0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x1b, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x73, 0x2e,
	0x42, 0x75, 0x73, 0x4c, 0x6f, 0x63, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x52, 0x65, 0x73, 0x70, 0x6f,
	0x6e, 0x73, 0x65, 0x12, 0x55, 0x0a, 0x10, 0x52, 0x65, 0x70, 0x6f, 0x72, 0x74, 0x54, 0x72, 0x61,
	0x66, 0x66, 0x69, 0x63, 0x4a, 0x61, 0x6d, 0x12, 0x1f, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x73,
	0x2e, 0x54, 0x72, 0x61, 0x66, 0x66, 0x69, 0x63, 0x4a, 0x61, 0x6d, 0x52, 0x65, 0x70, 0x6f, 0x72,
	0x74, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x20, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x73, 0x2e, 0x54, 0x72, 0x61, 0x66, 0x66, 0x69, 0x63, 0x4a, 0x61, 0x6d, 0x52, 0x65, 0x70, 0x6f,
	0x72, 0x74, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x42, 0x11, 0x5a, 0x0f, 0x2e, 0x2f,
	0x67, 0x65, 0x6e, 0x65, 0x72, 0x61, 0x74, 0x65, 0x64, 0x2f, 0x62, 0x75, 0x73, 0x62, 0x06, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_bus_proto_rawDescOnce sync.Once
	file_bus_proto_rawDescData = file_bus_proto_rawDesc
)

func file_bus_proto_rawDescGZIP() []byte {
	file_bus_proto_rawDescOnce.Do(func() {
		file_bus_proto_rawDescData = protoimpl.X.CompressGZIP(file_bus_proto_rawDescData)
	})
	return file_bus_proto_rawDescData
}

var file_bus_proto_msgTypes = make([]protoimpl.MessageInfo, 5)
var file_bus_proto_goTypes = []interface{}{
	(*BusRequest)(nil),               // 0: protos.BusRequest
	(*BusScheduleResponse)(nil),      // 1: protos.BusScheduleResponse
	(*BusLocationResponse)(nil),      // 2: protos.BusLocationResponse
	(*TrafficJamReportRequest)(nil),  // 3: protos.TrafficJamReportRequest
	(*TrafficJamReportResponse)(nil), // 4: protos.TrafficJamReportResponse
}
var file_bus_proto_depIdxs = []int32{
	0, // 0: protos.TransportService.GetBusSchedule:input_type -> protos.BusRequest
	0, // 1: protos.TransportService.TrackBusLocation:input_type -> protos.BusRequest
	3, // 2: protos.TransportService.ReportTrafficJam:input_type -> protos.TrafficJamReportRequest
	1, // 3: protos.TransportService.GetBusSchedule:output_type -> protos.BusScheduleResponse
	2, // 4: protos.TransportService.TrackBusLocation:output_type -> protos.BusLocationResponse
	4, // 5: protos.TransportService.ReportTrafficJam:output_type -> protos.TrafficJamReportResponse
	3, // [3:6] is the sub-list for method output_type
	0, // [0:3] is the sub-list for method input_type
	0, // [0:0] is the sub-list for extension type_name
	0, // [0:0] is the sub-list for extension extendee
	0, // [0:0] is the sub-list for field type_name
}

func init() { file_bus_proto_init() }
func file_bus_proto_init() {
	if File_bus_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_bus_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*BusRequest); i {
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
		file_bus_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*BusScheduleResponse); i {
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
		file_bus_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*BusLocationResponse); i {
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
		file_bus_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*TrafficJamReportRequest); i {
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
		file_bus_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*TrafficJamReportResponse); i {
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
			RawDescriptor: file_bus_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   5,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_bus_proto_goTypes,
		DependencyIndexes: file_bus_proto_depIdxs,
		MessageInfos:      file_bus_proto_msgTypes,
	}.Build()
	File_bus_proto = out.File
	file_bus_proto_rawDesc = nil
	file_bus_proto_goTypes = nil
	file_bus_proto_depIdxs = nil
}