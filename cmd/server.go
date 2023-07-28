package main

import (
	"flag"
	"log"
	"net"

	"google.golang.org/grpc"

	v1 "github.com/panyam/onehub/gen/go/onehub/v1"
	svc "github.com/panyam/onehub/services"

	// This is needed to enable the use of the grpc_cli tool
	"google.golang.org/grpc/reflection"
)

var (
	addr = flag.String("addr", ":9000", "Address to start the onehub grpc server on.")
)

func startGRPCServer(addr string) {
	// create new gRPC server
	server := grpc.NewServer()
	v1.RegisterTopicServiceServer(server, svc.NewTopicService(nil))
	// v1.RegisterMessageServiceServer(server, svc.NewMessageService(nil))
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

func main() {
	flag.Parse()
	startGRPCServer(*addr)
}
