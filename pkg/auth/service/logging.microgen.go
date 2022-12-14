// Code generated by microgen 1.0.5. DO NOT EDIT.

package service

import (
	"context"
	log "github.com/go-kit/kit/log"
	service "github.com/nullc4t/auth-service/pkg/auth"
	types "github.com/nullc4t/auth-service/pkg/types"
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

func (M loggingMiddleware) Register(arg0 context.Context, arg1 string, arg2 string, arg3 string, arg4 uint32) (res0 bool, res1 error) {
	defer func(begin time.Time) {
		M.logger.Log(
			"method", "Register",
			"message", "Register called",
			"request", logRegisterRequest{
				AccountId: arg4,
				Login:     arg1,
				Password:  arg2,
				Service:   arg3,
			},
			"response", logRegisterResponse{Ok: res0},
			"err", res1,
			"took", time.Since(begin))
	}(time.Now())
	return M.next.Register(arg0, arg1, arg2, arg3, arg4)
}

func (M loggingMiddleware) Login(arg0 context.Context, arg1 string, arg2 string, arg3 string) (res0 *types.AccessToken, res1 error) {
	defer func(begin time.Time) {
		M.logger.Log(
			"method", "Login",
			"message", "Login called",
			"request", logLoginRequest{
				Login:    arg1,
				Password: arg2,
				Service:  arg3,
			},
			"response", logLoginResponse{Token: res0},
			"err", res1,
			"took", time.Since(begin))
	}(time.Now())
	return M.next.Login(arg0, arg1, arg2, arg3)
}

func (M loggingMiddleware) PublicKey(arg0 context.Context) (res0 []byte, res1 error) {
	defer func(begin time.Time) {
		M.logger.Log(
			"method", "PublicKey",
			"message", "PublicKey called",
			"response", logPublicKeyResponse{Pub: res0},
			"err", res1,
			"took", time.Since(begin))
	}(time.Now())
	return M.next.PublicKey(arg0)
}

func (M loggingMiddleware) GetPermissionsForService(arg0 context.Context, arg1 string) (res0 []*types.Permission, res1 error) {
	defer func(begin time.Time) {
		M.logger.Log(
			"method", "GetPermissionsForService",
			"message", "GetPermissionsForService called",
			"request", logGetPermissionsForServiceRequest{Name: arg1},
			"response", logGetPermissionsForServiceResponse{Permissions: res0},
			"err", res1,
			"took", time.Since(begin))
	}(time.Now())
	return M.next.GetPermissionsForService(arg0, arg1)
}

type (
	logRegisterRequest struct {
		Login     string
		Password  string
		Service   string
		AccountId uint32
	}
	logRegisterResponse struct {
		Ok bool
	}
	logLoginRequest struct {
		Login    string
		Password string
		Service  string
	}
	logLoginResponse struct {
		Token *types.AccessToken
	}
	logPublicKeyResponse struct {
		Pub []byte
	}
	logGetPermissionsForServiceRequest struct {
		Name string
	}
	logGetPermissionsForServiceResponse struct {
		Permissions []*types.Permission
	}
)
