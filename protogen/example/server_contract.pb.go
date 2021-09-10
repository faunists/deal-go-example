// Code generated by protoc-gen-go-deal. DO NOT EDIT.
//
// versions:
//   - protoc

package example

import (
	context "context"
	grpc "google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
	bufconn "google.golang.org/grpc/test/bufconn"
	proto "google.golang.org/protobuf/proto"
	log "log"
	net "net"
	testing "testing"
)

type MyServiceContractClient struct{}

func (_ MyServiceContractClient) MyMethod(ctx context.Context, in *RequestMessage, opts ...grpc.CallOption) (*ResponseMessage, error) {
	switch {
	case proto.Equal(in, &RequestMessage{RequestField: "VALUE"}):
		// Description: Should do something
		return &ResponseMessage{ResponseField: 42}, nil
	case proto.Equal(in, &RequestMessage{RequestField: "ANOTHER_VALUE"}):
		// Description: Some description here
		return nil, status.Errorf(codes.NotFound, "ANOTHER_VALUE NotFound")
	default:
		return nil, nil
	}
}

func MyServiceContractTest(t *testing.T, ctx context.Context, server *grpc.Server) {
	// gRPC Server setup
	bufSize := 1024 * 1024
	bufferListener := bufconn.Listen(bufSize)
	go func() {
		if err := server.Serve(bufferListener); err != nil {
			log.Fatalf("Contract Server test exited with error: %v", err)
		}
	}()
	defer server.Stop()

	// gRPC Client setup
	dialer := func(_ context.Context, _ string) (net.Conn, error) { return bufferListener.Dial() }
	clientConn, err := grpc.DialContext(ctx, "bufnet", grpc.WithContextDialer(dialer), grpc.WithInsecure())
	if err != nil {
		t.Fatalf("Failed to dial bufnet: %v", err)
	}
	defer clientConn.Close()

	client := NewMyServiceClient(clientConn)
	runMyServiceTests(t, ctx, client)
}

func runMyServiceTests(t *testing.T, ctx context.Context, client MyServiceClient) {
	t.Run("Contract test for 'MyMethod' method", func(t *testing.T) {
		t.Run("Success Cases", func(t *testing.T) {
			tests := []struct {
				name             string
				request          *RequestMessage
				expectedResponse *ResponseMessage
			}{
				{
					name:             "Should do something",
					request:          &RequestMessage{RequestField: "VALUE"},
					expectedResponse: &ResponseMessage{ResponseField: 42},
				},
			}

			for _, test := range tests {
				t.Run(test.name, func(t *testing.T) {
					response, err := client.MyMethod(ctx, test.request)
					if err != nil {
						t.Fatalf("unexpected error happened: %w", err)
					}

					if !proto.Equal(response, test.expectedResponse) {
						t.Fatalf(
							"expected response: %v, given response: %v",
							test.expectedResponse, response,
						)
					}
				})
			}
		})

		t.Run("Failure Cases", func(t *testing.T) {
			tests := []struct {
				name          string
				request       *RequestMessage
				expectedError string
			}{
				{
					name:          "Some description here",
					request:       &RequestMessage{RequestField: "ANOTHER_VALUE"},
					expectedError: "rpc error: code = NotFound desc = ANOTHER_VALUE NotFound",
				},
			}

			for _, test := range tests {
				t.Run(test.name, func(t *testing.T) {
					_, err := client.MyMethod(ctx, test.request)
					if err == nil {
						t.Fatalf("an error was expected but no one was returned")
					}

					if err.Error() != test.expectedError {
						t.Fatalf("expected error: %s, given error: %s", test.expectedError, err)
					}
				})
			}
		})
	})
}
