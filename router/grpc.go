package router

import (
	"context"
	"fmt"
	"log"
	"net"
	"strconv"

	"github.com/ajikamaludin/go-grpc_basic/api/v1/auth"
	"github.com/ajikamaludin/go-grpc_basic/api/v1/health"
	"github.com/ajikamaludin/go-grpc_basic/configs"
	"github.com/ajikamaludin/go-grpc_basic/pkg/v1/config"
	"github.com/ajikamaludin/go-grpc_basic/pkg/v1/jwt"
	"github.com/ajikamaludin/go-grpc_basic/pkg/v1/utils/constants"
	"github.com/ajikamaludin/go-grpc_basic/pkg/v1/utils/errors"
	"github.com/ajikamaludin/go-grpc_basic/pkg/v1/utils/logger"
	authpb "github.com/ajikamaludin/go-grpc_basic/proto/v1/auth"
	hlpb "github.com/ajikamaludin/go-grpc_basic/proto/v1/health"
	grpc_middleware "github.com/grpc-ecosystem/go-grpc-middleware"
	"github.com/grpc-ecosystem/grpc-gateway/runtime"
	"github.com/sirupsen/logrus"
	"google.golang.org/grpc"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/metadata"
	"google.golang.org/grpc/reflection"
)

func NewGRPCServer(configs *configs.Configs, logger *logrus.Logger) error {
	listen, err := net.Listen("tcp", fmt.Sprintf(":%v", configs.Config.Server.Grpc.Port))
	if err != nil {
		return err
	}

	// register grpc service server
	// grpcServer := grpc.NewServer()

	grpcServer := grpc.NewServer(
		grpc.UnaryInterceptor(
			grpc_middleware.ChainUnaryServer(
				authInterceptor,
			),
		),
	)

	hlpb.RegisterHealthServiceServer(grpcServer, health.New(configs, logger))
	authpb.RegisterAuthServiceServer(grpcServer, auth.New(configs, logger))

	// add reflection service
	reflection.Register(grpcServer)

	// running gRPC server
	log.Println("[SERVER] GRPC server is running")

	grpcServer.Serve(listen)

	return nil
}

func authInterceptor(ctx context.Context, req interface{}, info *grpc.UnaryServerInfo, handler grpc.UnaryHandler) (interface{}, error) {
	fmt.Println(info)
	if !(info.FullMethod == constants.Endpoint_Auth_Login ||
		info.FullMethod == constants.Endpoint_Auth_Register) {
		// var userId string

		// // create config and logger
		// config, err := config.New()
		// if err != nil {
		// 	return nil, err
		// }

		// read header from incoming request
		headers, ok := metadata.FromIncomingContext(ctx)
		if !ok {
			return nil, errors.New("error get context")
		}

		if len(headers.Get("Authorization")) < 1 || headers.Get("Authorization")[0] == "" {
			return nil, errors.FormatError(codes.InvalidArgument, &errors.Response{
				Code: "1000",
				Msg:  "header Authorization is empty",
			})
		}

		config, err := config.GetInstance()
		if err != nil {
			return nil, err
		}

		//Verifying token
		_, err = jwt.ClaimToken(config, headers.Get("Authorization")[0], false)
		if err != nil {
			return nil, errors.FormatError(codes.Unauthenticated, &errors.Response{
				Code: strconv.Itoa(runtime.HTTPStatusFromCode(codes.Unauthenticated)),
				Msg:  err.Error(),
			})
		}
	}

	// store request log
	err := logger.StoreRestRequest(ctx, req, info, "")
	if err != nil {
		return nil, err
	}

	return handler(ctx, req)
}
