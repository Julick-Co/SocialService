package main

import (
	"SocialService/internal/social"
	pb "SocialService/proto/pkg/api/social"
	"log"
	"net"

	"google.golang.org/grpc"
)

func main() {
	address := ":9091"

	lis, err := net.Listen("tcp", address)
	if err != nil {
		log.Fatalf("listen error: %v", err)
	}

	server := grpc.NewServer()

	pb.RegisterSocialServiceServer(server, &social.SocialServiceImpl{})

	log.Printf("gRPC server running on %s\n", address)
	if err := server.Serve(lis); err != nil {
		log.Fatalf("serve error: %v", err)
	}
}
