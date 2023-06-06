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

// StrIsKebab returns true if the string is in kebab-case.
//
// Kebab-case is a naming convention in which words are separated by hyphens,
// and all letters are lowercase.
//
// Example usage:
//
//	scs.StrIsKebab("hello-world")  // returns true
//	scs.StrIsKebab("HelloWorld")   // returns false
//	scs.StrIsKebab("hello_world")  // returns false
//	scs.StrIsKebab("Hello-World")  // returns false
func StrIsKebab(s string) bool {
	return isKebabCase.Match([]byte(s))
}

// StrToKebab converts a string to kebab-case.
//
// This function transforms the input string to kebab-case format. Kebab-case
// is a naming convention in which words are separated by hyphens, and all
// letters are lowercase.
//
// Example usage:
//
//	scs.StrToKebab("helloWorld")   // returns "hello-world"
//	scs.StrToKebab("HelloWorld")   // returns "hello-world"
//	scs.StrToKebab("hello_world")  // returns "hello-world"
//	scs.StrToKebab("Hello-World")  // returns "hello-world"
func StrToKebab(s string) string {
	return toSeparate(s, "-")
}

// ToKebab converts a string to kebab-case.
//
// This function converts the input string to kebab-case format. Kebab-case
// is a naming convention in which words are separated by hyphens, and all
// letters are lowercase.
//
// If the source string already has a certain format, such as camelCase,
// kebab-case, PascalCase, or snake_case, it will be correctly converted
// to kebab-case.
//
// Example usage:
//
//	scs.ToKebab("helloWorld")   // returns "hello-world"
//	scs.ToKebab("HelloWorld")   // returns "hello-world"
//	scs.ToKebab("hello_world")  // returns "hello-world"
//	scs.ToKebab("Hello-World")  // returns "hello-world"
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
//
// This function checks if the input string is in kebab-case. If it's not,
// it returns an error.
//
// If the input string is in kebab-case, it replaces all hyphens with spaces
// and converts the resulting string to camelCase. CamelCase is a naming
// convention in which the first letter of each word (including the first
// word) is capitalized, and there are no spaces or underscores between words.
//
// Note that this conversion could fail if the input string is not in
// kebab-case style. In that case, an error will be returned along with
// an empty string.
//
// Example usage:
//
//	result, err := KebabToCamel("hello-world")
//	// result: "helloWorld", err: nil
//
//	result, err := KebabToCamel("Hello-World")
//	// result: "", err: error (not kebab-case)
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
//
// This function checks if the input string is in kebab-case. If it's not,
// it returns an error.
//
// If the input string is in kebab-case, it replaces all hyphens with
// underscores, effectively converting the string from kebab-case to
// snake_case. Snake_case is a naming convention in which words are
// separated by underscores and there are no capital letters.
//
// Note that this conversion could fail if the input string is not in
// kebab-case style. In that case, an error will be returned along with
// an empty string.
//
// Example usage:
//
//	result, err := KebabToSnake("hello-world")
//	// result: "hello_world", err: nil
//
//	result, err := KebabToSnake("Hello-World")
//	// result: "", err: error (not kebab-case)
func KebabToSnake(kebab string) (string, error) {
	if !StrIsKebab(kebab) {
		return "", fmt.Errorf("value %s isn't kebab-case style", kebab)
	}

	return strings.ReplaceAll(kebab, "-", "_"), nil
}

// KebabToPascal converts a kebab-case-style string to PascalCase.
//
// This function checks if the input string is in kebab-case. If it's not,
// it returns an error.
//
// If the input string is in kebab-case, it replaces all hyphens with spaces,
// performs title casing on each word, and joins the words together to form
// a PascalCase string. PascalCase is a naming convention in which the first
// letter of each word (including the first word) is capitalized, and there
// are no underscores or spaces between words.
//
// Note that this conversion could fail if the input string is not in
// kebab-case style. In that case, an error will be returned along with
// an empty string.
//
// Example usage:
//
//	result, err := KebabToPascal("hello-world")
//	// result: "HelloWorld", err: nil
//
//	result, err := KebabToPascal("Hello-World")
//	// result: "", err: error (not kebab-case)
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
