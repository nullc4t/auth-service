// Code generated by microgen 1.0.5. DO NOT EDIT.

package service

import (
	"context"
	log "github.com/go-kit/kit/log"
	access "github.com/nullc4t/auth-service/pkg/access"
	service "github.com/nullc4t/auth-service/pkg/mgmt"
	types "github.com/nullc4t/auth-service/pkg/types"
)

// ErrorLoggingMiddleware writes to logger any error, if it is not nil.
func ErrorLoggingMiddleware(logger log.Logger) Middleware {
	return func(next service.Service) service.Service {
		return &errorLoggingMiddleware{
			logger: logger,
			next:   next,
		}
	}
}

type errorLoggingMiddleware struct {
	logger log.Logger
	next   service.Service
}

func (M errorLoggingMiddleware) CreateUserWithLoginPassword(ctx context.Context, login string, pass string) (user *types.User, err error) {
	defer func() {
		if err != nil {
			M.logger.Log("method", "CreateUserWithLoginPassword", "message", err)
		}
	}()
	return M.next.CreateUserWithLoginPassword(ctx, login, pass)
}

func (M errorLoggingMiddleware) CreateUserWithTelegram(ctx context.Context, id uint64, name string, userN string) (user *types.User, err error) {
	defer func() {
		if err != nil {
			M.logger.Log("method", "CreateUserWithTelegram", "message", err)
		}
	}()
	return M.next.CreateUserWithTelegram(ctx, id, name, userN)
}

func (M errorLoggingMiddleware) GetAllUsers(ctx context.Context) (users []*types.User, err error) {
	defer func() {
		if err != nil {
			M.logger.Log("method", "GetAllUsers", "message", err)
		}
	}()
	return M.next.GetAllUsers(ctx)
}

func (M errorLoggingMiddleware) GetUser(ctx context.Context, userReq *types.User) (user *types.User, err error) {
	defer func() {
		if err != nil {
			M.logger.Log("method", "GetUser", "message", err)
		}
	}()
	return M.next.GetUser(ctx, userReq)
}

func (M errorLoggingMiddleware) UpdateUser(ctx context.Context, userReq *types.User) (user *types.User, err error) {
	defer func() {
		if err != nil {
			M.logger.Log("method", "UpdateUser", "message", err)
		}
	}()
	return M.next.UpdateUser(ctx, userReq)
}

func (M errorLoggingMiddleware) BlockUser(ctx context.Context, userId uint32) (ok bool, err error) {
	defer func() {
		if err != nil {
			M.logger.Log("method", "BlockUser", "message", err)
		}
	}()
	return M.next.BlockUser(ctx, userId)
}

func (M errorLoggingMiddleware) UnblockUser(ctx context.Context, userId uint32) (ok bool, err error) {
	defer func() {
		if err != nil {
			M.logger.Log("method", "UnblockUser", "message", err)
		}
	}()
	return M.next.UnblockUser(ctx, userId)
}

func (M errorLoggingMiddleware) CreateService(ctx context.Context, name string) (s *types.Service, err error) {
	defer func() {
		if err != nil {
			M.logger.Log("method", "CreateService", "message", err)
		}
	}()
	return M.next.CreateService(ctx, name)
}

func (M errorLoggingMiddleware) GetAllServices(ctx context.Context) (ss []*types.Service, err error) {
	defer func() {
		if err != nil {
			M.logger.Log("method", "GetAllServices", "message", err)
		}
	}()
	return M.next.GetAllServices(ctx)
}

func (M errorLoggingMiddleware) GetService(ctx context.Context, svc *types.Service) (s *types.Service, err error) {
	defer func() {
		if err != nil {
			M.logger.Log("method", "GetService", "message", err)
		}
	}()
	return M.next.GetService(ctx, svc)
}

func (M errorLoggingMiddleware) CreateAccount(ctx context.Context) (a *types.Account, err error) {
	defer func() {
		if err != nil {
			M.logger.Log("method", "CreateAccount", "message", err)
		}
	}()
	return M.next.CreateAccount(ctx)
}

func (M errorLoggingMiddleware) CreateAccountWithName(ctx context.Context, name string) (a *types.Account, err error) {
	defer func() {
		if err != nil {
			M.logger.Log("method", "CreateAccountWithName", "message", err)
		}
	}()
	return M.next.CreateAccountWithName(ctx, name)
}

func (M errorLoggingMiddleware) GetAllAccounts(ctx context.Context) (as []*types.Account, err error) {
	defer func() {
		if err != nil {
			M.logger.Log("method", "GetAllAccounts", "message", err)
		}
	}()
	return M.next.GetAllAccounts(ctx)
}

func (M errorLoggingMiddleware) GetAccount(ctx context.Context, acc *types.Account) (a *types.Account, err error) {
	defer func() {
		if err != nil {
			M.logger.Log("method", "GetAccount", "message", err)
		}
	}()
	return M.next.GetAccount(ctx, acc)
}

func (M errorLoggingMiddleware) UpdateAccount(ctx context.Context, acc *types.Account) (a *types.Account, err error) {
	defer func() {
		if err != nil {
			M.logger.Log("method", "UpdateAccount", "message", err)
		}
	}()
	return M.next.UpdateAccount(ctx, acc)
}

func (M errorLoggingMiddleware) AttachUserToAccount(ctx context.Context, userId uint32, accountId uint32) (ok bool, err error) {
	defer func() {
		if err != nil {
			M.logger.Log("method", "AttachUserToAccount", "message", err)
		}
	}()
	return M.next.AttachUserToAccount(ctx, userId, accountId)
}

func (M errorLoggingMiddleware) AttachAccountToService(ctx context.Context, serviceId uint32, accountId uint32) (ok bool, err error) {
	defer func() {
		if err != nil {
			M.logger.Log("method", "AttachAccountToService", "message", err)
		}
	}()
	return M.next.AttachAccountToService(ctx, serviceId, accountId)
}

func (M errorLoggingMiddleware) RemoveAccountFromService(ctx context.Context, serviceId uint32, accountId uint32) (ok bool, err error) {
	defer func() {
		if err != nil {
			M.logger.Log("method", "RemoveAccountFromService", "message", err)
		}
	}()
	return M.next.RemoveAccountFromService(ctx, serviceId, accountId)
}

func (M errorLoggingMiddleware) CreatePermission(ctx context.Context, serviceId uint32, name string, access *access.Access) (p *types.Permission, err error) {
	defer func() {
		if err != nil {
			M.logger.Log("method", "CreatePermission", "message", err)
		}
	}()
	return M.next.CreatePermission(ctx, serviceId, name, access)
}

func (M errorLoggingMiddleware) GetPermission(ctx context.Context, p *types.Permission) (perm *types.Permission, err error) {
	defer func() {
		if err != nil {
			M.logger.Log("method", "GetPermission", "message", err)
		}
	}()
	return M.next.GetPermission(ctx, p)
}

func (M errorLoggingMiddleware) GetAllPermission(ctx context.Context) (p []*types.Permission, err error) {
	defer func() {
		if err != nil {
			M.logger.Log("method", "GetAllPermission", "message", err)
		}
	}()
	return M.next.GetAllPermission(ctx)
}

func (M errorLoggingMiddleware) GetFilteredPermissions(ctx context.Context, p *types.Permission) (perm []*types.Permission, err error) {
	defer func() {
		if err != nil {
			M.logger.Log("method", "GetFilteredPermissions", "message", err)
		}
	}()
	return M.next.GetFilteredPermissions(ctx, p)
}

func (M errorLoggingMiddleware) DeletePermission(ctx context.Context, p *types.Permission) (ok bool, err error) {
	defer func() {
		if err != nil {
			M.logger.Log("method", "DeletePermission", "message", err)
		}
	}()
	return M.next.DeletePermission(ctx, p)
}

func (M errorLoggingMiddleware) GetUserPermissions(ctx context.Context, userId uint32) (permissions []*types.Permission, err error) {
	defer func() {
		if err != nil {
			M.logger.Log("method", "GetUserPermissions", "message", err)
		}
	}()
	return M.next.GetUserPermissions(ctx, userId)
}

func (M errorLoggingMiddleware) AddUserPermission(ctx context.Context, p *types.Permission, userId uint32) (ok bool, err error) {
	defer func() {
		if err != nil {
			M.logger.Log("method", "AddUserPermission", "message", err)
		}
	}()
	return M.next.AddUserPermission(ctx, p, userId)
}

func (M errorLoggingMiddleware) RemoveUserPermission(ctx context.Context, permId uint32, userId uint32) (ok bool, err error) {
	defer func() {
		if err != nil {
			M.logger.Log("method", "RemoveUserPermission", "message", err)
		}
	}()
	return M.next.RemoveUserPermission(ctx, permId, userId)
}
