package keeper

import (
	"github.com/CaptoshkaLA/AuctionOnCosmos/x/auctiononcosmos/types"
)

var _ types.QueryServer = Keeper{}
