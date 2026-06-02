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
	ob.Add(Order{ID: uuid.NewString(), Type: Market, Side: Buy, Price: 109, Quantity: 1, Timestamp: time.Now()})
	ob.Add(Order{ID: uuid.NewString(), Type: Market, Side: Sell, Price: 103, Quantity: 1, Timestamp: time.Now()})
	ob.Add(Order{ID: uuid.NewString(), Type: Market, Side: Sell, Price: 101, Quantity: 1, Timestamp: time.Now()})
	ob.Add(Order{ID: uuid.NewString(), Type: Market, Side: Sell, Price: 100, Quantity: 1, Timestamp: time.Now()})
	ob.Add(Order{ID: uuid.NewString(), Type: Market, Side: Sell, Price: 106, Quantity: 1, Timestamp: time.Now()})

	order := Order{ID: uuid.NewString(), Type: Market, Side: Sell, Price: 108, Quantity: 1, Timestamp: time.Now()}
	ob.Add(order)
	ob.Print()
	fmt.Println("DELETING .....")
	ob.Cancel(order.ID)
	ob.Print()
}
