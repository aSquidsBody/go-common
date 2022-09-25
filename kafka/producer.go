package kafka

import (
	"fmt"
	"net/http"
	"time"

	"github.com/Shopify/sarama"
)

type SyncProducer interface {
	Send(msg *ProducerMessage) error
}

type syncProducer struct {
	Config   *config
	producer *sarama.SyncProducer
}

func (p *syncProducer) Send(msg *ProducerMessage) error {
	partition, offset, err := (*p.producer).SendMessage(&sarama.ProducerMessage{
		Topic: p.Config.producerTopic,
		Value: msg,
	})

	if err != nil {
		return fmt.Errorf("Failed to send Kafka message: %w", err)
	}

	fmt.Printf("Topic %s: Message published on partition=%d, offset=%d\n", p.Config.producerTopic, partition, offset)
	return nil
}

func NewSyncProducer(brokers []string, config *config) (SyncProducer, error) {

	producer, err := sarama.NewSyncProducer(brokers, config.config)
	if err != nil {
		return nil, fmt.Errorf("Failed to create Sarama SyncProducer: %w", err)
	}

	return &syncProducer{Config: config, producer: &producer}, nil
}

type AsyncProducer interface {
	Send(msg *ProducerMessage)
}

type asyncProducer struct {
	Config        *config
	producer      *sarama.AsyncProducer
	errorCallback func(err error)
}

func (ap *asyncProducer) Send(msg *ProducerMessage) {
	(*ap.producer).Input() <- &sarama.ProducerMessage{
		Topic: ap.Config.producerTopic,
		Value: msg,
	}
}

func newAsyncProducer(brokers []string, config *config, errorCallback func(err error)) (*asyncProducer, error) {
	producer, err := sarama.NewAsyncProducer(brokers, config.config)
	if err != nil {
		return nil, fmt.Errorf("Failed to create Sarama AsyncProducer: %w", err)
	}

	ap := &asyncProducer{Config: config, producer: &producer, errorCallback: errorCallback}
	go func() {
		for err := range (*ap.producer).Errors() {
			ap.errorCallback(err)
		}
	}()

	return ap, nil
}

func NewAsyncProducer(brokers []string, config *config, errorCallback func(err error)) (AsyncProducer, error) {
	return newAsyncProducer(brokers, config, errorCallback)
}

type AccessLogger interface {
	HandlerFunc(next http.Handler) http.Handler
}

type accessLogger struct {
	*asyncProducer
}
type accessLogEntry struct {
	Method       string  `json:"method"`
	Host         string  `json:"host"`
	Path         string  `json:"path"`
	IP           string  `json:"ip"`
	ResponseTime float64 `json:"response_time"`
}

func (al *accessLogger) HandlerFunc(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {

		started := time.Now()
		next.ServeHTTP(w, r)

		entry := accessLogEntry{
			Method:       r.Method,
			Host:         r.Host,
			Path:         r.RequestURI,
			IP:           r.RemoteAddr,
			ResponseTime: float64(time.Since(started)) / float64(time.Second),
		}

		(*al.producer).Input() <- &sarama.ProducerMessage{
			Topic: al.Config.producerTopic,
			Key:   sarama.StringEncoder(r.RemoteAddr),
			Value: NewMessage(entry),
		}
	})
}

func NewAccessLogger(brokers []string, config *config, errorCallback func(err error)) (AccessLogger, error) {
	ap, err := newAsyncProducer(brokers, config, errorCallback)
	if err != nil {
		return nil, err
	}

	return &accessLogger{ap}, nil
}
