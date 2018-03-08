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

type countMutex struct {
	mu    *sync.Mutex
	count int32
}

func CountWorker(count *countMutex, wg *sync.WaitGroup, jobChan <-chan int) {
	r := rand.New(rand.NewSource(time.Now().Unix()))
	for j := range jobChan {
		sum := int32(0)
		for i := 0; i <= j; i++ {
			sum += r.Int31n(100)
		}
		count.mu.Lock()
		count.count += sum
		count.mu.Unlock()
		wg.Done()
	}
}

// Count finds the sum of random numbers added together N times
func (s *Server) Count(ctx context.Context, times *pb.Times) (*pb.Sum, error) {
	// totalSum := int32(0)
	wg := &sync.WaitGroup{}

	jobSize := 100000
	numWorkers := 10
	countMutex := &countMutex{
		mu:    &sync.Mutex{},
		count: 0,
	}
	jobChan := make(chan int, 100)
	for i := 0; i < numWorkers; i++ {
		go CountWorker(countMutex, wg, jobChan)
	}

	numJobs := int(times.Times) / jobSize
	for w := 0; w < numJobs; w++ {
		wg.Add(1)
		jobChan <- jobSize
	}
	close(jobChan)

	// for w := 0; w < numJobs; w++ {
	// 	wg.Add(1)
	// 	go func() {
	// 		sum := int32(0)
	// 		r := rand.New(rand.NewSource(time.Now().Unix()))
	// 		for i := 0; i <= jobSize; i++ {
	// 			sum += r.Int31n(100)
	// 		}
	// 		mutex.Lock()
	// 		totalSum += sum
	// 		mutex.Unlock()
	// 		wg.Done()
	// 	}()
	// }
	wg.Wait()

	return &pb.Sum{
		Sum: countMutex.count,
	}, nil
}
