// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.2.0
// - protoc             v3.12.4
// source: quanta.proto

package shared

import (
	context "context"
	empty "github.com/golang/protobuf/ptypes/empty"
	wrappers "github.com/golang/protobuf/ptypes/wrappers"
	grpc "google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
)

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
// Requires gRPC-Go v1.32.0 or later.
const _ = grpc.SupportPackageIsVersion7

// ClusterAdminClient is the client API for ClusterAdmin service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type ClusterAdminClient interface {
	Status(ctx context.Context, in *empty.Empty, opts ...grpc.CallOption) (*StatusMessage, error)
}

type clusterAdminClient struct {
	cc grpc.ClientConnInterface
}

func NewClusterAdminClient(cc grpc.ClientConnInterface) ClusterAdminClient {
	return &clusterAdminClient{cc}
}

func (c *clusterAdminClient) Status(ctx context.Context, in *empty.Empty, opts ...grpc.CallOption) (*StatusMessage, error) {
	out := new(StatusMessage)
	err := c.cc.Invoke(ctx, "/shared.ClusterAdmin/Status", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// ClusterAdminServer is the server API for ClusterAdmin service.
// All implementations should embed UnimplementedClusterAdminServer
// for forward compatibility
type ClusterAdminServer interface {
	Status(context.Context, *empty.Empty) (*StatusMessage, error)
}

// UnimplementedClusterAdminServer should be embedded to have forward compatible implementations.
type UnimplementedClusterAdminServer struct {
}

func (UnimplementedClusterAdminServer) Status(context.Context, *empty.Empty) (*StatusMessage, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Status not implemented")
}

// UnsafeClusterAdminServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to ClusterAdminServer will
// result in compilation errors.
type UnsafeClusterAdminServer interface {
	mustEmbedUnimplementedClusterAdminServer()
}

func RegisterClusterAdminServer(s grpc.ServiceRegistrar, srv ClusterAdminServer) {
	s.RegisterService(&ClusterAdmin_ServiceDesc, srv)
}

func _ClusterAdmin_Status_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(empty.Empty)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ClusterAdminServer).Status(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/shared.ClusterAdmin/Status",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ClusterAdminServer).Status(ctx, req.(*empty.Empty))
	}
	return interceptor(ctx, in, info, handler)
}

// ClusterAdmin_ServiceDesc is the grpc.ServiceDesc for ClusterAdmin service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var ClusterAdmin_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "shared.ClusterAdmin",
	HandlerType: (*ClusterAdminServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "Status",
			Handler:    _ClusterAdmin_Status_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "quanta.proto",
}

// KVStoreClient is the client API for KVStore service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type KVStoreClient interface {
	Put(ctx context.Context, in *IndexKVPair, opts ...grpc.CallOption) (*empty.Empty, error)
	BatchPut(ctx context.Context, opts ...grpc.CallOption) (KVStore_BatchPutClient, error)
	Lookup(ctx context.Context, in *IndexKVPair, opts ...grpc.CallOption) (*IndexKVPair, error)
	BatchLookup(ctx context.Context, opts ...grpc.CallOption) (KVStore_BatchLookupClient, error)
	Items(ctx context.Context, in *wrappers.StringValue, opts ...grpc.CallOption) (KVStore_ItemsClient, error)
	PutStringEnum(ctx context.Context, in *StringEnum, opts ...grpc.CallOption) (*wrappers.UInt64Value, error)
	DeleteIndicesWithPrefix(ctx context.Context, in *DeleteIndicesWithPrefixRequest, opts ...grpc.CallOption) (*empty.Empty, error)
}

type kVStoreClient struct {
	cc grpc.ClientConnInterface
}

func NewKVStoreClient(cc grpc.ClientConnInterface) KVStoreClient {
	return &kVStoreClient{cc}
}

func (c *kVStoreClient) Put(ctx context.Context, in *IndexKVPair, opts ...grpc.CallOption) (*empty.Empty, error) {
	out := new(empty.Empty)
	err := c.cc.Invoke(ctx, "/shared.KVStore/Put", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *kVStoreClient) BatchPut(ctx context.Context, opts ...grpc.CallOption) (KVStore_BatchPutClient, error) {
	stream, err := c.cc.NewStream(ctx, &KVStore_ServiceDesc.Streams[0], "/shared.KVStore/BatchPut", opts...)
	if err != nil {
		return nil, err
	}
	x := &kVStoreBatchPutClient{stream}
	return x, nil
}

type KVStore_BatchPutClient interface {
	Send(*IndexKVPair) error
	CloseAndRecv() (*empty.Empty, error)
	grpc.ClientStream
}

type kVStoreBatchPutClient struct {
	grpc.ClientStream
}

func (x *kVStoreBatchPutClient) Send(m *IndexKVPair) error {
	return x.ClientStream.SendMsg(m)
}

func (x *kVStoreBatchPutClient) CloseAndRecv() (*empty.Empty, error) {
	if err := x.ClientStream.CloseSend(); err != nil {
		return nil, err
	}
	m := new(empty.Empty)
	if err := x.ClientStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

func (c *kVStoreClient) Lookup(ctx context.Context, in *IndexKVPair, opts ...grpc.CallOption) (*IndexKVPair, error) {
	out := new(IndexKVPair)
	err := c.cc.Invoke(ctx, "/shared.KVStore/Lookup", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *kVStoreClient) BatchLookup(ctx context.Context, opts ...grpc.CallOption) (KVStore_BatchLookupClient, error) {
	stream, err := c.cc.NewStream(ctx, &KVStore_ServiceDesc.Streams[1], "/shared.KVStore/BatchLookup", opts...)
	if err != nil {
		return nil, err
	}
	x := &kVStoreBatchLookupClient{stream}
	return x, nil
}

type KVStore_BatchLookupClient interface {
	Send(*IndexKVPair) error
	Recv() (*IndexKVPair, error)
	grpc.ClientStream
}

type kVStoreBatchLookupClient struct {
	grpc.ClientStream
}

func (x *kVStoreBatchLookupClient) Send(m *IndexKVPair) error {
	return x.ClientStream.SendMsg(m)
}

func (x *kVStoreBatchLookupClient) Recv() (*IndexKVPair, error) {
	m := new(IndexKVPair)
	if err := x.ClientStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

func (c *kVStoreClient) Items(ctx context.Context, in *wrappers.StringValue, opts ...grpc.CallOption) (KVStore_ItemsClient, error) {
	stream, err := c.cc.NewStream(ctx, &KVStore_ServiceDesc.Streams[2], "/shared.KVStore/Items", opts...)
	if err != nil {
		return nil, err
	}
	x := &kVStoreItemsClient{stream}
	if err := x.ClientStream.SendMsg(in); err != nil {
		return nil, err
	}
	if err := x.ClientStream.CloseSend(); err != nil {
		return nil, err
	}
	return x, nil
}

type KVStore_ItemsClient interface {
	Recv() (*IndexKVPair, error)
	grpc.ClientStream
}

type kVStoreItemsClient struct {
	grpc.ClientStream
}

func (x *kVStoreItemsClient) Recv() (*IndexKVPair, error) {
	m := new(IndexKVPair)
	if err := x.ClientStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

func (c *kVStoreClient) PutStringEnum(ctx context.Context, in *StringEnum, opts ...grpc.CallOption) (*wrappers.UInt64Value, error) {
	out := new(wrappers.UInt64Value)
	err := c.cc.Invoke(ctx, "/shared.KVStore/PutStringEnum", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *kVStoreClient) DeleteIndicesWithPrefix(ctx context.Context, in *DeleteIndicesWithPrefixRequest, opts ...grpc.CallOption) (*empty.Empty, error) {
	out := new(empty.Empty)
	err := c.cc.Invoke(ctx, "/shared.KVStore/DeleteIndicesWithPrefix", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// KVStoreServer is the server API for KVStore service.
// All implementations should embed UnimplementedKVStoreServer
// for forward compatibility
type KVStoreServer interface {
	Put(context.Context, *IndexKVPair) (*empty.Empty, error)
	BatchPut(KVStore_BatchPutServer) error
	Lookup(context.Context, *IndexKVPair) (*IndexKVPair, error)
	BatchLookup(KVStore_BatchLookupServer) error
	Items(*wrappers.StringValue, KVStore_ItemsServer) error
	PutStringEnum(context.Context, *StringEnum) (*wrappers.UInt64Value, error)
	DeleteIndicesWithPrefix(context.Context, *DeleteIndicesWithPrefixRequest) (*empty.Empty, error)
}

// UnimplementedKVStoreServer should be embedded to have forward compatible implementations.
type UnimplementedKVStoreServer struct {
}

func (UnimplementedKVStoreServer) Put(context.Context, *IndexKVPair) (*empty.Empty, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Put not implemented")
}
func (UnimplementedKVStoreServer) BatchPut(KVStore_BatchPutServer) error {
	return status.Errorf(codes.Unimplemented, "method BatchPut not implemented")
}
func (UnimplementedKVStoreServer) Lookup(context.Context, *IndexKVPair) (*IndexKVPair, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Lookup not implemented")
}
func (UnimplementedKVStoreServer) BatchLookup(KVStore_BatchLookupServer) error {
	return status.Errorf(codes.Unimplemented, "method BatchLookup not implemented")
}
func (UnimplementedKVStoreServer) Items(*wrappers.StringValue, KVStore_ItemsServer) error {
	return status.Errorf(codes.Unimplemented, "method Items not implemented")
}
func (UnimplementedKVStoreServer) PutStringEnum(context.Context, *StringEnum) (*wrappers.UInt64Value, error) {
	return nil, status.Errorf(codes.Unimplemented, "method PutStringEnum not implemented")
}
func (UnimplementedKVStoreServer) DeleteIndicesWithPrefix(context.Context, *DeleteIndicesWithPrefixRequest) (*empty.Empty, error) {
	return nil, status.Errorf(codes.Unimplemented, "method DeleteIndicesWithPrefix not implemented")
}

// UnsafeKVStoreServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to KVStoreServer will
// result in compilation errors.
type UnsafeKVStoreServer interface {
	mustEmbedUnimplementedKVStoreServer()
}

func RegisterKVStoreServer(s grpc.ServiceRegistrar, srv KVStoreServer) {
	s.RegisterService(&KVStore_ServiceDesc, srv)
}

func _KVStore_Put_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(IndexKVPair)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(KVStoreServer).Put(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/shared.KVStore/Put",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(KVStoreServer).Put(ctx, req.(*IndexKVPair))
	}
	return interceptor(ctx, in, info, handler)
}

func _KVStore_BatchPut_Handler(srv interface{}, stream grpc.ServerStream) error {
	return srv.(KVStoreServer).BatchPut(&kVStoreBatchPutServer{stream})
}

type KVStore_BatchPutServer interface {
	SendAndClose(*empty.Empty) error
	Recv() (*IndexKVPair, error)
	grpc.ServerStream
}

type kVStoreBatchPutServer struct {
	grpc.ServerStream
}

func (x *kVStoreBatchPutServer) SendAndClose(m *empty.Empty) error {
	return x.ServerStream.SendMsg(m)
}

func (x *kVStoreBatchPutServer) Recv() (*IndexKVPair, error) {
	m := new(IndexKVPair)
	if err := x.ServerStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

func _KVStore_Lookup_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(IndexKVPair)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(KVStoreServer).Lookup(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/shared.KVStore/Lookup",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(KVStoreServer).Lookup(ctx, req.(*IndexKVPair))
	}
	return interceptor(ctx, in, info, handler)
}

func _KVStore_BatchLookup_Handler(srv interface{}, stream grpc.ServerStream) error {
	return srv.(KVStoreServer).BatchLookup(&kVStoreBatchLookupServer{stream})
}

type KVStore_BatchLookupServer interface {
	Send(*IndexKVPair) error
	Recv() (*IndexKVPair, error)
	grpc.ServerStream
}

type kVStoreBatchLookupServer struct {
	grpc.ServerStream
}

func (x *kVStoreBatchLookupServer) Send(m *IndexKVPair) error {
	return x.ServerStream.SendMsg(m)
}

func (x *kVStoreBatchLookupServer) Recv() (*IndexKVPair, error) {
	m := new(IndexKVPair)
	if err := x.ServerStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

func _KVStore_Items_Handler(srv interface{}, stream grpc.ServerStream) error {
	m := new(wrappers.StringValue)
	if err := stream.RecvMsg(m); err != nil {
		return err
	}
	return srv.(KVStoreServer).Items(m, &kVStoreItemsServer{stream})
}

type KVStore_ItemsServer interface {
	Send(*IndexKVPair) error
	grpc.ServerStream
}

type kVStoreItemsServer struct {
	grpc.ServerStream
}

func (x *kVStoreItemsServer) Send(m *IndexKVPair) error {
	return x.ServerStream.SendMsg(m)
}

func _KVStore_PutStringEnum_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(StringEnum)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(KVStoreServer).PutStringEnum(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/shared.KVStore/PutStringEnum",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(KVStoreServer).PutStringEnum(ctx, req.(*StringEnum))
	}
	return interceptor(ctx, in, info, handler)
}

func _KVStore_DeleteIndicesWithPrefix_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(DeleteIndicesWithPrefixRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(KVStoreServer).DeleteIndicesWithPrefix(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/shared.KVStore/DeleteIndicesWithPrefix",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(KVStoreServer).DeleteIndicesWithPrefix(ctx, req.(*DeleteIndicesWithPrefixRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// KVStore_ServiceDesc is the grpc.ServiceDesc for KVStore service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var KVStore_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "shared.KVStore",
	HandlerType: (*KVStoreServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "Put",
			Handler:    _KVStore_Put_Handler,
		},
		{
			MethodName: "Lookup",
			Handler:    _KVStore_Lookup_Handler,
		},
		{
			MethodName: "PutStringEnum",
			Handler:    _KVStore_PutStringEnum_Handler,
		},
		{
			MethodName: "DeleteIndicesWithPrefix",
			Handler:    _KVStore_DeleteIndicesWithPrefix_Handler,
		},
	},
	Streams: []grpc.StreamDesc{
		{
			StreamName:    "BatchPut",
			Handler:       _KVStore_BatchPut_Handler,
			ClientStreams: true,
		},
		{
			StreamName:    "BatchLookup",
			Handler:       _KVStore_BatchLookup_Handler,
			ServerStreams: true,
			ClientStreams: true,
		},
		{
			StreamName:    "Items",
			Handler:       _KVStore_Items_Handler,
			ServerStreams: true,
		},
	},
	Metadata: "quanta.proto",
}

// StringSearchClient is the client API for StringSearch service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type StringSearchClient interface {
	BatchIndex(ctx context.Context, opts ...grpc.CallOption) (StringSearch_BatchIndexClient, error)
	Search(ctx context.Context, in *wrappers.StringValue, opts ...grpc.CallOption) (StringSearch_SearchClient, error)
}

type stringSearchClient struct {
	cc grpc.ClientConnInterface
}

func NewStringSearchClient(cc grpc.ClientConnInterface) StringSearchClient {
	return &stringSearchClient{cc}
}

func (c *stringSearchClient) BatchIndex(ctx context.Context, opts ...grpc.CallOption) (StringSearch_BatchIndexClient, error) {
	stream, err := c.cc.NewStream(ctx, &StringSearch_ServiceDesc.Streams[0], "/shared.StringSearch/BatchIndex", opts...)
	if err != nil {
		return nil, err
	}
	x := &stringSearchBatchIndexClient{stream}
	return x, nil
}

type StringSearch_BatchIndexClient interface {
	Send(*wrappers.StringValue) error
	CloseAndRecv() (*empty.Empty, error)
	grpc.ClientStream
}

type stringSearchBatchIndexClient struct {
	grpc.ClientStream
}

func (x *stringSearchBatchIndexClient) Send(m *wrappers.StringValue) error {
	return x.ClientStream.SendMsg(m)
}

func (x *stringSearchBatchIndexClient) CloseAndRecv() (*empty.Empty, error) {
	if err := x.ClientStream.CloseSend(); err != nil {
		return nil, err
	}
	m := new(empty.Empty)
	if err := x.ClientStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

func (c *stringSearchClient) Search(ctx context.Context, in *wrappers.StringValue, opts ...grpc.CallOption) (StringSearch_SearchClient, error) {
	stream, err := c.cc.NewStream(ctx, &StringSearch_ServiceDesc.Streams[1], "/shared.StringSearch/Search", opts...)
	if err != nil {
		return nil, err
	}
	x := &stringSearchSearchClient{stream}
	if err := x.ClientStream.SendMsg(in); err != nil {
		return nil, err
	}
	if err := x.ClientStream.CloseSend(); err != nil {
		return nil, err
	}
	return x, nil
}

type StringSearch_SearchClient interface {
	Recv() (*wrappers.UInt64Value, error)
	grpc.ClientStream
}

type stringSearchSearchClient struct {
	grpc.ClientStream
}

func (x *stringSearchSearchClient) Recv() (*wrappers.UInt64Value, error) {
	m := new(wrappers.UInt64Value)
	if err := x.ClientStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

// StringSearchServer is the server API for StringSearch service.
// All implementations should embed UnimplementedStringSearchServer
// for forward compatibility
type StringSearchServer interface {
	BatchIndex(StringSearch_BatchIndexServer) error
	Search(*wrappers.StringValue, StringSearch_SearchServer) error
}

// UnimplementedStringSearchServer should be embedded to have forward compatible implementations.
type UnimplementedStringSearchServer struct {
}

func (UnimplementedStringSearchServer) BatchIndex(StringSearch_BatchIndexServer) error {
	return status.Errorf(codes.Unimplemented, "method BatchIndex not implemented")
}
func (UnimplementedStringSearchServer) Search(*wrappers.StringValue, StringSearch_SearchServer) error {
	return status.Errorf(codes.Unimplemented, "method Search not implemented")
}

// UnsafeStringSearchServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to StringSearchServer will
// result in compilation errors.
type UnsafeStringSearchServer interface {
	mustEmbedUnimplementedStringSearchServer()
}

func RegisterStringSearchServer(s grpc.ServiceRegistrar, srv StringSearchServer) {
	s.RegisterService(&StringSearch_ServiceDesc, srv)
}

func _StringSearch_BatchIndex_Handler(srv interface{}, stream grpc.ServerStream) error {
	return srv.(StringSearchServer).BatchIndex(&stringSearchBatchIndexServer{stream})
}

type StringSearch_BatchIndexServer interface {
	SendAndClose(*empty.Empty) error
	Recv() (*wrappers.StringValue, error)
	grpc.ServerStream
}

type stringSearchBatchIndexServer struct {
	grpc.ServerStream
}

func (x *stringSearchBatchIndexServer) SendAndClose(m *empty.Empty) error {
	return x.ServerStream.SendMsg(m)
}

func (x *stringSearchBatchIndexServer) Recv() (*wrappers.StringValue, error) {
	m := new(wrappers.StringValue)
	if err := x.ServerStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

func _StringSearch_Search_Handler(srv interface{}, stream grpc.ServerStream) error {
	m := new(wrappers.StringValue)
	if err := stream.RecvMsg(m); err != nil {
		return err
	}
	return srv.(StringSearchServer).Search(m, &stringSearchSearchServer{stream})
}

type StringSearch_SearchServer interface {
	Send(*wrappers.UInt64Value) error
	grpc.ServerStream
}

type stringSearchSearchServer struct {
	grpc.ServerStream
}

func (x *stringSearchSearchServer) Send(m *wrappers.UInt64Value) error {
	return x.ServerStream.SendMsg(m)
}

// StringSearch_ServiceDesc is the grpc.ServiceDesc for StringSearch service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var StringSearch_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "shared.StringSearch",
	HandlerType: (*StringSearchServer)(nil),
	Methods:     []grpc.MethodDesc{},
	Streams: []grpc.StreamDesc{
		{
			StreamName:    "BatchIndex",
			Handler:       _StringSearch_BatchIndex_Handler,
			ClientStreams: true,
		},
		{
			StreamName:    "Search",
			Handler:       _StringSearch_Search_Handler,
			ServerStreams: true,
		},
	},
	Metadata: "quanta.proto",
}

// BitmapIndexClient is the client API for BitmapIndex service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type BitmapIndexClient interface {
	Update(ctx context.Context, in *UpdateRequest, opts ...grpc.CallOption) (*empty.Empty, error)
	BatchMutate(ctx context.Context, opts ...grpc.CallOption) (BitmapIndex_BatchMutateClient, error)
	BulkClear(ctx context.Context, in *BulkClearRequest, opts ...grpc.CallOption) (*empty.Empty, error)
	Query(ctx context.Context, in *BitmapQuery, opts ...grpc.CallOption) (*QueryResult, error)
	Join(ctx context.Context, in *JoinRequest, opts ...grpc.CallOption) (*JoinResponse, error)
	Projection(ctx context.Context, in *ProjectionRequest, opts ...grpc.CallOption) (*ProjectionResponse, error)
	CheckoutSequence(ctx context.Context, in *CheckoutSequenceRequest, opts ...grpc.CallOption) (*CheckoutSequenceResponse, error)
	TableOperation(ctx context.Context, in *TableOperationRequest, opts ...grpc.CallOption) (*empty.Empty, error)
}

type bitmapIndexClient struct {
	cc grpc.ClientConnInterface
}

func NewBitmapIndexClient(cc grpc.ClientConnInterface) BitmapIndexClient {
	return &bitmapIndexClient{cc}
}

func (c *bitmapIndexClient) Update(ctx context.Context, in *UpdateRequest, opts ...grpc.CallOption) (*empty.Empty, error) {
	out := new(empty.Empty)
	err := c.cc.Invoke(ctx, "/shared.BitmapIndex/Update", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *bitmapIndexClient) BatchMutate(ctx context.Context, opts ...grpc.CallOption) (BitmapIndex_BatchMutateClient, error) {
	stream, err := c.cc.NewStream(ctx, &BitmapIndex_ServiceDesc.Streams[0], "/shared.BitmapIndex/BatchMutate", opts...)
	if err != nil {
		return nil, err
	}
	x := &bitmapIndexBatchMutateClient{stream}
	return x, nil
}

type BitmapIndex_BatchMutateClient interface {
	Send(*IndexKVPair) error
	CloseAndRecv() (*empty.Empty, error)
	grpc.ClientStream
}

type bitmapIndexBatchMutateClient struct {
	grpc.ClientStream
}

func (x *bitmapIndexBatchMutateClient) Send(m *IndexKVPair) error {
	return x.ClientStream.SendMsg(m)
}

func (x *bitmapIndexBatchMutateClient) CloseAndRecv() (*empty.Empty, error) {
	if err := x.ClientStream.CloseSend(); err != nil {
		return nil, err
	}
	m := new(empty.Empty)
	if err := x.ClientStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

func (c *bitmapIndexClient) BulkClear(ctx context.Context, in *BulkClearRequest, opts ...grpc.CallOption) (*empty.Empty, error) {
	out := new(empty.Empty)
	err := c.cc.Invoke(ctx, "/shared.BitmapIndex/BulkClear", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *bitmapIndexClient) Query(ctx context.Context, in *BitmapQuery, opts ...grpc.CallOption) (*QueryResult, error) {
	out := new(QueryResult)
	err := c.cc.Invoke(ctx, "/shared.BitmapIndex/Query", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *bitmapIndexClient) Join(ctx context.Context, in *JoinRequest, opts ...grpc.CallOption) (*JoinResponse, error) {
	out := new(JoinResponse)
	err := c.cc.Invoke(ctx, "/shared.BitmapIndex/Join", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *bitmapIndexClient) Projection(ctx context.Context, in *ProjectionRequest, opts ...grpc.CallOption) (*ProjectionResponse, error) {
	out := new(ProjectionResponse)
	err := c.cc.Invoke(ctx, "/shared.BitmapIndex/Projection", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *bitmapIndexClient) CheckoutSequence(ctx context.Context, in *CheckoutSequenceRequest, opts ...grpc.CallOption) (*CheckoutSequenceResponse, error) {
	out := new(CheckoutSequenceResponse)
	err := c.cc.Invoke(ctx, "/shared.BitmapIndex/CheckoutSequence", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *bitmapIndexClient) TableOperation(ctx context.Context, in *TableOperationRequest, opts ...grpc.CallOption) (*empty.Empty, error) {
	out := new(empty.Empty)
	err := c.cc.Invoke(ctx, "/shared.BitmapIndex/TableOperation", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// BitmapIndexServer is the server API for BitmapIndex service.
// All implementations should embed UnimplementedBitmapIndexServer
// for forward compatibility
type BitmapIndexServer interface {
	Update(context.Context, *UpdateRequest) (*empty.Empty, error)
	BatchMutate(BitmapIndex_BatchMutateServer) error
	BulkClear(context.Context, *BulkClearRequest) (*empty.Empty, error)
	Query(context.Context, *BitmapQuery) (*QueryResult, error)
	Join(context.Context, *JoinRequest) (*JoinResponse, error)
	Projection(context.Context, *ProjectionRequest) (*ProjectionResponse, error)
	CheckoutSequence(context.Context, *CheckoutSequenceRequest) (*CheckoutSequenceResponse, error)
	TableOperation(context.Context, *TableOperationRequest) (*empty.Empty, error)
}

// UnimplementedBitmapIndexServer should be embedded to have forward compatible implementations.
type UnimplementedBitmapIndexServer struct {
}

func (UnimplementedBitmapIndexServer) Update(context.Context, *UpdateRequest) (*empty.Empty, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Update not implemented")
}
func (UnimplementedBitmapIndexServer) BatchMutate(BitmapIndex_BatchMutateServer) error {
	return status.Errorf(codes.Unimplemented, "method BatchMutate not implemented")
}
func (UnimplementedBitmapIndexServer) BulkClear(context.Context, *BulkClearRequest) (*empty.Empty, error) {
	return nil, status.Errorf(codes.Unimplemented, "method BulkClear not implemented")
}
func (UnimplementedBitmapIndexServer) Query(context.Context, *BitmapQuery) (*QueryResult, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Query not implemented")
}
func (UnimplementedBitmapIndexServer) Join(context.Context, *JoinRequest) (*JoinResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Join not implemented")
}
func (UnimplementedBitmapIndexServer) Projection(context.Context, *ProjectionRequest) (*ProjectionResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Projection not implemented")
}
func (UnimplementedBitmapIndexServer) CheckoutSequence(context.Context, *CheckoutSequenceRequest) (*CheckoutSequenceResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method CheckoutSequence not implemented")
}
func (UnimplementedBitmapIndexServer) TableOperation(context.Context, *TableOperationRequest) (*empty.Empty, error) {
	return nil, status.Errorf(codes.Unimplemented, "method TableOperation not implemented")
}

// UnsafeBitmapIndexServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to BitmapIndexServer will
// result in compilation errors.
type UnsafeBitmapIndexServer interface {
	mustEmbedUnimplementedBitmapIndexServer()
}

func RegisterBitmapIndexServer(s grpc.ServiceRegistrar, srv BitmapIndexServer) {
	s.RegisterService(&BitmapIndex_ServiceDesc, srv)
}

func _BitmapIndex_Update_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(UpdateRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(BitmapIndexServer).Update(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/shared.BitmapIndex/Update",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(BitmapIndexServer).Update(ctx, req.(*UpdateRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _BitmapIndex_BatchMutate_Handler(srv interface{}, stream grpc.ServerStream) error {
	return srv.(BitmapIndexServer).BatchMutate(&bitmapIndexBatchMutateServer{stream})
}

type BitmapIndex_BatchMutateServer interface {
	SendAndClose(*empty.Empty) error
	Recv() (*IndexKVPair, error)
	grpc.ServerStream
}

type bitmapIndexBatchMutateServer struct {
	grpc.ServerStream
}

func (x *bitmapIndexBatchMutateServer) SendAndClose(m *empty.Empty) error {
	return x.ServerStream.SendMsg(m)
}

func (x *bitmapIndexBatchMutateServer) Recv() (*IndexKVPair, error) {
	m := new(IndexKVPair)
	if err := x.ServerStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

func _BitmapIndex_BulkClear_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(BulkClearRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(BitmapIndexServer).BulkClear(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/shared.BitmapIndex/BulkClear",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(BitmapIndexServer).BulkClear(ctx, req.(*BulkClearRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _BitmapIndex_Query_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(BitmapQuery)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(BitmapIndexServer).Query(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/shared.BitmapIndex/Query",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(BitmapIndexServer).Query(ctx, req.(*BitmapQuery))
	}
	return interceptor(ctx, in, info, handler)
}

func _BitmapIndex_Join_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(JoinRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(BitmapIndexServer).Join(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/shared.BitmapIndex/Join",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(BitmapIndexServer).Join(ctx, req.(*JoinRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _BitmapIndex_Projection_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ProjectionRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(BitmapIndexServer).Projection(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/shared.BitmapIndex/Projection",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(BitmapIndexServer).Projection(ctx, req.(*ProjectionRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _BitmapIndex_CheckoutSequence_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(CheckoutSequenceRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(BitmapIndexServer).CheckoutSequence(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/shared.BitmapIndex/CheckoutSequence",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(BitmapIndexServer).CheckoutSequence(ctx, req.(*CheckoutSequenceRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _BitmapIndex_TableOperation_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(TableOperationRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(BitmapIndexServer).TableOperation(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/shared.BitmapIndex/TableOperation",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(BitmapIndexServer).TableOperation(ctx, req.(*TableOperationRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// BitmapIndex_ServiceDesc is the grpc.ServiceDesc for BitmapIndex service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var BitmapIndex_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "shared.BitmapIndex",
	HandlerType: (*BitmapIndexServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "Update",
			Handler:    _BitmapIndex_Update_Handler,
		},
		{
			MethodName: "BulkClear",
			Handler:    _BitmapIndex_BulkClear_Handler,
		},
		{
			MethodName: "Query",
			Handler:    _BitmapIndex_Query_Handler,
		},
		{
			MethodName: "Join",
			Handler:    _BitmapIndex_Join_Handler,
		},
		{
			MethodName: "Projection",
			Handler:    _BitmapIndex_Projection_Handler,
		},
		{
			MethodName: "CheckoutSequence",
			Handler:    _BitmapIndex_CheckoutSequence_Handler,
		},
		{
			MethodName: "TableOperation",
			Handler:    _BitmapIndex_TableOperation_Handler,
		},
	},
	Streams: []grpc.StreamDesc{
		{
			StreamName:    "BatchMutate",
			Handler:       _BitmapIndex_BatchMutate_Handler,
			ClientStreams: true,
		},
	},
	Metadata: "quanta.proto",
}
