// Package tabs implements a simple tab interface for HTML documents
package tabs

import (
	"github.com/MJKWoolnough/gopherjs/style"
	"github.com/MJKWoolnough/gopherjs/xdom"
	"github.com/MJKWoolnough/gopherjs/xjs"
	"honnef.co/go/js/dom"
)

const css = `.tabs {
	line-height : 24px;
	position : relative;
	width : 100%;
	overflow : hidden;
	margin : 0;
	padding : 0 0 0 20px;
	z-index : 0;
}

.tabs:after {
	position : absolute;
	content : "";
	width : 100%;
	bottom : 0;
	left : 0;
	border-bottom: 1px solid #000;
	z-index : 1;
	overflow : hidden;
	text-align : center;
	transform : translateX(-20px);
}

.tabs > div {
	border : 1px solid #000;
	display : inline-block;
	position : relative;
	z-index : 1;
	margin : 0 -5px;
	padding : 0 20px;
	border-top-right-radius: 6px;
	border-top-left-radius: 6px;
	background : linear-gradient(to bottom, #ececec 50%, #d1d1d1 100%);
	box-shadow : 0 3px 3px rgba(0, 0, 0, 0.4), inset 0 1px 0 #fff;
	text-shadow : 0 1px #fff;
	user-select : none;
	-moz-user-select : none;
	-webkit-user-select : none;
}

.tabs > div:hover {
	background : linear-gradient(to bottom, #faa 1%, #ffecec 50%, #d1d1d1 100%);
	cursor : pointer;
}

.tabs > div:before, .tabs > div:after {
	position : absolute;
	bottom : -1px;
	width : 6px;
	height : 6px;
	content : " ";
	border : 1px solid #000;
}

.tabs > div:before {
	left : -7px;
	border-bottom-right-radius : 6px;
	border-width : 0 1px 1px 0;
	box-shadow: 2px 2px 0 #d1d1d1;
}

.tabs > div:after {
	right : -7px;
	border-bottom-left-radius : 6px;
	border-width : 0 0 1px 1px;
	box-shadow: -2px 2px 0 #d1d1d1;
}

.tabs > div.selected {
	border-bottom-color : #fff;
	z-index : 2;
	background : #fff;
}

.tabs > div.selected:hover {
	background : #fff;
	cursor : default;
}

.tabs > div.selected:before {
	box-shadow: 2px 2px 0 #fff;
}

.tabs > div.selected:after {
	box-shadow: -2px 2px 0 #fff;
}

.tabs + .content {
	margin-top : 10px;
	padding : 10px 0 0 0;
}`

func init() {
	style.Add(css)
}

// Tab represents a tab in a menu
type Tab struct {
	// The name of the Tab
	Name string
	// Func to run when the tab is selected
	Func func(dom.Element)
}

// MakeTabs takes a list of tabs and generates a tabbed interface, which is return as a document fragment
func MakeTabs(t []Tab) dom.Node {
	f := xjs.DocumentFragment()
	if len(t) < 0 {
		return f
	}
	tabsDiv := xdom.Div()
	contents := xdom.Div()
	tabsDiv.Class().SetString("tabs")
	contents.Class().SetString("content")
	tabs := make([]dom.Element, len(t))
	for n := range t {
		tabs[n] = xjs.SetInnerText(xdom.Div(), t[n].Name).(dom.Element)
		tabs[n].AddEventListener("click", false, func() func(dom.Event) {
			i := n
			return func(e dom.Event) {
				if tabs[i].Class().String() == "selected" {
					return
				}
				for _, tab := range tabs {
					tab.Class().Remove("selected")
				}
				newContents := contents.CloneNode(false).(*dom.HTMLDivElement)
				contents.ParentNode().ReplaceChild(newContents, contents)
				contents = newContents
				tabs[i].Class().Add("selected")
				go t[i].Func(contents)
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
