package client

import (
	"context"

	"github.com/cosmos/cosmos-sdk/types/tx"
	sdktx "github.com/cosmos/cosmos-sdk/types/tx"
)

// BroadcastTx broadcasts transaction.
func (cp *ClientProxy) BroadcastTx(ctx context.Context, txBytes []byte, mode sdktx.BroadcastMode) (*tx.BroadcastTxResponse, error) {
	client := tx.NewServiceClient(cp.grpcConn)

	req := &tx.BroadcastTxRequest{
		TxBytes: txBytes,
		Mode:    mode,
	}

	return client.BroadcastTx(ctx, req)
}
