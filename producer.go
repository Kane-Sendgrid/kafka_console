package main

import (
	"bufio"
	"fmt"
	"github.com/Shopify/sarama"
	"github.com/kelseyhightower/envconfig"
	"log"
	"os"
)

type configuration struct {
	Topic     string
	Broker    string
	Partition int32
}

type partitioner struct {}

func (self *partitioner) Partition(key sarama.Encoder, numPartitions int32) int32 {
	return appConfig.Partition
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

	producerConfig := sarama.NewProducerConfig()
	producerConfig.Partitioner = &partitioner{}
	producer, err := sarama.NewProducer(client, producerConfig)
	if err != nil {
		panic(err)
	}
	defer producer.Close()

	reader := bufio.NewReader(os.Stdin)
	for {
		line, err := reader.ReadString(byte('\n'))
		if err != nil {
			break
		}

		err = producer.SendMessage(appConfig.Topic, nil, sarama.StringEncoder(line))
		fmt.Printf("Sent message %v, error: %v\n", line, err)
	}
}
