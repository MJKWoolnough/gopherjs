// Package overlay creates a simple 'window' overlay system for gopherjs
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
	background-color : rgba(0, 0, 0, 0.25);
}

.mw-overlay > *:not(.closer) {
	position : absolute;
	top : 10px;
	left : 10px;
	right : 10px;
	background-color : #fff;
}

.mw-overlay > .closer {
	color : #fff;
	font-family : serif;
	right : 12px;
	position : absolute;
	text-shadow : #000 -1px -1px 0, #000 1px -1px 0, #000 1px 1px 0, #000 -1px 1px 0, #000 -1px 0 0, #000 1px 0 0;
	top : 10px;
	transition : color 0.25s;
	transition-timing-function : ease;
	user-select : none;
	z-index : 1;
	-moz-user-select : none;
	-webkit-user-select : none;
}

.mw-overlay > .closer:hover {
	color : #f00;
	cursor : pointer;	
}

* > .mw-overlay:not(:last-of-type) {
	background-color : transparent;
}
`

func init() {
	style.Add(css)
}

// Overlay is a dom.Element which
type Overlay struct {
	dom.Element
	onClose []func()
}

// OnClose adds an onClose handler for the overlay
func (o *Overlay) OnClose(f func()) {
	o.onClose = append(o.onClose, f)
}

// Close closes the overlay, and runs any given onClose handle funcs
func (o *Overlay) Close() error {
	p := o.ParentNode()
	if p != nil {
		p.RemoveChild(o)
		for _, f := range o.onClose {
			f()
		}
	}
	return nil
}

// New wraps a dom.Element in an Overlay
func New(e dom.Element) *Overlay {
	o := &Overlay{Element: xjs.CreateElement("div")}
	o.SetAttribute("class", "mw-overlay")
	c := xjs.CreateElement("div")
	xjs.SetInnerText(c, "X")
	c.SetAttribute("class", "closer")
	c.AddEventListener("click", false, func(dom.Event) {
		o.Close()
	})
	o.AppendChild(c)
	o.AppendChild(e)
	return o
}
