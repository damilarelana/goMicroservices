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

func main() {

	// start a server listener
	listener, err := net.Listen("tcp", port)
	if err != nil {
		log.Fatal(errors.Wrap(err, fmt.Sprintf("Failed to listen on port %v due to %v", port)))
	}

	// instanstiate a gRPC server
	server := grpc.NewServer()
	ms.RegisterMathServiceServer(server, &mathService{})
	reflection.Register(server) // provides publicly accessible information about this `server`

	// start the gRPC Server
	err = server.Serve(listener)
	if err != nil {
		log.Fatal(errors.Wrap(err, "Failed to start gRPC Server"))
	}
}
