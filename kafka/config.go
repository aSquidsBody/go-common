package kafka

import (
	"time"

	"github.com/Shopify/sarama"
)

type config struct {
	config *sarama.Config
	topic  string
}

func (c *config) ForSync() *config {
	c.config.Producer.Retry.Max = 10
	c.config.Producer.RequiredAcks = sarama.WaitForAll // wait for all in-sync replicas to ack the message
	c.config.Producer.Return.Successes = true
	return c
}

func (c *config) ForAccessLog() *config {
	c.config.Producer.RequiredAcks = sarama.WaitForLocal       // only wait for the leader to ack
	c.config.Producer.Compression = sarama.CompressionSnappy   // compress messages
	c.config.Producer.Flush.Frequency = 500 * time.Millisecond // flush batches every 500ms
	return c
}

func Config(topic string) *config {
	return &config{sarama.NewConfig(), topic}
}
