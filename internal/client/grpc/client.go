package grpcclient

import lsuserspb "github.com/p1xray/love-signal-protos/gen/go/users"

// GRPCClient provides gRPC clients.
type GRPCClient struct {
	Users lsuserspb.UsersClient
}

// New creates new gRPC client instance.
func New(users lsuserspb.UsersClient) *GRPCClient {
	return &GRPCClient{
		Users: users,
	}
}
