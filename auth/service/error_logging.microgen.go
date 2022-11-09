// Code generated by microgen 1.0.5. DO NOT EDIT.

package service

import (
	service "auth/auth"
	types "auth/pkg/types"
	"context"
	log "github.com/go-kit/kit/log"
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

func (M errorLoggingMiddleware) Register(ctx context.Context, login string, password string, service string, accountID uint) (ok bool, err error) {
	defer func() {
		if err != nil {
			M.logger.Log("method", "Register", "message", err)
		}
	}()
	return M.next.Register(ctx, login, password, service, accountID)
}

func (M errorLoggingMiddleware) Login(ctx context.Context, login string, password string, service string) (at types.AccessToken, err error) {
	defer func() {
		if err != nil {
			M.logger.Log("method", "Login", "message", err)
		}
	}()
	return M.next.Login(ctx, login, password, service)
}
