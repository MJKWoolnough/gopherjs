#!/bin/bash

(
	echo "// Package xdom provides simple functions to create HTML elements with the correct type";
	echo "package xdom";
	echo "";
	echo "// File automatically generated with ./genMap.sh";
	echo "";
	echo "import ("
	echo "	\"github.com/MJKWoolnough/gopherjs/xjs\"";
	echo "	\"honnef.co/go/js/dom\"";
	echo ")";

	while read element; do
		interface="";
		if [ -z "${element/*:*/}" ]; then
			interface="HTML${element#*:}Element";
			element="${element%:*}";
		fi;
		cElement="${element^}";
		echo "";
		if [ -z "$interface" ]; then
			echo "// $cElement returns an $element with type dom.Element";
			echo "func $cElement() dom.Element {";
			echo "	return xjs.CreateElement(\"$element\")";
			echo "}"
		else 
			echo "// $cElement returns an $element with type *dom.$interface";
			echo "func $cElement() *dom.$interface {";
			echo "	return xjs.CreateElement(\"$element\").(*dom.$interface)";
			echo "}"
		fi;
	done < map.gen
) > elements.go
