package grpc

import (
	"google.golang.org/grpc"
	"log"
	"net"

	pb "github.com/Suhach/test_protoc-cont/proto/user"
)

func RegisterServer(s *grpc.Server, h *Handler) {
	pb.RegisterUserServiceServer(s, h)
}

func RunServer(h *Handler, port string) {
	lis, err := net.Listen("tcp", port)
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}

	s := grpc.NewServer()
	RegisterServer(s, h)
	log.Printf("User service running on port %s", port)
	if err := s.Serve(lis); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}
}
