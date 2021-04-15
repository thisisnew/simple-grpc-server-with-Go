// Code generated by protoc-gen-go-grpc. DO NOT EDIT.

package vehicle

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

// VehicleClient is the client API for Vehicle service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type VehicleClient interface {
	GetVehicle(ctx context.Context, in *GetVehicleRequest, opts ...grpc.CallOption) (*GetVehicleResponse, error)
	ListVehicles(ctx context.Context, in *ListVehiclesRequest, opts ...grpc.CallOption) (*ListVehiclesResponse, error)
}

type vehicleClient struct {
	cc grpc.ClientConnInterface
}

func NewVehicleClient(cc grpc.ClientConnInterface) VehicleClient {
	return &vehicleClient{cc}
}

func (c *vehicleClient) GetVehicle(ctx context.Context, in *GetVehicleRequest, opts ...grpc.CallOption) (*GetVehicleResponse, error) {
	out := new(GetVehicleResponse)
	err := c.cc.Invoke(ctx, "/protos.vehicle.Vehicle/GetVehicle", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *vehicleClient) ListVehicles(ctx context.Context, in *ListVehiclesRequest, opts ...grpc.CallOption) (*ListVehiclesResponse, error) {
	out := new(ListVehiclesResponse)
	err := c.cc.Invoke(ctx, "/protos.vehicle.Vehicle/ListVehicles", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// VehicleServer is the server API for Vehicle service.
// All implementations must embed UnimplementedVehicleServer
// for forward compatibility
type VehicleServer interface {
	GetVehicle(context.Context, *GetVehicleRequest) (*GetVehicleResponse, error)
	ListVehicles(context.Context, *ListVehiclesRequest) (*ListVehiclesResponse, error)
	mustEmbedUnimplementedVehicleServer()
}

// UnimplementedVehicleServer must be embedded to have forward compatible implementations.
type UnimplementedVehicleServer struct {
}

func (UnimplementedVehicleServer) GetVehicle(context.Context, *GetVehicleRequest) (*GetVehicleResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetVehicle not implemented")
}
func (UnimplementedVehicleServer) ListVehicles(context.Context, *ListVehiclesRequest) (*ListVehiclesResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method ListVehicles not implemented")
}
func (UnimplementedVehicleServer) mustEmbedUnimplementedVehicleServer() {}

// UnsafeVehicleServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to VehicleServer will
// result in compilation errors.
type UnsafeVehicleServer interface {
	mustEmbedUnimplementedVehicleServer()
}

func RegisterVehicleServer(s grpc.ServiceRegistrar, srv VehicleServer) {
	s.RegisterService(&Vehicle_ServiceDesc, srv)
}

func _Vehicle_GetVehicle_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetVehicleRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(VehicleServer).GetVehicle(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/protos.vehicle.Vehicle/GetVehicle",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(VehicleServer).GetVehicle(ctx, req.(*GetVehicleRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Vehicle_ListVehicles_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ListVehiclesRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(VehicleServer).ListVehicles(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/protos.vehicle.Vehicle/ListVehicles",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(VehicleServer).ListVehicles(ctx, req.(*ListVehiclesRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// Vehicle_ServiceDesc is the grpc.ServiceDesc for Vehicle service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var Vehicle_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "protos.vehicle.Vehicle",
	HandlerType: (*VehicleServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "GetVehicle",
			Handler:    _Vehicle_GetVehicle_Handler,
		},
		{
			MethodName: "ListVehicles",
			Handler:    _Vehicle_ListVehicles_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "protos/vehicle.proto",
}
