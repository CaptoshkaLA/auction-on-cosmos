package types

import (
	"errors"
)

// Validate checks if the fields in an Auction are valid.
func (a Auction) Validate() error {
	if a.Index == 0 {
		return errors.New("Invalid auction index")
	}

	if a.AuctionOwner == "" {
		return errors.New("Invalid auction owner")
	}

	if a.ItemDescription == "" {
		return errors.New("Invalid item description")
	}

	if a.InitialPrice == 0 {
		return errors.New("Initial price should be greater than zero")
	}

	if a.CreatedAt == 0 {
		return errors.New("Invalid creation time")
	}

	if a.EndTime <= a.CreatedAt {
		return errors.New("End time should be after creation time")
	}

	if a.AuctionStatus != "active" && a.AuctionStatus != "completed" {
		return errors.New("Invalid auction status")
	}

	return nil
}
