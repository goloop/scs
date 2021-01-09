package scs

import (
	"regexp"
	"strings"
)

var (
	pascalHead = regexp.MustCompile("(.)([A-Z][a-z]+)")
	pascalBody = regexp.MustCompile("([a-z0-9])([A-Z])")
)

// ToPascal converts a string to PascalCase.
func ToPascal(s string) string {
	return toUnited(s, false)
}

// PascalToKebab converts a PascalCase-style string to kebab-case.
// The conversion will be invalid if the input string is not PascalCase style.
func PascalToKebab(pascal string) string {
	kebab := pascalHead.ReplaceAllString(pascal, "${1}-${2}")
	kebab = pascalBody.ReplaceAllString(kebab, "${1}-${2}")
	return strings.ToLower(kebab)
}

// PascalToCamel converts a PascalCase-style string to camelCase.
// The conversion will be invalid if the input string is not PascalCase style.
func PascalToCamel(pascal string) string {
	camel := pascalHead.ReplaceAllString(pascal, "${1} ${2}")
	camel = pascalBody.ReplaceAllString(camel, "${1} ${2}")
	return ToCamel(camel)
}

// PascalToSnake converts a PascalCase-style string to snake_case.
// The conversion will be invalid if the input string is not PascalCase style.
func PascalToSnake(pascal string) string {
	snake := pascalHead.ReplaceAllString(pascal, "${1}_${2}")
	snake = pascalBody.ReplaceAllString(snake, "${1}_${2}")
	return strings.ToLower(snake)
}
