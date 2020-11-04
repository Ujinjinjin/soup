package services

import (
	"context"
	"log"

	pb "github.com/ujinjinjin/user_service/interface"
	"github.com/ujinjinjin/user_service/repository"
)

type UserRpcService struct {
	pb.UnimplementedUserServiceServer
	*repository.UserRepository
}

func InitService(userRepository *repository.UserRepository) *UserRpcService {
	return &UserRpcService{
		UserRepository: userRepository,
	}
}

// CreateUser creates user with specified fields
func (s *UserRpcService) CreateUser(context context.Context, request *pb.CreateUserRequest) (*pb.CreateUserReply, error) {

	log.Print("New user created:")
	log.Printf("\tEmail: %s", request.Email)
	log.Printf("\tFirstName: %s", request.FirstName)
	log.Printf("\tLastName: %s", request.LastName)
	log.Printf("\tMiddleName: %s", request.MiddleName)

	testResult, err := s.Test()
	if err != nil {
		return nil, err
	}

	log.Printf("\tRepository response: %s", testResult)

	var result = &pb.CreateUserReply{
		Id: 1,
	}
	return result, nil
}
