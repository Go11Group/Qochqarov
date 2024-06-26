// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.2.0
// - protoc             v3.12.4
// source: library/library.proto

package library

import (
	context "context"
	grpc "google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
)

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
// Requires gRPC-Go v1.32.0 or later.
const _ = grpc.SupportPackageIsVersion7

// LibraryServerClient is the client API for LibraryServer service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type LibraryServerClient interface {
	AddBook(ctx context.Context, in *AddBookRequest, opts ...grpc.CallOption) (*AddBookRespons, error)
	SearchBook(ctx context.Context, in *SearchBookRequest, opts ...grpc.CallOption) (*SearchBookRespons, error)
	BorrowBook(ctx context.Context, in *BorrowBookRequest, opts ...grpc.CallOption) (*BorrowBookRespons, error)
}

type libraryServerClient struct {
	cc grpc.ClientConnInterface
}

func NewLibraryServerClient(cc grpc.ClientConnInterface) LibraryServerClient {
	return &libraryServerClient{cc}
}

func (c *libraryServerClient) AddBook(ctx context.Context, in *AddBookRequest, opts ...grpc.CallOption) (*AddBookRespons, error) {
	out := new(AddBookRespons)
	err := c.cc.Invoke(ctx, "/library.LibraryServer/AddBook", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *libraryServerClient) SearchBook(ctx context.Context, in *SearchBookRequest, opts ...grpc.CallOption) (*SearchBookRespons, error) {
	out := new(SearchBookRespons)
	err := c.cc.Invoke(ctx, "/library.LibraryServer/SearchBook", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *libraryServerClient) BorrowBook(ctx context.Context, in *BorrowBookRequest, opts ...grpc.CallOption) (*BorrowBookRespons, error) {
	out := new(BorrowBookRespons)
	err := c.cc.Invoke(ctx, "/library.LibraryServer/BorrowBook", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// LibraryServerServer is the server API for LibraryServer service.
// All implementations must embed UnimplementedLibraryServerServer
// for forward compatibility
type LibraryServerServer interface {
	AddBook(context.Context, *AddBookRequest) (*AddBookRespons, error)
	SearchBook(context.Context, *SearchBookRequest) (*SearchBookRespons, error)
	BorrowBook(context.Context, *BorrowBookRequest) (*BorrowBookRespons, error)
	mustEmbedUnimplementedLibraryServerServer()
}

// UnimplementedLibraryServerServer must be embedded to have forward compatible implementations.
type UnimplementedLibraryServerServer struct {
}

func (UnimplementedLibraryServerServer) AddBook(context.Context, *AddBookRequest) (*AddBookRespons, error) {
	return nil, status.Errorf(codes.Unimplemented, "method AddBook not implemented")
}
func (UnimplementedLibraryServerServer) SearchBook(context.Context, *SearchBookRequest) (*SearchBookRespons, error) {
	return nil, status.Errorf(codes.Unimplemented, "method SearchBook not implemented")
}
func (UnimplementedLibraryServerServer) BorrowBook(context.Context, *BorrowBookRequest) (*BorrowBookRespons, error) {
	return nil, status.Errorf(codes.Unimplemented, "method BorrowBook not implemented")
}
func (UnimplementedLibraryServerServer) mustEmbedUnimplementedLibraryServerServer() {}

// UnsafeLibraryServerServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to LibraryServerServer will
// result in compilation errors.
type UnsafeLibraryServerServer interface {
	mustEmbedUnimplementedLibraryServerServer()
}

func RegisterLibraryServerServer(s grpc.ServiceRegistrar, srv LibraryServerServer) {
	s.RegisterService(&LibraryServer_ServiceDesc, srv)
}

func _LibraryServer_AddBook_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(AddBookRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(LibraryServerServer).AddBook(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/library.LibraryServer/AddBook",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(LibraryServerServer).AddBook(ctx, req.(*AddBookRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _LibraryServer_SearchBook_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(SearchBookRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(LibraryServerServer).SearchBook(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/library.LibraryServer/SearchBook",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(LibraryServerServer).SearchBook(ctx, req.(*SearchBookRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _LibraryServer_BorrowBook_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(BorrowBookRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(LibraryServerServer).BorrowBook(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/library.LibraryServer/BorrowBook",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(LibraryServerServer).BorrowBook(ctx, req.(*BorrowBookRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// LibraryServer_ServiceDesc is the grpc.ServiceDesc for LibraryServer service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var LibraryServer_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "library.LibraryServer",
	HandlerType: (*LibraryServerServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "AddBook",
			Handler:    _LibraryServer_AddBook_Handler,
		},
		{
			MethodName: "SearchBook",
			Handler:    _LibraryServer_SearchBook_Handler,
		},
		{
			MethodName: "BorrowBook",
			Handler:    _LibraryServer_BorrowBook_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "library/library.proto",
}
