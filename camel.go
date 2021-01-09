package scs

import (
	"regexp"
	"strings"
)

var (
	camelHead    = regexp.MustCompile("(.)([A-Z][a-z]+)")
	camelBody    = regexp.MustCompile("([a-z0-9])([A-Z])")
	camelNumbers = regexp.MustCompile("([0-9]+)")
)

// ToCamel converts a string to camelCase.
func ToCamel(s string) string {
	return toUnited(s, true)
}

// CamelToKebab converts a camelCase-style string to kebab-case.
// The conversion will be invalid if the input string is not camelCase style.
func CamelToKebab(camel string) string {
	kebab := camelHead.ReplaceAllString(camel, "${1}-${2}")
	kebab = camelBody.ReplaceAllString(kebab, "${1}-${2}")
	return strings.ToLower(kebab)
}

// CamelToPascal converts a camelCase-style string to PascalCase.
// The conversion will be invalid if the input string is not camelCase style.
func CamelToPascal(camel string) string {
	pascal := camelHead.ReplaceAllString(camel, "${1} ${2}")
	pascal = camelBody.ReplaceAllString(pascal, "${1} ${2}")
	pascal = camelNumbers.ReplaceAllString(pascal, " ${1} ")
	return ToPascal(pascal)
}

// CamelToSnake converts a camelCase-style string to snake_case.
// The conversion will be invalid if the input string is not camelCase style.
func CamelToSnake(camel string) string {
	snake := camelHead.ReplaceAllString(camel, "${1}_${2}")
	snake = camelBody.ReplaceAllString(snake, "${1}_${2}")
	return strings.ToLower(snake)
}
