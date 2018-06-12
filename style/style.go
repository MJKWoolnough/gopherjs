// Package style implements a simple way to handle styling in an HTML document
package style // import "vimagination.zapto.org/gopherjs/style"

import (
	"honnef.co/go/js/dom"
	"vimagination.zapto.org/gopherjs/xdom"
	"vimagination.zapto.org/gopherjs/xjs"
)

// Add adds the given CSS string to the DOM
func Add(css string) dom.Element {
	h := dom.GetWindow().Document().GetElementsByTagName("head")
	if len(h) == 0 {
		return nil
	}
	s := xdom.Style()
	h[0].AppendChild(xjs.SetInnerText(s, css))
	return s
}
