// Package snaker provides methods to convert CamelCase to and from snake_case.
//
// snaker takes into takes into consideration common initialisms (ie, ID, HTTP,
// ACL, etc) when converting to/from CamelCase and snake_case.
package snaker

//go:generate ./gen.sh --update

import (
	"fmt"
	"strings"
	"unicode"
)

// CamelToSnake converts s to snake_case.
func CamelToSnake(s string) string {
	if s == "" {
		return ""
	}
	rs := []rune(s)
	var r strings.Builder
	var lastWasUpper, lastWasLetter, lastWasIsm, isUpper, isLetter bool
	for i := 0; i < len(rs); {
		isUpper, isLetter = unicode.IsUpper(rs[i]), unicode.IsLetter(rs[i])

		// append _ when last was not upper and not letter
		if (lastWasLetter && isUpper) || (lastWasIsm && isLetter) {
			r.WriteString("_")
		}

		// determine next to append to r
		var next string
		if ism := peekInitialism(rs[i:]); ism != "" && (!lastWasUpper || lastWasIsm) {
			next = ism
		} else {
			next = string(rs[i])
		}

		// save for next iteration
		lastWasIsm = len(next) > 1
		lastWasUpper, lastWasLetter = isUpper, isLetter
		r.WriteString(next)
		i += len(next)
	}
	return strings.ToLower(r.String())
}

// CamelToSnakeIdentifier converts s to its snake_case identifier.
func CamelToSnakeIdentifier(s string) string {
	return toIdentifier(CamelToSnake(s))
}

// SnakeToCamel converts s to CamelCase.
func SnakeToCamel(s string) string {
	var r strings.Builder
	for _, w := range strings.Split(s, "_") {
		if w == "" {
			continue
		}
		u := strings.ToUpper(w)
		if ok := commonInitialisms[u]; ok {
			r.WriteString(u)
		} else {
			r.WriteString(strings.ToUpper(w[:1]) + strings.ToLower(w[1:]))
		}
	}
	return r.String()
}

// SnakeToCamelIdentifier converts s to its CamelCase identifier (first
// letter is capitalized).
func SnakeToCamelIdentifier(s string) string {
	return SnakeToCamel(toIdentifier(s))
}

// ForceCamelIdentifier forces CamelCase specific to Go.
func ForceCamelIdentifier(s string) string {
	if s == "" {
		return ""
	}
	return SnakeToCamelIdentifier(CamelToSnake(s))
}

// ForceLowerCamelIdentifier forces the first portion of an identifier to be
// lower case.
func ForceLowerCamelIdentifier(s string) string {
	if s == "" {
		return ""
	}
	s = CamelToSnake(s)
	first := strings.SplitN(s, "_", -1)[0]
	s = SnakeToCamelIdentifier(s)
	return strings.ToLower(first) + s[len(first):]
}

// AddInitialisms adds initialisms to the recognized initialisms.
func AddInitialisms(initialisms ...string) error {
	for _, s := range initialisms {
		if len(s) < minInitialismLen || len(s) > maxInitialismLen {
			return fmt.Errorf("%s does not have length between %d and %d", s, minInitialismLen, maxInitialismLen)
		}
		commonInitialisms[s] = true
	}
	return nil
}

// IsInitialism indicates whether or not an initialism is registered as an
// identified initialism.
func IsInitialism(initialism string) bool {
	return commonInitialisms[strings.ToUpper(initialism)]
}
