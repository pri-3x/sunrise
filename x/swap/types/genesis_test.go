package types_test

import (
	"testing"

	"github.com/stretchr/testify/require"
	"github.com/sunriselayer/sunrise/x/swap/types"
)

func TestGenesisState_Validate(t *testing.T) {
	tests := []struct {
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

				InFlightPacketList: []types.InFlightPacket{
					{
						SrcPortId:    "0",
						SrcChannelId: "0",
						Sequence:     0,
					},
					{
						SrcPortId:    "1",
						SrcChannelId: "1",
						Sequence:     1,
					},
				},
				// this line is used by starport scaffolding # types/genesis/validField
			},
			valid: true,
		},
		{
			desc: "duplicated inFlightPacket",
			genState: &types.GenesisState{
				InFlightPacketList: []types.InFlightPacket{
					{
						SrcPortId:    "0",
						SrcChannelId: "0",
						Sequence:     0,
					},
					{
						SrcPortId:    "0",
						SrcChannelId: "0",
						Sequence:     0,
					},
				},
			},
			valid: false,
		},
		// this line is used by starport scaffolding # types/genesis/testcase
	}
	for _, tc := range tests {
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