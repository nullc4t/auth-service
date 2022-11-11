// Code generated by microgen 1.0.5. DO NOT EDIT.

// Please, do not change functions names!
package transportgrpc

import (
	pb "auth/mgmt/proto"
	transport "auth/mgmt/transport"
	"context"
	"errors"
	empty "github.com/golang/protobuf/ptypes/empty"
)

func _Encode_CreateService_Request(ctx context.Context, request interface{}) (interface{}, error) {
	if request == nil {
		return nil, errors.New("nil CreateServiceRequest")
	}
	req := request.(*transport.CreateServiceRequest)
	return &pb.CreateServiceRequest{Name: req.Name}, nil
}

func _Encode_GetAllServices_Request(ctx context.Context, request interface{}) (interface{}, error) {
	return &empty.Empty{}, nil
}

func _Encode_GetService_Request(ctx context.Context, request interface{}) (interface{}, error) {
	if request == nil {
		return nil, errors.New("nil GetServiceRequest")
	}
	req := request.(*transport.GetServiceRequest)
	reqSvc, err := PtrTypesServiceToProto(req.Svc)
	if err != nil {
		return nil, err
	}
	return &pb.GetServiceRequest{Svc: reqSvc}, nil
}

func _Encode_CreateAccount_Request(ctx context.Context, request interface{}) (interface{}, error) {
	return &empty.Empty{}, nil
}

func _Encode_CreateAccountWithName_Request(ctx context.Context, request interface{}) (interface{}, error) {
	if request == nil {
		return nil, errors.New("nil CreateAccountWithNameRequest")
	}
	req := request.(*transport.CreateAccountWithNameRequest)
	return &pb.CreateAccountWithNameRequest{Name: req.Name}, nil
}

func _Encode_GetAllAccounts_Request(ctx context.Context, request interface{}) (interface{}, error) {
	return &empty.Empty{}, nil
}

func _Encode_GetAccount_Request(ctx context.Context, request interface{}) (interface{}, error) {
	if request == nil {
		return nil, errors.New("nil GetAccountRequest")
	}
	req := request.(*transport.GetAccountRequest)
	reqAcc, err := PtrTypesAccountToProto(req.Acc)
	if err != nil {
		return nil, err
	}
	return &pb.GetAccountRequest{Acc: reqAcc}, nil
}

func _Encode_UpdateAccount_Request(ctx context.Context, request interface{}) (interface{}, error) {
	if request == nil {
		return nil, errors.New("nil UpdateAccountRequest")
	}
	req := request.(*transport.UpdateAccountRequest)
	reqAcc, err := PtrTypesAccountToProto(req.Acc)
	if err != nil {
		return nil, err
	}
	return &pb.UpdateAccountRequest{Acc: reqAcc}, nil
}

func _Encode_AttachAccountToService_Request(ctx context.Context, request interface{}) (interface{}, error) {
	if request == nil {
		return nil, errors.New("nil AttachAccountToServiceRequest")
	}
	req := request.(*transport.AttachAccountToServiceRequest)
	return &pb.AttachAccountToServiceRequest{
		AccountId: req.AccountId,
		ServiceId: req.ServiceId,
	}, nil
}

func _Encode_RemoveAccountFromService_Request(ctx context.Context, request interface{}) (interface{}, error) {
	if request == nil {
		return nil, errors.New("nil RemoveAccountFromServiceRequest")
	}
	req := request.(*transport.RemoveAccountFromServiceRequest)
	return &pb.RemoveAccountFromServiceRequest{
		AccountId: req.AccountId,
		ServiceId: req.ServiceId,
	}, nil
}

func _Encode_CreatePermission_Request(ctx context.Context, request interface{}) (interface{}, error) {
	if request == nil {
		return nil, errors.New("nil CreatePermissionRequest")
	}
	req := request.(*transport.CreatePermissionRequest)
	reqAccess, err := PtrAccessAccessToProto(req.Access)
	if err != nil {
		return nil, err
	}
	return &pb.CreatePermissionRequest{
		Access:    reqAccess,
		Name:      req.Name,
		ServiceId: req.ServiceId,
	}, nil
}

func _Encode_GetPermission_Request(ctx context.Context, request interface{}) (interface{}, error) {
	if request == nil {
		return nil, errors.New("nil GetPermissionRequest")
	}
	req := request.(*transport.GetPermissionRequest)
	reqP, err := PtrTypesPermissionToProto(req.P)
	if err != nil {
		return nil, err
	}
	return &pb.GetPermissionRequest{P: reqP}, nil
}

func _Encode_GetAllPermission_Request(ctx context.Context, request interface{}) (interface{}, error) {
	return &empty.Empty{}, nil
}

func _Encode_GetFilteredPermissions_Request(ctx context.Context, request interface{}) (interface{}, error) {
	if request == nil {
		return nil, errors.New("nil GetFilteredPermissionsRequest")
	}
	req := request.(*transport.GetFilteredPermissionsRequest)
	reqP, err := PtrTypesPermissionToProto(req.P)
	if err != nil {
		return nil, err
	}
	return &pb.GetFilteredPermissionsRequest{P: reqP}, nil
}

func _Encode_DeletePermission_Request(ctx context.Context, request interface{}) (interface{}, error) {
	if request == nil {
		return nil, errors.New("nil DeletePermissionRequest")
	}
	req := request.(*transport.DeletePermissionRequest)
	reqP, err := PtrTypesPermissionToProto(req.P)
	if err != nil {
		return nil, err
	}
	return &pb.DeletePermissionRequest{P: reqP}, nil
}

func _Encode_GetUserPermissions_Request(ctx context.Context, request interface{}) (interface{}, error) {
	if request == nil {
		return nil, errors.New("nil GetUserPermissionsRequest")
	}
	req := request.(*transport.GetUserPermissionsRequest)
	return &pb.GetUserPermissionsRequest{UserId: req.UserId}, nil
}

func _Encode_AddUserPermission_Request(ctx context.Context, request interface{}) (interface{}, error) {
	if request == nil {
		return nil, errors.New("nil AddUserPermissionRequest")
	}
	req := request.(*transport.AddUserPermissionRequest)
	reqP, err := PtrTypesPermissionToProto(req.P)
	if err != nil {
		return nil, err
	}
	return &pb.AddUserPermissionRequest{
		P:      reqP,
		UserId: req.UserId,
	}, nil
}

func _Encode_RemoveUserPermission_Request(ctx context.Context, request interface{}) (interface{}, error) {
	if request == nil {
		return nil, errors.New("nil RemoveUserPermissionRequest")
	}
	req := request.(*transport.RemoveUserPermissionRequest)
	return &pb.RemoveUserPermissionRequest{
		PermId: req.PermId,
		UserId: req.UserId,
	}, nil
}

func _Encode_CreateService_Response(ctx context.Context, response interface{}) (interface{}, error) {
	if response == nil {
		return nil, errors.New("nil CreateServiceResponse")
	}
	resp := response.(*transport.CreateServiceResponse)
	respS, err := PtrTypesServiceToProto(resp.S)
	if err != nil {
		return nil, err
	}
	return &pb.CreateServiceResponse{S: respS}, nil
}

func _Encode_GetAllServices_Response(ctx context.Context, response interface{}) (interface{}, error) {
	if response == nil {
		return nil, errors.New("nil GetAllServicesResponse")
	}
	resp := response.(*transport.GetAllServicesResponse)
	respSs, err := ListPtrTypesServiceToProto(resp.Ss)
	if err != nil {
		return nil, err
	}
	return &pb.GetAllServicesResponse{Ss: respSs}, nil
}

func _Encode_GetService_Response(ctx context.Context, response interface{}) (interface{}, error) {
	if response == nil {
		return nil, errors.New("nil GetServiceResponse")
	}
	resp := response.(*transport.GetServiceResponse)
	respS, err := PtrTypesServiceToProto(resp.S)
	if err != nil {
		return nil, err
	}
	return &pb.GetServiceResponse{S: respS}, nil
}

func _Encode_CreateAccount_Response(ctx context.Context, response interface{}) (interface{}, error) {
	if response == nil {
		return nil, errors.New("nil CreateAccountResponse")
	}
	resp := response.(*transport.CreateAccountResponse)
	respA, err := PtrTypesAccountToProto(resp.A)
	if err != nil {
		return nil, err
	}
	return &pb.CreateAccountResponse{A: respA}, nil
}

func _Encode_CreateAccountWithName_Response(ctx context.Context, response interface{}) (interface{}, error) {
	if response == nil {
		return nil, errors.New("nil CreateAccountWithNameResponse")
	}
	resp := response.(*transport.CreateAccountWithNameResponse)
	respA, err := PtrTypesAccountToProto(resp.A)
	if err != nil {
		return nil, err
	}
	return &pb.CreateAccountWithNameResponse{A: respA}, nil
}

func _Encode_GetAllAccounts_Response(ctx context.Context, response interface{}) (interface{}, error) {
	if response == nil {
		return nil, errors.New("nil GetAllAccountsResponse")
	}
	resp := response.(*transport.GetAllAccountsResponse)
	respAs, err := ListPtrTypesAccountToProto(resp.As)
	if err != nil {
		return nil, err
	}
	return &pb.GetAllAccountsResponse{As: respAs}, nil
}

func _Encode_GetAccount_Response(ctx context.Context, response interface{}) (interface{}, error) {
	if response == nil {
		return nil, errors.New("nil GetAccountResponse")
	}
	resp := response.(*transport.GetAccountResponse)
	respA, err := PtrTypesAccountToProto(resp.A)
	if err != nil {
		return nil, err
	}
	return &pb.GetAccountResponse{A: respA}, nil
}

func _Encode_UpdateAccount_Response(ctx context.Context, response interface{}) (interface{}, error) {
	if response == nil {
		return nil, errors.New("nil UpdateAccountResponse")
	}
	resp := response.(*transport.UpdateAccountResponse)
	respA, err := PtrTypesAccountToProto(resp.A)
	if err != nil {
		return nil, err
	}
	return &pb.UpdateAccountResponse{A: respA}, nil
}

func _Encode_AttachAccountToService_Response(ctx context.Context, response interface{}) (interface{}, error) {
	if response == nil {
		return nil, errors.New("nil AttachAccountToServiceResponse")
	}
	resp := response.(*transport.AttachAccountToServiceResponse)
	return &pb.AttachAccountToServiceResponse{Ok: resp.Ok}, nil
}

func _Encode_RemoveAccountFromService_Response(ctx context.Context, response interface{}) (interface{}, error) {
	if response == nil {
		return nil, errors.New("nil RemoveAccountFromServiceResponse")
	}
	resp := response.(*transport.RemoveAccountFromServiceResponse)
	return &pb.RemoveAccountFromServiceResponse{Ok: resp.Ok}, nil
}

func _Encode_CreatePermission_Response(ctx context.Context, response interface{}) (interface{}, error) {
	if response == nil {
		return nil, errors.New("nil CreatePermissionResponse")
	}
	resp := response.(*transport.CreatePermissionResponse)
	respP, err := PtrTypesPermissionToProto(resp.P)
	if err != nil {
		return nil, err
	}
	return &pb.CreatePermissionResponse{P: respP}, nil
}

func _Encode_GetPermission_Response(ctx context.Context, response interface{}) (interface{}, error) {
	if response == nil {
		return nil, errors.New("nil GetPermissionResponse")
	}
	resp := response.(*transport.GetPermissionResponse)
	respPerm, err := PtrTypesPermissionToProto(resp.Perm)
	if err != nil {
		return nil, err
	}
	return &pb.GetPermissionResponse{Perm: respPerm}, nil
}

func _Encode_GetAllPermission_Response(ctx context.Context, response interface{}) (interface{}, error) {
	if response == nil {
		return nil, errors.New("nil GetAllPermissionResponse")
	}
	resp := response.(*transport.GetAllPermissionResponse)
	respP, err := ListPtrTypesPermissionToProto(resp.P)
	if err != nil {
		return nil, err
	}
	return &pb.GetAllPermissionResponse{P: respP}, nil
}

func _Encode_GetFilteredPermissions_Response(ctx context.Context, response interface{}) (interface{}, error) {
	if response == nil {
		return nil, errors.New("nil GetFilteredPermissionsResponse")
	}
	resp := response.(*transport.GetFilteredPermissionsResponse)
	respPerm, err := ListPtrTypesPermissionToProto(resp.Perm)
	if err != nil {
		return nil, err
	}
	return &pb.GetFilteredPermissionsResponse{Perm: respPerm}, nil
}

func _Encode_DeletePermission_Response(ctx context.Context, response interface{}) (interface{}, error) {
	if response == nil {
		return nil, errors.New("nil DeletePermissionResponse")
	}
	resp := response.(*transport.DeletePermissionResponse)
	return &pb.DeletePermissionResponse{Ok: resp.Ok}, nil
}

func _Encode_GetUserPermissions_Response(ctx context.Context, response interface{}) (interface{}, error) {
	if response == nil {
		return nil, errors.New("nil GetUserPermissionsResponse")
	}
	resp := response.(*transport.GetUserPermissionsResponse)
	respPermissions, err := ListPtrTypesPermissionToProto(resp.Permissions)
	if err != nil {
		return nil, err
	}
	return &pb.GetUserPermissionsResponse{Permissions: respPermissions}, nil
}

func _Encode_AddUserPermission_Response(ctx context.Context, response interface{}) (interface{}, error) {
	if response == nil {
		return nil, errors.New("nil AddUserPermissionResponse")
	}
	resp := response.(*transport.AddUserPermissionResponse)
	return &pb.AddUserPermissionResponse{Ok: resp.Ok}, nil
}

func _Encode_RemoveUserPermission_Response(ctx context.Context, response interface{}) (interface{}, error) {
	if response == nil {
		return nil, errors.New("nil RemoveUserPermissionResponse")
	}
	resp := response.(*transport.RemoveUserPermissionResponse)
	return &pb.RemoveUserPermissionResponse{Ok: resp.Ok}, nil
}

func _Decode_CreateService_Request(ctx context.Context, request interface{}) (interface{}, error) {
	if request == nil {
		return nil, errors.New("nil CreateServiceRequest")
	}
	req := request.(*pb.CreateServiceRequest)
	return &transport.CreateServiceRequest{Name: string(req.Name)}, nil
}

func _Decode_GetAllServices_Request(ctx context.Context, request interface{}) (interface{}, error) {
	return &empty.Empty{}, nil
}

func _Decode_GetService_Request(ctx context.Context, request interface{}) (interface{}, error) {
	if request == nil {
		return nil, errors.New("nil GetServiceRequest")
	}
	req := request.(*pb.GetServiceRequest)
	reqSvc, err := ProtoToPtrTypesService(req.Svc)
	if err != nil {
		return nil, err
	}
	return &transport.GetServiceRequest{Svc: reqSvc}, nil
}

func _Decode_CreateAccount_Request(ctx context.Context, request interface{}) (interface{}, error) {
	return &empty.Empty{}, nil
}

func _Decode_CreateAccountWithName_Request(ctx context.Context, request interface{}) (interface{}, error) {
	if request == nil {
		return nil, errors.New("nil CreateAccountWithNameRequest")
	}
	req := request.(*pb.CreateAccountWithNameRequest)
	return &transport.CreateAccountWithNameRequest{Name: string(req.Name)}, nil
}

func _Decode_GetAllAccounts_Request(ctx context.Context, request interface{}) (interface{}, error) {
	return &empty.Empty{}, nil
}

func _Decode_GetAccount_Request(ctx context.Context, request interface{}) (interface{}, error) {
	if request == nil {
		return nil, errors.New("nil GetAccountRequest")
	}
	req := request.(*pb.GetAccountRequest)
	reqAcc, err := ProtoToPtrTypesAccount(req.Acc)
	if err != nil {
		return nil, err
	}
	return &transport.GetAccountRequest{Acc: reqAcc}, nil
}

func _Decode_UpdateAccount_Request(ctx context.Context, request interface{}) (interface{}, error) {
	if request == nil {
		return nil, errors.New("nil UpdateAccountRequest")
	}
	req := request.(*pb.UpdateAccountRequest)
	reqAcc, err := ProtoToPtrTypesAccount(req.Acc)
	if err != nil {
		return nil, err
	}
	return &transport.UpdateAccountRequest{Acc: reqAcc}, nil
}

func _Decode_AttachAccountToService_Request(ctx context.Context, request interface{}) (interface{}, error) {
	if request == nil {
		return nil, errors.New("nil AttachAccountToServiceRequest")
	}
	req := request.(*pb.AttachAccountToServiceRequest)
	return &transport.AttachAccountToServiceRequest{
		AccountId: uint32(req.AccountId),
		ServiceId: uint32(req.ServiceId),
	}, nil
}

func _Decode_RemoveAccountFromService_Request(ctx context.Context, request interface{}) (interface{}, error) {
	if request == nil {
		return nil, errors.New("nil RemoveAccountFromServiceRequest")
	}
	req := request.(*pb.RemoveAccountFromServiceRequest)
	return &transport.RemoveAccountFromServiceRequest{
		AccountId: uint32(req.AccountId),
		ServiceId: uint32(req.ServiceId),
	}, nil
}

func _Decode_CreatePermission_Request(ctx context.Context, request interface{}) (interface{}, error) {
	if request == nil {
		return nil, errors.New("nil CreatePermissionRequest")
	}
	req := request.(*pb.CreatePermissionRequest)
	reqAccess, err := ProtoToPtrAccessAccess(req.Access)
	if err != nil {
		return nil, err
	}
	return &transport.CreatePermissionRequest{
		Access:    reqAccess,
		Name:      string(req.Name),
		ServiceId: uint32(req.ServiceId),
	}, nil
}

func _Decode_GetPermission_Request(ctx context.Context, request interface{}) (interface{}, error) {
	if request == nil {
		return nil, errors.New("nil GetPermissionRequest")
	}
	req := request.(*pb.GetPermissionRequest)
	reqP, err := ProtoToPtrTypesPermission(req.P)
	if err != nil {
		return nil, err
	}
	return &transport.GetPermissionRequest{P: reqP}, nil
}

func _Decode_GetAllPermission_Request(ctx context.Context, request interface{}) (interface{}, error) {
	return &empty.Empty{}, nil
}

func _Decode_GetFilteredPermissions_Request(ctx context.Context, request interface{}) (interface{}, error) {
	if request == nil {
		return nil, errors.New("nil GetFilteredPermissionsRequest")
	}
	req := request.(*pb.GetFilteredPermissionsRequest)
	reqP, err := ProtoToPtrTypesPermission(req.P)
	if err != nil {
		return nil, err
	}
	return &transport.GetFilteredPermissionsRequest{P: reqP}, nil
}

func _Decode_DeletePermission_Request(ctx context.Context, request interface{}) (interface{}, error) {
	if request == nil {
		return nil, errors.New("nil DeletePermissionRequest")
	}
	req := request.(*pb.DeletePermissionRequest)
	reqP, err := ProtoToPtrTypesPermission(req.P)
	if err != nil {
		return nil, err
	}
	return &transport.DeletePermissionRequest{P: reqP}, nil
}

func _Decode_GetUserPermissions_Request(ctx context.Context, request interface{}) (interface{}, error) {
	if request == nil {
		return nil, errors.New("nil GetUserPermissionsRequest")
	}
	req := request.(*pb.GetUserPermissionsRequest)
	return &transport.GetUserPermissionsRequest{UserId: uint32(req.UserId)}, nil
}

func _Decode_AddUserPermission_Request(ctx context.Context, request interface{}) (interface{}, error) {
	if request == nil {
		return nil, errors.New("nil AddUserPermissionRequest")
	}
	req := request.(*pb.AddUserPermissionRequest)
	reqP, err := ProtoToPtrTypesPermission(req.P)
	if err != nil {
		return nil, err
	}
	return &transport.AddUserPermissionRequest{
		P:      reqP,
		UserId: uint32(req.UserId),
	}, nil
}

func _Decode_RemoveUserPermission_Request(ctx context.Context, request interface{}) (interface{}, error) {
	if request == nil {
		return nil, errors.New("nil RemoveUserPermissionRequest")
	}
	req := request.(*pb.RemoveUserPermissionRequest)
	return &transport.RemoveUserPermissionRequest{
		PermId: uint32(req.PermId),
		UserId: uint32(req.UserId),
	}, nil
}

func _Decode_CreateService_Response(ctx context.Context, response interface{}) (interface{}, error) {
	if response == nil {
		return nil, errors.New("nil CreateServiceResponse")
	}
	resp := response.(*pb.CreateServiceResponse)
	respS, err := ProtoToPtrTypesService(resp.S)
	if err != nil {
		return nil, err
	}
	return &transport.CreateServiceResponse{S: respS}, nil
}

func _Decode_GetAllServices_Response(ctx context.Context, response interface{}) (interface{}, error) {
	if response == nil {
		return nil, errors.New("nil GetAllServicesResponse")
	}
	resp := response.(*pb.GetAllServicesResponse)
	respSs, err := ProtoToListPtrTypesService(resp.Ss)
	if err != nil {
		return nil, err
	}
	return &transport.GetAllServicesResponse{Ss: respSs}, nil
}

func _Decode_GetService_Response(ctx context.Context, response interface{}) (interface{}, error) {
	if response == nil {
		return nil, errors.New("nil GetServiceResponse")
	}
	resp := response.(*pb.GetServiceResponse)
	respS, err := ProtoToPtrTypesService(resp.S)
	if err != nil {
		return nil, err
	}
	return &transport.GetServiceResponse{S: respS}, nil
}

func _Decode_CreateAccount_Response(ctx context.Context, response interface{}) (interface{}, error) {
	if response == nil {
		return nil, errors.New("nil CreateAccountResponse")
	}
	resp := response.(*pb.CreateAccountResponse)
	respA, err := ProtoToPtrTypesAccount(resp.A)
	if err != nil {
		return nil, err
	}
	return &transport.CreateAccountResponse{A: respA}, nil
}

func _Decode_CreateAccountWithName_Response(ctx context.Context, response interface{}) (interface{}, error) {
	if response == nil {
		return nil, errors.New("nil CreateAccountWithNameResponse")
	}
	resp := response.(*pb.CreateAccountWithNameResponse)
	respA, err := ProtoToPtrTypesAccount(resp.A)
	if err != nil {
		return nil, err
	}
	return &transport.CreateAccountWithNameResponse{A: respA}, nil
}

func _Decode_GetAllAccounts_Response(ctx context.Context, response interface{}) (interface{}, error) {
	if response == nil {
		return nil, errors.New("nil GetAllAccountsResponse")
	}
	resp := response.(*pb.GetAllAccountsResponse)
	respAs, err := ProtoToListPtrTypesAccount(resp.As)
	if err != nil {
		return nil, err
	}
	return &transport.GetAllAccountsResponse{As: respAs}, nil
}

func _Decode_GetAccount_Response(ctx context.Context, response interface{}) (interface{}, error) {
	if response == nil {
		return nil, errors.New("nil GetAccountResponse")
	}
	resp := response.(*pb.GetAccountResponse)
	respA, err := ProtoToPtrTypesAccount(resp.A)
	if err != nil {
		return nil, err
	}
	return &transport.GetAccountResponse{A: respA}, nil
}

func _Decode_UpdateAccount_Response(ctx context.Context, response interface{}) (interface{}, error) {
	if response == nil {
		return nil, errors.New("nil UpdateAccountResponse")
	}
	resp := response.(*pb.UpdateAccountResponse)
	respA, err := ProtoToPtrTypesAccount(resp.A)
	if err != nil {
		return nil, err
	}
	return &transport.UpdateAccountResponse{A: respA}, nil
}

func _Decode_AttachAccountToService_Response(ctx context.Context, response interface{}) (interface{}, error) {
	if response == nil {
		return nil, errors.New("nil AttachAccountToServiceResponse")
	}
	resp := response.(*pb.AttachAccountToServiceResponse)
	return &transport.AttachAccountToServiceResponse{Ok: bool(resp.Ok)}, nil
}

func _Decode_RemoveAccountFromService_Response(ctx context.Context, response interface{}) (interface{}, error) {
	if response == nil {
		return nil, errors.New("nil RemoveAccountFromServiceResponse")
	}
	resp := response.(*pb.RemoveAccountFromServiceResponse)
	return &transport.RemoveAccountFromServiceResponse{Ok: bool(resp.Ok)}, nil
}

func _Decode_CreatePermission_Response(ctx context.Context, response interface{}) (interface{}, error) {
	if response == nil {
		return nil, errors.New("nil CreatePermissionResponse")
	}
	resp := response.(*pb.CreatePermissionResponse)
	respP, err := ProtoToPtrTypesPermission(resp.P)
	if err != nil {
		return nil, err
	}
	return &transport.CreatePermissionResponse{P: respP}, nil
}

func _Decode_GetPermission_Response(ctx context.Context, response interface{}) (interface{}, error) {
	if response == nil {
		return nil, errors.New("nil GetPermissionResponse")
	}
	resp := response.(*pb.GetPermissionResponse)
	respPerm, err := ProtoToPtrTypesPermission(resp.Perm)
	if err != nil {
		return nil, err
	}
	return &transport.GetPermissionResponse{Perm: respPerm}, nil
}

func _Decode_GetAllPermission_Response(ctx context.Context, response interface{}) (interface{}, error) {
	if response == nil {
		return nil, errors.New("nil GetAllPermissionResponse")
	}
	resp := response.(*pb.GetAllPermissionResponse)
	respP, err := ProtoToListPtrTypesPermission(resp.P)
	if err != nil {
		return nil, err
	}
	return &transport.GetAllPermissionResponse{P: respP}, nil
}

func _Decode_GetFilteredPermissions_Response(ctx context.Context, response interface{}) (interface{}, error) {
	if response == nil {
		return nil, errors.New("nil GetFilteredPermissionsResponse")
	}
	resp := response.(*pb.GetFilteredPermissionsResponse)
	respPerm, err := ProtoToListPtrTypesPermission(resp.Perm)
	if err != nil {
		return nil, err
	}
	return &transport.GetFilteredPermissionsResponse{Perm: respPerm}, nil
}

func _Decode_DeletePermission_Response(ctx context.Context, response interface{}) (interface{}, error) {
	if response == nil {
		return nil, errors.New("nil DeletePermissionResponse")
	}
	resp := response.(*pb.DeletePermissionResponse)
	return &transport.DeletePermissionResponse{Ok: bool(resp.Ok)}, nil
}

func _Decode_GetUserPermissions_Response(ctx context.Context, response interface{}) (interface{}, error) {
	if response == nil {
		return nil, errors.New("nil GetUserPermissionsResponse")
	}
	resp := response.(*pb.GetUserPermissionsResponse)
	respPermissions, err := ProtoToListPtrTypesPermission(resp.Permissions)
	if err != nil {
		return nil, err
	}
	return &transport.GetUserPermissionsResponse{Permissions: respPermissions}, nil
}

func _Decode_AddUserPermission_Response(ctx context.Context, response interface{}) (interface{}, error) {
	if response == nil {
		return nil, errors.New("nil AddUserPermissionResponse")
	}
	resp := response.(*pb.AddUserPermissionResponse)
	return &transport.AddUserPermissionResponse{Ok: bool(resp.Ok)}, nil
}

func _Decode_RemoveUserPermission_Response(ctx context.Context, response interface{}) (interface{}, error) {
	if response == nil {
		return nil, errors.New("nil RemoveUserPermissionResponse")
	}
	resp := response.(*pb.RemoveUserPermissionResponse)
	return &transport.RemoveUserPermissionResponse{Ok: bool(resp.Ok)}, nil
}
