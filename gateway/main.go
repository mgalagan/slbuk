package main

import (
	"context"
	"flag"
	"fmt"
	"log"
	"net/http"

	"github.com/grpc-ecosystem/grpc-gateway/v2/runtime"
	grpcgw "github.com/mgalagan/slbuk/grpc"
	"google.golang.org/grpc"
)

func main() {
	gatewayAddress := flag.String("gateway-address", "localhost:9002", "The gateway server, format host:port")
	grpcServerAddress := flag.String("grpc-server-address", "localhost:9001", "gRPC server server format host:port")
	flag.Parse()

	ctx := context.Background()
	ctx, cancel := context.WithCancel(ctx)
	defer cancel()

	mux := runtime.NewServeMux()
	opts := []grpc.DialOption{grpc.WithInsecure()}
	err := grpcgw.RegisterSlbUkEntitiesHandlerFromEndpoint(ctx, mux, *grpcServerAddress, opts)
	if err != nil {
		log.Fatalf("Can't register endpoint %v", err)
	}
	fmt.Printf("Run gateway on %s to gRPC %s", *gatewayAddress, *grpcServerAddress)
	http.ListenAndServe(*gatewayAddress, mux)
}
