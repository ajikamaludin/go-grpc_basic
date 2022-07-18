package auth

import (
	"context"

	"github.com/ajikamaludin/go-grpc_basic/pkg/v1/jwt"
	"github.com/ajikamaludin/go-grpc_basic/pkg/v1/utils/constants"
	"github.com/ajikamaludin/go-grpc_basic/pkg/v1/utils/errors"
	authpb "github.com/ajikamaludin/go-grpc_basic/proto/v1/auth"
	"google.golang.org/grpc/codes"
)

func (s *Server) Login(ctx context.Context, req *authpb.Request) (*authpb.Response, error) {
	err := isValidRequest(LOGIN, req)
	if err != nil {
		s.logger.Errorf("[AUTH][LOGIN] ERROR %v", err)
		return nil, err
	}

	// TODO: database logic to match user

	auth, err := jwt.GenerateToken(s.config.Config, req.GetUserId())
	if err != nil {
		s.logger.Errorf("[AUTH][LOGIN] ERROR %v", err)
		return nil, errors.FormatError(codes.Internal, &errors.Response{
			Code: "1001",
			Msg:  err.Error(),
		})
	}

	s.logger.Infof("[AUTH][LOGIN] SUCCESS")

	return &authpb.Response{
		Success: true,
		Code:    constants.SuccessCode,
		Desc:    constants.SuccesDesc,
		Auth: &authpb.Auth{
			Type:           auth.Type,
			Access:         auth.Access,
			ExpiredPeriode: int32(auth.ExpiredPeriode),
			Refresh:        auth.Refresh,
		},
	}, nil

}
