package types

import "fmt"

const (
	// ModuleName defines the module name
	ModuleName = "auctiononcosmos"

	// StoreKey defines the primary module store key
	StoreKey = ModuleName

	// RouterKey defines the module's message routing key
	RouterKey = ModuleName

	// MemStoreKey defines the in-memory store key
	MemStoreKey = "mem_auctiononcosmos"

	// KeyLastAuctionID stores the last used ID for the auction
	KeyLastAuctionID = "LastAuctionID"
)

func KeyPrefix(p string) []byte {
	return []byte(p)
}

// KeyAuction returns the key for storing a specific auction
func KeyAuction(index uint64) []byte {
	return []byte(fmt.Sprintf("Auction-%d", index))
}
