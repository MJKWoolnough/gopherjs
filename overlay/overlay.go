package overlay

import (
	"github.com/MJKWoolnough/gopherjs/style"
	"github.com/MJKWoolnough/gopherjs/xjs"
	"honnef.co/go/js/dom"
)

const css = `
window, body {
	width : 100%;
	margin : 0;
	overflow-x : hidden;
}

.mw-overlay {
	position : fixed;
	top : 0;
	left : 0;
	right : 0;
	bottom : 0;
	background : rgba(0, 0, 0, 0.25);
}

.mw-overlay > * {
	position : absolute;
	top : 10px;
	left : 10px;
	right : 10px;
	background : #fff;
}

.mw-overlay .closer {
	color : #fff;
	font-family : serif;
	right : 2px;
	position : absolute;
	text-shadow : #000 -1px -1px 0, #000 1px -1px 0, #000 1px 1px 0, #000 -1px 1px 0, #000 -1px 0 0, #000 1px 0 0;
	top : 0;
	transition : color 0.25s;
	transition-timing-function : ease;
	user-select : none;
	z-index : 1;
	-moz-user-select : none;
	-webkit-user-select : none;
}

.mw-overlay .closer:hover {
	color : #f00;
	cursor : pointer;	
}`

func init() {
	style.Add(css)
}

func Overlay(e dom.Element) dom.Element {
	o := xjs.CreateElement("div")
	o.SetAttribute("class", "mw-overlay")
	c := xjs.CreateElement("div")
	xjs.SetInnerText(c, "X")
	c.SetAttribute("class", "closer")
	c.AddEventListener("click", false, func(dom.Event) {
		p := o.ParentNode()
		if p != nil {
			p.RemoveChild(o)
		}
	})
	e.InsertBefore(c, e.FirstChild())
	o.AppendChild(e)
	return o
}
