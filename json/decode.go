package json

import (
	"errors"

	"github.com/gopherjs/gopherjs/js"
)

func Unmarshal(data []byte, v interface{}) error {
	return UnmarshalString(string(data), v)
}

func UnmarshalString(data string, v interface{}) (err error) {
	defer func() {
		if e := recover(); e != nil {
			if er, ok := e.(*js.Error); ok && er != nil {
				err = er
			}
		}
	}()
	io := js.InternalObject(v)
	if io.Get("constructor").Get("kind").Int() != ptrKind { // necessary???
		return ErrPointerRequired
	}

	parsed := js.Global.Get("JSON").Call("parse", data)
	unjsonify(io, parsed)
	return nil
}

func unjsonify(v, parsed *js.Object) {
	t := v.Get("constructor")
	v, t = unwrap(v, t)
}

// Errors
var (
	ErrPointerRequired = errors.New("pointer required")
)
