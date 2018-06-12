// Package xdom provides simple functions to create HTML elements with the correct type
package xdom // import "vimagination.zapto.org/gopherjs/xdom"

// File automatically generated with ./genMap.sh

import (
	"github.com/gopherjs/gopherjs/js"
	"honnef.co/go/js/dom"
)

// Text returns a text node with the text set to the given string
func Text(s string) *dom.Text {
	return &dom.Text{&dom.BasicNode{js.Global.Get("document").Call("createTextNode", s)}}
}

// DocumentFragment returns a new DocumentFragment
func DocumentFragment() *dom.BasicHTMLElement {
	return &dom.BasicHTMLElement{&dom.BasicElement{&dom.BasicNode{js.Global.Get("document").Call("createDocumentFragment")}}}
}

// A returns a "a" element with type *dom.HTMLAnchorElement
func A() *dom.HTMLAnchorElement {
	o := js.Global.Get("document").Call("createElement", "a")
	return &dom.HTMLAnchorElement{BasicHTMLElement: &dom.BasicHTMLElement{&dom.BasicElement{&dom.BasicNode{o}}}, URLUtils: &dom.URLUtils{Object: o}}
}

// Abbr returns a "abbr" element with type *dom.BasicHTMLElement
func Abbr() *dom.BasicHTMLElement {
	return &dom.BasicHTMLElement{&dom.BasicElement{&dom.BasicNode{js.Global.Get("document").Call("createElement", "abbr")}}}
}

// Acronym returns a "acronym" element with type *dom.BasicHTMLElement
func Acronym() *dom.BasicHTMLElement {
	return &dom.BasicHTMLElement{&dom.BasicElement{&dom.BasicNode{js.Global.Get("document").Call("createElement", "acronym")}}}
}

// Address returns a "address" element with type *dom.BasicHTMLElement
func Address() *dom.BasicHTMLElement {
	return &dom.BasicHTMLElement{&dom.BasicElement{&dom.BasicNode{js.Global.Get("document").Call("createElement", "address")}}}
}

// Applet returns a "applet" element with type *dom.HTMLAppletElement
func Applet() *dom.HTMLAppletElement {
	return &dom.HTMLAppletElement{BasicHTMLElement: &dom.BasicHTMLElement{&dom.BasicElement{&dom.BasicNode{js.Global.Get("document").Call("createElement", "applet")}}}}
}

// Area returns a "area" element with type *dom.HTMLAreaElement
func Area() *dom.HTMLAreaElement {
	o := js.Global.Get("document").Call("createElement", "area")
	return &dom.HTMLAreaElement{BasicHTMLElement: &dom.BasicHTMLElement{&dom.BasicElement{&dom.BasicNode{o}}}, URLUtils: &dom.URLUtils{Object: o}}
}

// Article returns a "article" element with type *dom.BasicHTMLElement
func Article() *dom.BasicHTMLElement {
	return &dom.BasicHTMLElement{&dom.BasicElement{&dom.BasicNode{js.Global.Get("document").Call("createElement", "article")}}}
}

// Aside returns a "aside" element with type *dom.BasicHTMLElement
func Aside() *dom.BasicHTMLElement {
	return &dom.BasicHTMLElement{&dom.BasicElement{&dom.BasicNode{js.Global.Get("document").Call("createElement", "aside")}}}
}

// Audio returns a "audio" element with type *dom.HTMLAudioElement
func Audio() *dom.HTMLAudioElement {
	return &dom.HTMLAudioElement{HTMLMediaElement: &dom.HTMLMediaElement{BasicHTMLElement: &dom.BasicHTMLElement{&dom.BasicElement{&dom.BasicNode{js.Global.Get("document").Call("createElement", "audio")}}}}}
}

// B returns a "b" element with type *dom.BasicHTMLElement
func B() *dom.BasicHTMLElement {
	return &dom.BasicHTMLElement{&dom.BasicElement{&dom.BasicNode{js.Global.Get("document").Call("createElement", "b")}}}
}

// Base returns a "base" element with type *dom.HTMLBaseElement
func Base() *dom.HTMLBaseElement {
	return &dom.HTMLBaseElement{BasicHTMLElement: &dom.BasicHTMLElement{&dom.BasicElement{&dom.BasicNode{js.Global.Get("document").Call("createElement", "base")}}}}
}

// Basefont returns a "basefont" element with type *dom.BasicHTMLElement
func Basefont() *dom.BasicHTMLElement {
	return &dom.BasicHTMLElement{&dom.BasicElement{&dom.BasicNode{js.Global.Get("document").Call("createElement", "basefont")}}}
}

// Bdi returns a "bdi" element with type *dom.BasicHTMLElement
func Bdi() *dom.BasicHTMLElement {
	return &dom.BasicHTMLElement{&dom.BasicElement{&dom.BasicNode{js.Global.Get("document").Call("createElement", "bdi")}}}
}

// Bdo returns a "bdo" element with type *dom.BasicHTMLElement
func Bdo() *dom.BasicHTMLElement {
	return &dom.BasicHTMLElement{&dom.BasicElement{&dom.BasicNode{js.Global.Get("document").Call("createElement", "bdo")}}}
}

// Big returns a "big" element with type *dom.BasicHTMLElement
func Big() *dom.BasicHTMLElement {
	return &dom.BasicHTMLElement{&dom.BasicElement{&dom.BasicNode{js.Global.Get("document").Call("createElement", "big")}}}
}

// Blockquote returns a "blockquote" element with type *dom.HTMLQuoteElement
func Blockquote() *dom.HTMLQuoteElement {
	return &dom.HTMLQuoteElement{BasicHTMLElement: &dom.BasicHTMLElement{&dom.BasicElement{&dom.BasicNode{js.Global.Get("document").Call("createElement", "blockquote")}}}}
}

// Body returns a "body" element with type *dom.HTMLBodyElement
func Body() *dom.HTMLBodyElement {
	return &dom.HTMLBodyElement{BasicHTMLElement: &dom.BasicHTMLElement{&dom.BasicElement{&dom.BasicNode{js.Global.Get("document").Call("createElement", "body")}}}}
}

// Br returns a "br" element with type *dom.HTMLBRElement
func Br() *dom.HTMLBRElement {
	return &dom.HTMLBRElement{BasicHTMLElement: &dom.BasicHTMLElement{&dom.BasicElement{&dom.BasicNode{js.Global.Get("document").Call("createElement", "br")}}}}
}

// Button returns a "button" element with type *dom.HTMLButtonElement
func Button() *dom.HTMLButtonElement {
	return &dom.HTMLButtonElement{BasicHTMLElement: &dom.BasicHTMLElement{&dom.BasicElement{&dom.BasicNode{js.Global.Get("document").Call("createElement", "button")}}}}
}

// Canvas returns a "canvas" element with type *dom.HTMLCanvasElement
func Canvas() *dom.HTMLCanvasElement {
	return &dom.HTMLCanvasElement{BasicHTMLElement: &dom.BasicHTMLElement{&dom.BasicElement{&dom.BasicNode{js.Global.Get("document").Call("createElement", "canvas")}}}}
}

// Caption returns a "caption" element with type *dom.HTMLTableCaptionElement
func Caption() *dom.HTMLTableCaptionElement {
	return &dom.HTMLTableCaptionElement{BasicHTMLElement: &dom.BasicHTMLElement{&dom.BasicElement{&dom.BasicNode{js.Global.Get("document").Call("createElement", "caption")}}}}
}

// Center returns a "center" element with type *dom.BasicHTMLElement
func Center() *dom.BasicHTMLElement {
	return &dom.BasicHTMLElement{&dom.BasicElement{&dom.BasicNode{js.Global.Get("document").Call("createElement", "center")}}}
}

// Cite returns a "cite" element with type *dom.HTMLQuoteElement
func Cite() *dom.HTMLQuoteElement {
	return &dom.HTMLQuoteElement{BasicHTMLElement: &dom.BasicHTMLElement{&dom.BasicElement{&dom.BasicNode{js.Global.Get("document").Call("createElement", "cite")}}}}
}

// Code returns a "code" element with type *dom.BasicHTMLElement
func Code() *dom.BasicHTMLElement {
	return &dom.BasicHTMLElement{&dom.BasicElement{&dom.BasicNode{js.Global.Get("document").Call("createElement", "code")}}}
}

// Col returns a "col" element with type *dom.HTMLTableColElement
func Col() *dom.HTMLTableColElement {
	return &dom.HTMLTableColElement{BasicHTMLElement: &dom.BasicHTMLElement{&dom.BasicElement{&dom.BasicNode{js.Global.Get("document").Call("createElement", "col")}}}}
}

// Colgroup returns a "colgroup" element with type *dom.HTMLTableColElement
func Colgroup() *dom.HTMLTableColElement {
	return &dom.HTMLTableColElement{BasicHTMLElement: &dom.BasicHTMLElement{&dom.BasicElement{&dom.BasicNode{js.Global.Get("document").Call("createElement", "colgroup")}}}}
}

// Content returns a "content" element with type *dom.BasicHTMLElement
func Content() *dom.BasicHTMLElement {
	return &dom.BasicHTMLElement{&dom.BasicElement{&dom.BasicNode{js.Global.Get("document").Call("createElement", "content")}}}
}

// Data returns a "data" element with type *dom.HTMLDataElement
func Data() *dom.HTMLDataElement {
	return &dom.HTMLDataElement{BasicHTMLElement: &dom.BasicHTMLElement{&dom.BasicElement{&dom.BasicNode{js.Global.Get("document").Call("createElement", "data")}}}}
}

// Datalist returns a "datalist" element with type *dom.HTMLDataListElement
func Datalist() *dom.HTMLDataListElement {
	return &dom.HTMLDataListElement{BasicHTMLElement: &dom.BasicHTMLElement{&dom.BasicElement{&dom.BasicNode{js.Global.Get("document").Call("createElement", "datalist")}}}}
}

// Dd returns a "dd" element with type *dom.BasicHTMLElement
func Dd() *dom.BasicHTMLElement {
	return &dom.BasicHTMLElement{&dom.BasicElement{&dom.BasicNode{js.Global.Get("document").Call("createElement", "dd")}}}
}

// Del returns a "del" element with type *dom.HTMLModElement
func Del() *dom.HTMLModElement {
	return &dom.HTMLModElement{BasicHTMLElement: &dom.BasicHTMLElement{&dom.BasicElement{&dom.BasicNode{js.Global.Get("document").Call("createElement", "del")}}}}
}

// Details returns a "details" element with type *dom.BasicHTMLElement
func Details() *dom.BasicHTMLElement {
	return &dom.BasicHTMLElement{&dom.BasicElement{&dom.BasicNode{js.Global.Get("document").Call("createElement", "details")}}}
}

// Dfn returns a "dfn" element with type *dom.BasicHTMLElement
func Dfn() *dom.BasicHTMLElement {
	return &dom.BasicHTMLElement{&dom.BasicElement{&dom.BasicNode{js.Global.Get("document").Call("createElement", "dfn")}}}
}

// Dialog returns a "dialog" element with type *dom.BasicHTMLElement
func Dialog() *dom.BasicHTMLElement {
	return &dom.BasicHTMLElement{&dom.BasicElement{&dom.BasicNode{js.Global.Get("document").Call("createElement", "dialog")}}}
}

// Dir returns a "dir" element with type *dom.HTMLDirectoryElement
func Dir() *dom.HTMLDirectoryElement {
	return &dom.HTMLDirectoryElement{BasicHTMLElement: &dom.BasicHTMLElement{&dom.BasicElement{&dom.BasicNode{js.Global.Get("document").Call("createElement", "dir")}}}}
}

// Div returns a "div" element with type *dom.HTMLDivElement
func Div() *dom.HTMLDivElement {
	return &dom.HTMLDivElement{BasicHTMLElement: &dom.BasicHTMLElement{&dom.BasicElement{&dom.BasicNode{js.Global.Get("document").Call("createElement", "div")}}}}
}

// Dl returns a "dl" element with type *dom.HTMLDListElement
func Dl() *dom.HTMLDListElement {
	return &dom.HTMLDListElement{BasicHTMLElement: &dom.BasicHTMLElement{&dom.BasicElement{&dom.BasicNode{js.Global.Get("document").Call("createElement", "dl")}}}}
}

// Dt returns a "dt" element with type *dom.BasicHTMLElement
func Dt() *dom.BasicHTMLElement {
	return &dom.BasicHTMLElement{&dom.BasicElement{&dom.BasicNode{js.Global.Get("document").Call("createElement", "dt")}}}
}

// Em returns a "em" element with type *dom.BasicHTMLElement
func Em() *dom.BasicHTMLElement {
	return &dom.BasicHTMLElement{&dom.BasicElement{&dom.BasicNode{js.Global.Get("document").Call("createElement", "em")}}}
}

// Embed returns a "embed" element with type *dom.HTMLEmbedElement
func Embed() *dom.HTMLEmbedElement {
	return &dom.HTMLEmbedElement{BasicHTMLElement: &dom.BasicHTMLElement{&dom.BasicElement{&dom.BasicNode{js.Global.Get("document").Call("createElement", "embed")}}}}
}

// Fieldset returns a "fieldset" element with type *dom.HTMLFieldSetElement
func Fieldset() *dom.HTMLFieldSetElement {
	return &dom.HTMLFieldSetElement{BasicHTMLElement: &dom.BasicHTMLElement{&dom.BasicElement{&dom.BasicNode{js.Global.Get("document").Call("createElement", "fieldset")}}}}
}

// Figcaption returns a "figcaption" element with type *dom.BasicHTMLElement
func Figcaption() *dom.BasicHTMLElement {
	return &dom.BasicHTMLElement{&dom.BasicElement{&dom.BasicNode{js.Global.Get("document").Call("createElement", "figcaption")}}}
}

// Figure returns a "figure" element with type *dom.BasicHTMLElement
func Figure() *dom.BasicHTMLElement {
	return &dom.BasicHTMLElement{&dom.BasicElement{&dom.BasicNode{js.Global.Get("document").Call("createElement", "figure")}}}
}

// Font returns a "font" element with type *dom.HTMLFontElement
func Font() *dom.HTMLFontElement {
	return &dom.HTMLFontElement{BasicHTMLElement: &dom.BasicHTMLElement{&dom.BasicElement{&dom.BasicNode{js.Global.Get("document").Call("createElement", "font")}}}}
}

// Footer returns a "footer" element with type *dom.BasicHTMLElement
func Footer() *dom.BasicHTMLElement {
	return &dom.BasicHTMLElement{&dom.BasicElement{&dom.BasicNode{js.Global.Get("document").Call("createElement", "footer")}}}
}

// Form returns a "form" element with type *dom.HTMLFormElement
func Form() *dom.HTMLFormElement {
	return &dom.HTMLFormElement{BasicHTMLElement: &dom.BasicHTMLElement{&dom.BasicElement{&dom.BasicNode{js.Global.Get("document").Call("createElement", "form")}}}}
}

// Frame returns a "frame" element with type *dom.HTMLFrameElement
func Frame() *dom.HTMLFrameElement {
	return &dom.HTMLFrameElement{BasicHTMLElement: &dom.BasicHTMLElement{&dom.BasicElement{&dom.BasicNode{js.Global.Get("document").Call("createElement", "frame")}}}}
}

// Frameset returns a "frameset" element with type *dom.HTMLFrameSetElement
func Frameset() *dom.HTMLFrameSetElement {
	return &dom.HTMLFrameSetElement{BasicHTMLElement: &dom.BasicHTMLElement{&dom.BasicElement{&dom.BasicNode{js.Global.Get("document").Call("createElement", "frameset")}}}}
}

// H1 returns a "h1" element with type *dom.HTMLHeadingElement
func H1() *dom.HTMLHeadingElement {
	return &dom.HTMLHeadingElement{BasicHTMLElement: &dom.BasicHTMLElement{&dom.BasicElement{&dom.BasicNode{js.Global.Get("document").Call("createElement", "h1")}}}}
}

// H2 returns a "h2" element with type *dom.HTMLHeadingElement
func H2() *dom.HTMLHeadingElement {
	return &dom.HTMLHeadingElement{BasicHTMLElement: &dom.BasicHTMLElement{&dom.BasicElement{&dom.BasicNode{js.Global.Get("document").Call("createElement", "h2")}}}}
}

// H3 returns a "h3" element with type *dom.HTMLHeadingElement
func H3() *dom.HTMLHeadingElement {
	return &dom.HTMLHeadingElement{BasicHTMLElement: &dom.BasicHTMLElement{&dom.BasicElement{&dom.BasicNode{js.Global.Get("document").Call("createElement", "h3")}}}}
}

// H4 returns a "h4" element with type *dom.HTMLHeadingElement
func H4() *dom.HTMLHeadingElement {
	return &dom.HTMLHeadingElement{BasicHTMLElement: &dom.BasicHTMLElement{&dom.BasicElement{&dom.BasicNode{js.Global.Get("document").Call("createElement", "h4")}}}}
}

// H5 returns a "h5" element with type *dom.HTMLHeadingElement
func H5() *dom.HTMLHeadingElement {
	return &dom.HTMLHeadingElement{BasicHTMLElement: &dom.BasicHTMLElement{&dom.BasicElement{&dom.BasicNode{js.Global.Get("document").Call("createElement", "h5")}}}}
}

// H6 returns a "h6" element with type *dom.HTMLHeadingElement
func H6() *dom.HTMLHeadingElement {
	return &dom.HTMLHeadingElement{BasicHTMLElement: &dom.BasicHTMLElement{&dom.BasicElement{&dom.BasicNode{js.Global.Get("document").Call("createElement", "h6")}}}}
}

// Head returns a "head" element with type *dom.HTMLHeadElement
func Head() *dom.HTMLHeadElement {
	return &dom.HTMLHeadElement{BasicHTMLElement: &dom.BasicHTMLElement{&dom.BasicElement{&dom.BasicNode{js.Global.Get("document").Call("createElement", "head")}}}}
}

// Header returns a "header" element with type *dom.BasicHTMLElement
func Header() *dom.BasicHTMLElement {
	return &dom.BasicHTMLElement{&dom.BasicElement{&dom.BasicNode{js.Global.Get("document").Call("createElement", "header")}}}
}

// Hr returns a "hr" element with type *dom.HTMLHRElement
func Hr() *dom.HTMLHRElement {
	return &dom.HTMLHRElement{BasicHTMLElement: &dom.BasicHTMLElement{&dom.BasicElement{&dom.BasicNode{js.Global.Get("document").Call("createElement", "hr")}}}}
}

// Html returns a "html" element with type *dom.HTMLHtmlElement
func Html() *dom.HTMLHtmlElement {
	return &dom.HTMLHtmlElement{BasicHTMLElement: &dom.BasicHTMLElement{&dom.BasicElement{&dom.BasicNode{js.Global.Get("document").Call("createElement", "html")}}}}
}

// I returns a "i" element with type *dom.BasicHTMLElement
func I() *dom.BasicHTMLElement {
	return &dom.BasicHTMLElement{&dom.BasicElement{&dom.BasicNode{js.Global.Get("document").Call("createElement", "i")}}}
}

// Iframe returns a "iframe" element with type *dom.HTMLIFrameElement
func Iframe() *dom.HTMLIFrameElement {
	return &dom.HTMLIFrameElement{BasicHTMLElement: &dom.BasicHTMLElement{&dom.BasicElement{&dom.BasicNode{js.Global.Get("document").Call("createElement", "iframe")}}}}
}

// Img returns a "img" element with type *dom.HTMLImageElement
func Img() *dom.HTMLImageElement {
	return &dom.HTMLImageElement{BasicHTMLElement: &dom.BasicHTMLElement{&dom.BasicElement{&dom.BasicNode{js.Global.Get("document").Call("createElement", "img")}}}}
}

// Input returns a "input" element with type *dom.HTMLInputElement
func Input() *dom.HTMLInputElement {
	return &dom.HTMLInputElement{BasicHTMLElement: &dom.BasicHTMLElement{&dom.BasicElement{&dom.BasicNode{js.Global.Get("document").Call("createElement", "input")}}}}
}

// Ins returns a "ins" element with type *dom.HTMLModElement
func Ins() *dom.HTMLModElement {
	return &dom.HTMLModElement{BasicHTMLElement: &dom.BasicHTMLElement{&dom.BasicElement{&dom.BasicNode{js.Global.Get("document").Call("createElement", "ins")}}}}
}

// Kbd returns a "kbd" element with type *dom.BasicHTMLElement
func Kbd() *dom.BasicHTMLElement {
	return &dom.BasicHTMLElement{&dom.BasicElement{&dom.BasicNode{js.Global.Get("document").Call("createElement", "kbd")}}}
}

// Keygen returns a "keygen" element with type *dom.HTMLKeygenElement
func Keygen() *dom.HTMLKeygenElement {
	return &dom.HTMLKeygenElement{BasicHTMLElement: &dom.BasicHTMLElement{&dom.BasicElement{&dom.BasicNode{js.Global.Get("document").Call("createElement", "keygen")}}}}
}

// Label returns a "label" element with type *dom.HTMLLabelElement
func Label() *dom.HTMLLabelElement {
	return &dom.HTMLLabelElement{BasicHTMLElement: &dom.BasicHTMLElement{&dom.BasicElement{&dom.BasicNode{js.Global.Get("document").Call("createElement", "label")}}}}
}

// Legend returns a "legend" element with type *dom.HTMLLegendElement
func Legend() *dom.HTMLLegendElement {
	return &dom.HTMLLegendElement{BasicHTMLElement: &dom.BasicHTMLElement{&dom.BasicElement{&dom.BasicNode{js.Global.Get("document").Call("createElement", "legend")}}}}
}

// Li returns a "li" element with type *dom.HTMLLIElement
func Li() *dom.HTMLLIElement {
	return &dom.HTMLLIElement{BasicHTMLElement: &dom.BasicHTMLElement{&dom.BasicElement{&dom.BasicNode{js.Global.Get("document").Call("createElement", "li")}}}}
}

// Link returns a "link" element with type *dom.HTMLLinkElement
func Link() *dom.HTMLLinkElement {
	return &dom.HTMLLinkElement{BasicHTMLElement: &dom.BasicHTMLElement{&dom.BasicElement{&dom.BasicNode{js.Global.Get("document").Call("createElement", "link")}}}}
}

// Main returns a "main" element with type *dom.BasicHTMLElement
func Main() *dom.BasicHTMLElement {
	return &dom.BasicHTMLElement{&dom.BasicElement{&dom.BasicNode{js.Global.Get("document").Call("createElement", "main")}}}
}

// Map returns a "map" element with type *dom.HTMLMapElement
func Map() *dom.HTMLMapElement {
	return &dom.HTMLMapElement{BasicHTMLElement: &dom.BasicHTMLElement{&dom.BasicElement{&dom.BasicNode{js.Global.Get("document").Call("createElement", "map")}}}}
}

// Mark returns a "mark" element with type *dom.BasicHTMLElement
func Mark() *dom.BasicHTMLElement {
	return &dom.BasicHTMLElement{&dom.BasicElement{&dom.BasicNode{js.Global.Get("document").Call("createElement", "mark")}}}
}

// Menu returns a "menu" element with type *dom.HTMLMenuElement
func Menu() *dom.HTMLMenuElement {
	return &dom.HTMLMenuElement{BasicHTMLElement: &dom.BasicHTMLElement{&dom.BasicElement{&dom.BasicNode{js.Global.Get("document").Call("createElement", "menu")}}}}
}

// Menuitem returns a "menuitem" element with type *dom.BasicHTMLElement
func Menuitem() *dom.BasicHTMLElement {
	return &dom.BasicHTMLElement{&dom.BasicElement{&dom.BasicNode{js.Global.Get("document").Call("createElement", "menuitem")}}}
}

// Meta returns a "meta" element with type *dom.HTMLMetaElement
func Meta() *dom.HTMLMetaElement {
	return &dom.HTMLMetaElement{BasicHTMLElement: &dom.BasicHTMLElement{&dom.BasicElement{&dom.BasicNode{js.Global.Get("document").Call("createElement", "meta")}}}}
}

// Meter returns a "meter" element with type *dom.HTMLMeterElement
func Meter() *dom.HTMLMeterElement {
	return &dom.HTMLMeterElement{BasicHTMLElement: &dom.BasicHTMLElement{&dom.BasicElement{&dom.BasicNode{js.Global.Get("document").Call("createElement", "meter")}}}}
}

// Nav returns a "nav" element with type *dom.BasicHTMLElement
func Nav() *dom.BasicHTMLElement {
	return &dom.BasicHTMLElement{&dom.BasicElement{&dom.BasicNode{js.Global.Get("document").Call("createElement", "nav")}}}
}

// Noframes returns a "noframes" element with type *dom.BasicHTMLElement
func Noframes() *dom.BasicHTMLElement {
	return &dom.BasicHTMLElement{&dom.BasicElement{&dom.BasicNode{js.Global.Get("document").Call("createElement", "noframes")}}}
}

// Noscript returns a "noscript" element with type *dom.BasicHTMLElement
func Noscript() *dom.BasicHTMLElement {
	return &dom.BasicHTMLElement{&dom.BasicElement{&dom.BasicNode{js.Global.Get("document").Call("createElement", "noscript")}}}
}

// Object returns a "object" element with type *dom.HTMLObjectElement
func Object() *dom.HTMLObjectElement {
	return &dom.HTMLObjectElement{BasicHTMLElement: &dom.BasicHTMLElement{&dom.BasicElement{&dom.BasicNode{js.Global.Get("document").Call("createElement", "object")}}}}
}

// Ol returns a "ol" element with type *dom.HTMLOListElement
func Ol() *dom.HTMLOListElement {
	return &dom.HTMLOListElement{BasicHTMLElement: &dom.BasicHTMLElement{&dom.BasicElement{&dom.BasicNode{js.Global.Get("document").Call("createElement", "ol")}}}}
}

// Optgroup returns a "optgroup" element with type *dom.HTMLOptGroupElement
func Optgroup() *dom.HTMLOptGroupElement {
	return &dom.HTMLOptGroupElement{BasicHTMLElement: &dom.BasicHTMLElement{&dom.BasicElement{&dom.BasicNode{js.Global.Get("document").Call("createElement", "optgroup")}}}}
}

// Option returns a "option" element with type *dom.HTMLOptionElement
func Option() *dom.HTMLOptionElement {
	return &dom.HTMLOptionElement{BasicHTMLElement: &dom.BasicHTMLElement{&dom.BasicElement{&dom.BasicNode{js.Global.Get("document").Call("createElement", "option")}}}}
}

// Output returns a "output" element with type *dom.HTMLOutputElement
func Output() *dom.HTMLOutputElement {
	return &dom.HTMLOutputElement{BasicHTMLElement: &dom.BasicHTMLElement{&dom.BasicElement{&dom.BasicNode{js.Global.Get("document").Call("createElement", "output")}}}}
}

// P returns a "p" element with type *dom.HTMLParagraphElement
func P() *dom.HTMLParagraphElement {
	return &dom.HTMLParagraphElement{BasicHTMLElement: &dom.BasicHTMLElement{&dom.BasicElement{&dom.BasicNode{js.Global.Get("document").Call("createElement", "p")}}}}
}

// Param returns a "param" element with type *dom.HTMLParamElement
func Param() *dom.HTMLParamElement {
	return &dom.HTMLParamElement{BasicHTMLElement: &dom.BasicHTMLElement{&dom.BasicElement{&dom.BasicNode{js.Global.Get("document").Call("createElement", "param")}}}}
}

// Pre returns a "pre" element with type *dom.HTMLPreElement
func Pre() *dom.HTMLPreElement {
	return &dom.HTMLPreElement{BasicHTMLElement: &dom.BasicHTMLElement{&dom.BasicElement{&dom.BasicNode{js.Global.Get("document").Call("createElement", "pre")}}}}
}

// Progress returns a "progress" element with type *dom.HTMLProgressElement
func Progress() *dom.HTMLProgressElement {
	return &dom.HTMLProgressElement{BasicHTMLElement: &dom.BasicHTMLElement{&dom.BasicElement{&dom.BasicNode{js.Global.Get("document").Call("createElement", "progress")}}}}
}

// Q returns a "q" element with type *dom.HTMLQuoteElement
func Q() *dom.HTMLQuoteElement {
	return &dom.HTMLQuoteElement{BasicHTMLElement: &dom.BasicHTMLElement{&dom.BasicElement{&dom.BasicNode{js.Global.Get("document").Call("createElement", "q")}}}}
}

// Rp returns a "rp" element with type *dom.BasicHTMLElement
func Rp() *dom.BasicHTMLElement {
	return &dom.BasicHTMLElement{&dom.BasicElement{&dom.BasicNode{js.Global.Get("document").Call("createElement", "rp")}}}
}

// Rt returns a "rt" element with type *dom.BasicHTMLElement
func Rt() *dom.BasicHTMLElement {
	return &dom.BasicHTMLElement{&dom.BasicElement{&dom.BasicNode{js.Global.Get("document").Call("createElement", "rt")}}}
}

// Ruby returns a "ruby" element with type *dom.BasicHTMLElement
func Ruby() *dom.BasicHTMLElement {
	return &dom.BasicHTMLElement{&dom.BasicElement{&dom.BasicNode{js.Global.Get("document").Call("createElement", "ruby")}}}
}

// S returns a "s" element with type *dom.BasicHTMLElement
func S() *dom.BasicHTMLElement {
	return &dom.BasicHTMLElement{&dom.BasicElement{&dom.BasicNode{js.Global.Get("document").Call("createElement", "s")}}}
}

// Samp returns a "samp" element with type *dom.BasicHTMLElement
func Samp() *dom.BasicHTMLElement {
	return &dom.BasicHTMLElement{&dom.BasicElement{&dom.BasicNode{js.Global.Get("document").Call("createElement", "samp")}}}
}

// Script returns a "script" element with type *dom.HTMLScriptElement
func Script() *dom.HTMLScriptElement {
	return &dom.HTMLScriptElement{BasicHTMLElement: &dom.BasicHTMLElement{&dom.BasicElement{&dom.BasicNode{js.Global.Get("document").Call("createElement", "script")}}}}
}

// Section returns a "section" element with type *dom.BasicHTMLElement
func Section() *dom.BasicHTMLElement {
	return &dom.BasicHTMLElement{&dom.BasicElement{&dom.BasicNode{js.Global.Get("document").Call("createElement", "section")}}}
}

// Select returns a "select" element with type *dom.HTMLSelectElement
func Select() *dom.HTMLSelectElement {
	return &dom.HTMLSelectElement{BasicHTMLElement: &dom.BasicHTMLElement{&dom.BasicElement{&dom.BasicNode{js.Global.Get("document").Call("createElement", "select")}}}}
}

// Shadow returns a "shadow" element with type *dom.BasicHTMLElement
func Shadow() *dom.BasicHTMLElement {
	return &dom.BasicHTMLElement{&dom.BasicElement{&dom.BasicNode{js.Global.Get("document").Call("createElement", "shadow")}}}
}

// Small returns a "small" element with type *dom.BasicHTMLElement
func Small() *dom.BasicHTMLElement {
	return &dom.BasicHTMLElement{&dom.BasicElement{&dom.BasicNode{js.Global.Get("document").Call("createElement", "small")}}}
}

// Source returns a "source" element with type *dom.HTMLSourceElement
func Source() *dom.HTMLSourceElement {
	return &dom.HTMLSourceElement{BasicHTMLElement: &dom.BasicHTMLElement{&dom.BasicElement{&dom.BasicNode{js.Global.Get("document").Call("createElement", "source")}}}}
}

// Span returns a "span" element with type *dom.HTMLSpanElement
func Span() *dom.HTMLSpanElement {
	return &dom.HTMLSpanElement{BasicHTMLElement: &dom.BasicHTMLElement{&dom.BasicElement{&dom.BasicNode{js.Global.Get("document").Call("createElement", "span")}}}}
}

// Strike returns a "strike" element with type *dom.BasicHTMLElement
func Strike() *dom.BasicHTMLElement {
	return &dom.BasicHTMLElement{&dom.BasicElement{&dom.BasicNode{js.Global.Get("document").Call("createElement", "strike")}}}
}

// Strong returns a "strong" element with type *dom.BasicHTMLElement
func Strong() *dom.BasicHTMLElement {
	return &dom.BasicHTMLElement{&dom.BasicElement{&dom.BasicNode{js.Global.Get("document").Call("createElement", "strong")}}}
}

// Style returns a "style" element with type *dom.HTMLStyleElement
func Style() *dom.HTMLStyleElement {
	return &dom.HTMLStyleElement{BasicHTMLElement: &dom.BasicHTMLElement{&dom.BasicElement{&dom.BasicNode{js.Global.Get("document").Call("createElement", "style")}}}}
}

// Sub returns a "sub" element with type *dom.BasicHTMLElement
func Sub() *dom.BasicHTMLElement {
	return &dom.BasicHTMLElement{&dom.BasicElement{&dom.BasicNode{js.Global.Get("document").Call("createElement", "sub")}}}
}

// Summary returns a "summary" element with type *dom.BasicHTMLElement
func Summary() *dom.BasicHTMLElement {
	return &dom.BasicHTMLElement{&dom.BasicElement{&dom.BasicNode{js.Global.Get("document").Call("createElement", "summary")}}}
}

// Sup returns a "sup" element with type *dom.BasicHTMLElement
func Sup() *dom.BasicHTMLElement {
	return &dom.BasicHTMLElement{&dom.BasicElement{&dom.BasicNode{js.Global.Get("document").Call("createElement", "sup")}}}
}

// Table returns a "table" element with type *dom.HTMLTableElement
func Table() *dom.HTMLTableElement {
	return &dom.HTMLTableElement{BasicHTMLElement: &dom.BasicHTMLElement{&dom.BasicElement{&dom.BasicNode{js.Global.Get("document").Call("createElement", "table")}}}}
}

// Tbody returns a "tbody" element with type *dom.HTMLTableSectionElement
func Tbody() *dom.HTMLTableSectionElement {
	return &dom.HTMLTableSectionElement{BasicHTMLElement: &dom.BasicHTMLElement{&dom.BasicElement{&dom.BasicNode{js.Global.Get("document").Call("createElement", "tbody")}}}}
}

// Td returns a "td" element with type *dom.HTMLTableCellElement
func Td() *dom.HTMLTableCellElement {
	return &dom.HTMLTableCellElement{BasicHTMLElement: &dom.BasicHTMLElement{&dom.BasicElement{&dom.BasicNode{js.Global.Get("document").Call("createElement", "td")}}}}
}

// Template returns a "template" element with type *dom.BasicHTMLElement
func Template() *dom.BasicHTMLElement {
	return &dom.BasicHTMLElement{&dom.BasicElement{&dom.BasicNode{js.Global.Get("document").Call("createElement", "template")}}}
}

// Textarea returns a "textarea" element with type *dom.HTMLTextAreaElement
func Textarea() *dom.HTMLTextAreaElement {
	return &dom.HTMLTextAreaElement{BasicHTMLElement: &dom.BasicHTMLElement{&dom.BasicElement{&dom.BasicNode{js.Global.Get("document").Call("createElement", "textarea")}}}}
}

// Tfoot returns a "tfoot" element with type *dom.HTMLTableSectionElement
func Tfoot() *dom.HTMLTableSectionElement {
	return &dom.HTMLTableSectionElement{BasicHTMLElement: &dom.BasicHTMLElement{&dom.BasicElement{&dom.BasicNode{js.Global.Get("document").Call("createElement", "tfoot")}}}}
}

// Th returns a "th" element with type *dom.HTMLTableCellElement
func Th() *dom.HTMLTableCellElement {
	return &dom.HTMLTableCellElement{BasicHTMLElement: &dom.BasicHTMLElement{&dom.BasicElement{&dom.BasicNode{js.Global.Get("document").Call("createElement", "th")}}}}
}

// Thead returns a "thead" element with type *dom.HTMLTableSectionElement
func Thead() *dom.HTMLTableSectionElement {
	return &dom.HTMLTableSectionElement{BasicHTMLElement: &dom.BasicHTMLElement{&dom.BasicElement{&dom.BasicNode{js.Global.Get("document").Call("createElement", "thead")}}}}
}

// Time returns a "time" element with type *dom.HTMLTimeElement
func Time() *dom.HTMLTimeElement {
	return &dom.HTMLTimeElement{BasicHTMLElement: &dom.BasicHTMLElement{&dom.BasicElement{&dom.BasicNode{js.Global.Get("document").Call("createElement", "time")}}}}
}

// Title returns a "title" element with type *dom.HTMLTitleElement
func Title() *dom.HTMLTitleElement {
	return &dom.HTMLTitleElement{BasicHTMLElement: &dom.BasicHTMLElement{&dom.BasicElement{&dom.BasicNode{js.Global.Get("document").Call("createElement", "title")}}}}
}

// Tr returns a "tr" element with type *dom.HTMLTableRowElement
func Tr() *dom.HTMLTableRowElement {
	return &dom.HTMLTableRowElement{BasicHTMLElement: &dom.BasicHTMLElement{&dom.BasicElement{&dom.BasicNode{js.Global.Get("document").Call("createElement", "tr")}}}}
}

// Track returns a "track" element with type *dom.HTMLTrackElement
func Track() *dom.HTMLTrackElement {
	return &dom.HTMLTrackElement{BasicHTMLElement: &dom.BasicHTMLElement{&dom.BasicElement{&dom.BasicNode{js.Global.Get("document").Call("createElement", "track")}}}}
}

// Tt returns a "tt" element with type *dom.BasicHTMLElement
func Tt() *dom.BasicHTMLElement {
	return &dom.BasicHTMLElement{&dom.BasicElement{&dom.BasicNode{js.Global.Get("document").Call("createElement", "tt")}}}
}

// U returns a "u" element with type *dom.BasicHTMLElement
func U() *dom.BasicHTMLElement {
	return &dom.BasicHTMLElement{&dom.BasicElement{&dom.BasicNode{js.Global.Get("document").Call("createElement", "u")}}}
}

// Ul returns a "ul" element with type *dom.HTMLUListElement
func Ul() *dom.HTMLUListElement {
	return &dom.HTMLUListElement{BasicHTMLElement: &dom.BasicHTMLElement{&dom.BasicElement{&dom.BasicNode{js.Global.Get("document").Call("createElement", "ul")}}}}
}

// Var returns a "var" element with type *dom.BasicHTMLElement
func Var() *dom.BasicHTMLElement {
	return &dom.BasicHTMLElement{&dom.BasicElement{&dom.BasicNode{js.Global.Get("document").Call("createElement", "var")}}}
}

// Video returns a "video" element with type *dom.HTMLVideoElement
func Video() *dom.HTMLVideoElement {
	return &dom.HTMLVideoElement{HTMLMediaElement: &dom.HTMLMediaElement{BasicHTMLElement: &dom.BasicHTMLElement{&dom.BasicElement{&dom.BasicNode{js.Global.Get("document").Call("createElement", "video")}}}}}
}

// Wbr returns a "wbr" element with type *dom.BasicHTMLElement
func Wbr() *dom.BasicHTMLElement {
	return &dom.BasicHTMLElement{&dom.BasicElement{&dom.BasicNode{js.Global.Get("document").Call("createElement", "wbr")}}}
}
