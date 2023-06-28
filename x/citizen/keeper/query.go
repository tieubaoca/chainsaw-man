package keeper

import (
	"gitlab.com/onechain/saw/x/citizen/types"
)

var _ types.QueryServer = Keeper{}
