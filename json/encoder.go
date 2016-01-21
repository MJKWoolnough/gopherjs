package json

import "github.com/gopherjs/gopherjs/js"

func Encode(v interface{}) string {
	run := true
	return js.Global.Get("JSON").Call("stringify", v, func(key, value *js.Object) *js.Object {
		if run {
			run = false
			filter(js.InternalObject(v).Get("constructor"), value)
		}
		return value
	}).String()
}

const (
	arrayKind  = 17
	mapKind    = 21
	ptrKind    = 22
	sliceKind  = 23
	structKind = 25
)

func filter(t, value *js.Object) {
	switch t.Get("kind").Int() {
	case arrayKind, sliceKind:
		len := value.Length()
		for i := 0; i < len; i++ {
			filter(t.Get("elem"), value.Index(i))
		}
	case mapKind:
		keys := js.Global.Get("Object").Call("keys", value)
		len := keys.Length()
		for i := 0; i < len; i++ {
			filter(t.Get("elem"), value.Get(keys.Index(i).String()))
		}
	case ptrKind:
		filter(t.Get("elem"), value)
	case structKind:
		nf := t.Get("fields").Length()
		for i := 0; i < nf; i++ {
			field := t.Get("fields").Index(i)
			tag := field.Get("tag").Call("replace", js.Global.Get("RegExp").New(".*json:\"([^\"]*)\".*"), "$1") // need to do this better
			if tag.String() == "-" {
				value.Delete(field.Get("name").String())
				continue
			}
			name := field.Get("name")
			filter(field.Get("typ"), value.Get(name.String()))
			if tag.Length() != 0 {
				p := tag.Call("split", ",", 2)
				if p.Index(0).Length() != 0 {
					if p.Index(0) != name {
						value.Set(p.Index(0).String(), value.Get(name.String()))
						value.Delete(name.String())
					}
				}
			}
		}
	}
}
