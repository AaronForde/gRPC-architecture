package main

import (
	"context"
	"log"
	"net"

	"github.com/AaronForde/gRCP-architecture.git/pb"

	"google.golang.org/grpc"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/reflection"
	"google.golang.org/grpc/status"
)

type server struct {
	pb.UnimplementedCalculatorServer
}

func (s *server) Add(
	ctx context.Context, in *pb.CalculationRequest,
) (*pb.CalculationResponse, error) {
	return &pb.CalculationResponse{
		Result: in.A + in.B,
	}, nil
}

func (s *server) Divide(
	ctx context.Context, in *pb.CalculationRequest,
) (*pb.CalculationResponse, error) {

	if in.B == 0 {
		return nil, status.Errorf(
			codes.InvalidArgument, "cannot divide by zero",
		)
	}

	return &pb.CalculationResponse{
		Result: in.A / in.B,
	}, nil
}

func main() {
	listener, err := net.Listen("tcp", ":8080")
	if err != nil {
		log.Fatalln("Failed to create listener:", err)
	}

	s := grpc.NewServer()
	reflection.Register(s)

	pb.RegisterCalculatorServer(s, &server{})
	if err := s.Serve(listener); err != nil {
		log.Fatalln("Failed to serve:", err)
	}
}
