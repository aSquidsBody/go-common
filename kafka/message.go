package kafka

import (
	"encoding/json"

	"github.com/Shopify/sarama"
)

const (
	CREATED = "created"
	UPDATED = "updated"
	DELETED = "deleted"
)

type ProducerMessage struct {
	content interface{}

	encoded []byte
	err     error
}
type ConsumerMessage struct {
	*sarama.ConsumerMessage
}

type EventMessage struct {
	EventType string `json:"event_type"`
}

func (m *ConsumerMessage) Unmarshal(v interface{}) error {
	return json.Unmarshal(m.Value, v)
}

func (m *ProducerMessage) ensureEncoded() {
	if m.encoded == nil && m.err == nil {
		m.encoded, m.err = json.Marshal(m.content)
	}
}

func (m *ProducerMessage) Length() int {
	m.ensureEncoded()
	return len(m.encoded)
}

func (m *ProducerMessage) Encode() ([]byte, error) {
	m.ensureEncoded()
	return m.encoded, m.err
}

func NewMessage(content interface{}) *ProducerMessage {
	return &ProducerMessage{
		content: content,
		encoded: nil,
		err:     nil,
	}
}
