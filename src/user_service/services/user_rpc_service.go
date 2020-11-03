package services

import (
	"context"
	"log"

	pb "github.com/ujinjinjin/user_service/interface"
)

type UserRpcService struct {
	pb.UnimplementedUserServiceServer
}

func InitService() *UserRpcService {
	return &UserRpcService{}
}

// CreateUser creates user with specified fields
func (s *UserRpcService) CreateUser(context context.Context, request *pb.CreateUserRequest) (*pb.CreateUserReply, error) {

	log.Print("New user created:")
	log.Printf("\tEmail: %s", request.Email)
	log.Printf("\tFirstName: %s", request.FirstName)
	log.Printf("\tLastName: %s", request.LastName)
	log.Printf("\tMiddleName: %s", request.MiddleName)

	var result = &pb.CreateUserReply{
		Id: 1,
	}
	return result, nil
}
