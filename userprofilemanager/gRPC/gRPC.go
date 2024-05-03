package gRPC

import (
	"context"
	"log"
	"net"
	"userprofilemanager/api"
	"userprofilemanager/core"
	pb "userprofilemanager/gen/go"
	"userprofilemanager/services"

	"google.golang.org/grpc"
)

type server struct {
	pb.UnimplementedProfileServiceServer
}

func (s *server) SignUp(ctx context.Context, req *pb.SignUpRequest) (*pb.SignUpResponse, error) {
	username := req.GetUsername()
	password := req.GetPassword()

	if api.NewUserProfile(username, password) {
		return &pb.SignUpResponse{Notification: "User created successfully"}, nil
	}
	return &pb.SignUpResponse{Notification: "Failed to create user"}, nil

}

func (s *server) UpdateProfile(ctx context.Context, req *pb.UpdateProfileRequest) (*pb.UpdateUserProfileResponse, error) {
	fullName := req.GetFullName()
	birthday := req.GetBirthday()
	address := req.GetAddress()
	email := req.GetEmail()
	phone := req.GetPhone()
	accesstoken := req.GetAccessToken()

	birthdayTime := services.TimestampProtoToTime(birthday)

	id, _, _, err := services.ValidateAuthToken(accesstoken)
	if err != nil {
		log.Println(err)
		return &pb.UpdateUserProfileResponse{Notification: "Failed to update profile"}, nil
	}

	if api.UpdateProfile(fullName, birthdayTime, address, email, phone, id) {
		return &pb.UpdateUserProfileResponse{Notification: "Profile updated successfully"}, nil
	}
	return &pb.UpdateUserProfileResponse{Notification: "Failed to update profile"}, nil

}

func (s *server) IsUserCorrect(ctx context.Context, req *pb.IsUserCorrectRequest) (*pb.IsUserCorrectResponse, error) {
	username := req.GetUsername()
	password := req.GetPassword()
	isCorrect, id := api.CheckUser(username, password)
	if isCorrect {
		return &pb.IsUserCorrectResponse{IsCorrect: isCorrect, Id: id}, nil
	}
	return &pb.IsUserCorrectResponse{IsCorrect: isCorrect, Id: id}, nil
}

func (s *server) ViewProfile(ctx context.Context, req *pb.ViewProfileRequest) (*pb.ViewProfileResponse, error) {
	accesstoken := req.GetAccessToken()

	id, _, _, err := services.ValidateAuthToken(accesstoken)
	if err != nil {
		log.Println(err)
		return &pb.ViewProfileResponse{
			FullName: "",
			Birthday: nil,
			Address:  "",
			Email:    "",
			Phone:    "",
		}, nil
	}

	var userProfile core.UserProfile = api.ViewProfile(id)

	return &pb.ViewProfileResponse{
		FullName: userProfile.FullName,
		Birthday: services.TimeToTimestampProto(userProfile.Birthday),
		Address:  userProfile.Address,
		Email:    userProfile.Email,
		Phone:    userProfile.Phone,
	}, nil

}

func GRPC() {
	lis, err := net.Listen("tcp", ":50051")
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}
	s := grpc.NewServer()
	pb.RegisterProfileServiceServer(s, &server{})
	log.Println("Server is running on port 50051")
	if err := s.Serve(lis); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}
}
