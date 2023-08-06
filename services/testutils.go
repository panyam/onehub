package services

import (
	"context"
	"log"
	"net"
	"testing"

	"google.golang.org/grpc"
	"google.golang.org/grpc/test/bufconn"
)

type DialerFunc = func(context.Context, string) (net.Conn, error)
type ServerInitFunc = func(server *grpc.Server)
type TestBodyFunc = func(ctx context.Context, conn *grpc.ClientConn)

func TestDialer() (*grpc.Server, DialerFunc) {
	listener := bufconn.Listen(1024 * 1024)
	server := grpc.NewServer()
	go func() {
		if err := server.Serve(listener); err != nil {
			log.Fatal(err)
		}
	}()
	return server, func(context.Context, string) (net.Conn, error) {
		return listener.Dial()
	}
}

func RunTest(t *testing.T, init_server ServerInitFunc, test_body TestBodyFunc) {
	ctx := context.Background()
	server, dialer_func := TestDialer()
	init_server(server)
	conn, err := grpc.DialContext(ctx, "", grpc.WithInsecure(), grpc.WithContextDialer(dialer_func))
	if err != nil {
		log.Fatal(err)
	}
	defer conn.Close()
	test_body(ctx, conn)
}
