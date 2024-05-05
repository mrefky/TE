package main
import (
	"github.com/Shopify/sarama"
	//	cluster "github.com/bsm/sarama-cluster"
	"log"
	"fmt"

)
func createProducer() sarama.AsyncProducer {
	config := sarama.NewConfig()
	config.Producer.Return.Successes = false         // fire and forget
	config.Producer.Return.Errors = true             // notify on failed
	config.Producer.RequiredAcks = sarama.WaitForAll // waits for all insync replicas to commit

	producer, err := sarama.NewAsyncProducer([]string{Kafka}, config)

	if err != nil {
		log.Fatal("Unable to connect producer to kafka server")
	}
	fmt.Println("<<<<<<<<<<<<<<<<<<<<<<<<<<<<<<<<<<<<<<<<<<<<<<<<<<<<<<<<<<<<<<<<")
log.Printf("Connected to Kafka to publish trades")
fmt.Println("<<<<<<<<<<<<<<<<<<<<<<<<<<<<<<<<<<<<<<<<<<<<<<<<<<<<<<<<<<<<<<<<")
	return producer

}
