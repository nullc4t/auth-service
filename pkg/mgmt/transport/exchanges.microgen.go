// Code generated by microgen 1.0.5. DO NOT EDIT.

package transport

import (
	access "github.com/nullc4t/auth-service/pkg/access"
	types "github.com/nullc4t/auth-service/pkg/types"
)

type (
	CreateUserWithLoginPasswordRequest struct {
		Login string `json:"login"`
		Pass  string `json:"pass"`
	}
	CreateUserWithLoginPasswordResponse struct {
		User *types.User `json:"user"`
	}

	CreateUserWithTelegramRequest struct {
		Id    uint64 `json:"id"`
		Name  string `json:"name"`
		UserN string `json:"user_n"`
	}
	CreateUserWithTelegramResponse struct {
		User *types.User `json:"user"`
	}

	// Formal exchange type, please do not delete.
	GetAllUsersRequest  struct{}
	GetAllUsersResponse struct {
		Users []*types.User `json:"users"`
	}

	GetUserRequest struct {
		UserReq *types.User `json:"user_req"`
	}
	GetUserResponse struct {
		User *types.User `json:"user"`
	}

	UpdateUserRequest struct {
		UserReq *types.User `json:"user_req"`
	}
	UpdateUserResponse struct {
		User *types.User `json:"user"`
	}

	BlockUserRequest struct {
		UserId uint32 `json:"user_id"`
	}
	BlockUserResponse struct {
		Ok bool `json:"ok"`
	}

	UnblockUserRequest struct {
		UserId uint32 `json:"user_id"`
	}
	UnblockUserResponse struct {
		Ok bool `json:"ok"`
	}

	CreateServiceRequest struct {
		Name string `json:"name"`
	}
	CreateServiceResponse struct {
		S *types.Service `json:"s"`
	}

	// Formal exchange type, please do not delete.
	GetAllServicesRequest  struct{}
	GetAllServicesResponse struct {
		Ss []*types.Service `json:"ss"`
	}

	GetServiceRequest struct {
		Svc *types.Service `json:"svc"`
	}
	GetServiceResponse struct {
		S *types.Service `json:"s"`
	}

	// Formal exchange type, please do not delete.
	CreateAccountRequest  struct{}
	CreateAccountResponse struct {
		A *types.Account `json:"a"`
	}

	CreateAccountWithNameRequest struct {
		Name string `json:"name"`
	}
	CreateAccountWithNameResponse struct {
		A *types.Account `json:"a"`
	}

	// Formal exchange type, please do not delete.
	GetAllAccountsRequest  struct{}
	GetAllAccountsResponse struct {
		As []*types.Account `json:"as"`
	}

	GetAccountRequest struct {
		Acc *types.Account `json:"acc"`
	}
	GetAccountResponse struct {
		A *types.Account `json:"a"`
	}

	UpdateAccountRequest struct {
		Acc *types.Account `json:"acc"`
	}
	UpdateAccountResponse struct {
		A *types.Account `json:"a"`
	}

	AttachUserToAccountRequest struct {
		UserId    uint32 `json:"user_id"`
		AccountId uint32 `json:"account_id"`
	}
	AttachUserToAccountResponse struct {
		Ok bool `json:"ok"`
	}

	AttachAccountToServiceRequest struct {
		ServiceId uint32 `json:"service_id"`
		AccountId uint32 `json:"account_id"`
	}
	AttachAccountToServiceResponse struct {
		Ok bool `json:"ok"`
	}

	RemoveAccountFromServiceRequest struct {
		ServiceId uint32 `json:"service_id"`
		AccountId uint32 `json:"account_id"`
	}
	RemoveAccountFromServiceResponse struct {
		Ok bool `json:"ok"`
	}

	CreatePermissionRequest struct {
		ServiceId uint32         `json:"service_id"`
		Name      string         `json:"name"`
		Access    *access.Access `json:"access"`
	}
	CreatePermissionResponse struct {
		P *types.Permission `json:"p"`
	}

	GetPermissionRequest struct {
		P *types.Permission `json:"p"`
	}
	GetPermissionResponse struct {
		Perm *types.Permission `json:"perm"`
	}

	// Formal exchange type, please do not delete.
	GetAllPermissionRequest  struct{}
	GetAllPermissionResponse struct {
		P []*types.Permission `json:"p"`
	}

	GetFilteredPermissionsRequest struct {
		P *types.Permission `json:"p"`
	}
	GetFilteredPermissionsResponse struct {
		Perm []*types.Permission `json:"perm"`
	}

	DeletePermissionRequest struct {
		P *types.Permission `json:"p"`
	}
	DeletePermissionResponse struct {
		Ok bool `json:"ok"`
	}

	GetUserPermissionsRequest struct {
		UserId uint32 `json:"user_id"`
	}
	GetUserPermissionsResponse struct {
		Permissions []*types.Permission `json:"permissions"`
	}

	AddUserPermissionRequest struct {
		P      *types.Permission `json:"p"`
		UserId uint32            `json:"user_id"`
	}
	AddUserPermissionResponse struct {
		Ok bool `json:"ok"`
	}

	RemoveUserPermissionRequest struct {
		PermId uint32 `json:"perm_id"`
		UserId uint32 `json:"user_id"`
	}
	RemoveUserPermissionResponse struct {
		Ok bool `json:"ok"`
	}
)
