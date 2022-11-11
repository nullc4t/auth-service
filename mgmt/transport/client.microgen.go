// Code generated by microgen 1.0.5. DO NOT EDIT.

package transport

import (
	access "auth/access"
	types "auth/pkg/types"
	"context"
	"errors"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
)

func (set EndpointsSet) CreateUserWithLoginPassword(arg0 context.Context, arg1 string, arg2 string) (res0 *types.User, res1 error) {
	request := CreateUserWithLoginPasswordRequest{
		Login: arg1,
		Pass:  arg2,
	}
	response, res1 := set.CreateUserWithLoginPasswordEndpoint(arg0, &request)
	if res1 != nil {
		if e, ok := status.FromError(res1); ok || e.Code() == codes.Internal || e.Code() == codes.Unknown {
			res1 = errors.New(e.Message())
		}
		return
	}
	return response.(*CreateUserWithLoginPasswordResponse).User, res1
}

func (set EndpointsSet) CreateUserWithTelegram(arg0 context.Context, arg1 uint64, arg2 string, arg3 string) (res0 *types.User, res1 error) {
	request := CreateUserWithTelegramRequest{
		Id:    arg1,
		Name:  arg2,
		UserN: arg3,
	}
	response, res1 := set.CreateUserWithTelegramEndpoint(arg0, &request)
	if res1 != nil {
		if e, ok := status.FromError(res1); ok || e.Code() == codes.Internal || e.Code() == codes.Unknown {
			res1 = errors.New(e.Message())
		}
		return
	}
	return response.(*CreateUserWithTelegramResponse).User, res1
}

func (set EndpointsSet) GetAllUsers(arg0 context.Context) (res0 []*types.User, res1 error) {
	request := GetAllUsersRequest{}
	response, res1 := set.GetAllUsersEndpoint(arg0, &request)
	if res1 != nil {
		if e, ok := status.FromError(res1); ok || e.Code() == codes.Internal || e.Code() == codes.Unknown {
			res1 = errors.New(e.Message())
		}
		return
	}
	return response.(*GetAllUsersResponse).Users, res1
}

func (set EndpointsSet) GetUser(arg0 context.Context, arg1 *types.User) (res0 *types.User, res1 error) {
	request := GetUserRequest{UserReq: arg1}
	response, res1 := set.GetUserEndpoint(arg0, &request)
	if res1 != nil {
		if e, ok := status.FromError(res1); ok || e.Code() == codes.Internal || e.Code() == codes.Unknown {
			res1 = errors.New(e.Message())
		}
		return
	}
	return response.(*GetUserResponse).User, res1
}

func (set EndpointsSet) UpdateUser(arg0 context.Context, arg1 *types.User) (res0 *types.User, res1 error) {
	request := UpdateUserRequest{UserReq: arg1}
	response, res1 := set.UpdateUserEndpoint(arg0, &request)
	if res1 != nil {
		if e, ok := status.FromError(res1); ok || e.Code() == codes.Internal || e.Code() == codes.Unknown {
			res1 = errors.New(e.Message())
		}
		return
	}
	return response.(*UpdateUserResponse).User, res1
}

func (set EndpointsSet) BlockUser(arg0 context.Context, arg1 uint32) (res0 bool, res1 error) {
	request := BlockUserRequest{UserId: arg1}
	response, res1 := set.BlockUserEndpoint(arg0, &request)
	if res1 != nil {
		if e, ok := status.FromError(res1); ok || e.Code() == codes.Internal || e.Code() == codes.Unknown {
			res1 = errors.New(e.Message())
		}
		return
	}
	return response.(*BlockUserResponse).Ok, res1
}

func (set EndpointsSet) UnblockUser(arg0 context.Context, arg1 uint32) (res0 bool, res1 error) {
	request := UnblockUserRequest{UserId: arg1}
	response, res1 := set.UnblockUserEndpoint(arg0, &request)
	if res1 != nil {
		if e, ok := status.FromError(res1); ok || e.Code() == codes.Internal || e.Code() == codes.Unknown {
			res1 = errors.New(e.Message())
		}
		return
	}
	return response.(*UnblockUserResponse).Ok, res1
}

func (set EndpointsSet) CreateService(arg0 context.Context, arg1 string) (res0 *types.Service, res1 error) {
	request := CreateServiceRequest{Name: arg1}
	response, res1 := set.CreateServiceEndpoint(arg0, &request)
	if res1 != nil {
		if e, ok := status.FromError(res1); ok || e.Code() == codes.Internal || e.Code() == codes.Unknown {
			res1 = errors.New(e.Message())
		}
		return
	}
	return response.(*CreateServiceResponse).S, res1
}

func (set EndpointsSet) GetAllServices(arg0 context.Context) (res0 []*types.Service, res1 error) {
	request := GetAllServicesRequest{}
	response, res1 := set.GetAllServicesEndpoint(arg0, &request)
	if res1 != nil {
		if e, ok := status.FromError(res1); ok || e.Code() == codes.Internal || e.Code() == codes.Unknown {
			res1 = errors.New(e.Message())
		}
		return
	}
	return response.(*GetAllServicesResponse).Ss, res1
}

func (set EndpointsSet) GetService(arg0 context.Context, arg1 *types.Service) (res0 *types.Service, res1 error) {
	request := GetServiceRequest{Svc: arg1}
	response, res1 := set.GetServiceEndpoint(arg0, &request)
	if res1 != nil {
		if e, ok := status.FromError(res1); ok || e.Code() == codes.Internal || e.Code() == codes.Unknown {
			res1 = errors.New(e.Message())
		}
		return
	}
	return response.(*GetServiceResponse).S, res1
}

func (set EndpointsSet) CreateAccount(arg0 context.Context) (res0 *types.Account, res1 error) {
	request := CreateAccountRequest{}
	response, res1 := set.CreateAccountEndpoint(arg0, &request)
	if res1 != nil {
		if e, ok := status.FromError(res1); ok || e.Code() == codes.Internal || e.Code() == codes.Unknown {
			res1 = errors.New(e.Message())
		}
		return
	}
	return response.(*CreateAccountResponse).A, res1
}

func (set EndpointsSet) CreateAccountWithName(arg0 context.Context, arg1 string) (res0 *types.Account, res1 error) {
	request := CreateAccountWithNameRequest{Name: arg1}
	response, res1 := set.CreateAccountWithNameEndpoint(arg0, &request)
	if res1 != nil {
		if e, ok := status.FromError(res1); ok || e.Code() == codes.Internal || e.Code() == codes.Unknown {
			res1 = errors.New(e.Message())
		}
		return
	}
	return response.(*CreateAccountWithNameResponse).A, res1
}

func (set EndpointsSet) GetAllAccounts(arg0 context.Context) (res0 []*types.Account, res1 error) {
	request := GetAllAccountsRequest{}
	response, res1 := set.GetAllAccountsEndpoint(arg0, &request)
	if res1 != nil {
		if e, ok := status.FromError(res1); ok || e.Code() == codes.Internal || e.Code() == codes.Unknown {
			res1 = errors.New(e.Message())
		}
		return
	}
	return response.(*GetAllAccountsResponse).As, res1
}

func (set EndpointsSet) GetAccount(arg0 context.Context, arg1 *types.Account) (res0 *types.Account, res1 error) {
	request := GetAccountRequest{Acc: arg1}
	response, res1 := set.GetAccountEndpoint(arg0, &request)
	if res1 != nil {
		if e, ok := status.FromError(res1); ok || e.Code() == codes.Internal || e.Code() == codes.Unknown {
			res1 = errors.New(e.Message())
		}
		return
	}
	return response.(*GetAccountResponse).A, res1
}

func (set EndpointsSet) UpdateAccount(arg0 context.Context, arg1 *types.Account) (res0 *types.Account, res1 error) {
	request := UpdateAccountRequest{Acc: arg1}
	response, res1 := set.UpdateAccountEndpoint(arg0, &request)
	if res1 != nil {
		if e, ok := status.FromError(res1); ok || e.Code() == codes.Internal || e.Code() == codes.Unknown {
			res1 = errors.New(e.Message())
		}
		return
	}
	return response.(*UpdateAccountResponse).A, res1
}

func (set EndpointsSet) AttachUserToAccount(arg0 context.Context, arg1 uint32, arg2 uint32) (res0 bool, res1 error) {
	request := AttachUserToAccountRequest{
		AccountId: arg2,
		UserId:    arg1,
	}
	response, res1 := set.AttachUserToAccountEndpoint(arg0, &request)
	if res1 != nil {
		if e, ok := status.FromError(res1); ok || e.Code() == codes.Internal || e.Code() == codes.Unknown {
			res1 = errors.New(e.Message())
		}
		return
	}
	return response.(*AttachUserToAccountResponse).Ok, res1
}

func (set EndpointsSet) AttachAccountToService(arg0 context.Context, arg1 uint32, arg2 uint32) (res0 bool, res1 error) {
	request := AttachAccountToServiceRequest{
		AccountId: arg2,
		ServiceId: arg1,
	}
	response, res1 := set.AttachAccountToServiceEndpoint(arg0, &request)
	if res1 != nil {
		if e, ok := status.FromError(res1); ok || e.Code() == codes.Internal || e.Code() == codes.Unknown {
			res1 = errors.New(e.Message())
		}
		return
	}
	return response.(*AttachAccountToServiceResponse).Ok, res1
}

func (set EndpointsSet) RemoveAccountFromService(arg0 context.Context, arg1 uint32, arg2 uint32) (res0 bool, res1 error) {
	request := RemoveAccountFromServiceRequest{
		AccountId: arg2,
		ServiceId: arg1,
	}
	response, res1 := set.RemoveAccountFromServiceEndpoint(arg0, &request)
	if res1 != nil {
		if e, ok := status.FromError(res1); ok || e.Code() == codes.Internal || e.Code() == codes.Unknown {
			res1 = errors.New(e.Message())
		}
		return
	}
	return response.(*RemoveAccountFromServiceResponse).Ok, res1
}

func (set EndpointsSet) CreatePermission(arg0 context.Context, arg1 uint32, arg2 string, arg3 *access.Access) (res0 *types.Permission, res1 error) {
	request := CreatePermissionRequest{
		Access:    arg3,
		Name:      arg2,
		ServiceId: arg1,
	}
	response, res1 := set.CreatePermissionEndpoint(arg0, &request)
	if res1 != nil {
		if e, ok := status.FromError(res1); ok || e.Code() == codes.Internal || e.Code() == codes.Unknown {
			res1 = errors.New(e.Message())
		}
		return
	}
	return response.(*CreatePermissionResponse).P, res1
}

func (set EndpointsSet) GetPermission(arg0 context.Context, arg1 *types.Permission) (res0 *types.Permission, res1 error) {
	request := GetPermissionRequest{P: arg1}
	response, res1 := set.GetPermissionEndpoint(arg0, &request)
	if res1 != nil {
		if e, ok := status.FromError(res1); ok || e.Code() == codes.Internal || e.Code() == codes.Unknown {
			res1 = errors.New(e.Message())
		}
		return
	}
	return response.(*GetPermissionResponse).Perm, res1
}

func (set EndpointsSet) GetAllPermission(arg0 context.Context) (res0 []*types.Permission, res1 error) {
	request := GetAllPermissionRequest{}
	response, res1 := set.GetAllPermissionEndpoint(arg0, &request)
	if res1 != nil {
		if e, ok := status.FromError(res1); ok || e.Code() == codes.Internal || e.Code() == codes.Unknown {
			res1 = errors.New(e.Message())
		}
		return
	}
	return response.(*GetAllPermissionResponse).P, res1
}

func (set EndpointsSet) GetFilteredPermissions(arg0 context.Context, arg1 *types.Permission) (res0 []*types.Permission, res1 error) {
	request := GetFilteredPermissionsRequest{P: arg1}
	response, res1 := set.GetFilteredPermissionsEndpoint(arg0, &request)
	if res1 != nil {
		if e, ok := status.FromError(res1); ok || e.Code() == codes.Internal || e.Code() == codes.Unknown {
			res1 = errors.New(e.Message())
		}
		return
	}
	return response.(*GetFilteredPermissionsResponse).Perm, res1
}

func (set EndpointsSet) DeletePermission(arg0 context.Context, arg1 *types.Permission) (res0 bool, res1 error) {
	request := DeletePermissionRequest{P: arg1}
	response, res1 := set.DeletePermissionEndpoint(arg0, &request)
	if res1 != nil {
		if e, ok := status.FromError(res1); ok || e.Code() == codes.Internal || e.Code() == codes.Unknown {
			res1 = errors.New(e.Message())
		}
		return
	}
	return response.(*DeletePermissionResponse).Ok, res1
}

func (set EndpointsSet) GetUserPermissions(arg0 context.Context, arg1 uint32) (res0 []*types.Permission, res1 error) {
	request := GetUserPermissionsRequest{UserId: arg1}
	response, res1 := set.GetUserPermissionsEndpoint(arg0, &request)
	if res1 != nil {
		if e, ok := status.FromError(res1); ok || e.Code() == codes.Internal || e.Code() == codes.Unknown {
			res1 = errors.New(e.Message())
		}
		return
	}
	return response.(*GetUserPermissionsResponse).Permissions, res1
}

func (set EndpointsSet) AddUserPermission(arg0 context.Context, arg1 *types.Permission, arg2 uint32) (res0 bool, res1 error) {
	request := AddUserPermissionRequest{
		P:      arg1,
		UserId: arg2,
	}
	response, res1 := set.AddUserPermissionEndpoint(arg0, &request)
	if res1 != nil {
		if e, ok := status.FromError(res1); ok || e.Code() == codes.Internal || e.Code() == codes.Unknown {
			res1 = errors.New(e.Message())
		}
		return
	}
	return response.(*AddUserPermissionResponse).Ok, res1
}

func (set EndpointsSet) RemoveUserPermission(arg0 context.Context, arg1 uint32, arg2 uint32) (res0 bool, res1 error) {
	request := RemoveUserPermissionRequest{
		PermId: arg1,
		UserId: arg2,
	}
	response, res1 := set.RemoveUserPermissionEndpoint(arg0, &request)
	if res1 != nil {
		if e, ok := status.FromError(res1); ok || e.Code() == codes.Internal || e.Code() == codes.Unknown {
			res1 = errors.New(e.Message())
		}
		return
	}
	return response.(*RemoveUserPermissionResponse).Ok, res1
}
