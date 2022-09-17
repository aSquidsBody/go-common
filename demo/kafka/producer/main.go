package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/aSquidsBody/go-common/kafka"
	res "github.com/aSquidsBody/go-common/response"
)

var producer kafka.SyncProducer

func rootHandler(w http.ResponseWriter, r *http.Request) {
	msg := kafka.NewMessage("Test Value")
	err := producer.Send(msg)
	if err != nil {
		fmt.Println("Error occurred: ", err)
		return
	}
	res.WriteOk(w, "response")
}

func main() {
	brokers := []string{"kafka.kafka.svc.cluster.local:9092"}
	topic := "test.topic"

	var err error
	config := kafka.Config().ForSync(topic)
	producer, err = kafka.NewSyncProducer(brokers, config)
	if err != nil {
		log.Fatal("could not connect to kafka")
		return
	}

	fmt.Println("Listening on port 5000")
	log.Fatal(http.ListenAndServe(":5000", http.HandlerFunc(rootHandler)))
}
