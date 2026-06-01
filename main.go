package main

import (
	"fmt"
	"github.com/google/uuid"
	"time"
)

func main() {
	fmt.Println("Creating OrderBook")
	ob := OrderBook{}
	ob.Add(Order{ID: uuid.New().String(), Type: Market, Side: Buy, Price: 100, Quantity: 1, Timestamp: time.Now()})
	ob.Print()
}
