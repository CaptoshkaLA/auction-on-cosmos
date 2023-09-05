package keeper

import (
	"fmt"

	"github.com/CaptoshkaLA/AuctionOnCosmos/x/auctiononcosmos/types"
	sdk "github.com/cosmos/cosmos-sdk/types"
)

// CreateAuction creates a new auction
func (k Keeper) CreateAuctionAuxiliary(ctx sdk.Context, auction types.Auction) uint64 {

	if err := auction.Validate(); err != nil {
		panic(err)
	}

	// Getting the current auction index
	store := ctx.KVStore(k.storeKey)
	var auctionIndex uint64
	//bz := store.Get(types.KeyLastAuctionID)
	bz := store.Get([]byte(types.KeyLastAuctionID))
	if bz == nil {
		auctionIndex = 0
	} else {
		auctionIndex = sdk.BigEndianToUint64(bz)
	}
	// Setting a new index
	newAuctionIndex := auctionIndex + 1
	//store.Set(types.KeyLastAuctionID, sdk.Uint64ToBigEndian(newAuctionIndex))
	store.Set([]byte(types.KeyLastAuctionID), sdk.Uint64ToBigEndian(newAuctionIndex))

	// Saving the auction
	store.Set(types.KeyAuction(newAuctionIndex), k.cdc.MustMarshal(&auction))

	return newAuctionIndex
}

// GetAuction returns an auction by index
func (k Keeper) GetAuctionAuxiliary(ctx sdk.Context, index uint64) (types.Auction, bool) {
	store := ctx.KVStore(k.storeKey)
	bz := store.Get(types.KeyAuction(index))
	if bz == nil {
		return types.Auction{}, false
	}
	var auction types.Auction
	k.cdc.MustUnmarshal(bz, &auction)
	return auction, true
}

// EndAuction completes the auction
func (k Keeper) EndAuctionAuxiliary(ctx sdk.Context, index uint64) error {
	store := ctx.KVStore(k.storeKey)
	if store.Has(types.KeyAuction(index)) {
		store.Delete(types.KeyAuction(index))
		return nil
	}
	return fmt.Errorf("Auction with index %d not found", index)
}

// TODO: Implement a function to cancel the auction
