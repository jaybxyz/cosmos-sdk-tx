package client

import (
	"context"
	"fmt"
	"net/http"
	"time"

	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials"

	rpcclient "github.com/tendermint/tendermint/rpc/client"
	rpc "github.com/tendermint/tendermint/rpc/client/http"
	tmctypes "github.com/tendermint/tendermint/rpc/core/types"
)

// ClientProxy is a wrapper for various clients.
type ClientProxy struct {
	rpcClient rpcclient.Client
	grpcConn  *grpc.ClientConn
}

// NewClientProxy creates a new Client with the given configuration.
func NewClientProxy(rpcURL, rpcToken, grpcURL, grpcToken string, insecure bool) (*ClientProxy, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	opts := []grpc.DialOption{grpc.WithBlock()}

	// handle TLS/SSL connection
	if insecure {
		opts = append(opts, grpc.WithInsecure())
	} else {
		opts = append(opts, grpc.WithTransportCredentials(credentials.NewTLS(nil)))
	}

	grpcConn, err := grpc.DialContext(ctx, grpcURL, opts...)
	if err != nil {
		return nil, fmt.Errorf("failed to create gRPC client connection : %w", err)
	}

	httpClient := &http.Client{
		CheckRedirect: nil,
		Jar:           nil,
		Timeout:       0,
	}

	if rpcToken != "" {
		httpClient.Transport = AddTokenRoundTripper{
			rt:    http.DefaultTransport,
			token: rpcToken,
		}
	}

	rpcClient, err := rpc.NewWithClient(rpcURL, "/websocket", httpClient)
	if err != nil {
		return nil, fmt.Errorf("failed to create RPC client connection: %w", err)
	}

	return &ClientProxy{
		rpcClient: rpcClient,
		grpcConn:  grpcConn,
	}, nil
}

// Close sloses the node stop execution to the RPC and GRPC clients.
func (cp *ClientProxy) Close() error {
	if err := cp.rpcClient.Stop(); err != nil {
		return err
	}
	if err := cp.grpcConn.Close(); err != nil {
		return err
	}
	return nil
}

// NetworkChainID returns network chain id.
func (cp *ClientProxy) NetworkChainID(ctx context.Context) (string, error) {
	status, err := cp.rpcClient.Status(ctx)
	if err != nil {
		return "", fmt.Errorf("failed to get status: %v", err)
	}

	return status.NodeInfo.Network, nil
}

// Status returns the status of the blockchain network.
func (cp *ClientProxy) Status(ctx context.Context) (*tmctypes.ResultStatus, error) {
	return cp.rpcClient.Status(ctx)
}
