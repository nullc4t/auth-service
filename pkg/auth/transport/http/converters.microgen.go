// Code generated by microgen 1.0.5. DO NOT EDIT.

// Please, do not change functions names!
package transporthttp

import (
	"bytes"
	"context"
	"encoding/json"
	transport "github.com/nullc4t/auth-service/pkg/auth/transport"
	"io/ioutil"
	"net/http"
	"path"
)

func CommonHTTPRequestEncoder(_ context.Context, r *http.Request, request interface{}) error {
	var buf bytes.Buffer
	if err := json.NewEncoder(&buf).Encode(request); err != nil {
		return err
	}
	r.Body = ioutil.NopCloser(&buf)
	return nil
}

func CommonHTTPResponseEncoder(_ context.Context, w http.ResponseWriter, response interface{}) error {
	w.Header().Set("Content-Type", "application/json; charset=utf-8")
	return json.NewEncoder(w).Encode(response)
}

func _Decode_Register_Request(_ context.Context, r *http.Request) (interface{}, error) {
	var req transport.RegisterRequest
	err := json.NewDecoder(r.Body).Decode(&req)
	return &req, err
}

func _Decode_Login_Request(_ context.Context, r *http.Request) (interface{}, error) {
	var req transport.LoginRequest
	err := json.NewDecoder(r.Body).Decode(&req)
	return &req, err
}

func _Decode_PublicKey_Request(_ context.Context, r *http.Request) (interface{}, error) {
	var req transport.PublicKeyRequest
	err := json.NewDecoder(r.Body).Decode(&req)
	return &req, err
}

func _Decode_GetPermissionsForService_Request(_ context.Context, r *http.Request) (interface{}, error) {
	var req transport.GetPermissionsForServiceRequest
	err := json.NewDecoder(r.Body).Decode(&req)
	return &req, err
}

func _Decode_Register_Response(_ context.Context, r *http.Response) (interface{}, error) {
	var resp transport.RegisterResponse
	err := json.NewDecoder(r.Body).Decode(&resp)
	return &resp, err
}

func _Decode_Login_Response(_ context.Context, r *http.Response) (interface{}, error) {
	var resp transport.LoginResponse
	err := json.NewDecoder(r.Body).Decode(&resp)
	return &resp, err
}

func _Decode_PublicKey_Response(_ context.Context, r *http.Response) (interface{}, error) {
	var resp transport.PublicKeyResponse
	err := json.NewDecoder(r.Body).Decode(&resp)
	return &resp, err
}

func _Decode_GetPermissionsForService_Response(_ context.Context, r *http.Response) (interface{}, error) {
	var resp transport.GetPermissionsForServiceResponse
	err := json.NewDecoder(r.Body).Decode(&resp)
	return &resp, err
}

func _Encode_Register_Request(ctx context.Context, r *http.Request, request interface{}) error {
	r.URL.Path = path.Join(r.URL.Path, "register")
	return CommonHTTPRequestEncoder(ctx, r, request)
}

func _Encode_Login_Request(ctx context.Context, r *http.Request, request interface{}) error {
	r.URL.Path = path.Join(r.URL.Path, "login")
	return CommonHTTPRequestEncoder(ctx, r, request)
}

func _Encode_PublicKey_Request(ctx context.Context, r *http.Request, request interface{}) error {
	r.URL.Path = path.Join(r.URL.Path, "public-key")
	return CommonHTTPRequestEncoder(ctx, r, request)
}

func _Encode_GetPermissionsForService_Request(ctx context.Context, r *http.Request, request interface{}) error {
	r.URL.Path = path.Join(r.URL.Path, "get-permissions-for-service")
	return CommonHTTPRequestEncoder(ctx, r, request)
}

func _Encode_Register_Response(ctx context.Context, w http.ResponseWriter, response interface{}) error {
	return CommonHTTPResponseEncoder(ctx, w, response)
}

func _Encode_Login_Response(ctx context.Context, w http.ResponseWriter, response interface{}) error {
	return CommonHTTPResponseEncoder(ctx, w, response)
}

func _Encode_PublicKey_Response(ctx context.Context, w http.ResponseWriter, response interface{}) error {
	return CommonHTTPResponseEncoder(ctx, w, response)
}

func _Encode_GetPermissionsForService_Response(ctx context.Context, w http.ResponseWriter, response interface{}) error {
	return CommonHTTPResponseEncoder(ctx, w, response)
}
