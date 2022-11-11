// Code generated by microgen 1.0.5. DO NOT EDIT.

package transporthttp

import (
	transport "auth/mgmt/transport"
	http "github.com/go-kit/kit/transport/http"
	mux "github.com/gorilla/mux"
	http1 "net/http"
)

func NewHTTPHandler(endpoints *transport.EndpointsSet, opts ...http.ServerOption) http1.Handler {
	mux := mux.NewRouter()
	mux.Methods("POST").Path("/create-user-with-login-password").Handler(
		http.NewServer(
			endpoints.CreateUserWithLoginPasswordEndpoint,
			_Decode_CreateUserWithLoginPassword_Request,
			_Encode_CreateUserWithLoginPassword_Response,
			opts...))
	mux.Methods("POST").Path("/create-user-with-telegram").Handler(
		http.NewServer(
			endpoints.CreateUserWithTelegramEndpoint,
			_Decode_CreateUserWithTelegram_Request,
			_Encode_CreateUserWithTelegram_Response,
			opts...))
	mux.Methods("POST").Path("/get-all-users").Handler(
		http.NewServer(
			endpoints.GetAllUsersEndpoint,
			_Decode_GetAllUsers_Request,
			_Encode_GetAllUsers_Response,
			opts...))
	mux.Methods("POST").Path("/get-user").Handler(
		http.NewServer(
			endpoints.GetUserEndpoint,
			_Decode_GetUser_Request,
			_Encode_GetUser_Response,
			opts...))
	mux.Methods("POST").Path("/update-user").Handler(
		http.NewServer(
			endpoints.UpdateUserEndpoint,
			_Decode_UpdateUser_Request,
			_Encode_UpdateUser_Response,
			opts...))
	mux.Methods("POST").Path("/block-user").Handler(
		http.NewServer(
			endpoints.BlockUserEndpoint,
			_Decode_BlockUser_Request,
			_Encode_BlockUser_Response,
			opts...))
	mux.Methods("POST").Path("/unblock-user").Handler(
		http.NewServer(
			endpoints.UnblockUserEndpoint,
			_Decode_UnblockUser_Request,
			_Encode_UnblockUser_Response,
			opts...))
	mux.Methods("POST").Path("/create-service").Handler(
		http.NewServer(
			endpoints.CreateServiceEndpoint,
			_Decode_CreateService_Request,
			_Encode_CreateService_Response,
			opts...))
	mux.Methods("POST").Path("/get-all-services").Handler(
		http.NewServer(
			endpoints.GetAllServicesEndpoint,
			_Decode_GetAllServices_Request,
			_Encode_GetAllServices_Response,
			opts...))
	mux.Methods("POST").Path("/get-service").Handler(
		http.NewServer(
			endpoints.GetServiceEndpoint,
			_Decode_GetService_Request,
			_Encode_GetService_Response,
			opts...))
	mux.Methods("POST").Path("/create-account").Handler(
		http.NewServer(
			endpoints.CreateAccountEndpoint,
			_Decode_CreateAccount_Request,
			_Encode_CreateAccount_Response,
			opts...))
	mux.Methods("POST").Path("/create-account-with-name").Handler(
		http.NewServer(
			endpoints.CreateAccountWithNameEndpoint,
			_Decode_CreateAccountWithName_Request,
			_Encode_CreateAccountWithName_Response,
			opts...))
	mux.Methods("POST").Path("/get-all-accounts").Handler(
		http.NewServer(
			endpoints.GetAllAccountsEndpoint,
			_Decode_GetAllAccounts_Request,
			_Encode_GetAllAccounts_Response,
			opts...))
	mux.Methods("POST").Path("/get-account").Handler(
		http.NewServer(
			endpoints.GetAccountEndpoint,
			_Decode_GetAccount_Request,
			_Encode_GetAccount_Response,
			opts...))
	mux.Methods("POST").Path("/update-account").Handler(
		http.NewServer(
			endpoints.UpdateAccountEndpoint,
			_Decode_UpdateAccount_Request,
			_Encode_UpdateAccount_Response,
			opts...))
	mux.Methods("POST").Path("/attach-user-toaccount").Handler(
		http.NewServer(
			endpoints.AttachUserToAccountEndpoint,
			_Decode_AttachUserToAccount_Request,
			_Encode_AttachUserToAccount_Response,
			opts...))
	mux.Methods("POST").Path("/attach-account-toservice").Handler(
		http.NewServer(
			endpoints.AttachAccountToServiceEndpoint,
			_Decode_AttachAccountToService_Request,
			_Encode_AttachAccountToService_Response,
			opts...))
	mux.Methods("POST").Path("/remove-account-from-service").Handler(
		http.NewServer(
			endpoints.RemoveAccountFromServiceEndpoint,
			_Decode_RemoveAccountFromService_Request,
			_Encode_RemoveAccountFromService_Response,
			opts...))
	mux.Methods("POST").Path("/create-permission").Handler(
		http.NewServer(
			endpoints.CreatePermissionEndpoint,
			_Decode_CreatePermission_Request,
			_Encode_CreatePermission_Response,
			opts...))
	mux.Methods("POST").Path("/get-permission").Handler(
		http.NewServer(
			endpoints.GetPermissionEndpoint,
			_Decode_GetPermission_Request,
			_Encode_GetPermission_Response,
			opts...))
	mux.Methods("POST").Path("/get-all-permission").Handler(
		http.NewServer(
			endpoints.GetAllPermissionEndpoint,
			_Decode_GetAllPermission_Request,
			_Encode_GetAllPermission_Response,
			opts...))
	mux.Methods("POST").Path("/get-filtered-permissions").Handler(
		http.NewServer(
			endpoints.GetFilteredPermissionsEndpoint,
			_Decode_GetFilteredPermissions_Request,
			_Encode_GetFilteredPermissions_Response,
			opts...))
	mux.Methods("POST").Path("/delete-permission").Handler(
		http.NewServer(
			endpoints.DeletePermissionEndpoint,
			_Decode_DeletePermission_Request,
			_Encode_DeletePermission_Response,
			opts...))
	mux.Methods("POST").Path("/get-user-permissions").Handler(
		http.NewServer(
			endpoints.GetUserPermissionsEndpoint,
			_Decode_GetUserPermissions_Request,
			_Encode_GetUserPermissions_Response,
			opts...))
	mux.Methods("POST").Path("/add-user-permission").Handler(
		http.NewServer(
			endpoints.AddUserPermissionEndpoint,
			_Decode_AddUserPermission_Request,
			_Encode_AddUserPermission_Response,
			opts...))
	mux.Methods("POST").Path("/remove-user-permission").Handler(
		http.NewServer(
			endpoints.RemoveUserPermissionEndpoint,
			_Decode_RemoveUserPermission_Request,
			_Encode_RemoveUserPermission_Response,
			opts...))
	return mux
}