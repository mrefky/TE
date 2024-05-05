package main

import (
	"fmt"
	"log"
	//"github.com/shopspring/decimal"
	//	"strconv"
)

// Process an order and return the trades generated before adding the remaining amount to the market
func (book *OrderBook) Process(order Order) []Trade {
	DebugME(book, order.Seccode)
	if order.Side.String() == "buy" {
		//DebugME(book)
		fmt.Println("Process buy order")
		return book.processLimitBuyOrder(order)
	}
	//DebugME(book)
	fmt.Println("Process Sell order")
	return book.processLimitSellOrder(order)
}

func (book *OrderBook) processLimitBuyOrder(order Order) []Trade {
	fmt.Println("Processing LIMIT BUY ORDER of quantity=", order.Quantity, " and Price=", order.Price)
	//DebugME(book)
	//DebugME(book)
	// create a trade object
	trades := make([]Trade, 0, 1)

	n := len(book.Asks)
	//fmt.Println("Found ", n, " sell orders")
	if n == 0 {
		//	fmt.Println("Adding the buy order to the book since ther are no sell orders")
		book.addBuyOrder(order)
		DebugME(book, order.Seccode)

		return trades
	}
	//////fmt.Println("check if we have atleast one matching order")
	//DebugME(book)
	// first compare with last sell price
	// loop only when the new order price is less than last sell price
	if n != 0 || book.Asks[n-1].Price < (order.Price) {
		// traverse all orders that match
		for i := n - 1; i >= 0; i-- {
			sellOrder := book.Asks[i]
			//fmt.Println("checking sell order of qty=", sellOrder.Quantity, "@", sellOrder.Price)
			// last sell price is 10 limit = 9
			// return as the highest sell order is gt limit price
			if sellOrder.Price > (order.Price) {
				//	fmt.Println("sell order cant match since it has a higher price")
				break
			}
			// fill the entire order when sellOrder has higher quantity
			if sellOrder.Quantity >= (order.Quantity) {
				DebugME(book, order.Seccode)
				trades = append(
					trades,
					Trade{
						TakerOrderID: order.ID,     // TakerOrderID
						MakerOrderID: sellOrder.ID, // Maker OrderID
						Quantity:     order.Quantity,
						Price:        sellOrder.Price,
						Timestamp:    order.Timestamp,
					},
				)
				DebugME(book, order.Seccode)
				fmt.Println("New buy order fully matched, and sellordr is patially or fully matched")

				sellOrder.Quantity = sellOrder.Quantity - (order.Quantity)
				book.Asks[i].Quantity = sellOrder.Quantity
				order.Quantity = order.Quantity - (order.Quantity)
				if sellOrder.Quantity == 0 {
					fmt.Println("Removing fully matched sell order")
					book.removeSellOrder(i)
				}
				//fmt.Println("we have ", len(book.Bids), " buy orders", " And ", len(book.Asks), " Sell orders in the book")
				DebugME(book, order.Seccode)
				return trades

			}
			// partially fill the order and continue
			if sellOrder.Quantity < (order.Quantity) {
				trades = append(
					trades,
					Trade{
						TakerOrderID: order.ID,
						MakerOrderID: sellOrder.ID,
						Quantity:     sellOrder.Quantity,
						Price:        sellOrder.Price,
						Timestamp:    order.Timestamp,
					},
				)
				order.Quantity = order.Quantity - (sellOrder.Quantity)
				// remove the sell Order as all quantities are filled by bid
				fmt.Println("Removing sell order from the book siince it is fully matched")
				book.removeSellOrder(i)
				if order.Quantity == 0 {
					fmt.Println("Removing fully matched order of", order.Quantity, "@", order.Price)

				}
				continue
			}
		}
	}
	// finally add the order with remaining qty to book
	//fmt.Println("q=", (order.Quantity))

	if order.Quantity > 0 {
		fmt.Println("Adding the new order to book", order.Side.String(), order.Quantity, ",@", order.Price)
		book.addBuyOrder(order)
	}
	//fmt.Println("we have ", len(book.Bids), " buy orders", " And ", len(book.Asks), " Sell orders in the book")
	DebugME(book, order.Seccode)
	return trades
}

func (book *OrderBook) processLimitSellOrder(order Order) []Trade {
	log.Println("Processing LIMIT SELL ORDER of ", order.Quantity, "@", order.Price)

	trades := make([]Trade, 0, 1)
	n := len(book.Bids)
	//fmt.Println("Found ", n, " buy orders in the book")
	if n == 0 {
		DebugME(book, order.Seccode)
		book.addSellOrder(order)
		fmt.Println("Adding the new order to the sell book since no buy orders exists")
		DebugME(book, order.Seccode)
		return trades
	}

	// Proceed only if the sell Price is Greather than user highest buy Price
	if n != 0 || book.Bids[n-1].Price >= (order.Price) {
		// travers all bids that match
		for i := n - 1; i >= 0; i-- {
			buyOrder := book.Bids[i]

			if buyOrder.Price < (order.Price) {
				fmt.Println("NO morder orders that can match ....")
				break // exit
			}

			// fill the entire order of buy order is gte
			if buyOrder.Quantity >= (order.Quantity) {
				trades = append(
					trades,
					Trade{
						TakerOrderID: order.ID,
						MakerOrderID: buyOrder.ID,
						Quantity:     order.Quantity,
						Price:        buyOrder.Price,
						Timestamp:    order.Timestamp,
					},
				)

				fmt.Println("Comming sell order is fully Matched")

				buyOrder.Quantity = buyOrder.Quantity - (order.Quantity)
				order.Quantity = 0 // order.Quantity - (buyOrder.Quantity)
				book.Bids[i].Quantity = buyOrder.Quantity
				//DebugME(book, order.Seccode)
				if buyOrder.Quantity == 0 {
					fmt.Println("Removing order #", buyOrder.ID)
					book.removeBuyOrder(i)
				//	DebugME(book, order.Seccode)
				}
				fmt.Println("Order MAtched")
				//DebugME(book)
				fmt.Println("Sell order remaining quantity ", order.Quantity)
				fmt.Println(order.Quantity)

				//DebugME(book, order.Seccode)

				//buyOrder.Quantity = 0
				return trades
			}

			// fill a partial order and continue
			if buyOrder.Quantity < (order.Quantity) {
				trades = append(
					trades,
					Trade{
						TakerOrderID: order.ID,
						MakerOrderID: buyOrder.ID,
						Quantity:     buyOrder.Quantity,
						Price:        buyOrder.Price,
						Timestamp:    order.Timestamp,
					},
				)
				//	fmt.Println(order.Quantity.String())

				order.Quantity = order.Quantity - (buyOrder.Quantity)
				//buyOrder.Quantity=buyOrder.Quantity.Sub(order.Quantity)
				if order.Quantity == 0 {
					//need to also remove the order
					//order.Quantity = 0
				//	DebugME(book, order.Seccode)
					return trades
				}
				book.removeBuyOrder(i)

				continue
			}
		}
	}

	//fmt.Println("<<<<<<<<", order.Quantity)
	book.addSellOrder(order)
	//DebugME(book, order.Seccode)
	return trades
}

func DebugME(book *OrderBook, seccode string) {
	///return
	fmt.Println("---------------- " + seccode + " Order Book----------------------")
	fmt.Println("Asks =", len(book.Asks), " Bids =", len(book.Bids))

	return
	if len(book.Asks) > 0 {
		fmt.Println("Asks")

		for i := len(book.Asks) - 1; i >= 0; i-- {
			fmt.Println(book.Asks[i].ID, ":", book.Asks[i].Quantity, "@", book.Asks[i].Price)

		}
		fmt.Println()
	}
	if len(book.Bids) > 0 {
		fmt.Println("Bids")
		for i := len(book.Bids) - 1; i >= 0; i-- {
			fmt.Println(book.Bids[i].ID, ":", book.Bids[i].Quantity, "@", book.Bids[i].Price)
		}

	}
	fmt.Println("--------------------------------------")
}
