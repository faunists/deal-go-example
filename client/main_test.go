package main_test

import (
	"context"
	"testing"

	"github.com/faunists/deal-go-example/protogen/proto/example"
	"github.com/stretchr/testify/require"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	"google.golang.org/protobuf/proto"
)

func TestClient(t *testing.T) {
	t.Run("should return a response", func(t *testing.T) {
		ctx := context.Background()
		expectedResp := &example.ResponseMessage{ResponseField: 42}
		client := example.MyServiceContractClient{}

		actualResp, err := client.MyMethod(ctx, &example.RequestMessage{
			RequestField: "VALUE",
		})

		require.NoError(t, err)
		require.True(t, proto.Equal(expectedResp, actualResp))
	})

	t.Run("should return an error", func(t *testing.T) {
		ctx := context.Background()
		expectedError := status.Error(codes.NotFound, "ANOTHER_VALUE NotFound")
		client := example.MyServiceContractClient{}

		_, err := client.MyMethod(ctx, &example.RequestMessage{
			RequestField: "ANOTHER_VALUE",
		})

		require.EqualError(t, err, expectedError.Error())
	})
}
