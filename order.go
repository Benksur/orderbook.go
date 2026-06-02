package main

import (
	"time"
)

const (
	Buy Side = iota
	Sell
)

const (
	Market OrderType = iota
	Limit
	GoodTillCancelled
	FillOrKill
)

type Order struct {
	ID        OrderID
	Type      OrderType
	Side      Side
	Price     Price
	Quantity  Quantity
	Timestamp time.Time
}
