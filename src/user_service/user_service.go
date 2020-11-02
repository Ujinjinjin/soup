package main

import (
	"context"
	"flag"
	"fmt"
	"log"
	"net"
	//"sync"

	"google.golang.org/grpc"

	pb "github.com/ujinjinjin/user_service/interface"
)

//var (
//	address = flag.String("port", "localhost", "The server address")
//	port = flag.Int("port", 10000, "The server port")
//)

type userServiceServer struct {
	pb.UnimplementedUserServiceServer
}

func newServer() *userServiceServer {
	s := &userServiceServer{}
	return s
}

// CreateUser creates user with specified fields
func (s *userServiceServer) CreateUser(context context.Context, request *pb.CreateUserRequest) (*pb.CreateUserReply, error) {

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

func main() {
	flag.Parse()
	log.Printf("Starting server localhost:10000...")
	lis, err := net.Listen("tcp", fmt.Sprint("localhost:10000"))
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}
	grpcServer := grpc.NewServer()

	pb.RegisterUserServiceServer(grpcServer, newServer())

	var serveResult = grpcServer.Serve(lis)
	if serveResult != nil {
		log.Fatalf("failed to serve: %v", err)
	}
}

