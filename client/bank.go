package client

import (
	"context"

	sdktypes "github.com/cosmos/cosmos-sdk/types"
	banktypes "github.com/cosmos/cosmos-sdk/x/bank/types"
)

// Balance returns balance of a given account for staking denom.
func (cp *ClientProxy) Balance(ctx context.Context, address string, denom string) (sdktypes.Coin, error) {
	client := banktypes.NewQueryClient(cp.grpcConn)

	req := banktypes.QueryBalanceRequest{
		Address: address,
		Denom:   denom,
	}

	resp, err := client.Balance(ctx, &req)
	if err != nil {
		return sdktypes.Coin{}, err
	}

	return *resp.Balance, nil
}

// Balances returns all balances of a given account.
func (cp *ClientProxy) Balances(ctx context.Context, address string) (sdktypes.Coins, error) {
	client := banktypes.NewQueryClient(cp.grpcConn)

	req := banktypes.QueryAllBalancesRequest{
		Address: address,
	}

	resp, err := client.AllBalances(ctx, &req)
	if err != nil {
		return sdktypes.Coins{}, err
	}

	return resp.Balances, nil
}
