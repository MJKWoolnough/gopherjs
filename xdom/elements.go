// Package xdom provides simple functions to create HTML elements with the correct type
package xdom

// File automatically generated with ./genMap.sh

import "honnef.co/go/js/dom"

// A returns a "a" element with type *dom.HTMLAnchorElement
func A() *dom.HTMLAnchorElement {
	return dom.GetWindow().Document().CreateElement("a").(*dom.HTMLAnchorElement)
}

// Abbr returns a "abbr" element with type dom.Element
func Abbr() dom.Element {
	return dom.GetWindow().Document().CreateElement("abbr")
}

// Acronym returns a "acronym" element with type dom.Element
func Acronym() dom.Element {
	return dom.GetWindow().Document().CreateElement("acronym")
}

// Address returns a "address" element with type dom.Element
func Address() dom.Element {
	return dom.GetWindow().Document().CreateElement("address")
}

// Applet returns a "applet" element with type *dom.HTMLAppletElement
func Applet() *dom.HTMLAppletElement {
	return dom.GetWindow().Document().CreateElement("applet").(*dom.HTMLAppletElement)
}

// Area returns a "area" element with type *dom.HTMLAreaElement
func Area() *dom.HTMLAreaElement {
	return dom.GetWindow().Document().CreateElement("area").(*dom.HTMLAreaElement)
}

// Article returns a "article" element with type dom.Element
func Article() dom.Element {
	return dom.GetWindow().Document().CreateElement("article")
}

// Aside returns a "aside" element with type dom.Element
func Aside() dom.Element {
	return dom.GetWindow().Document().CreateElement("aside")
}

// Audio returns a "audio" element with type *dom.HTMLAudioElement
func Audio() *dom.HTMLAudioElement {
	return dom.GetWindow().Document().CreateElement("audio").(*dom.HTMLAudioElement)
}

// B returns a "b" element with type dom.Element
func B() dom.Element {
	return dom.GetWindow().Document().CreateElement("b")
}

// Base returns a "base" element with type *dom.HTMLBaseElement
func Base() *dom.HTMLBaseElement {
	return dom.GetWindow().Document().CreateElement("base").(*dom.HTMLBaseElement)
}

// Basefont returns a "basefont" element with type dom.Element
func Basefont() dom.Element {
	return dom.GetWindow().Document().CreateElement("basefont")
}

// Bdi returns a "bdi" element with type dom.Element
func Bdi() dom.Element {
	return dom.GetWindow().Document().CreateElement("bdi")
}

// Bdo returns a "bdo" element with type dom.Element
func Bdo() dom.Element {
	return dom.GetWindow().Document().CreateElement("bdo")
}

// Big returns a "big" element with type dom.Element
func Big() dom.Element {
	return dom.GetWindow().Document().CreateElement("big")
}

// Blockquote returns a "blockquote" element with type *dom.HTMLQuoteElement
func Blockquote() *dom.HTMLQuoteElement {
	return dom.GetWindow().Document().CreateElement("blockquote").(*dom.HTMLQuoteElement)
}

// Body returns a "body" element with type *dom.HTMLBodyElement
func Body() *dom.HTMLBodyElement {
	return dom.GetWindow().Document().CreateElement("body").(*dom.HTMLBodyElement)
}

// Br returns a "br" element with type *dom.HTMLBRElement
func Br() *dom.HTMLBRElement {
	return dom.GetWindow().Document().CreateElement("br").(*dom.HTMLBRElement)
}

// Button returns a "button" element with type *dom.HTMLButtonElement
func Button() *dom.HTMLButtonElement {
	return dom.GetWindow().Document().CreateElement("button").(*dom.HTMLButtonElement)
}

// Canvas returns a "canvas" element with type *dom.HTMLCanvasElement
func Canvas() *dom.HTMLCanvasElement {
	return dom.GetWindow().Document().CreateElement("canvas").(*dom.HTMLCanvasElement)
}

// Caption returns a "caption" element with type *dom.HTMLTableCaptionElement
func Caption() *dom.HTMLTableCaptionElement {
	return dom.GetWindow().Document().CreateElement("caption").(*dom.HTMLTableCaptionElement)
}

// Center returns a "center" element with type dom.Element
func Center() dom.Element {
	return dom.GetWindow().Document().CreateElement("center")
}

// Cite returns a "cite" element with type *dom.HTMLQuoteElement
func Cite() *dom.HTMLQuoteElement {
	return dom.GetWindow().Document().CreateElement("cite").(*dom.HTMLQuoteElement)
}

// Code returns a "code" element with type dom.Element
func Code() dom.Element {
	return dom.GetWindow().Document().CreateElement("code")
}

// Col returns a "col" element with type *dom.HTMLTableColElement
func Col() *dom.HTMLTableColElement {
	return dom.GetWindow().Document().CreateElement("col").(*dom.HTMLTableColElement)
}

// Colgroup returns a "colgroup" element with type *dom.HTMLTableColElement
func Colgroup() *dom.HTMLTableColElement {
	return dom.GetWindow().Document().CreateElement("colgroup").(*dom.HTMLTableColElement)
}

// Content returns a "content" element with type dom.Element
func Content() dom.Element {
	return dom.GetWindow().Document().CreateElement("content")
}

// Data returns a "data" element with type *dom.HTMLDataElement
func Data() *dom.HTMLDataElement {
	return dom.GetWindow().Document().CreateElement("data").(*dom.HTMLDataElement)
}

// Datalist returns a "datalist" element with type *dom.HTMLDataListElement
func Datalist() *dom.HTMLDataListElement {
	return dom.GetWindow().Document().CreateElement("datalist").(*dom.HTMLDataListElement)
}

// Dd returns a "dd" element with type dom.Element
func Dd() dom.Element {
	return dom.GetWindow().Document().CreateElement("dd")
}

// Del returns a "del" element with type *dom.HTMLModElement
func Del() *dom.HTMLModElement {
	return dom.GetWindow().Document().CreateElement("del").(*dom.HTMLModElement)
}

// Details returns a "details" element with type dom.Element
func Details() dom.Element {
	return dom.GetWindow().Document().CreateElement("details")
}

// Dfn returns a "dfn" element with type dom.Element
func Dfn() dom.Element {
	return dom.GetWindow().Document().CreateElement("dfn")
}

// Dialog returns a "dialog" element with type dom.Element
func Dialog() dom.Element {
	return dom.GetWindow().Document().CreateElement("dialog")
}

// Dir returns a "dir" element with type *dom.HTMLDirectoryElement
func Dir() *dom.HTMLDirectoryElement {
	return dom.GetWindow().Document().CreateElement("dir").(*dom.HTMLDirectoryElement)
}

// Div returns a "div" element with type *dom.HTMLDivElement
func Div() *dom.HTMLDivElement {
	return dom.GetWindow().Document().CreateElement("div").(*dom.HTMLDivElement)
}

// Dl returns a "dl" element with type *dom.HTMLDListElement
func Dl() *dom.HTMLDListElement {
	return dom.GetWindow().Document().CreateElement("dl").(*dom.HTMLDListElement)
}

// Dt returns a "dt" element with type dom.Element
func Dt() dom.Element {
	return dom.GetWindow().Document().CreateElement("dt")
}

// Em returns a "em" element with type dom.Element
func Em() dom.Element {
	return dom.GetWindow().Document().CreateElement("em")
}

// Embed returns a "embed" element with type *dom.HTMLEmbedElement
func Embed() *dom.HTMLEmbedElement {
	return dom.GetWindow().Document().CreateElement("embed").(*dom.HTMLEmbedElement)
}

// Fieldset returns a "fieldset" element with type *dom.HTMLFieldSetElement
func Fieldset() *dom.HTMLFieldSetElement {
	return dom.GetWindow().Document().CreateElement("fieldset").(*dom.HTMLFieldSetElement)
}

// Figcaption returns a "figcaption" element with type dom.Element
func Figcaption() dom.Element {
	return dom.GetWindow().Document().CreateElement("figcaption")
}

// Figure returns a "figure" element with type dom.Element
func Figure() dom.Element {
	return dom.GetWindow().Document().CreateElement("figure")
}

// Font returns a "font" element with type *dom.HTMLFontElement
func Font() *dom.HTMLFontElement {
	return dom.GetWindow().Document().CreateElement("font").(*dom.HTMLFontElement)
}

// Footer returns a "footer" element with type dom.Element
func Footer() dom.Element {
	return dom.GetWindow().Document().CreateElement("footer")
}

// Form returns a "form" element with type *dom.HTMLFormElement
func Form() *dom.HTMLFormElement {
	return dom.GetWindow().Document().CreateElement("form").(*dom.HTMLFormElement)
}

// Frame returns a "frame" element with type *dom.HTMLFrameElement
func Frame() *dom.HTMLFrameElement {
	return dom.GetWindow().Document().CreateElement("frame").(*dom.HTMLFrameElement)
}

// Frameset returns a "frameset" element with type *dom.HTMLFrameSetElement
func Frameset() *dom.HTMLFrameSetElement {
	return dom.GetWindow().Document().CreateElement("frameset").(*dom.HTMLFrameSetElement)
}

// H1 returns a "h1" element with type *dom.HTMLHeadingElement
func H1() *dom.HTMLHeadingElement {
	return dom.GetWindow().Document().CreateElement("h1").(*dom.HTMLHeadingElement)
}

// H2 returns a "h2" element with type *dom.HTMLHeadingElement
func H2() *dom.HTMLHeadingElement {
	return dom.GetWindow().Document().CreateElement("h2").(*dom.HTMLHeadingElement)
}

// H3 returns a "h3" element with type *dom.HTMLHeadingElement
func H3() *dom.HTMLHeadingElement {
	return dom.GetWindow().Document().CreateElement("h3").(*dom.HTMLHeadingElement)
}

// H4 returns a "h4" element with type *dom.HTMLHeadingElement
func H4() *dom.HTMLHeadingElement {
	return dom.GetWindow().Document().CreateElement("h4").(*dom.HTMLHeadingElement)
}

// H5 returns a "h5" element with type *dom.HTMLHeadingElement
func H5() *dom.HTMLHeadingElement {
	return dom.GetWindow().Document().CreateElement("h5").(*dom.HTMLHeadingElement)
}

// H6 returns a "h6" element with type *dom.HTMLHeadingElement
func H6() *dom.HTMLHeadingElement {
	return dom.GetWindow().Document().CreateElement("h6").(*dom.HTMLHeadingElement)
}

// Head returns a "head" element with type *dom.HTMLHeadElement
func Head() *dom.HTMLHeadElement {
	return dom.GetWindow().Document().CreateElement("head").(*dom.HTMLHeadElement)
}

// Header returns a "header" element with type dom.Element
func Header() dom.Element {
	return dom.GetWindow().Document().CreateElement("header")
}

// Hr returns a "hr" element with type *dom.HTMLHRElement
func Hr() *dom.HTMLHRElement {
	return dom.GetWindow().Document().CreateElement("hr").(*dom.HTMLHRElement)
}

// Html returns a "html" element with type *dom.HTMLHtmlElement
func Html() *dom.HTMLHtmlElement {
	return dom.GetWindow().Document().CreateElement("html").(*dom.HTMLHtmlElement)
}

// I returns a "i" element with type dom.Element
func I() dom.Element {
	return dom.GetWindow().Document().CreateElement("i")
}

// Iframe returns a "iframe" element with type *dom.HTMLIFrameElement
func Iframe() *dom.HTMLIFrameElement {
	return dom.GetWindow().Document().CreateElement("iframe").(*dom.HTMLIFrameElement)
}

// Img returns a "img" element with type *dom.HTMLImageElement
func Img() *dom.HTMLImageElement {
	return dom.GetWindow().Document().CreateElement("img").(*dom.HTMLImageElement)
}

// Input returns a "input" element with type *dom.HTMLInputElement
func Input() *dom.HTMLInputElement {
	return dom.GetWindow().Document().CreateElement("input").(*dom.HTMLInputElement)
}

// Ins returns a "ins" element with type *dom.HTMLModElement
func Ins() *dom.HTMLModElement {
	return dom.GetWindow().Document().CreateElement("ins").(*dom.HTMLModElement)
}

// Kbd returns a "kbd" element with type dom.Element
func Kbd() dom.Element {
	return dom.GetWindow().Document().CreateElement("kbd")
}

// Keygen returns a "keygen" element with type *dom.HTMLKeygenElement
func Keygen() *dom.HTMLKeygenElement {
	return dom.GetWindow().Document().CreateElement("keygen").(*dom.HTMLKeygenElement)
}

// Label returns a "label" element with type *dom.HTMLLabelElement
func Label() *dom.HTMLLabelElement {
	return dom.GetWindow().Document().CreateElement("label").(*dom.HTMLLabelElement)
}

// Legend returns a "legend" element with type *dom.HTMLLegendElement
func Legend() *dom.HTMLLegendElement {
	return dom.GetWindow().Document().CreateElement("legend").(*dom.HTMLLegendElement)
}

// Li returns a "li" element with type *dom.HTMLLIElement
func Li() *dom.HTMLLIElement {
	return dom.GetWindow().Document().CreateElement("li").(*dom.HTMLLIElement)
}

// Link returns a "link" element with type *dom.HTMLLinkElement
func Link() *dom.HTMLLinkElement {
	return dom.GetWindow().Document().CreateElement("link").(*dom.HTMLLinkElement)
}

// Main returns a "main" element with type dom.Element
func Main() dom.Element {
	return dom.GetWindow().Document().CreateElement("main")
}

// Map returns a "map" element with type *dom.HTMLMapElement
func Map() *dom.HTMLMapElement {
	return dom.GetWindow().Document().CreateElement("map").(*dom.HTMLMapElement)
}

// Mark returns a "mark" element with type dom.Element
func Mark() dom.Element {
	return dom.GetWindow().Document().CreateElement("mark")
}

// Menu returns a "menu" element with type *dom.HTMLMenuElement
func Menu() *dom.HTMLMenuElement {
	return dom.GetWindow().Document().CreateElement("menu").(*dom.HTMLMenuElement)
}

// Menuitem returns a "menuitem" element with type dom.Element
func Menuitem() dom.Element {
	return dom.GetWindow().Document().CreateElement("menuitem")
}

// Meta returns a "meta" element with type *dom.HTMLMetaElement
func Meta() *dom.HTMLMetaElement {
	return dom.GetWindow().Document().CreateElement("meta").(*dom.HTMLMetaElement)
}

// Meter returns a "meter" element with type *dom.HTMLMeterElement
func Meter() *dom.HTMLMeterElement {
	return dom.GetWindow().Document().CreateElement("meter").(*dom.HTMLMeterElement)
}

// Nav returns a "nav" element with type dom.Element
func Nav() dom.Element {
	return dom.GetWindow().Document().CreateElement("nav")
}

// Noframes returns a "noframes" element with type dom.Element
func Noframes() dom.Element {
	return dom.GetWindow().Document().CreateElement("noframes")
}

// Noscript returns a "noscript" element with type dom.Element
func Noscript() dom.Element {
	return dom.GetWindow().Document().CreateElement("noscript")
}

// Object returns a "object" element with type *dom.HTMLObjectElement
func Object() *dom.HTMLObjectElement {
	return dom.GetWindow().Document().CreateElement("object").(*dom.HTMLObjectElement)
}

// Ol returns a "ol" element with type *dom.HTMLOListElement
func Ol() *dom.HTMLOListElement {
	return dom.GetWindow().Document().CreateElement("ol").(*dom.HTMLOListElement)
}

// Optgroup returns a "optgroup" element with type *dom.HTMLOptGroupElement
func Optgroup() *dom.HTMLOptGroupElement {
	return dom.GetWindow().Document().CreateElement("optgroup").(*dom.HTMLOptGroupElement)
}

// Option returns a "option" element with type *dom.HTMLOptionElement
func Option() *dom.HTMLOptionElement {
	return dom.GetWindow().Document().CreateElement("option").(*dom.HTMLOptionElement)
}

// Output returns a "output" element with type *dom.HTMLOutputElement
func Output() *dom.HTMLOutputElement {
	return dom.GetWindow().Document().CreateElement("output").(*dom.HTMLOutputElement)
}

// P returns a "p" element with type *dom.HTMLParagraphElement
func P() *dom.HTMLParagraphElement {
	return dom.GetWindow().Document().CreateElement("p").(*dom.HTMLParagraphElement)
}

// Param returns a "param" element with type *dom.HTMLParamElement
func Param() *dom.HTMLParamElement {
	return dom.GetWindow().Document().CreateElement("param").(*dom.HTMLParamElement)
}

// Pre returns a "pre" element with type *dom.HTMLPreElement
func Pre() *dom.HTMLPreElement {
	return dom.GetWindow().Document().CreateElement("pre").(*dom.HTMLPreElement)
}

// Progress returns a "progress" element with type *dom.HTMLProgressElement
func Progress() *dom.HTMLProgressElement {
	return dom.GetWindow().Document().CreateElement("progress").(*dom.HTMLProgressElement)
}

// Q returns a "q" element with type *dom.HTMLQuoteElement
func Q() *dom.HTMLQuoteElement {
	return dom.GetWindow().Document().CreateElement("q").(*dom.HTMLQuoteElement)
}

// Rp returns a "rp" element with type dom.Element
func Rp() dom.Element {
	return dom.GetWindow().Document().CreateElement("rp")
}

// Rt returns a "rt" element with type dom.Element
func Rt() dom.Element {
	return dom.GetWindow().Document().CreateElement("rt")
}

// Ruby returns a "ruby" element with type dom.Element
func Ruby() dom.Element {
	return dom.GetWindow().Document().CreateElement("ruby")
}

// S returns a "s" element with type dom.Element
func S() dom.Element {
	return dom.GetWindow().Document().CreateElement("s")
}

// Samp returns a "samp" element with type dom.Element
func Samp() dom.Element {
	return dom.GetWindow().Document().CreateElement("samp")
}

// Script returns a "script" element with type *dom.HTMLScriptElement
func Script() *dom.HTMLScriptElement {
	return dom.GetWindow().Document().CreateElement("script").(*dom.HTMLScriptElement)
}

// Section returns a "section" element with type dom.Element
func Section() dom.Element {
	return dom.GetWindow().Document().CreateElement("section")
}

// Select returns a "select" element with type *dom.HTMLSelectElement
func Select() *dom.HTMLSelectElement {
	return dom.GetWindow().Document().CreateElement("select").(*dom.HTMLSelectElement)
}

// Shadow returns a "shadow" element with type dom.Element
func Shadow() dom.Element {
	return dom.GetWindow().Document().CreateElement("shadow")
}

// Small returns a "small" element with type dom.Element
func Small() dom.Element {
	return dom.GetWindow().Document().CreateElement("small")
}

// Source returns a "source" element with type *dom.HTMLSourceElement
func Source() *dom.HTMLSourceElement {
	return dom.GetWindow().Document().CreateElement("source").(*dom.HTMLSourceElement)
}

// Span returns a "span" element with type *dom.HTMLSpanElement
func Span() *dom.HTMLSpanElement {
	return dom.GetWindow().Document().CreateElement("span").(*dom.HTMLSpanElement)
}

// Strike returns a "strike" element with type dom.Element
func Strike() dom.Element {
	return dom.GetWindow().Document().CreateElement("strike")
}

// Strong returns a "strong" element with type dom.Element
func Strong() dom.Element {
	return dom.GetWindow().Document().CreateElement("strong")
}

// Style returns a "style" element with type *dom.HTMLStyleElement
func Style() *dom.HTMLStyleElement {
	return dom.GetWindow().Document().CreateElement("style").(*dom.HTMLStyleElement)
}

// Sub returns a "sub" element with type dom.Element
func Sub() dom.Element {
	return dom.GetWindow().Document().CreateElement("sub")
}

// Summary returns a "summary" element with type dom.Element
func Summary() dom.Element {
	return dom.GetWindow().Document().CreateElement("summary")
}

// Sup returns a "sup" element with type dom.Element
func Sup() dom.Element {
	return dom.GetWindow().Document().CreateElement("sup")
}

// Table returns a "table" element with type *dom.HTMLTableElement
func Table() *dom.HTMLTableElement {
	return dom.GetWindow().Document().CreateElement("table").(*dom.HTMLTableElement)
}

// Tbody returns a "tbody" element with type *dom.HTMLTableSectionElement
func Tbody() *dom.HTMLTableSectionElement {
	return dom.GetWindow().Document().CreateElement("tbody").(*dom.HTMLTableSectionElement)
}

// Td returns a "td" element with type *dom.HTMLTableCellElement
func Td() *dom.HTMLTableCellElement {
	return dom.GetWindow().Document().CreateElement("td").(*dom.HTMLTableCellElement)
}

// Template returns a "template" element with type dom.Element
func Template() dom.Element {
	return dom.GetWindow().Document().CreateElement("template")
}

// Textarea returns a "textarea" element with type *dom.HTMLTextAreaElement
func Textarea() *dom.HTMLTextAreaElement {
	return dom.GetWindow().Document().CreateElement("textarea").(*dom.HTMLTextAreaElement)
}

// Tfoot returns a "tfoot" element with type *dom.HTMLTableSectionElement
func Tfoot() *dom.HTMLTableSectionElement {
	return dom.GetWindow().Document().CreateElement("tfoot").(*dom.HTMLTableSectionElement)
}

// Th returns a "th" element with type *dom.HTMLTableCellElement
func Th() *dom.HTMLTableCellElement {
	return dom.GetWindow().Document().CreateElement("th").(*dom.HTMLTableCellElement)
}

// Thead returns a "thead" element with type *dom.HTMLTableSectionElement
func Thead() *dom.HTMLTableSectionElement {
	return dom.GetWindow().Document().CreateElement("thead").(*dom.HTMLTableSectionElement)
}

// Time returns a "time" element with type *dom.HTMLTimeElement
func Time() *dom.HTMLTimeElement {
	return dom.GetWindow().Document().CreateElement("time").(*dom.HTMLTimeElement)
}

// Title returns a "title" element with type *dom.HTMLTitleElement
func Title() *dom.HTMLTitleElement {
	return dom.GetWindow().Document().CreateElement("title").(*dom.HTMLTitleElement)
}

// Tr returns a "tr" element with type *dom.HTMLTableRowElement
func Tr() *dom.HTMLTableRowElement {
	return dom.GetWindow().Document().CreateElement("tr").(*dom.HTMLTableRowElement)
}

// Track returns a "track" element with type *dom.HTMLTrackElement
func Track() *dom.HTMLTrackElement {
	return dom.GetWindow().Document().CreateElement("track").(*dom.HTMLTrackElement)
}

// Tt returns a "tt" element with type dom.Element
func Tt() dom.Element {
	return dom.GetWindow().Document().CreateElement("tt")
}

// U returns a "u" element with type dom.Element
func U() dom.Element {
	return dom.GetWindow().Document().CreateElement("u")
}

// Ul returns a "ul" element with type *dom.HTMLUListElement
func Ul() *dom.HTMLUListElement {
	return dom.GetWindow().Document().CreateElement("ul").(*dom.HTMLUListElement)
}

// Var returns a "var" element with type dom.Element
func Var() dom.Element {
	return dom.GetWindow().Document().CreateElement("var")
}

// Video returns a "video" element with type *dom.HTMLVideoElement
func Video() *dom.HTMLVideoElement {
	return dom.GetWindow().Document().CreateElement("video").(*dom.HTMLVideoElement)
}

// Wbr returns a "wbr" element with type dom.Element
func Wbr() dom.Element {
	return dom.GetWindow().Document().CreateElement("wbr")
}
