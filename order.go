package main

import (
	"time"
)

type OrderSide uint8

const (
	Buy OrderSide = iota
	Sell
)

type OrderType uint8

const (
	Market OrderType = iota
	Limit
	GoodTillCancelled
	FillOrKill
)

type Order struct {
	ID        string
	Type      OrderType
	Side      OrderSide
	Price     int64
	Quantity  uint64
	Timestamp time.Time
}
