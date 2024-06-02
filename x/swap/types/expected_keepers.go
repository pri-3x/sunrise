package types

import (
	"context"

	sdk "github.com/cosmos/cosmos-sdk/types"
	transfertypes "github.com/cosmos/ibc-go/v8/modules/apps/transfer/types"
	lptypes "github.com/sunriselayer/sunrise/x/liquiditypool/types"
)

// AccountKeeper defines the expected interface for the Account module.
type AccountKeeper interface {
	GetAccount(context.Context, sdk.AccAddress) sdk.AccountI // only used for simulation
	// Methods imported from account should be defined here
}

// BankKeeper defines the expected interface for the Bank module.
type BankKeeper interface {
	SpendableCoins(context.Context, sdk.AccAddress) sdk.Coins
	// Methods imported from bank should be defined here
}

// TransferKeeper defines the expected interface for the IBC Transfer module.
type TransferKeeper interface {
	Transfer(ctx context.Context, msg *transfertypes.MsgTransfer) (*transfertypes.MsgTransferResponse, error)
	DenomPathFromHash(ctx sdk.Context, denom string) (string, error)
	GetTotalEscrowForDenom(ctx sdk.Context, denom string) sdk.Coin
	SetTotalEscrowForDenom(ctx sdk.Context, coin sdk.Coin)
}

// LiquidityPoolKeeper defines the expected interface for the liquiditypool module.
type LiquidityPoolKeeper interface {
	GetPool(ctx context.Context, id uint64) (val lptypes.Pool, found bool)
	CalcOutAmtGivenIn(ctx sdk.Context, pool lptypes.Pool, tokenIn sdk.Coin, tokenOutDenom string, spreadFactor math.LegacyDec) (tokenOut sdk.Coin, err error)
	CalcInAmtGivenOut(ctx sdk.Context, pool lptypes.Pool, tokenOut sdk.Coin, tokenInDenom string, spreadFactor math.LegacyDec) (sdk.Coin, error)
	SwapExactAmountIn(ctx sdk.Context, sender sdk.AccAddress, pool lptypes.Pool, tokenIn sdk.Coin, tokenOutDenom string, tokenOutMinAmount math.Int, spreadFactor math.LegacyDec) (tokenOutAmount math.Int, err error)
	SwapExactAmountOut(ctx sdk.Context, sender sdk.AccAddress, pool lptypes.Pool, tokenInDenom string, tokenInMaxAmount math.Int, tokenOut sdk.Coin, spreadFactor math.LegacyDec) (tokenInAmount math.Int, err error)
}

// ParamSubspace defines the expected Subspace interface for parameters.
type ParamSubspace interface {
	Get(context.Context, []byte, interface{})
	Set(context.Context, []byte, interface{})
}
