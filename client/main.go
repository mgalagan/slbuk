package main

import (
	"context"
	"flag"
	"io"
	"log"
	"time"

	grpcpb "github.com/mgalagan/slbuk/grpc"
	"google.golang.org/grpc"
)

func main() {
	serverAddr := flag.String("server-address", "localhost:9001", "The server address in the format of host:port")
	flag.Parse()

	conn, err := grpc.Dial(*serverAddr, grpc.WithInsecure())
	if err != nil {
		log.Fatalf("fail to dial: %v", err)
	}
	defer conn.Close()
	client := grpcpb.NewSlbUkEntitiesClient(conn)

	ctx, cancel := context.WithTimeout(context.Background(), 20*time.Second)
	defer cancel()

	stream, err := client.GetEntitiesStream(ctx, &grpcpb.Empty{})
	if err != nil {
		log.Fatalf("can't stream stream, %v", err)
	}
	for {
		entity, err := stream.Recv()
		if err == io.EOF {
			break
		}
		if err != nil {
			log.Fatalf("can't get entity %v", err)
		}
		log.Printf("Entity, id: %s, name: %s", entity.Id, entity.Name)
	}
}
