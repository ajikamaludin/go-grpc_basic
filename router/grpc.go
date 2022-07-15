package router

import (
	"fmt"
	"log"
	"net"

	"github.com/ajikamaludin/go-grpc_basic/api/v1/health"
	"github.com/ajikamaludin/go-grpc_basic/configs"
	hlpb "github.com/ajikamaludin/go-grpc_basic/proto/v1/health"
	"github.com/sirupsen/logrus"
	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
)

func NewGRPCServer(configs *configs.Configs, logger *logrus.Logger) error {
	listen, err := net.Listen("tcp", fmt.Sprintf(":%v", configs.Config.Server.Grpc.Port))
	if err != nil {
		return err
	}

	// register grpc service server
	grpcServer := grpc.NewServer()

	hlpb.RegisterHealthServiceServer(grpcServer, health.New(configs, logger))

	// add reflection service
	reflection.Register(grpcServer)

	// running gRPC server
	log.Println("[SERVER] GRPC server is running")

	grpcServer.Serve(listen)

	return nil
}
