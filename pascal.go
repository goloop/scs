package scs

import (
	"fmt"
	"regexp"
	"strings"
)

var (
	pascalHead   = regexp.MustCompile("(.)([A-Z][a-z]+)")
	pascalBody   = regexp.MustCompile("([a-z0-9])([A-Z])")
	isPascalCase = regexp.MustCompile(`^[A-Z]+((\d)|([A-Za-z0-9]+))*([A-Z])?$`)
)

// StrIsPascal returns true if string is PascalCase.
func StrIsPascal(s string) bool {
	return isPascalCase.Match([]byte(s))
}

// StrToPascal converts a string to PascalCase.
func StrToPascal(s string) string {
	return toUnited(s, false)
}

// ToPascal converts a string to PascalCase.
// Unlike the StrToPascal function, if the source string already
// has a certain format, it will be correctly converted to PascalCase.
func ToPascal(s string) string {
	switch {
	case StrIsCamel(s):
		r, _ := CamelToPascal(s)
		return r
	case StrIsKebab(s):
		r, _ := KebabToPascal(s)
		return r
	case StrIsPascal(s):
		// r, _ := PascalToSnake(s)
		// return r
		return s
	case StrIsSnake(s):
		r, _ := SnakeToPascal(s)
		return r
	}

	return StrToPascal(s)
}

// PascalToKebab converts a PascalCase-style string to kebab-case.
// The conversion will be invalid if the input string is not PascalCase style.
func PascalToKebab(pascal string) (string, error) {
	if !StrIsPascal(pascal) {
		return "", fmt.Errorf("value %s isn't PascalCase style", pascal)
	}

	kebab := pascalHead.ReplaceAllString(pascal, "${1}-${2}")
	kebab = pascalBody.ReplaceAllString(kebab, "${1}-${2}")
	return strings.ToLower(kebab), nil
}

// PascalToCamel converts a PascalCase-style string to camelCase.
// The conversion will be invalid if the input string is not PascalCase style.
func PascalToCamel(pascal string) (string, error) {
	if !StrIsPascal(pascal) {
		return "", fmt.Errorf("value %s isn't PascalCase style", pascal)
	}

	camel := pascalHead.ReplaceAllString(pascal, "${1} ${2}")
	camel = pascalBody.ReplaceAllString(camel, "${1} ${2}")
	return StrToCamel(camel), nil
}

// PascalToSnake converts a PascalCase-style string to snake_case.
// The conversion will be invalid if the input string is not PascalCase style.
func PascalToSnake(pascal string) (string, error) {
	if !StrIsPascal(pascal) {
		return "", fmt.Errorf("value %s isn't PascalCase style", pascal)
	}

	snake := pascalHead.ReplaceAllString(pascal, "${1}_${2}")
	snake = pascalBody.ReplaceAllString(snake, "${1}_${2}")
	return strings.ToLower(snake), nil
}
