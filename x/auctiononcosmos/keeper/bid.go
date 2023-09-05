package keeper

import (
	//"fmt"

	"github.com/CaptoshkaLA/AuctionOnCosmos/x/auctiononcosmos/types"
	"github.com/cosmos/cosmos-sdk/store/prefix"
	sdk "github.com/cosmos/cosmos-sdk/types"
)

// CreateBid creates a new bid
func (k Keeper) CreateBidAuxiliary(ctx sdk.Context, index uint64, auctionID uint64, bidder string, bidAmount uint64, createdAt uint64) types.Bid {
	bid := types.Bid{
		Index:     index,
		AuctionID: auctionID,
		Bidder:    bidder,
		BidAmount: bidAmount,
		CreatedAt: createdAt,
	}

	store := prefix.NewStore(ctx.KVStore(k.storeKey), types.KeyPrefix(types.BidKey))
	key := types.KeyBid(index)
	value := k.cdc.MustMarshal(&bid)
	store.Set(key, value)

	return bid
}

// GetBid returns the bid based on the index
func (k Keeper) GetBid(ctx sdk.Context, index uint64) (types.Bid, bool) {
	store := prefix.NewStore(ctx.KVStore(k.storeKey), types.KeyPrefix(types.BidKey))
	key := types.KeyBid(index)
	value := store.Get(key)

	if value == nil {
		return types.Bid{}, false
	}

	bid := types.Bid{}
	k.cdc.MustUnmarshal(value, &bid)
	return bid, true
}

// GetAllBidsByAuction returns all bids for the specified auction
func (k Keeper) GetAllBidsByAuction(ctx sdk.Context, auctionID uint64) []types.Bid {
	var bids []types.Bid
	store := prefix.NewStore(ctx.KVStore(k.storeKey), types.KeyPrefix(types.BidKey))
	iterator := sdk.KVStorePrefixIterator(store, nil)

	defer iterator.Close()
	for ; iterator.Valid(); iterator.Next() {
		var bid types.Bid
		k.cdc.MustUnmarshal(iterator.Value(), &bid)
		if bid.AuctionID == auctionID {
			bids = append(bids, bid)
		}
	}

	return bids
}

// GetCurrentHighestBidForAuction returns the current highest bid for the specified auction
func (k Keeper) GetCurrentHighestBidForAuction(ctx sdk.Context, auctionID uint64) (types.Bid, bool) {
	bids := k.GetAllBidsByAuction(ctx, auctionID)
	if len(bids) == 0 {
		return types.Bid{}, false
	}

	highestBid := bids[0]
	for _, bid := range bids {
		if bid.BidAmount > highestBid.BidAmount {
			highestBid = bid
		}
	}

	return highestBid, true
}
