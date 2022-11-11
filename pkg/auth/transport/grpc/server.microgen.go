// Code generated by microgen 1.0.5. DO NOT EDIT.

// DO NOT EDIT.
package transportgrpc

import (
	pb "auth/pkg/auth/proto"
	transport "auth/pkg/auth/transport"
	grpc "github.com/go-kit/kit/transport/grpc"
	empty "github.com/golang/protobuf/ptypes/empty"
	context "golang.org/x/net/context"
)

type serviceServer struct {
	pb.UnimplementedServiceServer
	register                 grpc.Handler
	login                    grpc.Handler
	publicKey                grpc.Handler
	getPermissionsForService grpc.Handler
}

func NewGRPCServer(endpoints *transport.EndpointsSet, opts ...grpc.ServerOption) pb.ServiceServer {
	return &serviceServer{
		getPermissionsForService: grpc.NewServer(
			endpoints.GetPermissionsForServiceEndpoint,
			_Decode_GetPermissionsForService_Request,
			_Encode_GetPermissionsForService_Response,
			opts...,
		),
		login: grpc.NewServer(
			endpoints.LoginEndpoint,
			_Decode_Login_Request,
			_Encode_Login_Response,
			opts...,
		),
		publicKey: grpc.NewServer(
			endpoints.PublicKeyEndpoint,
			_Decode_PublicKey_Request,
			_Encode_PublicKey_Response,
			opts...,
		),
		register: grpc.NewServer(
			endpoints.RegisterEndpoint,
			_Decode_Register_Request,
			_Encode_Register_Response,
			opts...,
		),
	}
}

func newOneToManyStreamServer(endpoint transport.OneToManyStreamEndpoint) transport.OneToManyStreamEndpoint {
	return endpoint
}

func newManyToOneStreamServer(endpoint transport.ManyToOneStreamEndpoint) transport.ManyToOneStreamEndpoint {
	return endpoint
}

func newManyToManyStreamServer(endpoint transport.ManyToManyStreamEndpoint) transport.ManyToManyStreamEndpoint {
	return endpoint
}

func (S *serviceServer) Register(ctx context.Context, req *pb.RegisterRequest) (*pb.RegisterResponse, error) {
	_, resp, err := S.register.ServeGRPC(ctx, req)
	if err != nil {
		return nil, err
	}
	return resp.(*pb.RegisterResponse), nil
}

func (S *serviceServer) Login(ctx context.Context, req *pb.LoginRequest) (*pb.LoginResponse, error) {
	_, resp, err := S.login.ServeGRPC(ctx, req)
	if err != nil {
		return nil, err
	}
	return resp.(*pb.LoginResponse), nil
}

func (S *serviceServer) PublicKey(ctx context.Context, req *empty.Empty) (*pb.PublicKeyResponse, error) {
	_, resp, err := S.publicKey.ServeGRPC(ctx, req)
	if err != nil {
		return nil, err
	}
	return resp.(*pb.PublicKeyResponse), nil
}

func (S *serviceServer) GetPermissionsForService(ctx context.Context, req *pb.GetPermissionsForServiceRequest) (*pb.GetPermissionsForServiceResponse, error) {
	_, resp, err := S.getPermissionsForService.ServeGRPC(ctx, req)
	if err != nil {
		return nil, err
	}
	return resp.(*pb.GetPermissionsForServiceResponse), nil
}
