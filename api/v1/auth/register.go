package auth

import (
	"context"

	"github.com/ajikamaludin/go-grpc_basic/pkg/v1/utils/constants"
	authpb "github.com/ajikamaludin/go-grpc_basic/proto/v1/auth"
)

func (s *Server) Register(ctx context.Context, req *authpb.Request) (*authpb.Response, error) {
	s.logger.Infof("[AUTH][REGISTER] SUCCESS")

	return &authpb.Response{
		Success: true,
		Code:    constants.SuccessCode,
		Desc:    constants.SuccesDesc,
	}, nil

}
