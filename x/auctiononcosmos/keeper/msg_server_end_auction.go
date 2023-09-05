package keeper

import (
	"context"

	"github.com/CaptoshkaLA/AuctionOnCosmos/x/auctiononcosmos/types"
	sdk "github.com/cosmos/cosmos-sdk/types"
)

func (k msgServer) EndAuction(goCtx context.Context, msg *types.MsgEndAuction) (*types.MsgEndAuctionResponse, error) {
	ctx := sdk.UnwrapSDKContext(goCtx)

	if err := k.EndAuctionAuxiliary(ctx, msg.AuctionId); err != nil {
		return nil, err
	}

	return &types.MsgEndAuctionResponse{}, nil
}
