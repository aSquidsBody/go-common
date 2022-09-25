package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/aSquidsBody/go-common/kafka"
	res "github.com/aSquidsBody/go-common/response"
)

var consumer kafka.Consumer
var messages []string

func rootHandler(w http.ResponseWriter, r *http.Request) {
	res.WriteOk(w, messages)
}

func main() {
	brokers := []string{"kafka.kafka.svc.cluster.local:9092"}
	topics := []string{"test.topic"}

	var err error
	config := kafka.Config().ForConsumer("3.2.1", topics)
	consumer, err = kafka.NewConsumer("test-group", brokers, config)
	if err != nil {
		log.Fatal("could not create consumer ", err)
	}

	consumer.SetHandler(func(msg kafka.ConsumerMessage) {
		var value string
		err := msg.Unmarshal(&value)
		if err != nil {
			fmt.Println("An error occurred unmarshalling", err)
			return
		}
		messages = append(messages, value)
	})
	consumer.Consume()
}
