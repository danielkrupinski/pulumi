// Copyright 2016 Marapongo, Inc. All rights reserved.

package aws

import (
	"unicode"
	"unicode/utf8"
)

// makeAWSFriendlyName returns a name part that is suitable for inclusion in a CloudFormation string.  This includes
// stripping out unacceptable characters, and using either PascalCasing or camelCasing, based on the camel parameter.
func makeAWSFriendlyName(s string, pascal bool) string {
	if len(s) == 0 {
		return s
	}
	t := []rune{}
	b := []byte(s)
	first := true
	capnext := pascal
	for {
		r, s := utf8.DecodeRune(b)
		if s == 0 {
			break
		}

		if (r >= 'a' && r <= 'z') || (r >= 'A' && r <= 'Z') || (r >= '0' && r <= '9') {
			if capnext {
				// We must capitalize, either due to PascalCasing, or because of a non-printable rune.
				r = unicode.ToUpper(r)
				capnext = false
			} else if first {
				// For the first letter, we'll have PascalCased (thanks to capnext), but need to camelCase.
				r = unicode.ToLower(r)
			}
			t = append(t, r)
			first = false
		} else if !first {
			capnext = true
		}

		b = b[s:]
	}

	return string(t)
}