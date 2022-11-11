// Code generated by microgen 1.0.5. DO NOT EDIT.

// It is better for you if you do not change functions names!
// This file will never be overwritten.
package transportgrpc

import (
	pb "auth/pkg/auth/proto"
	types "auth/pkg/types"
)

func PtrTypesAccessTokenToProto(token *types.AccessToken) (*pb.AccessToken, error) {
	return &pb.AccessToken{AccessToken: token.AccessToken}, nil
}

func ProtoToPtrTypesAccessToken(protoToken *pb.AccessToken) (*types.AccessToken, error) {
	return &types.AccessToken{AccessToken: protoToken.AccessToken}, nil
}

func ListByteToProto(pub []byte) ([]byte, error) {
	return pub, nil
}

func ProtoToListByte(protoPub []byte) ([]byte, error) {
	return protoPub, nil
}