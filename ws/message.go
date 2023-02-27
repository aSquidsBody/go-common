package ws

import (
	"encoding/json"
	"fmt"

	"github.com/aSquidsBody/go-common/logs"
	"github.com/gorilla/websocket"
)

const (
	TextMessage  = websocket.TextMessage
	BytesMessage = websocket.BinaryMessage
	PongMessage  = websocket.PongMessage
)

var InvalidContent = fmt.Errorf("Unexpected WsMessage.Content type.")

func NewTextMessage(text string) *WsMessage {
	return &WsMessage{
		messageType: TextMessage,
		Content:     text,
	}
}

func NewJsonMessage(content interface{}) (*WsMessage, error) {
	data, err := json.Marshal(content)
	if err != nil {
		logs.Errorf(err, "Could not marshal WsMessage contents")
		return nil, err
	}

	return &WsMessage{
		messageType: TextMessage,
		Content:     string(data),
	}, nil
}

// WsMessage is a message that can be sent
// - from the client to the server
// - from the server to a client
// - from a server to all client
type WsMessage struct {
	// 1 for text, 2 for bytes
	messageType int
	Content     interface{} `json:"content"`
}

func (ws *WsMessage) IsText() bool {
	return ws.messageType == 1
}

func (ws *WsMessage) ContentText() (string, error) {
	s, ok := ws.Content.(string)
	if !ok {
		return "", InvalidContent
	}
	return s, nil
}

func (ws *WsMessage) ContentBytes() ([]byte, error) {
	b, ok := ws.Content.([]byte)
	if !ok {
		return nil, InvalidContent
	}
	return b, nil
}
