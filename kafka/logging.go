package kafka

import (
	"log"
	"os"

	"github.com/Shopify/sarama"
)

func verbose() {
	sarama.Logger = log.New(os.Stdout, "[Sarama] ", log.LstdFlags)
}
