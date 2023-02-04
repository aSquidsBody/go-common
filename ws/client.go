package ws

import (
	"bytes"
	"fmt"
	"time"

	"github.com/aSquidsBody/go-common/logs"
	"github.com/gorilla/websocket"
)

type WsClient struct {
	// websocket connection
	conn *websocket.Conn

	// Buffered channel to send the messages to the end-user
	send chan *WsMessage

	// Buffered channel to receive messages from the end user
	receive chan *WsMessage
}

// Write a message to the end-user
func (wc *WsClient) Write(msg *WsMessage) {
	wc.send <- msg
}

// Messages received from the end user
func (wc *WsClient) Messages() <-chan *WsMessage {
	return wc.receive
}

func NewWsClient(conn *websocket.Conn) *WsClient {
	wc := &WsClient{
		conn, make(chan *WsMessage, 256), make(chan *WsMessage, 256),
	}
	return wc
}

// Enable will start the readloop and write loop for the client
func (wc *WsClient) Enable() {
	go readloop(wc)
	go writeloop(wc)
}

// Close the connection with the end user
func (wc *WsClient) Close() {
	// close the conn
	wc.conn.Close()
}

// Read pong message from the end user
// Read any other data from the end user
func readloop(wc *WsClient) {
	defer wc.Close()

	wc.conn.SetReadLimit(maxMessageSize)
	wc.conn.SetReadDeadline(time.Now().Add(pongWait))
	wc.conn.SetPongHandler(func(string) error { wc.conn.SetReadDeadline(time.Now().Add(pongWait)); return nil })

	// infinite loop to read the messages from end user
	for {
		messageType, message, err := wc.conn.ReadMessage()
		// if error, then stop the loop
		if err != nil {
			if websocket.IsUnexpectedCloseError(err, websocket.CloseGoingAway, websocket.CloseAbnormalClosure) {
				fmt.Printf("Error cause closure of websocket. Error = %s.\n", err.Error())
			}
			break
		}
		if messageType == PongMessage {
			logs.Info("Client pong")
			continue
		}

		if messageType == TextMessage {
			message = bytes.TrimSpace(bytes.Replace(message, []byte{'\n'}, []byte{' '}, -1))
		}

		fmt.Printf("Received message on client. Message = '%s'\n", message)
		wc.receive <- &WsMessage{
			messageType: messageType,
			Content:     message,
		}
	}
}

// write pong messages
// write any other data to the end user
func writeloop(wc *WsClient) {
	// regular ping interval ticker
	pingTicker := time.NewTicker(pingPeriod)

	// close the connection and stop the ping ticker when done
	defer func() {
		pingTicker.Stop()
		wc.Close()
	}()

	// infinite loop to write the messages to the end user
	for {
		select {
		// case that there is info to send to the user
		case message, ok := <-wc.send:
			wc.conn.SetWriteDeadline(time.Now().Add(writeWait))

			// if the send channel was closed, then tell the client that the socket is closed and stop the loop
			if !ok {
				wc.conn.WriteMessage(websocket.CloseMessage, []byte{})
				return
			}

			var err error
			if message.messageType == TextMessage {
				err = wc.conn.WriteMessage(TextMessage, []byte(message.Content.(string)))
			} else {
				err = wc.conn.WriteMessage(BytesMessage, message.Content.([]byte))
			}

			if err != nil {
				logs.Errorf(err, "Could not write message to client")
				return
			}

		// case that it is time to send a new ping
		case <-pingTicker.C:
			wc.conn.SetWriteDeadline(time.Now().Add(writeWait)) // update the write deadline (can't have a write take too long)
			if err := wc.conn.WriteMessage(websocket.PingMessage, nil); err != nil {
				return // if failure during ping, then stop the connection
			}
		}
	}
}
