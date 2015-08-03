package xjs

import (
	"strings"

	"github.com/gopherjs/gopherjs/js"
	"honnef.co/go/js/dom"
)

func DocumentFragment() dom.Node {
	return dom.WrapNode(js.Global.Get("document").Call("createDocumentFragment"))
}

func RemoveChildren(node dom.Node) dom.Node {
	for node.HasChildNodes() {
		node.RemoveChild(node.LastChild())
	}
	return node
}

func SetInnerText(node dom.Node, text string) dom.Node {
	RemoveChildren(node)
	node.AppendChild(dom.GetWindow().Document().CreateTextNode(text))
	return node
}

func SetPreText(node dom.Node, text string) dom.Node {
	RemoveChildren(node)
	for n, part := range strings.Split(text, "\n") {
		if n > 0 {
			node.AppendChild(CreateElement("br"))
		}
		node.AppendChild(dom.GetWindow().Document().CreateTextNode(part))
	}
	return node
}

func CreateElement(name string) dom.Element {
	return dom.GetWindow().Document().CreateElement(name)
}
