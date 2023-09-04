package auctiononcosmos_test

import (
	"testing"

	keepertest "github.com/CaptoshkaLA/AuctionOnCosmos/testutil/keeper"
	"github.com/CaptoshkaLA/AuctionOnCosmos/testutil/nullify"
	"github.com/CaptoshkaLA/AuctionOnCosmos/x/auctiononcosmos"
	"github.com/CaptoshkaLA/AuctionOnCosmos/x/auctiononcosmos/types"
	"github.com/stretchr/testify/require"
)

func TestGenesis(t *testing.T) {
	genesisState := types.GenesisState{
		Params: types.DefaultParams(),

		// this line is used by starport scaffolding # genesis/test/state
	}

	k, ctx := keepertest.AuctiononcosmosKeeper(t)
	auctiononcosmos.InitGenesis(ctx, *k, genesisState)
	got := auctiononcosmos.ExportGenesis(ctx, *k)
	require.NotNil(t, got)

	nullify.Fill(&genesisState)
	nullify.Fill(got)

	// this line is used by starport scaffolding # genesis/test/assert
}
