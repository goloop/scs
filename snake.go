package scs

import (
	"fmt"
	"regexp"
	"strings"
)

var (
	snakeBody   = regexp.MustCompile("(^[A-Za-z0-9])|_([A-Za-z0-9])")
	isSnakeCase = regexp.MustCompile("(^[a-z0-9]+_[a-z0-9_]+$)|(^[a-z0-9]+$)")
)

// StrIsSnake returns true if string is snake_case.
func StrIsSnake(s string) bool {
	return isSnakeCase.Match([]byte(s))
}

// StrToSnake converts a string to snake_case.
func StrToSnake(s string) string {
	return toSeparate(s, "_")
}

// ToSnake converts a string to snake_case.
// Unlike the StrToSnake function, if the source string already
// has a certain format, it will be correctly converted to snake_case.
func ToSnake(s string) string {
	switch {
	case StrIsCamel(s):
		r, _ := CamelToSnake(s)
		return r
	case StrIsKebab(s):
		r, _ := KebabToSnake(s)
		return r
	case StrIsPascal(s):
		r, _ := PascalToSnake(s)
		return r
	case StrIsSnake(s):
		return s
	}

	return StrToSnake(s)
}

// SnakeToCamel converts a snake_case-style string to camelCase.
// The conversion will be invalid if the input string is not snake_case style.
func SnakeToCamel(snake string) (string, error) {
	if !StrIsSnake(snake) {
		return "", fmt.Errorf("value %s isn't snake_case style", snake)
	}

	result := StrToCamel(
		snakeBody.ReplaceAllStringFunc(
			snake,
			func(s string) string { return strings.Replace(s, "_", " ", -1) },
		),
	)

	return result, nil
}

// SnakeToKebab converts a snake_case-style string to kebab-case.
// The conversion will be invalid if the input string is not snake_case style.
func SnakeToKebab(snake string) (string, error) {
	if !StrIsSnake(snake) {
		return "", fmt.Errorf("value %s isn't snake_case style", snake)
	}

	return strings.ReplaceAll(snake, "_", "-"), nil
}

// SnakeToPascal converts a snake_case-style string to PascalCase.
// The conversion will be invalid if the input string is not snake_case style.
func SnakeToPascal(snake string) (string, error) {
	if !StrIsSnake(snake) {
		return "", fmt.Errorf("value %s isn't snake_case style", snake)
	}

	result := StrToPascal(
		snakeBody.ReplaceAllStringFunc(
			snake,
			func(s string) string { return strings.Replace(s, "_", " ", -1) },
		),
	)

	return result, nil
}
