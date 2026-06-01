package main

import (
	"fmt"
	"time"
)

type OrderBook struct {
	Bids []Order
	Asks []Order
}

func (ob *OrderBook) Add(order Order) {
	switch order.Side {
	case Buy:
		ob.Bids = append(ob.Bids, order)
	case Sell:
		ob.Asks = append(ob.Asks, order)
	}
}

func (ob *OrderBook) Cancel(orderId string) {

}

func (ob *OrderBook) GetOrder(orderId string) {

}

func (ob OrderBook) Print() {
	fmt.Println("========== BIDS ==========")
	for _, order := range ob.Bids {
		fmt.Printf("ID: %s, Type: %d, Price: %d, Quantity: %d, Time: %s\n",
			order.ID,
			order.Type,
			order.Price,
			order.Quantity,
			order.Timestamp.Format(time.TimeOnly))
	}
	fmt.Println("========== ASKS ==========")
	for _, order := range ob.Asks {
		fmt.Printf("ID: %s, Type: %d, Price: %d, Quantity: %d, Time: %s\n",
			order.ID,
			order.Type,
			order.Price,
			order.Quantity,
			order.Timestamp.Format(time.TimeOnly))
	}
}
