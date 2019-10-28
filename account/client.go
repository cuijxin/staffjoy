package account

import (
	fmt "fmt"

	grpc "google.golang.org/grpc"
)

// After calling it, run a defer close on the close function
func NewClient() (AccountServiceClient, func() error, error) {
	conn, err := grpc.Dial(Endpoint, grpc.WithInsecure())
	if err != nil {
		return nil, nil, fmt.Errorf("did not connect: %v", err)
	}

	return NewAccountServiceClient(conn), conn.Close, nil
}
