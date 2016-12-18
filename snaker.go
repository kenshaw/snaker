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

const (
	// minInitialismLen is the min length of any of the commonInitialisms
	// below.
	minInitialismLen = 2

	// maxInitialismLen is the max length of any of the commonInitialisms
	// below.
	maxInitialismLen = 5
)

// min returns the minimum of a, b.
func min(a, b int) int {
	if a < b {
		return a
	}

	return b
}

// nextInitialism returns the next longest possible initialism in rs.
func peekInitialism(rs []rune) string {
	// do no work
	if len(rs) < minInitialismLen {
		return ""
	}

	// grab at most next maxInitialismLen uppercase characters
	l := min(len(rs), maxInitialismLen)
	var z []rune
	for i := 0; i < l; i++ {
		if !unicode.IsUpper(rs[i]) {
			break
		}
		z = append(z, rs[i])
	}

	// bail if next few characters were not uppercase.
	if len(z) < minInitialismLen {
		return ""
	}

	// determine if common initialism
	for i := min(maxInitialismLen, len(z)); i >= minInitialismLen; i-- {
		if r := string(z[:i]); commonInitialisms[r] {
			return r
		}
	}

	return ""
}

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
