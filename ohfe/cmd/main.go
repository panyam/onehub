package main

import (
	_ "html/template"
	"os"

	"github.com/panyam/onehub/ohfe"
)

func main() {
	grpcEndpoint := os.Getenv("ONEHUB_GRPC_ENDPOINT")
	if grpcEndpoint == "" {
		grpcEndpoint = ":9000"
	}
	web := ohfe.Web{GrpcEndpoint: grpcEndpoint}
	web.Start(":5050")
}
