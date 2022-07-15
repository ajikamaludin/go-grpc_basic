package health

import (
	"context"
	"errors"

	"github.com/ajikamaludin/go-grpc_basic/pkg/v1/postgres"
	"github.com/ajikamaludin/go-grpc_basic/pkg/v1/utils/constants"
	hlpb "github.com/ajikamaludin/go-grpc_basic/proto/v1/health"
	"github.com/golang/protobuf/ptypes/empty"
)

func (s *Server) CallDb(ctx context.Context, req *empty.Empty) (*hlpb.Response, error) {
	// select db
	rows, err := s.config.Pg.CustomMainSelect(&postgres.CustomMain{})

	if err != nil {
		s.logger.Errorf("[HEALTH][GET] ERROR %v", err)
	}

	if len(rows) <= 0 {
		s.logger.Errorf("[HEALTH][GET] EMPTY ROW")
		return nil, errors.New("Invoke empty rows")
	}
	var data []*hlpb.Data
	for _, v := range rows {
		data = append(data, &hlpb.Data{
			Email:    v.UserId,
			Username: v.Pass,
		})
	}

	return &hlpb.Response{
		Success: true,
		Code:    constants.SuccessCode,
		Desc:    constants.SuccesDesc,
		Data:    data,
	}, nil
}
