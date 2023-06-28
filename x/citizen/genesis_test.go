package citizen_test

import (
	"testing"

	"github.com/stretchr/testify/require"
	keepertest "gitlab.com/onechain/saw/testutil/keeper"
	"gitlab.com/onechain/saw/testutil/nullify"
	"gitlab.com/onechain/saw/x/citizen"
	"gitlab.com/onechain/saw/x/citizen/types"
)

func TestGenesis(t *testing.T) {
	genesisState := types.GenesisState{
		Params: types.DefaultParams(),

		// this line is used by starport scaffolding # genesis/test/state
	}

	k, ctx := keepertest.CitizenKeeper(t)
	citizen.InitGenesis(ctx, *k, genesisState)
	got := citizen.ExportGenesis(ctx, *k)
	require.NotNil(t, got)

	nullify.Fill(&genesisState)
	nullify.Fill(got)

	// this line is used by starport scaffolding # genesis/test/assert
}
