package main

import (
	"log"
	"net"

	"github.com/faunists/deal-go-example/api/server"
	"github.com/faunists/deal-go-example/protogen/example"
	"google.golang.org/grpc"
)

func main() {
	listener, err := net.Listen("tcp", ":50051")
	if err != nil {
		log.Fatalf("error opening the listener: %v", err)
	}
	defer listener.Close()

	grpcServer := grpc.NewServer()
	example.RegisterMyServiceServer(grpcServer, &server.MyServer{})

	log.Printf("grpc server listening at %v", listener.Addr())
	if err = grpcServer.Serve(listener); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}
	grpcServer.GracefulStop()
}
