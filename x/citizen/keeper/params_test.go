package keeper_test

import (
	"testing"

	"github.com/stretchr/testify/require"
	testkeeper "gitlab.com/onechain/saw/testutil/keeper"
	"gitlab.com/onechain/saw/x/citizen/types"
)

func TestGetParams(t *testing.T) {
	k, ctx := testkeeper.CitizenKeeper(t)
	params := types.DefaultParams()

	k.SetParams(ctx, params)

	require.EqualValues(t, params, k.GetParams(ctx))
}
