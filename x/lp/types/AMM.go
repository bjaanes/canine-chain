package types

import (
	"errors"
	"fmt"

	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
)

// Every liquidity pools have an amm model that balances the assets.
// All amm models implements AMM interface.

// During a swap transaction or query, it calls this function to get their amm model.
// Returns an amm model with an interface to estimate swap in and out.
// When an invalid id is given backup amm model is returned with an error.
func GetAMM(id uint32) (AMM, error) {
	switch id {
	// NOTE: append new amm models here, do not change existing amm id
	// e.g.
	// case x:
	// 	return amm_name{}, nil
	case 0:
		return backup{}, nil
	default:
		return backup{}, errors.New(fmt.Sprintf("AMM model id not found: %v", id))
	}
}

// All amm models should implement this interface.
type AMM interface {
	// Estimate amount of coins returned with given deposit and pool coins.
	// Returns amount of coin returned.
	// Returns error when deposit coins are not subset of pool coins or
	// 	return amount less than or equal to zero.
	EstimateReturn(poolCoins sdk.Coins, depositCoins sdk.Coins) (sdk.Coins, error)

	// Estimate amount of coins to deposit to get desired coin return.
	// Takes desired amount of coins (sdk.Coins).
	// Returns amount of coins to deposit to get desired coins.
	EstimateDeposit(poolCoins sdk.Coins, desiredCoins sdk.Coins) (sdk.Coins, error)
}

// *******AMMs********

// Supports pool with two denoms.
// Model: xy=k
type backup struct{}

func (b backup) EstimateReturn(poolCoins sdk.Coins, depositCoins sdk.Coins) (sdk.Coins, error) {
	// How to get poolCoins denoms NOT in depositCoins denoms:
	// Subtract coins with deposit denom and amount in poolCoins from poolCoins
	depositDenom := depositCoins.GetDenomByIndex(0)
	subCoin := sdk.NewCoin(depositDenom, poolCoins.AmountOf(depositDenom))
	denomsGive := poolCoins.Sub(sdk.NewCoins(subCoin))

	x := depositDenom
	y := denomsGive.GetDenomByIndex(0)
	k := poolCoins.AmountOf(x).Mul(poolCoins.AmountOf(y))
	xAfter := poolCoins.AmountOf(x).Add(depositCoins.AmountOf(x))
	yAfter := k.Quo(xAfter)
	yAmt := poolCoins.AmountOf(y).Sub(yAfter)
	if yAmt == poolCoins.AmountOf(y) {
		yAmt = yAmt.Sub(sdk.OneInt())
	}
	if yAmt.LTE(sdk.ZeroInt()) {
		return sdk.NewCoins(sdk.NewCoin(y, sdk.ZeroInt())), sdkerrors.Wrapf(
			sdkerrors.ErrLogic,
			"AMM run result is zero",
		)
	}
	return sdk.NewCoins(sdk.NewCoin(y, yAmt)), nil
}

func (b backup) EstimateDeposit(poolCoins sdk.Coins, desiredCoins sdk.Coins) (sdk.Coins, error) {
	desiredDenom := desiredCoins.GetDenomByIndex(0)
	subCoin := sdk.NewCoin(desiredDenom, poolCoins.AmountOf(desiredDenom))
	depositEstDenom := poolCoins.Sub(sdk.NewCoins(subCoin)).GetDenomByIndex(0)

	x := poolCoins.AmountOf(depositEstDenom)
	var xEst sdk.Int
	y := poolCoins.AmountOf(desiredDenom)
	yDep := desiredCoins.AmountOf(desiredDenom)
	k := x.Mul(y)

	// k = (x + xEst)(y - yDep)
	yAfter := y.Sub(yDep)
	// k/yAfter = (x + xEst)
	z := k.Quo(yAfter)
	// z - x = xEst
	xEst = z.Sub(x)

	return sdk.NewCoins(sdk.NewCoin(desiredDenom, xEst)), nil
}
