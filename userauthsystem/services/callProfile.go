package services

import (
	"context"
	"log"
	"time"
	pb "userauthsystem/gen/go"

	"google.golang.org/grpc"
)

func CallIsUserCorrect(host string, isusername string, ispassword string) (bool, string) {
	conn, err := grpc.Dial(host, grpc.WithInsecure(), grpc.WithBlock())
	if err != nil {
		log.Fatalf("Connection error: %v", err)
		return false, ""
	}
	defer conn.Close()
	c := pb.NewProfileServiceClient(conn)

	ctx, cancel := context.WithTimeout(context.Background(), time.Second)
	defer cancel()
	r, err := c.IsUserCorrect(ctx, &pb.IsUserCorrectRequest{Username: isusername, Password: ispassword})
	if err != nil {
		log.Fatalf("Fail IsUserCorrect: %v", err)
		return false, ""
	}
	return r.GetIsCorrect(), r.GetId()
}
