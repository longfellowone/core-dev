// Code generated by protoc-gen-go. DO NOT EDIT.
// source: todo.proto

package Todo

import (
	context "context"
	fmt "fmt"
	proto "github.com/golang/protobuf/proto"
	grpc "google.golang.org/grpc"
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
const _ = proto.ProtoPackageIsVersion2 // please upgrade the proto package

type Empty struct {
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *Empty) Reset()         { *m = Empty{} }
func (m *Empty) String() string { return proto.CompactTextString(m) }
func (*Empty) ProtoMessage()    {}
func (*Empty) Descriptor() ([]byte, []int) {
	return fileDescriptor_0e4b95d0c4e09639, []int{0}
}

func (m *Empty) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Empty.Unmarshal(m, b)
}
func (m *Empty) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Empty.Marshal(b, m, deterministic)
}
func (m *Empty) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Empty.Merge(m, src)
}
func (m *Empty) XXX_Size() int {
	return xxx_messageInfo_Empty.Size(m)
}
func (m *Empty) XXX_DiscardUnknown() {
	xxx_messageInfo_Empty.DiscardUnknown(m)
}

var xxx_messageInfo_Empty proto.InternalMessageInfo

type Task struct {
	Uuid                 string   `protobuf:"bytes,1,opt,name=uuid,proto3" json:"uuid,omitempty"`
	Message              string   `protobuf:"bytes,2,opt,name=message,proto3" json:"message,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *Task) Reset()         { *m = Task{} }
func (m *Task) String() string { return proto.CompactTextString(m) }
func (*Task) ProtoMessage()    {}
func (*Task) Descriptor() ([]byte, []int) {
	return fileDescriptor_0e4b95d0c4e09639, []int{1}
}

func (m *Task) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Task.Unmarshal(m, b)
}
func (m *Task) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Task.Marshal(b, m, deterministic)
}
func (m *Task) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Task.Merge(m, src)
}
func (m *Task) XXX_Size() int {
	return xxx_messageInfo_Task.Size(m)
}
func (m *Task) XXX_DiscardUnknown() {
	xxx_messageInfo_Task.DiscardUnknown(m)
}

var xxx_messageInfo_Task proto.InternalMessageInfo

func (m *Task) GetUuid() string {
	if m != nil {
		return m.Uuid
	}
	return ""
}

func (m *Task) GetMessage() string {
	if m != nil {
		return m.Message
	}
	return ""
}

type TaskResponse struct {
	Tasks                []*Task  `protobuf:"bytes,1,rep,name=tasks,proto3" json:"tasks,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *TaskResponse) Reset()         { *m = TaskResponse{} }
func (m *TaskResponse) String() string { return proto.CompactTextString(m) }
func (*TaskResponse) ProtoMessage()    {}
func (*TaskResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_0e4b95d0c4e09639, []int{2}
}

func (m *TaskResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_TaskResponse.Unmarshal(m, b)
}
func (m *TaskResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_TaskResponse.Marshal(b, m, deterministic)
}
func (m *TaskResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_TaskResponse.Merge(m, src)
}
func (m *TaskResponse) XXX_Size() int {
	return xxx_messageInfo_TaskResponse.Size(m)
}
func (m *TaskResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_TaskResponse.DiscardUnknown(m)
}

var xxx_messageInfo_TaskResponse proto.InternalMessageInfo

func (m *TaskResponse) GetTasks() []*Task {
	if m != nil {
		return m.Tasks
	}
	return nil
}

type RemoveTaskRequest struct {
	Uuid                 string   `protobuf:"bytes,1,opt,name=uuid,proto3" json:"uuid,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *RemoveTaskRequest) Reset()         { *m = RemoveTaskRequest{} }
func (m *RemoveTaskRequest) String() string { return proto.CompactTextString(m) }
func (*RemoveTaskRequest) ProtoMessage()    {}
func (*RemoveTaskRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_0e4b95d0c4e09639, []int{3}
}

func (m *RemoveTaskRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_RemoveTaskRequest.Unmarshal(m, b)
}
func (m *RemoveTaskRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_RemoveTaskRequest.Marshal(b, m, deterministic)
}
func (m *RemoveTaskRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_RemoveTaskRequest.Merge(m, src)
}
func (m *RemoveTaskRequest) XXX_Size() int {
	return xxx_messageInfo_RemoveTaskRequest.Size(m)
}
func (m *RemoveTaskRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_RemoveTaskRequest.DiscardUnknown(m)
}

var xxx_messageInfo_RemoveTaskRequest proto.InternalMessageInfo

func (m *RemoveTaskRequest) GetUuid() string {
	if m != nil {
		return m.Uuid
	}
	return ""
}

type Product struct {
	Uuid                 string   `protobuf:"bytes,1,opt,name=uuid,proto3" json:"uuid,omitempty"`
	Name                 string   `protobuf:"bytes,2,opt,name=name,proto3" json:"name,omitempty"`
	Indexes              []*Index `protobuf:"bytes,3,rep,name=indexes,proto3" json:"indexes,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *Product) Reset()         { *m = Product{} }
func (m *Product) String() string { return proto.CompactTextString(m) }
func (*Product) ProtoMessage()    {}
func (*Product) Descriptor() ([]byte, []int) {
	return fileDescriptor_0e4b95d0c4e09639, []int{4}
}

func (m *Product) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Product.Unmarshal(m, b)
}
func (m *Product) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Product.Marshal(b, m, deterministic)
}
func (m *Product) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Product.Merge(m, src)
}
func (m *Product) XXX_Size() int {
	return xxx_messageInfo_Product.Size(m)
}
func (m *Product) XXX_DiscardUnknown() {
	xxx_messageInfo_Product.DiscardUnknown(m)
}

var xxx_messageInfo_Product proto.InternalMessageInfo

func (m *Product) GetUuid() string {
	if m != nil {
		return m.Uuid
	}
	return ""
}

func (m *Product) GetName() string {
	if m != nil {
		return m.Name
	}
	return ""
}

func (m *Product) GetIndexes() []*Index {
	if m != nil {
		return m.Indexes
	}
	return nil
}

type Index struct {
	Index                string   `protobuf:"bytes,1,opt,name=index,proto3" json:"index,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *Index) Reset()         { *m = Index{} }
func (m *Index) String() string { return proto.CompactTextString(m) }
func (*Index) ProtoMessage()    {}
func (*Index) Descriptor() ([]byte, []int) {
	return fileDescriptor_0e4b95d0c4e09639, []int{5}
}

func (m *Index) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Index.Unmarshal(m, b)
}
func (m *Index) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Index.Marshal(b, m, deterministic)
}
func (m *Index) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Index.Merge(m, src)
}
func (m *Index) XXX_Size() int {
	return xxx_messageInfo_Index.Size(m)
}
func (m *Index) XXX_DiscardUnknown() {
	xxx_messageInfo_Index.DiscardUnknown(m)
}

var xxx_messageInfo_Index proto.InternalMessageInfo

func (m *Index) GetIndex() string {
	if m != nil {
		return m.Index
	}
	return ""
}

type FindProductRequest struct {
	Name                 string   `protobuf:"bytes,1,opt,name=name,proto3" json:"name,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *FindProductRequest) Reset()         { *m = FindProductRequest{} }
func (m *FindProductRequest) String() string { return proto.CompactTextString(m) }
func (*FindProductRequest) ProtoMessage()    {}
func (*FindProductRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_0e4b95d0c4e09639, []int{6}
}

func (m *FindProductRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_FindProductRequest.Unmarshal(m, b)
}
func (m *FindProductRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_FindProductRequest.Marshal(b, m, deterministic)
}
func (m *FindProductRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_FindProductRequest.Merge(m, src)
}
func (m *FindProductRequest) XXX_Size() int {
	return xxx_messageInfo_FindProductRequest.Size(m)
}
func (m *FindProductRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_FindProductRequest.DiscardUnknown(m)
}

var xxx_messageInfo_FindProductRequest proto.InternalMessageInfo

func (m *FindProductRequest) GetName() string {
	if m != nil {
		return m.Name
	}
	return ""
}

type FindProductResponse struct {
	Products             []*Product `protobuf:"bytes,1,rep,name=products,proto3" json:"products,omitempty"`
	XXX_NoUnkeyedLiteral struct{}   `json:"-"`
	XXX_unrecognized     []byte     `json:"-"`
	XXX_sizecache        int32      `json:"-"`
}

func (m *FindProductResponse) Reset()         { *m = FindProductResponse{} }
func (m *FindProductResponse) String() string { return proto.CompactTextString(m) }
func (*FindProductResponse) ProtoMessage()    {}
func (*FindProductResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_0e4b95d0c4e09639, []int{7}
}

func (m *FindProductResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_FindProductResponse.Unmarshal(m, b)
}
func (m *FindProductResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_FindProductResponse.Marshal(b, m, deterministic)
}
func (m *FindProductResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_FindProductResponse.Merge(m, src)
}
func (m *FindProductResponse) XXX_Size() int {
	return xxx_messageInfo_FindProductResponse.Size(m)
}
func (m *FindProductResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_FindProductResponse.DiscardUnknown(m)
}

var xxx_messageInfo_FindProductResponse proto.InternalMessageInfo

func (m *FindProductResponse) GetProducts() []*Product {
	if m != nil {
		return m.Products
	}
	return nil
}

func init() {
	proto.RegisterType((*Empty)(nil), "Todo.Empty")
	proto.RegisterType((*Task)(nil), "Todo.Task")
	proto.RegisterType((*TaskResponse)(nil), "Todo.TaskResponse")
	proto.RegisterType((*RemoveTaskRequest)(nil), "Todo.RemoveTaskRequest")
	proto.RegisterType((*Product)(nil), "Todo.Product")
	proto.RegisterType((*Index)(nil), "Todo.Index")
	proto.RegisterType((*FindProductRequest)(nil), "Todo.FindProductRequest")
	proto.RegisterType((*FindProductResponse)(nil), "Todo.FindProductResponse")
}

func init() { proto.RegisterFile("todo.proto", fileDescriptor_0e4b95d0c4e09639) }

var fileDescriptor_0e4b95d0c4e09639 = []byte{
	// 322 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x6c, 0x52, 0x4f, 0x4b, 0xfc, 0x30,
	0x14, 0xdc, 0xfe, 0x76, 0xfb, 0xab, 0xfb, 0xaa, 0x07, 0x9f, 0x82, 0xb1, 0x20, 0x94, 0xa0, 0x58,
	0x2f, 0x45, 0xd6, 0xfd, 0x00, 0x1e, 0x54, 0x10, 0x44, 0xa4, 0xec, 0xc1, 0x6b, 0x35, 0x41, 0xca,
	0xd2, 0xa6, 0x6e, 0x52, 0xff, 0x7c, 0x53, 0x3f, 0x8e, 0xf4, 0xa5, 0x71, 0xab, 0xdb, 0x5b, 0x32,
	0xf3, 0xa6, 0x33, 0x9d, 0x17, 0x00, 0xa3, 0x84, 0x4a, 0xeb, 0x95, 0x32, 0x0a, 0x27, 0x0b, 0x25,
	0x14, 0x0f, 0xc0, 0xbf, 0x2e, 0x6b, 0xf3, 0xc9, 0xe7, 0x30, 0x59, 0xe4, 0x7a, 0x89, 0x08, 0x93,
	0xa6, 0x29, 0x04, 0xf3, 0x62, 0x2f, 0x99, 0x66, 0x74, 0x46, 0x06, 0x41, 0x29, 0xb5, 0xce, 0x5f,
	0x24, 0xfb, 0x47, 0xb0, 0xbb, 0xf2, 0x73, 0xd8, 0x6e, 0x55, 0x99, 0xd4, 0xb5, 0xaa, 0xb4, 0xc4,
	0x18, 0x7c, 0x93, 0xeb, 0xa5, 0x66, 0x5e, 0x3c, 0x4e, 0xc2, 0x19, 0xa4, 0xad, 0x49, 0x4a, 0x23,
	0x96, 0xe0, 0xa7, 0xb0, 0x9b, 0xc9, 0x52, 0xbd, 0x49, 0xab, 0x7b, 0x6d, 0xa4, 0x36, 0x43, 0xa6,
	0xfc, 0x11, 0x82, 0x87, 0x95, 0x12, 0xcd, 0xf3, 0x20, 0xdd, 0x62, 0x55, 0x5e, 0xba, 0x40, 0x74,
	0xc6, 0x13, 0x08, 0x8a, 0x4a, 0xc8, 0x0f, 0xa9, 0xd9, 0x98, 0xfc, 0x43, 0xeb, 0x7f, 0xdb, 0x82,
	0x99, 0xe3, 0xf8, 0x11, 0xf8, 0x84, 0xe0, 0x3e, 0xf8, 0x84, 0x75, 0x1f, 0xb6, 0x17, 0x9e, 0x00,
	0xde, 0x14, 0x95, 0xe8, 0xcc, 0x7b, 0x11, 0xc9, 0xcf, 0x5b, 0xfb, 0xf1, 0x4b, 0xd8, 0xfb, 0x35,
	0xd9, 0x95, 0x70, 0x06, 0x5b, 0xb5, 0x85, 0x5c, 0x0f, 0x3b, 0x36, 0x87, 0x1b, 0xfc, 0xa1, 0x67,
	0x5f, 0x1e, 0xd0, 0x1e, 0x30, 0x85, 0xe9, 0x5d, 0xa1, 0x4d, 0x5b, 0x8a, 0xc6, 0x2e, 0x36, 0x2d,
	0x26, 0xc2, 0x5e, 0x87, 0x9d, 0x03, 0x1f, 0xe1, 0x31, 0x04, 0xf7, 0xf2, 0x9d, 0x36, 0xd6, 0x2b,
	0x39, 0xea, 0x2b, 0xf9, 0x08, 0xe7, 0x00, 0xeb, 0xb2, 0xf1, 0xc0, 0x92, 0x1b, 0xf5, 0xff, 0x55,
	0x5d, 0x41, 0xd8, 0xfb, 0x2d, 0x64, 0x96, 0xdd, 0xec, 0x24, 0x3a, 0x1c, 0x60, 0x5c, 0xc2, 0xa7,
	0xff, 0xf4, 0xcc, 0x2e, 0xbe, 0x03, 0x00, 0x00, 0xff, 0xff, 0xe2, 0x65, 0xdb, 0x09, 0x74, 0x02,
	0x00, 0x00,
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConn

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion4

// TodoClient is the client API for Todo service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type TodoClient interface {
	ListTasks(ctx context.Context, in *Empty, opts ...grpc.CallOption) (*TaskResponse, error)
	NewTask(ctx context.Context, in *Task, opts ...grpc.CallOption) (*Empty, error)
	RemoveTask(ctx context.Context, in *RemoveTaskRequest, opts ...grpc.CallOption) (*Empty, error)
	FindProduct(ctx context.Context, in *FindProductRequest, opts ...grpc.CallOption) (*FindProductResponse, error)
}

type todoClient struct {
	cc *grpc.ClientConn
}

func NewTodoClient(cc *grpc.ClientConn) TodoClient {
	return &todoClient{cc}
}

func (c *todoClient) ListTasks(ctx context.Context, in *Empty, opts ...grpc.CallOption) (*TaskResponse, error) {
	out := new(TaskResponse)
	err := c.cc.Invoke(ctx, "/Todo.Todo/ListTasks", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *todoClient) NewTask(ctx context.Context, in *Task, opts ...grpc.CallOption) (*Empty, error) {
	out := new(Empty)
	err := c.cc.Invoke(ctx, "/Todo.Todo/NewTask", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *todoClient) RemoveTask(ctx context.Context, in *RemoveTaskRequest, opts ...grpc.CallOption) (*Empty, error) {
	out := new(Empty)
	err := c.cc.Invoke(ctx, "/Todo.Todo/RemoveTask", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *todoClient) FindProduct(ctx context.Context, in *FindProductRequest, opts ...grpc.CallOption) (*FindProductResponse, error) {
	out := new(FindProductResponse)
	err := c.cc.Invoke(ctx, "/Todo.Todo/FindProduct", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// TodoServer is the server API for Todo service.
type TodoServer interface {
	ListTasks(context.Context, *Empty) (*TaskResponse, error)
	NewTask(context.Context, *Task) (*Empty, error)
	RemoveTask(context.Context, *RemoveTaskRequest) (*Empty, error)
	FindProduct(context.Context, *FindProductRequest) (*FindProductResponse, error)
}

func RegisterTodoServer(s *grpc.Server, srv TodoServer) {
	s.RegisterService(&_Todo_serviceDesc, srv)
}

func _Todo_ListTasks_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(Empty)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(TodoServer).ListTasks(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/Todo.Todo/ListTasks",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(TodoServer).ListTasks(ctx, req.(*Empty))
	}
	return interceptor(ctx, in, info, handler)
}

func _Todo_NewTask_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(Task)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(TodoServer).NewTask(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/Todo.Todo/NewTask",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(TodoServer).NewTask(ctx, req.(*Task))
	}
	return interceptor(ctx, in, info, handler)
}

func _Todo_RemoveTask_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(RemoveTaskRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(TodoServer).RemoveTask(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/Todo.Todo/RemoveTask",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(TodoServer).RemoveTask(ctx, req.(*RemoveTaskRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Todo_FindProduct_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(FindProductRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(TodoServer).FindProduct(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/Todo.Todo/FindProduct",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(TodoServer).FindProduct(ctx, req.(*FindProductRequest))
	}
	return interceptor(ctx, in, info, handler)
}

var _Todo_serviceDesc = grpc.ServiceDesc{
	ServiceName: "Todo.Todo",
	HandlerType: (*TodoServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "ListTasks",
			Handler:    _Todo_ListTasks_Handler,
		},
		{
			MethodName: "NewTask",
			Handler:    _Todo_NewTask_Handler,
		},
		{
			MethodName: "RemoveTask",
			Handler:    _Todo_RemoveTask_Handler,
		},
		{
			MethodName: "FindProduct",
			Handler:    _Todo_FindProduct_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "todo.proto",
}
