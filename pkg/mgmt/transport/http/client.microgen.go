// Code generated by microgen 1.0.5. DO NOT EDIT.

package transporthttp

import (
	httpkit "github.com/go-kit/kit/transport/http"
	transport "github.com/nullc4t/auth-service/pkg/mgmt/transport"
	"net/url"
)

func NewHTTPClient(u *url.URL, opts ...httpkit.ClientOption) transport.EndpointsSet {
	return transport.EndpointsSet{
		AddUserPermissionEndpoint: httpkit.NewClient(
			"POST", u,
			_Encode_AddUserPermission_Request,
			_Decode_AddUserPermission_Response,
			opts...,
		).Endpoint(),
		AttachAccountToServiceEndpoint: httpkit.NewClient(
			"POST", u,
			_Encode_AttachAccountToService_Request,
			_Decode_AttachAccountToService_Response,
			opts...,
		).Endpoint(),
		AttachUserToAccountEndpoint: httpkit.NewClient(
			"POST", u,
			_Encode_AttachUserToAccount_Request,
			_Decode_AttachUserToAccount_Response,
			opts...,
		).Endpoint(),
		BlockUserEndpoint: httpkit.NewClient(
			"POST", u,
			_Encode_BlockUser_Request,
			_Decode_BlockUser_Response,
			opts...,
		).Endpoint(),
		CreateAccountEndpoint: httpkit.NewClient(
			"POST", u,
			_Encode_CreateAccount_Request,
			_Decode_CreateAccount_Response,
			opts...,
		).Endpoint(),
		CreateAccountWithNameEndpoint: httpkit.NewClient(
			"POST", u,
			_Encode_CreateAccountWithName_Request,
			_Decode_CreateAccountWithName_Response,
			opts...,
		).Endpoint(),
		CreatePermissionEndpoint: httpkit.NewClient(
			"POST", u,
			_Encode_CreatePermission_Request,
			_Decode_CreatePermission_Response,
			opts...,
		).Endpoint(),
		CreateServiceEndpoint: httpkit.NewClient(
			"POST", u,
			_Encode_CreateService_Request,
			_Decode_CreateService_Response,
			opts...,
		).Endpoint(),
		CreateUserWithLoginPasswordEndpoint: httpkit.NewClient(
			"POST", u,
			_Encode_CreateUserWithLoginPassword_Request,
			_Decode_CreateUserWithLoginPassword_Response,
			opts...,
		).Endpoint(),
		CreateUserWithTelegramEndpoint: httpkit.NewClient(
			"POST", u,
			_Encode_CreateUserWithTelegram_Request,
			_Decode_CreateUserWithTelegram_Response,
			opts...,
		).Endpoint(),
		DeletePermissionEndpoint: httpkit.NewClient(
			"POST", u,
			_Encode_DeletePermission_Request,
			_Decode_DeletePermission_Response,
			opts...,
		).Endpoint(),
		GetAccountEndpoint: httpkit.NewClient(
			"POST", u,
			_Encode_GetAccount_Request,
			_Decode_GetAccount_Response,
			opts...,
		).Endpoint(),
		GetAllAccountsEndpoint: httpkit.NewClient(
			"POST", u,
			_Encode_GetAllAccounts_Request,
			_Decode_GetAllAccounts_Response,
			opts...,
		).Endpoint(),
		GetAllPermissionEndpoint: httpkit.NewClient(
			"POST", u,
			_Encode_GetAllPermission_Request,
			_Decode_GetAllPermission_Response,
			opts...,
		).Endpoint(),
		GetAllServicesEndpoint: httpkit.NewClient(
			"POST", u,
			_Encode_GetAllServices_Request,
			_Decode_GetAllServices_Response,
			opts...,
		).Endpoint(),
		GetAllUsersEndpoint: httpkit.NewClient(
			"POST", u,
			_Encode_GetAllUsers_Request,
			_Decode_GetAllUsers_Response,
			opts...,
		).Endpoint(),
		GetFilteredPermissionsEndpoint: httpkit.NewClient(
			"POST", u,
			_Encode_GetFilteredPermissions_Request,
			_Decode_GetFilteredPermissions_Response,
			opts...,
		).Endpoint(),
		GetPermissionEndpoint: httpkit.NewClient(
			"POST", u,
			_Encode_GetPermission_Request,
			_Decode_GetPermission_Response,
			opts...,
		).Endpoint(),
		GetServiceEndpoint: httpkit.NewClient(
			"POST", u,
			_Encode_GetService_Request,
			_Decode_GetService_Response,
			opts...,
		).Endpoint(),
		GetUserEndpoint: httpkit.NewClient(
			"POST", u,
			_Encode_GetUser_Request,
			_Decode_GetUser_Response,
			opts...,
		).Endpoint(),
		GetUserPermissionsEndpoint: httpkit.NewClient(
			"POST", u,
			_Encode_GetUserPermissions_Request,
			_Decode_GetUserPermissions_Response,
			opts...,
		).Endpoint(),
		RemoveAccountFromServiceEndpoint: httpkit.NewClient(
			"POST", u,
			_Encode_RemoveAccountFromService_Request,
			_Decode_RemoveAccountFromService_Response,
			opts...,
		).Endpoint(),
		RemoveUserPermissionEndpoint: httpkit.NewClient(
			"POST", u,
			_Encode_RemoveUserPermission_Request,
			_Decode_RemoveUserPermission_Response,
			opts...,
		).Endpoint(),
		UnblockUserEndpoint: httpkit.NewClient(
			"POST", u,
			_Encode_UnblockUser_Request,
			_Decode_UnblockUser_Response,
			opts...,
		).Endpoint(),
		UpdateAccountEndpoint: httpkit.NewClient(
			"POST", u,
			_Encode_UpdateAccount_Request,
			_Decode_UpdateAccount_Response,
			opts...,
		).Endpoint(),
		UpdateUserEndpoint: httpkit.NewClient(
			"POST", u,
			_Encode_UpdateUser_Request,
			_Decode_UpdateUser_Response,
			opts...,
		).Endpoint(),
	}
}
