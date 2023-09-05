package types

import (
	"testing"
)

func TestMsgCreateAuctionValidateBasic(t *testing.T) {
	tests := []struct {
		name      string
		msg       MsgCreateAuction
		expectErr bool
	}{
		{
			"Valid",
			*NewMsgCreateAuction("Anton", "An amazing item", "open", 10, 1, 2),
			false,
		},
		{
			"Empty Owner",
			*NewMsgCreateAuction("", "An amazing item", "open", 10, 1, 2),
			true,
		},
		{
			"Empty Description",
			*NewMsgCreateAuction("Anton", "", "open", 10, 1, 2),
			true,
		},
		{
			"Zero Price",
			*NewMsgCreateAuction("Anton", "An amazing item", "open", 0, 1, 2),
			true,
		},
		{
			"Invalid End Time",
			*NewMsgCreateAuction("Anton", "An item", "open", 10, 2, 1),
			true,
		},
		{
			"Invalid Status",
			*NewMsgCreateAuction("Anton", "An item", "unknown", 10, 1, 2),
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
