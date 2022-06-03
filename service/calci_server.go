package service

import (
	"context"

	"github.com/bodocoder/grpc-calci/pb"
)

type CalciServer struct {
	pb.UnimplementedCalculatorServiceServer
}

func (server *CalciServer) PerformAddition(ctx context.Context, req *pb.CalculateRequest) (*pb.CalculateResponse, error) {
	return &pb.CalculateResponse{Res: req.X + req.Y}, nil
}

func (server *CalciServer) PerformSubtraction(ctx context.Context, req *pb.CalculateRequest) (*pb.CalculateResponse, error) {
	return &pb.CalculateResponse{Res: req.X - req.Y}, nil
}
