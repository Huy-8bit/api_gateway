package gRPC

import (
	"context"
	"log"
	"net"
	"userauthsystem/api"
	pb "userauthsystem/gen/go"

	"google.golang.org/grpc"
)

type server struct {
	pb.UnimplementedAuthServiceServer
}

func (s *server) SignIn(ctx context.Context, in *pb.SignInRequest) (*pb.SignInResponse, error) {
	log.Printf("Received: %v", in.GetUsername())
	username := in.GetUsername()
	password := in.GetPassword()
	token, err := api.SignIn(username, password)
	if err != nil {
		return nil, err
	}
	log.Printf("Token: %v", token)
	if token == "" {
		return &pb.SignInResponse{IsLogin: false, AccessToken: ""}, nil
	}
	return &pb.SignInResponse{IsLogin: true, AccessToken: token}, nil

}

func GRPC() {
	lis, err := net.Listen("tcp", ":50052")
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}
	s := grpc.NewServer()
	pb.RegisterAuthServiceServer(s, &server{})
	log.Println("Server is running on port 50052")
	if err := s.Serve(lis); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}
}
