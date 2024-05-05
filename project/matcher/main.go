package main

import (
	//	"encoding/csv"
	"fmt"
	"log"

	"os"

	//"sys"

	//	"bytes"
	// "io"

	"database/sql"
	"strconv"
	"time"

	_ "github.com/go-sql-driver/mysql"

	"github.com/Shopify/sarama"
	// cluster "github.com/bsm/sarama-cluster"
	// "github.com/kailashyogeshwar85/slim-orderbook/cmd/slim-orderbook/engine"
)

var Kafka string = "192.168.169.70:9094"
var MYSQL string = "root@tcp(192.168.169.50:3306)/test"

//
//var Kafka  ="kafka.default.svc.cluster.local:9092"

//var MYSQL="root@tcp(mysql-0.default.svc.cluster.local:3306)/test"

const BG bool = true

type ShoppingRecord struct {
	ID      string
	Seccode string
}
type Egx struct {
	Secid     string
	Seccode   string
	orderbook OrderBook
}

func getorderbook(ex []Egx, seccode string) int {
	//var u int
	//u = 0
	for i := 0; i < len(ex); i++ {
		var y Egx = ex[i]

		if y.Seccode == seccode {
	//		u = i+1
			return i
		}

	}
	return -1
}

func Validate(ho []Holding, ord Order) bool {
return true
	if (ord.Side.String()) == "buy" {
		return true
	}

	for i := 0; i < len(ho); i++ {
		if ord.Quantity <= ho[i].Quantity && ord.Custodian == ho[i].Custodian && ord.TrdAcc == ho[i].Tradingaccount {
			return true

		}
		
	}

	fmt.Println("No Holdings , Order neglected >>>>>>>>>>>>>>>>>")

		return false
}


func main() {
	if BG {

		Kafka = "kafka.default.svc.cluster.local:9092"
		MYSQL = "root@tcp(mysql-0.default.svc.cluster.local:3306)/test"

		f, err := os.OpenFile("/mylogs/matcher.log", os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)

		//	//f, err := os.Create("/mylogs/myfile")
		if err != nil {
			log.Fatal(err)
		}
		defer f.Close()

		_, os.Stdout = os.Stdout, f
		log.SetOutput(f)
	}
	start := time.Now()
	fmt.Println("-----------------------------------------")
	fmt.Println("        EGX new matching Engine 1.0               ")
	fmt.Println("-----------------------------------------")
	// create consumer for consuming order
	var EGX []Egx
	//var EGX Egx= make("",[] OrderBook,0,300)
	//records := readCsvFile("./Securities.csv")
	//shoppingList := createShoppingList(records)
	var Dbcon string = MYSQL
	//Dbcon="root@tcp(192.168.1.50:3306)/test"
	db, err := sql.Open("mysql", Dbcon)

	if err != nil {
		panic(err.Error())
	}
	defer db.Close()
	var s string
	fmt.Println("Connecting to MySQL Success!")

	s = "select * from  Security"

	insert, err := db.Query(s)
	if err != nil {
		panic(err.Error())
	}
	//defer insert.Close()

	if err != nil {
		panic(err.Error())
	}
	defer insert.Close()
	var n, m string

	rows, err := db.Query("select secid,seccode from test.Security")
	if err != nil {
		log.Fatal(err)
	}
	defer rows.Close()
	var ord OrderBook = OrderBook{}
	//Ex := Egx{Secid: n, Seccode: m, orderbook: ord}
	for rows.Next() {
		err := rows.Scan(&n, &m)
		if err != nil {
			log.Fatal(err)
		}
		//	fmt.Println(n, m)
		Ex := Egx{Secid: n, Seccode: m, orderbook: ord}
		//EX =append(EX,Ex)
		EGX = append(EGX, Ex)
	}
	err = rows.Err()
	if err != nil {
		log.Fatal(err)
	}
	//	for p:=0;p<len(EGX);p++{
	//fmt.Println(EGX[p])
	//}
	fmt.Println("Prepared order books for " + strconv.Itoa(len(EGX)) + " Securities")
	// Now I need to read the outstanding open orders to cater for failure
	EGX = Outstanding(EGX)
	elapsed := time.Since(start)
	fmt.Println("****************************************************")
	log.Print("This took ", elapsed)
	fmt.Println("****************************************************")
	DebugAll(EGX)
	consumer := createConsumer()

	// create producer to send trades and orders data
	producer := createProducer()
	//producer2 := createProducer2()
	// create the order book
	book := OrderBook{
		Bids: make([]Order, 0, 10000),
		Asks: make([]Order, 0, 10000),
	}

	// create a channel to know when done
	done := make(chan bool)

	// start processing order
	go func() {
		validate(EGX)
		for msg := range consumer.Messages() {
			log.Print("_________GOT--", string(msg.Value[:]))
			var order Order
			// deserialize the message
			order.FromJSON(msg.Value)
			
		//	if order.MsgType !=""{ 
				
		//		if order.MsgType =="C" {
					//need to cancel that order
		//			return
		//		}

		//		if order.MsgType =="A" {
					//need to Amend that order
		//			return
		//		}



		//	}
			//Kinsert(order)
			// process the order
			fmt.Sprintln(order.Seccode)
			var xyz int = getorderbook(EGX, order.Seccode)
			if xyz == -1 {
				fmt.Println("Unkown Security; message neglected")
				//done <- true
				//consumer.MarkOffset(msg, "")
				//return

			}
			if xyz != -1 {
				book = EGX[xyz].orderbook
				//fmt.Println("Order book of " + EGX[xyz].Seccode)
				if Validate(EGX[xyz].orderbook.Holdings, order) {

					trades := book.Process(order)

					EGX[xyz].orderbook = book
					//log.Println("Trades length: ", len(trades))

					// send trades to message queue
					for _, trade := range trades {
						rawTrade := trade.ToJSON()
						log.Println("Publishing trade on topic -> trades")

						// publish the message over receiving channel
						producer.Input() <- &sarama.ProducerMessage{
							Topic: "trades",
							Value: sarama.ByteEncoder(rawTrade),
						}
					}
				}

				// mark the offset as commited
				consumer.MarkOffset(msg, "")
			}
		}

		done <- true
	}()
	<-done
}

var ii int

func DebugAll(egx []Egx) {
	for i := 0; i < len(egx); i++ {
		ii = i
		if len(egx[i].orderbook.Bids) > 0 || (len(egx[i].orderbook.Asks) > 0) {
			fmt.Println(egx[i].Seccode, " Asks=", len(egx[i].orderbook.Asks), " Bids=", len(egx[i].orderbook.Bids))
		}

	}
	fmt.Println("Order books for ", ii, " Securities created")

}
