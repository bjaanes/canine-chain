package lp_test

import (
	"testing"

	keepertest "canine-LP/testutil/keeper"
	"canine-LP/testutil/nullify"
	"canine-LP/x/lp"
	"canine-LP/x/lp/types"
	"github.com/stretchr/testify/require"
)

func TestGenesis(t *testing.T) {
	genesisState := types.GenesisState{
		Params: types.DefaultParams(),

		LPoolList: []types.LPool{
			{
				Index: "0",
			},
			{
				Index: "1",
			},
		},
		LProviderRecordList: []types.LProviderRecord{
			{
				Index: "0",
			},
			{
				Index: "1",
			},
		},
		// this line is used by starport scaffolding # genesis/test/state
	}

	k, ctx := keepertest.LpKeeper(t)
	lp.InitGenesis(ctx, *k, genesisState)
	got := lp.ExportGenesis(ctx, *k)
	require.NotNil(t, got)

	nullify.Fill(&genesisState)
	nullify.Fill(got)

	require.ElementsMatch(t, genesisState.LPoolList, got.LPoolList)
	require.ElementsMatch(t, genesisState.LProviderRecordList, got.LProviderRecordList)
	// this line is used by starport scaffolding # genesis/test/assert
}
