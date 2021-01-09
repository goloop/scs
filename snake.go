package scs

import (
	"regexp"
	"strings"
)

var snakeBody = regexp.MustCompile("(^[A-Za-z])|_([A-Za-z])")

// ToSnake converts a string to snake_case.
func ToSnake(s string) string {
	return toSeparate(s, "_")
}

// SnakeToCamel converts a snake_case-style string to camelCase.
// The conversion will be invalid if the input string is not snake_case style.
func SnakeToCamel(snake string) string {
	return ToCamel(
		snakeBody.ReplaceAllStringFunc(
			snake,
			func(s string) string { return strings.Replace(s, "_", " ", -1) },
		),
	)
}

// SnakeToKebab converts a snake_case-style string to kebab-case.
// The conversion will be invalid if the input string is not snake_case style.
func SnakeToKebab(snake string) string {
	return strings.ReplaceAll(snake, "_", "-")
}

// SnakeToPascal converts a snake_case-style string to PascalCase.
// The conversion will be invalid if the input string is not snake_case style.
func SnakeToPascal(snake string) string {
	return ToPascal(
		snakeBody.ReplaceAllStringFunc(
			snake,
			func(s string) string { return strings.Replace(s, "_", " ", -1) },
		),
	)
}
