package keeper

import (
	"crypto/sha256"
	"encoding/hex"

	"github.com/cosmos/cosmos-sdk/store/prefix"
	sdk "github.com/cosmos/cosmos-sdk/types"
	"gitlab.com/onechain/saw/x/citizen/types"
)

func (k Keeper) CreateNewCitizen(ctx sdk.Context, addressOwner, name, avatar string) error {
	hasher := sha256.New()
	hasher.Write(append(ctx.TxBytes(), []byte(addressOwner)...))
	saIdByte := hasher.Sum(nil)[:12]
	saId := hex.EncodeToString(saIdByte)
	if _, ok := k.GetCitizen(ctx, saId); ok {
		return types.ErrCitizenAlreadyExist
	}
	citizen := types.Citizen{
		SaId:         saId,
		CitizenCode:  saId,
		Name:         name,
		AddressOwner: addressOwner,
		CreateAt:     uint64(ctx.BlockTime().Second()),
		UpdateAt:     uint64(ctx.BlockTime().Second()),
		AvatarUrl:    avatar,
	}
	store := prefix.NewStore(ctx.KVStore(k.storeKey), types.KeyPrefix(types.CitizenKeyPrefix))
	b := k.cdc.MustMarshal(&citizen)
	store.Set(saIdByte, b)
	return nil
}

func (k Keeper) GetCitizen(ctx sdk.Context, saId string) (types.Citizen, bool) {
	store := prefix.NewStore(ctx.KVStore(k.storeKey), types.KeyPrefix(types.CitizenKeyPrefix))
	var citizen types.Citizen
	b := store.Get([]byte(saId))
	if b == nil {
		return citizen, false
	}
	k.cdc.MustUnmarshal(b, &citizen)
	return citizen, true
}

func (k Keeper) RemoveCitizen(ctx sdk.Context, saId string) error {
	if _, ok := k.GetCitizen(ctx, saId); !ok {
		return types.ErrCitizenNotFound
	}
	store := prefix.NewStore(ctx.KVStore(k.storeKey), types.KeyPrefix(types.CitizenKeyPrefix))
	store.Delete([]byte(saId))
	return nil
}

func (k Keeper) UpdateCitizen(ctx sdk.Context, saId string, citizen types.Citizen) error {
	if _, ok := k.GetCitizen(ctx, saId); !ok {
		return types.ErrCitizenNotFound
	}
	store := prefix.NewStore(ctx.KVStore(k.storeKey), types.KeyPrefix(types.CitizenKeyPrefix))
	b := k.cdc.MustMarshal(&citizen)
	store.Set([]byte(saId), b)
	return nil
}
