package main

import (
	"context"
	"flag"
	"fmt"
	"log"
	"net"
	"net/http"
	"runtime/debug"
	"strings"

	"google.golang.org/grpc"

	"github.com/grpc-ecosystem/grpc-gateway/v2/runtime"
	cmdutils "github.com/panyam/onehub/cmd/utils"
	ds "github.com/panyam/onehub/datastore"
	v1 "github.com/panyam/onehub/gen/go/onehub/v1"
	svc "github.com/panyam/onehub/services"

	// This is needed to enable the use of the grpc_cli tool
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/metadata"
	"google.golang.org/grpc/reflection"
	"google.golang.org/grpc/status"
)

const DEFAULT_DB_ENDPOINT = "postgres://postgres:docker@localhost:5432/onehubdb"

var (
	addr    = flag.String("addr", ":9000", "Address to start the onehub grpc server on.")
	gw_addr = flag.String("gw_addr", ":9080", "Address to start the grpc gateway server on.")

	db_endpoint = flag.String("db_endpoint", "", fmt.Sprintf("Endpoint of DB where all topics/messages state are persisted.  Default value: ONEHUB_DB_ENDPOINT environment variable or %s", DEFAULT_DB_ENDPOINT))
)

func startGRPCServer(addr string, db *ds.OneHubDB) {
	// create new gRPC server
	server := grpc.NewServer(
		grpc.ChainUnaryInterceptor(
			ErrorLogger(),
			EnsureAuthIsValid,
		),
	)
	v1.RegisterTopicServiceServer(server, svc.NewTopicService(db))
	v1.RegisterUserServiceServer(server, svc.NewUserService(db))
	v1.RegisterMessageServiceServer(server, svc.NewMessageService(db))
	if l, err := net.Listen("tcp", addr); err != nil {
		log.Fatalf("error in listening on port %s: %v", addr, err)
	} else {
		// the gRPC server
		log.Printf("Starting grpc endpoint on %s:", addr)
		reflection.Register(server)
		if err := server.Serve(l); err != nil {
			log.Fatal("unable to start server", err)
		}
	}
}

func startGatewayServer(gw_addr, grpc_addr string) {
	ctx := context.Background()
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

	opts := []grpc.DialOption{grpc.WithInsecure()}
	err := v1.RegisterTopicServiceHandlerFromEndpoint(ctx, mux, grpc_addr, opts)
	if err != nil {
		log.Fatal(err)
	}
	if err := v1.RegisterMessageServiceHandlerFromEndpoint(ctx, mux, grpc_addr, opts); err != nil {
		log.Fatal(err)
	}
	if err := v1.RegisterUserServiceHandlerFromEndpoint(ctx, mux, grpc_addr, opts); err != nil {
		log.Fatal(err)
	}

	log.Println("Starting grpc gateway server on: ", gw_addr)
	http.ListenAndServe(gw_addr, mux)
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

func main() {
	flag.Parse()
	ohdb := OpenOHDB()
	go startGRPCServer(*addr, ohdb)
	startGatewayServer(*gw_addr, *addr)
}

func OpenOHDB() *ds.OneHubDB {
	if *db_endpoint == "" {
		*db_endpoint = cmdutils.GetEnvOrDefault("ONEHUB_DB_ENDPOINT", DEFAULT_DB_ENDPOINT)
	}
	db, err := cmdutils.OpenDB(*db_endpoint)
	if err != nil {
		log.Fatal(err)
		panic(err)
	}
	return ds.NewOneHubDB(db)
}
