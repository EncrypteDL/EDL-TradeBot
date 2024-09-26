package internal

import "time"

// Limit represents an order that a limit type
type Limit struct {
	ClientID string
	Pair     Pair
	Side     Side
	BaseSize string
	Price    string
	PostOnly bool
	Expires  *time.Time
}

// Side represents the side of the trade that the order is placing.
type Side string

const (
	// SideBuy specifies the buy side of an order
	SideBuy Side = "BUY"

	// SideSell specifies the sell side of an order
	SideSell = "SELL"
)
