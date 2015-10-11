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
		interface="dom.Element";
		if [ -z "${element/*:*/}" ]; then
			interface="*dom.HTML${element#*:}Element";
			element="${element%:*}";
		fi;
		cElement="${element^}";
		echo "";
		echo "// $cElement returns a \"$element\" element with type $interface"
		echo "func $cElement() $interface {";
		echo -n "	return xjs.CreateElement(\"$element\")";
		if [ "$interface" == "dom.Element" ]; then
			echo;
		else
			echo ".($interface)";
		fi;
		echo "}";
	done < "map.gen";
) > elements.go
