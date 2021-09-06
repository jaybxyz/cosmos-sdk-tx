package client

import (
	"net/http"

	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

type AddTokenRoundTripper struct {
	rt    http.RoundTripper
	token string
}

func (rt AddTokenRoundTripper) RoundTrip(req *http.Request) (*http.Response, error) {
	req.Header.Add("Authorization", rt.token)
	return rt.rt.RoundTrip(req)
}

type ClientOptions struct {
	blockHeight *int64
}

type ClientOption func(*ClientOptions)

func WithBlockHeight(blockHeight int64) ClientOption {
	return func(opts *ClientOptions) {
		opts.blockHeight = &blockHeight
	}
}

// IsNotFound returns not found status.
func IsNotFound(err error) bool {
	return status.Convert(err).Code() == codes.NotFound
}
