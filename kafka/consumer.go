package kafka

import (
	"context"
	"fmt"
	"log"
	"sync"

	"github.com/Shopify/sarama"
)

type handler func(ConsumerMessage)

type Consumer interface {
	SetHandler(handler)
	Consume()
}

type consumer struct {
	ready         chan bool
	handler       handler
	config        *config
	consumerGroup sarama.ConsumerGroup
	group         string
}

func NewConsumer(group string, brokers []string, config *config) (Consumer, error) {
	consumerGroup, err := sarama.NewConsumerGroup(brokers, group, config.config)
	if err != nil {
		return nil, err
	}

	c := &consumer{
		ready:         make(chan bool),
		config:        config,
		handler:       nil,
		consumerGroup: consumerGroup,
		group:         group,
	}
	return c, nil
}

func (c *consumer) Consume() {
	ctx, cancel := context.WithCancel(context.Background())
	wg := &sync.WaitGroup{}
	wg.Add(1)
	go func() {
		defer wg.Done()
		for {
			// `Consume` should be called inside an infinite loop, when a
			// server-side rebalance happens, the consumer session will need to be
			// recreated to get the new claims
			if err := c.consumerGroup.Consume(ctx, c.config.consumerTopics, c); err != nil {
				log.Panicf("Error from the consumer: %v", err)
			}

			// check if context was cancelled, signalling that the consumer should stop
			if ctx.Err() != nil {
				return
			}
			c.ready = make(chan bool)
		}
	}()

	<-c.ready // wait until consumer has been setup
	fmt.Printf("Consumer in group %s listening to topics %s\n", c.group, c.config.consumerTopics)

	<-ctx.Done() // wait for context to terminate
	cancel()
	wg.Wait()
	if err := c.consumerGroup.Close(); err != nil {
		log.Panicf("Error closing client: %v", err)
	}
}

func (c *consumer) SetHandler(handlerFunc handler) {
	c.handler = handlerFunc
}

// Setup runs any setup code and then emits a ready signal to the consumer.
func (c *consumer) Setup(sarama.ConsumerGroupSession) error {

	// Setup code...

	close(c.ready)
	return nil
}

// Cleanup runs any cleanup code
func (c *consumer) Cleanup(sarama.ConsumerGroupSession) error {

	// Cleanup code...

	return nil
}

// ConsumeClaim continually listens to messages that come through the consumer claim
func (c *consumer) ConsumeClaim(session sarama.ConsumerGroupSession, claim sarama.ConsumerGroupClaim) error {
	for {
		select {
		case m := <-claim.Messages():
			message := ConsumerMessage{m}
			c.handler(message)
			session.MarkMessage(m, fmt.Sprintf("Consumed topic = %s, partition = %d, timestamp = %s", message.Topic, message.Partition, message.Timestamp))

			// should return when `session.Context()` is done.
			// If not, will raise `ErrRebalanceInProgress` or `read tcp <ip>:<port>: i/o timeout` when kafka re-balance
		case <-session.Context().Done():
			return nil
		}
	}
}
