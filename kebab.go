package scs

import (
	"fmt"
	"regexp"
	"strings"
)

var (
	kebabBody   = regexp.MustCompile("(^[A-Za-z])|-([A-Za-z])")
	isKebabCase = regexp.MustCompile("(^[a-z0-9]+-[a-z0-9-]+$)|(^[a-z0-9]+$)")
)

// StrIsKebab returns true if string is kebab-case.
func StrIsKebab(s string) bool {
	return isKebabCase.Match([]byte(s))
}

// StrToKebab converts a string to kebab-case.
func StrToKebab(s string) string {
	return toSeparate(s, "-")
}

// ToKebab converts a string to kebab-case.
// Unlike the StrToKebab function, if the source string already
// has a certain format, it will be correctly converted to kebab-case.
func ToKebab(s string) string {
	switch {
	case StrIsCamel(s):
		r, _ := CamelToKebab(s)
		return r
	case StrIsKebab(s):
		return s
	case StrIsPascal(s):
		r, _ := PascalToKebab(s)
		return r
	case StrIsSnake(s):
		r, _ := SnakeToKebab(s)
		return r
	}

	return StrToKebab(s)
}

// KebabToCamel converts a kebab-case-style string to camelCase.
// The conversion will be invalid if the input string is not kebab-case style.
func KebabToCamel(kebab string) (string, error) {
	if !StrIsKebab(kebab) {
		return "", fmt.Errorf("value %s isn't kebab-case style", kebab)
	}

	result := StrToCamel(
		kebabBody.ReplaceAllStringFunc(
			kebab,
			func(s string) string { return strings.Replace(s, "-", " ", -1) },
		),
	)

	return result, nil
}

// KebabToSnake converts a kebab-case-style string to snake_case.
// The conversion will be invalid if the input string is not kebab-case style.
func KebabToSnake(kebab string) (string, error) {
	if !StrIsKebab(kebab) {
		return "", fmt.Errorf("value %s isn't kebab-case style", kebab)
	}

	return strings.ReplaceAll(kebab, "-", "_"), nil
}

// KebabToPascal converts a kebab-case-style string to PascalCase.
// The conversion will be invalid if the input string is not kebab-case style.
func KebabToPascal(kebab string) (string, error) {
	if !StrIsKebab(kebab) {
		return "", fmt.Errorf("value %s isn't kebab-case style", kebab)
	}

	result := StrToPascal(
		kebabBody.ReplaceAllStringFunc(
			kebab,
			func(s string) string { return strings.Replace(s, "-", " ", -1) },
		),
	)

	return result, nil
}
