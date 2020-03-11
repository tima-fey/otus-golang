// Code generated by protoc-gen-go. DO NOT EDIT.
// source: scheme.proto

package scheme

import (
	context "context"
	fmt "fmt"
	proto "github.com/golang/protobuf/proto"
	timestamp "github.com/golang/protobuf/ptypes/timestamp"
	grpc "google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
	math "math"
)

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// This is a compile-time assertion to ensure that this generated file
// is compatible with the proto package it is being compiled against.
// A compilation error at this line likely means your copy of the
// proto package needs to be updated.
const _ = proto.ProtoPackageIsVersion3 // please upgrade the proto package

type EventId struct {
	Id                   int32    `protobuf:"varint,1,opt,name=id,proto3" json:"id,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *EventId) Reset()         { *m = EventId{} }
func (m *EventId) String() string { return proto.CompactTextString(m) }
func (*EventId) ProtoMessage()    {}
func (*EventId) Descriptor() ([]byte, []int) {
	return fileDescriptor_8bf3faeb8468324e, []int{0}
}

func (m *EventId) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_EventId.Unmarshal(m, b)
}
func (m *EventId) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_EventId.Marshal(b, m, deterministic)
}
func (m *EventId) XXX_Merge(src proto.Message) {
	xxx_messageInfo_EventId.Merge(m, src)
}
func (m *EventId) XXX_Size() int {
	return xxx_messageInfo_EventId.Size(m)
}
func (m *EventId) XXX_DiscardUnknown() {
	xxx_messageInfo_EventId.DiscardUnknown(m)
}

var xxx_messageInfo_EventId proto.InternalMessageInfo

func (m *EventId) GetId() int32 {
	if m != nil {
		return m.Id
	}
	return 0
}

type Event struct {
	Id                   int32                `protobuf:"varint,1,opt,name=id,proto3" json:"id,omitempty"`
	Name                 string               `protobuf:"bytes,2,opt,name=name,proto3" json:"name,omitempty"`
	StarTtime            *timestamp.Timestamp `protobuf:"bytes,3,opt,name=starTtime,proto3" json:"starTtime,omitempty"`
	EndTime              *timestamp.Timestamp `protobuf:"bytes,4,opt,name=endTime,proto3" json:"endTime,omitempty"`
	XXX_NoUnkeyedLiteral struct{}             `json:"-"`
	XXX_unrecognized     []byte               `json:"-"`
	XXX_sizecache        int32                `json:"-"`
}

func (m *Event) Reset()         { *m = Event{} }
func (m *Event) String() string { return proto.CompactTextString(m) }
func (*Event) ProtoMessage()    {}
func (*Event) Descriptor() ([]byte, []int) {
	return fileDescriptor_8bf3faeb8468324e, []int{1}
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

func (m *Event) GetId() int32 {
	if m != nil {
		return m.Id
	}
	return 0
}

func (m *Event) GetName() string {
	if m != nil {
		return m.Name
	}
	return ""
}

func (m *Event) GetStarTtime() *timestamp.Timestamp {
	if m != nil {
		return m.StarTtime
	}
	return nil
}

func (m *Event) GetEndTime() *timestamp.Timestamp {
	if m != nil {
		return m.EndTime
	}
	return nil
}

func init() {
	proto.RegisterType((*EventId)(nil), "scheme.EventId")
	proto.RegisterType((*Event)(nil), "scheme.Event")
}

func init() { proto.RegisterFile("scheme.proto", fileDescriptor_8bf3faeb8468324e) }

var fileDescriptor_8bf3faeb8468324e = []byte{
	// 247 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x84, 0x90, 0xcf, 0x4a, 0xc3, 0x40,
	0x10, 0xc6, 0xd9, 0xb4, 0x4d, 0xe9, 0x68, 0x15, 0xe6, 0x14, 0x73, 0x31, 0x14, 0x94, 0xe0, 0x61,
	0x0b, 0xd5, 0x83, 0x57, 0x41, 0xd1, 0x5e, 0x43, 0x7d, 0x80, 0xad, 0x33, 0xd6, 0x40, 0xfe, 0x91,
	0x1d, 0x7d, 0x19, 0x9f, 0xc2, 0x37, 0x94, 0x6c, 0x0c, 0x92, 0xf5, 0xe0, 0x6d, 0xe7, 0xfb, 0x7e,
	0xf3, 0x7d, 0xbb, 0x0b, 0xc7, 0xf6, 0xe5, 0x8d, 0x4b, 0xd6, 0x4d, 0x5b, 0x4b, 0x8d, 0x61, 0x3f,
	0xc5, 0xe7, 0x87, 0xba, 0x3e, 0x14, 0xbc, 0x76, 0xea, 0xfe, 0xfd, 0x75, 0x2d, 0x79, 0xc9, 0x56,
	0x4c, 0xd9, 0xf4, 0xe0, 0xea, 0x0c, 0xe6, 0x0f, 0x1f, 0x5c, 0xc9, 0x96, 0xf0, 0x04, 0x82, 0x9c,
	0x22, 0x95, 0xa8, 0x74, 0x96, 0x05, 0x39, 0xad, 0x3e, 0x15, 0xcc, 0x9c, 0xe7, 0x3b, 0x88, 0x30,
	0xad, 0x4c, 0xc9, 0x51, 0x90, 0xa8, 0x74, 0x91, 0xb9, 0x33, 0xde, 0xc2, 0xc2, 0x8a, 0x69, 0x77,
	0x5d, 0x41, 0x34, 0x49, 0x54, 0x7a, 0xb4, 0x89, 0x75, 0xdf, 0xae, 0x87, 0x76, 0xbd, 0x1b, 0xda,
	0xb3, 0x5f, 0x18, 0x6f, 0x60, 0xce, 0x15, 0x75, 0x56, 0x34, 0xfd, 0x77, 0x6f, 0x40, 0x37, 0x5f,
	0x0a, 0x96, 0xee, 0x76, 0xf6, 0xc9, 0x54, 0x54, 0x70, 0x8b, 0x17, 0x30, 0xb9, 0x23, 0xc2, 0xa5,
	0xfe, 0xf9, 0x09, 0xe7, 0xc6, 0xa7, 0xa3, 0x71, 0x4b, 0x1d, 0xf6, 0xc8, 0x82, 0xbe, 0x1e, 0x8f,
	0xf7, 0xf0, 0x12, 0xc2, 0xe7, 0x86, 0x8c, 0xb0, 0x1f, 0xe8, 0x71, 0x57, 0x10, 0xde, 0x73, 0xc1,
	0xc2, 0x7f, 0x13, 0x7d, 0x61, 0x1f, 0xba, 0x07, 0x5d, 0x7f, 0x07, 0x00, 0x00, 0xff, 0xff, 0xcf,
	0xd2, 0xed, 0x06, 0xac, 0x01, 0x00, 0x00,
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConnInterface

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion6

// EventsHandlerClient is the client API for EventsHandler service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type EventsHandlerClient interface {
	Add(ctx context.Context, in *Event, opts ...grpc.CallOption) (*EventId, error)
	Get(ctx context.Context, in *EventId, opts ...grpc.CallOption) (*Event, error)
	Update(ctx context.Context, in *Event, opts ...grpc.CallOption) (*Event, error)
	Delete(ctx context.Context, in *EventId, opts ...grpc.CallOption) (*EventId, error)
}

type eventsHandlerClient struct {
	cc grpc.ClientConnInterface
}

func NewEventsHandlerClient(cc grpc.ClientConnInterface) EventsHandlerClient {
	return &eventsHandlerClient{cc}
}

func (c *eventsHandlerClient) Add(ctx context.Context, in *Event, opts ...grpc.CallOption) (*EventId, error) {
	out := new(EventId)
	err := c.cc.Invoke(ctx, "/scheme.EventsHandler/Add", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *eventsHandlerClient) Get(ctx context.Context, in *EventId, opts ...grpc.CallOption) (*Event, error) {
	out := new(Event)
	err := c.cc.Invoke(ctx, "/scheme.EventsHandler/Get", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *eventsHandlerClient) Update(ctx context.Context, in *Event, opts ...grpc.CallOption) (*Event, error) {
	out := new(Event)
	err := c.cc.Invoke(ctx, "/scheme.EventsHandler/Update", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *eventsHandlerClient) Delete(ctx context.Context, in *EventId, opts ...grpc.CallOption) (*EventId, error) {
	out := new(EventId)
	err := c.cc.Invoke(ctx, "/scheme.EventsHandler/Delete", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// EventsHandlerServer is the server API for EventsHandler service.
type EventsHandlerServer interface {
	Add(context.Context, *Event) (*EventId, error)
	Get(context.Context, *EventId) (*Event, error)
	Update(context.Context, *Event) (*Event, error)
	Delete(context.Context, *EventId) (*EventId, error)
}

// UnimplementedEventsHandlerServer can be embedded to have forward compatible implementations.
type UnimplementedEventsHandlerServer struct {
}

func (*UnimplementedEventsHandlerServer) Add(ctx context.Context, req *Event) (*EventId, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Add not implemented")
}
func (*UnimplementedEventsHandlerServer) Get(ctx context.Context, req *EventId) (*Event, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Get not implemented")
}
func (*UnimplementedEventsHandlerServer) Update(ctx context.Context, req *Event) (*Event, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Update not implemented")
}
func (*UnimplementedEventsHandlerServer) Delete(ctx context.Context, req *EventId) (*EventId, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Delete not implemented")
}

func RegisterEventsHandlerServer(s *grpc.Server, srv EventsHandlerServer) {
	s.RegisterService(&_EventsHandler_serviceDesc, srv)
}

func _EventsHandler_Add_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(Event)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(EventsHandlerServer).Add(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/scheme.EventsHandler/Add",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(EventsHandlerServer).Add(ctx, req.(*Event))
	}
	return interceptor(ctx, in, info, handler)
}

func _EventsHandler_Get_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(EventId)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(EventsHandlerServer).Get(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/scheme.EventsHandler/Get",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(EventsHandlerServer).Get(ctx, req.(*EventId))
	}
	return interceptor(ctx, in, info, handler)
}

func _EventsHandler_Update_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(Event)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(EventsHandlerServer).Update(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/scheme.EventsHandler/Update",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(EventsHandlerServer).Update(ctx, req.(*Event))
	}
	return interceptor(ctx, in, info, handler)
}

func _EventsHandler_Delete_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(EventId)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(EventsHandlerServer).Delete(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/scheme.EventsHandler/Delete",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(EventsHandlerServer).Delete(ctx, req.(*EventId))
	}
	return interceptor(ctx, in, info, handler)
}

var _EventsHandler_serviceDesc = grpc.ServiceDesc{
	ServiceName: "scheme.EventsHandler",
	HandlerType: (*EventsHandlerServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "Add",
			Handler:    _EventsHandler_Add_Handler,
		},
		{
			MethodName: "Get",
			Handler:    _EventsHandler_Get_Handler,
		},
		{
			MethodName: "Update",
			Handler:    _EventsHandler_Update_Handler,
		},
		{
			MethodName: "Delete",
			Handler:    _EventsHandler_Delete_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "scheme.proto",
}
