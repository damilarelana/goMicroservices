package main

import (
	"context"
	"fmt"
	"log"
	"net"

	mf "github.com/damilarelana/goMicroservice/mathFunctions"
	ms "github.com/damilarelana/goMicroservice/mathService"
	"github.com/pkg/errors"
	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
)

// set the Port that the server would be listening/serving from
const port = ":9090"

// Server to run the service [from mathService.proto]
type mathService struct {
	// ...
}

// Add service handler method
func (server *mathService) Add(ctx context.Context, r *ms.AddRequest) (*ms.AddResponse, error) {
	log.Printf("Received input integer %v and %v", r.X, r.Y)
	return &ms.AddResponse{Addition: mf.Add(r.X, r.Y)}, nil
}

// Average service handler method
func (server *mathService) Average(ctx context.Context, r *ms.AverageRequest) (*ms.AverageResponse, error) {
	log.Printf("Received input array %v", r.Array)
	return &ms.AverageResponse{Average: mf.Average(r.Array)}, nil
}

// Max service handler method
func (server *mathService) Max(ctx context.Context, r *ms.MaxRequest) (*ms.MaxResponse, error) {
	log.Printf("Received input array %v", r.Array)
	return &ms.MaxResponse{Maximum: mf.Max(r.Array)}, nil
}

// Min service handler method
func (server *mathService) Min(ctx context.Context, r *ms.MinRequest) (*ms.MinResponse, error) {
	log.Printf("Received input array %v", r.Array)
	return &ms.MinResponse{Minimum: mf.Min(r.Array)}, nil
}

// Sort service handler method
func (server *mathService) Sort(ctx context.Context, r *ms.SortRequest) (*ms.SortResponse, error) {
	log.Printf("Received input array %v", r.Array)
	return &ms.SortResponse{SortedArray: mf.Bubblesort(r.Array)}, nil
}

// Sum service handler method
func (server *mathService) Sum(ctx context.Context, r *ms.SumRequest) (*ms.SumResponse, error) {
	log.Printf("Received input array %v", r.Array)
	return &ms.SumResponse{ArrayValuesSum: mf.Sum(r.Array)}, nil
}

func main() {

	// start a server listener
	listener, err := net.Listen("tcp", port)
	if err != nil {
		log.Fatal(errors.Wrap(err, fmt.Sprintf("Failed to listen on port %v", port)))
	}

	// instanstiate a gRPC server
	server := grpc.NewServer()
	ms.RegisterMathServiceServer(server, &mathService{})
	reflection.Register(server) // provides publicly accessible information about this `server`

	// start the gRPC Server
	fmt.Println("Microservice is up and running at http://127.0.0.1:9090")
	err = server.Serve(listener)
	if err != nil {
		log.Fatal(errors.Wrap(err, "Failed to start gRPC Server"))
	}
}
