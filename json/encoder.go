package json

import (
	"reflect"
	"strings"

	"github.com/gopherjs/gopherjs/js"
)

func Encode(v interface{}) string {
	run := true
	t := reflect.TypeOf(v)
	return js.Global.Get("JSON").Call("stringify", v, func(key, value *js.Object) *js.Object {
		if run {
			run = false
			filter(t, value)
		}
		return value
	}).String()
}

func filter(t reflect.Type, value *js.Object) {
	switch t.Kind() {
	case reflect.Struct:
		nf := t.NumField()
		for i := 0; i < nf; i++ {
			field := t.Field(i)
			tag := field.Tag.Get("json")
			if tag == "-" {
				value.Delete(field.Name)
				continue
			}
			filter(field.Type, value.Get(field.Name))
			if tag != "" {
				p := strings.SplitN(tag, ",", 2)
				if p[0] != "" {
					if p[0] != field.Name {
						value.Set(p[0], value.Get(field.Name))
						value.Delete(field.Name)
					}
				}
			}

		}
	case reflect.Array, reflect.Slice:
		len := value.Length()
		for i := 0; i < len; i++ {
			filter(t.Elem(), value.Index(i))
		}
	case reflect.Map:
		keys := js.Global.Get("Object").Call("keys", value)
		len := keys.Length()
		for i := 0; i < len; i++ {
			filter(t.Elem(), value.Get(keys.Index(i).String()))
		}
	}
}
