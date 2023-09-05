package types

import (
	"fmt"

	sdk "github.com/cosmos/cosmos-sdk/types"
)

const (
	TypeMsgEndAuction = "end_auction"
)

var _ sdk.Msg = &MsgEndAuction{}

func NewMsgEndAuction(auctionId uint64, auctionOwner string) *MsgEndAuction {
	return &MsgEndAuction{
		AuctionOwner: auctionOwner,
		AuctionId:    auctionId,
	}
}

func (msg MsgEndAuction) Route() string { return RouterKey }

func (msg MsgEndAuction) Type() string { return TypeMsgEndAuction }

func (msg MsgEndAuction) GetSignBytes() []byte {
	bz := ModuleCdc.MustMarshalJSON(&msg)
	return sdk.MustSortJSON(bz)
}

func (msg MsgEndAuction) GetSigners() []sdk.AccAddress {
	creator, err := sdk.AccAddressFromBech32(msg.AuctionOwner)
	if err != nil {
		panic(err)
	}
	return []sdk.AccAddress{creator}
}

// ValidateBasic makes simple validation of message checks
func (msg MsgEndAuction) ValidateBasic() error {
	if msg.AuctionId == 0 {
		return fmt.Errorf("Auction Id cannot be 0")
	}
	return nil
}
