// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: cosmos/authz/v1beta1/query.proto

package types

import (
	context "context"
	fmt "fmt"
	_ "github.com/cosmos/cosmos-sdk/codec/types"
	query "github.com/cosmos/cosmos-sdk/types/query"
	_ "github.com/gogo/protobuf/gogoproto"
	grpc1 "github.com/gogo/protobuf/grpc"
	proto "github.com/gogo/protobuf/proto"
	_ "github.com/regen-network/cosmos-proto"
	_ "google.golang.org/genproto/googleapis/api/annotations"
	grpc "google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
	io "io"
	math "math"
	math_bits "math/bits"
)

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// This is a compile-time assertion to ensure that this generated file
// is compatible with the proto package it is being compiled against.
// A compilation error at this line likely means your copy of the
// proto package needs to be updated.
const _ = proto.GoGoProtoPackageIsVersion3 // please upgrade the proto package

// QueryAuthorizationRequest is the request type for the Query/Authorization RPC method.
type QueryAuthorizationRequest struct {
	Granter string `protobuf:"bytes,1,opt,name=granter,proto3" json:"granter,omitempty"`
	Grantee string `protobuf:"bytes,2,opt,name=grantee,proto3" json:"grantee,omitempty"`
	MsgType string `protobuf:"bytes,3,opt,name=msg_type,json=msgType,proto3" json:"msg_type,omitempty"`
}

func (m *QueryAuthorizationRequest) Reset()         { *m = QueryAuthorizationRequest{} }
func (m *QueryAuthorizationRequest) String() string { return proto.CompactTextString(m) }
func (*QueryAuthorizationRequest) ProtoMessage()    {}
func (*QueryAuthorizationRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_376d714ffdeb1545, []int{0}
}
func (m *QueryAuthorizationRequest) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *QueryAuthorizationRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_QueryAuthorizationRequest.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *QueryAuthorizationRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_QueryAuthorizationRequest.Merge(m, src)
}
func (m *QueryAuthorizationRequest) XXX_Size() int {
	return m.Size()
}
func (m *QueryAuthorizationRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_QueryAuthorizationRequest.DiscardUnknown(m)
}

var xxx_messageInfo_QueryAuthorizationRequest proto.InternalMessageInfo

func (m *QueryAuthorizationRequest) GetGranter() string {
	if m != nil {
		return m.Granter
	}
	return ""
}

func (m *QueryAuthorizationRequest) GetGrantee() string {
	if m != nil {
		return m.Grantee
	}
	return ""
}

func (m *QueryAuthorizationRequest) GetMsgType() string {
	if m != nil {
		return m.MsgType
	}
	return ""
}

// QueryAuthorizationResponse is the response type for the Query/Authorization RPC method.
type QueryAuthorizationResponse struct {
	// authorization is a authorization granted for grantee by granter.
	Authorization *AuthorizationGrant `protobuf:"bytes,1,opt,name=authorization,proto3" json:"authorization,omitempty"`
}

func (m *QueryAuthorizationResponse) Reset()         { *m = QueryAuthorizationResponse{} }
func (m *QueryAuthorizationResponse) String() string { return proto.CompactTextString(m) }
func (*QueryAuthorizationResponse) ProtoMessage()    {}
func (*QueryAuthorizationResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_376d714ffdeb1545, []int{1}
}
func (m *QueryAuthorizationResponse) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *QueryAuthorizationResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_QueryAuthorizationResponse.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *QueryAuthorizationResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_QueryAuthorizationResponse.Merge(m, src)
}
func (m *QueryAuthorizationResponse) XXX_Size() int {
	return m.Size()
}
func (m *QueryAuthorizationResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_QueryAuthorizationResponse.DiscardUnknown(m)
}

var xxx_messageInfo_QueryAuthorizationResponse proto.InternalMessageInfo

func (m *QueryAuthorizationResponse) GetAuthorization() *AuthorizationGrant {
	if m != nil {
		return m.Authorization
	}
	return nil
}

// QueryAuthorizationsRequest is the request type for the Query/Authorizations RPC method.
type QueryAuthorizationsRequest struct {
	Granter string `protobuf:"bytes,1,opt,name=granter,proto3" json:"granter,omitempty"`
	Grantee string `protobuf:"bytes,2,opt,name=grantee,proto3" json:"grantee,omitempty"`
	// pagination defines an pagination for the request.
	Pagination *query.PageRequest `protobuf:"bytes,3,opt,name=pagination,proto3" json:"pagination,omitempty"`
}

func (m *QueryAuthorizationsRequest) Reset()         { *m = QueryAuthorizationsRequest{} }
func (m *QueryAuthorizationsRequest) String() string { return proto.CompactTextString(m) }
func (*QueryAuthorizationsRequest) ProtoMessage()    {}
func (*QueryAuthorizationsRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_376d714ffdeb1545, []int{2}
}
func (m *QueryAuthorizationsRequest) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *QueryAuthorizationsRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_QueryAuthorizationsRequest.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *QueryAuthorizationsRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_QueryAuthorizationsRequest.Merge(m, src)
}
func (m *QueryAuthorizationsRequest) XXX_Size() int {
	return m.Size()
}
func (m *QueryAuthorizationsRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_QueryAuthorizationsRequest.DiscardUnknown(m)
}

var xxx_messageInfo_QueryAuthorizationsRequest proto.InternalMessageInfo

func (m *QueryAuthorizationsRequest) GetGranter() string {
	if m != nil {
		return m.Granter
	}
	return ""
}

func (m *QueryAuthorizationsRequest) GetGrantee() string {
	if m != nil {
		return m.Grantee
	}
	return ""
}

func (m *QueryAuthorizationsRequest) GetPagination() *query.PageRequest {
	if m != nil {
		return m.Pagination
	}
	return nil
}

// QueryAuthorizationsResponse is the response type for the Query/Authorizations RPC method.
type QueryAuthorizationsResponse struct {
	// authorizations is a list of grants granted for grantee by granter.
	Authorizations []*AuthorizationGrant `protobuf:"bytes,1,rep,name=authorizations,proto3" json:"authorizations,omitempty"`
	// pagination defines an pagination for the response.
	Pagination *query.PageResponse `protobuf:"bytes,2,opt,name=pagination,proto3" json:"pagination,omitempty"`
}

func (m *QueryAuthorizationsResponse) Reset()         { *m = QueryAuthorizationsResponse{} }
func (m *QueryAuthorizationsResponse) String() string { return proto.CompactTextString(m) }
func (*QueryAuthorizationsResponse) ProtoMessage()    {}
func (*QueryAuthorizationsResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_376d714ffdeb1545, []int{3}
}
func (m *QueryAuthorizationsResponse) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *QueryAuthorizationsResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_QueryAuthorizationsResponse.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *QueryAuthorizationsResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_QueryAuthorizationsResponse.Merge(m, src)
}
func (m *QueryAuthorizationsResponse) XXX_Size() int {
	return m.Size()
}
func (m *QueryAuthorizationsResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_QueryAuthorizationsResponse.DiscardUnknown(m)
}

var xxx_messageInfo_QueryAuthorizationsResponse proto.InternalMessageInfo

func (m *QueryAuthorizationsResponse) GetAuthorizations() []*AuthorizationGrant {
	if m != nil {
		return m.Authorizations
	}
	return nil
}

func (m *QueryAuthorizationsResponse) GetPagination() *query.PageResponse {
	if m != nil {
		return m.Pagination
	}
	return nil
}

func init() {
	proto.RegisterType((*QueryAuthorizationRequest)(nil), "cosmos.authz.v1beta1.QueryAuthorizationRequest")
	proto.RegisterType((*QueryAuthorizationResponse)(nil), "cosmos.authz.v1beta1.QueryAuthorizationResponse")
	proto.RegisterType((*QueryAuthorizationsRequest)(nil), "cosmos.authz.v1beta1.QueryAuthorizationsRequest")
	proto.RegisterType((*QueryAuthorizationsResponse)(nil), "cosmos.authz.v1beta1.QueryAuthorizationsResponse")
}

func init() { proto.RegisterFile("cosmos/authz/v1beta1/query.proto", fileDescriptor_376d714ffdeb1545) }

var fileDescriptor_376d714ffdeb1545 = []byte{
	// 488 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xa4, 0x94, 0x3d, 0x8f, 0xd3, 0x30,
	0x18, 0xc7, 0xeb, 0x56, 0xbc, 0xf9, 0x74, 0x37, 0x58, 0x37, 0xa4, 0x01, 0x45, 0x55, 0x06, 0x38,
	0x9d, 0x44, 0x4c, 0xcb, 0x27, 0xb8, 0xe3, 0x44, 0x05, 0x03, 0x2a, 0x11, 0x13, 0x4b, 0xe5, 0x14,
	0xe3, 0x06, 0x9a, 0x38, 0x8d, 0x1d, 0x44, 0x8b, 0x58, 0x58, 0x59, 0x90, 0x58, 0xf8, 0x28, 0x2c,
	0x0c, 0x6c, 0x8c, 0x95, 0x58, 0x18, 0x51, 0xcb, 0x07, 0x41, 0xb1, 0x9d, 0x36, 0x41, 0x41, 0xd7,
	0xaa, 0x53, 0x13, 0x3f, 0x2f, 0xff, 0xdf, 0xff, 0xf1, 0xd3, 0xc0, 0xce, 0x88, 0x8b, 0x88, 0x0b,
	0x4c, 0x32, 0x39, 0x9e, 0xe3, 0x37, 0xdd, 0x80, 0x4a, 0xd2, 0xc5, 0xd3, 0x8c, 0xa6, 0x33, 0x2f,
	0x49, 0xb9, 0xe4, 0xe8, 0x58, 0x67, 0x78, 0x2a, 0xc3, 0x33, 0x19, 0xf6, 0x31, 0xe3, 0x8c, 0xab,
	0x04, 0x9c, 0x3f, 0xe9, 0x5c, 0xbb, 0xcd, 0x38, 0x67, 0x13, 0x8a, 0xd5, 0x5b, 0x90, 0xbd, 0xc4,
	0x24, 0x36, 0x6d, 0xec, 0x5b, 0x26, 0x44, 0x92, 0x10, 0x93, 0x38, 0xe6, 0x92, 0xc8, 0x90, 0xc7,
	0xa2, 0x28, 0xd4, 0x22, 0x43, 0xdd, 0xd1, 0x28, 0xea, 0xd0, 0xa9, 0x21, 0x0c, 0x88, 0xa0, 0x1a,
	0x6c, 0x8d, 0x99, 0x10, 0x16, 0xc6, 0xaa, 0x8f, 0xc9, 0xad, 0x77, 0xa3, 0xc9, 0x55, 0x86, 0xfb,
	0x0a, 0xb6, 0x9f, 0xe6, 0x3d, 0xce, 0x32, 0x39, 0xe6, 0x69, 0x38, 0x57, 0xd5, 0x3e, 0x9d, 0x66,
	0x54, 0x48, 0x64, 0xc1, 0x6b, 0x2c, 0x25, 0xb1, 0xa4, 0xa9, 0x05, 0x3a, 0xe0, 0xe4, 0x86, 0x5f,
	0xbc, 0x6e, 0x22, 0xd4, 0x6a, 0x96, 0x23, 0x14, 0xb5, 0xe1, 0xf5, 0x48, 0xb0, 0xa1, 0x9c, 0x25,
	0xd4, 0x6a, 0xe9, 0x50, 0x24, 0xd8, 0xb3, 0x59, 0x42, 0xdd, 0x09, 0xb4, 0xeb, 0xb4, 0x44, 0xc2,
	0x63, 0x41, 0xd1, 0x13, 0x78, 0x48, 0xca, 0x01, 0x25, 0x79, 0xd0, 0x3b, 0xf1, 0xea, 0xe6, 0xed,
	0x55, 0x7a, 0xf4, 0x73, 0x6d, 0xbf, 0x5a, 0xee, 0x7e, 0x01, 0x75, 0x72, 0x62, 0x1f, 0x6f, 0x0f,
	0x21, 0xdc, 0x8c, 0x58, 0xb9, 0x3b, 0xe8, 0xdd, 0x2e, 0xf8, 0xf2, 0xfb, 0xf0, 0xf4, 0xa2, 0x14,
	0x90, 0x03, 0xc2, 0xa8, 0xd1, 0xf3, 0x4b, 0x95, 0xee, 0x57, 0x00, 0x6f, 0xd6, 0xa2, 0x99, 0x51,
	0x0c, 0xe0, 0x51, 0xc5, 0x8b, 0xb0, 0x40, 0xa7, 0xb5, 0xd3, 0x2c, 0xfe, 0xa9, 0x47, 0xfd, 0x0a,
	0x79, 0x53, 0x91, 0xdf, 0xb9, 0x94, 0x5c, 0xe3, 0x94, 0xd1, 0x7b, 0x1f, 0x5b, 0xf0, 0x8a, 0x42,
	0x47, 0xdf, 0x00, 0x3c, 0xac, 0x28, 0x23, 0x5c, 0x8f, 0xf7, 0xdf, 0xfd, 0xb2, 0xef, 0x6d, 0x5f,
	0xa0, 0x51, 0xdc, 0x47, 0x1f, 0x7e, 0xfe, 0xf9, 0xdc, 0x7c, 0x80, 0xce, 0x70, 0xed, 0x66, 0x9b,
	0x2b, 0x14, 0xf8, 0x9d, 0x79, 0x7a, 0x6f, 0x8e, 0xe8, 0xfa, 0x88, 0x9a, 0x23, 0xf4, 0x1d, 0xc0,
	0xa3, 0xea, 0xfc, 0xd1, 0xd6, 0x3c, 0xc5, 0x16, 0xd9, 0xdd, 0x1d, 0x2a, 0x8c, 0x85, 0xc7, 0xca,
	0xc2, 0x05, 0x3a, 0xdf, 0xdb, 0x82, 0x38, 0xbf, 0xf8, 0xb1, 0x74, 0xc0, 0x62, 0xe9, 0x80, 0xdf,
	0x4b, 0x07, 0x7c, 0x5a, 0x39, 0x8d, 0xc5, 0xca, 0x69, 0xfc, 0x5a, 0x39, 0x8d, 0xe7, 0xa7, 0x2c,
	0x94, 0xe3, 0x2c, 0xf0, 0x46, 0x3c, 0x2a, 0x74, 0xf4, 0xcf, 0x5d, 0xf1, 0xe2, 0x35, 0x7e, 0x6b,
	0x44, 0xf3, 0x7f, 0xa9, 0x08, 0xae, 0xaa, 0x4f, 0xc1, 0xfd, 0xbf, 0x01, 0x00, 0x00, 0xff, 0xff,
	0x2a, 0x99, 0xdb, 0x07, 0xfc, 0x04, 0x00, 0x00,
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConn

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion4

// QueryClient is the client API for Query service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type QueryClient interface {
	// Returns any `Authorization` (or `nil`), with the expiration time, granted to the grantee by the granter for the
	// provided msg type.
	Authorization(ctx context.Context, in *QueryAuthorizationRequest, opts ...grpc.CallOption) (*QueryAuthorizationResponse, error)
	// Returns list of `Authorization`, granted to the grantee by the granter.
	Authorizations(ctx context.Context, in *QueryAuthorizationsRequest, opts ...grpc.CallOption) (*QueryAuthorizationsResponse, error)
}

type queryClient struct {
	cc grpc1.ClientConn
}

func NewQueryClient(cc grpc1.ClientConn) QueryClient {
	return &queryClient{cc}
}

func (c *queryClient) Authorization(ctx context.Context, in *QueryAuthorizationRequest, opts ...grpc.CallOption) (*QueryAuthorizationResponse, error) {
	out := new(QueryAuthorizationResponse)
	err := c.cc.Invoke(ctx, "/cosmos.authz.v1beta1.Query/Authorization", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *queryClient) Authorizations(ctx context.Context, in *QueryAuthorizationsRequest, opts ...grpc.CallOption) (*QueryAuthorizationsResponse, error) {
	out := new(QueryAuthorizationsResponse)
	err := c.cc.Invoke(ctx, "/cosmos.authz.v1beta1.Query/Authorizations", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// QueryServer is the server API for Query service.
type QueryServer interface {
	// Returns any `Authorization` (or `nil`), with the expiration time, granted to the grantee by the granter for the
	// provided msg type.
	Authorization(context.Context, *QueryAuthorizationRequest) (*QueryAuthorizationResponse, error)
	// Returns list of `Authorization`, granted to the grantee by the granter.
	Authorizations(context.Context, *QueryAuthorizationsRequest) (*QueryAuthorizationsResponse, error)
}

// UnimplementedQueryServer can be embedded to have forward compatible implementations.
type UnimplementedQueryServer struct {
}

func (*UnimplementedQueryServer) Authorization(ctx context.Context, req *QueryAuthorizationRequest) (*QueryAuthorizationResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Authorization not implemented")
}
func (*UnimplementedQueryServer) Authorizations(ctx context.Context, req *QueryAuthorizationsRequest) (*QueryAuthorizationsResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Authorizations not implemented")
}

func RegisterQueryServer(s grpc1.Server, srv QueryServer) {
	s.RegisterService(&_Query_serviceDesc, srv)
}

func _Query_Authorization_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(QueryAuthorizationRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(QueryServer).Authorization(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/cosmos.authz.v1beta1.Query/Authorization",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(QueryServer).Authorization(ctx, req.(*QueryAuthorizationRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Query_Authorizations_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(QueryAuthorizationsRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(QueryServer).Authorizations(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/cosmos.authz.v1beta1.Query/Authorizations",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(QueryServer).Authorizations(ctx, req.(*QueryAuthorizationsRequest))
	}
	return interceptor(ctx, in, info, handler)
}

var _Query_serviceDesc = grpc.ServiceDesc{
	ServiceName: "cosmos.authz.v1beta1.Query",
	HandlerType: (*QueryServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "Authorization",
			Handler:    _Query_Authorization_Handler,
		},
		{
			MethodName: "Authorizations",
			Handler:    _Query_Authorizations_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "cosmos/authz/v1beta1/query.proto",
}

func (m *QueryAuthorizationRequest) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *QueryAuthorizationRequest) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *QueryAuthorizationRequest) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if len(m.MsgType) > 0 {
		i -= len(m.MsgType)
		copy(dAtA[i:], m.MsgType)
		i = encodeVarintQuery(dAtA, i, uint64(len(m.MsgType)))
		i--
		dAtA[i] = 0x1a
	}
	if len(m.Grantee) > 0 {
		i -= len(m.Grantee)
		copy(dAtA[i:], m.Grantee)
		i = encodeVarintQuery(dAtA, i, uint64(len(m.Grantee)))
		i--
		dAtA[i] = 0x12
	}
	if len(m.Granter) > 0 {
		i -= len(m.Granter)
		copy(dAtA[i:], m.Granter)
		i = encodeVarintQuery(dAtA, i, uint64(len(m.Granter)))
		i--
		dAtA[i] = 0xa
	}
	return len(dAtA) - i, nil
}

func (m *QueryAuthorizationResponse) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *QueryAuthorizationResponse) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *QueryAuthorizationResponse) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if m.Authorization != nil {
		{
			size, err := m.Authorization.MarshalToSizedBuffer(dAtA[:i])
			if err != nil {
				return 0, err
			}
			i -= size
			i = encodeVarintQuery(dAtA, i, uint64(size))
		}
		i--
		dAtA[i] = 0xa
	}
	return len(dAtA) - i, nil
}

func (m *QueryAuthorizationsRequest) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *QueryAuthorizationsRequest) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *QueryAuthorizationsRequest) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if m.Pagination != nil {
		{
			size, err := m.Pagination.MarshalToSizedBuffer(dAtA[:i])
			if err != nil {
				return 0, err
			}
			i -= size
			i = encodeVarintQuery(dAtA, i, uint64(size))
		}
		i--
		dAtA[i] = 0x1a
	}
	if len(m.Grantee) > 0 {
		i -= len(m.Grantee)
		copy(dAtA[i:], m.Grantee)
		i = encodeVarintQuery(dAtA, i, uint64(len(m.Grantee)))
		i--
		dAtA[i] = 0x12
	}
	if len(m.Granter) > 0 {
		i -= len(m.Granter)
		copy(dAtA[i:], m.Granter)
		i = encodeVarintQuery(dAtA, i, uint64(len(m.Granter)))
		i--
		dAtA[i] = 0xa
	}
	return len(dAtA) - i, nil
}

func (m *QueryAuthorizationsResponse) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *QueryAuthorizationsResponse) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *QueryAuthorizationsResponse) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if m.Pagination != nil {
		{
			size, err := m.Pagination.MarshalToSizedBuffer(dAtA[:i])
			if err != nil {
				return 0, err
			}
			i -= size
			i = encodeVarintQuery(dAtA, i, uint64(size))
		}
		i--
		dAtA[i] = 0x12
	}
	if len(m.Authorizations) > 0 {
		for iNdEx := len(m.Authorizations) - 1; iNdEx >= 0; iNdEx-- {
			{
				size, err := m.Authorizations[iNdEx].MarshalToSizedBuffer(dAtA[:i])
				if err != nil {
					return 0, err
				}
				i -= size
				i = encodeVarintQuery(dAtA, i, uint64(size))
			}
			i--
			dAtA[i] = 0xa
		}
	}
	return len(dAtA) - i, nil
}

func encodeVarintQuery(dAtA []byte, offset int, v uint64) int {
	offset -= sovQuery(v)
	base := offset
	for v >= 1<<7 {
		dAtA[offset] = uint8(v&0x7f | 0x80)
		v >>= 7
		offset++
	}
	dAtA[offset] = uint8(v)
	return base
}
func (m *QueryAuthorizationRequest) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	l = len(m.Granter)
	if l > 0 {
		n += 1 + l + sovQuery(uint64(l))
	}
	l = len(m.Grantee)
	if l > 0 {
		n += 1 + l + sovQuery(uint64(l))
	}
	l = len(m.MsgType)
	if l > 0 {
		n += 1 + l + sovQuery(uint64(l))
	}
	return n
}

func (m *QueryAuthorizationResponse) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	if m.Authorization != nil {
		l = m.Authorization.Size()
		n += 1 + l + sovQuery(uint64(l))
	}
	return n
}

func (m *QueryAuthorizationsRequest) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	l = len(m.Granter)
	if l > 0 {
		n += 1 + l + sovQuery(uint64(l))
	}
	l = len(m.Grantee)
	if l > 0 {
		n += 1 + l + sovQuery(uint64(l))
	}
	if m.Pagination != nil {
		l = m.Pagination.Size()
		n += 1 + l + sovQuery(uint64(l))
	}
	return n
}

func (m *QueryAuthorizationsResponse) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	if len(m.Authorizations) > 0 {
		for _, e := range m.Authorizations {
			l = e.Size()
			n += 1 + l + sovQuery(uint64(l))
		}
	}
	if m.Pagination != nil {
		l = m.Pagination.Size()
		n += 1 + l + sovQuery(uint64(l))
	}
	return n
}

func sovQuery(x uint64) (n int) {
	return (math_bits.Len64(x|1) + 6) / 7
}
func sozQuery(x uint64) (n int) {
	return sovQuery(uint64((x << 1) ^ uint64((int64(x) >> 63))))
}
func (m *QueryAuthorizationRequest) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowQuery
			}
			if iNdEx >= l {
				return io.ErrUnexpectedEOF
			}
			b := dAtA[iNdEx]
			iNdEx++
			wire |= uint64(b&0x7F) << shift
			if b < 0x80 {
				break
			}
		}
		fieldNum := int32(wire >> 3)
		wireType := int(wire & 0x7)
		if wireType == 4 {
			return fmt.Errorf("proto: QueryAuthorizationRequest: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: QueryAuthorizationRequest: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Granter", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowQuery
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				stringLen |= uint64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			intStringLen := int(stringLen)
			if intStringLen < 0 {
				return ErrInvalidLengthQuery
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthQuery
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Granter = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 2:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Grantee", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowQuery
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				stringLen |= uint64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			intStringLen := int(stringLen)
			if intStringLen < 0 {
				return ErrInvalidLengthQuery
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthQuery
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Grantee = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 3:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field MsgType", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowQuery
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				stringLen |= uint64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			intStringLen := int(stringLen)
			if intStringLen < 0 {
				return ErrInvalidLengthQuery
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthQuery
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.MsgType = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipQuery(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if skippy < 0 {
				return ErrInvalidLengthQuery
			}
			if (iNdEx + skippy) < 0 {
				return ErrInvalidLengthQuery
			}
			if (iNdEx + skippy) > l {
				return io.ErrUnexpectedEOF
			}
			iNdEx += skippy
		}
	}

	if iNdEx > l {
		return io.ErrUnexpectedEOF
	}
	return nil
}
func (m *QueryAuthorizationResponse) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowQuery
			}
			if iNdEx >= l {
				return io.ErrUnexpectedEOF
			}
			b := dAtA[iNdEx]
			iNdEx++
			wire |= uint64(b&0x7F) << shift
			if b < 0x80 {
				break
			}
		}
		fieldNum := int32(wire >> 3)
		wireType := int(wire & 0x7)
		if wireType == 4 {
			return fmt.Errorf("proto: QueryAuthorizationResponse: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: QueryAuthorizationResponse: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Authorization", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowQuery
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				msglen |= int(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			if msglen < 0 {
				return ErrInvalidLengthQuery
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthQuery
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if m.Authorization == nil {
				m.Authorization = &AuthorizationGrant{}
			}
			if err := m.Authorization.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipQuery(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if skippy < 0 {
				return ErrInvalidLengthQuery
			}
			if (iNdEx + skippy) < 0 {
				return ErrInvalidLengthQuery
			}
			if (iNdEx + skippy) > l {
				return io.ErrUnexpectedEOF
			}
			iNdEx += skippy
		}
	}

	if iNdEx > l {
		return io.ErrUnexpectedEOF
	}
	return nil
}
func (m *QueryAuthorizationsRequest) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowQuery
			}
			if iNdEx >= l {
				return io.ErrUnexpectedEOF
			}
			b := dAtA[iNdEx]
			iNdEx++
			wire |= uint64(b&0x7F) << shift
			if b < 0x80 {
				break
			}
		}
		fieldNum := int32(wire >> 3)
		wireType := int(wire & 0x7)
		if wireType == 4 {
			return fmt.Errorf("proto: QueryAuthorizationsRequest: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: QueryAuthorizationsRequest: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Granter", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowQuery
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				stringLen |= uint64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			intStringLen := int(stringLen)
			if intStringLen < 0 {
				return ErrInvalidLengthQuery
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthQuery
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Granter = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 2:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Grantee", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowQuery
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				stringLen |= uint64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			intStringLen := int(stringLen)
			if intStringLen < 0 {
				return ErrInvalidLengthQuery
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthQuery
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Grantee = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 3:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Pagination", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowQuery
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				msglen |= int(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			if msglen < 0 {
				return ErrInvalidLengthQuery
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthQuery
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if m.Pagination == nil {
				m.Pagination = &query.PageRequest{}
			}
			if err := m.Pagination.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipQuery(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if skippy < 0 {
				return ErrInvalidLengthQuery
			}
			if (iNdEx + skippy) < 0 {
				return ErrInvalidLengthQuery
			}
			if (iNdEx + skippy) > l {
				return io.ErrUnexpectedEOF
			}
			iNdEx += skippy
		}
	}

	if iNdEx > l {
		return io.ErrUnexpectedEOF
	}
	return nil
}
func (m *QueryAuthorizationsResponse) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowQuery
			}
			if iNdEx >= l {
				return io.ErrUnexpectedEOF
			}
			b := dAtA[iNdEx]
			iNdEx++
			wire |= uint64(b&0x7F) << shift
			if b < 0x80 {
				break
			}
		}
		fieldNum := int32(wire >> 3)
		wireType := int(wire & 0x7)
		if wireType == 4 {
			return fmt.Errorf("proto: QueryAuthorizationsResponse: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: QueryAuthorizationsResponse: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Authorizations", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowQuery
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				msglen |= int(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			if msglen < 0 {
				return ErrInvalidLengthQuery
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthQuery
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Authorizations = append(m.Authorizations, &AuthorizationGrant{})
			if err := m.Authorizations[len(m.Authorizations)-1].Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 2:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Pagination", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowQuery
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				msglen |= int(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			if msglen < 0 {
				return ErrInvalidLengthQuery
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthQuery
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if m.Pagination == nil {
				m.Pagination = &query.PageResponse{}
			}
			if err := m.Pagination.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipQuery(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if skippy < 0 {
				return ErrInvalidLengthQuery
			}
			if (iNdEx + skippy) < 0 {
				return ErrInvalidLengthQuery
			}
			if (iNdEx + skippy) > l {
				return io.ErrUnexpectedEOF
			}
			iNdEx += skippy
		}
	}

	if iNdEx > l {
		return io.ErrUnexpectedEOF
	}
	return nil
}
func skipQuery(dAtA []byte) (n int, err error) {
	l := len(dAtA)
	iNdEx := 0
	depth := 0
	for iNdEx < l {
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return 0, ErrIntOverflowQuery
			}
			if iNdEx >= l {
				return 0, io.ErrUnexpectedEOF
			}
			b := dAtA[iNdEx]
			iNdEx++
			wire |= (uint64(b) & 0x7F) << shift
			if b < 0x80 {
				break
			}
		}
		wireType := int(wire & 0x7)
		switch wireType {
		case 0:
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return 0, ErrIntOverflowQuery
				}
				if iNdEx >= l {
					return 0, io.ErrUnexpectedEOF
				}
				iNdEx++
				if dAtA[iNdEx-1] < 0x80 {
					break
				}
			}
		case 1:
			iNdEx += 8
		case 2:
			var length int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return 0, ErrIntOverflowQuery
				}
				if iNdEx >= l {
					return 0, io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				length |= (int(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			if length < 0 {
				return 0, ErrInvalidLengthQuery
			}
			iNdEx += length
		case 3:
			depth++
		case 4:
			if depth == 0 {
				return 0, ErrUnexpectedEndOfGroupQuery
			}
			depth--
		case 5:
			iNdEx += 4
		default:
			return 0, fmt.Errorf("proto: illegal wireType %d", wireType)
		}
		if iNdEx < 0 {
			return 0, ErrInvalidLengthQuery
		}
		if depth == 0 {
			return iNdEx, nil
		}
	}
	return 0, io.ErrUnexpectedEOF
}

var (
	ErrInvalidLengthQuery        = fmt.Errorf("proto: negative length found during unmarshaling")
	ErrIntOverflowQuery          = fmt.Errorf("proto: integer overflow")
	ErrUnexpectedEndOfGroupQuery = fmt.Errorf("proto: unexpected end of group")
)