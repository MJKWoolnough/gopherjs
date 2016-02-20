package rpc

import (
	"encoding/json"

	"github.com/gopherjs/websocket"
)

type Call struct {
	ServiceMethod string
	Args, Reply   interface{}
	Error         error
	Done          chan *Call
}

type request struct {
	Method string         `json:"method"`
	ID     uint           `json:"id"`
	Params [1]interface{} `json:"params"`
}

type Client struct {
	ws     *websocket.Websocket
	nextID uint
	reqs   map[uint]func(json.RawMessage, error)
}

func NewClient(addr string) *Client {
	w, err := websocket.New(addr)

	return &Client{
		ws:   w,
		reqs: make(map[uint]func(json.RawMessage, error)),
	}
}

func (c *Client) Call(method string, args interface{}, reply interface{}) error {
	call := <-c.Go(method, args, reply, make(chan *Call, 1)).Done
	return call.Error
}

func (c *Client) Close() error {
	return c.w.Close()
}

func (c *Client) Go(method string, args interface{}, reply interface{}, done chan *Call) *Call {
	call := &Call{
		ServiceMethod: method,
		Args:          args,
		Reply:         reply,
	}
	if done == nil {
		call.Done = make(chan *Call, 1)
	} else {
		call.Done = done
	}
	c.nextID++
	id := c.nextID
	c.reqs[id] = func(rm json.RawMessage, err error) {
		if err != nil {
			call.Error = err
		} else if err = json.Unmarshal(rm, reply); err != nil {
			call.Error = err
		}
		call.Done <- call
	}
	return call
}
