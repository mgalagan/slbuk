package main

import (
	"context"
	"flag"
	"fmt"
	"log"
	"net"

	grpcpb "github.com/mgalagan/slbuk/grpc"
	"google.golang.org/grpc"
)

// SlbUkEntitiesServer is entities server implemetation
type SlbUkEntitiesServer struct {
	grpcpb.UnimplementedSlbUkEntitiesServer

	savedEntities []*grpcpb.Entity
}

// GetEntities returns entities
func (s *SlbUkEntitiesServer) GetEntities(context.Context, *grpcpb.Empty) (*grpcpb.GetEntityResponse, error) {
	return &grpcpb.GetEntityResponse{Entities: s.savedEntities}, nil
}

// GetEntitiesStream returns entities stream
func (s *SlbUkEntitiesServer) GetEntitiesStream(_ *grpcpb.Empty, stream grpcpb.SlbUkEntities_GetEntitiesStreamServer) error {
	for _, entity := range s.savedEntities {
		if err := stream.Send(entity); err != nil {
			return err
		}
	}

	return nil
}

// NewServer creates new EntitiesServer
func NewServer() *SlbUkEntitiesServer {
	s := &SlbUkEntitiesServer{}

	s.savedEntities = []*grpcpb.Entity{
		{Id: "idmax", Name: "Max"},
		{Id: "idsteave", Name: "Steave"},
		{Id: "idwill", Name: "Will"},
	}

	return s
}

func main() {
	serverAddress := flag.String("server-address", "localhost:9001", "The server port, format host:port")
	flag.Parse()

	lis, err := net.Listen("tcp", *serverAddress)
	if err != nil {
		log.Fatalf("Cant start listen port %v", err.Error())
	}

	grpcServer := grpc.NewServer()
	grpcpb.RegisterSlbUkEntitiesServer(grpcServer, NewServer())
	fmt.Printf("Server started on port %s", *serverAddress)
	grpcServer.Serve(lis)
}
