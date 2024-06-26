package main

import (
	_ "html/template"
	"os"

	"github.com/panyam/onehub/ohfe"
)

func main() {
	web := ohfe.Web{GrpcEndpoint: os.Getenv("ONEHUB_GRPC_ENDPOINT"),
		ApiEndpoint: os.Getenv("ONEHUB_API_ENDPOINT")}
	web.Start(":5050")
}
