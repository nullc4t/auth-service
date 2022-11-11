// Code generated by microgen 1.0.5. DO NOT EDIT.

package transport

import types "auth/pkg/types"

type (
	RegisterRequest struct {
		Login     string `json:"login"`
		Password  string `json:"password"`
		Service   string `json:"service"`
		AccountId uint32 `json:"account_id"`
	}
	RegisterResponse struct {
		Ok bool `json:"ok"`
	}

	LoginRequest struct {
		Login    string `json:"login"`
		Password string `json:"password"`
		Service  string `json:"service"`
	}
	LoginResponse struct {
		Token *types.AccessToken `json:"token"`
	}

	// Formal exchange type, please do not delete.
	PublicKeyRequest  struct{}
	PublicKeyResponse struct {
		Pub []byte `json:"pub"`
	}

	GetPermissionsForServiceRequest struct {
		Name string `json:"name"`
	}
	GetPermissionsForServiceResponse struct {
		Permissions []*types.Permission `json:"permissions"`
	}
)
