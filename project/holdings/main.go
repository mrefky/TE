package main
import (
	"fmt"
	"github.com/confluentinc/confluent-kafka-go/v2/kafka"
	
	"encoding/json"
	//"github.com/Shopify/sarama"
	//cluster "github.com/bsm/sarama-cluster"
	//"log"
	"time"
//	"math"
//	"math/rand"
)


type Holding struct {
Secid string  `json:"Secid"`
Tradingaccount string `json:"TradACC"`
Custodian string `json:"Custodian"`
UserID string `json:"User"`
Quantity int32 `json:"qty"`
}



func main() {

p, err := kafka.NewProducer(&kafka.ConfigMap{"bootstrap.servers": "192.168.1.70:9094"})
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
var hol1 Holding
hol1.Secid="EGMFC001MM19"
hol1.Custodian="4501"
hol1.Tradingaccount="10100400"
hol1.Quantity=1000

	// Produce messages to topic (asynchronously)
	topic := "Holdings"
		//min:=1
		//max:=100
		//r := min + rand.Float64() * (max - min)
		//}
		msg:=hol1.toJSON()
		fmt.Println(string(msg[:]))

		p.Produce(&kafka.Message{
			TopicPartition: kafka.TopicPartition{Topic: &topic, Partition: kafka.PartitionAny},
			Value:   msg,
			Key: []byte(hol1.Secid), 
		}, nil)
	 //if (i % 10 == 0){
	time.Sleep(1 * time.Second)	
//	}



	// Wait for message deliveries before shutting down
	p.Flush(15 * 1000)
}
// Convert order to json from order struct
func (holding *Holding) toJSON() []byte {
	str, _ := json.Marshal(holding)
	return str
}
