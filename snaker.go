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
	var lastWasUpper, lastWasLetter, lastWasIsm, isUpper, isLetter bool
	for i := 0; i < len(rs); {
		isUpper = unicode.IsUpper(rs[i])
		isLetter = unicode.IsLetter(rs[i])

		// append _ when last was not upper and not letter
		if (lastWasLetter && isUpper) || (lastWasIsm && isLetter) {
			r += "_"
		}

		// determine next to append to r
		var next string
		if ism := peekInitialism(rs[i:]); ism != "" && (!lastWasUpper || lastWasIsm) {
			next = ism
		} else {
			next = string(rs[i])
		}

		lastWasIsm = false
		if len(next) > 1 {
			lastWasIsm = true
		}

		lastWasUpper = isUpper
		lastWasLetter = isLetter

		r += next
		i += len(next)
	}

	return strings.ToLower(r)
}

// CamelToSnakeIdentifier converts s to its snake_case identifier.
func CamelToSnakeIdentifier(s string) string {
	return toIdentifier(CamelToSnake(s))
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

// SnakeToCamelIdentifier converts s to its CamelCase identifier (first
// letter is capitalized).
func SnakeToCamelIdentifier(s string) string {
	return SnakeToCamel(toIdentifier(s))
}
