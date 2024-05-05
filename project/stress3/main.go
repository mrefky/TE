package main
import (
	"fmt"
	"github.com/confluentinc/confluent-kafka-go/v2/kafka"
	
	"encoding/json"
	//"github.com/Shopify/sarama"
	//cluster "github.com/bsm/sarama-cluster"
	//"log"
	"time"
	"math"
	"math/rand"
)
type Order struct {
	ID        int32 `json:"id"`
	Side      string  `json:"Side"`
	Quantity  int32  `json:"Quantity"`
	Price     float64  `json:"Price"`
	Timestamp int64 `json:"Timestamp"`
	Seccode   string `json:"Seccode"`
  	Custodian string `json:"Custodian"`
        HQty int32 `json:"HQty"`
        User string `json:"User"`
        TrdAcc string `json:"TrdAcc"`
        MsgType string `json:"MsgType"`
		TimeInForce string `json:"TimeInForce"`
		OrdType string `json:"OrdType"`


}
func randFloats(min, max float64, n int) []float64 {

    res := make([]float64, n)
    for i := range res {
        res[i] = min + rand.Float64() * (max - min)
    }
    return res
}
func randomBool() bool {
    rand.Seed(time.Now().UnixNano())
    return rand.Intn(2) == 1
}

func main() {

	p, err := kafka.NewProducer(&kafka.ConfigMap{"bootstrap.servers": "192.168.169.70:9094"})
	if err != nil {
		panic(err)
	}

	defer p.Close()

	// Delivery report handler for produced messages
	go func() {
		for e := range p.Events() {
			switch ev := e.(type) {
			case *kafka.Message:
				if ev.TopicPartition.Error != nil {
					fmt.Printf("Delivery failed: %v\n", ev.TopicPartition)
				} else {
					fmt.Printf("Delivered message to %v\n", ev.TopicPartition)
					
				}
			}
		}
	}()

	// Produce messages to topic (asynchronously)
	topic := "orders"
	fmt.Println("Enter number of required prders then press enter:")
	var ix int=0
	_,_= fmt.Scanf("%d", &ix)
	for  i:=1;i<ix+1;i++{
		var ord Order
		//min:=1
		//max:=100
		//r := min + rand.Float64() * (max - min)
		ord.ID=int32(1000+i)
		ord.Price= math.Round(100*randFloats(1,100,1)[0])/100
		ord.Quantity=int32(1+rand.Intn(1000))
		ord.Seccode="EGS21451C017"
		
		if randomBool() {
		ord.Side="sell"
		}else{
				ord.Side="buy"
		}
		ord.Timestamp=time.Now().UnixNano()
		//var msg []
		msg:=ord.toJSON()
		fmt.Println(string(msg[:]))

		p.Produce(&kafka.Message{
			TopicPartition: kafka.TopicPartition{Topic: &topic, Partition: kafka.PartitionAny},
			Value:   msg,
			Key: []byte(ord.Seccode), 
		}, nil)
if (i % 10 == 0){
//time.Sleep(1 * time.Second)	
}}

	// Wait for message deliveries before shutting down
	p.Flush(15 * 1000)
}
// Convert order to json from order struct
func (order *Order) toJSON() []byte {
	str, _ := json.Marshal(order)
	return str
}
