package main

import (
	"fmt"
	"github.com/google/uuid"
	"time"
)

func main() {
	fmt.Println("Creating OrderBook")
	bids := BookSide{Map: make(map[Price][]Order), Index: make([]Price, 0)}
	asks := BookSide{Map: make(map[Price][]Order), Index: make([]Price, 0)}
	ob := OrderBook{Bids: bids, Asks: asks}
	ob.Add(Order{ID: uuid.NewString(), Type: Market, Side: Buy, Price: 100, Quantity: 1, Timestamp: time.Now()})
	ob.Print()
}
