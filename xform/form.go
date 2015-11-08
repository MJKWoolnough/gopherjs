package xform

import (
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

.sizeableInput {` + dom.GetWindow().Document().(dom.HTMLDocument).DefaultView().GetComputedStyle(xdom.Input(), "").String() + `}
`)
}

func InputSizeable(id, value string) *dom.HTMLSpanElement {
	s := xdom.Span()
	s.Class().SetString("sizeableInput")
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
	remove := xdom.Button()
	remove.Value = "-"
	remove.AddEventListener("click", false, func(dom.Event) {
		l := len(sl.contents) - 1
		d.RemoveChild(sl.contents[l])
		sl.contents = sl.contents[:l]
	})
	add := xdom.Button()
	add.Value = "+"
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
	return i
}

func InputSubmit(name string) *dom.HTMLInputElement {
	i := xdom.Input()
	i.Type = "submit"
	i.Value = name
	return i
}

type Option struct {
	Label, Value string
}

func SelectBox(id string, values ...Option) *dom.HTMLSelectElement {
	s := xdom.Select()
	if id != "" {
		s.SetID(id)
	}
	for _, v := range values {
		o := xdom.Option()
		o.Value = v.Value
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
