package json

import (
	"reflect"

	"github.com/gopherjs/gopherjs/js"
)

type objectInfo struct {
	v reflect.Value
	c int
}

func Encode(v interface{}) string {
	objects := make([]objectInfo, 1, 16) // worth working out actual maximum depth?
	objects[0] = objectInfo{v: reflect.ValueOf(v)}
	return js.Global.Get("JSON").Get("stringify").Invoke(v, func(key, value *js.Object) *js.Object {
		v := objects[len(objects)-1].v
		objects[len(objects)-1].c--
		if objects[len(objects)-1].c == 0 {
			objects = objects[:len(objects)-1]
		}
		switch k := v.Kind(); k {
		case reflect.Struct:
			v = v.FieldByName(key.String())
		case reflect.Array, reflect.Slice:
			v = v.Index(key.Int())
		case reflect.Map:
			v = v.MapIndex(reflect.ValueOf(key.String()))
		}
		switch k := v.Kind(); k {
		case reflect.Struct:
			c := 0
			nf := v.NumField()
			t := v.Type()
			for i := 0; i < nf; i++ {
				field := t.Field(i)
				if field.PkgPath == "" {
					tag := field.Tag.Get("json")
					if tag == "-" {
						value.Set(field.Name, js.Undefined)
						value.Delete(field.Name)
						continue
					}
					c++
				}
			}
			objects = append(objects, objectInfo{v, c})
		case reflect.Array, reflect.Slice:
			objects = append(objects, objectInfo{v, v.Len()})
		case reflect.Map:
			objects = append(objects, objectInfo{v, v.Len()})
		}
		return value
	}).String()
}
