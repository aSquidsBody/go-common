package kafka

import (
	"encoding/json"

	"github.com/Shopify/sarama"
)

// message used in the producer
type message struct {
	content interface{}

	encoded []byte
	err     error
}

// Message is used in the consumer, exported so users can
// make use of it
type Message struct {
	*sarama.ConsumerMessage
}

func (m *Message) Unmarshal(v interface{}) error {
	return json.Unmarshal(m.Value, v)
}

func (m *message) ensureEncoded() {
	if m.encoded == nil && m.err == nil {
		m.encoded, m.err = json.Marshal(m.content)
	}
}

func (m *message) Length() int {
	m.ensureEncoded()
	return len(m.encoded)
}

func (m *message) Encode() ([]byte, error) {
	m.ensureEncoded()
	return m.encoded, m.err
}

func NewMessage(content interface{}) *message {
	return &message{
		content: content,
		encoded: nil,
		err:     nil,
	}
}
