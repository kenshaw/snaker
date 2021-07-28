// Package snaker provides methods to convert CamelCase to and from snake_case.
//
// snaker takes into takes into consideration common initialisms (ie, ID, HTTP,
// ACL, etc) when converting to/from CamelCase and snake_case.
package snaker

import (
	"fmt"
	"regexp"
	"strings"
	"unicode"
)

// CamelToSnake converts name from camel case ("AnIdentifier") to snake case
// ("an_identifier").
func CamelToSnake(name string) string {
	if name == "" {
		return ""
	}
	s, r := "", []rune(name)
	var lastWasUpper, lastWasLetter, lastWasIsm, isUpper, isLetter bool
	for i := 0; i < len(r); {
		isUpper, isLetter = unicode.IsUpper(r[i]), unicode.IsLetter(r[i])
		// append _ when last was not upper and not letter
		if (lastWasLetter && isUpper) || (lastWasIsm && isLetter) {
			s += "_"
		}
		// determine next to append to r
		var next string
		if ism := peekInitialism(r[i:]); ism != "" && (!lastWasUpper || lastWasIsm) {
			next = ism
		} else {
			next = string(r[i])
		}
		// save for next iteration
		lastWasIsm = len(next) > 1
		lastWasUpper, lastWasLetter = isUpper, isLetter
		s += next
		i += len(next)
	}
	return strings.ToLower(s)
}

// CamelToSnakeIdentifier converts name from camel case to a snake case
// identifier.
func CamelToSnakeIdentifier(name string) string {
	return toIdentifier(CamelToSnake(name))
}

// SnakeToCamel converts name to CamelCase.
func SnakeToCamel(name string) string {
	var s string
	for _, word := range strings.Split(name, "_") {
		if word == "" {
			continue
		}
		u := strings.ToUpper(word)
		if ok := commonInitialisms[u]; ok {
			s += u
		} else {
			s += strings.ToUpper(word[:1]) + strings.ToLower(word[1:])
		}
	}
	return s
}

// SnakeToCamelIdentifier converts name to its CamelCase identifier (first
// letter is capitalized).
func SnakeToCamelIdentifier(name string) string {
	return SnakeToCamel(toIdentifier(name))
}

// ForceCamelIdentifier forces name to its CamelCase specific to Go
// ("AnIdentifier").
func ForceCamelIdentifier(name string) string {
	if name == "" {
		return ""
	}
	return SnakeToCamelIdentifier(CamelToSnake(name))
}

// ForceLowerCamelIdentifier forces the first portion of an identifier to be
// lower case ("anIdentifier").
func ForceLowerCamelIdentifier(name string) string {
	if name == "" {
		return ""
	}
	name = CamelToSnake(name)
	first := strings.SplitN(name, "_", -1)[0]
	name = SnakeToCamelIdentifier(name)
	return strings.ToLower(first) + name[len(first):]
}

// AddInitialisms adds initialisms to the recognized initialisms.
func AddInitialisms(initialisms ...string) error {
	for _, s := range initialisms {
		if len(s) < minInitialismLen {
			return fmt.Errorf("%q must have length of at least %d", s, minInitialismLen)
		}
		if len(s) > maxInitialismLen {
			maxInitialismLen = len(s)
		}
		commonInitialisms[strings.ToUpper(s)] = true
	}
	return nil
}

// IsInitialism indicates whether or not an initialism is registered as an
// identified initialism.
func IsInitialism(initialism string) bool {
	return commonInitialisms[strings.ToUpper(initialism)]
}

var (
	// minInitialismLen is the min length of any of the commonInitialisms.
	minInitialismLen = 2
	// maxInitialismLen is the default max length of any of the commonInitialisms.
	// if  an initialism is added with a greater length; maxInitialismLen is
	// increased.
	maxInitialismLen = 5
)

// min returns the minimum of a, b.
func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}

// peekInitialism returns the next longest possible initialism in r.
func peekInitialism(r []rune) string {
	// do no work
	if len(r) < minInitialismLen {
		return ""
	}
	// grab at most next maxInitialismLen uppercase characters
	l := min(len(r), maxInitialismLen)
	var z []rune
	for i := 0; i < l; i++ {
		if !unicode.IsUpper(r[i]) {
			break
		}
		z = append(z, r[i])
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

// isIdentifierChar determines if ch is a valid character for a Go identifier.
//
// see: go/src/go/scanner/scanner.go
func isIdentifierChar(ch rune) bool {
	return 'a' <= ch && ch <= 'z' || 'A' <= ch && ch <= 'Z' || ch == '_' || ch >= 0x80 && unicode.IsLetter(ch) ||
		'0' <= ch && ch <= '9' || ch >= 0x80 && unicode.IsDigit(ch)
}

// replaceBadChars strips characters and character sequences that are invalid
// characters for Go identifiers.
func replaceBadChars(s string) string {
	// strip bad characters
	var r []rune
	for _, c := range s {
		if isIdentifierChar(c) {
			r = append(r, c)
		} else {
			r = append(r, '_')
		}
	}
	return string(r)
}

// underscoreRE matches underscores.
var underscoreRE = regexp.MustCompile(`_+`)

// leadingRE matches leading numbers.
var leadingRE = regexp.MustCompile(`^[0-9_]+`)

// toIdentifier cleans up a string so that it is usable as an identifier.
func toIdentifier(s string) string {
	// replace bad chars with _
	s = replaceBadChars(strings.TrimSpace(s))
	// fix 2 or more __ and remove leading numbers/underscores
	s = underscoreRE.ReplaceAllString(s, "_")
	s = leadingRE.ReplaceAllString(s, "_")
	// remove leading/trailing underscores
	s = strings.TrimLeft(s, "_")
	s = strings.TrimRight(s, "_")
	return s
}

// commonInitialisms is the set of commonInitialisms.
//
// Originally built from the list in golang.org/x/lint @ 738671d, however
// golang.org/x/lint has since been deprecated.
var commonInitialisms = map[string]bool{
	"ACL":   true,
	"API":   true,
	"ASCII": true,
	"CPU":   true,
	"CSS":   true,
	"DNS":   true,
	"EOF":   true,
	"GUID":  true,
	"HTML":  true,
	"HTTPS": true,
	"HTTP":  true,
	"ID":    true,
	"IP":    true,
	"JSON":  true,
	"LHS":   true,
	"QPS":   true,
	"RAM":   true,
	"RHS":   true,
	"RPC":   true,
	"SLA":   true,
	"SMTP":  true,
	"SQL":   true,
	"SSH":   true,
	"TCP":   true,
	"TLS":   true,
	"TTL":   true,
	"UDP":   true,
	"UID":   true,
	"UI":    true,
	"URI":   true,
	"URL":   true,
	"UTC":   true,
	"UTF8":  true,
	"UUID":  true,
	"VM":    true,
	"XML":   true,
	"XMPP":  true,
	"XSRF":  true,
	"XSS":   true,
	"YAML":  true,
}
