package keeper_test

import (
	"testing"

	testkeeper "github.com/jackal-dao/canine/testutil/keeper"
	"github.com/jackal-dao/canine/x/jklmint/types"
	"github.com/stretchr/testify/require"
)

func TestGetParams(t *testing.T) {
	k, ctx := testkeeper.JklmintKeeper(t)
	params := types.DefaultParams()

	k.SetParams(ctx, params)

	require.EqualValues(t, params, k.GetParams(ctx))
	require.EqualValues(t, params.MintDenom, k.MintDenom(ctx))
}
