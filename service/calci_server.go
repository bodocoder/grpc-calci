package service

import (
	"context"

	"github.com/bodocoder/grpc-calci/pb"
)

type CalciServer struct {
	pb.UnimplementedCalculatorServiceServer
}

func (server *CalciServer) PerformAddition(ctx context.Context, in *pb.CalculateRequest) (*pb.CalculateResponse, error) {
	return &pb.CalculateResponse{Res: in.X + in.Y}, nil
}
