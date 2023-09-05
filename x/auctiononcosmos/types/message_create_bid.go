package types

import (
	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
)

const (
	TypeMsgCreateBid = "create_bid"
)

func NewMsgCreateBid(index uint64, auctionId uint64, bidder string, bidAmount uint64, createdAt uint64) *MsgCreateBid {
	return &MsgCreateBid{
		Index:     index,
		AuctionId: auctionId,
		Bidder:    bidder,
		BidAmount: bidAmount,
		CreatedAt: createdAt,
	}
}

func (msg MsgCreateBid) Route() string { return RouterKey }
func (msg MsgCreateBid) Type() string  { return TypeMsgCreateBid }

func (msg MsgCreateBid) ValidateBasic() error {
	if msg.Index == 0 || msg.AuctionId == 0 || msg.Bidder == "" || msg.BidAmount == 0 || msg.CreatedAt == 0 {
		return sdkerrors.Wrap(sdkerrors.ErrInvalidRequest, "one or more arguments are zero or empty")
	}
	return nil
}

func (msg MsgCreateBid) GetSignBytes() []byte {
	return sdk.MustSortJSON(ModuleCdc.MustMarshalJSON(&msg))
}

func (msg MsgCreateBid) GetSigners() []sdk.AccAddress {
	creator, err := sdk.AccAddressFromBech32(msg.Bidder)
	if err != nil {
		panic(err)
	}
	return []sdk.AccAddress{creator}
}
