package auth

import (
	"github.com/ajikamaludin/go-grpc_basic/configs"
	"github.com/ajikamaludin/go-grpc_basic/pkg/v1/utils/errors"
	authpb "github.com/ajikamaludin/go-grpc_basic/proto/v1/auth"
	"github.com/sirupsen/logrus"
	"google.golang.org/grpc/codes"
)

// Method is the methode type
type Method int

const (
	// SchemeCode of different Methods
	GET Method = iota
	REGISTER
	LOGIN
)

type Server struct {
	config *configs.Configs
	logger *logrus.Logger
}

func New(config *configs.Configs, logger *logrus.Logger) *Server {
	return &Server{
		config: config,
		logger: logger,
	}
}

// isValidRequest validates the status request
func isValidRequest(m Method, req *authpb.Request) error {
	switch m {
	case REGISTER, LOGIN:
		if req.GetUserId() == "" {
			return errors.FormatError(codes.InvalidArgument, &errors.Response{
				Code: "1000",
				Msg:  "userId is empty",
			})
		}
		if req.GetPassword() == "" {
			return errors.FormatError(codes.InvalidArgument, &errors.Response{
				Code: "1000",
				Msg:  "password is empty",
			})
		}
	}

	return nil
}
