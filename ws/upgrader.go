package ws

import (
	"net/http"

	"github.com/aSquidsBody/go-common/logs"
	"github.com/gorilla/websocket"
)

var upgrader = websocket.Upgrader{
	ReadBufferSize:  1024,
	WriteBufferSize: 1024,
	CheckOrigin: func(r *http.Request) bool {
		// TODO: check origin
		return true
	},
}

func UpgradeToWebSocket(w http.ResponseWriter, r *http.Request) (*WsClient, error) {
	conn, err := upgrader.Upgrade(w, r, nil)
	if err != nil {
		logs.Errorf(err, "Could not upgrade connection to websocket")
		return nil, err
	}

	client := NewWsClient(conn)
	client.Enable()
	return client, nil
}
