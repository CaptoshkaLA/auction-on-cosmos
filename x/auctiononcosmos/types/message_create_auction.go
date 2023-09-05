package types

import (
	"fmt"
	sdk "github.com/cosmos/cosmos-sdk/types"
)

const (
	TypeMsgCreateAuction = "create_auction"
)

// NewMsgCreateAuction creates a new instance of the msg
func NewMsgCreateAuction(auctionOwner, itemDescription, auctionStatus string, initialPrice, createdAt, endTime uint64) *MsgCreateAuction {
	return &MsgCreateAuction{
		AuctionOwner:    auctionOwner,
		ItemDescription: itemDescription,
		InitialPrice:    initialPrice,
		CreatedAt:       createdAt,
		EndTime:         endTime,
		AuctionStatus:   auctionStatus,
	}
}

func (msg MsgCreateAuction) Route() string { return RouterKey }

func (msg MsgCreateAuction) Type() string { return TypeMsgCreateAuction }

func (msg MsgCreateAuction) ValidateBasic() error {
	if msg.AuctionOwner == "" {
		return fmt.Errorf("Auction Owner cannot be empty")
	}

	if len(msg.ItemDescription) == 0 {
		return fmt.Errorf("Item Description cannot be empty")
	}

	if msg.InitialPrice == 0 {
		return fmt.Errorf("Initial Price must be greater than zero")
	}

	if msg.CreatedAt >= msg.EndTime {
		return fmt.Errorf("The end time should be longer than the creation time")
	}

	if msg.AuctionStatus != "open" && msg.AuctionStatus != "closed" {
		return fmt.Errorf("Invalid auction status")
	}

	return nil
}

func (msg MsgCreateAuction) GetSignBytes() []byte {
	return sdk.MustSortJSON(ModuleCdc.MustMarshalJSON(&msg))
}

func (msg MsgCreateAuction) GetSigners() []sdk.AccAddress {
	creator, err := sdk.AccAddressFromBech32(msg.AuctionOwner)
	if err != nil {
		panic(err)
	}
	return []sdk.AccAddress{creator}
}
