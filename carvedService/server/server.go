package server

import (
	"context"
	"math/rand"

	pb "github.com/zachgoldstein/serviceCarving/carvedService/rpc"
)

// Server implements the UniversalPaperclips service
type Server struct{}

// NewServer creates an instance of our server
func NewServer() *Server {
	return &Server{}
}

// Count finds the sum of random numbers added together N times
func (s *Server) Count(ctx context.Context, times *pb.Times) (*pb.Sum, error) {
	sum := int32(0)
	for i := 0; i <= int(times.Times); i++ {
		sum += rand.Int31n(100)
	}

	return &pb.Sum{
		Sum: sum,
	}, nil
}
