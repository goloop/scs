package scs

import (
	"fmt"
	"regexp"
	"strings"
)

var (
	camelHead    = regexp.MustCompile("(.)([A-Z][a-z]+)")
	camelBody    = regexp.MustCompile("([a-z0-9])([A-Z])")
	camelNumbers = regexp.MustCompile("([0-9]+)")
	isCamelCase  = regexp.MustCompile(`^[a-z]+((\d)|([A-Za-z0-9]+))*([A-Z])?$`)
)

// StrIsCamel returns true if string is camelCase.
func StrIsCamel(s string) bool {
	return isCamelCase.Match([]byte(s))
}

// StrToCamel converts a string to camelCase.
func StrToCamel(s string) string {
	return toUnited(s, true)
}

// ToCamel converts a string to camelCase.
// Unlike the StrToCamel function, if the source string already
// has a certain format, it will be correctly converted to camelCase.
func ToCamel(s string) string {
	switch {
	case StrIsCamel(s):
		return s
	case StrIsKebab(s):
		r, _ := KebabToCamel(s)
		return r
	case StrIsPascal(s):
		r, _ := PascalToCamel(s)
		return r
	case StrIsSnake(s):
		r, _ := SnakeToCamel(s)
		return r
	}

	return StrToCamel(s)
}

// CamelToKebab converts a camelCase-style string to kebab-case.
// The conversion will be invalid if the input string is not camelCase style.
func CamelToKebab(camel string) (string, error) {
	if !StrIsCamel(camel) {
		return "", fmt.Errorf("value %s isn't camelCase style", camel)
	}

	kebab := camelHead.ReplaceAllString(camel, "${1}-${2}")
	kebab = camelBody.ReplaceAllString(kebab, "${1}-${2}")
	return strings.ToLower(kebab), nil
}

// CamelToPascal converts a camelCase-style string to PascalCase.
// The conversion will be invalid if the input string is not camelCase style.
func CamelToPascal(camel string) (string, error) {
	if !StrIsCamel(camel) {
		return "", fmt.Errorf("value %s isn't camelCase style", camel)
	}

	pascal := camelHead.ReplaceAllString(camel, "${1} ${2}")
	pascal = camelBody.ReplaceAllString(pascal, "${1} ${2}")
	pascal = camelNumbers.ReplaceAllString(pascal, " ${1} ")
	return StrToPascal(pascal), nil
}

// CamelToSnake converts a camelCase-style string to snake_case.
// The conversion will be invalid if the input string is not camelCase style.
func CamelToSnake(camel string) (string, error) {
	if !StrIsCamel(camel) {
		return "", fmt.Errorf("value %s isn't camelCase style", camel)
	}

	snake := camelHead.ReplaceAllString(camel, "${1}_${2}")
	snake = camelBody.ReplaceAllString(snake, "${1}_${2}")
	return strings.ToLower(snake), nil
}
