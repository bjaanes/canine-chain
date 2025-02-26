package types_test

import (
	"testing"

	"github.com/jackal-dao/canine/x/filetree/types"
	"github.com/stretchr/testify/require"
)

func TestGenesisState_Validate(t *testing.T) {
	for _, tc := range []struct {
		desc     string
		genState *types.GenesisState
		valid    bool
	}{
		{
			desc:     "default is valid",
			genState: types.DefaultGenesis(),
			valid:    true,
		},
		{
			desc: "valid genesis state",
			genState: &types.GenesisState{

				FilesList: []types.Files{
					{
						Address: "0",
					},
					{
						Address: "1",
					},
				},
				PubkeyList: []types.Pubkey{
					{
						Address: "0",
					},
					{
						Address: "1",
					},
				},
				Tracker: &types.Tracker{
					TrackingNumber: 51,
				},
				// this line is used by starport scaffolding # types/genesis/validField
			},
			valid: true,
		},
		{
			desc: "duplicated files",
			genState: &types.GenesisState{
				FilesList: []types.Files{
					{
						Address: "0",
					},
					{
						Address: "0",
					},
				},
			},
			valid: false,
		},
		{
			desc: "duplicated pubkey",
			genState: &types.GenesisState{
				PubkeyList: []types.Pubkey{
					{
						Address: "0",
					},
					{
						Address: "0",
					},
				},
			},
			valid: false,
		},
		// this line is used by starport scaffolding # types/genesis/testcase
	} {
		t.Run(tc.desc, func(t *testing.T) {
			err := tc.genState.Validate()
			if tc.valid {
				require.NoError(t, err)
			} else {
				require.Error(t, err)
			}
		})
	}
}
