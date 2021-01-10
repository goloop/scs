package scs

import (
	"fmt"
	"regexp"
	"strings"
)

var (
	snakeBody   = regexp.MustCompile("(^[A-Za-z])|_([A-Za-z])")
	isSnakeCase = regexp.MustCompile("(^[a-z0-9]+_[a-z0-9_]+$)|(^[a-z0-9]+$)")
)

// StrIsSnake returns true if string is snake_case.
func StrIsSnake(s string) bool {
	return isSnakeCase.Match([]byte(s))
}

// StrToSnake converts a string to snake_case.
func StrToSnake(s string) (string, error) {
	return toSeparate(s, "_")
}

// SnakeToCamel converts a snake_case-style string to camelCase.
// The conversion will be invalid if the input string is not snake_case style.
func SnakeToCamel(snake string) (string, error) {
	if !StrIsSnake(snake) {
		return "", fmt.Errorf("value %s isn't snake_case style", snake)
	}

	return StrToCamel(
		snakeBody.ReplaceAllStringFunc(
			snake,
			func(s string) string { return strings.Replace(s, "_", " ", -1) },
		),
	)
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

	return StrToPascal(
		snakeBody.ReplaceAllStringFunc(
			snake,
			func(s string) string { return strings.Replace(s, "_", " ", -1) },
		),
	)
}
