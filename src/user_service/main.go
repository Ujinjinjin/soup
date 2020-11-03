package main

import (
	"flag"
	"fmt"
	"log"
	"net"

	"google.golang.org/grpc"

	pb "github.com/ujinjinjin/user_service/interface"
	"github.com/ujinjinjin/user_service/services"
)

var (
	address = flag.String("address", "localhost", "The server address")
	port = flag.Int("port", 10000, "The server port")
)

func main() {
	flag.Parse()
	log.Printf("Starting server at %s:%d", *address, *port)
	lis, err := net.Listen("tcp", fmt.Sprintf("%s:%d", *address, *port))
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}
	grpcServer := grpc.NewServer()

	pb.RegisterUserServiceServer(grpcServer, services.InitService())

	var serveResult = grpcServer.Serve(lis)
	if serveResult != nil {
		log.Fatalf("failed to serve: %v", err)
	}
}

