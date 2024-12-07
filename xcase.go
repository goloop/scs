package scs

import (
	"strings"
	"unicode"
)

// The getChunks clears the string of special characters and splits the
// string by whitespace and returns a list of words ignoring empty elements.
func getChunks(s string) []string {
	chunks := make([]string, 0, strings.Count(s, " ")+1)
	var builder strings.Builder
	builder.Grow(len(s))

	for _, r := range strings.ToLower(s) {
		if unicode.IsLetter(r) || unicode.IsNumber(r) {
			builder.WriteRune(r)
		} else if builder.Len() > 0 {
			chunks = append(chunks, builder.String())
			builder.Reset()
		}
	}

	if builder.Len() > 0 {
		chunks = append(chunks, builder.String())
	}

	return chunks
}

// The toUnited converts a string to a format similar to camel or PascalCase.
func toUnited(s string, firstWordIsLower bool) string {
	chunks := getChunks(s)
	if len(chunks) == 0 {
		return ""
	}

	var builder strings.Builder
	builder.Grow(len(s))

	// Перше слово
	if firstWordIsLower {
		builder.WriteString(chunks[0])
	} else if v, ok := abbreviations[chunks[0]]; ok {
		builder.WriteString(v)
	} else {
		builder.WriteString(strings.Title(chunks[0]))
	}

	// Решта слів
	for _, chunk := range chunks[1:] {
		if v, ok := abbreviations[chunk]; ok {
			builder.WriteString(v)
		} else {
			builder.WriteString(strings.Title(chunk))
		}
	}

	return builder.String()
}

// The toSeparate converts a string to a format similar to snake or kebab-case.
func toSeparate(s, delimiter string) string {
	chunks := getChunks(s)
	if len(chunks) == 0 {
		return ""
	}

	var builder strings.Builder
	totalLen := len(s) + len(chunks) - 1
	builder.Grow(totalLen)

	builder.WriteString(chunks[0])
	for _, chunk := range chunks[1:] {
		builder.WriteString(delimiter)
		builder.WriteString(chunk)
	}

	return builder.String()
}
