package ws

import (
	"time"
)

const (
	// Time allowed to write the data to the client
	writeWait = 10 * time.Second

	// Time allowed to read the next pong message from the client
	pongWait = 60 * time.Second

	// Send pings to the client with this period. Must be less than pongWait
	pingPeriod = pongWait * 9 / 10

	// Max message size receive from user
	maxMessageSize = 4096
)
