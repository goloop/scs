package scs

import (
	"fmt"
	"regexp"
	"strings"
)

var (
	pascalPrep   = regexp.MustCompile("([0-9]+)")
	pascalHead   = regexp.MustCompile("(.)([A-Z][a-z]+)")
	pascalBody   = regexp.MustCompile("([a-z0-9])([A-Z])")
	isPascalCase = regexp.MustCompile(`^[A-Z]+((\d)|([A-Za-z0-9]+))*([A-Z])?$`)
)

// StrIsPascal returns true if the string is in PascalCase.
//
// PascalCase is a naming convention in which the first letter of each word
// (including the first word) is capitalized, and there are no underscores
// between words.
//
// Example usage:
//
//	scs.StrIsPascal("HelloWorld")    // returns true
//	scs.StrIsPascal("helloWorld")    // returns false
//	scs.StrIsPascal("hello_world")   // returns false
//	scs.StrIsPascal("Hello_World")   // returns false
func StrIsPascal(s string) bool {
	return isPascalCase.Match([]byte(s))
}

// StrToPascal converts a string to PascalCase.
//
// This function transforms the input string to PascalCase format.
// PascalCase is a naming convention in which the first letter of each word,
// including the first word, is capitalized, and there are no underscores
// or spaces between words.
//
// Note that this function will convert all characters to lowercase, except
// for the first letter of each word, which will be capitalized.
//
// Example usage:
//
//	scs.StrToPascal("hello_world") // returns "HelloWorld"
//	scs.StrToPascal("hello world") // returns "HelloWorld"
//	scs.StrToPascal("helloWorld")  // returns "Helloworld"
func StrToPascal(s string) string {
	return toUnited(s, false)
}

// ToPascal converts a string to PascalCase.
//
// This function converts the input string to PascalCase format. PascalCase
// is a naming convention in which the first letter of each word (including
// the first word) is capitalized, and there are no underscores, spaces, or
// hyphens between words.
//
// If the source string already has a certain format, such as camelCase,
// kebab-case, or snake_case, it will be correctly converted to PascalCase.
//
// Example usage:
//
//	scs.ToPascal("hello_world")   // returns "HelloWorld"
//	scs.ToPascal("hello-world")   // returns "HelloWorld"
//	scs.ToPascal("helloWorld")    // returns "HelloWorld"
//	scs.ToPascal("helloWorld123") // returns "HelloWorld123"
func ToPascal(s string) string {
	switch {
	case StrIsCamel(s):
		r, _ := CamelToPascal(s)
		return r
	case StrIsKebab(s):
		r, _ := KebabToPascal(s)
		return r
	case StrIsPascal(s):
		return s
	case StrIsSnake(s):
		r, _ := SnakeToPascal(s)
		return r
	}

	return StrToPascal(s)
}

// PascalToKebab converts a PascalCase-style string to kebab-case.
//
// This function checks if the input string is in PascalCase. If it's not,
// it returns an error.
//
// If the input string is in PascalCase, it performs the conversion by
// inserting hyphens before each capital letter (except the first one)
// and then converts the string to lowercase.
//
// Note that this conversion could fail if the input string is not in
// PascalCase style. In that case, an error will be returned along with
// an empty string.
//
// Example usage:
//
//	result, err := PascalToKebab("HelloWorld")
//	// result: "hello-world", err: nil
//
//	result, err := PascalToKebab("helloWorld")
//	// result: "", err: error (not PascalCase)
func PascalToKebab(pascal string) (string, error) {
	if !StrIsPascal(pascal) {
		return "", fmt.Errorf("value %s isn't PascalCase style", pascal)
	}

	kebab := pascalPrep.ReplaceAllString(pascal, "-${1}-")
	kebab = pascalHead.ReplaceAllString(kebab, "${1}-${2}")
	kebab = pascalBody.ReplaceAllString(kebab, "${1}-${2}")
	return strings.ToLower(strings.Trim(kebab, "-")), nil
}

// PascalToCamel converts a PascalCase-style string to camelCase.
//
// This function checks if the input string is in PascalCase. If it's not,
// it returns an error.
//
// If the input string is in PascalCase, it performs the conversion by
// inserting a space before each capital letter (except the first one),
// and then converts the string to camelCase.
//
// Note that this conversion could fail if the input string is not in
// PascalCase style. In that case, an error will be returned along with
// an empty string.
//
// Example usage:
//
//	result, err := PascalToCamel("HelloWorld")
//	// result: "helloWorld", err: nil
//
//	result, err := PascalToCamel("helloWorld")
//	// result: "", err: error (not PascalCase)
func PascalToCamel(pascal string) (string, error) {
	if !StrIsPascal(pascal) {
		return "", fmt.Errorf("value %s isn't PascalCase style", pascal)
	}

	camel := pascalHead.ReplaceAllString(pascal, "${1} ${2}")
	camel = pascalBody.ReplaceAllString(camel, "${1} ${2}")
	return StrToCamel(camel), nil
}

// PascalToSnake converts a PascalCase-style string to snake_case.
//
// This function checks if the input string is in PascalCase. If it's not,
// it returns an error.
//
// If the input string is in PascalCase, it performs the conversion by
// inserting underscores before each capital letter (except the first one),
// and then converts the string to snake_case.
//
// Note that this conversion could fail if the input string is not in
// PascalCase style. In that case, an error will be returned along with
// an empty string.
//
// Example usage:
//
//	result, err := PascalToSnake("HelloWorld")
//	// result: "hello_world", err: nil
//
//	result, err := PascalToSnake("helloWorld")
//	// result: "", err: error (not PascalCase)
func PascalToSnake(pascal string) (string, error) {
	if !StrIsPascal(pascal) {
		return "", fmt.Errorf("value %s isn't PascalCase style", pascal)
	}

	snake := pascalPrep.ReplaceAllString(pascal, "_${1}_")
	snake = pascalHead.ReplaceAllString(snake, "${1}_${2}")
	snake = pascalBody.ReplaceAllString(snake, "${1}_${2}")
	return strings.ToLower(strings.Trim(snake, "_")), nil
}
