package kafka

import "encoding/json"

type message struct {
	content interface{}

	encoded []byte
	err     error
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
