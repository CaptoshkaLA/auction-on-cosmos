package types

import (
	"testing"
)

func TestMsgEndBidValidateBasic(t *testing.T) {
	tests := []struct {
		name      string
		msg       MsgCreateBid
		expectErr bool
	}{
		{
			"Valid",
			*NewMsgCreateBid(1, 1, "Anton", 100, 1628814387),
			false,
		},
		{
			"Zero Index",
			*NewMsgCreateBid(0, 1, "cAnton", 100, 1628814387),
			true,
		},
		{
			"Zero Auction ID",
			*NewMsgCreateBid(1, 0, "Anton", 100, 1628814387),
			true,
		},
		{
			"Empty Bidder",
			*NewMsgCreateBid(1, 1, "", 100, 1628814387),
			true,
		},
		{
			"Zero Bid Amount",
			*NewMsgCreateBid(1, 1, "Anton", 0, 1628814387),
			true,
		},
		{
			"Zero CreatedAt",
			*NewMsgCreateBid(1, 1, "Anton", 100, 0),
			true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			err := tt.msg.ValidateBasic()
			if (err != nil) != tt.expectErr {
				t.Errorf("ValidateBasic() error = %v, expectErr %v", err, tt.expectErr)
			}
		})
	}
}
