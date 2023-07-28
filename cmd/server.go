package main

import (
	"context"
	"flag"
	"log"
	"net"
	"net/http"

	"google.golang.org/grpc"

	"github.com/grpc-ecosystem/grpc-gateway/v2/runtime"
	v1 "github.com/panyam/onehub/gen/go/onehub/v1"
	svc "github.com/panyam/onehub/services"

	// This is needed to enable the use of the grpc_cli tool
	"google.golang.org/grpc/reflection"
)

var (
	addr    = flag.String("addr", ":9000", "Address to start the onehub grpc server on.")
	gw_addr = flag.String("gw_addr", ":8080", "Address to start the grpc gateway server on.")
)

func startGRPCServer(addr string) {
	// create new gRPC server
	server := grpc.NewServer()
	v1.RegisterTopicServiceServer(server, svc.NewTopicService(nil))
	v1.RegisterMessageServiceServer(server, svc.NewMessageService(nil))
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
	mux := runtime.NewServeMux()

	opts := []grpc.DialOption{grpc.WithInsecure()}
	err := v1.RegisterTopicServiceHandlerFromEndpoint(ctx, mux, grpc_addr, opts)
	if err != nil {
		log.Fatal(err)
	}
	if err := v1.RegisterMessageServiceHandlerFromEndpoint(ctx, mux, grpc_addr, opts); err != nil {
		log.Fatal(err)
	}

	log.Println("Starting grpc gateway server on: ", gw_addr)
	http.ListenAndServe(gw_addr, mux)
}

func main() {
	flag.Parse()
	go startGRPCServer(*addr)
	startGatewayServer(*gw_addr, *addr)
}
