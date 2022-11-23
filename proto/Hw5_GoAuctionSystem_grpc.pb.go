// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.2.0
// - protoc             v3.6.1
// source: proto/Hw5_GoAuctionSystem.proto

package GoAuctionSystem

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

// AuctionSystemClient is the client API for AuctionSystem service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type AuctionSystemClient interface {
	Bid(ctx context.Context, in *BidPost, opts ...grpc.CallOption) (*Ack, error)
	Result(ctx context.Context, in *Empty, opts ...grpc.CallOption) (*Outcome, error)
}

type auctionSystemClient struct {
	cc grpc.ClientConnInterface
}

func NewAuctionSystemClient(cc grpc.ClientConnInterface) AuctionSystemClient {
	return &auctionSystemClient{cc}
}

func (c *auctionSystemClient) Bid(ctx context.Context, in *BidPost, opts ...grpc.CallOption) (*Ack, error) {
	out := new(Ack)
	err := c.cc.Invoke(ctx, "/GoAuctionSystem.AuctionSystem/Bid", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *auctionSystemClient) Result(ctx context.Context, in *Empty, opts ...grpc.CallOption) (*Outcome, error) {
	out := new(Outcome)
	err := c.cc.Invoke(ctx, "/GoAuctionSystem.AuctionSystem/Result", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// AuctionSystemServer is the server API for AuctionSystem service.
// All implementations must embed UnimplementedAuctionSystemServer
// for forward compatibility
type AuctionSystemServer interface {
	Bid(context.Context, *BidPost) (*Ack, error)
	Result(context.Context, *Empty) (*Outcome, error)
	mustEmbedUnimplementedAuctionSystemServer()
}

// UnimplementedAuctionSystemServer must be embedded to have forward compatible implementations.
type UnimplementedAuctionSystemServer struct {
}

func (UnimplementedAuctionSystemServer) Bid(context.Context, *BidPost) (*Ack, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Bid not implemented")
}
func (UnimplementedAuctionSystemServer) Result(context.Context, *Empty) (*Outcome, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Result not implemented")
}
func (UnimplementedAuctionSystemServer) mustEmbedUnimplementedAuctionSystemServer() {}

// UnsafeAuctionSystemServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to AuctionSystemServer will
// result in compilation errors.
type UnsafeAuctionSystemServer interface {
	mustEmbedUnimplementedAuctionSystemServer()
}

func RegisterAuctionSystemServer(s grpc.ServiceRegistrar, srv AuctionSystemServer) {
	s.RegisterService(&AuctionSystem_ServiceDesc, srv)
}

func _AuctionSystem_Bid_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(BidPost)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(AuctionSystemServer).Bid(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/GoAuctionSystem.AuctionSystem/Bid",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(AuctionSystemServer).Bid(ctx, req.(*BidPost))
	}
	return interceptor(ctx, in, info, handler)
}

func _AuctionSystem_Result_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(Empty)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(AuctionSystemServer).Result(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/GoAuctionSystem.AuctionSystem/Result",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(AuctionSystemServer).Result(ctx, req.(*Empty))
	}
	return interceptor(ctx, in, info, handler)
}

// AuctionSystem_ServiceDesc is the grpc.ServiceDesc for AuctionSystem service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var AuctionSystem_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "GoAuctionSystem.AuctionSystem",
	HandlerType: (*AuctionSystemServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "Bid",
			Handler:    _AuctionSystem_Bid_Handler,
		},
		{
			MethodName: "Result",
			Handler:    _AuctionSystem_Result_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "proto/Hw5_GoAuctionSystem.proto",
}
