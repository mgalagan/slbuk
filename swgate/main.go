package main

import (
	"context"
	"flag"
	"log"
	"time"

	"github.com/go-openapi/loads"
	"github.com/go-openapi/runtime/middleware"

	"github.com/mgalagan/slbuk/swgate/models"
	"github.com/mgalagan/slbuk/swgate/restapi"
	"github.com/mgalagan/slbuk/swgate/restapi/operations"
	"github.com/mgalagan/slbuk/swgate/restapi/operations/slb_uk_entities"

	grpcpb "github.com/mgalagan/slbuk/grpc"
	"google.golang.org/grpc"
)

func convertToSwagger(response *grpcpb.GetEntityResponse) *models.GrpcGetEntityResponse {
	result := &models.GrpcGetEntityResponse{}
	for _, entity := range response.Entities {
		result.Entities = append(result.Entities, &models.GrpcEntity{ID: entity.Id, Name: entity.Name})
	}
	return result
}

func main() {
	var (
		gatewayPort       = flag.Int("port", 9002, "Gateway port")
		grpcServerAddress = flag.String("grpc-server-address", "localhost:9001", "gRPC server server format host:port")
	)
	flag.Parse()

	conn, err := grpc.Dial(*grpcServerAddress, grpc.WithInsecure(), grpc.WithBlock())
	if err != nil {
		log.Fatalf("fail to dial: %v", err)
	}
	defer conn.Close()
	client := grpcpb.NewSlbUkEntitiesClient(conn)

	swaggerSpec, err := loads.Analyzed(restapi.SwaggerJSON, "")
	if err != nil {
		log.Fatalln(err)
	}

	api := operations.NewSlbukswgateAPI(swaggerSpec)

	api.SlbUkEntitiesSlbUkEntitiesGetEntitiesHandler = slb_uk_entities.SlbUkEntitiesGetEntitiesHandlerFunc(
		func(params slb_uk_entities.SlbUkEntitiesGetEntitiesParams) middleware.Responder {
			ctx, cancel := context.WithTimeout(context.Background(), 1*time.Minute)
			defer cancel()
			entityResponse, _ := client.GetEntities(ctx, &grpcpb.Empty{})

			return slb_uk_entities.NewSlbUkEntitiesGetEntitiesOK().WithPayload(convertToSwagger(entityResponse))
		})

	server := restapi.NewServer(api)
	defer server.Shutdown()
	server.Port = *gatewayPort

	if err := server.Serve(); err != nil {
		log.Fatalln(err)
	}
}
