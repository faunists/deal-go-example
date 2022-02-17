package main_test

import (
	"context"
	"testing"

	"github.com/faunists/deal-go-example/protogen/proto/example"

	"github.com/faunists/deal-go-example/api/server"
	"google.golang.org/grpc"
)

func TestMyServiceContract(t *testing.T) {
	grpcServer := grpc.NewServer()
	example.RegisterMyServiceServer(grpcServer, &server.MyServer{})

	example.MyServiceContractTest(t, context.Background(), grpcServer)
}
