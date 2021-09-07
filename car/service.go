package car

import (
	"context"
	"errors"
	"io"
	"log"
	"time"

	pb "github.com/domahidizoltan/playground-grpc-gateway/generated/car"
	"google.golang.org/grpc/metadata"
)

type CarService struct {
	pb.UnimplementedCarServiceServer
}

func (s *CarService) EchoCar(ctx context.Context, car *pb.Car) (*pb.CarEcho, error) {
	if err := authenticate(ctx); err != nil {
		log.Printf("EchoCar: %v", err)
		return nil, err
	}

	log.Printf("EchoCar %v", car)
	return echo(car), nil
}

func (s *CarService) EchoCars(server pb.CarService_EchoCarsServer) error {
	if err := authenticate(server.Context()); err != nil {
		log.Printf("EchoCars: %v", err)
		return err
	}

	log.Printf("EchoCars:")
	for {
		car, err := server.Recv()
		if err == io.EOF {
			break
		} else if err != nil {
			return err
		}
		log.Printf("\t%v", car)
		server.Send(echo(car))
	}
	return nil
}

func (s *CarService) mustEmbedUnimplementedCarServiceServer() {
	log.Printf("mustEmbedUnimplementedCarServiceServer")
}

func echo(car *pb.Car) *pb.CarEcho {
	return &pb.CarEcho{
		Car: car,
		Timestamp: time.Now().Unix(),
	}
}

func authenticate(ctx context.Context) error {
	md, ok := metadata.FromIncomingContext(ctx)
	if !ok {
		return errors.New("authentication failed: no metadata")
	}

	log.Printf("md %+v", md)
	key := md.Get("x-api-key")
	if len(key) ==0 || key[0] != "123" {
		return errors.New("authentication failed: x-api-key mismatch")
	}

	return nil
}
