package benchmarking

import (
	"strings"
	"unicode"
	"unicode/utf8"
)

// The "Bad" $O(n^2)$ Catastrophe
func RemoveWhitespaceThrasher(s string) string {
	var result string
	for _, r := range s {
		if !unicode.IsSpace(r) {
			result += string(r)
		}
	}
	return result
}

// The "Good" Idiomatic Way
func RemoveWhitespaceBuilder(s string) string {
	var b strings.Builder
	b.Grow(len(s))
	for _, r := range s {
		if !unicode.IsSpace(r) {
			b.WriteRune(r)
		}
	}
	return b.String()
}

// The "Optimized" In-Place Way
func RemoveWhitespaceInPlace(s string) string {
	b := []byte(s)
	n := 0
	for i := 0; i < len(b); {
		r, size := utf8.DecodeRune(b[i:])
		if !unicode.IsSpace(r) {
			if n != i {
				utf8.EncodeRune(b[n:], r)
			}
			n += size
		}
		i += size
	}
	return string(b[:n])
}
