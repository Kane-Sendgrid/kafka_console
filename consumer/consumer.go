package main

import (
	"fmt"
	"github.com/Shopify/sarama"
	"github.com/kelseyhightower/envconfig"
	"log"
)

type configuration struct {
	Topic     string
	Broker    string
	Partition int32
}

var appConfig configuration

func main() {
	err := envconfig.Process("KC", &appConfig)
	if err != nil {
		log.Fatal(err.Error())
	}
	fmt.Println(appConfig)
	client, err := sarama.NewClient("my_client", []string{appConfig.Broker}, nil)

	if err != nil {
		panic(err)
	} else {
		fmt.Println("> connected")
	}
	defer client.Close()

	consumerConfig := sarama.NewConsumerConfig()
	consumerConfig.OffsetMethod = sarama.OffsetMethodNewest
	consumer, err := sarama.NewConsumer(client, appConfig.Topic,
		appConfig.Partition, "my_consumer_group", consumerConfig)
	if err != nil {
		panic(err)
	} else {
		fmt.Println("> consumer ready")
	}
	defer consumer.Close()

	var event *sarama.ConsumerEvent
	for {
		select {
		case event = <-consumer.Events():
			if event.Err != nil {
				panic(event.Err)
			}
			fmt.Println("got message:", string(event.Value))
		}
	}
}
