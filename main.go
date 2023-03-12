package main

import (
	"context"
	"log"
	"net"
	"net/http"
	"runtime"

	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
	api "volumefi.com/volume_fi/api"
	pb "volumefi.com/volume_fi/pb"
)

func main() {
	log.Println("Entrypoint/Start for the exam question")
	nCPU := runtime.NumCPU()
	runtime.GOMAXPROCS(nCPU)

	// Required API
	http.HandleFunc("/", api.HandleDefault)
	http.HandleFunc("/calculate", api.HandleCalculateFlightTracker)

	// Start GRPC Servre
	go startGRPCServer()

	// Start the server on port 8080
	log.Fatal(http.ListenAndServe(":8080", nil))
}

type server struct {
	pb.FlightServer
}

func (s *server) Calculate(ctx context.Context, in *pb.CalculatePathRequest) (*pb.CalculatePathResponse, error) {
	return &pb.CalculatePathResponse{
		Response: &pb.Response{FlightPath: []string{"SFO", "ATL", "GSO", "IND", "EWR"}},
	}, nil
}

func startGRPCServer() {
	listener, err := net.Listen("tcp", ":8089")
	if err != nil {
		panic(err)
	}

	grpcServer := grpc.NewServer()
	reflection.Register(grpcServer)
	pb.RegisterFlightServer(grpcServer, &server{})
	if err := grpcServer.Serve(listener); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}
}
