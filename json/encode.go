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

func MarshalIndentString(v interface{}, prefix, indent string) (json string, err error) {
	defer func() {
		if e := recover(); e != nil {
			if er, ok := e.(*js.Error); ok {
				err = er
			}
		}
	}()
	str := js.Global.Get("JSON").Call("stringify", js.InternalObject(jsonify).Invoke(js.InternalObject(v)), nil, indent)
	if len(prefix) > 0 {
		str = str.Call("replace", "\n", "\n"+prefix)
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

func unwrap(v, t *js.Object) (*js.Object, *js.Object) {
Loop:
	for {
		switch t.Get("kind").Int() {
		case ptrKind:
			if v.Get("$val") != t.Get("nil") {
				v = v.Call("$get")
				t = t.Get("elem")
			} else {
				break Loop
			}
		case interfaceKind:
			t = v.Get("constructor")
			v = v.Get("$val")
		default:
			break Loop
		}
	}
	return v, t
}

func wrap(v, t *js.Object) *js.Object {
	nt := js.Global.Get("Object").New()
	nt.Set("elem", t)
	nt.Set("kind", ptrKind)
	nv := js.Global.Get("Object").New()
	nv.Set("$val", v)
	nv.Set("constructor", nt)
	nv.Set("$get", js.MakeFunc(func(this *js.Object, _ []*js.Object) interface{} {
		return this.Get("$val")
	}))
	return nv
}

func jsonify(v *js.Object) *js.Object {
	t := v.Get("constructor")
	v, t = unwrap(v, t)
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
	case mapKind:
		m := js.Global.Get("Object").New()
		keys := js.Global.Get("Object").Call("keys", v)
		len := keys.Length()
		for i := 0; i < len; i++ {
			val := v.Get(keys.Index(i).String())
			m.Set(val.Get("k").String(), js.InternalObject(jsonify).Invoke(val.Get("v")))
		}
		return m
	case structKind:
		s := js.Global.Get("Object").New()
		nextLevel := [][2]*js.Object{{v, t}}
		fieldsTodo := make(map[string]*js.Object)
		for len(nextLevel) > 0 {
			level := nextLevel
			//nextLevel = nextLevel[:0]
			nextLevel = make([][2]*js.Object, 0)
			for len(level) > 0 {
				v := level[0][0]
				t := level[0][1]
				level = level[1:]
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
					if anon {
						nv, nt := unwrap(v.Get(fName), f.Get("typ"))
						if nv != nt.Get("nil") {
							nextLevel = append(nextLevel, [2]*js.Object{nv, nt})
						}
					} else if _, ok := fieldsTodo[jName]; ok {
						fieldsTodo[jName] = nil
					} else {
						val := v.Get(fName)
						vt := f.Get("typ")
						val, vt = unwrap(val, vt)
						if o.Contains("string") && stringable(vt) {
							s := js.Global.Get("Object").New()
							s.Set("kind", stringKind)
							if vt.Get("kind").Int() == stringKind {
								val = wrap(js.Global.Get("JSON").Call("stringify", val), s)
							} else {
								val = wrap(val.Call("toString"), s)
							}
						} else if val.Get("constructor").Get("kind") == js.Undefined {
							val = wrap(val, vt)
						}
						fieldsTodo[jName] = val
					}
				}
			}
			for name, f := range fieldsTodo {
				if f != nil {
					if isMarshaler(f) {
						tup := f.Call("MarshalJSON")
						if tup.Index(1) != js.Global.Get("$ifaceNil") {
							panic(tup.Index(1))
						}
						s.Set(name, js.Global.Get("JSON").Call("parse", js.Global.Call("$bytesToString", tup.Index(0)).String()))
					} else {
						s.Set(name, js.InternalObject(jsonify).Invoke(f))
					}
					fieldsTodo[name] = nil
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

func isMarshaler(v *js.Object) bool {
	_, t := unwrap(v, v.Get("constructor"))
	methods := t.Get("methods")
	if methods == js.Undefined {
		return false
	}
	methodNum := methods.Length()
	for i := 0; i < methodNum; i++ {
		method := methods.Index(i)
		if method.Get("name").String() == "MarshalJSON" {
			mt := method.Get("typ")
			return mt.Get("params").Length() == 0 && mt.Get("results").Length() == 2 && mt.Get("results").Index(0).Get("string").String() == "[]uint8" && mt.Get("results").Index(1).Get("string").String() == "error"
		}
	}
	return false
}
