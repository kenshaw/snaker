// Package snaker provides methods to convert CamelCase to and from snake_case.
//
// snaker takes into takes into consideration common initialisms (ie, ID, HTTP,
// ACL, etc) when converting to/from CamelCase and snake_case.
package snaker

//go:generate ./gen.sh --update

import (
	"strings"
	"unicode"
)

// CamelToSnake converts s to snake_case.
func CamelToSnake(s string) string {
	if s == "" {
		return ""
	}

	rs := []rune(s)

	var r string
	var lastWasLetter, lastWasIsm, isUpper, isLetter bool
	for i := 0; i < len(rs); {
		isUpper = unicode.IsUpper(rs[i])
		isLetter = unicode.IsLetter(rs[i])

		// append _ when last was not upper and not letter
		if (lastWasLetter && isUpper) || (lastWasIsm && isLetter) {
			r += "_"
			lastWasIsm = false
		}

		// determine next to append to r
		var next string
		if ism := peekInitialism(rs[i:]); ism != "" {
			next = ism
			lastWasIsm = true
		} else {
			next = string(rs[i])
			lastWasIsm = false
		}

		lastWasLetter = isLetter

		r += next
		i += len(next)
	}

	return strings.ToLower(r)
}

// SnakeToCamel converts s to CamelCase.
func SnakeToCamel(s string) string {
	var r string

	for _, w := range strings.Split(s, "_") {
		if w == "" {
			continue
		}

		u := strings.ToUpper(w)
		if ok := commonInitialisms[u]; ok {
			r += u
		} else {
			r += strings.ToUpper(w[:1]) + strings.ToLower(w[1:])
		}
	}

	return r
}

// SnakeToGoIdentifier converts s into a Go safe identifier (first letter will
// be capitalized).
func SnakeToGoIdentifier(s string) string {
	// replace bad chars with _
	s = replaceBadChars(s)

	// fix 2 or more __
	s = underscoreRE.ReplaceAllString(s, "_")

	// remove leading/trailing underscores
	s = strings.TrimLeft(s, "_")
	s = strings.TrimRight(s, "_")

	// convert to camel
	s = SnakeToCamel(s)

	// remove leading numbers
	s = numberRE.ReplaceAllString(s, "_")
	if s == "" {
		s = "_"
	}

	return s
}
