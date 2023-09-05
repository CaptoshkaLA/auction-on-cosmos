package keeper

import (
	"context"
	"fmt"
	"time"

	"github.com/CaptoshkaLA/AuctionOnCosmos/x/auctiononcosmos/types"
	sdk "github.com/cosmos/cosmos-sdk/types"
)

func (k msgServer) CreateBid(goCtx context.Context, msg *types.MsgCreateBid) (*types.MsgCreateBidResponse, error) {
	ctx := sdk.UnwrapSDKContext(goCtx)

	// Receive the auction by ID and check its status
	auction, found := k.GetAuctionAuxiliary(ctx, msg.AuctionId)
	if !found {
		return nil, fmt.Errorf("auction with ID %d not found", msg.AuctionId)
	}

	// Get the current highest bid for the auction
	highestBid, _ := k.GetCurrentHighestBidForAuction(ctx, msg.AuctionId)

	// Creating a new bid
	bid := k.CreateBidAuxiliary(
		ctx,
		msg.Index,
		msg.AuctionId,
		msg.Bidder,
		msg.BidAmount,
		uint64(time.Now().Unix()), //msg.CreatedAt,
	)

	// Validation of a new bid
	if err := types.ValidateBid(bid, highestBid, auction); err != nil {
		return nil, err
	}

	return &types.MsgCreateBidResponse{
		Index: bid.Index,
	}, nil
}
