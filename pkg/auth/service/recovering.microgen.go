// Code generated by microgen 1.0.5. DO NOT EDIT.

package service

import (
	"context"
	"fmt"
	log "github.com/go-kit/kit/log"
	service "github.com/nullc4t/auth-service/pkg/auth"
	types "github.com/nullc4t/auth-service/pkg/types"
)

// RecoveringMiddleware recovers panics from method calls, writes to provided logger and returns the error of panic as method error.
func RecoveringMiddleware(logger log.Logger) Middleware {
	return func(next service.Service) service.Service {
		return &recoveringMiddleware{
			logger: logger,
			next:   next,
		}
	}
}

type recoveringMiddleware struct {
	logger log.Logger
	next   service.Service
}

func (M recoveringMiddleware) Register(ctx context.Context, login string, password string, service string, accountId uint32) (ok bool, err error) {
	defer func() {
		if r := recover(); r != nil {
			M.logger.Log("method", "Register", "message", r)
			err = fmt.Errorf("%v", r)
		}
	}()
	return M.next.Register(ctx, login, password, service, accountId)
}

func (M recoveringMiddleware) Login(ctx context.Context, login string, password string, service string) (token *types.AccessToken, err error) {
	defer func() {
		if r := recover(); r != nil {
			M.logger.Log("method", "Login", "message", r)
			err = fmt.Errorf("%v", r)
		}
	}()
	return M.next.Login(ctx, login, password, service)
}

func (M recoveringMiddleware) PublicKey(ctx context.Context) (pub []byte, err error) {
	defer func() {
		if r := recover(); r != nil {
			M.logger.Log("method", "PublicKey", "message", r)
			err = fmt.Errorf("%v", r)
		}
	}()
	return M.next.PublicKey(ctx)
}

func (M recoveringMiddleware) GetPermissionsForService(ctx context.Context, name string) (permissions []*types.Permission, err error) {
	defer func() {
		if r := recover(); r != nil {
			M.logger.Log("method", "GetPermissionsForService", "message", r)
			err = fmt.Errorf("%v", r)
		}
	}()
	return M.next.GetPermissionsForService(ctx, name)
}
