package main

import (
	"log"
	"net"

	"github.com/bodocoder/grpc-calci/pb"
	"github.com/bodocoder/grpc-calci/service"
	"google.golang.org/grpc"
)

func main() {
	lis, err := net.Listen("tcp", ":3000")
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}

	grpcServer := grpc.NewServer()

	pb.RegisterCalculatorServiceServer(grpcServer, &service.CalciServer{})
	log.Printf("server listening at %v", lis.Addr())
	if err := grpcServer.Serve(lis); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}
}
