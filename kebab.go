package scs

import (
	"regexp"
	"strings"
)

var kebabBody = regexp.MustCompile("(^[A-Za-z])|-([A-Za-z])")

// ToKebab converts a string to kebab-case.
func ToKebab(s string) string {
	return toSeparate(s, "-")
}

// KebabToCamel converts a kebab-case-style string to camelCase.
// The conversion will be invalid if the input string is not kebab-case style.
func KebabToCamel(kebab string) string {
	return ToCamel(
		kebabBody.ReplaceAllStringFunc(
			kebab,
			func(s string) string { return strings.Replace(s, "-", " ", -1) },
		),
	)
}

// KebabToSnake converts a kebab-case-style string to snake_case.
// The conversion will be invalid if the input string is not kebab-case style.
func KebabToSnake(kebab string) string {
	return strings.ReplaceAll(kebab, "-", "_")
}

// KebabToPascal converts a kebab-case-style string to PascalCase.
// The conversion will be invalid if the input string is not kebab-case style.
func KebabToPascal(kebab string) string {
	return ToPascal(
		kebabBody.ReplaceAllStringFunc(
			kebab,
			func(s string) string { return strings.Replace(s, "-", " ", -1) },
		),
	)
}
