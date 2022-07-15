package router

import (
	"context"
	"fmt"
	"log"
	"net/http"

	"github.com/ajikamaludin/go-grpc_basic/configs"
	"github.com/ajikamaludin/go-grpc_basic/pkg/v1/utils/constants"
	"github.com/ajikamaludin/go-grpc_basic/pkg/v1/utils/errors"
	hlpb "github.com/ajikamaludin/go-grpc_basic/proto/v1/health"
	"github.com/gorilla/handlers"
	"github.com/grpc-ecosystem/grpc-gateway/runtime"
	"github.com/sirupsen/logrus"
	"google.golang.org/grpc"
)

func NewHTTPServer(configs *configs.Configs, loggger *logrus.Logger) error {
	// create custom http errors
	runtime.HTTPError = errors.CustomHTTPError

	ctx := context.Background()
	ctx, cancel := context.WithCancel(ctx)
	defer cancel()

	// Connect to the GRPC server
	address := fmt.Sprintf("0.0.0.0:%v", configs.Config.Server.Grpc.Port)
	conn, err := grpc.Dial(address, grpc.WithInsecure())
	if err != nil {
		return err
	}
	defer conn.Close()

	// Create new grpc-gateway
	rmux := runtime.NewServeMux()

	// register gateway endpoints
	for _, f := range []func(ctx context.Context, mux *runtime.ServeMux, conn *grpc.ClientConn) error{
		// register grpc service handler
		hlpb.RegisterHealthServiceHandler,
	} {
		if err = f(ctx, rmux, conn); err != nil {
			return err
		}
	}

	mux := http.NewServeMux()
	mux.Handle("/", rmux)

	if configs.Config.Env != constants.EnvProduction {
		CreateSwagger(mux)
	}

	headerOk := handlers.AllowedHeaders([]string{"Accept", "Accept-Language", "Content-Language", "Content-Type", "X-Requested-With", "Content-Length", "Accept-Encoding", "X-CSRF-Token", "Authorization", "Timezone-Offset"})
	originsOk := handlers.AllowedOrigins([]string{"*"})
	methodsOk := handlers.AllowedMethods([]string{"GET", "HEAD", "POST", "PUT", "OPTIONS"})

	// running rest http server
	log.Println("[SERVER] REST HTTP server is ready")

	err = http.ListenAndServe(fmt.Sprintf("0.0.0.0:%v", configs.Config.Server.Rest.Port), handlers.CORS(headerOk, originsOk, methodsOk)(mux))
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

	// 	// load swagger-ui file
	// 	fs := http.FileServer(http.Dir("swagger/swagger-ui"))
	// 	gwmux.Handle("/api/health/docs/", http.StripPrefix("/api/health/docs", fs))
}
