package scs

import (
	"regexp"
	"strings"
)

var chunkLayout = regexp.MustCompile(`[^0-9a-zA-Z ]+`)

// The getChunks clears the string of special characters and splits the
// string by whitespace and returns a list of words ignoring empty elements.
func getChunks(s string) []string {
	return strings.FieldsFunc(
		chunkLayout.ReplaceAllString(strings.ToLower(s), ""),
		func(c rune) bool { return c == ' ' },
	)
}

// The toUnited converts a string to a format similar to camel or PascaCase.
func toUnited(s string, firstWordIsLower bool) string {
	var result string

	for i, chunk := range getChunks(s) {
		if i == 0 && firstWordIsLower {
			result = chunk
			continue
		}

		if v, ok := acronyms[chunk]; ok {
			result += v
			continue
		}

		result += strings.Title(chunk)
	}

	return result
}

// The toSeparate converts a string to a format similar to snake or kebab-case.
func toSeparate(s, delimiter string) string {
	var result string

	for i, chunk := range getChunks(s) {
		if i == 0 {
			result = chunk
			continue
		}

		result += delimiter + chunk
	}

	return result
}
