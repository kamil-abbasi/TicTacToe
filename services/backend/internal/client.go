package internal

import (
	"log"
	"sync"
	"time"

	"github.com/gorilla/websocket"
)

type Client struct {
	once  sync.Once
	name  string
	conn  *websocket.Conn
	write chan []byte
	read  chan []byte
}

func NewClient(name string, conn *websocket.Conn) *Client {
	return &Client{
		once:  sync.Once{},
		name:  name,
		conn:  conn,
		write: make(chan []byte),
		read:  make(chan []byte),
	}
}

func (c *Client) Name() string {
	return c.name
}

func (c *Client) Read() <-chan []byte {
	return c.read
}

func (c *Client) Write(message []byte) {
	c.write <- message
}

func (c *Client) Close() {
	c.once.Do(func() {
		c.conn.WriteControl(websocket.CloseMessage, websocket.FormatCloseMessage(websocket.CloseNormalClosure, ""), time.Now().Add(2*time.Second))
		c.conn.Close()
		close(c.read)
		close(c.write)
	})
}

func (c *Client) ReadPump() {
	defer func() {
		c.Close()
	}()

	for {
		_, message, err := c.conn.ReadMessage()

		if err != nil {
			if websocket.IsUnexpectedCloseError(
				err,
				websocket.CloseGoingAway,
				websocket.CloseNormalClosure,
				websocket.CloseNoStatusReceived,
			) {
				log.Println(err.Error())
			}
			break
		}

		c.read <- message
	}
}

func (c *Client) WritePump() {
	defer func() {
		c.Close()
	}()

	for message := range c.write {
		err := c.conn.WriteMessage(websocket.TextMessage, message)

		if err != nil {
			log.Println(err.Error())
			break
		}
	}
}
