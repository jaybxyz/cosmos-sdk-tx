package main

import (
	"github.com/cosmos/cosmos-sdk/simapp"
	sdk "github.com/cosmos/cosmos-sdk/types"
	banktypes "github.com/cosmos/cosmos-sdk/x/bank/types"

	"github.com/kogisin/cosmos-sdk-tx/wallet"
)

var (
	rpcAddr  = "http://localhost:26657"
	grpcAddr = "localhost:9090"

	mnemonic = "guard cream sadness conduct invite crumble clock pudding hole grit liar hotel maid produce squeeze return argue turtle know drive eight casino maze host"
	password = ""

	gasLimit  = uint64(1_000_000)
	feeAmount = sdk.NewCoins(sdk.NewInt64Coin("stake", 5000))
	memo      = "By Kogisin cosmos-sdk-tx program"
)

func main() {

	// test
	_, _, _ = wallet.RecoverAccountFromMnemonic(mnemonic, password)
}

// sendTx
func sendTx() error {
	// choose your codec: Amino or Protobuf. Here, we use Protobuf.
	encCfg := simapp.MakeTestEncodingConfig()

	// create a new TxBuilder instance
	txBuilder := encCfg.TxConfig.NewTxBuilder()

	fromAddr, _ := sdk.AccAddressFromBech32("")
	toAddr, _ := sdk.AccAddressFromBech32("")
	amount := sdk.NewCoins()

	msg := banktypes.NewMsgSend(fromAddr, toAddr, amount)

	if err := txBuilder.SetMsgs(msg); err != nil {
		return err
	}
	txBuilder.SetGasLimit(gasLimit)
	txBuilder.SetFeeAmount(feeAmount)
	txBuilder.SetMemo(memo)

	return nil
}

// simulateTx
func simulateTx() error {
	return nil
}
