package client

import (
	"context"

	authtypes "github.com/cosmos/cosmos-sdk/x/auth/types"
)

// BaseAccountInfo returns base account information
func (cp *ClientProxy) BaseAccountInfo(ctx context.Context, address string) (authtypes.BaseAccount, error) {
	client := authtypes.NewQueryClient(cp.grpcConn)

	req := authtypes.QueryAccountRequest{
		Address: address,
	}

	resp, err := client.Account(ctx, &req)
	if err != nil {
		return authtypes.BaseAccount{}, err
	}

	var acc authtypes.BaseAccount
	if err := acc.Unmarshal(resp.GetAccount().Value); err != nil {
		return authtypes.BaseAccount{}, err
	}

	return acc, nil
}
