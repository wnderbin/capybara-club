package nats_client

import (
	"github.com/nats-io/nats.go"
)

type Client struct {
	Conn *nats.Conn
}

func New(url string) (*Client, error) {
	nc, err := nats.Connect(url)
	if err != nil {
		return nil, err
	}
	return &Client{Conn: nc}, nil
}

func (c *Client) Close() {
	c.Conn.Close()
}
