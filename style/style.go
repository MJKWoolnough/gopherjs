package style

import (
	"github.com/MJKWoolnough/gopherjs/xjs"
	"honnef.co/go/js/dom"
)

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
