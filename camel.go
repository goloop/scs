package scs

import (
	"fmt"
	"regexp"
	"strings"
)

var (
	camelPrep    = regexp.MustCompile("([0-9]+)")
	camelHead    = regexp.MustCompile("(.)([A-Z][a-z]+)")
	camelBody    = regexp.MustCompile("([a-z0-9])([A-Z])")
	camelNumbers = regexp.MustCompile("([0-9]+)")
	isCamelCase  = regexp.MustCompile(`^[a-z]+((\d)|([A-Za-z0-9]+))*([A-Z])?$`)
)

// StrIsCamel returns true if the string is in camelCase.
//
// CamelCase is a naming convention in which the first letter of the first
// word is lowercase, and the first letter of each subsequent word is
// capitalized. There are no spaces or underscores between words.
//
// Example usage:
//
//	scs.StrIsCamel("helloWorld")  // returns true
//	scs.StrIsCamel("HelloWorld")  // returns true
//	scs.StrIsCamel("hello_world") // returns false
//	scs.StrIsCamel("hello-World") // returns false
func StrIsCamel(s string) bool {
	return isCamelCase.Match([]byte(s))
}

// StrToCamel converts a string to camelCase.
//
// This function converts the input string to camelCase format. CamelCase is
// a naming convention in which the first letter of the first word is
// lowercase, and the first letter of each subsequent word is capitalized.
// There are no spaces or underscores between words.
//
// Note that this function will convert all characters to lowercase, except
// for the first letter of the first word, which will be lowercase.
//
// Example usage:
//
//	scs.StrToCamel("hello_world")  // returns "helloWorld"
//	scs.StrToCamel("hello-world")  // returns "helloWorld"
//	scs.StrToCamel("HelloWorld")   // returns "helloWorld"
func StrToCamel(s string) string {
	return toUnited(s, true)
}

// ToCamel converts a string to camelCase.
//
// This function converts the input string to camelCase format. CamelCase is
// a naming convention in which the first letter of the first word is
// lowercase, and the first letter of each subsequent word is capitalized.
// There are no spaces or underscores between words.
//
// Unlike the StrToCamel function, if the source string already has a certain
// format such as camelCase, kebab-case, PascalCase, or snake_case, it will
// be correctly converted to camelCase.
//
// Example usage:
//
//	scs.ToCamel("hello_world")  // returns "helloWorld"
//	scs.ToCamel("hello-world")  // returns "helloWorld"
//	scs.ToCamel("HelloWorld")   // returns "helloWorld"
//	scs.ToCamel("snake_case")   // returns "snakeCase"
//	scs.ToCamel("kebab-case")   // returns "kebabCase"
//	scs.ToCamel("PascalCase")   // returns "pascalCase"
//	scs.ToCamel("camelCase")    // returns "camelCase"
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
//
// This function checks if the input string is in camelCase. If it's not,
// it returns an error.
//
// If the input string is in camelCase, it performs the conversion by
// inserting hyphens before each capital letter (except the first one)
// and then converts the string to lowercase.
//
// Note that this conversion could fail if the input string is not in
// camelCase style. In that case, an error will be returned along with
// an empty string.
//
// Example usage:
//
//	result, err := CamelToKebab("helloWorld")
//	// result: "hello-world", err: nil
//
//	result, err := CamelToKebab("HelloWorld")
//	// result: "", err: error (not camelCase)
func CamelToKebab(camel string) (string, error) {
	if !StrIsCamel(camel) {
		return "", fmt.Errorf("value %s isn't camelCase style", camel)
	}

	kebab := camelPrep.ReplaceAllString(camel, "-${1}-")
	kebab = camelHead.ReplaceAllString(kebab, "${1}-${2}")
	kebab = camelBody.ReplaceAllString(kebab, "${1}-${2}")
	return strings.ToLower(strings.Trim(kebab, "-")), nil
}

// CamelToPascal converts a camelCase-style string to PascalCase.
//
// This function checks if the input string is in camelCase. If it's not,
// it returns an error.
//
// If the input string is in camelCase, it performs the conversion by
// inserting spaces before each capital letter (except the first one)
// and then converts the string to PascalCase.
//
// Note that this conversion could fail if the input string is not in
// camelCase style. In that case, an error will be returned along with
// an empty string.
//
// Example usage:
//
//	result, err := CamelToPascal("helloWorld")
//	// result: "HelloWorld", err: nil
//
//	result, err := CamelToPascal("HelloWorld")
//	// result: "HelloWorld", err: nil
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
//
// This function checks if the input string is in camelCase. If it's not,
// it returns an error.
//
// If the input string is in camelCase, it performs the conversion by
// inserting underscores before each capital letter (except the first one)
// and then converts the string to snake_case.
//
// Note that this conversion could fail if the input string is not in
// camelCase style. In that case, an error will be returned along with
// an empty string.
//
// Example usage:
//
//	result, err := CamelToSnake("helloWorld")
//	// result: "hello_world", err: nil
//
//	result, err := CamelToSnake("HelloWorld")
//	// result: "hello_world", err: nil
func CamelToSnake(camel string) (string, error) {
	if !StrIsCamel(camel) {
		return "", fmt.Errorf("value %s isn't camelCase style", camel)
	}

	snake := camelPrep.ReplaceAllString(camel, "_${1}_")
	snake = camelHead.ReplaceAllString(snake, "${1}_${2}")
	snake = camelBody.ReplaceAllString(snake, "${1}_${2}")
	return strings.ToLower(strings.Trim(snake, "_")), nil
}
