package main

import (
	"fmt"
	"sort"
)


// making field lowercase makes it private property
type OrderBook struct {
	Bids []Order `json:"bids"`
	Asks []Order `json:"asks"`
	Holdings []Holding `json:"Holdings"`
}

type Holding struct {
Secid string  `json:"Secid"`
Tradingaccount string `json:"TradACC"`
Custodian string `json:"Custodian"`
UserID string `json:"User"`
Quantity int32 `json:"qty"`
}

// APIs
// addBuyOrder(order)
// addSellOrder(order)
// removeBuyOrder(orderId)

// Add the new Order to end of orderbook in bids
func (book *OrderBook) addBuyOrder(order Order) {
	book.Bids = append(book.Bids, order)
	//then sort the order book
	sort.Slice(book.Bids, func(i, j int) bool {
		//fmt.Println("Adding buy order #", order.ID)
		if book.Bids[i].Price != book.Bids[j].Price {
			return book.Bids[i].Price < book.Bids[j].Price
		}
		return book.Bids[i].ID > book.Bids[j].ID
	})
}

func (book *OrderBook) addSellOrder(order Order) {

	book.Asks = append(book.Asks, order)
	sort.Slice(book.Asks, func(i, j int) bool {
		//fmt.Println("Adding Sell order #", order.ID)
		if book.Asks[i].Price != book.Asks[j].Price {
			return book.Asks[i].Price > book.Asks[j].Price
		}
		return book.Asks[i].ID > book.Asks[j].ID

	})
}

func (book *OrderBook) removeBuyOrder(index int) {
	fmt.Println("Removing Buy order #", book.Bids[index].ID)
	book.Bids = append(book.Bids[:index], book.Bids[index+1:]...)
}

func (book *OrderBook) removeSellOrder(index int) {
	fmt.Println("Removing Sell order #", book.Asks[index].ID)
	book.Asks = append(book.Asks[:index], book.Asks[index+1:]...)
}
