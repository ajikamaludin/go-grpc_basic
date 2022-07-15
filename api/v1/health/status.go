package health

import (
	"context"

	"github.com/ajikamaludin/go-grpc_basic/pkg/v1/utils/constants"
	hlpb "github.com/ajikamaludin/go-grpc_basic/proto/v1/health"
	"github.com/golang/protobuf/ptypes/empty"
)

func (s *Server) Status(ctx context.Context, req *empty.Empty) (*hlpb.Response, error) {

	return &hlpb.Response{
		Success: true,
		Code:    constants.SuccessCode,
		Desc:    constants.SuccesDesc,
	}, nil
}
