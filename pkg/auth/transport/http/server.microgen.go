// Code generated by microgen 1.0.5. DO NOT EDIT.

package transporthttp

import (
	transport "auth/pkg/auth/transport"
	http "github.com/go-kit/kit/transport/http"
	mux "github.com/gorilla/mux"
	http1 "net/http"
)

func NewHTTPHandler(endpoints *transport.EndpointsSet, opts ...http.ServerOption) http1.Handler {
	mux := mux.NewRouter()
	mux.Methods("POST").Path("/register").Handler(
		http.NewServer(
			endpoints.RegisterEndpoint,
			_Decode_Register_Request,
			_Encode_Register_Response,
			opts...))
	mux.Methods("POST").Path("/login").Handler(
		http.NewServer(
			endpoints.LoginEndpoint,
			_Decode_Login_Request,
			_Encode_Login_Response,
			opts...))
	mux.Methods("POST").Path("/public-key").Handler(
		http.NewServer(
			endpoints.PublicKeyEndpoint,
			_Decode_PublicKey_Request,
			_Encode_PublicKey_Response,
			opts...))
	mux.Methods("POST").Path("/get-permissions-for-service").Handler(
		http.NewServer(
			endpoints.GetPermissionsForServiceEndpoint,
			_Decode_GetPermissionsForService_Request,
			_Encode_GetPermissionsForService_Response,
			opts...))
	return mux
}
