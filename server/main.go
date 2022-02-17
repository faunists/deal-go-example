package main

import (
	"log"
	"net"

	"github.com/faunists/deal-go-example/protogen/proto/example"
	"google.golang.org/grpc"
)

func main() {
	listener, err := net.Listen("tcp", ":50051")
	if err != nil {
		log.Fatalf("error opening the listener: %v", err)
	}
	defer func() { _ = listener.Close() }()

	grpcServer := grpc.NewServer()
	stubServer := example.MyServiceStubServer{}
	example.RegisterMyServiceServer(grpcServer, &stubServer)

	log.Printf("grpc server listening at %v", listener.Addr())
	if err = grpcServer.Serve(listener); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}
	grpcServer.GracefulStop()
}
