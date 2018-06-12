#!/bin/bash

(
	cat <<HEREDOC
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
HEREDOC
	while read element; do
		eType="BasicHTMLElement";
		special=
		if [ -z "${element/*:*/}" ]; then
			OFS="$IFS";
			IFS=":";
			data=($element);
			element="${data[0]}";
			eType="HTML${data[1]}Element";
			special="${data[2]}";
			IFS="$OFS";
		fi;
		name="${element^}";
		cat <<HEREDOC

// $name returns a "$element" element with type *dom.$eType
func $name() *dom.$eType {
HEREDOC
		if [ "$special" = "URL" ]; then
			cat <<HEREDOC
	o := js.Global.Get("document").Call("createElement", "$element")
	return &dom.$eType{BasicHTMLElement: &dom.BasicHTMLElement{&dom.BasicElement{&dom.BasicNode{o}}}, URLUtils: &dom.URLUtils{Object: o}}
}
HEREDOC
		elif [ "$special" = "MEDIA" ]; then
			cat <<HEREDOC
	return &dom.$eType{HTMLMediaElement: &dom.HTMLMediaElement{BasicHTMLElement: &dom.BasicHTMLElement{&dom.BasicElement{&dom.BasicNode{js.Global.Get("document").Call("createElement", "$element")}}}}}
}
HEREDOC
		elif [ "$eType" = "BasicHTMLElement" ]; then
			cat <<HEREDOC
	return &dom.BasicHTMLElement{&dom.BasicElement{&dom.BasicNode{js.Global.Get("document").Call("createElement", "$element")}}}
}
HEREDOC
		else
			cat <<HEREDOC
	return &dom.$eType{BasicHTMLElement: &dom.BasicHTMLElement{&dom.BasicElement{&dom.BasicNode{js.Global.Get("document").Call("createElement", "$element")}}}}
}
HEREDOC
		fi;
	done < "map.gen";
) > elements.go
