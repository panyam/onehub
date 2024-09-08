package services

import (
	"context"
	"fmt"
	"log"
	"log/slog"
	"net"
	"net/http"
	"runtime/debug"
	"strings"

	"github.com/grpc-ecosystem/grpc-gateway/v2/runtime"
	ds "github.com/panyam/onehub/datastore"
	v1 "github.com/panyam/onehub/gen/go/onehub/v1"
	"go.opentelemetry.io/contrib/instrumentation/google.golang.org/grpc/otelgrpc"
	"google.golang.org/grpc"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/metadata"
	"google.golang.org/grpc/reflection"
	"google.golang.org/grpc/status"
)

/**
 * Create the gateway mux for for all our services.
 */
func CreateGatewayMux(ctx context.Context, grpc_addr string, srvErr chan error, opts ...grpc.DialOption) *runtime.ServeMux {
	mux := runtime.NewServeMux(
		runtime.WithMetadata(func(ctx context.Context, request *http.Request) metadata.MD {

			//
			// Step 2 - Extend the context
			//
			ctx = metadata.AppendToOutgoingContext(ctx)

			//
			// Step 3 - get the basic auth params
			//
			username, password, ok := request.BasicAuth()
			if !ok {
				return nil
			}
			md := metadata.Pairs()
			md.Append("OneHubUsername", username)
			md.Append("OneHubPassword", password)
			return md
		}))

	// Use the OpenTelemetry gRPC client interceptor for tracing
	conn, err := grpc.NewClient(grpc_addr, opts...)
	if err != nil {
		srvErr <- err
	}

	err = v1.RegisterTopicServiceHandler(ctx, mux, conn)
	if err != nil {
		srvErr <- err
	}
	if err = v1.RegisterMessageServiceHandler(ctx, mux, conn); err != nil {
		srvErr <- err
	}
	if err := v1.RegisterUserServiceHandler(ctx, mux, conn); err != nil {
		srvErr <- err
	}
	return mux
}

func StartGRPCServer(addr string, db *ds.OneHubDB, srvErr chan error, stopChan chan bool) {
	// create new gRPC server with otel enabled
	server := grpc.NewServer(
		grpc.StatsHandler(otelgrpc.NewServerHandler()),
		grpc.ChainUnaryInterceptor(
			ErrorLogger(),
			EnsureAuthIsValid,
		),
	)
	v1.RegisterTopicServiceServer(server, NewTopicService(db))
	v1.RegisterUserServiceServer(server, NewUserService(db))
	v1.RegisterMessageServiceServer(server, NewMessageService(db))
	if l, err := net.Listen("tcp", addr); err != nil {
		slog.Error("error in listening on port: ", "addr", addr, "err", err)
		srvErr <- err
	} else {
		// the gRPC server
		log.Printf("Starting grpc endpoint on %s:", addr)
		reflection.Register(server)

		go func() {
			<-stopChan
			server.GracefulStop()
		}()
		if err := server.Serve(l); err != nil {
			slog.Error("unable to start grpc server", "err", err)
			srvErr <- err
		}
	}
}

func EnsureAuthExists(ctx context.Context,
	method string, // Method to be invoked on the service (eg GetAlbums)
	req, // Request payload  (eg GetAlbumsRequest)
	reply interface{}, // Response payload (eg GetAlbumsResponse)
	cc *grpc.ClientConn, // the underlying connection to the service
	invoker grpc.UnaryInvoker, // The next handler
	opts ...grpc.CallOption) error {

	md, ok := metadata.FromOutgoingContext(ctx)
	if ok {
		usernames := md.Get("OneHubUsername")
		passwords := md.Get("OneHubPassword")
		log.Println("UP: ", usernames, passwords)
		if len(usernames) > 0 && len(passwords) > 0 {
			username := strings.TrimSpace(usernames[0])
			password := strings.TrimSpace(passwords[0])
			if len(username) > 0 && len(password) > 0 {
				// All fine - just call the invoker
				return invoker(ctx, method, req, reply, cc, opts...)
			}
		}
	}
	return status.Error(codes.NotFound, "BasicAuth params not found")
}

func EnsureAuthIsValid(ctx context.Context,
	req interface{},
	info *grpc.UnaryServerInfo,
	handler grpc.UnaryHandler) (resp interface{}, err error) {
	md, ok := metadata.FromIncomingContext(ctx)
	if ok {
		usernames := md.Get("OneHubUsername")
		passwords := md.Get("OneHubPassword")
		log.Println("UP: ", usernames, passwords)
		if len(usernames) > 0 && len(passwords) > 0 {
			username := strings.TrimSpace(usernames[0])
			password := strings.TrimSpace(passwords[0])

			// Make sure you use better passwords than this!
			if len(username) > 0 && password == fmt.Sprintf("%s123", username) {
				// All fine - just call the invoker
				return handler(ctx, req)
			}
		}
	}
	return nil, status.Error(codes.Unauthenticated, "Invalid username/password")
}

func ErrorLogger( /* Add configs here */ ) grpc.UnaryServerInterceptor {
	return func(ctx context.Context,
		req interface{},
		info *grpc.UnaryServerInfo,
		handler grpc.UnaryHandler) (resp interface{}, err error) {

		onPanic := func() {
			r := recover()
			if r != nil {
				err = status.Errorf(codes.Internal, "panic: %s", r)
				errmsg := fmt.Sprintf("[PANIC] %s\n\n%s", r, string(debug.Stack()))
				log.Println(errmsg)
			}
		}
		defer onPanic()

		resp, err = handler(ctx, req)
		errCode := status.Code(err)
		if errCode == codes.Unknown || errCode == codes.Internal {
			log.Println("Request handler returned an internal error - reporting it")
			return
		}
		return
	}
}
