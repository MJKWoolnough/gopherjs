package rpc

import "github.com/gopherjs/websocket"

type Client struct {
	nextID uint
}

func NewClient(addr string) *Client {
	w, err := websocket.Dial(addr)
}

type call struct {
	Method string         `json:"method"`
	ID     uint           `json:"id"`
	Params [1]interface{} `json:"params"`
}

func (c *Client) Call(method string, args interface{}, reply interface{}) error {
	return nil
}

func (c *Client) Close() error {
	return nil
}

func (c *Client) Go(method string, args interface{}, reply interface{}, done chan *Call) *Call {
	return nil
}
