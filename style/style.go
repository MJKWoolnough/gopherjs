// Package style implements a simple way to handle styling in an HTML document
package style

import (
	"github.com/MJKWoolnough/gopherjs/xjs"
	"honnef.co/go/js/dom"
)

// Add adds the given CSS string to the DOM
func Add(css string) dom.Element {
	h := dom.GetWindow().Document().GetElementsByTagName("head")
	if len(h) == 0 {
		return nil
	}
	s := xjs.CreateElement("style")
	xjs.SetInnerText(s, css)
	h[0].AppendChild(s)
	return s
}
