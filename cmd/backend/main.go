package main

import (
	"context"
	"flag"
	"fmt"
	"log"
	"net"
	"net/http"
	"os"
	"os/signal"

	"go.opentelemetry.io/contrib/instrumentation/google.golang.org/grpc/otelgrpc"
	"go.opentelemetry.io/contrib/instrumentation/net/http/otelhttp"
	"google.golang.org/grpc"

	cmdutils "github.com/panyam/onehub/cmd/utils"
	ds "github.com/panyam/onehub/datastore"
	svc "github.com/panyam/onehub/services"

	// This is needed to enable the use of the grpc_cli tool
	"github.com/grpc-ecosystem/grpc-gateway/v2/runtime"
	_ "github.com/panyam/onehub/obs"
	"google.golang.org/grpc/credentials/insecure"
)

const DEFAULT_DB_ENDPOINT = "postgres://postgres:docker@localhost:5432/onehubdb"

var (
	addr    = flag.String("addr", ":9000", "Address to start the onehub grpc server on.")
	gw_addr = flag.String("gw_addr", ":9080", "Address to start the grpc gateway server on.")

	db_endpoint = flag.String("db_endpoint", "", fmt.Sprintf("Endpoint of DB where all topics/messages state are persisted.  Default value: ONEHUB_DB_ENDPOINT environment variable or %s", DEFAULT_DB_ENDPOINT))
)

func StartGatewayServer(ctx context.Context, mux *runtime.ServeMux, gw_addr string, srvErr chan error, stopChan chan bool) {
	log.Println("Starting grpc gateway server on: ", gw_addr)
	server := &http.Server{
		Addr:        gw_addr,
		BaseContext: func(_ net.Listener) context.Context { return ctx },
		Handler: otelhttp.NewHandler(mux, "gateway", otelhttp.WithSpanNameFormatter(func(operation string, r *http.Request) string {
			return fmt.Sprintf("%s %s %s", operation, r.Method, r.URL.Path)
		})),
	}

	go func() {
		<-stopChan
		if err := server.Shutdown(context.Background()); err != nil {
			log.Fatalln(err)
		}
	}()
	srvErr <- server.ListenAndServe()
}

func main() {
	flag.Parse()

	// Handle SIGINT (CTRL+C) gracefully.
	ctx, stop := signal.NotifyContext(context.Background(), os.Interrupt)
	defer stop()

	collectorAddr := cmdutils.GetEnvOrDefault("OTEL_COLLECTOR_ADDR", "otel-collector:4317")
	conn, err := grpc.NewClient(collectorAddr,
		// Note the use of insecure transport here. TLS is recommended in production.
		grpc.WithTransportCredentials(insecure.NewCredentials()),
	)
	if err != nil {
		log.Println("failed to create gRPC connection to collector: %w", err)
		return
	}
	setup := NewOTELSetupWithCollector(conn)
	err = setup.Setup(ctx)
	if err != nil {
		log.Println("error setting up otel: ", err)
	}

	defer func() {
		err = setup.Shutdown(context.Background())
	}()

	ohdb := OpenOHDB()

	srvErr := make(chan error, 2)

	httpSrvChan := make(chan bool)
	grpcSrvChan := make(chan bool)
	trclient := grpc.WithStatsHandler(otelgrpc.NewClientHandler())
	mux := svc.CreateGatewayMux(ctx, *addr, srvErr, grpc.WithTransportCredentials(insecure.NewCredentials()), trclient)
	go svc.StartGRPCServer(*addr, ohdb, srvErr, grpcSrvChan)
	go StartGatewayServer(ctx, mux, *gw_addr, srvErr, httpSrvChan)

	// Wait for interruption.
	select {
	case err = <-srvErr:
		log.Println("Server error: ", err)
		// Error when starting HTTP server or GRPC server
		return
	case <-ctx.Done():
		// Wait for first CTRL+C.
		// Stop receiving signal notifications as soon as possible.
		stop()
	}

	// When Shutdown is called, ListenAndServe immediately returns ErrServerClosed.
	httpSrvChan <- true
	grpcSrvChan <- true
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
