package xform

import (
	"strconv"

	"github.com/MJKWoolnough/gopherjs/style"
	"github.com/MJKWoolnough/gopherjs/xdom"
	"github.com/MJKWoolnough/gopherjs/xjs"
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

func InputSizeable(id, value string) *dom.HTMLSpanElement {
	s := xdom.Span()
	s.Class().SetString("sizeableInput")
	s.SetContentEditable("true")
	s.Set("spellcheck", "false")
	if id != "" {
		s.SetID(id)
	}
	xjs.SetInnerText(s, value)
	return s
}

type SizeableList struct {
	*dom.HTMLDivElement
	contents []*dom.HTMLSpanElement
}

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

func (s *SizeableList) Values() []string {
	v := make([]string, len(s.contents))
	for i, s := range s.contents {
		v[i] = s.TextContent()
	}
	return v
}

func Label(label, forID string) *dom.HTMLLabelElement {
	l := xdom.Label()
	l.For = forID
	xjs.SetInnerText(l, label)
	return l
}

func InputText(id, value string) *dom.HTMLInputElement {
	i := xdom.Input()
	i.Type = "text"
	if id != "" {
		i.SetID(id)
	}
	i.Value = value
	return i
}

func InputCheckbox(id string, value bool) *dom.HTMLInputElement {
	i := xdom.Input()
	i.Type = "checkbox"
	if id != "" {
		i.SetID(id)
	}
	i.Checked = value
	return i
}

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

func InputUpload(id string) *dom.HTMLInputElement {
	i := xdom.Input()
	i.Type = "file"
	if id != "" {
		i.SetID(id)
	}
	return i
}

func InputButton(id, name string) *dom.HTMLInputElement {
	i := xdom.Input()
	i.Type = "button"
	i.Value = name
	if id != "" {
		i.SetID(id)
	}
	return i
}

func InputSubmit(name string) *dom.HTMLInputElement {
	i := xdom.Input()
	i.Type = "submit"
	i.Value = name
	return i
}

func InputPassword(id, value string) *dom.HTMLInputElement {
	i := xdom.Input()
	i.Type = "password"
	i.Value = value
	if id != "" {
		i.SetID(id)
	}
	return i
}

func InputNumber(id string, min, max float64) *dom.HTMLInputElement {
	i := xdom.Input()
	i.Type = "number"
	i.Min = strconv.FormatFloat(min, 'f', -1, 64)
	i.Max = strconv.FormatFloat(min, 'f', -1, 64)
	if id != "" {
		i.SetID(id)
	}
	return i
}

func InputDate(id string) *dom.HTMLInputElement {
	i := xdom.Input()
	i.Type = "date"
	if id != "" {
		i.SetID(id)
	}
	return i
}

func InputDateTime(id string) *dom.HTMLInputElement {
	i := xdom.Input()
	i.Type = "datetime"
	if id != "" {
		i.SetID(id)
	}
	return i
}

func InputDateTimeLocal(id string) *dom.HTMLInputElement {
	i := xdom.Input()
	i.Type = "datetime-local"
	if id != "" {
		i.SetID(id)
	}
	return i
}

func InputMonth(id string) *dom.HTMLInputElement {
	i := xdom.Input()
	i.Type = "month"
	if id != "" {
		i.SetID(id)
	}
	return i
}

func InputWeek(id string) *dom.HTMLInputElement {
	i := xdom.Input()
	i.Type = "week"
	if id != "" {
		i.SetID(id)
	}
	return i
}

func InputTime(id string) *dom.HTMLInputElement {
	i := xdom.Input()
	i.Type = "time"
	if id != "" {
		i.SetID(id)
	}
	return i
}

func InputColor(id string) *dom.HTMLInputElement {
	i := xdom.Input()
	i.Type = "color"
	if id != "" {
		i.SetID(id)
	}
	return i
}

func InputRange(id string, min, max, step float64) *dom.HTMLInputElement {
	i := xdom.Input()
	i.Type = "range"
	i.Min = strconv.FormatFloat(min, 'f', -1, 64)
	i.Max = strconv.FormatFloat(min, 'f', -1, 64)
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

func InputEmail(id, value string) *dom.HTMLInputElement {
	i := xdom.Input()
	i.Type = "email"
	i.Value = value
	if id != "" {
		i.SetID(id)
	}
	return i
}

func InputURL(id, value string) *dom.HTMLInputElement {
	i := xdom.Input()
	i.Type = "url"
	i.Value = value
	if id != "" {
		i.SetID(id)
	}
	return i
}

type Option struct {
	Label, Value string
	Selected     bool
}

func SelectBox(id string, values ...Option) *dom.HTMLSelectElement {
	s := xdom.Select()
	if id != "" {
		s.SetID(id)
	}
	selected := false
	for _, v := range values {
		o := xdom.Option()
		o.Value = v.Value
		if v.Selected && !selected {
			selected = true
			o.Selected = true
		}
		s.AppendChild(xjs.SetInnerText(o, v.Label))
	}
	return s
}

func TextArea(id string, value string) *dom.HTMLTextAreaElement {
	t := xdom.Textarea()
	if id != "" {
		t.SetID(id)
	}
	xjs.SetInnerText(t, value)
	return t
}
