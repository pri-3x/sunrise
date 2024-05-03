package keeper

import (
	"context"

	errorsmod "cosmossdk.io/errors"
	"cosmossdk.io/math"
	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
	"github.com/sunriselayer/sunrise/x/liquiditypool/types"
)

func (k Keeper) TransferFromAccountToPoolModule(ctx context.Context, token sdk.Coin, address sdk.AccAddress, poolId uint64) error {
	if token.Amount.IsZero() {
		return nil
	}

	moduleName := types.PoolModuleName(poolId)
	err := k.bankKeeper.SendCoinsFromAccountToModule(ctx, address, moduleName, sdk.NewCoins(token))

	return err
}

func (k Keeper) TransferFromPoolModuleToAccount(ctx context.Context, token sdk.Coin, address sdk.AccAddress, poolId uint64) error {
	if token.Amount.IsZero() {
		return nil
	}

	moduleName := types.PoolModuleName(poolId)
	err := k.bankKeeper.SendCoinsFromModuleToAccount(ctx, moduleName, address, sdk.NewCoins(token))

	return err
}

func (k Keeper) TransferFromPoolModuleToPoolTreasuryModule(ctx context.Context, token sdk.Coin, poolId uint64) error {
	if token.Amount.IsZero() {
		return nil
	}

	moduleName := types.PoolModuleName(poolId)
	treasuryModuleName := types.PoolTreasuryModuleName(poolId)
	err := k.bankKeeper.SendCoinsFromModuleToModule(ctx, moduleName, treasuryModuleName, sdk.NewCoins(token))

	return err
}

func (k Keeper) EmitEventPoolFee(ctx context.Context, poolId uint64, poolFee sdk.Coin) {
	if poolFee.Amount.IsZero() {
		return
	}

	sdkCtx := sdk.UnwrapSDKContext(ctx)
	sdkCtx.EventManager().EmitTypedEvent(nil)
}

// ExactAmountIn

func (k Keeper) SwapExactAmountInSinglePool(ctx context.Context, poolId uint64, tokenIn sdk.Coin, denomOutConfirmation string, dryRun bool, address *sdk.AccAddress) (*math.Int, error) {
	// equal to zero is not caught here
	if tokenIn.Amount.LT(math.ZeroInt()) {
		return nil, errorsmod.Wrapf(sdkerrors.ErrInvalidRequest, "invalid tokenIn amount %s", tokenIn.Amount.String())
	}

	// get pool
	pool, found := k.GetPool(ctx, poolId)

	if !found {
		return nil, errorsmod.Wrapf(sdkerrors.ErrKeyNotFound, "pool id %d not found", poolId)
	}

	// check denom
	if tokenIn.Denom != pool.BaseDenom && tokenIn.Denom != pool.QuoteDenom {
		return nil, errorsmod.Wrapf(sdkerrors.ErrInvalidRequest, "invalid tokenIn denom %s", tokenIn.Denom)
	}

	// check denom
	var denomOut string
	if tokenIn.Denom == pool.BaseDenom {
		denomOut = pool.QuoteDenom
	} else {
		denomOut = pool.BaseDenom
	}

	if denomOut != denomOutConfirmation {
		return nil, errorsmod.Wrapf(sdkerrors.ErrInvalidRequest, "invalid denomOutConfirmation %s", denomOutConfirmation)
	}

	// pass zero input
	if tokenIn.Amount.Equal(math.ZeroInt()) {
		amountOut := math.ZeroInt()
		return &amountOut, nil
	}

	// calculate amount
	x, y := k.GetPoolBalance(ctx, pool)
	kValue, err := types.CalculateK(x, y, pool.FK)
	if err != nil {
		return nil, errorsmod.Wrapf(sdkerrors.ErrInvalidRequest, "error calculating k: %s", err.Error())
	}

	var amountOutNeg *math.Int
	if tokenIn.Denom == pool.BaseDenom {
		amountOutNeg, err = types.CalculateDy(x, tokenIn.Amount, *kValue, pool.FY)
	} else {
		amountOutNeg, err = types.CalculateDx(y, tokenIn.Amount, *kValue, pool.FX)
	}

	if err != nil {
		return nil, errorsmod.Wrapf(sdkerrors.ErrInvalidRequest, "error calculating amountOut: %s", err.Error())
	}

	// equal to zero is not caught here
	amountOut := amountOutNeg.Neg()
	if amountOut.LT(math.ZeroInt()) {
		return nil, errorsmod.Wrapf(sdkerrors.ErrInvalidRequest, "amountOut is negative")
	}

	// calculate fees
	treasuryTaxAmount := k.GetParams(ctx).TreasuryTaxRate.MulInt(amountOut).RoundInt()
	poolFeeAmount := pool.FeeRate.MulInt(amountOut).RoundInt()
	amountOut = amountOut.Sub(treasuryTaxAmount).Sub(poolFeeAmount)

	if !dryRun {
		// transfer tokenIn from address to module
		if err := k.TransferFromAccountToPoolModule(ctx, tokenIn, *address, poolId); err != nil {
			return nil, errorsmod.Wrapf(sdkerrors.ErrInsufficientFunds, "error transferring tokenIn: %s", err.Error())
		}

		// transfer tokenOut from module to address
		tokenOut := sdk.NewCoin(denomOut, amountOut)
		if err := k.TransferFromPoolModuleToAccount(ctx, tokenOut, *address, poolId); err != nil {
			return nil, err
		}

		// transfer treasury tax to treasury
		treasuryTax := sdk.NewCoin(denomOut, treasuryTaxAmount)
		if err := k.TransferFromPoolModuleToPoolTreasuryModule(ctx, treasuryTax, poolId); err != nil {
			return nil, err
		}

		// emit event of pool fee
		poolFee := sdk.NewCoin(tokenOut.Denom, poolFeeAmount)
		k.EmitEventPoolFee(ctx, poolId, poolFee)
	}

	return &amountOut, nil
}

func (k Keeper) SwapExactAmountInMultiPool(ctx context.Context, route types.SwapRoute, tokenIn sdk.Coin, dryRun bool, address *sdk.AccAddress) (*sdk.Coin, error) {
	amounts := types.AmountsFromWeights(route.PoolWeights, tokenIn.Amount)

	var tokenOut sdk.Coin
	if tokenIn.Denom == route.BaseDenom {
		tokenOut = sdk.NewCoin(route.QuoteDenom, math.ZeroInt())
	} else {
		tokenOut = sdk.NewCoin(route.BaseDenom, math.ZeroInt())
	}

	for i := range route.PoolWeights {
		amountOut, err := k.SwapExactAmountInSinglePool(
			ctx,
			route.PoolWeights[i].PoolId,
			sdk.NewCoin(tokenIn.Denom, amounts[i]),
			tokenOut.Denom,
			dryRun,
			address,
		)

		if err != nil {
			return nil, errorsmod.Wrapf(sdkerrors.ErrInvalidRequest, "error swapping in pool: %s", err.Error())
		}

		tokenOut.Amount = tokenOut.Amount.Add(*amountOut)
	}

	return &tokenOut, nil
}

func (k Keeper) SwapExactAmountInMultiRoute(ctx context.Context, routes []types.SwapRoute, tokenIn sdk.Coin, dryRun bool, address *sdk.AccAddress) ([]sdk.Coin, error) {
	tokensOut := []sdk.Coin{}
	for _, route := range routes {
		tokenOut, err := k.SwapExactAmountInMultiPool(ctx, route, tokenIn, dryRun, address)
		if err != nil {
			return nil, errorsmod.Wrapf(sdkerrors.ErrInvalidRequest, "error swapping: %s", err.Error())
		}

		tokensOut = append(tokensOut, *tokenOut)

		tokenIn = *tokenOut
	}

	return tokensOut, nil
}

// ExactAmountOut

func (k Keeper) SwapExactAmountOutSinglePool(ctx context.Context, poolId uint64, tokenOut sdk.Coin, denomIntConfirmation string, dryRun bool, address *sdk.AccAddress) (*math.Int, error) {
	// equal to zero is caught here
	if tokenOut.Amount.LTE(math.ZeroInt()) {
		return nil, errorsmod.Wrapf(sdkerrors.ErrInvalidRequest, "invalid tokenOut amount %s", tokenOut.Amount.String())
	}

	// get pool
	pool, found := k.GetPool(ctx, poolId)

	if !found {
		return nil, errorsmod.Wrapf(sdkerrors.ErrKeyNotFound, "pool id %d not found", poolId)
	}

	// check denom
	if tokenOut.Denom != pool.BaseDenom && tokenOut.Denom != pool.QuoteDenom {
		return nil, errorsmod.Wrapf(sdkerrors.ErrInvalidRequest, "invalid tokenOut denom %s", tokenOut.Denom)
	}

	// check denom
	var denomIn string
	if tokenOut.Denom == pool.BaseDenom {
		denomIn = pool.QuoteDenom
	} else {
		denomIn = pool.BaseDenom
	}

	if denomIn != denomIntConfirmation {
		return nil, errorsmod.Wrapf(sdkerrors.ErrInvalidRequest, "invalid denomIntConfirmation %s", denomIntConfirmation)
	}

	// calculate amount
	treasuryTaxRate := k.GetParams(ctx).TreasuryTaxRate
	tokenOutBeforeFee := math.LegacyOneDec().Quo(math.LegacyOneDec().Sub(treasuryTaxRate).Sub(pool.FeeRate)).MulInt(tokenOut.Amount).RoundInt()
	if tokenOutBeforeFee.LTE(math.ZeroInt()) {
		return nil, errorsmod.Wrapf(sdkerrors.ErrInvalidRequest, "tokenOutBeforeFee is negative")
	}
	tokenOutBeforeFeeNeg := tokenOutBeforeFee.Neg()

	x, y := k.GetPoolBalance(ctx, pool)
	kValue, err := types.CalculateK(x, y, pool.FK)
	if err != nil {
		return nil, errorsmod.Wrapf(sdkerrors.ErrInvalidRequest, "error calculating k: %s", err.Error())
	}

	var amountIn *math.Int
	if tokenOut.Denom == pool.BaseDenom {
		amountIn, err = types.CalculateDy(x, tokenOutBeforeFeeNeg, *kValue, pool.FY)
	} else {
		amountIn, err = types.CalculateDx(y, tokenOutBeforeFeeNeg, *kValue, pool.FX)
	}

	if err != nil {
		return nil, errorsmod.Wrapf(sdkerrors.ErrInvalidRequest, "error calculating amountOut: %s", err.Error())
	}

	// equal to zero is caught here
	if amountIn.LTE(math.ZeroInt()) {
		return nil, errorsmod.Wrapf(sdkerrors.ErrInvalidRequest, "amountIn is negative")
	}

	// calculate fees
	treasuryTaxAmount := treasuryTaxRate.MulInt(tokenOutBeforeFee).RoundInt()
	poolFeeAmount := tokenOutBeforeFee.Sub(tokenOut.Amount).Sub(treasuryTaxAmount)

	if !dryRun {
		// transfer tokenIn from address to module
		tokenIn := sdk.NewCoin(denomIn, *amountIn)
		if err := k.TransferFromAccountToPoolModule(ctx, tokenIn, *address, poolId); err != nil {
			return nil, err
		}

		// transfer tokenOut from module to address
		// no need to check zero case
		if err := k.TransferFromPoolModuleToAccount(ctx, tokenOut, *address, poolId); err != nil {
			return nil, err
		}

		// transfer treasury tax to treasury
		treasuryTax := sdk.NewCoin(tokenOut.Denom, treasuryTaxAmount)
		if err := k.TransferFromPoolModuleToPoolTreasuryModule(ctx, treasuryTax, poolId); err != nil {
			return nil, err
		}

		// emit event of pool fee
		poolFee := sdk.NewCoin(tokenOut.Denom, poolFeeAmount)
		k.EmitEventPoolFee(ctx, poolId, poolFee)
	}

	return amountIn, nil
}

func (k Keeper) SwapExactAmountOutMultiPool(ctx context.Context, route types.SwapRoute, tokenOut sdk.Coin, dryRun bool, address *sdk.AccAddress, maxAmountIn *math.Int) (*sdk.Coin, error) {
	amounts := types.AmountsFromWeights(route.PoolWeights, tokenOut.Amount)

	var tokenIn sdk.Coin
	if tokenOut.Denom == route.BaseDenom {
		tokenIn = sdk.NewCoin(route.QuoteDenom, math.ZeroInt())
	} else {
		tokenIn = sdk.NewCoin(route.BaseDenom, math.ZeroInt())
	}

	for i := range route.PoolWeights {
		amountIn, err := k.SwapExactAmountOutSinglePool(
			ctx,
			route.PoolWeights[i].PoolId,
			sdk.NewCoin(tokenOut.Denom, amounts[i]),
			tokenIn.Denom,
			dryRun,
			address,
		)

		if err != nil {
			return nil, errorsmod.Wrapf(sdkerrors.ErrInvalidRequest, "error swapping in pool: %s", err.Error())
		}

		tokenIn.Amount = tokenIn.Amount.Add(*amountIn)

		if !dryRun {
			if tokenIn.Amount.GT(*maxAmountIn) {
				return nil, errorsmod.Wrap(sdkerrors.ErrInvalidRequest, "slippage exceeded")
			}
		}
	}

	return &tokenIn, nil
}
