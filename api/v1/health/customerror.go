package health

import (
	"context"

	"github.com/ajikamaludin/go-grpc_basic/pkg/v1/utils/errors"
	hlpb "github.com/ajikamaludin/go-grpc_basic/proto/v1/health"
	"github.com/golang/protobuf/ptypes/empty"
	"google.golang.org/grpc/codes"
)

func (s *Server) CallCustomError(ctx context.Context, req *empty.Empty) (*hlpb.Response, error) {

	// return nil, errors.FormatError(codes.NotFound, &errors.Response{
	// 	Code: "404",
	// 	Msg:  "Not Found",
	// })

	return nil, errors.FormatError(codes.Unauthenticated, &errors.Response{
		Code: "401",
		Msg:  "User not found",
	})
}
