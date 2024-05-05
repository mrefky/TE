package main

import (
	"fmt"
	"time"
	//"sort"
	//"encoding/json"
	//	"github.com/shopspring/decimal"
	"strconv"
	//"github.com/eapache/go-resiliency/breaker"
	"github.com/Shopify/sarama"
	//cluster "github.com/bsm/sarama-cluster"
	//	"sync"
)

var Egx []Orderbook

func main() {

	producer := createProducer()
	records := readCsvFile("/home/mrefky/project/sec/Securities.csv")

	SecList := CreateSecurityList(records)

	var ord Orderbook
	var n string

	for i := 0; i < len(SecList); i++ {

		fmt.Sscan(SecList[i].id, &n)
		ord = Orderbook{seccode: SecList[i].Seccode}

		Egx = append(Egx, ord)
		//fmt.Println(ord)

	}

	records2 := readCsvFile("/home/mrefky/project/sec/Book1.csv")

	rawOrders := CreateOrdersList(records2)
	for i := 0; i < len(rawOrders); i++ {
		//fmt.Println(rawOrders[i])

		var h int
		h = getorderbook(Egx, rawOrders[i].Seccode)
		if h != 0 {
			var h1 Order
			h1.Price, _ = strconv.ParseFloat(rawOrders[i].Price, 10)
			h1.Myside = rawOrders[i].Myside
			fmt.Sscan(rawOrders[i].Ord_qty, &h1.Quantity)
			//h1.Quantity,_= strconv.ParseInt(rawOrders[i].Ord_qty)

		//	layout := "19/6/2007 11:00:20 AM"
		//	str := rawOrders[i].Timestamp

		//	t, _ := time.Parse(layout,str)

			//t, _ := time.Parse(layout , str)

			h1.Seccode = rawOrders[i].Seccode
			//ma, _ := strconv.ParseInt(rawOrders[i].Timestamp, 10, 0)
			h1.Timestamp=time.Now().Unix()
			fmt.Sscan(rawOrders[i].Id, &h1.Id)

			fmt.Println(h1.Id)
			Egx[h].Buys = append(Egx[h].Buys, h1)
			//fmt.Println(Egx[h].buys[0].id)
		}

	}

	var t int
	for i := 0; i < 10; i++ {
		//fmt.Println(Egx[i])
	}
	t = getorderbook(Egx, "EGMFC001MM19")

	fmt.Println("10 Best buys")
	for i := 0; i < 10; i++ {
		fmt.Println(i, " ", Egx[t].Buys[i].Price, Egx[t].Buys[i].Timestamp)

	}

	//	fmt.Println("10 Best Sells")

	//var wg sync.WaitGroup

	//	defer wg.Done()
	//wg.Add(1)
	for i := 0; i < len(Egx[t].Buys); i++ {

		//fmt.Println(i," ",Egx[t].sells[i].Price)
		//fmt.Println(Egx[t].sells[i].myside)

		raw := Egx[t].Buys[i].toJSON()

		fmt.Println(string(raw[:]))
		//fmt.Println(_json.Marshal)
		producer.Input() <- &sarama.ProducerMessage{
			Topic: "orders",
			Value: sarama.ByteEncoder(raw),
		
			Key: sarama.ByteEncoder("X-Stream"),
		}
		//fmt.Println(i)

	}

	//fmt.Println(Egx[t].Buys)
	fmt.Scanln()

	//wg.Done()
}

func createProducer() sarama.AsyncProducer {
	config := sarama.NewConfig()
	config.Producer.Return.Successes = false         // fire and forget
	config.Producer.Return.Errors = true             // notify on failed
	config.Producer.RequiredAcks = sarama.WaitForAll // waits for all insync replicas to commit

	producer, err := sarama.NewAsyncProducer([]string{"192.168.169.70:9094"}, config)

	if err != nil {
		fmt.Println("Unable to connect producer to kafka server")
	}

	return producer

}
