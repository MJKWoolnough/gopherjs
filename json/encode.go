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
	boolKind       = 1
	intKind        = 2
	int8Kind       = 3
	int16Kind      = 4
	int32Kind      = 5
	int64Kind      = 6
	uintKind       = 7
	uint8Kind      = 8
	uint16Kind     = 9
	uint32Kind     = 10
	uint64Kind     = 11
	uintptrKind    = 12
	float32Kind    = 13
	float64Kind    = 14
	complex64Kind  = 15
	complex128Kind = 16
	arrayKind      = 17
	interfaceKind  = 20
	mapKind        = 21
	ptrKind        = 22
	sliceKind      = 23
	stringKind     = 24
	structKind     = 25
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
		keys := js.Global.Get("Object").Call("keys", v)
		len := keys.Length()
		for i := 0; i < len; i++ {
			o := v.Get(keys.Index(i).String())
			filter(o.Get("v"), t.Get("elem"), value.Get(o.Get("k").String()))
		}
	case ptrKind:
		filter(v, t.Get("elem"), value)
	case structKind:
		nf := t.Get("fields").Length()
		for i := 0; i < nf; i++ {
			field := t.Get("fields").Index(i)
			fieldType := field.Get("typ")
			tag, tagOptions := parseTag(getJSONTag(field.Get("tag")))
			if tag == "-" {
				value.Delete(field.Get("name").String())
				continue
			}
			name := field.Get("name").String()
			if tagOptions.Contains("omitempty") {
				if isEmpty(v.Get(name), fieldType) {
					value.Delete(name)
					continue
				}
			}
			if tagOptions.Contains("string") && stringable(fieldType) {
				if fieldType.Get("kind").Int() == stringKind {
					value.Set(name, js.Global.Get("JSON").Call("stringify", v.Get(name)))
				} else {
					value.Set(name, value.Get(name).Call("toString"))
				}
			} else {
				filter(v.Get(name), fieldType, value.Get(name))
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

func isEmpty(v, t *js.Object) bool {
	switch t.Get("kind").Int() {
	case boolKind:
		return !v.Bool()
	case int8Kind, int16Kind, int32Kind, intKind, uint8Kind, uint16Kind, uint32Kind, uintKind:
		return v.Int() == 0
	case int64Kind, uint64Kind:
		return v.Int64() == 0
	case float32Kind, float64Kind:
		return v.Float() == 0
	case arrayKind, sliceKind, stringKind:
		return v.Length() == 0
	case ptrKind:
		if v == t.Get("nil") {
			return true
		}
		return isEmpty(v.Call("$get"), t.Get("elem"))
	case interfaceKind:
		if v.Get("$val") == js.Undefined {
			return true
		}
		return isEmpty(v.Get("$val"), v.Get("constructor"))
	case mapKind:
		return js.Global.Get("Object").Call("keys", v).Length() == 0
	}
	return false
}

func stringable(t *js.Object) bool {
	switch t.Get("kind").Int() {
	case boolKind, int8Kind, int16Kind, int32Kind, intKind, uint8Kind, uint16Kind, uint32Kind, uintKind, int64Kind, uint64Kind, float32Kind, float64Kind, stringKind:
		return true
	}
	return false
}
