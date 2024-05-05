package main

import (
	"database/sql"
	"encoding/json"
	"fmt"
	"log"
	"strconv"
	"sync"
	"time"
"os"
	"github.com/Shopify/sarama"
	cluster "github.com/bsm/sarama-cluster"
	_ "github.com/go-sql-driver/mysql"
	//"os"
)

type Order struct {
	ID        int32 `json:"id"`
	Side      string  `json:"side"`
	Quantity  int32  `json:"Quantity"`
	Price     float64 `json:"Price"`
	Timestamp string `json:"Timestamp"`
	Seccode  string `json:"seccode"`
}

type Trade struct {
	Taker_Order_ID int32
	Maker_Order_ID int32
	Quantity       int32
	Price          float64
	Timestamp      int64
	Seccode string `json:"seccode"`

}

func main() {

 f1, err := os.OpenFile("/mylogs/sql.log", os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)

        //f, err := os.Create("/mylogs/myfile")
        if err != nil {
                log.Fatal(err)
        }
        defer f1.Close()

        _, os.Stdout = os.Stdout, f1
        log.SetOutput(f1)
	var wg sync.WaitGroup
	wg.Add(2)
	go func() {
		f()

		wg.Done()

	}()

	//wg1.Add(1)
	go func() {
		ftrade()

		wg.Done()

	}()
fmt.Println(">>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>")
	log.Printf("Staring to Update Mysql")
	db, err := sql.Open("mysql", "root:pass@tcp(mysql-0.default.svc.cluster.local:3306)/test")
	if err != nil {
		panic(err.Error())
	}
	defer db.Close()
	log.Printf("Connecting to MySQL Success!")
	fmt.Println(">>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>")
	//done := make(chan string)
	//go f()

	//fmt.Println(done)

	//go f()

	//go  ftrade()
	//go ftrade()

	wg.Wait()
}
func (order *Order) FromJSON(msg []byte) error {
	return json.Unmarshal(msg, order)
}

func (trd *Trade) FromJSON(msg []byte) error {
	return json.Unmarshal(msg, trd)
}

func f() {
	//endless loop reading orders topis and writing orders to orders table
	consumer := createConsumer()
	//if len(sql) != 0 {

	for msg1 := range consumer.Messages() {
		//	var myString string

		//	fmt.Println(string(msg.Value))

		consumer.MarkOffset(msg1, "")

		var ord Order
		//fmt.Println(msg)
		ord.FromJSON(msg1.Value)
		//fmt.Println(ord)
		var s string
		s = "INSERT INTO orders set orderid=" + fmt.Sprint(ord.ID) + ",Quantity='" + fmt.Sprint(ord.Quantity) + "',Price='" + fmt.Sprintf("%f", ord.Price) + "'"
		s = s+ ",ordtype='" + ord.Side+"'"  
		s=s+",ID=" + fmt.Sprint(int64(ord.ID))+",timestamp='" + ord.Timestamp + "'"
		s=s+",secid='"+ord.Seccode+"'"
		fmt.Print("o")

		//s1= "INSERT INTO matches set orderid='" + ord.ID + "',Quantity=0"
		//fmt.Println(s)
		//fmt.Print(".")
		//call the inset into orders function
		//s=s+";"+s1+";"
		//done := make(chan string)
		kupdate(s)
		//fmt.Println(<-done)
		// kupdate(s1)
		//fmt.Println(<-done)
		//need to wait tilll the above finishe
		//kupdate(string(msg.Value))
		//time.Sleep(time.Second * 1)
		//var x string
		//x=s + s1

	}
}

func ftrade() {
	var trdID int
	trdID = 0
	consumer2 := createConsumer2()
	//if len(sql) != 0 {

	for msg := range consumer2.Messages() {
		//	var myString string

		//	fmt.Println(string(msg.Value))

		consumer2.MarkOffset(msg, "")
		var trd Trade
		//fmt.Println(msg)
		trd.FromJSON(msg.Value)
		//fmt.Println(trd)
		//return
		var s1 string
		//var mqty string = fmt.Sprint(trd.Quantity)
		//s = "update matches set Quantity=Quantity+" + mqty
		//s = s + ",Matched_Quantity= Matched_Quantity+" + mqty
		//s = s + " where orderid=" + trd.Maker_Order_ID
		//s = s + " or orderid=" + trd.Taker_Order_ID
		 fmt.Print("T")
		s1 = "insert into trades set MakerID=" + fmt.Sprint(trd.Maker_Order_ID)
		s1 = s1 + ",TakerID=" + fmt.Sprint(trd.Taker_Order_ID)
		s1 = s1 + ",Price='" + fmt.Sprintf("%f", trd.Price) + "'"
		s1 = s1 + ",Quantity='" + fmt.Sprint(trd.Quantity) + "'"
		s1 = s1 + ",timestamp='" + strconv.FormatInt(trd.Timestamp, 10) + "'"
		s1 = s1 + ",tradeid='" + fmt.Sprint(trdID) + "';"
		//fmt.Println(s1)
		//fmt.Fprintf(file, "%v\n", )
		trdID = trdID + 1
		//s1 = "update order set qunatity="
		//fmt.Println(s1)
		//fmt.Println(s1)
		//s=s1+";"+s+";"
		//done1 := make(chan string)
		kupdate(s1)
		//done2 := make(chan string)

		//kupdate(s)
		//fmt.Println(<-done2)
		//kupdate(string(msg.Value))
		//time.Sleep(time.Second * 1)
	}
}

func createConsumer() *cluster.Consumer {
	// define the configuration for our cluster
	config := cluster.NewConfig()
	config.Consumer.Return.Errors = true
	config.Group.Return.Notifications = false
	config.Consumer.Offsets.Initial = sarama.OffsetOldest // earliest uncommited offset
	config.Consumer.Offsets.CommitInterval = time.Second

	orderTopic := []string{"orders"}

	// log.Println("Listening for SQL on topic -> ", orderTopic)
	// create the consumer
	consumer, err := cluster.NewConsumer(
		[]string{"kafka.default.svc.cluster.local:9092"},
		"SQL-Orders",
		orderTopic,
		config,
	)

	if err != nil {
		log.Fatal("Unable to connect to kafka cluster")
	}else{
	 fmt.Println("<<<<<<<<<<<<<<<<<<<<<<<<<<<<<<<<<<<<<<<<<<<<<<<<<<<<<<<<<<<<<<<<")
	 log.Printf("Connected to Kafka")
	 fmt.Println("<<<<<<<<<<<<<<<<<<<<<<<<<<<<<<<<<<<<<<<<<<<<<<<<<<<<<<<<<<<<<<<<")
 }
	go handleErrors(consumer)
	go handleNotifications(consumer)
	return consumer
}

func handleErrors(consumer *cluster.Consumer) {
	for err := range consumer.Errors() {
		log.Printf("Error: %s\n", err.Error())
	}
}

func handleNotifications(consumer *cluster.Consumer) {
	for ntf := range consumer.Notifications() {
		log.Printf("Rebalanced %+v\n", ntf)
	}
}
func kupdate(sql1 string) {

	//fmt.Println(sql1)
	//fmt.Print("T")
	db, err := sql.Open("mysql", "root@tcp(mysql-0.default.svc.cluster.local:3306)/test")

	if err != nil {
		panic(err.Error())
	}
	defer db.Close()
	//var s string
	//var newqty int64=ord.Quantity-qty
	//s="update orders set Quantity=Quantity-"+strconv.FormatInt(mqty,10) +", Matched_Quantity= Matched_Quantity+'"+strconv.FormatInt(mqty,10)+"' where orderid="+Id
	//fmt.Println(">>>>>>>>>>>>>>",s)

	insert, err := db.Query(sql1)
	if err != nil {
		panic(err.Error())
	}
	defer insert.Close()
	//      fmt.Println(Id," Order updaeted")
	//done <- sql1
}

//()
//<-done
//}
func createConsumer2() *cluster.Consumer {
	// define the configuration for our cluster
	config := cluster.NewConfig()
	config.Consumer.Return.Errors = true
	config.Group.Return.Notifications = false
	config.Consumer.Offsets.Initial = sarama.OffsetOldest // earliest uncommited offset
	config.Consumer.Offsets.CommitInterval = time.Second

	orderTopic := []string{"trades"}

	// log.Println("Listening for SQL on topic -> ", orderTopic)
	// create the consumer
	consumer, err := cluster.NewConsumer(
		[]string{"kafka.default.svc.cluster.local:9092"},
		"SQL-Trades",
		orderTopic,
		config,
	)

	if err != nil {
		log.Fatal("Unable to connect to kafka cluster")
	}

	go handleErrors2(consumer)
	go handleNotifications2(consumer)
	return consumer
}

func handleErrors2(consumer2 *cluster.Consumer) {
	for err := range consumer2.Errors() {
		log.Printf("Error: %s\n", err.Error())
	}
}

func handleNotifications2(consumer2 *cluster.Consumer) {
	for ntf := range consumer2.Notifications() {
		log.Printf("Rebalanced %+v\n", ntf)
	}
}

