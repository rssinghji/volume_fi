// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.29.0
// 	protoc        v3.21.12
// source: flight.proto

package pb

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

type Request struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	FlightPath []string `protobuf:"bytes,1,rep,name=flight_path,json=flightPath,proto3" json:"flight_path,omitempty"`
}

func (x *Request) Reset() {
	*x = Request{}
	if protoimpl.UnsafeEnabled {
		mi := &file_flight_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Request) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Request) ProtoMessage() {}

func (x *Request) ProtoReflect() protoreflect.Message {
	mi := &file_flight_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Request.ProtoReflect.Descriptor instead.
func (*Request) Descriptor() ([]byte, []int) {
	return file_flight_proto_rawDescGZIP(), []int{0}
}

func (x *Request) GetFlightPath() []string {
	if x != nil {
		return x.FlightPath
	}
	return nil
}

type Response struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	FlightPath []string `protobuf:"bytes,1,rep,name=flight_path,json=flightPath,proto3" json:"flight_path,omitempty"`
}

func (x *Response) Reset() {
	*x = Response{}
	if protoimpl.UnsafeEnabled {
		mi := &file_flight_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Response) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Response) ProtoMessage() {}

func (x *Response) ProtoReflect() protoreflect.Message {
	mi := &file_flight_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Response.ProtoReflect.Descriptor instead.
func (*Response) Descriptor() ([]byte, []int) {
	return file_flight_proto_rawDescGZIP(), []int{1}
}

func (x *Response) GetFlightPath() []string {
	if x != nil {
		return x.FlightPath
	}
	return nil
}

type CalculatePathRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Request []*Request `protobuf:"bytes,1,rep,name=request,proto3" json:"request,omitempty"`
}

func (x *CalculatePathRequest) Reset() {
	*x = CalculatePathRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_flight_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *CalculatePathRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CalculatePathRequest) ProtoMessage() {}

func (x *CalculatePathRequest) ProtoReflect() protoreflect.Message {
	mi := &file_flight_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CalculatePathRequest.ProtoReflect.Descriptor instead.
func (*CalculatePathRequest) Descriptor() ([]byte, []int) {
	return file_flight_proto_rawDescGZIP(), []int{2}
}

func (x *CalculatePathRequest) GetRequest() []*Request {
	if x != nil {
		return x.Request
	}
	return nil
}

type CalculatePathResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Response *Response `protobuf:"bytes,1,opt,name=response,proto3" json:"response,omitempty"`
}

func (x *CalculatePathResponse) Reset() {
	*x = CalculatePathResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_flight_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *CalculatePathResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CalculatePathResponse) ProtoMessage() {}

func (x *CalculatePathResponse) ProtoReflect() protoreflect.Message {
	mi := &file_flight_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CalculatePathResponse.ProtoReflect.Descriptor instead.
func (*CalculatePathResponse) Descriptor() ([]byte, []int) {
	return file_flight_proto_rawDescGZIP(), []int{3}
}

func (x *CalculatePathResponse) GetResponse() *Response {
	if x != nil {
		return x.Response
	}
	return nil
}

var File_flight_proto protoreflect.FileDescriptor

var file_flight_proto_rawDesc = []byte{
	0x0a, 0x0c, 0x66, 0x6c, 0x69, 0x67, 0x68, 0x74, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0x2a,
	0x0a, 0x07, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x1f, 0x0a, 0x0b, 0x66, 0x6c, 0x69,
	0x67, 0x68, 0x74, 0x5f, 0x70, 0x61, 0x74, 0x68, 0x18, 0x01, 0x20, 0x03, 0x28, 0x09, 0x52, 0x0a,
	0x66, 0x6c, 0x69, 0x67, 0x68, 0x74, 0x50, 0x61, 0x74, 0x68, 0x22, 0x2b, 0x0a, 0x08, 0x52, 0x65,
	0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x1f, 0x0a, 0x0b, 0x66, 0x6c, 0x69, 0x67, 0x68, 0x74,
	0x5f, 0x70, 0x61, 0x74, 0x68, 0x18, 0x01, 0x20, 0x03, 0x28, 0x09, 0x52, 0x0a, 0x66, 0x6c, 0x69,
	0x67, 0x68, 0x74, 0x50, 0x61, 0x74, 0x68, 0x22, 0x3a, 0x0a, 0x14, 0x43, 0x61, 0x6c, 0x63, 0x75,
	0x6c, 0x61, 0x74, 0x65, 0x50, 0x61, 0x74, 0x68, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12,
	0x22, 0x0a, 0x07, 0x72, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x18, 0x01, 0x20, 0x03, 0x28, 0x0b,
	0x32, 0x08, 0x2e, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x52, 0x07, 0x72, 0x65, 0x71, 0x75,
	0x65, 0x73, 0x74, 0x22, 0x3e, 0x0a, 0x15, 0x43, 0x61, 0x6c, 0x63, 0x75, 0x6c, 0x61, 0x74, 0x65,
	0x50, 0x61, 0x74, 0x68, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x25, 0x0a, 0x08,
	0x72, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x09,
	0x2e, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x52, 0x08, 0x72, 0x65, 0x73, 0x70, 0x6f,
	0x6e, 0x73, 0x65, 0x32, 0x46, 0x0a, 0x06, 0x46, 0x6c, 0x69, 0x67, 0x68, 0x74, 0x12, 0x3c, 0x0a,
	0x09, 0x43, 0x61, 0x6c, 0x63, 0x75, 0x6c, 0x61, 0x74, 0x65, 0x12, 0x15, 0x2e, 0x43, 0x61, 0x6c,
	0x63, 0x75, 0x6c, 0x61, 0x74, 0x65, 0x50, 0x61, 0x74, 0x68, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73,
	0x74, 0x1a, 0x16, 0x2e, 0x43, 0x61, 0x6c, 0x63, 0x75, 0x6c, 0x61, 0x74, 0x65, 0x50, 0x61, 0x74,
	0x68, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x00, 0x42, 0x05, 0x5a, 0x03, 0x70,
	0x62, 0x2f, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_flight_proto_rawDescOnce sync.Once
	file_flight_proto_rawDescData = file_flight_proto_rawDesc
)

func file_flight_proto_rawDescGZIP() []byte {
	file_flight_proto_rawDescOnce.Do(func() {
		file_flight_proto_rawDescData = protoimpl.X.CompressGZIP(file_flight_proto_rawDescData)
	})
	return file_flight_proto_rawDescData
}

var file_flight_proto_msgTypes = make([]protoimpl.MessageInfo, 4)
var file_flight_proto_goTypes = []interface{}{
	(*Request)(nil),               // 0: Request
	(*Response)(nil),              // 1: Response
	(*CalculatePathRequest)(nil),  // 2: CalculatePathRequest
	(*CalculatePathResponse)(nil), // 3: CalculatePathResponse
}
var file_flight_proto_depIdxs = []int32{
	0, // 0: CalculatePathRequest.request:type_name -> Request
	1, // 1: CalculatePathResponse.response:type_name -> Response
	2, // 2: Flight.Calculate:input_type -> CalculatePathRequest
	3, // 3: Flight.Calculate:output_type -> CalculatePathResponse
	3, // [3:4] is the sub-list for method output_type
	2, // [2:3] is the sub-list for method input_type
	2, // [2:2] is the sub-list for extension type_name
	2, // [2:2] is the sub-list for extension extendee
	0, // [0:2] is the sub-list for field type_name
}

func init() { file_flight_proto_init() }
func file_flight_proto_init() {
	if File_flight_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_flight_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Request); i {
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
		file_flight_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Response); i {
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
		file_flight_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*CalculatePathRequest); i {
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
		file_flight_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*CalculatePathResponse); i {
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
			RawDescriptor: file_flight_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   4,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_flight_proto_goTypes,
		DependencyIndexes: file_flight_proto_depIdxs,
		MessageInfos:      file_flight_proto_msgTypes,
	}.Build()
	File_flight_proto = out.File
	file_flight_proto_rawDesc = nil
	file_flight_proto_goTypes = nil
	file_flight_proto_depIdxs = nil
}
