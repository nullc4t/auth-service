// Code generated by microgen 1.0.5. DO NOT EDIT.

package service

import (
	access "auth/pkg/access"
	service "auth/pkg/mgmt"
	types "auth/pkg/types"
	"context"
	log "github.com/go-kit/kit/log"
	"time"
)

// LoggingMiddleware writes params, results and working time of method call to provided logger after its execution.
func LoggingMiddleware(logger log.Logger) Middleware {
	return func(next service.Service) service.Service {
		return &loggingMiddleware{
			logger: logger,
			next:   next,
		}
	}
}

type loggingMiddleware struct {
	logger log.Logger
	next   service.Service
}

func (M loggingMiddleware) CreateUserWithLoginPassword(arg0 context.Context, arg1 string, arg2 string) (res0 *types.User, res1 error) {
	defer func(begin time.Time) {
		M.logger.Log(
			"method", "CreateUserWithLoginPassword",
			"message", "CreateUserWithLoginPassword called",
			"request", logCreateUserWithLoginPasswordRequest{
				Login: arg1,
				Pass:  arg2,
			},
			"response", logCreateUserWithLoginPasswordResponse{User: res0},
			"err", res1,
			"took", time.Since(begin))
	}(time.Now())
	return M.next.CreateUserWithLoginPassword(arg0, arg1, arg2)
}

func (M loggingMiddleware) CreateUserWithTelegram(arg0 context.Context, arg1 uint64, arg2 string, arg3 string) (res0 *types.User, res1 error) {
	defer func(begin time.Time) {
		M.logger.Log(
			"method", "CreateUserWithTelegram",
			"message", "CreateUserWithTelegram called",
			"request", logCreateUserWithTelegramRequest{
				Id:    arg1,
				Name:  arg2,
				UserN: arg3,
			},
			"response", logCreateUserWithTelegramResponse{User: res0},
			"err", res1,
			"took", time.Since(begin))
	}(time.Now())
	return M.next.CreateUserWithTelegram(arg0, arg1, arg2, arg3)
}

func (M loggingMiddleware) GetAllUsers(arg0 context.Context) (res0 []*types.User, res1 error) {
	defer func(begin time.Time) {
		M.logger.Log(
			"method", "GetAllUsers",
			"message", "GetAllUsers called",
			"response", logGetAllUsersResponse{Users: res0},
			"err", res1,
			"took", time.Since(begin))
	}(time.Now())
	return M.next.GetAllUsers(arg0)
}

func (M loggingMiddleware) GetUser(arg0 context.Context, arg1 *types.User) (res0 *types.User, res1 error) {
	defer func(begin time.Time) {
		M.logger.Log(
			"method", "GetUser",
			"message", "GetUser called",
			"request", logGetUserRequest{UserReq: arg1},
			"response", logGetUserResponse{User: res0},
			"err", res1,
			"took", time.Since(begin))
	}(time.Now())
	return M.next.GetUser(arg0, arg1)
}

func (M loggingMiddleware) UpdateUser(arg0 context.Context, arg1 *types.User) (res0 *types.User, res1 error) {
	defer func(begin time.Time) {
		M.logger.Log(
			"method", "UpdateUser",
			"message", "UpdateUser called",
			"request", logUpdateUserRequest{UserReq: arg1},
			"response", logUpdateUserResponse{User: res0},
			"err", res1,
			"took", time.Since(begin))
	}(time.Now())
	return M.next.UpdateUser(arg0, arg1)
}

func (M loggingMiddleware) BlockUser(arg0 context.Context, arg1 uint32) (res0 bool, res1 error) {
	defer func(begin time.Time) {
		M.logger.Log(
			"method", "BlockUser",
			"message", "BlockUser called",
			"request", logBlockUserRequest{UserId: arg1},
			"response", logBlockUserResponse{Ok: res0},
			"err", res1,
			"took", time.Since(begin))
	}(time.Now())
	return M.next.BlockUser(arg0, arg1)
}

func (M loggingMiddleware) UnblockUser(arg0 context.Context, arg1 uint32) (res0 bool, res1 error) {
	defer func(begin time.Time) {
		M.logger.Log(
			"method", "UnblockUser",
			"message", "UnblockUser called",
			"request", logUnblockUserRequest{UserId: arg1},
			"response", logUnblockUserResponse{Ok: res0},
			"err", res1,
			"took", time.Since(begin))
	}(time.Now())
	return M.next.UnblockUser(arg0, arg1)
}

func (M loggingMiddleware) CreateService(arg0 context.Context, arg1 string) (res0 *types.Service, res1 error) {
	defer func(begin time.Time) {
		M.logger.Log(
			"method", "CreateService",
			"message", "CreateService called",
			"request", logCreateServiceRequest{Name: arg1},
			"response", logCreateServiceResponse{S: res0},
			"err", res1,
			"took", time.Since(begin))
	}(time.Now())
	return M.next.CreateService(arg0, arg1)
}

func (M loggingMiddleware) GetAllServices(arg0 context.Context) (res0 []*types.Service, res1 error) {
	defer func(begin time.Time) {
		M.logger.Log(
			"method", "GetAllServices",
			"message", "GetAllServices called",
			"response", logGetAllServicesResponse{Ss: res0},
			"err", res1,
			"took", time.Since(begin))
	}(time.Now())
	return M.next.GetAllServices(arg0)
}

func (M loggingMiddleware) GetService(arg0 context.Context, arg1 *types.Service) (res0 *types.Service, res1 error) {
	defer func(begin time.Time) {
		M.logger.Log(
			"method", "GetService",
			"message", "GetService called",
			"request", logGetServiceRequest{Svc: arg1},
			"response", logGetServiceResponse{S: res0},
			"err", res1,
			"took", time.Since(begin))
	}(time.Now())
	return M.next.GetService(arg0, arg1)
}

func (M loggingMiddleware) CreateAccount(arg0 context.Context) (res0 *types.Account, res1 error) {
	defer func(begin time.Time) {
		M.logger.Log(
			"method", "CreateAccount",
			"message", "CreateAccount called",
			"response", logCreateAccountResponse{A: res0},
			"err", res1,
			"took", time.Since(begin))
	}(time.Now())
	return M.next.CreateAccount(arg0)
}

func (M loggingMiddleware) CreateAccountWithName(arg0 context.Context, arg1 string) (res0 *types.Account, res1 error) {
	defer func(begin time.Time) {
		M.logger.Log(
			"method", "CreateAccountWithName",
			"message", "CreateAccountWithName called",
			"request", logCreateAccountWithNameRequest{Name: arg1},
			"response", logCreateAccountWithNameResponse{A: res0},
			"err", res1,
			"took", time.Since(begin))
	}(time.Now())
	return M.next.CreateAccountWithName(arg0, arg1)
}

func (M loggingMiddleware) GetAllAccounts(arg0 context.Context) (res0 []*types.Account, res1 error) {
	defer func(begin time.Time) {
		M.logger.Log(
			"method", "GetAllAccounts",
			"message", "GetAllAccounts called",
			"response", logGetAllAccountsResponse{As: res0},
			"err", res1,
			"took", time.Since(begin))
	}(time.Now())
	return M.next.GetAllAccounts(arg0)
}

func (M loggingMiddleware) GetAccount(arg0 context.Context, arg1 *types.Account) (res0 *types.Account, res1 error) {
	defer func(begin time.Time) {
		M.logger.Log(
			"method", "GetAccount",
			"message", "GetAccount called",
			"request", logGetAccountRequest{Acc: arg1},
			"response", logGetAccountResponse{A: res0},
			"err", res1,
			"took", time.Since(begin))
	}(time.Now())
	return M.next.GetAccount(arg0, arg1)
}

func (M loggingMiddleware) UpdateAccount(arg0 context.Context, arg1 *types.Account) (res0 *types.Account, res1 error) {
	defer func(begin time.Time) {
		M.logger.Log(
			"method", "UpdateAccount",
			"message", "UpdateAccount called",
			"request", logUpdateAccountRequest{Acc: arg1},
			"response", logUpdateAccountResponse{A: res0},
			"err", res1,
			"took", time.Since(begin))
	}(time.Now())
	return M.next.UpdateAccount(arg0, arg1)
}

func (M loggingMiddleware) AttachUserToAccount(arg0 context.Context, arg1 uint32, arg2 uint32) (res0 bool, res1 error) {
	defer func(begin time.Time) {
		M.logger.Log(
			"method", "AttachUserToAccount",
			"message", "AttachUserToAccount called",
			"request", logAttachUserToAccountRequest{
				AccountId: arg2,
				UserId:    arg1,
			},
			"response", logAttachUserToAccountResponse{Ok: res0},
			"err", res1,
			"took", time.Since(begin))
	}(time.Now())
	return M.next.AttachUserToAccount(arg0, arg1, arg2)
}

func (M loggingMiddleware) AttachAccountToService(arg0 context.Context, arg1 uint32, arg2 uint32) (res0 bool, res1 error) {
	defer func(begin time.Time) {
		M.logger.Log(
			"method", "AttachAccountToService",
			"message", "AttachAccountToService called",
			"request", logAttachAccountToServiceRequest{
				AccountId: arg2,
				ServiceId: arg1,
			},
			"response", logAttachAccountToServiceResponse{Ok: res0},
			"err", res1,
			"took", time.Since(begin))
	}(time.Now())
	return M.next.AttachAccountToService(arg0, arg1, arg2)
}

func (M loggingMiddleware) RemoveAccountFromService(arg0 context.Context, arg1 uint32, arg2 uint32) (res0 bool, res1 error) {
	defer func(begin time.Time) {
		M.logger.Log(
			"method", "RemoveAccountFromService",
			"message", "RemoveAccountFromService called",
			"request", logRemoveAccountFromServiceRequest{
				AccountId: arg2,
				ServiceId: arg1,
			},
			"response", logRemoveAccountFromServiceResponse{Ok: res0},
			"err", res1,
			"took", time.Since(begin))
	}(time.Now())
	return M.next.RemoveAccountFromService(arg0, arg1, arg2)
}

func (M loggingMiddleware) CreatePermission(arg0 context.Context, arg1 uint32, arg2 string, arg3 *access.Access) (res0 *types.Permission, res1 error) {
	defer func(begin time.Time) {
		M.logger.Log(
			"method", "CreatePermission",
			"message", "CreatePermission called",
			"request", logCreatePermissionRequest{
				Access:    arg3,
				Name:      arg2,
				ServiceId: arg1,
			},
			"response", logCreatePermissionResponse{P: res0},
			"err", res1,
			"took", time.Since(begin))
	}(time.Now())
	return M.next.CreatePermission(arg0, arg1, arg2, arg3)
}

func (M loggingMiddleware) GetPermission(arg0 context.Context, arg1 *types.Permission) (res0 *types.Permission, res1 error) {
	defer func(begin time.Time) {
		M.logger.Log(
			"method", "GetPermission",
			"message", "GetPermission called",
			"request", logGetPermissionRequest{P: arg1},
			"response", logGetPermissionResponse{Perm: res0},
			"err", res1,
			"took", time.Since(begin))
	}(time.Now())
	return M.next.GetPermission(arg0, arg1)
}

func (M loggingMiddleware) GetAllPermission(arg0 context.Context) (res0 []*types.Permission, res1 error) {
	defer func(begin time.Time) {
		M.logger.Log(
			"method", "GetAllPermission",
			"message", "GetAllPermission called",
			"response", logGetAllPermissionResponse{P: res0},
			"err", res1,
			"took", time.Since(begin))
	}(time.Now())
	return M.next.GetAllPermission(arg0)
}

func (M loggingMiddleware) GetFilteredPermissions(arg0 context.Context, arg1 *types.Permission) (res0 []*types.Permission, res1 error) {
	defer func(begin time.Time) {
		M.logger.Log(
			"method", "GetFilteredPermissions",
			"message", "GetFilteredPermissions called",
			"request", logGetFilteredPermissionsRequest{P: arg1},
			"response", logGetFilteredPermissionsResponse{Perm: res0},
			"err", res1,
			"took", time.Since(begin))
	}(time.Now())
	return M.next.GetFilteredPermissions(arg0, arg1)
}

func (M loggingMiddleware) DeletePermission(arg0 context.Context, arg1 *types.Permission) (res0 bool, res1 error) {
	defer func(begin time.Time) {
		M.logger.Log(
			"method", "DeletePermission",
			"message", "DeletePermission called",
			"request", logDeletePermissionRequest{P: arg1},
			"response", logDeletePermissionResponse{Ok: res0},
			"err", res1,
			"took", time.Since(begin))
	}(time.Now())
	return M.next.DeletePermission(arg0, arg1)
}

func (M loggingMiddleware) GetUserPermissions(arg0 context.Context, arg1 uint32) (res0 []*types.Permission, res1 error) {
	defer func(begin time.Time) {
		M.logger.Log(
			"method", "GetUserPermissions",
			"message", "GetUserPermissions called",
			"request", logGetUserPermissionsRequest{UserId: arg1},
			"response", logGetUserPermissionsResponse{Permissions: res0},
			"err", res1,
			"took", time.Since(begin))
	}(time.Now())
	return M.next.GetUserPermissions(arg0, arg1)
}

func (M loggingMiddleware) AddUserPermission(arg0 context.Context, arg1 *types.Permission, arg2 uint32) (res0 bool, res1 error) {
	defer func(begin time.Time) {
		M.logger.Log(
			"method", "AddUserPermission",
			"message", "AddUserPermission called",
			"request", logAddUserPermissionRequest{
				P:      arg1,
				UserId: arg2,
			},
			"response", logAddUserPermissionResponse{Ok: res0},
			"err", res1,
			"took", time.Since(begin))
	}(time.Now())
	return M.next.AddUserPermission(arg0, arg1, arg2)
}

func (M loggingMiddleware) RemoveUserPermission(arg0 context.Context, arg1 uint32, arg2 uint32) (res0 bool, res1 error) {
	defer func(begin time.Time) {
		M.logger.Log(
			"method", "RemoveUserPermission",
			"message", "RemoveUserPermission called",
			"request", logRemoveUserPermissionRequest{
				PermId: arg1,
				UserId: arg2,
			},
			"response", logRemoveUserPermissionResponse{Ok: res0},
			"err", res1,
			"took", time.Since(begin))
	}(time.Now())
	return M.next.RemoveUserPermission(arg0, arg1, arg2)
}

type (
	logCreateUserWithLoginPasswordRequest struct {
		Login string
		Pass  string
	}
	logCreateUserWithLoginPasswordResponse struct {
		User *types.User
	}
	logCreateUserWithTelegramRequest struct {
		Id    uint64
		Name  string
		UserN string
	}
	logCreateUserWithTelegramResponse struct {
		User *types.User
	}
	logGetAllUsersResponse struct {
		Users []*types.User
	}
	logGetUserRequest struct {
		UserReq *types.User
	}
	logGetUserResponse struct {
		User *types.User
	}
	logUpdateUserRequest struct {
		UserReq *types.User
	}
	logUpdateUserResponse struct {
		User *types.User
	}
	logBlockUserRequest struct {
		UserId uint32
	}
	logBlockUserResponse struct {
		Ok bool
	}
	logUnblockUserRequest struct {
		UserId uint32
	}
	logUnblockUserResponse struct {
		Ok bool
	}
	logCreateServiceRequest struct {
		Name string
	}
	logCreateServiceResponse struct {
		S *types.Service
	}
	logGetAllServicesResponse struct {
		Ss []*types.Service
	}
	logGetServiceRequest struct {
		Svc *types.Service
	}
	logGetServiceResponse struct {
		S *types.Service
	}
	logCreateAccountResponse struct {
		A *types.Account
	}
	logCreateAccountWithNameRequest struct {
		Name string
	}
	logCreateAccountWithNameResponse struct {
		A *types.Account
	}
	logGetAllAccountsResponse struct {
		As []*types.Account
	}
	logGetAccountRequest struct {
		Acc *types.Account
	}
	logGetAccountResponse struct {
		A *types.Account
	}
	logUpdateAccountRequest struct {
		Acc *types.Account
	}
	logUpdateAccountResponse struct {
		A *types.Account
	}
	logAttachUserToAccountRequest struct {
		UserId    uint32
		AccountId uint32
	}
	logAttachUserToAccountResponse struct {
		Ok bool
	}
	logAttachAccountToServiceRequest struct {
		ServiceId uint32
		AccountId uint32
	}
	logAttachAccountToServiceResponse struct {
		Ok bool
	}
	logRemoveAccountFromServiceRequest struct {
		ServiceId uint32
		AccountId uint32
	}
	logRemoveAccountFromServiceResponse struct {
		Ok bool
	}
	logCreatePermissionRequest struct {
		ServiceId uint32
		Name      string
		Access    *access.Access
	}
	logCreatePermissionResponse struct {
		P *types.Permission
	}
	logGetPermissionRequest struct {
		P *types.Permission
	}
	logGetPermissionResponse struct {
		Perm *types.Permission
	}
	logGetAllPermissionResponse struct {
		P []*types.Permission
	}
	logGetFilteredPermissionsRequest struct {
		P *types.Permission
	}
	logGetFilteredPermissionsResponse struct {
		Perm []*types.Permission
	}
	logDeletePermissionRequest struct {
		P *types.Permission
	}
	logDeletePermissionResponse struct {
		Ok bool
	}
	logGetUserPermissionsRequest struct {
		UserId uint32
	}
	logGetUserPermissionsResponse struct {
		Permissions []*types.Permission
	}
	logAddUserPermissionRequest struct {
		P      *types.Permission
		UserId uint32
	}
	logAddUserPermissionResponse struct {
		Ok bool
	}
	logRemoveUserPermissionRequest struct {
		PermId uint32
		UserId uint32
	}
	logRemoveUserPermissionResponse struct {
		Ok bool
	}
)
