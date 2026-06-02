package main

import (
	"fmt"
	"github.com/google/uuid"
	"time"
)

func main() {
	fmt.Println("Creating OrderBook")
	bids := BookSide{Levels: make(map[Price][]Order), PriceIndex: make([]Price, 0), OrderIndex: make(map[OrderID]Price)}
	asks := BookSide{Levels: make(map[Price][]Order), PriceIndex: make([]Price, 0), OrderIndex: make(map[OrderID]Price)}
	ob := OrderBook{Bids: bids, Asks: asks}
	ob.Add(Order{ID: uuid.NewString(), Type: Market, Side: Buy, Price: 100, Quantity: 1, Timestamp: time.Now()})
	ob.Add(Order{ID: uuid.NewString(), Type: Market, Side: Buy, Price: 105, Quantity: 1, Timestamp: time.Now()})
	ob.Add(Order{ID: uuid.NewString(), Type: Market, Side: Buy, Price: 102, Quantity: 1, Timestamp: time.Now()})
	ob.Add(Order{ID: uuid.NewString(), Type: Market, Side: Buy, Price: 107, Quantity: 1, Timestamp: time.Now()})
	ob.Add(Order{ID: uuid.NewString(), Type: Market, Side: Buy, Price: 109, Quantity: 3, Timestamp: time.Now()})
	ob.Add(Order{ID: uuid.NewString(), Type: Market, Side: Sell, Price: 103, Quantity: 2, Timestamp: time.Now()})
	ob.Add(Order{ID: uuid.NewString(), Type: Market, Side: Sell, Price: 101, Quantity: 1, Timestamp: time.Now()})
	ob.Add(Order{ID: uuid.NewString(), Type: Market, Side: Sell, Price: 100, Quantity: 1, Timestamp: time.Now()})
	ob.Add(Order{ID: uuid.NewString(), Type: Market, Side: Sell, Price: 106, Quantity: 1, Timestamp: time.Now()})

	ob.Print()
	fmt.Println()
	fmt.Println()
	fmt.Println("=============== Begin Matching ==============")
	ob.Match()
	fmt.Println("=============== After Matching ==============")
	fmt.Println()
	fmt.Println()
	ob.Print()
}
