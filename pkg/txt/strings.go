package txt

import (
	"regexp"
	"strings"
	"unicode"
)

var ContainsNumberRegexp = regexp.MustCompile("\\d+")

// ContainsNumber returns true if string contains a number.
func ContainsNumber(s string) bool {
	return ContainsNumberRegexp.MatchString(s)
}

// isSeparator reports whether the rune could mark a word boundary.
func isSeparator(r rune) bool {
	// ASCII alphanumerics and underscore are not separators
	if r <= 0x7F {
		switch {
		case '0' <= r && r <= '9':
			return false
		case 'a' <= r && r <= 'z':
			return false
		case 'A' <= r && r <= 'Z':
			return false
		case r == '_', r == '\'':
			return false
		}
		return true
	}
	// Letters and digits are not separators
	if unicode.IsLetter(r) || unicode.IsDigit(r) {
		return false
	}
	// Otherwise, all we can do for now is treat spaces as separators.
	return unicode.IsSpace(r)
}

// UcFirst returns the string with the first character converted to uppercase.
func UcFirst(str string) string {
	for i, v := range str {
		return string(unicode.ToUpper(v)) + str[i+1:]
	}
	return ""
}

// Title returns the string with the first characters of each word converted to uppercase.
func Title(s string) string {
	prev := ' '
	return strings.Map(
		func(r rune) rune {
			if isSeparator(prev) {
				prev = r
				return unicode.ToTitle(r)
			}
			prev = r
			return r
		},
		s)
}

// Bool casts a string to bool.
func Bool(s string) bool {
	s = strings.TrimSpace(s)

	if s == "" || s == "0" || s == "false" || s == "no" {
		return false
	}

	return true
}
