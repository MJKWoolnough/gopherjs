package json

import "github.com/gopherjs/gopherjs/js"

type StringWriter interface {
	WriteString(string) (int, error)
}

func Compact(dst StringWriter, src []byte) error {
	return Indent(dst, src, "", "")
}

func Indent(dst StringWriter, src []byte, prefix, indent string) (err error) {
	defer func() {
		if e := recover(); e != nil {
			if er, ok := e.(*js.Error); ok && er != nil {
				err = er
			} else {
				panic(e)
			}
		}
	}()
	str := js.Global.Get("JSON").Call("stringify", js.Global.Get("JSON").Call("parse", string(src)), nil, indent)
	if len(prefix) > 0 {
		str = str.Call("replace", js.Global.Get("RegExp").New("\n", "g"), "\n"+prefix)
	}
	_, err = dst.WriteString(str.String())
	return err
}
