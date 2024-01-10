package websocket

import "github.com/gorilla/websocket"

type Client struct {
	Conn   *websocket.Conn
	UserId int64 `json:"user_id"`
}

func CreateClient(conn *websocket.Conn, id int64) *Client {
	return &Client{
		Conn:   conn,
		UserId: id,
	}
}
