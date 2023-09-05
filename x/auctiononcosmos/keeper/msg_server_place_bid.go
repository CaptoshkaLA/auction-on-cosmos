package keeper

import (
	"context"
	"fmt"

	"github.com/CaptoshkaLA/AuctionOnCosmos/x/auctiononcosmos/types"
	sdk "github.com/cosmos/cosmos-sdk/types"
)

func (k msgServer) CreateBid(goCtx context.Context, msg *types.MsgCreateBid) (*types.MsgCreateBidResponse, error) {
	ctx := sdk.UnwrapSDKContext(goCtx)

	// Получаем аукцион по ID и проверяем его статус
	auction, found := k.GetAuctionAuxiliary(ctx, msg.AuctionId)
	if !found {
		return nil, fmt.Errorf("auction with ID %d not found", msg.AuctionId)
	}

	if auction.AuctionStatus != "active" {
		return nil, fmt.Errorf("auction is not active")
	}

	// Создаем новую ставку
	bid := k.CreateBidAuxiliary(
		ctx,
		msg.Index,
		msg.AuctionId,
		msg.Bidder,
		msg.BidAmount,
		msg.CreatedAt,
	)

	return &types.MsgCreateBidResponse{
		Index: bid.Index,
	}, nil
}
