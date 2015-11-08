// Package xjs provides some simple, but often needed shortcut funcs for gopherJS
package xjs

import (
	"fmt"
	"strings"

	"honnef.co/go/js/dom"
)

var docNode = dom.GetWindow().Document()

// Get the document's body element
func Body() dom.HTMLElement {
	return docNode.(dom.HTMLDocument).Body()
}

// DocumentFragment returns a new DocumentFragment as a dom.Node
func DocumentFragment() dom.Node {
	return dom.WrapNode(docNode.Underlying().Call("createDocumentFragment"))
}

// RemoveChildren removes all of the child nodes of the node given
func RemoveChildren(node dom.Node) dom.Node {
	for node.HasChildNodes() {
		node.RemoveChild(node.LastChild())
	}
	return node
}

// SetInnerText removes all child nodes from the given node and sets a single
// Text Node with the given string
func SetInnerText(node dom.Node, text string) dom.Node {
	RemoveChildren(node)
	node.AppendChild(Text(text))
	return node
}

// SetPreText does similar to SetInnerText, but linebreaks are converted to <br />s
func SetPreText(node dom.Node, text string) dom.Node {
	RemoveChildren(node)
	for n, part := range strings.Split(text, "\n") {
		if n > 0 {
			node.AppendChild(CreateElement("br"))
		}
		node.AppendChild(Text(part))
	}
	return node
}

// CreateElement is a shortcut to create an element with the given name
func CreateElement(name string) dom.Element {
	return dom.GetWindow().Document().CreateElement(name)
}

// Alert provides for formated alert boxes
func Alert(format string, params ...interface{}) {
	dom.GetWindow().Alert(fmt.Sprintf(format, params...))
}

// TextNode creates a text node containing the givin text
func Text(text string) *dom.Text {
	return docNode.CreateTextNode(text)
}

// AppendChildren appends all the given children to the parent.
func AppendChildren(parent dom.Node, children ...dom.Node) dom.Node {
	for _, child := range children {
		parent.AppendChild(child)
	}
	return parent
}
