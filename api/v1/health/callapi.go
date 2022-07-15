package health

import (
	"context"

	"github.com/ajikamaludin/go-grpc_basic/pkg/v1/utils/constants"
	hlpb "github.com/ajikamaludin/go-grpc_basic/proto/v1/health"
	"github.com/ajikamaludin/go-grpc_basic/services/v1/jsonplaceholder"
	"github.com/golang/protobuf/ptypes/empty"
)

func (s *Server) CallApi(ctx context.Context, req *empty.Empty) (*hlpb.Response, error) {
	// call reqres
	res, err := jsonplaceholder.GetListUser()

	if err != nil {
		s.logger.Errorf("[HEALTH][GET] ERROR %v", err)
	}

	var data []*hlpb.Data
	for _, v := range *res {
		data = append(data, &hlpb.Data{
			Id:       uint32(v.ID),
			Name:     v.Name,
			Username: v.Username,
			Email:    v.Email,
			Phone:    v.Phone,
			Website:  v.Website,
		})
	}

	return &hlpb.Response{
		Success: true,
		Code:    constants.SuccessCode,
		Desc:    constants.SuccesDesc,
		Data:    data,
	}, nil
}
