package server

import (
	"context"
	"math/rand"
	"sync"
	"time"

	pb "github.com/zachgoldstein/serviceCarving/carvedService/rpc"
)

// Server implements the UniversalPaperclips service
type Server struct{}

// NewServer creates an instance of our server
func NewServer() *Server {
	return &Server{}
}

type job struct {
	times int
}

// Count finds the sum of random numbers added together N times
func (s *Server) Count(ctx context.Context, times *pb.Times) (*pb.Sum, error) {
	totalSum := int32(0)
	mutex := &sync.Mutex{}
	wg := &sync.WaitGroup{}

	jobSize := 100000
	numJobs := int(times.Times) / jobSize
	for w := 0; w < numJobs; w++ {
		wg.Add(1)
		go func() {
			sum := int32(0)
			r := rand.New(rand.NewSource(time.Now().Unix()))
			for i := 0; i <= jobSize; i++ {
				sum += r.Int31n(100)
			}
			mutex.Lock()
			totalSum += sum
			mutex.Unlock()
			wg.Done()
		}()
	}
	wg.Wait()

	return &pb.Sum{
		Sum: totalSum,
	}, nil
}
