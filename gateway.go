package main

import (
	"fmt"
	"context"
	"net/http"
	"log"

	"google.golang.org/grpc"


	"github.com/grpc-ecosystem/grpc-gateway/v2/runtime"
	pb "github.com/domahidizoltan/playground-grpc-gateway/generated/car"

)

const httpPort = 2020

func withGateway(serving func()) {
	go func() {
		serving()
	}()

	dialAddr := fmt.Sprintf("dns:///0.0.0.0:%d", grpcPort)
	// Create a client connection to the gRPC Server we just started.
	// This is where the gRPC-Gateway proxies the requests.
	conn, err := grpc.DialContext(
		context.Background(),
		dialAddr,
		grpc.WithInsecure(),
		grpc.WithBlock(),
	)
	if err != nil {
		fmt.Errorf("failed to dial server: %w", err)
	}
	log.Printf("grpc server listening on port localhost:%d", grpcPort)

	mux := runtime.NewServeMux()
	pb.RegisterCarServiceHandler(context.Background(), mux, conn)
	s := &http.Server{
		Addr:    fmt.Sprintf("0.0.0.0:%d", httpPort),
		Handler: http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			mux.ServeHTTP(w, r)
			return
		}),
	}
	log.Printf("http server listening on port localhost:%d", httpPort)
	if err := s.ListenAndServe(); err != http.ErrServerClosed {
		log.Fatalf("Failed to listen and serve: %v", err)
	}

}