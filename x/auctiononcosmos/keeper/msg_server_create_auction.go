package keeper

import (
	"context"
	//"fmt"
	"time"

	"github.com/CaptoshkaLA/AuctionOnCosmos/x/auctiononcosmos/types"
	sdk "github.com/cosmos/cosmos-sdk/types"
)

func (k msgServer) CreateAuction(goCtx context.Context, msg *types.MsgCreateAuction) (*types.MsgCreateAuctionResponse, error) {
	ctx := sdk.UnwrapSDKContext(goCtx)

	auction := types.Auction{
		Index:           0, // Later we will update the index
		AuctionOwner:    msg.AuctionOwner,
		ItemDescription: msg.ItemDescription,
		InitialPrice:    msg.InitialPrice,
		CreatedAt:       uint64(time.Now().Unix()), //msg.CreatedAt,
		EndTime:         msg.EndTime,
		AuctionStatus:   msg.AuctionStatus,
	}

	// Validate the auction before creating
	if err := auction.Validate(); err != nil {
		panic(err)
	}

	index := k.CreateAuctionAuxiliary(ctx, auction)

	return &types.MsgCreateAuctionResponse{
		AuctionId: index,
	}, nil
}
