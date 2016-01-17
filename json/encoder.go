package json

import (
	"reflect"

	"github.com/gopherjs/gopherjs/js"
)

func Encode(v interface{}) string {
	return js.Global.Get("JSON").Get("stringify").Invoke(processType(reflect.ValueOf(v))).String()
}

func processStruct(v reflect.Value) *js.Object {
	t := v.Type()
	o := js.Global.Get("Object").New()
	for i := 0; i < t.NumField; i++ {
		f := t.Field(i)
		if f.PkgPath == "" { // exported
			name := f.Name
			tag := f.Tag.Get("json")
			if tag != "" {
				if tag == "-" {
					continue
				}
				// set name and check tag options
			}
			o.Set(name, processType(v.Field(i)))
		}
	}
	return o
}

func processSlice(v reflect.Value) *js.Object {
	a := js.Global.Get("Array").New(v.Len())
	for i := 0; i < v.Len(); i++ {
		a.SetIndex(i, processType(v.Index(i)))
	}
	return a
}

func processType(v reflect.Value) *js.Object {
	for v.Kind() == reflect.Ptr {
		if v.IsNil() {
			return nil
		}
		v = v.Elem()
	}
	if v.IsNil() {
		return nil
	}
	switch v.Kind() {
	case reflect.Bool:
		return js.Global.Get("Bool").Invoke(v.Interface())
	case reflect.Uint8, reflect.Uint16, reflect.Uint32, reflect.Uint64, reflect.Uint, reflect.Uintptr, reflect.Int8, reflect.Int16, reflect.Int32, reflect.Int64, reflect.Int, reflect.Float32, reflect.Float64:
		return js.Global.Get("Number").Invoke(v.Interface())
	case reflect.String:
		return js.Global.Get("String").Invoke(v.Interface())
	case reflect.Struct:
		return processStruct(v)
	case reflect.Array, reflect.Slice:
		return processSlice(v)
	default:
		return nil
	}
}
