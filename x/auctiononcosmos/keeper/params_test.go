package keeper_test

import (
	"testing"

	testkeeper "github.com/CaptoshkaLA/AuctionOnCosmos/testutil/keeper"
	"github.com/CaptoshkaLA/AuctionOnCosmos/x/auctiononcosmos/types"
	"github.com/stretchr/testify/require"
)

func TestGetParams(t *testing.T) {
	k, ctx := testkeeper.AuctiononcosmosKeeper(t)
	params := types.DefaultParams()

	k.SetParams(ctx, params)

	require.EqualValues(t, params, k.GetParams(ctx))
}
