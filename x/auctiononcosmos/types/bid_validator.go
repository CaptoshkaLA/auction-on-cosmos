package types

import (
	"errors"
)

// ValidateBid checks if the fields in a Bid are valid.
func ValidateBid(bid Bid, highestBid Bid, auction Auction) error {
	// Basic Field Validation
	if bid.Index == 0 {
		return errors.New("Invalid bid index")
	}

	if bid.AuctionID == 0 {
		return errors.New("Invalid auction ID")
	}

	if bid.Bidder == "" {
		return errors.New("Invalid bidder")
	}

	if bid.BidAmount == 0 {
		return errors.New("Bid amount should be greater than zero")
	}

	if bid.CreatedAt == 0 {
		return errors.New("Invalid creation time")
	}

	// Check if the auction is active
	if auction.AuctionStatus != "active" {
		return errors.New("Cannot place bid on inactive or completed auction")
	}

	// Check if the new bid is higher than the current highest bid
	if highestBid.BidAmount >= bid.BidAmount {
		return errors.New("New bid amount must be higher than the current highest bid")
	}

	// TODO: checking that the user has enough money to place this bet

	return nil
}
