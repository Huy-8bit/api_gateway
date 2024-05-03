package gRPC

import (
	"context"
	"fmt"
	"log"
	"net"

	pb "userauthsystem/gen/go"

	"google.golang.org/grpc"
)

type server struct {
	pb.UnimplementedHelloServiceServer
}

func (s *server) SayHello(ctx context.Context, in *pb.HelloRequest) (*pb.HelloResponse, error) {
	fmt.Println("SayHello")
	log.Println("Request received from client: ", in.Name)
	return &pb.HelloResponse{Message: "Hello " + in.Name}, nil
}

func GRPC() {
	lis, err := net.Listen("tcp", ":9000")
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}
	s := grpc.NewServer()
	pb.RegisterHelloServiceServer(s, &server{})
	log.Println("Server is running on port 9000")
	if err := s.Serve(lis); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}
}
