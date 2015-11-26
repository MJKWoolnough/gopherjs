// Package xform provides some shortcut funcs for various form related activites
package xform

import (
	"strconv"

	"github.com/MJKWoolnough/gopherjs/style"
	"github.com/MJKWoolnough/gopherjs/xdom"
	"github.com/gopherjs/gopherjs/js"
	"honnef.co/go/js/dom"
)

func init() {
	style.Add(`label {
	display : block;
	float : left;
	text-align : right;
	width : 200px;
}

label:after {
	content : ':';
}

.sizeableInput {
	border : 2px inset #DCDAD5;
	padding-left : 3px;
	padding-right : 3px;
	min-width : 50px;
	height : 20px;
	margin-top : 2px;
}
`)
}

// InputSizeable returns a content-editable span that is style to look a text
// input box
func InputSizeable(id, value string) *dom.HTMLSpanElement {
	s := xdom.Span()
	s.Class().SetString("sizeableInput")
	s.SetContentEditable("true")
	s.Set("spellcheck", "false")
	if id != "" {
		s.SetID(id)
	}
	s.Underlying().Call("appendChild", js.Global.Get("document").Call("createTextNode", value))
	return s
}

// SizeableList is a collection of InputSizable elements
type SizeableList struct {
	*dom.HTMLDivElement
	contents []*dom.HTMLSpanElement
}

// InputSizeableList creates a list of InputSizeable elements, wrapped in a div
func InputSizeableList(values ...string) *SizeableList {
	d := xdom.Div()
	d.Class().SetString("sizeableList")
	contents := make([]*dom.HTMLSpanElement, len(values))
	for i, value := range values {
		s := InputSizeable("", value)
		d.AppendChild(s)
		contents[i] = s
	}
	sl := &SizeableList{
		d,
		contents,
	}
	remove := InputButton("", "-")
	remove.AddEventListener("click", false, func(dom.Event) {
		l := len(sl.contents) - 1
		d.RemoveChild(sl.contents[l])
		sl.contents = sl.contents[:l]
	})
	add := InputButton("", "+")
	add.AddEventListener("click", false, func(dom.Event) {
		s := InputSizeable("", "")
		d.InsertBefore(s, remove)
		sl.contents = append(sl.contents, s)
	})
	d.AppendChild(remove)
	d.AppendChild(add)
	return sl
}

// Values returns the values of the enclose InputSizeable's
func (s *SizeableList) Values() []string {
	v := make([]string, len(s.contents))
	for i, s := range s.contents {
		v[i] = s.TextContent()
	}
	return v
}

// Label create a form label
func Label(label, forID string) *dom.HTMLLabelElement {
	l := xdom.Label()
	l.For = forID
	l.Underlying().Call("appendChild", js.Global.Get("document").Call("createTextNode", label))
	return l
}

// InputText creates a text input box
func InputText(id, value string) *dom.HTMLInputElement {
	i := xdom.Input()
	i.Type = "text"
	if id != "" {
		i.SetID(id)
	}
	i.Value = value
	return i
}

// InputCheckbox creates a checkbox input
func InputCheckbox(id string, value bool) *dom.HTMLInputElement {
	i := xdom.Input()
	i.Type = "checkbox"
	if id != "" {
		i.SetID(id)
	}
	i.Checked = value
	return i
}

// InputRadio create a radio button input
func InputRadio(id, name string, value bool) *dom.HTMLInputElement {
	i := xdom.Input()
	i.Type = "radio"
	i.Name = name
	if id != "" {
		i.SetID(id)
	}
	i.Checked = value
	return i
}

// InputUpload creates an upload input field
func InputUpload(id string) *dom.HTMLInputElement {
	i := xdom.Input()
	i.Type = "file"
	if id != "" {
		i.SetID(id)
	}
	return i
}

// InputButton creates a button input
func InputButton(id, name string) *dom.HTMLInputElement {
	i := xdom.Input()
	i.Type = "button"
	i.Value = name
	if id != "" {
		i.SetID(id)
	}
	return i
}

// InputSubmit creates a submit input
func InputSubmit(name string) *dom.HTMLInputElement {
	i := xdom.Input()
	i.Type = "submit"
	i.Value = name
	return i
}

// InputPassword creates a password input
func InputPassword(id, value string) *dom.HTMLInputElement {
	i := xdom.Input()
	i.Type = "password"
	i.Value = value
	if id != "" {
		i.SetID(id)
	}
	return i
}

// InputNumber creates a text input that only allows number to be entered
func InputNumber(id string, min, max, value float64) *dom.HTMLInputElement {
	i := xdom.Input()
	i.Type = "number"
	i.Min = strconv.FormatFloat(min, 'f', -1, 64)
	i.Max = strconv.FormatFloat(max, 'f', -1, 64)
	i.ValueAsNumber = value
	if id != "" {
		i.SetID(id)
	}
	return i
}

// InputDate create a date based input, the workings of which are implementation
// specific
func InputDate(id string) *dom.HTMLInputElement {
	i := xdom.Input()
	i.Type = "date"
	if id != "" {
		i.SetID(id)
	}
	return i
}

// InputDateTime create a datetime based input, the workings of which are
// implementation specific
func InputDateTime(id string) *dom.HTMLInputElement {
	i := xdom.Input()
	i.Type = "datetime"
	if id != "" {
		i.SetID(id)
	}
	return i
}

// InputDateTimeLocal create a local datetime based input, the workings of
// which are implementation specific
func InputDateTimeLocal(id string) *dom.HTMLInputElement {
	i := xdom.Input()
	i.Type = "datetime-local"
	if id != "" {
		i.SetID(id)
	}
	return i
}

// InputMonth creates a month based input box
func InputMonth(id string) *dom.HTMLInputElement {
	i := xdom.Input()
	i.Type = "month"
	if id != "" {
		i.SetID(id)
	}
	return i
}

// InputWeek creates a week based input box
func InputWeek(id string) *dom.HTMLInputElement {
	i := xdom.Input()
	i.Type = "week"
	if id != "" {
		i.SetID(id)
	}
	return i
}

// InputTime creates a time based input box
func InputTime(id string) *dom.HTMLInputElement {
	i := xdom.Input()
	i.Type = "time"
	if id != "" {
		i.SetID(id)
	}
	return i
}

// InputColor creates a colour based input box, the workings of which are
// implementation specific
func InputColor(id string) *dom.HTMLInputElement {
	i := xdom.Input()
	i.Type = "color"
	if id != "" {
		i.SetID(id)
	}
	return i
}

// InputRange creates a sliding rule with which a number in the given range
// can be selected
func InputRange(id string, min, max, step, value float64) *dom.HTMLInputElement {
	i := xdom.Input()
	i.Type = "range"
	i.Min = strconv.FormatFloat(min, 'f', -1, 64)
	i.Max = strconv.FormatFloat(max, 'f', -1, 64)
	i.Value = strconv.FormatFloat(value, 'f', -1, 64)
	if step != step {
		i.Step = "all"
	} else {
		i.Step = strconv.FormatFloat(min, 'f', -1, 64)
	}
	if id != "" {
		i.SetID(id)
	}
	return i
}

// InputEmail is a text box that validates as an email address
func InputEmail(id, value string) *dom.HTMLInputElement {
	i := xdom.Input()
	i.Type = "email"
	i.Value = value
	if id != "" {
		i.SetID(id)
	}
	return i
}

// InputURL is a text box that validates as a URL
func InputURL(id, value string) *dom.HTMLInputElement {
	i := xdom.Input()
	i.Type = "url"
	i.Value = value
	if id != "" {
		i.SetID(id)
	}
	return i
}

// Option is a structure to define a 'select's option.
type Option struct {
	Label, Value string
	Selected     bool
}

// SelectBox provides a select input, filled with the given options
func SelectBox(id string, values ...Option) *dom.HTMLSelectElement {
	s := xdom.Select()
	if id != "" {
		s.SetID(id)
	}
	selected := false
	su := s.Underlying()
	for _, v := range values {
		o := xdom.Option()
		o.Value = v.Value
		if v.Selected && !selected {
			selected = true
			o.Selected = true
		}
		o.Underlying().Call("appendChild", js.Global.Get("document").Call("createTextNode", v.Label))
		su.Call("appendChild", o.Underlying())
	}
	return s
}

// TextArea provides a textarea input
func TextArea(id string, value string) *dom.HTMLTextAreaElement {
	t := xdom.Textarea()
	if id != "" {
		t.SetID(id)
	}
	t.Underlying().Call("appendChild", js.Global.Get("document").Call("createTextNode", value))
	return t
}
