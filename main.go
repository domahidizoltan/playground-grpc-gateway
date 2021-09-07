package main

import (
	"fmt"
	"net"
	"log"

	"google.golang.org/grpc"

	"github.com/domahidizoltan/playground-grpc-gateway/car"
	pb "github.com/domahidizoltan/playground-grpc-gateway/generated/car"
)

const grpcPort = 2000

func main() {
	lis, err := net.Listen("tcp", fmt.Sprintf("0.0.0.0:%d", grpcPort))
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}

	server := grpc.NewServer()
	service := &car.CarService{}
	pb.RegisterCarServiceServer(server, service)

	withGateway(func() {
		log.Printf("grpc server listening on 0.0.0.0:%d", grpcPort)
		log.Fatal(server.Serve(lis)) 
	})

}