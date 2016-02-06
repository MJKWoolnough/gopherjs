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
	str := js.Global.Get("JSON").Call("stringify", js.InternalObject(jsonify).Invoke(js.InternalObject(v)), nil, indent)
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

func jsonify(v *js.Object) *js.Object {
	t := v.Get("constructor")
	for t.Get("kind").Int() == ptrKind {
		if v.Get("$val") != t.Get("nil") {
			v = v.Call("$get")
			t = t.Get("elem")
		}
	}
	switch t.Get("kind").Int() {
	case boolKind, intKind, int8Kind, int16Kind, int32Kind, uintKind, uint8Kind, uint16Kind, uint32Kind, uintptrKind, float32Kind, float64Kind, int64Kind, uint64Kind, complex64Kind, complex128Kind, ptrKind, stringKind:
		return js.Global.Call("$externalize", v, t)
	case arrayKind, sliceKind:
		l := v.Get("$length").Int()
		a := js.Global.Get("Array").New(l)
		for i := 0; i < l; i++ {
			a.SetIndex(i, js.InternalObject(jsonify).Invoke(v.Get("$array").Index(i)))
		}
		return a
	case interfaceKind:
		return js.InternalObject(jsonify).Invoke(v.Get("$val"))
	case mapKind:
		m := js.Global.Get("Object").New()
		keys := js.Global.Get("Object").Call("keys", v)
		len := keys.Length()
		for i := 0; i < len; i++ {
			k := keys.Index(i).String()
			m.Set(k, js.InternalObject(jsonify).Invoke(v.Get(k)))
		}
		return m
	case structKind:
		s := js.Global.Get("Object").New()
		todo := [][2]*js.Object{{v, t}}
		fieldsTodo := make(map[string]*js.Object)
		for _, val := range todo {
			v := val[0]
			t := val[1]
			fields := t.Get("fields")
			fieldsLen := fields.Length()
			for i := 0; i < fieldsLen; i++ {
				f := fields.Index(i)
				if f.Get("pkg").Length() > 0 {
					continue
				}
				fName := f.Get("prop").String()
				name := f.Get("name").String()
				anon := false
				if name == "" {
					name = fName // filter $x?
					anon = true
				}
				n, o := parseTag(getJSONTag(f.Get("tag")))
				if n == "-" || (o.Contains("omitempty") && isEmpty(v.Get(fName), f.Get("typ"))) {
					continue
				}
				jName := name
				if n != "" {
					jName = n
				}
				if _, ok := fieldsTodo[jName]; ok {
					fieldsTodo[jName] = nil
				} else if anon {
					todo = append(todo, [2]*js.Object{v.Get(fName), f.Get("typ")})
				} else if s.Get(jName) == js.Undefined {
					val := v.Get(fName)
					if val.Get("constructor").Get("kind") == js.Undefined {
						t := js.Global.Get("Object").New()
						t.Set("kind", f.Get("kind"))
						val.Set("constructor", t)
					}
					fieldsTodo[jName] = val
				}
			}
			for name, f := range fieldsTodo {
				if f != nil {
					s.Set(name, js.InternalObject(jsonify).Invoke(f))
					delete(fieldsTodo, name)
				}
			}
		}
		return s
	}
	return js.Undefined
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
		return v == t.Get("nil")
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

func isMarshaler(t *js.Object) bool {
	methodNum := t.Get("methods").Length()
	for i := 0; i < methodNum; i++ {
		method := t.Get("methods").Index(i)
		if method.Get("name").String() == "MarshalJSON" {
			mt := method.Get("typ")
			return mt.Get("params").Length() == 0 && mt.Get("results").Length() == 2 && mt.Get("results").Index(0).Get("string").String() == "[]uint8" && mt.Get("results").Index(1).Get("string").String() == "error"
		}
	}
	return false
}
