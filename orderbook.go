package main

import (
	"fmt"
	"slices"
	"sort"
	"time"
)

type BookSide struct {
	Levels     map[Price][]Order
	PriceIndex []Price
	OrderIndex map[OrderID]Price
}

type OrderBook struct {
	Bids BookSide
	Asks BookSide
}

/* ================================= PRIVATE API =========================================== */

func insertSorted(prices []Price, target Price, less func(a, b Price) bool) []Price {
	i := sort.Search(len(prices), func(i int) bool {
		return !less(prices[i], target)
	})
	prices = append(prices, 0)
	copy(prices[i+1:], prices[i:])
	prices[i] = target
	return prices
}

func addToBookSide(bs *BookSide, o Order, less func(a, b Price) bool) {
	if _, ok := bs.Levels[o.Price]; !ok {
		bs.PriceIndex = insertSorted(bs.PriceIndex, o.Price, less)
	}
	bs.Levels[o.Price] = append(bs.Levels[o.Price], o)
	bs.OrderIndex[o.ID] = o.Price
}

func remove(orderId OrderID, bs *BookSide, p Price) {
	level := bs.Levels[p]
	for idx, order := range level {
		if order.ID == orderId {
			bs.Levels[p] = slices.Delete(level, idx, idx+1)
			// if no more orders at that price
			if len(bs.Levels[p]) == 0 {
				delete(bs.Levels, p)
				bs.PriceIndex = slices.DeleteFunc(bs.PriceIndex, func(n Price) bool {
					return n == p
				})
			}
		}
	}
	delete(bs.OrderIndex, orderId)
}

/* ================================= PUBLIC API =========================================== */

func (ob *OrderBook) Add(order Order) {
	switch order.Side {
	case Buy:
		addToBookSide(&ob.Bids, order, func(a, b Price) bool { return a > b }) // descending
	case Sell:
		addToBookSide(&ob.Asks, order, func(a, b Price) bool { return a < b }) // ascending
	}
}

func (ob *OrderBook) Cancel(orderId OrderID) {
	if price, ok := ob.Asks.OrderIndex[orderId]; ok {
		remove(orderId, &ob.Asks, price)
		return
	}
	if price, ok := ob.Bids.OrderIndex[orderId]; ok {
		remove(orderId, &ob.Bids, price)
		return
	}
}

func (ob *OrderBook) Match() {
	for {
		//always pop from head of queue. will need to track indices for FOK
		bestBidLevel := ob.Bids.Levels[ob.Bids.PriceIndex[0]]
		bestBid := &bestBidLevel[0]
		bestAskLevel := ob.Asks.Levels[ob.Asks.PriceIndex[0]]
		bestAsk := &bestAskLevel[0]

		fmt.Printf("matching bid with price: %d, quantity %d\n", bestBid.Price, bestBid.Quantity)
		fmt.Printf("matching ask with price: %d, quantity %d\n", bestAsk.Price, bestAsk.Quantity)

		if bestAsk.Price <= bestBid.Price {
			fmt.Println("made trade")

			if bestBid.Quantity > bestAsk.Quantity {
				bestBid.Quantity -= bestAsk.Quantity
				remove(bestAsk.ID, &ob.Asks, bestAsk.Price)
			} else if bestAsk.Quantity > bestBid.Quantity {
				bestAsk.Quantity -= bestBid.Quantity
				remove(bestBid.ID, &ob.Bids, bestBid.Price)
			} else {
				remove(bestAsk.ID, &ob.Asks, bestAsk.Price)
				remove(bestBid.ID, &ob.Bids, bestBid.Price)
			}
		} else {
			break
		}

	}
}

func (ob OrderBook) Print() {
	fmt.Println("========== BIDS ==========")
	for _, level := range ob.Bids.PriceIndex {
		for _, order := range ob.Bids.Levels[level] {
			fmt.Printf("ID: %s, Type: %d, Price: %d, Quantity: %d, Time: %s\n",
				order.ID,
				order.Type,
				order.Price,
				order.Quantity,
				order.Timestamp.Format(time.TimeOnly))

		}
	}
	fmt.Println("========== ASKS ==========")
	for _, level := range ob.Asks.PriceIndex {
		for _, order := range ob.Asks.Levels[level] {
			fmt.Printf("ID: %s, Type: %d, Price: %d, Quantity: %d, Time: %s\n",
				order.ID,
				order.Type,
				order.Price,
				order.Quantity,
				order.Timestamp.Format(time.TimeOnly))
		}
	}
}
