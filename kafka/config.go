package kafka

import (
	"log"
	"time"

	"github.com/Shopify/sarama"
)

type config struct {
	config         *sarama.Config
	producerTopic  string
	consumerTopics []string
}

func (c *config) ForSync(topic string) *config {
	c.producerTopic = topic
	c.config.Producer.Retry.Max = 10
	c.config.Producer.RequiredAcks = sarama.WaitForAll // wait for all in-sync replicas to ack the message
	c.config.Producer.Return.Successes = true
	return c
}

func (c *config) ForAccessLog(topic string) *config {
	c.producerTopic = topic
	c.config.Producer.RequiredAcks = sarama.WaitForLocal       // only wait for the leader to ack
	c.config.Producer.Compression = sarama.CompressionSnappy   // compress messages
	c.config.Producer.Flush.Frequency = 500 * time.Millisecond // flush batches every 500ms
	return c
}

func (c *config) ForConsumer(topics []string, group string) *config {
	c.consumerTopics = topics

	// if verbose
	// sarama.Logger = log.New(os.Stdout, "[sarama]", log.Lstdflags)

	version, err := sarama.ParseKafkaVersion("2.1.1")
	if err != nil {
		log.Panicf("Error parsing kafka version: %v", err)
	}

	c.config.Version = version
	c.config.Consumer.Group.Rebalance.Strategy = sarama.BalanceStrategyRoundRobin
	c.config.Consumer.Offsets.Initial = sarama.OffsetOldest
	return c
}

func Config() *config {
	return &config{config: sarama.NewConfig()}
}
