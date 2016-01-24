package json

import "github.com/gopherjs/gopherjs/js"

func Marshal(v interface{}) ([]byte, error) {
	return MarshalIndent(v, "", "")
}

func MarshalIndent(v interface{}, prefix, indent string) ([]byte, error) {
	s, err := MarshalIndentString(v, prefix, indent)
	return []byte(s), err
}

func MarshalString(v interface{}) (string, error) {
	return MarshalIndentString(v, "", "")
}

func MarshalIndentString(v interface{}, prefix, indent string) (string, error) {
	str := js.Global.Get("JSON").Call("stringify", js.InternalObject(filterIt).Invoke(js.InternalObject(v), v), nil, indent)
	if len(prefix) > 0 {
		str = str.Call("replace", js.Global.Get("RegExp").New("\n", "g"), "\n"+prefix)
	}
	return str.String(), nil
}

const (
	arrayKind     = 17
	interfaceKind = 20
	mapKind       = 21
	ptrKind       = 22
	sliceKind     = 23
	structKind    = 25
)

func filterIt(v, value *js.Object) *js.Object {
	filter(v, v.Get("constructor"), value)
	return value
}

func filter(v, t, value *js.Object) {
	switch t.Get("kind").Int() {
	case arrayKind, sliceKind:
		len := value.Length()
		for i := 0; i < len; i++ {
			filter(v.Get("$array").Index(i), t.Get("elem"), value.Index(i))
		}
	case mapKind:
		keys := js.Global.Get("Object").Call("keys", value)
		len := keys.Length()
		for i := 0; i < len; i++ {
			name := keys.Index(i).String()
			filter(v.Get("$map").Get(name), t.Get("elem"), value.Get(name))
		}
	case ptrKind:
		filter(v, t.Get("elem"), value)
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
			filter(v.Get(name.String()), field.Get("typ"), value.Get(name.String()))
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
	case interfaceKind:
		filter(v.Get("$val"), v.Get("constructor"), value)
	}
}
