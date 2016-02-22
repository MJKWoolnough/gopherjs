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

func Dial(addr string) (*Client, error) {
	w, err := websocket.New(addr)
	if err != nil {
		return nil, err
	}
	reqs := make(map[uint]func(json.RawMessage, error))
	w.AddEventListener("message", false, func(e *js.Object) {
		var (
			r request
			m json.RawMessage
		)
		r.Params[0] = &m
		err := json.UnmarshalString(e.Get("data").String(), &r)
		f, ok := reqs[r.ID]
		if ok {
			delete(reqs, r.ID)
			f(m, err)
		}
	})
	return &Client{
		ws:   w,
		reqs: reqs,
	}, nil
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
		if cap(done) < 1 {
			panic("invalid channel capacity")
		}
		call.Done = done
	}
	str, err := json.MarshalString(request{
		Method: method,
		ID:     c.nextID,
		Params: [1]interface{}{args},
	})
	if err == nil {
		err = c.ws.Send(str)
	}
	if err != nil {
		call.Error = err
		call.Done <- call
		return call
	}
	c.reqs[c.nextID] = func(rm json.RawMessage, err error) {
		if err != nil {
			call.Error = err
		} else if err = json.Unmarshal(rm, reply); err != nil {
			call.Error = err
		}
		call.Done <- call
	}
	c.nextID++
	return call
}
