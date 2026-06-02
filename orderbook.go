package main

import (
	"fmt"
	"sort"
	"time"
)

type BookSide struct {
	Map   map[Price][]Order
	Index []Price
}

type OrderBook struct {
	Bids BookSide
	Asks BookSide
}

func insertSorted(prices []Price, target Price) []Price {
	i := sort.Search(len(prices), func(i int) bool {
		return prices[i] >= target
	})
	prices = append(prices, 0)
	copy(prices[i+1:], prices[i:])
	prices[i] = target
	return prices
}

func addToBook(bs *BookSide, o Order) {
	if _, ok := bs.Map[o.Price]; !ok {
		bs.Index = insertSorted(bs.Index, o.Price)
	}
	bs.Map[o.Price] = append(bs.Map[o.Price], o)
}

func (ob *OrderBook) Add(order Order) {
	switch order.Side {
	case Buy:
		addToBook(&ob.Asks, order)
	case Sell:
		addToBook(&ob.Bids, order)
	}
}

func (ob *OrderBook) Cancel(orderId string) {

}

func (ob *OrderBook) GetOrder(orderId string) {

}

func (ob OrderBook) Print() {
	fmt.Println("========== BIDS ==========")
	for _, level := range ob.Bids.Map {
		for _, order := range level {
			fmt.Printf("ID: %s, Type: %d, Price: %d, Quantity: %d, Time: %s\n",
				order.ID,
				order.Type,
				order.Price,
				order.Quantity,
				order.Timestamp.Format(time.TimeOnly))

		}
	}
	fmt.Println("========== ASKS ==========")
	for _, level := range ob.Asks.Map {
		for _, order := range level {
			fmt.Printf("ID: %s, Type: %d, Price: %d, Quantity: %d, Time: %s\n",
				order.ID,
				order.Type,
				order.Price,
				order.Quantity,
				order.Timestamp.Format(time.TimeOnly))
		}
	}
}
