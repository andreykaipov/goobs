package general

import (
	"encoding/json"
	"fmt"

	"github.com/gorilla/websocket"
)

func (c *Client) GetVersion() (bytes []byte, err error) {
	if bytes, err = json.Marshal(
		map[string]string{
			"request-type": "GetSceneItemProperties",
			"message-id":   fmt.Sprintf("%v", 1),
			"item":         "Chat",
		},
	); err != nil {
		return nil, err
	}

	if err = c.conn.WriteMessage(websocket.TextMessage, bytes); err != nil {
		return nil, err
	}

	_, message, err := c.conn.ReadMessage()

	return message, nil
}
