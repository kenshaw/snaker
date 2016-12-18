package snaker

import (
	"regexp"
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
	r := []rune{}
	for _, ch := range s {
		if isIdentifierChar(ch) {
			r = append(r, ch)
		} else {
			r = append(r, '_')
		}
	}

	return string(r)
}

// underscoreRE matches underscores.
var underscoreRE = regexp.MustCompile(`_+`)

// numberRE matches leading numbers.
var numberRE = regexp.MustCompile(`^[0-9]+`)
