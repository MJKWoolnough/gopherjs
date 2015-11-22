#!/bin/bash

(
	cat <<HEREDOC
// Package xdom provides simple functions to create HTML elements with the correct type
package xdom

// File automatically generated with ./genMap.sh

import "honnef.co/go/js/dom"
HEREDOC
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
		echo -n "	return dom.GetWindow().Document().CreateElement(\"$element\")";
		if [ "$interface" == "dom.Element" ]; then
			echo;
		else
			echo ".($interface)";
		fi;
		echo "}";
	done < "map.gen";
) > elements.go
