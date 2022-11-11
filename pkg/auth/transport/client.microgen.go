// Code generated by microgen 1.0.5. DO NOT EDIT.

package transport

import (
	types "auth/pkg/types"
	"context"
	"errors"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
)

func (set EndpointsSet) Register(arg0 context.Context, arg1 string, arg2 string, arg3 string, arg4 uint32) (res0 bool, res1 error) {
	request := RegisterRequest{
		AccountId: arg4,
		Login:     arg1,
		Password:  arg2,
		Service:   arg3,
	}
	response, res1 := set.RegisterEndpoint(arg0, &request)
	if res1 != nil {
		if e, ok := status.FromError(res1); ok || e.Code() == codes.Internal || e.Code() == codes.Unknown {
			res1 = errors.New(e.Message())
		}
		return
	}
	return response.(*RegisterResponse).Ok, res1
}

func (set EndpointsSet) Login(arg0 context.Context, arg1 string, arg2 string, arg3 string) (res0 *types.AccessToken, res1 error) {
	request := LoginRequest{
		Login:    arg1,
		Password: arg2,
		Service:  arg3,
	}
	response, res1 := set.LoginEndpoint(arg0, &request)
	if res1 != nil {
		if e, ok := status.FromError(res1); ok || e.Code() == codes.Internal || e.Code() == codes.Unknown {
			res1 = errors.New(e.Message())
		}
		return
	}
	return response.(*LoginResponse).Token, res1
}

func (set EndpointsSet) PublicKey(arg0 context.Context) (res0 []byte, res1 error) {
	request := PublicKeyRequest{}
	response, res1 := set.PublicKeyEndpoint(arg0, &request)
	if res1 != nil {
		if e, ok := status.FromError(res1); ok || e.Code() == codes.Internal || e.Code() == codes.Unknown {
			res1 = errors.New(e.Message())
		}
		return
	}
	return response.(*PublicKeyResponse).Pub, res1
}

func (set EndpointsSet) GetPermissionsForService(arg0 context.Context, arg1 string) (res0 []*types.Permission, res1 error) {
	request := GetPermissionsForServiceRequest{Name: arg1}
	response, res1 := set.GetPermissionsForServiceEndpoint(arg0, &request)
	if res1 != nil {
		if e, ok := status.FromError(res1); ok || e.Code() == codes.Internal || e.Code() == codes.Unknown {
			res1 = errors.New(e.Message())
		}
		return
	}
	return response.(*GetPermissionsForServiceResponse).Permissions, res1
}
