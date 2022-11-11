// Code generated by microgen 1.0.5. DO NOT EDIT.

package transporthttp

import (
	transport "auth/auth/transport"
	httpkit "github.com/go-kit/kit/transport/http"
	"net/url"
)

func NewHTTPClient(u *url.URL, opts ...httpkit.ClientOption) transport.EndpointsSet {
	return transport.EndpointsSet{
		LoginEndpoint: httpkit.NewClient(
			"POST", u,
			_Encode_Login_Request,
			_Decode_Login_Response,
			opts...,
		).Endpoint(),
		PublicKeyEndpoint: httpkit.NewClient(
			"POST", u,
			_Encode_PublicKey_Request,
			_Decode_PublicKey_Response,
			opts...,
		).Endpoint(),
		RegisterEndpoint: httpkit.NewClient(
			"POST", u,
			_Encode_Register_Request,
			_Decode_Register_Response,
			opts...,
		).Endpoint(),
	}
}
