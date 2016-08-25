// Package rpc is a client-only RPC package using websockets for javascript
package rpc

import (
	"github.com/MJKWoolnough/gopherjs/json"
	"github.com/gopherjs/gopherjs/js"
	"github.com/gopherjs/websocket"
)

// Call represents the necessary data for an RPC call
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

// Client is an RPC client
type Client struct {
	ws     *websocket.WebSocket
	nextID uint
	reqs   map[uint]func(json.RawMessage, error)
}

// Dial connects a websocket to the given address and creates the client
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
			go f(m, err)
		}
	})
	return &Client{
		ws:   w,
		reqs: reqs,
	}, nil
}

// Call make an RPC request to the given method name
func (c *Client) Call(method string, args interface{}, reply interface{}) error {
	call := <-c.Go(method, args, reply, make(chan *Call, 1)).Done
	return call.Error
}

// Close closes the websocket - it does not currently do anything special with
// outstanding requests
func (c *Client) Close() error {
	return c.ws.Close()
}

// Go makes an RPC request in a goroutine, returning a Call for the user to
// be notified upon completion and to retrieve any errors returned.
//
// If a nil done chan is given, one will be created
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
