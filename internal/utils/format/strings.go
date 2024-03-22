package format

import "unicode/utf8"

// Truncate truncates a given string to the specified length.
func Truncate(s string, length int) string {
	if length <= 0 {
		return ""
	}

	if utf8.RuneCountInString(s) <= length {
		return s
	}

	runes := []rune(s)

	return string(runes[:length])
}
