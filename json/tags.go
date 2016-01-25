package json

import "github.com/gopherjs/gopherjs/js"

func parseTag(tag *js.Object) (*js.Object, *js.Object) {
	strLen := tag.Length()
	var val *js.Object
	for i := 0; i < strLen; i++ {
		start := i
		for i < strLen && tag.Call("charCodeAt", i).Int() == ' ' {
			i++
		}
		var char int
		for ; i < strLen; i++ {
			char = tag.Call("charCodeAt", i).Int()
			if char <= ' ' || char == ':' || char == '"' {
				break
			}
		}
		if i+1 >= strLen || char != ':' || tag.Call("charCodeAt", i+1).Int() {
			break
		}
		name := tag.Call("substring", start, i)
		i += 2
		start = i
		for ; i < strLen; i++ {
			char = tag.Call("charCodeAt", i).Int()
			if char == '\\' {
				i++
			} else if char == '"' {
				break
			}
		}
		if name.String() == "json" {
			val = tag.Call("substring", start, i)
			break
		}
	}
	if val == nil {
		return nil, nil
	}
	split := val.Call("indexOf", ",").Int()
	if split < 0 {
		return val, nil
	}
	return val.Call("substr", 0, split), val.Call("substr", split+1)
}

func tagContains(str *js.Object, option string) bool {
	if str == nil || str.Length() == 0 || str.Length() < len(option) {
		return false
	}
	if str.Length() == len(option) {
		return str.String() == option
	}
	return str.Call("substr", 0, len(option)+1) == option+"," || str.Call("substr", -1-len(option)) == ","+option || str.Call("includes", ","+option+",")
}
