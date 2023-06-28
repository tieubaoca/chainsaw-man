package types

// DONTCOVER

import (
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
)

// x/citizen module sentinel errors
var (
	ErrSample = sdkerrors.Register(ModuleName, 1100, "sample error")
)

var (
	ErrCitizenAlreadyExist = sdkerrors.Register(ModuleName, 1101, "citizen already exist")
	ErrCitizenNotFound     = sdkerrors.Register(ModuleName, 1102, "citizen not found")
)
