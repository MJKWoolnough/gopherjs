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
			tag, tagOptions := parseTag(getJSONTag(field.Get("tag")))
			if tag == "-" {
				value.Delete(field.Get("name").String())
				continue
			}
			name := field.Get("name").String()
			if tagOptions.Contains("omitempty") {
				if isEmpty(v.Get(name)) {
					value.Delete(name)
					continue
				}
			}
			if tagOptions.Contains("string") {

			} else {
				filter(v.Get(name), field.Get("typ"), value.Get(name))
			}
			if len(tag) != 0 && tag != name {
				value.Set(tag, value.Get(name))
				value.Delete(name)
			}
		}
	case interfaceKind:
		filter(v.Get("$val"), v.Get("constructor"), value)
	}
}

func isEmpty(v *js.Object) bool {
	return false
}
