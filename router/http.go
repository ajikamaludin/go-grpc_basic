package router

import (
	"context"
	"fmt"
	"log"
	"net/http"

	"github.com/ajikamaludin/go-grpc_basic/configs"
	"github.com/ajikamaludin/go-grpc_basic/pkg/v1/utils/constants"
	"github.com/ajikamaludin/go-grpc_basic/pkg/v1/utils/errors"
	"github.com/ajikamaludin/go-grpc_basic/pkg/v1/utils/logger"
	"github.com/gorilla/handlers"
	"github.com/grpc-ecosystem/grpc-gateway/runtime"
	"github.com/sirupsen/logrus"
	"google.golang.org/grpc"

	"github.com/golang/protobuf/proto"

	authpb "github.com/ajikamaludin/go-grpc_basic/proto/v1/auth"
	hlpb "github.com/ajikamaludin/go-grpc_basic/proto/v1/health"
)

func NewHTTPServer(configs *configs.Configs, loggger *logrus.Logger) error {
	// create custom http errors
	runtime.HTTPError = errors.CustomHTTPError

	ctx := context.Background()
	ctx, cancel := context.WithCancel(ctx)
	defer cancel()

	// Connect to the GRPC server
	conn, err := grpc.Dial(fmt.Sprintf("0.0.0.0:%v", configs.Config.Server.Grpc.Port), grpc.WithInsecure())
	if err != nil {
		return err
	}
	defer conn.Close()

	// Create new grpc-gateway
	rmux := runtime.NewServeMux(runtime.WithForwardResponseOption(httpResponseModifier))

	// register gateway endpoints
	for _, f := range []func(ctx context.Context, mux *runtime.ServeMux, conn *grpc.ClientConn) error{
		// register grpc service handler
		hlpb.RegisterHealthServiceHandler,
		authpb.RegisterAuthServiceHandler,
	} {
		if err = f(ctx, rmux, conn); err != nil {
			return err
		}
	}

	// create http server mux
	mux := http.NewServeMux()
	mux.Handle("/", rmux)

	// run swagger server
	if configs.Config.Env != constants.EnvProduction {
		CreateSwagger(mux)
	}

	// Where ORIGIN_ALLOWED is like `scheme://dns[:port]`, or `*` (insecure)
	headersOk := handlers.AllowedHeaders([]string{"Accept", "Accept-Language", "Content-Language", "Content-Type", "X-Requested-With", "Content-Length", "Accept-Encoding", "X-CSRF-Token", "Authorization", "Timezone-Offset"})
	originsOk := handlers.AllowedOrigins([]string{"*"})
	methodsOk := handlers.AllowedMethods([]string{"GET", "HEAD", "POST", "PUT", "OPTIONS"})

	// running rest http server
	log.Println("[SERVER] REST HTTP server is ready")

	err = http.ListenAndServe(fmt.Sprintf("0.0.0.0:%v", configs.Config.Server.Rest.Port), handlers.CORS(headersOk, originsOk, methodsOk)(mux))
	if err != nil {
		return err
	}

	return nil
}

// // CreateSwagger creates the swagger server serve mux.
func CreateSwagger(gwmux *http.ServeMux) {
	// 	// register swagger service server
	gwmux.HandleFunc("/api/health/docs.json", func(w http.ResponseWriter, r *http.Request) {
		http.ServeFile(w, r, "swagger/docs.json")
	})
}

func httpResponseModifier(ctx context.Context, w http.ResponseWriter, p proto.Message) error {
	// store response log
	err := logger.StoreRestResponse(ctx, w, p)
	if err != nil {
		return err
	}

	return nil
}
