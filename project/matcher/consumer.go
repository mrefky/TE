package main
import (

	"log"
	"time"
	"github.com/Shopify/sarama"
	cluster "github.com/bsm/sarama-cluster"

)
func createConsumer1() *cluster.Consumer {
	// define the configuration for our cluster
	config := cluster.NewConfig()
	config.Consumer.Return.Errors = false
	config.Group.Return.Notifications = false
	config.Consumer.Offsets.Initial = sarama.OffsetOldest // earliest uncommited offset
	config.Consumer.Offsets.CommitInterval = time.Second

	orderTopic := []string{"pre-orders"}

	log.Println("Listening for orders on topic -> ", orderTopic)
	// create the consumer
	consumer, err := cluster.NewConsumer(
		[]string{Kafka},
		"Matcher-pre-Orders",
		orderTopic,
		config,
	)

	if err != nil {
		log.Fatal("Unable to connect to kafka cluster")
	}

	go handleErrors(consumer)
	go handleNotifications(consumer)
	return consumer
}

 func createConsumer() *cluster.Consumer {
	// define the configuration for our cluster
	config := cluster.NewConfig()
	config.Consumer.Return.Errors = false
	config.Group.Return.Notifications = false
	config.Consumer.Offsets.Initial = sarama.OffsetOldest // earliest uncommited offset
	config.Consumer.Offsets.CommitInterval = time.Second

	orderTopic := []string{"orders"}

	log.Println("Listening for orders on topic -> ", orderTopic)
	// create the consumer
	consumer, err := cluster.NewConsumer(
		[]string{Kafka},
		"Matcher-Orders",
		orderTopic,
		config,
	)

	if err != nil {
		log.Fatal("Unable to connect to kafka cluster")
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