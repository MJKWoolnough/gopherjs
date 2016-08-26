// Package xjs provides some simple, but often needed shortcut funcs for gopherJS
package xjs

import (
	"fmt"
	"strings"

	"github.com/gopherjs/gopherjs/js"

	"honnef.co/go/js/dom"
)

// body var used to cache the wrapped body once Body is sucessfully called
var body *dom.HTMLBodyElement

// Body gets the document's body element
func Body() *dom.HTMLBodyElement {
	if body == nil {
		o := js.Global.Get("document").Get("body")
		if o == nil || o == js.Undefined {
			return nil
		}
		body = &dom.HTMLBodyElement{BasicHTMLElement: &dom.BasicHTMLElement{&dom.BasicElement{&dom.BasicNode{js.Global.Get("document").Get("body")}}}}
	}
	return body
}

// RemoveChildren removes all of the child nodes of the node given
func RemoveChildren(node dom.Node) dom.Node {
	n := node.Underlying()
	for n.Call("hasChildNodes").Bool() {
		n.Call("removeChild", n.Get("lastChild"))
	}
	return node
}

// SetInnerText removes all child nodes from the given node and sets a single
// Text Node with the given string
func SetInnerText(node dom.Node, text string) dom.Node {
	n := node.Underlying()
	for n.Call("hasChildNodes").Bool() {
		n.Call("removeChild", n.Get("lastChild"))
	}
	n.Call("appendChild", js.Global.Get("document").Call("createTextNode", text))
	return node
}

// SetPreText does similar to SetInnerText, but linebreaks are converted to <br />s
func SetPreText(node dom.Node, text string) dom.Node {
	n := node.Underlying()
	for n.Call("hasChildNodes").Bool() {
		n.Call("removeChild", n.Get("lastChild"))
	}
	for i, part := range strings.Split(text, "\n") {
		if i > 0 {
			n.Call("appendChild", js.Global.Get("document").Call("createElement", "br"))
		}
		n.Call("appendChild", js.Global.Get("document").Call("createTextNode", part))
	}
	return node
}

// CreateElement is a shortcut to create an element with the given name
func CreateElement(name string) dom.Element {
	return dom.WrapElement(js.Global.Get("document").Call("createElement", name))
}

// Alert provides for formated alert boxes
func Alert(format string, params ...interface{}) {
	js.Global.Get("document").Call("alert", fmt.Sprintf(format, params...))
}

// Text creates a text node containing the givin text
func Text(text string) *dom.Text {
	return &dom.Text{&dom.BasicNode{js.Global.Get("document").Call("createTextNode", text)}}
}

// AppendChildren appends all the given children to the parent.
func AppendChildren(parent dom.Node, children ...dom.Node) dom.Node {
	for _, child := range children {
		parent.AppendChild(child)
	}
	return parent
}

// Log prints a formatted string to the javascript console
func Log(format string, params ...interface{}) {
	js.Global.Get("console").Call("log", fmt.Sprintf(format, params...))
}
