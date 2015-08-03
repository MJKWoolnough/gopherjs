package tabs

import (
	"github.com/MJKWoolnough/gopherjs/xjs"
	"honnef.co/go/js/dom"
)

type Tab struct {
	Name string
	Func func(dom.Element)
}

func MakeTabs(t []Tab) dom.Node {
	f := xjs.DocumentFragment()
	if len(t) < 0 {
		return f
	}
	tabsDiv := xjs.CreateElement("div")
	contents := xjs.CreateElement("div")
	tabsDiv.Class().SetString("tabs")
	contents.Class().SetString("content")
	tabs := make([]dom.Element, len(t))
	for n := range t {
		tabs[n] = xjs.SetInnerText(xjs.CreateElement("div"), t[n].Name).(dom.Element)
		tabs[n].AddEventListener("click", false, func() func(dom.Event) {
			i := n
			return func(e dom.Event) {
				if tabs[i].Class().String() == "selected" {
					return
				}
				for _, tab := range tabs {
					tab.Class().Remove("selected")
				}
				xjs.RemoveChildren(contents)
				tabs[i].Class().Add("selected")
				t[i].Func(contents)
			}
		}())
		tabsDiv.AppendChild(tabs[n])
	}
	t[0].Func(contents)
	tabs[0].Class().Add("selected")
	f.AppendChild(tabsDiv)
	f.AppendChild(contents)
	return f
}
