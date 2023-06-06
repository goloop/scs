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

// StrIsSnake returns true if the string is in snake_case.
//
// Snake case represents words separated by underscores and does
// not have any capital letters.
//
// Example usage:
//
//	scs.StrIsSnake("hello_world") // returns true
//	scs.StrIsSnake("HelloWorld")  // returns false
//	scs.StrIsSnake("hello-world") // returns false
func StrIsSnake(s string) bool {
	return isSnakeCase.Match([]byte(s))
}

// StrToSnake converts a string to snake_case.
//
// This function first transforms the input string by replacing all
// non-alphanumeric characters with spaces.
//
// It then splits the string into chunks at the spaces, converts
// each chunk to lower case, and finally joins the chunks back together
// with underscores.
//
// Note that this function will convert upper case letters to lower case,
// so the output string will be all lower case even if the input string
// contains upper case letters.
//
// Example usage:
//
//	scs.StrToSnake("Hello World") // returns "hello_world"
//	scs.StrToSnake("hello world") // returns "hello_world"
//	scs.StrToSnake("hello-world") // returns "hello_world"
func StrToSnake(s string) string {
	return toSeparate(s, "_")
}

// ToSnake converts a string to snake_case.
//
// Unlike the StrToSnake function, this function attempts to correctly
// handle strings that are already in a certain format (such as CamelCase,
// KebabCase, or PascalCase) and convert them into snake_case.
//
// The function first checks if the string is in CamelCase, KebabCase,
// PascalCase, or  SnakeCase format. If it is, it uses the corresponding
// conversion function to convert the string into snake_case.
//
// If the string is not in any recognized format, it defaults to using the
// StrToSnake function to attempt a conversion.
//
// Example usage:
//
//	scs.ToSnake("helloWorld")   // returns "hello_world"
//	scs.ToSnake("HelloWorld")   // returns "hello_world"
//	scs.ToSnake("hello-world")  // returns "hello_world"
//	scs.ToSnake("hello_world")  // returns "hello_world"
//	scs.ToSnake("Hello World")  // returns "hello_world"
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
//
// This function checks if the input string is in snake_case.
// If not, it returns an error.
//
// If the input string is in snake_case, it replaces all underscores
// with spaces and converts the resulting string to camelCase.
//
// Note that the first word in the output string will be in lower case,
// and the first letter of each subsequent word will be in upper case.
// All other letters will be lower case.
//
// This conversion could fail if the input string is not in snake_case style.
// In that case, an error will be returned along with an empty string.
//
// Example usage:
//
//	scs.SnakeToCamel("hello_world") // returns "helloWorld", nil
//	scs.SnakeToCamel("HelloWorld")  // returns "", error
//	scs.SnakeToCamel("hello-world") // returns "", error
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
//
// This function checks if the input string is in snake_case. If it's not,
// it returns an error.
//
// If the input string is in snake_case, it replaces all underscores with
// hyphens, effectively converting the string from snake_case to kebab-case.
//
// Note that this conversion could fail if the input string is not in
// snake_case style. In that case, an error will be returned along with
// an empty string.
//
// Example usage:
//
//	scs.SnakeToKebab("hello_world") // returns "hello-world", nil
//	scs.SnakeToKebab("HelloWorld")  // returns "", error
//	scs.SnakeToKebab("helloWorld")  // returns "", error
func SnakeToKebab(snake string) (string, error) {
	if !StrIsSnake(snake) {
		return "", fmt.Errorf("value %s isn't snake_case style", snake)
	}

	return strings.ReplaceAll(snake, "_", "-"), nil
}

// SnakeToPascal converts a snake_case-style string to PascalCase.
//
// This function checks if the input string is in snake_case. If it's not,
// it returns an error.
//
// If the input string is in snake_case, it replaces all underscores with
// spaces and converts the resulting string to PascalCase. The PascalCase
// format starts with an uppercase letter, and includes an uppercase letter
// at the beginning of each new word.
//
// Note that this conversion could fail if the input string is not in
// snake_case style. In that case, an error will be returned along with
// an empty string.
//
// Example usage:
//
//	scs.SnakeToPascal("hello_world") // returns "HelloWorld", nil
//	scs.SnakeToPascal("HelloWorld")  // returns "", error
//	scs.SnakeToPascal("hello-world") // returns "", error
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
