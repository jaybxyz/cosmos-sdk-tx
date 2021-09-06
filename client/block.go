package client

import (
	"context"
	"time"

	tmctypes "github.com/tendermint/tendermint/rpc/core/types"
)

// Block returns block information for the height.
func (cp *ClientProxy) Block(ctx context.Context, height int64) (*tmctypes.ResultBlock, error) {
	return cp.rpcClient.Block(ctx, &height)
}

// LatestBlockHeight returns the latest block height on the network.
func (cp *ClientProxy) LatestBlockHeight(ctx context.Context) (int64, error) {
	resp, err := cp.rpcClient.Status(ctx)
	if err != nil {
		return 0, err
	}

	return resp.SyncInfo.LatestBlockHeight, nil
}

// SubscribeNewBlocks subscribes to the new block event handler through the RPC
// client with the given subscriber name. An receiving only channel, context
// cancel function and an error is returned. It is up to the caller to cancel
// the context and handle any errors appropriately.
func (cp *ClientProxy) SubscribeNewBlocks(subscriber string) (<-chan tmctypes.ResultEvent, context.CancelFunc, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	eventCh, err := cp.rpcClient.Subscribe(ctx, subscriber, "tm.event = 'NewBlock'")
	return eventCh, cancel, err
}
