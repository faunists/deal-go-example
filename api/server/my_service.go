package server

import (
	"context"

	"github.com/faunists/deal-go-example/protogen/proto/example"

	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

type MyServer struct {
	example.UnimplementedMyServiceServer
}

func (svc *MyServer) MyMethod(
	_ context.Context,
	req *example.RequestMessage,
) (*example.ResponseMessage, error) {
	if req.RequestField == "ANOTHER_VALUE" {
		return nil, status.Error(codes.NotFound, "ANOTHER_VALUE NotFound")
	}

	return &example.ResponseMessage{ResponseField: 42}, nil
}
